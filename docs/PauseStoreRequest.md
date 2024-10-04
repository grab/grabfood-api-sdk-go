# PauseStoreRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**MerchantID** | **string** | The merchant&#39;s ID that is in GrabFood&#39;s database. | 
**IsPause** | **bool** | Boolean value to pause or unpause store. | 
**Duration** | Pointer to **string** | The duration to pause the store. Only required when &#x60;isPause&#x3D;true&#x60;. | [optional] 

## Methods

### NewPauseStoreRequest

`func NewPauseStoreRequest(merchantID string, isPause bool, ) *PauseStoreRequest`

NewPauseStoreRequest instantiates a new PauseStoreRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPauseStoreRequestWithDefaults

`func NewPauseStoreRequestWithDefaults() *PauseStoreRequest`

NewPauseStoreRequestWithDefaults instantiates a new PauseStoreRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMerchantID

`func (o *PauseStoreRequest) GetMerchantID() string`

GetMerchantID returns the MerchantID field if non-nil, zero value otherwise.

### GetMerchantIDOk

`func (o *PauseStoreRequest) GetMerchantIDOk() (*string, bool)`

GetMerchantIDOk returns a tuple with the MerchantID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMerchantID

`func (o *PauseStoreRequest) SetMerchantID(v string)`

SetMerchantID sets MerchantID field to given value.


### GetIsPause

`func (o *PauseStoreRequest) GetIsPause() bool`

GetIsPause returns the IsPause field if non-nil, zero value otherwise.

### GetIsPauseOk

`func (o *PauseStoreRequest) GetIsPauseOk() (*bool, bool)`

GetIsPauseOk returns a tuple with the IsPause field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsPause

`func (o *PauseStoreRequest) SetIsPause(v bool)`

SetIsPause sets IsPause field to given value.


### GetDuration

`func (o *PauseStoreRequest) GetDuration() string`

GetDuration returns the Duration field if non-nil, zero value otherwise.

### GetDurationOk

`func (o *PauseStoreRequest) GetDurationOk() (*string, bool)`

GetDurationOk returns a tuple with the Duration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDuration

`func (o *PauseStoreRequest) SetDuration(v string)`

SetDuration sets Duration field to given value.

### HasDuration

`func (o *PauseStoreRequest) HasDuration() bool`

HasDuration returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


