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
)

// checks if the UpdatePurchasability type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UpdatePurchasability{}

// UpdatePurchasability struct for UpdatePurchasability
type UpdatePurchasability struct {
	// Available service type.
	Key *string `json:"key,omitempty"`
	Purchasable *bool `json:"purchasable,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _UpdatePurchasability UpdatePurchasability

// NewUpdatePurchasability instantiates a new UpdatePurchasability object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUpdatePurchasability() *UpdatePurchasability {
	this := UpdatePurchasability{}
	return &this
}

// NewUpdatePurchasabilityWithDefaults instantiates a new UpdatePurchasability object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUpdatePurchasabilityWithDefaults() *UpdatePurchasability {
	this := UpdatePurchasability{}
	return &this
}

// GetKey returns the Key field value if set, zero value otherwise.
func (o *UpdatePurchasability) GetKey() string {
	if o == nil || IsNil(o.Key) {
		var ret string
		return ret
	}
	return *o.Key
}

// GetKeyOk returns a tuple with the Key field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdatePurchasability) GetKeyOk() (*string, bool) {
	if o == nil || IsNil(o.Key) {
		return nil, false
	}
	return o.Key, true
}

// HasKey returns a boolean if a field has been set.
func (o *UpdatePurchasability) HasKey() bool {
	if o != nil && !IsNil(o.Key) {
		return true
	}

	return false
}

// SetKey gets a reference to the given string and assigns it to the Key field.
func (o *UpdatePurchasability) SetKey(v string) {
	o.Key = &v
}

// GetPurchasable returns the Purchasable field value if set, zero value otherwise.
func (o *UpdatePurchasability) GetPurchasable() bool {
	if o == nil || IsNil(o.Purchasable) {
		var ret bool
		return ret
	}
	return *o.Purchasable
}

// GetPurchasableOk returns a tuple with the Purchasable field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdatePurchasability) GetPurchasableOk() (*bool, bool) {
	if o == nil || IsNil(o.Purchasable) {
		return nil, false
	}
	return o.Purchasable, true
}

// HasPurchasable returns a boolean if a field has been set.
func (o *UpdatePurchasability) HasPurchasable() bool {
	if o != nil && !IsNil(o.Purchasable) {
		return true
	}

	return false
}

// SetPurchasable gets a reference to the given bool and assigns it to the Purchasable field.
func (o *UpdatePurchasability) SetPurchasable(v bool) {
	o.Purchasable = &v
}

func (o UpdatePurchasability) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UpdatePurchasability) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Key) {
		toSerialize["key"] = o.Key
	}
	if !IsNil(o.Purchasable) {
		toSerialize["purchasable"] = o.Purchasable
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *UpdatePurchasability) UnmarshalJSON(data []byte) (err error) {
	varUpdatePurchasability := _UpdatePurchasability{}

	err = json.Unmarshal(data, &varUpdatePurchasability)

	if err != nil {
		return err
	}

	*o = UpdatePurchasability(varUpdatePurchasability)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "key")
		delete(additionalProperties, "purchasable")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableUpdatePurchasability struct {
	value *UpdatePurchasability
	isSet bool
}

func (v NullableUpdatePurchasability) Get() *UpdatePurchasability {
	return v.value
}

func (v *NullableUpdatePurchasability) Set(val *UpdatePurchasability) {
	v.value = val
	v.isSet = true
}

func (v NullableUpdatePurchasability) IsSet() bool {
	return v.isSet
}

func (v *NullableUpdatePurchasability) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUpdatePurchasability(val *UpdatePurchasability) *NullableUpdatePurchasability {
	return &NullableUpdatePurchasability{value: val, isSet: true}
}

func (v NullableUpdatePurchasability) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUpdatePurchasability) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


