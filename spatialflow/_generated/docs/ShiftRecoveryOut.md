# ShiftRecoveryOut

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DeviceUuid** | **string** |  | 
**DeviceId** | **string** |  | 
**SourceSessionId** | **string** |  | 
**Outcome** | **string** |  | 
**Reason** | Pointer to **NullableString** |  | [optional] 
**Replacement** | Pointer to [**NullableShiftRecoveryReplacementOut**](ShiftRecoveryReplacementOut.md) |  | [optional] 
**Current** | [**ShiftRecoveryCurrentOut**](ShiftRecoveryCurrentOut.md) |  | 

## Methods

### NewShiftRecoveryOut

`func NewShiftRecoveryOut(deviceUuid string, deviceId string, sourceSessionId string, outcome string, current ShiftRecoveryCurrentOut, ) *ShiftRecoveryOut`

NewShiftRecoveryOut instantiates a new ShiftRecoveryOut object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewShiftRecoveryOutWithDefaults

`func NewShiftRecoveryOutWithDefaults() *ShiftRecoveryOut`

NewShiftRecoveryOutWithDefaults instantiates a new ShiftRecoveryOut object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDeviceUuid

`func (o *ShiftRecoveryOut) GetDeviceUuid() string`

GetDeviceUuid returns the DeviceUuid field if non-nil, zero value otherwise.

### GetDeviceUuidOk

`func (o *ShiftRecoveryOut) GetDeviceUuidOk() (*string, bool)`

GetDeviceUuidOk returns a tuple with the DeviceUuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeviceUuid

`func (o *ShiftRecoveryOut) SetDeviceUuid(v string)`

SetDeviceUuid sets DeviceUuid field to given value.


### GetDeviceId

`func (o *ShiftRecoveryOut) GetDeviceId() string`

GetDeviceId returns the DeviceId field if non-nil, zero value otherwise.

### GetDeviceIdOk

`func (o *ShiftRecoveryOut) GetDeviceIdOk() (*string, bool)`

GetDeviceIdOk returns a tuple with the DeviceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeviceId

`func (o *ShiftRecoveryOut) SetDeviceId(v string)`

SetDeviceId sets DeviceId field to given value.


### GetSourceSessionId

`func (o *ShiftRecoveryOut) GetSourceSessionId() string`

GetSourceSessionId returns the SourceSessionId field if non-nil, zero value otherwise.

### GetSourceSessionIdOk

`func (o *ShiftRecoveryOut) GetSourceSessionIdOk() (*string, bool)`

GetSourceSessionIdOk returns a tuple with the SourceSessionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceSessionId

`func (o *ShiftRecoveryOut) SetSourceSessionId(v string)`

SetSourceSessionId sets SourceSessionId field to given value.


### GetOutcome

`func (o *ShiftRecoveryOut) GetOutcome() string`

GetOutcome returns the Outcome field if non-nil, zero value otherwise.

### GetOutcomeOk

`func (o *ShiftRecoveryOut) GetOutcomeOk() (*string, bool)`

GetOutcomeOk returns a tuple with the Outcome field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOutcome

`func (o *ShiftRecoveryOut) SetOutcome(v string)`

SetOutcome sets Outcome field to given value.


### GetReason

`func (o *ShiftRecoveryOut) GetReason() string`

GetReason returns the Reason field if non-nil, zero value otherwise.

### GetReasonOk

`func (o *ShiftRecoveryOut) GetReasonOk() (*string, bool)`

GetReasonOk returns a tuple with the Reason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReason

`func (o *ShiftRecoveryOut) SetReason(v string)`

SetReason sets Reason field to given value.

### HasReason

`func (o *ShiftRecoveryOut) HasReason() bool`

HasReason returns a boolean if a field has been set.

### SetReasonNil

`func (o *ShiftRecoveryOut) SetReasonNil(b bool)`

 SetReasonNil sets the value for Reason to be an explicit nil

### UnsetReason
`func (o *ShiftRecoveryOut) UnsetReason()`

UnsetReason ensures that no value is present for Reason, not even an explicit nil
### GetReplacement

`func (o *ShiftRecoveryOut) GetReplacement() ShiftRecoveryReplacementOut`

GetReplacement returns the Replacement field if non-nil, zero value otherwise.

### GetReplacementOk

`func (o *ShiftRecoveryOut) GetReplacementOk() (*ShiftRecoveryReplacementOut, bool)`

GetReplacementOk returns a tuple with the Replacement field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReplacement

`func (o *ShiftRecoveryOut) SetReplacement(v ShiftRecoveryReplacementOut)`

SetReplacement sets Replacement field to given value.

### HasReplacement

`func (o *ShiftRecoveryOut) HasReplacement() bool`

HasReplacement returns a boolean if a field has been set.

### SetReplacementNil

`func (o *ShiftRecoveryOut) SetReplacementNil(b bool)`

 SetReplacementNil sets the value for Replacement to be an explicit nil

### UnsetReplacement
`func (o *ShiftRecoveryOut) UnsetReplacement()`

UnsetReplacement ensures that no value is present for Replacement, not even an explicit nil
### GetCurrent

`func (o *ShiftRecoveryOut) GetCurrent() ShiftRecoveryCurrentOut`

GetCurrent returns the Current field if non-nil, zero value otherwise.

### GetCurrentOk

`func (o *ShiftRecoveryOut) GetCurrentOk() (*ShiftRecoveryCurrentOut, bool)`

GetCurrentOk returns a tuple with the Current field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrent

`func (o *ShiftRecoveryOut) SetCurrent(v ShiftRecoveryCurrentOut)`

SetCurrent sets Current field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


