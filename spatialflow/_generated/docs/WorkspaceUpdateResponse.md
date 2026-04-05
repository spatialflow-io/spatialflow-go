# WorkspaceUpdateResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Message** | **string** |  | 
**WorkspaceId** | **string** |  | 
**Name** | **string** |  | 
**Changes** | **map[string]interface{}** |  | 

## Methods

### NewWorkspaceUpdateResponse

`func NewWorkspaceUpdateResponse(message string, workspaceId string, name string, changes map[string]interface{}, ) *WorkspaceUpdateResponse`

NewWorkspaceUpdateResponse instantiates a new WorkspaceUpdateResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWorkspaceUpdateResponseWithDefaults

`func NewWorkspaceUpdateResponseWithDefaults() *WorkspaceUpdateResponse`

NewWorkspaceUpdateResponseWithDefaults instantiates a new WorkspaceUpdateResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMessage

`func (o *WorkspaceUpdateResponse) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *WorkspaceUpdateResponse) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *WorkspaceUpdateResponse) SetMessage(v string)`

SetMessage sets Message field to given value.


### GetWorkspaceId

`func (o *WorkspaceUpdateResponse) GetWorkspaceId() string`

GetWorkspaceId returns the WorkspaceId field if non-nil, zero value otherwise.

### GetWorkspaceIdOk

`func (o *WorkspaceUpdateResponse) GetWorkspaceIdOk() (*string, bool)`

GetWorkspaceIdOk returns a tuple with the WorkspaceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWorkspaceId

`func (o *WorkspaceUpdateResponse) SetWorkspaceId(v string)`

SetWorkspaceId sets WorkspaceId field to given value.


### GetName

`func (o *WorkspaceUpdateResponse) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *WorkspaceUpdateResponse) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *WorkspaceUpdateResponse) SetName(v string)`

SetName sets Name field to given value.


### GetChanges

`func (o *WorkspaceUpdateResponse) GetChanges() map[string]interface{}`

GetChanges returns the Changes field if non-nil, zero value otherwise.

### GetChangesOk

`func (o *WorkspaceUpdateResponse) GetChangesOk() (*map[string]interface{}, bool)`

GetChangesOk returns a tuple with the Changes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChanges

`func (o *WorkspaceUpdateResponse) SetChanges(v map[string]interface{})`

SetChanges sets Changes field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


