# {{classname}}

All URIs are relative to *https://beta-api.mmvpay.com/{product}/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ConfirmUsersAuthenticationDocuments**](UsersIdentityAuthenticationsApi.md#ConfirmUsersAuthenticationDocuments) | **Put** /users/authentications/documents | Confirm submission of KYC (Know Your Customer) documents of the user in the system
[**CreateUsersAuthenticationDocuments**](UsersIdentityAuthenticationsApi.md#CreateUsersAuthenticationDocuments) | **Post** /users/authentications/documents | Create KYC
[**DeleteUsersAuthenticationsDocuments**](UsersIdentityAuthenticationsApi.md#DeleteUsersAuthenticationsDocuments) | **Delete** /users/authentications/documents | Delete an uploaded image of the user
[**GetUsersAuthenticationDocuments**](UsersIdentityAuthenticationsApi.md#GetUsersAuthenticationDocuments) | **Get** /users/authentications/documents | Get KYC
[**UpdateUsersAuthenticationDocumentsStatus**](UsersIdentityAuthenticationsApi.md#UpdateUsersAuthenticationDocumentsStatus) | **Put** /users/authentications/statuses | Update KYC status

# **ConfirmUsersAuthenticationDocuments**
> InlineResponse20016 ConfirmUsersAuthenticationDocuments(ctx, xAuthUserID)
Confirm submission of KYC (Know Your Customer) documents of the user in the system

Used to confirm the process of Identity Authentication documents. Post this the application is processed via a manual process or based on region specific third party integrations done. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **xAuthUserID** | **string**| MatchMove provided hash ID for the user | 

### Return type

[**InlineResponse20016**](inline_response_200_16.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CreateUsersAuthenticationDocuments**
> InlineResponse20017 CreateUsersAuthenticationDocuments(ctx, xAuthUserID, optional)
Create KYC

Initialize a submissions session for users Identity authentication.  Do note the resource has the following pre-requisites too be present in the system for the user information  - `title` - `id_type` - `id_number` - `country_of_issue` - `gender`  Depending on geography, certain geographies might have additional attributes as required for the user model.  In addition, the residential address of the users needs to be updated and present in the system.  Also note the count of images uploaded is restricted to *4* for a particular The image types are restricted to .jpg , .png , .pdf and each file can be up to 8MB in size. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **xAuthUserID** | **string**| MatchMove provided hash ID for the user | 
 **optional** | ***UsersIdentityAuthenticationsApiCreateUsersAuthenticationDocumentsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a UsersIdentityAuthenticationsApiCreateUsersAuthenticationDocumentsOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **data** | **optional.**|  | 
 **documentType** | **optional.**|  | 
 **kycType** | **optional.**|  | 

### Return type

[**InlineResponse20017**](inline_response_200_17.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: application/x-www-form-urlencoded
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteUsersAuthenticationsDocuments**
> InlineResponse20010 DeleteUsersAuthenticationsDocuments(ctx, xAuthUserID, optional)
Delete an uploaded image of the user

End point to delete user kyc document

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **xAuthUserID** | **string**| MatchMove provided hash ID for the user | 
 **optional** | ***UsersIdentityAuthenticationsApiDeleteUsersAuthenticationsDocumentsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a UsersIdentityAuthenticationsApiDeleteUsersAuthenticationsDocumentsOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **document** | **optional.**|  | 

### Return type

[**InlineResponse20010**](inline_response_200_10.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: application/x-www-form-urlencoded
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetUsersAuthenticationDocuments**
> InlineResponse20015 GetUsersAuthenticationDocuments(ctx, xAuthUserID)
Get KYC

Retrieve details of users Identity Authentication Documents currently enrolled in the system

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **xAuthUserID** | **string**| MatchMove provided hash ID for the user | 

### Return type

[**InlineResponse20015**](inline_response_200_15.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateUsersAuthenticationDocumentsStatus**
> InlineResponse20018 UpdateUsersAuthenticationDocumentsStatus(ctx, optional)
Update KYC status

Update  KYC (Know your  Customer) status of the user in the system 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***UsersIdentityAuthenticationsApiUpdateUsersAuthenticationDocumentsStatusOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a UsersIdentityAuthenticationsApiUpdateUsersAuthenticationDocumentsStatusOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **hashId** | **optional.**|  | 
 **status** | **optional.**|  | 
 **subStatus** | **optional.**|  | 
 **message** | **optional.**|  | 
 **kycDetails** | **optional.**|  | 
 **kycType** | **optional.**|  | 

### Return type

[**InlineResponse20018**](inline_response_200_18.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: application/x-www-form-urlencoded
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

