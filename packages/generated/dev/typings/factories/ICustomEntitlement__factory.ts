/* Autogenerated file. Do not edit manually. */
/* tslint:disable */
/* eslint-disable */

import { Contract, Signer, utils } from "ethers";
import type { Provider } from "@ethersproject/providers";
import type {
  ICustomEntitlement,
  ICustomEntitlementInterface,
} from "../ICustomEntitlement";

const _abi = [
  {
    type: "function",
    name: "isEntitled",
    inputs: [
      {
        name: "user",
        type: "address[]",
        internalType: "address[]",
      },
    ],
    outputs: [
      {
        name: "",
        type: "bool",
        internalType: "bool",
      },
    ],
    stateMutability: "view",
  },
  {
    type: "function",
    name: "supportsInterface",
    inputs: [
      {
        name: "interfaceId",
        type: "bytes4",
        internalType: "bytes4",
      },
    ],
    outputs: [
      {
        name: "",
        type: "bool",
        internalType: "bool",
      },
    ],
    stateMutability: "view",
  },
] as const;

export class ICustomEntitlement__factory {
  static readonly abi = _abi;
  static createInterface(): ICustomEntitlementInterface {
    return new utils.Interface(_abi) as ICustomEntitlementInterface;
  }
  static connect(
    address: string,
    signerOrProvider: Signer | Provider
  ): ICustomEntitlement {
    return new Contract(address, _abi, signerOrProvider) as ICustomEntitlement;
  }
}
