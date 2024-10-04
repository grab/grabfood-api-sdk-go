# ListOrdersResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**More** | Pointer to **bool** | The boolean value to indicate if there is more order data. If &#x60;true&#x60;, you can continue requesting with page+1. | [optional] 
**Orders** | Pointer to [**[]Order**](Order.md) |  | [optional] 

## Methods

### NewListOrdersResponse

`func NewListOrdersResponse() *ListOrdersResponse`

NewListOrdersResponse instantiates a new ListOrdersResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListOrdersResponseWithDefaults

`func NewListOrdersResponseWithDefaults() *ListOrdersResponse`

NewListOrdersResponseWithDefaults instantiates a new ListOrdersResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMore

`func (o *ListOrdersResponse) GetMore() bool`

GetMore returns the More field if non-nil, zero value otherwise.

### GetMoreOk

`func (o *ListOrdersResponse) GetMoreOk() (*bool, bool)`

GetMoreOk returns a tuple with the More field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMore

`func (o *ListOrdersResponse) SetMore(v bool)`

SetMore sets More field to given value.

### HasMore

`func (o *ListOrdersResponse) HasMore() bool`

HasMore returns a boolean if a field has been set.

### GetOrders

`func (o *ListOrdersResponse) GetOrders() []Order`

GetOrders returns the Orders field if non-nil, zero value otherwise.

### GetOrdersOk

`func (o *ListOrdersResponse) GetOrdersOk() (*[]Order, bool)`

GetOrdersOk returns a tuple with the Orders field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrders

`func (o *ListOrdersResponse) SetOrders(v []Order)`

SetOrders sets Orders field to given value.

### HasOrders

`func (o *ListOrdersResponse) HasOrders() bool`

HasOrders returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


