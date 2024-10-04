# UpdateMenuModifier

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**MerchantID** | **string** | The merchant&#39;s ID that is in GrabFood&#39;s database. | 
**Field** | **string** | The record type that you want to update. | 
**Id** | **string** | The record&#39;s ID on the partner system. For example, the modifier id in case type is MODIFIER. | 
**Price** | Pointer to **int64** | The record&#39;s price in minor unit format. | [optional] 
**Name** | **string** | **Only required when updating modifiers.** The record&#39;s name. Used as identifier to locate the correct record. | 
**AvailableStatus** | Pointer to **string** | The record&#39;s availableStatus. | [optional] 
**IsFree** | Pointer to **bool** | Allows the modifier&#39;s price to be explicitly set as zero. Possible values are as follows:   * &#x60;isFree&#x60; &amp;&amp; &#x60;price &#x3D;&#x3D; 0&#x60; sets the modifier&#39;s price to zero.   * &#x60;isFree&#x60; &amp;&amp; &#x60;price &gt; 0&#x60; returns an error message that \&quot;price cannot be set to &gt; 0, if modifier is free”.   * &#x60;!isFree&#x60; &amp;&amp; &#x60;price &gt; 0&#x60; sets the modifier&#39;s price to the defined price.   * &#x60;!isFree&#x60; &amp;&amp; &#x60;price &#x3D;&#x3D; 0&#x60; does not update the modifier&#39;s price and reuses the existing price.  | [optional] 
**AdvancedPricings** | Pointer to [**[]UpdateAdvancedPricing**](UpdateAdvancedPricing.md) | Price configuration (in minor unit) for different service, order type and channel combination. If a service type does not have a specified price, it will utilize the default price of the item.  | [optional] 

## Methods

### NewUpdateMenuModifier

`func NewUpdateMenuModifier(merchantID string, field string, id string, name string, ) *UpdateMenuModifier`

NewUpdateMenuModifier instantiates a new UpdateMenuModifier object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateMenuModifierWithDefaults

`func NewUpdateMenuModifierWithDefaults() *UpdateMenuModifier`

NewUpdateMenuModifierWithDefaults instantiates a new UpdateMenuModifier object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMerchantID

`func (o *UpdateMenuModifier) GetMerchantID() string`

GetMerchantID returns the MerchantID field if non-nil, zero value otherwise.

### GetMerchantIDOk

`func (o *UpdateMenuModifier) GetMerchantIDOk() (*string, bool)`

GetMerchantIDOk returns a tuple with the MerchantID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMerchantID

`func (o *UpdateMenuModifier) SetMerchantID(v string)`

SetMerchantID sets MerchantID field to given value.


### GetField

`func (o *UpdateMenuModifier) GetField() string`

GetField returns the Field field if non-nil, zero value otherwise.

### GetFieldOk

`func (o *UpdateMenuModifier) GetFieldOk() (*string, bool)`

GetFieldOk returns a tuple with the Field field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetField

`func (o *UpdateMenuModifier) SetField(v string)`

SetField sets Field field to given value.


### GetId

`func (o *UpdateMenuModifier) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *UpdateMenuModifier) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *UpdateMenuModifier) SetId(v string)`

SetId sets Id field to given value.


### GetPrice

`func (o *UpdateMenuModifier) GetPrice() int64`

GetPrice returns the Price field if non-nil, zero value otherwise.

### GetPriceOk

`func (o *UpdateMenuModifier) GetPriceOk() (*int64, bool)`

GetPriceOk returns a tuple with the Price field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrice

`func (o *UpdateMenuModifier) SetPrice(v int64)`

SetPrice sets Price field to given value.

### HasPrice

`func (o *UpdateMenuModifier) HasPrice() bool`

HasPrice returns a boolean if a field has been set.

### GetName

`func (o *UpdateMenuModifier) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *UpdateMenuModifier) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *UpdateMenuModifier) SetName(v string)`

SetName sets Name field to given value.


### GetAvailableStatus

`func (o *UpdateMenuModifier) GetAvailableStatus() string`

GetAvailableStatus returns the AvailableStatus field if non-nil, zero value otherwise.

### GetAvailableStatusOk

`func (o *UpdateMenuModifier) GetAvailableStatusOk() (*string, bool)`

GetAvailableStatusOk returns a tuple with the AvailableStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvailableStatus

`func (o *UpdateMenuModifier) SetAvailableStatus(v string)`

SetAvailableStatus sets AvailableStatus field to given value.

### HasAvailableStatus

`func (o *UpdateMenuModifier) HasAvailableStatus() bool`

HasAvailableStatus returns a boolean if a field has been set.

### GetIsFree

`func (o *UpdateMenuModifier) GetIsFree() bool`

GetIsFree returns the IsFree field if non-nil, zero value otherwise.

### GetIsFreeOk

`func (o *UpdateMenuModifier) GetIsFreeOk() (*bool, bool)`

GetIsFreeOk returns a tuple with the IsFree field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsFree

`func (o *UpdateMenuModifier) SetIsFree(v bool)`

SetIsFree sets IsFree field to given value.

### HasIsFree

`func (o *UpdateMenuModifier) HasIsFree() bool`

HasIsFree returns a boolean if a field has been set.

### GetAdvancedPricings

`func (o *UpdateMenuModifier) GetAdvancedPricings() []UpdateAdvancedPricing`

GetAdvancedPricings returns the AdvancedPricings field if non-nil, zero value otherwise.

### GetAdvancedPricingsOk

`func (o *UpdateMenuModifier) GetAdvancedPricingsOk() (*[]UpdateAdvancedPricing, bool)`

GetAdvancedPricingsOk returns a tuple with the AdvancedPricings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdvancedPricings

`func (o *UpdateMenuModifier) SetAdvancedPricings(v []UpdateAdvancedPricing)`

SetAdvancedPricings sets AdvancedPricings field to given value.

### HasAdvancedPricings

`func (o *UpdateMenuModifier) HasAdvancedPricings() bool`

HasAdvancedPricings returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


