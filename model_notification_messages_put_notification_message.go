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

// checks if the NotificationMessagesPutNotificationMessage type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &NotificationMessagesPutNotificationMessage{}

// NotificationMessagesPutNotificationMessage struct for NotificationMessagesPutNotificationMessage
type NotificationMessagesPutNotificationMessage struct {
	// Note: This is a Primary Key.<pk/>
	Id *int32 `json:"id,omitempty"`
	// Note: This is a Foreign Key to `users.id`.<fk table='users' column='id'/>
	FromUserId *int32 `json:"from_user_id,omitempty"`
	// Note: This is a Foreign Key to `users.id`.<fk table='users' column='id'/>
	ToUserId int32 `json:"to_user_id"`
	Type *string `json:"type,omitempty"`
	Status *string `json:"status,omitempty"`
	IsRead bool `json:"is_read"`
	Message string `json:"message"`
	RawData *string `json:"raw_data,omitempty"`
	CreatedAt *string `json:"created_at,omitempty"`
	UpdatedAt *string `json:"updated_at,omitempty"`
	DeletedAt *string `json:"deleted_at,omitempty"`
}

type _NotificationMessagesPutNotificationMessage NotificationMessagesPutNotificationMessage

// NewNotificationMessagesPutNotificationMessage instantiates a new NotificationMessagesPutNotificationMessage object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewNotificationMessagesPutNotificationMessage(toUserId int32, isRead bool, message string) *NotificationMessagesPutNotificationMessage {
	this := NotificationMessagesPutNotificationMessage{}
	this.ToUserId = toUserId
	this.IsRead = isRead
	this.Message = message
	return &this
}

// NewNotificationMessagesPutNotificationMessageWithDefaults instantiates a new NotificationMessagesPutNotificationMessage object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewNotificationMessagesPutNotificationMessageWithDefaults() *NotificationMessagesPutNotificationMessage {
	this := NotificationMessagesPutNotificationMessage{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *NotificationMessagesPutNotificationMessage) GetId() int32 {
	if o == nil || IsNil(o.Id) {
		var ret int32
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NotificationMessagesPutNotificationMessage) GetIdOk() (*int32, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *NotificationMessagesPutNotificationMessage) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given int32 and assigns it to the Id field.
func (o *NotificationMessagesPutNotificationMessage) SetId(v int32) {
	o.Id = &v
}

// GetFromUserId returns the FromUserId field value if set, zero value otherwise.
func (o *NotificationMessagesPutNotificationMessage) GetFromUserId() int32 {
	if o == nil || IsNil(o.FromUserId) {
		var ret int32
		return ret
	}
	return *o.FromUserId
}

// GetFromUserIdOk returns a tuple with the FromUserId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NotificationMessagesPutNotificationMessage) GetFromUserIdOk() (*int32, bool) {
	if o == nil || IsNil(o.FromUserId) {
		return nil, false
	}
	return o.FromUserId, true
}

// HasFromUserId returns a boolean if a field has been set.
func (o *NotificationMessagesPutNotificationMessage) HasFromUserId() bool {
	if o != nil && !IsNil(o.FromUserId) {
		return true
	}

	return false
}

// SetFromUserId gets a reference to the given int32 and assigns it to the FromUserId field.
func (o *NotificationMessagesPutNotificationMessage) SetFromUserId(v int32) {
	o.FromUserId = &v
}

// GetToUserId returns the ToUserId field value
func (o *NotificationMessagesPutNotificationMessage) GetToUserId() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.ToUserId
}

// GetToUserIdOk returns a tuple with the ToUserId field value
// and a boolean to check if the value has been set.
func (o *NotificationMessagesPutNotificationMessage) GetToUserIdOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ToUserId, true
}

