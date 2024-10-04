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
	"time"
	"fmt"
)

// checks if the NewOrderTimeRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &NewOrderTimeRequest{}

// NewOrderTimeRequest This request updates an order with a new ready time on GrabFood. 
type NewOrderTimeRequest struct {
	// The order's ID that is returned from GrabFood. Refer to FAQs for more details about [orderID and shortOrderNumber](#section/Order/What's-the-difference-between-orderID-and-shortOrderNumber).
	OrderID string `json:"orderID"`
	// The new order ready time for this order, based on ISO_8601/RFC3339.
	NewOrderReadyTime time.Time `json:"newOrderReadyTime"`
	AdditionalProperties map[string]interface{}
}

type _NewOrderTimeRequest NewOrderTimeRequest

// NewNewOrderTimeRequest instantiates a new NewOrderTimeRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewNewOrderTimeRequest(orderID string, newOrderReadyTime time.Time) *NewOrderTimeRequest {
	this := NewOrderTimeRequest{}
	this.OrderID = orderID
	this.NewOrderReadyTime = newOrderReadyTime
	return &this
}

// NewNewOrderTimeRequestWithDefaults instantiates a new NewOrderTimeRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewNewOrderTimeRequestWithDefaults() *NewOrderTimeRequest {
	this := NewOrderTimeRequest{}
	return &this
}

// GetOrderID returns the OrderID field value
func (o *NewOrderTimeRequest) GetOrderID() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.OrderID
}

// GetOrderIDOk returns a tuple with the OrderID field value
// and a boolean to check if the value has been set.
func (o *NewOrderTimeRequest) GetOrderIDOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.OrderID, true
}

// SetOrderID sets field value
func (o *NewOrderTimeRequest) SetOrderID(v string) {
	o.OrderID = v
}

// GetNewOrderReadyTime returns the NewOrderReadyTime field value
func (o *NewOrderTimeRequest) GetNewOrderReadyTime() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.NewOrderReadyTime
}

// GetNewOrderReadyTimeOk returns a tuple with the NewOrderReadyTime field value
// and a boolean to check if the value has been set.
func (o *NewOrderTimeRequest) GetNewOrderReadyTimeOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.NewOrderReadyTime, true
}

// SetNewOrderReadyTime sets field value
func (o *NewOrderTimeRequest) SetNewOrderReadyTime(v time.Time) {
	o.NewOrderReadyTime = v
}

func (o NewOrderTimeRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o NewOrderTimeRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["orderID"] = o.OrderID
	toSerialize["newOrderReadyTime"] = o.NewOrderReadyTime

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *NewOrderTimeRequest) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"orderID",
		"newOrderReadyTime",
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

	varNewOrderTimeRequest := _NewOrderTimeRequest{}

	err = json.Unmarshal(data, &varNewOrderTimeRequest)

	if err != nil {
		return err
	}

	*o = NewOrderTimeRequest(varNewOrderTimeRequest)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "orderID")
		delete(additionalProperties, "newOrderReadyTime")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableNewOrderTimeRequest struct {
	value *NewOrderTimeRequest
	isSet bool
}

func (v NullableNewOrderTimeRequest) Get() *NewOrderTimeRequest {
	return v.value
}

func (v *NullableNewOrderTimeRequest) Set(val *NewOrderTimeRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableNewOrderTimeRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableNewOrderTimeRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableNewOrderTimeRequest(val *NewOrderTimeRequest) *NullableNewOrderTimeRequest {
	return &NullableNewOrderTimeRequest{value: val, isSet: true}
}

func (v NullableNewOrderTimeRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableNewOrderTimeRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


