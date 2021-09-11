# {{classname}}

All URIs are relative to *https://beta-api.mmvpay.com/{product}/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeleteOauthConsumerTransactionsRuleHashI**](TransactionRulesApi.md#DeleteOauthConsumerTransactionsRuleHashI) | **Delete** /oauth/consumer/transactions/rule/{hash_id} | 
[**GetOauthConsumerTransactionsRules**](TransactionRulesApi.md#GetOauthConsumerTransactionsRules) | **Get** /oauth/consumer/transactions/rules | 
[**PostOauthConsumerTransactionsRuleHashId**](TransactionRulesApi.md#PostOauthConsumerTransactionsRuleHashId) | **Post** /oauth/consumer/transactions/rule/{hash_id} | 
[**PostOauthConsumerTransactionsRules**](TransactionRulesApi.md#PostOauthConsumerTransactionsRules) | **Post** /oauth/consumer/transactions/rules | 
[**PutOauthConsumerTransactionsRuleHashId**](TransactionRulesApi.md#PutOauthConsumerTransactionsRuleHashId) | **Put** /oauth/consumer/transactions/rule/{hash_id} | 

# **DeleteOauthConsumerTransactionsRuleHashI**
> InlineResponse20035 DeleteOauthConsumerTransactionsRuleHashI(ctx, hashId)


Delete transaction rule entry

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **hashId** | **string**|  | 

### Return type

[**InlineResponse20035**](inline_response_200_35.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetOauthConsumerTransactionsRules**
> []TransactionRulesV1Inner GetOauthConsumerTransactionsRules(ctx, )


Creates new transaction rules in db

### Required Parameters
This endpoint does not need any parameter.

### Return type

[**[]TransactionRulesV1Inner**](array.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PostOauthConsumerTransactionsRuleHashId**
> []TransactionRulesV1Inner PostOauthConsumerTransactionsRuleHashId(ctx, hashId, optional)


Update transaction rule entry

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **hashId** | **string**|  | 
 **optional** | ***TransactionRulesApiPostOauthConsumerTransactionsRuleHashIdOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a TransactionRulesApiPostOauthConsumerTransactionsRuleHashIdOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **id** | **optional.**|  | 

### Return type

[**[]TransactionRulesV1Inner**](array.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: application/x-www-form-urlencoded
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PostOauthConsumerTransactionsRules**
> []TransactionRulesV1Inner PostOauthConsumerTransactionsRules(ctx, optional)


Creates new transaction rules in db

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***TransactionRulesApiPostOauthConsumerTransactionsRulesOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a TransactionRulesApiPostOauthConsumerTransactionsRulesOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **attribute** | **optional.**|  | 
 **action** | **optional.**|  | 
 **value** | **optional.**|  | 
 **cardTypeCode** | **optional.**|  | 
 **userState** | **optional.**|  | 

### Return type

[**[]TransactionRulesV1Inner**](array.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: application/x-www-form-urlencoded
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PutOauthConsumerTransactionsRuleHashId**
> []TransactionRulesV1Inner PutOauthConsumerTransactionsRuleHashId(ctx, hashId, optional)


Update transaction rule entry

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **hashId** | **string**|  | 
 **optional** | ***TransactionRulesApiPutOauthConsumerTransactionsRuleHashIdOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a TransactionRulesApiPutOauthConsumerTransactionsRuleHashIdOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **attribute** | **optional.**|  | 
 **action** | **optional.**|  | 
 **value** | **optional.**|  | 
 **cardTypeCode** | **optional.**|  | 
 **userState** | **optional.**|  | 

### Return type

[**[]TransactionRulesV1Inner**](array.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: application/x-www-form-urlencoded
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

