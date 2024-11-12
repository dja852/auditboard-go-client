# ControlsPutPreviousValues

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int32** | Note: This is a Primary Key.&lt;pk/&gt; | [optional] [readonly] 
**CreatedAt** | Pointer to **string** |  | [optional] [readonly] 
**UpdatedAt** | Pointer to **string** |  | [optional] [readonly] 
**DeletedAt** | Pointer to **string** |  | [optional] [readonly] 
**Uid** | Pointer to **string** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**SubprocessId** | Pointer to **int32** | Note: This is a Foreign Key to &#x60;subprocesses.id&#x60;.&lt;fk table&#x3D;&#39;subprocesses&#39; column&#x3D;&#39;id&#39;/&gt; | [optional] 
**Description** | Pointer to **string** |  | [optional] [default to ""]
**IdCode** | Pointer to **string** |  | [optional] 
**FieldData** | Pointer to **interface{}** |  | [optional] 
**ComplianceControlCount** | Pointer to **int32** |  | [optional] [default to 0]
**ControlObjective** | Pointer to **string** |  | [optional] 
**RiskStatement** | Pointer to **string** |  | [optional] 
**LibraryControlClassificationId** | Pointer to **int32** | Note: This is a Foreign Key to &#x60;library_control_classifications.id&#x60;.&lt;fk table&#x3D;&#39;library_control_classifications&#39; column&#x3D;&#39;id&#39;/&gt; | [optional] 
**LibraryControlNatureId** | Pointer to **int32** | Note: This is a Foreign Key to &#x60;library_control_natures.id&#x60;.&lt;fk table&#x3D;&#39;library_control_natures&#39; column&#x3D;&#39;id&#39;/&gt; | [optional] 
**LibraryControlTypeId** | Pointer to **int32** | Note: This is a Foreign Key to &#x60;library_control_types.id&#x60;.&lt;fk table&#x3D;&#39;library_control_types&#39; column&#x3D;&#39;id&#39;/&gt; | [optional] 
**ControlCategoryId** | Pointer to **int32** | Note: This is a Foreign Key to &#x60;control_categories.id&#x60;.&lt;fk table&#x3D;&#39;control_categories&#39; column&#x3D;&#39;id&#39;/&gt; | [optional] 

## Methods

### NewControlsPutPreviousValues

`func NewControlsPutPreviousValues() *ControlsPutPreviousValues`

NewControlsPutPreviousValues instantiates a new ControlsPutPreviousValues object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewControlsPutPreviousValuesWithDefaults

`func NewControlsPutPreviousValuesWithDefaults() *ControlsPutPreviousValues`

NewControlsPutPreviousValuesWithDefaults instantiates a new ControlsPutPreviousValues object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ControlsPutPreviousValues) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ControlsPutPreviousValues) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ControlsPutPreviousValues) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *ControlsPutPreviousValues) HasId() bool`

HasId returns a boolean if a field has been set.

### GetCreatedAt

`func (o *ControlsPutPreviousValues) GetCreatedAt() string`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *ControlsPutPreviousValues) GetCreatedAtOk() (*string, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *ControlsPutPreviousValues) SetCreatedAt(v string)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *ControlsPutPreviousValues) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *ControlsPutPreviousValues) GetUpdatedAt() string`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *ControlsPutPreviousValues) GetUpdatedAtOk() (*string, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *ControlsPutPreviousValues) SetUpdatedAt(v string)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *ControlsPutPreviousValues) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.

### GetDeletedAt

`func (o *ControlsPutPreviousValues) GetDeletedAt() string`

GetDeletedAt returns the DeletedAt field if non-nil, zero value otherwise.

### GetDeletedAtOk

`func (o *ControlsPutPreviousValues) GetDeletedAtOk() (*string, bool)`

GetDeletedAtOk returns a tuple with the DeletedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeletedAt

`func (o *ControlsPutPreviousValues) SetDeletedAt(v string)`

SetDeletedAt sets DeletedAt field to given value.

### HasDeletedAt

`func (o *ControlsPutPreviousValues) HasDeletedAt() bool`

HasDeletedAt returns a boolean if a field has been set.

### GetUid

`func (o *ControlsPutPreviousValues) GetUid() string`

GetUid returns the Uid field if non-nil, zero value otherwise.

