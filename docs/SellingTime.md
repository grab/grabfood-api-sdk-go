# SellingTime

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**StartTime** | Pointer to **string** | The selling time group start time in date &amp; time format.  | [optional] 
**EndTime** | Pointer to **string** | The selling time group end time in date &amp; time format.  | [optional] 
**Id** | Pointer to **string** | The selling time&#39;s ID on the partner system. This ID should be unique with length min 1 and max 64.  | [optional] 
**Name** | Pointer to **string** | The name of the selling time.  | [optional] 
**ServiceHours** | Pointer to [**ServiceHours**](ServiceHours.md) |  | [optional] 

## Methods

### NewSellingTime

`func NewSellingTime() *SellingTime`

NewSellingTime instantiates a new SellingTime object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSellingTimeWithDefaults

`func NewSellingTimeWithDefaults() *SellingTime`

NewSellingTimeWithDefaults instantiates a new SellingTime object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetStartTime

`func (o *SellingTime) GetStartTime() string`

GetStartTime returns the StartTime field if non-nil, zero value otherwise.

### GetStartTimeOk

`func (o *SellingTime) GetStartTimeOk() (*string, bool)`

GetStartTimeOk returns a tuple with the StartTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartTime

`func (o *SellingTime) SetStartTime(v string)`

SetStartTime sets StartTime field to given value.

### HasStartTime

`func (o *SellingTime) HasStartTime() bool`

HasStartTime returns a boolean if a field has been set.

### GetEndTime

`func (o *SellingTime) GetEndTime() string`

GetEndTime returns the EndTime field if non-nil, zero value otherwise.

### GetEndTimeOk

`func (o *SellingTime) GetEndTimeOk() (*string, bool)`

GetEndTimeOk returns a tuple with the EndTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndTime

`func (o *SellingTime) SetEndTime(v string)`

SetEndTime sets EndTime field to given value.

### HasEndTime

`func (o *SellingTime) HasEndTime() bool`

HasEndTime returns a boolean if a field has been set.

### GetId

`func (o *SellingTime) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *SellingTime) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *SellingTime) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *SellingTime) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *SellingTime) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *SellingTime) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *SellingTime) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *SellingTime) HasName() bool`

HasName returns a boolean if a field has been set.

### GetServiceHours

`func (o *SellingTime) GetServiceHours() ServiceHours`

GetServiceHours returns the ServiceHours field if non-nil, zero value otherwise.

### GetServiceHoursOk

`func (o *SellingTime) GetServiceHoursOk() (*ServiceHours, bool)`

GetServiceHoursOk returns a tuple with the ServiceHours field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceHours

`func (o *SellingTime) SetServiceHours(v ServiceHours)`

SetServiceHours sets ServiceHours field to given value.

### HasServiceHours

`func (o *SellingTime) HasServiceHours() bool`

HasServiceHours returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


