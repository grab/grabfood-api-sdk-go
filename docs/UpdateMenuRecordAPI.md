# \UpdateMenuRecordAPI

All URIs are relative to *https://partner-api.stg-myteksi.com/grabfood-sandbox*

Method | HTTP request | Description
------------- | ------------- | -------------
[**BatchUpdateMenu**](UpdateMenuRecordAPI.md#BatchUpdateMenu) | **Put** /partner/v1/batch/menu | Batch Update Menu
[**UpdateMenu**](UpdateMenuRecordAPI.md#UpdateMenu) | **Put** /partner/v1/menu | Update menu record



## BatchUpdateMenu

> BatchUpdateMenuResponse BatchUpdateMenu(ctx).ContentType(contentType).Authorization(authorization).BatchUpdateMenuItem(batchUpdateMenuItem).Execute()

Batch Update Menu

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
	batchUpdateMenuItem := *grabfood.NewBatchUpdateMenuItem("1-CYNGRUNGSBCCC", "ITEM") // BatchUpdateMenuItem | 

	configuration := grabfood.NewConfiguration()
	apiClient := grabfood.NewAPIClient(configuration)
	resp, r, err := apiClient.UpdateMenuRecordAPI.BatchUpdateMenu(context.Background()).ContentType(contentType).Authorization(authorization).BatchUpdateMenuItem(batchUpdateMenuItem).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UpdateMenuRecordAPI.BatchUpdateMenu``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `BatchUpdateMenu`: BatchUpdateMenuResponse
	fmt.Fprintf(os.Stdout, "Response from `UpdateMenuRecordAPI.BatchUpdateMenu`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiBatchUpdateMenuRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **contentType** | **string** | The content type of the request body. You must use &#x60;application/json&#x60; for this header as GrabFood API currently does not support other formats. | 
 **authorization** | **string** | Specify the generated authorization token of the bearer type. | 
 **batchUpdateMenuItem** | [**BatchUpdateMenuItem**](BatchUpdateMenuItem.md) |  | 

### Return type

[**BatchUpdateMenuResponse**](BatchUpdateMenuResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateMenu

> UpdateMenu(ctx).ContentType(contentType).Authorization(authorization).UpdateMenuRequest(updateMenuRequest).Execute()

Update menu record

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
	updateMenuRequest := grabfood.update_menu_request{UpdateMenuItem: grabfood.NewUpdateMenuItem("1-CYNGRUNGSBCCC", "ITEM", "ITEM-1")} // UpdateMenuRequest | 

	configuration := grabfood.NewConfiguration()
	apiClient := grabfood.NewAPIClient(configuration)
	r, err := apiClient.UpdateMenuRecordAPI.UpdateMenu(context.Background()).ContentType(contentType).Authorization(authorization).UpdateMenuRequest(updateMenuRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UpdateMenuRecordAPI.UpdateMenu``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiUpdateMenuRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **contentType** | **string** | The content type of the request body. You must use &#x60;application/json&#x60; for this header as GrabFood API currently does not support other formats. | 
 **authorization** | **string** | Specify the generated authorization token of the bearer type. | 
 **updateMenuRequest** | [**UpdateMenuRequest**](UpdateMenuRequest.md) |  | 

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

