# DeviceSessionOut

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** |  | 
**StartedAt** | **time.Time** |  | 
**EndedAt** | Pointer to **NullableTime** |  | [optional] 
**DurationSeconds** | Pointer to **NullableInt32** |  | [optional] 
**LocationCount** | **int32** |  | 
**DistanceMeters** | Pointer to **NullableFloat32** |  | [optional] 
**HasTrackGeometry** | Pointer to **bool** |  | [optional] [default to false]
**AutoClosedAt** | Pointer to **NullableTime** |  | [optional] 
**CloseReason** | Pointer to **string** |  | [optional] [default to ""]
**AutoClosedFromShiftStatus** | Pointer to **NullableString** |  | [optional] 
**PhotoCount** | **NullableInt32** |  | 
**NoteCount** | **NullableInt32** |  | 

## Methods

### NewDeviceSessionOut

`func NewDeviceSessionOut(id string, startedAt time.Time, locationCount int32, photoCount NullableInt32, noteCount NullableInt32, ) *DeviceSessionOut`

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

### HasEndedAt

`func (o *DeviceSessionOut) HasEndedAt() bool`

HasEndedAt returns a boolean if a field has been set.

### SetEndedAtNil

`func (o *DeviceSessionOut) SetEndedAtNil(b bool)`

 SetEndedAtNil sets the value for EndedAt to be an explicit nil

### UnsetEndedAt
`func (o *DeviceSessionOut) UnsetEndedAt()`

UnsetEndedAt ensures that no value is present for EndedAt, not even an explicit nil
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

### HasDurationSeconds

`func (o *DeviceSessionOut) HasDurationSeconds() bool`

HasDurationSeconds returns a boolean if a field has been set.

### SetDurationSecondsNil

`func (o *DeviceSessionOut) SetDurationSecondsNil(b bool)`

 SetDurationSecondsNil sets the value for DurationSeconds to be an explicit nil

### UnsetDurationSeconds
`func (o *DeviceSessionOut) UnsetDurationSeconds()`

UnsetDurationSeconds ensures that no value is present for DurationSeconds, not even an explicit nil
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

### GetAutoClosedAt

`func (o *DeviceSessionOut) GetAutoClosedAt() time.Time`

GetAutoClosedAt returns the AutoClosedAt field if non-nil, zero value otherwise.

### GetAutoClosedAtOk

`func (o *DeviceSessionOut) GetAutoClosedAtOk() (*time.Time, bool)`

GetAutoClosedAtOk returns a tuple with the AutoClosedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAutoClosedAt

`func (o *DeviceSessionOut) SetAutoClosedAt(v time.Time)`

SetAutoClosedAt sets AutoClosedAt field to given value.

### HasAutoClosedAt

`func (o *DeviceSessionOut) HasAutoClosedAt() bool`

HasAutoClosedAt returns a boolean if a field has been set.

### SetAutoClosedAtNil

`func (o *DeviceSessionOut) SetAutoClosedAtNil(b bool)`

 SetAutoClosedAtNil sets the value for AutoClosedAt to be an explicit nil

### UnsetAutoClosedAt
`func (o *DeviceSessionOut) UnsetAutoClosedAt()`

UnsetAutoClosedAt ensures that no value is present for AutoClosedAt, not even an explicit nil
### GetCloseReason

`func (o *DeviceSessionOut) GetCloseReason() string`

GetCloseReason returns the CloseReason field if non-nil, zero value otherwise.

### GetCloseReasonOk

`func (o *DeviceSessionOut) GetCloseReasonOk() (*string, bool)`

GetCloseReasonOk returns a tuple with the CloseReason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCloseReason

`func (o *DeviceSessionOut) SetCloseReason(v string)`

SetCloseReason sets CloseReason field to given value.

### HasCloseReason

`func (o *DeviceSessionOut) HasCloseReason() bool`

HasCloseReason returns a boolean if a field has been set.

### GetAutoClosedFromShiftStatus

`func (o *DeviceSessionOut) GetAutoClosedFromShiftStatus() string`

GetAutoClosedFromShiftStatus returns the AutoClosedFromShiftStatus field if non-nil, zero value otherwise.

### GetAutoClosedFromShiftStatusOk

`func (o *DeviceSessionOut) GetAutoClosedFromShiftStatusOk() (*string, bool)`

GetAutoClosedFromShiftStatusOk returns a tuple with the AutoClosedFromShiftStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAutoClosedFromShiftStatus

`func (o *DeviceSessionOut) SetAutoClosedFromShiftStatus(v string)`

SetAutoClosedFromShiftStatus sets AutoClosedFromShiftStatus field to given value.

### HasAutoClosedFromShiftStatus

`func (o *DeviceSessionOut) HasAutoClosedFromShiftStatus() bool`

HasAutoClosedFromShiftStatus returns a boolean if a field has been set.

### SetAutoClosedFromShiftStatusNil

`func (o *DeviceSessionOut) SetAutoClosedFromShiftStatusNil(b bool)`

 SetAutoClosedFromShiftStatusNil sets the value for AutoClosedFromShiftStatus to be an explicit nil

### UnsetAutoClosedFromShiftStatus
`func (o *DeviceSessionOut) UnsetAutoClosedFromShiftStatus()`

UnsetAutoClosedFromShiftStatus ensures that no value is present for AutoClosedFromShiftStatus, not even an explicit nil
### GetPhotoCount

`func (o *DeviceSessionOut) GetPhotoCount() int32`

GetPhotoCount returns the PhotoCount field if non-nil, zero value otherwise.

### GetPhotoCountOk

`func (o *DeviceSessionOut) GetPhotoCountOk() (*int32, bool)`

GetPhotoCountOk returns a tuple with the PhotoCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPhotoCount

`func (o *DeviceSessionOut) SetPhotoCount(v int32)`

SetPhotoCount sets PhotoCount field to given value.


### SetPhotoCountNil

`func (o *DeviceSessionOut) SetPhotoCountNil(b bool)`

 SetPhotoCountNil sets the value for PhotoCount to be an explicit nil

### UnsetPhotoCount
`func (o *DeviceSessionOut) UnsetPhotoCount()`

UnsetPhotoCount ensures that no value is present for PhotoCount, not even an explicit nil
### GetNoteCount

`func (o *DeviceSessionOut) GetNoteCount() int32`

GetNoteCount returns the NoteCount field if non-nil, zero value otherwise.

### GetNoteCountOk

`func (o *DeviceSessionOut) GetNoteCountOk() (*int32, bool)`

GetNoteCountOk returns a tuple with the NoteCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNoteCount

`func (o *DeviceSessionOut) SetNoteCount(v int32)`

SetNoteCount sets NoteCount field to given value.


### SetNoteCountNil

`func (o *DeviceSessionOut) SetNoteCountNil(b bool)`

 SetNoteCountNil sets the value for NoteCount to be an explicit nil

### UnsetNoteCount
`func (o *DeviceSessionOut) UnsetNoteCount()`

UnsetNoteCount ensures that no value is present for NoteCount, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


