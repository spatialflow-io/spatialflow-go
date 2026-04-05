# SignalEventDetailOut

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** |  | 
**DeviceId** | **string** |  | 
**DeviceName** | **string** |  | 
**SignalType** | **string** |  | 
**State** | **string** |  | 
**StartedAt** | **time.Time** |  | 
**ConfirmedAt** | Pointer to **NullableTime** |  | [optional] 
**EndedAt** | Pointer to **NullableTime** |  | [optional] 
**GeofenceId** | Pointer to **NullableString** |  | [optional] 
**GeofenceName** | Pointer to **NullableString** |  | [optional] 
**HasExplanation** | **bool** |  | 
**Metadata** | Pointer to **map[string]interface{}** |  | [optional] [default to {}]
**CreatedAt** | **time.Time** |  | 
**UpdatedAt** | **time.Time** |  | 
**ReasonCodes** | Pointer to **[]interface{}** |  | [optional] [default to []]
**Explanation** | Pointer to **map[string]interface{}** |  | [optional] [default to {}]
**ContributingLocations** | Pointer to [**[]ContributingLocationOut**](ContributingLocationOut.md) |  | [optional] [default to []]
**GeofenceGeometry** | Pointer to [**NullableGeofenceGeometryOut**](GeofenceGeometryOut.md) |  | [optional] 

## Methods

### NewSignalEventDetailOut

`func NewSignalEventDetailOut(id string, deviceId string, deviceName string, signalType string, state string, startedAt time.Time, hasExplanation bool, createdAt time.Time, updatedAt time.Time, ) *SignalEventDetailOut`

NewSignalEventDetailOut instantiates a new SignalEventDetailOut object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSignalEventDetailOutWithDefaults

`func NewSignalEventDetailOutWithDefaults() *SignalEventDetailOut`

NewSignalEventDetailOutWithDefaults instantiates a new SignalEventDetailOut object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *SignalEventDetailOut) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *SignalEventDetailOut) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *SignalEventDetailOut) SetId(v string)`

SetId sets Id field to given value.


### GetDeviceId

`func (o *SignalEventDetailOut) GetDeviceId() string`

GetDeviceId returns the DeviceId field if non-nil, zero value otherwise.

### GetDeviceIdOk

`func (o *SignalEventDetailOut) GetDeviceIdOk() (*string, bool)`

GetDeviceIdOk returns a tuple with the DeviceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeviceId

`func (o *SignalEventDetailOut) SetDeviceId(v string)`

SetDeviceId sets DeviceId field to given value.


### GetDeviceName

`func (o *SignalEventDetailOut) GetDeviceName() string`

GetDeviceName returns the DeviceName field if non-nil, zero value otherwise.

### GetDeviceNameOk

`func (o *SignalEventDetailOut) GetDeviceNameOk() (*string, bool)`

GetDeviceNameOk returns a tuple with the DeviceName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeviceName

`func (o *SignalEventDetailOut) SetDeviceName(v string)`

SetDeviceName sets DeviceName field to given value.


### GetSignalType

`func (o *SignalEventDetailOut) GetSignalType() string`

GetSignalType returns the SignalType field if non-nil, zero value otherwise.

### GetSignalTypeOk

`func (o *SignalEventDetailOut) GetSignalTypeOk() (*string, bool)`

GetSignalTypeOk returns a tuple with the SignalType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSignalType

`func (o *SignalEventDetailOut) SetSignalType(v string)`

SetSignalType sets SignalType field to given value.


### GetState

`func (o *SignalEventDetailOut) GetState() string`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *SignalEventDetailOut) GetStateOk() (*string, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *SignalEventDetailOut) SetState(v string)`

SetState sets State field to given value.


### GetStartedAt

`func (o *SignalEventDetailOut) GetStartedAt() time.Time`

GetStartedAt returns the StartedAt field if non-nil, zero value otherwise.

### GetStartedAtOk

`func (o *SignalEventDetailOut) GetStartedAtOk() (*time.Time, bool)`

GetStartedAtOk returns a tuple with the StartedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartedAt

`func (o *SignalEventDetailOut) SetStartedAt(v time.Time)`

SetStartedAt sets StartedAt field to given value.


### GetConfirmedAt

`func (o *SignalEventDetailOut) GetConfirmedAt() time.Time`

