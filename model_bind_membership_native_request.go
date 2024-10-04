// Copyright 2024 Grabtaxi Holdings PTE LTE (GRAB), All rights reserved.
// Use of this source code is governed by an MIT-style license that can be found in the LICENSE file

/*
GrabFood

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: 1.1.3
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package grabfood

import (
	"encoding/json"
	"fmt"
)

// checks if the BindMembershipNativeRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &BindMembershipNativeRequest{}

// BindMembershipNativeRequest This request submits membership binding request to partner. 
type BindMembershipNativeRequest struct {
	// The id used to identify an unique grab user.
	GrabID string `json:"grabID"`
	// Grab user phone number used to login.
	PhoneNumber *string `json:"phoneNumber,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _BindMembershipNativeRequest BindMembershipNativeRequest

// NewBindMembershipNativeRequest instantiates a new BindMembershipNativeRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBindMembershipNativeRequest(grabID string) *BindMembershipNativeRequest {
	this := BindMembershipNativeRequest{}
	this.GrabID = grabID
	return &this
}

// NewBindMembershipNativeRequestWithDefaults instantiates a new BindMembershipNativeRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBindMembershipNativeRequestWithDefaults() *BindMembershipNativeRequest {
	this := BindMembershipNativeRequest{}
	return &this
}

// GetGrabID returns the GrabID field value
func (o *BindMembershipNativeRequest) GetGrabID() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.GrabID
}

// GetGrabIDOk returns a tuple with the GrabID field value
// and a boolean to check if the value has been set.
func (o *BindMembershipNativeRequest) GetGrabIDOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.GrabID, true
}

// SetGrabID sets field value
func (o *BindMembershipNativeRequest) SetGrabID(v string) {
	o.GrabID = v
}

// GetPhoneNumber returns the PhoneNumber field value if set, zero value otherwise.
func (o *BindMembershipNativeRequest) GetPhoneNumber() string {
	if o == nil || IsNil(o.PhoneNumber) {
		var ret string
		return ret
	}
	return *o.PhoneNumber
}

// GetPhoneNumberOk returns a tuple with the PhoneNumber field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BindMembershipNativeRequest) GetPhoneNumberOk() (*string, bool) {
	if o == nil || IsNil(o.PhoneNumber) {
		return nil, false
	}
	return o.PhoneNumber, true
}

// HasPhoneNumber returns a boolean if a field has been set.
func (o *BindMembershipNativeRequest) HasPhoneNumber() bool {
	if o != nil && !IsNil(o.PhoneNumber) {
		return true
	}

	return false
}

// SetPhoneNumber gets a reference to the given string and assigns it to the PhoneNumber field.
func (o *BindMembershipNativeRequest) SetPhoneNumber(v string) {
	o.PhoneNumber = &v
}

func (o BindMembershipNativeRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o BindMembershipNativeRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["grabID"] = o.GrabID
	if !IsNil(o.PhoneNumber) {
		toSerialize["phoneNumber"] = o.PhoneNumber
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *BindMembershipNativeRequest) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"grabID",
	}

	allProperties := make(map[string]interface{})

	err = json.Unmarshal(data, &allProperties)

	if err != nil {
		return err;
	}

	for _, requiredProperty := range(requiredProperties) {
		if _, exists := allProperties[requiredProperty]; !exists {
			return fmt.Errorf("no value given for required property %v", requiredProperty)
		}
	}

	varBindMembershipNativeRequest := _BindMembershipNativeRequest{}

	err = json.Unmarshal(data, &varBindMembershipNativeRequest)

	if err != nil {
		return err
	}

	*o = BindMembershipNativeRequest(varBindMembershipNativeRequest)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "grabID")
		delete(additionalProperties, "phoneNumber")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableBindMembershipNativeRequest struct {
	value *BindMembershipNativeRequest
	isSet bool
}

func (v NullableBindMembershipNativeRequest) Get() *BindMembershipNativeRequest {
	return v.value
}

func (v *NullableBindMembershipNativeRequest) Set(val *BindMembershipNativeRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableBindMembershipNativeRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableBindMembershipNativeRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBindMembershipNativeRequest(val *BindMembershipNativeRequest) *NullableBindMembershipNativeRequest {
	return &NullableBindMembershipNativeRequest{value: val, isSet: true}
}

func (v NullableBindMembershipNativeRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBindMembershipNativeRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


