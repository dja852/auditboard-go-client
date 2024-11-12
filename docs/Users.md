# Users

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

### NewUsers

`func NewUsers(firstName string, lastName string, email string, mfaEnabled bool, cellPhone string, loginAttempt int32, coreModules interface{}, ) *Users`

NewUsers instantiates a new Users object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUsersWithDefaults

`func NewUsersWithDefaults() *Users`

NewUsersWithDefaults instantiates a new Users object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *Users) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Users) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Users) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *Users) HasId() bool`

HasId returns a boolean if a field has been set.

### GetCreatedAt

`func (o *Users) GetCreatedAt() string`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *Users) GetCreatedAtOk() (*string, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *Users) SetCreatedAt(v string)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *Users) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *Users) GetUpdatedAt() string`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *Users) GetUpdatedAtOk() (*string, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *Users) SetUpdatedAt(v string)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *Users) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.

### GetDeletedAt

`func (o *Users) GetDeletedAt() string`

GetDeletedAt returns the DeletedAt field if non-nil, zero value otherwise.

### GetDeletedAtOk

`func (o *Users) GetDeletedAtOk() (*string, bool)`

GetDeletedAtOk returns a tuple with the DeletedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeletedAt

`func (o *Users) SetDeletedAt(v string)`

SetDeletedAt sets DeletedAt field to given value.

### HasDeletedAt

`func (o *Users) HasDeletedAt() bool`

HasDeletedAt returns a boolean if a field has been set.

### GetFirstName

`func (o *Users) GetFirstName() string`

GetFirstName returns the FirstName field if non-nil, zero value otherwise.

### GetFirstNameOk

`func (o *Users) GetFirstNameOk() (*string, bool)`

GetFirstNameOk returns a tuple with the FirstName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFirstName

`func (o *Users) SetFirstName(v string)`

SetFirstName sets FirstName field to given value.


### GetLastName

`func (o *Users) GetLastName() string`

GetLastName returns the LastName field if non-nil, zero value otherwise.

### GetLastNameOk

`func (o *Users) GetLastNameOk() (*string, bool)`

GetLastNameOk returns a tuple with the LastName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastName

`func (o *Users) SetLastName(v string)`

SetLastName sets LastName field to given value.


### GetEmail

