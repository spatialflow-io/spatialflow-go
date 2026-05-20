# WorkflowTemplateDetailOut

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** |  | 
**Name** | **string** |  | 
**Description** | **string** |  | 
**Category** | **string** |  | 
**Lane** | **string** |  | 
**TriggerConfig** | **map[string]interface{}** |  | 
**Steps** | **[]interface{}** |  | 
**Tags** | Pointer to **[]string** |  | [optional] 
**IsFeatured** | **bool** |  | 
**UsageCount** | **int32** |  | 

## Methods

### NewWorkflowTemplateDetailOut

`func NewWorkflowTemplateDetailOut(id string, name string, description string, category string, lane string, triggerConfig map[string]interface{}, steps []interface{}, isFeatured bool, usageCount int32, ) *WorkflowTemplateDetailOut`

NewWorkflowTemplateDetailOut instantiates a new WorkflowTemplateDetailOut object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWorkflowTemplateDetailOutWithDefaults

`func NewWorkflowTemplateDetailOutWithDefaults() *WorkflowTemplateDetailOut`

NewWorkflowTemplateDetailOutWithDefaults instantiates a new WorkflowTemplateDetailOut object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *WorkflowTemplateDetailOut) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *WorkflowTemplateDetailOut) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *WorkflowTemplateDetailOut) SetId(v string)`

SetId sets Id field to given value.


### GetName

`func (o *WorkflowTemplateDetailOut) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *WorkflowTemplateDetailOut) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *WorkflowTemplateDetailOut) SetName(v string)`

SetName sets Name field to given value.


### GetDescription

`func (o *WorkflowTemplateDetailOut) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *WorkflowTemplateDetailOut) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *WorkflowTemplateDetailOut) SetDescription(v string)`

SetDescription sets Description field to given value.


### GetCategory

`func (o *WorkflowTemplateDetailOut) GetCategory() string`

GetCategory returns the Category field if non-nil, zero value otherwise.

### GetCategoryOk

`func (o *WorkflowTemplateDetailOut) GetCategoryOk() (*string, bool)`

GetCategoryOk returns a tuple with the Category field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCategory

`func (o *WorkflowTemplateDetailOut) SetCategory(v string)`

SetCategory sets Category field to given value.


### GetLane

`func (o *WorkflowTemplateDetailOut) GetLane() string`

GetLane returns the Lane field if non-nil, zero value otherwise.

### GetLaneOk

`func (o *WorkflowTemplateDetailOut) GetLaneOk() (*string, bool)`

GetLaneOk returns a tuple with the Lane field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLane

`func (o *WorkflowTemplateDetailOut) SetLane(v string)`

SetLane sets Lane field to given value.


### GetTriggerConfig

`func (o *WorkflowTemplateDetailOut) GetTriggerConfig() map[string]interface{}`

GetTriggerConfig returns the TriggerConfig field if non-nil, zero value otherwise.

### GetTriggerConfigOk

`func (o *WorkflowTemplateDetailOut) GetTriggerConfigOk() (*map[string]interface{}, bool)`

GetTriggerConfigOk returns a tuple with the TriggerConfig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTriggerConfig

`func (o *WorkflowTemplateDetailOut) SetTriggerConfig(v map[string]interface{})`

SetTriggerConfig sets TriggerConfig field to given value.


### GetSteps

`func (o *WorkflowTemplateDetailOut) GetSteps() []interface{}`

GetSteps returns the Steps field if non-nil, zero value otherwise.

### GetStepsOk

`func (o *WorkflowTemplateDetailOut) GetStepsOk() (*[]interface{}, bool)`

GetStepsOk returns a tuple with the Steps field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSteps

`func (o *WorkflowTemplateDetailOut) SetSteps(v []interface{})`

SetSteps sets Steps field to given value.


### GetTags

`func (o *WorkflowTemplateDetailOut) GetTags() []string`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *WorkflowTemplateDetailOut) GetTagsOk() (*[]string, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *WorkflowTemplateDetailOut) SetTags(v []string)`

SetTags sets Tags field to given value.

### HasTags

`func (o *WorkflowTemplateDetailOut) HasTags() bool`

HasTags returns a boolean if a field has been set.

### SetTagsNil

`func (o *WorkflowTemplateDetailOut) SetTagsNil(b bool)`

 SetTagsNil sets the value for Tags to be an explicit nil

### UnsetTags
`func (o *WorkflowTemplateDetailOut) UnsetTags()`

UnsetTags ensures that no value is present for Tags, not even an explicit nil
### GetIsFeatured

`func (o *WorkflowTemplateDetailOut) GetIsFeatured() bool`

GetIsFeatured returns the IsFeatured field if non-nil, zero value otherwise.

### GetIsFeaturedOk

`func (o *WorkflowTemplateDetailOut) GetIsFeaturedOk() (*bool, bool)`

GetIsFeaturedOk returns a tuple with the IsFeatured field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsFeatured

`func (o *WorkflowTemplateDetailOut) SetIsFeatured(v bool)`

SetIsFeatured sets IsFeatured field to given value.


### GetUsageCount

`func (o *WorkflowTemplateDetailOut) GetUsageCount() int32`

GetUsageCount returns the UsageCount field if non-nil, zero value otherwise.

### GetUsageCountOk

`func (o *WorkflowTemplateDetailOut) GetUsageCountOk() (*int32, bool)`

GetUsageCountOk returns a tuple with the UsageCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsageCount

`func (o *WorkflowTemplateDetailOut) SetUsageCount(v int32)`

SetUsageCount sets UsageCount field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


