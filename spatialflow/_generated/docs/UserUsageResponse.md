# UserUsageResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**UserId** | **string** |  | 
**Email** | **string** |  | 
**Workspace** | Pointer to [**NullableWorkspaceDetail**](WorkspaceDetail.md) |  | [optional] 
**ApiUsage** | [**APIUsageStats**](APIUsageStats.md) |  | 
**EmailStats** | [**EmailStats**](EmailStats.md) |  | 
**GeofenceStats** | [**GeofenceStats**](GeofenceStats.md) |  | 
**ActivitySummary** | [**[]ActivitySummary**](ActivitySummary.md) |  | 
**RecentActivities** | [**[]RecentActivity**](RecentActivity.md) |  | 
**AccountCreated** | **NullableString** |  | 
**LastLogin** | **NullableString** |  | 

## Methods

### NewUserUsageResponse

`func NewUserUsageResponse(userId string, email string, apiUsage APIUsageStats, emailStats EmailStats, geofenceStats GeofenceStats, activitySummary []ActivitySummary, recentActivities []RecentActivity, accountCreated NullableString, lastLogin NullableString, ) *UserUsageResponse`

NewUserUsageResponse instantiates a new UserUsageResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUserUsageResponseWithDefaults

`func NewUserUsageResponseWithDefaults() *UserUsageResponse`

NewUserUsageResponseWithDefaults instantiates a new UserUsageResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUserId

`func (o *UserUsageResponse) GetUserId() string`

GetUserId returns the UserId field if non-nil, zero value otherwise.

### GetUserIdOk

`func (o *UserUsageResponse) GetUserIdOk() (*string, bool)`

GetUserIdOk returns a tuple with the UserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserId

`func (o *UserUsageResponse) SetUserId(v string)`

SetUserId sets UserId field to given value.


### GetEmail

`func (o *UserUsageResponse) GetEmail() string`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *UserUsageResponse) GetEmailOk() (*string, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *UserUsageResponse) SetEmail(v string)`

SetEmail sets Email field to given value.


### GetWorkspace

`func (o *UserUsageResponse) GetWorkspace() WorkspaceDetail`

GetWorkspace returns the Workspace field if non-nil, zero value otherwise.

### GetWorkspaceOk

`func (o *UserUsageResponse) GetWorkspaceOk() (*WorkspaceDetail, bool)`

GetWorkspaceOk returns a tuple with the Workspace field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWorkspace

`func (o *UserUsageResponse) SetWorkspace(v WorkspaceDetail)`

SetWorkspace sets Workspace field to given value.

### HasWorkspace

`func (o *UserUsageResponse) HasWorkspace() bool`

HasWorkspace returns a boolean if a field has been set.

### SetWorkspaceNil

`func (o *UserUsageResponse) SetWorkspaceNil(b bool)`

 SetWorkspaceNil sets the value for Workspace to be an explicit nil

### UnsetWorkspace
`func (o *UserUsageResponse) UnsetWorkspace()`

UnsetWorkspace ensures that no value is present for Workspace, not even an explicit nil
### GetApiUsage

`func (o *UserUsageResponse) GetApiUsage() APIUsageStats`

GetApiUsage returns the ApiUsage field if non-nil, zero value otherwise.

### GetApiUsageOk

`func (o *UserUsageResponse) GetApiUsageOk() (*APIUsageStats, bool)`

GetApiUsageOk returns a tuple with the ApiUsage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApiUsage

`func (o *UserUsageResponse) SetApiUsage(v APIUsageStats)`

SetApiUsage sets ApiUsage field to given value.


### GetEmailStats

`func (o *UserUsageResponse) GetEmailStats() EmailStats`

GetEmailStats returns the EmailStats field if non-nil, zero value otherwise.

### GetEmailStatsOk

`func (o *UserUsageResponse) GetEmailStatsOk() (*EmailStats, bool)`

GetEmailStatsOk returns a tuple with the EmailStats field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmailStats

`func (o *UserUsageResponse) SetEmailStats(v EmailStats)`

SetEmailStats sets EmailStats field to given value.


### GetGeofenceStats

`func (o *UserUsageResponse) GetGeofenceStats() GeofenceStats`

GetGeofenceStats returns the GeofenceStats field if non-nil, zero value otherwise.

### GetGeofenceStatsOk

`func (o *UserUsageResponse) GetGeofenceStatsOk() (*GeofenceStats, bool)`

GetGeofenceStatsOk returns a tuple with the GeofenceStats field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGeofenceStats

`func (o *UserUsageResponse) SetGeofenceStats(v GeofenceStats)`

SetGeofenceStats sets GeofenceStats field to given value.


### GetActivitySummary

`func (o *UserUsageResponse) GetActivitySummary() []ActivitySummary`

GetActivitySummary returns the ActivitySummary field if non-nil, zero value otherwise.

### GetActivitySummaryOk

`func (o *UserUsageResponse) GetActivitySummaryOk() (*[]ActivitySummary, bool)`

GetActivitySummaryOk returns a tuple with the ActivitySummary field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActivitySummary

`func (o *UserUsageResponse) SetActivitySummary(v []ActivitySummary)`

SetActivitySummary sets ActivitySummary field to given value.


### GetRecentActivities

`func (o *UserUsageResponse) GetRecentActivities() []RecentActivity`

GetRecentActivities returns the RecentActivities field if non-nil, zero value otherwise.

### GetRecentActivitiesOk

`func (o *UserUsageResponse) GetRecentActivitiesOk() (*[]RecentActivity, bool)`

GetRecentActivitiesOk returns a tuple with the RecentActivities field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecentActivities

`func (o *UserUsageResponse) SetRecentActivities(v []RecentActivity)`

SetRecentActivities sets RecentActivities field to given value.


### GetAccountCreated

`func (o *UserUsageResponse) GetAccountCreated() string`

GetAccountCreated returns the AccountCreated field if non-nil, zero value otherwise.

### GetAccountCreatedOk

`func (o *UserUsageResponse) GetAccountCreatedOk() (*string, bool)`

GetAccountCreatedOk returns a tuple with the AccountCreated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountCreated

`func (o *UserUsageResponse) SetAccountCreated(v string)`

SetAccountCreated sets AccountCreated field to given value.


### SetAccountCreatedNil

`func (o *UserUsageResponse) SetAccountCreatedNil(b bool)`

 SetAccountCreatedNil sets the value for AccountCreated to be an explicit nil

### UnsetAccountCreated
`func (o *UserUsageResponse) UnsetAccountCreated()`

UnsetAccountCreated ensures that no value is present for AccountCreated, not even an explicit nil
### GetLastLogin

`func (o *UserUsageResponse) GetLastLogin() string`

GetLastLogin returns the LastLogin field if non-nil, zero value otherwise.

### GetLastLoginOk

`func (o *UserUsageResponse) GetLastLoginOk() (*string, bool)`

GetLastLoginOk returns a tuple with the LastLogin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastLogin

`func (o *UserUsageResponse) SetLastLogin(v string)`

SetLastLogin sets LastLogin field to given value.


### SetLastLoginNil

`func (o *UserUsageResponse) SetLastLoginNil(b bool)`

 SetLastLoginNil sets the value for LastLogin to be an explicit nil

### UnsetLastLogin
`func (o *UserUsageResponse) UnsetLastLogin()`

UnsetLastLogin ensures that no value is present for LastLogin, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