`func (o *Users) GetEmail() string`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *Users) GetEmailOk() (*string, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *Users) SetEmail(v string)`

SetEmail sets Email field to given value.


### GetPassword

`func (o *Users) GetPassword() string`

GetPassword returns the Password field if non-nil, zero value otherwise.

### GetPasswordOk

`func (o *Users) GetPasswordOk() (*string, bool)`

GetPasswordOk returns a tuple with the Password field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPassword

`func (o *Users) SetPassword(v string)`

SetPassword sets Password field to given value.

### HasPassword

`func (o *Users) HasPassword() bool`

HasPassword returns a boolean if a field has been set.

### GetSalt

`func (o *Users) GetSalt() string`

GetSalt returns the Salt field if non-nil, zero value otherwise.

### GetSaltOk

`func (o *Users) GetSaltOk() (*string, bool)`

GetSaltOk returns a tuple with the Salt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSalt

`func (o *Users) SetSalt(v string)`

SetSalt sets Salt field to given value.

### HasSalt

`func (o *Users) HasSalt() bool`

HasSalt returns a boolean if a field has been set.

### GetRoleId

`func (o *Users) GetRoleId() int32`

GetRoleId returns the RoleId field if non-nil, zero value otherwise.

### GetRoleIdOk

`func (o *Users) GetRoleIdOk() (*int32, bool)`

GetRoleIdOk returns a tuple with the RoleId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRoleId

`func (o *Users) SetRoleId(v int32)`

SetRoleId sets RoleId field to given value.

### HasRoleId

`func (o *Users) HasRoleId() bool`

HasRoleId returns a boolean if a field has been set.

### GetAcceptedTerms

`func (o *Users) GetAcceptedTerms() bool`

GetAcceptedTerms returns the AcceptedTerms field if non-nil, zero value otherwise.

### GetAcceptedTermsOk

`func (o *Users) GetAcceptedTermsOk() (*bool, bool)`

GetAcceptedTermsOk returns a tuple with the AcceptedTerms field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAcceptedTerms

`func (o *Users) SetAcceptedTerms(v bool)`

SetAcceptedTerms sets AcceptedTerms field to given value.

### HasAcceptedTerms

`func (o *Users) HasAcceptedTerms() bool`

HasAcceptedTerms returns a boolean if a field has been set.

### GetLastLogin

`func (o *Users) GetLastLogin() string`

GetLastLogin returns the LastLogin field if non-nil, zero value otherwise.

### GetLastLoginOk

`func (o *Users) GetLastLoginOk() (*string, bool)`

GetLastLoginOk returns a tuple with the LastLogin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastLogin

`func (o *Users) SetLastLogin(v string)`

SetLastLogin sets LastLogin field to given value.

### HasLastLogin

`func (o *Users) HasLastLogin() bool`

HasLastLogin returns a boolean if a field has been set.

### GetLastPasswordChange

`func (o *Users) GetLastPasswordChange() string`

GetLastPasswordChange returns the LastPasswordChange field if non-nil, zero value otherwise.

### GetLastPasswordChangeOk

`func (o *Users) GetLastPasswordChangeOk() (*string, bool)`

GetLastPasswordChangeOk returns a tuple with the LastPasswordChange field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastPasswordChange

`func (o *Users) SetLastPasswordChange(v string)`

SetLastPasswordChange sets LastPasswordChange field to given value.

### HasLastPasswordChange

`func (o *Users) HasLastPasswordChange() bool`

HasLastPasswordChange returns a boolean if a field has been set.

### GetLastIpAddress

`func (o *Users) GetLastIpAddress() string`

GetLastIpAddress returns the LastIpAddress field if non-nil, zero value otherwise.

### GetLastIpAddressOk

`func (o *Users) GetLastIpAddressOk() (*string, bool)`

GetLastIpAddressOk returns a tuple with the LastIpAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastIpAddress

`func (o *Users) SetLastIpAddress(v string)`

SetLastIpAddress sets LastIpAddress field to given value.

### HasLastIpAddress

`func (o *Users) HasLastIpAddress() bool`

HasLastIpAddress returns a boolean if a field has been set.

### GetPasswordHistory

`func (o *Users) GetPasswordHistory() string`

GetPasswordHistory returns the PasswordHistory field if non-nil, zero value otherwise.

### GetPasswordHistoryOk

`func (o *Users) GetPasswordHistoryOk() (*string, bool)`

GetPasswordHistoryOk returns a tuple with the PasswordHistory field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPasswordHistory

`func (o *Users) SetPasswordHistory(v string)`

SetPasswordHistory sets PasswordHistory field to given value.

### HasPasswordHistory

`func (o *Users) HasPasswordHistory() bool`

HasPasswordHistory returns a boolean if a field has been set.

### GetVerifiedIpHistory

`func (o *Users) GetVerifiedIpHistory() string`

GetVerifiedIpHistory returns the VerifiedIpHistory field if non-nil, zero value otherwise.

### GetVerifiedIpHistoryOk

`func (o *Users) GetVerifiedIpHistoryOk() (*string, bool)`

GetVerifiedIpHistoryOk returns a tuple with the VerifiedIpHistory field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVerifiedIpHistory

`func (o *Users) SetVerifiedIpHistory(v string)`

SetVerifiedIpHistory sets VerifiedIpHistory field to given value.

### HasVerifiedIpHistory

`func (o *Users) HasVerifiedIpHistory() bool`

HasVerifiedIpHistory returns a boolean if a field has been set.

### GetRequiresPasswordChange

`func (o *Users) GetRequiresPasswordChange() bool`

GetRequiresPasswordChange returns the RequiresPasswordChange field if non-nil, zero value otherwise.

### GetRequiresPasswordChangeOk

`func (o *Users) GetRequiresPasswordChangeOk() (*bool, bool)`

GetRequiresPasswordChangeOk returns a tuple with the RequiresPasswordChange field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequiresPasswordChange

`func (o *Users) SetRequiresPasswordChange(v bool)`

SetRequiresPasswordChange sets RequiresPasswordChange field to given value.

### HasRequiresPasswordChange

`func (o *Users) HasRequiresPasswordChange() bool`

HasRequiresPasswordChange returns a boolean if a field has been set.

### GetIconBgColor

`func (o *Users) GetIconBgColor() string`

GetIconBgColor returns the IconBgColor field if non-nil, zero value otherwise.

### GetIconBgColorOk

`func (o *Users) GetIconBgColorOk() (*string, bool)`

GetIconBgColorOk returns a tuple with the IconBgColor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIconBgColor

`func (o *Users) SetIconBgColor(v string)`

SetIconBgColor sets IconBgColor field to given value.

### HasIconBgColor

`func (o *Users) HasIconBgColor() bool`

HasIconBgColor returns a boolean if a field has been set.

### GetIconFgColor

`func (o *Users) GetIconFgColor() string`

GetIconFgColor returns the IconFgColor field if non-nil, zero value otherwise.

### GetIconFgColorOk

`func (o *Users) GetIconFgColorOk() (*string, bool)`

GetIconFgColorOk returns a tuple with the IconFgColor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIconFgColor

`func (o *Users) SetIconFgColor(v string)`

SetIconFgColor sets IconFgColor field to given value.

### HasIconFgColor

`func (o *Users) HasIconFgColor() bool`

HasIconFgColor returns a boolean if a field has been set.

### GetCompetencyEvaluation

`func (o *Users) GetCompetencyEvaluation() string`

GetCompetencyEvaluation returns the CompetencyEvaluation field if non-nil, zero value otherwise.

### GetCompetencyEvaluationOk

`func (o *Users) GetCompetencyEvaluationOk() (*string, bool)`

GetCompetencyEvaluationOk returns a tuple with the CompetencyEvaluation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompetencyEvaluation

`func (o *Users) SetCompetencyEvaluation(v string)`

SetCompetencyEvaluation sets CompetencyEvaluation field to given value.

### HasCompetencyEvaluation

`func (o *Users) HasCompetencyEvaluation() bool`

HasCompetencyEvaluation returns a boolean if a field has been set.

### GetUserStatId

`func (o *Users) GetUserStatId() int32`

GetUserStatId returns the UserStatId field if non-nil, zero value otherwise.

### GetUserStatIdOk

`func (o *Users) GetUserStatIdOk() (*int32, bool)`

GetUserStatIdOk returns a tuple with the UserStatId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserStatId

`func (o *Users) SetUserStatId(v int32)`

SetUserStatId sets UserStatId field to given value.

### HasUserStatId

`func (o *Users) HasUserStatId() bool`

HasUserStatId returns a boolean if a field has been set.

### GetPasswordResetToken

`func (o *Users) GetPasswordResetToken() string`

GetPasswordResetToken returns the PasswordResetToken field if non-nil, zero value otherwise.

### GetPasswordResetTokenOk

`func (o *Users) GetPasswordResetTokenOk() (*string, bool)`

GetPasswordResetTokenOk returns a tuple with the PasswordResetToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPasswordResetToken

`func (o *Users) SetPasswordResetToken(v string)`

SetPasswordResetToken sets PasswordResetToken field to given value.

### HasPasswordResetToken

`func (o *Users) HasPasswordResetToken() bool`

HasPasswordResetToken returns a boolean if a field has been set.

### GetPasswordResetTokenExpires

`func (o *Users) GetPasswordResetTokenExpires() string`

GetPasswordResetTokenExpires returns the PasswordResetTokenExpires field if non-nil, zero value otherwise.

### GetPasswordResetTokenExpiresOk

`func (o *Users) GetPasswordResetTokenExpiresOk() (*string, bool)`

GetPasswordResetTokenExpiresOk returns a tuple with the PasswordResetTokenExpires field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPasswordResetTokenExpires

`func (o *Users) SetPasswordResetTokenExpires(v string)`

SetPasswordResetTokenExpires sets PasswordResetTokenExpires field to given value.

### HasPasswordResetTokenExpires

`func (o *Users) HasPasswordResetTokenExpires() bool`

HasPasswordResetTokenExpires returns a boolean if a field has been set.

### GetAuthType

`func (o *Users) GetAuthType() string`

GetAuthType returns the AuthType field if non-nil, zero value otherwise.

### GetAuthTypeOk

`func (o *Users) GetAuthTypeOk() (*string, bool)`

GetAuthTypeOk returns a tuple with the AuthType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthType

`func (o *Users) SetAuthType(v string)`

SetAuthType sets AuthType field to given value.

### HasAuthType

`func (o *Users) HasAuthType() bool`

HasAuthType returns a boolean if a field has been set.

### GetLocationId

`func (o *Users) GetLocationId() int32`

GetLocationId returns the LocationId field if non-nil, zero value otherwise.

### GetLocationIdOk

`func (o *Users) GetLocationIdOk() (*int32, bool)`

GetLocationIdOk returns a tuple with the LocationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocationId

`func (o *Users) SetLocationId(v int32)`

SetLocationId sets LocationId field to given value.

### HasLocationId

`func (o *Users) HasLocationId() bool`

HasLocationId returns a boolean if a field has been set.

### GetDepartmentId

`func (o *Users) GetDepartmentId() int32`

GetDepartmentId returns the DepartmentId field if non-nil, zero value otherwise.

### GetDepartmentIdOk

`func (o *Users) GetDepartmentIdOk() (*int32, bool)`

GetDepartmentIdOk returns a tuple with the DepartmentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDepartmentId

`func (o *Users) SetDepartmentId(v int32)`

SetDepartmentId sets DepartmentId field to given value.

### HasDepartmentId

`func (o *Users) HasDepartmentId() bool`

HasDepartmentId returns a boolean if a field has been set.

### GetJobTitle

`func (o *Users) GetJobTitle() string`

GetJobTitle returns the JobTitle field if non-nil, zero value otherwise.

### GetJobTitleOk

`func (o *Users) GetJobTitleOk() (*string, bool)`

GetJobTitleOk returns a tuple with the JobTitle field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJobTitle

`func (o *Users) SetJobTitle(v string)`

SetJobTitle sets JobTitle field to given value.

### HasJobTitle

`func (o *Users) HasJobTitle() bool`

HasJobTitle returns a boolean if a field has been set.

### GetMfaEnabled

`func (o *Users) GetMfaEnabled() bool`

GetMfaEnabled returns the MfaEnabled field if non-nil, zero value otherwise.

### GetMfaEnabledOk

`func (o *Users) GetMfaEnabledOk() (*bool, bool)`

GetMfaEnabledOk returns a tuple with the MfaEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMfaEnabled

`func (o *Users) SetMfaEnabled(v bool)`

SetMfaEnabled sets MfaEnabled field to given value.


### GetCellPhone

`func (o *Users) GetCellPhone() string`

GetCellPhone returns the CellPhone field if non-nil, zero value otherwise.

### GetCellPhoneOk

`func (o *Users) GetCellPhoneOk() (*string, bool)`

GetCellPhoneOk returns a tuple with the CellPhone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCellPhone

`func (o *Users) SetCellPhone(v string)`

SetCellPhone sets CellPhone field to given value.


### GetAuthyId

`func (o *Users) GetAuthyId() int32`

GetAuthyId returns the AuthyId field if non-nil, zero value otherwise.

### GetAuthyIdOk

`func (o *Users) GetAuthyIdOk() (*int32, bool)`

GetAuthyIdOk returns a tuple with the AuthyId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthyId

`func (o *Users) SetAuthyId(v int32)`

SetAuthyId sets AuthyId field to given value.

### HasAuthyId

`func (o *Users) HasAuthyId() bool`

HasAuthyId returns a boolean if a field has been set.

### GetEmployeeNumber

`func (o *Users) GetEmployeeNumber() string`

GetEmployeeNumber returns the EmployeeNumber field if non-nil, zero value otherwise.

### GetEmployeeNumberOk

`func (o *Users) GetEmployeeNumberOk() (*string, bool)`

GetEmployeeNumberOk returns a tuple with the EmployeeNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmployeeNumber

`func (o *Users) SetEmployeeNumber(v string)`

SetEmployeeNumber sets EmployeeNumber field to given value.

### HasEmployeeNumber

`func (o *Users) HasEmployeeNumber() bool`

HasEmployeeNumber returns a boolean if a field has been set.

### GetUsername

`func (o *Users) GetUsername() string`

GetUsername returns the Username field if non-nil, zero value otherwise.

### GetUsernameOk

`func (o *Users) GetUsernameOk() (*string, bool)`

GetUsernameOk returns a tuple with the Username field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsername

`func (o *Users) SetUsername(v string)`

SetUsername sets Username field to given value.

### HasUsername

`func (o *Users) HasUsername() bool`

HasUsername returns a boolean if a field has been set.

### GetLoginAttempt

`func (o *Users) GetLoginAttempt() int32`

GetLoginAttempt returns the LoginAttempt field if non-nil, zero value otherwise.

### GetLoginAttemptOk

`func (o *Users) GetLoginAttemptOk() (*int32, bool)`

GetLoginAttemptOk returns a tuple with the LoginAttempt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLoginAttempt

`func (o *Users) SetLoginAttempt(v int32)`

SetLoginAttempt sets LoginAttempt field to given value.


### GetLastLoginFailedAttempt

`func (o *Users) GetLastLoginFailedAttempt() string`

GetLastLoginFailedAttempt returns the LastLoginFailedAttempt field if non-nil, zero value otherwise.

### GetLastLoginFailedAttemptOk

`func (o *Users) GetLastLoginFailedAttemptOk() (*string, bool)`

GetLastLoginFailedAttemptOk returns a tuple with the LastLoginFailedAttempt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastLoginFailedAttempt

`func (o *Users) SetLastLoginFailedAttempt(v string)`

SetLastLoginFailedAttempt sets LastLoginFailedAttempt field to given value.

### HasLastLoginFailedAttempt

`func (o *Users) HasLastLoginFailedAttempt() bool`

HasLastLoginFailedAttempt returns a boolean if a field has been set.

### GetFirstLogin

`func (o *Users) GetFirstLogin() string`

GetFirstLogin returns the FirstLogin field if non-nil, zero value otherwise.

### GetFirstLoginOk

`func (o *Users) GetFirstLoginOk() (*string, bool)`

GetFirstLoginOk returns a tuple with the FirstLogin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFirstLogin

`func (o *Users) SetFirstLogin(v string)`

SetFirstLogin sets FirstLogin field to given value.

### HasFirstLogin

`func (o *Users) HasFirstLogin() bool`

HasFirstLogin returns a boolean if a field has been set.

### GetChameleonIdentifier

`func (o *Users) GetChameleonIdentifier() string`

GetChameleonIdentifier returns the ChameleonIdentifier field if non-nil, zero value otherwise.

### GetChameleonIdentifierOk

`func (o *Users) GetChameleonIdentifierOk() (*string, bool)`

GetChameleonIdentifierOk returns a tuple with the ChameleonIdentifier field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChameleonIdentifier

`func (o *Users) SetChameleonIdentifier(v string)`

SetChameleonIdentifier sets ChameleonIdentifier field to given value.

### HasChameleonIdentifier

`func (o *Users) HasChameleonIdentifier() bool`

HasChameleonIdentifier returns a boolean if a field has been set.

### GetSlackUserId

`func (o *Users) GetSlackUserId() string`

GetSlackUserId returns the SlackUserId field if non-nil, zero value otherwise.

### GetSlackUserIdOk

`func (o *Users) GetSlackUserIdOk() (*string, bool)`

GetSlackUserIdOk returns a tuple with the SlackUserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSlackUserId

`func (o *Users) SetSlackUserId(v string)`

SetSlackUserId sets SlackUserId field to given value.

### HasSlackUserId

`func (o *Users) HasSlackUserId() bool`

HasSlackUserId returns a boolean if a field has been set.

### GetSentSlackWelcomeNotification

`func (o *Users) GetSentSlackWelcomeNotification() bool`

GetSentSlackWelcomeNotification returns the SentSlackWelcomeNotification field if non-nil, zero value otherwise.

### GetSentSlackWelcomeNotificationOk

`func (o *Users) GetSentSlackWelcomeNotificationOk() (*bool, bool)`

GetSentSlackWelcomeNotificationOk returns a tuple with the SentSlackWelcomeNotification field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSentSlackWelcomeNotification

`func (o *Users) SetSentSlackWelcomeNotification(v bool)`

SetSentSlackWelcomeNotification sets SentSlackWelcomeNotification field to given value.

### HasSentSlackWelcomeNotification

`func (o *Users) HasSentSlackWelcomeNotification() bool`

HasSentSlackWelcomeNotification returns a boolean if a field has been set.

### GetAvatarKey

`func (o *Users) GetAvatarKey() string`

GetAvatarKey returns the AvatarKey field if non-nil, zero value otherwise.

### GetAvatarKeyOk

`func (o *Users) GetAvatarKeyOk() (*string, bool)`

GetAvatarKeyOk returns a tuple with the AvatarKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvatarKey

`func (o *Users) SetAvatarKey(v string)`

SetAvatarKey sets AvatarKey field to given value.

### HasAvatarKey

`func (o *Users) HasAvatarKey() bool`

HasAvatarKey returns a boolean if a field has been set.

### GetExternalApiToken

`func (o *Users) GetExternalApiToken() string`

GetExternalApiToken returns the ExternalApiToken field if non-nil, zero value otherwise.

### GetExternalApiTokenOk

`func (o *Users) GetExternalApiTokenOk() (*string, bool)`

GetExternalApiTokenOk returns a tuple with the ExternalApiToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalApiToken

`func (o *Users) SetExternalApiToken(v string)`

SetExternalApiToken sets ExternalApiToken field to given value.

### HasExternalApiToken

`func (o *Users) HasExternalApiToken() bool`

HasExternalApiToken returns a boolean if a field has been set.

### GetAzureConversationData

`func (o *Users) GetAzureConversationData() interface{}`

GetAzureConversationData returns the AzureConversationData field if non-nil, zero value otherwise.

### GetAzureConversationDataOk

`func (o *Users) GetAzureConversationDataOk() (*interface{}, bool)`

GetAzureConversationDataOk returns a tuple with the AzureConversationData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAzureConversationData

`func (o *Users) SetAzureConversationData(v interface{})`

SetAzureConversationData sets AzureConversationData field to given value.

### HasAzureConversationData

`func (o *Users) HasAzureConversationData() bool`

HasAzureConversationData returns a boolean if a field has been set.

### SetAzureConversationDataNil

`func (o *Users) SetAzureConversationDataNil(b bool)`

 SetAzureConversationDataNil sets the value for AzureConversationData to be an explicit nil

### UnsetAzureConversationData
`func (o *Users) UnsetAzureConversationData()`

UnsetAzureConversationData ensures that no value is present for AzureConversationData, not even an explicit nil
### GetScimUserActive

`func (o *Users) GetScimUserActive() string`

GetScimUserActive returns the ScimUserActive field if non-nil, zero value otherwise.

### GetScimUserActiveOk

`func (o *Users) GetScimUserActiveOk() (*string, bool)`

GetScimUserActiveOk returns a tuple with the ScimUserActive field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScimUserActive

`func (o *Users) SetScimUserActive(v string)`

SetScimUserActive sets ScimUserActive field to given value.

### HasScimUserActive

`func (o *Users) HasScimUserActive() bool`

HasScimUserActive returns a boolean if a field has been set.

### GetTeamsNotificationsEnabled

`func (o *Users) GetTeamsNotificationsEnabled() bool`

GetTeamsNotificationsEnabled returns the TeamsNotificationsEnabled field if non-nil, zero value otherwise.

### GetTeamsNotificationsEnabledOk

`func (o *Users) GetTeamsNotificationsEnabledOk() (*bool, bool)`

GetTeamsNotificationsEnabledOk returns a tuple with the TeamsNotificationsEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTeamsNotificationsEnabled

`func (o *Users) SetTeamsNotificationsEnabled(v bool)`

SetTeamsNotificationsEnabled sets TeamsNotificationsEnabled field to given value.

### HasTeamsNotificationsEnabled

`func (o *Users) HasTeamsNotificationsEnabled() bool`

HasTeamsNotificationsEnabled returns a boolean if a field has been set.

### GetCoreModules

`func (o *Users) GetCoreModules() interface{}`

GetCoreModules returns the CoreModules field if non-nil, zero value otherwise.

### GetCoreModulesOk

`func (o *Users) GetCoreModulesOk() (*interface{}, bool)`

GetCoreModulesOk returns a tuple with the CoreModules field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCoreModules

`func (o *Users) SetCoreModules(v interface{})`

SetCoreModules sets CoreModules field to given value.


### SetCoreModulesNil

`func (o *Users) SetCoreModulesNil(b bool)`

 SetCoreModulesNil sets the value for CoreModules to be an explicit nil

### UnsetCoreModules
`func (o *Users) UnsetCoreModules()`

UnsetCoreModules ensures that no value is present for CoreModules, not even an explicit nil
### GetLocale

`func (o *Users) GetLocale() string`

GetLocale returns the Locale field if non-nil, zero value otherwise.

### GetLocaleOk

`func (o *Users) GetLocaleOk() (*string, bool)`

GetLocaleOk returns a tuple with the Locale field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocale

`func (o *Users) SetLocale(v string)`

SetLocale sets Locale field to given value.

### HasLocale

`func (o *Users) HasLocale() bool`

HasLocale returns a boolean if a field has been set.

### GetAccountExpirationDate

`func (o *Users) GetAccountExpirationDate() string`

GetAccountExpirationDate returns the AccountExpirationDate field if non-nil, zero value otherwise.

### GetAccountExpirationDateOk

`func (o *Users) GetAccountExpirationDateOk() (*string, bool)`

GetAccountExpirationDateOk returns a tuple with the AccountExpirationDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountExpirationDate

`func (o *Users) SetAccountExpirationDate(v string)`

SetAccountExpirationDate sets AccountExpirationDate field to given value.

### HasAccountExpirationDate

`func (o *Users) HasAccountExpirationDate() bool`

HasAccountExpirationDate returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


