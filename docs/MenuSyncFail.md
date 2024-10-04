# MenuSyncFail

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** |  | [optional] 
**Errors** | Pointer to **[]string** | An array of strings of error message. | [optional] 
**ServiceHours** | Pointer to [**MenuSyncFailServiceHours**](MenuSyncFailServiceHours.md) |  | [optional] 
**Categories** | Pointer to [**[]MenuSyncFailCategory**](MenuSyncFailCategory.md) |  | [optional] 

## Methods

### NewMenuSyncFail

`func NewMenuSyncFail() *MenuSyncFail`

NewMenuSyncFail instantiates a new MenuSyncFail object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMenuSyncFailWithDefaults

`func NewMenuSyncFailWithDefaults() *MenuSyncFail`

NewMenuSyncFailWithDefaults instantiates a new MenuSyncFail object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *MenuSyncFail) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *MenuSyncFail) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *MenuSyncFail) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *MenuSyncFail) HasId() bool`

HasId returns a boolean if a field has been set.

### GetErrors

`func (o *MenuSyncFail) GetErrors() []string`

GetErrors returns the Errors field if non-nil, zero value otherwise.

### GetErrorsOk

`func (o *MenuSyncFail) GetErrorsOk() (*[]string, bool)`

GetErrorsOk returns a tuple with the Errors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrors

`func (o *MenuSyncFail) SetErrors(v []string)`

SetErrors sets Errors field to given value.

### HasErrors

`func (o *MenuSyncFail) HasErrors() bool`

HasErrors returns a boolean if a field has been set.

### GetServiceHours

`func (o *MenuSyncFail) GetServiceHours() MenuSyncFailServiceHours`

GetServiceHours returns the ServiceHours field if non-nil, zero value otherwise.

### GetServiceHoursOk

`func (o *MenuSyncFail) GetServiceHoursOk() (*MenuSyncFailServiceHours, bool)`

GetServiceHoursOk returns a tuple with the ServiceHours field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceHours

`func (o *MenuSyncFail) SetServiceHours(v MenuSyncFailServiceHours)`

SetServiceHours sets ServiceHours field to given value.

### HasServiceHours

`func (o *MenuSyncFail) HasServiceHours() bool`

HasServiceHours returns a boolean if a field has been set.

### GetCategories

`func (o *MenuSyncFail) GetCategories() []MenuSyncFailCategory`

GetCategories returns the Categories field if non-nil, zero value otherwise.

### GetCategoriesOk

`func (o *MenuSyncFail) GetCategoriesOk() (*[]MenuSyncFailCategory, bool)`

GetCategoriesOk returns a tuple with the Categories field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCategories

`func (o *MenuSyncFail) SetCategories(v []MenuSyncFailCategory)`

SetCategories sets Categories field to given value.

### HasCategories

`func (o *MenuSyncFail) HasCategories() bool`

HasCategories returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


