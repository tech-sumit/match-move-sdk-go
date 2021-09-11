
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

type CardsTransactionalLimitsApiService service
/*
CardsTransactionalLimitsApiService Delete rules for a particular transaction category
 * @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param categoryHashId
@return InlineResponse2007
*/
func (a *CardsTransactionalLimitsApiService) DeleteOauthConsumerTransactionCategoriesCategoryHashIdRule(ctx context.Context, categoryHashId string) (InlineResponse2007, *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Delete")
		localVarPostBody   interface{}
		localVarFileName   string
		localVarFileBytes  []byte
		localVarReturnValue InlineResponse2007
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath + "/oauth/consumer/transaction/categories/{category_hash_id}/rule"
	localVarPath = strings.Replace(localVarPath, "{"+"category_hash_id"+"}", fmt.Sprintf("%v", categoryHashId), -1)

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
			var v InlineResponse2007
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
		return localVarReturnValue, localVarHttpResponse, newErr
	}

	return localVarReturnValue, localVarHttpResponse, nil
}
/*
CardsTransactionalLimitsApiService Disable transactional limits for a particular card and category
Disable transactional limits for a particular card and category 
 * @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param cardHashId
 * @param categoryHashId Transaction category Identifier
 * @param xAuthUserID User Identifier
@return UserLimitV1
*/
func (a *CardsTransactionalLimitsApiService) DeleteUsersWalletsCardsCardHashIdTransactionCategoryHashIdLimits(ctx context.Context, cardHashId string, categoryHashId string, xAuthUserID string) (UserLimitV1, *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Delete")
		localVarPostBody   interface{}
		localVarFileName   string
		localVarFileBytes  []byte
		localVarReturnValue UserLimitV1
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath + "/users/wallets/cards/{card_hash_id}/transactions/categories/{category_hash_id}/limit"
	localVarPath = strings.Replace(localVarPath, "{"+"card_hash_id"+"}", fmt.Sprintf("%v", cardHashId), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"category_hash_id"+"}", fmt.Sprintf("%v", categoryHashId), -1)

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
	localVarHeaderParams["X-Auth-User-ID"] = parameterToString(xAuthUserID, "")
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
			var v UserLimitV1
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
		return localVarReturnValue, localVarHttpResponse, newErr
	}

	return localVarReturnValue, localVarHttpResponse, nil
}
/*
CardsTransactionalLimitsApiService Retrieve rules for a particular transaction category
 * @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param categoryHashId
@return TransactionCategoriesRulesResponseV1
*/
func (a *CardsTransactionalLimitsApiService) GetOauthConsumerTransactionCategoriesCategoryHashIdRule(ctx context.Context, categoryHashId string) (TransactionCategoriesRulesResponseV1, *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Get")
		localVarPostBody   interface{}
		localVarFileName   string
		localVarFileBytes  []byte
		localVarReturnValue TransactionCategoriesRulesResponseV1
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath + "/oauth/consumer/transaction/categories/{category_hash_id}/rule"
	localVarPath = strings.Replace(localVarPath, "{"+"category_hash_id"+"}", fmt.Sprintf("%v", categoryHashId), -1)

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
			var v TransactionCategoriesRulesResponseV1
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
		return localVarReturnValue, localVarHttpResponse, newErr
	}

	return localVarReturnValue, localVarHttpResponse, nil
}
/*
CardsTransactionalLimitsApiService Retrieve transactional limits for categories
Retrieve transactional limits for a card for a particular category
 * @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param cardHashId
 * @param categoryHashId Transaction category Identifier
 * @param xAuthUserID User Identifier
@return UserLimitV1
*/
func (a *CardsTransactionalLimitsApiService) GetUsersWalletsCardsCardHashIdTransactionCategoryHashIdLimits(ctx context.Context, cardHashId string, categoryHashId string, xAuthUserID string) (UserLimitV1, *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Get")
		localVarPostBody   interface{}
		localVarFileName   string
		localVarFileBytes  []byte
		localVarReturnValue UserLimitV1
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath + "/users/wallets/cards/{card_hash_id}/transactions/categories/{category_hash_id}/limit"
	localVarPath = strings.Replace(localVarPath, "{"+"card_hash_id"+"}", fmt.Sprintf("%v", cardHashId), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"category_hash_id"+"}", fmt.Sprintf("%v", categoryHashId), -1)

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
	localVarHeaderParams["X-Auth-User-ID"] = parameterToString(xAuthUserID, "")
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
			var v UserLimitV1
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
		return localVarReturnValue, localVarHttpResponse, newErr
	}

	return localVarReturnValue, localVarHttpResponse, nil
}
/*
CardsTransactionalLimitsApiService Retrieve transactional limits for a card
Retrieve transactional limits for a card categories
 * @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param cardHashId
@return InlineResponse20053
*/
func (a *CardsTransactionalLimitsApiService) GetUsersWalletsCardsCardHashIdTransactionsCategoriesLimit(ctx context.Context, cardHashId string) (InlineResponse20053, *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Get")
		localVarPostBody   interface{}
		localVarFileName   string
		localVarFileBytes  []byte
		localVarReturnValue InlineResponse20053
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath + "/users/wallets/cards/{card_hash_id}/transactions/categories/limit"
	localVarPath = strings.Replace(localVarPath, "{"+"card_hash_id"+"}", fmt.Sprintf("%v", cardHashId), -1)

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
			var v InlineResponse20053
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
		return localVarReturnValue, localVarHttpResponse, newErr
	}

	return localVarReturnValue, localVarHttpResponse, nil
}
/*
CardsTransactionalLimitsApiService Create rules for a particular transaction category
 * @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param categoryHashId
 * @param optional nil or *CardsTransactionalLimitsApiPostOauthConsumerTransactionCategoriesCategoryHashIdRuleOpts - Optional Parameters:
     * @param "EventCodes" (optional.String) - 
     * @param "Channel" (optional.String) - 
     * @param "AcquiringCountryCode" (optional.String) - 
     * @param "Currency" (optional.String) - 
     * @param "Network" (optional.String) - 
     * @param "TerminalId" (optional.String) - 
     * @param "MerchantId" (optional.String) - 
     * @param "PosEntryModeValue" (optional.String) - 
     * @param "Pp" (optional.String) - 
     * @param "MerchantCategoryCode" (optional.String) - 
@return TransactionCategoriesRulesResponseV1
*/

