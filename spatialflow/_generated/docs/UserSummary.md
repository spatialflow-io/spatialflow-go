# UserSummary

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** |  | 
**Email** | **string** |  | 
**Name** | **NullableString** |  | 
**Role** | **string** |  | 
**EmailVerified** | **bool** |  | 
**IsBetaUser** | Pointer to **bool** |  | [optional] [default to false]
**AdminApproved** | Pointer to **bool** |  | [optional] [default to false]
**AdminApprovedAt** | Pointer to **NullableString** |  | [optional] 
**CreatedAt** | **NullableString** |  | 
**LastLogin** | **NullableString** |  | 
**UpdatedAt** | Pointer to **NullableString** |  | [optional] 
**SubscriptionTier** | Pointer to **string** |  | [optional] [default to "free"]
**ApiKeysCount** | **int32** |  | 
**Workspace** | Pointer to [**NullableWorkspaceSummary**](WorkspaceSummary.md) |  | [optional] 

## Methods

### NewUserSummary

`func NewUserSummary(id string, email string, name NullableString, role string, emailVerified bool, createdAt NullableString, lastLogin NullableString, apiKeysCount int32, ) *UserSummary`

NewUserSummary instantiates a new UserSummary object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUserSummaryWithDefaults

`func NewUserSummaryWithDefaults() *UserSummary`

NewUserSummaryWithDefaults instantiates a new UserSummary object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *UserSummary) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *UserSummary) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *UserSummary) SetId(v string)`

SetId sets Id field to given value.


### GetEmail

`func (o *UserSummary) GetEmail() string`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *UserSummary) GetEmailOk() (*string, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *UserSummary) SetEmail(v string)`

SetEmail sets Email field to given value.


### GetName

`func (o *UserSummary) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *UserSummary) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *UserSummary) SetName(v string)`

SetName sets Name field to given value.


### SetNameNil

`func (o *UserSummary) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *UserSummary) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetRole

`func (o *UserSummary) GetRole() string`

GetRole returns the Role field if non-nil, zero value otherwise.

### GetRoleOk

`func (o *UserSummary) GetRoleOk() (*string, bool)`

GetRoleOk returns a tuple with the Role field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRole

`func (o *UserSummary) SetRole(v string)`

SetRole sets Role field to given value.


### GetEmailVerified

`func (o *UserSummary) GetEmailVerified() bool`

GetEmailVerified returns the EmailVerified field if non-nil, zero value otherwise.

### GetEmailVerifiedOk

`func (o *UserSummary) GetEmailVerifiedOk() (*bool, bool)`

GetEmailVerifiedOk returns a tuple with the EmailVerified field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmailVerified

`func (o *UserSummary) SetEmailVerified(v bool)`

SetEmailVerified sets EmailVerified field to given value.


### GetIsBetaUser

`func (o *UserSummary) GetIsBetaUser() bool`

GetIsBetaUser returns the IsBetaUser field if non-nil, zero value otherwise.

### GetIsBetaUserOk

`func (o *UserSummary) GetIsBetaUserOk() (*bool, bool)`

GetIsBetaUserOk returns a tuple with the IsBetaUser field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsBetaUser

`func (o *UserSummary) SetIsBetaUser(v bool)`

SetIsBetaUser sets IsBetaUser field to given value.

### HasIsBetaUser

`func (o *UserSummary) HasIsBetaUser() bool`

HasIsBetaUser returns a boolean if a field has been set.

### GetAdminApproved

`func (o *UserSummary) GetAdminApproved() bool`

GetAdminApproved returns the AdminApproved field if non-nil, zero value otherwise.

### GetAdminApprovedOk

`func (o *UserSummary) GetAdminApprovedOk() (*bool, bool)`

GetAdminApprovedOk returns a tuple with the AdminApproved field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdminApproved

`func (o *UserSummary) SetAdminApproved(v bool)`

SetAdminApproved sets AdminApproved field to given value.

### HasAdminApproved

`func (o *UserSummary) HasAdminApproved() bool`

HasAdminApproved returns a boolean if a field has been set.

### GetAdminApprovedAt

`func (o *UserSummary) GetAdminApprovedAt() string`

GetAdminApprovedAt returns the AdminApprovedAt field if non-nil, zero value otherwise.

### GetAdminApprovedAtOk

`func (o *UserSummary) GetAdminApprovedAtOk() (*string, bool)`

GetAdminApprovedAtOk returns a tuple with the AdminApprovedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdminApprovedAt

`func (o *UserSummary) SetAdminApprovedAt(v string)`

SetAdminApprovedAt sets AdminApprovedAt field to given value.

### HasAdminApprovedAt

`func (o *UserSummary) HasAdminApprovedAt() bool`

HasAdminApprovedAt returns a boolean if a field has been set.

### SetAdminApprovedAtNil

`func (o *UserSummary) SetAdminApprovedAtNil(b bool)`

 SetAdminApprovedAtNil sets the value for AdminApprovedAt to be an explicit nil

### UnsetAdminApprovedAt
`func (o *UserSummary) UnsetAdminApprovedAt()`

UnsetAdminApprovedAt ensures that no value is present for AdminApprovedAt, not even an explicit nil
### GetCreatedAt

`func (o *UserSummary) GetCreatedAt() string`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *UserSummary) GetCreatedAtOk() (*string, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *UserSummary) SetCreatedAt(v string)`

