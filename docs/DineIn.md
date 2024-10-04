# DineIn

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**TableID** | Pointer to **string** | Table number. | [optional] 
**EaterCount** | Pointer to **int64** | The number of eaters. | [optional] 

## Methods

### NewDineIn

`func NewDineIn() *DineIn`

NewDineIn instantiates a new DineIn object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDineInWithDefaults

`func NewDineInWithDefaults() *DineIn`

NewDineInWithDefaults instantiates a new DineIn object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTableID

`func (o *DineIn) GetTableID() string`

GetTableID returns the TableID field if non-nil, zero value otherwise.

### GetTableIDOk

`func (o *DineIn) GetTableIDOk() (*string, bool)`

GetTableIDOk returns a tuple with the TableID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTableID

`func (o *DineIn) SetTableID(v string)`

SetTableID sets TableID field to given value.

### HasTableID

`func (o *DineIn) HasTableID() bool`

HasTableID returns a boolean if a field has been set.

### GetEaterCount

`func (o *DineIn) GetEaterCount() int64`

GetEaterCount returns the EaterCount field if non-nil, zero value otherwise.

### GetEaterCountOk

`func (o *DineIn) GetEaterCountOk() (*int64, bool)`

GetEaterCountOk returns a tuple with the EaterCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEaterCount

`func (o *DineIn) SetEaterCount(v int64)`

SetEaterCount sets EaterCount field to given value.

### HasEaterCount

`func (o *DineIn) HasEaterCount() bool`

HasEaterCount returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


