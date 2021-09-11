# {{classname}}

All URIs are relative to */*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ConsumerConsumerHashProvidersGet**](ConsumerProvidersApi.md#ConsumerConsumerHashProvidersGet) | **Get** /consumer/{consumer_hash}/providers | Reference list for the providers

# **ConsumerConsumerHashProvidersGet**
> Response7 ConsumerConsumerHashProvidersGet(ctx, xAuthHeader, consumerHash, optional)
Reference list for the providers

This resource is used for creating reference list for the providers

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **xAuthHeader** | **string**| Consumer ACL Key | [default to b334822cc19e3c7dd86a1db2a4041c2f2f092f826dfc9b79e567ff69b299151f]
  **consumerHash** | **string**| Consumer Hash ID | 
 **optional** | ***ConsumerProvidersApiConsumerConsumerHashProvidersGetOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ConsumerProvidersApiConsumerConsumerHashProvidersGetOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **filterField** | **optional.String**| Filter Field; Allowable: [hash_id, type, currency, name] | [default to provider_hash]
 **filterCondition** | **optional.String**| Filter Condition; Allowable: [&#x3D;] | [default to &#x3D;]
 **filterValue** | **optional.String**| Filter Value | [default to 782c2b60e8bcd853b86c49a4144fca62]
 **sortField** | **optional.String**| Sort Field; Allowable: [date_created] | [default to date_created]
 **sortDirection** | **optional.String**| Direction For Sorting; Allowable: [asc, desc] | [default to asc]

### Return type

[**Response7**](response_7.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

