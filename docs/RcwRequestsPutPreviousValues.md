# RcwRequestsPutPreviousValues

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int32** | Note: This is a Primary Key.&lt;pk/&gt; | [optional] [readonly] 
**Title** | Pointer to **string** |  | [optional] 
**RcwLibraryRequestId** | Pointer to **int32** | Note: This is a Foreign Key to &#x60;rcw_library_requests.id&#x60;.&lt;fk table&#x3D;&#39;rcw_library_requests&#39; column&#x3D;&#39;id&#39;/&gt; | [optional] 
**RcwProjectId** | Pointer to **int32** | Note: This is a Foreign Key to &#x60;rcw_projects.id&#x60;.&lt;fk table&#x3D;&#39;rcw_projects&#39; column&#x3D;&#39;id&#39;/&gt; | [optional] 
**AssigneeUserId** | Pointer to **int32** | Note: This is a Foreign Key to &#x60;users.id&#x60;.&lt;fk table&#x3D;&#39;users&#39; column&#x3D;&#39;id&#39;/&gt; | [optional] 
**Status** | Pointer to **string** |  | [optional] 
**OpenDate** | Pointer to **string** |  | [optional] 
**OpenByUserId** | Pointer to **int32** | Note: This is a Foreign Key to &#x60;users.id&#x60;.&lt;fk table&#x3D;&#39;users&#39; column&#x3D;&#39;id&#39;/&gt; | [optional] 
**PendingReviewDate** | Pointer to **string** |  | [optional] 
**PendingReviewByUserId** | Pointer to **int32** | Note: This is a Foreign Key to &#x60;users.id&#x60;.&lt;fk table&#x3D;&#39;users&#39; column&#x3D;&#39;id&#39;/&gt; | [optional] 
**ReopenedDate** | Pointer to **string** |  | [optional] 
**ReopenedByUserId** | Pointer to **int32** | Note: This is a Foreign Key to &#x60;users.id&#x60;.&lt;fk table&#x3D;&#39;users&#39; column&#x3D;&#39;id&#39;/&gt; | [optional] 
**CompletedDate** | Pointer to **string** |  | [optional] 
**CompletedByUserId** | Pointer to **int32** | Note: This is a Foreign Key to &#x60;users.id&#x60;.&lt;fk table&#x3D;&#39;users&#39; column&#x3D;&#39;id&#39;/&gt; | [optional] 
**DueDate** | Pointer to **string** |  | [optional] 
**CreatedAt** | Pointer to **string** |  | [optional] [readonly] 
**UpdatedAt** | Pointer to **string** |  | [optional] [readonly] 
**DeletedAt** | Pointer to **string** |  | [optional] [readonly] 
**RcwRequestRatingId** | Pointer to **int32** | Note: This is a Foreign Key to &#x60;rcw_request_ratings.id&#x60;.&lt;fk table&#x3D;&#39;rcw_request_ratings&#39; column&#x3D;&#39;id&#39;/&gt; | [optional] 
**DidNotify** | Pointer to **bool** |  | [optional] [default to false]
**ExplanationText** | Pointer to **string** |  | [optional] 
**IsAutomated** | Pointer to **bool** |  | [optional] 
**LaunchDate** | Pointer to **string** |  | [optional] 
**PeriodStart** | Pointer to **string** |  | [optional] 
**PeriodEnd** | Pointer to **string** |  | [optional] 
**ExternalIntegrationUrl** | Pointer to **string** |  | [optional] 
**PeriodDependent** | Pointer to **bool** |  | [optional] [default to false]
**Description** | Pointer to **string** |  | [optional] 
**AdditionalInformation** | Pointer to **string** |  | [optional] 
**Scopes** | Pointer to **interface{}** |  | [optional] 
**IsRecurrence** | Pointer to **bool** |  | [optional] 
**SubmittedByUserId** | Pointer to **int32** | Note: This is a Foreign Key to &#x60;users.id&#x60;.&lt;fk table&#x3D;&#39;users&#39; column&#x3D;&#39;id&#39;/&gt; | [optional] 

## Methods

### NewRcwRequestsPutPreviousValues

`func NewRcwRequestsPutPreviousValues() *RcwRequestsPutPreviousValues`

NewRcwRequestsPutPreviousValues instantiates a new RcwRequestsPutPreviousValues object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRcwRequestsPutPreviousValuesWithDefaults

