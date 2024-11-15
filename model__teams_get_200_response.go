/*
AuditBoard Developer Portal API Documentation

External API reference documentation.

API version: 23.8.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package auditboard

import (
	"encoding/json"
)

// checks if the TeamsGet200Response type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &TeamsGet200Response{}

// TeamsGet200Response struct for TeamsGet200Response
type TeamsGet200Response struct {
	Teams []Teams `json:"teams,omitempty"`
}

// NewTeamsGet200Response instantiates a new TeamsGet200Response object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTeamsGet200Response() *TeamsGet200Response {
	this := TeamsGet200Response{}
	return &this
}

// NewTeamsGet200ResponseWithDefaults instantiates a new TeamsGet200Response object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTeamsGet200ResponseWithDefaults() *TeamsGet200Response {
	this := TeamsGet200Response{}
	return &this
}

// GetTeams returns the Teams field value if set, zero value otherwise.
func (o *TeamsGet200Response) GetTeams() []Teams {
	if o == nil || IsNil(o.Teams) {
		var ret []Teams
		return ret
	}
	return o.Teams
}

// GetTeamsOk returns a tuple with the Teams field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TeamsGet200Response) GetTeamsOk() ([]Teams, bool) {
	if o == nil || IsNil(o.Teams) {
		return nil, false
	}
	return o.Teams, true
}

// HasTeams returns a boolean if a field has been set.
func (o *TeamsGet200Response) HasTeams() bool {
	if o != nil && !IsNil(o.Teams) {
		return true
	}

	return false
}

// SetTeams gets a reference to the given []Teams and assigns it to the Teams field.
func (o *TeamsGet200Response) SetTeams(v []Teams) {
	o.Teams = v
}

func (o TeamsGet200Response) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o TeamsGet200Response) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Teams) {
		toSerialize["teams"] = o.Teams
	}
	return toSerialize, nil
}

type NullableTeamsGet200Response struct {
	value *TeamsGet200Response
	isSet bool
}

func (v NullableTeamsGet200Response) Get() *TeamsGet200Response {
	return v.value
}

func (v *NullableTeamsGet200Response) Set(val *TeamsGet200Response) {
	v.value = val
	v.isSet = true
}

func (v NullableTeamsGet200Response) IsSet() bool {
	return v.isSet
}

func (v *NullableTeamsGet200Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTeamsGet200Response(val *TeamsGet200Response) *NullableTeamsGet200Response {
	return &NullableTeamsGet200Response{value: val, isSet: true}
}

func (v NullableTeamsGet200Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTeamsGet200Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


