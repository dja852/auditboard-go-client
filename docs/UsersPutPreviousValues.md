# UsersPutPreviousValues

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int32** | Note: This is a Primary Key.&lt;pk/&gt; | [optional] [readonly] 
**CreatedAt** | Pointer to **string** |  | [optional] [readonly] 
**UpdatedAt** | Pointer to **string** |  | [optional] [readonly] 
**DeletedAt** | Pointer to **string** |  | [optional] [readonly] 
**FirstName** | Pointer to **string** |  | [optional] 
**LastName** | Pointer to **string** |  | [optional] 
**Email** | Pointer to **string** |  | [optional] 
**Password** | Pointer to **string** |  | [optional] 
**Salt** | Pointer to **string** |  | [optional] 
**RoleId** | Pointer to **int32** | Note: This is a Foreign Key to &#x60;roles.id&#x60;.&lt;fk table&#x3D;&#39;roles&#39; column&#x3D;&#39;id&#39;/&gt; | [optional] 
**AcceptedTerms** | Pointer to **bool** |  | [optional] 
**LastLogin** | Pointer to **string** |  | [optional] 
**LastPasswordChange** | Pointer to **string** |  | [optional] 
**LastIpAddress** | Pointer to **string** |  | [optional] 
**PasswordHistory** | Pointer to **string** |  | [optional] 
**VerifiedIpHistory** | Pointer to **string** |  | [optional] 
**RequiresPasswordChange** | Pointer to **bool** |  | [optional] 
**IconBgColor** | Pointer to **string** |  | [optional] 
**IconFgColor** | Pointer to **string** |  | [optional] 
**CompetencyEvaluation** | Pointer to **string** |  | [optional] 
**UserStatId** | Pointer to **int32** | Note: This is a Foreign Key to &#x60;user_stats.id&#x60;.&lt;fk table&#x3D;&#39;user_stats&#39; column&#x3D;&#39;id&#39;/&gt; | [optional] 
**PasswordResetToken** | Pointer to **string** |  | [optional] 
**PasswordResetTokenExpires** | Pointer to **string** |  | [optional] 
**AuthType** | Pointer to **string** |  | [optional] 
**LocationId** | Pointer to **int32** | Note: This is a Foreign Key to &#x60;locations.id&#x60;.&lt;fk table&#x3D;&#39;locations&#39; column&#x3D;&#39;id&#39;/&gt; | [optional] 
**DepartmentId** | Pointer to **int32** | Note: This is a Foreign Key to &#x60;departments.id&#x60;.&lt;fk table&#x3D;&#39;departments&#39; column&#x3D;&#39;id&#39;/&gt; | [optional] 
**JobTitle** | Pointer to **string** |  | [optional] 
**MfaEnabled** | Pointer to **bool** |  | [optional] [default to false]
**CellPhone** | Pointer to **string** |  | [optional] [default to ""]
**AuthyId** | Pointer to **int32** |  | [optional] 
**EmployeeNumber** | Pointer to **string** |  | [optional] 
**Username** | Pointer to **string** |  | [optional] 
**LoginAttempt** | Pointer to **int32** |  | [optional] [default to 0]
**LastLoginFailedAttempt** | Pointer to **string** |  | [optional] 
**FirstLogin** | Pointer to **string** |  | [optional] 
**ChameleonIdentifier** | Pointer to **string** |  | [optional] 
**SlackUserId** | Pointer to **string** |  | [optional] 
**SentSlackWelcomeNotification** | Pointer to **bool** |  | [optional] 
**AvatarKey** | Pointer to **string** |  | [optional] 
**ExternalApiToken** | Pointer to **string** |  | [optional] 
**AzureConversationData** | Pointer to **interface{}** |  | [optional] 
**ScimUserActive** | Pointer to **string** |  | [optional] 
**TeamsNotificationsEnabled** | Pointer to **bool** |  | [optional] 
**CoreModules** | Pointer to **interface{}** |  | [optional] 
**Locale** | Pointer to **string** |  | [optional] [default to "NULL"]
**AccountExpirationDate** | Pointer to **string** |  | [optional] 

## Methods

### NewUsersPutPreviousValues

`func NewUsersPutPreviousValues() *UsersPutPreviousValues`

NewUsersPutPreviousValues instantiates a new UsersPutPreviousValues object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUsersPutPreviousValuesWithDefaults

`func NewUsersPutPreviousValuesWithDefaults() *UsersPutPreviousValues`

