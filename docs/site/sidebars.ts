import type { SidebarsConfig } from "@docusaurus/plugin-content-docs";

const sidebars: SidebarsConfig = {
  learnSidebar: [
    {
      type: "autogenerated",
      dirName: "learn",
    },
    {
      type: "html",
      value: "<div class='sidebar-separator'></div>",
    },
    {
      type: "doc",
      id: "resources/glossary",
    },
  ],
  protocolSidebar: [
    "protocol/overview",
    {
      type: "html",
      value: "<div class='sidebar-separator'></div>",
    },
    {
      type: "category",
      label: "Architecture",
      className: "sidebar-title",
      collapsible: false,
      items: [
        {
          type: "autogenerated",
          dirName: "protocol/architecture",
        }
      ]
    },
    {
      type: "html",
      value: "<div class='sidebar-separator'></div>",
    },
    {
      type: "category",
      label: "Restaking",
      className: "sidebar-title",
      collapsible: false,
      items: [
        {
          type: "autogenerated",
          dirName: "protocol/restaking",
        }
      ]
    },
    {
      type: "html",
      value: "<div class='sidebar-separator'></div>",
    },
    "protocol/future"
  ],
  developSidebar: [
    "develop/introduction",
    {
      type: "html",
      value: "<div class='sidebar-separator'></div>",
    },
    {
      type: "category",
      label: "XApp",
      className: "sidebar-title",
      collapsible: false,
      items: [
        {
          type: "autogenerated",
          dirName: "develop/xapp",
        }
      ]
    },
    {
      type: "html",
      value: "<div class='sidebar-separator'></div>",
    },
    "develop/contracts",
  ],
  operateSidebar: [
    {
      type: "autogenerated",
      dirName: "operate",
    },
  ],
};

export default sidebars;
