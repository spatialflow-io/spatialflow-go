# DashboardComparisonMetrics

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**EventsChange** | **int32** | Absolute change in events vs previous period | 
**WorkflowsChange** | **int32** | Absolute change in active workflows vs previous period | 
**SuccessRateChange** | Pointer to **NullableFloat32** |  | [optional] 

## Methods

### NewDashboardComparisonMetrics

`func NewDashboardComparisonMetrics(eventsChange int32, workflowsChange int32, ) *DashboardComparisonMetrics`

NewDashboardComparisonMetrics instantiates a new DashboardComparisonMetrics object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDashboardComparisonMetricsWithDefaults

`func NewDashboardComparisonMetricsWithDefaults() *DashboardComparisonMetrics`

NewDashboardComparisonMetricsWithDefaults instantiates a new DashboardComparisonMetrics object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEventsChange

`func (o *DashboardComparisonMetrics) GetEventsChange() int32`

GetEventsChange returns the EventsChange field if non-nil, zero value otherwise.

### GetEventsChangeOk

`func (o *DashboardComparisonMetrics) GetEventsChangeOk() (*int32, bool)`

GetEventsChangeOk returns a tuple with the EventsChange field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEventsChange

`func (o *DashboardComparisonMetrics) SetEventsChange(v int32)`

SetEventsChange sets EventsChange field to given value.


### GetWorkflowsChange

`func (o *DashboardComparisonMetrics) GetWorkflowsChange() int32`

GetWorkflowsChange returns the WorkflowsChange field if non-nil, zero value otherwise.

### GetWorkflowsChangeOk

`func (o *DashboardComparisonMetrics) GetWorkflowsChangeOk() (*int32, bool)`

GetWorkflowsChangeOk returns a tuple with the WorkflowsChange field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWorkflowsChange

`func (o *DashboardComparisonMetrics) SetWorkflowsChange(v int32)`

SetWorkflowsChange sets WorkflowsChange field to given value.


### GetSuccessRateChange

`func (o *DashboardComparisonMetrics) GetSuccessRateChange() float32`

GetSuccessRateChange returns the SuccessRateChange field if non-nil, zero value otherwise.

### GetSuccessRateChangeOk

`func (o *DashboardComparisonMetrics) GetSuccessRateChangeOk() (*float32, bool)`

GetSuccessRateChangeOk returns a tuple with the SuccessRateChange field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSuccessRateChange

`func (o *DashboardComparisonMetrics) SetSuccessRateChange(v float32)`

SetSuccessRateChange sets SuccessRateChange field to given value.

### HasSuccessRateChange

`func (o *DashboardComparisonMetrics) HasSuccessRateChange() bool`

HasSuccessRateChange returns a boolean if a field has been set.

### SetSuccessRateChangeNil

`func (o *DashboardComparisonMetrics) SetSuccessRateChangeNil(b bool)`

 SetSuccessRateChangeNil sets the value for SuccessRateChange to be an explicit nil

### UnsetSuccessRateChange
`func (o *DashboardComparisonMetrics) UnsetSuccessRateChange()`

UnsetSuccessRateChange ensures that no value is present for SuccessRateChange, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


