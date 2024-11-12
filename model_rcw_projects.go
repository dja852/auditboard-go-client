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

// checks if the RcwProjects type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &RcwProjects{}

// RcwProjects struct for RcwProjects
type RcwProjects struct {
	// Note: This is a Primary Key.<pk/>
	Id *int32 `json:"id,omitempty"`
	Name string `json:"name"`
	Status string `json:"status"`
	// Note: This is a Foreign Key to `users.id`.<fk table='users' column='id'/>
	CreatedByUserId *int32 `json:"created_by_user_id,omitempty"`
	CreatedAt *string `json:"created_at,omitempty"`
	UpdatedAt *string `json:"updated_at,omitempty"`
	DeletedAt *string `json:"deleted_at,omitempty"`
	LaunchDate *string `json:"launch_date,omitempty"`
	// Note: This is a Foreign Key to `users.id`.<fk table='users' column='id'/>
	LaunchByUserId *int32 `json:"launch_by_user_id,omitempty"`
	CompletedDate *string `json:"completed_date,omitempty"`
	// Note: This is a Foreign Key to `users.id`.<fk table='users' column='id'/>
	CompletedByUserId *int32 `json:"completed_by_user_id,omitempty"`
	DueDate *string `json:"due_date,omitempty"`
	PeriodStart *string `json:"period_start,omitempty"`
	PeriodEnd *string `json:"period_end,omitempty"`
	Totals interface{} `json:"totals,omitempty"`
	Type string `json:"type"`
	IsDraft bool `json:"is_draft"`
}

type _RcwProjects RcwProjects

// NewRcwProjects instantiates a new RcwProjects object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRcwProjects(name string, status string, type_ string, isDraft bool) *RcwProjects {
	this := RcwProjects{}
	this.Name = name
	this.Status = status
	this.Type = type_
	this.IsDraft = isDraft
	return &this
}

// NewRcwProjectsWithDefaults instantiates a new RcwProjects object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRcwProjectsWithDefaults() *RcwProjects {
	this := RcwProjects{}
	var type_ string = "Generic"
	this.Type = type_
	var isDraft bool = true
	this.IsDraft = isDraft
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *RcwProjects) GetId() int32 {
	if o == nil || IsNil(o.Id) {
		var ret int32
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RcwProjects) GetIdOk() (*int32, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *RcwProjects) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given int32 and assigns it to the Id field.
func (o *RcwProjects) SetId(v int32) {
	o.Id = &v
}

// GetName returns the Name field value
func (o *RcwProjects) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *RcwProjects) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *RcwProjects) SetName(v string) {
	o.Name = v
}

// GetStatus returns the Status field value
func (o *RcwProjects) GetStatus() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Status
}

// GetStatusOk returns a tuple with the Status field value
// and a boolean to check if the value has been set.
func (o *RcwProjects) GetStatusOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Status, true
}

// SetStatus sets field value
func (o *RcwProjects) SetStatus(v string) {
	o.Status = v
}

// GetCreatedByUserId returns the CreatedByUserId field value if set, zero value otherwise.
func (o *RcwProjects) GetCreatedByUserId() int32 {
	if o == nil || IsNil(o.CreatedByUserId) {
		var ret int32
		return ret
	}
	return *o.CreatedByUserId
}

// GetCreatedByUserIdOk returns a tuple with the CreatedByUserId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RcwProjects) GetCreatedByUserIdOk() (*int32, bool) {
	if o == nil || IsNil(o.CreatedByUserId) {
		return nil, false
	}
	return o.CreatedByUserId, true
}

// HasCreatedByUserId returns a boolean if a field has been set.
func (o *RcwProjects) HasCreatedByUserId() bool {
	if o != nil && !IsNil(o.CreatedByUserId) {
		return true
	}

	return false
}

// SetCreatedByUserId gets a reference to the given int32 and assigns it to the CreatedByUserId field.
func (o *RcwProjects) SetCreatedByUserId(v int32) {
	o.CreatedByUserId = &v
}

