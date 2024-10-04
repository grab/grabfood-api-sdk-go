# AcceptOrderRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**OrderID** | **string** | The order&#39;s ID that is returned from GrabFood. Refer to FAQs for more details about [orderID and shortOrderNumber](#section/Order/What&#39;s-the-difference-between-orderID-and-shortOrderNumber). | 
**ToState** | **string** | The order&#39;s updated state. | 

## Methods

### NewAcceptOrderRequest

`func NewAcceptOrderRequest(orderID string, toState string, ) *AcceptOrderRequest`

NewAcceptOrderRequest instantiates a new AcceptOrderRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAcceptOrderRequestWithDefaults

`func NewAcceptOrderRequestWithDefaults() *AcceptOrderRequest`

NewAcceptOrderRequestWithDefaults instantiates a new AcceptOrderRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetOrderID

`func (o *AcceptOrderRequest) GetOrderID() string`

GetOrderID returns the OrderID field if non-nil, zero value otherwise.

### GetOrderIDOk

`func (o *AcceptOrderRequest) GetOrderIDOk() (*string, bool)`

GetOrderIDOk returns a tuple with the OrderID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrderID

`func (o *AcceptOrderRequest) SetOrderID(v string)`

SetOrderID sets OrderID field to given value.


### GetToState

`func (o *AcceptOrderRequest) GetToState() string`

GetToState returns the ToState field if non-nil, zero value otherwise.

### GetToStateOk

`func (o *AcceptOrderRequest) GetToStateOk() (*string, bool)`

GetToStateOk returns a tuple with the ToState field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToState

`func (o *AcceptOrderRequest) SetToState(v string)`

SetToState sets ToState field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


