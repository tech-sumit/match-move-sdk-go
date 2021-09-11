# TransactionCategoriesRulesRequestV1

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**EventCodes** | **string** | Event code from event_codes table | [default to null]
**Channel** | **string** | Transcation channel like atm,pos,ecom | [optional] [default to null]
**AcquiringCountryCode** | **string** | 3 Digits iso defined country code number like 356 | [optional] [default to null]
**Currency** | **string** | currency code like USD | [optional] [default to null]
**Network** | **string** | Transcation network like visa,mastercard | [optional] [default to null]
**TerminalId** | **string** | terminal_id defined in transcation rules table | [optional] [default to null]
**MerchantId** | **string** | merchant_id defined in transcation rules table | [optional] [default to null]
**PosEntryModeValue** | **string** | Point of service entry mode like MANUALLY_KEYED | [optional] [default to null]
**Pp** | **string** | pp value that goes into pp cosnumer table | [optional] [default to null]
**MerchantCategoryCode** | **string** | merchant_category_code defined in transcation rules table | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

