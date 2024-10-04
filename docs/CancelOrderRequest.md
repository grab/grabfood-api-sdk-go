# CancelOrderRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**OrderID** | **string** | The order&#39;s ID that is returned from GrabFood. Refer to FAQs for more details about [orderID and shortOrderNumber](#section/Order/What&#39;s-the-difference-between-orderID-and-shortOrderNumber). | 
**MerchantID** | **string** | The merchant&#39;s ID that is in GrabFood&#39;s database. | 
**CancelCode** | [**CancelCode**](CancelCode.md) |  | 

## Methods

### NewCancelOrderRequest

`func NewCancelOrderRequest(orderID string, merchantID string, cancelCode CancelCode, ) *CancelOrderRequest`

NewCancelOrderRequest instantiates a new CancelOrderRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCancelOrderRequestWithDefaults

`func NewCancelOrderRequestWithDefaults() *CancelOrderRequest`

NewCancelOrderRequestWithDefaults instantiates a new CancelOrderRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetOrderID

`func (o *CancelOrderRequest) GetOrderID() string`

GetOrderID returns the OrderID field if non-nil, zero value otherwise.

### GetOrderIDOk

`func (o *CancelOrderRequest) GetOrderIDOk() (*string, bool)`

GetOrderIDOk returns a tuple with the OrderID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrderID

`func (o *CancelOrderRequest) SetOrderID(v string)`

SetOrderID sets OrderID field to given value.


### GetMerchantID

`func (o *CancelOrderRequest) GetMerchantID() string`

GetMerchantID returns the MerchantID field if non-nil, zero value otherwise.

### GetMerchantIDOk

`func (o *CancelOrderRequest) GetMerchantIDOk() (*string, bool)`

GetMerchantIDOk returns a tuple with the MerchantID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMerchantID

`func (o *CancelOrderRequest) SetMerchantID(v string)`

SetMerchantID sets MerchantID field to given value.


### GetCancelCode

`func (o *CancelOrderRequest) GetCancelCode() CancelCode`

GetCancelCode returns the CancelCode field if non-nil, zero value otherwise.

### GetCancelCodeOk

`func (o *CancelOrderRequest) GetCancelCodeOk() (*CancelCode, bool)`

GetCancelCodeOk returns a tuple with the CancelCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCancelCode

`func (o *CancelOrderRequest) SetCancelCode(v CancelCode)`

SetCancelCode sets CancelCode field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


