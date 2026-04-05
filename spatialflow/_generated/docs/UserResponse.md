# UserResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** |  | 
**Email** | **string** |  | 
**Name** | **string** |  | 
**Role** | **string** |  | 
**EmailVerified** | **bool** |  | 
**SelectedPlan** | **string** |  | 
**Company** | **NullableString** |  | 
**LanguagePreference** | **string** |  | 
**CreatedAt** | **string** |  | 
**UpdatedAt** | **string** |  | 
**IsSuperuser** | **bool** |  | 
**IsStaff** | **bool** |  | 
**Roles** | Pointer to **[]string** |  | [optional] 
**IsBetaUser** | Pointer to **NullableBool** |  | [optional] 
**AdminApproved** | Pointer to **NullableBool** |  | [optional] 
**AdminApprovedAt** | Pointer to **NullableString** |  | [optional] 
**WorkspaceId** | Pointer to **NullableString** |  | [optional] 
**WorkspaceSlug** | Pointer to **NullableString** |  | [optional] 
**WorkspaceRole** | Pointer to **NullableString** |  | [optional] 

## Methods

### NewUserResponse

`func NewUserResponse(id string, email string, name string, role string, emailVerified bool, selectedPlan string, company NullableString, languagePreference string, createdAt string, updatedAt string, isSuperuser bool, isStaff bool, ) *UserResponse`

NewUserResponse instantiates a new UserResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUserResponseWithDefaults

`func NewUserResponseWithDefaults() *UserResponse`

NewUserResponseWithDefaults instantiates a new UserResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *UserResponse) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *UserResponse) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *UserResponse) SetId(v string)`

SetId sets Id field to given value.


### GetEmail

`func (o *UserResponse) GetEmail() string`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *UserResponse) GetEmailOk() (*string, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *UserResponse) SetEmail(v string)`

SetEmail sets Email field to given value.


### GetName

`func (o *UserResponse) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *UserResponse) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *UserResponse) SetName(v string)`

SetName sets Name field to given value.


### GetRole

`func (o *UserResponse) GetRole() string`

GetRole returns the Role field if non-nil, zero value otherwise.

### GetRoleOk

`func (o *UserResponse) GetRoleOk() (*string, bool)`

GetRoleOk returns a tuple with the Role field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRole

`func (o *UserResponse) SetRole(v string)`

SetRole sets Role field to given value.


### GetEmailVerified

`func (o *UserResponse) GetEmailVerified() bool`

GetEmailVerified returns the EmailVerified field if non-nil, zero value otherwise.

### GetEmailVerifiedOk

`func (o *UserResponse) GetEmailVerifiedOk() (*bool, bool)`

GetEmailVerifiedOk returns a tuple with the EmailVerified field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmailVerified

`func (o *UserResponse) SetEmailVerified(v bool)`

SetEmailVerified sets EmailVerified field to given value.


### GetSelectedPlan

`func (o *UserResponse) GetSelectedPlan() string`

GetSelectedPlan returns the SelectedPlan field if non-nil, zero value otherwise.

### GetSelectedPlanOk

`func (o *UserResponse) GetSelectedPlanOk() (*string, bool)`

GetSelectedPlanOk returns a tuple with the SelectedPlan field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSelectedPlan

`func (o *UserResponse) SetSelectedPlan(v string)`

SetSelectedPlan sets SelectedPlan field to given value.


### GetCompany

`func (o *UserResponse) GetCompany() string`

GetCompany returns the Company field if non-nil, zero value otherwise.

### GetCompanyOk

`func (o *UserResponse) GetCompanyOk() (*string, bool)`

GetCompanyOk returns a tuple with the Company field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompany

`func (o *UserResponse) SetCompany(v string)`

SetCompany sets Company field to given value.


### SetCompanyNil

`func (o *UserResponse) SetCompanyNil(b bool)`

 SetCompanyNil sets the value for Company to be an explicit nil

### UnsetCompany
`func (o *UserResponse) UnsetCompany()`

UnsetCompany ensures that no value is present for Company, not even an explicit nil
### GetLanguagePreference

`func (o *UserResponse) GetLanguagePreference() string`

GetLanguagePreference returns the LanguagePreference field if non-nil, zero value otherwise.

### GetLanguagePreferenceOk

`func (o *UserResponse) GetLanguagePreferenceOk() (*string, bool)`

GetLanguagePreferenceOk returns a tuple with the LanguagePreference field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLanguagePreference

