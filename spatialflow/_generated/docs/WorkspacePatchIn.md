# WorkspacePatchIn

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **NullableString** |  | [optional] 
**LogoUrl** | Pointer to **NullableString** |  | [optional] 
**Website** | Pointer to **NullableString** |  | [optional] 
**BillingEmail** | Pointer to **NullableString** |  | [optional] 
**Description** | Pointer to **NullableString** |  | [optional] 
**Timezone** | Pointer to **NullableString** |  | [optional] 
**SupportEmail** | Pointer to **NullableString** |  | [optional] 
**SlackConnectUrl** | Pointer to **NullableString** |  | [optional] 
**UnitSystem** | Pointer to **NullableString** |  | [optional] 
**MapHome** | Pointer to **map[string]interface{}** |  | [optional] 

## Methods

### NewWorkspacePatchIn

`func NewWorkspacePatchIn() *WorkspacePatchIn`

NewWorkspacePatchIn instantiates a new WorkspacePatchIn object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWorkspacePatchInWithDefaults

`func NewWorkspacePatchInWithDefaults() *WorkspacePatchIn`

NewWorkspacePatchInWithDefaults instantiates a new WorkspacePatchIn object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *WorkspacePatchIn) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *WorkspacePatchIn) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *WorkspacePatchIn) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *WorkspacePatchIn) HasName() bool`

HasName returns a boolean if a field has been set.

### SetNameNil

`func (o *WorkspacePatchIn) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *WorkspacePatchIn) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetLogoUrl

`func (o *WorkspacePatchIn) GetLogoUrl() string`

GetLogoUrl returns the LogoUrl field if non-nil, zero value otherwise.

### GetLogoUrlOk

`func (o *WorkspacePatchIn) GetLogoUrlOk() (*string, bool)`

GetLogoUrlOk returns a tuple with the LogoUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLogoUrl

`func (o *WorkspacePatchIn) SetLogoUrl(v string)`

SetLogoUrl sets LogoUrl field to given value.

### HasLogoUrl

`func (o *WorkspacePatchIn) HasLogoUrl() bool`

HasLogoUrl returns a boolean if a field has been set.

### SetLogoUrlNil

`func (o *WorkspacePatchIn) SetLogoUrlNil(b bool)`

 SetLogoUrlNil sets the value for LogoUrl to be an explicit nil

### UnsetLogoUrl
`func (o *WorkspacePatchIn) UnsetLogoUrl()`

UnsetLogoUrl ensures that no value is present for LogoUrl, not even an explicit nil
### GetWebsite

`func (o *WorkspacePatchIn) GetWebsite() string`

GetWebsite returns the Website field if non-nil, zero value otherwise.

### GetWebsiteOk

`func (o *WorkspacePatchIn) GetWebsiteOk() (*string, bool)`

GetWebsiteOk returns a tuple with the Website field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWebsite

`func (o *WorkspacePatchIn) SetWebsite(v string)`

SetWebsite sets Website field to given value.

### HasWebsite

`func (o *WorkspacePatchIn) HasWebsite() bool`

HasWebsite returns a boolean if a field has been set.

### SetWebsiteNil

`func (o *WorkspacePatchIn) SetWebsiteNil(b bool)`

 SetWebsiteNil sets the value for Website to be an explicit nil

### UnsetWebsite
`func (o *WorkspacePatchIn) UnsetWebsite()`

UnsetWebsite ensures that no value is present for Website, not even an explicit nil
### GetBillingEmail

`func (o *WorkspacePatchIn) GetBillingEmail() string`

GetBillingEmail returns the BillingEmail field if non-nil, zero value otherwise.

### GetBillingEmailOk

`func (o *WorkspacePatchIn) GetBillingEmailOk() (*string, bool)`

GetBillingEmailOk returns a tuple with the BillingEmail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBillingEmail

`func (o *WorkspacePatchIn) SetBillingEmail(v string)`

SetBillingEmail sets BillingEmail field to given value.

### HasBillingEmail

`func (o *WorkspacePatchIn) HasBillingEmail() bool`

HasBillingEmail returns a boolean if a field has been set.

### SetBillingEmailNil

`func (o *WorkspacePatchIn) SetBillingEmailNil(b bool)`

 SetBillingEmailNil sets the value for BillingEmail to be an explicit nil

### UnsetBillingEmail
`func (o *WorkspacePatchIn) UnsetBillingEmail()`

UnsetBillingEmail ensures that no value is present for BillingEmail, not even an explicit nil
### GetDescription

`func (o *WorkspacePatchIn) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *WorkspacePatchIn) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *WorkspacePatchIn) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *WorkspacePatchIn) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescriptionNil

`func (o *WorkspacePatchIn) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *WorkspacePatchIn) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetTimezone

`func (o *WorkspacePatchIn) GetTimezone() string`

GetTimezone returns the Timezone field if non-nil, zero value otherwise.

### GetTimezoneOk

`func (o *WorkspacePatchIn) GetTimezoneOk() (*string, bool)`

