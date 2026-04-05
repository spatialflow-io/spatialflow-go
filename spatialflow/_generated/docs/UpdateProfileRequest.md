# UpdateProfileRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **NullableString** |  | [optional] 
**Bio** | Pointer to **NullableString** |  | [optional] 
**AvatarUrl** | Pointer to **NullableString** |  | [optional] 
**PhoneNumber** | Pointer to **NullableString** |  | [optional] 
**Timezone** | Pointer to **NullableString** |  | [optional] 
**DateFormat** | Pointer to **NullableString** |  | [optional] 
**TimeFormat** | Pointer to **NullableString** |  | [optional] 
**EmailNotifications** | Pointer to **NullableBool** |  | [optional] 
**WebhookFailureNotifications** | Pointer to **NullableBool** |  | [optional] 
**WorkflowFailureNotifications** | Pointer to **NullableBool** |  | [optional] 
**UsageAlertNotifications** | Pointer to **NullableBool** |  | [optional] 
**MarketingEmails** | Pointer to **NullableBool** |  | [optional] 
**DefaultMapStyle** | Pointer to **NullableString** |  | [optional] 
**DefaultGeofenceColor** | Pointer to **NullableString** |  | [optional] 
**ShowTutorialTooltips** | Pointer to **NullableBool** |  | [optional] 
**DefaultApiVersion** | Pointer to **NullableString** |  | [optional] 

## Methods

### NewUpdateProfileRequest

`func NewUpdateProfileRequest() *UpdateProfileRequest`

NewUpdateProfileRequest instantiates a new UpdateProfileRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateProfileRequestWithDefaults

`func NewUpdateProfileRequestWithDefaults() *UpdateProfileRequest`

NewUpdateProfileRequestWithDefaults instantiates a new UpdateProfileRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *UpdateProfileRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *UpdateProfileRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *UpdateProfileRequest) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *UpdateProfileRequest) HasName() bool`

HasName returns a boolean if a field has been set.

### SetNameNil

`func (o *UpdateProfileRequest) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *UpdateProfileRequest) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetBio

`func (o *UpdateProfileRequest) GetBio() string`

GetBio returns the Bio field if non-nil, zero value otherwise.

### GetBioOk

`func (o *UpdateProfileRequest) GetBioOk() (*string, bool)`

GetBioOk returns a tuple with the Bio field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBio

`func (o *UpdateProfileRequest) SetBio(v string)`

SetBio sets Bio field to given value.

### HasBio

`func (o *UpdateProfileRequest) HasBio() bool`

HasBio returns a boolean if a field has been set.

### SetBioNil

`func (o *UpdateProfileRequest) SetBioNil(b bool)`

 SetBioNil sets the value for Bio to be an explicit nil

### UnsetBio
`func (o *UpdateProfileRequest) UnsetBio()`

UnsetBio ensures that no value is present for Bio, not even an explicit nil
### GetAvatarUrl

`func (o *UpdateProfileRequest) GetAvatarUrl() string`

GetAvatarUrl returns the AvatarUrl field if non-nil, zero value otherwise.

### GetAvatarUrlOk

`func (o *UpdateProfileRequest) GetAvatarUrlOk() (*string, bool)`

GetAvatarUrlOk returns a tuple with the AvatarUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvatarUrl

`func (o *UpdateProfileRequest) SetAvatarUrl(v string)`

SetAvatarUrl sets AvatarUrl field to given value.

### HasAvatarUrl

`func (o *UpdateProfileRequest) HasAvatarUrl() bool`

HasAvatarUrl returns a boolean if a field has been set.

### SetAvatarUrlNil

`func (o *UpdateProfileRequest) SetAvatarUrlNil(b bool)`

 SetAvatarUrlNil sets the value for AvatarUrl to be an explicit nil

### UnsetAvatarUrl
`func (o *UpdateProfileRequest) UnsetAvatarUrl()`

