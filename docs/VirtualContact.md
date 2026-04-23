# VirtualContact

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PhoneNumber** | Pointer to **string** | The generated virtual phone number that will forward calls to the customer&#39;s actual phone number. Will be omitted if the &#x60;status&#x60; is not &#x60;ACTIVE&#x60;. | [optional] 
**PIN** | Pointer to **string** | A unique PIN required when forwarding calls to the customer for verification and security. Will be omitted if the &#x60;status&#x60; is not &#x60;ACTIVE&#x60;. | [optional] 
**ExpiredAt** | Pointer to **string** | The expiry time of the virtual contact in UTC based on ISO_8601/RFC3339. Will be omitted if the &#x60;status&#x60; is not &#x60;ACTIVE&#x60;. | [optional] 
**Status** | Pointer to **string** | Indicates the current status and validity of the virtual contact. * &#x60;ACTIVE&#x60; - The virtual contact is valid. * &#x60;EXPIRED&#x60; - The virtual contact has expired and is no longer valid. * &#x60;UNAVAILABLE&#x60; - An internal error occurred while generating or retrieving the virtual contact.  | [optional] 

## Methods

### NewVirtualContact

`func NewVirtualContact() *VirtualContact`

NewVirtualContact instantiates a new VirtualContact object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVirtualContactWithDefaults

`func NewVirtualContactWithDefaults() *VirtualContact`

NewVirtualContactWithDefaults instantiates a new VirtualContact object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPhoneNumber

`func (o *VirtualContact) GetPhoneNumber() string`

GetPhoneNumber returns the PhoneNumber field if non-nil, zero value otherwise.

### GetPhoneNumberOk

`func (o *VirtualContact) GetPhoneNumberOk() (*string, bool)`

GetPhoneNumberOk returns a tuple with the PhoneNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPhoneNumber

`func (o *VirtualContact) SetPhoneNumber(v string)`

SetPhoneNumber sets PhoneNumber field to given value.

### HasPhoneNumber

`func (o *VirtualContact) HasPhoneNumber() bool`

HasPhoneNumber returns a boolean if a field has been set.

### GetPIN

`func (o *VirtualContact) GetPIN() string`

GetPIN returns the PIN field if non-nil, zero value otherwise.

### GetPINOk

`func (o *VirtualContact) GetPINOk() (*string, bool)`

GetPINOk returns a tuple with the PIN field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPIN

`func (o *VirtualContact) SetPIN(v string)`

SetPIN sets PIN field to given value.

### HasPIN

`func (o *VirtualContact) HasPIN() bool`

HasPIN returns a boolean if a field has been set.

### GetExpiredAt

`func (o *VirtualContact) GetExpiredAt() string`

GetExpiredAt returns the ExpiredAt field if non-nil, zero value otherwise.

### GetExpiredAtOk

`func (o *VirtualContact) GetExpiredAtOk() (*string, bool)`

GetExpiredAtOk returns a tuple with the ExpiredAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpiredAt

`func (o *VirtualContact) SetExpiredAt(v string)`

SetExpiredAt sets ExpiredAt field to given value.

### HasExpiredAt

`func (o *VirtualContact) HasExpiredAt() bool`

HasExpiredAt returns a boolean if a field has been set.

### GetStatus

`func (o *VirtualContact) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *VirtualContact) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *VirtualContact) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *VirtualContact) HasStatus() bool`

HasStatus returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


