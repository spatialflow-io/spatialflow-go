# StartPlaybackRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SpeedMultiplier** | Pointer to **float32** |  | [optional] [default to 1.0]
**LoopEnabled** | Pointer to **bool** |  | [optional] [default to false]

## Methods

### NewStartPlaybackRequest

`func NewStartPlaybackRequest() *StartPlaybackRequest`

NewStartPlaybackRequest instantiates a new StartPlaybackRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewStartPlaybackRequestWithDefaults

`func NewStartPlaybackRequestWithDefaults() *StartPlaybackRequest`

NewStartPlaybackRequestWithDefaults instantiates a new StartPlaybackRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSpeedMultiplier

`func (o *StartPlaybackRequest) GetSpeedMultiplier() float32`

GetSpeedMultiplier returns the SpeedMultiplier field if non-nil, zero value otherwise.

### GetSpeedMultiplierOk

`func (o *StartPlaybackRequest) GetSpeedMultiplierOk() (*float32, bool)`

GetSpeedMultiplierOk returns a tuple with the SpeedMultiplier field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpeedMultiplier

`func (o *StartPlaybackRequest) SetSpeedMultiplier(v float32)`

SetSpeedMultiplier sets SpeedMultiplier field to given value.

### HasSpeedMultiplier

`func (o *StartPlaybackRequest) HasSpeedMultiplier() bool`

HasSpeedMultiplier returns a boolean if a field has been set.

### GetLoopEnabled

`func (o *StartPlaybackRequest) GetLoopEnabled() bool`

GetLoopEnabled returns the LoopEnabled field if non-nil, zero value otherwise.

### GetLoopEnabledOk

`func (o *StartPlaybackRequest) GetLoopEnabledOk() (*bool, bool)`

GetLoopEnabledOk returns a tuple with the LoopEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLoopEnabled

`func (o *StartPlaybackRequest) SetLoopEnabled(v bool)`

SetLoopEnabled sets LoopEnabled field to given value.

### HasLoopEnabled

`func (o *StartPlaybackRequest) HasLoopEnabled() bool`

HasLoopEnabled returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


