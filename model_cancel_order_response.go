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

// checks if the CancelOrderResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CancelOrderResponse{}

// CancelOrderResponse 
type CancelOrderResponse struct {
	LimitType *CancelOrderLimitType `json:"limitType,omitempty"`
	// The remaining cancellation quota for the merchant. A value is only returned when the nearest remaining cancellation limit is approaching, else it returns 0.
	LimitTimes *int64 `json:"limitTimes,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _CancelOrderResponse CancelOrderResponse

// NewCancelOrderResponse instantiates a new CancelOrderResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCancelOrderResponse() *CancelOrderResponse {
	this := CancelOrderResponse{}
	return &this
}

// NewCancelOrderResponseWithDefaults instantiates a new CancelOrderResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCancelOrderResponseWithDefaults() *CancelOrderResponse {
	this := CancelOrderResponse{}
	return &this
}

// GetLimitType returns the LimitType field value if set, zero value otherwise.
func (o *CancelOrderResponse) GetLimitType() CancelOrderLimitType {
	if o == nil || IsNil(o.LimitType) {
		var ret CancelOrderLimitType
		return ret
	}
	return *o.LimitType
}

// GetLimitTypeOk returns a tuple with the LimitType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CancelOrderResponse) GetLimitTypeOk() (*CancelOrderLimitType, bool) {
	if o == nil || IsNil(o.LimitType) {
		return nil, false
	}
	return o.LimitType, true
}

// HasLimitType returns a boolean if a field has been set.
func (o *CancelOrderResponse) HasLimitType() bool {
	if o != nil && !IsNil(o.LimitType) {
		return true
	}

	return false
}

// SetLimitType gets a reference to the given CancelOrderLimitType and assigns it to the LimitType field.
func (o *CancelOrderResponse) SetLimitType(v CancelOrderLimitType) {
	o.LimitType = &v
}

// GetLimitTimes returns the LimitTimes field value if set, zero value otherwise.
func (o *CancelOrderResponse) GetLimitTimes() int64 {
	if o == nil || IsNil(o.LimitTimes) {
		var ret int64
		return ret
	}
	return *o.LimitTimes
}

// GetLimitTimesOk returns a tuple with the LimitTimes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CancelOrderResponse) GetLimitTimesOk() (*int64, bool) {
	if o == nil || IsNil(o.LimitTimes) {
		return nil, false
	}
	return o.LimitTimes, true
}

// HasLimitTimes returns a boolean if a field has been set.
func (o *CancelOrderResponse) HasLimitTimes() bool {
	if o != nil && !IsNil(o.LimitTimes) {
		return true
	}

	return false
}

// SetLimitTimes gets a reference to the given int64 and assigns it to the LimitTimes field.
func (o *CancelOrderResponse) SetLimitTimes(v int64) {
	o.LimitTimes = &v
}

func (o CancelOrderResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CancelOrderResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.LimitType) {
		toSerialize["limitType"] = o.LimitType
	}
	if !IsNil(o.LimitTimes) {
		toSerialize["limitTimes"] = o.LimitTimes
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *CancelOrderResponse) UnmarshalJSON(data []byte) (err error) {
	varCancelOrderResponse := _CancelOrderResponse{}

	err = json.Unmarshal(data, &varCancelOrderResponse)

	if err != nil {
		return err
	}

	*o = CancelOrderResponse(varCancelOrderResponse)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "limitType")
		delete(additionalProperties, "limitTimes")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableCancelOrderResponse struct {
	value *CancelOrderResponse
	isSet bool
}

func (v NullableCancelOrderResponse) Get() *CancelOrderResponse {
	return v.value
}

func (v *NullableCancelOrderResponse) Set(val *CancelOrderResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableCancelOrderResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableCancelOrderResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCancelOrderResponse(val *CancelOrderResponse) *NullableCancelOrderResponse {
	return &NullableCancelOrderResponse{value: val, isSet: true}
}

func (v NullableCancelOrderResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCancelOrderResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


