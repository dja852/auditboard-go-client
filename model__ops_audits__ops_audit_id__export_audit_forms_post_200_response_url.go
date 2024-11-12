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

// checks if the OpsAuditsOpsAuditIdExportAuditFormsPost200ResponseUrl type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &OpsAuditsOpsAuditIdExportAuditFormsPost200ResponseUrl{}

// OpsAuditsOpsAuditIdExportAuditFormsPost200ResponseUrl struct for OpsAuditsOpsAuditIdExportAuditFormsPost200ResponseUrl
type OpsAuditsOpsAuditIdExportAuditFormsPost200ResponseUrl struct {
	Id *int32 `json:"id,omitempty"`
	Type *string `json:"type,omitempty"`
	Data *OpsAuditsOpsAuditIdExportAuditFormsPost200ResponseUrlData `json:"data,omitempty"`
	Priority *int32 `json:"priority,omitempty"`
	Progress *int32 `json:"progress,omitempty"`
	State *string `json:"state,omitempty"`
	CreatedAt *int64 `json:"created_at,omitempty"`
	PromoteAt *int64 `json:"promote_at,omitempty"`
	Attempts *OpsAuditsOpsAuditIdExportAuditFormsPost200ResponseUrlAttempts `json:"attempts,omitempty"`
}

// NewOpsAuditsOpsAuditIdExportAuditFormsPost200ResponseUrl instantiates a new OpsAuditsOpsAuditIdExportAuditFormsPost200ResponseUrl object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOpsAuditsOpsAuditIdExportAuditFormsPost200ResponseUrl() *OpsAuditsOpsAuditIdExportAuditFormsPost200ResponseUrl {
	this := OpsAuditsOpsAuditIdExportAuditFormsPost200ResponseUrl{}
	return &this
}

