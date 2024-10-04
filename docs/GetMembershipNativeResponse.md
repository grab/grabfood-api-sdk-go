# GetMembershipNativeResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**MembershipStatus** | Pointer to **string** | Status of the memberID. | [optional] 
**PointInfo** | Pointer to [**GetMembershipNativeResponsePointInfo**](GetMembershipNativeResponsePointInfo.md) |  | [optional] 
**PointsExpireDate** | Pointer to **string** | Earliest points expiry date. In &#x60;yyyy-mm-dd&#x60; format | [optional] 

## Methods

### NewGetMembershipNativeResponse

`func NewGetMembershipNativeResponse() *GetMembershipNativeResponse`

NewGetMembershipNativeResponse instantiates a new GetMembershipNativeResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetMembershipNativeResponseWithDefaults

`func NewGetMembershipNativeResponseWithDefaults() *GetMembershipNativeResponse`

NewGetMembershipNativeResponseWithDefaults instantiates a new GetMembershipNativeResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMembershipStatus

`func (o *GetMembershipNativeResponse) GetMembershipStatus() string`

GetMembershipStatus returns the MembershipStatus field if non-nil, zero value otherwise.

### GetMembershipStatusOk

`func (o *GetMembershipNativeResponse) GetMembershipStatusOk() (*string, bool)`

GetMembershipStatusOk returns a tuple with the MembershipStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMembershipStatus

`func (o *GetMembershipNativeResponse) SetMembershipStatus(v string)`

SetMembershipStatus sets MembershipStatus field to given value.

### HasMembershipStatus

`func (o *GetMembershipNativeResponse) HasMembershipStatus() bool`

HasMembershipStatus returns a boolean if a field has been set.

### GetPointInfo

`func (o *GetMembershipNativeResponse) GetPointInfo() GetMembershipNativeResponsePointInfo`

GetPointInfo returns the PointInfo field if non-nil, zero value otherwise.

### GetPointInfoOk

`func (o *GetMembershipNativeResponse) GetPointInfoOk() (*GetMembershipNativeResponsePointInfo, bool)`

GetPointInfoOk returns a tuple with the PointInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPointInfo

`func (o *GetMembershipNativeResponse) SetPointInfo(v GetMembershipNativeResponsePointInfo)`

SetPointInfo sets PointInfo field to given value.

### HasPointInfo

`func (o *GetMembershipNativeResponse) HasPointInfo() bool`

HasPointInfo returns a boolean if a field has been set.

### GetPointsExpireDate

`func (o *GetMembershipNativeResponse) GetPointsExpireDate() string`

GetPointsExpireDate returns the PointsExpireDate field if non-nil, zero value otherwise.

### GetPointsExpireDateOk

`func (o *GetMembershipNativeResponse) GetPointsExpireDateOk() (*string, bool)`

GetPointsExpireDateOk returns a tuple with the PointsExpireDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPointsExpireDate

`func (o *GetMembershipNativeResponse) SetPointsExpireDate(v string)`

SetPointsExpireDate sets PointsExpireDate field to given value.

### HasPointsExpireDate

`func (o *GetMembershipNativeResponse) HasPointsExpireDate() bool`

HasPointsExpireDate returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


