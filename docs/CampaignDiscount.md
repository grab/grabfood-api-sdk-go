# CampaignDiscount

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **string** | The type of discount  | 
**Cap** | Pointer to **float64** | The maximum discount dollar amount. It is **not required** and will be ignored when the &#x60;discount.type&#x60; is: - &#x60;net&#x60; - &#x60;delivery&#x60; - &#x60;freeItem&#x60; - &#x60;bundleSameNet&#x60; - &#x60;bundleSamePercentage&#x60; - &#x60;bundleSameFixPrice&#x60; - &#x60;bundleDiffNet&#x60; - &#x60;bundleDiffPercentage&#x60; - &#x60;bundleDiffFixPrice&#x60;  | [optional] 
**Value** | Pointer to **float64** | Specify the discount amount. Decimal number is not supported For VN, ID and TH. For example, &#x60;10.5&#x60; is not allowed and it should be &#x60;10.0&#x60;. * Dollar amount value when &#x60;discount.type&#x60; is &#x60;net&#x60;, &#x60;delivery&#x60;, &#x60;bundleSameNet&#x60;, &#x60;bundleSameFixPrice&#x60;, &#x60;bundleDiffNet&#x60;, &#x60;bundleDiffFixPrice&#x60;. * Percentage value (0-100) when &#x60;discount.type&#x60; is &#x60;percentage&#x60;, &#x60;bundleSamePercentage&#x60;, &#x60;bundleDiffPercentage&#x60;. * **Not required** when &#x60;discount.type&#x60; is &#x60;freeItem&#x60;.  | [optional] 
**Scope** | [**CampaignScope**](CampaignScope.md) |  | 

## Methods

### NewCampaignDiscount

`func NewCampaignDiscount(type_ string, scope CampaignScope, ) *CampaignDiscount`

NewCampaignDiscount instantiates a new CampaignDiscount object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCampaignDiscountWithDefaults

`func NewCampaignDiscountWithDefaults() *CampaignDiscount`

NewCampaignDiscountWithDefaults instantiates a new CampaignDiscount object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *CampaignDiscount) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *CampaignDiscount) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *CampaignDiscount) SetType(v string)`

SetType sets Type field to given value.


### GetCap

`func (o *CampaignDiscount) GetCap() float64`

GetCap returns the Cap field if non-nil, zero value otherwise.

### GetCapOk

`func (o *CampaignDiscount) GetCapOk() (*float64, bool)`

GetCapOk returns a tuple with the Cap field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCap

`func (o *CampaignDiscount) SetCap(v float64)`

SetCap sets Cap field to given value.

### HasCap

`func (o *CampaignDiscount) HasCap() bool`

HasCap returns a boolean if a field has been set.

### GetValue

`func (o *CampaignDiscount) GetValue() float64`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *CampaignDiscount) GetValueOk() (*float64, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *CampaignDiscount) SetValue(v float64)`

SetValue sets Value field to given value.

### HasValue

`func (o *CampaignDiscount) HasValue() bool`

HasValue returns a boolean if a field has been set.

### GetScope

`func (o *CampaignDiscount) GetScope() CampaignScope`

GetScope returns the Scope field if non-nil, zero value otherwise.

### GetScopeOk

`func (o *CampaignDiscount) GetScopeOk() (*CampaignScope, bool)`

GetScopeOk returns a tuple with the Scope field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScope

`func (o *CampaignDiscount) SetScope(v CampaignScope)`

SetScope sets Scope field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


