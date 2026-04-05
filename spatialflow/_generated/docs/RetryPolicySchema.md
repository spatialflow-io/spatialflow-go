# RetryPolicySchema

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Enabled** | Pointer to **bool** |  | [optional] [default to true]
**MaxAttempts** | Pointer to **int32** |  | [optional] [default to 3]
**Strategy** | Pointer to [**RetryStrategyEnum**](RetryStrategyEnum.md) |  | [optional] [default to EXPONENTIAL_BACKOFF]
**BaseDelayMs** | Pointer to **int32** |  | [optional] [default to 1000]
**MaxDelayMs** | Pointer to **int32** |  | [optional] [default to 30000]
**Jitter** | Pointer to **bool** |  | [optional] [default to true]

## Methods

### NewRetryPolicySchema

`func NewRetryPolicySchema() *RetryPolicySchema`

NewRetryPolicySchema instantiates a new RetryPolicySchema object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRetryPolicySchemaWithDefaults

`func NewRetryPolicySchemaWithDefaults() *RetryPolicySchema`

NewRetryPolicySchemaWithDefaults instantiates a new RetryPolicySchema object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEnabled

`func (o *RetryPolicySchema) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *RetryPolicySchema) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *RetryPolicySchema) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *RetryPolicySchema) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.

### GetMaxAttempts

`func (o *RetryPolicySchema) GetMaxAttempts() int32`

GetMaxAttempts returns the MaxAttempts field if non-nil, zero value otherwise.

### GetMaxAttemptsOk

`func (o *RetryPolicySchema) GetMaxAttemptsOk() (*int32, bool)`

GetMaxAttemptsOk returns a tuple with the MaxAttempts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxAttempts

`func (o *RetryPolicySchema) SetMaxAttempts(v int32)`

SetMaxAttempts sets MaxAttempts field to given value.

### HasMaxAttempts

`func (o *RetryPolicySchema) HasMaxAttempts() bool`

HasMaxAttempts returns a boolean if a field has been set.

### GetStrategy

`func (o *RetryPolicySchema) GetStrategy() RetryStrategyEnum`

GetStrategy returns the Strategy field if non-nil, zero value otherwise.

### GetStrategyOk

`func (o *RetryPolicySchema) GetStrategyOk() (*RetryStrategyEnum, bool)`

GetStrategyOk returns a tuple with the Strategy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStrategy

`func (o *RetryPolicySchema) SetStrategy(v RetryStrategyEnum)`

SetStrategy sets Strategy field to given value.

### HasStrategy

`func (o *RetryPolicySchema) HasStrategy() bool`

HasStrategy returns a boolean if a field has been set.

### GetBaseDelayMs

`func (o *RetryPolicySchema) GetBaseDelayMs() int32`

GetBaseDelayMs returns the BaseDelayMs field if non-nil, zero value otherwise.

### GetBaseDelayMsOk

`func (o *RetryPolicySchema) GetBaseDelayMsOk() (*int32, bool)`

GetBaseDelayMsOk returns a tuple with the BaseDelayMs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBaseDelayMs

`func (o *RetryPolicySchema) SetBaseDelayMs(v int32)`

SetBaseDelayMs sets BaseDelayMs field to given value.

### HasBaseDelayMs

`func (o *RetryPolicySchema) HasBaseDelayMs() bool`

HasBaseDelayMs returns a boolean if a field has been set.

### GetMaxDelayMs

`func (o *RetryPolicySchema) GetMaxDelayMs() int32`

GetMaxDelayMs returns the MaxDelayMs field if non-nil, zero value otherwise.

### GetMaxDelayMsOk

`func (o *RetryPolicySchema) GetMaxDelayMsOk() (*int32, bool)`

GetMaxDelayMsOk returns a tuple with the MaxDelayMs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxDelayMs

`func (o *RetryPolicySchema) SetMaxDelayMs(v int32)`

SetMaxDelayMs sets MaxDelayMs field to given value.

### HasMaxDelayMs

`func (o *RetryPolicySchema) HasMaxDelayMs() bool`

HasMaxDelayMs returns a boolean if a field has been set.

### GetJitter

`func (o *RetryPolicySchema) GetJitter() bool`

GetJitter returns the Jitter field if non-nil, zero value otherwise.

### GetJitterOk

`func (o *RetryPolicySchema) GetJitterOk() (*bool, bool)`

GetJitterOk returns a tuple with the Jitter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJitter

`func (o *RetryPolicySchema) SetJitter(v bool)`

SetJitter sets Jitter field to given value.

### HasJitter

`func (o *RetryPolicySchema) HasJitter() bool`

HasJitter returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


