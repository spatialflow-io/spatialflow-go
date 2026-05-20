# GeofenceGroupItemOut

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**GroupId** | **string** |  | 
**GroupName** | Pointer to **NullableString** |  | [optional] 
**GeofenceCount** | **int32** |  | 

## Methods

### NewGeofenceGroupItemOut

`func NewGeofenceGroupItemOut(groupId string, geofenceCount int32, ) *GeofenceGroupItemOut`

NewGeofenceGroupItemOut instantiates a new GeofenceGroupItemOut object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGeofenceGroupItemOutWithDefaults

`func NewGeofenceGroupItemOutWithDefaults() *GeofenceGroupItemOut`

NewGeofenceGroupItemOutWithDefaults instantiates a new GeofenceGroupItemOut object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetGroupId

`func (o *GeofenceGroupItemOut) GetGroupId() string`

GetGroupId returns the GroupId field if non-nil, zero value otherwise.

### GetGroupIdOk

`func (o *GeofenceGroupItemOut) GetGroupIdOk() (*string, bool)`

GetGroupIdOk returns a tuple with the GroupId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroupId

`func (o *GeofenceGroupItemOut) SetGroupId(v string)`

SetGroupId sets GroupId field to given value.


### GetGroupName

`func (o *GeofenceGroupItemOut) GetGroupName() string`

GetGroupName returns the GroupName field if non-nil, zero value otherwise.

### GetGroupNameOk

`func (o *GeofenceGroupItemOut) GetGroupNameOk() (*string, bool)`

GetGroupNameOk returns a tuple with the GroupName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroupName

`func (o *GeofenceGroupItemOut) SetGroupName(v string)`

SetGroupName sets GroupName field to given value.

### HasGroupName

`func (o *GeofenceGroupItemOut) HasGroupName() bool`

HasGroupName returns a boolean if a field has been set.

### SetGroupNameNil

`func (o *GeofenceGroupItemOut) SetGroupNameNil(b bool)`

 SetGroupNameNil sets the value for GroupName to be an explicit nil

### UnsetGroupName
`func (o *GeofenceGroupItemOut) UnsetGroupName()`

UnsetGroupName ensures that no value is present for GroupName, not even an explicit nil
### GetGeofenceCount

`func (o *GeofenceGroupItemOut) GetGeofenceCount() int32`

GetGeofenceCount returns the GeofenceCount field if non-nil, zero value otherwise.

### GetGeofenceCountOk

`func (o *GeofenceGroupItemOut) GetGeofenceCountOk() (*int32, bool)`

GetGeofenceCountOk returns a tuple with the GeofenceCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGeofenceCount

`func (o *GeofenceGroupItemOut) SetGeofenceCount(v int32)`

SetGeofenceCount sets GeofenceCount field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


