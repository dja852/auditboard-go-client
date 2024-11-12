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

// checks if the GlobalAudits type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GlobalAudits{}

// GlobalAudits struct for GlobalAudits
type GlobalAudits struct {
	Id *int32 `json:"id,omitempty"`
	CreatedAt *string `json:"created_at,omitempty"`
	UpdatedAt *string `json:"updated_at,omitempty"`
	// Note: This is a Foreign Key to `users.id`.<fk table='users' column='id'/>
	UserId *int32 `json:"user_id,omitempty"`
	LoggableId int32 `json:"loggable_id"`
	LoggableType string `json:"loggable_type"`
	Action string `json:"action"`
	Field string `json:"field"`
	FromValue *string `json:"from_value,omitempty"`
	ToValue *string `json:"to_value,omitempty"`
	Message *string `json:"message,omitempty"`
}

type _GlobalAudits GlobalAudits

// NewGlobalAudits instantiates a new GlobalAudits object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGlobalAudits(loggableId int32, loggableType string, action string, field string) *GlobalAudits {
	this := GlobalAudits{}
	this.LoggableId = loggableId
	this.LoggableType = loggableType
	this.Action = action
	this.Field = field
	return &this
}

// NewGlobalAuditsWithDefaults instantiates a new GlobalAudits object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGlobalAuditsWithDefaults() *GlobalAudits {
	this := GlobalAudits{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *GlobalAudits) GetId() int32 {
	if o == nil || IsNil(o.Id) {
		var ret int32
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GlobalAudits) GetIdOk() (*int32, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *GlobalAudits) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given int32 and assigns it to the Id field.
func (o *GlobalAudits) SetId(v int32) {
	o.Id = &v
}

// GetCreatedAt returns the CreatedAt field value if set, zero value otherwise.
func (o *GlobalAudits) GetCreatedAt() string {
	if o == nil || IsNil(o.CreatedAt) {
		var ret string
		return ret
	}
	return *o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GlobalAudits) GetCreatedAtOk() (*string, bool) {
	if o == nil || IsNil(o.CreatedAt) {
		return nil, false
	}
	return o.CreatedAt, true
}

// HasCreatedAt returns a boolean if a field has been set.
func (o *GlobalAudits) HasCreatedAt() bool {
	if o != nil && !IsNil(o.CreatedAt) {
		return true
	}

	return false
}

// SetCreatedAt gets a reference to the given string and assigns it to the CreatedAt field.
func (o *GlobalAudits) SetCreatedAt(v string) {
	o.CreatedAt = &v
}

// GetUpdatedAt returns the UpdatedAt field value if set, zero value otherwise.
func (o *GlobalAudits) GetUpdatedAt() string {
	if o == nil || IsNil(o.UpdatedAt) {
		var ret string
		return ret
	}
	return *o.UpdatedAt
}

// GetUpdatedAtOk returns a tuple with the UpdatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GlobalAudits) GetUpdatedAtOk() (*string, bool) {
	if o == nil || IsNil(o.UpdatedAt) {
		return nil, false
	}
	return o.UpdatedAt, true
}

// HasUpdatedAt returns a boolean if a field has been set.
func (o *GlobalAudits) HasUpdatedAt() bool {
	if o != nil && !IsNil(o.UpdatedAt) {
		return true
	}

	return false
}

// SetUpdatedAt gets a reference to the given string and assigns it to the UpdatedAt field.
func (o *GlobalAudits) SetUpdatedAt(v string) {
	o.UpdatedAt = &v
}

// GetUserId returns the UserId field value if set, zero value otherwise.
func (o *GlobalAudits) GetUserId() int32 {
	if o == nil || IsNil(o.UserId) {
		var ret int32
		return ret
	}
	return *o.UserId
}

// GetUserIdOk returns a tuple with the UserId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GlobalAudits) GetUserIdOk() (*int32, bool) {
	if o == nil || IsNil(o.UserId) {
		return nil, false
	}
	return o.UserId, true
}

// HasUserId returns a boolean if a field has been set.
func (o *GlobalAudits) HasUserId() bool {
	if o != nil && !IsNil(o.UserId) {
		return true
	}

	return false
}

// SetUserId gets a reference to the given int32 and assigns it to the UserId field.
func (o *GlobalAudits) SetUserId(v int32) {
	o.UserId = &v
}

// GetLoggableId returns the LoggableId field value
func (o *GlobalAudits) GetLoggableId() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.LoggableId
}

// GetLoggableIdOk returns a tuple with the LoggableId field value
// and a boolean to check if the value has been set.
func (o *GlobalAudits) GetLoggableIdOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.LoggableId, true
}

// SetLoggableId sets field value
func (o *GlobalAudits) SetLoggableId(v int32) {
	o.LoggableId = v
}

// GetLoggableType returns the LoggableType field value
func (o *GlobalAudits) GetLoggableType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.LoggableType
}

