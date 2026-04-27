# POSItem

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | The item&#39;s externalID in the partner system.  | [optional] 
**GrabItemID** | Pointer to **string** | The item&#39;s ID in Grab system. Partner can use this field in the &#x60;EditOrder&#x60; endpoint. Note: The index number (after &#x60;#&#x60;) is different for the same item with different modifiers. This helps identify items when editing complex orders. This is currently controlled by feature flag until full rollout, for non whitelisted partners, the &#39;#&#39; and index number will not be included.  | [optional] 
**Name** | Pointer to **string** | The name of the item. | [optional] 
**Quantity** | Pointer to **int32** | The number of the item ordered. | [optional] 
**Modifiers** | Pointer to [**[]PosItemModifier**](PosItemModifier.md) | The ordered items in an array of JSON Object.  | [optional] 
**Price** | Pointer to **int64** | The price for a single item along with its associated modifiers in minor unit and tax-inclusive.  &#x60;&#x60;&#x60; price &#x3D; Item price(tax inclusive) + Modifier price(tax inclusive) | (2241*1.06)+(165*1.06)&#x3D;2550  | [optional] 
**Tax** | Pointer to **int64** | Tax in minor format for a single item along with its associated modifiers. &#x60;0&#x60; if tax configuration is absent. Refer to FAQs for more details about [tax](#section/Order/How-is-tax-calculated). &#x60;&#x60;&#x60; tax &#x3D; Item tax + Modifier tax | (2241*0.06)+(165*0.06)&#x3D;144  | [optional] 

## Methods

### NewPOSItem

`func NewPOSItem() *POSItem`

NewPOSItem instantiates a new POSItem object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPOSItemWithDefaults

`func NewPOSItemWithDefaults() *POSItem`

NewPOSItemWithDefaults instantiates a new POSItem object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *POSItem) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *POSItem) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *POSItem) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *POSItem) HasId() bool`

HasId returns a boolean if a field has been set.

### GetGrabItemID

`func (o *POSItem) GetGrabItemID() string`

GetGrabItemID returns the GrabItemID field if non-nil, zero value otherwise.

### GetGrabItemIDOk

`func (o *POSItem) GetGrabItemIDOk() (*string, bool)`

GetGrabItemIDOk returns a tuple with the GrabItemID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGrabItemID

`func (o *POSItem) SetGrabItemID(v string)`

SetGrabItemID sets GrabItemID field to given value.

### HasGrabItemID

`func (o *POSItem) HasGrabItemID() bool`

HasGrabItemID returns a boolean if a field has been set.

### GetName

`func (o *POSItem) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *POSItem) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *POSItem) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *POSItem) HasName() bool`

HasName returns a boolean if a field has been set.

### GetQuantity

`func (o *POSItem) GetQuantity() int32`

GetQuantity returns the Quantity field if non-nil, zero value otherwise.

### GetQuantityOk

`func (o *POSItem) GetQuantityOk() (*int32, bool)`

GetQuantityOk returns a tuple with the Quantity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuantity

`func (o *POSItem) SetQuantity(v int32)`

SetQuantity sets Quantity field to given value.

### HasQuantity

`func (o *POSItem) HasQuantity() bool`

HasQuantity returns a boolean if a field has been set.

### GetModifiers

`func (o *POSItem) GetModifiers() []PosItemModifier`

GetModifiers returns the Modifiers field if non-nil, zero value otherwise.

### GetModifiersOk

`func (o *POSItem) GetModifiersOk() (*[]PosItemModifier, bool)`

GetModifiersOk returns a tuple with the Modifiers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModifiers

`func (o *POSItem) SetModifiers(v []PosItemModifier)`

SetModifiers sets Modifiers field to given value.

### HasModifiers

`func (o *POSItem) HasModifiers() bool`

HasModifiers returns a boolean if a field has been set.

### GetPrice

`func (o *POSItem) GetPrice() int64`

GetPrice returns the Price field if non-nil, zero value otherwise.

### GetPriceOk

`func (o *POSItem) GetPriceOk() (*int64, bool)`

GetPriceOk returns a tuple with the Price field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrice

`func (o *POSItem) SetPrice(v int64)`

SetPrice sets Price field to given value.

### HasPrice

`func (o *POSItem) HasPrice() bool`

HasPrice returns a boolean if a field has been set.

### GetTax

`func (o *POSItem) GetTax() int64`

GetTax returns the Tax field if non-nil, zero value otherwise.

### GetTaxOk

`func (o *POSItem) GetTaxOk() (*int64, bool)`

GetTaxOk returns a tuple with the Tax field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTax

`func (o *POSItem) SetTax(v int64)`

SetTax sets Tax field to given value.

### HasTax

`func (o *POSItem) HasTax() bool`

HasTax returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


