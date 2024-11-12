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

// checks if the KeyRiskIndicatorValue type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &KeyRiskIndicatorValue{}

// KeyRiskIndicatorValue struct for KeyRiskIndicatorValue
type KeyRiskIndicatorValue struct {
	HistoricalDate *string `json:"historical_date,omitempty"`
	Id int64 `json:"id"`
	KeyRiskIndicatorId *int64 `json:"key_risk_indicator_id,omitempty"`
	KriStatus *string `json:"kri_status,omitempty"`
	KriFormula *string `json:"kri_formula,omitempty"`
	Value *string `json:"value,omitempty"`
	Comments *string `json:"comments,omitempty"`
	UpdatedAt *string `json:"updated_at,omitempty"`
	CreatedAt *string `json:"created_at,omitempty"`
	DeletedAt *string `json:"deleted_at,omitempty"`
	UpdatedByUserId interface{} `json:"updated_by_user_id,omitempty"`
	NormalizedValue *string `json:"normalized_value,omitempty"`
}

type _KeyRiskIndicatorValue KeyRiskIndicatorValue

// NewKeyRiskIndicatorValue instantiates a new KeyRiskIndicatorValue object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewKeyRiskIndicatorValue(id int64) *KeyRiskIndicatorValue {
	this := KeyRiskIndicatorValue{}
	this.Id = id
	return &this
}

// NewKeyRiskIndicatorValueWithDefaults instantiates a new KeyRiskIndicatorValue object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewKeyRiskIndicatorValueWithDefaults() *KeyRiskIndicatorValue {
	this := KeyRiskIndicatorValue{}
	return &this
}

// GetHistoricalDate returns the HistoricalDate field value if set, zero value otherwise.
func (o *KeyRiskIndicatorValue) GetHistoricalDate() string {
	if o == nil || IsNil(o.HistoricalDate) {
		var ret string
		return ret
	}
	return *o.HistoricalDate
}

// GetHistoricalDateOk returns a tuple with the HistoricalDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KeyRiskIndicatorValue) GetHistoricalDateOk() (*string, bool) {
	if o == nil || IsNil(o.HistoricalDate) {
		return nil, false
	}
	return o.HistoricalDate, true
}

// HasHistoricalDate returns a boolean if a field has been set.
func (o *KeyRiskIndicatorValue) HasHistoricalDate() bool {
	if o != nil && !IsNil(o.HistoricalDate) {
		return true
	}

	return false
}

// SetHistoricalDate gets a reference to the given string and assigns it to the HistoricalDate field.
func (o *KeyRiskIndicatorValue) SetHistoricalDate(v string) {
	o.HistoricalDate = &v
}

// GetId returns the Id field value
func (o *KeyRiskIndicatorValue) GetId() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *KeyRiskIndicatorValue) GetIdOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *KeyRiskIndicatorValue) SetId(v int64) {
	o.Id = v
}

// GetKeyRiskIndicatorId returns the KeyRiskIndicatorId field value if set, zero value otherwise.
func (o *KeyRiskIndicatorValue) GetKeyRiskIndicatorId() int64 {
	if o == nil || IsNil(o.KeyRiskIndicatorId) {
		var ret int64
		return ret
	}
	return *o.KeyRiskIndicatorId
}

// GetKeyRiskIndicatorIdOk returns a tuple with the KeyRiskIndicatorId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KeyRiskIndicatorValue) GetKeyRiskIndicatorIdOk() (*int64, bool) {
	if o == nil || IsNil(o.KeyRiskIndicatorId) {
		return nil, false
	}
	return o.KeyRiskIndicatorId, true
}

// HasKeyRiskIndicatorId returns a boolean if a field has been set.
func (o *KeyRiskIndicatorValue) HasKeyRiskIndicatorId() bool {
	if o != nil && !IsNil(o.KeyRiskIndicatorId) {
		return true
	}

	return false
}

// SetKeyRiskIndicatorId gets a reference to the given int64 and assigns it to the KeyRiskIndicatorId field.
func (o *KeyRiskIndicatorValue) SetKeyRiskIndicatorId(v int64) {
	o.KeyRiskIndicatorId = &v
}

// GetKriStatus returns the KriStatus field value if set, zero value otherwise.
func (o *KeyRiskIndicatorValue) GetKriStatus() string {
	if o == nil || IsNil(o.KriStatus) {
		var ret string
		return ret
	}
	return *o.KriStatus
}

// GetKriStatusOk returns a tuple with the KriStatus field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KeyRiskIndicatorValue) GetKriStatusOk() (*string, bool) {
	if o == nil || IsNil(o.KriStatus) {
		return nil, false
	}
	return o.KriStatus, true
}

// HasKriStatus returns a boolean if a field has been set.
func (o *KeyRiskIndicatorValue) HasKriStatus() bool {
	if o != nil && !IsNil(o.KriStatus) {
		return true
	}

	return false
}

