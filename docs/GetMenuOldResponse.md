# GetMenuOldResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**MerchantID** | Pointer to **string** | The merchant&#39;s ID that is in GrabFood&#39;s database. | [optional] 
**PartnerMerchantID** | Pointer to **string** | The merchant&#39;s ID that is on the partner&#39;s database. | [optional] 
**Currency** | [**Currency**](Currency.md) |  | 
**Sections** | [**[]MenuSection**](MenuSection.md) | An array of section JSON objects. Max 7 allowed. Refer to [Sections](#sections) for more information. | 

## Methods

### NewGetMenuOldResponse

`func NewGetMenuOldResponse(currency Currency, sections []MenuSection, ) *GetMenuOldResponse`

NewGetMenuOldResponse instantiates a new GetMenuOldResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetMenuOldResponseWithDefaults

`func NewGetMenuOldResponseWithDefaults() *GetMenuOldResponse`

NewGetMenuOldResponseWithDefaults instantiates a new GetMenuOldResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMerchantID

`func (o *GetMenuOldResponse) GetMerchantID() string`

GetMerchantID returns the MerchantID field if non-nil, zero value otherwise.

### GetMerchantIDOk

`func (o *GetMenuOldResponse) GetMerchantIDOk() (*string, bool)`

GetMerchantIDOk returns a tuple with the MerchantID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMerchantID

`func (o *GetMenuOldResponse) SetMerchantID(v string)`

SetMerchantID sets MerchantID field to given value.

### HasMerchantID

`func (o *GetMenuOldResponse) HasMerchantID() bool`

HasMerchantID returns a boolean if a field has been set.

### GetPartnerMerchantID

`func (o *GetMenuOldResponse) GetPartnerMerchantID() string`

GetPartnerMerchantID returns the PartnerMerchantID field if non-nil, zero value otherwise.

### GetPartnerMerchantIDOk

`func (o *GetMenuOldResponse) GetPartnerMerchantIDOk() (*string, bool)`

GetPartnerMerchantIDOk returns a tuple with the PartnerMerchantID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPartnerMerchantID

`func (o *GetMenuOldResponse) SetPartnerMerchantID(v string)`

SetPartnerMerchantID sets PartnerMerchantID field to given value.

### HasPartnerMerchantID

`func (o *GetMenuOldResponse) HasPartnerMerchantID() bool`

HasPartnerMerchantID returns a boolean if a field has been set.

### GetCurrency

`func (o *GetMenuOldResponse) GetCurrency() Currency`

GetCurrency returns the Currency field if non-nil, zero value otherwise.

### GetCurrencyOk

`func (o *GetMenuOldResponse) GetCurrencyOk() (*Currency, bool)`

GetCurrencyOk returns a tuple with the Currency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrency

`func (o *GetMenuOldResponse) SetCurrency(v Currency)`

SetCurrency sets Currency field to given value.


### GetSections

`func (o *GetMenuOldResponse) GetSections() []MenuSection`

GetSections returns the Sections field if non-nil, zero value otherwise.

### GetSectionsOk

`func (o *GetMenuOldResponse) GetSectionsOk() (*[]MenuSection, bool)`

GetSectionsOk returns a tuple with the Sections field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSections

`func (o *GetMenuOldResponse) SetSections(v []MenuSection)`

SetSections sets Sections field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


