# RemittanceTestRequestV1

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AgentId** | **string** | Retrieved from the list of supported payout agents for the program.  Based on the agent id, we should be able to get the following details:  - provider - receiving agent type (consumer or corporate) - payment_mode - routing param - routing type - send currency - receive currency - receive country  | [optional] [default to null]
**CalculationMode** | **string** | Determines whether the amount would be the source or receive amount Acceptable values: - source - receive     | [optional] [default to null]
**Amount** | **string** | Amount to send or calculate.  - The amount precision should be in sync on the precision that is supported for the country. - For example, you send an SGD currency, the amount should have up to 2 decimal places only  | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

