# ControlsPutControl

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int32** | Note: This is a Primary Key.&lt;pk/&gt; | [optional] [readonly] 
**CreatedAt** | Pointer to **string** |  | [optional] [readonly] 
**UpdatedAt** | Pointer to **string** |  | [optional] [readonly] 
**DeletedAt** | Pointer to **string** |  | [optional] [readonly] 
**Uid** | **string** |  | 
**Name** | **string** |  | 
**SubprocessId** | Pointer to **int32** | Note: This is a Foreign Key to &#x60;subprocesses.id&#x60;.&lt;fk table&#x3D;&#39;subprocesses&#39; column&#x3D;&#39;id&#39;/&gt; | [optional] 
**Description** | Pointer to **string** |  | [optional] [default to ""]
**IdCode** | **string** |  | 
**FieldData** | Pointer to **interface{}** |  | [optional] 
**ComplianceControlCount** | **int32** |  | [default to 0]
**ControlObjective** | Pointer to **string** |  | [optional] 
**RiskStatement** | Pointer to **string** |  | [optional] 
**LibraryControlClassificationId** | Pointer to **int32** | Note: This is a Foreign Key to &#x60;library_control_classifications.id&#x60;.&lt;fk table&#x3D;&#39;library_control_classifications&#39; column&#x3D;&#39;id&#39;/&gt; | [optional] 
**LibraryControlNatureId** | Pointer to **int32** | Note: This is a Foreign Key to &#x60;library_control_natures.id&#x60;.&lt;fk table&#x3D;&#39;library_control_natures&#39; column&#x3D;&#39;id&#39;/&gt; | [optional] 
**LibraryControlTypeId** | Pointer to **int32** | Note: This is a Foreign Key to &#x60;library_control_types.id&#x60;.&lt;fk table&#x3D;&#39;library_control_types&#39; column&#x3D;&#39;id&#39;/&gt; | [optional] 
**ControlCategoryId** | **int32** | Note: This is a Foreign Key to &#x60;control_categories.id&#x60;.&lt;fk table&#x3D;&#39;control_categories&#39; column&#x3D;&#39;id&#39;/&gt; | 

## Methods

### NewControlsPutControl

`func NewControlsPutControl(uid string, name string, idCode string, complianceControlCount int32, controlCategoryId int32, ) *ControlsPutControl`

NewControlsPutControl instantiates a new ControlsPutControl object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewControlsPutControlWithDefaults

`func NewControlsPutControlWithDefaults() *ControlsPutControl`

NewControlsPutControlWithDefaults instantiates a new ControlsPutControl object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ControlsPutControl) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ControlsPutControl) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ControlsPutControl) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *ControlsPutControl) HasId() bool`

HasId returns a boolean if a field has been set.

### GetCreatedAt

`func (o *ControlsPutControl) GetCreatedAt() string`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *ControlsPutControl) GetCreatedAtOk() (*string, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *ControlsPutControl) SetCreatedAt(v string)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *ControlsPutControl) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *ControlsPutControl) GetUpdatedAt() string`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *ControlsPutControl) GetUpdatedAtOk() (*string, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *ControlsPutControl) SetUpdatedAt(v string)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *ControlsPutControl) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.

### GetDeletedAt

`func (o *ControlsPutControl) GetDeletedAt() string`

GetDeletedAt returns the DeletedAt field if non-nil, zero value otherwise.

### GetDeletedAtOk

`func (o *ControlsPutControl) GetDeletedAtOk() (*string, bool)`

GetDeletedAtOk returns a tuple with the DeletedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeletedAt

`func (o *ControlsPutControl) SetDeletedAt(v string)`

SetDeletedAt sets DeletedAt field to given value.

### HasDeletedAt

`func (o *ControlsPutControl) HasDeletedAt() bool`

HasDeletedAt returns a boolean if a field has been set.

### GetUid

`func (o *ControlsPutControl) GetUid() string`

GetUid returns the Uid field if non-nil, zero value otherwise.

### GetUidOk

`func (o *ControlsPutControl) GetUidOk() (*string, bool)`

GetUidOk returns a tuple with the Uid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUid

`func (o *ControlsPutControl) SetUid(v string)`

SetUid sets Uid field to given value.


### GetName

`func (o *ControlsPutControl) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ControlsPutControl) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ControlsPutControl) SetName(v string)`

SetName sets Name field to given value.


### GetSubprocessId

`func (o *ControlsPutControl) GetSubprocessId() int32`

GetSubprocessId returns the SubprocessId field if non-nil, zero value otherwise.

### GetSubprocessIdOk

`func (o *ControlsPutControl) GetSubprocessIdOk() (*int32, bool)`

GetSubprocessIdOk returns a tuple with the SubprocessId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubprocessId

`func (o *ControlsPutControl) SetSubprocessId(v int32)`

SetSubprocessId sets SubprocessId field to given value.

### HasSubprocessId

`func (o *ControlsPutControl) HasSubprocessId() bool`

HasSubprocessId returns a boolean if a field has been set.

### GetDescription

