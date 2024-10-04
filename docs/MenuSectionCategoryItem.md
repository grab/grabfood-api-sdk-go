# MenuSectionCategoryItem

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | The item&#39;s ID in the partner system.  | 
**Name** | **string** | The name of the item. | 
**NameTranslation** | Pointer to **map[string]string** | Translation of the item name. Only support up to 1 translated language. Refer [Menu Translation](#section/Menu-Translation). | [optional] 
**AvailableStatus** | **string** | The status for the item that is in the category. Refer to FAQs for more details about [availableStatus](#section/Menu/What-is-availableStatus).  Note: In order to set an item as \&quot;UNAVAILABLE\&quot;, it is required to update both the &#x60;availableStatus&#x60; and &#x60;maxStock&#x60; fields, whereby the &#x60;maxStock&#x60; value should be set to 0.  | 
**Description** | Pointer to **string** | The description of the item. There is a custom length limit of 2000 for &#x60;VN&#x60;.  | [optional] 
**DescriptionTranslation** | Pointer to **map[string]string** | Translation of the item description. Only support up to 1 translated language. Refer [Menu Translation](#section/Menu-Translation). | [optional] 
**Price** | **int32** | The item&#39;s price (excluding tax) in minor format. For example: 1900 means $19 with &#x60;currency.exponent&#x60; as 2. Refer to [FAQ](#section/Menu/Is-the-item-price-with-or-without-tax) for more details.  | 
**Photos** | Pointer to **[]string** | An array string for the item’s image URL links. Refer to FAQs for more details about [images](#section/Menu/What-are-the-recommended-formats-for-an-item-image).  | [optional] 
**SpecialType** | Pointer to **string** | The item&#39;s special Tag. Refer to FAQs for more details about [specialType](#section/Menu/What&#39;s-specialType).  | [optional] 
**Taxable** | Pointer to **bool** | **For Indonesia only.** This field allows the configuration for an item to be marked as tax applicable, and marked item would then be included in a commercial invoice to consumers as per the government&#39;s regulations.  | [optional] 
**Barcode** | Pointer to **string** | The barcode Number (GTIN). Max 64 allowed. GTIN must be 8, 12, 13, 14 numeric digits.  | [optional] 
**MaxStock** | Pointer to **int32** | Available stocks under inventory for this item. Auto reduce when there is order placed for this item. Empty value implies no limit.  Note: It is necessary to set &#x60;maxStock&#x60; to 0 if the &#x60;availableStatus&#x60; of the item is \&quot;UNAVAILABLE\&quot;. Item will be set to \&quot;AVAILABLE\&quot; if &#x60;maxStock&#x60; &gt; 0.  | [optional] 
**AdvancedPricing** | Pointer to [**AdvancedPricing**](AdvancedPricing.md) |  | [optional] 
**Purchasability** | Pointer to [**Purchasability**](Purchasability.md) |  | [optional] 
**ModifierGroups** | Pointer to [**[]ModifierGroup**](ModifierGroup.md) | An array of the modifierGroup JSON objects. Max 30 allowed per item. Refer to [Modifier groups](#modifier-groups) for more information. | [optional] 

## Methods

### NewMenuSectionCategoryItem

`func NewMenuSectionCategoryItem(id string, name string, availableStatus string, price int32, ) *MenuSectionCategoryItem`

NewMenuSectionCategoryItem instantiates a new MenuSectionCategoryItem object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMenuSectionCategoryItemWithDefaults

`func NewMenuSectionCategoryItemWithDefaults() *MenuSectionCategoryItem`

NewMenuSectionCategoryItemWithDefaults instantiates a new MenuSectionCategoryItem object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *MenuSectionCategoryItem) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *MenuSectionCategoryItem) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *MenuSectionCategoryItem) SetId(v string)`

SetId sets Id field to given value.


### GetName

`func (o *MenuSectionCategoryItem) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *MenuSectionCategoryItem) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *MenuSectionCategoryItem) SetName(v string)`

SetName sets Name field to given value.


### GetNameTranslation

`func (o *MenuSectionCategoryItem) GetNameTranslation() map[string]string`

GetNameTranslation returns the NameTranslation field if non-nil, zero value otherwise.

### GetNameTranslationOk

`func (o *MenuSectionCategoryItem) GetNameTranslationOk() (*map[string]string, bool)`

GetNameTranslationOk returns a tuple with the NameTranslation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNameTranslation

`func (o *MenuSectionCategoryItem) SetNameTranslation(v map[string]string)`

SetNameTranslation sets NameTranslation field to given value.

### HasNameTranslation

`func (o *MenuSectionCategoryItem) HasNameTranslation() bool`

HasNameTranslation returns a boolean if a field has been set.

### GetAvailableStatus

`func (o *MenuSectionCategoryItem) GetAvailableStatus() string`

GetAvailableStatus returns the AvailableStatus field if non-nil, zero value otherwise.

### GetAvailableStatusOk

`func (o *MenuSectionCategoryItem) GetAvailableStatusOk() (*string, bool)`

GetAvailableStatusOk returns a tuple with the AvailableStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvailableStatus

`func (o *MenuSectionCategoryItem) SetAvailableStatus(v string)`

SetAvailableStatus sets AvailableStatus field to given value.


### GetDescription

`func (o *MenuSectionCategoryItem) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *MenuSectionCategoryItem) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *MenuSectionCategoryItem) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *MenuSectionCategoryItem) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetDescriptionTranslation

