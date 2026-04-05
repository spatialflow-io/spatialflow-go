# RouteTestResultsOut

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**TestId** | **string** |  | 
**Status** | **string** |  | 
**FileName** | **string** |  | 
**TotalPoints** | **int32** |  | 
**TotalDistanceKm** | **float32** |  | 
**DurationSeconds** | **int32** |  | 
**TrackPoints** | Pointer to **[]map[string]interface{}** |  | [optional] 
**Events** | **[]map[string]interface{}** |  | 
**Summary** | [**RouteTestSummary**](RouteTestSummary.md) |  | 

## Methods

### NewRouteTestResultsOut

`func NewRouteTestResultsOut(testId string, status string, fileName string, totalPoints int32, totalDistanceKm float32, durationSeconds int32, events []*map[string]interface{}, summary RouteTestSummary, ) *RouteTestResultsOut`

NewRouteTestResultsOut instantiates a new RouteTestResultsOut object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRouteTestResultsOutWithDefaults

`func NewRouteTestResultsOutWithDefaults() *RouteTestResultsOut`

NewRouteTestResultsOutWithDefaults instantiates a new RouteTestResultsOut object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTestId

`func (o *RouteTestResultsOut) GetTestId() string`

GetTestId returns the TestId field if non-nil, zero value otherwise.

### GetTestIdOk

`func (o *RouteTestResultsOut) GetTestIdOk() (*string, bool)`

GetTestIdOk returns a tuple with the TestId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTestId

`func (o *RouteTestResultsOut) SetTestId(v string)`

SetTestId sets TestId field to given value.


### GetStatus

`func (o *RouteTestResultsOut) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *RouteTestResultsOut) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *RouteTestResultsOut) SetStatus(v string)`

SetStatus sets Status field to given value.


### GetFileName

`func (o *RouteTestResultsOut) GetFileName() string`

GetFileName returns the FileName field if non-nil, zero value otherwise.

### GetFileNameOk

`func (o *RouteTestResultsOut) GetFileNameOk() (*string, bool)`

GetFileNameOk returns a tuple with the FileName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFileName

`func (o *RouteTestResultsOut) SetFileName(v string)`

SetFileName sets FileName field to given value.


### GetTotalPoints

`func (o *RouteTestResultsOut) GetTotalPoints() int32`

GetTotalPoints returns the TotalPoints field if non-nil, zero value otherwise.

### GetTotalPointsOk

`func (o *RouteTestResultsOut) GetTotalPointsOk() (*int32, bool)`

GetTotalPointsOk returns a tuple with the TotalPoints field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalPoints

`func (o *RouteTestResultsOut) SetTotalPoints(v int32)`

SetTotalPoints sets TotalPoints field to given value.


### GetTotalDistanceKm

`func (o *RouteTestResultsOut) GetTotalDistanceKm() float32`

GetTotalDistanceKm returns the TotalDistanceKm field if non-nil, zero value otherwise.

### GetTotalDistanceKmOk

`func (o *RouteTestResultsOut) GetTotalDistanceKmOk() (*float32, bool)`

GetTotalDistanceKmOk returns a tuple with the TotalDistanceKm field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalDistanceKm

`func (o *RouteTestResultsOut) SetTotalDistanceKm(v float32)`

SetTotalDistanceKm sets TotalDistanceKm field to given value.


### GetDurationSeconds

`func (o *RouteTestResultsOut) GetDurationSeconds() int32`

GetDurationSeconds returns the DurationSeconds field if non-nil, zero value otherwise.

### GetDurationSecondsOk

`func (o *RouteTestResultsOut) GetDurationSecondsOk() (*int32, bool)`

GetDurationSecondsOk returns a tuple with the DurationSeconds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDurationSeconds

`func (o *RouteTestResultsOut) SetDurationSeconds(v int32)`

SetDurationSeconds sets DurationSeconds field to given value.


### GetTrackPoints

`func (o *RouteTestResultsOut) GetTrackPoints() []map[string]interface{}`

GetTrackPoints returns the TrackPoints field if non-nil, zero value otherwise.

### GetTrackPointsOk

`func (o *RouteTestResultsOut) GetTrackPointsOk() (*[]map[string]interface{}, bool)`

GetTrackPointsOk returns a tuple with the TrackPoints field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTrackPoints

`func (o *RouteTestResultsOut) SetTrackPoints(v []map[string]interface{})`

SetTrackPoints sets TrackPoints field to given value.

### HasTrackPoints

`func (o *RouteTestResultsOut) HasTrackPoints() bool`

HasTrackPoints returns a boolean if a field has been set.

### SetTrackPointsNil

`func (o *RouteTestResultsOut) SetTrackPointsNil(b bool)`

 SetTrackPointsNil sets the value for TrackPoints to be an explicit nil

### UnsetTrackPoints
`func (o *RouteTestResultsOut) UnsetTrackPoints()`

UnsetTrackPoints ensures that no value is present for TrackPoints, not even an explicit nil
### GetEvents

`func (o *RouteTestResultsOut) GetEvents() []*map[string]interface{}`

GetEvents returns the Events field if non-nil, zero value otherwise.

### GetEventsOk

`func (o *RouteTestResultsOut) GetEventsOk() (*[]*map[string]interface{}, bool)`

GetEventsOk returns a tuple with the Events field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEvents

`func (o *RouteTestResultsOut) SetEvents(v []*map[string]interface{})`

SetEvents sets Events field to given value.


### GetSummary

`func (o *RouteTestResultsOut) GetSummary() RouteTestSummary`

GetSummary returns the Summary field if non-nil, zero value otherwise.

### GetSummaryOk

`func (o *RouteTestResultsOut) GetSummaryOk() (*RouteTestSummary, bool)`

GetSummaryOk returns a tuple with the Summary field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSummary

`func (o *RouteTestResultsOut) SetSummary(v RouteTestSummary)`

SetSummary sets Summary field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


