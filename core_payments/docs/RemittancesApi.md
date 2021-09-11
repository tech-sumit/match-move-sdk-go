# {{classname}}

All URIs are relative to *https://beta-api.mmvpay.com/{product}/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetOauthConsumersFundsTransfersOverseasAttachmentForm**](RemittancesApi.md#GetOauthConsumersFundsTransfersOverseasAttachmentForm) | **Get** /oauth/consumers/funds/transfers/overseas/attachment/form | Get Form for Remittance Attachments
[**GetOauthConsumersFundsTransfersOverseasBeneficiaryForm**](RemittancesApi.md#GetOauthConsumersFundsTransfersOverseasBeneficiaryForm) | **Get** /oauth/consumer/funds/transfers/overseas/beneficiary/form | Remittance Overseas Beneficiary Form
[**GetOauthConsumersFundsTransfersOverseasCountries**](RemittancesApi.md#GetOauthConsumersFundsTransfersOverseasCountries) | **Get** /oauth/consumer/funds/transfers/overseas/countries | Remittance Countries
[**GetOauthConsumersFundsTransfersOverseasModes**](RemittancesApi.md#GetOauthConsumersFundsTransfersOverseasModes) | **Get** /oauth/consumer/funds/transfers/overseas/modes | Remittance Payment Modes
[**GetUsersWalletsFundsTransfersOverseasHistoryLimitOffset**](RemittancesApi.md#GetUsersWalletsFundsTransfersOverseasHistoryLimitOffset) | **Get** /users/wallets/funds/transfers/overseas/history/{limit}]/{offset} | Get Remittance History
[**GetUsersWalletsFundsTransfersOverseasIdAttachments**](RemittancesApi.md#GetUsersWalletsFundsTransfersOverseasIdAttachments) | **Get** /users/wallets/funds/transfers/overseas/{id}/attachments | Get Uploaded Remittance Attachments
[**GetUsersWalletsFundsTransfersOverseasProviders**](RemittancesApi.md#GetUsersWalletsFundsTransfersOverseasProviders) | **Get** /users/wallets/funds/transfers/overseas/providers/{limit}]/{offset} | Remittance Provider Agents
[**GetUsersWalletsOverseasTransfersFees**](RemittancesApi.md#GetUsersWalletsOverseasTransfersFees) | **Get** /users/wallets/funds/transfers/overseas/{type}/fees/{limit}/{offset} | Get Remittance Fees
[**GetUsersWalletsOverseasTransfersTypes**](RemittancesApi.md#GetUsersWalletsOverseasTransfersTypes) | **Get** /users/wallets/funds/transfers/overseas/types/{code} | Get Remittance Providers
[**PostOauthConsumerFundsTransfersOverseas**](RemittancesApi.md#PostOauthConsumerFundsTransfersOverseas) | **Post** /oauth/consumer/funds/transfers/overseas | Confirm Remittance transaction
[**PostUsersWalletsFundsTransfersOverseas**](RemittancesApi.md#PostUsersWalletsFundsTransfersOverseas) | **Post** /users/wallets/funds/transfers/overseas | Send Overseas Transfers without providing Rails
[**PostUsersWalletsFundsTransfersOverseasIdAttachment**](RemittancesApi.md#PostUsersWalletsFundsTransfersOverseasIdAttachment) | **Post** /users/wallets/funds/transfers/overseas/{id}/attachment | Upload Remittance Attachment
[**PostUsersWalletsFundsTransfersOverseasTest**](RemittancesApi.md#PostUsersWalletsFundsTransfersOverseasTest) | **Post** /users/wallets/funds/transfers/overseas/test | Calculate Fees for Overseas Transfers without providing Rails
[**PostUsersWalletsFundsTransfersOverseasType**](RemittancesApi.md#PostUsersWalletsFundsTransfersOverseasType) | **Post** /users/wallets/funds/transfers/overseas/{type} | Calculate or Create remittance transaction
[**PostUsersWalletsFundsTransfersOverseasTypeTest**](RemittancesApi.md#PostUsersWalletsFundsTransfersOverseasTypeTest) | **Post** /users/wallets/funds/transfers/overseas/{type}/test | Calculate Rates for Remittance Transaction

# **GetOauthConsumersFundsTransfersOverseasAttachmentForm**
> RemittanceAttachmentFormV1 GetOauthConsumersFundsTransfersOverseasAttachmentForm(ctx, payoutAgent)
Get Form for Remittance Attachments

Retrieve the attachment form details and supported document types to attach

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **payoutAgent** | **string**| Payout agent hash id | 

### Return type

[**RemittanceAttachmentFormV1**](Remittance_attachment_form.v1.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetOauthConsumersFundsTransfersOverseasBeneficiaryForm**
> InlineResponse20055 GetOauthConsumersFundsTransfersOverseasBeneficiaryForm(ctx, optional)
Remittance Overseas Beneficiary Form

Retrieve beneficiary form of a given provider and payout agent

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***RemittancesApiGetOauthConsumersFundsTransfersOverseasBeneficiaryFormOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a RemittancesApiGetOauthConsumersFundsTransfersOverseasBeneficiaryFormOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **provider** | **optional.String**| Provider code from &#x60;GET /users/wallets/funds/transfers/overseas/types&#x60; | 
 **type_** | **optional.String**| Beneficiary Type | 
 **payoutAgent** | **optional.String**| Payout Agent | 

### Return type

[**InlineResponse20055**](inline_response_200_55.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetOauthConsumersFundsTransfersOverseasCountries**
> RemittanceCountriesV1 GetOauthConsumersFundsTransfersOverseasCountries(ctx, optional)
Remittance Countries

Retrieve all countries supported by the program for remittance

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***RemittancesApiGetOauthConsumersFundsTransfersOverseasCountriesOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a RemittancesApiGetOauthConsumersFundsTransfersOverseasCountriesOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **provider** | **optional.String**| Provider code from &#x60;GET /users/wallets/funds/transfers/overseas/types&#x60; | 
 **type_** | **optional.String**| Beneficiary Type | 
 **name** | **optional.String**| Country Name | 
 **code** | **optional.String**| Country Code | 

### Return type

[**RemittanceCountriesV1**](Remittance_countries.v1.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetOauthConsumersFundsTransfersOverseasModes**
> RemittancePaymentModesV1 GetOauthConsumersFundsTransfersOverseasModes(ctx, optional)
Remittance Payment Modes

Retrieve all modes with specified countries per mode supported by the program for remittance

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***RemittancesApiGetOauthConsumersFundsTransfersOverseasModesOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a RemittancesApiGetOauthConsumersFundsTransfersOverseasModesOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **provider** | **optional.String**| Provider code from &#x60;GET /users/wallets/funds/transfers/overseas/types&#x60; | 
 **type_** | **optional.String**| Beneficiary Type | 
 **name** | **optional.String**| Payment channel name | 
 **code** | **optional.String**| Payment channel code | 
 **countryCode** | **optional.String**| Country Code | 

### Return type

[**RemittancePaymentModesV1**](Remittance_payment_modes.v1.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetUsersWalletsFundsTransfersOverseasHistoryLimitOffset**
> InlineResponse20050 GetUsersWalletsFundsTransfersOverseasHistoryLimitOffset(ctx, limit, offset, optional)
Get Remittance History

Get Remittance history against the wallet of the user 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **limit** | **string**|  | 
  **offset** | **string**|  | 
 **optional** | ***RemittancesApiGetUsersWalletsFundsTransfersOverseasHistoryLimitOffsetOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a RemittancesApiGetUsersWalletsFundsTransfersOverseasHistoryLimitOffsetOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **sort** | **optional.String**| Comma delimited sorting of the result according to id and date.added | 
 **deliveryMethod** | **optional.String**| Filter transactions based on the delivery method | 
 **provider** | **optional.String**| Filter transactions based on the [provider](#get-users-wallets-funds-transfers-overseas-types) | 
 **dateAdded** | **optional.String**| Filter transactions based on the date added | 
 **dateExpiry** | **optional.String**| Filter transactions based on the date expiry | 
 **status** | **optional.String**| Filter transactions based status | 
 **search** | **optional.String**| Filter transactions based on the transaction code , reference id, or transfer reference id | 
 **senderFirstName** | **optional.String**| Filter transactions based on the sender first name | 
 **senderMiddleName** | **optional.String**| Filter transactions based on the sender middle name | 
 **senderLastName** | **optional.String**| Filter transactions based on the sender last name | 
 **senderMobileNumber** | **optional.String**| Filter transactions based on the sender mobile number | 
 **receiverFirstName** | **optional.String**| Filter transactions based on the beneficiary first name | 
 **receiverMiddleName** | **optional.String**| Filter transactions based on the beneficiary middle name | 
 **receiverLastName** | **optional.String**| Filter transactions based on the beneficiary last name | 
 **businessRegisteredName** | **optional.String**| Filter transactions based on the corporate registered name | 
 **beneficiaryHashId** | **optional.String**| Filter transactions based on the corporate/consumer beneficiary hash id | 
 **clientRefId** | **optional.String**| Filter transactions based on the client reference id | 

### Return type

[**InlineResponse20050**](inline_response_200_50.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetUsersWalletsFundsTransfersOverseasIdAttachments**
> RemittanceAttachmentsV1 GetUsersWalletsFundsTransfersOverseasIdAttachments(ctx, id)
Get Uploaded Remittance Attachments

Retrieve the attached documents against the given remittance hash id

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**| Remittance Hash ID | 

### Return type

[**RemittanceAttachmentsV1**](Remittance_attachments.v1.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetUsersWalletsFundsTransfersOverseasProviders**
> InlineResponse20054 GetUsersWalletsFundsTransfersOverseasProviders(ctx, limit, offset, optional)
Remittance Provider Agents

Retrieve all the payout agents supported on the system for remittance.  Data retrieved for this endpoint will be getting from a cache object which is populated on a background task. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **limit** | **int32**| Pagination limit | 
  **offset** | **int32**| Pagination offset | 
 **optional** | ***RemittancesApiGetUsersWalletsFundsTransfersOverseasProvidersOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a RemittancesApiGetUsersWalletsFundsTransfersOverseasProvidersOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **type_** | **optional.String**| Beneficiary type | 
 **agentId** | **optional.String**| Payout Agent ID | 
 **receiveCountry** | **optional.String**| Receiving Country Code (ISO alpha-3) | 
 **mode** | **optional.String**| Mode from response of GET oauth/consumer/funds/transfers/overseas/modes | 
 **provider** | **optional.String**| Remittance Rails | 

### Return type

[**InlineResponse20054**](inline_response_200_54.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetUsersWalletsOverseasTransfersFees**
> RemittanceFeesV1 GetUsersWalletsOverseasTransfersFees(ctx, type_, limit, offset)
Get Remittance Fees

Retrieve list of fees for remittance agents configured in the system.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **type_** | **string**|  | 
  **limit** | **string**|  | 
  **offset** | **string**|  | 

### Return type

[**RemittanceFeesV1**](Remittance_fees.v1.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetUsersWalletsOverseasTransfersTypes**
> RemittanceProvidersV1 GetUsersWalletsOverseasTransfersTypes(ctx, code)
Get Remittance Providers

Lists all overseas providers offered. You may retrieve information for a specific provider by appending its code. This resource serves as a prerequisite to list rates, payers and destinations. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **code** | **string**|  | 

### Return type

[**RemittanceProvidersV1**](Remittance_providers.v1.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PostOauthConsumerFundsTransfersOverseas**
> InlineResponse20051 PostOauthConsumerFundsTransfersOverseas(ctx, ids)
Confirm Remittance transaction

Confirm a remittance transaction , given the id(s)

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **ids** | **string**|  | 

### Return type

[**InlineResponse20051**](inline_response_200_51.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: application/x-www-form-urlencoded
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PostUsersWalletsFundsTransfersOverseas**
> RemittanceSendV1 PostUsersWalletsFundsTransfersOverseas(ctx, optional)
Send Overseas Transfers without providing Rails

Send a money transfer request with the specified agent.  Any details required on this endpoint will be based from the sender/beneficiary form on the calculate (/test) endpoint.  Any transaction made through this resource will be on pending validation. Confirmation must be made to clear the transaction. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***RemittancesApiPostUsersWalletsFundsTransfersOverseasOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a RemittancesApiPostUsersWalletsFundsTransfersOverseasOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **agentId** | **optional.**|  | 
 **calculationMode** | **optional.**|  | 
 **transactionType** | **optional.**|  | 
 **amount** | **optional.**|  | 
 **sendCurrency** | **optional.**|  | 
 **receiveCurrency** | **optional.**|  | 
 **receiveCountry** | **optional.**|  | 
 **paymentMode** | **optional.**|  | 
 **routingType** | **optional.**|  | 
 **routingParam** | **optional.**|  | 
 **quotationId** | **optional.**|  | 
 **payoutAgent** | **optional.**|  | 
 **hashId** | **optional.**|  | 
 **firstName** | **optional.**|  | 
 **middleName** | **optional.**|  | 
 **lastName** | **optional.**|  | 
 **chineseName** | **optional.**|  | 
 **receiveMobileNumber** | **optional.**|  | 
 **receiveGender** | **optional.**|  | 
 **idNumber** | **optional.**|  | 
 **idType** | **optional.**|  | 
 **occupation** | **optional.**|  | 
 **address1** | **optional.**|  | 
 **address2** | **optional.**|  | 
 **birthDate** | **optional.**|  | 
 **city** | **optional.**|  | 
 **state** | **optional.**|  | 
 **country** | **optional.**|  | 
 **nationality** | **optional.**|  | 
 **zipcode** | **optional.**|  | 
 **beneficiaryHashId** | **optional.**|  | 
 **receivingBusinessRegisteredName** | **optional.**|  | 
 **receivingBusinessTradingName** | **optional.**|  | 
 **receivingBusinessEmail** | **optional.**|  | 
 **receivingBusinessMsisdn** | **optional.**|  | 
 **receivingBusinessRegistrationNumber** | **optional.**|  | 
 **receivingBusinessTaxId** | **optional.**|  | 
 **receivingBusinessAddress1** | **optional.**|  | 
 **receivingBusinessAddress2** | **optional.**|  | 
 **receivingBusinessCity** | **optional.**|  | 
 **receivingBusinessState** | **optional.**|  | 
 **receivingBusinessCountry** | **optional.**|  | 
 **receivingBusinessRepresentativeFirstName** | **optional.**|  | 
 **receivingBusinessRepresentativeMiddleName** | **optional.**|  | 
 **receivingBusinessRepresentativeLastName** | **optional.**|  | 
 **receivingBusinessRepresentativeLastName2** | **optional.**|  | 
 **receivingBusinessRepresentativeNativeName** | **optional.**|  | 
 **receivingBusinessIdType** | **optional.**|  | 
 **receivingBusinessIdNumber** | **optional.**|  | 
 **receivingBusinessCountryOfIssue** | **optional.**|  | 
 **receivingBusinessIssueDate** | **optional.**|  | 
 **receivingBusinessExpireDate** | **optional.**|  | 
 **receivingBusinessDateOfIncorporation** | **optional.**|  | 
 **sendingBusinessHashId** | **optional.**|  | 
 **sendingBusinessRegisteredName** | **optional.**|  | 
 **sendingBusinessTradingName** | **optional.**|  | 
 **sendingBusinessEmail** | **optional.**|  | 
 **sendingBusinessMsisdn** | **optional.**|  | 
 **sendingBusinessRegistrationNumber** | **optional.**|  | 
 **sendingBusinessTaxId** | **optional.**|  | 
 **sendingBusinessAddress1** | **optional.**|  | 
 **sendingBusinessAddress2** | **optional.**|  | 
 **sendingBusinessCity** | **optional.**|  | 
 **sendingBusinessState** | **optional.**|  | 
 **sendingBusinessZipcode** | **optional.**|  | 
 **sendingBusinessCountry** | **optional.**|  | 
 **sendingBusinessRepresentativeFirstName** | **optional.**|  | 
 **sendingBusinessRepresentativeMiddleName** | **optional.**|  | 
 **sendingBusinessRepresentativeLastName** | **optional.**|  | 
 **sendingBusinessRepresentativeLastName2** | **optional.**|  | 
 **sendingBusinessRepresentativeNativeName** | **optional.**|  | 
 **sendingBusinessIdType** | **optional.**|  | 
 **sendingBusinessIdNumber** | **optional.**|  | 
 **sendingBusinessCountryOfIssue** | **optional.**|  | 
 **sendingBusinessIssueDate** | **optional.**|  | 
 **sendingBusinessExpireDate** | **optional.**|  | 
 **sendingBusinessDateOfIncorporation** | **optional.**|  | 
 **documentReferenceNumber** | **optional.**|  | 
 **swiftCode** | **optional.**|  | 
 **bankIfcCode** | **optional.**|  | 
 **bankBranchName** | **optional.**|  | 
 **bankBranchCode** | **optional.**|  | 
 **bankAccountNumber** | **optional.**|  | 
 **cbcBankName** | **optional.**|  | 
 **bpiBankName** | **optional.**|  | 
 **partnerName** | **optional.**|  | 
 **sourceOfIncome** | **optional.**|  | 
 **relationship** | **optional.**|  | 
 **reason** | **optional.**|  | 
 **senderReference** | **optional.**|  | 
 **additionalInformation** | **optional.**|  | 

### Return type

[**RemittanceSendV1**](Remittance_send.v1.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: application/x-www-form-urlencoded
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PostUsersWalletsFundsTransfersOverseasIdAttachment**
> RemittanceAttachmentV1 PostUsersWalletsFundsTransfersOverseasIdAttachment(ctx, data, documentType, documentName, xAuthUserID, id)
Upload Remittance Attachment

Attaches document on the given remittance hash id

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **data** | **string**|  | 
  **documentType** | **string**|  | 
  **documentName** | **string**|  | 
  **xAuthUserID** | **string**| MatchMove provided hash ID for the user | 
  **id** | **string**| Remittance Hash ID retrieved after (#post-users-wallets-funds-transfers-overseas-type) | 

### Return type

[**RemittanceAttachmentV1**](Remittance_attachment.v1.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: application/x-www-form-urlencoded
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PostUsersWalletsFundsTransfersOverseasTest**
> RemittanceCalculateV1 PostUsersWalletsFundsTransfersOverseasTest(ctx, optional)
Calculate Fees for Overseas Transfers without providing Rails

Calculate the exchange rate and fees applicable for the provided agent before an actual remittance transaction is created.  A beneficiary form is also returned as part of the validation process when creating the actual remittance transaction.  No transaction record is created yet when this endpoint is called. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***RemittancesApiPostUsersWalletsFundsTransfersOverseasTestOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a RemittancesApiPostUsersWalletsFundsTransfersOverseasTestOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **agentId** | **optional.**|  | 
 **calculationMode** | **optional.**|  | 
 **amount** | **optional.**|  | 

### Return type

[**RemittanceCalculateV1**](Remittance_calculate.v1.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: application/x-www-form-urlencoded
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PostUsersWalletsFundsTransfersOverseasType**
> InlineResponse20052 PostUsersWalletsFundsTransfersOverseasType(ctx, type_, optional)
Calculate or Create remittance transaction

Send a money transfer request with the specified overseas provider.  Any transaction made through this resource will be on pending validation. Confirmation must be made to clear the transaction.  Appending `test` to the (users/wallets/funds/transfers/overseas/{type}/test) sends a calculation test (ONLY) which useful for retrieving the fees applicable before creating a real transaction.  With administrative scope, this api can be accessed as a public resource. The additional parameter `hash_id` must be supplied when using this as a public resource. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **type_** | **string**| Overseas Provider Code | 
 **optional** | ***RemittancesApiPostUsersWalletsFundsTransfersOverseasTypeOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a RemittancesApiPostUsersWalletsFundsTransfersOverseasTypeOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **agentId** | **optional.**|  | 
 **calculationMode** | **optional.**|  | 
 **transactionType** | **optional.**|  | 
 **amount** | **optional.**|  | 
 **sendCurrency** | **optional.**|  | 
 **receiveCurrency** | **optional.**|  | 
 **receiveCountry** | **optional.**|  | 
 **paymentMode** | **optional.**|  | 
 **routingType** | **optional.**|  | 
 **routingParam** | **optional.**|  | 
 **quotationId** | **optional.**|  | 
 **payoutAgent** | **optional.**|  | 
 **hashId** | **optional.**|  | 
 **firstName** | **optional.**|  | 
 **middleName** | **optional.**|  | 
 **lastName** | **optional.**|  | 
 **chineseName** | **optional.**|  | 
 **receiveMobileNumber** | **optional.**|  | 
 **receiveGender** | **optional.**|  | 
 **idNumber** | **optional.**|  | 
 **idType** | **optional.**|  | 
 **occupation** | **optional.**|  | 
 **address1** | **optional.**|  | 
 **address2** | **optional.**|  | 
 **birthDate** | **optional.**|  | 
 **city** | **optional.**|  | 
 **state** | **optional.**|  | 
 **country** | **optional.**|  | 
 **nationality** | **optional.**|  | 
 **zipcode** | **optional.**|  | 
 **beneficiaryHashId** | **optional.**|  | 
 **receivingBusinessRegisteredName** | **optional.**|  | 
 **receivingBusinessTradingName** | **optional.**|  | 
 **receivingBusinessEmail** | **optional.**|  | 
 **receivingBusinessMsisdn** | **optional.**|  | 
 **receivingBusinessRegistrationNumber** | **optional.**|  | 
 **receivingBusinessTaxId** | **optional.**|  | 
 **receivingBusinessAddress1** | **optional.**|  | 
 **receivingBusinessAddress2** | **optional.**|  | 
 **receivingBusinessCity** | **optional.**|  | 
 **receivingBusinessState** | **optional.**|  | 
 **receivingBusinessCountry** | **optional.**|  | 
 **receivingBusinessRepresentativeFirstName** | **optional.**|  | 
 **receivingBusinessRepresentativeMiddleName** | **optional.**|  | 
 **receivingBusinessRepresentativeLastName** | **optional.**|  | 
 **receivingBusinessRepresentativeLastName2** | **optional.**|  | 
 **receivingBusinessRepresentativeNativeName** | **optional.**|  | 
 **receivingBusinessIdType** | **optional.**|  | 
 **receivingBusinessIdNumber** | **optional.**|  | 
 **receivingBusinessCountryOfIssue** | **optional.**|  | 
 **receivingBusinessIssueDate** | **optional.**|  | 
 **receivingBusinessExpireDate** | **optional.**|  | 
 **receivingBusinessDateOfIncorporation** | **optional.**|  | 
 **sendingBusinessHashId** | **optional.**|  | 
 **sendingBusinessRegisteredName** | **optional.**|  | 
 **sendingBusinessTradingName** | **optional.**|  | 
 **sendingBusinessEmail** | **optional.**|  | 
 **sendingBusinessMsisdn** | **optional.**|  | 
 **sendingBusinessRegistrationNumber** | **optional.**|  | 
 **sendingBusinessTaxId** | **optional.**|  | 
 **sendingBusinessAddress1** | **optional.**|  | 
 **sendingBusinessAddress2** | **optional.**|  | 
 **sendingBusinessCity** | **optional.**|  | 
 **sendingBusinessState** | **optional.**|  | 
 **sendingBusinessZipcode** | **optional.**|  | 
 **sendingBusinessCountry** | **optional.**|  | 
 **sendingBusinessRepresentativeFirstName** | **optional.**|  | 
 **sendingBusinessRepresentativeMiddleName** | **optional.**|  | 
 **sendingBusinessRepresentativeLastName** | **optional.**|  | 
 **sendingBusinessRepresentativeLastName2** | **optional.**|  | 
 **sendingBusinessRepresentativeNativeName** | **optional.**|  | 
 **sendingBusinessIdType** | **optional.**|  | 
 **sendingBusinessIdNumber** | **optional.**|  | 
 **sendingBusinessCountryOfIssue** | **optional.**|  | 
 **sendingBusinessIssueDate** | **optional.**|  | 
 **sendingBusinessExpireDate** | **optional.**|  | 
 **sendingBusinessDateOfIncorporation** | **optional.**|  | 
 **documentReferenceNumber** | **optional.**|  | 
 **swiftCode** | **optional.**|  | 
 **bankIfcCode** | **optional.**|  | 
 **bankBranchName** | **optional.**|  | 
 **bankBranchCode** | **optional.**|  | 
 **bankAccountNumber** | **optional.**|  | 
 **cbcBankName** | **optional.**|  | 
 **bpiBankName** | **optional.**|  | 
 **partnerName** | **optional.**|  | 
 **sourceOfIncome** | **optional.**|  | 
 **relationship** | **optional.**|  | 
 **reason** | **optional.**|  | 
 **senderReference** | **optional.**|  | 
 **additionalInformation** | **optional.**|  | 

### Return type

[**InlineResponse20052**](inline_response_200_52.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: application/x-www-form-urlencoded
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PostUsersWalletsFundsTransfersOverseasTypeTest**
> InlineResponse20052 PostUsersWalletsFundsTransfersOverseasTypeTest(ctx, type_, optional)
Calculate Rates for Remittance Transaction

Send a money transfer request with the specified overseas provider.  Any transaction made through this resource will be on pending validation. Confirmation must be made to clear the transaction.  Appending `test` to the (users/wallets/funds/transfers/overseas/{type}/test) sends a calculation test (ONLY) which is useful for retrieving the fees applicable before creating a real transaction.  With administrative scope, this api can be accessed as a public resource. The additional parameter `hash_id` must be supplied when using this as a public resource. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **type_** | **string**| Overseas Provider Code | 
 **optional** | ***RemittancesApiPostUsersWalletsFundsTransfersOverseasTypeTestOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a RemittancesApiPostUsersWalletsFundsTransfersOverseasTypeTestOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **agentId** | **optional.**|  | 
 **calculationMode** | **optional.**|  | 
 **transactionType** | **optional.**|  | 
 **amount** | **optional.**|  | 
 **sendCurrency** | **optional.**|  | 
 **receiveCurrency** | **optional.**|  | 
 **receiveCountry** | **optional.**|  | 
 **paymentMode** | **optional.**|  | 
 **routingType** | **optional.**|  | 
 **routingParam** | **optional.**|  | 
 **quotationId** | **optional.**|  | 
 **payoutAgent** | **optional.**|  | 
 **hashId** | **optional.**|  | 
 **firstName** | **optional.**|  | 
 **middleName** | **optional.**|  | 
 **lastName** | **optional.**|  | 
 **chineseName** | **optional.**|  | 
 **receiveMobileNumber** | **optional.**|  | 
 **receiveGender** | **optional.**|  | 
 **idNumber** | **optional.**|  | 
 **idType** | **optional.**|  | 
 **occupation** | **optional.**|  | 
 **address1** | **optional.**|  | 
 **address2** | **optional.**|  | 
 **birthDate** | **optional.**|  | 
 **city** | **optional.**|  | 
 **state** | **optional.**|  | 
 **country** | **optional.**|  | 
 **nationality** | **optional.**|  | 
 **zipcode** | **optional.**|  | 
 **beneficiaryHashId** | **optional.**|  | 
 **receivingBusinessRegisteredName** | **optional.**|  | 
 **receivingBusinessTradingName** | **optional.**|  | 
 **receivingBusinessEmail** | **optional.**|  | 
 **receivingBusinessMsisdn** | **optional.**|  | 
 **receivingBusinessRegistrationNumber** | **optional.**|  | 
 **receivingBusinessTaxId** | **optional.**|  | 
 **receivingBusinessAddress1** | **optional.**|  | 
 **receivingBusinessAddress2** | **optional.**|  | 
 **receivingBusinessCity** | **optional.**|  | 
 **receivingBusinessState** | **optional.**|  | 
 **receivingBusinessCountry** | **optional.**|  | 
 **receivingBusinessRepresentativeFirstName** | **optional.**|  | 
 **receivingBusinessRepresentativeMiddleName** | **optional.**|  | 
 **receivingBusinessRepresentativeLastName** | **optional.**|  | 
 **receivingBusinessRepresentativeLastName2** | **optional.**|  | 
 **receivingBusinessRepresentativeNativeName** | **optional.**|  | 
 **receivingBusinessIdType** | **optional.**|  | 
 **receivingBusinessIdNumber** | **optional.**|  | 
 **receivingBusinessCountryOfIssue** | **optional.**|  | 
 **receivingBusinessIssueDate** | **optional.**|  | 
 **receivingBusinessExpireDate** | **optional.**|  | 
 **receivingBusinessDateOfIncorporation** | **optional.**|  | 
 **sendingBusinessHashId** | **optional.**|  | 
 **sendingBusinessRegisteredName** | **optional.**|  | 
 **sendingBusinessTradingName** | **optional.**|  | 
 **sendingBusinessEmail** | **optional.**|  | 
 **sendingBusinessMsisdn** | **optional.**|  | 
 **sendingBusinessRegistrationNumber** | **optional.**|  | 
 **sendingBusinessTaxId** | **optional.**|  | 
 **sendingBusinessAddress1** | **optional.**|  | 
 **sendingBusinessAddress2** | **optional.**|  | 
 **sendingBusinessCity** | **optional.**|  | 
 **sendingBusinessState** | **optional.**|  | 
 **sendingBusinessZipcode** | **optional.**|  | 
 **sendingBusinessCountry** | **optional.**|  | 
 **sendingBusinessRepresentativeFirstName** | **optional.**|  | 
 **sendingBusinessRepresentativeMiddleName** | **optional.**|  | 
 **sendingBusinessRepresentativeLastName** | **optional.**|  | 
 **sendingBusinessRepresentativeLastName2** | **optional.**|  | 
 **sendingBusinessRepresentativeNativeName** | **optional.**|  | 
 **sendingBusinessIdType** | **optional.**|  | 
 **sendingBusinessIdNumber** | **optional.**|  | 
 **sendingBusinessCountryOfIssue** | **optional.**|  | 
 **sendingBusinessIssueDate** | **optional.**|  | 
 **sendingBusinessExpireDate** | **optional.**|  | 
 **sendingBusinessDateOfIncorporation** | **optional.**|  | 
 **documentReferenceNumber** | **optional.**|  | 
 **swiftCode** | **optional.**|  | 
 **bankIfcCode** | **optional.**|  | 
 **bankBranchName** | **optional.**|  | 
 **bankBranchCode** | **optional.**|  | 
 **bankAccountNumber** | **optional.**|  | 
 **cbcBankName** | **optional.**|  | 
 **bpiBankName** | **optional.**|  | 
 **partnerName** | **optional.**|  | 
 **sourceOfIncome** | **optional.**|  | 
 **relationship** | **optional.**|  | 
 **reason** | **optional.**|  | 
 **senderReference** | **optional.**|  | 
 **additionalInformation** | **optional.**|  | 

### Return type

[**InlineResponse20052**](inline_response_200_52.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: application/x-www-form-urlencoded
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

