# BulkInviteItem

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Email** | **string** |  | 
**WorkspaceId** | Pointer to **NullableString** |  | [optional] 
**Name** | Pointer to **NullableString** |  | [optional] 
**Role** | Pointer to **string** |  | [optional] [default to "field_worker"]

## Methods

### NewBulkInviteItem

`func NewBulkInviteItem(email string, ) *BulkInviteItem`

NewBulkInviteItem instantiates a new BulkInviteItem object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBulkInviteItemWithDefaults

`func NewBulkInviteItemWithDefaults() *BulkInviteItem`

NewBulkInviteItemWithDefaults instantiates a new BulkInviteItem object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEmail

`func (o *BulkInviteItem) GetEmail() string`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *BulkInviteItem) GetEmailOk() (*string, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *BulkInviteItem) SetEmail(v string)`

SetEmail sets Email field to given value.


### GetWorkspaceId

`func (o *BulkInviteItem) GetWorkspaceId() string`

GetWorkspaceId returns the WorkspaceId field if non-nil, zero value otherwise.

### GetWorkspaceIdOk

`func (o *BulkInviteItem) GetWorkspaceIdOk() (*string, bool)`

GetWorkspaceIdOk returns a tuple with the WorkspaceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWorkspaceId

`func (o *BulkInviteItem) SetWorkspaceId(v string)`

SetWorkspaceId sets WorkspaceId field to given value.

### HasWorkspaceId

`func (o *BulkInviteItem) HasWorkspaceId() bool`

HasWorkspaceId returns a boolean if a field has been set.

### SetWorkspaceIdNil

`func (o *BulkInviteItem) SetWorkspaceIdNil(b bool)`

 SetWorkspaceIdNil sets the value for WorkspaceId to be an explicit nil

### UnsetWorkspaceId
`func (o *BulkInviteItem) UnsetWorkspaceId()`

UnsetWorkspaceId ensures that no value is present for WorkspaceId, not even an explicit nil
### GetName

`func (o *BulkInviteItem) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *BulkInviteItem) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *BulkInviteItem) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *BulkInviteItem) HasName() bool`

HasName returns a boolean if a field has been set.

### SetNameNil

`func (o *BulkInviteItem) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *BulkInviteItem) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetRole

`func (o *BulkInviteItem) GetRole() string`

GetRole returns the Role field if non-nil, zero value otherwise.

### GetRoleOk

`func (o *BulkInviteItem) GetRoleOk() (*string, bool)`

GetRoleOk returns a tuple with the Role field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRole

`func (o *BulkInviteItem) SetRole(v string)`

SetRole sets Role field to given value.

### HasRole

`func (o *BulkInviteItem) HasRole() bool`

HasRole returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


