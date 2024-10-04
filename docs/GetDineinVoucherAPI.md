# \GetDineinVoucherAPI

All URIs are relative to *https://partner-api.stg-myteksi.com/grabfood-sandbox*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetDineinVoucher**](GetDineinVoucherAPI.md#GetDineinVoucher) | **Get** /partner/v1/dinein/voucher | Get Dine In Voucher



## GetDineinVoucher

> GetDineInVoucherResponse GetDineinVoucher(ctx).Authorization(authorization).MerchantID(merchantID).VoucherCode(voucherCode).CertificateID(certificateID).Execute()

Get Dine In Voucher

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
	merchantID := "merchantID_example" // string | 
	voucherCode := "voucherCode_example" // string | A short code for the dine-in voucher purchased by the user. Required if `certificateID` is not specified. (optional)
	certificateID := "certificateID_example" // string | This certificateID is decoded from scanning the QR code, and 1:1 mapping with voucherCode. Required if `voucherCode` is not specified. (optional)

	configuration := grabfood.NewConfiguration()
	apiClient := grabfood.NewAPIClient(configuration)
	resp, r, err := apiClient.GetDineinVoucherAPI.GetDineinVoucher(context.Background()).Authorization(authorization).MerchantID(merchantID).VoucherCode(voucherCode).CertificateID(certificateID).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `GetDineinVoucherAPI.GetDineinVoucher``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetDineinVoucher`: GetDineInVoucherResponse
	fmt.Fprintf(os.Stdout, "Response from `GetDineinVoucherAPI.GetDineinVoucher`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetDineinVoucherRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **authorization** | **string** | Specify the generated authorization token of the bearer type. | 
 **merchantID** | **string** |  | 
 **voucherCode** | **string** | A short code for the dine-in voucher purchased by the user. Required if &#x60;certificateID&#x60; is not specified. | 
 **certificateID** | **string** | This certificateID is decoded from scanning the QR code, and 1:1 mapping with voucherCode. Required if &#x60;voucherCode&#x60; is not specified. | 

### Return type

[**GetDineInVoucherResponse**](GetDineInVoucherResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

