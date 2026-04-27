# EditOrderItemModifier

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | The modifier&#39;s external ID in partner system. | [optional] 
**Quantity** | Pointer to **int32** | The modifier&#39;s quantity. 0 to remove the modifier, 1 to add or keep the modifier. | [optional] 

## Methods

### NewEditOrderItemModifier

`func NewEditOrderItemModifier() *EditOrderItemModifier`

NewEditOrderItemModifier instantiates a new EditOrderItemModifier object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEditOrderItemModifierWithDefaults

`func NewEditOrderItemModifierWithDefaults() *EditOrderItemModifier`

NewEditOrderItemModifierWithDefaults instantiates a new EditOrderItemModifier object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *EditOrderItemModifier) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *EditOrderItemModifier) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *EditOrderItemModifier) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *EditOrderItemModifier) HasId() bool`

HasId returns a boolean if a field has been set.

### GetQuantity

`func (o *EditOrderItemModifier) GetQuantity() int32`

GetQuantity returns the Quantity field if non-nil, zero value otherwise.

### GetQuantityOk

`func (o *EditOrderItemModifier) GetQuantityOk() (*int32, bool)`

GetQuantityOk returns a tuple with the Quantity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuantity

`func (o *EditOrderItemModifier) SetQuantity(v int32)`

SetQuantity sets Quantity field to given value.

### HasQuantity

`func (o *EditOrderItemModifier) HasQuantity() bool`

HasQuantity returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


