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

// checks if the AttachmentsPutPreviousValues type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AttachmentsPutPreviousValues{}

// AttachmentsPutPreviousValues struct for AttachmentsPutPreviousValues
type AttachmentsPutPreviousValues struct {
	// Note: This is a Primary Key.<pk/>
	Id *int32 `json:"id,omitempty"`
	CreatedAt *string `json:"created_at,omitempty"`
	UpdatedAt *string `json:"updated_at,omitempty"`
	DeletedAt *string `json:"deleted_at,omitempty"`
	Type *string `json:"type,omitempty"`
	// Note: This is a Foreign Key to `sheets.id`.<fk table='sheets' column='id'/>
	SheetId *int32 `json:"sheet_id,omitempty"`
	// Note: This is a Foreign Key to `documents.id`.<fk table='documents' column='id'/>
	DocumentId *int32 `json:"document_id,omitempty"`
	AttachableId *int32 `json:"attachable_id,omitempty"`
	AttachableType *string `json:"attachable_type,omitempty"`
	// Note: This is a Foreign Key to `users.id`.<fk table='users' column='id'/>
	LockUserId *int32 `json:"lock_user_id,omitempty"`
	LockedAt *string `json:"locked_at,omitempty"`
	Name *string `json:"name,omitempty"`
	// Note: This is a Foreign Key to `files.id`.<fk table='files' column='id'/>
	FileId *int32 `json:"file_id,omitempty"`
	SortOrder *int32 `json:"sort_order,omitempty"`
	// Note: This is a Foreign Key to `files.id`.<fk table='files' column='id'/>
	LinkifySourceFileId *int32 `json:"linkify_source_file_id,omitempty"`
	IsLinkified *bool `json:"is_linkified,omitempty"`
	LastLockUserId *int32 `json:"last_lock_user_id,omitempty"`
	WordReportOptions interface{} `json:"word_report_options,omitempty"`
	Status *string `json:"status,omitempty"`
	// Note: This is a Foreign Key to `users.id`.<fk table='users' column='id'/>
	ReviewedByUserId *int32 `json:"reviewed_by_user_id,omitempty"`
	ReviewedDate *string `json:"reviewed_date,omitempty"`
	// Note: This is a Foreign Key to `users.id`.<fk table='users' column='id'/>
	AmendedByUserId *int32 `json:"amended_by_user_id,omitempty"`
	AmendedDate *string `json:"amended_date,omitempty"`
	BrokenLinkCount *int32 `json:"broken_link_count,omitempty"`
}

// NewAttachmentsPutPreviousValues instantiates a new AttachmentsPutPreviousValues object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAttachmentsPutPreviousValues() *AttachmentsPutPreviousValues {
	this := AttachmentsPutPreviousValues{}
	var isLinkified bool = false
	this.IsLinkified = &isLinkified
	var status string = "unverified"
	this.Status = &status
	return &this
}

// NewAttachmentsPutPreviousValuesWithDefaults instantiates a new AttachmentsPutPreviousValues object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAttachmentsPutPreviousValuesWithDefaults() *AttachmentsPutPreviousValues {
	this := AttachmentsPutPreviousValues{}
	var isLinkified bool = false
	this.IsLinkified = &isLinkified
	var status string = "unverified"
	this.Status = &status
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *AttachmentsPutPreviousValues) GetId() int32 {
	if o == nil || IsNil(o.Id) {
		var ret int32
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AttachmentsPutPreviousValues) GetIdOk() (*int32, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *AttachmentsPutPreviousValues) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given int32 and assigns it to the Id field.
func (o *AttachmentsPutPreviousValues) SetId(v int32) {
	o.Id = &v
}

// GetCreatedAt returns the CreatedAt field value if set, zero value otherwise.
func (o *AttachmentsPutPreviousValues) GetCreatedAt() string {
	if o == nil || IsNil(o.CreatedAt) {
		var ret string
		return ret
	}
	return *o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AttachmentsPutPreviousValues) GetCreatedAtOk() (*string, bool) {
	if o == nil || IsNil(o.CreatedAt) {
		return nil, false
	}
	return o.CreatedAt, true
}

// HasCreatedAt returns a boolean if a field has been set.
func (o *AttachmentsPutPreviousValues) HasCreatedAt() bool {
	if o != nil && !IsNil(o.CreatedAt) {
		return true
	}

	return false
}

// SetCreatedAt gets a reference to the given string and assigns it to the CreatedAt field.
func (o *AttachmentsPutPreviousValues) SetCreatedAt(v string) {
	o.CreatedAt = &v
}

// GetUpdatedAt returns the UpdatedAt field value if set, zero value otherwise.
func (o *AttachmentsPutPreviousValues) GetUpdatedAt() string {
	if o == nil || IsNil(o.UpdatedAt) {
		var ret string
		return ret
	}
	return *o.UpdatedAt
}

// GetUpdatedAtOk returns a tuple with the UpdatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AttachmentsPutPreviousValues) GetUpdatedAtOk() (*string, bool) {
	if o == nil || IsNil(o.UpdatedAt) {
		return nil, false
	}
	return o.UpdatedAt, true
}

