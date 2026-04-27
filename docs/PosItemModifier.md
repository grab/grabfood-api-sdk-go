# PosItemModifier

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | The modifier’s ID in the partner system.  | [optional] 
**Name** | Pointer to **string** | The name of the modifier. | [optional] 
**Quantity** | Pointer to **int32** | The number of the modifier ordered. | [optional] 
**Price** | Pointer to **int64** | The modifier’s price in minor units and tax-inclusive in the partner system.  | [optional] 
**Tax** | Pointer to **int64** | This is the tax amount on the modifier. If grab needs to show modifier excl-tax price, we’ll use this value to calculate. &#x60;&#x60;&#x60; Price excl-tax &#x3D; price - tax  | [optional] 

## Methods

### NewPosItemModifier

`func NewPosItemModifier() *PosItemModifier`

NewPosItemModifier instantiates a new PosItemModifier object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPosItemModifierWithDefaults

`func NewPosItemModifierWithDefaults() *PosItemModifier`

NewPosItemModifierWithDefaults instantiates a new PosItemModifier object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *PosItemModifier) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *PosItemModifier) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *PosItemModifier) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *PosItemModifier) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *PosItemModifier) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *PosItemModifier) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *PosItemModifier) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *PosItemModifier) HasName() bool`

HasName returns a boolean if a field has been set.

### GetQuantity

`func (o *PosItemModifier) GetQuantity() int32`

GetQuantity returns the Quantity field if non-nil, zero value otherwise.

### GetQuantityOk

`func (o *PosItemModifier) GetQuantityOk() (*int32, bool)`

GetQuantityOk returns a tuple with the Quantity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuantity

`func (o *PosItemModifier) SetQuantity(v int32)`

SetQuantity sets Quantity field to given value.

### HasQuantity

`func (o *PosItemModifier) HasQuantity() bool`

HasQuantity returns a boolean if a field has been set.

### GetPrice

`func (o *PosItemModifier) GetPrice() int64`

GetPrice returns the Price field if non-nil, zero value otherwise.

### GetPriceOk

`func (o *PosItemModifier) GetPriceOk() (*int64, bool)`

GetPriceOk returns a tuple with the Price field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrice

`func (o *PosItemModifier) SetPrice(v int64)`

SetPrice sets Price field to given value.

### HasPrice

`func (o *PosItemModifier) HasPrice() bool`

HasPrice returns a boolean if a field has been set.

### GetTax

`func (o *PosItemModifier) GetTax() int64`

GetTax returns the Tax field if non-nil, zero value otherwise.

### GetTaxOk

`func (o *PosItemModifier) GetTaxOk() (*int64, bool)`

GetTaxOk returns a tuple with the Tax field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTax

`func (o *PosItemModifier) SetTax(v int64)`

SetTax sets Tax field to given value.

### HasTax

`func (o *PosItemModifier) HasTax() bool`

HasTax returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


