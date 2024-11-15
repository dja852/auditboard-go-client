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

// checks if the EntityRisksPutPreviousValues type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &EntityRisksPutPreviousValues{}

// EntityRisksPutPreviousValues struct for EntityRisksPutPreviousValues
type EntityRisksPutPreviousValues struct {
	// Note: This is a Primary Key.<pk/>
	Id *int32 `json:"id,omitempty"`
	// Note: This is a Foreign Key to `entities.id`.<fk table='entities' column='id'/>
	EntityId *int32 `json:"entity_id,omitempty"`
	// Note: This is a Foreign Key to `risks.id`.<fk table='risks' column='id'/>
	RiskId *int32 `json:"risk_id,omitempty"`
	Status *string `json:"status,omitempty"`
	CreatedAt *string `json:"created_at,omitempty"`
	UpdatedAt *string `json:"updated_at,omitempty"`
	DeletedAt *string `json:"deleted_at,omitempty"`
	Uid *string `json:"uid,omitempty"`
	Description *string `json:"description,omitempty"`
	Activity *string `json:"activity,omitempty"`
	MitigationFactors *string `json:"mitigation_factors,omitempty"`
	CustomText1 *string `json:"custom_text1,omitempty"`
	CustomText2 *string `json:"custom_text2,omitempty"`
	// Note: This is a Foreign Key to `er_custom_select1_options.id`.<fk table='er_custom_select1_options' column='id'/>
	ErCustomSelect1OptionId *int32 `json:"er_custom_select1_option_id,omitempty"`
	// Note: This is a Foreign Key to `er_custom_select2_options.id`.<fk table='er_custom_select2_options' column='id'/>
	ErCustomSelect2OptionId *int32 `json:"er_custom_select2_option_id,omitempty"`
	// Note: This is a Foreign Key to `risk_categories.id`.<fk table='risk_categories' column='id'/>
	RiskCategoryId *int32 `json:"risk_category_id,omitempty"`
}

// NewEntityRisksPutPreviousValues instantiates a new EntityRisksPutPreviousValues object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewEntityRisksPutPreviousValues() *EntityRisksPutPreviousValues {
	this := EntityRisksPutPreviousValues{}
	return &this
}

// NewEntityRisksPutPreviousValuesWithDefaults instantiates a new EntityRisksPutPreviousValues object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewEntityRisksPutPreviousValuesWithDefaults() *EntityRisksPutPreviousValues {
	this := EntityRisksPutPreviousValues{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *EntityRisksPutPreviousValues) GetId() int32 {
	if o == nil || IsNil(o.Id) {
		var ret int32
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EntityRisksPutPreviousValues) GetIdOk() (*int32, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *EntityRisksPutPreviousValues) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given int32 and assigns it to the Id field.
func (o *EntityRisksPutPreviousValues) SetId(v int32) {
	o.Id = &v
}

// GetEntityId returns the EntityId field value if set, zero value otherwise.
func (o *EntityRisksPutPreviousValues) GetEntityId() int32 {
	if o == nil || IsNil(o.EntityId) {
		var ret int32
		return ret
	}
	return *o.EntityId
}

// GetEntityIdOk returns a tuple with the EntityId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EntityRisksPutPreviousValues) GetEntityIdOk() (*int32, bool) {
	if o == nil || IsNil(o.EntityId) {
		return nil, false
	}
	return o.EntityId, true
}

// HasEntityId returns a boolean if a field has been set.
func (o *EntityRisksPutPreviousValues) HasEntityId() bool {
	if o != nil && !IsNil(o.EntityId) {
		return true
	}

	return false
}

// SetEntityId gets a reference to the given int32 and assigns it to the EntityId field.
func (o *EntityRisksPutPreviousValues) SetEntityId(v int32) {
	o.EntityId = &v
}

