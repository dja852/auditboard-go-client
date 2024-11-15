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

// checks if the DepartmentsPostRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &DepartmentsPostRequest{}

// DepartmentsPostRequest struct for DepartmentsPostRequest
type DepartmentsPostRequest struct {
	Department *Departments `json:"department,omitempty"`
}

// NewDepartmentsPostRequest instantiates a new DepartmentsPostRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDepartmentsPostRequest() *DepartmentsPostRequest {
	this := DepartmentsPostRequest{}
	return &this
}

// NewDepartmentsPostRequestWithDefaults instantiates a new DepartmentsPostRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDepartmentsPostRequestWithDefaults() *DepartmentsPostRequest {
	this := DepartmentsPostRequest{}
	return &this
}

// GetDepartment returns the Department field value if set, zero value otherwise.
func (o *DepartmentsPostRequest) GetDepartment() Departments {
	if o == nil || IsNil(o.Department) {
		var ret Departments
		return ret
	}
	return *o.Department
}

// GetDepartmentOk returns a tuple with the Department field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DepartmentsPostRequest) GetDepartmentOk() (*Departments, bool) {
	if o == nil || IsNil(o.Department) {
		return nil, false
	}
	return o.Department, true
}

// HasDepartment returns a boolean if a field has been set.
func (o *DepartmentsPostRequest) HasDepartment() bool {
	if o != nil && !IsNil(o.Department) {
		return true
	}

	return false
}

// SetDepartment gets a reference to the given Departments and assigns it to the Department field.
func (o *DepartmentsPostRequest) SetDepartment(v Departments) {
	o.Department = &v
}

func (o DepartmentsPostRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o DepartmentsPostRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Department) {
		toSerialize["department"] = o.Department
	}
	return toSerialize, nil
}

type NullableDepartmentsPostRequest struct {
	value *DepartmentsPostRequest
	isSet bool
}

func (v NullableDepartmentsPostRequest) Get() *DepartmentsPostRequest {
	return v.value
}

func (v *NullableDepartmentsPostRequest) Set(val *DepartmentsPostRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableDepartmentsPostRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableDepartmentsPostRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDepartmentsPostRequest(val *DepartmentsPostRequest) *NullableDepartmentsPostRequest {
	return &NullableDepartmentsPostRequest{value: val, isSet: true}
}

func (v NullableDepartmentsPostRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDepartmentsPostRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


