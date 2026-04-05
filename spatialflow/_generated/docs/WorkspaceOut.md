# WorkspaceOut

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **NullableString** |  | [optional] 
**Name** | **string** |  | 
**Slug** | **string** |  | 
**LogoUrl** | Pointer to **NullableString** |  | [optional] 
**Website** | Pointer to **NullableString** |  | [optional] 
**BillingEmail** | Pointer to **NullableString** |  | [optional] 
**Description** | Pointer to **NullableString** |  | [optional] 
**Timezone** | Pointer to **string** | Workspace default timezone | [optional] [default to "UTC"]
**SupportEmail** | Pointer to **NullableString** |  | [optional] 
**SlackConnectUrl** | Pointer to **NullableString** |  | [optional] 
**UnitSystem** | Pointer to **string** | Unit system for display (imperial: mi/mph/ft, metric: km/kph/m) | [optional] [default to "imperial"]
**CreatedAt** | **time.Time** |  | 
**UpdatedAt** | **time.Time** |  | 

## Methods

### NewWorkspaceOut

`func NewWorkspaceOut(name string, slug string, createdAt time.Time, updatedAt time.Time, ) *WorkspaceOut`

NewWorkspaceOut instantiates a new WorkspaceOut object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWorkspaceOutWithDefaults

`func NewWorkspaceOutWithDefaults() *WorkspaceOut`

NewWorkspaceOutWithDefaults instantiates a new WorkspaceOut object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *WorkspaceOut) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *WorkspaceOut) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *WorkspaceOut) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *WorkspaceOut) HasId() bool`

HasId returns a boolean if a field has been set.

### SetIdNil

`func (o *WorkspaceOut) SetIdNil(b bool)`

 SetIdNil sets the value for Id to be an explicit nil

### UnsetId
`func (o *WorkspaceOut) UnsetId()`

UnsetId ensures that no value is present for Id, not even an explicit nil
### GetName

`func (o *WorkspaceOut) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *WorkspaceOut) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *WorkspaceOut) SetName(v string)`

SetName sets Name field to given value.


### GetSlug

`func (o *WorkspaceOut) GetSlug() string`

GetSlug returns the Slug field if non-nil, zero value otherwise.

### GetSlugOk

`func (o *WorkspaceOut) GetSlugOk() (*string, bool)`

GetSlugOk returns a tuple with the Slug field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSlug

`func (o *WorkspaceOut) SetSlug(v string)`

SetSlug sets Slug field to given value.


### GetLogoUrl

`func (o *WorkspaceOut) GetLogoUrl() string`

GetLogoUrl returns the LogoUrl field if non-nil, zero value otherwise.

### GetLogoUrlOk

`func (o *WorkspaceOut) GetLogoUrlOk() (*string, bool)`

GetLogoUrlOk returns a tuple with the LogoUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLogoUrl

`func (o *WorkspaceOut) SetLogoUrl(v string)`

SetLogoUrl sets LogoUrl field to given value.

### HasLogoUrl

`func (o *WorkspaceOut) HasLogoUrl() bool`

HasLogoUrl returns a boolean if a field has been set.

### SetLogoUrlNil

`func (o *WorkspaceOut) SetLogoUrlNil(b bool)`

 SetLogoUrlNil sets the value for LogoUrl to be an explicit nil

### UnsetLogoUrl
`func (o *WorkspaceOut) UnsetLogoUrl()`

UnsetLogoUrl ensures that no value is present for LogoUrl, not even an explicit nil
### GetWebsite

`func (o *WorkspaceOut) GetWebsite() string`

GetWebsite returns the Website field if non-nil, zero value otherwise.

### GetWebsiteOk

`func (o *WorkspaceOut) GetWebsiteOk() (*string, bool)`

GetWebsiteOk returns a tuple with the Website field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWebsite

`func (o *WorkspaceOut) SetWebsite(v string)`

SetWebsite sets Website field to given value.

### HasWebsite

`func (o *WorkspaceOut) HasWebsite() bool`

HasWebsite returns a boolean if a field has been set.

### SetWebsiteNil

`func (o *WorkspaceOut) SetWebsiteNil(b bool)`

 SetWebsiteNil sets the value for Website to be an explicit nil

### UnsetWebsite
`func (o *WorkspaceOut) UnsetWebsite()`

UnsetWebsite ensures that no value is present for Website, not even an explicit nil
### GetBillingEmail

`func (o *WorkspaceOut) GetBillingEmail() string`

GetBillingEmail returns the BillingEmail field if non-nil, zero value otherwise.

### GetBillingEmailOk

`func (o *WorkspaceOut) GetBillingEmailOk() (*string, bool)`

GetBillingEmailOk returns a tuple with the BillingEmail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBillingEmail

`func (o *WorkspaceOut) SetBillingEmail(v string)`

SetBillingEmail sets BillingEmail field to given value.

### HasBillingEmail

`func (o *WorkspaceOut) HasBillingEmail() bool`

HasBillingEmail returns a boolean if a field has been set.

### SetBillingEmailNil

`func (o *WorkspaceOut) SetBillingEmailNil(b bool)`

 SetBillingEmailNil sets the value for BillingEmail to be an explicit nil

### UnsetBillingEmail
`func (o *WorkspaceOut) UnsetBillingEmail()`

UnsetBillingEmail ensures that no value is present for BillingEmail, not even an explicit nil
### GetDescription

`func (o *WorkspaceOut) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *WorkspaceOut) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *WorkspaceOut) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *WorkspaceOut) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescriptionNil

