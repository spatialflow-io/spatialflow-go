# GeofenceGroupsOut

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Groups** | [**[]GeofenceGroupItemOut**](GeofenceGroupItemOut.md) |  | 
**Count** | **int32** |  | 

## Methods

### NewGeofenceGroupsOut

`func NewGeofenceGroupsOut(groups []GeofenceGroupItemOut, count int32, ) *GeofenceGroupsOut`

NewGeofenceGroupsOut instantiates a new GeofenceGroupsOut object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGeofenceGroupsOutWithDefaults

`func NewGeofenceGroupsOutWithDefaults() *GeofenceGroupsOut`

NewGeofenceGroupsOutWithDefaults instantiates a new GeofenceGroupsOut object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetGroups

`func (o *GeofenceGroupsOut) GetGroups() []GeofenceGroupItemOut`

GetGroups returns the Groups field if non-nil, zero value otherwise.

### GetGroupsOk

`func (o *GeofenceGroupsOut) GetGroupsOk() (*[]GeofenceGroupItemOut, bool)`

GetGroupsOk returns a tuple with the Groups field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroups

`func (o *GeofenceGroupsOut) SetGroups(v []GeofenceGroupItemOut)`

SetGroups sets Groups field to given value.


### GetCount

`func (o *GeofenceGroupsOut) GetCount() int32`

GetCount returns the Count field if non-nil, zero value otherwise.

### GetCountOk

`func (o *GeofenceGroupsOut) GetCountOk() (*int32, bool)`

GetCountOk returns a tuple with the Count field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCount

`func (o *GeofenceGroupsOut) SetCount(v int32)`

SetCount sets Count field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


