# DashboardStatsOut

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ActiveDeviceCount** | **int32** |  | 
**LiveCount** | **int32** |  | 
**ReportingCount** | **int32** |  | 
**OfflineStaleCount** | **int32** |  | 
**ExpectedReportingCount** | **int32** |  | 
**AttentionCount** | **int32** |  | 
**ParkedCount** | **int32** |  | 
**ReportingWindowMinutes** | **int32** |  | 
**PausedCount** | **int32** |  | 
**OffShiftCount** | **int32** |  | 
**LowBatteryCount** | **int32** |  | 
**LowBatteryWindowMinutes** | **int32** |  | 
**InGeofenceCount** | **int32** |  | 
**LastKnownGeofenceCount** | **int32** |  | 
**ConfirmedCurrentGeofenceCount** | **int32** |  | 
**AlertsOpen** | **int32** |  | 
**WorkflowFailures1h** | Pointer to **NullableInt32** |  | [optional] 
**WebhookRetries1h** | Pointer to **NullableInt32** |  | [optional] 

## Methods

### NewDashboardStatsOut

`func NewDashboardStatsOut(activeDeviceCount int32, liveCount int32, reportingCount int32, offlineStaleCount int32, expectedReportingCount int32, attentionCount int32, parkedCount int32, reportingWindowMinutes int32, pausedCount int32, offShiftCount int32, lowBatteryCount int32, lowBatteryWindowMinutes int32, inGeofenceCount int32, lastKnownGeofenceCount int32, confirmedCurrentGeofenceCount int32, alertsOpen int32, ) *DashboardStatsOut`

NewDashboardStatsOut instantiates a new DashboardStatsOut object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDashboardStatsOutWithDefaults

`func NewDashboardStatsOutWithDefaults() *DashboardStatsOut`

NewDashboardStatsOutWithDefaults instantiates a new DashboardStatsOut object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetActiveDeviceCount

`func (o *DashboardStatsOut) GetActiveDeviceCount() int32`

GetActiveDeviceCount returns the ActiveDeviceCount field if non-nil, zero value otherwise.

### GetActiveDeviceCountOk

`func (o *DashboardStatsOut) GetActiveDeviceCountOk() (*int32, bool)`

GetActiveDeviceCountOk returns a tuple with the ActiveDeviceCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActiveDeviceCount

`func (o *DashboardStatsOut) SetActiveDeviceCount(v int32)`

SetActiveDeviceCount sets ActiveDeviceCount field to given value.


### GetLiveCount

`func (o *DashboardStatsOut) GetLiveCount() int32`

GetLiveCount returns the LiveCount field if non-nil, zero value otherwise.

### GetLiveCountOk

`func (o *DashboardStatsOut) GetLiveCountOk() (*int32, bool)`

GetLiveCountOk returns a tuple with the LiveCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLiveCount

`func (o *DashboardStatsOut) SetLiveCount(v int32)`

SetLiveCount sets LiveCount field to given value.


### GetReportingCount

`func (o *DashboardStatsOut) GetReportingCount() int32`

GetReportingCount returns the ReportingCount field if non-nil, zero value otherwise.

### GetReportingCountOk

`func (o *DashboardStatsOut) GetReportingCountOk() (*int32, bool)`

GetReportingCountOk returns a tuple with the ReportingCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReportingCount

`func (o *DashboardStatsOut) SetReportingCount(v int32)`

SetReportingCount sets ReportingCount field to given value.


### GetOfflineStaleCount

`func (o *DashboardStatsOut) GetOfflineStaleCount() int32`

GetOfflineStaleCount returns the OfflineStaleCount field if non-nil, zero value otherwise.

### GetOfflineStaleCountOk

`func (o *DashboardStatsOut) GetOfflineStaleCountOk() (*int32, bool)`

GetOfflineStaleCountOk returns a tuple with the OfflineStaleCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOfflineStaleCount

`func (o *DashboardStatsOut) SetOfflineStaleCount(v int32)`

SetOfflineStaleCount sets OfflineStaleCount field to given value.


### GetExpectedReportingCount

`func (o *DashboardStatsOut) GetExpectedReportingCount() int32`

GetExpectedReportingCount returns the ExpectedReportingCount field if non-nil, zero value otherwise.

### GetExpectedReportingCountOk

`func (o *DashboardStatsOut) GetExpectedReportingCountOk() (*int32, bool)`

GetExpectedReportingCountOk returns a tuple with the ExpectedReportingCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpectedReportingCount

`func (o *DashboardStatsOut) SetExpectedReportingCount(v int32)`

