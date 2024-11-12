# AuditableEntityTypesPutPreviousValues

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int32** | Note: This is a Primary Key.&lt;pk/&gt; | [optional] [readonly] 
**SortOrder** | Pointer to **int32** |  | [optional] 
**CreatedAt** | Pointer to **string** |  | [optional] [readonly] 
**UpdatedAt** | Pointer to **string** |  | [optional] [readonly] 
**DeletedAt** | Pointer to **string** |  | [optional] [readonly] 
**Name** | Pointer to **string** |  | [optional] 
**FormTemplateId** | Pointer to **int32** |  | [optional] 
**AllowedSections** | Pointer to **interface{}** |  | [optional] 
**EnabledAttributes** | Pointer to **interface{}** |  | [optional] 
**ExcludedAttributes** | Pointer to **interface{}** |  | [optional] 
**Key** | Pointer to **string** |  | [optional] 
**IntakeFormTemplateId** | Pointer to **int32** | Note: This is a Foreign Key to &#x60;audit_survey_templates.id&#x60;.&lt;fk table&#x3D;&#39;audit_survey_templates&#39; column&#x3D;&#39;id&#39;/&gt; | [optional] 
**Level** | Pointer to **int32** |  | [optional] 
**InventoryClass** | Pointer to **string** |  | [optional] 

## Methods

### NewAuditableEntityTypesPutPreviousValues

`func NewAuditableEntityTypesPutPreviousValues() *AuditableEntityTypesPutPreviousValues`

NewAuditableEntityTypesPutPreviousValues instantiates a new AuditableEntityTypesPutPreviousValues object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAuditableEntityTypesPutPreviousValuesWithDefaults

`func NewAuditableEntityTypesPutPreviousValuesWithDefaults() *AuditableEntityTypesPutPreviousValues`

NewAuditableEntityTypesPutPreviousValuesWithDefaults instantiates a new AuditableEntityTypesPutPreviousValues object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *AuditableEntityTypesPutPreviousValues) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *AuditableEntityTypesPutPreviousValues) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *AuditableEntityTypesPutPreviousValues) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *AuditableEntityTypesPutPreviousValues) HasId() bool`

HasId returns a boolean if a field has been set.

### GetSortOrder

`func (o *AuditableEntityTypesPutPreviousValues) GetSortOrder() int32`

GetSortOrder returns the SortOrder field if non-nil, zero value otherwise.

### GetSortOrderOk

`func (o *AuditableEntityTypesPutPreviousValues) GetSortOrderOk() (*int32, bool)`

GetSortOrderOk returns a tuple with the SortOrder field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSortOrder

`func (o *AuditableEntityTypesPutPreviousValues) SetSortOrder(v int32)`

SetSortOrder sets SortOrder field to given value.

### HasSortOrder

`func (o *AuditableEntityTypesPutPreviousValues) HasSortOrder() bool`

HasSortOrder returns a boolean if a field has been set.

### GetCreatedAt

`func (o *AuditableEntityTypesPutPreviousValues) GetCreatedAt() string`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *AuditableEntityTypesPutPreviousValues) GetCreatedAtOk() (*string, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *AuditableEntityTypesPutPreviousValues) SetCreatedAt(v string)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *AuditableEntityTypesPutPreviousValues) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *AuditableEntityTypesPutPreviousValues) GetUpdatedAt() string`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *AuditableEntityTypesPutPreviousValues) GetUpdatedAtOk() (*string, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *AuditableEntityTypesPutPreviousValues) SetUpdatedAt(v string)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *AuditableEntityTypesPutPreviousValues) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.

### GetDeletedAt

`func (o *AuditableEntityTypesPutPreviousValues) GetDeletedAt() string`

GetDeletedAt returns the DeletedAt field if non-nil, zero value otherwise.

### GetDeletedAtOk

`func (o *AuditableEntityTypesPutPreviousValues) GetDeletedAtOk() (*string, bool)`

GetDeletedAtOk returns a tuple with the DeletedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeletedAt

`func (o *AuditableEntityTypesPutPreviousValues) SetDeletedAt(v string)`

