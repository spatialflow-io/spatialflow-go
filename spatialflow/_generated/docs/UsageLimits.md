# UsageLimits

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**LocationEvents** | Pointer to **int32** |  | [optional] [default to 0]
**ActionDeliveries** | Pointer to **int32** |  | [optional] [default to 0]
**EventUnits** | Pointer to **float32** |  | [optional] [default to 0.0]
**MaxGeofences** | Pointer to **NullableInt32** |  | [optional] 
**MaxDevices** | Pointer to **NullableInt32** |  | [optional] 
**MaxEventsPerMonth** | Pointer to **NullableInt32** |  | [optional] 
**GeofenceCount** | Pointer to **int32** |  | [optional] [default to 0]
**DeviceCount** | Pointer to **int32** |  | [optional] [default to 0]

## Methods

### NewUsageLimits

`func NewUsageLimits() *UsageLimits`

NewUsageLimits instantiates a new UsageLimits object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUsageLimitsWithDefaults

`func NewUsageLimitsWithDefaults() *UsageLimits`

NewUsageLimitsWithDefaults instantiates a new UsageLimits object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLocationEvents

`func (o *UsageLimits) GetLocationEvents() int32`

GetLocationEvents returns the LocationEvents field if non-nil, zero value otherwise.

### GetLocationEventsOk

`func (o *UsageLimits) GetLocationEventsOk() (*int32, bool)`

GetLocationEventsOk returns a tuple with the LocationEvents field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocationEvents

`func (o *UsageLimits) SetLocationEvents(v int32)`

SetLocationEvents sets LocationEvents field to given value.

### HasLocationEvents

`func (o *UsageLimits) HasLocationEvents() bool`

HasLocationEvents returns a boolean if a field has been set.

### GetActionDeliveries

`func (o *UsageLimits) GetActionDeliveries() int32`

GetActionDeliveries returns the ActionDeliveries field if non-nil, zero value otherwise.

### GetActionDeliveriesOk

`func (o *UsageLimits) GetActionDeliveriesOk() (*int32, bool)`

GetActionDeliveriesOk returns a tuple with the ActionDeliveries field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActionDeliveries

`func (o *UsageLimits) SetActionDeliveries(v int32)`

SetActionDeliveries sets ActionDeliveries field to given value.

### HasActionDeliveries

`func (o *UsageLimits) HasActionDeliveries() bool`

HasActionDeliveries returns a boolean if a field has been set.

### GetEventUnits

`func (o *UsageLimits) GetEventUnits() float32`

GetEventUnits returns the EventUnits field if non-nil, zero value otherwise.

### GetEventUnitsOk

`func (o *UsageLimits) GetEventUnitsOk() (*float32, bool)`

GetEventUnitsOk returns a tuple with the EventUnits field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEventUnits

`func (o *UsageLimits) SetEventUnits(v float32)`

SetEventUnits sets EventUnits field to given value.

### HasEventUnits

`func (o *UsageLimits) HasEventUnits() bool`

HasEventUnits returns a boolean if a field has been set.

### GetMaxGeofences

`func (o *UsageLimits) GetMaxGeofences() int32`

GetMaxGeofences returns the MaxGeofences field if non-nil, zero value otherwise.

### GetMaxGeofencesOk

`func (o *UsageLimits) GetMaxGeofencesOk() (*int32, bool)`

GetMaxGeofencesOk returns a tuple with the MaxGeofences field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxGeofences

`func (o *UsageLimits) SetMaxGeofences(v int32)`

SetMaxGeofences sets MaxGeofences field to given value.

### HasMaxGeofences

`func (o *UsageLimits) HasMaxGeofences() bool`

HasMaxGeofences returns a boolean if a field has been set.

### SetMaxGeofencesNil

`func (o *UsageLimits) SetMaxGeofencesNil(b bool)`

 SetMaxGeofencesNil sets the value for MaxGeofences to be an explicit nil

### UnsetMaxGeofences
`func (o *UsageLimits) UnsetMaxGeofences()`

UnsetMaxGeofences ensures that no value is present for MaxGeofences, not even an explicit nil
### GetMaxDevices

`func (o *UsageLimits) GetMaxDevices() int32`

GetMaxDevices returns the MaxDevices field if non-nil, zero value otherwise.

### GetMaxDevicesOk

`func (o *UsageLimits) GetMaxDevicesOk() (*int32, bool)`

GetMaxDevicesOk returns a tuple with the MaxDevices field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxDevices

`func (o *UsageLimits) SetMaxDevices(v int32)`

SetMaxDevices sets MaxDevices field to given value.

### HasMaxDevices

`func (o *UsageLimits) HasMaxDevices() bool`

HasMaxDevices returns a boolean if a field has been set.

### SetMaxDevicesNil

`func (o *UsageLimits) SetMaxDevicesNil(b bool)`

 SetMaxDevicesNil sets the value for MaxDevices to be an explicit nil

### UnsetMaxDevices
`func (o *UsageLimits) UnsetMaxDevices()`

UnsetMaxDevices ensures that no value is present for MaxDevices, not even an explicit nil
### GetMaxEventsPerMonth

`func (o *UsageLimits) GetMaxEventsPerMonth() int32`

GetMaxEventsPerMonth returns the MaxEventsPerMonth field if non-nil, zero value otherwise.

### GetMaxEventsPerMonthOk

`func (o *UsageLimits) GetMaxEventsPerMonthOk() (*int32, bool)`

GetMaxEventsPerMonthOk returns a tuple with the MaxEventsPerMonth field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxEventsPerMonth

`func (o *UsageLimits) SetMaxEventsPerMonth(v int32)`

SetMaxEventsPerMonth sets MaxEventsPerMonth field to given value.

### HasMaxEventsPerMonth

`func (o *UsageLimits) HasMaxEventsPerMonth() bool`

HasMaxEventsPerMonth returns a boolean if a field has been set.

### SetMaxEventsPerMonthNil

`func (o *UsageLimits) SetMaxEventsPerMonthNil(b bool)`

 SetMaxEventsPerMonthNil sets the value for MaxEventsPerMonth to be an explicit nil

### UnsetMaxEventsPerMonth
`func (o *UsageLimits) UnsetMaxEventsPerMonth()`

UnsetMaxEventsPerMonth ensures that no value is present for MaxEventsPerMonth, not even an explicit nil
### GetGeofenceCount

`func (o *UsageLimits) GetGeofenceCount() int32`

GetGeofenceCount returns the GeofenceCount field if non-nil, zero value otherwise.

### GetGeofenceCountOk

`func (o *UsageLimits) GetGeofenceCountOk() (*int32, bool)`

GetGeofenceCountOk returns a tuple with the GeofenceCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGeofenceCount

`func (o *UsageLimits) SetGeofenceCount(v int32)`

SetGeofenceCount sets GeofenceCount field to given value.

### HasGeofenceCount

`func (o *UsageLimits) HasGeofenceCount() bool`

HasGeofenceCount returns a boolean if a field has been set.

### GetDeviceCount

`func (o *UsageLimits) GetDeviceCount() int32`

GetDeviceCount returns the DeviceCount field if non-nil, zero value otherwise.

### GetDeviceCountOk

`func (o *UsageLimits) GetDeviceCountOk() (*int32, bool)`

GetDeviceCountOk returns a tuple with the DeviceCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeviceCount

`func (o *UsageLimits) SetDeviceCount(v int32)`

SetDeviceCount sets DeviceCount field to given value.

### HasDeviceCount

`func (o *UsageLimits) HasDeviceCount() bool`

HasDeviceCount returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


