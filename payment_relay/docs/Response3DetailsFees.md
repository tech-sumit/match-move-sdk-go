# Response3DetailsFees

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Currency** | **string** | Currency code [ISO 3166](https://www.iso.org/iso-3166-country-codes.html) alpha-3 | [optional] [default to null]
**Amount** | [**float64**](decimal.md) | Amount entered. Net amount and credit amount will differ depending on the current cppc fee settings. | [optional] [default to null]
**TotalFees** | [**float64**](decimal.md) | Total fees (provider, tax, charge, service) | [optional] [default to null]
**Discount** | [**float64**](decimal.md) | Discount offered | [optional] [default to null]
**NetAmount** | [**float64**](decimal.md) | Total amount to pay | [optional] [default to null]
**CreditAmount** | [**float64**](decimal.md) | Actual amount to be credited to MM Wallet | [optional] [default to null]
**CalculationMode** | **string** | Mode on how the calculation of credit amount will be. &#x60;credit&#x60; means credit amount will be the exact amount entered; &#x60;net&#x60; means the amount entered less fees | [optional] [default to null]
**Computation** | **string** | Flag on how the fee will be deducted, whether by &#x60;flat&#x60; or &#x60;percentage&#x60; | [optional] [default to null]
**FeeFlat** | [**float64**](decimal.md) | Fee in Flat | [optional] [default to null]
**FeePercentage** | [**float64**](decimal.md) | Fee in Percentage of the credit amount | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

