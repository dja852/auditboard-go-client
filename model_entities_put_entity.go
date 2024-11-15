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

// checks if the EntitiesPutEntity type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &EntitiesPutEntity{}

// EntitiesPutEntity struct for EntitiesPutEntity
type EntitiesPutEntity struct {
	// Note: This is a Primary Key.<pk/>
	Id *int32 `json:"id,omitempty"`
	CreatedAt *string `json:"created_at,omitempty"`
	UpdatedAt *string `json:"updated_at,omitempty"`
	DeletedAt *string `json:"deleted_at,omitempty"`
	SortOrder int32 `json:"sort_order"`
	EntityCode string `json:"entity_code"`
	Name string `json:"name"`
	// Note: This is a Foreign Key to `regions.id`.<fk table='regions' column='id'/>
	RegionId int32 `json:"region_id"`
	// Note: This is a Foreign Key to `auditable_entities.id`.<fk table='auditable_entities' column='id'/>
	AuditableEntityId *int32 `json:"auditable_entity_id,omitempty"`
	IsGroup bool `json:"is_group"`
}

type _EntitiesPutEntity EntitiesPutEntity

// NewEntitiesPutEntity instantiates a new EntitiesPutEntity object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewEntitiesPutEntity(sortOrder int32, entityCode string, name string, regionId int32, isGroup bool) *EntitiesPutEntity {
	this := EntitiesPutEntity{}
	this.SortOrder = sortOrder
	this.EntityCode = entityCode
	this.Name = name
	this.RegionId = regionId
	this.IsGroup = isGroup
	return &this
}

// NewEntitiesPutEntityWithDefaults instantiates a new EntitiesPutEntity object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewEntitiesPutEntityWithDefaults() *EntitiesPutEntity {
	this := EntitiesPutEntity{}
	var isGroup bool = false
	this.IsGroup = isGroup
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *EntitiesPutEntity) GetId() int32 {
	if o == nil || IsNil(o.Id) {
		var ret int32
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EntitiesPutEntity) GetIdOk() (*int32, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *EntitiesPutEntity) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given int32 and assigns it to the Id field.
func (o *EntitiesPutEntity) SetId(v int32) {
	o.Id = &v
}

// GetCreatedAt returns the CreatedAt field value if set, zero value otherwise.
func (o *EntitiesPutEntity) GetCreatedAt() string {
	if o == nil || IsNil(o.CreatedAt) {
		var ret string
		return ret
	}
	return *o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EntitiesPutEntity) GetCreatedAtOk() (*string, bool) {
	if o == nil || IsNil(o.CreatedAt) {
		return nil, false
	}
	return o.CreatedAt, true
}

// HasCreatedAt returns a boolean if a field has been set.
func (o *EntitiesPutEntity) HasCreatedAt() bool {
	if o != nil && !IsNil(o.CreatedAt) {
		return true
	}

	return false
}

// SetCreatedAt gets a reference to the given string and assigns it to the CreatedAt field.
func (o *EntitiesPutEntity) SetCreatedAt(v string) {
	o.CreatedAt = &v
}

// GetUpdatedAt returns the UpdatedAt field value if set, zero value otherwise.
func (o *EntitiesPutEntity) GetUpdatedAt() string {
	if o == nil || IsNil(o.UpdatedAt) {
		var ret string
		return ret
	}
	return *o.UpdatedAt
}

// GetUpdatedAtOk returns a tuple with the UpdatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EntitiesPutEntity) GetUpdatedAtOk() (*string, bool) {
	if o == nil || IsNil(o.UpdatedAt) {
		return nil, false
	}
	return o.UpdatedAt, true
}

// HasUpdatedAt returns a boolean if a field has been set.
func (o *EntitiesPutEntity) HasUpdatedAt() bool {
	if o != nil && !IsNil(o.UpdatedAt) {
		return true
	}

	return false
}

// SetUpdatedAt gets a reference to the given string and assigns it to the UpdatedAt field.
func (o *EntitiesPutEntity) SetUpdatedAt(v string) {
	o.UpdatedAt = &v
}

// GetDeletedAt returns the DeletedAt field value if set, zero value otherwise.
func (o *EntitiesPutEntity) GetDeletedAt() string {
	if o == nil || IsNil(o.DeletedAt) {
		var ret string
		return ret
	}
	return *o.DeletedAt
}

// GetDeletedAtOk returns a tuple with the DeletedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EntitiesPutEntity) GetDeletedAtOk() (*string, bool) {
	if o == nil || IsNil(o.DeletedAt) {
		return nil, false
	}
	return o.DeletedAt, true
}

