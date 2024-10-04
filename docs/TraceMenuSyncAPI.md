# \TraceMenuSyncAPI

All URIs are relative to *https://partner-api.stg-myteksi.com/grabfood-sandbox*

Method | HTTP request | Description
------------- | ------------- | -------------
[**TraceMenuSync**](TraceMenuSyncAPI.md#TraceMenuSync) | **Get** /partner/v1/merchant/menu/trace | Trace menu sync



## TraceMenuSync

> MenuSyncResponse TraceMenuSync(ctx).Authorization(authorization).MerchantID(merchantID).Execute()

Trace menu sync

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
	resp, r, err := apiClient.TraceMenuSyncAPI.TraceMenuSync(context.Background()).Authorization(authorization).MerchantID(merchantID).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TraceMenuSyncAPI.TraceMenuSync``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `TraceMenuSync`: MenuSyncResponse
	fmt.Fprintf(os.Stdout, "Response from `TraceMenuSyncAPI.TraceMenuSync`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiTraceMenuSyncRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **authorization** | **string** | Specify the generated authorization token of the bearer type. | 
 **merchantID** | **string** | The merchant&#39;s ID that is in GrabFood&#39;s database. | 

### Return type

[**MenuSyncResponse**](MenuSyncResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