// GetRiskId returns the RiskId field value if set, zero value otherwise.
func (o *EntityRisksPutPreviousValues) GetRiskId() int32 {
	if o == nil || IsNil(o.RiskId) {
		var ret int32
		return ret
	}
	return *o.RiskId
}

// GetRiskIdOk returns a tuple with the RiskId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EntityRisksPutPreviousValues) GetRiskIdOk() (*int32, bool) {
	if o == nil || IsNil(o.RiskId) {
		return nil, false
	}
	return o.RiskId, true
}

// HasRiskId returns a boolean if a field has been set.
func (o *EntityRisksPutPreviousValues) HasRiskId() bool {
	if o != nil && !IsNil(o.RiskId) {
		return true
	}

	return false
}

// SetRiskId gets a reference to the given int32 and assigns it to the RiskId field.
func (o *EntityRisksPutPreviousValues) SetRiskId(v int32) {
	o.RiskId = &v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *EntityRisksPutPreviousValues) GetStatus() string {
	if o == nil || IsNil(o.Status) {
		var ret string
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EntityRisksPutPreviousValues) GetStatusOk() (*string, bool) {
	if o == nil || IsNil(o.Status) {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *EntityRisksPutPreviousValues) HasStatus() bool {
	if o != nil && !IsNil(o.Status) {
		return true
	}

	return false
}

// SetStatus gets a reference to the given string and assigns it to the Status field.
func (o *EntityRisksPutPreviousValues) SetStatus(v string) {
	o.Status = &v
}

// GetCreatedAt returns the CreatedAt field value if set, zero value otherwise.
func (o *EntityRisksPutPreviousValues) GetCreatedAt() string {
	if o == nil || IsNil(o.CreatedAt) {
		var ret string
		return ret
	}
	return *o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EntityRisksPutPreviousValues) GetCreatedAtOk() (*string, bool) {
	if o == nil || IsNil(o.CreatedAt) {
		return nil, false
	}
	return o.CreatedAt, true
}

// HasCreatedAt returns a boolean if a field has been set.
func (o *EntityRisksPutPreviousValues) HasCreatedAt() bool {
	if o != nil && !IsNil(o.CreatedAt) {
		return true
	}

	return false
}

// SetCreatedAt gets a reference to the given string and assigns it to the CreatedAt field.
func (o *EntityRisksPutPreviousValues) SetCreatedAt(v string) {
	o.CreatedAt = &v
}

// GetUpdatedAt returns the UpdatedAt field value if set, zero value otherwise.
func (o *EntityRisksPutPreviousValues) GetUpdatedAt() string {
	if o == nil || IsNil(o.UpdatedAt) {
		var ret string
		return ret
	}
	return *o.UpdatedAt
}

// GetUpdatedAtOk returns a tuple with the UpdatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EntityRisksPutPreviousValues) GetUpdatedAtOk() (*string, bool) {
	if o == nil || IsNil(o.UpdatedAt) {
		return nil, false
	}
	return o.UpdatedAt, true
}

// HasUpdatedAt returns a boolean if a field has been set.
func (o *EntityRisksPutPreviousValues) HasUpdatedAt() bool {
	if o != nil && !IsNil(o.UpdatedAt) {
		return true
	}

	return false
}

// SetUpdatedAt gets a reference to the given string and assigns it to the UpdatedAt field.
func (o *EntityRisksPutPreviousValues) SetUpdatedAt(v string) {
	o.UpdatedAt = &v
}

// GetDeletedAt returns the DeletedAt field value if set, zero value otherwise.
func (o *EntityRisksPutPreviousValues) GetDeletedAt() string {
	if o == nil || IsNil(o.DeletedAt) {
		var ret string
		return ret
	}
	return *o.DeletedAt
}

// GetDeletedAtOk returns a tuple with the DeletedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EntityRisksPutPreviousValues) GetDeletedAtOk() (*string, bool) {
	if o == nil || IsNil(o.DeletedAt) {
		return nil, false
	}
	return o.DeletedAt, true
}

