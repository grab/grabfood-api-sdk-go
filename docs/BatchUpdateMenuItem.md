# BatchUpdateMenuItem

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**MerchantID** | **string** | The merchant&#39;s ID that is in GrabFood&#39;s database. | 
**Field** | **string** | The record type that you want to update. | 
**MenuEntities** | Pointer to [**[]MenuEntity**](MenuEntity.md) | The items in an array of JSON Object.  | [optional] 

## Methods

### NewBatchUpdateMenuItem

`func NewBatchUpdateMenuItem(merchantID string, field string, ) *BatchUpdateMenuItem`

NewBatchUpdateMenuItem instantiates a new BatchUpdateMenuItem object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBatchUpdateMenuItemWithDefaults

`func NewBatchUpdateMenuItemWithDefaults() *BatchUpdateMenuItem`

NewBatchUpdateMenuItemWithDefaults instantiates a new BatchUpdateMenuItem object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMerchantID

`func (o *BatchUpdateMenuItem) GetMerchantID() string`

GetMerchantID returns the MerchantID field if non-nil, zero value otherwise.

### GetMerchantIDOk

`func (o *BatchUpdateMenuItem) GetMerchantIDOk() (*string, bool)`

GetMerchantIDOk returns a tuple with the MerchantID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMerchantID

`func (o *BatchUpdateMenuItem) SetMerchantID(v string)`

SetMerchantID sets MerchantID field to given value.


### GetField

`func (o *BatchUpdateMenuItem) GetField() string`

GetField returns the Field field if non-nil, zero value otherwise.

### GetFieldOk

`func (o *BatchUpdateMenuItem) GetFieldOk() (*string, bool)`

GetFieldOk returns a tuple with the Field field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetField

`func (o *BatchUpdateMenuItem) SetField(v string)`

SetField sets Field field to given value.


### GetMenuEntities

`func (o *BatchUpdateMenuItem) GetMenuEntities() []MenuEntity`

GetMenuEntities returns the MenuEntities field if non-nil, zero value otherwise.

### GetMenuEntitiesOk

`func (o *BatchUpdateMenuItem) GetMenuEntitiesOk() (*[]MenuEntity, bool)`

GetMenuEntitiesOk returns a tuple with the MenuEntities field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMenuEntities

`func (o *BatchUpdateMenuItem) SetMenuEntities(v []MenuEntity)`

SetMenuEntities sets MenuEntities field to given value.

### HasMenuEntities

`func (o *BatchUpdateMenuItem) HasMenuEntities() bool`

HasMenuEntities returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


