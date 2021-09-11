# {{classname}}

All URIs are relative to *https://beta-api.mmvpay.com/{product}/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetOauthConsumerKycProviders**](KYCApi.md#GetOauthConsumerKycProviders) | **Get** oauth/consumer/kyc/providers | Get KYC Providers enabled for the program
[**PostUsersAuthenticationsDocumentsProcess**](KYCApi.md#PostUsersAuthenticationsDocumentsProcess) | **Post** users/authentications/documents/process | Process KYC Verification

# **GetOauthConsumerKycProviders**
> GetKycProvidersResponseV1 GetOauthConsumerKycProviders(ctx, )
Get KYC Providers enabled for the program

This endpoint calls the KYC Service to retrieve the providers configured for the program.  The provider id and user type would be retrieved from this endpoint to use it for the processing of KYC verification. 

### Required Parameters
This endpoint does not need any parameter.

### Return type

[**GetKycProvidersResponseV1**](Get_kyc_providers_response.v1.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PostUsersAuthenticationsDocumentsProcess**
> KycProcessResponseV1 PostUsersAuthenticationsDocumentsProcess(ctx, optional)
Process KYC Verification

This endpoint will initiate the KYC verification flow and will return the redirect web URL where end users can initialize the verification process for verifying their identities. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***KYCApiPostUsersAuthenticationsDocumentsProcessOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a KYCApiPostUsersAuthenticationsDocumentsProcessOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **providerId** | **optional.**|  | 
 **userType** | **optional.**|  | 
 **userId** | **optional.**|  | 
 **programCode** | **optional.**|  | 

### Return type

[**KycProcessResponseV1**](Kyc_process_response.v1.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: application/x-www-form-urlencoded
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

