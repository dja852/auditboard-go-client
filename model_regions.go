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

// checks if the Regions type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &Regions{}

// Regions struct for Regions
type Regions struct {
	// Note: This is a Primary Key.<pk/>
	Id *int32 `json:"id,omitempty"`
	CreatedAt *string `json:"created_at,omitempty"`
	UpdatedAt *string `json:"updated_at,omitempty"`
	DeletedAt *string `json:"deleted_at,omitempty"`
	SortOrder int32 `json:"sort_order"`
	RegionCode string `json:"region_code"`
	Name string `json:"name"`
	// Note: This is a Foreign Key to `workspaces.id`.<fk table='workspaces' column='id'/>
	WorkspaceId int32 `json:"workspace_id"`
	FormTemplates interface{} `json:"form_templates,omitempty"`
	EnabledAttributes interface{} `json:"enabled_attributes,omitempty"`
	ExcludedAttributes interface{} `json:"excluded_attributes,omitempty"`
	Scopes interface{} `json:"scopes"`
	Options interface{} `json:"options,omitempty"`
}

type _Regions Regions

// NewRegions instantiates a new Regions object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRegions(sortOrder int32, regionCode string, name string, workspaceId int32, scopes interface{}) *Regions {
	this := Regions{}
	this.SortOrder = sortOrder
	this.RegionCode = regionCode
	this.Name = name
	this.WorkspaceId = workspaceId
	this.Scopes = scopes
	return &this
}

// NewRegionsWithDefaults instantiates a new Regions object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRegionsWithDefaults() *Regions {
	this := Regions{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *Regions) GetId() int32 {
	if o == nil || IsNil(o.Id) {
		var ret int32
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Regions) GetIdOk() (*int32, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *Regions) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given int32 and assigns it to the Id field.
func (o *Regions) SetId(v int32) {
	o.Id = &v
}

// GetCreatedAt returns the CreatedAt field value if set, zero value otherwise.
func (o *Regions) GetCreatedAt() string {
	if o == nil || IsNil(o.CreatedAt) {
		var ret string
		return ret
	}
	return *o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Regions) GetCreatedAtOk() (*string, bool) {
	if o == nil || IsNil(o.CreatedAt) {
		return nil, false
	}
	return o.CreatedAt, true
}

// HasCreatedAt returns a boolean if a field has been set.
func (o *Regions) HasCreatedAt() bool {
	if o != nil && !IsNil(o.CreatedAt) {
		return true
	}

	return false
}

// SetCreatedAt gets a reference to the given string and assigns it to the CreatedAt field.
func (o *Regions) SetCreatedAt(v string) {
	o.CreatedAt = &v
}

// GetUpdatedAt returns the UpdatedAt field value if set, zero value otherwise.
func (o *Regions) GetUpdatedAt() string {
	if o == nil || IsNil(o.UpdatedAt) {
		var ret string
		return ret
	}
	return *o.UpdatedAt
}

// GetUpdatedAtOk returns a tuple with the UpdatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Regions) GetUpdatedAtOk() (*string, bool) {
	if o == nil || IsNil(o.UpdatedAt) {
		return nil, false
	}
	return o.UpdatedAt, true
}

// HasUpdatedAt returns a boolean if a field has been set.
func (o *Regions) HasUpdatedAt() bool {
	if o != nil && !IsNil(o.UpdatedAt) {
		return true
	}

	return false
}

// SetUpdatedAt gets a reference to the given string and assigns it to the UpdatedAt field.
func (o *Regions) SetUpdatedAt(v string) {
	o.UpdatedAt = &v
}

// GetDeletedAt returns the DeletedAt field value if set, zero value otherwise.
func (o *Regions) GetDeletedAt() string {
	if o == nil || IsNil(o.DeletedAt) {
		var ret string
		return ret
	}
	return *o.DeletedAt
}

// GetDeletedAtOk returns a tuple with the DeletedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Regions) GetDeletedAtOk() (*string, bool) {
	if o == nil || IsNil(o.DeletedAt) {
		return nil, false
	}
	return o.DeletedAt, true
}

// HasDeletedAt returns a boolean if a field has been set.
func (o *Regions) HasDeletedAt() bool {
	if o != nil && !IsNil(o.DeletedAt) {
		return true
	}

	return false
}

// SetDeletedAt gets a reference to the given string and assigns it to the DeletedAt field.
func (o *Regions) SetDeletedAt(v string) {
	o.DeletedAt = &v
}

// GetSortOrder returns the SortOrder field value
func (o *Regions) GetSortOrder() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.SortOrder
}

// GetSortOrderOk returns a tuple with the SortOrder field value
// and a boolean to check if the value has been set.
func (o *Regions) GetSortOrderOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SortOrder, true
}

// SetSortOrder sets field value
func (o *Regions) SetSortOrder(v int32) {
	o.SortOrder = v
}

// GetRegionCode returns the RegionCode field value
func (o *Regions) GetRegionCode() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.RegionCode
}

// GetRegionCodeOk returns a tuple with the RegionCode field value
// and a boolean to check if the value has been set.
func (o *Regions) GetRegionCodeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.RegionCode, true
}

// SetRegionCode sets field value
func (o *Regions) SetRegionCode(v string) {
	o.RegionCode = v
}

// GetName returns the Name field value
func (o *Regions) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *Regions) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *Regions) SetName(v string) {
	o.Name = v
}

// GetWorkspaceId returns the WorkspaceId field value
func (o *Regions) GetWorkspaceId() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.WorkspaceId
}