SetExpectedReportingCount sets ExpectedReportingCount field to given value.


### GetAttentionCount

`func (o *DashboardStatsOut) GetAttentionCount() int32`

GetAttentionCount returns the AttentionCount field if non-nil, zero value otherwise.

### GetAttentionCountOk

`func (o *DashboardStatsOut) GetAttentionCountOk() (*int32, bool)`

GetAttentionCountOk returns a tuple with the AttentionCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttentionCount

`func (o *DashboardStatsOut) SetAttentionCount(v int32)`

SetAttentionCount sets AttentionCount field to given value.


### GetParkedCount

`func (o *DashboardStatsOut) GetParkedCount() int32`

GetParkedCount returns the ParkedCount field if non-nil, zero value otherwise.

### GetParkedCountOk

`func (o *DashboardStatsOut) GetParkedCountOk() (*int32, bool)`

GetParkedCountOk returns a tuple with the ParkedCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParkedCount

`func (o *DashboardStatsOut) SetParkedCount(v int32)`

SetParkedCount sets ParkedCount field to given value.


### GetReportingWindowMinutes

`func (o *DashboardStatsOut) GetReportingWindowMinutes() int32`

GetReportingWindowMinutes returns the ReportingWindowMinutes field if non-nil, zero value otherwise.

### GetReportingWindowMinutesOk

`func (o *DashboardStatsOut) GetReportingWindowMinutesOk() (*int32, bool)`

GetReportingWindowMinutesOk returns a tuple with the ReportingWindowMinutes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReportingWindowMinutes

`func (o *DashboardStatsOut) SetReportingWindowMinutes(v int32)`

SetReportingWindowMinutes sets ReportingWindowMinutes field to given value.


### GetPausedCount

`func (o *DashboardStatsOut) GetPausedCount() int32`

GetPausedCount returns the PausedCount field if non-nil, zero value otherwise.

### GetPausedCountOk

`func (o *DashboardStatsOut) GetPausedCountOk() (*int32, bool)`

GetPausedCountOk returns a tuple with the PausedCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPausedCount

`func (o *DashboardStatsOut) SetPausedCount(v int32)`

SetPausedCount sets PausedCount field to given value.


### GetOffShiftCount

`func (o *DashboardStatsOut) GetOffShiftCount() int32`

GetOffShiftCount returns the OffShiftCount field if non-nil, zero value otherwise.

### GetOffShiftCountOk

`func (o *DashboardStatsOut) GetOffShiftCountOk() (*int32, bool)`

GetOffShiftCountOk returns a tuple with the OffShiftCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOffShiftCount

`func (o *DashboardStatsOut) SetOffShiftCount(v int32)`

SetOffShiftCount sets OffShiftCount field to given value.


### GetLowBatteryCount

`func (o *DashboardStatsOut) GetLowBatteryCount() int32`

GetLowBatteryCount returns the LowBatteryCount field if non-nil, zero value otherwise.

### GetLowBatteryCountOk

`func (o *DashboardStatsOut) GetLowBatteryCountOk() (*int32, bool)`

GetLowBatteryCountOk returns a tuple with the LowBatteryCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLowBatteryCount

`func (o *DashboardStatsOut) SetLowBatteryCount(v int32)`

SetLowBatteryCount sets LowBatteryCount field to given value.


### GetLowBatteryWindowMinutes

`func (o *DashboardStatsOut) GetLowBatteryWindowMinutes() int32`

GetLowBatteryWindowMinutes returns the LowBatteryWindowMinutes field if non-nil, zero value otherwise.

### GetLowBatteryWindowMinutesOk

`func (o *DashboardStatsOut) GetLowBatteryWindowMinutesOk() (*int32, bool)`

GetLowBatteryWindowMinutesOk returns a tuple with the LowBatteryWindowMinutes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLowBatteryWindowMinutes

`func (o *DashboardStatsOut) SetLowBatteryWindowMinutes(v int32)`

SetLowBatteryWindowMinutes sets LowBatteryWindowMinutes field to given value.


### GetInGeofenceCount

`func (o *DashboardStatsOut) GetInGeofenceCount() int32`

GetInGeofenceCount returns the InGeofenceCount field if non-nil, zero value otherwise.

### GetInGeofenceCountOk

`func (o *DashboardStatsOut) GetInGeofenceCountOk() (*int32, bool)`

GetInGeofenceCountOk returns a tuple with the InGeofenceCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInGeofenceCount

`func (o *DashboardStatsOut) SetInGeofenceCount(v int32)`