GetConfirmedAt returns the ConfirmedAt field if non-nil, zero value otherwise.

### GetConfirmedAtOk

`func (o *SignalEventDetailOut) GetConfirmedAtOk() (*time.Time, bool)`

GetConfirmedAtOk returns a tuple with the ConfirmedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfirmedAt

`func (o *SignalEventDetailOut) SetConfirmedAt(v time.Time)`

SetConfirmedAt sets ConfirmedAt field to given value.

### HasConfirmedAt

`func (o *SignalEventDetailOut) HasConfirmedAt() bool`

HasConfirmedAt returns a boolean if a field has been set.

### SetConfirmedAtNil

`func (o *SignalEventDetailOut) SetConfirmedAtNil(b bool)`

 SetConfirmedAtNil sets the value for ConfirmedAt to be an explicit nil

### UnsetConfirmedAt
`func (o *SignalEventDetailOut) UnsetConfirmedAt()`

UnsetConfirmedAt ensures that no value is present for ConfirmedAt, not even an explicit nil
### GetEndedAt

`func (o *SignalEventDetailOut) GetEndedAt() time.Time`

GetEndedAt returns the EndedAt field if non-nil, zero value otherwise.

### GetEndedAtOk

`func (o *SignalEventDetailOut) GetEndedAtOk() (*time.Time, bool)`

GetEndedAtOk returns a tuple with the EndedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndedAt

`func (o *SignalEventDetailOut) SetEndedAt(v time.Time)`

SetEndedAt sets EndedAt field to given value.

### HasEndedAt

`func (o *SignalEventDetailOut) HasEndedAt() bool`

HasEndedAt returns a boolean if a field has been set.

### SetEndedAtNil

`func (o *SignalEventDetailOut) SetEndedAtNil(b bool)`

 SetEndedAtNil sets the value for EndedAt to be an explicit nil

### UnsetEndedAt
`func (o *SignalEventDetailOut) UnsetEndedAt()`

UnsetEndedAt ensures that no value is present for EndedAt, not even an explicit nil
### GetGeofenceId

`func (o *SignalEventDetailOut) GetGeofenceId() string`

GetGeofenceId returns the GeofenceId field if non-nil, zero value otherwise.

### GetGeofenceIdOk

`func (o *SignalEventDetailOut) GetGeofenceIdOk() (*string, bool)`

GetGeofenceIdOk returns a tuple with the GeofenceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGeofenceId

`func (o *SignalEventDetailOut) SetGeofenceId(v string)`

SetGeofenceId sets GeofenceId field to given value.

### HasGeofenceId

`func (o *SignalEventDetailOut) HasGeofenceId() bool`

HasGeofenceId returns a boolean if a field has been set.

### SetGeofenceIdNil

`func (o *SignalEventDetailOut) SetGeofenceIdNil(b bool)`

 SetGeofenceIdNil sets the value for GeofenceId to be an explicit nil

### UnsetGeofenceId
`func (o *SignalEventDetailOut) UnsetGeofenceId()`

UnsetGeofenceId ensures that no value is present for GeofenceId, not even an explicit nil
### GetGeofenceName

`func (o *SignalEventDetailOut) GetGeofenceName() string`

GetGeofenceName returns the GeofenceName field if non-nil, zero value otherwise.

### GetGeofenceNameOk

`func (o *SignalEventDetailOut) GetGeofenceNameOk() (*string, bool)`

GetGeofenceNameOk returns a tuple with the GeofenceName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGeofenceName

`func (o *SignalEventDetailOut) SetGeofenceName(v string)`

SetGeofenceName sets GeofenceName field to given value.

### HasGeofenceName

`func (o *SignalEventDetailOut) HasGeofenceName() bool`

HasGeofenceName returns a boolean if a field has been set.

### SetGeofenceNameNil

`func (o *SignalEventDetailOut) SetGeofenceNameNil(b bool)`

 SetGeofenceNameNil sets the value for GeofenceName to be an explicit nil

### UnsetGeofenceName
`func (o *SignalEventDetailOut) UnsetGeofenceName()`

UnsetGeofenceName ensures that no value is present for GeofenceName, not even an explicit nil
### GetHasExplanation

`func (o *SignalEventDetailOut) GetHasExplanation() bool`

GetHasExplanation returns the HasExplanation field if non-nil, zero value otherwise.