`func (o *WorkspaceOut) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *WorkspaceOut) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetTimezone

`func (o *WorkspaceOut) GetTimezone() string`

GetTimezone returns the Timezone field if non-nil, zero value otherwise.

### GetTimezoneOk

`func (o *WorkspaceOut) GetTimezoneOk() (*string, bool)`

GetTimezoneOk returns a tuple with the Timezone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimezone

`func (o *WorkspaceOut) SetTimezone(v string)`

SetTimezone sets Timezone field to given value.

### HasTimezone

`func (o *WorkspaceOut) HasTimezone() bool`

HasTimezone returns a boolean if a field has been set.

### GetSupportEmail

`func (o *WorkspaceOut) GetSupportEmail() string`

GetSupportEmail returns the SupportEmail field if non-nil, zero value otherwise.

### GetSupportEmailOk

`func (o *WorkspaceOut) GetSupportEmailOk() (*string, bool)`

GetSupportEmailOk returns a tuple with the SupportEmail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSupportEmail

`func (o *WorkspaceOut) SetSupportEmail(v string)`

SetSupportEmail sets SupportEmail field to given value.

### HasSupportEmail

`func (o *WorkspaceOut) HasSupportEmail() bool`

HasSupportEmail returns a boolean if a field has been set.

### SetSupportEmailNil

`func (o *WorkspaceOut) SetSupportEmailNil(b bool)`

 SetSupportEmailNil sets the value for SupportEmail to be an explicit nil

### UnsetSupportEmail
`func (o *WorkspaceOut) UnsetSupportEmail()`

UnsetSupportEmail ensures that no value is present for SupportEmail, not even an explicit nil
### GetSlackConnectUrl

`func (o *WorkspaceOut) GetSlackConnectUrl() string`

GetSlackConnectUrl returns the SlackConnectUrl field if non-nil, zero value otherwise.

### GetSlackConnectUrlOk

`func (o *WorkspaceOut) GetSlackConnectUrlOk() (*string, bool)`

GetSlackConnectUrlOk returns a tuple with the SlackConnectUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSlackConnectUrl

`func (o *WorkspaceOut) SetSlackConnectUrl(v string)`

SetSlackConnectUrl sets SlackConnectUrl field to given value.

### HasSlackConnectUrl

`func (o *WorkspaceOut) HasSlackConnectUrl() bool`

HasSlackConnectUrl returns a boolean if a field has been set.

### SetSlackConnectUrlNil

`func (o *WorkspaceOut) SetSlackConnectUrlNil(b bool)`

 SetSlackConnectUrlNil sets the value for SlackConnectUrl to be an explicit nil

### UnsetSlackConnectUrl
`func (o *WorkspaceOut) UnsetSlackConnectUrl()`

UnsetSlackConnectUrl ensures that no value is present for SlackConnectUrl, not even an explicit nil
### GetUnitSystem

`func (o *WorkspaceOut) GetUnitSystem() string`

GetUnitSystem returns the UnitSystem field if non-nil, zero value otherwise.

### GetUnitSystemOk

`func (o *WorkspaceOut) GetUnitSystemOk() (*string, bool)`

GetUnitSystemOk returns a tuple with the UnitSystem field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnitSystem

`func (o *WorkspaceOut) SetUnitSystem(v string)`

SetUnitSystem sets UnitSystem field to given value.

### HasUnitSystem

`func (o *WorkspaceOut) HasUnitSystem() bool`

HasUnitSystem returns a boolean if a field has been set.

### GetCreatedAt

`func (o *WorkspaceOut) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *WorkspaceOut) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *WorkspaceOut) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.


### GetUpdatedAt

`func (o *WorkspaceOut) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *WorkspaceOut) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *WorkspaceOut) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


