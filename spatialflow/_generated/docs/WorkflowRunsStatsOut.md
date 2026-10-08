# WorkflowRunsStatsOut

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**TotalRuns24h** | **int32** |  | 
**SuccessfulRuns24h** | **int32** |  | 
**FailedRuns24h** | **int32** |  | 
**SuccessRate** | Pointer to **NullableFloat32** |  | [optional] 

## Methods

### NewWorkflowRunsStatsOut

`func NewWorkflowRunsStatsOut(totalRuns24h int32, successfulRuns24h int32, failedRuns24h int32, ) *WorkflowRunsStatsOut`

NewWorkflowRunsStatsOut instantiates a new WorkflowRunsStatsOut object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWorkflowRunsStatsOutWithDefaults

`func NewWorkflowRunsStatsOutWithDefaults() *WorkflowRunsStatsOut`

NewWorkflowRunsStatsOutWithDefaults instantiates a new WorkflowRunsStatsOut object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTotalRuns24h

`func (o *WorkflowRunsStatsOut) GetTotalRuns24h() int32`

GetTotalRuns24h returns the TotalRuns24h field if non-nil, zero value otherwise.

### GetTotalRuns24hOk

`func (o *WorkflowRunsStatsOut) GetTotalRuns24hOk() (*int32, bool)`

GetTotalRuns24hOk returns a tuple with the TotalRuns24h field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalRuns24h

`func (o *WorkflowRunsStatsOut) SetTotalRuns24h(v int32)`

SetTotalRuns24h sets TotalRuns24h field to given value.


### GetSuccessfulRuns24h

`func (o *WorkflowRunsStatsOut) GetSuccessfulRuns24h() int32`

GetSuccessfulRuns24h returns the SuccessfulRuns24h field if non-nil, zero value otherwise.

### GetSuccessfulRuns24hOk

`func (o *WorkflowRunsStatsOut) GetSuccessfulRuns24hOk() (*int32, bool)`

GetSuccessfulRuns24hOk returns a tuple with the SuccessfulRuns24h field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSuccessfulRuns24h

`func (o *WorkflowRunsStatsOut) SetSuccessfulRuns24h(v int32)`

SetSuccessfulRuns24h sets SuccessfulRuns24h field to given value.


### GetFailedRuns24h

`func (o *WorkflowRunsStatsOut) GetFailedRuns24h() int32`

GetFailedRuns24h returns the FailedRuns24h field if non-nil, zero value otherwise.

### GetFailedRuns24hOk

`func (o *WorkflowRunsStatsOut) GetFailedRuns24hOk() (*int32, bool)`

GetFailedRuns24hOk returns a tuple with the FailedRuns24h field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFailedRuns24h

`func (o *WorkflowRunsStatsOut) SetFailedRuns24h(v int32)`

SetFailedRuns24h sets FailedRuns24h field to given value.


### GetSuccessRate

`func (o *WorkflowRunsStatsOut) GetSuccessRate() float32`

GetSuccessRate returns the SuccessRate field if non-nil, zero value otherwise.

### GetSuccessRateOk

`func (o *WorkflowRunsStatsOut) GetSuccessRateOk() (*float32, bool)`

GetSuccessRateOk returns a tuple with the SuccessRate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSuccessRate

`func (o *WorkflowRunsStatsOut) SetSuccessRate(v float32)`

SetSuccessRate sets SuccessRate field to given value.

### HasSuccessRate

`func (o *WorkflowRunsStatsOut) HasSuccessRate() bool`

HasSuccessRate returns a boolean if a field has been set.

### SetSuccessRateNil

`func (o *WorkflowRunsStatsOut) SetSuccessRateNil(b bool)`

 SetSuccessRateNil sets the value for SuccessRate to be an explicit nil

### UnsetSuccessRate
`func (o *WorkflowRunsStatsOut) UnsetSuccessRate()`

UnsetSuccessRate ensures that no value is present for SuccessRate, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