// GetLoggableTypeOk returns a tuple with the LoggableType field value
// and a boolean to check if the value has been set.
func (o *GlobalAudits) GetLoggableTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.LoggableType, true
}

// SetLoggableType sets field value
func (o *GlobalAudits) SetLoggableType(v string) {
	o.LoggableType = v
}

// GetAction returns the Action field value
func (o *GlobalAudits) GetAction() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Action
}

// GetActionOk returns a tuple with the Action field value
// and a boolean to check if the value has been set.
func (o *GlobalAudits) GetActionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Action, true
}

// SetAction sets field value
func (o *GlobalAudits) SetAction(v string) {
	o.Action = v
}

// GetField returns the Field field value
func (o *GlobalAudits) GetField() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Field
}

// GetFieldOk returns a tuple with the Field field value
// and a boolean to check if the value has been set.
func (o *GlobalAudits) GetFieldOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Field, true
}

// SetField sets field value
func (o *GlobalAudits) SetField(v string) {
	o.Field = v
}

// GetFromValue returns the FromValue field value if set, zero value otherwise.
func (o *GlobalAudits) GetFromValue() string {
	if o == nil || IsNil(o.FromValue) {
		var ret string
		return ret
	}
	return *o.FromValue
}

// GetFromValueOk returns a tuple with the FromValue field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GlobalAudits) GetFromValueOk() (*string, bool) {
	if o == nil || IsNil(o.FromValue) {
		return nil, false
	}
	return o.FromValue, true
}

// HasFromValue returns a boolean if a field has been set.
func (o *GlobalAudits) HasFromValue() bool {
	if o != nil && !IsNil(o.FromValue) {
		return true
	}

	return false
}

// SetFromValue gets a reference to the given string and assigns it to the FromValue field.
func (o *GlobalAudits) SetFromValue(v string) {
	o.FromValue = &v
}

// GetToValue returns the ToValue field value if set, zero value otherwise.
func (o *GlobalAudits) GetToValue() string {
	if o == nil || IsNil(o.ToValue) {
		var ret string
		return ret
	}
	return *o.ToValue
}

// GetToValueOk returns a tuple with the ToValue field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GlobalAudits) GetToValueOk() (*string, bool) {
	if o == nil || IsNil(o.ToValue) {
		return nil, false
	}
	return o.ToValue, true
}

// HasToValue returns a boolean if a field has been set.
func (o *GlobalAudits) HasToValue() bool {
	if o != nil && !IsNil(o.ToValue) {
		return true
	}

	return false
}

// SetToValue gets a reference to the given string and assigns it to the ToValue field.
func (o *GlobalAudits) SetToValue(v string) {
	o.ToValue = &v
}

// GetMessage returns the Message field value if set, zero value otherwise.
func (o *GlobalAudits) GetMessage() string {
	if o == nil || IsNil(o.Message) {
		var ret string
		return ret
	}
	return *o.Message
}

// GetMessageOk returns a tuple with the Message field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GlobalAudits) GetMessageOk() (*string, bool) {
	if o == nil || IsNil(o.Message) {
		return nil, false
	}
	return o.Message, true
}

// HasMessage returns a boolean if a field has been set.
func (o *GlobalAudits) HasMessage() bool {
	if o != nil && !IsNil(o.Message) {
		return true
	}

	return false
}

// SetMessage gets a reference to the given string and assigns it to the Message field.
func (o *GlobalAudits) SetMessage(v string) {
	o.Message = &v
}

func (o GlobalAudits) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GlobalAudits) ToMap() (map[string]interface{}, error) {
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
	if !IsNil(o.UserId) {
		toSerialize["user_id"] = o.UserId
	}
	toSerialize["loggable_id"] = o.LoggableId
	toSerialize["loggable_type"] = o.LoggableType
	toSerialize["action"] = o.Action
	toSerialize["field"] = o.Field
	if !IsNil(o.FromValue) {
		toSerialize["from_value"] = o.FromValue
	}
	if !IsNil(o.ToValue) {
		toSerialize["to_value"] = o.ToValue
	}
	if !IsNil(o.Message) {
		toSerialize["message"] = o.Message
	}
	return toSerialize, nil
}

type NullableGlobalAudits struct {
	value *GlobalAudits
	isSet bool
}

func (v NullableGlobalAudits) Get() *GlobalAudits {
	return v.value
}

func (v *NullableGlobalAudits) Set(val *GlobalAudits) {
	v.value = val
	v.isSet = true
}

func (v NullableGlobalAudits) IsSet() bool {
	return v.isSet
}

func (v *NullableGlobalAudits) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGlobalAudits(val *GlobalAudits) *NullableGlobalAudits {
	return &NullableGlobalAudits{value: val, isSet: true}
}

func (v NullableGlobalAudits) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGlobalAudits) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


