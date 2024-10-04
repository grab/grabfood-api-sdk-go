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

// checks if the CancelOrderRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CancelOrderRequest{}

// CancelOrderRequest This request cancels an order on GrabFood. 
type CancelOrderRequest struct {
	// The order's ID that is returned from GrabFood. Refer to FAQs for more details about [orderID and shortOrderNumber](#section/Order/What's-the-difference-between-orderID-and-shortOrderNumber).
	OrderID string `json:"orderID"`
	// The merchant's ID that is in GrabFood's database.
	MerchantID string `json:"merchantID"`
	CancelCode CancelCode `json:"cancelCode"`
	AdditionalProperties map[string]interface{}
}

type _CancelOrderRequest CancelOrderRequest

// NewCancelOrderRequest instantiates a new CancelOrderRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCancelOrderRequest(orderID string, merchantID string, cancelCode CancelCode) *CancelOrderRequest {
	this := CancelOrderRequest{}
	this.OrderID = orderID
	this.MerchantID = merchantID
	this.CancelCode = cancelCode
	return &this
}

// NewCancelOrderRequestWithDefaults instantiates a new CancelOrderRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCancelOrderRequestWithDefaults() *CancelOrderRequest {
	this := CancelOrderRequest{}
	return &this
}

// GetOrderID returns the OrderID field value
func (o *CancelOrderRequest) GetOrderID() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.OrderID
}

// GetOrderIDOk returns a tuple with the OrderID field value
// and a boolean to check if the value has been set.
func (o *CancelOrderRequest) GetOrderIDOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.OrderID, true
}

// SetOrderID sets field value
func (o *CancelOrderRequest) SetOrderID(v string) {
	o.OrderID = v
}

// GetMerchantID returns the MerchantID field value
func (o *CancelOrderRequest) GetMerchantID() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.MerchantID
}

// GetMerchantIDOk returns a tuple with the MerchantID field value
// and a boolean to check if the value has been set.
func (o *CancelOrderRequest) GetMerchantIDOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.MerchantID, true
}

// SetMerchantID sets field value
func (o *CancelOrderRequest) SetMerchantID(v string) {
	o.MerchantID = v
}

// GetCancelCode returns the CancelCode field value
func (o *CancelOrderRequest) GetCancelCode() CancelCode {
	if o == nil {
		var ret CancelCode
		return ret
	}

	return o.CancelCode
}

// GetCancelCodeOk returns a tuple with the CancelCode field value
// and a boolean to check if the value has been set.
func (o *CancelOrderRequest) GetCancelCodeOk() (*CancelCode, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CancelCode, true
}

// SetCancelCode sets field value
func (o *CancelOrderRequest) SetCancelCode(v CancelCode) {
	o.CancelCode = v
}

func (o CancelOrderRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CancelOrderRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["orderID"] = o.OrderID
	toSerialize["merchantID"] = o.MerchantID
	toSerialize["cancelCode"] = o.CancelCode

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *CancelOrderRequest) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"orderID",
		"merchantID",
		"cancelCode",
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

	varCancelOrderRequest := _CancelOrderRequest{}

	err = json.Unmarshal(data, &varCancelOrderRequest)

	if err != nil {
		return err
	}

	*o = CancelOrderRequest(varCancelOrderRequest)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "orderID")
		delete(additionalProperties, "merchantID")
		delete(additionalProperties, "cancelCode")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableCancelOrderRequest struct {
	value *CancelOrderRequest
	isSet bool
}

func (v NullableCancelOrderRequest) Get() *CancelOrderRequest {
	return v.value
}

func (v *NullableCancelOrderRequest) Set(val *CancelOrderRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableCancelOrderRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableCancelOrderRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCancelOrderRequest(val *CancelOrderRequest) *NullableCancelOrderRequest {
	return &NullableCancelOrderRequest{value: val, isSet: true}
}

func (v NullableCancelOrderRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCancelOrderRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