// HasDeletedAt returns a boolean if a field has been set.
func (o *EntitiesPutEntity) HasDeletedAt() bool {
	if o != nil && !IsNil(o.DeletedAt) {
		return true
	}

	return false
}

// SetDeletedAt gets a reference to the given string and assigns it to the DeletedAt field.
func (o *EntitiesPutEntity) SetDeletedAt(v string) {
	o.DeletedAt = &v
}

// GetSortOrder returns the SortOrder field value
func (o *EntitiesPutEntity) GetSortOrder() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.SortOrder
}

// GetSortOrderOk returns a tuple with the SortOrder field value
// and a boolean to check if the value has been set.
func (o *EntitiesPutEntity) GetSortOrderOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SortOrder, true
}

// SetSortOrder sets field value
func (o *EntitiesPutEntity) SetSortOrder(v int32) {
	o.SortOrder = v
}

// GetEntityCode returns the EntityCode field value
func (o *EntitiesPutEntity) GetEntityCode() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.EntityCode
}

// GetEntityCodeOk returns a tuple with the EntityCode field value
// and a boolean to check if the value has been set.
func (o *EntitiesPutEntity) GetEntityCodeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.EntityCode, true
}

// SetEntityCode sets field value
func (o *EntitiesPutEntity) SetEntityCode(v string) {
	o.EntityCode = v
}

// GetName returns the Name field value
func (o *EntitiesPutEntity) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *EntitiesPutEntity) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *EntitiesPutEntity) SetName(v string) {
	o.Name = v
}

// GetRegionId returns the RegionId field value
func (o *EntitiesPutEntity) GetRegionId() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.RegionId
}

// GetRegionIdOk returns a tuple with the RegionId field value
// and a boolean to check if the value has been set.
func (o *EntitiesPutEntity) GetRegionIdOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.RegionId, true
}

// SetRegionId sets field value
func (o *EntitiesPutEntity) SetRegionId(v int32) {
	o.RegionId = v
}

// GetAuditableEntityId returns the AuditableEntityId field value if set, zero value otherwise.
func (o *EntitiesPutEntity) GetAuditableEntityId() int32 {
	if o == nil || IsNil(o.AuditableEntityId) {
		var ret int32
		return ret
	}
	return *o.AuditableEntityId
}

// GetAuditableEntityIdOk returns a tuple with the AuditableEntityId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EntitiesPutEntity) GetAuditableEntityIdOk() (*int32, bool) {
	if o == nil || IsNil(o.AuditableEntityId) {
		return nil, false
	}
	return o.AuditableEntityId, true
}

// HasAuditableEntityId returns a boolean if a field has been set.
func (o *EntitiesPutEntity) HasAuditableEntityId() bool {
	if o != nil && !IsNil(o.AuditableEntityId) {
		return true
	}

	return false
}

// SetAuditableEntityId gets a reference to the given int32 and assigns it to the AuditableEntityId field.
func (o *EntitiesPutEntity) SetAuditableEntityId(v int32) {
	o.AuditableEntityId = &v
}

// GetIsGroup returns the IsGroup field value
func (o *EntitiesPutEntity) GetIsGroup() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.IsGroup
}

// GetIsGroupOk returns a tuple with the IsGroup field value
// and a boolean to check if the value has been set.
func (o *EntitiesPutEntity) GetIsGroupOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.IsGroup, true
}

// SetIsGroup sets field value
func (o *EntitiesPutEntity) SetIsGroup(v bool) {
	o.IsGroup = v
}

func (o EntitiesPutEntity) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o EntitiesPutEntity) ToMap() (map[string]interface{}, error) {
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
	toSerialize["entity_code"] = o.EntityCode
	toSerialize["name"] = o.Name
	toSerialize["region_id"] = o.RegionId
	if !IsNil(o.AuditableEntityId) {
		toSerialize["auditable_entity_id"] = o.AuditableEntityId
	}
	toSerialize["is_group"] = o.IsGroup
	return toSerialize, nil
}

type NullableEntitiesPutEntity struct {
	value *EntitiesPutEntity
	isSet bool
}

func (v NullableEntitiesPutEntity) Get() *EntitiesPutEntity {
	return v.value
}

func (v *NullableEntitiesPutEntity) Set(val *EntitiesPutEntity) {
	v.value = val
	v.isSet = true
}

func (v NullableEntitiesPutEntity) IsSet() bool {
	return v.isSet
}

func (v *NullableEntitiesPutEntity) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEntitiesPutEntity(val *EntitiesPutEntity) *NullableEntitiesPutEntity {
	return &NullableEntitiesPutEntity{value: val, isSet: true}
}

func (v NullableEntitiesPutEntity) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEntitiesPutEntity) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


