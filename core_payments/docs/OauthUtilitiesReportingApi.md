# {{classname}}

All URIs are relative to *https://beta-api.mmvpay.com/{product}/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetConsumerUserInstruments**](OauthUtilitiesReportingApi.md#GetConsumerUserInstruments) | **Get** /oauth/consumer/users/instruments | Retrieve all payments instruments of the users in the system
[**GetOauthConsumerConsumerKeyTransactionsPage**](OauthUtilitiesReportingApi.md#GetOauthConsumerConsumerKeyTransactionsPage) | **Get** /oauth/consumer/{consumer_key}/transactions/{page} | Get Consumer transaction
[**GetOauthConsumerFunds**](OauthUtilitiesReportingApi.md#GetOauthConsumerFunds) | **Get** /oauth/consumer/funds | Retrieve Oauth Consumer Funds
[**GetOauthConsumerSearchUsers**](OauthUtilitiesReportingApi.md#GetOauthConsumerSearchUsers) | **Get** /oauth/consumer/search/users | Gets detailed information of a user
[**GetOauthConsumerWalletFunds**](OauthUtilitiesReportingApi.md#GetOauthConsumerWalletFunds) | **Get** /oauth/consumer/wallets/funds | Closed Loop Transactions
[**GetOauthConsumerWalletReverseLimitOffset**](OauthUtilitiesReportingApi.md#GetOauthConsumerWalletReverseLimitOffset) | **Get** /oauth/consumer/wallet/reverse/{limit}/{offset} | Get oauth wallet reversals
[**GetOauthUser**](OauthUtilitiesReportingApi.md#GetOauthUser) | **Get** /oauth/user | Get oauth user
[**GetOauthUsersSearch**](OauthUtilitiesReportingApi.md#GetOauthUsersSearch) | **Get** /oauth/users/search | Get Oauth Users Search
[**GetProgramBankAccountDebitCreditTransactions**](OauthUtilitiesReportingApi.md#GetProgramBankAccountDebitCreditTransactions) | **Get** /oauth/consumer/wallets/fund_transfers | Get all fund transfer transactions
[**GetProgramBankAccountStandingInstruction**](OauthUtilitiesReportingApi.md#GetProgramBankAccountStandingInstruction) | **Get** /oauth/consumer/wallets/standing_instructions | Get all Standing Instructions
[**GetProgramBankTransferTransactions**](OauthUtilitiesReportingApi.md#GetProgramBankTransferTransactions) | **Get** /oauth/consumer/account/transactions | Get Program Bank Transfer Transactions
[**GetProgramBeneficiaryBankAccount**](OauthUtilitiesReportingApi.md#GetProgramBeneficiaryBankAccount) | **Get** /oauth/consumer/wallets/bank_accounts | Get all Bank Accounts
[**GetProgramVirtualAccounts**](OauthUtilitiesReportingApi.md#GetProgramVirtualAccounts) | **Get** /oauth/consumer/wallets/virtual_accounts | Get all Virtual Accounts

# **GetConsumerUserInstruments**
> InlineResponse20014 GetConsumerUserInstruments(ctx, optional)
Retrieve all payments instruments of the users in the system

A lightweight endpoint to return the User list of payment instruments. Suggested to use in case of repeat calls.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***OauthUtilitiesReportingApiGetConsumerUserInstrumentsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a OauthUtilitiesReportingApiGetConsumerUserInstrumentsOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **userId** | **optional.String**| User hash ID of the user returned from the MatchMove system | 
 **mobileCountryCode** | **optional.String**| If mobile is passed this is a required parameter | 
 **mobile** | **optional.String**| if mobile_country_code is passed this is a required parameter | 
 **cardTypeCode** | **optional.String**| The list of supported values is retrievable by calling GET /users/wallets/cards/types | 
 **cardId** | **optional.String**| Card hash ID of the user returned from card creation | 

### Return type

[**InlineResponse20014**](inline_response_200_14.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetOauthConsumerConsumerKeyTransactionsPage**
> []OauthTransactionV1 GetOauthConsumerConsumerKeyTransactionsPage(ctx, consumerKey, page)
Get Consumer transaction

Get Consumer Transactions 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **consumerKey** | **string**|  | 
  **page** | **string**|  | 

### Return type

[**[]OauthTransactionV1**](Oauth_transaction.v1.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetOauthConsumerFunds**
> InlineResponse20037 GetOauthConsumerFunds(ctx, ids, optional)
Retrieve Oauth Consumer Funds

Retrieve users funds transactions 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **ids** | **string**| Reference identifiers for the top up transaction(s). Multiple identifiers in CSV format are accepted. | 
 **optional** | ***OauthUtilitiesReportingApiGetOauthConsumerFundsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a OauthUtilitiesReportingApiGetOauthConsumerFundsOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **paymentRef** | **optional.String**| Payment transaction identifier | 

### Return type

[**InlineResponse20037**](inline_response_200_37.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetOauthConsumerSearchUsers**
> UserReportV1 GetOauthConsumerSearchUsers(ctx, optional)
Gets detailed information of a user

Gets detailed information of a user 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***OauthUtilitiesReportingApiGetOauthConsumerSearchUsersOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a OauthUtilitiesReportingApiGetOauthConsumerSearchUsersOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **email** | **optional.String**| Email Address (from users table) | 
 **mobile** | **optional.String**| Full mobile (concatenated value of country code and mobile number) | 
 **userHashId** | **optional.String**| User identifier | 
 **userStatus** | **optional.String**| User State | 
 **kycStatus** | **optional.String**| User KYC State | 
 **firstName** | **optional.String**| first name to search | 
 **lastName** | **optional.String**| last name to search | 
 **customerId** | **optional.String**| Customer ID to search | 
 **idType** | **optional.String**| ID Type to search | 
 **idNumber** | **optional.String**| ID number to search | 
 **idDateIssue** | **optional.String**| ID issue date to search | 
 **idDateExpiry** | **optional.String**| ID expiry date to search | 
 **birthday** | **optional.String**| Birthday to search | 
 **cardIdentifier** | **optional.String**| Card Identifier:   * &#x60;proxy&#x60; - Card Proxy number to search   * &#x60;mask&#x60; - Masked card number to search   * &#x60;full&#x60; - Full(actual) card number to search  | 
 **cardIdentifierValue** | **optional.String**| Value is dependent on the &#x60;card_identifier&#x60; parameter | 
 **balanceCondition** | **optional.String**| Balances Condition:   * &#x60;gt&#x60; - Amount greater than or equal to specified balance   * &#x60;lt&#x60; - Amount lesser than or equal to specified balance   * &#x60;eq&#x60; - Amount equal to specified balance  | [default to eq]
 **cardAvailableBalance** | **optional.Float32**| Get all users where its per card available balance is within the specified &#x60;balance_condition&#x60; parameter | 
 **cardWithholdingBalance** | **optional.Float32**| Get all users where its per card withholding balance is within the specified &#x60;balance_condition&#x60; parameter | 
 **walletAvailableBalance** | **optional.Float32**| Get all users where its per wallet available balance is within the specified &#x60;balance_condition&#x60; parameter | 
 **walletWithholdingBalance** | **optional.Float32**| Get all users where its per wallet withholding balance is within the specified &#x60;balance_condition&#x60; parameter | 
 **transactionReference** | **optional.String**| Get all users given the transaction reference (will apply to all transaction types) | 
 **matchType** | **optional.String**| Matching Condition:   * &#x60;fuzzy&#x60; - endpoint will perform a &#x60;LIKE&#x60; condition in query   * &#x60;exact&#x60; - endpoint will perform an &#x60;EQUAL&#x60; condition in query  | 
 **responseType** | **optional.String**| Matching Condition:   * &#x60;expand&#x60; - response will show an expandable/complete information for address, documents, wallets object   * &#x60;condense&#x60; - response will show a url link on information for address, documents, wallets object  | 
 **filterCondition** | **optional.String**| Matching Condition:   * &#x60;AND&#x60; - for 2 or more filter parameters, the query will perform an &#x60;AND&#x60; operation in query    * &#x60;OR&#x60; - for 2 or more filter parameters, the query will perform an &#x60;OR&#x60; operation in query  | [default to AND]
 **page** | **optional.Float64**| Page Number to display | [default to 1]
 **recordsPerPage** | **optional.Float64**| Records per page to display | [default to 10]
 **partnerId** | **optional.String**| Partner Id | 

### Return type

[**UserReportV1**](User_report.v1.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetOauthConsumerWalletFunds**
> ClosedLoopV1 GetOauthConsumerWalletFunds(ctx, optional)
Closed Loop Transactions

Gets detailed information of a closed-loop transactions

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***OauthUtilitiesReportingApiGetOauthConsumerWalletFundsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a OauthUtilitiesReportingApiGetOauthConsumerWalletFundsOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **type_** | **optional.String**| Acceptable Transaction Type:   * &#x60;Topup&#x60; - all top up transactions   * &#x60;Transfer&#x60; - all wallet to wallet money transfer transactions   * &#x60;Load&#x60; - all wallet to card transfer transactions   * &#x60;Unload&#x60; - all card to wallet transfer transactions   * &#x60;Reversal&#x60; - all  transactions   * &#x60;Refund&#x60; - all card to wallet transfer transactions   * &#x60;All&#x60; - all transactions above  | [default to All]
 **userHashId** | **optional.String**|  | 
 **email** | **optional.String**|  | 
 **mobile** | **optional.String**|  | 
 **cardIdentifier** | **optional.String**| Card Identifier:   * &#x60;proxy&#x60; - Card/Wallet Proxy number to search   * &#x60;mask&#x60; - Masked card number to search   * &#x60;full&#x60; - Full(actual) card number to search  | 
 **cardIdentifierValue** | **optional.String**| Value is dependent on the &#x60;card_identifier&#x60; parameter; get all transactions done by the given card_identifier_value | 
 **transactionRefHash** | **optional.String**| Transaction Reference; get all transactions given the reference hash | 
 **transactionStatus** | **optional.String**| Transaction Status; get all transactions given the status:   * &#x60;hold&#x60; - all withholding transaction   * &#x60;active&#x60; - all success transaction   * &#x60;failed&#x60; - all failed transaction  | 
 **addedDate** | **optional.String**| Transaction Date Added (YYYY-MM-DD); get all transactions added on the given date | 
 **updatedDate** | **optional.String**| Transaction Date Updated (YYYY-MM-DD); get all transactions updaed on the given date | 
 **reversalDate** | **optional.String**| Transaction Date Reversal (YYYY-MM-DD); get all transactions reversed on the given date | 
 **refundDate** | **optional.String**| Transaction Date Refund (YYYY-MM-DD); get all transactions refunded on the given date | 

### Return type

[**ClosedLoopV1**](Closed_loop.v1.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetOauthConsumerWalletReverseLimitOffset**
> InlineResponse20036 GetOauthConsumerWalletReverseLimitOffset(ctx, limit, offset)
Get oauth wallet reversals

Gets Wallet reversals 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **limit** | **string**|  | 
  **offset** | **string**|  | 

### Return type

[**InlineResponse20036**](inline_response_200_36.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetOauthUser**
> []UserFullV1 GetOauthUser(ctx, )
Get oauth user

Get oauth user 

### Required Parameters
This endpoint does not need any parameter.

### Return type

[**[]UserFullV1**](User_full.v1.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetOauthUsersSearch**
> InlineResponse20040 GetOauthUsersSearch(ctx, )
Get Oauth Users Search

Get Oauth Users Search 

### Required Parameters
This endpoint does not need any parameter.

### Return type

[**InlineResponse20040**](inline_response_200_40.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetProgramBankAccountDebitCreditTransactions**
> InlineResponse20049 GetProgramBankAccountDebitCreditTransactions(ctx, optional)
Get all fund transfer transactions

Get all fund transfer processed under the program

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***OauthUtilitiesReportingApiGetProgramBankAccountDebitCreditTransactionsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a OauthUtilitiesReportingApiGetProgramBankAccountDebitCreditTransactionsOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **userHashId** | **optional.String**| af11e6342686a34d37a58e7130ec5e11 | 
 **walletHashId** | **optional.String**| df11e6342686a34d37a58e7130ec5e11 | 
 **bankAccountNumber** | **optional.String**| 123456 | 
 **bankAccountId** | **optional.String**| df11e6342686a34d37a58e7130ec5e11 | 
 **uniquePaymentId** | **optional.String**| 123456 | 
 **status** | **optional.String**| active | inactive | terminated | 
 **accountType** | **optional.String**| consumer or corporate | 
 **fromDate** | **optional.String**| From Date | 
 **toDate** | **optional.String**| To Date | 
 **page** | **optional.Float64**| The page number | 
 **recordsPerPage** | **optional.Float64**| The number of records to return per page | 

### Return type

[**InlineResponse20049**](inline_response_200_49.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetProgramBankAccountStandingInstruction**
> InlineResponse20048 GetProgramBankAccountStandingInstruction(ctx, optional)
Get all Standing Instructions

Get all standing instructions registered under the program

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***OauthUtilitiesReportingApiGetProgramBankAccountStandingInstructionOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a OauthUtilitiesReportingApiGetProgramBankAccountStandingInstructionOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **userHashId** | **optional.String**| af11e6342686a34d37a58e7130ec5e11 | 
 **walletHashId** | **optional.String**| df11e6342686a34d37a58e7130ec5e11 | 
 **bankAccountNumber** | **optional.String**| 123456 | 
 **bankAccountId** | **optional.String**| df11e6342686a34d37a58e7130ec5e11 | 
 **status** | **optional.String**| pending | under-review | approved | expired | revoked | disabled | 
 **accountType** | **optional.String**| consumer or corporate | 
 **fromDate** | **optional.String**| From Date | 
 **toDate** | **optional.String**| To Date | 
 **page** | **optional.Float64**| Page Number | 
 **recordsPerPage** | **optional.Float64**| Number of records per page | 

### Return type

[**InlineResponse20048**](inline_response_200_48.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetProgramBankTransferTransactions**
> InlineResponse20044 GetProgramBankTransferTransactions(ctx, )
Get Program Bank Transfer Transactions

Retrieve list of all Bank Transfer transactions (Top  up and Transfer out) done in the system 

### Required Parameters
This endpoint does not need any parameter.

### Return type

[**InlineResponse20044**](inline_response_200_44.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetProgramBeneficiaryBankAccount**
> InlineResponse20047 GetProgramBeneficiaryBankAccount(ctx, optional)
Get all Bank Accounts

Get all Bank Accounts registered under the program

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***OauthUtilitiesReportingApiGetProgramBeneficiaryBankAccountOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a OauthUtilitiesReportingApiGetProgramBeneficiaryBankAccountOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **userHashId** | **optional.String**| af11e6342686a34d37a58e7130ec5e11 | 
 **walletHashId** | **optional.String**| df11e6342686a34d37a58e7130ec5e11 | 
 **bankAccountNumber** | **optional.String**| 123456 | 
 **bankAccountId** | **optional.String**| adf11e6342686a34d37a58e7130ec5e11 | 
 **accountType** | **optional.String**| consumer or corporate | 
 **fromDate** | **optional.String**| From Date | 
 **toDate** | **optional.String**| To Date | 
 **page** | **optional.Float64**| Page Number | 
 **recordsPerPage** | **optional.Float64**| Records per page | 

### Return type

[**InlineResponse20047**](inline_response_200_47.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetProgramVirtualAccounts**
> InlineResponse20043 GetProgramVirtualAccounts(ctx, optional)
Get all Virtual Accounts

Get all Virtual Accounts registered under the program

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***OauthUtilitiesReportingApiGetProgramVirtualAccountsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a OauthUtilitiesReportingApiGetProgramVirtualAccountsOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **userHashId** | **optional.String**| af11e6342686a34d37a58e7130ec5e11 | 
 **walletHashId** | **optional.String**| df11e6342686a34d37a58e7130ec5e11 | 
 **accountNumber** | **optional.String**| 123456 | 
 **status** | **optional.String**| active | inactive | terminated | 
 **accountType** | **optional.String**| consumer or corporate | 
 **fromDate** | **optional.String**| From Date | 
 **toDate** | **optional.String**| To Date | 
 **page** | **optional.Float64**| Page Number | 
 **recordsPerPage** | **optional.Float64**| Number of records per page | 

### Return type

[**InlineResponse20043**](inline_response_200_43.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

