# UserProfileResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** |  | 
**Email** | **string** |  | 
**Name** | **string** |  | 
**Role** | **string** |  | 
**IsStaff** | Pointer to **bool** |  | [optional] [default to false]
**IsSuperuser** | Pointer to **bool** |  | [optional] [default to false]
**EmailVerified** | **bool** |  | 
**CreatedAt** | **time.Time** |  | 
**UpdatedAt** | **time.Time** |  | 
**Bio** | Pointer to **NullableString** |  | [optional] 
**AvatarUrl** | Pointer to **NullableString** |  | [optional] 
**PhoneNumber** | Pointer to **NullableString** |  | [optional] 
**Timezone** | Pointer to **string** |  | [optional] [default to "UTC"]
**DateFormat** | Pointer to **string** |  | [optional] [default to "YYYY-MM-DD"]
**TimeFormat** | Pointer to **string** |  | [optional] [default to "24h"]
**EmailNotifications** | Pointer to **bool** |  | [optional] [default to true]
**WebhookFailureNotifications** | Pointer to **bool** |  | [optional] [default to true]
**WorkflowFailureNotifications** | Pointer to **bool** |  | [optional] [default to true]
**UsageAlertNotifications** | Pointer to **bool** |  | [optional] [default to true]
**MarketingEmails** | Pointer to **bool** |  | [optional] [default to false]
**DefaultMapStyle** | Pointer to **string** |  | [optional] [default to "streets"]
**DefaultGeofenceColor** | Pointer to **string** |  | [optional] [default to "#3B82F6"]
**ShowTutorialTooltips** | Pointer to **bool** |  | [optional] [default to true]
**DefaultApiVersion** | Pointer to **string** |  | [optional] [default to "v1"]
**WorkspaceRole** | Pointer to **NullableString** |  | [optional] 
**WorkspaceId** | Pointer to **NullableString** |  | [optional] 
**SelectedPlan** | Pointer to **string** |  | [optional] [default to "free"]
**Company** | Pointer to **NullableString** |  | [optional] 

## Methods

### NewUserProfileResponse

`func NewUserProfileResponse(id string, email string, name string, role string, emailVerified bool, createdAt time.Time, updatedAt time.Time, ) *UserProfileResponse`

NewUserProfileResponse instantiates a new UserProfileResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUserProfileResponseWithDefaults

`func NewUserProfileResponseWithDefaults() *UserProfileResponse`

NewUserProfileResponseWithDefaults instantiates a new UserProfileResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *UserProfileResponse) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *UserProfileResponse) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *UserProfileResponse) SetId(v string)`

SetId sets Id field to given value.


### GetEmail

`func (o *UserProfileResponse) GetEmail() string`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *UserProfileResponse) GetEmailOk() (*string, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *UserProfileResponse) SetEmail(v string)`

SetEmail sets Email field to given value.


### GetName

`func (o *UserProfileResponse) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *UserProfileResponse) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *UserProfileResponse) SetName(v string)`

SetName sets Name field to given value.


### GetRole

`func (o *UserProfileResponse) GetRole() string`

GetRole returns the Role field if non-nil, zero value otherwise.

### GetRoleOk

`func (o *UserProfileResponse) GetRoleOk() (*string, bool)`

GetRoleOk returns a tuple with the Role field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRole

`func (o *UserProfileResponse) SetRole(v string)`

SetRole sets Role field to given value.


### GetIsStaff

`func (o *UserProfileResponse) GetIsStaff() bool`

GetIsStaff returns the IsStaff field if non-nil, zero value otherwise.

### GetIsStaffOk

`func (o *UserProfileResponse) GetIsStaffOk() (*bool, bool)`

GetIsStaffOk returns a tuple with the IsStaff field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsStaff

`func (o *UserProfileResponse) SetIsStaff(v bool)`

SetIsStaff sets IsStaff field to given value.

### HasIsStaff

`func (o *UserProfileResponse) HasIsStaff() bool`

HasIsStaff returns a boolean if a field has been set.

### GetIsSuperuser

`func (o *UserProfileResponse) GetIsSuperuser() bool`

GetIsSuperuser returns the IsSuperuser field if non-nil, zero value otherwise.

### GetIsSuperuserOk

