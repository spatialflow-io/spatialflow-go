# GroupGeofencesOut

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Geofences** | [**[]GroupGeofenceItemOut**](GroupGeofenceItemOut.md) |  | 
**Count** | **int32** |  | 
**GroupId** | **string** |  | 

## Methods

### NewGroupGeofencesOut

`func NewGroupGeofencesOut(geofences []GroupGeofenceItemOut, count int32, groupId string, ) *GroupGeofencesOut`

NewGroupGeofencesOut instantiates a new GroupGeofencesOut object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGroupGeofencesOutWithDefaults

`func NewGroupGeofencesOutWithDefaults() *GroupGeofencesOut`

NewGroupGeofencesOutWithDefaults instantiates a new GroupGeofencesOut object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetGeofences

`func (o *GroupGeofencesOut) GetGeofences() []GroupGeofenceItemOut`

GetGeofences returns the Geofences field if non-nil, zero value otherwise.

### GetGeofencesOk

`func (o *GroupGeofencesOut) GetGeofencesOk() (*[]GroupGeofenceItemOut, bool)`

GetGeofencesOk returns a tuple with the Geofences field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGeofences

`func (o *GroupGeofencesOut) SetGeofences(v []GroupGeofenceItemOut)`

SetGeofences sets Geofences field to given value.


### GetCount

`func (o *GroupGeofencesOut) GetCount() int32`

GetCount returns the Count field if non-nil, zero value otherwise.

### GetCountOk

`func (o *GroupGeofencesOut) GetCountOk() (*int32, bool)`

GetCountOk returns a tuple with the Count field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCount

`func (o *GroupGeofencesOut) SetCount(v int32)`

SetCount sets Count field to given value.


### GetGroupId

`func (o *GroupGeofencesOut) GetGroupId() string`

GetGroupId returns the GroupId field if non-nil, zero value otherwise.

### GetGroupIdOk

`func (o *GroupGeofencesOut) GetGroupIdOk() (*string, bool)`

GetGroupIdOk returns a tuple with the GroupId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroupId

`func (o *GroupGeofencesOut) SetGroupId(v string)`

SetGroupId sets GroupId field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


