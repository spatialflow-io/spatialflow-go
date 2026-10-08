# ActionDeliveryOut

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**StepIndex** | **int32** |  | 
**NodeId** | Pointer to **NullableString** |  | [optional] 
**StepType** | **string** |  | 
**StepName** | **string** |  | 
**StepStatus** | **string** |  | 
**DeliveryStatus** | Pointer to **NullableString** |  | [optional] 
**ErrorMessage** | Pointer to **NullableString** |  | [optional] 
**ConditionPassed** | Pointer to **NullableBool** |  | [optional] 

## Methods

### NewActionDeliveryOut

`func NewActionDeliveryOut(stepIndex int32, stepType string, stepName string, stepStatus string, ) *ActionDeliveryOut`

NewActionDeliveryOut instantiates a new ActionDeliveryOut object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewActionDeliveryOutWithDefaults

`func NewActionDeliveryOutWithDefaults() *ActionDeliveryOut`

NewActionDeliveryOutWithDefaults instantiates a new ActionDeliveryOut object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetStepIndex

`func (o *ActionDeliveryOut) GetStepIndex() int32`

GetStepIndex returns the StepIndex field if non-nil, zero value otherwise.

### GetStepIndexOk

`func (o *ActionDeliveryOut) GetStepIndexOk() (*int32, bool)`

GetStepIndexOk returns a tuple with the StepIndex field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStepIndex

`func (o *ActionDeliveryOut) SetStepIndex(v int32)`

SetStepIndex sets StepIndex field to given value.


### GetNodeId

`func (o *ActionDeliveryOut) GetNodeId() string`

GetNodeId returns the NodeId field if non-nil, zero value otherwise.

### GetNodeIdOk

`func (o *ActionDeliveryOut) GetNodeIdOk() (*string, bool)`

GetNodeIdOk returns a tuple with the NodeId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNodeId

`func (o *ActionDeliveryOut) SetNodeId(v string)`

SetNodeId sets NodeId field to given value.

### HasNodeId

`func (o *ActionDeliveryOut) HasNodeId() bool`

HasNodeId returns a boolean if a field has been set.

### SetNodeIdNil

`func (o *ActionDeliveryOut) SetNodeIdNil(b bool)`

 SetNodeIdNil sets the value for NodeId to be an explicit nil

### UnsetNodeId
`func (o *ActionDeliveryOut) UnsetNodeId()`

UnsetNodeId ensures that no value is present for NodeId, not even an explicit nil
### GetStepType

`func (o *ActionDeliveryOut) GetStepType() string`

GetStepType returns the StepType field if non-nil, zero value otherwise.

### GetStepTypeOk

`func (o *ActionDeliveryOut) GetStepTypeOk() (*string, bool)`

GetStepTypeOk returns a tuple with the StepType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStepType

`func (o *ActionDeliveryOut) SetStepType(v string)`

SetStepType sets StepType field to given value.


### GetStepName

`func (o *ActionDeliveryOut) GetStepName() string`

GetStepName returns the StepName field if non-nil, zero value otherwise.

### GetStepNameOk

`func (o *ActionDeliveryOut) GetStepNameOk() (*string, bool)`

GetStepNameOk returns a tuple with the StepName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStepName

`func (o *ActionDeliveryOut) SetStepName(v string)`

SetStepName sets StepName field to given value.


### GetStepStatus

`func (o *ActionDeliveryOut) GetStepStatus() string`

GetStepStatus returns the StepStatus field if non-nil, zero value otherwise.

### GetStepStatusOk

`func (o *ActionDeliveryOut) GetStepStatusOk() (*string, bool)`

GetStepStatusOk returns a tuple with the StepStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStepStatus

`func (o *ActionDeliveryOut) SetStepStatus(v string)`

SetStepStatus sets StepStatus field to given value.


### GetDeliveryStatus

`func (o *ActionDeliveryOut) GetDeliveryStatus() string`

GetDeliveryStatus returns the DeliveryStatus field if non-nil, zero value otherwise.

### GetDeliveryStatusOk

`func (o *ActionDeliveryOut) GetDeliveryStatusOk() (*string, bool)`

GetDeliveryStatusOk returns a tuple with the DeliveryStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeliveryStatus

`func (o *ActionDeliveryOut) SetDeliveryStatus(v string)`

SetDeliveryStatus sets DeliveryStatus field to given value.

### HasDeliveryStatus

`func (o *ActionDeliveryOut) HasDeliveryStatus() bool`

HasDeliveryStatus returns a boolean if a field has been set.

### SetDeliveryStatusNil

`func (o *ActionDeliveryOut) SetDeliveryStatusNil(b bool)`

 SetDeliveryStatusNil sets the value for DeliveryStatus to be an explicit nil

### UnsetDeliveryStatus
`func (o *ActionDeliveryOut) UnsetDeliveryStatus()`

UnsetDeliveryStatus ensures that no value is present for DeliveryStatus, not even an explicit nil
### GetErrorMessage

`func (o *ActionDeliveryOut) GetErrorMessage() string`

GetErrorMessage returns the ErrorMessage field if non-nil, zero value otherwise.

### GetErrorMessageOk

`func (o *ActionDeliveryOut) GetErrorMessageOk() (*string, bool)`

GetErrorMessageOk returns a tuple with the ErrorMessage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrorMessage

`func (o *ActionDeliveryOut) SetErrorMessage(v string)`

SetErrorMessage sets ErrorMessage field to given value.

### HasErrorMessage

`func (o *ActionDeliveryOut) HasErrorMessage() bool`

HasErrorMessage returns a boolean if a field has been set.

### SetErrorMessageNil

`func (o *ActionDeliveryOut) SetErrorMessageNil(b bool)`

 SetErrorMessageNil sets the value for ErrorMessage to be an explicit nil

### UnsetErrorMessage
`func (o *ActionDeliveryOut) UnsetErrorMessage()`

UnsetErrorMessage ensures that no value is present for ErrorMessage, not even an explicit nil
### GetConditionPassed

`func (o *ActionDeliveryOut) GetConditionPassed() bool`

GetConditionPassed returns the ConditionPassed field if non-nil, zero value otherwise.

### GetConditionPassedOk

`func (o *ActionDeliveryOut) GetConditionPassedOk() (*bool, bool)`

GetConditionPassedOk returns a tuple with the ConditionPassed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConditionPassed

`func (o *ActionDeliveryOut) SetConditionPassed(v bool)`

SetConditionPassed sets ConditionPassed field to given value.

### HasConditionPassed

`func (o *ActionDeliveryOut) HasConditionPassed() bool`

HasConditionPassed returns a boolean if a field has been set.

### SetConditionPassedNil

`func (o *ActionDeliveryOut) SetConditionPassedNil(b bool)`

 SetConditionPassedNil sets the value for ConditionPassed to be an explicit nil

### UnsetConditionPassed
`func (o *ActionDeliveryOut) UnsetConditionPassed()`

UnsetConditionPassed ensures that no value is present for ConditionPassed, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


