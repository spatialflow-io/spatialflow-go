# IncidentOut

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** |  | 
**SignalType** | **string** |  | 
**SourceIds** | **[]string** |  | 
**Severity** | **string** |  | 
**OpenedAt** | **time.Time** |  | 
**AcknowledgedAt** | Pointer to **NullableTime** |  | [optional] 
**AcknowledgedBy** | Pointer to **NullableString** |  | [optional] 
**MutedUntil** | Pointer to **NullableTime** |  | [optional] 
**OwnerId** | Pointer to **NullableString** |  | [optional] 
**ResolvedAt** | Pointer to **NullableTime** |  | [optional] 
**Payload** | **map[string]interface{}** |  | 
**Status** | **string** |  | 

## Methods

### NewIncidentOut

`func NewIncidentOut(id string, signalType string, sourceIds []string, severity string, openedAt time.Time, payload map[string]interface{}, status string, ) *IncidentOut`

NewIncidentOut instantiates a new IncidentOut object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIncidentOutWithDefaults

`func NewIncidentOutWithDefaults() *IncidentOut`

NewIncidentOutWithDefaults instantiates a new IncidentOut object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *IncidentOut) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *IncidentOut) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *IncidentOut) SetId(v string)`

SetId sets Id field to given value.


### GetSignalType

`func (o *IncidentOut) GetSignalType() string`

GetSignalType returns the SignalType field if non-nil, zero value otherwise.

### GetSignalTypeOk

`func (o *IncidentOut) GetSignalTypeOk() (*string, bool)`

GetSignalTypeOk returns a tuple with the SignalType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSignalType

`func (o *IncidentOut) SetSignalType(v string)`

SetSignalType sets SignalType field to given value.


### GetSourceIds

`func (o *IncidentOut) GetSourceIds() []string`

GetSourceIds returns the SourceIds field if non-nil, zero value otherwise.

### GetSourceIdsOk

`func (o *IncidentOut) GetSourceIdsOk() (*[]string, bool)`

GetSourceIdsOk returns a tuple with the SourceIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceIds

`func (o *IncidentOut) SetSourceIds(v []string)`

SetSourceIds sets SourceIds field to given value.


### GetSeverity

`func (o *IncidentOut) GetSeverity() string`

GetSeverity returns the Severity field if non-nil, zero value otherwise.

### GetSeverityOk

`func (o *IncidentOut) GetSeverityOk() (*string, bool)`

GetSeverityOk returns a tuple with the Severity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSeverity

`func (o *IncidentOut) SetSeverity(v string)`

SetSeverity sets Severity field to given value.


### GetOpenedAt

`func (o *IncidentOut) GetOpenedAt() time.Time`

GetOpenedAt returns the OpenedAt field if non-nil, zero value otherwise.

### GetOpenedAtOk

`func (o *IncidentOut) GetOpenedAtOk() (*time.Time, bool)`

GetOpenedAtOk returns a tuple with the OpenedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOpenedAt

`func (o *IncidentOut) SetOpenedAt(v time.Time)`

SetOpenedAt sets OpenedAt field to given value.


### GetAcknowledgedAt

`func (o *IncidentOut) GetAcknowledgedAt() time.Time`

GetAcknowledgedAt returns the AcknowledgedAt field if non-nil, zero value otherwise.

### GetAcknowledgedAtOk

`func (o *IncidentOut) GetAcknowledgedAtOk() (*time.Time, bool)`

GetAcknowledgedAtOk returns a tuple with the AcknowledgedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAcknowledgedAt

`func (o *IncidentOut) SetAcknowledgedAt(v time.Time)`

SetAcknowledgedAt sets AcknowledgedAt field to given value.

### HasAcknowledgedAt

`func (o *IncidentOut) HasAcknowledgedAt() bool`

HasAcknowledgedAt returns a boolean if a field has been set.

### SetAcknowledgedAtNil

`func (o *IncidentOut) SetAcknowledgedAtNil(b bool)`

 SetAcknowledgedAtNil sets the value for AcknowledgedAt to be an explicit nil

### UnsetAcknowledgedAt
`func (o *IncidentOut) UnsetAcknowledgedAt()`

UnsetAcknowledgedAt ensures that no value is present for AcknowledgedAt, not even an explicit nil
### GetAcknowledgedBy

`func (o *IncidentOut) GetAcknowledgedBy() string`

GetAcknowledgedBy returns the AcknowledgedBy field if non-nil, zero value otherwise.

### GetAcknowledgedByOk

`func (o *IncidentOut) GetAcknowledgedByOk() (*string, bool)`

GetAcknowledgedByOk returns a tuple with the AcknowledgedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAcknowledgedBy

`func (o *IncidentOut) SetAcknowledgedBy(v string)`

SetAcknowledgedBy sets AcknowledgedBy field to given value.

### HasAcknowledgedBy

`func (o *IncidentOut) HasAcknowledgedBy() bool`

HasAcknowledgedBy returns a boolean if a field has been set.

### SetAcknowledgedByNil

`func (o *IncidentOut) SetAcknowledgedByNil(b bool)`

 SetAcknowledgedByNil sets the value for AcknowledgedBy to be an explicit nil

### UnsetAcknowledgedBy
`func (o *IncidentOut) UnsetAcknowledgedBy()`

UnsetAcknowledgedBy ensures that no value is present for AcknowledgedBy, not even an explicit nil
### GetMutedUntil

`func (o *IncidentOut) GetMutedUntil() time.Time`

GetMutedUntil returns the MutedUntil field if non-nil, zero value otherwise.

### GetMutedUntilOk

`func (o *IncidentOut) GetMutedUntilOk() (*time.Time, bool)`

GetMutedUntilOk returns a tuple with the MutedUntil field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMutedUntil

`func (o *IncidentOut) SetMutedUntil(v time.Time)`

SetMutedUntil sets MutedUntil field to given value.

### HasMutedUntil

`func (o *IncidentOut) HasMutedUntil() bool`

HasMutedUntil returns a boolean if a field has been set.

### SetMutedUntilNil

`func (o *IncidentOut) SetMutedUntilNil(b bool)`

 SetMutedUntilNil sets the value for MutedUntil to be an explicit nil

### UnsetMutedUntil
`func (o *IncidentOut) UnsetMutedUntil()`

UnsetMutedUntil ensures that no value is present for MutedUntil, not even an explicit nil
### GetOwnerId

`func (o *IncidentOut) GetOwnerId() string`

GetOwnerId returns the OwnerId field if non-nil, zero value otherwise.

### GetOwnerIdOk

`func (o *IncidentOut) GetOwnerIdOk() (*string, bool)`

GetOwnerIdOk returns a tuple with the OwnerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwnerId

`func (o *IncidentOut) SetOwnerId(v string)`

SetOwnerId sets OwnerId field to given value.

### HasOwnerId

`func (o *IncidentOut) HasOwnerId() bool`

HasOwnerId returns a boolean if a field has been set.

### SetOwnerIdNil

`func (o *IncidentOut) SetOwnerIdNil(b bool)`

 SetOwnerIdNil sets the value for OwnerId to be an explicit nil

### UnsetOwnerId
`func (o *IncidentOut) UnsetOwnerId()`

UnsetOwnerId ensures that no value is present for OwnerId, not even an explicit nil
### GetResolvedAt

`func (o *IncidentOut) GetResolvedAt() time.Time`

GetResolvedAt returns the ResolvedAt field if non-nil, zero value otherwise.

### GetResolvedAtOk

`func (o *IncidentOut) GetResolvedAtOk() (*time.Time, bool)`

GetResolvedAtOk returns a tuple with the ResolvedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResolvedAt

`func (o *IncidentOut) SetResolvedAt(v time.Time)`

SetResolvedAt sets ResolvedAt field to given value.

### HasResolvedAt

`func (o *IncidentOut) HasResolvedAt() bool`

HasResolvedAt returns a boolean if a field has been set.

### SetResolvedAtNil

`func (o *IncidentOut) SetResolvedAtNil(b bool)`

 SetResolvedAtNil sets the value for ResolvedAt to be an explicit nil

### UnsetResolvedAt
`func (o *IncidentOut) UnsetResolvedAt()`

UnsetResolvedAt ensures that no value is present for ResolvedAt, not even an explicit nil
### GetPayload

`func (o *IncidentOut) GetPayload() map[string]interface{}`

GetPayload returns the Payload field if non-nil, zero value otherwise.

### GetPayloadOk

`func (o *IncidentOut) GetPayloadOk() (*map[string]interface{}, bool)`

GetPayloadOk returns a tuple with the Payload field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPayload

`func (o *IncidentOut) SetPayload(v map[string]interface{})`

SetPayload sets Payload field to given value.


### GetStatus

`func (o *IncidentOut) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *IncidentOut) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *IncidentOut) SetStatus(v string)`

SetStatus sets Status field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


