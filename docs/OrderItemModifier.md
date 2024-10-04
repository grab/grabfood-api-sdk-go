# OrderItemModifier

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | The modifier&#39;s ID that is on the partner&#39;s system. | [optional] 
**Price** | Pointer to **int64** | The modifier&#39;s price (tax-inclusive) in minor format. &#x60;round(165 * (1 + 0.06)) &#x3D; 175&#x60;. | [optional] 
**Tax** | Pointer to **int64** | Tax in minor format for 1 modifier. &#x60;165*0.06&#x3D;10&#x60;. Refer to FAQs for more details about [tax](#section/Order/How-is-tax-calculated). | [optional] 
**Quantity** | Pointer to **int32** | The number of modifiers present. The value is always 1. | [optional] 

## Methods

### NewOrderItemModifier

`func NewOrderItemModifier() *OrderItemModifier`

NewOrderItemModifier instantiates a new OrderItemModifier object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOrderItemModifierWithDefaults

`func NewOrderItemModifierWithDefaults() *OrderItemModifier`

NewOrderItemModifierWithDefaults instantiates a new OrderItemModifier object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *OrderItemModifier) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *OrderItemModifier) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *OrderItemModifier) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *OrderItemModifier) HasId() bool`

HasId returns a boolean if a field has been set.

### GetPrice

`func (o *OrderItemModifier) GetPrice() int64`

GetPrice returns the Price field if non-nil, zero value otherwise.

### GetPriceOk

`func (o *OrderItemModifier) GetPriceOk() (*int64, bool)`

GetPriceOk returns a tuple with the Price field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrice

`func (o *OrderItemModifier) SetPrice(v int64)`

SetPrice sets Price field to given value.

### HasPrice

`func (o *OrderItemModifier) HasPrice() bool`

HasPrice returns a boolean if a field has been set.

### GetTax

`func (o *OrderItemModifier) GetTax() int64`

GetTax returns the Tax field if non-nil, zero value otherwise.

### GetTaxOk

`func (o *OrderItemModifier) GetTaxOk() (*int64, bool)`

GetTaxOk returns a tuple with the Tax field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTax

`func (o *OrderItemModifier) SetTax(v int64)`

SetTax sets Tax field to given value.

### HasTax

`func (o *OrderItemModifier) HasTax() bool`

HasTax returns a boolean if a field has been set.

### GetQuantity

`func (o *OrderItemModifier) GetQuantity() int32`

GetQuantity returns the Quantity field if non-nil, zero value otherwise.

### GetQuantityOk

`func (o *OrderItemModifier) GetQuantityOk() (*int32, bool)`

GetQuantityOk returns a tuple with the Quantity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuantity

`func (o *OrderItemModifier) SetQuantity(v int32)`

SetQuantity sets Quantity field to given value.

### HasQuantity

`func (o *OrderItemModifier) HasQuantity() bool`

HasQuantity returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


