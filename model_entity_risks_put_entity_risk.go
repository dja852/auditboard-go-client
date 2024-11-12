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

// checks if the EntityRisksPutEntityRisk type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &EntityRisksPutEntityRisk{}

// EntityRisksPutEntityRisk struct for EntityRisksPutEntityRisk
type EntityRisksPutEntityRisk struct {
	// Note: This is a Primary Key.<pk/>
	Id *int32 `json:"id,omitempty"`
	// Note: This is a Foreign Key to `entities.id`.<fk table='entities' column='id'/>
	EntityId int32 `json:"entity_id"`
	// Note: This is a Foreign Key to `risks.id`.<fk table='risks' column='id'/>
	RiskId int32 `json:"risk_id"`
	Status string `json:"status"`
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

type _EntityRisksPutEntityRisk EntityRisksPutEntityRisk

// NewEntityRisksPutEntityRisk instantiates a new EntityRisksPutEntityRisk object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewEntityRisksPutEntityRisk(entityId int32, riskId int32, status string) *EntityRisksPutEntityRisk {
	this := EntityRisksPutEntityRisk{}
	this.EntityId = entityId
	this.RiskId = riskId
	this.Status = status
	return &this
}

// NewEntityRisksPutEntityRiskWithDefaults instantiates a new EntityRisksPutEntityRisk object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewEntityRisksPutEntityRiskWithDefaults() *EntityRisksPutEntityRisk {
	this := EntityRisksPutEntityRisk{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *EntityRisksPutEntityRisk) GetId() int32 {
	if o == nil || IsNil(o.Id) {
		var ret int32
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EntityRisksPutEntityRisk) GetIdOk() (*int32, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *EntityRisksPutEntityRisk) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given int32 and assigns it to the Id field.
func (o *EntityRisksPutEntityRisk) SetId(v int32) {
	o.Id = &v
}

// GetEntityId returns the EntityId field value
func (o *EntityRisksPutEntityRisk) GetEntityId() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.EntityId
}

// GetEntityIdOk returns a tuple with the EntityId field value
// and a boolean to check if the value has been set.
func (o *EntityRisksPutEntityRisk) GetEntityIdOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.EntityId, true
}

// SetEntityId sets field value
func (o *EntityRisksPutEntityRisk) SetEntityId(v int32) {
	o.EntityId = v
}

// GetRiskId returns the RiskId field value
func (o *EntityRisksPutEntityRisk) GetRiskId() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.RiskId
}

// GetRiskIdOk returns a tuple with the RiskId field value
// and a boolean to check if the value has been set.
func (o *EntityRisksPutEntityRisk) GetRiskIdOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.RiskId, true
}

// SetRiskId sets field value
func (o *EntityRisksPutEntityRisk) SetRiskId(v int32) {
	o.RiskId = v
}

// GetStatus returns the Status field value
func (o *EntityRisksPutEntityRisk) GetStatus() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Status
}

// GetStatusOk returns a tuple with the Status field value
// and a boolean to check if the value has been set.
func (o *EntityRisksPutEntityRisk) GetStatusOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Status, true
}

// SetStatus sets field value
func (o *EntityRisksPutEntityRisk) SetStatus(v string) {
	o.Status = v
}

// GetCreatedAt returns the CreatedAt field value if set, zero value otherwise.
func (o *EntityRisksPutEntityRisk) GetCreatedAt() string {
	if o == nil || IsNil(o.CreatedAt) {
		var ret string
		return ret
	}
	return *o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EntityRisksPutEntityRisk) GetCreatedAtOk() (*string, bool) {
	if o == nil || IsNil(o.CreatedAt) {
		return nil, false
	}
	return o.CreatedAt, true
}

// HasCreatedAt returns a boolean if a field has been set.
func (o *EntityRisksPutEntityRisk) HasCreatedAt() bool {
	if o != nil && !IsNil(o.CreatedAt) {
		return true
	}

	return false
}

// SetCreatedAt gets a reference to the given string and assigns it to the CreatedAt field.
func (o *EntityRisksPutEntityRisk) SetCreatedAt(v string) {
	o.CreatedAt = &v
}

