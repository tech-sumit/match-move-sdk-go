# RemittanceQuotationV1

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | Quotation ID from Remit Service | [default to null]
**ExternalId** | **string** | External ID from Rails | [default to null]
**QuotationId** | **string** | Quotation ID from Rails | [default to null]
**PayoutAgent** | **string** | Payout Agent Identifier | [default to null]
**CalculationMode** | **string** | Calculation Mode | [default to null]
**TransactionType** | **string** | Transaction Type - C2C or B2B | [default to null]
**Initial** | [***AmountDetailsV1**](Amount_details.v1.md) |  | [default to null]
**Send** | [***AmountDetailsV1**](Amount_details.v1.md) |  | [default to null]
**Receive** | [***AmountDetailsV1**](Amount_details.v1.md) |  | [default to null]
**Fee** | [***AmountDetailsV1**](Amount_details.v1.md) |  | [default to null]
**ExchangeRate** | **string** | Calculated Exchange Rate after fx markup | [default to null]
**ProviderExchangeRate** | **string** | Exchange Rate from Rails | [default to null]
**Status** | **string** | Quotation Status | [default to null]
**DateExpiry** | **string** | Quotation Expiry (after an hour it was created for T2) | [default to null]
**DateCreated** | **string** | Quotation Date Created | [default to null]
**SubscriptionPricing** | **bool** | Set to true if subscription pricing was applied, else false | [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

