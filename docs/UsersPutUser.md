# UsersPutUser

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int32** | Note: This is a Primary Key.&lt;pk/&gt; | [optional] [readonly] 
**CreatedAt** | Pointer to **string** |  | [optional] [readonly] 
**UpdatedAt** | Pointer to **string** |  | [optional] [readonly] 
**DeletedAt** | Pointer to **string** |  | [optional] [readonly] 
**FirstName** | **string** |  | 
**LastName** | **string** |  | 
**Email** | **string** |  | 
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
**MfaEnabled** | **bool** |  | [default to false]
**CellPhone** | **string** |  | [default to ""]
**AuthyId** | Pointer to **int32** |  | [optional] 
**EmployeeNumber** | Pointer to **string** |  | [optional] 
**Username** | Pointer to **string** |  | [optional] 
**LoginAttempt** | **int32** |  | [default to 0]
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
**CoreModules** | **interface{}** |  | 
**Locale** | Pointer to **string** |  | [optional] [default to "NULL"]
**AccountExpirationDate** | Pointer to **string** |  | [optional] 

## Methods

### NewUsersPutUser

`func NewUsersPutUser(firstName string, lastName string, email string, mfaEnabled bool, cellPhone string, loginAttempt int32, coreModules interface{}, ) *UsersPutUser`

NewUsersPutUser instantiates a new UsersPutUser object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUsersPutUserWithDefaults

`func NewUsersPutUserWithDefaults() *UsersPutUser`

NewUsersPutUserWithDefaults instantiates a new UsersPutUser object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *UsersPutUser) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *UsersPutUser) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *UsersPutUser) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *UsersPutUser) HasId() bool`

HasId returns a boolean if a field has been set.

### GetCreatedAt

`func (o *UsersPutUser) GetCreatedAt() string`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *UsersPutUser) GetCreatedAtOk() (*string, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *UsersPutUser) SetCreatedAt(v string)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *UsersPutUser) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *UsersPutUser) GetUpdatedAt() string`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *UsersPutUser) GetUpdatedAtOk() (*string, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *UsersPutUser) SetUpdatedAt(v string)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *UsersPutUser) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.

### GetDeletedAt

`func (o *UsersPutUser) GetDeletedAt() string`

GetDeletedAt returns the DeletedAt field if non-nil, zero value otherwise.

### GetDeletedAtOk

`func (o *UsersPutUser) GetDeletedAtOk() (*string, bool)`

GetDeletedAtOk returns a tuple with the DeletedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeletedAt

`func (o *UsersPutUser) SetDeletedAt(v string)`

SetDeletedAt sets DeletedAt field to given value.

### HasDeletedAt

`func (o *UsersPutUser) HasDeletedAt() bool`

HasDeletedAt returns a boolean if a field has been set.

### GetFirstName

`func (o *UsersPutUser) GetFirstName() string`

GetFirstName returns the FirstName field if non-nil, zero value otherwise.

### GetFirstNameOk

`func (o *UsersPutUser) GetFirstNameOk() (*string, bool)`

GetFirstNameOk returns a tuple with the FirstName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFirstName

`func (o *UsersPutUser) SetFirstName(v string)`

SetFirstName sets FirstName field to given value.


### GetLastName

`func (o *UsersPutUser) GetLastName() string`

GetLastName returns the LastName field if non-nil, zero value otherwise.

### GetLastNameOk

`func (o *UsersPutUser) GetLastNameOk() (*string, bool)`

GetLastNameOk returns a tuple with the LastName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastName

`func (o *UsersPutUser) SetLastName(v string)`

SetLastName sets LastName field to given value.


### GetEmail

