# CaptureActivityOut

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Kind** | **string** |  | 
**DeviceUuid** | **string** |  | 
**DeviceName** | **string** |  | 
**WorkerName** | **string** |  | 
**SessionId** | **string** |  | 
**Count** | **int32** |  | 
**At** | **time.Time** | Latest upload or note time in the group. | 

## Methods

### NewCaptureActivityOut

`func NewCaptureActivityOut(kind string, deviceUuid string, deviceName string, workerName string, sessionId string, count int32, at time.Time, ) *CaptureActivityOut`

NewCaptureActivityOut instantiates a new CaptureActivityOut object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCaptureActivityOutWithDefaults

`func NewCaptureActivityOutWithDefaults() *CaptureActivityOut`

NewCaptureActivityOutWithDefaults instantiates a new CaptureActivityOut object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetKind

`func (o *CaptureActivityOut) GetKind() string`

GetKind returns the Kind field if non-nil, zero value otherwise.

### GetKindOk

`func (o *CaptureActivityOut) GetKindOk() (*string, bool)`

GetKindOk returns a tuple with the Kind field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKind

`func (o *CaptureActivityOut) SetKind(v string)`

SetKind sets Kind field to given value.


### GetDeviceUuid

`func (o *CaptureActivityOut) GetDeviceUuid() string`

GetDeviceUuid returns the DeviceUuid field if non-nil, zero value otherwise.

### GetDeviceUuidOk

`func (o *CaptureActivityOut) GetDeviceUuidOk() (*string, bool)`

GetDeviceUuidOk returns a tuple with the DeviceUuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeviceUuid

`func (o *CaptureActivityOut) SetDeviceUuid(v string)`

SetDeviceUuid sets DeviceUuid field to given value.


### GetDeviceName

`func (o *CaptureActivityOut) GetDeviceName() string`

GetDeviceName returns the DeviceName field if non-nil, zero value otherwise.

### GetDeviceNameOk

`func (o *CaptureActivityOut) GetDeviceNameOk() (*string, bool)`

GetDeviceNameOk returns a tuple with the DeviceName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeviceName

`func (o *CaptureActivityOut) SetDeviceName(v string)`

SetDeviceName sets DeviceName field to given value.


### GetWorkerName

`func (o *CaptureActivityOut) GetWorkerName() string`

GetWorkerName returns the WorkerName field if non-nil, zero value otherwise.

### GetWorkerNameOk

`func (o *CaptureActivityOut) GetWorkerNameOk() (*string, bool)`

GetWorkerNameOk returns a tuple with the WorkerName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWorkerName

`func (o *CaptureActivityOut) SetWorkerName(v string)`

SetWorkerName sets WorkerName field to given value.


### GetSessionId

`func (o *CaptureActivityOut) GetSessionId() string`

GetSessionId returns the SessionId field if non-nil, zero value otherwise.

### GetSessionIdOk

`func (o *CaptureActivityOut) GetSessionIdOk() (*string, bool)`

GetSessionIdOk returns a tuple with the SessionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSessionId

`func (o *CaptureActivityOut) SetSessionId(v string)`

SetSessionId sets SessionId field to given value.


### GetCount

`func (o *CaptureActivityOut) GetCount() int32`

GetCount returns the Count field if non-nil, zero value otherwise.

### GetCountOk

`func (o *CaptureActivityOut) GetCountOk() (*int32, bool)`

GetCountOk returns a tuple with the Count field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCount

`func (o *CaptureActivityOut) SetCount(v int32)`

SetCount sets Count field to given value.


### GetAt

`func (o *CaptureActivityOut) GetAt() time.Time`

GetAt returns the At field if non-nil, zero value otherwise.

### GetAtOk

`func (o *CaptureActivityOut) GetAtOk() (*time.Time, bool)`

GetAtOk returns a tuple with the At field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAt

`func (o *CaptureActivityOut) SetAt(v time.Time)`

SetAt sets At field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


