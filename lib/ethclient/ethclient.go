package ethclient

import (
	"context"

	"github.com/omni-network/omni/lib/errors"

	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/ethereum/go-ethereum/rpc"
)

//go:generate go run genwrap/genwrap.go

var _ Client = Wrapper{}

type HeadType string

func (h HeadType) String() string {
	return string(h)
}

const (
	HeadLatest    HeadType = "latest"
	HeadEarliest  HeadType = "earliest"
	HeadPending   HeadType = "pending"
	HeadSafe      HeadType = "safe"
	HeadFinalized HeadType = "finalized"
)

// Wrapper wraps an ethclient.Client adding metrics and wrapped errors.
type Wrapper struct {
	cl    *ethclient.Client
	chain string
}

// NewClient wraps an *rpc.Client adding metrics and wrapped errors.
func NewClient(cl *rpc.Client, chain string) Wrapper {
	return Wrapper{
		cl:    ethclient.NewClient(cl),
		chain: chain,
	}
}

// Dial connects a client to the given URL.
//
// Note if the URL is http(s), it doesn't return an error if it cannot connect to the URL.
// It will retry connecting on every call to a wrapped method. It will only return an error if the
// url is invalid.
func Dial(chainName string, url string) (Wrapper, error) {
	cl, err := ethclient.Dial(url)
	if err != nil {
		return Wrapper{}, errors.Wrap(err, "dial", "chain", chainName, "url", url)
	}

	return Wrapper{
		cl:    cl,
		chain: chainName,
	}, nil
}

// Close closes the underlying RPC connection.
func (w Wrapper) Close() {
	w.cl.Close()
}

// HeaderByType returns the block header for the given head type.
func (w Wrapper) HeaderByType(ctx context.Context, typ HeadType) (*types.Header, error) {
	const endpoint = "eth_getBlockByNumber"
	defer latency(w.chain, endpoint)() //nolint:revive // Defer chain is fine here.

	var header *types.Header
	err := w.cl.Client().CallContext(
		ctx,
		&header,
		endpoint,
		typ.String(),
		false,
	)
	if err != nil {
		incError(w.chain, endpoint)
		return nil, errors.Wrap(err, "get block")
	}

	return header, nil
}