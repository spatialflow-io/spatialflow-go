# UserWorkspaceResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Message** | **string** |  | 
**UserId** | **string** |  | 
**WorkspaceId** | Pointer to **NullableString** |  | [optional] 
**Role** | Pointer to **NullableString** |  | [optional] 
**PreviousWorkspaceName** | Pointer to **NullableString** |  | [optional] 

## Methods

### NewUserWorkspaceResponse

`func NewUserWorkspaceResponse(message string, userId string, ) *UserWorkspaceResponse`

NewUserWorkspaceResponse instantiates a new UserWorkspaceResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUserWorkspaceResponseWithDefaults

`func NewUserWorkspaceResponseWithDefaults() *UserWorkspaceResponse`

NewUserWorkspaceResponseWithDefaults instantiates a new UserWorkspaceResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMessage

`func (o *UserWorkspaceResponse) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *UserWorkspaceResponse) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *UserWorkspaceResponse) SetMessage(v string)`

SetMessage sets Message field to given value.


### GetUserId

`func (o *UserWorkspaceResponse) GetUserId() string`

GetUserId returns the UserId field if non-nil, zero value otherwise.

### GetUserIdOk

`func (o *UserWorkspaceResponse) GetUserIdOk() (*string, bool)`

GetUserIdOk returns a tuple with the UserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserId

`func (o *UserWorkspaceResponse) SetUserId(v string)`

SetUserId sets UserId field to given value.


### GetWorkspaceId

`func (o *UserWorkspaceResponse) GetWorkspaceId() string`

GetWorkspaceId returns the WorkspaceId field if non-nil, zero value otherwise.

### GetWorkspaceIdOk

`func (o *UserWorkspaceResponse) GetWorkspaceIdOk() (*string, bool)`

GetWorkspaceIdOk returns a tuple with the WorkspaceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWorkspaceId

`func (o *UserWorkspaceResponse) SetWorkspaceId(v string)`

SetWorkspaceId sets WorkspaceId field to given value.

### HasWorkspaceId

`func (o *UserWorkspaceResponse) HasWorkspaceId() bool`

HasWorkspaceId returns a boolean if a field has been set.

### SetWorkspaceIdNil

`func (o *UserWorkspaceResponse) SetWorkspaceIdNil(b bool)`

 SetWorkspaceIdNil sets the value for WorkspaceId to be an explicit nil

### UnsetWorkspaceId
`func (o *UserWorkspaceResponse) UnsetWorkspaceId()`

UnsetWorkspaceId ensures that no value is present for WorkspaceId, not even an explicit nil
### GetRole

`func (o *UserWorkspaceResponse) GetRole() string`

GetRole returns the Role field if non-nil, zero value otherwise.

### GetRoleOk

`func (o *UserWorkspaceResponse) GetRoleOk() (*string, bool)`

GetRoleOk returns a tuple with the Role field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRole

`func (o *UserWorkspaceResponse) SetRole(v string)`

SetRole sets Role field to given value.

### HasRole

`func (o *UserWorkspaceResponse) HasRole() bool`

HasRole returns a boolean if a field has been set.

### SetRoleNil

`func (o *UserWorkspaceResponse) SetRoleNil(b bool)`

 SetRoleNil sets the value for Role to be an explicit nil

### UnsetRole
`func (o *UserWorkspaceResponse) UnsetRole()`

UnsetRole ensures that no value is present for Role, not even an explicit nil
### GetPreviousWorkspaceName

`func (o *UserWorkspaceResponse) GetPreviousWorkspaceName() string`

GetPreviousWorkspaceName returns the PreviousWorkspaceName field if non-nil, zero value otherwise.

### GetPreviousWorkspaceNameOk

`func (o *UserWorkspaceResponse) GetPreviousWorkspaceNameOk() (*string, bool)`

GetPreviousWorkspaceNameOk returns a tuple with the PreviousWorkspaceName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPreviousWorkspaceName

`func (o *UserWorkspaceResponse) SetPreviousWorkspaceName(v string)`

SetPreviousWorkspaceName sets PreviousWorkspaceName field to given value.

### HasPreviousWorkspaceName

`func (o *UserWorkspaceResponse) HasPreviousWorkspaceName() bool`

HasPreviousWorkspaceName returns a boolean if a field has been set.

### SetPreviousWorkspaceNameNil

`func (o *UserWorkspaceResponse) SetPreviousWorkspaceNameNil(b bool)`

 SetPreviousWorkspaceNameNil sets the value for PreviousWorkspaceName to be an explicit nil

### UnsetPreviousWorkspaceName
`func (o *UserWorkspaceResponse) UnsetPreviousWorkspaceName()`

UnsetPreviousWorkspaceName ensures that no value is present for PreviousWorkspaceName, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


