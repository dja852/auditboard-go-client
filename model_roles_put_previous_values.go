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

// checks if the RolesPutPreviousValues type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &RolesPutPreviousValues{}

// RolesPutPreviousValues struct for RolesPutPreviousValues
type RolesPutPreviousValues struct {
	// Note: This is a Primary Key.<pk/>
	Id *int32 `json:"id,omitempty"`
	CreatedAt *string `json:"created_at,omitempty"`
	UpdatedAt *string `json:"updated_at,omitempty"`
	DeletedAt *string `json:"deleted_at,omitempty"`
	Name *string `json:"name,omitempty"`
	ShortName *string `json:"short_name,omitempty"`
	DefaultRoute *string `json:"default_route,omitempty"`
	Description *string `json:"description,omitempty"`
	IsImplementer *bool `json:"is_implementer,omitempty"`
	IsLimited *bool `json:"is_limited,omitempty"`
}

// NewRolesPutPreviousValues instantiates a new RolesPutPreviousValues object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRolesPutPreviousValues() *RolesPutPreviousValues {
	this := RolesPutPreviousValues{}
	var isImplementer bool = false
	this.IsImplementer = &isImplementer
	var isLimited bool = false
	this.IsLimited = &isLimited
	return &this
}

// NewRolesPutPreviousValuesWithDefaults instantiates a new RolesPutPreviousValues object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRolesPutPreviousValuesWithDefaults() *RolesPutPreviousValues {
	this := RolesPutPreviousValues{}
	var isImplementer bool = false
	this.IsImplementer = &isImplementer
	var isLimited bool = false
	this.IsLimited = &isLimited
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *RolesPutPreviousValues) GetId() int32 {
	if o == nil || IsNil(o.Id) {
		var ret int32
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RolesPutPreviousValues) GetIdOk() (*int32, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *RolesPutPreviousValues) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given int32 and assigns it to the Id field.
func (o *RolesPutPreviousValues) SetId(v int32) {
	o.Id = &v
}

// GetCreatedAt returns the CreatedAt field value if set, zero value otherwise.
func (o *RolesPutPreviousValues) GetCreatedAt() string {
	if o == nil || IsNil(o.CreatedAt) {
		var ret string
		return ret
	}
	return *o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RolesPutPreviousValues) GetCreatedAtOk() (*string, bool) {
	if o == nil || IsNil(o.CreatedAt) {
		return nil, false
	}
	return o.CreatedAt, true
}

// HasCreatedAt returns a boolean if a field has been set.
func (o *RolesPutPreviousValues) HasCreatedAt() bool {
	if o != nil && !IsNil(o.CreatedAt) {
		return true
	}

	return false
}

// SetCreatedAt gets a reference to the given string and assigns it to the CreatedAt field.
func (o *RolesPutPreviousValues) SetCreatedAt(v string) {
	o.CreatedAt = &v
}

// GetUpdatedAt returns the UpdatedAt field value if set, zero value otherwise.
func (o *RolesPutPreviousValues) GetUpdatedAt() string {
	if o == nil || IsNil(o.UpdatedAt) {
		var ret string
		return ret
	}
	return *o.UpdatedAt
}

// GetUpdatedAtOk returns a tuple with the UpdatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RolesPutPreviousValues) GetUpdatedAtOk() (*string, bool) {
	if o == nil || IsNil(o.UpdatedAt) {
		return nil, false
	}
	return o.UpdatedAt, true
}

// HasUpdatedAt returns a boolean if a field has been set.
func (o *RolesPutPreviousValues) HasUpdatedAt() bool {
	if o != nil && !IsNil(o.UpdatedAt) {
		return true
	}

	return false
}

// SetUpdatedAt gets a reference to the given string and assigns it to the UpdatedAt field.
func (o *RolesPutPreviousValues) SetUpdatedAt(v string) {
	o.UpdatedAt = &v
}

// GetDeletedAt returns the DeletedAt field value if set, zero value otherwise.
func (o *RolesPutPreviousValues) GetDeletedAt() string {
	if o == nil || IsNil(o.DeletedAt) {
		var ret string
		return ret
	}
	return *o.DeletedAt
}

// GetDeletedAtOk returns a tuple with the DeletedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RolesPutPreviousValues) GetDeletedAtOk() (*string, bool) {
	if o == nil || IsNil(o.DeletedAt) {
		return nil, false
	}
	return o.DeletedAt, true
}

// HasDeletedAt returns a boolean if a field has been set.
func (o *RolesPutPreviousValues) HasDeletedAt() bool {
	if o != nil && !IsNil(o.DeletedAt) {
		return true
	}

	return false
}

// SetDeletedAt gets a reference to the given string and assigns it to the DeletedAt field.
func (o *RolesPutPreviousValues) SetDeletedAt(v string) {
	o.DeletedAt = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *RolesPutPreviousValues) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RolesPutPreviousValues) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *RolesPutPreviousValues) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *RolesPutPreviousValues) SetName(v string) {
	o.Name = &v
}

// GetShortName returns the ShortName field value if set, zero value otherwise.
func (o *RolesPutPreviousValues) GetShortName() string {
	if o == nil || IsNil(o.ShortName) {
		var ret string
		return ret
	}
	return *o.ShortName
}

