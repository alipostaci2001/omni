package keeper

import (
	"bytes"
	"context"

	"github.com/omni-network/omni/halo2/attest/types"
	"github.com/omni-network/omni/lib/engine"
	"github.com/omni-network/omni/lib/errors"
	"github.com/omni-network/omni/lib/k1util"

	abci "github.com/cometbft/cometbft/abci/types"
	k1 "github.com/cometbft/cometbft/crypto/secp256k1"

	"github.com/ethereum/go-ethereum/common"

	ormv1alpha1 "cosmossdk.io/api/cosmos/orm/v1alpha1"
	"cosmossdk.io/core/store"
	"cosmossdk.io/orm/model/ormdb"
	"cosmossdk.io/orm/types/ormerrors"
	"github.com/cosmos/cosmos-sdk/codec"
	skeeper "github.com/cosmos/cosmos-sdk/x/staking/keeper"
)

// Keeper is the aggregate attestation keeper.
// It keeps tracks of all attestations included on-chain and detects when they are approved.
type Keeper struct {
	aggTable     AggAttestationTable
	sigTable     AttSignatureTable
	cdc          codec.BinaryCodec
	storeService store.KVStoreService
	ethCl        engine.API
	skeeper      *skeeper.Keeper // TODO(corver): Define a interface for the methods we use.
}

// NewKeeper returns a new aggregate attestation keeper.
func NewKeeper(
	cdc codec.BinaryCodec,
	storeSvc store.KVStoreService,
	ethCl engine.API,
	skeeper *skeeper.Keeper,
) (Keeper, error) {
	schema := &ormv1alpha1.ModuleSchemaDescriptor{SchemaFile: []*ormv1alpha1.ModuleSchemaDescriptor_FileEntry{
		{Id: 1, ProtoFileName: File_halo2_attest_keeper_aggregate_proto.Path()},
	}}

	modDB, err := ormdb.NewModuleDB(schema, ormdb.ModuleDBOptions{KVStoreService: storeSvc})
	if err != nil {
		return Keeper{}, errors.Wrap(err, "create module db")
	}

	aggstore, err := NewAggregateStore(modDB)
	if err != nil {
		return Keeper{}, errors.Wrap(err, "create aggregate store")
	}

	return Keeper{
		aggTable:     aggstore.AggAttestationTable(),
		sigTable:     aggstore.AttSignatureTable(),
		cdc:          cdc,
		storeService: storeSvc,
		ethCl:        ethCl,
		skeeper:      skeeper,
	}, nil
}

// Add adds the given aggregate attestations to the store.
// It merges the aggregate if it already exists.
func (k Keeper) Add(ctx context.Context, msg *types.MsgAggAttestation) error {
	header := msg.BlockHeader

	var aggID uint64
	exiting, err := k.aggTable.GetByChainIdHeightHash(ctx, header.ChainId, header.Height, header.Hash)
	if ormerrors.IsNotFound(err) {
		// Insert new aggregate
		aggID, err = k.aggTable.InsertReturningId(ctx, aggAttestationToORM(msg))
		if err != nil {
			return errors.Wrap(err, "insert")
		}
	} else if err != nil {
		return errors.Wrap(err, "by block header")
	} else if !bytes.Equal(exiting.GetBlockRoot(), msg.BlockRoot) {
		return errors.New("mismatching block root")
	} else {
		aggID = exiting.GetId()
	}

	// Insert signatures
	for _, sig := range msg.Signatures {
		err := k.sigTable.Insert(ctx, &AttSignature{
			Signature:        sig.Signature,
			ValidatorAddress: sig.ValidatorAddress,
			AggId:            aggID,
		})
		if err != nil {
			return errors.Wrap(err, "insert signature")
		}
	}

	return nil
}

// Approve approves any pending aggregate attestations that have quorum signatures form the provided set.
func (k Keeper) Approve(ctx context.Context, valSetID uint64, validators abci.ValidatorUpdates) error {
	approvedIdx := AggAttestationStatusChainIdHeightIndexKey{}.WithStatus(int32(AggStatus_Approved))
	iter, err := k.aggTable.List(ctx, approvedIdx)
	if err != nil {
		return errors.Wrap(err, "list pending")
	}
	defer iter.Close()

	for iter.Next() {
		agg, err := iter.Value()
		if err != nil {
			return errors.Wrap(err, "value")
		}

		sigs, err := k.getAggSigs(ctx, agg.GetId())
		if err != nil {
			return errors.Wrap(err, "get agg validators")
		}

		toDelete, ok := isApproved(validators, sigs)
		if !ok {
			continue
		}

		for _, sig := range toDelete {
			err := k.sigTable.Delete(ctx, sig)
			if err != nil {
				return errors.Wrap(err, "delete sig")
			}
		}

		// Update status
		agg.Status = int32(AggStatus_Approved)
		agg.ValidatorSetId = valSetID
		err = k.aggTable.Save(ctx, agg)
		if err != nil {
			return errors.Wrap(err, "save")
		}
	}

	return nil
}

