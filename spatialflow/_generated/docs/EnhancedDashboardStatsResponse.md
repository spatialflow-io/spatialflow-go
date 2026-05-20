# EnhancedDashboardStatsResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**TotalUsers** | **int32** |  | 
**ActiveUsers** | **int32** |  | 
**NewUsersThisWeek** | **int32** |  | 
**TotalWorkspaces** | **int32** |  | 
**PaidSubscriptions** | Pointer to **int32** |  | [optional] [default to 0]
**TrialSubscriptions** | Pointer to **int32** |  | [optional] [default to 0]
**CanceledSubscriptions** | Pointer to **int32** |  | [optional] [default to 0]
**MonthlyRecurringRevenue** | Pointer to **float32** |  | [optional] [default to 0.0]

## Methods

### NewEnhancedDashboardStatsResponse

`func NewEnhancedDashboardStatsResponse(totalUsers int32, activeUsers int32, newUsersThisWeek int32, totalWorkspaces int32, ) *EnhancedDashboardStatsResponse`

NewEnhancedDashboardStatsResponse instantiates a new EnhancedDashboardStatsResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEnhancedDashboardStatsResponseWithDefaults

`func NewEnhancedDashboardStatsResponseWithDefaults() *EnhancedDashboardStatsResponse`

NewEnhancedDashboardStatsResponseWithDefaults instantiates a new EnhancedDashboardStatsResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTotalUsers

`func (o *EnhancedDashboardStatsResponse) GetTotalUsers() int32`

GetTotalUsers returns the TotalUsers field if non-nil, zero value otherwise.

### GetTotalUsersOk

`func (o *EnhancedDashboardStatsResponse) GetTotalUsersOk() (*int32, bool)`

GetTotalUsersOk returns a tuple with the TotalUsers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalUsers

`func (o *EnhancedDashboardStatsResponse) SetTotalUsers(v int32)`

SetTotalUsers sets TotalUsers field to given value.


### GetActiveUsers

`func (o *EnhancedDashboardStatsResponse) GetActiveUsers() int32`

GetActiveUsers returns the ActiveUsers field if non-nil, zero value otherwise.

### GetActiveUsersOk

`func (o *EnhancedDashboardStatsResponse) GetActiveUsersOk() (*int32, bool)`

GetActiveUsersOk returns a tuple with the ActiveUsers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActiveUsers

`func (o *EnhancedDashboardStatsResponse) SetActiveUsers(v int32)`

SetActiveUsers sets ActiveUsers field to given value.


### GetNewUsersThisWeek

`func (o *EnhancedDashboardStatsResponse) GetNewUsersThisWeek() int32`

GetNewUsersThisWeek returns the NewUsersThisWeek field if non-nil, zero value otherwise.

### GetNewUsersThisWeekOk

`func (o *EnhancedDashboardStatsResponse) GetNewUsersThisWeekOk() (*int32, bool)`

GetNewUsersThisWeekOk returns a tuple with the NewUsersThisWeek field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNewUsersThisWeek

`func (o *EnhancedDashboardStatsResponse) SetNewUsersThisWeek(v int32)`

SetNewUsersThisWeek sets NewUsersThisWeek field to given value.


### GetTotalWorkspaces

`func (o *EnhancedDashboardStatsResponse) GetTotalWorkspaces() int32`

GetTotalWorkspaces returns the TotalWorkspaces field if non-nil, zero value otherwise.

### GetTotalWorkspacesOk

`func (o *EnhancedDashboardStatsResponse) GetTotalWorkspacesOk() (*int32, bool)`

GetTotalWorkspacesOk returns a tuple with the TotalWorkspaces field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalWorkspaces

`func (o *EnhancedDashboardStatsResponse) SetTotalWorkspaces(v int32)`

SetTotalWorkspaces sets TotalWorkspaces field to given value.


### GetPaidSubscriptions

`func (o *EnhancedDashboardStatsResponse) GetPaidSubscriptions() int32`

GetPaidSubscriptions returns the PaidSubscriptions field if non-nil, zero value otherwise.

### GetPaidSubscriptionsOk

`func (o *EnhancedDashboardStatsResponse) GetPaidSubscriptionsOk() (*int32, bool)`

GetPaidSubscriptionsOk returns a tuple with the PaidSubscriptions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaidSubscriptions

`func (o *EnhancedDashboardStatsResponse) SetPaidSubscriptions(v int32)`

SetPaidSubscriptions sets PaidSubscriptions field to given value.

### HasPaidSubscriptions

`func (o *EnhancedDashboardStatsResponse) HasPaidSubscriptions() bool`

HasPaidSubscriptions returns a boolean if a field has been set.

### GetTrialSubscriptions

`func (o *EnhancedDashboardStatsResponse) GetTrialSubscriptions() int32`

GetTrialSubscriptions returns the TrialSubscriptions field if non-nil, zero value otherwise.

### GetTrialSubscriptionsOk

`func (o *EnhancedDashboardStatsResponse) GetTrialSubscriptionsOk() (*int32, bool)`

GetTrialSubscriptionsOk returns a tuple with the TrialSubscriptions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTrialSubscriptions

`func (o *EnhancedDashboardStatsResponse) SetTrialSubscriptions(v int32)`

SetTrialSubscriptions sets TrialSubscriptions field to given value.

### HasTrialSubscriptions

`func (o *EnhancedDashboardStatsResponse) HasTrialSubscriptions() bool`

HasTrialSubscriptions returns a boolean if a field has been set.

### GetCanceledSubscriptions

`func (o *EnhancedDashboardStatsResponse) GetCanceledSubscriptions() int32`

GetCanceledSubscriptions returns the CanceledSubscriptions field if non-nil, zero value otherwise.

### GetCanceledSubscriptionsOk

`func (o *EnhancedDashboardStatsResponse) GetCanceledSubscriptionsOk() (*int32, bool)`

GetCanceledSubscriptionsOk returns a tuple with the CanceledSubscriptions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCanceledSubscriptions

`func (o *EnhancedDashboardStatsResponse) SetCanceledSubscriptions(v int32)`

SetCanceledSubscriptions sets CanceledSubscriptions field to given value.

### HasCanceledSubscriptions

`func (o *EnhancedDashboardStatsResponse) HasCanceledSubscriptions() bool`

HasCanceledSubscriptions returns a boolean if a field has been set.

### GetMonthlyRecurringRevenue

`func (o *EnhancedDashboardStatsResponse) GetMonthlyRecurringRevenue() float32`

GetMonthlyRecurringRevenue returns the MonthlyRecurringRevenue field if non-nil, zero value otherwise.

### GetMonthlyRecurringRevenueOk

`func (o *EnhancedDashboardStatsResponse) GetMonthlyRecurringRevenueOk() (*float32, bool)`

GetMonthlyRecurringRevenueOk returns a tuple with the MonthlyRecurringRevenue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMonthlyRecurringRevenue

`func (o *EnhancedDashboardStatsResponse) SetMonthlyRecurringRevenue(v float32)`

SetMonthlyRecurringRevenue sets MonthlyRecurringRevenue field to given value.

### HasMonthlyRecurringRevenue

`func (o *EnhancedDashboardStatsResponse) HasMonthlyRecurringRevenue() bool`

HasMonthlyRecurringRevenue returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


