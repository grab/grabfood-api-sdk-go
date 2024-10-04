# MenuSyncFailItem

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | The item&#39;s ID in the partner system.  | [optional] 
**Errors** | Pointer to **[]string** | An array of strings of error message. | [optional] 
**ModifierGroups** | Pointer to [**[]MenuSyncFailModifierGroup**](MenuSyncFailModifierGroup.md) |  | [optional] 

## Methods

### NewMenuSyncFailItem

`func NewMenuSyncFailItem() *MenuSyncFailItem`

NewMenuSyncFailItem instantiates a new MenuSyncFailItem object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMenuSyncFailItemWithDefaults

`func NewMenuSyncFailItemWithDefaults() *MenuSyncFailItem`

NewMenuSyncFailItemWithDefaults instantiates a new MenuSyncFailItem object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *MenuSyncFailItem) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *MenuSyncFailItem) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *MenuSyncFailItem) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *MenuSyncFailItem) HasId() bool`

HasId returns a boolean if a field has been set.

### GetErrors

`func (o *MenuSyncFailItem) GetErrors() []string`

GetErrors returns the Errors field if non-nil, zero value otherwise.

### GetErrorsOk

`func (o *MenuSyncFailItem) GetErrorsOk() (*[]string, bool)`

GetErrorsOk returns a tuple with the Errors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrors

`func (o *MenuSyncFailItem) SetErrors(v []string)`

SetErrors sets Errors field to given value.

### HasErrors

`func (o *MenuSyncFailItem) HasErrors() bool`

HasErrors returns a boolean if a field has been set.

### GetModifierGroups

`func (o *MenuSyncFailItem) GetModifierGroups() []MenuSyncFailModifierGroup`

GetModifierGroups returns the ModifierGroups field if non-nil, zero value otherwise.

### GetModifierGroupsOk

`func (o *MenuSyncFailItem) GetModifierGroupsOk() (*[]MenuSyncFailModifierGroup, bool)`

GetModifierGroupsOk returns a tuple with the ModifierGroups field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModifierGroups

`func (o *MenuSyncFailItem) SetModifierGroups(v []MenuSyncFailModifierGroup)`

SetModifierGroups sets ModifierGroups field to given value.

### HasModifierGroups

`func (o *MenuSyncFailItem) HasModifierGroups() bool`

HasModifierGroups returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


