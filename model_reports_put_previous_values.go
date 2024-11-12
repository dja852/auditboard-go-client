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

// checks if the ReportsPutPreviousValues type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ReportsPutPreviousValues{}

// ReportsPutPreviousValues struct for ReportsPutPreviousValues
type ReportsPutPreviousValues struct {
	// Note: This is a Primary Key.<pk/>
	Id *int32 `json:"id,omitempty"`
	CreatedAt *string `json:"created_at,omitempty"`
	UpdatedAt *string `json:"updated_at,omitempty"`
	DeletedAt *string `json:"deleted_at,omitempty"`
	Name *string `json:"name,omitempty"`
	QueryFields *string `json:"query_fields,omitempty"`
	QueryFilters *string `json:"query_filters,omitempty"`
	ShareToken *string `json:"share_token,omitempty"`
	Type *string `json:"type,omitempty"`
	FieldsGrouping *string `json:"fields_grouping,omitempty"`
	CanonicalFields interface{} `json:"canonical_fields,omitempty"`
	// Note: This is a Foreign Key to `users.id`.<fk table='users' column='id'/>
	ShareLinkUserId *int32 `json:"share_link_user_id,omitempty"`
	ReportOptions interface{} `json:"report_options,omitempty"`
}

// NewReportsPutPreviousValues instantiates a new ReportsPutPreviousValues object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewReportsPutPreviousValues() *ReportsPutPreviousValues {
	this := ReportsPutPreviousValues{}
	var type_ string = "Control"
	this.Type = &type_
	return &this
}

// NewReportsPutPreviousValuesWithDefaults instantiates a new ReportsPutPreviousValues object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewReportsPutPreviousValuesWithDefaults() *ReportsPutPreviousValues {
	this := ReportsPutPreviousValues{}
	var type_ string = "Control"
	this.Type = &type_
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *ReportsPutPreviousValues) GetId() int32 {
	if o == nil || IsNil(o.Id) {
		var ret int32
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ReportsPutPreviousValues) GetIdOk() (*int32, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *ReportsPutPreviousValues) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given int32 and assigns it to the Id field.
func (o *ReportsPutPreviousValues) SetId(v int32) {
	o.Id = &v
}

// GetCreatedAt returns the CreatedAt field value if set, zero value otherwise.
func (o *ReportsPutPreviousValues) GetCreatedAt() string {
	if o == nil || IsNil(o.CreatedAt) {
		var ret string
		return ret
	}
	return *o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ReportsPutPreviousValues) GetCreatedAtOk() (*string, bool) {
	if o == nil || IsNil(o.CreatedAt) {
		return nil, false
	}
	return o.CreatedAt, true
}

// HasCreatedAt returns a boolean if a field has been set.
func (o *ReportsPutPreviousValues) HasCreatedAt() bool {
	if o != nil && !IsNil(o.CreatedAt) {
		return true
	}

	return false
}

// SetCreatedAt gets a reference to the given string and assigns it to the CreatedAt field.
func (o *ReportsPutPreviousValues) SetCreatedAt(v string) {
	o.CreatedAt = &v
}

// GetUpdatedAt returns the UpdatedAt field value if set, zero value otherwise.
func (o *ReportsPutPreviousValues) GetUpdatedAt() string {
	if o == nil || IsNil(o.UpdatedAt) {
		var ret string
		return ret
	}
	return *o.UpdatedAt
}

// GetUpdatedAtOk returns a tuple with the UpdatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ReportsPutPreviousValues) GetUpdatedAtOk() (*string, bool) {
	if o == nil || IsNil(o.UpdatedAt) {
		return nil, false
	}
	return o.UpdatedAt, true
}

// HasUpdatedAt returns a boolean if a field has been set.
func (o *ReportsPutPreviousValues) HasUpdatedAt() bool {
	if o != nil && !IsNil(o.UpdatedAt) {
		return true
	}

	return false
}

// SetUpdatedAt gets a reference to the given string and assigns it to the UpdatedAt field.
func (o *ReportsPutPreviousValues) SetUpdatedAt(v string) {
	o.UpdatedAt = &v
}

// GetDeletedAt returns the DeletedAt field value if set, zero value otherwise.
func (o *ReportsPutPreviousValues) GetDeletedAt() string {
	if o == nil || IsNil(o.DeletedAt) {
		var ret string
		return ret
	}
	return *o.DeletedAt
}

// GetDeletedAtOk returns a tuple with the DeletedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ReportsPutPreviousValues) GetDeletedAtOk() (*string, bool) {
	if o == nil || IsNil(o.DeletedAt) {
		return nil, false
	}
	return o.DeletedAt, true
}

// HasDeletedAt returns a boolean if a field has been set.
func (o *ReportsPutPreviousValues) HasDeletedAt() bool {
	if o != nil && !IsNil(o.DeletedAt) {
		return true
	}

	return false
}

