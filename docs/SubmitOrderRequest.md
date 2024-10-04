# SubmitOrderRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**OrderID** | **string** | The order&#39;s ID that is returned from GrabFood. Refer to FAQs for more details about [orderID and shortOrderNumber](#section/Order/What&#39;s-the-difference-between-orderID-and-shortOrderNumber). | 
**ShortOrderNumber** | **string** | The GrabFood short order number. This is unique for each merchant per day. Refer to FAQs for more details about [orderID and shortOrderNumber](#section/Order/What&#39;s-the-difference-between-orderID-and-shortOrderNumber). | 
**MerchantID** | **string** | The merchant&#39;s ID that is in GrabFood&#39;s database. | 
**PartnerMerchantID** | Pointer to **string** | The merchant&#39;s ID that is on the partner&#39;s database. | [optional] 
**PaymentType** | **string** | The payment method used. Refer to FAQs for more details about [paymentType](#section/Order/Does-the-paymentType-affect-partners). | 
**Cutlery** | **bool** | The boolean value to indicate whether cutlery are needed or not. Refer to FAQs for more details about [cutlery](#section/Order/What-do-the-true-or-false-values-mean-for-cutlery). | 
**OrderTime** | **string** | The UTC time that a consumer places the order, based on ISO_8601/RFC3339. | 
**SubmitTime** | Pointer to **time.Time** | The order submit time, based on ISO_8601/RFC3339. Only present in the [List Orders](#tag/list-order) response. | [optional] 
**CompleteTime** | Pointer to **time.Time** | The order complete time, based on ISO_8601/RFC3339. Only present in the [List Orders](#tag/list-order) response. | [optional] 
**ScheduledTime** | Pointer to **string** | The order scheduled time, based on ISO_8601/RFC3339. Empty for non-scheduled orders. | [optional] 
**OrderState** | Pointer to **string** | The state of the order. Only present in the [List Orders](#tag/list-order) response. Refer to [Order States](#section/Order-states). | [optional] 
**Currency** | [**Currency**](Currency.md) |  | 
**FeatureFlags** | [**OrderFeatureFlags**](OrderFeatureFlags.md) |  | 
**Items** | [**[]OrderItem**](OrderItem.md) | The items in an array of JSON Object. Refer to [Items](#items) for more information. | 
**Campaigns** | Pointer to [**[]OrderCampaign**](OrderCampaign.md) | The campaigns that are applicable for the order.&#x60;null&#x60; when there is no campaign applied.  | [optional] 
**Promos** | Pointer to [**[]OrderPromo**](OrderPromo.md) | An array of promotion objects. Only promotions that are funded by merchants will be sent. | [optional] 
**Price** | [**OrderPrice**](OrderPrice.md) |  | 
**DineIn** | Pointer to [**DineIn**](DineIn.md) |  | [optional] 
**Receiver** | Pointer to [**Receiver**](Receiver.md) |  | [optional] 
**OrderReadyEstimation** | Pointer to [**OrderReadyEstimation**](OrderReadyEstimation.md) |  | [optional] 
**MembershipID** | Pointer to **string** | Membership ID for loyalty project. Only present for loyalty program partners. | [optional] 

## Methods

### NewSubmitOrderRequest

`func NewSubmitOrderRequest(orderID string, shortOrderNumber string, merchantID string, paymentType string, cutlery bool, orderTime string, currency Currency, featureFlags OrderFeatureFlags, items []OrderItem, price OrderPrice, ) *SubmitOrderRequest`

NewSubmitOrderRequest instantiates a new SubmitOrderRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSubmitOrderRequestWithDefaults

`func NewSubmitOrderRequestWithDefaults() *SubmitOrderRequest`

NewSubmitOrderRequestWithDefaults instantiates a new SubmitOrderRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetOrderID

`func (o *SubmitOrderRequest) GetOrderID() string`

GetOrderID returns the OrderID field if non-nil, zero value otherwise.

### GetOrderIDOk

`func (o *SubmitOrderRequest) GetOrderIDOk() (*string, bool)`

GetOrderIDOk returns a tuple with the OrderID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrderID

`func (o *SubmitOrderRequest) SetOrderID(v string)`

SetOrderID sets OrderID field to given value.


### GetShortOrderNumber

`func (o *SubmitOrderRequest) GetShortOrderNumber() string`

GetShortOrderNumber returns the ShortOrderNumber field if non-nil, zero value otherwise.

### GetShortOrderNumberOk

`func (o *SubmitOrderRequest) GetShortOrderNumberOk() (*string, bool)`

GetShortOrderNumberOk returns a tuple with the ShortOrderNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShortOrderNumber

`func (o *SubmitOrderRequest) SetShortOrderNumber(v string)`

SetShortOrderNumber sets ShortOrderNumber field to given value.


### GetMerchantID

`func (o *SubmitOrderRequest) GetMerchantID() string`

GetMerchantID returns the MerchantID field if non-nil, zero value otherwise.

### GetMerchantIDOk

`func (o *SubmitOrderRequest) GetMerchantIDOk() (*string, bool)`

GetMerchantIDOk returns a tuple with the MerchantID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMerchantID

`func (o *SubmitOrderRequest) SetMerchantID(v string)`

SetMerchantID sets MerchantID field to given value.


### GetPartnerMerchantID

`func (o *SubmitOrderRequest) GetPartnerMerchantID() string`

GetPartnerMerchantID returns the PartnerMerchantID field if non-nil, zero value otherwise.

### GetPartnerMerchantIDOk

`func (o *SubmitOrderRequest) GetPartnerMerchantIDOk() (*string, bool)`

GetPartnerMerchantIDOk returns a tuple with the PartnerMerchantID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPartnerMerchantID

`func (o *SubmitOrderRequest) SetPartnerMerchantID(v string)`

SetPartnerMerchantID sets PartnerMerchantID field to given value.

### HasPartnerMerchantID

`func (o *SubmitOrderRequest) HasPartnerMerchantID() bool`

HasPartnerMerchantID returns a boolean if a field has been set.

### GetPaymentType

`func (o *SubmitOrderRequest) GetPaymentType() string`

GetPaymentType returns the PaymentType field if non-nil, zero value otherwise.

### GetPaymentTypeOk

`func (o *SubmitOrderRequest) GetPaymentTypeOk() (*string, bool)`

GetPaymentTypeOk returns a tuple with the PaymentType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaymentType

`func (o *SubmitOrderRequest) SetPaymentType(v string)`

SetPaymentType sets PaymentType field to given value.


### GetCutlery

`func (o *SubmitOrderRequest) GetCutlery() bool`

GetCutlery returns the Cutlery field if non-nil, zero value otherwise.

### GetCutleryOk

`func (o *SubmitOrderRequest) GetCutleryOk() (*bool, bool)`

GetCutleryOk returns a tuple with the Cutlery field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCutlery

`func (o *SubmitOrderRequest) SetCutlery(v bool)`

SetCutlery sets Cutlery field to given value.


### GetOrderTime

`func (o *SubmitOrderRequest) GetOrderTime() string`

GetOrderTime returns the OrderTime field if non-nil, zero value otherwise.

### GetOrderTimeOk

`func (o *SubmitOrderRequest) GetOrderTimeOk() (*string, bool)`

GetOrderTimeOk returns a tuple with the OrderTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrderTime

`func (o *SubmitOrderRequest) SetOrderTime(v string)`

SetOrderTime sets OrderTime field to given value.


### GetSubmitTime

`func (o *SubmitOrderRequest) GetSubmitTime() time.Time`

GetSubmitTime returns the SubmitTime field if non-nil, zero value otherwise.

### GetSubmitTimeOk

`func (o *SubmitOrderRequest) GetSubmitTimeOk() (*time.Time, bool)`

GetSubmitTimeOk returns a tuple with the SubmitTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubmitTime

`func (o *SubmitOrderRequest) SetSubmitTime(v time.Time)`

SetSubmitTime sets SubmitTime field to given value.

### HasSubmitTime

`func (o *SubmitOrderRequest) HasSubmitTime() bool`

HasSubmitTime returns a boolean if a field has been set.

### GetCompleteTime

`func (o *SubmitOrderRequest) GetCompleteTime() time.Time`

GetCompleteTime returns the CompleteTime field if non-nil, zero value otherwise.

### GetCompleteTimeOk

`func (o *SubmitOrderRequest) GetCompleteTimeOk() (*time.Time, bool)`

GetCompleteTimeOk returns a tuple with the CompleteTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompleteTime

`func (o *SubmitOrderRequest) SetCompleteTime(v time.Time)`

SetCompleteTime sets CompleteTime field to given value.

### HasCompleteTime

`func (o *SubmitOrderRequest) HasCompleteTime() bool`

HasCompleteTime returns a boolean if a field has been set.

### GetScheduledTime

`func (o *SubmitOrderRequest) GetScheduledTime() string`

GetScheduledTime returns the ScheduledTime field if non-nil, zero value otherwise.

### GetScheduledTimeOk

`func (o *SubmitOrderRequest) GetScheduledTimeOk() (*string, bool)`

GetScheduledTimeOk returns a tuple with the ScheduledTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScheduledTime

`func (o *SubmitOrderRequest) SetScheduledTime(v string)`

SetScheduledTime sets ScheduledTime field to given value.

### HasScheduledTime

`func (o *SubmitOrderRequest) HasScheduledTime() bool`

HasScheduledTime returns a boolean if a field has been set.

### GetOrderState

`func (o *SubmitOrderRequest) GetOrderState() string`

GetOrderState returns the OrderState field if non-nil, zero value otherwise.

### GetOrderStateOk

`func (o *SubmitOrderRequest) GetOrderStateOk() (*string, bool)`

GetOrderStateOk returns a tuple with the OrderState field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrderState

`func (o *SubmitOrderRequest) SetOrderState(v string)`

SetOrderState sets OrderState field to given value.

### HasOrderState

`func (o *SubmitOrderRequest) HasOrderState() bool`

HasOrderState returns a boolean if a field has been set.

### GetCurrency

`func (o *SubmitOrderRequest) GetCurrency() Currency`

GetCurrency returns the Currency field if non-nil, zero value otherwise.

### GetCurrencyOk

`func (o *SubmitOrderRequest) GetCurrencyOk() (*Currency, bool)`

GetCurrencyOk returns a tuple with the Currency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrency

`func (o *SubmitOrderRequest) SetCurrency(v Currency)`

SetCurrency sets Currency field to given value.


### GetFeatureFlags

`func (o *SubmitOrderRequest) GetFeatureFlags() OrderFeatureFlags`

GetFeatureFlags returns the FeatureFlags field if non-nil, zero value otherwise.

### GetFeatureFlagsOk

`func (o *SubmitOrderRequest) GetFeatureFlagsOk() (*OrderFeatureFlags, bool)`

GetFeatureFlagsOk returns a tuple with the FeatureFlags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFeatureFlags

`func (o *SubmitOrderRequest) SetFeatureFlags(v OrderFeatureFlags)`

SetFeatureFlags sets FeatureFlags field to given value.


### GetItems

`func (o *SubmitOrderRequest) GetItems() []OrderItem`

GetItems returns the Items field if non-nil, zero value otherwise.

### GetItemsOk

`func (o *SubmitOrderRequest) GetItemsOk() (*[]OrderItem, bool)`

GetItemsOk returns a tuple with the Items field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItems

`func (o *SubmitOrderRequest) SetItems(v []OrderItem)`

SetItems sets Items field to given value.


### GetCampaigns

`func (o *SubmitOrderRequest) GetCampaigns() []OrderCampaign`

GetCampaigns returns the Campaigns field if non-nil, zero value otherwise.

### GetCampaignsOk

`func (o *SubmitOrderRequest) GetCampaignsOk() (*[]OrderCampaign, bool)`

GetCampaignsOk returns a tuple with the Campaigns field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCampaigns

`func (o *SubmitOrderRequest) SetCampaigns(v []OrderCampaign)`

SetCampaigns sets Campaigns field to given value.

### HasCampaigns

`func (o *SubmitOrderRequest) HasCampaigns() bool`

HasCampaigns returns a boolean if a field has been set.

### SetCampaignsNil

`func (o *SubmitOrderRequest) SetCampaignsNil(b bool)`

 SetCampaignsNil sets the value for Campaigns to be an explicit nil

### UnsetCampaigns
`func (o *SubmitOrderRequest) UnsetCampaigns()`

UnsetCampaigns ensures that no value is present for Campaigns, not even an explicit nil
### GetPromos

`func (o *SubmitOrderRequest) GetPromos() []OrderPromo`

GetPromos returns the Promos field if non-nil, zero value otherwise.

### GetPromosOk

`func (o *SubmitOrderRequest) GetPromosOk() (*[]OrderPromo, bool)`

GetPromosOk returns a tuple with the Promos field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPromos

`func (o *SubmitOrderRequest) SetPromos(v []OrderPromo)`

SetPromos sets Promos field to given value.

### HasPromos

`func (o *SubmitOrderRequest) HasPromos() bool`

HasPromos returns a boolean if a field has been set.

### GetPrice

`func (o *SubmitOrderRequest) GetPrice() OrderPrice`

GetPrice returns the Price field if non-nil, zero value otherwise.

### GetPriceOk

`func (o *SubmitOrderRequest) GetPriceOk() (*OrderPrice, bool)`

GetPriceOk returns a tuple with the Price field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrice

`func (o *SubmitOrderRequest) SetPrice(v OrderPrice)`

SetPrice sets Price field to given value.


### GetDineIn

`func (o *SubmitOrderRequest) GetDineIn() DineIn`

GetDineIn returns the DineIn field if non-nil, zero value otherwise.

### GetDineInOk

`func (o *SubmitOrderRequest) GetDineInOk() (*DineIn, bool)`

GetDineInOk returns a tuple with the DineIn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDineIn

`func (o *SubmitOrderRequest) SetDineIn(v DineIn)`

SetDineIn sets DineIn field to given value.

### HasDineIn

`func (o *SubmitOrderRequest) HasDineIn() bool`

HasDineIn returns a boolean if a field has been set.

### GetReceiver

`func (o *SubmitOrderRequest) GetReceiver() Receiver`

GetReceiver returns the Receiver field if non-nil, zero value otherwise.

### GetReceiverOk

`func (o *SubmitOrderRequest) GetReceiverOk() (*Receiver, bool)`

GetReceiverOk returns a tuple with the Receiver field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReceiver

`func (o *SubmitOrderRequest) SetReceiver(v Receiver)`

SetReceiver sets Receiver field to given value.

### HasReceiver

`func (o *SubmitOrderRequest) HasReceiver() bool`

HasReceiver returns a boolean if a field has been set.

### GetOrderReadyEstimation

`func (o *SubmitOrderRequest) GetOrderReadyEstimation() OrderReadyEstimation`

GetOrderReadyEstimation returns the OrderReadyEstimation field if non-nil, zero value otherwise.

### GetOrderReadyEstimationOk

`func (o *SubmitOrderRequest) GetOrderReadyEstimationOk() (*OrderReadyEstimation, bool)`

GetOrderReadyEstimationOk returns a tuple with the OrderReadyEstimation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrderReadyEstimation

`func (o *SubmitOrderRequest) SetOrderReadyEstimation(v OrderReadyEstimation)`

SetOrderReadyEstimation sets OrderReadyEstimation field to given value.

### HasOrderReadyEstimation

`func (o *SubmitOrderRequest) HasOrderReadyEstimation() bool`

HasOrderReadyEstimation returns a boolean if a field has been set.

### GetMembershipID

`func (o *SubmitOrderRequest) GetMembershipID() string`

GetMembershipID returns the MembershipID field if non-nil, zero value otherwise.

### GetMembershipIDOk

`func (o *SubmitOrderRequest) GetMembershipIDOk() (*string, bool)`

GetMembershipIDOk returns a tuple with the MembershipID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMembershipID

`func (o *SubmitOrderRequest) SetMembershipID(v string)`

SetMembershipID sets MembershipID field to given value.

### HasMembershipID

`func (o *SubmitOrderRequest) HasMembershipID() bool`

HasMembershipID returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


