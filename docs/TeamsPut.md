# TeamsPut

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Team** | Pointer to [**TeamsPutTeam**](TeamsPutTeam.md) |  | [optional] 
**PreviousValues** | Pointer to [**TeamsPutPreviousValues**](TeamsPutPreviousValues.md) |  | [optional] 

## Methods

### NewTeamsPut

`func NewTeamsPut() *TeamsPut`

NewTeamsPut instantiates a new TeamsPut object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTeamsPutWithDefaults

`func NewTeamsPutWithDefaults() *TeamsPut`

NewTeamsPutWithDefaults instantiates a new TeamsPut object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTeam

`func (o *TeamsPut) GetTeam() TeamsPutTeam`

GetTeam returns the Team field if non-nil, zero value otherwise.

### GetTeamOk

`func (o *TeamsPut) GetTeamOk() (*TeamsPutTeam, bool)`

GetTeamOk returns a tuple with the Team field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTeam

`func (o *TeamsPut) SetTeam(v TeamsPutTeam)`

SetTeam sets Team field to given value.

### HasTeam

`func (o *TeamsPut) HasTeam() bool`

HasTeam returns a boolean if a field has been set.

### GetPreviousValues

`func (o *TeamsPut) GetPreviousValues() TeamsPutPreviousValues`

GetPreviousValues returns the PreviousValues field if non-nil, zero value otherwise.

### GetPreviousValuesOk

`func (o *TeamsPut) GetPreviousValuesOk() (*TeamsPutPreviousValues, bool)`

GetPreviousValuesOk returns a tuple with the PreviousValues field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPreviousValues

`func (o *TeamsPut) SetPreviousValues(v TeamsPutPreviousValues)`

SetPreviousValues sets PreviousValues field to given value.

### HasPreviousValues

`func (o *TeamsPut) HasPreviousValues() bool`

HasPreviousValues returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


