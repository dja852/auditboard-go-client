# TeamsPostRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Team** | Pointer to [**Teams**](Teams.md) |  | [optional] 

## Methods

### NewTeamsPostRequest

`func NewTeamsPostRequest() *TeamsPostRequest`

NewTeamsPostRequest instantiates a new TeamsPostRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTeamsPostRequestWithDefaults

`func NewTeamsPostRequestWithDefaults() *TeamsPostRequest`

NewTeamsPostRequestWithDefaults instantiates a new TeamsPostRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTeam

`func (o *TeamsPostRequest) GetTeam() Teams`

GetTeam returns the Team field if non-nil, zero value otherwise.

### GetTeamOk

`func (o *TeamsPostRequest) GetTeamOk() (*Teams, bool)`

GetTeamOk returns a tuple with the Team field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTeam

`func (o *TeamsPostRequest) SetTeam(v Teams)`

SetTeam sets Team field to given value.

### HasTeam

`func (o *TeamsPostRequest) HasTeam() bool`

HasTeam returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


