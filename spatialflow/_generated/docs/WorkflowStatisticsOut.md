# WorkflowStatisticsOut

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**TotalExecutions** | Pointer to **int32** |  | [optional] [default to 0]
**SuccessfulExecutions** | Pointer to **int32** |  | [optional] [default to 0]
**FailedExecutions** | Pointer to **int32** |  | [optional] [default to 0]
**RunningExecutions** | Pointer to **int32** |  | [optional] [default to 0]
**CancelledExecutions** | Pointer to **int32** |  | [optional] [default to 0]
**SuccessRate** | Pointer to **float32** |  | [optional] [default to 0.0]
**AverageDurationSeconds** | Pointer to **NullableFloat32** |  | [optional] 
**RecentExecutions** | Pointer to [**[]RecentExecutionOut**](RecentExecutionOut.md) |  | [optional] 
**Last7Days** | Pointer to **map[string]interface{}** |  | [optional] 
**Last30Days** | Pointer to **map[string]interface{}** |  | [optional] 
**ExecutionsByDay** | Pointer to **[]interface{}** |  | [optional] 
**CommonErrors** | Pointer to **[]interface{}** |  | [optional] 

## Methods

### NewWorkflowStatisticsOut

`func NewWorkflowStatisticsOut() *WorkflowStatisticsOut`

NewWorkflowStatisticsOut instantiates a new WorkflowStatisticsOut object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWorkflowStatisticsOutWithDefaults

`func NewWorkflowStatisticsOutWithDefaults() *WorkflowStatisticsOut`

NewWorkflowStatisticsOutWithDefaults instantiates a new WorkflowStatisticsOut object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTotalExecutions

`func (o *WorkflowStatisticsOut) GetTotalExecutions() int32`

GetTotalExecutions returns the TotalExecutions field if non-nil, zero value otherwise.

### GetTotalExecutionsOk

`func (o *WorkflowStatisticsOut) GetTotalExecutionsOk() (*int32, bool)`

GetTotalExecutionsOk returns a tuple with the TotalExecutions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalExecutions

`func (o *WorkflowStatisticsOut) SetTotalExecutions(v int32)`

SetTotalExecutions sets TotalExecutions field to given value.

### HasTotalExecutions

`func (o *WorkflowStatisticsOut) HasTotalExecutions() bool`

HasTotalExecutions returns a boolean if a field has been set.

### GetSuccessfulExecutions

`func (o *WorkflowStatisticsOut) GetSuccessfulExecutions() int32`

GetSuccessfulExecutions returns the SuccessfulExecutions field if non-nil, zero value otherwise.

### GetSuccessfulExecutionsOk

`func (o *WorkflowStatisticsOut) GetSuccessfulExecutionsOk() (*int32, bool)`

GetSuccessfulExecutionsOk returns a tuple with the SuccessfulExecutions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSuccessfulExecutions

`func (o *WorkflowStatisticsOut) SetSuccessfulExecutions(v int32)`

SetSuccessfulExecutions sets SuccessfulExecutions field to given value.

### HasSuccessfulExecutions

`func (o *WorkflowStatisticsOut) HasSuccessfulExecutions() bool`

HasSuccessfulExecutions returns a boolean if a field has been set.

### GetFailedExecutions

`func (o *WorkflowStatisticsOut) GetFailedExecutions() int32`

GetFailedExecutions returns the FailedExecutions field if non-nil, zero value otherwise.

### GetFailedExecutionsOk

`func (o *WorkflowStatisticsOut) GetFailedExecutionsOk() (*int32, bool)`

GetFailedExecutionsOk returns a tuple with the FailedExecutions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFailedExecutions

`func (o *WorkflowStatisticsOut) SetFailedExecutions(v int32)`

SetFailedExecutions sets FailedExecutions field to given value.

### HasFailedExecutions

`func (o *WorkflowStatisticsOut) HasFailedExecutions() bool`

HasFailedExecutions returns a boolean if a field has been set.

### GetRunningExecutions

`func (o *WorkflowStatisticsOut) GetRunningExecutions() int32`

GetRunningExecutions returns the RunningExecutions field if non-nil, zero value otherwise.

### GetRunningExecutionsOk

`func (o *WorkflowStatisticsOut) GetRunningExecutionsOk() (*int32, bool)`

GetRunningExecutionsOk returns a tuple with the RunningExecutions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRunningExecutions

