# ExecutionDetailOut

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

## Methods

### NewExecutionDetailOut

`func NewExecutionDetailOut(id string, workflowId string, workflowName string, executionId string, triggerSource string, status string, steps []ExecutionStepDetailOut, ) *ExecutionDetailOut`

NewExecutionDetailOut instantiates a new ExecutionDetailOut object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewExecutionDetailOutWithDefaults

`func NewExecutionDetailOutWithDefaults() *ExecutionDetailOut`

NewExecutionDetailOutWithDefaults instantiates a new ExecutionDetailOut object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ExecutionDetailOut) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ExecutionDetailOut) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ExecutionDetailOut) SetId(v string)`

SetId sets Id field to given value.


### GetWorkflowId

`func (o *ExecutionDetailOut) GetWorkflowId() string`

GetWorkflowId returns the WorkflowId field if non-nil, zero value otherwise.

### GetWorkflowIdOk

`func (o *ExecutionDetailOut) GetWorkflowIdOk() (*string, bool)`

GetWorkflowIdOk returns a tuple with the WorkflowId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWorkflowId

`func (o *ExecutionDetailOut) SetWorkflowId(v string)`

SetWorkflowId sets WorkflowId field to given value.


### GetWorkflowName

`func (o *ExecutionDetailOut) GetWorkflowName() string`

GetWorkflowName returns the WorkflowName field if non-nil, zero value otherwise.

### GetWorkflowNameOk

`func (o *ExecutionDetailOut) GetWorkflowNameOk() (*string, bool)`

GetWorkflowNameOk returns a tuple with the WorkflowName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWorkflowName

`func (o *ExecutionDetailOut) SetWorkflowName(v string)`

SetWorkflowName sets WorkflowName field to given value.


### GetExecutionId

`func (o *ExecutionDetailOut) GetExecutionId() string`

GetExecutionId returns the ExecutionId field if non-nil, zero value otherwise.

### GetExecutionIdOk

`func (o *ExecutionDetailOut) GetExecutionIdOk() (*string, bool)`

GetExecutionIdOk returns a tuple with the ExecutionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExecutionId

`func (o *ExecutionDetailOut) SetExecutionId(v string)`

SetExecutionId sets ExecutionId field to given value.


### GetTriggerSource

`func (o *ExecutionDetailOut) GetTriggerSource() string`

GetTriggerSource returns the TriggerSource field if non-nil, zero value otherwise.

### GetTriggerSourceOk

`func (o *ExecutionDetailOut) GetTriggerSourceOk() (*string, bool)`

GetTriggerSourceOk returns a tuple with the TriggerSource field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTriggerSource

`func (o *ExecutionDetailOut) SetTriggerSource(v string)`

SetTriggerSource sets TriggerSource field to given value.


### GetTriggerData

`func (o *ExecutionDetailOut) GetTriggerData() map[string]interface{}`

GetTriggerData returns the TriggerData field if non-nil, zero value otherwise.

### GetTriggerDataOk

`func (o *ExecutionDetailOut) GetTriggerDataOk() (*map[string]interface{}, bool)`

GetTriggerDataOk returns a tuple with the TriggerData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTriggerData

`func (o *ExecutionDetailOut) SetTriggerData(v map[string]interface{})`

SetTriggerData sets TriggerData field to given value.

### HasTriggerData

`func (o *ExecutionDetailOut) HasTriggerData() bool`

HasTriggerData returns a boolean if a field has been set.

### SetTriggerDataNil

`func (o *ExecutionDetailOut) SetTriggerDataNil(b bool)`

 SetTriggerDataNil sets the value for TriggerData to be an explicit nil

### UnsetTriggerData
`func (o *ExecutionDetailOut) UnsetTriggerData()`

UnsetTriggerData ensures that no value is present for TriggerData, not even an explicit nil
### GetStatus

`func (o *ExecutionDetailOut) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *ExecutionDetailOut) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *ExecutionDetailOut) SetStatus(v string)`

SetStatus sets Status field to given value.


### GetCurrentStep

`func (o *ExecutionDetailOut) GetCurrentStep() int32`

GetCurrentStep returns the CurrentStep field if non-nil, zero value otherwise.

### GetCurrentStepOk

`func (o *ExecutionDetailOut) GetCurrentStepOk() (*int32, bool)`

GetCurrentStepOk returns a tuple with the CurrentStep field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrentStep

`func (o *ExecutionDetailOut) SetCurrentStep(v int32)`

SetCurrentStep sets CurrentStep field to given value.

### HasCurrentStep

`func (o *ExecutionDetailOut) HasCurrentStep() bool`

HasCurrentStep returns a boolean if a field has been set.

### SetCurrentStepNil

`func (o *ExecutionDetailOut) SetCurrentStepNil(b bool)`

 SetCurrentStepNil sets the value for CurrentStep to be an explicit nil

### UnsetCurrentStep
`func (o *ExecutionDetailOut) UnsetCurrentStep()`

UnsetCurrentStep ensures that no value is present for CurrentStep, not even an explicit nil
### GetErrorMessage

`func (o *ExecutionDetailOut) GetErrorMessage() string`

GetErrorMessage returns the ErrorMessage field if non-nil, zero value otherwise.

### GetErrorMessageOk

`func (o *ExecutionDetailOut) GetErrorMessageOk() (*string, bool)`

GetErrorMessageOk returns a tuple with the ErrorMessage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrorMessage

`func (o *ExecutionDetailOut) SetErrorMessage(v string)`

SetErrorMessage sets ErrorMessage field to given value.

### HasErrorMessage

`func (o *ExecutionDetailOut) HasErrorMessage() bool`

HasErrorMessage returns a boolean if a field has been set.

### SetErrorMessageNil

`func (o *ExecutionDetailOut) SetErrorMessageNil(b bool)`

 SetErrorMessageNil sets the value for ErrorMessage to be an explicit nil

### UnsetErrorMessage
`func (o *ExecutionDetailOut) UnsetErrorMessage()`

UnsetErrorMessage ensures that no value is present for ErrorMessage, not even an explicit nil
### GetStartedAt

`func (o *ExecutionDetailOut) GetStartedAt() string`

GetStartedAt returns the StartedAt field if non-nil, zero value otherwise.

### GetStartedAtOk

`func (o *ExecutionDetailOut) GetStartedAtOk() (*string, bool)`

GetStartedAtOk returns a tuple with the StartedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartedAt

`func (o *ExecutionDetailOut) SetStartedAt(v string)`

SetStartedAt sets StartedAt field to given value.

### HasStartedAt

`func (o *ExecutionDetailOut) HasStartedAt() bool`

HasStartedAt returns a boolean if a field has been set.

### SetStartedAtNil

`func (o *ExecutionDetailOut) SetStartedAtNil(b bool)`

 SetStartedAtNil sets the value for StartedAt to be an explicit nil

### UnsetStartedAt
`func (o *ExecutionDetailOut) UnsetStartedAt()`

UnsetStartedAt ensures that no value is present for StartedAt, not even an explicit nil
### GetCompletedAt

`func (o *ExecutionDetailOut) GetCompletedAt() string`

GetCompletedAt returns the CompletedAt field if non-nil, zero value otherwise.

### GetCompletedAtOk

`func (o *ExecutionDetailOut) GetCompletedAtOk() (*string, bool)`

GetCompletedAtOk returns a tuple with the CompletedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompletedAt

`func (o *ExecutionDetailOut) SetCompletedAt(v string)`

SetCompletedAt sets CompletedAt field to given value.

### HasCompletedAt

`func (o *ExecutionDetailOut) HasCompletedAt() bool`

HasCompletedAt returns a boolean if a field has been set.

### SetCompletedAtNil

`func (o *ExecutionDetailOut) SetCompletedAtNil(b bool)`

 SetCompletedAtNil sets the value for CompletedAt to be an explicit nil

### UnsetCompletedAt
`func (o *ExecutionDetailOut) UnsetCompletedAt()`

UnsetCompletedAt ensures that no value is present for CompletedAt, not even an explicit nil
### GetDurationSeconds

`func (o *ExecutionDetailOut) GetDurationSeconds() float32`

GetDurationSeconds returns the DurationSeconds field if non-nil, zero value otherwise.

### GetDurationSecondsOk

`func (o *ExecutionDetailOut) GetDurationSecondsOk() (*float32, bool)`

GetDurationSecondsOk returns a tuple with the DurationSeconds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDurationSeconds

`func (o *ExecutionDetailOut) SetDurationSeconds(v float32)`

SetDurationSeconds sets DurationSeconds field to given value.

### HasDurationSeconds

`func (o *ExecutionDetailOut) HasDurationSeconds() bool`

HasDurationSeconds returns a boolean if a field has been set.

### SetDurationSecondsNil

`func (o *ExecutionDetailOut) SetDurationSecondsNil(b bool)`

 SetDurationSecondsNil sets the value for DurationSeconds to be an explicit nil

### UnsetDurationSeconds
`func (o *ExecutionDetailOut) UnsetDurationSeconds()`

UnsetDurationSeconds ensures that no value is present for DurationSeconds, not even an explicit nil
### GetSteps

`func (o *ExecutionDetailOut) GetSteps() []ExecutionStepDetailOut`

GetSteps returns the Steps field if non-nil, zero value otherwise.

### GetStepsOk

`func (o *ExecutionDetailOut) GetStepsOk() (*[]ExecutionStepDetailOut, bool)`

GetStepsOk returns a tuple with the Steps field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSteps

`func (o *ExecutionDetailOut) SetSteps(v []ExecutionStepDetailOut)`

SetSteps sets Steps field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


