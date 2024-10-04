# \GetOauthGrabAPI

All URIs are relative to *https://partner-api.stg-myteksi.com/grabfood-sandbox*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetOauthGrab**](GetOauthGrabAPI.md#GetOauthGrab) | **Post** /grabid/v1/oauth2/token | Get Oauth access token



## GetOauthGrab

> GrabOauthResponse GetOauthGrab(ctx).ContentType(contentType).GrabOauthRequest(grabOauthRequest).Execute()

Get Oauth access token

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	grabfood "github.com/grab/grabfood-api-sdk-go"
)

func main() {
	contentType := "application/json" // string | The content type of the request body. You must use `application/json` for this header as GrabFood API currently does not support other formats.
	grabOauthRequest := *grabfood.NewGrabOauthRequest("<CLIENT_ID>", "<CLIENT_SECRET>", "client_credentials", "food.partner_api") // GrabOauthRequest | 

	configuration := grabfood.NewConfiguration()
	apiClient := grabfood.NewAPIClient(configuration)
	resp, r, err := apiClient.GetOauthGrabAPI.GetOauthGrab(context.Background()).ContentType(contentType).GrabOauthRequest(grabOauthRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `GetOauthGrabAPI.GetOauthGrab``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetOauthGrab`: GrabOauthResponse
	fmt.Fprintf(os.Stdout, "Response from `GetOauthGrabAPI.GetOauthGrab`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetOauthGrabRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **contentType** | **string** | The content type of the request body. You must use &#x60;application/json&#x60; for this header as GrabFood API currently does not support other formats. | 
 **grabOauthRequest** | [**GrabOauthRequest**](GrabOauthRequest.md) |  | 

### Return type

[**GrabOauthResponse**](GrabOauthResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

