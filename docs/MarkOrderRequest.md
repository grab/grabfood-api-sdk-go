# MarkOrderRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**OrderID** | **string** | The order&#39;s ID that is returned from GrabFood. Refer to FAQs for more details about [orderID and shortOrderNumber](#section/Order/What&#39;s-the-difference-between-orderID-and-shortOrderNumber). | 
**MarkStatus** | **int32** | The status to be marked accordingly.  * &#x60;1&#x60; - mark order as ready  * &#x60;2&#x60; - mark order as completed and only applicable to **dine-in** orders  | 

## Methods

### NewMarkOrderRequest

`func NewMarkOrderRequest(orderID string, markStatus int32, ) *MarkOrderRequest`

NewMarkOrderRequest instantiates a new MarkOrderRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMarkOrderRequestWithDefaults

`func NewMarkOrderRequestWithDefaults() *MarkOrderRequest`

NewMarkOrderRequestWithDefaults instantiates a new MarkOrderRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetOrderID

`func (o *MarkOrderRequest) GetOrderID() string`

GetOrderID returns the OrderID field if non-nil, zero value otherwise.

### GetOrderIDOk

`func (o *MarkOrderRequest) GetOrderIDOk() (*string, bool)`

GetOrderIDOk returns a tuple with the OrderID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrderID

`func (o *MarkOrderRequest) SetOrderID(v string)`

SetOrderID sets OrderID field to given value.


### GetMarkStatus

`func (o *MarkOrderRequest) GetMarkStatus() int32`

GetMarkStatus returns the MarkStatus field if non-nil, zero value otherwise.

### GetMarkStatusOk

`func (o *MarkOrderRequest) GetMarkStatusOk() (*int32, bool)`

GetMarkStatusOk returns a tuple with the MarkStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMarkStatus

`func (o *MarkOrderRequest) SetMarkStatus(v int32)`

SetMarkStatus sets MarkStatus field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