`func NewRcwRequestsPutPreviousValuesWithDefaults() *RcwRequestsPutPreviousValues`

NewRcwRequestsPutPreviousValuesWithDefaults instantiates a new RcwRequestsPutPreviousValues object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *RcwRequestsPutPreviousValues) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *RcwRequestsPutPreviousValues) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *RcwRequestsPutPreviousValues) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *RcwRequestsPutPreviousValues) HasId() bool`

HasId returns a boolean if a field has been set.

### GetTitle

`func (o *RcwRequestsPutPreviousValues) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *RcwRequestsPutPreviousValues) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *RcwRequestsPutPreviousValues) SetTitle(v string)`

SetTitle sets Title field to given value.

### HasTitle

`func (o *RcwRequestsPutPreviousValues) HasTitle() bool`

HasTitle returns a boolean if a field has been set.

### GetRcwLibraryRequestId

`func (o *RcwRequestsPutPreviousValues) GetRcwLibraryRequestId() int32`

GetRcwLibraryRequestId returns the RcwLibraryRequestId field if non-nil, zero value otherwise.

### GetRcwLibraryRequestIdOk

`func (o *RcwRequestsPutPreviousValues) GetRcwLibraryRequestIdOk() (*int32, bool)`

GetRcwLibraryRequestIdOk returns a tuple with the RcwLibraryRequestId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRcwLibraryRequestId

`func (o *RcwRequestsPutPreviousValues) SetRcwLibraryRequestId(v int32)`

SetRcwLibraryRequestId sets RcwLibraryRequestId field to given value.

### HasRcwLibraryRequestId

`func (o *RcwRequestsPutPreviousValues) HasRcwLibraryRequestId() bool`

HasRcwLibraryRequestId returns a boolean if a field has been set.

### GetRcwProjectId

`func (o *RcwRequestsPutPreviousValues) GetRcwProjectId() int32`

GetRcwProjectId returns the RcwProjectId field if non-nil, zero value otherwise.

### GetRcwProjectIdOk

`func (o *RcwRequestsPutPreviousValues) GetRcwProjectIdOk() (*int32, bool)`

GetRcwProjectIdOk returns a tuple with the RcwProjectId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRcwProjectId

`func (o *RcwRequestsPutPreviousValues) SetRcwProjectId(v int32)`

SetRcwProjectId sets RcwProjectId field to given value.

### HasRcwProjectId

`func (o *RcwRequestsPutPreviousValues) HasRcwProjectId() bool`

HasRcwProjectId returns a boolean if a field has been set.

### GetAssigneeUserId

`func (o *RcwRequestsPutPreviousValues) GetAssigneeUserId() int32`

GetAssigneeUserId returns the AssigneeUserId field if non-nil, zero value otherwise.

### GetAssigneeUserIdOk

`func (o *RcwRequestsPutPreviousValues) GetAssigneeUserIdOk() (*int32, bool)`

GetAssigneeUserIdOk returns a tuple with the AssigneeUserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssigneeUserId

`func (o *RcwRequestsPutPreviousValues) SetAssigneeUserId(v int32)`

SetAssigneeUserId sets AssigneeUserId field to given value.

### HasAssigneeUserId

`func (o *RcwRequestsPutPreviousValues) HasAssigneeUserId() bool`

HasAssigneeUserId returns a boolean if a field has been set.

### GetStatus

`func (o *RcwRequestsPutPreviousValues) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *RcwRequestsPutPreviousValues) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *RcwRequestsPutPreviousValues) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *RcwRequestsPutPreviousValues) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetOpenDate

`func (o *RcwRequestsPutPreviousValues) GetOpenDate() string`

GetOpenDate returns the OpenDate field if non-nil, zero value otherwise.

### GetOpenDateOk

`func (o *RcwRequestsPutPreviousValues) GetOpenDateOk() (*string, bool)`

GetOpenDateOk returns a tuple with the OpenDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOpenDate

`func (o *RcwRequestsPutPreviousValues) SetOpenDate(v string)`

SetOpenDate sets OpenDate field to given value.

### HasOpenDate

`func (o *RcwRequestsPutPreviousValues) HasOpenDate() bool`

HasOpenDate returns a boolean if a field has been set.

### GetOpenByUserId

`func (o *RcwRequestsPutPreviousValues) GetOpenByUserId() int32`

GetOpenByUserId returns the OpenByUserId field if non-nil, zero value otherwise.

### GetOpenByUserIdOk

