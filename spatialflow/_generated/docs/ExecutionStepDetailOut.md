# ExecutionStepDetailOut

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**StepIndex** | **int32** |  | 
**NodeId** | Pointer to **NullableString** |  | [optional] 
**StepName** | **string** |  | 
**StepType** | **string** |  | 
**Status** | **string** |  | 
**StartedAt** | Pointer to **NullableString** |  | [optional] 
**CompletedAt** | Pointer to **NullableString** |  | [optional] 
**DurationMs** | Pointer to **NullableFloat32** |  | [optional] 
**ErrorMessage** | Pointer to **NullableString** |  | [optional] 
**InputData** | Pointer to **map[string]interface{}** |  | [optional] 
**OutputData** | Pointer to **map[string]interface{}** |  | [optional] 

## Methods

### NewExecutionStepDetailOut

`func NewExecutionStepDetailOut(stepIndex int32, stepName string, stepType string, status string, ) *ExecutionStepDetailOut`

NewExecutionStepDetailOut instantiates a new ExecutionStepDetailOut object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewExecutionStepDetailOutWithDefaults

`func NewExecutionStepDetailOutWithDefaults() *ExecutionStepDetailOut`

NewExecutionStepDetailOutWithDefaults instantiates a new ExecutionStepDetailOut object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetStepIndex

`func (o *ExecutionStepDetailOut) GetStepIndex() int32`

GetStepIndex returns the StepIndex field if non-nil, zero value otherwise.

### GetStepIndexOk

`func (o *ExecutionStepDetailOut) GetStepIndexOk() (*int32, bool)`

GetStepIndexOk returns a tuple with the StepIndex field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStepIndex

`func (o *ExecutionStepDetailOut) SetStepIndex(v int32)`

SetStepIndex sets StepIndex field to given value.


### GetNodeId

`func (o *ExecutionStepDetailOut) GetNodeId() string`

GetNodeId returns the NodeId field if non-nil, zero value otherwise.

### GetNodeIdOk

`func (o *ExecutionStepDetailOut) GetNodeIdOk() (*string, bool)`

GetNodeIdOk returns a tuple with the NodeId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNodeId

`func (o *ExecutionStepDetailOut) SetNodeId(v string)`

SetNodeId sets NodeId field to given value.

### HasNodeId

`func (o *ExecutionStepDetailOut) HasNodeId() bool`

HasNodeId returns a boolean if a field has been set.

### SetNodeIdNil

`func (o *ExecutionStepDetailOut) SetNodeIdNil(b bool)`

 SetNodeIdNil sets the value for NodeId to be an explicit nil

### UnsetNodeId
`func (o *ExecutionStepDetailOut) UnsetNodeId()`

UnsetNodeId ensures that no value is present for NodeId, not even an explicit nil
### GetStepName

`func (o *ExecutionStepDetailOut) GetStepName() string`

GetStepName returns the StepName field if non-nil, zero value otherwise.

### GetStepNameOk

`func (o *ExecutionStepDetailOut) GetStepNameOk() (*string, bool)`

GetStepNameOk returns a tuple with the StepName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStepName

`func (o *ExecutionStepDetailOut) SetStepName(v string)`

SetStepName sets StepName field to given value.


### GetStepType

`func (o *ExecutionStepDetailOut) GetStepType() string`

GetStepType returns the StepType field if non-nil, zero value otherwise.

### GetStepTypeOk

`func (o *ExecutionStepDetailOut) GetStepTypeOk() (*string, bool)`

GetStepTypeOk returns a tuple with the StepType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStepType

`func (o *ExecutionStepDetailOut) SetStepType(v string)`

SetStepType sets StepType field to given value.


### GetStatus

`func (o *ExecutionStepDetailOut) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *ExecutionStepDetailOut) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *ExecutionStepDetailOut) SetStatus(v string)`

SetStatus sets Status field to given value.


### GetStartedAt

`func (o *ExecutionStepDetailOut) GetStartedAt() string`

GetStartedAt returns the StartedAt field if non-nil, zero value otherwise.

### GetStartedAtOk

`func (o *ExecutionStepDetailOut) GetStartedAtOk() (*string, bool)`

GetStartedAtOk returns a tuple with the StartedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartedAt

`func (o *ExecutionStepDetailOut) SetStartedAt(v string)`

SetStartedAt sets StartedAt field to given value.

### HasStartedAt

`func (o *ExecutionStepDetailOut) HasStartedAt() bool`

HasStartedAt returns a boolean if a field has been set.

### SetStartedAtNil

`func (o *ExecutionStepDetailOut) SetStartedAtNil(b bool)`

 SetStartedAtNil sets the value for StartedAt to be an explicit nil

### UnsetStartedAt
`func (o *ExecutionStepDetailOut) UnsetStartedAt()`

UnsetStartedAt ensures that no value is present for StartedAt, not even an explicit nil
### GetCompletedAt

`func (o *ExecutionStepDetailOut) GetCompletedAt() string`

GetCompletedAt returns the CompletedAt field if non-nil, zero value otherwise.