// GetUpdatedAt returns the UpdatedAt field value if set, zero value otherwise.
func (o *EntityRisksPutEntityRisk) GetUpdatedAt() string {
	if o == nil || IsNil(o.UpdatedAt) {
		var ret string
		return ret
	}
	return *o.UpdatedAt
}

// GetUpdatedAtOk returns a tuple with the UpdatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EntityRisksPutEntityRisk) GetUpdatedAtOk() (*string, bool) {
	if o == nil || IsNil(o.UpdatedAt) {
		return nil, false
	}
	return o.UpdatedAt, true
}

// HasUpdatedAt returns a boolean if a field has been set.
func (o *EntityRisksPutEntityRisk) HasUpdatedAt() bool {
	if o != nil && !IsNil(o.UpdatedAt) {
		return true
	}

	return false
}

// SetUpdatedAt gets a reference to the given string and assigns it to the UpdatedAt field.
func (o *EntityRisksPutEntityRisk) SetUpdatedAt(v string) {
	o.UpdatedAt = &v
}

// GetDeletedAt returns the DeletedAt field value if set, zero value otherwise.
func (o *EntityRisksPutEntityRisk) GetDeletedAt() string {
	if o == nil || IsNil(o.DeletedAt) {
		var ret string
		return ret
	}
	return *o.DeletedAt
}

// GetDeletedAtOk returns a tuple with the DeletedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EntityRisksPutEntityRisk) GetDeletedAtOk() (*string, bool) {
	if o == nil || IsNil(o.DeletedAt) {
		return nil, false
	}
	return o.DeletedAt, true
}

// HasDeletedAt returns a boolean if a field has been set.
func (o *EntityRisksPutEntityRisk) HasDeletedAt() bool {
	if o != nil && !IsNil(o.DeletedAt) {
		return true
	}

	return false
}

// SetDeletedAt gets a reference to the given string and assigns it to the DeletedAt field.
func (o *EntityRisksPutEntityRisk) SetDeletedAt(v string) {
	o.DeletedAt = &v
}

// GetUid returns the Uid field value if set, zero value otherwise.
func (o *EntityRisksPutEntityRisk) GetUid() string {
	if o == nil || IsNil(o.Uid) {
		var ret string
		return ret
	}
	return *o.Uid
}

// GetUidOk returns a tuple with the Uid field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EntityRisksPutEntityRisk) GetUidOk() (*string, bool) {
	if o == nil || IsNil(o.Uid) {
		return nil, false
	}
	return o.Uid, true
}

// HasUid returns a boolean if a field has been set.
func (o *EntityRisksPutEntityRisk) HasUid() bool {
	if o != nil && !IsNil(o.Uid) {
		return true
	}

	return false
}

// SetUid gets a reference to the given string and assigns it to the Uid field.
func (o *EntityRisksPutEntityRisk) SetUid(v string) {
	o.Uid = &v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *EntityRisksPutEntityRisk) GetDescription() string {
	if o == nil || IsNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EntityRisksPutEntityRisk) GetDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *EntityRisksPutEntityRisk) HasDescription() bool {
	if o != nil && !IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *EntityRisksPutEntityRisk) SetDescription(v string) {
	o.Description = &v
}

// GetActivity returns the Activity field value if set, zero value otherwise.
func (o *EntityRisksPutEntityRisk) GetActivity() string {
	if o == nil || IsNil(o.Activity) {
		var ret string
		return ret
	}
	return *o.Activity
}

// GetActivityOk returns a tuple with the Activity field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EntityRisksPutEntityRisk) GetActivityOk() (*string, bool) {
	if o == nil || IsNil(o.Activity) {
		return nil, false
	}
	return o.Activity, true
}

// HasActivity returns a boolean if a field has been set.
func (o *EntityRisksPutEntityRisk) HasActivity() bool {
	if o != nil && !IsNil(o.Activity) {
		return true
	}

	return false
}

// SetActivity gets a reference to the given string and assigns it to the Activity field.
func (o *EntityRisksPutEntityRisk) SetActivity(v string) {
	o.Activity = &v
}

