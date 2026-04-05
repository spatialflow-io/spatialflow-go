# CircuitBreakerSchema

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Enabled** | Pointer to **bool** |  | [optional] [default to true]
**FailureThreshold** | Pointer to **int32** |  | [optional] [default to 5]
**RecoveryTimeoutMs** | Pointer to **int32** |  | [optional] [default to 60000]
**SuccessThreshold** | Pointer to **int32** |  | [optional] [default to 2]

## Methods

### NewCircuitBreakerSchema

`func NewCircuitBreakerSchema() *CircuitBreakerSchema`

NewCircuitBreakerSchema instantiates a new CircuitBreakerSchema object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCircuitBreakerSchemaWithDefaults

`func NewCircuitBreakerSchemaWithDefaults() *CircuitBreakerSchema`

NewCircuitBreakerSchemaWithDefaults instantiates a new CircuitBreakerSchema object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEnabled

`func (o *CircuitBreakerSchema) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *CircuitBreakerSchema) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *CircuitBreakerSchema) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *CircuitBreakerSchema) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.

### GetFailureThreshold

`func (o *CircuitBreakerSchema) GetFailureThreshold() int32`

GetFailureThreshold returns the FailureThreshold field if non-nil, zero value otherwise.

### GetFailureThresholdOk

`func (o *CircuitBreakerSchema) GetFailureThresholdOk() (*int32, bool)`

GetFailureThresholdOk returns a tuple with the FailureThreshold field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFailureThreshold

`func (o *CircuitBreakerSchema) SetFailureThreshold(v int32)`

SetFailureThreshold sets FailureThreshold field to given value.

### HasFailureThreshold

`func (o *CircuitBreakerSchema) HasFailureThreshold() bool`

HasFailureThreshold returns a boolean if a field has been set.

### GetRecoveryTimeoutMs

`func (o *CircuitBreakerSchema) GetRecoveryTimeoutMs() int32`

GetRecoveryTimeoutMs returns the RecoveryTimeoutMs field if non-nil, zero value otherwise.

### GetRecoveryTimeoutMsOk

`func (o *CircuitBreakerSchema) GetRecoveryTimeoutMsOk() (*int32, bool)`

GetRecoveryTimeoutMsOk returns a tuple with the RecoveryTimeoutMs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecoveryTimeoutMs

`func (o *CircuitBreakerSchema) SetRecoveryTimeoutMs(v int32)`

SetRecoveryTimeoutMs sets RecoveryTimeoutMs field to given value.

### HasRecoveryTimeoutMs

`func (o *CircuitBreakerSchema) HasRecoveryTimeoutMs() bool`

HasRecoveryTimeoutMs returns a boolean if a field has been set.

### GetSuccessThreshold

`func (o *CircuitBreakerSchema) GetSuccessThreshold() int32`

GetSuccessThreshold returns the SuccessThreshold field if non-nil, zero value otherwise.

### GetSuccessThresholdOk

`func (o *CircuitBreakerSchema) GetSuccessThresholdOk() (*int32, bool)`

GetSuccessThresholdOk returns a tuple with the SuccessThreshold field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSuccessThreshold

`func (o *CircuitBreakerSchema) SetSuccessThreshold(v int32)`

SetSuccessThreshold sets SuccessThreshold field to given value.

### HasSuccessThreshold

`func (o *CircuitBreakerSchema) HasSuccessThreshold() bool`

HasSuccessThreshold returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


