# UpdateMenuItem

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**MerchantID** | **string** | The merchant&#39;s ID that is in GrabFood&#39;s database. | 
**Field** | **string** | The record type that you want to update. | 
**Id** | **string** | The record&#39;s ID on the partner system. For example, the item id in case type is ITEM. | 
**Price** | Pointer to **int64** | The record&#39;s price in minor unit format. | [optional] 
**AvailableStatus** | Pointer to **string** | The record&#39;s availableStatus.   Note: In order to set an item as \&quot;UNAVAILABLE\&quot;, it is required to update both the &#x60;availableStatus&#x60; and &#x60;maxStock&#x60; fields, whereby the &#x60;maxStock&#x60; value should be set to 0.  | [optional] 
**MaxStock** | Pointer to **int64** | Available stocks under inventory for this item. Auto reduce when there is order placed for this item.  Note: It is necessary to set &#x60;maxStock&#x60; to 0 if the &#x60;availableStatus&#x60; of the item is \&quot;UNAVAILABLE\&quot;. Item will be set to \&quot;AVAILABLE\&quot; if &#x60;maxStock&#x60; &gt; 0.  | [optional] 
**AdvancedPricings** | Pointer to [**[]UpdateAdvancedPricing**](UpdateAdvancedPricing.md) | Price configuration (in minor unit) for different service, order type and channel combination. If a service type does not have a specified price, it will utilize the default price of the item.  | [optional] 
**Purchasabilities** | Pointer to [**[]UpdatePurchasability**](UpdatePurchasability.md) | Purchasability is set to true by default for all service type, unless it is explicitly set to false. Modifier will reuse it’s item’s purchasability.  | [optional] 

## Methods

### NewUpdateMenuItem

`func NewUpdateMenuItem(merchantID string, field string, id string, ) *UpdateMenuItem`

NewUpdateMenuItem instantiates a new UpdateMenuItem object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateMenuItemWithDefaults

`func NewUpdateMenuItemWithDefaults() *UpdateMenuItem`

NewUpdateMenuItemWithDefaults instantiates a new UpdateMenuItem object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMerchantID

`func (o *UpdateMenuItem) GetMerchantID() string`

GetMerchantID returns the MerchantID field if non-nil, zero value otherwise.

### GetMerchantIDOk

`func (o *UpdateMenuItem) GetMerchantIDOk() (*string, bool)`

GetMerchantIDOk returns a tuple with the MerchantID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMerchantID

`func (o *UpdateMenuItem) SetMerchantID(v string)`

SetMerchantID sets MerchantID field to given value.


### GetField

`func (o *UpdateMenuItem) GetField() string`

GetField returns the Field field if non-nil, zero value otherwise.

### GetFieldOk

`func (o *UpdateMenuItem) GetFieldOk() (*string, bool)`

GetFieldOk returns a tuple with the Field field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetField

`func (o *UpdateMenuItem) SetField(v string)`

SetField sets Field field to given value.


### GetId

`func (o *UpdateMenuItem) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *UpdateMenuItem) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *UpdateMenuItem) SetId(v string)`

SetId sets Id field to given value.


### GetPrice

`func (o *UpdateMenuItem) GetPrice() int64`

GetPrice returns the Price field if non-nil, zero value otherwise.

### GetPriceOk

`func (o *UpdateMenuItem) GetPriceOk() (*int64, bool)`

GetPriceOk returns a tuple with the Price field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrice

`func (o *UpdateMenuItem) SetPrice(v int64)`

SetPrice sets Price field to given value.

### HasPrice

`func (o *UpdateMenuItem) HasPrice() bool`

HasPrice returns a boolean if a field has been set.

### GetAvailableStatus

`func (o *UpdateMenuItem) GetAvailableStatus() string`

GetAvailableStatus returns the AvailableStatus field if non-nil, zero value otherwise.

### GetAvailableStatusOk

`func (o *UpdateMenuItem) GetAvailableStatusOk() (*string, bool)`

GetAvailableStatusOk returns a tuple with the AvailableStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvailableStatus

`func (o *UpdateMenuItem) SetAvailableStatus(v string)`

SetAvailableStatus sets AvailableStatus field to given value.

### HasAvailableStatus

`func (o *UpdateMenuItem) HasAvailableStatus() bool`

HasAvailableStatus returns a boolean if a field has been set.

### GetMaxStock

`func (o *UpdateMenuItem) GetMaxStock() int64`

GetMaxStock returns the MaxStock field if non-nil, zero value otherwise.

### GetMaxStockOk

`func (o *UpdateMenuItem) GetMaxStockOk() (*int64, bool)`

GetMaxStockOk returns a tuple with the MaxStock field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxStock

`func (o *UpdateMenuItem) SetMaxStock(v int64)`

SetMaxStock sets MaxStock field to given value.

### HasMaxStock

`func (o *UpdateMenuItem) HasMaxStock() bool`

HasMaxStock returns a boolean if a field has been set.

### GetAdvancedPricings

`func (o *UpdateMenuItem) GetAdvancedPricings() []UpdateAdvancedPricing`

GetAdvancedPricings returns the AdvancedPricings field if non-nil, zero value otherwise.

### GetAdvancedPricingsOk

`func (o *UpdateMenuItem) GetAdvancedPricingsOk() (*[]UpdateAdvancedPricing, bool)`

GetAdvancedPricingsOk returns a tuple with the AdvancedPricings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdvancedPricings

`func (o *UpdateMenuItem) SetAdvancedPricings(v []UpdateAdvancedPricing)`

SetAdvancedPricings sets AdvancedPricings field to given value.

### HasAdvancedPricings

`func (o *UpdateMenuItem) HasAdvancedPricings() bool`

HasAdvancedPricings returns a boolean if a field has been set.

### GetPurchasabilities

`func (o *UpdateMenuItem) GetPurchasabilities() []UpdatePurchasability`

GetPurchasabilities returns the Purchasabilities field if non-nil, zero value otherwise.

### GetPurchasabilitiesOk

`func (o *UpdateMenuItem) GetPurchasabilitiesOk() (*[]UpdatePurchasability, bool)`

GetPurchasabilitiesOk returns a tuple with the Purchasabilities field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPurchasabilities

`func (o *UpdateMenuItem) SetPurchasabilities(v []UpdatePurchasability)`

SetPurchasabilities sets Purchasabilities field to given value.

### HasPurchasabilities

`func (o *UpdateMenuItem) HasPurchasabilities() bool`

HasPurchasabilities returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


