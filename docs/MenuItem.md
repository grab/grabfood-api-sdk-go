# MenuItem

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | The item&#39;s ID in the partner system.  | 
**Name** | **string** | The name of the item. | 
**NameTranslation** | Pointer to **map[string]string** | Translation of the item name. Only support up to 1 translated language. Refer [Menu Translation](#section/Menu-Translation). | [optional] 
**AvailableStatus** | **string** | The status for the item that is in the category.  Note: In order to set an item as \&quot;UNAVAILABLE\&quot;, it is required to update both the &#x60;availableStatus&#x60; and &#x60;maxStock&#x60; fields, whereby the &#x60;maxStock&#x60; value should be set to 0.  | 
**Description** | Pointer to **string** | The description of the item. There is a custom length limit of 2000 for &#x60;VN&#x60;.  | [optional] 
**DescriptionTranslation** | Pointer to **map[string]string** | Translation of the item description. Only support up to 1 translated language. Refer [Menu Translation](#section/Menu-Translation). | [optional] 
**Price** | **int64** | The item&#39;s price (excluding tax) in minor format. For example: 1900 means $19 with &#x60;currency.exponent&#x60; as 2. Refer to [FAQ](#section/Menu/Is-the-item-price-with-or-without-tax) for more details.  | 
**Photos** | Pointer to **[]string** | An array string for the item’s image URL links. Refer to FAQs for more details about [images](#section/Menu/What-are-the-recommended-formats-for-an-item-image).  | [optional] 
**SpecialType** | Pointer to **string** | The item&#39;s special Tag. Refer to FAQs for more details about [specialType](#section/Menu/What&#39;s-specialType).  | [optional] 
**Taxable** | Pointer to **bool** | **For Indonesia only.** This field allows the configuration for an item to be marked as tax applicable, and marked item would then be included in a commercial invoice to consumers as per the government&#39;s regulations.  | [optional] 
**Barcode** | Pointer to **string** | The barcode Number (GTIN). Max 64 allowed. GTIN must be 8, 12, 13, 14 numeric digits.  | [optional] 
**SellingTimeID** | Pointer to **string** | The selling time&#39;s ID for the item. This value overwrites the category&#39;s selling time if it is different. Empty value implies the category&#39;s selling time will be applied.  | [optional] 
**MaxStock** | Pointer to **int64** | Available stocks under inventory for this item. Auto reduce when there is order placed for this item. Empty value implies no limit.  Note: It is necessary to set &#x60;maxStock&#x60; to 0 if the &#x60;availableStatus&#x60; of the item is \&quot;UNAVAILABLE\&quot;. Item will be set to \&quot;AVAILABLE\&quot; if &#x60;maxStock&#x60; &gt; 0.  | [optional] 
**AdvancedPricing** | Pointer to [**AdvancedPricing**](AdvancedPricing.md) |  | [optional] 
**Purchasability** | Pointer to [**Purchasability**](Purchasability.md) |  | [optional] 
**ModifierGroups** | Pointer to [**[]ModifierGroup**](ModifierGroup.md) | An array of the modifierGroup JSON objects. Max 30 allowed per item. Refer to [Modifier groups](#modifier-groups) for more information. | [optional] 

## Methods

### NewMenuItem

`func NewMenuItem(id string, name string, availableStatus string, price int64, ) *MenuItem`

NewMenuItem instantiates a new MenuItem object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMenuItemWithDefaults

`func NewMenuItemWithDefaults() *MenuItem`

NewMenuItemWithDefaults instantiates a new MenuItem object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *MenuItem) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *MenuItem) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *MenuItem) SetId(v string)`

SetId sets Id field to given value.


### GetName

`func (o *MenuItem) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *MenuItem) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *MenuItem) SetName(v string)`

SetName sets Name field to given value.


### GetNameTranslation

`func (o *MenuItem) GetNameTranslation() map[string]string`

GetNameTranslation returns the NameTranslation field if non-nil, zero value otherwise.

### GetNameTranslationOk

`func (o *MenuItem) GetNameTranslationOk() (*map[string]string, bool)`

GetNameTranslationOk returns a tuple with the NameTranslation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNameTranslation

`func (o *MenuItem) SetNameTranslation(v map[string]string)`

SetNameTranslation sets NameTranslation field to given value.

### HasNameTranslation

`func (o *MenuItem) HasNameTranslation() bool`

HasNameTranslation returns a boolean if a field has been set.

### GetAvailableStatus

`func (o *MenuItem) GetAvailableStatus() string`

GetAvailableStatus returns the AvailableStatus field if non-nil, zero value otherwise.

### GetAvailableStatusOk

`func (o *MenuItem) GetAvailableStatusOk() (*string, bool)`

GetAvailableStatusOk returns a tuple with the AvailableStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvailableStatus

`func (o *MenuItem) SetAvailableStatus(v string)`

SetAvailableStatus sets AvailableStatus field to given value.


### GetDescription

`func (o *MenuItem) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *MenuItem) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *MenuItem) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *MenuItem) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetDescriptionTranslation

