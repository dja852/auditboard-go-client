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

// checks if the CommentsPutComment type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CommentsPutComment{}

// CommentsPutComment struct for CommentsPutComment
type CommentsPutComment struct {
	// Note: This is a Primary Key.<pk/>
	Id *int32 `json:"id,omitempty"`
	// Note: This is a Foreign Key to `comments.id`.<fk table='comments' column='id'/>
	ParentCommentId *int32 `json:"parent_comment_id,omitempty"`
	// Note: This is a Foreign Key to `users.id`.<fk table='users' column='id'/>
	OwnerUserId *int32 `json:"owner_user_id,omitempty"`
	Type *string `json:"type,omitempty"`
	Status *string `json:"status,omitempty"`
	MessageRaw string `json:"message_raw"`
	MessageHtml string `json:"message_html"`
	CreatedAt *string `json:"created_at,omitempty"`
	UpdatedAt *string `json:"updated_at,omitempty"`
	DeletedAt *string `json:"deleted_at,omitempty"`
	CommentableId *int32 `json:"commentable_id,omitempty"`
	CommentableType *string `json:"commentable_type,omitempty"`
	ReferenceMeta interface{} `json:"reference_meta,omitempty"`
	// Note: This is a Foreign Key to `attachments.id`.<fk table='attachments' column='id'/>
	OriginAttachmentId *int32 `json:"origin_attachment_id,omitempty"`
	// Note: This is a Foreign Key to `files.id`.<fk table='files' column='id'/>
	OriginFileId *int32 `json:"origin_file_id,omitempty"`
	Scopes interface{} `json:"scopes"`
	ExternalIntegrationId string `json:"external_integration_id"`
	ExternalIntegrationAuthor interface{} `json:"external_integration_author,omitempty"`
	OriginatingPath *string `json:"originating_path,omitempty"`
}

type _CommentsPutComment CommentsPutComment

// NewCommentsPutComment instantiates a new CommentsPutComment object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCommentsPutComment(messageRaw string, messageHtml string, scopes interface{}, externalIntegrationId string) *CommentsPutComment {
	this := CommentsPutComment{}
	this.MessageRaw = messageRaw
	this.MessageHtml = messageHtml
	this.Scopes = scopes
	this.ExternalIntegrationId = externalIntegrationId
	return &this
}

// NewCommentsPutCommentWithDefaults instantiates a new CommentsPutComment object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCommentsPutCommentWithDefaults() *CommentsPutComment {
	this := CommentsPutComment{}
	var externalIntegrationId string = ""
	this.ExternalIntegrationId = externalIntegrationId
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *CommentsPutComment) GetId() int32 {
	if o == nil || IsNil(o.Id) {
		var ret int32
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CommentsPutComment) GetIdOk() (*int32, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *CommentsPutComment) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given int32 and assigns it to the Id field.
func (o *CommentsPutComment) SetId(v int32) {
	o.Id = &v
}

// GetParentCommentId returns the ParentCommentId field value if set, zero value otherwise.
func (o *CommentsPutComment) GetParentCommentId() int32 {
	if o == nil || IsNil(o.ParentCommentId) {
		var ret int32
		return ret
	}
	return *o.ParentCommentId
}

// GetParentCommentIdOk returns a tuple with the ParentCommentId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CommentsPutComment) GetParentCommentIdOk() (*int32, bool) {
	if o == nil || IsNil(o.ParentCommentId) {
		return nil, false
	}
	return o.ParentCommentId, true
}

// HasParentCommentId returns a boolean if a field has been set.
func (o *CommentsPutComment) HasParentCommentId() bool {
	if o != nil && !IsNil(o.ParentCommentId) {
		return true
	}

	return false
}

// SetParentCommentId gets a reference to the given int32 and assigns it to the ParentCommentId field.
func (o *CommentsPutComment) SetParentCommentId(v int32) {
	o.ParentCommentId = &v
}

// GetOwnerUserId returns the OwnerUserId field value if set, zero value otherwise.
func (o *CommentsPutComment) GetOwnerUserId() int32 {
	if o == nil || IsNil(o.OwnerUserId) {
		var ret int32
		return ret
	}
	return *o.OwnerUserId
}

