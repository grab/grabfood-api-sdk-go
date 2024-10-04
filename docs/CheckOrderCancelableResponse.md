# CheckOrderCancelableResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CancelAble** | Pointer to **bool** | The boolean value to indicate whether an order can be cancelled. | [optional] 
**NonCancellationReason** | Pointer to **string** | The reason for the order to be non-cancelable. | [optional] 
**LimitType** | Pointer to [**CancelOrderLimitType**](CancelOrderLimitType.md) |  | [optional] 
**LimitTimes** | Pointer to **int64** | The remaining cancellation quota for the merchant. A value is only returned when the nearest remaining cancellation limit is approaching, else it returns 0. | [optional] 
**CancelReasons** | Pointer to [**[]CancelReason**](CancelReason.md) | An array of cancel order reasons JSON objects. | [optional] 

## Methods

### NewCheckOrderCancelableResponse

`func NewCheckOrderCancelableResponse() *CheckOrderCancelableResponse`

NewCheckOrderCancelableResponse instantiates a new CheckOrderCancelableResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCheckOrderCancelableResponseWithDefaults

`func NewCheckOrderCancelableResponseWithDefaults() *CheckOrderCancelableResponse`

NewCheckOrderCancelableResponseWithDefaults instantiates a new CheckOrderCancelableResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCancelAble

`func (o *CheckOrderCancelableResponse) GetCancelAble() bool`

GetCancelAble returns the CancelAble field if non-nil, zero value otherwise.

### GetCancelAbleOk

`func (o *CheckOrderCancelableResponse) GetCancelAbleOk() (*bool, bool)`

GetCancelAbleOk returns a tuple with the CancelAble field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCancelAble

`func (o *CheckOrderCancelableResponse) SetCancelAble(v bool)`

SetCancelAble sets CancelAble field to given value.

### HasCancelAble

`func (o *CheckOrderCancelableResponse) HasCancelAble() bool`

HasCancelAble returns a boolean if a field has been set.

### GetNonCancellationReason

`func (o *CheckOrderCancelableResponse) GetNonCancellationReason() string`

GetNonCancellationReason returns the NonCancellationReason field if non-nil, zero value otherwise.

### GetNonCancellationReasonOk

`func (o *CheckOrderCancelableResponse) GetNonCancellationReasonOk() (*string, bool)`

GetNonCancellationReasonOk returns a tuple with the NonCancellationReason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNonCancellationReason

`func (o *CheckOrderCancelableResponse) SetNonCancellationReason(v string)`

SetNonCancellationReason sets NonCancellationReason field to given value.

### HasNonCancellationReason

`func (o *CheckOrderCancelableResponse) HasNonCancellationReason() bool`

HasNonCancellationReason returns a boolean if a field has been set.

### GetLimitType

`func (o *CheckOrderCancelableResponse) GetLimitType() CancelOrderLimitType`

GetLimitType returns the LimitType field if non-nil, zero value otherwise.

### GetLimitTypeOk

`func (o *CheckOrderCancelableResponse) GetLimitTypeOk() (*CancelOrderLimitType, bool)`

GetLimitTypeOk returns a tuple with the LimitType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLimitType

`func (o *CheckOrderCancelableResponse) SetLimitType(v CancelOrderLimitType)`

SetLimitType sets LimitType field to given value.

### HasLimitType

`func (o *CheckOrderCancelableResponse) HasLimitType() bool`

HasLimitType returns a boolean if a field has been set.

### GetLimitTimes

`func (o *CheckOrderCancelableResponse) GetLimitTimes() int64`

GetLimitTimes returns the LimitTimes field if non-nil, zero value otherwise.

### GetLimitTimesOk

`func (o *CheckOrderCancelableResponse) GetLimitTimesOk() (*int64, bool)`

GetLimitTimesOk returns a tuple with the LimitTimes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLimitTimes

`func (o *CheckOrderCancelableResponse) SetLimitTimes(v int64)`

SetLimitTimes sets LimitTimes field to given value.

### HasLimitTimes

`func (o *CheckOrderCancelableResponse) HasLimitTimes() bool`

HasLimitTimes returns a boolean if a field has been set.

### GetCancelReasons

`func (o *CheckOrderCancelableResponse) GetCancelReasons() []CancelReason`

GetCancelReasons returns the CancelReasons field if non-nil, zero value otherwise.

### GetCancelReasonsOk

`func (o *CheckOrderCancelableResponse) GetCancelReasonsOk() (*[]CancelReason, bool)`

GetCancelReasonsOk returns a tuple with the CancelReasons field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCancelReasons

`func (o *CheckOrderCancelableResponse) SetCancelReasons(v []CancelReason)`

SetCancelReasons sets CancelReasons field to given value.

### HasCancelReasons

`func (o *CheckOrderCancelableResponse) HasCancelReasons() bool`

HasCancelReasons returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


