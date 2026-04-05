# ConfigurationItem

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Key** | **string** |  | 
**Name** | **string** |  | 
**Description** | **string** |  | 
**Category** | **string** |  | 
**ValueType** | **string** |  | 
**Value** | Pointer to [**NullableAnyOf**](anyOf&lt;&gt;.md) |  | [optional] 
**Source** | **string** |  | 
**RequiresRestart** | **bool** |  | 
**IsSensitive** | **bool** |  | 
**IsReadonly** | **bool** |  | 
**IsWriteOnly** | **bool** |  | 
**AllowEmpty** | Pointer to **bool** |  | [optional] [default to false]
**ValidationRules** | **map[string]interface{}** |  | 
**UpdatedAt** | Pointer to **NullableString** |  | [optional] 
**Error** | Pointer to **NullableString** |  | [optional] 

## Methods

### NewConfigurationItem

`func NewConfigurationItem(key string, name string, description string, category string, valueType string, source string, requiresRestart bool, isSensitive bool, isReadonly bool, isWriteOnly bool, validationRules map[string]interface{}, ) *ConfigurationItem`

NewConfigurationItem instantiates a new ConfigurationItem object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewConfigurationItemWithDefaults

`func NewConfigurationItemWithDefaults() *ConfigurationItem`

NewConfigurationItemWithDefaults instantiates a new ConfigurationItem object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetKey

`func (o *ConfigurationItem) GetKey() string`

GetKey returns the Key field if non-nil, zero value otherwise.

### GetKeyOk

`func (o *ConfigurationItem) GetKeyOk() (*string, bool)`

GetKeyOk returns a tuple with the Key field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKey

`func (o *ConfigurationItem) SetKey(v string)`

SetKey sets Key field to given value.


### GetName

`func (o *ConfigurationItem) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ConfigurationItem) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ConfigurationItem) SetName(v string)`

SetName sets Name field to given value.


### GetDescription

`func (o *ConfigurationItem) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *ConfigurationItem) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *ConfigurationItem) SetDescription(v string)`

SetDescription sets Description field to given value.


### GetCategory

`func (o *ConfigurationItem) GetCategory() string`

GetCategory returns the Category field if non-nil, zero value otherwise.

### GetCategoryOk

`func (o *ConfigurationItem) GetCategoryOk() (*string, bool)`

GetCategoryOk returns a tuple with the Category field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCategory

`func (o *ConfigurationItem) SetCategory(v string)`

SetCategory sets Category field to given value.


### GetValueType

`func (o *ConfigurationItem) GetValueType() string`

GetValueType returns the ValueType field if non-nil, zero value otherwise.

### GetValueTypeOk

`func (o *ConfigurationItem) GetValueTypeOk() (*string, bool)`

GetValueTypeOk returns a tuple with the ValueType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValueType

`func (o *ConfigurationItem) SetValueType(v string)`

SetValueType sets ValueType field to given value.


### GetValue