`func (o *WorkflowStatisticsOut) SetRunningExecutions(v int32)`

SetRunningExecutions sets RunningExecutions field to given value.

### HasRunningExecutions

`func (o *WorkflowStatisticsOut) HasRunningExecutions() bool`

HasRunningExecutions returns a boolean if a field has been set.

### GetCancelledExecutions

`func (o *WorkflowStatisticsOut) GetCancelledExecutions() int32`

GetCancelledExecutions returns the CancelledExecutions field if non-nil, zero value otherwise.

### GetCancelledExecutionsOk

`func (o *WorkflowStatisticsOut) GetCancelledExecutionsOk() (*int32, bool)`

GetCancelledExecutionsOk returns a tuple with the CancelledExecutions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCancelledExecutions

`func (o *WorkflowStatisticsOut) SetCancelledExecutions(v int32)`

SetCancelledExecutions sets CancelledExecutions field to given value.

### HasCancelledExecutions

`func (o *WorkflowStatisticsOut) HasCancelledExecutions() bool`

HasCancelledExecutions returns a boolean if a field has been set.

### GetSuccessRate

`func (o *WorkflowStatisticsOut) GetSuccessRate() float32`

GetSuccessRate returns the SuccessRate field if non-nil, zero value otherwise.

### GetSuccessRateOk

`func (o *WorkflowStatisticsOut) GetSuccessRateOk() (*float32, bool)`

GetSuccessRateOk returns a tuple with the SuccessRate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSuccessRate

`func (o *WorkflowStatisticsOut) SetSuccessRate(v float32)`

SetSuccessRate sets SuccessRate field to given value.

### HasSuccessRate

`func (o *WorkflowStatisticsOut) HasSuccessRate() bool`

HasSuccessRate returns a boolean if a field has been set.

### GetAverageDurationSeconds

`func (o *WorkflowStatisticsOut) GetAverageDurationSeconds() float32`

GetAverageDurationSeconds returns the AverageDurationSeconds field if non-nil, zero value otherwise.

### GetAverageDurationSecondsOk

`func (o *WorkflowStatisticsOut) GetAverageDurationSecondsOk() (*float32, bool)`

GetAverageDurationSecondsOk returns a tuple with the AverageDurationSeconds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAverageDurationSeconds

`func (o *WorkflowStatisticsOut) SetAverageDurationSeconds(v float32)`

SetAverageDurationSeconds sets AverageDurationSeconds field to given value.

### HasAverageDurationSeconds

`func (o *WorkflowStatisticsOut) HasAverageDurationSeconds() bool`

HasAverageDurationSeconds returns a boolean if a field has been set.

### SetAverageDurationSecondsNil

`func (o *WorkflowStatisticsOut) SetAverageDurationSecondsNil(b bool)`

 SetAverageDurationSecondsNil sets the value for AverageDurationSeconds to be an explicit nil

### UnsetAverageDurationSeconds
`func (o *WorkflowStatisticsOut) UnsetAverageDurationSeconds()`

UnsetAverageDurationSeconds ensures that no value is present for AverageDurationSeconds, not even an explicit nil
### GetRecentExecutions

`func (o *WorkflowStatisticsOut) GetRecentExecutions() []RecentExecutionOut`

GetRecentExecutions returns the RecentExecutions field if non-nil, zero value otherwise.

### GetRecentExecutionsOk

`func (o *WorkflowStatisticsOut) GetRecentExecutionsOk() (*[]RecentExecutionOut, bool)`

GetRecentExecutionsOk returns a tuple with the RecentExecutions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecentExecutions

`func (o *WorkflowStatisticsOut) SetRecentExecutions(v []RecentExecutionOut)`

SetRecentExecutions sets RecentExecutions field to given value.

### HasRecentExecutions

`func (o *WorkflowStatisticsOut) HasRecentExecutions() bool`

HasRecentExecutions returns a boolean if a field has been set.

### SetRecentExecutionsNil

`func (o *WorkflowStatisticsOut) SetRecentExecutionsNil(b bool)`

 SetRecentExecutionsNil sets the value for RecentExecutions to be an explicit nil

### UnsetRecentExecutions
`func (o *WorkflowStatisticsOut) UnsetRecentExecutions()`

UnsetRecentExecutions ensures that no value is present for RecentExecutions, not even an explicit nil
### GetLast7Days

`func (o *WorkflowStatisticsOut) GetLast7Days() map[string]interface{}`

