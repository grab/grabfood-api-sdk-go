# Payment

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Method** | Pointer to **string** | payment name, eg, visa, masterCard. | [optional] 
**FundingType** | Pointer to **string** | fundingType name, eg, credit, debit. | [optional] 
**AmountInMin** | Pointer to **int64** | amount paid by this payment method. | [optional] 

## Methods

### NewPayment

`func NewPayment() *Payment`

NewPayment instantiates a new Payment object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPaymentWithDefaults

`func NewPaymentWithDefaults() *Payment`

NewPaymentWithDefaults instantiates a new Payment object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMethod

`func (o *Payment) GetMethod() string`

GetMethod returns the Method field if non-nil, zero value otherwise.

### GetMethodOk

`func (o *Payment) GetMethodOk() (*string, bool)`

GetMethodOk returns a tuple with the Method field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMethod

`func (o *Payment) SetMethod(v string)`

SetMethod sets Method field to given value.

### HasMethod

`func (o *Payment) HasMethod() bool`

HasMethod returns a boolean if a field has been set.

### GetFundingType

`func (o *Payment) GetFundingType() string`

GetFundingType returns the FundingType field if non-nil, zero value otherwise.

### GetFundingTypeOk

`func (o *Payment) GetFundingTypeOk() (*string, bool)`

GetFundingTypeOk returns a tuple with the FundingType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFundingType

`func (o *Payment) SetFundingType(v string)`

SetFundingType sets FundingType field to given value.

### HasFundingType

`func (o *Payment) HasFundingType() bool`

HasFundingType returns a boolean if a field has been set.

### GetAmountInMin

`func (o *Payment) GetAmountInMin() int64`

GetAmountInMin returns the AmountInMin field if non-nil, zero value otherwise.

### GetAmountInMinOk

`func (o *Payment) GetAmountInMinOk() (*int64, bool)`

GetAmountInMinOk returns a tuple with the AmountInMin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmountInMin

`func (o *Payment) SetAmountInMin(v int64)`

SetAmountInMin sets AmountInMin field to given value.

### HasAmountInMin

`func (o *Payment) HasAmountInMin() bool`

HasAmountInMin returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


