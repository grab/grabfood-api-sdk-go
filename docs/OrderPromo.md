# OrderPromo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Code** | Pointer to **string** | Promo code applied in the order. | [optional] 
**Description** | Pointer to **string** | Promo description. | [optional] 
**Name** | Pointer to **string** | Name of the promotion. | [optional] 
**PromoAmount** | Pointer to **int64** | Promo amount applied in the order, in local currency. This amount is rounded into whole number. | [optional] 
**MexFundedRatio** | Pointer to **int32** | The merchant&#39;s funded ratio of the promo in percentage. | [optional] 
**MexFundedAmount** | Pointer to **int64** | The merchant&#39;s promo fund in the minor unit. Calculated based on merchant funded ratio. | [optional] 
**TargetedPrice** | Pointer to **int64** | The subtotal of the order basket in minor unit. | [optional] 
**PromoAmountInMin** | Pointer to **int64** | Promo amount applied in the order in minor unit. | [optional] 

## Methods

### NewOrderPromo

`func NewOrderPromo() *OrderPromo`

NewOrderPromo instantiates a new OrderPromo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOrderPromoWithDefaults

`func NewOrderPromoWithDefaults() *OrderPromo`

NewOrderPromoWithDefaults instantiates a new OrderPromo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCode

`func (o *OrderPromo) GetCode() string`

GetCode returns the Code field if non-nil, zero value otherwise.

### GetCodeOk

`func (o *OrderPromo) GetCodeOk() (*string, bool)`

GetCodeOk returns a tuple with the Code field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCode

`func (o *OrderPromo) SetCode(v string)`

SetCode sets Code field to given value.

### HasCode

`func (o *OrderPromo) HasCode() bool`

HasCode returns a boolean if a field has been set.

### GetDescription

`func (o *OrderPromo) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *OrderPromo) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *OrderPromo) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *OrderPromo) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetName

`func (o *OrderPromo) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *OrderPromo) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *OrderPromo) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *OrderPromo) HasName() bool`

HasName returns a boolean if a field has been set.

### GetPromoAmount

`func (o *OrderPromo) GetPromoAmount() int64`

GetPromoAmount returns the PromoAmount field if non-nil, zero value otherwise.

### GetPromoAmountOk

`func (o *OrderPromo) GetPromoAmountOk() (*int64, bool)`

GetPromoAmountOk returns a tuple with the PromoAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPromoAmount

`func (o *OrderPromo) SetPromoAmount(v int64)`

SetPromoAmount sets PromoAmount field to given value.

### HasPromoAmount

`func (o *OrderPromo) HasPromoAmount() bool`

HasPromoAmount returns a boolean if a field has been set.

### GetMexFundedRatio

`func (o *OrderPromo) GetMexFundedRatio() int32`

GetMexFundedRatio returns the MexFundedRatio field if non-nil, zero value otherwise.

### GetMexFundedRatioOk

`func (o *OrderPromo) GetMexFundedRatioOk() (*int32, bool)`

GetMexFundedRatioOk returns a tuple with the MexFundedRatio field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMexFundedRatio

`func (o *OrderPromo) SetMexFundedRatio(v int32)`

SetMexFundedRatio sets MexFundedRatio field to given value.

### HasMexFundedRatio

`func (o *OrderPromo) HasMexFundedRatio() bool`

HasMexFundedRatio returns a boolean if a field has been set.

### GetMexFundedAmount

`func (o *OrderPromo) GetMexFundedAmount() int64`

GetMexFundedAmount returns the MexFundedAmount field if non-nil, zero value otherwise.

### GetMexFundedAmountOk

`func (o *OrderPromo) GetMexFundedAmountOk() (*int64, bool)`

GetMexFundedAmountOk returns a tuple with the MexFundedAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMexFundedAmount

`func (o *OrderPromo) SetMexFundedAmount(v int64)`

SetMexFundedAmount sets MexFundedAmount field to given value.

### HasMexFundedAmount

`func (o *OrderPromo) HasMexFundedAmount() bool`

HasMexFundedAmount returns a boolean if a field has been set.

### GetTargetedPrice

`func (o *OrderPromo) GetTargetedPrice() int64`

GetTargetedPrice returns the TargetedPrice field if non-nil, zero value otherwise.

### GetTargetedPriceOk

`func (o *OrderPromo) GetTargetedPriceOk() (*int64, bool)`

GetTargetedPriceOk returns a tuple with the TargetedPrice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetedPrice

`func (o *OrderPromo) SetTargetedPrice(v int64)`

SetTargetedPrice sets TargetedPrice field to given value.

### HasTargetedPrice

`func (o *OrderPromo) HasTargetedPrice() bool`

HasTargetedPrice returns a boolean if a field has been set.

### GetPromoAmountInMin

`func (o *OrderPromo) GetPromoAmountInMin() int64`

GetPromoAmountInMin returns the PromoAmountInMin field if non-nil, zero value otherwise.

### GetPromoAmountInMinOk

`func (o *OrderPromo) GetPromoAmountInMinOk() (*int64, bool)`

GetPromoAmountInMinOk returns a tuple with the PromoAmountInMin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPromoAmountInMin

`func (o *OrderPromo) SetPromoAmountInMin(v int64)`

SetPromoAmountInMin sets PromoAmountInMin field to given value.

### HasPromoAmountInMin

`func (o *OrderPromo) HasPromoAmountInMin() bool`

HasPromoAmountInMin returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


