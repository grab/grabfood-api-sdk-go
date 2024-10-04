# RedeemDineInVoucherResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CertificateID** | Pointer to **string** | This certificateID is decoded from scanning the QR code, and 1:1 mapping with &#x60;voucherCode&#x60;. | [optional] 
**VoucherCode** | Pointer to **string** | A short code for the dine-in voucher purchased by the user. | [optional] 
**Voucher** | Pointer to [**Voucher**](Voucher.md) |  | [optional] 
**RedeemResult** | Pointer to [**RedeemResult**](RedeemResult.md) |  | [optional] 
**MerchantID** | Pointer to **string** | The merchant&#39;s ID that is in GrabFood&#39;s database. | [optional] 
**CampaignID** | Pointer to **string** | The dine-in voucher campaign&#39;s ID in GrabFood&#39;s database. | [optional] 

## Methods

### NewRedeemDineInVoucherResponse

`func NewRedeemDineInVoucherResponse() *RedeemDineInVoucherResponse`

NewRedeemDineInVoucherResponse instantiates a new RedeemDineInVoucherResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRedeemDineInVoucherResponseWithDefaults

`func NewRedeemDineInVoucherResponseWithDefaults() *RedeemDineInVoucherResponse`

NewRedeemDineInVoucherResponseWithDefaults instantiates a new RedeemDineInVoucherResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCertificateID

`func (o *RedeemDineInVoucherResponse) GetCertificateID() string`

GetCertificateID returns the CertificateID field if non-nil, zero value otherwise.

### GetCertificateIDOk

`func (o *RedeemDineInVoucherResponse) GetCertificateIDOk() (*string, bool)`

GetCertificateIDOk returns a tuple with the CertificateID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCertificateID

`func (o *RedeemDineInVoucherResponse) SetCertificateID(v string)`

SetCertificateID sets CertificateID field to given value.

### HasCertificateID

`func (o *RedeemDineInVoucherResponse) HasCertificateID() bool`

HasCertificateID returns a boolean if a field has been set.

### GetVoucherCode

`func (o *RedeemDineInVoucherResponse) GetVoucherCode() string`

GetVoucherCode returns the VoucherCode field if non-nil, zero value otherwise.

### GetVoucherCodeOk

`func (o *RedeemDineInVoucherResponse) GetVoucherCodeOk() (*string, bool)`

GetVoucherCodeOk returns a tuple with the VoucherCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVoucherCode

`func (o *RedeemDineInVoucherResponse) SetVoucherCode(v string)`

SetVoucherCode sets VoucherCode field to given value.

### HasVoucherCode

`func (o *RedeemDineInVoucherResponse) HasVoucherCode() bool`

HasVoucherCode returns a boolean if a field has been set.

### GetVoucher

`func (o *RedeemDineInVoucherResponse) GetVoucher() Voucher`

GetVoucher returns the Voucher field if non-nil, zero value otherwise.

### GetVoucherOk

`func (o *RedeemDineInVoucherResponse) GetVoucherOk() (*Voucher, bool)`

GetVoucherOk returns a tuple with the Voucher field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVoucher

`func (o *RedeemDineInVoucherResponse) SetVoucher(v Voucher)`

SetVoucher sets Voucher field to given value.

### HasVoucher

`func (o *RedeemDineInVoucherResponse) HasVoucher() bool`

HasVoucher returns a boolean if a field has been set.

### GetRedeemResult

`func (o *RedeemDineInVoucherResponse) GetRedeemResult() RedeemResult`

GetRedeemResult returns the RedeemResult field if non-nil, zero value otherwise.

### GetRedeemResultOk

`func (o *RedeemDineInVoucherResponse) GetRedeemResultOk() (*RedeemResult, bool)`

GetRedeemResultOk returns a tuple with the RedeemResult field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRedeemResult

`func (o *RedeemDineInVoucherResponse) SetRedeemResult(v RedeemResult)`

SetRedeemResult sets RedeemResult field to given value.

### HasRedeemResult

`func (o *RedeemDineInVoucherResponse) HasRedeemResult() bool`

HasRedeemResult returns a boolean if a field has been set.

### GetMerchantID

`func (o *RedeemDineInVoucherResponse) GetMerchantID() string`

GetMerchantID returns the MerchantID field if non-nil, zero value otherwise.

### GetMerchantIDOk

`func (o *RedeemDineInVoucherResponse) GetMerchantIDOk() (*string, bool)`

GetMerchantIDOk returns a tuple with the MerchantID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMerchantID

`func (o *RedeemDineInVoucherResponse) SetMerchantID(v string)`

SetMerchantID sets MerchantID field to given value.

### HasMerchantID

`func (o *RedeemDineInVoucherResponse) HasMerchantID() bool`

HasMerchantID returns a boolean if a field has been set.

### GetCampaignID

`func (o *RedeemDineInVoucherResponse) GetCampaignID() string`

GetCampaignID returns the CampaignID field if non-nil, zero value otherwise.

### GetCampaignIDOk

`func (o *RedeemDineInVoucherResponse) GetCampaignIDOk() (*string, bool)`

GetCampaignIDOk returns a tuple with the CampaignID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCampaignID

`func (o *RedeemDineInVoucherResponse) SetCampaignID(v string)`

SetCampaignID sets CampaignID field to given value.

### HasCampaignID

`func (o *RedeemDineInVoucherResponse) HasCampaignID() bool`

HasCampaignID returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


