# WorkflowExecutionDetailOut

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** |  | 
**WorkflowId** | **string** |  | 
**WorkflowName** | **string** |  | 
**ExecutionId** | **string** |  | 
**TriggerSource** | **string** |  | 
**TriggerData** | Pointer to **map[string]interface{}** |  | [optional] 
**Status** | **string** |  | 
**CurrentStep** | Pointer to **NullableInt32** |  | [optional] 
**ErrorMessage** | Pointer to **NullableString** |  | [optional] 
**StartedAt** | Pointer to **NullableString** |  | [optional] 
**CompletedAt** | Pointer to **NullableString** |  | [optional] 
**DurationSeconds** | Pointer to **NullableFloat32** |  | [optional] 
**Steps** | [**[]ExecutionStepDetailOut**](ExecutionStepDetailOut.md) |  | 
**ExecutionData** | Pointer to **map[string]interface{}** |  | [optional] 
**WorkflowConfig** | Pointer to **map[string]interface{}** |  | [optional] 

## Methods

### NewWorkflowExecutionDetailOut

`func NewWorkflowExecutionDetailOut(id string, workflowId string, workflowName string, executionId string, triggerSource string, status string, steps []ExecutionStepDetailOut, ) *WorkflowExecutionDetailOut`

NewWorkflowExecutionDetailOut instantiates a new WorkflowExecutionDetailOut object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWorkflowExecutionDetailOutWithDefaults

`func NewWorkflowExecutionDetailOutWithDefaults() *WorkflowExecutionDetailOut`

