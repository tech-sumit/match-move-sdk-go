# CardMiniV1

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | Unique Reference for Wallet in the MatchMove System | [default to null]
**Status** | [***CardFullV1Status**](Card_full.v1_status.md) |  | [default to null]
**Holder** | [***CardFullV1Holder**](Card_full.v1_holder.md) |  | [default to null]
**Funds** | [***CardFullV1Funds**](Card_full.v1_funds.md) |  | [optional] [default to null]
**Number** | **string** | Account Identification Number for the Wallet | [default to null]
**Network** | **string** | Network the card is associated to Example : visa, mastercard, rupay | [optional] [default to null]
**CardTypeCode** | **string** | One of the card types available via GET /users/wallets/cards/types | [optional] [default to null]
**Image** | [***CardImageV1**](Card_image.v1.md) |  | [optional] [default to null]
**Date** | [***CardFullV1Date**](Card_full.v1_date.md) |  | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

