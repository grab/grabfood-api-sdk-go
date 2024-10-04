# \NotifyMembershipWebviewAPI

All URIs are relative to *https://partner-api.stg-myteksi.com/grabfood-sandbox*

Method | HTTP request | Description
------------- | ------------- | -------------
[**NotifyMembershipWebview**](NotifyMembershipWebviewAPI.md#NotifyMembershipWebview) | **Post** /partner/v1/membership/notify | Notify Membership



## NotifyMembershipWebview

> NotifyMembershipWebview(ctx).Authorization(authorization).ContentType(contentType).NotifyMembershipWebviewRequest(notifyMembershipWebviewRequest).Execute()

Notify Membership

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
	notifyMembershipWebviewRequest := *grabfood.NewNotifyMembershipWebviewRequest() // NotifyMembershipWebviewRequest | 

	configuration := grabfood.NewConfiguration()
	apiClient := grabfood.NewAPIClient(configuration)
	r, err := apiClient.NotifyMembershipWebviewAPI.NotifyMembershipWebview(context.Background()).Authorization(authorization).ContentType(contentType).NotifyMembershipWebviewRequest(notifyMembershipWebviewRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `NotifyMembershipWebviewAPI.NotifyMembershipWebview``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiNotifyMembershipWebviewRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **authorization** | **string** | Specify the generated authorization token of the bearer type. | 
 **contentType** | **string** | The content type of the request body. You must use &#x60;application/json&#x60; for this header as GrabFood API currently does not support other formats. | 
 **notifyMembershipWebviewRequest** | [**NotifyMembershipWebviewRequest**](NotifyMembershipWebviewRequest.md) |  | 

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

