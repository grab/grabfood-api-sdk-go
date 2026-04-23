# \GenerateStoQrCodeAPI

All URIs are relative to *https://partner-api.grab.com/grabfood-sandbox*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GenerateStoQrCode**](GenerateStoQrCodeAPI.md#GenerateStoQrCode) | **Get** /partner/v1/dinein/sto/qrcode | Generate STO QR code



## GenerateStoQrCode

> GenerateSTOQRCodeResponse GenerateStoQrCode(ctx).Authorization(authorization).ContentType(contentType).MerchantID(merchantID).QrType(qrType).TableNumber(tableNumber).Execute()

Generate STO QR code

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
	merchantID := "merchantID_example" // string | 
	qrType := "qrType_example" // string | 
	tableNumber := "tableNumber_example" // string | 

	configuration := grabfood.NewConfiguration()
	apiClient := grabfood.NewAPIClient(configuration)
	resp, r, err := apiClient.GenerateStoQrCodeAPI.GenerateStoQrCode(context.Background()).Authorization(authorization).ContentType(contentType).MerchantID(merchantID).QrType(qrType).TableNumber(tableNumber).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `GenerateStoQrCodeAPI.GenerateStoQrCode``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GenerateStoQrCode`: GenerateSTOQRCodeResponse
	fmt.Fprintf(os.Stdout, "Response from `GenerateStoQrCodeAPI.GenerateStoQrCode`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGenerateStoQrCodeRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **authorization** | **string** | Specify the generated authorization token of the bearer type. | 
 **contentType** | **string** | The content type of the request body. You must use &#x60;application/json&#x60; for this header as GrabFood API currently does not support other formats. | 
 **merchantID** | **string** |  | 
 **qrType** | **string** |  | 
 **tableNumber** | **string** |  | 

### Return type

[**GenerateSTOQRCodeResponse**](GenerateSTOQRCodeResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