`func (o *MenuSectionCategoryItem) GetDescriptionTranslation() map[string]string`

GetDescriptionTranslation returns the DescriptionTranslation field if non-nil, zero value otherwise.

### GetDescriptionTranslationOk

`func (o *MenuSectionCategoryItem) GetDescriptionTranslationOk() (*map[string]string, bool)`

GetDescriptionTranslationOk returns a tuple with the DescriptionTranslation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescriptionTranslation

`func (o *MenuSectionCategoryItem) SetDescriptionTranslation(v map[string]string)`

SetDescriptionTranslation sets DescriptionTranslation field to given value.

### HasDescriptionTranslation

`func (o *MenuSectionCategoryItem) HasDescriptionTranslation() bool`

HasDescriptionTranslation returns a boolean if a field has been set.

### GetPrice

`func (o *MenuSectionCategoryItem) GetPrice() int32`

GetPrice returns the Price field if non-nil, zero value otherwise.

### GetPriceOk

`func (o *MenuSectionCategoryItem) GetPriceOk() (*int32, bool)`

GetPriceOk returns a tuple with the Price field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrice

`func (o *MenuSectionCategoryItem) SetPrice(v int32)`

SetPrice sets Price field to given value.


### GetPhotos

`func (o *MenuSectionCategoryItem) GetPhotos() []string`

GetPhotos returns the Photos field if non-nil, zero value otherwise.

### GetPhotosOk

`func (o *MenuSectionCategoryItem) GetPhotosOk() (*[]string, bool)`

GetPhotosOk returns a tuple with the Photos field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPhotos

`func (o *MenuSectionCategoryItem) SetPhotos(v []string)`

SetPhotos sets Photos field to given value.

### HasPhotos

`func (o *MenuSectionCategoryItem) HasPhotos() bool`

HasPhotos returns a boolean if a field has been set.

### GetSpecialType

`func (o *MenuSectionCategoryItem) GetSpecialType() string`

GetSpecialType returns the SpecialType field if non-nil, zero value otherwise.

### GetSpecialTypeOk

`func (o *MenuSectionCategoryItem) GetSpecialTypeOk() (*string, bool)`

GetSpecialTypeOk returns a tuple with the SpecialType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpecialType

`func (o *MenuSectionCategoryItem) SetSpecialType(v string)`

SetSpecialType sets SpecialType field to given value.

### HasSpecialType

`func (o *MenuSectionCategoryItem) HasSpecialType() bool`

HasSpecialType returns a boolean if a field has been set.

### GetTaxable

`func (o *MenuSectionCategoryItem) GetTaxable() bool`

