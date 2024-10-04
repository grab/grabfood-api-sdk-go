# OrderDeliveryRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**OrderID** | **string** | The order&#39;s ID that is returned from GrabFood. Refer to FAQs for more details about [orderID and shortOrderNumber](#section/Order/What&#39;s-the-difference-between-orderID-and-shortOrderNumber). | 
**FromState** | **string** | Specify the order&#39;s initial state. | 
**ToState** | **string** | Specify the order&#39;s new state. | 

## Methods

### NewOrderDeliveryRequest

`func NewOrderDeliveryRequest(orderID string, fromState string, toState string, ) *OrderDeliveryRequest`

NewOrderDeliveryRequest instantiates a new OrderDeliveryRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOrderDeliveryRequestWithDefaults

`func NewOrderDeliveryRequestWithDefaults() *OrderDeliveryRequest`

NewOrderDeliveryRequestWithDefaults instantiates a new OrderDeliveryRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetOrderID

`func (o *OrderDeliveryRequest) GetOrderID() string`

GetOrderID returns the OrderID field if non-nil, zero value otherwise.

### GetOrderIDOk

`func (o *OrderDeliveryRequest) GetOrderIDOk() (*string, bool)`

GetOrderIDOk returns a tuple with the OrderID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrderID

`func (o *OrderDeliveryRequest) SetOrderID(v string)`

SetOrderID sets OrderID field to given value.


### GetFromState

`func (o *OrderDeliveryRequest) GetFromState() string`

GetFromState returns the FromState field if non-nil, zero value otherwise.

### GetFromStateOk

`func (o *OrderDeliveryRequest) GetFromStateOk() (*string, bool)`

GetFromStateOk returns a tuple with the FromState field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFromState

`func (o *OrderDeliveryRequest) SetFromState(v string)`

SetFromState sets FromState field to given value.


### GetToState

`func (o *OrderDeliveryRequest) GetToState() string`

GetToState returns the ToState field if non-nil, zero value otherwise.

### GetToStateOk

`func (o *OrderDeliveryRequest) GetToStateOk() (*string, bool)`

GetToStateOk returns a tuple with the ToState field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToState

`func (o *OrderDeliveryRequest) SetToState(v string)`

SetToState sets ToState field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