// NewOpsAuditsOpsAuditIdExportAuditFormsPost200ResponseUrlWithDefaults instantiates a new OpsAuditsOpsAuditIdExportAuditFormsPost200ResponseUrl object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOpsAuditsOpsAuditIdExportAuditFormsPost200ResponseUrlWithDefaults() *OpsAuditsOpsAuditIdExportAuditFormsPost200ResponseUrl {
	this := OpsAuditsOpsAuditIdExportAuditFormsPost200ResponseUrl{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *OpsAuditsOpsAuditIdExportAuditFormsPost200ResponseUrl) GetId() int32 {
	if o == nil || IsNil(o.Id) {
		var ret int32
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OpsAuditsOpsAuditIdExportAuditFormsPost200ResponseUrl) GetIdOk() (*int32, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *OpsAuditsOpsAuditIdExportAuditFormsPost200ResponseUrl) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given int32 and assigns it to the Id field.
func (o *OpsAuditsOpsAuditIdExportAuditFormsPost200ResponseUrl) SetId(v int32) {
	o.Id = &v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *OpsAuditsOpsAuditIdExportAuditFormsPost200ResponseUrl) GetType() string {
	if o == nil || IsNil(o.Type) {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OpsAuditsOpsAuditIdExportAuditFormsPost200ResponseUrl) GetTypeOk() (*string, bool) {
	if o == nil || IsNil(o.Type) {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *OpsAuditsOpsAuditIdExportAuditFormsPost200ResponseUrl) HasType() bool {
	if o != nil && !IsNil(o.Type) {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *OpsAuditsOpsAuditIdExportAuditFormsPost200ResponseUrl) SetType(v string) {
	o.Type = &v
}

// GetData returns the Data field value if set, zero value otherwise.
func (o *OpsAuditsOpsAuditIdExportAuditFormsPost200ResponseUrl) GetData() OpsAuditsOpsAuditIdExportAuditFormsPost200ResponseUrlData {
	if o == nil || IsNil(o.Data) {
		var ret OpsAuditsOpsAuditIdExportAuditFormsPost200ResponseUrlData
		return ret
	}
	return *o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OpsAuditsOpsAuditIdExportAuditFormsPost200ResponseUrl) GetDataOk() (*OpsAuditsOpsAuditIdExportAuditFormsPost200ResponseUrlData, bool) {
	if o == nil || IsNil(o.Data) {
		return nil, false
	}
	return o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *OpsAuditsOpsAuditIdExportAuditFormsPost200ResponseUrl) HasData() bool {
	if o != nil && !IsNil(o.Data) {
		return true
	}

	return false
}

// SetData gets a reference to the given OpsAuditsOpsAuditIdExportAuditFormsPost200ResponseUrlData and assigns it to the Data field.
func (o *OpsAuditsOpsAuditIdExportAuditFormsPost200ResponseUrl) SetData(v OpsAuditsOpsAuditIdExportAuditFormsPost200ResponseUrlData) {
	o.Data = &v
}

// GetPriority returns the Priority field value if set, zero value otherwise.
func (o *OpsAuditsOpsAuditIdExportAuditFormsPost200ResponseUrl) GetPriority() int32 {
	if o == nil || IsNil(o.Priority) {
		var ret int32
		return ret
	}
	return *o.Priority
}

// GetPriorityOk returns a tuple with the Priority field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OpsAuditsOpsAuditIdExportAuditFormsPost200ResponseUrl) GetPriorityOk() (*int32, bool) {
	if o == nil || IsNil(o.Priority) {
		return nil, false
	}
	return o.Priority, true
}

// HasPriority returns a boolean if a field has been set.
func (o *OpsAuditsOpsAuditIdExportAuditFormsPost200ResponseUrl) HasPriority() bool {
	if o != nil && !IsNil(o.Priority) {
		return true
	}

	return false
}

// SetPriority gets a reference to the given int32 and assigns it to the Priority field.
func (o *OpsAuditsOpsAuditIdExportAuditFormsPost200ResponseUrl) SetPriority(v int32) {
	o.Priority = &v
}

// GetProgress returns the Progress field value if set, zero value otherwise.
func (o *OpsAuditsOpsAuditIdExportAuditFormsPost200ResponseUrl) GetProgress() int32 {
	if o == nil || IsNil(o.Progress) {
		var ret int32
		return ret
	}
	return *o.Progress
}

// GetProgressOk returns a tuple with the Progress field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OpsAuditsOpsAuditIdExportAuditFormsPost200ResponseUrl) GetProgressOk() (*int32, bool) {
	if o == nil || IsNil(o.Progress) {
		return nil, false
	}
	return o.Progress, true
}

// HasProgress returns a boolean if a field has been set.
func (o *OpsAuditsOpsAuditIdExportAuditFormsPost200ResponseUrl) HasProgress() bool {
	if o != nil && !IsNil(o.Progress) {
		return true
	}

	return false
}

// SetProgress gets a reference to the given int32 and assigns it to the Progress field.
func (o *OpsAuditsOpsAuditIdExportAuditFormsPost200ResponseUrl) SetProgress(v int32) {
	o.Progress = &v
}

// GetState returns the State field value if set, zero value otherwise.
func (o *OpsAuditsOpsAuditIdExportAuditFormsPost200ResponseUrl) GetState() string {
	if o == nil || IsNil(o.State) {
		var ret string
		return ret
	}
	return *o.State
}

// GetStateOk returns a tuple with the State field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OpsAuditsOpsAuditIdExportAuditFormsPost200ResponseUrl) GetStateOk() (*string, bool) {
	if o == nil || IsNil(o.State) {
		return nil, false
	}
	return o.State, true
}

// HasState returns a boolean if a field has been set.
func (o *OpsAuditsOpsAuditIdExportAuditFormsPost200ResponseUrl) HasState() bool {
	if o != nil && !IsNil(o.State) {
		return true
	}

	return false
}

// SetState gets a reference to the given string and assigns it to the State field.
func (o *OpsAuditsOpsAuditIdExportAuditFormsPost200ResponseUrl) SetState(v string) {
	o.State = &v
}

// GetCreatedAt returns the CreatedAt field value if set, zero value otherwise.
func (o *OpsAuditsOpsAuditIdExportAuditFormsPost200ResponseUrl) GetCreatedAt() int64 {
	if o == nil || IsNil(o.CreatedAt) {
		var ret int64
		return ret
	}
	return *o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OpsAuditsOpsAuditIdExportAuditFormsPost200ResponseUrl) GetCreatedAtOk() (*int64, bool) {
	if o == nil || IsNil(o.CreatedAt) {
		return nil, false
	}
	return o.CreatedAt, true
}

// HasCreatedAt returns a boolean if a field has been set.
func (o *OpsAuditsOpsAuditIdExportAuditFormsPost200ResponseUrl) HasCreatedAt() bool {
	if o != nil && !IsNil(o.CreatedAt) {
		return true
	}

	return false
}

// SetCreatedAt gets a reference to the given int64 and assigns it to the CreatedAt field.
func (o *OpsAuditsOpsAuditIdExportAuditFormsPost200ResponseUrl) SetCreatedAt(v int64) {
	o.CreatedAt = &v
}

// GetPromoteAt returns the PromoteAt field value if set, zero value otherwise.
func (o *OpsAuditsOpsAuditIdExportAuditFormsPost200ResponseUrl) GetPromoteAt() int64 {
	if o == nil || IsNil(o.PromoteAt) {
		var ret int64
		return ret
	}
	return *o.PromoteAt
}

// GetPromoteAtOk returns a tuple with the PromoteAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OpsAuditsOpsAuditIdExportAuditFormsPost200ResponseUrl) GetPromoteAtOk() (*int64, bool) {
	if o == nil || IsNil(o.PromoteAt) {
		return nil, false
	}
	return o.PromoteAt, true
}

