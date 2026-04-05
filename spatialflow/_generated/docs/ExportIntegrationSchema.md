# ExportIntegrationSchema

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** |  | 
**Type** | **string** |  | 
**Description** | **string** |  | 
**Config** | **map[string]interface{}** |  | 
**Tags** | **[]string** |  | 
**Version** | Pointer to **string** |  | [optional] [default to "1.0"]

## Methods

### NewExportIntegrationSchema

`func NewExportIntegrationSchema(name string, type_ string, description string, config map[string]interface{}, tags []string, ) *ExportIntegrationSchema`

NewExportIntegrationSchema instantiates a new ExportIntegrationSchema object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewExportIntegrationSchemaWithDefaults

`func NewExportIntegrationSchemaWithDefaults() *ExportIntegrationSchema`

NewExportIntegrationSchemaWithDefaults instantiates a new ExportIntegrationSchema object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *ExportIntegrationSchema) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ExportIntegrationSchema) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ExportIntegrationSchema) SetName(v string)`

SetName sets Name field to given value.


### GetType

`func (o *ExportIntegrationSchema) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *ExportIntegrationSchema) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *ExportIntegrationSchema) SetType(v string)`

SetType sets Type field to given value.


### GetDescription

`func (o *ExportIntegrationSchema) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *ExportIntegrationSchema) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *ExportIntegrationSchema) SetDescription(v string)`

SetDescription sets Description field to given value.


### GetConfig

`func (o *ExportIntegrationSchema) GetConfig() map[string]interface{}`

GetConfig returns the Config field if non-nil, zero value otherwise.

### GetConfigOk

`func (o *ExportIntegrationSchema) GetConfigOk() (*map[string]interface{}, bool)`

GetConfigOk returns a tuple with the Config field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfig

`func (o *ExportIntegrationSchema) SetConfig(v map[string]interface{})`

SetConfig sets Config field to given value.


### GetTags

`func (o *ExportIntegrationSchema) GetTags() []string`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *ExportIntegrationSchema) GetTagsOk() (*[]string, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *ExportIntegrationSchema) SetTags(v []string)`

SetTags sets Tags field to given value.


### GetVersion

`func (o *ExportIntegrationSchema) GetVersion() string`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *ExportIntegrationSchema) GetVersionOk() (*string, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *ExportIntegrationSchema) SetVersion(v string)`

SetVersion sets Version field to given value.

### HasVersion

`func (o *ExportIntegrationSchema) HasVersion() bool`

HasVersion returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


