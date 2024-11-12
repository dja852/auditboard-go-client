# TestStop

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Tests** | Pointer to [**[]Tests**](Tests.md) |  | [optional] 
**GlobalAudits** | Pointer to [**[]GlobalAudits**](GlobalAudits.md) |  | [optional] 

## Methods

### NewTestStop

`func NewTestStop() *TestStop`

NewTestStop instantiates a new TestStop object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTestStopWithDefaults

`func NewTestStopWithDefaults() *TestStop`

NewTestStopWithDefaults instantiates a new TestStop object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTests

`func (o *TestStop) GetTests() []Tests`

GetTests returns the Tests field if non-nil, zero value otherwise.

### GetTestsOk

`func (o *TestStop) GetTestsOk() (*[]Tests, bool)`

GetTestsOk returns a tuple with the Tests field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTests

`func (o *TestStop) SetTests(v []Tests)`

SetTests sets Tests field to given value.

### HasTests

`func (o *TestStop) HasTests() bool`

HasTests returns a boolean if a field has been set.

### GetGlobalAudits

`func (o *TestStop) GetGlobalAudits() []GlobalAudits`

GetGlobalAudits returns the GlobalAudits field if non-nil, zero value otherwise.

### GetGlobalAuditsOk

`func (o *TestStop) GetGlobalAuditsOk() (*[]GlobalAudits, bool)`

GetGlobalAuditsOk returns a tuple with the GlobalAudits field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGlobalAudits

`func (o *TestStop) SetGlobalAudits(v []GlobalAudits)`

SetGlobalAudits sets GlobalAudits field to given value.

### HasGlobalAudits

`func (o *TestStop) HasGlobalAudits() bool`

HasGlobalAudits returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


