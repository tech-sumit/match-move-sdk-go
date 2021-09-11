
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

type RemittancesApiService service
/*
RemittancesApiService Get Form for Remittance Attachments
Retrieve the attachment form details and supported document types to attach
 * @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param payoutAgent Payout agent hash id
@return RemittanceAttachmentFormV1
*/
func (a *RemittancesApiService) GetOauthConsumersFundsTransfersOverseasAttachmentForm(ctx context.Context, payoutAgent string) (RemittanceAttachmentFormV1, *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Get")
		localVarPostBody   interface{}
		localVarFileName   string
		localVarFileBytes  []byte
		localVarReturnValue RemittanceAttachmentFormV1
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath + "/oauth/consumers/funds/transfers/overseas/attachment/form"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	localVarQueryParams.Add("payout_agent", parameterToString(payoutAgent, ""))
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
			var v RemittanceAttachmentFormV1
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
RemittancesApiService Remittance Overseas Beneficiary Form
Retrieve beneficiary form of a given provider and payout agent
 * @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param optional nil or *RemittancesApiGetOauthConsumersFundsTransfersOverseasBeneficiaryFormOpts - Optional Parameters:
     * @param "Provider" (optional.String) -  Provider code from &#x60;GET /users/wallets/funds/transfers/overseas/types&#x60;
     * @param "Type_" (optional.String) -  Beneficiary Type
     * @param "PayoutAgent" (optional.String) -  Payout Agent
@return InlineResponse20055
*/

type RemittancesApiGetOauthConsumersFundsTransfersOverseasBeneficiaryFormOpts struct {
    Provider optional.String
    Type_ optional.String
    PayoutAgent optional.String
}

func (a *RemittancesApiService) GetOauthConsumersFundsTransfersOverseasBeneficiaryForm(ctx context.Context, localVarOptionals *RemittancesApiGetOauthConsumersFundsTransfersOverseasBeneficiaryFormOpts) (InlineResponse20055, *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Get")
		localVarPostBody   interface{}
		localVarFileName   string
		localVarFileBytes  []byte
		localVarReturnValue InlineResponse20055
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath + "/oauth/consumer/funds/transfers/overseas/beneficiary/form"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if localVarOptionals != nil && localVarOptionals.Provider.IsSet() {
		localVarQueryParams.Add("provider", parameterToString(localVarOptionals.Provider.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.Type_.IsSet() {
		localVarQueryParams.Add("type", parameterToString(localVarOptionals.Type_.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.PayoutAgent.IsSet() {
		localVarQueryParams.Add("payout_agent", parameterToString(localVarOptionals.PayoutAgent.Value(), ""))
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
			var v InlineResponse20055
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
RemittancesApiService Remittance Countries
Retrieve all countries supported by the program for remittance
 * @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param optional nil or *RemittancesApiGetOauthConsumersFundsTransfersOverseasCountriesOpts - Optional Parameters:
     * @param "Provider" (optional.String) -  Provider code from &#x60;GET /users/wallets/funds/transfers/overseas/types&#x60;
     * @param "Type_" (optional.String) -  Beneficiary Type
     * @param "Name" (optional.String) -  Country Name
     * @param "Code" (optional.String) -  Country Code
@return RemittanceCountriesV1
*/

type RemittancesApiGetOauthConsumersFundsTransfersOverseasCountriesOpts struct {
    Provider optional.String
    Type_ optional.String
    Name optional.String
    Code optional.String
}

func (a *RemittancesApiService) GetOauthConsumersFundsTransfersOverseasCountries(ctx context.Context, localVarOptionals *RemittancesApiGetOauthConsumersFundsTransfersOverseasCountriesOpts) (RemittanceCountriesV1, *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Get")
		localVarPostBody   interface{}
		localVarFileName   string
		localVarFileBytes  []byte
		localVarReturnValue RemittanceCountriesV1
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath + "/oauth/consumer/funds/transfers/overseas/countries"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if localVarOptionals != nil && localVarOptionals.Provider.IsSet() {
		localVarQueryParams.Add("provider", parameterToString(localVarOptionals.Provider.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.Type_.IsSet() {
		localVarQueryParams.Add("type", parameterToString(localVarOptionals.Type_.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.Name.IsSet() {
		localVarQueryParams.Add("name", parameterToString(localVarOptionals.Name.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.Code.IsSet() {
		localVarQueryParams.Add("code", parameterToString(localVarOptionals.Code.Value(), ""))
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
			var v RemittanceCountriesV1
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
RemittancesApiService Remittance Payment Modes
Retrieve all modes with specified countries per mode supported by the program for remittance
 * @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param optional nil or *RemittancesApiGetOauthConsumersFundsTransfersOverseasModesOpts - Optional Parameters:
     * @param "Provider" (optional.String) -  Provider code from &#x60;GET /users/wallets/funds/transfers/overseas/types&#x60;
     * @param "Type_" (optional.String) -  Beneficiary Type
     * @param "Name" (optional.String) -  Payment channel name
     * @param "Code" (optional.String) -  Payment channel code
     * @param "CountryCode" (optional.String) -  Country Code
@return RemittancePaymentModesV1
*/

type RemittancesApiGetOauthConsumersFundsTransfersOverseasModesOpts struct {
    Provider optional.String
    Type_ optional.String
    Name optional.String
    Code optional.String
    CountryCode optional.String
}

func (a *RemittancesApiService) GetOauthConsumersFundsTransfersOverseasModes(ctx context.Context, localVarOptionals *RemittancesApiGetOauthConsumersFundsTransfersOverseasModesOpts) (RemittancePaymentModesV1, *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Get")
		localVarPostBody   interface{}
		localVarFileName   string
		localVarFileBytes  []byte
		localVarReturnValue RemittancePaymentModesV1
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath + "/oauth/consumer/funds/transfers/overseas/modes"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if localVarOptionals != nil && localVarOptionals.Provider.IsSet() {
		localVarQueryParams.Add("provider", parameterToString(localVarOptionals.Provider.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.Type_.IsSet() {
		localVarQueryParams.Add("type", parameterToString(localVarOptionals.Type_.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.Name.IsSet() {
		localVarQueryParams.Add("name", parameterToString(localVarOptionals.Name.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.Code.IsSet() {
		localVarQueryParams.Add("code", parameterToString(localVarOptionals.Code.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.CountryCode.IsSet() {
		localVarQueryParams.Add("country_code", parameterToString(localVarOptionals.CountryCode.Value(), ""))
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
			var v RemittancePaymentModesV1
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
RemittancesApiService Get Remittance History
Get Remittance history against the wallet of the user 
 * @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param limit
 * @param offset
 * @param optional nil or *RemittancesApiGetUsersWalletsFundsTransfersOverseasHistoryLimitOffsetOpts - Optional Parameters:
     * @param "Sort" (optional.String) -  Comma delimited sorting of the result according to id and date.added
     * @param "DeliveryMethod" (optional.String) -  Filter transactions based on the delivery method
     * @param "Provider" (optional.String) -  Filter transactions based on the [provider](#get-users-wallets-funds-transfers-overseas-types)
     * @param "DateAdded" (optional.String) -  Filter transactions based on the date added
     * @param "DateExpiry" (optional.String) -  Filter transactions based on the date expiry
     * @param "Status" (optional.String) -  Filter transactions based status
     * @param "Search" (optional.String) -  Filter transactions based on the transaction code , reference id, or transfer reference id
     * @param "SenderFirstName" (optional.String) -  Filter transactions based on the sender first name
     * @param "SenderMiddleName" (optional.String) -  Filter transactions based on the sender middle name
     * @param "SenderLastName" (optional.String) -  Filter transactions based on the sender last name
     * @param "SenderMobileNumber" (optional.String) -  Filter transactions based on the sender mobile number
     * @param "ReceiverFirstName" (optional.String) -  Filter transactions based on the beneficiary first name
     * @param "ReceiverMiddleName" (optional.String) -  Filter transactions based on the beneficiary middle name
     * @param "ReceiverLastName" (optional.String) -  Filter transactions based on the beneficiary last name
     * @param "BusinessRegisteredName" (optional.String) -  Filter transactions based on the corporate registered name
     * @param "BeneficiaryHashId" (optional.String) -  Filter transactions based on the corporate/consumer beneficiary hash id
     * @param "ClientRefId" (optional.String) -  Filter transactions based on the client reference id
@return InlineResponse20050
*/

type RemittancesApiGetUsersWalletsFundsTransfersOverseasHistoryLimitOffsetOpts struct {
    Sort optional.String
    DeliveryMethod optional.String
    Provider optional.String
    DateAdded optional.String
    DateExpiry optional.String
    Status optional.String
    Search optional.String
    SenderFirstName optional.String
    SenderMiddleName optional.String
    SenderLastName optional.String
    SenderMobileNumber optional.String
    ReceiverFirstName optional.String
    ReceiverMiddleName optional.String
    ReceiverLastName optional.String
    BusinessRegisteredName optional.String
    BeneficiaryHashId optional.String
    ClientRefId optional.String
}

func (a *RemittancesApiService) GetUsersWalletsFundsTransfersOverseasHistoryLimitOffset(ctx context.Context, limit string, offset string, localVarOptionals *RemittancesApiGetUsersWalletsFundsTransfersOverseasHistoryLimitOffsetOpts) (InlineResponse20050, *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Get")
		localVarPostBody   interface{}
		localVarFileName   string
		localVarFileBytes  []byte
		localVarReturnValue InlineResponse20050
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath + "/users/wallets/funds/transfers/overseas/history/{limit}]/{offset}"
	localVarPath = strings.Replace(localVarPath, "{"+"limit"+"}", fmt.Sprintf("%v", limit), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"offset"+"}", fmt.Sprintf("%v", offset), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if localVarOptionals != nil && localVarOptionals.Sort.IsSet() {
		localVarQueryParams.Add("sort", parameterToString(localVarOptionals.Sort.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.DeliveryMethod.IsSet() {
		localVarQueryParams.Add("delivery_method", parameterToString(localVarOptionals.DeliveryMethod.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.Provider.IsSet() {
		localVarQueryParams.Add("provider", parameterToString(localVarOptionals.Provider.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.DateAdded.IsSet() {
		localVarQueryParams.Add("date_added", parameterToString(localVarOptionals.DateAdded.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.DateExpiry.IsSet() {
		localVarQueryParams.Add("date_expiry", parameterToString(localVarOptionals.DateExpiry.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.Status.IsSet() {
		localVarQueryParams.Add("status", parameterToString(localVarOptionals.Status.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.Search.IsSet() {
		localVarQueryParams.Add("search", parameterToString(localVarOptionals.Search.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.SenderFirstName.IsSet() {
		localVarQueryParams.Add("sender_first_name", parameterToString(localVarOptionals.SenderFirstName.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.SenderMiddleName.IsSet() {
		localVarQueryParams.Add("sender_middle_name", parameterToString(localVarOptionals.SenderMiddleName.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.SenderLastName.IsSet() {
		localVarQueryParams.Add("sender_last_name", parameterToString(localVarOptionals.SenderLastName.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.SenderMobileNumber.IsSet() {
		localVarQueryParams.Add("sender_mobile_number", parameterToString(localVarOptionals.SenderMobileNumber.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.ReceiverFirstName.IsSet() {
		localVarQueryParams.Add("receiver_first_name", parameterToString(localVarOptionals.ReceiverFirstName.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.ReceiverMiddleName.IsSet() {
		localVarQueryParams.Add("receiver_middle_name", parameterToString(localVarOptionals.ReceiverMiddleName.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.ReceiverLastName.IsSet() {
		localVarQueryParams.Add("receiver_last_name", parameterToString(localVarOptionals.ReceiverLastName.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.BusinessRegisteredName.IsSet() {
		localVarQueryParams.Add("business_registered_name", parameterToString(localVarOptionals.BusinessRegisteredName.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.BeneficiaryHashId.IsSet() {
		localVarQueryParams.Add("beneficiary_hash_id", parameterToString(localVarOptionals.BeneficiaryHashId.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.ClientRefId.IsSet() {
		localVarQueryParams.Add("client_ref_id", parameterToString(localVarOptionals.ClientRefId.Value(), ""))
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
			var v InlineResponse20050
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
RemittancesApiService Get Uploaded Remittance Attachments
Retrieve the attached documents against the given remittance hash id
 * @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param id Remittance Hash ID
@return RemittanceAttachmentsV1
*/
func (a *RemittancesApiService) GetUsersWalletsFundsTransfersOverseasIdAttachments(ctx context.Context, id string) (RemittanceAttachmentsV1, *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Get")
		localVarPostBody   interface{}
		localVarFileName   string
		localVarFileBytes  []byte
		localVarReturnValue RemittanceAttachmentsV1
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath + "/users/wallets/funds/transfers/overseas/{id}/attachments"
	localVarPath = strings.Replace(localVarPath, "{"+"id"+"}", fmt.Sprintf("%v", id), -1)

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
			var v RemittanceAttachmentsV1
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
RemittancesApiService Remittance Provider Agents
Retrieve all the payout agents supported on the system for remittance.  Data retrieved for this endpoint will be getting from a cache object which is populated on a background task. 
 * @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param limit Pagination limit
 * @param offset Pagination offset
 * @param optional nil or *RemittancesApiGetUsersWalletsFundsTransfersOverseasProvidersOpts - Optional Parameters:
     * @param "Type_" (optional.String) -  Beneficiary type
     * @param "AgentId" (optional.String) -  Payout Agent ID
     * @param "ReceiveCountry" (optional.String) -  Receiving Country Code (ISO alpha-3)
     * @param "Mode" (optional.String) -  Mode from response of GET oauth/consumer/funds/transfers/overseas/modes
     * @param "Provider" (optional.String) -  Remittance Rails
@return InlineResponse20054
*/

type RemittancesApiGetUsersWalletsFundsTransfersOverseasProvidersOpts struct {
    Type_ optional.String
    AgentId optional.String
    ReceiveCountry optional.String
    Mode optional.String
    Provider optional.String
}

func (a *RemittancesApiService) GetUsersWalletsFundsTransfersOverseasProviders(ctx context.Context, limit int32, offset int32, localVarOptionals *RemittancesApiGetUsersWalletsFundsTransfersOverseasProvidersOpts) (InlineResponse20054, *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Get")
		localVarPostBody   interface{}
		localVarFileName   string
		localVarFileBytes  []byte
		localVarReturnValue InlineResponse20054
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath + "/users/wallets/funds/transfers/overseas/providers/{limit}]/{offset}"
	localVarPath = strings.Replace(localVarPath, "{"+"limit"+"}", fmt.Sprintf("%v", limit), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"offset"+"}", fmt.Sprintf("%v", offset), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if localVarOptionals != nil && localVarOptionals.Type_.IsSet() {
		localVarQueryParams.Add("type", parameterToString(localVarOptionals.Type_.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.AgentId.IsSet() {
		localVarQueryParams.Add("agent_id", parameterToString(localVarOptionals.AgentId.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.ReceiveCountry.IsSet() {
		localVarQueryParams.Add("receive_country", parameterToString(localVarOptionals.ReceiveCountry.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.Mode.IsSet() {
		localVarQueryParams.Add("mode", parameterToString(localVarOptionals.Mode.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.Provider.IsSet() {
		localVarQueryParams.Add("provider", parameterToString(localVarOptionals.Provider.Value(), ""))
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
			var v InlineResponse20054
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
RemittancesApiService Get Remittance Fees
Retrieve list of fees for remittance agents configured in the system.
 * @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param type_
 * @param limit
 * @param offset
@return RemittanceFeesV1
*/
func (a *RemittancesApiService) GetUsersWalletsOverseasTransfersFees(ctx context.Context, type_ string, limit string, offset string) (RemittanceFeesV1, *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Get")
		localVarPostBody   interface{}
		localVarFileName   string
		localVarFileBytes  []byte
		localVarReturnValue RemittanceFeesV1
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath + "/users/wallets/funds/transfers/overseas/{type}/fees/{limit}/{offset}"
	localVarPath = strings.Replace(localVarPath, "{"+"type"+"}", fmt.Sprintf("%v", type_), -1)
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
			var v RemittanceFeesV1
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
RemittancesApiService Get Remittance Providers
Lists all overseas providers offered. You may retrieve information for a specific provider by appending its code. This resource serves as a prerequisite to list rates, payers and destinations. 
 * @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param code
@return RemittanceProvidersV1
*/
func (a *RemittancesApiService) GetUsersWalletsOverseasTransfersTypes(ctx context.Context, code string) (RemittanceProvidersV1, *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Get")
		localVarPostBody   interface{}
		localVarFileName   string
		localVarFileBytes  []byte
		localVarReturnValue RemittanceProvidersV1
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath + "/users/wallets/funds/transfers/overseas/types/{code}"
	localVarPath = strings.Replace(localVarPath, "{"+"code"+"}", fmt.Sprintf("%v", code), -1)

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
			var v RemittanceProvidersV1
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
RemittancesApiService Confirm Remittance transaction
Confirm a remittance transaction , given the id(s)
 * @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param ids
@return InlineResponse20051
*/
func (a *RemittancesApiService) PostOauthConsumerFundsTransfersOverseas(ctx context.Context, ids string) (InlineResponse20051, *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Post")
		localVarPostBody   interface{}
		localVarFileName   string
		localVarFileBytes  []byte
		localVarReturnValue InlineResponse20051
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath + "/oauth/consumer/funds/transfers/overseas"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	// to determine the Content-Type header
	localVarHttpContentTypes := []string{"application/x-www-form-urlencoded"}

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
	localVarFormParams.Add("ids", parameterToString(ids, ""))
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
			var v InlineResponse20051
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
RemittancesApiService Send Overseas Transfers without providing Rails
Send a money transfer request with the specified agent.  Any details required on this endpoint will be based from the sender/beneficiary form on the calculate (/test) endpoint.  Any transaction made through this resource will be on pending validation. Confirmation must be made to clear the transaction. 
 * @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param optional nil or *RemittancesApiPostUsersWalletsFundsTransfersOverseasOpts - Optional Parameters:
     * @param "AgentId" (optional.String) - 
     * @param "CalculationMode" (optional.String) - 
     * @param "TransactionType" (optional.String) - 
     * @param "Amount" (optional.String) - 
     * @param "SendCurrency" (optional.String) - 
     * @param "ReceiveCurrency" (optional.String) - 
     * @param "ReceiveCountry" (optional.String) - 
     * @param "PaymentMode" (optional.String) - 
     * @param "RoutingType" (optional.String) - 
     * @param "RoutingParam" (optional.String) - 
     * @param "QuotationId" (optional.String) - 
     * @param "PayoutAgent" (optional.String) - 
     * @param "HashId" (optional.String) - 
     * @param "FirstName" (optional.String) - 
     * @param "MiddleName" (optional.String) - 
     * @param "LastName" (optional.String) - 
     * @param "ChineseName" (optional.String) - 
     * @param "ReceiveMobileNumber" (optional.String) - 
     * @param "ReceiveGender" (optional.String) - 
     * @param "IdNumber" (optional.String) - 
     * @param "IdType" (optional.String) - 
     * @param "Occupation" (optional.String) - 
     * @param "Address1" (optional.String) - 
     * @param "Address2" (optional.String) - 
     * @param "BirthDate" (optional.String) - 
     * @param "City" (optional.String) - 
     * @param "State" (optional.String) - 
     * @param "Country" (optional.String) - 
     * @param "Nationality" (optional.String) - 
     * @param "Zipcode" (optional.String) - 
     * @param "BeneficiaryHashId" (optional.String) - 
     * @param "ReceivingBusinessRegisteredName" (optional.String) - 
     * @param "ReceivingBusinessTradingName" (optional.String) - 
     * @param "ReceivingBusinessEmail" (optional.String) - 
     * @param "ReceivingBusinessMsisdn" (optional.String) - 
     * @param "ReceivingBusinessRegistrationNumber" (optional.String) - 
     * @param "ReceivingBusinessTaxId" (optional.String) - 
     * @param "ReceivingBusinessAddress1" (optional.String) - 
     * @param "ReceivingBusinessAddress2" (optional.String) - 
     * @param "ReceivingBusinessCity" (optional.String) - 
     * @param "ReceivingBusinessState" (optional.String) - 
     * @param "ReceivingBusinessCountry" (optional.String) - 
     * @param "ReceivingBusinessRepresentativeFirstName" (optional.String) - 
     * @param "ReceivingBusinessRepresentativeMiddleName" (optional.String) - 
     * @param "ReceivingBusinessRepresentativeLastName" (optional.String) - 
     * @param "ReceivingBusinessRepresentativeLastName2" (optional.String) - 
     * @param "ReceivingBusinessRepresentativeNativeName" (optional.String) - 
     * @param "ReceivingBusinessIdType" (optional.String) - 
     * @param "ReceivingBusinessIdNumber" (optional.String) - 
     * @param "ReceivingBusinessCountryOfIssue" (optional.String) - 
     * @param "ReceivingBusinessIssueDate" (optional.String) - 
     * @param "ReceivingBusinessExpireDate" (optional.String) - 
     * @param "ReceivingBusinessDateOfIncorporation" (optional.String) - 
     * @param "SendingBusinessHashId" (optional.String) - 
     * @param "SendingBusinessRegisteredName" (optional.String) - 
     * @param "SendingBusinessTradingName" (optional.String) - 
     * @param "SendingBusinessEmail" (optional.String) - 
     * @param "SendingBusinessMsisdn" (optional.String) - 
     * @param "SendingBusinessRegistrationNumber" (optional.String) - 
     * @param "SendingBusinessTaxId" (optional.String) - 
     * @param "SendingBusinessAddress1" (optional.String) - 
     * @param "SendingBusinessAddress2" (optional.String) - 
     * @param "SendingBusinessCity" (optional.String) - 
     * @param "SendingBusinessState" (optional.String) - 
     * @param "SendingBusinessZipcode" (optional.String) - 
     * @param "SendingBusinessCountry" (optional.String) - 
     * @param "SendingBusinessRepresentativeFirstName" (optional.String) - 
     * @param "SendingBusinessRepresentativeMiddleName" (optional.String) - 
     * @param "SendingBusinessRepresentativeLastName" (optional.String) - 
     * @param "SendingBusinessRepresentativeLastName2" (optional.String) - 
     * @param "SendingBusinessRepresentativeNativeName" (optional.String) - 
     * @param "SendingBusinessIdType" (optional.String) - 
     * @param "SendingBusinessIdNumber" (optional.String) - 
     * @param "SendingBusinessCountryOfIssue" (optional.String) - 
     * @param "SendingBusinessIssueDate" (optional.String) - 
     * @param "SendingBusinessExpireDate" (optional.String) - 
     * @param "SendingBusinessDateOfIncorporation" (optional.String) - 
     * @param "DocumentReferenceNumber" (optional.String) - 
     * @param "SwiftCode" (optional.String) - 
     * @param "BankIfcCode" (optional.String) - 
     * @param "BankBranchName" (optional.String) - 
     * @param "BankBranchCode" (optional.String) - 
     * @param "BankAccountNumber" (optional.String) - 
     * @param "CbcBankName" (optional.String) - 
     * @param "BpiBankName" (optional.String) - 
     * @param "PartnerName" (optional.String) - 
     * @param "SourceOfIncome" (optional.String) - 
     * @param "Relationship" (optional.String) - 
     * @param "Reason" (optional.String) - 
     * @param "SenderReference" (optional.String) - 
     * @param "AdditionalInformation" (optional.String) - 
@return RemittanceSendV1
*/

type RemittancesApiPostUsersWalletsFundsTransfersOverseasOpts struct {
    AgentId optional.String
    CalculationMode optional.String
    TransactionType optional.String
    Amount optional.String
    SendCurrency optional.String
    ReceiveCurrency optional.String
    ReceiveCountry optional.String
    PaymentMode optional.String
    RoutingType optional.String
    RoutingParam optional.String
    QuotationId optional.String
    PayoutAgent optional.String
    HashId optional.String
    FirstName optional.String
    MiddleName optional.String
    LastName optional.String
    ChineseName optional.String
    ReceiveMobileNumber optional.String
    ReceiveGender optional.String
    IdNumber optional.String
    IdType optional.String
    Occupation optional.String
    Address1 optional.String
    Address2 optional.String
    BirthDate optional.String
    City optional.String
    State optional.String
    Country optional.String
    Nationality optional.String
    Zipcode optional.String
    BeneficiaryHashId optional.String
    ReceivingBusinessRegisteredName optional.String
    ReceivingBusinessTradingName optional.String
    ReceivingBusinessEmail optional.String
    ReceivingBusinessMsisdn optional.String
    ReceivingBusinessRegistrationNumber optional.String
    ReceivingBusinessTaxId optional.String
    ReceivingBusinessAddress1 optional.String
    ReceivingBusinessAddress2 optional.String
    ReceivingBusinessCity optional.String
    ReceivingBusinessState optional.String
    ReceivingBusinessCountry optional.String
    ReceivingBusinessRepresentativeFirstName optional.String
    ReceivingBusinessRepresentativeMiddleName optional.String
    ReceivingBusinessRepresentativeLastName optional.String
    ReceivingBusinessRepresentativeLastName2 optional.String
    ReceivingBusinessRepresentativeNativeName optional.String
    ReceivingBusinessIdType optional.String
    ReceivingBusinessIdNumber optional.String
    ReceivingBusinessCountryOfIssue optional.String
    ReceivingBusinessIssueDate optional.String
    ReceivingBusinessExpireDate optional.String
    ReceivingBusinessDateOfIncorporation optional.String
    SendingBusinessHashId optional.String
    SendingBusinessRegisteredName optional.String
    SendingBusinessTradingName optional.String
    SendingBusinessEmail optional.String
    SendingBusinessMsisdn optional.String
    SendingBusinessRegistrationNumber optional.String
    SendingBusinessTaxId optional.String
    SendingBusinessAddress1 optional.String
    SendingBusinessAddress2 optional.String
    SendingBusinessCity optional.String
    SendingBusinessState optional.String
    SendingBusinessZipcode optional.String
    SendingBusinessCountry optional.String
    SendingBusinessRepresentativeFirstName optional.String
    SendingBusinessRepresentativeMiddleName optional.String
    SendingBusinessRepresentativeLastName optional.String
    SendingBusinessRepresentativeLastName2 optional.String
    SendingBusinessRepresentativeNativeName optional.String
    SendingBusinessIdType optional.String
    SendingBusinessIdNumber optional.String
    SendingBusinessCountryOfIssue optional.String
    SendingBusinessIssueDate optional.String
    SendingBusinessExpireDate optional.String
    SendingBusinessDateOfIncorporation optional.String
    DocumentReferenceNumber optional.String
    SwiftCode optional.String
    BankIfcCode optional.String
    BankBranchName optional.String
    BankBranchCode optional.String
    BankAccountNumber optional.String
    CbcBankName optional.String
    BpiBankName optional.String
    PartnerName optional.String
    SourceOfIncome optional.String
    Relationship optional.String
    Reason optional.String
    SenderReference optional.String
    AdditionalInformation optional.String
}

func (a *RemittancesApiService) PostUsersWalletsFundsTransfersOverseas(ctx context.Context, localVarOptionals *RemittancesApiPostUsersWalletsFundsTransfersOverseasOpts) (RemittanceSendV1, *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Post")
		localVarPostBody   interface{}
		localVarFileName   string
		localVarFileBytes  []byte
		localVarReturnValue RemittanceSendV1
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath + "/users/wallets/funds/transfers/overseas"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	// to determine the Content-Type header
	localVarHttpContentTypes := []string{"application/x-www-form-urlencoded"}

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
	if localVarOptionals != nil && localVarOptionals.AgentId.IsSet() {
		localVarFormParams.Add("agent_id", parameterToString(localVarOptionals.AgentId.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.CalculationMode.IsSet() {
		localVarFormParams.Add("calculation_mode", parameterToString(localVarOptionals.CalculationMode.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.TransactionType.IsSet() {
		localVarFormParams.Add("transaction_type", parameterToString(localVarOptionals.TransactionType.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.Amount.IsSet() {
		localVarFormParams.Add("amount", parameterToString(localVarOptionals.Amount.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.SendCurrency.IsSet() {
		localVarFormParams.Add("send_currency", parameterToString(localVarOptionals.SendCurrency.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.ReceiveCurrency.IsSet() {
		localVarFormParams.Add("receive_currency", parameterToString(localVarOptionals.ReceiveCurrency.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.ReceiveCountry.IsSet() {
		localVarFormParams.Add("receive_country", parameterToString(localVarOptionals.ReceiveCountry.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.PaymentMode.IsSet() {
		localVarFormParams.Add("payment_mode", parameterToString(localVarOptionals.PaymentMode.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.RoutingType.IsSet() {
		localVarFormParams.Add("routing_type", parameterToString(localVarOptionals.RoutingType.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.RoutingParam.IsSet() {
		localVarFormParams.Add("routing_param", parameterToString(localVarOptionals.RoutingParam.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.QuotationId.IsSet() {
		localVarFormParams.Add("quotation_id", parameterToString(localVarOptionals.QuotationId.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.PayoutAgent.IsSet() {
		localVarFormParams.Add("payout_agent", parameterToString(localVarOptionals.PayoutAgent.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.HashId.IsSet() {
		localVarFormParams.Add("hash_id", parameterToString(localVarOptionals.HashId.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.FirstName.IsSet() {
		localVarFormParams.Add("first_name", parameterToString(localVarOptionals.FirstName.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.MiddleName.IsSet() {
		localVarFormParams.Add("middle_name", parameterToString(localVarOptionals.MiddleName.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.LastName.IsSet() {
		localVarFormParams.Add("last_name", parameterToString(localVarOptionals.LastName.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.ChineseName.IsSet() {
		localVarFormParams.Add("chinese_name", parameterToString(localVarOptionals.ChineseName.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.ReceiveMobileNumber.IsSet() {
		localVarFormParams.Add("receive_mobile_number", parameterToString(localVarOptionals.ReceiveMobileNumber.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.ReceiveGender.IsSet() {
		localVarFormParams.Add("receive_gender", parameterToString(localVarOptionals.ReceiveGender.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.IdNumber.IsSet() {
		localVarFormParams.Add("id_number", parameterToString(localVarOptionals.IdNumber.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.IdType.IsSet() {
		localVarFormParams.Add("id_type", parameterToString(localVarOptionals.IdType.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.Occupation.IsSet() {
		localVarFormParams.Add("occupation", parameterToString(localVarOptionals.Occupation.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.Address1.IsSet() {
		localVarFormParams.Add("address_1", parameterToString(localVarOptionals.Address1.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.Address2.IsSet() {
		localVarFormParams.Add("address_2", parameterToString(localVarOptionals.Address2.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.BirthDate.IsSet() {
		localVarFormParams.Add("birth_date", parameterToString(localVarOptionals.BirthDate.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.City.IsSet() {
		localVarFormParams.Add("city", parameterToString(localVarOptionals.City.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.State.IsSet() {
		localVarFormParams.Add("state", parameterToString(localVarOptionals.State.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.Country.IsSet() {
		localVarFormParams.Add("country", parameterToString(localVarOptionals.Country.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.Nationality.IsSet() {
		localVarFormParams.Add("nationality", parameterToString(localVarOptionals.Nationality.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.Zipcode.IsSet() {
		localVarFormParams.Add("zipcode", parameterToString(localVarOptionals.Zipcode.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.BeneficiaryHashId.IsSet() {
		localVarFormParams.Add("beneficiary_hash_id", parameterToString(localVarOptionals.BeneficiaryHashId.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.ReceivingBusinessRegisteredName.IsSet() {
		localVarFormParams.Add("receiving_business_registered_name", parameterToString(localVarOptionals.ReceivingBusinessRegisteredName.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.ReceivingBusinessTradingName.IsSet() {
		localVarFormParams.Add("receiving_business_trading_name", parameterToString(localVarOptionals.ReceivingBusinessTradingName.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.ReceivingBusinessEmail.IsSet() {
		localVarFormParams.Add("receiving_business_email", parameterToString(localVarOptionals.ReceivingBusinessEmail.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.ReceivingBusinessMsisdn.IsSet() {
		localVarFormParams.Add("receiving_business_msisdn", parameterToString(localVarOptionals.ReceivingBusinessMsisdn.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.ReceivingBusinessRegistrationNumber.IsSet() {
		localVarFormParams.Add("receiving_business_registration_number", parameterToString(localVarOptionals.ReceivingBusinessRegistrationNumber.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.ReceivingBusinessTaxId.IsSet() {
		localVarFormParams.Add("receiving_business_tax_id", parameterToString(localVarOptionals.ReceivingBusinessTaxId.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.ReceivingBusinessAddress1.IsSet() {
		localVarFormParams.Add("receiving_business_address1", parameterToString(localVarOptionals.ReceivingBusinessAddress1.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.ReceivingBusinessAddress2.IsSet() {
		localVarFormParams.Add("receiving_business_address2", parameterToString(localVarOptionals.ReceivingBusinessAddress2.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.ReceivingBusinessCity.IsSet() {
		localVarFormParams.Add("receiving_business_city", parameterToString(localVarOptionals.ReceivingBusinessCity.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.ReceivingBusinessState.IsSet() {
		localVarFormParams.Add("receiving_business_state", parameterToString(localVarOptionals.ReceivingBusinessState.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.ReceivingBusinessCountry.IsSet() {
		localVarFormParams.Add("receiving_business_country", parameterToString(localVarOptionals.ReceivingBusinessCountry.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.ReceivingBusinessRepresentativeFirstName.IsSet() {
		localVarFormParams.Add("receiving_business_representative_first_name", parameterToString(localVarOptionals.ReceivingBusinessRepresentativeFirstName.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.ReceivingBusinessRepresentativeMiddleName.IsSet() {
		localVarFormParams.Add("receiving_business_representative_middle_name", parameterToString(localVarOptionals.ReceivingBusinessRepresentativeMiddleName.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.ReceivingBusinessRepresentativeLastName.IsSet() {
		localVarFormParams.Add("receiving_business_representative_last_name", parameterToString(localVarOptionals.ReceivingBusinessRepresentativeLastName.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.ReceivingBusinessRepresentativeLastName2.IsSet() {
		localVarFormParams.Add("receiving_business_representative_last_name2", parameterToString(localVarOptionals.ReceivingBusinessRepresentativeLastName2.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.ReceivingBusinessRepresentativeNativeName.IsSet() {
		localVarFormParams.Add("receiving_business_representative_native_name", parameterToString(localVarOptionals.ReceivingBusinessRepresentativeNativeName.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.ReceivingBusinessIdType.IsSet() {
		localVarFormParams.Add("receiving_business_id_type", parameterToString(localVarOptionals.ReceivingBusinessIdType.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.ReceivingBusinessIdNumber.IsSet() {
		localVarFormParams.Add("receiving_business_id_number", parameterToString(localVarOptionals.ReceivingBusinessIdNumber.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.ReceivingBusinessCountryOfIssue.IsSet() {
		localVarFormParams.Add("receiving_business_country_of_issue", parameterToString(localVarOptionals.ReceivingBusinessCountryOfIssue.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.ReceivingBusinessIssueDate.IsSet() {
		localVarFormParams.Add("receiving_business_issue_date", parameterToString(localVarOptionals.ReceivingBusinessIssueDate.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.ReceivingBusinessExpireDate.IsSet() {
		localVarFormParams.Add("receiving_business_expire_date", parameterToString(localVarOptionals.ReceivingBusinessExpireDate.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.ReceivingBusinessDateOfIncorporation.IsSet() {
		localVarFormParams.Add("receiving_business_date_of_incorporation", parameterToString(localVarOptionals.ReceivingBusinessDateOfIncorporation.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.SendingBusinessHashId.IsSet() {
		localVarFormParams.Add("sending_business_hash_id", parameterToString(localVarOptionals.SendingBusinessHashId.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.SendingBusinessRegisteredName.IsSet() {
		localVarFormParams.Add("sending_business_registered_name", parameterToString(localVarOptionals.SendingBusinessRegisteredName.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.SendingBusinessTradingName.IsSet() {
		localVarFormParams.Add("sending_business_trading_name", parameterToString(localVarOptionals.SendingBusinessTradingName.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.SendingBusinessEmail.IsSet() {
		localVarFormParams.Add("sending_business_email", parameterToString(localVarOptionals.SendingBusinessEmail.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.SendingBusinessMsisdn.IsSet() {
		localVarFormParams.Add("sending_business_msisdn", parameterToString(localVarOptionals.SendingBusinessMsisdn.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.SendingBusinessRegistrationNumber.IsSet() {
		localVarFormParams.Add("sending_business_registration_number", parameterToString(localVarOptionals.SendingBusinessRegistrationNumber.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.SendingBusinessTaxId.IsSet() {
		localVarFormParams.Add("sending_business_tax_id", parameterToString(localVarOptionals.SendingBusinessTaxId.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.SendingBusinessAddress1.IsSet() {
		localVarFormParams.Add("sending_business_address1", parameterToString(localVarOptionals.SendingBusinessAddress1.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.SendingBusinessAddress2.IsSet() {
		localVarFormParams.Add("sending_business_address2", parameterToString(localVarOptionals.SendingBusinessAddress2.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.SendingBusinessCity.IsSet() {
		localVarFormParams.Add("sending_business_city", parameterToString(localVarOptionals.SendingBusinessCity.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.SendingBusinessState.IsSet() {
		localVarFormParams.Add("sending_business_state", parameterToString(localVarOptionals.SendingBusinessState.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.SendingBusinessZipcode.IsSet() {
		localVarFormParams.Add("sending_business_zipcode", parameterToString(localVarOptionals.SendingBusinessZipcode.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.SendingBusinessCountry.IsSet() {
		localVarFormParams.Add("sending_business_country", parameterToString(localVarOptionals.SendingBusinessCountry.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.SendingBusinessRepresentativeFirstName.IsSet() {
		localVarFormParams.Add("sending_business_representative_first_name", parameterToString(localVarOptionals.SendingBusinessRepresentativeFirstName.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.SendingBusinessRepresentativeMiddleName.IsSet() {
		localVarFormParams.Add("sending_business_representative_middle_name", parameterToString(localVarOptionals.SendingBusinessRepresentativeMiddleName.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.SendingBusinessRepresentativeLastName.IsSet() {
		localVarFormParams.Add("sending_business_representative_last_name", parameterToString(localVarOptionals.SendingBusinessRepresentativeLastName.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.SendingBusinessRepresentativeLastName2.IsSet() {
		localVarFormParams.Add("sending_business_representative_last_name2", parameterToString(localVarOptionals.SendingBusinessRepresentativeLastName2.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.SendingBusinessRepresentativeNativeName.IsSet() {
		localVarFormParams.Add("sending_business_representative_native_name", parameterToString(localVarOptionals.SendingBusinessRepresentativeNativeName.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.SendingBusinessIdType.IsSet() {
		localVarFormParams.Add("sending_business_id_type", parameterToString(localVarOptionals.SendingBusinessIdType.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.SendingBusinessIdNumber.IsSet() {
		localVarFormParams.Add("sending_business_id_number", parameterToString(localVarOptionals.SendingBusinessIdNumber.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.SendingBusinessCountryOfIssue.IsSet() {
		localVarFormParams.Add("sending_business_country_of_issue", parameterToString(localVarOptionals.SendingBusinessCountryOfIssue.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.SendingBusinessIssueDate.IsSet() {
		localVarFormParams.Add("sending_business_issue_date", parameterToString(localVarOptionals.SendingBusinessIssueDate.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.SendingBusinessExpireDate.IsSet() {
		localVarFormParams.Add("sending_business_expire_date", parameterToString(localVarOptionals.SendingBusinessExpireDate.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.SendingBusinessDateOfIncorporation.IsSet() {
		localVarFormParams.Add("sending_business_date_of_incorporation", parameterToString(localVarOptionals.SendingBusinessDateOfIncorporation.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.DocumentReferenceNumber.IsSet() {
		localVarFormParams.Add("document_reference_number", parameterToString(localVarOptionals.DocumentReferenceNumber.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.SwiftCode.IsSet() {
		localVarFormParams.Add("swift_code", parameterToString(localVarOptionals.SwiftCode.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.BankIfcCode.IsSet() {
		localVarFormParams.Add("bank_ifc_code", parameterToString(localVarOptionals.BankIfcCode.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.BankBranchName.IsSet() {
		localVarFormParams.Add("bank_branch_name", parameterToString(localVarOptionals.BankBranchName.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.BankBranchCode.IsSet() {
		localVarFormParams.Add("bank_branch_code", parameterToString(localVarOptionals.BankBranchCode.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.BankAccountNumber.IsSet() {
		localVarFormParams.Add("bank_account_number", parameterToString(localVarOptionals.BankAccountNumber.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.CbcBankName.IsSet() {
		localVarFormParams.Add("cbc_bank_name", parameterToString(localVarOptionals.CbcBankName.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.BpiBankName.IsSet() {
		localVarFormParams.Add("bpi_bank_name", parameterToString(localVarOptionals.BpiBankName.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.PartnerName.IsSet() {
		localVarFormParams.Add("partner_name", parameterToString(localVarOptionals.PartnerName.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.SourceOfIncome.IsSet() {
		localVarFormParams.Add("source_of_income", parameterToString(localVarOptionals.SourceOfIncome.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.Relationship.IsSet() {
		localVarFormParams.Add("relationship", parameterToString(localVarOptionals.Relationship.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.Reason.IsSet() {
		localVarFormParams.Add("reason", parameterToString(localVarOptionals.Reason.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.SenderReference.IsSet() {
		localVarFormParams.Add("sender_reference", parameterToString(localVarOptionals.SenderReference.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.AdditionalInformation.IsSet() {
		localVarFormParams.Add("additional_information", parameterToString(localVarOptionals.AdditionalInformation.Value(), ""))
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
			var v RemittanceSendV1
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
RemittancesApiService Upload Remittance Attachment
Attaches document on the given remittance hash id
 * @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param data
 * @param documentType
 * @param documentName
 * @param xAuthUserID MatchMove provided hash ID for the user
 * @param id Remittance Hash ID retrieved after (#post-users-wallets-funds-transfers-overseas-type)
@return RemittanceAttachmentV1
*/
func (a *RemittancesApiService) PostUsersWalletsFundsTransfersOverseasIdAttachment(ctx context.Context, data string, documentType string, documentName string, xAuthUserID string, id string) (RemittanceAttachmentV1, *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Post")
		localVarPostBody   interface{}
		localVarFileName   string
		localVarFileBytes  []byte
		localVarReturnValue RemittanceAttachmentV1
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath + "/users/wallets/funds/transfers/overseas/{id}/attachment"
	localVarPath = strings.Replace(localVarPath, "{"+"id"+"}", fmt.Sprintf("%v", id), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if strlen(xAuthUserID) > 32 {
		return localVarReturnValue, nil, reportError("xAuthUserID must have less than 32 elements")
	}

	// to determine the Content-Type header
	localVarHttpContentTypes := []string{"application/x-www-form-urlencoded"}

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
	localVarHeaderParams["X-Auth-User-ID"] = parameterToString(xAuthUserID, "")
	localVarFormParams.Add("data", parameterToString(data, ""))
	localVarFormParams.Add("document_type", parameterToString(documentType, ""))
	localVarFormParams.Add("document_name", parameterToString(documentName, ""))
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
			var v RemittanceAttachmentV1
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
RemittancesApiService Calculate Fees for Overseas Transfers without providing Rails
Calculate the exchange rate and fees applicable for the provided agent before an actual remittance transaction is created.  A beneficiary form is also returned as part of the validation process when creating the actual remittance transaction.  No transaction record is created yet when this endpoint is called. 
 * @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param optional nil or *RemittancesApiPostUsersWalletsFundsTransfersOverseasTestOpts - Optional Parameters:
     * @param "AgentId" (optional.String) - 
     * @param "CalculationMode" (optional.String) - 
     * @param "Amount" (optional.String) - 
@return RemittanceCalculateV1
*/

type RemittancesApiPostUsersWalletsFundsTransfersOverseasTestOpts struct {
    AgentId optional.String
    CalculationMode optional.String
    Amount optional.String
}

func (a *RemittancesApiService) PostUsersWalletsFundsTransfersOverseasTest(ctx context.Context, localVarOptionals *RemittancesApiPostUsersWalletsFundsTransfersOverseasTestOpts) (RemittanceCalculateV1, *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Post")
		localVarPostBody   interface{}
		localVarFileName   string
		localVarFileBytes  []byte
		localVarReturnValue RemittanceCalculateV1
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath + "/users/wallets/funds/transfers/overseas/test"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	// to determine the Content-Type header
	localVarHttpContentTypes := []string{"application/x-www-form-urlencoded"}

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
	if localVarOptionals != nil && localVarOptionals.AgentId.IsSet() {
		localVarFormParams.Add("agent_id", parameterToString(localVarOptionals.AgentId.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.CalculationMode.IsSet() {
		localVarFormParams.Add("calculation_mode", parameterToString(localVarOptionals.CalculationMode.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.Amount.IsSet() {
		localVarFormParams.Add("amount", parameterToString(localVarOptionals.Amount.Value(), ""))
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
			var v RemittanceCalculateV1
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
RemittancesApiService Calculate or Create remittance transaction
Send a money transfer request with the specified overseas provider.  Any transaction made through this resource will be on pending validation. Confirmation must be made to clear the transaction.  Appending &#x60;test&#x60; to the (users/wallets/funds/transfers/overseas/{type}/test) sends a calculation test (ONLY) which useful for retrieving the fees applicable before creating a real transaction.  With administrative scope, this api can be accessed as a public resource. The additional parameter &#x60;hash_id&#x60; must be supplied when using this as a public resource. 
 * @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param type_ Overseas Provider Code
 * @param optional nil or *RemittancesApiPostUsersWalletsFundsTransfersOverseasTypeOpts - Optional Parameters:
     * @param "AgentId" (optional.String) - 
     * @param "CalculationMode" (optional.String) - 
     * @param "TransactionType" (optional.String) - 
     * @param "Amount" (optional.String) - 
     * @param "SendCurrency" (optional.String) - 
     * @param "ReceiveCurrency" (optional.String) - 
     * @param "ReceiveCountry" (optional.String) - 
     * @param "PaymentMode" (optional.String) - 
     * @param "RoutingType" (optional.String) - 
     * @param "RoutingParam" (optional.String) - 
     * @param "QuotationId" (optional.String) - 
     * @param "PayoutAgent" (optional.String) - 
     * @param "HashId" (optional.String) - 
     * @param "FirstName" (optional.String) - 
     * @param "MiddleName" (optional.String) - 
     * @param "LastName" (optional.String) - 
     * @param "ChineseName" (optional.String) - 
     * @param "ReceiveMobileNumber" (optional.String) - 
     * @param "ReceiveGender" (optional.String) - 
     * @param "IdNumber" (optional.String) - 
     * @param "IdType" (optional.String) - 
     * @param "Occupation" (optional.String) - 
     * @param "Address1" (optional.String) - 
     * @param "Address2" (optional.String) - 
     * @param "BirthDate" (optional.String) - 
     * @param "City" (optional.String) - 
     * @param "State" (optional.String) - 
     * @param "Country" (optional.String) - 
     * @param "Nationality" (optional.String) - 
     * @param "Zipcode" (optional.String) - 
     * @param "BeneficiaryHashId" (optional.String) - 
     * @param "ReceivingBusinessRegisteredName" (optional.String) - 
     * @param "ReceivingBusinessTradingName" (optional.String) - 
     * @param "ReceivingBusinessEmail" (optional.String) - 
     * @param "ReceivingBusinessMsisdn" (optional.String) - 
     * @param "ReceivingBusinessRegistrationNumber" (optional.String) - 
     * @param "ReceivingBusinessTaxId" (optional.String) - 
     * @param "ReceivingBusinessAddress1" (optional.String) - 
     * @param "ReceivingBusinessAddress2" (optional.String) - 
     * @param "ReceivingBusinessCity" (optional.String) - 
     * @param "ReceivingBusinessState" (optional.String) - 
     * @param "ReceivingBusinessCountry" (optional.String) - 
     * @param "ReceivingBusinessRepresentativeFirstName" (optional.String) - 
     * @param "ReceivingBusinessRepresentativeMiddleName" (optional.String) - 
     * @param "ReceivingBusinessRepresentativeLastName" (optional.String) - 
     * @param "ReceivingBusinessRepresentativeLastName2" (optional.String) - 
     * @param "ReceivingBusinessRepresentativeNativeName" (optional.String) - 
     * @param "ReceivingBusinessIdType" (optional.String) - 
     * @param "ReceivingBusinessIdNumber" (optional.String) - 
     * @param "ReceivingBusinessCountryOfIssue" (optional.String) - 
     * @param "ReceivingBusinessIssueDate" (optional.String) - 
     * @param "ReceivingBusinessExpireDate" (optional.String) - 
     * @param "ReceivingBusinessDateOfIncorporation" (optional.String) - 
     * @param "SendingBusinessHashId" (optional.String) - 
     * @param "SendingBusinessRegisteredName" (optional.String) - 
     * @param "SendingBusinessTradingName" (optional.String) - 
     * @param "SendingBusinessEmail" (optional.String) - 
     * @param "SendingBusinessMsisdn" (optional.String) - 
     * @param "SendingBusinessRegistrationNumber" (optional.String) - 
     * @param "SendingBusinessTaxId" (optional.String) - 
     * @param "SendingBusinessAddress1" (optional.String) - 
     * @param "SendingBusinessAddress2" (optional.String) - 
     * @param "SendingBusinessCity" (optional.String) - 
     * @param "SendingBusinessState" (optional.String) - 
     * @param "SendingBusinessZipcode" (optional.String) - 
     * @param "SendingBusinessCountry" (optional.String) - 
     * @param "SendingBusinessRepresentativeFirstName" (optional.String) - 
     * @param "SendingBusinessRepresentativeMiddleName" (optional.String) - 
     * @param "SendingBusinessRepresentativeLastName" (optional.String) - 
     * @param "SendingBusinessRepresentativeLastName2" (optional.String) - 
     * @param "SendingBusinessRepresentativeNativeName" (optional.String) - 
     * @param "SendingBusinessIdType" (optional.String) - 
     * @param "SendingBusinessIdNumber" (optional.String) - 
     * @param "SendingBusinessCountryOfIssue" (optional.String) - 
     * @param "SendingBusinessIssueDate" (optional.String) - 
     * @param "SendingBusinessExpireDate" (optional.String) - 
     * @param "SendingBusinessDateOfIncorporation" (optional.String) - 
     * @param "DocumentReferenceNumber" (optional.String) - 
     * @param "SwiftCode" (optional.String) - 
     * @param "BankIfcCode" (optional.String) - 
     * @param "BankBranchName" (optional.String) - 
     * @param "BankBranchCode" (optional.String) - 
     * @param "BankAccountNumber" (optional.String) - 
     * @param "CbcBankName" (optional.String) - 
     * @param "BpiBankName" (optional.String) - 
     * @param "PartnerName" (optional.String) - 
     * @param "SourceOfIncome" (optional.String) - 
     * @param "Relationship" (optional.String) - 
     * @param "Reason" (optional.String) - 
     * @param "SenderReference" (optional.String) - 
     * @param "AdditionalInformation" (optional.String) - 
@return InlineResponse20052
*/

type RemittancesApiPostUsersWalletsFundsTransfersOverseasTypeOpts struct {
    AgentId optional.String
    CalculationMode optional.String
    TransactionType optional.String
    Amount optional.String
    SendCurrency optional.String
    ReceiveCurrency optional.String
    ReceiveCountry optional.String
    PaymentMode optional.String
    RoutingType optional.String
    RoutingParam optional.String
    QuotationId optional.String
    PayoutAgent optional.String
    HashId optional.String
    FirstName optional.String
    MiddleName optional.String
    LastName optional.String
    ChineseName optional.String
    ReceiveMobileNumber optional.String
    ReceiveGender optional.String
    IdNumber optional.String
    IdType optional.String
    Occupation optional.String
    Address1 optional.String
    Address2 optional.String
    BirthDate optional.String
    City optional.String
    State optional.String
    Country optional.String
    Nationality optional.String
    Zipcode optional.String
    BeneficiaryHashId optional.String
    ReceivingBusinessRegisteredName optional.String
    ReceivingBusinessTradingName optional.String
    ReceivingBusinessEmail optional.String
    ReceivingBusinessMsisdn optional.String
    ReceivingBusinessRegistrationNumber optional.String
    ReceivingBusinessTaxId optional.String
    ReceivingBusinessAddress1 optional.String
    ReceivingBusinessAddress2 optional.String
    ReceivingBusinessCity optional.String
    ReceivingBusinessState optional.String
    ReceivingBusinessCountry optional.String
    ReceivingBusinessRepresentativeFirstName optional.String
    ReceivingBusinessRepresentativeMiddleName optional.String
    ReceivingBusinessRepresentativeLastName optional.String
    ReceivingBusinessRepresentativeLastName2 optional.String
    ReceivingBusinessRepresentativeNativeName optional.String
    ReceivingBusinessIdType optional.String
    ReceivingBusinessIdNumber optional.String
    ReceivingBusinessCountryOfIssue optional.String
    ReceivingBusinessIssueDate optional.String
    ReceivingBusinessExpireDate optional.String
    ReceivingBusinessDateOfIncorporation optional.String
    SendingBusinessHashId optional.String
    SendingBusinessRegisteredName optional.String
    SendingBusinessTradingName optional.String
    SendingBusinessEmail optional.String
    SendingBusinessMsisdn optional.String
    SendingBusinessRegistrationNumber optional.String
    SendingBusinessTaxId optional.String
    SendingBusinessAddress1 optional.String
    SendingBusinessAddress2 optional.String
    SendingBusinessCity optional.String
    SendingBusinessState optional.String
    SendingBusinessZipcode optional.String
    SendingBusinessCountry optional.String
    SendingBusinessRepresentativeFirstName optional.String
    SendingBusinessRepresentativeMiddleName optional.String
    SendingBusinessRepresentativeLastName optional.String
    SendingBusinessRepresentativeLastName2 optional.String
    SendingBusinessRepresentativeNativeName optional.String
    SendingBusinessIdType optional.String
    SendingBusinessIdNumber optional.String
    SendingBusinessCountryOfIssue optional.String
    SendingBusinessIssueDate optional.String
    SendingBusinessExpireDate optional.String
    SendingBusinessDateOfIncorporation optional.String
    DocumentReferenceNumber optional.String
    SwiftCode optional.String
    BankIfcCode optional.String
    BankBranchName optional.String
    BankBranchCode optional.String
    BankAccountNumber optional.String
    CbcBankName optional.String
    BpiBankName optional.String
    PartnerName optional.String
    SourceOfIncome optional.String
    Relationship optional.String
    Reason optional.String
    SenderReference optional.String
    AdditionalInformation optional.String
}

func (a *RemittancesApiService) PostUsersWalletsFundsTransfersOverseasType(ctx context.Context, type_ string, localVarOptionals *RemittancesApiPostUsersWalletsFundsTransfersOverseasTypeOpts) (InlineResponse20052, *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Post")
		localVarPostBody   interface{}
		localVarFileName   string
		localVarFileBytes  []byte
		localVarReturnValue InlineResponse20052
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath + "/users/wallets/funds/transfers/overseas/{type}"
	localVarPath = strings.Replace(localVarPath, "{"+"type"+"}", fmt.Sprintf("%v", type_), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	// to determine the Content-Type header
	localVarHttpContentTypes := []string{"application/x-www-form-urlencoded"}

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
	if localVarOptionals != nil && localVarOptionals.AgentId.IsSet() {
		localVarFormParams.Add("agent_id", parameterToString(localVarOptionals.AgentId.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.CalculationMode.IsSet() {
		localVarFormParams.Add("calculation_mode", parameterToString(localVarOptionals.CalculationMode.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.TransactionType.IsSet() {
		localVarFormParams.Add("transaction_type", parameterToString(localVarOptionals.TransactionType.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.Amount.IsSet() {
		localVarFormParams.Add("amount", parameterToString(localVarOptionals.Amount.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.SendCurrency.IsSet() {
		localVarFormParams.Add("send_currency", parameterToString(localVarOptionals.SendCurrency.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.ReceiveCurrency.IsSet() {
		localVarFormParams.Add("receive_currency", parameterToString(localVarOptionals.ReceiveCurrency.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.ReceiveCountry.IsSet() {
		localVarFormParams.Add("receive_country", parameterToString(localVarOptionals.ReceiveCountry.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.PaymentMode.IsSet() {
		localVarFormParams.Add("payment_mode", parameterToString(localVarOptionals.PaymentMode.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.RoutingType.IsSet() {
		localVarFormParams.Add("routing_type", parameterToString(localVarOptionals.RoutingType.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.RoutingParam.IsSet() {
		localVarFormParams.Add("routing_param", parameterToString(localVarOptionals.RoutingParam.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.QuotationId.IsSet() {
		localVarFormParams.Add("quotation_id", parameterToString(localVarOptionals.QuotationId.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.PayoutAgent.IsSet() {
		localVarFormParams.Add("payout_agent", parameterToString(localVarOptionals.PayoutAgent.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.HashId.IsSet() {
		localVarFormParams.Add("hash_id", parameterToString(localVarOptionals.HashId.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.FirstName.IsSet() {
		localVarFormParams.Add("first_name", parameterToString(localVarOptionals.FirstName.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.MiddleName.IsSet() {
		localVarFormParams.Add("middle_name", parameterToString(localVarOptionals.MiddleName.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.LastName.IsSet() {
		localVarFormParams.Add("last_name", parameterToString(localVarOptionals.LastName.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.ChineseName.IsSet() {
		localVarFormParams.Add("chinese_name", parameterToString(localVarOptionals.ChineseName.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.ReceiveMobileNumber.IsSet() {
		localVarFormParams.Add("receive_mobile_number", parameterToString(localVarOptionals.ReceiveMobileNumber.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.ReceiveGender.IsSet() {
		localVarFormParams.Add("receive_gender", parameterToString(localVarOptionals.ReceiveGender.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.IdNumber.IsSet() {
		localVarFormParams.Add("id_number", parameterToString(localVarOptionals.IdNumber.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.IdType.IsSet() {
		localVarFormParams.Add("id_type", parameterToString(localVarOptionals.IdType.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.Occupation.IsSet() {
		localVarFormParams.Add("occupation", parameterToString(localVarOptionals.Occupation.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.Address1.IsSet() {
		localVarFormParams.Add("address_1", parameterToString(localVarOptionals.Address1.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.Address2.IsSet() {
		localVarFormParams.Add("address_2", parameterToString(localVarOptionals.Address2.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.BirthDate.IsSet() {
		localVarFormParams.Add("birth_date", parameterToString(localVarOptionals.BirthDate.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.City.IsSet() {
		localVarFormParams.Add("city", parameterToString(localVarOptionals.City.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.State.IsSet() {
		localVarFormParams.Add("state", parameterToString(localVarOptionals.State.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.Country.IsSet() {
		localVarFormParams.Add("country", parameterToString(localVarOptionals.Country.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.Nationality.IsSet() {
		localVarFormParams.Add("nationality", parameterToString(localVarOptionals.Nationality.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.Zipcode.IsSet() {
		localVarFormParams.Add("zipcode", parameterToString(localVarOptionals.Zipcode.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.BeneficiaryHashId.IsSet() {
		localVarFormParams.Add("beneficiary_hash_id", parameterToString(localVarOptionals.BeneficiaryHashId.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.ReceivingBusinessRegisteredName.IsSet() {
		localVarFormParams.Add("receiving_business_registered_name", parameterToString(localVarOptionals.ReceivingBusinessRegisteredName.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.ReceivingBusinessTradingName.IsSet() {
		localVarFormParams.Add("receiving_business_trading_name", parameterToString(localVarOptionals.ReceivingBusinessTradingName.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.ReceivingBusinessEmail.IsSet() {
		localVarFormParams.Add("receiving_business_email", parameterToString(localVarOptionals.ReceivingBusinessEmail.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.ReceivingBusinessMsisdn.IsSet() {
		localVarFormParams.Add("receiving_business_msisdn", parameterToString(localVarOptionals.ReceivingBusinessMsisdn.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.ReceivingBusinessRegistrationNumber.IsSet() {
		localVarFormParams.Add("receiving_business_registration_number", parameterToString(localVarOptionals.ReceivingBusinessRegistrationNumber.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.ReceivingBusinessTaxId.IsSet() {
		localVarFormParams.Add("receiving_business_tax_id", parameterToString(localVarOptionals.ReceivingBusinessTaxId.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.ReceivingBusinessAddress1.IsSet() {
		localVarFormParams.Add("receiving_business_address1", parameterToString(localVarOptionals.ReceivingBusinessAddress1.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.ReceivingBusinessAddress2.IsSet() {
		localVarFormParams.Add("receiving_business_address2", parameterToString(localVarOptionals.ReceivingBusinessAddress2.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.ReceivingBusinessCity.IsSet() {
		localVarFormParams.Add("receiving_business_city", parameterToString(localVarOptionals.ReceivingBusinessCity.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.ReceivingBusinessState.IsSet() {
		localVarFormParams.Add("receiving_business_state", parameterToString(localVarOptionals.ReceivingBusinessState.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.ReceivingBusinessCountry.IsSet() {
		localVarFormParams.Add("receiving_business_country", parameterToString(localVarOptionals.ReceivingBusinessCountry.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.ReceivingBusinessRepresentativeFirstName.IsSet() {
		localVarFormParams.Add("receiving_business_representative_first_name", parameterToString(localVarOptionals.ReceivingBusinessRepresentativeFirstName.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.ReceivingBusinessRepresentativeMiddleName.IsSet() {
		localVarFormParams.Add("receiving_business_representative_middle_name", parameterToString(localVarOptionals.ReceivingBusinessRepresentativeMiddleName.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.ReceivingBusinessRepresentativeLastName.IsSet() {
		localVarFormParams.Add("receiving_business_representative_last_name", parameterToString(localVarOptionals.ReceivingBusinessRepresentativeLastName.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.ReceivingBusinessRepresentativeLastName2.IsSet() {
		localVarFormParams.Add("receiving_business_representative_last_name2", parameterToString(localVarOptionals.ReceivingBusinessRepresentativeLastName2.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.ReceivingBusinessRepresentativeNativeName.IsSet() {
		localVarFormParams.Add("receiving_business_representative_native_name", parameterToString(localVarOptionals.ReceivingBusinessRepresentativeNativeName.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.ReceivingBusinessIdType.IsSet() {
		localVarFormParams.Add("receiving_business_id_type", parameterToString(localVarOptionals.ReceivingBusinessIdType.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.ReceivingBusinessIdNumber.IsSet() {
		localVarFormParams.Add("receiving_business_id_number", parameterToString(localVarOptionals.ReceivingBusinessIdNumber.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.ReceivingBusinessCountryOfIssue.IsSet() {
		localVarFormParams.Add("receiving_business_country_of_issue", parameterToString(localVarOptionals.ReceivingBusinessCountryOfIssue.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.ReceivingBusinessIssueDate.IsSet() {
		localVarFormParams.Add("receiving_business_issue_date", parameterToString(localVarOptionals.ReceivingBusinessIssueDate.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.ReceivingBusinessExpireDate.IsSet() {
		localVarFormParams.Add("receiving_business_expire_date", parameterToString(localVarOptionals.ReceivingBusinessExpireDate.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.ReceivingBusinessDateOfIncorporation.IsSet() {
		localVarFormParams.Add("receiving_business_date_of_incorporation", parameterToString(localVarOptionals.ReceivingBusinessDateOfIncorporation.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.SendingBusinessHashId.IsSet() {
		localVarFormParams.Add("sending_business_hash_id", parameterToString(localVarOptionals.SendingBusinessHashId.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.SendingBusinessRegisteredName.IsSet() {
		localVarFormParams.Add("sending_business_registered_name", parameterToString(localVarOptionals.SendingBusinessRegisteredName.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.SendingBusinessTradingName.IsSet() {
		localVarFormParams.Add("sending_business_trading_name", parameterToString(localVarOptionals.SendingBusinessTradingName.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.SendingBusinessEmail.IsSet() {
		localVarFormParams.Add("sending_business_email", parameterToString(localVarOptionals.SendingBusinessEmail.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.SendingBusinessMsisdn.IsSet() {
		localVarFormParams.Add("sending_business_msisdn", parameterToString(localVarOptionals.SendingBusinessMsisdn.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.SendingBusinessRegistrationNumber.IsSet() {
		localVarFormParams.Add("sending_business_registration_number", parameterToString(localVarOptionals.SendingBusinessRegistrationNumber.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.SendingBusinessTaxId.IsSet() {
		localVarFormParams.Add("sending_business_tax_id", parameterToString(localVarOptionals.SendingBusinessTaxId.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.SendingBusinessAddress1.IsSet() {
		localVarFormParams.Add("sending_business_address1", parameterToString(localVarOptionals.SendingBusinessAddress1.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.SendingBusinessAddress2.IsSet() {
		localVarFormParams.Add("sending_business_address2", parameterToString(localVarOptionals.SendingBusinessAddress2.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.SendingBusinessCity.IsSet() {
		localVarFormParams.Add("sending_business_city", parameterToString(localVarOptionals.SendingBusinessCity.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.SendingBusinessState.IsSet() {
		localVarFormParams.Add("sending_business_state", parameterToString(localVarOptionals.SendingBusinessState.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.SendingBusinessZipcode.IsSet() {
		localVarFormParams.Add("sending_business_zipcode", parameterToString(localVarOptionals.SendingBusinessZipcode.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.SendingBusinessCountry.IsSet() {
		localVarFormParams.Add("sending_business_country", parameterToString(localVarOptionals.SendingBusinessCountry.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.SendingBusinessRepresentativeFirstName.IsSet() {
		localVarFormParams.Add("sending_business_representative_first_name", parameterToString(localVarOptionals.SendingBusinessRepresentativeFirstName.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.SendingBusinessRepresentativeMiddleName.IsSet() {
		localVarFormParams.Add("sending_business_representative_middle_name", parameterToString(localVarOptionals.SendingBusinessRepresentativeMiddleName.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.SendingBusinessRepresentativeLastName.IsSet() {
		localVarFormParams.Add("sending_business_representative_last_name", parameterToString(localVarOptionals.SendingBusinessRepresentativeLastName.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.SendingBusinessRepresentativeLastName2.IsSet() {
		localVarFormParams.Add("sending_business_representative_last_name2", parameterToString(localVarOptionals.SendingBusinessRepresentativeLastName2.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.SendingBusinessRepresentativeNativeName.IsSet() {
		localVarFormParams.Add("sending_business_representative_native_name", parameterToString(localVarOptionals.SendingBusinessRepresentativeNativeName.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.SendingBusinessIdType.IsSet() {
		localVarFormParams.Add("sending_business_id_type", parameterToString(localVarOptionals.SendingBusinessIdType.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.SendingBusinessIdNumber.IsSet() {
		localVarFormParams.Add("sending_business_id_number", parameterToString(localVarOptionals.SendingBusinessIdNumber.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.SendingBusinessCountryOfIssue.IsSet() {
		localVarFormParams.Add("sending_business_country_of_issue", parameterToString(localVarOptionals.SendingBusinessCountryOfIssue.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.SendingBusinessIssueDate.IsSet() {
		localVarFormParams.Add("sending_business_issue_date", parameterToString(localVarOptionals.SendingBusinessIssueDate.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.SendingBusinessExpireDate.IsSet() {
		localVarFormParams.Add("sending_business_expire_date", parameterToString(localVarOptionals.SendingBusinessExpireDate.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.SendingBusinessDateOfIncorporation.IsSet() {
		localVarFormParams.Add("sending_business_date_of_incorporation", parameterToString(localVarOptionals.SendingBusinessDateOfIncorporation.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.DocumentReferenceNumber.IsSet() {
		localVarFormParams.Add("document_reference_number", parameterToString(localVarOptionals.DocumentReferenceNumber.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.SwiftCode.IsSet() {
		localVarFormParams.Add("swift_code", parameterToString(localVarOptionals.SwiftCode.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.BankIfcCode.IsSet() {
		localVarFormParams.Add("bank_ifc_code", parameterToString(localVarOptionals.BankIfcCode.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.BankBranchName.IsSet() {
		localVarFormParams.Add("bank_branch_name", parameterToString(localVarOptionals.BankBranchName.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.BankBranchCode.IsSet() {
		localVarFormParams.Add("bank_branch_code", parameterToString(localVarOptionals.BankBranchCode.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.BankAccountNumber.IsSet() {
		localVarFormParams.Add("bank_account_number", parameterToString(localVarOptionals.BankAccountNumber.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.CbcBankName.IsSet() {
		localVarFormParams.Add("cbc_bank_name", parameterToString(localVarOptionals.CbcBankName.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.BpiBankName.IsSet() {
		localVarFormParams.Add("bpi_bank_name", parameterToString(localVarOptionals.BpiBankName.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.PartnerName.IsSet() {
		localVarFormParams.Add("partner_name", parameterToString(localVarOptionals.PartnerName.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.SourceOfIncome.IsSet() {
		localVarFormParams.Add("source_of_income", parameterToString(localVarOptionals.SourceOfIncome.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.Relationship.IsSet() {
		localVarFormParams.Add("relationship", parameterToString(localVarOptionals.Relationship.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.Reason.IsSet() {
		localVarFormParams.Add("reason", parameterToString(localVarOptionals.Reason.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.SenderReference.IsSet() {
		localVarFormParams.Add("sender_reference", parameterToString(localVarOptionals.SenderReference.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.AdditionalInformation.IsSet() {
		localVarFormParams.Add("additional_information", parameterToString(localVarOptionals.AdditionalInformation.Value(), ""))
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
			var v InlineResponse20052
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
RemittancesApiService Calculate Rates for Remittance Transaction
Send a money transfer request with the specified overseas provider.  Any transaction made through this resource will be on pending validation. Confirmation must be made to clear the transaction.  Appending &#x60;test&#x60; to the (users/wallets/funds/transfers/overseas/{type}/test) sends a calculation test (ONLY) which is useful for retrieving the fees applicable before creating a real transaction.  With administrative scope, this api can be accessed as a public resource. The additional parameter &#x60;hash_id&#x60; must be supplied when using this as a public resource. 
 * @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param type_ Overseas Provider Code
 * @param optional nil or *RemittancesApiPostUsersWalletsFundsTransfersOverseasTypeTestOpts - Optional Parameters:
     * @param "AgentId" (optional.String) - 
     * @param "CalculationMode" (optional.String) - 
     * @param "TransactionType" (optional.String) - 
     * @param "Amount" (optional.String) - 
     * @param "SendCurrency" (optional.String) - 
     * @param "ReceiveCurrency" (optional.String) - 
     * @param "ReceiveCountry" (optional.String) - 
     * @param "PaymentMode" (optional.String) - 
     * @param "RoutingType" (optional.String) - 
     * @param "RoutingParam" (optional.String) - 
     * @param "QuotationId" (optional.String) - 
     * @param "PayoutAgent" (optional.String) - 
     * @param "HashId" (optional.String) - 
     * @param "FirstName" (optional.String) - 
     * @param "MiddleName" (optional.String) - 
     * @param "LastName" (optional.String) - 
     * @param "ChineseName" (optional.String) - 
     * @param "ReceiveMobileNumber" (optional.String) - 
     * @param "ReceiveGender" (optional.String) - 
     * @param "IdNumber" (optional.String) - 
     * @param "IdType" (optional.String) - 
     * @param "Occupation" (optional.String) - 
     * @param "Address1" (optional.String) - 
     * @param "Address2" (optional.String) - 
     * @param "BirthDate" (optional.String) - 
     * @param "City" (optional.String) - 
     * @param "State" (optional.String) - 
     * @param "Country" (optional.String) - 
     * @param "Nationality" (optional.String) - 
     * @param "Zipcode" (optional.String) - 
     * @param "BeneficiaryHashId" (optional.String) - 
     * @param "ReceivingBusinessRegisteredName" (optional.String) - 
     * @param "ReceivingBusinessTradingName" (optional.String) - 
     * @param "ReceivingBusinessEmail" (optional.String) - 
     * @param "ReceivingBusinessMsisdn" (optional.String) - 
     * @param "ReceivingBusinessRegistrationNumber" (optional.String) - 
     * @param "ReceivingBusinessTaxId" (optional.String) - 
     * @param "ReceivingBusinessAddress1" (optional.String) - 
     * @param "ReceivingBusinessAddress2" (optional.String) - 
     * @param "ReceivingBusinessCity" (optional.String) - 
     * @param "ReceivingBusinessState" (optional.String) - 
     * @param "ReceivingBusinessCountry" (optional.String) - 
     * @param "ReceivingBusinessRepresentativeFirstName" (optional.String) - 
     * @param "ReceivingBusinessRepresentativeMiddleName" (optional.String) - 
     * @param "ReceivingBusinessRepresentativeLastName" (optional.String) - 
     * @param "ReceivingBusinessRepresentativeLastName2" (optional.String) - 
     * @param "ReceivingBusinessRepresentativeNativeName" (optional.String) - 
     * @param "ReceivingBusinessIdType" (optional.String) - 
     * @param "ReceivingBusinessIdNumber" (optional.String) - 
     * @param "ReceivingBusinessCountryOfIssue" (optional.String) - 
     * @param "ReceivingBusinessIssueDate" (optional.String) - 
     * @param "ReceivingBusinessExpireDate" (optional.String) - 
     * @param "ReceivingBusinessDateOfIncorporation" (optional.String) - 
     * @param "SendingBusinessHashId" (optional.String) - 
     * @param "SendingBusinessRegisteredName" (optional.String) - 
     * @param "SendingBusinessTradingName" (optional.String) - 
     * @param "SendingBusinessEmail" (optional.String) - 
     * @param "SendingBusinessMsisdn" (optional.String) - 
     * @param "SendingBusinessRegistrationNumber" (optional.String) - 
     * @param "SendingBusinessTaxId" (optional.String) - 
     * @param "SendingBusinessAddress1" (optional.String) - 
     * @param "SendingBusinessAddress2" (optional.String) - 
     * @param "SendingBusinessCity" (optional.String) - 
     * @param "SendingBusinessState" (optional.String) - 
     * @param "SendingBusinessZipcode" (optional.String) - 
     * @param "SendingBusinessCountry" (optional.String) - 
     * @param "SendingBusinessRepresentativeFirstName" (optional.String) - 
     * @param "SendingBusinessRepresentativeMiddleName" (optional.String) - 
     * @param "SendingBusinessRepresentativeLastName" (optional.String) - 
     * @param "SendingBusinessRepresentativeLastName2" (optional.String) - 
     * @param "SendingBusinessRepresentativeNativeName" (optional.String) - 
     * @param "SendingBusinessIdType" (optional.String) - 
     * @param "SendingBusinessIdNumber" (optional.String) - 
     * @param "SendingBusinessCountryOfIssue" (optional.String) - 
     * @param "SendingBusinessIssueDate" (optional.String) - 
     * @param "SendingBusinessExpireDate" (optional.String) - 
     * @param "SendingBusinessDateOfIncorporation" (optional.String) - 
     * @param "DocumentReferenceNumber" (optional.String) - 
     * @param "SwiftCode" (optional.String) - 
     * @param "BankIfcCode" (optional.String) - 
     * @param "BankBranchName" (optional.String) - 
     * @param "BankBranchCode" (optional.String) - 
     * @param "BankAccountNumber" (optional.String) - 
     * @param "CbcBankName" (optional.String) - 
     * @param "BpiBankName" (optional.String) - 
     * @param "PartnerName" (optional.String) - 
     * @param "SourceOfIncome" (optional.String) - 
     * @param "Relationship" (optional.String) - 
     * @param "Reason" (optional.String) - 
     * @param "SenderReference" (optional.String) - 
     * @param "AdditionalInformation" (optional.String) - 
@return InlineResponse20052
*/

type RemittancesApiPostUsersWalletsFundsTransfersOverseasTypeTestOpts struct {
    AgentId optional.String
    CalculationMode optional.String
    TransactionType optional.String
    Amount optional.String
    SendCurrency optional.String
    ReceiveCurrency optional.String
    ReceiveCountry optional.String
    PaymentMode optional.String
    RoutingType optional.String
    RoutingParam optional.String
    QuotationId optional.String
    PayoutAgent optional.String
    HashId optional.String
    FirstName optional.String
    MiddleName optional.String
    LastName optional.String
    ChineseName optional.String
    ReceiveMobileNumber optional.String
    ReceiveGender optional.String
    IdNumber optional.String
    IdType optional.String
    Occupation optional.String
    Address1 optional.String
    Address2 optional.String
    BirthDate optional.String
    City optional.String
    State optional.String
    Country optional.String
    Nationality optional.String
    Zipcode optional.String
    BeneficiaryHashId optional.String
    ReceivingBusinessRegisteredName optional.String
    ReceivingBusinessTradingName optional.String
    ReceivingBusinessEmail optional.String
    ReceivingBusinessMsisdn optional.String
    ReceivingBusinessRegistrationNumber optional.String
    ReceivingBusinessTaxId optional.String
    ReceivingBusinessAddress1 optional.String
    ReceivingBusinessAddress2 optional.String
    ReceivingBusinessCity optional.String
    ReceivingBusinessState optional.String
    ReceivingBusinessCountry optional.String
    ReceivingBusinessRepresentativeFirstName optional.String
    ReceivingBusinessRepresentativeMiddleName optional.String
    ReceivingBusinessRepresentativeLastName optional.String
    ReceivingBusinessRepresentativeLastName2 optional.String
    ReceivingBusinessRepresentativeNativeName optional.String
    ReceivingBusinessIdType optional.String
    ReceivingBusinessIdNumber optional.String
    ReceivingBusinessCountryOfIssue optional.String
    ReceivingBusinessIssueDate optional.String
    ReceivingBusinessExpireDate optional.String
    ReceivingBusinessDateOfIncorporation optional.String
    SendingBusinessHashId optional.String
    SendingBusinessRegisteredName optional.String
    SendingBusinessTradingName optional.String
    SendingBusinessEmail optional.String
    SendingBusinessMsisdn optional.String
    SendingBusinessRegistrationNumber optional.String
    SendingBusinessTaxId optional.String
    SendingBusinessAddress1 optional.String
    SendingBusinessAddress2 optional.String
    SendingBusinessCity optional.String
    SendingBusinessState optional.String
    SendingBusinessZipcode optional.String
    SendingBusinessCountry optional.String
    SendingBusinessRepresentativeFirstName optional.String
    SendingBusinessRepresentativeMiddleName optional.String
    SendingBusinessRepresentativeLastName optional.String
    SendingBusinessRepresentativeLastName2 optional.String
    SendingBusinessRepresentativeNativeName optional.String
    SendingBusinessIdType optional.String
    SendingBusinessIdNumber optional.String
    SendingBusinessCountryOfIssue optional.String
    SendingBusinessIssueDate optional.String
    SendingBusinessExpireDate optional.String
    SendingBusinessDateOfIncorporation optional.String
    DocumentReferenceNumber optional.String
    SwiftCode optional.String
    BankIfcCode optional.String
    BankBranchName optional.String
    BankBranchCode optional.String
    BankAccountNumber optional.String
    CbcBankName optional.String
    BpiBankName optional.String
    PartnerName optional.String
    SourceOfIncome optional.String
    Relationship optional.String
    Reason optional.String
    SenderReference optional.String
    AdditionalInformation optional.String
}

func (a *RemittancesApiService) PostUsersWalletsFundsTransfersOverseasTypeTest(ctx context.Context, type_ string, localVarOptionals *RemittancesApiPostUsersWalletsFundsTransfersOverseasTypeTestOpts) (InlineResponse20052, *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Post")
		localVarPostBody   interface{}
		localVarFileName   string
		localVarFileBytes  []byte
		localVarReturnValue InlineResponse20052
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath + "/users/wallets/funds/transfers/overseas/{type}/test"
	localVarPath = strings.Replace(localVarPath, "{"+"type"+"}", fmt.Sprintf("%v", type_), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	// to determine the Content-Type header
	localVarHttpContentTypes := []string{"application/x-www-form-urlencoded"}

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
	if localVarOptionals != nil && localVarOptionals.AgentId.IsSet() {
		localVarFormParams.Add("agent_id", parameterToString(localVarOptionals.AgentId.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.CalculationMode.IsSet() {
		localVarFormParams.Add("calculation_mode", parameterToString(localVarOptionals.CalculationMode.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.TransactionType.IsSet() {
		localVarFormParams.Add("transaction_type", parameterToString(localVarOptionals.TransactionType.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.Amount.IsSet() {
		localVarFormParams.Add("amount", parameterToString(localVarOptionals.Amount.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.SendCurrency.IsSet() {
		localVarFormParams.Add("send_currency", parameterToString(localVarOptionals.SendCurrency.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.ReceiveCurrency.IsSet() {
		localVarFormParams.Add("receive_currency", parameterToString(localVarOptionals.ReceiveCurrency.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.ReceiveCountry.IsSet() {
		localVarFormParams.Add("receive_country", parameterToString(localVarOptionals.ReceiveCountry.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.PaymentMode.IsSet() {
		localVarFormParams.Add("payment_mode", parameterToString(localVarOptionals.PaymentMode.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.RoutingType.IsSet() {
		localVarFormParams.Add("routing_type", parameterToString(localVarOptionals.RoutingType.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.RoutingParam.IsSet() {
		localVarFormParams.Add("routing_param", parameterToString(localVarOptionals.RoutingParam.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.QuotationId.IsSet() {
		localVarFormParams.Add("quotation_id", parameterToString(localVarOptionals.QuotationId.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.PayoutAgent.IsSet() {
		localVarFormParams.Add("payout_agent", parameterToString(localVarOptionals.PayoutAgent.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.HashId.IsSet() {
		localVarFormParams.Add("hash_id", parameterToString(localVarOptionals.HashId.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.FirstName.IsSet() {
		localVarFormParams.Add("first_name", parameterToString(localVarOptionals.FirstName.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.MiddleName.IsSet() {
		localVarFormParams.Add("middle_name", parameterToString(localVarOptionals.MiddleName.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.LastName.IsSet() {
		localVarFormParams.Add("last_name", parameterToString(localVarOptionals.LastName.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.ChineseName.IsSet() {
		localVarFormParams.Add("chinese_name", parameterToString(localVarOptionals.ChineseName.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.ReceiveMobileNumber.IsSet() {
		localVarFormParams.Add("receive_mobile_number", parameterToString(localVarOptionals.ReceiveMobileNumber.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.ReceiveGender.IsSet() {
		localVarFormParams.Add("receive_gender", parameterToString(localVarOptionals.ReceiveGender.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.IdNumber.IsSet() {
		localVarFormParams.Add("id_number", parameterToString(localVarOptionals.IdNumber.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.IdType.IsSet() {
		localVarFormParams.Add("id_type", parameterToString(localVarOptionals.IdType.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.Occupation.IsSet() {
		localVarFormParams.Add("occupation", parameterToString(localVarOptionals.Occupation.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.Address1.IsSet() {
		localVarFormParams.Add("address_1", parameterToString(localVarOptionals.Address1.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.Address2.IsSet() {
		localVarFormParams.Add("address_2", parameterToString(localVarOptionals.Address2.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.BirthDate.IsSet() {
		localVarFormParams.Add("birth_date", parameterToString(localVarOptionals.BirthDate.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.City.IsSet() {
		localVarFormParams.Add("city", parameterToString(localVarOptionals.City.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.State.IsSet() {
		localVarFormParams.Add("state", parameterToString(localVarOptionals.State.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.Country.IsSet() {
		localVarFormParams.Add("country", parameterToString(localVarOptionals.Country.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.Nationality.IsSet() {
		localVarFormParams.Add("nationality", parameterToString(localVarOptionals.Nationality.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.Zipcode.IsSet() {
		localVarFormParams.Add("zipcode", parameterToString(localVarOptionals.Zipcode.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.BeneficiaryHashId.IsSet() {
		localVarFormParams.Add("beneficiary_hash_id", parameterToString(localVarOptionals.BeneficiaryHashId.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.ReceivingBusinessRegisteredName.IsSet() {
		localVarFormParams.Add("receiving_business_registered_name", parameterToString(localVarOptionals.ReceivingBusinessRegisteredName.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.ReceivingBusinessTradingName.IsSet() {
		localVarFormParams.Add("receiving_business_trading_name", parameterToString(localVarOptionals.ReceivingBusinessTradingName.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.ReceivingBusinessEmail.IsSet() {
		localVarFormParams.Add("receiving_business_email", parameterToString(localVarOptionals.ReceivingBusinessEmail.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.ReceivingBusinessMsisdn.IsSet() {
		localVarFormParams.Add("receiving_business_msisdn", parameterToString(localVarOptionals.ReceivingBusinessMsisdn.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.ReceivingBusinessRegistrationNumber.IsSet() {
		localVarFormParams.Add("receiving_business_registration_number", parameterToString(localVarOptionals.ReceivingBusinessRegistrationNumber.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.ReceivingBusinessTaxId.IsSet() {
		localVarFormParams.Add("receiving_business_tax_id", parameterToString(localVarOptionals.ReceivingBusinessTaxId.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.ReceivingBusinessAddress1.IsSet() {
		localVarFormParams.Add("receiving_business_address1", parameterToString(localVarOptionals.ReceivingBusinessAddress1.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.ReceivingBusinessAddress2.IsSet() {
		localVarFormParams.Add("receiving_business_address2", parameterToString(localVarOptionals.ReceivingBusinessAddress2.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.ReceivingBusinessCity.IsSet() {
		localVarFormParams.Add("receiving_business_city", parameterToString(localVarOptionals.ReceivingBusinessCity.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.ReceivingBusinessState.IsSet() {
		localVarFormParams.Add("receiving_business_state", parameterToString(localVarOptionals.ReceivingBusinessState.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.ReceivingBusinessCountry.IsSet() {
		localVarFormParams.Add("receiving_business_country", parameterToString(localVarOptionals.ReceivingBusinessCountry.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.ReceivingBusinessRepresentativeFirstName.IsSet() {
		localVarFormParams.Add("receiving_business_representative_first_name", parameterToString(localVarOptionals.ReceivingBusinessRepresentativeFirstName.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.ReceivingBusinessRepresentativeMiddleName.IsSet() {
		localVarFormParams.Add("receiving_business_representative_middle_name", parameterToString(localVarOptionals.ReceivingBusinessRepresentativeMiddleName.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.ReceivingBusinessRepresentativeLastName.IsSet() {
		localVarFormParams.Add("receiving_business_representative_last_name", parameterToString(localVarOptionals.ReceivingBusinessRepresentativeLastName.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.ReceivingBusinessRepresentativeLastName2.IsSet() {
		localVarFormParams.Add("receiving_business_representative_last_name2", parameterToString(localVarOptionals.ReceivingBusinessRepresentativeLastName2.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.ReceivingBusinessRepresentativeNativeName.IsSet() {
		localVarFormParams.Add("receiving_business_representative_native_name", parameterToString(localVarOptionals.ReceivingBusinessRepresentativeNativeName.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.ReceivingBusinessIdType.IsSet() {
		localVarFormParams.Add("receiving_business_id_type", parameterToString(localVarOptionals.ReceivingBusinessIdType.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.ReceivingBusinessIdNumber.IsSet() {
		localVarFormParams.Add("receiving_business_id_number", parameterToString(localVarOptionals.ReceivingBusinessIdNumber.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.ReceivingBusinessCountryOfIssue.IsSet() {
		localVarFormParams.Add("receiving_business_country_of_issue", parameterToString(localVarOptionals.ReceivingBusinessCountryOfIssue.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.ReceivingBusinessIssueDate.IsSet() {
		localVarFormParams.Add("receiving_business_issue_date", parameterToString(localVarOptionals.ReceivingBusinessIssueDate.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.ReceivingBusinessExpireDate.IsSet() {
		localVarFormParams.Add("receiving_business_expire_date", parameterToString(localVarOptionals.ReceivingBusinessExpireDate.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.ReceivingBusinessDateOfIncorporation.IsSet() {
		localVarFormParams.Add("receiving_business_date_of_incorporation", parameterToString(localVarOptionals.ReceivingBusinessDateOfIncorporation.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.SendingBusinessHashId.IsSet() {
		localVarFormParams.Add("sending_business_hash_id", parameterToString(localVarOptionals.SendingBusinessHashId.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.SendingBusinessRegisteredName.IsSet() {
		localVarFormParams.Add("sending_business_registered_name", parameterToString(localVarOptionals.SendingBusinessRegisteredName.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.SendingBusinessTradingName.IsSet() {
		localVarFormParams.Add("sending_business_trading_name", parameterToString(localVarOptionals.SendingBusinessTradingName.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.SendingBusinessEmail.IsSet() {
		localVarFormParams.Add("sending_business_email", parameterToString(localVarOptionals.SendingBusinessEmail.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.SendingBusinessMsisdn.IsSet() {
		localVarFormParams.Add("sending_business_msisdn", parameterToString(localVarOptionals.SendingBusinessMsisdn.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.SendingBusinessRegistrationNumber.IsSet() {
		localVarFormParams.Add("sending_business_registration_number", parameterToString(localVarOptionals.SendingBusinessRegistrationNumber.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.SendingBusinessTaxId.IsSet() {
		localVarFormParams.Add("sending_business_tax_id", parameterToString(localVarOptionals.SendingBusinessTaxId.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.SendingBusinessAddress1.IsSet() {
		localVarFormParams.Add("sending_business_address1", parameterToString(localVarOptionals.SendingBusinessAddress1.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.SendingBusinessAddress2.IsSet() {
		localVarFormParams.Add("sending_business_address2", parameterToString(localVarOptionals.SendingBusinessAddress2.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.SendingBusinessCity.IsSet() {
		localVarFormParams.Add("sending_business_city", parameterToString(localVarOptionals.SendingBusinessCity.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.SendingBusinessState.IsSet() {
		localVarFormParams.Add("sending_business_state", parameterToString(localVarOptionals.SendingBusinessState.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.SendingBusinessZipcode.IsSet() {
		localVarFormParams.Add("sending_business_zipcode", parameterToString(localVarOptionals.SendingBusinessZipcode.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.SendingBusinessCountry.IsSet() {
		localVarFormParams.Add("sending_business_country", parameterToString(localVarOptionals.SendingBusinessCountry.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.SendingBusinessRepresentativeFirstName.IsSet() {
		localVarFormParams.Add("sending_business_representative_first_name", parameterToString(localVarOptionals.SendingBusinessRepresentativeFirstName.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.SendingBusinessRepresentativeMiddleName.IsSet() {
		localVarFormParams.Add("sending_business_representative_middle_name", parameterToString(localVarOptionals.SendingBusinessRepresentativeMiddleName.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.SendingBusinessRepresentativeLastName.IsSet() {
		localVarFormParams.Add("sending_business_representative_last_name", parameterToString(localVarOptionals.SendingBusinessRepresentativeLastName.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.SendingBusinessRepresentativeLastName2.IsSet() {
		localVarFormParams.Add("sending_business_representative_last_name2", parameterToString(localVarOptionals.SendingBusinessRepresentativeLastName2.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.SendingBusinessRepresentativeNativeName.IsSet() {
		localVarFormParams.Add("sending_business_representative_native_name", parameterToString(localVarOptionals.SendingBusinessRepresentativeNativeName.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.SendingBusinessIdType.IsSet() {
		localVarFormParams.Add("sending_business_id_type", parameterToString(localVarOptionals.SendingBusinessIdType.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.SendingBusinessIdNumber.IsSet() {
		localVarFormParams.Add("sending_business_id_number", parameterToString(localVarOptionals.SendingBusinessIdNumber.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.SendingBusinessCountryOfIssue.IsSet() {
		localVarFormParams.Add("sending_business_country_of_issue", parameterToString(localVarOptionals.SendingBusinessCountryOfIssue.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.SendingBusinessIssueDate.IsSet() {
		localVarFormParams.Add("sending_business_issue_date", parameterToString(localVarOptionals.SendingBusinessIssueDate.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.SendingBusinessExpireDate.IsSet() {
		localVarFormParams.Add("sending_business_expire_date", parameterToString(localVarOptionals.SendingBusinessExpireDate.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.SendingBusinessDateOfIncorporation.IsSet() {
		localVarFormParams.Add("sending_business_date_of_incorporation", parameterToString(localVarOptionals.SendingBusinessDateOfIncorporation.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.DocumentReferenceNumber.IsSet() {
		localVarFormParams.Add("document_reference_number", parameterToString(localVarOptionals.DocumentReferenceNumber.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.SwiftCode.IsSet() {
		localVarFormParams.Add("swift_code", parameterToString(localVarOptionals.SwiftCode.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.BankIfcCode.IsSet() {
		localVarFormParams.Add("bank_ifc_code", parameterToString(localVarOptionals.BankIfcCode.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.BankBranchName.IsSet() {
		localVarFormParams.Add("bank_branch_name", parameterToString(localVarOptionals.BankBranchName.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.BankBranchCode.IsSet() {
		localVarFormParams.Add("bank_branch_code", parameterToString(localVarOptionals.BankBranchCode.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.BankAccountNumber.IsSet() {
		localVarFormParams.Add("bank_account_number", parameterToString(localVarOptionals.BankAccountNumber.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.CbcBankName.IsSet() {
		localVarFormParams.Add("cbc_bank_name", parameterToString(localVarOptionals.CbcBankName.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.BpiBankName.IsSet() {
		localVarFormParams.Add("bpi_bank_name", parameterToString(localVarOptionals.BpiBankName.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.PartnerName.IsSet() {
		localVarFormParams.Add("partner_name", parameterToString(localVarOptionals.PartnerName.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.SourceOfIncome.IsSet() {
		localVarFormParams.Add("source_of_income", parameterToString(localVarOptionals.SourceOfIncome.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.Relationship.IsSet() {
		localVarFormParams.Add("relationship", parameterToString(localVarOptionals.Relationship.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.Reason.IsSet() {
		localVarFormParams.Add("reason", parameterToString(localVarOptionals.Reason.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.SenderReference.IsSet() {
		localVarFormParams.Add("sender_reference", parameterToString(localVarOptionals.SenderReference.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.AdditionalInformation.IsSet() {
		localVarFormParams.Add("additional_information", parameterToString(localVarOptionals.AdditionalInformation.Value(), ""))
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
			var v InlineResponse20052
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