`func (o *RcwRequestsPutPreviousValues) GetOpenByUserIdOk() (*int32, bool)`

GetOpenByUserIdOk returns a tuple with the OpenByUserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOpenByUserId

`func (o *RcwRequestsPutPreviousValues) SetOpenByUserId(v int32)`

SetOpenByUserId sets OpenByUserId field to given value.

### HasOpenByUserId

`func (o *RcwRequestsPutPreviousValues) HasOpenByUserId() bool`

HasOpenByUserId returns a boolean if a field has been set.

### GetPendingReviewDate

`func (o *RcwRequestsPutPreviousValues) GetPendingReviewDate() string`

GetPendingReviewDate returns the PendingReviewDate field if non-nil, zero value otherwise.

### GetPendingReviewDateOk

`func (o *RcwRequestsPutPreviousValues) GetPendingReviewDateOk() (*string, bool)`

GetPendingReviewDateOk returns a tuple with the PendingReviewDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPendingReviewDate

`func (o *RcwRequestsPutPreviousValues) SetPendingReviewDate(v string)`

SetPendingReviewDate sets PendingReviewDate field to given value.

### HasPendingReviewDate

`func (o *RcwRequestsPutPreviousValues) HasPendingReviewDate() bool`

HasPendingReviewDate returns a boolean if a field has been set.

### GetPendingReviewByUserId

`func (o *RcwRequestsPutPreviousValues) GetPendingReviewByUserId() int32`

GetPendingReviewByUserId returns the PendingReviewByUserId field if non-nil, zero value otherwise.

### GetPendingReviewByUserIdOk

`func (o *RcwRequestsPutPreviousValues) GetPendingReviewByUserIdOk() (*int32, bool)`

GetPendingReviewByUserIdOk returns a tuple with the PendingReviewByUserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPendingReviewByUserId

`func (o *RcwRequestsPutPreviousValues) SetPendingReviewByUserId(v int32)`

SetPendingReviewByUserId sets PendingReviewByUserId field to given value.

### HasPendingReviewByUserId

`func (o *RcwRequestsPutPreviousValues) HasPendingReviewByUserId() bool`

HasPendingReviewByUserId returns a boolean if a field has been set.

### GetReopenedDate

`func (o *RcwRequestsPutPreviousValues) GetReopenedDate() string`

GetReopenedDate returns the ReopenedDate field if non-nil, zero value otherwise.

### GetReopenedDateOk

`func (o *RcwRequestsPutPreviousValues) GetReopenedDateOk() (*string, bool)`

GetReopenedDateOk returns a tuple with the ReopenedDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReopenedDate

`func (o *RcwRequestsPutPreviousValues) SetReopenedDate(v string)`

SetReopenedDate sets ReopenedDate field to given value.

### HasReopenedDate

`func (o *RcwRequestsPutPreviousValues) HasReopenedDate() bool`

HasReopenedDate returns a boolean if a field has been set.

### GetReopenedByUserId

`func (o *RcwRequestsPutPreviousValues) GetReopenedByUserId() int32`

GetReopenedByUserId returns the ReopenedByUserId field if non-nil, zero value otherwise.

### GetReopenedByUserIdOk

`func (o *RcwRequestsPutPreviousValues) GetReopenedByUserIdOk() (*int32, bool)`

GetReopenedByUserIdOk returns a tuple with the ReopenedByUserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReopenedByUserId

`func (o *RcwRequestsPutPreviousValues) SetReopenedByUserId(v int32)`

SetReopenedByUserId sets ReopenedByUserId field to given value.

### HasReopenedByUserId

`func (o *RcwRequestsPutPreviousValues) HasReopenedByUserId() bool`

HasReopenedByUserId returns a boolean if a field has been set.

### GetCompletedDate

`func (o *RcwRequestsPutPreviousValues) GetCompletedDate() string`

GetCompletedDate returns the CompletedDate field if non-nil, zero value otherwise.

### GetCompletedDateOk

`func (o *RcwRequestsPutPreviousValues) GetCompletedDateOk() (*string, bool)`

GetCompletedDateOk returns a tuple with the CompletedDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompletedDate

`func (o *RcwRequestsPutPreviousValues) SetCompletedDate(v string)`

SetCompletedDate sets CompletedDate field to given value.

### HasCompletedDate

`func (o *RcwRequestsPutPreviousValues) HasCompletedDate() bool`

