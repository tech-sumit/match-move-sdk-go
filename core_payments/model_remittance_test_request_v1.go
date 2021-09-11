package core_payments

type RemittanceTestRequestV1 struct {
	// Retrieved from the list of supported payout agents for the program.  Based on the agent id, we should be able to get the following details:  - provider - receiving agent type (consumer or corporate) - payment_mode - routing param - routing type - send currency - receive currency - receive country 
	AgentId string `json:"agent_id,omitempty"`
	// Determines whether the amount would be the source or receive amount Acceptable values: - source - receive    
	CalculationMode string `json:"calculation_mode,omitempty"`
	// Amount to send or calculate.  - The amount precision should be in sync on the precision that is supported for the country. - For example, you send an SGD currency, the amount should have up to 2 decimal places only 
	Amount string `json:"amount,omitempty"`
}
