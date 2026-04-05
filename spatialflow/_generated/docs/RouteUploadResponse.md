# RouteUploadResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Route** | [**SimulationRouteOut**](SimulationRouteOut.md) |  | 
**GpxRouteId** | **string** |  | 
**TotalPoints** | **int32** |  | 
**TotalDistanceKm** | **float32** |  | 
**TotalDurationSeconds** | **float32** |  | 

## Methods

### NewRouteUploadResponse

`func NewRouteUploadResponse(route SimulationRouteOut, gpxRouteId string, totalPoints int32, totalDistanceKm float32, totalDurationSeconds float32, ) *RouteUploadResponse`

NewRouteUploadResponse instantiates a new RouteUploadResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRouteUploadResponseWithDefaults

`func NewRouteUploadResponseWithDefaults() *RouteUploadResponse`

NewRouteUploadResponseWithDefaults instantiates a new RouteUploadResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRoute

`func (o *RouteUploadResponse) GetRoute() SimulationRouteOut`

GetRoute returns the Route field if non-nil, zero value otherwise.

### GetRouteOk

`func (o *RouteUploadResponse) GetRouteOk() (*SimulationRouteOut, bool)`

GetRouteOk returns a tuple with the Route field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRoute

`func (o *RouteUploadResponse) SetRoute(v SimulationRouteOut)`

SetRoute sets Route field to given value.


### GetGpxRouteId

`func (o *RouteUploadResponse) GetGpxRouteId() string`

GetGpxRouteId returns the GpxRouteId field if non-nil, zero value otherwise.

### GetGpxRouteIdOk

`func (o *RouteUploadResponse) GetGpxRouteIdOk() (*string, bool)`

GetGpxRouteIdOk returns a tuple with the GpxRouteId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGpxRouteId

`func (o *RouteUploadResponse) SetGpxRouteId(v string)`

SetGpxRouteId sets GpxRouteId field to given value.


### GetTotalPoints

`func (o *RouteUploadResponse) GetTotalPoints() int32`

GetTotalPoints returns the TotalPoints field if non-nil, zero value otherwise.

### GetTotalPointsOk

`func (o *RouteUploadResponse) GetTotalPointsOk() (*int32, bool)`

GetTotalPointsOk returns a tuple with the TotalPoints field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalPoints

`func (o *RouteUploadResponse) SetTotalPoints(v int32)`

SetTotalPoints sets TotalPoints field to given value.


### GetTotalDistanceKm

`func (o *RouteUploadResponse) GetTotalDistanceKm() float32`

GetTotalDistanceKm returns the TotalDistanceKm field if non-nil, zero value otherwise.

### GetTotalDistanceKmOk

`func (o *RouteUploadResponse) GetTotalDistanceKmOk() (*float32, bool)`

GetTotalDistanceKmOk returns a tuple with the TotalDistanceKm field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalDistanceKm

`func (o *RouteUploadResponse) SetTotalDistanceKm(v float32)`

SetTotalDistanceKm sets TotalDistanceKm field to given value.


### GetTotalDurationSeconds

`func (o *RouteUploadResponse) GetTotalDurationSeconds() float32`

GetTotalDurationSeconds returns the TotalDurationSeconds field if non-nil, zero value otherwise.

### GetTotalDurationSecondsOk

`func (o *RouteUploadResponse) GetTotalDurationSecondsOk() (*float32, bool)`

GetTotalDurationSecondsOk returns a tuple with the TotalDurationSeconds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalDurationSeconds

`func (o *RouteUploadResponse) SetTotalDurationSeconds(v float32)`

SetTotalDurationSeconds sets TotalDurationSeconds field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


