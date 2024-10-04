# \DeleteCampaignAPI

All URIs are relative to *https://partner-api.stg-myteksi.com/grabfood-sandbox*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeleteCampaign**](DeleteCampaignAPI.md#DeleteCampaign) | **Delete** /partner/v1/campaigns/{campaign_id} | Delete campaigns



## DeleteCampaign

> DeleteCampaign(ctx, campaignId).Authorization(authorization).Execute()

Delete campaigns

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
	campaignId := "campaignId_example" // string | 

	configuration := grabfood.NewConfiguration()
	apiClient := grabfood.NewAPIClient(configuration)
	r, err := apiClient.DeleteCampaignAPI.DeleteCampaign(context.Background(), campaignId).Authorization(authorization).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DeleteCampaignAPI.DeleteCampaign``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**campaignId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteCampaignRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **authorization** | **string** | Specify the generated authorization token of the bearer type. | 


### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

