# RedeemDineInVoucherRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CertificateID** | **string** | This certificateID is decoded from scanning the QR code, and 1:1 mapping with &#x60;voucherCode&#x60;. | 
**MerchantID** | **string** | The merchant&#39;s ID that is in GrabFood&#39;s database. | 

## Methods

### NewRedeemDineInVoucherRequest

`func NewRedeemDineInVoucherRequest(certificateID string, merchantID string, ) *RedeemDineInVoucherRequest`

NewRedeemDineInVoucherRequest instantiates a new RedeemDineInVoucherRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRedeemDineInVoucherRequestWithDefaults

`func NewRedeemDineInVoucherRequestWithDefaults() *RedeemDineInVoucherRequest`

NewRedeemDineInVoucherRequestWithDefaults instantiates a new RedeemDineInVoucherRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCertificateID

`func (o *RedeemDineInVoucherRequest) GetCertificateID() string`

GetCertificateID returns the CertificateID field if non-nil, zero value otherwise.

### GetCertificateIDOk

`func (o *RedeemDineInVoucherRequest) GetCertificateIDOk() (*string, bool)`

GetCertificateIDOk returns a tuple with the CertificateID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCertificateID

`func (o *RedeemDineInVoucherRequest) SetCertificateID(v string)`

SetCertificateID sets CertificateID field to given value.


### GetMerchantID

`func (o *RedeemDineInVoucherRequest) GetMerchantID() string`

GetMerchantID returns the MerchantID field if non-nil, zero value otherwise.

### GetMerchantIDOk

`func (o *RedeemDineInVoucherRequest) GetMerchantIDOk() (*string, bool)`

GetMerchantIDOk returns a tuple with the MerchantID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMerchantID

`func (o *RedeemDineInVoucherRequest) SetMerchantID(v string)`

SetMerchantID sets MerchantID field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


