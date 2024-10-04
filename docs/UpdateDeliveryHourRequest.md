# UpdateDeliveryHourRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**OpeningHour** | [**StoreHour**](StoreHour.md) |  | 
**Force** | **bool** | To enable force update store delivery hours. Error will be returned if set to false while there is ongoing order. | 

## Methods

### NewUpdateDeliveryHourRequest

`func NewUpdateDeliveryHourRequest(openingHour StoreHour, force bool, ) *UpdateDeliveryHourRequest`

NewUpdateDeliveryHourRequest instantiates a new UpdateDeliveryHourRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateDeliveryHourRequestWithDefaults

`func NewUpdateDeliveryHourRequestWithDefaults() *UpdateDeliveryHourRequest`

NewUpdateDeliveryHourRequestWithDefaults instantiates a new UpdateDeliveryHourRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetOpeningHour

`func (o *UpdateDeliveryHourRequest) GetOpeningHour() StoreHour`

GetOpeningHour returns the OpeningHour field if non-nil, zero value otherwise.

### GetOpeningHourOk

`func (o *UpdateDeliveryHourRequest) GetOpeningHourOk() (*StoreHour, bool)`

GetOpeningHourOk returns a tuple with the OpeningHour field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOpeningHour

`func (o *UpdateDeliveryHourRequest) SetOpeningHour(v StoreHour)`

SetOpeningHour sets OpeningHour field to given value.


### GetForce

`func (o *UpdateDeliveryHourRequest) GetForce() bool`

GetForce returns the Force field if non-nil, zero value otherwise.

### GetForceOk

`func (o *UpdateDeliveryHourRequest) GetForceOk() (*bool, bool)`

GetForceOk returns a tuple with the Force field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetForce

`func (o *UpdateDeliveryHourRequest) SetForce(v bool)`

SetForce sets Force field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