// SetDeletedAt gets a reference to the given string and assigns it to the DeletedAt field.
func (o *ReportsPutPreviousValues) SetDeletedAt(v string) {
	o.DeletedAt = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *ReportsPutPreviousValues) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ReportsPutPreviousValues) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *ReportsPutPreviousValues) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *ReportsPutPreviousValues) SetName(v string) {
	o.Name = &v
}

// GetQueryFields returns the QueryFields field value if set, zero value otherwise.
func (o *ReportsPutPreviousValues) GetQueryFields() string {
	if o == nil || IsNil(o.QueryFields) {
		var ret string
		return ret
	}
	return *o.QueryFields
}

// GetQueryFieldsOk returns a tuple with the QueryFields field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ReportsPutPreviousValues) GetQueryFieldsOk() (*string, bool) {
	if o == nil || IsNil(o.QueryFields) {
		return nil, false
	}
	return o.QueryFields, true
}

// HasQueryFields returns a boolean if a field has been set.
func (o *ReportsPutPreviousValues) HasQueryFields() bool {
	if o != nil && !IsNil(o.QueryFields) {
		return true
	}

	return false
}

// SetQueryFields gets a reference to the given string and assigns it to the QueryFields field.
func (o *ReportsPutPreviousValues) SetQueryFields(v string) {
	o.QueryFields = &v
}

// GetQueryFilters returns the QueryFilters field value if set, zero value otherwise.
func (o *ReportsPutPreviousValues) GetQueryFilters() string {
	if o == nil || IsNil(o.QueryFilters) {
		var ret string
		return ret
	}
	return *o.QueryFilters
}

// GetQueryFiltersOk returns a tuple with the QueryFilters field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ReportsPutPreviousValues) GetQueryFiltersOk() (*string, bool) {
	if o == nil || IsNil(o.QueryFilters) {
		return nil, false
	}
	return o.QueryFilters, true
}

// HasQueryFilters returns a boolean if a field has been set.
func (o *ReportsPutPreviousValues) HasQueryFilters() bool {
	if o != nil && !IsNil(o.QueryFilters) {
		return true
	}

	return false
}

// SetQueryFilters gets a reference to the given string and assigns it to the QueryFilters field.
func (o *ReportsPutPreviousValues) SetQueryFilters(v string) {
	o.QueryFilters = &v
}

// GetShareToken returns the ShareToken field value if set, zero value otherwise.
func (o *ReportsPutPreviousValues) GetShareToken() string {
	if o == nil || IsNil(o.ShareToken) {
		var ret string
		return ret
	}
	return *o.ShareToken
}

// GetShareTokenOk returns a tuple with the ShareToken field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ReportsPutPreviousValues) GetShareTokenOk() (*string, bool) {
	if o == nil || IsNil(o.ShareToken) {
		return nil, false
	}
	return o.ShareToken, true
}

// HasShareToken returns a boolean if a field has been set.
func (o *ReportsPutPreviousValues) HasShareToken() bool {
	if o != nil && !IsNil(o.ShareToken) {
		return true
	}

	return false
}

// SetShareToken gets a reference to the given string and assigns it to the ShareToken field.
func (o *ReportsPutPreviousValues) SetShareToken(v string) {
	o.ShareToken = &v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *ReportsPutPreviousValues) GetType() string {
	if o == nil || IsNil(o.Type) {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ReportsPutPreviousValues) GetTypeOk() (*string, bool) {
	if o == nil || IsNil(o.Type) {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *ReportsPutPreviousValues) HasType() bool {
	if o != nil && !IsNil(o.Type) {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *ReportsPutPreviousValues) SetType(v string) {
	o.Type = &v
}

// GetFieldsGrouping returns the FieldsGrouping field value if set, zero value otherwise.
func (o *ReportsPutPreviousValues) GetFieldsGrouping() string {
	if o == nil || IsNil(o.FieldsGrouping) {
		var ret string
		return ret
	}
	return *o.FieldsGrouping
}

// GetFieldsGroupingOk returns a tuple with the FieldsGrouping field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ReportsPutPreviousValues) GetFieldsGroupingOk() (*string, bool) {
	if o == nil || IsNil(o.FieldsGrouping) {
		return nil, false
	}
	return o.FieldsGrouping, true
}

// HasFieldsGrouping returns a boolean if a field has been set.
func (o *ReportsPutPreviousValues) HasFieldsGrouping() bool {
	if o != nil && !IsNil(o.FieldsGrouping) {
		return true
	}

	return false
}

// SetFieldsGrouping gets a reference to the given string and assigns it to the FieldsGrouping field.
func (o *ReportsPutPreviousValues) SetFieldsGrouping(v string) {
	o.FieldsGrouping = &v
}

// GetCanonicalFields returns the CanonicalFields field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ReportsPutPreviousValues) GetCanonicalFields() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.CanonicalFields
}

// GetCanonicalFieldsOk returns a tuple with the CanonicalFields field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ReportsPutPreviousValues) GetCanonicalFieldsOk() (*interface{}, bool) {
	if o == nil || IsNil(o.CanonicalFields) {
		return nil, false
	}
	return &o.CanonicalFields, true
}

