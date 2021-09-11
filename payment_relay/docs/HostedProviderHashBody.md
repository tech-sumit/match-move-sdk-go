# HostedProviderHashBody

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ConsumerHash** | **string** | Consumer Hash ID | [default to 9e050afeed738f820b36a935c8ff1ed5]
**ProductCode** | **string** | OP Product Code | [default to sgmcpr]
**Amount** | **float32** | Topup amount | [default to null]
**KycStatus** | **string** | User KYC Status [&#x60;pre_kyc&#x60;, &#x60;post_kyc&#x60;] | [default to pre_kyc]
**UserHashId** | **string** | User Unique Identifier | [default to b234cabd0cae5ea241589e4788ac727c]
**Email** | **string** | User Email Address (For Razorpay) Can be used as auto-populated contact email in their client-side form | [optional] [default to test@email.com]
**MobileCountryCode** | **int32** | User Mobile Country Code (For Razorpay) Can be used as auto-populated contact number in their client-side form | [optional] [default to null]
**Mobile** | **int32** | User Mobile Number (For Razorpay) Can be used as auto-populated contact number in their client-side form | [optional] [default to null]
**PreferredName** | **string** | (For Razorpay) Can be used as auto-populated user name in their client-side form | [optional] [default to Jane Doe]
**InRoute** | **int32** | (For Ingenico) (Required) Ingenico Product ID | [optional] [default to null]
**Currency** | **string** | Currency code [ISO 3166](https://www.iso.org/iso-3166-country-codes.html) alpha-3 | [optional] [default to SGD]
**CountryCode** | **string** | (For Ingenico) [ISO 3166](https://www.iso.org/iso-3166-country-codes.html) alpha-2 country code | [optional] [default to SG]
**Pan** | **string** | User&#x27;s Card Primary Account Number | [optional] [default to null]
**ClientRefId** | **string** | Consumer Identifier for a Payment Transaction | [optional] [default to null]
**ClientIp** | **string** | Client IP Address | [optional] 
**ActionOrigin** | **string** | Origin on where the payment transaction is executed [&#x60;admin&#x60;, &#x60;client&#x60;, &#x60;cron&#x60;, &#x60;payment_provider&#x60;] | [optional] [default to client]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