### GetCompletedAtOk

`func (o *ExecutionStepDetailOut) GetCompletedAtOk() (*string, bool)`

GetCompletedAtOk returns a tuple with the CompletedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompletedAt

`func (o *ExecutionStepDetailOut) SetCompletedAt(v string)`

SetCompletedAt sets CompletedAt field to given value.

### HasCompletedAt

`func (o *ExecutionStepDetailOut) HasCompletedAt() bool`

HasCompletedAt returns a boolean if a field has been set.

### SetCompletedAtNil

`func (o *ExecutionStepDetailOut) SetCompletedAtNil(b bool)`

 SetCompletedAtNil sets the value for CompletedAt to be an explicit nil

### UnsetCompletedAt
`func (o *ExecutionStepDetailOut) UnsetCompletedAt()`

UnsetCompletedAt ensures that no value is present for CompletedAt, not even an explicit nil
### GetDurationMs

`func (o *ExecutionStepDetailOut) GetDurationMs() float32`

GetDurationMs returns the DurationMs field if non-nil, zero value otherwise.

### GetDurationMsOk

`func (o *ExecutionStepDetailOut) GetDurationMsOk() (*float32, bool)`

GetDurationMsOk returns a tuple with the DurationMs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDurationMs

`func (o *ExecutionStepDetailOut) SetDurationMs(v float32)`

SetDurationMs sets DurationMs field to given value.

### HasDurationMs

`func (o *ExecutionStepDetailOut) HasDurationMs() bool`

HasDurationMs returns a boolean if a field has been set.

### SetDurationMsNil

`func (o *ExecutionStepDetailOut) SetDurationMsNil(b bool)`

 SetDurationMsNil sets the value for DurationMs to be an explicit nil

### UnsetDurationMs
`func (o *ExecutionStepDetailOut) UnsetDurationMs()`

UnsetDurationMs ensures that no value is present for DurationMs, not even an explicit nil
### GetErrorMessage

`func (o *ExecutionStepDetailOut) GetErrorMessage() string`

GetErrorMessage returns the ErrorMessage field if non-nil, zero value otherwise.

### GetErrorMessageOk

`func (o *ExecutionStepDetailOut) GetErrorMessageOk() (*string, bool)`

GetErrorMessageOk returns a tuple with the ErrorMessage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrorMessage

`func (o *ExecutionStepDetailOut) SetErrorMessage(v string)`

SetErrorMessage sets ErrorMessage field to given value.

### HasErrorMessage

`func (o *ExecutionStepDetailOut) HasErrorMessage() bool`

HasErrorMessage returns a boolean if a field has been set.

### SetErrorMessageNil

`func (o *ExecutionStepDetailOut) SetErrorMessageNil(b bool)`

 SetErrorMessageNil sets the value for ErrorMessage to be an explicit nil

### UnsetErrorMessage
`func (o *ExecutionStepDetailOut) UnsetErrorMessage()`

UnsetErrorMessage ensures that no value is present for ErrorMessage, not even an explicit nil
### GetInputData

`func (o *ExecutionStepDetailOut) GetInputData() map[string]interface{}`

GetInputData returns the InputData field if non-nil, zero value otherwise.

### GetInputDataOk

`func (o *ExecutionStepDetailOut) GetInputDataOk() (*map[string]interface{}, bool)`

GetInputDataOk returns a tuple with the InputData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInputData

`func (o *ExecutionStepDetailOut) SetInputData(v map[string]interface{})`

SetInputData sets InputData field to given value.

### HasInputData

`func (o *ExecutionStepDetailOut) HasInputData() bool`

HasInputData returns a boolean if a field has been set.

### SetInputDataNil

`func (o *ExecutionStepDetailOut) SetInputDataNil(b bool)`

 SetInputDataNil sets the value for InputData to be an explicit nil

### UnsetInputData
`func (o *ExecutionStepDetailOut) UnsetInputData()`

UnsetInputData ensures that no value is present for InputData, not even an explicit nil
### GetOutputData

`func (o *ExecutionStepDetailOut) GetOutputData() map[string]interface{}`

GetOutputData returns the OutputData field if non-nil, zero value otherwise.

### GetOutputDataOk

`func (o *ExecutionStepDetailOut) GetOutputDataOk() (*map[string]interface{}, bool)`

GetOutputDataOk returns a tuple with the OutputData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOutputData

`func (o *ExecutionStepDetailOut) SetOutputData(v map[string]interface{})`

SetOutputData sets OutputData field to given value.

### HasOutputData

`func (o *ExecutionStepDetailOut) HasOutputData() bool`

HasOutputData returns a boolean if a field has been set.

### SetOutputDataNil

`func (o *ExecutionStepDetailOut) SetOutputDataNil(b bool)`

 SetOutputDataNil sets the value for OutputData to be an explicit nil

### UnsetOutputData
`func (o *ExecutionStepDetailOut) UnsetOutputData()`

UnsetOutputData ensures that no value is present for OutputData, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


