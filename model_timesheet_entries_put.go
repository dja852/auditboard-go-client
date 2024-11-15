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

// checks if the TimesheetEntriesPut type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &TimesheetEntriesPut{}

// TimesheetEntriesPut struct for TimesheetEntriesPut
type TimesheetEntriesPut struct {
	TimesheetEntry *TimesheetEntriesPutTimesheetEntry `json:"timesheet_entry,omitempty"`
	PreviousValues *TimesheetEntriesPutPreviousValues `json:"previous_values,omitempty"`
}

// NewTimesheetEntriesPut instantiates a new TimesheetEntriesPut object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTimesheetEntriesPut() *TimesheetEntriesPut {
	this := TimesheetEntriesPut{}
	return &this
}

// NewTimesheetEntriesPutWithDefaults instantiates a new TimesheetEntriesPut object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTimesheetEntriesPutWithDefaults() *TimesheetEntriesPut {
	this := TimesheetEntriesPut{}
	return &this
}

// GetTimesheetEntry returns the TimesheetEntry field value if set, zero value otherwise.
func (o *TimesheetEntriesPut) GetTimesheetEntry() TimesheetEntriesPutTimesheetEntry {
	if o == nil || IsNil(o.TimesheetEntry) {
		var ret TimesheetEntriesPutTimesheetEntry
		return ret
	}
	return *o.TimesheetEntry
}

// GetTimesheetEntryOk returns a tuple with the TimesheetEntry field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TimesheetEntriesPut) GetTimesheetEntryOk() (*TimesheetEntriesPutTimesheetEntry, bool) {
	if o == nil || IsNil(o.TimesheetEntry) {
		return nil, false
	}
	return o.TimesheetEntry, true
}

// HasTimesheetEntry returns a boolean if a field has been set.
func (o *TimesheetEntriesPut) HasTimesheetEntry() bool {
	if o != nil && !IsNil(o.TimesheetEntry) {
		return true
	}

	return false
}

// SetTimesheetEntry gets a reference to the given TimesheetEntriesPutTimesheetEntry and assigns it to the TimesheetEntry field.
func (o *TimesheetEntriesPut) SetTimesheetEntry(v TimesheetEntriesPutTimesheetEntry) {
	o.TimesheetEntry = &v
}

// GetPreviousValues returns the PreviousValues field value if set, zero value otherwise.
func (o *TimesheetEntriesPut) GetPreviousValues() TimesheetEntriesPutPreviousValues {
	if o == nil || IsNil(o.PreviousValues) {
		var ret TimesheetEntriesPutPreviousValues
		return ret
	}
	return *o.PreviousValues
}

// GetPreviousValuesOk returns a tuple with the PreviousValues field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TimesheetEntriesPut) GetPreviousValuesOk() (*TimesheetEntriesPutPreviousValues, bool) {
	if o == nil || IsNil(o.PreviousValues) {
		return nil, false
	}
	return o.PreviousValues, true
}

// HasPreviousValues returns a boolean if a field has been set.
func (o *TimesheetEntriesPut) HasPreviousValues() bool {
	if o != nil && !IsNil(o.PreviousValues) {
		return true
	}

	return false
}

// SetPreviousValues gets a reference to the given TimesheetEntriesPutPreviousValues and assigns it to the PreviousValues field.
func (o *TimesheetEntriesPut) SetPreviousValues(v TimesheetEntriesPutPreviousValues) {
	o.PreviousValues = &v
}

func (o TimesheetEntriesPut) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o TimesheetEntriesPut) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.TimesheetEntry) {
		toSerialize["timesheet_entry"] = o.TimesheetEntry
	}
	if !IsNil(o.PreviousValues) {
		toSerialize["previous_values"] = o.PreviousValues
	}
	return toSerialize, nil
}

type NullableTimesheetEntriesPut struct {
	value *TimesheetEntriesPut
	isSet bool
}

func (v NullableTimesheetEntriesPut) Get() *TimesheetEntriesPut {
	return v.value
}

func (v *NullableTimesheetEntriesPut) Set(val *TimesheetEntriesPut) {
	v.value = val
	v.isSet = true
}

func (v NullableTimesheetEntriesPut) IsSet() bool {
	return v.isSet
}

func (v *NullableTimesheetEntriesPut) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTimesheetEntriesPut(val *TimesheetEntriesPut) *NullableTimesheetEntriesPut {
	return &NullableTimesheetEntriesPut{value: val, isSet: true}
}

func (v NullableTimesheetEntriesPut) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTimesheetEntriesPut) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


