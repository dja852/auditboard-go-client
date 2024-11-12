# Frameworks

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int32** | Note: This is a Primary Key.&lt;pk/&gt; | [optional] [readonly] 
**Name** | Pointer to **string** |  | [optional] 
**SortOrder** | Pointer to **int32** |  | [optional] [default to 0]
**CreatedAt** | Pointer to **string** |  | [optional] [readonly] 
**UpdatedAt** | Pointer to **string** |  | [optional] [readonly] 
**DeletedAt** | Pointer to **string** |  | [optional] [readonly] 
**FormTemplateId** | Pointer to **int32** |  | [optional] 
**ContentServiceId** | Pointer to **int32** |  | [optional] 
**ShortName** | Pointer to **string** |  | [optional] 
**ContentSource** | Pointer to **string** |  | [optional] 
**IsCcf** | Pointer to **bool** |  | [optional] 
**RequirementLevel** | **int32** |  | [default to 0]
**FrameworkCategoryId** | Pointer to **int32** | Note: This is a Foreign Key to &#x60;framework_categories.id&#x60;.&lt;fk table&#x3D;&#39;framework_categories&#39; column&#x3D;&#39;id&#39;/&gt; | [optional] 
**Status** | **string** |  | [default to "Active"]
**ScopeDescription** | Pointer to **string** |  | [optional] 
**Version** | Pointer to **string** |  | [optional] 
**ImplementationDate** | Pointer to **string** |  | [optional] 
**ImplementationStatus** | Pointer to **string** |  | [optional] 
**ImplementationTargetDate** | Pointer to **string** |  | [optional] 
**LastAuditDate** | Pointer to **string** |  | [optional] 
**NextAuditDate** | Pointer to **string** |  | [optional] 
**EffectiveDate** | Pointer to **string** |  | [optional] 
**OwnerUserId** | Pointer to **int32** | Note: This is a Foreign Key to &#x60;users.id&#x60;.&lt;fk table&#x3D;&#39;users&#39; column&#x3D;&#39;id&#39;/&gt; | [optional] 
**Scopes** | Pointer to **interface{}** |  | [optional] 
**XcTotals** | Pointer to **interface{}** |  | [optional] 
**RatingType** | Pointer to **string** |  | [optional] 
**LicenseRequired** | **bool** |  | [default to false]
**LicenseApprovedBy** | Pointer to **int32** | Note: This is a Foreign Key to &#x60;users.id&#x60;.&lt;fk table&#x3D;&#39;users&#39; column&#x3D;&#39;id&#39;/&gt; | [optional] 
**LicenseApprovedOn** | Pointer to **string** |  | [optional] 

## Methods

### NewFrameworks

`func NewFrameworks(requirementLevel int32, status string, licenseRequired bool, ) *Frameworks`

NewFrameworks instantiates a new Frameworks object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFrameworksWithDefaults

`func NewFrameworksWithDefaults() *Frameworks`

NewFrameworksWithDefaults instantiates a new Frameworks object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *Frameworks) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Frameworks) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Frameworks) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *Frameworks) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *Frameworks) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *Frameworks) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *Frameworks) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *Frameworks) HasName() bool`

HasName returns a boolean if a field has been set.

### GetSortOrder

`func (o *Frameworks) GetSortOrder() int32`

GetSortOrder returns the SortOrder field if non-nil, zero value otherwise.

### GetSortOrderOk

`func (o *Frameworks) GetSortOrderOk() (*int32, bool)`

GetSortOrderOk returns a tuple with the SortOrder field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSortOrder

`func (o *Frameworks) SetSortOrder(v int32)`

SetSortOrder sets SortOrder field to given value.

### HasSortOrder

`func (o *Frameworks) HasSortOrder() bool`

HasSortOrder returns a boolean if a field has been set.

### GetCreatedAt

`func (o *Frameworks) GetCreatedAt() string`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *Frameworks) GetCreatedAtOk() (*string, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *Frameworks) SetCreatedAt(v string)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *Frameworks) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *Frameworks) GetUpdatedAt() string`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *Frameworks) GetUpdatedAtOk() (*string, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *Frameworks) SetUpdatedAt(v string)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *Frameworks) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.

### GetDeletedAt

`func (o *Frameworks) GetDeletedAt() string`

GetDeletedAt returns the DeletedAt field if non-nil, zero value otherwise.

### GetDeletedAtOk

`func (o *Frameworks) GetDeletedAtOk() (*string, bool)`

GetDeletedAtOk returns a tuple with the DeletedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeletedAt

`func (o *Frameworks) SetDeletedAt(v string)`

SetDeletedAt sets DeletedAt field to given value.

### HasDeletedAt

`func (o *Frameworks) HasDeletedAt() bool`

HasDeletedAt returns a boolean if a field has been set.

