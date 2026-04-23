# Receiver

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** | The name of the receiver. | [optional] 
**Phones** | Pointer to **string** | The receiver&#39;s phone number. Only applicable for orders that are delivered by the restaurant. &#x60;null&#x60; if not applicable.  &gt; Note: The &#x60;phones&#x60; field will be deprecated once the virtualContact feature is fully rolled out.  | [optional] 
**Address** | Pointer to [**Address**](Address.md) |  | [optional] 
**VirtualContact** | Pointer to [**VirtualContact**](VirtualContact.md) |  | [optional] 

## Methods

### NewReceiver

`func NewReceiver() *Receiver`

NewReceiver instantiates a new Receiver object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewReceiverWithDefaults

`func NewReceiverWithDefaults() *Receiver`

NewReceiverWithDefaults instantiates a new Receiver object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *Receiver) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *Receiver) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *Receiver) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *Receiver) HasName() bool`

HasName returns a boolean if a field has been set.

### GetPhones

`func (o *Receiver) GetPhones() string`

GetPhones returns the Phones field if non-nil, zero value otherwise.

### GetPhonesOk

`func (o *Receiver) GetPhonesOk() (*string, bool)`

GetPhonesOk returns a tuple with the Phones field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPhones

`func (o *Receiver) SetPhones(v string)`

SetPhones sets Phones field to given value.

### HasPhones

`func (o *Receiver) HasPhones() bool`

HasPhones returns a boolean if a field has been set.

### GetAddress

`func (o *Receiver) GetAddress() Address`

GetAddress returns the Address field if non-nil, zero value otherwise.

### GetAddressOk

`func (o *Receiver) GetAddressOk() (*Address, bool)`

GetAddressOk returns a tuple with the Address field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddress

`func (o *Receiver) SetAddress(v Address)`

SetAddress sets Address field to given value.

### HasAddress

`func (o *Receiver) HasAddress() bool`

HasAddress returns a boolean if a field has been set.

### GetVirtualContact

`func (o *Receiver) GetVirtualContact() VirtualContact`

GetVirtualContact returns the VirtualContact field if non-nil, zero value otherwise.

### GetVirtualContactOk

`func (o *Receiver) GetVirtualContactOk() (*VirtualContact, bool)`

GetVirtualContactOk returns a tuple with the VirtualContact field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVirtualContact

`func (o *Receiver) SetVirtualContact(v VirtualContact)`

SetVirtualContact sets VirtualContact field to given value.

### HasVirtualContact

`func (o *Receiver) HasVirtualContact() bool`

HasVirtualContact returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