GetTimezoneOk returns a tuple with the Timezone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimezone

`func (o *WorkspacePatchIn) SetTimezone(v string)`

SetTimezone sets Timezone field to given value.

### HasTimezone

`func (o *WorkspacePatchIn) HasTimezone() bool`

HasTimezone returns a boolean if a field has been set.

### SetTimezoneNil

`func (o *WorkspacePatchIn) SetTimezoneNil(b bool)`

 SetTimezoneNil sets the value for Timezone to be an explicit nil

### UnsetTimezone
`func (o *WorkspacePatchIn) UnsetTimezone()`

UnsetTimezone ensures that no value is present for Timezone, not even an explicit nil
### GetSupportEmail

`func (o *WorkspacePatchIn) GetSupportEmail() string`

GetSupportEmail returns the SupportEmail field if non-nil, zero value otherwise.

### GetSupportEmailOk

`func (o *WorkspacePatchIn) GetSupportEmailOk() (*string, bool)`

GetSupportEmailOk returns a tuple with the SupportEmail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSupportEmail

`func (o *WorkspacePatchIn) SetSupportEmail(v string)`

SetSupportEmail sets SupportEmail field to given value.

### HasSupportEmail

`func (o *WorkspacePatchIn) HasSupportEmail() bool`

HasSupportEmail returns a boolean if a field has been set.

### SetSupportEmailNil

`func (o *WorkspacePatchIn) SetSupportEmailNil(b bool)`

 SetSupportEmailNil sets the value for SupportEmail to be an explicit nil

### UnsetSupportEmail
`func (o *WorkspacePatchIn) UnsetSupportEmail()`

UnsetSupportEmail ensures that no value is present for SupportEmail, not even an explicit nil
### GetSlackConnectUrl

`func (o *WorkspacePatchIn) GetSlackConnectUrl() string`

GetSlackConnectUrl returns the SlackConnectUrl field if non-nil, zero value otherwise.

### GetSlackConnectUrlOk

`func (o *WorkspacePatchIn) GetSlackConnectUrlOk() (*string, bool)`

GetSlackConnectUrlOk returns a tuple with the SlackConnectUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSlackConnectUrl

`func (o *WorkspacePatchIn) SetSlackConnectUrl(v string)`

SetSlackConnectUrl sets SlackConnectUrl field to given value.

### HasSlackConnectUrl

`func (o *WorkspacePatchIn) HasSlackConnectUrl() bool`

HasSlackConnectUrl returns a boolean if a field has been set.

### SetSlackConnectUrlNil

`func (o *WorkspacePatchIn) SetSlackConnectUrlNil(b bool)`

 SetSlackConnectUrlNil sets the value for SlackConnectUrl to be an explicit nil

### UnsetSlackConnectUrl
`func (o *WorkspacePatchIn) UnsetSlackConnectUrl()`

UnsetSlackConnectUrl ensures that no value is present for SlackConnectUrl, not even an explicit nil
### GetUnitSystem

`func (o *WorkspacePatchIn) GetUnitSystem() string`

GetUnitSystem returns the UnitSystem field if non-nil, zero value otherwise.

### GetUnitSystemOk

`func (o *WorkspacePatchIn) GetUnitSystemOk() (*string, bool)`

GetUnitSystemOk returns a tuple with the UnitSystem field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnitSystem

`func (o *WorkspacePatchIn) SetUnitSystem(v string)`

SetUnitSystem sets UnitSystem field to given value.

### HasUnitSystem

`func (o *WorkspacePatchIn) HasUnitSystem() bool`

HasUnitSystem returns a boolean if a field has been set.

### SetUnitSystemNil

`func (o *WorkspacePatchIn) SetUnitSystemNil(b bool)`

 SetUnitSystemNil sets the value for UnitSystem to be an explicit nil

### UnsetUnitSystem
`func (o *WorkspacePatchIn) UnsetUnitSystem()`

UnsetUnitSystem ensures that no value is present for UnitSystem, not even an explicit nil
### GetMapHome

`func (o *WorkspacePatchIn) GetMapHome() map[string]interface{}`

GetMapHome returns the MapHome field if non-nil, zero value otherwise.

### GetMapHomeOk

`func (o *WorkspacePatchIn) GetMapHomeOk() (*map[string]interface{}, bool)`

GetMapHomeOk returns a tuple with the MapHome field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMapHome

`func (o *WorkspacePatchIn) SetMapHome(v map[string]interface{})`

SetMapHome sets MapHome field to given value.

### HasMapHome

`func (o *WorkspacePatchIn) HasMapHome() bool`

HasMapHome returns a boolean if a field has been set.

### SetMapHomeNil

`func (o *WorkspacePatchIn) SetMapHomeNil(b bool)`

 SetMapHomeNil sets the value for MapHome to be an explicit nil

### UnsetMapHome
`func (o *WorkspacePatchIn) UnsetMapHome()`

UnsetMapHome ensures that no value is present for MapHome, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