// GetOwnerUserIdOk returns a tuple with the OwnerUserId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CommentsPutComment) GetOwnerUserIdOk() (*int32, bool) {
	if o == nil || IsNil(o.OwnerUserId) {
		return nil, false
	}
	return o.OwnerUserId, true
}

// HasOwnerUserId returns a boolean if a field has been set.
func (o *CommentsPutComment) HasOwnerUserId() bool {
	if o != nil && !IsNil(o.OwnerUserId) {
		return true
	}

	return false
}

// SetOwnerUserId gets a reference to the given int32 and assigns it to the OwnerUserId field.
func (o *CommentsPutComment) SetOwnerUserId(v int32) {
	o.OwnerUserId = &v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *CommentsPutComment) GetType() string {
	if o == nil || IsNil(o.Type) {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CommentsPutComment) GetTypeOk() (*string, bool) {
	if o == nil || IsNil(o.Type) {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *CommentsPutComment) HasType() bool {
	if o != nil && !IsNil(o.Type) {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *CommentsPutComment) SetType(v string) {
	o.Type = &v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *CommentsPutComment) GetStatus() string {
	if o == nil || IsNil(o.Status) {
		var ret string
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CommentsPutComment) GetStatusOk() (*string, bool) {
	if o == nil || IsNil(o.Status) {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *CommentsPutComment) HasStatus() bool {
	if o != nil && !IsNil(o.Status) {
		return true
	}

	return false
}

// SetStatus gets a reference to the given string and assigns it to the Status field.
func (o *CommentsPutComment) SetStatus(v string) {
	o.Status = &v
}

// GetMessageRaw returns the MessageRaw field value
func (o *CommentsPutComment) GetMessageRaw() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.MessageRaw
}

// GetMessageRawOk returns a tuple with the MessageRaw field value
// and a boolean to check if the value has been set.
func (o *CommentsPutComment) GetMessageRawOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.MessageRaw, true
}

// SetMessageRaw sets field value
func (o *CommentsPutComment) SetMessageRaw(v string) {
	o.MessageRaw = v
}

// GetMessageHtml returns the MessageHtml field value
func (o *CommentsPutComment) GetMessageHtml() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.MessageHtml
}

// GetMessageHtmlOk returns a tuple with the MessageHtml field value
// and a boolean to check if the value has been set.
func (o *CommentsPutComment) GetMessageHtmlOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.MessageHtml, true
}

// SetMessageHtml sets field value
func (o *CommentsPutComment) SetMessageHtml(v string) {
	o.MessageHtml = v
}

// GetCreatedAt returns the CreatedAt field value if set, zero value otherwise.
func (o *CommentsPutComment) GetCreatedAt() string {
	if o == nil || IsNil(o.CreatedAt) {
		var ret string
		return ret
	}
	return *o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CommentsPutComment) GetCreatedAtOk() (*string, bool) {
	if o == nil || IsNil(o.CreatedAt) {
		return nil, false
	}
	return o.CreatedAt, true
}

// HasCreatedAt returns a boolean if a field has been set.
func (o *CommentsPutComment) HasCreatedAt() bool {
	if o != nil && !IsNil(o.CreatedAt) {
		return true
	}

	return false
}

// SetCreatedAt gets a reference to the given string and assigns it to the CreatedAt field.
func (o *CommentsPutComment) SetCreatedAt(v string) {
	o.CreatedAt = &v
}

// GetUpdatedAt returns the UpdatedAt field value if set, zero value otherwise.
func (o *CommentsPutComment) GetUpdatedAt() string {
	if o == nil || IsNil(o.UpdatedAt) {
		var ret string
		return ret
	}
	return *o.UpdatedAt
}

// GetUpdatedAtOk returns a tuple with the UpdatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CommentsPutComment) GetUpdatedAtOk() (*string, bool) {
	if o == nil || IsNil(o.UpdatedAt) {
		return nil, false
	}
	return o.UpdatedAt, true
}

