
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

type DefaultApiService service
/*
DefaultApiService
This resource is used for the initial process of payment for providers involving tokens from Relay to the Payment Gateway 
 * @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param xAuthHeader Consumer ACL Key
 * @param providerHash Provider Hash ID
 * @param optional nil or *DefaultApiCheckoutTokenProviderHashPostOpts - Optional Parameters:
     * @param "ConsumerHash" (optional.String) - 
     * @param "ProductCode" (optional.String) - 
     * @param "ProviderChannel" (optional.String) - 
     * @param "Amount" (optional.Float64) - 
     * @param "KycStatus" (optional.String) - 
     * @param "UserHashId" (optional.String) - 
     * @param "Email" (optional.String) - 
     * @param "MobileCountryCode" (optional.Int32) - 
     * @param "Mobile" (optional.Int32) - 
     * @param "Currency" (optional.String) - 
     * @param "Pan" (optional.String) - 
     * @param "ClientRefId" (optional.String) - 
     * @param "ActionOrigin" (optional.String) - 
@return InlineResponse200
*/

type DefaultApiCheckoutTokenProviderHashPostOpts struct {
    ConsumerHash optional.String
    ProductCode optional.String
    ProviderChannel optional.String
    Amount optional.Float64
    KycStatus optional.String
    UserHashId optional.String
    Email optional.String
    MobileCountryCode optional.Int32
    Mobile optional.Int32
    Currency optional.String
    Pan optional.String
    ClientRefId optional.String
    ActionOrigin optional.String
}

func (a *DefaultApiService) CheckoutTokenProviderHashPost(ctx context.Context, xAuthHeader string, providerHash string, localVarOptionals *DefaultApiCheckoutTokenProviderHashPostOpts) (InlineResponse200, *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Post")
		localVarPostBody   interface{}
		localVarFileName   string
		localVarFileBytes  []byte
		localVarReturnValue InlineResponse200
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath + "/checkout/token/{provider_hash}"
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
	if localVarOptionals != nil && localVarOptionals.ConsumerHash.IsSet() {
		localVarFormParams.Add("consumer_hash", parameterToString(localVarOptionals.ConsumerHash.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.ProductCode.IsSet() {
		localVarFormParams.Add("product_code", parameterToString(localVarOptionals.ProductCode.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.ProviderChannel.IsSet() {
		localVarFormParams.Add("provider_channel", parameterToString(localVarOptionals.ProviderChannel.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.Amount.IsSet() {
		localVarFormParams.Add("amount", parameterToString(localVarOptionals.Amount.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.KycStatus.IsSet() {
		localVarFormParams.Add("kyc_status", parameterToString(localVarOptionals.KycStatus.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.UserHashId.IsSet() {
		localVarFormParams.Add("user_hash_id", parameterToString(localVarOptionals.UserHashId.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.Email.IsSet() {
		localVarFormParams.Add("email", parameterToString(localVarOptionals.Email.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.MobileCountryCode.IsSet() {
		localVarFormParams.Add("mobile_country_code", parameterToString(localVarOptionals.MobileCountryCode.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.Mobile.IsSet() {
		localVarFormParams.Add("mobile", parameterToString(localVarOptionals.Mobile.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.Currency.IsSet() {
		localVarFormParams.Add("currency", parameterToString(localVarOptionals.Currency.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.Pan.IsSet() {
		localVarFormParams.Add("pan", parameterToString(localVarOptionals.Pan.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.ClientRefId.IsSet() {
		localVarFormParams.Add("client_ref_id", parameterToString(localVarOptionals.ClientRefId.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.ActionOrigin.IsSet() {
		localVarFormParams.Add("action_origin", parameterToString(localVarOptionals.ActionOrigin.Value(), ""))
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
			var v InlineResponse200
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
