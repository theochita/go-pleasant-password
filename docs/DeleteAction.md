# DeleteAction

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Action** | Pointer to **string** |  | [optional] 
**Comment** | Pointer to **string** |  | [optional] 

## Methods

### NewDeleteAction

`func NewDeleteAction() *DeleteAction`

NewDeleteAction instantiates a new DeleteAction object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDeleteActionWithDefaults

`func NewDeleteActionWithDefaults() *DeleteAction`

NewDeleteActionWithDefaults instantiates a new DeleteAction object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAction

`func (o *DeleteAction) GetAction() string`

GetAction returns the Action field if non-nil, zero value otherwise.

### GetActionOk

`func (o *DeleteAction) GetActionOk() (*string, bool)`

GetActionOk returns a tuple with the Action field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAction

`func (o *DeleteAction) SetAction(v string)`

SetAction sets Action field to given value.

### HasAction

`func (o *DeleteAction) HasAction() bool`

HasAction returns a boolean if a field has been set.

### GetComment

`func (o *DeleteAction) GetComment() string`

GetComment returns the Comment field if non-nil, zero value otherwise.

### GetCommentOk

`func (o *DeleteAction) GetCommentOk() (*string, bool)`

GetCommentOk returns a tuple with the Comment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComment

`func (o *DeleteAction) SetComment(v string)`

SetComment sets Comment field to given value.

### HasComment

`func (o *DeleteAction) HasComment() bool`

HasComment returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


