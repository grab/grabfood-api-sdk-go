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


// GetOauthGrabAPIService GetOauthGrabAPI service
type GetOauthGrabAPIService service

type ApiGetOauthGrabRequest struct {
	ctx context.Context
	ApiService *GetOauthGrabAPIService
	contentType *string
	grabOauthRequest *GrabOauthRequest
}

// The content type of the request body. You must use &#x60;application/json&#x60; for this header as GrabFood API currently does not support other formats.
func (r ApiGetOauthGrabRequest) ContentType(contentType string) ApiGetOauthGrabRequest {
	r.contentType = &contentType
	return r
}

// 
func (r ApiGetOauthGrabRequest) GrabOauthRequest(grabOauthRequest GrabOauthRequest) ApiGetOauthGrabRequest {
	r.grabOauthRequest = &grabOauthRequest
	return r
}

func (r ApiGetOauthGrabRequest) Execute() (*GrabOauthResponse, *http.Response, error) {
	return r.ApiService.GetOauthGrabExecute(r)
}

/*
GetOauthGrab Get Oauth access token

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiGetOauthGrabRequest
*/
func (a *GetOauthGrabAPIService) GetOauthGrab(ctx context.Context) ApiGetOauthGrabRequest {
	return ApiGetOauthGrabRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return GrabOauthResponse
func (a *GetOauthGrabAPIService) GetOauthGrabExecute(r ApiGetOauthGrabRequest) (*GrabOauthResponse, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *GrabOauthResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "GetOauthGrabAPIService.GetOauthGrab")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/grabid/v1/oauth2/token"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.contentType == nil {
		return localVarReturnValue, nil, reportError("contentType is required and must be specified")
	}
	if r.grabOauthRequest == nil {
		return localVarReturnValue, nil, reportError("grabOauthRequest is required and must be specified")
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
	// body params
	localVarPostBody = r.grabOauthRequest
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
