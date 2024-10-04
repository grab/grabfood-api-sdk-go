# StoreHourResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DineInHour** | Pointer to [**StoreHour**](StoreHour.md) |  | [optional] 
**OpeningHour** | Pointer to [**StoreHour**](StoreHour.md) |  | [optional] 
**SpecialOpeningHours** | Pointer to [**[]SpecialOpeningHour**](SpecialOpeningHour.md) | The store&#39;s special opening hours. | [optional] 

## Methods

### NewStoreHourResponse

`func NewStoreHourResponse() *StoreHourResponse`

NewStoreHourResponse instantiates a new StoreHourResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewStoreHourResponseWithDefaults

`func NewStoreHourResponseWithDefaults() *StoreHourResponse`

NewStoreHourResponseWithDefaults instantiates a new StoreHourResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDineInHour

`func (o *StoreHourResponse) GetDineInHour() StoreHour`

GetDineInHour returns the DineInHour field if non-nil, zero value otherwise.

### GetDineInHourOk

`func (o *StoreHourResponse) GetDineInHourOk() (*StoreHour, bool)`

GetDineInHourOk returns a tuple with the DineInHour field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDineInHour

`func (o *StoreHourResponse) SetDineInHour(v StoreHour)`

SetDineInHour sets DineInHour field to given value.

### HasDineInHour

`func (o *StoreHourResponse) HasDineInHour() bool`

HasDineInHour returns a boolean if a field has been set.

### GetOpeningHour

`func (o *StoreHourResponse) GetOpeningHour() StoreHour`

GetOpeningHour returns the OpeningHour field if non-nil, zero value otherwise.

### GetOpeningHourOk

`func (o *StoreHourResponse) GetOpeningHourOk() (*StoreHour, bool)`

GetOpeningHourOk returns a tuple with the OpeningHour field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOpeningHour

`func (o *StoreHourResponse) SetOpeningHour(v StoreHour)`

SetOpeningHour sets OpeningHour field to given value.

### HasOpeningHour

`func (o *StoreHourResponse) HasOpeningHour() bool`

HasOpeningHour returns a boolean if a field has been set.

### GetSpecialOpeningHours

`func (o *StoreHourResponse) GetSpecialOpeningHours() []SpecialOpeningHour`

GetSpecialOpeningHours returns the SpecialOpeningHours field if non-nil, zero value otherwise.

### GetSpecialOpeningHoursOk

`func (o *StoreHourResponse) GetSpecialOpeningHoursOk() (*[]SpecialOpeningHour, bool)`

GetSpecialOpeningHoursOk returns a tuple with the SpecialOpeningHours field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpecialOpeningHours

`func (o *StoreHourResponse) SetSpecialOpeningHours(v []SpecialOpeningHour)`

SetSpecialOpeningHours sets SpecialOpeningHours field to given value.

### HasSpecialOpeningHours

`func (o *StoreHourResponse) HasSpecialOpeningHours() bool`

HasSpecialOpeningHours returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