`func (o *UsersPutUser) GetEmail() string`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *UsersPutUser) GetEmailOk() (*string, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *UsersPutUser) SetEmail(v string)`

SetEmail sets Email field to given value.


### GetPassword

`func (o *UsersPutUser) GetPassword() string`

GetPassword returns the Password field if non-nil, zero value otherwise.

### GetPasswordOk

`func (o *UsersPutUser) GetPasswordOk() (*string, bool)`

GetPasswordOk returns a tuple with the Password field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPassword

`func (o *UsersPutUser) SetPassword(v string)`

SetPassword sets Password field to given value.

### HasPassword

`func (o *UsersPutUser) HasPassword() bool`

HasPassword returns a boolean if a field has been set.

### GetSalt

`func (o *UsersPutUser) GetSalt() string`

GetSalt returns the Salt field if non-nil, zero value otherwise.

### GetSaltOk

`func (o *UsersPutUser) GetSaltOk() (*string, bool)`

GetSaltOk returns a tuple with the Salt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSalt

`func (o *UsersPutUser) SetSalt(v string)`

SetSalt sets Salt field to given value.

### HasSalt

`func (o *UsersPutUser) HasSalt() bool`

HasSalt returns a boolean if a field has been set.

### GetRoleId

`func (o *UsersPutUser) GetRoleId() int32`

GetRoleId returns the RoleId field if non-nil, zero value otherwise.

### GetRoleIdOk

`func (o *UsersPutUser) GetRoleIdOk() (*int32, bool)`

GetRoleIdOk returns a tuple with the RoleId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRoleId

`func (o *UsersPutUser) SetRoleId(v int32)`

SetRoleId sets RoleId field to given value.

### HasRoleId

`func (o *UsersPutUser) HasRoleId() bool`

HasRoleId returns a boolean if a field has been set.

### GetAcceptedTerms

`func (o *UsersPutUser) GetAcceptedTerms() bool`

GetAcceptedTerms returns the AcceptedTerms field if non-nil, zero value otherwise.

### GetAcceptedTermsOk

`func (o *UsersPutUser) GetAcceptedTermsOk() (*bool, bool)`

GetAcceptedTermsOk returns a tuple with the AcceptedTerms field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAcceptedTerms

`func (o *UsersPutUser) SetAcceptedTerms(v bool)`

SetAcceptedTerms sets AcceptedTerms field to given value.

### HasAcceptedTerms

`func (o *UsersPutUser) HasAcceptedTerms() bool`

HasAcceptedTerms returns a boolean if a field has been set.

### GetLastLogin

`func (o *UsersPutUser) GetLastLogin() string`

GetLastLogin returns the LastLogin field if non-nil, zero value otherwise.

### GetLastLoginOk

`func (o *UsersPutUser) GetLastLoginOk() (*string, bool)`

GetLastLoginOk returns a tuple with the LastLogin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastLogin

`func (o *UsersPutUser) SetLastLogin(v string)`

SetLastLogin sets LastLogin field to given value.

### HasLastLogin

`func (o *UsersPutUser) HasLastLogin() bool`

HasLastLogin returns a boolean if a field has been set.

### GetLastPasswordChange

`func (o *UsersPutUser) GetLastPasswordChange() string`

GetLastPasswordChange returns the LastPasswordChange field if non-nil, zero value otherwise.

### GetLastPasswordChangeOk

`func (o *UsersPutUser) GetLastPasswordChangeOk() (*string, bool)`

GetLastPasswordChangeOk returns a tuple with the LastPasswordChange field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastPasswordChange

`func (o *UsersPutUser) SetLastPasswordChange(v string)`

SetLastPasswordChange sets LastPasswordChange field to given value.

### HasLastPasswordChange

`func (o *UsersPutUser) HasLastPasswordChange() bool`

HasLastPasswordChange returns a boolean if a field has been set.

### GetLastIpAddress

`func (o *UsersPutUser) GetLastIpAddress() string`

GetLastIpAddress returns the LastIpAddress field if non-nil, zero value otherwise.

### GetLastIpAddressOk

`func (o *UsersPutUser) GetLastIpAddressOk() (*string, bool)`

GetLastIpAddressOk returns a tuple with the LastIpAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastIpAddress

`func (o *UsersPutUser) SetLastIpAddress(v string)`

SetLastIpAddress sets LastIpAddress field to given value.

### HasLastIpAddress

`func (o *UsersPutUser) HasLastIpAddress() bool`

HasLastIpAddress returns a boolean if a field has been set.

### GetPasswordHistory

`func (o *UsersPutUser) GetPasswordHistory() string`

GetPasswordHistory returns the PasswordHistory field if non-nil, zero value otherwise.

### GetPasswordHistoryOk

`func (o *UsersPutUser) GetPasswordHistoryOk() (*string, bool)`

GetPasswordHistoryOk returns a tuple with the PasswordHistory field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPasswordHistory

`func (o *UsersPutUser) SetPasswordHistory(v string)`

SetPasswordHistory sets PasswordHistory field to given value.

### HasPasswordHistory

`func (o *UsersPutUser) HasPasswordHistory() bool`

HasPasswordHistory returns a boolean if a field has been set.

### GetVerifiedIpHistory

`func (o *UsersPutUser) GetVerifiedIpHistory() string`

GetVerifiedIpHistory returns the VerifiedIpHistory field if non-nil, zero value otherwise.

### GetVerifiedIpHistoryOk

`func (o *UsersPutUser) GetVerifiedIpHistoryOk() (*string, bool)`

GetVerifiedIpHistoryOk returns a tuple with the VerifiedIpHistory field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVerifiedIpHistory

`func (o *UsersPutUser) SetVerifiedIpHistory(v string)`

SetVerifiedIpHistory sets VerifiedIpHistory field to given value.

### HasVerifiedIpHistory

`func (o *UsersPutUser) HasVerifiedIpHistory() bool`

HasVerifiedIpHistory returns a boolean if a field has been set.

### GetRequiresPasswordChange

`func (o *UsersPutUser) GetRequiresPasswordChange() bool`

GetRequiresPasswordChange returns the RequiresPasswordChange field if non-nil, zero value otherwise.

### GetRequiresPasswordChangeOk

`func (o *UsersPutUser) GetRequiresPasswordChangeOk() (*bool, bool)`

GetRequiresPasswordChangeOk returns a tuple with the RequiresPasswordChange field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequiresPasswordChange

`func (o *UsersPutUser) SetRequiresPasswordChange(v bool)`

SetRequiresPasswordChange sets RequiresPasswordChange field to given value.

### HasRequiresPasswordChange

`func (o *UsersPutUser) HasRequiresPasswordChange() bool`

HasRequiresPasswordChange returns a boolean if a field has been set.

### GetIconBgColor

`func (o *UsersPutUser) GetIconBgColor() string`

GetIconBgColor returns the IconBgColor field if non-nil, zero value otherwise.

### GetIconBgColorOk

`func (o *UsersPutUser) GetIconBgColorOk() (*string, bool)`

GetIconBgColorOk returns a tuple with the IconBgColor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIconBgColor

`func (o *UsersPutUser) SetIconBgColor(v string)`

SetIconBgColor sets IconBgColor field to given value.

### HasIconBgColor

`func (o *UsersPutUser) HasIconBgColor() bool`

HasIconBgColor returns a boolean if a field has been set.

### GetIconFgColor

`func (o *UsersPutUser) GetIconFgColor() string`

GetIconFgColor returns the IconFgColor field if non-nil, zero value otherwise.

### GetIconFgColorOk

`func (o *UsersPutUser) GetIconFgColorOk() (*string, bool)`

GetIconFgColorOk returns a tuple with the IconFgColor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIconFgColor

`func (o *UsersPutUser) SetIconFgColor(v string)`

SetIconFgColor sets IconFgColor field to given value.

### HasIconFgColor

`func (o *UsersPutUser) HasIconFgColor() bool`

HasIconFgColor returns a boolean if a field has been set.

### GetCompetencyEvaluation

`func (o *UsersPutUser) GetCompetencyEvaluation() string`

GetCompetencyEvaluation returns the CompetencyEvaluation field if non-nil, zero value otherwise.

### GetCompetencyEvaluationOk

`func (o *UsersPutUser) GetCompetencyEvaluationOk() (*string, bool)`

GetCompetencyEvaluationOk returns a tuple with the CompetencyEvaluation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompetencyEvaluation

`func (o *UsersPutUser) SetCompetencyEvaluation(v string)`

SetCompetencyEvaluation sets CompetencyEvaluation field to given value.

### HasCompetencyEvaluation

`func (o *UsersPutUser) HasCompetencyEvaluation() bool`

HasCompetencyEvaluation returns a boolean if a field has been set.

### GetUserStatId

`func (o *UsersPutUser) GetUserStatId() int32`

GetUserStatId returns the UserStatId field if non-nil, zero value otherwise.

### GetUserStatIdOk

`func (o *UsersPutUser) GetUserStatIdOk() (*int32, bool)`

GetUserStatIdOk returns a tuple with the UserStatId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserStatId

`func (o *UsersPutUser) SetUserStatId(v int32)`

SetUserStatId sets UserStatId field to given value.

### HasUserStatId

`func (o *UsersPutUser) HasUserStatId() bool`

HasUserStatId returns a boolean if a field has been set.

### GetPasswordResetToken

`func (o *UsersPutUser) GetPasswordResetToken() string`

GetPasswordResetToken returns the PasswordResetToken field if non-nil, zero value otherwise.

### GetPasswordResetTokenOk

`func (o *UsersPutUser) GetPasswordResetTokenOk() (*string, bool)`

GetPasswordResetTokenOk returns a tuple with the PasswordResetToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPasswordResetToken

`func (o *UsersPutUser) SetPasswordResetToken(v string)`

SetPasswordResetToken sets PasswordResetToken field to given value.

### HasPasswordResetToken

`func (o *UsersPutUser) HasPasswordResetToken() bool`

HasPasswordResetToken returns a boolean if a field has been set.

### GetPasswordResetTokenExpires

`func (o *UsersPutUser) GetPasswordResetTokenExpires() string`

GetPasswordResetTokenExpires returns the PasswordResetTokenExpires field if non-nil, zero value otherwise.

### GetPasswordResetTokenExpiresOk

`func (o *UsersPutUser) GetPasswordResetTokenExpiresOk() (*string, bool)`

GetPasswordResetTokenExpiresOk returns a tuple with the PasswordResetTokenExpires field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPasswordResetTokenExpires

`func (o *UsersPutUser) SetPasswordResetTokenExpires(v string)`

SetPasswordResetTokenExpires sets PasswordResetTokenExpires field to given value.

### HasPasswordResetTokenExpires

`func (o *UsersPutUser) HasPasswordResetTokenExpires() bool`

HasPasswordResetTokenExpires returns a boolean if a field has been set.

### GetAuthType

`func (o *UsersPutUser) GetAuthType() string`

GetAuthType returns the AuthType field if non-nil, zero value otherwise.

### GetAuthTypeOk

`func (o *UsersPutUser) GetAuthTypeOk() (*string, bool)`

GetAuthTypeOk returns a tuple with the AuthType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthType

`func (o *UsersPutUser) SetAuthType(v string)`

SetAuthType sets AuthType field to given value.

### HasAuthType

`func (o *UsersPutUser) HasAuthType() bool`

HasAuthType returns a boolean if a field has been set.

### GetLocationId

`func (o *UsersPutUser) GetLocationId() int32`

GetLocationId returns the LocationId field if non-nil, zero value otherwise.

### GetLocationIdOk

`func (o *UsersPutUser) GetLocationIdOk() (*int32, bool)`

GetLocationIdOk returns a tuple with the LocationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocationId

`func (o *UsersPutUser) SetLocationId(v int32)`

SetLocationId sets LocationId field to given value.

### HasLocationId

`func (o *UsersPutUser) HasLocationId() bool`

HasLocationId returns a boolean if a field has been set.

### GetDepartmentId

`func (o *UsersPutUser) GetDepartmentId() int32`

GetDepartmentId returns the DepartmentId field if non-nil, zero value otherwise.

### GetDepartmentIdOk

`func (o *UsersPutUser) GetDepartmentIdOk() (*int32, bool)`

GetDepartmentIdOk returns a tuple with the DepartmentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDepartmentId

`func (o *UsersPutUser) SetDepartmentId(v int32)`

SetDepartmentId sets DepartmentId field to given value.

### HasDepartmentId

`func (o *UsersPutUser) HasDepartmentId() bool`

HasDepartmentId returns a boolean if a field has been set.

### GetJobTitle

`func (o *UsersPutUser) GetJobTitle() string`

GetJobTitle returns the JobTitle field if non-nil, zero value otherwise.

### GetJobTitleOk

`func (o *UsersPutUser) GetJobTitleOk() (*string, bool)`

GetJobTitleOk returns a tuple with the JobTitle field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJobTitle

`func (o *UsersPutUser) SetJobTitle(v string)`

SetJobTitle sets JobTitle field to given value.

### HasJobTitle

`func (o *UsersPutUser) HasJobTitle() bool`

HasJobTitle returns a boolean if a field has been set.

### GetMfaEnabled

`func (o *UsersPutUser) GetMfaEnabled() bool`

GetMfaEnabled returns the MfaEnabled field if non-nil, zero value otherwise.

### GetMfaEnabledOk

`func (o *UsersPutUser) GetMfaEnabledOk() (*bool, bool)`

GetMfaEnabledOk returns a tuple with the MfaEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMfaEnabled

`func (o *UsersPutUser) SetMfaEnabled(v bool)`

SetMfaEnabled sets MfaEnabled field to given value.


### GetCellPhone

`func (o *UsersPutUser) GetCellPhone() string`

GetCellPhone returns the CellPhone field if non-nil, zero value otherwise.

### GetCellPhoneOk

`func (o *UsersPutUser) GetCellPhoneOk() (*string, bool)`

GetCellPhoneOk returns a tuple with the CellPhone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCellPhone

`func (o *UsersPutUser) SetCellPhone(v string)`

SetCellPhone sets CellPhone field to given value.


### GetAuthyId

`func (o *UsersPutUser) GetAuthyId() int32`

GetAuthyId returns the AuthyId field if non-nil, zero value otherwise.

### GetAuthyIdOk

`func (o *UsersPutUser) GetAuthyIdOk() (*int32, bool)`

GetAuthyIdOk returns a tuple with the AuthyId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthyId

`func (o *UsersPutUser) SetAuthyId(v int32)`

SetAuthyId sets AuthyId field to given value.

### HasAuthyId

`func (o *UsersPutUser) HasAuthyId() bool`

HasAuthyId returns a boolean if a field has been set.

### GetEmployeeNumber

`func (o *UsersPutUser) GetEmployeeNumber() string`

GetEmployeeNumber returns the EmployeeNumber field if non-nil, zero value otherwise.

### GetEmployeeNumberOk

`func (o *UsersPutUser) GetEmployeeNumberOk() (*string, bool)`

GetEmployeeNumberOk returns a tuple with the EmployeeNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmployeeNumber

`func (o *UsersPutUser) SetEmployeeNumber(v string)`

SetEmployeeNumber sets EmployeeNumber field to given value.

### HasEmployeeNumber

`func (o *UsersPutUser) HasEmployeeNumber() bool`

HasEmployeeNumber returns a boolean if a field has been set.

### GetUsername

`func (o *UsersPutUser) GetUsername() string`

GetUsername returns the Username field if non-nil, zero value otherwise.

### GetUsernameOk

`func (o *UsersPutUser) GetUsernameOk() (*string, bool)`

GetUsernameOk returns a tuple with the Username field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsername

`func (o *UsersPutUser) SetUsername(v string)`

SetUsername sets Username field to given value.

### HasUsername

`func (o *UsersPutUser) HasUsername() bool`

HasUsername returns a boolean if a field has been set.

### GetLoginAttempt

`func (o *UsersPutUser) GetLoginAttempt() int32`

GetLoginAttempt returns the LoginAttempt field if non-nil, zero value otherwise.

### GetLoginAttemptOk

`func (o *UsersPutUser) GetLoginAttemptOk() (*int32, bool)`

GetLoginAttemptOk returns a tuple with the LoginAttempt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLoginAttempt

`func (o *UsersPutUser) SetLoginAttempt(v int32)`

SetLoginAttempt sets LoginAttempt field to given value.


### GetLastLoginFailedAttempt

`func (o *UsersPutUser) GetLastLoginFailedAttempt() string`

GetLastLoginFailedAttempt returns the LastLoginFailedAttempt field if non-nil, zero value otherwise.

### GetLastLoginFailedAttemptOk

`func (o *UsersPutUser) GetLastLoginFailedAttemptOk() (*string, bool)`

GetLastLoginFailedAttemptOk returns a tuple with the LastLoginFailedAttempt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastLoginFailedAttempt

`func (o *UsersPutUser) SetLastLoginFailedAttempt(v string)`

SetLastLoginFailedAttempt sets LastLoginFailedAttempt field to given value.

### HasLastLoginFailedAttempt

`func (o *UsersPutUser) HasLastLoginFailedAttempt() bool`

HasLastLoginFailedAttempt returns a boolean if a field has been set.

### GetFirstLogin

`func (o *UsersPutUser) GetFirstLogin() string`

GetFirstLogin returns the FirstLogin field if non-nil, zero value otherwise.

### GetFirstLoginOk

`func (o *UsersPutUser) GetFirstLoginOk() (*string, bool)`

GetFirstLoginOk returns a tuple with the FirstLogin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFirstLogin

`func (o *UsersPutUser) SetFirstLogin(v string)`

SetFirstLogin sets FirstLogin field to given value.

### HasFirstLogin

`func (o *UsersPutUser) HasFirstLogin() bool`

HasFirstLogin returns a boolean if a field has been set.

### GetChameleonIdentifier

`func (o *UsersPutUser) GetChameleonIdentifier() string`

GetChameleonIdentifier returns the ChameleonIdentifier field if non-nil, zero value otherwise.

### GetChameleonIdentifierOk

`func (o *UsersPutUser) GetChameleonIdentifierOk() (*string, bool)`

GetChameleonIdentifierOk returns a tuple with the ChameleonIdentifier field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChameleonIdentifier

`func (o *UsersPutUser) SetChameleonIdentifier(v string)`

SetChameleonIdentifier sets ChameleonIdentifier field to given value.

### HasChameleonIdentifier

`func (o *UsersPutUser) HasChameleonIdentifier() bool`

HasChameleonIdentifier returns a boolean if a field has been set.

### GetSlackUserId

`func (o *UsersPutUser) GetSlackUserId() string`

GetSlackUserId returns the SlackUserId field if non-nil, zero value otherwise.

### GetSlackUserIdOk

`func (o *UsersPutUser) GetSlackUserIdOk() (*string, bool)`

GetSlackUserIdOk returns a tuple with the SlackUserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSlackUserId

`func (o *UsersPutUser) SetSlackUserId(v string)`

SetSlackUserId sets SlackUserId field to given value.

### HasSlackUserId

`func (o *UsersPutUser) HasSlackUserId() bool`

HasSlackUserId returns a boolean if a field has been set.

### GetSentSlackWelcomeNotification

`func (o *UsersPutUser) GetSentSlackWelcomeNotification() bool`

GetSentSlackWelcomeNotification returns the SentSlackWelcomeNotification field if non-nil, zero value otherwise.

### GetSentSlackWelcomeNotificationOk

`func (o *UsersPutUser) GetSentSlackWelcomeNotificationOk() (*bool, bool)`

GetSentSlackWelcomeNotificationOk returns a tuple with the SentSlackWelcomeNotification field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSentSlackWelcomeNotification

`func (o *UsersPutUser) SetSentSlackWelcomeNotification(v bool)`

SetSentSlackWelcomeNotification sets SentSlackWelcomeNotification field to given value.

### HasSentSlackWelcomeNotification

`func (o *UsersPutUser) HasSentSlackWelcomeNotification() bool`

HasSentSlackWelcomeNotification returns a boolean if a field has been set.

### GetAvatarKey

`func (o *UsersPutUser) GetAvatarKey() string`

GetAvatarKey returns the AvatarKey field if non-nil, zero value otherwise.

### GetAvatarKeyOk

`func (o *UsersPutUser) GetAvatarKeyOk() (*string, bool)`

GetAvatarKeyOk returns a tuple with the AvatarKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvatarKey

`func (o *UsersPutUser) SetAvatarKey(v string)`

SetAvatarKey sets AvatarKey field to given value.

### HasAvatarKey

`func (o *UsersPutUser) HasAvatarKey() bool`

HasAvatarKey returns a boolean if a field has been set.

### GetExternalApiToken

`func (o *UsersPutUser) GetExternalApiToken() string`

GetExternalApiToken returns the ExternalApiToken field if non-nil, zero value otherwise.

### GetExternalApiTokenOk

`func (o *UsersPutUser) GetExternalApiTokenOk() (*string, bool)`

GetExternalApiTokenOk returns a tuple with the ExternalApiToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalApiToken

`func (o *UsersPutUser) SetExternalApiToken(v string)`

SetExternalApiToken sets ExternalApiToken field to given value.

### HasExternalApiToken

`func (o *UsersPutUser) HasExternalApiToken() bool`

HasExternalApiToken returns a boolean if a field has been set.

### GetAzureConversationData

`func (o *UsersPutUser) GetAzureConversationData() interface{}`

GetAzureConversationData returns the AzureConversationData field if non-nil, zero value otherwise.

### GetAzureConversationDataOk

`func (o *UsersPutUser) GetAzureConversationDataOk() (*interface{}, bool)`

GetAzureConversationDataOk returns a tuple with the AzureConversationData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAzureConversationData

`func (o *UsersPutUser) SetAzureConversationData(v interface{})`

SetAzureConversationData sets AzureConversationData field to given value.

### HasAzureConversationData

`func (o *UsersPutUser) HasAzureConversationData() bool`

HasAzureConversationData returns a boolean if a field has been set.

### SetAzureConversationDataNil

`func (o *UsersPutUser) SetAzureConversationDataNil(b bool)`

 SetAzureConversationDataNil sets the value for AzureConversationData to be an explicit nil

### UnsetAzureConversationData
`func (o *UsersPutUser) UnsetAzureConversationData()`

UnsetAzureConversationData ensures that no value is present for AzureConversationData, not even an explicit nil
### GetScimUserActive

`func (o *UsersPutUser) GetScimUserActive() string`

GetScimUserActive returns the ScimUserActive field if non-nil, zero value otherwise.

### GetScimUserActiveOk

`func (o *UsersPutUser) GetScimUserActiveOk() (*string, bool)`

GetScimUserActiveOk returns a tuple with the ScimUserActive field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScimUserActive

`func (o *UsersPutUser) SetScimUserActive(v string)`

SetScimUserActive sets ScimUserActive field to given value.

### HasScimUserActive

`func (o *UsersPutUser) HasScimUserActive() bool`

HasScimUserActive returns a boolean if a field has been set.

### GetTeamsNotificationsEnabled

`func (o *UsersPutUser) GetTeamsNotificationsEnabled() bool`

GetTeamsNotificationsEnabled returns the TeamsNotificationsEnabled field if non-nil, zero value otherwise.

### GetTeamsNotificationsEnabledOk

`func (o *UsersPutUser) GetTeamsNotificationsEnabledOk() (*bool, bool)`

GetTeamsNotificationsEnabledOk returns a tuple with the TeamsNotificationsEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTeamsNotificationsEnabled

`func (o *UsersPutUser) SetTeamsNotificationsEnabled(v bool)`

SetTeamsNotificationsEnabled sets TeamsNotificationsEnabled field to given value.

### HasTeamsNotificationsEnabled

`func (o *UsersPutUser) HasTeamsNotificationsEnabled() bool`

HasTeamsNotificationsEnabled returns a boolean if a field has been set.

### GetCoreModules

`func (o *UsersPutUser) GetCoreModules() interface{}`

GetCoreModules returns the CoreModules field if non-nil, zero value otherwise.

### GetCoreModulesOk

`func (o *UsersPutUser) GetCoreModulesOk() (*interface{}, bool)`

GetCoreModulesOk returns a tuple with the CoreModules field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCoreModules

`func (o *UsersPutUser) SetCoreModules(v interface{})`

SetCoreModules sets CoreModules field to given value.


### SetCoreModulesNil

`func (o *UsersPutUser) SetCoreModulesNil(b bool)`

 SetCoreModulesNil sets the value for CoreModules to be an explicit nil

### UnsetCoreModules
`func (o *UsersPutUser) UnsetCoreModules()`

UnsetCoreModules ensures that no value is present for CoreModules, not even an explicit nil
### GetLocale

`func (o *UsersPutUser) GetLocale() string`

GetLocale returns the Locale field if non-nil, zero value otherwise.

### GetLocaleOk

`func (o *UsersPutUser) GetLocaleOk() (*string, bool)`

GetLocaleOk returns a tuple with the Locale field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocale

`func (o *UsersPutUser) SetLocale(v string)`

SetLocale sets Locale field to given value.

### HasLocale

`func (o *UsersPutUser) HasLocale() bool`

HasLocale returns a boolean if a field has been set.

### GetAccountExpirationDate

`func (o *UsersPutUser) GetAccountExpirationDate() string`

GetAccountExpirationDate returns the AccountExpirationDate field if non-nil, zero value otherwise.

### GetAccountExpirationDateOk

`func (o *UsersPutUser) GetAccountExpirationDateOk() (*string, bool)`

GetAccountExpirationDateOk returns a tuple with the AccountExpirationDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountExpirationDate

`func (o *UsersPutUser) SetAccountExpirationDate(v string)`

SetAccountExpirationDate sets AccountExpirationDate field to given value.

### HasAccountExpirationDate

`func (o *UsersPutUser) HasAccountExpirationDate() bool`

HasAccountExpirationDate returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


