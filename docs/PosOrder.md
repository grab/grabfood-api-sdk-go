# PosOrder

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**OrderID** | Pointer to **string** | The long orderID in grab system. | [optional] 
**PartnerOrderID** | Pointer to **string** | The orderID in pos system. | [optional] 
**MerchantID** | Pointer to **string** | The merchant&#39;s ID is the one in GrabFood&#39;s database. | [optional] 
**PartnerMerchantID** | Pointer to **string** | The merchant ID in pos system. | [optional] 
**OrderTime** | Pointer to **time.Time** | The UTC time that a consumer places the order, based on ISO_8601/RFC3339. | [optional] 
**OrderState** | Pointer to **string** | The order state in POS system, eg, COMPLETED. | [optional] 
**Currency** | Pointer to [**Currency**](Currency.md) |  | [optional] 
**Items** | Pointer to [**[]POSItem**](POSItem.md) | The ordered items in an array of JSON Object.  | [optional] 
**Price** | Pointer to [**PosPriceDetails**](PosPriceDetails.md) |  | [optional] 
**DineIn** | Pointer to [**NullableDineIn**](DineIn.md) |  | [optional] 
**Payments** | Pointer to [**[]Payment**](Payment.md) | An array of payment objects. &#x60;null&#x60; when there is no payment info from pos. This is only applicable for STO order | [optional] 

## Methods

### NewPosOrder

`func NewPosOrder() *PosOrder`

NewPosOrder instantiates a new PosOrder object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPosOrderWithDefaults

`func NewPosOrderWithDefaults() *PosOrder`

NewPosOrderWithDefaults instantiates a new PosOrder object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetOrderID

`func (o *PosOrder) GetOrderID() string`

GetOrderID returns the OrderID field if non-nil, zero value otherwise.

### GetOrderIDOk

`func (o *PosOrder) GetOrderIDOk() (*string, bool)`

GetOrderIDOk returns a tuple with the OrderID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrderID

`func (o *PosOrder) SetOrderID(v string)`

SetOrderID sets OrderID field to given value.

### HasOrderID

`func (o *PosOrder) HasOrderID() bool`

HasOrderID returns a boolean if a field has been set.

### GetPartnerOrderID

`func (o *PosOrder) GetPartnerOrderID() string`

GetPartnerOrderID returns the PartnerOrderID field if non-nil, zero value otherwise.

### GetPartnerOrderIDOk

`func (o *PosOrder) GetPartnerOrderIDOk() (*string, bool)`

GetPartnerOrderIDOk returns a tuple with the PartnerOrderID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPartnerOrderID

`func (o *PosOrder) SetPartnerOrderID(v string)`

SetPartnerOrderID sets PartnerOrderID field to given value.

### HasPartnerOrderID

`func (o *PosOrder) HasPartnerOrderID() bool`

HasPartnerOrderID returns a boolean if a field has been set.

### GetMerchantID

`func (o *PosOrder) GetMerchantID() string`

GetMerchantID returns the MerchantID field if non-nil, zero value otherwise.

### GetMerchantIDOk

`func (o *PosOrder) GetMerchantIDOk() (*string, bool)`

GetMerchantIDOk returns a tuple with the MerchantID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMerchantID

`func (o *PosOrder) SetMerchantID(v string)`

SetMerchantID sets MerchantID field to given value.

### HasMerchantID

`func (o *PosOrder) HasMerchantID() bool`

HasMerchantID returns a boolean if a field has been set.

### GetPartnerMerchantID

`func (o *PosOrder) GetPartnerMerchantID() string`

GetPartnerMerchantID returns the PartnerMerchantID field if non-nil, zero value otherwise.

### GetPartnerMerchantIDOk

`func (o *PosOrder) GetPartnerMerchantIDOk() (*string, bool)`

GetPartnerMerchantIDOk returns a tuple with the PartnerMerchantID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPartnerMerchantID

`func (o *PosOrder) SetPartnerMerchantID(v string)`

SetPartnerMerchantID sets PartnerMerchantID field to given value.

### HasPartnerMerchantID

`func (o *PosOrder) HasPartnerMerchantID() bool`

HasPartnerMerchantID returns a boolean if a field has been set.

### GetOrderTime

`func (o *PosOrder) GetOrderTime() time.Time`

GetOrderTime returns the OrderTime field if non-nil, zero value otherwise.