// GetCreatedAt returns the CreatedAt field value if set, zero value otherwise.
func (o *RcwProjects) GetCreatedAt() string {
	if o == nil || IsNil(o.CreatedAt) {
		var ret string
		return ret
	}
	return *o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RcwProjects) GetCreatedAtOk() (*string, bool) {
	if o == nil || IsNil(o.CreatedAt) {
		return nil, false
	}
	return o.CreatedAt, true
}

// HasCreatedAt returns a boolean if a field has been set.
func (o *RcwProjects) HasCreatedAt() bool {
	if o != nil && !IsNil(o.CreatedAt) {
		return true
	}

	return false
}

// SetCreatedAt gets a reference to the given string and assigns it to the CreatedAt field.
func (o *RcwProjects) SetCreatedAt(v string) {
	o.CreatedAt = &v
}

// GetUpdatedAt returns the UpdatedAt field value if set, zero value otherwise.
func (o *RcwProjects) GetUpdatedAt() string {
	if o == nil || IsNil(o.UpdatedAt) {
		var ret string
		return ret
	}
	return *o.UpdatedAt
}

// GetUpdatedAtOk returns a tuple with the UpdatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RcwProjects) GetUpdatedAtOk() (*string, bool) {
	if o == nil || IsNil(o.UpdatedAt) {
		return nil, false
	}
	return o.UpdatedAt, true
}

// HasUpdatedAt returns a boolean if a field has been set.
func (o *RcwProjects) HasUpdatedAt() bool {
	if o != nil && !IsNil(o.UpdatedAt) {
		return true
	}

	return false
}

// SetUpdatedAt gets a reference to the given string and assigns it to the UpdatedAt field.
func (o *RcwProjects) SetUpdatedAt(v string) {
	o.UpdatedAt = &v
}

// GetDeletedAt returns the DeletedAt field value if set, zero value otherwise.
func (o *RcwProjects) GetDeletedAt() string {
	if o == nil || IsNil(o.DeletedAt) {
		var ret string
		return ret
	}
	return *o.DeletedAt
}

// GetDeletedAtOk returns a tuple with the DeletedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RcwProjects) GetDeletedAtOk() (*string, bool) {
	if o == nil || IsNil(o.DeletedAt) {
		return nil, false
	}
	return o.DeletedAt, true
}

// HasDeletedAt returns a boolean if a field has been set.
func (o *RcwProjects) HasDeletedAt() bool {
	if o != nil && !IsNil(o.DeletedAt) {
		return true
	}

	return false
}

// SetDeletedAt gets a reference to the given string and assigns it to the DeletedAt field.
func (o *RcwProjects) SetDeletedAt(v string) {
	o.DeletedAt = &v
}

// GetLaunchDate returns the LaunchDate field value if set, zero value otherwise.
func (o *RcwProjects) GetLaunchDate() string {
	if o == nil || IsNil(o.LaunchDate) {
		var ret string
		return ret
	}
	return *o.LaunchDate
}

// GetLaunchDateOk returns a tuple with the LaunchDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RcwProjects) GetLaunchDateOk() (*string, bool) {
	if o == nil || IsNil(o.LaunchDate) {
		return nil, false
	}
	return o.LaunchDate, true
}

// HasLaunchDate returns a boolean if a field has been set.
func (o *RcwProjects) HasLaunchDate() bool {
	if o != nil && !IsNil(o.LaunchDate) {
		return true
	}

	return false
}

// SetLaunchDate gets a reference to the given string and assigns it to the LaunchDate field.
func (o *RcwProjects) SetLaunchDate(v string) {
	o.LaunchDate = &v
}

// GetLaunchByUserId returns the LaunchByUserId field value if set, zero value otherwise.
func (o *RcwProjects) GetLaunchByUserId() int32 {
	if o == nil || IsNil(o.LaunchByUserId) {
		var ret int32
		return ret
	}
	return *o.LaunchByUserId
}

// GetLaunchByUserIdOk returns a tuple with the LaunchByUserId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RcwProjects) GetLaunchByUserIdOk() (*int32, bool) {
	if o == nil || IsNil(o.LaunchByUserId) {
		return nil, false
	}
	return o.LaunchByUserId, true
}

// HasLaunchByUserId returns a boolean if a field has been set.
func (o *RcwProjects) HasLaunchByUserId() bool {
	if o != nil && !IsNil(o.LaunchByUserId) {
		return true
	}

	return false
}

