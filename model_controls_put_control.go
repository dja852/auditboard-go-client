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

// checks if the ControlsPutControl type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ControlsPutControl{}

// ControlsPutControl struct for ControlsPutControl
type ControlsPutControl struct {
	// Note: This is a Primary Key.<pk/>
	Id *int32 `json:"id,omitempty"`
	CreatedAt *string `json:"created_at,omitempty"`
	UpdatedAt *string `json:"updated_at,omitempty"`
	DeletedAt *string `json:"deleted_at,omitempty"`
	Uid string `json:"uid"`
	Name string `json:"name"`
	// Note: This is a Foreign Key to `subprocesses.id`.<fk table='subprocesses' column='id'/>
	SubprocessId *int32 `json:"subprocess_id,omitempty"`
	Description *string `json:"description,omitempty"`
	IdCode string `json:"id_code"`
	FieldData interface{} `json:"field_data,omitempty"`
	ComplianceControlCount int32 `json:"compliance_control_count"`
	ControlObjective *string `json:"control_objective,omitempty"`
	RiskStatement *string `json:"risk_statement,omitempty"`
	// Note: This is a Foreign Key to `library_control_classifications.id`.<fk table='library_control_classifications' column='id'/>
	LibraryControlClassificationId *int32 `json:"library_control_classification_id,omitempty"`
	// Note: This is a Foreign Key to `library_control_natures.id`.<fk table='library_control_natures' column='id'/>
	LibraryControlNatureId *int32 `json:"library_control_nature_id,omitempty"`
	// Note: This is a Foreign Key to `library_control_types.id`.<fk table='library_control_types' column='id'/>
	LibraryControlTypeId *int32 `json:"library_control_type_id,omitempty"`
	// Note: This is a Foreign Key to `control_categories.id`.<fk table='control_categories' column='id'/>
	ControlCategoryId int32 `json:"control_category_id"`
}

type _ControlsPutControl ControlsPutControl

// NewControlsPutControl instantiates a new ControlsPutControl object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewControlsPutControl(uid string, name string, idCode string, complianceControlCount int32, controlCategoryId int32) *ControlsPutControl {
	this := ControlsPutControl{}
	this.Uid = uid
	this.Name = name
	var description string = ""
	this.Description = &description
	this.IdCode = idCode
	this.ComplianceControlCount = complianceControlCount
	this.ControlCategoryId = controlCategoryId
	return &this
}

// NewControlsPutControlWithDefaults instantiates a new ControlsPutControl object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewControlsPutControlWithDefaults() *ControlsPutControl {
	this := ControlsPutControl{}
	var description string = ""
	this.Description = &description
	var complianceControlCount int32 = 0
	this.ComplianceControlCount = complianceControlCount
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *ControlsPutControl) GetId() int32 {
	if o == nil || IsNil(o.Id) {
		var ret int32
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ControlsPutControl) GetIdOk() (*int32, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *ControlsPutControl) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given int32 and assigns it to the Id field.
func (o *ControlsPutControl) SetId(v int32) {
	o.Id = &v
}

// GetCreatedAt returns the CreatedAt field value if set, zero value otherwise.
func (o *ControlsPutControl) GetCreatedAt() string {
	if o == nil || IsNil(o.CreatedAt) {
		var ret string
		return ret
	}
	return *o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ControlsPutControl) GetCreatedAtOk() (*string, bool) {
	if o == nil || IsNil(o.CreatedAt) {
		return nil, false
	}
	return o.CreatedAt, true
}

// HasCreatedAt returns a boolean if a field has been set.
func (o *ControlsPutControl) HasCreatedAt() bool {
	if o != nil && !IsNil(o.CreatedAt) {
		return true
	}

	return false
}

// SetCreatedAt gets a reference to the given string and assigns it to the CreatedAt field.
func (o *ControlsPutControl) SetCreatedAt(v string) {
	o.CreatedAt = &v
}

// GetUpdatedAt returns the UpdatedAt field value if set, zero value otherwise.
func (o *ControlsPutControl) GetUpdatedAt() string {
	if o == nil || IsNil(o.UpdatedAt) {
		var ret string
		return ret
	}
	return *o.UpdatedAt
}

// GetUpdatedAtOk returns a tuple with the UpdatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ControlsPutControl) GetUpdatedAtOk() (*string, bool) {
	if o == nil || IsNil(o.UpdatedAt) {
		return nil, false
	}
	return o.UpdatedAt, true
}

