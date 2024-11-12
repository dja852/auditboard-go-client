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

// checks if the AssessmentTemplates type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AssessmentTemplates{}

// AssessmentTemplates struct for AssessmentTemplates
type AssessmentTemplates struct {
	// Note: This is a Primary Key.<pk/>
	Id *int32 `json:"id,omitempty"`
	Name *string `json:"name,omitempty"`
	Status *string `json:"status,omitempty"`
	TemplateJson interface{} `json:"template_json,omitempty"`
	CreatedAt *string `json:"created_at,omitempty"`
	UpdatedAt *string `json:"updated_at,omitempty"`
	DeletedAt *string `json:"deleted_at,omitempty"`
	ProjectOptions interface{} `json:"project_options,omitempty"`
	PowerbiMapping interface{} `json:"powerbi_mapping,omitempty"`
	RiskMapping interface{} `json:"risk_mapping,omitempty"`
	Active *bool `json:"active,omitempty"`
	Type *string `json:"type,omitempty"`
	AuditFrequencyCalculationKey *string `json:"audit_frequency_calculation_key,omitempty"`
	PackagingIsStandard bool `json:"packaging_is_standard"`
	FrequenciesNeedRefresh *bool `json:"frequencies_need_refresh,omitempty"`
	DecimalPrecision int32 `json:"decimal_precision"`
}

type _AssessmentTemplates AssessmentTemplates

// NewAssessmentTemplates instantiates a new AssessmentTemplates object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAssessmentTemplates(packagingIsStandard bool, decimalPrecision int32) *AssessmentTemplates {
	this := AssessmentTemplates{}
	this.PackagingIsStandard = packagingIsStandard
	this.DecimalPrecision = decimalPrecision
	return &this
}

// NewAssessmentTemplatesWithDefaults instantiates a new AssessmentTemplates object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAssessmentTemplatesWithDefaults() *AssessmentTemplates {
	this := AssessmentTemplates{}
	var packagingIsStandard bool = false
	this.PackagingIsStandard = packagingIsStandard
	var decimalPrecision int32 = 4
	this.DecimalPrecision = decimalPrecision
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *AssessmentTemplates) GetId() int32 {
	if o == nil || IsNil(o.Id) {
		var ret int32
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AssessmentTemplates) GetIdOk() (*int32, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *AssessmentTemplates) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given int32 and assigns it to the Id field.
func (o *AssessmentTemplates) SetId(v int32) {
	o.Id = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *AssessmentTemplates) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AssessmentTemplates) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *AssessmentTemplates) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *AssessmentTemplates) SetName(v string) {
	o.Name = &v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *AssessmentTemplates) GetStatus() string {
	if o == nil || IsNil(o.Status) {
		var ret string
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AssessmentTemplates) GetStatusOk() (*string, bool) {
	if o == nil || IsNil(o.Status) {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *AssessmentTemplates) HasStatus() bool {
	if o != nil && !IsNil(o.Status) {
		return true
	}

	return false
}

// SetStatus gets a reference to the given string and assigns it to the Status field.
func (o *AssessmentTemplates) SetStatus(v string) {
	o.Status = &v
}

// GetTemplateJson returns the TemplateJson field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *AssessmentTemplates) GetTemplateJson() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.TemplateJson
}

// GetTemplateJsonOk returns a tuple with the TemplateJson field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *AssessmentTemplates) GetTemplateJsonOk() (*interface{}, bool) {
	if o == nil || IsNil(o.TemplateJson) {
		return nil, false
	}
	return &o.TemplateJson, true
}

// HasTemplateJson returns a boolean if a field has been set.
func (o *AssessmentTemplates) HasTemplateJson() bool {
	if o != nil && !IsNil(o.TemplateJson) {
		return true
	}

	return false
}

// SetTemplateJson gets a reference to the given interface{} and assigns it to the TemplateJson field.
func (o *AssessmentTemplates) SetTemplateJson(v interface{}) {
	o.TemplateJson = v
}

