# CampaignQuotas

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**TotalCount** | Pointer to **int64** | The maximum number of redemptions. Default is unlimited if unspecified.  | [optional] 
**TotalCountPerUser** | Pointer to **int64** | The maximum number of redemptions per user. Default is unlimited if unspecified. | [optional] 

## Methods

### NewCampaignQuotas

`func NewCampaignQuotas() *CampaignQuotas`

NewCampaignQuotas instantiates a new CampaignQuotas object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCampaignQuotasWithDefaults

`func NewCampaignQuotasWithDefaults() *CampaignQuotas`

NewCampaignQuotasWithDefaults instantiates a new CampaignQuotas object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTotalCount

`func (o *CampaignQuotas) GetTotalCount() int64`

GetTotalCount returns the TotalCount field if non-nil, zero value otherwise.

### GetTotalCountOk

`func (o *CampaignQuotas) GetTotalCountOk() (*int64, bool)`

GetTotalCountOk returns a tuple with the TotalCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalCount

`func (o *CampaignQuotas) SetTotalCount(v int64)`

SetTotalCount sets TotalCount field to given value.

### HasTotalCount

`func (o *CampaignQuotas) HasTotalCount() bool`

HasTotalCount returns a boolean if a field has been set.

### GetTotalCountPerUser

`func (o *CampaignQuotas) GetTotalCountPerUser() int64`

GetTotalCountPerUser returns the TotalCountPerUser field if non-nil, zero value otherwise.

### GetTotalCountPerUserOk

`func (o *CampaignQuotas) GetTotalCountPerUserOk() (*int64, bool)`

GetTotalCountPerUserOk returns a tuple with the TotalCountPerUser field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalCountPerUser

`func (o *CampaignQuotas) SetTotalCountPerUser(v int64)`

SetTotalCountPerUser sets TotalCountPerUser field to given value.

### HasTotalCountPerUser

`func (o *CampaignQuotas) HasTotalCountPerUser() bool`

HasTotalCountPerUser returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


