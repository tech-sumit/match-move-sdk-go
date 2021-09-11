# {{classname}}

All URIs are relative to *https://beta-api.mmvpay.com/{product}/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeleteUserWalletCardIdFunds**](FundsApi.md#DeleteUserWalletCardIdFunds) | **Delete** /users/wallets/cards/{card_id}/funds | Create Un-load transaction
[**DeleteUserWalletFunds**](FundsApi.md#DeleteUserWalletFunds) | **Delete** /users/wallets/funds | Create debit transaction
[**DeleteUserWalletFundsHolding**](FundsApi.md#DeleteUserWalletFundsHolding) | **Delete** /users/wallets/funds/holdings/{ref_hash} | Cancel pending debits/credits
[**GetUserWalletCardIdFunds**](FundsApi.md#GetUserWalletCardIdFunds) | **Get** /users/wallets/cards/{card_id}/funds | Get Debits/Credits of cards
[**GetUserWalletFunds**](FundsApi.md#GetUserWalletFunds) | **Get** /users/wallets/funds | Get funds
[**GetUserWalletFundsHolding**](FundsApi.md#GetUserWalletFundsHolding) | **Get** /users/wallets/funds/holdings/{limit}/{offset} | Get pending debit/credits
[**GetUserWalletFundsTransfer**](FundsApi.md#GetUserWalletFundsTransfer) | **Get** /users/wallets/funds/transfers | Get P2P transactions
[**PostUserWalletCardIdFunds**](FundsApi.md#PostUserWalletCardIdFunds) | **Post** /users/wallets/cards/{card_id}/funds | Create load transaction
[**PostUserWalletFunds**](FundsApi.md#PostUserWalletFunds) | **Post** /users/wallets/funds | Create credit or transfer transaction
[**PostUserWalletFundsHolding**](FundsApi.md#PostUserWalletFundsHolding) | **Post** /users/wallets/funds/holdings/{ref_hash} | Retry pending debits/credits
[**PostUserWalletRefunds**](FundsApi.md#PostUserWalletRefunds) | **Post** /users/wallets/refund/{type} | Create Refund
[**PutUserWalletFunds**](FundsApi.md#PutUserWalletFunds) | **Put** /users/wallets/funds | Update transfer transaction
[**PutUserWalletFundsTransfers**](FundsApi.md#PutUserWalletFundsTransfers) | **Put** /users/wallets/funds/transfers | Confirm pending transfer

# **DeleteUserWalletCardIdFunds**
> InlineResponse20020 DeleteUserWalletCardIdFunds(ctx, xAuthUserID, cardId, optional)
Create Un-load transaction

If the X-Auth-User-ID is passed, then this API will transfer credits from the card back to the wallet, ELSE, this API will initiate a closed loop debit transaction on a particular card linked to users wallet.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **xAuthUserID** | **string**| The &#x60;user_id&#x60; of the user returned via  GET /users | 
  **cardId** | **string**| The &#x60;card_id&#x60; for which the information is being retrieved.  | 
 **optional** | ***FundsApiDeleteUserWalletCardIdFundsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a FundsApiDeleteUserWalletCardIdFundsOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **amount** | **optional.**|  | 
 **message** | **optional.**|  | 
 **fundCategoryName** | **optional.**|  | 

### Return type

[**InlineResponse20020**](inline_response_200_20.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: application/x-www-form-urlencoded
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteUserWalletFunds**
> InlineResponse20022 DeleteUserWalletFunds(ctx, optional)
Create debit transaction

Initiate a  closed loop debit transaction for a particular users wallet in the system.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***FundsApiDeleteUserWalletFundsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a FundsApiDeleteUserWalletFundsOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **email** | **optional.**|  | 
 **mobileCountryCode** | **optional.**|  | 
 **mobile** | **optional.**|  | 
 **hashId** | **optional.**|  | 
 **amount** | **optional.**|  | 
 **details** | **optional.**|  | 
 **fundCategoryName** | **optional.**|  | 

### Return type

[**InlineResponse20022**](inline_response_200_22.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: application/x-www-form-urlencoded
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteUserWalletFundsHolding**
> InlineResponse20020 DeleteUserWalletFundsHolding(ctx, refHash)
Cancel pending debits/credits

Cancel  a closed loop transaction in the system which is in intermediate state.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **refHash** | **string**|  | 

### Return type

[**InlineResponse20020**](inline_response_200_20.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetUserWalletCardIdFunds**
> InlineResponse20020 GetUserWalletCardIdFunds(ctx, xAuthUserID, cardId)
Get Debits/Credits of cards

Retrieve closed loop debit and credits for a particular card linked to a users wallet

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **xAuthUserID** | **string**| The &#x60;user_id&#x60; of the user returned via  GET /users | 
  **cardId** | **string**| The &#x60;card_id&#x60; for which the information is being retrieved.  | 

### Return type

[**InlineResponse20020**](inline_response_200_20.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetUserWalletFunds**
> InlineResponse20019 GetUserWalletFunds(ctx, xAuthUserID)
Get funds

Retrieve closed loop debit and credits for a particular users

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **xAuthUserID** | **string**| MatchMove provided hash ID for the user | 

### Return type

[**InlineResponse20019**](inline_response_200_19.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetUserWalletFundsHolding**
> InlineResponse20020 GetUserWalletFundsHolding(ctx, limit, offset)
Get pending debit/credits

Retrieve closed loop transactions in the system which are in intermediate state.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **limit** | **int32**|  | 
  **offset** | **int32**|  | 

### Return type

[**InlineResponse20020**](inline_response_200_20.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetUserWalletFundsTransfer**
> InlineResponse20024 GetUserWalletFundsTransfer(ctx, xAuthUserID)
Get P2P transactions

Retrieve p2p closed loop transactions for a users wallet in the system

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **xAuthUserID** | **string**| MatchMove provided hash ID for the user | 

### Return type

[**InlineResponse20024**](inline_response_200_24.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PostUserWalletCardIdFunds**
> InlineResponse20020 PostUserWalletCardIdFunds(ctx, xAuthUserID, cardId, optional)
Create load transaction

Initiate a closed loop load transaction from the wallet to the specified card

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **xAuthUserID** | **string**| The &#x60;user_id&#x60; of the user returned via  GET /users | 
  **cardId** | **string**| The &#x60;card_id&#x60; for which the information is being retrieved.  | 
 **optional** | ***FundsApiPostUserWalletCardIdFundsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a FundsApiPostUserWalletCardIdFundsOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **amount** | **optional.**|  | 
 **message** | **optional.**|  | 
 **fundCategoryName** | **optional.**|  | 

### Return type

[**InlineResponse20020**](inline_response_200_20.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: application/x-www-form-urlencoded
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PostUserWalletFunds**
> InlineResponse20021 PostUserWalletFunds(ctx, optional)
Create credit or transfer transaction

Initiate a  closed loop credit transaction involving a particular user in the system.  A. Credit If the `X-Auth-User-Id` parameter is not passed as part of the header, this API will credit the recipient's wallet as per the users `email`/  `mobile` / `user_id` parameter defined as part of the request.  B. Transfer If the `X-Auth-User-Id` is passed,  this API will transfer amount from the wallet of the user defined as part of `X-Auth-User-id` to another wallet as per the users `email` / `mobile` / `user_id` parameter specified as part of the request 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***FundsApiPostUserWalletFundsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a FundsApiPostUserWalletFundsOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **email** | **optional.**|  | 
 **mobileCountryCode** | **optional.**|  | 
 **mobile** | **optional.**|  | 
 **userId** | **optional.**|  | 
 **amount** | **optional.**|  | 
 **details** | **optional.**|  | 
 **forwardedFor** | **optional.**|  | 
 **hashedPan** | **optional.**|  | 
 **fundCategoryName** | **optional.**|  | 
 **type_** | **optional.**|  | 
 **loadCard** | **optional.**|  | 
 **cardId** | **optional.**|  | 

### Return type

[**InlineResponse20021**](inline_response_200_21.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: application/x-www-form-urlencoded
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PostUserWalletFundsHolding**
> InlineResponse20020 PostUserWalletFundsHolding(ctx, refHash)
Retry pending debits/credits

Retry confirming a closed loop transaction in the system which is in intermediate state.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **refHash** | **string**|  | 

### Return type

[**InlineResponse20020**](inline_response_200_20.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PostUserWalletRefunds**
> InlineResponse20023 PostUserWalletRefunds(ctx, type_, optional)
Create Refund

Issue a partial or full refund out of the wallet. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **type_** | **string**|  | 
 **optional** | ***FundsApiPostUserWalletRefundsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a FundsApiPostUserWalletRefundsOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **amount** | **optional.**|  | 
 **details** | **optional.**|  | 

### Return type

[**InlineResponse20023**](inline_response_200_23.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: application/x-www-form-urlencoded
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PutUserWalletFunds**
> InlineResponse20020 PutUserWalletFunds(ctx, xAuthUserID)
Update transfer transaction

Confirm the two staged transfer transaction for legacy systems

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **xAuthUserID** | **string**| MatchMove provided hash ID for the user | 

### Return type

[**InlineResponse20020**](inline_response_200_20.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PutUserWalletFundsTransfers**
> InlineResponse20020 PutUserWalletFundsTransfers(ctx, xAuthUserID)
Confirm pending transfer

Confirm a p2p closed loop transactions for a user in the system.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **xAuthUserID** | **string**| MatchMove provided hash ID for the user | 

### Return type

[**InlineResponse20020**](inline_response_200_20.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