// HasUpdatedAt returns a boolean if a field has been set.
func (o *AttachmentsPutPreviousValues) HasUpdatedAt() bool {
	if o != nil && !IsNil(o.UpdatedAt) {
		return true
	}

	return false
}

// SetUpdatedAt gets a reference to the given string and assigns it to the UpdatedAt field.
func (o *AttachmentsPutPreviousValues) SetUpdatedAt(v string) {
	o.UpdatedAt = &v
}

// GetDeletedAt returns the DeletedAt field value if set, zero value otherwise.
func (o *AttachmentsPutPreviousValues) GetDeletedAt() string {
	if o == nil || IsNil(o.DeletedAt) {
		var ret string
		return ret
	}
	return *o.DeletedAt
}

// GetDeletedAtOk returns a tuple with the DeletedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AttachmentsPutPreviousValues) GetDeletedAtOk() (*string, bool) {
	if o == nil || IsNil(o.DeletedAt) {
		return nil, false
	}
	return o.DeletedAt, true
}

// HasDeletedAt returns a boolean if a field has been set.
func (o *AttachmentsPutPreviousValues) HasDeletedAt() bool {
	if o != nil && !IsNil(o.DeletedAt) {
		return true
	}

	return false
}

// SetDeletedAt gets a reference to the given string and assigns it to the DeletedAt field.
func (o *AttachmentsPutPreviousValues) SetDeletedAt(v string) {
	o.DeletedAt = &v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *AttachmentsPutPreviousValues) GetType() string {
	if o == nil || IsNil(o.Type) {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AttachmentsPutPreviousValues) GetTypeOk() (*string, bool) {
	if o == nil || IsNil(o.Type) {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *AttachmentsPutPreviousValues) HasType() bool {
	if o != nil && !IsNil(o.Type) {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *AttachmentsPutPreviousValues) SetType(v string) {
	o.Type = &v
}

// GetSheetId returns the SheetId field value if set, zero value otherwise.
func (o *AttachmentsPutPreviousValues) GetSheetId() int32 {
	if o == nil || IsNil(o.SheetId) {
		var ret int32
		return ret
	}
	return *o.SheetId
}

// GetSheetIdOk returns a tuple with the SheetId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AttachmentsPutPreviousValues) GetSheetIdOk() (*int32, bool) {
	if o == nil || IsNil(o.SheetId) {
		return nil, false
	}
	return o.SheetId, true
}

// HasSheetId returns a boolean if a field has been set.
func (o *AttachmentsPutPreviousValues) HasSheetId() bool {
	if o != nil && !IsNil(o.SheetId) {
		return true
	}

	return false
}

// SetSheetId gets a reference to the given int32 and assigns it to the SheetId field.
func (o *AttachmentsPutPreviousValues) SetSheetId(v int32) {
	o.SheetId = &v
}

// GetDocumentId returns the DocumentId field value if set, zero value otherwise.
func (o *AttachmentsPutPreviousValues) GetDocumentId() int32 {
	if o == nil || IsNil(o.DocumentId) {
		var ret int32
		return ret
	}
	return *o.DocumentId
}

// GetDocumentIdOk returns a tuple with the DocumentId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AttachmentsPutPreviousValues) GetDocumentIdOk() (*int32, bool) {
	if o == nil || IsNil(o.DocumentId) {
		return nil, false
	}
	return o.DocumentId, true
}

// HasDocumentId returns a boolean if a field has been set.
func (o *AttachmentsPutPreviousValues) HasDocumentId() bool {
	if o != nil && !IsNil(o.DocumentId) {
		return true
	}

	return false
}

// SetDocumentId gets a reference to the given int32 and assigns it to the DocumentId field.
func (o *AttachmentsPutPreviousValues) SetDocumentId(v int32) {
	o.DocumentId = &v
}

// GetAttachableId returns the AttachableId field value if set, zero value otherwise.
func (o *AttachmentsPutPreviousValues) GetAttachableId() int32 {
	if o == nil || IsNil(o.AttachableId) {
		var ret int32
		return ret
	}
	return *o.AttachableId
}

// GetAttachableIdOk returns a tuple with the AttachableId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AttachmentsPutPreviousValues) GetAttachableIdOk() (*int32, bool) {
	if o == nil || IsNil(o.AttachableId) {
		return nil, false
	}
	return o.AttachableId, true
}

