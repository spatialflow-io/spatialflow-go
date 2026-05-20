# IntegrationErrorStatsOut

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ErrorCount24h** | **int32** |  | 
**LastErrorAt** | Pointer to **NullableString** |  | [optional] 

## Methods

### NewIntegrationErrorStatsOut

`func NewIntegrationErrorStatsOut(errorCount24h int32, ) *IntegrationErrorStatsOut`

NewIntegrationErrorStatsOut instantiates a new IntegrationErrorStatsOut object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIntegrationErrorStatsOutWithDefaults

`func NewIntegrationErrorStatsOutWithDefaults() *IntegrationErrorStatsOut`

NewIntegrationErrorStatsOutWithDefaults instantiates a new IntegrationErrorStatsOut object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetErrorCount24h

`func (o *IntegrationErrorStatsOut) GetErrorCount24h() int32`

GetErrorCount24h returns the ErrorCount24h field if non-nil, zero value otherwise.

### GetErrorCount24hOk

`func (o *IntegrationErrorStatsOut) GetErrorCount24hOk() (*int32, bool)`

GetErrorCount24hOk returns a tuple with the ErrorCount24h field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrorCount24h

`func (o *IntegrationErrorStatsOut) SetErrorCount24h(v int32)`

SetErrorCount24h sets ErrorCount24h field to given value.


### GetLastErrorAt

`func (o *IntegrationErrorStatsOut) GetLastErrorAt() string`

GetLastErrorAt returns the LastErrorAt field if non-nil, zero value otherwise.

### GetLastErrorAtOk

`func (o *IntegrationErrorStatsOut) GetLastErrorAtOk() (*string, bool)`

GetLastErrorAtOk returns a tuple with the LastErrorAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastErrorAt

`func (o *IntegrationErrorStatsOut) SetLastErrorAt(v string)`

SetLastErrorAt sets LastErrorAt field to given value.

### HasLastErrorAt

`func (o *IntegrationErrorStatsOut) HasLastErrorAt() bool`

HasLastErrorAt returns a boolean if a field has been set.

### SetLastErrorAtNil

`func (o *IntegrationErrorStatsOut) SetLastErrorAtNil(b bool)`

 SetLastErrorAtNil sets the value for LastErrorAt to be an explicit nil

### UnsetLastErrorAt
`func (o *IntegrationErrorStatsOut) UnsetLastErrorAt()`

UnsetLastErrorAt ensures that no value is present for LastErrorAt, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


