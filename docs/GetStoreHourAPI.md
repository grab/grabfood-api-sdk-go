# \GetStoreHourAPI

All URIs are relative to *https://partner-api.stg-myteksi.com/grabfood-sandbox*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetStoreHour**](GetStoreHourAPI.md#GetStoreHour) | **Get** /partner/v2/merchants/{merchantID}/store/hours | Get Store Hours



## GetStoreHour

> StoreHourResponse GetStoreHour(ctx, merchantID).Authorization(authorization).Execute()

Get Store Hours

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
	resp, r, err := apiClient.GetStoreHourAPI.GetStoreHour(context.Background(), merchantID).Authorization(authorization).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `GetStoreHourAPI.GetStoreHour``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetStoreHour`: StoreHourResponse
	fmt.Fprintf(os.Stdout, "Response from `GetStoreHourAPI.GetStoreHour`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**merchantID** | **string** | The merchant&#39;s ID that is in GrabFood&#39;s database. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetStoreHourRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **authorization** | **string** | Specify the generated authorization token of the bearer type. | 


### Return type

[**StoreHourResponse**](StoreHourResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

