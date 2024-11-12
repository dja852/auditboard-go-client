# FilesPut

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**File** | Pointer to [**FilesPutFile**](FilesPutFile.md) |  | [optional] 
**PreviousValues** | Pointer to [**FilesPutPreviousValues**](FilesPutPreviousValues.md) |  | [optional] 

## Methods

### NewFilesPut

`func NewFilesPut() *FilesPut`

NewFilesPut instantiates a new FilesPut object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFilesPutWithDefaults

`func NewFilesPutWithDefaults() *FilesPut`

NewFilesPutWithDefaults instantiates a new FilesPut object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFile

`func (o *FilesPut) GetFile() FilesPutFile`

GetFile returns the File field if non-nil, zero value otherwise.

### GetFileOk

`func (o *FilesPut) GetFileOk() (*FilesPutFile, bool)`

GetFileOk returns a tuple with the File field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFile

`func (o *FilesPut) SetFile(v FilesPutFile)`

SetFile sets File field to given value.

### HasFile

`func (o *FilesPut) HasFile() bool`

HasFile returns a boolean if a field has been set.

### GetPreviousValues

`func (o *FilesPut) GetPreviousValues() FilesPutPreviousValues`

GetPreviousValues returns the PreviousValues field if non-nil, zero value otherwise.

### GetPreviousValuesOk

`func (o *FilesPut) GetPreviousValuesOk() (*FilesPutPreviousValues, bool)`

GetPreviousValuesOk returns a tuple with the PreviousValues field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPreviousValues

`func (o *FilesPut) SetPreviousValues(v FilesPutPreviousValues)`

SetPreviousValues sets PreviousValues field to given value.

### HasPreviousValues

`func (o *FilesPut) HasPreviousValues() bool`

HasPreviousValues returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


