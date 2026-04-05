# MemberActionResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Message** | **string** |  | 
**UserId** | **string** |  | 
**WorkspaceId** | **string** |  | 
**NewRole** | Pointer to **NullableString** |  | [optional] 

## Methods

### NewMemberActionResponse

`func NewMemberActionResponse(message string, userId string, workspaceId string, ) *MemberActionResponse`

NewMemberActionResponse instantiates a new MemberActionResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMemberActionResponseWithDefaults

`func NewMemberActionResponseWithDefaults() *MemberActionResponse`

NewMemberActionResponseWithDefaults instantiates a new MemberActionResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMessage

`func (o *MemberActionResponse) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *MemberActionResponse) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *MemberActionResponse) SetMessage(v string)`

SetMessage sets Message field to given value.


### GetUserId

`func (o *MemberActionResponse) GetUserId() string`

GetUserId returns the UserId field if non-nil, zero value otherwise.

### GetUserIdOk

`func (o *MemberActionResponse) GetUserIdOk() (*string, bool)`

GetUserIdOk returns a tuple with the UserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserId

`func (o *MemberActionResponse) SetUserId(v string)`

SetUserId sets UserId field to given value.


### GetWorkspaceId

`func (o *MemberActionResponse) GetWorkspaceId() string`

GetWorkspaceId returns the WorkspaceId field if non-nil, zero value otherwise.

### GetWorkspaceIdOk

`func (o *MemberActionResponse) GetWorkspaceIdOk() (*string, bool)`

GetWorkspaceIdOk returns a tuple with the WorkspaceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWorkspaceId

`func (o *MemberActionResponse) SetWorkspaceId(v string)`

SetWorkspaceId sets WorkspaceId field to given value.


### GetNewRole

`func (o *MemberActionResponse) GetNewRole() string`

GetNewRole returns the NewRole field if non-nil, zero value otherwise.

### GetNewRoleOk

`func (o *MemberActionResponse) GetNewRoleOk() (*string, bool)`

GetNewRoleOk returns a tuple with the NewRole field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNewRole

`func (o *MemberActionResponse) SetNewRole(v string)`

SetNewRole sets NewRole field to given value.

### HasNewRole

`func (o *MemberActionResponse) HasNewRole() bool`

HasNewRole returns a boolean if a field has been set.

### SetNewRoleNil

`func (o *MemberActionResponse) SetNewRoleNil(b bool)`

 SetNewRoleNil sets the value for NewRole to be an explicit nil

### UnsetNewRole
`func (o *MemberActionResponse) UnsetNewRole()`

UnsetNewRole ensures that no value is present for NewRole, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


