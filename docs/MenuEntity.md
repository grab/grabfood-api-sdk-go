# MenuEntity

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | The record&#39;s ID on the partner system. For example, the item id in case type is ITEM, modifier id is MODIFIER. | [optional] 
**Price** | Pointer to **int64** | The record&#39;s price in minor unit format. | [optional] 
**AvailableStatus** | Pointer to **string** | The record&#39;s availableStatus.  Note: In order to set an item as \&quot;UNAVAILABLE\&quot;, it is required to update both the &#x60;availableStatus&#x60; and &#x60;maxStock&#x60; fields, whereby the &#x60;maxStock&#x60; value should be set to 0.  | [optional] 
**MaxStock** | Pointer to **int64** | Available stocks under inventory for this item. Auto reduce when there is order placed for this item.  Note: It is necessary to set &#x60;maxStock&#x60; to 0 if the &#x60;availableStatus&#x60; of the item is \&quot;UNAVAILABLE\&quot;. Item will be set to \&quot;AVAILABLE\&quot; if &#x60;maxStock&#x60; &gt; 0.  | [optional] 
**AdvancedPricings** | Pointer to [**[]UpdateAdvancedPricing**](UpdateAdvancedPricing.md) | Price configuration (in minor unit) for different service, order type and channel combination. If a service type does not have a specified price, it will utilize the default price of the item.  | [optional] 
**Purchasabilities** | Pointer to [**[]UpdatePurchasability**](UpdatePurchasability.md) | Purchasability is set to true by default for all service type, unless it is explicitly set to false. Modifier will reuse it’s item’s purchasability.  | [optional] 

## Methods

### NewMenuEntity

`func NewMenuEntity() *MenuEntity`

NewMenuEntity instantiates a new MenuEntity object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMenuEntityWithDefaults

`func NewMenuEntityWithDefaults() *MenuEntity`

NewMenuEntityWithDefaults instantiates a new MenuEntity object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *MenuEntity) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *MenuEntity) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *MenuEntity) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *MenuEntity) HasId() bool`

HasId returns a boolean if a field has been set.

### GetPrice

`func (o *MenuEntity) GetPrice() int64`

GetPrice returns the Price field if non-nil, zero value otherwise.

### GetPriceOk

`func (o *MenuEntity) GetPriceOk() (*int64, bool)`

GetPriceOk returns a tuple with the Price field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrice

`func (o *MenuEntity) SetPrice(v int64)`

SetPrice sets Price field to given value.

### HasPrice

`func (o *MenuEntity) HasPrice() bool`

HasPrice returns a boolean if a field has been set.

### GetAvailableStatus

`func (o *MenuEntity) GetAvailableStatus() string`

GetAvailableStatus returns the AvailableStatus field if non-nil, zero value otherwise.

### GetAvailableStatusOk

`func (o *MenuEntity) GetAvailableStatusOk() (*string, bool)`

GetAvailableStatusOk returns a tuple with the AvailableStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvailableStatus

`func (o *MenuEntity) SetAvailableStatus(v string)`

SetAvailableStatus sets AvailableStatus field to given value.

### HasAvailableStatus

`func (o *MenuEntity) HasAvailableStatus() bool`

HasAvailableStatus returns a boolean if a field has been set.

### GetMaxStock

`func (o *MenuEntity) GetMaxStock() int64`

GetMaxStock returns the MaxStock field if non-nil, zero value otherwise.

### GetMaxStockOk

`func (o *MenuEntity) GetMaxStockOk() (*int64, bool)`

GetMaxStockOk returns a tuple with the MaxStock field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxStock

`func (o *MenuEntity) SetMaxStock(v int64)`

SetMaxStock sets MaxStock field to given value.

### HasMaxStock

`func (o *MenuEntity) HasMaxStock() bool`

HasMaxStock returns a boolean if a field has been set.

### GetAdvancedPricings

`func (o *MenuEntity) GetAdvancedPricings() []UpdateAdvancedPricing`

GetAdvancedPricings returns the AdvancedPricings field if non-nil, zero value otherwise.

### GetAdvancedPricingsOk

`func (o *MenuEntity) GetAdvancedPricingsOk() (*[]UpdateAdvancedPricing, bool)`

GetAdvancedPricingsOk returns a tuple with the AdvancedPricings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdvancedPricings

`func (o *MenuEntity) SetAdvancedPricings(v []UpdateAdvancedPricing)`

SetAdvancedPricings sets AdvancedPricings field to given value.

### HasAdvancedPricings

`func (o *MenuEntity) HasAdvancedPricings() bool`

HasAdvancedPricings returns a boolean if a field has been set.

### GetPurchasabilities

`func (o *MenuEntity) GetPurchasabilities() []UpdatePurchasability`

GetPurchasabilities returns the Purchasabilities field if non-nil, zero value otherwise.

### GetPurchasabilitiesOk

`func (o *MenuEntity) GetPurchasabilitiesOk() (*[]UpdatePurchasability, bool)`

GetPurchasabilitiesOk returns a tuple with the Purchasabilities field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPurchasabilities

`func (o *MenuEntity) SetPurchasabilities(v []UpdatePurchasability)`

SetPurchasabilities sets Purchasabilities field to given value.

### HasPurchasabilities

`func (o *MenuEntity) HasPurchasabilities() bool`

HasPurchasabilities returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