// SetToUserId sets field value
func (o *NotificationMessagesPutNotificationMessage) SetToUserId(v int32) {
	o.ToUserId = v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *NotificationMessagesPutNotificationMessage) GetType() string {
	if o == nil || IsNil(o.Type) {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NotificationMessagesPutNotificationMessage) GetTypeOk() (*string, bool) {
	if o == nil || IsNil(o.Type) {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *NotificationMessagesPutNotificationMessage) HasType() bool {
	if o != nil && !IsNil(o.Type) {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *NotificationMessagesPutNotificationMessage) SetType(v string) {
	o.Type = &v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *NotificationMessagesPutNotificationMessage) GetStatus() string {
	if o == nil || IsNil(o.Status) {
		var ret string
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NotificationMessagesPutNotificationMessage) GetStatusOk() (*string, bool) {
	if o == nil || IsNil(o.Status) {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *NotificationMessagesPutNotificationMessage) HasStatus() bool {
	if o != nil && !IsNil(o.Status) {
		return true
	}

	return false
}

// SetStatus gets a reference to the given string and assigns it to the Status field.
func (o *NotificationMessagesPutNotificationMessage) SetStatus(v string) {
	o.Status = &v
}

// GetIsRead returns the IsRead field value
func (o *NotificationMessagesPutNotificationMessage) GetIsRead() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.IsRead
}

// GetIsReadOk returns a tuple with the IsRead field value
// and a boolean to check if the value has been set.
func (o *NotificationMessagesPutNotificationMessage) GetIsReadOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.IsRead, true
}

// SetIsRead sets field value
func (o *NotificationMessagesPutNotificationMessage) SetIsRead(v bool) {
	o.IsRead = v
}

// GetMessage returns the Message field value
func (o *NotificationMessagesPutNotificationMessage) GetMessage() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Message
}

// GetMessageOk returns a tuple with the Message field value
// and a boolean to check if the value has been set.
func (o *NotificationMessagesPutNotificationMessage) GetMessageOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Message, true
}

// SetMessage sets field value
func (o *NotificationMessagesPutNotificationMessage) SetMessage(v string) {
	o.Message = v
}

// GetRawData returns the RawData field value if set, zero value otherwise.
func (o *NotificationMessagesPutNotificationMessage) GetRawData() string {
	if o == nil || IsNil(o.RawData) {
		var ret string
		return ret
	}
	return *o.RawData
}

// GetRawDataOk returns a tuple with the RawData field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NotificationMessagesPutNotificationMessage) GetRawDataOk() (*string, bool) {
	if o == nil || IsNil(o.RawData) {
		return nil, false
	}
	return o.RawData, true
}

// HasRawData returns a boolean if a field has been set.
func (o *NotificationMessagesPutNotificationMessage) HasRawData() bool {
	if o != nil && !IsNil(o.RawData) {
		return true
	}

	return false
}

// SetRawData gets a reference to the given string and assigns it to the RawData field.
func (o *NotificationMessagesPutNotificationMessage) SetRawData(v string) {
	o.RawData = &v
}

// GetCreatedAt returns the CreatedAt field value if set, zero value otherwise.
func (o *NotificationMessagesPutNotificationMessage) GetCreatedAt() string {
	if o == nil || IsNil(o.CreatedAt) {
		var ret string
		return ret
	}
	return *o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NotificationMessagesPutNotificationMessage) GetCreatedAtOk() (*string, bool) {
	if o == nil || IsNil(o.CreatedAt) {
		return nil, false
	}
	return o.CreatedAt, true
}

// HasCreatedAt returns a boolean if a field has been set.
func (o *NotificationMessagesPutNotificationMessage) HasCreatedAt() bool {
	if o != nil && !IsNil(o.CreatedAt) {
		return true
	}

	return false
}

// SetCreatedAt gets a reference to the given string and assigns it to the CreatedAt field.
func (o *NotificationMessagesPutNotificationMessage) SetCreatedAt(v string) {
	o.CreatedAt = &v
}

// GetUpdatedAt returns the UpdatedAt field value if set, zero value otherwise.
func (o *NotificationMessagesPutNotificationMessage) GetUpdatedAt() string {
	if o == nil || IsNil(o.UpdatedAt) {
		var ret string
		return ret
	}
	return *o.UpdatedAt
}

