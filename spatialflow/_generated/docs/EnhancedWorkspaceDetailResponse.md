# EnhancedWorkspaceDetailResponse

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
**Subscription** | [**SubscriptionInfo**](SubscriptionInfo.md) |  | 
**Usage** | [**UsageLimits**](UsageLimits.md) |  | 
**CreatedAt** | **NullableString** |  | 
**UpdatedAt** | **NullableString** |  | 

## Methods

### NewEnhancedWorkspaceDetailResponse

`func NewEnhancedWorkspaceDetailResponse(id string, name string, slug string, billingEmail NullableString, website NullableString, logoUrl NullableString, timezone string, stripeCustomerId NullableString, memberCount int32, subscription SubscriptionInfo, usage UsageLimits, createdAt NullableString, updatedAt NullableString, ) *EnhancedWorkspaceDetailResponse`

NewEnhancedWorkspaceDetailResponse instantiates a new EnhancedWorkspaceDetailResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEnhancedWorkspaceDetailResponseWithDefaults

`func NewEnhancedWorkspaceDetailResponseWithDefaults() *EnhancedWorkspaceDetailResponse`

NewEnhancedWorkspaceDetailResponseWithDefaults instantiates a new EnhancedWorkspaceDetailResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *EnhancedWorkspaceDetailResponse) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *EnhancedWorkspaceDetailResponse) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *EnhancedWorkspaceDetailResponse) SetId(v string)`

SetId sets Id field to given value.


### GetName

`func (o *EnhancedWorkspaceDetailResponse) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *EnhancedWorkspaceDetailResponse) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *EnhancedWorkspaceDetailResponse) SetName(v string)`

SetName sets Name field to given value.


### GetSlug

`func (o *EnhancedWorkspaceDetailResponse) GetSlug() string`

GetSlug returns the Slug field if non-nil, zero value otherwise.

### GetSlugOk

`func (o *EnhancedWorkspaceDetailResponse) GetSlugOk() (*string, bool)`

GetSlugOk returns a tuple with the Slug field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSlug

`func (o *EnhancedWorkspaceDetailResponse) SetSlug(v string)`

SetSlug sets Slug field to given value.


### GetBillingEmail

`func (o *EnhancedWorkspaceDetailResponse) GetBillingEmail() string`

GetBillingEmail returns the BillingEmail field if non-nil, zero value otherwise.

### GetBillingEmailOk

`func (o *EnhancedWorkspaceDetailResponse) GetBillingEmailOk() (*string, bool)`

GetBillingEmailOk returns a tuple with the BillingEmail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBillingEmail

`func (o *EnhancedWorkspaceDetailResponse) SetBillingEmail(v string)`

SetBillingEmail sets BillingEmail field to given value.


### SetBillingEmailNil

`func (o *EnhancedWorkspaceDetailResponse) SetBillingEmailNil(b bool)`

 SetBillingEmailNil sets the value for BillingEmail to be an explicit nil

### UnsetBillingEmail
`func (o *EnhancedWorkspaceDetailResponse) UnsetBillingEmail()`

UnsetBillingEmail ensures that no value is present for BillingEmail, not even an explicit nil
### GetWebsite

`func (o *EnhancedWorkspaceDetailResponse) GetWebsite() string`

GetWebsite returns the Website field if non-nil, zero value otherwise.

### GetWebsiteOk

`func (o *EnhancedWorkspaceDetailResponse) GetWebsiteOk() (*string, bool)`

GetWebsiteOk returns a tuple with the Website field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWebsite

`func (o *EnhancedWorkspaceDetailResponse) SetWebsite(v string)`

SetWebsite sets Website field to given value.


### SetWebsiteNil

`func (o *EnhancedWorkspaceDetailResponse) SetWebsiteNil(b bool)`

 SetWebsiteNil sets the value for Website to be an explicit nil

### UnsetWebsite
`func (o *EnhancedWorkspaceDetailResponse) UnsetWebsite()`

UnsetWebsite ensures that no value is present for Website, not even an explicit nil
### GetLogoUrl

`func (o *EnhancedWorkspaceDetailResponse) GetLogoUrl() string`

GetLogoUrl returns the LogoUrl field if non-nil, zero value otherwise.

### GetLogoUrlOk

`func (o *EnhancedWorkspaceDetailResponse) GetLogoUrlOk() (*string, bool)`

GetLogoUrlOk returns a tuple with the LogoUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLogoUrl

`func (o *EnhancedWorkspaceDetailResponse) SetLogoUrl(v string)`

SetLogoUrl sets LogoUrl field to given value.


### SetLogoUrlNil

`func (o *EnhancedWorkspaceDetailResponse) SetLogoUrlNil(b bool)`

 SetLogoUrlNil sets the value for LogoUrl to be an explicit nil

### UnsetLogoUrl
`func (o *EnhancedWorkspaceDetailResponse) UnsetLogoUrl()`

UnsetLogoUrl ensures that no value is present for LogoUrl, not even an explicit nil
### GetTimezone

`func (o *EnhancedWorkspaceDetailResponse) GetTimezone() string`

GetTimezone returns the Timezone field if non-nil, zero value otherwise.

### GetTimezoneOk

`func (o *EnhancedWorkspaceDetailResponse) GetTimezoneOk() (*string, bool)`

GetTimezoneOk returns a tuple with the Timezone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimezone

`func (o *EnhancedWorkspaceDetailResponse) SetTimezone(v string)`

SetTimezone sets Timezone field to given value.


### GetStripeCustomerId

`func (o *EnhancedWorkspaceDetailResponse) GetStripeCustomerId() string`

GetStripeCustomerId returns the StripeCustomerId field if non-nil, zero value otherwise.

### GetStripeCustomerIdOk

`func (o *EnhancedWorkspaceDetailResponse) GetStripeCustomerIdOk() (*string, bool)`

GetStripeCustomerIdOk returns a tuple with the StripeCustomerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStripeCustomerId

`func (o *EnhancedWorkspaceDetailResponse) SetStripeCustomerId(v string)`

SetStripeCustomerId sets StripeCustomerId field to given value.


### SetStripeCustomerIdNil

`func (o *EnhancedWorkspaceDetailResponse) SetStripeCustomerIdNil(b bool)`

 SetStripeCustomerIdNil sets the value for StripeCustomerId to be an explicit nil

### UnsetStripeCustomerId
`func (o *EnhancedWorkspaceDetailResponse) UnsetStripeCustomerId()`

UnsetStripeCustomerId ensures that no value is present for StripeCustomerId, not even an explicit nil
### GetMemberCount

`func (o *EnhancedWorkspaceDetailResponse) GetMemberCount() int32`

GetMemberCount returns the MemberCount field if non-nil, zero value otherwise.

### GetMemberCountOk

`func (o *EnhancedWorkspaceDetailResponse) GetMemberCountOk() (*int32, bool)`

GetMemberCountOk returns a tuple with the MemberCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMemberCount

`func (o *EnhancedWorkspaceDetailResponse) SetMemberCount(v int32)`

SetMemberCount sets MemberCount field to given value.


### GetSubscription

`func (o *EnhancedWorkspaceDetailResponse) GetSubscription() SubscriptionInfo`

GetSubscription returns the Subscription field if non-nil, zero value otherwise.

### GetSubscriptionOk

`func (o *EnhancedWorkspaceDetailResponse) GetSubscriptionOk() (*SubscriptionInfo, bool)`

GetSubscriptionOk returns a tuple with the Subscription field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubscription

`func (o *EnhancedWorkspaceDetailResponse) SetSubscription(v SubscriptionInfo)`

SetSubscription sets Subscription field to given value.


### GetUsage

`func (o *EnhancedWorkspaceDetailResponse) GetUsage() UsageLimits`

GetUsage returns the Usage field if non-nil, zero value otherwise.

### GetUsageOk

`func (o *EnhancedWorkspaceDetailResponse) GetUsageOk() (*UsageLimits, bool)`

GetUsageOk returns a tuple with the Usage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsage

`func (o *EnhancedWorkspaceDetailResponse) SetUsage(v UsageLimits)`

SetUsage sets Usage field to given value.


### GetCreatedAt

`func (o *EnhancedWorkspaceDetailResponse) GetCreatedAt() string`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *EnhancedWorkspaceDetailResponse) GetCreatedAtOk() (*string, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *EnhancedWorkspaceDetailResponse) SetCreatedAt(v string)`

SetCreatedAt sets CreatedAt field to given value.


### SetCreatedAtNil

`func (o *EnhancedWorkspaceDetailResponse) SetCreatedAtNil(b bool)`

 SetCreatedAtNil sets the value for CreatedAt to be an explicit nil

### UnsetCreatedAt
`func (o *EnhancedWorkspaceDetailResponse) UnsetCreatedAt()`

UnsetCreatedAt ensures that no value is present for CreatedAt, not even an explicit nil
### GetUpdatedAt

`func (o *EnhancedWorkspaceDetailResponse) GetUpdatedAt() string`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *EnhancedWorkspaceDetailResponse) GetUpdatedAtOk() (*string, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *EnhancedWorkspaceDetailResponse) SetUpdatedAt(v string)`

SetUpdatedAt sets UpdatedAt field to given value.


### SetUpdatedAtNil

`func (o *EnhancedWorkspaceDetailResponse) SetUpdatedAtNil(b bool)`

 SetUpdatedAtNil sets the value for UpdatedAt to be an explicit nil

### UnsetUpdatedAt
`func (o *EnhancedWorkspaceDetailResponse) UnsetUpdatedAt()`

UnsetUpdatedAt ensures that no value is present for UpdatedAt, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


