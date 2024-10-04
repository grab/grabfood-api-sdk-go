# OutOfStockInstruction

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Title** | Pointer to **string** | The short instruction message. | [optional] 
**InstructionType** | Pointer to **string** | Type of out-of-stock instruction chosen by customer. &#x60;CONTACT&#x60; is disabled by default, kindly reach out to your integration manager if you wish to receive this instruction. | [optional] 
**ReplacementItemID** | Pointer to **string** | The preferred item&#39;s ID in the partner system. Only applicable when the instructionType is &#x60;SPECIFIC_ITEM&#x60;. | [optional] 
**ReplacementGrabItemID** | Pointer to **string** | The preferred item&#39;s ID in the Grab system. Only applicable when the instructionType is &#x60;SPECIFIC_ITEM&#x60;. | [optional] 

## Methods

### NewOutOfStockInstruction

`func NewOutOfStockInstruction() *OutOfStockInstruction`

NewOutOfStockInstruction instantiates a new OutOfStockInstruction object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOutOfStockInstructionWithDefaults

`func NewOutOfStockInstructionWithDefaults() *OutOfStockInstruction`

NewOutOfStockInstructionWithDefaults instantiates a new OutOfStockInstruction object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTitle

`func (o *OutOfStockInstruction) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *OutOfStockInstruction) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *OutOfStockInstruction) SetTitle(v string)`

SetTitle sets Title field to given value.

### HasTitle

`func (o *OutOfStockInstruction) HasTitle() bool`

HasTitle returns a boolean if a field has been set.

### GetInstructionType

`func (o *OutOfStockInstruction) GetInstructionType() string`

GetInstructionType returns the InstructionType field if non-nil, zero value otherwise.

### GetInstructionTypeOk

`func (o *OutOfStockInstruction) GetInstructionTypeOk() (*string, bool)`

GetInstructionTypeOk returns a tuple with the InstructionType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstructionType

`func (o *OutOfStockInstruction) SetInstructionType(v string)`

SetInstructionType sets InstructionType field to given value.

### HasInstructionType

`func (o *OutOfStockInstruction) HasInstructionType() bool`

HasInstructionType returns a boolean if a field has been set.

### GetReplacementItemID

`func (o *OutOfStockInstruction) GetReplacementItemID() string`

GetReplacementItemID returns the ReplacementItemID field if non-nil, zero value otherwise.

### GetReplacementItemIDOk

`func (o *OutOfStockInstruction) GetReplacementItemIDOk() (*string, bool)`

GetReplacementItemIDOk returns a tuple with the ReplacementItemID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReplacementItemID

`func (o *OutOfStockInstruction) SetReplacementItemID(v string)`

SetReplacementItemID sets ReplacementItemID field to given value.

### HasReplacementItemID

`func (o *OutOfStockInstruction) HasReplacementItemID() bool`

HasReplacementItemID returns a boolean if a field has been set.

### GetReplacementGrabItemID

`func (o *OutOfStockInstruction) GetReplacementGrabItemID() string`

GetReplacementGrabItemID returns the ReplacementGrabItemID field if non-nil, zero value otherwise.

### GetReplacementGrabItemIDOk

`func (o *OutOfStockInstruction) GetReplacementGrabItemIDOk() (*string, bool)`

GetReplacementGrabItemIDOk returns a tuple with the ReplacementGrabItemID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReplacementGrabItemID

`func (o *OutOfStockInstruction) SetReplacementGrabItemID(v string)`

SetReplacementGrabItemID sets ReplacementGrabItemID field to given value.

### HasReplacementGrabItemID

`func (o *OutOfStockInstruction) HasReplacementGrabItemID() bool`

HasReplacementGrabItemID returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