// GetMitigationFactors returns the MitigationFactors field value if set, zero value otherwise.
func (o *EntityRisksPutEntityRisk) GetMitigationFactors() string {
	if o == nil || IsNil(o.MitigationFactors) {
		var ret string
		return ret
	}
	return *o.MitigationFactors
}

// GetMitigationFactorsOk returns a tuple with the MitigationFactors field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EntityRisksPutEntityRisk) GetMitigationFactorsOk() (*string, bool) {
	if o == nil || IsNil(o.MitigationFactors) {
		return nil, false
	}
	return o.MitigationFactors, true
}

// HasMitigationFactors returns a boolean if a field has been set.
func (o *EntityRisksPutEntityRisk) HasMitigationFactors() bool {
	if o != nil && !IsNil(o.MitigationFactors) {
		return true
	}

	return false
}

// SetMitigationFactors gets a reference to the given string and assigns it to the MitigationFactors field.
func (o *EntityRisksPutEntityRisk) SetMitigationFactors(v string) {
	o.MitigationFactors = &v
}

// GetCustomText1 returns the CustomText1 field value if set, zero value otherwise.
func (o *EntityRisksPutEntityRisk) GetCustomText1() string {
	if o == nil || IsNil(o.CustomText1) {
		var ret string
		return ret
	}
	return *o.CustomText1
}

// GetCustomText1Ok returns a tuple with the CustomText1 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EntityRisksPutEntityRisk) GetCustomText1Ok() (*string, bool) {
	if o == nil || IsNil(o.CustomText1) {
		return nil, false
	}
	return o.CustomText1, true
}

// HasCustomText1 returns a boolean if a field has been set.
func (o *EntityRisksPutEntityRisk) HasCustomText1() bool {
	if o != nil && !IsNil(o.CustomText1) {
		return true
	}

	return false
}

// SetCustomText1 gets a reference to the given string and assigns it to the CustomText1 field.
func (o *EntityRisksPutEntityRisk) SetCustomText1(v string) {
	o.CustomText1 = &v
}

// GetCustomText2 returns the CustomText2 field value if set, zero value otherwise.
func (o *EntityRisksPutEntityRisk) GetCustomText2() string {
	if o == nil || IsNil(o.CustomText2) {
		var ret string
		return ret
	}
	return *o.CustomText2
}

// GetCustomText2Ok returns a tuple with the CustomText2 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EntityRisksPutEntityRisk) GetCustomText2Ok() (*string, bool) {
	if o == nil || IsNil(o.CustomText2) {
		return nil, false
	}
	return o.CustomText2, true
}

// HasCustomText2 returns a boolean if a field has been set.
func (o *EntityRisksPutEntityRisk) HasCustomText2() bool {
	if o != nil && !IsNil(o.CustomText2) {
		return true
	}

	return false
}

// SetCustomText2 gets a reference to the given string and assigns it to the CustomText2 field.
func (o *EntityRisksPutEntityRisk) SetCustomText2(v string) {
	o.CustomText2 = &v
}

// GetErCustomSelect1OptionId returns the ErCustomSelect1OptionId field value if set, zero value otherwise.
func (o *EntityRisksPutEntityRisk) GetErCustomSelect1OptionId() int32 {
	if o == nil || IsNil(o.ErCustomSelect1OptionId) {
		var ret int32
		return ret
	}
	return *o.ErCustomSelect1OptionId
}

// GetErCustomSelect1OptionIdOk returns a tuple with the ErCustomSelect1OptionId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EntityRisksPutEntityRisk) GetErCustomSelect1OptionIdOk() (*int32, bool) {
	if o == nil || IsNil(o.ErCustomSelect1OptionId) {
		return nil, false
	}
	return o.ErCustomSelect1OptionId, true
}

// HasErCustomSelect1OptionId returns a boolean if a field has been set.
func (o *EntityRisksPutEntityRisk) HasErCustomSelect1OptionId() bool {
	if o != nil && !IsNil(o.ErCustomSelect1OptionId) {
		return true
	}

	return false
}

