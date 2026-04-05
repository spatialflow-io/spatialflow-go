# ConfigFieldDefinitionRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | Field name in the configuration | 
**Label** | **string** | Display label for this field | 
**FieldType** | Pointer to **string** | Field type: text, password, url, email, tel, number, select, boolean, json, textarea | [optional] [default to "text"]
**Required** | Pointer to **bool** | Whether this field is required | [optional] [default to false]
**Placeholder** | Pointer to **NullableString** |  | [optional] 
**HelpText** | Pointer to **NullableString** |  | [optional] 
**DefaultValue** | Pointer to **NullableString** |  | [optional] 
**ValidationRules** | Pointer to **map[string]interface{}** | JSON validation rules | [optional] 
**Options** | Pointer to **[]map[string]string** | Options for select fields | [optional] 
**Order** | Pointer to **int32** | Display order for this field | [optional] [default to 0]
**DependsOn** | Pointer to **NullableString** |  | [optional] 
**DependsOnValue** | Pointer to **NullableString** |  | [optional] 

## Methods

### NewConfigFieldDefinitionRequest

`func NewConfigFieldDefinitionRequest(name string, label string, ) *ConfigFieldDefinitionRequest`

NewConfigFieldDefinitionRequest instantiates a new ConfigFieldDefinitionRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewConfigFieldDefinitionRequestWithDefaults

`func NewConfigFieldDefinitionRequestWithDefaults() *ConfigFieldDefinitionRequest`

