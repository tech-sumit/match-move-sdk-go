# {{classname}}

All URIs are relative to *https://beta-api.mmvpay.com/{product}/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeleteOauthConsumerConsumerKeyWebhooksWebhokId**](WebhooksApi.md#DeleteOauthConsumerConsumerKeyWebhooksWebhokId) | **Delete** /oauth/consumer/{consumer_key}/webhooks/{webhook_id} | Delete consumer webhook
[**GetOauthConsumerConsumerKeyWebhooks**](WebhooksApi.md#GetOauthConsumerConsumerKeyWebhooks) | **Get** /oauth/consumer/{consumer_key}/webhooks | Retrieve list of webhooks
[**GetOauthConsumerConsumerKeyWebhooksWebhokId**](WebhooksApi.md#GetOauthConsumerConsumerKeyWebhooksWebhokId) | **Get** /oauth/consumer/{consumer_key}/webhooks/{webhook_id} | Retrieve details of a particular webhook
[**GetOauthConsumerKeyEventCategories**](WebhooksApi.md#GetOauthConsumerKeyEventCategories) | **Get** /oauth/consumer/{consumer_key}/event/categories | List available webhook categories
[**PostOauthConsumerConsumerKeyWebhooks**](WebhooksApi.md#PostOauthConsumerConsumerKeyWebhooks) | **Post** /oauth/consumer/{consumer_key}/webhooks | Create a webhook for your consumer
[**PutOauthConsumerConsumerKeyWebhooksWebhokId**](WebhooksApi.md#PutOauthConsumerConsumerKeyWebhooksWebhokId) | **Put** /oauth/consumer/{consumer_key}/webhooks/{webhook_id} | Update details of a particular webhook

# **DeleteOauthConsumerConsumerKeyWebhooksWebhokId**
> []WebhooksV1Inner DeleteOauthConsumerConsumerKeyWebhooksWebhokId(ctx, consumerKey, webhookId)
Delete consumer webhook

Delete webhook

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **consumerKey** | **string**|  | 
  **webhookId** | **string**|  | 

### Return type

[**[]WebhooksV1Inner**](array.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetOauthConsumerConsumerKeyWebhooks**
> []WebhooksV1Inner GetOauthConsumerConsumerKeyWebhooks(ctx, consumerKey)
Retrieve list of webhooks

Retrieve the list of webhooks

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **consumerKey** | **string**|  | 

### Return type

[**[]WebhooksV1Inner**](array.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetOauthConsumerConsumerKeyWebhooksWebhokId**
> []InlineResponse20034 GetOauthConsumerConsumerKeyWebhooksWebhokId(ctx, consumerKey, webhookId)
Retrieve details of a particular webhook

Get a webhook

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **consumerKey** | **string**|  | 
  **webhookId** | **string**|  | 

### Return type

[**[]InlineResponse20034**](inline_response_200_34.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetOauthConsumerKeyEventCategories**
> []InlineResponse20034 GetOauthConsumerKeyEventCategories(ctx, consumerKey)
List available webhook categories

List available webhook categories available in the system 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **consumerKey** | **string**|  | 

### Return type

[**[]InlineResponse20034**](inline_response_200_34.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PostOauthConsumerConsumerKeyWebhooks**
> []WebhooksV1Inner PostOauthConsumerConsumerKeyWebhooks(ctx, consumerKey, optional)
Create a webhook for your consumer

Create a webhook entry

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **consumerKey** | **string**|  | 
 **optional** | ***WebhooksApiPostOauthConsumerConsumerKeyWebhooksOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a WebhooksApiPostOauthConsumerConsumerKeyWebhooksOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **url** | **optional.**|  | 
 **eventHash** | **optional.**|  | 
 **categoryHash** | **optional.**|  | 

### Return type

[**[]WebhooksV1Inner**](array.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: application/x-www-form-urlencoded
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PutOauthConsumerConsumerKeyWebhooksWebhokId**
> []WebhooksV1Inner PutOauthConsumerConsumerKeyWebhooksWebhokId(ctx, consumerKey, webhookId, optional)
Update details of a particular webhook

Update webhook

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **consumerKey** | **string**|  | 
  **webhookId** | **string**|  | 
 **optional** | ***WebhooksApiPutOauthConsumerConsumerKeyWebhooksWebhokIdOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a WebhooksApiPutOauthConsumerConsumerKeyWebhooksWebhokIdOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **url** | **optional.**|  | 

### Return type

[**[]WebhooksV1Inner**](array.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: application/x-www-form-urlencoded
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

