# RecentExecutionOut

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** |  | 
**ExecutionId** | **string** |  | 
**Status** | **string** |  | 
**TriggerSource** | **string** |  | 
**StartedAt** | Pointer to **NullableString** |  | [optional] 
**DurationSeconds** | Pointer to **NullableFloat32** |  | [optional] 

## Methods

### NewRecentExecutionOut

`func NewRecentExecutionOut(id string, executionId string, status string, triggerSource string, ) *RecentExecutionOut`

NewRecentExecutionOut instantiates a new RecentExecutionOut object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRecentExecutionOutWithDefaults

`func NewRecentExecutionOutWithDefaults() *RecentExecutionOut`

NewRecentExecutionOutWithDefaults instantiates a new RecentExecutionOut object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *RecentExecutionOut) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *RecentExecutionOut) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *RecentExecutionOut) SetId(v string)`

SetId sets Id field to given value.


### GetExecutionId

`func (o *RecentExecutionOut) GetExecutionId() string`

GetExecutionId returns the ExecutionId field if non-nil, zero value otherwise.

### GetExecutionIdOk

`func (o *RecentExecutionOut) GetExecutionIdOk() (*string, bool)`

GetExecutionIdOk returns a tuple with the ExecutionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExecutionId

`func (o *RecentExecutionOut) SetExecutionId(v string)`

SetExecutionId sets ExecutionId field to given value.


### GetStatus

`func (o *RecentExecutionOut) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *RecentExecutionOut) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *RecentExecutionOut) SetStatus(v string)`

SetStatus sets Status field to given value.


### GetTriggerSource

`func (o *RecentExecutionOut) GetTriggerSource() string`

GetTriggerSource returns the TriggerSource field if non-nil, zero value otherwise.

### GetTriggerSourceOk

`func (o *RecentExecutionOut) GetTriggerSourceOk() (*string, bool)`

GetTriggerSourceOk returns a tuple with the TriggerSource field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTriggerSource

`func (o *RecentExecutionOut) SetTriggerSource(v string)`

SetTriggerSource sets TriggerSource field to given value.


### GetStartedAt

`func (o *RecentExecutionOut) GetStartedAt() string`

GetStartedAt returns the StartedAt field if non-nil, zero value otherwise.

### GetStartedAtOk

`func (o *RecentExecutionOut) GetStartedAtOk() (*string, bool)`

GetStartedAtOk returns a tuple with the StartedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartedAt

`func (o *RecentExecutionOut) SetStartedAt(v string)`

SetStartedAt sets StartedAt field to given value.

### HasStartedAt

`func (o *RecentExecutionOut) HasStartedAt() bool`

HasStartedAt returns a boolean if a field has been set.

### SetStartedAtNil

`func (o *RecentExecutionOut) SetStartedAtNil(b bool)`

 SetStartedAtNil sets the value for StartedAt to be an explicit nil

### UnsetStartedAt
`func (o *RecentExecutionOut) UnsetStartedAt()`

UnsetStartedAt ensures that no value is present for StartedAt, not even an explicit nil
### GetDurationSeconds

`func (o *RecentExecutionOut) GetDurationSeconds() float32`

GetDurationSeconds returns the DurationSeconds field if non-nil, zero value otherwise.

### GetDurationSecondsOk

`func (o *RecentExecutionOut) GetDurationSecondsOk() (*float32, bool)`

GetDurationSecondsOk returns a tuple with the DurationSeconds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDurationSeconds

`func (o *RecentExecutionOut) SetDurationSeconds(v float32)`

SetDurationSeconds sets DurationSeconds field to given value.

### HasDurationSeconds

`func (o *RecentExecutionOut) HasDurationSeconds() bool`

HasDurationSeconds returns a boolean if a field has been set.

### SetDurationSecondsNil

`func (o *RecentExecutionOut) SetDurationSecondsNil(b bool)`

 SetDurationSecondsNil sets the value for DurationSeconds to be an explicit nil

### UnsetDurationSeconds
`func (o *RecentExecutionOut) UnsetDurationSeconds()`

UnsetDurationSeconds ensures that no value is present for DurationSeconds, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


