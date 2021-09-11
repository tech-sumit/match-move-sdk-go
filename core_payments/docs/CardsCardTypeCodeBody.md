# CardsCardTypeCodeBody

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AssocNumber** | **string** | Required value supported only for non personalized physical cards. Physical Card Proxy Number [format: PY + Proxy Number (eg. &#x27;PY000000000123&#x27;)] or Account identification number. When present, this will bind the card to the current user.  | [optional] [default to null]
**RefId** | **string** | Transaction reference identifier obtained after using &#x60;POST /users/wallets/funds&#x60; when the consumer is pre-funded or after using &#x60;POST /oauth/consumer/funds&#x60; when the consumer is not pre-funded.  | [optional] [default to null]
**CardDesign** | **string** |  | [optional] [default to null]
**NameOnCard** | **string** |  | [optional] [default to null]
**AdditionalDetails** | **string** |  | [optional] [default to null]
**AutoActivate** | **bool** | If passed as true the POST request does the activation of the card along with the linkage of the card. Do note this parameter is only supported for non personalized card. Also depending on the card management partner used only option &#x60;true&#x60; might be supported  | [optional] [default to null]
**Var2faMethod** | **string** |  | [optional] [default to null]
**Var2faDelivery** | **string** |  | [optional] [default to null]
**Var2faValue** | **string** |  | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