type CardsTransactionalLimitsApiPostOauthConsumerTransactionCategoriesCategoryHashIdRuleOpts struct {
    EventCodes optional.String
    Channel optional.String
    AcquiringCountryCode optional.String
    Currency optional.String
    Network optional.String
    TerminalId optional.String
    MerchantId optional.String
    PosEntryModeValue optional.String
    Pp optional.String
    MerchantCategoryCode optional.String
}

func (a *CardsTransactionalLimitsApiService) PostOauthConsumerTransactionCategoriesCategoryHashIdRule(ctx context.Context, categoryHashId string, localVarOptionals *CardsTransactionalLimitsApiPostOauthConsumerTransactionCategoriesCategoryHashIdRuleOpts) (TransactionCategoriesRulesResponseV1, *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Post")
		localVarPostBody   interface{}
		localVarFileName   string
		localVarFileBytes  []byte
		localVarReturnValue TransactionCategoriesRulesResponseV1
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath + "/oauth/consumer/transaction/categories/{category_hash_id}/rule"
	localVarPath = strings.Replace(localVarPath, "{"+"category_hash_id"+"}", fmt.Sprintf("%v", categoryHashId), -1)

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
	if localVarOptionals != nil && localVarOptionals.EventCodes.IsSet() {
		localVarFormParams.Add("event_codes", parameterToString(localVarOptionals.EventCodes.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.Channel.IsSet() {
		localVarFormParams.Add("channel", parameterToString(localVarOptionals.Channel.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.AcquiringCountryCode.IsSet() {
		localVarFormParams.Add("acquiring_country_code", parameterToString(localVarOptionals.AcquiringCountryCode.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.Currency.IsSet() {
		localVarFormParams.Add("currency", parameterToString(localVarOptionals.Currency.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.Network.IsSet() {
		localVarFormParams.Add("network", parameterToString(localVarOptionals.Network.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.TerminalId.IsSet() {
		localVarFormParams.Add("terminal_id", parameterToString(localVarOptionals.TerminalId.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.MerchantId.IsSet() {
		localVarFormParams.Add("merchant_id", parameterToString(localVarOptionals.MerchantId.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.PosEntryModeValue.IsSet() {
		localVarFormParams.Add("pos_entry_mode_value", parameterToString(localVarOptionals.PosEntryModeValue.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.Pp.IsSet() {
		localVarFormParams.Add("pp", parameterToString(localVarOptionals.Pp.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.MerchantCategoryCode.IsSet() {
		localVarFormParams.Add("merchant_category_code", parameterToString(localVarOptionals.MerchantCategoryCode.Value(), ""))
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
			var v TransactionCategoriesRulesResponseV1
			err = a.client.decode(&v, localVarBody, localVarHttpResponse.Header.Get("Content-Type"));
				if err != nil {
					newErr.error = err.Error()
					return localVarReturnValue, localVarHttpResponse, newErr
				}
				newErr.model = v
				return localVarReturnValue, localVarHttpResponse, newErr
		}
		if localVarHttpResponse.StatusCode == 400 {
			var v TransactionCategoriesRulesResponseV1
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
		return localVarReturnValue, localVarHttpResponse, newErr
	}

	return localVarReturnValue, localVarHttpResponse, nil
}
/*
CardsTransactionalLimitsApiService Configure transactional limits for a particular card and category
To set user limits for card
 * @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param xAuthUserID User Identifier
 * @param cardHashId
 * @param categoryHashId Transaction category Identifier
 * @param optional nil or *CardsTransactionalLimitsApiPostUsersWalletsCardsCardHashIdTransactionCategoryHashIdLimitsOpts - Optional Parameters:
     * @param "Enable" (optional.Bool) - 
     * @param "Values" (optional.Interface of []UserLimitV1CardLimits) - 
@return UserLimitV1
*/

type CardsTransactionalLimitsApiPostUsersWalletsCardsCardHashIdTransactionCategoryHashIdLimitsOpts struct {
    Enable optional.Bool
    Values optional.Interface
}

func (a *CardsTransactionalLimitsApiService) PostUsersWalletsCardsCardHashIdTransactionCategoryHashIdLimits(ctx context.Context, xAuthUserID string, cardHashId string, categoryHashId string, localVarOptionals *CardsTransactionalLimitsApiPostUsersWalletsCardsCardHashIdTransactionCategoryHashIdLimitsOpts) (UserLimitV1, *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Post")
		localVarPostBody   interface{}
		localVarFileName   string
		localVarFileBytes  []byte
		localVarReturnValue UserLimitV1
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath + "/users/wallets/cards/{card_hash_id}/transactions/categories/{category_hash_id}/limit"
	localVarPath = strings.Replace(localVarPath, "{"+"card_hash_id"+"}", fmt.Sprintf("%v", cardHashId), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"category_hash_id"+"}", fmt.Sprintf("%v", categoryHashId), -1)

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
	localVarHeaderParams["X-Auth-User-ID"] = parameterToString(xAuthUserID, "")
	if localVarOptionals != nil && localVarOptionals.Enable.IsSet() {
		localVarFormParams.Add("enable", parameterToString(localVarOptionals.Enable.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.Values.IsSet() {
		localVarFormParams.Add("values", parameterToString(localVarOptionals.Values.Value(), "multi"))
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
			var v UserLimitV1
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
		return localVarReturnValue, localVarHttpResponse, newErr
	}

	return localVarReturnValue, localVarHttpResponse, nil
}
/*
CardsTransactionalLimitsApiService Create Transaction Category Limits
Create a user wallet card transaction category limit 
 * @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param optional nil or *CardsTransactionalLimitsApiPostUsersWalletsCardsTransactionsCategoriesLimitsDetailOpts - Optional Parameters:
     * @param "CardHashId" (optional.String) - 
     * @param "UserHashId" (optional.String) - 
     * @param "TransactionAmount" (optional.Float64) - 
     * @param "Amount" (optional.Float64) - 
     * @param "CardCurrency" (optional.String) - 
     * @param "TransactionCurrency" (optional.String) - 
     * @param "Channel" (optional.String) - 
     * @param "PosEntryModeValue" (optional.String) - 
     * @param "AcquiringCountryCode" (optional.String) - 
     * @param "Network" (optional.String) - 
     * @param "MerchantCategoryCode" (optional.String) - 
     * @param "MerchantId" (optional.String) - 
     * @param "TerminalId" (optional.String) - 

*/

type CardsTransactionalLimitsApiPostUsersWalletsCardsTransactionsCategoriesLimitsDetailOpts struct {
    CardHashId optional.String
    UserHashId optional.String
    TransactionAmount optional.Float64
    Amount optional.Float64
    CardCurrency optional.String
    TransactionCurrency optional.String
    Channel optional.String
    PosEntryModeValue optional.String
    AcquiringCountryCode optional.String
    Network optional.String
    MerchantCategoryCode optional.String
    MerchantId optional.String
    TerminalId optional.String
}

func (a *CardsTransactionalLimitsApiService) PostUsersWalletsCardsTransactionsCategoriesLimitsDetail(ctx context.Context, localVarOptionals *CardsTransactionalLimitsApiPostUsersWalletsCardsTransactionsCategoriesLimitsDetailOpts) (*http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Post")
		localVarPostBody   interface{}
		localVarFileName   string
		localVarFileBytes  []byte
		
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath + "/users/wallets/cards/transactions/categories/limits/detail"

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
	localVarHttpHeaderAccepts := []string{}

	// set Accept header
	localVarHttpHeaderAccept := selectHeaderAccept(localVarHttpHeaderAccepts)
	if localVarHttpHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHttpHeaderAccept
	}
	if localVarOptionals != nil && localVarOptionals.CardHashId.IsSet() {
		localVarFormParams.Add("card_hash_id", parameterToString(localVarOptionals.CardHashId.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.UserHashId.IsSet() {
		localVarFormParams.Add("user_hash_id", parameterToString(localVarOptionals.UserHashId.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.TransactionAmount.IsSet() {
		localVarFormParams.Add("transaction_amount", parameterToString(localVarOptionals.TransactionAmount.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.Amount.IsSet() {
		localVarFormParams.Add("amount", parameterToString(localVarOptionals.Amount.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.CardCurrency.IsSet() {
		localVarFormParams.Add("card_currency", parameterToString(localVarOptionals.CardCurrency.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.TransactionCurrency.IsSet() {
		localVarFormParams.Add("transaction_currency", parameterToString(localVarOptionals.TransactionCurrency.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.Channel.IsSet() {
		localVarFormParams.Add("channel", parameterToString(localVarOptionals.Channel.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.PosEntryModeValue.IsSet() {
		localVarFormParams.Add("pos_entry_mode_value", parameterToString(localVarOptionals.PosEntryModeValue.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.AcquiringCountryCode.IsSet() {
		localVarFormParams.Add("acquiring_country_code", parameterToString(localVarOptionals.AcquiringCountryCode.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.Network.IsSet() {
		localVarFormParams.Add("network", parameterToString(localVarOptionals.Network.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.MerchantCategoryCode.IsSet() {
		localVarFormParams.Add("merchant_category_code", parameterToString(localVarOptionals.MerchantCategoryCode.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.MerchantId.IsSet() {
		localVarFormParams.Add("merchant_id", parameterToString(localVarOptionals.MerchantId.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.TerminalId.IsSet() {
		localVarFormParams.Add("terminal_id", parameterToString(localVarOptionals.TerminalId.Value(), ""))
	}
	r, err := a.client.prepareRequest(ctx, localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFileName, localVarFileBytes)
	if err != nil {
		return nil, err
	}

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return localVarHttpResponse, err
	}

	localVarBody, err := ioutil.ReadAll(localVarHttpResponse.Body)
	localVarHttpResponse.Body.Close()
	if err != nil {
		return localVarHttpResponse, err
	}


	if localVarHttpResponse.StatusCode >= 300 {
		newErr := GenericSwaggerError{
			body: localVarBody,
			error: localVarHttpResponse.Status,
		}
		return localVarHttpResponse, newErr
	}

	return localVarHttpResponse, nil
}
/*
CardsTransactionalLimitsApiService Update rules for a particular transaction category
 * @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param categoryHashId
 * @param optional nil or *CardsTransactionalLimitsApiPutOauthConsumerTransactionCategoriesCategoryHashIdRuleOpts - Optional Parameters:
     * @param "EventCodes" (optional.String) - 
     * @param "Channel" (optional.String) - 
     * @param "AcquiringCountryCode" (optional.String) - 
     * @param "Currency" (optional.String) - 
     * @param "Network" (optional.String) - 
     * @param "TerminalId" (optional.String) - 
     * @param "MerchantId" (optional.String) - 
     * @param "PosEntryModeValue" (optional.String) - 
     * @param "Pp" (optional.String) - 
     * @param "MerchantCategoryCode" (optional.String) - 
@return TransactionCategoriesRulesResponseV1
*/

type CardsTransactionalLimitsApiPutOauthConsumerTransactionCategoriesCategoryHashIdRuleOpts struct {
    EventCodes optional.String
    Channel optional.String
    AcquiringCountryCode optional.String
    Currency optional.String
    Network optional.String
    TerminalId optional.String
    MerchantId optional.String
    PosEntryModeValue optional.String
    Pp optional.String
    MerchantCategoryCode optional.String
}

func (a *CardsTransactionalLimitsApiService) PutOauthConsumerTransactionCategoriesCategoryHashIdRule(ctx context.Context, categoryHashId string, localVarOptionals *CardsTransactionalLimitsApiPutOauthConsumerTransactionCategoriesCategoryHashIdRuleOpts) (TransactionCategoriesRulesResponseV1, *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Put")
		localVarPostBody   interface{}
		localVarFileName   string
		localVarFileBytes  []byte
		localVarReturnValue TransactionCategoriesRulesResponseV1
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath + "/oauth/consumer/transaction/categories/{category_hash_id}/rule"
	localVarPath = strings.Replace(localVarPath, "{"+"category_hash_id"+"}", fmt.Sprintf("%v", categoryHashId), -1)

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
	if localVarOptionals != nil && localVarOptionals.EventCodes.IsSet() {
		localVarFormParams.Add("event_codes", parameterToString(localVarOptionals.EventCodes.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.Channel.IsSet() {
		localVarFormParams.Add("channel", parameterToString(localVarOptionals.Channel.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.AcquiringCountryCode.IsSet() {
		localVarFormParams.Add("acquiring_country_code", parameterToString(localVarOptionals.AcquiringCountryCode.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.Currency.IsSet() {
		localVarFormParams.Add("currency", parameterToString(localVarOptionals.Currency.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.Network.IsSet() {
		localVarFormParams.Add("network", parameterToString(localVarOptionals.Network.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.TerminalId.IsSet() {
		localVarFormParams.Add("terminal_id", parameterToString(localVarOptionals.TerminalId.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.MerchantId.IsSet() {
		localVarFormParams.Add("merchant_id", parameterToString(localVarOptionals.MerchantId.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.PosEntryModeValue.IsSet() {
		localVarFormParams.Add("pos_entry_mode_value", parameterToString(localVarOptionals.PosEntryModeValue.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.Pp.IsSet() {
		localVarFormParams.Add("pp", parameterToString(localVarOptionals.Pp.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.MerchantCategoryCode.IsSet() {
		localVarFormParams.Add("merchant_category_code", parameterToString(localVarOptionals.MerchantCategoryCode.Value(), ""))
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
			var v TransactionCategoriesRulesResponseV1
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
		return localVarReturnValue, localVarHttpResponse, newErr
	}

	return localVarReturnValue, localVarHttpResponse, nil
}
/*
CardsTransactionalLimitsApiService Update transactional limits for a particular card and category
To set user limits for card
 * @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param xAuthUserID User Identifier
 * @param cardHashId
 * @param categoryHashId Transaction category Identifier
 * @param optional nil or *CardsTransactionalLimitsApiPutUsersWalletsCardsCardHashIdTransactionCategoryHashIdLimitsOpts - Optional Parameters:
     * @param "Enable" (optional.Bool) - 
     * @param "Values" (optional.Interface of []UserLimitV1CardLimits) - 
@return UserLimitV1
*/

type CardsTransactionalLimitsApiPutUsersWalletsCardsCardHashIdTransactionCategoryHashIdLimitsOpts struct {
    Enable optional.Bool
    Values optional.Interface
}

func (a *CardsTransactionalLimitsApiService) PutUsersWalletsCardsCardHashIdTransactionCategoryHashIdLimits(ctx context.Context, xAuthUserID string, cardHashId string, categoryHashId string, localVarOptionals *CardsTransactionalLimitsApiPutUsersWalletsCardsCardHashIdTransactionCategoryHashIdLimitsOpts) (UserLimitV1, *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Put")
		localVarPostBody   interface{}
		localVarFileName   string
		localVarFileBytes  []byte
		localVarReturnValue UserLimitV1
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath + "/users/wallets/cards/{card_hash_id}/transactions/categories/{category_hash_id}/limit"
	localVarPath = strings.Replace(localVarPath, "{"+"card_hash_id"+"}", fmt.Sprintf("%v", cardHashId), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"category_hash_id"+"}", fmt.Sprintf("%v", categoryHashId), -1)

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
	localVarHeaderParams["X-Auth-User-ID"] = parameterToString(xAuthUserID, "")
	if localVarOptionals != nil && localVarOptionals.Enable.IsSet() {
		localVarFormParams.Add("enable", parameterToString(localVarOptionals.Enable.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.Values.IsSet() {
		localVarFormParams.Add("values", parameterToString(localVarOptionals.Values.Value(), "multi"))
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
			var v UserLimitV1
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
		return localVarReturnValue, localVarHttpResponse, newErr
	}

	return localVarReturnValue, localVarHttpResponse, nil
}
