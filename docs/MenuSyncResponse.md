# MenuSyncResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CreatedTime** | **time.Time** | The Unix time the specified menu was created in GrabFood&#39;s database. | 
**UpdatedTime** | **time.Time** | The Unix time the specified menu was created in GrabFood&#39;s database. | 
**Code** | **string** | The status code for this request. See [Menu sync response statuses](#section/Menu-sync-response-statuses) for more information. | 
**Errors** | Pointer to **[]string** | An array of strings of error message. | [optional] 
**Sections** | Pointer to [**[]MenuSyncFail**](MenuSyncFail.md) |  | [optional] 

## Methods

### NewMenuSyncResponse

`func NewMenuSyncResponse(createdTime time.Time, updatedTime time.Time, code string, ) *MenuSyncResponse`

NewMenuSyncResponse instantiates a new MenuSyncResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMenuSyncResponseWithDefaults

`func NewMenuSyncResponseWithDefaults() *MenuSyncResponse`

NewMenuSyncResponseWithDefaults instantiates a new MenuSyncResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCreatedTime

`func (o *MenuSyncResponse) GetCreatedTime() time.Time`

GetCreatedTime returns the CreatedTime field if non-nil, zero value otherwise.

### GetCreatedTimeOk

`func (o *MenuSyncResponse) GetCreatedTimeOk() (*time.Time, bool)`

GetCreatedTimeOk returns a tuple with the CreatedTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedTime

`func (o *MenuSyncResponse) SetCreatedTime(v time.Time)`

SetCreatedTime sets CreatedTime field to given value.


### GetUpdatedTime

`func (o *MenuSyncResponse) GetUpdatedTime() time.Time`

GetUpdatedTime returns the UpdatedTime field if non-nil, zero value otherwise.

### GetUpdatedTimeOk

`func (o *MenuSyncResponse) GetUpdatedTimeOk() (*time.Time, bool)`

GetUpdatedTimeOk returns a tuple with the UpdatedTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedTime

`func (o *MenuSyncResponse) SetUpdatedTime(v time.Time)`

SetUpdatedTime sets UpdatedTime field to given value.


### GetCode

`func (o *MenuSyncResponse) GetCode() string`

GetCode returns the Code field if non-nil, zero value otherwise.

### GetCodeOk

`func (o *MenuSyncResponse) GetCodeOk() (*string, bool)`

GetCodeOk returns a tuple with the Code field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCode

`func (o *MenuSyncResponse) SetCode(v string)`

SetCode sets Code field to given value.


### GetErrors

`func (o *MenuSyncResponse) GetErrors() []string`

GetErrors returns the Errors field if non-nil, zero value otherwise.

### GetErrorsOk

`func (o *MenuSyncResponse) GetErrorsOk() (*[]string, bool)`

GetErrorsOk returns a tuple with the Errors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrors

`func (o *MenuSyncResponse) SetErrors(v []string)`

SetErrors sets Errors field to given value.

### HasErrors

`func (o *MenuSyncResponse) HasErrors() bool`

HasErrors returns a boolean if a field has been set.

### SetErrorsNil

`func (o *MenuSyncResponse) SetErrorsNil(b bool)`

 SetErrorsNil sets the value for Errors to be an explicit nil

### UnsetErrors
`func (o *MenuSyncResponse) UnsetErrors()`

UnsetErrors ensures that no value is present for Errors, not even an explicit nil
### GetSections

`func (o *MenuSyncResponse) GetSections() []MenuSyncFail`

GetSections returns the Sections field if non-nil, zero value otherwise.

### GetSectionsOk

`func (o *MenuSyncResponse) GetSectionsOk() (*[]MenuSyncFail, bool)`

GetSectionsOk returns a tuple with the Sections field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSections

`func (o *MenuSyncResponse) SetSections(v []MenuSyncFail)`

SetSections sets Sections field to given value.

### HasSections

`func (o *MenuSyncResponse) HasSections() bool`

HasSections returns a boolean if a field has been set.

### SetSectionsNil

`func (o *MenuSyncResponse) SetSectionsNil(b bool)`

 SetSectionsNil sets the value for Sections to be an explicit nil

### UnsetSections
`func (o *MenuSyncResponse) UnsetSections()`

UnsetSections ensures that no value is present for Sections, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


