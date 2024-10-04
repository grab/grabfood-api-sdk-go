# OrderStateRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**MerchantID** | **string** | The merchant&#39;s ID that is in GrabFood&#39;s database. | 
**OrderID** | **string** | The order&#39;s ID that is returned from GrabFood. Refer to FAQs for more details about [orderID and shortOrderNumber](#section/Order/What&#39;s-the-difference-between-orderID-and-shortOrderNumber). | 
**State** | **string** | The current order state. For takeaway orders, only &#x60;DELIVERED&#x60; and &#x60;CANCELLED&#x60; states are pushed. | 
**DriverETA** | Pointer to **NullableInt32** | The driver&#39;s estimated of arrival (ETA) in seconds when the state is &#x60;DRIVER_ALLOCATED&#x60;. | [optional] 
**Code** | Pointer to **string** | The current order&#39;s sub-state. This is in free text so you should only use for reference. Grab may use this for troubleshooting. If you want some analysis, kindly use &#x60;state&#x60; instead. | [optional] 
**Message** | Pointer to **string** | Additional information to explain the current order state. May be system status or human entered message. | [optional] 

## Methods

### NewOrderStateRequest

`func NewOrderStateRequest(merchantID string, orderID string, state string, ) *OrderStateRequest`

NewOrderStateRequest instantiates a new OrderStateRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOrderStateRequestWithDefaults

`func NewOrderStateRequestWithDefaults() *OrderStateRequest`

NewOrderStateRequestWithDefaults instantiates a new OrderStateRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMerchantID

`func (o *OrderStateRequest) GetMerchantID() string`

GetMerchantID returns the MerchantID field if non-nil, zero value otherwise.

### GetMerchantIDOk

`func (o *OrderStateRequest) GetMerchantIDOk() (*string, bool)`

GetMerchantIDOk returns a tuple with the MerchantID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMerchantID

`func (o *OrderStateRequest) SetMerchantID(v string)`

SetMerchantID sets MerchantID field to given value.


### GetOrderID

`func (o *OrderStateRequest) GetOrderID() string`

GetOrderID returns the OrderID field if non-nil, zero value otherwise.

### GetOrderIDOk

`func (o *OrderStateRequest) GetOrderIDOk() (*string, bool)`

GetOrderIDOk returns a tuple with the OrderID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrderID

`func (o *OrderStateRequest) SetOrderID(v string)`

SetOrderID sets OrderID field to given value.


### GetState

`func (o *OrderStateRequest) GetState() string`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *OrderStateRequest) GetStateOk() (*string, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *OrderStateRequest) SetState(v string)`

SetState sets State field to given value.


### GetDriverETA

`func (o *OrderStateRequest) GetDriverETA() int32`

GetDriverETA returns the DriverETA field if non-nil, zero value otherwise.

### GetDriverETAOk

`func (o *OrderStateRequest) GetDriverETAOk() (*int32, bool)`

GetDriverETAOk returns a tuple with the DriverETA field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDriverETA

`func (o *OrderStateRequest) SetDriverETA(v int32)`

SetDriverETA sets DriverETA field to given value.

### HasDriverETA

`func (o *OrderStateRequest) HasDriverETA() bool`

HasDriverETA returns a boolean if a field has been set.

### SetDriverETANil

`func (o *OrderStateRequest) SetDriverETANil(b bool)`

 SetDriverETANil sets the value for DriverETA to be an explicit nil

### UnsetDriverETA
`func (o *OrderStateRequest) UnsetDriverETA()`

UnsetDriverETA ensures that no value is present for DriverETA, not even an explicit nil
### GetCode

`func (o *OrderStateRequest) GetCode() string`

GetCode returns the Code field if non-nil, zero value otherwise.

### GetCodeOk

`func (o *OrderStateRequest) GetCodeOk() (*string, bool)`

GetCodeOk returns a tuple with the Code field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCode

`func (o *OrderStateRequest) SetCode(v string)`

SetCode sets Code field to given value.

### HasCode

`func (o *OrderStateRequest) HasCode() bool`

HasCode returns a boolean if a field has been set.

### GetMessage

`func (o *OrderStateRequest) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *OrderStateRequest) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *OrderStateRequest) SetMessage(v string)`

SetMessage sets Message field to given value.

### HasMessage

`func (o *OrderStateRequest) HasMessage() bool`

HasMessage returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