UnsetAvatarUrl ensures that no value is present for AvatarUrl, not even an explicit nil
### GetPhoneNumber

`func (o *UpdateProfileRequest) GetPhoneNumber() string`

GetPhoneNumber returns the PhoneNumber field if non-nil, zero value otherwise.

### GetPhoneNumberOk

`func (o *UpdateProfileRequest) GetPhoneNumberOk() (*string, bool)`

GetPhoneNumberOk returns a tuple with the PhoneNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPhoneNumber

`func (o *UpdateProfileRequest) SetPhoneNumber(v string)`

SetPhoneNumber sets PhoneNumber field to given value.

### HasPhoneNumber

`func (o *UpdateProfileRequest) HasPhoneNumber() bool`

HasPhoneNumber returns a boolean if a field has been set.

### SetPhoneNumberNil

`func (o *UpdateProfileRequest) SetPhoneNumberNil(b bool)`

 SetPhoneNumberNil sets the value for PhoneNumber to be an explicit nil

### UnsetPhoneNumber
`func (o *UpdateProfileRequest) UnsetPhoneNumber()`

UnsetPhoneNumber ensures that no value is present for PhoneNumber, not even an explicit nil
### GetTimezone

`func (o *UpdateProfileRequest) GetTimezone() string`

GetTimezone returns the Timezone field if non-nil, zero value otherwise.

### GetTimezoneOk

`func (o *UpdateProfileRequest) GetTimezoneOk() (*string, bool)`

GetTimezoneOk returns a tuple with the Timezone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimezone

`func (o *UpdateProfileRequest) SetTimezone(v string)`

SetTimezone sets Timezone field to given value.

### HasTimezone

`func (o *UpdateProfileRequest) HasTimezone() bool`

HasTimezone returns a boolean if a field has been set.

### SetTimezoneNil

`func (o *UpdateProfileRequest) SetTimezoneNil(b bool)`

 SetTimezoneNil sets the value for Timezone to be an explicit nil

### UnsetTimezone
`func (o *UpdateProfileRequest) UnsetTimezone()`

UnsetTimezone ensures that no value is present for Timezone, not even an explicit nil
### GetDateFormat

`func (o *UpdateProfileRequest) GetDateFormat() string`

GetDateFormat returns the DateFormat field if non-nil, zero value otherwise.

### GetDateFormatOk

`func (o *UpdateProfileRequest) GetDateFormatOk() (*string, bool)`

GetDateFormatOk returns a tuple with the DateFormat field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDateFormat

`func (o *UpdateProfileRequest) SetDateFormat(v string)`

SetDateFormat sets DateFormat field to given value.

### HasDateFormat

`func (o *UpdateProfileRequest) HasDateFormat() bool`

HasDateFormat returns a boolean if a field has been set.

### SetDateFormatNil

`func (o *UpdateProfileRequest) SetDateFormatNil(b bool)`

 SetDateFormatNil sets the value for DateFormat to be an explicit nil

### UnsetDateFormat
`func (o *UpdateProfileRequest) UnsetDateFormat()`

UnsetDateFormat ensures that no value is present for DateFormat, not even an explicit nil
### GetTimeFormat

`func (o *UpdateProfileRequest) GetTimeFormat() string`

GetTimeFormat returns the TimeFormat field if non-nil, zero value otherwise.

### GetTimeFormatOk

`func (o *UpdateProfileRequest) GetTimeFormatOk() (*string, bool)`

GetTimeFormatOk returns a tuple with the TimeFormat field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeFormat

`func (o *UpdateProfileRequest) SetTimeFormat(v string)`

SetTimeFormat sets TimeFormat field to given value.

### HasTimeFormat

`func (o *UpdateProfileRequest) HasTimeFormat() bool`

HasTimeFormat returns a boolean if a field has been set.

### SetTimeFormatNil

