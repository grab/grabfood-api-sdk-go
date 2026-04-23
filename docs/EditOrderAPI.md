# \EditOrderAPI

All URIs are relative to *https://partner-api.grab.com/grabfood-sandbox*

Method | HTTP request | Description
------------- | ------------- | -------------
[**EditOrderV1**](EditOrderAPI.md#EditOrderV1) | **Put** /partner/v1/orders/{orderID} | Edit Order V1
[**EditOrderV2**](EditOrderAPI.md#EditOrderV2) | **Put** /partner/v2/orders/{orderID} | Edit Order V2



## EditOrderV1

> EditOrderV1(ctx, orderID).ContentType(contentType).Authorization(authorization).EditOrderRequest(editOrderRequest).Execute()

Edit Order V1

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
	editOrderRequest := *grabfood.NewEditOrderRequest("123-CYNKLPCVRN5", []grabfood.EditOrderItem{*grabfood.NewEditOrderItem("IDGFSTI000004qy1490868132306763533#0", "UPDATED")}) // EditOrderRequest | 

	configuration := grabfood.NewConfiguration()
	apiClient := grabfood.NewAPIClient(configuration)
	r, err := apiClient.EditOrderAPI.EditOrderV1(context.Background(), orderID).ContentType(contentType).Authorization(authorization).EditOrderRequest(editOrderRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `EditOrderAPI.EditOrderV1``: %v\n", err)
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

Other parameters are passed through a pointer to a apiEditOrderV1Request struct via the builder pattern


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


## EditOrderV2

> EditOrderV2Response EditOrderV2(ctx, orderID).ContentType(contentType).Authorization(authorization).EditOrderRequest(editOrderRequest).Execute()

Edit Order V2

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
	editOrderRequest := *grabfood.NewEditOrderRequest("123-CYNKLPCVRN5", []grabfood.EditOrderItem{*grabfood.NewEditOrderItem("IDGFSTI000004qy1490868132306763533#0", "UPDATED")}) // EditOrderRequest | 

	configuration := grabfood.NewConfiguration()
	apiClient := grabfood.NewAPIClient(configuration)
	resp, r, err := apiClient.EditOrderAPI.EditOrderV2(context.Background(), orderID).ContentType(contentType).Authorization(authorization).EditOrderRequest(editOrderRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `EditOrderAPI.EditOrderV2``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `EditOrderV2`: EditOrderV2Response
	fmt.Fprintf(os.Stdout, "Response from `EditOrderAPI.EditOrderV2`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orderID** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiEditOrderV2Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **contentType** | **string** | The content type of the request body. You must use &#x60;application/json&#x60; for this header as GrabFood API currently does not support other formats. | 
 **authorization** | **string** | Specify the generated authorization token of the bearer type. | 

 **editOrderRequest** | [**EditOrderRequest**](EditOrderRequest.md) |  | 

### Return type

[**EditOrderV2Response**](EditOrderV2Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

