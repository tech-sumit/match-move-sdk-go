
package core_payments

import (
	"context"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
	"fmt"
	"github.com/antihax/optional"
)

// Linger please
var (
	_ context.Context
)

type OauthUtilitiesReportingApiService service
/*
OauthUtilitiesReportingApiService Retrieve all payments instruments of the users in the system
A lightweight endpoint to return the User list of payment instruments. Suggested to use in case of repeat calls.
 * @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param optional nil or *OauthUtilitiesReportingApiGetConsumerUserInstrumentsOpts - Optional Parameters:
     * @param "UserId" (optional.String) -  User hash ID of the user returned from the MatchMove system
     * @param "MobileCountryCode" (optional.String) -  If mobile is passed this is a required parameter
     * @param "Mobile" (optional.String) -  if mobile_country_code is passed this is a required parameter
     * @param "CardTypeCode" (optional.String) -  The list of supported values is retrievable by calling GET /users/wallets/cards/types
     * @param "CardId" (optional.String) -  Card hash ID of the user returned from card creation
@return InlineResponse20014
*/

type OauthUtilitiesReportingApiGetConsumerUserInstrumentsOpts struct {
    UserId optional.String
    MobileCountryCode optional.String
    Mobile optional.String
    CardTypeCode optional.String
    CardId optional.String
}