### GetFormTemplateId

`func (o *Frameworks) GetFormTemplateId() int32`

GetFormTemplateId returns the FormTemplateId field if non-nil, zero value otherwise.

### GetFormTemplateIdOk

`func (o *Frameworks) GetFormTemplateIdOk() (*int32, bool)`

GetFormTemplateIdOk returns a tuple with the FormTemplateId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFormTemplateId

`func (o *Frameworks) SetFormTemplateId(v int32)`

SetFormTemplateId sets FormTemplateId field to given value.

### HasFormTemplateId

`func (o *Frameworks) HasFormTemplateId() bool`

HasFormTemplateId returns a boolean if a field has been set.

### GetContentServiceId

`func (o *Frameworks) GetContentServiceId() int32`

GetContentServiceId returns the ContentServiceId field if non-nil, zero value otherwise.

### GetContentServiceIdOk

`func (o *Frameworks) GetContentServiceIdOk() (*int32, bool)`

GetContentServiceIdOk returns a tuple with the ContentServiceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContentServiceId

`func (o *Frameworks) SetContentServiceId(v int32)`

SetContentServiceId sets ContentServiceId field to given value.

### HasContentServiceId

`func (o *Frameworks) HasContentServiceId() bool`

HasContentServiceId returns a boolean if a field has been set.

### GetShortName

`func (o *Frameworks) GetShortName() string`

GetShortName returns the ShortName field if non-nil, zero value otherwise.

### GetShortNameOk

`func (o *Frameworks) GetShortNameOk() (*string, bool)`

GetShortNameOk returns a tuple with the ShortName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShortName

`func (o *Frameworks) SetShortName(v string)`

SetShortName sets ShortName field to given value.

### HasShortName

`func (o *Frameworks) HasShortName() bool`

HasShortName returns a boolean if a field has been set.

### GetContentSource

`func (o *Frameworks) GetContentSource() string`

GetContentSource returns the ContentSource field if non-nil, zero value otherwise.

### GetContentSourceOk

`func (o *Frameworks) GetContentSourceOk() (*string, bool)`

GetContentSourceOk returns a tuple with the ContentSource field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContentSource

`func (o *Frameworks) SetContentSource(v string)`

SetContentSource sets ContentSource field to given value.

### HasContentSource

`func (o *Frameworks) HasContentSource() bool`

HasContentSource returns a boolean if a field has been set.

### GetIsCcf

`func (o *Frameworks) GetIsCcf() bool`

GetIsCcf returns the IsCcf field if non-nil, zero value otherwise.

### GetIsCcfOk

`func (o *Frameworks) GetIsCcfOk() (*bool, bool)`

GetIsCcfOk returns a tuple with the IsCcf field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsCcf

`func (o *Frameworks) SetIsCcf(v bool)`

SetIsCcf sets IsCcf field to given value.

### HasIsCcf

`func (o *Frameworks) HasIsCcf() bool`

HasIsCcf returns a boolean if a field has been set.

### GetRequirementLevel

`func (o *Frameworks) GetRequirementLevel() int32`

GetRequirementLevel returns the RequirementLevel field if non-nil, zero value otherwise.

### GetRequirementLevelOk

`func (o *Frameworks) GetRequirementLevelOk() (*int32, bool)`

GetRequirementLevelOk returns a tuple with the RequirementLevel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequirementLevel

`func (o *Frameworks) SetRequirementLevel(v int32)`

SetRequirementLevel sets RequirementLevel field to given value.


### GetFrameworkCategoryId

`func (o *Frameworks) GetFrameworkCategoryId() int32`

GetFrameworkCategoryId returns the FrameworkCategoryId field if non-nil, zero value otherwise.

### GetFrameworkCategoryIdOk

`func (o *Frameworks) GetFrameworkCategoryIdOk() (*int32, bool)`

GetFrameworkCategoryIdOk returns a tuple with the FrameworkCategoryId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFrameworkCategoryId

`func (o *Frameworks) SetFrameworkCategoryId(v int32)`

SetFrameworkCategoryId sets FrameworkCategoryId field to given value.

### HasFrameworkCategoryId

`func (o *Frameworks) HasFrameworkCategoryId() bool`

HasFrameworkCategoryId returns a boolean if a field has been set.

### GetStatus

`func (o *Frameworks) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *Frameworks) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *Frameworks) SetStatus(v string)`

SetStatus sets Status field to given value.


### GetScopeDescription

`func (o *Frameworks) GetScopeDescription() string`

GetScopeDescription returns the ScopeDescription field if non-nil, zero value otherwise.

### GetScopeDescriptionOk

`func (o *Frameworks) GetScopeDescriptionOk() (*string, bool)`

GetScopeDescriptionOk returns a tuple with the ScopeDescription field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScopeDescription

