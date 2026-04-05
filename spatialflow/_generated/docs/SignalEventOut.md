# SignalEventOut

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

## Methods

### NewSignalEventOut

`func NewSignalEventOut(id string, deviceId string, deviceName string, signalType string, state string, startedAt time.Time, hasExplanation bool, createdAt time.Time, updatedAt time.Time, ) *SignalEventOut`

NewSignalEventOut instantiates a new SignalEventOut object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSignalEventOutWithDefaults

`func NewSignalEventOutWithDefaults() *SignalEventOut`

NewSignalEventOutWithDefaults instantiates a new SignalEventOut object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *SignalEventOut) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *SignalEventOut) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *SignalEventOut) SetId(v string)`

SetId sets Id field to given value.


### GetDeviceId

`func (o *SignalEventOut) GetDeviceId() string`

GetDeviceId returns the DeviceId field if non-nil, zero value otherwise.

### GetDeviceIdOk

`func (o *SignalEventOut) GetDeviceIdOk() (*string, bool)`

GetDeviceIdOk returns a tuple with the DeviceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeviceId

`func (o *SignalEventOut) SetDeviceId(v string)`

SetDeviceId sets DeviceId field to given value.


### GetDeviceName

`func (o *SignalEventOut) GetDeviceName() string`

GetDeviceName returns the DeviceName field if non-nil, zero value otherwise.

### GetDeviceNameOk

`func (o *SignalEventOut) GetDeviceNameOk() (*string, bool)`

GetDeviceNameOk returns a tuple with the DeviceName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeviceName

`func (o *SignalEventOut) SetDeviceName(v string)`

SetDeviceName sets DeviceName field to given value.


### GetSignalType

`func (o *SignalEventOut) GetSignalType() string`

GetSignalType returns the SignalType field if non-nil, zero value otherwise.

### GetSignalTypeOk

`func (o *SignalEventOut) GetSignalTypeOk() (*string, bool)`

GetSignalTypeOk returns a tuple with the SignalType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSignalType

`func (o *SignalEventOut) SetSignalType(v string)`

SetSignalType sets SignalType field to given value.


### GetState

`func (o *SignalEventOut) GetState() string`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *SignalEventOut) GetStateOk() (*string, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *SignalEventOut) SetState(v string)`

SetState sets State field to given value.


### GetStartedAt

`func (o *SignalEventOut) GetStartedAt() time.Time`

GetStartedAt returns the StartedAt field if non-nil, zero value otherwise.

### GetStartedAtOk

`func (o *SignalEventOut) GetStartedAtOk() (*time.Time, bool)`

GetStartedAtOk returns a tuple with the StartedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartedAt

`func (o *SignalEventOut) SetStartedAt(v time.Time)`

SetStartedAt sets StartedAt field to given value.


### GetConfirmedAt

`func (o *SignalEventOut) GetConfirmedAt() time.Time`

GetConfirmedAt returns the ConfirmedAt field if non-nil, zero value otherwise.

### GetConfirmedAtOk

`func (o *SignalEventOut) GetConfirmedAtOk() (*time.Time, bool)`

GetConfirmedAtOk returns a tuple with the ConfirmedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfirmedAt

`func (o *SignalEventOut) SetConfirmedAt(v time.Time)`

SetConfirmedAt sets ConfirmedAt field to given value.

### HasConfirmedAt

`func (o *SignalEventOut) HasConfirmedAt() bool`

HasConfirmedAt returns a boolean if a field has been set.

### SetConfirmedAtNil

`func (o *SignalEventOut) SetConfirmedAtNil(b bool)`

 SetConfirmedAtNil sets the value for ConfirmedAt to be an explicit nil

### UnsetConfirmedAt
`func (o *SignalEventOut) UnsetConfirmedAt()`

UnsetConfirmedAt ensures that no value is present for ConfirmedAt, not even an explicit nil
### GetEndedAt

`func (o *SignalEventOut) GetEndedAt() time.Time`

