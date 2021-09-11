package core_payments

// Available parameters supported for creating remittance transactions.  All supported parameters will be based on the beneficiary/sender form from the calculate endpoint (/test)
type RemittanceSendRequestV1 struct {
	// Payout Agent ID; required on the new remittance endpoint (where no rails is provided)
	AgentId string `json:"agent_id,omitempty"`
	// Determines whether the amount would be the source or receive amount.
	CalculationMode string `json:"calculation_mode,omitempty"`
	// Transaction Type
	TransactionType string `json:"transaction_type,omitempty"`
	// Transfer amount
	Amount string `json:"amount,omitempty"`
	// Send Amount Currency Code
	SendCurrency string `json:"send_currency,omitempty"`
	// Receive Amount Currency Code
	ReceiveCurrency string `json:"receive_currency,omitempty"`
	// Receiving Country
	ReceiveCountry string `json:"receive_country,omitempty"`
	// Payment Mode
	PaymentMode string `json:"payment_mode,omitempty"`
	// Routing Type
	RoutingType string `json:"routing_type,omitempty"`
	// Routing Param
	RoutingParam string `json:"routing_param,omitempty"`
	// Quotation ID
	QuotationId string `json:"quotation_id,omitempty"`
	// Payout Agent ID
	PayoutAgent string `json:"payout_agent,omitempty"`
	// User ID of the sender
	HashId string `json:"hash_id,omitempty"`
	// Receiver First Name
	FirstName string `json:"first_name,omitempty"`
	// Receiver Middle Name
	MiddleName string `json:"middle_name,omitempty"`
	// Receiver Last Name
	LastName string `json:"last_name,omitempty"`
	// Receiver Chinese Name
	ChineseName string `json:"chinese_name,omitempty"`
	// Receiver Mobile Number
	ReceiveMobileNumber string `json:"receive_mobile_number,omitempty"`
	// Receiver Gender
	ReceiveGender string `json:"receive_gender,omitempty"`
	// Receiver ID Number
	IdNumber string `json:"id_number,omitempty"`
	// Receiver ID Type
	IdType string `json:"id_type,omitempty"`
	// Receiver Occupation
	Occupation string `json:"occupation,omitempty"`
	// Receiver Address 1
	Address1 string `json:"address_1,omitempty"`
	// Receiver Address 2
	Address2 string `json:"address_2,omitempty"`
	// Receiver birth date
	BirthDate string `json:"birth_date,omitempty"`
	// Receiver Address City
	City string `json:"city,omitempty"`
	// Receiver Address State
	State string `json:"state,omitempty"`
	// Receiver Address Country
	Country string `json:"country,omitempty"`
	// Receiver Nationality
	Nationality string `json:"nationality,omitempty"`
	// Receiver Zipcode
	Zipcode string `json:"zipcode,omitempty"`
	// Beneficiary Hash ID
	BeneficiaryHashId string `json:"beneficiary_hash_id,omitempty"`
	// Receiving Business Registered Name
	ReceivingBusinessRegisteredName string `json:"receiving_business_registered_name,omitempty"`
	// Receiving Business Trading Name
	ReceivingBusinessTradingName string `json:"receiving_business_trading_name,omitempty"`
	// Receiving Business Email
	ReceivingBusinessEmail string `json:"receiving_business_email,omitempty"`
	// Receiving Business MSISDN
	ReceivingBusinessMsisdn string `json:"receiving_business_msisdn,omitempty"`
	// Receiving Business Registration Number
	ReceivingBusinessRegistrationNumber string `json:"receiving_business_registration_number,omitempty"`
	// Receiving Business Tax ID
	ReceivingBusinessTaxId string `json:"receiving_business_tax_id,omitempty"`
	// Receiving Business Address 1
	ReceivingBusinessAddress1 string `json:"receiving_business_address1,omitempty"`
	// Receiving Business Address 2
	ReceivingBusinessAddress2 string `json:"receiving_business_address2,omitempty"`
	// Receiving Business Address City
	ReceivingBusinessCity string `json:"receiving_business_city,omitempty"`
	// Receiving Business Address State
	ReceivingBusinessState string `json:"receiving_business_state,omitempty"`
	// Receiving Business Address Country
	ReceivingBusinessCountry string `json:"receiving_business_country,omitempty"`
	// Receiving Business Representative First Name
	ReceivingBusinessRepresentativeFirstName string `json:"receiving_business_representative_first_name,omitempty"`
	// Receiving Business Representative Middle Name
	ReceivingBusinessRepresentativeMiddleName string `json:"receiving_business_representative_middle_name,omitempty"`
	// Receiving Business Representative Last Name
	ReceivingBusinessRepresentativeLastName string `json:"receiving_business_representative_last_name,omitempty"`
	// Receiving Business Representative Last Name 2
	ReceivingBusinessRepresentativeLastName2 string `json:"receiving_business_representative_last_name2,omitempty"`
	// Receiving Business Representative Native Name
	ReceivingBusinessRepresentativeNativeName string `json:"receiving_business_representative_native_name,omitempty"`
	// Receiving Business Representative ID Type
	ReceivingBusinessIdType string `json:"receiving_business_id_type,omitempty"`
	// Receiving Business Representative ID Number
	ReceivingBusinessIdNumber string `json:"receiving_business_id_number,omitempty"`
	// Receiving Business Representative ID Country of Issue
	ReceivingBusinessCountryOfIssue string `json:"receiving_business_country_of_issue,omitempty"`
	// Receiving Business Representative ID Issue Date
	ReceivingBusinessIssueDate string `json:"receiving_business_issue_date,omitempty"`
	// Receiving Business Representative ID Expiry Date
	ReceivingBusinessExpireDate string `json:"receiving_business_expire_date,omitempty"`
	// Receiving Business Date of Incorporation
	ReceivingBusinessDateOfIncorporation string `json:"receiving_business_date_of_incorporation,omitempty"`
	// Sending Business Hash ID
	SendingBusinessHashId string `json:"sending_business_hash_id,omitempty"`
	// Sending Business Registered Name
	SendingBusinessRegisteredName string `json:"sending_business_registered_name,omitempty"`
	// Sending Business Trading Name
	SendingBusinessTradingName string `json:"sending_business_trading_name,omitempty"`
	// Sending Business Email
	SendingBusinessEmail string `json:"sending_business_email,omitempty"`
	// Sending Business MSISDN
	SendingBusinessMsisdn string `json:"sending_business_msisdn,omitempty"`
	// Sending Business Registration Number
	SendingBusinessRegistrationNumber string `json:"sending_business_registration_number,omitempty"`
	// Sending Business Tax ID
	SendingBusinessTaxId string `json:"sending_business_tax_id,omitempty"`
	// Sending Business Address 1
	SendingBusinessAddress1 string `json:"sending_business_address1,omitempty"`
	// Sending Business Address 2
	SendingBusinessAddress2 string `json:"sending_business_address2,omitempty"`
	// Sending Business Address City
	SendingBusinessCity string `json:"sending_business_city,omitempty"`
	// Sending Business Address State
	SendingBusinessState string `json:"sending_business_state,omitempty"`
	// Sending Business Address Zipcode
	SendingBusinessZipcode string `json:"sending_business_zipcode,omitempty"`
	// Sending Business Address Country
	SendingBusinessCountry string `json:"sending_business_country,omitempty"`
	// Sending Business Representative First Name
	SendingBusinessRepresentativeFirstName string `json:"sending_business_representative_first_name,omitempty"`
	// Sending Business Representative Middle Name
	SendingBusinessRepresentativeMiddleName string `json:"sending_business_representative_middle_name,omitempty"`
	// Sending Business Representative Last Name
	SendingBusinessRepresentativeLastName string `json:"sending_business_representative_last_name,omitempty"`
	// Sending Business Representative Last Name 2
	SendingBusinessRepresentativeLastName2 string `json:"sending_business_representative_last_name2,omitempty"`
	// Sending Business Representative Native Name
	SendingBusinessRepresentativeNativeName string `json:"sending_business_representative_native_name,omitempty"`
	// Sending Business Representative ID Type
	SendingBusinessIdType string `json:"sending_business_id_type,omitempty"`
	// Sending Business Representative ID Number
	SendingBusinessIdNumber string `json:"sending_business_id_number,omitempty"`
	// Sending Business Representative ID Country of Issue
	SendingBusinessCountryOfIssue string `json:"sending_business_country_of_issue,omitempty"`
	// Sending Business Representative ID Issue Date
	SendingBusinessIssueDate string `json:"sending_business_issue_date,omitempty"`
	// Sending Business Representative ID Expiry Date
	SendingBusinessExpireDate string `json:"sending_business_expire_date,omitempty"`
	// Sending Business Date of Incoporation
	SendingBusinessDateOfIncorporation string `json:"sending_business_date_of_incorporation,omitempty"`
	// Document Reference Number
	DocumentReferenceNumber string `json:"document_reference_number,omitempty"`
	// Swift code
	SwiftCode string `json:"swift_code,omitempty"`
	// Indian Bank IFC Code
	BankIfcCode string `json:"bank_ifc_code,omitempty"`
	// Branch name of the recipient bank account
	BankBranchName string `json:"bank_branch_name,omitempty"`
	// Branch code of the recipient bank account
	BankBranchCode string `json:"bank_branch_code,omitempty"`
	// Bank account number if paying via bank
	BankAccountNumber string `json:"bank_account_number,omitempty"`
	// CBC bank name of the recipient bank account
	CbcBankName string `json:"cbc_bank_name,omitempty"`
	// BPI Bank name of the recipient bank account
	BpiBankName string `json:"bpi_bank_name,omitempty"`
	// Partner name of the service provider
	PartnerName string `json:"partner_name,omitempty"`
	// Source of income
	SourceOfIncome string `json:"source_of_income,omitempty"`
	// Relationship with the Receiver
	Relationship string `json:"relationship,omitempty"`
	// Reason
	Reason string `json:"reason,omitempty"`
	// Sender reference
	SenderReference string `json:"sender_reference,omitempty"`
	// Additional Information
	AdditionalInformation string `json:"additional_information,omitempty"`
}
