# PolicyCreateIn

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** |  | 
**GeofenceId** | **string** |  | 
**Description** | Pointer to **NullableString** |  | [optional] 
**IsEnabled** | Pointer to **bool** |  | [optional] [default to true]
**TimeWindow** | Pointer to **map[string]interface{}** |  | [optional] 
**DeviceFilter** | Pointer to **map[string]interface{}** |  | [optional] 
**RoleFilter** | Pointer to **[]string** |  | [optional] 
**TemplateId** | Pointer to **NullableString** |  | [optional] 

## Methods

### NewPolicyCreateIn

`func NewPolicyCreateIn(name string, geofenceId string, ) *PolicyCreateIn`

NewPolicyCreateIn instantiates a new PolicyCreateIn object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPolicyCreateInWithDefaults

`func NewPolicyCreateInWithDefaults() *PolicyCreateIn`

NewPolicyCreateInWithDefaults instantiates a new PolicyCreateIn object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *PolicyCreateIn) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *PolicyCreateIn) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *PolicyCreateIn) SetName(v string)`

SetName sets Name field to given value.


### GetGeofenceId

`func (o *PolicyCreateIn) GetGeofenceId() string`

GetGeofenceId returns the GeofenceId field if non-nil, zero value otherwise.

### GetGeofenceIdOk

`func (o *PolicyCreateIn) GetGeofenceIdOk() (*string, bool)`

GetGeofenceIdOk returns a tuple with the GeofenceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGeofenceId

`func (o *PolicyCreateIn) SetGeofenceId(v string)`

SetGeofenceId sets GeofenceId field to given value.


### GetDescription

`func (o *PolicyCreateIn) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *PolicyCreateIn) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *PolicyCreateIn) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *PolicyCreateIn) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescriptionNil

`func (o *PolicyCreateIn) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *PolicyCreateIn) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetIsEnabled

`func (o *PolicyCreateIn) GetIsEnabled() bool`

GetIsEnabled returns the IsEnabled field if non-nil, zero value otherwise.

### GetIsEnabledOk

`func (o *PolicyCreateIn) GetIsEnabledOk() (*bool, bool)`

GetIsEnabledOk returns a tuple with the IsEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsEnabled

`func (o *PolicyCreateIn) SetIsEnabled(v bool)`

SetIsEnabled sets IsEnabled field to given value.

### HasIsEnabled

`func (o *PolicyCreateIn) HasIsEnabled() bool`

HasIsEnabled returns a boolean if a field has been set.

### GetTimeWindow

`func (o *PolicyCreateIn) GetTimeWindow() map[string]interface{}`

GetTimeWindow returns the TimeWindow field if non-nil, zero value otherwise.

### GetTimeWindowOk

`func (o *PolicyCreateIn) GetTimeWindowOk() (*map[string]interface{}, bool)`

GetTimeWindowOk returns a tuple with the TimeWindow field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeWindow

`func (o *PolicyCreateIn) SetTimeWindow(v map[string]interface{})`

SetTimeWindow sets TimeWindow field to given value.

### HasTimeWindow

`func (o *PolicyCreateIn) HasTimeWindow() bool`

HasTimeWindow returns a boolean if a field has been set.

### SetTimeWindowNil

`func (o *PolicyCreateIn) SetTimeWindowNil(b bool)`

 SetTimeWindowNil sets the value for TimeWindow to be an explicit nil

### UnsetTimeWindow
`func (o *PolicyCreateIn) UnsetTimeWindow()`

UnsetTimeWindow ensures that no value is present for TimeWindow, not even an explicit nil
### GetDeviceFilter

`func (o *PolicyCreateIn) GetDeviceFilter() map[string]interface{}`

GetDeviceFilter returns the DeviceFilter field if non-nil, zero value otherwise.

### GetDeviceFilterOk

`func (o *PolicyCreateIn) GetDeviceFilterOk() (*map[string]interface{}, bool)`

GetDeviceFilterOk returns a tuple with the DeviceFilter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeviceFilter

`func (o *PolicyCreateIn) SetDeviceFilter(v map[string]interface{})`

SetDeviceFilter sets DeviceFilter field to given value.

### HasDeviceFilter

`func (o *PolicyCreateIn) HasDeviceFilter() bool`

HasDeviceFilter returns a boolean if a field has been set.

### SetDeviceFilterNil

`func (o *PolicyCreateIn) SetDeviceFilterNil(b bool)`

 SetDeviceFilterNil sets the value for DeviceFilter to be an explicit nil

### UnsetDeviceFilter
`func (o *PolicyCreateIn) UnsetDeviceFilter()`

UnsetDeviceFilter ensures that no value is present for DeviceFilter, not even an explicit nil
### GetRoleFilter

`func (o *PolicyCreateIn) GetRoleFilter() []string`

GetRoleFilter returns the RoleFilter field if non-nil, zero value otherwise.

### GetRoleFilterOk

`func (o *PolicyCreateIn) GetRoleFilterOk() (*[]string, bool)`

GetRoleFilterOk returns a tuple with the RoleFilter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRoleFilter

`func (o *PolicyCreateIn) SetRoleFilter(v []string)`

SetRoleFilter sets RoleFilter field to given value.

### HasRoleFilter

`func (o *PolicyCreateIn) HasRoleFilter() bool`

HasRoleFilter returns a boolean if a field has been set.

### SetRoleFilterNil

`func (o *PolicyCreateIn) SetRoleFilterNil(b bool)`

 SetRoleFilterNil sets the value for RoleFilter to be an explicit nil

### UnsetRoleFilter
`func (o *PolicyCreateIn) UnsetRoleFilter()`

UnsetRoleFilter ensures that no value is present for RoleFilter, not even an explicit nil
### GetTemplateId

`func (o *PolicyCreateIn) GetTemplateId() string`

GetTemplateId returns the TemplateId field if non-nil, zero value otherwise.

### GetTemplateIdOk

`func (o *PolicyCreateIn) GetTemplateIdOk() (*string, bool)`

GetTemplateIdOk returns a tuple with the TemplateId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTemplateId

`func (o *PolicyCreateIn) SetTemplateId(v string)`

SetTemplateId sets TemplateId field to given value.

### HasTemplateId

`func (o *PolicyCreateIn) HasTemplateId() bool`

HasTemplateId returns a boolean if a field has been set.

### SetTemplateIdNil

`func (o *PolicyCreateIn) SetTemplateIdNil(b bool)`

 SetTemplateIdNil sets the value for TemplateId to be an explicit nil

### UnsetTemplateId
`func (o *PolicyCreateIn) UnsetTemplateId()`

UnsetTemplateId ensures that no value is present for TemplateId, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


