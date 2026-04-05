# ConfigFieldDefinitionResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** |  | 
**IntegrationTypeId** | **string** |  | 
**Name** | **string** |  | 
**Label** | **string** |  | 
**FieldType** | **string** |  | 
**Required** | **bool** |  | 
**Placeholder** | **NullableString** |  | 
**HelpText** | **NullableString** |  | 
**DefaultValue** | **NullableString** |  | 
**ValidationRules** | **map[string]interface{}** |  | 
**Options** | **[]map[string]string** |  | 
**Order** | **int32** |  | 
**DependsOn** | **NullableString** |  | 
**DependsOnValue** | **NullableString** |  | 

## Methods

### NewConfigFieldDefinitionResponse

`func NewConfigFieldDefinitionResponse(id string, integrationTypeId string, name string, label string, fieldType string, required bool, placeholder NullableString, helpText NullableString, defaultValue NullableString, validationRules map[string]interface{}, options []map[string]string, order int32, dependsOn NullableString, dependsOnValue NullableString, ) *ConfigFieldDefinitionResponse`

NewConfigFieldDefinitionResponse instantiates a new ConfigFieldDefinitionResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewConfigFieldDefinitionResponseWithDefaults

`func NewConfigFieldDefinitionResponseWithDefaults() *ConfigFieldDefinitionResponse`

