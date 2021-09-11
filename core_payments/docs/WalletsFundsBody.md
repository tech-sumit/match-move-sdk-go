# WalletsFundsBody

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Email** | **string** | The registered &#x60;email&#x60; address of the user for which the fund transaction is being initiated | [optional] [default to null]
**MobileCountryCode** | **string** | The &#x60;mobile_country_code&#x60; of the user for which the fund transaction is being initiated | [optional] [default to null]
**Mobile** | **string** | The &#x60;mobile&#x60; of the user for which the fund transaction is being initiated | [optional] [default to null]
**UserId** | **string** | The &#x60;user_id&#x60; of the user for which the fund transaction is being initiated | [optional] [default to null]
**Amount** | **float64** | The amount for the transaction. The format supported would be based on the precision of the base currency | [default to null]
**Details** | **string** | JSON encoded details to be passed as meta data to be stored in MatchMove sytem | [optional] [default to null]
**ForwardedFor** | **string** |  | [optional] [default to null]
**HashedPan** | **string** |  | [optional] [default to null]
**FundCategoryName** | **string** | Optional category to specify where the credit/transfer of funds happens. | [optional] [default to null]
**Type_** | **string** |  | [optional] [default to null]
**LoadCard** | **bool** | To be specified if the funds need to be directly moved to the card in non-shared card balance. | [optional] [default to null]
**CardId** | **string** | The card_id of the card in non-shared balance to which the funds are to be moved. Required if the value for &#x60;load_card&#x60; is true. | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

