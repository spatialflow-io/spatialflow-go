# GroupTestPointOut

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**GroupId** | **string** |  | 
**TestPoint** | [**TestPointCoordinateOut**](TestPointCoordinateOut.md) |  | 
**Results** | [**[]TestPointResultOut**](TestPointResultOut.md) |  | 
**TotalGeofences** | **int32** |  | 
**MatchingGeofences** | **int32** |  | 

## Methods

### NewGroupTestPointOut

`func NewGroupTestPointOut(groupId string, testPoint TestPointCoordinateOut, results []TestPointResultOut, totalGeofences int32, matchingGeofences int32, ) *GroupTestPointOut`

NewGroupTestPointOut instantiates a new GroupTestPointOut object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGroupTestPointOutWithDefaults

`func NewGroupTestPointOutWithDefaults() *GroupTestPointOut`

NewGroupTestPointOutWithDefaults instantiates a new GroupTestPointOut object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetGroupId

`func (o *GroupTestPointOut) GetGroupId() string`

GetGroupId returns the GroupId field if non-nil, zero value otherwise.

### GetGroupIdOk

`func (o *GroupTestPointOut) GetGroupIdOk() (*string, bool)`

GetGroupIdOk returns a tuple with the GroupId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroupId

`func (o *GroupTestPointOut) SetGroupId(v string)`

SetGroupId sets GroupId field to given value.


### GetTestPoint

`func (o *GroupTestPointOut) GetTestPoint() TestPointCoordinateOut`

GetTestPoint returns the TestPoint field if non-nil, zero value otherwise.

### GetTestPointOk

`func (o *GroupTestPointOut) GetTestPointOk() (*TestPointCoordinateOut, bool)`

GetTestPointOk returns a tuple with the TestPoint field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTestPoint

`func (o *GroupTestPointOut) SetTestPoint(v TestPointCoordinateOut)`

SetTestPoint sets TestPoint field to given value.


### GetResults

`func (o *GroupTestPointOut) GetResults() []TestPointResultOut`

GetResults returns the Results field if non-nil, zero value otherwise.

### GetResultsOk

`func (o *GroupTestPointOut) GetResultsOk() (*[]TestPointResultOut, bool)`

GetResultsOk returns a tuple with the Results field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResults

`func (o *GroupTestPointOut) SetResults(v []TestPointResultOut)`

SetResults sets Results field to given value.


### GetTotalGeofences

`func (o *GroupTestPointOut) GetTotalGeofences() int32`

GetTotalGeofences returns the TotalGeofences field if non-nil, zero value otherwise.

### GetTotalGeofencesOk

`func (o *GroupTestPointOut) GetTotalGeofencesOk() (*int32, bool)`

GetTotalGeofencesOk returns a tuple with the TotalGeofences field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalGeofences

`func (o *GroupTestPointOut) SetTotalGeofences(v int32)`

SetTotalGeofences sets TotalGeofences field to given value.


### GetMatchingGeofences

`func (o *GroupTestPointOut) GetMatchingGeofences() int32`

GetMatchingGeofences returns the MatchingGeofences field if non-nil, zero value otherwise.

### GetMatchingGeofencesOk

`func (o *GroupTestPointOut) GetMatchingGeofencesOk() (*int32, bool)`

GetMatchingGeofencesOk returns a tuple with the MatchingGeofences field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMatchingGeofences

`func (o *GroupTestPointOut) SetMatchingGeofences(v int32)`

SetMatchingGeofences sets MatchingGeofences field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


