# CreateSimulationRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** |  | 
**Description** | Pointer to **string** |  | [optional] [default to ""]
**SpeedMultiplier** | Pointer to **float32** |  | [optional] [default to 1.0]
**LoopEnabled** | Pointer to **bool** |  | [optional] [default to false]
**MuteExternalActions** | Pointer to **bool** |  | [optional] [default to true]

## Methods

### NewCreateSimulationRequest

`func NewCreateSimulationRequest(name string, ) *CreateSimulationRequest`

NewCreateSimulationRequest instantiates a new CreateSimulationRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateSimulationRequestWithDefaults

`func NewCreateSimulationRequestWithDefaults() *CreateSimulationRequest`

NewCreateSimulationRequestWithDefaults instantiates a new CreateSimulationRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *CreateSimulationRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CreateSimulationRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CreateSimulationRequest) SetName(v string)`

SetName sets Name field to given value.


### GetDescription

`func (o *CreateSimulationRequest) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *CreateSimulationRequest) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *CreateSimulationRequest) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *CreateSimulationRequest) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetSpeedMultiplier

`func (o *CreateSimulationRequest) GetSpeedMultiplier() float32`

GetSpeedMultiplier returns the SpeedMultiplier field if non-nil, zero value otherwise.

### GetSpeedMultiplierOk

`func (o *CreateSimulationRequest) GetSpeedMultiplierOk() (*float32, bool)`

GetSpeedMultiplierOk returns a tuple with the SpeedMultiplier field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpeedMultiplier

`func (o *CreateSimulationRequest) SetSpeedMultiplier(v float32)`

SetSpeedMultiplier sets SpeedMultiplier field to given value.

### HasSpeedMultiplier

`func (o *CreateSimulationRequest) HasSpeedMultiplier() bool`

HasSpeedMultiplier returns a boolean if a field has been set.

### GetLoopEnabled

`func (o *CreateSimulationRequest) GetLoopEnabled() bool`

GetLoopEnabled returns the LoopEnabled field if non-nil, zero value otherwise.

### GetLoopEnabledOk

`func (o *CreateSimulationRequest) GetLoopEnabledOk() (*bool, bool)`

GetLoopEnabledOk returns a tuple with the LoopEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLoopEnabled

`func (o *CreateSimulationRequest) SetLoopEnabled(v bool)`

SetLoopEnabled sets LoopEnabled field to given value.

### HasLoopEnabled

`func (o *CreateSimulationRequest) HasLoopEnabled() bool`

HasLoopEnabled returns a boolean if a field has been set.

### GetMuteExternalActions

`func (o *CreateSimulationRequest) GetMuteExternalActions() bool`

GetMuteExternalActions returns the MuteExternalActions field if non-nil, zero value otherwise.

### GetMuteExternalActionsOk

`func (o *CreateSimulationRequest) GetMuteExternalActionsOk() (*bool, bool)`

GetMuteExternalActionsOk returns a tuple with the MuteExternalActions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMuteExternalActions

`func (o *CreateSimulationRequest) SetMuteExternalActions(v bool)`

SetMuteExternalActions sets MuteExternalActions field to given value.

### HasMuteExternalActions

`func (o *CreateSimulationRequest) HasMuteExternalActions() bool`

HasMuteExternalActions returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


