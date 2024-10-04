# \UpdateDeliveryStateAPI

All URIs are relative to *https://partner-api.stg-myteksi.com/grabfood-sandbox*

Method | HTTP request | Description
------------- | ------------- | -------------
[**UpdateDeliveryState**](UpdateDeliveryStateAPI.md#UpdateDeliveryState) | **Post** /partner/v1/order/delivery | Update delivery state



## UpdateDeliveryState

> UpdateDeliveryState(ctx).ContentType(contentType).Authorization(authorization).OrderDeliveryRequest(orderDeliveryRequest).Execute()

Update delivery state

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
	orderDeliveryRequest := *grabfood.NewOrderDeliveryRequest("123-CYNKLPCVRN5", "Collect", "Delivered") // OrderDeliveryRequest | 

	configuration := grabfood.NewConfiguration()
	apiClient := grabfood.NewAPIClient(configuration)
	r, err := apiClient.UpdateDeliveryStateAPI.UpdateDeliveryState(context.Background()).ContentType(contentType).Authorization(authorization).OrderDeliveryRequest(orderDeliveryRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UpdateDeliveryStateAPI.UpdateDeliveryState``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiUpdateDeliveryStateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **contentType** | **string** | The content type of the request body. You must use &#x60;application/json&#x60; for this header as GrabFood API currently does not support other formats. | 
 **authorization** | **string** | Specify the generated authorization token of the bearer type. | 
 **orderDeliveryRequest** | [**OrderDeliveryRequest**](OrderDeliveryRequest.md) |  | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

