# DirectProviderHashBody

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ConsumerHash** | **string** | Consumer Hash ID | [default to 9e050afeed738f820b36a935c8ff1ed5]
**ProductCode** | **string** | OP Product Code | [default to sgmcpr]
**Amount** | **float32** | Topup amount | [default to null]
**KycStatus** | **string** | User KYC Status [&#x60;pre_kyc&#x60;, &#x60;post_kyc&#x60;] | [default to pre_kyc]
**UserHashId** | **string** | User Unique Identifier | [default to c4432940228d7fc04bb1656e10103c60]
**Email** | **string** | User Email Address | [optional] [default to louie+relay1508151862@matchmove.com]
**MobileCountryCode** | **int32** | User Mobile Country Code | [optional] [default to null]
**Mobile** | **int32** | User Mobile Number | [optional] [default to null]
**Currency** | **string** | Currency code ISO 4217 alpha code | [optional] [default to SGD]
**CardToken** | **string** | Token assigned to a particular card of the current user | [optional] [default to null]
**Pan** | **string** | User&#x27;s Card Primary Account Number | [optional] [default to null]
**ExpiryMonth** | **string** | Expiry Month of User&#x27;s Card Primary Account Number | [optional] [default to null]
**ExpiryYear** | **string** | Expiry Year of User&#x27;s Card Primary Account Number | [optional] [default to null]
**SecurityCode** | **string** | Security Code of User&#x27;s Card Primary Account Number | [optional] [default to null]
**Source** | **string** | Source on where we received the transaction [&#x60;CALL_CENTRE&#x60;, &#x60;CARD_PRESENT&#x60;, &#x60;INTERNET&#x60;, &#x60;MAIL_ORDER&#x60;, &#x60;MOTO&#x60;, &#x60;TELEPHONE_ORDER&#x60;, &#x60;VOICE_RESPONSE&#x60;] | [optional] [default to INTERNET]
**BillingAddress** | **string** | JSON-encoded value of User&#x27;s billing address which should contain [street, street2, city, stateProvince, country(3 letter ISO standard alpha country code), postcodeZip] | [optional] 
**Customer** | **string** | JSON-encoded value of User&#x27;s Details which should contain [firstName, lastName] | [optional] 
**ClientRefId** | **string** | Consumer Identifier for a Payment Transaction | [optional] [default to null]
**ClientIp** | **string** | Client IP Address | [optional] 
**ActionOrigin** | **string** | Origin on where the payment transaction is executed [&#x60;admin&#x60;, &#x60;client&#x60;, &#x60;cron&#x60;, &#x60;payment_provider&#x60;] | [optional] [default to client]
**SaveCard** | **string** | Will save card | [optional] [default to SAVE_CARD.0_]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