// HasAttachableId returns a boolean if a field has been set.
func (o *AttachmentsPutPreviousValues) HasAttachableId() bool {
	if o != nil && !IsNil(o.AttachableId) {
		return true
	}

	return false
}

// SetAttachableId gets a reference to the given int32 and assigns it to the AttachableId field.
func (o *AttachmentsPutPreviousValues) SetAttachableId(v int32) {
	o.AttachableId = &v
}

// GetAttachableType returns the AttachableType field value if set, zero value otherwise.
func (o *AttachmentsPutPreviousValues) GetAttachableType() string {
	if o == nil || IsNil(o.AttachableType) {
		var ret string
		return ret
	}
	return *o.AttachableType
}

// GetAttachableTypeOk returns a tuple with the AttachableType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AttachmentsPutPreviousValues) GetAttachableTypeOk() (*string, bool) {
	if o == nil || IsNil(o.AttachableType) {
		return nil, false
	}
	return o.AttachableType, true
}

// HasAttachableType returns a boolean if a field has been set.
func (o *AttachmentsPutPreviousValues) HasAttachableType() bool {
	if o != nil && !IsNil(o.AttachableType) {
		return true
	}

	return false
}

// SetAttachableType gets a reference to the given string and assigns it to the AttachableType field.
func (o *AttachmentsPutPreviousValues) SetAttachableType(v string) {
	o.AttachableType = &v
}

// GetLockUserId returns the LockUserId field value if set, zero value otherwise.
func (o *AttachmentsPutPreviousValues) GetLockUserId() int32 {
	if o == nil || IsNil(o.LockUserId) {
		var ret int32
		return ret
	}
	return *o.LockUserId
}

// GetLockUserIdOk returns a tuple with the LockUserId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AttachmentsPutPreviousValues) GetLockUserIdOk() (*int32, bool) {
	if o == nil || IsNil(o.LockUserId) {
		return nil, false
	}
	return o.LockUserId, true
}

// HasLockUserId returns a boolean if a field has been set.
func (o *AttachmentsPutPreviousValues) HasLockUserId() bool {
	if o != nil && !IsNil(o.LockUserId) {
		return true
	}

	return false
}

// SetLockUserId gets a reference to the given int32 and assigns it to the LockUserId field.
func (o *AttachmentsPutPreviousValues) SetLockUserId(v int32) {
	o.LockUserId = &v
}

// GetLockedAt returns the LockedAt field value if set, zero value otherwise.
func (o *AttachmentsPutPreviousValues) GetLockedAt() string {
	if o == nil || IsNil(o.LockedAt) {
		var ret string
		return ret
	}
	return *o.LockedAt
}

