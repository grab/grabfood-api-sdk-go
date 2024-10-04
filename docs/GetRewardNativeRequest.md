# GetRewardNativeRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**MemberID** | Pointer to **string** | The unique member ID on the partner&#39;s database. | [optional] 
**MerchantID** | Pointer to **string** | Grab merchant&#39;s ID. | [optional] 
**Items** | Pointer to [**[]RewardItem**](RewardItem.md) |  | [optional] 
**OrderValue** | Pointer to **int64** | The post-discount order value. | [optional] 

## Methods

### NewGetRewardNativeRequest

`func NewGetRewardNativeRequest() *GetRewardNativeRequest`

NewGetRewardNativeRequest instantiates a new GetRewardNativeRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetRewardNativeRequestWithDefaults

`func NewGetRewardNativeRequestWithDefaults() *GetRewardNativeRequest`

NewGetRewardNativeRequestWithDefaults instantiates a new GetRewardNativeRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMemberID

`func (o *GetRewardNativeRequest) GetMemberID() string`

GetMemberID returns the MemberID field if non-nil, zero value otherwise.

### GetMemberIDOk

`func (o *GetRewardNativeRequest) GetMemberIDOk() (*string, bool)`

GetMemberIDOk returns a tuple with the MemberID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMemberID

`func (o *GetRewardNativeRequest) SetMemberID(v string)`

SetMemberID sets MemberID field to given value.

### HasMemberID

`func (o *GetRewardNativeRequest) HasMemberID() bool`

HasMemberID returns a boolean if a field has been set.

### GetMerchantID

`func (o *GetRewardNativeRequest) GetMerchantID() string`

GetMerchantID returns the MerchantID field if non-nil, zero value otherwise.

### GetMerchantIDOk

`func (o *GetRewardNativeRequest) GetMerchantIDOk() (*string, bool)`

GetMerchantIDOk returns a tuple with the MerchantID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMerchantID

`func (o *GetRewardNativeRequest) SetMerchantID(v string)`

SetMerchantID sets MerchantID field to given value.

### HasMerchantID

`func (o *GetRewardNativeRequest) HasMerchantID() bool`

HasMerchantID returns a boolean if a field has been set.

### GetItems

`func (o *GetRewardNativeRequest) GetItems() []RewardItem`

GetItems returns the Items field if non-nil, zero value otherwise.

### GetItemsOk

`func (o *GetRewardNativeRequest) GetItemsOk() (*[]RewardItem, bool)`

GetItemsOk returns a tuple with the Items field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItems

`func (o *GetRewardNativeRequest) SetItems(v []RewardItem)`

SetItems sets Items field to given value.

### HasItems

`func (o *GetRewardNativeRequest) HasItems() bool`

HasItems returns a boolean if a field has been set.

### GetOrderValue

`func (o *GetRewardNativeRequest) GetOrderValue() int64`

GetOrderValue returns the OrderValue field if non-nil, zero value otherwise.

### GetOrderValueOk

`func (o *GetRewardNativeRequest) GetOrderValueOk() (*int64, bool)`

GetOrderValueOk returns a tuple with the OrderValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrderValue

`func (o *GetRewardNativeRequest) SetOrderValue(v int64)`

SetOrderValue sets OrderValue field to given value.

### HasOrderValue

`func (o *GetRewardNativeRequest) HasOrderValue() bool`

HasOrderValue returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


