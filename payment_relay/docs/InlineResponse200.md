# InlineResponse200

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PaymentRefId** | **string** | Unique Payment Reference Id | [optional] [default to null]
**Consumer** | [**[]InlineResponse200Consumer**](inline_response_200_consumer.md) | Consumer Details | [default to null]
**Product** | [**[]InlineResponse200Product**](inline_response_200_product.md) | Product Details | [default to null]
**Provider** | [**[]InlineResponse200Provider**](inline_response_200_provider.md) | Provider Details | [default to null]
**User** | [**[]InlineResponse200User**](inline_response_200_user.md) | User Details like hash id, email and any other user related info | [default to null]
**Fees** | [**[]InlineResponse200Fees**](inline_response_200_fees.md) | Fees breakdown | [default to null]
**ExpiresAt** | **string** | Expiry date of the transaction | [default to null]
**Message** | **string** | Instuction or message from Relay to customer | [default to null]
**ExpiryDuration** | **int32** | Expiry duration in seconds | [optional] [default to null]
**QrDetails** | [**[]InlineResponse200QrDetails**](inline_response_200_qr_details.md) |  | [optional] [default to null]
**QrImage** | [**[]InlineResponse200QrImage**](inline_response_200_qr_image.md) |  | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

