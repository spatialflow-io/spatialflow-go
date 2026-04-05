# ConfigurationResetResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Success** | **bool** |  | 
**NewSource** | **string** |  | 
**EffectiveValue** | Pointer to [**NullableAnyOf**](anyOf&lt;&gt;.md) |  | [optional] 
**Error** | Pointer to **NullableString** |  | [optional] 

## Methods

### NewConfigurationResetResponse

`func NewConfigurationResetResponse(success bool, newSource string, ) *ConfigurationResetResponse`

NewConfigurationResetResponse instantiates a new ConfigurationResetResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewConfigurationResetResponseWithDefaults

`func NewConfigurationResetResponseWithDefaults() *ConfigurationResetResponse`

NewConfigurationResetResponseWithDefaults instantiates a new ConfigurationResetResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSuccess

`func (o *ConfigurationResetResponse) GetSuccess() bool`

GetSuccess returns the Success field if non-nil, zero value otherwise.

### GetSuccessOk

`func (o *ConfigurationResetResponse) GetSuccessOk() (*bool, bool)`

GetSuccessOk returns a tuple with the Success field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSuccess

`func (o *ConfigurationResetResponse) SetSuccess(v bool)`

SetSuccess sets Success field to given value.


### GetNewSource

`func (o *ConfigurationResetResponse) GetNewSource() string`

GetNewSource returns the NewSource field if non-nil, zero value otherwise.

### GetNewSourceOk

`func (o *ConfigurationResetResponse) GetNewSourceOk() (*string, bool)`

GetNewSourceOk returns a tuple with the NewSource field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNewSource

`func (o *ConfigurationResetResponse) SetNewSource(v string)`

SetNewSource sets NewSource field to given value.


### GetEffectiveValue

`func (o *ConfigurationResetResponse) GetEffectiveValue() AnyOf`

GetEffectiveValue returns the EffectiveValue field if non-nil, zero value otherwise.

### GetEffectiveValueOk

`func (o *ConfigurationResetResponse) GetEffectiveValueOk() (*AnyOf, bool)`

GetEffectiveValueOk returns a tuple with the EffectiveValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEffectiveValue

`func (o *ConfigurationResetResponse) SetEffectiveValue(v AnyOf)`

SetEffectiveValue sets EffectiveValue field to given value.

### HasEffectiveValue

`func (o *ConfigurationResetResponse) HasEffectiveValue() bool`

HasEffectiveValue returns a boolean if a field has been set.

### SetEffectiveValueNil

`func (o *ConfigurationResetResponse) SetEffectiveValueNil(b bool)`

 SetEffectiveValueNil sets the value for EffectiveValue to be an explicit nil

### UnsetEffectiveValue
`func (o *ConfigurationResetResponse) UnsetEffectiveValue()`

UnsetEffectiveValue ensures that no value is present for EffectiveValue, not even an explicit nil
### GetError

`func (o *ConfigurationResetResponse) GetError() string`

GetError returns the Error field if non-nil, zero value otherwise.

### GetErrorOk

`func (o *ConfigurationResetResponse) GetErrorOk() (*string, bool)`

GetErrorOk returns a tuple with the Error field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetError

`func (o *ConfigurationResetResponse) SetError(v string)`

SetError sets Error field to given value.

### HasError

`func (o *ConfigurationResetResponse) HasError() bool`

HasError returns a boolean if a field has been set.

### SetErrorNil

`func (o *ConfigurationResetResponse) SetErrorNil(b bool)`

 SetErrorNil sets the value for Error to be an explicit nil

### UnsetError
`func (o *ConfigurationResetResponse) UnsetError()`

UnsetError ensures that no value is present for Error, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


