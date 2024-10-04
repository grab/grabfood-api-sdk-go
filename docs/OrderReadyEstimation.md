# OrderReadyEstimation

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AllowChange** | **bool** | A boolean value indicating if the order ready time can be changed. | 
**EstimatedOrderReadyTime** | **time.Time** | The order ready time for this order estimated by GrabFood, based on ISO_8601/RFC3339. | 
**MaxOrderReadyTime** | **time.Time** | The max allowed order ready time for this order, based on ISO_8601/RFC3339. | 
**NewOrderReadyTime** | Pointer to **NullableTime** | The new order ready time for this order. Only present after a new order ready time is set (default will be null), based on ISO_8601/RFC3339. | [optional] 

## Methods

### NewOrderReadyEstimation

`func NewOrderReadyEstimation(allowChange bool, estimatedOrderReadyTime time.Time, maxOrderReadyTime time.Time, ) *OrderReadyEstimation`

NewOrderReadyEstimation instantiates a new OrderReadyEstimation object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOrderReadyEstimationWithDefaults

`func NewOrderReadyEstimationWithDefaults() *OrderReadyEstimation`

NewOrderReadyEstimationWithDefaults instantiates a new OrderReadyEstimation object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAllowChange

`func (o *OrderReadyEstimation) GetAllowChange() bool`

GetAllowChange returns the AllowChange field if non-nil, zero value otherwise.

### GetAllowChangeOk

`func (o *OrderReadyEstimation) GetAllowChangeOk() (*bool, bool)`

GetAllowChangeOk returns a tuple with the AllowChange field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowChange

`func (o *OrderReadyEstimation) SetAllowChange(v bool)`

SetAllowChange sets AllowChange field to given value.


### GetEstimatedOrderReadyTime

`func (o *OrderReadyEstimation) GetEstimatedOrderReadyTime() time.Time`

GetEstimatedOrderReadyTime returns the EstimatedOrderReadyTime field if non-nil, zero value otherwise.

### GetEstimatedOrderReadyTimeOk

`func (o *OrderReadyEstimation) GetEstimatedOrderReadyTimeOk() (*time.Time, bool)`

GetEstimatedOrderReadyTimeOk returns a tuple with the EstimatedOrderReadyTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEstimatedOrderReadyTime

`func (o *OrderReadyEstimation) SetEstimatedOrderReadyTime(v time.Time)`

SetEstimatedOrderReadyTime sets EstimatedOrderReadyTime field to given value.


### GetMaxOrderReadyTime

`func (o *OrderReadyEstimation) GetMaxOrderReadyTime() time.Time`

GetMaxOrderReadyTime returns the MaxOrderReadyTime field if non-nil, zero value otherwise.

### GetMaxOrderReadyTimeOk

`func (o *OrderReadyEstimation) GetMaxOrderReadyTimeOk() (*time.Time, bool)`

GetMaxOrderReadyTimeOk returns a tuple with the MaxOrderReadyTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxOrderReadyTime

`func (o *OrderReadyEstimation) SetMaxOrderReadyTime(v time.Time)`

SetMaxOrderReadyTime sets MaxOrderReadyTime field to given value.


### GetNewOrderReadyTime

`func (o *OrderReadyEstimation) GetNewOrderReadyTime() time.Time`

GetNewOrderReadyTime returns the NewOrderReadyTime field if non-nil, zero value otherwise.

### GetNewOrderReadyTimeOk

`func (o *OrderReadyEstimation) GetNewOrderReadyTimeOk() (*time.Time, bool)`

GetNewOrderReadyTimeOk returns a tuple with the NewOrderReadyTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNewOrderReadyTime

`func (o *OrderReadyEstimation) SetNewOrderReadyTime(v time.Time)`

SetNewOrderReadyTime sets NewOrderReadyTime field to given value.

### HasNewOrderReadyTime

`func (o *OrderReadyEstimation) HasNewOrderReadyTime() bool`

HasNewOrderReadyTime returns a boolean if a field has been set.

### SetNewOrderReadyTimeNil

`func (o *OrderReadyEstimation) SetNewOrderReadyTimeNil(b bool)`

 SetNewOrderReadyTimeNil sets the value for NewOrderReadyTime to be an explicit nil

### UnsetNewOrderReadyTime
`func (o *OrderReadyEstimation) UnsetNewOrderReadyTime()`

UnsetNewOrderReadyTime ensures that no value is present for NewOrderReadyTime, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


