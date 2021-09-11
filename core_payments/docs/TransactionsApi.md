# {{classname}}

All URIs are relative to *https://beta-api.mmvpay.com/{product}/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetUserWalletCardTransactions**](TransactionsApi.md#GetUserWalletCardTransactions) | **Get** /users/wallets/cards/{card_id}/transactions/{page} | Get Card transactions
[**GetUserWalletTransaction**](TransactionsApi.md#GetUserWalletTransaction) | **Get** /users/wallets/transactions/{page} | Get Wallet Transactions

# **GetUserWalletCardTransactions**
> InlineResponse20012 GetUserWalletCardTransactions(ctx, cardId, xAuthUserID, page, optional)
Get Card transactions

Retrieve transactions performed on a particular users card in the system

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **cardId** | [**string**](.md)| The &#x60;card_id&#x60; for which the information is being retrieved.  | 
  **xAuthUserID** | **string**| MatchMove provided hash ID for the user | 
  **page** | **string**| the one-based page number (maximum value of 99) to be displayed. | 
 **optional** | ***TransactionsApiGetUserWalletCardTransactionsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a TransactionsApiGetUserWalletCardTransactionsOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **limit** | **optional.String**| Number of transactions return per page (Default 5) | [default to 5]
 **type_** | **optional.String**| Filter transactions according to transaction types. Values can be: [debit , credit].  | 
 **dateRange** | **optional.String**| Limit transactions based on date range. Example: &#x60;date_range&#x3D;20141220,20150131&#x60;.  | 
 **sort** | **optional.String**| Comma delimited sorting of the result according to amount, date.added and date.expiry. Example: sort&#x3D;&#x60;+amount&#x60;,&#x60;-date.added&#x60;,&#x60;-date.expiry&#x60;. Only works for in house balance products. | 

### Return type

[**InlineResponse20012**](inline_response_200_12.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetUserWalletTransaction**
> InlineResponse2001 GetUserWalletTransaction(ctx, page, xAuthUserID, optional)
Get Wallet Transactions

Retrieves the transaction history of the User's Virtual Wallet

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **page** | **string**| the one-based page number (maximum value of 99) to be displayed. | 
  **xAuthUserID** | **string**| MatchMove provided hash ID for the user | 
 **optional** | ***TransactionsApiGetUserWalletTransactionOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a TransactionsApiGetUserWalletTransactionOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **limit** | **optional.String**| Number of transactions return per page (Default 5) | [default to 5]
 **type_** | **optional.String**| Filter transactions according to transaction types. Values can be: [debit , credit].  | 
 **dateRange** | **optional.String**| Limit transactions based on date range. Example: &#x60;date_range&#x3D;20141220,20150131&#x60;.  | 
 **sort** | **optional.String**| Comma delimited sorting of the result according to amount, date.added and date.expiry. Example: sort&#x3D;&#x60;+amount&#x60;,&#x60;-date.added&#x60;,&#x60;-date.expiry&#x60;. Only works for in house balance products. | 

### Return type

[**InlineResponse2001**](inline_response_200_1.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