// HasUpdatedAt returns a boolean if a field has been set.
func (o *CommentsPutComment) HasUpdatedAt() bool {
	if o != nil && !IsNil(o.UpdatedAt) {
		return true
	}

	return false
}

// SetUpdatedAt gets a reference to the given string and assigns it to the UpdatedAt field.
func (o *CommentsPutComment) SetUpdatedAt(v string) {
	o.UpdatedAt = &v
}

// GetDeletedAt returns the DeletedAt field value if set, zero value otherwise.
func (o *CommentsPutComment) GetDeletedAt() string {
	if o == nil || IsNil(o.DeletedAt) {
		var ret string
		return ret
	}
	return *o.DeletedAt
}

// GetDeletedAtOk returns a tuple with the DeletedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CommentsPutComment) GetDeletedAtOk() (*string, bool) {
	if o == nil || IsNil(o.DeletedAt) {
		return nil, false
	}
	return o.DeletedAt, true
}

// HasDeletedAt returns a boolean if a field has been set.
func (o *CommentsPutComment) HasDeletedAt() bool {
	if o != nil && !IsNil(o.DeletedAt) {
		return true
	}

	return false
}

// SetDeletedAt gets a reference to the given string and assigns it to the DeletedAt field.
func (o *CommentsPutComment) SetDeletedAt(v string) {
	o.DeletedAt = &v
}

// GetCommentableId returns the CommentableId field value if set, zero value otherwise.
func (o *CommentsPutComment) GetCommentableId() int32 {
	if o == nil || IsNil(o.CommentableId) {
		var ret int32
		return ret
	}
	return *o.CommentableId
}

// GetCommentableIdOk returns a tuple with the CommentableId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CommentsPutComment) GetCommentableIdOk() (*int32, bool) {
	if o == nil || IsNil(o.CommentableId) {
		return nil, false
	}
	return o.CommentableId, true
}

// HasCommentableId returns a boolean if a field has been set.
func (o *CommentsPutComment) HasCommentableId() bool {
	if o != nil && !IsNil(o.CommentableId) {
		return true
	}

	return false
}

// SetCommentableId gets a reference to the given int32 and assigns it to the CommentableId field.
func (o *CommentsPutComment) SetCommentableId(v int32) {
	o.CommentableId = &v
}

// GetCommentableType returns the CommentableType field value if set, zero value otherwise.
func (o *CommentsPutComment) GetCommentableType() string {
	if o == nil || IsNil(o.CommentableType) {
		var ret string
		return ret
	}
	return *o.CommentableType
}

// GetCommentableTypeOk returns a tuple with the CommentableType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CommentsPutComment) GetCommentableTypeOk() (*string, bool) {
	if o == nil || IsNil(o.CommentableType) {
		return nil, false
	}
	return o.CommentableType, true
}

// HasCommentableType returns a boolean if a field has been set.
func (o *CommentsPutComment) HasCommentableType() bool {
	if o != nil && !IsNil(o.CommentableType) {
		return true
	}

	return false
}

// SetCommentableType gets a reference to the given string and assigns it to the CommentableType field.
func (o *CommentsPutComment) SetCommentableType(v string) {
	o.CommentableType = &v
}

// GetReferenceMeta returns the ReferenceMeta field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CommentsPutComment) GetReferenceMeta() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.ReferenceMeta
}

// GetReferenceMetaOk returns a tuple with the ReferenceMeta field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CommentsPutComment) GetReferenceMetaOk() (*interface{}, bool) {
	if o == nil || IsNil(o.ReferenceMeta) {
		return nil, false
	}
	return &o.ReferenceMeta, true
}

// HasReferenceMeta returns a boolean if a field has been set.
func (o *CommentsPutComment) HasReferenceMeta() bool {
	if o != nil && !IsNil(o.ReferenceMeta) {
		return true
	}

	return false
}

// SetReferenceMeta gets a reference to the given interface{} and assigns it to the ReferenceMeta field.
func (o *CommentsPutComment) SetReferenceMeta(v interface{}) {
	o.ReferenceMeta = v
}