// HasUpdatedAt returns a boolean if a field has been set.
func (o *ControlsPutControl) HasUpdatedAt() bool {
	if o != nil && !IsNil(o.UpdatedAt) {
		return true
	}

	return false
}

// SetUpdatedAt gets a reference to the given string and assigns it to the UpdatedAt field.
func (o *ControlsPutControl) SetUpdatedAt(v string) {
	o.UpdatedAt = &v
}

// GetDeletedAt returns the DeletedAt field value if set, zero value otherwise.
func (o *ControlsPutControl) GetDeletedAt() string {
	if o == nil || IsNil(o.DeletedAt) {
		var ret string
		return ret
	}
	return *o.DeletedAt
}

// GetDeletedAtOk returns a tuple with the DeletedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ControlsPutControl) GetDeletedAtOk() (*string, bool) {
	if o == nil || IsNil(o.DeletedAt) {
		return nil, false
	}
	return o.DeletedAt, true
}

// HasDeletedAt returns a boolean if a field has been set.
func (o *ControlsPutControl) HasDeletedAt() bool {
	if o != nil && !IsNil(o.DeletedAt) {
		return true
	}

	return false
}

// SetDeletedAt gets a reference to the given string and assigns it to the DeletedAt field.
func (o *ControlsPutControl) SetDeletedAt(v string) {
	o.DeletedAt = &v
}

// GetUid returns the Uid field value
func (o *ControlsPutControl) GetUid() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Uid
}

// GetUidOk returns a tuple with the Uid field value
// and a boolean to check if the value has been set.
func (o *ControlsPutControl) GetUidOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Uid, true
}

// SetUid sets field value
func (o *ControlsPutControl) SetUid(v string) {
	o.Uid = v
}

// GetName returns the Name field value
func (o *ControlsPutControl) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *ControlsPutControl) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *ControlsPutControl) SetName(v string) {
	o.Name = v
}

// GetSubprocessId returns the SubprocessId field value if set, zero value otherwise.
func (o *ControlsPutControl) GetSubprocessId() int32 {
	if o == nil || IsNil(o.SubprocessId) {
		var ret int32
		return ret
	}
	return *o.SubprocessId
}

// GetSubprocessIdOk returns a tuple with the SubprocessId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ControlsPutControl) GetSubprocessIdOk() (*int32, bool) {
	if o == nil || IsNil(o.SubprocessId) {
		return nil, false
	}
	return o.SubprocessId, true
}

// HasSubprocessId returns a boolean if a field has been set.
func (o *ControlsPutControl) HasSubprocessId() bool {
	if o != nil && !IsNil(o.SubprocessId) {
		return true
	}

	return false
}

// SetSubprocessId gets a reference to the given int32 and assigns it to the SubprocessId field.
func (o *ControlsPutControl) SetSubprocessId(v int32) {
	o.SubprocessId = &v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *ControlsPutControl) GetDescription() string {
	if o == nil || IsNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ControlsPutControl) GetDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *ControlsPutControl) HasDescription() bool {
	if o != nil && !IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *ControlsPutControl) SetDescription(v string) {
	o.Description = &v
}

// GetIdCode returns the IdCode field value
func (o *ControlsPutControl) GetIdCode() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.IdCode
}

// GetIdCodeOk returns a tuple with the IdCode field value
// and a boolean to check if the value has been set.
func (o *ControlsPutControl) GetIdCodeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.IdCode, true
}

// SetIdCode sets field value
func (o *ControlsPutControl) SetIdCode(v string) {
	o.IdCode = v
}

// GetFieldData returns the FieldData field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ControlsPutControl) GetFieldData() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.FieldData
}

// GetFieldDataOk returns a tuple with the FieldData field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ControlsPutControl) GetFieldDataOk() (*interface{}, bool) {
	if o == nil || IsNil(o.FieldData) {
		return nil, false
	}
	return &o.FieldData, true
}

// HasFieldData returns a boolean if a field has been set.
func (o *ControlsPutControl) HasFieldData() bool {
	if o != nil && !IsNil(o.FieldData) {
		return true
	}

	return false
}

