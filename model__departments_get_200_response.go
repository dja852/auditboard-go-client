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

// checks if the DepartmentsGet200Response type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &DepartmentsGet200Response{}

// DepartmentsGet200Response struct for DepartmentsGet200Response
type DepartmentsGet200Response struct {
	Departments []Departments `json:"departments,omitempty"`
}

// NewDepartmentsGet200Response instantiates a new DepartmentsGet200Response object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDepartmentsGet200Response() *DepartmentsGet200Response {
	this := DepartmentsGet200Response{}
	return &this
}

// NewDepartmentsGet200ResponseWithDefaults instantiates a new DepartmentsGet200Response object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDepartmentsGet200ResponseWithDefaults() *DepartmentsGet200Response {
	this := DepartmentsGet200Response{}
	return &this
}

// GetDepartments returns the Departments field value if set, zero value otherwise.
func (o *DepartmentsGet200Response) GetDepartments() []Departments {
	if o == nil || IsNil(o.Departments) {
		var ret []Departments
		return ret
	}
	return o.Departments
}

// GetDepartmentsOk returns a tuple with the Departments field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DepartmentsGet200Response) GetDepartmentsOk() ([]Departments, bool) {
	if o == nil || IsNil(o.Departments) {
		return nil, false
	}
	return o.Departments, true
}

// HasDepartments returns a boolean if a field has been set.
func (o *DepartmentsGet200Response) HasDepartments() bool {
	if o != nil && !IsNil(o.Departments) {
		return true
	}

	return false
}

// SetDepartments gets a reference to the given []Departments and assigns it to the Departments field.
func (o *DepartmentsGet200Response) SetDepartments(v []Departments) {
	o.Departments = v
}

func (o DepartmentsGet200Response) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o DepartmentsGet200Response) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Departments) {
		toSerialize["departments"] = o.Departments
	}
	return toSerialize, nil
}

type NullableDepartmentsGet200Response struct {
	value *DepartmentsGet200Response
	isSet bool
}

func (v NullableDepartmentsGet200Response) Get() *DepartmentsGet200Response {
	return v.value
}

func (v *NullableDepartmentsGet200Response) Set(val *DepartmentsGet200Response) {
	v.value = val
	v.isSet = true
}

func (v NullableDepartmentsGet200Response) IsSet() bool {
	return v.isSet
}

func (v *NullableDepartmentsGet200Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDepartmentsGet200Response(val *DepartmentsGet200Response) *NullableDepartmentsGet200Response {
	return &NullableDepartmentsGet200Response{value: val, isSet: true}
}

func (v NullableDepartmentsGet200Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDepartmentsGet200Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


