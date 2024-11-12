# AuditableEntityTypesPutAuditableEntityType

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int32** | Note: This is a Primary Key.&lt;pk/&gt; | [optional] [readonly] 
**SortOrder** | Pointer to **int32** |  | [optional] 
**CreatedAt** | Pointer to **string** |  | [optional] [readonly] 
**UpdatedAt** | Pointer to **string** |  | [optional] [readonly] 
**DeletedAt** | Pointer to **string** |  | [optional] [readonly] 
**Name** | **string** |  | 
**FormTemplateId** | Pointer to **int32** |  | [optional] 
**AllowedSections** | **interface{}** |  | 
**EnabledAttributes** | Pointer to **interface{}** |  | [optional] 
**ExcludedAttributes** | Pointer to **interface{}** |  | [optional] 
**Key** | Pointer to **string** |  | [optional] 
**IntakeFormTemplateId** | Pointer to **int32** | Note: This is a Foreign Key to &#x60;audit_survey_templates.id&#x60;.&lt;fk table&#x3D;&#39;audit_survey_templates&#39; column&#x3D;&#39;id&#39;/&gt; | [optional] 
**Level** | Pointer to **int32** |  | [optional] 
**InventoryClass** | Pointer to **string** |  | [optional] 

## Methods

### NewAuditableEntityTypesPutAuditableEntityType

`func NewAuditableEntityTypesPutAuditableEntityType(name string, allowedSections interface{}, ) *AuditableEntityTypesPutAuditableEntityType`

NewAuditableEntityTypesPutAuditableEntityType instantiates a new AuditableEntityTypesPutAuditableEntityType object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAuditableEntityTypesPutAuditableEntityTypeWithDefaults

`func NewAuditableEntityTypesPutAuditableEntityTypeWithDefaults() *AuditableEntityTypesPutAuditableEntityType`

NewAuditableEntityTypesPutAuditableEntityTypeWithDefaults instantiates a new AuditableEntityTypesPutAuditableEntityType object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *AuditableEntityTypesPutAuditableEntityType) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *AuditableEntityTypesPutAuditableEntityType) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *AuditableEntityTypesPutAuditableEntityType) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *AuditableEntityTypesPutAuditableEntityType) HasId() bool`

HasId returns a boolean if a field has been set.

### GetSortOrder

`func (o *AuditableEntityTypesPutAuditableEntityType) GetSortOrder() int32`

GetSortOrder returns the SortOrder field if non-nil, zero value otherwise.

### GetSortOrderOk

`func (o *AuditableEntityTypesPutAuditableEntityType) GetSortOrderOk() (*int32, bool)`

GetSortOrderOk returns a tuple with the SortOrder field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSortOrder

`func (o *AuditableEntityTypesPutAuditableEntityType) SetSortOrder(v int32)`

SetSortOrder sets SortOrder field to given value.

### HasSortOrder

`func (o *AuditableEntityTypesPutAuditableEntityType) HasSortOrder() bool`

HasSortOrder returns a boolean if a field has been set.

### GetCreatedAt

`func (o *AuditableEntityTypesPutAuditableEntityType) GetCreatedAt() string`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *AuditableEntityTypesPutAuditableEntityType) GetCreatedAtOk() (*string, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *AuditableEntityTypesPutAuditableEntityType) SetCreatedAt(v string)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *AuditableEntityTypesPutAuditableEntityType) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *AuditableEntityTypesPutAuditableEntityType) GetUpdatedAt() string`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *AuditableEntityTypesPutAuditableEntityType) GetUpdatedAtOk() (*string, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *AuditableEntityTypesPutAuditableEntityType) SetUpdatedAt(v string)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *AuditableEntityTypesPutAuditableEntityType) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.

### GetDeletedAt

`func (o *AuditableEntityTypesPutAuditableEntityType) GetDeletedAt() string`

GetDeletedAt returns the DeletedAt field if non-nil, zero value otherwise.

### GetDeletedAtOk

`func (o *AuditableEntityTypesPutAuditableEntityType) GetDeletedAtOk() (*string, bool)`

