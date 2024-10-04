# \GetStoreStatusAPI

All URIs are relative to *https://partner-api.stg-myteksi.com/grabfood-sandbox*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetStoreStatus**](GetStoreStatusAPI.md#GetStoreStatus) | **Get** /partner/v1/merchants/{merchantID}/store/status | Get Store Status



## GetStoreStatus

> StoreStatusResponse GetStoreStatus(ctx, merchantID).Authorization(authorization).Execute()

Get Store Status

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

	configuration := grabfood.NewConfiguration()
	apiClient := grabfood.NewAPIClient(configuration)
	resp, r, err := apiClient.GetStoreStatusAPI.GetStoreStatus(context.Background(), merchantID).Authorization(authorization).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `GetStoreStatusAPI.GetStoreStatus``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetStoreStatus`: StoreStatusResponse
	fmt.Fprintf(os.Stdout, "Response from `GetStoreStatusAPI.GetStoreStatus`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**merchantID** | **string** | The merchant&#39;s ID that is in GrabFood&#39;s database. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetStoreStatusRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **authorization** | **string** | Specify the generated authorization token of the bearer type. | 


### Return type

[**StoreStatusResponse**](StoreStatusResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

