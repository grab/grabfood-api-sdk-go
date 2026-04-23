# OrderPrice

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Subtotal** | **int64** | Total item and modifier price (tax-inclusive) in the minor unit. &#x60;&#x60;&#x60; subtotal &#x3D; Sum of all (item price * quantity) | 2550*1&#x3D;2550  | 
**Tax** | Pointer to **int64** | GrabFood&#39;s tax in the minor unit. Refer to FAQs for more details about [tax](#section/Order/How-is-tax-calculated). &#x60;&#x60;&#x60; tax &#x3D; (subtotal + merchantChargeFee - merchantFundPromo) * Tax / (1+Tax) | (2550-475)*0.06/1.06&#x3D;117  | [optional] 
**MerchantChargeFee** | Pointer to **int64** | Any additional fee charged by merchant (tax-inclusive), which is 100% paid out to the merchant. Reach out to your integration support team for the configuration. Eg. Takeaway, packaging costs, dine-in charge.  | [optional] 
**ServiceChargeFee** | Pointer to **int64** | Additional service charge fee charged by merchant (tax-inclusive), which is 100% paid out to the merchant. Reach out to your integration support team for the configuration.  | [optional] 
**GrabFundPromo** | Pointer to **int64** | GrabFood&#39;s promo fund in the minor unit. Calculated based on funded ratio. Only present when &#x60;paymentType:CASH&#x60; or &#x60;orderType:DeliveredByRestaurant&#x60;. Otherwise, it will be set to &#x60;0&#x60;. | [optional] 
**MerchantFundPromo** | Pointer to **int64** | The merchant&#39;s promo fund in the minor unit. Calculated based on funded ratio. | [optional] 
**BasketPromo** | Pointer to **int64** | The total amount promo applied to the basket items only (item level/order level) in the minor unit, excluding delivery fee. Only present when &#x60;paymentType: CASH&#x60; or &#x60;orderType: DeliveredByRestaurant&#x60;. Otherwise, it will be set to &#x60;0&#x60;.  &#x60;&#x60;&#x60; basketPromo &#x3D; (grabFundPromo + merchantFundPromo) | 300 + 475 &#x3D; 775  | [optional] 
**DeliveryFee** | Pointer to **int64** | The delivery fee in the minor unit. Only present when &#x60;paymentType:CASH&#x60; or &#x60;orderType:DeliveredByRestaurant&#x60;. Otherwise, it will be set to &#x60;0&#x60;. | [optional] 
**SmallOrderFee** | Pointer to **int64** | The fee charged by GrabFood for order that does not meet a certain minimum order value. Only present when &#x60;paymentType:CASH&#x60; and &#x60;orderType:DeliveredByRestaurant&#x60;. | [optional] 
**EaterPayment** | Pointer to **int64** | The total amount paid by the consumer in the minor unit, excluding some additional fees charged by GrabFood. Only present when &#x60;paymentType:CASH&#x60; or &#x60;orderType:DeliveredByRestaurant&#x60;. Otherwise, it will be set to &#x60;0&#x60;.  &#x60;&#x60;&#x60; eaterPayment &#x3D; (subtotal + merchantChargeFee + deliveryFee) - (sum of all promo) | (2550+400)-775&#x3D;2175  | [optional] 
**Total** | Pointer to **int64** | The total merchant-related amount calculated exclusive of commission charges. Formulae is the same for all delivery method.  &#x60;&#x60;&#x60; total &#x3D; subtotal + merchantChargeFee - merchantFundPromo | 2550+0-475&#x3D;2075  | [optional] 
**MerchantEarning** | Pointer to [**NullableMerchantEarning**](MerchantEarning.md) |  | [optional] 

## Methods

### NewOrderPrice

`func NewOrderPrice(subtotal int64, ) *OrderPrice`

NewOrderPrice instantiates a new OrderPrice object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOrderPriceWithDefaults

`func NewOrderPriceWithDefaults() *OrderPrice`

NewOrderPriceWithDefaults instantiates a new OrderPrice object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSubtotal

