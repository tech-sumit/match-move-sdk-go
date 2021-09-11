# {{classname}}

All URIs are relative to *https://beta-api.mmvpay.com/{product}/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetUserAddresses**](UserAddressesApi.md#GetUserAddresses) | **Get** /users/addresses/{type} | Get users addresses
[**UpdateUserAddresses**](UserAddressesApi.md#UpdateUserAddresses) | **Put** /users/addresses/{type} | Update users addresses

# **GetUserAddresses**
> UserAddressV1 GetUserAddresses(ctx, type_, xAuthUserID)
Get users addresses

Retrieve the specified address type 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **type_** | **string**| The type of address for which there is a request to be made | 
  **xAuthUserID** | **string**| MatchMove provided hash ID for the user | 

### Return type

[**UserAddressV1**](User_address.v1.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateUserAddresses**
> UserAddressV1 UpdateUserAddresses(ctx, type_, xAuthUserID, optional)
Update users addresses

Update the specified address type 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **type_** | **string**| The type of address for which there is a request to be made | 
  **xAuthUserID** | **string**| MatchMove provided hash ID for the user | 
 **optional** | ***UserAddressesApiUpdateUserAddressesOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a UserAddressesApiUpdateUserAddressesOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **address1** | **optional.**|  | 
 **address2** | **optional.**|  | 
 **address3** | **optional.**|  | 
 **address4** | **optional.**|  | 
 **city** | **optional.**|  | 
 **state** | **optional.**|  | 
 **country** | **optional.**|  | 
 **countryCode** | **optional.**|  | 
 **zipcode** | **optional.**|  | 

### Return type

[**UserAddressV1**](User_address.v1.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: application/x-www-form-urlencoded
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