// GetCreatedAt returns the CreatedAt field value if set, zero value otherwise.
func (o *AssessmentTemplates) GetCreatedAt() string {
	if o == nil || IsNil(o.CreatedAt) {
		var ret string
		return ret
	}
	return *o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AssessmentTemplates) GetCreatedAtOk() (*string, bool) {
	if o == nil || IsNil(o.CreatedAt) {
		return nil, false
	}
	return o.CreatedAt, true
}

// HasCreatedAt returns a boolean if a field has been set.
func (o *AssessmentTemplates) HasCreatedAt() bool {
	if o != nil && !IsNil(o.CreatedAt) {
		return true
	}

	return false
}

// SetCreatedAt gets a reference to the given string and assigns it to the CreatedAt field.
func (o *AssessmentTemplates) SetCreatedAt(v string) {
	o.CreatedAt = &v
}

// GetUpdatedAt returns the UpdatedAt field value if set, zero value otherwise.
func (o *AssessmentTemplates) GetUpdatedAt() string {
	if o == nil || IsNil(o.UpdatedAt) {
		var ret string
		return ret
	}
	return *o.UpdatedAt
}

// GetUpdatedAtOk returns a tuple with the UpdatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AssessmentTemplates) GetUpdatedAtOk() (*string, bool) {
	if o == nil || IsNil(o.UpdatedAt) {
		return nil, false
	}
	return o.UpdatedAt, true
}

// HasUpdatedAt returns a boolean if a field has been set.
func (o *AssessmentTemplates) HasUpdatedAt() bool {
	if o != nil && !IsNil(o.UpdatedAt) {
		return true
	}

	return false
}

// SetUpdatedAt gets a reference to the given string and assigns it to the UpdatedAt field.
func (o *AssessmentTemplates) SetUpdatedAt(v string) {
	o.UpdatedAt = &v
}

// GetDeletedAt returns the DeletedAt field value if set, zero value otherwise.
func (o *AssessmentTemplates) GetDeletedAt() string {
	if o == nil || IsNil(o.DeletedAt) {
		var ret string
		return ret
	}
	return *o.DeletedAt
}

// GetDeletedAtOk returns a tuple with the DeletedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AssessmentTemplates) GetDeletedAtOk() (*string, bool) {
	if o == nil || IsNil(o.DeletedAt) {
		return nil, false
	}
	return o.DeletedAt, true
}

// HasDeletedAt returns a boolean if a field has been set.
func (o *AssessmentTemplates) HasDeletedAt() bool {
	if o != nil && !IsNil(o.DeletedAt) {
		return true
	}

	return false
}

// SetDeletedAt gets a reference to the given string and assigns it to the DeletedAt field.
func (o *AssessmentTemplates) SetDeletedAt(v string) {
	o.DeletedAt = &v
}

// GetProjectOptions returns the ProjectOptions field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *AssessmentTemplates) GetProjectOptions() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.ProjectOptions
}

// GetProjectOptionsOk returns a tuple with the ProjectOptions field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *AssessmentTemplates) GetProjectOptionsOk() (*interface{}, bool) {
	if o == nil || IsNil(o.ProjectOptions) {
		return nil, false
	}
	return &o.ProjectOptions, true
}

// HasProjectOptions returns a boolean if a field has been set.
func (o *AssessmentTemplates) HasProjectOptions() bool {
	if o != nil && !IsNil(o.ProjectOptions) {
		return true
	}

	return false
}

// SetProjectOptions gets a reference to the given interface{} and assigns it to the ProjectOptions field.
func (o *AssessmentTemplates) SetProjectOptions(v interface{}) {
	o.ProjectOptions = v
}

// GetPowerbiMapping returns the PowerbiMapping field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *AssessmentTemplates) GetPowerbiMapping() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.PowerbiMapping
}

// GetPowerbiMappingOk returns a tuple with the PowerbiMapping field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *AssessmentTemplates) GetPowerbiMappingOk() (*interface{}, bool) {
	if o == nil || IsNil(o.PowerbiMapping) {
		return nil, false
	}
	return &o.PowerbiMapping, true
}

