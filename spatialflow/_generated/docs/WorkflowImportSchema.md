# WorkflowImportSchema

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Version** | **string** |  | 
**Workflow** | [**WorkflowImportDataSchema**](WorkflowImportDataSchema.md) |  | 

## Methods

### NewWorkflowImportSchema

`func NewWorkflowImportSchema(version string, workflow WorkflowImportDataSchema, ) *WorkflowImportSchema`

NewWorkflowImportSchema instantiates a new WorkflowImportSchema object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWorkflowImportSchemaWithDefaults

`func NewWorkflowImportSchemaWithDefaults() *WorkflowImportSchema`

NewWorkflowImportSchemaWithDefaults instantiates a new WorkflowImportSchema object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetVersion

`func (o *WorkflowImportSchema) GetVersion() string`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *WorkflowImportSchema) GetVersionOk() (*string, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *WorkflowImportSchema) SetVersion(v string)`

SetVersion sets Version field to given value.


### GetWorkflow

`func (o *WorkflowImportSchema) GetWorkflow() WorkflowImportDataSchema`

GetWorkflow returns the Workflow field if non-nil, zero value otherwise.

### GetWorkflowOk

`func (o *WorkflowImportSchema) GetWorkflowOk() (*WorkflowImportDataSchema, bool)`

GetWorkflowOk returns a tuple with the Workflow field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWorkflow

`func (o *WorkflowImportSchema) SetWorkflow(v WorkflowImportDataSchema)`

SetWorkflow sets Workflow field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


