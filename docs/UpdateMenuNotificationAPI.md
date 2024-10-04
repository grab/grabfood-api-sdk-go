# \UpdateMenuNotificationAPI

All URIs are relative to *https://partner-api.stg-myteksi.com/grabfood-sandbox*

Method | HTTP request | Description
------------- | ------------- | -------------
[**UpdateMenuNotification**](UpdateMenuNotificationAPI.md#UpdateMenuNotification) | **Post** /partner/v1/merchant/menu/notification | Notify Grab of updated menu



## UpdateMenuNotification

> UpdateMenuNotification(ctx).ContentType(contentType).Authorization(authorization).UpdateMenuNotifRequest(updateMenuNotifRequest).Execute()

Notify Grab of updated menu

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
	updateMenuNotifRequest := *grabfood.NewUpdateMenuNotifRequest("1-CYNGRUNGSBCCC") // UpdateMenuNotifRequest | 

	configuration := grabfood.NewConfiguration()
	apiClient := grabfood.NewAPIClient(configuration)
	r, err := apiClient.UpdateMenuNotificationAPI.UpdateMenuNotification(context.Background()).ContentType(contentType).Authorization(authorization).UpdateMenuNotifRequest(updateMenuNotifRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UpdateMenuNotificationAPI.UpdateMenuNotification``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiUpdateMenuNotificationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **contentType** | **string** | The content type of the request body. You must use &#x60;application/json&#x60; for this header as GrabFood API currently does not support other formats. | 
 **authorization** | **string** | Specify the generated authorization token of the bearer type. | 
 **updateMenuNotifRequest** | [**UpdateMenuNotifRequest**](UpdateMenuNotifRequest.md) |  | 

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

