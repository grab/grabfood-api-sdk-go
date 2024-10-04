# CancelReason

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Code** | Pointer to [**CancelCode**](CancelCode.md) |  | [optional] 
**Reason** | Pointer to **string** | The detailed cancel reason for the specific cancel code. - Items are unavailable &lt;code 1001&gt; - I have too many orders now &lt;code 1002&gt; - My shop is closed &lt;code 1003&gt; - My shop is closing soon &lt;code 1004&gt;  | [optional] 

## Methods

### NewCancelReason

`func NewCancelReason() *CancelReason`

NewCancelReason instantiates a new CancelReason object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCancelReasonWithDefaults

`func NewCancelReasonWithDefaults() *CancelReason`

NewCancelReasonWithDefaults instantiates a new CancelReason object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCode

`func (o *CancelReason) GetCode() CancelCode`

GetCode returns the Code field if non-nil, zero value otherwise.

### GetCodeOk

`func (o *CancelReason) GetCodeOk() (*CancelCode, bool)`

GetCodeOk returns a tuple with the Code field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCode

`func (o *CancelReason) SetCode(v CancelCode)`

SetCode sets Code field to given value.

### HasCode

`func (o *CancelReason) HasCode() bool`

HasCode returns a boolean if a field has been set.

### GetReason

`func (o *CancelReason) GetReason() string`

GetReason returns the Reason field if non-nil, zero value otherwise.

### GetReasonOk

`func (o *CancelReason) GetReasonOk() (*string, bool)`

GetReasonOk returns a tuple with the Reason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReason

`func (o *CancelReason) SetReason(v string)`

SetReason sets Reason field to given value.

### HasReason

`func (o *CancelReason) HasReason() bool`

HasReason returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


