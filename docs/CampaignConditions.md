# CampaignConditions

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**StartTime** | **time.Time** | The campaign&#39;s start time in UTC format. For example, 2021-09-23T03:30:00Z means 2021-09-23 11:30:00 (UTC +08:00). | 
**EndTime** | **time.Time** | The campaign&#39;s end time in UTC format. | 
**EaterType** | **string** | The type of eater eligible for the campaign.  * &#x60;all&#x60; - campaign will be applied to everyone. No limitation on campaign type. * &#x60;new&#x60; - campaign will be applied to consumers who have not ordered from this store in the last three months. Only applicable to **order-level** campaign.  | 
**MinBasketAmount** | Pointer to **float64** | The minimum basket amount to be eligible for the campaign. Only applicable for **order-level** campaign. For example, 10.5 means the basket amount has to be at least $10.50. | [optional] 
**BundleQuantity** | Pointer to **int32** | Specify the bundle quantity for bundle offer campaign. | [optional] 
**WorkingHour** | Pointer to [**WorkingHour**](WorkingHour.md) |  | [optional] 

## Methods

### NewCampaignConditions

`func NewCampaignConditions(startTime time.Time, endTime time.Time, eaterType string, ) *CampaignConditions`

NewCampaignConditions instantiates a new CampaignConditions object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCampaignConditionsWithDefaults

`func NewCampaignConditionsWithDefaults() *CampaignConditions`

NewCampaignConditionsWithDefaults instantiates a new CampaignConditions object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetStartTime

`func (o *CampaignConditions) GetStartTime() time.Time`

GetStartTime returns the StartTime field if non-nil, zero value otherwise.

### GetStartTimeOk

`func (o *CampaignConditions) GetStartTimeOk() (*time.Time, bool)`

GetStartTimeOk returns a tuple with the StartTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartTime

`func (o *CampaignConditions) SetStartTime(v time.Time)`

SetStartTime sets StartTime field to given value.


### GetEndTime

`func (o *CampaignConditions) GetEndTime() time.Time`

GetEndTime returns the EndTime field if non-nil, zero value otherwise.

### GetEndTimeOk

`func (o *CampaignConditions) GetEndTimeOk() (*time.Time, bool)`

GetEndTimeOk returns a tuple with the EndTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndTime

`func (o *CampaignConditions) SetEndTime(v time.Time)`

SetEndTime sets EndTime field to given value.


### GetEaterType

`func (o *CampaignConditions) GetEaterType() string`

GetEaterType returns the EaterType field if non-nil, zero value otherwise.

### GetEaterTypeOk

`func (o *CampaignConditions) GetEaterTypeOk() (*string, bool)`

GetEaterTypeOk returns a tuple with the EaterType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEaterType

`func (o *CampaignConditions) SetEaterType(v string)`

SetEaterType sets EaterType field to given value.


### GetMinBasketAmount

`func (o *CampaignConditions) GetMinBasketAmount() float64`

GetMinBasketAmount returns the MinBasketAmount field if non-nil, zero value otherwise.

### GetMinBasketAmountOk

`func (o *CampaignConditions) GetMinBasketAmountOk() (*float64, bool)`

GetMinBasketAmountOk returns a tuple with the MinBasketAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinBasketAmount

`func (o *CampaignConditions) SetMinBasketAmount(v float64)`

SetMinBasketAmount sets MinBasketAmount field to given value.

### HasMinBasketAmount

`func (o *CampaignConditions) HasMinBasketAmount() bool`

HasMinBasketAmount returns a boolean if a field has been set.

### GetBundleQuantity

`func (o *CampaignConditions) GetBundleQuantity() int32`

GetBundleQuantity returns the BundleQuantity field if non-nil, zero value otherwise.

### GetBundleQuantityOk

`func (o *CampaignConditions) GetBundleQuantityOk() (*int32, bool)`

GetBundleQuantityOk returns a tuple with the BundleQuantity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBundleQuantity

`func (o *CampaignConditions) SetBundleQuantity(v int32)`

SetBundleQuantity sets BundleQuantity field to given value.

### HasBundleQuantity

`func (o *CampaignConditions) HasBundleQuantity() bool`

HasBundleQuantity returns a boolean if a field has been set.

### GetWorkingHour

`func (o *CampaignConditions) GetWorkingHour() WorkingHour`

GetWorkingHour returns the WorkingHour field if non-nil, zero value otherwise.

### GetWorkingHourOk

`func (o *CampaignConditions) GetWorkingHourOk() (*WorkingHour, bool)`

GetWorkingHourOk returns a tuple with the WorkingHour field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWorkingHour

`func (o *CampaignConditions) SetWorkingHour(v WorkingHour)`

SetWorkingHour sets WorkingHour field to given value.

### HasWorkingHour

`func (o *CampaignConditions) HasWorkingHour() bool`

HasWorkingHour returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


