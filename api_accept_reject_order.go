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
)


// AcceptRejectOrderAPIService AcceptRejectOrderAPI service
type AcceptRejectOrderAPIService service

type ApiAcceptRejectOrderRequest struct {
	ctx context.Context
	ApiService *AcceptRejectOrderAPIService
	authorization *string
	contentType *string
	acceptOrderRequest *AcceptOrderRequest
}

// Specify the generated authorization token of the bearer type.
func (r ApiAcceptRejectOrderRequest) Authorization(authorization string) ApiAcceptRejectOrderRequest {
	r.authorization = &authorization
	return r
}

// The content type of the request body. You must use &#x60;application/json&#x60; for this header as GrabFood API currently does not support other formats.
func (r ApiAcceptRejectOrderRequest) ContentType(contentType string) ApiAcceptRejectOrderRequest {
	r.contentType = &contentType
	return r
}

// 
func (r ApiAcceptRejectOrderRequest) AcceptOrderRequest(acceptOrderRequest AcceptOrderRequest) ApiAcceptRejectOrderRequest {
	r.acceptOrderRequest = &acceptOrderRequest
	return r
}

func (r ApiAcceptRejectOrderRequest) Execute() (*http.Response, error) {
	return r.ApiService.AcceptRejectOrderExecute(r)
}

/*
AcceptRejectOrder Manually accept/reject orders

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiAcceptRejectOrderRequest
*/
func (a *AcceptRejectOrderAPIService) AcceptRejectOrder(ctx context.Context) ApiAcceptRejectOrderRequest {
	return ApiAcceptRejectOrderRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
func (a *AcceptRejectOrderAPIService) AcceptRejectOrderExecute(r ApiAcceptRejectOrderRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "AcceptRejectOrderAPIService.AcceptRejectOrder")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/partner/v1/order/prepare"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.authorization == nil {
		return nil, reportError("authorization is required and must be specified")
	}
	if r.contentType == nil {
		return nil, reportError("contentType is required and must be specified")
	}
	if r.acceptOrderRequest == nil {
		return nil, reportError("acceptOrderRequest is required and must be specified")
	}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{"application/json"}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	parameterAddToHeaderOrQuery(localVarHeaderParams, "Authorization", r.authorization, "simple", "")
	parameterAddToHeaderOrQuery(localVarHeaderParams, "Content-Type", r.contentType, "simple", "")
	// body params
	localVarPostBody = r.acceptOrderRequest
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarHTTPResponse, err
	}

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		return localVarHTTPResponse, newErr
	}

	return localVarHTTPResponse, nil
}
