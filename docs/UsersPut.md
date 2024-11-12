# UsersPut

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**User** | Pointer to [**UsersPutUser**](UsersPutUser.md) |  | [optional] 
**PreviousValues** | Pointer to [**UsersPutPreviousValues**](UsersPutPreviousValues.md) |  | [optional] 

## Methods

### NewUsersPut

`func NewUsersPut() *UsersPut`

NewUsersPut instantiates a new UsersPut object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUsersPutWithDefaults

`func NewUsersPutWithDefaults() *UsersPut`

NewUsersPutWithDefaults instantiates a new UsersPut object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUser

`func (o *UsersPut) GetUser() UsersPutUser`

GetUser returns the User field if non-nil, zero value otherwise.

### GetUserOk

`func (o *UsersPut) GetUserOk() (*UsersPutUser, bool)`

GetUserOk returns a tuple with the User field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUser

`func (o *UsersPut) SetUser(v UsersPutUser)`

SetUser sets User field to given value.

### HasUser

`func (o *UsersPut) HasUser() bool`

HasUser returns a boolean if a field has been set.

### GetPreviousValues

`func (o *UsersPut) GetPreviousValues() UsersPutPreviousValues`

GetPreviousValues returns the PreviousValues field if non-nil, zero value otherwise.

### GetPreviousValuesOk

`func (o *UsersPut) GetPreviousValuesOk() (*UsersPutPreviousValues, bool)`

GetPreviousValuesOk returns a tuple with the PreviousValues field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPreviousValues

`func (o *UsersPut) SetPreviousValues(v UsersPutPreviousValues)`

SetPreviousValues sets PreviousValues field to given value.

### HasPreviousValues

`func (o *UsersPut) HasPreviousValues() bool`

HasPreviousValues returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


