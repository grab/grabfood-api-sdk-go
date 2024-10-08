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

// checks if the GetMembershipRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetMembershipRequest{}

// GetMembershipRequest This request submits membership unbind request to partner. 
type GetMembershipRequest struct {
	// The unique member ID on the partner's database.
	MemberID string `json:"memberID"`
	AdditionalProperties map[string]interface{}
}

type _GetMembershipRequest GetMembershipRequest

// NewGetMembershipRequest instantiates a new GetMembershipRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetMembershipRequest(memberID string) *GetMembershipRequest {
	this := GetMembershipRequest{}
	this.MemberID = memberID
	return &this
}

// NewGetMembershipRequestWithDefaults instantiates a new GetMembershipRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetMembershipRequestWithDefaults() *GetMembershipRequest {
	this := GetMembershipRequest{}
	return &this
}

// GetMemberID returns the MemberID field value
func (o *GetMembershipRequest) GetMemberID() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.MemberID
}

// GetMemberIDOk returns a tuple with the MemberID field value
// and a boolean to check if the value has been set.
func (o *GetMembershipRequest) GetMemberIDOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.MemberID, true
}

// SetMemberID sets field value
func (o *GetMembershipRequest) SetMemberID(v string) {
	o.MemberID = v
}

func (o GetMembershipRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetMembershipRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["memberID"] = o.MemberID

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *GetMembershipRequest) UnmarshalJSON(data []byte) (err error) {
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

	varGetMembershipRequest := _GetMembershipRequest{}

	err = json.Unmarshal(data, &varGetMembershipRequest)

	if err != nil {
		return err
	}

	*o = GetMembershipRequest(varGetMembershipRequest)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "memberID")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableGetMembershipRequest struct {
	value *GetMembershipRequest
	isSet bool
}

func (v NullableGetMembershipRequest) Get() *GetMembershipRequest {
	return v.value
}

func (v *NullableGetMembershipRequest) Set(val *GetMembershipRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableGetMembershipRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableGetMembershipRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetMembershipRequest(val *GetMembershipRequest) *NullableGetMembershipRequest {
	return &NullableGetMembershipRequest{value: val, isSet: true}
}

func (v NullableGetMembershipRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetMembershipRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