// GetLockedAtOk returns a tuple with the LockedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AttachmentsPutPreviousValues) GetLockedAtOk() (*string, bool) {
	if o == nil || IsNil(o.LockedAt) {
		return nil, false
	}
	return o.LockedAt, true
}

// HasLockedAt returns a boolean if a field has been set.
func (o *AttachmentsPutPreviousValues) HasLockedAt() bool {
	if o != nil && !IsNil(o.LockedAt) {
		return true
	}

	return false
}

// SetLockedAt gets a reference to the given string and assigns it to the LockedAt field.
func (o *AttachmentsPutPreviousValues) SetLockedAt(v string) {
	o.LockedAt = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *AttachmentsPutPreviousValues) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AttachmentsPutPreviousValues) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *AttachmentsPutPreviousValues) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *AttachmentsPutPreviousValues) SetName(v string) {
	o.Name = &v
}

// GetFileId returns the FileId field value if set, zero value otherwise.
func (o *AttachmentsPutPreviousValues) GetFileId() int32 {
	if o == nil || IsNil(o.FileId) {
		var ret int32
		return ret
	}
	return *o.FileId
}

// GetFileIdOk returns a tuple with the FileId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AttachmentsPutPreviousValues) GetFileIdOk() (*int32, bool) {
	if o == nil || IsNil(o.FileId) {
		return nil, false
	}
	return o.FileId, true
}

// HasFileId returns a boolean if a field has been set.
func (o *AttachmentsPutPreviousValues) HasFileId() bool {
	if o != nil && !IsNil(o.FileId) {
		return true
	}

	return false
}

// SetFileId gets a reference to the given int32 and assigns it to the FileId field.
func (o *AttachmentsPutPreviousValues) SetFileId(v int32) {
	o.FileId = &v
}

// GetSortOrder returns the SortOrder field value if set, zero value otherwise.
func (o *AttachmentsPutPreviousValues) GetSortOrder() int32 {
	if o == nil || IsNil(o.SortOrder) {
		var ret int32
		return ret
	}
	return *o.SortOrder
}

// GetSortOrderOk returns a tuple with the SortOrder field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AttachmentsPutPreviousValues) GetSortOrderOk() (*int32, bool) {
	if o == nil || IsNil(o.SortOrder) {
		return nil, false
	}
	return o.SortOrder, true
}

// HasSortOrder returns a boolean if a field has been set.
func (o *AttachmentsPutPreviousValues) HasSortOrder() bool {
	if o != nil && !IsNil(o.SortOrder) {
		return true
	}

	return false
}

// SetSortOrder gets a reference to the given int32 and assigns it to the SortOrder field.
func (o *AttachmentsPutPreviousValues) SetSortOrder(v int32) {
	o.SortOrder = &v
}

// GetLinkifySourceFileId returns the LinkifySourceFileId field value if set, zero value otherwise.
func (o *AttachmentsPutPreviousValues) GetLinkifySourceFileId() int32 {
	if o == nil || IsNil(o.LinkifySourceFileId) {
		var ret int32
		return ret
	}
	return *o.LinkifySourceFileId
}

// GetLinkifySourceFileIdOk returns a tuple with the LinkifySourceFileId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AttachmentsPutPreviousValues) GetLinkifySourceFileIdOk() (*int32, bool) {
	if o == nil || IsNil(o.LinkifySourceFileId) {
		return nil, false
	}
	return o.LinkifySourceFileId, true
}

// HasLinkifySourceFileId returns a boolean if a field has been set.
func (o *AttachmentsPutPreviousValues) HasLinkifySourceFileId() bool {
	if o != nil && !IsNil(o.LinkifySourceFileId) {
		return true
	}

	return false
}

// SetLinkifySourceFileId gets a reference to the given int32 and assigns it to the LinkifySourceFileId field.
func (o *AttachmentsPutPreviousValues) SetLinkifySourceFileId(v int32) {
	o.LinkifySourceFileId = &v
}

