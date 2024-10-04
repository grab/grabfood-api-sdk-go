# \EditOrderAPI

All URIs are relative to *https://partner-api.stg-myteksi.com/grabfood-sandbox*

Method | HTTP request | Description
------------- | ------------- | -------------
[**EditOrder**](EditOrderAPI.md#EditOrder) | **Put** /partner/v1/orders/{orderID} | Edit Order



## EditOrder

> EditOrder(ctx, orderID).ContentType(contentType).Authorization(authorization).EditOrderRequest(editOrderRequest).Execute()

Edit Order

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
	orderID := "orderID_example" // string | 
	editOrderRequest := *grabfood.NewEditOrderRequest("123-CYNKLPCVRN5", []grabfood.EditOrderItem{*grabfood.NewEditOrderItem("IDGFSTI000004qy1490868132306763533", "UPDATED")}) // EditOrderRequest | 

	configuration := grabfood.NewConfiguration()
	apiClient := grabfood.NewAPIClient(configuration)
	r, err := apiClient.EditOrderAPI.EditOrder(context.Background(), orderID).ContentType(contentType).Authorization(authorization).EditOrderRequest(editOrderRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `EditOrderAPI.EditOrder``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orderID** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiEditOrderRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **contentType** | **string** | The content type of the request body. You must use &#x60;application/json&#x60; for this header as GrabFood API currently does not support other formats. | 
 **authorization** | **string** | Specify the generated authorization token of the bearer type. | 

 **editOrderRequest** | [**EditOrderRequest**](EditOrderRequest.md) |  | 

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

