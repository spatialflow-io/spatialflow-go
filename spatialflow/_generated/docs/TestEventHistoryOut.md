# TestEventHistoryOut

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**GeofenceId** | **string** |  | 
**GeofenceName** | **string** |  | 
**TestEvents** | [**[]TestEventItemOut**](TestEventItemOut.md) |  | 
**Count** | **int32** |  | 
**Total** | **int32** |  | 

## Methods

### NewTestEventHistoryOut

`func NewTestEventHistoryOut(geofenceId string, geofenceName string, testEvents []TestEventItemOut, count int32, total int32, ) *TestEventHistoryOut`

NewTestEventHistoryOut instantiates a new TestEventHistoryOut object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTestEventHistoryOutWithDefaults

`func NewTestEventHistoryOutWithDefaults() *TestEventHistoryOut`

NewTestEventHistoryOutWithDefaults instantiates a new TestEventHistoryOut object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetGeofenceId

`func (o *TestEventHistoryOut) GetGeofenceId() string`

GetGeofenceId returns the GeofenceId field if non-nil, zero value otherwise.

### GetGeofenceIdOk

`func (o *TestEventHistoryOut) GetGeofenceIdOk() (*string, bool)`

GetGeofenceIdOk returns a tuple with the GeofenceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGeofenceId

`func (o *TestEventHistoryOut) SetGeofenceId(v string)`

SetGeofenceId sets GeofenceId field to given value.


### GetGeofenceName

`func (o *TestEventHistoryOut) GetGeofenceName() string`

GetGeofenceName returns the GeofenceName field if non-nil, zero value otherwise.

### GetGeofenceNameOk

`func (o *TestEventHistoryOut) GetGeofenceNameOk() (*string, bool)`

GetGeofenceNameOk returns a tuple with the GeofenceName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGeofenceName

`func (o *TestEventHistoryOut) SetGeofenceName(v string)`

SetGeofenceName sets GeofenceName field to given value.


### GetTestEvents

`func (o *TestEventHistoryOut) GetTestEvents() []TestEventItemOut`

GetTestEvents returns the TestEvents field if non-nil, zero value otherwise.

### GetTestEventsOk

`func (o *TestEventHistoryOut) GetTestEventsOk() (*[]TestEventItemOut, bool)`

GetTestEventsOk returns a tuple with the TestEvents field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTestEvents

`func (o *TestEventHistoryOut) SetTestEvents(v []TestEventItemOut)`

SetTestEvents sets TestEvents field to given value.


### GetCount

`func (o *TestEventHistoryOut) GetCount() int32`

GetCount returns the Count field if non-nil, zero value otherwise.

### GetCountOk

`func (o *TestEventHistoryOut) GetCountOk() (*int32, bool)`

GetCountOk returns a tuple with the Count field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCount

`func (o *TestEventHistoryOut) SetCount(v int32)`

SetCount sets Count field to given value.


### GetTotal

`func (o *TestEventHistoryOut) GetTotal() int32`

GetTotal returns the Total field if non-nil, zero value otherwise.

### GetTotalOk

`func (o *TestEventHistoryOut) GetTotalOk() (*int32, bool)`

GetTotalOk returns a tuple with the Total field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotal

`func (o *TestEventHistoryOut) SetTotal(v int32)`

SetTotal sets Total field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


