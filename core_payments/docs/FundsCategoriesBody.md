# FundsCategoriesBody

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | Unique category label which accepts case insensitive alphanumeric combinations with hyphens | [default to null]
**DefaultAmount** | **string** | Serves as the maximum amount allowable for the category | [default to null]
**Factor** | **string** | Defines the default_amount type or context (UNIT, PERCENT). Defaults to UNIT | [optional] [default to null]
**Mcc** | **string** | If a comma-delimited list of Merchant Category Codes is passed, this category will use them as filters. Value with no record matches are ignored. | [optional] [default to null]
**Country** | **string** | A comma-delimited list of country abbreviations in [ISO 3166 alpha-3](http://www.iso.org/iso/country_codes) format. If provided, the currency of these countries will be used as filters. Value with no record matches are ignored. | [optional] [default to null]
**Priority** | **string** | Spending priority ( The lower the number the hight the priority ) | [optional] [default to null]
**Card** | **string** | If a comma-delimited list of card ID is passed, this category will use them as filters. Value with no record matches are ignored. | [optional] [default to null]
**Merchant** | **string** | If a comma-delimited list of Merchant Codes is passed, this category will use them as filters | [optional] [default to null]
**Terminal** | **string** | If a comma-delimited list of Terminal Codes is passed, this category will use them as filters | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

