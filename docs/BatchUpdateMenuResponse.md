# BatchUpdateMenuResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**MerchantID** | Pointer to **string** | The merchant&#39;s ID that is in GrabFood&#39;s database. | [optional] 
**Status** | Pointer to **string** | The status of this request. | [optional] 
**Errors** | Pointer to [**[]MenuEntityError**](MenuEntityError.md) | The error messages when batch update menu record was partial success/fail. &#x60;null&#x60; when the request was success. | [optional] 

## Methods

### NewBatchUpdateMenuResponse

`func NewBatchUpdateMenuResponse() *BatchUpdateMenuResponse`

NewBatchUpdateMenuResponse instantiates a new BatchUpdateMenuResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBatchUpdateMenuResponseWithDefaults

`func NewBatchUpdateMenuResponseWithDefaults() *BatchUpdateMenuResponse`

NewBatchUpdateMenuResponseWithDefaults instantiates a new BatchUpdateMenuResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMerchantID

`func (o *BatchUpdateMenuResponse) GetMerchantID() string`

GetMerchantID returns the MerchantID field if non-nil, zero value otherwise.

### GetMerchantIDOk

`func (o *BatchUpdateMenuResponse) GetMerchantIDOk() (*string, bool)`

GetMerchantIDOk returns a tuple with the MerchantID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMerchantID

`func (o *BatchUpdateMenuResponse) SetMerchantID(v string)`

SetMerchantID sets MerchantID field to given value.

### HasMerchantID

`func (o *BatchUpdateMenuResponse) HasMerchantID() bool`

HasMerchantID returns a boolean if a field has been set.

### GetStatus

`func (o *BatchUpdateMenuResponse) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *BatchUpdateMenuResponse) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *BatchUpdateMenuResponse) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *BatchUpdateMenuResponse) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetErrors

`func (o *BatchUpdateMenuResponse) GetErrors() []MenuEntityError`

GetErrors returns the Errors field if non-nil, zero value otherwise.

### GetErrorsOk

`func (o *BatchUpdateMenuResponse) GetErrorsOk() (*[]MenuEntityError, bool)`

GetErrorsOk returns a tuple with the Errors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrors

`func (o *BatchUpdateMenuResponse) SetErrors(v []MenuEntityError)`

SetErrors sets Errors field to given value.

### HasErrors

`func (o *BatchUpdateMenuResponse) HasErrors() bool`

HasErrors returns a boolean if a field has been set.

### SetErrorsNil

`func (o *BatchUpdateMenuResponse) SetErrorsNil(b bool)`

 SetErrorsNil sets the value for Errors to be an explicit nil

### UnsetErrors
`func (o *BatchUpdateMenuResponse) UnsetErrors()`

UnsetErrors ensures that no value is present for Errors, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


