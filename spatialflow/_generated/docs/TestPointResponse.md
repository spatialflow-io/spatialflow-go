# TestPointResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Point** | **map[string]interface{}** |  | 
**InsideGeofences** | **int32** |  | 
**TotalGeofences** | **int32** |  | 
**Results** | [**[]GeofenceTestResult**](GeofenceTestResult.md) |  | 
**MatchedGeofences** | Pointer to [**[]MatchedGeofenceItem**](MatchedGeofenceItem.md) | Geofences containing the test point, with group info. Ordered by distance_meters ASC, then geofence_id ASC. | [optional] 
**RequestMetadata** | Pointer to **map[string]interface{}** |  | [optional] 

## Methods

### NewTestPointResponse

`func NewTestPointResponse(point map[string]interface{}, insideGeofences int32, totalGeofences int32, results []GeofenceTestResult, ) *TestPointResponse`

NewTestPointResponse instantiates a new TestPointResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTestPointResponseWithDefaults

`func NewTestPointResponseWithDefaults() *TestPointResponse`

NewTestPointResponseWithDefaults instantiates a new TestPointResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPoint

`func (o *TestPointResponse) GetPoint() map[string]interface{}`

GetPoint returns the Point field if non-nil, zero value otherwise.

### GetPointOk

`func (o *TestPointResponse) GetPointOk() (*map[string]interface{}, bool)`

GetPointOk returns a tuple with the Point field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPoint

`func (o *TestPointResponse) SetPoint(v map[string]interface{})`

SetPoint sets Point field to given value.


### GetInsideGeofences

`func (o *TestPointResponse) GetInsideGeofences() int32`

GetInsideGeofences returns the InsideGeofences field if non-nil, zero value otherwise.

### GetInsideGeofencesOk

`func (o *TestPointResponse) GetInsideGeofencesOk() (*int32, bool)`

GetInsideGeofencesOk returns a tuple with the InsideGeofences field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInsideGeofences

`func (o *TestPointResponse) SetInsideGeofences(v int32)`

SetInsideGeofences sets InsideGeofences field to given value.


### GetTotalGeofences

`func (o *TestPointResponse) GetTotalGeofences() int32`

GetTotalGeofences returns the TotalGeofences field if non-nil, zero value otherwise.

### GetTotalGeofencesOk

`func (o *TestPointResponse) GetTotalGeofencesOk() (*int32, bool)`

GetTotalGeofencesOk returns a tuple with the TotalGeofences field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalGeofences

`func (o *TestPointResponse) SetTotalGeofences(v int32)`

SetTotalGeofences sets TotalGeofences field to given value.


### GetResults

`func (o *TestPointResponse) GetResults() []GeofenceTestResult`

GetResults returns the Results field if non-nil, zero value otherwise.

### GetResultsOk

`func (o *TestPointResponse) GetResultsOk() (*[]GeofenceTestResult, bool)`

GetResultsOk returns a tuple with the Results field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResults

`func (o *TestPointResponse) SetResults(v []GeofenceTestResult)`

SetResults sets Results field to given value.


### GetMatchedGeofences

`func (o *TestPointResponse) GetMatchedGeofences() []MatchedGeofenceItem`

GetMatchedGeofences returns the MatchedGeofences field if non-nil, zero value otherwise.

### GetMatchedGeofencesOk

`func (o *TestPointResponse) GetMatchedGeofencesOk() (*[]MatchedGeofenceItem, bool)`

GetMatchedGeofencesOk returns a tuple with the MatchedGeofences field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMatchedGeofences

`func (o *TestPointResponse) SetMatchedGeofences(v []MatchedGeofenceItem)`

SetMatchedGeofences sets MatchedGeofences field to given value.

### HasMatchedGeofences

`func (o *TestPointResponse) HasMatchedGeofences() bool`

HasMatchedGeofences returns a boolean if a field has been set.

### GetRequestMetadata

`func (o *TestPointResponse) GetRequestMetadata() map[string]interface{}`

GetRequestMetadata returns the RequestMetadata field if non-nil, zero value otherwise.

### GetRequestMetadataOk

`func (o *TestPointResponse) GetRequestMetadataOk() (*map[string]interface{}, bool)`

GetRequestMetadataOk returns a tuple with the RequestMetadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestMetadata

`func (o *TestPointResponse) SetRequestMetadata(v map[string]interface{})`

SetRequestMetadata sets RequestMetadata field to given value.

### HasRequestMetadata

`func (o *TestPointResponse) HasRequestMetadata() bool`

HasRequestMetadata returns a boolean if a field has been set.

### SetRequestMetadataNil

`func (o *TestPointResponse) SetRequestMetadataNil(b bool)`

 SetRequestMetadataNil sets the value for RequestMetadata to be an explicit nil

### UnsetRequestMetadata
`func (o *TestPointResponse) UnsetRequestMetadata()`

UnsetRequestMetadata ensures that no value is present for RequestMetadata, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