// SetFieldData gets a reference to the given interface{} and assigns it to the FieldData field.
func (o *ControlsPutControl) SetFieldData(v interface{}) {
	o.FieldData = v
}

// GetComplianceControlCount returns the ComplianceControlCount field value
func (o *ControlsPutControl) GetComplianceControlCount() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.ComplianceControlCount
}

// GetComplianceControlCountOk returns a tuple with the ComplianceControlCount field value
// and a boolean to check if the value has been set.
func (o *ControlsPutControl) GetComplianceControlCountOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ComplianceControlCount, true
}

// SetComplianceControlCount sets field value
func (o *ControlsPutControl) SetComplianceControlCount(v int32) {
	o.ComplianceControlCount = v
}

// GetControlObjective returns the ControlObjective field value if set, zero value otherwise.
func (o *ControlsPutControl) GetControlObjective() string {
	if o == nil || IsNil(o.ControlObjective) {
		var ret string
		return ret
	}
	return *o.ControlObjective
}

// GetControlObjectiveOk returns a tuple with the ControlObjective field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ControlsPutControl) GetControlObjectiveOk() (*string, bool) {
	if o == nil || IsNil(o.ControlObjective) {
		return nil, false
	}
	return o.ControlObjective, true
}

// HasControlObjective returns a boolean if a field has been set.
func (o *ControlsPutControl) HasControlObjective() bool {
	if o != nil && !IsNil(o.ControlObjective) {
		return true
	}

	return false
}

// SetControlObjective gets a reference to the given string and assigns it to the ControlObjective field.
func (o *ControlsPutControl) SetControlObjective(v string) {
	o.ControlObjective = &v
}

// GetRiskStatement returns the RiskStatement field value if set, zero value otherwise.
func (o *ControlsPutControl) GetRiskStatement() string {
	if o == nil || IsNil(o.RiskStatement) {
		var ret string
		return ret
	}
	return *o.RiskStatement
}

// GetRiskStatementOk returns a tuple with the RiskStatement field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ControlsPutControl) GetRiskStatementOk() (*string, bool) {
	if o == nil || IsNil(o.RiskStatement) {
		return nil, false
	}
	return o.RiskStatement, true
}

// HasRiskStatement returns a boolean if a field has been set.
func (o *ControlsPutControl) HasRiskStatement() bool {
	if o != nil && !IsNil(o.RiskStatement) {
		return true
	}

	return false
}

// SetRiskStatement gets a reference to the given string and assigns it to the RiskStatement field.
func (o *ControlsPutControl) SetRiskStatement(v string) {
	o.RiskStatement = &v
}

// GetLibraryControlClassificationId returns the LibraryControlClassificationId field value if set, zero value otherwise.
func (o *ControlsPutControl) GetLibraryControlClassificationId() int32 {
	if o == nil || IsNil(o.LibraryControlClassificationId) {
		var ret int32
		return ret
	}
	return *o.LibraryControlClassificationId
}

// GetLibraryControlClassificationIdOk returns a tuple with the LibraryControlClassificationId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ControlsPutControl) GetLibraryControlClassificationIdOk() (*int32, bool) {
	if o == nil || IsNil(o.LibraryControlClassificationId) {
		return nil, false
	}
	return o.LibraryControlClassificationId, true
}

// HasLibraryControlClassificationId returns a boolean if a field has been set.
func (o *ControlsPutControl) HasLibraryControlClassificationId() bool {
	if o != nil && !IsNil(o.LibraryControlClassificationId) {
		return true
	}

	return false
}

// SetLibraryControlClassificationId gets a reference to the given int32 and assigns it to the LibraryControlClassificationId field.
func (o *ControlsPutControl) SetLibraryControlClassificationId(v int32) {
	o.LibraryControlClassificationId = &v
}

// GetLibraryControlNatureId returns the LibraryControlNatureId field value if set, zero value otherwise.
func (o *ControlsPutControl) GetLibraryControlNatureId() int32 {
	if o == nil || IsNil(o.LibraryControlNatureId) {
		var ret int32
		return ret
	}
	return *o.LibraryControlNatureId
}

// GetLibraryControlNatureIdOk returns a tuple with the LibraryControlNatureId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ControlsPutControl) GetLibraryControlNatureIdOk() (*int32, bool) {
	if o == nil || IsNil(o.LibraryControlNatureId) {
		return nil, false
	}
	return o.LibraryControlNatureId, true
}

