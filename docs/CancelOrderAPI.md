# \CancelOrderAPI

All URIs are relative to *https://partner-api.stg-myteksi.com/grabfood-sandbox*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CancelOrder**](CancelOrderAPI.md#CancelOrder) | **Put** /partner/v1/order/cancel | Cancel an order



## CancelOrder

> CancelOrderResponse CancelOrder(ctx).ContentType(contentType).Authorization(authorization).CancelOrderRequest(cancelOrderRequest).Execute()

Cancel an order

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
	cancelOrderRequest := *grabfood.NewCancelOrderRequest("123-CYNKLPCVRN5", "1-CYNGRUNGSBCCC", grabfood.cancelCode(1001)) // CancelOrderRequest | 

	configuration := grabfood.NewConfiguration()
	apiClient := grabfood.NewAPIClient(configuration)
	resp, r, err := apiClient.CancelOrderAPI.CancelOrder(context.Background()).ContentType(contentType).Authorization(authorization).CancelOrderRequest(cancelOrderRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CancelOrderAPI.CancelOrder``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CancelOrder`: CancelOrderResponse
	fmt.Fprintf(os.Stdout, "Response from `CancelOrderAPI.CancelOrder`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCancelOrderRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **contentType** | **string** | The content type of the request body. You must use &#x60;application/json&#x60; for this header as GrabFood API currently does not support other formats. | 
 **authorization** | **string** | Specify the generated authorization token of the bearer type. | 
 **cancelOrderRequest** | [**CancelOrderRequest**](CancelOrderRequest.md) |  | 

### Return type

[**CancelOrderResponse**](CancelOrderResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

