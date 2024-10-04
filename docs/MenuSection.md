# MenuSection

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | The section&#39;s ID in the partner system.  | 
**Name** | **string** | The name of the section. | 
**ServiceHours** | [**ServiceHours**](ServiceHours.md) |  | 
**Categories** | [**[]MenuSectionCategory**](MenuSectionCategory.md) | An array of category JSON objects. Max 100 allowed per section. Refer to [Categories](#categories) for more information. | 

## Methods

### NewMenuSection

`func NewMenuSection(id string, name string, serviceHours ServiceHours, categories []MenuSectionCategory, ) *MenuSection`

NewMenuSection instantiates a new MenuSection object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMenuSectionWithDefaults

`func NewMenuSectionWithDefaults() *MenuSection`

NewMenuSectionWithDefaults instantiates a new MenuSection object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *MenuSection) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *MenuSection) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *MenuSection) SetId(v string)`

SetId sets Id field to given value.


### GetName

`func (o *MenuSection) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *MenuSection) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *MenuSection) SetName(v string)`

SetName sets Name field to given value.


### GetServiceHours

`func (o *MenuSection) GetServiceHours() ServiceHours`

GetServiceHours returns the ServiceHours field if non-nil, zero value otherwise.

### GetServiceHoursOk

`func (o *MenuSection) GetServiceHoursOk() (*ServiceHours, bool)`

GetServiceHoursOk returns a tuple with the ServiceHours field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceHours

`func (o *MenuSection) SetServiceHours(v ServiceHours)`

SetServiceHours sets ServiceHours field to given value.


### GetCategories

`func (o *MenuSection) GetCategories() []MenuSectionCategory`

GetCategories returns the Categories field if non-nil, zero value otherwise.

### GetCategoriesOk

`func (o *MenuSection) GetCategoriesOk() (*[]MenuSectionCategory, bool)`

GetCategoriesOk returns a tuple with the Categories field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCategories

`func (o *MenuSection) SetCategories(v []MenuSectionCategory)`

SetCategories sets Categories field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