SetDeletedAt sets DeletedAt field to given value.

### HasDeletedAt

`func (o *AuditableEntityTypesPutPreviousValues) HasDeletedAt() bool`

HasDeletedAt returns a boolean if a field has been set.

### GetName

`func (o *AuditableEntityTypesPutPreviousValues) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *AuditableEntityTypesPutPreviousValues) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *AuditableEntityTypesPutPreviousValues) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *AuditableEntityTypesPutPreviousValues) HasName() bool`

HasName returns a boolean if a field has been set.

### GetFormTemplateId

`func (o *AuditableEntityTypesPutPreviousValues) GetFormTemplateId() int32`

GetFormTemplateId returns the FormTemplateId field if non-nil, zero value otherwise.

### GetFormTemplateIdOk

`func (o *AuditableEntityTypesPutPreviousValues) GetFormTemplateIdOk() (*int32, bool)`

GetFormTemplateIdOk returns a tuple with the FormTemplateId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFormTemplateId

`func (o *AuditableEntityTypesPutPreviousValues) SetFormTemplateId(v int32)`

SetFormTemplateId sets FormTemplateId field to given value.

### HasFormTemplateId

`func (o *AuditableEntityTypesPutPreviousValues) HasFormTemplateId() bool`

HasFormTemplateId returns a boolean if a field has been set.

### GetAllowedSections

`func (o *AuditableEntityTypesPutPreviousValues) GetAllowedSections() interface{}`

GetAllowedSections returns the AllowedSections field if non-nil, zero value otherwise.

### GetAllowedSectionsOk

`func (o *AuditableEntityTypesPutPreviousValues) GetAllowedSectionsOk() (*interface{}, bool)`

GetAllowedSectionsOk returns a tuple with the AllowedSections field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowedSections

`func (o *AuditableEntityTypesPutPreviousValues) SetAllowedSections(v interface{})`

SetAllowedSections sets AllowedSections field to given value.

### HasAllowedSections

`func (o *AuditableEntityTypesPutPreviousValues) HasAllowedSections() bool`

HasAllowedSections returns a boolean if a field has been set.

### SetAllowedSectionsNil

`func (o *AuditableEntityTypesPutPreviousValues) SetAllowedSectionsNil(b bool)`

 SetAllowedSectionsNil sets the value for AllowedSections to be an explicit nil

### UnsetAllowedSections
`func (o *AuditableEntityTypesPutPreviousValues) UnsetAllowedSections()`

UnsetAllowedSections ensures that no value is present for AllowedSections, not even an explicit nil
### GetEnabledAttributes

`func (o *AuditableEntityTypesPutPreviousValues) GetEnabledAttributes() interface{}`

GetEnabledAttributes returns the EnabledAttributes field if non-nil, zero value otherwise.

### GetEnabledAttributesOk

`func (o *AuditableEntityTypesPutPreviousValues) GetEnabledAttributesOk() (*interface{}, bool)`

GetEnabledAttributesOk returns a tuple with the EnabledAttributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabledAttributes

`func (o *AuditableEntityTypesPutPreviousValues) SetEnabledAttributes(v interface{})`

SetEnabledAttributes sets EnabledAttributes field to given value.

### HasEnabledAttributes

`func (o *AuditableEntityTypesPutPreviousValues) HasEnabledAttributes() bool`

HasEnabledAttributes returns a boolean if a field has been set.

### SetEnabledAttributesNil

`func (o *AuditableEntityTypesPutPreviousValues) SetEnabledAttributesNil(b bool)`

 SetEnabledAttributesNil sets the value for EnabledAttributes to be an explicit nil

### UnsetEnabledAttributes
`func (o *AuditableEntityTypesPutPreviousValues) UnsetEnabledAttributes()`

UnsetEnabledAttributes ensures that no value is present for EnabledAttributes, not even an explicit nil
### GetExcludedAttributes

`func (o *AuditableEntityTypesPutPreviousValues) GetExcludedAttributes() interface{}`