`func (o *OrderPrice) GetSubtotal() int64`

GetSubtotal returns the Subtotal field if non-nil, zero value otherwise.

### GetSubtotalOk

`func (o *OrderPrice) GetSubtotalOk() (*int64, bool)`

GetSubtotalOk returns a tuple with the Subtotal field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubtotal

`func (o *OrderPrice) SetSubtotal(v int64)`

SetSubtotal sets Subtotal field to given value.


### GetTax

`func (o *OrderPrice) GetTax() int64`

GetTax returns the Tax field if non-nil, zero value otherwise.

### GetTaxOk

`func (o *OrderPrice) GetTaxOk() (*int64, bool)`

GetTaxOk returns a tuple with the Tax field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTax

`func (o *OrderPrice) SetTax(v int64)`

SetTax sets Tax field to given value.

### HasTax

`func (o *OrderPrice) HasTax() bool`

HasTax returns a boolean if a field has been set.

### GetMerchantChargeFee

`func (o *OrderPrice) GetMerchantChargeFee() int64`

GetMerchantChargeFee returns the MerchantChargeFee field if non-nil, zero value otherwise.

### GetMerchantChargeFeeOk

`func (o *OrderPrice) GetMerchantChargeFeeOk() (*int64, bool)`

GetMerchantChargeFeeOk returns a tuple with the MerchantChargeFee field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMerchantChargeFee

`func (o *OrderPrice) SetMerchantChargeFee(v int64)`

SetMerchantChargeFee sets MerchantChargeFee field to given value.

### HasMerchantChargeFee

`func (o *OrderPrice) HasMerchantChargeFee() bool`

HasMerchantChargeFee returns a boolean if a field has been set.

### GetServiceChargeFee

`func (o *OrderPrice) GetServiceChargeFee() int64`

GetServiceChargeFee returns the ServiceChargeFee field if non-nil, zero value otherwise.

### GetServiceChargeFeeOk

`func (o *OrderPrice) GetServiceChargeFeeOk() (*int64, bool)`

GetServiceChargeFeeOk returns a tuple with the ServiceChargeFee field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceChargeFee

`func (o *OrderPrice) SetServiceChargeFee(v int64)`

SetServiceChargeFee sets ServiceChargeFee field to given value.

### HasServiceChargeFee

`func (o *OrderPrice) HasServiceChargeFee() bool`

HasServiceChargeFee returns a boolean if a field has been set.

### GetGrabFundPromo

`func (o *OrderPrice) GetGrabFundPromo() int64`

GetGrabFundPromo returns the GrabFundPromo field if non-nil, zero value otherwise.

### GetGrabFundPromoOk

`func (o *OrderPrice) GetGrabFundPromoOk() (*int64, bool)`

GetGrabFundPromoOk returns a tuple with the GrabFundPromo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGrabFundPromo

`func (o *OrderPrice) SetGrabFundPromo(v int64)`

SetGrabFundPromo sets GrabFundPromo field to given value.

### HasGrabFundPromo

`func (o *OrderPrice) HasGrabFundPromo() bool`

HasGrabFundPromo returns a boolean if a field has been set.

### GetMerchantFundPromo

`func (o *OrderPrice) GetMerchantFundPromo() int64`

GetMerchantFundPromo returns the MerchantFundPromo field if non-nil, zero value otherwise.

### GetMerchantFundPromoOk

`func (o *OrderPrice) GetMerchantFundPromoOk() (*int64, bool)`

GetMerchantFundPromoOk returns a tuple with the MerchantFundPromo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMerchantFundPromo

`func (o *OrderPrice) SetMerchantFundPromo(v int64)`

SetMerchantFundPromo sets MerchantFundPromo field to given value.

### HasMerchantFundPromo

`func (o *OrderPrice) HasMerchantFundPromo() bool`

HasMerchantFundPromo returns a boolean if a field has been set.

### GetBasketPromo

`func (o *OrderPrice) GetBasketPromo() int64`

GetBasketPromo returns the BasketPromo field if non-nil, zero value otherwise.

### GetBasketPromoOk