### GetHasExplanationOk

`func (o *SignalEventDetailOut) GetHasExplanationOk() (*bool, bool)`

GetHasExplanationOk returns a tuple with the HasExplanation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasExplanation

`func (o *SignalEventDetailOut) SetHasExplanation(v bool)`

SetHasExplanation sets HasExplanation field to given value.


### GetMetadata

`func (o *SignalEventDetailOut) GetMetadata() map[string]interface{}`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *SignalEventDetailOut) GetMetadataOk() (*map[string]interface{}, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *SignalEventDetailOut) SetMetadata(v map[string]interface{})`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *SignalEventDetailOut) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### GetCreatedAt

`func (o *SignalEventDetailOut) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *SignalEventDetailOut) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *SignalEventDetailOut) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.


### GetUpdatedAt

`func (o *SignalEventDetailOut) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *SignalEventDetailOut) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *SignalEventDetailOut) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.


### GetReasonCodes

`func (o *SignalEventDetailOut) GetReasonCodes() []interface{}`

GetReasonCodes returns the ReasonCodes field if non-nil, zero value otherwise.

### GetReasonCodesOk

`func (o *SignalEventDetailOut) GetReasonCodesOk() (*[]interface{}, bool)`

GetReasonCodesOk returns a tuple with the ReasonCodes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReasonCodes

`func (o *SignalEventDetailOut) SetReasonCodes(v []interface{})`

SetReasonCodes sets ReasonCodes field to given value.

### HasReasonCodes

`func (o *SignalEventDetailOut) HasReasonCodes() bool`

HasReasonCodes returns a boolean if a field has been set.

### GetExplanation

`func (o *SignalEventDetailOut) GetExplanation() map[string]interface{}`

GetExplanation returns the Explanation field if non-nil, zero value otherwise.

### GetExplanationOk

`func (o *SignalEventDetailOut) GetExplanationOk() (*map[string]interface{}, bool)`

GetExplanationOk returns a tuple with the Explanation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExplanation

`func (o *SignalEventDetailOut) SetExplanation(v map[string]interface{})`

SetExplanation sets Explanation field to given value.

### HasExplanation

`func (o *SignalEventDetailOut) HasExplanation() bool`

HasExplanation returns a boolean if a field has been set.

### GetContributingLocations

`func (o *SignalEventDetailOut) GetContributingLocations() []ContributingLocationOut`

GetContributingLocations returns the ContributingLocations field if non-nil, zero value otherwise.

### GetContributingLocationsOk

`func (o *SignalEventDetailOut) GetContributingLocationsOk() (*[]ContributingLocationOut, bool)`

GetContributingLocationsOk returns a tuple with the ContributingLocations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContributingLocations

`func (o *SignalEventDetailOut) SetContributingLocations(v []ContributingLocationOut)`

SetContributingLocations sets ContributingLocations field to given value.

### HasContributingLocations

`func (o *SignalEventDetailOut) HasContributingLocations() bool`

HasContributingLocations returns a boolean if a field has been set.

### GetGeofenceGeometry

`func (o *SignalEventDetailOut) GetGeofenceGeometry() GeofenceGeometryOut`

GetGeofenceGeometry returns the GeofenceGeometry field if non-nil, zero value otherwise.

### GetGeofenceGeometryOk

`func (o *SignalEventDetailOut) GetGeofenceGeometryOk() (*GeofenceGeometryOut, bool)`

GetGeofenceGeometryOk returns a tuple with the GeofenceGeometry field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGeofenceGeometry

`func (o *SignalEventDetailOut) SetGeofenceGeometry(v GeofenceGeometryOut)`

SetGeofenceGeometry sets GeofenceGeometry field to given value.

### HasGeofenceGeometry

`func (o *SignalEventDetailOut) HasGeofenceGeometry() bool`

HasGeofenceGeometry returns a boolean if a field has been set.

### SetGeofenceGeometryNil

`func (o *SignalEventDetailOut) SetGeofenceGeometryNil(b bool)`

 SetGeofenceGeometryNil sets the value for GeofenceGeometry to be an explicit nil

### UnsetGeofenceGeometry
`func (o *SignalEventDetailOut) UnsetGeofenceGeometry()`

UnsetGeofenceGeometry ensures that no value is present for GeofenceGeometry, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


