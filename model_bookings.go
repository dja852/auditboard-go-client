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

// checks if the Bookings type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &Bookings{}

// Bookings struct for Bookings
type Bookings struct {
	// Note: This is a Primary Key.<pk/>
	Id *int32 `json:"id,omitempty"`
	// Note: This is a Foreign Key to `resources.id`.<fk table='resources' column='id'/>
	ResourceId int32 `json:"resource_id"`
	// Note: This is a Foreign Key to `project_plans.id`.<fk table='project_plans' column='id'/>
	ProjectPlanId int32 `json:"project_plan_id"`
	StartDate *string `json:"start_date,omitempty"`
	EndDate *string `json:"end_date,omitempty"`
	CreatedAt *string `json:"created_at,omitempty"`
	UpdatedAt *string `json:"updated_at,omitempty"`
	DeletedAt *string `json:"deleted_at,omitempty"`
	BookedHours int32 `json:"booked_hours"`
	// Note: This is a Foreign Key to `project_plan_phases.id`.<fk table='project_plan_phases' column='id'/>
	ProjectPlanPhaseId *int32 `json:"project_plan_phase_id,omitempty"`
}

type _Bookings Bookings

// NewBookings instantiates a new Bookings object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBookings(resourceId int32, projectPlanId int32, bookedHours int32) *Bookings {
	this := Bookings{}
	this.ResourceId = resourceId
	this.ProjectPlanId = projectPlanId
	this.BookedHours = bookedHours
	return &this
}

// NewBookingsWithDefaults instantiates a new Bookings object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBookingsWithDefaults() *Bookings {
	this := Bookings{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *Bookings) GetId() int32 {
	if o == nil || IsNil(o.Id) {
		var ret int32
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Bookings) GetIdOk() (*int32, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *Bookings) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given int32 and assigns it to the Id field.
func (o *Bookings) SetId(v int32) {
	o.Id = &v
}

// GetResourceId returns the ResourceId field value
func (o *Bookings) GetResourceId() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.ResourceId
}

// GetResourceIdOk returns a tuple with the ResourceId field value
// and a boolean to check if the value has been set.
func (o *Bookings) GetResourceIdOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ResourceId, true
}

// SetResourceId sets field value
func (o *Bookings) SetResourceId(v int32) {
	o.ResourceId = v
}

// GetProjectPlanId returns the ProjectPlanId field value
func (o *Bookings) GetProjectPlanId() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.ProjectPlanId
}

// GetProjectPlanIdOk returns a tuple with the ProjectPlanId field value
// and a boolean to check if the value has been set.
func (o *Bookings) GetProjectPlanIdOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ProjectPlanId, true
}

// SetProjectPlanId sets field value
func (o *Bookings) SetProjectPlanId(v int32) {
	o.ProjectPlanId = v
}

// GetStartDate returns the StartDate field value if set, zero value otherwise.
func (o *Bookings) GetStartDate() string {
	if o == nil || IsNil(o.StartDate) {
		var ret string
		return ret
	}
	return *o.StartDate
}

// GetStartDateOk returns a tuple with the StartDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Bookings) GetStartDateOk() (*string, bool) {
	if o == nil || IsNil(o.StartDate) {
		return nil, false
	}
	return o.StartDate, true
}

// HasStartDate returns a boolean if a field has been set.
func (o *Bookings) HasStartDate() bool {
	if o != nil && !IsNil(o.StartDate) {
		return true
	}

	return false
}

// SetStartDate gets a reference to the given string and assigns it to the StartDate field.
func (o *Bookings) SetStartDate(v string) {
	o.StartDate = &v
}

// GetEndDate returns the EndDate field value if set, zero value otherwise.
func (o *Bookings) GetEndDate() string {
	if o == nil || IsNil(o.EndDate) {
		var ret string
		return ret
	}
	return *o.EndDate
}

// GetEndDateOk returns a tuple with the EndDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Bookings) GetEndDateOk() (*string, bool) {
	if o == nil || IsNil(o.EndDate) {
		return nil, false
	}
	return o.EndDate, true
}

// HasEndDate returns a boolean if a field has been set.
func (o *Bookings) HasEndDate() bool {
	if o != nil && !IsNil(o.EndDate) {
		return true
	}

	return false
}

// SetEndDate gets a reference to the given string and assigns it to the EndDate field.
func (o *Bookings) SetEndDate(v string) {
	o.EndDate = &v
}

// GetCreatedAt returns the CreatedAt field value if set, zero value otherwise.
func (o *Bookings) GetCreatedAt() string {
	if o == nil || IsNil(o.CreatedAt) {
		var ret string
		return ret
	}
	return *o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Bookings) GetCreatedAtOk() (*string, bool) {
	if o == nil || IsNil(o.CreatedAt) {
		return nil, false
	}
	return o.CreatedAt, true
}

// HasCreatedAt returns a boolean if a field has been set.
func (o *Bookings) HasCreatedAt() bool {
	if o != nil && !IsNil(o.CreatedAt) {
		return true
	}

	return false
}

