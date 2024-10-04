# RewardItem

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ItemID** | Pointer to **string** | The item&#39;s ID in partner system. | [optional] 
**Quantity** | Pointer to **int32** | The item&#39;s quantity. | [optional] 

## Methods

### NewRewardItem

`func NewRewardItem() *RewardItem`

NewRewardItem instantiates a new RewardItem object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRewardItemWithDefaults

`func NewRewardItemWithDefaults() *RewardItem`

NewRewardItemWithDefaults instantiates a new RewardItem object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetItemID

`func (o *RewardItem) GetItemID() string`

GetItemID returns the ItemID field if non-nil, zero value otherwise.

### GetItemIDOk

`func (o *RewardItem) GetItemIDOk() (*string, bool)`

GetItemIDOk returns a tuple with the ItemID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItemID

`func (o *RewardItem) SetItemID(v string)`

SetItemID sets ItemID field to given value.

### HasItemID

`func (o *RewardItem) HasItemID() bool`

HasItemID returns a boolean if a field has been set.

### GetQuantity

`func (o *RewardItem) GetQuantity() int32`

GetQuantity returns the Quantity field if non-nil, zero value otherwise.

### GetQuantityOk

`func (o *RewardItem) GetQuantityOk() (*int32, bool)`

GetQuantityOk returns a tuple with the Quantity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuantity

`func (o *RewardItem) SetQuantity(v int32)`

SetQuantity sets Quantity field to given value.

### HasQuantity

`func (o *RewardItem) HasQuantity() bool`

HasQuantity returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