// SetKriStatus gets a reference to the given string and assigns it to the KriStatus field.
func (o *KeyRiskIndicatorValue) SetKriStatus(v string) {
	o.KriStatus = &v
}

// GetKriFormula returns the KriFormula field value if set, zero value otherwise.
func (o *KeyRiskIndicatorValue) GetKriFormula() string {
	if o == nil || IsNil(o.KriFormula) {
		var ret string
		return ret
	}
	return *o.KriFormula
}

// GetKriFormulaOk returns a tuple with the KriFormula field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KeyRiskIndicatorValue) GetKriFormulaOk() (*string, bool) {
	if o == nil || IsNil(o.KriFormula) {
		return nil, false
	}
	return o.KriFormula, true
}

// HasKriFormula returns a boolean if a field has been set.
func (o *KeyRiskIndicatorValue) HasKriFormula() bool {
	if o != nil && !IsNil(o.KriFormula) {
		return true
	}

	return false
}

// SetKriFormula gets a reference to the given string and assigns it to the KriFormula field.
func (o *KeyRiskIndicatorValue) SetKriFormula(v string) {
	o.KriFormula = &v
}

// GetValue returns the Value field value if set, zero value otherwise.
func (o *KeyRiskIndicatorValue) GetValue() string {
	if o == nil || IsNil(o.Value) {
		var ret string
		return ret
	}
	return *o.Value
}

// GetValueOk returns a tuple with the Value field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KeyRiskIndicatorValue) GetValueOk() (*string, bool) {
	if o == nil || IsNil(o.Value) {
		return nil, false
	}
	return o.Value, true
}

// HasValue returns a boolean if a field has been set.
func (o *KeyRiskIndicatorValue) HasValue() bool {
	if o != nil && !IsNil(o.Value) {
		return true
	}

	return false
}

// SetValue gets a reference to the given string and assigns it to the Value field.
func (o *KeyRiskIndicatorValue) SetValue(v string) {
	o.Value = &v
}

// GetComments returns the Comments field value if set, zero value otherwise.
func (o *KeyRiskIndicatorValue) GetComments() string {
	if o == nil || IsNil(o.Comments) {
		var ret string
		return ret
	}
	return *o.Comments
}

// GetCommentsOk returns a tuple with the Comments field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KeyRiskIndicatorValue) GetCommentsOk() (*string, bool) {
	if o == nil || IsNil(o.Comments) {
		return nil, false
	}
	return o.Comments, true
}

// HasComments returns a boolean if a field has been set.
func (o *KeyRiskIndicatorValue) HasComments() bool {
	if o != nil && !IsNil(o.Comments) {
		return true
	}

	return false
}

// SetComments gets a reference to the given string and assigns it to the Comments field.
func (o *KeyRiskIndicatorValue) SetComments(v string) {
	o.Comments = &v
}

// GetUpdatedAt returns the UpdatedAt field value if set, zero value otherwise.
func (o *KeyRiskIndicatorValue) GetUpdatedAt() string {
	if o == nil || IsNil(o.UpdatedAt) {
		var ret string
		return ret
	}
	return *o.UpdatedAt
}

// GetUpdatedAtOk returns a tuple with the UpdatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KeyRiskIndicatorValue) GetUpdatedAtOk() (*string, bool) {
	if o == nil || IsNil(o.UpdatedAt) {
		return nil, false
	}
	return o.UpdatedAt, true
}

// HasUpdatedAt returns a boolean if a field has been set.
func (o *KeyRiskIndicatorValue) HasUpdatedAt() bool {
	if o != nil && !IsNil(o.UpdatedAt) {
		return true
	}

	return false
}

// SetUpdatedAt gets a reference to the given string and assigns it to the UpdatedAt field.
func (o *KeyRiskIndicatorValue) SetUpdatedAt(v string) {
	o.UpdatedAt = &v
}

// GetCreatedAt returns the CreatedAt field value if set, zero value otherwise.
func (o *KeyRiskIndicatorValue) GetCreatedAt() string {
	if o == nil || IsNil(o.CreatedAt) {
		var ret string
		return ret
	}
	return *o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KeyRiskIndicatorValue) GetCreatedAtOk() (*string, bool) {
	if o == nil || IsNil(o.CreatedAt) {
		return nil, false
	}
	return o.CreatedAt, true
}

// HasCreatedAt returns a boolean if a field has been set.
func (o *KeyRiskIndicatorValue) HasCreatedAt() bool {
	if o != nil && !IsNil(o.CreatedAt) {
		return true
	}

	return false
}

// SetCreatedAt gets a reference to the given string and assigns it to the CreatedAt field.
func (o *KeyRiskIndicatorValue) SetCreatedAt(v string) {
	o.CreatedAt = &v
}

// GetDeletedAt returns the DeletedAt field value if set, zero value otherwise.
func (o *KeyRiskIndicatorValue) GetDeletedAt() string {
	if o == nil || IsNil(o.DeletedAt) {
		var ret string
		return ret
	}
	return *o.DeletedAt
}

