# GetMembershipNativeResponsePointInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CurrentPoints** | Pointer to **int64** | Point that user currently obtained. | [optional] 
**MaxPoints** | Pointer to **int64** | Maximum point that user can obtain. | [optional] 
**ExpirePoints** | Pointer to **int64** | Points that would get expired in future. | [optional] 

## Methods

### NewGetMembershipNativeResponsePointInfo

`func NewGetMembershipNativeResponsePointInfo() *GetMembershipNativeResponsePointInfo`

NewGetMembershipNativeResponsePointInfo instantiates a new GetMembershipNativeResponsePointInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetMembershipNativeResponsePointInfoWithDefaults

`func NewGetMembershipNativeResponsePointInfoWithDefaults() *GetMembershipNativeResponsePointInfo`

NewGetMembershipNativeResponsePointInfoWithDefaults instantiates a new GetMembershipNativeResponsePointInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCurrentPoints

`func (o *GetMembershipNativeResponsePointInfo) GetCurrentPoints() int64`

GetCurrentPoints returns the CurrentPoints field if non-nil, zero value otherwise.

### GetCurrentPointsOk

`func (o *GetMembershipNativeResponsePointInfo) GetCurrentPointsOk() (*int64, bool)`

GetCurrentPointsOk returns a tuple with the CurrentPoints field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrentPoints

`func (o *GetMembershipNativeResponsePointInfo) SetCurrentPoints(v int64)`

SetCurrentPoints sets CurrentPoints field to given value.

### HasCurrentPoints

`func (o *GetMembershipNativeResponsePointInfo) HasCurrentPoints() bool`

HasCurrentPoints returns a boolean if a field has been set.

### GetMaxPoints

`func (o *GetMembershipNativeResponsePointInfo) GetMaxPoints() int64`

GetMaxPoints returns the MaxPoints field if non-nil, zero value otherwise.

### GetMaxPointsOk

`func (o *GetMembershipNativeResponsePointInfo) GetMaxPointsOk() (*int64, bool)`

GetMaxPointsOk returns a tuple with the MaxPoints field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxPoints

`func (o *GetMembershipNativeResponsePointInfo) SetMaxPoints(v int64)`

SetMaxPoints sets MaxPoints field to given value.

### HasMaxPoints

`func (o *GetMembershipNativeResponsePointInfo) HasMaxPoints() bool`

HasMaxPoints returns a boolean if a field has been set.

### GetExpirePoints

`func (o *GetMembershipNativeResponsePointInfo) GetExpirePoints() int64`

GetExpirePoints returns the ExpirePoints field if non-nil, zero value otherwise.

### GetExpirePointsOk

`func (o *GetMembershipNativeResponsePointInfo) GetExpirePointsOk() (*int64, bool)`

GetExpirePointsOk returns a tuple with the ExpirePoints field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpirePoints

`func (o *GetMembershipNativeResponsePointInfo) SetExpirePoints(v int64)`

SetExpirePoints sets ExpirePoints field to given value.

### HasExpirePoints

`func (o *GetMembershipNativeResponsePointInfo) HasExpirePoints() bool`

HasExpirePoints returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


