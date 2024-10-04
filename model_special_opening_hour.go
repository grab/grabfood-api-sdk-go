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

// checks if the SpecialOpeningHour type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SpecialOpeningHour{}

// SpecialOpeningHour struct for SpecialOpeningHour
type SpecialOpeningHour struct {
	// The start date of store special opening hours.
	StartDate *string `json:"startDate,omitempty"`
	// The end date of store special opening hours.
	EndDate *string `json:"endDate,omitempty"`
	Metadata *SpecialOpeningHourMetadata `json:"metadata,omitempty"`
	OpeningHours *SpecialOpeningHourOpeningHours `json:"openingHours,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _SpecialOpeningHour SpecialOpeningHour

// NewSpecialOpeningHour instantiates a new SpecialOpeningHour object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSpecialOpeningHour() *SpecialOpeningHour {
	this := SpecialOpeningHour{}
	return &this
}

// NewSpecialOpeningHourWithDefaults instantiates a new SpecialOpeningHour object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSpecialOpeningHourWithDefaults() *SpecialOpeningHour {
	this := SpecialOpeningHour{}
	return &this
}

// GetStartDate returns the StartDate field value if set, zero value otherwise.
func (o *SpecialOpeningHour) GetStartDate() string {
	if o == nil || IsNil(o.StartDate) {
		var ret string
		return ret
	}
	return *o.StartDate
}

// GetStartDateOk returns a tuple with the StartDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SpecialOpeningHour) GetStartDateOk() (*string, bool) {
	if o == nil || IsNil(o.StartDate) {
		return nil, false
	}
	return o.StartDate, true
}

// HasStartDate returns a boolean if a field has been set.
func (o *SpecialOpeningHour) HasStartDate() bool {
	if o != nil && !IsNil(o.StartDate) {
		return true
	}

	return false
}

// SetStartDate gets a reference to the given string and assigns it to the StartDate field.
func (o *SpecialOpeningHour) SetStartDate(v string) {
	o.StartDate = &v
}

// GetEndDate returns the EndDate field value if set, zero value otherwise.
func (o *SpecialOpeningHour) GetEndDate() string {
	if o == nil || IsNil(o.EndDate) {
		var ret string
		return ret
	}
	return *o.EndDate
}

// GetEndDateOk returns a tuple with the EndDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SpecialOpeningHour) GetEndDateOk() (*string, bool) {
	if o == nil || IsNil(o.EndDate) {
		return nil, false
	}
	return o.EndDate, true
}

// HasEndDate returns a boolean if a field has been set.
func (o *SpecialOpeningHour) HasEndDate() bool {
	if o != nil && !IsNil(o.EndDate) {
		return true
	}

	return false
}

// SetEndDate gets a reference to the given string and assigns it to the EndDate field.
func (o *SpecialOpeningHour) SetEndDate(v string) {
	o.EndDate = &v
}

// GetMetadata returns the Metadata field value if set, zero value otherwise.
func (o *SpecialOpeningHour) GetMetadata() SpecialOpeningHourMetadata {
	if o == nil || IsNil(o.Metadata) {
		var ret SpecialOpeningHourMetadata
		return ret
	}
	return *o.Metadata
}

// GetMetadataOk returns a tuple with the Metadata field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SpecialOpeningHour) GetMetadataOk() (*SpecialOpeningHourMetadata, bool) {
	if o == nil || IsNil(o.Metadata) {
		return nil, false
	}
	return o.Metadata, true
}

// HasMetadata returns a boolean if a field has been set.
func (o *SpecialOpeningHour) HasMetadata() bool {
	if o != nil && !IsNil(o.Metadata) {
		return true
	}

	return false
}

// SetMetadata gets a reference to the given SpecialOpeningHourMetadata and assigns it to the Metadata field.
func (o *SpecialOpeningHour) SetMetadata(v SpecialOpeningHourMetadata) {
	o.Metadata = &v
}

// GetOpeningHours returns the OpeningHours field value if set, zero value otherwise.
func (o *SpecialOpeningHour) GetOpeningHours() SpecialOpeningHourOpeningHours {
	if o == nil || IsNil(o.OpeningHours) {
		var ret SpecialOpeningHourOpeningHours
		return ret
	}
	return *o.OpeningHours
}

// GetOpeningHoursOk returns a tuple with the OpeningHours field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SpecialOpeningHour) GetOpeningHoursOk() (*SpecialOpeningHourOpeningHours, bool) {
	if o == nil || IsNil(o.OpeningHours) {
		return nil, false
	}
	return o.OpeningHours, true
}

// HasOpeningHours returns a boolean if a field has been set.
func (o *SpecialOpeningHour) HasOpeningHours() bool {
	if o != nil && !IsNil(o.OpeningHours) {
		return true
	}

	return false
}

// SetOpeningHours gets a reference to the given SpecialOpeningHourOpeningHours and assigns it to the OpeningHours field.
func (o *SpecialOpeningHour) SetOpeningHours(v SpecialOpeningHourOpeningHours) {
	o.OpeningHours = &v
}

func (o SpecialOpeningHour) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SpecialOpeningHour) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.StartDate) {
		toSerialize["startDate"] = o.StartDate
	}
	if !IsNil(o.EndDate) {
		toSerialize["endDate"] = o.EndDate
	}
	if !IsNil(o.Metadata) {
		toSerialize["metadata"] = o.Metadata
	}
	if !IsNil(o.OpeningHours) {
		toSerialize["openingHours"] = o.OpeningHours
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *SpecialOpeningHour) UnmarshalJSON(data []byte) (err error) {
	varSpecialOpeningHour := _SpecialOpeningHour{}

	err = json.Unmarshal(data, &varSpecialOpeningHour)

	if err != nil {
		return err
	}

	*o = SpecialOpeningHour(varSpecialOpeningHour)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "startDate")
		delete(additionalProperties, "endDate")
		delete(additionalProperties, "metadata")
		delete(additionalProperties, "openingHours")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableSpecialOpeningHour struct {
	value *SpecialOpeningHour
	isSet bool
}

func (v NullableSpecialOpeningHour) Get() *SpecialOpeningHour {
	return v.value
}

func (v *NullableSpecialOpeningHour) Set(val *SpecialOpeningHour) {
	v.value = val
	v.isSet = true
}

func (v NullableSpecialOpeningHour) IsSet() bool {
	return v.isSet
}

func (v *NullableSpecialOpeningHour) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSpecialOpeningHour(val *SpecialOpeningHour) *NullableSpecialOpeningHour {
	return &NullableSpecialOpeningHour{value: val, isSet: true}
}

func (v NullableSpecialOpeningHour) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSpecialOpeningHour) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