// GetWorkspaceIdOk returns a tuple with the WorkspaceId field value
// and a boolean to check if the value has been set.
func (o *Regions) GetWorkspaceIdOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.WorkspaceId, true
}

// SetWorkspaceId sets field value
func (o *Regions) SetWorkspaceId(v int32) {
	o.WorkspaceId = v
}

// GetFormTemplates returns the FormTemplates field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Regions) GetFormTemplates() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.FormTemplates
}

// GetFormTemplatesOk returns a tuple with the FormTemplates field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Regions) GetFormTemplatesOk() (*interface{}, bool) {
	if o == nil || IsNil(o.FormTemplates) {
		return nil, false
	}
	return &o.FormTemplates, true
}

// HasFormTemplates returns a boolean if a field has been set.
func (o *Regions) HasFormTemplates() bool {
	if o != nil && !IsNil(o.FormTemplates) {
		return true
	}

	return false
}

// SetFormTemplates gets a reference to the given interface{} and assigns it to the FormTemplates field.
func (o *Regions) SetFormTemplates(v interface{}) {
	o.FormTemplates = v
}

// GetEnabledAttributes returns the EnabledAttributes field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Regions) GetEnabledAttributes() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.EnabledAttributes
}

// GetEnabledAttributesOk returns a tuple with the EnabledAttributes field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Regions) GetEnabledAttributesOk() (*interface{}, bool) {
	if o == nil || IsNil(o.EnabledAttributes) {
		return nil, false
	}
	return &o.EnabledAttributes, true
}

// HasEnabledAttributes returns a boolean if a field has been set.
func (o *Regions) HasEnabledAttributes() bool {
	if o != nil && !IsNil(o.EnabledAttributes) {
		return true
	}

	return false
}

// SetEnabledAttributes gets a reference to the given interface{} and assigns it to the EnabledAttributes field.
func (o *Regions) SetEnabledAttributes(v interface{}) {
	o.EnabledAttributes = v
}

// GetExcludedAttributes returns the ExcludedAttributes field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Regions) GetExcludedAttributes() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.ExcludedAttributes
}

// GetExcludedAttributesOk returns a tuple with the ExcludedAttributes field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Regions) GetExcludedAttributesOk() (*interface{}, bool) {
	if o == nil || IsNil(o.ExcludedAttributes) {
		return nil, false
	}
	return &o.ExcludedAttributes, true
}

// HasExcludedAttributes returns a boolean if a field has been set.
func (o *Regions) HasExcludedAttributes() bool {
	if o != nil && !IsNil(o.ExcludedAttributes) {
		return true
	}

	return false
}

// SetExcludedAttributes gets a reference to the given interface{} and assigns it to the ExcludedAttributes field.
func (o *Regions) SetExcludedAttributes(v interface{}) {
	o.ExcludedAttributes = v
}

// GetScopes returns the Scopes field value
// If the value is explicit nil, the zero value for interface{} will be returned
func (o *Regions) GetScopes() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}

	return o.Scopes
}

// GetScopesOk returns a tuple with the Scopes field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Regions) GetScopesOk() (*interface{}, bool) {
	if o == nil || IsNil(o.Scopes) {
		return nil, false
	}
	return &o.Scopes, true
}

// SetScopes sets field value
func (o *Regions) SetScopes(v interface{}) {
	o.Scopes = v
}

// GetOptions returns the Options field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Regions) GetOptions() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.Options
}

// GetOptionsOk returns a tuple with the Options field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Regions) GetOptionsOk() (*interface{}, bool) {
	if o == nil || IsNil(o.Options) {
		return nil, false
	}
	return &o.Options, true
}

// HasOptions returns a boolean if a field has been set.
func (o *Regions) HasOptions() bool {
	if o != nil && !IsNil(o.Options) {
		return true
	}

	return false
}

// SetOptions gets a reference to the given interface{} and assigns it to the Options field.
func (o *Regions) SetOptions(v interface{}) {
	o.Options = v
}

func (o Regions) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o Regions) ToMap() (map[string]interface{}, error) {
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
	toSerialize["sort_order"] = o.SortOrder
	toSerialize["region_code"] = o.RegionCode
	toSerialize["name"] = o.Name
	toSerialize["workspace_id"] = o.WorkspaceId
	if o.FormTemplates != nil {
		toSerialize["form_templates"] = o.FormTemplates
	}
	if o.EnabledAttributes != nil {
		toSerialize["enabled_attributes"] = o.EnabledAttributes
	}
	if o.ExcludedAttributes != nil {
		toSerialize["excluded_attributes"] = o.ExcludedAttributes
	}
	if o.Scopes != nil {
		toSerialize["scopes"] = o.Scopes
	}
	if o.Options != nil {
		toSerialize["options"] = o.Options
	}
	return toSerialize, nil
}

func (o *Regions) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"sort_order",
		"region_code",
		"name",
		"workspace_id",
		"scopes",
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

	varRegions := _Regions{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varRegions)

	if err != nil {
		return err
	}

	*o = Regions(varRegions)

	return err
}

type NullableRegions struct {
	value *Regions
	isSet bool
}

func (v NullableRegions) Get() *Regions {
	return v.value
}

func (v *NullableRegions) Set(val *Regions) {
	v.value = val
	v.isSet = true
}

func (v NullableRegions) IsSet() bool {
	return v.isSet
}

func (v *NullableRegions) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRegions(val *Regions) *NullableRegions {
	return &NullableRegions{value: val, isSet: true}
}

func (v NullableRegions) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRegions) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