HasCompletedDate returns a boolean if a field has been set.

### GetCompletedByUserId

`func (o *RcwRequestsPutPreviousValues) GetCompletedByUserId() int32`

GetCompletedByUserId returns the CompletedByUserId field if non-nil, zero value otherwise.

### GetCompletedByUserIdOk

`func (o *RcwRequestsPutPreviousValues) GetCompletedByUserIdOk() (*int32, bool)`

GetCompletedByUserIdOk returns a tuple with the CompletedByUserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompletedByUserId

`func (o *RcwRequestsPutPreviousValues) SetCompletedByUserId(v int32)`

SetCompletedByUserId sets CompletedByUserId field to given value.

### HasCompletedByUserId

`func (o *RcwRequestsPutPreviousValues) HasCompletedByUserId() bool`

HasCompletedByUserId returns a boolean if a field has been set.

### GetDueDate

`func (o *RcwRequestsPutPreviousValues) GetDueDate() string`

GetDueDate returns the DueDate field if non-nil, zero value otherwise.

### GetDueDateOk

`func (o *RcwRequestsPutPreviousValues) GetDueDateOk() (*string, bool)`

GetDueDateOk returns a tuple with the DueDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDueDate

`func (o *RcwRequestsPutPreviousValues) SetDueDate(v string)`

SetDueDate sets DueDate field to given value.

### HasDueDate

`func (o *RcwRequestsPutPreviousValues) HasDueDate() bool`

HasDueDate returns a boolean if a field has been set.

### GetCreatedAt

