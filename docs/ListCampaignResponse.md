# ListCampaignResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Ongoing** | Pointer to [**[]Campaign**](Campaign.md) |  | [optional] 
**Upcoming** | Pointer to [**[]Campaign**](Campaign.md) |  | [optional] 

## Methods

### NewListCampaignResponse

`func NewListCampaignResponse() *ListCampaignResponse`

NewListCampaignResponse instantiates a new ListCampaignResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListCampaignResponseWithDefaults

`func NewListCampaignResponseWithDefaults() *ListCampaignResponse`

NewListCampaignResponseWithDefaults instantiates a new ListCampaignResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetOngoing

`func (o *ListCampaignResponse) GetOngoing() []Campaign`

GetOngoing returns the Ongoing field if non-nil, zero value otherwise.

### GetOngoingOk

`func (o *ListCampaignResponse) GetOngoingOk() (*[]Campaign, bool)`

GetOngoingOk returns a tuple with the Ongoing field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOngoing

`func (o *ListCampaignResponse) SetOngoing(v []Campaign)`

SetOngoing sets Ongoing field to given value.

### HasOngoing

`func (o *ListCampaignResponse) HasOngoing() bool`

HasOngoing returns a boolean if a field has been set.

### GetUpcoming

`func (o *ListCampaignResponse) GetUpcoming() []Campaign`

GetUpcoming returns the Upcoming field if non-nil, zero value otherwise.

### GetUpcomingOk

`func (o *ListCampaignResponse) GetUpcomingOk() (*[]Campaign, bool)`

GetUpcomingOk returns a tuple with the Upcoming field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpcoming

`func (o *ListCampaignResponse) SetUpcoming(v []Campaign)`

SetUpcoming sets Upcoming field to given value.

### HasUpcoming

`func (o *ListCampaignResponse) HasUpcoming() bool`

HasUpcoming returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


