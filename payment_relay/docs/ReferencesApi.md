# {{classname}}

All URIs are relative to */*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ReferencesFeesConsumerHashProviderHashGet**](ReferencesApi.md#ReferencesFeesConsumerHashProviderHashGet) | **Get** /references/fees/{consumer_hash}/{provider_hash} | Fee data for the consumer, product, provider, channel combination
[**ReferencesProvidersChannelsGet**](ReferencesApi.md#ReferencesProvidersChannelsGet) | **Get** /references/providers/channels | Get provider channels

# **ReferencesFeesConsumerHashProviderHashGet**
> Response8 ReferencesFeesConsumerHashProviderHashGet(ctx, xAuthHeader, consumerHash, providerHash, optional)
Fee data for the consumer, product, provider, channel combination

This resource is used for getting fee for given consumer, product, provider, channel combination

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **xAuthHeader** | **string**| Consumer ACL Key | [default to 717763e4d8384ff1a5e4c7b68d8032d7176b9f1168c180c3d6c9f045ab6bf95e]
  **consumerHash** | **string**| Consumer Hash ID | 
  **providerHash** | **string**| Provider Hash ID | 
 **optional** | ***ReferencesApiReferencesFeesConsumerHashProviderHashGetOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ReferencesApiReferencesFeesConsumerHashProviderHashGetOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **productCode** | **optional.String**| Product Code | [default to sgmcpr]
 **providerChannel** | **optional.String**| Provider Channel | 

### Return type

[**Response8**](response_8.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ReferencesProvidersChannelsGet**
> Response9 ReferencesProvidersChannelsGet(ctx, xAuthHeader)
Get provider channels

Get provider channels

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **xAuthHeader** | **string**| Consumer ACL Key | [default to 717763e4d8384ff1a5e4c7b68d8032d7176b9f1168c180c3d6c9f045ab6bf95e]

### Return type

[**Response9**](response_9.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