`func (o *UserProfileResponse) GetIsSuperuserOk() (*bool, bool)`

GetIsSuperuserOk returns a tuple with the IsSuperuser field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsSuperuser

`func (o *UserProfileResponse) SetIsSuperuser(v bool)`

SetIsSuperuser sets IsSuperuser field to given value.

### HasIsSuperuser

`func (o *UserProfileResponse) HasIsSuperuser() bool`

HasIsSuperuser returns a boolean if a field has been set.

### GetEmailVerified

`func (o *UserProfileResponse) GetEmailVerified() bool`

GetEmailVerified returns the EmailVerified field if non-nil, zero value otherwise.

### GetEmailVerifiedOk

`func (o *UserProfileResponse) GetEmailVerifiedOk() (*bool, bool)`

GetEmailVerifiedOk returns a tuple with the EmailVerified field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmailVerified

`func (o *UserProfileResponse) SetEmailVerified(v bool)`

SetEmailVerified sets EmailVerified field to given value.


### GetCreatedAt

`func (o *UserProfileResponse) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *UserProfileResponse) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *UserProfileResponse) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.


### GetUpdatedAt

`func (o *UserProfileResponse) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *UserProfileResponse) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *UserProfileResponse) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.


### GetBio

`func (o *UserProfileResponse) GetBio() string`

GetBio returns the Bio field if non-nil, zero value otherwise.

### GetBioOk

`func (o *UserProfileResponse) GetBioOk() (*string, bool)`

GetBioOk returns a tuple with the Bio field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBio

`func (o *UserProfileResponse) SetBio(v string)`

SetBio sets Bio field to given value.

### HasBio

`func (o *UserProfileResponse) HasBio() bool`

HasBio returns a boolean if a field has been set.

### SetBioNil

`func (o *UserProfileResponse) SetBioNil(b bool)`

 SetBioNil sets the value for Bio to be an explicit nil

### UnsetBio
`func (o *UserProfileResponse) UnsetBio()`

UnsetBio ensures that no value is present for Bio, not even an explicit nil
### GetAvatarUrl

`func (o *UserProfileResponse) GetAvatarUrl() string`

GetAvatarUrl returns the AvatarUrl field if non-nil, zero value otherwise.

### GetAvatarUrlOk

`func (o *UserProfileResponse) GetAvatarUrlOk() (*string, bool)`

GetAvatarUrlOk returns a tuple with the AvatarUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvatarUrl

`func (o *UserProfileResponse) SetAvatarUrl(v string)`

SetAvatarUrl sets AvatarUrl field to given value.

### HasAvatarUrl

`func (o *UserProfileResponse) HasAvatarUrl() bool`

HasAvatarUrl returns a boolean if a field has been set.

### SetAvatarUrlNil

`func (o *UserProfileResponse) SetAvatarUrlNil(b bool)`

 SetAvatarUrlNil sets the value for AvatarUrl to be an explicit nil

### UnsetAvatarUrl
`func (o *UserProfileResponse) UnsetAvatarUrl()`

UnsetAvatarUrl ensures that no value is present for AvatarUrl, not even an explicit nil
### GetPhoneNumber

`func (o *UserProfileResponse) GetPhoneNumber() string`

GetPhoneNumber returns the PhoneNumber field if non-nil, zero value otherwise.

### GetPhoneNumberOk

`func (o *UserProfileResponse) GetPhoneNumberOk() (*string, bool)`

GetPhoneNumberOk returns a tuple with the PhoneNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPhoneNumber

`func (o *UserProfileResponse) SetPhoneNumber(v string)`

SetPhoneNumber sets PhoneNumber field to given value.

### HasPhoneNumber

`func (o *UserProfileResponse) HasPhoneNumber() bool`

HasPhoneNumber returns a boolean if a field has been set.

### SetPhoneNumberNil

`func (o *UserProfileResponse) SetPhoneNumberNil(b bool)`

 SetPhoneNumberNil sets the value for PhoneNumber to be an explicit nil

### UnsetPhoneNumber
`func (o *UserProfileResponse) UnsetPhoneNumber()`

UnsetPhoneNumber ensures that no value is present for PhoneNumber, not even an explicit nil
### GetTimezone