`func (o *MenuItem) GetDescriptionTranslation() map[string]string`

GetDescriptionTranslation returns the DescriptionTranslation field if non-nil, zero value otherwise.

### GetDescriptionTranslationOk

`func (o *MenuItem) GetDescriptionTranslationOk() (*map[string]string, bool)`

GetDescriptionTranslationOk returns a tuple with the DescriptionTranslation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescriptionTranslation

`func (o *MenuItem) SetDescriptionTranslation(v map[string]string)`

SetDescriptionTranslation sets DescriptionTranslation field to given value.

### HasDescriptionTranslation

`func (o *MenuItem) HasDescriptionTranslation() bool`

HasDescriptionTranslation returns a boolean if a field has been set.

### GetPrice

`func (o *MenuItem) GetPrice() int64`

GetPrice returns the Price field if non-nil, zero value otherwise.

### GetPriceOk

`func (o *MenuItem) GetPriceOk() (*int64, bool)`

GetPriceOk returns a tuple with the Price field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrice

`func (o *MenuItem) SetPrice(v int64)`

SetPrice sets Price field to given value.


### GetPhotos

`func (o *MenuItem) GetPhotos() []string`

GetPhotos returns the Photos field if non-nil, zero value otherwise.

### GetPhotosOk

`func (o *MenuItem) GetPhotosOk() (*[]string, bool)`

GetPhotosOk returns a tuple with the Photos field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPhotos

`func (o *MenuItem) SetPhotos(v []string)`

SetPhotos sets Photos field to given value.

### HasPhotos

`func (o *MenuItem) HasPhotos() bool`

HasPhotos returns a boolean if a field has been set.

### GetSpecialType

`func (o *MenuItem) GetSpecialType() string`

GetSpecialType returns the SpecialType field if non-nil, zero value otherwise.

### GetSpecialTypeOk

`func (o *MenuItem) GetSpecialTypeOk() (*string, bool)`

GetSpecialTypeOk returns a tuple with the SpecialType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpecialType

`func (o *MenuItem) SetSpecialType(v string)`

SetSpecialType sets SpecialType field to given value.

### HasSpecialType

`func (o *MenuItem) HasSpecialType() bool`

HasSpecialType returns a boolean if a field has been set.

### GetTaxable

`func (o *MenuItem) GetTaxable() bool`

GetTaxable returns the Taxable field if non-nil, zero value otherwise.

### GetTaxableOk

`func (o *MenuItem) GetTaxableOk() (*bool, bool)`

GetTaxableOk returns a tuple with the Taxable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaxable

`func (o *MenuItem) SetTaxable(v bool)`

SetTaxable sets Taxable field to given value.

### HasTaxable

`func (o *MenuItem) HasTaxable() bool`