// HasPromoteAt returns a boolean if a field has been set.
func (o *OpsAuditsOpsAuditIdExportAuditFormsPost200ResponseUrl) HasPromoteAt() bool {
	if o != nil && !IsNil(o.PromoteAt) {
		return true
	}

	return false
}

// SetPromoteAt gets a reference to the given int64 and assigns it to the PromoteAt field.
func (o *OpsAuditsOpsAuditIdExportAuditFormsPost200ResponseUrl) SetPromoteAt(v int64) {
	o.PromoteAt = &v
}

// GetAttempts returns the Attempts field value if set, zero value otherwise.
func (o *OpsAuditsOpsAuditIdExportAuditFormsPost200ResponseUrl) GetAttempts() OpsAuditsOpsAuditIdExportAuditFormsPost200ResponseUrlAttempts {
	if o == nil || IsNil(o.Attempts) {
		var ret OpsAuditsOpsAuditIdExportAuditFormsPost200ResponseUrlAttempts
		return ret
	}
	return *o.Attempts
}

// GetAttemptsOk returns a tuple with the Attempts field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OpsAuditsOpsAuditIdExportAuditFormsPost200ResponseUrl) GetAttemptsOk() (*OpsAuditsOpsAuditIdExportAuditFormsPost200ResponseUrlAttempts, bool) {
	if o == nil || IsNil(o.Attempts) {
		return nil, false
	}
	return o.Attempts, true
}

// HasAttempts returns a boolean if a field has been set.
func (o *OpsAuditsOpsAuditIdExportAuditFormsPost200ResponseUrl) HasAttempts() bool {
	if o != nil && !IsNil(o.Attempts) {
		return true
	}

	return false
}

// SetAttempts gets a reference to the given OpsAuditsOpsAuditIdExportAuditFormsPost200ResponseUrlAttempts and assigns it to the Attempts field.
func (o *OpsAuditsOpsAuditIdExportAuditFormsPost200ResponseUrl) SetAttempts(v OpsAuditsOpsAuditIdExportAuditFormsPost200ResponseUrlAttempts) {
	o.Attempts = &v
}

func (o OpsAuditsOpsAuditIdExportAuditFormsPost200ResponseUrl) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o OpsAuditsOpsAuditIdExportAuditFormsPost200ResponseUrl) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !IsNil(o.Type) {
		toSerialize["type"] = o.Type
	}
	if !IsNil(o.Data) {
		toSerialize["data"] = o.Data
	}
	if !IsNil(o.Priority) {
		toSerialize["priority"] = o.Priority
	}
	if !IsNil(o.Progress) {
		toSerialize["progress"] = o.Progress
	}
	if !IsNil(o.State) {
		toSerialize["state"] = o.State
	}
	if !IsNil(o.CreatedAt) {
		toSerialize["created_at"] = o.CreatedAt
	}
	if !IsNil(o.PromoteAt) {
		toSerialize["promote_at"] = o.PromoteAt
	}
	if !IsNil(o.Attempts) {
		toSerialize["attempts"] = o.Attempts
	}
	return toSerialize, nil
}

type NullableOpsAuditsOpsAuditIdExportAuditFormsPost200ResponseUrl struct {
	value *OpsAuditsOpsAuditIdExportAuditFormsPost200ResponseUrl
	isSet bool
}

func (v NullableOpsAuditsOpsAuditIdExportAuditFormsPost200ResponseUrl) Get() *OpsAuditsOpsAuditIdExportAuditFormsPost200ResponseUrl {
	return v.value
}

func (v *NullableOpsAuditsOpsAuditIdExportAuditFormsPost200ResponseUrl) Set(val *OpsAuditsOpsAuditIdExportAuditFormsPost200ResponseUrl) {
	v.value = val
	v.isSet = true
}

func (v NullableOpsAuditsOpsAuditIdExportAuditFormsPost200ResponseUrl) IsSet() bool {
	return v.isSet
}

func (v *NullableOpsAuditsOpsAuditIdExportAuditFormsPost200ResponseUrl) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOpsAuditsOpsAuditIdExportAuditFormsPost200ResponseUrl(val *OpsAuditsOpsAuditIdExportAuditFormsPost200ResponseUrl) *NullableOpsAuditsOpsAuditIdExportAuditFormsPost200ResponseUrl {
	return &NullableOpsAuditsOpsAuditIdExportAuditFormsPost200ResponseUrl{value: val, isSet: true}
}

func (v NullableOpsAuditsOpsAuditIdExportAuditFormsPost200ResponseUrl) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOpsAuditsOpsAuditIdExportAuditFormsPost200ResponseUrl) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


