# CardsCardIdBody1

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | Card unique identifier. When present, this will unblock the card associated with this parameter. | [default to null]
**AssocNumber** | **string** | Physical Card Proxy Number [format: PY + Proxy Number (eg. &#x27;PY000000000123&#x27;)] or Account identification number. When present, this will bind the card to the current user.  | [optional] [default to null]
**RefId** | **string** | Transaction reference identifier obtained after using POST /users/wallets/funds when the consumer is pre-funded or after using POST /oauth/consumer/funds when the consumer is not pre-funded. | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