// GetDeletedAtOk returns a tuple with the DeletedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KeyRiskIndicatorValue) GetDeletedAtOk() (*string, bool) {
	if o == nil || IsNil(o.DeletedAt) {
		return nil, false
	}
	return o.DeletedAt, true
}

// HasDeletedAt returns a boolean if a field has been set.
func (o *KeyRiskIndicatorValue) HasDeletedAt() bool {
	if o != nil && !IsNil(o.DeletedAt) {
		return true
	}

	return false
}

// SetDeletedAt gets a reference to the given string and assigns it to the DeletedAt field.
func (o *KeyRiskIndicatorValue) SetDeletedAt(v string) {
	o.DeletedAt = &v
}

// GetUpdatedByUserId returns the UpdatedByUserId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *KeyRiskIndicatorValue) GetUpdatedByUserId() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.UpdatedByUserId
}

// GetUpdatedByUserIdOk returns a tuple with the UpdatedByUserId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *KeyRiskIndicatorValue) GetUpdatedByUserIdOk() (*interface{}, bool) {
	if o == nil || IsNil(o.UpdatedByUserId) {
		return nil, false
	}
	return &o.UpdatedByUserId, true
}

// HasUpdatedByUserId returns a boolean if a field has been set.
func (o *KeyRiskIndicatorValue) HasUpdatedByUserId() bool {
	if o != nil && !IsNil(o.UpdatedByUserId) {
		return true
	}

	return false
}

// SetUpdatedByUserId gets a reference to the given interface{} and assigns it to the UpdatedByUserId field.
func (o *KeyRiskIndicatorValue) SetUpdatedByUserId(v interface{}) {
	o.UpdatedByUserId = v
}

// GetNormalizedValue returns the NormalizedValue field value if set, zero value otherwise.
func (o *KeyRiskIndicatorValue) GetNormalizedValue() string {
	if o == nil || IsNil(o.NormalizedValue) {
		var ret string
		return ret
	}
	return *o.NormalizedValue
}

// GetNormalizedValueOk returns a tuple with the NormalizedValue field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KeyRiskIndicatorValue) GetNormalizedValueOk() (*string, bool) {
	if o == nil || IsNil(o.NormalizedValue) {
		return nil, false
	}
	return o.NormalizedValue, true
}

// HasNormalizedValue returns a boolean if a field has been set.
func (o *KeyRiskIndicatorValue) HasNormalizedValue() bool {
	if o != nil && !IsNil(o.NormalizedValue) {
		return true
	}

	return false
}

// SetNormalizedValue gets a reference to the given string and assigns it to the NormalizedValue field.
func (o *KeyRiskIndicatorValue) SetNormalizedValue(v string) {
	o.NormalizedValue = &v
}

func (o KeyRiskIndicatorValue) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o KeyRiskIndicatorValue) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.HistoricalDate) {
		toSerialize["historical_date"] = o.HistoricalDate
	}
	toSerialize["id"] = o.Id
	if !IsNil(o.KeyRiskIndicatorId) {
		toSerialize["key_risk_indicator_id"] = o.KeyRiskIndicatorId
	}
	if !IsNil(o.KriStatus) {
		toSerialize["kri_status"] = o.KriStatus
	}
	if !IsNil(o.KriFormula) {
		toSerialize["kri_formula"] = o.KriFormula
	}
	if !IsNil(o.Value) {
		toSerialize["value"] = o.Value
	}
	if !IsNil(o.Comments) {
		toSerialize["comments"] = o.Comments
	}
	if !IsNil(o.UpdatedAt) {
		toSerialize["updated_at"] = o.UpdatedAt
	}
	if !IsNil(o.CreatedAt) {
		toSerialize["created_at"] = o.CreatedAt
	}
	if !IsNil(o.DeletedAt) {
		toSerialize["deleted_at"] = o.DeletedAt
	}
	if o.UpdatedByUserId != nil {
		toSerialize["updated_by_user_id"] = o.UpdatedByUserId
	}
	if !IsNil(o.NormalizedValue) {
		toSerialize["normalized_value"] = o.NormalizedValue
	}
	return toSerialize, nil
}

type NullableKeyRiskIndicatorValue struct {
	value *KeyRiskIndicatorValue
	isSet bool
}

func (v NullableKeyRiskIndicatorValue) Get() *KeyRiskIndicatorValue {
	return v.value
}

func (v *NullableKeyRiskIndicatorValue) Set(val *KeyRiskIndicatorValue) {
	v.value = val
	v.isSet = true
}

func (v NullableKeyRiskIndicatorValue) IsSet() bool {
	return v.isSet
}

func (v *NullableKeyRiskIndicatorValue) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableKeyRiskIndicatorValue(val *KeyRiskIndicatorValue) *NullableKeyRiskIndicatorValue {
	return &NullableKeyRiskIndicatorValue{value: val, isSet: true}
}

func (v NullableKeyRiskIndicatorValue) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableKeyRiskIndicatorValue) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