`func (o *UserProfileResponse) GetTimezone() string`

GetTimezone returns the Timezone field if non-nil, zero value otherwise.

### GetTimezoneOk

`func (o *UserProfileResponse) GetTimezoneOk() (*string, bool)`

GetTimezoneOk returns a tuple with the Timezone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimezone

`func (o *UserProfileResponse) SetTimezone(v string)`

SetTimezone sets Timezone field to given value.

### HasTimezone

`func (o *UserProfileResponse) HasTimezone() bool`

HasTimezone returns a boolean if a field has been set.

### GetDateFormat

`func (o *UserProfileResponse) GetDateFormat() string`

GetDateFormat returns the DateFormat field if non-nil, zero value otherwise.

### GetDateFormatOk

`func (o *UserProfileResponse) GetDateFormatOk() (*string, bool)`

GetDateFormatOk returns a tuple with the DateFormat field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDateFormat

`func (o *UserProfileResponse) SetDateFormat(v string)`

SetDateFormat sets DateFormat field to given value.

### HasDateFormat

`func (o *UserProfileResponse) HasDateFormat() bool`

HasDateFormat returns a boolean if a field has been set.

### GetTimeFormat

`func (o *UserProfileResponse) GetTimeFormat() string`

GetTimeFormat returns the TimeFormat field if non-nil, zero value otherwise.

### GetTimeFormatOk

`func (o *UserProfileResponse) GetTimeFormatOk() (*string, bool)`

GetTimeFormatOk returns a tuple with the TimeFormat field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeFormat

`func (o *UserProfileResponse) SetTimeFormat(v string)`

SetTimeFormat sets TimeFormat field to given value.

### HasTimeFormat

`func (o *UserProfileResponse) HasTimeFormat() bool`

HasTimeFormat returns a boolean if a field has been set.

### GetEmailNotifications

`func (o *UserProfileResponse) GetEmailNotifications() bool`

GetEmailNotifications returns the EmailNotifications field if non-nil, zero value otherwise.

### GetEmailNotificationsOk

`func (o *UserProfileResponse) GetEmailNotificationsOk() (*bool, bool)`

GetEmailNotificationsOk returns a tuple with the EmailNotifications field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmailNotifications

`func (o *UserProfileResponse) SetEmailNotifications(v bool)`

SetEmailNotifications sets EmailNotifications field to given value.

### HasEmailNotifications

`func (o *UserProfileResponse) HasEmailNotifications() bool`

HasEmailNotifications returns a boolean if a field has been set.

### GetWebhookFailureNotifications

`func (o *UserProfileResponse) GetWebhookFailureNotifications() bool`

GetWebhookFailureNotifications returns the WebhookFailureNotifications field if non-nil, zero value otherwise.

### GetWebhookFailureNotificationsOk

`func (o *UserProfileResponse) GetWebhookFailureNotificationsOk() (*bool, bool)`

GetWebhookFailureNotificationsOk returns a tuple with the WebhookFailureNotifications field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWebhookFailureNotifications

`func (o *UserProfileResponse) SetWebhookFailureNotifications(v bool)`

SetWebhookFailureNotifications sets WebhookFailureNotifications field to given value.

### HasWebhookFailureNotifications

`func (o *UserProfileResponse) HasWebhookFailureNotifications() bool`

HasWebhookFailureNotifications returns a boolean if a field has been set.

### GetWorkflowFailureNotifications

`func (o *UserProfileResponse) GetWorkflowFailureNotifications() bool`

GetWorkflowFailureNotifications returns the WorkflowFailureNotifications field if non-nil, zero value otherwise.

### GetWorkflowFailureNotificationsOk

`func (o *UserProfileResponse) GetWorkflowFailureNotificationsOk() (*bool, bool)`

GetWorkflowFailureNotificationsOk returns a tuple with the WorkflowFailureNotifications field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWorkflowFailureNotifications

`func (o *UserProfileResponse) SetWorkflowFailureNotifications(v bool)`

SetWorkflowFailureNotifications sets WorkflowFailureNotifications field to given value.

### HasWorkflowFailureNotifications

`func (o *UserProfileResponse) HasWorkflowFailureNotifications() bool`

HasWorkflowFailureNotifications returns a boolean if a field has been set.