// GetShortNameOk returns a tuple with the ShortName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RolesPutPreviousValues) GetShortNameOk() (*string, bool) {
	if o == nil || IsNil(o.ShortName) {
		return nil, false
	}
	return o.ShortName, true
}

// HasShortName returns a boolean if a field has been set.
func (o *RolesPutPreviousValues) HasShortName() bool {
	if o != nil && !IsNil(o.ShortName) {
		return true
	}

	return false
}

// SetShortName gets a reference to the given string and assigns it to the ShortName field.
func (o *RolesPutPreviousValues) SetShortName(v string) {
	o.ShortName = &v
}

// GetDefaultRoute returns the DefaultRoute field value if set, zero value otherwise.
func (o *RolesPutPreviousValues) GetDefaultRoute() string {
	if o == nil || IsNil(o.DefaultRoute) {
		var ret string
		return ret
	}
	return *o.DefaultRoute
}

// GetDefaultRouteOk returns a tuple with the DefaultRoute field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RolesPutPreviousValues) GetDefaultRouteOk() (*string, bool) {
	if o == nil || IsNil(o.DefaultRoute) {
		return nil, false
	}
	return o.DefaultRoute, true
}

// HasDefaultRoute returns a boolean if a field has been set.
func (o *RolesPutPreviousValues) HasDefaultRoute() bool {
	if o != nil && !IsNil(o.DefaultRoute) {
		return true
	}

	return false
}

// SetDefaultRoute gets a reference to the given string and assigns it to the DefaultRoute field.
func (o *RolesPutPreviousValues) SetDefaultRoute(v string) {
	o.DefaultRoute = &v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *RolesPutPreviousValues) GetDescription() string {
	if o == nil || IsNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RolesPutPreviousValues) GetDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *RolesPutPreviousValues) HasDescription() bool {
	if o != nil && !IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *RolesPutPreviousValues) SetDescription(v string) {
	o.Description = &v
}

// GetIsImplementer returns the IsImplementer field value if set, zero value otherwise.
func (o *RolesPutPreviousValues) GetIsImplementer() bool {
	if o == nil || IsNil(o.IsImplementer) {
		var ret bool
		return ret
	}
	return *o.IsImplementer
}

// GetIsImplementerOk returns a tuple with the IsImplementer field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RolesPutPreviousValues) GetIsImplementerOk() (*bool, bool) {
	if o == nil || IsNil(o.IsImplementer) {
		return nil, false
	}
	return o.IsImplementer, true
}

// HasIsImplementer returns a boolean if a field has been set.
func (o *RolesPutPreviousValues) HasIsImplementer() bool {
	if o != nil && !IsNil(o.IsImplementer) {
		return true
	}

	return false
}

// SetIsImplementer gets a reference to the given bool and assigns it to the IsImplementer field.
func (o *RolesPutPreviousValues) SetIsImplementer(v bool) {
	o.IsImplementer = &v
}

// GetIsLimited returns the IsLimited field value if set, zero value otherwise.
func (o *RolesPutPreviousValues) GetIsLimited() bool {
	if o == nil || IsNil(o.IsLimited) {
		var ret bool
		return ret
	}
	return *o.IsLimited
}

// GetIsLimitedOk returns a tuple with the IsLimited field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RolesPutPreviousValues) GetIsLimitedOk() (*bool, bool) {
	if o == nil || IsNil(o.IsLimited) {
		return nil, false
	}
	return o.IsLimited, true
}

// HasIsLimited returns a boolean if a field has been set.
func (o *RolesPutPreviousValues) HasIsLimited() bool {
	if o != nil && !IsNil(o.IsLimited) {
		return true
	}

	return false
}

// SetIsLimited gets a reference to the given bool and assigns it to the IsLimited field.
func (o *RolesPutPreviousValues) SetIsLimited(v bool) {
	o.IsLimited = &v
}

func (o RolesPutPreviousValues) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o RolesPutPreviousValues) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
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
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !IsNil(o.ShortName) {
		toSerialize["short_name"] = o.ShortName
	}
	if !IsNil(o.DefaultRoute) {
		toSerialize["default_route"] = o.DefaultRoute
	}
	if !IsNil(o.Description) {
		toSerialize["description"] = o.Description
	}
	if !IsNil(o.IsImplementer) {
		toSerialize["is_implementer"] = o.IsImplementer
	}
	if !IsNil(o.IsLimited) {
		toSerialize["is_limited"] = o.IsLimited
	}
	return toSerialize, nil
}

type NullableRolesPutPreviousValues struct {
	value *RolesPutPreviousValues
	isSet bool
}

func (v NullableRolesPutPreviousValues) Get() *RolesPutPreviousValues {
	return v.value
}

func (v *NullableRolesPutPreviousValues) Set(val *RolesPutPreviousValues) {
	v.value = val
	v.isSet = true
}

func (v NullableRolesPutPreviousValues) IsSet() bool {
	return v.isSet
}

func (v *NullableRolesPutPreviousValues) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRolesPutPreviousValues(val *RolesPutPreviousValues) *NullableRolesPutPreviousValues {
	return &NullableRolesPutPreviousValues{value: val, isSet: true}
}

func (v NullableRolesPutPreviousValues) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRolesPutPreviousValues) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


