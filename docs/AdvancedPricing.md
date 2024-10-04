# AdvancedPricing

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DeliveryOnDemandGrabApp** | Pointer to **int64** | Service type: &#x60;Delivery&#x60;, Order type: &#x60;Instant&#x60;, Channel: &#x60;Grab App&#x60;  | [optional] 
**DeliveryScheduledGrabApp** | Pointer to **int64** | Service type: &#x60;Delivery&#x60;, Order type: &#x60;Scheduled&#x60;, Channel: &#x60;Grab App&#x60;  | [optional] 
**SelfPickUpOnDemandGrabApp** | Pointer to **int64** | Service type: &#x60;Self Pick Up&#x60;, Order type: &#x60;Instant&#x60;, Channel: &#x60;Grab App&#x60;  | [optional] 
**DineInOnDemandGrabApp** | Pointer to **int64** | Service type: &#x60;Dine In&#x60;, Order type: &#x60;Instant&#x60;, Channel: &#x60;Grab App&#x60;  | [optional] 
**DeliveryOnDemandStoreFront** | Pointer to **int64** | Service type: &#x60;Delivery&#x60;, Order type: &#x60;Instant&#x60;, Channel: &#x60;Store Front&#x60;  | [optional] 
**DeliveryScheduledStoreFront** | Pointer to **int64** | Service type: &#x60;Delivery&#x60;, Order type: &#x60;Scheduled&#x60;, Channel: &#x60;Store Front&#x60;  | [optional] 
**SelfPickUpOnDemandStoreFront** | Pointer to **int64** | Service type: &#x60;Self Pick Up&#x60;, Order type: &#x60;Instant&#x60;, Channel: &#x60;Store Front&#x60;  | [optional] 

## Methods

### NewAdvancedPricing

`func NewAdvancedPricing() *AdvancedPricing`

NewAdvancedPricing instantiates a new AdvancedPricing object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAdvancedPricingWithDefaults

`func NewAdvancedPricingWithDefaults() *AdvancedPricing`

NewAdvancedPricingWithDefaults instantiates a new AdvancedPricing object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDeliveryOnDemandGrabApp

`func (o *AdvancedPricing) GetDeliveryOnDemandGrabApp() int64`

GetDeliveryOnDemandGrabApp returns the DeliveryOnDemandGrabApp field if non-nil, zero value otherwise.

### GetDeliveryOnDemandGrabAppOk

`func (o *AdvancedPricing) GetDeliveryOnDemandGrabAppOk() (*int64, bool)`

GetDeliveryOnDemandGrabAppOk returns a tuple with the DeliveryOnDemandGrabApp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeliveryOnDemandGrabApp

`func (o *AdvancedPricing) SetDeliveryOnDemandGrabApp(v int64)`

SetDeliveryOnDemandGrabApp sets DeliveryOnDemandGrabApp field to given value.

### HasDeliveryOnDemandGrabApp

`func (o *AdvancedPricing) HasDeliveryOnDemandGrabApp() bool`

HasDeliveryOnDemandGrabApp returns a boolean if a field has been set.

### GetDeliveryScheduledGrabApp

`func (o *AdvancedPricing) GetDeliveryScheduledGrabApp() int64`

GetDeliveryScheduledGrabApp returns the DeliveryScheduledGrabApp field if non-nil, zero value otherwise.

### GetDeliveryScheduledGrabAppOk

`func (o *AdvancedPricing) GetDeliveryScheduledGrabAppOk() (*int64, bool)`

GetDeliveryScheduledGrabAppOk returns a tuple with the DeliveryScheduledGrabApp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeliveryScheduledGrabApp

`func (o *AdvancedPricing) SetDeliveryScheduledGrabApp(v int64)`

SetDeliveryScheduledGrabApp sets DeliveryScheduledGrabApp field to given value.

### HasDeliveryScheduledGrabApp

`func (o *AdvancedPricing) HasDeliveryScheduledGrabApp() bool`

HasDeliveryScheduledGrabApp returns a boolean if a field has been set.

### GetSelfPickUpOnDemandGrabApp

`func (o *AdvancedPricing) GetSelfPickUpOnDemandGrabApp() int64`

GetSelfPickUpOnDemandGrabApp returns the SelfPickUpOnDemandGrabApp field if non-nil, zero value otherwise.

### GetSelfPickUpOnDemandGrabAppOk

`func (o *AdvancedPricing) GetSelfPickUpOnDemandGrabAppOk() (*int64, bool)`

