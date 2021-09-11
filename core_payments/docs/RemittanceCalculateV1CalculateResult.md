# RemittanceCalculateV1CalculateResult

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ResultCode** | **string** |  | [optional] [default to null]
**InitialAmount** | **string** |  | [optional] [default to null]
**AmountToSend** | **string** |  | [optional] [default to null]
**AmountToSendCurrency** | **string** |  | [optional] [default to null]
**AmountToReceive** | **string** |  | [optional] [default to null]
**AmountToReceiveCurrency** | **string** |  | [optional] [default to null]
**ReceiveCurrency** | **string** |  | [optional] [default to null]
**ProviderExchangeRate** | **string** |  | [optional] [default to null]
**ExchangeRate** | **string** |  | [optional] [default to null]
**FixedFee** | **string** |  | [optional] [default to null]
**Commission** | **string** |  | [optional] [default to null]
**TotalAmount** | **string** |  | [optional] [default to null]
**AmountToBeDeducted** | **string** |  | [optional] [default to null]
**AmountDetails** | [***RemittanceCalculateV1CalculateResultAmountDetails**](Remittance_calculate.v1_calculate_result_amount_details.md) |  | [optional] [default to null]
**Quotation** | [***RemittanceQuotationV1**](Remittance_quotation.v1.md) |  | [optional] [default to null]
**SubscriptionPricing** | **string** |  | [optional] [default to null]
**RequiredFields** | **string** | For new calculate endpoint, this will not be present on response. Please check on beneficiary_form or sender_form if a particular field is required or not | [optional] [default to null]
**AttachmentForm** | [**[]RemittanceBasicFormV1**](Remittance_basic_form.v1.md) |  | [optional] [default to null]
**SenderForm** | [**[]RemittanceBasicFormV1**](Remittance_basic_form.v1.md) |  | [optional] [default to null]
**BeneficiaryForm** | [**[]RemittanceBasicFormV1**](Remittance_basic_form.v1.md) |  | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

