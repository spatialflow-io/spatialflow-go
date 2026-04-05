# PolicyTemplateOut

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**TemplateId** | **string** |  | 
**Name** | **string** |  | 
**Description** | **string** |  | 
**DefaultTimeWindow** | Pointer to **map[string]interface{}** |  | [optional] 

## Methods

### NewPolicyTemplateOut

`func NewPolicyTemplateOut(templateId string, name string, description string, ) *PolicyTemplateOut`

NewPolicyTemplateOut instantiates a new PolicyTemplateOut object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPolicyTemplateOutWithDefaults

`func NewPolicyTemplateOutWithDefaults() *PolicyTemplateOut`

NewPolicyTemplateOutWithDefaults instantiates a new PolicyTemplateOut object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTemplateId

`func (o *PolicyTemplateOut) GetTemplateId() string`

GetTemplateId returns the TemplateId field if non-nil, zero value otherwise.

### GetTemplateIdOk

`func (o *PolicyTemplateOut) GetTemplateIdOk() (*string, bool)`

GetTemplateIdOk returns a tuple with the TemplateId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTemplateId

`func (o *PolicyTemplateOut) SetTemplateId(v string)`

SetTemplateId sets TemplateId field to given value.


### GetName

`func (o *PolicyTemplateOut) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *PolicyTemplateOut) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *PolicyTemplateOut) SetName(v string)`

SetName sets Name field to given value.


### GetDescription

`func (o *PolicyTemplateOut) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *PolicyTemplateOut) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *PolicyTemplateOut) SetDescription(v string)`

SetDescription sets Description field to given value.


### GetDefaultTimeWindow

`func (o *PolicyTemplateOut) GetDefaultTimeWindow() map[string]interface{}`

GetDefaultTimeWindow returns the DefaultTimeWindow field if non-nil, zero value otherwise.

### GetDefaultTimeWindowOk

`func (o *PolicyTemplateOut) GetDefaultTimeWindowOk() (*map[string]interface{}, bool)`

GetDefaultTimeWindowOk returns a tuple with the DefaultTimeWindow field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultTimeWindow

`func (o *PolicyTemplateOut) SetDefaultTimeWindow(v map[string]interface{})`

SetDefaultTimeWindow sets DefaultTimeWindow field to given value.

### HasDefaultTimeWindow

`func (o *PolicyTemplateOut) HasDefaultTimeWindow() bool`

HasDefaultTimeWindow returns a boolean if a field has been set.

### SetDefaultTimeWindowNil

`func (o *PolicyTemplateOut) SetDefaultTimeWindowNil(b bool)`

 SetDefaultTimeWindowNil sets the value for DefaultTimeWindow to be an explicit nil

### UnsetDefaultTimeWindow
`func (o *PolicyTemplateOut) UnsetDefaultTimeWindow()`

UnsetDefaultTimeWindow ensures that no value is present for DefaultTimeWindow, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


