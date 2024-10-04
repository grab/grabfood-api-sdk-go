# \AcceptRejectOrderAPI

All URIs are relative to *https://partner-api.stg-myteksi.com/grabfood-sandbox*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AcceptRejectOrder**](AcceptRejectOrderAPI.md#AcceptRejectOrder) | **Post** /partner/v1/order/prepare | Manually accept/reject orders



## AcceptRejectOrder

> AcceptRejectOrder(ctx).Authorization(authorization).ContentType(contentType).AcceptOrderRequest(acceptOrderRequest).Execute()

Manually accept/reject orders

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
	authorization := "Bearer <ACCESS_TOKEN_HERE>" // string | Specify the generated authorization token of the bearer type.
	contentType := "application/json" // string | The content type of the request body. You must use `application/json` for this header as GrabFood API currently does not support other formats.
	acceptOrderRequest := *grabfood.NewAcceptOrderRequest("123-CYNKLPCVRN5", "ToState_example") // AcceptOrderRequest | 

	configuration := grabfood.NewConfiguration()
	apiClient := grabfood.NewAPIClient(configuration)
	r, err := apiClient.AcceptRejectOrderAPI.AcceptRejectOrder(context.Background()).Authorization(authorization).ContentType(contentType).AcceptOrderRequest(acceptOrderRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AcceptRejectOrderAPI.AcceptRejectOrder``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAcceptRejectOrderRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **authorization** | **string** | Specify the generated authorization token of the bearer type. | 
 **contentType** | **string** | The content type of the request body. You must use &#x60;application/json&#x60; for this header as GrabFood API currently does not support other formats. | 
 **acceptOrderRequest** | [**AcceptOrderRequest**](AcceptOrderRequest.md) |  | 

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