// GetIsLinkified returns the IsLinkified field value if set, zero value otherwise.
func (o *AttachmentsPutPreviousValues) GetIsLinkified() bool {
	if o == nil || IsNil(o.IsLinkified) {
		var ret bool
		return ret
	}
	return *o.IsLinkified
}

// GetIsLinkifiedOk returns a tuple with the IsLinkified field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AttachmentsPutPreviousValues) GetIsLinkifiedOk() (*bool, bool) {
	if o == nil || IsNil(o.IsLinkified) {
		return nil, false
	}
	return o.IsLinkified, true
}

// HasIsLinkified returns a boolean if a field has been set.
func (o *AttachmentsPutPreviousValues) HasIsLinkified() bool {
	if o != nil && !IsNil(o.IsLinkified) {
		return true
	}

	return false
}

// SetIsLinkified gets a reference to the given bool and assigns it to the IsLinkified field.
func (o *AttachmentsPutPreviousValues) SetIsLinkified(v bool) {
	o.IsLinkified = &v
}

// GetLastLockUserId returns the LastLockUserId field value if set, zero value otherwise.
func (o *AttachmentsPutPreviousValues) GetLastLockUserId() int32 {
	if o == nil || IsNil(o.LastLockUserId) {
		var ret int32
		return ret
	}
	return *o.LastLockUserId
}

// GetLastLockUserIdOk returns a tuple with the LastLockUserId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AttachmentsPutPreviousValues) GetLastLockUserIdOk() (*int32, bool) {
	if o == nil || IsNil(o.LastLockUserId) {
		return nil, false
	}
	return o.LastLockUserId, true
}

// HasLastLockUserId returns a boolean if a field has been set.
func (o *AttachmentsPutPreviousValues) HasLastLockUserId() bool {
	if o != nil && !IsNil(o.LastLockUserId) {
		return true
	}

	return false
}

// SetLastLockUserId gets a reference to the given int32 and assigns it to the LastLockUserId field.
func (o *AttachmentsPutPreviousValues) SetLastLockUserId(v int32) {
	o.LastLockUserId = &v
}

// GetWordReportOptions returns the WordReportOptions field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *AttachmentsPutPreviousValues) GetWordReportOptions() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.WordReportOptions
}

// GetWordReportOptionsOk returns a tuple with the WordReportOptions field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *AttachmentsPutPreviousValues) GetWordReportOptionsOk() (*interface{}, bool) {
	if o == nil || IsNil(o.WordReportOptions) {
		return nil, false
	}
	return &o.WordReportOptions, true
}

// HasWordReportOptions returns a boolean if a field has been set.
func (o *AttachmentsPutPreviousValues) HasWordReportOptions() bool {
	if o != nil && !IsNil(o.WordReportOptions) {
		return true
	}

	return false
}

// SetWordReportOptions gets a reference to the given interface{} and assigns it to the WordReportOptions field.
func (o *AttachmentsPutPreviousValues) SetWordReportOptions(v interface{}) {
	o.WordReportOptions = v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *AttachmentsPutPreviousValues) GetStatus() string {
	if o == nil || IsNil(o.Status) {
		var ret string
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AttachmentsPutPreviousValues) GetStatusOk() (*string, bool) {
	if o == nil || IsNil(o.Status) {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *AttachmentsPutPreviousValues) HasStatus() bool {
	if o != nil && !IsNil(o.Status) {
		return true
	}

	return false
}

// SetStatus gets a reference to the given string and assigns it to the Status field.
func (o *AttachmentsPutPreviousValues) SetStatus(v string) {
	o.Status = &v
}

// GetReviewedByUserId returns the ReviewedByUserId field value if set, zero value otherwise.
func (o *AttachmentsPutPreviousValues) GetReviewedByUserId() int32 {
	if o == nil || IsNil(o.ReviewedByUserId) {
		var ret int32
		return ret
	}
	return *o.ReviewedByUserId
}

// GetReviewedByUserIdOk returns a tuple with the ReviewedByUserId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AttachmentsPutPreviousValues) GetReviewedByUserIdOk() (*int32, bool) {
	if o == nil || IsNil(o.ReviewedByUserId) {
		return nil, false
	}
	return o.ReviewedByUserId, true
}

// HasReviewedByUserId returns a boolean if a field has been set.
func (o *AttachmentsPutPreviousValues) HasReviewedByUserId() bool {
	if o != nil && !IsNil(o.ReviewedByUserId) {
		return true
	}

	return false
}

// SetReviewedByUserId gets a reference to the given int32 and assigns it to the ReviewedByUserId field.
func (o *AttachmentsPutPreviousValues) SetReviewedByUserId(v int32) {
	o.ReviewedByUserId = &v
}

// GetReviewedDate returns the ReviewedDate field value if set, zero value otherwise.
func (o *AttachmentsPutPreviousValues) GetReviewedDate() string {
	if o == nil || IsNil(o.ReviewedDate) {
		var ret string
		return ret
	}
	return *o.ReviewedDate
}

// GetReviewedDateOk returns a tuple with the ReviewedDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AttachmentsPutPreviousValues) GetReviewedDateOk() (*string, bool) {
	if o == nil || IsNil(o.ReviewedDate) {
		return nil, false
	}
	return o.ReviewedDate, true
}

