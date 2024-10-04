# EditOrderItem

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ItemID** | **string** | The item&#39;s ID in Grab system that can be obtained from the [Submit Order Webhook](#tag/submit-order-webhook/operation/submit-order-webhook) request payload parameters under &#x60;items[].grabItemID&#x60;, or &#x60;items[].outOfStockInstruction.replacementGrabItemID&#x60; for item replacement. External item ID from Partner system is only supported when &#x60;ADDED&#x60; status and &#x60;isExternalItemID: true&#x60;. | 
**Status** | **string** | The item&#39;s edited status. Leave empty string if there is no change to the item. | 
**Quantity** | Pointer to **int64** | The item&#39;s quantity. If the item is not being updated or deleted, use the original quantity. | [optional] 
**IsExternalItemID** | Pointer to **bool** | Only applicable for &#x60;ADDED&#x60;status. Indicate if the &#x60;itemID&#x60; is an external item ID. Grab checks for the items that are mapped to the provided item ID, considering their availability. If multiple Grab items are found to be mapped to the provided external item ID, the last updated item will be chosen. If no suitable record is found, an 400 error will be returned to the partner, indicating that the submitted external item ID cannot be edited. | [optional] 

## Methods

### NewEditOrderItem

`func NewEditOrderItem(itemID string, status string, ) *EditOrderItem`

NewEditOrderItem instantiates a new EditOrderItem object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEditOrderItemWithDefaults

`func NewEditOrderItemWithDefaults() *EditOrderItem`

NewEditOrderItemWithDefaults instantiates a new EditOrderItem object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetItemID

`func (o *EditOrderItem) GetItemID() string`

GetItemID returns the ItemID field if non-nil, zero value otherwise.

### GetItemIDOk

`func (o *EditOrderItem) GetItemIDOk() (*string, bool)`

GetItemIDOk returns a tuple with the ItemID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItemID

`func (o *EditOrderItem) SetItemID(v string)`

SetItemID sets ItemID field to given value.


### GetStatus

`func (o *EditOrderItem) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *EditOrderItem) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *EditOrderItem) SetStatus(v string)`

SetStatus sets Status field to given value.


### GetQuantity

`func (o *EditOrderItem) GetQuantity() int64`

GetQuantity returns the Quantity field if non-nil, zero value otherwise.

### GetQuantityOk

`func (o *EditOrderItem) GetQuantityOk() (*int64, bool)`

GetQuantityOk returns a tuple with the Quantity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuantity

`func (o *EditOrderItem) SetQuantity(v int64)`

SetQuantity sets Quantity field to given value.

### HasQuantity

`func (o *EditOrderItem) HasQuantity() bool`

HasQuantity returns a boolean if a field has been set.

### GetIsExternalItemID

`func (o *EditOrderItem) GetIsExternalItemID() bool`

GetIsExternalItemID returns the IsExternalItemID field if non-nil, zero value otherwise.

### GetIsExternalItemIDOk

`func (o *EditOrderItem) GetIsExternalItemIDOk() (*bool, bool)`

GetIsExternalItemIDOk returns a tuple with the IsExternalItemID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsExternalItemID

`func (o *EditOrderItem) SetIsExternalItemID(v bool)`

SetIsExternalItemID sets IsExternalItemID field to given value.

### HasIsExternalItemID

`func (o *EditOrderItem) HasIsExternalItemID() bool`

HasIsExternalItemID returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


