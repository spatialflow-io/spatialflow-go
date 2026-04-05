# ConfigurationUpdateResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Success** | **bool** |  | 
**RequiresRestart** | **bool** |  | 
**Error** | Pointer to **NullableString** |  | [optional] 
**EffectiveValue** | Pointer to [**NullableAnyOf**](anyOf&lt;&gt;.md) |  | [optional] 

## Methods

### NewConfigurationUpdateResponse

`func NewConfigurationUpdateResponse(success bool, requiresRestart bool, ) *ConfigurationUpdateResponse`

NewConfigurationUpdateResponse instantiates a new ConfigurationUpdateResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewConfigurationUpdateResponseWithDefaults

`func NewConfigurationUpdateResponseWithDefaults() *ConfigurationUpdateResponse`

NewConfigurationUpdateResponseWithDefaults instantiates a new ConfigurationUpdateResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSuccess

`func (o *ConfigurationUpdateResponse) GetSuccess() bool`

GetSuccess returns the Success field if non-nil, zero value otherwise.

### GetSuccessOk

`func (o *ConfigurationUpdateResponse) GetSuccessOk() (*bool, bool)`

GetSuccessOk returns a tuple with the Success field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSuccess

`func (o *ConfigurationUpdateResponse) SetSuccess(v bool)`

SetSuccess sets Success field to given value.


### GetRequiresRestart

`func (o *ConfigurationUpdateResponse) GetRequiresRestart() bool`

GetRequiresRestart returns the RequiresRestart field if non-nil, zero value otherwise.

### GetRequiresRestartOk

`func (o *ConfigurationUpdateResponse) GetRequiresRestartOk() (*bool, bool)`

GetRequiresRestartOk returns a tuple with the RequiresRestart field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequiresRestart

`func (o *ConfigurationUpdateResponse) SetRequiresRestart(v bool)`

SetRequiresRestart sets RequiresRestart field to given value.


### GetError

`func (o *ConfigurationUpdateResponse) GetError() string`

GetError returns the Error field if non-nil, zero value otherwise.

### GetErrorOk

`func (o *ConfigurationUpdateResponse) GetErrorOk() (*string, bool)`

GetErrorOk returns a tuple with the Error field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetError

`func (o *ConfigurationUpdateResponse) SetError(v string)`

SetError sets Error field to given value.

### HasError

`func (o *ConfigurationUpdateResponse) HasError() bool`

HasError returns a boolean if a field has been set.

### SetErrorNil

`func (o *ConfigurationUpdateResponse) SetErrorNil(b bool)`

 SetErrorNil sets the value for Error to be an explicit nil

### UnsetError
`func (o *ConfigurationUpdateResponse) UnsetError()`

UnsetError ensures that no value is present for Error, not even an explicit nil
### GetEffectiveValue

`func (o *ConfigurationUpdateResponse) GetEffectiveValue() AnyOf`

GetEffectiveValue returns the EffectiveValue field if non-nil, zero value otherwise.

### GetEffectiveValueOk

`func (o *ConfigurationUpdateResponse) GetEffectiveValueOk() (*AnyOf, bool)`

GetEffectiveValueOk returns a tuple with the EffectiveValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEffectiveValue

`func (o *ConfigurationUpdateResponse) SetEffectiveValue(v AnyOf)`

SetEffectiveValue sets EffectiveValue field to given value.

### HasEffectiveValue

`func (o *ConfigurationUpdateResponse) HasEffectiveValue() bool`

HasEffectiveValue returns a boolean if a field has been set.

### SetEffectiveValueNil

`func (o *ConfigurationUpdateResponse) SetEffectiveValueNil(b bool)`

 SetEffectiveValueNil sets the value for EffectiveValue to be an explicit nil

### UnsetEffectiveValue
`func (o *ConfigurationUpdateResponse) UnsetEffectiveValue()`

UnsetEffectiveValue ensures that no value is present for EffectiveValue, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