// approvedFrom returns the subsequent approved attestations from the provided height (inclusive).
func (k Keeper) approvedFrom(ctx context.Context, chainID uint64, height uint64, max uint64,
) ([]*types.MsgAggAttestation, error) {
	from := AggAttestationStatusChainIdHeightIndexKey{}.WithStatusChainIdHeight(
		int32(AggStatus_Approved), chainID, height)
	to := AggAttestationStatusChainIdHeightIndexKey{}.WithStatusChainIdHeight(
		int32(AggStatus_Approved), chainID, height+max)

	iter, err := k.aggTable.ListRange(ctx, from, to)
	if err != nil {
		return nil, errors.Wrap(err, "list range")
	}
	defer iter.Close()

	var resp []*types.MsgAggAttestation
	next := height
	for iter.Next() {
		agg, err := iter.Value()
		if err != nil {
			return nil, errors.Wrap(err, "value")
		}

		if agg.GetHeight() != next {
			break
		}
		next++

		pbsigs, err := k.getAggSigs(ctx, agg.GetId())
		if err != nil {
			return nil, errors.Wrap(err, "get agg sigs")
		}

		var sigs []*types.SigTuple
		for _, pbsig := range pbsigs {
			sigs = append(sigs, &types.SigTuple{
				ValidatorAddress: pbsig.GetValidatorAddress(),
				Signature:        pbsig.GetSignature(),
			})
		}

		resp = append(resp, &types.MsgAggAttestation{
			BlockHeader: &types.BlockHeader{
				ChainId: agg.GetChainId(),
				Height:  agg.GetHeight(),
				Hash:    agg.GetHash(),
			},
			ValidatorSetId: agg.GetValidatorSetId(),
			BlockRoot:      agg.GetBlockRoot(),
			Signatures:     sigs,
		})
	}

	return resp, nil
}

// getAggSigs returns the signatures for the given aggregate ID.
func (k Keeper) getAggSigs(ctx context.Context, aggID uint64) ([]*AttSignature, error) {
	aggIDIdx := AttSignatureAggIdIndexKey{}.WithAggId(aggID)
	sigIter, err := k.sigTable.List(ctx, aggIDIdx)
	if err != nil {
		return nil, errors.Wrap(err, "list sig")
	}
	defer sigIter.Close()

	var sigs []*AttSignature
	for sigIter.Next() {
		sig, err := sigIter.Value()
		if err != nil {
			return nil, errors.Wrap(err, "value sig")
		}

		sigs = append(sigs, sig)
	}

	return sigs, nil
}

func (k Keeper) EndBlock(ctx context.Context) error {
	reduction := k.skeeper.PowerReduction(ctx)

	// TODO(corver): We need to fetch the validators for the previous block. Maybe get them from cometBFT directly.
	vals, err := k.skeeper.GetBondedValidatorsByPower(ctx)
	if err != nil {
		return errors.Wrap(err, "get last validators")
	}

	valUpdates := make(abci.ValidatorUpdates, 0, len(vals))
	for _, val := range vals {
		cosmosPubKey, err := val.ConsPubKey()
		if err != nil {
			return errors.Wrap(err, "consensus pubkey")
		}

		valUpdates = append(valUpdates, abci.UpdateValidator(
			cosmosPubKey.Bytes(),
			val.ConsensusPower(reduction),
			k1.KeyType,
		))
	}

	// TODO(corver): Figure out where this is stored and fetch it.
	var valSetID uint64 = 1

	return k.Approve(ctx, valSetID, valUpdates)
}

// isApproved returns whether the given signatures are approved by the given validators.
// It also returns the signatures to delete (not in the validator set).
func isApproved(validators abci.ValidatorUpdates, sigs []*AttSignature) ([]*AttSignature, bool) {
	valSet := make(map[common.Address]int64)
	var total int64
	for _, val := range validators {
		addr, err := k1util.PubKeyPBToAddress(val.PubKey)
		if err != nil {
			return nil, false
		}

		total += val.Power
		valSet[addr] = val.Power
	}

	var sum int64
	var toDelete []*AttSignature
	for _, sig := range sigs {
		power, ok := valSet[common.Address(sig.GetValidatorAddress())]
		if !ok {
			toDelete = append(toDelete, sig)
			continue
		}

		sum += power
	}

	return toDelete, sum > total*2/3
}