// HasPowerbiMapping returns a boolean if a field has been set.
func (o *AssessmentTemplates) HasPowerbiMapping() bool {
	if o != nil && !IsNil(o.PowerbiMapping) {
		return true
	}

	return false
}

// SetPowerbiMapping gets a reference to the given interface{} and assigns it to the PowerbiMapping field.
func (o *AssessmentTemplates) SetPowerbiMapping(v interface{}) {
	o.PowerbiMapping = v
}

// GetRiskMapping returns the RiskMapping field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *AssessmentTemplates) GetRiskMapping() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.RiskMapping
}

// GetRiskMappingOk returns a tuple with the RiskMapping field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *AssessmentTemplates) GetRiskMappingOk() (*interface{}, bool) {
	if o == nil || IsNil(o.RiskMapping) {
		return nil, false
	}
	return &o.RiskMapping, true
}

// HasRiskMapping returns a boolean if a field has been set.
func (o *AssessmentTemplates) HasRiskMapping() bool {
	if o != nil && !IsNil(o.RiskMapping) {
		return true
	}

	return false
}

// SetRiskMapping gets a reference to the given interface{} and assigns it to the RiskMapping field.
func (o *AssessmentTemplates) SetRiskMapping(v interface{}) {
	o.RiskMapping = v
}

// GetActive returns the Active field value if set, zero value otherwise.
func (o *AssessmentTemplates) GetActive() bool {
	if o == nil || IsNil(o.Active) {
		var ret bool
		return ret
	}
	return *o.Active
}

// GetActiveOk returns a tuple with the Active field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AssessmentTemplates) GetActiveOk() (*bool, bool) {
	if o == nil || IsNil(o.Active) {
		return nil, false
	}
	return o.Active, true
}

// HasActive returns a boolean if a field has been set.
func (o *AssessmentTemplates) HasActive() bool {
	if o != nil && !IsNil(o.Active) {
		return true
	}

	return false
}

// SetActive gets a reference to the given bool and assigns it to the Active field.
func (o *AssessmentTemplates) SetActive(v bool) {
	o.Active = &v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *AssessmentTemplates) GetType() string {
	if o == nil || IsNil(o.Type) {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AssessmentTemplates) GetTypeOk() (*string, bool) {
	if o == nil || IsNil(o.Type) {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *AssessmentTemplates) HasType() bool {
	if o != nil && !IsNil(o.Type) {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *AssessmentTemplates) SetType(v string) {
	o.Type = &v
}

// GetAuditFrequencyCalculationKey returns the AuditFrequencyCalculationKey field value if set, zero value otherwise.
func (o *AssessmentTemplates) GetAuditFrequencyCalculationKey() string {
	if o == nil || IsNil(o.AuditFrequencyCalculationKey) {
		var ret string
		return ret
	}
	return *o.AuditFrequencyCalculationKey
}

// GetAuditFrequencyCalculationKeyOk returns a tuple with the AuditFrequencyCalculationKey field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AssessmentTemplates) GetAuditFrequencyCalculationKeyOk() (*string, bool) {
	if o == nil || IsNil(o.AuditFrequencyCalculationKey) {
		return nil, false
	}
	return o.AuditFrequencyCalculationKey, true
}

// HasAuditFrequencyCalculationKey returns a boolean if a field has been set.
func (o *AssessmentTemplates) HasAuditFrequencyCalculationKey() bool {
	if o != nil && !IsNil(o.AuditFrequencyCalculationKey) {
		return true
	}

	return false
}

// SetAuditFrequencyCalculationKey gets a reference to the given string and assigns it to the AuditFrequencyCalculationKey field.
func (o *AssessmentTemplates) SetAuditFrequencyCalculationKey(v string) {
	o.AuditFrequencyCalculationKey = &v
}

// GetPackagingIsStandard returns the PackagingIsStandard field value
func (o *AssessmentTemplates) GetPackagingIsStandard() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.PackagingIsStandard
}

// GetPackagingIsStandardOk returns a tuple with the PackagingIsStandard field value
// and a boolean to check if the value has been set.
func (o *AssessmentTemplates) GetPackagingIsStandardOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.PackagingIsStandard, true
}