// HasLibraryControlNatureId returns a boolean if a field has been set.
func (o *ControlsPutControl) HasLibraryControlNatureId() bool {
	if o != nil && !IsNil(o.LibraryControlNatureId) {
		return true
	}

	return false
}

// SetLibraryControlNatureId gets a reference to the given int32 and assigns it to the LibraryControlNatureId field.
func (o *ControlsPutControl) SetLibraryControlNatureId(v int32) {
	o.LibraryControlNatureId = &v
}

// GetLibraryControlTypeId returns the LibraryControlTypeId field value if set, zero value otherwise.
func (o *ControlsPutControl) GetLibraryControlTypeId() int32 {
	if o == nil || IsNil(o.LibraryControlTypeId) {
		var ret int32
		return ret
	}
	return *o.LibraryControlTypeId
}

// GetLibraryControlTypeIdOk returns a tuple with the LibraryControlTypeId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ControlsPutControl) GetLibraryControlTypeIdOk() (*int32, bool) {
	if o == nil || IsNil(o.LibraryControlTypeId) {
		return nil, false
	}
	return o.LibraryControlTypeId, true
}

// HasLibraryControlTypeId returns a boolean if a field has been set.
func (o *ControlsPutControl) HasLibraryControlTypeId() bool {
	if o != nil && !IsNil(o.LibraryControlTypeId) {
		return true
	}

	return false
}

// SetLibraryControlTypeId gets a reference to the given int32 and assigns it to the LibraryControlTypeId field.
func (o *ControlsPutControl) SetLibraryControlTypeId(v int32) {
	o.LibraryControlTypeId = &v
}

// GetControlCategoryId returns the ControlCategoryId field value
func (o *ControlsPutControl) GetControlCategoryId() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.ControlCategoryId
}

// GetControlCategoryIdOk returns a tuple with the ControlCategoryId field value
// and a boolean to check if the value has been set.
func (o *ControlsPutControl) GetControlCategoryIdOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ControlCategoryId, true
}

// SetControlCategoryId sets field value
func (o *ControlsPutControl) SetControlCategoryId(v int32) {
	o.ControlCategoryId = v
}

func (o ControlsPutControl) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ControlsPutControl) ToMap() (map[string]interface{}, error) {
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
	toSerialize["uid"] = o.Uid
	toSerialize["name"] = o.Name
	if !IsNil(o.SubprocessId) {
		toSerialize["subprocess_id"] = o.SubprocessId
	}
	if !IsNil(o.Description) {
		toSerialize["description"] = o.Description
	}
	toSerialize["id_code"] = o.IdCode
	if o.FieldData != nil {
		toSerialize["field_data"] = o.FieldData
	}
	toSerialize["compliance_control_count"] = o.ComplianceControlCount
	if !IsNil(o.ControlObjective) {
		toSerialize["control_objective"] = o.ControlObjective
	}
	if !IsNil(o.RiskStatement) {
		toSerialize["risk_statement"] = o.RiskStatement
	}
	if !IsNil(o.LibraryControlClassificationId) {
		toSerialize["library_control_classification_id"] = o.LibraryControlClassificationId
	}
	if !IsNil(o.LibraryControlNatureId) {
		toSerialize["library_control_nature_id"] = o.LibraryControlNatureId
	}
	if !IsNil(o.LibraryControlTypeId) {
		toSerialize["library_control_type_id"] = o.LibraryControlTypeId
	}
	toSerialize["control_category_id"] = o.ControlCategoryId
	return toSerialize, nil
}

type NullableControlsPutControl struct {
	value *ControlsPutControl
	isSet bool
}

func (v NullableControlsPutControl) Get() *ControlsPutControl {
	return v.value
}

func (v *NullableControlsPutControl) Set(val *ControlsPutControl) {
	v.value = val
	v.isSet = true
}

func (v NullableControlsPutControl) IsSet() bool {
	return v.isSet
}

func (v *NullableControlsPutControl) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableControlsPutControl(val *ControlsPutControl) *NullableControlsPutControl {
	return &NullableControlsPutControl{value: val, isSet: true}
}

func (v NullableControlsPutControl) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableControlsPutControl) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


