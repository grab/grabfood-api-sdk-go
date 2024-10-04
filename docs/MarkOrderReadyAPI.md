# \MarkOrderReadyAPI

All URIs are relative to *https://partner-api.stg-myteksi.com/grabfood-sandbox*

Method | HTTP request | Description
------------- | ------------- | -------------
[**MarkOrderReady**](MarkOrderReadyAPI.md#MarkOrderReady) | **Post** /partner/v1/orders/mark | Mark order as ready



## MarkOrderReady

> MarkOrderReady(ctx).ContentType(contentType).Authorization(authorization).MarkOrderRequest(markOrderRequest).Execute()

Mark order as ready

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
	authorization := "Bearer <ACCESS_TOKEN_HERE>" // string | Specify the generated authorization token of the bearer type.
	markOrderRequest := *grabfood.NewMarkOrderRequest("123-CYNKLPCVRN5", int32(123)) // MarkOrderRequest | 

	configuration := grabfood.NewConfiguration()
	apiClient := grabfood.NewAPIClient(configuration)
	r, err := apiClient.MarkOrderReadyAPI.MarkOrderReady(context.Background()).ContentType(contentType).Authorization(authorization).MarkOrderRequest(markOrderRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MarkOrderReadyAPI.MarkOrderReady``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiMarkOrderReadyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **contentType** | **string** | The content type of the request body. You must use &#x60;application/json&#x60; for this header as GrabFood API currently does not support other formats. | 
 **authorization** | **string** | Specify the generated authorization token of the bearer type. | 
 **markOrderRequest** | [**MarkOrderRequest**](MarkOrderRequest.md) |  | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

