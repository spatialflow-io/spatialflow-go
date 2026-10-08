# DeviceOut

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** |  | 
**DeviceId** | **string** |  | 
**Name** | **string** |  | 
**DeviceType** | **string** |  | 
**RegistrationSource** | **string** |  | 
**IsActive** | **bool** |  | 
**IsExample** | Pointer to **bool** |  | [optional] [default to false]
**ShiftStatus** | **string** |  | 
**ShiftStartedAt** | Pointer to **NullableTime** |  | [optional] 
**ShiftPausedAt** | Pointer to **NullableTime** |  | [optional] 
**ShiftEndedAt** | Pointer to **NullableTime** |  | [optional] 
**ShiftResumedAt** | Pointer to **NullableTime** |  | [optional] 
**LastLocation** | Pointer to [**NullableLatLonOut**](LatLonOut.md) |  | [optional] 
**LastLocationTime** | Pointer to **NullableTime** |  | [optional] 
**LastAccuracy** | Pointer to **NullableFloat32** |  | [optional] 
**LastHeading** | Pointer to **NullableFloat32** |  | [optional] 
**LastBatteryLevel** | Pointer to **NullableInt32** |  | [optional] 
**LastBatteryCharging** | Pointer to **NullableBool** |  | [optional] 
**LastBatteryTime** | Pointer to **NullableTime** |  | [optional] 
**DriverName** | Pointer to **string** |  | [optional] [default to ""]
**VehicleLabel** | Pointer to **string** |  | [optional] [default to ""]
**EmployeeId** | Pointer to **string** |  | [optional] [default to ""]
**LicensePlate** | Pointer to **string** |  | [optional] [default to ""]
**Group** | Pointer to **string** |  | [optional] [default to ""]
**Status** | **string** |  | 
**LastSeenSeconds** | Pointer to **NullableInt32** |  | [optional] 
**Parked** | Pointer to **bool** |  | [optional] [default to false]
**TodayDistanceMeters** | Pointer to **NullableFloat32** |  | [optional] 
**CurrentSessionNotes** | Pointer to **string** |  | [optional] [default to ""]
**CurrentSession** | Pointer to [**NullableCurrentSessionOut**](CurrentSessionOut.md) |  | [optional] 
**InGeofenceIds** | Pointer to **[]string** |  | [optional] [default to []]
**InGeofenceEntries** | Pointer to **map[string]string** |  | [optional] [default to {}]
**CreatedAt** | **time.Time** |  | 
**UpdatedAt** | **time.Time** |  | 

## Methods

### NewDeviceOut

`func NewDeviceOut(id string, deviceId string, name string, deviceType string, registrationSource string, isActive bool, shiftStatus string, status string, createdAt time.Time, updatedAt time.Time, ) *DeviceOut`

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


### GetRegistrationSource

`func (o *DeviceOut) GetRegistrationSource() string`

GetRegistrationSource returns the RegistrationSource field if non-nil, zero value otherwise.

### GetRegistrationSourceOk

`func (o *DeviceOut) GetRegistrationSourceOk() (*string, bool)`

GetRegistrationSourceOk returns a tuple with the RegistrationSource field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegistrationSource

`func (o *DeviceOut) SetRegistrationSource(v string)`

SetRegistrationSource sets RegistrationSource field to given value.


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


### GetIsExample

`func (o *DeviceOut) GetIsExample() bool`

GetIsExample returns the IsExample field if non-nil, zero value otherwise.

### GetIsExampleOk

`func (o *DeviceOut) GetIsExampleOk() (*bool, bool)`

GetIsExampleOk returns a tuple with the IsExample field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsExample

`func (o *DeviceOut) SetIsExample(v bool)`

SetIsExample sets IsExample field to given value.

### HasIsExample

`func (o *DeviceOut) HasIsExample() bool`

HasIsExample returns a boolean if a field has been set.

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

`func (o *DeviceOut) GetLastLocation() LatLonOut`

GetLastLocation returns the LastLocation field if non-nil, zero value otherwise.

### GetLastLocationOk

`func (o *DeviceOut) GetLastLocationOk() (*LatLonOut, bool)`