// HasReviewedDate returns a boolean if a field has been set.
func (o *AttachmentsPutPreviousValues) HasReviewedDate() bool {
	if o != nil && !IsNil(o.ReviewedDate) {
		return true
	}

	return false
}

// SetReviewedDate gets a reference to the given string and assigns it to the ReviewedDate field.
func (o *AttachmentsPutPreviousValues) SetReviewedDate(v string) {
	o.ReviewedDate = &v
}

// GetAmendedByUserId returns the AmendedByUserId field value if set, zero value otherwise.
func (o *AttachmentsPutPreviousValues) GetAmendedByUserId() int32 {
	if o == nil || IsNil(o.AmendedByUserId) {
		var ret int32
		return ret
	}
	return *o.AmendedByUserId
}

// GetAmendedByUserIdOk returns a tuple with the AmendedByUserId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AttachmentsPutPreviousValues) GetAmendedByUserIdOk() (*int32, bool) {
	if o == nil || IsNil(o.AmendedByUserId) {
		return nil, false
	}
	return o.AmendedByUserId, true
}

// HasAmendedByUserId returns a boolean if a field has been set.
func (o *AttachmentsPutPreviousValues) HasAmendedByUserId() bool {
	if o != nil && !IsNil(o.AmendedByUserId) {
		return true
	}

	return false
}

// SetAmendedByUserId gets a reference to the given int32 and assigns it to the AmendedByUserId field.
func (o *AttachmentsPutPreviousValues) SetAmendedByUserId(v int32) {
	o.AmendedByUserId = &v
}

// GetAmendedDate returns the AmendedDate field value if set, zero value otherwise.
func (o *AttachmentsPutPreviousValues) GetAmendedDate() string {
	if o == nil || IsNil(o.AmendedDate) {
		var ret string
		return ret
	}
	return *o.AmendedDate
}

// GetAmendedDateOk returns a tuple with the AmendedDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AttachmentsPutPreviousValues) GetAmendedDateOk() (*string, bool) {
	if o == nil || IsNil(o.AmendedDate) {
		return nil, false
	}
	return o.AmendedDate, true
}

// HasAmendedDate returns a boolean if a field has been set.
func (o *AttachmentsPutPreviousValues) HasAmendedDate() bool {
	if o != nil && !IsNil(o.AmendedDate) {
		return true
	}

	return false
}

// SetAmendedDate gets a reference to the given string and assigns it to the AmendedDate field.
func (o *AttachmentsPutPreviousValues) SetAmendedDate(v string) {
	o.AmendedDate = &v
}

// GetBrokenLinkCount returns the BrokenLinkCount field value if set, zero value otherwise.
func (o *AttachmentsPutPreviousValues) GetBrokenLinkCount() int32 {
	if o == nil || IsNil(o.BrokenLinkCount) {
		var ret int32
		return ret
	}
	return *o.BrokenLinkCount
}

