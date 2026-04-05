# PolicyOut

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** |  | 
**Name** | **string** |  | 
**Description** | **string** |  | 
**IsEnabled** | **bool** |  | 
**GeofenceId** | **string** |  | 
**GeofenceName** | **string** |  | 
**TimeWindow** | Pointer to **map[string]interface{}** |  | [optional] 
**DeviceFilter** | Pointer to **map[string]interface{}** |  | [optional] 
**RoleFilter** | Pointer to **[]string** |  | [optional] 
**TemplateId** | Pointer to **string** |  | [optional] [default to ""]
**CreatedAt** | **time.Time** |  | 
**UpdatedAt** | **time.Time** |  | 

## Methods

### NewPolicyOut

`func NewPolicyOut(id string, name string, description string, isEnabled bool, geofenceId string, geofenceName string, createdAt time.Time, updatedAt time.Time, ) *PolicyOut`

NewPolicyOut instantiates a new PolicyOut object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPolicyOutWithDefaults

`func NewPolicyOutWithDefaults() *PolicyOut`

NewPolicyOutWithDefaults instantiates a new PolicyOut object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *PolicyOut) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *PolicyOut) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *PolicyOut) SetId(v string)`

SetId sets Id field to given value.


### GetName

`func (o *PolicyOut) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *PolicyOut) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *PolicyOut) SetName(v string)`

SetName sets Name field to given value.


### GetDescription

`func (o *PolicyOut) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *PolicyOut) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *PolicyOut) SetDescription(v string)`

SetDescription sets Description field to given value.


### GetIsEnabled

`func (o *PolicyOut) GetIsEnabled() bool`

GetIsEnabled returns the IsEnabled field if non-nil, zero value otherwise.

### GetIsEnabledOk

`func (o *PolicyOut) GetIsEnabledOk() (*bool, bool)`

GetIsEnabledOk returns a tuple with the IsEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsEnabled

`func (o *PolicyOut) SetIsEnabled(v bool)`

SetIsEnabled sets IsEnabled field to given value.


### GetGeofenceId

`func (o *PolicyOut) GetGeofenceId() string`

GetGeofenceId returns the GeofenceId field if non-nil, zero value otherwise.

### GetGeofenceIdOk

`func (o *PolicyOut) GetGeofenceIdOk() (*string, bool)`

GetGeofenceIdOk returns a tuple with the GeofenceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGeofenceId

`func (o *PolicyOut) SetGeofenceId(v string)`

SetGeofenceId sets GeofenceId field to given value.


### GetGeofenceName

`func (o *PolicyOut) GetGeofenceName() string`

GetGeofenceName returns the GeofenceName field if non-nil, zero value otherwise.

### GetGeofenceNameOk

`func (o *PolicyOut) GetGeofenceNameOk() (*string, bool)`

GetGeofenceNameOk returns a tuple with the GeofenceName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGeofenceName

`func (o *PolicyOut) SetGeofenceName(v string)`

SetGeofenceName sets GeofenceName field to given value.


### GetTimeWindow

`func (o *PolicyOut) GetTimeWindow() map[string]interface{}`

GetTimeWindow returns the TimeWindow field if non-nil, zero value otherwise.

### GetTimeWindowOk

`func (o *PolicyOut) GetTimeWindowOk() (*map[string]interface{}, bool)`

GetTimeWindowOk returns a tuple with the TimeWindow field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeWindow

`func (o *PolicyOut) SetTimeWindow(v map[string]interface{})`

SetTimeWindow sets TimeWindow field to given value.

### HasTimeWindow

`func (o *PolicyOut) HasTimeWindow() bool`

HasTimeWindow returns a boolean if a field has been set.

### SetTimeWindowNil

`func (o *PolicyOut) SetTimeWindowNil(b bool)`

 SetTimeWindowNil sets the value for TimeWindow to be an explicit nil

### UnsetTimeWindow
`func (o *PolicyOut) UnsetTimeWindow()`

UnsetTimeWindow ensures that no value is present for TimeWindow, not even an explicit nil
### GetDeviceFilter

`func (o *PolicyOut) GetDeviceFilter() map[string]interface{}`

GetDeviceFilter returns the DeviceFilter field if non-nil, zero value otherwise.

### GetDeviceFilterOk

`func (o *PolicyOut) GetDeviceFilterOk() (*map[string]interface{}, bool)`

GetDeviceFilterOk returns a tuple with the DeviceFilter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeviceFilter

`func (o *PolicyOut) SetDeviceFilter(v map[string]interface{})`

SetDeviceFilter sets DeviceFilter field to given value.

### HasDeviceFilter

`func (o *PolicyOut) HasDeviceFilter() bool`

HasDeviceFilter returns a boolean if a field has been set.

### SetDeviceFilterNil

`func (o *PolicyOut) SetDeviceFilterNil(b bool)`

 SetDeviceFilterNil sets the value for DeviceFilter to be an explicit nil

### UnsetDeviceFilter
`func (o *PolicyOut) UnsetDeviceFilter()`

UnsetDeviceFilter ensures that no value is present for DeviceFilter, not even an explicit nil
### GetRoleFilter

`func (o *PolicyOut) GetRoleFilter() []string`

GetRoleFilter returns the RoleFilter field if non-nil, zero value otherwise.

### GetRoleFilterOk

`func (o *PolicyOut) GetRoleFilterOk() (*[]string, bool)`

GetRoleFilterOk returns a tuple with the RoleFilter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRoleFilter

`func (o *PolicyOut) SetRoleFilter(v []string)`

SetRoleFilter sets RoleFilter field to given value.

### HasRoleFilter

`func (o *PolicyOut) HasRoleFilter() bool`

HasRoleFilter returns a boolean if a field has been set.

### SetRoleFilterNil

`func (o *PolicyOut) SetRoleFilterNil(b bool)`

 SetRoleFilterNil sets the value for RoleFilter to be an explicit nil

### UnsetRoleFilter
`func (o *PolicyOut) UnsetRoleFilter()`

UnsetRoleFilter ensures that no value is present for RoleFilter, not even an explicit nil
### GetTemplateId

`func (o *PolicyOut) GetTemplateId() string`

GetTemplateId returns the TemplateId field if non-nil, zero value otherwise.

### GetTemplateIdOk

`func (o *PolicyOut) GetTemplateIdOk() (*string, bool)`

GetTemplateIdOk returns a tuple with the TemplateId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTemplateId

`func (o *PolicyOut) SetTemplateId(v string)`

SetTemplateId sets TemplateId field to given value.

### HasTemplateId

`func (o *PolicyOut) HasTemplateId() bool`

HasTemplateId returns a boolean if a field has been set.

### GetCreatedAt

`func (o *PolicyOut) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *PolicyOut) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *PolicyOut) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.


### GetUpdatedAt

`func (o *PolicyOut) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *PolicyOut) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *PolicyOut) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


