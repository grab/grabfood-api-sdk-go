# PosPriceDetails

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Subtotal** | Pointer to **int64** | Total item and modifier price  in minor units and tax-inclusive. | [optional] 
**Tax** | Pointer to **int64** | Total tax in the minor unit. &#x60;&#x60;&#x60; Formula Tax &#x3D; total item/modifier tax  | [optional] 
**MerchantChargeFeeInMin** | Pointer to **int64** | Any additional fee charged by the merchant  in minor units and tax-inclusive, which is 100% paid out to the merchant. Eg. service charge | [optional] 
**DepositAmountInMin** | Pointer to **int64** | This field represents the reservation depositAmount, which is the amount paid upfront by the diner during the reservation process. It can be applied towards the final payment when the order is completed. It’s in minor units and tax-inclusive. | [optional] 
**OfflinePOSDiscountInMin** | Pointer to **int64** | Offline discount that is provided to the diner in minor units and tax-inclusive. | [optional] 
**BillRoundingInMin** | Pointer to **int64** | The rounding amount in minor units.  &#x60;&#x60;&#x60; Round down should be in negative value Round up should be in positive value &#x60;&#x60;&#x60;  | [optional] 
**EaterPayment** | Pointer to **int64** | The total bill value in minor units and tax-inclusive. &#x60;&#x60;&#x60; Formula:   eaterPayment &#x3D;   + subtotal   + merchantChargeFeeInMin   - depositAmountInMin   - offlinePOSDiscountInMin  | [optional] 

## Methods

### NewPosPriceDetails

`func NewPosPriceDetails() *PosPriceDetails`

NewPosPriceDetails instantiates a new PosPriceDetails object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPosPriceDetailsWithDefaults

`func NewPosPriceDetailsWithDefaults() *PosPriceDetails`

NewPosPriceDetailsWithDefaults instantiates a new PosPriceDetails object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSubtotal

`func (o *PosPriceDetails) GetSubtotal() int64`

GetSubtotal returns the Subtotal field if non-nil, zero value otherwise.

### GetSubtotalOk

`func (o *PosPriceDetails) GetSubtotalOk() (*int64, bool)`

GetSubtotalOk returns a tuple with the Subtotal field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubtotal

`func (o *PosPriceDetails) SetSubtotal(v int64)`

SetSubtotal sets Subtotal field to given value.

### HasSubtotal

`func (o *PosPriceDetails) HasSubtotal() bool`

HasSubtotal returns a boolean if a field has been set.

### GetTax

`func (o *PosPriceDetails) GetTax() int64`

GetTax returns the Tax field if non-nil, zero value otherwise.

### GetTaxOk

`func (o *PosPriceDetails) GetTaxOk() (*int64, bool)`

GetTaxOk returns a tuple with the Tax field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTax

`func (o *PosPriceDetails) SetTax(v int64)`

SetTax sets Tax field to given value.

### HasTax

`func (o *PosPriceDetails) HasTax() bool`

HasTax returns a boolean if a field has been set.

### GetMerchantChargeFeeInMin

`func (o *PosPriceDetails) GetMerchantChargeFeeInMin() int64`

GetMerchantChargeFeeInMin returns the MerchantChargeFeeInMin field if non-nil, zero value otherwise.

### GetMerchantChargeFeeInMinOk

`func (o *PosPriceDetails) GetMerchantChargeFeeInMinOk() (*int64, bool)`

GetMerchantChargeFeeInMinOk returns a tuple with the MerchantChargeFeeInMin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMerchantChargeFeeInMin

`func (o *PosPriceDetails) SetMerchantChargeFeeInMin(v int64)`

SetMerchantChargeFeeInMin sets MerchantChargeFeeInMin field to given value.

### HasMerchantChargeFeeInMin

`func (o *PosPriceDetails) HasMerchantChargeFeeInMin() bool`

HasMerchantChargeFeeInMin returns a boolean if a field has been set.

### GetDepositAmountInMin

`func (o *PosPriceDetails) GetDepositAmountInMin() int64`

GetDepositAmountInMin returns the DepositAmountInMin field if non-nil, zero value otherwise.

### GetDepositAmountInMinOk

`func (o *PosPriceDetails) GetDepositAmountInMinOk() (*int64, bool)`

GetDepositAmountInMinOk returns a tuple with the DepositAmountInMin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDepositAmountInMin

`func (o *PosPriceDetails) SetDepositAmountInMin(v int64)`

SetDepositAmountInMin sets DepositAmountInMin field to given value.

### HasDepositAmountInMin

`func (o *PosPriceDetails) HasDepositAmountInMin() bool`

HasDepositAmountInMin returns a boolean if a field has been set.

### GetOfflinePOSDiscountInMin

`func (o *PosPriceDetails) GetOfflinePOSDiscountInMin() int64`

GetOfflinePOSDiscountInMin returns the OfflinePOSDiscountInMin field if non-nil, zero value otherwise.

### GetOfflinePOSDiscountInMinOk

`func (o *PosPriceDetails) GetOfflinePOSDiscountInMinOk() (*int64, bool)`

GetOfflinePOSDiscountInMinOk returns a tuple with the OfflinePOSDiscountInMin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOfflinePOSDiscountInMin

`func (o *PosPriceDetails) SetOfflinePOSDiscountInMin(v int64)`

SetOfflinePOSDiscountInMin sets OfflinePOSDiscountInMin field to given value.

### HasOfflinePOSDiscountInMin

`func (o *PosPriceDetails) HasOfflinePOSDiscountInMin() bool`

HasOfflinePOSDiscountInMin returns a boolean if a field has been set.

### GetBillRoundingInMin

`func (o *PosPriceDetails) GetBillRoundingInMin() int64`

GetBillRoundingInMin returns the BillRoundingInMin field if non-nil, zero value otherwise.

### GetBillRoundingInMinOk

`func (o *PosPriceDetails) GetBillRoundingInMinOk() (*int64, bool)`

GetBillRoundingInMinOk returns a tuple with the BillRoundingInMin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBillRoundingInMin

`func (o *PosPriceDetails) SetBillRoundingInMin(v int64)`

SetBillRoundingInMin sets BillRoundingInMin field to given value.

### HasBillRoundingInMin

`func (o *PosPriceDetails) HasBillRoundingInMin() bool`

HasBillRoundingInMin returns a boolean if a field has been set.

### GetEaterPayment

`func (o *PosPriceDetails) GetEaterPayment() int64`

GetEaterPayment returns the EaterPayment field if non-nil, zero value otherwise.

### GetEaterPaymentOk

`func (o *PosPriceDetails) GetEaterPaymentOk() (*int64, bool)`

GetEaterPaymentOk returns a tuple with the EaterPayment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEaterPayment

`func (o *PosPriceDetails) SetEaterPayment(v int64)`

SetEaterPayment sets EaterPayment field to given value.

### HasEaterPayment

`func (o *PosPriceDetails) HasEaterPayment() bool`

HasEaterPayment returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


