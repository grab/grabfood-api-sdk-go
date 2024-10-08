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

// checks if the Voucher type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &Voucher{}

// Voucher A JSON object containing dine-in voucher details.
type Voucher struct {
	// The title of the voucher.
	Title *string `json:"title,omitempty"`
	// The amount paid after discount is applied in local currency.
	DiscountedPrice *string `json:"discountedPrice,omitempty"`
	// The original amount before discount is applied in local currency.
	OriginalPrice *string `json:"originalPrice,omitempty"`
	DescriptionInfo *VoucherDescriptionInfo `json:"descriptionInfo,omitempty"`
	// The type of the dine-in voucher.
	VoucherType *string `json:"voucherType,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _Voucher Voucher

// NewVoucher instantiates a new Voucher object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewVoucher() *Voucher {
	this := Voucher{}
	return &this
}

// NewVoucherWithDefaults instantiates a new Voucher object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewVoucherWithDefaults() *Voucher {
	this := Voucher{}
	return &this
}

// GetTitle returns the Title field value if set, zero value otherwise.
func (o *Voucher) GetTitle() string {
	if o == nil || IsNil(o.Title) {
		var ret string
		return ret
	}
	return *o.Title
}

// GetTitleOk returns a tuple with the Title field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Voucher) GetTitleOk() (*string, bool) {
	if o == nil || IsNil(o.Title) {
		return nil, false
	}
	return o.Title, true
}

// HasTitle returns a boolean if a field has been set.
func (o *Voucher) HasTitle() bool {
	if o != nil && !IsNil(o.Title) {
		return true
	}

	return false
}

// SetTitle gets a reference to the given string and assigns it to the Title field.
func (o *Voucher) SetTitle(v string) {
	o.Title = &v
}

// GetDiscountedPrice returns the DiscountedPrice field value if set, zero value otherwise.
func (o *Voucher) GetDiscountedPrice() string {
	if o == nil || IsNil(o.DiscountedPrice) {
		var ret string
		return ret
	}
	return *o.DiscountedPrice
}

// GetDiscountedPriceOk returns a tuple with the DiscountedPrice field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Voucher) GetDiscountedPriceOk() (*string, bool) {
	if o == nil || IsNil(o.DiscountedPrice) {
		return nil, false
	}
	return o.DiscountedPrice, true
}

// HasDiscountedPrice returns a boolean if a field has been set.
func (o *Voucher) HasDiscountedPrice() bool {
	if o != nil && !IsNil(o.DiscountedPrice) {
		return true
	}

	return false
}

// SetDiscountedPrice gets a reference to the given string and assigns it to the DiscountedPrice field.
func (o *Voucher) SetDiscountedPrice(v string) {
	o.DiscountedPrice = &v
}

// GetOriginalPrice returns the OriginalPrice field value if set, zero value otherwise.
func (o *Voucher) GetOriginalPrice() string {
	if o == nil || IsNil(o.OriginalPrice) {
		var ret string
		return ret
	}
	return *o.OriginalPrice
}

// GetOriginalPriceOk returns a tuple with the OriginalPrice field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Voucher) GetOriginalPriceOk() (*string, bool) {
	if o == nil || IsNil(o.OriginalPrice) {
		return nil, false
	}
	return o.OriginalPrice, true
}

// HasOriginalPrice returns a boolean if a field has been set.
func (o *Voucher) HasOriginalPrice() bool {
	if o != nil && !IsNil(o.OriginalPrice) {
		return true
	}

	return false
}

// SetOriginalPrice gets a reference to the given string and assigns it to the OriginalPrice field.
func (o *Voucher) SetOriginalPrice(v string) {
	o.OriginalPrice = &v
}

// GetDescriptionInfo returns the DescriptionInfo field value if set, zero value otherwise.
func (o *Voucher) GetDescriptionInfo() VoucherDescriptionInfo {
	if o == nil || IsNil(o.DescriptionInfo) {
		var ret VoucherDescriptionInfo
		return ret
	}
	return *o.DescriptionInfo
}

// GetDescriptionInfoOk returns a tuple with the DescriptionInfo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Voucher) GetDescriptionInfoOk() (*VoucherDescriptionInfo, bool) {
	if o == nil || IsNil(o.DescriptionInfo) {
		return nil, false
	}
	return o.DescriptionInfo, true
}

// HasDescriptionInfo returns a boolean if a field has been set.
func (o *Voucher) HasDescriptionInfo() bool {
	if o != nil && !IsNil(o.DescriptionInfo) {
		return true
	}

	return false
}

// SetDescriptionInfo gets a reference to the given VoucherDescriptionInfo and assigns it to the DescriptionInfo field.
func (o *Voucher) SetDescriptionInfo(v VoucherDescriptionInfo) {
	o.DescriptionInfo = &v
}

// GetVoucherType returns the VoucherType field value if set, zero value otherwise.
func (o *Voucher) GetVoucherType() string {
	if o == nil || IsNil(o.VoucherType) {
		var ret string
		return ret
	}
	return *o.VoucherType
}

// GetVoucherTypeOk returns a tuple with the VoucherType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Voucher) GetVoucherTypeOk() (*string, bool) {
	if o == nil || IsNil(o.VoucherType) {
		return nil, false
	}
	return o.VoucherType, true
}

// HasVoucherType returns a boolean if a field has been set.
func (o *Voucher) HasVoucherType() bool {
	if o != nil && !IsNil(o.VoucherType) {
		return true
	}

	return false
}

// SetVoucherType gets a reference to the given string and assigns it to the VoucherType field.
func (o *Voucher) SetVoucherType(v string) {
	o.VoucherType = &v
}

func (o Voucher) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o Voucher) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Title) {
		toSerialize["title"] = o.Title
	}
	if !IsNil(o.DiscountedPrice) {
		toSerialize["discountedPrice"] = o.DiscountedPrice
	}
	if !IsNil(o.OriginalPrice) {
		toSerialize["originalPrice"] = o.OriginalPrice
	}
	if !IsNil(o.DescriptionInfo) {
		toSerialize["descriptionInfo"] = o.DescriptionInfo
	}
	if !IsNil(o.VoucherType) {
		toSerialize["voucherType"] = o.VoucherType
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *Voucher) UnmarshalJSON(data []byte) (err error) {
	varVoucher := _Voucher{}

	err = json.Unmarshal(data, &varVoucher)

	if err != nil {
		return err
	}

	*o = Voucher(varVoucher)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "title")
		delete(additionalProperties, "discountedPrice")
		delete(additionalProperties, "originalPrice")
		delete(additionalProperties, "descriptionInfo")
		delete(additionalProperties, "voucherType")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableVoucher struct {
	value *Voucher
	isSet bool
}

func (v NullableVoucher) Get() *Voucher {
	return v.value
}

func (v *NullableVoucher) Set(val *Voucher) {
	v.value = val
	v.isSet = true
}

func (v NullableVoucher) IsSet() bool {
	return v.isSet
}

func (v *NullableVoucher) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableVoucher(val *Voucher) *NullableVoucher {
	return &NullableVoucher{value: val, isSet: true}
}

func (v NullableVoucher) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableVoucher) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


