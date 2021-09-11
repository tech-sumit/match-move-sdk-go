# {{classname}}

All URIs are relative to *https://beta-api.mmvpay.com/{product}/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateUserWallet**](WalletsApi.md#CreateUserWallet) | **Post** /users/wallets | Create Wallet
[**GetUserWallet**](WalletsApi.md#GetUserWallet) | **Get** /users/wallets | Retrieve Wallet

# **CreateUserWallet**
> WalletFullV1 CreateUserWallet(ctx, xAuthUserID)
Create Wallet

Create a wallet for a customer created in the system   Do note the resource has the following pre-requisites too be present in the system for the user information depending on regions   For Singapore - `Nationality`  For India - `id_type` - `id_number` 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **xAuthUserID** | **string**| MatchMove provided hash ID for the user | 

### Return type

[**WalletFullV1**](Wallet_full.v1.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetUserWallet**
> WalletFullV1 GetUserWallet(ctx, xAuthUserID)
Retrieve Wallet

Retrieve user's wallet details and information.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **xAuthUserID** | **string**| MatchMove provided hash ID for the user | 

### Return type

[**WalletFullV1**](Wallet_full.v1.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

