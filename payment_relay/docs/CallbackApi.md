# {{classname}}

All URIs are relative to */*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CallbackHookTypePaymentRefIdGet**](CallbackApi.md#CallbackHookTypePaymentRefIdGet) | **Get** /callback/hook/{type}/{payment_ref_id} | Transaction confirmation hook
[**CallbackPaymentRefIdSecureKeyGet**](CallbackApi.md#CallbackPaymentRefIdSecureKeyGet) | **Get** /callback/{payment_ref_id}/{secure_key} | Provider call back
[**CallbackPaymentRefIdSignaturePost**](CallbackApi.md#CallbackPaymentRefIdSignaturePost) | **Post** /callback/{payment_ref_id}/{signature} | Provider call back
[**CallbackTokensTokenPost**](CallbackApi.md#CallbackTokensTokenPost) | **Post** /callback/tokens/{token} | Callback for Providers with Token Reference
[**RedirectGet**](CallbackApi.md#RedirectGet) | **Get** /redirect | Redirect after final payment(success/fail)
[**RedirectProviderHashGet**](CallbackApi.md#RedirectProviderHashGet) | **Get** /redirect/{provider_hash} | Redirect after final payment(success/fail)

# **CallbackHookTypePaymentRefIdGet**
> Response2 CallbackHookTypePaymentRefIdGet(ctx, type_, paymentRefId)
Transaction confirmation hook

Calls consumer pre or post confirmation urls for additional logic before and after finalization of payment and topup

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **type_** | **string**| Transaction (&#x60;pre&#x60;, &#x60;post&#x60;) | 
  **paymentRefId** | **string**| Payment Reference Transaction | 

### Return type

[**Response2**](response_2.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CallbackPaymentRefIdSecureKeyGet**
> Response1 CallbackPaymentRefIdSecureKeyGet(ctx, paymentRefId, secureKey, respCode, reasonCode)
Provider call back

This resource finalizes the payment and processes the topup. This will be triggered from provider side.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **paymentRefId** | **string**| Payment Reference Transaction | 
  **secureKey** | **string**| Callback secure key for security | 
  **respCode** | **string**| Callback RespCode | [default to 1]
  **reasonCode** | **string**| Callback ReasonCode | [default to 1]

### Return type

[**Response1**](response_1.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CallbackPaymentRefIdSignaturePost**
> Response CallbackPaymentRefIdSignaturePost(ctx, razorpay, easypay2, enets, paymentRefId, signature, transaction)
Provider call back

This resource finalizes the payment and processes the topup. This will be triggered from provider side.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **razorpay** | [**RazorpayCallbackRequest**](.md)|  | 
  **easypay2** | [**Easypay2CallbackRequest**](.md)|  | 
  **enets** | [**EnetsCallbackRequest**](.md)|  | 
  **paymentRefId** | **string**| Payment Reference Transaction | 
  **signature** | **string**| Callback Signature for security | 
  **transaction** | **string**| Transaction type [&#x60;success&#x60;, &#x60;failed&#x60;, &#x60;cancelled&#x60;] | 

### Return type

[**Response**](response.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/x-www-form-urlencoded
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CallbackTokensTokenPost**
> Response3 CallbackTokensTokenPost(ctx, xAuthHeader, token, optional)
Callback for Providers with Token Reference

This resource receives token and process the topup

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **xAuthHeader** | **string**| Consumer ACL Key | [default to 717763e4d8384ff1a5e4c7b68d8032d7176b9f1168c180c3d6c9f045ab6bf95e]
  **token** | **string**| Token Payment Reference Transaction | 
 **optional** | ***CallbackApiCallbackTokensTokenPostOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a CallbackApiCallbackTokensTokenPostOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **signature** | **optional.**|  | 
 **amount** | **optional.**|  | 
 **currency** | **optional.**|  | 
 **clientRefId** | **optional.**|  | 
 **additionalHook** | **optional.**|  | 

### Return type

[**Response3**](response_3.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/x-www-form-urlencoded
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **RedirectGet**
> RedirectGet(ctx, )
Redirect after final payment(success/fail)

Redirect after final payment(success/fail)

### Required Parameters
This endpoint does not need any parameter.

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **RedirectProviderHashGet**
> RedirectProviderHashGet(ctx, )
Redirect after final payment(success/fail)

Redirect after final payment(success/fail)

### Required Parameters
This endpoint does not need any parameter.

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

