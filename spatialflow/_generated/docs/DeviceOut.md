# DeviceOut

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** |  | 
**DeviceId** | **string** |  | 
**Name** | **string** |  | 
**DeviceType** | **string** |  | 
**IsActive** | **bool** |  | 
**ShiftStatus** | **string** |  | 
**ShiftStartedAt** | Pointer to **NullableTime** |  | [optional] 
**ShiftPausedAt** | Pointer to **NullableTime** |  | [optional] 
**ShiftEndedAt** | Pointer to **NullableTime** |  | [optional] 
**ShiftResumedAt** | Pointer to **NullableTime** |  | [optional] 
**LastLocation** | Pointer to **map[string]interface{}** |  | [optional] 
**LastLocationTime** | Pointer to **NullableTime** |  | [optional] 
**LastHeading** | Pointer to **NullableFloat32** |  | [optional] 
**CurrentSessionNotes** | Pointer to **string** |  | [optional] [default to ""]
**InGeofenceIds** | Pointer to **[]string** |  | [optional] [default to []]
**InGeofenceEntries** | Pointer to **map[string]string** |  | [optional] [default to {}]
**CreatedAt** | **time.Time** |  | 
**UpdatedAt** | **time.Time** |  | 

## Methods

### NewDeviceOut

`func NewDeviceOut(id string, deviceId string, name string, deviceType string, isActive bool, shiftStatus string, createdAt time.Time, updatedAt time.Time, ) *DeviceOut`

NewDeviceOut instantiates a new DeviceOut object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDeviceOutWithDefaults

`func NewDeviceOutWithDefaults() *DeviceOut`

