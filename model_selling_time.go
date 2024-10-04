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

// checks if the SellingTime type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SellingTime{}

// SellingTime struct for SellingTime
type SellingTime struct {
	// The selling time group start time in date & time format. 
	StartTime *string `json:"startTime,omitempty"`
	// The selling time group end time in date & time format. 
	EndTime *string `json:"endTime,omitempty"`
	// The selling time's ID on the partner system. This ID should be unique with length min 1 and max 64. 
	Id *string `json:"id,omitempty"`
	// The name of the selling time. 
	Name *string `json:"name,omitempty"`
	ServiceHours *ServiceHours `json:"serviceHours,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _SellingTime SellingTime

// NewSellingTime instantiates a new SellingTime object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSellingTime() *SellingTime {
	this := SellingTime{}
	return &this
}

// NewSellingTimeWithDefaults instantiates a new SellingTime object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSellingTimeWithDefaults() *SellingTime {
	this := SellingTime{}
	return &this
}

// GetStartTime returns the StartTime field value if set, zero value otherwise.
func (o *SellingTime) GetStartTime() string {
	if o == nil || IsNil(o.StartTime) {
		var ret string
		return ret
	}
	return *o.StartTime
}

// GetStartTimeOk returns a tuple with the StartTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SellingTime) GetStartTimeOk() (*string, bool) {
	if o == nil || IsNil(o.StartTime) {
		return nil, false
	}
	return o.StartTime, true
}

// HasStartTime returns a boolean if a field has been set.
func (o *SellingTime) HasStartTime() bool {
	if o != nil && !IsNil(o.StartTime) {
		return true
	}

	return false
}

// SetStartTime gets a reference to the given string and assigns it to the StartTime field.
func (o *SellingTime) SetStartTime(v string) {
	o.StartTime = &v
}

// GetEndTime returns the EndTime field value if set, zero value otherwise.
func (o *SellingTime) GetEndTime() string {
	if o == nil || IsNil(o.EndTime) {
		var ret string
		return ret
	}
	return *o.EndTime
}

// GetEndTimeOk returns a tuple with the EndTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SellingTime) GetEndTimeOk() (*string, bool) {
	if o == nil || IsNil(o.EndTime) {
		return nil, false
	}
	return o.EndTime, true
}

// HasEndTime returns a boolean if a field has been set.
func (o *SellingTime) HasEndTime() bool {
	if o != nil && !IsNil(o.EndTime) {
		return true
	}

	return false
}

// SetEndTime gets a reference to the given string and assigns it to the EndTime field.
func (o *SellingTime) SetEndTime(v string) {
	o.EndTime = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *SellingTime) GetId() string {
	if o == nil || IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SellingTime) GetIdOk() (*string, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *SellingTime) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *SellingTime) SetId(v string) {
	o.Id = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *SellingTime) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SellingTime) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *SellingTime) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *SellingTime) SetName(v string) {
	o.Name = &v
}

// GetServiceHours returns the ServiceHours field value if set, zero value otherwise.
func (o *SellingTime) GetServiceHours() ServiceHours {
	if o == nil || IsNil(o.ServiceHours) {
		var ret ServiceHours
		return ret
	}
	return *o.ServiceHours
}

// GetServiceHoursOk returns a tuple with the ServiceHours field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SellingTime) GetServiceHoursOk() (*ServiceHours, bool) {
	if o == nil || IsNil(o.ServiceHours) {
		return nil, false
	}
	return o.ServiceHours, true
}

// HasServiceHours returns a boolean if a field has been set.
func (o *SellingTime) HasServiceHours() bool {
	if o != nil && !IsNil(o.ServiceHours) {
		return true
	}

	return false
}

// SetServiceHours gets a reference to the given ServiceHours and assigns it to the ServiceHours field.
func (o *SellingTime) SetServiceHours(v ServiceHours) {
	o.ServiceHours = &v
}

func (o SellingTime) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SellingTime) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.StartTime) {
		toSerialize["startTime"] = o.StartTime
	}
	if !IsNil(o.EndTime) {
		toSerialize["endTime"] = o.EndTime
	}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !IsNil(o.ServiceHours) {
		toSerialize["serviceHours"] = o.ServiceHours
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *SellingTime) UnmarshalJSON(data []byte) (err error) {
	varSellingTime := _SellingTime{}

	err = json.Unmarshal(data, &varSellingTime)

	if err != nil {
		return err
	}

	*o = SellingTime(varSellingTime)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "startTime")
		delete(additionalProperties, "endTime")
		delete(additionalProperties, "id")
		delete(additionalProperties, "name")
		delete(additionalProperties, "serviceHours")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableSellingTime struct {
	value *SellingTime
	isSet bool
}

func (v NullableSellingTime) Get() *SellingTime {
	return v.value
}

func (v *NullableSellingTime) Set(val *SellingTime) {
	v.value = val
	v.isSet = true
}

func (v NullableSellingTime) IsSet() bool {
	return v.isSet
}

func (v *NullableSellingTime) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSellingTime(val *SellingTime) *NullableSellingTime {
	return &NullableSellingTime{value: val, isSet: true}
}

func (v NullableSellingTime) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSellingTime) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