HasTaxable returns a boolean if a field has been set.

### GetBarcode

`func (o *MenuItem) GetBarcode() string`

GetBarcode returns the Barcode field if non-nil, zero value otherwise.

### GetBarcodeOk

`func (o *MenuItem) GetBarcodeOk() (*string, bool)`

GetBarcodeOk returns a tuple with the Barcode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBarcode

`func (o *MenuItem) SetBarcode(v string)`

SetBarcode sets Barcode field to given value.

### HasBarcode

`func (o *MenuItem) HasBarcode() bool`

HasBarcode returns a boolean if a field has been set.

### GetSellingTimeID

`func (o *MenuItem) GetSellingTimeID() string`

GetSellingTimeID returns the SellingTimeID field if non-nil, zero value otherwise.

### GetSellingTimeIDOk

`func (o *MenuItem) GetSellingTimeIDOk() (*string, bool)`

GetSellingTimeIDOk returns a tuple with the SellingTimeID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSellingTimeID

`func (o *MenuItem) SetSellingTimeID(v string)`

SetSellingTimeID sets SellingTimeID field to given value.

### HasSellingTimeID

`func (o *MenuItem) HasSellingTimeID() bool`

HasSellingTimeID returns a boolean if a field has been set.

### GetMaxStock

`func (o *MenuItem) GetMaxStock() int64`

GetMaxStock returns the MaxStock field if non-nil, zero value otherwise.

### GetMaxStockOk

`func (o *MenuItem) GetMaxStockOk() (*int64, bool)`

GetMaxStockOk returns a tuple with the MaxStock field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxStock

`func (o *MenuItem) SetMaxStock(v int64)`

SetMaxStock sets MaxStock field to given value.

### HasMaxStock

`func (o *MenuItem) HasMaxStock() bool`

HasMaxStock returns a boolean if a field has been set.

### GetAdvancedPricing

`func (o *MenuItem) GetAdvancedPricing() AdvancedPricing`

GetAdvancedPricing returns the AdvancedPricing field if non-nil, zero value otherwise.

### GetAdvancedPricingOk

`func (o *MenuItem) GetAdvancedPricingOk() (*AdvancedPricing, bool)`

GetAdvancedPricingOk returns a tuple with the AdvancedPricing field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdvancedPricing

`func (o *MenuItem) SetAdvancedPricing(v AdvancedPricing)`

SetAdvancedPricing sets AdvancedPricing field to given value.

### HasAdvancedPricing

`func (o *MenuItem) HasAdvancedPricing() bool`

HasAdvancedPricing returns a boolean if a field has been set.

### GetPurchasability

`func (o *MenuItem) GetPurchasability() Purchasability`

GetPurchasability returns the Purchasability field if non-nil, zero value otherwise.

### GetPurchasabilityOk

`func (o *MenuItem) GetPurchasabilityOk() (*Purchasability, bool)`

GetPurchasabilityOk returns a tuple with the Purchasability field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPurchasability

`func (o *MenuItem) SetPurchasability(v Purchasability)`

SetPurchasability sets Purchasability field to given value.

### HasPurchasability

`func (o *MenuItem) HasPurchasability() bool`

HasPurchasability returns a boolean if a field has been set.

### GetModifierGroups

`func (o *MenuItem) GetModifierGroups() []ModifierGroup`

GetModifierGroups returns the ModifierGroups field if non-nil, zero value otherwise.

### GetModifierGroupsOk

`func (o *MenuItem) GetModifierGroupsOk() (*[]ModifierGroup, bool)`

GetModifierGroupsOk returns a tuple with the ModifierGroups field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModifierGroups

`func (o *MenuItem) SetModifierGroups(v []ModifierGroup)`

SetModifierGroups sets ModifierGroups field to given value.

### HasModifierGroups

`func (o *MenuItem) HasModifierGroups() bool`

HasModifierGroups returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


