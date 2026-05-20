# EnhancedWorkspaceListItem

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** |  | 
**Name** | **string** |  | 
**Slug** | **string** |  | 
**BillingEmail** | **NullableString** |  | 
**Website** | **NullableString** |  | 
**LogoUrl** | **NullableString** |  | 
**Timezone** | **string** |  | 
**MemberCount** | **int32** |  | 
**SubscriptionTier** | Pointer to **string** |  | [optional] [default to "free"]
**SubscriptionStatus** | Pointer to **string** |  | [optional] [default to "none"]
**UsageThisMonth** | Pointer to **float32** |  | [optional] [default to 0.0]
**LastActivity** | Pointer to **NullableString** |  | [optional] 
**CreatedAt** | **NullableString** |  | 
**UpdatedAt** | **NullableString** |  | 

## Methods

### NewEnhancedWorkspaceListItem

`func NewEnhancedWorkspaceListItem(id string, name string, slug string, billingEmail NullableString, website NullableString, logoUrl NullableString, timezone string, memberCount int32, createdAt NullableString, updatedAt NullableString, ) *EnhancedWorkspaceListItem`

NewEnhancedWorkspaceListItem instantiates a new EnhancedWorkspaceListItem object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEnhancedWorkspaceListItemWithDefaults

`func NewEnhancedWorkspaceListItemWithDefaults() *EnhancedWorkspaceListItem`

