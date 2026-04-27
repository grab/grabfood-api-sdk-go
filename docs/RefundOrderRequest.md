# RefundOrderRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**OrderID** | **string** | This order ID in grab system. | 
**MerchantID** | **string** | This merchant ID in grab system. | 
**IsFullRefund** | Pointer to **bool** | currently we only support fully refund. | [optional] 
**RefundAmountInMin** | Pointer to **int64** | The total amount the POS want to refund for STO order. | [optional] 

## Methods

### NewRefundOrderRequest

`func NewRefundOrderRequest(orderID string, merchantID string, ) *RefundOrderRequest`

NewRefundOrderRequest instantiates a new RefundOrderRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRefundOrderRequestWithDefaults

`func NewRefundOrderRequestWithDefaults() *RefundOrderRequest`

NewRefundOrderRequestWithDefaults instantiates a new RefundOrderRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetOrderID

`func (o *RefundOrderRequest) GetOrderID() string`

GetOrderID returns the OrderID field if non-nil, zero value otherwise.

### GetOrderIDOk

`func (o *RefundOrderRequest) GetOrderIDOk() (*string, bool)`

GetOrderIDOk returns a tuple with the OrderID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrderID

`func (o *RefundOrderRequest) SetOrderID(v string)`

SetOrderID sets OrderID field to given value.


### GetMerchantID

`func (o *RefundOrderRequest) GetMerchantID() string`

GetMerchantID returns the MerchantID field if non-nil, zero value otherwise.

### GetMerchantIDOk

`func (o *RefundOrderRequest) GetMerchantIDOk() (*string, bool)`

GetMerchantIDOk returns a tuple with the MerchantID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMerchantID

`func (o *RefundOrderRequest) SetMerchantID(v string)`

SetMerchantID sets MerchantID field to given value.


### GetIsFullRefund

`func (o *RefundOrderRequest) GetIsFullRefund() bool`

GetIsFullRefund returns the IsFullRefund field if non-nil, zero value otherwise.

### GetIsFullRefundOk

`func (o *RefundOrderRequest) GetIsFullRefundOk() (*bool, bool)`

GetIsFullRefundOk returns a tuple with the IsFullRefund field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsFullRefund

`func (o *RefundOrderRequest) SetIsFullRefund(v bool)`

SetIsFullRefund sets IsFullRefund field to given value.

### HasIsFullRefund

`func (o *RefundOrderRequest) HasIsFullRefund() bool`

HasIsFullRefund returns a boolean if a field has been set.

### GetRefundAmountInMin

`func (o *RefundOrderRequest) GetRefundAmountInMin() int64`

GetRefundAmountInMin returns the RefundAmountInMin field if non-nil, zero value otherwise.

### GetRefundAmountInMinOk

`func (o *RefundOrderRequest) GetRefundAmountInMinOk() (*int64, bool)`

GetRefundAmountInMinOk returns a tuple with the RefundAmountInMin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRefundAmountInMin

`func (o *RefundOrderRequest) SetRefundAmountInMin(v int64)`

SetRefundAmountInMin sets RefundAmountInMin field to given value.

### HasRefundAmountInMin

`func (o *RefundOrderRequest) HasRefundAmountInMin() bool`

HasRefundAmountInMin returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