// GetUpdatedAtOk returns a tuple with the UpdatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NotificationMessagesPutNotificationMessage) GetUpdatedAtOk() (*string, bool) {
	if o == nil || IsNil(o.UpdatedAt) {
		return nil, false
	}
	return o.UpdatedAt, true
}

// HasUpdatedAt returns a boolean if a field has been set.
func (o *NotificationMessagesPutNotificationMessage) HasUpdatedAt() bool {
	if o != nil && !IsNil(o.UpdatedAt) {
		return true
	}

	return false
}

// SetUpdatedAt gets a reference to the given string and assigns it to the UpdatedAt field.
func (o *NotificationMessagesPutNotificationMessage) SetUpdatedAt(v string) {
	o.UpdatedAt = &v
}

// GetDeletedAt returns the DeletedAt field value if set, zero value otherwise.
func (o *NotificationMessagesPutNotificationMessage) GetDeletedAt() string {
	if o == nil || IsNil(o.DeletedAt) {
		var ret string
		return ret
	}
	return *o.DeletedAt
}

// GetDeletedAtOk returns a tuple with the DeletedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NotificationMessagesPutNotificationMessage) GetDeletedAtOk() (*string, bool) {
	if o == nil || IsNil(o.DeletedAt) {
		return nil, false
	}
	return o.DeletedAt, true
}

// HasDeletedAt returns a boolean if a field has been set.
func (o *NotificationMessagesPutNotificationMessage) HasDeletedAt() bool {
	if o != nil && !IsNil(o.DeletedAt) {
		return true
	}

	return false
}

// SetDeletedAt gets a reference to the given string and assigns it to the DeletedAt field.
func (o *NotificationMessagesPutNotificationMessage) SetDeletedAt(v string) {
	o.DeletedAt = &v
}

func (o NotificationMessagesPutNotificationMessage) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o NotificationMessagesPutNotificationMessage) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !IsNil(o.FromUserId) {
		toSerialize["from_user_id"] = o.FromUserId
	}
	toSerialize["to_user_id"] = o.ToUserId
	if !IsNil(o.Type) {
		toSerialize["type"] = o.Type
	}
	if !IsNil(o.Status) {
		toSerialize["status"] = o.Status
	}
	toSerialize["is_read"] = o.IsRead
	toSerialize["message"] = o.Message
	if !IsNil(o.RawData) {
		toSerialize["raw_data"] = o.RawData
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
	return toSerialize, nil
}

func (o *NotificationMessagesPutNotificationMessage) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"to_user_id",
		"is_read",
		"message",
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

	varNotificationMessagesPutNotificationMessage := _NotificationMessagesPutNotificationMessage{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varNotificationMessagesPutNotificationMessage)

	if err != nil {
		return err
	}

	*o = NotificationMessagesPutNotificationMessage(varNotificationMessagesPutNotificationMessage)

	return err
}

type NullableNotificationMessagesPutNotificationMessage struct {
	value *NotificationMessagesPutNotificationMessage
	isSet bool
}

func (v NullableNotificationMessagesPutNotificationMessage) Get() *NotificationMessagesPutNotificationMessage {
	return v.value
}

func (v *NullableNotificationMessagesPutNotificationMessage) Set(val *NotificationMessagesPutNotificationMessage) {
	v.value = val
	v.isSet = true
}

func (v NullableNotificationMessagesPutNotificationMessage) IsSet() bool {
	return v.isSet
}

func (v *NullableNotificationMessagesPutNotificationMessage) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableNotificationMessagesPutNotificationMessage(val *NotificationMessagesPutNotificationMessage) *NullableNotificationMessagesPutNotificationMessage {
	return &NullableNotificationMessagesPutNotificationMessage{value: val, isSet: true}
}

func (v NullableNotificationMessagesPutNotificationMessage) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableNotificationMessagesPutNotificationMessage) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


