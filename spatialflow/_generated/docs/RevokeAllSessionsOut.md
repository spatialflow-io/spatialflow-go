# RevokeAllSessionsOut

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Message** | **string** |  | 
**MembersAffected** | **int32** |  | 
**SessionsRevoked** | **int32** |  | 
**WorkspaceId** | **string** |  | 

## Methods

### NewRevokeAllSessionsOut

`func NewRevokeAllSessionsOut(message string, membersAffected int32, sessionsRevoked int32, workspaceId string, ) *RevokeAllSessionsOut`

NewRevokeAllSessionsOut instantiates a new RevokeAllSessionsOut object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRevokeAllSessionsOutWithDefaults

`func NewRevokeAllSessionsOutWithDefaults() *RevokeAllSessionsOut`

NewRevokeAllSessionsOutWithDefaults instantiates a new RevokeAllSessionsOut object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMessage

`func (o *RevokeAllSessionsOut) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *RevokeAllSessionsOut) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *RevokeAllSessionsOut) SetMessage(v string)`

SetMessage sets Message field to given value.


### GetMembersAffected

`func (o *RevokeAllSessionsOut) GetMembersAffected() int32`

GetMembersAffected returns the MembersAffected field if non-nil, zero value otherwise.

### GetMembersAffectedOk

`func (o *RevokeAllSessionsOut) GetMembersAffectedOk() (*int32, bool)`

GetMembersAffectedOk returns a tuple with the MembersAffected field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMembersAffected

`func (o *RevokeAllSessionsOut) SetMembersAffected(v int32)`

SetMembersAffected sets MembersAffected field to given value.


### GetSessionsRevoked

`func (o *RevokeAllSessionsOut) GetSessionsRevoked() int32`

GetSessionsRevoked returns the SessionsRevoked field if non-nil, zero value otherwise.

### GetSessionsRevokedOk

`func (o *RevokeAllSessionsOut) GetSessionsRevokedOk() (*int32, bool)`

GetSessionsRevokedOk returns a tuple with the SessionsRevoked field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSessionsRevoked

`func (o *RevokeAllSessionsOut) SetSessionsRevoked(v int32)`

SetSessionsRevoked sets SessionsRevoked field to given value.


### GetWorkspaceId

`func (o *RevokeAllSessionsOut) GetWorkspaceId() string`

GetWorkspaceId returns the WorkspaceId field if non-nil, zero value otherwise.

### GetWorkspaceIdOk

`func (o *RevokeAllSessionsOut) GetWorkspaceIdOk() (*string, bool)`

GetWorkspaceIdOk returns a tuple with the WorkspaceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWorkspaceId

`func (o *RevokeAllSessionsOut) SetWorkspaceId(v string)`

SetWorkspaceId sets WorkspaceId field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


