# EditOrderRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**OrderID** | **string** | The order&#39;s ID that is returned from GrabFood. Refer to FAQs for more details about [orderID and shortOrderNumber](#section/Order/What&#39;s-the-difference-between-orderID-and-shortOrderNumber). | 
**Items** | [**[]EditOrderItem**](EditOrderItem.md) | Specify the array of all items in the order, including deleted, added, updated and unchanged items. | 
**OnlyRecalculate** | Pointer to **bool** | This parameter specifies whether to recalculate the edited order without submitting it. It is intended for testing purposes only. This parameter is set to false by default, which means the edited order will be recalculated and re-submitted to partners.  | [optional] 

## Methods

### NewEditOrderRequest

`func NewEditOrderRequest(orderID string, items []EditOrderItem, ) *EditOrderRequest`

NewEditOrderRequest instantiates a new EditOrderRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEditOrderRequestWithDefaults

`func NewEditOrderRequestWithDefaults() *EditOrderRequest`

NewEditOrderRequestWithDefaults instantiates a new EditOrderRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetOrderID

`func (o *EditOrderRequest) GetOrderID() string`

GetOrderID returns the OrderID field if non-nil, zero value otherwise.

### GetOrderIDOk

`func (o *EditOrderRequest) GetOrderIDOk() (*string, bool)`

GetOrderIDOk returns a tuple with the OrderID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrderID

`func (o *EditOrderRequest) SetOrderID(v string)`

SetOrderID sets OrderID field to given value.


### GetItems

`func (o *EditOrderRequest) GetItems() []EditOrderItem`

GetItems returns the Items field if non-nil, zero value otherwise.

### GetItemsOk

`func (o *EditOrderRequest) GetItemsOk() (*[]EditOrderItem, bool)`

GetItemsOk returns a tuple with the Items field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItems

`func (o *EditOrderRequest) SetItems(v []EditOrderItem)`

SetItems sets Items field to given value.


### GetOnlyRecalculate

`func (o *EditOrderRequest) GetOnlyRecalculate() bool`

GetOnlyRecalculate returns the OnlyRecalculate field if non-nil, zero value otherwise.

### GetOnlyRecalculateOk

`func (o *EditOrderRequest) GetOnlyRecalculateOk() (*bool, bool)`

GetOnlyRecalculateOk returns a tuple with the OnlyRecalculate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOnlyRecalculate

`func (o *EditOrderRequest) SetOnlyRecalculate(v bool)`

SetOnlyRecalculate sets OnlyRecalculate field to given value.

### HasOnlyRecalculate

`func (o *EditOrderRequest) HasOnlyRecalculate() bool`

HasOnlyRecalculate returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


