# Consumers

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ConsumerHash** | **string** | Consumer Hash ID | [default to null]
**AclKey** | **string** | Consumer ACK Key | [default to null]
**Code** | **string** | Consumer code name | [default to null]
**Description** | **string** | Consumer description | [default to null]
**PreCreditConfirm** | **string** | Will call a client&#x27;s external pre-credit url before doing payment and topup [&#x60;0&#x60;, &#x60;1&#x60;] | [optional] [default to 0]
**PreCreditUrlMethod** | **string** | Client&#x27;s external pre-credit url method to call before doing payment and topup [&#x60;GET&#x60;, &#x60;POST&#x60;, &#x60;PUT&#x60;, &#x60;DELETE&#x60;] | [optional] 
**PreCreditUrl** | **string** | Client&#x27;s external pre-credit url to call before doing payment and topup | [optional] 
**PostCreditConfirm** | **string** | Will call a client&#x27;s external pre-credit url after doing payment and topup [&#x60;0&#x60;, &#x60;1&#x60;] | [optional] [default to 0]
**PostCreditUrlMethod** | **string** | Client&#x27;s external pre-credit url method to call after doing payment and topup [&#x60;GET&#x60;, &#x60;POST&#x60;, &#x60;PUT&#x60;, &#x60;DELETE&#x60;] | [optional] 
**PostCreditUrl** | **string** | Client&#x27;s external pre-credit url to call after doing payment and topup | [optional] 
**Timezone** | **string** | Timezone in [TZ format](https://en.wikipedia.org/wiki/List_of_tz_database_time_zones) | [optional] [default to Asia/Singapore]
**TimeOffset** | **string** | Time Offset in [UTC format](https://en.wikipedia.org/wiki/List_of_tz_database_time_zones) | [optional] [default to +08:00]
**Status** | **string** | Status [&#x60;active&#x60;, &#x60;inactive&#x60;, &#x60;terminated&#x60;] | [default to active]
**IsAdmin** | **int32** | Is admin | [default to null]
**DateCreated** | **string** | Date created &#x60;(YYYY-MM-DD HH:MM:SS)&#x60; | [optional] [default to null]
**DateModified** | **string** | Date modified &#x60;(YYYY-MM-DD HH:MM:SS)&#x60; | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

