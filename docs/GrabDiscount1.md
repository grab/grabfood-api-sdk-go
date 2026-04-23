# GrabDiscount1

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Code** | Pointer to **string** | discount code. | [optional] 
**Id** | Pointer to **string** | discount id. | [optional] 
**Name** | Pointer to **string** | discount name. | [optional] 
**DeductAmountInMin** | Pointer to **int64** | The total discount amount in minor unit. | [optional] 
**Level** | Pointer to **string** | discount level, eg, order / item level. | [optional] 
**Type** | Pointer to **string** | discount type, eg, Promo / DineOutVoucher / DineOutDiscount. | [optional] 
**MexFundedAmountInMin** | Pointer to **int64** | The mexFundDiscount in minor unit. | [optional] 
**AppliedItemIDs** | Pointer to **[]string** | An array of item IDs that get discount under this grabDiscount. &#x60;null&#x60; if no item applied in this grabDiscount. | [optional] 

## Methods

### NewGrabDiscount1

`func NewGrabDiscount1() *GrabDiscount1`

NewGrabDiscount1 instantiates a new GrabDiscount1 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGrabDiscount1WithDefaults

`func NewGrabDiscount1WithDefaults() *GrabDiscount1`

NewGrabDiscount1WithDefaults instantiates a new GrabDiscount1 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCode

`func (o *GrabDiscount1) GetCode() string`

GetCode returns the Code field if non-nil, zero value otherwise.

### GetCodeOk

`func (o *GrabDiscount1) GetCodeOk() (*string, bool)`

GetCodeOk returns a tuple with the Code field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCode

`func (o *GrabDiscount1) SetCode(v string)`

SetCode sets Code field to given value.

### HasCode

`func (o *GrabDiscount1) HasCode() bool`

HasCode returns a boolean if a field has been set.

### GetId

`func (o *GrabDiscount1) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *GrabDiscount1) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *GrabDiscount1) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *GrabDiscount1) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *GrabDiscount1) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *GrabDiscount1) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *GrabDiscount1) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *GrabDiscount1) HasName() bool`

HasName returns a boolean if a field has been set.

### GetDeductAmountInMin

`func (o *GrabDiscount1) GetDeductAmountInMin() int64`

GetDeductAmountInMin returns the DeductAmountInMin field if non-nil, zero value otherwise.

### GetDeductAmountInMinOk

`func (o *GrabDiscount1) GetDeductAmountInMinOk() (*int64, bool)`

GetDeductAmountInMinOk returns a tuple with the DeductAmountInMin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeductAmountInMin

`func (o *GrabDiscount1) SetDeductAmountInMin(v int64)`

SetDeductAmountInMin sets DeductAmountInMin field to given value.

### HasDeductAmountInMin

`func (o *GrabDiscount1) HasDeductAmountInMin() bool`

HasDeductAmountInMin returns a boolean if a field has been set.

### GetLevel

`func (o *GrabDiscount1) GetLevel() string`

GetLevel returns the Level field if non-nil, zero value otherwise.

### GetLevelOk

`func (o *GrabDiscount1) GetLevelOk() (*string, bool)`

GetLevelOk returns a tuple with the Level field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLevel

`func (o *GrabDiscount1) SetLevel(v string)`

SetLevel sets Level field to given value.

### HasLevel

`func (o *GrabDiscount1) HasLevel() bool`

HasLevel returns a boolean if a field has been set.

### GetType

`func (o *GrabDiscount1) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *GrabDiscount1) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *GrabDiscount1) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *GrabDiscount1) HasType() bool`

HasType returns a boolean if a field has been set.

### GetMexFundedAmountInMin

`func (o *GrabDiscount1) GetMexFundedAmountInMin() int64`

GetMexFundedAmountInMin returns the MexFundedAmountInMin field if non-nil, zero value otherwise.

### GetMexFundedAmountInMinOk

`func (o *GrabDiscount1) GetMexFundedAmountInMinOk() (*int64, bool)`

GetMexFundedAmountInMinOk returns a tuple with the MexFundedAmountInMin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMexFundedAmountInMin

`func (o *GrabDiscount1) SetMexFundedAmountInMin(v int64)`

SetMexFundedAmountInMin sets MexFundedAmountInMin field to given value.

### HasMexFundedAmountInMin

`func (o *GrabDiscount1) HasMexFundedAmountInMin() bool`

HasMexFundedAmountInMin returns a boolean if a field has been set.

### GetAppliedItemIDs

`func (o *GrabDiscount1) GetAppliedItemIDs() []string`

GetAppliedItemIDs returns the AppliedItemIDs field if non-nil, zero value otherwise.

### GetAppliedItemIDsOk

`func (o *GrabDiscount1) GetAppliedItemIDsOk() (*[]string, bool)`

GetAppliedItemIDsOk returns a tuple with the AppliedItemIDs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAppliedItemIDs

`func (o *GrabDiscount1) SetAppliedItemIDs(v []string)`

SetAppliedItemIDs sets AppliedItemIDs field to given value.

### HasAppliedItemIDs

`func (o *GrabDiscount1) HasAppliedItemIDs() bool`

HasAppliedItemIDs returns a boolean if a field has been set.

### SetAppliedItemIDsNil

`func (o *GrabDiscount1) SetAppliedItemIDsNil(b bool)`

 SetAppliedItemIDsNil sets the value for AppliedItemIDs to be an explicit nil

### UnsetAppliedItemIDs
`func (o *GrabDiscount1) UnsetAppliedItemIDs()`

UnsetAppliedItemIDs ensures that no value is present for AppliedItemIDs, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