NewConfigFieldDefinitionResponseWithDefaults instantiates a new ConfigFieldDefinitionResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ConfigFieldDefinitionResponse) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ConfigFieldDefinitionResponse) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ConfigFieldDefinitionResponse) SetId(v string)`

SetId sets Id field to given value.


### GetIntegrationTypeId

`func (o *ConfigFieldDefinitionResponse) GetIntegrationTypeId() string`

GetIntegrationTypeId returns the IntegrationTypeId field if non-nil, zero value otherwise.

### GetIntegrationTypeIdOk

`func (o *ConfigFieldDefinitionResponse) GetIntegrationTypeIdOk() (*string, bool)`

GetIntegrationTypeIdOk returns a tuple with the IntegrationTypeId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIntegrationTypeId

`func (o *ConfigFieldDefinitionResponse) SetIntegrationTypeId(v string)`

SetIntegrationTypeId sets IntegrationTypeId field to given value.


### GetName

`func (o *ConfigFieldDefinitionResponse) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ConfigFieldDefinitionResponse) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ConfigFieldDefinitionResponse) SetName(v string)`

SetName sets Name field to given value.


### GetLabel

`func (o *ConfigFieldDefinitionResponse) GetLabel() string`

GetLabel returns the Label field if non-nil, zero value otherwise.

### GetLabelOk

`func (o *ConfigFieldDefinitionResponse) GetLabelOk() (*string, bool)`

GetLabelOk returns a tuple with the Label field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabel

`func (o *ConfigFieldDefinitionResponse) SetLabel(v string)`

SetLabel sets Label field to given value.


### GetFieldType

`func (o *ConfigFieldDefinitionResponse) GetFieldType() string`

GetFieldType returns the FieldType field if non-nil, zero value otherwise.

### GetFieldTypeOk

`func (o *ConfigFieldDefinitionResponse) GetFieldTypeOk() (*string, bool)`

GetFieldTypeOk returns a tuple with the FieldType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFieldType

`func (o *ConfigFieldDefinitionResponse) SetFieldType(v string)`

SetFieldType sets FieldType field to given value.


### GetRequired

`func (o *ConfigFieldDefinitionResponse) GetRequired() bool`

GetRequired returns the Required field if non-nil, zero value otherwise.

### GetRequiredOk

`func (o *ConfigFieldDefinitionResponse) GetRequiredOk() (*bool, bool)`

GetRequiredOk returns a tuple with the Required field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequired

`func (o *ConfigFieldDefinitionResponse) SetRequired(v bool)`

SetRequired sets Required field to given value.


### GetPlaceholder

`func (o *ConfigFieldDefinitionResponse) GetPlaceholder() string`

GetPlaceholder returns the Placeholder field if non-nil, zero value otherwise.

### GetPlaceholderOk

`func (o *ConfigFieldDefinitionResponse) GetPlaceholderOk() (*string, bool)`

GetPlaceholderOk returns a tuple with the Placeholder field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlaceholder

`func (o *ConfigFieldDefinitionResponse) SetPlaceholder(v string)`

SetPlaceholder sets Placeholder field to given value.


### SetPlaceholderNil

`func (o *ConfigFieldDefinitionResponse) SetPlaceholderNil(b bool)`

 SetPlaceholderNil sets the value for Placeholder to be an explicit nil

### UnsetPlaceholder
`func (o *ConfigFieldDefinitionResponse) UnsetPlaceholder()`

UnsetPlaceholder ensures that no value is present for Placeholder, not even an explicit nil
### GetHelpText

`func (o *ConfigFieldDefinitionResponse) GetHelpText() string`

GetHelpText returns the HelpText field if non-nil, zero value otherwise.

### GetHelpTextOk

`func (o *ConfigFieldDefinitionResponse) GetHelpTextOk() (*string, bool)`

GetHelpTextOk returns a tuple with the HelpText field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHelpText

`func (o *ConfigFieldDefinitionResponse) SetHelpText(v string)`

SetHelpText sets HelpText field to given value.


### SetHelpTextNil

`func (o *ConfigFieldDefinitionResponse) SetHelpTextNil(b bool)`

 SetHelpTextNil sets the value for HelpText to be an explicit nil

### UnsetHelpText
`func (o *ConfigFieldDefinitionResponse) UnsetHelpText()`

UnsetHelpText ensures that no value is present for HelpText, not even an explicit nil
### GetDefaultValue

`func (o *ConfigFieldDefinitionResponse) GetDefaultValue() string`

GetDefaultValue returns the DefaultValue field if non-nil, zero value otherwise.

### GetDefaultValueOk

`func (o *ConfigFieldDefinitionResponse) GetDefaultValueOk() (*string, bool)`

GetDefaultValueOk returns a tuple with the DefaultValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultValue

`func (o *ConfigFieldDefinitionResponse) SetDefaultValue(v string)`

SetDefaultValue sets DefaultValue field to given value.


### SetDefaultValueNil

`func (o *ConfigFieldDefinitionResponse) SetDefaultValueNil(b bool)`

 SetDefaultValueNil sets the value for DefaultValue to be an explicit nil

### UnsetDefaultValue
`func (o *ConfigFieldDefinitionResponse) UnsetDefaultValue()`

UnsetDefaultValue ensures that no value is present for DefaultValue, not even an explicit nil
### GetValidationRules

`func (o *ConfigFieldDefinitionResponse) GetValidationRules() map[string]interface{}`

GetValidationRules returns the ValidationRules field if non-nil, zero value otherwise.

### GetValidationRulesOk

`func (o *ConfigFieldDefinitionResponse) GetValidationRulesOk() (*map[string]interface{}, bool)`

GetValidationRulesOk returns a tuple with the ValidationRules field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValidationRules

`func (o *ConfigFieldDefinitionResponse) SetValidationRules(v map[string]interface{})`

SetValidationRules sets ValidationRules field to given value.


### GetOptions

`func (o *ConfigFieldDefinitionResponse) GetOptions() []map[string]string`

GetOptions returns the Options field if non-nil, zero value otherwise.

### GetOptionsOk

`func (o *ConfigFieldDefinitionResponse) GetOptionsOk() (*[]map[string]string, bool)`

GetOptionsOk returns a tuple with the Options field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOptions

`func (o *ConfigFieldDefinitionResponse) SetOptions(v []map[string]string)`

SetOptions sets Options field to given value.


### GetOrder

`func (o *ConfigFieldDefinitionResponse) GetOrder() int32`

GetOrder returns the Order field if non-nil, zero value otherwise.

### GetOrderOk

`func (o *ConfigFieldDefinitionResponse) GetOrderOk() (*int32, bool)`

GetOrderOk returns a tuple with the Order field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrder

`func (o *ConfigFieldDefinitionResponse) SetOrder(v int32)`

SetOrder sets Order field to given value.


### GetDependsOn

`func (o *ConfigFieldDefinitionResponse) GetDependsOn() string`

GetDependsOn returns the DependsOn field if non-nil, zero value otherwise.

### GetDependsOnOk

`func (o *ConfigFieldDefinitionResponse) GetDependsOnOk() (*string, bool)`

GetDependsOnOk returns a tuple with the DependsOn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDependsOn

`func (o *ConfigFieldDefinitionResponse) SetDependsOn(v string)`

SetDependsOn sets DependsOn field to given value.


### SetDependsOnNil

`func (o *ConfigFieldDefinitionResponse) SetDependsOnNil(b bool)`

 SetDependsOnNil sets the value for DependsOn to be an explicit nil

### UnsetDependsOn
`func (o *ConfigFieldDefinitionResponse) UnsetDependsOn()`

UnsetDependsOn ensures that no value is present for DependsOn, not even an explicit nil
### GetDependsOnValue

`func (o *ConfigFieldDefinitionResponse) GetDependsOnValue() string`

GetDependsOnValue returns the DependsOnValue field if non-nil, zero value otherwise.

### GetDependsOnValueOk

`func (o *ConfigFieldDefinitionResponse) GetDependsOnValueOk() (*string, bool)`

GetDependsOnValueOk returns a tuple with the DependsOnValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDependsOnValue

`func (o *ConfigFieldDefinitionResponse) SetDependsOnValue(v string)`

SetDependsOnValue sets DependsOnValue field to given value.


### SetDependsOnValueNil

`func (o *ConfigFieldDefinitionResponse) SetDependsOnValueNil(b bool)`

 SetDependsOnValueNil sets the value for DependsOnValue to be an explicit nil

### UnsetDependsOnValue
`func (o *ConfigFieldDefinitionResponse) UnsetDependsOnValue()`

UnsetDependsOnValue ensures that no value is present for DependsOnValue, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


