# \CheckOrderCancelableAPI

All URIs are relative to *https://partner-api.stg-myteksi.com/grabfood-sandbox*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CheckOrderCancelable**](CheckOrderCancelableAPI.md#CheckOrderCancelable) | **Get** /partner/v1/order/cancelable | Check order cancelable



## CheckOrderCancelable

> CheckOrderCancelableResponse CheckOrderCancelable(ctx).Authorization(authorization).OrderID(orderID).MerchantID(merchantID).Execute()

Check order cancelable

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
	orderID := "orderID_example" // string | 
	merchantID := "1-CYNGRUNGSBCCC" // string | The merchant's ID that is in GrabFood's database.

	configuration := grabfood.NewConfiguration()
	apiClient := grabfood.NewAPIClient(configuration)
	resp, r, err := apiClient.CheckOrderCancelableAPI.CheckOrderCancelable(context.Background()).Authorization(authorization).OrderID(orderID).MerchantID(merchantID).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CheckOrderCancelableAPI.CheckOrderCancelable``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CheckOrderCancelable`: CheckOrderCancelableResponse
	fmt.Fprintf(os.Stdout, "Response from `CheckOrderCancelableAPI.CheckOrderCancelable`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCheckOrderCancelableRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **authorization** | **string** | Specify the generated authorization token of the bearer type. | 
 **orderID** | **string** |  | 
 **merchantID** | **string** | The merchant&#39;s ID that is in GrabFood&#39;s database. | 

### Return type

[**CheckOrderCancelableResponse**](CheckOrderCancelableResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

