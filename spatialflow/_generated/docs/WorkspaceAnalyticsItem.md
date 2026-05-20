# WorkspaceAnalyticsItem

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** |  | 
**Name** | **string** |  | 
**Slug** | **string** |  | 
**MemberCount** | **int32** |  | 
**SubscriptionTier** | Pointer to **string** |  | [optional] [default to "free"]
**SubscriptionStatus** | Pointer to **string** |  | [optional] [default to "none"]
**LocationEvents** | Pointer to **int32** |  | [optional] [default to 0]
**ActionDeliveries** | Pointer to **int32** |  | [optional] [default to 0]
**EventUnits** | Pointer to **float32** |  | [optional] [default to 0.0]
**LastActivity** | Pointer to **NullableString** |  | [optional] 

## Methods

### NewWorkspaceAnalyticsItem

`func NewWorkspaceAnalyticsItem(id string, name string, slug string, memberCount int32, ) *WorkspaceAnalyticsItem`

NewWorkspaceAnalyticsItem instantiates a new WorkspaceAnalyticsItem object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWorkspaceAnalyticsItemWithDefaults

`func NewWorkspaceAnalyticsItemWithDefaults() *WorkspaceAnalyticsItem`

NewWorkspaceAnalyticsItemWithDefaults instantiates a new WorkspaceAnalyticsItem object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *WorkspaceAnalyticsItem) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *WorkspaceAnalyticsItem) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *WorkspaceAnalyticsItem) SetId(v string)`

SetId sets Id field to given value.


### GetName

`func (o *WorkspaceAnalyticsItem) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *WorkspaceAnalyticsItem) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *WorkspaceAnalyticsItem) SetName(v string)`

SetName sets Name field to given value.


### GetSlug

`func (o *WorkspaceAnalyticsItem) GetSlug() string`

GetSlug returns the Slug field if non-nil, zero value otherwise.

### GetSlugOk

`func (o *WorkspaceAnalyticsItem) GetSlugOk() (*string, bool)`

GetSlugOk returns a tuple with the Slug field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSlug

`func (o *WorkspaceAnalyticsItem) SetSlug(v string)`

SetSlug sets Slug field to given value.


### GetMemberCount

`func (o *WorkspaceAnalyticsItem) GetMemberCount() int32`

GetMemberCount returns the MemberCount field if non-nil, zero value otherwise.

### GetMemberCountOk

`func (o *WorkspaceAnalyticsItem) GetMemberCountOk() (*int32, bool)`

GetMemberCountOk returns a tuple with the MemberCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMemberCount

`func (o *WorkspaceAnalyticsItem) SetMemberCount(v int32)`

SetMemberCount sets MemberCount field to given value.


### GetSubscriptionTier

`func (o *WorkspaceAnalyticsItem) GetSubscriptionTier() string`

GetSubscriptionTier returns the SubscriptionTier field if non-nil, zero value otherwise.

### GetSubscriptionTierOk

`func (o *WorkspaceAnalyticsItem) GetSubscriptionTierOk() (*string, bool)`

GetSubscriptionTierOk returns a tuple with the SubscriptionTier field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubscriptionTier

`func (o *WorkspaceAnalyticsItem) SetSubscriptionTier(v string)`

SetSubscriptionTier sets SubscriptionTier field to given value.

### HasSubscriptionTier

`func (o *WorkspaceAnalyticsItem) HasSubscriptionTier() bool`

HasSubscriptionTier returns a boolean if a field has been set.

### GetSubscriptionStatus

`func (o *WorkspaceAnalyticsItem) GetSubscriptionStatus() string`

GetSubscriptionStatus returns the SubscriptionStatus field if non-nil, zero value otherwise.

### GetSubscriptionStatusOk

`func (o *WorkspaceAnalyticsItem) GetSubscriptionStatusOk() (*string, bool)`

GetSubscriptionStatusOk returns a tuple with the SubscriptionStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubscriptionStatus

`func (o *WorkspaceAnalyticsItem) SetSubscriptionStatus(v string)`

SetSubscriptionStatus sets SubscriptionStatus field to given value.

### HasSubscriptionStatus

`func (o *WorkspaceAnalyticsItem) HasSubscriptionStatus() bool`

HasSubscriptionStatus returns a boolean if a field has been set.

### GetLocationEvents

`func (o *WorkspaceAnalyticsItem) GetLocationEvents() int32`

GetLocationEvents returns the LocationEvents field if non-nil, zero value otherwise.

### GetLocationEventsOk

`func (o *WorkspaceAnalyticsItem) GetLocationEventsOk() (*int32, bool)`

GetLocationEventsOk returns a tuple with the LocationEvents field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocationEvents

`func (o *WorkspaceAnalyticsItem) SetLocationEvents(v int32)`

SetLocationEvents sets LocationEvents field to given value.

### HasLocationEvents

`func (o *WorkspaceAnalyticsItem) HasLocationEvents() bool`

HasLocationEvents returns a boolean if a field has been set.

### GetActionDeliveries

`func (o *WorkspaceAnalyticsItem) GetActionDeliveries() int32`

GetActionDeliveries returns the ActionDeliveries field if non-nil, zero value otherwise.

### GetActionDeliveriesOk

`func (o *WorkspaceAnalyticsItem) GetActionDeliveriesOk() (*int32, bool)`

GetActionDeliveriesOk returns a tuple with the ActionDeliveries field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActionDeliveries

`func (o *WorkspaceAnalyticsItem) SetActionDeliveries(v int32)`

SetActionDeliveries sets ActionDeliveries field to given value.

### HasActionDeliveries

`func (o *WorkspaceAnalyticsItem) HasActionDeliveries() bool`

HasActionDeliveries returns a boolean if a field has been set.

### GetEventUnits

`func (o *WorkspaceAnalyticsItem) GetEventUnits() float32`

GetEventUnits returns the EventUnits field if non-nil, zero value otherwise.

### GetEventUnitsOk

`func (o *WorkspaceAnalyticsItem) GetEventUnitsOk() (*float32, bool)`

GetEventUnitsOk returns a tuple with the EventUnits field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEventUnits

`func (o *WorkspaceAnalyticsItem) SetEventUnits(v float32)`

SetEventUnits sets EventUnits field to given value.

### HasEventUnits

`func (o *WorkspaceAnalyticsItem) HasEventUnits() bool`

HasEventUnits returns a boolean if a field has been set.

### GetLastActivity

`func (o *WorkspaceAnalyticsItem) GetLastActivity() string`

GetLastActivity returns the LastActivity field if non-nil, zero value otherwise.

### GetLastActivityOk

`func (o *WorkspaceAnalyticsItem) GetLastActivityOk() (*string, bool)`

GetLastActivityOk returns a tuple with the LastActivity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastActivity

`func (o *WorkspaceAnalyticsItem) SetLastActivity(v string)`

SetLastActivity sets LastActivity field to given value.

### HasLastActivity

`func (o *WorkspaceAnalyticsItem) HasLastActivity() bool`

HasLastActivity returns a boolean if a field has been set.

### SetLastActivityNil

`func (o *WorkspaceAnalyticsItem) SetLastActivityNil(b bool)`

 SetLastActivityNil sets the value for LastActivity to be an explicit nil

### UnsetLastActivity
`func (o *WorkspaceAnalyticsItem) UnsetLastActivity()`

UnsetLastActivity ensures that no value is present for LastActivity, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


