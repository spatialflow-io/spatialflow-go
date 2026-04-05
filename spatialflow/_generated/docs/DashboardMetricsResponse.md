# DashboardMetricsResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**TimeRange** | **string** | Time range: today, 7d, 30d, or custom | 
**PeriodStart** | **time.Time** | Start of the period | 
**PeriodEnd** | **time.Time** | End of the period | 
**ActiveWorkflows** | **int32** | Number of active workflows in period | 
**EventsTotal** | **int32** | Total geofence events (entry/exit) in period | 
**ActionDeliverySuccess** | [**ActionDeliverySuccessMetrics**](ActionDeliverySuccessMetrics.md) | Action delivery success metrics (north-star KPI) | 
**Comparison** | [**DashboardComparisonMetrics**](DashboardComparisonMetrics.md) | Comparison vs previous period | 

## Methods

### NewDashboardMetricsResponse

`func NewDashboardMetricsResponse(timeRange string, periodStart time.Time, periodEnd time.Time, activeWorkflows int32, eventsTotal int32, actionDeliverySuccess ActionDeliverySuccessMetrics, comparison DashboardComparisonMetrics, ) *DashboardMetricsResponse`

NewDashboardMetricsResponse instantiates a new DashboardMetricsResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDashboardMetricsResponseWithDefaults

`func NewDashboardMetricsResponseWithDefaults() *DashboardMetricsResponse`

NewDashboardMetricsResponseWithDefaults instantiates a new DashboardMetricsResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTimeRange

`func (o *DashboardMetricsResponse) GetTimeRange() string`

GetTimeRange returns the TimeRange field if non-nil, zero value otherwise.

### GetTimeRangeOk

`func (o *DashboardMetricsResponse) GetTimeRangeOk() (*string, bool)`

GetTimeRangeOk returns a tuple with the TimeRange field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeRange

`func (o *DashboardMetricsResponse) SetTimeRange(v string)`

SetTimeRange sets TimeRange field to given value.


### GetPeriodStart

`func (o *DashboardMetricsResponse) GetPeriodStart() time.Time`

GetPeriodStart returns the PeriodStart field if non-nil, zero value otherwise.

### GetPeriodStartOk

`func (o *DashboardMetricsResponse) GetPeriodStartOk() (*time.Time, bool)`

GetPeriodStartOk returns a tuple with the PeriodStart field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPeriodStart

`func (o *DashboardMetricsResponse) SetPeriodStart(v time.Time)`

SetPeriodStart sets PeriodStart field to given value.


### GetPeriodEnd

`func (o *DashboardMetricsResponse) GetPeriodEnd() time.Time`

GetPeriodEnd returns the PeriodEnd field if non-nil, zero value otherwise.

### GetPeriodEndOk

`func (o *DashboardMetricsResponse) GetPeriodEndOk() (*time.Time, bool)`

GetPeriodEndOk returns a tuple with the PeriodEnd field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPeriodEnd

`func (o *DashboardMetricsResponse) SetPeriodEnd(v time.Time)`

SetPeriodEnd sets PeriodEnd field to given value.


### GetActiveWorkflows

`func (o *DashboardMetricsResponse) GetActiveWorkflows() int32`

GetActiveWorkflows returns the ActiveWorkflows field if non-nil, zero value otherwise.

### GetActiveWorkflowsOk

`func (o *DashboardMetricsResponse) GetActiveWorkflowsOk() (*int32, bool)`

GetActiveWorkflowsOk returns a tuple with the ActiveWorkflows field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActiveWorkflows

`func (o *DashboardMetricsResponse) SetActiveWorkflows(v int32)`

SetActiveWorkflows sets ActiveWorkflows field to given value.


### GetEventsTotal

`func (o *DashboardMetricsResponse) GetEventsTotal() int32`

GetEventsTotal returns the EventsTotal field if non-nil, zero value otherwise.

### GetEventsTotalOk

`func (o *DashboardMetricsResponse) GetEventsTotalOk() (*int32, bool)`

GetEventsTotalOk returns a tuple with the EventsTotal field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEventsTotal

`func (o *DashboardMetricsResponse) SetEventsTotal(v int32)`

SetEventsTotal sets EventsTotal field to given value.


### GetActionDeliverySuccess

`func (o *DashboardMetricsResponse) GetActionDeliverySuccess() ActionDeliverySuccessMetrics`

GetActionDeliverySuccess returns the ActionDeliverySuccess field if non-nil, zero value otherwise.

### GetActionDeliverySuccessOk

`func (o *DashboardMetricsResponse) GetActionDeliverySuccessOk() (*ActionDeliverySuccessMetrics, bool)`

GetActionDeliverySuccessOk returns a tuple with the ActionDeliverySuccess field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActionDeliverySuccess

`func (o *DashboardMetricsResponse) SetActionDeliverySuccess(v ActionDeliverySuccessMetrics)`

SetActionDeliverySuccess sets ActionDeliverySuccess field to given value.


### GetComparison

`func (o *DashboardMetricsResponse) GetComparison() DashboardComparisonMetrics`

GetComparison returns the Comparison field if non-nil, zero value otherwise.

### GetComparisonOk

`func (o *DashboardMetricsResponse) GetComparisonOk() (*DashboardComparisonMetrics, bool)`

GetComparisonOk returns a tuple with the Comparison field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComparison

`func (o *DashboardMetricsResponse) SetComparison(v DashboardComparisonMetrics)`

SetComparison sets Comparison field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


