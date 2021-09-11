# {{classname}}

All URIs are relative to *https://beta-api.mmvpay.com/{product}/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeleteOauthConsumerTransactionCategoriesCategoryHashIdRule**](CardsTransactionalLimitsApi.md#DeleteOauthConsumerTransactionCategoriesCategoryHashIdRule) | **Delete** /oauth/consumer/transaction/categories/{category_hash_id}/rule | Delete rules for a particular transaction category
[**DeleteUsersWalletsCardsCardHashIdTransactionCategoryHashIdLimits**](CardsTransactionalLimitsApi.md#DeleteUsersWalletsCardsCardHashIdTransactionCategoryHashIdLimits) | **Delete** /users/wallets/cards/{card_hash_id}/transactions/categories/{category_hash_id}/limit | Disable transactional limits for a particular card and category
[**GetOauthConsumerTransactionCategoriesCategoryHashIdRule**](CardsTransactionalLimitsApi.md#GetOauthConsumerTransactionCategoriesCategoryHashIdRule) | **Get** /oauth/consumer/transaction/categories/{category_hash_id}/rule | Retrieve rules for a particular transaction category
[**GetUsersWalletsCardsCardHashIdTransactionCategoryHashIdLimits**](CardsTransactionalLimitsApi.md#GetUsersWalletsCardsCardHashIdTransactionCategoryHashIdLimits) | **Get** /users/wallets/cards/{card_hash_id}/transactions/categories/{category_hash_id}/limit | Retrieve transactional limits for categories
[**GetUsersWalletsCardsCardHashIdTransactionsCategoriesLimit**](CardsTransactionalLimitsApi.md#GetUsersWalletsCardsCardHashIdTransactionsCategoriesLimit) | **Get** /users/wallets/cards/{card_hash_id}/transactions/categories/limit | Retrieve transactional limits for a card
[**PostOauthConsumerTransactionCategoriesCategoryHashIdRule**](CardsTransactionalLimitsApi.md#PostOauthConsumerTransactionCategoriesCategoryHashIdRule) | **Post** /oauth/consumer/transaction/categories/{category_hash_id}/rule | Create rules for a particular transaction category
[**PostUsersWalletsCardsCardHashIdTransactionCategoryHashIdLimits**](CardsTransactionalLimitsApi.md#PostUsersWalletsCardsCardHashIdTransactionCategoryHashIdLimits) | **Post** /users/wallets/cards/{card_hash_id}/transactions/categories/{category_hash_id}/limit | Configure transactional limits for a particular card and category
[**PostUsersWalletsCardsTransactionsCategoriesLimitsDetail**](CardsTransactionalLimitsApi.md#PostUsersWalletsCardsTransactionsCategoriesLimitsDetail) | **Post** /users/wallets/cards/transactions/categories/limits/detail | Create Transaction Category Limits
[**PutOauthConsumerTransactionCategoriesCategoryHashIdRule**](CardsTransactionalLimitsApi.md#PutOauthConsumerTransactionCategoriesCategoryHashIdRule) | **Put** /oauth/consumer/transaction/categories/{category_hash_id}/rule | Update rules for a particular transaction category
[**PutUsersWalletsCardsCardHashIdTransactionCategoryHashIdLimits**](CardsTransactionalLimitsApi.md#PutUsersWalletsCardsCardHashIdTransactionCategoryHashIdLimits) | **Put** /users/wallets/cards/{card_hash_id}/transactions/categories/{category_hash_id}/limit | Update transactional limits for a particular card and category

# **DeleteOauthConsumerTransactionCategoriesCategoryHashIdRule**
> InlineResponse2007 DeleteOauthConsumerTransactionCategoriesCategoryHashIdRule(ctx, categoryHashId)
Delete rules for a particular transaction category

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **categoryHashId** | **string**|  | 

### Return type

[**InlineResponse2007**](inline_response_200_7.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteUsersWalletsCardsCardHashIdTransactionCategoryHashIdLimits**
> UserLimitV1 DeleteUsersWalletsCardsCardHashIdTransactionCategoryHashIdLimits(ctx, cardHashId, categoryHashId, xAuthUserID)
Disable transactional limits for a particular card and category

Disable transactional limits for a particular card and category 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **cardHashId** | **string**|  | 
  **categoryHashId** | **string**| Transaction category Identifier | 
  **xAuthUserID** | **string**| User Identifier | 

### Return type

[**UserLimitV1**](User_limit.v1.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetOauthConsumerTransactionCategoriesCategoryHashIdRule**
> TransactionCategoriesRulesResponseV1 GetOauthConsumerTransactionCategoriesCategoryHashIdRule(ctx, categoryHashId)
Retrieve rules for a particular transaction category

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **categoryHashId** | **string**|  | 

### Return type

[**TransactionCategoriesRulesResponseV1**](Transaction_categories_rules_response.v1.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetUsersWalletsCardsCardHashIdTransactionCategoryHashIdLimits**
> UserLimitV1 GetUsersWalletsCardsCardHashIdTransactionCategoryHashIdLimits(ctx, cardHashId, categoryHashId, xAuthUserID)
Retrieve transactional limits for categories

Retrieve transactional limits for a card for a particular category

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **cardHashId** | **string**|  | 
  **categoryHashId** | **string**| Transaction category Identifier | 
  **xAuthUserID** | **string**| User Identifier | 

### Return type

[**UserLimitV1**](User_limit.v1.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetUsersWalletsCardsCardHashIdTransactionsCategoriesLimit**
> InlineResponse20053 GetUsersWalletsCardsCardHashIdTransactionsCategoriesLimit(ctx, cardHashId)
Retrieve transactional limits for a card

Retrieve transactional limits for a card categories

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **cardHashId** | **string**|  | 

### Return type

[**InlineResponse20053**](inline_response_200_53.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PostOauthConsumerTransactionCategoriesCategoryHashIdRule**
> TransactionCategoriesRulesResponseV1 PostOauthConsumerTransactionCategoriesCategoryHashIdRule(ctx, categoryHashId, optional)
Create rules for a particular transaction category

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **categoryHashId** | **string**|  | 
 **optional** | ***CardsTransactionalLimitsApiPostOauthConsumerTransactionCategoriesCategoryHashIdRuleOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a CardsTransactionalLimitsApiPostOauthConsumerTransactionCategoriesCategoryHashIdRuleOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **eventCodes** | **optional.**|  | 
 **channel** | **optional.**|  | 
 **acquiringCountryCode** | **optional.**|  | 
 **currency** | **optional.**|  | 
 **network** | **optional.**|  | 
 **terminalId** | **optional.**|  | 
 **merchantId** | **optional.**|  | 
 **posEntryModeValue** | **optional.**|  | 
 **pp** | **optional.**|  | 
 **merchantCategoryCode** | **optional.**|  | 

### Return type

[**TransactionCategoriesRulesResponseV1**](Transaction_categories_rules_response.v1.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: application/x-www-form-urlencoded
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PostUsersWalletsCardsCardHashIdTransactionCategoryHashIdLimits**
> UserLimitV1 PostUsersWalletsCardsCardHashIdTransactionCategoryHashIdLimits(ctx, cardHashId, categoryHashId, xAuthUserID, optional)
Configure transactional limits for a particular card and category

To set user limits for card

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **cardHashId** | **string**|  | 
  **categoryHashId** | **string**| Transaction category Identifier | 
  **xAuthUserID** | **string**| User Identifier | 
 **optional** | ***CardsTransactionalLimitsApiPostUsersWalletsCardsCardHashIdTransactionCategoryHashIdLimitsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a CardsTransactionalLimitsApiPostUsersWalletsCardsCardHashIdTransactionCategoryHashIdLimitsOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **enable** | **optional.**|  | 
 **values** | [**optional.Interface of []UserLimitV1CardLimits**](UserLimitV1CardLimits.md)|  | 

### Return type

[**UserLimitV1**](User_limit.v1.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: application/x-www-form-urlencoded
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PostUsersWalletsCardsTransactionsCategoriesLimitsDetail**
> PostUsersWalletsCardsTransactionsCategoriesLimitsDetail(ctx, optional)
Create Transaction Category Limits

Create a user wallet card transaction category limit 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***CardsTransactionalLimitsApiPostUsersWalletsCardsTransactionsCategoriesLimitsDetailOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a CardsTransactionalLimitsApiPostUsersWalletsCardsTransactionsCategoriesLimitsDetailOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **cardHashId** | **optional.**|  | 
 **userHashId** | **optional.**|  | 
 **transactionAmount** | **optional.**|  | 
 **amount** | **optional.**|  | 
 **cardCurrency** | **optional.**|  | 
 **transactionCurrency** | **optional.**|  | 
 **channel** | **optional.**|  | 
 **posEntryModeValue** | **optional.**|  | 
 **acquiringCountryCode** | **optional.**|  | 
 **network** | **optional.**|  | 
 **merchantCategoryCode** | **optional.**|  | 
 **merchantId** | **optional.**|  | 
 **terminalId** | **optional.**|  | 

### Return type

 (empty response body)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: application/x-www-form-urlencoded
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PutOauthConsumerTransactionCategoriesCategoryHashIdRule**
> TransactionCategoriesRulesResponseV1 PutOauthConsumerTransactionCategoriesCategoryHashIdRule(ctx, categoryHashId, optional)
Update rules for a particular transaction category

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **categoryHashId** | **string**|  | 
 **optional** | ***CardsTransactionalLimitsApiPutOauthConsumerTransactionCategoriesCategoryHashIdRuleOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a CardsTransactionalLimitsApiPutOauthConsumerTransactionCategoriesCategoryHashIdRuleOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **eventCodes** | **optional.**|  | 
 **channel** | **optional.**|  | 
 **acquiringCountryCode** | **optional.**|  | 
 **currency** | **optional.**|  | 
 **network** | **optional.**|  | 
 **terminalId** | **optional.**|  | 
 **merchantId** | **optional.**|  | 
 **posEntryModeValue** | **optional.**|  | 
 **pp** | **optional.**|  | 
 **merchantCategoryCode** | **optional.**|  | 

### Return type

[**TransactionCategoriesRulesResponseV1**](Transaction_categories_rules_response.v1.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: application/x-www-form-urlencoded
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PutUsersWalletsCardsCardHashIdTransactionCategoryHashIdLimits**
> UserLimitV1 PutUsersWalletsCardsCardHashIdTransactionCategoryHashIdLimits(ctx, cardHashId, categoryHashId, xAuthUserID, optional)
Update transactional limits for a particular card and category

To set user limits for card

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **cardHashId** | **string**|  | 
  **categoryHashId** | **string**| Transaction category Identifier | 
  **xAuthUserID** | **string**| User Identifier | 
 **optional** | ***CardsTransactionalLimitsApiPutUsersWalletsCardsCardHashIdTransactionCategoryHashIdLimitsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a CardsTransactionalLimitsApiPutUsersWalletsCardsCardHashIdTransactionCategoryHashIdLimitsOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **enable** | **optional.**|  | 
 **values** | [**optional.Interface of []UserLimitV1CardLimits**](UserLimitV1CardLimits.md)|  | 

### Return type

[**UserLimitV1**](User_limit.v1.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: application/x-www-form-urlencoded
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

