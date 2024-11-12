# GlobalAuditsPutGlobalAudit

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int32** |  | [optional] [readonly] 
**CreatedAt** | Pointer to **string** |  | [optional] [readonly] 
**UpdatedAt** | Pointer to **string** |  | [optional] [readonly] 
**UserId** | Pointer to **int32** | Note: This is a Foreign Key to &#x60;users.id&#x60;.&lt;fk table&#x3D;&#39;users&#39; column&#x3D;&#39;id&#39;/&gt; | [optional] 
**LoggableId** | **int32** |  | 
**LoggableType** | **string** |  | 
**Action** | **string** |  | 
**Field** | **string** |  | 
**FromValue** | Pointer to **string** |  | [optional] 
**ToValue** | Pointer to **string** |  | [optional] 
**Message** | Pointer to **string** |  | [optional] 

## Methods

### NewGlobalAuditsPutGlobalAudit

`func NewGlobalAuditsPutGlobalAudit(loggableId int32, loggableType string, action string, field string, ) *GlobalAuditsPutGlobalAudit`

NewGlobalAuditsPutGlobalAudit instantiates a new GlobalAuditsPutGlobalAudit object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGlobalAuditsPutGlobalAuditWithDefaults

`func NewGlobalAuditsPutGlobalAuditWithDefaults() *GlobalAuditsPutGlobalAudit`

NewGlobalAuditsPutGlobalAuditWithDefaults instantiates a new GlobalAuditsPutGlobalAudit object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *GlobalAuditsPutGlobalAudit) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *GlobalAuditsPutGlobalAudit) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *GlobalAuditsPutGlobalAudit) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *GlobalAuditsPutGlobalAudit) HasId() bool`

HasId returns a boolean if a field has been set.

### GetCreatedAt

`func (o *GlobalAuditsPutGlobalAudit) GetCreatedAt() string`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *GlobalAuditsPutGlobalAudit) GetCreatedAtOk() (*string, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *GlobalAuditsPutGlobalAudit) SetCreatedAt(v string)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *GlobalAuditsPutGlobalAudit) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *GlobalAuditsPutGlobalAudit) GetUpdatedAt() string`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *GlobalAuditsPutGlobalAudit) GetUpdatedAtOk() (*string, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *GlobalAuditsPutGlobalAudit) SetUpdatedAt(v string)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *GlobalAuditsPutGlobalAudit) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.

### GetUserId

`func (o *GlobalAuditsPutGlobalAudit) GetUserId() int32`

GetUserId returns the UserId field if non-nil, zero value otherwise.

### GetUserIdOk

`func (o *GlobalAuditsPutGlobalAudit) GetUserIdOk() (*int32, bool)`

GetUserIdOk returns a tuple with the UserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserId

`func (o *GlobalAuditsPutGlobalAudit) SetUserId(v int32)`

SetUserId sets UserId field to given value.

### HasUserId

`func (o *GlobalAuditsPutGlobalAudit) HasUserId() bool`

HasUserId returns a boolean if a field has been set.

### GetLoggableId

`func (o *GlobalAuditsPutGlobalAudit) GetLoggableId() int32`

GetLoggableId returns the LoggableId field if non-nil, zero value otherwise.

### GetLoggableIdOk

`func (o *GlobalAuditsPutGlobalAudit) GetLoggableIdOk() (*int32, bool)`

GetLoggableIdOk returns a tuple with the LoggableId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLoggableId

`func (o *GlobalAuditsPutGlobalAudit) SetLoggableId(v int32)`

SetLoggableId sets LoggableId field to given value.


### GetLoggableType

`func (o *GlobalAuditsPutGlobalAudit) GetLoggableType() string`

GetLoggableType returns the LoggableType field if non-nil, zero value otherwise.

### GetLoggableTypeOk

`func (o *GlobalAuditsPutGlobalAudit) GetLoggableTypeOk() (*string, bool)`

GetLoggableTypeOk returns a tuple with the LoggableType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLoggableType

`func (o *GlobalAuditsPutGlobalAudit) SetLoggableType(v string)`

SetLoggableType sets LoggableType field to given value.


### GetAction

`func (o *GlobalAuditsPutGlobalAudit) GetAction() string`

GetAction returns the Action field if non-nil, zero value otherwise.

### GetActionOk

`func (o *GlobalAuditsPutGlobalAudit) GetActionOk() (*string, bool)`

GetActionOk returns a tuple with the Action field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAction

`func (o *GlobalAuditsPutGlobalAudit) SetAction(v string)`

SetAction sets Action field to given value.


### GetField

`func (o *GlobalAuditsPutGlobalAudit) GetField() string`

GetField returns the Field field if non-nil, zero value otherwise.

### GetFieldOk

`func (o *GlobalAuditsPutGlobalAudit) GetFieldOk() (*string, bool)`

GetFieldOk returns a tuple with the Field field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetField

`func (o *GlobalAuditsPutGlobalAudit) SetField(v string)`

SetField sets Field field to given value.


### GetFromValue

`func (o *GlobalAuditsPutGlobalAudit) GetFromValue() string`

GetFromValue returns the FromValue field if non-nil, zero value otherwise.

### GetFromValueOk

`func (o *GlobalAuditsPutGlobalAudit) GetFromValueOk() (*string, bool)`

GetFromValueOk returns a tuple with the FromValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFromValue

`func (o *GlobalAuditsPutGlobalAudit) SetFromValue(v string)`

SetFromValue sets FromValue field to given value.

### HasFromValue

`func (o *GlobalAuditsPutGlobalAudit) HasFromValue() bool`

HasFromValue returns a boolean if a field has been set.

### GetToValue

`func (o *GlobalAuditsPutGlobalAudit) GetToValue() string`

GetToValue returns the ToValue field if non-nil, zero value otherwise.

### GetToValueOk

`func (o *GlobalAuditsPutGlobalAudit) GetToValueOk() (*string, bool)`

GetToValueOk returns a tuple with the ToValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToValue

`func (o *GlobalAuditsPutGlobalAudit) SetToValue(v string)`

SetToValue sets ToValue field to given value.

### HasToValue

`func (o *GlobalAuditsPutGlobalAudit) HasToValue() bool`

HasToValue returns a boolean if a field has been set.

### GetMessage

`func (o *GlobalAuditsPutGlobalAudit) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *GlobalAuditsPutGlobalAudit) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *GlobalAuditsPutGlobalAudit) SetMessage(v string)`

SetMessage sets Message field to given value.

### HasMessage

`func (o *GlobalAuditsPutGlobalAudit) HasMessage() bool`

HasMessage returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


