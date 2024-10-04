# CreateCampaignRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**MerchantID** | **string** | The merchant&#39;s ID that is in GrabFood&#39;s database. | 
**Name** | **string** | The campaign&#39;s name. | 
**Quotas** | Pointer to [**CampaignQuotas**](CampaignQuotas.md) |  | [optional] 
**Conditions** | [**CampaignConditions**](CampaignConditions.md) |  | 
**Discount** | [**CampaignDiscount**](CampaignDiscount.md) |  | 
**CustomTag** | Pointer to **string** | Specify the tag for custom bundle offer campaign. Only whitelisted partner is supported as of now. | [optional] 

## Methods

### NewCreateCampaignRequest

`func NewCreateCampaignRequest(merchantID string, name string, conditions CampaignConditions, discount CampaignDiscount, ) *CreateCampaignRequest`

NewCreateCampaignRequest instantiates a new CreateCampaignRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateCampaignRequestWithDefaults

`func NewCreateCampaignRequestWithDefaults() *CreateCampaignRequest`

NewCreateCampaignRequestWithDefaults instantiates a new CreateCampaignRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMerchantID

`func (o *CreateCampaignRequest) GetMerchantID() string`

GetMerchantID returns the MerchantID field if non-nil, zero value otherwise.

### GetMerchantIDOk

`func (o *CreateCampaignRequest) GetMerchantIDOk() (*string, bool)`

GetMerchantIDOk returns a tuple with the MerchantID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMerchantID

`func (o *CreateCampaignRequest) SetMerchantID(v string)`

SetMerchantID sets MerchantID field to given value.


### GetName

`func (o *CreateCampaignRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CreateCampaignRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CreateCampaignRequest) SetName(v string)`

SetName sets Name field to given value.


### GetQuotas

`func (o *CreateCampaignRequest) GetQuotas() CampaignQuotas`

GetQuotas returns the Quotas field if non-nil, zero value otherwise.

### GetQuotasOk

`func (o *CreateCampaignRequest) GetQuotasOk() (*CampaignQuotas, bool)`

GetQuotasOk returns a tuple with the Quotas field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuotas

`func (o *CreateCampaignRequest) SetQuotas(v CampaignQuotas)`

SetQuotas sets Quotas field to given value.

### HasQuotas

`func (o *CreateCampaignRequest) HasQuotas() bool`

HasQuotas returns a boolean if a field has been set.

### GetConditions

`func (o *CreateCampaignRequest) GetConditions() CampaignConditions`

GetConditions returns the Conditions field if non-nil, zero value otherwise.

### GetConditionsOk

`func (o *CreateCampaignRequest) GetConditionsOk() (*CampaignConditions, bool)`

GetConditionsOk returns a tuple with the Conditions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConditions

`func (o *CreateCampaignRequest) SetConditions(v CampaignConditions)`

SetConditions sets Conditions field to given value.


### GetDiscount

`func (o *CreateCampaignRequest) GetDiscount() CampaignDiscount`

GetDiscount returns the Discount field if non-nil, zero value otherwise.

### GetDiscountOk

`func (o *CreateCampaignRequest) GetDiscountOk() (*CampaignDiscount, bool)`

GetDiscountOk returns a tuple with the Discount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDiscount

`func (o *CreateCampaignRequest) SetDiscount(v CampaignDiscount)`

SetDiscount sets Discount field to given value.


### GetCustomTag

`func (o *CreateCampaignRequest) GetCustomTag() string`

GetCustomTag returns the CustomTag field if non-nil, zero value otherwise.

### GetCustomTagOk

`func (o *CreateCampaignRequest) GetCustomTagOk() (*string, bool)`

GetCustomTagOk returns a tuple with the CustomTag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomTag

`func (o *CreateCampaignRequest) SetCustomTag(v string)`

SetCustomTag sets CustomTag field to given value.

### HasCustomTag

`func (o *CreateCampaignRequest) HasCustomTag() bool`

HasCustomTag returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


