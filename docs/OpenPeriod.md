# OpenPeriod

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**StartTime** | **string** | The open start time in 24h format. | 
**EndTime** | **string** | The open start time in 24h format. | 

## Methods

### NewOpenPeriod

`func NewOpenPeriod(startTime string, endTime string, ) *OpenPeriod`

NewOpenPeriod instantiates a new OpenPeriod object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOpenPeriodWithDefaults

`func NewOpenPeriodWithDefaults() *OpenPeriod`

NewOpenPeriodWithDefaults instantiates a new OpenPeriod object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetStartTime

`func (o *OpenPeriod) GetStartTime() string`

GetStartTime returns the StartTime field if non-nil, zero value otherwise.

### GetStartTimeOk

`func (o *OpenPeriod) GetStartTimeOk() (*string, bool)`

GetStartTimeOk returns a tuple with the StartTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartTime

`func (o *OpenPeriod) SetStartTime(v string)`

SetStartTime sets StartTime field to given value.


### GetEndTime

`func (o *OpenPeriod) GetEndTime() string`

GetEndTime returns the EndTime field if non-nil, zero value otherwise.

### GetEndTimeOk

`func (o *OpenPeriod) GetEndTimeOk() (*string, bool)`

GetEndTimeOk returns a tuple with the EndTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndTime

`func (o *OpenPeriod) SetEndTime(v string)`

SetEndTime sets EndTime field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


