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

// checks if the GlobalAuditsPutGlobalAudit type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GlobalAuditsPutGlobalAudit{}

// GlobalAuditsPutGlobalAudit struct for GlobalAuditsPutGlobalAudit
type GlobalAuditsPutGlobalAudit struct {
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

type _GlobalAuditsPutGlobalAudit GlobalAuditsPutGlobalAudit

// NewGlobalAuditsPutGlobalAudit instantiates a new GlobalAuditsPutGlobalAudit object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGlobalAuditsPutGlobalAudit(loggableId int32, loggableType string, action string, field string) *GlobalAuditsPutGlobalAudit {
	this := GlobalAuditsPutGlobalAudit{}
	this.LoggableId = loggableId
	this.LoggableType = loggableType
	this.Action = action
	this.Field = field
	return &this
}

// NewGlobalAuditsPutGlobalAuditWithDefaults instantiates a new GlobalAuditsPutGlobalAudit object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGlobalAuditsPutGlobalAuditWithDefaults() *GlobalAuditsPutGlobalAudit {
	this := GlobalAuditsPutGlobalAudit{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *GlobalAuditsPutGlobalAudit) GetId() int32 {
	if o == nil || IsNil(o.Id) {
		var ret int32
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GlobalAuditsPutGlobalAudit) GetIdOk() (*int32, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *GlobalAuditsPutGlobalAudit) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given int32 and assigns it to the Id field.
func (o *GlobalAuditsPutGlobalAudit) SetId(v int32) {
	o.Id = &v
}

// GetCreatedAt returns the CreatedAt field value if set, zero value otherwise.
func (o *GlobalAuditsPutGlobalAudit) GetCreatedAt() string {
	if o == nil || IsNil(o.CreatedAt) {
		var ret string
		return ret
	}
	return *o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GlobalAuditsPutGlobalAudit) GetCreatedAtOk() (*string, bool) {
	if o == nil || IsNil(o.CreatedAt) {
		return nil, false
	}
	return o.CreatedAt, true
}

// HasCreatedAt returns a boolean if a field has been set.
func (o *GlobalAuditsPutGlobalAudit) HasCreatedAt() bool {
	if o != nil && !IsNil(o.CreatedAt) {
		return true
	}

	return false
}

// SetCreatedAt gets a reference to the given string and assigns it to the CreatedAt field.
func (o *GlobalAuditsPutGlobalAudit) SetCreatedAt(v string) {
	o.CreatedAt = &v
}

// GetUpdatedAt returns the UpdatedAt field value if set, zero value otherwise.
func (o *GlobalAuditsPutGlobalAudit) GetUpdatedAt() string {
	if o == nil || IsNil(o.UpdatedAt) {
		var ret string
		return ret
	}
	return *o.UpdatedAt
}

// GetUpdatedAtOk returns a tuple with the UpdatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GlobalAuditsPutGlobalAudit) GetUpdatedAtOk() (*string, bool) {
	if o == nil || IsNil(o.UpdatedAt) {
		return nil, false
	}
	return o.UpdatedAt, true
}

// HasUpdatedAt returns a boolean if a field has been set.
func (o *GlobalAuditsPutGlobalAudit) HasUpdatedAt() bool {
	if o != nil && !IsNil(o.UpdatedAt) {
		return true
	}

	return false
}

// SetUpdatedAt gets a reference to the given string and assigns it to the UpdatedAt field.
func (o *GlobalAuditsPutGlobalAudit) SetUpdatedAt(v string) {
	o.UpdatedAt = &v
}

// GetUserId returns the UserId field value if set, zero value otherwise.
func (o *GlobalAuditsPutGlobalAudit) GetUserId() int32 {
	if o == nil || IsNil(o.UserId) {
		var ret int32
		return ret
	}
	return *o.UserId
}

// GetUserIdOk returns a tuple with the UserId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GlobalAuditsPutGlobalAudit) GetUserIdOk() (*int32, bool) {
	if o == nil || IsNil(o.UserId) {
		return nil, false
	}
	return o.UserId, true
}

// HasUserId returns a boolean if a field has been set.
func (o *GlobalAuditsPutGlobalAudit) HasUserId() bool {
	if o != nil && !IsNil(o.UserId) {
		return true
	}

	return false
}

// SetUserId gets a reference to the given int32 and assigns it to the UserId field.
func (o *GlobalAuditsPutGlobalAudit) SetUserId(v int32) {
	o.UserId = &v
}

// GetLoggableId returns the LoggableId field value
func (o *GlobalAuditsPutGlobalAudit) GetLoggableId() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.LoggableId
}

// GetLoggableIdOk returns a tuple with the LoggableId field value
// and a boolean to check if the value has been set.
func (o *GlobalAuditsPutGlobalAudit) GetLoggableIdOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.LoggableId, true
}

// SetLoggableId sets field value
func (o *GlobalAuditsPutGlobalAudit) SetLoggableId(v int32) {
	o.LoggableId = v
}

// GetLoggableType returns the LoggableType field value
func (o *GlobalAuditsPutGlobalAudit) GetLoggableType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.LoggableType
}

// GetLoggableTypeOk returns a tuple with the LoggableType field value
// and a boolean to check if the value has been set.
func (o *GlobalAuditsPutGlobalAudit) GetLoggableTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.LoggableType, true
}