### GetUidOk

`func (o *ControlsPutPreviousValues) GetUidOk() (*string, bool)`

GetUidOk returns a tuple with the Uid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUid

`func (o *ControlsPutPreviousValues) SetUid(v string)`

SetUid sets Uid field to given value.

### HasUid

`func (o *ControlsPutPreviousValues) HasUid() bool`

HasUid returns a boolean if a field has been set.

### GetName

`func (o *ControlsPutPreviousValues) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ControlsPutPreviousValues) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ControlsPutPreviousValues) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *ControlsPutPreviousValues) HasName() bool`

HasName returns a boolean if a field has been set.

### GetSubprocessId

`func (o *ControlsPutPreviousValues) GetSubprocessId() int32`

GetSubprocessId returns the SubprocessId field if non-nil, zero value otherwise.

### GetSubprocessIdOk

`func (o *ControlsPutPreviousValues) GetSubprocessIdOk() (*int32, bool)`

GetSubprocessIdOk returns a tuple with the SubprocessId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubprocessId

`func (o *ControlsPutPreviousValues) SetSubprocessId(v int32)`

SetSubprocessId sets SubprocessId field to given value.

### HasSubprocessId

`func (o *ControlsPutPreviousValues) HasSubprocessId() bool`

HasSubprocessId returns a boolean if a field has been set.

### GetDescription

