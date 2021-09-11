# {{classname}}

All URIs are relative to */*

Method | HTTP request | Description
------------- | ------------- | -------------
[**TransactionStatusGet**](TransactionApi.md#TransactionStatusGet) | **Get** /transaction/status | Get transaction status

# **TransactionStatusGet**
> TransactionStatus TransactionStatusGet(ctx, xAuthHeader, optional)
Get transaction status

Resource that will retrieve transaction status based on topup reference id, payment reference id, or client reference id

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **xAuthHeader** | **string**| Consumer ACL Key | [default to 717763e4d8384ff1a5e4c7b68d8032d7176b9f1168c180c3d6c9f045ab6bf95e]
 **optional** | ***TransactionApiTransactionStatusGetOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a TransactionApiTransactionStatusGetOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **paymentRefId** | **optional.String**| IH Relay Transaction Reference ID | 
 **clientRefId** | **optional.String**| Client Transaction Reference ID | 
 **providerRefId** | **optional.String**| Reference ID generated by Provider to help identify the transaction as needed in settlement reports and such | 
 **paymentId** | **optional.String**| Topup Reference ID | 
 **refId** | **optional.String**| (For Prefund Consumer) Topup Processor Reference ID | 

### Return type

[**TransactionStatus**](TransactionStatus.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)
