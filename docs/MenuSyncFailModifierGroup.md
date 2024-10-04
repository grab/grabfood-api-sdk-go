# MenuSyncFailModifierGroup

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** |  | [optional] 
**Errors** | Pointer to **[]string** | An array of strings of error message. | [optional] 
**Modifiers** | Pointer to [**[]MenuSyncFailModifier**](MenuSyncFailModifier.md) |  | [optional] 

## Methods

### NewMenuSyncFailModifierGroup

`func NewMenuSyncFailModifierGroup() *MenuSyncFailModifierGroup`

NewMenuSyncFailModifierGroup instantiates a new MenuSyncFailModifierGroup object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMenuSyncFailModifierGroupWithDefaults

`func NewMenuSyncFailModifierGroupWithDefaults() *MenuSyncFailModifierGroup`

NewMenuSyncFailModifierGroupWithDefaults instantiates a new MenuSyncFailModifierGroup object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *MenuSyncFailModifierGroup) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *MenuSyncFailModifierGroup) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *MenuSyncFailModifierGroup) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *MenuSyncFailModifierGroup) HasId() bool`

HasId returns a boolean if a field has been set.

### GetErrors

`func (o *MenuSyncFailModifierGroup) GetErrors() []string`

GetErrors returns the Errors field if non-nil, zero value otherwise.

### GetErrorsOk

`func (o *MenuSyncFailModifierGroup) GetErrorsOk() (*[]string, bool)`

GetErrorsOk returns a tuple with the Errors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrors

`func (o *MenuSyncFailModifierGroup) SetErrors(v []string)`

SetErrors sets Errors field to given value.

### HasErrors

`func (o *MenuSyncFailModifierGroup) HasErrors() bool`

HasErrors returns a boolean if a field has been set.

### GetModifiers

`func (o *MenuSyncFailModifierGroup) GetModifiers() []MenuSyncFailModifier`

GetModifiers returns the Modifiers field if non-nil, zero value otherwise.

### GetModifiersOk

`func (o *MenuSyncFailModifierGroup) GetModifiersOk() (*[]MenuSyncFailModifier, bool)`

GetModifiersOk returns a tuple with the Modifiers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModifiers

`func (o *MenuSyncFailModifierGroup) SetModifiers(v []MenuSyncFailModifier)`

SetModifiers sets Modifiers field to given value.

### HasModifiers

`func (o *MenuSyncFailModifierGroup) HasModifiers() bool`

HasModifiers returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


