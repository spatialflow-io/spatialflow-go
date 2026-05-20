# DeviceStatsOut

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**TotalDevices** | **int32** |  | 
**ActiveDevices** | **int32** |  | 
**InactiveDevices** | **int32** |  | 
**EventsToday** | **int32** |  | 
**EventsYesterday** | **int32** |  | 
**EventsThisWeek** | **int32** |  | 
**EventsLastWeek** | **int32** |  | 
**EntriesToday** | **int32** |  | 
**ExitsToday** | **int32** |  | 
**ActiveGeofences** | **int32** |  | 

## Methods

### NewDeviceStatsOut

`func NewDeviceStatsOut(totalDevices int32, activeDevices int32, inactiveDevices int32, eventsToday int32, eventsYesterday int32, eventsThisWeek int32, eventsLastWeek int32, entriesToday int32, exitsToday int32, activeGeofences int32, ) *DeviceStatsOut`

NewDeviceStatsOut instantiates a new DeviceStatsOut object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDeviceStatsOutWithDefaults

`func NewDeviceStatsOutWithDefaults() *DeviceStatsOut`

NewDeviceStatsOutWithDefaults instantiates a new DeviceStatsOut object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTotalDevices

`func (o *DeviceStatsOut) GetTotalDevices() int32`

GetTotalDevices returns the TotalDevices field if non-nil, zero value otherwise.

### GetTotalDevicesOk

`func (o *DeviceStatsOut) GetTotalDevicesOk() (*int32, bool)`

GetTotalDevicesOk returns a tuple with the TotalDevices field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalDevices

`func (o *DeviceStatsOut) SetTotalDevices(v int32)`

SetTotalDevices sets TotalDevices field to given value.


### GetActiveDevices

`func (o *DeviceStatsOut) GetActiveDevices() int32`

GetActiveDevices returns the ActiveDevices field if non-nil, zero value otherwise.

### GetActiveDevicesOk

`func (o *DeviceStatsOut) GetActiveDevicesOk() (*int32, bool)`

GetActiveDevicesOk returns a tuple with the ActiveDevices field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActiveDevices

`func (o *DeviceStatsOut) SetActiveDevices(v int32)`

SetActiveDevices sets ActiveDevices field to given value.


### GetInactiveDevices

`func (o *DeviceStatsOut) GetInactiveDevices() int32`

GetInactiveDevices returns the InactiveDevices field if non-nil, zero value otherwise.

### GetInactiveDevicesOk

`func (o *DeviceStatsOut) GetInactiveDevicesOk() (*int32, bool)`

GetInactiveDevicesOk returns a tuple with the InactiveDevices field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInactiveDevices

`func (o *DeviceStatsOut) SetInactiveDevices(v int32)`

SetInactiveDevices sets InactiveDevices field to given value.


### GetEventsToday

`func (o *DeviceStatsOut) GetEventsToday() int32`

GetEventsToday returns the EventsToday field if non-nil, zero value otherwise.

### GetEventsTodayOk

`func (o *DeviceStatsOut) GetEventsTodayOk() (*int32, bool)`

GetEventsTodayOk returns a tuple with the EventsToday field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEventsToday

`func (o *DeviceStatsOut) SetEventsToday(v int32)`

SetEventsToday sets EventsToday field to given value.


### GetEventsYesterday

`func (o *DeviceStatsOut) GetEventsYesterday() int32`

GetEventsYesterday returns the EventsYesterday field if non-nil, zero value otherwise.

### GetEventsYesterdayOk

`func (o *DeviceStatsOut) GetEventsYesterdayOk() (*int32, bool)`

GetEventsYesterdayOk returns a tuple with the EventsYesterday field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEventsYesterday

`func (o *DeviceStatsOut) SetEventsYesterday(v int32)`

SetEventsYesterday sets EventsYesterday field to given value.


### GetEventsThisWeek

`func (o *DeviceStatsOut) GetEventsThisWeek() int32`

GetEventsThisWeek returns the EventsThisWeek field if non-nil, zero value otherwise.

### GetEventsThisWeekOk

`func (o *DeviceStatsOut) GetEventsThisWeekOk() (*int32, bool)`

GetEventsThisWeekOk returns a tuple with the EventsThisWeek field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEventsThisWeek

`func (o *DeviceStatsOut) SetEventsThisWeek(v int32)`

SetEventsThisWeek sets EventsThisWeek field to given value.


### GetEventsLastWeek

`func (o *DeviceStatsOut) GetEventsLastWeek() int32`

GetEventsLastWeek returns the EventsLastWeek field if non-nil, zero value otherwise.

### GetEventsLastWeekOk

`func (o *DeviceStatsOut) GetEventsLastWeekOk() (*int32, bool)`

GetEventsLastWeekOk returns a tuple with the EventsLastWeek field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEventsLastWeek

`func (o *DeviceStatsOut) SetEventsLastWeek(v int32)`

SetEventsLastWeek sets EventsLastWeek field to given value.


### GetEntriesToday

`func (o *DeviceStatsOut) GetEntriesToday() int32`

GetEntriesToday returns the EntriesToday field if non-nil, zero value otherwise.

### GetEntriesTodayOk

`func (o *DeviceStatsOut) GetEntriesTodayOk() (*int32, bool)`

GetEntriesTodayOk returns a tuple with the EntriesToday field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEntriesToday

`func (o *DeviceStatsOut) SetEntriesToday(v int32)`

SetEntriesToday sets EntriesToday field to given value.


### GetExitsToday

`func (o *DeviceStatsOut) GetExitsToday() int32`

GetExitsToday returns the ExitsToday field if non-nil, zero value otherwise.

### GetExitsTodayOk

`func (o *DeviceStatsOut) GetExitsTodayOk() (*int32, bool)`

GetExitsTodayOk returns a tuple with the ExitsToday field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExitsToday

`func (o *DeviceStatsOut) SetExitsToday(v int32)`

SetExitsToday sets ExitsToday field to given value.


### GetActiveGeofences

`func (o *DeviceStatsOut) GetActiveGeofences() int32`

GetActiveGeofences returns the ActiveGeofences field if non-nil, zero value otherwise.

### GetActiveGeofencesOk

`func (o *DeviceStatsOut) GetActiveGeofencesOk() (*int32, bool)`

GetActiveGeofencesOk returns a tuple with the ActiveGeofences field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActiveGeofences

`func (o *DeviceStatsOut) SetActiveGeofences(v int32)`

SetActiveGeofences sets ActiveGeofences field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