SetInGeofenceCount sets InGeofenceCount field to given value.


### GetLastKnownGeofenceCount

`func (o *DashboardStatsOut) GetLastKnownGeofenceCount() int32`

GetLastKnownGeofenceCount returns the LastKnownGeofenceCount field if non-nil, zero value otherwise.

### GetLastKnownGeofenceCountOk

`func (o *DashboardStatsOut) GetLastKnownGeofenceCountOk() (*int32, bool)`

GetLastKnownGeofenceCountOk returns a tuple with the LastKnownGeofenceCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastKnownGeofenceCount

`func (o *DashboardStatsOut) SetLastKnownGeofenceCount(v int32)`

SetLastKnownGeofenceCount sets LastKnownGeofenceCount field to given value.


### GetConfirmedCurrentGeofenceCount

`func (o *DashboardStatsOut) GetConfirmedCurrentGeofenceCount() int32`

GetConfirmedCurrentGeofenceCount returns the ConfirmedCurrentGeofenceCount field if non-nil, zero value otherwise.

### GetConfirmedCurrentGeofenceCountOk

`func (o *DashboardStatsOut) GetConfirmedCurrentGeofenceCountOk() (*int32, bool)`

GetConfirmedCurrentGeofenceCountOk returns a tuple with the ConfirmedCurrentGeofenceCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfirmedCurrentGeofenceCount

`func (o *DashboardStatsOut) SetConfirmedCurrentGeofenceCount(v int32)`

SetConfirmedCurrentGeofenceCount sets ConfirmedCurrentGeofenceCount field to given value.


### GetAlertsOpen

`func (o *DashboardStatsOut) GetAlertsOpen() int32`

GetAlertsOpen returns the AlertsOpen field if non-nil, zero value otherwise.

### GetAlertsOpenOk

`func (o *DashboardStatsOut) GetAlertsOpenOk() (*int32, bool)`

GetAlertsOpenOk returns a tuple with the AlertsOpen field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAlertsOpen

`func (o *DashboardStatsOut) SetAlertsOpen(v int32)`

SetAlertsOpen sets AlertsOpen field to given value.


### GetWorkflowFailures1h

`func (o *DashboardStatsOut) GetWorkflowFailures1h() int32`

GetWorkflowFailures1h returns the WorkflowFailures1h field if non-nil, zero value otherwise.

### GetWorkflowFailures1hOk

`func (o *DashboardStatsOut) GetWorkflowFailures1hOk() (*int32, bool)`

GetWorkflowFailures1hOk returns a tuple with the WorkflowFailures1h field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWorkflowFailures1h

`func (o *DashboardStatsOut) SetWorkflowFailures1h(v int32)`

SetWorkflowFailures1h sets WorkflowFailures1h field to given value.

### HasWorkflowFailures1h

`func (o *DashboardStatsOut) HasWorkflowFailures1h() bool`

HasWorkflowFailures1h returns a boolean if a field has been set.

### SetWorkflowFailures1hNil

`func (o *DashboardStatsOut) SetWorkflowFailures1hNil(b bool)`

 SetWorkflowFailures1hNil sets the value for WorkflowFailures1h to be an explicit nil

### UnsetWorkflowFailures1h
`func (o *DashboardStatsOut) UnsetWorkflowFailures1h()`

UnsetWorkflowFailures1h ensures that no value is present for WorkflowFailures1h, not even an explicit nil
### GetWebhookRetries1h

`func (o *DashboardStatsOut) GetWebhookRetries1h() int32`

GetWebhookRetries1h returns the WebhookRetries1h field if non-nil, zero value otherwise.

### GetWebhookRetries1hOk

`func (o *DashboardStatsOut) GetWebhookRetries1hOk() (*int32, bool)`

GetWebhookRetries1hOk returns a tuple with the WebhookRetries1h field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWebhookRetries1h

`func (o *DashboardStatsOut) SetWebhookRetries1h(v int32)`

SetWebhookRetries1h sets WebhookRetries1h field to given value.

### HasWebhookRetries1h

`func (o *DashboardStatsOut) HasWebhookRetries1h() bool`

HasWebhookRetries1h returns a boolean if a field has been set.

### SetWebhookRetries1hNil

`func (o *DashboardStatsOut) SetWebhookRetries1hNil(b bool)`

 SetWebhookRetries1hNil sets the value for WebhookRetries1h to be an explicit nil

### UnsetWebhookRetries1h
`func (o *DashboardStatsOut) UnsetWebhookRetries1h()`

UnsetWebhookRetries1h ensures that no value is present for WebhookRetries1h, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