NewWorkflowExecutionDetailOutWithDefaults instantiates a new WorkflowExecutionDetailOut object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *WorkflowExecutionDetailOut) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *WorkflowExecutionDetailOut) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *WorkflowExecutionDetailOut) SetId(v string)`

SetId sets Id field to given value.


### GetWorkflowId

`func (o *WorkflowExecutionDetailOut) GetWorkflowId() string`

GetWorkflowId returns the WorkflowId field if non-nil, zero value otherwise.

### GetWorkflowIdOk

`func (o *WorkflowExecutionDetailOut) GetWorkflowIdOk() (*string, bool)`

GetWorkflowIdOk returns a tuple with the WorkflowId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWorkflowId

`func (o *WorkflowExecutionDetailOut) SetWorkflowId(v string)`

SetWorkflowId sets WorkflowId field to given value.


### GetWorkflowName

`func (o *WorkflowExecutionDetailOut) GetWorkflowName() string`

GetWorkflowName returns the WorkflowName field if non-nil, zero value otherwise.

### GetWorkflowNameOk

`func (o *WorkflowExecutionDetailOut) GetWorkflowNameOk() (*string, bool)`

GetWorkflowNameOk returns a tuple with the WorkflowName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWorkflowName

`func (o *WorkflowExecutionDetailOut) SetWorkflowName(v string)`

SetWorkflowName sets WorkflowName field to given value.


### GetExecutionId

`func (o *WorkflowExecutionDetailOut) GetExecutionId() string`

GetExecutionId returns the ExecutionId field if non-nil, zero value otherwise.

### GetExecutionIdOk

`func (o *WorkflowExecutionDetailOut) GetExecutionIdOk() (*string, bool)`

GetExecutionIdOk returns a tuple with the ExecutionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExecutionId

`func (o *WorkflowExecutionDetailOut) SetExecutionId(v string)`

SetExecutionId sets ExecutionId field to given value.


### GetTriggerSource

`func (o *WorkflowExecutionDetailOut) GetTriggerSource() string`

GetTriggerSource returns the TriggerSource field if non-nil, zero value otherwise.

### GetTriggerSourceOk

`func (o *WorkflowExecutionDetailOut) GetTriggerSourceOk() (*string, bool)`

GetTriggerSourceOk returns a tuple with the TriggerSource field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTriggerSource

`func (o *WorkflowExecutionDetailOut) SetTriggerSource(v string)`

SetTriggerSource sets TriggerSource field to given value.


### GetTriggerData

`func (o *WorkflowExecutionDetailOut) GetTriggerData() map[string]interface{}`

GetTriggerData returns the TriggerData field if non-nil, zero value otherwise.

### GetTriggerDataOk

`func (o *WorkflowExecutionDetailOut) GetTriggerDataOk() (*map[string]interface{}, bool)`

GetTriggerDataOk returns a tuple with the TriggerData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTriggerData

`func (o *WorkflowExecutionDetailOut) SetTriggerData(v map[string]interface{})`

SetTriggerData sets TriggerData field to given value.

### HasTriggerData

`func (o *WorkflowExecutionDetailOut) HasTriggerData() bool`

HasTriggerData returns a boolean if a field has been set.

### SetTriggerDataNil

`func (o *WorkflowExecutionDetailOut) SetTriggerDataNil(b bool)`

 SetTriggerDataNil sets the value for TriggerData to be an explicit nil

### UnsetTriggerData
`func (o *WorkflowExecutionDetailOut) UnsetTriggerData()`

UnsetTriggerData ensures that no value is present for TriggerData, not even an explicit nil
### GetStatus

`func (o *WorkflowExecutionDetailOut) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *WorkflowExecutionDetailOut) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *WorkflowExecutionDetailOut) SetStatus(v string)`

SetStatus sets Status field to given value.


### GetCurrentStep

`func (o *WorkflowExecutionDetailOut) GetCurrentStep() int32`

GetCurrentStep returns the CurrentStep field if non-nil, zero value otherwise.

### GetCurrentStepOk

`func (o *WorkflowExecutionDetailOut) GetCurrentStepOk() (*int32, bool)`

GetCurrentStepOk returns a tuple with the CurrentStep field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrentStep

`func (o *WorkflowExecutionDetailOut) SetCurrentStep(v int32)`

SetCurrentStep sets CurrentStep field to given value.

### HasCurrentStep

`func (o *WorkflowExecutionDetailOut) HasCurrentStep() bool`

HasCurrentStep returns a boolean if a field has been set.

### SetCurrentStepNil

`func (o *WorkflowExecutionDetailOut) SetCurrentStepNil(b bool)`

 SetCurrentStepNil sets the value for CurrentStep to be an explicit nil

### UnsetCurrentStep
`func (o *WorkflowExecutionDetailOut) UnsetCurrentStep()`

UnsetCurrentStep ensures that no value is present for CurrentStep, not even an explicit nil
### GetErrorMessage

`func (o *WorkflowExecutionDetailOut) GetErrorMessage() string`

GetErrorMessage returns the ErrorMessage field if non-nil, zero value otherwise.

### GetErrorMessageOk

`func (o *WorkflowExecutionDetailOut) GetErrorMessageOk() (*string, bool)`

GetErrorMessageOk returns a tuple with the ErrorMessage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrorMessage

`func (o *WorkflowExecutionDetailOut) SetErrorMessage(v string)`

SetErrorMessage sets ErrorMessage field to given value.

### HasErrorMessage

`func (o *WorkflowExecutionDetailOut) HasErrorMessage() bool`

HasErrorMessage returns a boolean if a field has been set.

### SetErrorMessageNil

`func (o *WorkflowExecutionDetailOut) SetErrorMessageNil(b bool)`

 SetErrorMessageNil sets the value for ErrorMessage to be an explicit nil

### UnsetErrorMessage
`func (o *WorkflowExecutionDetailOut) UnsetErrorMessage()`

UnsetErrorMessage ensures that no value is present for ErrorMessage, not even an explicit nil
### GetStartedAt

`func (o *WorkflowExecutionDetailOut) GetStartedAt() string`

GetStartedAt returns the StartedAt field if non-nil, zero value otherwise.

### GetStartedAtOk

`func (o *WorkflowExecutionDetailOut) GetStartedAtOk() (*string, bool)`

GetStartedAtOk returns a tuple with the StartedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartedAt

`func (o *WorkflowExecutionDetailOut) SetStartedAt(v string)`

SetStartedAt sets StartedAt field to given value.

### HasStartedAt

`func (o *WorkflowExecutionDetailOut) HasStartedAt() bool`

HasStartedAt returns a boolean if a field has been set.

### SetStartedAtNil

`func (o *WorkflowExecutionDetailOut) SetStartedAtNil(b bool)`

 SetStartedAtNil sets the value for StartedAt to be an explicit nil

### UnsetStartedAt
`func (o *WorkflowExecutionDetailOut) UnsetStartedAt()`

UnsetStartedAt ensures that no value is present for StartedAt, not even an explicit nil
### GetCompletedAt

`func (o *WorkflowExecutionDetailOut) GetCompletedAt() string`

GetCompletedAt returns the CompletedAt field if non-nil, zero value otherwise.

### GetCompletedAtOk

`func (o *WorkflowExecutionDetailOut) GetCompletedAtOk() (*string, bool)`

GetCompletedAtOk returns a tuple with the CompletedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompletedAt

`func (o *WorkflowExecutionDetailOut) SetCompletedAt(v string)`

SetCompletedAt sets CompletedAt field to given value.

### HasCompletedAt

`func (o *WorkflowExecutionDetailOut) HasCompletedAt() bool`

HasCompletedAt returns a boolean if a field has been set.

### SetCompletedAtNil

`func (o *WorkflowExecutionDetailOut) SetCompletedAtNil(b bool)`

 SetCompletedAtNil sets the value for CompletedAt to be an explicit nil

### UnsetCompletedAt
`func (o *WorkflowExecutionDetailOut) UnsetCompletedAt()`

UnsetCompletedAt ensures that no value is present for CompletedAt, not even an explicit nil
### GetDurationSeconds

`func (o *WorkflowExecutionDetailOut) GetDurationSeconds() float32`

GetDurationSeconds returns the DurationSeconds field if non-nil, zero value otherwise.

### GetDurationSecondsOk

`func (o *WorkflowExecutionDetailOut) GetDurationSecondsOk() (*float32, bool)`

GetDurationSecondsOk returns a tuple with the DurationSeconds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDurationSeconds

`func (o *WorkflowExecutionDetailOut) SetDurationSeconds(v float32)`

SetDurationSeconds sets DurationSeconds field to given value.

### HasDurationSeconds

`func (o *WorkflowExecutionDetailOut) HasDurationSeconds() bool`

HasDurationSeconds returns a boolean if a field has been set.

### SetDurationSecondsNil

`func (o *WorkflowExecutionDetailOut) SetDurationSecondsNil(b bool)`

 SetDurationSecondsNil sets the value for DurationSeconds to be an explicit nil

### UnsetDurationSeconds
`func (o *WorkflowExecutionDetailOut) UnsetDurationSeconds()`

UnsetDurationSeconds ensures that no value is present for DurationSeconds, not even an explicit nil
### GetSteps

`func (o *WorkflowExecutionDetailOut) GetSteps() []ExecutionStepDetailOut`

GetSteps returns the Steps field if non-nil, zero value otherwise.

### GetStepsOk

`func (o *WorkflowExecutionDetailOut) GetStepsOk() (*[]ExecutionStepDetailOut, bool)`

GetStepsOk returns a tuple with the Steps field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSteps

`func (o *WorkflowExecutionDetailOut) SetSteps(v []ExecutionStepDetailOut)`

SetSteps sets Steps field to given value.


### GetExecutionData

`func (o *WorkflowExecutionDetailOut) GetExecutionData() map[string]interface{}`

GetExecutionData returns the ExecutionData field if non-nil, zero value otherwise.

### GetExecutionDataOk

`func (o *WorkflowExecutionDetailOut) GetExecutionDataOk() (*map[string]interface{}, bool)`

GetExecutionDataOk returns a tuple with the ExecutionData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExecutionData

`func (o *WorkflowExecutionDetailOut) SetExecutionData(v map[string]interface{})`

SetExecutionData sets ExecutionData field to given value.

### HasExecutionData

`func (o *WorkflowExecutionDetailOut) HasExecutionData() bool`

HasExecutionData returns a boolean if a field has been set.

### SetExecutionDataNil

`func (o *WorkflowExecutionDetailOut) SetExecutionDataNil(b bool)`

 SetExecutionDataNil sets the value for ExecutionData to be an explicit nil

### UnsetExecutionData
`func (o *WorkflowExecutionDetailOut) UnsetExecutionData()`

UnsetExecutionData ensures that no value is present for ExecutionData, not even an explicit nil
### GetWorkflowConfig

`func (o *WorkflowExecutionDetailOut) GetWorkflowConfig() map[string]interface{}`

GetWorkflowConfig returns the WorkflowConfig field if non-nil, zero value otherwise.

### GetWorkflowConfigOk

`func (o *WorkflowExecutionDetailOut) GetWorkflowConfigOk() (*map[string]interface{}, bool)`

GetWorkflowConfigOk returns a tuple with the WorkflowConfig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWorkflowConfig

`func (o *WorkflowExecutionDetailOut) SetWorkflowConfig(v map[string]interface{})`

SetWorkflowConfig sets WorkflowConfig field to given value.

### HasWorkflowConfig

`func (o *WorkflowExecutionDetailOut) HasWorkflowConfig() bool`

HasWorkflowConfig returns a boolean if a field has been set.

### SetWorkflowConfigNil

`func (o *WorkflowExecutionDetailOut) SetWorkflowConfigNil(b bool)`

 SetWorkflowConfigNil sets the value for WorkflowConfig to be an explicit nil

### UnsetWorkflowConfig
`func (o *WorkflowExecutionDetailOut) UnsetWorkflowConfig()`

UnsetWorkflowConfig ensures that no value is present for WorkflowConfig, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


