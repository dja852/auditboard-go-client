# RcwRequestsPutRcwRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int32** | Note: This is a Primary Key.&lt;pk/&gt; | [optional] [readonly] 
**Title** | Pointer to **string** |  | [optional] 
**RcwLibraryRequestId** | **int32** | Note: This is a Foreign Key to &#x60;rcw_library_requests.id&#x60;.&lt;fk table&#x3D;&#39;rcw_library_requests&#39; column&#x3D;&#39;id&#39;/&gt; | 
**RcwProjectId** | Pointer to **int32** | Note: This is a Foreign Key to &#x60;rcw_projects.id&#x60;.&lt;fk table&#x3D;&#39;rcw_projects&#39; column&#x3D;&#39;id&#39;/&gt; | [optional] 
**AssigneeUserId** | Pointer to **int32** | Note: This is a Foreign Key to &#x60;users.id&#x60;.&lt;fk table&#x3D;&#39;users&#39; column&#x3D;&#39;id&#39;/&gt; | [optional] 
**Status** | **string** |  | 
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
**DidNotify** | **bool** |  | [default to false]
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

### NewRcwRequestsPutRcwRequest

`func NewRcwRequestsPutRcwRequest(rcwLibraryRequestId int32, status string, didNotify bool, ) *RcwRequestsPutRcwRequest`

NewRcwRequestsPutRcwRequest instantiates a new RcwRequestsPutRcwRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRcwRequestsPutRcwRequestWithDefaults

`func NewRcwRequestsPutRcwRequestWithDefaults() *RcwRequestsPutRcwRequest`

