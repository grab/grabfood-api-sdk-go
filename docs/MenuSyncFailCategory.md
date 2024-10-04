# MenuSyncFailCategory

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | The category&#39;s ID that is on the partner system. This ID should be unique with a min length of 1 and max of 64. | [optional] 
**Errors** | Pointer to **[]string** | An array of strings of error message. | [optional] 
**Items** | Pointer to [**[]MenuSyncFailItem**](MenuSyncFailItem.md) | An array of item JSON objects. Max 300 allowed per category. Refer to [Items](#items) for more information. | [optional] 

## Methods

### NewMenuSyncFailCategory

`func NewMenuSyncFailCategory() *MenuSyncFailCategory`

NewMenuSyncFailCategory instantiates a new MenuSyncFailCategory object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMenuSyncFailCategoryWithDefaults

`func NewMenuSyncFailCategoryWithDefaults() *MenuSyncFailCategory`

NewMenuSyncFailCategoryWithDefaults instantiates a new MenuSyncFailCategory object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *MenuSyncFailCategory) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *MenuSyncFailCategory) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *MenuSyncFailCategory) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *MenuSyncFailCategory) HasId() bool`

HasId returns a boolean if a field has been set.

### GetErrors

`func (o *MenuSyncFailCategory) GetErrors() []string`

GetErrors returns the Errors field if non-nil, zero value otherwise.

### GetErrorsOk

`func (o *MenuSyncFailCategory) GetErrorsOk() (*[]string, bool)`

GetErrorsOk returns a tuple with the Errors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrors

`func (o *MenuSyncFailCategory) SetErrors(v []string)`

SetErrors sets Errors field to given value.

### HasErrors

`func (o *MenuSyncFailCategory) HasErrors() bool`

HasErrors returns a boolean if a field has been set.

### GetItems

`func (o *MenuSyncFailCategory) GetItems() []MenuSyncFailItem`

GetItems returns the Items field if non-nil, zero value otherwise.

### GetItemsOk

`func (o *MenuSyncFailCategory) GetItemsOk() (*[]MenuSyncFailItem, bool)`

GetItemsOk returns a tuple with the Items field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItems

`func (o *MenuSyncFailCategory) SetItems(v []MenuSyncFailItem)`

SetItems sets Items field to given value.

### HasItems

`func (o *MenuSyncFailCategory) HasItems() bool`

HasItems returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


