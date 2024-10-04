# GetDineInVoucherResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CertificateID** | Pointer to **string** | This certificateID is decoded from scanning the QR code, and 1:1 mapping with &#x60;voucherCode&#x60;. | [optional] 
**VoucherCode** | Pointer to **string** | A short code for the dine-in voucher purchased by the user. | [optional] 
**Voucher** | Pointer to [**Voucher**](Voucher.md) |  | [optional] 
**VoucherStatus** | Pointer to **string** | The status of the dine-in voucher purchased. Only active voucher is eligible for redemption. | [optional] 
**MerchantID** | Pointer to **string** | The merchant&#39;s ID that is in GrabFood&#39;s database. | [optional] 
**CampaignID** | Pointer to **string** | The dine-in voucher campaign&#39;s ID in GrabFood&#39;s database. | [optional] 

## Methods

### NewGetDineInVoucherResponse

`func NewGetDineInVoucherResponse() *GetDineInVoucherResponse`

NewGetDineInVoucherResponse instantiates a new GetDineInVoucherResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetDineInVoucherResponseWithDefaults

`func NewGetDineInVoucherResponseWithDefaults() *GetDineInVoucherResponse`

NewGetDineInVoucherResponseWithDefaults instantiates a new GetDineInVoucherResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCertificateID

`func (o *GetDineInVoucherResponse) GetCertificateID() string`

GetCertificateID returns the CertificateID field if non-nil, zero value otherwise.

### GetCertificateIDOk

`func (o *GetDineInVoucherResponse) GetCertificateIDOk() (*string, bool)`

GetCertificateIDOk returns a tuple with the CertificateID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCertificateID

`func (o *GetDineInVoucherResponse) SetCertificateID(v string)`

SetCertificateID sets CertificateID field to given value.

### HasCertificateID

`func (o *GetDineInVoucherResponse) HasCertificateID() bool`

HasCertificateID returns a boolean if a field has been set.

### GetVoucherCode

`func (o *GetDineInVoucherResponse) GetVoucherCode() string`

GetVoucherCode returns the VoucherCode field if non-nil, zero value otherwise.

### GetVoucherCodeOk

`func (o *GetDineInVoucherResponse) GetVoucherCodeOk() (*string, bool)`

GetVoucherCodeOk returns a tuple with the VoucherCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVoucherCode

`func (o *GetDineInVoucherResponse) SetVoucherCode(v string)`

SetVoucherCode sets VoucherCode field to given value.

### HasVoucherCode

`func (o *GetDineInVoucherResponse) HasVoucherCode() bool`

HasVoucherCode returns a boolean if a field has been set.

### GetVoucher

`func (o *GetDineInVoucherResponse) GetVoucher() Voucher`

GetVoucher returns the Voucher field if non-nil, zero value otherwise.

### GetVoucherOk

`func (o *GetDineInVoucherResponse) GetVoucherOk() (*Voucher, bool)`

GetVoucherOk returns a tuple with the Voucher field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVoucher

`func (o *GetDineInVoucherResponse) SetVoucher(v Voucher)`

SetVoucher sets Voucher field to given value.

### HasVoucher

`func (o *GetDineInVoucherResponse) HasVoucher() bool`

HasVoucher returns a boolean if a field has been set.

### GetVoucherStatus

`func (o *GetDineInVoucherResponse) GetVoucherStatus() string`

GetVoucherStatus returns the VoucherStatus field if non-nil, zero value otherwise.

### GetVoucherStatusOk

`func (o *GetDineInVoucherResponse) GetVoucherStatusOk() (*string, bool)`

GetVoucherStatusOk returns a tuple with the VoucherStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVoucherStatus

`func (o *GetDineInVoucherResponse) SetVoucherStatus(v string)`

SetVoucherStatus sets VoucherStatus field to given value.

### HasVoucherStatus

`func (o *GetDineInVoucherResponse) HasVoucherStatus() bool`

HasVoucherStatus returns a boolean if a field has been set.

### GetMerchantID

`func (o *GetDineInVoucherResponse) GetMerchantID() string`

GetMerchantID returns the MerchantID field if non-nil, zero value otherwise.

### GetMerchantIDOk

`func (o *GetDineInVoucherResponse) GetMerchantIDOk() (*string, bool)`

GetMerchantIDOk returns a tuple with the MerchantID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMerchantID

`func (o *GetDineInVoucherResponse) SetMerchantID(v string)`

SetMerchantID sets MerchantID field to given value.

### HasMerchantID

`func (o *GetDineInVoucherResponse) HasMerchantID() bool`

HasMerchantID returns a boolean if a field has been set.

### GetCampaignID

`func (o *GetDineInVoucherResponse) GetCampaignID() string`

GetCampaignID returns the CampaignID field if non-nil, zero value otherwise.

### GetCampaignIDOk

`func (o *GetDineInVoucherResponse) GetCampaignIDOk() (*string, bool)`

GetCampaignIDOk returns a tuple with the CampaignID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCampaignID

`func (o *GetDineInVoucherResponse) SetCampaignID(v string)`

SetCampaignID sets CampaignID field to given value.

### HasCampaignID

`func (o *GetDineInVoucherResponse) HasCampaignID() bool`

HasCampaignID returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


