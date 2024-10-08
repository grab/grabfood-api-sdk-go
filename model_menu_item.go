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

// checks if the MenuItem type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &MenuItem{}

// MenuItem struct for MenuItem
type MenuItem struct {
	// The item's ID in the partner system. 
	Id string `json:"id"`
	// The name of the item.
	Name string `json:"name"`
	// Translation of the item name. Only support up to 1 translated language. Refer [Menu Translation](#section/Menu-Translation).
	NameTranslation *map[string]string `json:"nameTranslation,omitempty"`
	// The status for the item that is in the category.  Note: In order to set an item as \"UNAVAILABLE\", it is required to update both the `availableStatus` and `maxStock` fields, whereby the `maxStock` value should be set to 0. 
	AvailableStatus string `json:"availableStatus"`
	// The description of the item. There is a custom length limit of 2000 for `VN`. 
	Description *string `json:"description,omitempty"`
	// Translation of the item description. Only support up to 1 translated language. Refer [Menu Translation](#section/Menu-Translation).
	DescriptionTranslation *map[string]string `json:"descriptionTranslation,omitempty"`
	// The item's price (excluding tax) in minor format. For example: 1900 means $19 with `currency.exponent` as 2. Refer to [FAQ](#section/Menu/Is-the-item-price-with-or-without-tax) for more details. 
	Price int64 `json:"price"`
	// An array string for the item’s image URL links. Refer to FAQs for more details about [images](#section/Menu/What-are-the-recommended-formats-for-an-item-image). 
	Photos []string `json:"photos,omitempty"`
	// The item's special Tag. Refer to FAQs for more details about [specialType](#section/Menu/What's-specialType). 
	SpecialType *string `json:"specialType,omitempty"`
	// **For Indonesia only.** This field allows the configuration for an item to be marked as tax applicable, and marked item would then be included in a commercial invoice to consumers as per the government's regulations. 
	Taxable *bool `json:"taxable,omitempty"`
	// The barcode Number (GTIN). Max 64 allowed. GTIN must be 8, 12, 13, 14 numeric digits. 
	Barcode *string `json:"barcode,omitempty"`
	// The selling time's ID for the item. This value overwrites the category's selling time if it is different. Empty value implies the category's selling time will be applied. 
	SellingTimeID *string `json:"sellingTimeID,omitempty"`
	// Available stocks under inventory for this item. Auto reduce when there is order placed for this item. Empty value implies no limit.  Note: It is necessary to set `maxStock` to 0 if the `availableStatus` of the item is \"UNAVAILABLE\". Item will be set to \"AVAILABLE\" if `maxStock` > 0. 
	MaxStock *int64 `json:"maxStock,omitempty"`
	AdvancedPricing *AdvancedPricing `json:"advancedPricing,omitempty"`
	Purchasability *Purchasability `json:"purchasability,omitempty"`
	// An array of the modifierGroup JSON objects. Max 30 allowed per item. Refer to [Modifier groups](#modifier-groups) for more information.
	ModifierGroups []ModifierGroup `json:"modifierGroups,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _MenuItem MenuItem

// NewMenuItem instantiates a new MenuItem object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMenuItem(id string, name string, availableStatus string, price int64) *MenuItem {
	this := MenuItem{}
	this.Id = id
	this.Name = name
	this.AvailableStatus = availableStatus
	this.Price = price
	return &this
}

// NewMenuItemWithDefaults instantiates a new MenuItem object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMenuItemWithDefaults() *MenuItem {
	this := MenuItem{}
	return &this
}

// GetId returns the Id field value
func (o *MenuItem) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *MenuItem) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *MenuItem) SetId(v string) {
	o.Id = v
}

// GetName returns the Name field value
func (o *MenuItem) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *MenuItem) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *MenuItem) SetName(v string) {
	o.Name = v
}

// GetNameTranslation returns the NameTranslation field value if set, zero value otherwise.
func (o *MenuItem) GetNameTranslation() map[string]string {
	if o == nil || IsNil(o.NameTranslation) {
		var ret map[string]string
		return ret
	}
	return *o.NameTranslation
}

// GetNameTranslationOk returns a tuple with the NameTranslation field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MenuItem) GetNameTranslationOk() (*map[string]string, bool) {
	if o == nil || IsNil(o.NameTranslation) {
		return nil, false
	}
	return o.NameTranslation, true
}

// HasNameTranslation returns a boolean if a field has been set.
func (o *MenuItem) HasNameTranslation() bool {
	if o != nil && !IsNil(o.NameTranslation) {
		return true
	}

	return false
}

// SetNameTranslation gets a reference to the given map[string]string and assigns it to the NameTranslation field.
func (o *MenuItem) SetNameTranslation(v map[string]string) {
	o.NameTranslation = &v
}

// GetAvailableStatus returns the AvailableStatus field value
func (o *MenuItem) GetAvailableStatus() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.AvailableStatus
}

// GetAvailableStatusOk returns a tuple with the AvailableStatus field value
// and a boolean to check if the value has been set.
func (o *MenuItem) GetAvailableStatusOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AvailableStatus, true
}

// SetAvailableStatus sets field value
func (o *MenuItem) SetAvailableStatus(v string) {
	o.AvailableStatus = v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *MenuItem) GetDescription() string {
	if o == nil || IsNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MenuItem) GetDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *MenuItem) HasDescription() bool {
	if o != nil && !IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *MenuItem) SetDescription(v string) {
	o.Description = &v
}

// GetDescriptionTranslation returns the DescriptionTranslation field value if set, zero value otherwise.
func (o *MenuItem) GetDescriptionTranslation() map[string]string {
	if o == nil || IsNil(o.DescriptionTranslation) {
		var ret map[string]string
		return ret
	}
	return *o.DescriptionTranslation
}

// GetDescriptionTranslationOk returns a tuple with the DescriptionTranslation field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MenuItem) GetDescriptionTranslationOk() (*map[string]string, bool) {
	if o == nil || IsNil(o.DescriptionTranslation) {
		return nil, false
	}
	return o.DescriptionTranslation, true
}

// HasDescriptionTranslation returns a boolean if a field has been set.
func (o *MenuItem) HasDescriptionTranslation() bool {
	if o != nil && !IsNil(o.DescriptionTranslation) {
		return true
	}

	return false
}

// SetDescriptionTranslation gets a reference to the given map[string]string and assigns it to the DescriptionTranslation field.
func (o *MenuItem) SetDescriptionTranslation(v map[string]string) {
	o.DescriptionTranslation = &v
}

// GetPrice returns the Price field value
func (o *MenuItem) GetPrice() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.Price
}

// GetPriceOk returns a tuple with the Price field value
// and a boolean to check if the value has been set.
func (o *MenuItem) GetPriceOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Price, true
}

// SetPrice sets field value
func (o *MenuItem) SetPrice(v int64) {
	o.Price = v
}

// GetPhotos returns the Photos field value if set, zero value otherwise.
func (o *MenuItem) GetPhotos() []string {
	if o == nil || IsNil(o.Photos) {
		var ret []string
		return ret
	}
	return o.Photos
}

// GetPhotosOk returns a tuple with the Photos field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MenuItem) GetPhotosOk() ([]string, bool) {
	if o == nil || IsNil(o.Photos) {
		return nil, false
	}
	return o.Photos, true
}

// HasPhotos returns a boolean if a field has been set.
func (o *MenuItem) HasPhotos() bool {
	if o != nil && !IsNil(o.Photos) {
		return true
	}

	return false
}

// SetPhotos gets a reference to the given []string and assigns it to the Photos field.
func (o *MenuItem) SetPhotos(v []string) {
	o.Photos = v
}

// GetSpecialType returns the SpecialType field value if set, zero value otherwise.
func (o *MenuItem) GetSpecialType() string {
	if o == nil || IsNil(o.SpecialType) {
		var ret string
		return ret
	}
	return *o.SpecialType
}

// GetSpecialTypeOk returns a tuple with the SpecialType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MenuItem) GetSpecialTypeOk() (*string, bool) {
	if o == nil || IsNil(o.SpecialType) {
		return nil, false
	}
	return o.SpecialType, true
}

// HasSpecialType returns a boolean if a field has been set.
func (o *MenuItem) HasSpecialType() bool {
	if o != nil && !IsNil(o.SpecialType) {
		return true
	}

	return false
}

// SetSpecialType gets a reference to the given string and assigns it to the SpecialType field.
func (o *MenuItem) SetSpecialType(v string) {
	o.SpecialType = &v
}

// GetTaxable returns the Taxable field value if set, zero value otherwise.
func (o *MenuItem) GetTaxable() bool {
	if o == nil || IsNil(o.Taxable) {
		var ret bool
		return ret
	}
	return *o.Taxable
}

// GetTaxableOk returns a tuple with the Taxable field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MenuItem) GetTaxableOk() (*bool, bool) {
	if o == nil || IsNil(o.Taxable) {
		return nil, false
	}
	return o.Taxable, true
}

// HasTaxable returns a boolean if a field has been set.
func (o *MenuItem) HasTaxable() bool {
	if o != nil && !IsNil(o.Taxable) {
		return true
	}

	return false
}

// SetTaxable gets a reference to the given bool and assigns it to the Taxable field.
func (o *MenuItem) SetTaxable(v bool) {
	o.Taxable = &v
}

// GetBarcode returns the Barcode field value if set, zero value otherwise.
func (o *MenuItem) GetBarcode() string {
	if o == nil || IsNil(o.Barcode) {
		var ret string
		return ret
	}
	return *o.Barcode
}

// GetBarcodeOk returns a tuple with the Barcode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MenuItem) GetBarcodeOk() (*string, bool) {
	if o == nil || IsNil(o.Barcode) {
		return nil, false
	}
	return o.Barcode, true
}

// HasBarcode returns a boolean if a field has been set.
func (o *MenuItem) HasBarcode() bool {
	if o != nil && !IsNil(o.Barcode) {
		return true
	}

	return false
}

// SetBarcode gets a reference to the given string and assigns it to the Barcode field.
func (o *MenuItem) SetBarcode(v string) {
	o.Barcode = &v
}

// GetSellingTimeID returns the SellingTimeID field value if set, zero value otherwise.
func (o *MenuItem) GetSellingTimeID() string {
	if o == nil || IsNil(o.SellingTimeID) {
		var ret string
		return ret
	}
	return *o.SellingTimeID
}

// GetSellingTimeIDOk returns a tuple with the SellingTimeID field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MenuItem) GetSellingTimeIDOk() (*string, bool) {
	if o == nil || IsNil(o.SellingTimeID) {
		return nil, false
	}
	return o.SellingTimeID, true
}

// HasSellingTimeID returns a boolean if a field has been set.
func (o *MenuItem) HasSellingTimeID() bool {
	if o != nil && !IsNil(o.SellingTimeID) {
		return true
	}

	return false
}

// SetSellingTimeID gets a reference to the given string and assigns it to the SellingTimeID field.
func (o *MenuItem) SetSellingTimeID(v string) {
	o.SellingTimeID = &v
}

// GetMaxStock returns the MaxStock field value if set, zero value otherwise.
func (o *MenuItem) GetMaxStock() int64 {
	if o == nil || IsNil(o.MaxStock) {
		var ret int64
		return ret
	}
	return *o.MaxStock
}

// GetMaxStockOk returns a tuple with the MaxStock field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MenuItem) GetMaxStockOk() (*int64, bool) {
	if o == nil || IsNil(o.MaxStock) {
		return nil, false
	}
	return o.MaxStock, true
}

// HasMaxStock returns a boolean if a field has been set.
func (o *MenuItem) HasMaxStock() bool {
	if o != nil && !IsNil(o.MaxStock) {
		return true
	}

	return false
}

// SetMaxStock gets a reference to the given int64 and assigns it to the MaxStock field.
func (o *MenuItem) SetMaxStock(v int64) {
	o.MaxStock = &v
}

// GetAdvancedPricing returns the AdvancedPricing field value if set, zero value otherwise.
func (o *MenuItem) GetAdvancedPricing() AdvancedPricing {
	if o == nil || IsNil(o.AdvancedPricing) {
		var ret AdvancedPricing
		return ret
	}
	return *o.AdvancedPricing
}

// GetAdvancedPricingOk returns a tuple with the AdvancedPricing field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MenuItem) GetAdvancedPricingOk() (*AdvancedPricing, bool) {
	if o == nil || IsNil(o.AdvancedPricing) {
		return nil, false
	}
	return o.AdvancedPricing, true
}

// HasAdvancedPricing returns a boolean if a field has been set.
func (o *MenuItem) HasAdvancedPricing() bool {
	if o != nil && !IsNil(o.AdvancedPricing) {
		return true
	}

	return false
}

// SetAdvancedPricing gets a reference to the given AdvancedPricing and assigns it to the AdvancedPricing field.
func (o *MenuItem) SetAdvancedPricing(v AdvancedPricing) {
	o.AdvancedPricing = &v
}

// GetPurchasability returns the Purchasability field value if set, zero value otherwise.
func (o *MenuItem) GetPurchasability() Purchasability {
	if o == nil || IsNil(o.Purchasability) {
		var ret Purchasability
		return ret
	}
	return *o.Purchasability
}

// GetPurchasabilityOk returns a tuple with the Purchasability field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MenuItem) GetPurchasabilityOk() (*Purchasability, bool) {
	if o == nil || IsNil(o.Purchasability) {
		return nil, false
	}
	return o.Purchasability, true
}

// HasPurchasability returns a boolean if a field has been set.
func (o *MenuItem) HasPurchasability() bool {
	if o != nil && !IsNil(o.Purchasability) {
		return true
	}

	return false
}

// SetPurchasability gets a reference to the given Purchasability and assigns it to the Purchasability field.
func (o *MenuItem) SetPurchasability(v Purchasability) {
	o.Purchasability = &v
}

// GetModifierGroups returns the ModifierGroups field value if set, zero value otherwise.
func (o *MenuItem) GetModifierGroups() []ModifierGroup {
	if o == nil || IsNil(o.ModifierGroups) {
		var ret []ModifierGroup
		return ret
	}
	return o.ModifierGroups
}

// GetModifierGroupsOk returns a tuple with the ModifierGroups field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MenuItem) GetModifierGroupsOk() ([]ModifierGroup, bool) {
	if o == nil || IsNil(o.ModifierGroups) {
		return nil, false
	}
	return o.ModifierGroups, true
}

// HasModifierGroups returns a boolean if a field has been set.
func (o *MenuItem) HasModifierGroups() bool {
	if o != nil && !IsNil(o.ModifierGroups) {
		return true
	}

	return false
}

// SetModifierGroups gets a reference to the given []ModifierGroup and assigns it to the ModifierGroups field.
func (o *MenuItem) SetModifierGroups(v []ModifierGroup) {
	o.ModifierGroups = v
}

func (o MenuItem) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o MenuItem) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["id"] = o.Id
	toSerialize["name"] = o.Name
	if !IsNil(o.NameTranslation) {
		toSerialize["nameTranslation"] = o.NameTranslation
	}
	toSerialize["availableStatus"] = o.AvailableStatus
	if !IsNil(o.Description) {
		toSerialize["description"] = o.Description
	}
	if !IsNil(o.DescriptionTranslation) {
		toSerialize["descriptionTranslation"] = o.DescriptionTranslation
	}
	toSerialize["price"] = o.Price
	if !IsNil(o.Photos) {
		toSerialize["photos"] = o.Photos
	}
	if !IsNil(o.SpecialType) {
		toSerialize["specialType"] = o.SpecialType
	}
	if !IsNil(o.Taxable) {
		toSerialize["taxable"] = o.Taxable
	}
	if !IsNil(o.Barcode) {
		toSerialize["barcode"] = o.Barcode
	}
	if !IsNil(o.SellingTimeID) {
		toSerialize["sellingTimeID"] = o.SellingTimeID
	}
	if !IsNil(o.MaxStock) {
		toSerialize["maxStock"] = o.MaxStock
	}
	if !IsNil(o.AdvancedPricing) {
		toSerialize["advancedPricing"] = o.AdvancedPricing
	}
	if !IsNil(o.Purchasability) {
		toSerialize["purchasability"] = o.Purchasability
	}
	if !IsNil(o.ModifierGroups) {
		toSerialize["modifierGroups"] = o.ModifierGroups
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *MenuItem) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"id",
		"name",
		"availableStatus",
		"price",
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

	varMenuItem := _MenuItem{}

	err = json.Unmarshal(data, &varMenuItem)

	if err != nil {
		return err
	}

	*o = MenuItem(varMenuItem)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "id")
		delete(additionalProperties, "name")
		delete(additionalProperties, "nameTranslation")
		delete(additionalProperties, "availableStatus")
		delete(additionalProperties, "description")
		delete(additionalProperties, "descriptionTranslation")
		delete(additionalProperties, "price")
		delete(additionalProperties, "photos")
		delete(additionalProperties, "specialType")
		delete(additionalProperties, "taxable")
		delete(additionalProperties, "barcode")
		delete(additionalProperties, "sellingTimeID")
		delete(additionalProperties, "maxStock")
		delete(additionalProperties, "advancedPricing")
		delete(additionalProperties, "purchasability")
		delete(additionalProperties, "modifierGroups")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableMenuItem struct {
	value *MenuItem
	isSet bool
}

func (v NullableMenuItem) Get() *MenuItem {
	return v.value
}

func (v *NullableMenuItem) Set(val *MenuItem) {
	v.value = val
	v.isSet = true
}

func (v NullableMenuItem) IsSet() bool {
	return v.isSet
}

func (v *NullableMenuItem) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMenuItem(val *MenuItem) *NullableMenuItem {
	return &NullableMenuItem{value: val, isSet: true}
}

func (v NullableMenuItem) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMenuItem) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


