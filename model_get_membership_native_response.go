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

// checks if the GetMembershipNativeResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetMembershipNativeResponse{}

// GetMembershipNativeResponse This response returns membership detail. 
type GetMembershipNativeResponse struct {
	// Status of the memberID.
	MembershipStatus *string `json:"membershipStatus,omitempty"`
	PointInfo *GetMembershipNativeResponsePointInfo `json:"pointInfo,omitempty"`
	// Earliest points expiry date. In `yyyy-mm-dd` format
	PointsExpireDate *string `json:"pointsExpireDate,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _GetMembershipNativeResponse GetMembershipNativeResponse

// NewGetMembershipNativeResponse instantiates a new GetMembershipNativeResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetMembershipNativeResponse() *GetMembershipNativeResponse {
	this := GetMembershipNativeResponse{}
	return &this
}

// NewGetMembershipNativeResponseWithDefaults instantiates a new GetMembershipNativeResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetMembershipNativeResponseWithDefaults() *GetMembershipNativeResponse {
	this := GetMembershipNativeResponse{}
	return &this
}

// GetMembershipStatus returns the MembershipStatus field value if set, zero value otherwise.
func (o *GetMembershipNativeResponse) GetMembershipStatus() string {
	if o == nil || IsNil(o.MembershipStatus) {
		var ret string
		return ret
	}
	return *o.MembershipStatus
}

// GetMembershipStatusOk returns a tuple with the MembershipStatus field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetMembershipNativeResponse) GetMembershipStatusOk() (*string, bool) {
	if o == nil || IsNil(o.MembershipStatus) {
		return nil, false
	}
	return o.MembershipStatus, true
}

// HasMembershipStatus returns a boolean if a field has been set.
func (o *GetMembershipNativeResponse) HasMembershipStatus() bool {
	if o != nil && !IsNil(o.MembershipStatus) {
		return true
	}

	return false
}

// SetMembershipStatus gets a reference to the given string and assigns it to the MembershipStatus field.
func (o *GetMembershipNativeResponse) SetMembershipStatus(v string) {
	o.MembershipStatus = &v
}

// GetPointInfo returns the PointInfo field value if set, zero value otherwise.
func (o *GetMembershipNativeResponse) GetPointInfo() GetMembershipNativeResponsePointInfo {
	if o == nil || IsNil(o.PointInfo) {
		var ret GetMembershipNativeResponsePointInfo
		return ret
	}
	return *o.PointInfo
}

// GetPointInfoOk returns a tuple with the PointInfo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetMembershipNativeResponse) GetPointInfoOk() (*GetMembershipNativeResponsePointInfo, bool) {
	if o == nil || IsNil(o.PointInfo) {
		return nil, false
	}
	return o.PointInfo, true
}

// HasPointInfo returns a boolean if a field has been set.
func (o *GetMembershipNativeResponse) HasPointInfo() bool {
	if o != nil && !IsNil(o.PointInfo) {
		return true
	}

	return false
}

// SetPointInfo gets a reference to the given GetMembershipNativeResponsePointInfo and assigns it to the PointInfo field.
func (o *GetMembershipNativeResponse) SetPointInfo(v GetMembershipNativeResponsePointInfo) {
	o.PointInfo = &v
}

// GetPointsExpireDate returns the PointsExpireDate field value if set, zero value otherwise.
func (o *GetMembershipNativeResponse) GetPointsExpireDate() string {
	if o == nil || IsNil(o.PointsExpireDate) {
		var ret string
		return ret
	}
	return *o.PointsExpireDate
}

// GetPointsExpireDateOk returns a tuple with the PointsExpireDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetMembershipNativeResponse) GetPointsExpireDateOk() (*string, bool) {
	if o == nil || IsNil(o.PointsExpireDate) {
		return nil, false
	}
	return o.PointsExpireDate, true
}

// HasPointsExpireDate returns a boolean if a field has been set.
func (o *GetMembershipNativeResponse) HasPointsExpireDate() bool {
	if o != nil && !IsNil(o.PointsExpireDate) {
		return true
	}

	return false
}

// SetPointsExpireDate gets a reference to the given string and assigns it to the PointsExpireDate field.
func (o *GetMembershipNativeResponse) SetPointsExpireDate(v string) {
	o.PointsExpireDate = &v
}

func (o GetMembershipNativeResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetMembershipNativeResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.MembershipStatus) {
		toSerialize["membershipStatus"] = o.MembershipStatus
	}
	if !IsNil(o.PointInfo) {
		toSerialize["pointInfo"] = o.PointInfo
	}
	if !IsNil(o.PointsExpireDate) {
		toSerialize["pointsExpireDate"] = o.PointsExpireDate
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *GetMembershipNativeResponse) UnmarshalJSON(data []byte) (err error) {
	varGetMembershipNativeResponse := _GetMembershipNativeResponse{}

	err = json.Unmarshal(data, &varGetMembershipNativeResponse)

	if err != nil {
		return err
	}

	*o = GetMembershipNativeResponse(varGetMembershipNativeResponse)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "membershipStatus")
		delete(additionalProperties, "pointInfo")
		delete(additionalProperties, "pointsExpireDate")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableGetMembershipNativeResponse struct {
	value *GetMembershipNativeResponse
	isSet bool
}

func (v NullableGetMembershipNativeResponse) Get() *GetMembershipNativeResponse {
	return v.value
}

func (v *NullableGetMembershipNativeResponse) Set(val *GetMembershipNativeResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableGetMembershipNativeResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableGetMembershipNativeResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetMembershipNativeResponse(val *GetMembershipNativeResponse) *NullableGetMembershipNativeResponse {
	return &NullableGetMembershipNativeResponse{value: val, isSet: true}
}

func (v NullableGetMembershipNativeResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetMembershipNativeResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