SetCreatedAt sets CreatedAt field to given value.


### SetCreatedAtNil

`func (o *UserSummary) SetCreatedAtNil(b bool)`

 SetCreatedAtNil sets the value for CreatedAt to be an explicit nil

### UnsetCreatedAt
`func (o *UserSummary) UnsetCreatedAt()`

UnsetCreatedAt ensures that no value is present for CreatedAt, not even an explicit nil
### GetLastLogin

`func (o *UserSummary) GetLastLogin() string`

GetLastLogin returns the LastLogin field if non-nil, zero value otherwise.

### GetLastLoginOk

`func (o *UserSummary) GetLastLoginOk() (*string, bool)`

GetLastLoginOk returns a tuple with the LastLogin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastLogin

`func (o *UserSummary) SetLastLogin(v string)`

SetLastLogin sets LastLogin field to given value.


### SetLastLoginNil

`func (o *UserSummary) SetLastLoginNil(b bool)`

 SetLastLoginNil sets the value for LastLogin to be an explicit nil

### UnsetLastLogin
`func (o *UserSummary) UnsetLastLogin()`

UnsetLastLogin ensures that no value is present for LastLogin, not even an explicit nil
### GetUpdatedAt

`func (o *UserSummary) GetUpdatedAt() string`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *UserSummary) GetUpdatedAtOk() (*string, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *UserSummary) SetUpdatedAt(v string)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *UserSummary) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.

### SetUpdatedAtNil

`func (o *UserSummary) SetUpdatedAtNil(b bool)`

 SetUpdatedAtNil sets the value for UpdatedAt to be an explicit nil

### UnsetUpdatedAt
`func (o *UserSummary) UnsetUpdatedAt()`

UnsetUpdatedAt ensures that no value is present for UpdatedAt, not even an explicit nil
### GetSubscriptionTier

`func (o *UserSummary) GetSubscriptionTier() string`

GetSubscriptionTier returns the SubscriptionTier field if non-nil, zero value otherwise.

### GetSubscriptionTierOk

`func (o *UserSummary) GetSubscriptionTierOk() (*string, bool)`

GetSubscriptionTierOk returns a tuple with the SubscriptionTier field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubscriptionTier

`func (o *UserSummary) SetSubscriptionTier(v string)`

SetSubscriptionTier sets SubscriptionTier field to given value.

### HasSubscriptionTier

`func (o *UserSummary) HasSubscriptionTier() bool`

HasSubscriptionTier returns a boolean if a field has been set.

### GetApiKeysCount

`func (o *UserSummary) GetApiKeysCount() int32`

GetApiKeysCount returns the ApiKeysCount field if non-nil, zero value otherwise.

### GetApiKeysCountOk

`func (o *UserSummary) GetApiKeysCountOk() (*int32, bool)`

GetApiKeysCountOk returns a tuple with the ApiKeysCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApiKeysCount

`func (o *UserSummary) SetApiKeysCount(v int32)`

SetApiKeysCount sets ApiKeysCount field to given value.


### GetWorkspace

`func (o *UserSummary) GetWorkspace() WorkspaceSummary`

GetWorkspace returns the Workspace field if non-nil, zero value otherwise.

### GetWorkspaceOk

`func (o *UserSummary) GetWorkspaceOk() (*WorkspaceSummary, bool)`

GetWorkspaceOk returns a tuple with the Workspace field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWorkspace

`func (o *UserSummary) SetWorkspace(v WorkspaceSummary)`

SetWorkspace sets Workspace field to given value.

### HasWorkspace

`func (o *UserSummary) HasWorkspace() bool`

HasWorkspace returns a boolean if a field has been set.

### SetWorkspaceNil

`func (o *UserSummary) SetWorkspaceNil(b bool)`

 SetWorkspaceNil sets the value for Workspace to be an explicit nil

### UnsetWorkspace
`func (o *UserSummary) UnsetWorkspace()`

UnsetWorkspace ensures that no value is present for Workspace, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


