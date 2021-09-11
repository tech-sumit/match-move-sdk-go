# {{classname}}

All URIs are relative to *https://beta-api.mmvpay.com/{product}/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateUserWalletCard**](CardsApi.md#CreateUserWalletCard) | **Post** /users/wallets/cards/{card_type_code} | Create Card
[**CreateUserWalletCardPin**](CardsApi.md#CreateUserWalletCardPin) | **Post** /users/wallets/cards/{card_id}/pins | Set card pin
[**DeleteUserWalletCard**](CardsApi.md#DeleteUserWalletCard) | **Delete** /users/wallets/cards/{card_id} | Suspend / Block card
[**GetUserWalletCardId**](CardsApi.md#GetUserWalletCardId) | **Get** /users/wallets/cards/{card_id} | Get card details
[**GetUserWalletCardNumer**](CardsApi.md#GetUserWalletCardNumer) | **Get** users/wallets/cards/{card_id}/number | Get the encrypted card number
[**GetUserWalletCardPins**](CardsApi.md#GetUserWalletCardPins) | **Get** /users/wallets/cards/{card_id}/pins/reset | Reset card pin
[**GetUserWalletCardSecuritiesTokens**](CardsApi.md#GetUserWalletCardSecuritiesTokens) | **Get** /users/wallets/cards/{card_id}/securities/tokens | Retrieve card security code
[**GetUserWalletCardTypes**](CardsApi.md#GetUserWalletCardTypes) | **Get** /users/wallets/cards/types | Get configured card types
[**GetUserWalletCards**](CardsApi.md#GetUserWalletCards) | **Get** /users/wallets/cards | Get cards for user
[**PutUserWalletCardId**](CardsApi.md#PutUserWalletCardId) | **Put** /users/wallets/cards/{card_id} | Activate card
[**ReactivateUserWalletCardId**](CardsApi.md#ReactivateUserWalletCardId) | **Post** /users/wallets/cards/{card_id} | Re-activate card

# **CreateUserWalletCard**
> CardFullV1 CreateUserWalletCard(ctx, assocNumber, refId, cardDesign, nameOnCard, additionalDetails, autoActivate, var2faMethod, var2faDelivery, var2faValue, cardTypeCode, xAuthUserID)
Create Card

This Api is used for fulfilling the following use cases:   - Creating Virtual Card in the system   - Linking a non personalized card to the wallet of the user in the system   - Linking and activating a non personalized card to the wallet of the user in the system (requires `auto_activate` parameter with value `true` passed)   - Placing a personalized card order in the system  <!-- theme: info --> > ### Note > If mask_number config is enabled the `number` will be in the encrypted form (123456 XX XXXX 1234).   Learn more about [cards resource](https://matchmove.stoplight.io/docs/optimus-prime/docs/Guides/Cards/Cards-Overview.md) 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **assocNumber** | **string**|  | 
  **refId** | **string**|  | 
  **cardDesign** | **string**|  | 
  **nameOnCard** | **string**|  | 
  **additionalDetails** | **string**|  | 
  **autoActivate** | **bool**|  | 
  **var2faMethod** | **string**|  | 
  **var2faDelivery** | **string**|  | 
  **var2faValue** | **string**|  | 
  **cardTypeCode** | **string**| The card type available for the program in the system. Available choices can be retrieved via [Get configured card types](https://developer.matchmove.com/docs/optimus-prime/op_api.yaml/paths/~1users~1wallets~1cards~1types/get) | 
  **xAuthUserID** | **string**| MatchMove provided hash ID for the user | 

### Return type

[**CardFullV1**](Card_full.v1.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: application/x-www-form-urlencoded
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CreateUserWalletCardPin**
> InlineResponse2007 CreateUserWalletCardPin(ctx, pinblock, xAuthUserID, cardId)
Set card pin

Setting a pin for the user in the system. Due to security considerations and guidelines imposed by PCIDSS, the PIN can never be transmitted using clear text.   To use this API , clients will be provided with a ZPK , that would be used for generating the Pinblock in a secure manner.  In addition for security considerations and prevent fraud. There are thresholds implemented with respect to the count of PIN resets that can be attempted within a given period.  *Note* It is highly recommended to implement a KMS or HSM for storage of the ZPK provisioned for you.  The encryption logic is as follows  PINBLOCK FORMAT : ISO FORMAT ZERO  Encryption Method : ANSI X 9.17 – Triple DES – ECB  Key Management : Static PIN Encryption Key  Key Length : Double Length (32 hexadecimal characters)  Sample Data with Input and Output  ### Input PIN Encryption Key (Clear) : `0123456789ABCDEF0123456789ABCDEF` Pin entered by Customer : `1234` Mobile no : `00 65 12345678` Outgoing PINBLOCK(ISO FORMAT ZERO) : `1D51A078AC3F7081`  *** mobile number will be always 12 digits, lets say if system was having a mobile number with only 10 digits (6512345678) then should we consider (006512345678)  *Phone Number (12 digits)*: `123456789012`  *Key (16 or 24 bytes, 16 bytes in this example)*: `112233445566778899AABBCCDDEEFF00`  *Padding (DCBA) + Phone Number + ~Padding () + ~Phone Number*: `DCBA` + `123456789012` + `2345` + `EDCBA9876FED Diversified Key (3DES CBC): 4CB178CC3FB92FB8693D31F1A39DB8C9`    *Pin*: `1234`  *Data* 1: `0 4 1234 FFFFFFFFFF`  *Data 2*: `0000 123456789012`  *XOR (Data 1, Data 2)*: `041226CBA9876FE`  *Final PIN Block*: `4D316E71853CFB3A` 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **pinblock** | **string**|  | 
  **xAuthUserID** | **string**| The &#x60;user_id&#x60; of the user returned via  GET /users  | 
  **cardId** | [**string**](.md)| The &#x60;card_id&#x60; for which the information is being retrieved.  | 

### Return type

[**InlineResponse2007**](inline_response_200_7.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: application/x-www-form-urlencoded
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteUserWalletCard**
> InlineResponse2006 DeleteUserWalletCard(ctx, cardId, optional)
Suspend / Block card

There are two kinds of suspensions which are supported in the system currently. Depending on the state passed as part of the request the behavior changes. ### Suspended Suspends the card prevents fresh authorizations on that card. If the card has been already pre-validated by a merchant, an authorization can still be made. ### Blocked Blocking cards will disable all transaction made by that card, and all purchase authorization. <!-- theme: info --> > ### Note > A card with pending settlements cannot be blocked. You are recommended to first suspend the account and wait for pending settlements before blocking  <!-- theme: info --> > ### Note > If mask_number config is enabled the `number` will be in the encrypted form (123456 XX XXXX 1234).   <!-- theme: danger --> > ### Note > Passing values `lost`, `stolen` and `damaged`  will permanently block the card and is irreversible 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **cardId** | **string**| Cards id which can be obtained using GET &#x60;/users/wallets&#x60; or GET &#x60;/users/wallets/cards&#x60;  | 
 **optional** | ***CardsApiDeleteUserWalletCardOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a CardsApiDeleteUserWalletCardOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **type_** | **optional.**|  | 

### Return type

[**InlineResponse2006**](inline_response_200_6.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: application/x-www-form-urlencoded
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetUserWalletCardId**
> CardFullV1 GetUserWalletCardId(ctx, cardId, maskNumber, xAuthUserID)
Get card details

Retrieve the details of a particular card in the system via card_id.  <!-- theme: info --> > ### Note > If mask_number config is enabled the `number` will be in the encrypted form (123456 XX XXXX 1234). 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **cardId** | **string**| Cards id which can be obtained using GET &#x60;/users/wallets&#x60; or GET &#x60;/users/wallets/cards&#x60;  | 
  **maskNumber** | **string**|  | 
  **xAuthUserID** | **string**| MatchMove provided hash ID for the user | 

### Return type

[**CardFullV1**](Card_full.v1.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetUserWalletCardNumer**
> InlineResponse20013 GetUserWalletCardNumer(ctx, cardId, xAuthUserID)
Get the encrypted card number

Get the encrypted card number 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **cardId** | **string**| The &#x60;card_id&#x60; for which the information is being retrieved.  | 
  **xAuthUserID** | **string**| MatchMove provided hash ID for the user | 

### Return type

[**InlineResponse20013**](inline_response_200_13.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetUserWalletCardPins**
> InlineResponse2007 GetUserWalletCardPins(ctx, cardId, mode)
Reset card pin

Reset Card Pin 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **cardId** | [**string**](.md)| The &#x60;card_id&#x60; for which the information is being retrieved.  | 
  **mode** | **string**| Mode denotes how the system generated PIN is sent to the end user.  | Value | Name | Description | |-----|-----|-------| | M   | Email   | The PIN is delivered via email   | | S   | SMS     | The PIN is delivered via SMS   |  | 

### Return type

[**InlineResponse2007**](inline_response_200_7.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetUserWalletCardSecuritiesTokens**
> InlineResponse2008 GetUserWalletCardSecuritiesTokens(ctx, xAuthUserID, cardId)
Retrieve card security code

All card entities attached to the users wallet are industry standard cards which come with their own card security code to process card not present transactions to reduce fraud and limit the issuer liability.  Depending on the network they are referred by scheme specific names like - CSC (American Express) - CVC (MasterCard) - CVV (Visa) - CVV (Rupay)  These are typically the 3 or 4 digit numbers that are printed at the back of the card. Do note the networks are the one who dictate printing of these numbers on the back of the card. Lately with the launch of newer card products , they are more open to not showing this information as part of the printed card.  Virtual cards since mirror the physical cards in all respects , are also provided with this security code and its up to the program manager how to implement the display of the same.  In addition for Singapore, MatchMove has also received approval from the Scheme partner (MasterCard) to have a dynamic CVC in which case the value retrieved via API is re-generated on each request. This CVC generated is only valid for one time use or 10 minutes. This provides an additional layer of security , though admittedly it introduces a greater friction in checkout experience for the end users. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **xAuthUserID** | **string**| MatchMove provided hash ID for the user | 
  **cardId** | [**string**](.md)| The &#x60;card_id&#x60; for which the information is being retrieved.  | 

### Return type

[**InlineResponse2008**](inline_response_200_8.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetUserWalletCardTypes**
> InlineResponse2002 GetUserWalletCardTypes(ctx, )
Get configured card types

Retrieve the list of card types supported in the system

### Required Parameters
This endpoint does not need any parameter.

### Return type

[**InlineResponse2002**](inline_response_200_2.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetUserWalletCards**
> InlineResponse2003 GetUserWalletCards(ctx, maskNumber, xAuthUserID)
Get cards for user

Retrieve the list of card types supported in the system <!-- theme: danger -->  > ### Deprecation Notice > > Do note this API is no longer actively supported and partners are expected to consume get /users/wallets/cards/card_id explicitly.  <!-- theme: info --> > ### Note > If mask_number config is enabled the `number` will be in the encrypted form (123456 XX XXXX 1234).  

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **maskNumber** | **string**|  | 
  **xAuthUserID** | **string**| MatchMove provided hash ID for the user | 

### Return type

[**InlineResponse2003**](inline_response_200_3.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PutUserWalletCardId**
> InlineResponse2004 PutUserWalletCardId(ctx, activationCode, var2faMethod, var2faDelivery, var2faValue, cardId)
Activate card

Process for activation of physical card (Personalized / Non Personalized). This API is not applicable for virtual card type. If 2fa is enabled, then in response the details related to the same is returned. <!-- theme: info --> > ### Note > Do note if `auto_activate=true` is passed during `POST /users/wallets/{card_type_code}` this step is not applicable. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **activationCode** | **string**|  | 
  **var2faMethod** | **string**|  | 
  **var2faDelivery** | **string**|  | 
  **var2faValue** | **string**|  | 
  **cardId** | **string**| Cards id which can be obtained using GET &#x60;/users/wallets&#x60; or GET &#x60;/users/wallets/cards&#x60;  | 

### Return type

[**InlineResponse2004**](inline_response_200_4.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: application/x-www-form-urlencoded
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ReactivateUserWalletCardId**
> InlineResponse2005 ReactivateUserWalletCardId(ctx, id, assocNumber, refId, cardId)
Re-activate card

Re-Activate a locked card  <!-- theme: info --> > ### Note > If mask_number config is enabled the `number` will be in the encrypted form (123456 XX XXXX 1234).  

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**|  | 
  **assocNumber** | **string**|  | 
  **refId** | **string**|  | 
  **cardId** | **string**| Cards id which can be obtained using GET &#x60;/users/wallets&#x60; or GET &#x60;/users/wallets/cards&#x60;  | 

### Return type

[**InlineResponse2005**](inline_response_200_5.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: application/x-www-form-urlencoded
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