NewUsersPutPreviousValuesWithDefaults instantiates a new UsersPutPreviousValues object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *UsersPutPreviousValues) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *UsersPutPreviousValues) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *UsersPutPreviousValues) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *UsersPutPreviousValues) HasId() bool`

HasId returns a boolean if a field has been set.

### GetCreatedAt

`func (o *UsersPutPreviousValues) GetCreatedAt() string`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *UsersPutPreviousValues) GetCreatedAtOk() (*string, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *UsersPutPreviousValues) SetCreatedAt(v string)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *UsersPutPreviousValues) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *UsersPutPreviousValues) GetUpdatedAt() string`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *UsersPutPreviousValues) GetUpdatedAtOk() (*string, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *UsersPutPreviousValues) SetUpdatedAt(v string)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *UsersPutPreviousValues) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.

### GetDeletedAt

`func (o *UsersPutPreviousValues) GetDeletedAt() string`

GetDeletedAt returns the DeletedAt field if non-nil, zero value otherwise.

### GetDeletedAtOk

`func (o *UsersPutPreviousValues) GetDeletedAtOk() (*string, bool)`

GetDeletedAtOk returns a tuple with the DeletedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeletedAt

`func (o *UsersPutPreviousValues) SetDeletedAt(v string)`

SetDeletedAt sets DeletedAt field to given value.

### HasDeletedAt

`func (o *UsersPutPreviousValues) HasDeletedAt() bool`

HasDeletedAt returns a boolean if a field has been set.

### GetFirstName

`func (o *UsersPutPreviousValues) GetFirstName() string`

GetFirstName returns the FirstName field if non-nil, zero value otherwise.

### GetFirstNameOk

`func (o *UsersPutPreviousValues) GetFirstNameOk() (*string, bool)`

GetFirstNameOk returns a tuple with the FirstName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFirstName

`func (o *UsersPutPreviousValues) SetFirstName(v string)`

SetFirstName sets FirstName field to given value.

### HasFirstName

`func (o *UsersPutPreviousValues) HasFirstName() bool`

HasFirstName returns a boolean if a field has been set.

### GetLastName

`func (o *UsersPutPreviousValues) GetLastName() string`

GetLastName returns the LastName field if non-nil, zero value otherwise.

### GetLastNameOk

`func (o *UsersPutPreviousValues) GetLastNameOk() (*string, bool)`

GetLastNameOk returns a tuple with the LastName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastName

`func (o *UsersPutPreviousValues) SetLastName(v string)`

SetLastName sets LastName field to given value.

### HasLastName

`func (o *UsersPutPreviousValues) HasLastName() bool`

HasLastName returns a boolean if a field has been set.

### GetEmail

