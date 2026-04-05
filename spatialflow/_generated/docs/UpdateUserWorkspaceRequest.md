# UpdateUserWorkspaceRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**WorkspaceId** | **string** |  | 
**Role** | Pointer to **string** |  | [optional] [default to "field_worker"]

## Methods

### NewUpdateUserWorkspaceRequest

`func NewUpdateUserWorkspaceRequest(workspaceId string, ) *UpdateUserWorkspaceRequest`

NewUpdateUserWorkspaceRequest instantiates a new UpdateUserWorkspaceRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateUserWorkspaceRequestWithDefaults

`func NewUpdateUserWorkspaceRequestWithDefaults() *UpdateUserWorkspaceRequest`

NewUpdateUserWorkspaceRequestWithDefaults instantiates a new UpdateUserWorkspaceRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetWorkspaceId

`func (o *UpdateUserWorkspaceRequest) GetWorkspaceId() string`

GetWorkspaceId returns the WorkspaceId field if non-nil, zero value otherwise.

### GetWorkspaceIdOk

`func (o *UpdateUserWorkspaceRequest) GetWorkspaceIdOk() (*string, bool)`

GetWorkspaceIdOk returns a tuple with the WorkspaceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWorkspaceId

`func (o *UpdateUserWorkspaceRequest) SetWorkspaceId(v string)`

SetWorkspaceId sets WorkspaceId field to given value.


### GetRole

`func (o *UpdateUserWorkspaceRequest) GetRole() string`

GetRole returns the Role field if non-nil, zero value otherwise.

### GetRoleOk

`func (o *UpdateUserWorkspaceRequest) GetRoleOk() (*string, bool)`

GetRoleOk returns a tuple with the Role field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRole

`func (o *UpdateUserWorkspaceRequest) SetRole(v string)`

SetRole sets Role field to given value.

### HasRole

`func (o *UpdateUserWorkspaceRequest) HasRole() bool`

HasRole returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