// GetBrokenLinkCountOk returns a tuple with the BrokenLinkCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AttachmentsPutPreviousValues) GetBrokenLinkCountOk() (*int32, bool) {
	if o == nil || IsNil(o.BrokenLinkCount) {
		return nil, false
	}
	return o.BrokenLinkCount, true
}

// HasBrokenLinkCount returns a boolean if a field has been set.
func (o *AttachmentsPutPreviousValues) HasBrokenLinkCount() bool {
	if o != nil && !IsNil(o.BrokenLinkCount) {
		return true
	}

	return false
}

// SetBrokenLinkCount gets a reference to the given int32 and assigns it to the BrokenLinkCount field.
func (o *AttachmentsPutPreviousValues) SetBrokenLinkCount(v int32) {
	o.BrokenLinkCount = &v
}

func (o AttachmentsPutPreviousValues) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AttachmentsPutPreviousValues) ToMap() (map[string]interface{}, error) {
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
	if !IsNil(o.Type) {
		toSerialize["type"] = o.Type
	}
	if !IsNil(o.SheetId) {
		toSerialize["sheet_id"] = o.SheetId
	}
	if !IsNil(o.DocumentId) {
		toSerialize["document_id"] = o.DocumentId
	}
	if !IsNil(o.AttachableId) {
		toSerialize["attachable_id"] = o.AttachableId
	}
	if !IsNil(o.AttachableType) {
		toSerialize["attachable_type"] = o.AttachableType
	}
	if !IsNil(o.LockUserId) {
		toSerialize["lock_user_id"] = o.LockUserId
	}
	if !IsNil(o.LockedAt) {
		toSerialize["locked_at"] = o.LockedAt
	}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !IsNil(o.FileId) {
		toSerialize["file_id"] = o.FileId
	}
	if !IsNil(o.SortOrder) {
		toSerialize["sort_order"] = o.SortOrder
	}
	if !IsNil(o.LinkifySourceFileId) {
		toSerialize["linkify_source_file_id"] = o.LinkifySourceFileId
	}
	if !IsNil(o.IsLinkified) {
		toSerialize["is_linkified"] = o.IsLinkified
	}
	if !IsNil(o.LastLockUserId) {
		toSerialize["last_lock_user_id"] = o.LastLockUserId
	}
	if o.WordReportOptions != nil {
		toSerialize["word_report_options"] = o.WordReportOptions
	}
	if !IsNil(o.Status) {
		toSerialize["status"] = o.Status
	}
	if !IsNil(o.ReviewedByUserId) {
		toSerialize["reviewed_by_user_id"] = o.ReviewedByUserId
	}
	if !IsNil(o.ReviewedDate) {
		toSerialize["reviewed_date"] = o.ReviewedDate
	}
	if !IsNil(o.AmendedByUserId) {
		toSerialize["amended_by_user_id"] = o.AmendedByUserId
	}
	if !IsNil(o.AmendedDate) {
		toSerialize["amended_date"] = o.AmendedDate
	}
	if !IsNil(o.BrokenLinkCount) {
		toSerialize["broken_link_count"] = o.BrokenLinkCount
	}
	return toSerialize, nil
}

type NullableAttachmentsPutPreviousValues struct {
	value *AttachmentsPutPreviousValues
	isSet bool
}

func (v NullableAttachmentsPutPreviousValues) Get() *AttachmentsPutPreviousValues {
	return v.value
}

func (v *NullableAttachmentsPutPreviousValues) Set(val *AttachmentsPutPreviousValues) {
	v.value = val
	v.isSet = true
}

func (v NullableAttachmentsPutPreviousValues) IsSet() bool {
	return v.isSet
}

func (v *NullableAttachmentsPutPreviousValues) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAttachmentsPutPreviousValues(val *AttachmentsPutPreviousValues) *NullableAttachmentsPutPreviousValues {
	return &NullableAttachmentsPutPreviousValues{value: val, isSet: true}
}

func (v NullableAttachmentsPutPreviousValues) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAttachmentsPutPreviousValues) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


