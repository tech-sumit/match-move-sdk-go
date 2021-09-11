
package payment_relay

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

type ConsumerProvidersApiService service
/*
ConsumerProvidersApiService Reference list for the providers
This resource is used for creating reference list for the providers
 * @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param xAuthHeader Consumer ACL Key
 * @param consumerHash Consumer Hash ID
 * @param optional nil or *ConsumerProvidersApiConsumerConsumerHashProvidersGetOpts - Optional Parameters:
     * @param "FilterField" (optional.String) -  Filter Field; Allowable: [hash_id, type, currency, name]
     * @param "FilterCondition" (optional.String) -  Filter Condition; Allowable: [&#x3D;]
     * @param "FilterValue" (optional.String) -  Filter Value
     * @param "SortField" (optional.String) -  Sort Field; Allowable: [date_created]
     * @param "SortDirection" (optional.String) -  Direction For Sorting; Allowable: [asc, desc]
@return Response7
*/

type ConsumerProvidersApiConsumerConsumerHashProvidersGetOpts struct {
    FilterField optional.String
    FilterCondition optional.String
    FilterValue optional.String
    SortField optional.String
    SortDirection optional.String
}

func (a *ConsumerProvidersApiService) ConsumerConsumerHashProvidersGet(ctx context.Context, xAuthHeader string, consumerHash string, localVarOptionals *ConsumerProvidersApiConsumerConsumerHashProvidersGetOpts) (Response7, *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Get")
		localVarPostBody   interface{}
		localVarFileName   string
		localVarFileBytes  []byte
		localVarReturnValue Response7
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath + "/consumer/{consumer_hash}/providers"
	localVarPath = strings.Replace(localVarPath, "{"+"consumer_hash"+"}", fmt.Sprintf("%v", consumerHash), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if localVarOptionals != nil && localVarOptionals.FilterField.IsSet() {
		localVarQueryParams.Add("filter_field", parameterToString(localVarOptionals.FilterField.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.FilterCondition.IsSet() {
		localVarQueryParams.Add("filter_condition", parameterToString(localVarOptionals.FilterCondition.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.FilterValue.IsSet() {
		localVarQueryParams.Add("filter_value", parameterToString(localVarOptionals.FilterValue.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.SortField.IsSet() {
		localVarQueryParams.Add("sort_field", parameterToString(localVarOptionals.SortField.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.SortDirection.IsSet() {
		localVarQueryParams.Add("sort_direction", parameterToString(localVarOptionals.SortDirection.Value(), ""))
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
	localVarHeaderParams["X-Auth-Header"] = parameterToString(xAuthHeader, "")
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
			var v Response7
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
