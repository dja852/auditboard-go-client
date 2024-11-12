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

// checks if the ReportsPutReport type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ReportsPutReport{}

// ReportsPutReport struct for ReportsPutReport
type ReportsPutReport struct {
	// Note: This is a Primary Key.<pk/>
	Id *int32 `json:"id,omitempty"`
	CreatedAt *string `json:"created_at,omitempty"`
	UpdatedAt *string `json:"updated_at,omitempty"`
	DeletedAt *string `json:"deleted_at,omitempty"`
	Name string `json:"name"`
	QueryFields string `json:"query_fields"`
	QueryFilters string `json:"query_filters"`
	ShareToken *string `json:"share_token,omitempty"`
	Type string `json:"type"`
	FieldsGrouping *string `json:"fields_grouping,omitempty"`
	CanonicalFields interface{} `json:"canonical_fields,omitempty"`
	// Note: This is a Foreign Key to `users.id`.<fk table='users' column='id'/>
	ShareLinkUserId *int32 `json:"share_link_user_id,omitempty"`
	ReportOptions interface{} `json:"report_options,omitempty"`
}

type _ReportsPutReport ReportsPutReport

// NewReportsPutReport instantiates a new ReportsPutReport object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewReportsPutReport(name string, queryFields string, queryFilters string, type_ string) *ReportsPutReport {
	this := ReportsPutReport{}
	this.Name = name
	this.QueryFields = queryFields
	this.QueryFilters = queryFilters
	this.Type = type_
	return &this
}

// NewReportsPutReportWithDefaults instantiates a new ReportsPutReport object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewReportsPutReportWithDefaults() *ReportsPutReport {
	this := ReportsPutReport{}
	var type_ string = "Control"
	this.Type = type_
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *ReportsPutReport) GetId() int32 {
	if o == nil || IsNil(o.Id) {
		var ret int32
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ReportsPutReport) GetIdOk() (*int32, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *ReportsPutReport) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given int32 and assigns it to the Id field.
func (o *ReportsPutReport) SetId(v int32) {
	o.Id = &v
}

// GetCreatedAt returns the CreatedAt field value if set, zero value otherwise.
func (o *ReportsPutReport) GetCreatedAt() string {
	if o == nil || IsNil(o.CreatedAt) {
		var ret string
		return ret
	}
	return *o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ReportsPutReport) GetCreatedAtOk() (*string, bool) {
	if o == nil || IsNil(o.CreatedAt) {
		return nil, false
	}
	return o.CreatedAt, true
}

// HasCreatedAt returns a boolean if a field has been set.
func (o *ReportsPutReport) HasCreatedAt() bool {
	if o != nil && !IsNil(o.CreatedAt) {
		return true
	}

	return false
}

// SetCreatedAt gets a reference to the given string and assigns it to the CreatedAt field.
func (o *ReportsPutReport) SetCreatedAt(v string) {
	o.CreatedAt = &v
}

// GetUpdatedAt returns the UpdatedAt field value if set, zero value otherwise.
func (o *ReportsPutReport) GetUpdatedAt() string {
	if o == nil || IsNil(o.UpdatedAt) {
		var ret string
		return ret
	}
	return *o.UpdatedAt
}

// GetUpdatedAtOk returns a tuple with the UpdatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ReportsPutReport) GetUpdatedAtOk() (*string, bool) {
	if o == nil || IsNil(o.UpdatedAt) {
		return nil, false
	}
	return o.UpdatedAt, true
}

// HasUpdatedAt returns a boolean if a field has been set.
func (o *ReportsPutReport) HasUpdatedAt() bool {
	if o != nil && !IsNil(o.UpdatedAt) {
		return true
	}

	return false
}

// SetUpdatedAt gets a reference to the given string and assigns it to the UpdatedAt field.
func (o *ReportsPutReport) SetUpdatedAt(v string) {
	o.UpdatedAt = &v
}

