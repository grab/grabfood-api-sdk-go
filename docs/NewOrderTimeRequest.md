# NewOrderTimeRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**OrderID** | **string** | The order&#39;s ID that is returned from GrabFood. Refer to FAQs for more details about [orderID and shortOrderNumber](#section/Order/What&#39;s-the-difference-between-orderID-and-shortOrderNumber). | 
**NewOrderReadyTime** | **time.Time** | The new order ready time for this order, based on ISO_8601/RFC3339. | 

## Methods

### NewNewOrderTimeRequest

`func NewNewOrderTimeRequest(orderID string, newOrderReadyTime time.Time, ) *NewOrderTimeRequest`

NewNewOrderTimeRequest instantiates a new NewOrderTimeRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNewOrderTimeRequestWithDefaults

`func NewNewOrderTimeRequestWithDefaults() *NewOrderTimeRequest`

NewNewOrderTimeRequestWithDefaults instantiates a new NewOrderTimeRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetOrderID

`func (o *NewOrderTimeRequest) GetOrderID() string`

GetOrderID returns the OrderID field if non-nil, zero value otherwise.

### GetOrderIDOk

`func (o *NewOrderTimeRequest) GetOrderIDOk() (*string, bool)`

GetOrderIDOk returns a tuple with the OrderID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrderID

`func (o *NewOrderTimeRequest) SetOrderID(v string)`

SetOrderID sets OrderID field to given value.


### GetNewOrderReadyTime

`func (o *NewOrderTimeRequest) GetNewOrderReadyTime() time.Time`

GetNewOrderReadyTime returns the NewOrderReadyTime field if non-nil, zero value otherwise.

### GetNewOrderReadyTimeOk

`func (o *NewOrderTimeRequest) GetNewOrderReadyTimeOk() (*time.Time, bool)`

GetNewOrderReadyTimeOk returns a tuple with the NewOrderReadyTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNewOrderReadyTime

`func (o *NewOrderTimeRequest) SetNewOrderReadyTime(v time.Time)`

SetNewOrderReadyTime sets NewOrderReadyTime field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


