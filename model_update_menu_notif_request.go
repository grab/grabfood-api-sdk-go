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

// checks if the UpdateMenuNotifRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UpdateMenuNotifRequest{}

// UpdateMenuNotifRequest This request notifies GrabFood about the partner's updated food menu. 
type UpdateMenuNotifRequest struct {
	// The merchant's ID that is in GrabFood's database.
	MerchantID string `json:"merchantID"`
	AdditionalProperties map[string]interface{}
}

type _UpdateMenuNotifRequest UpdateMenuNotifRequest

// NewUpdateMenuNotifRequest instantiates a new UpdateMenuNotifRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUpdateMenuNotifRequest(merchantID string) *UpdateMenuNotifRequest {
	this := UpdateMenuNotifRequest{}
	this.MerchantID = merchantID
	return &this
}

// NewUpdateMenuNotifRequestWithDefaults instantiates a new UpdateMenuNotifRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUpdateMenuNotifRequestWithDefaults() *UpdateMenuNotifRequest {
	this := UpdateMenuNotifRequest{}
	return &this
}

// GetMerchantID returns the MerchantID field value
func (o *UpdateMenuNotifRequest) GetMerchantID() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.MerchantID
}

// GetMerchantIDOk returns a tuple with the MerchantID field value
// and a boolean to check if the value has been set.
func (o *UpdateMenuNotifRequest) GetMerchantIDOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.MerchantID, true
}

// SetMerchantID sets field value
func (o *UpdateMenuNotifRequest) SetMerchantID(v string) {
	o.MerchantID = v
}

func (o UpdateMenuNotifRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UpdateMenuNotifRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["merchantID"] = o.MerchantID

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *UpdateMenuNotifRequest) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"merchantID",
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

	varUpdateMenuNotifRequest := _UpdateMenuNotifRequest{}

	err = json.Unmarshal(data, &varUpdateMenuNotifRequest)

	if err != nil {
		return err
	}

	*o = UpdateMenuNotifRequest(varUpdateMenuNotifRequest)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "merchantID")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableUpdateMenuNotifRequest struct {
	value *UpdateMenuNotifRequest
	isSet bool
}

func (v NullableUpdateMenuNotifRequest) Get() *UpdateMenuNotifRequest {
	return v.value
}

func (v *NullableUpdateMenuNotifRequest) Set(val *UpdateMenuNotifRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableUpdateMenuNotifRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableUpdateMenuNotifRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUpdateMenuNotifRequest(val *UpdateMenuNotifRequest) *NullableUpdateMenuNotifRequest {
	return &NullableUpdateMenuNotifRequest{value: val, isSet: true}
}

func (v NullableUpdateMenuNotifRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUpdateMenuNotifRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


