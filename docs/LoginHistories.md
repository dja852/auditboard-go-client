# LoginHistories

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int32** | Note: This is a Primary Key.&lt;pk/&gt; | [optional] [readonly] 
**Timestamp** | Pointer to **string** |  | [optional] 
**UserId** | Pointer to **int32** | Note: This is a Foreign Key to &#x60;users.id&#x60;.&lt;fk table&#x3D;&#39;users&#39; column&#x3D;&#39;id&#39;/&gt; | [optional] 
**IpAddress** | Pointer to **string** |  | [optional] 
**Action** | Pointer to **string** |  | [optional] 
**Message** | Pointer to **string** |  | [optional] 
**Location** | Pointer to **string** |  | [optional] 
**CreatedAt** | Pointer to **string** |  | [optional] [readonly] 
**UpdatedAt** | Pointer to **string** |  | [optional] [readonly] 
**DeletedAt** | Pointer to **string** |  | [optional] [readonly] 
**LoginEmail** | Pointer to **string** |  | [optional] 

## Methods

### NewLoginHistories

`func NewLoginHistories() *LoginHistories`

NewLoginHistories instantiates a new LoginHistories object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLoginHistoriesWithDefaults

`func NewLoginHistoriesWithDefaults() *LoginHistories`

NewLoginHistoriesWithDefaults instantiates a new LoginHistories object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *LoginHistories) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *LoginHistories) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *LoginHistories) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *LoginHistories) HasId() bool`

HasId returns a boolean if a field has been set.

### GetTimestamp

`func (o *LoginHistories) GetTimestamp() string`

GetTimestamp returns the Timestamp field if non-nil, zero value otherwise.

### GetTimestampOk

`func (o *LoginHistories) GetTimestampOk() (*string, bool)`

GetTimestampOk returns a tuple with the Timestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimestamp

`func (o *LoginHistories) SetTimestamp(v string)`

SetTimestamp sets Timestamp field to given value.

### HasTimestamp

`func (o *LoginHistories) HasTimestamp() bool`

HasTimestamp returns a boolean if a field has been set.

### GetUserId

`func (o *LoginHistories) GetUserId() int32`

GetUserId returns the UserId field if non-nil, zero value otherwise.

### GetUserIdOk

`func (o *LoginHistories) GetUserIdOk() (*int32, bool)`

GetUserIdOk returns a tuple with the UserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserId

`func (o *LoginHistories) SetUserId(v int32)`

SetUserId sets UserId field to given value.

### HasUserId

`func (o *LoginHistories) HasUserId() bool`

HasUserId returns a boolean if a field has been set.

### GetIpAddress

`func (o *LoginHistories) GetIpAddress() string`

GetIpAddress returns the IpAddress field if non-nil, zero value otherwise.

### GetIpAddressOk

`func (o *LoginHistories) GetIpAddressOk() (*string, bool)`

GetIpAddressOk returns a tuple with the IpAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpAddress

`func (o *LoginHistories) SetIpAddress(v string)`

SetIpAddress sets IpAddress field to given value.

### HasIpAddress

`func (o *LoginHistories) HasIpAddress() bool`

HasIpAddress returns a boolean if a field has been set.

### GetAction

`func (o *LoginHistories) GetAction() string`

GetAction returns the Action field if non-nil, zero value otherwise.

### GetActionOk

`func (o *LoginHistories) GetActionOk() (*string, bool)`

GetActionOk returns a tuple with the Action field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAction

`func (o *LoginHistories) SetAction(v string)`

SetAction sets Action field to given value.

### HasAction

`func (o *LoginHistories) HasAction() bool`

HasAction returns a boolean if a field has been set.

### GetMessage

`func (o *LoginHistories) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *LoginHistories) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *LoginHistories) SetMessage(v string)`

SetMessage sets Message field to given value.

### HasMessage

`func (o *LoginHistories) HasMessage() bool`

HasMessage returns a boolean if a field has been set.

### GetLocation

`func (o *LoginHistories) GetLocation() string`

GetLocation returns the Location field if non-nil, zero value otherwise.

### GetLocationOk

`func (o *LoginHistories) GetLocationOk() (*string, bool)`

GetLocationOk returns a tuple with the Location field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocation

`func (o *LoginHistories) SetLocation(v string)`

SetLocation sets Location field to given value.

### HasLocation

`func (o *LoginHistories) HasLocation() bool`

HasLocation returns a boolean if a field has been set.

### GetCreatedAt

`func (o *LoginHistories) GetCreatedAt() string`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *LoginHistories) GetCreatedAtOk() (*string, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *LoginHistories) SetCreatedAt(v string)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *LoginHistories) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *LoginHistories) GetUpdatedAt() string`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *LoginHistories) GetUpdatedAtOk() (*string, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *LoginHistories) SetUpdatedAt(v string)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *LoginHistories) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.

### GetDeletedAt

`func (o *LoginHistories) GetDeletedAt() string`

GetDeletedAt returns the DeletedAt field if non-nil, zero value otherwise.

### GetDeletedAtOk

`func (o *LoginHistories) GetDeletedAtOk() (*string, bool)`

GetDeletedAtOk returns a tuple with the DeletedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeletedAt

`func (o *LoginHistories) SetDeletedAt(v string)`

SetDeletedAt sets DeletedAt field to given value.

### HasDeletedAt

`func (o *LoginHistories) HasDeletedAt() bool`

HasDeletedAt returns a boolean if a field has been set.

### GetLoginEmail

`func (o *LoginHistories) GetLoginEmail() string`

GetLoginEmail returns the LoginEmail field if non-nil, zero value otherwise.

### GetLoginEmailOk

`func (o *LoginHistories) GetLoginEmailOk() (*string, bool)`

GetLoginEmailOk returns a tuple with the LoginEmail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLoginEmail

`func (o *LoginHistories) SetLoginEmail(v string)`

SetLoginEmail sets LoginEmail field to given value.

### HasLoginEmail

`func (o *LoginHistories) HasLoginEmail() bool`

HasLoginEmail returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


