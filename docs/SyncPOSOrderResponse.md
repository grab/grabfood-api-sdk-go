# SyncPOSOrderResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**OrderID** | Pointer to **string** | The orderID in grab system. | [optional] 
**PaybillQRCode** | Pointer to **string** | The paybill QR Code. | [optional] 

## Methods

### NewSyncPOSOrderResponse

`func NewSyncPOSOrderResponse() *SyncPOSOrderResponse`

NewSyncPOSOrderResponse instantiates a new SyncPOSOrderResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSyncPOSOrderResponseWithDefaults

`func NewSyncPOSOrderResponseWithDefaults() *SyncPOSOrderResponse`

NewSyncPOSOrderResponseWithDefaults instantiates a new SyncPOSOrderResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetOrderID

`func (o *SyncPOSOrderResponse) GetOrderID() string`

GetOrderID returns the OrderID field if non-nil, zero value otherwise.

### GetOrderIDOk

`func (o *SyncPOSOrderResponse) GetOrderIDOk() (*string, bool)`

GetOrderIDOk returns a tuple with the OrderID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrderID

`func (o *SyncPOSOrderResponse) SetOrderID(v string)`

SetOrderID sets OrderID field to given value.

### HasOrderID

`func (o *SyncPOSOrderResponse) HasOrderID() bool`

HasOrderID returns a boolean if a field has been set.

### GetPaybillQRCode

`func (o *SyncPOSOrderResponse) GetPaybillQRCode() string`

GetPaybillQRCode returns the PaybillQRCode field if non-nil, zero value otherwise.

### GetPaybillQRCodeOk

`func (o *SyncPOSOrderResponse) GetPaybillQRCodeOk() (*string, bool)`

GetPaybillQRCodeOk returns a tuple with the PaybillQRCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaybillQRCode

`func (o *SyncPOSOrderResponse) SetPaybillQRCode(v string)`

SetPaybillQRCode sets PaybillQRCode field to given value.

### HasPaybillQRCode

`func (o *SyncPOSOrderResponse) HasPaybillQRCode() bool`

HasPaybillQRCode returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


