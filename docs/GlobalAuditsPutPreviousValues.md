# GlobalAuditsPutPreviousValues

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int32** |  | [optional] [readonly] 
**CreatedAt** | Pointer to **string** |  | [optional] [readonly] 
**UpdatedAt** | Pointer to **string** |  | [optional] [readonly] 
**UserId** | Pointer to **int32** | Note: This is a Foreign Key to &#x60;users.id&#x60;.&lt;fk table&#x3D;&#39;users&#39; column&#x3D;&#39;id&#39;/&gt; | [optional] 
**LoggableId** | Pointer to **int32** |  | [optional] 
**LoggableType** | Pointer to **string** |  | [optional] 
**Action** | Pointer to **string** |  | [optional] 
**Field** | Pointer to **string** |  | [optional] 
**FromValue** | Pointer to **string** |  | [optional] 
**ToValue** | Pointer to **string** |  | [optional] 
**Message** | Pointer to **string** |  | [optional] 

## Methods

### NewGlobalAuditsPutPreviousValues

`func NewGlobalAuditsPutPreviousValues() *GlobalAuditsPutPreviousValues`

NewGlobalAuditsPutPreviousValues instantiates a new GlobalAuditsPutPreviousValues object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGlobalAuditsPutPreviousValuesWithDefaults

`func NewGlobalAuditsPutPreviousValuesWithDefaults() *GlobalAuditsPutPreviousValues`

NewGlobalAuditsPutPreviousValuesWithDefaults instantiates a new GlobalAuditsPutPreviousValues object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *GlobalAuditsPutPreviousValues) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *GlobalAuditsPutPreviousValues) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *GlobalAuditsPutPreviousValues) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *GlobalAuditsPutPreviousValues) HasId() bool`

HasId returns a boolean if a field has been set.

### GetCreatedAt

`func (o *GlobalAuditsPutPreviousValues) GetCreatedAt() string`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *GlobalAuditsPutPreviousValues) GetCreatedAtOk() (*string, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *GlobalAuditsPutPreviousValues) SetCreatedAt(v string)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *GlobalAuditsPutPreviousValues) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *GlobalAuditsPutPreviousValues) GetUpdatedAt() string`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *GlobalAuditsPutPreviousValues) GetUpdatedAtOk() (*string, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *GlobalAuditsPutPreviousValues) SetUpdatedAt(v string)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *GlobalAuditsPutPreviousValues) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.

### GetUserId

`func (o *GlobalAuditsPutPreviousValues) GetUserId() int32`

GetUserId returns the UserId field if non-nil, zero value otherwise.

### GetUserIdOk

`func (o *GlobalAuditsPutPreviousValues) GetUserIdOk() (*int32, bool)`

GetUserIdOk returns a tuple with the UserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserId

`func (o *GlobalAuditsPutPreviousValues) SetUserId(v int32)`

SetUserId sets UserId field to given value.

### HasUserId

`func (o *GlobalAuditsPutPreviousValues) HasUserId() bool`

HasUserId returns a boolean if a field has been set.

### GetLoggableId

`func (o *GlobalAuditsPutPreviousValues) GetLoggableId() int32`

GetLoggableId returns the LoggableId field if non-nil, zero value otherwise.

### GetLoggableIdOk

`func (o *GlobalAuditsPutPreviousValues) GetLoggableIdOk() (*int32, bool)`

GetLoggableIdOk returns a tuple with the LoggableId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLoggableId

`func (o *GlobalAuditsPutPreviousValues) SetLoggableId(v int32)`

SetLoggableId sets LoggableId field to given value.

### HasLoggableId

`func (o *GlobalAuditsPutPreviousValues) HasLoggableId() bool`

HasLoggableId returns a boolean if a field has been set.

### GetLoggableType

`func (o *GlobalAuditsPutPreviousValues) GetLoggableType() string`

GetLoggableType returns the LoggableType field if non-nil, zero value otherwise.

### GetLoggableTypeOk

`func (o *GlobalAuditsPutPreviousValues) GetLoggableTypeOk() (*string, bool)`

GetLoggableTypeOk returns a tuple with the LoggableType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLoggableType

`func (o *GlobalAuditsPutPreviousValues) SetLoggableType(v string)`

SetLoggableType sets LoggableType field to given value.

### HasLoggableType

`func (o *GlobalAuditsPutPreviousValues) HasLoggableType() bool`

HasLoggableType returns a boolean if a field has been set.

### GetAction

`func (o *GlobalAuditsPutPreviousValues) GetAction() string`

GetAction returns the Action field if non-nil, zero value otherwise.

### GetActionOk

`func (o *GlobalAuditsPutPreviousValues) GetActionOk() (*string, bool)`

GetActionOk returns a tuple with the Action field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAction

`func (o *GlobalAuditsPutPreviousValues) SetAction(v string)`

SetAction sets Action field to given value.

### HasAction

`func (o *GlobalAuditsPutPreviousValues) HasAction() bool`

HasAction returns a boolean if a field has been set.

### GetField

`func (o *GlobalAuditsPutPreviousValues) GetField() string`

GetField returns the Field field if non-nil, zero value otherwise.

### GetFieldOk

`func (o *GlobalAuditsPutPreviousValues) GetFieldOk() (*string, bool)`

GetFieldOk returns a tuple with the Field field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetField

`func (o *GlobalAuditsPutPreviousValues) SetField(v string)`

SetField sets Field field to given value.

### HasField

`func (o *GlobalAuditsPutPreviousValues) HasField() bool`

HasField returns a boolean if a field has been set.

### GetFromValue

`func (o *GlobalAuditsPutPreviousValues) GetFromValue() string`

GetFromValue returns the FromValue field if non-nil, zero value otherwise.

### GetFromValueOk

`func (o *GlobalAuditsPutPreviousValues) GetFromValueOk() (*string, bool)`

GetFromValueOk returns a tuple with the FromValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFromValue

`func (o *GlobalAuditsPutPreviousValues) SetFromValue(v string)`

SetFromValue sets FromValue field to given value.

### HasFromValue

`func (o *GlobalAuditsPutPreviousValues) HasFromValue() bool`

HasFromValue returns a boolean if a field has been set.

### GetToValue

`func (o *GlobalAuditsPutPreviousValues) GetToValue() string`

GetToValue returns the ToValue field if non-nil, zero value otherwise.

### GetToValueOk

`func (o *GlobalAuditsPutPreviousValues) GetToValueOk() (*string, bool)`

GetToValueOk returns a tuple with the ToValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToValue

`func (o *GlobalAuditsPutPreviousValues) SetToValue(v string)`

SetToValue sets ToValue field to given value.

### HasToValue

`func (o *GlobalAuditsPutPreviousValues) HasToValue() bool`

HasToValue returns a boolean if a field has been set.

### GetMessage

`func (o *GlobalAuditsPutPreviousValues) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *GlobalAuditsPutPreviousValues) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *GlobalAuditsPutPreviousValues) SetMessage(v string)`

SetMessage sets Message field to given value.

### HasMessage

`func (o *GlobalAuditsPutPreviousValues) HasMessage() bool`

HasMessage returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


