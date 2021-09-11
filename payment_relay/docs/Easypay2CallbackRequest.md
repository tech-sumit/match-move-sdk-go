# Easypay2CallbackRequest

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**TMMCode** | **string** | Merchant ID assigned by Easypay2 | [default to null]
**TMRefNo** | **string** | Payment Reference ID -- The value in this field is an echo of the &#x60;ref&#x60; field submitted from IH-Relay to Easypay2 | [default to null]
**TMDebitAmt** | **float32** | Amount -- The value in this field is an echo of the &#x60;amt&#x60; field submitted from IH-Relay to Easypay2 (In json format, this must not be enclosed in quotation marks serving as string) | [default to null]
**TMCurrency** | **string** | Currency -- The value in this field is an echo of the &#x60;cur&#x60; field submitted from IH-Relay to Easypay2 | [default to null]
**TMStatus** | **string** | Status of the Transaction : Valid values         &#x60;NO&#x60; – Unsuccessful (Bank rejected transaction), &#x60;YES&#x60; – Successful (Bank has authorized the authorization / payment) | [default to null]
**TMApprovalCode** | **string** | Bank’s approval code for successful transactions NULL or empty if the transaction failed. | [default to null]
**TMBankRespCode** | **string** | Response code returned by the acquiring bank         &#x60;00&#x60; for successful transactions, &#x60;NN&#x60; denotes failure reason code | [default to null]
**TMError** | **string** | Error details | [optional] [default to null]
**TMUserField1** | **string** | Additional user field 1 | [optional] [default to null]
**TMUserField2** | **string** | Additional user field 2 | [optional] [default to null]
**TMUserField3** | **string** | Additional user field 3 | [optional] [default to null]
**TMUserField4** | **string** | Additional user field 4 | [optional] [default to null]
**TMUserField5** | **string** | Additional user field 5 : Callback Signature | [default to null]
**TMCCLast4Digit** | **string** | Last 4 digits of credit card number | [default to null]
**TMExpiryDate** | **string** | Credit card expiry date specified in request | [default to null]
**TMTrnType** | **string** | Type of transaction performed | [default to null]
**TMSubTrnType** | **string** | Sub-type of transaction performed | [optional] [default to null]
**TMCardNum** | **string** | This value is returned if rcard is specified from IH-Relay to Easypay2 | [default to null]
**TMOriginalPayType** | **string** | Transaction pay type | [optional] [default to null]
**AdditionalHook** | **string** | JSON encoded text string containing additional parameters for confirmation hook | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

