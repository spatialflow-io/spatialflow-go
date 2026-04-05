# ActionRetryConfigSchema

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**RetryPolicy** | Pointer to [**NullableRetryPolicySchema**](RetryPolicySchema.md) |  | [optional] 
**CircuitBreaker** | Pointer to [**NullableCircuitBreakerSchema**](CircuitBreakerSchema.md) |  | [optional] 
**RetryPolicyName** | Pointer to **NullableString** |  | [optional] 

## Methods

### NewActionRetryConfigSchema

`func NewActionRetryConfigSchema() *ActionRetryConfigSchema`

NewActionRetryConfigSchema instantiates a new ActionRetryConfigSchema object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewActionRetryConfigSchemaWithDefaults

`func NewActionRetryConfigSchemaWithDefaults() *ActionRetryConfigSchema`

NewActionRetryConfigSchemaWithDefaults instantiates a new ActionRetryConfigSchema object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRetryPolicy

`func (o *ActionRetryConfigSchema) GetRetryPolicy() RetryPolicySchema`

GetRetryPolicy returns the RetryPolicy field if non-nil, zero value otherwise.

### GetRetryPolicyOk

`func (o *ActionRetryConfigSchema) GetRetryPolicyOk() (*RetryPolicySchema, bool)`

GetRetryPolicyOk returns a tuple with the RetryPolicy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRetryPolicy

`func (o *ActionRetryConfigSchema) SetRetryPolicy(v RetryPolicySchema)`

SetRetryPolicy sets RetryPolicy field to given value.

### HasRetryPolicy

`func (o *ActionRetryConfigSchema) HasRetryPolicy() bool`

HasRetryPolicy returns a boolean if a field has been set.

### SetRetryPolicyNil

`func (o *ActionRetryConfigSchema) SetRetryPolicyNil(b bool)`

 SetRetryPolicyNil sets the value for RetryPolicy to be an explicit nil

### UnsetRetryPolicy
`func (o *ActionRetryConfigSchema) UnsetRetryPolicy()`

UnsetRetryPolicy ensures that no value is present for RetryPolicy, not even an explicit nil
### GetCircuitBreaker

`func (o *ActionRetryConfigSchema) GetCircuitBreaker() CircuitBreakerSchema`

GetCircuitBreaker returns the CircuitBreaker field if non-nil, zero value otherwise.

### GetCircuitBreakerOk

`func (o *ActionRetryConfigSchema) GetCircuitBreakerOk() (*CircuitBreakerSchema, bool)`

GetCircuitBreakerOk returns a tuple with the CircuitBreaker field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCircuitBreaker

`func (o *ActionRetryConfigSchema) SetCircuitBreaker(v CircuitBreakerSchema)`

SetCircuitBreaker sets CircuitBreaker field to given value.

### HasCircuitBreaker

`func (o *ActionRetryConfigSchema) HasCircuitBreaker() bool`

HasCircuitBreaker returns a boolean if a field has been set.

### SetCircuitBreakerNil

`func (o *ActionRetryConfigSchema) SetCircuitBreakerNil(b bool)`

 SetCircuitBreakerNil sets the value for CircuitBreaker to be an explicit nil

### UnsetCircuitBreaker
`func (o *ActionRetryConfigSchema) UnsetCircuitBreaker()`

UnsetCircuitBreaker ensures that no value is present for CircuitBreaker, not even an explicit nil
### GetRetryPolicyName

`func (o *ActionRetryConfigSchema) GetRetryPolicyName() string`

GetRetryPolicyName returns the RetryPolicyName field if non-nil, zero value otherwise.

### GetRetryPolicyNameOk

`func (o *ActionRetryConfigSchema) GetRetryPolicyNameOk() (*string, bool)`

GetRetryPolicyNameOk returns a tuple with the RetryPolicyName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRetryPolicyName

`func (o *ActionRetryConfigSchema) SetRetryPolicyName(v string)`

SetRetryPolicyName sets RetryPolicyName field to given value.

### HasRetryPolicyName

`func (o *ActionRetryConfigSchema) HasRetryPolicyName() bool`

HasRetryPolicyName returns a boolean if a field has been set.

### SetRetryPolicyNameNil

`func (o *ActionRetryConfigSchema) SetRetryPolicyNameNil(b bool)`

 SetRetryPolicyNameNil sets the value for RetryPolicyName to be an explicit nil

### UnsetRetryPolicyName
`func (o *ActionRetryConfigSchema) UnsetRetryPolicyName()`

UnsetRetryPolicyName ensures that no value is present for RetryPolicyName, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


