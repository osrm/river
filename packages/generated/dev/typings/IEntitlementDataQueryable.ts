/* Autogenerated file. Do not edit manually. */
/* tslint:disable */
/* eslint-disable */
import type {
  BaseContract,
  BigNumber,
  BytesLike,
  CallOverrides,
  PopulatedTransaction,
  Signer,
  utils,
} from "ethers";
import type { FunctionFragment, Result } from "@ethersproject/abi";
import type { Listener, Provider } from "@ethersproject/providers";
import type {
  TypedEventFilter,
  TypedEvent,
  TypedListener,
  OnEvent,
  PromiseOrValue,
} from "./common";

export declare namespace IEntitlementDataQueryableBase {
  export type EntitlementDataStruct = {
    entitlementType: PromiseOrValue<string>;
    entitlementData: PromiseOrValue<BytesLike>;
  };

  export type EntitlementDataStructOutput = [string, string] & {
    entitlementType: string;
    entitlementData: string;
  };
}

export interface IEntitlementDataQueryableInterface extends utils.Interface {
  functions: {
    "getChannelEntitlementDataByPermission(bytes32,string)": FunctionFragment;
    "getEntitlementDataByPermission(string)": FunctionFragment;
  };

  getFunction(
    nameOrSignatureOrTopic:
      | "getChannelEntitlementDataByPermission"
      | "getEntitlementDataByPermission"
  ): FunctionFragment;

  encodeFunctionData(
    functionFragment: "getChannelEntitlementDataByPermission",
    values: [PromiseOrValue<BytesLike>, PromiseOrValue<string>]
  ): string;
  encodeFunctionData(
    functionFragment: "getEntitlementDataByPermission",
    values: [PromiseOrValue<string>]
  ): string;

  decodeFunctionResult(
    functionFragment: "getChannelEntitlementDataByPermission",
    data: BytesLike
  ): Result;
  decodeFunctionResult(
    functionFragment: "getEntitlementDataByPermission",
    data: BytesLike
  ): Result;

  events: {};
}

export interface IEntitlementDataQueryable extends BaseContract {
  connect(signerOrProvider: Signer | Provider | string): this;
  attach(addressOrName: string): this;
  deployed(): Promise<this>;

  interface: IEntitlementDataQueryableInterface;

  queryFilter<TEvent extends TypedEvent>(
    event: TypedEventFilter<TEvent>,
    fromBlockOrBlockhash?: string | number | undefined,
    toBlock?: string | number | undefined
  ): Promise<Array<TEvent>>;

  listeners<TEvent extends TypedEvent>(
    eventFilter?: TypedEventFilter<TEvent>
  ): Array<TypedListener<TEvent>>;
  listeners(eventName?: string): Array<Listener>;
  removeAllListeners<TEvent extends TypedEvent>(
    eventFilter: TypedEventFilter<TEvent>
  ): this;
  removeAllListeners(eventName?: string): this;
  off: OnEvent<this>;
  on: OnEvent<this>;
  once: OnEvent<this>;
  removeListener: OnEvent<this>;

  functions: {
    getChannelEntitlementDataByPermission(
      channelId: PromiseOrValue<BytesLike>,
      permission: PromiseOrValue<string>,
      overrides?: CallOverrides
    ): Promise<[IEntitlementDataQueryableBase.EntitlementDataStructOutput[]]>;

    getEntitlementDataByPermission(
      permission: PromiseOrValue<string>,
      overrides?: CallOverrides
    ): Promise<[IEntitlementDataQueryableBase.EntitlementDataStructOutput[]]>;
  };

  getChannelEntitlementDataByPermission(
    channelId: PromiseOrValue<BytesLike>,
    permission: PromiseOrValue<string>,
    overrides?: CallOverrides
  ): Promise<IEntitlementDataQueryableBase.EntitlementDataStructOutput[]>;

  getEntitlementDataByPermission(
    permission: PromiseOrValue<string>,
    overrides?: CallOverrides
  ): Promise<IEntitlementDataQueryableBase.EntitlementDataStructOutput[]>;

  callStatic: {
    getChannelEntitlementDataByPermission(
      channelId: PromiseOrValue<BytesLike>,
      permission: PromiseOrValue<string>,
      overrides?: CallOverrides
    ): Promise<IEntitlementDataQueryableBase.EntitlementDataStructOutput[]>;

    getEntitlementDataByPermission(
      permission: PromiseOrValue<string>,
      overrides?: CallOverrides
    ): Promise<IEntitlementDataQueryableBase.EntitlementDataStructOutput[]>;
  };

  filters: {};

  estimateGas: {
    getChannelEntitlementDataByPermission(
      channelId: PromiseOrValue<BytesLike>,
      permission: PromiseOrValue<string>,
      overrides?: CallOverrides
    ): Promise<BigNumber>;

    getEntitlementDataByPermission(
      permission: PromiseOrValue<string>,
      overrides?: CallOverrides
    ): Promise<BigNumber>;
  };

  populateTransaction: {
    getChannelEntitlementDataByPermission(
      channelId: PromiseOrValue<BytesLike>,
      permission: PromiseOrValue<string>,
      overrides?: CallOverrides
    ): Promise<PopulatedTransaction>;

    getEntitlementDataByPermission(
      permission: PromiseOrValue<string>,
      overrides?: CallOverrides
    ): Promise<PopulatedTransaction>;
  };
}