NewRcwRequestsPutRcwRequestWithDefaults instantiates a new RcwRequestsPutRcwRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *RcwRequestsPutRcwRequest) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *RcwRequestsPutRcwRequest) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *RcwRequestsPutRcwRequest) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *RcwRequestsPutRcwRequest) HasId() bool`

HasId returns a boolean if a field has been set.

### GetTitle

`func (o *RcwRequestsPutRcwRequest) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *RcwRequestsPutRcwRequest) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *RcwRequestsPutRcwRequest) SetTitle(v string)`

SetTitle sets Title field to given value.

### HasTitle

`func (o *RcwRequestsPutRcwRequest) HasTitle() bool`

HasTitle returns a boolean if a field has been set.

### GetRcwLibraryRequestId

`func (o *RcwRequestsPutRcwRequest) GetRcwLibraryRequestId() int32`

GetRcwLibraryRequestId returns the RcwLibraryRequestId field if non-nil, zero value otherwise.

### GetRcwLibraryRequestIdOk

`func (o *RcwRequestsPutRcwRequest) GetRcwLibraryRequestIdOk() (*int32, bool)`

GetRcwLibraryRequestIdOk returns a tuple with the RcwLibraryRequestId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRcwLibraryRequestId

`func (o *RcwRequestsPutRcwRequest) SetRcwLibraryRequestId(v int32)`

SetRcwLibraryRequestId sets RcwLibraryRequestId field to given value.


### GetRcwProjectId

`func (o *RcwRequestsPutRcwRequest) GetRcwProjectId() int32`

GetRcwProjectId returns the RcwProjectId field if non-nil, zero value otherwise.

### GetRcwProjectIdOk

`func (o *RcwRequestsPutRcwRequest) GetRcwProjectIdOk() (*int32, bool)`

GetRcwProjectIdOk returns a tuple with the RcwProjectId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRcwProjectId

`func (o *RcwRequestsPutRcwRequest) SetRcwProjectId(v int32)`

SetRcwProjectId sets RcwProjectId field to given value.

### HasRcwProjectId

`func (o *RcwRequestsPutRcwRequest) HasRcwProjectId() bool`

HasRcwProjectId returns a boolean if a field has been set.

### GetAssigneeUserId

`func (o *RcwRequestsPutRcwRequest) GetAssigneeUserId() int32`

GetAssigneeUserId returns the AssigneeUserId field if non-nil, zero value otherwise.

### GetAssigneeUserIdOk

`func (o *RcwRequestsPutRcwRequest) GetAssigneeUserIdOk() (*int32, bool)`

GetAssigneeUserIdOk returns a tuple with the AssigneeUserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssigneeUserId

`func (o *RcwRequestsPutRcwRequest) SetAssigneeUserId(v int32)`

SetAssigneeUserId sets AssigneeUserId field to given value.

### HasAssigneeUserId

`func (o *RcwRequestsPutRcwRequest) HasAssigneeUserId() bool`

HasAssigneeUserId returns a boolean if a field has been set.

### GetStatus

`func (o *RcwRequestsPutRcwRequest) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *RcwRequestsPutRcwRequest) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *RcwRequestsPutRcwRequest) SetStatus(v string)`

SetStatus sets Status field to given value.


### GetOpenDate

`func (o *RcwRequestsPutRcwRequest) GetOpenDate() string`

GetOpenDate returns the OpenDate field if non-nil, zero value otherwise.

### GetOpenDateOk

`func (o *RcwRequestsPutRcwRequest) GetOpenDateOk() (*string, bool)`

GetOpenDateOk returns a tuple with the OpenDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOpenDate

`func (o *RcwRequestsPutRcwRequest) SetOpenDate(v string)`

SetOpenDate sets OpenDate field to given value.

### HasOpenDate

`func (o *RcwRequestsPutRcwRequest) HasOpenDate() bool`

HasOpenDate returns a boolean if a field has been set.

### GetOpenByUserId

`func (o *RcwRequestsPutRcwRequest) GetOpenByUserId() int32`

GetOpenByUserId returns the OpenByUserId field if non-nil, zero value otherwise.

### GetOpenByUserIdOk

`func (o *RcwRequestsPutRcwRequest) GetOpenByUserIdOk() (*int32, bool)`

GetOpenByUserIdOk returns a tuple with the OpenByUserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOpenByUserId

`func (o *RcwRequestsPutRcwRequest) SetOpenByUserId(v int32)`

SetOpenByUserId sets OpenByUserId field to given value.

### HasOpenByUserId

`func (o *RcwRequestsPutRcwRequest) HasOpenByUserId() bool`

HasOpenByUserId returns a boolean if a field has been set.

### GetPendingReviewDate

`func (o *RcwRequestsPutRcwRequest) GetPendingReviewDate() string`

GetPendingReviewDate returns the PendingReviewDate field if non-nil, zero value otherwise.

### GetPendingReviewDateOk

`func (o *RcwRequestsPutRcwRequest) GetPendingReviewDateOk() (*string, bool)`

GetPendingReviewDateOk returns a tuple with the PendingReviewDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPendingReviewDate

`func (o *RcwRequestsPutRcwRequest) SetPendingReviewDate(v string)`

SetPendingReviewDate sets PendingReviewDate field to given value.

### HasPendingReviewDate

`func (o *RcwRequestsPutRcwRequest) HasPendingReviewDate() bool`

HasPendingReviewDate returns a boolean if a field has been set.

### GetPendingReviewByUserId

`func (o *RcwRequestsPutRcwRequest) GetPendingReviewByUserId() int32`

GetPendingReviewByUserId returns the PendingReviewByUserId field if non-nil, zero value otherwise.

### GetPendingReviewByUserIdOk

`func (o *RcwRequestsPutRcwRequest) GetPendingReviewByUserIdOk() (*int32, bool)`

GetPendingReviewByUserIdOk returns a tuple with the PendingReviewByUserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPendingReviewByUserId

`func (o *RcwRequestsPutRcwRequest) SetPendingReviewByUserId(v int32)`

SetPendingReviewByUserId sets PendingReviewByUserId field to given value.

### HasPendingReviewByUserId

`func (o *RcwRequestsPutRcwRequest) HasPendingReviewByUserId() bool`

HasPendingReviewByUserId returns a boolean if a field has been set.

### GetReopenedDate

`func (o *RcwRequestsPutRcwRequest) GetReopenedDate() string`

GetReopenedDate returns the ReopenedDate field if non-nil, zero value otherwise.

### GetReopenedDateOk

`func (o *RcwRequestsPutRcwRequest) GetReopenedDateOk() (*string, bool)`

GetReopenedDateOk returns a tuple with the ReopenedDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReopenedDate

`func (o *RcwRequestsPutRcwRequest) SetReopenedDate(v string)`

SetReopenedDate sets ReopenedDate field to given value.

### HasReopenedDate

`func (o *RcwRequestsPutRcwRequest) HasReopenedDate() bool`

HasReopenedDate returns a boolean if a field has been set.

### GetReopenedByUserId

`func (o *RcwRequestsPutRcwRequest) GetReopenedByUserId() int32`

GetReopenedByUserId returns the ReopenedByUserId field if non-nil, zero value otherwise.

### GetReopenedByUserIdOk

`func (o *RcwRequestsPutRcwRequest) GetReopenedByUserIdOk() (*int32, bool)`

GetReopenedByUserIdOk returns a tuple with the ReopenedByUserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReopenedByUserId

`func (o *RcwRequestsPutRcwRequest) SetReopenedByUserId(v int32)`

SetReopenedByUserId sets ReopenedByUserId field to given value.

### HasReopenedByUserId

`func (o *RcwRequestsPutRcwRequest) HasReopenedByUserId() bool`

HasReopenedByUserId returns a boolean if a field has been set.

### GetCompletedDate

`func (o *RcwRequestsPutRcwRequest) GetCompletedDate() string`

GetCompletedDate returns the CompletedDate field if non-nil, zero value otherwise.

### GetCompletedDateOk

`func (o *RcwRequestsPutRcwRequest) GetCompletedDateOk() (*string, bool)`

GetCompletedDateOk returns a tuple with the CompletedDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompletedDate

`func (o *RcwRequestsPutRcwRequest) SetCompletedDate(v string)`

SetCompletedDate sets CompletedDate field to given value.

### HasCompletedDate

`func (o *RcwRequestsPutRcwRequest) HasCompletedDate() bool`

HasCompletedDate returns a boolean if a field has been set.

### GetCompletedByUserId

`func (o *RcwRequestsPutRcwRequest) GetCompletedByUserId() int32`

GetCompletedByUserId returns the CompletedByUserId field if non-nil, zero value otherwise.

### GetCompletedByUserIdOk

`func (o *RcwRequestsPutRcwRequest) GetCompletedByUserIdOk() (*int32, bool)`

GetCompletedByUserIdOk returns a tuple with the CompletedByUserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompletedByUserId

`func (o *RcwRequestsPutRcwRequest) SetCompletedByUserId(v int32)`

SetCompletedByUserId sets CompletedByUserId field to given value.

### HasCompletedByUserId

`func (o *RcwRequestsPutRcwRequest) HasCompletedByUserId() bool`

HasCompletedByUserId returns a boolean if a field has been set.

### GetDueDate

`func (o *RcwRequestsPutRcwRequest) GetDueDate() string`

GetDueDate returns the DueDate field if non-nil, zero value otherwise.

### GetDueDateOk

`func (o *RcwRequestsPutRcwRequest) GetDueDateOk() (*string, bool)`

GetDueDateOk returns a tuple with the DueDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDueDate

`func (o *RcwRequestsPutRcwRequest) SetDueDate(v string)`

SetDueDate sets DueDate field to given value.

### HasDueDate

`func (o *RcwRequestsPutRcwRequest) HasDueDate() bool`

HasDueDate returns a boolean if a field has been set.

### GetCreatedAt

`func (o *RcwRequestsPutRcwRequest) GetCreatedAt() string`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *RcwRequestsPutRcwRequest) GetCreatedAtOk() (*string, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *RcwRequestsPutRcwRequest) SetCreatedAt(v string)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *RcwRequestsPutRcwRequest) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *RcwRequestsPutRcwRequest) GetUpdatedAt() string`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *RcwRequestsPutRcwRequest) GetUpdatedAtOk() (*string, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *RcwRequestsPutRcwRequest) SetUpdatedAt(v string)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *RcwRequestsPutRcwRequest) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.

