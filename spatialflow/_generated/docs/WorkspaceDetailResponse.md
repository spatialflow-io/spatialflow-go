# WorkspaceDetailResponse

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
**StripeCustomerId** | **NullableString** |  | 
**MemberCount** | **int32** |  | 
**UsageStats** | [**UsageStats**](UsageStats.md) |  | 
**CreatedAt** | **NullableString** |  | 
**UpdatedAt** | **NullableString** |  | 

## Methods

### NewWorkspaceDetailResponse

`func NewWorkspaceDetailResponse(id string, name string, slug string, billingEmail NullableString, website NullableString, logoUrl NullableString, timezone string, stripeCustomerId NullableString, memberCount int32, usageStats UsageStats, createdAt NullableString, updatedAt NullableString, ) *WorkspaceDetailResponse`

NewWorkspaceDetailResponse instantiates a new WorkspaceDetailResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWorkspaceDetailResponseWithDefaults

`func NewWorkspaceDetailResponseWithDefaults() *WorkspaceDetailResponse`

NewWorkspaceDetailResponseWithDefaults instantiates a new WorkspaceDetailResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *WorkspaceDetailResponse) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *WorkspaceDetailResponse) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *WorkspaceDetailResponse) SetId(v string)`

SetId sets Id field to given value.


### GetName

`func (o *WorkspaceDetailResponse) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *WorkspaceDetailResponse) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *WorkspaceDetailResponse) SetName(v string)`

SetName sets Name field to given value.


### GetSlug

`func (o *WorkspaceDetailResponse) GetSlug() string`

GetSlug returns the Slug field if non-nil, zero value otherwise.

### GetSlugOk

`func (o *WorkspaceDetailResponse) GetSlugOk() (*string, bool)`

GetSlugOk returns a tuple with the Slug field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSlug

`func (o *WorkspaceDetailResponse) SetSlug(v string)`

SetSlug sets Slug field to given value.


### GetBillingEmail

`func (o *WorkspaceDetailResponse) GetBillingEmail() string`

GetBillingEmail returns the BillingEmail field if non-nil, zero value otherwise.

### GetBillingEmailOk

`func (o *WorkspaceDetailResponse) GetBillingEmailOk() (*string, bool)`

GetBillingEmailOk returns a tuple with the BillingEmail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBillingEmail

`func (o *WorkspaceDetailResponse) SetBillingEmail(v string)`

SetBillingEmail sets BillingEmail field to given value.


### SetBillingEmailNil

`func (o *WorkspaceDetailResponse) SetBillingEmailNil(b bool)`

 SetBillingEmailNil sets the value for BillingEmail to be an explicit nil

### UnsetBillingEmail
`func (o *WorkspaceDetailResponse) UnsetBillingEmail()`

UnsetBillingEmail ensures that no value is present for BillingEmail, not even an explicit nil
### GetWebsite

`func (o *WorkspaceDetailResponse) GetWebsite() string`

GetWebsite returns the Website field if non-nil, zero value otherwise.

### GetWebsiteOk

`func (o *WorkspaceDetailResponse) GetWebsiteOk() (*string, bool)`

GetWebsiteOk returns a tuple with the Website field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWebsite

`func (o *WorkspaceDetailResponse) SetWebsite(v string)`

SetWebsite sets Website field to given value.


### SetWebsiteNil

`func (o *WorkspaceDetailResponse) SetWebsiteNil(b bool)`

 SetWebsiteNil sets the value for Website to be an explicit nil

### UnsetWebsite
`func (o *WorkspaceDetailResponse) UnsetWebsite()`

UnsetWebsite ensures that no value is present for Website, not even an explicit nil
### GetLogoUrl

`func (o *WorkspaceDetailResponse) GetLogoUrl() string`

GetLogoUrl returns the LogoUrl field if non-nil, zero value otherwise.

### GetLogoUrlOk

`func (o *WorkspaceDetailResponse) GetLogoUrlOk() (*string, bool)`

GetLogoUrlOk returns a tuple with the LogoUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLogoUrl

`func (o *WorkspaceDetailResponse) SetLogoUrl(v string)`

SetLogoUrl sets LogoUrl field to given value.


### SetLogoUrlNil

`func (o *WorkspaceDetailResponse) SetLogoUrlNil(b bool)`

 SetLogoUrlNil sets the value for LogoUrl to be an explicit nil

### UnsetLogoUrl
`func (o *WorkspaceDetailResponse) UnsetLogoUrl()`

UnsetLogoUrl ensures that no value is present for LogoUrl, not even an explicit nil
### GetTimezone

`func (o *WorkspaceDetailResponse) GetTimezone() string`

GetTimezone returns the Timezone field if non-nil, zero value otherwise.

### GetTimezoneOk

`func (o *WorkspaceDetailResponse) GetTimezoneOk() (*string, bool)`

GetTimezoneOk returns a tuple with the Timezone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimezone

`func (o *WorkspaceDetailResponse) SetTimezone(v string)`

SetTimezone sets Timezone field to given value.


### GetStripeCustomerId

`func (o *WorkspaceDetailResponse) GetStripeCustomerId() string`

GetStripeCustomerId returns the StripeCustomerId field if non-nil, zero value otherwise.

### GetStripeCustomerIdOk

`func (o *WorkspaceDetailResponse) GetStripeCustomerIdOk() (*string, bool)`

GetStripeCustomerIdOk returns a tuple with the StripeCustomerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStripeCustomerId

`func (o *WorkspaceDetailResponse) SetStripeCustomerId(v string)`

SetStripeCustomerId sets StripeCustomerId field to given value.


### SetStripeCustomerIdNil

`func (o *WorkspaceDetailResponse) SetStripeCustomerIdNil(b bool)`

 SetStripeCustomerIdNil sets the value for StripeCustomerId to be an explicit nil

### UnsetStripeCustomerId
`func (o *WorkspaceDetailResponse) UnsetStripeCustomerId()`

UnsetStripeCustomerId ensures that no value is present for StripeCustomerId, not even an explicit nil
### GetMemberCount

`func (o *WorkspaceDetailResponse) GetMemberCount() int32`

GetMemberCount returns the MemberCount field if non-nil, zero value otherwise.

### GetMemberCountOk

`func (o *WorkspaceDetailResponse) GetMemberCountOk() (*int32, bool)`

GetMemberCountOk returns a tuple with the MemberCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMemberCount

`func (o *WorkspaceDetailResponse) SetMemberCount(v int32)`

SetMemberCount sets MemberCount field to given value.


### GetUsageStats

`func (o *WorkspaceDetailResponse) GetUsageStats() UsageStats`

GetUsageStats returns the UsageStats field if non-nil, zero value otherwise.

### GetUsageStatsOk

`func (o *WorkspaceDetailResponse) GetUsageStatsOk() (*UsageStats, bool)`

GetUsageStatsOk returns a tuple with the UsageStats field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsageStats

`func (o *WorkspaceDetailResponse) SetUsageStats(v UsageStats)`

SetUsageStats sets UsageStats field to given value.


### GetCreatedAt

`func (o *WorkspaceDetailResponse) GetCreatedAt() string`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *WorkspaceDetailResponse) GetCreatedAtOk() (*string, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *WorkspaceDetailResponse) SetCreatedAt(v string)`

SetCreatedAt sets CreatedAt field to given value.


### SetCreatedAtNil

`func (o *WorkspaceDetailResponse) SetCreatedAtNil(b bool)`

 SetCreatedAtNil sets the value for CreatedAt to be an explicit nil

### UnsetCreatedAt
`func (o *WorkspaceDetailResponse) UnsetCreatedAt()`

UnsetCreatedAt ensures that no value is present for CreatedAt, not even an explicit nil
### GetUpdatedAt

`func (o *WorkspaceDetailResponse) GetUpdatedAt() string`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *WorkspaceDetailResponse) GetUpdatedAtOk() (*string, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *WorkspaceDetailResponse) SetUpdatedAt(v string)`

SetUpdatedAt sets UpdatedAt field to given value.


### SetUpdatedAtNil

`func (o *WorkspaceDetailResponse) SetUpdatedAtNil(b bool)`

 SetUpdatedAtNil sets the value for UpdatedAt to be an explicit nil

### UnsetUpdatedAt
`func (o *WorkspaceDetailResponse) UnsetUpdatedAt()`

UnsetUpdatedAt ensures that no value is present for UpdatedAt, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


