# MenuModifier

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | The modifier&#39;s ID that is on the partner&#39;s system. This ID should be unique with a min length of 1 and max of 64. | 
**Name** | **string** | The name of the modifier. | 
**NameTranslation** | Pointer to **map[string]string** | Translation of the modifier name. Only support up to 1 translated language. Refer [Menu Translation](#section/Menu-Translation). | [optional] 
**AvailableStatus** | **string** | The status for the modifier that is in the ModifierGroup. | 
**Price** | Pointer to **int64** | The modifier&#39;s price (excluding tax) in minor format. Refer to FAQs for more details about [tax](#section/Order/How-is-tax-calculated). | [optional] 
**Barcode** | Pointer to **string** | The barcode Number (GTIN). GTIN must be 8, 12, 13, 14 numeric digits. | [optional] 
**AdvancedPricing** | Pointer to [**AdvancedPricing**](AdvancedPricing.md) |  | [optional] 

## Methods

### NewMenuModifier

`func NewMenuModifier(id string, name string, availableStatus string, ) *MenuModifier`

NewMenuModifier instantiates a new MenuModifier object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMenuModifierWithDefaults

`func NewMenuModifierWithDefaults() *MenuModifier`

NewMenuModifierWithDefaults instantiates a new MenuModifier object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *MenuModifier) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *MenuModifier) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *MenuModifier) SetId(v string)`

SetId sets Id field to given value.


### GetName

`func (o *MenuModifier) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *MenuModifier) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *MenuModifier) SetName(v string)`

SetName sets Name field to given value.


### GetNameTranslation

`func (o *MenuModifier) GetNameTranslation() map[string]string`

GetNameTranslation returns the NameTranslation field if non-nil, zero value otherwise.

### GetNameTranslationOk

`func (o *MenuModifier) GetNameTranslationOk() (*map[string]string, bool)`

GetNameTranslationOk returns a tuple with the NameTranslation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNameTranslation

`func (o *MenuModifier) SetNameTranslation(v map[string]string)`

SetNameTranslation sets NameTranslation field to given value.

### HasNameTranslation

`func (o *MenuModifier) HasNameTranslation() bool`

HasNameTranslation returns a boolean if a field has been set.

### GetAvailableStatus

`func (o *MenuModifier) GetAvailableStatus() string`

GetAvailableStatus returns the AvailableStatus field if non-nil, zero value otherwise.

### GetAvailableStatusOk

`func (o *MenuModifier) GetAvailableStatusOk() (*string, bool)`

GetAvailableStatusOk returns a tuple with the AvailableStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvailableStatus

`func (o *MenuModifier) SetAvailableStatus(v string)`

SetAvailableStatus sets AvailableStatus field to given value.


### GetPrice

`func (o *MenuModifier) GetPrice() int64`

GetPrice returns the Price field if non-nil, zero value otherwise.

### GetPriceOk

`func (o *MenuModifier) GetPriceOk() (*int64, bool)`

GetPriceOk returns a tuple with the Price field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrice

`func (o *MenuModifier) SetPrice(v int64)`

SetPrice sets Price field to given value.

### HasPrice

`func (o *MenuModifier) HasPrice() bool`

HasPrice returns a boolean if a field has been set.

### GetBarcode

`func (o *MenuModifier) GetBarcode() string`

GetBarcode returns the Barcode field if non-nil, zero value otherwise.

### GetBarcodeOk

`func (o *MenuModifier) GetBarcodeOk() (*string, bool)`

GetBarcodeOk returns a tuple with the Barcode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBarcode

`func (o *MenuModifier) SetBarcode(v string)`

SetBarcode sets Barcode field to given value.

### HasBarcode

`func (o *MenuModifier) HasBarcode() bool`

HasBarcode returns a boolean if a field has been set.

### GetAdvancedPricing

`func (o *MenuModifier) GetAdvancedPricing() AdvancedPricing`

GetAdvancedPricing returns the AdvancedPricing field if non-nil, zero value otherwise.

### GetAdvancedPricingOk

`func (o *MenuModifier) GetAdvancedPricingOk() (*AdvancedPricing, bool)`

GetAdvancedPricingOk returns a tuple with the AdvancedPricing field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdvancedPricing

`func (o *MenuModifier) SetAdvancedPricing(v AdvancedPricing)`

SetAdvancedPricing sets AdvancedPricing field to given value.

### HasAdvancedPricing

`func (o *MenuModifier) HasAdvancedPricing() bool`

HasAdvancedPricing returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


