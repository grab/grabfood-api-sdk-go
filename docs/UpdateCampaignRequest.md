# UpdateCampaignRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**MerchantID** | Pointer to **string** | The merchant&#39;s ID that is in GrabFood&#39;s database. | [optional] 
**Name** | Pointer to **string** | The campaign&#39;s name. | [optional] 
**Quotas** | Pointer to [**CampaignQuotas**](CampaignQuotas.md) |  | [optional] 
**Conditions** | Pointer to [**CampaignConditions**](CampaignConditions.md) |  | [optional] 
**Discount** | Pointer to [**CampaignDiscount**](CampaignDiscount.md) |  | [optional] 
**CustomTag** | Pointer to **string** | Specify the tag for custom bundle offer campaign. Only whitelisted partner is supported as of now. | [optional] 

## Methods

### NewUpdateCampaignRequest

`func NewUpdateCampaignRequest() *UpdateCampaignRequest`

NewUpdateCampaignRequest instantiates a new UpdateCampaignRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateCampaignRequestWithDefaults

`func NewUpdateCampaignRequestWithDefaults() *UpdateCampaignRequest`

NewUpdateCampaignRequestWithDefaults instantiates a new UpdateCampaignRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMerchantID

`func (o *UpdateCampaignRequest) GetMerchantID() string`

GetMerchantID returns the MerchantID field if non-nil, zero value otherwise.

### GetMerchantIDOk

`func (o *UpdateCampaignRequest) GetMerchantIDOk() (*string, bool)`

GetMerchantIDOk returns a tuple with the MerchantID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMerchantID

`func (o *UpdateCampaignRequest) SetMerchantID(v string)`

SetMerchantID sets MerchantID field to given value.

### HasMerchantID

`func (o *UpdateCampaignRequest) HasMerchantID() bool`

HasMerchantID returns a boolean if a field has been set.

### GetName

`func (o *UpdateCampaignRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *UpdateCampaignRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *UpdateCampaignRequest) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *UpdateCampaignRequest) HasName() bool`

HasName returns a boolean if a field has been set.

### GetQuotas

`func (o *UpdateCampaignRequest) GetQuotas() CampaignQuotas`

GetQuotas returns the Quotas field if non-nil, zero value otherwise.

### GetQuotasOk

`func (o *UpdateCampaignRequest) GetQuotasOk() (*CampaignQuotas, bool)`

GetQuotasOk returns a tuple with the Quotas field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuotas

`func (o *UpdateCampaignRequest) SetQuotas(v CampaignQuotas)`

SetQuotas sets Quotas field to given value.

### HasQuotas

`func (o *UpdateCampaignRequest) HasQuotas() bool`

HasQuotas returns a boolean if a field has been set.

### GetConditions

`func (o *UpdateCampaignRequest) GetConditions() CampaignConditions`

GetConditions returns the Conditions field if non-nil, zero value otherwise.

### GetConditionsOk

`func (o *UpdateCampaignRequest) GetConditionsOk() (*CampaignConditions, bool)`

GetConditionsOk returns a tuple with the Conditions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConditions

`func (o *UpdateCampaignRequest) SetConditions(v CampaignConditions)`

SetConditions sets Conditions field to given value.

### HasConditions

`func (o *UpdateCampaignRequest) HasConditions() bool`

HasConditions returns a boolean if a field has been set.

### GetDiscount

`func (o *UpdateCampaignRequest) GetDiscount() CampaignDiscount`

GetDiscount returns the Discount field if non-nil, zero value otherwise.

### GetDiscountOk

`func (o *UpdateCampaignRequest) GetDiscountOk() (*CampaignDiscount, bool)`

GetDiscountOk returns a tuple with the Discount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDiscount

`func (o *UpdateCampaignRequest) SetDiscount(v CampaignDiscount)`

SetDiscount sets Discount field to given value.

### HasDiscount

`func (o *UpdateCampaignRequest) HasDiscount() bool`

HasDiscount returns a boolean if a field has been set.

### GetCustomTag

`func (o *UpdateCampaignRequest) GetCustomTag() string`

GetCustomTag returns the CustomTag field if non-nil, zero value otherwise.

### GetCustomTagOk

`func (o *UpdateCampaignRequest) GetCustomTagOk() (*string, bool)`

GetCustomTagOk returns a tuple with the CustomTag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomTag

`func (o *UpdateCampaignRequest) SetCustomTag(v string)`

SetCustomTag sets CustomTag field to given value.

### HasCustomTag

`func (o *UpdateCampaignRequest) HasCustomTag() bool`

HasCustomTag returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


