# SpecialOpeningHourOpeningHours

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**OpenPeriodType** | Pointer to **string** | The period type for when the outlet is open. | [optional] 
**Periods** | Pointer to [**[]OpenPeriod**](OpenPeriod.md) | An array of open periods. Maximum of 3 periods. | [optional] 

## Methods

### NewSpecialOpeningHourOpeningHours

`func NewSpecialOpeningHourOpeningHours() *SpecialOpeningHourOpeningHours`

NewSpecialOpeningHourOpeningHours instantiates a new SpecialOpeningHourOpeningHours object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSpecialOpeningHourOpeningHoursWithDefaults

`func NewSpecialOpeningHourOpeningHoursWithDefaults() *SpecialOpeningHourOpeningHours`

NewSpecialOpeningHourOpeningHoursWithDefaults instantiates a new SpecialOpeningHourOpeningHours object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetOpenPeriodType

`func (o *SpecialOpeningHourOpeningHours) GetOpenPeriodType() string`

GetOpenPeriodType returns the OpenPeriodType field if non-nil, zero value otherwise.

### GetOpenPeriodTypeOk

`func (o *SpecialOpeningHourOpeningHours) GetOpenPeriodTypeOk() (*string, bool)`

GetOpenPeriodTypeOk returns a tuple with the OpenPeriodType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOpenPeriodType

`func (o *SpecialOpeningHourOpeningHours) SetOpenPeriodType(v string)`

SetOpenPeriodType sets OpenPeriodType field to given value.

### HasOpenPeriodType

`func (o *SpecialOpeningHourOpeningHours) HasOpenPeriodType() bool`

HasOpenPeriodType returns a boolean if a field has been set.

### GetPeriods

`func (o *SpecialOpeningHourOpeningHours) GetPeriods() []OpenPeriod`

GetPeriods returns the Periods field if non-nil, zero value otherwise.

### GetPeriodsOk

`func (o *SpecialOpeningHourOpeningHours) GetPeriodsOk() (*[]OpenPeriod, bool)`

GetPeriodsOk returns a tuple with the Periods field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPeriods

`func (o *SpecialOpeningHourOpeningHours) SetPeriods(v []OpenPeriod)`

SetPeriods sets Periods field to given value.

### HasPeriods

`func (o *SpecialOpeningHourOpeningHours) HasPeriods() bool`

HasPeriods returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