// HasDeletedAt returns a boolean if a field has been set.
func (o *EntityRisksPutPreviousValues) HasDeletedAt() bool {
	if o != nil && !IsNil(o.DeletedAt) {
		return true
	}

	return false
}

// SetDeletedAt gets a reference to the given string and assigns it to the DeletedAt field.
func (o *EntityRisksPutPreviousValues) SetDeletedAt(v string) {
	o.DeletedAt = &v
}

// GetUid returns the Uid field value if set, zero value otherwise.
func (o *EntityRisksPutPreviousValues) GetUid() string {
	if o == nil || IsNil(o.Uid) {
		var ret string
		return ret
	}
	return *o.Uid
}

// GetUidOk returns a tuple with the Uid field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EntityRisksPutPreviousValues) GetUidOk() (*string, bool) {
	if o == nil || IsNil(o.Uid) {
		return nil, false
	}
	return o.Uid, true
}

// HasUid returns a boolean if a field has been set.
func (o *EntityRisksPutPreviousValues) HasUid() bool {
	if o != nil && !IsNil(o.Uid) {
		return true
	}

	return false
}

// SetUid gets a reference to the given string and assigns it to the Uid field.
func (o *EntityRisksPutPreviousValues) SetUid(v string) {
	o.Uid = &v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *EntityRisksPutPreviousValues) GetDescription() string {
	if o == nil || IsNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EntityRisksPutPreviousValues) GetDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *EntityRisksPutPreviousValues) HasDescription() bool {
	if o != nil && !IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *EntityRisksPutPreviousValues) SetDescription(v string) {
	o.Description = &v
}

// GetActivity returns the Activity field value if set, zero value otherwise.
func (o *EntityRisksPutPreviousValues) GetActivity() string {
	if o == nil || IsNil(o.Activity) {
		var ret string
		return ret
	}
	return *o.Activity
}

// GetActivityOk returns a tuple with the Activity field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EntityRisksPutPreviousValues) GetActivityOk() (*string, bool) {
	if o == nil || IsNil(o.Activity) {
		return nil, false
	}
	return o.Activity, true
}

// HasActivity returns a boolean if a field has been set.
func (o *EntityRisksPutPreviousValues) HasActivity() bool {
	if o != nil && !IsNil(o.Activity) {
		return true
	}

	return false
}

// SetActivity gets a reference to the given string and assigns it to the Activity field.
func (o *EntityRisksPutPreviousValues) SetActivity(v string) {
	o.Activity = &v
}

// GetMitigationFactors returns the MitigationFactors field value if set, zero value otherwise.
func (o *EntityRisksPutPreviousValues) GetMitigationFactors() string {
	if o == nil || IsNil(o.MitigationFactors) {
		var ret string
		return ret
	}
	return *o.MitigationFactors
}

// GetMitigationFactorsOk returns a tuple with the MitigationFactors field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EntityRisksPutPreviousValues) GetMitigationFactorsOk() (*string, bool) {
	if o == nil || IsNil(o.MitigationFactors) {
		return nil, false
	}
	return o.MitigationFactors, true
}

// HasMitigationFactors returns a boolean if a field has been set.
func (o *EntityRisksPutPreviousValues) HasMitigationFactors() bool {
	if o != nil && !IsNil(o.MitigationFactors) {
		return true
	}

	return false
}

// SetMitigationFactors gets a reference to the given string and assigns it to the MitigationFactors field.
func (o *EntityRisksPutPreviousValues) SetMitigationFactors(v string) {
	o.MitigationFactors = &v
}

// GetCustomText1 returns the CustomText1 field value if set, zero value otherwise.
func (o *EntityRisksPutPreviousValues) GetCustomText1() string {
	if o == nil || IsNil(o.CustomText1) {
		var ret string
		return ret
	}
	return *o.CustomText1
}

