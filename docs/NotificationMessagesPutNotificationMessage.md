# NotificationMessagesPutNotificationMessage

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int32** | Note: This is a Primary Key.&lt;pk/&gt; | [optional] [readonly] 
**FromUserId** | Pointer to **int32** | Note: This is a Foreign Key to &#x60;users.id&#x60;.&lt;fk table&#x3D;&#39;users&#39; column&#x3D;&#39;id&#39;/&gt; | [optional] 
**ToUserId** | **int32** | Note: This is a Foreign Key to &#x60;users.id&#x60;.&lt;fk table&#x3D;&#39;users&#39; column&#x3D;&#39;id&#39;/&gt; | 
**Type** | Pointer to **string** |  | [optional] 
**Status** | Pointer to **string** |  | [optional] 
**IsRead** | **bool** |  | 
**Message** | **string** |  | 
**RawData** | Pointer to **string** |  | [optional] 
**CreatedAt** | Pointer to **string** |  | [optional] [readonly] 
**UpdatedAt** | Pointer to **string** |  | [optional] [readonly] 
**DeletedAt** | Pointer to **string** |  | [optional] [readonly] 

## Methods

### NewNotificationMessagesPutNotificationMessage

`func NewNotificationMessagesPutNotificationMessage(toUserId int32, isRead bool, message string, ) *NotificationMessagesPutNotificationMessage`

NewNotificationMessagesPutNotificationMessage instantiates a new NotificationMessagesPutNotificationMessage object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNotificationMessagesPutNotificationMessageWithDefaults

`func NewNotificationMessagesPutNotificationMessageWithDefaults() *NotificationMessagesPutNotificationMessage`

NewNotificationMessagesPutNotificationMessageWithDefaults instantiates a new NotificationMessagesPutNotificationMessage object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *NotificationMessagesPutNotificationMessage) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *NotificationMessagesPutNotificationMessage) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *NotificationMessagesPutNotificationMessage) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *NotificationMessagesPutNotificationMessage) HasId() bool`

HasId returns a boolean if a field has been set.

### GetFromUserId

`func (o *NotificationMessagesPutNotificationMessage) GetFromUserId() int32`

GetFromUserId returns the FromUserId field if non-nil, zero value otherwise.

### GetFromUserIdOk

`func (o *NotificationMessagesPutNotificationMessage) GetFromUserIdOk() (*int32, bool)`

GetFromUserIdOk returns a tuple with the FromUserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFromUserId

`func (o *NotificationMessagesPutNotificationMessage) SetFromUserId(v int32)`

SetFromUserId sets FromUserId field to given value.

### HasFromUserId

`func (o *NotificationMessagesPutNotificationMessage) HasFromUserId() bool`

HasFromUserId returns a boolean if a field has been set.

### GetToUserId

`func (o *NotificationMessagesPutNotificationMessage) GetToUserId() int32`

GetToUserId returns the ToUserId field if non-nil, zero value otherwise.

### GetToUserIdOk

`func (o *NotificationMessagesPutNotificationMessage) GetToUserIdOk() (*int32, bool)`

GetToUserIdOk returns a tuple with the ToUserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToUserId

`func (o *NotificationMessagesPutNotificationMessage) SetToUserId(v int32)`

SetToUserId sets ToUserId field to given value.


### GetType

`func (o *NotificationMessagesPutNotificationMessage) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *NotificationMessagesPutNotificationMessage) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *NotificationMessagesPutNotificationMessage) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *NotificationMessagesPutNotificationMessage) HasType() bool`

HasType returns a boolean if a field has been set.

### GetStatus

`func (o *NotificationMessagesPutNotificationMessage) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *NotificationMessagesPutNotificationMessage) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *NotificationMessagesPutNotificationMessage) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *NotificationMessagesPutNotificationMessage) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetIsRead

`func (o *NotificationMessagesPutNotificationMessage) GetIsRead() bool`

GetIsRead returns the IsRead field if non-nil, zero value otherwise.

### GetIsReadOk

`func (o *NotificationMessagesPutNotificationMessage) GetIsReadOk() (*bool, bool)`

GetIsReadOk returns a tuple with the IsRead field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsRead

`func (o *NotificationMessagesPutNotificationMessage) SetIsRead(v bool)`

SetIsRead sets IsRead field to given value.


### GetMessage

`func (o *NotificationMessagesPutNotificationMessage) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *NotificationMessagesPutNotificationMessage) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *NotificationMessagesPutNotificationMessage) SetMessage(v string)`

SetMessage sets Message field to given value.


### GetRawData

`func (o *NotificationMessagesPutNotificationMessage) GetRawData() string`

GetRawData returns the RawData field if non-nil, zero value otherwise.

### GetRawDataOk

`func (o *NotificationMessagesPutNotificationMessage) GetRawDataOk() (*string, bool)`

GetRawDataOk returns a tuple with the RawData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRawData

`func (o *NotificationMessagesPutNotificationMessage) SetRawData(v string)`

SetRawData sets RawData field to given value.

### HasRawData

`func (o *NotificationMessagesPutNotificationMessage) HasRawData() bool`

HasRawData returns a boolean if a field has been set.

### GetCreatedAt

`func (o *NotificationMessagesPutNotificationMessage) GetCreatedAt() string`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *NotificationMessagesPutNotificationMessage) GetCreatedAtOk() (*string, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *NotificationMessagesPutNotificationMessage) SetCreatedAt(v string)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *NotificationMessagesPutNotificationMessage) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *NotificationMessagesPutNotificationMessage) GetUpdatedAt() string`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *NotificationMessagesPutNotificationMessage) GetUpdatedAtOk() (*string, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *NotificationMessagesPutNotificationMessage) SetUpdatedAt(v string)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *NotificationMessagesPutNotificationMessage) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.

### GetDeletedAt

`func (o *NotificationMessagesPutNotificationMessage) GetDeletedAt() string`

GetDeletedAt returns the DeletedAt field if non-nil, zero value otherwise.

### GetDeletedAtOk

`func (o *NotificationMessagesPutNotificationMessage) GetDeletedAtOk() (*string, bool)`

GetDeletedAtOk returns a tuple with the DeletedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeletedAt

`func (o *NotificationMessagesPutNotificationMessage) SetDeletedAt(v string)`

SetDeletedAt sets DeletedAt field to given value.

### HasDeletedAt

`func (o *NotificationMessagesPutNotificationMessage) HasDeletedAt() bool`

HasDeletedAt returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


