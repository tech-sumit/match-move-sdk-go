# {{classname}}

All URIs are relative to *https://beta-api.mmvpay.com/{product}/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetProgramEnumerationCountries**](EnumerationsUsersApi.md#GetProgramEnumerationCountries) | **Get** /users/enumerations/countries | Retrieve supported &#x60;countries&#x60;
[**GetProgramEnumerationDocumentType**](EnumerationsUsersApi.md#GetProgramEnumerationDocumentType) | **Get** /users/enumerations/document_types | Retrieve supported &#x60;document_types&#x60;
[**GetProgramEnumerationGender**](EnumerationsUsersApi.md#GetProgramEnumerationGender) | **Get** /users/enumerations/genders | Retrieve supported &#x60;genders&#x60;
[**GetProgramEnumerationIdType**](EnumerationsUsersApi.md#GetProgramEnumerationIdType) | **Get** /users/enumerations/id_types | Retrieve supported &#x60;id_types&#x60;
[**GetProgramEnumerationKycType**](EnumerationsUsersApi.md#GetProgramEnumerationKycType) | **Get** /users/enumerations/kyc_types | Retrieve supported &#x60;kyc_types&#x60;
[**GetProgramEnumerationMaritalStatus**](EnumerationsUsersApi.md#GetProgramEnumerationMaritalStatus) | **Get** /users/enumerations/marital_status | Retrieve supported &#x60;marital_status&#x60;
[**GetProgramEnumerationMobileCountryCode**](EnumerationsUsersApi.md#GetProgramEnumerationMobileCountryCode) | **Get** /users/enumerations/mobile_country_codes | Retrieve supported &#x60;mobile_country_codes&#x60;
[**GetProgramEnumerationNationality**](EnumerationsUsersApi.md#GetProgramEnumerationNationality) | **Get** /users/enumerations/nationalities | Retrieve supported &#x60;nationalities&#x60;
[**GetProgramEnumerationTitle**](EnumerationsUsersApi.md#GetProgramEnumerationTitle) | **Get** /users/enumerations/titles | Retrieve supported &#x60;titles&#x60;

# **GetProgramEnumerationCountries**
> InlineResponse20033 GetProgramEnumerationCountries(ctx, optional)
Retrieve supported `countries`

Retrieve supported `countries` 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***EnumerationsUsersApiGetProgramEnumerationCountriesOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a EnumerationsUsersApiGetProgramEnumerationCountriesOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **type_** | **optional.String**|  | 

### Return type

[**InlineResponse20033**](inline_response_200_33.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetProgramEnumerationDocumentType**
> InlineResponse20031 GetProgramEnumerationDocumentType(ctx, )
Retrieve supported `document_types`

Retrieve supported `document_types` 

### Required Parameters
This endpoint does not need any parameter.

### Return type

[**InlineResponse20031**](inline_response_200_31.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetProgramEnumerationGender**
> InlineResponse20026 GetProgramEnumerationGender(ctx, )
Retrieve supported `genders`

Retrieve supported `genders` 

### Required Parameters
This endpoint does not need any parameter.

### Return type

[**InlineResponse20026**](inline_response_200_26.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetProgramEnumerationIdType**
> InlineResponse20029 GetProgramEnumerationIdType(ctx, )
Retrieve supported `id_types`

Retrieve supported `id_types` 

### Required Parameters
This endpoint does not need any parameter.

### Return type

[**InlineResponse20029**](inline_response_200_29.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetProgramEnumerationKycType**
> InlineResponse20030 GetProgramEnumerationKycType(ctx, )
Retrieve supported `kyc_types`

Retrieve supported `kyc_types` 

### Required Parameters
This endpoint does not need any parameter.

### Return type

[**InlineResponse20030**](inline_response_200_30.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetProgramEnumerationMaritalStatus**
> InlineResponse20026 GetProgramEnumerationMaritalStatus(ctx, )
Retrieve supported `marital_status`

Retrieve supported `marital_status` 

### Required Parameters
This endpoint does not need any parameter.

### Return type

[**InlineResponse20026**](inline_response_200_26.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetProgramEnumerationMobileCountryCode**
> InlineResponse20032 GetProgramEnumerationMobileCountryCode(ctx, )
Retrieve supported `mobile_country_codes`

Retrieve supported `mobile_country_codes` 

### Required Parameters
This endpoint does not need any parameter.

### Return type

[**InlineResponse20032**](inline_response_200_32.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetProgramEnumerationNationality**
> InlineResponse20028 GetProgramEnumerationNationality(ctx, )
Retrieve supported `nationalities`

Retrieve supported `nationalities` 

### Required Parameters
This endpoint does not need any parameter.

### Return type

[**InlineResponse20028**](inline_response_200_28.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetProgramEnumerationTitle**
> InlineResponse20027 GetProgramEnumerationTitle(ctx, )
Retrieve supported `titles`

Retrieve supported `titles` 

### Required Parameters
This endpoint does not need any parameter.

### Return type

[**InlineResponse20027**](inline_response_200_27.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

