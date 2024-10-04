# MenuSyncWebhookRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**RequestID** | Pointer to **string** | An universally unique identifier (UUID) string. Used to uniquely identify a webhook request. Partners should use this value to distinguish between different webhook requests. If two requests contain the same requestID, only the first request should be considered and later requests **must** be ignored or discarded.  | [optional] 
**MerchantID** | Pointer to **string** | The merchant&#39;s ID that is in GrabFood&#39;s database. | [optional] 
**PartnerMerchantID** | Pointer to **string** | The merchant&#39;s ID that is on the partner&#39;s database. | [optional] 
**JobID** | Pointer to **string** | An UUID string. Uniquely identifies a menu sync job. This can be found from the [Menu Update Notification](#tag/update-menu-noti) API response header.  | [optional] 
**UpdatedAt** | Pointer to **string** | Indicates the time of menu sync status change. This is based on ISO_8601/RFC3339. For example: &#x60;2022-07-29T15:55:59Z&#x60;.  | [optional] 
**Status** | Pointer to **string** | Indicates the state of the menu sync job. | [optional] 
**Errors** | Pointer to **[]string** | A string array of errors that occurred during processing. This array is empty if the status is not &#x60;FAILED&#x60;. | [optional] 

## Methods

### NewMenuSyncWebhookRequest

`func NewMenuSyncWebhookRequest() *MenuSyncWebhookRequest`

NewMenuSyncWebhookRequest instantiates a new MenuSyncWebhookRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMenuSyncWebhookRequestWithDefaults

`func NewMenuSyncWebhookRequestWithDefaults() *MenuSyncWebhookRequest`

NewMenuSyncWebhookRequestWithDefaults instantiates a new MenuSyncWebhookRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRequestID

`func (o *MenuSyncWebhookRequest) GetRequestID() string`

GetRequestID returns the RequestID field if non-nil, zero value otherwise.

### GetRequestIDOk

`func (o *MenuSyncWebhookRequest) GetRequestIDOk() (*string, bool)`

GetRequestIDOk returns a tuple with the RequestID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestID

`func (o *MenuSyncWebhookRequest) SetRequestID(v string)`

SetRequestID sets RequestID field to given value.

### HasRequestID

`func (o *MenuSyncWebhookRequest) HasRequestID() bool`

HasRequestID returns a boolean if a field has been set.

### GetMerchantID

`func (o *MenuSyncWebhookRequest) GetMerchantID() string`

GetMerchantID returns the MerchantID field if non-nil, zero value otherwise.

### GetMerchantIDOk

`func (o *MenuSyncWebhookRequest) GetMerchantIDOk() (*string, bool)`

GetMerchantIDOk returns a tuple with the MerchantID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMerchantID

`func (o *MenuSyncWebhookRequest) SetMerchantID(v string)`

SetMerchantID sets MerchantID field to given value.

### HasMerchantID

`func (o *MenuSyncWebhookRequest) HasMerchantID() bool`

HasMerchantID returns a boolean if a field has been set.

### GetPartnerMerchantID

`func (o *MenuSyncWebhookRequest) GetPartnerMerchantID() string`

GetPartnerMerchantID returns the PartnerMerchantID field if non-nil, zero value otherwise.

### GetPartnerMerchantIDOk

`func (o *MenuSyncWebhookRequest) GetPartnerMerchantIDOk() (*string, bool)`

GetPartnerMerchantIDOk returns a tuple with the PartnerMerchantID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPartnerMerchantID

`func (o *MenuSyncWebhookRequest) SetPartnerMerchantID(v string)`

SetPartnerMerchantID sets PartnerMerchantID field to given value.

### HasPartnerMerchantID

`func (o *MenuSyncWebhookRequest) HasPartnerMerchantID() bool`

HasPartnerMerchantID returns a boolean if a field has been set.

### GetJobID

`func (o *MenuSyncWebhookRequest) GetJobID() string`

GetJobID returns the JobID field if non-nil, zero value otherwise.

### GetJobIDOk

`func (o *MenuSyncWebhookRequest) GetJobIDOk() (*string, bool)`

GetJobIDOk returns a tuple with the JobID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJobID

`func (o *MenuSyncWebhookRequest) SetJobID(v string)`

SetJobID sets JobID field to given value.

### HasJobID

`func (o *MenuSyncWebhookRequest) HasJobID() bool`

HasJobID returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *MenuSyncWebhookRequest) GetUpdatedAt() string`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *MenuSyncWebhookRequest) GetUpdatedAtOk() (*string, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *MenuSyncWebhookRequest) SetUpdatedAt(v string)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *MenuSyncWebhookRequest) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.

### GetStatus

`func (o *MenuSyncWebhookRequest) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *MenuSyncWebhookRequest) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *MenuSyncWebhookRequest) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *MenuSyncWebhookRequest) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetErrors

`func (o *MenuSyncWebhookRequest) GetErrors() []string`

GetErrors returns the Errors field if non-nil, zero value otherwise.

### GetErrorsOk

`func (o *MenuSyncWebhookRequest) GetErrorsOk() (*[]string, bool)`

GetErrorsOk returns a tuple with the Errors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrors

`func (o *MenuSyncWebhookRequest) SetErrors(v []string)`

SetErrors sets Errors field to given value.

### HasErrors

`func (o *MenuSyncWebhookRequest) HasErrors() bool`

HasErrors returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