### GetDeletedAt

`func (o *RcwRequestsPutRcwRequest) GetDeletedAt() string`

GetDeletedAt returns the DeletedAt field if non-nil, zero value otherwise.

### GetDeletedAtOk

`func (o *RcwRequestsPutRcwRequest) GetDeletedAtOk() (*string, bool)`

GetDeletedAtOk returns a tuple with the DeletedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeletedAt

`func (o *RcwRequestsPutRcwRequest) SetDeletedAt(v string)`

SetDeletedAt sets DeletedAt field to given value.

### HasDeletedAt

`func (o *RcwRequestsPutRcwRequest) HasDeletedAt() bool`

HasDeletedAt returns a boolean if a field has been set.

### GetRcwRequestRatingId

`func (o *RcwRequestsPutRcwRequest) GetRcwRequestRatingId() int32`

GetRcwRequestRatingId returns the RcwRequestRatingId field if non-nil, zero value otherwise.

### GetRcwRequestRatingIdOk

`func (o *RcwRequestsPutRcwRequest) GetRcwRequestRatingIdOk() (*int32, bool)`

GetRcwRequestRatingIdOk returns a tuple with the RcwRequestRatingId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRcwRequestRatingId

`func (o *RcwRequestsPutRcwRequest) SetRcwRequestRatingId(v int32)`

SetRcwRequestRatingId sets RcwRequestRatingId field to given value.

### HasRcwRequestRatingId

`func (o *RcwRequestsPutRcwRequest) HasRcwRequestRatingId() bool`

HasRcwRequestRatingId returns a boolean if a field has been set.

### GetDidNotify

`func (o *RcwRequestsPutRcwRequest) GetDidNotify() bool`