`func (o *UserResponse) SetLanguagePreference(v string)`

SetLanguagePreference sets LanguagePreference field to given value.


### GetCreatedAt

`func (o *UserResponse) GetCreatedAt() string`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *UserResponse) GetCreatedAtOk() (*string, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *UserResponse) SetCreatedAt(v string)`

SetCreatedAt sets CreatedAt field to given value.


### GetUpdatedAt

`func (o *UserResponse) GetUpdatedAt() string`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *UserResponse) GetUpdatedAtOk() (*string, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *UserResponse) SetUpdatedAt(v string)`

SetUpdatedAt sets UpdatedAt field to given value.


### GetIsSuperuser

`func (o *UserResponse) GetIsSuperuser() bool`

GetIsSuperuser returns the IsSuperuser field if non-nil, zero value otherwise.

### GetIsSuperuserOk

`func (o *UserResponse) GetIsSuperuserOk() (*bool, bool)`

GetIsSuperuserOk returns a tuple with the IsSuperuser field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsSuperuser

`func (o *UserResponse) SetIsSuperuser(v bool)`

SetIsSuperuser sets IsSuperuser field to given value.


### GetIsStaff

`func (o *UserResponse) GetIsStaff() bool`

GetIsStaff returns the IsStaff field if non-nil, zero value otherwise.

### GetIsStaffOk

`func (o *UserResponse) GetIsStaffOk() (*bool, bool)`

GetIsStaffOk returns a tuple with the IsStaff field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsStaff

`func (o *UserResponse) SetIsStaff(v bool)`

SetIsStaff sets IsStaff field to given value.


### GetRoles

`func (o *UserResponse) GetRoles() []string`

GetRoles returns the Roles field if non-nil, zero value otherwise.

### GetRolesOk

`func (o *UserResponse) GetRolesOk() (*[]string, bool)`

GetRolesOk returns a tuple with the Roles field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRoles

`func (o *UserResponse) SetRoles(v []string)`

SetRoles sets Roles field to given value.

### HasRoles

`func (o *UserResponse) HasRoles() bool`

HasRoles returns a boolean if a field has been set.

### SetRolesNil

`func (o *UserResponse) SetRolesNil(b bool)`

 SetRolesNil sets the value for Roles to be an explicit nil

### UnsetRoles
`func (o *UserResponse) UnsetRoles()`

UnsetRoles ensures that no value is present for Roles, not even an explicit nil
### GetIsBetaUser

`func (o *UserResponse) GetIsBetaUser() bool`

GetIsBetaUser returns the IsBetaUser field if non-nil, zero value otherwise.

### GetIsBetaUserOk

`func (o *UserResponse) GetIsBetaUserOk() (*bool, bool)`

GetIsBetaUserOk returns a tuple with the IsBetaUser field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsBetaUser

`func (o *UserResponse) SetIsBetaUser(v bool)`

SetIsBetaUser sets IsBetaUser field to given value.

### HasIsBetaUser

`func (o *UserResponse) HasIsBetaUser() bool`

HasIsBetaUser returns a boolean if a field has been set.

### SetIsBetaUserNil

`func (o *UserResponse) SetIsBetaUserNil(b bool)`

 SetIsBetaUserNil sets the value for IsBetaUser to be an explicit nil

### UnsetIsBetaUser
`func (o *UserResponse) UnsetIsBetaUser()`

UnsetIsBetaUser ensures that no value is present for IsBetaUser, not even an explicit nil
### GetAdminApproved

`func (o *UserResponse) GetAdminApproved() bool`

GetAdminApproved returns the AdminApproved field if non-nil, zero value otherwise.

### GetAdminApprovedOk

`func (o *UserResponse) GetAdminApprovedOk() (*bool, bool)`

GetAdminApprovedOk returns a tuple with the AdminApproved field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdminApproved

`func (o *UserResponse) SetAdminApproved(v bool)`

SetAdminApproved sets AdminApproved field to given value.

### HasAdminApproved

`func (o *UserResponse) HasAdminApproved() bool`

HasAdminApproved returns a boolean if a field has been set.

### SetAdminApprovedNil

`func (o *UserResponse) SetAdminApprovedNil(b bool)`

 SetAdminApprovedNil sets the value for AdminApproved to be an explicit nil

### UnsetAdminApproved
`func (o *UserResponse) UnsetAdminApproved()`

UnsetAdminApproved ensures that no value is present for AdminApproved, not even an explicit nil
### GetAdminApprovedAt

