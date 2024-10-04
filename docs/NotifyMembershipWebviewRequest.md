# NotifyMembershipWebviewRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**MemberID** | Pointer to **string** | The unique member ID on the partner&#39;s database. | [optional] 
**GrabID** | Pointer to **string** | The id used to identify an unique grab user. | [optional] 
**MerchantID** | Pointer to **string** | The merchant&#39;s ID that is in GrabFood&#39;s database. | [optional] 
**Action** | Pointer to **string** | Action completed in partner&#39;s webview. | [optional] 

## Methods

### NewNotifyMembershipWebviewRequest

`func NewNotifyMembershipWebviewRequest() *NotifyMembershipWebviewRequest`

NewNotifyMembershipWebviewRequest instantiates a new NotifyMembershipWebviewRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNotifyMembershipWebviewRequestWithDefaults

`func NewNotifyMembershipWebviewRequestWithDefaults() *NotifyMembershipWebviewRequest`

NewNotifyMembershipWebviewRequestWithDefaults instantiates a new NotifyMembershipWebviewRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMemberID

`func (o *NotifyMembershipWebviewRequest) GetMemberID() string`

GetMemberID returns the MemberID field if non-nil, zero value otherwise.

### GetMemberIDOk

`func (o *NotifyMembershipWebviewRequest) GetMemberIDOk() (*string, bool)`

GetMemberIDOk returns a tuple with the MemberID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMemberID

`func (o *NotifyMembershipWebviewRequest) SetMemberID(v string)`

SetMemberID sets MemberID field to given value.

### HasMemberID

`func (o *NotifyMembershipWebviewRequest) HasMemberID() bool`

HasMemberID returns a boolean if a field has been set.

### GetGrabID

`func (o *NotifyMembershipWebviewRequest) GetGrabID() string`

GetGrabID returns the GrabID field if non-nil, zero value otherwise.

### GetGrabIDOk

`func (o *NotifyMembershipWebviewRequest) GetGrabIDOk() (*string, bool)`

GetGrabIDOk returns a tuple with the GrabID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGrabID

`func (o *NotifyMembershipWebviewRequest) SetGrabID(v string)`

SetGrabID sets GrabID field to given value.

### HasGrabID

`func (o *NotifyMembershipWebviewRequest) HasGrabID() bool`

HasGrabID returns a boolean if a field has been set.

### GetMerchantID

`func (o *NotifyMembershipWebviewRequest) GetMerchantID() string`

GetMerchantID returns the MerchantID field if non-nil, zero value otherwise.

### GetMerchantIDOk

`func (o *NotifyMembershipWebviewRequest) GetMerchantIDOk() (*string, bool)`

GetMerchantIDOk returns a tuple with the MerchantID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMerchantID

`func (o *NotifyMembershipWebviewRequest) SetMerchantID(v string)`

SetMerchantID sets MerchantID field to given value.

### HasMerchantID

`func (o *NotifyMembershipWebviewRequest) HasMerchantID() bool`

HasMerchantID returns a boolean if a field has been set.

### GetAction

`func (o *NotifyMembershipWebviewRequest) GetAction() string`

GetAction returns the Action field if non-nil, zero value otherwise.

### GetActionOk

`func (o *NotifyMembershipWebviewRequest) GetActionOk() (*string, bool)`

GetActionOk returns a tuple with the Action field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAction

`func (o *NotifyMembershipWebviewRequest) SetAction(v string)`

SetAction sets Action field to given value.

### HasAction

`func (o *NotifyMembershipWebviewRequest) HasAction() bool`

HasAction returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


