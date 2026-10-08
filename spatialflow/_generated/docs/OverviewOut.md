# OverviewOut

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Badges** | [**BadgeCountsOut**](BadgeCountsOut.md) |  | 
**Incidents** | [**[]IncidentOut**](IncidentOut.md) |  | 
**DashboardStats** | [**OverviewDashboardStatsOut**](OverviewDashboardStatsOut.md) |  | 
**DashboardMetrics** | [**OverviewDashboardMetricsOut**](OverviewDashboardMetricsOut.md) |  | 
**DlqStats** | Pointer to [**NullableOverviewDlqStatsOut**](OverviewDlqStatsOut.md) |  | [optional] 
**WorkflowStats** | Pointer to [**NullableOverviewWorkflowStatsOut**](OverviewWorkflowStatsOut.md) |  | [optional] 
**RecentEvents** | [**[]OverviewEventOut**](OverviewEventOut.md) |  | 

## Methods

### NewOverviewOut

`func NewOverviewOut(badges BadgeCountsOut, incidents []IncidentOut, dashboardStats OverviewDashboardStatsOut, dashboardMetrics OverviewDashboardMetricsOut, recentEvents []OverviewEventOut, ) *OverviewOut`

NewOverviewOut instantiates a new OverviewOut object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOverviewOutWithDefaults

`func NewOverviewOutWithDefaults() *OverviewOut`

NewOverviewOutWithDefaults instantiates a new OverviewOut object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBadges

`func (o *OverviewOut) GetBadges() BadgeCountsOut`

GetBadges returns the Badges field if non-nil, zero value otherwise.

### GetBadgesOk

`func (o *OverviewOut) GetBadgesOk() (*BadgeCountsOut, bool)`

GetBadgesOk returns a tuple with the Badges field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBadges

`func (o *OverviewOut) SetBadges(v BadgeCountsOut)`

SetBadges sets Badges field to given value.


### GetIncidents

`func (o *OverviewOut) GetIncidents() []IncidentOut`

GetIncidents returns the Incidents field if non-nil, zero value otherwise.

### GetIncidentsOk

`func (o *OverviewOut) GetIncidentsOk() (*[]IncidentOut, bool)`

GetIncidentsOk returns a tuple with the Incidents field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncidents

`func (o *OverviewOut) SetIncidents(v []IncidentOut)`

SetIncidents sets Incidents field to given value.


### GetDashboardStats

`func (o *OverviewOut) GetDashboardStats() OverviewDashboardStatsOut`

GetDashboardStats returns the DashboardStats field if non-nil, zero value otherwise.

### GetDashboardStatsOk

`func (o *OverviewOut) GetDashboardStatsOk() (*OverviewDashboardStatsOut, bool)`

GetDashboardStatsOk returns a tuple with the DashboardStats field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDashboardStats

`func (o *OverviewOut) SetDashboardStats(v OverviewDashboardStatsOut)`

SetDashboardStats sets DashboardStats field to given value.


### GetDashboardMetrics

`func (o *OverviewOut) GetDashboardMetrics() OverviewDashboardMetricsOut`

GetDashboardMetrics returns the DashboardMetrics field if non-nil, zero value otherwise.

### GetDashboardMetricsOk

`func (o *OverviewOut) GetDashboardMetricsOk() (*OverviewDashboardMetricsOut, bool)`

GetDashboardMetricsOk returns a tuple with the DashboardMetrics field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDashboardMetrics

`func (o *OverviewOut) SetDashboardMetrics(v OverviewDashboardMetricsOut)`

SetDashboardMetrics sets DashboardMetrics field to given value.


### GetDlqStats

`func (o *OverviewOut) GetDlqStats() OverviewDlqStatsOut`

GetDlqStats returns the DlqStats field if non-nil, zero value otherwise.

### GetDlqStatsOk

`func (o *OverviewOut) GetDlqStatsOk() (*OverviewDlqStatsOut, bool)`

GetDlqStatsOk returns a tuple with the DlqStats field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDlqStats

`func (o *OverviewOut) SetDlqStats(v OverviewDlqStatsOut)`

SetDlqStats sets DlqStats field to given value.

### HasDlqStats

`func (o *OverviewOut) HasDlqStats() bool`

HasDlqStats returns a boolean if a field has been set.

### SetDlqStatsNil

`func (o *OverviewOut) SetDlqStatsNil(b bool)`

 SetDlqStatsNil sets the value for DlqStats to be an explicit nil

### UnsetDlqStats
`func (o *OverviewOut) UnsetDlqStats()`

UnsetDlqStats ensures that no value is present for DlqStats, not even an explicit nil
### GetWorkflowStats

`func (o *OverviewOut) GetWorkflowStats() OverviewWorkflowStatsOut`

GetWorkflowStats returns the WorkflowStats field if non-nil, zero value otherwise.

### GetWorkflowStatsOk

`func (o *OverviewOut) GetWorkflowStatsOk() (*OverviewWorkflowStatsOut, bool)`

GetWorkflowStatsOk returns a tuple with the WorkflowStats field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWorkflowStats

`func (o *OverviewOut) SetWorkflowStats(v OverviewWorkflowStatsOut)`

SetWorkflowStats sets WorkflowStats field to given value.

### HasWorkflowStats

`func (o *OverviewOut) HasWorkflowStats() bool`

HasWorkflowStats returns a boolean if a field has been set.

### SetWorkflowStatsNil

`func (o *OverviewOut) SetWorkflowStatsNil(b bool)`

 SetWorkflowStatsNil sets the value for WorkflowStats to be an explicit nil

### UnsetWorkflowStats
`func (o *OverviewOut) UnsetWorkflowStats()`

UnsetWorkflowStats ensures that no value is present for WorkflowStats, not even an explicit nil
### GetRecentEvents

`func (o *OverviewOut) GetRecentEvents() []OverviewEventOut`

GetRecentEvents returns the RecentEvents field if non-nil, zero value otherwise.

### GetRecentEventsOk

`func (o *OverviewOut) GetRecentEventsOk() (*[]OverviewEventOut, bool)`

GetRecentEventsOk returns a tuple with the RecentEvents field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecentEvents

`func (o *OverviewOut) SetRecentEvents(v []OverviewEventOut)`

SetRecentEvents sets RecentEvents field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


