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

// checks if the TeamsPut type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &TeamsPut{}

// TeamsPut struct for TeamsPut
type TeamsPut struct {
	Team *TeamsPutTeam `json:"team,omitempty"`
	PreviousValues *TeamsPutPreviousValues `json:"previous_values,omitempty"`
}

// NewTeamsPut instantiates a new TeamsPut object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTeamsPut() *TeamsPut {
	this := TeamsPut{}
	return &this
}

// NewTeamsPutWithDefaults instantiates a new TeamsPut object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTeamsPutWithDefaults() *TeamsPut {
	this := TeamsPut{}
	return &this
}

// GetTeam returns the Team field value if set, zero value otherwise.
func (o *TeamsPut) GetTeam() TeamsPutTeam {
	if o == nil || IsNil(o.Team) {
		var ret TeamsPutTeam
		return ret
	}
	return *o.Team
}

// GetTeamOk returns a tuple with the Team field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TeamsPut) GetTeamOk() (*TeamsPutTeam, bool) {
	if o == nil || IsNil(o.Team) {
		return nil, false
	}
	return o.Team, true
}

// HasTeam returns a boolean if a field has been set.
func (o *TeamsPut) HasTeam() bool {
	if o != nil && !IsNil(o.Team) {
		return true
	}

	return false
}

// SetTeam gets a reference to the given TeamsPutTeam and assigns it to the Team field.
func (o *TeamsPut) SetTeam(v TeamsPutTeam) {
	o.Team = &v
}

// GetPreviousValues returns the PreviousValues field value if set, zero value otherwise.
func (o *TeamsPut) GetPreviousValues() TeamsPutPreviousValues {
	if o == nil || IsNil(o.PreviousValues) {
		var ret TeamsPutPreviousValues
		return ret
	}
	return *o.PreviousValues
}

// GetPreviousValuesOk returns a tuple with the PreviousValues field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TeamsPut) GetPreviousValuesOk() (*TeamsPutPreviousValues, bool) {
	if o == nil || IsNil(o.PreviousValues) {
		return nil, false
	}
	return o.PreviousValues, true
}

// HasPreviousValues returns a boolean if a field has been set.
func (o *TeamsPut) HasPreviousValues() bool {
	if o != nil && !IsNil(o.PreviousValues) {
		return true
	}

	return false
}

// SetPreviousValues gets a reference to the given TeamsPutPreviousValues and assigns it to the PreviousValues field.
func (o *TeamsPut) SetPreviousValues(v TeamsPutPreviousValues) {
	o.PreviousValues = &v
}

func (o TeamsPut) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o TeamsPut) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Team) {
		toSerialize["team"] = o.Team
	}
	if !IsNil(o.PreviousValues) {
		toSerialize["previous_values"] = o.PreviousValues
	}
	return toSerialize, nil
}

type NullableTeamsPut struct {
	value *TeamsPut
	isSet bool
}

func (v NullableTeamsPut) Get() *TeamsPut {
	return v.value
}

func (v *NullableTeamsPut) Set(val *TeamsPut) {
	v.value = val
	v.isSet = true
}

func (v NullableTeamsPut) IsSet() bool {
	return v.isSet
}

func (v *NullableTeamsPut) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTeamsPut(val *TeamsPut) *NullableTeamsPut {
	return &NullableTeamsPut{value: val, isSet: true}
}

func (v NullableTeamsPut) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTeamsPut) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


