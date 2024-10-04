# BindMembershipNativeRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**GrabID** | **string** | The id used to identify an unique grab user. | 
**PhoneNumber** | Pointer to **string** | Grab user phone number used to login. | [optional] 

## Methods

### NewBindMembershipNativeRequest

`func NewBindMembershipNativeRequest(grabID string, ) *BindMembershipNativeRequest`

NewBindMembershipNativeRequest instantiates a new BindMembershipNativeRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBindMembershipNativeRequestWithDefaults

`func NewBindMembershipNativeRequestWithDefaults() *BindMembershipNativeRequest`

NewBindMembershipNativeRequestWithDefaults instantiates a new BindMembershipNativeRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetGrabID

`func (o *BindMembershipNativeRequest) GetGrabID() string`

GetGrabID returns the GrabID field if non-nil, zero value otherwise.

### GetGrabIDOk

`func (o *BindMembershipNativeRequest) GetGrabIDOk() (*string, bool)`

GetGrabIDOk returns a tuple with the GrabID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGrabID

`func (o *BindMembershipNativeRequest) SetGrabID(v string)`

SetGrabID sets GrabID field to given value.


### GetPhoneNumber

`func (o *BindMembershipNativeRequest) GetPhoneNumber() string`

GetPhoneNumber returns the PhoneNumber field if non-nil, zero value otherwise.

### GetPhoneNumberOk

`func (o *BindMembershipNativeRequest) GetPhoneNumberOk() (*string, bool)`

GetPhoneNumberOk returns a tuple with the PhoneNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPhoneNumber

`func (o *BindMembershipNativeRequest) SetPhoneNumber(v string)`

SetPhoneNumber sets PhoneNumber field to given value.

### HasPhoneNumber

`func (o *BindMembershipNativeRequest) HasPhoneNumber() bool`

HasPhoneNumber returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