GetSelfPickUpOnDemandGrabAppOk returns a tuple with the SelfPickUpOnDemandGrabApp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSelfPickUpOnDemandGrabApp

`func (o *AdvancedPricing) SetSelfPickUpOnDemandGrabApp(v int64)`

SetSelfPickUpOnDemandGrabApp sets SelfPickUpOnDemandGrabApp field to given value.

### HasSelfPickUpOnDemandGrabApp

`func (o *AdvancedPricing) HasSelfPickUpOnDemandGrabApp() bool`

HasSelfPickUpOnDemandGrabApp returns a boolean if a field has been set.

### GetDineInOnDemandGrabApp

`func (o *AdvancedPricing) GetDineInOnDemandGrabApp() int64`

GetDineInOnDemandGrabApp returns the DineInOnDemandGrabApp field if non-nil, zero value otherwise.

### GetDineInOnDemandGrabAppOk

`func (o *AdvancedPricing) GetDineInOnDemandGrabAppOk() (*int64, bool)`

GetDineInOnDemandGrabAppOk returns a tuple with the DineInOnDemandGrabApp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDineInOnDemandGrabApp

`func (o *AdvancedPricing) SetDineInOnDemandGrabApp(v int64)`

SetDineInOnDemandGrabApp sets DineInOnDemandGrabApp field to given value.

### HasDineInOnDemandGrabApp

`func (o *AdvancedPricing) HasDineInOnDemandGrabApp() bool`

HasDineInOnDemandGrabApp returns a boolean if a field has been set.

### GetDeliveryOnDemandStoreFront

`func (o *AdvancedPricing) GetDeliveryOnDemandStoreFront() int64`

GetDeliveryOnDemandStoreFront returns the DeliveryOnDemandStoreFront field if non-nil, zero value otherwise.

### GetDeliveryOnDemandStoreFrontOk

`func (o *AdvancedPricing) GetDeliveryOnDemandStoreFrontOk() (*int64, bool)`

GetDeliveryOnDemandStoreFrontOk returns a tuple with the DeliveryOnDemandStoreFront field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeliveryOnDemandStoreFront

`func (o *AdvancedPricing) SetDeliveryOnDemandStoreFront(v int64)`

SetDeliveryOnDemandStoreFront sets DeliveryOnDemandStoreFront field to given value.

### HasDeliveryOnDemandStoreFront

`func (o *AdvancedPricing) HasDeliveryOnDemandStoreFront() bool`

HasDeliveryOnDemandStoreFront returns a boolean if a field has been set.

### GetDeliveryScheduledStoreFront

`func (o *AdvancedPricing) GetDeliveryScheduledStoreFront() int64`

GetDeliveryScheduledStoreFront returns the DeliveryScheduledStoreFront field if non-nil, zero value otherwise.

### GetDeliveryScheduledStoreFrontOk

`func (o *AdvancedPricing) GetDeliveryScheduledStoreFrontOk() (*int64, bool)`

GetDeliveryScheduledStoreFrontOk returns a tuple with the DeliveryScheduledStoreFront field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeliveryScheduledStoreFront

`func (o *AdvancedPricing) SetDeliveryScheduledStoreFront(v int64)`

SetDeliveryScheduledStoreFront sets DeliveryScheduledStoreFront field to given value.

### HasDeliveryScheduledStoreFront

`func (o *AdvancedPricing) HasDeliveryScheduledStoreFront() bool`

HasDeliveryScheduledStoreFront returns a boolean if a field has been set.

### GetSelfPickUpOnDemandStoreFront

`func (o *AdvancedPricing) GetSelfPickUpOnDemandStoreFront() int64`

GetSelfPickUpOnDemandStoreFront returns the SelfPickUpOnDemandStoreFront field if non-nil, zero value otherwise.

### GetSelfPickUpOnDemandStoreFrontOk

`func (o *AdvancedPricing) GetSelfPickUpOnDemandStoreFrontOk() (*int64, bool)`

GetSelfPickUpOnDemandStoreFrontOk returns a tuple with the SelfPickUpOnDemandStoreFront field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSelfPickUpOnDemandStoreFront

`func (o *AdvancedPricing) SetSelfPickUpOnDemandStoreFront(v int64)`

SetSelfPickUpOnDemandStoreFront sets SelfPickUpOnDemandStoreFront field to given value.

### HasSelfPickUpOnDemandStoreFront

`func (o *AdvancedPricing) HasSelfPickUpOnDemandStoreFront() bool`

HasSelfPickUpOnDemandStoreFront returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