GetLastLocationOk returns a tuple with the LastLocation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastLocation

`func (o *DeviceOut) SetLastLocation(v LatLonOut)`

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
### GetLastAccuracy

`func (o *DeviceOut) GetLastAccuracy() float32`

GetLastAccuracy returns the LastAccuracy field if non-nil, zero value otherwise.

### GetLastAccuracyOk

`func (o *DeviceOut) GetLastAccuracyOk() (*float32, bool)`

GetLastAccuracyOk returns a tuple with the LastAccuracy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastAccuracy

`func (o *DeviceOut) SetLastAccuracy(v float32)`

SetLastAccuracy sets LastAccuracy field to given value.

### HasLastAccuracy

`func (o *DeviceOut) HasLastAccuracy() bool`

HasLastAccuracy returns a boolean if a field has been set.

### SetLastAccuracyNil

`func (o *DeviceOut) SetLastAccuracyNil(b bool)`

 SetLastAccuracyNil sets the value for LastAccuracy to be an explicit nil

### UnsetLastAccuracy
`func (o *DeviceOut) UnsetLastAccuracy()`

UnsetLastAccuracy ensures that no value is present for LastAccuracy, not even an explicit nil
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
### GetLastBatteryLevel

`func (o *DeviceOut) GetLastBatteryLevel() int32`

GetLastBatteryLevel returns the LastBatteryLevel field if non-nil, zero value otherwise.

### GetLastBatteryLevelOk

`func (o *DeviceOut) GetLastBatteryLevelOk() (*int32, bool)`

GetLastBatteryLevelOk returns a tuple with the LastBatteryLevel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastBatteryLevel

`func (o *DeviceOut) SetLastBatteryLevel(v int32)`

SetLastBatteryLevel sets LastBatteryLevel field to given value.

### HasLastBatteryLevel

`func (o *DeviceOut) HasLastBatteryLevel() bool`

HasLastBatteryLevel returns a boolean if a field has been set.

### SetLastBatteryLevelNil

`func (o *DeviceOut) SetLastBatteryLevelNil(b bool)`

 SetLastBatteryLevelNil sets the value for LastBatteryLevel to be an explicit nil

### UnsetLastBatteryLevel
`func (o *DeviceOut) UnsetLastBatteryLevel()`

UnsetLastBatteryLevel ensures that no value is present for LastBatteryLevel, not even an explicit nil
### GetLastBatteryCharging

`func (o *DeviceOut) GetLastBatteryCharging() bool`

GetLastBatteryCharging returns the LastBatteryCharging field if non-nil, zero value otherwise.

### GetLastBatteryChargingOk

`func (o *DeviceOut) GetLastBatteryChargingOk() (*bool, bool)`

GetLastBatteryChargingOk returns a tuple with the LastBatteryCharging field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastBatteryCharging

`func (o *DeviceOut) SetLastBatteryCharging(v bool)`

SetLastBatteryCharging sets LastBatteryCharging field to given value.

### HasLastBatteryCharging

`func (o *DeviceOut) HasLastBatteryCharging() bool`

HasLastBatteryCharging returns a boolean if a field has been set.

### SetLastBatteryChargingNil

`func (o *DeviceOut) SetLastBatteryChargingNil(b bool)`

 SetLastBatteryChargingNil sets the value for LastBatteryCharging to be an explicit nil

### UnsetLastBatteryCharging
`func (o *DeviceOut) UnsetLastBatteryCharging()`

UnsetLastBatteryCharging ensures that no value is present for LastBatteryCharging, not even an explicit nil
### GetLastBatteryTime

`func (o *DeviceOut) GetLastBatteryTime() time.Time`

GetLastBatteryTime returns the LastBatteryTime field if non-nil, zero value otherwise.

### GetLastBatteryTimeOk

`func (o *DeviceOut) GetLastBatteryTimeOk() (*time.Time, bool)`

GetLastBatteryTimeOk returns a tuple with the LastBatteryTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastBatteryTime

`func (o *DeviceOut) SetLastBatteryTime(v time.Time)`