// SetErCustomSelect1OptionId gets a reference to the given int32 and assigns it to the ErCustomSelect1OptionId field.
func (o *EntityRisksPutEntityRisk) SetErCustomSelect1OptionId(v int32) {
	o.ErCustomSelect1OptionId = &v
}

// GetErCustomSelect2OptionId returns the ErCustomSelect2OptionId field value if set, zero value otherwise.
func (o *EntityRisksPutEntityRisk) GetErCustomSelect2OptionId() int32 {
	if o == nil || IsNil(o.ErCustomSelect2OptionId) {
		var ret int32
		return ret
	}
	return *o.ErCustomSelect2OptionId
}

// GetErCustomSelect2OptionIdOk returns a tuple with the ErCustomSelect2OptionId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EntityRisksPutEntityRisk) GetErCustomSelect2OptionIdOk() (*int32, bool) {
	if o == nil || IsNil(o.ErCustomSelect2OptionId) {
		return nil, false
	}
	return o.ErCustomSelect2OptionId, true
}

// HasErCustomSelect2OptionId returns a boolean if a field has been set.
func (o *EntityRisksPutEntityRisk) HasErCustomSelect2OptionId() bool {
	if o != nil && !IsNil(o.ErCustomSelect2OptionId) {
		return true
	}

	return false
}

// SetErCustomSelect2OptionId gets a reference to the given int32 and assigns it to the ErCustomSelect2OptionId field.
func (o *EntityRisksPutEntityRisk) SetErCustomSelect2OptionId(v int32) {
	o.ErCustomSelect2OptionId = &v
}

// GetRiskCategoryId returns the RiskCategoryId field value if set, zero value otherwise.
func (o *EntityRisksPutEntityRisk) GetRiskCategoryId() int32 {
	if o == nil || IsNil(o.RiskCategoryId) {
		var ret int32
		return ret
	}
	return *o.RiskCategoryId
}

// GetRiskCategoryIdOk returns a tuple with the RiskCategoryId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EntityRisksPutEntityRisk) GetRiskCategoryIdOk() (*int32, bool) {
	if o == nil || IsNil(o.RiskCategoryId) {
		return nil, false
	}
	return o.RiskCategoryId, true
}

// HasRiskCategoryId returns a boolean if a field has been set.
func (o *EntityRisksPutEntityRisk) HasRiskCategoryId() bool {
	if o != nil && !IsNil(o.RiskCategoryId) {
		return true
	}

	return false
}

// SetRiskCategoryId gets a reference to the given int32 and assigns it to the RiskCategoryId field.
func (o *EntityRisksPutEntityRisk) SetRiskCategoryId(v int32) {
	o.RiskCategoryId = &v
}

func (o EntityRisksPutEntityRisk) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o EntityRisksPutEntityRisk) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	toSerialize["entity_id"] = o.EntityId
	toSerialize["risk_id"] = o.RiskId
	toSerialize["status"] = o.Status
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

func (o *EntityRisksPutEntityRisk) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"entity_id",
		"risk_id",
		"status",
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

	varEntityRisksPutEntityRisk := _EntityRisksPutEntityRisk{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varEntityRisksPutEntityRisk)

	if err != nil {
		return err
	}

	*o = EntityRisksPutEntityRisk(varEntityRisksPutEntityRisk)

	return err
}

type NullableEntityRisksPutEntityRisk struct {
	value *EntityRisksPutEntityRisk
	isSet bool
}

func (v NullableEntityRisksPutEntityRisk) Get() *EntityRisksPutEntityRisk {
	return v.value
}

func (v *NullableEntityRisksPutEntityRisk) Set(val *EntityRisksPutEntityRisk) {
	v.value = val
	v.isSet = true
}

func (v NullableEntityRisksPutEntityRisk) IsSet() bool {
	return v.isSet
}

func (v *NullableEntityRisksPutEntityRisk) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEntityRisksPutEntityRisk(val *EntityRisksPutEntityRisk) *NullableEntityRisksPutEntityRisk {
	return &NullableEntityRisksPutEntityRisk{value: val, isSet: true}
}

func (v NullableEntityRisksPutEntityRisk) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEntityRisksPutEntityRisk) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