// HasCanonicalFields returns a boolean if a field has been set.
func (o *ReportsPutPreviousValues) HasCanonicalFields() bool {
	if o != nil && !IsNil(o.CanonicalFields) {
		return true
	}

	return false
}

// SetCanonicalFields gets a reference to the given interface{} and assigns it to the CanonicalFields field.
func (o *ReportsPutPreviousValues) SetCanonicalFields(v interface{}) {
	o.CanonicalFields = v
}

// GetShareLinkUserId returns the ShareLinkUserId field value if set, zero value otherwise.
func (o *ReportsPutPreviousValues) GetShareLinkUserId() int32 {
	if o == nil || IsNil(o.ShareLinkUserId) {
		var ret int32
		return ret
	}
	return *o.ShareLinkUserId
}

// GetShareLinkUserIdOk returns a tuple with the ShareLinkUserId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ReportsPutPreviousValues) GetShareLinkUserIdOk() (*int32, bool) {
	if o == nil || IsNil(o.ShareLinkUserId) {
		return nil, false
	}
	return o.ShareLinkUserId, true
}

// HasShareLinkUserId returns a boolean if a field has been set.
func (o *ReportsPutPreviousValues) HasShareLinkUserId() bool {
	if o != nil && !IsNil(o.ShareLinkUserId) {
		return true
	}

	return false
}

// SetShareLinkUserId gets a reference to the given int32 and assigns it to the ShareLinkUserId field.
func (o *ReportsPutPreviousValues) SetShareLinkUserId(v int32) {
	o.ShareLinkUserId = &v
}

// GetReportOptions returns the ReportOptions field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ReportsPutPreviousValues) GetReportOptions() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.ReportOptions
}

// GetReportOptionsOk returns a tuple with the ReportOptions field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ReportsPutPreviousValues) GetReportOptionsOk() (*interface{}, bool) {
	if o == nil || IsNil(o.ReportOptions) {
		return nil, false
	}
	return &o.ReportOptions, true
}

// HasReportOptions returns a boolean if a field has been set.
func (o *ReportsPutPreviousValues) HasReportOptions() bool {
	if o != nil && !IsNil(o.ReportOptions) {
		return true
	}

	return false
}

// SetReportOptions gets a reference to the given interface{} and assigns it to the ReportOptions field.
func (o *ReportsPutPreviousValues) SetReportOptions(v interface{}) {
	o.ReportOptions = v
}

func (o ReportsPutPreviousValues) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ReportsPutPreviousValues) ToMap() (map[string]interface{}, error) {
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
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !IsNil(o.QueryFields) {
		toSerialize["query_fields"] = o.QueryFields
	}
	if !IsNil(o.QueryFilters) {
		toSerialize["query_filters"] = o.QueryFilters
	}
	if !IsNil(o.ShareToken) {
		toSerialize["share_token"] = o.ShareToken
	}
	if !IsNil(o.Type) {
		toSerialize["type"] = o.Type
	}
	if !IsNil(o.FieldsGrouping) {
		toSerialize["fields_grouping"] = o.FieldsGrouping
	}
	if o.CanonicalFields != nil {
		toSerialize["canonical_fields"] = o.CanonicalFields
	}
	if !IsNil(o.ShareLinkUserId) {
		toSerialize["share_link_user_id"] = o.ShareLinkUserId
	}
	if o.ReportOptions != nil {
		toSerialize["report_options"] = o.ReportOptions
	}
	return toSerialize, nil
}

type NullableReportsPutPreviousValues struct {
	value *ReportsPutPreviousValues
	isSet bool
}

func (v NullableReportsPutPreviousValues) Get() *ReportsPutPreviousValues {
	return v.value
}

func (v *NullableReportsPutPreviousValues) Set(val *ReportsPutPreviousValues) {
	v.value = val
	v.isSet = true
}

func (v NullableReportsPutPreviousValues) IsSet() bool {
	return v.isSet
}

func (v *NullableReportsPutPreviousValues) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableReportsPutPreviousValues(val *ReportsPutPreviousValues) *NullableReportsPutPreviousValues {
	return &NullableReportsPutPreviousValues{value: val, isSet: true}
}

func (v NullableReportsPutPreviousValues) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableReportsPutPreviousValues) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

