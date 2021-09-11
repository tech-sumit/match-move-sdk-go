# {{classname}}

All URIs are relative to *https://beta-api.mmvpay.com/{product}/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateUsers**](UsersApi.md#CreateUsers) | **Post** /users | Create/Re-activate user
[**DeleteUserId**](UsersApi.md#DeleteUserId) | **Delete** /users/{user_id} | De-activate user 
[**GetUser**](UsersApi.md#GetUser) | **Get** /users | Get user
[**UpdateUser**](UsersApi.md#UpdateUser) | **Put** /users | Update user

# **CreateUsers**
> UserFullV1 CreateUsers(ctx, id, firstName, lastName, middleName, preferredName, email, mobileCountryCode, password, mobile, title, idType, idNumber, idDateExpiry, idDateIssue, countryOfIssue, altIdType, altIdNumber, altIdDateExpiry, altIdDateIssue, altIdCountryOfIssue, birthday, placeOfBirth, nationality, gender, maritalStatus, mothersMaidenName, natureOfBusiness, customerId, purposeOfAccount, partnerId)
Create/Re-activate user

This endpoint used for creation of new user in the system. In addition when called just using `user_id` the same endpoint is used for re-activating the user in the system. If the user is in locked state.  The user information that needs to be colleted in the system depends on the country of issuance of the program. <!-- theme: info --> > ### Note > > Mobile number and email combination is unique individually in the system. For a given program if mobile number is already in use and that account is in active or locked state the same mobile number cannot be used for new account creation. It can only be re-used if the original account is in blocked state. Similarly for a given program if email is already in use and the account is in active or locked state the same email cannot be used for another user. It can only be re-used if the account is in blocked state.   <!-- theme: info --> > ### Note > > Do note when this endpoint is called in context of user re-activation only user_id is required.   <!-- theme: danger --> > ### Note > > Blocked accounts cannot be re-activated 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**|  | 
  **firstName** | **string**|  | 
  **lastName** | **string**|  | 
  **middleName** | **string**|  | 
  **preferredName** | **string**|  | 
  **email** | **string**|  | 
  **mobileCountryCode** | **string**|  | 
  **password** | **string**|  | 
  **mobile** | **string**|  | 
  **title** | **string**|  | 
  **idType** | **string**|  | 
  **idNumber** | **string**|  | 
  **idDateExpiry** | **string**|  | 
  **idDateIssue** | **string**|  | 
  **countryOfIssue** | **string**|  | 
  **altIdType** | **string**|  | 
  **altIdNumber** | **string**|  | 
  **altIdDateExpiry** | **string**|  | 
  **altIdDateIssue** | **string**|  | 
  **altIdCountryOfIssue** | **string**|  | 
  **birthday** | **string**|  | 
  **placeOfBirth** | **string**|  | 
  **nationality** | **string**|  | 
  **gender** | **string**|  | 
  **maritalStatus** | **string**|  | 
  **mothersMaidenName** | **string**|  | 
  **natureOfBusiness** | **string**|  | 
  **customerId** | **string**|  | 
  **purposeOfAccount** | **string**|  | 
  **partnerId** | **string**|  | 

### Return type

[**UserFullV1**](User_full.v1.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: application/x-www-form-urlencoded
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteUserId**
> InlineResponse200 DeleteUserId(ctx, userId, optional)
De-activate user 

There are two kinds of suspensions which are supported in the system currently. depending on the state passed as part of the request the behavior changes. ### Suspended   disable all resources of the user which includes the generation of the security token. This will also invalidate all active access associated to the user. If the card has been already pre-validated by a merchant, an authorization can still be made. ### Blocked   will suspend the account and block all cards associated to the user. The blocked cards will disable all transaction made by that card, and all purchase authorization. <!-- theme: info --> > ### Note > An account with cards or payment account with pending settlements cannot be closed. You are recommended to first suspend the account and wait for pending settlements before blocking 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **userId** | **string**|  | 
 **optional** | ***UsersApiDeleteUserIdOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a UsersApiDeleteUserIdOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **state** | **optional.**|  | 

### Return type

[**InlineResponse200**](inline_response_200.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: application/x-www-form-urlencoded
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetUser**
> UserFullV1 GetUser(ctx, xAuthUserID)
Get user

Retrieve details of a particular user created in the system 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **xAuthUserID** | **string**| MatchMove provided hash ID for the user | 

### Return type

[**UserFullV1**](User_full.v1.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateUser**
> UserFullV1 UpdateUser(ctx, id, firstName, lastName, middleName, preferredName, email, mobileCountryCode, password, mobile, title, idType, idNumber, idDateExpiry, idDateIssue, countryOfIssue, altIdType, altIdNumber, altIdDateExpiry, altIdDateIssue, altIdCountryOfIssue, birthday, placeOfBirth, nationality, gender, maritalStatus, mothersMaidenName, natureOfBusiness, customerId, purposeOfAccount, partnerId, xAuthUserID)
Update user

Update details of a user created in the system  The `preferred_name` can only be changed before wallet/card creation.  Modifying `first_name`, `middle_name`, `last_name`, `id_type`, `id_number` and/or `id_date_expiry` will result to the removal of any risk exception created for the user, thus, a new exception must be added using <a href=\"#post-users-risk-exceptions\">POST /users/risk/exceptions</a>.  This resource may function as a validation for the given parameters, provided appending `test` to the url (eg. users/test). 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**|  | 
  **firstName** | **string**|  | 
  **lastName** | **string**|  | 
  **middleName** | **string**|  | 
  **preferredName** | **string**|  | 
  **email** | **string**|  | 
  **mobileCountryCode** | **string**|  | 
  **password** | **string**|  | 
  **mobile** | **string**|  | 
  **title** | **string**|  | 
  **idType** | **string**|  | 
  **idNumber** | **string**|  | 
  **idDateExpiry** | **string**|  | 
  **idDateIssue** | **string**|  | 
  **countryOfIssue** | **string**|  | 
  **altIdType** | **string**|  | 
  **altIdNumber** | **string**|  | 
  **altIdDateExpiry** | **string**|  | 
  **altIdDateIssue** | **string**|  | 
  **altIdCountryOfIssue** | **string**|  | 
  **birthday** | **string**|  | 
  **placeOfBirth** | **string**|  | 
  **nationality** | **string**|  | 
  **gender** | **string**|  | 
  **maritalStatus** | **string**|  | 
  **mothersMaidenName** | **string**|  | 
  **natureOfBusiness** | **string**|  | 
  **customerId** | **string**|  | 
  **purposeOfAccount** | **string**|  | 
  **partnerId** | **string**|  | 
  **xAuthUserID** | **string**| MatchMove provided hash ID for the user | 

### Return type

[**UserFullV1**](User_full.v1.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: application/x-www-form-urlencoded
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