// GetDeletedAt returns the DeletedAt field value if set, zero value otherwise.
func (o *ReportsPutReport) GetDeletedAt() string {
	if o == nil || IsNil(o.DeletedAt) {
		var ret string
		return ret
	}
	return *o.DeletedAt
}

// GetDeletedAtOk returns a tuple with the DeletedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ReportsPutReport) GetDeletedAtOk() (*string, bool) {
	if o == nil || IsNil(o.DeletedAt) {
		return nil, false
	}
	return o.DeletedAt, true
}

// HasDeletedAt returns a boolean if a field has been set.
func (o *ReportsPutReport) HasDeletedAt() bool {
	if o != nil && !IsNil(o.DeletedAt) {
		return true
	}

	return false
}

// SetDeletedAt gets a reference to the given string and assigns it to the DeletedAt field.
func (o *ReportsPutReport) SetDeletedAt(v string) {
	o.DeletedAt = &v
}

// GetName returns the Name field value
func (o *ReportsPutReport) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *ReportsPutReport) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *ReportsPutReport) SetName(v string) {
	o.Name = v
}

// GetQueryFields returns the QueryFields field value
func (o *ReportsPutReport) GetQueryFields() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.QueryFields
}

// GetQueryFieldsOk returns a tuple with the QueryFields field value
// and a boolean to check if the value has been set.
func (o *ReportsPutReport) GetQueryFieldsOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.QueryFields, true
}

// SetQueryFields sets field value
func (o *ReportsPutReport) SetQueryFields(v string) {
	o.QueryFields = v
}

// GetQueryFilters returns the QueryFilters field value
func (o *ReportsPutReport) GetQueryFilters() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.QueryFilters
}

// GetQueryFiltersOk returns a tuple with the QueryFilters field value
// and a boolean to check if the value has been set.
func (o *ReportsPutReport) GetQueryFiltersOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.QueryFilters, true
}

// SetQueryFilters sets field value
func (o *ReportsPutReport) SetQueryFilters(v string) {
	o.QueryFilters = v
}

// GetShareToken returns the ShareToken field value if set, zero value otherwise.
func (o *ReportsPutReport) GetShareToken() string {
	if o == nil || IsNil(o.ShareToken) {
		var ret string
		return ret
	}
	return *o.ShareToken
}

// GetShareTokenOk returns a tuple with the ShareToken field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ReportsPutReport) GetShareTokenOk() (*string, bool) {
	if o == nil || IsNil(o.ShareToken) {
		return nil, false
	}
	return o.ShareToken, true
}

// HasShareToken returns a boolean if a field has been set.
func (o *ReportsPutReport) HasShareToken() bool {
	if o != nil && !IsNil(o.ShareToken) {
		return true
	}

	return false
}

// SetShareToken gets a reference to the given string and assigns it to the ShareToken field.
func (o *ReportsPutReport) SetShareToken(v string) {
	o.ShareToken = &v
}

// GetType returns the Type field value
func (o *ReportsPutReport) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *ReportsPutReport) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *ReportsPutReport) SetType(v string) {
	o.Type = v
}

// GetFieldsGrouping returns the FieldsGrouping field value if set, zero value otherwise.
func (o *ReportsPutReport) GetFieldsGrouping() string {
	if o == nil || IsNil(o.FieldsGrouping) {
		var ret string
		return ret
	}
	return *o.FieldsGrouping
}

// GetFieldsGroupingOk returns a tuple with the FieldsGrouping field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ReportsPutReport) GetFieldsGroupingOk() (*string, bool) {
	if o == nil || IsNil(o.FieldsGrouping) {
		return nil, false
	}
	return o.FieldsGrouping, true
}

// HasFieldsGrouping returns a boolean if a field has been set.
func (o *ReportsPutReport) HasFieldsGrouping() bool {
	if o != nil && !IsNil(o.FieldsGrouping) {
		return true
	}

	return false
}

