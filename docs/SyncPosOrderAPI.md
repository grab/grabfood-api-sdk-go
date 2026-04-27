# \SyncPosOrderAPI

All URIs are relative to *https://partner-api.grab.com/grabfood-sandbox*

Method | HTTP request | Description
------------- | ------------- | -------------
[**SyncPosOrder**](SyncPosOrderAPI.md#SyncPosOrder) | **Post** /partner/v1/pos/order | Sync POS order



## SyncPosOrder

> SyncPOSOrderResponse SyncPosOrder(ctx).Authorization(authorization).ContentType(contentType).SyncPOSOrderRequest(syncPOSOrderRequest).Execute()

Sync POS order

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
	syncPOSOrderRequest := *grabfood.NewSyncPOSOrderRequest("BILL_GENERATED", *grabfood.NewPosOrder()) // SyncPOSOrderRequest | 

	configuration := grabfood.NewConfiguration()
	apiClient := grabfood.NewAPIClient(configuration)
	resp, r, err := apiClient.SyncPosOrderAPI.SyncPosOrder(context.Background()).Authorization(authorization).ContentType(contentType).SyncPOSOrderRequest(syncPOSOrderRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SyncPosOrderAPI.SyncPosOrder``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SyncPosOrder`: SyncPOSOrderResponse
	fmt.Fprintf(os.Stdout, "Response from `SyncPosOrderAPI.SyncPosOrder`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSyncPosOrderRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **authorization** | **string** | Specify the generated authorization token of the bearer type. | 
 **contentType** | **string** | The content type of the request body. You must use &#x60;application/json&#x60; for this header as GrabFood API currently does not support other formats. | 
 **syncPOSOrderRequest** | [**SyncPOSOrderRequest**](SyncPOSOrderRequest.md) |  | 

### Return type

[**SyncPOSOrderResponse**](SyncPOSOrderResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

