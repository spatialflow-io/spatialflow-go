# IntegrationStatsSchema

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**TotalUses** | **int32** |  | 
**SuccessfulUses** | **int32** |  | 
**SuccessRate** | **float32** |  | 
**AverageDurationMs** | **NullableFloat32** |  | 
**HealthStatus** | **string** |  | 
**LastUsedAt** | **NullableString** |  | 
**LastHealthCheckAt** | **NullableString** |  | 
**RecentErrors** | **[]map[string]interface{}** |  | 

## Methods

### NewIntegrationStatsSchema

`func NewIntegrationStatsSchema(totalUses int32, successfulUses int32, successRate float32, averageDurationMs NullableFloat32, healthStatus string, lastUsedAt NullableString, lastHealthCheckAt NullableString, recentErrors []*map[string]interface{}, ) *IntegrationStatsSchema`

NewIntegrationStatsSchema instantiates a new IntegrationStatsSchema object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIntegrationStatsSchemaWithDefaults

`func NewIntegrationStatsSchemaWithDefaults() *IntegrationStatsSchema`

NewIntegrationStatsSchemaWithDefaults instantiates a new IntegrationStatsSchema object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTotalUses

`func (o *IntegrationStatsSchema) GetTotalUses() int32`

GetTotalUses returns the TotalUses field if non-nil, zero value otherwise.

### GetTotalUsesOk

`func (o *IntegrationStatsSchema) GetTotalUsesOk() (*int32, bool)`

GetTotalUsesOk returns a tuple with the TotalUses field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalUses

`func (o *IntegrationStatsSchema) SetTotalUses(v int32)`

SetTotalUses sets TotalUses field to given value.


### GetSuccessfulUses

`func (o *IntegrationStatsSchema) GetSuccessfulUses() int32`

GetSuccessfulUses returns the SuccessfulUses field if non-nil, zero value otherwise.

### GetSuccessfulUsesOk

`func (o *IntegrationStatsSchema) GetSuccessfulUsesOk() (*int32, bool)`

GetSuccessfulUsesOk returns a tuple with the SuccessfulUses field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSuccessfulUses

`func (o *IntegrationStatsSchema) SetSuccessfulUses(v int32)`

SetSuccessfulUses sets SuccessfulUses field to given value.


### GetSuccessRate

`func (o *IntegrationStatsSchema) GetSuccessRate() float32`

GetSuccessRate returns the SuccessRate field if non-nil, zero value otherwise.

### GetSuccessRateOk

`func (o *IntegrationStatsSchema) GetSuccessRateOk() (*float32, bool)`

GetSuccessRateOk returns a tuple with the SuccessRate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSuccessRate

`func (o *IntegrationStatsSchema) SetSuccessRate(v float32)`

SetSuccessRate sets SuccessRate field to given value.


### GetAverageDurationMs

`func (o *IntegrationStatsSchema) GetAverageDurationMs() float32`

GetAverageDurationMs returns the AverageDurationMs field if non-nil, zero value otherwise.

### GetAverageDurationMsOk

`func (o *IntegrationStatsSchema) GetAverageDurationMsOk() (*float32, bool)`

GetAverageDurationMsOk returns a tuple with the AverageDurationMs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAverageDurationMs

`func (o *IntegrationStatsSchema) SetAverageDurationMs(v float32)`

SetAverageDurationMs sets AverageDurationMs field to given value.


### SetAverageDurationMsNil

`func (o *IntegrationStatsSchema) SetAverageDurationMsNil(b bool)`

 SetAverageDurationMsNil sets the value for AverageDurationMs to be an explicit nil

### UnsetAverageDurationMs
`func (o *IntegrationStatsSchema) UnsetAverageDurationMs()`

UnsetAverageDurationMs ensures that no value is present for AverageDurationMs, not even an explicit nil
### GetHealthStatus

`func (o *IntegrationStatsSchema) GetHealthStatus() string`

GetHealthStatus returns the HealthStatus field if non-nil, zero value otherwise.

### GetHealthStatusOk

`func (o *IntegrationStatsSchema) GetHealthStatusOk() (*string, bool)`

GetHealthStatusOk returns a tuple with the HealthStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHealthStatus

`func (o *IntegrationStatsSchema) SetHealthStatus(v string)`

SetHealthStatus sets HealthStatus field to given value.


### GetLastUsedAt

`func (o *IntegrationStatsSchema) GetLastUsedAt() string`

GetLastUsedAt returns the LastUsedAt field if non-nil, zero value otherwise.

### GetLastUsedAtOk

`func (o *IntegrationStatsSchema) GetLastUsedAtOk() (*string, bool)`

GetLastUsedAtOk returns a tuple with the LastUsedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastUsedAt

`func (o *IntegrationStatsSchema) SetLastUsedAt(v string)`

SetLastUsedAt sets LastUsedAt field to given value.


### SetLastUsedAtNil

`func (o *IntegrationStatsSchema) SetLastUsedAtNil(b bool)`

 SetLastUsedAtNil sets the value for LastUsedAt to be an explicit nil

### UnsetLastUsedAt
`func (o *IntegrationStatsSchema) UnsetLastUsedAt()`

UnsetLastUsedAt ensures that no value is present for LastUsedAt, not even an explicit nil
### GetLastHealthCheckAt

`func (o *IntegrationStatsSchema) GetLastHealthCheckAt() string`

GetLastHealthCheckAt returns the LastHealthCheckAt field if non-nil, zero value otherwise.

### GetLastHealthCheckAtOk

`func (o *IntegrationStatsSchema) GetLastHealthCheckAtOk() (*string, bool)`

GetLastHealthCheckAtOk returns a tuple with the LastHealthCheckAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastHealthCheckAt

`func (o *IntegrationStatsSchema) SetLastHealthCheckAt(v string)`

SetLastHealthCheckAt sets LastHealthCheckAt field to given value.


### SetLastHealthCheckAtNil

`func (o *IntegrationStatsSchema) SetLastHealthCheckAtNil(b bool)`

 SetLastHealthCheckAtNil sets the value for LastHealthCheckAt to be an explicit nil

### UnsetLastHealthCheckAt
`func (o *IntegrationStatsSchema) UnsetLastHealthCheckAt()`

UnsetLastHealthCheckAt ensures that no value is present for LastHealthCheckAt, not even an explicit nil
### GetRecentErrors

`func (o *IntegrationStatsSchema) GetRecentErrors() []*map[string]interface{}`

GetRecentErrors returns the RecentErrors field if non-nil, zero value otherwise.

### GetRecentErrorsOk

`func (o *IntegrationStatsSchema) GetRecentErrorsOk() (*[]*map[string]interface{}, bool)`

GetRecentErrorsOk returns a tuple with the RecentErrors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecentErrors

`func (o *IntegrationStatsSchema) SetRecentErrors(v []*map[string]interface{})`

SetRecentErrors sets RecentErrors field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