GetDidNotify returns the DidNotify field if non-nil, zero value otherwise.

### GetDidNotifyOk

`func (o *RcwRequestsPutRcwRequest) GetDidNotifyOk() (*bool, bool)`

GetDidNotifyOk returns a tuple with the DidNotify field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDidNotify

`func (o *RcwRequestsPutRcwRequest) SetDidNotify(v bool)`

SetDidNotify sets DidNotify field to given value.


### GetExplanationText

`func (o *RcwRequestsPutRcwRequest) GetExplanationText() string`

GetExplanationText returns the ExplanationText field if non-nil, zero value otherwise.

### GetExplanationTextOk

`func (o *RcwRequestsPutRcwRequest) GetExplanationTextOk() (*string, bool)`

GetExplanationTextOk returns a tuple with the ExplanationText field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExplanationText

`func (o *RcwRequestsPutRcwRequest) SetExplanationText(v string)`

SetExplanationText sets ExplanationText field to given value.

### HasExplanationText

`func (o *RcwRequestsPutRcwRequest) HasExplanationText() bool`

HasExplanationText returns a boolean if a field has been set.

### GetIsAutomated

`func (o *RcwRequestsPutRcwRequest) GetIsAutomated() bool`

GetIsAutomated returns the IsAutomated field if non-nil, zero value otherwise.

### GetIsAutomatedOk

`func (o *RcwRequestsPutRcwRequest) GetIsAutomatedOk() (*bool, bool)`

GetIsAutomatedOk returns a tuple with the IsAutomated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsAutomated

`func (o *RcwRequestsPutRcwRequest) SetIsAutomated(v bool)`

SetIsAutomated sets IsAutomated field to given value.

### HasIsAutomated

`func (o *RcwRequestsPutRcwRequest) HasIsAutomated() bool`

HasIsAutomated returns a boolean if a field has been set.

### GetLaunchDate

`func (o *RcwRequestsPutRcwRequest) GetLaunchDate() string`

GetLaunchDate returns the LaunchDate field if non-nil, zero value otherwise.

### GetLaunchDateOk

`func (o *RcwRequestsPutRcwRequest) GetLaunchDateOk() (*string, bool)`

GetLaunchDateOk returns a tuple with the LaunchDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLaunchDate

`func (o *RcwRequestsPutRcwRequest) SetLaunchDate(v string)`

SetLaunchDate sets LaunchDate field to given value.

### HasLaunchDate

`func (o *RcwRequestsPutRcwRequest) HasLaunchDate() bool`

HasLaunchDate returns a boolean if a field has been set.

### GetPeriodStart

`func (o *RcwRequestsPutRcwRequest) GetPeriodStart() string`

GetPeriodStart returns the PeriodStart field if non-nil, zero value otherwise.

### GetPeriodStartOk

`func (o *RcwRequestsPutRcwRequest) GetPeriodStartOk() (*string, bool)`

GetPeriodStartOk returns a tuple with the PeriodStart field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPeriodStart

`func (o *RcwRequestsPutRcwRequest) SetPeriodStart(v string)`

SetPeriodStart sets PeriodStart field to given value.

### HasPeriodStart

`func (o *RcwRequestsPutRcwRequest) HasPeriodStart() bool`

HasPeriodStart returns a boolean if a field has been set.

### GetPeriodEnd

`func (o *RcwRequestsPutRcwRequest) GetPeriodEnd() string`

GetPeriodEnd returns the PeriodEnd field if non-nil, zero value otherwise.

### GetPeriodEndOk

`func (o *RcwRequestsPutRcwRequest) GetPeriodEndOk() (*string, bool)`

GetPeriodEndOk returns a tuple with the PeriodEnd field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPeriodEnd

`func (o *RcwRequestsPutRcwRequest) SetPeriodEnd(v string)`

SetPeriodEnd sets PeriodEnd field to given value.

### HasPeriodEnd

`func (o *RcwRequestsPutRcwRequest) HasPeriodEnd() bool`

HasPeriodEnd returns a boolean if a field has been set.

### GetExternalIntegrationUrl

`func (o *RcwRequestsPutRcwRequest) GetExternalIntegrationUrl() string`

GetExternalIntegrationUrl returns the ExternalIntegrationUrl field if non-nil, zero value otherwise.

### GetExternalIntegrationUrlOk

`func (o *RcwRequestsPutRcwRequest) GetExternalIntegrationUrlOk() (*string, bool)`

GetExternalIntegrationUrlOk returns a tuple with the ExternalIntegrationUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalIntegrationUrl

`func (o *RcwRequestsPutRcwRequest) SetExternalIntegrationUrl(v string)`

