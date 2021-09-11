# {{classname}}

All URIs are relative to */*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CheckoutTokenProviderHashPost**](DefaultApi.md#CheckoutTokenProviderHashPost) | **Post** /checkout/token/{provider_hash} | 

# **CheckoutTokenProviderHashPost**
> InlineResponse200 CheckoutTokenProviderHashPost(ctx, providerHash, xAuthHeader, optional)


This resource is used for the initial process of payment for providers involving tokens from Relay to the Payment Gateway 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **providerHash** | **string**| Provider Hash ID | 
  **xAuthHeader** | **string**| Consumer ACL Key | 
 **optional** | ***DefaultApiCheckoutTokenProviderHashPostOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a DefaultApiCheckoutTokenProviderHashPostOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **consumerHash** | **optional.**|  | 
 **productCode** | **optional.**|  | 
 **providerChannel** | **optional.**|  | 
 **amount** | **optional.**|  | 
 **kycStatus** | **optional.**|  | 
 **userHashId** | **optional.**|  | 
 **email** | **optional.**|  | 
 **mobileCountryCode** | **optional.**|  | 
 **mobile** | **optional.**|  | 
 **currency** | **optional.**|  | 
 **pan** | **optional.**|  | 
 **clientRefId** | **optional.**|  | 
 **actionOrigin** | **optional.**|  | 

### Return type

[**InlineResponse200**](inline_response_200.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/x-www-form-urlencoded
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