`func (o *UpdateProfileRequest) SetTimeFormatNil(b bool)`

 SetTimeFormatNil sets the value for TimeFormat to be an explicit nil

### UnsetTimeFormat
`func (o *UpdateProfileRequest) UnsetTimeFormat()`

UnsetTimeFormat ensures that no value is present for TimeFormat, not even an explicit nil
### GetEmailNotifications

`func (o *UpdateProfileRequest) GetEmailNotifications() bool`

GetEmailNotifications returns the EmailNotifications field if non-nil, zero value otherwise.

### GetEmailNotificationsOk

`func (o *UpdateProfileRequest) GetEmailNotificationsOk() (*bool, bool)`

GetEmailNotificationsOk returns a tuple with the EmailNotifications field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmailNotifications

`func (o *UpdateProfileRequest) SetEmailNotifications(v bool)`

SetEmailNotifications sets EmailNotifications field to given value.

### HasEmailNotifications

`func (o *UpdateProfileRequest) HasEmailNotifications() bool`

HasEmailNotifications returns a boolean if a field has been set.

### SetEmailNotificationsNil

`func (o *UpdateProfileRequest) SetEmailNotificationsNil(b bool)`

 SetEmailNotificationsNil sets the value for EmailNotifications to be an explicit nil

### UnsetEmailNotifications
`func (o *UpdateProfileRequest) UnsetEmailNotifications()`

UnsetEmailNotifications ensures that no value is present for EmailNotifications, not even an explicit nil
### GetWebhookFailureNotifications

`func (o *UpdateProfileRequest) GetWebhookFailureNotifications() bool`

GetWebhookFailureNotifications returns the WebhookFailureNotifications field if non-nil, zero value otherwise.

### GetWebhookFailureNotificationsOk

`func (o *UpdateProfileRequest) GetWebhookFailureNotificationsOk() (*bool, bool)`

GetWebhookFailureNotificationsOk returns a tuple with the WebhookFailureNotifications field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWebhookFailureNotifications

`func (o *UpdateProfileRequest) SetWebhookFailureNotifications(v bool)`

SetWebhookFailureNotifications sets WebhookFailureNotifications field to given value.

### HasWebhookFailureNotifications

`func (o *UpdateProfileRequest) HasWebhookFailureNotifications() bool`

HasWebhookFailureNotifications returns a boolean if a field has been set.

### SetWebhookFailureNotificationsNil

`func (o *UpdateProfileRequest) SetWebhookFailureNotificationsNil(b bool)`

 SetWebhookFailureNotificationsNil sets the value for WebhookFailureNotifications to be an explicit nil

### UnsetWebhookFailureNotifications
`func (o *UpdateProfileRequest) UnsetWebhookFailureNotifications()`

UnsetWebhookFailureNotifications ensures that no value is present for WebhookFailureNotifications, not even an explicit nil
### GetWorkflowFailureNotifications

`func (o *UpdateProfileRequest) GetWorkflowFailureNotifications() bool`

GetWorkflowFailureNotifications returns the WorkflowFailureNotifications field if non-nil, zero value otherwise.

### GetWorkflowFailureNotificationsOk

`func (o *UpdateProfileRequest) GetWorkflowFailureNotificationsOk() (*bool, bool)`

GetWorkflowFailureNotificationsOk returns a tuple with the WorkflowFailureNotifications field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWorkflowFailureNotifications

`func (o *UpdateProfileRequest) SetWorkflowFailureNotifications(v bool)`

SetWorkflowFailureNotifications sets WorkflowFailureNotifications field to given value.

### HasWorkflowFailureNotifications

`func (o *UpdateProfileRequest) HasWorkflowFailureNotifications() bool`

HasWorkflowFailureNotifications returns a boolean if a field has been set.

### SetWorkflowFailureNotificationsNil

`func (o *UpdateProfileRequest) SetWorkflowFailureNotificationsNil(b bool)`

 SetWorkflowFailureNotificationsNil sets the value for WorkflowFailureNotifications to be an explicit nil

