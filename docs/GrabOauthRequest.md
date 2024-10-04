# GrabOauthRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClientId** | **string** | The client identifier issued to the client to obtain the OAuth 2.0 access_token. | 
**ClientSecret** | **string** | The client secret issued to the client to obtain the OAuth 2.0 access_token. | 
**GrantType** | **string** | The grant type for the client to obtain the OAuth 2.0 access_token. | 
**Scope** | **string** | The scope for the client to obtain the OAuth 2.0 access_token. | 

## Methods

### NewGrabOauthRequest

`func NewGrabOauthRequest(clientId string, clientSecret string, grantType string, scope string, ) *GrabOauthRequest`

NewGrabOauthRequest instantiates a new GrabOauthRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGrabOauthRequestWithDefaults

`func NewGrabOauthRequestWithDefaults() *GrabOauthRequest`

NewGrabOauthRequestWithDefaults instantiates a new GrabOauthRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClientId

`func (o *GrabOauthRequest) GetClientId() string`

GetClientId returns the ClientId field if non-nil, zero value otherwise.

### GetClientIdOk

`func (o *GrabOauthRequest) GetClientIdOk() (*string, bool)`

GetClientIdOk returns a tuple with the ClientId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientId

`func (o *GrabOauthRequest) SetClientId(v string)`

SetClientId sets ClientId field to given value.


### GetClientSecret

`func (o *GrabOauthRequest) GetClientSecret() string`

GetClientSecret returns the ClientSecret field if non-nil, zero value otherwise.

### GetClientSecretOk

`func (o *GrabOauthRequest) GetClientSecretOk() (*string, bool)`

GetClientSecretOk returns a tuple with the ClientSecret field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientSecret

`func (o *GrabOauthRequest) SetClientSecret(v string)`

SetClientSecret sets ClientSecret field to given value.


### GetGrantType

`func (o *GrabOauthRequest) GetGrantType() string`

GetGrantType returns the GrantType field if non-nil, zero value otherwise.

### GetGrantTypeOk

`func (o *GrabOauthRequest) GetGrantTypeOk() (*string, bool)`

GetGrantTypeOk returns a tuple with the GrantType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGrantType

`func (o *GrabOauthRequest) SetGrantType(v string)`

SetGrantType sets GrantType field to given value.


### GetScope

`func (o *GrabOauthRequest) GetScope() string`

GetScope returns the Scope field if non-nil, zero value otherwise.

### GetScopeOk

`func (o *GrabOauthRequest) GetScopeOk() (*string, bool)`

GetScopeOk returns a tuple with the Scope field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScope

`func (o *GrabOauthRequest) SetScope(v string)`

SetScope sets Scope field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


