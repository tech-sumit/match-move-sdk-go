# {{classname}}

All URIs are relative to *https://beta-api.mmvpay.com/{product}/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeleteOauthConsumerFunds**](OauthUtilitiesApi.md#DeleteOauthConsumerFunds) | **Delete** /oauth/consumer/funds | Delete Oauth Consumer Funds
[**PostOauthConsumerFunds**](OauthUtilitiesApi.md#PostOauthConsumerFunds) | **Post** /oauth/consumer/funds | Create Oauth Consumer Funds

# **DeleteOauthConsumerFunds**
> InlineResponse20039 DeleteOauthConsumerFunds(ctx, optional)
Delete Oauth Consumer Funds

Confirm a cloesd loop debit transaction 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***OauthUtilitiesApiDeleteOauthConsumerFundsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a OauthUtilitiesApiDeleteOauthConsumerFundsOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ids** | **optional.**|  | 

### Return type

[**InlineResponse20039**](inline_response_200_39.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: application/x-www-form-urlencoded
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PostOauthConsumerFunds**
> InlineResponse20038 PostOauthConsumerFunds(ctx, optional)
Create Oauth Consumer Funds

Confirm a top up transaction given its ref_id

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***OauthUtilitiesApiPostOauthConsumerFundsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a OauthUtilitiesApiPostOauthConsumerFundsOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ids** | **optional.**|  | 

### Return type

[**InlineResponse20038**](inline_response_200_38.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: application/x-www-form-urlencoded
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

