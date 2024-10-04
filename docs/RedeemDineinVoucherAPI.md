# \RedeemDineinVoucherAPI

All URIs are relative to *https://partner-api.stg-myteksi.com/grabfood-sandbox*

Method | HTTP request | Description
------------- | ------------- | -------------
[**RedeemDineinVoucher**](RedeemDineinVoucherAPI.md#RedeemDineinVoucher) | **Post** /partner/v1/dinein/voucher/redeem | Redeem Dine In Voucher



## RedeemDineinVoucher

> RedeemDineInVoucherResponse RedeemDineinVoucher(ctx).Authorization(authorization).ContentType(contentType).RedeemDineInVoucherRequest(redeemDineInVoucherRequest).Execute()

Redeem Dine In Voucher

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
	redeemDineInVoucherRequest := *grabfood.NewRedeemDineInVoucherRequest("PPD#h9YygDAZhWkZVD50bT_xnM0bEFqTayuOdmHEhqJ4YAeDFPD3OKsEVg==", "1-CYNGRUNGSBCCC") // RedeemDineInVoucherRequest | 

	configuration := grabfood.NewConfiguration()
	apiClient := grabfood.NewAPIClient(configuration)
	resp, r, err := apiClient.RedeemDineinVoucherAPI.RedeemDineinVoucher(context.Background()).Authorization(authorization).ContentType(contentType).RedeemDineInVoucherRequest(redeemDineInVoucherRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RedeemDineinVoucherAPI.RedeemDineinVoucher``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `RedeemDineinVoucher`: RedeemDineInVoucherResponse
	fmt.Fprintf(os.Stdout, "Response from `RedeemDineinVoucherAPI.RedeemDineinVoucher`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiRedeemDineinVoucherRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **authorization** | **string** | Specify the generated authorization token of the bearer type. | 
 **contentType** | **string** | The content type of the request body. You must use &#x60;application/json&#x60; for this header as GrabFood API currently does not support other formats. | 
 **redeemDineInVoucherRequest** | [**RedeemDineInVoucherRequest**](RedeemDineInVoucherRequest.md) |  | 

### Return type

[**RedeemDineInVoucherResponse**](RedeemDineInVoucherResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

