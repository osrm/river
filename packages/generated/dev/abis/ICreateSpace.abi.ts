export default [
  {
    "type": "function",
    "name": "createSpace",
    "inputs": [
      {
        "name": "SpaceInfo",
        "type": "tuple",
        "internalType": "struct IArchitectBase.SpaceInfo",
        "components": [
          {
            "name": "name",
            "type": "string",
            "internalType": "string"
          },
          {
            "name": "uri",
            "type": "string",
            "internalType": "string"
          },
          {
            "name": "shortDescription",
            "type": "string",
            "internalType": "string"
          },
          {
            "name": "longDescription",
            "type": "string",
            "internalType": "string"
          },
          {
            "name": "membership",
            "type": "tuple",
            "internalType": "struct IArchitectBase.Membership",
            "components": [
              {
                "name": "settings",
                "type": "tuple",
                "internalType": "struct IMembershipBase.Membership",
                "components": [
                  {
                    "name": "name",
                    "type": "string",
                    "internalType": "string"
                  },
                  {
                    "name": "symbol",
                    "type": "string",
                    "internalType": "string"
                  },
                  {
                    "name": "price",
                    "type": "uint256",
                    "internalType": "uint256"
                  },
                  {
                    "name": "maxSupply",
                    "type": "uint256",
                    "internalType": "uint256"
                  },
                  {
                    "name": "duration",
                    "type": "uint64",
                    "internalType": "uint64"
                  },
                  {
                    "name": "currency",
                    "type": "address",
                    "internalType": "address"
                  },
                  {
                    "name": "feeRecipient",
                    "type": "address",
                    "internalType": "address"
                  },
                  {
                    "name": "freeAllocation",
                    "type": "uint256",
                    "internalType": "uint256"
                  },
                  {
                    "name": "pricingModule",
                    "type": "address",
                    "internalType": "address"
                  }
                ]
              },
              {
                "name": "requirements",
                "type": "tuple",
                "internalType": "struct IArchitectBase.MembershipRequirements",
                "components": [
                  {
                    "name": "everyone",
                    "type": "bool",
                    "internalType": "bool"
                  },
                  {
                    "name": "users",
                    "type": "address[]",
                    "internalType": "address[]"
                  },
                  {
                    "name": "ruleData",
                    "type": "bytes",
                    "internalType": "bytes"
                  },
                  {
                    "name": "syncEntitlements",
                    "type": "bool",
                    "internalType": "bool"
                  }
                ]
              },
              {
                "name": "permissions",
                "type": "string[]",
                "internalType": "string[]"
              }
            ]
          },
          {
            "name": "channel",
            "type": "tuple",
            "internalType": "struct IArchitectBase.ChannelInfo",
            "components": [
              {
                "name": "metadata",
                "type": "string",
                "internalType": "string"
              }
            ]
          }
        ]
      }
    ],
    "outputs": [
      {
        "name": "",
        "type": "address",
        "internalType": "address"
      }
    ],
    "stateMutability": "nonpayable"
  },
  {
    "type": "function",
    "name": "createSpaceWithPrepay",
    "inputs": [
      {
        "name": "createSpace",
        "type": "tuple",
        "internalType": "struct IArchitectBase.CreateSpace",
        "components": [
          {
            "name": "metadata",
            "type": "tuple",
            "internalType": "struct IArchitectBase.Metadata",
            "components": [
              {
                "name": "name",
                "type": "string",
                "internalType": "string"
              },
              {
                "name": "uri",
                "type": "string",
                "internalType": "string"
              },
              {
                "name": "shortDescription",
                "type": "string",
                "internalType": "string"
              },
              {
                "name": "longDescription",
                "type": "string",
                "internalType": "string"
              }
            ]
          },
          {
            "name": "membership",
            "type": "tuple",
            "internalType": "struct IArchitectBase.Membership",
            "components": [
              {
                "name": "settings",
                "type": "tuple",
                "internalType": "struct IMembershipBase.Membership",
                "components": [
                  {
                    "name": "name",
                    "type": "string",
                    "internalType": "string"
                  },
                  {
                    "name": "symbol",
                    "type": "string",
                    "internalType": "string"
                  },
                  {
                    "name": "price",
                    "type": "uint256",
                    "internalType": "uint256"
                  },
                  {
                    "name": "maxSupply",
                    "type": "uint256",
                    "internalType": "uint256"
                  },
                  {
                    "name": "duration",
                    "type": "uint64",
                    "internalType": "uint64"
                  },
                  {
                    "name": "currency",
                    "type": "address",
                    "internalType": "address"
                  },
                  {
                    "name": "feeRecipient",
                    "type": "address",
                    "internalType": "address"
                  },
                  {
                    "name": "freeAllocation",
                    "type": "uint256",
                    "internalType": "uint256"
                  },
                  {
                    "name": "pricingModule",
                    "type": "address",
                    "internalType": "address"
                  }
                ]
              },
              {
                "name": "requirements",
                "type": "tuple",
                "internalType": "struct IArchitectBase.MembershipRequirements",
                "components": [
                  {
                    "name": "everyone",
                    "type": "bool",
                    "internalType": "bool"
                  },
                  {
                    "name": "users",
                    "type": "address[]",
                    "internalType": "address[]"
                  },
                  {
                    "name": "ruleData",
                    "type": "bytes",
                    "internalType": "bytes"
                  },
                  {
                    "name": "syncEntitlements",
                    "type": "bool",
                    "internalType": "bool"
                  }
                ]
              },
              {
                "name": "permissions",
                "type": "string[]",
                "internalType": "string[]"
              }
            ]
          },
          {
            "name": "channel",
            "type": "tuple",
            "internalType": "struct IArchitectBase.ChannelInfo",
            "components": [
              {
                "name": "metadata",
                "type": "string",
                "internalType": "string"
              }
            ]
          },
          {
            "name": "prepay",
            "type": "tuple",
            "internalType": "struct IArchitectBase.Prepay",
            "components": [
              {
                "name": "supply",
                "type": "uint256",
                "internalType": "uint256"
              }
            ]
          }
        ]
      }
    ],
    "outputs": [
      {
        "name": "",
        "type": "address",
        "internalType": "address"
      }
    ],
    "stateMutability": "payable"
  },
  {
    "type": "event",
    "name": "Architect__ProxyInitializerSet",
    "inputs": [
      {
        "name": "proxyInitializer",
        "type": "address",
        "indexed": true,
        "internalType": "address"
      }
    ],
    "anonymous": false
  },
  {
    "type": "event",
    "name": "SpaceCreated",
    "inputs": [
      {
        "name": "owner",
        "type": "address",
        "indexed": true,
        "internalType": "address"
      },
      {
        "name": "tokenId",
        "type": "uint256",
        "indexed": true,
        "internalType": "uint256"
      },
      {
        "name": "space",
        "type": "address",
        "indexed": true,
        "internalType": "address"
      }
    ],
    "anonymous": false
  },
  {
    "type": "error",
    "name": "Architect__InvalidAddress",
    "inputs": []
  },
  {
    "type": "error",
    "name": "Architect__InvalidNetworkId",
    "inputs": []
  },
  {
    "type": "error",
    "name": "Architect__InvalidPricingModule",
    "inputs": []
  },
  {
    "type": "error",
    "name": "Architect__InvalidStringLength",
    "inputs": []
  },
  {
    "type": "error",
    "name": "Architect__NotContract",
    "inputs": []
  }
] as const