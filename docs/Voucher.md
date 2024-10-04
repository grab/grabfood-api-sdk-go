# Voucher

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Title** | Pointer to **string** | The title of the voucher. | [optional] 
**DiscountedPrice** | Pointer to **string** | The amount paid after discount is applied in local currency. | [optional] 
**OriginalPrice** | Pointer to **string** | The original amount before discount is applied in local currency. | [optional] 
**DescriptionInfo** | Pointer to [**VoucherDescriptionInfo**](VoucherDescriptionInfo.md) |  | [optional] 
**VoucherType** | Pointer to **string** | The type of the dine-in voucher. | [optional] 

## Methods

### NewVoucher

`func NewVoucher() *Voucher`

NewVoucher instantiates a new Voucher object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVoucherWithDefaults

`func NewVoucherWithDefaults() *Voucher`

NewVoucherWithDefaults instantiates a new Voucher object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTitle

`func (o *Voucher) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *Voucher) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *Voucher) SetTitle(v string)`

SetTitle sets Title field to given value.

### HasTitle

`func (o *Voucher) HasTitle() bool`

HasTitle returns a boolean if a field has been set.

### GetDiscountedPrice

`func (o *Voucher) GetDiscountedPrice() string`

GetDiscountedPrice returns the DiscountedPrice field if non-nil, zero value otherwise.

### GetDiscountedPriceOk

`func (o *Voucher) GetDiscountedPriceOk() (*string, bool)`

GetDiscountedPriceOk returns a tuple with the DiscountedPrice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDiscountedPrice

`func (o *Voucher) SetDiscountedPrice(v string)`

SetDiscountedPrice sets DiscountedPrice field to given value.

### HasDiscountedPrice

`func (o *Voucher) HasDiscountedPrice() bool`

HasDiscountedPrice returns a boolean if a field has been set.

### GetOriginalPrice

`func (o *Voucher) GetOriginalPrice() string`

GetOriginalPrice returns the OriginalPrice field if non-nil, zero value otherwise.

### GetOriginalPriceOk

`func (o *Voucher) GetOriginalPriceOk() (*string, bool)`

GetOriginalPriceOk returns a tuple with the OriginalPrice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOriginalPrice

`func (o *Voucher) SetOriginalPrice(v string)`

SetOriginalPrice sets OriginalPrice field to given value.

### HasOriginalPrice

`func (o *Voucher) HasOriginalPrice() bool`

HasOriginalPrice returns a boolean if a field has been set.

### GetDescriptionInfo

`func (o *Voucher) GetDescriptionInfo() VoucherDescriptionInfo`

GetDescriptionInfo returns the DescriptionInfo field if non-nil, zero value otherwise.

### GetDescriptionInfoOk

`func (o *Voucher) GetDescriptionInfoOk() (*VoucherDescriptionInfo, bool)`

GetDescriptionInfoOk returns a tuple with the DescriptionInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescriptionInfo

`func (o *Voucher) SetDescriptionInfo(v VoucherDescriptionInfo)`

SetDescriptionInfo sets DescriptionInfo field to given value.

### HasDescriptionInfo

`func (o *Voucher) HasDescriptionInfo() bool`

HasDescriptionInfo returns a boolean if a field has been set.

### GetVoucherType

`func (o *Voucher) GetVoucherType() string`

GetVoucherType returns the VoucherType field if non-nil, zero value otherwise.

### GetVoucherTypeOk

`func (o *Voucher) GetVoucherTypeOk() (*string, bool)`

GetVoucherTypeOk returns a tuple with the VoucherType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVoucherType

`func (o *Voucher) SetVoucherType(v string)`

SetVoucherType sets VoucherType field to given value.

### HasVoucherType

`func (o *Voucher) HasVoucherType() bool`

HasVoucherType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


