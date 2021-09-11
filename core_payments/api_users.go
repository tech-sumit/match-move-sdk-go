
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

type UsersApiService service
/*
UsersApiService Create/Re-activate user
This endpoint used for creation of new user in the system. In addition when called just using &#x60;user_id&#x60; the same endpoint is used for re-activating the user in the system. If the user is in locked state.  The user information that needs to be colleted in the system depends on the country of issuance of the program. &lt;!-- theme: info --&gt; &gt; ### Note &gt; &gt; Mobile number and email combination is unique individually in the system. For a given program if mobile number is already in use and that account is in active or locked state the same mobile number cannot be used for new account creation. It can only be re-used if the original account is in blocked state. Similarly for a given program if email is already in use and the account is in active or locked state the same email cannot be used for another user. It can only be re-used if the account is in blocked state.   &lt;!-- theme: info --&gt; &gt; ### Note &gt; &gt; Do note when this endpoint is called in context of user re-activation only user_id is required.   &lt;!-- theme: danger --&gt; &gt; ### Note &gt; &gt; Blocked accounts cannot be re-activated 
 * @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param id
 * @param firstName
 * @param lastName
 * @param middleName
 * @param preferredName
 * @param email
 * @param mobileCountryCode
 * @param password
 * @param mobile
 * @param title
 * @param idType
 * @param idNumber
 * @param idDateExpiry
 * @param idDateIssue
 * @param countryOfIssue
 * @param altIdType
 * @param altIdNumber
 * @param altIdDateExpiry
 * @param altIdDateIssue
 * @param altIdCountryOfIssue
 * @param birthday
 * @param placeOfBirth
 * @param nationality
 * @param gender
 * @param maritalStatus
 * @param mothersMaidenName
 * @param natureOfBusiness
 * @param customerId
 * @param purposeOfAccount
 * @param partnerId
@return UserFullV1
*/
func (a *UsersApiService) CreateUsers(ctx context.Context, id string, firstName string, lastName string, middleName string, preferredName string, email string, mobileCountryCode string, password string, mobile string, title string, idType string, idNumber string, idDateExpiry string, idDateIssue string, countryOfIssue string, altIdType string, altIdNumber string, altIdDateExpiry string, altIdDateIssue string, altIdCountryOfIssue string, birthday string, placeOfBirth string, nationality string, gender string, maritalStatus string, mothersMaidenName string, natureOfBusiness string, customerId string, purposeOfAccount string, partnerId string) (UserFullV1, *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Post")
		localVarPostBody   interface{}
		localVarFileName   string
		localVarFileBytes  []byte
		localVarReturnValue UserFullV1
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath + "/users"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if strlen(firstName) > 50 {
		return localVarReturnValue, nil, reportError("firstName must have less than 50 elements")
	}
	if strlen(lastName) > 50 {
		return localVarReturnValue, nil, reportError("lastName must have less than 50 elements")
	}
	if strlen(middleName) > 50 {
		return localVarReturnValue, nil, reportError("middleName must have less than 50 elements")
	}
	if strlen(preferredName) > 25 {
		return localVarReturnValue, nil, reportError("preferredName must have less than 25 elements")
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
	localVarFormParams.Add("id", parameterToString(id, ""))
	localVarFormParams.Add("first_name", parameterToString(firstName, ""))
	localVarFormParams.Add("last_name", parameterToString(lastName, ""))
	localVarFormParams.Add("middle_name", parameterToString(middleName, ""))
	localVarFormParams.Add("preferred_name", parameterToString(preferredName, ""))
	localVarFormParams.Add("email", parameterToString(email, ""))
	localVarFormParams.Add("mobile_country_code", parameterToString(mobileCountryCode, ""))
	localVarFormParams.Add("password", parameterToString(password, ""))
	localVarFormParams.Add("mobile", parameterToString(mobile, ""))
	localVarFormParams.Add("title", parameterToString(title, ""))
	localVarFormParams.Add("id_type", parameterToString(idType, ""))
	localVarFormParams.Add("id_number", parameterToString(idNumber, ""))
	localVarFormParams.Add("id_date_expiry", parameterToString(idDateExpiry, ""))
	localVarFormParams.Add("id_date_issue", parameterToString(idDateIssue, ""))
	localVarFormParams.Add("country_of_issue", parameterToString(countryOfIssue, ""))
	localVarFormParams.Add("alt_id_type", parameterToString(altIdType, ""))
	localVarFormParams.Add("alt_id_number", parameterToString(altIdNumber, ""))
	localVarFormParams.Add("alt_id_date_expiry", parameterToString(altIdDateExpiry, ""))
	localVarFormParams.Add("alt_id_date_issue", parameterToString(altIdDateIssue, ""))
	localVarFormParams.Add("alt_id_country_of_issue", parameterToString(altIdCountryOfIssue, ""))
	localVarFormParams.Add("birthday", parameterToString(birthday, ""))
	localVarFormParams.Add("place_of_birth", parameterToString(placeOfBirth, ""))
	localVarFormParams.Add("nationality", parameterToString(nationality, ""))
	localVarFormParams.Add("gender", parameterToString(gender, ""))
	localVarFormParams.Add("marital_status", parameterToString(maritalStatus, ""))
	localVarFormParams.Add("mothers_maiden_name", parameterToString(mothersMaidenName, ""))
	localVarFormParams.Add("nature_of_business", parameterToString(natureOfBusiness, ""))
	localVarFormParams.Add("customer_id", parameterToString(customerId, ""))
	localVarFormParams.Add("purpose_of_account", parameterToString(purposeOfAccount, ""))
	localVarFormParams.Add("partner_id", parameterToString(partnerId, ""))
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
			var v UserFullV1
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
UsersApiService De-activate user 
There are two kinds of suspensions which are supported in the system currently. depending on the state passed as part of the request the behavior changes. ### Suspended   disable all resources of the user which includes the generation of the security token. This will also invalidate all active access associated to the user. If the card has been already pre-validated by a merchant, an authorization can still be made. ### Blocked   will suspend the account and block all cards associated to the user. The blocked cards will disable all transaction made by that card, and all purchase authorization. &lt;!-- theme: info --&gt; &gt; ### Note &gt; An account with cards or payment account with pending settlements cannot be closed. You are recommended to first suspend the account and wait for pending settlements before blocking 
 * @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param userId
 * @param optional nil or *UsersApiDeleteUserIdOpts - Optional Parameters:
     * @param "State" (optional.String) - 
@return InlineResponse200
*/

type UsersApiDeleteUserIdOpts struct {
    State optional.String
}

func (a *UsersApiService) DeleteUserId(ctx context.Context, userId string, localVarOptionals *UsersApiDeleteUserIdOpts) (InlineResponse200, *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Delete")
		localVarPostBody   interface{}
		localVarFileName   string
		localVarFileBytes  []byte
		localVarReturnValue InlineResponse200
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath + "/users/{user_id}"
	localVarPath = strings.Replace(localVarPath, "{"+"user_id"+"}", fmt.Sprintf("%v", userId), -1)

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
	if localVarOptionals != nil && localVarOptionals.State.IsSet() {
		localVarFormParams.Add("state", parameterToString(localVarOptionals.State.Value(), ""))
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
UsersApiService Get user
Retrieve details of a particular user created in the system 
 * @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param xAuthUserID MatchMove provided hash ID for the user
@return UserFullV1
*/
func (a *UsersApiService) GetUser(ctx context.Context, xAuthUserID string) (UserFullV1, *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Get")
		localVarPostBody   interface{}
		localVarFileName   string
		localVarFileBytes  []byte
		localVarReturnValue UserFullV1
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath + "/users"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if strlen(xAuthUserID) > 32 {
		return localVarReturnValue, nil, reportError("xAuthUserID must have less than 32 elements")
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
			var v UserFullV1
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
UsersApiService Update user
Update details of a user created in the system  The &#x60;preferred_name&#x60; can only be changed before wallet/card creation.  Modifying &#x60;first_name&#x60;, &#x60;middle_name&#x60;, &#x60;last_name&#x60;, &#x60;id_type&#x60;, &#x60;id_number&#x60; and/or &#x60;id_date_expiry&#x60; will result to the removal of any risk exception created for the user, thus, a new exception must be added using &lt;a href&#x3D;\&quot;#post-users-risk-exceptions\&quot;&gt;POST /users/risk/exceptions&lt;/a&gt;.  This resource may function as a validation for the given parameters, provided appending &#x60;test&#x60; to the url (eg. users/test). 
 * @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param id
 * @param firstName
 * @param lastName
 * @param middleName
 * @param preferredName
 * @param email
 * @param mobileCountryCode
 * @param password
 * @param mobile
 * @param title
 * @param idType
 * @param idNumber
 * @param idDateExpiry
 * @param idDateIssue
 * @param countryOfIssue
 * @param altIdType
 * @param altIdNumber
 * @param altIdDateExpiry
 * @param altIdDateIssue
 * @param altIdCountryOfIssue
 * @param birthday
 * @param placeOfBirth
 * @param nationality
 * @param gender
 * @param maritalStatus
 * @param mothersMaidenName
 * @param natureOfBusiness
 * @param customerId
 * @param purposeOfAccount
 * @param partnerId
 * @param xAuthUserID MatchMove provided hash ID for the user
@return UserFullV1
*/
func (a *UsersApiService) UpdateUser(ctx context.Context, id string, firstName string, lastName string, middleName string, preferredName string, email string, mobileCountryCode string, password string, mobile string, title string, idType string, idNumber string, idDateExpiry string, idDateIssue string, countryOfIssue string, altIdType string, altIdNumber string, altIdDateExpiry string, altIdDateIssue string, altIdCountryOfIssue string, birthday string, placeOfBirth string, nationality string, gender string, maritalStatus string, mothersMaidenName string, natureOfBusiness string, customerId string, purposeOfAccount string, partnerId string, xAuthUserID string) (UserFullV1, *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Put")
		localVarPostBody   interface{}
		localVarFileName   string
		localVarFileBytes  []byte
		localVarReturnValue UserFullV1
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath + "/users"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if strlen(firstName) > 50 {
		return localVarReturnValue, nil, reportError("firstName must have less than 50 elements")
	}
	if strlen(lastName) > 50 {
		return localVarReturnValue, nil, reportError("lastName must have less than 50 elements")
	}
	if strlen(middleName) > 50 {
		return localVarReturnValue, nil, reportError("middleName must have less than 50 elements")
	}
	if strlen(preferredName) > 25 {
		return localVarReturnValue, nil, reportError("preferredName must have less than 25 elements")
	}
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
	localVarFormParams.Add("id", parameterToString(id, ""))
	localVarFormParams.Add("first_name", parameterToString(firstName, ""))
	localVarFormParams.Add("last_name", parameterToString(lastName, ""))
	localVarFormParams.Add("middle_name", parameterToString(middleName, ""))
	localVarFormParams.Add("preferred_name", parameterToString(preferredName, ""))
	localVarFormParams.Add("email", parameterToString(email, ""))
	localVarFormParams.Add("mobile_country_code", parameterToString(mobileCountryCode, ""))
	localVarFormParams.Add("password", parameterToString(password, ""))
	localVarFormParams.Add("mobile", parameterToString(mobile, ""))
	localVarFormParams.Add("title", parameterToString(title, ""))
	localVarFormParams.Add("id_type", parameterToString(idType, ""))
	localVarFormParams.Add("id_number", parameterToString(idNumber, ""))
	localVarFormParams.Add("id_date_expiry", parameterToString(idDateExpiry, ""))
	localVarFormParams.Add("id_date_issue", parameterToString(idDateIssue, ""))
	localVarFormParams.Add("country_of_issue", parameterToString(countryOfIssue, ""))
	localVarFormParams.Add("alt_id_type", parameterToString(altIdType, ""))
	localVarFormParams.Add("alt_id_number", parameterToString(altIdNumber, ""))
	localVarFormParams.Add("alt_id_date_expiry", parameterToString(altIdDateExpiry, ""))
	localVarFormParams.Add("alt_id_date_issue", parameterToString(altIdDateIssue, ""))
	localVarFormParams.Add("alt_id_country_of_issue", parameterToString(altIdCountryOfIssue, ""))
	localVarFormParams.Add("birthday", parameterToString(birthday, ""))
	localVarFormParams.Add("place_of_birth", parameterToString(placeOfBirth, ""))
	localVarFormParams.Add("nationality", parameterToString(nationality, ""))
	localVarFormParams.Add("gender", parameterToString(gender, ""))
	localVarFormParams.Add("marital_status", parameterToString(maritalStatus, ""))
	localVarFormParams.Add("mothers_maiden_name", parameterToString(mothersMaidenName, ""))
	localVarFormParams.Add("nature_of_business", parameterToString(natureOfBusiness, ""))
	localVarFormParams.Add("customer_id", parameterToString(customerId, ""))
	localVarFormParams.Add("purpose_of_account", parameterToString(purposeOfAccount, ""))
	localVarFormParams.Add("partner_id", parameterToString(partnerId, ""))
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
			var v UserFullV1
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
