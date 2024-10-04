# OrderFeatureFlags

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**OrderAcceptedType** | **string** | The acceptance type for the order. Refer to FAQs for more details about [orderAcceptedType](#section/Order/How-do-I-identify-if-a-particular-order-is-auto-or-manual-acceptance).  | 
**OrderType** | **string** | The type of order.  | 
**IsMexEditOrder** | Pointer to **bool** | A boolean value that indicates if the order is edited.  | [optional] 

## Methods

### NewOrderFeatureFlags

`func NewOrderFeatureFlags(orderAcceptedType string, orderType string, ) *OrderFeatureFlags`

NewOrderFeatureFlags instantiates a new OrderFeatureFlags object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOrderFeatureFlagsWithDefaults

`func NewOrderFeatureFlagsWithDefaults() *OrderFeatureFlags`

NewOrderFeatureFlagsWithDefaults instantiates a new OrderFeatureFlags object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetOrderAcceptedType

`func (o *OrderFeatureFlags) GetOrderAcceptedType() string`

GetOrderAcceptedType returns the OrderAcceptedType field if non-nil, zero value otherwise.

### GetOrderAcceptedTypeOk

`func (o *OrderFeatureFlags) GetOrderAcceptedTypeOk() (*string, bool)`

GetOrderAcceptedTypeOk returns a tuple with the OrderAcceptedType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrderAcceptedType

`func (o *OrderFeatureFlags) SetOrderAcceptedType(v string)`

SetOrderAcceptedType sets OrderAcceptedType field to given value.


### GetOrderType

`func (o *OrderFeatureFlags) GetOrderType() string`

GetOrderType returns the OrderType field if non-nil, zero value otherwise.

### GetOrderTypeOk

`func (o *OrderFeatureFlags) GetOrderTypeOk() (*string, bool)`

GetOrderTypeOk returns a tuple with the OrderType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrderType

`func (o *OrderFeatureFlags) SetOrderType(v string)`

SetOrderType sets OrderType field to given value.


### GetIsMexEditOrder

`func (o *OrderFeatureFlags) GetIsMexEditOrder() bool`

GetIsMexEditOrder returns the IsMexEditOrder field if non-nil, zero value otherwise.

### GetIsMexEditOrderOk

`func (o *OrderFeatureFlags) GetIsMexEditOrderOk() (*bool, bool)`

GetIsMexEditOrderOk returns a tuple with the IsMexEditOrder field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsMexEditOrder

`func (o *OrderFeatureFlags) SetIsMexEditOrder(v bool)`

SetIsMexEditOrder sets IsMexEditOrder field to given value.

### HasIsMexEditOrder

`func (o *OrderFeatureFlags) HasIsMexEditOrder() bool`

HasIsMexEditOrder returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