GetEndedAt returns the EndedAt field if non-nil, zero value otherwise.

### GetEndedAtOk

`func (o *SignalEventOut) GetEndedAtOk() (*time.Time, bool)`

GetEndedAtOk returns a tuple with the EndedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndedAt

`func (o *SignalEventOut) SetEndedAt(v time.Time)`

SetEndedAt sets EndedAt field to given value.

### HasEndedAt

`func (o *SignalEventOut) HasEndedAt() bool`

HasEndedAt returns a boolean if a field has been set.

### SetEndedAtNil

`func (o *SignalEventOut) SetEndedAtNil(b bool)`

 SetEndedAtNil sets the value for EndedAt to be an explicit nil

### UnsetEndedAt
`func (o *SignalEventOut) UnsetEndedAt()`

UnsetEndedAt ensures that no value is present for EndedAt, not even an explicit nil
### GetGeofenceId

`func (o *SignalEventOut) GetGeofenceId() string`

GetGeofenceId returns the GeofenceId field if non-nil, zero value otherwise.

### GetGeofenceIdOk

`func (o *SignalEventOut) GetGeofenceIdOk() (*string, bool)`

GetGeofenceIdOk returns a tuple with the GeofenceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGeofenceId

`func (o *SignalEventOut) SetGeofenceId(v string)`

SetGeofenceId sets GeofenceId field to given value.

### HasGeofenceId

`func (o *SignalEventOut) HasGeofenceId() bool`

HasGeofenceId returns a boolean if a field has been set.

### SetGeofenceIdNil

`func (o *SignalEventOut) SetGeofenceIdNil(b bool)`

 SetGeofenceIdNil sets the value for GeofenceId to be an explicit nil

### UnsetGeofenceId
`func (o *SignalEventOut) UnsetGeofenceId()`

UnsetGeofenceId ensures that no value is present for GeofenceId, not even an explicit nil
### GetGeofenceName

`func (o *SignalEventOut) GetGeofenceName() string`

GetGeofenceName returns the GeofenceName field if non-nil, zero value otherwise.

### GetGeofenceNameOk

`func (o *SignalEventOut) GetGeofenceNameOk() (*string, bool)`

GetGeofenceNameOk returns a tuple with the GeofenceName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGeofenceName

`func (o *SignalEventOut) SetGeofenceName(v string)`

SetGeofenceName sets GeofenceName field to given value.

### HasGeofenceName

`func (o *SignalEventOut) HasGeofenceName() bool`

HasGeofenceName returns a boolean if a field has been set.

### SetGeofenceNameNil

`func (o *SignalEventOut) SetGeofenceNameNil(b bool)`

 SetGeofenceNameNil sets the value for GeofenceName to be an explicit nil

### UnsetGeofenceName
`func (o *SignalEventOut) UnsetGeofenceName()`

UnsetGeofenceName ensures that no value is present for GeofenceName, not even an explicit nil
### GetHasExplanation

`func (o *SignalEventOut) GetHasExplanation() bool`

GetHasExplanation returns the HasExplanation field if non-nil, zero value otherwise.

### GetHasExplanationOk

`func (o *SignalEventOut) GetHasExplanationOk() (*bool, bool)`

GetHasExplanationOk returns a tuple with the HasExplanation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasExplanation

`func (o *SignalEventOut) SetHasExplanation(v bool)`

SetHasExplanation sets HasExplanation field to given value.


### GetMetadata

`func (o *SignalEventOut) GetMetadata() map[string]interface{}`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *SignalEventOut) GetMetadataOk() (*map[string]interface{}, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *SignalEventOut) SetMetadata(v map[string]interface{})`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *SignalEventOut) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### GetCreatedAt

`func (o *SignalEventOut) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *SignalEventOut) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *SignalEventOut) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.


### GetUpdatedAt

`func (o *SignalEventOut) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *SignalEventOut) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *SignalEventOut) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