GetLast7Days returns the Last7Days field if non-nil, zero value otherwise.

### GetLast7DaysOk

`func (o *WorkflowStatisticsOut) GetLast7DaysOk() (*map[string]interface{}, bool)`

GetLast7DaysOk returns a tuple with the Last7Days field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLast7Days

`func (o *WorkflowStatisticsOut) SetLast7Days(v map[string]interface{})`

SetLast7Days sets Last7Days field to given value.

### HasLast7Days

`func (o *WorkflowStatisticsOut) HasLast7Days() bool`

HasLast7Days returns a boolean if a field has been set.

### SetLast7DaysNil

`func (o *WorkflowStatisticsOut) SetLast7DaysNil(b bool)`

 SetLast7DaysNil sets the value for Last7Days to be an explicit nil

### UnsetLast7Days
`func (o *WorkflowStatisticsOut) UnsetLast7Days()`

UnsetLast7Days ensures that no value is present for Last7Days, not even an explicit nil
### GetLast30Days

`func (o *WorkflowStatisticsOut) GetLast30Days() map[string]interface{}`

GetLast30Days returns the Last30Days field if non-nil, zero value otherwise.

### GetLast30DaysOk

`func (o *WorkflowStatisticsOut) GetLast30DaysOk() (*map[string]interface{}, bool)`

GetLast30DaysOk returns a tuple with the Last30Days field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLast30Days

`func (o *WorkflowStatisticsOut) SetLast30Days(v map[string]interface{})`

SetLast30Days sets Last30Days field to given value.

### HasLast30Days

`func (o *WorkflowStatisticsOut) HasLast30Days() bool`

HasLast30Days returns a boolean if a field has been set.

### SetLast30DaysNil

`func (o *WorkflowStatisticsOut) SetLast30DaysNil(b bool)`

 SetLast30DaysNil sets the value for Last30Days to be an explicit nil

### UnsetLast30Days
`func (o *WorkflowStatisticsOut) UnsetLast30Days()`

UnsetLast30Days ensures that no value is present for Last30Days, not even an explicit nil
### GetExecutionsByDay

`func (o *WorkflowStatisticsOut) GetExecutionsByDay() []interface{}`

GetExecutionsByDay returns the ExecutionsByDay field if non-nil, zero value otherwise.

### GetExecutionsByDayOk

`func (o *WorkflowStatisticsOut) GetExecutionsByDayOk() (*[]interface{}, bool)`

GetExecutionsByDayOk returns a tuple with the ExecutionsByDay field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExecutionsByDay

`func (o *WorkflowStatisticsOut) SetExecutionsByDay(v []interface{})`

SetExecutionsByDay sets ExecutionsByDay field to given value.

### HasExecutionsByDay

`func (o *WorkflowStatisticsOut) HasExecutionsByDay() bool`

HasExecutionsByDay returns a boolean if a field has been set.

### SetExecutionsByDayNil

`func (o *WorkflowStatisticsOut) SetExecutionsByDayNil(b bool)`

 SetExecutionsByDayNil sets the value for ExecutionsByDay to be an explicit nil

### UnsetExecutionsByDay
`func (o *WorkflowStatisticsOut) UnsetExecutionsByDay()`

UnsetExecutionsByDay ensures that no value is present for ExecutionsByDay, not even an explicit nil
### GetCommonErrors

`func (o *WorkflowStatisticsOut) GetCommonErrors() []interface{}`

GetCommonErrors returns the CommonErrors field if non-nil, zero value otherwise.

### GetCommonErrorsOk

`func (o *WorkflowStatisticsOut) GetCommonErrorsOk() (*[]interface{}, bool)`

GetCommonErrorsOk returns a tuple with the CommonErrors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCommonErrors

`func (o *WorkflowStatisticsOut) SetCommonErrors(v []interface{})`

SetCommonErrors sets CommonErrors field to given value.

### HasCommonErrors

`func (o *WorkflowStatisticsOut) HasCommonErrors() bool`

HasCommonErrors returns a boolean if a field has been set.

### SetCommonErrorsNil

`func (o *WorkflowStatisticsOut) SetCommonErrorsNil(b bool)`

 SetCommonErrorsNil sets the value for CommonErrors to be an explicit nil

### UnsetCommonErrors
`func (o *WorkflowStatisticsOut) UnsetCommonErrors()`

UnsetCommonErrors ensures that no value is present for CommonErrors, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


