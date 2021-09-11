# Refunds

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **int32** | Refund ID | [default to null]
**CreditId** | **int32** | Topup/Credit ID | [default to null]
**PaymentRefId** | **string** | Payment reference ID | [default to null]
**RefundRefId** | **string** | Refund reference ID | [default to null]
**ProviderRefId** | **string** | Provider reference ID | [default to null]
**Type_** | **string** | Refund type (full, partial) | [default to null]
**Action** | **string** | During which action refund transaction is created [topup, stale, blacklist, cron, settlement] | [default to null]
**Amount** | **int32** | Refund amount | [default to null]
**Reason** | **string** | reason for refund | [default to null]
**Attempts** | **int32** | Number of refund attempts | [default to null]
**Status** | **int32** | Status [&#x60;0&#x60;, &#x60;1&#x60;, &#x60;2&#x60;, &#x60;3&#x60;, &#x60;4&#x60;, &#x60;5&#x60;, &#x60;6&#x60;] | [default to null]
**DateCreated** | **string** | Date created &#x60;(YYYY-MM-DD HH:MM:SS)&#x60; | [default to null]
**DateModified** | **string** | Date modified &#x60;(YYYY-MM-DD HH:MM:SS)&#x60; | [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

