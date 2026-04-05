# WorkflowOut

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** |  | 
**Name** | **string** |  | 
**Description** | **NullableString** |  | 
**Status** | **string** |  | 
**Nodes** | **[]map[string]interface{}** |  | 
**Edges** | **[]map[string]interface{}** |  | 
**CreatedAt** | **time.Time** |  | 
**UpdatedAt** | **time.Time** |  | 
**LastRun** | **NullableString** |  | 
**RunCount** | **int32** |  | 
**SuccessRate** | **float32** |  | 
**UserId** | **string** |  | 
**WorkspaceId** | Pointer to **NullableString** |  | [optional] 
**Version** | **int32** |  | 
**TemplateWarnings** | Pointer to [**NullableTemplateWarnings**](TemplateWarnings.md) |  | [optional] 

## Methods

### NewWorkflowOut

`func NewWorkflowOut(id string, name string, description NullableString, status string, nodes []*map[string]interface{}, edges []*map[string]interface{}, createdAt time.Time, updatedAt time.Time, lastRun NullableString, runCount int32, successRate float32, userId string, version int32, ) *WorkflowOut`

NewWorkflowOut instantiates a new WorkflowOut object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWorkflowOutWithDefaults

`func NewWorkflowOutWithDefaults() *WorkflowOut`

NewWorkflowOutWithDefaults instantiates a new WorkflowOut object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *WorkflowOut) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *WorkflowOut) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *WorkflowOut) SetId(v string)`

SetId sets Id field to given value.


### GetName

`func (o *WorkflowOut) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *WorkflowOut) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *WorkflowOut) SetName(v string)`

SetName sets Name field to given value.


### GetDescription

`func (o *WorkflowOut) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *WorkflowOut) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *WorkflowOut) SetDescription(v string)`

SetDescription sets Description field to given value.


### SetDescriptionNil

`func (o *WorkflowOut) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *WorkflowOut) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetStatus

`func (o *WorkflowOut) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *WorkflowOut) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *WorkflowOut) SetStatus(v string)`

SetStatus sets Status field to given value.


### GetNodes

`func (o *WorkflowOut) GetNodes() []*map[string]interface{}`

GetNodes returns the Nodes field if non-nil, zero value otherwise.

### GetNodesOk

`func (o *WorkflowOut) GetNodesOk() (*[]*map[string]interface{}, bool)`

GetNodesOk returns a tuple with the Nodes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNodes

`func (o *WorkflowOut) SetNodes(v []*map[string]interface{})`

SetNodes sets Nodes field to given value.


### GetEdges

`func (o *WorkflowOut) GetEdges() []*map[string]interface{}`

GetEdges returns the Edges field if non-nil, zero value otherwise.

### GetEdgesOk

`func (o *WorkflowOut) GetEdgesOk() (*[]*map[string]interface{}, bool)`

GetEdgesOk returns a tuple with the Edges field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEdges

`func (o *WorkflowOut) SetEdges(v []*map[string]interface{})`

SetEdges sets Edges field to given value.


### GetCreatedAt

`func (o *WorkflowOut) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *WorkflowOut) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *WorkflowOut) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.


### GetUpdatedAt

`func (o *WorkflowOut) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *WorkflowOut) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *WorkflowOut) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.


### GetLastRun

`func (o *WorkflowOut) GetLastRun() string`

GetLastRun returns the LastRun field if non-nil, zero value otherwise.

### GetLastRunOk

`func (o *WorkflowOut) GetLastRunOk() (*string, bool)`

GetLastRunOk returns a tuple with the LastRun field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastRun

`func (o *WorkflowOut) SetLastRun(v string)`

SetLastRun sets LastRun field to given value.


### SetLastRunNil

`func (o *WorkflowOut) SetLastRunNil(b bool)`

 SetLastRunNil sets the value for LastRun to be an explicit nil

### UnsetLastRun
`func (o *WorkflowOut) UnsetLastRun()`