GetTaxable returns the Taxable field if non-nil, zero value otherwise.

### GetTaxableOk

`func (o *MenuSectionCategoryItem) GetTaxableOk() (*bool, bool)`

GetTaxableOk returns a tuple with the Taxable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaxable

`func (o *MenuSectionCategoryItem) SetTaxable(v bool)`

SetTaxable sets Taxable field to given value.

### HasTaxable

`func (o *MenuSectionCategoryItem) HasTaxable() bool`

HasTaxable returns a boolean if a field has been set.

### GetBarcode

`func (o *MenuSectionCategoryItem) GetBarcode() string`

GetBarcode returns the Barcode field if non-nil, zero value otherwise.

### GetBarcodeOk

`func (o *MenuSectionCategoryItem) GetBarcodeOk() (*string, bool)`

GetBarcodeOk returns a tuple with the Barcode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBarcode

`func (o *MenuSectionCategoryItem) SetBarcode(v string)`

SetBarcode sets Barcode field to given value.

### HasBarcode

`func (o *MenuSectionCategoryItem) HasBarcode() bool`

HasBarcode returns a boolean if a field has been set.

### GetMaxStock

`func (o *MenuSectionCategoryItem) GetMaxStock() int32`

GetMaxStock returns the MaxStock field if non-nil, zero value otherwise.

### GetMaxStockOk

`func (o *MenuSectionCategoryItem) GetMaxStockOk() (*int32, bool)`

GetMaxStockOk returns a tuple with the MaxStock field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxStock

`func (o *MenuSectionCategoryItem) SetMaxStock(v int32)`

SetMaxStock sets MaxStock field to given value.

### HasMaxStock

`func (o *MenuSectionCategoryItem) HasMaxStock() bool`

HasMaxStock returns a boolean if a field has been set.

### GetAdvancedPricing

`func (o *MenuSectionCategoryItem) GetAdvancedPricing() AdvancedPricing`

GetAdvancedPricing returns the AdvancedPricing field if non-nil, zero value otherwise.

### GetAdvancedPricingOk

`func (o *MenuSectionCategoryItem) GetAdvancedPricingOk() (*AdvancedPricing, bool)`

GetAdvancedPricingOk returns a tuple with the AdvancedPricing field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdvancedPricing

`func (o *MenuSectionCategoryItem) SetAdvancedPricing(v AdvancedPricing)`

SetAdvancedPricing sets AdvancedPricing field to given value.

### HasAdvancedPricing

`func (o *MenuSectionCategoryItem) HasAdvancedPricing() bool`

HasAdvancedPricing returns a boolean if a field has been set.

### GetPurchasability

`func (o *MenuSectionCategoryItem) GetPurchasability() Purchasability`

GetPurchasability returns the Purchasability field if non-nil, zero value otherwise.

### GetPurchasabilityOk

`func (o *MenuSectionCategoryItem) GetPurchasabilityOk() (*Purchasability, bool)`

GetPurchasabilityOk returns a tuple with the Purchasability field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPurchasability

`func (o *MenuSectionCategoryItem) SetPurchasability(v Purchasability)`

SetPurchasability sets Purchasability field to given value.

### HasPurchasability

`func (o *MenuSectionCategoryItem) HasPurchasability() bool`

HasPurchasability returns a boolean if a field has been set.

### GetModifierGroups

`func (o *MenuSectionCategoryItem) GetModifierGroups() []ModifierGroup`

GetModifierGroups returns the ModifierGroups field if non-nil, zero value otherwise.

### GetModifierGroupsOk

`func (o *MenuSectionCategoryItem) GetModifierGroupsOk() (*[]ModifierGroup, bool)`

GetModifierGroupsOk returns a tuple with the ModifierGroups field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModifierGroups

`func (o *MenuSectionCategoryItem) SetModifierGroups(v []ModifierGroup)`

SetModifierGroups sets ModifierGroups field to given value.

### HasModifierGroups

`func (o *MenuSectionCategoryItem) HasModifierGroups() bool`

HasModifierGroups returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


