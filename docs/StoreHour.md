# StoreHour

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Mon** | [**[]OpenPeriod**](OpenPeriod.md) | An array of open periods. Maximum of 3 periods. Blank indicates store close. | 
**Tue** | [**[]OpenPeriod**](OpenPeriod.md) | An array of open periods. Maximum of 3 periods. Blank indicates store close. | 
**Wed** | [**[]OpenPeriod**](OpenPeriod.md) | An array of open periods. Maximum of 3 periods. Blank indicates store close. | 
**Thu** | [**[]OpenPeriod**](OpenPeriod.md) | An array of open periods. Maximum of 3 periods. Blank indicates store close. | 
**Fri** | [**[]OpenPeriod**](OpenPeriod.md) | An array of open periods. Maximum of 3 periods. Blank indicates store close. | 
**Sat** | [**[]OpenPeriod**](OpenPeriod.md) | An array of open periods. Maximum of 3 periods. Blank indicates store close. | 
**Sun** | [**[]OpenPeriod**](OpenPeriod.md) | An array of open periods. Maximum of 3 periods. Blank indicates store close. | 

## Methods

### NewStoreHour

`func NewStoreHour(mon []OpenPeriod, tue []OpenPeriod, wed []OpenPeriod, thu []OpenPeriod, fri []OpenPeriod, sat []OpenPeriod, sun []OpenPeriod, ) *StoreHour`

NewStoreHour instantiates a new StoreHour object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewStoreHourWithDefaults

`func NewStoreHourWithDefaults() *StoreHour`

NewStoreHourWithDefaults instantiates a new StoreHour object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMon

`func (o *StoreHour) GetMon() []OpenPeriod`

GetMon returns the Mon field if non-nil, zero value otherwise.

### GetMonOk

`func (o *StoreHour) GetMonOk() (*[]OpenPeriod, bool)`

GetMonOk returns a tuple with the Mon field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMon

`func (o *StoreHour) SetMon(v []OpenPeriod)`

SetMon sets Mon field to given value.


### GetTue

`func (o *StoreHour) GetTue() []OpenPeriod`

GetTue returns the Tue field if non-nil, zero value otherwise.

### GetTueOk

`func (o *StoreHour) GetTueOk() (*[]OpenPeriod, bool)`

GetTueOk returns a tuple with the Tue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTue

`func (o *StoreHour) SetTue(v []OpenPeriod)`

SetTue sets Tue field to given value.


### GetWed

`func (o *StoreHour) GetWed() []OpenPeriod`

GetWed returns the Wed field if non-nil, zero value otherwise.

### GetWedOk

`func (o *StoreHour) GetWedOk() (*[]OpenPeriod, bool)`

GetWedOk returns a tuple with the Wed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWed

`func (o *StoreHour) SetWed(v []OpenPeriod)`

SetWed sets Wed field to given value.


### GetThu

`func (o *StoreHour) GetThu() []OpenPeriod`

GetThu returns the Thu field if non-nil, zero value otherwise.

### GetThuOk

`func (o *StoreHour) GetThuOk() (*[]OpenPeriod, bool)`

GetThuOk returns a tuple with the Thu field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetThu

`func (o *StoreHour) SetThu(v []OpenPeriod)`

SetThu sets Thu field to given value.


### GetFri

`func (o *StoreHour) GetFri() []OpenPeriod`

GetFri returns the Fri field if non-nil, zero value otherwise.

### GetFriOk

`func (o *StoreHour) GetFriOk() (*[]OpenPeriod, bool)`

GetFriOk returns a tuple with the Fri field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFri

`func (o *StoreHour) SetFri(v []OpenPeriod)`

SetFri sets Fri field to given value.


### GetSat

`func (o *StoreHour) GetSat() []OpenPeriod`

GetSat returns the Sat field if non-nil, zero value otherwise.

### GetSatOk

`func (o *StoreHour) GetSatOk() (*[]OpenPeriod, bool)`

GetSatOk returns a tuple with the Sat field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSat

`func (o *StoreHour) SetSat(v []OpenPeriod)`

SetSat sets Sat field to given value.


### GetSun

`func (o *StoreHour) GetSun() []OpenPeriod`

GetSun returns the Sun field if non-nil, zero value otherwise.

### GetSunOk

`func (o *StoreHour) GetSunOk() (*[]OpenPeriod, bool)`

GetSunOk returns a tuple with the Sun field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSun

`func (o *StoreHour) SetSun(v []OpenPeriod)`

SetSun sets Sun field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