// SetFieldsGrouping gets a reference to the given string and assigns it to the FieldsGrouping field.
func (o *ReportsPutReport) SetFieldsGrouping(v string) {
	o.FieldsGrouping = &v
}

// GetCanonicalFields returns the CanonicalFields field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ReportsPutReport) GetCanonicalFields() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.CanonicalFields
}

// GetCanonicalFieldsOk returns a tuple with the CanonicalFields field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ReportsPutReport) GetCanonicalFieldsOk() (*interface{}, bool) {
	if o == nil || IsNil(o.CanonicalFields) {
		return nil, false
	}
	return &o.CanonicalFields, true
}

// HasCanonicalFields returns a boolean if a field has been set.
func (o *ReportsPutReport) HasCanonicalFields() bool {
	if o != nil && !IsNil(o.CanonicalFields) {
		return true
	}

	return false
}

// SetCanonicalFields gets a reference to the given interface{} and assigns it to the CanonicalFields field.
func (o *ReportsPutReport) SetCanonicalFields(v interface{}) {
	o.CanonicalFields = v
}

// GetShareLinkUserId returns the ShareLinkUserId field value if set, zero value otherwise.
func (o *ReportsPutReport) GetShareLinkUserId() int32 {
	if o == nil || IsNil(o.ShareLinkUserId) {
		var ret int32
		return ret
	}
	return *o.ShareLinkUserId
}

// GetShareLinkUserIdOk returns a tuple with the ShareLinkUserId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ReportsPutReport) GetShareLinkUserIdOk() (*int32, bool) {
	if o == nil || IsNil(o.ShareLinkUserId) {
		return nil, false
	}
	return o.ShareLinkUserId, true
}

// HasShareLinkUserId returns a boolean if a field has been set.
func (o *ReportsPutReport) HasShareLinkUserId() bool {
	if o != nil && !IsNil(o.ShareLinkUserId) {
		return true
	}

	return false
}

// SetShareLinkUserId gets a reference to the given int32 and assigns it to the ShareLinkUserId field.
func (o *ReportsPutReport) SetShareLinkUserId(v int32) {
	o.ShareLinkUserId = &v
}

// GetReportOptions returns the ReportOptions field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ReportsPutReport) GetReportOptions() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.ReportOptions
}

// GetReportOptionsOk returns a tuple with the ReportOptions field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ReportsPutReport) GetReportOptionsOk() (*interface{}, bool) {
	if o == nil || IsNil(o.ReportOptions) {
		return nil, false
	}
	return &o.ReportOptions, true
}

// HasReportOptions returns a boolean if a field has been set.
func (o *ReportsPutReport) HasReportOptions() bool {
	if o != nil && !IsNil(o.ReportOptions) {
		return true
	}

	return false
}

// SetReportOptions gets a reference to the given interface{} and assigns it to the ReportOptions field.
func (o *ReportsPutReport) SetReportOptions(v interface{}) {
	o.ReportOptions = v
}

func (o ReportsPutReport) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ReportsPutReport) ToMap() (map[string]interface{}, error) {
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
	toSerialize["name"] = o.Name
	toSerialize["query_fields"] = o.QueryFields
	toSerialize["query_filters"] = o.QueryFilters
	if !IsNil(o.ShareToken) {
		toSerialize["share_token"] = o.ShareToken
	}
	toSerialize["type"] = o.Type
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

type NullableReportsPutReport struct {
	value *ReportsPutReport
	isSet bool
}

func (v NullableReportsPutReport) Get() *ReportsPutReport {
	return v.value
}

func (v *NullableReportsPutReport) Set(val *ReportsPutReport) {
	v.value = val
	v.isSet = true
}

func (v NullableReportsPutReport) IsSet() bool {
	return v.isSet
}

func (v *NullableReportsPutReport) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableReportsPutReport(val *ReportsPutReport) *NullableReportsPutReport {
	return &NullableReportsPutReport{value: val, isSet: true}
}

func (v NullableReportsPutReport) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableReportsPutReport) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