`func (o *OrderPrice) GetBasketPromoOk() (*int64, bool)`

GetBasketPromoOk returns a tuple with the BasketPromo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBasketPromo

`func (o *OrderPrice) SetBasketPromo(v int64)`

SetBasketPromo sets BasketPromo field to given value.

### HasBasketPromo

`func (o *OrderPrice) HasBasketPromo() bool`

HasBasketPromo returns a boolean if a field has been set.

### GetDeliveryFee

`func (o *OrderPrice) GetDeliveryFee() int64`

GetDeliveryFee returns the DeliveryFee field if non-nil, zero value otherwise.

### GetDeliveryFeeOk

`func (o *OrderPrice) GetDeliveryFeeOk() (*int64, bool)`

GetDeliveryFeeOk returns a tuple with the DeliveryFee field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeliveryFee

`func (o *OrderPrice) SetDeliveryFee(v int64)`

SetDeliveryFee sets DeliveryFee field to given value.

### HasDeliveryFee

`func (o *OrderPrice) HasDeliveryFee() bool`

HasDeliveryFee returns a boolean if a field has been set.

### GetSmallOrderFee

`func (o *OrderPrice) GetSmallOrderFee() int64`

GetSmallOrderFee returns the SmallOrderFee field if non-nil, zero value otherwise.

### GetSmallOrderFeeOk

`func (o *OrderPrice) GetSmallOrderFeeOk() (*int64, bool)`

GetSmallOrderFeeOk returns a tuple with the SmallOrderFee field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSmallOrderFee

`func (o *OrderPrice) SetSmallOrderFee(v int64)`

SetSmallOrderFee sets SmallOrderFee field to given value.

### HasSmallOrderFee

`func (o *OrderPrice) HasSmallOrderFee() bool`

HasSmallOrderFee returns a boolean if a field has been set.

### GetEaterPayment

`func (o *OrderPrice) GetEaterPayment() int64`

GetEaterPayment returns the EaterPayment field if non-nil, zero value otherwise.

### GetEaterPaymentOk

`func (o *OrderPrice) GetEaterPaymentOk() (*int64, bool)`

GetEaterPaymentOk returns a tuple with the EaterPayment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEaterPayment

`func (o *OrderPrice) SetEaterPayment(v int64)`

SetEaterPayment sets EaterPayment field to given value.

### HasEaterPayment

`func (o *OrderPrice) HasEaterPayment() bool`

HasEaterPayment returns a boolean if a field has been set.

### GetTotal

`func (o *OrderPrice) GetTotal() int64`

GetTotal returns the Total field if non-nil, zero value otherwise.

### GetTotalOk

`func (o *OrderPrice) GetTotalOk() (*int64, bool)`

GetTotalOk returns a tuple with the Total field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotal

`func (o *OrderPrice) SetTotal(v int64)`

SetTotal sets Total field to given value.

### HasTotal

`func (o *OrderPrice) HasTotal() bool`

HasTotal returns a boolean if a field has been set.

### GetMerchantEarning

`func (o *OrderPrice) GetMerchantEarning() MerchantEarning`

GetMerchantEarning returns the MerchantEarning field if non-nil, zero value otherwise.

### GetMerchantEarningOk

`func (o *OrderPrice) GetMerchantEarningOk() (*MerchantEarning, bool)`

GetMerchantEarningOk returns a tuple with the MerchantEarning field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMerchantEarning

`func (o *OrderPrice) SetMerchantEarning(v MerchantEarning)`

SetMerchantEarning sets MerchantEarning field to given value.

### HasMerchantEarning

`func (o *OrderPrice) HasMerchantEarning() bool`

HasMerchantEarning returns a boolean if a field has been set.

### SetMerchantEarningNil

`func (o *OrderPrice) SetMerchantEarningNil(b bool)`

 SetMerchantEarningNil sets the value for MerchantEarning to be an explicit nil

### UnsetMerchantEarning
`func (o *OrderPrice) UnsetMerchantEarning()`

UnsetMerchantEarning ensures that no value is present for MerchantEarning, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


