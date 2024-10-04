# CampaignScope

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **string** | The scope type for this campaign.  * &#x60;order&#x60; - order level campaign. * &#x60;items&#x60; - item level campaign or bundle offer.  * &#x60;category&#x60; - category level campaign where all items within applies the same discount.  | 
**ObjectIDs** | Pointer to **[]string** | The list of item IDs in the partner&#39;s database applicable for discount when &#x60;discount.scope.type&#x60; is &#x60;items&#x60; (or category IDs for &#x60;category&#x60;).  One and only 1 item supported when &#x60;discount.type&#x60; is: - &#x60;freeItem&#x60; - &#x60;bundleSameNet&#x60; - &#x60;bundleSamePercentage&#x60; - &#x60;bundleSameFixPrice&#x60;  Minimum 2 - Maximum 20 items supported when &#x60;discount.type&#x60; is: - &#x60;bundleDiffNet&#x60; - &#x60;bundleDiffPercentage&#x60; - &#x60;bundleDiffFixPrice&#x60;  | [optional] 

## Methods

### NewCampaignScope

`func NewCampaignScope(type_ string, ) *CampaignScope`

NewCampaignScope instantiates a new CampaignScope object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCampaignScopeWithDefaults

`func NewCampaignScopeWithDefaults() *CampaignScope`

NewCampaignScopeWithDefaults instantiates a new CampaignScope object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *CampaignScope) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *CampaignScope) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *CampaignScope) SetType(v string)`

SetType sets Type field to given value.


### GetObjectIDs

`func (o *CampaignScope) GetObjectIDs() []string`

GetObjectIDs returns the ObjectIDs field if non-nil, zero value otherwise.

### GetObjectIDsOk

`func (o *CampaignScope) GetObjectIDsOk() (*[]string, bool)`

GetObjectIDsOk returns a tuple with the ObjectIDs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectIDs

`func (o *CampaignScope) SetObjectIDs(v []string)`

SetObjectIDs sets ObjectIDs field to given value.

### HasObjectIDs

`func (o *CampaignScope) HasObjectIDs() bool`

HasObjectIDs returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