SetLastBatteryTime sets LastBatteryTime field to given value.

### HasLastBatteryTime

`func (o *DeviceOut) HasLastBatteryTime() bool`

HasLastBatteryTime returns a boolean if a field has been set.

### SetLastBatteryTimeNil

`func (o *DeviceOut) SetLastBatteryTimeNil(b bool)`

 SetLastBatteryTimeNil sets the value for LastBatteryTime to be an explicit nil

### UnsetLastBatteryTime
`func (o *DeviceOut) UnsetLastBatteryTime()`

UnsetLastBatteryTime ensures that no value is present for LastBatteryTime, not even an explicit nil
### GetDriverName

`func (o *DeviceOut) GetDriverName() string`

GetDriverName returns the DriverName field if non-nil, zero value otherwise.

### GetDriverNameOk

`func (o *DeviceOut) GetDriverNameOk() (*string, bool)`

GetDriverNameOk returns a tuple with the DriverName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDriverName

`func (o *DeviceOut) SetDriverName(v string)`

SetDriverName sets DriverName field to given value.

### HasDriverName

`func (o *DeviceOut) HasDriverName() bool`

HasDriverName returns a boolean if a field has been set.

### GetVehicleLabel

`func (o *DeviceOut) GetVehicleLabel() string`

GetVehicleLabel returns the VehicleLabel field if non-nil, zero value otherwise.

### GetVehicleLabelOk

`func (o *DeviceOut) GetVehicleLabelOk() (*string, bool)`

GetVehicleLabelOk returns a tuple with the VehicleLabel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVehicleLabel

`func (o *DeviceOut) SetVehicleLabel(v string)`

SetVehicleLabel sets VehicleLabel field to given value.

### HasVehicleLabel

`func (o *DeviceOut) HasVehicleLabel() bool`

HasVehicleLabel returns a boolean if a field has been set.

### GetEmployeeId

`func (o *DeviceOut) GetEmployeeId() string`

GetEmployeeId returns the EmployeeId field if non-nil, zero value otherwise.

### GetEmployeeIdOk

`func (o *DeviceOut) GetEmployeeIdOk() (*string, bool)`

GetEmployeeIdOk returns a tuple with the EmployeeId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmployeeId

`func (o *DeviceOut) SetEmployeeId(v string)`

SetEmployeeId sets EmployeeId field to given value.

### HasEmployeeId

`func (o *DeviceOut) HasEmployeeId() bool`

HasEmployeeId returns a boolean if a field has been set.

### GetLicensePlate

`func (o *DeviceOut) GetLicensePlate() string`

GetLicensePlate returns the LicensePlate field if non-nil, zero value otherwise.

### GetLicensePlateOk

`func (o *DeviceOut) GetLicensePlateOk() (*string, bool)`

GetLicensePlateOk returns a tuple with the LicensePlate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLicensePlate

`func (o *DeviceOut) SetLicensePlate(v string)`

SetLicensePlate sets LicensePlate field to given value.

### HasLicensePlate

`func (o *DeviceOut) HasLicensePlate() bool`

HasLicensePlate returns a boolean if a field has been set.

### GetGroup

`func (o *DeviceOut) GetGroup() string`

GetGroup returns the Group field if non-nil, zero value otherwise.

### GetGroupOk

`func (o *DeviceOut) GetGroupOk() (*string, bool)`

GetGroupOk returns a tuple with the Group field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroup

`func (o *DeviceOut) SetGroup(v string)`

SetGroup sets Group field to given value.

### HasGroup

`func (o *DeviceOut) HasGroup() bool`

HasGroup returns a boolean if a field has been set.

### GetStatus

`func (o *DeviceOut) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *DeviceOut) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *DeviceOut) SetStatus(v string)`

SetStatus sets Status field to given value.


### GetLastSeenSeconds

`func (o *DeviceOut) GetLastSeenSeconds() int32`

GetLastSeenSeconds returns the LastSeenSeconds field if non-nil, zero value otherwise.

### GetLastSeenSecondsOk

`func (o *DeviceOut) GetLastSeenSecondsOk() (*int32, bool)`

