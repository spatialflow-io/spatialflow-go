# DeviceSessionOut

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** |  | 
**StartedAt** | **time.Time** |  | 
**EndedAt** | **time.Time** |  | 
**DurationSeconds** | **int32** |  | 
**LocationCount** | **int32** |  | 
**DistanceMeters** | Pointer to **NullableFloat32** |  | [optional] 
**Notes** | **string** |  | 
**HasTrackGeometry** | Pointer to **bool** |  | [optional] [default to false]

## Methods

### NewDeviceSessionOut

`func NewDeviceSessionOut(id string, startedAt time.Time, endedAt time.Time, durationSeconds int32, locationCount int32, notes string, ) *DeviceSessionOut`

NewDeviceSessionOut instantiates a new DeviceSessionOut object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDeviceSessionOutWithDefaults

`func NewDeviceSessionOutWithDefaults() *DeviceSessionOut`

NewDeviceSessionOutWithDefaults instantiates a new DeviceSessionOut object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *DeviceSessionOut) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *DeviceSessionOut) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *DeviceSessionOut) SetId(v string)`

SetId sets Id field to given value.


### GetStartedAt

`func (o *DeviceSessionOut) GetStartedAt() time.Time`

GetStartedAt returns the StartedAt field if non-nil, zero value otherwise.

### GetStartedAtOk

`func (o *DeviceSessionOut) GetStartedAtOk() (*time.Time, bool)`

GetStartedAtOk returns a tuple with the StartedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartedAt

`func (o *DeviceSessionOut) SetStartedAt(v time.Time)`

SetStartedAt sets StartedAt field to given value.


### GetEndedAt

`func (o *DeviceSessionOut) GetEndedAt() time.Time`

GetEndedAt returns the EndedAt field if non-nil, zero value otherwise.

### GetEndedAtOk

`func (o *DeviceSessionOut) GetEndedAtOk() (*time.Time, bool)`

GetEndedAtOk returns a tuple with the EndedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndedAt

`func (o *DeviceSessionOut) SetEndedAt(v time.Time)`

SetEndedAt sets EndedAt field to given value.


### GetDurationSeconds

`func (o *DeviceSessionOut) GetDurationSeconds() int32`

GetDurationSeconds returns the DurationSeconds field if non-nil, zero value otherwise.

### GetDurationSecondsOk

`func (o *DeviceSessionOut) GetDurationSecondsOk() (*int32, bool)`

GetDurationSecondsOk returns a tuple with the DurationSeconds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDurationSeconds

`func (o *DeviceSessionOut) SetDurationSeconds(v int32)`

SetDurationSeconds sets DurationSeconds field to given value.


### GetLocationCount

`func (o *DeviceSessionOut) GetLocationCount() int32`

GetLocationCount returns the LocationCount field if non-nil, zero value otherwise.

### GetLocationCountOk

`func (o *DeviceSessionOut) GetLocationCountOk() (*int32, bool)`

GetLocationCountOk returns a tuple with the LocationCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocationCount

`func (o *DeviceSessionOut) SetLocationCount(v int32)`

SetLocationCount sets LocationCount field to given value.


### GetDistanceMeters

`func (o *DeviceSessionOut) GetDistanceMeters() float32`

GetDistanceMeters returns the DistanceMeters field if non-nil, zero value otherwise.

### GetDistanceMetersOk

`func (o *DeviceSessionOut) GetDistanceMetersOk() (*float32, bool)`

GetDistanceMetersOk returns a tuple with the DistanceMeters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDistanceMeters

`func (o *DeviceSessionOut) SetDistanceMeters(v float32)`

SetDistanceMeters sets DistanceMeters field to given value.

### HasDistanceMeters

`func (o *DeviceSessionOut) HasDistanceMeters() bool`

HasDistanceMeters returns a boolean if a field has been set.

### SetDistanceMetersNil

`func (o *DeviceSessionOut) SetDistanceMetersNil(b bool)`

 SetDistanceMetersNil sets the value for DistanceMeters to be an explicit nil

### UnsetDistanceMeters
`func (o *DeviceSessionOut) UnsetDistanceMeters()`

UnsetDistanceMeters ensures that no value is present for DistanceMeters, not even an explicit nil
### GetNotes

`func (o *DeviceSessionOut) GetNotes() string`

GetNotes returns the Notes field if non-nil, zero value otherwise.

### GetNotesOk

`func (o *DeviceSessionOut) GetNotesOk() (*string, bool)`

GetNotesOk returns a tuple with the Notes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotes

`func (o *DeviceSessionOut) SetNotes(v string)`

SetNotes sets Notes field to given value.


### GetHasTrackGeometry

`func (o *DeviceSessionOut) GetHasTrackGeometry() bool`

GetHasTrackGeometry returns the HasTrackGeometry field if non-nil, zero value otherwise.

### GetHasTrackGeometryOk

`func (o *DeviceSessionOut) GetHasTrackGeometryOk() (*bool, bool)`

GetHasTrackGeometryOk returns a tuple with the HasTrackGeometry field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasTrackGeometry

`func (o *DeviceSessionOut) SetHasTrackGeometry(v bool)`

SetHasTrackGeometry sets HasTrackGeometry field to given value.

### HasHasTrackGeometry

`func (o *DeviceSessionOut) HasHasTrackGeometry() bool`

HasHasTrackGeometry returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


