# CommentsPostRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Comment** | Pointer to [**Comments**](Comments.md) |  | [optional] 

## Methods

### NewCommentsPostRequest

`func NewCommentsPostRequest() *CommentsPostRequest`

NewCommentsPostRequest instantiates a new CommentsPostRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCommentsPostRequestWithDefaults

`func NewCommentsPostRequestWithDefaults() *CommentsPostRequest`

NewCommentsPostRequestWithDefaults instantiates a new CommentsPostRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetComment

`func (o *CommentsPostRequest) GetComment() Comments`

GetComment returns the Comment field if non-nil, zero value otherwise.

### GetCommentOk

`func (o *CommentsPostRequest) GetCommentOk() (*Comments, bool)`

GetCommentOk returns a tuple with the Comment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComment

`func (o *CommentsPostRequest) SetComment(v Comments)`

SetComment sets Comment field to given value.

### HasComment

`func (o *CommentsPostRequest) HasComment() bool`

HasComment returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


