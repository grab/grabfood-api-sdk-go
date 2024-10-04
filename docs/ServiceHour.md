# ServiceHour

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**OpenPeriodType** | **string** | The period type for when the outlet is open. - OpenPeriod &#x3D; open only in given periods - OpenAllDay &#x3D; open 24 hours - CloseAllDay &#x3D; closed 24 hours  | 
**Periods** | Pointer to [**[]OpenPeriod**](OpenPeriod.md) | An array of open periods. Only required when &#x60;openPeriodType&#x60; is **OpenPeriod** | [optional] 

## Methods

### NewServiceHour

`func NewServiceHour(openPeriodType string, ) *ServiceHour`

NewServiceHour instantiates a new ServiceHour object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewServiceHourWithDefaults

`func NewServiceHourWithDefaults() *ServiceHour`

NewServiceHourWithDefaults instantiates a new ServiceHour object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetOpenPeriodType

`func (o *ServiceHour) GetOpenPeriodType() string`

GetOpenPeriodType returns the OpenPeriodType field if non-nil, zero value otherwise.

### GetOpenPeriodTypeOk

`func (o *ServiceHour) GetOpenPeriodTypeOk() (*string, bool)`

GetOpenPeriodTypeOk returns a tuple with the OpenPeriodType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOpenPeriodType

`func (o *ServiceHour) SetOpenPeriodType(v string)`

SetOpenPeriodType sets OpenPeriodType field to given value.


### GetPeriods

`func (o *ServiceHour) GetPeriods() []OpenPeriod`

GetPeriods returns the Periods field if non-nil, zero value otherwise.

### GetPeriodsOk

`func (o *ServiceHour) GetPeriodsOk() (*[]OpenPeriod, bool)`

GetPeriodsOk returns a tuple with the Periods field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPeriods

`func (o *ServiceHour) SetPeriods(v []OpenPeriod)`

SetPeriods sets Periods field to given value.

### HasPeriods

`func (o *ServiceHour) HasPeriods() bool`

HasPeriods returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


