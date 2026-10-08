# EventDeviceOut

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** |  | 
**DeviceId** | **string** |  | 
**Name** | **string** |  | 
**DriverName** | Pointer to **NullableString** |  | [optional] 
**Type** | **string** |  | 
**IsExample** | Pointer to **bool** |  | [optional] [default to false]
**LastAccuracy** | Pointer to **NullableFloat32** |  | [optional] 
**LastBatteryLevel** | Pointer to **NullableInt32** |  | [optional] 
**LastBatteryCharging** | Pointer to **NullableBool** |  | [optional] 
**LastBatteryTime** | Pointer to **NullableTime** |  | [optional] 

## Methods

### NewEventDeviceOut

`func NewEventDeviceOut(id string, deviceId string, name string, type_ string, ) *EventDeviceOut`

NewEventDeviceOut instantiates a new EventDeviceOut object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEventDeviceOutWithDefaults

`func NewEventDeviceOutWithDefaults() *EventDeviceOut`

NewEventDeviceOutWithDefaults instantiates a new EventDeviceOut object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *EventDeviceOut) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *EventDeviceOut) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *EventDeviceOut) SetId(v string)`

SetId sets Id field to given value.


### GetDeviceId

`func (o *EventDeviceOut) GetDeviceId() string`

GetDeviceId returns the DeviceId field if non-nil, zero value otherwise.

### GetDeviceIdOk

`func (o *EventDeviceOut) GetDeviceIdOk() (*string, bool)`

GetDeviceIdOk returns a tuple with the DeviceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeviceId

`func (o *EventDeviceOut) SetDeviceId(v string)`

SetDeviceId sets DeviceId field to given value.


### GetName

`func (o *EventDeviceOut) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *EventDeviceOut) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *EventDeviceOut) SetName(v string)`

SetName sets Name field to given value.


### GetDriverName

`func (o *EventDeviceOut) GetDriverName() string`

GetDriverName returns the DriverName field if non-nil, zero value otherwise.

### GetDriverNameOk

`func (o *EventDeviceOut) GetDriverNameOk() (*string, bool)`

GetDriverNameOk returns a tuple with the DriverName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDriverName

`func (o *EventDeviceOut) SetDriverName(v string)`

SetDriverName sets DriverName field to given value.

### HasDriverName

`func (o *EventDeviceOut) HasDriverName() bool`

HasDriverName returns a boolean if a field has been set.

### SetDriverNameNil

`func (o *EventDeviceOut) SetDriverNameNil(b bool)`

 SetDriverNameNil sets the value for DriverName to be an explicit nil

### UnsetDriverName
`func (o *EventDeviceOut) UnsetDriverName()`

UnsetDriverName ensures that no value is present for DriverName, not even an explicit nil
### GetType

`func (o *EventDeviceOut) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *EventDeviceOut) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *EventDeviceOut) SetType(v string)`

SetType sets Type field to given value.


### GetIsExample

`func (o *EventDeviceOut) GetIsExample() bool`

GetIsExample returns the IsExample field if non-nil, zero value otherwise.

### GetIsExampleOk

`func (o *EventDeviceOut) GetIsExampleOk() (*bool, bool)`

GetIsExampleOk returns a tuple with the IsExample field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsExample

`func (o *EventDeviceOut) SetIsExample(v bool)`

SetIsExample sets IsExample field to given value.

### HasIsExample

`func (o *EventDeviceOut) HasIsExample() bool`

HasIsExample returns a boolean if a field has been set.

### GetLastAccuracy

`func (o *EventDeviceOut) GetLastAccuracy() float32`

GetLastAccuracy returns the LastAccuracy field if non-nil, zero value otherwise.

### GetLastAccuracyOk

`func (o *EventDeviceOut) GetLastAccuracyOk() (*float32, bool)`

GetLastAccuracyOk returns a tuple with the LastAccuracy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastAccuracy

`func (o *EventDeviceOut) SetLastAccuracy(v float32)`

SetLastAccuracy sets LastAccuracy field to given value.

### HasLastAccuracy

`func (o *EventDeviceOut) HasLastAccuracy() bool`

HasLastAccuracy returns a boolean if a field has been set.

### SetLastAccuracyNil

`func (o *EventDeviceOut) SetLastAccuracyNil(b bool)`

 SetLastAccuracyNil sets the value for LastAccuracy to be an explicit nil

### UnsetLastAccuracy
`func (o *EventDeviceOut) UnsetLastAccuracy()`

UnsetLastAccuracy ensures that no value is present for LastAccuracy, not even an explicit nil
### GetLastBatteryLevel

`func (o *EventDeviceOut) GetLastBatteryLevel() int32`

GetLastBatteryLevel returns the LastBatteryLevel field if non-nil, zero value otherwise.

### GetLastBatteryLevelOk

`func (o *EventDeviceOut) GetLastBatteryLevelOk() (*int32, bool)`

GetLastBatteryLevelOk returns a tuple with the LastBatteryLevel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastBatteryLevel

`func (o *EventDeviceOut) SetLastBatteryLevel(v int32)`

SetLastBatteryLevel sets LastBatteryLevel field to given value.

### HasLastBatteryLevel

`func (o *EventDeviceOut) HasLastBatteryLevel() bool`

HasLastBatteryLevel returns a boolean if a field has been set.

### SetLastBatteryLevelNil

`func (o *EventDeviceOut) SetLastBatteryLevelNil(b bool)`

 SetLastBatteryLevelNil sets the value for LastBatteryLevel to be an explicit nil

### UnsetLastBatteryLevel
`func (o *EventDeviceOut) UnsetLastBatteryLevel()`

UnsetLastBatteryLevel ensures that no value is present for LastBatteryLevel, not even an explicit nil
### GetLastBatteryCharging

`func (o *EventDeviceOut) GetLastBatteryCharging() bool`

GetLastBatteryCharging returns the LastBatteryCharging field if non-nil, zero value otherwise.

### GetLastBatteryChargingOk

`func (o *EventDeviceOut) GetLastBatteryChargingOk() (*bool, bool)`

GetLastBatteryChargingOk returns a tuple with the LastBatteryCharging field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastBatteryCharging

`func (o *EventDeviceOut) SetLastBatteryCharging(v bool)`

SetLastBatteryCharging sets LastBatteryCharging field to given value.

### HasLastBatteryCharging

`func (o *EventDeviceOut) HasLastBatteryCharging() bool`

HasLastBatteryCharging returns a boolean if a field has been set.

### SetLastBatteryChargingNil

`func (o *EventDeviceOut) SetLastBatteryChargingNil(b bool)`

 SetLastBatteryChargingNil sets the value for LastBatteryCharging to be an explicit nil

### UnsetLastBatteryCharging
`func (o *EventDeviceOut) UnsetLastBatteryCharging()`

UnsetLastBatteryCharging ensures that no value is present for LastBatteryCharging, not even an explicit nil
### GetLastBatteryTime

`func (o *EventDeviceOut) GetLastBatteryTime() time.Time`

GetLastBatteryTime returns the LastBatteryTime field if non-nil, zero value otherwise.

### GetLastBatteryTimeOk

`func (o *EventDeviceOut) GetLastBatteryTimeOk() (*time.Time, bool)`

GetLastBatteryTimeOk returns a tuple with the LastBatteryTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastBatteryTime

`func (o *EventDeviceOut) SetLastBatteryTime(v time.Time)`

SetLastBatteryTime sets LastBatteryTime field to given value.

### HasLastBatteryTime

`func (o *EventDeviceOut) HasLastBatteryTime() bool`

HasLastBatteryTime returns a boolean if a field has been set.

### SetLastBatteryTimeNil

`func (o *EventDeviceOut) SetLastBatteryTimeNil(b bool)`

 SetLastBatteryTimeNil sets the value for LastBatteryTime to be an explicit nil

### UnsetLastBatteryTime
`func (o *EventDeviceOut) UnsetLastBatteryTime()`

UnsetLastBatteryTime ensures that no value is present for LastBatteryTime, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


