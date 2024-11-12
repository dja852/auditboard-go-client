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

// checks if the SubprocessesDataPut type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SubprocessesDataPut{}

// SubprocessesDataPut struct for SubprocessesDataPut
type SubprocessesDataPut struct {
	SubprocessesDatum *SubprocessesDataPutSubprocessesDatum `json:"subprocesses_datum,omitempty"`
	PreviousValues *SubprocessesDataPutPreviousValues `json:"previous_values,omitempty"`
}

// NewSubprocessesDataPut instantiates a new SubprocessesDataPut object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSubprocessesDataPut() *SubprocessesDataPut {
	this := SubprocessesDataPut{}
	return &this
}

// NewSubprocessesDataPutWithDefaults instantiates a new SubprocessesDataPut object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSubprocessesDataPutWithDefaults() *SubprocessesDataPut {
	this := SubprocessesDataPut{}
	return &this
}

// GetSubprocessesDatum returns the SubprocessesDatum field value if set, zero value otherwise.
func (o *SubprocessesDataPut) GetSubprocessesDatum() SubprocessesDataPutSubprocessesDatum {
	if o == nil || IsNil(o.SubprocessesDatum) {
		var ret SubprocessesDataPutSubprocessesDatum
		return ret
	}
	return *o.SubprocessesDatum
}

// GetSubprocessesDatumOk returns a tuple with the SubprocessesDatum field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SubprocessesDataPut) GetSubprocessesDatumOk() (*SubprocessesDataPutSubprocessesDatum, bool) {
	if o == nil || IsNil(o.SubprocessesDatum) {
		return nil, false
	}
	return o.SubprocessesDatum, true
}

// HasSubprocessesDatum returns a boolean if a field has been set.
func (o *SubprocessesDataPut) HasSubprocessesDatum() bool {
	if o != nil && !IsNil(o.SubprocessesDatum) {
		return true
	}

	return false
}

// SetSubprocessesDatum gets a reference to the given SubprocessesDataPutSubprocessesDatum and assigns it to the SubprocessesDatum field.
func (o *SubprocessesDataPut) SetSubprocessesDatum(v SubprocessesDataPutSubprocessesDatum) {
	o.SubprocessesDatum = &v
}

// GetPreviousValues returns the PreviousValues field value if set, zero value otherwise.
func (o *SubprocessesDataPut) GetPreviousValues() SubprocessesDataPutPreviousValues {
	if o == nil || IsNil(o.PreviousValues) {
		var ret SubprocessesDataPutPreviousValues
		return ret
	}
	return *o.PreviousValues
}

// GetPreviousValuesOk returns a tuple with the PreviousValues field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SubprocessesDataPut) GetPreviousValuesOk() (*SubprocessesDataPutPreviousValues, bool) {
	if o == nil || IsNil(o.PreviousValues) {
		return nil, false
	}
	return o.PreviousValues, true
}

// HasPreviousValues returns a boolean if a field has been set.
func (o *SubprocessesDataPut) HasPreviousValues() bool {
	if o != nil && !IsNil(o.PreviousValues) {
		return true
	}

	return false
}

// SetPreviousValues gets a reference to the given SubprocessesDataPutPreviousValues and assigns it to the PreviousValues field.
func (o *SubprocessesDataPut) SetPreviousValues(v SubprocessesDataPutPreviousValues) {
	o.PreviousValues = &v
}

func (o SubprocessesDataPut) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SubprocessesDataPut) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.SubprocessesDatum) {
		toSerialize["subprocesses_datum"] = o.SubprocessesDatum
	}
	if !IsNil(o.PreviousValues) {
		toSerialize["previous_values"] = o.PreviousValues
	}
	return toSerialize, nil
}

type NullableSubprocessesDataPut struct {
	value *SubprocessesDataPut
	isSet bool
}

func (v NullableSubprocessesDataPut) Get() *SubprocessesDataPut {
	return v.value
}

func (v *NullableSubprocessesDataPut) Set(val *SubprocessesDataPut) {
	v.value = val
	v.isSet = true
}

func (v NullableSubprocessesDataPut) IsSet() bool {
	return v.isSet
}

func (v *NullableSubprocessesDataPut) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSubprocessesDataPut(val *SubprocessesDataPut) *NullableSubprocessesDataPut {
	return &NullableSubprocessesDataPut{value: val, isSet: true}
}

func (v NullableSubprocessesDataPut) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSubprocessesDataPut) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


