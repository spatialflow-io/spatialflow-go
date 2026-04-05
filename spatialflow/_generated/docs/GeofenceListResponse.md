# GeofenceListResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Geofences** | [**[]GeofenceResponse**](GeofenceResponse.md) |  | 
**Count** | **int32** |  | 
**TotalCount** | **int32** |  | 

## Methods

### NewGeofenceListResponse

`func NewGeofenceListResponse(geofences []GeofenceResponse, count int32, totalCount int32, ) *GeofenceListResponse`

NewGeofenceListResponse instantiates a new GeofenceListResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGeofenceListResponseWithDefaults

`func NewGeofenceListResponseWithDefaults() *GeofenceListResponse`

NewGeofenceListResponseWithDefaults instantiates a new GeofenceListResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetGeofences

`func (o *GeofenceListResponse) GetGeofences() []GeofenceResponse`

GetGeofences returns the Geofences field if non-nil, zero value otherwise.

### GetGeofencesOk

`func (o *GeofenceListResponse) GetGeofencesOk() (*[]GeofenceResponse, bool)`

GetGeofencesOk returns a tuple with the Geofences field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGeofences

`func (o *GeofenceListResponse) SetGeofences(v []GeofenceResponse)`

SetGeofences sets Geofences field to given value.


### GetCount

`func (o *GeofenceListResponse) GetCount() int32`

GetCount returns the Count field if non-nil, zero value otherwise.

### GetCountOk

`func (o *GeofenceListResponse) GetCountOk() (*int32, bool)`

GetCountOk returns a tuple with the Count field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCount

`func (o *GeofenceListResponse) SetCount(v int32)`

SetCount sets Count field to given value.


### GetTotalCount

`func (o *GeofenceListResponse) GetTotalCount() int32`

GetTotalCount returns the TotalCount field if non-nil, zero value otherwise.

### GetTotalCountOk

`func (o *GeofenceListResponse) GetTotalCountOk() (*int32, bool)`

GetTotalCountOk returns a tuple with the TotalCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalCount

`func (o *GeofenceListResponse) SetTotalCount(v int32)`

SetTotalCount sets TotalCount field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