`func (o *RcwRequestsPutPreviousValues) GetCreatedAt() string`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *RcwRequestsPutPreviousValues) GetCreatedAtOk() (*string, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *RcwRequestsPutPreviousValues) SetCreatedAt(v string)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *RcwRequestsPutPreviousValues) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *RcwRequestsPutPreviousValues) GetUpdatedAt() string`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *RcwRequestsPutPreviousValues) GetUpdatedAtOk() (*string, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *RcwRequestsPutPreviousValues) SetUpdatedAt(v string)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *RcwRequestsPutPreviousValues) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.

### GetDeletedAt

`func (o *RcwRequestsPutPreviousValues) GetDeletedAt() string`

GetDeletedAt returns the DeletedAt field if non-nil, zero value otherwise.

### GetDeletedAtOk

`func (o *RcwRequestsPutPreviousValues) GetDeletedAtOk() (*string, bool)`

GetDeletedAtOk returns a tuple with the DeletedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeletedAt

`func (o *RcwRequestsPutPreviousValues) SetDeletedAt(v string)`

SetDeletedAt sets DeletedAt field to given value.

### HasDeletedAt

`func (o *RcwRequestsPutPreviousValues) HasDeletedAt() bool`

HasDeletedAt returns a boolean if a field has been set.

### GetRcwRequestRatingId

`func (o *RcwRequestsPutPreviousValues) GetRcwRequestRatingId() int32`

GetRcwRequestRatingId returns the RcwRequestRatingId field if non-nil, zero value otherwise.

### GetRcwRequestRatingIdOk

`func (o *RcwRequestsPutPreviousValues) GetRcwRequestRatingIdOk() (*int32, bool)`

GetRcwRequestRatingIdOk returns a tuple with the RcwRequestRatingId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRcwRequestRatingId

`func (o *RcwRequestsPutPreviousValues) SetRcwRequestRatingId(v int32)`

SetRcwRequestRatingId sets RcwRequestRatingId field to given value.

### HasRcwRequestRatingId

`func (o *RcwRequestsPutPreviousValues) HasRcwRequestRatingId() bool`

HasRcwRequestRatingId returns a boolean if a field has been set.

### GetDidNotify

`func (o *RcwRequestsPutPreviousValues) GetDidNotify() bool`

GetDidNotify returns the DidNotify field if non-nil, zero value otherwise.

### GetDidNotifyOk

`func (o *RcwRequestsPutPreviousValues) GetDidNotifyOk() (*bool, bool)`

GetDidNotifyOk returns a tuple with the DidNotify field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDidNotify

`func (o *RcwRequestsPutPreviousValues) SetDidNotify(v bool)`

SetDidNotify sets DidNotify field to given value.

### HasDidNotify

`func (o *RcwRequestsPutPreviousValues) HasDidNotify() bool`

HasDidNotify returns a boolean if a field has been set.

### GetExplanationText

`func (o *RcwRequestsPutPreviousValues) GetExplanationText() string`

GetExplanationText returns the ExplanationText field if non-nil, zero value otherwise.

### GetExplanationTextOk

`func (o *RcwRequestsPutPreviousValues) GetExplanationTextOk() (*string, bool)`

GetExplanationTextOk returns a tuple with the ExplanationText field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExplanationText

`func (o *RcwRequestsPutPreviousValues) SetExplanationText(v string)`

SetExplanationText sets ExplanationText field to given value.

### HasExplanationText

`func (o *RcwRequestsPutPreviousValues) HasExplanationText() bool`

HasExplanationText returns a boolean if a field has been set.

### GetIsAutomated

`func (o *RcwRequestsPutPreviousValues) GetIsAutomated() bool`

GetIsAutomated returns the IsAutomated field if non-nil, zero value otherwise.

### GetIsAutomatedOk

`func (o *RcwRequestsPutPreviousValues) GetIsAutomatedOk() (*bool, bool)`

GetIsAutomatedOk returns a tuple with the IsAutomated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsAutomated

`func (o *RcwRequestsPutPreviousValues) SetIsAutomated(v bool)`

SetIsAutomated sets IsAutomated field to given value.

### HasIsAutomated

`func (o *RcwRequestsPutPreviousValues) HasIsAutomated() bool`

HasIsAutomated returns a boolean if a field has been set.

### GetLaunchDate

`func (o *RcwRequestsPutPreviousValues) GetLaunchDate() string`

GetLaunchDate returns the LaunchDate field if non-nil, zero value otherwise.

### GetLaunchDateOk

`func (o *RcwRequestsPutPreviousValues) GetLaunchDateOk() (*string, bool)`

GetLaunchDateOk returns a tuple with the LaunchDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLaunchDate

`func (o *RcwRequestsPutPreviousValues) SetLaunchDate(v string)`

SetLaunchDate sets LaunchDate field to given value.

### HasLaunchDate

`func (o *RcwRequestsPutPreviousValues) HasLaunchDate() bool`

HasLaunchDate returns a boolean if a field has been set.

### GetPeriodStart

`func (o *RcwRequestsPutPreviousValues) GetPeriodStart() string`

GetPeriodStart returns the PeriodStart field if non-nil, zero value otherwise.

### GetPeriodStartOk

`func (o *RcwRequestsPutPreviousValues) GetPeriodStartOk() (*string, bool)`

GetPeriodStartOk returns a tuple with the PeriodStart field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPeriodStart

`func (o *RcwRequestsPutPreviousValues) SetPeriodStart(v string)`

SetPeriodStart sets PeriodStart field to given value.

### HasPeriodStart

`func (o *RcwRequestsPutPreviousValues) HasPeriodStart() bool`

HasPeriodStart returns a boolean if a field has been set.

### GetPeriodEnd

`func (o *RcwRequestsPutPreviousValues) GetPeriodEnd() string`

GetPeriodEnd returns the PeriodEnd field if non-nil, zero value otherwise.

### GetPeriodEndOk

`func (o *RcwRequestsPutPreviousValues) GetPeriodEndOk() (*string, bool)`

GetPeriodEndOk returns a tuple with the PeriodEnd field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPeriodEnd

`func (o *RcwRequestsPutPreviousValues) SetPeriodEnd(v string)`

SetPeriodEnd sets PeriodEnd field to given value.

### HasPeriodEnd

`func (o *RcwRequestsPutPreviousValues) HasPeriodEnd() bool`

HasPeriodEnd returns a boolean if a field has been set.

### GetExternalIntegrationUrl

`func (o *RcwRequestsPutPreviousValues) GetExternalIntegrationUrl() string`

GetExternalIntegrationUrl returns the ExternalIntegrationUrl field if non-nil, zero value otherwise.

### GetExternalIntegrationUrlOk

`func (o *RcwRequestsPutPreviousValues) GetExternalIntegrationUrlOk() (*string, bool)`

GetExternalIntegrationUrlOk returns a tuple with the ExternalIntegrationUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalIntegrationUrl

`func (o *RcwRequestsPutPreviousValues) SetExternalIntegrationUrl(v string)`

SetExternalIntegrationUrl sets ExternalIntegrationUrl field to given value.

### HasExternalIntegrationUrl

`func (o *RcwRequestsPutPreviousValues) HasExternalIntegrationUrl() bool`

HasExternalIntegrationUrl returns a boolean if a field has been set.

### GetPeriodDependent

`func (o *RcwRequestsPutPreviousValues) GetPeriodDependent() bool`

GetPeriodDependent returns the PeriodDependent field if non-nil, zero value otherwise.

### GetPeriodDependentOk

`func (o *RcwRequestsPutPreviousValues) GetPeriodDependentOk() (*bool, bool)`

GetPeriodDependentOk returns a tuple with the PeriodDependent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPeriodDependent

`func (o *RcwRequestsPutPreviousValues) SetPeriodDependent(v bool)`

SetPeriodDependent sets PeriodDependent field to given value.

### HasPeriodDependent

`func (o *RcwRequestsPutPreviousValues) HasPeriodDependent() bool`

HasPeriodDependent returns a boolean if a field has been set.

### GetDescription

`func (o *RcwRequestsPutPreviousValues) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *RcwRequestsPutPreviousValues) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *RcwRequestsPutPreviousValues) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *RcwRequestsPutPreviousValues) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetAdditionalInformation

