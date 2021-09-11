
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

type CheckoutApiService service
/*
CheckoutApiService Initial Payment Processing for Direct Payment Providers
This resource is used for the initial process of payment from Relay to the Payment Gateway
 * @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param xAuthHeader Consumer ACL Key
 * @param providerHash Provider Hash ID
 * @param optional nil or *CheckoutApiCheckoutDirectProviderHashPostOpts - Optional Parameters:
     * @param "ConsumerHash" (optional.String) - 
     * @param "ProductCode" (optional.String) - 
     * @param "Amount" (optional.Float32) - 
     * @param "KycStatus" (optional.String) - 
     * @param "UserHashId" (optional.String) - 
     * @param "Email" (optional.String) - 
     * @param "MobileCountryCode" (optional.Int32) - 
     * @param "Mobile" (optional.Int32) - 
     * @param "Currency" (optional.String) - 
     * @param "CardToken" (optional.String) - 
     * @param "Pan" (optional.String) - 
     * @param "ExpiryMonth" (optional.String) - 
     * @param "ExpiryYear" (optional.String) - 
     * @param "SecurityCode" (optional.String) - 
     * @param "Source" (optional.String) - 
     * @param "BillingAddress" (optional.String) - 
     * @param "Customer" (optional.String) - 
     * @param "ClientRefId" (optional.String) - 
     * @param "ClientIp" (optional.String) - 
     * @param "ActionOrigin" (optional.String) - 
     * @param "SaveCard" (optional.String) - 
@return Response4
*/

type CheckoutApiCheckoutDirectProviderHashPostOpts struct {
    ConsumerHash optional.String
    ProductCode optional.String
    Amount optional.Float32
    KycStatus optional.String
    UserHashId optional.String
    Email optional.String
    MobileCountryCode optional.Int32
    Mobile optional.Int32
    Currency optional.String
    CardToken optional.String
    Pan optional.String
    ExpiryMonth optional.String
    ExpiryYear optional.String
    SecurityCode optional.String
    Source optional.String
    BillingAddress optional.String
    Customer optional.String
    ClientRefId optional.String
    ClientIp optional.String
    ActionOrigin optional.String
    SaveCard optional.String
}