// GetOriginAttachmentId returns the OriginAttachmentId field value if set, zero value otherwise.
func (o *CommentsPutComment) GetOriginAttachmentId() int32 {
	if o == nil || IsNil(o.OriginAttachmentId) {
		var ret int32
		return ret
	}
	return *o.OriginAttachmentId
}

// GetOriginAttachmentIdOk returns a tuple with the OriginAttachmentId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CommentsPutComment) GetOriginAttachmentIdOk() (*int32, bool) {
	if o == nil || IsNil(o.OriginAttachmentId) {
		return nil, false
	}
	return o.OriginAttachmentId, true
}

// HasOriginAttachmentId returns a boolean if a field has been set.
func (o *CommentsPutComment) HasOriginAttachmentId() bool {
	if o != nil && !IsNil(o.OriginAttachmentId) {
		return true
	}

	return false
}

// SetOriginAttachmentId gets a reference to the given int32 and assigns it to the OriginAttachmentId field.
func (o *CommentsPutComment) SetOriginAttachmentId(v int32) {
	o.OriginAttachmentId = &v
}

// GetOriginFileId returns the OriginFileId field value if set, zero value otherwise.
func (o *CommentsPutComment) GetOriginFileId() int32 {
	if o == nil || IsNil(o.OriginFileId) {
		var ret int32
		return ret
	}
	return *o.OriginFileId
}

// GetOriginFileIdOk returns a tuple with the OriginFileId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CommentsPutComment) GetOriginFileIdOk() (*int32, bool) {
	if o == nil || IsNil(o.OriginFileId) {
		return nil, false
	}
	return o.OriginFileId, true
}

// HasOriginFileId returns a boolean if a field has been set.
func (o *CommentsPutComment) HasOriginFileId() bool {
	if o != nil && !IsNil(o.OriginFileId) {
		return true
	}

	return false
}

// SetOriginFileId gets a reference to the given int32 and assigns it to the OriginFileId field.
func (o *CommentsPutComment) SetOriginFileId(v int32) {
	o.OriginFileId = &v
}

// GetScopes returns the Scopes field value
// If the value is explicit nil, the zero value for interface{} will be returned
func (o *CommentsPutComment) GetScopes() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}

	return o.Scopes
}

// GetScopesOk returns a tuple with the Scopes field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CommentsPutComment) GetScopesOk() (*interface{}, bool) {
	if o == nil || IsNil(o.Scopes) {
		return nil, false
	}
	return &o.Scopes, true
}

// SetScopes sets field value
func (o *CommentsPutComment) SetScopes(v interface{}) {
	o.Scopes = v
}

// GetExternalIntegrationId returns the ExternalIntegrationId field value
func (o *CommentsPutComment) GetExternalIntegrationId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ExternalIntegrationId
}

// GetExternalIntegrationIdOk returns a tuple with the ExternalIntegrationId field value
// and a boolean to check if the value has been set.
func (o *CommentsPutComment) GetExternalIntegrationIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ExternalIntegrationId, true
}

// SetExternalIntegrationId sets field value
func (o *CommentsPutComment) SetExternalIntegrationId(v string) {
	o.ExternalIntegrationId = v
}

// GetExternalIntegrationAuthor returns the ExternalIntegrationAuthor field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CommentsPutComment) GetExternalIntegrationAuthor() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.ExternalIntegrationAuthor
}

// GetExternalIntegrationAuthorOk returns a tuple with the ExternalIntegrationAuthor field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CommentsPutComment) GetExternalIntegrationAuthorOk() (*interface{}, bool) {
	if o == nil || IsNil(o.ExternalIntegrationAuthor) {
		return nil, false
	}
	return &o.ExternalIntegrationAuthor, true
}

// HasExternalIntegrationAuthor returns a boolean if a field has been set.
func (o *CommentsPutComment) HasExternalIntegrationAuthor() bool {
	if o != nil && !IsNil(o.ExternalIntegrationAuthor) {
		return true
	}

	return false
}

// SetExternalIntegrationAuthor gets a reference to the given interface{} and assigns it to the ExternalIntegrationAuthor field.
func (o *CommentsPutComment) SetExternalIntegrationAuthor(v interface{}) {
	o.ExternalIntegrationAuthor = v
}