### UnsetWorkflowFailureNotifications
`func (o *UpdateProfileRequest) UnsetWorkflowFailureNotifications()`

UnsetWorkflowFailureNotifications ensures that no value is present for WorkflowFailureNotifications, not even an explicit nil
### GetUsageAlertNotifications

`func (o *UpdateProfileRequest) GetUsageAlertNotifications() bool`

GetUsageAlertNotifications returns the UsageAlertNotifications field if non-nil, zero value otherwise.

### GetUsageAlertNotificationsOk

`func (o *UpdateProfileRequest) GetUsageAlertNotificationsOk() (*bool, bool)`

GetUsageAlertNotificationsOk returns a tuple with the UsageAlertNotifications field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsageAlertNotifications

`func (o *UpdateProfileRequest) SetUsageAlertNotifications(v bool)`

SetUsageAlertNotifications sets UsageAlertNotifications field to given value.

### HasUsageAlertNotifications

`func (o *UpdateProfileRequest) HasUsageAlertNotifications() bool`

HasUsageAlertNotifications returns a boolean if a field has been set.

### SetUsageAlertNotificationsNil

`func (o *UpdateProfileRequest) SetUsageAlertNotificationsNil(b bool)`

 SetUsageAlertNotificationsNil sets the value for UsageAlertNotifications to be an explicit nil

### UnsetUsageAlertNotifications
`func (o *UpdateProfileRequest) UnsetUsageAlertNotifications()`

UnsetUsageAlertNotifications ensures that no value is present for UsageAlertNotifications, not even an explicit nil
### GetMarketingEmails

`func (o *UpdateProfileRequest) GetMarketingEmails() bool`

GetMarketingEmails returns the MarketingEmails field if non-nil, zero value otherwise.

### GetMarketingEmailsOk

`func (o *UpdateProfileRequest) GetMarketingEmailsOk() (*bool, bool)`

GetMarketingEmailsOk returns a tuple with the MarketingEmails field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMarketingEmails

`func (o *UpdateProfileRequest) SetMarketingEmails(v bool)`

SetMarketingEmails sets MarketingEmails field to given value.

### HasMarketingEmails

`func (o *UpdateProfileRequest) HasMarketingEmails() bool`

HasMarketingEmails returns a boolean if a field has been set.

### SetMarketingEmailsNil

`func (o *UpdateProfileRequest) SetMarketingEmailsNil(b bool)`

 SetMarketingEmailsNil sets the value for MarketingEmails to be an explicit nil

### UnsetMarketingEmails
`func (o *UpdateProfileRequest) UnsetMarketingEmails()`

UnsetMarketingEmails ensures that no value is present for MarketingEmails, not even an explicit nil
### GetDefaultMapStyle

`func (o *UpdateProfileRequest) GetDefaultMapStyle() string`

GetDefaultMapStyle returns the DefaultMapStyle field if non-nil, zero value otherwise.

### GetDefaultMapStyleOk

`func (o *UpdateProfileRequest) GetDefaultMapStyleOk() (*string, bool)`

GetDefaultMapStyleOk returns a tuple with the DefaultMapStyle field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultMapStyle

`func (o *UpdateProfileRequest) SetDefaultMapStyle(v string)`

SetDefaultMapStyle sets DefaultMapStyle field to given value.

### HasDefaultMapStyle

`func (o *UpdateProfileRequest) HasDefaultMapStyle() bool`

HasDefaultMapStyle returns a boolean if a field has been set.

### SetDefaultMapStyleNil

`func (o *UpdateProfileRequest) SetDefaultMapStyleNil(b bool)`

 SetDefaultMapStyleNil sets the value for DefaultMapStyle to be an explicit nil

### UnsetDefaultMapStyle
`func (o *UpdateProfileRequest) UnsetDefaultMapStyle()`

