# UpdateIntegrationSchema

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **NullableString** |  | [optional] 
**Description** | Pointer to **NullableString** |  | [optional] 
**Config** | Pointer to **map[string]interface{}** |  | [optional] 
**Tags** | Pointer to **[]string** |  | [optional] 
**IsActive** | Pointer to **NullableBool** |  | [optional] 

## Methods

### NewUpdateIntegrationSchema

`func NewUpdateIntegrationSchema() *UpdateIntegrationSchema`

NewUpdateIntegrationSchema instantiates a new UpdateIntegrationSchema object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateIntegrationSchemaWithDefaults

`func NewUpdateIntegrationSchemaWithDefaults() *UpdateIntegrationSchema`

NewUpdateIntegrationSchemaWithDefaults instantiates a new UpdateIntegrationSchema object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *UpdateIntegrationSchema) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *UpdateIntegrationSchema) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *UpdateIntegrationSchema) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *UpdateIntegrationSchema) HasName() bool`

HasName returns a boolean if a field has been set.

### SetNameNil

`func (o *UpdateIntegrationSchema) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *UpdateIntegrationSchema) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetDescription

`func (o *UpdateIntegrationSchema) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *UpdateIntegrationSchema) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *UpdateIntegrationSchema) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *UpdateIntegrationSchema) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescriptionNil

`func (o *UpdateIntegrationSchema) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *UpdateIntegrationSchema) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetConfig

`func (o *UpdateIntegrationSchema) GetConfig() map[string]interface{}`

GetConfig returns the Config field if non-nil, zero value otherwise.

### GetConfigOk

`func (o *UpdateIntegrationSchema) GetConfigOk() (*map[string]interface{}, bool)`

GetConfigOk returns a tuple with the Config field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfig

`func (o *UpdateIntegrationSchema) SetConfig(v map[string]interface{})`

SetConfig sets Config field to given value.

### HasConfig

`func (o *UpdateIntegrationSchema) HasConfig() bool`

HasConfig returns a boolean if a field has been set.

### SetConfigNil

`func (o *UpdateIntegrationSchema) SetConfigNil(b bool)`

 SetConfigNil sets the value for Config to be an explicit nil

### UnsetConfig
`func (o *UpdateIntegrationSchema) UnsetConfig()`

UnsetConfig ensures that no value is present for Config, not even an explicit nil
### GetTags

`func (o *UpdateIntegrationSchema) GetTags() []string`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *UpdateIntegrationSchema) GetTagsOk() (*[]string, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *UpdateIntegrationSchema) SetTags(v []string)`

SetTags sets Tags field to given value.

### HasTags

`func (o *UpdateIntegrationSchema) HasTags() bool`

HasTags returns a boolean if a field has been set.

### SetTagsNil

`func (o *UpdateIntegrationSchema) SetTagsNil(b bool)`

 SetTagsNil sets the value for Tags to be an explicit nil

### UnsetTags
`func (o *UpdateIntegrationSchema) UnsetTags()`

UnsetTags ensures that no value is present for Tags, not even an explicit nil
### GetIsActive

`func (o *UpdateIntegrationSchema) GetIsActive() bool`

GetIsActive returns the IsActive field if non-nil, zero value otherwise.

### GetIsActiveOk

`func (o *UpdateIntegrationSchema) GetIsActiveOk() (*bool, bool)`

GetIsActiveOk returns a tuple with the IsActive field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsActive

`func (o *UpdateIntegrationSchema) SetIsActive(v bool)`

SetIsActive sets IsActive field to given value.

### HasIsActive

`func (o *UpdateIntegrationSchema) HasIsActive() bool`

HasIsActive returns a boolean if a field has been set.

### SetIsActiveNil

`func (o *UpdateIntegrationSchema) SetIsActiveNil(b bool)`

 SetIsActiveNil sets the value for IsActive to be an explicit nil

### UnsetIsActive
`func (o *UpdateIntegrationSchema) UnsetIsActive()`

UnsetIsActive ensures that no value is present for IsActive, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