GetDeletedAtOk returns a tuple with the DeletedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeletedAt

`func (o *AuditableEntityTypesPutAuditableEntityType) SetDeletedAt(v string)`

SetDeletedAt sets DeletedAt field to given value.

### HasDeletedAt

`func (o *AuditableEntityTypesPutAuditableEntityType) HasDeletedAt() bool`

HasDeletedAt returns a boolean if a field has been set.

### GetName

`func (o *AuditableEntityTypesPutAuditableEntityType) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *AuditableEntityTypesPutAuditableEntityType) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *AuditableEntityTypesPutAuditableEntityType) SetName(v string)`

SetName sets Name field to given value.


### GetFormTemplateId

`func (o *AuditableEntityTypesPutAuditableEntityType) GetFormTemplateId() int32`

GetFormTemplateId returns the FormTemplateId field if non-nil, zero value otherwise.

### GetFormTemplateIdOk

`func (o *AuditableEntityTypesPutAuditableEntityType) GetFormTemplateIdOk() (*int32, bool)`

GetFormTemplateIdOk returns a tuple with the FormTemplateId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFormTemplateId

`func (o *AuditableEntityTypesPutAuditableEntityType) SetFormTemplateId(v int32)`

SetFormTemplateId sets FormTemplateId field to given value.

### HasFormTemplateId

`func (o *AuditableEntityTypesPutAuditableEntityType) HasFormTemplateId() bool`

HasFormTemplateId returns a boolean if a field has been set.

### GetAllowedSections

`func (o *AuditableEntityTypesPutAuditableEntityType) GetAllowedSections() interface{}`

GetAllowedSections returns the AllowedSections field if non-nil, zero value otherwise.

### GetAllowedSectionsOk

`func (o *AuditableEntityTypesPutAuditableEntityType) GetAllowedSectionsOk() (*interface{}, bool)`

GetAllowedSectionsOk returns a tuple with the AllowedSections field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowedSections

`func (o *AuditableEntityTypesPutAuditableEntityType) SetAllowedSections(v interface{})`

SetAllowedSections sets AllowedSections field to given value.


### SetAllowedSectionsNil

`func (o *AuditableEntityTypesPutAuditableEntityType) SetAllowedSectionsNil(b bool)`

 SetAllowedSectionsNil sets the value for AllowedSections to be an explicit nil

### UnsetAllowedSections
`func (o *AuditableEntityTypesPutAuditableEntityType) UnsetAllowedSections()`

UnsetAllowedSections ensures that no value is present for AllowedSections, not even an explicit nil
### GetEnabledAttributes

`func (o *AuditableEntityTypesPutAuditableEntityType) GetEnabledAttributes() interface{}`

GetEnabledAttributes returns the EnabledAttributes field if non-nil, zero value otherwise.

### GetEnabledAttributesOk

`func (o *AuditableEntityTypesPutAuditableEntityType) GetEnabledAttributesOk() (*interface{}, bool)`

GetEnabledAttributesOk returns a tuple with the EnabledAttributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabledAttributes

`func (o *AuditableEntityTypesPutAuditableEntityType) SetEnabledAttributes(v interface{})`

SetEnabledAttributes sets EnabledAttributes field to given value.

### HasEnabledAttributes

`func (o *AuditableEntityTypesPutAuditableEntityType) HasEnabledAttributes() bool`

HasEnabledAttributes returns a boolean if a field has been set.

### SetEnabledAttributesNil

`func (o *AuditableEntityTypesPutAuditableEntityType) SetEnabledAttributesNil(b bool)`

 SetEnabledAttributesNil sets the value for EnabledAttributes to be an explicit nil

### UnsetEnabledAttributes
`func (o *AuditableEntityTypesPutAuditableEntityType) UnsetEnabledAttributes()`

UnsetEnabledAttributes ensures that no value is present for EnabledAttributes, not even an explicit nil
### GetExcludedAttributes

`func (o *AuditableEntityTypesPutAuditableEntityType) GetExcludedAttributes() interface{}`

GetExcludedAttributes returns the ExcludedAttributes field if non-nil, zero value otherwise.