// SetLaunchByUserId gets a reference to the given int32 and assigns it to the LaunchByUserId field.
func (o *RcwProjects) SetLaunchByUserId(v int32) {
	o.LaunchByUserId = &v
}

// GetCompletedDate returns the CompletedDate field value if set, zero value otherwise.
func (o *RcwProjects) GetCompletedDate() string {
	if o == nil || IsNil(o.CompletedDate) {
		var ret string
		return ret
	}
	return *o.CompletedDate
}

// GetCompletedDateOk returns a tuple with the CompletedDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RcwProjects) GetCompletedDateOk() (*string, bool) {
	if o == nil || IsNil(o.CompletedDate) {
		return nil, false
	}
	return o.CompletedDate, true
}

// HasCompletedDate returns a boolean if a field has been set.
func (o *RcwProjects) HasCompletedDate() bool {
	if o != nil && !IsNil(o.CompletedDate) {
		return true
	}

	return false
}

// SetCompletedDate gets a reference to the given string and assigns it to the CompletedDate field.
func (o *RcwProjects) SetCompletedDate(v string) {
	o.CompletedDate = &v
}

// GetCompletedByUserId returns the CompletedByUserId field value if set, zero value otherwise.
func (o *RcwProjects) GetCompletedByUserId() int32 {
	if o == nil || IsNil(o.CompletedByUserId) {
		var ret int32
		return ret
	}
	return *o.CompletedByUserId
}

// GetCompletedByUserIdOk returns a tuple with the CompletedByUserId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RcwProjects) GetCompletedByUserIdOk() (*int32, bool) {
	if o == nil || IsNil(o.CompletedByUserId) {
		return nil, false
	}
	return o.CompletedByUserId, true
}

// HasCompletedByUserId returns a boolean if a field has been set.
func (o *RcwProjects) HasCompletedByUserId() bool {
	if o != nil && !IsNil(o.CompletedByUserId) {
		return true
	}

	return false
}

// SetCompletedByUserId gets a reference to the given int32 and assigns it to the CompletedByUserId field.
func (o *RcwProjects) SetCompletedByUserId(v int32) {
	o.CompletedByUserId = &v
}

// GetDueDate returns the DueDate field value if set, zero value otherwise.
func (o *RcwProjects) GetDueDate() string {
	if o == nil || IsNil(o.DueDate) {
		var ret string
		return ret
	}
	return *o.DueDate
}

// GetDueDateOk returns a tuple with the DueDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RcwProjects) GetDueDateOk() (*string, bool) {
	if o == nil || IsNil(o.DueDate) {
		return nil, false
	}
	return o.DueDate, true
}

// HasDueDate returns a boolean if a field has been set.
func (o *RcwProjects) HasDueDate() bool {
	if o != nil && !IsNil(o.DueDate) {
		return true
	}

	return false
}

// SetDueDate gets a reference to the given string and assigns it to the DueDate field.
func (o *RcwProjects) SetDueDate(v string) {
	o.DueDate = &v
}

// GetPeriodStart returns the PeriodStart field value if set, zero value otherwise.
func (o *RcwProjects) GetPeriodStart() string {
	if o == nil || IsNil(o.PeriodStart) {
		var ret string
		return ret
	}
	return *o.PeriodStart
}

// GetPeriodStartOk returns a tuple with the PeriodStart field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RcwProjects) GetPeriodStartOk() (*string, bool) {
	if o == nil || IsNil(o.PeriodStart) {
		return nil, false
	}
	return o.PeriodStart, true
}

// HasPeriodStart returns a boolean if a field has been set.
func (o *RcwProjects) HasPeriodStart() bool {
	if o != nil && !IsNil(o.PeriodStart) {
		return true
	}

	return false
}

// SetPeriodStart gets a reference to the given string and assigns it to the PeriodStart field.
func (o *RcwProjects) SetPeriodStart(v string) {
	o.PeriodStart = &v
}

// GetPeriodEnd returns the PeriodEnd field value if set, zero value otherwise.
func (o *RcwProjects) GetPeriodEnd() string {
	if o == nil || IsNil(o.PeriodEnd) {
		var ret string
		return ret
	}
	return *o.PeriodEnd
}