`func (o *ControlsPutControl) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *ControlsPutControl) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *ControlsPutControl) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *ControlsPutControl) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetIdCode

`func (o *ControlsPutControl) GetIdCode() string`

GetIdCode returns the IdCode field if non-nil, zero value otherwise.

### GetIdCodeOk

`func (o *ControlsPutControl) GetIdCodeOk() (*string, bool)`

GetIdCodeOk returns a tuple with the IdCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdCode

`func (o *ControlsPutControl) SetIdCode(v string)`

SetIdCode sets IdCode field to given value.


### GetFieldData

`func (o *ControlsPutControl) GetFieldData() interface{}`

GetFieldData returns the FieldData field if non-nil, zero value otherwise.

### GetFieldDataOk

`func (o *ControlsPutControl) GetFieldDataOk() (*interface{}, bool)`

GetFieldDataOk returns a tuple with the FieldData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFieldData

`func (o *ControlsPutControl) SetFieldData(v interface{})`

SetFieldData sets FieldData field to given value.

### HasFieldData

`func (o *ControlsPutControl) HasFieldData() bool`

HasFieldData returns a boolean if a field has been set.

### SetFieldDataNil

`func (o *ControlsPutControl) SetFieldDataNil(b bool)`

 SetFieldDataNil sets the value for FieldData to be an explicit nil

### UnsetFieldData
`func (o *ControlsPutControl) UnsetFieldData()`

UnsetFieldData ensures that no value is present for FieldData, not even an explicit nil
### GetComplianceControlCount

`func (o *ControlsPutControl) GetComplianceControlCount() int32`

GetComplianceControlCount returns the ComplianceControlCount field if non-nil, zero value otherwise.

### GetComplianceControlCountOk

`func (o *ControlsPutControl) GetComplianceControlCountOk() (*int32, bool)`

GetComplianceControlCountOk returns a tuple with the ComplianceControlCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComplianceControlCount

`func (o *ControlsPutControl) SetComplianceControlCount(v int32)`

SetComplianceControlCount sets ComplianceControlCount field to given value.


### GetControlObjective

`func (o *ControlsPutControl) GetControlObjective() string`

GetControlObjective returns the ControlObjective field if non-nil, zero value otherwise.

### GetControlObjectiveOk

`func (o *ControlsPutControl) GetControlObjectiveOk() (*string, bool)`

GetControlObjectiveOk returns a tuple with the ControlObjective field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetControlObjective

`func (o *ControlsPutControl) SetControlObjective(v string)`

SetControlObjective sets ControlObjective field to given value.

### HasControlObjective

`func (o *ControlsPutControl) HasControlObjective() bool`

HasControlObjective returns a boolean if a field has been set.

### GetRiskStatement

`func (o *ControlsPutControl) GetRiskStatement() string`

GetRiskStatement returns the RiskStatement field if non-nil, zero value otherwise.

### GetRiskStatementOk

`func (o *ControlsPutControl) GetRiskStatementOk() (*string, bool)`

GetRiskStatementOk returns a tuple with the RiskStatement field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRiskStatement

`func (o *ControlsPutControl) SetRiskStatement(v string)`

SetRiskStatement sets RiskStatement field to given value.

### HasRiskStatement

`func (o *ControlsPutControl) HasRiskStatement() bool`

HasRiskStatement returns a boolean if a field has been set.

### GetLibraryControlClassificationId

`func (o *ControlsPutControl) GetLibraryControlClassificationId() int32`

GetLibraryControlClassificationId returns the LibraryControlClassificationId field if non-nil, zero value otherwise.

### GetLibraryControlClassificationIdOk

`func (o *ControlsPutControl) GetLibraryControlClassificationIdOk() (*int32, bool)`

GetLibraryControlClassificationIdOk returns a tuple with the LibraryControlClassificationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLibraryControlClassificationId

`func (o *ControlsPutControl) SetLibraryControlClassificationId(v int32)`

SetLibraryControlClassificationId sets LibraryControlClassificationId field to given value.

### HasLibraryControlClassificationId

`func (o *ControlsPutControl) HasLibraryControlClassificationId() bool`

HasLibraryControlClassificationId returns a boolean if a field has been set.

### GetLibraryControlNatureId

`func (o *ControlsPutControl) GetLibraryControlNatureId() int32`

GetLibraryControlNatureId returns the LibraryControlNatureId field if non-nil, zero value otherwise.

### GetLibraryControlNatureIdOk

`func (o *ControlsPutControl) GetLibraryControlNatureIdOk() (*int32, bool)`

GetLibraryControlNatureIdOk returns a tuple with the LibraryControlNatureId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLibraryControlNatureId

`func (o *ControlsPutControl) SetLibraryControlNatureId(v int32)`

SetLibraryControlNatureId sets LibraryControlNatureId field to given value.

### HasLibraryControlNatureId

`func (o *ControlsPutControl) HasLibraryControlNatureId() bool`

HasLibraryControlNatureId returns a boolean if a field has been set.

### GetLibraryControlTypeId

`func (o *ControlsPutControl) GetLibraryControlTypeId() int32`

GetLibraryControlTypeId returns the LibraryControlTypeId field if non-nil, zero value otherwise.

### GetLibraryControlTypeIdOk

`func (o *ControlsPutControl) GetLibraryControlTypeIdOk() (*int32, bool)`

GetLibraryControlTypeIdOk returns a tuple with the LibraryControlTypeId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLibraryControlTypeId

`func (o *ControlsPutControl) SetLibraryControlTypeId(v int32)`

SetLibraryControlTypeId sets LibraryControlTypeId field to given value.

### HasLibraryControlTypeId

`func (o *ControlsPutControl) HasLibraryControlTypeId() bool`

HasLibraryControlTypeId returns a boolean if a field has been set.

### GetControlCategoryId

`func (o *ControlsPutControl) GetControlCategoryId() int32`

GetControlCategoryId returns the ControlCategoryId field if non-nil, zero value otherwise.

### GetControlCategoryIdOk

`func (o *ControlsPutControl) GetControlCategoryIdOk() (*int32, bool)`

GetControlCategoryIdOk returns a tuple with the ControlCategoryId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetControlCategoryId

`func (o *ControlsPutControl) SetControlCategoryId(v int32)`

SetControlCategoryId sets ControlCategoryId field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