func (a *OauthUtilitiesReportingApiService) GetConsumerUserInstruments(ctx context.Context, localVarOptionals *OauthUtilitiesReportingApiGetConsumerUserInstrumentsOpts) (InlineResponse20014, *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Get")
		localVarPostBody   interface{}
		localVarFileName   string
		localVarFileBytes  []byte
		localVarReturnValue InlineResponse20014
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath + "/oauth/consumer/users/instruments"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if localVarOptionals != nil && localVarOptionals.UserId.IsSet() {
		localVarQueryParams.Add("user_id", parameterToString(localVarOptionals.UserId.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.MobileCountryCode.IsSet() {
		localVarQueryParams.Add("mobile_country_code", parameterToString(localVarOptionals.MobileCountryCode.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.Mobile.IsSet() {
		localVarQueryParams.Add("mobile", parameterToString(localVarOptionals.Mobile.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.CardTypeCode.IsSet() {
		localVarQueryParams.Add("card_type_code", parameterToString(localVarOptionals.CardTypeCode.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.CardId.IsSet() {
		localVarQueryParams.Add("card_id", parameterToString(localVarOptionals.CardId.Value(), ""))
	}
	// to determine the Content-Type header
	localVarHttpContentTypes := []string{}

	// set Content-Type header
	localVarHttpContentType := selectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHttpContentType
	}

	// to determine the Accept header
	localVarHttpHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHttpHeaderAccept := selectHeaderAccept(localVarHttpHeaderAccepts)
	if localVarHttpHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHttpHeaderAccept
	}
	r, err := a.client.prepareRequest(ctx, localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFileName, localVarFileBytes)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return localVarReturnValue, localVarHttpResponse, err
	}

	localVarBody, err := ioutil.ReadAll(localVarHttpResponse.Body)
	localVarHttpResponse.Body.Close()
	if err != nil {
		return localVarReturnValue, localVarHttpResponse, err
	}

	if localVarHttpResponse.StatusCode < 300 {
		// If we succeed, return the data, otherwise pass on to decode error.
		err = a.client.decode(&localVarReturnValue, localVarBody, localVarHttpResponse.Header.Get("Content-Type"));
		if err == nil { 
			return localVarReturnValue, localVarHttpResponse, err
		}
	}

	if localVarHttpResponse.StatusCode >= 300 {
		newErr := GenericSwaggerError{
			body: localVarBody,
			error: localVarHttpResponse.Status,
		}
		if localVarHttpResponse.StatusCode == 200 {
			var v InlineResponse20014
			err = a.client.decode(&v, localVarBody, localVarHttpResponse.Header.Get("Content-Type"));
				if err != nil {
					newErr.error = err.Error()
					return localVarReturnValue, localVarHttpResponse, newErr
				}
				newErr.model = v
				return localVarReturnValue, localVarHttpResponse, newErr
		}
		if localVarHttpResponse.StatusCode == 400 {
			var v ErrorV1
			err = a.client.decode(&v, localVarBody, localVarHttpResponse.Header.Get("Content-Type"));
				if err != nil {
					newErr.error = err.Error()
					return localVarReturnValue, localVarHttpResponse, newErr
				}
				newErr.model = v
				return localVarReturnValue, localVarHttpResponse, newErr
		}
		if localVarHttpResponse.StatusCode == 401 {
			var v ErrorV1
			err = a.client.decode(&v, localVarBody, localVarHttpResponse.Header.Get("Content-Type"));
				if err != nil {
					newErr.error = err.Error()
					return localVarReturnValue, localVarHttpResponse, newErr
				}
				newErr.model = v
				return localVarReturnValue, localVarHttpResponse, newErr
		}
		if localVarHttpResponse.StatusCode == 403 {
			var v ErrorV1
			err = a.client.decode(&v, localVarBody, localVarHttpResponse.Header.Get("Content-Type"));
				if err != nil {
					newErr.error = err.Error()
					return localVarReturnValue, localVarHttpResponse, newErr
				}
				newErr.model = v
				return localVarReturnValue, localVarHttpResponse, newErr
		}
		if localVarHttpResponse.StatusCode == 404 {
			var v ErrorV1
			err = a.client.decode(&v, localVarBody, localVarHttpResponse.Header.Get("Content-Type"));
				if err != nil {
					newErr.error = err.Error()
					return localVarReturnValue, localVarHttpResponse, newErr
				}
				newErr.model = v
				return localVarReturnValue, localVarHttpResponse, newErr
		}
		if localVarHttpResponse.StatusCode == 500 {
			var v ErrorV1
			err = a.client.decode(&v, localVarBody, localVarHttpResponse.Header.Get("Content-Type"));
				if err != nil {
					newErr.error = err.Error()
					return localVarReturnValue, localVarHttpResponse, newErr
				}
				newErr.model = v
				return localVarReturnValue, localVarHttpResponse, newErr
		}
		if localVarHttpResponse.StatusCode == 502 {
			var v ErrorV1
			err = a.client.decode(&v, localVarBody, localVarHttpResponse.Header.Get("Content-Type"));
				if err != nil {
					newErr.error = err.Error()
					return localVarReturnValue, localVarHttpResponse, newErr
				}
				newErr.model = v
				return localVarReturnValue, localVarHttpResponse, newErr
		}
		if localVarHttpResponse.StatusCode == 503 {
			var v ErrorV1
			err = a.client.decode(&v, localVarBody, localVarHttpResponse.Header.Get("Content-Type"));
				if err != nil {
					newErr.error = err.Error()
					return localVarReturnValue, localVarHttpResponse, newErr
				}
				newErr.model = v
				return localVarReturnValue, localVarHttpResponse, newErr
		}
		if localVarHttpResponse.StatusCode == 504 {
			var v ErrorV1
			err = a.client.decode(&v, localVarBody, localVarHttpResponse.Header.Get("Content-Type"));
				if err != nil {
					newErr.error = err.Error()
					return localVarReturnValue, localVarHttpResponse, newErr
				}
				newErr.model = v
				return localVarReturnValue, localVarHttpResponse, newErr
		}
		return localVarReturnValue, localVarHttpResponse, newErr
	}

	return localVarReturnValue, localVarHttpResponse, nil
}
/*
OauthUtilitiesReportingApiService Get Consumer transaction
Get Consumer Transactions 
 * @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param consumerKey
 * @param page
@return []OauthTransactionV1
*/
func (a *OauthUtilitiesReportingApiService) GetOauthConsumerConsumerKeyTransactionsPage(ctx context.Context, consumerKey string, page string) ([]OauthTransactionV1, *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Get")
		localVarPostBody   interface{}
		localVarFileName   string
		localVarFileBytes  []byte
		localVarReturnValue []OauthTransactionV1
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath + "/oauth/consumer/{consumer_key}/transactions/{page}"
	localVarPath = strings.Replace(localVarPath, "{"+"consumer_key"+"}", fmt.Sprintf("%v", consumerKey), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"page"+"}", fmt.Sprintf("%v", page), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	// to determine the Content-Type header
	localVarHttpContentTypes := []string{}

	// set Content-Type header
	localVarHttpContentType := selectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHttpContentType
	}

	// to determine the Accept header
	localVarHttpHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHttpHeaderAccept := selectHeaderAccept(localVarHttpHeaderAccepts)
	if localVarHttpHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHttpHeaderAccept
	}
	r, err := a.client.prepareRequest(ctx, localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFileName, localVarFileBytes)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return localVarReturnValue, localVarHttpResponse, err
	}

	localVarBody, err := ioutil.ReadAll(localVarHttpResponse.Body)
	localVarHttpResponse.Body.Close()
	if err != nil {
		return localVarReturnValue, localVarHttpResponse, err
	}

	if localVarHttpResponse.StatusCode < 300 {
		// If we succeed, return the data, otherwise pass on to decode error.
		err = a.client.decode(&localVarReturnValue, localVarBody, localVarHttpResponse.Header.Get("Content-Type"));
		if err == nil { 
			return localVarReturnValue, localVarHttpResponse, err
		}
	}

	if localVarHttpResponse.StatusCode >= 300 {
		newErr := GenericSwaggerError{
			body: localVarBody,
			error: localVarHttpResponse.Status,
		}
		if localVarHttpResponse.StatusCode == 200 {
			var v []OauthTransactionV1
			err = a.client.decode(&v, localVarBody, localVarHttpResponse.Header.Get("Content-Type"));
				if err != nil {
					newErr.error = err.Error()
					return localVarReturnValue, localVarHttpResponse, newErr
				}
				newErr.model = v
				return localVarReturnValue, localVarHttpResponse, newErr
		}
		if localVarHttpResponse.StatusCode == 400 {
			var v ErrorV1
			err = a.client.decode(&v, localVarBody, localVarHttpResponse.Header.Get("Content-Type"));
				if err != nil {
					newErr.error = err.Error()
					return localVarReturnValue, localVarHttpResponse, newErr
				}
				newErr.model = v
				return localVarReturnValue, localVarHttpResponse, newErr
		}
		if localVarHttpResponse.StatusCode == 403 {
			var v ErrorV1
			err = a.client.decode(&v, localVarBody, localVarHttpResponse.Header.Get("Content-Type"));
				if err != nil {
					newErr.error = err.Error()
					return localVarReturnValue, localVarHttpResponse, newErr
				}
				newErr.model = v
				return localVarReturnValue, localVarHttpResponse, newErr
		}
		if localVarHttpResponse.StatusCode == 404 {
			var v ErrorV1
			err = a.client.decode(&v, localVarBody, localVarHttpResponse.Header.Get("Content-Type"));
				if err != nil {
					newErr.error = err.Error()
					return localVarReturnValue, localVarHttpResponse, newErr
				}
				newErr.model = v
				return localVarReturnValue, localVarHttpResponse, newErr
		}
		if localVarHttpResponse.StatusCode == 500 {
			var v ErrorV1
			err = a.client.decode(&v, localVarBody, localVarHttpResponse.Header.Get("Content-Type"));
				if err != nil {
					newErr.error = err.Error()
					return localVarReturnValue, localVarHttpResponse, newErr
				}
				newErr.model = v
				return localVarReturnValue, localVarHttpResponse, newErr
		}
		if localVarHttpResponse.StatusCode == 501 {
			var v ErrorV1
			err = a.client.decode(&v, localVarBody, localVarHttpResponse.Header.Get("Content-Type"));
				if err != nil {
					newErr.error = err.Error()
					return localVarReturnValue, localVarHttpResponse, newErr
				}
				newErr.model = v
				return localVarReturnValue, localVarHttpResponse, newErr
		}
		if localVarHttpResponse.StatusCode == 503 {
			var v ErrorV1
			err = a.client.decode(&v, localVarBody, localVarHttpResponse.Header.Get("Content-Type"));
				if err != nil {
					newErr.error = err.Error()
					return localVarReturnValue, localVarHttpResponse, newErr
				}
				newErr.model = v
				return localVarReturnValue, localVarHttpResponse, newErr
		}
		return localVarReturnValue, localVarHttpResponse, newErr
	}

	return localVarReturnValue, localVarHttpResponse, nil
}
/*
OauthUtilitiesReportingApiService Retrieve Oauth Consumer Funds
Retrieve users funds transactions 
 * @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param ids Reference identifiers for the top up transaction(s). Multiple identifiers in CSV format are accepted.
 * @param optional nil or *OauthUtilitiesReportingApiGetOauthConsumerFundsOpts - Optional Parameters:
     * @param "PaymentRef" (optional.String) -  Payment transaction identifier
@return InlineResponse20037
*/

type OauthUtilitiesReportingApiGetOauthConsumerFundsOpts struct {
    PaymentRef optional.String
}

func (a *OauthUtilitiesReportingApiService) GetOauthConsumerFunds(ctx context.Context, ids string, localVarOptionals *OauthUtilitiesReportingApiGetOauthConsumerFundsOpts) (InlineResponse20037, *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Get")
		localVarPostBody   interface{}
		localVarFileName   string
		localVarFileBytes  []byte
		localVarReturnValue InlineResponse20037
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath + "/oauth/consumer/funds"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	localVarQueryParams.Add("ids", parameterToString(ids, ""))
	if localVarOptionals != nil && localVarOptionals.PaymentRef.IsSet() {
		localVarQueryParams.Add("payment_ref", parameterToString(localVarOptionals.PaymentRef.Value(), ""))
	}
	// to determine the Content-Type header
	localVarHttpContentTypes := []string{}

	// set Content-Type header
	localVarHttpContentType := selectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHttpContentType
	}

	// to determine the Accept header
	localVarHttpHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHttpHeaderAccept := selectHeaderAccept(localVarHttpHeaderAccepts)
	if localVarHttpHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHttpHeaderAccept
	}
	r, err := a.client.prepareRequest(ctx, localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFileName, localVarFileBytes)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return localVarReturnValue, localVarHttpResponse, err
	}

	localVarBody, err := ioutil.ReadAll(localVarHttpResponse.Body)
	localVarHttpResponse.Body.Close()
	if err != nil {
		return localVarReturnValue, localVarHttpResponse, err
	}

	if localVarHttpResponse.StatusCode < 300 {
		// If we succeed, return the data, otherwise pass on to decode error.
		err = a.client.decode(&localVarReturnValue, localVarBody, localVarHttpResponse.Header.Get("Content-Type"));
		if err == nil { 
			return localVarReturnValue, localVarHttpResponse, err
		}
	}

	if localVarHttpResponse.StatusCode >= 300 {
		newErr := GenericSwaggerError{
			body: localVarBody,
			error: localVarHttpResponse.Status,
		}
		if localVarHttpResponse.StatusCode == 200 {
			var v InlineResponse20037
			err = a.client.decode(&v, localVarBody, localVarHttpResponse.Header.Get("Content-Type"));
				if err != nil {
					newErr.error = err.Error()
					return localVarReturnValue, localVarHttpResponse, newErr
				}
				newErr.model = v
				return localVarReturnValue, localVarHttpResponse, newErr
		}
		return localVarReturnValue, localVarHttpResponse, newErr
	}

	return localVarReturnValue, localVarHttpResponse, nil
}
/*
OauthUtilitiesReportingApiService Gets detailed information of a user
Gets detailed information of a user 
 * @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param optional nil or *OauthUtilitiesReportingApiGetOauthConsumerSearchUsersOpts - Optional Parameters:
     * @param "Email" (optional.String) -  Email Address (from users table)
     * @param "Mobile" (optional.String) -  Full mobile (concatenated value of country code and mobile number)
     * @param "UserHashId" (optional.String) -  User identifier
     * @param "UserStatus" (optional.String) -  User State
     * @param "KycStatus" (optional.String) -  User KYC State
     * @param "FirstName" (optional.String) -  first name to search
     * @param "LastName" (optional.String) -  last name to search
     * @param "CustomerId" (optional.String) -  Customer ID to search
     * @param "IdType" (optional.String) -  ID Type to search
     * @param "IdNumber" (optional.String) -  ID number to search
     * @param "IdDateIssue" (optional.String) -  ID issue date to search
     * @param "IdDateExpiry" (optional.String) -  ID expiry date to search
     * @param "Birthday" (optional.String) -  Birthday to search
     * @param "CardIdentifier" (optional.String) -  Card Identifier:   * &#x60;proxy&#x60; - Card Proxy number to search   * &#x60;mask&#x60; - Masked card number to search   * &#x60;full&#x60; - Full(actual) card number to search 
     * @param "CardIdentifierValue" (optional.String) -  Value is dependent on the &#x60;card_identifier&#x60; parameter
     * @param "BalanceCondition" (optional.String) -  Balances Condition:   * &#x60;gt&#x60; - Amount greater than or equal to specified balance   * &#x60;lt&#x60; - Amount lesser than or equal to specified balance   * &#x60;eq&#x60; - Amount equal to specified balance 
     * @param "CardAvailableBalance" (optional.Float32) -  Get all users where its per card available balance is within the specified &#x60;balance_condition&#x60; parameter
     * @param "CardWithholdingBalance" (optional.Float32) -  Get all users where its per card withholding balance is within the specified &#x60;balance_condition&#x60; parameter
     * @param "WalletAvailableBalance" (optional.Float32) -  Get all users where its per wallet available balance is within the specified &#x60;balance_condition&#x60; parameter
     * @param "WalletWithholdingBalance" (optional.Float32) -  Get all users where its per wallet withholding balance is within the specified &#x60;balance_condition&#x60; parameter
     * @param "TransactionReference" (optional.String) -  Get all users given the transaction reference (will apply to all transaction types)
     * @param "MatchType" (optional.String) -  Matching Condition:   * &#x60;fuzzy&#x60; - endpoint will perform a &#x60;LIKE&#x60; condition in query   * &#x60;exact&#x60; - endpoint will perform an &#x60;EQUAL&#x60; condition in query 
     * @param "ResponseType" (optional.String) -  Matching Condition:   * &#x60;expand&#x60; - response will show an expandable/complete information for address, documents, wallets object   * &#x60;condense&#x60; - response will show a url link on information for address, documents, wallets object 
     * @param "FilterCondition" (optional.String) -  Matching Condition:   * &#x60;AND&#x60; - for 2 or more filter parameters, the query will perform an &#x60;AND&#x60; operation in query    * &#x60;OR&#x60; - for 2 or more filter parameters, the query will perform an &#x60;OR&#x60; operation in query 
     * @param "Page" (optional.Float64) -  Page Number to display
     * @param "RecordsPerPage" (optional.Float64) -  Records per page to display
     * @param "PartnerId" (optional.String) -  Partner Id
@return UserReportV1
*/

type OauthUtilitiesReportingApiGetOauthConsumerSearchUsersOpts struct {
    Email optional.String
    Mobile optional.String
    UserHashId optional.String
    UserStatus optional.String
    KycStatus optional.String
    FirstName optional.String
    LastName optional.String
    CustomerId optional.String
    IdType optional.String
    IdNumber optional.String
    IdDateIssue optional.String
    IdDateExpiry optional.String
    Birthday optional.String
    CardIdentifier optional.String
    CardIdentifierValue optional.String
    BalanceCondition optional.String
    CardAvailableBalance optional.Float32
    CardWithholdingBalance optional.Float32
    WalletAvailableBalance optional.Float32
    WalletWithholdingBalance optional.Float32
    TransactionReference optional.String
    MatchType optional.String
    ResponseType optional.String
    FilterCondition optional.String
    Page optional.Float64
    RecordsPerPage optional.Float64
    PartnerId optional.String
}

func (a *OauthUtilitiesReportingApiService) GetOauthConsumerSearchUsers(ctx context.Context, localVarOptionals *OauthUtilitiesReportingApiGetOauthConsumerSearchUsersOpts) (UserReportV1, *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Get")
		localVarPostBody   interface{}
		localVarFileName   string
		localVarFileBytes  []byte
		localVarReturnValue UserReportV1
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath + "/oauth/consumer/search/users"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if localVarOptionals != nil && localVarOptionals.Email.IsSet() {
		localVarQueryParams.Add("email", parameterToString(localVarOptionals.Email.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.Mobile.IsSet() {
		localVarQueryParams.Add("mobile", parameterToString(localVarOptionals.Mobile.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.UserHashId.IsSet() {
		localVarQueryParams.Add("user_hash_id", parameterToString(localVarOptionals.UserHashId.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.UserStatus.IsSet() {
		localVarQueryParams.Add("user_status", parameterToString(localVarOptionals.UserStatus.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.KycStatus.IsSet() {
		localVarQueryParams.Add("kyc_status", parameterToString(localVarOptionals.KycStatus.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.FirstName.IsSet() {
		localVarQueryParams.Add("first_name", parameterToString(localVarOptionals.FirstName.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.LastName.IsSet() {
		localVarQueryParams.Add("last_name", parameterToString(localVarOptionals.LastName.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.CustomerId.IsSet() {
		localVarQueryParams.Add("customer_id", parameterToString(localVarOptionals.CustomerId.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.IdType.IsSet() {
		localVarQueryParams.Add("id_type", parameterToString(localVarOptionals.IdType.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.IdNumber.IsSet() {
		localVarQueryParams.Add("id_number", parameterToString(localVarOptionals.IdNumber.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.IdDateIssue.IsSet() {
		localVarQueryParams.Add("id_date_issue", parameterToString(localVarOptionals.IdDateIssue.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.IdDateExpiry.IsSet() {
		localVarQueryParams.Add("id_date_expiry", parameterToString(localVarOptionals.IdDateExpiry.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.Birthday.IsSet() {
		localVarQueryParams.Add("birthday", parameterToString(localVarOptionals.Birthday.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.CardIdentifier.IsSet() {
		localVarQueryParams.Add("card_identifier", parameterToString(localVarOptionals.CardIdentifier.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.CardIdentifierValue.IsSet() {
		localVarQueryParams.Add("card_identifier_value", parameterToString(localVarOptionals.CardIdentifierValue.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.BalanceCondition.IsSet() {
		localVarQueryParams.Add("balance_condition", parameterToString(localVarOptionals.BalanceCondition.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.CardAvailableBalance.IsSet() {
		localVarQueryParams.Add("card_available_balance", parameterToString(localVarOptionals.CardAvailableBalance.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.CardWithholdingBalance.IsSet() {
		localVarQueryParams.Add("card_withholding_balance", parameterToString(localVarOptionals.CardWithholdingBalance.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.WalletAvailableBalance.IsSet() {
		localVarQueryParams.Add("wallet_available_balance", parameterToString(localVarOptionals.WalletAvailableBalance.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.WalletWithholdingBalance.IsSet() {
		localVarQueryParams.Add("wallet_withholding_balance", parameterToString(localVarOptionals.WalletWithholdingBalance.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.TransactionReference.IsSet() {
		localVarQueryParams.Add("transaction_reference", parameterToString(localVarOptionals.TransactionReference.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.MatchType.IsSet() {
		localVarQueryParams.Add("match_type", parameterToString(localVarOptionals.MatchType.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.ResponseType.IsSet() {
		localVarQueryParams.Add("response_type", parameterToString(localVarOptionals.ResponseType.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.FilterCondition.IsSet() {
		localVarQueryParams.Add("filter_condition", parameterToString(localVarOptionals.FilterCondition.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.Page.IsSet() {
		localVarQueryParams.Add("page", parameterToString(localVarOptionals.Page.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.RecordsPerPage.IsSet() {
		localVarQueryParams.Add("records_per_page", parameterToString(localVarOptionals.RecordsPerPage.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.PartnerId.IsSet() {
		localVarQueryParams.Add("partner_id", parameterToString(localVarOptionals.PartnerId.Value(), ""))
	}
	// to determine the Content-Type header
	localVarHttpContentTypes := []string{}

	// set Content-Type header
	localVarHttpContentType := selectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHttpContentType
	}

	// to determine the Accept header
	localVarHttpHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHttpHeaderAccept := selectHeaderAccept(localVarHttpHeaderAccepts)
	if localVarHttpHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHttpHeaderAccept
	}
	r, err := a.client.prepareRequest(ctx, localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFileName, localVarFileBytes)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return localVarReturnValue, localVarHttpResponse, err
	}

	localVarBody, err := ioutil.ReadAll(localVarHttpResponse.Body)
	localVarHttpResponse.Body.Close()
	if err != nil {
		return localVarReturnValue, localVarHttpResponse, err
	}

	if localVarHttpResponse.StatusCode < 300 {
		// If we succeed, return the data, otherwise pass on to decode error.
		err = a.client.decode(&localVarReturnValue, localVarBody, localVarHttpResponse.Header.Get("Content-Type"));
		if err == nil { 
			return localVarReturnValue, localVarHttpResponse, err
		}
	}

	if localVarHttpResponse.StatusCode >= 300 {
		newErr := GenericSwaggerError{
			body: localVarBody,
			error: localVarHttpResponse.Status,
		}
		if localVarHttpResponse.StatusCode == 200 {
			var v UserReportV1
			err = a.client.decode(&v, localVarBody, localVarHttpResponse.Header.Get("Content-Type"));
				if err != nil {
					newErr.error = err.Error()
					return localVarReturnValue, localVarHttpResponse, newErr
				}
				newErr.model = v
				return localVarReturnValue, localVarHttpResponse, newErr
		}
		if localVarHttpResponse.StatusCode == 400 {
			var v ErrorV1
			err = a.client.decode(&v, localVarBody, localVarHttpResponse.Header.Get("Content-Type"));
				if err != nil {
					newErr.error = err.Error()
					return localVarReturnValue, localVarHttpResponse, newErr
				}
				newErr.model = v
				return localVarReturnValue, localVarHttpResponse, newErr
		}
		if localVarHttpResponse.StatusCode == 401 {
			var v ErrorV1
			err = a.client.decode(&v, localVarBody, localVarHttpResponse.Header.Get("Content-Type"));
				if err != nil {
					newErr.error = err.Error()
					return localVarReturnValue, localVarHttpResponse, newErr
				}
				newErr.model = v
				return localVarReturnValue, localVarHttpResponse, newErr
		}
		if localVarHttpResponse.StatusCode == 403 {
			var v ErrorV1
			err = a.client.decode(&v, localVarBody, localVarHttpResponse.Header.Get("Content-Type"));
				if err != nil {
					newErr.error = err.Error()
					return localVarReturnValue, localVarHttpResponse, newErr
				}
				newErr.model = v
				return localVarReturnValue, localVarHttpResponse, newErr
		}
		if localVarHttpResponse.StatusCode == 404 {
			var v ErrorV1
			err = a.client.decode(&v, localVarBody, localVarHttpResponse.Header.Get("Content-Type"));
				if err != nil {
					newErr.error = err.Error()
					return localVarReturnValue, localVarHttpResponse, newErr
				}
				newErr.model = v
				return localVarReturnValue, localVarHttpResponse, newErr
		}
		if localVarHttpResponse.StatusCode == 500 {
			var v ErrorV1
			err = a.client.decode(&v, localVarBody, localVarHttpResponse.Header.Get("Content-Type"));
				if err != nil {
					newErr.error = err.Error()
					return localVarReturnValue, localVarHttpResponse, newErr
				}
				newErr.model = v
				return localVarReturnValue, localVarHttpResponse, newErr
		}
		if localVarHttpResponse.StatusCode == 502 {
			var v ErrorV1
			err = a.client.decode(&v, localVarBody, localVarHttpResponse.Header.Get("Content-Type"));
				if err != nil {
					newErr.error = err.Error()
					return localVarReturnValue, localVarHttpResponse, newErr
				}
				newErr.model = v
				return localVarReturnValue, localVarHttpResponse, newErr
		}
		if localVarHttpResponse.StatusCode == 503 {
			var v ErrorV1
			err = a.client.decode(&v, localVarBody, localVarHttpResponse.Header.Get("Content-Type"));
				if err != nil {
					newErr.error = err.Error()
					return localVarReturnValue, localVarHttpResponse, newErr
				}
				newErr.model = v
				return localVarReturnValue, localVarHttpResponse, newErr
		}
		if localVarHttpResponse.StatusCode == 504 {
			var v ErrorV1
			err = a.client.decode(&v, localVarBody, localVarHttpResponse.Header.Get("Content-Type"));
				if err != nil {
					newErr.error = err.Error()
					return localVarReturnValue, localVarHttpResponse, newErr
				}
				newErr.model = v
				return localVarReturnValue, localVarHttpResponse, newErr
		}
		return localVarReturnValue, localVarHttpResponse, newErr
	}

	return localVarReturnValue, localVarHttpResponse, nil
}
/*
OauthUtilitiesReportingApiService Closed Loop Transactions
Gets detailed information of a closed-loop transactions
 * @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param optional nil or *OauthUtilitiesReportingApiGetOauthConsumerWalletFundsOpts - Optional Parameters:
     * @param "Type_" (optional.String) -  Acceptable Transaction Type:   * &#x60;Topup&#x60; - all top up transactions   * &#x60;Transfer&#x60; - all wallet to wallet money transfer transactions   * &#x60;Load&#x60; - all wallet to card transfer transactions   * &#x60;Unload&#x60; - all card to wallet transfer transactions   * &#x60;Reversal&#x60; - all  transactions   * &#x60;Refund&#x60; - all card to wallet transfer transactions   * &#x60;All&#x60; - all transactions above 
     * @param "UserHashId" (optional.String) - 
     * @param "Email" (optional.String) - 
     * @param "Mobile" (optional.String) - 
     * @param "CardIdentifier" (optional.String) -  Card Identifier:   * &#x60;proxy&#x60; - Card/Wallet Proxy number to search   * &#x60;mask&#x60; - Masked card number to search   * &#x60;full&#x60; - Full(actual) card number to search 
     * @param "CardIdentifierValue" (optional.String) -  Value is dependent on the &#x60;card_identifier&#x60; parameter; get all transactions done by the given card_identifier_value
     * @param "TransactionRefHash" (optional.String) -  Transaction Reference; get all transactions given the reference hash
     * @param "TransactionStatus" (optional.String) -  Transaction Status; get all transactions given the status:   * &#x60;hold&#x60; - all withholding transaction   * &#x60;active&#x60; - all success transaction   * &#x60;failed&#x60; - all failed transaction 
     * @param "AddedDate" (optional.String) -  Transaction Date Added (YYYY-MM-DD); get all transactions added on the given date
     * @param "UpdatedDate" (optional.String) -  Transaction Date Updated (YYYY-MM-DD); get all transactions updaed on the given date
     * @param "ReversalDate" (optional.String) -  Transaction Date Reversal (YYYY-MM-DD); get all transactions reversed on the given date
     * @param "RefundDate" (optional.String) -  Transaction Date Refund (YYYY-MM-DD); get all transactions refunded on the given date
@return ClosedLoopV1
*/

type OauthUtilitiesReportingApiGetOauthConsumerWalletFundsOpts struct {
    Type_ optional.String
    UserHashId optional.String
    Email optional.String
    Mobile optional.String
    CardIdentifier optional.String
    CardIdentifierValue optional.String
    TransactionRefHash optional.String
    TransactionStatus optional.String
    AddedDate optional.String
    UpdatedDate optional.String
    ReversalDate optional.String
    RefundDate optional.String
}

func (a *OauthUtilitiesReportingApiService) GetOauthConsumerWalletFunds(ctx context.Context, localVarOptionals *OauthUtilitiesReportingApiGetOauthConsumerWalletFundsOpts) (ClosedLoopV1, *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Get")
		localVarPostBody   interface{}
		localVarFileName   string
		localVarFileBytes  []byte
		localVarReturnValue ClosedLoopV1
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath + "/oauth/consumer/wallets/funds"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if localVarOptionals != nil && localVarOptionals.Type_.IsSet() {
		localVarQueryParams.Add("type", parameterToString(localVarOptionals.Type_.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.UserHashId.IsSet() {
		localVarQueryParams.Add("user_hash_id", parameterToString(localVarOptionals.UserHashId.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.Email.IsSet() {
		localVarQueryParams.Add("email", parameterToString(localVarOptionals.Email.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.Mobile.IsSet() {
		localVarQueryParams.Add("mobile", parameterToString(localVarOptionals.Mobile.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.CardIdentifier.IsSet() {
		localVarQueryParams.Add("card_identifier", parameterToString(localVarOptionals.CardIdentifier.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.CardIdentifierValue.IsSet() {
		localVarQueryParams.Add("card_identifier_value", parameterToString(localVarOptionals.CardIdentifierValue.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.TransactionRefHash.IsSet() {
		localVarQueryParams.Add("transaction_ref_hash", parameterToString(localVarOptionals.TransactionRefHash.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.TransactionStatus.IsSet() {
		localVarQueryParams.Add("transaction_status", parameterToString(localVarOptionals.TransactionStatus.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.AddedDate.IsSet() {
		localVarQueryParams.Add("added_date", parameterToString(localVarOptionals.AddedDate.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.UpdatedDate.IsSet() {
		localVarQueryParams.Add("updated_date", parameterToString(localVarOptionals.UpdatedDate.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.ReversalDate.IsSet() {
		localVarQueryParams.Add("reversal_date", parameterToString(localVarOptionals.ReversalDate.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.RefundDate.IsSet() {
		localVarQueryParams.Add("refund_date", parameterToString(localVarOptionals.RefundDate.Value(), ""))
	}
	// to determine the Content-Type header
	localVarHttpContentTypes := []string{}

	// set Content-Type header
	localVarHttpContentType := selectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHttpContentType
	}

	// to determine the Accept header
	localVarHttpHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHttpHeaderAccept := selectHeaderAccept(localVarHttpHeaderAccepts)
	if localVarHttpHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHttpHeaderAccept
	}
	r, err := a.client.prepareRequest(ctx, localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFileName, localVarFileBytes)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return localVarReturnValue, localVarHttpResponse, err
	}

	localVarBody, err := ioutil.ReadAll(localVarHttpResponse.Body)
	localVarHttpResponse.Body.Close()
	if err != nil {
		return localVarReturnValue, localVarHttpResponse, err
	}

	if localVarHttpResponse.StatusCode < 300 {
		// If we succeed, return the data, otherwise pass on to decode error.
		err = a.client.decode(&localVarReturnValue, localVarBody, localVarHttpResponse.Header.Get("Content-Type"));
		if err == nil { 
			return localVarReturnValue, localVarHttpResponse, err
		}
	}

	if localVarHttpResponse.StatusCode >= 300 {
		newErr := GenericSwaggerError{
			body: localVarBody,
			error: localVarHttpResponse.Status,
		}
		if localVarHttpResponse.StatusCode == 200 {
			var v ClosedLoopV1
			err = a.client.decode(&v, localVarBody, localVarHttpResponse.Header.Get("Content-Type"));
				if err != nil {
					newErr.error = err.Error()
					return localVarReturnValue, localVarHttpResponse, newErr
				}
				newErr.model = v
				return localVarReturnValue, localVarHttpResponse, newErr
		}
		if localVarHttpResponse.StatusCode == 400 {
			var v ErrorV1
			err = a.client.decode(&v, localVarBody, localVarHttpResponse.Header.Get("Content-Type"));
				if err != nil {
					newErr.error = err.Error()
					return localVarReturnValue, localVarHttpResponse, newErr
				}
				newErr.model = v
				return localVarReturnValue, localVarHttpResponse, newErr
		}
		if localVarHttpResponse.StatusCode == 403 {
			var v ErrorV1
			err = a.client.decode(&v, localVarBody, localVarHttpResponse.Header.Get("Content-Type"));
				if err != nil {
					newErr.error = err.Error()
					return localVarReturnValue, localVarHttpResponse, newErr
				}
				newErr.model = v
				return localVarReturnValue, localVarHttpResponse, newErr
		}
		if localVarHttpResponse.StatusCode == 404 {
			var v ErrorV1
			err = a.client.decode(&v, localVarBody, localVarHttpResponse.Header.Get("Content-Type"));
				if err != nil {
					newErr.error = err.Error()
					return localVarReturnValue, localVarHttpResponse, newErr
				}
				newErr.model = v
				return localVarReturnValue, localVarHttpResponse, newErr
		}
		if localVarHttpResponse.StatusCode == 500 {
			var v ErrorV1
			err = a.client.decode(&v, localVarBody, localVarHttpResponse.Header.Get("Content-Type"));
				if err != nil {
					newErr.error = err.Error()
					return localVarReturnValue, localVarHttpResponse, newErr
				}
				newErr.model = v
				return localVarReturnValue, localVarHttpResponse, newErr
		}
		if localVarHttpResponse.StatusCode == 503 {
			var v ErrorV1
			err = a.client.decode(&v, localVarBody, localVarHttpResponse.Header.Get("Content-Type"));
				if err != nil {
					newErr.error = err.Error()
					return localVarReturnValue, localVarHttpResponse, newErr
				}
				newErr.model = v
				return localVarReturnValue, localVarHttpResponse, newErr
		}
		return localVarReturnValue, localVarHttpResponse, newErr
	}

	return localVarReturnValue, localVarHttpResponse, nil
}
/*
OauthUtilitiesReportingApiService Get oauth wallet reversals
Gets Wallet reversals 
 * @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param limit
 * @param offset
@return InlineResponse20036
*/
func (a *OauthUtilitiesReportingApiService) GetOauthConsumerWalletReverseLimitOffset(ctx context.Context, limit string, offset string) (InlineResponse20036, *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Get")
		localVarPostBody   interface{}
		localVarFileName   string
		localVarFileBytes  []byte
		localVarReturnValue InlineResponse20036
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath + "/oauth/consumer/wallet/reverse/{limit}/{offset}"
	localVarPath = strings.Replace(localVarPath, "{"+"limit"+"}", fmt.Sprintf("%v", limit), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"offset"+"}", fmt.Sprintf("%v", offset), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	// to determine the Content-Type header
	localVarHttpContentTypes := []string{}

	// set Content-Type header
	localVarHttpContentType := selectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHttpContentType
	}

	// to determine the Accept header
	localVarHttpHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHttpHeaderAccept := selectHeaderAccept(localVarHttpHeaderAccepts)
	if localVarHttpHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHttpHeaderAccept
	}
	r, err := a.client.prepareRequest(ctx, localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFileName, localVarFileBytes)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return localVarReturnValue, localVarHttpResponse, err
	}

	localVarBody, err := ioutil.ReadAll(localVarHttpResponse.Body)
	localVarHttpResponse.Body.Close()
	if err != nil {
		return localVarReturnValue, localVarHttpResponse, err
	}

	if localVarHttpResponse.StatusCode < 300 {
		// If we succeed, return the data, otherwise pass on to decode error.
		err = a.client.decode(&localVarReturnValue, localVarBody, localVarHttpResponse.Header.Get("Content-Type"));
		if err == nil { 
			return localVarReturnValue, localVarHttpResponse, err
		}
	}

	if localVarHttpResponse.StatusCode >= 300 {
		newErr := GenericSwaggerError{
			body: localVarBody,
			error: localVarHttpResponse.Status,
		}
		if localVarHttpResponse.StatusCode == 200 {
			var v InlineResponse20036
			err = a.client.decode(&v, localVarBody, localVarHttpResponse.Header.Get("Content-Type"));
				if err != nil {
					newErr.error = err.Error()
					return localVarReturnValue, localVarHttpResponse, newErr
				}
				newErr.model = v
				return localVarReturnValue, localVarHttpResponse, newErr
		}
		return localVarReturnValue, localVarHttpResponse, newErr
	}

	return localVarReturnValue, localVarHttpResponse, nil
}
/*
OauthUtilitiesReportingApiService Get oauth user
Get oauth user 
 * @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
@return []UserFullV1
*/
func (a *OauthUtilitiesReportingApiService) GetOauthUser(ctx context.Context) ([]UserFullV1, *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Get")
		localVarPostBody   interface{}
		localVarFileName   string
		localVarFileBytes  []byte
		localVarReturnValue []UserFullV1
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath + "/oauth/user"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	// to determine the Content-Type header
	localVarHttpContentTypes := []string{}

	// set Content-Type header
	localVarHttpContentType := selectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHttpContentType
	}

	// to determine the Accept header
	localVarHttpHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHttpHeaderAccept := selectHeaderAccept(localVarHttpHeaderAccepts)
	if localVarHttpHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHttpHeaderAccept
	}
	r, err := a.client.prepareRequest(ctx, localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFileName, localVarFileBytes)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return localVarReturnValue, localVarHttpResponse, err
	}

	localVarBody, err := ioutil.ReadAll(localVarHttpResponse.Body)
	localVarHttpResponse.Body.Close()
	if err != nil {
		return localVarReturnValue, localVarHttpResponse, err
	}

	if localVarHttpResponse.StatusCode < 300 {
		// If we succeed, return the data, otherwise pass on to decode error.
		err = a.client.decode(&localVarReturnValue, localVarBody, localVarHttpResponse.Header.Get("Content-Type"));
		if err == nil { 
			return localVarReturnValue, localVarHttpResponse, err
		}
	}

	if localVarHttpResponse.StatusCode >= 300 {
		newErr := GenericSwaggerError{
			body: localVarBody,
			error: localVarHttpResponse.Status,
		}
		if localVarHttpResponse.StatusCode == 200 {
			var v []UserFullV1
			err = a.client.decode(&v, localVarBody, localVarHttpResponse.Header.Get("Content-Type"));
				if err != nil {
					newErr.error = err.Error()
					return localVarReturnValue, localVarHttpResponse, newErr
				}
				newErr.model = v
				return localVarReturnValue, localVarHttpResponse, newErr
		}
		if localVarHttpResponse.StatusCode == 400 {
			var v ErrorV1
			err = a.client.decode(&v, localVarBody, localVarHttpResponse.Header.Get("Content-Type"));
				if err != nil {
					newErr.error = err.Error()
					return localVarReturnValue, localVarHttpResponse, newErr
				}
				newErr.model = v
				return localVarReturnValue, localVarHttpResponse, newErr
		}
		if localVarHttpResponse.StatusCode == 403 {
			var v ErrorV1
			err = a.client.decode(&v, localVarBody, localVarHttpResponse.Header.Get("Content-Type"));
				if err != nil {
					newErr.error = err.Error()
					return localVarReturnValue, localVarHttpResponse, newErr
				}
				newErr.model = v
				return localVarReturnValue, localVarHttpResponse, newErr
		}
		if localVarHttpResponse.StatusCode == 500 {
			var v ErrorV1
			err = a.client.decode(&v, localVarBody, localVarHttpResponse.Header.Get("Content-Type"));
				if err != nil {
					newErr.error = err.Error()
					return localVarReturnValue, localVarHttpResponse, newErr
				}
				newErr.model = v
				return localVarReturnValue, localVarHttpResponse, newErr
		}
		if localVarHttpResponse.StatusCode == 503 {
			var v ErrorV1
			err = a.client.decode(&v, localVarBody, localVarHttpResponse.Header.Get("Content-Type"));
				if err != nil {
					newErr.error = err.Error()
					return localVarReturnValue, localVarHttpResponse, newErr
				}
				newErr.model = v
				return localVarReturnValue, localVarHttpResponse, newErr
		}
		return localVarReturnValue, localVarHttpResponse, newErr
	}

	return localVarReturnValue, localVarHttpResponse, nil
}
/*
OauthUtilitiesReportingApiService Get Oauth Users Search
Get Oauth Users Search 
 * @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
@return InlineResponse20040
*/
func (a *OauthUtilitiesReportingApiService) GetOauthUsersSearch(ctx context.Context) (InlineResponse20040, *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Get")
		localVarPostBody   interface{}
		localVarFileName   string
		localVarFileBytes  []byte
		localVarReturnValue InlineResponse20040
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath + "/oauth/users/search"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	// to determine the Content-Type header
	localVarHttpContentTypes := []string{}

	// set Content-Type header
	localVarHttpContentType := selectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHttpContentType
	}

	// to determine the Accept header
	localVarHttpHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHttpHeaderAccept := selectHeaderAccept(localVarHttpHeaderAccepts)
	if localVarHttpHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHttpHeaderAccept
	}
	r, err := a.client.prepareRequest(ctx, localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFileName, localVarFileBytes)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return localVarReturnValue, localVarHttpResponse, err
	}

	localVarBody, err := ioutil.ReadAll(localVarHttpResponse.Body)
	localVarHttpResponse.Body.Close()
	if err != nil {
		return localVarReturnValue, localVarHttpResponse, err
	}

	if localVarHttpResponse.StatusCode < 300 {
		// If we succeed, return the data, otherwise pass on to decode error.
		err = a.client.decode(&localVarReturnValue, localVarBody, localVarHttpResponse.Header.Get("Content-Type"));
		if err == nil { 
			return localVarReturnValue, localVarHttpResponse, err
		}
	}

	if localVarHttpResponse.StatusCode >= 300 {
		newErr := GenericSwaggerError{
			body: localVarBody,
			error: localVarHttpResponse.Status,
		}
		if localVarHttpResponse.StatusCode == 200 {
			var v InlineResponse20040
			err = a.client.decode(&v, localVarBody, localVarHttpResponse.Header.Get("Content-Type"));
				if err != nil {
					newErr.error = err.Error()
					return localVarReturnValue, localVarHttpResponse, newErr
				}
				newErr.model = v
				return localVarReturnValue, localVarHttpResponse, newErr
		}
		if localVarHttpResponse.StatusCode == 400 {
			var v ErrorV1
			err = a.client.decode(&v, localVarBody, localVarHttpResponse.Header.Get("Content-Type"));
				if err != nil {
					newErr.error = err.Error()
					return localVarReturnValue, localVarHttpResponse, newErr
				}
				newErr.model = v
				return localVarReturnValue, localVarHttpResponse, newErr
		}
		if localVarHttpResponse.StatusCode == 403 {
			var v ErrorV1
			err = a.client.decode(&v, localVarBody, localVarHttpResponse.Header.Get("Content-Type"));
				if err != nil {
					newErr.error = err.Error()
					return localVarReturnValue, localVarHttpResponse, newErr
				}
				newErr.model = v
				return localVarReturnValue, localVarHttpResponse, newErr
		}
		if localVarHttpResponse.StatusCode == 500 {
			var v ErrorV1
			err = a.client.decode(&v, localVarBody, localVarHttpResponse.Header.Get("Content-Type"));
				if err != nil {
					newErr.error = err.Error()
					return localVarReturnValue, localVarHttpResponse, newErr
				}
				newErr.model = v
				return localVarReturnValue, localVarHttpResponse, newErr
		}
		if localVarHttpResponse.StatusCode == 503 {
			var v ErrorV1
			err = a.client.decode(&v, localVarBody, localVarHttpResponse.Header.Get("Content-Type"));
				if err != nil {
					newErr.error = err.Error()
					return localVarReturnValue, localVarHttpResponse, newErr
				}
				newErr.model = v
				return localVarReturnValue, localVarHttpResponse, newErr
		}
		return localVarReturnValue, localVarHttpResponse, newErr
	}

	return localVarReturnValue, localVarHttpResponse, nil
}
/*
OauthUtilitiesReportingApiService Get all fund transfer transactions
Get all fund transfer processed under the program
 * @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param optional nil or *OauthUtilitiesReportingApiGetProgramBankAccountDebitCreditTransactionsOpts - Optional Parameters:
     * @param "UserHashId" (optional.String) -  af11e6342686a34d37a58e7130ec5e11
     * @param "WalletHashId" (optional.String) -  df11e6342686a34d37a58e7130ec5e11
     * @param "BankAccountNumber" (optional.String) -  123456
     * @param "BankAccountId" (optional.String) -  df11e6342686a34d37a58e7130ec5e11
     * @param "UniquePaymentId" (optional.String) -  123456
     * @param "Status" (optional.String) -  active | inactive | terminated
     * @param "AccountType" (optional.String) -  consumer or corporate
     * @param "FromDate" (optional.String) -  From Date
     * @param "ToDate" (optional.String) -  To Date
     * @param "Page" (optional.Float64) -  The page number
     * @param "RecordsPerPage" (optional.Float64) -  The number of records to return per page
@return InlineResponse20049
*/

type OauthUtilitiesReportingApiGetProgramBankAccountDebitCreditTransactionsOpts struct {
    UserHashId optional.String
    WalletHashId optional.String
    BankAccountNumber optional.String
    BankAccountId optional.String
    UniquePaymentId optional.String
    Status optional.String
    AccountType optional.String
    FromDate optional.String
    ToDate optional.String
    Page optional.Float64
    RecordsPerPage optional.Float64
}

func (a *OauthUtilitiesReportingApiService) GetProgramBankAccountDebitCreditTransactions(ctx context.Context, localVarOptionals *OauthUtilitiesReportingApiGetProgramBankAccountDebitCreditTransactionsOpts) (InlineResponse20049, *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Get")
		localVarPostBody   interface{}
		localVarFileName   string
		localVarFileBytes  []byte
		localVarReturnValue InlineResponse20049
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath + "/oauth/consumer/wallets/fund_transfers"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if localVarOptionals != nil && localVarOptionals.UserHashId.IsSet() {
		localVarQueryParams.Add("user_hash_id", parameterToString(localVarOptionals.UserHashId.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.WalletHashId.IsSet() {
		localVarQueryParams.Add("wallet_hash_id", parameterToString(localVarOptionals.WalletHashId.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.BankAccountNumber.IsSet() {
		localVarQueryParams.Add("bank_account_number", parameterToString(localVarOptionals.BankAccountNumber.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.BankAccountId.IsSet() {
		localVarQueryParams.Add("bank_account_id", parameterToString(localVarOptionals.BankAccountId.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.UniquePaymentId.IsSet() {
		localVarQueryParams.Add("unique_payment_id", parameterToString(localVarOptionals.UniquePaymentId.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.Status.IsSet() {
		localVarQueryParams.Add("status", parameterToString(localVarOptionals.Status.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.AccountType.IsSet() {
		localVarQueryParams.Add("account_type", parameterToString(localVarOptionals.AccountType.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.FromDate.IsSet() {
		localVarQueryParams.Add("from_date", parameterToString(localVarOptionals.FromDate.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.ToDate.IsSet() {
		localVarQueryParams.Add("to_date", parameterToString(localVarOptionals.ToDate.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.Page.IsSet() {
		localVarQueryParams.Add("page", parameterToString(localVarOptionals.Page.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.RecordsPerPage.IsSet() {
		localVarQueryParams.Add("records_per_page", parameterToString(localVarOptionals.RecordsPerPage.Value(), ""))
	}
	// to determine the Content-Type header
	localVarHttpContentTypes := []string{}

	// set Content-Type header
	localVarHttpContentType := selectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHttpContentType
	}

	// to determine the Accept header
	localVarHttpHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHttpHeaderAccept := selectHeaderAccept(localVarHttpHeaderAccepts)
	if localVarHttpHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHttpHeaderAccept
	}
	r, err := a.client.prepareRequest(ctx, localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFileName, localVarFileBytes)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return localVarReturnValue, localVarHttpResponse, err
	}

	localVarBody, err := ioutil.ReadAll(localVarHttpResponse.Body)
	localVarHttpResponse.Body.Close()
	if err != nil {
		return localVarReturnValue, localVarHttpResponse, err
	}

	if localVarHttpResponse.StatusCode < 300 {
		// If we succeed, return the data, otherwise pass on to decode error.
		err = a.client.decode(&localVarReturnValue, localVarBody, localVarHttpResponse.Header.Get("Content-Type"));
		if err == nil { 
			return localVarReturnValue, localVarHttpResponse, err
		}
	}

	if localVarHttpResponse.StatusCode >= 300 {
		newErr := GenericSwaggerError{
			body: localVarBody,
			error: localVarHttpResponse.Status,
		}
		if localVarHttpResponse.StatusCode == 200 {
			var v InlineResponse20049
			err = a.client.decode(&v, localVarBody, localVarHttpResponse.Header.Get("Content-Type"));
				if err != nil {
					newErr.error = err.Error()
					return localVarReturnValue, localVarHttpResponse, newErr
				}
				newErr.model = v
				return localVarReturnValue, localVarHttpResponse, newErr
		}
		if localVarHttpResponse.StatusCode == 400 {
			var v ErrorV1
			err = a.client.decode(&v, localVarBody, localVarHttpResponse.Header.Get("Content-Type"));
				if err != nil {
					newErr.error = err.Error()
					return localVarReturnValue, localVarHttpResponse, newErr
				}
				newErr.model = v
				return localVarReturnValue, localVarHttpResponse, newErr
		}
		if localVarHttpResponse.StatusCode == 401 {
			var v ErrorV1
			err = a.client.decode(&v, localVarBody, localVarHttpResponse.Header.Get("Content-Type"));
				if err != nil {
					newErr.error = err.Error()
					return localVarReturnValue, localVarHttpResponse, newErr
				}
				newErr.model = v
				return localVarReturnValue, localVarHttpResponse, newErr
		}
		if localVarHttpResponse.StatusCode == 403 {
			var v ErrorV1
			err = a.client.decode(&v, localVarBody, localVarHttpResponse.Header.Get("Content-Type"));
				if err != nil {
					newErr.error = err.Error()
					return localVarReturnValue, localVarHttpResponse, newErr
				}
				newErr.model = v
				return localVarReturnValue, localVarHttpResponse, newErr
		}
		if localVarHttpResponse.StatusCode == 404 {
			var v ErrorV1
			err = a.client.decode(&v, localVarBody, localVarHttpResponse.Header.Get("Content-Type"));
				if err != nil {
					newErr.error = err.Error()
					return localVarReturnValue, localVarHttpResponse, newErr
				}
				newErr.model = v
				return localVarReturnValue, localVarHttpResponse, newErr
		}
		if localVarHttpResponse.StatusCode == 500 {
			var v ErrorV1
			err = a.client.decode(&v, localVarBody, localVarHttpResponse.Header.Get("Content-Type"));
				if err != nil {
					newErr.error = err.Error()
					return localVarReturnValue, localVarHttpResponse, newErr
				}
				newErr.model = v
				return localVarReturnValue, localVarHttpResponse, newErr
		}
		if localVarHttpResponse.StatusCode == 502 {
			var v ErrorV1
			err = a.client.decode(&v, localVarBody, localVarHttpResponse.Header.Get("Content-Type"));
				if err != nil {
					newErr.error = err.Error()
					return localVarReturnValue, localVarHttpResponse, newErr
				}
				newErr.model = v
				return localVarReturnValue, localVarHttpResponse, newErr
		}
		if localVarHttpResponse.StatusCode == 503 {
			var v ErrorV1
			err = a.client.decode(&v, localVarBody, localVarHttpResponse.Header.Get("Content-Type"));
				if err != nil {
					newErr.error = err.Error()
					return localVarReturnValue, localVarHttpResponse, newErr
				}
				newErr.model = v
				return localVarReturnValue, localVarHttpResponse, newErr
		}
		if localVarHttpResponse.StatusCode == 504 {
			var v ErrorV1
			err = a.client.decode(&v, localVarBody, localVarHttpResponse.Header.Get("Content-Type"));
				if err != nil {
					newErr.error = err.Error()
					return localVarReturnValue, localVarHttpResponse, newErr
				}
				newErr.model = v
				return localVarReturnValue, localVarHttpResponse, newErr
		}
		return localVarReturnValue, localVarHttpResponse, newErr
	}

	return localVarReturnValue, localVarHttpResponse, nil
}
/*
OauthUtilitiesReportingApiService Get all Standing Instructions
Get all standing instructions registered under the program
 * @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param optional nil or *OauthUtilitiesReportingApiGetProgramBankAccountStandingInstructionOpts - Optional Parameters:
     * @param "UserHashId" (optional.String) -  af11e6342686a34d37a58e7130ec5e11
     * @param "WalletHashId" (optional.String) -  df11e6342686a34d37a58e7130ec5e11
     * @param "BankAccountNumber" (optional.String) -  123456
     * @param "BankAccountId" (optional.String) -  df11e6342686a34d37a58e7130ec5e11
     * @param "Status" (optional.String) -  pending | under-review | approved | expired | revoked | disabled
     * @param "AccountType" (optional.String) -  consumer or corporate
     * @param "FromDate" (optional.String) -  From Date
     * @param "ToDate" (optional.String) -  To Date
     * @param "Page" (optional.Float64) -  Page Number
     * @param "RecordsPerPage" (optional.Float64) -  Number of records per page
@return InlineResponse20048
*/

type OauthUtilitiesReportingApiGetProgramBankAccountStandingInstructionOpts struct {
    UserHashId optional.String
    WalletHashId optional.String
    BankAccountNumber optional.String
    BankAccountId optional.String
    Status optional.String
    AccountType optional.String
    FromDate optional.String
    ToDate optional.String
    Page optional.Float64
    RecordsPerPage optional.Float64
}

func (a *OauthUtilitiesReportingApiService) GetProgramBankAccountStandingInstruction(ctx context.Context, localVarOptionals *OauthUtilitiesReportingApiGetProgramBankAccountStandingInstructionOpts) (InlineResponse20048, *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Get")
		localVarPostBody   interface{}
		localVarFileName   string
		localVarFileBytes  []byte
		localVarReturnValue InlineResponse20048
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath + "/oauth/consumer/wallets/standing_instructions"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if localVarOptionals != nil && localVarOptionals.UserHashId.IsSet() {
		localVarQueryParams.Add("user_hash_id", parameterToString(localVarOptionals.UserHashId.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.WalletHashId.IsSet() {
		localVarQueryParams.Add("wallet_hash_id", parameterToString(localVarOptionals.WalletHashId.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.BankAccountNumber.IsSet() {
		localVarQueryParams.Add("bank_account_number", parameterToString(localVarOptionals.BankAccountNumber.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.BankAccountId.IsSet() {
		localVarQueryParams.Add("bank_account_id", parameterToString(localVarOptionals.BankAccountId.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.Status.IsSet() {
		localVarQueryParams.Add("status", parameterToString(localVarOptionals.Status.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.AccountType.IsSet() {
		localVarQueryParams.Add("account_type", parameterToString(localVarOptionals.AccountType.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.FromDate.IsSet() {
		localVarQueryParams.Add("from_date", parameterToString(localVarOptionals.FromDate.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.ToDate.IsSet() {
		localVarQueryParams.Add("to_date", parameterToString(localVarOptionals.ToDate.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.Page.IsSet() {
		localVarQueryParams.Add("page", parameterToString(localVarOptionals.Page.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.RecordsPerPage.IsSet() {
		localVarQueryParams.Add("records_per_page", parameterToString(localVarOptionals.RecordsPerPage.Value(), ""))
	}
	// to determine the Content-Type header
	localVarHttpContentTypes := []string{}

	// set Content-Type header
	localVarHttpContentType := selectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHttpContentType
	}

	// to determine the Accept header
	localVarHttpHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHttpHeaderAccept := selectHeaderAccept(localVarHttpHeaderAccepts)
	if localVarHttpHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHttpHeaderAccept
	}
	r, err := a.client.prepareRequest(ctx, localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFileName, localVarFileBytes)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return localVarReturnValue, localVarHttpResponse, err
	}

	localVarBody, err := ioutil.ReadAll(localVarHttpResponse.Body)
	localVarHttpResponse.Body.Close()
	if err != nil {
		return localVarReturnValue, localVarHttpResponse, err
	}

	if localVarHttpResponse.StatusCode < 300 {
		// If we succeed, return the data, otherwise pass on to decode error.
		err = a.client.decode(&localVarReturnValue, localVarBody, localVarHttpResponse.Header.Get("Content-Type"));
		if err == nil { 
			return localVarReturnValue, localVarHttpResponse, err
		}
	}

	if localVarHttpResponse.StatusCode >= 300 {
		newErr := GenericSwaggerError{
			body: localVarBody,
			error: localVarHttpResponse.Status,
		}
		if localVarHttpResponse.StatusCode == 200 {
			var v InlineResponse20048
			err = a.client.decode(&v, localVarBody, localVarHttpResponse.Header.Get("Content-Type"));
				if err != nil {
					newErr.error = err.Error()
					return localVarReturnValue, localVarHttpResponse, newErr
				}
				newErr.model = v
				return localVarReturnValue, localVarHttpResponse, newErr
		}
		if localVarHttpResponse.StatusCode == 400 {
			var v ErrorV1
			err = a.client.decode(&v, localVarBody, localVarHttpResponse.Header.Get("Content-Type"));
				if err != nil {
					newErr.error = err.Error()
					return localVarReturnValue, localVarHttpResponse, newErr
				}
				newErr.model = v
				return localVarReturnValue, localVarHttpResponse, newErr
		}
		if localVarHttpResponse.StatusCode == 401 {
			var v ErrorV1
			err = a.client.decode(&v, localVarBody, localVarHttpResponse.Header.Get("Content-Type"));
				if err != nil {
					newErr.error = err.Error()
					return localVarReturnValue, localVarHttpResponse, newErr
				}
				newErr.model = v
				return localVarReturnValue, localVarHttpResponse, newErr
		}
		if localVarHttpResponse.StatusCode == 403 {
			var v ErrorV1
			err = a.client.decode(&v, localVarBody, localVarHttpResponse.Header.Get("Content-Type"));
				if err != nil {
					newErr.error = err.Error()
					return localVarReturnValue, localVarHttpResponse, newErr
				}
				newErr.model = v
				return localVarReturnValue, localVarHttpResponse, newErr
		}
		if localVarHttpResponse.StatusCode == 404 {
			var v ErrorV1
			err = a.client.decode(&v, localVarBody, localVarHttpResponse.Header.Get("Content-Type"));
				if err != nil {
					newErr.error = err.Error()
					return localVarReturnValue, localVarHttpResponse, newErr
				}
				newErr.model = v
				return localVarReturnValue, localVarHttpResponse, newErr
		}
		if localVarHttpResponse.StatusCode == 500 {
			var v ErrorV1
			err = a.client.decode(&v, localVarBody, localVarHttpResponse.Header.Get("Content-Type"));
				if err != nil {
					newErr.error = err.Error()
					return localVarReturnValue, localVarHttpResponse, newErr
				}
				newErr.model = v
				return localVarReturnValue, localVarHttpResponse, newErr
		}
		if localVarHttpResponse.StatusCode == 502 {
			var v ErrorV1
			err = a.client.decode(&v, localVarBody, localVarHttpResponse.Header.Get("Content-Type"));
				if err != nil {
					newErr.error = err.Error()
					return localVarReturnValue, localVarHttpResponse, newErr
				}
				newErr.model = v
				return localVarReturnValue, localVarHttpResponse, newErr
		}
		if localVarHttpResponse.StatusCode == 503 {
			var v ErrorV1
			err = a.client.decode(&v, localVarBody, localVarHttpResponse.Header.Get("Content-Type"));
				if err != nil {
					newErr.error = err.Error()
					return localVarReturnValue, localVarHttpResponse, newErr
				}
				newErr.model = v
				return localVarReturnValue, localVarHttpResponse, newErr
		}
		if localVarHttpResponse.StatusCode == 504 {
			var v ErrorV1
			err = a.client.decode(&v, localVarBody, localVarHttpResponse.Header.Get("Content-Type"));
				if err != nil {
					newErr.error = err.Error()
					return localVarReturnValue, localVarHttpResponse, newErr
				}
				newErr.model = v
				return localVarReturnValue, localVarHttpResponse, newErr
		}
		return localVarReturnValue, localVarHttpResponse, newErr
	}

	return localVarReturnValue, localVarHttpResponse, nil
}
/*
OauthUtilitiesReportingApiService Get Program Bank Transfer Transactions
Retrieve list of all Bank Transfer transactions (Top  up and Transfer out) done in the system 
 * @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
@return InlineResponse20044
*/
func (a *OauthUtilitiesReportingApiService) GetProgramBankTransferTransactions(ctx context.Context) (InlineResponse20044, *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Get")
		localVarPostBody   interface{}
		localVarFileName   string
		localVarFileBytes  []byte
		localVarReturnValue InlineResponse20044
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath + "/oauth/consumer/account/transactions"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	// to determine the Content-Type header
	localVarHttpContentTypes := []string{}

	// set Content-Type header
	localVarHttpContentType := selectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHttpContentType
	}

	// to determine the Accept header
	localVarHttpHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHttpHeaderAccept := selectHeaderAccept(localVarHttpHeaderAccepts)
	if localVarHttpHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHttpHeaderAccept
	}
	r, err := a.client.prepareRequest(ctx, localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFileName, localVarFileBytes)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return localVarReturnValue, localVarHttpResponse, err
	}

	localVarBody, err := ioutil.ReadAll(localVarHttpResponse.Body)
	localVarHttpResponse.Body.Close()
	if err != nil {
		return localVarReturnValue, localVarHttpResponse, err
	}

	if localVarHttpResponse.StatusCode < 300 {
		// If we succeed, return the data, otherwise pass on to decode error.
		err = a.client.decode(&localVarReturnValue, localVarBody, localVarHttpResponse.Header.Get("Content-Type"));
		if err == nil { 
			return localVarReturnValue, localVarHttpResponse, err
		}
	}

	if localVarHttpResponse.StatusCode >= 300 {
		newErr := GenericSwaggerError{
			body: localVarBody,
			error: localVarHttpResponse.Status,
		}
		if localVarHttpResponse.StatusCode == 200 {
			var v InlineResponse20044
			err = a.client.decode(&v, localVarBody, localVarHttpResponse.Header.Get("Content-Type"));
				if err != nil {
					newErr.error = err.Error()
					return localVarReturnValue, localVarHttpResponse, newErr
				}
				newErr.model = v
				return localVarReturnValue, localVarHttpResponse, newErr
		}
		return localVarReturnValue, localVarHttpResponse, newErr
	}

	return localVarReturnValue, localVarHttpResponse, nil
}
/*
OauthUtilitiesReportingApiService Get all Bank Accounts
Get all Bank Accounts registered under the program
 * @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param optional nil or *OauthUtilitiesReportingApiGetProgramBeneficiaryBankAccountOpts - Optional Parameters:
     * @param "UserHashId" (optional.String) -  af11e6342686a34d37a58e7130ec5e11
     * @param "WalletHashId" (optional.String) -  df11e6342686a34d37a58e7130ec5e11
     * @param "BankAccountNumber" (optional.String) -  123456
     * @param "BankAccountId" (optional.String) -  adf11e6342686a34d37a58e7130ec5e11
     * @param "AccountType" (optional.String) -  consumer or corporate
     * @param "FromDate" (optional.String) -  From Date
     * @param "ToDate" (optional.String) -  To Date
     * @param "Page" (optional.Float64) -  Page Number
     * @param "RecordsPerPage" (optional.Float64) -  Records per page
@return InlineResponse20047
*/

type OauthUtilitiesReportingApiGetProgramBeneficiaryBankAccountOpts struct {
    UserHashId optional.String
    WalletHashId optional.String
    BankAccountNumber optional.String
    BankAccountId optional.String
    AccountType optional.String
    FromDate optional.String
    ToDate optional.String
    Page optional.Float64
    RecordsPerPage optional.Float64
}

func (a *OauthUtilitiesReportingApiService) GetProgramBeneficiaryBankAccount(ctx context.Context, localVarOptionals *OauthUtilitiesReportingApiGetProgramBeneficiaryBankAccountOpts) (InlineResponse20047, *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Get")
		localVarPostBody   interface{}
		localVarFileName   string
		localVarFileBytes  []byte
		localVarReturnValue InlineResponse20047
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath + "/oauth/consumer/wallets/bank_accounts"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if localVarOptionals != nil && localVarOptionals.UserHashId.IsSet() {
		localVarQueryParams.Add("user_hash_id", parameterToString(localVarOptionals.UserHashId.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.WalletHashId.IsSet() {
		localVarQueryParams.Add("wallet_hash_id", parameterToString(localVarOptionals.WalletHashId.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.BankAccountNumber.IsSet() {
		localVarQueryParams.Add("bank_account_number", parameterToString(localVarOptionals.BankAccountNumber.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.BankAccountId.IsSet() {
		localVarQueryParams.Add("bank_account_id", parameterToString(localVarOptionals.BankAccountId.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.AccountType.IsSet() {
		localVarQueryParams.Add("account_type", parameterToString(localVarOptionals.AccountType.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.FromDate.IsSet() {
		localVarQueryParams.Add("from_date", parameterToString(localVarOptionals.FromDate.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.ToDate.IsSet() {
		localVarQueryParams.Add("to_date", parameterToString(localVarOptionals.ToDate.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.Page.IsSet() {
		localVarQueryParams.Add("page", parameterToString(localVarOptionals.Page.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.RecordsPerPage.IsSet() {
		localVarQueryParams.Add("records_per_page", parameterToString(localVarOptionals.RecordsPerPage.Value(), ""))
	}
	// to determine the Content-Type header
	localVarHttpContentTypes := []string{}

	// set Content-Type header
	localVarHttpContentType := selectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHttpContentType
	}

	// to determine the Accept header
	localVarHttpHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHttpHeaderAccept := selectHeaderAccept(localVarHttpHeaderAccepts)
	if localVarHttpHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHttpHeaderAccept
	}
	r, err := a.client.prepareRequest(ctx, localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFileName, localVarFileBytes)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return localVarReturnValue, localVarHttpResponse, err
	}

	localVarBody, err := ioutil.ReadAll(localVarHttpResponse.Body)
	localVarHttpResponse.Body.Close()
	if err != nil {
		return localVarReturnValue, localVarHttpResponse, err
	}

	if localVarHttpResponse.StatusCode < 300 {
		// If we succeed, return the data, otherwise pass on to decode error.
		err = a.client.decode(&localVarReturnValue, localVarBody, localVarHttpResponse.Header.Get("Content-Type"));
		if err == nil { 
			return localVarReturnValue, localVarHttpResponse, err
		}
	}

	if localVarHttpResponse.StatusCode >= 300 {
		newErr := GenericSwaggerError{
			body: localVarBody,
			error: localVarHttpResponse.Status,
		}
		if localVarHttpResponse.StatusCode == 200 {
			var v InlineResponse20047
			err = a.client.decode(&v, localVarBody, localVarHttpResponse.Header.Get("Content-Type"));
				if err != nil {
					newErr.error = err.Error()
					return localVarReturnValue, localVarHttpResponse, newErr
				}
				newErr.model = v
				return localVarReturnValue, localVarHttpResponse, newErr
		}
		if localVarHttpResponse.StatusCode == 400 {
			var v ErrorV1
			err = a.client.decode(&v, localVarBody, localVarHttpResponse.Header.Get("Content-Type"));
				if err != nil {
					newErr.error = err.Error()
					return localVarReturnValue, localVarHttpResponse, newErr
				}
				newErr.model = v
				return localVarReturnValue, localVarHttpResponse, newErr
		}
		if localVarHttpResponse.StatusCode == 401 {
			var v ErrorV1
			err = a.client.decode(&v, localVarBody, localVarHttpResponse.Header.Get("Content-Type"));
				if err != nil {
					newErr.error = err.Error()
					return localVarReturnValue, localVarHttpResponse, newErr
				}
				newErr.model = v
				return localVarReturnValue, localVarHttpResponse, newErr
		}
		if localVarHttpResponse.StatusCode == 403 {
			var v ErrorV1
			err = a.client.decode(&v, localVarBody, localVarHttpResponse.Header.Get("Content-Type"));
				if err != nil {
					newErr.error = err.Error()
					return localVarReturnValue, localVarHttpResponse, newErr
				}
				newErr.model = v
				return localVarReturnValue, localVarHttpResponse, newErr
		}
		if localVarHttpResponse.StatusCode == 404 {
			var v ErrorV1
			err = a.client.decode(&v, localVarBody, localVarHttpResponse.Header.Get("Content-Type"));
				if err != nil {
					newErr.error = err.Error()
					return localVarReturnValue, localVarHttpResponse, newErr
				}
				newErr.model = v
				return localVarReturnValue, localVarHttpResponse, newErr
		}
		if localVarHttpResponse.StatusCode == 500 {
			var v ErrorV1
			err = a.client.decode(&v, localVarBody, localVarHttpResponse.Header.Get("Content-Type"));
				if err != nil {
					newErr.error = err.Error()
					return localVarReturnValue, localVarHttpResponse, newErr
				}
				newErr.model = v
				return localVarReturnValue, localVarHttpResponse, newErr
		}
		if localVarHttpResponse.StatusCode == 502 {
			var v ErrorV1
			err = a.client.decode(&v, localVarBody, localVarHttpResponse.Header.Get("Content-Type"));
				if err != nil {
					newErr.error = err.Error()
					return localVarReturnValue, localVarHttpResponse, newErr
				}
				newErr.model = v
				return localVarReturnValue, localVarHttpResponse, newErr
		}
		if localVarHttpResponse.StatusCode == 503 {
			var v ErrorV1
			err = a.client.decode(&v, localVarBody, localVarHttpResponse.Header.Get("Content-Type"));
				if err != nil {
					newErr.error = err.Error()
					return localVarReturnValue, localVarHttpResponse, newErr
				}
				newErr.model = v
				return localVarReturnValue, localVarHttpResponse, newErr
		}
		if localVarHttpResponse.StatusCode == 504 {
			var v ErrorV1
			err = a.client.decode(&v, localVarBody, localVarHttpResponse.Header.Get("Content-Type"));
				if err != nil {
					newErr.error = err.Error()
					return localVarReturnValue, localVarHttpResponse, newErr
				}
				newErr.model = v
				return localVarReturnValue, localVarHttpResponse, newErr
		}
		return localVarReturnValue, localVarHttpResponse, newErr
	}

	return localVarReturnValue, localVarHttpResponse, nil
}
/*
OauthUtilitiesReportingApiService Get all Virtual Accounts
Get all Virtual Accounts registered under the program
 * @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param optional nil or *OauthUtilitiesReportingApiGetProgramVirtualAccountsOpts - Optional Parameters:
     * @param "UserHashId" (optional.String) -  af11e6342686a34d37a58e7130ec5e11
     * @param "WalletHashId" (optional.String) -  df11e6342686a34d37a58e7130ec5e11
     * @param "AccountNumber" (optional.String) -  123456
     * @param "Status" (optional.String) -  active | inactive | terminated
     * @param "AccountType" (optional.String) -  consumer or corporate
     * @param "FromDate" (optional.String) -  From Date
     * @param "ToDate" (optional.String) -  To Date
     * @param "Page" (optional.Float64) -  Page Number
     * @param "RecordsPerPage" (optional.Float64) -  Number of records per page
@return InlineResponse20043
*/

type OauthUtilitiesReportingApiGetProgramVirtualAccountsOpts struct {
    UserHashId optional.String
    WalletHashId optional.String
    AccountNumber optional.String
    Status optional.String
    AccountType optional.String
    FromDate optional.String
    ToDate optional.String
    Page optional.Float64
    RecordsPerPage optional.Float64
}

func (a *OauthUtilitiesReportingApiService) GetProgramVirtualAccounts(ctx context.Context, localVarOptionals *OauthUtilitiesReportingApiGetProgramVirtualAccountsOpts) (InlineResponse20043, *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Get")
		localVarPostBody   interface{}
		localVarFileName   string
		localVarFileBytes  []byte
		localVarReturnValue InlineResponse20043
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath + "/oauth/consumer/wallets/virtual_accounts"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if localVarOptionals != nil && localVarOptionals.UserHashId.IsSet() {
		localVarQueryParams.Add("user_hash_id", parameterToString(localVarOptionals.UserHashId.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.WalletHashId.IsSet() {
		localVarQueryParams.Add("wallet_hash_id", parameterToString(localVarOptionals.WalletHashId.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.AccountNumber.IsSet() {
		localVarQueryParams.Add("account_number", parameterToString(localVarOptionals.AccountNumber.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.Status.IsSet() {
		localVarQueryParams.Add("status", parameterToString(localVarOptionals.Status.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.AccountType.IsSet() {
		localVarQueryParams.Add("account_type", parameterToString(localVarOptionals.AccountType.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.FromDate.IsSet() {
		localVarQueryParams.Add("from_date", parameterToString(localVarOptionals.FromDate.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.ToDate.IsSet() {
		localVarQueryParams.Add("to_date", parameterToString(localVarOptionals.ToDate.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.Page.IsSet() {
		localVarQueryParams.Add("page", parameterToString(localVarOptionals.Page.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.RecordsPerPage.IsSet() {
		localVarQueryParams.Add("records_per_page", parameterToString(localVarOptionals.RecordsPerPage.Value(), ""))
	}
	// to determine the Content-Type header
	localVarHttpContentTypes := []string{}

	// set Content-Type header
	localVarHttpContentType := selectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHttpContentType
	}

	// to determine the Accept header
	localVarHttpHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHttpHeaderAccept := selectHeaderAccept(localVarHttpHeaderAccepts)
	if localVarHttpHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHttpHeaderAccept
	}
	r, err := a.client.prepareRequest(ctx, localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFileName, localVarFileBytes)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return localVarReturnValue, localVarHttpResponse, err
	}

	localVarBody, err := ioutil.ReadAll(localVarHttpResponse.Body)
	localVarHttpResponse.Body.Close()
	if err != nil {
		return localVarReturnValue, localVarHttpResponse, err
	}

	if localVarHttpResponse.StatusCode < 300 {
		// If we succeed, return the data, otherwise pass on to decode error.
		err = a.client.decode(&localVarReturnValue, localVarBody, localVarHttpResponse.Header.Get("Content-Type"));
		if err == nil { 
			return localVarReturnValue, localVarHttpResponse, err
		}
	}

	if localVarHttpResponse.StatusCode >= 300 {
		newErr := GenericSwaggerError{
			body: localVarBody,
			error: localVarHttpResponse.Status,
		}
		if localVarHttpResponse.StatusCode == 200 {
			var v InlineResponse20043
			err = a.client.decode(&v, localVarBody, localVarHttpResponse.Header.Get("Content-Type"));
				if err != nil {
					newErr.error = err.Error()
					return localVarReturnValue, localVarHttpResponse, newErr
				}
				newErr.model = v
				return localVarReturnValue, localVarHttpResponse, newErr
		}
		if localVarHttpResponse.StatusCode == 400 {
			var v ErrorV1
			err = a.client.decode(&v, localVarBody, localVarHttpResponse.Header.Get("Content-Type"));
				if err != nil {
					newErr.error = err.Error()
					return localVarReturnValue, localVarHttpResponse, newErr
				}
				newErr.model = v
				return localVarReturnValue, localVarHttpResponse, newErr
		}
		if localVarHttpResponse.StatusCode == 401 {
			var v ErrorV1
			err = a.client.decode(&v, localVarBody, localVarHttpResponse.Header.Get("Content-Type"));
				if err != nil {
					newErr.error = err.Error()
					return localVarReturnValue, localVarHttpResponse, newErr
				}
				newErr.model = v
				return localVarReturnValue, localVarHttpResponse, newErr
		}
		if localVarHttpResponse.StatusCode == 403 {
			var v ErrorV1
			err = a.client.decode(&v, localVarBody, localVarHttpResponse.Header.Get("Content-Type"));
				if err != nil {
					newErr.error = err.Error()
					return localVarReturnValue, localVarHttpResponse, newErr
				}
				newErr.model = v
				return localVarReturnValue, localVarHttpResponse, newErr
		}
		if localVarHttpResponse.StatusCode == 404 {
			var v ErrorV1
			err = a.client.decode(&v, localVarBody, localVarHttpResponse.Header.Get("Content-Type"));
				if err != nil {
					newErr.error = err.Error()
					return localVarReturnValue, localVarHttpResponse, newErr
				}
				newErr.model = v
				return localVarReturnValue, localVarHttpResponse, newErr
		}
		if localVarHttpResponse.StatusCode == 500 {
			var v ErrorV1
			err = a.client.decode(&v, localVarBody, localVarHttpResponse.Header.Get("Content-Type"));
				if err != nil {
					newErr.error = err.Error()
					return localVarReturnValue, localVarHttpResponse, newErr
				}
				newErr.model = v
				return localVarReturnValue, localVarHttpResponse, newErr
		}
		if localVarHttpResponse.StatusCode == 502 {
			var v ErrorV1
			err = a.client.decode(&v, localVarBody, localVarHttpResponse.Header.Get("Content-Type"));
				if err != nil {
					newErr.error = err.Error()
					return localVarReturnValue, localVarHttpResponse, newErr
				}
				newErr.model = v
				return localVarReturnValue, localVarHttpResponse, newErr
		}
		if localVarHttpResponse.StatusCode == 503 {
			var v ErrorV1
			err = a.client.decode(&v, localVarBody, localVarHttpResponse.Header.Get("Content-Type"));
				if err != nil {
					newErr.error = err.Error()
					return localVarReturnValue, localVarHttpResponse, newErr
				}
				newErr.model = v
				return localVarReturnValue, localVarHttpResponse, newErr
		}
		if localVarHttpResponse.StatusCode == 504 {
			var v ErrorV1
			err = a.client.decode(&v, localVarBody, localVarHttpResponse.Header.Get("Content-Type"));
				if err != nil {
					newErr.error = err.Error()
					return localVarReturnValue, localVarHttpResponse, newErr
				}
				newErr.model = v
				return localVarReturnValue, localVarHttpResponse, newErr
		}
		return localVarReturnValue, localVarHttpResponse, newErr
	}

	return localVarReturnValue, localVarHttpResponse, nil
}
