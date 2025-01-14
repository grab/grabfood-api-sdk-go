# \CreateSelfServeJourneyAPI

All URIs are relative to *https://partner-api.grab.com/grabfood-sandbox*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateSelfServeJourney**](CreateSelfServeJourneyAPI.md#CreateSelfServeJourney) | **Post** /partner/v1/self-serve/activation | Create self serve journey



## CreateSelfServeJourney

> CreateSelfServeJourneyResponse CreateSelfServeJourney(ctx).ContentType(contentType).Authorization(authorization).CreateSelfServeJourneyRequest(createSelfServeJourneyRequest).Execute()

Create self serve journey

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
	createSelfServeJourneyRequest := *grabfood.NewCreateSelfServeJourneyRequest(*grabfood.NewCreateSelfServeJourneyRequestPartner("Partner-ABECU")) // CreateSelfServeJourneyRequest | 

	configuration := grabfood.NewConfiguration()
	apiClient := grabfood.NewAPIClient(configuration)
	resp, r, err := apiClient.CreateSelfServeJourneyAPI.CreateSelfServeJourney(context.Background()).ContentType(contentType).Authorization(authorization).CreateSelfServeJourneyRequest(createSelfServeJourneyRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CreateSelfServeJourneyAPI.CreateSelfServeJourney``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateSelfServeJourney`: CreateSelfServeJourneyResponse
	fmt.Fprintf(os.Stdout, "Response from `CreateSelfServeJourneyAPI.CreateSelfServeJourney`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateSelfServeJourneyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **contentType** | **string** | The content type of the request body. You must use &#x60;application/json&#x60; for this header as GrabFood API currently does not support other formats. | 
 **authorization** | **string** | Specify the generated authorization token of the bearer type. | 
 **createSelfServeJourneyRequest** | [**CreateSelfServeJourneyRequest**](CreateSelfServeJourneyRequest.md) |  | 

### Return type

[**CreateSelfServeJourneyResponse**](CreateSelfServeJourneyResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

