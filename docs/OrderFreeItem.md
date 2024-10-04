# OrderFreeItem

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | The item&#39;s ID | [optional] 
**Name** | Pointer to **string** | The name of the item.  | [optional] 
**Quantity** | Pointer to **int32** | The item&#39;s quantity. Maximum is **1**. | [optional] 
**Price** | Pointer to **int64** | The item&#39;s price in minor unit format. | [optional] 

## Methods

### NewOrderFreeItem

`func NewOrderFreeItem() *OrderFreeItem`

NewOrderFreeItem instantiates a new OrderFreeItem object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOrderFreeItemWithDefaults

`func NewOrderFreeItemWithDefaults() *OrderFreeItem`

NewOrderFreeItemWithDefaults instantiates a new OrderFreeItem object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *OrderFreeItem) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *OrderFreeItem) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *OrderFreeItem) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *OrderFreeItem) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *OrderFreeItem) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *OrderFreeItem) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *OrderFreeItem) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *OrderFreeItem) HasName() bool`

HasName returns a boolean if a field has been set.

### GetQuantity

`func (o *OrderFreeItem) GetQuantity() int32`

GetQuantity returns the Quantity field if non-nil, zero value otherwise.

### GetQuantityOk

`func (o *OrderFreeItem) GetQuantityOk() (*int32, bool)`

GetQuantityOk returns a tuple with the Quantity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuantity

`func (o *OrderFreeItem) SetQuantity(v int32)`

SetQuantity sets Quantity field to given value.

### HasQuantity

`func (o *OrderFreeItem) HasQuantity() bool`

HasQuantity returns a boolean if a field has been set.

### GetPrice

`func (o *OrderFreeItem) GetPrice() int64`

GetPrice returns the Price field if non-nil, zero value otherwise.

### GetPriceOk

`func (o *OrderFreeItem) GetPriceOk() (*int64, bool)`

GetPriceOk returns a tuple with the Price field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrice

`func (o *OrderFreeItem) SetPrice(v int64)`

SetPrice sets Price field to given value.

### HasPrice

`func (o *OrderFreeItem) HasPrice() bool`

HasPrice returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


