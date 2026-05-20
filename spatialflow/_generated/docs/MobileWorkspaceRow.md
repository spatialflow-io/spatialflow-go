# MobileWorkspaceRow

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** |  | 
**Name** | **string** |  | 
**Slug** | **string** |  | 
**Role** | **string** |  | 
**LogoUrl** | Pointer to **NullableString** |  | [optional] 
**Timezone** | **string** |  | 
**UnitSystem** | **string** |  | 
**IsSelected** | **bool** |  | 
**MemberCount** | Pointer to **int32** |  | [optional] [default to 0]

## Methods

### NewMobileWorkspaceRow

`func NewMobileWorkspaceRow(id string, name string, slug string, role string, timezone string, unitSystem string, isSelected bool, ) *MobileWorkspaceRow`

NewMobileWorkspaceRow instantiates a new MobileWorkspaceRow object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMobileWorkspaceRowWithDefaults

`func NewMobileWorkspaceRowWithDefaults() *MobileWorkspaceRow`

NewMobileWorkspaceRowWithDefaults instantiates a new MobileWorkspaceRow object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *MobileWorkspaceRow) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *MobileWorkspaceRow) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *MobileWorkspaceRow) SetId(v string)`

SetId sets Id field to given value.


### GetName

`func (o *MobileWorkspaceRow) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *MobileWorkspaceRow) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *MobileWorkspaceRow) SetName(v string)`

SetName sets Name field to given value.


### GetSlug

`func (o *MobileWorkspaceRow) GetSlug() string`

GetSlug returns the Slug field if non-nil, zero value otherwise.

### GetSlugOk

`func (o *MobileWorkspaceRow) GetSlugOk() (*string, bool)`

GetSlugOk returns a tuple with the Slug field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSlug

`func (o *MobileWorkspaceRow) SetSlug(v string)`

SetSlug sets Slug field to given value.


### GetRole

`func (o *MobileWorkspaceRow) GetRole() string`

GetRole returns the Role field if non-nil, zero value otherwise.

### GetRoleOk

`func (o *MobileWorkspaceRow) GetRoleOk() (*string, bool)`

GetRoleOk returns a tuple with the Role field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRole

`func (o *MobileWorkspaceRow) SetRole(v string)`

SetRole sets Role field to given value.


### GetLogoUrl

`func (o *MobileWorkspaceRow) GetLogoUrl() string`

GetLogoUrl returns the LogoUrl field if non-nil, zero value otherwise.

### GetLogoUrlOk

`func (o *MobileWorkspaceRow) GetLogoUrlOk() (*string, bool)`

GetLogoUrlOk returns a tuple with the LogoUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLogoUrl

`func (o *MobileWorkspaceRow) SetLogoUrl(v string)`

SetLogoUrl sets LogoUrl field to given value.

### HasLogoUrl

`func (o *MobileWorkspaceRow) HasLogoUrl() bool`

HasLogoUrl returns a boolean if a field has been set.

### SetLogoUrlNil

`func (o *MobileWorkspaceRow) SetLogoUrlNil(b bool)`

 SetLogoUrlNil sets the value for LogoUrl to be an explicit nil

### UnsetLogoUrl
`func (o *MobileWorkspaceRow) UnsetLogoUrl()`

UnsetLogoUrl ensures that no value is present for LogoUrl, not even an explicit nil
### GetTimezone

`func (o *MobileWorkspaceRow) GetTimezone() string`

GetTimezone returns the Timezone field if non-nil, zero value otherwise.

### GetTimezoneOk

`func (o *MobileWorkspaceRow) GetTimezoneOk() (*string, bool)`

GetTimezoneOk returns a tuple with the Timezone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimezone

`func (o *MobileWorkspaceRow) SetTimezone(v string)`

SetTimezone sets Timezone field to given value.


### GetUnitSystem

`func (o *MobileWorkspaceRow) GetUnitSystem() string`

GetUnitSystem returns the UnitSystem field if non-nil, zero value otherwise.

### GetUnitSystemOk

`func (o *MobileWorkspaceRow) GetUnitSystemOk() (*string, bool)`

GetUnitSystemOk returns a tuple with the UnitSystem field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnitSystem

`func (o *MobileWorkspaceRow) SetUnitSystem(v string)`

SetUnitSystem sets UnitSystem field to given value.


### GetIsSelected

`func (o *MobileWorkspaceRow) GetIsSelected() bool`

GetIsSelected returns the IsSelected field if non-nil, zero value otherwise.

### GetIsSelectedOk

`func (o *MobileWorkspaceRow) GetIsSelectedOk() (*bool, bool)`

GetIsSelectedOk returns a tuple with the IsSelected field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsSelected

`func (o *MobileWorkspaceRow) SetIsSelected(v bool)`

SetIsSelected sets IsSelected field to given value.


### GetMemberCount

`func (o *MobileWorkspaceRow) GetMemberCount() int32`

GetMemberCount returns the MemberCount field if non-nil, zero value otherwise.

### GetMemberCountOk

`func (o *MobileWorkspaceRow) GetMemberCountOk() (*int32, bool)`

GetMemberCountOk returns a tuple with the MemberCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMemberCount

`func (o *MobileWorkspaceRow) SetMemberCount(v int32)`

SetMemberCount sets MemberCount field to given value.

### HasMemberCount

`func (o *MobileWorkspaceRow) HasMemberCount() bool`

HasMemberCount returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


