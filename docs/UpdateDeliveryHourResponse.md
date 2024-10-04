# UpdateDeliveryHourResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ErrorReasons** | Pointer to **[]string** | Error message when updating store delivery hour. &#x60;null&#x60; indicates no error. | [optional] 
**OrderCount** | **int64** | Total active order for the day. | 
**ScheduledOrderCount** | **int64** | Total scheduled order during store close period. | 
**CloseImmediately** | **bool** | Indicate the store status after updating delivery hours. | 

## Methods

### NewUpdateDeliveryHourResponse

`func NewUpdateDeliveryHourResponse(orderCount int64, scheduledOrderCount int64, closeImmediately bool, ) *UpdateDeliveryHourResponse`

NewUpdateDeliveryHourResponse instantiates a new UpdateDeliveryHourResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateDeliveryHourResponseWithDefaults

`func NewUpdateDeliveryHourResponseWithDefaults() *UpdateDeliveryHourResponse`

NewUpdateDeliveryHourResponseWithDefaults instantiates a new UpdateDeliveryHourResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetErrorReasons

`func (o *UpdateDeliveryHourResponse) GetErrorReasons() []string`

GetErrorReasons returns the ErrorReasons field if non-nil, zero value otherwise.

### GetErrorReasonsOk

`func (o *UpdateDeliveryHourResponse) GetErrorReasonsOk() (*[]string, bool)`

GetErrorReasonsOk returns a tuple with the ErrorReasons field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrorReasons

`func (o *UpdateDeliveryHourResponse) SetErrorReasons(v []string)`

SetErrorReasons sets ErrorReasons field to given value.

### HasErrorReasons

`func (o *UpdateDeliveryHourResponse) HasErrorReasons() bool`

HasErrorReasons returns a boolean if a field has been set.

### SetErrorReasonsNil

`func (o *UpdateDeliveryHourResponse) SetErrorReasonsNil(b bool)`

 SetErrorReasonsNil sets the value for ErrorReasons to be an explicit nil

### UnsetErrorReasons
`func (o *UpdateDeliveryHourResponse) UnsetErrorReasons()`

UnsetErrorReasons ensures that no value is present for ErrorReasons, not even an explicit nil
### GetOrderCount

`func (o *UpdateDeliveryHourResponse) GetOrderCount() int64`

GetOrderCount returns the OrderCount field if non-nil, zero value otherwise.

### GetOrderCountOk

`func (o *UpdateDeliveryHourResponse) GetOrderCountOk() (*int64, bool)`

GetOrderCountOk returns a tuple with the OrderCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrderCount

`func (o *UpdateDeliveryHourResponse) SetOrderCount(v int64)`

SetOrderCount sets OrderCount field to given value.


### GetScheduledOrderCount

`func (o *UpdateDeliveryHourResponse) GetScheduledOrderCount() int64`

GetScheduledOrderCount returns the ScheduledOrderCount field if non-nil, zero value otherwise.

### GetScheduledOrderCountOk

`func (o *UpdateDeliveryHourResponse) GetScheduledOrderCountOk() (*int64, bool)`

GetScheduledOrderCountOk returns a tuple with the ScheduledOrderCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScheduledOrderCount

`func (o *UpdateDeliveryHourResponse) SetScheduledOrderCount(v int64)`

SetScheduledOrderCount sets ScheduledOrderCount field to given value.


### GetCloseImmediately

`func (o *UpdateDeliveryHourResponse) GetCloseImmediately() bool`

GetCloseImmediately returns the CloseImmediately field if non-nil, zero value otherwise.

### GetCloseImmediatelyOk

`func (o *UpdateDeliveryHourResponse) GetCloseImmediatelyOk() (*bool, bool)`

GetCloseImmediatelyOk returns a tuple with the CloseImmediately field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCloseImmediately

`func (o *UpdateDeliveryHourResponse) SetCloseImmediately(v bool)`

SetCloseImmediately sets CloseImmediately field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


