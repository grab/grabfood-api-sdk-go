# UpdateSpecialHourRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SpecialOpeningHours** | [**[]SpecialOpeningHour**](SpecialOpeningHour.md) | An array of objects contain store special hours. Max. 3 array elements. | 

## Methods

### NewUpdateSpecialHourRequest

`func NewUpdateSpecialHourRequest(specialOpeningHours []SpecialOpeningHour, ) *UpdateSpecialHourRequest`

NewUpdateSpecialHourRequest instantiates a new UpdateSpecialHourRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateSpecialHourRequestWithDefaults

`func NewUpdateSpecialHourRequestWithDefaults() *UpdateSpecialHourRequest`

NewUpdateSpecialHourRequestWithDefaults instantiates a new UpdateSpecialHourRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSpecialOpeningHours

`func (o *UpdateSpecialHourRequest) GetSpecialOpeningHours() []SpecialOpeningHour`

GetSpecialOpeningHours returns the SpecialOpeningHours field if non-nil, zero value otherwise.

### GetSpecialOpeningHoursOk

`func (o *UpdateSpecialHourRequest) GetSpecialOpeningHoursOk() (*[]SpecialOpeningHour, bool)`

GetSpecialOpeningHoursOk returns a tuple with the SpecialOpeningHours field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpecialOpeningHours

`func (o *UpdateSpecialHourRequest) SetSpecialOpeningHours(v []SpecialOpeningHour)`

SetSpecialOpeningHours sets SpecialOpeningHours field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