### GetUsageAlertNotifications

`func (o *UserProfileResponse) GetUsageAlertNotifications() bool`

GetUsageAlertNotifications returns the UsageAlertNotifications field if non-nil, zero value otherwise.

### GetUsageAlertNotificationsOk

`func (o *UserProfileResponse) GetUsageAlertNotificationsOk() (*bool, bool)`

GetUsageAlertNotificationsOk returns a tuple with the UsageAlertNotifications field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsageAlertNotifications

`func (o *UserProfileResponse) SetUsageAlertNotifications(v bool)`

SetUsageAlertNotifications sets UsageAlertNotifications field to given value.

### HasUsageAlertNotifications

`func (o *UserProfileResponse) HasUsageAlertNotifications() bool`

HasUsageAlertNotifications returns a boolean if a field has been set.

### GetMarketingEmails

`func (o *UserProfileResponse) GetMarketingEmails() bool`

GetMarketingEmails returns the MarketingEmails field if non-nil, zero value otherwise.

### GetMarketingEmailsOk

`func (o *UserProfileResponse) GetMarketingEmailsOk() (*bool, bool)`

GetMarketingEmailsOk returns a tuple with the MarketingEmails field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMarketingEmails

`func (o *UserProfileResponse) SetMarketingEmails(v bool)`

SetMarketingEmails sets MarketingEmails field to given value.

### HasMarketingEmails

`func (o *UserProfileResponse) HasMarketingEmails() bool`

HasMarketingEmails returns a boolean if a field has been set.

### GetDefaultMapStyle

`func (o *UserProfileResponse) GetDefaultMapStyle() string`

GetDefaultMapStyle returns the DefaultMapStyle field if non-nil, zero value otherwise.

### GetDefaultMapStyleOk

`func (o *UserProfileResponse) GetDefaultMapStyleOk() (*string, bool)`

GetDefaultMapStyleOk returns a tuple with the DefaultMapStyle field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultMapStyle

`func (o *UserProfileResponse) SetDefaultMapStyle(v string)`

SetDefaultMapStyle sets DefaultMapStyle field to given value.

### HasDefaultMapStyle

`func (o *UserProfileResponse) HasDefaultMapStyle() bool`

HasDefaultMapStyle returns a boolean if a field has been set.

### GetDefaultGeofenceColor

`func (o *UserProfileResponse) GetDefaultGeofenceColor() string`

GetDefaultGeofenceColor returns the DefaultGeofenceColor field if non-nil, zero value otherwise.

### GetDefaultGeofenceColorOk

`func (o *UserProfileResponse) GetDefaultGeofenceColorOk() (*string, bool)`

GetDefaultGeofenceColorOk returns a tuple with the DefaultGeofenceColor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultGeofenceColor

`func (o *UserProfileResponse) SetDefaultGeofenceColor(v string)`

SetDefaultGeofenceColor sets DefaultGeofenceColor field to given value.

### HasDefaultGeofenceColor

`func (o *UserProfileResponse) HasDefaultGeofenceColor() bool`

HasDefaultGeofenceColor returns a boolean if a field has been set.

### GetShowTutorialTooltips

`func (o *UserProfileResponse) GetShowTutorialTooltips() bool`

GetShowTutorialTooltips returns the ShowTutorialTooltips field if non-nil, zero value otherwise.

### GetShowTutorialTooltipsOk

`func (o *UserProfileResponse) GetShowTutorialTooltipsOk() (*bool, bool)`

GetShowTutorialTooltipsOk returns a tuple with the ShowTutorialTooltips field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShowTutorialTooltips

`func (o *UserProfileResponse) SetShowTutorialTooltips(v bool)`

SetShowTutorialTooltips sets ShowTutorialTooltips field to given value.

### HasShowTutorialTooltips

`func (o *UserProfileResponse) HasShowTutorialTooltips() bool`

HasShowTutorialTooltips returns a boolean if a field has been set.

### GetDefaultApiVersion

`func (o *UserProfileResponse) GetDefaultApiVersion() string`

GetDefaultApiVersion returns the DefaultApiVersion field if non-nil, zero value otherwise.

### GetDefaultApiVersionOk