### GetExcludedAttributesOk

`func (o *AuditableEntityTypesPutAuditableEntityType) GetExcludedAttributesOk() (*interface{}, bool)`

GetExcludedAttributesOk returns a tuple with the ExcludedAttributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExcludedAttributes

`func (o *AuditableEntityTypesPutAuditableEntityType) SetExcludedAttributes(v interface{})`

SetExcludedAttributes sets ExcludedAttributes field to given value.

### HasExcludedAttributes

`func (o *AuditableEntityTypesPutAuditableEntityType) HasExcludedAttributes() bool`

HasExcludedAttributes returns a boolean if a field has been set.

### SetExcludedAttributesNil

`func (o *AuditableEntityTypesPutAuditableEntityType) SetExcludedAttributesNil(b bool)`

 SetExcludedAttributesNil sets the value for ExcludedAttributes to be an explicit nil

### UnsetExcludedAttributes
`func (o *AuditableEntityTypesPutAuditableEntityType) UnsetExcludedAttributes()`

UnsetExcludedAttributes ensures that no value is present for ExcludedAttributes, not even an explicit nil
### GetKey

`func (o *AuditableEntityTypesPutAuditableEntityType) GetKey() string`

GetKey returns the Key field if non-nil, zero value otherwise.

### GetKeyOk

`func (o *AuditableEntityTypesPutAuditableEntityType) GetKeyOk() (*string, bool)`

GetKeyOk returns a tuple with the Key field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKey

`func (o *AuditableEntityTypesPutAuditableEntityType) SetKey(v string)`

SetKey sets Key field to given value.

### HasKey

`func (o *AuditableEntityTypesPutAuditableEntityType) HasKey() bool`

HasKey returns a boolean if a field has been set.

### GetIntakeFormTemplateId

`func (o *AuditableEntityTypesPutAuditableEntityType) GetIntakeFormTemplateId() int32`

GetIntakeFormTemplateId returns the IntakeFormTemplateId field if non-nil, zero value otherwise.

### GetIntakeFormTemplateIdOk

`func (o *AuditableEntityTypesPutAuditableEntityType) GetIntakeFormTemplateIdOk() (*int32, bool)`

GetIntakeFormTemplateIdOk returns a tuple with the IntakeFormTemplateId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIntakeFormTemplateId

`func (o *AuditableEntityTypesPutAuditableEntityType) SetIntakeFormTemplateId(v int32)`

SetIntakeFormTemplateId sets IntakeFormTemplateId field to given value.

### HasIntakeFormTemplateId

`func (o *AuditableEntityTypesPutAuditableEntityType) HasIntakeFormTemplateId() bool`

HasIntakeFormTemplateId returns a boolean if a field has been set.

### GetLevel

`func (o *AuditableEntityTypesPutAuditableEntityType) GetLevel() int32`

GetLevel returns the Level field if non-nil, zero value otherwise.

### GetLevelOk

`func (o *AuditableEntityTypesPutAuditableEntityType) GetLevelOk() (*int32, bool)`

GetLevelOk returns a tuple with the Level field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLevel

`func (o *AuditableEntityTypesPutAuditableEntityType) SetLevel(v int32)`

SetLevel sets Level field to given value.

### HasLevel

`func (o *AuditableEntityTypesPutAuditableEntityType) HasLevel() bool`

HasLevel returns a boolean if a field has been set.

### GetInventoryClass

`func (o *AuditableEntityTypesPutAuditableEntityType) GetInventoryClass() string`

GetInventoryClass returns the InventoryClass field if non-nil, zero value otherwise.

### GetInventoryClassOk

`func (o *AuditableEntityTypesPutAuditableEntityType) GetInventoryClassOk() (*string, bool)`

GetInventoryClassOk returns a tuple with the InventoryClass field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInventoryClass

`func (o *AuditableEntityTypesPutAuditableEntityType) SetInventoryClass(v string)`

SetInventoryClass sets InventoryClass field to given value.

### HasInventoryClass

`func (o *AuditableEntityTypesPutAuditableEntityType) HasInventoryClass() bool`

HasInventoryClass returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


