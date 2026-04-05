# WorkflowStepRetrySchema

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**StepIndex** | **int32** |  | 
**StepName** | Pointer to **NullableString** |  | [optional] 
**ActionType** | **string** |  | 
**RetryConfig** | [**ActionRetryConfigSchema**](ActionRetryConfigSchema.md) |  | 

## Methods

### NewWorkflowStepRetrySchema

`func NewWorkflowStepRetrySchema(stepIndex int32, actionType string, retryConfig ActionRetryConfigSchema, ) *WorkflowStepRetrySchema`

NewWorkflowStepRetrySchema instantiates a new WorkflowStepRetrySchema object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWorkflowStepRetrySchemaWithDefaults

`func NewWorkflowStepRetrySchemaWithDefaults() *WorkflowStepRetrySchema`

NewWorkflowStepRetrySchemaWithDefaults instantiates a new WorkflowStepRetrySchema object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetStepIndex

`func (o *WorkflowStepRetrySchema) GetStepIndex() int32`

GetStepIndex returns the StepIndex field if non-nil, zero value otherwise.

### GetStepIndexOk

`func (o *WorkflowStepRetrySchema) GetStepIndexOk() (*int32, bool)`

GetStepIndexOk returns a tuple with the StepIndex field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStepIndex

`func (o *WorkflowStepRetrySchema) SetStepIndex(v int32)`

SetStepIndex sets StepIndex field to given value.


### GetStepName

`func (o *WorkflowStepRetrySchema) GetStepName() string`

GetStepName returns the StepName field if non-nil, zero value otherwise.

### GetStepNameOk

`func (o *WorkflowStepRetrySchema) GetStepNameOk() (*string, bool)`

GetStepNameOk returns a tuple with the StepName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStepName

`func (o *WorkflowStepRetrySchema) SetStepName(v string)`

SetStepName sets StepName field to given value.

### HasStepName

`func (o *WorkflowStepRetrySchema) HasStepName() bool`

HasStepName returns a boolean if a field has been set.

### SetStepNameNil

`func (o *WorkflowStepRetrySchema) SetStepNameNil(b bool)`

 SetStepNameNil sets the value for StepName to be an explicit nil

### UnsetStepName
`func (o *WorkflowStepRetrySchema) UnsetStepName()`

UnsetStepName ensures that no value is present for StepName, not even an explicit nil
### GetActionType

`func (o *WorkflowStepRetrySchema) GetActionType() string`

GetActionType returns the ActionType field if non-nil, zero value otherwise.

### GetActionTypeOk

`func (o *WorkflowStepRetrySchema) GetActionTypeOk() (*string, bool)`

GetActionTypeOk returns a tuple with the ActionType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActionType

`func (o *WorkflowStepRetrySchema) SetActionType(v string)`

SetActionType sets ActionType field to given value.


### GetRetryConfig

`func (o *WorkflowStepRetrySchema) GetRetryConfig() ActionRetryConfigSchema`

GetRetryConfig returns the RetryConfig field if non-nil, zero value otherwise.

### GetRetryConfigOk

`func (o *WorkflowStepRetrySchema) GetRetryConfigOk() (*ActionRetryConfigSchema, bool)`

GetRetryConfigOk returns a tuple with the RetryConfig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRetryConfig

`func (o *WorkflowStepRetrySchema) SetRetryConfig(v ActionRetryConfigSchema)`

SetRetryConfig sets RetryConfig field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


