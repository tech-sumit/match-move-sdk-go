
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

type AdminApiService service
/*
AdminApiService Gets rate per consumer and provider
Will show the rate per consumer and provider
 * @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param xAuthHeader Consumer ACL Key
 * @param consumerHash Consumer Hash ID
 * @param providerHash provider Hash ID
 * @param productCode product code
 * @param optional nil or *AdminApiManageConsumerRatesConsumerHashProviderHashGetOpts - Optional Parameters:
     * @param "ProviderChannel" (optional.String) -  provider channel
@return Fees
*/

type AdminApiManageConsumerRatesConsumerHashProviderHashGetOpts struct {
    ProviderChannel optional.String
}

func (a *AdminApiService) ManageConsumerRatesConsumerHashProviderHashGet(ctx context.Context, xAuthHeader string, consumerHash string, providerHash string, productCode string, localVarOptionals *AdminApiManageConsumerRatesConsumerHashProviderHashGetOpts) (Fees, *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Get")
		localVarPostBody   interface{}
		localVarFileName   string
		localVarFileBytes  []byte
		localVarReturnValue Fees
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath + "/manage/consumer/rates/{consumer_hash}/{provider_hash}"
	localVarPath = strings.Replace(localVarPath, "{"+"consumer_hash"+"}", fmt.Sprintf("%v", consumerHash), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"provider_hash"+"}", fmt.Sprintf("%v", providerHash), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	localVarQueryParams.Add("product_code", parameterToString(productCode, ""))
	if localVarOptionals != nil && localVarOptionals.ProviderChannel.IsSet() {
		localVarQueryParams.Add("provider_channel", parameterToString(localVarOptionals.ProviderChannel.Value(), ""))
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
			var v Fees
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
AdminApiService Add rate per consumer and provider
Add rate per consumer and provider.
 * @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param xAuthHeader Consumer ACL Key
 * @param consumerHash Consumer Hash ID
 * @param providerHash provider Hash ID
 * @param optional nil or *AdminApiManageConsumerRatesConsumerHashProviderHashPostOpts - Optional Parameters:
     * @param "ProductCode" (optional.String) - 
     * @param "FeeFlat" (optional.String) - 
     * @param "FeePercentage" (optional.String) - 
     * @param "TaxFlat" (optional.String) - 
     * @param "TaxPercentage" (optional.String) - 
     * @param "ChargeFlat" (optional.String) - 
     * @param "ChargePercentage" (optional.String) - 
     * @param "ServiceFlat" (optional.String) - 
     * @param "ServicePercentage" (optional.String) - 
@return Fees
*/

type AdminApiManageConsumerRatesConsumerHashProviderHashPostOpts struct {
    ProductCode optional.String
    FeeFlat optional.String
    FeePercentage optional.String
    TaxFlat optional.String
    TaxPercentage optional.String
    ChargeFlat optional.String
    ChargePercentage optional.String
    ServiceFlat optional.String
    ServicePercentage optional.String
}

func (a *AdminApiService) ManageConsumerRatesConsumerHashProviderHashPost(ctx context.Context, xAuthHeader string, consumerHash string, providerHash string, localVarOptionals *AdminApiManageConsumerRatesConsumerHashProviderHashPostOpts) (Fees, *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Post")
		localVarPostBody   interface{}
		localVarFileName   string
		localVarFileBytes  []byte
		localVarReturnValue Fees
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath + "/manage/consumer/rates/{consumer_hash}/{provider_hash}"
	localVarPath = strings.Replace(localVarPath, "{"+"consumer_hash"+"}", fmt.Sprintf("%v", consumerHash), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"provider_hash"+"}", fmt.Sprintf("%v", providerHash), -1)

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
	localVarHeaderParams["X-Auth-Header"] = parameterToString(xAuthHeader, "")
	if localVarOptionals != nil && localVarOptionals.ProductCode.IsSet() {
		localVarFormParams.Add("product_code", parameterToString(localVarOptionals.ProductCode.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.FeeFlat.IsSet() {
		localVarFormParams.Add("fee_flat", parameterToString(localVarOptionals.FeeFlat.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.FeePercentage.IsSet() {
		localVarFormParams.Add("fee_percentage", parameterToString(localVarOptionals.FeePercentage.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.TaxFlat.IsSet() {
		localVarFormParams.Add("tax_flat", parameterToString(localVarOptionals.TaxFlat.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.TaxPercentage.IsSet() {
		localVarFormParams.Add("tax_percentage", parameterToString(localVarOptionals.TaxPercentage.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.ChargeFlat.IsSet() {
		localVarFormParams.Add("charge_flat", parameterToString(localVarOptionals.ChargeFlat.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.ChargePercentage.IsSet() {
		localVarFormParams.Add("charge_percentage", parameterToString(localVarOptionals.ChargePercentage.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.ServiceFlat.IsSet() {
		localVarFormParams.Add("service_flat", parameterToString(localVarOptionals.ServiceFlat.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.ServicePercentage.IsSet() {
		localVarFormParams.Add("service_percentage", parameterToString(localVarOptionals.ServicePercentage.Value(), ""))
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
			var v Fees
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
AdminApiService Inactivate consumer rate
Inactivates rate for specific consumer product provider
 * @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param xAuthHeader Consumer ACL Key
 * @param consumerHash Consumer Hash ID
 * @param providerHash Provider Hash ID
 * @param optional nil or *AdminApiManageConsumerRatesConsumerHashProviderHashPutOpts - Optional Parameters:
     * @param "ProductCode" (optional.String) - 
     * @param "Inactivate" (optional.String) - 
@return Fees
*/

type AdminApiManageConsumerRatesConsumerHashProviderHashPutOpts struct {
    ProductCode optional.String
    Inactivate optional.String
}

func (a *AdminApiService) ManageConsumerRatesConsumerHashProviderHashPut(ctx context.Context, xAuthHeader string, consumerHash string, providerHash string, localVarOptionals *AdminApiManageConsumerRatesConsumerHashProviderHashPutOpts) (Fees, *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Put")
		localVarPostBody   interface{}
		localVarFileName   string
		localVarFileBytes  []byte
		localVarReturnValue Fees
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath + "/manage/consumer/rates/{consumer_hash}/{provider_hash}"
	localVarPath = strings.Replace(localVarPath, "{"+"consumer_hash"+"}", fmt.Sprintf("%v", consumerHash), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"provider_hash"+"}", fmt.Sprintf("%v", providerHash), -1)

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
	localVarHeaderParams["X-Auth-Header"] = parameterToString(xAuthHeader, "")
	if localVarOptionals != nil && localVarOptionals.ProductCode.IsSet() {
		localVarFormParams.Add("product_code", parameterToString(localVarOptionals.ProductCode.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.Inactivate.IsSet() {
		localVarFormParams.Add("inactivate", parameterToString(localVarOptionals.Inactivate.Value(), ""))
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
			var v Fees
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
AdminApiService Add provider
Add provider
 * @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param xAuthHeader Consumer ACL Key
 * @param optional nil or *AdminApiManageProvidersPostOpts - Optional Parameters:
     * @param "ProviderName" (optional.String) - 
     * @param "Type_" (optional.String) - 
     * @param "ProviderMode" (optional.String) - 
     * @param "Currency" (optional.String) - 
     * @param "Algorithm" (optional.String) - 
     * @param "Description" (optional.String) - 

*/

type AdminApiManageProvidersPostOpts struct {
    ProviderName optional.String
    Type_ optional.String
    ProviderMode optional.String
    Currency optional.String
    Algorithm optional.String
    Description optional.String
}

func (a *AdminApiService) ManageProvidersPost(ctx context.Context, xAuthHeader string, localVarOptionals *AdminApiManageProvidersPostOpts) (*http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Post")
		localVarPostBody   interface{}
		localVarFileName   string
		localVarFileBytes  []byte
		
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath + "/manage/providers"

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
	localVarHeaderParams["X-Auth-Header"] = parameterToString(xAuthHeader, "")
	if localVarOptionals != nil && localVarOptionals.ProviderName.IsSet() {
		localVarFormParams.Add("provider_name", parameterToString(localVarOptionals.ProviderName.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.Type_.IsSet() {
		localVarFormParams.Add("type", parameterToString(localVarOptionals.Type_.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.ProviderMode.IsSet() {
		localVarFormParams.Add("provider_mode", parameterToString(localVarOptionals.ProviderMode.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.Currency.IsSet() {
		localVarFormParams.Add("currency", parameterToString(localVarOptionals.Currency.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.Algorithm.IsSet() {
		localVarFormParams.Add("algorithm", parameterToString(localVarOptionals.Algorithm.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.Description.IsSet() {
		localVarFormParams.Add("description", parameterToString(localVarOptionals.Description.Value(), ""))
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
