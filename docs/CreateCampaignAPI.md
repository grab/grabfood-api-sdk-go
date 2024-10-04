# \CreateCampaignAPI

All URIs are relative to *https://partner-api.stg-myteksi.com/grabfood-sandbox*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateCampaign**](CreateCampaignAPI.md#CreateCampaign) | **Post** /partner/v1/campaigns | Create campaign



## CreateCampaign

> CreateCampaignResponse CreateCampaign(ctx).ContentType(contentType).Authorization(authorization).CreateCampaignRequest(createCampaignRequest).Execute()

Create campaign

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
    "time"
	grabfood "github.com/grab/grabfood-api-sdk-go"
)

func main() {
	contentType := "application/json" // string | The content type of the request body. You must use `application/json` for this header as GrabFood API currently does not support other formats.
	authorization := "Bearer <ACCESS_TOKEN_HERE>" // string | Specify the generated authorization token of the bearer type.
	createCampaignRequest := *grabfood.NewCreateCampaignRequest("1-CYNGRUNGSBCCC", "$4 off with min $10 order for all users within weekday", *grabfood.NewCampaignConditions(time.Now(), time.Now(), "EaterType_example"), *grabfood.NewCampaignDiscount("percentage", *grabfood.NewCampaignScope("items"))) // CreateCampaignRequest | 

	configuration := grabfood.NewConfiguration()
	apiClient := grabfood.NewAPIClient(configuration)
	resp, r, err := apiClient.CreateCampaignAPI.CreateCampaign(context.Background()).ContentType(contentType).Authorization(authorization).CreateCampaignRequest(createCampaignRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CreateCampaignAPI.CreateCampaign``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateCampaign`: CreateCampaignResponse
	fmt.Fprintf(os.Stdout, "Response from `CreateCampaignAPI.CreateCampaign`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateCampaignRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **contentType** | **string** | The content type of the request body. You must use &#x60;application/json&#x60; for this header as GrabFood API currently does not support other formats. | 
 **authorization** | **string** | Specify the generated authorization token of the bearer type. | 
 **createCampaignRequest** | [**CreateCampaignRequest**](CreateCampaignRequest.md) |  | 

### Return type

[**CreateCampaignResponse**](CreateCampaignResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

