# WorkflowRetryPolicyUpdateSchema

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**WorkflowId** | **string** |  | 
**DefaultRetryPolicy** | Pointer to **NullableString** |  | [optional] 
**StepOverrides** | Pointer to [**[]WorkflowStepRetrySchema**](WorkflowStepRetrySchema.md) |  | [optional] 

## Methods

### NewWorkflowRetryPolicyUpdateSchema

`func NewWorkflowRetryPolicyUpdateSchema(workflowId string, ) *WorkflowRetryPolicyUpdateSchema`

NewWorkflowRetryPolicyUpdateSchema instantiates a new WorkflowRetryPolicyUpdateSchema object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWorkflowRetryPolicyUpdateSchemaWithDefaults

`func NewWorkflowRetryPolicyUpdateSchemaWithDefaults() *WorkflowRetryPolicyUpdateSchema`

NewWorkflowRetryPolicyUpdateSchemaWithDefaults instantiates a new WorkflowRetryPolicyUpdateSchema object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetWorkflowId

`func (o *WorkflowRetryPolicyUpdateSchema) GetWorkflowId() string`

GetWorkflowId returns the WorkflowId field if non-nil, zero value otherwise.

### GetWorkflowIdOk

`func (o *WorkflowRetryPolicyUpdateSchema) GetWorkflowIdOk() (*string, bool)`

GetWorkflowIdOk returns a tuple with the WorkflowId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWorkflowId

`func (o *WorkflowRetryPolicyUpdateSchema) SetWorkflowId(v string)`

SetWorkflowId sets WorkflowId field to given value.


### GetDefaultRetryPolicy

`func (o *WorkflowRetryPolicyUpdateSchema) GetDefaultRetryPolicy() string`

GetDefaultRetryPolicy returns the DefaultRetryPolicy field if non-nil, zero value otherwise.

### GetDefaultRetryPolicyOk

`func (o *WorkflowRetryPolicyUpdateSchema) GetDefaultRetryPolicyOk() (*string, bool)`

GetDefaultRetryPolicyOk returns a tuple with the DefaultRetryPolicy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultRetryPolicy

`func (o *WorkflowRetryPolicyUpdateSchema) SetDefaultRetryPolicy(v string)`

SetDefaultRetryPolicy sets DefaultRetryPolicy field to given value.

### HasDefaultRetryPolicy

`func (o *WorkflowRetryPolicyUpdateSchema) HasDefaultRetryPolicy() bool`

HasDefaultRetryPolicy returns a boolean if a field has been set.

### SetDefaultRetryPolicyNil

`func (o *WorkflowRetryPolicyUpdateSchema) SetDefaultRetryPolicyNil(b bool)`

 SetDefaultRetryPolicyNil sets the value for DefaultRetryPolicy to be an explicit nil

### UnsetDefaultRetryPolicy
`func (o *WorkflowRetryPolicyUpdateSchema) UnsetDefaultRetryPolicy()`

UnsetDefaultRetryPolicy ensures that no value is present for DefaultRetryPolicy, not even an explicit nil
### GetStepOverrides

`func (o *WorkflowRetryPolicyUpdateSchema) GetStepOverrides() []WorkflowStepRetrySchema`

GetStepOverrides returns the StepOverrides field if non-nil, zero value otherwise.

### GetStepOverridesOk

`func (o *WorkflowRetryPolicyUpdateSchema) GetStepOverridesOk() (*[]WorkflowStepRetrySchema, bool)`

GetStepOverridesOk returns a tuple with the StepOverrides field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStepOverrides

`func (o *WorkflowRetryPolicyUpdateSchema) SetStepOverrides(v []WorkflowStepRetrySchema)`

SetStepOverrides sets StepOverrides field to given value.

### HasStepOverrides

`func (o *WorkflowRetryPolicyUpdateSchema) HasStepOverrides() bool`

HasStepOverrides returns a boolean if a field has been set.

### SetStepOverridesNil

`func (o *WorkflowRetryPolicyUpdateSchema) SetStepOverridesNil(b bool)`

 SetStepOverridesNil sets the value for StepOverrides to be an explicit nil

### UnsetStepOverrides
`func (o *WorkflowRetryPolicyUpdateSchema) UnsetStepOverrides()`

UnsetStepOverrides ensures that no value is present for StepOverrides, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