// GetCustomText1Ok returns a tuple with the CustomText1 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EntityRisksPutPreviousValues) GetCustomText1Ok() (*string, bool) {
	if o == nil || IsNil(o.CustomText1) {
		return nil, false
	}
	return o.CustomText1, true
}

// HasCustomText1 returns a boolean if a field has been set.
func (o *EntityRisksPutPreviousValues) HasCustomText1() bool {
	if o != nil && !IsNil(o.CustomText1) {
		return true
	}

	return false
}

// SetCustomText1 gets a reference to the given string and assigns it to the CustomText1 field.
func (o *EntityRisksPutPreviousValues) SetCustomText1(v string) {
	o.CustomText1 = &v
}

// GetCustomText2 returns the CustomText2 field value if set, zero value otherwise.
func (o *EntityRisksPutPreviousValues) GetCustomText2() string {
	if o == nil || IsNil(o.CustomText2) {
		var ret string
		return ret
	}
	return *o.CustomText2
}

// GetCustomText2Ok returns a tuple with the CustomText2 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EntityRisksPutPreviousValues) GetCustomText2Ok() (*string, bool) {
	if o == nil || IsNil(o.CustomText2) {
		return nil, false
	}
	return o.CustomText2, true
}

// HasCustomText2 returns a boolean if a field has been set.
func (o *EntityRisksPutPreviousValues) HasCustomText2() bool {
	if o != nil && !IsNil(o.CustomText2) {
		return true
	}

	return false
}

// SetCustomText2 gets a reference to the given string and assigns it to the CustomText2 field.
func (o *EntityRisksPutPreviousValues) SetCustomText2(v string) {
	o.CustomText2 = &v
}

// GetErCustomSelect1OptionId returns the ErCustomSelect1OptionId field value if set, zero value otherwise.
func (o *EntityRisksPutPreviousValues) GetErCustomSelect1OptionId() int32 {
	if o == nil || IsNil(o.ErCustomSelect1OptionId) {
		var ret int32
		return ret
	}
	return *o.ErCustomSelect1OptionId
}

// GetErCustomSelect1OptionIdOk returns a tuple with the ErCustomSelect1OptionId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EntityRisksPutPreviousValues) GetErCustomSelect1OptionIdOk() (*int32, bool) {
	if o == nil || IsNil(o.ErCustomSelect1OptionId) {
		return nil, false
	}
	return o.ErCustomSelect1OptionId, true
}

// HasErCustomSelect1OptionId returns a boolean if a field has been set.
func (o *EntityRisksPutPreviousValues) HasErCustomSelect1OptionId() bool {
	if o != nil && !IsNil(o.ErCustomSelect1OptionId) {
		return true
	}

	return false
}

// SetErCustomSelect1OptionId gets a reference to the given int32 and assigns it to the ErCustomSelect1OptionId field.
func (o *EntityRisksPutPreviousValues) SetErCustomSelect1OptionId(v int32) {
	o.ErCustomSelect1OptionId = &v
}

// GetErCustomSelect2OptionId returns the ErCustomSelect2OptionId field value if set, zero value otherwise.
func (o *EntityRisksPutPreviousValues) GetErCustomSelect2OptionId() int32 {
	if o == nil || IsNil(o.ErCustomSelect2OptionId) {
		var ret int32
		return ret
	}
	return *o.ErCustomSelect2OptionId
}

// GetErCustomSelect2OptionIdOk returns a tuple with the ErCustomSelect2OptionId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EntityRisksPutPreviousValues) GetErCustomSelect2OptionIdOk() (*int32, bool) {
	if o == nil || IsNil(o.ErCustomSelect2OptionId) {
		return nil, false
	}
	return o.ErCustomSelect2OptionId, true
}

// HasErCustomSelect2OptionId returns a boolean if a field has been set.
func (o *EntityRisksPutPreviousValues) HasErCustomSelect2OptionId() bool {
	if o != nil && !IsNil(o.ErCustomSelect2OptionId) {
		return true
	}

	return false
}

