# \ListOrdersAPI

All URIs are relative to *https://partner-api.stg-myteksi.com/grabfood-sandbox*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ListOrders**](ListOrdersAPI.md#ListOrders) | **Get** /partner/v1/orders | List orders



## ListOrders

> ListOrdersResponse ListOrders(ctx).Authorization(authorization).MerchantID(merchantID).Date(date).Page(page).Execute()

List orders

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
	merchantID := "1-CYNGRUNGSBCCC" // string | The merchant's ID that is in GrabFood's database.
	date := "date_example" // string | 
	page := int32(1) // int32 | Specify the page number for the report.

	configuration := grabfood.NewConfiguration()
	apiClient := grabfood.NewAPIClient(configuration)
	resp, r, err := apiClient.ListOrdersAPI.ListOrders(context.Background()).Authorization(authorization).MerchantID(merchantID).Date(date).Page(page).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ListOrdersAPI.ListOrders``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListOrders`: ListOrdersResponse
	fmt.Fprintf(os.Stdout, "Response from `ListOrdersAPI.ListOrders`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListOrdersRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **authorization** | **string** | Specify the generated authorization token of the bearer type. | 
 **merchantID** | **string** | The merchant&#39;s ID that is in GrabFood&#39;s database. | 
 **date** | **string** |  | 
 **page** | **int32** | Specify the page number for the report. | 

### Return type

[**ListOrdersResponse**](ListOrdersResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

