# WorkflowImportDataSchema

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** |  | 
**Description** | Pointer to **NullableString** |  | [optional] 
**TriggerConfig** | Pointer to **map[string]interface{}** |  | [optional] 
**Steps** | Pointer to **[]map[string]interface{}** |  | [optional] 
**Tags** | Pointer to **[]string** |  | [optional] 
**Category** | Pointer to **NullableString** |  | [optional] 
**IsTestMode** | Pointer to **NullableBool** |  | [optional] 

## Methods

### NewWorkflowImportDataSchema

`func NewWorkflowImportDataSchema(name string, ) *WorkflowImportDataSchema`

NewWorkflowImportDataSchema instantiates a new WorkflowImportDataSchema object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWorkflowImportDataSchemaWithDefaults

`func NewWorkflowImportDataSchemaWithDefaults() *WorkflowImportDataSchema`

NewWorkflowImportDataSchemaWithDefaults instantiates a new WorkflowImportDataSchema object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *WorkflowImportDataSchema) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *WorkflowImportDataSchema) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *WorkflowImportDataSchema) SetName(v string)`

SetName sets Name field to given value.


### GetDescription

`func (o *WorkflowImportDataSchema) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *WorkflowImportDataSchema) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *WorkflowImportDataSchema) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *WorkflowImportDataSchema) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescriptionNil

`func (o *WorkflowImportDataSchema) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *WorkflowImportDataSchema) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetTriggerConfig

`func (o *WorkflowImportDataSchema) GetTriggerConfig() map[string]interface{}`

GetTriggerConfig returns the TriggerConfig field if non-nil, zero value otherwise.

### GetTriggerConfigOk

`func (o *WorkflowImportDataSchema) GetTriggerConfigOk() (*map[string]interface{}, bool)`

GetTriggerConfigOk returns a tuple with the TriggerConfig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTriggerConfig

`func (o *WorkflowImportDataSchema) SetTriggerConfig(v map[string]interface{})`

SetTriggerConfig sets TriggerConfig field to given value.

### HasTriggerConfig

`func (o *WorkflowImportDataSchema) HasTriggerConfig() bool`

HasTriggerConfig returns a boolean if a field has been set.

### SetTriggerConfigNil

`func (o *WorkflowImportDataSchema) SetTriggerConfigNil(b bool)`

 SetTriggerConfigNil sets the value for TriggerConfig to be an explicit nil

### UnsetTriggerConfig
`func (o *WorkflowImportDataSchema) UnsetTriggerConfig()`

UnsetTriggerConfig ensures that no value is present for TriggerConfig, not even an explicit nil
### GetSteps

`func (o *WorkflowImportDataSchema) GetSteps() []map[string]interface{}`

GetSteps returns the Steps field if non-nil, zero value otherwise.

### GetStepsOk

`func (o *WorkflowImportDataSchema) GetStepsOk() (*[]map[string]interface{}, bool)`

GetStepsOk returns a tuple with the Steps field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSteps

`func (o *WorkflowImportDataSchema) SetSteps(v []map[string]interface{})`

SetSteps sets Steps field to given value.

### HasSteps

`func (o *WorkflowImportDataSchema) HasSteps() bool`

HasSteps returns a boolean if a field has been set.

### SetStepsNil

`func (o *WorkflowImportDataSchema) SetStepsNil(b bool)`

 SetStepsNil sets the value for Steps to be an explicit nil

### UnsetSteps
`func (o *WorkflowImportDataSchema) UnsetSteps()`

UnsetSteps ensures that no value is present for Steps, not even an explicit nil
### GetTags

`func (o *WorkflowImportDataSchema) GetTags() []string`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *WorkflowImportDataSchema) GetTagsOk() (*[]string, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *WorkflowImportDataSchema) SetTags(v []string)`

SetTags sets Tags field to given value.

### HasTags

`func (o *WorkflowImportDataSchema) HasTags() bool`

HasTags returns a boolean if a field has been set.

### SetTagsNil

`func (o *WorkflowImportDataSchema) SetTagsNil(b bool)`

 SetTagsNil sets the value for Tags to be an explicit nil

### UnsetTags
`func (o *WorkflowImportDataSchema) UnsetTags()`

UnsetTags ensures that no value is present for Tags, not even an explicit nil
### GetCategory

`func (o *WorkflowImportDataSchema) GetCategory() string`

GetCategory returns the Category field if non-nil, zero value otherwise.

### GetCategoryOk

`func (o *WorkflowImportDataSchema) GetCategoryOk() (*string, bool)`

GetCategoryOk returns a tuple with the Category field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCategory

`func (o *WorkflowImportDataSchema) SetCategory(v string)`

SetCategory sets Category field to given value.

### HasCategory

`func (o *WorkflowImportDataSchema) HasCategory() bool`

HasCategory returns a boolean if a field has been set.

### SetCategoryNil

`func (o *WorkflowImportDataSchema) SetCategoryNil(b bool)`

 SetCategoryNil sets the value for Category to be an explicit nil

### UnsetCategory
`func (o *WorkflowImportDataSchema) UnsetCategory()`

UnsetCategory ensures that no value is present for Category, not even an explicit nil
### GetIsTestMode

`func (o *WorkflowImportDataSchema) GetIsTestMode() bool`

GetIsTestMode returns the IsTestMode field if non-nil, zero value otherwise.

### GetIsTestModeOk

`func (o *WorkflowImportDataSchema) GetIsTestModeOk() (*bool, bool)`

GetIsTestModeOk returns a tuple with the IsTestMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsTestMode

`func (o *WorkflowImportDataSchema) SetIsTestMode(v bool)`

SetIsTestMode sets IsTestMode field to given value.

### HasIsTestMode

`func (o *WorkflowImportDataSchema) HasIsTestMode() bool`

HasIsTestMode returns a boolean if a field has been set.

### SetIsTestModeNil

`func (o *WorkflowImportDataSchema) SetIsTestModeNil(b bool)`

 SetIsTestModeNil sets the value for IsTestMode to be an explicit nil

### UnsetIsTestMode
`func (o *WorkflowImportDataSchema) UnsetIsTestMode()`

UnsetIsTestMode ensures that no value is present for IsTestMode, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