NewConfigFieldDefinitionRequestWithDefaults instantiates a new ConfigFieldDefinitionRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *ConfigFieldDefinitionRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ConfigFieldDefinitionRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ConfigFieldDefinitionRequest) SetName(v string)`

SetName sets Name field to given value.


### GetLabel

`func (o *ConfigFieldDefinitionRequest) GetLabel() string`

GetLabel returns the Label field if non-nil, zero value otherwise.

### GetLabelOk

`func (o *ConfigFieldDefinitionRequest) GetLabelOk() (*string, bool)`

GetLabelOk returns a tuple with the Label field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabel

`func (o *ConfigFieldDefinitionRequest) SetLabel(v string)`

SetLabel sets Label field to given value.


### GetFieldType

`func (o *ConfigFieldDefinitionRequest) GetFieldType() string`

GetFieldType returns the FieldType field if non-nil, zero value otherwise.

### GetFieldTypeOk

`func (o *ConfigFieldDefinitionRequest) GetFieldTypeOk() (*string, bool)`

GetFieldTypeOk returns a tuple with the FieldType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFieldType

`func (o *ConfigFieldDefinitionRequest) SetFieldType(v string)`

SetFieldType sets FieldType field to given value.

### HasFieldType

`func (o *ConfigFieldDefinitionRequest) HasFieldType() bool`

HasFieldType returns a boolean if a field has been set.

### GetRequired

`func (o *ConfigFieldDefinitionRequest) GetRequired() bool`

GetRequired returns the Required field if non-nil, zero value otherwise.

### GetRequiredOk

`func (o *ConfigFieldDefinitionRequest) GetRequiredOk() (*bool, bool)`

GetRequiredOk returns a tuple with the Required field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequired

`func (o *ConfigFieldDefinitionRequest) SetRequired(v bool)`

SetRequired sets Required field to given value.

### HasRequired

`func (o *ConfigFieldDefinitionRequest) HasRequired() bool`

HasRequired returns a boolean if a field has been set.

### GetPlaceholder

`func (o *ConfigFieldDefinitionRequest) GetPlaceholder() string`

GetPlaceholder returns the Placeholder field if non-nil, zero value otherwise.

### GetPlaceholderOk

`func (o *ConfigFieldDefinitionRequest) GetPlaceholderOk() (*string, bool)`

GetPlaceholderOk returns a tuple with the Placeholder field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlaceholder

`func (o *ConfigFieldDefinitionRequest) SetPlaceholder(v string)`

SetPlaceholder sets Placeholder field to given value.

### HasPlaceholder

`func (o *ConfigFieldDefinitionRequest) HasPlaceholder() bool`

HasPlaceholder returns a boolean if a field has been set.

### SetPlaceholderNil

`func (o *ConfigFieldDefinitionRequest) SetPlaceholderNil(b bool)`

 SetPlaceholderNil sets the value for Placeholder to be an explicit nil

### UnsetPlaceholder
`func (o *ConfigFieldDefinitionRequest) UnsetPlaceholder()`

UnsetPlaceholder ensures that no value is present for Placeholder, not even an explicit nil
### GetHelpText

`func (o *ConfigFieldDefinitionRequest) GetHelpText() string`

GetHelpText returns the HelpText field if non-nil, zero value otherwise.

### GetHelpTextOk

`func (o *ConfigFieldDefinitionRequest) GetHelpTextOk() (*string, bool)`

GetHelpTextOk returns a tuple with the HelpText field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHelpText

`func (o *ConfigFieldDefinitionRequest) SetHelpText(v string)`

SetHelpText sets HelpText field to given value.

### HasHelpText

`func (o *ConfigFieldDefinitionRequest) HasHelpText() bool`

HasHelpText returns a boolean if a field has been set.

### SetHelpTextNil

`func (o *ConfigFieldDefinitionRequest) SetHelpTextNil(b bool)`

 SetHelpTextNil sets the value for HelpText to be an explicit nil

### UnsetHelpText
`func (o *ConfigFieldDefinitionRequest) UnsetHelpText()`

UnsetHelpText ensures that no value is present for HelpText, not even an explicit nil
### GetDefaultValue

`func (o *ConfigFieldDefinitionRequest) GetDefaultValue() string`

GetDefaultValue returns the DefaultValue field if non-nil, zero value otherwise.

### GetDefaultValueOk

`func (o *ConfigFieldDefinitionRequest) GetDefaultValueOk() (*string, bool)`

GetDefaultValueOk returns a tuple with the DefaultValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultValue

`func (o *ConfigFieldDefinitionRequest) SetDefaultValue(v string)`

SetDefaultValue sets DefaultValue field to given value.

### HasDefaultValue

`func (o *ConfigFieldDefinitionRequest) HasDefaultValue() bool`

HasDefaultValue returns a boolean if a field has been set.

### SetDefaultValueNil

`func (o *ConfigFieldDefinitionRequest) SetDefaultValueNil(b bool)`

 SetDefaultValueNil sets the value for DefaultValue to be an explicit nil

### UnsetDefaultValue
`func (o *ConfigFieldDefinitionRequest) UnsetDefaultValue()`

UnsetDefaultValue ensures that no value is present for DefaultValue, not even an explicit nil
### GetValidationRules

`func (o *ConfigFieldDefinitionRequest) GetValidationRules() map[string]interface{}`

GetValidationRules returns the ValidationRules field if non-nil, zero value otherwise.

### GetValidationRulesOk

`func (o *ConfigFieldDefinitionRequest) GetValidationRulesOk() (*map[string]interface{}, bool)`

GetValidationRulesOk returns a tuple with the ValidationRules field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValidationRules

`func (o *ConfigFieldDefinitionRequest) SetValidationRules(v map[string]interface{})`

SetValidationRules sets ValidationRules field to given value.

### HasValidationRules

`func (o *ConfigFieldDefinitionRequest) HasValidationRules() bool`

HasValidationRules returns a boolean if a field has been set.

### GetOptions

`func (o *ConfigFieldDefinitionRequest) GetOptions() []map[string]string`

GetOptions returns the Options field if non-nil, zero value otherwise.

### GetOptionsOk

`func (o *ConfigFieldDefinitionRequest) GetOptionsOk() (*[]map[string]string, bool)`

GetOptionsOk returns a tuple with the Options field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOptions

`func (o *ConfigFieldDefinitionRequest) SetOptions(v []map[string]string)`

SetOptions sets Options field to given value.

### HasOptions

`func (o *ConfigFieldDefinitionRequest) HasOptions() bool`

HasOptions returns a boolean if a field has been set.

### GetOrder

`func (o *ConfigFieldDefinitionRequest) GetOrder() int32`

GetOrder returns the Order field if non-nil, zero value otherwise.

### GetOrderOk

`func (o *ConfigFieldDefinitionRequest) GetOrderOk() (*int32, bool)`

GetOrderOk returns a tuple with the Order field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrder

`func (o *ConfigFieldDefinitionRequest) SetOrder(v int32)`

SetOrder sets Order field to given value.

### HasOrder

`func (o *ConfigFieldDefinitionRequest) HasOrder() bool`

HasOrder returns a boolean if a field has been set.

### GetDependsOn

`func (o *ConfigFieldDefinitionRequest) GetDependsOn() string`

GetDependsOn returns the DependsOn field if non-nil, zero value otherwise.

### GetDependsOnOk

`func (o *ConfigFieldDefinitionRequest) GetDependsOnOk() (*string, bool)`

GetDependsOnOk returns a tuple with the DependsOn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDependsOn

`func (o *ConfigFieldDefinitionRequest) SetDependsOn(v string)`

SetDependsOn sets DependsOn field to given value.

### HasDependsOn

`func (o *ConfigFieldDefinitionRequest) HasDependsOn() bool`

HasDependsOn returns a boolean if a field has been set.

### SetDependsOnNil

`func (o *ConfigFieldDefinitionRequest) SetDependsOnNil(b bool)`

 SetDependsOnNil sets the value for DependsOn to be an explicit nil

### UnsetDependsOn
`func (o *ConfigFieldDefinitionRequest) UnsetDependsOn()`

UnsetDependsOn ensures that no value is present for DependsOn, not even an explicit nil
### GetDependsOnValue

`func (o *ConfigFieldDefinitionRequest) GetDependsOnValue() string`

GetDependsOnValue returns the DependsOnValue field if non-nil, zero value otherwise.

### GetDependsOnValueOk

`func (o *ConfigFieldDefinitionRequest) GetDependsOnValueOk() (*string, bool)`

GetDependsOnValueOk returns a tuple with the DependsOnValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDependsOnValue

`func (o *ConfigFieldDefinitionRequest) SetDependsOnValue(v string)`

SetDependsOnValue sets DependsOnValue field to given value.

### HasDependsOnValue

`func (o *ConfigFieldDefinitionRequest) HasDependsOnValue() bool`

HasDependsOnValue returns a boolean if a field has been set.

### SetDependsOnValueNil

`func (o *ConfigFieldDefinitionRequest) SetDependsOnValueNil(b bool)`

 SetDependsOnValueNil sets the value for DependsOnValue to be an explicit nil

### UnsetDependsOnValue
`func (o *ConfigFieldDefinitionRequest) UnsetDependsOnValue()`

UnsetDependsOnValue ensures that no value is present for DependsOnValue, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


