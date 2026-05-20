# DormantWorkspaceItem

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** |  | 
**Name** | **string** |  | 
**Slug** | **string** |  | 
**MemberCount** | **int32** |  | 
**SubscriptionTier** | Pointer to **string** |  | [optional] [default to "free"]
**LastActivity** | Pointer to **NullableString** |  | [optional] 
**InactiveDays** | **int32** |  | 
**CreatedAt** | **NullableString** |  | 

## Methods

### NewDormantWorkspaceItem

`func NewDormantWorkspaceItem(id string, name string, slug string, memberCount int32, inactiveDays int32, createdAt NullableString, ) *DormantWorkspaceItem`

NewDormantWorkspaceItem instantiates a new DormantWorkspaceItem object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDormantWorkspaceItemWithDefaults

`func NewDormantWorkspaceItemWithDefaults() *DormantWorkspaceItem`

NewDormantWorkspaceItemWithDefaults instantiates a new DormantWorkspaceItem object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *DormantWorkspaceItem) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *DormantWorkspaceItem) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *DormantWorkspaceItem) SetId(v string)`

SetId sets Id field to given value.


### GetName

`func (o *DormantWorkspaceItem) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *DormantWorkspaceItem) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *DormantWorkspaceItem) SetName(v string)`

SetName sets Name field to given value.


### GetSlug

`func (o *DormantWorkspaceItem) GetSlug() string`

GetSlug returns the Slug field if non-nil, zero value otherwise.

### GetSlugOk

`func (o *DormantWorkspaceItem) GetSlugOk() (*string, bool)`

GetSlugOk returns a tuple with the Slug field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSlug

`func (o *DormantWorkspaceItem) SetSlug(v string)`

SetSlug sets Slug field to given value.


### GetMemberCount

`func (o *DormantWorkspaceItem) GetMemberCount() int32`

GetMemberCount returns the MemberCount field if non-nil, zero value otherwise.

### GetMemberCountOk

`func (o *DormantWorkspaceItem) GetMemberCountOk() (*int32, bool)`

GetMemberCountOk returns a tuple with the MemberCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMemberCount

`func (o *DormantWorkspaceItem) SetMemberCount(v int32)`

SetMemberCount sets MemberCount field to given value.


### GetSubscriptionTier

`func (o *DormantWorkspaceItem) GetSubscriptionTier() string`

GetSubscriptionTier returns the SubscriptionTier field if non-nil, zero value otherwise.

### GetSubscriptionTierOk

`func (o *DormantWorkspaceItem) GetSubscriptionTierOk() (*string, bool)`

GetSubscriptionTierOk returns a tuple with the SubscriptionTier field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubscriptionTier

`func (o *DormantWorkspaceItem) SetSubscriptionTier(v string)`

SetSubscriptionTier sets SubscriptionTier field to given value.

### HasSubscriptionTier

`func (o *DormantWorkspaceItem) HasSubscriptionTier() bool`

HasSubscriptionTier returns a boolean if a field has been set.

### GetLastActivity

`func (o *DormantWorkspaceItem) GetLastActivity() string`

GetLastActivity returns the LastActivity field if non-nil, zero value otherwise.

### GetLastActivityOk

`func (o *DormantWorkspaceItem) GetLastActivityOk() (*string, bool)`

GetLastActivityOk returns a tuple with the LastActivity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastActivity

`func (o *DormantWorkspaceItem) SetLastActivity(v string)`

SetLastActivity sets LastActivity field to given value.

### HasLastActivity

`func (o *DormantWorkspaceItem) HasLastActivity() bool`

HasLastActivity returns a boolean if a field has been set.

### SetLastActivityNil

`func (o *DormantWorkspaceItem) SetLastActivityNil(b bool)`

 SetLastActivityNil sets the value for LastActivity to be an explicit nil

### UnsetLastActivity
`func (o *DormantWorkspaceItem) UnsetLastActivity()`

UnsetLastActivity ensures that no value is present for LastActivity, not even an explicit nil
### GetInactiveDays

`func (o *DormantWorkspaceItem) GetInactiveDays() int32`

GetInactiveDays returns the InactiveDays field if non-nil, zero value otherwise.

### GetInactiveDaysOk

`func (o *DormantWorkspaceItem) GetInactiveDaysOk() (*int32, bool)`

GetInactiveDaysOk returns a tuple with the InactiveDays field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInactiveDays

`func (o *DormantWorkspaceItem) SetInactiveDays(v int32)`

SetInactiveDays sets InactiveDays field to given value.


### GetCreatedAt

`func (o *DormantWorkspaceItem) GetCreatedAt() string`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *DormantWorkspaceItem) GetCreatedAtOk() (*string, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *DormantWorkspaceItem) SetCreatedAt(v string)`

SetCreatedAt sets CreatedAt field to given value.


### SetCreatedAtNil

`func (o *DormantWorkspaceItem) SetCreatedAtNil(b bool)`

 SetCreatedAtNil sets the value for CreatedAt to be an explicit nil

### UnsetCreatedAt
`func (o *DormantWorkspaceItem) UnsetCreatedAt()`

UnsetCreatedAt ensures that no value is present for CreatedAt, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


