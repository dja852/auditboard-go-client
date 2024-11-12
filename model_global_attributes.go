/*
AuditBoard Developer Portal API Documentation

External API reference documentation.

API version: 23.8.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package auditboard

import (
	"encoding/json"
	"bytes"
	"fmt"
)

// checks if the GlobalAttributes type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GlobalAttributes{}

// GlobalAttributes struct for GlobalAttributes
type GlobalAttributes struct {
	// Note: This is a Primary Key.<pk/>
	Id *int32 `json:"id,omitempty"`
	Section *string `json:"section,omitempty"`
	ObjectType *string `json:"object_type,omitempty"`
	Key *string `json:"key,omitempty"`
	DefaultLabel *string `json:"default_label,omitempty"`
	Label *string `json:"label,omitempty"`
	CreatedAt *string `json:"created_at,omitempty"`
	UpdatedAt *string `json:"updated_at,omitempty"`
	DeletedAt *string `json:"deleted_at,omitempty"`
	Enabled bool `json:"enabled"`
	Scope interface{} `json:"scope"`
}

type _GlobalAttributes GlobalAttributes

// NewGlobalAttributes instantiates a new GlobalAttributes object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGlobalAttributes(enabled bool, scope interface{}) *GlobalAttributes {
	this := GlobalAttributes{}
	this.Enabled = enabled
	this.Scope = scope
	return &this
}

// NewGlobalAttributesWithDefaults instantiates a new GlobalAttributes object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGlobalAttributesWithDefaults() *GlobalAttributes {
	this := GlobalAttributes{}
	var enabled bool = true
	this.Enabled = enabled
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *GlobalAttributes) GetId() int32 {
	if o == nil || IsNil(o.Id) {
		var ret int32
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GlobalAttributes) GetIdOk() (*int32, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *GlobalAttributes) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given int32 and assigns it to the Id field.
func (o *GlobalAttributes) SetId(v int32) {
	o.Id = &v
}

// GetSection returns the Section field value if set, zero value otherwise.
func (o *GlobalAttributes) GetSection() string {
	if o == nil || IsNil(o.Section) {
		var ret string
		return ret
	}
	return *o.Section
}

// GetSectionOk returns a tuple with the Section field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GlobalAttributes) GetSectionOk() (*string, bool) {
	if o == nil || IsNil(o.Section) {
		return nil, false
	}
	return o.Section, true
}

// HasSection returns a boolean if a field has been set.
func (o *GlobalAttributes) HasSection() bool {
	if o != nil && !IsNil(o.Section) {
		return true
	}

	return false
}

// SetSection gets a reference to the given string and assigns it to the Section field.
func (o *GlobalAttributes) SetSection(v string) {
	o.Section = &v
}

// GetObjectType returns the ObjectType field value if set, zero value otherwise.
func (o *GlobalAttributes) GetObjectType() string {
	if o == nil || IsNil(o.ObjectType) {
		var ret string
		return ret
	}
	return *o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GlobalAttributes) GetObjectTypeOk() (*string, bool) {
	if o == nil || IsNil(o.ObjectType) {
		return nil, false
	}
	return o.ObjectType, true
}

// HasObjectType returns a boolean if a field has been set.
func (o *GlobalAttributes) HasObjectType() bool {
	if o != nil && !IsNil(o.ObjectType) {
		return true
	}

	return false
}

// SetObjectType gets a reference to the given string and assigns it to the ObjectType field.
func (o *GlobalAttributes) SetObjectType(v string) {
	o.ObjectType = &v
}

// GetKey returns the Key field value if set, zero value otherwise.
func (o *GlobalAttributes) GetKey() string {
	if o == nil || IsNil(o.Key) {
		var ret string
		return ret
	}
	return *o.Key
}

// GetKeyOk returns a tuple with the Key field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GlobalAttributes) GetKeyOk() (*string, bool) {
	if o == nil || IsNil(o.Key) {
		return nil, false
	}
	return o.Key, true
}

// HasKey returns a boolean if a field has been set.
func (o *GlobalAttributes) HasKey() bool {
	if o != nil && !IsNil(o.Key) {
		return true
	}

	return false
}

// SetKey gets a reference to the given string and assigns it to the Key field.
func (o *GlobalAttributes) SetKey(v string) {
	o.Key = &v
}

// GetDefaultLabel returns the DefaultLabel field value if set, zero value otherwise.
func (o *GlobalAttributes) GetDefaultLabel() string {
	if o == nil || IsNil(o.DefaultLabel) {
		var ret string
		return ret
	}
	return *o.DefaultLabel
}

// GetDefaultLabelOk returns a tuple with the DefaultLabel field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GlobalAttributes) GetDefaultLabelOk() (*string, bool) {
	if o == nil || IsNil(o.DefaultLabel) {
		return nil, false
	}
	return o.DefaultLabel, true
}

// HasDefaultLabel returns a boolean if a field has been set.
func (o *GlobalAttributes) HasDefaultLabel() bool {
	if o != nil && !IsNil(o.DefaultLabel) {
		return true
	}

	return false
}

// SetDefaultLabel gets a reference to the given string and assigns it to the DefaultLabel field.
func (o *GlobalAttributes) SetDefaultLabel(v string) {
	o.DefaultLabel = &v
}

// GetLabel returns the Label field value if set, zero value otherwise.
func (o *GlobalAttributes) GetLabel() string {
	if o == nil || IsNil(o.Label) {
		var ret string
		return ret
	}
	return *o.Label
}

// GetLabelOk returns a tuple with the Label field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GlobalAttributes) GetLabelOk() (*string, bool) {
	if o == nil || IsNil(o.Label) {
		return nil, false
	}
	return o.Label, true
}

// HasLabel returns a boolean if a field has been set.
func (o *GlobalAttributes) HasLabel() bool {
	if o != nil && !IsNil(o.Label) {
		return true
	}

	return false
}

// SetLabel gets a reference to the given string and assigns it to the Label field.
func (o *GlobalAttributes) SetLabel(v string) {
	o.Label = &v
}

// GetCreatedAt returns the CreatedAt field value if set, zero value otherwise.
func (o *GlobalAttributes) GetCreatedAt() string {
	if o == nil || IsNil(o.CreatedAt) {
		var ret string
		return ret
	}
	return *o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GlobalAttributes) GetCreatedAtOk() (*string, bool) {
	if o == nil || IsNil(o.CreatedAt) {
		return nil, false
	}
	return o.CreatedAt, true
}

// HasCreatedAt returns a boolean if a field has been set.
func (o *GlobalAttributes) HasCreatedAt() bool {
	if o != nil && !IsNil(o.CreatedAt) {
		return true
	}

	return false
}

// SetCreatedAt gets a reference to the given string and assigns it to the CreatedAt field.
func (o *GlobalAttributes) SetCreatedAt(v string) {
	o.CreatedAt = &v
}

// GetUpdatedAt returns the UpdatedAt field value if set, zero value otherwise.
func (o *GlobalAttributes) GetUpdatedAt() string {
	if o == nil || IsNil(o.UpdatedAt) {
		var ret string
		return ret
	}
	return *o.UpdatedAt
}

// GetUpdatedAtOk returns a tuple with the UpdatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GlobalAttributes) GetUpdatedAtOk() (*string, bool) {
	if o == nil || IsNil(o.UpdatedAt) {
		return nil, false
	}
	return o.UpdatedAt, true
}

// HasUpdatedAt returns a boolean if a field has been set.
func (o *GlobalAttributes) HasUpdatedAt() bool {
	if o != nil && !IsNil(o.UpdatedAt) {
		return true
	}

	return false
}

// SetUpdatedAt gets a reference to the given string and assigns it to the UpdatedAt field.
func (o *GlobalAttributes) SetUpdatedAt(v string) {
	o.UpdatedAt = &v
}

// GetDeletedAt returns the DeletedAt field value if set, zero value otherwise.
func (o *GlobalAttributes) GetDeletedAt() string {
	if o == nil || IsNil(o.DeletedAt) {
		var ret string
		return ret
	}
	return *o.DeletedAt
}

// GetDeletedAtOk returns a tuple with the DeletedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GlobalAttributes) GetDeletedAtOk() (*string, bool) {
	if o == nil || IsNil(o.DeletedAt) {
		return nil, false
	}
	return o.DeletedAt, true
}

// HasDeletedAt returns a boolean if a field has been set.
func (o *GlobalAttributes) HasDeletedAt() bool {
	if o != nil && !IsNil(o.DeletedAt) {
		return true
	}

	return false
}

// SetDeletedAt gets a reference to the given string and assigns it to the DeletedAt field.
func (o *GlobalAttributes) SetDeletedAt(v string) {
	o.DeletedAt = &v
}

// GetEnabled returns the Enabled field value
func (o *GlobalAttributes) GetEnabled() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.Enabled
}

// GetEnabledOk returns a tuple with the Enabled field value
// and a boolean to check if the value has been set.
func (o *GlobalAttributes) GetEnabledOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Enabled, true
}

// SetEnabled sets field value
func (o *GlobalAttributes) SetEnabled(v bool) {
	o.Enabled = v
}

// GetScope returns the Scope field value
// If the value is explicit nil, the zero value for interface{} will be returned
func (o *GlobalAttributes) GetScope() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}

	return o.Scope
}

// GetScopeOk returns a tuple with the Scope field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *GlobalAttributes) GetScopeOk() (*interface{}, bool) {
	if o == nil || IsNil(o.Scope) {
		return nil, false
	}
	return &o.Scope, true
}

// SetScope sets field value
func (o *GlobalAttributes) SetScope(v interface{}) {
	o.Scope = v
}

func (o GlobalAttributes) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GlobalAttributes) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !IsNil(o.Section) {
		toSerialize["section"] = o.Section
	}
	if !IsNil(o.ObjectType) {
		toSerialize["object_type"] = o.ObjectType
	}
	if !IsNil(o.Key) {
		toSerialize["key"] = o.Key
	}
	if !IsNil(o.DefaultLabel) {
		toSerialize["default_label"] = o.DefaultLabel
	}
	if !IsNil(o.Label) {
		toSerialize["label"] = o.Label
	}
	if !IsNil(o.CreatedAt) {
		toSerialize["created_at"] = o.CreatedAt
	}
	if !IsNil(o.UpdatedAt) {
		toSerialize["updated_at"] = o.UpdatedAt
	}
	if !IsNil(o.DeletedAt) {
		toSerialize["deleted_at"] = o.DeletedAt
	}
	toSerialize["enabled"] = o.Enabled
	if o.Scope != nil {
		toSerialize["scope"] = o.Scope
	}
	return toSerialize, nil
}

func (o *GlobalAttributes) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"enabled",
		"scope",
	}

	allProperties := make(map[string]interface{})

	err = json.Unmarshal(data, &allProperties)

	if err != nil {
		return err;
	}

	for _, requiredProperty := range(requiredProperties) {
		if _, exists := allProperties[requiredProperty]; !exists {
			return fmt.Errorf("no value given for required property %v", requiredProperty)
		}
	}

	varGlobalAttributes := _GlobalAttributes{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varGlobalAttributes)

	if err != nil {
		return err
	}

	*o = GlobalAttributes(varGlobalAttributes)

	return err
}

type NullableGlobalAttributes struct {
	value *GlobalAttributes
	isSet bool
}

func (v NullableGlobalAttributes) Get() *GlobalAttributes {
	return v.value
}

func (v *NullableGlobalAttributes) Set(val *GlobalAttributes) {
	v.value = val
	v.isSet = true
}

func (v NullableGlobalAttributes) IsSet() bool {
	return v.isSet
}

func (v *NullableGlobalAttributes) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGlobalAttributes(val *GlobalAttributes) *NullableGlobalAttributes {
	return &NullableGlobalAttributes{value: val, isSet: true}
}

func (v NullableGlobalAttributes) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGlobalAttributes) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


