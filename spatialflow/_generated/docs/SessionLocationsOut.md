# SessionLocationsOut

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SessionId** | **string** |  | 
**SnapshotAt** | Pointer to **NullableTime** |  | [optional] 
**Locations** | [**[]LocationPointOut**](LocationPointOut.md) |  | 
**TotalCount** | **int32** |  | 
**Offset** | **int32** |  | 
**Limit** | **int32** |  | 
**Simplified** | Pointer to **bool** |  | [optional] [default to false]
**RenderedTrack** | Pointer to **[][]float32** |  | [optional] 
**TrackSource** | Pointer to **string** |  | [optional] [default to "raw"]

## Methods

### NewSessionLocationsOut

`func NewSessionLocationsOut(sessionId string, locations []LocationPointOut, totalCount int32, offset int32, limit int32, ) *SessionLocationsOut`

NewSessionLocationsOut instantiates a new SessionLocationsOut object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSessionLocationsOutWithDefaults

`func NewSessionLocationsOutWithDefaults() *SessionLocationsOut`

NewSessionLocationsOutWithDefaults instantiates a new SessionLocationsOut object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSessionId

`func (o *SessionLocationsOut) GetSessionId() string`

GetSessionId returns the SessionId field if non-nil, zero value otherwise.

### GetSessionIdOk

`func (o *SessionLocationsOut) GetSessionIdOk() (*string, bool)`

GetSessionIdOk returns a tuple with the SessionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSessionId

`func (o *SessionLocationsOut) SetSessionId(v string)`

SetSessionId sets SessionId field to given value.


### GetSnapshotAt

`func (o *SessionLocationsOut) GetSnapshotAt() time.Time`

GetSnapshotAt returns the SnapshotAt field if non-nil, zero value otherwise.

### GetSnapshotAtOk

`func (o *SessionLocationsOut) GetSnapshotAtOk() (*time.Time, bool)`

GetSnapshotAtOk returns a tuple with the SnapshotAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSnapshotAt

`func (o *SessionLocationsOut) SetSnapshotAt(v time.Time)`

SetSnapshotAt sets SnapshotAt field to given value.

### HasSnapshotAt

`func (o *SessionLocationsOut) HasSnapshotAt() bool`

HasSnapshotAt returns a boolean if a field has been set.

### SetSnapshotAtNil

`func (o *SessionLocationsOut) SetSnapshotAtNil(b bool)`

 SetSnapshotAtNil sets the value for SnapshotAt to be an explicit nil

### UnsetSnapshotAt
`func (o *SessionLocationsOut) UnsetSnapshotAt()`

UnsetSnapshotAt ensures that no value is present for SnapshotAt, not even an explicit nil
### GetLocations

`func (o *SessionLocationsOut) GetLocations() []LocationPointOut`

GetLocations returns the Locations field if non-nil, zero value otherwise.

### GetLocationsOk

`func (o *SessionLocationsOut) GetLocationsOk() (*[]LocationPointOut, bool)`

GetLocationsOk returns a tuple with the Locations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocations

`func (o *SessionLocationsOut) SetLocations(v []LocationPointOut)`

SetLocations sets Locations field to given value.


### GetTotalCount

`func (o *SessionLocationsOut) GetTotalCount() int32`

GetTotalCount returns the TotalCount field if non-nil, zero value otherwise.

### GetTotalCountOk

`func (o *SessionLocationsOut) GetTotalCountOk() (*int32, bool)`

GetTotalCountOk returns a tuple with the TotalCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalCount

`func (o *SessionLocationsOut) SetTotalCount(v int32)`

SetTotalCount sets TotalCount field to given value.


### GetOffset

`func (o *SessionLocationsOut) GetOffset() int32`

GetOffset returns the Offset field if non-nil, zero value otherwise.

### GetOffsetOk

`func (o *SessionLocationsOut) GetOffsetOk() (*int32, bool)`

GetOffsetOk returns a tuple with the Offset field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOffset

`func (o *SessionLocationsOut) SetOffset(v int32)`

SetOffset sets Offset field to given value.


### GetLimit

`func (o *SessionLocationsOut) GetLimit() int32`

GetLimit returns the Limit field if non-nil, zero value otherwise.

### GetLimitOk

`func (o *SessionLocationsOut) GetLimitOk() (*int32, bool)`

GetLimitOk returns a tuple with the Limit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLimit

`func (o *SessionLocationsOut) SetLimit(v int32)`

SetLimit sets Limit field to given value.


### GetSimplified

`func (o *SessionLocationsOut) GetSimplified() bool`

GetSimplified returns the Simplified field if non-nil, zero value otherwise.

### GetSimplifiedOk

`func (o *SessionLocationsOut) GetSimplifiedOk() (*bool, bool)`

GetSimplifiedOk returns a tuple with the Simplified field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSimplified

`func (o *SessionLocationsOut) SetSimplified(v bool)`

SetSimplified sets Simplified field to given value.

### HasSimplified

`func (o *SessionLocationsOut) HasSimplified() bool`

HasSimplified returns a boolean if a field has been set.

### GetRenderedTrack

`func (o *SessionLocationsOut) GetRenderedTrack() [][]float32`

GetRenderedTrack returns the RenderedTrack field if non-nil, zero value otherwise.

### GetRenderedTrackOk

`func (o *SessionLocationsOut) GetRenderedTrackOk() (*[][]float32, bool)`

GetRenderedTrackOk returns a tuple with the RenderedTrack field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRenderedTrack

`func (o *SessionLocationsOut) SetRenderedTrack(v [][]float32)`

SetRenderedTrack sets RenderedTrack field to given value.

### HasRenderedTrack

`func (o *SessionLocationsOut) HasRenderedTrack() bool`

HasRenderedTrack returns a boolean if a field has been set.

### SetRenderedTrackNil

`func (o *SessionLocationsOut) SetRenderedTrackNil(b bool)`

 SetRenderedTrackNil sets the value for RenderedTrack to be an explicit nil

### UnsetRenderedTrack
`func (o *SessionLocationsOut) UnsetRenderedTrack()`

UnsetRenderedTrack ensures that no value is present for RenderedTrack, not even an explicit nil
### GetTrackSource

`func (o *SessionLocationsOut) GetTrackSource() string`

GetTrackSource returns the TrackSource field if non-nil, zero value otherwise.

### GetTrackSourceOk

`func (o *SessionLocationsOut) GetTrackSourceOk() (*string, bool)`

GetTrackSourceOk returns a tuple with the TrackSource field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTrackSource

`func (o *SessionLocationsOut) SetTrackSource(v string)`

SetTrackSource sets TrackSource field to given value.

### HasTrackSource

`func (o *SessionLocationsOut) HasTrackSource() bool`

HasTrackSource returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


