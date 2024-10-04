# OrderItem

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | The item&#39;s ExternalID in the partner system.  | 
**GrabItemID** | **string** | The item&#39;s ID in Grab system. Partner can use this field in the &#x60;EditOrder&#x60; endpoint. | 
**Quantity** | **int32** | The number of the item ordered. | 
**Price** | **int64** | The price (tax-inclusive) in minor format for 1 item and modifiers under it. &#x60;Item price(tax inclusive) + Modifier price(tax inclusive) | (2241*1.06)+(165*1.06)&#x3D;2550&#x60;.  | 
**Tax** | Pointer to **int64** | Tax in minor format for 1 item and all modifiers under it. &#x60;0&#x60; if tax configuration is absent. &#x60;Item tax + Modifier tax | (2241*0.06)+(165*0.06)&#x3D;144&#x60;. Refer to FAQs for more details about [tax](#section/Order/How-is-tax-calculated). | [optional] 
**Specifications** | Pointer to **string** | An extra note for the merchant. &#x60;Blank&#x60; if no note from consumer.  | [optional] 
**OutOfStockInstruction** | Pointer to [**NullableOutOfStockInstruction**](OutOfStockInstruction.md) |  | [optional] 
**Modifiers** | Pointer to [**[]OrderItemModifier**](OrderItemModifier.md) | An array of JSON objects modifiers. Read [this](#categories) for more information. | [optional] 

## Methods

### NewOrderItem

`func NewOrderItem(id string, grabItemID string, quantity int32, price int64, ) *OrderItem`

NewOrderItem instantiates a new OrderItem object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOrderItemWithDefaults

`func NewOrderItemWithDefaults() *OrderItem`

NewOrderItemWithDefaults instantiates a new OrderItem object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *OrderItem) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *OrderItem) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *OrderItem) SetId(v string)`

SetId sets Id field to given value.


### GetGrabItemID

`func (o *OrderItem) GetGrabItemID() string`

GetGrabItemID returns the GrabItemID field if non-nil, zero value otherwise.

### GetGrabItemIDOk

`func (o *OrderItem) GetGrabItemIDOk() (*string, bool)`

GetGrabItemIDOk returns a tuple with the GrabItemID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGrabItemID

`func (o *OrderItem) SetGrabItemID(v string)`

SetGrabItemID sets GrabItemID field to given value.


### GetQuantity

`func (o *OrderItem) GetQuantity() int32`

GetQuantity returns the Quantity field if non-nil, zero value otherwise.

### GetQuantityOk

`func (o *OrderItem) GetQuantityOk() (*int32, bool)`

GetQuantityOk returns a tuple with the Quantity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuantity

`func (o *OrderItem) SetQuantity(v int32)`

SetQuantity sets Quantity field to given value.


### GetPrice

`func (o *OrderItem) GetPrice() int64`

GetPrice returns the Price field if non-nil, zero value otherwise.

### GetPriceOk

`func (o *OrderItem) GetPriceOk() (*int64, bool)`

GetPriceOk returns a tuple with the Price field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrice

`func (o *OrderItem) SetPrice(v int64)`

SetPrice sets Price field to given value.


### GetTax

`func (o *OrderItem) GetTax() int64`

GetTax returns the Tax field if non-nil, zero value otherwise.

### GetTaxOk

`func (o *OrderItem) GetTaxOk() (*int64, bool)`

GetTaxOk returns a tuple with the Tax field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTax

`func (o *OrderItem) SetTax(v int64)`

SetTax sets Tax field to given value.

### HasTax

`func (o *OrderItem) HasTax() bool`

HasTax returns a boolean if a field has been set.

### GetSpecifications

`func (o *OrderItem) GetSpecifications() string`

GetSpecifications returns the Specifications field if non-nil, zero value otherwise.

### GetSpecificationsOk

`func (o *OrderItem) GetSpecificationsOk() (*string, bool)`

GetSpecificationsOk returns a tuple with the Specifications field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpecifications

`func (o *OrderItem) SetSpecifications(v string)`

SetSpecifications sets Specifications field to given value.

### HasSpecifications

`func (o *OrderItem) HasSpecifications() bool`

HasSpecifications returns a boolean if a field has been set.

### GetOutOfStockInstruction

`func (o *OrderItem) GetOutOfStockInstruction() OutOfStockInstruction`

GetOutOfStockInstruction returns the OutOfStockInstruction field if non-nil, zero value otherwise.

### GetOutOfStockInstructionOk

`func (o *OrderItem) GetOutOfStockInstructionOk() (*OutOfStockInstruction, bool)`

GetOutOfStockInstructionOk returns a tuple with the OutOfStockInstruction field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOutOfStockInstruction

`func (o *OrderItem) SetOutOfStockInstruction(v OutOfStockInstruction)`

SetOutOfStockInstruction sets OutOfStockInstruction field to given value.

### HasOutOfStockInstruction

`func (o *OrderItem) HasOutOfStockInstruction() bool`

HasOutOfStockInstruction returns a boolean if a field has been set.

### SetOutOfStockInstructionNil

`func (o *OrderItem) SetOutOfStockInstructionNil(b bool)`

 SetOutOfStockInstructionNil sets the value for OutOfStockInstruction to be an explicit nil

### UnsetOutOfStockInstruction
`func (o *OrderItem) UnsetOutOfStockInstruction()`

UnsetOutOfStockInstruction ensures that no value is present for OutOfStockInstruction, not even an explicit nil
### GetModifiers

`func (o *OrderItem) GetModifiers() []OrderItemModifier`

GetModifiers returns the Modifiers field if non-nil, zero value otherwise.

### GetModifiersOk

`func (o *OrderItem) GetModifiersOk() (*[]OrderItemModifier, bool)`

GetModifiersOk returns a tuple with the Modifiers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModifiers

`func (o *OrderItem) SetModifiers(v []OrderItemModifier)`

SetModifiers sets Modifiers field to given value.

### HasModifiers

`func (o *OrderItem) HasModifiers() bool`

HasModifiers returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


