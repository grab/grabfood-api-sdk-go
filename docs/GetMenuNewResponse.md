# GetMenuNewResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**MerchantID** | Pointer to **string** | The merchant&#39;s ID that is in GrabFood&#39;s database. | [optional] 
**PartnerMerchantID** | Pointer to **string** | The merchant&#39;s ID that is on the partner&#39;s database. | [optional] 
**Currency** | [**Currency**](Currency.md) |  | 
**SellingTimes** | [**[]SellingTime**](SellingTime.md) | An array of sellingTimes JSON objects. Max 20 allowed. Refer to [Selling Times](#selling-times) for more information. | 
**Categories** | [**[]MenuCategory**](MenuCategory.md) | An array of category JSON objects. Max 100 allowed per section. Refer to [Categories](#categories) for more information. | 

## Methods

### NewGetMenuNewResponse

`func NewGetMenuNewResponse(currency Currency, sellingTimes []SellingTime, categories []MenuCategory, ) *GetMenuNewResponse`

NewGetMenuNewResponse instantiates a new GetMenuNewResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetMenuNewResponseWithDefaults

`func NewGetMenuNewResponseWithDefaults() *GetMenuNewResponse`

NewGetMenuNewResponseWithDefaults instantiates a new GetMenuNewResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMerchantID

`func (o *GetMenuNewResponse) GetMerchantID() string`

GetMerchantID returns the MerchantID field if non-nil, zero value otherwise.

### GetMerchantIDOk

`func (o *GetMenuNewResponse) GetMerchantIDOk() (*string, bool)`

GetMerchantIDOk returns a tuple with the MerchantID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMerchantID

`func (o *GetMenuNewResponse) SetMerchantID(v string)`

SetMerchantID sets MerchantID field to given value.

### HasMerchantID

`func (o *GetMenuNewResponse) HasMerchantID() bool`

HasMerchantID returns a boolean if a field has been set.

### GetPartnerMerchantID

`func (o *GetMenuNewResponse) GetPartnerMerchantID() string`

GetPartnerMerchantID returns the PartnerMerchantID field if non-nil, zero value otherwise.

### GetPartnerMerchantIDOk

`func (o *GetMenuNewResponse) GetPartnerMerchantIDOk() (*string, bool)`

GetPartnerMerchantIDOk returns a tuple with the PartnerMerchantID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPartnerMerchantID

`func (o *GetMenuNewResponse) SetPartnerMerchantID(v string)`

SetPartnerMerchantID sets PartnerMerchantID field to given value.

### HasPartnerMerchantID

`func (o *GetMenuNewResponse) HasPartnerMerchantID() bool`

HasPartnerMerchantID returns a boolean if a field has been set.

### GetCurrency

`func (o *GetMenuNewResponse) GetCurrency() Currency`

GetCurrency returns the Currency field if non-nil, zero value otherwise.

### GetCurrencyOk

`func (o *GetMenuNewResponse) GetCurrencyOk() (*Currency, bool)`

GetCurrencyOk returns a tuple with the Currency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrency

`func (o *GetMenuNewResponse) SetCurrency(v Currency)`

SetCurrency sets Currency field to given value.


### GetSellingTimes

`func (o *GetMenuNewResponse) GetSellingTimes() []SellingTime`

GetSellingTimes returns the SellingTimes field if non-nil, zero value otherwise.

### GetSellingTimesOk

`func (o *GetMenuNewResponse) GetSellingTimesOk() (*[]SellingTime, bool)`

GetSellingTimesOk returns a tuple with the SellingTimes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSellingTimes

`func (o *GetMenuNewResponse) SetSellingTimes(v []SellingTime)`

SetSellingTimes sets SellingTimes field to given value.


### GetCategories

`func (o *GetMenuNewResponse) GetCategories() []MenuCategory`

GetCategories returns the Categories field if non-nil, zero value otherwise.

### GetCategoriesOk

`func (o *GetMenuNewResponse) GetCategoriesOk() (*[]MenuCategory, bool)`

GetCategoriesOk returns a tuple with the Categories field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCategories

`func (o *GetMenuNewResponse) SetCategories(v []MenuCategory)`

SetCategories sets Categories field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


