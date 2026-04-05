# TrackPointOut

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Lat** | **float32** |  | 
**Lon** | **float32** |  | 
**Ele** | Pointer to **NullableFloat32** |  | [optional] 
**Time** | Pointer to **NullableString** |  | [optional] 

## Methods

### NewTrackPointOut

`func NewTrackPointOut(lat float32, lon float32, ) *TrackPointOut`

NewTrackPointOut instantiates a new TrackPointOut object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTrackPointOutWithDefaults

`func NewTrackPointOutWithDefaults() *TrackPointOut`

NewTrackPointOutWithDefaults instantiates a new TrackPointOut object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLat

`func (o *TrackPointOut) GetLat() float32`

GetLat returns the Lat field if non-nil, zero value otherwise.

### GetLatOk

`func (o *TrackPointOut) GetLatOk() (*float32, bool)`

GetLatOk returns a tuple with the Lat field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLat

`func (o *TrackPointOut) SetLat(v float32)`

SetLat sets Lat field to given value.


### GetLon

`func (o *TrackPointOut) GetLon() float32`

GetLon returns the Lon field if non-nil, zero value otherwise.

### GetLonOk

`func (o *TrackPointOut) GetLonOk() (*float32, bool)`

GetLonOk returns a tuple with the Lon field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLon

`func (o *TrackPointOut) SetLon(v float32)`

SetLon sets Lon field to given value.


### GetEle

`func (o *TrackPointOut) GetEle() float32`

GetEle returns the Ele field if non-nil, zero value otherwise.

### GetEleOk

`func (o *TrackPointOut) GetEleOk() (*float32, bool)`

GetEleOk returns a tuple with the Ele field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEle

`func (o *TrackPointOut) SetEle(v float32)`

SetEle sets Ele field to given value.

### HasEle

`func (o *TrackPointOut) HasEle() bool`

HasEle returns a boolean if a field has been set.

### SetEleNil

`func (o *TrackPointOut) SetEleNil(b bool)`

 SetEleNil sets the value for Ele to be an explicit nil

### UnsetEle
`func (o *TrackPointOut) UnsetEle()`

UnsetEle ensures that no value is present for Ele, not even an explicit nil
### GetTime

`func (o *TrackPointOut) GetTime() string`

GetTime returns the Time field if non-nil, zero value otherwise.

### GetTimeOk

`func (o *TrackPointOut) GetTimeOk() (*string, bool)`

GetTimeOk returns a tuple with the Time field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTime

`func (o *TrackPointOut) SetTime(v string)`

SetTime sets Time field to given value.

### HasTime

`func (o *TrackPointOut) HasTime() bool`

HasTime returns a boolean if a field has been set.

### SetTimeNil

`func (o *TrackPointOut) SetTimeNil(b bool)`

 SetTimeNil sets the value for Time to be an explicit nil

### UnsetTime
`func (o *TrackPointOut) UnsetTime()`

UnsetTime ensures that no value is present for Time, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


