# MatchedGeofenceItem

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**GeofenceId** | **string** |  | 
**GeofenceName** | **string** |  | 
**GroupId** | Pointer to **NullableString** |  | [optional] 
**GroupName** | Pointer to **NullableString** |  | [optional] 
**DistanceMeters** | Pointer to **NullableFloat32** |  | [optional] 

## Methods

### NewMatchedGeofenceItem

`func NewMatchedGeofenceItem(geofenceId string, geofenceName string, ) *MatchedGeofenceItem`

NewMatchedGeofenceItem instantiates a new MatchedGeofenceItem object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMatchedGeofenceItemWithDefaults

`func NewMatchedGeofenceItemWithDefaults() *MatchedGeofenceItem`

NewMatchedGeofenceItemWithDefaults instantiates a new MatchedGeofenceItem object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetGeofenceId

`func (o *MatchedGeofenceItem) GetGeofenceId() string`

GetGeofenceId returns the GeofenceId field if non-nil, zero value otherwise.

### GetGeofenceIdOk

`func (o *MatchedGeofenceItem) GetGeofenceIdOk() (*string, bool)`

GetGeofenceIdOk returns a tuple with the GeofenceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGeofenceId

`func (o *MatchedGeofenceItem) SetGeofenceId(v string)`

SetGeofenceId sets GeofenceId field to given value.


### GetGeofenceName

`func (o *MatchedGeofenceItem) GetGeofenceName() string`

GetGeofenceName returns the GeofenceName field if non-nil, zero value otherwise.

### GetGeofenceNameOk

`func (o *MatchedGeofenceItem) GetGeofenceNameOk() (*string, bool)`

GetGeofenceNameOk returns a tuple with the GeofenceName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGeofenceName

`func (o *MatchedGeofenceItem) SetGeofenceName(v string)`

SetGeofenceName sets GeofenceName field to given value.


### GetGroupId

`func (o *MatchedGeofenceItem) GetGroupId() string`

GetGroupId returns the GroupId field if non-nil, zero value otherwise.

### GetGroupIdOk

`func (o *MatchedGeofenceItem) GetGroupIdOk() (*string, bool)`

GetGroupIdOk returns a tuple with the GroupId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroupId

`func (o *MatchedGeofenceItem) SetGroupId(v string)`

SetGroupId sets GroupId field to given value.

### HasGroupId

`func (o *MatchedGeofenceItem) HasGroupId() bool`

HasGroupId returns a boolean if a field has been set.

### SetGroupIdNil

`func (o *MatchedGeofenceItem) SetGroupIdNil(b bool)`

 SetGroupIdNil sets the value for GroupId to be an explicit nil

### UnsetGroupId
`func (o *MatchedGeofenceItem) UnsetGroupId()`

UnsetGroupId ensures that no value is present for GroupId, not even an explicit nil
### GetGroupName

`func (o *MatchedGeofenceItem) GetGroupName() string`

GetGroupName returns the GroupName field if non-nil, zero value otherwise.

### GetGroupNameOk

`func (o *MatchedGeofenceItem) GetGroupNameOk() (*string, bool)`

GetGroupNameOk returns a tuple with the GroupName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroupName

`func (o *MatchedGeofenceItem) SetGroupName(v string)`

SetGroupName sets GroupName field to given value.

### HasGroupName

`func (o *MatchedGeofenceItem) HasGroupName() bool`

HasGroupName returns a boolean if a field has been set.

### SetGroupNameNil

`func (o *MatchedGeofenceItem) SetGroupNameNil(b bool)`

 SetGroupNameNil sets the value for GroupName to be an explicit nil

### UnsetGroupName
`func (o *MatchedGeofenceItem) UnsetGroupName()`

UnsetGroupName ensures that no value is present for GroupName, not even an explicit nil
### GetDistanceMeters

`func (o *MatchedGeofenceItem) GetDistanceMeters() float32`

GetDistanceMeters returns the DistanceMeters field if non-nil, zero value otherwise.

### GetDistanceMetersOk

`func (o *MatchedGeofenceItem) GetDistanceMetersOk() (*float32, bool)`

GetDistanceMetersOk returns a tuple with the DistanceMeters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDistanceMeters

`func (o *MatchedGeofenceItem) SetDistanceMeters(v float32)`

SetDistanceMeters sets DistanceMeters field to given value.

### HasDistanceMeters

`func (o *MatchedGeofenceItem) HasDistanceMeters() bool`

HasDistanceMeters returns a boolean if a field has been set.

### SetDistanceMetersNil

`func (o *MatchedGeofenceItem) SetDistanceMetersNil(b bool)`

 SetDistanceMetersNil sets the value for DistanceMeters to be an explicit nil

### UnsetDistanceMeters
`func (o *MatchedGeofenceItem) UnsetDistanceMeters()`

UnsetDistanceMeters ensures that no value is present for DistanceMeters, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


