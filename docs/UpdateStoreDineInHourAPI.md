# \UpdateStoreDineInHourAPI

All URIs are relative to *https://partner-api.stg-myteksi.com/grabfood-sandbox*

Method | HTTP request | Description
------------- | ------------- | -------------
[**UpdateStoreDineInHour**](UpdateStoreDineInHourAPI.md#UpdateStoreDineInHour) | **Put** /partner/v1/merchants/{merchantID}/store/dine-in-hours | Update Store Dine-in Hours



## UpdateStoreDineInHour

> UpdateDineInHourResponse UpdateStoreDineInHour(ctx, merchantID).ContentType(contentType).Authorization(authorization).UpdateDineInHourRequest(updateDineInHourRequest).Execute()

Update Store Dine-in Hours

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
	updateDineInHourRequest := *grabfood.NewUpdateDineInHourRequest(*grabfood.NewStoreHour([]grabfood.OpenPeriod{*grabfood.NewOpenPeriod("11:30", "21:30")}, []grabfood.OpenPeriod{*grabfood.NewOpenPeriod("11:30", "21:30")}, []grabfood.OpenPeriod{*grabfood.NewOpenPeriod("11:30", "21:30")}, []grabfood.OpenPeriod{*grabfood.NewOpenPeriod("11:30", "21:30")}, []grabfood.OpenPeriod{*grabfood.NewOpenPeriod("11:30", "21:30")}, []grabfood.OpenPeriod{*grabfood.NewOpenPeriod("11:30", "21:30")}, []grabfood.OpenPeriod{*grabfood.NewOpenPeriod("11:30", "21:30")})) // UpdateDineInHourRequest | 

	configuration := grabfood.NewConfiguration()
	apiClient := grabfood.NewAPIClient(configuration)
	resp, r, err := apiClient.UpdateStoreDineInHourAPI.UpdateStoreDineInHour(context.Background(), merchantID).ContentType(contentType).Authorization(authorization).UpdateDineInHourRequest(updateDineInHourRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UpdateStoreDineInHourAPI.UpdateStoreDineInHour``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateStoreDineInHour`: UpdateDineInHourResponse
	fmt.Fprintf(os.Stdout, "Response from `UpdateStoreDineInHourAPI.UpdateStoreDineInHour`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**merchantID** | **string** | The merchant&#39;s ID that is in GrabFood&#39;s database. | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateStoreDineInHourRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **contentType** | **string** | The content type of the request body. You must use &#x60;application/json&#x60; for this header as GrabFood API currently does not support other formats. | 
 **authorization** | **string** | Specify the generated authorization token of the bearer type. | 

 **updateDineInHourRequest** | [**UpdateDineInHourRequest**](UpdateDineInHourRequest.md) |  | 

### Return type

[**UpdateDineInHourResponse**](UpdateDineInHourResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

