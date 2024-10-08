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

// checks if the BatchUpdateMenuResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &BatchUpdateMenuResponse{}

// BatchUpdateMenuResponse struct for BatchUpdateMenuResponse
type BatchUpdateMenuResponse struct {
	// The merchant's ID that is in GrabFood's database.
	MerchantID *string `json:"merchantID,omitempty"`
	// The status of this request.
	Status *string `json:"status,omitempty"`
	// The error messages when batch update menu record was partial success/fail. `null` when the request was success.
	Errors []MenuEntityError `json:"errors,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _BatchUpdateMenuResponse BatchUpdateMenuResponse

// NewBatchUpdateMenuResponse instantiates a new BatchUpdateMenuResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBatchUpdateMenuResponse() *BatchUpdateMenuResponse {
	this := BatchUpdateMenuResponse{}
	return &this
}

// NewBatchUpdateMenuResponseWithDefaults instantiates a new BatchUpdateMenuResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBatchUpdateMenuResponseWithDefaults() *BatchUpdateMenuResponse {
	this := BatchUpdateMenuResponse{}
	return &this
}

// GetMerchantID returns the MerchantID field value if set, zero value otherwise.
func (o *BatchUpdateMenuResponse) GetMerchantID() string {
	if o == nil || IsNil(o.MerchantID) {
		var ret string
		return ret
	}
	return *o.MerchantID
}

// GetMerchantIDOk returns a tuple with the MerchantID field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BatchUpdateMenuResponse) GetMerchantIDOk() (*string, bool) {
	if o == nil || IsNil(o.MerchantID) {
		return nil, false
	}
	return o.MerchantID, true
}

// HasMerchantID returns a boolean if a field has been set.
func (o *BatchUpdateMenuResponse) HasMerchantID() bool {
	if o != nil && !IsNil(o.MerchantID) {
		return true
	}

	return false
}

// SetMerchantID gets a reference to the given string and assigns it to the MerchantID field.
func (o *BatchUpdateMenuResponse) SetMerchantID(v string) {
	o.MerchantID = &v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *BatchUpdateMenuResponse) GetStatus() string {
	if o == nil || IsNil(o.Status) {
		var ret string
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BatchUpdateMenuResponse) GetStatusOk() (*string, bool) {
	if o == nil || IsNil(o.Status) {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *BatchUpdateMenuResponse) HasStatus() bool {
	if o != nil && !IsNil(o.Status) {
		return true
	}

	return false
}

// SetStatus gets a reference to the given string and assigns it to the Status field.
func (o *BatchUpdateMenuResponse) SetStatus(v string) {
	o.Status = &v
}

// GetErrors returns the Errors field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *BatchUpdateMenuResponse) GetErrors() []MenuEntityError {
	if o == nil {
		var ret []MenuEntityError
		return ret
	}
	return o.Errors
}

// GetErrorsOk returns a tuple with the Errors field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *BatchUpdateMenuResponse) GetErrorsOk() ([]MenuEntityError, bool) {
	if o == nil || IsNil(o.Errors) {
		return nil, false
	}
	return o.Errors, true
}

// HasErrors returns a boolean if a field has been set.
func (o *BatchUpdateMenuResponse) HasErrors() bool {
	if o != nil && !IsNil(o.Errors) {
		return true
	}

	return false
}

// SetErrors gets a reference to the given []MenuEntityError and assigns it to the Errors field.
func (o *BatchUpdateMenuResponse) SetErrors(v []MenuEntityError) {
	o.Errors = v
}

func (o BatchUpdateMenuResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o BatchUpdateMenuResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.MerchantID) {
		toSerialize["merchantID"] = o.MerchantID
	}
	if !IsNil(o.Status) {
		toSerialize["status"] = o.Status
	}
	if o.Errors != nil {
		toSerialize["errors"] = o.Errors
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *BatchUpdateMenuResponse) UnmarshalJSON(data []byte) (err error) {
	varBatchUpdateMenuResponse := _BatchUpdateMenuResponse{}

	err = json.Unmarshal(data, &varBatchUpdateMenuResponse)

	if err != nil {
		return err
	}

	*o = BatchUpdateMenuResponse(varBatchUpdateMenuResponse)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "merchantID")
		delete(additionalProperties, "status")
		delete(additionalProperties, "errors")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableBatchUpdateMenuResponse struct {
	value *BatchUpdateMenuResponse
	isSet bool
}

func (v NullableBatchUpdateMenuResponse) Get() *BatchUpdateMenuResponse {
	return v.value
}

func (v *NullableBatchUpdateMenuResponse) Set(val *BatchUpdateMenuResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableBatchUpdateMenuResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableBatchUpdateMenuResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBatchUpdateMenuResponse(val *BatchUpdateMenuResponse) *NullableBatchUpdateMenuResponse {
	return &NullableBatchUpdateMenuResponse{value: val, isSet: true}
}

func (v NullableBatchUpdateMenuResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBatchUpdateMenuResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


