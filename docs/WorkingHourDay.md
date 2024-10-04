# WorkingHourDay

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Periods** | Pointer to [**[]OpenPeriod**](OpenPeriod.md) | An array of open periods. &#x60;null&#x60; in [ListCampaign](#tag/list-campaign) response if the campaign is available all day. | [optional] 

## Methods

### NewWorkingHourDay

`func NewWorkingHourDay() *WorkingHourDay`

NewWorkingHourDay instantiates a new WorkingHourDay object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWorkingHourDayWithDefaults

`func NewWorkingHourDayWithDefaults() *WorkingHourDay`

NewWorkingHourDayWithDefaults instantiates a new WorkingHourDay object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPeriods

`func (o *WorkingHourDay) GetPeriods() []OpenPeriod`

GetPeriods returns the Periods field if non-nil, zero value otherwise.

### GetPeriodsOk

`func (o *WorkingHourDay) GetPeriodsOk() (*[]OpenPeriod, bool)`

GetPeriodsOk returns a tuple with the Periods field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPeriods

`func (o *WorkingHourDay) SetPeriods(v []OpenPeriod)`

SetPeriods sets Periods field to given value.

### HasPeriods

`func (o *WorkingHourDay) HasPeriods() bool`

HasPeriods returns a boolean if a field has been set.

### SetPeriodsNil

`func (o *WorkingHourDay) SetPeriodsNil(b bool)`

 SetPeriodsNil sets the value for Periods to be an explicit nil

### UnsetPeriods
`func (o *WorkingHourDay) UnsetPeriods()`

UnsetPeriods ensures that no value is present for Periods, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


