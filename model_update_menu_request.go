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
	"gopkg.in/validator.v2"
	"fmt"
)

// UpdateMenuRequest - struct for UpdateMenuRequest
type UpdateMenuRequest struct {
	UpdateMenuItem *UpdateMenuItem
	UpdateMenuModifier *UpdateMenuModifier
}

// UpdateMenuItemAsUpdateMenuRequest is a convenience function that returns UpdateMenuItem wrapped in UpdateMenuRequest
func UpdateMenuItemAsUpdateMenuRequest(v *UpdateMenuItem) UpdateMenuRequest {
	return UpdateMenuRequest{
		UpdateMenuItem: v,
	}
}

// UpdateMenuModifierAsUpdateMenuRequest is a convenience function that returns UpdateMenuModifier wrapped in UpdateMenuRequest
func UpdateMenuModifierAsUpdateMenuRequest(v *UpdateMenuModifier) UpdateMenuRequest {
	return UpdateMenuRequest{
		UpdateMenuModifier: v,
	}
}


// Unmarshal JSON data into one of the pointers in the struct
func (dst *UpdateMenuRequest) UnmarshalJSON(data []byte) error {
	var err error
	match := 0
	// try to unmarshal data into UpdateMenuItem
	err = newStrictDecoder(data).Decode(&dst.UpdateMenuItem)
	if err == nil {
		jsonUpdateMenuItem, _ := json.Marshal(dst.UpdateMenuItem)
		if string(jsonUpdateMenuItem) == "{}" { // empty struct
			dst.UpdateMenuItem = nil
		} else {
			if err = validator.Validate(dst.UpdateMenuItem); err != nil {
				dst.UpdateMenuItem = nil
			} else {
				match++
			}
		}
	} else {
		dst.UpdateMenuItem = nil
	}

	// try to unmarshal data into UpdateMenuModifier
	err = newStrictDecoder(data).Decode(&dst.UpdateMenuModifier)
	if err == nil {
		jsonUpdateMenuModifier, _ := json.Marshal(dst.UpdateMenuModifier)
		if string(jsonUpdateMenuModifier) == "{}" { // empty struct
			dst.UpdateMenuModifier = nil
		} else {
			if err = validator.Validate(dst.UpdateMenuModifier); err != nil {
				dst.UpdateMenuModifier = nil
			} else {
				match++
			}
		}
	} else {
		dst.UpdateMenuModifier = nil
	}

	if match > 1 { // more than 1 match
		// reset to nil
		dst.UpdateMenuItem = nil
		dst.UpdateMenuModifier = nil

		return fmt.Errorf("data matches more than one schema in oneOf(UpdateMenuRequest)")
	} else if match == 1 {
		return nil // exactly one match
	} else { // no match
		return fmt.Errorf("data failed to match schemas in oneOf(UpdateMenuRequest)")
	}
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src UpdateMenuRequest) MarshalJSON() ([]byte, error) {
	if src.UpdateMenuItem != nil {
		return json.Marshal(&src.UpdateMenuItem)
	}

	if src.UpdateMenuModifier != nil {
		return json.Marshal(&src.UpdateMenuModifier)
	}

	return nil, nil // no data in oneOf schemas
}

// Get the actual instance
func (obj *UpdateMenuRequest) GetActualInstance() (interface{}) {
	if obj == nil {
		return nil
	}
	if obj.UpdateMenuItem != nil {
		return obj.UpdateMenuItem
	}

	if obj.UpdateMenuModifier != nil {
		return obj.UpdateMenuModifier
	}

	// all schemas are nil
	return nil
}

type NullableUpdateMenuRequest struct {
	value *UpdateMenuRequest
	isSet bool
}

func (v NullableUpdateMenuRequest) Get() *UpdateMenuRequest {
	return v.value
}

func (v *NullableUpdateMenuRequest) Set(val *UpdateMenuRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableUpdateMenuRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableUpdateMenuRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUpdateMenuRequest(val *UpdateMenuRequest) *NullableUpdateMenuRequest {
	return &NullableUpdateMenuRequest{value: val, isSet: true}
}

func (v NullableUpdateMenuRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUpdateMenuRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


