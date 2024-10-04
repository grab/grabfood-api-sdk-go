# MenuSectionCategory

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | The category&#39;s ID that is on the partner system. This ID should be unique with a min length of 1 and max of 64. | 
**Name** | **string** | The name of the category. | 
**NameTranslation** | Pointer to **map[string]string** | Translation of the category name. Only support up to 1 translated language. Refer [Menu Translation](#section/Menu-Translation). | [optional] 
**AvailableStatus** | **string** | The status for the category. Refer to FAQs for more details about [availableStatus](#section/Menu/What-is-availableStatus). | 
**Items** | [**[]MenuSectionCategoryItem**](MenuSectionCategoryItem.md) | An array of item JSON objects. Max 300 allowed per category. Refer to [Items](#items) for more information. | 

## Methods

### NewMenuSectionCategory

`func NewMenuSectionCategory(id string, name string, availableStatus string, items []MenuSectionCategoryItem, ) *MenuSectionCategory`

NewMenuSectionCategory instantiates a new MenuSectionCategory object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMenuSectionCategoryWithDefaults

`func NewMenuSectionCategoryWithDefaults() *MenuSectionCategory`

NewMenuSectionCategoryWithDefaults instantiates a new MenuSectionCategory object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *MenuSectionCategory) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *MenuSectionCategory) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *MenuSectionCategory) SetId(v string)`

SetId sets Id field to given value.


### GetName

`func (o *MenuSectionCategory) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *MenuSectionCategory) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *MenuSectionCategory) SetName(v string)`

SetName sets Name field to given value.


### GetNameTranslation

`func (o *MenuSectionCategory) GetNameTranslation() map[string]string`

GetNameTranslation returns the NameTranslation field if non-nil, zero value otherwise.

### GetNameTranslationOk

`func (o *MenuSectionCategory) GetNameTranslationOk() (*map[string]string, bool)`

GetNameTranslationOk returns a tuple with the NameTranslation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNameTranslation

`func (o *MenuSectionCategory) SetNameTranslation(v map[string]string)`

SetNameTranslation sets NameTranslation field to given value.

### HasNameTranslation

`func (o *MenuSectionCategory) HasNameTranslation() bool`

HasNameTranslation returns a boolean if a field has been set.

### GetAvailableStatus

`func (o *MenuSectionCategory) GetAvailableStatus() string`

GetAvailableStatus returns the AvailableStatus field if non-nil, zero value otherwise.

### GetAvailableStatusOk

`func (o *MenuSectionCategory) GetAvailableStatusOk() (*string, bool)`

GetAvailableStatusOk returns a tuple with the AvailableStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvailableStatus

`func (o *MenuSectionCategory) SetAvailableStatus(v string)`

SetAvailableStatus sets AvailableStatus field to given value.


### GetItems

`func (o *MenuSectionCategory) GetItems() []MenuSectionCategoryItem`

GetItems returns the Items field if non-nil, zero value otherwise.

### GetItemsOk

`func (o *MenuSectionCategory) GetItemsOk() (*[]MenuSectionCategoryItem, bool)`

GetItemsOk returns a tuple with the Items field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItems

`func (o *MenuSectionCategory) SetItems(v []MenuSectionCategoryItem)`

SetItems sets Items field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