`func (o *UsersPutPreviousValues) GetEmail() string`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *UsersPutPreviousValues) GetEmailOk() (*string, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *UsersPutPreviousValues) SetEmail(v string)`

SetEmail sets Email field to given value.

### HasEmail

`func (o *UsersPutPreviousValues) HasEmail() bool`

HasEmail returns a boolean if a field has been set.

### GetPassword

`func (o *UsersPutPreviousValues) GetPassword() string`

GetPassword returns the Password field if non-nil, zero value otherwise.

### GetPasswordOk

`func (o *UsersPutPreviousValues) GetPasswordOk() (*string, bool)`

GetPasswordOk returns a tuple with the Password field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPassword

`func (o *UsersPutPreviousValues) SetPassword(v string)`

SetPassword sets Password field to given value.

### HasPassword

`func (o *UsersPutPreviousValues) HasPassword() bool`

HasPassword returns a boolean if a field has been set.

### GetSalt

`func (o *UsersPutPreviousValues) GetSalt() string`

GetSalt returns the Salt field if non-nil, zero value otherwise.

### GetSaltOk

`func (o *UsersPutPreviousValues) GetSaltOk() (*string, bool)`

GetSaltOk returns a tuple with the Salt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSalt

`func (o *UsersPutPreviousValues) SetSalt(v string)`

SetSalt sets Salt field to given value.

### HasSalt

`func (o *UsersPutPreviousValues) HasSalt() bool`

HasSalt returns a boolean if a field has been set.

### GetRoleId

`func (o *UsersPutPreviousValues) GetRoleId() int32`

GetRoleId returns the RoleId field if non-nil, zero value otherwise.

### GetRoleIdOk

`func (o *UsersPutPreviousValues) GetRoleIdOk() (*int32, bool)`

GetRoleIdOk returns a tuple with the RoleId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRoleId

`func (o *UsersPutPreviousValues) SetRoleId(v int32)`

SetRoleId sets RoleId field to given value.

### HasRoleId

`func (o *UsersPutPreviousValues) HasRoleId() bool`

HasRoleId returns a boolean if a field has been set.

### GetAcceptedTerms

`func (o *UsersPutPreviousValues) GetAcceptedTerms() bool`

GetAcceptedTerms returns the AcceptedTerms field if non-nil, zero value otherwise.

### GetAcceptedTermsOk

`func (o *UsersPutPreviousValues) GetAcceptedTermsOk() (*bool, bool)`

GetAcceptedTermsOk returns a tuple with the AcceptedTerms field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAcceptedTerms

`func (o *UsersPutPreviousValues) SetAcceptedTerms(v bool)`

SetAcceptedTerms sets AcceptedTerms field to given value.

### HasAcceptedTerms

`func (o *UsersPutPreviousValues) HasAcceptedTerms() bool`

HasAcceptedTerms returns a boolean if a field has been set.

### GetLastLogin

`func (o *UsersPutPreviousValues) GetLastLogin() string`

GetLastLogin returns the LastLogin field if non-nil, zero value otherwise.

### GetLastLoginOk

`func (o *UsersPutPreviousValues) GetLastLoginOk() (*string, bool)`

GetLastLoginOk returns a tuple with the LastLogin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastLogin

`func (o *UsersPutPreviousValues) SetLastLogin(v string)`

SetLastLogin sets LastLogin field to given value.

### HasLastLogin

`func (o *UsersPutPreviousValues) HasLastLogin() bool`

HasLastLogin returns a boolean if a field has been set.

### GetLastPasswordChange

`func (o *UsersPutPreviousValues) GetLastPasswordChange() string`

GetLastPasswordChange returns the LastPasswordChange field if non-nil, zero value otherwise.

### GetLastPasswordChangeOk

`func (o *UsersPutPreviousValues) GetLastPasswordChangeOk() (*string, bool)`

GetLastPasswordChangeOk returns a tuple with the LastPasswordChange field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastPasswordChange

`func (o *UsersPutPreviousValues) SetLastPasswordChange(v string)`

SetLastPasswordChange sets LastPasswordChange field to given value.

### HasLastPasswordChange

`func (o *UsersPutPreviousValues) HasLastPasswordChange() bool`

HasLastPasswordChange returns a boolean if a field has been set.

### GetLastIpAddress

`func (o *UsersPutPreviousValues) GetLastIpAddress() string`

GetLastIpAddress returns the LastIpAddress field if non-nil, zero value otherwise.

### GetLastIpAddressOk

`func (o *UsersPutPreviousValues) GetLastIpAddressOk() (*string, bool)`

GetLastIpAddressOk returns a tuple with the LastIpAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastIpAddress

`func (o *UsersPutPreviousValues) SetLastIpAddress(v string)`

SetLastIpAddress sets LastIpAddress field to given value.

### HasLastIpAddress

`func (o *UsersPutPreviousValues) HasLastIpAddress() bool`

HasLastIpAddress returns a boolean if a field has been set.

### GetPasswordHistory

`func (o *UsersPutPreviousValues) GetPasswordHistory() string`

GetPasswordHistory returns the PasswordHistory field if non-nil, zero value otherwise.

### GetPasswordHistoryOk

`func (o *UsersPutPreviousValues) GetPasswordHistoryOk() (*string, bool)`

GetPasswordHistoryOk returns a tuple with the PasswordHistory field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPasswordHistory

`func (o *UsersPutPreviousValues) SetPasswordHistory(v string)`

SetPasswordHistory sets PasswordHistory field to given value.

### HasPasswordHistory

`func (o *UsersPutPreviousValues) HasPasswordHistory() bool`

HasPasswordHistory returns a boolean if a field has been set.

### GetVerifiedIpHistory

`func (o *UsersPutPreviousValues) GetVerifiedIpHistory() string`

GetVerifiedIpHistory returns the VerifiedIpHistory field if non-nil, zero value otherwise.

### GetVerifiedIpHistoryOk

`func (o *UsersPutPreviousValues) GetVerifiedIpHistoryOk() (*string, bool)`

GetVerifiedIpHistoryOk returns a tuple with the VerifiedIpHistory field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVerifiedIpHistory

`func (o *UsersPutPreviousValues) SetVerifiedIpHistory(v string)`

SetVerifiedIpHistory sets VerifiedIpHistory field to given value.

### HasVerifiedIpHistory

`func (o *UsersPutPreviousValues) HasVerifiedIpHistory() bool`

HasVerifiedIpHistory returns a boolean if a field has been set.

### GetRequiresPasswordChange

`func (o *UsersPutPreviousValues) GetRequiresPasswordChange() bool`

GetRequiresPasswordChange returns the RequiresPasswordChange field if non-nil, zero value otherwise.

### GetRequiresPasswordChangeOk

`func (o *UsersPutPreviousValues) GetRequiresPasswordChangeOk() (*bool, bool)`

GetRequiresPasswordChangeOk returns a tuple with the RequiresPasswordChange field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequiresPasswordChange

`func (o *UsersPutPreviousValues) SetRequiresPasswordChange(v bool)`

SetRequiresPasswordChange sets RequiresPasswordChange field to given value.

### HasRequiresPasswordChange

`func (o *UsersPutPreviousValues) HasRequiresPasswordChange() bool`

HasRequiresPasswordChange returns a boolean if a field has been set.

### GetIconBgColor

`func (o *UsersPutPreviousValues) GetIconBgColor() string`

GetIconBgColor returns the IconBgColor field if non-nil, zero value otherwise.

### GetIconBgColorOk

`func (o *UsersPutPreviousValues) GetIconBgColorOk() (*string, bool)`

GetIconBgColorOk returns a tuple with the IconBgColor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIconBgColor

`func (o *UsersPutPreviousValues) SetIconBgColor(v string)`

SetIconBgColor sets IconBgColor field to given value.

### HasIconBgColor

`func (o *UsersPutPreviousValues) HasIconBgColor() bool`

HasIconBgColor returns a boolean if a field has been set.

### GetIconFgColor

`func (o *UsersPutPreviousValues) GetIconFgColor() string`

GetIconFgColor returns the IconFgColor field if non-nil, zero value otherwise.

### GetIconFgColorOk

`func (o *UsersPutPreviousValues) GetIconFgColorOk() (*string, bool)`

GetIconFgColorOk returns a tuple with the IconFgColor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIconFgColor

`func (o *UsersPutPreviousValues) SetIconFgColor(v string)`

SetIconFgColor sets IconFgColor field to given value.

### HasIconFgColor

`func (o *UsersPutPreviousValues) HasIconFgColor() bool`

HasIconFgColor returns a boolean if a field has been set.

### GetCompetencyEvaluation

`func (o *UsersPutPreviousValues) GetCompetencyEvaluation() string`

GetCompetencyEvaluation returns the CompetencyEvaluation field if non-nil, zero value otherwise.

### GetCompetencyEvaluationOk

`func (o *UsersPutPreviousValues) GetCompetencyEvaluationOk() (*string, bool)`

GetCompetencyEvaluationOk returns a tuple with the CompetencyEvaluation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompetencyEvaluation

`func (o *UsersPutPreviousValues) SetCompetencyEvaluation(v string)`

SetCompetencyEvaluation sets CompetencyEvaluation field to given value.

### HasCompetencyEvaluation

`func (o *UsersPutPreviousValues) HasCompetencyEvaluation() bool`

HasCompetencyEvaluation returns a boolean if a field has been set.

### GetUserStatId

`func (o *UsersPutPreviousValues) GetUserStatId() int32`

GetUserStatId returns the UserStatId field if non-nil, zero value otherwise.

### GetUserStatIdOk

`func (o *UsersPutPreviousValues) GetUserStatIdOk() (*int32, bool)`

GetUserStatIdOk returns a tuple with the UserStatId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserStatId

`func (o *UsersPutPreviousValues) SetUserStatId(v int32)`

SetUserStatId sets UserStatId field to given value.

### HasUserStatId

`func (o *UsersPutPreviousValues) HasUserStatId() bool`

HasUserStatId returns a boolean if a field has been set.

### GetPasswordResetToken

`func (o *UsersPutPreviousValues) GetPasswordResetToken() string`

GetPasswordResetToken returns the PasswordResetToken field if non-nil, zero value otherwise.

### GetPasswordResetTokenOk

`func (o *UsersPutPreviousValues) GetPasswordResetTokenOk() (*string, bool)`

GetPasswordResetTokenOk returns a tuple with the PasswordResetToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPasswordResetToken

`func (o *UsersPutPreviousValues) SetPasswordResetToken(v string)`

SetPasswordResetToken sets PasswordResetToken field to given value.

### HasPasswordResetToken

`func (o *UsersPutPreviousValues) HasPasswordResetToken() bool`

HasPasswordResetToken returns a boolean if a field has been set.

### GetPasswordResetTokenExpires

`func (o *UsersPutPreviousValues) GetPasswordResetTokenExpires() string`

GetPasswordResetTokenExpires returns the PasswordResetTokenExpires field if non-nil, zero value otherwise.

### GetPasswordResetTokenExpiresOk

`func (o *UsersPutPreviousValues) GetPasswordResetTokenExpiresOk() (*string, bool)`

GetPasswordResetTokenExpiresOk returns a tuple with the PasswordResetTokenExpires field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPasswordResetTokenExpires

`func (o *UsersPutPreviousValues) SetPasswordResetTokenExpires(v string)`

SetPasswordResetTokenExpires sets PasswordResetTokenExpires field to given value.

### HasPasswordResetTokenExpires

`func (o *UsersPutPreviousValues) HasPasswordResetTokenExpires() bool`

HasPasswordResetTokenExpires returns a boolean if a field has been set.

### GetAuthType

`func (o *UsersPutPreviousValues) GetAuthType() string`

GetAuthType returns the AuthType field if non-nil, zero value otherwise.

### GetAuthTypeOk

`func (o *UsersPutPreviousValues) GetAuthTypeOk() (*string, bool)`

GetAuthTypeOk returns a tuple with the AuthType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthType

`func (o *UsersPutPreviousValues) SetAuthType(v string)`

SetAuthType sets AuthType field to given value.

### HasAuthType

`func (o *UsersPutPreviousValues) HasAuthType() bool`

HasAuthType returns a boolean if a field has been set.

### GetLocationId

`func (o *UsersPutPreviousValues) GetLocationId() int32`

GetLocationId returns the LocationId field if non-nil, zero value otherwise.

### GetLocationIdOk

`func (o *UsersPutPreviousValues) GetLocationIdOk() (*int32, bool)`

GetLocationIdOk returns a tuple with the LocationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocationId

`func (o *UsersPutPreviousValues) SetLocationId(v int32)`

SetLocationId sets LocationId field to given value.

### HasLocationId

`func (o *UsersPutPreviousValues) HasLocationId() bool`

HasLocationId returns a boolean if a field has been set.

### GetDepartmentId

`func (o *UsersPutPreviousValues) GetDepartmentId() int32`

GetDepartmentId returns the DepartmentId field if non-nil, zero value otherwise.

### GetDepartmentIdOk

`func (o *UsersPutPreviousValues) GetDepartmentIdOk() (*int32, bool)`

GetDepartmentIdOk returns a tuple with the DepartmentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDepartmentId

`func (o *UsersPutPreviousValues) SetDepartmentId(v int32)`

SetDepartmentId sets DepartmentId field to given value.

### HasDepartmentId

`func (o *UsersPutPreviousValues) HasDepartmentId() bool`

HasDepartmentId returns a boolean if a field has been set.

### GetJobTitle

`func (o *UsersPutPreviousValues) GetJobTitle() string`

GetJobTitle returns the JobTitle field if non-nil, zero value otherwise.

### GetJobTitleOk

`func (o *UsersPutPreviousValues) GetJobTitleOk() (*string, bool)`

GetJobTitleOk returns a tuple with the JobTitle field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJobTitle

`func (o *UsersPutPreviousValues) SetJobTitle(v string)`

SetJobTitle sets JobTitle field to given value.

### HasJobTitle

`func (o *UsersPutPreviousValues) HasJobTitle() bool`

HasJobTitle returns a boolean if a field has been set.

### GetMfaEnabled

`func (o *UsersPutPreviousValues) GetMfaEnabled() bool`

GetMfaEnabled returns the MfaEnabled field if non-nil, zero value otherwise.

### GetMfaEnabledOk

`func (o *UsersPutPreviousValues) GetMfaEnabledOk() (*bool, bool)`

GetMfaEnabledOk returns a tuple with the MfaEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMfaEnabled

`func (o *UsersPutPreviousValues) SetMfaEnabled(v bool)`

SetMfaEnabled sets MfaEnabled field to given value.

### HasMfaEnabled

`func (o *UsersPutPreviousValues) HasMfaEnabled() bool`

HasMfaEnabled returns a boolean if a field has been set.

### GetCellPhone

`func (o *UsersPutPreviousValues) GetCellPhone() string`

GetCellPhone returns the CellPhone field if non-nil, zero value otherwise.

### GetCellPhoneOk

`func (o *UsersPutPreviousValues) GetCellPhoneOk() (*string, bool)`

GetCellPhoneOk returns a tuple with the CellPhone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCellPhone

`func (o *UsersPutPreviousValues) SetCellPhone(v string)`

SetCellPhone sets CellPhone field to given value.

### HasCellPhone

`func (o *UsersPutPreviousValues) HasCellPhone() bool`

HasCellPhone returns a boolean if a field has been set.

### GetAuthyId

`func (o *UsersPutPreviousValues) GetAuthyId() int32`

GetAuthyId returns the AuthyId field if non-nil, zero value otherwise.

### GetAuthyIdOk

`func (o *UsersPutPreviousValues) GetAuthyIdOk() (*int32, bool)`

GetAuthyIdOk returns a tuple with the AuthyId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthyId

`func (o *UsersPutPreviousValues) SetAuthyId(v int32)`

SetAuthyId sets AuthyId field to given value.

### HasAuthyId

`func (o *UsersPutPreviousValues) HasAuthyId() bool`

HasAuthyId returns a boolean if a field has been set.

### GetEmployeeNumber

`func (o *UsersPutPreviousValues) GetEmployeeNumber() string`

GetEmployeeNumber returns the EmployeeNumber field if non-nil, zero value otherwise.

### GetEmployeeNumberOk

`func (o *UsersPutPreviousValues) GetEmployeeNumberOk() (*string, bool)`

GetEmployeeNumberOk returns a tuple with the EmployeeNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmployeeNumber

`func (o *UsersPutPreviousValues) SetEmployeeNumber(v string)`

SetEmployeeNumber sets EmployeeNumber field to given value.

### HasEmployeeNumber

`func (o *UsersPutPreviousValues) HasEmployeeNumber() bool`

HasEmployeeNumber returns a boolean if a field has been set.

### GetUsername

`func (o *UsersPutPreviousValues) GetUsername() string`

GetUsername returns the Username field if non-nil, zero value otherwise.

### GetUsernameOk

`func (o *UsersPutPreviousValues) GetUsernameOk() (*string, bool)`

GetUsernameOk returns a tuple with the Username field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsername

`func (o *UsersPutPreviousValues) SetUsername(v string)`

SetUsername sets Username field to given value.

### HasUsername

`func (o *UsersPutPreviousValues) HasUsername() bool`

HasUsername returns a boolean if a field has been set.

### GetLoginAttempt

`func (o *UsersPutPreviousValues) GetLoginAttempt() int32`

GetLoginAttempt returns the LoginAttempt field if non-nil, zero value otherwise.

### GetLoginAttemptOk

`func (o *UsersPutPreviousValues) GetLoginAttemptOk() (*int32, bool)`

GetLoginAttemptOk returns a tuple with the LoginAttempt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLoginAttempt

`func (o *UsersPutPreviousValues) SetLoginAttempt(v int32)`

SetLoginAttempt sets LoginAttempt field to given value.

### HasLoginAttempt

`func (o *UsersPutPreviousValues) HasLoginAttempt() bool`

HasLoginAttempt returns a boolean if a field has been set.

### GetLastLoginFailedAttempt

`func (o *UsersPutPreviousValues) GetLastLoginFailedAttempt() string`

GetLastLoginFailedAttempt returns the LastLoginFailedAttempt field if non-nil, zero value otherwise.

### GetLastLoginFailedAttemptOk

`func (o *UsersPutPreviousValues) GetLastLoginFailedAttemptOk() (*string, bool)`

GetLastLoginFailedAttemptOk returns a tuple with the LastLoginFailedAttempt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastLoginFailedAttempt

`func (o *UsersPutPreviousValues) SetLastLoginFailedAttempt(v string)`

SetLastLoginFailedAttempt sets LastLoginFailedAttempt field to given value.

### HasLastLoginFailedAttempt

`func (o *UsersPutPreviousValues) HasLastLoginFailedAttempt() bool`

HasLastLoginFailedAttempt returns a boolean if a field has been set.

### GetFirstLogin

`func (o *UsersPutPreviousValues) GetFirstLogin() string`

GetFirstLogin returns the FirstLogin field if non-nil, zero value otherwise.

### GetFirstLoginOk

`func (o *UsersPutPreviousValues) GetFirstLoginOk() (*string, bool)`

GetFirstLoginOk returns a tuple with the FirstLogin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFirstLogin

`func (o *UsersPutPreviousValues) SetFirstLogin(v string)`

SetFirstLogin sets FirstLogin field to given value.

### HasFirstLogin

`func (o *UsersPutPreviousValues) HasFirstLogin() bool`

HasFirstLogin returns a boolean if a field has been set.

### GetChameleonIdentifier

`func (o *UsersPutPreviousValues) GetChameleonIdentifier() string`

GetChameleonIdentifier returns the ChameleonIdentifier field if non-nil, zero value otherwise.

### GetChameleonIdentifierOk

`func (o *UsersPutPreviousValues) GetChameleonIdentifierOk() (*string, bool)`

GetChameleonIdentifierOk returns a tuple with the ChameleonIdentifier field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChameleonIdentifier

`func (o *UsersPutPreviousValues) SetChameleonIdentifier(v string)`

SetChameleonIdentifier sets ChameleonIdentifier field to given value.

### HasChameleonIdentifier

`func (o *UsersPutPreviousValues) HasChameleonIdentifier() bool`

HasChameleonIdentifier returns a boolean if a field has been set.

### GetSlackUserId

`func (o *UsersPutPreviousValues) GetSlackUserId() string`

GetSlackUserId returns the SlackUserId field if non-nil, zero value otherwise.

### GetSlackUserIdOk

`func (o *UsersPutPreviousValues) GetSlackUserIdOk() (*string, bool)`

GetSlackUserIdOk returns a tuple with the SlackUserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSlackUserId

`func (o *UsersPutPreviousValues) SetSlackUserId(v string)`

SetSlackUserId sets SlackUserId field to given value.

### HasSlackUserId

`func (o *UsersPutPreviousValues) HasSlackUserId() bool`

HasSlackUserId returns a boolean if a field has been set.

### GetSentSlackWelcomeNotification

`func (o *UsersPutPreviousValues) GetSentSlackWelcomeNotification() bool`

GetSentSlackWelcomeNotification returns the SentSlackWelcomeNotification field if non-nil, zero value otherwise.

### GetSentSlackWelcomeNotificationOk

`func (o *UsersPutPreviousValues) GetSentSlackWelcomeNotificationOk() (*bool, bool)`

GetSentSlackWelcomeNotificationOk returns a tuple with the SentSlackWelcomeNotification field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSentSlackWelcomeNotification

`func (o *UsersPutPreviousValues) SetSentSlackWelcomeNotification(v bool)`

SetSentSlackWelcomeNotification sets SentSlackWelcomeNotification field to given value.

### HasSentSlackWelcomeNotification

`func (o *UsersPutPreviousValues) HasSentSlackWelcomeNotification() bool`

HasSentSlackWelcomeNotification returns a boolean if a field has been set.

### GetAvatarKey

`func (o *UsersPutPreviousValues) GetAvatarKey() string`

GetAvatarKey returns the AvatarKey field if non-nil, zero value otherwise.

### GetAvatarKeyOk

`func (o *UsersPutPreviousValues) GetAvatarKeyOk() (*string, bool)`

GetAvatarKeyOk returns a tuple with the AvatarKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvatarKey

`func (o *UsersPutPreviousValues) SetAvatarKey(v string)`

SetAvatarKey sets AvatarKey field to given value.

### HasAvatarKey

`func (o *UsersPutPreviousValues) HasAvatarKey() bool`

HasAvatarKey returns a boolean if a field has been set.

### GetExternalApiToken

`func (o *UsersPutPreviousValues) GetExternalApiToken() string`

GetExternalApiToken returns the ExternalApiToken field if non-nil, zero value otherwise.

### GetExternalApiTokenOk

`func (o *UsersPutPreviousValues) GetExternalApiTokenOk() (*string, bool)`

GetExternalApiTokenOk returns a tuple with the ExternalApiToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalApiToken

`func (o *UsersPutPreviousValues) SetExternalApiToken(v string)`

SetExternalApiToken sets ExternalApiToken field to given value.

### HasExternalApiToken

`func (o *UsersPutPreviousValues) HasExternalApiToken() bool`

HasExternalApiToken returns a boolean if a field has been set.

### GetAzureConversationData

`func (o *UsersPutPreviousValues) GetAzureConversationData() interface{}`

GetAzureConversationData returns the AzureConversationData field if non-nil, zero value otherwise.

### GetAzureConversationDataOk

`func (o *UsersPutPreviousValues) GetAzureConversationDataOk() (*interface{}, bool)`

GetAzureConversationDataOk returns a tuple with the AzureConversationData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAzureConversationData

`func (o *UsersPutPreviousValues) SetAzureConversationData(v interface{})`

SetAzureConversationData sets AzureConversationData field to given value.

### HasAzureConversationData

`func (o *UsersPutPreviousValues) HasAzureConversationData() bool`

HasAzureConversationData returns a boolean if a field has been set.

### SetAzureConversationDataNil

`func (o *UsersPutPreviousValues) SetAzureConversationDataNil(b bool)`

 SetAzureConversationDataNil sets the value for AzureConversationData to be an explicit nil

### UnsetAzureConversationData
`func (o *UsersPutPreviousValues) UnsetAzureConversationData()`

UnsetAzureConversationData ensures that no value is present for AzureConversationData, not even an explicit nil
### GetScimUserActive

`func (o *UsersPutPreviousValues) GetScimUserActive() string`

GetScimUserActive returns the ScimUserActive field if non-nil, zero value otherwise.

### GetScimUserActiveOk

`func (o *UsersPutPreviousValues) GetScimUserActiveOk() (*string, bool)`

GetScimUserActiveOk returns a tuple with the ScimUserActive field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScimUserActive

`func (o *UsersPutPreviousValues) SetScimUserActive(v string)`

SetScimUserActive sets ScimUserActive field to given value.

### HasScimUserActive

`func (o *UsersPutPreviousValues) HasScimUserActive() bool`

HasScimUserActive returns a boolean if a field has been set.

### GetTeamsNotificationsEnabled

`func (o *UsersPutPreviousValues) GetTeamsNotificationsEnabled() bool`

GetTeamsNotificationsEnabled returns the TeamsNotificationsEnabled field if non-nil, zero value otherwise.

### GetTeamsNotificationsEnabledOk

`func (o *UsersPutPreviousValues) GetTeamsNotificationsEnabledOk() (*bool, bool)`

GetTeamsNotificationsEnabledOk returns a tuple with the TeamsNotificationsEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTeamsNotificationsEnabled

`func (o *UsersPutPreviousValues) SetTeamsNotificationsEnabled(v bool)`

SetTeamsNotificationsEnabled sets TeamsNotificationsEnabled field to given value.

### HasTeamsNotificationsEnabled

`func (o *UsersPutPreviousValues) HasTeamsNotificationsEnabled() bool`

HasTeamsNotificationsEnabled returns a boolean if a field has been set.

### GetCoreModules

`func (o *UsersPutPreviousValues) GetCoreModules() interface{}`

GetCoreModules returns the CoreModules field if non-nil, zero value otherwise.

### GetCoreModulesOk

`func (o *UsersPutPreviousValues) GetCoreModulesOk() (*interface{}, bool)`

GetCoreModulesOk returns a tuple with the CoreModules field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCoreModules

`func (o *UsersPutPreviousValues) SetCoreModules(v interface{})`

SetCoreModules sets CoreModules field to given value.

### HasCoreModules

`func (o *UsersPutPreviousValues) HasCoreModules() bool`

HasCoreModules returns a boolean if a field has been set.

### SetCoreModulesNil

`func (o *UsersPutPreviousValues) SetCoreModulesNil(b bool)`

 SetCoreModulesNil sets the value for CoreModules to be an explicit nil

### UnsetCoreModules
`func (o *UsersPutPreviousValues) UnsetCoreModules()`

UnsetCoreModules ensures that no value is present for CoreModules, not even an explicit nil
### GetLocale

`func (o *UsersPutPreviousValues) GetLocale() string`

GetLocale returns the Locale field if non-nil, zero value otherwise.

### GetLocaleOk

`func (o *UsersPutPreviousValues) GetLocaleOk() (*string, bool)`

GetLocaleOk returns a tuple with the Locale field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocale

`func (o *UsersPutPreviousValues) SetLocale(v string)`

SetLocale sets Locale field to given value.

### HasLocale

`func (o *UsersPutPreviousValues) HasLocale() bool`

HasLocale returns a boolean if a field has been set.

### GetAccountExpirationDate

`func (o *UsersPutPreviousValues) GetAccountExpirationDate() string`

GetAccountExpirationDate returns the AccountExpirationDate field if non-nil, zero value otherwise.

### GetAccountExpirationDateOk

`func (o *UsersPutPreviousValues) GetAccountExpirationDateOk() (*string, bool)`

GetAccountExpirationDateOk returns a tuple with the AccountExpirationDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountExpirationDate

`func (o *UsersPutPreviousValues) SetAccountExpirationDate(v string)`

SetAccountExpirationDate sets AccountExpirationDate field to given value.

### HasAccountExpirationDate

`func (o *UsersPutPreviousValues) HasAccountExpirationDate() bool`

HasAccountExpirationDate returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


