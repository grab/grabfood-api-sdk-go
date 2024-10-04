# RegisterMembershipNativeRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**GrabID** | **string** | The id used to identify an unique grab user. | 
**PhoneNumber** | Pointer to **string** | Grab user&#39;s phone number for registration. | [optional] 
**Name** | Pointer to **string** | Grab user&#39;s name for registration. | [optional] 
**Email** | Pointer to **string** | Grab user&#39;s email address for registration. | [optional] 

## Methods

### NewRegisterMembershipNativeRequest

`func NewRegisterMembershipNativeRequest(grabID string, ) *RegisterMembershipNativeRequest`

NewRegisterMembershipNativeRequest instantiates a new RegisterMembershipNativeRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRegisterMembershipNativeRequestWithDefaults

`func NewRegisterMembershipNativeRequestWithDefaults() *RegisterMembershipNativeRequest`

NewRegisterMembershipNativeRequestWithDefaults instantiates a new RegisterMembershipNativeRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetGrabID

`func (o *RegisterMembershipNativeRequest) GetGrabID() string`

GetGrabID returns the GrabID field if non-nil, zero value otherwise.

### GetGrabIDOk

`func (o *RegisterMembershipNativeRequest) GetGrabIDOk() (*string, bool)`

GetGrabIDOk returns a tuple with the GrabID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGrabID

`func (o *RegisterMembershipNativeRequest) SetGrabID(v string)`

SetGrabID sets GrabID field to given value.


### GetPhoneNumber

`func (o *RegisterMembershipNativeRequest) GetPhoneNumber() string`

GetPhoneNumber returns the PhoneNumber field if non-nil, zero value otherwise.

### GetPhoneNumberOk

`func (o *RegisterMembershipNativeRequest) GetPhoneNumberOk() (*string, bool)`

GetPhoneNumberOk returns a tuple with the PhoneNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPhoneNumber

`func (o *RegisterMembershipNativeRequest) SetPhoneNumber(v string)`

SetPhoneNumber sets PhoneNumber field to given value.

### HasPhoneNumber

`func (o *RegisterMembershipNativeRequest) HasPhoneNumber() bool`

HasPhoneNumber returns a boolean if a field has been set.

### GetName

`func (o *RegisterMembershipNativeRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *RegisterMembershipNativeRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *RegisterMembershipNativeRequest) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *RegisterMembershipNativeRequest) HasName() bool`

HasName returns a boolean if a field has been set.

### GetEmail

`func (o *RegisterMembershipNativeRequest) GetEmail() string`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *RegisterMembershipNativeRequest) GetEmailOk() (*string, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *RegisterMembershipNativeRequest) SetEmail(v string)`

SetEmail sets Email field to given value.

### HasEmail

`func (o *RegisterMembershipNativeRequest) HasEmail() bool`

HasEmail returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


