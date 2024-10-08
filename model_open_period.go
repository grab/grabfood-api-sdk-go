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

// checks if the OpenPeriod type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &OpenPeriod{}

// OpenPeriod struct for OpenPeriod
type OpenPeriod struct {
	// The open start time in 24h format.
	StartTime string `json:"startTime"`
	// The open start time in 24h format.
	EndTime string `json:"endTime"`
	AdditionalProperties map[string]interface{}
}

type _OpenPeriod OpenPeriod

// NewOpenPeriod instantiates a new OpenPeriod object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOpenPeriod(startTime string, endTime string) *OpenPeriod {
	this := OpenPeriod{}
	this.StartTime = startTime
	this.EndTime = endTime
	return &this
}

// NewOpenPeriodWithDefaults instantiates a new OpenPeriod object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOpenPeriodWithDefaults() *OpenPeriod {
	this := OpenPeriod{}
	return &this
}

// GetStartTime returns the StartTime field value
func (o *OpenPeriod) GetStartTime() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.StartTime
}

// GetStartTimeOk returns a tuple with the StartTime field value
// and a boolean to check if the value has been set.
func (o *OpenPeriod) GetStartTimeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.StartTime, true
}

// SetStartTime sets field value
func (o *OpenPeriod) SetStartTime(v string) {
	o.StartTime = v
}

// GetEndTime returns the EndTime field value
func (o *OpenPeriod) GetEndTime() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.EndTime
}

// GetEndTimeOk returns a tuple with the EndTime field value
// and a boolean to check if the value has been set.
func (o *OpenPeriod) GetEndTimeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.EndTime, true
}

// SetEndTime sets field value
func (o *OpenPeriod) SetEndTime(v string) {
	o.EndTime = v
}

func (o OpenPeriod) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o OpenPeriod) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["startTime"] = o.StartTime
	toSerialize["endTime"] = o.EndTime

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *OpenPeriod) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"startTime",
		"endTime",
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

	varOpenPeriod := _OpenPeriod{}

	err = json.Unmarshal(data, &varOpenPeriod)

	if err != nil {
		return err
	}

	*o = OpenPeriod(varOpenPeriod)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "startTime")
		delete(additionalProperties, "endTime")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableOpenPeriod struct {
	value *OpenPeriod
	isSet bool
}

func (v NullableOpenPeriod) Get() *OpenPeriod {
	return v.value
}

func (v *NullableOpenPeriod) Set(val *OpenPeriod) {
	v.value = val
	v.isSet = true
}

func (v NullableOpenPeriod) IsSet() bool {
	return v.isSet
}

func (v *NullableOpenPeriod) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOpenPeriod(val *OpenPeriod) *NullableOpenPeriod {
	return &NullableOpenPeriod{value: val, isSet: true}
}

func (v NullableOpenPeriod) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOpenPeriod) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


