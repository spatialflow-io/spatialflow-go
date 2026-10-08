# IncidentBulkIn

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Action** | **string** |  | 
**SignalType** | Pointer to **NullableString** |  | [optional] 
**Ids** | Pointer to **[]string** |  | [optional] 
**Status** | Pointer to **string** |  | [optional] [default to "open"]
**Severity** | Pointer to **NullableString** |  | [optional] 
**AsOf** | **time.Time** |  | 
**ExpectedCount** | Pointer to **NullableInt32** |  | [optional] 

## Methods

### NewIncidentBulkIn

`func NewIncidentBulkIn(action string, asOf time.Time, ) *IncidentBulkIn`

NewIncidentBulkIn instantiates a new IncidentBulkIn object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIncidentBulkInWithDefaults

`func NewIncidentBulkInWithDefaults() *IncidentBulkIn`

NewIncidentBulkInWithDefaults instantiates a new IncidentBulkIn object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAction

`func (o *IncidentBulkIn) GetAction() string`

GetAction returns the Action field if non-nil, zero value otherwise.

### GetActionOk

`func (o *IncidentBulkIn) GetActionOk() (*string, bool)`

GetActionOk returns a tuple with the Action field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAction

`func (o *IncidentBulkIn) SetAction(v string)`

SetAction sets Action field to given value.


### GetSignalType

`func (o *IncidentBulkIn) GetSignalType() string`

GetSignalType returns the SignalType field if non-nil, zero value otherwise.

### GetSignalTypeOk

`func (o *IncidentBulkIn) GetSignalTypeOk() (*string, bool)`

GetSignalTypeOk returns a tuple with the SignalType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSignalType

`func (o *IncidentBulkIn) SetSignalType(v string)`

SetSignalType sets SignalType field to given value.

### HasSignalType

`func (o *IncidentBulkIn) HasSignalType() bool`

HasSignalType returns a boolean if a field has been set.

### SetSignalTypeNil

`func (o *IncidentBulkIn) SetSignalTypeNil(b bool)`

 SetSignalTypeNil sets the value for SignalType to be an explicit nil

### UnsetSignalType
`func (o *IncidentBulkIn) UnsetSignalType()`

UnsetSignalType ensures that no value is present for SignalType, not even an explicit nil
### GetIds

`func (o *IncidentBulkIn) GetIds() []string`

GetIds returns the Ids field if non-nil, zero value otherwise.

### GetIdsOk

`func (o *IncidentBulkIn) GetIdsOk() (*[]string, bool)`

GetIdsOk returns a tuple with the Ids field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIds

`func (o *IncidentBulkIn) SetIds(v []string)`

SetIds sets Ids field to given value.

### HasIds

`func (o *IncidentBulkIn) HasIds() bool`

HasIds returns a boolean if a field has been set.

### SetIdsNil

`func (o *IncidentBulkIn) SetIdsNil(b bool)`

 SetIdsNil sets the value for Ids to be an explicit nil

### UnsetIds
`func (o *IncidentBulkIn) UnsetIds()`

UnsetIds ensures that no value is present for Ids, not even an explicit nil
### GetStatus

`func (o *IncidentBulkIn) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *IncidentBulkIn) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *IncidentBulkIn) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *IncidentBulkIn) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetSeverity

`func (o *IncidentBulkIn) GetSeverity() string`

GetSeverity returns the Severity field if non-nil, zero value otherwise.

### GetSeverityOk

`func (o *IncidentBulkIn) GetSeverityOk() (*string, bool)`

GetSeverityOk returns a tuple with the Severity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSeverity

`func (o *IncidentBulkIn) SetSeverity(v string)`

SetSeverity sets Severity field to given value.

### HasSeverity

`func (o *IncidentBulkIn) HasSeverity() bool`

HasSeverity returns a boolean if a field has been set.

### SetSeverityNil

`func (o *IncidentBulkIn) SetSeverityNil(b bool)`

 SetSeverityNil sets the value for Severity to be an explicit nil

### UnsetSeverity
`func (o *IncidentBulkIn) UnsetSeverity()`

UnsetSeverity ensures that no value is present for Severity, not even an explicit nil
### GetAsOf

`func (o *IncidentBulkIn) GetAsOf() time.Time`

GetAsOf returns the AsOf field if non-nil, zero value otherwise.

### GetAsOfOk

`func (o *IncidentBulkIn) GetAsOfOk() (*time.Time, bool)`

GetAsOfOk returns a tuple with the AsOf field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAsOf

`func (o *IncidentBulkIn) SetAsOf(v time.Time)`

SetAsOf sets AsOf field to given value.


### GetExpectedCount

`func (o *IncidentBulkIn) GetExpectedCount() int32`

GetExpectedCount returns the ExpectedCount field if non-nil, zero value otherwise.

### GetExpectedCountOk

`func (o *IncidentBulkIn) GetExpectedCountOk() (*int32, bool)`

GetExpectedCountOk returns a tuple with the ExpectedCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpectedCount

`func (o *IncidentBulkIn) SetExpectedCount(v int32)`

SetExpectedCount sets ExpectedCount field to given value.

### HasExpectedCount

`func (o *IncidentBulkIn) HasExpectedCount() bool`

HasExpectedCount returns a boolean if a field has been set.

### SetExpectedCountNil

`func (o *IncidentBulkIn) SetExpectedCountNil(b bool)`

 SetExpectedCountNil sets the value for ExpectedCount to be an explicit nil

### UnsetExpectedCount
`func (o *IncidentBulkIn) UnsetExpectedCount()`

UnsetExpectedCount ensures that no value is present for ExpectedCount, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


