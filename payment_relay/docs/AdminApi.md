# {{classname}}

All URIs are relative to */*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ManageConsumerRatesConsumerHashProviderHashGet**](AdminApi.md#ManageConsumerRatesConsumerHashProviderHashGet) | **Get** /manage/consumer/rates/{consumer_hash}/{provider_hash} | Gets rate per consumer and provider
[**ManageConsumerRatesConsumerHashProviderHashPost**](AdminApi.md#ManageConsumerRatesConsumerHashProviderHashPost) | **Post** /manage/consumer/rates/{consumer_hash}/{provider_hash} | Add rate per consumer and provider
[**ManageConsumerRatesConsumerHashProviderHashPut**](AdminApi.md#ManageConsumerRatesConsumerHashProviderHashPut) | **Put** /manage/consumer/rates/{consumer_hash}/{provider_hash} | Inactivate consumer rate
[**ManageProvidersPost**](AdminApi.md#ManageProvidersPost) | **Post** /manage/providers | Add provider

# **ManageConsumerRatesConsumerHashProviderHashGet**
> Fees ManageConsumerRatesConsumerHashProviderHashGet(ctx, xAuthHeader, consumerHash, providerHash, productCode, optional)
Gets rate per consumer and provider

Will show the rate per consumer and provider

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **xAuthHeader** | **string**| Consumer ACL Key | [default to 717763e4d8384ff1a5e4c7b68d8032d7176b9f1168c180c3d6c9f045ab6bf95e]
  **consumerHash** | **string**| Consumer Hash ID | 
  **providerHash** | **string**| provider Hash ID | 
  **productCode** | **string**| product code | [default to sgmcpr]
 **optional** | ***AdminApiManageConsumerRatesConsumerHashProviderHashGetOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a AdminApiManageConsumerRatesConsumerHashProviderHashGetOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




 **providerChannel** | **optional.String**| provider channel | 

### Return type

[**Fees**](Fees.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ManageConsumerRatesConsumerHashProviderHashPost**
> Fees ManageConsumerRatesConsumerHashProviderHashPost(ctx, xAuthHeader, consumerHash, providerHash, optional)
Add rate per consumer and provider

Add rate per consumer and provider.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **xAuthHeader** | **string**| Consumer ACL Key | [default to b334822cc19e3c7dd86a1db2a4041c2f2f092f826dfc9b79e567ff69b299151f]
  **consumerHash** | **string**| Consumer Hash ID | 
  **providerHash** | **string**| provider Hash ID | 
 **optional** | ***AdminApiManageConsumerRatesConsumerHashProviderHashPostOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a AdminApiManageConsumerRatesConsumerHashProviderHashPostOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **productCode** | **optional.**|  | 
 **feeFlat** | **optional.**|  | 
 **feePercentage** | **optional.**|  | 
 **taxFlat** | **optional.**|  | 
 **taxPercentage** | **optional.**|  | 
 **chargeFlat** | **optional.**|  | 
 **chargePercentage** | **optional.**|  | 
 **serviceFlat** | **optional.**|  | 
 **servicePercentage** | **optional.**|  | 

### Return type

[**Fees**](Fees.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/x-www-form-urlencoded
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ManageConsumerRatesConsumerHashProviderHashPut**
> Fees ManageConsumerRatesConsumerHashProviderHashPut(ctx, xAuthHeader, consumerHash, providerHash, optional)
Inactivate consumer rate

Inactivates rate for specific consumer product provider

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **xAuthHeader** | **string**| Consumer ACL Key | [default to 717763e4d8384ff1a5e4c7b68d8032d7176b9f1168c180c3d6c9f045ab6bf95e]
  **consumerHash** | **string**| Consumer Hash ID | 
  **providerHash** | **string**| Provider Hash ID | 
 **optional** | ***AdminApiManageConsumerRatesConsumerHashProviderHashPutOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a AdminApiManageConsumerRatesConsumerHashProviderHashPutOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **productCode** | **optional.**|  | 
 **inactivate** | **optional.**|  | 

### Return type

[**Fees**](Fees.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/x-www-form-urlencoded
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ManageProvidersPost**
> ManageProvidersPost(ctx, xAuthHeader, optional)
Add provider

Add provider

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **xAuthHeader** | **string**| Consumer ACL Key | [default to f88acc4dd9848a63d1a726e6a205626c4c8ae873fa099715db4e27e1ad0f2586]
 **optional** | ***AdminApiManageProvidersPostOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a AdminApiManageProvidersPostOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **providerName** | **optional.**|  | 
 **type_** | **optional.**|  | 
 **providerMode** | **optional.**|  | 
 **currency** | **optional.**|  | 
 **algorithm** | **optional.**|  | 
 **description** | **optional.**|  | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/x-www-form-urlencoded
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