// SetPackagingIsStandard sets field value
func (o *AssessmentTemplates) SetPackagingIsStandard(v bool) {
	o.PackagingIsStandard = v
}

// GetFrequenciesNeedRefresh returns the FrequenciesNeedRefresh field value if set, zero value otherwise.
func (o *AssessmentTemplates) GetFrequenciesNeedRefresh() bool {
	if o == nil || IsNil(o.FrequenciesNeedRefresh) {
		var ret bool
		return ret
	}
	return *o.FrequenciesNeedRefresh
}

// GetFrequenciesNeedRefreshOk returns a tuple with the FrequenciesNeedRefresh field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AssessmentTemplates) GetFrequenciesNeedRefreshOk() (*bool, bool) {
	if o == nil || IsNil(o.FrequenciesNeedRefresh) {
		return nil, false
	}
	return o.FrequenciesNeedRefresh, true
}

// HasFrequenciesNeedRefresh returns a boolean if a field has been set.
func (o *AssessmentTemplates) HasFrequenciesNeedRefresh() bool {
	if o != nil && !IsNil(o.FrequenciesNeedRefresh) {
		return true
	}

	return false
}

// SetFrequenciesNeedRefresh gets a reference to the given bool and assigns it to the FrequenciesNeedRefresh field.
func (o *AssessmentTemplates) SetFrequenciesNeedRefresh(v bool) {
	o.FrequenciesNeedRefresh = &v
}

// GetDecimalPrecision returns the DecimalPrecision field value
func (o *AssessmentTemplates) GetDecimalPrecision() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.DecimalPrecision
}

// GetDecimalPrecisionOk returns a tuple with the DecimalPrecision field value
// and a boolean to check if the value has been set.
func (o *AssessmentTemplates) GetDecimalPrecisionOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.DecimalPrecision, true
}

// SetDecimalPrecision sets field value
func (o *AssessmentTemplates) SetDecimalPrecision(v int32) {
	o.DecimalPrecision = v
}

func (o AssessmentTemplates) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AssessmentTemplates) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !IsNil(o.Status) {
		toSerialize["status"] = o.Status
	}
	if o.TemplateJson != nil {
		toSerialize["template_json"] = o.TemplateJson
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
	if o.ProjectOptions != nil {
		toSerialize["project_options"] = o.ProjectOptions
	}
	if o.PowerbiMapping != nil {
		toSerialize["powerbi_mapping"] = o.PowerbiMapping
	}
	if o.RiskMapping != nil {
		toSerialize["risk_mapping"] = o.RiskMapping
	}
	if !IsNil(o.Active) {
		toSerialize["active"] = o.Active
	}
	if !IsNil(o.Type) {
		toSerialize["type"] = o.Type
	}
	if !IsNil(o.AuditFrequencyCalculationKey) {
		toSerialize["audit_frequency_calculation_key"] = o.AuditFrequencyCalculationKey
	}
	toSerialize["packaging_is_standard"] = o.PackagingIsStandard
	if !IsNil(o.FrequenciesNeedRefresh) {
		toSerialize["frequencies_need_refresh"] = o.FrequenciesNeedRefresh
	}
	toSerialize["decimal_precision"] = o.DecimalPrecision
	return toSerialize, nil
}

type NullableAssessmentTemplates struct {
	value *AssessmentTemplates
	isSet bool
}

func (v NullableAssessmentTemplates) Get() *AssessmentTemplates {
	return v.value
}

func (v *NullableAssessmentTemplates) Set(val *AssessmentTemplates) {
	v.value = val
	v.isSet = true
}

func (v NullableAssessmentTemplates) IsSet() bool {
	return v.isSet
}

func (v *NullableAssessmentTemplates) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAssessmentTemplates(val *AssessmentTemplates) *NullableAssessmentTemplates {
	return &NullableAssessmentTemplates{value: val, isSet: true}
}

func (v NullableAssessmentTemplates) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAssessmentTemplates) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


