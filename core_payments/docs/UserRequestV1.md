# UserRequestV1

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** |  | [optional] [default to null]
**FirstName** | **string** | First name or given name | [optional] [default to null]
**LastName** | **string** | Last name or family name | [optional] [default to null]
**MiddleName** | **string** | User Middle name | [optional] [default to null]
**PreferredName** | **string** | Prefereed name | [optional] [default to null]
**Email** | **string** | Registration email for the user to be created in the MatchMove system for a particular program. | [optional] [default to null]
**MobileCountryCode** | **string** | Supported Country codes for the program can be retrieved from &#x60;GET /users/enumerations/mobile_country_codes&#x60;  | [optional] [default to null]
**Password** | **string** | Minimum 8 Characters , One special character , One number , One small character  | [optional] [default to null]
**Mobile** | **string** | Mobile numbe rof the user | [optional] [default to null]
**Title** | **string** | Title to be set like mr/miss etc | [optional] [default to null]
**IdType** | **string** | Id type to be set. Available values can be retrieved from &#x60;GET /users/enumerations/id_types&#x60; | [optional] [default to null]
**IdNumber** | **string** | Id number to be unique | [optional] [default to null]
**IdDateExpiry** | **string** | Id type expiry date | [optional] [default to null]
**IdDateIssue** | **string** | Issued date of id type | [optional] [default to null]
**CountryOfIssue** | **string** | Country that issued your id type | [optional] [default to null]
**AltIdType** | **string** | alternative id type | [optional] [default to null]
**AltIdNumber** | **string** | Alertnative id type number to be unique | [optional] [default to null]
**AltIdDateExpiry** | **string** | Alternative id type expiry date | [optional] [default to null]
**AltIdDateIssue** | **string** | Alternative id type issued date | [optional] [default to null]
**AltIdCountryOfIssue** | **string** | Alternative Id type isunace country | [optional] [default to null]
**Birthday** | **string** | Birthday of the user registred in matchmove system | [optional] [default to null]
**PlaceOfBirth** | **string** | Place where you have born | [optional] [default to null]
**Nationality** | **string** | User Nationality. Supported values can be retrieved from &#x60;GET /users/enumerations/nationalities&#x60; | [optional] [default to null]
**Gender** | **string** | User Gender. Supported values can be retreived from &#x60;GET /users/enumerations/genders&#x60; | [optional] [default to null]
**MaritalStatus** | **string** | User marital Status. Supported values can be retreived from &#x60;GET /users/enumerations/marital_status&#x60; | [optional] [default to null]
**MothersMaidenName** | **string** | Mother&#x27;s maiden name | [optional] [default to null]
**NatureOfBusiness** | **string** | Nature of your business | [optional] [default to null]
**CustomerId** | **string** | Customer id to be passed | [optional] [default to null]
**PurposeOfAccount** | **string** | Purpose of your account | [optional] [default to null]
**PartnerId** | **string** | Partner id to be passed | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

