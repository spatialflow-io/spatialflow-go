# APIUsageStats

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**TotalApiKeys** | **int32** |  | 
**ActiveApiKeys** | **int32** |  | 
**TotalApiCalls** | **int32** |  | 
**LastApiCall** | **NullableString** |  | 

## Methods

### NewAPIUsageStats

`func NewAPIUsageStats(totalApiKeys int32, activeApiKeys int32, totalApiCalls int32, lastApiCall NullableString, ) *APIUsageStats`

NewAPIUsageStats instantiates a new APIUsageStats object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAPIUsageStatsWithDefaults

`func NewAPIUsageStatsWithDefaults() *APIUsageStats`

NewAPIUsageStatsWithDefaults instantiates a new APIUsageStats object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTotalApiKeys

`func (o *APIUsageStats) GetTotalApiKeys() int32`

GetTotalApiKeys returns the TotalApiKeys field if non-nil, zero value otherwise.

### GetTotalApiKeysOk

`func (o *APIUsageStats) GetTotalApiKeysOk() (*int32, bool)`

GetTotalApiKeysOk returns a tuple with the TotalApiKeys field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalApiKeys

`func (o *APIUsageStats) SetTotalApiKeys(v int32)`

SetTotalApiKeys sets TotalApiKeys field to given value.


### GetActiveApiKeys

`func (o *APIUsageStats) GetActiveApiKeys() int32`

GetActiveApiKeys returns the ActiveApiKeys field if non-nil, zero value otherwise.

### GetActiveApiKeysOk

`func (o *APIUsageStats) GetActiveApiKeysOk() (*int32, bool)`

GetActiveApiKeysOk returns a tuple with the ActiveApiKeys field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActiveApiKeys

`func (o *APIUsageStats) SetActiveApiKeys(v int32)`

SetActiveApiKeys sets ActiveApiKeys field to given value.


### GetTotalApiCalls

`func (o *APIUsageStats) GetTotalApiCalls() int32`

GetTotalApiCalls returns the TotalApiCalls field if non-nil, zero value otherwise.

### GetTotalApiCallsOk

`func (o *APIUsageStats) GetTotalApiCallsOk() (*int32, bool)`

GetTotalApiCallsOk returns a tuple with the TotalApiCalls field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalApiCalls

`func (o *APIUsageStats) SetTotalApiCalls(v int32)`

SetTotalApiCalls sets TotalApiCalls field to given value.


### GetLastApiCall

`func (o *APIUsageStats) GetLastApiCall() string`

GetLastApiCall returns the LastApiCall field if non-nil, zero value otherwise.

### GetLastApiCallOk

`func (o *APIUsageStats) GetLastApiCallOk() (*string, bool)`

GetLastApiCallOk returns a tuple with the LastApiCall field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastApiCall

`func (o *APIUsageStats) SetLastApiCall(v string)`

SetLastApiCall sets LastApiCall field to given value.


### SetLastApiCallNil

`func (o *APIUsageStats) SetLastApiCallNil(b bool)`

 SetLastApiCallNil sets the value for LastApiCall to be an explicit nil

### UnsetLastApiCall
`func (o *APIUsageStats) UnsetLastApiCall()`

UnsetLastApiCall ensures that no value is present for LastApiCall, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


