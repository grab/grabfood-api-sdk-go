# ModifierGroup

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | The ModifierGroup&#39;s ID that is on the partner system. This ID should be unique with a min length of 1 and max of 64. | 
**Name** | **string** | The name of the ModifierGroup for the item that is in the parent category and section. | 
**NameTranslation** | Pointer to **map[string]string** | Translation of the modifier group name. Only support up to 1 translated language. Refer [Menu Translation](#section/Menu-Translation). | [optional] 
**AvailableStatus** | **string** | The status for the ModifierGroup that is in the item. | 
**SelectionRangeMin** | Pointer to **int32** | The minimum quantity of the attribute. Refer to FAQs for more details about [selection range](#section/Menu/What-does-the-selection-range-do). | [optional] 
**SelectionRangeMax** | **int32** | The maximum quantity of the attribute. Refer to FAQs for more details about [selection range](#section/Menu/What-does-the-selection-range-do). | 
**Modifiers** | Pointer to [**[]MenuModifier**](MenuModifier.md) | An array of modifier JSON objects. Max 100 per modifierGroup. Refer to [Modifiers](#modifiers) for more information. | [optional] 

## Methods

### NewModifierGroup

`func NewModifierGroup(id string, name string, availableStatus string, selectionRangeMax int32, ) *ModifierGroup`

NewModifierGroup instantiates a new ModifierGroup object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewModifierGroupWithDefaults

`func NewModifierGroupWithDefaults() *ModifierGroup`

NewModifierGroupWithDefaults instantiates a new ModifierGroup object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ModifierGroup) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ModifierGroup) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ModifierGroup) SetId(v string)`

SetId sets Id field to given value.


### GetName

`func (o *ModifierGroup) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ModifierGroup) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ModifierGroup) SetName(v string)`

SetName sets Name field to given value.


### GetNameTranslation

`func (o *ModifierGroup) GetNameTranslation() map[string]string`

GetNameTranslation returns the NameTranslation field if non-nil, zero value otherwise.

### GetNameTranslationOk

`func (o *ModifierGroup) GetNameTranslationOk() (*map[string]string, bool)`

GetNameTranslationOk returns a tuple with the NameTranslation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNameTranslation

`func (o *ModifierGroup) SetNameTranslation(v map[string]string)`

SetNameTranslation sets NameTranslation field to given value.

### HasNameTranslation

`func (o *ModifierGroup) HasNameTranslation() bool`

HasNameTranslation returns a boolean if a field has been set.

### GetAvailableStatus

`func (o *ModifierGroup) GetAvailableStatus() string`

GetAvailableStatus returns the AvailableStatus field if non-nil, zero value otherwise.

### GetAvailableStatusOk

`func (o *ModifierGroup) GetAvailableStatusOk() (*string, bool)`

GetAvailableStatusOk returns a tuple with the AvailableStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvailableStatus

`func (o *ModifierGroup) SetAvailableStatus(v string)`

SetAvailableStatus sets AvailableStatus field to given value.


### GetSelectionRangeMin

`func (o *ModifierGroup) GetSelectionRangeMin() int32`

GetSelectionRangeMin returns the SelectionRangeMin field if non-nil, zero value otherwise.

### GetSelectionRangeMinOk

`func (o *ModifierGroup) GetSelectionRangeMinOk() (*int32, bool)`

GetSelectionRangeMinOk returns a tuple with the SelectionRangeMin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSelectionRangeMin

`func (o *ModifierGroup) SetSelectionRangeMin(v int32)`

SetSelectionRangeMin sets SelectionRangeMin field to given value.

### HasSelectionRangeMin

`func (o *ModifierGroup) HasSelectionRangeMin() bool`

HasSelectionRangeMin returns a boolean if a field has been set.

### GetSelectionRangeMax

`func (o *ModifierGroup) GetSelectionRangeMax() int32`

GetSelectionRangeMax returns the SelectionRangeMax field if non-nil, zero value otherwise.

### GetSelectionRangeMaxOk

`func (o *ModifierGroup) GetSelectionRangeMaxOk() (*int32, bool)`

GetSelectionRangeMaxOk returns a tuple with the SelectionRangeMax field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSelectionRangeMax

`func (o *ModifierGroup) SetSelectionRangeMax(v int32)`

SetSelectionRangeMax sets SelectionRangeMax field to given value.


### GetModifiers

`func (o *ModifierGroup) GetModifiers() []MenuModifier`

GetModifiers returns the Modifiers field if non-nil, zero value otherwise.

### GetModifiersOk

`func (o *ModifierGroup) GetModifiersOk() (*[]MenuModifier, bool)`

GetModifiersOk returns a tuple with the Modifiers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModifiers

`func (o *ModifierGroup) SetModifiers(v []MenuModifier)`

SetModifiers sets Modifiers field to given value.

### HasModifiers

`func (o *ModifierGroup) HasModifiers() bool`

HasModifiers returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