`func (o *UserProfileResponse) GetDefaultApiVersionOk() (*string, bool)`

GetDefaultApiVersionOk returns a tuple with the DefaultApiVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultApiVersion

`func (o *UserProfileResponse) SetDefaultApiVersion(v string)`

SetDefaultApiVersion sets DefaultApiVersion field to given value.

### HasDefaultApiVersion

`func (o *UserProfileResponse) HasDefaultApiVersion() bool`

HasDefaultApiVersion returns a boolean if a field has been set.

### GetWorkspaceRole

`func (o *UserProfileResponse) GetWorkspaceRole() string`

GetWorkspaceRole returns the WorkspaceRole field if non-nil, zero value otherwise.

### GetWorkspaceRoleOk

`func (o *UserProfileResponse) GetWorkspaceRoleOk() (*string, bool)`

GetWorkspaceRoleOk returns a tuple with the WorkspaceRole field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWorkspaceRole

`func (o *UserProfileResponse) SetWorkspaceRole(v string)`

SetWorkspaceRole sets WorkspaceRole field to given value.

### HasWorkspaceRole

`func (o *UserProfileResponse) HasWorkspaceRole() bool`

HasWorkspaceRole returns a boolean if a field has been set.

### SetWorkspaceRoleNil

`func (o *UserProfileResponse) SetWorkspaceRoleNil(b bool)`

 SetWorkspaceRoleNil sets the value for WorkspaceRole to be an explicit nil

### UnsetWorkspaceRole
`func (o *UserProfileResponse) UnsetWorkspaceRole()`

UnsetWorkspaceRole ensures that no value is present for WorkspaceRole, not even an explicit nil
### GetWorkspaceId

`func (o *UserProfileResponse) GetWorkspaceId() string`

GetWorkspaceId returns the WorkspaceId field if non-nil, zero value otherwise.

### GetWorkspaceIdOk

`func (o *UserProfileResponse) GetWorkspaceIdOk() (*string, bool)`

GetWorkspaceIdOk returns a tuple with the WorkspaceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWorkspaceId

`func (o *UserProfileResponse) SetWorkspaceId(v string)`

SetWorkspaceId sets WorkspaceId field to given value.

### HasWorkspaceId

`func (o *UserProfileResponse) HasWorkspaceId() bool`

HasWorkspaceId returns a boolean if a field has been set.

### SetWorkspaceIdNil

`func (o *UserProfileResponse) SetWorkspaceIdNil(b bool)`

 SetWorkspaceIdNil sets the value for WorkspaceId to be an explicit nil

### UnsetWorkspaceId
`func (o *UserProfileResponse) UnsetWorkspaceId()`

UnsetWorkspaceId ensures that no value is present for WorkspaceId, not even an explicit nil
### GetSelectedPlan

`func (o *UserProfileResponse) GetSelectedPlan() string`

GetSelectedPlan returns the SelectedPlan field if non-nil, zero value otherwise.

### GetSelectedPlanOk

`func (o *UserProfileResponse) GetSelectedPlanOk() (*string, bool)`

GetSelectedPlanOk returns a tuple with the SelectedPlan field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSelectedPlan

`func (o *UserProfileResponse) SetSelectedPlan(v string)`

SetSelectedPlan sets SelectedPlan field to given value.

### HasSelectedPlan

`func (o *UserProfileResponse) HasSelectedPlan() bool`

HasSelectedPlan returns a boolean if a field has been set.

### GetCompany

`func (o *UserProfileResponse) GetCompany() string`

GetCompany returns the Company field if non-nil, zero value otherwise.

### GetCompanyOk

`func (o *UserProfileResponse) GetCompanyOk() (*string, bool)`

GetCompanyOk returns a tuple with the Company field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompany

`func (o *UserProfileResponse) SetCompany(v string)`

SetCompany sets Company field to given value.

### HasCompany

`func (o *UserProfileResponse) HasCompany() bool`

HasCompany returns a boolean if a field has been set.

### SetCompanyNil

`func (o *UserProfileResponse) SetCompanyNil(b bool)`

 SetCompanyNil sets the value for Company to be an explicit nil

### UnsetCompany
`func (o *UserProfileResponse) UnsetCompany()`

UnsetCompany ensures that no value is present for Company, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