NewEnhancedWorkspaceListItemWithDefaults instantiates a new EnhancedWorkspaceListItem object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *EnhancedWorkspaceListItem) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *EnhancedWorkspaceListItem) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *EnhancedWorkspaceListItem) SetId(v string)`

SetId sets Id field to given value.


### GetName

`func (o *EnhancedWorkspaceListItem) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *EnhancedWorkspaceListItem) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *EnhancedWorkspaceListItem) SetName(v string)`

SetName sets Name field to given value.


### GetSlug

`func (o *EnhancedWorkspaceListItem) GetSlug() string`

GetSlug returns the Slug field if non-nil, zero value otherwise.

### GetSlugOk

`func (o *EnhancedWorkspaceListItem) GetSlugOk() (*string, bool)`

GetSlugOk returns a tuple with the Slug field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSlug

`func (o *EnhancedWorkspaceListItem) SetSlug(v string)`

SetSlug sets Slug field to given value.


### GetBillingEmail

`func (o *EnhancedWorkspaceListItem) GetBillingEmail() string`

GetBillingEmail returns the BillingEmail field if non-nil, zero value otherwise.

### GetBillingEmailOk

`func (o *EnhancedWorkspaceListItem) GetBillingEmailOk() (*string, bool)`

GetBillingEmailOk returns a tuple with the BillingEmail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBillingEmail

`func (o *EnhancedWorkspaceListItem) SetBillingEmail(v string)`

SetBillingEmail sets BillingEmail field to given value.


### SetBillingEmailNil

`func (o *EnhancedWorkspaceListItem) SetBillingEmailNil(b bool)`

 SetBillingEmailNil sets the value for BillingEmail to be an explicit nil

### UnsetBillingEmail
`func (o *EnhancedWorkspaceListItem) UnsetBillingEmail()`

UnsetBillingEmail ensures that no value is present for BillingEmail, not even an explicit nil
### GetWebsite

`func (o *EnhancedWorkspaceListItem) GetWebsite() string`

GetWebsite returns the Website field if non-nil, zero value otherwise.

### GetWebsiteOk

`func (o *EnhancedWorkspaceListItem) GetWebsiteOk() (*string, bool)`

GetWebsiteOk returns a tuple with the Website field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWebsite

`func (o *EnhancedWorkspaceListItem) SetWebsite(v string)`

SetWebsite sets Website field to given value.


### SetWebsiteNil

`func (o *EnhancedWorkspaceListItem) SetWebsiteNil(b bool)`

 SetWebsiteNil sets the value for Website to be an explicit nil

### UnsetWebsite
`func (o *EnhancedWorkspaceListItem) UnsetWebsite()`

UnsetWebsite ensures that no value is present for Website, not even an explicit nil
### GetLogoUrl

`func (o *EnhancedWorkspaceListItem) GetLogoUrl() string`

GetLogoUrl returns the LogoUrl field if non-nil, zero value otherwise.

### GetLogoUrlOk

`func (o *EnhancedWorkspaceListItem) GetLogoUrlOk() (*string, bool)`

GetLogoUrlOk returns a tuple with the LogoUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLogoUrl

`func (o *EnhancedWorkspaceListItem) SetLogoUrl(v string)`

SetLogoUrl sets LogoUrl field to given value.


### SetLogoUrlNil

`func (o *EnhancedWorkspaceListItem) SetLogoUrlNil(b bool)`

 SetLogoUrlNil sets the value for LogoUrl to be an explicit nil

### UnsetLogoUrl
`func (o *EnhancedWorkspaceListItem) UnsetLogoUrl()`

UnsetLogoUrl ensures that no value is present for LogoUrl, not even an explicit nil
### GetTimezone

`func (o *EnhancedWorkspaceListItem) GetTimezone() string`

GetTimezone returns the Timezone field if non-nil, zero value otherwise.

### GetTimezoneOk

`func (o *EnhancedWorkspaceListItem) GetTimezoneOk() (*string, bool)`

GetTimezoneOk returns a tuple with the Timezone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimezone

`func (o *EnhancedWorkspaceListItem) SetTimezone(v string)`

SetTimezone sets Timezone field to given value.


### GetMemberCount

`func (o *EnhancedWorkspaceListItem) GetMemberCount() int32`

GetMemberCount returns the MemberCount field if non-nil, zero value otherwise.

### GetMemberCountOk

`func (o *EnhancedWorkspaceListItem) GetMemberCountOk() (*int32, bool)`

GetMemberCountOk returns a tuple with the MemberCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMemberCount

`func (o *EnhancedWorkspaceListItem) SetMemberCount(v int32)`

SetMemberCount sets MemberCount field to given value.


### GetSubscriptionTier

`func (o *EnhancedWorkspaceListItem) GetSubscriptionTier() string`

GetSubscriptionTier returns the SubscriptionTier field if non-nil, zero value otherwise.

### GetSubscriptionTierOk

`func (o *EnhancedWorkspaceListItem) GetSubscriptionTierOk() (*string, bool)`

GetSubscriptionTierOk returns a tuple with the SubscriptionTier field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubscriptionTier

`func (o *EnhancedWorkspaceListItem) SetSubscriptionTier(v string)`

SetSubscriptionTier sets SubscriptionTier field to given value.

### HasSubscriptionTier

`func (o *EnhancedWorkspaceListItem) HasSubscriptionTier() bool`

HasSubscriptionTier returns a boolean if a field has been set.

### GetSubscriptionStatus

`func (o *EnhancedWorkspaceListItem) GetSubscriptionStatus() string`

GetSubscriptionStatus returns the SubscriptionStatus field if non-nil, zero value otherwise.

### GetSubscriptionStatusOk

`func (o *EnhancedWorkspaceListItem) GetSubscriptionStatusOk() (*string, bool)`

GetSubscriptionStatusOk returns a tuple with the SubscriptionStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubscriptionStatus

`func (o *EnhancedWorkspaceListItem) SetSubscriptionStatus(v string)`

SetSubscriptionStatus sets SubscriptionStatus field to given value.

### HasSubscriptionStatus

`func (o *EnhancedWorkspaceListItem) HasSubscriptionStatus() bool`

HasSubscriptionStatus returns a boolean if a field has been set.

### GetUsageThisMonth

`func (o *EnhancedWorkspaceListItem) GetUsageThisMonth() float32`

GetUsageThisMonth returns the UsageThisMonth field if non-nil, zero value otherwise.

### GetUsageThisMonthOk

`func (o *EnhancedWorkspaceListItem) GetUsageThisMonthOk() (*float32, bool)`

GetUsageThisMonthOk returns a tuple with the UsageThisMonth field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsageThisMonth

`func (o *EnhancedWorkspaceListItem) SetUsageThisMonth(v float32)`

SetUsageThisMonth sets UsageThisMonth field to given value.

### HasUsageThisMonth

`func (o *EnhancedWorkspaceListItem) HasUsageThisMonth() bool`

HasUsageThisMonth returns a boolean if a field has been set.

### GetLastActivity

`func (o *EnhancedWorkspaceListItem) GetLastActivity() string`

GetLastActivity returns the LastActivity field if non-nil, zero value otherwise.

### GetLastActivityOk

`func (o *EnhancedWorkspaceListItem) GetLastActivityOk() (*string, bool)`

GetLastActivityOk returns a tuple with the LastActivity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastActivity

`func (o *EnhancedWorkspaceListItem) SetLastActivity(v string)`

SetLastActivity sets LastActivity field to given value.

### HasLastActivity

`func (o *EnhancedWorkspaceListItem) HasLastActivity() bool`

HasLastActivity returns a boolean if a field has been set.

### SetLastActivityNil

`func (o *EnhancedWorkspaceListItem) SetLastActivityNil(b bool)`

 SetLastActivityNil sets the value for LastActivity to be an explicit nil

### UnsetLastActivity
`func (o *EnhancedWorkspaceListItem) UnsetLastActivity()`

UnsetLastActivity ensures that no value is present for LastActivity, not even an explicit nil
### GetCreatedAt

`func (o *EnhancedWorkspaceListItem) GetCreatedAt() string`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *EnhancedWorkspaceListItem) GetCreatedAtOk() (*string, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *EnhancedWorkspaceListItem) SetCreatedAt(v string)`

SetCreatedAt sets CreatedAt field to given value.


### SetCreatedAtNil

`func (o *EnhancedWorkspaceListItem) SetCreatedAtNil(b bool)`

 SetCreatedAtNil sets the value for CreatedAt to be an explicit nil

### UnsetCreatedAt
`func (o *EnhancedWorkspaceListItem) UnsetCreatedAt()`

UnsetCreatedAt ensures that no value is present for CreatedAt, not even an explicit nil
### GetUpdatedAt

`func (o *EnhancedWorkspaceListItem) GetUpdatedAt() string`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *EnhancedWorkspaceListItem) GetUpdatedAtOk() (*string, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *EnhancedWorkspaceListItem) SetUpdatedAt(v string)`

SetUpdatedAt sets UpdatedAt field to given value.


### SetUpdatedAtNil

`func (o *EnhancedWorkspaceListItem) SetUpdatedAtNil(b bool)`

 SetUpdatedAtNil sets the value for UpdatedAt to be an explicit nil

### UnsetUpdatedAt
`func (o *EnhancedWorkspaceListItem) UnsetUpdatedAt()`

UnsetUpdatedAt ensures that no value is present for UpdatedAt, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


