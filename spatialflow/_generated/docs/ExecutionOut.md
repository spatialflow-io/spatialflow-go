# ExecutionOut

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** |  | 
**WorkflowId** | **string** |  | 
**WorkflowName** | **string** |  | 
**ExecutionId** | **string** |  | 
**TriggerSource** | **string** |  | 
**Status** | **string** |  | 
**StartedAt** | **NullableTime** |  | 
**CompletedAt** | **NullableTime** |  | 
**DurationSeconds** | **NullableFloat32** |  | 
**ErrorMessage** | **NullableString** |  | 

## Methods

### NewExecutionOut

`func NewExecutionOut(id string, workflowId string, workflowName string, executionId string, triggerSource string, status string, startedAt NullableTime, completedAt NullableTime, durationSeconds NullableFloat32, errorMessage NullableString, ) *ExecutionOut`

NewExecutionOut instantiates a new ExecutionOut object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewExecutionOutWithDefaults

`func NewExecutionOutWithDefaults() *ExecutionOut`

NewExecutionOutWithDefaults instantiates a new ExecutionOut object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ExecutionOut) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ExecutionOut) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ExecutionOut) SetId(v string)`

SetId sets Id field to given value.


### GetWorkflowId

`func (o *ExecutionOut) GetWorkflowId() string`

GetWorkflowId returns the WorkflowId field if non-nil, zero value otherwise.

### GetWorkflowIdOk

`func (o *ExecutionOut) GetWorkflowIdOk() (*string, bool)`

GetWorkflowIdOk returns a tuple with the WorkflowId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWorkflowId

`func (o *ExecutionOut) SetWorkflowId(v string)`

SetWorkflowId sets WorkflowId field to given value.


### GetWorkflowName

`func (o *ExecutionOut) GetWorkflowName() string`

GetWorkflowName returns the WorkflowName field if non-nil, zero value otherwise.

### GetWorkflowNameOk

`func (o *ExecutionOut) GetWorkflowNameOk() (*string, bool)`

GetWorkflowNameOk returns a tuple with the WorkflowName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWorkflowName

`func (o *ExecutionOut) SetWorkflowName(v string)`

SetWorkflowName sets WorkflowName field to given value.


### GetExecutionId

`func (o *ExecutionOut) GetExecutionId() string`

GetExecutionId returns the ExecutionId field if non-nil, zero value otherwise.

### GetExecutionIdOk

`func (o *ExecutionOut) GetExecutionIdOk() (*string, bool)`

GetExecutionIdOk returns a tuple with the ExecutionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExecutionId

`func (o *ExecutionOut) SetExecutionId(v string)`

SetExecutionId sets ExecutionId field to given value.


### GetTriggerSource

`func (o *ExecutionOut) GetTriggerSource() string`

GetTriggerSource returns the TriggerSource field if non-nil, zero value otherwise.

### GetTriggerSourceOk

`func (o *ExecutionOut) GetTriggerSourceOk() (*string, bool)`

GetTriggerSourceOk returns a tuple with the TriggerSource field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTriggerSource

`func (o *ExecutionOut) SetTriggerSource(v string)`

SetTriggerSource sets TriggerSource field to given value.


### GetStatus

`func (o *ExecutionOut) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *ExecutionOut) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *ExecutionOut) SetStatus(v string)`

SetStatus sets Status field to given value.


### GetStartedAt

`func (o *ExecutionOut) GetStartedAt() time.Time`

GetStartedAt returns the StartedAt field if non-nil, zero value otherwise.

### GetStartedAtOk

`func (o *ExecutionOut) GetStartedAtOk() (*time.Time, bool)`

GetStartedAtOk returns a tuple with the StartedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartedAt

`func (o *ExecutionOut) SetStartedAt(v time.Time)`

SetStartedAt sets StartedAt field to given value.


### SetStartedAtNil

`func (o *ExecutionOut) SetStartedAtNil(b bool)`

 SetStartedAtNil sets the value for StartedAt to be an explicit nil

### UnsetStartedAt
`func (o *ExecutionOut) UnsetStartedAt()`

UnsetStartedAt ensures that no value is present for StartedAt, not even an explicit nil
### GetCompletedAt

`func (o *ExecutionOut) GetCompletedAt() time.Time`

GetCompletedAt returns the CompletedAt field if non-nil, zero value otherwise.

### GetCompletedAtOk

`func (o *ExecutionOut) GetCompletedAtOk() (*time.Time, bool)`

GetCompletedAtOk returns a tuple with the CompletedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompletedAt

`func (o *ExecutionOut) SetCompletedAt(v time.Time)`

SetCompletedAt sets CompletedAt field to given value.


### SetCompletedAtNil

`func (o *ExecutionOut) SetCompletedAtNil(b bool)`

 SetCompletedAtNil sets the value for CompletedAt to be an explicit nil

### UnsetCompletedAt
`func (o *ExecutionOut) UnsetCompletedAt()`

UnsetCompletedAt ensures that no value is present for CompletedAt, not even an explicit nil
### GetDurationSeconds

`func (o *ExecutionOut) GetDurationSeconds() float32`

GetDurationSeconds returns the DurationSeconds field if non-nil, zero value otherwise.

### GetDurationSecondsOk

`func (o *ExecutionOut) GetDurationSecondsOk() (*float32, bool)`

GetDurationSecondsOk returns a tuple with the DurationSeconds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDurationSeconds

`func (o *ExecutionOut) SetDurationSeconds(v float32)`

SetDurationSeconds sets DurationSeconds field to given value.


### SetDurationSecondsNil

`func (o *ExecutionOut) SetDurationSecondsNil(b bool)`

 SetDurationSecondsNil sets the value for DurationSeconds to be an explicit nil

### UnsetDurationSeconds
`func (o *ExecutionOut) UnsetDurationSeconds()`

UnsetDurationSeconds ensures that no value is present for DurationSeconds, not even an explicit nil
### GetErrorMessage

`func (o *ExecutionOut) GetErrorMessage() string`

GetErrorMessage returns the ErrorMessage field if non-nil, zero value otherwise.

### GetErrorMessageOk

`func (o *ExecutionOut) GetErrorMessageOk() (*string, bool)`

GetErrorMessageOk returns a tuple with the ErrorMessage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrorMessage

`func (o *ExecutionOut) SetErrorMessage(v string)`

SetErrorMessage sets ErrorMessage field to given value.


### SetErrorMessageNil

`func (o *ExecutionOut) SetErrorMessageNil(b bool)`

 SetErrorMessageNil sets the value for ErrorMessage to be an explicit nil

### UnsetErrorMessage
`func (o *ExecutionOut) UnsetErrorMessage()`

UnsetErrorMessage ensures that no value is present for ErrorMessage, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


