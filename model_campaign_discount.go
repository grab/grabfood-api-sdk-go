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

// checks if the CampaignDiscount type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CampaignDiscount{}

// CampaignDiscount The discount detail for a particular campaign when conditions are satisfied.
type CampaignDiscount struct {
	// The type of discount 
	Type string `json:"type"`
	// The maximum discount dollar amount. It is **not required** and will be ignored when the `discount.type` is: - `net` - `delivery` - `freeItem` - `bundleSameNet` - `bundleSamePercentage` - `bundleSameFixPrice` - `bundleDiffNet` - `bundleDiffPercentage` - `bundleDiffFixPrice` 
	Cap *float64 `json:"cap,omitempty"`
	// Specify the discount amount. Decimal number is not supported For VN, ID and TH. For example, `10.5` is not allowed and it should be `10.0`. * Dollar amount value when `discount.type` is `net`, `delivery`, `bundleSameNet`, `bundleSameFixPrice`, `bundleDiffNet`, `bundleDiffFixPrice`. * Percentage value (0-100) when `discount.type` is `percentage`, `bundleSamePercentage`, `bundleDiffPercentage`. * **Not required** when `discount.type` is `freeItem`. 
	Value *float64 `json:"value,omitempty"`
	Scope CampaignScope `json:"scope"`
	AdditionalProperties map[string]interface{}
}

type _CampaignDiscount CampaignDiscount

// NewCampaignDiscount instantiates a new CampaignDiscount object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCampaignDiscount(type_ string, scope CampaignScope) *CampaignDiscount {
	this := CampaignDiscount{}
	this.Type = type_
	this.Scope = scope
	return &this
}

// NewCampaignDiscountWithDefaults instantiates a new CampaignDiscount object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCampaignDiscountWithDefaults() *CampaignDiscount {
	this := CampaignDiscount{}
	return &this
}

// GetType returns the Type field value
func (o *CampaignDiscount) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *CampaignDiscount) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *CampaignDiscount) SetType(v string) {
	o.Type = v
}

// GetCap returns the Cap field value if set, zero value otherwise.
func (o *CampaignDiscount) GetCap() float64 {
	if o == nil || IsNil(o.Cap) {
		var ret float64
		return ret
	}
	return *o.Cap
}

// GetCapOk returns a tuple with the Cap field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CampaignDiscount) GetCapOk() (*float64, bool) {
	if o == nil || IsNil(o.Cap) {
		return nil, false
	}
	return o.Cap, true
}

// HasCap returns a boolean if a field has been set.
func (o *CampaignDiscount) HasCap() bool {
	if o != nil && !IsNil(o.Cap) {
		return true
	}

	return false
}

// SetCap gets a reference to the given float64 and assigns it to the Cap field.
func (o *CampaignDiscount) SetCap(v float64) {
	o.Cap = &v
}

// GetValue returns the Value field value if set, zero value otherwise.
func (o *CampaignDiscount) GetValue() float64 {
	if o == nil || IsNil(o.Value) {
		var ret float64
		return ret
	}
	return *o.Value
}

// GetValueOk returns a tuple with the Value field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CampaignDiscount) GetValueOk() (*float64, bool) {
	if o == nil || IsNil(o.Value) {
		return nil, false
	}
	return o.Value, true
}

// HasValue returns a boolean if a field has been set.
func (o *CampaignDiscount) HasValue() bool {
	if o != nil && !IsNil(o.Value) {
		return true
	}

	return false
}

// SetValue gets a reference to the given float64 and assigns it to the Value field.
func (o *CampaignDiscount) SetValue(v float64) {
	o.Value = &v
}

// GetScope returns the Scope field value
func (o *CampaignDiscount) GetScope() CampaignScope {
	if o == nil {
		var ret CampaignScope
		return ret
	}

	return o.Scope
}

// GetScopeOk returns a tuple with the Scope field value
// and a boolean to check if the value has been set.
func (o *CampaignDiscount) GetScopeOk() (*CampaignScope, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Scope, true
}

// SetScope sets field value
func (o *CampaignDiscount) SetScope(v CampaignScope) {
	o.Scope = v
}

func (o CampaignDiscount) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CampaignDiscount) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["type"] = o.Type
	if !IsNil(o.Cap) {
		toSerialize["cap"] = o.Cap
	}
	if !IsNil(o.Value) {
		toSerialize["value"] = o.Value
	}
	toSerialize["scope"] = o.Scope

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *CampaignDiscount) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"type",
		"scope",
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

	varCampaignDiscount := _CampaignDiscount{}

	err = json.Unmarshal(data, &varCampaignDiscount)

	if err != nil {
		return err
	}

	*o = CampaignDiscount(varCampaignDiscount)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "type")
		delete(additionalProperties, "cap")
		delete(additionalProperties, "value")
		delete(additionalProperties, "scope")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableCampaignDiscount struct {
	value *CampaignDiscount
	isSet bool
}

func (v NullableCampaignDiscount) Get() *CampaignDiscount {
	return v.value
}

func (v *NullableCampaignDiscount) Set(val *CampaignDiscount) {
	v.value = val
	v.isSet = true
}

func (v NullableCampaignDiscount) IsSet() bool {
	return v.isSet
}

func (v *NullableCampaignDiscount) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCampaignDiscount(val *CampaignDiscount) *NullableCampaignDiscount {
	return &NullableCampaignDiscount{value: val, isSet: true}
}

func (v NullableCampaignDiscount) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCampaignDiscount) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