`func (o *RcwRequestsPutPreviousValues) GetAdditionalInformation() string`

GetAdditionalInformation returns the AdditionalInformation field if non-nil, zero value otherwise.

### GetAdditionalInformationOk

`func (o *RcwRequestsPutPreviousValues) GetAdditionalInformationOk() (*string, bool)`

GetAdditionalInformationOk returns a tuple with the AdditionalInformation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdditionalInformation

`func (o *RcwRequestsPutPreviousValues) SetAdditionalInformation(v string)`

SetAdditionalInformation sets AdditionalInformation field to given value.

### HasAdditionalInformation

`func (o *RcwRequestsPutPreviousValues) HasAdditionalInformation() bool`

HasAdditionalInformation returns a boolean if a field has been set.

### GetScopes

`func (o *RcwRequestsPutPreviousValues) GetScopes() interface{}`

GetScopes returns the Scopes field if non-nil, zero value otherwise.

### GetScopesOk

`func (o *RcwRequestsPutPreviousValues) GetScopesOk() (*interface{}, bool)`

GetScopesOk returns a tuple with the Scopes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScopes

`func (o *RcwRequestsPutPreviousValues) SetScopes(v interface{})`

SetScopes sets Scopes field to given value.

### HasScopes

`func (o *RcwRequestsPutPreviousValues) HasScopes() bool`

HasScopes returns a boolean if a field has been set.

### SetScopesNil

`func (o *RcwRequestsPutPreviousValues) SetScopesNil(b bool)`

 SetScopesNil sets the value for Scopes to be an explicit nil

### UnsetScopes
`func (o *RcwRequestsPutPreviousValues) UnsetScopes()`

UnsetScopes ensures that no value is present for Scopes, not even an explicit nil
### GetIsRecurrence

`func (o *RcwRequestsPutPreviousValues) GetIsRecurrence() bool`

GetIsRecurrence returns the IsRecurrence field if non-nil, zero value otherwise.

### GetIsRecurrenceOk

`func (o *RcwRequestsPutPreviousValues) GetIsRecurrenceOk() (*bool, bool)`

GetIsRecurrenceOk returns a tuple with the IsRecurrence field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsRecurrence

`func (o *RcwRequestsPutPreviousValues) SetIsRecurrence(v bool)`

SetIsRecurrence sets IsRecurrence field to given value.

### HasIsRecurrence

`func (o *RcwRequestsPutPreviousValues) HasIsRecurrence() bool`

HasIsRecurrence returns a boolean if a field has been set.

### GetSubmittedByUserId

`func (o *RcwRequestsPutPreviousValues) GetSubmittedByUserId() int32`

GetSubmittedByUserId returns the SubmittedByUserId field if non-nil, zero value otherwise.

### GetSubmittedByUserIdOk

`func (o *RcwRequestsPutPreviousValues) GetSubmittedByUserIdOk() (*int32, bool)`

GetSubmittedByUserIdOk returns a tuple with the SubmittedByUserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubmittedByUserId

`func (o *RcwRequestsPutPreviousValues) SetSubmittedByUserId(v int32)`

SetSubmittedByUserId sets SubmittedByUserId field to given value.

### HasSubmittedByUserId

`func (o *RcwRequestsPutPreviousValues) HasSubmittedByUserId() bool`

HasSubmittedByUserId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


