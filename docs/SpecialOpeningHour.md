# SpecialOpeningHour

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**StartDate** | Pointer to **string** | The start date of store special opening hours. | [optional] 
**EndDate** | Pointer to **string** | The end date of store special opening hours. | [optional] 
**Metadata** | Pointer to [**SpecialOpeningHourMetadata**](SpecialOpeningHourMetadata.md) |  | [optional] 
**OpeningHours** | Pointer to [**SpecialOpeningHourOpeningHours**](SpecialOpeningHourOpeningHours.md) |  | [optional] 

## Methods

### NewSpecialOpeningHour

`func NewSpecialOpeningHour() *SpecialOpeningHour`

NewSpecialOpeningHour instantiates a new SpecialOpeningHour object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSpecialOpeningHourWithDefaults

`func NewSpecialOpeningHourWithDefaults() *SpecialOpeningHour`

NewSpecialOpeningHourWithDefaults instantiates a new SpecialOpeningHour object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetStartDate

`func (o *SpecialOpeningHour) GetStartDate() string`

GetStartDate returns the StartDate field if non-nil, zero value otherwise.

### GetStartDateOk

`func (o *SpecialOpeningHour) GetStartDateOk() (*string, bool)`

GetStartDateOk returns a tuple with the StartDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartDate

`func (o *SpecialOpeningHour) SetStartDate(v string)`

SetStartDate sets StartDate field to given value.

### HasStartDate

`func (o *SpecialOpeningHour) HasStartDate() bool`

HasStartDate returns a boolean if a field has been set.

### GetEndDate

`func (o *SpecialOpeningHour) GetEndDate() string`

GetEndDate returns the EndDate field if non-nil, zero value otherwise.

### GetEndDateOk

`func (o *SpecialOpeningHour) GetEndDateOk() (*string, bool)`

GetEndDateOk returns a tuple with the EndDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndDate

`func (o *SpecialOpeningHour) SetEndDate(v string)`

SetEndDate sets EndDate field to given value.

### HasEndDate

`func (o *SpecialOpeningHour) HasEndDate() bool`

HasEndDate returns a boolean if a field has been set.

### GetMetadata

`func (o *SpecialOpeningHour) GetMetadata() SpecialOpeningHourMetadata`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *SpecialOpeningHour) GetMetadataOk() (*SpecialOpeningHourMetadata, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *SpecialOpeningHour) SetMetadata(v SpecialOpeningHourMetadata)`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *SpecialOpeningHour) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### GetOpeningHours

`func (o *SpecialOpeningHour) GetOpeningHours() SpecialOpeningHourOpeningHours`

GetOpeningHours returns the OpeningHours field if non-nil, zero value otherwise.

### GetOpeningHoursOk

`func (o *SpecialOpeningHour) GetOpeningHoursOk() (*SpecialOpeningHourOpeningHours, bool)`

GetOpeningHoursOk returns a tuple with the OpeningHours field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOpeningHours

`func (o *SpecialOpeningHour) SetOpeningHours(v SpecialOpeningHourOpeningHours)`

SetOpeningHours sets OpeningHours field to given value.

### HasOpeningHours

`func (o *SpecialOpeningHour) HasOpeningHours() bool`

HasOpeningHours returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


