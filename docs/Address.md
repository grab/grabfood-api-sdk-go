# Address

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**UnitNumber** | Pointer to **string** | The delivery address&#39; unit number. | [optional] 
**DeliveryInstruction** | Pointer to **string** | Instructions for the delivery. | [optional] 
**PoiSource** | Pointer to **string** | POI source | [optional] 
**PoiID** | Pointer to **string** | POI ID, empty when poiSource is GRAB. | [optional] 
**Address** | Pointer to **string** | The delivery address containing street name, city, postal code, and country. Has value only when poiSource is &#x60;GRAB&#x60;. | [optional] 
**Postcode** | Pointer to **string** | The postcode of the delivery address. Has value only when poiSource is &#x60;GRAB&#x60;. | [optional] 
**Coordinates** | Pointer to [**Coordinates**](Coordinates.md) |  | [optional] 

## Methods

### NewAddress

`func NewAddress() *Address`

NewAddress instantiates a new Address object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAddressWithDefaults

`func NewAddressWithDefaults() *Address`

NewAddressWithDefaults instantiates a new Address object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUnitNumber

`func (o *Address) GetUnitNumber() string`

GetUnitNumber returns the UnitNumber field if non-nil, zero value otherwise.

### GetUnitNumberOk

`func (o *Address) GetUnitNumberOk() (*string, bool)`

GetUnitNumberOk returns a tuple with the UnitNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnitNumber

`func (o *Address) SetUnitNumber(v string)`

SetUnitNumber sets UnitNumber field to given value.

### HasUnitNumber

`func (o *Address) HasUnitNumber() bool`

HasUnitNumber returns a boolean if a field has been set.

### GetDeliveryInstruction

`func (o *Address) GetDeliveryInstruction() string`

GetDeliveryInstruction returns the DeliveryInstruction field if non-nil, zero value otherwise.

### GetDeliveryInstructionOk

`func (o *Address) GetDeliveryInstructionOk() (*string, bool)`

GetDeliveryInstructionOk returns a tuple with the DeliveryInstruction field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeliveryInstruction

`func (o *Address) SetDeliveryInstruction(v string)`

SetDeliveryInstruction sets DeliveryInstruction field to given value.

### HasDeliveryInstruction

`func (o *Address) HasDeliveryInstruction() bool`

HasDeliveryInstruction returns a boolean if a field has been set.

### GetPoiSource

`func (o *Address) GetPoiSource() string`

GetPoiSource returns the PoiSource field if non-nil, zero value otherwise.

### GetPoiSourceOk

`func (o *Address) GetPoiSourceOk() (*string, bool)`

GetPoiSourceOk returns a tuple with the PoiSource field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPoiSource

`func (o *Address) SetPoiSource(v string)`

SetPoiSource sets PoiSource field to given value.

### HasPoiSource

`func (o *Address) HasPoiSource() bool`

HasPoiSource returns a boolean if a field has been set.

### GetPoiID

`func (o *Address) GetPoiID() string`

GetPoiID returns the PoiID field if non-nil, zero value otherwise.

### GetPoiIDOk

`func (o *Address) GetPoiIDOk() (*string, bool)`

GetPoiIDOk returns a tuple with the PoiID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPoiID

`func (o *Address) SetPoiID(v string)`

SetPoiID sets PoiID field to given value.

### HasPoiID

`func (o *Address) HasPoiID() bool`

HasPoiID returns a boolean if a field has been set.

### GetAddress

`func (o *Address) GetAddress() string`

GetAddress returns the Address field if non-nil, zero value otherwise.

### GetAddressOk

`func (o *Address) GetAddressOk() (*string, bool)`

GetAddressOk returns a tuple with the Address field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddress

`func (o *Address) SetAddress(v string)`

SetAddress sets Address field to given value.

### HasAddress

`func (o *Address) HasAddress() bool`

HasAddress returns a boolean if a field has been set.

### GetPostcode

`func (o *Address) GetPostcode() string`

GetPostcode returns the Postcode field if non-nil, zero value otherwise.

### GetPostcodeOk

`func (o *Address) GetPostcodeOk() (*string, bool)`

GetPostcodeOk returns a tuple with the Postcode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPostcode

`func (o *Address) SetPostcode(v string)`

SetPostcode sets Postcode field to given value.

### HasPostcode

`func (o *Address) HasPostcode() bool`

HasPostcode returns a boolean if a field has been set.

### GetCoordinates

`func (o *Address) GetCoordinates() Coordinates`

GetCoordinates returns the Coordinates field if non-nil, zero value otherwise.

### GetCoordinatesOk

`func (o *Address) GetCoordinatesOk() (*Coordinates, bool)`

GetCoordinatesOk returns a tuple with the Coordinates field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCoordinates

`func (o *Address) SetCoordinates(v Coordinates)`

SetCoordinates sets Coordinates field to given value.

### HasCoordinates

`func (o *Address) HasCoordinates() bool`

HasCoordinates returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