`func (o *Frameworks) SetScopeDescription(v string)`

SetScopeDescription sets ScopeDescription field to given value.

### HasScopeDescription

`func (o *Frameworks) HasScopeDescription() bool`

HasScopeDescription returns a boolean if a field has been set.

### GetVersion

`func (o *Frameworks) GetVersion() string`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *Frameworks) GetVersionOk() (*string, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *Frameworks) SetVersion(v string)`

SetVersion sets Version field to given value.

### HasVersion

`func (o *Frameworks) HasVersion() bool`

HasVersion returns a boolean if a field has been set.

### GetImplementationDate

`func (o *Frameworks) GetImplementationDate() string`

GetImplementationDate returns the ImplementationDate field if non-nil, zero value otherwise.

### GetImplementationDateOk

`func (o *Frameworks) GetImplementationDateOk() (*string, bool)`

GetImplementationDateOk returns a tuple with the ImplementationDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImplementationDate

`func (o *Frameworks) SetImplementationDate(v string)`

SetImplementationDate sets ImplementationDate field to given value.

### HasImplementationDate

`func (o *Frameworks) HasImplementationDate() bool`

HasImplementationDate returns a boolean if a field has been set.

### GetImplementationStatus

`func (o *Frameworks) GetImplementationStatus() string`

GetImplementationStatus returns the ImplementationStatus field if non-nil, zero value otherwise.

### GetImplementationStatusOk

`func (o *Frameworks) GetImplementationStatusOk() (*string, bool)`

GetImplementationStatusOk returns a tuple with the ImplementationStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImplementationStatus

`func (o *Frameworks) SetImplementationStatus(v string)`

SetImplementationStatus sets ImplementationStatus field to given value.

### HasImplementationStatus

`func (o *Frameworks) HasImplementationStatus() bool`

HasImplementationStatus returns a boolean if a field has been set.

### GetImplementationTargetDate

`func (o *Frameworks) GetImplementationTargetDate() string`

GetImplementationTargetDate returns the ImplementationTargetDate field if non-nil, zero value otherwise.

### GetImplementationTargetDateOk

`func (o *Frameworks) GetImplementationTargetDateOk() (*string, bool)`

GetImplementationTargetDateOk returns a tuple with the ImplementationTargetDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImplementationTargetDate

`func (o *Frameworks) SetImplementationTargetDate(v string)`

SetImplementationTargetDate sets ImplementationTargetDate field to given value.

### HasImplementationTargetDate

`func (o *Frameworks) HasImplementationTargetDate() bool`

HasImplementationTargetDate returns a boolean if a field has been set.

### GetLastAuditDate

`func (o *Frameworks) GetLastAuditDate() string`

GetLastAuditDate returns the LastAuditDate field if non-nil, zero value otherwise.

### GetLastAuditDateOk

`func (o *Frameworks) GetLastAuditDateOk() (*string, bool)`

GetLastAuditDateOk returns a tuple with the LastAuditDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastAuditDate

`func (o *Frameworks) SetLastAuditDate(v string)`

SetLastAuditDate sets LastAuditDate field to given value.

### HasLastAuditDate

`func (o *Frameworks) HasLastAuditDate() bool`

HasLastAuditDate returns a boolean if a field has been set.

### GetNextAuditDate

`func (o *Frameworks) GetNextAuditDate() string`

GetNextAuditDate returns the NextAuditDate field if non-nil, zero value otherwise.

### GetNextAuditDateOk

`func (o *Frameworks) GetNextAuditDateOk() (*string, bool)`

GetNextAuditDateOk returns a tuple with the NextAuditDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNextAuditDate

`func (o *Frameworks) SetNextAuditDate(v string)`

SetNextAuditDate sets NextAuditDate field to given value.

### HasNextAuditDate

`func (o *Frameworks) HasNextAuditDate() bool`

HasNextAuditDate returns a boolean if a field has been set.

### GetEffectiveDate

`func (o *Frameworks) GetEffectiveDate() string`

GetEffectiveDate returns the EffectiveDate field if non-nil, zero value otherwise.

### GetEffectiveDateOk

`func (o *Frameworks) GetEffectiveDateOk() (*string, bool)`

GetEffectiveDateOk returns a tuple with the EffectiveDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEffectiveDate

`func (o *Frameworks) SetEffectiveDate(v string)`

SetEffectiveDate sets EffectiveDate field to given value.

### HasEffectiveDate

`func (o *Frameworks) HasEffectiveDate() bool`

HasEffectiveDate returns a boolean if a field has been set.

### GetOwnerUserId

`func (o *Frameworks) GetOwnerUserId() int32`

GetOwnerUserId returns the OwnerUserId field if non-nil, zero value otherwise.

### GetOwnerUserIdOk

`func (o *Frameworks) GetOwnerUserIdOk() (*int32, bool)`

GetOwnerUserIdOk returns a tuple with the OwnerUserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwnerUserId

`func (o *Frameworks) SetOwnerUserId(v int32)`

SetOwnerUserId sets OwnerUserId field to given value.

### HasOwnerUserId

`func (o *Frameworks) HasOwnerUserId() bool`

HasOwnerUserId returns a boolean if a field has been set.

### GetScopes

`func (o *Frameworks) GetScopes() interface{}`

GetScopes returns the Scopes field if non-nil, zero value otherwise.

### GetScopesOk

`func (o *Frameworks) GetScopesOk() (*interface{}, bool)`

GetScopesOk returns a tuple with the Scopes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScopes

`func (o *Frameworks) SetScopes(v interface{})`

SetScopes sets Scopes field to given value.

### HasScopes

`func (o *Frameworks) HasScopes() bool`

HasScopes returns a boolean if a field has been set.

### SetScopesNil

`func (o *Frameworks) SetScopesNil(b bool)`

 SetScopesNil sets the value for Scopes to be an explicit nil

### UnsetScopes
`func (o *Frameworks) UnsetScopes()`

UnsetScopes ensures that no value is present for Scopes, not even an explicit nil
### GetXcTotals

`func (o *Frameworks) GetXcTotals() interface{}`

GetXcTotals returns the XcTotals field if non-nil, zero value otherwise.

### GetXcTotalsOk

`func (o *Frameworks) GetXcTotalsOk() (*interface{}, bool)`

GetXcTotalsOk returns a tuple with the XcTotals field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetXcTotals

`func (o *Frameworks) SetXcTotals(v interface{})`

SetXcTotals sets XcTotals field to given value.

### HasXcTotals

`func (o *Frameworks) HasXcTotals() bool`

HasXcTotals returns a boolean if a field has been set.

### SetXcTotalsNil

`func (o *Frameworks) SetXcTotalsNil(b bool)`

 SetXcTotalsNil sets the value for XcTotals to be an explicit nil

### UnsetXcTotals
`func (o *Frameworks) UnsetXcTotals()`

UnsetXcTotals ensures that no value is present for XcTotals, not even an explicit nil
### GetRatingType

`func (o *Frameworks) GetRatingType() string`

GetRatingType returns the RatingType field if non-nil, zero value otherwise.

### GetRatingTypeOk

`func (o *Frameworks) GetRatingTypeOk() (*string, bool)`

GetRatingTypeOk returns a tuple with the RatingType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRatingType

`func (o *Frameworks) SetRatingType(v string)`

SetRatingType sets RatingType field to given value.

### HasRatingType

`func (o *Frameworks) HasRatingType() bool`

HasRatingType returns a boolean if a field has been set.

### GetLicenseRequired

`func (o *Frameworks) GetLicenseRequired() bool`

GetLicenseRequired returns the LicenseRequired field if non-nil, zero value otherwise.

### GetLicenseRequiredOk

`func (o *Frameworks) GetLicenseRequiredOk() (*bool, bool)`

GetLicenseRequiredOk returns a tuple with the LicenseRequired field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLicenseRequired

`func (o *Frameworks) SetLicenseRequired(v bool)`

SetLicenseRequired sets LicenseRequired field to given value.


### GetLicenseApprovedBy

`func (o *Frameworks) GetLicenseApprovedBy() int32`

GetLicenseApprovedBy returns the LicenseApprovedBy field if non-nil, zero value otherwise.

### GetLicenseApprovedByOk

`func (o *Frameworks) GetLicenseApprovedByOk() (*int32, bool)`

GetLicenseApprovedByOk returns a tuple with the LicenseApprovedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLicenseApprovedBy

`func (o *Frameworks) SetLicenseApprovedBy(v int32)`

SetLicenseApprovedBy sets LicenseApprovedBy field to given value.

### HasLicenseApprovedBy

`func (o *Frameworks) HasLicenseApprovedBy() bool`

HasLicenseApprovedBy returns a boolean if a field has been set.

### GetLicenseApprovedOn

`func (o *Frameworks) GetLicenseApprovedOn() string`

GetLicenseApprovedOn returns the LicenseApprovedOn field if non-nil, zero value otherwise.

### GetLicenseApprovedOnOk

`func (o *Frameworks) GetLicenseApprovedOnOk() (*string, bool)`

GetLicenseApprovedOnOk returns a tuple with the LicenseApprovedOn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLicenseApprovedOn

`func (o *Frameworks) SetLicenseApprovedOn(v string)`

SetLicenseApprovedOn sets LicenseApprovedOn field to given value.

### HasLicenseApprovedOn

`func (o *Frameworks) HasLicenseApprovedOn() bool`

HasLicenseApprovedOn returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