GetExcludedAttributes returns the ExcludedAttributes field if non-nil, zero value otherwise.

### GetExcludedAttributesOk

`func (o *AuditableEntityTypesPutPreviousValues) GetExcludedAttributesOk() (*interface{}, bool)`

GetExcludedAttributesOk returns a tuple with the ExcludedAttributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExcludedAttributes

`func (o *AuditableEntityTypesPutPreviousValues) SetExcludedAttributes(v interface{})`

SetExcludedAttributes sets ExcludedAttributes field to given value.

### HasExcludedAttributes

`func (o *AuditableEntityTypesPutPreviousValues) HasExcludedAttributes() bool`

HasExcludedAttributes returns a boolean if a field has been set.

### SetExcludedAttributesNil

`func (o *AuditableEntityTypesPutPreviousValues) SetExcludedAttributesNil(b bool)`

 SetExcludedAttributesNil sets the value for ExcludedAttributes to be an explicit nil

### UnsetExcludedAttributes
`func (o *AuditableEntityTypesPutPreviousValues) UnsetExcludedAttributes()`

UnsetExcludedAttributes ensures that no value is present for ExcludedAttributes, not even an explicit nil
### GetKey

`func (o *AuditableEntityTypesPutPreviousValues) GetKey() string`

GetKey returns the Key field if non-nil, zero value otherwise.

### GetKeyOk

`func (o *AuditableEntityTypesPutPreviousValues) GetKeyOk() (*string, bool)`

GetKeyOk returns a tuple with the Key field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKey

`func (o *AuditableEntityTypesPutPreviousValues) SetKey(v string)`

SetKey sets Key field to given value.

### HasKey

`func (o *AuditableEntityTypesPutPreviousValues) HasKey() bool`

HasKey returns a boolean if a field has been set.

### GetIntakeFormTemplateId

`func (o *AuditableEntityTypesPutPreviousValues) GetIntakeFormTemplateId() int32`

GetIntakeFormTemplateId returns the IntakeFormTemplateId field if non-nil, zero value otherwise.

### GetIntakeFormTemplateIdOk

`func (o *AuditableEntityTypesPutPreviousValues) GetIntakeFormTemplateIdOk() (*int32, bool)`

GetIntakeFormTemplateIdOk returns a tuple with the IntakeFormTemplateId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIntakeFormTemplateId

`func (o *AuditableEntityTypesPutPreviousValues) SetIntakeFormTemplateId(v int32)`

SetIntakeFormTemplateId sets IntakeFormTemplateId field to given value.

### HasIntakeFormTemplateId

`func (o *AuditableEntityTypesPutPreviousValues) HasIntakeFormTemplateId() bool`

HasIntakeFormTemplateId returns a boolean if a field has been set.

### GetLevel

`func (o *AuditableEntityTypesPutPreviousValues) GetLevel() int32`

GetLevel returns the Level field if non-nil, zero value otherwise.

### GetLevelOk

`func (o *AuditableEntityTypesPutPreviousValues) GetLevelOk() (*int32, bool)`

GetLevelOk returns a tuple with the Level field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLevel

`func (o *AuditableEntityTypesPutPreviousValues) SetLevel(v int32)`

SetLevel sets Level field to given value.

### HasLevel

`func (o *AuditableEntityTypesPutPreviousValues) HasLevel() bool`

HasLevel returns a boolean if a field has been set.

### GetInventoryClass

`func (o *AuditableEntityTypesPutPreviousValues) GetInventoryClass() string`

GetInventoryClass returns the InventoryClass field if non-nil, zero value otherwise.

### GetInventoryClassOk

`func (o *AuditableEntityTypesPutPreviousValues) GetInventoryClassOk() (*string, bool)`

GetInventoryClassOk returns a tuple with the InventoryClass field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInventoryClass

`func (o *AuditableEntityTypesPutPreviousValues) SetInventoryClass(v string)`

SetInventoryClass sets InventoryClass field to given value.

### HasInventoryClass

`func (o *AuditableEntityTypesPutPreviousValues) HasInventoryClass() bool`

HasInventoryClass returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


