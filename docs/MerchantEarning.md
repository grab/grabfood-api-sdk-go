# MerchantEarning

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Revenue** | Pointer to **int64** | The revenue of merchant should receive. revenue &#x3D; price.eaterpayment - mexfunddiscount | [optional] 
**NetEarning** | Pointer to **int64** | The netEarning of merchant should receive. netEarning &#x3D; revenue - commission | [optional] 
**MexFundDiscount** | Pointer to **int64** | The mexFundDiscount that user applied from grab app in this payment | [optional] 
**Commission** | Pointer to **int64** | The commission that grab need charge from this pay merchant transaction. | [optional] 

## Methods

### NewMerchantEarning

`func NewMerchantEarning() *MerchantEarning`

NewMerchantEarning instantiates a new MerchantEarning object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMerchantEarningWithDefaults

`func NewMerchantEarningWithDefaults() *MerchantEarning`

NewMerchantEarningWithDefaults instantiates a new MerchantEarning object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRevenue

`func (o *MerchantEarning) GetRevenue() int64`

GetRevenue returns the Revenue field if non-nil, zero value otherwise.

### GetRevenueOk

`func (o *MerchantEarning) GetRevenueOk() (*int64, bool)`

GetRevenueOk returns a tuple with the Revenue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRevenue

`func (o *MerchantEarning) SetRevenue(v int64)`

SetRevenue sets Revenue field to given value.

### HasRevenue

`func (o *MerchantEarning) HasRevenue() bool`

HasRevenue returns a boolean if a field has been set.

### GetNetEarning

`func (o *MerchantEarning) GetNetEarning() int64`

GetNetEarning returns the NetEarning field if non-nil, zero value otherwise.

### GetNetEarningOk

`func (o *MerchantEarning) GetNetEarningOk() (*int64, bool)`

GetNetEarningOk returns a tuple with the NetEarning field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetEarning

`func (o *MerchantEarning) SetNetEarning(v int64)`

SetNetEarning sets NetEarning field to given value.

### HasNetEarning

`func (o *MerchantEarning) HasNetEarning() bool`

HasNetEarning returns a boolean if a field has been set.

### GetMexFundDiscount

`func (o *MerchantEarning) GetMexFundDiscount() int64`

GetMexFundDiscount returns the MexFundDiscount field if non-nil, zero value otherwise.

### GetMexFundDiscountOk

`func (o *MerchantEarning) GetMexFundDiscountOk() (*int64, bool)`

GetMexFundDiscountOk returns a tuple with the MexFundDiscount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMexFundDiscount

`func (o *MerchantEarning) SetMexFundDiscount(v int64)`

SetMexFundDiscount sets MexFundDiscount field to given value.

### HasMexFundDiscount

`func (o *MerchantEarning) HasMexFundDiscount() bool`

HasMexFundDiscount returns a boolean if a field has been set.

### GetCommission

`func (o *MerchantEarning) GetCommission() int64`

GetCommission returns the Commission field if non-nil, zero value otherwise.

### GetCommissionOk

`func (o *MerchantEarning) GetCommissionOk() (*int64, bool)`

GetCommissionOk returns a tuple with the Commission field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCommission

`func (o *MerchantEarning) SetCommission(v int64)`

SetCommission sets Commission field to given value.

### HasCommission

`func (o *MerchantEarning) HasCommission() bool`

HasCommission returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