// GetOriginatingPath returns the OriginatingPath field value if set, zero value otherwise.
func (o *CommentsPutComment) GetOriginatingPath() string {
	if o == nil || IsNil(o.OriginatingPath) {
		var ret string
		return ret
	}
	return *o.OriginatingPath
}

// GetOriginatingPathOk returns a tuple with the OriginatingPath field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CommentsPutComment) GetOriginatingPathOk() (*string, bool) {
	if o == nil || IsNil(o.OriginatingPath) {
		return nil, false
	}
	return o.OriginatingPath, true
}

// HasOriginatingPath returns a boolean if a field has been set.
func (o *CommentsPutComment) HasOriginatingPath() bool {
	if o != nil && !IsNil(o.OriginatingPath) {
		return true
	}

	return false
}

// SetOriginatingPath gets a reference to the given string and assigns it to the OriginatingPath field.
func (o *CommentsPutComment) SetOriginatingPath(v string) {
	o.OriginatingPath = &v
}

func (o CommentsPutComment) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CommentsPutComment) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !IsNil(o.ParentCommentId) {
		toSerialize["parent_comment_id"] = o.ParentCommentId
	}
	if !IsNil(o.OwnerUserId) {
		toSerialize["owner_user_id"] = o.OwnerUserId
	}
	if !IsNil(o.Type) {
		toSerialize["type"] = o.Type
	}
	if !IsNil(o.Status) {
		toSerialize["status"] = o.Status
	}
	toSerialize["message_raw"] = o.MessageRaw
	toSerialize["message_html"] = o.MessageHtml
	if !IsNil(o.CreatedAt) {
		toSerialize["created_at"] = o.CreatedAt
	}
	if !IsNil(o.UpdatedAt) {
		toSerialize["updated_at"] = o.UpdatedAt
	}
	if !IsNil(o.DeletedAt) {
		toSerialize["deleted_at"] = o.DeletedAt
	}
	if !IsNil(o.CommentableId) {
		toSerialize["commentable_id"] = o.CommentableId
	}
	if !IsNil(o.CommentableType) {
		toSerialize["commentable_type"] = o.CommentableType
	}
	if o.ReferenceMeta != nil {
		toSerialize["reference_meta"] = o.ReferenceMeta
	}
	if !IsNil(o.OriginAttachmentId) {
		toSerialize["origin_attachment_id"] = o.OriginAttachmentId
	}
	if !IsNil(o.OriginFileId) {
		toSerialize["origin_file_id"] = o.OriginFileId
	}
	if o.Scopes != nil {
		toSerialize["scopes"] = o.Scopes
	}
	toSerialize["external_integration_id"] = o.ExternalIntegrationId
	if o.ExternalIntegrationAuthor != nil {
		toSerialize["external_integration_author"] = o.ExternalIntegrationAuthor
	}
	if !IsNil(o.OriginatingPath) {
		toSerialize["originating_path"] = o.OriginatingPath
	}
	return toSerialize, nil
}

func (o *CommentsPutComment) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"message_raw",
		"message_html",
		"scopes",
		"external_integration_id",
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

	varCommentsPutComment := _CommentsPutComment{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varCommentsPutComment)

	if err != nil {
		return err
	}

	*o = CommentsPutComment(varCommentsPutComment)

	return err
}

type NullableCommentsPutComment struct {
	value *CommentsPutComment
	isSet bool
}

func (v NullableCommentsPutComment) Get() *CommentsPutComment {
	return v.value
}

func (v *NullableCommentsPutComment) Set(val *CommentsPutComment) {
	v.value = val
	v.isSet = true
}

func (v NullableCommentsPutComment) IsSet() bool {
	return v.isSet
}

func (v *NullableCommentsPutComment) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCommentsPutComment(val *CommentsPutComment) *NullableCommentsPutComment {
	return &NullableCommentsPutComment{value: val, isSet: true}
}

func (v NullableCommentsPutComment) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCommentsPutComment) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