func (a *CheckoutApiService) CheckoutDirectProviderHashPost(ctx context.Context, xAuthHeader string, providerHash string, localVarOptionals *CheckoutApiCheckoutDirectProviderHashPostOpts) (Response4, *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Post")
		localVarPostBody   interface{}
		localVarFileName   string
		localVarFileBytes  []byte
		localVarReturnValue Response4
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath + "/checkout/direct/{provider_hash}"
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
	if localVarOptionals != nil && localVarOptionals.CardToken.IsSet() {
		localVarFormParams.Add("card_token", parameterToString(localVarOptionals.CardToken.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.Pan.IsSet() {
		localVarFormParams.Add("pan", parameterToString(localVarOptionals.Pan.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.ExpiryMonth.IsSet() {
		localVarFormParams.Add("expiry_month", parameterToString(localVarOptionals.ExpiryMonth.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.ExpiryYear.IsSet() {
		localVarFormParams.Add("expiry_year", parameterToString(localVarOptionals.ExpiryYear.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.SecurityCode.IsSet() {
		localVarFormParams.Add("security_code", parameterToString(localVarOptionals.SecurityCode.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.Source.IsSet() {
		localVarFormParams.Add("source", parameterToString(localVarOptionals.Source.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.BillingAddress.IsSet() {
		localVarFormParams.Add("billing_address", parameterToString(localVarOptionals.BillingAddress.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.Customer.IsSet() {
		localVarFormParams.Add("customer", parameterToString(localVarOptionals.Customer.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.ClientRefId.IsSet() {
		localVarFormParams.Add("client_ref_id", parameterToString(localVarOptionals.ClientRefId.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.ClientIp.IsSet() {
		localVarFormParams.Add("client_ip", parameterToString(localVarOptionals.ClientIp.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.ActionOrigin.IsSet() {
		localVarFormParams.Add("action_origin", parameterToString(localVarOptionals.ActionOrigin.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.SaveCard.IsSet() {
		localVarFormParams.Add("save_card", parameterToString(localVarOptionals.SaveCard.Value(), ""))
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
			var v Response4
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
CheckoutApiService Initial Payment Processing for Hosted Providers
This resource is used for the initial process of payment from Relay to the Payment Gateway
 * @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param xAuthHeader Consumer ACL Key
 * @param providerHash Provider Hash ID
 * @param optional nil or *CheckoutApiCheckoutHostedProviderHashPostOpts - Optional Parameters:
     * @param "ConsumerHash" (optional.String) - 
     * @param "ProductCode" (optional.String) - 
     * @param "Amount" (optional.Float32) - 
     * @param "KycStatus" (optional.String) - 
     * @param "UserHashId" (optional.String) - 
     * @param "Email" (optional.String) - 
     * @param "MobileCountryCode" (optional.Int32) - 
     * @param "Mobile" (optional.Int32) - 
     * @param "PreferredName" (optional.String) - 
     * @param "InRoute" (optional.Int32) - 
     * @param "Currency" (optional.String) - 
     * @param "CountryCode" (optional.String) - 
     * @param "Pan" (optional.String) - 
     * @param "ClientRefId" (optional.String) - 
     * @param "ClientIp" (optional.String) - 
     * @param "ActionOrigin" (optional.String) - 
@return Response5
*/

type CheckoutApiCheckoutHostedProviderHashPostOpts struct {
    ConsumerHash optional.String
    ProductCode optional.String
    Amount optional.Float32
    KycStatus optional.String
    UserHashId optional.String
    Email optional.String
    MobileCountryCode optional.Int32
    Mobile optional.Int32
    PreferredName optional.String
    InRoute optional.Int32
    Currency optional.String
    CountryCode optional.String
    Pan optional.String
    ClientRefId optional.String
    ClientIp optional.String
    ActionOrigin optional.String
}

func (a *CheckoutApiService) CheckoutHostedProviderHashPost(ctx context.Context, xAuthHeader string, providerHash string, localVarOptionals *CheckoutApiCheckoutHostedProviderHashPostOpts) (Response5, *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Post")
		localVarPostBody   interface{}
		localVarFileName   string
		localVarFileBytes  []byte
		localVarReturnValue Response5
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath + "/checkout/hosted/{provider_hash}"
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
	if localVarOptionals != nil && localVarOptionals.PreferredName.IsSet() {
		localVarFormParams.Add("preferred_name", parameterToString(localVarOptionals.PreferredName.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.InRoute.IsSet() {
		localVarFormParams.Add("in_route", parameterToString(localVarOptionals.InRoute.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.Currency.IsSet() {
		localVarFormParams.Add("currency", parameterToString(localVarOptionals.Currency.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.CountryCode.IsSet() {
		localVarFormParams.Add("country_code", parameterToString(localVarOptionals.CountryCode.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.Pan.IsSet() {
		localVarFormParams.Add("pan", parameterToString(localVarOptionals.Pan.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.ClientRefId.IsSet() {
		localVarFormParams.Add("client_ref_id", parameterToString(localVarOptionals.ClientRefId.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.ClientIp.IsSet() {
		localVarFormParams.Add("client_ip", parameterToString(localVarOptionals.ClientIp.Value(), ""))
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
			var v Response5
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
CheckoutApiService Checkout Validation Rules
This resource is used for validating the rules before top-up
 * @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param xAuthHeader Consumer ACL Key
 * @param providerHash Provider Hash ID
 * @param optional nil or *CheckoutApiCheckoutValidateProviderHashPostOpts - Optional Parameters:
     * @param "ConsumerHash" (optional.String) - 
     * @param "ProductCode" (optional.String) - 
     * @param "Amount" (optional.Int32) - 
     * @param "UserHashId" (optional.String) - 
     * @param "KycStatus" (optional.String) - 
@return Response6
*/

type CheckoutApiCheckoutValidateProviderHashPostOpts struct {
    ConsumerHash optional.String
    ProductCode optional.String
    Amount optional.Int32
    UserHashId optional.String
    KycStatus optional.String
}

func (a *CheckoutApiService) CheckoutValidateProviderHashPost(ctx context.Context, xAuthHeader string, providerHash string, localVarOptionals *CheckoutApiCheckoutValidateProviderHashPostOpts) (Response6, *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Post")
		localVarPostBody   interface{}
		localVarFileName   string
		localVarFileBytes  []byte
		localVarReturnValue Response6
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath + "/checkout/validate/{provider_hash}"
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
	if localVarOptionals != nil && localVarOptionals.Amount.IsSet() {
		localVarFormParams.Add("amount", parameterToString(localVarOptionals.Amount.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.UserHashId.IsSet() {
		localVarFormParams.Add("user_hash_id", parameterToString(localVarOptionals.UserHashId.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.KycStatus.IsSet() {
		localVarFormParams.Add("kyc_status", parameterToString(localVarOptionals.KycStatus.Value(), ""))
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
			var v Response6
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
CheckoutApiService Checkout Validation Rules
This resource is used for validating the rules before top-up
 * @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param xAuthHeader Consumer ACL Key
 * @param providerHash Provider Hash ID
 * @param type_ Type to validate, [&#x60;wallet&#x60;, &#x60;card&#x60;]
 * @param optional nil or *CheckoutApiCheckoutValidateProviderHashTypePostOpts - Optional Parameters:
     * @param "ConsumerHash" (optional.String) - 
     * @param "ProductCode" (optional.String) - 
     * @param "Amount" (optional.Int32) - 
     * @param "UserHashId" (optional.String) - 
     * @param "KycStatus" (optional.String) - 
@return Response6
*/

type CheckoutApiCheckoutValidateProviderHashTypePostOpts struct {
    ConsumerHash optional.String
    ProductCode optional.String
    Amount optional.Int32
    UserHashId optional.String
    KycStatus optional.String
}

func (a *CheckoutApiService) CheckoutValidateProviderHashTypePost(ctx context.Context, xAuthHeader string, providerHash string, type_ string, localVarOptionals *CheckoutApiCheckoutValidateProviderHashTypePostOpts) (Response6, *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Post")
		localVarPostBody   interface{}
		localVarFileName   string
		localVarFileBytes  []byte
		localVarReturnValue Response6
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath + "/checkout/validate/{provider_hash}/{type}"
	localVarPath = strings.Replace(localVarPath, "{"+"provider_hash"+"}", fmt.Sprintf("%v", providerHash), -1)
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
	localVarHeaderParams["X-Auth-Header"] = parameterToString(xAuthHeader, "")
	if localVarOptionals != nil && localVarOptionals.ConsumerHash.IsSet() {
		localVarFormParams.Add("consumer_hash", parameterToString(localVarOptionals.ConsumerHash.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.ProductCode.IsSet() {
		localVarFormParams.Add("product_code", parameterToString(localVarOptionals.ProductCode.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.Amount.IsSet() {
		localVarFormParams.Add("amount", parameterToString(localVarOptionals.Amount.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.UserHashId.IsSet() {
		localVarFormParams.Add("user_hash_id", parameterToString(localVarOptionals.UserHashId.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.KycStatus.IsSet() {
		localVarFormParams.Add("kyc_status", parameterToString(localVarOptionals.KycStatus.Value(), ""))
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
			var v Response6
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
