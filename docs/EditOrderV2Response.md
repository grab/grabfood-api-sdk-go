# EditOrderV2Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**OrderID** | Pointer to **string** | The order&#39;s ID that is returned from GrabFood. Refer to FAQs for more details about [orderID and shortOrderNumber](#section/Order/What&#39;s-the-difference-between-orderID-and-shortOrderNumber). | [optional] 
**ShortOrderNumber** | Pointer to **string** | The GrabFood short order number. This is unique for each merchant per day. Refer to FAQs for more details about [orderID and shortOrderNumber](#section/Order/What&#39;s-the-difference-between-orderID-and-shortOrderNumber). | [optional] 
**MerchantID** | Pointer to **string** | The merchant&#39;s ID that is in GrabFood&#39;s database. | [optional] 
**PartnerMerchantID** | Pointer to **string** | The merchant&#39;s ID that is on the partner&#39;s database. | [optional] 
**PaymentType** | Pointer to **string** | The payment method used. Refer to FAQs for more details about [paymentType](#section/Order/Does-the-paymentType-affect-partners). | [optional] 
**Cutlery** | Pointer to **bool** | The boolean value to indicate whether cutlery are needed or not. Refer to FAQs for more details about [cutlery](#section/Order/What-do-the-true-or-false-values-mean-for-cutlery). | [optional] 
**OrderTime** | Pointer to **string** | The UTC time that a consumer places the order, based on ISO_8601/RFC3339. | [optional] 
**SubmitTime** | Pointer to **time.Time** | The order submit time, based on ISO_8601/RFC3339. &#x60;null&#x60; in EditOrder V2 response. Only present in the [List Orders](#tag/list-order) response. | [optional] 
**CompleteTime** | Pointer to **time.Time** | The order complete time, based on ISO_8601/RFC3339. &#x60;null&#x60; in EditOrder V2 response. Only present in the [List Orders](#tag/list-order) response. | [optional] 
**ScheduledTime** | Pointer to **string** | The order scheduled time, based on ISO_8601/RFC3339. Empty for non-scheduled orders. | [optional] 
**OrderState** | Pointer to **string** | The state of the order. Empty in EditOrder V2 response. Only present in the [List Orders](#tag/list-order) response. Refer to [Order States](#section/Order-states). | [optional] 
**Currency** | Pointer to [**Currency**](Currency.md) |  | [optional] 
**FeatureFlags** | Pointer to [**OrderFeatureFlags**](OrderFeatureFlags.md) |  | [optional] 
**Items** | Pointer to [**[]OrderItem**](OrderItem.md) | The ordered items in an array of JSON Object.  | [optional] 
**Campaigns** | Pointer to [**[]OrderCampaign**](OrderCampaign.md) | The campaigns that are applicable for the order. &#x60;null&#x60; when there is no campaign applied. Only campaigns that are funded by merchants will be sent.  | [optional] 
**Promos** | Pointer to [**[]OrderPromo**](OrderPromo.md) | An array of promotion objects. &#x60;null&#x60; when there is no promo code applied. Only promotions that are funded by merchants will be sent. | [optional] 
**Price** | Pointer to [**OrderPrice**](OrderPrice.md) |  | [optional] 
**DineIn** | Pointer to [**NullableDineIn**](DineIn.md) |  | [optional] 
**Receiver** | Pointer to **map[string]interface{}** | This field is set to &#x60;null&#x60; in the EditOrder V2 response. Refer to the [Submit Order](#tag/submit-order-webhook/operation/submit-order-webhook) payload for the complete receiver information. | [optional] 
**OrderReadyEstimation** | Pointer to [**NullableOrderReadyEstimation**](OrderReadyEstimation.md) |  | [optional] 
**MembershipID** | Pointer to **string** | Membership ID for loyalty project. Only present for loyalty program partners. Empty if not applicable. | [optional] 
**Discounts** | Pointer to [**[]GrabDiscount1**](GrabDiscount1.md) | The discounts that are applicable for the paybill order in dineout STO case. &#x60;null&#x60; when there is no discount applied. This is only applicable for STO order  | [optional] 
**Payments** | Pointer to **[]string** | This field is set to &#x60;null&#x60; in the EditOrder V2 response. Refer to the [Submit Order](#tag/submit-order-webhook/operation/submit-order-webhook) payload for the complete payment details. | [optional] 

## Methods

### NewEditOrderV2Response

`func NewEditOrderV2Response() *EditOrderV2Response`

NewEditOrderV2Response instantiates a new EditOrderV2Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEditOrderV2ResponseWithDefaults

`func NewEditOrderV2ResponseWithDefaults() *EditOrderV2Response`

NewEditOrderV2ResponseWithDefaults instantiates a new EditOrderV2Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetOrderID

`func (o *EditOrderV2Response) GetOrderID() string`

GetOrderID returns the OrderID field if non-nil, zero value otherwise.

### GetOrderIDOk

`func (o *EditOrderV2Response) GetOrderIDOk() (*string, bool)`

GetOrderIDOk returns a tuple with the OrderID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrderID

`func (o *EditOrderV2Response) SetOrderID(v string)`

SetOrderID sets OrderID field to given value.

### HasOrderID

`func (o *EditOrderV2Response) HasOrderID() bool`

HasOrderID returns a boolean if a field has been set.

### GetShortOrderNumber

`func (o *EditOrderV2Response) GetShortOrderNumber() string`

GetShortOrderNumber returns the ShortOrderNumber field if non-nil, zero value otherwise.

### GetShortOrderNumberOk

`func (o *EditOrderV2Response) GetShortOrderNumberOk() (*string, bool)`

GetShortOrderNumberOk returns a tuple with the ShortOrderNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShortOrderNumber

`func (o *EditOrderV2Response) SetShortOrderNumber(v string)`

SetShortOrderNumber sets ShortOrderNumber field to given value.

### HasShortOrderNumber

`func (o *EditOrderV2Response) HasShortOrderNumber() bool`

HasShortOrderNumber returns a boolean if a field has been set.

### GetMerchantID

`func (o *EditOrderV2Response) GetMerchantID() string`

GetMerchantID returns the MerchantID field if non-nil, zero value otherwise.

### GetMerchantIDOk

`func (o *EditOrderV2Response) GetMerchantIDOk() (*string, bool)`

GetMerchantIDOk returns a tuple with the MerchantID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMerchantID

`func (o *EditOrderV2Response) SetMerchantID(v string)`

SetMerchantID sets MerchantID field to given value.

### HasMerchantID

`func (o *EditOrderV2Response) HasMerchantID() bool`

HasMerchantID returns a boolean if a field has been set.

### GetPartnerMerchantID

`func (o *EditOrderV2Response) GetPartnerMerchantID() string`

GetPartnerMerchantID returns the PartnerMerchantID field if non-nil, zero value otherwise.

### GetPartnerMerchantIDOk

`func (o *EditOrderV2Response) GetPartnerMerchantIDOk() (*string, bool)`

GetPartnerMerchantIDOk returns a tuple with the PartnerMerchantID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPartnerMerchantID

`func (o *EditOrderV2Response) SetPartnerMerchantID(v string)`

SetPartnerMerchantID sets PartnerMerchantID field to given value.

### HasPartnerMerchantID

`func (o *EditOrderV2Response) HasPartnerMerchantID() bool`

HasPartnerMerchantID returns a boolean if a field has been set.

### GetPaymentType

`func (o *EditOrderV2Response) GetPaymentType() string`

GetPaymentType returns the PaymentType field if non-nil, zero value otherwise.

### GetPaymentTypeOk

`func (o *EditOrderV2Response) GetPaymentTypeOk() (*string, bool)`

GetPaymentTypeOk returns a tuple with the PaymentType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaymentType

`func (o *EditOrderV2Response) SetPaymentType(v string)`

SetPaymentType sets PaymentType field to given value.

### HasPaymentType

`func (o *EditOrderV2Response) HasPaymentType() bool`

HasPaymentType returns a boolean if a field has been set.

### GetCutlery

`func (o *EditOrderV2Response) GetCutlery() bool`

GetCutlery returns the Cutlery field if non-nil, zero value otherwise.

### GetCutleryOk

`func (o *EditOrderV2Response) GetCutleryOk() (*bool, bool)`

GetCutleryOk returns a tuple with the Cutlery field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCutlery

`func (o *EditOrderV2Response) SetCutlery(v bool)`

SetCutlery sets Cutlery field to given value.

### HasCutlery

`func (o *EditOrderV2Response) HasCutlery() bool`

HasCutlery returns a boolean if a field has been set.

### GetOrderTime

`func (o *EditOrderV2Response) GetOrderTime() string`

GetOrderTime returns the OrderTime field if non-nil, zero value otherwise.

### GetOrderTimeOk

`func (o *EditOrderV2Response) GetOrderTimeOk() (*string, bool)`

GetOrderTimeOk returns a tuple with the OrderTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrderTime

`func (o *EditOrderV2Response) SetOrderTime(v string)`

SetOrderTime sets OrderTime field to given value.

### HasOrderTime

`func (o *EditOrderV2Response) HasOrderTime() bool`

HasOrderTime returns a boolean if a field has been set.

### GetSubmitTime

`func (o *EditOrderV2Response) GetSubmitTime() time.Time`

GetSubmitTime returns the SubmitTime field if non-nil, zero value otherwise.

### GetSubmitTimeOk

`func (o *EditOrderV2Response) GetSubmitTimeOk() (*time.Time, bool)`

GetSubmitTimeOk returns a tuple with the SubmitTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubmitTime

`func (o *EditOrderV2Response) SetSubmitTime(v time.Time)`

SetSubmitTime sets SubmitTime field to given value.

### HasSubmitTime

`func (o *EditOrderV2Response) HasSubmitTime() bool`

HasSubmitTime returns a boolean if a field has been set.

### GetCompleteTime

`func (o *EditOrderV2Response) GetCompleteTime() time.Time`

GetCompleteTime returns the CompleteTime field if non-nil, zero value otherwise.

### GetCompleteTimeOk

`func (o *EditOrderV2Response) GetCompleteTimeOk() (*time.Time, bool)`

GetCompleteTimeOk returns a tuple with the CompleteTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompleteTime

`func (o *EditOrderV2Response) SetCompleteTime(v time.Time)`

SetCompleteTime sets CompleteTime field to given value.

### HasCompleteTime

`func (o *EditOrderV2Response) HasCompleteTime() bool`

HasCompleteTime returns a boolean if a field has been set.

### GetScheduledTime

`func (o *EditOrderV2Response) GetScheduledTime() string`

GetScheduledTime returns the ScheduledTime field if non-nil, zero value otherwise.

### GetScheduledTimeOk

`func (o *EditOrderV2Response) GetScheduledTimeOk() (*string, bool)`

GetScheduledTimeOk returns a tuple with the ScheduledTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScheduledTime

`func (o *EditOrderV2Response) SetScheduledTime(v string)`

SetScheduledTime sets ScheduledTime field to given value.

### HasScheduledTime

`func (o *EditOrderV2Response) HasScheduledTime() bool`

HasScheduledTime returns a boolean if a field has been set.

### GetOrderState

`func (o *EditOrderV2Response) GetOrderState() string`

GetOrderState returns the OrderState field if non-nil, zero value otherwise.

### GetOrderStateOk

`func (o *EditOrderV2Response) GetOrderStateOk() (*string, bool)`

GetOrderStateOk returns a tuple with the OrderState field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrderState

`func (o *EditOrderV2Response) SetOrderState(v string)`

SetOrderState sets OrderState field to given value.

### HasOrderState

`func (o *EditOrderV2Response) HasOrderState() bool`

HasOrderState returns a boolean if a field has been set.

### GetCurrency

`func (o *EditOrderV2Response) GetCurrency() Currency`

GetCurrency returns the Currency field if non-nil, zero value otherwise.

### GetCurrencyOk

`func (o *EditOrderV2Response) GetCurrencyOk() (*Currency, bool)`

GetCurrencyOk returns a tuple with the Currency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrency

`func (o *EditOrderV2Response) SetCurrency(v Currency)`

SetCurrency sets Currency field to given value.

### HasCurrency

`func (o *EditOrderV2Response) HasCurrency() bool`

HasCurrency returns a boolean if a field has been set.

### GetFeatureFlags

`func (o *EditOrderV2Response) GetFeatureFlags() OrderFeatureFlags`

GetFeatureFlags returns the FeatureFlags field if non-nil, zero value otherwise.

### GetFeatureFlagsOk

`func (o *EditOrderV2Response) GetFeatureFlagsOk() (*OrderFeatureFlags, bool)`

GetFeatureFlagsOk returns a tuple with the FeatureFlags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFeatureFlags

`func (o *EditOrderV2Response) SetFeatureFlags(v OrderFeatureFlags)`

SetFeatureFlags sets FeatureFlags field to given value.

### HasFeatureFlags

`func (o *EditOrderV2Response) HasFeatureFlags() bool`

HasFeatureFlags returns a boolean if a field has been set.

### GetItems

`func (o *EditOrderV2Response) GetItems() []OrderItem`

GetItems returns the Items field if non-nil, zero value otherwise.

### GetItemsOk

`func (o *EditOrderV2Response) GetItemsOk() (*[]OrderItem, bool)`

GetItemsOk returns a tuple with the Items field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItems

`func (o *EditOrderV2Response) SetItems(v []OrderItem)`

SetItems sets Items field to given value.

### HasItems

`func (o *EditOrderV2Response) HasItems() bool`

HasItems returns a boolean if a field has been set.

### GetCampaigns

`func (o *EditOrderV2Response) GetCampaigns() []OrderCampaign`

GetCampaigns returns the Campaigns field if non-nil, zero value otherwise.

### GetCampaignsOk

`func (o *EditOrderV2Response) GetCampaignsOk() (*[]OrderCampaign, bool)`

GetCampaignsOk returns a tuple with the Campaigns field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCampaigns

`func (o *EditOrderV2Response) SetCampaigns(v []OrderCampaign)`

SetCampaigns sets Campaigns field to given value.

### HasCampaigns

`func (o *EditOrderV2Response) HasCampaigns() bool`

HasCampaigns returns a boolean if a field has been set.

### SetCampaignsNil

`func (o *EditOrderV2Response) SetCampaignsNil(b bool)`

 SetCampaignsNil sets the value for Campaigns to be an explicit nil

### UnsetCampaigns
`func (o *EditOrderV2Response) UnsetCampaigns()`

UnsetCampaigns ensures that no value is present for Campaigns, not even an explicit nil
### GetPromos

`func (o *EditOrderV2Response) GetPromos() []OrderPromo`

GetPromos returns the Promos field if non-nil, zero value otherwise.

### GetPromosOk

`func (o *EditOrderV2Response) GetPromosOk() (*[]OrderPromo, bool)`

GetPromosOk returns a tuple with the Promos field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPromos

`func (o *EditOrderV2Response) SetPromos(v []OrderPromo)`

SetPromos sets Promos field to given value.

### HasPromos

`func (o *EditOrderV2Response) HasPromos() bool`

HasPromos returns a boolean if a field has been set.

### SetPromosNil

`func (o *EditOrderV2Response) SetPromosNil(b bool)`

 SetPromosNil sets the value for Promos to be an explicit nil

### UnsetPromos
`func (o *EditOrderV2Response) UnsetPromos()`

UnsetPromos ensures that no value is present for Promos, not even an explicit nil
### GetPrice

`func (o *EditOrderV2Response) GetPrice() OrderPrice`

GetPrice returns the Price field if non-nil, zero value otherwise.

### GetPriceOk

`func (o *EditOrderV2Response) GetPriceOk() (*OrderPrice, bool)`

GetPriceOk returns a tuple with the Price field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrice

`func (o *EditOrderV2Response) SetPrice(v OrderPrice)`

SetPrice sets Price field to given value.

### HasPrice

`func (o *EditOrderV2Response) HasPrice() bool`

HasPrice returns a boolean if a field has been set.

### GetDineIn

`func (o *EditOrderV2Response) GetDineIn() DineIn`

GetDineIn returns the DineIn field if non-nil, zero value otherwise.

### GetDineInOk

`func (o *EditOrderV2Response) GetDineInOk() (*DineIn, bool)`

GetDineInOk returns a tuple with the DineIn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDineIn

`func (o *EditOrderV2Response) SetDineIn(v DineIn)`

SetDineIn sets DineIn field to given value.

### HasDineIn

`func (o *EditOrderV2Response) HasDineIn() bool`

HasDineIn returns a boolean if a field has been set.

### SetDineInNil

`func (o *EditOrderV2Response) SetDineInNil(b bool)`

 SetDineInNil sets the value for DineIn to be an explicit nil

### UnsetDineIn
`func (o *EditOrderV2Response) UnsetDineIn()`

UnsetDineIn ensures that no value is present for DineIn, not even an explicit nil
### GetReceiver

`func (o *EditOrderV2Response) GetReceiver() map[string]interface{}`

GetReceiver returns the Receiver field if non-nil, zero value otherwise.

### GetReceiverOk

`func (o *EditOrderV2Response) GetReceiverOk() (*map[string]interface{}, bool)`

GetReceiverOk returns a tuple with the Receiver field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReceiver

`func (o *EditOrderV2Response) SetReceiver(v map[string]interface{})`

SetReceiver sets Receiver field to given value.

### HasReceiver

`func (o *EditOrderV2Response) HasReceiver() bool`

HasReceiver returns a boolean if a field has been set.

### GetOrderReadyEstimation

`func (o *EditOrderV2Response) GetOrderReadyEstimation() OrderReadyEstimation`

GetOrderReadyEstimation returns the OrderReadyEstimation field if non-nil, zero value otherwise.

### GetOrderReadyEstimationOk

`func (o *EditOrderV2Response) GetOrderReadyEstimationOk() (*OrderReadyEstimation, bool)`

GetOrderReadyEstimationOk returns a tuple with the OrderReadyEstimation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrderReadyEstimation

`func (o *EditOrderV2Response) SetOrderReadyEstimation(v OrderReadyEstimation)`

SetOrderReadyEstimation sets OrderReadyEstimation field to given value.

### HasOrderReadyEstimation

`func (o *EditOrderV2Response) HasOrderReadyEstimation() bool`

HasOrderReadyEstimation returns a boolean if a field has been set.

### SetOrderReadyEstimationNil

`func (o *EditOrderV2Response) SetOrderReadyEstimationNil(b bool)`

 SetOrderReadyEstimationNil sets the value for OrderReadyEstimation to be an explicit nil

### UnsetOrderReadyEstimation
`func (o *EditOrderV2Response) UnsetOrderReadyEstimation()`

UnsetOrderReadyEstimation ensures that no value is present for OrderReadyEstimation, not even an explicit nil
### GetMembershipID

`func (o *EditOrderV2Response) GetMembershipID() string`

GetMembershipID returns the MembershipID field if non-nil, zero value otherwise.

### GetMembershipIDOk

`func (o *EditOrderV2Response) GetMembershipIDOk() (*string, bool)`

GetMembershipIDOk returns a tuple with the MembershipID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMembershipID

`func (o *EditOrderV2Response) SetMembershipID(v string)`

SetMembershipID sets MembershipID field to given value.

### HasMembershipID

`func (o *EditOrderV2Response) HasMembershipID() bool`

HasMembershipID returns a boolean if a field has been set.

### GetDiscounts

`func (o *EditOrderV2Response) GetDiscounts() []GrabDiscount1`

GetDiscounts returns the Discounts field if non-nil, zero value otherwise.

### GetDiscountsOk

`func (o *EditOrderV2Response) GetDiscountsOk() (*[]GrabDiscount1, bool)`

GetDiscountsOk returns a tuple with the Discounts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDiscounts

`func (o *EditOrderV2Response) SetDiscounts(v []GrabDiscount1)`

SetDiscounts sets Discounts field to given value.

### HasDiscounts

`func (o *EditOrderV2Response) HasDiscounts() bool`

HasDiscounts returns a boolean if a field has been set.

### SetDiscountsNil

`func (o *EditOrderV2Response) SetDiscountsNil(b bool)`

 SetDiscountsNil sets the value for Discounts to be an explicit nil

### UnsetDiscounts
`func (o *EditOrderV2Response) UnsetDiscounts()`

UnsetDiscounts ensures that no value is present for Discounts, not even an explicit nil
### GetPayments

`func (o *EditOrderV2Response) GetPayments() []string`

GetPayments returns the Payments field if non-nil, zero value otherwise.

### GetPaymentsOk

`func (o *EditOrderV2Response) GetPaymentsOk() (*[]string, bool)`

GetPaymentsOk returns a tuple with the Payments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPayments

`func (o *EditOrderV2Response) SetPayments(v []string)`

SetPayments sets Payments field to given value.

### HasPayments

`func (o *EditOrderV2Response) HasPayments() bool`

HasPayments returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


