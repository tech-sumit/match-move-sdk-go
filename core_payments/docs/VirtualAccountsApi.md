# {{classname}}

All URIs are relative to *https://beta-api.mmvpay.com/{product}/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateBankAccountCreditTransaction**](VirtualAccountsApi.md#CreateBankAccountCreditTransaction) | **Post** /users/wallets/fund_transfers/credit | Credit bank account
[**CreateBankAccountDebitTransaction**](VirtualAccountsApi.md#CreateBankAccountDebitTransaction) | **Post** /users/wallets/fund_transfers/debit | Debit bank account
[**CreateBankAccountStandingInstruction**](VirtualAccountsApi.md#CreateBankAccountStandingInstruction) | **Post** /users/wallets/standing_instructions | Create New Standing Instruction
[**CreateBeneficiaryBankAccount**](VirtualAccountsApi.md#CreateBeneficiaryBankAccount) | **Post** /users/wallets/bank_accounts | Create New Beneficiary Bank Account
[**CreateVirtualAccounts**](VirtualAccountsApi.md#CreateVirtualAccounts) | **Post** /users/wallets/virtual_accounts | Create virtual account
[**DeleteBankAccountStandingInstruction**](VirtualAccountsApi.md#DeleteBankAccountStandingInstruction) | **Delete** /users/wallets/standing_instructions | Delete existing instruction
[**DeleteBeneficiaryBankAccount**](VirtualAccountsApi.md#DeleteBeneficiaryBankAccount) | **Delete** /users/wallets/bank_accounts | Remove beneficiary bank account
[**GetLinkedBeneficiaryBankAccount**](VirtualAccountsApi.md#GetLinkedBeneficiaryBankAccount) | **Get** /users/wallets/bank_accounts | Get Beneficiary Bank Accounts
[**GetUsersWalletsAccountsTransactions**](VirtualAccountsApi.md#GetUsersWalletsAccountsTransactions) | **Get** users/wallets/account/transactions | Get Account Transactions
[**GetVirtualAccountStandingInstruction**](VirtualAccountsApi.md#GetVirtualAccountStandingInstruction) | **Get** /users/wallets/standing_instructions | Get Standing Instruction
[**GetVirtualAccounts**](VirtualAccountsApi.md#GetVirtualAccounts) | **Get** /users/wallets/virtual_accounts | Get linked virtual accounts
[**UpdateBankAccountStandingInstruction**](VirtualAccountsApi.md#UpdateBankAccountStandingInstruction) | **Put** /users/wallets/standing_instructions | Update standing instruction status

# **CreateBankAccountCreditTransaction**
> FundTransferV1 CreateBankAccountCreditTransaction(ctx, bankAccountId, clientRefId, amount, currency, purposeOfTransfer, xAuthUserID)
Credit bank account

Transfer money to provided beneficiary account

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **bankAccountId** | [**string**](.md)|  | 
  **clientRefId** | **string**|  | 
  **amount** | **float64**|  | 
  **currency** | **string**|  | 
  **purposeOfTransfer** | **string**|  | 
  **xAuthUserID** | **string**| MatchMove provided hash ID for the user | 

### Return type

[**FundTransferV1**](Fund_transfer.v1.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: application/x-www-form-urlencoded
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CreateBankAccountDebitTransaction**
> FundTransferV1 CreateBankAccountDebitTransaction(ctx, bankAccountId, clientRefId, amount, currency, purposeOfTransfer, xAuthUserID)
Debit bank account

Pull money from specified bank account according to the specified standing instruction.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **bankAccountId** | [**string**](.md)|  | 
  **clientRefId** | **string**|  | 
  **amount** | **float64**|  | 
  **currency** | **string**|  | 
  **purposeOfTransfer** | **string**|  | 
  **xAuthUserID** | **string**| MatchMove provided hash ID for the user | 

### Return type

[**FundTransferV1**](Fund_transfer.v1.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: application/x-www-form-urlencoded
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CreateBankAccountStandingInstruction**
> StandingInstructionV1 CreateBankAccountStandingInstruction(ctx, bankAccountId, number, maximumDebitableAmount, dateApplication, dateExpiry, status, consentFullName, consentNumbers, consentSign, xAuthUserID)
Create New Standing Instruction

Create new standing instruction for provided bank account

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **bankAccountId** | [**string**](.md)|  | 
  **number** | **float64**|  | 
  **maximumDebitableAmount** | **string**|  | 
  **dateApplication** | **string**|  | 
  **dateExpiry** | **string**|  | 
  **status** | **string**|  | 
  **consentFullName** | **string**|  | 
  **consentNumbers** | **string**|  | 
  **consentSign** | **string**|  | 
  **xAuthUserID** | **string**| MatchMove provided hash ID for the user | 

### Return type

[**StandingInstructionV1**](Standing_instruction.v1.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: application/x-www-form-urlencoded
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CreateBeneficiaryBankAccount**
> BankAccountV1 CreateBeneficiaryBankAccount(ctx, bankAccountNumber, bankHolderName, bankCode, xAuthUserID)
Create New Beneficiary Bank Account

Create new Beneficiary Bank Account for the provided user

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **bankAccountNumber** | **string**|  | 
  **bankHolderName** | **string**|  | 
  **bankCode** | **string**|  | 
  **xAuthUserID** | **string**| MatchMove provided hash ID for the user | 

### Return type

[**BankAccountV1**](Bank_account.v1.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: application/x-www-form-urlencoded
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CreateVirtualAccounts**
> VirtualAccountV1 CreateVirtualAccounts(ctx, xAuthUserID)
Create virtual account

Create a virtual account for the user already created in the system

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **xAuthUserID** | [**string**](.md)| MatchMove provided hash ID for the user | 

### Return type

[**VirtualAccountV1**](Virtual_account.v1.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteBankAccountStandingInstruction**
> InlineResponse20046 DeleteBankAccountStandingInstruction(ctx, bankAccountId, xAuthUserID)
Delete existing instruction

Delete existing instruction on provided bank account

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **bankAccountId** | [**string**](.md)|  | 
  **xAuthUserID** | **string**| MatchMove provided hash ID for the user | 

### Return type

[**InlineResponse20046**](inline_response_200_46.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: application/x-www-form-urlencoded
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteBeneficiaryBankAccount**
> InlineResponse20046 DeleteBeneficiaryBankAccount(ctx, id, xAuthUserID)
Remove beneficiary bank account

Remove linked beneficiary bank account for the user. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | [**string**](.md)|  | 
  **xAuthUserID** | **string**| MatchMove provided hash ID for the user | 

### Return type

[**InlineResponse20046**](inline_response_200_46.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: application/x-www-form-urlencoded
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetLinkedBeneficiaryBankAccount**
> InlineResponse20045 GetLinkedBeneficiaryBankAccount(ctx, xAuthUserID)
Get Beneficiary Bank Accounts

Get Beneficiary Bank Accounts 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **xAuthUserID** | **string**| MatchMove provided hash ID for the user | 

### Return type

[**InlineResponse20045**](inline_response_200_45.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetUsersWalletsAccountsTransactions**
> InlineResponse20044 GetUsersWalletsAccountsTransactions(ctx, xAuthUserID)
Get Account Transactions

Retrieve all transactions done by the user via Bank Transfer (Top up and Trasfer out) 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **xAuthUserID** | **string**| MatchMove provided hash ID for the user | 

### Return type

[**InlineResponse20044**](inline_response_200_44.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetVirtualAccountStandingInstruction**
> StandingInstructionV1 GetVirtualAccountStandingInstruction(ctx, xAuthUserID)
Get Standing Instruction

Get standing instructions configured against a wallet 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **xAuthUserID** | **string**| MatchMove provided hash ID for the user | 

### Return type

[**StandingInstructionV1**](Standing_instruction.v1.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetVirtualAccounts**
> InlineResponse20042 GetVirtualAccounts(ctx, xAuthUserID)
Get linked virtual accounts

Retrieve user's virtual accounts.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **xAuthUserID** | **string**| MatchMove provided hash ID for the user | 

### Return type

[**InlineResponse20042**](inline_response_200_42.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateBankAccountStandingInstruction**
> StandingInstructionV1 UpdateBankAccountStandingInstruction(ctx, bankAccountId, status, xAuthUserID)
Update standing instruction status

Update the standing instruction status of provided bank account.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **bankAccountId** | [**string**](.md)|  | 
  **status** | **string**|  | 
  **xAuthUserID** | **string**| MatchMove provided hash ID for the user | 

### Return type

[**StandingInstructionV1**](Standing_instruction.v1.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: application/x-www-form-urlencoded
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