// SetCreatedAt gets a reference to the given string and assigns it to the CreatedAt field.
func (o *Bookings) SetCreatedAt(v string) {
	o.CreatedAt = &v
}

// GetUpdatedAt returns the UpdatedAt field value if set, zero value otherwise.
func (o *Bookings) GetUpdatedAt() string {
	if o == nil || IsNil(o.UpdatedAt) {
		var ret string
		return ret
	}
	return *o.UpdatedAt
}

// GetUpdatedAtOk returns a tuple with the UpdatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Bookings) GetUpdatedAtOk() (*string, bool) {
	if o == nil || IsNil(o.UpdatedAt) {
		return nil, false
	}
	return o.UpdatedAt, true
}

// HasUpdatedAt returns a boolean if a field has been set.
func (o *Bookings) HasUpdatedAt() bool {
	if o != nil && !IsNil(o.UpdatedAt) {
		return true
	}

	return false
}

// SetUpdatedAt gets a reference to the given string and assigns it to the UpdatedAt field.
func (o *Bookings) SetUpdatedAt(v string) {
	o.UpdatedAt = &v
}

// GetDeletedAt returns the DeletedAt field value if set, zero value otherwise.
func (o *Bookings) GetDeletedAt() string {
	if o == nil || IsNil(o.DeletedAt) {
		var ret string
		return ret
	}
	return *o.DeletedAt
}

// GetDeletedAtOk returns a tuple with the DeletedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Bookings) GetDeletedAtOk() (*string, bool) {
	if o == nil || IsNil(o.DeletedAt) {
		return nil, false
	}
	return o.DeletedAt, true
}

// HasDeletedAt returns a boolean if a field has been set.
func (o *Bookings) HasDeletedAt() bool {
	if o != nil && !IsNil(o.DeletedAt) {
		return true
	}

	return false
}

// SetDeletedAt gets a reference to the given string and assigns it to the DeletedAt field.
func (o *Bookings) SetDeletedAt(v string) {
	o.DeletedAt = &v
}

// GetBookedHours returns the BookedHours field value
func (o *Bookings) GetBookedHours() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.BookedHours
}

// GetBookedHoursOk returns a tuple with the BookedHours field value
// and a boolean to check if the value has been set.
func (o *Bookings) GetBookedHoursOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.BookedHours, true
}

// SetBookedHours sets field value
func (o *Bookings) SetBookedHours(v int32) {
	o.BookedHours = v
}

// GetProjectPlanPhaseId returns the ProjectPlanPhaseId field value if set, zero value otherwise.
func (o *Bookings) GetProjectPlanPhaseId() int32 {
	if o == nil || IsNil(o.ProjectPlanPhaseId) {
		var ret int32
		return ret
	}
	return *o.ProjectPlanPhaseId
}

// GetProjectPlanPhaseIdOk returns a tuple with the ProjectPlanPhaseId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Bookings) GetProjectPlanPhaseIdOk() (*int32, bool) {
	if o == nil || IsNil(o.ProjectPlanPhaseId) {
		return nil, false
	}
	return o.ProjectPlanPhaseId, true
}

// HasProjectPlanPhaseId returns a boolean if a field has been set.
func (o *Bookings) HasProjectPlanPhaseId() bool {
	if o != nil && !IsNil(o.ProjectPlanPhaseId) {
		return true
	}

	return false
}

// SetProjectPlanPhaseId gets a reference to the given int32 and assigns it to the ProjectPlanPhaseId field.
func (o *Bookings) SetProjectPlanPhaseId(v int32) {
	o.ProjectPlanPhaseId = &v
}

func (o Bookings) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o Bookings) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	toSerialize["resource_id"] = o.ResourceId
	toSerialize["project_plan_id"] = o.ProjectPlanId
	if !IsNil(o.StartDate) {
		toSerialize["start_date"] = o.StartDate
	}
	if !IsNil(o.EndDate) {
		toSerialize["end_date"] = o.EndDate
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
	toSerialize["booked_hours"] = o.BookedHours
	if !IsNil(o.ProjectPlanPhaseId) {
		toSerialize["project_plan_phase_id"] = o.ProjectPlanPhaseId
	}
	return toSerialize, nil
}

func (o *Bookings) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"resource_id",
		"project_plan_id",
		"booked_hours",
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

	varBookings := _Bookings{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varBookings)

	if err != nil {
		return err
	}

	*o = Bookings(varBookings)

	return err
}

type NullableBookings struct {
	value *Bookings
	isSet bool
}

func (v NullableBookings) Get() *Bookings {
	return v.value
}

func (v *NullableBookings) Set(val *Bookings) {
	v.value = val
	v.isSet = true
}

func (v NullableBookings) IsSet() bool {
	return v.isSet
}

func (v *NullableBookings) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBookings(val *Bookings) *NullableBookings {
	return &NullableBookings{value: val, isSet: true}
}

func (v NullableBookings) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBookings) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

