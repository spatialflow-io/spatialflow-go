# CreateIntegrationSchema

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | User-friendly name for the integration | 
**Type** | **string** | Integration type (webhook, slack, twilio_sms, etc.) | 
**Description** | Pointer to **NullableString** |  | [optional] 
**Config** | **map[string]interface{}** | Integration-specific configuration | 
**Tags** | Pointer to **[]string** |  | [optional] 

## Methods

### NewCreateIntegrationSchema

`func NewCreateIntegrationSchema(name string, type_ string, config map[string]interface{}, ) *CreateIntegrationSchema`

NewCreateIntegrationSchema instantiates a new CreateIntegrationSchema object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateIntegrationSchemaWithDefaults

`func NewCreateIntegrationSchemaWithDefaults() *CreateIntegrationSchema`

NewCreateIntegrationSchemaWithDefaults instantiates a new CreateIntegrationSchema object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *CreateIntegrationSchema) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CreateIntegrationSchema) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CreateIntegrationSchema) SetName(v string)`

SetName sets Name field to given value.


### GetType

`func (o *CreateIntegrationSchema) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *CreateIntegrationSchema) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *CreateIntegrationSchema) SetType(v string)`

SetType sets Type field to given value.


### GetDescription

`func (o *CreateIntegrationSchema) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *CreateIntegrationSchema) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *CreateIntegrationSchema) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *CreateIntegrationSchema) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescriptionNil

`func (o *CreateIntegrationSchema) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *CreateIntegrationSchema) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetConfig

`func (o *CreateIntegrationSchema) GetConfig() map[string]interface{}`

GetConfig returns the Config field if non-nil, zero value otherwise.

### GetConfigOk

`func (o *CreateIntegrationSchema) GetConfigOk() (*map[string]interface{}, bool)`

GetConfigOk returns a tuple with the Config field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfig

`func (o *CreateIntegrationSchema) SetConfig(v map[string]interface{})`

SetConfig sets Config field to given value.


### GetTags

`func (o *CreateIntegrationSchema) GetTags() []string`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *CreateIntegrationSchema) GetTagsOk() (*[]string, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *CreateIntegrationSchema) SetTags(v []string)`

SetTags sets Tags field to given value.

### HasTags

`func (o *CreateIntegrationSchema) HasTags() bool`

HasTags returns a boolean if a field has been set.

### SetTagsNil

`func (o *CreateIntegrationSchema) SetTagsNil(b bool)`

 SetTagsNil sets the value for Tags to be an explicit nil

### UnsetTags
`func (o *CreateIntegrationSchema) UnsetTags()`

UnsetTags ensures that no value is present for Tags, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


