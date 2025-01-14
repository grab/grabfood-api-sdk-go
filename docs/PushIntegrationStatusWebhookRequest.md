# PushIntegrationStatusWebhookRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PartnerMerchantID** | **string** | The merchant&#39;s ID that is on the partner&#39;s database. | 
**GrabMerchantID** | **string** | The merchant&#39;s ID that is in GrabFood&#39;s database. | 
**IntegrationStatus** | **string** | The store integration status. - &#x60;INACTIVE&#x60;: Merchant integration deactivated - &#x60;ACTIVE&#x60;: Merchant integration activated - &#x60;SYNCING&#x60;: Merchant integration is syncing - &#x60;FAILED&#x60;: Merchant integration has failed  | 

## Methods

### NewPushIntegrationStatusWebhookRequest

`func NewPushIntegrationStatusWebhookRequest(partnerMerchantID string, grabMerchantID string, integrationStatus string, ) *PushIntegrationStatusWebhookRequest`

NewPushIntegrationStatusWebhookRequest instantiates a new PushIntegrationStatusWebhookRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPushIntegrationStatusWebhookRequestWithDefaults

`func NewPushIntegrationStatusWebhookRequestWithDefaults() *PushIntegrationStatusWebhookRequest`

NewPushIntegrationStatusWebhookRequestWithDefaults instantiates a new PushIntegrationStatusWebhookRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPartnerMerchantID

`func (o *PushIntegrationStatusWebhookRequest) GetPartnerMerchantID() string`

GetPartnerMerchantID returns the PartnerMerchantID field if non-nil, zero value otherwise.

### GetPartnerMerchantIDOk

`func (o *PushIntegrationStatusWebhookRequest) GetPartnerMerchantIDOk() (*string, bool)`

GetPartnerMerchantIDOk returns a tuple with the PartnerMerchantID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPartnerMerchantID

`func (o *PushIntegrationStatusWebhookRequest) SetPartnerMerchantID(v string)`

SetPartnerMerchantID sets PartnerMerchantID field to given value.


### GetGrabMerchantID

`func (o *PushIntegrationStatusWebhookRequest) GetGrabMerchantID() string`

GetGrabMerchantID returns the GrabMerchantID field if non-nil, zero value otherwise.

### GetGrabMerchantIDOk

`func (o *PushIntegrationStatusWebhookRequest) GetGrabMerchantIDOk() (*string, bool)`

GetGrabMerchantIDOk returns a tuple with the GrabMerchantID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGrabMerchantID

`func (o *PushIntegrationStatusWebhookRequest) SetGrabMerchantID(v string)`

SetGrabMerchantID sets GrabMerchantID field to given value.


### GetIntegrationStatus

`func (o *PushIntegrationStatusWebhookRequest) GetIntegrationStatus() string`

GetIntegrationStatus returns the IntegrationStatus field if non-nil, zero value otherwise.

### GetIntegrationStatusOk

`func (o *PushIntegrationStatusWebhookRequest) GetIntegrationStatusOk() (*string, bool)`

GetIntegrationStatusOk returns a tuple with the IntegrationStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIntegrationStatus

`func (o *PushIntegrationStatusWebhookRequest) SetIntegrationStatus(v string)`

SetIntegrationStatus sets IntegrationStatus field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


