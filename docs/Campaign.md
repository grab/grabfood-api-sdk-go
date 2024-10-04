# Campaign

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | The campaign&#39;s ID. | 
**CreatedBy** | **string** | The party who created the campaign. Can be created by partners via API, merchants via the merchant app or Grab. | 
**MerchantID** | **string** | The merchant&#39;s ID that is in GrabFood&#39;s database. | 
**Name** | **string** | The campaign&#39;s name. | 
**Quotas** | Pointer to [**CampaignQuotas**](CampaignQuotas.md) |  | [optional] 
**Conditions** | Pointer to [**CampaignConditions**](CampaignConditions.md) |  | [optional] 
**Discount** | Pointer to [**CampaignDiscount**](CampaignDiscount.md) |  | [optional] 
**CustomTag** | Pointer to **string** | Specify the tag for custom bundle offer campaign. Only whitelisted partner is supported as of now. | [optional] 

## Methods

### NewCampaign

`func NewCampaign(id string, createdBy string, merchantID string, name string, ) *Campaign`

NewCampaign instantiates a new Campaign object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCampaignWithDefaults

`func NewCampaignWithDefaults() *Campaign`

NewCampaignWithDefaults instantiates a new Campaign object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *Campaign) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Campaign) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Campaign) SetId(v string)`

SetId sets Id field to given value.


### GetCreatedBy

`func (o *Campaign) GetCreatedBy() string`

GetCreatedBy returns the CreatedBy field if non-nil, zero value otherwise.

### GetCreatedByOk

`func (o *Campaign) GetCreatedByOk() (*string, bool)`

GetCreatedByOk returns a tuple with the CreatedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedBy

`func (o *Campaign) SetCreatedBy(v string)`

SetCreatedBy sets CreatedBy field to given value.


### GetMerchantID

`func (o *Campaign) GetMerchantID() string`

GetMerchantID returns the MerchantID field if non-nil, zero value otherwise.

### GetMerchantIDOk

`func (o *Campaign) GetMerchantIDOk() (*string, bool)`

GetMerchantIDOk returns a tuple with the MerchantID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMerchantID

`func (o *Campaign) SetMerchantID(v string)`

SetMerchantID sets MerchantID field to given value.


### GetName

`func (o *Campaign) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *Campaign) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *Campaign) SetName(v string)`

SetName sets Name field to given value.


### GetQuotas

`func (o *Campaign) GetQuotas() CampaignQuotas`

GetQuotas returns the Quotas field if non-nil, zero value otherwise.

### GetQuotasOk

`func (o *Campaign) GetQuotasOk() (*CampaignQuotas, bool)`

GetQuotasOk returns a tuple with the Quotas field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuotas

`func (o *Campaign) SetQuotas(v CampaignQuotas)`

SetQuotas sets Quotas field to given value.

### HasQuotas

`func (o *Campaign) HasQuotas() bool`

HasQuotas returns a boolean if a field has been set.

### GetConditions

`func (o *Campaign) GetConditions() CampaignConditions`

GetConditions returns the Conditions field if non-nil, zero value otherwise.

### GetConditionsOk

`func (o *Campaign) GetConditionsOk() (*CampaignConditions, bool)`

GetConditionsOk returns a tuple with the Conditions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConditions

`func (o *Campaign) SetConditions(v CampaignConditions)`

SetConditions sets Conditions field to given value.

### HasConditions

`func (o *Campaign) HasConditions() bool`

HasConditions returns a boolean if a field has been set.

### GetDiscount

`func (o *Campaign) GetDiscount() CampaignDiscount`

GetDiscount returns the Discount field if non-nil, zero value otherwise.

### GetDiscountOk

`func (o *Campaign) GetDiscountOk() (*CampaignDiscount, bool)`

GetDiscountOk returns a tuple with the Discount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDiscount

`func (o *Campaign) SetDiscount(v CampaignDiscount)`

SetDiscount sets Discount field to given value.

### HasDiscount

`func (o *Campaign) HasDiscount() bool`

HasDiscount returns a boolean if a field has been set.

### GetCustomTag

`func (o *Campaign) GetCustomTag() string`

GetCustomTag returns the CustomTag field if non-nil, zero value otherwise.

### GetCustomTagOk

`func (o *Campaign) GetCustomTagOk() (*string, bool)`

GetCustomTagOk returns a tuple with the CustomTag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomTag

`func (o *Campaign) SetCustomTag(v string)`

SetCustomTag sets CustomTag field to given value.

### HasCustomTag

`func (o *Campaign) HasCustomTag() bool`

HasCustomTag returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


