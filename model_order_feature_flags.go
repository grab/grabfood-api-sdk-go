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

// checks if the OrderFeatureFlags type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &OrderFeatureFlags{}

// OrderFeatureFlags The featureFlag JSON object containing an order's feature related information.
type OrderFeatureFlags struct {
	// The acceptance type for the order. Refer to FAQs for more details about [orderAcceptedType](#section/Order/How-do-I-identify-if-a-particular-order-is-auto-or-manual-acceptance). 
	OrderAcceptedType string `json:"orderAcceptedType"`
	// The type of order. 
	OrderType string `json:"orderType"`
	// A boolean value that indicates if the order is edited. 
	IsMexEditOrder *bool `json:"isMexEditOrder,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _OrderFeatureFlags OrderFeatureFlags

// NewOrderFeatureFlags instantiates a new OrderFeatureFlags object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOrderFeatureFlags(orderAcceptedType string, orderType string) *OrderFeatureFlags {
	this := OrderFeatureFlags{}
	this.OrderAcceptedType = orderAcceptedType
	this.OrderType = orderType
	return &this
}

// NewOrderFeatureFlagsWithDefaults instantiates a new OrderFeatureFlags object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOrderFeatureFlagsWithDefaults() *OrderFeatureFlags {
	this := OrderFeatureFlags{}
	return &this
}

// GetOrderAcceptedType returns the OrderAcceptedType field value
func (o *OrderFeatureFlags) GetOrderAcceptedType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.OrderAcceptedType
}

// GetOrderAcceptedTypeOk returns a tuple with the OrderAcceptedType field value
// and a boolean to check if the value has been set.
func (o *OrderFeatureFlags) GetOrderAcceptedTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.OrderAcceptedType, true
}

// SetOrderAcceptedType sets field value
func (o *OrderFeatureFlags) SetOrderAcceptedType(v string) {
	o.OrderAcceptedType = v
}

// GetOrderType returns the OrderType field value
func (o *OrderFeatureFlags) GetOrderType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.OrderType
}

// GetOrderTypeOk returns a tuple with the OrderType field value
// and a boolean to check if the value has been set.
func (o *OrderFeatureFlags) GetOrderTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.OrderType, true
}

// SetOrderType sets field value
func (o *OrderFeatureFlags) SetOrderType(v string) {
	o.OrderType = v
}

// GetIsMexEditOrder returns the IsMexEditOrder field value if set, zero value otherwise.
func (o *OrderFeatureFlags) GetIsMexEditOrder() bool {
	if o == nil || IsNil(o.IsMexEditOrder) {
		var ret bool
		return ret
	}
	return *o.IsMexEditOrder
}

// GetIsMexEditOrderOk returns a tuple with the IsMexEditOrder field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrderFeatureFlags) GetIsMexEditOrderOk() (*bool, bool) {
	if o == nil || IsNil(o.IsMexEditOrder) {
		return nil, false
	}
	return o.IsMexEditOrder, true
}

// HasIsMexEditOrder returns a boolean if a field has been set.
func (o *OrderFeatureFlags) HasIsMexEditOrder() bool {
	if o != nil && !IsNil(o.IsMexEditOrder) {
		return true
	}

	return false
}

// SetIsMexEditOrder gets a reference to the given bool and assigns it to the IsMexEditOrder field.
func (o *OrderFeatureFlags) SetIsMexEditOrder(v bool) {
	o.IsMexEditOrder = &v
}

func (o OrderFeatureFlags) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o OrderFeatureFlags) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["orderAcceptedType"] = o.OrderAcceptedType
	toSerialize["orderType"] = o.OrderType
	if !IsNil(o.IsMexEditOrder) {
		toSerialize["isMexEditOrder"] = o.IsMexEditOrder
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *OrderFeatureFlags) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"orderAcceptedType",
		"orderType",
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

	varOrderFeatureFlags := _OrderFeatureFlags{}

	err = json.Unmarshal(data, &varOrderFeatureFlags)

	if err != nil {
		return err
	}

	*o = OrderFeatureFlags(varOrderFeatureFlags)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "orderAcceptedType")
		delete(additionalProperties, "orderType")
		delete(additionalProperties, "isMexEditOrder")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableOrderFeatureFlags struct {
	value *OrderFeatureFlags
	isSet bool
}

func (v NullableOrderFeatureFlags) Get() *OrderFeatureFlags {
	return v.value
}

func (v *NullableOrderFeatureFlags) Set(val *OrderFeatureFlags) {
	v.value = val
	v.isSet = true
}

func (v NullableOrderFeatureFlags) IsSet() bool {
	return v.isSet
}

func (v *NullableOrderFeatureFlags) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOrderFeatureFlags(val *OrderFeatureFlags) *NullableOrderFeatureFlags {
	return &NullableOrderFeatureFlags{value: val, isSet: true}
}

func (v NullableOrderFeatureFlags) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOrderFeatureFlags) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


