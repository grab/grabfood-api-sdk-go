# SyncPOSOrderRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Action** | **string** | This action indicates the target action POS wants to do, eg, BILL_GENERATED, COMPLETED. | 
**Order** | [**PosOrder**](PosOrder.md) |  | 

## Methods

### NewSyncPOSOrderRequest

`func NewSyncPOSOrderRequest(action string, order PosOrder, ) *SyncPOSOrderRequest`

NewSyncPOSOrderRequest instantiates a new SyncPOSOrderRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSyncPOSOrderRequestWithDefaults

`func NewSyncPOSOrderRequestWithDefaults() *SyncPOSOrderRequest`

NewSyncPOSOrderRequestWithDefaults instantiates a new SyncPOSOrderRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAction

`func (o *SyncPOSOrderRequest) GetAction() string`

GetAction returns the Action field if non-nil, zero value otherwise.

### GetActionOk

`func (o *SyncPOSOrderRequest) GetActionOk() (*string, bool)`

GetActionOk returns a tuple with the Action field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAction

`func (o *SyncPOSOrderRequest) SetAction(v string)`

SetAction sets Action field to given value.


### GetOrder

`func (o *SyncPOSOrderRequest) GetOrder() PosOrder`

GetOrder returns the Order field if non-nil, zero value otherwise.

### GetOrderOk

`func (o *SyncPOSOrderRequest) GetOrderOk() (*PosOrder, bool)`

GetOrderOk returns a tuple with the Order field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrder

`func (o *SyncPOSOrderRequest) SetOrder(v PosOrder)`

SetOrder sets Order field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