UnsetDefaultMapStyle ensures that no value is present for DefaultMapStyle, not even an explicit nil
### GetDefaultGeofenceColor

`func (o *UpdateProfileRequest) GetDefaultGeofenceColor() string`

GetDefaultGeofenceColor returns the DefaultGeofenceColor field if non-nil, zero value otherwise.

### GetDefaultGeofenceColorOk

`func (o *UpdateProfileRequest) GetDefaultGeofenceColorOk() (*string, bool)`

GetDefaultGeofenceColorOk returns a tuple with the DefaultGeofenceColor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultGeofenceColor

`func (o *UpdateProfileRequest) SetDefaultGeofenceColor(v string)`

SetDefaultGeofenceColor sets DefaultGeofenceColor field to given value.

### HasDefaultGeofenceColor

`func (o *UpdateProfileRequest) HasDefaultGeofenceColor() bool`

HasDefaultGeofenceColor returns a boolean if a field has been set.

### SetDefaultGeofenceColorNil

`func (o *UpdateProfileRequest) SetDefaultGeofenceColorNil(b bool)`

 SetDefaultGeofenceColorNil sets the value for DefaultGeofenceColor to be an explicit nil

### UnsetDefaultGeofenceColor
`func (o *UpdateProfileRequest) UnsetDefaultGeofenceColor()`

UnsetDefaultGeofenceColor ensures that no value is present for DefaultGeofenceColor, not even an explicit nil
### GetShowTutorialTooltips

`func (o *UpdateProfileRequest) GetShowTutorialTooltips() bool`

GetShowTutorialTooltips returns the ShowTutorialTooltips field if non-nil, zero value otherwise.

### GetShowTutorialTooltipsOk

`func (o *UpdateProfileRequest) GetShowTutorialTooltipsOk() (*bool, bool)`

GetShowTutorialTooltipsOk returns a tuple with the ShowTutorialTooltips field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShowTutorialTooltips

`func (o *UpdateProfileRequest) SetShowTutorialTooltips(v bool)`

SetShowTutorialTooltips sets ShowTutorialTooltips field to given value.

### HasShowTutorialTooltips

`func (o *UpdateProfileRequest) HasShowTutorialTooltips() bool`

HasShowTutorialTooltips returns a boolean if a field has been set.

### SetShowTutorialTooltipsNil

`func (o *UpdateProfileRequest) SetShowTutorialTooltipsNil(b bool)`

 SetShowTutorialTooltipsNil sets the value for ShowTutorialTooltips to be an explicit nil

### UnsetShowTutorialTooltips
`func (o *UpdateProfileRequest) UnsetShowTutorialTooltips()`

UnsetShowTutorialTooltips ensures that no value is present for ShowTutorialTooltips, not even an explicit nil
### GetDefaultApiVersion

`func (o *UpdateProfileRequest) GetDefaultApiVersion() string`

GetDefaultApiVersion returns the DefaultApiVersion field if non-nil, zero value otherwise.

### GetDefaultApiVersionOk

`func (o *UpdateProfileRequest) GetDefaultApiVersionOk() (*string, bool)`

GetDefaultApiVersionOk returns a tuple with the DefaultApiVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultApiVersion

`func (o *UpdateProfileRequest) SetDefaultApiVersion(v string)`

SetDefaultApiVersion sets DefaultApiVersion field to given value.

### HasDefaultApiVersion

`func (o *UpdateProfileRequest) HasDefaultApiVersion() bool`

HasDefaultApiVersion returns a boolean if a field has been set.

### SetDefaultApiVersionNil

`func (o *UpdateProfileRequest) SetDefaultApiVersionNil(b bool)`

 SetDefaultApiVersionNil sets the value for DefaultApiVersion to be an explicit nil

### UnsetDefaultApiVersion
`func (o *UpdateProfileRequest) UnsetDefaultApiVersion()`

UnsetDefaultApiVersion ensures that no value is present for DefaultApiVersion, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


