# WorkflowIn

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** |  | 
**Description** | Pointer to **NullableString** |  | [optional] 
**Nodes** | **[]map[string]interface{}** |  | 
**Edges** | **[]map[string]interface{}** |  | 

## Methods

### NewWorkflowIn

`func NewWorkflowIn(name string, nodes []map[string]interface{}, edges []map[string]interface{}, ) *WorkflowIn`

NewWorkflowIn instantiates a new WorkflowIn object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWorkflowInWithDefaults

`func NewWorkflowInWithDefaults() *WorkflowIn`

NewWorkflowInWithDefaults instantiates a new WorkflowIn object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *WorkflowIn) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *WorkflowIn) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *WorkflowIn) SetName(v string)`

SetName sets Name field to given value.


### GetDescription

`func (o *WorkflowIn) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *WorkflowIn) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *WorkflowIn) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *WorkflowIn) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescriptionNil

`func (o *WorkflowIn) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *WorkflowIn) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetNodes

`func (o *WorkflowIn) GetNodes() []map[string]interface{}`

GetNodes returns the Nodes field if non-nil, zero value otherwise.

### GetNodesOk

`func (o *WorkflowIn) GetNodesOk() (*[]map[string]interface{}, bool)`

GetNodesOk returns a tuple with the Nodes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNodes

`func (o *WorkflowIn) SetNodes(v []map[string]interface{})`

SetNodes sets Nodes field to given value.


### GetEdges

`func (o *WorkflowIn) GetEdges() []map[string]interface{}`

GetEdges returns the Edges field if non-nil, zero value otherwise.

### GetEdgesOk

`func (o *WorkflowIn) GetEdgesOk() (*[]map[string]interface{}, bool)`

GetEdgesOk returns a tuple with the Edges field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEdges

`func (o *WorkflowIn) SetEdges(v []map[string]interface{})`

SetEdges sets Edges field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