SetExternalIntegrationUrl sets ExternalIntegrationUrl field to given value.

### HasExternalIntegrationUrl

`func (o *RcwRequestsPutRcwRequest) HasExternalIntegrationUrl() bool`

HasExternalIntegrationUrl returns a boolean if a field has been set.

### GetPeriodDependent

`func (o *RcwRequestsPutRcwRequest) GetPeriodDependent() bool`

GetPeriodDependent returns the PeriodDependent field if non-nil, zero value otherwise.

### GetPeriodDependentOk

`func (o *RcwRequestsPutRcwRequest) GetPeriodDependentOk() (*bool, bool)`

GetPeriodDependentOk returns a tuple with the PeriodDependent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPeriodDependent

`func (o *RcwRequestsPutRcwRequest) SetPeriodDependent(v bool)`

SetPeriodDependent sets PeriodDependent field to given value.

### HasPeriodDependent

`func (o *RcwRequestsPutRcwRequest) HasPeriodDependent() bool`

HasPeriodDependent returns a boolean if a field has been set.

### GetDescription

`func (o *RcwRequestsPutRcwRequest) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *RcwRequestsPutRcwRequest) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *RcwRequestsPutRcwRequest) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *RcwRequestsPutRcwRequest) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetAdditionalInformation

`func (o *RcwRequestsPutRcwRequest) GetAdditionalInformation() string`

GetAdditionalInformation returns the AdditionalInformation field if non-nil, zero value otherwise.

### GetAdditionalInformationOk

`func (o *RcwRequestsPutRcwRequest) GetAdditionalInformationOk() (*string, bool)`

GetAdditionalInformationOk returns a tuple with the AdditionalInformation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdditionalInformation

`func (o *RcwRequestsPutRcwRequest) SetAdditionalInformation(v string)`

SetAdditionalInformation sets AdditionalInformation field to given value.

### HasAdditionalInformation

`func (o *RcwRequestsPutRcwRequest) HasAdditionalInformation() bool`

HasAdditionalInformation returns a boolean if a field has been set.

### GetScopes

`func (o *RcwRequestsPutRcwRequest) GetScopes() interface{}`

GetScopes returns the Scopes field if non-nil, zero value otherwise.

### GetScopesOk

`func (o *RcwRequestsPutRcwRequest) GetScopesOk() (*interface{}, bool)`

GetScopesOk returns a tuple with the Scopes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScopes

`func (o *RcwRequestsPutRcwRequest) SetScopes(v interface{})`

SetScopes sets Scopes field to given value.

### HasScopes

`func (o *RcwRequestsPutRcwRequest) HasScopes() bool`

HasScopes returns a boolean if a field has been set.

### SetScopesNil

`func (o *RcwRequestsPutRcwRequest) SetScopesNil(b bool)`

 SetScopesNil sets the value for Scopes to be an explicit nil

### UnsetScopes
`func (o *RcwRequestsPutRcwRequest) UnsetScopes()`

UnsetScopes ensures that no value is present for Scopes, not even an explicit nil
### GetIsRecurrence

`func (o *RcwRequestsPutRcwRequest) GetIsRecurrence() bool`

GetIsRecurrence returns the IsRecurrence field if non-nil, zero value otherwise.

### GetIsRecurrenceOk

`func (o *RcwRequestsPutRcwRequest) GetIsRecurrenceOk() (*bool, bool)`

GetIsRecurrenceOk returns a tuple with the IsRecurrence field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsRecurrence

`func (o *RcwRequestsPutRcwRequest) SetIsRecurrence(v bool)`

SetIsRecurrence sets IsRecurrence field to given value.

### HasIsRecurrence

`func (o *RcwRequestsPutRcwRequest) HasIsRecurrence() bool`

HasIsRecurrence returns a boolean if a field has been set.

### GetSubmittedByUserId

`func (o *RcwRequestsPutRcwRequest) GetSubmittedByUserId() int32`

GetSubmittedByUserId returns the SubmittedByUserId field if non-nil, zero value otherwise.

### GetSubmittedByUserIdOk

`func (o *RcwRequestsPutRcwRequest) GetSubmittedByUserIdOk() (*int32, bool)`

GetSubmittedByUserIdOk returns a tuple with the SubmittedByUserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubmittedByUserId

`func (o *RcwRequestsPutRcwRequest) SetSubmittedByUserId(v int32)`

SetSubmittedByUserId sets SubmittedByUserId field to given value.

### HasSubmittedByUserId

`func (o *RcwRequestsPutRcwRequest) HasSubmittedByUserId() bool`

HasSubmittedByUserId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


