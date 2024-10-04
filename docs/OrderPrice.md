# OrderPrice

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Subtotal** | **int64** | Total item and modifier price (tax-inclusive) in the minor unit. &#x60;Sum of all (Item price * quantity) | 2550*1&#x3D;2550&#x60;. | 
**Tax** | Pointer to **int64** | GrabFood&#39;s tax in the minor unit. &#x60;(subtotal-merchantFundPromo)* Tax /(1+Tax) | (2550-475)*0.06/1.06&#x3D;117&#x60;. Refer to FAQs for more details about [tax](#section/Order/How-is-tax-calculated). | [optional] 
**MerchantChargeFee** | Pointer to **int64** | Any additional fee charged by merchant, which is 100% paid out to the merchant. Eg. Takeaway, packaging costs, dine-in charge. | [optional] 
**GrabFundPromo** | Pointer to **int64** | GrabFood&#39;s promo fund in the minor unit. Calculated based on funded ratio. | [optional] 
**MerchantFundPromo** | Pointer to **int64** | The merchant&#39;s promo fund in the minor unit. Calculated based on funded ratio. | [optional] 
**BasketPromo** | Pointer to **int64** | The total amount promo applied to the basket items only (item level/order level) in the minor unit. Delivery fee is excluded. &#x60;(grabFundPromo + merchantFundPromo) | 300 + 475 &#x3D; 775&#x60;  | [optional] 
**DeliveryFee** | Pointer to **int64** | The delivery fee in the minor unit. | [optional] 
**EaterPayment** | **int64** | The total amount consumer paid in the minor unit. &#x60;(subtotal + deliveryFee) - (sum of all promo) | (2550+400)-775&#x3D;2175&#x60; | 

## Methods

### NewOrderPrice

`func NewOrderPrice(subtotal int64, eaterPayment int64, ) *OrderPrice`

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



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