`func (o *UserResponse) GetAdminApprovedAt() string`

GetAdminApprovedAt returns the AdminApprovedAt field if non-nil, zero value otherwise.

### GetAdminApprovedAtOk

`func (o *UserResponse) GetAdminApprovedAtOk() (*string, bool)`

GetAdminApprovedAtOk returns a tuple with the AdminApprovedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdminApprovedAt

`func (o *UserResponse) SetAdminApprovedAt(v string)`

SetAdminApprovedAt sets AdminApprovedAt field to given value.

### HasAdminApprovedAt

`func (o *UserResponse) HasAdminApprovedAt() bool`

HasAdminApprovedAt returns a boolean if a field has been set.

### SetAdminApprovedAtNil

`func (o *UserResponse) SetAdminApprovedAtNil(b bool)`

 SetAdminApprovedAtNil sets the value for AdminApprovedAt to be an explicit nil

### UnsetAdminApprovedAt
`func (o *UserResponse) UnsetAdminApprovedAt()`

UnsetAdminApprovedAt ensures that no value is present for AdminApprovedAt, not even an explicit nil
### GetWorkspaceId

`func (o *UserResponse) GetWorkspaceId() string`

GetWorkspaceId returns the WorkspaceId field if non-nil, zero value otherwise.

### GetWorkspaceIdOk

`func (o *UserResponse) GetWorkspaceIdOk() (*string, bool)`

GetWorkspaceIdOk returns a tuple with the WorkspaceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWorkspaceId

`func (o *UserResponse) SetWorkspaceId(v string)`

SetWorkspaceId sets WorkspaceId field to given value.

### HasWorkspaceId

`func (o *UserResponse) HasWorkspaceId() bool`

HasWorkspaceId returns a boolean if a field has been set.

### SetWorkspaceIdNil

`func (o *UserResponse) SetWorkspaceIdNil(b bool)`

 SetWorkspaceIdNil sets the value for WorkspaceId to be an explicit nil

### UnsetWorkspaceId
`func (o *UserResponse) UnsetWorkspaceId()`

UnsetWorkspaceId ensures that no value is present for WorkspaceId, not even an explicit nil
### GetWorkspaceSlug

`func (o *UserResponse) GetWorkspaceSlug() string`

GetWorkspaceSlug returns the WorkspaceSlug field if non-nil, zero value otherwise.

### GetWorkspaceSlugOk

`func (o *UserResponse) GetWorkspaceSlugOk() (*string, bool)`

GetWorkspaceSlugOk returns a tuple with the WorkspaceSlug field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWorkspaceSlug

`func (o *UserResponse) SetWorkspaceSlug(v string)`

SetWorkspaceSlug sets WorkspaceSlug field to given value.

### HasWorkspaceSlug

`func (o *UserResponse) HasWorkspaceSlug() bool`

HasWorkspaceSlug returns a boolean if a field has been set.

### SetWorkspaceSlugNil

`func (o *UserResponse) SetWorkspaceSlugNil(b bool)`

 SetWorkspaceSlugNil sets the value for WorkspaceSlug to be an explicit nil

### UnsetWorkspaceSlug
`func (o *UserResponse) UnsetWorkspaceSlug()`

UnsetWorkspaceSlug ensures that no value is present for WorkspaceSlug, not even an explicit nil
### GetWorkspaceRole

`func (o *UserResponse) GetWorkspaceRole() string`

GetWorkspaceRole returns the WorkspaceRole field if non-nil, zero value otherwise.

### GetWorkspaceRoleOk

`func (o *UserResponse) GetWorkspaceRoleOk() (*string, bool)`

GetWorkspaceRoleOk returns a tuple with the WorkspaceRole field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWorkspaceRole

`func (o *UserResponse) SetWorkspaceRole(v string)`

SetWorkspaceRole sets WorkspaceRole field to given value.

### HasWorkspaceRole

`func (o *UserResponse) HasWorkspaceRole() bool`

HasWorkspaceRole returns a boolean if a field has been set.

### SetWorkspaceRoleNil

`func (o *UserResponse) SetWorkspaceRoleNil(b bool)`

 SetWorkspaceRoleNil sets the value for WorkspaceRole to be an explicit nil

### UnsetWorkspaceRole
`func (o *UserResponse) UnsetWorkspaceRole()`

UnsetWorkspaceRole ensures that no value is present for WorkspaceRole, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


