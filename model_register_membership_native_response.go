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

// checks if the RegisterMembershipNativeResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &RegisterMembershipNativeResponse{}

// RegisterMembershipNativeResponse This response returns membershipID after membership binding is successful. 
type RegisterMembershipNativeResponse struct {
	// The unique member ID on the partner's database.
	MemberID string `json:"memberID"`
	AdditionalProperties map[string]interface{}
}

type _RegisterMembershipNativeResponse RegisterMembershipNativeResponse

// NewRegisterMembershipNativeResponse instantiates a new RegisterMembershipNativeResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRegisterMembershipNativeResponse(memberID string) *RegisterMembershipNativeResponse {
	this := RegisterMembershipNativeResponse{}
	this.MemberID = memberID
	return &this
}

// NewRegisterMembershipNativeResponseWithDefaults instantiates a new RegisterMembershipNativeResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRegisterMembershipNativeResponseWithDefaults() *RegisterMembershipNativeResponse {
	this := RegisterMembershipNativeResponse{}
	return &this
}

// GetMemberID returns the MemberID field value
func (o *RegisterMembershipNativeResponse) GetMemberID() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.MemberID
}

// GetMemberIDOk returns a tuple with the MemberID field value
// and a boolean to check if the value has been set.
func (o *RegisterMembershipNativeResponse) GetMemberIDOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.MemberID, true
}

// SetMemberID sets field value
func (o *RegisterMembershipNativeResponse) SetMemberID(v string) {
	o.MemberID = v
}

func (o RegisterMembershipNativeResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o RegisterMembershipNativeResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["memberID"] = o.MemberID

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *RegisterMembershipNativeResponse) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"memberID",
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

	varRegisterMembershipNativeResponse := _RegisterMembershipNativeResponse{}

	err = json.Unmarshal(data, &varRegisterMembershipNativeResponse)

	if err != nil {
		return err
	}

	*o = RegisterMembershipNativeResponse(varRegisterMembershipNativeResponse)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "memberID")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableRegisterMembershipNativeResponse struct {
	value *RegisterMembershipNativeResponse
	isSet bool
}

func (v NullableRegisterMembershipNativeResponse) Get() *RegisterMembershipNativeResponse {
	return v.value
}

func (v *NullableRegisterMembershipNativeResponse) Set(val *RegisterMembershipNativeResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableRegisterMembershipNativeResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableRegisterMembershipNativeResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRegisterMembershipNativeResponse(val *RegisterMembershipNativeResponse) *NullableRegisterMembershipNativeResponse {
	return &NullableRegisterMembershipNativeResponse{value: val, isSet: true}
}

func (v NullableRegisterMembershipNativeResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRegisterMembershipNativeResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