GetLastSeenSecondsOk returns a tuple with the LastSeenSeconds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastSeenSeconds

`func (o *DeviceOut) SetLastSeenSeconds(v int32)`

SetLastSeenSeconds sets LastSeenSeconds field to given value.

### HasLastSeenSeconds

`func (o *DeviceOut) HasLastSeenSeconds() bool`

HasLastSeenSeconds returns a boolean if a field has been set.

### SetLastSeenSecondsNil

`func (o *DeviceOut) SetLastSeenSecondsNil(b bool)`

 SetLastSeenSecondsNil sets the value for LastSeenSeconds to be an explicit nil

### UnsetLastSeenSeconds
`func (o *DeviceOut) UnsetLastSeenSeconds()`

UnsetLastSeenSeconds ensures that no value is present for LastSeenSeconds, not even an explicit nil
### GetParked

`func (o *DeviceOut) GetParked() bool`

GetParked returns the Parked field if non-nil, zero value otherwise.

### GetParkedOk

`func (o *DeviceOut) GetParkedOk() (*bool, bool)`

GetParkedOk returns a tuple with the Parked field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParked

`func (o *DeviceOut) SetParked(v bool)`

SetParked sets Parked field to given value.

### HasParked

`func (o *DeviceOut) HasParked() bool`

HasParked returns a boolean if a field has been set.

### GetTodayDistanceMeters

`func (o *DeviceOut) GetTodayDistanceMeters() float32`

GetTodayDistanceMeters returns the TodayDistanceMeters field if non-nil, zero value otherwise.

### GetTodayDistanceMetersOk

`func (o *DeviceOut) GetTodayDistanceMetersOk() (*float32, bool)`

GetTodayDistanceMetersOk returns a tuple with the TodayDistanceMeters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTodayDistanceMeters

`func (o *DeviceOut) SetTodayDistanceMeters(v float32)`

SetTodayDistanceMeters sets TodayDistanceMeters field to given value.

### HasTodayDistanceMeters

`func (o *DeviceOut) HasTodayDistanceMeters() bool`

HasTodayDistanceMeters returns a boolean if a field has been set.

### SetTodayDistanceMetersNil

`func (o *DeviceOut) SetTodayDistanceMetersNil(b bool)`

 SetTodayDistanceMetersNil sets the value for TodayDistanceMeters to be an explicit nil

### UnsetTodayDistanceMeters
`func (o *DeviceOut) UnsetTodayDistanceMeters()`

UnsetTodayDistanceMeters ensures that no value is present for TodayDistanceMeters, not even an explicit nil
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

### GetCurrentSession

`func (o *DeviceOut) GetCurrentSession() CurrentSessionOut`

GetCurrentSession returns the CurrentSession field if non-nil, zero value otherwise.

### GetCurrentSessionOk

`func (o *DeviceOut) GetCurrentSessionOk() (*CurrentSessionOut, bool)`

GetCurrentSessionOk returns a tuple with the CurrentSession field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrentSession

`func (o *DeviceOut) SetCurrentSession(v CurrentSessionOut)`

SetCurrentSession sets CurrentSession field to given value.

### HasCurrentSession

`func (o *DeviceOut) HasCurrentSession() bool`

HasCurrentSession returns a boolean if a field has been set.

### SetCurrentSessionNil

`func (o *DeviceOut) SetCurrentSessionNil(b bool)`

 SetCurrentSessionNil sets the value for CurrentSession to be an explicit nil

### UnsetCurrentSession
`func (o *DeviceOut) UnsetCurrentSession()`

UnsetCurrentSession ensures that no value is present for CurrentSession, not even an explicit nil
### GetInGeofenceIds

`func (o *DeviceOut) GetInGeofenceIds() []string`

GetInGeofenceIds returns the InGeofenceIds field if non-nil, zero value otherwise.

### GetInGeofenceIdsOk

`func (o *DeviceOut) GetInGeofenceIdsOk() (*[]string, bool)`

GetInGeofenceIdsOk returns a tuple with the InGeofenceIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInGeofenceIds

`func (o *DeviceOut) SetInGeofenceIds(v []string)`

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


