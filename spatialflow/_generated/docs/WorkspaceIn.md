# WorkspaceIn

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** |  | 
**LogoUrl** | Pointer to **NullableString** |  | [optional] 
**Website** | Pointer to **NullableString** |  | [optional] 
**BillingEmail** | Pointer to **NullableString** |  | [optional] 
**Description** | Pointer to **NullableString** |  | [optional] 
**Timezone** | Pointer to **NullableString** |  | [optional] 
**SupportEmail** | Pointer to **NullableString** |  | [optional] 
**SlackConnectUrl** | Pointer to **NullableString** |  | [optional] 
**UnitSystem** | Pointer to **NullableString** |  | [optional] 

## Methods

### NewWorkspaceIn

`func NewWorkspaceIn(name string, ) *WorkspaceIn`

NewWorkspaceIn instantiates a new WorkspaceIn object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWorkspaceInWithDefaults

`func NewWorkspaceInWithDefaults() *WorkspaceIn`

NewWorkspaceInWithDefaults instantiates a new WorkspaceIn object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *WorkspaceIn) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *WorkspaceIn) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *WorkspaceIn) SetName(v string)`

SetName sets Name field to given value.


### GetLogoUrl

`func (o *WorkspaceIn) GetLogoUrl() string`

GetLogoUrl returns the LogoUrl field if non-nil, zero value otherwise.

### GetLogoUrlOk

`func (o *WorkspaceIn) GetLogoUrlOk() (*string, bool)`

GetLogoUrlOk returns a tuple with the LogoUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLogoUrl

`func (o *WorkspaceIn) SetLogoUrl(v string)`

SetLogoUrl sets LogoUrl field to given value.

### HasLogoUrl

`func (o *WorkspaceIn) HasLogoUrl() bool`

HasLogoUrl returns a boolean if a field has been set.

### SetLogoUrlNil

`func (o *WorkspaceIn) SetLogoUrlNil(b bool)`

 SetLogoUrlNil sets the value for LogoUrl to be an explicit nil

### UnsetLogoUrl
`func (o *WorkspaceIn) UnsetLogoUrl()`

UnsetLogoUrl ensures that no value is present for LogoUrl, not even an explicit nil
### GetWebsite

`func (o *WorkspaceIn) GetWebsite() string`

GetWebsite returns the Website field if non-nil, zero value otherwise.

### GetWebsiteOk

`func (o *WorkspaceIn) GetWebsiteOk() (*string, bool)`

GetWebsiteOk returns a tuple with the Website field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWebsite

`func (o *WorkspaceIn) SetWebsite(v string)`

SetWebsite sets Website field to given value.

### HasWebsite

`func (o *WorkspaceIn) HasWebsite() bool`

HasWebsite returns a boolean if a field has been set.

### SetWebsiteNil

`func (o *WorkspaceIn) SetWebsiteNil(b bool)`

 SetWebsiteNil sets the value for Website to be an explicit nil

### UnsetWebsite
`func (o *WorkspaceIn) UnsetWebsite()`

UnsetWebsite ensures that no value is present for Website, not even an explicit nil
### GetBillingEmail

`func (o *WorkspaceIn) GetBillingEmail() string`

GetBillingEmail returns the BillingEmail field if non-nil, zero value otherwise.

### GetBillingEmailOk

`func (o *WorkspaceIn) GetBillingEmailOk() (*string, bool)`

GetBillingEmailOk returns a tuple with the BillingEmail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBillingEmail

`func (o *WorkspaceIn) SetBillingEmail(v string)`

SetBillingEmail sets BillingEmail field to given value.

### HasBillingEmail

`func (o *WorkspaceIn) HasBillingEmail() bool`

HasBillingEmail returns a boolean if a field has been set.

### SetBillingEmailNil

`func (o *WorkspaceIn) SetBillingEmailNil(b bool)`

 SetBillingEmailNil sets the value for BillingEmail to be an explicit nil

### UnsetBillingEmail
`func (o *WorkspaceIn) UnsetBillingEmail()`

UnsetBillingEmail ensures that no value is present for BillingEmail, not even an explicit nil
### GetDescription

`func (o *WorkspaceIn) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *WorkspaceIn) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *WorkspaceIn) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *WorkspaceIn) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescriptionNil

`func (o *WorkspaceIn) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *WorkspaceIn) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetTimezone

`func (o *WorkspaceIn) GetTimezone() string`

GetTimezone returns the Timezone field if non-nil, zero value otherwise.

### GetTimezoneOk

`func (o *WorkspaceIn) GetTimezoneOk() (*string, bool)`

GetTimezoneOk returns a tuple with the Timezone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimezone

`func (o *WorkspaceIn) SetTimezone(v string)`

SetTimezone sets Timezone field to given value.

### HasTimezone

`func (o *WorkspaceIn) HasTimezone() bool`

HasTimezone returns a boolean if a field has been set.

### SetTimezoneNil

`func (o *WorkspaceIn) SetTimezoneNil(b bool)`

 SetTimezoneNil sets the value for Timezone to be an explicit nil

### UnsetTimezone
`func (o *WorkspaceIn) UnsetTimezone()`

UnsetTimezone ensures that no value is present for Timezone, not even an explicit nil
### GetSupportEmail

`func (o *WorkspaceIn) GetSupportEmail() string`

GetSupportEmail returns the SupportEmail field if non-nil, zero value otherwise.

### GetSupportEmailOk

`func (o *WorkspaceIn) GetSupportEmailOk() (*string, bool)`

GetSupportEmailOk returns a tuple with the SupportEmail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSupportEmail

`func (o *WorkspaceIn) SetSupportEmail(v string)`

SetSupportEmail sets SupportEmail field to given value.

### HasSupportEmail

`func (o *WorkspaceIn) HasSupportEmail() bool`

HasSupportEmail returns a boolean if a field has been set.

### SetSupportEmailNil

`func (o *WorkspaceIn) SetSupportEmailNil(b bool)`

 SetSupportEmailNil sets the value for SupportEmail to be an explicit nil

### UnsetSupportEmail
`func (o *WorkspaceIn) UnsetSupportEmail()`

UnsetSupportEmail ensures that no value is present for SupportEmail, not even an explicit nil
### GetSlackConnectUrl

`func (o *WorkspaceIn) GetSlackConnectUrl() string`

GetSlackConnectUrl returns the SlackConnectUrl field if non-nil, zero value otherwise.

### GetSlackConnectUrlOk

`func (o *WorkspaceIn) GetSlackConnectUrlOk() (*string, bool)`

GetSlackConnectUrlOk returns a tuple with the SlackConnectUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSlackConnectUrl

`func (o *WorkspaceIn) SetSlackConnectUrl(v string)`

SetSlackConnectUrl sets SlackConnectUrl field to given value.

### HasSlackConnectUrl

`func (o *WorkspaceIn) HasSlackConnectUrl() bool`

HasSlackConnectUrl returns a boolean if a field has been set.

### SetSlackConnectUrlNil

`func (o *WorkspaceIn) SetSlackConnectUrlNil(b bool)`

 SetSlackConnectUrlNil sets the value for SlackConnectUrl to be an explicit nil

### UnsetSlackConnectUrl
`func (o *WorkspaceIn) UnsetSlackConnectUrl()`

UnsetSlackConnectUrl ensures that no value is present for SlackConnectUrl, not even an explicit nil
### GetUnitSystem

`func (o *WorkspaceIn) GetUnitSystem() string`

GetUnitSystem returns the UnitSystem field if non-nil, zero value otherwise.

### GetUnitSystemOk

`func (o *WorkspaceIn) GetUnitSystemOk() (*string, bool)`

GetUnitSystemOk returns a tuple with the UnitSystem field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnitSystem

`func (o *WorkspaceIn) SetUnitSystem(v string)`

SetUnitSystem sets UnitSystem field to given value.

### HasUnitSystem

`func (o *WorkspaceIn) HasUnitSystem() bool`

HasUnitSystem returns a boolean if a field has been set.

### SetUnitSystemNil

`func (o *WorkspaceIn) SetUnitSystemNil(b bool)`

 SetUnitSystemNil sets the value for UnitSystem to be an explicit nil

### UnsetUnitSystem
`func (o *WorkspaceIn) UnsetUnitSystem()`

UnsetUnitSystem ensures that no value is present for UnitSystem, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


