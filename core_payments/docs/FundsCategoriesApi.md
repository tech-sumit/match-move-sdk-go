# {{classname}}

All URIs are relative to *https://beta-api.mmvpay.com/{product}/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateUserWalletCardFundCategoryLimits**](FundsCategoriesApi.md#CreateUserWalletCardFundCategoryLimits) | **Post** /users/wallets/cards/funds/categories/{name}/limits | Create Fund Category Limits
[**CreateUserWalletCardFundCategoryTransfer**](FundsCategoriesApi.md#CreateUserWalletCardFundCategoryTransfer) | **Post** /users/wallets/cards/funds/categories/transfers | Transfer balance from one fund category to another
[**DeleteUserWalletCardFundCategory**](FundsCategoriesApi.md#DeleteUserWalletCardFundCategory) | **Delete** /users/wallets/cards/funds/categories/{name} | Delete a fund category defined in the system
[**DeleteUserWalletCardFundCategoryLimits**](FundsCategoriesApi.md#DeleteUserWalletCardFundCategoryLimits) | **Delete** /users/wallets/cards/funds/categories/{name}/limits | Delete fund categories limit
[**GetUserWalletCardFundCategories**](FundsCategoriesApi.md#GetUserWalletCardFundCategories) | **Get** /users/wallets/cards/funds/categories/{name} | Retrieve Fund Categories
[**GetUserWalletCardFundCategoryLimits**](FundsCategoriesApi.md#GetUserWalletCardFundCategoryLimits) | **Get** /users/wallets/cards/funds/categories/{name}/limits | Get Fund Category Limits
[**PostUserWalletCardFundCategory**](FundsCategoriesApi.md#PostUserWalletCardFundCategory) | **Post** /users/wallets/cards/funds/categories | Create a new fund category in the system
[**PutUserWalletCardFundCategory**](FundsCategoriesApi.md#PutUserWalletCardFundCategory) | **Put** /users/wallets/cards/funds/categories/{name} | Update attributes of fund category defined in the system
[**UpdateUserWalletCardFundCategoryLimits**](FundsCategoriesApi.md#UpdateUserWalletCardFundCategoryLimits) | **Put** /users/wallets/cards/funds/categories/{name}/limits | Update Fund Category Limits

# **CreateUserWalletCardFundCategoryLimits**
> FundCategoriesLimitResponseV1 CreateUserWalletCardFundCategoryLimits(ctx, name, optional)
Create Fund Category Limits

Create Fund Category Limits 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **name** | **string**|  | 
 **optional** | ***FundsCategoriesApiCreateUserWalletCardFundCategoryLimitsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a FundsCategoriesApiCreateUserWalletCardFundCategoryLimitsOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **settings** | [**optional.Interface of FundCategoriesLimitRequestV1Settings**](.md)|  | 

### Return type

[**FundCategoriesLimitResponseV1**](Fund_categories_limit_response.v1.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: application/x-www-form-urlencoded
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CreateUserWalletCardFundCategoryTransfer**
> InlineResponse20011 CreateUserWalletCardFundCategoryTransfer(ctx, xAuthUserID, optional)
Transfer balance from one fund category to another

Transfer funds from one category to another 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **xAuthUserID** | **string**| MatchMove provided hash ID for the user | 
 **optional** | ***FundsCategoriesApiCreateUserWalletCardFundCategoryTransferOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a FundsCategoriesApiCreateUserWalletCardFundCategoryTransferOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **amount** | **optional.**|  | 
 **from** | **optional.**|  | 
 **to** | **optional.**|  | 

### Return type

[**InlineResponse20011**](inline_response_200_11.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: application/x-www-form-urlencoded
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteUserWalletCardFundCategory**
> FundCategoriesV1 DeleteUserWalletCardFundCategory(ctx, name)
Delete a fund category defined in the system

Delete fund category in the system 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **name** | **string**|  | 

### Return type

[**FundCategoriesV1**](Fund_categories.v1.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteUserWalletCardFundCategoryLimits**
> InlineResponse20010 DeleteUserWalletCardFundCategoryLimits(ctx, name)
Delete fund categories limit

Delete Fund Category Limits 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **name** | **string**|  | 

### Return type

[**InlineResponse20010**](inline_response_200_10.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetUserWalletCardFundCategories**
> InlineResponse2009 GetUserWalletCardFundCategories(ctx, name)
Retrieve Fund Categories

Retrieve Fund Categories 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **name** | **string**|  | 

### Return type

[**InlineResponse2009**](inline_response_200_9.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetUserWalletCardFundCategoryLimits**
> FundCategoriesLimitRequestV1 GetUserWalletCardFundCategoryLimits(ctx, name)
Get Fund Category Limits

Get limits of a particular category by name 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **name** | **string**|  | 

### Return type

[**FundCategoriesLimitRequestV1**](Fund_categories_limit_request.v1.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PostUserWalletCardFundCategory**
> FundCategoriesV1 PostUserWalletCardFundCategory(ctx, optional)
Create a new fund category in the system

Create a fund category 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***FundsCategoriesApiPostUserWalletCardFundCategoryOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a FundsCategoriesApiPostUserWalletCardFundCategoryOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **name** | **optional.**|  | 
 **defaultAmount** | **optional.**|  | 
 **factor** | **optional.**|  | 
 **mcc** | **optional.**|  | 
 **country** | **optional.**|  | 
 **priority** | **optional.**|  | 
 **card** | **optional.**|  | 
 **merchant** | **optional.**|  | 
 **terminal** | **optional.**|  | 

### Return type

[**FundCategoriesV1**](Fund_categories.v1.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: application/x-www-form-urlencoded
 - **Accept**: application/json, application/x-www-form-urlencoded

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PutUserWalletCardFundCategory**
> FundCategoriesV1 PutUserWalletCardFundCategory(ctx, name, optional)
Update attributes of fund category defined in the system

Update fund category 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **name** | **string**|  | 
 **optional** | ***FundsCategoriesApiPutUserWalletCardFundCategoryOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a FundsCategoriesApiPutUserWalletCardFundCategoryOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **defaultAmount** | **optional.**|  | 
 **factor** | **optional.**|  | 
 **mcc** | **optional.**|  | 
 **country** | **optional.**|  | 
 **priority** | **optional.**|  | 
 **card** | **optional.**|  | 
 **merchant** | **optional.**|  | 
 **terminal** | **optional.**|  | 

### Return type

[**FundCategoriesV1**](Fund_categories.v1.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: application/x-www-form-urlencoded
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateUserWalletCardFundCategoryLimits**
> FundCategoriesLimitResponseV1 UpdateUserWalletCardFundCategoryLimits(ctx, name, optional)
Update Fund Category Limits

Update Fund Category Limits 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **name** | **string**|  | 
 **optional** | ***FundsCategoriesApiUpdateUserWalletCardFundCategoryLimitsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a FundsCategoriesApiUpdateUserWalletCardFundCategoryLimitsOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **settings** | [**optional.Interface of FundCategoriesLimitRequestV1Settings**](.md)|  | 

### Return type

[**FundCategoriesLimitResponseV1**](Fund_categories_limit_response.v1.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: application/x-www-form-urlencoded
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

