# WorkspaceDeleteResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Message** | **string** |  | 
**WorkspaceId** | **string** |  | 
**WorkspaceName** | **string** |  | 
**DeletedCounts** | **map[string]int32** |  | 

## Methods

### NewWorkspaceDeleteResponse

`func NewWorkspaceDeleteResponse(message string, workspaceId string, workspaceName string, deletedCounts map[string]int32, ) *WorkspaceDeleteResponse`

NewWorkspaceDeleteResponse instantiates a new WorkspaceDeleteResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWorkspaceDeleteResponseWithDefaults

`func NewWorkspaceDeleteResponseWithDefaults() *WorkspaceDeleteResponse`

NewWorkspaceDeleteResponseWithDefaults instantiates a new WorkspaceDeleteResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMessage

`func (o *WorkspaceDeleteResponse) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *WorkspaceDeleteResponse) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *WorkspaceDeleteResponse) SetMessage(v string)`

SetMessage sets Message field to given value.


### GetWorkspaceId

`func (o *WorkspaceDeleteResponse) GetWorkspaceId() string`

GetWorkspaceId returns the WorkspaceId field if non-nil, zero value otherwise.

### GetWorkspaceIdOk

`func (o *WorkspaceDeleteResponse) GetWorkspaceIdOk() (*string, bool)`

GetWorkspaceIdOk returns a tuple with the WorkspaceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWorkspaceId

`func (o *WorkspaceDeleteResponse) SetWorkspaceId(v string)`

SetWorkspaceId sets WorkspaceId field to given value.


### GetWorkspaceName

`func (o *WorkspaceDeleteResponse) GetWorkspaceName() string`

GetWorkspaceName returns the WorkspaceName field if non-nil, zero value otherwise.

### GetWorkspaceNameOk

`func (o *WorkspaceDeleteResponse) GetWorkspaceNameOk() (*string, bool)`

GetWorkspaceNameOk returns a tuple with the WorkspaceName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWorkspaceName

`func (o *WorkspaceDeleteResponse) SetWorkspaceName(v string)`

SetWorkspaceName sets WorkspaceName field to given value.


### GetDeletedCounts

`func (o *WorkspaceDeleteResponse) GetDeletedCounts() map[string]int32`

GetDeletedCounts returns the DeletedCounts field if non-nil, zero value otherwise.

### GetDeletedCountsOk

`func (o *WorkspaceDeleteResponse) GetDeletedCountsOk() (*map[string]int32, bool)`

GetDeletedCountsOk returns a tuple with the DeletedCounts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeletedCounts

`func (o *WorkspaceDeleteResponse) SetDeletedCounts(v map[string]int32)`

SetDeletedCounts sets DeletedCounts field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


