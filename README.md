# Go API client for waas2

The Cobo Wallet-as-a-Service (WaaS) 2.0 API is the latest version of Cobo’s WaaS API offering. It enables you to access Cobo’s full suite of crypto wallet technologies with powerful and flexible access controls. By encapsulating complex security protocols and streamlining blockchain interactions, this API allows you to concentrate on your core business activities without worrying about the safety of your assets. The WaaS 2.0 API presents the following key features:

- A unified API for Cobo’s [all four wallet types](https://manuals.cobo.com/en/portal/introduction#an-all-in-one-wallet-platform)
- Support for 80+ chains and 3000+ tokens
- A comprehensive selection of webhook events
- Flexible usage models for MPC wallets, including [Organization-Controlled Wallets](https://manuals.cobo.com/en/portal/mpc-wallets/ocw/introduction) and [User-Controlled Wallets](https://manuals.cobo.com/en/portal/mpc-wallets/ucw/introduction)
- Programmatic control of smart contract wallets such as Safe\\{Wallet\\} with fine-grained access controls
- Seamlessly transfer funds across multiple exchanges, including Binance, OKX, Bybit, Deribit, and more

For more information about the WaaS 2.0 API, see [Introduction to WaaS 2.0](https://www.cobo.com/developers/v2/guides/overview/introduction).


## Overview
This API client was generated by the [OpenAPI Generator](https://openapi-generator.tech) project.  By using the [OpenAPI-spec](https://www.openapis.org/) from a remote server, you can easily generate an API client.

- API version: v2
- Package version: 1.0.0
- Build package: org.openapitools.codegen.languages.GoClientCodegen
For more information, please visit [https://www.cobo.com/waas](https://www.cobo.com/waas)

## Installation

Install the following dependencies:

```sh
go get github.com/stretchr/testify/assert
go get golang.org/x/oauth2
go get golang.org/x/net/context
```

Put the package under your project folder and add the following in import:

```go
import waas2 "github.com/CoboGlobal/cobo-waas2-go-sdk"
```

To use a proxy, set the environment variable `HTTP_PROXY`:

```go
os.Setenv("HTTP_PROXY", "http://proxy_name:proxy_port")
```

## Configuration of Server URL

Default configuration comes with `Servers` field that contains server objects as defined in the OpenAPI specification.

### Select Server Configuration

For using Dev environment set context value `waas2.ContextEnv` of type `int`.

```go
ctx := context.WithValue(context.Background(), waas2.ContextEnv, waas2.DevEnv)
```

For using Prod environment set context value `waas2.ContextEnv` of type `int`.

```go
ctx := context.WithValue(context.Background(), waas2.ContextEnv, waas2.ProdEnv)
```

you can also set server URL using context values `waas2.ContextServerHost` of type `string`

```go
ctx := context.WithValue(context.Background(), waas2.ContextServerHost, "https://api[.xxxx].cobo.com/v2")
```

## Set Api Signer
```go
ctx = context.WithValue(ctx, waas2.ContextPortalSigner, crypto.Ed25519Signer{
    Secret: "<YOUR_API_PRIV_KEY_IN_HEX>",
})
```

## Documentation for API Endpoints

All URIs are relative to *https://api.dev.cobo.com/v2*

Class | Method | HTTP request | Description
------------ | ------------- | ------------- | -------------
*DevelopersWebhooksAPI* | [**CreateWebhookEndpoint**](docs/DevelopersWebhooksAPI.md#createwebhookendpoint) | **Post** /webhooks/endpoints | Register webhook endpoint
*DevelopersWebhooksAPI* | [**GetWebhookEndpointById**](docs/DevelopersWebhooksAPI.md#getwebhookendpointbyid) | **Get** /webhooks/endpoints/{endpoint_id} | Get webhook endpoint information
*DevelopersWebhooksAPI* | [**GetWebhookEventById**](docs/DevelopersWebhooksAPI.md#getwebhookeventbyid) | **Get** /webhooks/endpoints/{endpoint_id}/events/{event_id} | Retrieve event information
*DevelopersWebhooksAPI* | [**ListWebhookEndpoints**](docs/DevelopersWebhooksAPI.md#listwebhookendpoints) | **Get** /webhooks/endpoints | List webhook endpoints
*DevelopersWebhooksAPI* | [**ListWebhookEventDefinitions**](docs/DevelopersWebhooksAPI.md#listwebhookeventdefinitions) | **Get** /webhooks/events/definitions | Get webhook event types
*DevelopersWebhooksAPI* | [**ListWebhookEventLogs**](docs/DevelopersWebhooksAPI.md#listwebhookeventlogs) | **Get** /webhooks/endpoints/{endpoint_id}/events/{event_id}/logs | List webhook event logs
*DevelopersWebhooksAPI* | [**ListWebhookEvents**](docs/DevelopersWebhooksAPI.md#listwebhookevents) | **Get** /webhooks/endpoints/{endpoint_id}/events | List all webhook events
*DevelopersWebhooksAPI* | [**RetryWebhookEventById**](docs/DevelopersWebhooksAPI.md#retrywebhookeventbyid) | **Post** /webhooks/endpoints/{endpoint_id}/events/{event_id}/retry | Retry event
*DevelopersWebhooksAPI* | [**UpdateWebhookEndpointById**](docs/DevelopersWebhooksAPI.md#updatewebhookendpointbyid) | **Put** /webhooks/endpoints/{endpoint_id} | Update webhook endpoint
*TransactionsAPI* | [**CancelTransactionById**](docs/TransactionsAPI.md#canceltransactionbyid) | **Post** /transactions/{transaction_id}/cancel | Cancel transaction
*TransactionsAPI* | [**CreateContractCallTransaction**](docs/TransactionsAPI.md#createcontractcalltransaction) | **Post** /transactions/contract_call | Call smart contract
*TransactionsAPI* | [**CreateMessageSignTransaction**](docs/TransactionsAPI.md#createmessagesigntransaction) | **Post** /transactions/message_sign | Sign message
*TransactionsAPI* | [**CreateTransferTransaction**](docs/TransactionsAPI.md#createtransfertransaction) | **Post** /transactions/transfer | Transfer token
*TransactionsAPI* | [**DropTransactionById**](docs/TransactionsAPI.md#droptransactionbyid) | **Post** /transactions/{transaction_id}/drop | Drop transaction
*TransactionsAPI* | [**EstimateFee**](docs/TransactionsAPI.md#estimatefee) | **Post** /transactions/estimate_fee | Estimate transaction fee
*TransactionsAPI* | [**GetTransactionById**](docs/TransactionsAPI.md#gettransactionbyid) | **Get** /transactions/{transaction_id} | Get transaction information
*TransactionsAPI* | [**ListTransactions**](docs/TransactionsAPI.md#listtransactions) | **Get** /transactions | List all transactions
*TransactionsAPI* | [**ResendTransactionById**](docs/TransactionsAPI.md#resendtransactionbyid) | **Post** /transactions/{transaction_id}/resend | Resend transaction
*TransactionsAPI* | [**SpeedupTransactionById**](docs/TransactionsAPI.md#speeduptransactionbyid) | **Post** /transactions/{transaction_id}/speedup | Speed up transaction
*WalletsAPI* | [**CheckAddressValidity**](docs/WalletsAPI.md#checkaddressvalidity) | **Get** /wallets/check_address_validity | Check address validity
*WalletsAPI* | [**CreateAddress**](docs/WalletsAPI.md#createaddress) | **Post** /wallets/{wallet_id}/addresses | Create addresses in wallet
*WalletsAPI* | [**CreateWallet**](docs/WalletsAPI.md#createwallet) | **Post** /wallets | Create wallet
*WalletsAPI* | [**DeleteWalletById**](docs/WalletsAPI.md#deletewalletbyid) | **Post** /wallets/{wallet_id}/delete | Delete wallet
*WalletsAPI* | [**GetAddress**](docs/WalletsAPI.md#getaddress) | **Get** /wallets/{wallet_id}/addresses/{address} | Get address information
*WalletsAPI* | [**GetChainById**](docs/WalletsAPI.md#getchainbyid) | **Get** /wallets/chains/{chain_id} | Get chain information
*WalletsAPI* | [**GetMaxTransferableValue**](docs/WalletsAPI.md#getmaxtransferablevalue) | **Get** /wallets/{wallet_id}/max_transferable_value | Get maximum transferable value
*WalletsAPI* | [**GetTokenById**](docs/WalletsAPI.md#gettokenbyid) | **Get** /wallets/tokens/{token_id} | Get token information
*WalletsAPI* | [**GetWalletById**](docs/WalletsAPI.md#getwalletbyid) | **Get** /wallets/{wallet_id} | Get wallet information
*WalletsAPI* | [**ListAddresses**](docs/WalletsAPI.md#listaddresses) | **Get** /wallets/{wallet_id}/addresses | List wallet addresses
*WalletsAPI* | [**ListEnabledChains**](docs/WalletsAPI.md#listenabledchains) | **Get** /wallets/enabled_chains | List enabled chains
*WalletsAPI* | [**ListEnabledTokens**](docs/WalletsAPI.md#listenabledtokens) | **Get** /wallets/enabled_tokens | List enabled tokens
*WalletsAPI* | [**ListSupportedChains**](docs/WalletsAPI.md#listsupportedchains) | **Get** /wallets/chains | List supported chains
*WalletsAPI* | [**ListSupportedTokens**](docs/WalletsAPI.md#listsupportedtokens) | **Get** /wallets/tokens | List supported tokens
*WalletsAPI* | [**ListTokenBalancesForAddress**](docs/WalletsAPI.md#listtokenbalancesforaddress) | **Get** /wallets/{wallet_id}/addresses/{address}/tokens | List token balances by address
*WalletsAPI* | [**ListTokenBalancesForWallet**](docs/WalletsAPI.md#listtokenbalancesforwallet) | **Get** /wallets/{wallet_id}/tokens | List token balances by wallet
*WalletsAPI* | [**ListUtxos**](docs/WalletsAPI.md#listutxos) | **Get** /wallets/{wallet_id}/utxos | List UTXOs
*WalletsAPI* | [**ListWallets**](docs/WalletsAPI.md#listwallets) | **Get** /wallets | List all wallets
*WalletsAPI* | [**LockUtxos**](docs/WalletsAPI.md#lockutxos) | **Post** /wallets/{wallet_id}/utxos/lock | Lock UTXOs
*WalletsAPI* | [**UnlockUtxos**](docs/WalletsAPI.md#unlockutxos) | **Post** /wallets/{wallet_id}/utxos/unlock | Unlock UTXOs
*WalletsAPI* | [**UpdateWalletById**](docs/WalletsAPI.md#updatewalletbyid) | **Put** /wallets/{wallet_id} | Update wallet
*WalletsMPCWalletsAPI* | [**CancelTssRequestById**](docs/WalletsMPCWalletsAPI.md#canceltssrequestbyid) | **Post** /wallets/mpc/vaults/{vault_id}/tss_requests/{tss_request_id}/cancel | Cancel TSS request
*WalletsMPCWalletsAPI* | [**CreateKeyShareHolderGroup**](docs/WalletsMPCWalletsAPI.md#createkeyshareholdergroup) | **Post** /wallets/mpc/vaults/{vault_id}/key_share_holder_groups | Create key share holder group
*WalletsMPCWalletsAPI* | [**CreateMpcProject**](docs/WalletsMPCWalletsAPI.md#creatempcproject) | **Post** /wallets/mpc/projects | Create project
*WalletsMPCWalletsAPI* | [**CreateMpcVault**](docs/WalletsMPCWalletsAPI.md#creatempcvault) | **Post** /wallets/mpc/vaults | Create vault
*WalletsMPCWalletsAPI* | [**CreateTssRequest**](docs/WalletsMPCWalletsAPI.md#createtssrequest) | **Post** /wallets/mpc/vaults/{vault_id}/tss_requests | Create TSS request
*WalletsMPCWalletsAPI* | [**DeleteKeyShareHolderGroupById**](docs/WalletsMPCWalletsAPI.md#deletekeyshareholdergroupbyid) | **Post** /wallets/mpc/vaults/{vault_id}/key_share_holder_groups/{key_share_holder_group_id}/delete | Delete key share holder group
*WalletsMPCWalletsAPI* | [**GetKeyShareHolderGroupById**](docs/WalletsMPCWalletsAPI.md#getkeyshareholdergroupbyid) | **Get** /wallets/mpc/vaults/{vault_id}/key_share_holder_groups/{key_share_holder_group_id} | Get key share holder group information
*WalletsMPCWalletsAPI* | [**GetMpcProjectById**](docs/WalletsMPCWalletsAPI.md#getmpcprojectbyid) | **Get** /wallets/mpc/projects/{project_id} | Get project information
*WalletsMPCWalletsAPI* | [**GetMpcVaultById**](docs/WalletsMPCWalletsAPI.md#getmpcvaultbyid) | **Get** /wallets/mpc/vaults/{vault_id} | Get vault information
*WalletsMPCWalletsAPI* | [**GetTssRequestById**](docs/WalletsMPCWalletsAPI.md#gettssrequestbyid) | **Get** /wallets/mpc/vaults/{vault_id}/tss_requests/{tss_request_id} | Get TSS request
*WalletsMPCWalletsAPI* | [**ListCoboKeyHolders**](docs/WalletsMPCWalletsAPI.md#listcobokeyholders) | **Get** /wallets/mpc/cobo_key_share_holders | List all Cobo key share holders
*WalletsMPCWalletsAPI* | [**ListKeyShareHolderGroups**](docs/WalletsMPCWalletsAPI.md#listkeyshareholdergroups) | **Get** /wallets/mpc/vaults/{vault_id}/key_share_holder_groups | List all key share holder groups
*WalletsMPCWalletsAPI* | [**ListMpcProjects**](docs/WalletsMPCWalletsAPI.md#listmpcprojects) | **Get** /wallets/mpc/projects | List all projects
*WalletsMPCWalletsAPI* | [**ListMpcVaults**](docs/WalletsMPCWalletsAPI.md#listmpcvaults) | **Get** /wallets/mpc/vaults | List all vaults
*WalletsMPCWalletsAPI* | [**ListTssRequests**](docs/WalletsMPCWalletsAPI.md#listtssrequests) | **Get** /wallets/mpc/vaults/{vault_id}/tss_requests | List TSS requests
*WalletsMPCWalletsAPI* | [**UpdateKeyShareHolderGroupById**](docs/WalletsMPCWalletsAPI.md#updatekeyshareholdergroupbyid) | **Put** /wallets/mpc/vaults/{vault_id}/key_share_holder_groups/{key_share_holder_group_id} | Update key share holder group
*WalletsMPCWalletsAPI* | [**UpdateMpcProjectById**](docs/WalletsMPCWalletsAPI.md#updatempcprojectbyid) | **Put** /wallets/mpc/projects/{project_id} | Update project name
*WalletsMPCWalletsAPI* | [**UpdateMpcVaultById**](docs/WalletsMPCWalletsAPI.md#updatempcvaultbyid) | **Put** /wallets/mpc/vaults/{vault_id} | Update vault name


## Documentation For Models

 - [Activity](docs/Activity.md)
 - [ActivityAction](docs/ActivityAction.md)
 - [ActivityInitiator](docs/ActivityInitiator.md)
 - [ActivityStatus](docs/ActivityStatus.md)
 - [ActivityTimeline](docs/ActivityTimeline.md)
 - [ActivityType](docs/ActivityType.md)
 - [AddressEncoding](docs/AddressEncoding.md)
 - [AddressInfo](docs/AddressInfo.md)
 - [AddressTransferDestination](docs/AddressTransferDestination.md)
 - [AddressTransferDestinationAccountOutput](docs/AddressTransferDestinationAccountOutput.md)
 - [AddressTransferDestinationUtxoOutputsInner](docs/AddressTransferDestinationUtxoOutputsInner.md)
 - [AmountDetailsInner](docs/AmountDetailsInner.md)
 - [AmountStatus](docs/AmountStatus.md)
 - [AssetBalance](docs/AssetBalance.md)
 - [AssetInfo](docs/AssetInfo.md)
 - [BabylonStakeExtra](docs/BabylonStakeExtra.md)
 - [BabylonValidator](docs/BabylonValidator.md)
 - [BaseContractCallSource](docs/BaseContractCallSource.md)
 - [BaseEstimateStakingFee](docs/BaseEstimateStakingFee.md)
 - [BaseStakeExtra](docs/BaseStakeExtra.md)
 - [ChainInfo](docs/ChainInfo.md)
 - [CheckAddressValidity200Response](docs/CheckAddressValidity200Response.md)
 - [CoboSafeDelegate](docs/CoboSafeDelegate.md)
 - [CoboSafeDelegateType](docs/CoboSafeDelegateType.md)
 - [ContractCallDestination](docs/ContractCallDestination.md)
 - [ContractCallDestinationType](docs/ContractCallDestinationType.md)
 - [ContractCallParams](docs/ContractCallParams.md)
 - [ContractCallSource](docs/ContractCallSource.md)
 - [ContractCallSourceType](docs/ContractCallSourceType.md)
 - [CreateAddressRequest](docs/CreateAddressRequest.md)
 - [CreateCustodialWalletParams](docs/CreateCustodialWalletParams.md)
 - [CreateExchangeWalletParams](docs/CreateExchangeWalletParams.md)
 - [CreateKeyShareHolder](docs/CreateKeyShareHolder.md)
 - [CreateKeyShareHolderGroupRequest](docs/CreateKeyShareHolderGroupRequest.md)
 - [CreateMpcProjectRequest](docs/CreateMpcProjectRequest.md)
 - [CreateMpcVaultRequest](docs/CreateMpcVaultRequest.md)
 - [CreateMpcWalletParams](docs/CreateMpcWalletParams.md)
 - [CreateSafeWalletParams](docs/CreateSafeWalletParams.md)
 - [CreateSmartContractWalletParams](docs/CreateSmartContractWalletParams.md)
 - [CreateStakeActivity](docs/CreateStakeActivity.md)
 - [CreateStakeActivityExtra](docs/CreateStakeActivityExtra.md)
 - [CreateTransferTransaction201Response](docs/CreateTransferTransaction201Response.md)
 - [CreateTssRequestRequest](docs/CreateTssRequestRequest.md)
 - [CreateUnstakeActivity](docs/CreateUnstakeActivity.md)
 - [CreateWalletParams](docs/CreateWalletParams.md)
 - [CreateWebhookEndpointRequest](docs/CreateWebhookEndpointRequest.md)
 - [CreateWithdrawActivity](docs/CreateWithdrawActivity.md)
 - [CreatedWalletInfo](docs/CreatedWalletInfo.md)
 - [CurveType](docs/CurveType.md)
 - [CustodialTransferSource](docs/CustodialTransferSource.md)
 - [CustodialWalletInfo](docs/CustodialWalletInfo.md)
 - [DeleteKeyShareHolderGroupById201Response](docs/DeleteKeyShareHolderGroupById201Response.md)
 - [DeleteWalletById201Response](docs/DeleteWalletById201Response.md)
 - [EigenLayerLstStakeExtra](docs/EigenLayerLstStakeExtra.md)
 - [EigenLayerNativeStakeExtra](docs/EigenLayerNativeStakeExtra.md)
 - [EigenlayerValidator](docs/EigenlayerValidator.md)
 - [ErrorResponse](docs/ErrorResponse.md)
 - [EstimateContractCallFeeParams](docs/EstimateContractCallFeeParams.md)
 - [EstimateFeeParams](docs/EstimateFeeParams.md)
 - [EstimateFeeRequestType](docs/EstimateFeeRequestType.md)
 - [EstimateStakeFee](docs/EstimateStakeFee.md)
 - [EstimateTransferFeeParams](docs/EstimateTransferFeeParams.md)
 - [EstimateUnstakeFee](docs/EstimateUnstakeFee.md)
 - [EstimateWithdrawFee](docs/EstimateWithdrawFee.md)
 - [EstimatedEvmEip1559Fee](docs/EstimatedEvmEip1559Fee.md)
 - [EstimatedEvmEip1559FeeSlow](docs/EstimatedEvmEip1559FeeSlow.md)
 - [EstimatedEvmLegacyFee](docs/EstimatedEvmLegacyFee.md)
 - [EstimatedEvmLegacyFeeSlow](docs/EstimatedEvmLegacyFeeSlow.md)
 - [EstimatedFee](docs/EstimatedFee.md)
 - [EstimatedFixedFee](docs/EstimatedFixedFee.md)
 - [EstimatedUtxoFee](docs/EstimatedUtxoFee.md)
 - [EstimatedUtxoFeeSlow](docs/EstimatedUtxoFeeSlow.md)
 - [EvmContractCallDestination](docs/EvmContractCallDestination.md)
 - [EvmEIP191MessageSignDestination](docs/EvmEIP191MessageSignDestination.md)
 - [EvmEIP712MessageSignDestination](docs/EvmEIP712MessageSignDestination.md)
 - [EvmEip1559FeeBasePrice](docs/EvmEip1559FeeBasePrice.md)
 - [EvmEip1559FeeRate](docs/EvmEip1559FeeRate.md)
 - [EvmLegacyFeeBasePrice](docs/EvmLegacyFeeBasePrice.md)
 - [EvmLegacyFeeRate](docs/EvmLegacyFeeRate.md)
 - [ExchangeId](docs/ExchangeId.md)
 - [ExchangeTransferDestination](docs/ExchangeTransferDestination.md)
 - [ExchangeTransferSource](docs/ExchangeTransferSource.md)
 - [ExchangeWalletInfo](docs/ExchangeWalletInfo.md)
 - [ExchangeWalletInfoAllOfSubAccounts](docs/ExchangeWalletInfoAllOfSubAccounts.md)
 - [ExtendedTokenInfo](docs/ExtendedTokenInfo.md)
 - [FeeAmount](docs/FeeAmount.md)
 - [FeeGasLimit](docs/FeeGasLimit.md)
 - [FeeRate](docs/FeeRate.md)
 - [FeeType](docs/FeeType.md)
 - [FixedFeeRate](docs/FixedFeeRate.md)
 - [KeyShareHolder](docs/KeyShareHolder.md)
 - [KeyShareHolderGroup](docs/KeyShareHolderGroup.md)
 - [KeyShareHolderGroupStatus](docs/KeyShareHolderGroupStatus.md)
 - [KeyShareHolderGroupType](docs/KeyShareHolderGroupType.md)
 - [KeyShareHolderStatus](docs/KeyShareHolderStatus.md)
 - [KeyShareHolderType](docs/KeyShareHolderType.md)
 - [ListAddresses200Response](docs/ListAddresses200Response.md)
 - [ListKeyShareHolderGroups200Response](docs/ListKeyShareHolderGroups200Response.md)
 - [ListMpcProjects200Response](docs/ListMpcProjects200Response.md)
 - [ListMpcVaults200Response](docs/ListMpcVaults200Response.md)
 - [ListSupportedChains200Response](docs/ListSupportedChains200Response.md)
 - [ListSupportedTokens200Response](docs/ListSupportedTokens200Response.md)
 - [ListTokenBalancesForAddress200Response](docs/ListTokenBalancesForAddress200Response.md)
 - [ListTransactions200Response](docs/ListTransactions200Response.md)
 - [ListTssRequests200Response](docs/ListTssRequests200Response.md)
 - [ListUtxos200Response](docs/ListUtxos200Response.md)
 - [ListWallets200Response](docs/ListWallets200Response.md)
 - [ListWebhookEndpoints200Response](docs/ListWebhookEndpoints200Response.md)
 - [ListWebhookEventDefinitions200ResponseInner](docs/ListWebhookEventDefinitions200ResponseInner.md)
 - [ListWebhookEventLogs200Response](docs/ListWebhookEventLogs200Response.md)
 - [ListWebhookEvents200Response](docs/ListWebhookEvents200Response.md)
 - [LockUtxos201Response](docs/LockUtxos201Response.md)
 - [LockUtxosRequest](docs/LockUtxosRequest.md)
 - [LockUtxosRequestUtxosInner](docs/LockUtxosRequestUtxosInner.md)
 - [MPCDelegate](docs/MPCDelegate.md)
 - [MPCProject](docs/MPCProject.md)
 - [MPCVault](docs/MPCVault.md)
 - [MPCVaultType](docs/MPCVaultType.md)
 - [MPCWalletInfo](docs/MPCWalletInfo.md)
 - [MaxFeeAmount](docs/MaxFeeAmount.md)
 - [MaxTransferableValue](docs/MaxTransferableValue.md)
 - [MessageSignDestination](docs/MessageSignDestination.md)
 - [MessageSignDestinationType](docs/MessageSignDestinationType.md)
 - [MessageSignParams](docs/MessageSignParams.md)
 - [MessageSignSource](docs/MessageSignSource.md)
 - [MessageSignSourceType](docs/MessageSignSourceType.md)
 - [MpcContractCallSource](docs/MpcContractCallSource.md)
 - [MpcMessageSignSource](docs/MpcMessageSignSource.md)
 - [MpcSigningGroup](docs/MpcSigningGroup.md)
 - [MpcTransferSource](docs/MpcTransferSource.md)
 - [Pagination](docs/Pagination.md)
 - [PoolDetails](docs/PoolDetails.md)
 - [PoolDetailsAllOfValidatorsInfo](docs/PoolDetailsAllOfValidatorsInfo.md)
 - [PoolSummary](docs/PoolSummary.md)
 - [ReplaceType](docs/ReplaceType.md)
 - [RetryWebhookEventById201Response](docs/RetryWebhookEventById201Response.md)
 - [RootPubkey](docs/RootPubkey.md)
 - [SafeContractCallSource](docs/SafeContractCallSource.md)
 - [SafeTransferSource](docs/SafeTransferSource.md)
 - [SafeWallet](docs/SafeWallet.md)
 - [SmartContractInitiator](docs/SmartContractInitiator.md)
 - [SmartContractWalletInfo](docs/SmartContractWalletInfo.md)
 - [SmartContractWalletOperationType](docs/SmartContractWalletOperationType.md)
 - [SmartContractWalletType](docs/SmartContractWalletType.md)
 - [SourceGroup](docs/SourceGroup.md)
 - [StakingPoolType](docs/StakingPoolType.md)
 - [StakingSource](docs/StakingSource.md)
 - [Stakings](docs/Stakings.md)
 - [StakingsValidatorInfo](docs/StakingsValidatorInfo.md)
 - [TSSGroups](docs/TSSGroups.md)
 - [TSSRequest](docs/TSSRequest.md)
 - [TSSRequestStatus](docs/TSSRequestStatus.md)
 - [TSSRequestType](docs/TSSRequestType.md)
 - [TokenBalance](docs/TokenBalance.md)
 - [TokenBalanceBalance](docs/TokenBalanceBalance.md)
 - [TokenInfo](docs/TokenInfo.md)
 - [Transaction](docs/Transaction.md)
 - [TransactionApprover](docs/TransactionApprover.md)
 - [TransactionBlockInfo](docs/TransactionBlockInfo.md)
 - [TransactionCustodialAssetWalletSource](docs/TransactionCustodialAssetWalletSource.md)
 - [TransactionDepositFromAddressSource](docs/TransactionDepositFromAddressSource.md)
 - [TransactionDepositFromLoopSource](docs/TransactionDepositFromLoopSource.md)
 - [TransactionDepositFromWalletSource](docs/TransactionDepositFromWalletSource.md)
 - [TransactionDepositToAddressDestination](docs/TransactionDepositToAddressDestination.md)
 - [TransactionDepositToWalletDestination](docs/TransactionDepositToWalletDestination.md)
 - [TransactionDestination](docs/TransactionDestination.md)
 - [TransactionDestinationType](docs/TransactionDestinationType.md)
 - [TransactionDetail](docs/TransactionDetail.md)
 - [TransactionDetailAllOfTimeline](docs/TransactionDetailAllOfTimeline.md)
 - [TransactionDetails](docs/TransactionDetails.md)
 - [TransactionEvmContractDestination](docs/TransactionEvmContractDestination.md)
 - [TransactionEvmEip1559Fee](docs/TransactionEvmEip1559Fee.md)
 - [TransactionEvmLegacyFee](docs/TransactionEvmLegacyFee.md)
 - [TransactionExchangeWalletSource](docs/TransactionExchangeWalletSource.md)
 - [TransactionFee](docs/TransactionFee.md)
 - [TransactionFeeStationWalletSource](docs/TransactionFeeStationWalletSource.md)
 - [TransactionFixedFee](docs/TransactionFixedFee.md)
 - [TransactionInitiatorType](docs/TransactionInitiatorType.md)
 - [TransactionMPCWalletSource](docs/TransactionMPCWalletSource.md)
 - [TransactionMessageSignEIP191Destination](docs/TransactionMessageSignEIP191Destination.md)
 - [TransactionMessageSignEIP712Destination](docs/TransactionMessageSignEIP712Destination.md)
 - [TransactionRawTxInfo](docs/TransactionRawTxInfo.md)
 - [TransactionRbf](docs/TransactionRbf.md)
 - [TransactionRbfSource](docs/TransactionRbfSource.md)
 - [TransactionReplacement](docs/TransactionReplacement.md)
 - [TransactionRequestEvmEip1559Fee](docs/TransactionRequestEvmEip1559Fee.md)
 - [TransactionRequestEvmLegacyFee](docs/TransactionRequestEvmLegacyFee.md)
 - [TransactionRequestFee](docs/TransactionRequestFee.md)
 - [TransactionRequestFixedFee](docs/TransactionRequestFixedFee.md)
 - [TransactionRequestUtxoFee](docs/TransactionRequestUtxoFee.md)
 - [TransactionResend](docs/TransactionResend.md)
 - [TransactionResult](docs/TransactionResult.md)
 - [TransactionResultType](docs/TransactionResultType.md)
 - [TransactionSignatureResult](docs/TransactionSignatureResult.md)
 - [TransactionSigner](docs/TransactionSigner.md)
 - [TransactionSmartContractSafeWalletSource](docs/TransactionSmartContractSafeWalletSource.md)
 - [TransactionSource](docs/TransactionSource.md)
 - [TransactionSourceType](docs/TransactionSourceType.md)
 - [TransactionStatus](docs/TransactionStatus.md)
 - [TransactionSubStatus](docs/TransactionSubStatus.md)
 - [TransactionTimeline](docs/TransactionTimeline.md)
 - [TransactionTokeApproval](docs/TransactionTokeApproval.md)
 - [TransactionTokenAmount](docs/TransactionTokenAmount.md)
 - [TransactionTransferToAddressDestination](docs/TransactionTransferToAddressDestination.md)
 - [TransactionTransferToAddressDestinationAccountOutput](docs/TransactionTransferToAddressDestinationAccountOutput.md)
 - [TransactionTransferToAddressDestinationUtxoOutputsInner](docs/TransactionTransferToAddressDestinationUtxoOutputsInner.md)
 - [TransactionTransferToWalletDestination](docs/TransactionTransferToWalletDestination.md)
 - [TransactionType](docs/TransactionType.md)
 - [TransactionUtxo](docs/TransactionUtxo.md)
 - [TransactionUtxoFee](docs/TransactionUtxoFee.md)
 - [TransactionWebhookEventData](docs/TransactionWebhookEventData.md)
 - [TransferDestination](docs/TransferDestination.md)
 - [TransferDestinationType](docs/TransferDestinationType.md)
 - [TransferParams](docs/TransferParams.md)
 - [TransferSource](docs/TransferSource.md)
 - [UTXO](docs/UTXO.md)
 - [UpdateCustodialWalletParams](docs/UpdateCustodialWalletParams.md)
 - [UpdateExchangeWalletParams](docs/UpdateExchangeWalletParams.md)
 - [UpdateGroupAction](docs/UpdateGroupAction.md)
 - [UpdateKeyShareHolderGroupByIdRequest](docs/UpdateKeyShareHolderGroupByIdRequest.md)
 - [UpdateMpcProjectByIdRequest](docs/UpdateMpcProjectByIdRequest.md)
 - [UpdateMpcVaultByIdRequest](docs/UpdateMpcVaultByIdRequest.md)
 - [UpdateMpcWalletParams](docs/UpdateMpcWalletParams.md)
 - [UpdateSmartContractWalletParams](docs/UpdateSmartContractWalletParams.md)
 - [UpdateWalletParams](docs/UpdateWalletParams.md)
 - [UpdateWebhookEndpointByIdRequest](docs/UpdateWebhookEndpointByIdRequest.md)
 - [UtxoFeeBasePrice](docs/UtxoFeeBasePrice.md)
 - [UtxoFeeRate](docs/UtxoFeeRate.md)
 - [WalletInfo](docs/WalletInfo.md)
 - [WalletSubtype](docs/WalletSubtype.md)
 - [WalletType](docs/WalletType.md)
 - [WebhookEndpoint](docs/WebhookEndpoint.md)
 - [WebhookEndpointStatus](docs/WebhookEndpointStatus.md)
 - [WebhookEvent](docs/WebhookEvent.md)
 - [WebhookEventData](docs/WebhookEventData.md)
 - [WebhookEventDataType](docs/WebhookEventDataType.md)
 - [WebhookEventLog](docs/WebhookEventLog.md)
 - [WebhookEventStatus](docs/WebhookEventStatus.md)
 - [WebhookEventType](docs/WebhookEventType.md)


## Documentation For Authorization


Authentication schemes defined for the API:
### CoboAuth

- **Type**: API key
- **API key parameter name**: BIZ-API-KEY
- **Location**: HTTP header

Note, each API key must be added to a map of `map[string]APIKey` where the key is: BIZ-API-KEY and passed in as the auth context for each request.

Example

```go
auth := context.WithValue(
		context.Background(),
		waas2.ContextAPIKeys,
		map[string]waas2.APIKey{
			"BIZ-API-KEY": {Key: "API_KEY_STRING"},
		},
	)
r, err := client.Service.Operation(auth, args)
```

### OAuth2


- **Type**: OAuth
- **Flow**: accessCode
- **Authorization URL**: https://auth.cobo.com/authorize
- **Scopes**: 
 - **custodial_asset_wallet:create**: Create access to custodial asset wallets
 - **custodial_asset_wallet:add**: Generate address access to custodial asset wallets
 - **custodial_asset_wallet:edit**: Change wallet name access to custodial asset wallets
 - **custodial_asset_wallet:withdraw**: Withdraw access to custodial asset wallets
 - **mpc_organization_controlled_wallet:create**: Create access to MPC organization-controlled wallets
 - **mpc_organization_controlled_wallet:add**: Generate address access to MPC organization-controlled wallets
 - **mpc_organization_controlled_wallet:edit**: Change wallet name access to MPC organization-controlled wallets
 - **mpc_organization_controlled_wallet:withdraw**: Withdraw access to MPC organization-controlled wallets
 - **mpc_organization_controlled_wallet:contract_call**: Contract call access to MPC organization-controlled wallets
 - **mpc_organization_controlled_wallet:message_sign**: Message sign access to MPC organization-controlled wallets
 - **mpc_organization_controlled_vault:manage**: Create/Edit access to MPC organization-controlled vaults
 - **mpc_organization_controlled_key_group:manage**: Create/Edit/Delete access to MPC organization-controlled key groups
 - **mpc_organization_controlled_tss_request:manage**: Create/Cancel access to MPC organization-controlled tss requests
 - **mpc_user_controlled_wallet:create**: Create access to MPC user-controlled wallets
 - **mpc_user_controlled_wallet:add**: Generate address access to MPC user-controlled wallets
 - **mpc_user_controlled_wallet:edit**: Change wallet name access to MPC user-controlled wallets
 - **mpc_user_controlled_wallet:withdraw**: Withdraw access to MPC user-controlled wallets
 - **mpc_user_controlled_wallet:contract_call**: Contract call access to MPC user-controlled wallets
 - **mpc_user_controlled_wallet:message_sign**: Message sign access to MPC user-controlled wallets
 - **mpc_user_controlled_project:manage**: Create/Edit access to MPC user-controlled projects
 - **mpc_user_controlled_vault:manage**: Create/Edit access to MPC user-controlled vaults
 - **mpc_user_controlled_key_group:manage**: Create/Edit/Delete access to MPC user-controlled key groups
 - **mpc_user_controlled_tss_request:manage**: Create/Cancel access to MPC user-controlled tss requests
 - **webhook:resend**: Resend access to webhook events
 - **webhook_url:edit**: Create/Edit access to webhook urls

Example

```go
auth := context.WithValue(context.Background(), waas2.ContextAccessToken, "ACCESSTOKENSTRING")
r, err := client.Service.Operation(auth, args)
```

Or via OAuth2 module to automatically refresh tokens and perform user authentication.

```go
import "golang.org/x/oauth2"

/* Perform OAuth2 round trip request and obtain a token */

tokenSource := oauth2cfg.TokenSource(createContext(httpClient), &token)
auth := context.WithValue(oauth2.NoContext, waas2.ContextOAuth2, tokenSource)
r, err := client.Service.Operation(auth, args)
```


## Documentation for Utility Methods

Due to the fact that model structure members are all pointers, this package contains
a number of utility functions to easily obtain pointers to values of basic types.
Each of these functions takes a value of the given basic type and returns a pointer to it:

* `PtrBool`
* `PtrInt`
* `PtrInt32`
* `PtrInt64`
* `PtrFloat`
* `PtrFloat32`
* `PtrFloat64`
* `PtrString`
* `PtrTime`

## Author

support@cobo.com
