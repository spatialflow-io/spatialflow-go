# PolicyUpdateIn

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **NullableString** |  | [optional] 
**Description** | Pointer to **NullableString** |  | [optional] 
**IsEnabled** | Pointer to **NullableBool** |  | [optional] 
**TimeWindow** | Pointer to **map[string]interface{}** |  | [optional] 
**DeviceFilter** | Pointer to **map[string]interface{}** |  | [optional] 
**RoleFilter** | Pointer to **[]string** |  | [optional] 

## Methods

### NewPolicyUpdateIn

`func NewPolicyUpdateIn() *PolicyUpdateIn`

NewPolicyUpdateIn instantiates a new PolicyUpdateIn object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPolicyUpdateInWithDefaults

`func NewPolicyUpdateInWithDefaults() *PolicyUpdateIn`

NewPolicyUpdateInWithDefaults instantiates a new PolicyUpdateIn object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *PolicyUpdateIn) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *PolicyUpdateIn) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *PolicyUpdateIn) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *PolicyUpdateIn) HasName() bool`

HasName returns a boolean if a field has been set.

### SetNameNil

`func (o *PolicyUpdateIn) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *PolicyUpdateIn) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetDescription

`func (o *PolicyUpdateIn) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *PolicyUpdateIn) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *PolicyUpdateIn) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *PolicyUpdateIn) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescriptionNil

`func (o *PolicyUpdateIn) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *PolicyUpdateIn) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetIsEnabled

`func (o *PolicyUpdateIn) GetIsEnabled() bool`

GetIsEnabled returns the IsEnabled field if non-nil, zero value otherwise.

### GetIsEnabledOk

`func (o *PolicyUpdateIn) GetIsEnabledOk() (*bool, bool)`

GetIsEnabledOk returns a tuple with the IsEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsEnabled

`func (o *PolicyUpdateIn) SetIsEnabled(v bool)`

SetIsEnabled sets IsEnabled field to given value.

### HasIsEnabled

`func (o *PolicyUpdateIn) HasIsEnabled() bool`

HasIsEnabled returns a boolean if a field has been set.

### SetIsEnabledNil

`func (o *PolicyUpdateIn) SetIsEnabledNil(b bool)`

 SetIsEnabledNil sets the value for IsEnabled to be an explicit nil

### UnsetIsEnabled
`func (o *PolicyUpdateIn) UnsetIsEnabled()`

UnsetIsEnabled ensures that no value is present for IsEnabled, not even an explicit nil
### GetTimeWindow

`func (o *PolicyUpdateIn) GetTimeWindow() map[string]interface{}`

GetTimeWindow returns the TimeWindow field if non-nil, zero value otherwise.

### GetTimeWindowOk

`func (o *PolicyUpdateIn) GetTimeWindowOk() (*map[string]interface{}, bool)`

GetTimeWindowOk returns a tuple with the TimeWindow field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeWindow

`func (o *PolicyUpdateIn) SetTimeWindow(v map[string]interface{})`

SetTimeWindow sets TimeWindow field to given value.

### HasTimeWindow

`func (o *PolicyUpdateIn) HasTimeWindow() bool`

HasTimeWindow returns a boolean if a field has been set.

### SetTimeWindowNil

`func (o *PolicyUpdateIn) SetTimeWindowNil(b bool)`

 SetTimeWindowNil sets the value for TimeWindow to be an explicit nil

### UnsetTimeWindow
`func (o *PolicyUpdateIn) UnsetTimeWindow()`

UnsetTimeWindow ensures that no value is present for TimeWindow, not even an explicit nil
### GetDeviceFilter

`func (o *PolicyUpdateIn) GetDeviceFilter() map[string]interface{}`

GetDeviceFilter returns the DeviceFilter field if non-nil, zero value otherwise.

### GetDeviceFilterOk

`func (o *PolicyUpdateIn) GetDeviceFilterOk() (*map[string]interface{}, bool)`

GetDeviceFilterOk returns a tuple with the DeviceFilter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeviceFilter

`func (o *PolicyUpdateIn) SetDeviceFilter(v map[string]interface{})`

SetDeviceFilter sets DeviceFilter field to given value.

### HasDeviceFilter

`func (o *PolicyUpdateIn) HasDeviceFilter() bool`

HasDeviceFilter returns a boolean if a field has been set.

### SetDeviceFilterNil

`func (o *PolicyUpdateIn) SetDeviceFilterNil(b bool)`

 SetDeviceFilterNil sets the value for DeviceFilter to be an explicit nil

### UnsetDeviceFilter
`func (o *PolicyUpdateIn) UnsetDeviceFilter()`

UnsetDeviceFilter ensures that no value is present for DeviceFilter, not even an explicit nil
### GetRoleFilter

`func (o *PolicyUpdateIn) GetRoleFilter() []string`

GetRoleFilter returns the RoleFilter field if non-nil, zero value otherwise.

### GetRoleFilterOk

`func (o *PolicyUpdateIn) GetRoleFilterOk() (*[]string, bool)`

GetRoleFilterOk returns a tuple with the RoleFilter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRoleFilter

`func (o *PolicyUpdateIn) SetRoleFilter(v []string)`

SetRoleFilter sets RoleFilter field to given value.

### HasRoleFilter

`func (o *PolicyUpdateIn) HasRoleFilter() bool`

HasRoleFilter returns a boolean if a field has been set.

### SetRoleFilterNil

`func (o *PolicyUpdateIn) SetRoleFilterNil(b bool)`

 SetRoleFilterNil sets the value for RoleFilter to be an explicit nil

### UnsetRoleFilter
`func (o *PolicyUpdateIn) UnsetRoleFilter()`

UnsetRoleFilter ensures that no value is present for RoleFilter, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


