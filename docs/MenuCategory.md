# MenuCategory

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | The category&#39;s ID that is on the partner system. This ID should be unique with a min length of 1 and max of 64. | 
**Name** | **string** | The name of the category. | 
**NameTranslation** | Pointer to **map[string]string** | Translation of the category name. Only support up to 1 translated language. Refer [Menu Translation](#section/Menu-Translation). | [optional] 
**AvailableStatus** | **string** | The status for the category. Refer to FAQs for more details about [availableStatus](#section/Menu/What-is-availableStatus). | 
**SellingTimeID** | **string** | The selling time&#39;s ID for the category. All items within the category will apply the same selling time unless there is another selling time specified for the item. | 
**Items** | [**[]MenuItem**](MenuItem.md) | An array of item JSON objects. Max 300 allowed per category. Refer to [Items](#items) for more information. | 

## Methods

### NewMenuCategory

`func NewMenuCategory(id string, name string, availableStatus string, sellingTimeID string, items []MenuItem, ) *MenuCategory`

NewMenuCategory instantiates a new MenuCategory object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMenuCategoryWithDefaults

`func NewMenuCategoryWithDefaults() *MenuCategory`

NewMenuCategoryWithDefaults instantiates a new MenuCategory object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *MenuCategory) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *MenuCategory) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *MenuCategory) SetId(v string)`

SetId sets Id field to given value.


### GetName

`func (o *MenuCategory) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *MenuCategory) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *MenuCategory) SetName(v string)`

SetName sets Name field to given value.


### GetNameTranslation

`func (o *MenuCategory) GetNameTranslation() map[string]string`

GetNameTranslation returns the NameTranslation field if non-nil, zero value otherwise.

### GetNameTranslationOk

`func (o *MenuCategory) GetNameTranslationOk() (*map[string]string, bool)`

GetNameTranslationOk returns a tuple with the NameTranslation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNameTranslation

`func (o *MenuCategory) SetNameTranslation(v map[string]string)`

SetNameTranslation sets NameTranslation field to given value.

### HasNameTranslation

`func (o *MenuCategory) HasNameTranslation() bool`

HasNameTranslation returns a boolean if a field has been set.

### GetAvailableStatus

`func (o *MenuCategory) GetAvailableStatus() string`

GetAvailableStatus returns the AvailableStatus field if non-nil, zero value otherwise.

### GetAvailableStatusOk

`func (o *MenuCategory) GetAvailableStatusOk() (*string, bool)`

GetAvailableStatusOk returns a tuple with the AvailableStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvailableStatus

`func (o *MenuCategory) SetAvailableStatus(v string)`

SetAvailableStatus sets AvailableStatus field to given value.


### GetSellingTimeID

`func (o *MenuCategory) GetSellingTimeID() string`

GetSellingTimeID returns the SellingTimeID field if non-nil, zero value otherwise.

### GetSellingTimeIDOk

`func (o *MenuCategory) GetSellingTimeIDOk() (*string, bool)`

GetSellingTimeIDOk returns a tuple with the SellingTimeID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSellingTimeID

`func (o *MenuCategory) SetSellingTimeID(v string)`

SetSellingTimeID sets SellingTimeID field to given value.


### GetItems

`func (o *MenuCategory) GetItems() []MenuItem`

GetItems returns the Items field if non-nil, zero value otherwise.

### GetItemsOk

`func (o *MenuCategory) GetItemsOk() (*[]MenuItem, bool)`

GetItemsOk returns a tuple with the Items field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItems

`func (o *MenuCategory) SetItems(v []MenuItem)`

SetItems sets Items field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