// SetLoggableType sets field value
func (o *GlobalAuditsPutGlobalAudit) SetLoggableType(v string) {
	o.LoggableType = v
}

// GetAction returns the Action field value
func (o *GlobalAuditsPutGlobalAudit) GetAction() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Action
}

// GetActionOk returns a tuple with the Action field value
// and a boolean to check if the value has been set.
func (o *GlobalAuditsPutGlobalAudit) GetActionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Action, true
}

// SetAction sets field value
func (o *GlobalAuditsPutGlobalAudit) SetAction(v string) {
	o.Action = v
}

// GetField returns the Field field value
func (o *GlobalAuditsPutGlobalAudit) GetField() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Field
}

// GetFieldOk returns a tuple with the Field field value
// and a boolean to check if the value has been set.
func (o *GlobalAuditsPutGlobalAudit) GetFieldOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Field, true
}

// SetField sets field value
func (o *GlobalAuditsPutGlobalAudit) SetField(v string) {
	o.Field = v
}

// GetFromValue returns the FromValue field value if set, zero value otherwise.
func (o *GlobalAuditsPutGlobalAudit) GetFromValue() string {
	if o == nil || IsNil(o.FromValue) {
		var ret string
		return ret
	}
	return *o.FromValue
}

// GetFromValueOk returns a tuple with the FromValue field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GlobalAuditsPutGlobalAudit) GetFromValueOk() (*string, bool) {
	if o == nil || IsNil(o.FromValue) {
		return nil, false
	}
	return o.FromValue, true
}

// HasFromValue returns a boolean if a field has been set.
func (o *GlobalAuditsPutGlobalAudit) HasFromValue() bool {
	if o != nil && !IsNil(o.FromValue) {
		return true
	}

	return false
}

// SetFromValue gets a reference to the given string and assigns it to the FromValue field.
func (o *GlobalAuditsPutGlobalAudit) SetFromValue(v string) {
	o.FromValue = &v
}

// GetToValue returns the ToValue field value if set, zero value otherwise.
func (o *GlobalAuditsPutGlobalAudit) GetToValue() string {
	if o == nil || IsNil(o.ToValue) {
		var ret string
		return ret
	}
	return *o.ToValue
}

// GetToValueOk returns a tuple with the ToValue field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GlobalAuditsPutGlobalAudit) GetToValueOk() (*string, bool) {
	if o == nil || IsNil(o.ToValue) {
		return nil, false
	}
	return o.ToValue, true
}

// HasToValue returns a boolean if a field has been set.
func (o *GlobalAuditsPutGlobalAudit) HasToValue() bool {
	if o != nil && !IsNil(o.ToValue) {
		return true
	}

	return false
}

// SetToValue gets a reference to the given string and assigns it to the ToValue field.
func (o *GlobalAuditsPutGlobalAudit) SetToValue(v string) {
	o.ToValue = &v
}

// GetMessage returns the Message field value if set, zero value otherwise.
func (o *GlobalAuditsPutGlobalAudit) GetMessage() string {
	if o == nil || IsNil(o.Message) {
		var ret string
		return ret
	}
	return *o.Message
}

// GetMessageOk returns a tuple with the Message field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GlobalAuditsPutGlobalAudit) GetMessageOk() (*string, bool) {
	if o == nil || IsNil(o.Message) {
		return nil, false
	}
	return o.Message, true
}

// HasMessage returns a boolean if a field has been set.
func (o *GlobalAuditsPutGlobalAudit) HasMessage() bool {
	if o != nil && !IsNil(o.Message) {
		return true
	}

	return false
}

// SetMessage gets a reference to the given string and assigns it to the Message field.
func (o *GlobalAuditsPutGlobalAudit) SetMessage(v string) {
	o.Message = &v
}

func (o GlobalAuditsPutGlobalAudit) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GlobalAuditsPutGlobalAudit) ToMap() (map[string]interface{}, error) {
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

func (o *GlobalAuditsPutGlobalAudit) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"loggable_id",
		"loggable_type",
		"action",
		"field",
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

	varGlobalAuditsPutGlobalAudit := _GlobalAuditsPutGlobalAudit{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varGlobalAuditsPutGlobalAudit)

	if err != nil {
		return err
	}

	*o = GlobalAuditsPutGlobalAudit(varGlobalAuditsPutGlobalAudit)

	return err
}

type NullableGlobalAuditsPutGlobalAudit struct {
	value *GlobalAuditsPutGlobalAudit
	isSet bool
}

func (v NullableGlobalAuditsPutGlobalAudit) Get() *GlobalAuditsPutGlobalAudit {
	return v.value
}

func (v *NullableGlobalAuditsPutGlobalAudit) Set(val *GlobalAuditsPutGlobalAudit) {
	v.value = val
	v.isSet = true
}

func (v NullableGlobalAuditsPutGlobalAudit) IsSet() bool {
	return v.isSet
}

func (v *NullableGlobalAuditsPutGlobalAudit) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGlobalAuditsPutGlobalAudit(val *GlobalAuditsPutGlobalAudit) *NullableGlobalAuditsPutGlobalAudit {
	return &NullableGlobalAuditsPutGlobalAudit{value: val, isSet: true}
}

func (v NullableGlobalAuditsPutGlobalAudit) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGlobalAuditsPutGlobalAudit) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

