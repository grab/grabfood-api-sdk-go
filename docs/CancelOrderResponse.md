# CancelOrderResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**LimitType** | Pointer to [**CancelOrderLimitType**](CancelOrderLimitType.md) |  | [optional] 
**LimitTimes** | Pointer to **int64** | The remaining cancellation quota for the merchant. A value is only returned when the nearest remaining cancellation limit is approaching, else it returns 0. | [optional] 

## Methods

### NewCancelOrderResponse

`func NewCancelOrderResponse() *CancelOrderResponse`

NewCancelOrderResponse instantiates a new CancelOrderResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCancelOrderResponseWithDefaults

`func NewCancelOrderResponseWithDefaults() *CancelOrderResponse`

NewCancelOrderResponseWithDefaults instantiates a new CancelOrderResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLimitType

`func (o *CancelOrderResponse) GetLimitType() CancelOrderLimitType`

GetLimitType returns the LimitType field if non-nil, zero value otherwise.

### GetLimitTypeOk

`func (o *CancelOrderResponse) GetLimitTypeOk() (*CancelOrderLimitType, bool)`

GetLimitTypeOk returns a tuple with the LimitType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLimitType

`func (o *CancelOrderResponse) SetLimitType(v CancelOrderLimitType)`

SetLimitType sets LimitType field to given value.

### HasLimitType

`func (o *CancelOrderResponse) HasLimitType() bool`

HasLimitType returns a boolean if a field has been set.

### GetLimitTimes

`func (o *CancelOrderResponse) GetLimitTimes() int64`

GetLimitTimes returns the LimitTimes field if non-nil, zero value otherwise.

### GetLimitTimesOk

`func (o *CancelOrderResponse) GetLimitTimesOk() (*int64, bool)`

GetLimitTimesOk returns a tuple with the LimitTimes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLimitTimes

`func (o *CancelOrderResponse) SetLimitTimes(v int64)`

SetLimitTimes sets LimitTimes field to given value.

### HasLimitTimes

`func (o *CancelOrderResponse) HasLimitTimes() bool`

HasLimitTimes returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


