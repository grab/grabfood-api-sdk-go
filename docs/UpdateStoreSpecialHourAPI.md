# \UpdateStoreSpecialHourAPI

All URIs are relative to *https://partner-api.stg-myteksi.com/grabfood-sandbox*

Method | HTTP request | Description
------------- | ------------- | -------------
[**UpdateStoreSpecialHour**](UpdateStoreSpecialHourAPI.md#UpdateStoreSpecialHour) | **Put** /partner/v2/merchants/{merchantID}/store/special-opening-hour | Update Store Special Hours



## UpdateStoreSpecialHour

> UpdateSpecialHourResponse UpdateStoreSpecialHour(ctx, merchantID).ContentType(contentType).Authorization(authorization).UpdateSpecialHourRequest(updateSpecialHourRequest).Execute()

Update Store Special Hours

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
	merchantID := "1-CYNGRUNGSBCCC" // string | The merchant's ID that is in GrabFood's database.
	updateSpecialHourRequest := *grabfood.NewUpdateSpecialHourRequest([]grabfood.SpecialOpeningHour{*grabfood.NewSpecialOpeningHour()}) // UpdateSpecialHourRequest | 

	configuration := grabfood.NewConfiguration()
	apiClient := grabfood.NewAPIClient(configuration)
	resp, r, err := apiClient.UpdateStoreSpecialHourAPI.UpdateStoreSpecialHour(context.Background(), merchantID).ContentType(contentType).Authorization(authorization).UpdateSpecialHourRequest(updateSpecialHourRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UpdateStoreSpecialHourAPI.UpdateStoreSpecialHour``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateStoreSpecialHour`: UpdateSpecialHourResponse
	fmt.Fprintf(os.Stdout, "Response from `UpdateStoreSpecialHourAPI.UpdateStoreSpecialHour`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**merchantID** | **string** | The merchant&#39;s ID that is in GrabFood&#39;s database. | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateStoreSpecialHourRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **contentType** | **string** | The content type of the request body. You must use &#x60;application/json&#x60; for this header as GrabFood API currently does not support other formats. | 
 **authorization** | **string** | Specify the generated authorization token of the bearer type. | 

 **updateSpecialHourRequest** | [**UpdateSpecialHourRequest**](UpdateSpecialHourRequest.md) |  | 

### Return type

[**UpdateSpecialHourResponse**](UpdateSpecialHourResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