### GetOrderTimeOk

`func (o *PosOrder) GetOrderTimeOk() (*time.Time, bool)`

GetOrderTimeOk returns a tuple with the OrderTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrderTime

`func (o *PosOrder) SetOrderTime(v time.Time)`

SetOrderTime sets OrderTime field to given value.

### HasOrderTime

`func (o *PosOrder) HasOrderTime() bool`

HasOrderTime returns a boolean if a field has been set.

### GetOrderState

`func (o *PosOrder) GetOrderState() string`

GetOrderState returns the OrderState field if non-nil, zero value otherwise.

### GetOrderStateOk

`func (o *PosOrder) GetOrderStateOk() (*string, bool)`

GetOrderStateOk returns a tuple with the OrderState field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrderState

`func (o *PosOrder) SetOrderState(v string)`

SetOrderState sets OrderState field to given value.

### HasOrderState

`func (o *PosOrder) HasOrderState() bool`

HasOrderState returns a boolean if a field has been set.

### GetCurrency

`func (o *PosOrder) GetCurrency() Currency`

GetCurrency returns the Currency field if non-nil, zero value otherwise.

### GetCurrencyOk

`func (o *PosOrder) GetCurrencyOk() (*Currency, bool)`

GetCurrencyOk returns a tuple with the Currency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrency

`func (o *PosOrder) SetCurrency(v Currency)`

SetCurrency sets Currency field to given value.

### HasCurrency

`func (o *PosOrder) HasCurrency() bool`

HasCurrency returns a boolean if a field has been set.

### GetItems

`func (o *PosOrder) GetItems() []POSItem`

GetItems returns the Items field if non-nil, zero value otherwise.

### GetItemsOk

`func (o *PosOrder) GetItemsOk() (*[]POSItem, bool)`

GetItemsOk returns a tuple with the Items field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItems

`func (o *PosOrder) SetItems(v []POSItem)`

SetItems sets Items field to given value.

### HasItems

`func (o *PosOrder) HasItems() bool`

HasItems returns a boolean if a field has been set.

### GetPrice

`func (o *PosOrder) GetPrice() PosPriceDetails`

GetPrice returns the Price field if non-nil, zero value otherwise.

### GetPriceOk

`func (o *PosOrder) GetPriceOk() (*PosPriceDetails, bool)`

GetPriceOk returns a tuple with the Price field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrice

`func (o *PosOrder) SetPrice(v PosPriceDetails)`

SetPrice sets Price field to given value.

### HasPrice

`func (o *PosOrder) HasPrice() bool`

HasPrice returns a boolean if a field has been set.

### GetDineIn

`func (o *PosOrder) GetDineIn() DineIn`

GetDineIn returns the DineIn field if non-nil, zero value otherwise.

### GetDineInOk

`func (o *PosOrder) GetDineInOk() (*DineIn, bool)`

GetDineInOk returns a tuple with the DineIn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDineIn

`func (o *PosOrder) SetDineIn(v DineIn)`

SetDineIn sets DineIn field to given value.

### HasDineIn

`func (o *PosOrder) HasDineIn() bool`

HasDineIn returns a boolean if a field has been set.

### SetDineInNil

`func (o *PosOrder) SetDineInNil(b bool)`

 SetDineInNil sets the value for DineIn to be an explicit nil

### UnsetDineIn
`func (o *PosOrder) UnsetDineIn()`

UnsetDineIn ensures that no value is present for DineIn, not even an explicit nil
### GetPayments

`func (o *PosOrder) GetPayments() []Payment`

GetPayments returns the Payments field if non-nil, zero value otherwise.

### GetPaymentsOk

`func (o *PosOrder) GetPaymentsOk() (*[]Payment, bool)`

GetPaymentsOk returns a tuple with the Payments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPayments

`func (o *PosOrder) SetPayments(v []Payment)`

SetPayments sets Payments field to given value.

### HasPayments

`func (o *PosOrder) HasPayments() bool`

HasPayments returns a boolean if a field has been set.

### SetPaymentsNil

`func (o *PosOrder) SetPaymentsNil(b bool)`

 SetPaymentsNil sets the value for Payments to be an explicit nil

### UnsetPayments
`func (o *PosOrder) UnsetPayments()`

UnsetPayments ensures that no value is present for Payments, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