`func (o *ConfigurationItem) GetValue() AnyOf`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *ConfigurationItem) GetValueOk() (*AnyOf, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *ConfigurationItem) SetValue(v AnyOf)`

SetValue sets Value field to given value.

### HasValue

`func (o *ConfigurationItem) HasValue() bool`

HasValue returns a boolean if a field has been set.

### SetValueNil

`func (o *ConfigurationItem) SetValueNil(b bool)`

 SetValueNil sets the value for Value to be an explicit nil

### UnsetValue
`func (o *ConfigurationItem) UnsetValue()`

UnsetValue ensures that no value is present for Value, not even an explicit nil
### GetSource

`func (o *ConfigurationItem) GetSource() string`

GetSource returns the Source field if non-nil, zero value otherwise.

### GetSourceOk

`func (o *ConfigurationItem) GetSourceOk() (*string, bool)`

GetSourceOk returns a tuple with the Source field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSource

`func (o *ConfigurationItem) SetSource(v string)`

SetSource sets Source field to given value.


### GetRequiresRestart

`func (o *ConfigurationItem) GetRequiresRestart() bool`

GetRequiresRestart returns the RequiresRestart field if non-nil, zero value otherwise.

### GetRequiresRestartOk

`func (o *ConfigurationItem) GetRequiresRestartOk() (*bool, bool)`

GetRequiresRestartOk returns a tuple with the RequiresRestart field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequiresRestart

`func (o *ConfigurationItem) SetRequiresRestart(v bool)`

SetRequiresRestart sets RequiresRestart field to given value.


### GetIsSensitive

`func (o *ConfigurationItem) GetIsSensitive() bool`

GetIsSensitive returns the IsSensitive field if non-nil, zero value otherwise.

### GetIsSensitiveOk

`func (o *ConfigurationItem) GetIsSensitiveOk() (*bool, bool)`

GetIsSensitiveOk returns a tuple with the IsSensitive field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsSensitive

`func (o *ConfigurationItem) SetIsSensitive(v bool)`

SetIsSensitive sets IsSensitive field to given value.


### GetIsReadonly

`func (o *ConfigurationItem) GetIsReadonly() bool`

GetIsReadonly returns the IsReadonly field if non-nil, zero value otherwise.

### GetIsReadonlyOk

`func (o *ConfigurationItem) GetIsReadonlyOk() (*bool, bool)`

GetIsReadonlyOk returns a tuple with the IsReadonly field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsReadonly

`func (o *ConfigurationItem) SetIsReadonly(v bool)`

SetIsReadonly sets IsReadonly field to given value.


### GetIsWriteOnly

`func (o *ConfigurationItem) GetIsWriteOnly() bool`

GetIsWriteOnly returns the IsWriteOnly field if non-nil, zero value otherwise.

### GetIsWriteOnlyOk

`func (o *ConfigurationItem) GetIsWriteOnlyOk() (*bool, bool)`

GetIsWriteOnlyOk returns a tuple with the IsWriteOnly field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsWriteOnly

`func (o *ConfigurationItem) SetIsWriteOnly(v bool)`

SetIsWriteOnly sets IsWriteOnly field to given value.


### GetAllowEmpty

`func (o *ConfigurationItem) GetAllowEmpty() bool`

GetAllowEmpty returns the AllowEmpty field if non-nil, zero value otherwise.

### GetAllowEmptyOk

`func (o *ConfigurationItem) GetAllowEmptyOk() (*bool, bool)`

GetAllowEmptyOk returns a tuple with the AllowEmpty field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowEmpty

`func (o *ConfigurationItem) SetAllowEmpty(v bool)`

SetAllowEmpty sets AllowEmpty field to given value.

### HasAllowEmpty

`func (o *ConfigurationItem) HasAllowEmpty() bool`

HasAllowEmpty returns a boolean if a field has been set.

### GetValidationRules

`func (o *ConfigurationItem) GetValidationRules() map[string]interface{}`

GetValidationRules returns the ValidationRules field if non-nil, zero value otherwise.

### GetValidationRulesOk

`func (o *ConfigurationItem) GetValidationRulesOk() (*map[string]interface{}, bool)`

GetValidationRulesOk returns a tuple with the ValidationRules field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValidationRules

`func (o *ConfigurationItem) SetValidationRules(v map[string]interface{})`

SetValidationRules sets ValidationRules field to given value.


### GetUpdatedAt

`func (o *ConfigurationItem) GetUpdatedAt() string`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *ConfigurationItem) GetUpdatedAtOk() (*string, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *ConfigurationItem) SetUpdatedAt(v string)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *ConfigurationItem) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.

### SetUpdatedAtNil

`func (o *ConfigurationItem) SetUpdatedAtNil(b bool)`

 SetUpdatedAtNil sets the value for UpdatedAt to be an explicit nil

### UnsetUpdatedAt
`func (o *ConfigurationItem) UnsetUpdatedAt()`

UnsetUpdatedAt ensures that no value is present for UpdatedAt, not even an explicit nil
### GetError

`func (o *ConfigurationItem) GetError() string`

GetError returns the Error field if non-nil, zero value otherwise.

### GetErrorOk

`func (o *ConfigurationItem) GetErrorOk() (*string, bool)`

GetErrorOk returns a tuple with the Error field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetError

`func (o *ConfigurationItem) SetError(v string)`

SetError sets Error field to given value.

### HasError

`func (o *ConfigurationItem) HasError() bool`

HasError returns a boolean if a field has been set.

### SetErrorNil

`func (o *ConfigurationItem) SetErrorNil(b bool)`

 SetErrorNil sets the value for Error to be an explicit nil

### UnsetError
`func (o *ConfigurationItem) UnsetError()`

UnsetError ensures that no value is present for Error, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


