# TestEventItemOut

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**TestEventId** | **string** |  | 
**EventType** | **string** |  | 
**WebhookTriggered** | **bool** |  | 
**WorkflowTriggered** | **bool** |  | 
**ExecutionResults** | Pointer to **map[string]interface{}** |  | [optional] 
**TestMetadata** | Pointer to **map[string]interface{}** |  | [optional] 
**CreatedAt** | **string** |  | 

## Methods

### NewTestEventItemOut

`func NewTestEventItemOut(testEventId string, eventType string, webhookTriggered bool, workflowTriggered bool, createdAt string, ) *TestEventItemOut`

NewTestEventItemOut instantiates a new TestEventItemOut object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTestEventItemOutWithDefaults

`func NewTestEventItemOutWithDefaults() *TestEventItemOut`

NewTestEventItemOutWithDefaults instantiates a new TestEventItemOut object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTestEventId

`func (o *TestEventItemOut) GetTestEventId() string`

GetTestEventId returns the TestEventId field if non-nil, zero value otherwise.

### GetTestEventIdOk

`func (o *TestEventItemOut) GetTestEventIdOk() (*string, bool)`

GetTestEventIdOk returns a tuple with the TestEventId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTestEventId

`func (o *TestEventItemOut) SetTestEventId(v string)`

SetTestEventId sets TestEventId field to given value.


### GetEventType

`func (o *TestEventItemOut) GetEventType() string`

GetEventType returns the EventType field if non-nil, zero value otherwise.

### GetEventTypeOk

`func (o *TestEventItemOut) GetEventTypeOk() (*string, bool)`

GetEventTypeOk returns a tuple with the EventType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEventType

`func (o *TestEventItemOut) SetEventType(v string)`

SetEventType sets EventType field to given value.


### GetWebhookTriggered

`func (o *TestEventItemOut) GetWebhookTriggered() bool`

GetWebhookTriggered returns the WebhookTriggered field if non-nil, zero value otherwise.

### GetWebhookTriggeredOk

`func (o *TestEventItemOut) GetWebhookTriggeredOk() (*bool, bool)`

GetWebhookTriggeredOk returns a tuple with the WebhookTriggered field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWebhookTriggered

`func (o *TestEventItemOut) SetWebhookTriggered(v bool)`

SetWebhookTriggered sets WebhookTriggered field to given value.


### GetWorkflowTriggered

`func (o *TestEventItemOut) GetWorkflowTriggered() bool`

GetWorkflowTriggered returns the WorkflowTriggered field if non-nil, zero value otherwise.

### GetWorkflowTriggeredOk

`func (o *TestEventItemOut) GetWorkflowTriggeredOk() (*bool, bool)`

GetWorkflowTriggeredOk returns a tuple with the WorkflowTriggered field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWorkflowTriggered

`func (o *TestEventItemOut) SetWorkflowTriggered(v bool)`

SetWorkflowTriggered sets WorkflowTriggered field to given value.


### GetExecutionResults

`func (o *TestEventItemOut) GetExecutionResults() map[string]interface{}`

GetExecutionResults returns the ExecutionResults field if non-nil, zero value otherwise.

### GetExecutionResultsOk

`func (o *TestEventItemOut) GetExecutionResultsOk() (*map[string]interface{}, bool)`

GetExecutionResultsOk returns a tuple with the ExecutionResults field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExecutionResults

`func (o *TestEventItemOut) SetExecutionResults(v map[string]interface{})`

SetExecutionResults sets ExecutionResults field to given value.

### HasExecutionResults

`func (o *TestEventItemOut) HasExecutionResults() bool`

HasExecutionResults returns a boolean if a field has been set.

### SetExecutionResultsNil

`func (o *TestEventItemOut) SetExecutionResultsNil(b bool)`

 SetExecutionResultsNil sets the value for ExecutionResults to be an explicit nil

### UnsetExecutionResults
`func (o *TestEventItemOut) UnsetExecutionResults()`

UnsetExecutionResults ensures that no value is present for ExecutionResults, not even an explicit nil
### GetTestMetadata

`func (o *TestEventItemOut) GetTestMetadata() map[string]interface{}`

GetTestMetadata returns the TestMetadata field if non-nil, zero value otherwise.

### GetTestMetadataOk

`func (o *TestEventItemOut) GetTestMetadataOk() (*map[string]interface{}, bool)`

GetTestMetadataOk returns a tuple with the TestMetadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTestMetadata

`func (o *TestEventItemOut) SetTestMetadata(v map[string]interface{})`

SetTestMetadata sets TestMetadata field to given value.

### HasTestMetadata

`func (o *TestEventItemOut) HasTestMetadata() bool`

HasTestMetadata returns a boolean if a field has been set.

### SetTestMetadataNil

`func (o *TestEventItemOut) SetTestMetadataNil(b bool)`

 SetTestMetadataNil sets the value for TestMetadata to be an explicit nil

### UnsetTestMetadata
`func (o *TestEventItemOut) UnsetTestMetadata()`

UnsetTestMetadata ensures that no value is present for TestMetadata, not even an explicit nil
### GetCreatedAt

`func (o *TestEventItemOut) GetCreatedAt() string`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *TestEventItemOut) GetCreatedAtOk() (*string, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *TestEventItemOut) SetCreatedAt(v string)`

SetCreatedAt sets CreatedAt field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