NewDeviceOutWithDefaults instantiates a new DeviceOut object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *DeviceOut) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *DeviceOut) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *DeviceOut) SetId(v string)`

SetId sets Id field to given value.


### GetDeviceId

`func (o *DeviceOut) GetDeviceId() string`

GetDeviceId returns the DeviceId field if non-nil, zero value otherwise.

### GetDeviceIdOk

`func (o *DeviceOut) GetDeviceIdOk() (*string, bool)`

GetDeviceIdOk returns a tuple with the DeviceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeviceId

`func (o *DeviceOut) SetDeviceId(v string)`

SetDeviceId sets DeviceId field to given value.


### GetName

`func (o *DeviceOut) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *DeviceOut) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *DeviceOut) SetName(v string)`

SetName sets Name field to given value.


### GetDeviceType

`func (o *DeviceOut) GetDeviceType() string`

GetDeviceType returns the DeviceType field if non-nil, zero value otherwise.

### GetDeviceTypeOk

`func (o *DeviceOut) GetDeviceTypeOk() (*string, bool)`

GetDeviceTypeOk returns a tuple with the DeviceType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeviceType

`func (o *DeviceOut) SetDeviceType(v string)`

SetDeviceType sets DeviceType field to given value.


### GetIsActive

`func (o *DeviceOut) GetIsActive() bool`

GetIsActive returns the IsActive field if non-nil, zero value otherwise.

### GetIsActiveOk

`func (o *DeviceOut) GetIsActiveOk() (*bool, bool)`

GetIsActiveOk returns a tuple with the IsActive field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsActive

`func (o *DeviceOut) SetIsActive(v bool)`

SetIsActive sets IsActive field to given value.


### GetShiftStatus

`func (o *DeviceOut) GetShiftStatus() string`

GetShiftStatus returns the ShiftStatus field if non-nil, zero value otherwise.

### GetShiftStatusOk

`func (o *DeviceOut) GetShiftStatusOk() (*string, bool)`

GetShiftStatusOk returns a tuple with the ShiftStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShiftStatus

`func (o *DeviceOut) SetShiftStatus(v string)`

SetShiftStatus sets ShiftStatus field to given value.


### GetShiftStartedAt

`func (o *DeviceOut) GetShiftStartedAt() time.Time`

GetShiftStartedAt returns the ShiftStartedAt field if non-nil, zero value otherwise.

### GetShiftStartedAtOk

`func (o *DeviceOut) GetShiftStartedAtOk() (*time.Time, bool)`

GetShiftStartedAtOk returns a tuple with the ShiftStartedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShiftStartedAt

`func (o *DeviceOut) SetShiftStartedAt(v time.Time)`

SetShiftStartedAt sets ShiftStartedAt field to given value.

### HasShiftStartedAt

`func (o *DeviceOut) HasShiftStartedAt() bool`

HasShiftStartedAt returns a boolean if a field has been set.

### SetShiftStartedAtNil

`func (o *DeviceOut) SetShiftStartedAtNil(b bool)`

 SetShiftStartedAtNil sets the value for ShiftStartedAt to be an explicit nil

### UnsetShiftStartedAt
`func (o *DeviceOut) UnsetShiftStartedAt()`

UnsetShiftStartedAt ensures that no value is present for ShiftStartedAt, not even an explicit nil
### GetShiftPausedAt

`func (o *DeviceOut) GetShiftPausedAt() time.Time`

GetShiftPausedAt returns the ShiftPausedAt field if non-nil, zero value otherwise.

### GetShiftPausedAtOk

`func (o *DeviceOut) GetShiftPausedAtOk() (*time.Time, bool)`

GetShiftPausedAtOk returns a tuple with the ShiftPausedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShiftPausedAt

`func (o *DeviceOut) SetShiftPausedAt(v time.Time)`

SetShiftPausedAt sets ShiftPausedAt field to given value.

### HasShiftPausedAt

`func (o *DeviceOut) HasShiftPausedAt() bool`

HasShiftPausedAt returns a boolean if a field has been set.

### SetShiftPausedAtNil

`func (o *DeviceOut) SetShiftPausedAtNil(b bool)`

 SetShiftPausedAtNil sets the value for ShiftPausedAt to be an explicit nil

### UnsetShiftPausedAt
`func (o *DeviceOut) UnsetShiftPausedAt()`

UnsetShiftPausedAt ensures that no value is present for ShiftPausedAt, not even an explicit nil
### GetShiftEndedAt

`func (o *DeviceOut) GetShiftEndedAt() time.Time`

GetShiftEndedAt returns the ShiftEndedAt field if non-nil, zero value otherwise.

### GetShiftEndedAtOk

`func (o *DeviceOut) GetShiftEndedAtOk() (*time.Time, bool)`

GetShiftEndedAtOk returns a tuple with the ShiftEndedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShiftEndedAt

`func (o *DeviceOut) SetShiftEndedAt(v time.Time)`

SetShiftEndedAt sets ShiftEndedAt field to given value.

### HasShiftEndedAt

`func (o *DeviceOut) HasShiftEndedAt() bool`

HasShiftEndedAt returns a boolean if a field has been set.

### SetShiftEndedAtNil

`func (o *DeviceOut) SetShiftEndedAtNil(b bool)`

 SetShiftEndedAtNil sets the value for ShiftEndedAt to be an explicit nil

### UnsetShiftEndedAt
`func (o *DeviceOut) UnsetShiftEndedAt()`

UnsetShiftEndedAt ensures that no value is present for ShiftEndedAt, not even an explicit nil
### GetShiftResumedAt

`func (o *DeviceOut) GetShiftResumedAt() time.Time`

GetShiftResumedAt returns the ShiftResumedAt field if non-nil, zero value otherwise.

### GetShiftResumedAtOk

`func (o *DeviceOut) GetShiftResumedAtOk() (*time.Time, bool)`

GetShiftResumedAtOk returns a tuple with the ShiftResumedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShiftResumedAt

`func (o *DeviceOut) SetShiftResumedAt(v time.Time)`

SetShiftResumedAt sets ShiftResumedAt field to given value.

### HasShiftResumedAt

`func (o *DeviceOut) HasShiftResumedAt() bool`

HasShiftResumedAt returns a boolean if a field has been set.

### SetShiftResumedAtNil

`func (o *DeviceOut) SetShiftResumedAtNil(b bool)`

 SetShiftResumedAtNil sets the value for ShiftResumedAt to be an explicit nil

### UnsetShiftResumedAt
`func (o *DeviceOut) UnsetShiftResumedAt()`

UnsetShiftResumedAt ensures that no value is present for ShiftResumedAt, not even an explicit nil
### GetLastLocation

`func (o *DeviceOut) GetLastLocation() map[string]interface{}`

GetLastLocation returns the LastLocation field if non-nil, zero value otherwise.

### GetLastLocationOk

`func (o *DeviceOut) GetLastLocationOk() (*map[string]interface{}, bool)`

GetLastLocationOk returns a tuple with the LastLocation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastLocation

`func (o *DeviceOut) SetLastLocation(v map[string]interface{})`

SetLastLocation sets LastLocation field to given value.

### HasLastLocation

`func (o *DeviceOut) HasLastLocation() bool`

HasLastLocation returns a boolean if a field has been set.

### SetLastLocationNil

`func (o *DeviceOut) SetLastLocationNil(b bool)`

 SetLastLocationNil sets the value for LastLocation to be an explicit nil

### UnsetLastLocation
`func (o *DeviceOut) UnsetLastLocation()`

UnsetLastLocation ensures that no value is present for LastLocation, not even an explicit nil
### GetLastLocationTime

`func (o *DeviceOut) GetLastLocationTime() time.Time`

GetLastLocationTime returns the LastLocationTime field if non-nil, zero value otherwise.

### GetLastLocationTimeOk

`func (o *DeviceOut) GetLastLocationTimeOk() (*time.Time, bool)`

GetLastLocationTimeOk returns a tuple with the LastLocationTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastLocationTime

`func (o *DeviceOut) SetLastLocationTime(v time.Time)`

SetLastLocationTime sets LastLocationTime field to given value.

### HasLastLocationTime

`func (o *DeviceOut) HasLastLocationTime() bool`

HasLastLocationTime returns a boolean if a field has been set.

### SetLastLocationTimeNil

`func (o *DeviceOut) SetLastLocationTimeNil(b bool)`

 SetLastLocationTimeNil sets the value for LastLocationTime to be an explicit nil

### UnsetLastLocationTime
`func (o *DeviceOut) UnsetLastLocationTime()`

UnsetLastLocationTime ensures that no value is present for LastLocationTime, not even an explicit nil
### GetLastHeading

`func (o *DeviceOut) GetLastHeading() float32`

GetLastHeading returns the LastHeading field if non-nil, zero value otherwise.

### GetLastHeadingOk

`func (o *DeviceOut) GetLastHeadingOk() (*float32, bool)`

GetLastHeadingOk returns a tuple with the LastHeading field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastHeading

`func (o *DeviceOut) SetLastHeading(v float32)`

SetLastHeading sets LastHeading field to given value.

### HasLastHeading

`func (o *DeviceOut) HasLastHeading() bool`

HasLastHeading returns a boolean if a field has been set.

### SetLastHeadingNil

`func (o *DeviceOut) SetLastHeadingNil(b bool)`

 SetLastHeadingNil sets the value for LastHeading to be an explicit nil

### UnsetLastHeading
`func (o *DeviceOut) UnsetLastHeading()`

UnsetLastHeading ensures that no value is present for LastHeading, not even an explicit nil
### GetCurrentSessionNotes

`func (o *DeviceOut) GetCurrentSessionNotes() string`

GetCurrentSessionNotes returns the CurrentSessionNotes field if non-nil, zero value otherwise.

### GetCurrentSessionNotesOk

`func (o *DeviceOut) GetCurrentSessionNotesOk() (*string, bool)`

GetCurrentSessionNotesOk returns a tuple with the CurrentSessionNotes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrentSessionNotes

`func (o *DeviceOut) SetCurrentSessionNotes(v string)`

SetCurrentSessionNotes sets CurrentSessionNotes field to given value.

### HasCurrentSessionNotes

`func (o *DeviceOut) HasCurrentSessionNotes() bool`

HasCurrentSessionNotes returns a boolean if a field has been set.

### GetInGeofenceIds

`func (o *DeviceOut) GetInGeofenceIds() []*string`

GetInGeofenceIds returns the InGeofenceIds field if non-nil, zero value otherwise.

### GetInGeofenceIdsOk

`func (o *DeviceOut) GetInGeofenceIdsOk() (*[]*string, bool)`

GetInGeofenceIdsOk returns a tuple with the InGeofenceIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInGeofenceIds

`func (o *DeviceOut) SetInGeofenceIds(v []*string)`

SetInGeofenceIds sets InGeofenceIds field to given value.

### HasInGeofenceIds

`func (o *DeviceOut) HasInGeofenceIds() bool`

HasInGeofenceIds returns a boolean if a field has been set.

### GetInGeofenceEntries

`func (o *DeviceOut) GetInGeofenceEntries() map[string]string`

GetInGeofenceEntries returns the InGeofenceEntries field if non-nil, zero value otherwise.

### GetInGeofenceEntriesOk

`func (o *DeviceOut) GetInGeofenceEntriesOk() (*map[string]string, bool)`

GetInGeofenceEntriesOk returns a tuple with the InGeofenceEntries field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInGeofenceEntries

`func (o *DeviceOut) SetInGeofenceEntries(v map[string]string)`

SetInGeofenceEntries sets InGeofenceEntries field to given value.

### HasInGeofenceEntries

`func (o *DeviceOut) HasInGeofenceEntries() bool`

HasInGeofenceEntries returns a boolean if a field has been set.

### GetCreatedAt

`func (o *DeviceOut) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *DeviceOut) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *DeviceOut) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.


### GetUpdatedAt

`func (o *DeviceOut) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *DeviceOut) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *DeviceOut) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


