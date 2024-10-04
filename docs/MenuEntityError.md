# MenuEntityError

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**EntityID** | Pointer to **string** | The itemID. | [optional] 
**ErrMsg** | Pointer to **string** | The error message. | [optional] 

## Methods

### NewMenuEntityError

`func NewMenuEntityError() *MenuEntityError`

NewMenuEntityError instantiates a new MenuEntityError object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMenuEntityErrorWithDefaults

`func NewMenuEntityErrorWithDefaults() *MenuEntityError`

NewMenuEntityErrorWithDefaults instantiates a new MenuEntityError object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEntityID

`func (o *MenuEntityError) GetEntityID() string`

GetEntityID returns the EntityID field if non-nil, zero value otherwise.

### GetEntityIDOk

`func (o *MenuEntityError) GetEntityIDOk() (*string, bool)`

GetEntityIDOk returns a tuple with the EntityID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEntityID

`func (o *MenuEntityError) SetEntityID(v string)`

SetEntityID sets EntityID field to given value.

### HasEntityID

`func (o *MenuEntityError) HasEntityID() bool`

HasEntityID returns a boolean if a field has been set.

### GetErrMsg

`func (o *MenuEntityError) GetErrMsg() string`

GetErrMsg returns the ErrMsg field if non-nil, zero value otherwise.

### GetErrMsgOk

`func (o *MenuEntityError) GetErrMsgOk() (*string, bool)`

GetErrMsgOk returns a tuple with the ErrMsg field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrMsg

`func (o *MenuEntityError) SetErrMsg(v string)`

SetErrMsg sets ErrMsg field to given value.

### HasErrMsg

`func (o *MenuEntityError) HasErrMsg() bool`

HasErrMsg returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


