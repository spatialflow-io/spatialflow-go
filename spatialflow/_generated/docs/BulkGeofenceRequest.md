# BulkGeofenceRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Geofences** | [**[]CreateGeofenceRequest**](CreateGeofenceRequest.md) |  | 

## Methods

### NewBulkGeofenceRequest

`func NewBulkGeofenceRequest(geofences []CreateGeofenceRequest, ) *BulkGeofenceRequest`

NewBulkGeofenceRequest instantiates a new BulkGeofenceRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBulkGeofenceRequestWithDefaults

`func NewBulkGeofenceRequestWithDefaults() *BulkGeofenceRequest`

NewBulkGeofenceRequestWithDefaults instantiates a new BulkGeofenceRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetGeofences

`func (o *BulkGeofenceRequest) GetGeofences() []CreateGeofenceRequest`

GetGeofences returns the Geofences field if non-nil, zero value otherwise.

### GetGeofencesOk

`func (o *BulkGeofenceRequest) GetGeofencesOk() (*[]CreateGeofenceRequest, bool)`

GetGeofencesOk returns a tuple with the Geofences field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGeofences

`func (o *BulkGeofenceRequest) SetGeofences(v []CreateGeofenceRequest)`

SetGeofences sets Geofences field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