// SetErCustomSelect2OptionId gets a reference to the given int32 and assigns it to the ErCustomSelect2OptionId field.
func (o *EntityRisksPutPreviousValues) SetErCustomSelect2OptionId(v int32) {
	o.ErCustomSelect2OptionId = &v
}

// GetRiskCategoryId returns the RiskCategoryId field value if set, zero value otherwise.
func (o *EntityRisksPutPreviousValues) GetRiskCategoryId() int32 {
	if o == nil || IsNil(o.RiskCategoryId) {
		var ret int32
		return ret
	}
	return *o.RiskCategoryId
}

// GetRiskCategoryIdOk returns a tuple with the RiskCategoryId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EntityRisksPutPreviousValues) GetRiskCategoryIdOk() (*int32, bool) {
	if o == nil || IsNil(o.RiskCategoryId) {
		return nil, false
	}
	return o.RiskCategoryId, true
}

// HasRiskCategoryId returns a boolean if a field has been set.
func (o *EntityRisksPutPreviousValues) HasRiskCategoryId() bool {
	if o != nil && !IsNil(o.RiskCategoryId) {
		return true
	}

	return false
}

// SetRiskCategoryId gets a reference to the given int32 and assigns it to the RiskCategoryId field.
func (o *EntityRisksPutPreviousValues) SetRiskCategoryId(v int32) {
	o.RiskCategoryId = &v
}

func (o EntityRisksPutPreviousValues) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o EntityRisksPutPreviousValues) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !IsNil(o.EntityId) {
		toSerialize["entity_id"] = o.EntityId
	}
	if !IsNil(o.RiskId) {
		toSerialize["risk_id"] = o.RiskId
	}
	if !IsNil(o.Status) {
		toSerialize["status"] = o.Status
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
	if !IsNil(o.Uid) {
		toSerialize["uid"] = o.Uid
	}
	if !IsNil(o.Description) {
		toSerialize["description"] = o.Description
	}
	if !IsNil(o.Activity) {
		toSerialize["activity"] = o.Activity
	}
	if !IsNil(o.MitigationFactors) {
		toSerialize["mitigation_factors"] = o.MitigationFactors
	}
	if !IsNil(o.CustomText1) {
		toSerialize["custom_text1"] = o.CustomText1
	}
	if !IsNil(o.CustomText2) {
		toSerialize["custom_text2"] = o.CustomText2
	}
	if !IsNil(o.ErCustomSelect1OptionId) {
		toSerialize["er_custom_select1_option_id"] = o.ErCustomSelect1OptionId
	}
	if !IsNil(o.ErCustomSelect2OptionId) {
		toSerialize["er_custom_select2_option_id"] = o.ErCustomSelect2OptionId
	}
	if !IsNil(o.RiskCategoryId) {
		toSerialize["risk_category_id"] = o.RiskCategoryId
	}
	return toSerialize, nil
}

type NullableEntityRisksPutPreviousValues struct {
	value *EntityRisksPutPreviousValues
	isSet bool
}

func (v NullableEntityRisksPutPreviousValues) Get() *EntityRisksPutPreviousValues {
	return v.value
}

func (v *NullableEntityRisksPutPreviousValues) Set(val *EntityRisksPutPreviousValues) {
	v.value = val
	v.isSet = true
}

func (v NullableEntityRisksPutPreviousValues) IsSet() bool {
	return v.isSet
}

func (v *NullableEntityRisksPutPreviousValues) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEntityRisksPutPreviousValues(val *EntityRisksPutPreviousValues) *NullableEntityRisksPutPreviousValues {
	return &NullableEntityRisksPutPreviousValues{value: val, isSet: true}
}

func (v NullableEntityRisksPutPreviousValues) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEntityRisksPutPreviousValues) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