`func (o *ControlsPutPreviousValues) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *ControlsPutPreviousValues) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *ControlsPutPreviousValues) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *ControlsPutPreviousValues) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetIdCode

`func (o *ControlsPutPreviousValues) GetIdCode() string`

GetIdCode returns the IdCode field if non-nil, zero value otherwise.

### GetIdCodeOk

`func (o *ControlsPutPreviousValues) GetIdCodeOk() (*string, bool)`

GetIdCodeOk returns a tuple with the IdCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdCode

`func (o *ControlsPutPreviousValues) SetIdCode(v string)`

SetIdCode sets IdCode field to given value.

### HasIdCode

`func (o *ControlsPutPreviousValues) HasIdCode() bool`

HasIdCode returns a boolean if a field has been set.

### GetFieldData

`func (o *ControlsPutPreviousValues) GetFieldData() interface{}`

GetFieldData returns the FieldData field if non-nil, zero value otherwise.

### GetFieldDataOk

`func (o *ControlsPutPreviousValues) GetFieldDataOk() (*interface{}, bool)`

GetFieldDataOk returns a tuple with the FieldData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFieldData

`func (o *ControlsPutPreviousValues) SetFieldData(v interface{})`

SetFieldData sets FieldData field to given value.

### HasFieldData

`func (o *ControlsPutPreviousValues) HasFieldData() bool`

HasFieldData returns a boolean if a field has been set.

### SetFieldDataNil

`func (o *ControlsPutPreviousValues) SetFieldDataNil(b bool)`

 SetFieldDataNil sets the value for FieldData to be an explicit nil

### UnsetFieldData
`func (o *ControlsPutPreviousValues) UnsetFieldData()`

UnsetFieldData ensures that no value is present for FieldData, not even an explicit nil
### GetComplianceControlCount

`func (o *ControlsPutPreviousValues) GetComplianceControlCount() int32`

GetComplianceControlCount returns the ComplianceControlCount field if non-nil, zero value otherwise.

### GetComplianceControlCountOk

`func (o *ControlsPutPreviousValues) GetComplianceControlCountOk() (*int32, bool)`

GetComplianceControlCountOk returns a tuple with the ComplianceControlCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComplianceControlCount

`func (o *ControlsPutPreviousValues) SetComplianceControlCount(v int32)`

SetComplianceControlCount sets ComplianceControlCount field to given value.

### HasComplianceControlCount

`func (o *ControlsPutPreviousValues) HasComplianceControlCount() bool`

HasComplianceControlCount returns a boolean if a field has been set.

### GetControlObjective

`func (o *ControlsPutPreviousValues) GetControlObjective() string`

GetControlObjective returns the ControlObjective field if non-nil, zero value otherwise.

### GetControlObjectiveOk

`func (o *ControlsPutPreviousValues) GetControlObjectiveOk() (*string, bool)`

GetControlObjectiveOk returns a tuple with the ControlObjective field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetControlObjective

`func (o *ControlsPutPreviousValues) SetControlObjective(v string)`

SetControlObjective sets ControlObjective field to given value.

### HasControlObjective

`func (o *ControlsPutPreviousValues) HasControlObjective() bool`

HasControlObjective returns a boolean if a field has been set.

### GetRiskStatement

`func (o *ControlsPutPreviousValues) GetRiskStatement() string`

GetRiskStatement returns the RiskStatement field if non-nil, zero value otherwise.

### GetRiskStatementOk

`func (o *ControlsPutPreviousValues) GetRiskStatementOk() (*string, bool)`

GetRiskStatementOk returns a tuple with the RiskStatement field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRiskStatement

`func (o *ControlsPutPreviousValues) SetRiskStatement(v string)`

SetRiskStatement sets RiskStatement field to given value.

### HasRiskStatement

`func (o *ControlsPutPreviousValues) HasRiskStatement() bool`

HasRiskStatement returns a boolean if a field has been set.

### GetLibraryControlClassificationId

`func (o *ControlsPutPreviousValues) GetLibraryControlClassificationId() int32`

GetLibraryControlClassificationId returns the LibraryControlClassificationId field if non-nil, zero value otherwise.

### GetLibraryControlClassificationIdOk

`func (o *ControlsPutPreviousValues) GetLibraryControlClassificationIdOk() (*int32, bool)`

GetLibraryControlClassificationIdOk returns a tuple with the LibraryControlClassificationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLibraryControlClassificationId

`func (o *ControlsPutPreviousValues) SetLibraryControlClassificationId(v int32)`

SetLibraryControlClassificationId sets LibraryControlClassificationId field to given value.

### HasLibraryControlClassificationId

`func (o *ControlsPutPreviousValues) HasLibraryControlClassificationId() bool`

HasLibraryControlClassificationId returns a boolean if a field has been set.

### GetLibraryControlNatureId

`func (o *ControlsPutPreviousValues) GetLibraryControlNatureId() int32`

GetLibraryControlNatureId returns the LibraryControlNatureId field if non-nil, zero value otherwise.

### GetLibraryControlNatureIdOk

`func (o *ControlsPutPreviousValues) GetLibraryControlNatureIdOk() (*int32, bool)`

GetLibraryControlNatureIdOk returns a tuple with the LibraryControlNatureId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLibraryControlNatureId

`func (o *ControlsPutPreviousValues) SetLibraryControlNatureId(v int32)`

SetLibraryControlNatureId sets LibraryControlNatureId field to given value.

### HasLibraryControlNatureId

`func (o *ControlsPutPreviousValues) HasLibraryControlNatureId() bool`

HasLibraryControlNatureId returns a boolean if a field has been set.

### GetLibraryControlTypeId

`func (o *ControlsPutPreviousValues) GetLibraryControlTypeId() int32`

GetLibraryControlTypeId returns the LibraryControlTypeId field if non-nil, zero value otherwise.

### GetLibraryControlTypeIdOk

`func (o *ControlsPutPreviousValues) GetLibraryControlTypeIdOk() (*int32, bool)`

GetLibraryControlTypeIdOk returns a tuple with the LibraryControlTypeId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLibraryControlTypeId

`func (o *ControlsPutPreviousValues) SetLibraryControlTypeId(v int32)`

SetLibraryControlTypeId sets LibraryControlTypeId field to given value.

### HasLibraryControlTypeId

`func (o *ControlsPutPreviousValues) HasLibraryControlTypeId() bool`

HasLibraryControlTypeId returns a boolean if a field has been set.

### GetControlCategoryId

`func (o *ControlsPutPreviousValues) GetControlCategoryId() int32`

GetControlCategoryId returns the ControlCategoryId field if non-nil, zero value otherwise.

### GetControlCategoryIdOk

`func (o *ControlsPutPreviousValues) GetControlCategoryIdOk() (*int32, bool)`

GetControlCategoryIdOk returns a tuple with the ControlCategoryId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetControlCategoryId

`func (o *ControlsPutPreviousValues) SetControlCategoryId(v int32)`

SetControlCategoryId sets ControlCategoryId field to given value.

### HasControlCategoryId

`func (o *ControlsPutPreviousValues) HasControlCategoryId() bool`

HasControlCategoryId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


