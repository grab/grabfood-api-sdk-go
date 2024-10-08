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

// checks if the UpdateSpecialHourResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UpdateSpecialHourResponse{}

// UpdateSpecialHourResponse Object contain update store special hour response
type UpdateSpecialHourResponse struct {
	// Error message when updating store delivery hour. `null` indicates no error.
	ErrorReasons []string `json:"errorReasons,omitempty"`
	// Total active order for the day.
	OrderCount int64 `json:"orderCount"`
	// Total scheduled order during store close period.
	ScheduledOrderCount int64 `json:"scheduledOrderCount"`
	// Indicate the store status after updating special hours.
	CloseImmediately bool `json:"closeImmediately"`
	AdditionalProperties map[string]interface{}
}

type _UpdateSpecialHourResponse UpdateSpecialHourResponse

// NewUpdateSpecialHourResponse instantiates a new UpdateSpecialHourResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUpdateSpecialHourResponse(orderCount int64, scheduledOrderCount int64, closeImmediately bool) *UpdateSpecialHourResponse {
	this := UpdateSpecialHourResponse{}
	this.OrderCount = orderCount
	this.ScheduledOrderCount = scheduledOrderCount
	this.CloseImmediately = closeImmediately
	return &this
}

// NewUpdateSpecialHourResponseWithDefaults instantiates a new UpdateSpecialHourResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUpdateSpecialHourResponseWithDefaults() *UpdateSpecialHourResponse {
	this := UpdateSpecialHourResponse{}
	return &this
}

// GetErrorReasons returns the ErrorReasons field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *UpdateSpecialHourResponse) GetErrorReasons() []string {
	if o == nil {
		var ret []string
		return ret
	}
	return o.ErrorReasons
}

// GetErrorReasonsOk returns a tuple with the ErrorReasons field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *UpdateSpecialHourResponse) GetErrorReasonsOk() ([]string, bool) {
	if o == nil || IsNil(o.ErrorReasons) {
		return nil, false
	}
	return o.ErrorReasons, true
}

// HasErrorReasons returns a boolean if a field has been set.
func (o *UpdateSpecialHourResponse) HasErrorReasons() bool {
	if o != nil && !IsNil(o.ErrorReasons) {
		return true
	}

	return false
}

// SetErrorReasons gets a reference to the given []string and assigns it to the ErrorReasons field.
func (o *UpdateSpecialHourResponse) SetErrorReasons(v []string) {
	o.ErrorReasons = v
}

// GetOrderCount returns the OrderCount field value
func (o *UpdateSpecialHourResponse) GetOrderCount() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.OrderCount
}

// GetOrderCountOk returns a tuple with the OrderCount field value
// and a boolean to check if the value has been set.
func (o *UpdateSpecialHourResponse) GetOrderCountOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.OrderCount, true
}

// SetOrderCount sets field value
func (o *UpdateSpecialHourResponse) SetOrderCount(v int64) {
	o.OrderCount = v
}

// GetScheduledOrderCount returns the ScheduledOrderCount field value
func (o *UpdateSpecialHourResponse) GetScheduledOrderCount() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.ScheduledOrderCount
}

// GetScheduledOrderCountOk returns a tuple with the ScheduledOrderCount field value
// and a boolean to check if the value has been set.
func (o *UpdateSpecialHourResponse) GetScheduledOrderCountOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ScheduledOrderCount, true
}

// SetScheduledOrderCount sets field value
func (o *UpdateSpecialHourResponse) SetScheduledOrderCount(v int64) {
	o.ScheduledOrderCount = v
}

// GetCloseImmediately returns the CloseImmediately field value
func (o *UpdateSpecialHourResponse) GetCloseImmediately() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.CloseImmediately
}

// GetCloseImmediatelyOk returns a tuple with the CloseImmediately field value
// and a boolean to check if the value has been set.
func (o *UpdateSpecialHourResponse) GetCloseImmediatelyOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CloseImmediately, true
}

// SetCloseImmediately sets field value
func (o *UpdateSpecialHourResponse) SetCloseImmediately(v bool) {
	o.CloseImmediately = v
}

func (o UpdateSpecialHourResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UpdateSpecialHourResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.ErrorReasons != nil {
		toSerialize["errorReasons"] = o.ErrorReasons
	}
	toSerialize["orderCount"] = o.OrderCount
	toSerialize["scheduledOrderCount"] = o.ScheduledOrderCount
	toSerialize["closeImmediately"] = o.CloseImmediately

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *UpdateSpecialHourResponse) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"orderCount",
		"scheduledOrderCount",
		"closeImmediately",
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

	varUpdateSpecialHourResponse := _UpdateSpecialHourResponse{}

	err = json.Unmarshal(data, &varUpdateSpecialHourResponse)

	if err != nil {
		return err
	}

	*o = UpdateSpecialHourResponse(varUpdateSpecialHourResponse)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "errorReasons")
		delete(additionalProperties, "orderCount")
		delete(additionalProperties, "scheduledOrderCount")
		delete(additionalProperties, "closeImmediately")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableUpdateSpecialHourResponse struct {
	value *UpdateSpecialHourResponse
	isSet bool
}

func (v NullableUpdateSpecialHourResponse) Get() *UpdateSpecialHourResponse {
	return v.value
}

func (v *NullableUpdateSpecialHourResponse) Set(val *UpdateSpecialHourResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableUpdateSpecialHourResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableUpdateSpecialHourResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUpdateSpecialHourResponse(val *UpdateSpecialHourResponse) *NullableUpdateSpecialHourResponse {
	return &NullableUpdateSpecialHourResponse{value: val, isSet: true}
}

func (v NullableUpdateSpecialHourResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUpdateSpecialHourResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