UnsetLastRun ensures that no value is present for LastRun, not even an explicit nil
### GetRunCount

`func (o *WorkflowOut) GetRunCount() int32`

GetRunCount returns the RunCount field if non-nil, zero value otherwise.

### GetRunCountOk

`func (o *WorkflowOut) GetRunCountOk() (*int32, bool)`

GetRunCountOk returns a tuple with the RunCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRunCount

`func (o *WorkflowOut) SetRunCount(v int32)`

SetRunCount sets RunCount field to given value.


### GetSuccessRate

`func (o *WorkflowOut) GetSuccessRate() float32`

GetSuccessRate returns the SuccessRate field if non-nil, zero value otherwise.

### GetSuccessRateOk

`func (o *WorkflowOut) GetSuccessRateOk() (*float32, bool)`

GetSuccessRateOk returns a tuple with the SuccessRate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSuccessRate

`func (o *WorkflowOut) SetSuccessRate(v float32)`

SetSuccessRate sets SuccessRate field to given value.


### GetUserId

`func (o *WorkflowOut) GetUserId() string`

GetUserId returns the UserId field if non-nil, zero value otherwise.

### GetUserIdOk

`func (o *WorkflowOut) GetUserIdOk() (*string, bool)`

GetUserIdOk returns a tuple with the UserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserId

`func (o *WorkflowOut) SetUserId(v string)`

SetUserId sets UserId field to given value.


### GetWorkspaceId

`func (o *WorkflowOut) GetWorkspaceId() string`

GetWorkspaceId returns the WorkspaceId field if non-nil, zero value otherwise.

### GetWorkspaceIdOk

`func (o *WorkflowOut) GetWorkspaceIdOk() (*string, bool)`

GetWorkspaceIdOk returns a tuple with the WorkspaceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWorkspaceId

`func (o *WorkflowOut) SetWorkspaceId(v string)`

SetWorkspaceId sets WorkspaceId field to given value.

### HasWorkspaceId

`func (o *WorkflowOut) HasWorkspaceId() bool`

HasWorkspaceId returns a boolean if a field has been set.

### SetWorkspaceIdNil

`func (o *WorkflowOut) SetWorkspaceIdNil(b bool)`

 SetWorkspaceIdNil sets the value for WorkspaceId to be an explicit nil

### UnsetWorkspaceId
`func (o *WorkflowOut) UnsetWorkspaceId()`

UnsetWorkspaceId ensures that no value is present for WorkspaceId, not even an explicit nil
### GetVersion

`func (o *WorkflowOut) GetVersion() int32`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *WorkflowOut) GetVersionOk() (*int32, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *WorkflowOut) SetVersion(v int32)`

SetVersion sets Version field to given value.


### GetTemplateWarnings

`func (o *WorkflowOut) GetTemplateWarnings() TemplateWarnings`

GetTemplateWarnings returns the TemplateWarnings field if non-nil, zero value otherwise.

### GetTemplateWarningsOk

`func (o *WorkflowOut) GetTemplateWarningsOk() (*TemplateWarnings, bool)`

GetTemplateWarningsOk returns a tuple with the TemplateWarnings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTemplateWarnings

`func (o *WorkflowOut) SetTemplateWarnings(v TemplateWarnings)`

SetTemplateWarnings sets TemplateWarnings field to given value.

### HasTemplateWarnings

`func (o *WorkflowOut) HasTemplateWarnings() bool`

HasTemplateWarnings returns a boolean if a field has been set.

### SetTemplateWarningsNil

`func (o *WorkflowOut) SetTemplateWarningsNil(b bool)`

 SetTemplateWarningsNil sets the value for TemplateWarnings to be an explicit nil

### UnsetTemplateWarnings
`func (o *WorkflowOut) UnsetTemplateWarnings()`

UnsetTemplateWarnings ensures that no value is present for TemplateWarnings, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


