// Copyright 2024 Grabtaxi Holdings PTE LTE (GRAB), All rights reserved.
// Use of this source code is governed by an MIT-style license that can be found in the LICENSE file

/*
GrabFood

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: 1.1.3
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package grabfood

import (
	"bytes"
	"context"
	"io"
	"net/http"
	"net/url"
	"strings"
)


// UpdateStoreSpecialHourAPIService UpdateStoreSpecialHourAPI service
type UpdateStoreSpecialHourAPIService service

type ApiUpdateStoreSpecialHourRequest struct {
	ctx context.Context
	ApiService *UpdateStoreSpecialHourAPIService
	contentType *string
	authorization *string
	merchantID string
	updateSpecialHourRequest *UpdateSpecialHourRequest
}

// The content type of the request body. You must use &#x60;application/json&#x60; for this header as GrabFood API currently does not support other formats.
func (r ApiUpdateStoreSpecialHourRequest) ContentType(contentType string) ApiUpdateStoreSpecialHourRequest {
	r.contentType = &contentType
	return r
}

// Specify the generated authorization token of the bearer type.
func (r ApiUpdateStoreSpecialHourRequest) Authorization(authorization string) ApiUpdateStoreSpecialHourRequest {
	r.authorization = &authorization
	return r
}

func (r ApiUpdateStoreSpecialHourRequest) UpdateSpecialHourRequest(updateSpecialHourRequest UpdateSpecialHourRequest) ApiUpdateStoreSpecialHourRequest {
	r.updateSpecialHourRequest = &updateSpecialHourRequest
	return r
}

func (r ApiUpdateStoreSpecialHourRequest) Execute() (*UpdateSpecialHourResponse, *http.Response, error) {
	return r.ApiService.UpdateStoreSpecialHourExecute(r)
}

/*
UpdateStoreSpecialHour Update Store Special Hours

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param merchantID The merchant's ID that is in GrabFood's database.
 @return ApiUpdateStoreSpecialHourRequest
*/
func (a *UpdateStoreSpecialHourAPIService) UpdateStoreSpecialHour(ctx context.Context, merchantID string) ApiUpdateStoreSpecialHourRequest {
	return ApiUpdateStoreSpecialHourRequest{
		ApiService: a,
		ctx: ctx,
		merchantID: merchantID,
	}
}

// Execute executes the request
//  @return UpdateSpecialHourResponse
func (a *UpdateStoreSpecialHourAPIService) UpdateStoreSpecialHourExecute(r ApiUpdateStoreSpecialHourRequest) (*UpdateSpecialHourResponse, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPut
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *UpdateSpecialHourResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "UpdateStoreSpecialHourAPIService.UpdateStoreSpecialHour")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/partner/v2/merchants/{merchantID}/store/special-opening-hour"
	localVarPath = strings.Replace(localVarPath, "{"+"merchantID"+"}", url.PathEscape(parameterValueToString(r.merchantID, "merchantID")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.contentType == nil {
		return localVarReturnValue, nil, reportError("contentType is required and must be specified")
	}
	if r.authorization == nil {
		return localVarReturnValue, nil, reportError("authorization is required and must be specified")
	}
	if r.updateSpecialHourRequest == nil {
		return localVarReturnValue, nil, reportError("updateSpecialHourRequest is required and must be specified")
	}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{"application/json"}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	parameterAddToHeaderOrQuery(localVarHeaderParams, "Content-Type", r.contentType, "simple", "")
	parameterAddToHeaderOrQuery(localVarHeaderParams, "Authorization", r.authorization, "simple", "")
	// body params
	localVarPostBody = r.updateSpecialHourRequest
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}
