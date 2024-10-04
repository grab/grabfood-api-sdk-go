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

// checks if the MenuModifier type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &MenuModifier{}

// MenuModifier struct for MenuModifier
type MenuModifier struct {
	// The modifier's ID that is on the partner's system. This ID should be unique with a min length of 1 and max of 64.
	Id string `json:"id"`
	// The name of the modifier.
	Name string `json:"name"`
	// Translation of the modifier name. Only support up to 1 translated language. Refer [Menu Translation](#section/Menu-Translation).
	NameTranslation *map[string]string `json:"nameTranslation,omitempty"`
	// The status for the modifier that is in the ModifierGroup.
	AvailableStatus string `json:"availableStatus"`
	// The modifier's price (excluding tax) in minor format. Refer to FAQs for more details about [tax](#section/Order/How-is-tax-calculated).
	Price *int64 `json:"price,omitempty"`
	// The barcode Number (GTIN). GTIN must be 8, 12, 13, 14 numeric digits.
	Barcode *string `json:"barcode,omitempty"`
	AdvancedPricing *AdvancedPricing `json:"advancedPricing,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _MenuModifier MenuModifier

// NewMenuModifier instantiates a new MenuModifier object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMenuModifier(id string, name string, availableStatus string) *MenuModifier {
	this := MenuModifier{}
	this.Id = id
	this.Name = name
	this.AvailableStatus = availableStatus
	return &this
}

// NewMenuModifierWithDefaults instantiates a new MenuModifier object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMenuModifierWithDefaults() *MenuModifier {
	this := MenuModifier{}
	return &this
}

// GetId returns the Id field value
func (o *MenuModifier) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *MenuModifier) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *MenuModifier) SetId(v string) {
	o.Id = v
}

// GetName returns the Name field value
func (o *MenuModifier) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *MenuModifier) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *MenuModifier) SetName(v string) {
	o.Name = v
}

// GetNameTranslation returns the NameTranslation field value if set, zero value otherwise.
func (o *MenuModifier) GetNameTranslation() map[string]string {
	if o == nil || IsNil(o.NameTranslation) {
		var ret map[string]string
		return ret
	}
	return *o.NameTranslation
}

// GetNameTranslationOk returns a tuple with the NameTranslation field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MenuModifier) GetNameTranslationOk() (*map[string]string, bool) {
	if o == nil || IsNil(o.NameTranslation) {
		return nil, false
	}
	return o.NameTranslation, true
}

// HasNameTranslation returns a boolean if a field has been set.
func (o *MenuModifier) HasNameTranslation() bool {
	if o != nil && !IsNil(o.NameTranslation) {
		return true
	}

	return false
}

// SetNameTranslation gets a reference to the given map[string]string and assigns it to the NameTranslation field.
func (o *MenuModifier) SetNameTranslation(v map[string]string) {
	o.NameTranslation = &v
}

// GetAvailableStatus returns the AvailableStatus field value
func (o *MenuModifier) GetAvailableStatus() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.AvailableStatus
}

// GetAvailableStatusOk returns a tuple with the AvailableStatus field value
// and a boolean to check if the value has been set.
func (o *MenuModifier) GetAvailableStatusOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AvailableStatus, true
}

// SetAvailableStatus sets field value
func (o *MenuModifier) SetAvailableStatus(v string) {
	o.AvailableStatus = v
}

// GetPrice returns the Price field value if set, zero value otherwise.
func (o *MenuModifier) GetPrice() int64 {
	if o == nil || IsNil(o.Price) {
		var ret int64
		return ret
	}
	return *o.Price
}

// GetPriceOk returns a tuple with the Price field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MenuModifier) GetPriceOk() (*int64, bool) {
	if o == nil || IsNil(o.Price) {
		return nil, false
	}
	return o.Price, true
}

// HasPrice returns a boolean if a field has been set.
func (o *MenuModifier) HasPrice() bool {
	if o != nil && !IsNil(o.Price) {
		return true
	}

	return false
}

// SetPrice gets a reference to the given int64 and assigns it to the Price field.
func (o *MenuModifier) SetPrice(v int64) {
	o.Price = &v
}

// GetBarcode returns the Barcode field value if set, zero value otherwise.
func (o *MenuModifier) GetBarcode() string {
	if o == nil || IsNil(o.Barcode) {
		var ret string
		return ret
	}
	return *o.Barcode
}

// GetBarcodeOk returns a tuple with the Barcode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MenuModifier) GetBarcodeOk() (*string, bool) {
	if o == nil || IsNil(o.Barcode) {
		return nil, false
	}
	return o.Barcode, true
}

// HasBarcode returns a boolean if a field has been set.
func (o *MenuModifier) HasBarcode() bool {
	if o != nil && !IsNil(o.Barcode) {
		return true
	}

	return false
}

// SetBarcode gets a reference to the given string and assigns it to the Barcode field.
func (o *MenuModifier) SetBarcode(v string) {
	o.Barcode = &v
}

// GetAdvancedPricing returns the AdvancedPricing field value if set, zero value otherwise.
func (o *MenuModifier) GetAdvancedPricing() AdvancedPricing {
	if o == nil || IsNil(o.AdvancedPricing) {
		var ret AdvancedPricing
		return ret
	}
	return *o.AdvancedPricing
}

// GetAdvancedPricingOk returns a tuple with the AdvancedPricing field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MenuModifier) GetAdvancedPricingOk() (*AdvancedPricing, bool) {
	if o == nil || IsNil(o.AdvancedPricing) {
		return nil, false
	}
	return o.AdvancedPricing, true
}

// HasAdvancedPricing returns a boolean if a field has been set.
func (o *MenuModifier) HasAdvancedPricing() bool {
	if o != nil && !IsNil(o.AdvancedPricing) {
		return true
	}

	return false
}

// SetAdvancedPricing gets a reference to the given AdvancedPricing and assigns it to the AdvancedPricing field.
func (o *MenuModifier) SetAdvancedPricing(v AdvancedPricing) {
	o.AdvancedPricing = &v
}

func (o MenuModifier) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o MenuModifier) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["id"] = o.Id
	toSerialize["name"] = o.Name
	if !IsNil(o.NameTranslation) {
		toSerialize["nameTranslation"] = o.NameTranslation
	}
	toSerialize["availableStatus"] = o.AvailableStatus
	if !IsNil(o.Price) {
		toSerialize["price"] = o.Price
	}
	if !IsNil(o.Barcode) {
		toSerialize["barcode"] = o.Barcode
	}
	if !IsNil(o.AdvancedPricing) {
		toSerialize["advancedPricing"] = o.AdvancedPricing
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *MenuModifier) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"id",
		"name",
		"availableStatus",
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

	varMenuModifier := _MenuModifier{}

	err = json.Unmarshal(data, &varMenuModifier)

	if err != nil {
		return err
	}

	*o = MenuModifier(varMenuModifier)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "id")
		delete(additionalProperties, "name")
		delete(additionalProperties, "nameTranslation")
		delete(additionalProperties, "availableStatus")
		delete(additionalProperties, "price")
		delete(additionalProperties, "barcode")
		delete(additionalProperties, "advancedPricing")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableMenuModifier struct {
	value *MenuModifier
	isSet bool
}

func (v NullableMenuModifier) Get() *MenuModifier {
	return v.value
}

func (v *NullableMenuModifier) Set(val *MenuModifier) {
	v.value = val
	v.isSet = true
}

func (v NullableMenuModifier) IsSet() bool {
	return v.isSet
}

func (v *NullableMenuModifier) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMenuModifier(val *MenuModifier) *NullableMenuModifier {
	return &NullableMenuModifier{value: val, isSet: true}
}

func (v NullableMenuModifier) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMenuModifier) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


