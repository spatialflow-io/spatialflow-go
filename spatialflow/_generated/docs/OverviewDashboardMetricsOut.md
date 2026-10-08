# OverviewDashboardMetricsOut

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**TimeRange** | **string** |  | 
**PeriodStart** | **time.Time** |  | 
**PeriodEnd** | **time.Time** |  | 
**ActiveWorkflows** | **int32** |  | 
**EventsTotal** | **int32** |  | 
**ActionDeliverySuccess** | [**OverviewActionDeliveryOut**](OverviewActionDeliveryOut.md) |  | 
**Comparison** | [**OverviewComparisonOut**](OverviewComparisonOut.md) |  | 

## Methods

### NewOverviewDashboardMetricsOut

`func NewOverviewDashboardMetricsOut(timeRange string, periodStart time.Time, periodEnd time.Time, activeWorkflows int32, eventsTotal int32, actionDeliverySuccess OverviewActionDeliveryOut, comparison OverviewComparisonOut, ) *OverviewDashboardMetricsOut`

NewOverviewDashboardMetricsOut instantiates a new OverviewDashboardMetricsOut object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOverviewDashboardMetricsOutWithDefaults

`func NewOverviewDashboardMetricsOutWithDefaults() *OverviewDashboardMetricsOut`

NewOverviewDashboardMetricsOutWithDefaults instantiates a new OverviewDashboardMetricsOut object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTimeRange

`func (o *OverviewDashboardMetricsOut) GetTimeRange() string`

GetTimeRange returns the TimeRange field if non-nil, zero value otherwise.

### GetTimeRangeOk

`func (o *OverviewDashboardMetricsOut) GetTimeRangeOk() (*string, bool)`

GetTimeRangeOk returns a tuple with the TimeRange field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeRange

`func (o *OverviewDashboardMetricsOut) SetTimeRange(v string)`

SetTimeRange sets TimeRange field to given value.


### GetPeriodStart

`func (o *OverviewDashboardMetricsOut) GetPeriodStart() time.Time`

GetPeriodStart returns the PeriodStart field if non-nil, zero value otherwise.

### GetPeriodStartOk

`func (o *OverviewDashboardMetricsOut) GetPeriodStartOk() (*time.Time, bool)`

GetPeriodStartOk returns a tuple with the PeriodStart field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPeriodStart

`func (o *OverviewDashboardMetricsOut) SetPeriodStart(v time.Time)`

SetPeriodStart sets PeriodStart field to given value.


### GetPeriodEnd

`func (o *OverviewDashboardMetricsOut) GetPeriodEnd() time.Time`

GetPeriodEnd returns the PeriodEnd field if non-nil, zero value otherwise.

### GetPeriodEndOk

`func (o *OverviewDashboardMetricsOut) GetPeriodEndOk() (*time.Time, bool)`

GetPeriodEndOk returns a tuple with the PeriodEnd field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPeriodEnd

`func (o *OverviewDashboardMetricsOut) SetPeriodEnd(v time.Time)`

SetPeriodEnd sets PeriodEnd field to given value.


### GetActiveWorkflows

`func (o *OverviewDashboardMetricsOut) GetActiveWorkflows() int32`

GetActiveWorkflows returns the ActiveWorkflows field if non-nil, zero value otherwise.

### GetActiveWorkflowsOk

`func (o *OverviewDashboardMetricsOut) GetActiveWorkflowsOk() (*int32, bool)`

GetActiveWorkflowsOk returns a tuple with the ActiveWorkflows field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActiveWorkflows

`func (o *OverviewDashboardMetricsOut) SetActiveWorkflows(v int32)`

SetActiveWorkflows sets ActiveWorkflows field to given value.


### GetEventsTotal

`func (o *OverviewDashboardMetricsOut) GetEventsTotal() int32`

GetEventsTotal returns the EventsTotal field if non-nil, zero value otherwise.

### GetEventsTotalOk

`func (o *OverviewDashboardMetricsOut) GetEventsTotalOk() (*int32, bool)`

GetEventsTotalOk returns a tuple with the EventsTotal field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEventsTotal

`func (o *OverviewDashboardMetricsOut) SetEventsTotal(v int32)`

SetEventsTotal sets EventsTotal field to given value.


### GetActionDeliverySuccess

`func (o *OverviewDashboardMetricsOut) GetActionDeliverySuccess() OverviewActionDeliveryOut`

GetActionDeliverySuccess returns the ActionDeliverySuccess field if non-nil, zero value otherwise.

### GetActionDeliverySuccessOk

`func (o *OverviewDashboardMetricsOut) GetActionDeliverySuccessOk() (*OverviewActionDeliveryOut, bool)`

GetActionDeliverySuccessOk returns a tuple with the ActionDeliverySuccess field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActionDeliverySuccess

`func (o *OverviewDashboardMetricsOut) SetActionDeliverySuccess(v OverviewActionDeliveryOut)`

SetActionDeliverySuccess sets ActionDeliverySuccess field to given value.


### GetComparison

`func (o *OverviewDashboardMetricsOut) GetComparison() OverviewComparisonOut`

GetComparison returns the Comparison field if non-nil, zero value otherwise.

### GetComparisonOk

`func (o *OverviewDashboardMetricsOut) GetComparisonOk() (*OverviewComparisonOut, bool)`

GetComparisonOk returns a tuple with the Comparison field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComparison

`func (o *OverviewDashboardMetricsOut) SetComparison(v OverviewComparisonOut)`

SetComparison sets Comparison field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


