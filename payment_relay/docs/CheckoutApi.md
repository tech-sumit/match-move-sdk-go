# {{classname}}

All URIs are relative to */*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CheckoutDirectProviderHashPost**](CheckoutApi.md#CheckoutDirectProviderHashPost) | **Post** /checkout/direct/{provider_hash} | Initial Payment Processing for Direct Payment Providers
[**CheckoutHostedProviderHashPost**](CheckoutApi.md#CheckoutHostedProviderHashPost) | **Post** /checkout/hosted/{provider_hash} | Initial Payment Processing for Hosted Providers
[**CheckoutValidateProviderHashPost**](CheckoutApi.md#CheckoutValidateProviderHashPost) | **Post** /checkout/validate/{provider_hash} | Checkout Validation Rules
[**CheckoutValidateProviderHashTypePost**](CheckoutApi.md#CheckoutValidateProviderHashTypePost) | **Post** /checkout/validate/{provider_hash}/{type} | Checkout Validation Rules

# **CheckoutDirectProviderHashPost**
> Response4 CheckoutDirectProviderHashPost(ctx, xAuthHeader, providerHash, optional)
Initial Payment Processing for Direct Payment Providers

This resource is used for the initial process of payment from Relay to the Payment Gateway

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **xAuthHeader** | **string**| Consumer ACL Key | [default to b334822cc19e3c7dd86a1db2a4041c2f2f092f826dfc9b79e567ff69b299151f]
  **providerHash** | **string**| Provider Hash ID | 
 **optional** | ***CheckoutApiCheckoutDirectProviderHashPostOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a CheckoutApiCheckoutDirectProviderHashPostOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **consumerHash** | **optional.**|  | 
 **productCode** | **optional.**|  | 
 **amount** | **optional.**|  | 
 **kycStatus** | **optional.**|  | 
 **userHashId** | **optional.**|  | 
 **email** | **optional.**|  | 
 **mobileCountryCode** | **optional.**|  | 
 **mobile** | **optional.**|  | 
 **currency** | **optional.**|  | 
 **cardToken** | **optional.**|  | 
 **pan** | **optional.**|  | 
 **expiryMonth** | **optional.**|  | 
 **expiryYear** | **optional.**|  | 
 **securityCode** | **optional.**|  | 
 **source** | **optional.**|  | 
 **billingAddress** | **optional.**|  | 
 **customer** | **optional.**|  | 
 **clientRefId** | **optional.**|  | 
 **clientIp** | **optional.**|  | 
 **actionOrigin** | **optional.**|  | 
 **saveCard** | **optional.**|  | 

### Return type

[**Response4**](response_4.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/x-www-form-urlencoded
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CheckoutHostedProviderHashPost**
> Response5 CheckoutHostedProviderHashPost(ctx, xAuthHeader, providerHash, optional)
Initial Payment Processing for Hosted Providers

This resource is used for the initial process of payment from Relay to the Payment Gateway

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **xAuthHeader** | **string**| Consumer ACL Key | [default to 717763e4d8384ff1a5e4c7b68d8032d7176b9f1168c180c3d6c9f045ab6bf95e]
  **providerHash** | **string**| Provider Hash ID | 
 **optional** | ***CheckoutApiCheckoutHostedProviderHashPostOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a CheckoutApiCheckoutHostedProviderHashPostOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **consumerHash** | **optional.**|  | 
 **productCode** | **optional.**|  | 
 **amount** | **optional.**|  | 
 **kycStatus** | **optional.**|  | 
 **userHashId** | **optional.**|  | 
 **email** | **optional.**|  | 
 **mobileCountryCode** | **optional.**|  | 
 **mobile** | **optional.**|  | 
 **preferredName** | **optional.**|  | 
 **inRoute** | **optional.**|  | 
 **currency** | **optional.**|  | 
 **countryCode** | **optional.**|  | 
 **pan** | **optional.**|  | 
 **clientRefId** | **optional.**|  | 
 **clientIp** | **optional.**|  | 
 **actionOrigin** | **optional.**|  | 

### Return type

[**Response5**](response_5.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/x-www-form-urlencoded
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CheckoutValidateProviderHashPost**
> Response6 CheckoutValidateProviderHashPost(ctx, providerHash, xAuthHeader, optional)
Checkout Validation Rules

This resource is used for validating the rules before top-up

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **providerHash** | **string**| Provider Hash ID | 
  **xAuthHeader** | **string**| Consumer ACL Key | [default to 717763e4d8384ff1a5e4c7b68d8032d7176b9f1168c180c3d6c9f045ab6bf95e]
 **optional** | ***CheckoutApiCheckoutValidateProviderHashPostOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a CheckoutApiCheckoutValidateProviderHashPostOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **consumerHash** | **optional.**|  | 
 **productCode** | **optional.**|  | 
 **amount** | **optional.**|  | 
 **userHashId** | **optional.**|  | 
 **kycStatus** | **optional.**|  | 

### Return type

[**Response6**](response_6.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/x-www-form-urlencoded
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CheckoutValidateProviderHashTypePost**
> Response6 CheckoutValidateProviderHashTypePost(ctx, providerHash, type_, xAuthHeader, optional)
Checkout Validation Rules

This resource is used for validating the rules before top-up

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **providerHash** | **string**| Provider Hash ID | 
  **type_** | **string**| Type to validate, [&#x60;wallet&#x60;, &#x60;card&#x60;] | 
  **xAuthHeader** | **string**| Consumer ACL Key | [default to 717763e4d8384ff1a5e4c7b68d8032d7176b9f1168c180c3d6c9f045ab6bf95e]
 **optional** | ***CheckoutApiCheckoutValidateProviderHashTypePostOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a CheckoutApiCheckoutValidateProviderHashTypePostOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **consumerHash** | **optional.**|  | 
 **productCode** | **optional.**|  | 
 **amount** | **optional.**|  | 
 **userHashId** | **optional.**|  | 
 **kycStatus** | **optional.**|  | 

### Return type

[**Response6**](response_6.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/x-www-form-urlencoded
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

