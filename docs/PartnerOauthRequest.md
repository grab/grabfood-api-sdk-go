# PartnerOauthRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClientId** | **string** | The client identifier issued to obtain the OAuth access token. | 
**ClientSecret** | **string** | The client secret issued to obtain the OAuth access token. | 
**GrantType** | **string** | Refers to the way an application gets OAuth access token. | 
**Scope** | Pointer to **string** | A space delimited list of partner scopes. | [optional] 

## Methods

### NewPartnerOauthRequest

`func NewPartnerOauthRequest(clientId string, clientSecret string, grantType string, ) *PartnerOauthRequest`

NewPartnerOauthRequest instantiates a new PartnerOauthRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPartnerOauthRequestWithDefaults

`func NewPartnerOauthRequestWithDefaults() *PartnerOauthRequest`

NewPartnerOauthRequestWithDefaults instantiates a new PartnerOauthRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClientId

`func (o *PartnerOauthRequest) GetClientId() string`

GetClientId returns the ClientId field if non-nil, zero value otherwise.

### GetClientIdOk

`func (o *PartnerOauthRequest) GetClientIdOk() (*string, bool)`

GetClientIdOk returns a tuple with the ClientId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientId

`func (o *PartnerOauthRequest) SetClientId(v string)`

SetClientId sets ClientId field to given value.


### GetClientSecret

`func (o *PartnerOauthRequest) GetClientSecret() string`

GetClientSecret returns the ClientSecret field if non-nil, zero value otherwise.

### GetClientSecretOk

`func (o *PartnerOauthRequest) GetClientSecretOk() (*string, bool)`

GetClientSecretOk returns a tuple with the ClientSecret field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientSecret

`func (o *PartnerOauthRequest) SetClientSecret(v string)`

SetClientSecret sets ClientSecret field to given value.


### GetGrantType

`func (o *PartnerOauthRequest) GetGrantType() string`

GetGrantType returns the GrantType field if non-nil, zero value otherwise.

### GetGrantTypeOk

`func (o *PartnerOauthRequest) GetGrantTypeOk() (*string, bool)`

GetGrantTypeOk returns a tuple with the GrantType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGrantType

`func (o *PartnerOauthRequest) SetGrantType(v string)`

SetGrantType sets GrantType field to given value.


### GetScope

`func (o *PartnerOauthRequest) GetScope() string`

GetScope returns the Scope field if non-nil, zero value otherwise.

### GetScopeOk

`func (o *PartnerOauthRequest) GetScopeOk() (*string, bool)`

GetScopeOk returns a tuple with the Scope field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScope

`func (o *PartnerOauthRequest) SetScope(v string)`

SetScope sets Scope field to given value.

### HasScope

`func (o *PartnerOauthRequest) HasScope() bool`

HasScope returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


