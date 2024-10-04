# StoreStatusResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CloseReason** | **string** | The code of store close reason. Blank indicates store is currently open. | 
**IsInSpecialOpeningHourRange** | **bool** | Indicate whether the store is in special opening hour range. | 
**IsOpen** | **bool** | Indicate whether the store is open. | 

## Methods

### NewStoreStatusResponse

`func NewStoreStatusResponse(closeReason string, isInSpecialOpeningHourRange bool, isOpen bool, ) *StoreStatusResponse`

NewStoreStatusResponse instantiates a new StoreStatusResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewStoreStatusResponseWithDefaults

`func NewStoreStatusResponseWithDefaults() *StoreStatusResponse`

NewStoreStatusResponseWithDefaults instantiates a new StoreStatusResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCloseReason

`func (o *StoreStatusResponse) GetCloseReason() string`

GetCloseReason returns the CloseReason field if non-nil, zero value otherwise.

### GetCloseReasonOk

`func (o *StoreStatusResponse) GetCloseReasonOk() (*string, bool)`

GetCloseReasonOk returns a tuple with the CloseReason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCloseReason

`func (o *StoreStatusResponse) SetCloseReason(v string)`

SetCloseReason sets CloseReason field to given value.


### GetIsInSpecialOpeningHourRange

`func (o *StoreStatusResponse) GetIsInSpecialOpeningHourRange() bool`

GetIsInSpecialOpeningHourRange returns the IsInSpecialOpeningHourRange field if non-nil, zero value otherwise.

### GetIsInSpecialOpeningHourRangeOk

`func (o *StoreStatusResponse) GetIsInSpecialOpeningHourRangeOk() (*bool, bool)`

GetIsInSpecialOpeningHourRangeOk returns a tuple with the IsInSpecialOpeningHourRange field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsInSpecialOpeningHourRange

`func (o *StoreStatusResponse) SetIsInSpecialOpeningHourRange(v bool)`

SetIsInSpecialOpeningHourRange sets IsInSpecialOpeningHourRange field to given value.


### GetIsOpen

`func (o *StoreStatusResponse) GetIsOpen() bool`

GetIsOpen returns the IsOpen field if non-nil, zero value otherwise.

### GetIsOpenOk

`func (o *StoreStatusResponse) GetIsOpenOk() (*bool, bool)`

GetIsOpenOk returns a tuple with the IsOpen field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsOpen

`func (o *StoreStatusResponse) SetIsOpen(v bool)`

SetIsOpen sets IsOpen field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