// GetPeriodEndOk returns a tuple with the PeriodEnd field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RcwProjects) GetPeriodEndOk() (*string, bool) {
	if o == nil || IsNil(o.PeriodEnd) {
		return nil, false
	}
	return o.PeriodEnd, true
}

// HasPeriodEnd returns a boolean if a field has been set.
func (o *RcwProjects) HasPeriodEnd() bool {
	if o != nil && !IsNil(o.PeriodEnd) {
		return true
	}

	return false
}

// SetPeriodEnd gets a reference to the given string and assigns it to the PeriodEnd field.
func (o *RcwProjects) SetPeriodEnd(v string) {
	o.PeriodEnd = &v
}

// GetTotals returns the Totals field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *RcwProjects) GetTotals() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.Totals
}

// GetTotalsOk returns a tuple with the Totals field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *RcwProjects) GetTotalsOk() (*interface{}, bool) {
	if o == nil || IsNil(o.Totals) {
		return nil, false
	}
	return &o.Totals, true
}

// HasTotals returns a boolean if a field has been set.
func (o *RcwProjects) HasTotals() bool {
	if o != nil && !IsNil(o.Totals) {
		return true
	}

	return false
}

// SetTotals gets a reference to the given interface{} and assigns it to the Totals field.
func (o *RcwProjects) SetTotals(v interface{}) {
	o.Totals = v
}

// GetType returns the Type field value
func (o *RcwProjects) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *RcwProjects) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *RcwProjects) SetType(v string) {
	o.Type = v
}

// GetIsDraft returns the IsDraft field value
func (o *RcwProjects) GetIsDraft() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.IsDraft
}

// GetIsDraftOk returns a tuple with the IsDraft field value
// and a boolean to check if the value has been set.
func (o *RcwProjects) GetIsDraftOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.IsDraft, true
}

// SetIsDraft sets field value
func (o *RcwProjects) SetIsDraft(v bool) {
	o.IsDraft = v
}

func (o RcwProjects) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o RcwProjects) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	toSerialize["name"] = o.Name
	toSerialize["status"] = o.Status
	if !IsNil(o.CreatedByUserId) {
		toSerialize["created_by_user_id"] = o.CreatedByUserId
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
	if !IsNil(o.LaunchDate) {
		toSerialize["launch_date"] = o.LaunchDate
	}
	if !IsNil(o.LaunchByUserId) {
		toSerialize["launch_by_user_id"] = o.LaunchByUserId
	}
	if !IsNil(o.CompletedDate) {
		toSerialize["completed_date"] = o.CompletedDate
	}
	if !IsNil(o.CompletedByUserId) {
		toSerialize["completed_by_user_id"] = o.CompletedByUserId
	}
	if !IsNil(o.DueDate) {
		toSerialize["due_date"] = o.DueDate
	}
	if !IsNil(o.PeriodStart) {
		toSerialize["period_start"] = o.PeriodStart
	}
	if !IsNil(o.PeriodEnd) {
		toSerialize["period_end"] = o.PeriodEnd
	}
	if o.Totals != nil {
		toSerialize["totals"] = o.Totals
	}
	toSerialize["type"] = o.Type
	toSerialize["is_draft"] = o.IsDraft
	return toSerialize, nil
}

func (o *RcwProjects) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"name",
		"status",
		"type",
		"is_draft",
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

	varRcwProjects := _RcwProjects{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varRcwProjects)

	if err != nil {
		return err
	}

	*o = RcwProjects(varRcwProjects)

	return err
}

type NullableRcwProjects struct {
	value *RcwProjects
	isSet bool
}

func (v NullableRcwProjects) Get() *RcwProjects {
	return v.value
}

func (v *NullableRcwProjects) Set(val *RcwProjects) {
	v.value = val
	v.isSet = true
}

func (v NullableRcwProjects) IsSet() bool {
	return v.isSet
}

func (v *NullableRcwProjects) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRcwProjects(val *RcwProjects) *NullableRcwProjects {
	return &NullableRcwProjects{value: val, isSet: true}
}

func (v NullableRcwProjects) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRcwProjects) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

