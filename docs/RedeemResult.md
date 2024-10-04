# RedeemResult

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Success** | Pointer to **bool** | To indicate whether the dine-in voucher redemption succeeded. | [optional] 
**Code** | Pointer to **string** | The code for the reason of failed redemption. Empty if the &#x60;success&#x60; is true.  * &#x60;VOUCHER_REDEEMED&#x60; - The voucher has already been redeemed. * &#x60;INVALID_STATE&#x60; - The current status of voucher is EXPIRED or REFUNDED. * &#x60;REDEEM_FAILED&#x60; - Internal service error. * &#x60;INVALID_MERCHANT&#x60; - Voucher not applicable for this merchant. * &#x60;INVALID_ID&#x60; - Invalid certificateID.  | [optional] 

## Methods

### NewRedeemResult

`func NewRedeemResult() *RedeemResult`

NewRedeemResult instantiates a new RedeemResult object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRedeemResultWithDefaults

`func NewRedeemResultWithDefaults() *RedeemResult`

NewRedeemResultWithDefaults instantiates a new RedeemResult object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSuccess

`func (o *RedeemResult) GetSuccess() bool`

GetSuccess returns the Success field if non-nil, zero value otherwise.

### GetSuccessOk

`func (o *RedeemResult) GetSuccessOk() (*bool, bool)`

GetSuccessOk returns a tuple with the Success field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSuccess

`func (o *RedeemResult) SetSuccess(v bool)`

SetSuccess sets Success field to given value.

### HasSuccess

`func (o *RedeemResult) HasSuccess() bool`

HasSuccess returns a boolean if a field has been set.

### GetCode

`func (o *RedeemResult) GetCode() string`

GetCode returns the Code field if non-nil, zero value otherwise.

### GetCodeOk

`func (o *RedeemResult) GetCodeOk() (*string, bool)`

GetCodeOk returns a tuple with the Code field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCode

`func (o *RedeemResult) SetCode(v string)`

SetCode sets Code field to given value.

### HasCode

`func (o *RedeemResult) HasCode() bool`

HasCode returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


