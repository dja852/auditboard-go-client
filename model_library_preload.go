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

// checks if the LibraryPreload type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &LibraryPreload{}

// LibraryPreload struct for LibraryPreload
type LibraryPreload struct {
	Data []interface{} `json:"data,omitempty"`
	Included []LibraryPreloadIncludedInner `json:"included,omitempty"`
	Meta []LibraryPreloadMetaInner `json:"meta,omitempty"`
}

// NewLibraryPreload instantiates a new LibraryPreload object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewLibraryPreload() *LibraryPreload {
	this := LibraryPreload{}
	return &this
}

// NewLibraryPreloadWithDefaults instantiates a new LibraryPreload object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewLibraryPreloadWithDefaults() *LibraryPreload {
	this := LibraryPreload{}
	return &this
}

// GetData returns the Data field value if set, zero value otherwise.
func (o *LibraryPreload) GetData() []interface{} {
	if o == nil || IsNil(o.Data) {
		var ret []interface{}
		return ret
	}
	return o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LibraryPreload) GetDataOk() ([]interface{}, bool) {
	if o == nil || IsNil(o.Data) {
		return nil, false
	}
	return o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *LibraryPreload) HasData() bool {
	if o != nil && !IsNil(o.Data) {
		return true
	}

	return false
}

// SetData gets a reference to the given []interface{} and assigns it to the Data field.
func (o *LibraryPreload) SetData(v []interface{}) {
	o.Data = v
}

// GetIncluded returns the Included field value if set, zero value otherwise.
func (o *LibraryPreload) GetIncluded() []LibraryPreloadIncludedInner {
	if o == nil || IsNil(o.Included) {
		var ret []LibraryPreloadIncludedInner
		return ret
	}
	return o.Included
}

// GetIncludedOk returns a tuple with the Included field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LibraryPreload) GetIncludedOk() ([]LibraryPreloadIncludedInner, bool) {
	if o == nil || IsNil(o.Included) {
		return nil, false
	}
	return o.Included, true
}

// HasIncluded returns a boolean if a field has been set.
func (o *LibraryPreload) HasIncluded() bool {
	if o != nil && !IsNil(o.Included) {
		return true
	}

	return false
}

// SetIncluded gets a reference to the given []LibraryPreloadIncludedInner and assigns it to the Included field.
func (o *LibraryPreload) SetIncluded(v []LibraryPreloadIncludedInner) {
	o.Included = v
}

// GetMeta returns the Meta field value if set, zero value otherwise.
func (o *LibraryPreload) GetMeta() []LibraryPreloadMetaInner {
	if o == nil || IsNil(o.Meta) {
		var ret []LibraryPreloadMetaInner
		return ret
	}
	return o.Meta
}

// GetMetaOk returns a tuple with the Meta field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LibraryPreload) GetMetaOk() ([]LibraryPreloadMetaInner, bool) {
	if o == nil || IsNil(o.Meta) {
		return nil, false
	}
	return o.Meta, true
}

// HasMeta returns a boolean if a field has been set.
func (o *LibraryPreload) HasMeta() bool {
	if o != nil && !IsNil(o.Meta) {
		return true
	}

	return false
}

// SetMeta gets a reference to the given []LibraryPreloadMetaInner and assigns it to the Meta field.
func (o *LibraryPreload) SetMeta(v []LibraryPreloadMetaInner) {
	o.Meta = v
}

func (o LibraryPreload) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o LibraryPreload) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Data) {
		toSerialize["data"] = o.Data
	}
	if !IsNil(o.Included) {
		toSerialize["included"] = o.Included
	}
	if !IsNil(o.Meta) {
		toSerialize["meta"] = o.Meta
	}
	return toSerialize, nil
}

type NullableLibraryPreload struct {
	value *LibraryPreload
	isSet bool
}

func (v NullableLibraryPreload) Get() *LibraryPreload {
	return v.value
}

func (v *NullableLibraryPreload) Set(val *LibraryPreload) {
	v.value = val
	v.isSet = true
}

func (v NullableLibraryPreload) IsSet() bool {
	return v.isSet
}

func (v *NullableLibraryPreload) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableLibraryPreload(val *LibraryPreload) *NullableLibraryPreload {
	return &NullableLibraryPreload{value: val, isSet: true}
}

func (v NullableLibraryPreload) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableLibraryPreload) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


