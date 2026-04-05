# RetryPolicyResponseSchema

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AvailablePolicies** | [**map[string]RetryPolicySchema**](RetryPolicySchema.md) |  | 
**RetryStrategies** | **[]string** |  | 
**CircuitBreakerDefaults** | [**CircuitBreakerSchema**](CircuitBreakerSchema.md) |  | 

## Methods

### NewRetryPolicyResponseSchema

`func NewRetryPolicyResponseSchema(availablePolicies map[string]RetryPolicySchema, retryStrategies []string, circuitBreakerDefaults CircuitBreakerSchema, ) *RetryPolicyResponseSchema`

NewRetryPolicyResponseSchema instantiates a new RetryPolicyResponseSchema object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRetryPolicyResponseSchemaWithDefaults

`func NewRetryPolicyResponseSchemaWithDefaults() *RetryPolicyResponseSchema`

NewRetryPolicyResponseSchemaWithDefaults instantiates a new RetryPolicyResponseSchema object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAvailablePolicies

`func (o *RetryPolicyResponseSchema) GetAvailablePolicies() map[string]RetryPolicySchema`

GetAvailablePolicies returns the AvailablePolicies field if non-nil, zero value otherwise.

### GetAvailablePoliciesOk

`func (o *RetryPolicyResponseSchema) GetAvailablePoliciesOk() (*map[string]RetryPolicySchema, bool)`

GetAvailablePoliciesOk returns a tuple with the AvailablePolicies field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvailablePolicies

`func (o *RetryPolicyResponseSchema) SetAvailablePolicies(v map[string]RetryPolicySchema)`

SetAvailablePolicies sets AvailablePolicies field to given value.


### GetRetryStrategies

`func (o *RetryPolicyResponseSchema) GetRetryStrategies() []string`

GetRetryStrategies returns the RetryStrategies field if non-nil, zero value otherwise.

### GetRetryStrategiesOk

`func (o *RetryPolicyResponseSchema) GetRetryStrategiesOk() (*[]string, bool)`

GetRetryStrategiesOk returns a tuple with the RetryStrategies field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRetryStrategies

`func (o *RetryPolicyResponseSchema) SetRetryStrategies(v []string)`

SetRetryStrategies sets RetryStrategies field to given value.


### GetCircuitBreakerDefaults

`func (o *RetryPolicyResponseSchema) GetCircuitBreakerDefaults() CircuitBreakerSchema`

GetCircuitBreakerDefaults returns the CircuitBreakerDefaults field if non-nil, zero value otherwise.

### GetCircuitBreakerDefaultsOk

`func (o *RetryPolicyResponseSchema) GetCircuitBreakerDefaultsOk() (*CircuitBreakerSchema, bool)`

GetCircuitBreakerDefaultsOk returns a tuple with the CircuitBreakerDefaults field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCircuitBreakerDefaults

`func (o *RetryPolicyResponseSchema) SetCircuitBreakerDefaults(v CircuitBreakerSchema)`

SetCircuitBreakerDefaults sets CircuitBreakerDefaults field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


