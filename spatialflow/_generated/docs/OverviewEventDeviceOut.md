# OverviewEventDeviceOut

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** |  | 
**DeviceId** | **string** |  | 
**Name** | **string** |  | 
**DriverName** | Pointer to **NullableString** |  | [optional] 
**Type** | **string** |  | 
**LastAccuracy** | Pointer to **NullableFloat32** |  | [optional] 
**LastBatteryLevel** | Pointer to **NullableInt32** |  | [optional] 
**LastBatteryCharging** | Pointer to **NullableBool** |  | [optional] 
**LastBatteryTime** | Pointer to **NullableTime** |  | [optional] 

## Methods

### NewOverviewEventDeviceOut

`func NewOverviewEventDeviceOut(id string, deviceId string, name string, type_ string, ) *OverviewEventDeviceOut`

NewOverviewEventDeviceOut instantiates a new OverviewEventDeviceOut object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOverviewEventDeviceOutWithDefaults

`func NewOverviewEventDeviceOutWithDefaults() *OverviewEventDeviceOut`

NewOverviewEventDeviceOutWithDefaults instantiates a new OverviewEventDeviceOut object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *OverviewEventDeviceOut) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *OverviewEventDeviceOut) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *OverviewEventDeviceOut) SetId(v string)`

SetId sets Id field to given value.


### GetDeviceId

`func (o *OverviewEventDeviceOut) GetDeviceId() string`

GetDeviceId returns the DeviceId field if non-nil, zero value otherwise.

### GetDeviceIdOk

`func (o *OverviewEventDeviceOut) GetDeviceIdOk() (*string, bool)`

GetDeviceIdOk returns a tuple with the DeviceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeviceId

`func (o *OverviewEventDeviceOut) SetDeviceId(v string)`

SetDeviceId sets DeviceId field to given value.


### GetName

`func (o *OverviewEventDeviceOut) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *OverviewEventDeviceOut) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *OverviewEventDeviceOut) SetName(v string)`

SetName sets Name field to given value.


### GetDriverName

`func (o *OverviewEventDeviceOut) GetDriverName() string`

GetDriverName returns the DriverName field if non-nil, zero value otherwise.

### GetDriverNameOk

`func (o *OverviewEventDeviceOut) GetDriverNameOk() (*string, bool)`

GetDriverNameOk returns a tuple with the DriverName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDriverName

`func (o *OverviewEventDeviceOut) SetDriverName(v string)`

SetDriverName sets DriverName field to given value.

### HasDriverName

`func (o *OverviewEventDeviceOut) HasDriverName() bool`

HasDriverName returns a boolean if a field has been set.

### SetDriverNameNil

`func (o *OverviewEventDeviceOut) SetDriverNameNil(b bool)`

 SetDriverNameNil sets the value for DriverName to be an explicit nil

### UnsetDriverName
`func (o *OverviewEventDeviceOut) UnsetDriverName()`

UnsetDriverName ensures that no value is present for DriverName, not even an explicit nil
### GetType

`func (o *OverviewEventDeviceOut) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *OverviewEventDeviceOut) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *OverviewEventDeviceOut) SetType(v string)`

SetType sets Type field to given value.


### GetLastAccuracy

`func (o *OverviewEventDeviceOut) GetLastAccuracy() float32`

GetLastAccuracy returns the LastAccuracy field if non-nil, zero value otherwise.

### GetLastAccuracyOk

`func (o *OverviewEventDeviceOut) GetLastAccuracyOk() (*float32, bool)`

GetLastAccuracyOk returns a tuple with the LastAccuracy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastAccuracy

`func (o *OverviewEventDeviceOut) SetLastAccuracy(v float32)`

SetLastAccuracy sets LastAccuracy field to given value.

### HasLastAccuracy

`func (o *OverviewEventDeviceOut) HasLastAccuracy() bool`

HasLastAccuracy returns a boolean if a field has been set.

### SetLastAccuracyNil

`func (o *OverviewEventDeviceOut) SetLastAccuracyNil(b bool)`

 SetLastAccuracyNil sets the value for LastAccuracy to be an explicit nil

### UnsetLastAccuracy
`func (o *OverviewEventDeviceOut) UnsetLastAccuracy()`

UnsetLastAccuracy ensures that no value is present for LastAccuracy, not even an explicit nil
### GetLastBatteryLevel

`func (o *OverviewEventDeviceOut) GetLastBatteryLevel() int32`

GetLastBatteryLevel returns the LastBatteryLevel field if non-nil, zero value otherwise.

### GetLastBatteryLevelOk

`func (o *OverviewEventDeviceOut) GetLastBatteryLevelOk() (*int32, bool)`

GetLastBatteryLevelOk returns a tuple with the LastBatteryLevel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastBatteryLevel

`func (o *OverviewEventDeviceOut) SetLastBatteryLevel(v int32)`

SetLastBatteryLevel sets LastBatteryLevel field to given value.

### HasLastBatteryLevel

`func (o *OverviewEventDeviceOut) HasLastBatteryLevel() bool`

HasLastBatteryLevel returns a boolean if a field has been set.

### SetLastBatteryLevelNil

`func (o *OverviewEventDeviceOut) SetLastBatteryLevelNil(b bool)`

 SetLastBatteryLevelNil sets the value for LastBatteryLevel to be an explicit nil

### UnsetLastBatteryLevel
`func (o *OverviewEventDeviceOut) UnsetLastBatteryLevel()`

UnsetLastBatteryLevel ensures that no value is present for LastBatteryLevel, not even an explicit nil
### GetLastBatteryCharging

`func (o *OverviewEventDeviceOut) GetLastBatteryCharging() bool`

GetLastBatteryCharging returns the LastBatteryCharging field if non-nil, zero value otherwise.

### GetLastBatteryChargingOk

`func (o *OverviewEventDeviceOut) GetLastBatteryChargingOk() (*bool, bool)`

GetLastBatteryChargingOk returns a tuple with the LastBatteryCharging field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastBatteryCharging

`func (o *OverviewEventDeviceOut) SetLastBatteryCharging(v bool)`

SetLastBatteryCharging sets LastBatteryCharging field to given value.

### HasLastBatteryCharging

`func (o *OverviewEventDeviceOut) HasLastBatteryCharging() bool`

HasLastBatteryCharging returns a boolean if a field has been set.

### SetLastBatteryChargingNil

`func (o *OverviewEventDeviceOut) SetLastBatteryChargingNil(b bool)`

 SetLastBatteryChargingNil sets the value for LastBatteryCharging to be an explicit nil

### UnsetLastBatteryCharging
`func (o *OverviewEventDeviceOut) UnsetLastBatteryCharging()`

UnsetLastBatteryCharging ensures that no value is present for LastBatteryCharging, not even an explicit nil
### GetLastBatteryTime

`func (o *OverviewEventDeviceOut) GetLastBatteryTime() time.Time`

GetLastBatteryTime returns the LastBatteryTime field if non-nil, zero value otherwise.

### GetLastBatteryTimeOk

`func (o *OverviewEventDeviceOut) GetLastBatteryTimeOk() (*time.Time, bool)`

GetLastBatteryTimeOk returns a tuple with the LastBatteryTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastBatteryTime

`func (o *OverviewEventDeviceOut) SetLastBatteryTime(v time.Time)`

SetLastBatteryTime sets LastBatteryTime field to given value.

### HasLastBatteryTime

`func (o *OverviewEventDeviceOut) HasLastBatteryTime() bool`

HasLastBatteryTime returns a boolean if a field has been set.

### SetLastBatteryTimeNil

`func (o *OverviewEventDeviceOut) SetLastBatteryTimeNil(b bool)`

 SetLastBatteryTimeNil sets the value for LastBatteryTime to be an explicit nil

### UnsetLastBatteryTime
`func (o *OverviewEventDeviceOut) UnsetLastBatteryTime()`

UnsetLastBatteryTime ensures that no value is present for LastBatteryTime, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


