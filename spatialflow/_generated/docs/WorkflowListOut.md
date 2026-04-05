# WorkflowListOut

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** |  | 
**Name** | **string** |  | 
**Description** | **NullableString** |  | 
**Status** | **string** |  | 
**CreatedAt** | **time.Time** |  | 
**UpdatedAt** | **time.Time** |  | 
**LastRun** | **NullableString** |  | 
**RunCount** | **int32** |  | 
**SuccessRate** | **float32** |  | 
**UserId** | **string** |  | 
**Version** | **int32** |  | 

## Methods

### NewWorkflowListOut

`func NewWorkflowListOut(id string, name string, description NullableString, status string, createdAt time.Time, updatedAt time.Time, lastRun NullableString, runCount int32, successRate float32, userId string, version int32, ) *WorkflowListOut`

NewWorkflowListOut instantiates a new WorkflowListOut object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWorkflowListOutWithDefaults

`func NewWorkflowListOutWithDefaults() *WorkflowListOut`

NewWorkflowListOutWithDefaults instantiates a new WorkflowListOut object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *WorkflowListOut) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *WorkflowListOut) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *WorkflowListOut) SetId(v string)`

SetId sets Id field to given value.


### GetName

`func (o *WorkflowListOut) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *WorkflowListOut) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *WorkflowListOut) SetName(v string)`

SetName sets Name field to given value.


### GetDescription

`func (o *WorkflowListOut) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *WorkflowListOut) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *WorkflowListOut) SetDescription(v string)`

SetDescription sets Description field to given value.


### SetDescriptionNil

`func (o *WorkflowListOut) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *WorkflowListOut) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetStatus

`func (o *WorkflowListOut) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *WorkflowListOut) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *WorkflowListOut) SetStatus(v string)`

SetStatus sets Status field to given value.


### GetCreatedAt

`func (o *WorkflowListOut) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *WorkflowListOut) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *WorkflowListOut) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.


### GetUpdatedAt

`func (o *WorkflowListOut) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *WorkflowListOut) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *WorkflowListOut) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.


### GetLastRun

`func (o *WorkflowListOut) GetLastRun() string`

GetLastRun returns the LastRun field if non-nil, zero value otherwise.

### GetLastRunOk

`func (o *WorkflowListOut) GetLastRunOk() (*string, bool)`

GetLastRunOk returns a tuple with the LastRun field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastRun

`func (o *WorkflowListOut) SetLastRun(v string)`

SetLastRun sets LastRun field to given value.


### SetLastRunNil

`func (o *WorkflowListOut) SetLastRunNil(b bool)`

 SetLastRunNil sets the value for LastRun to be an explicit nil

### UnsetLastRun
`func (o *WorkflowListOut) UnsetLastRun()`

UnsetLastRun ensures that no value is present for LastRun, not even an explicit nil
### GetRunCount

`func (o *WorkflowListOut) GetRunCount() int32`

GetRunCount returns the RunCount field if non-nil, zero value otherwise.

### GetRunCountOk

`func (o *WorkflowListOut) GetRunCountOk() (*int32, bool)`

GetRunCountOk returns a tuple with the RunCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRunCount

`func (o *WorkflowListOut) SetRunCount(v int32)`

SetRunCount sets RunCount field to given value.


### GetSuccessRate

`func (o *WorkflowListOut) GetSuccessRate() float32`

GetSuccessRate returns the SuccessRate field if non-nil, zero value otherwise.

### GetSuccessRateOk

`func (o *WorkflowListOut) GetSuccessRateOk() (*float32, bool)`

GetSuccessRateOk returns a tuple with the SuccessRate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSuccessRate

`func (o *WorkflowListOut) SetSuccessRate(v float32)`

SetSuccessRate sets SuccessRate field to given value.


### GetUserId

`func (o *WorkflowListOut) GetUserId() string`

GetUserId returns the UserId field if non-nil, zero value otherwise.

### GetUserIdOk

`func (o *WorkflowListOut) GetUserIdOk() (*string, bool)`

GetUserIdOk returns a tuple with the UserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserId

`func (o *WorkflowListOut) SetUserId(v string)`

SetUserId sets UserId field to given value.


### GetVersion

`func (o *WorkflowListOut) GetVersion() int32`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *WorkflowListOut) GetVersionOk() (*int32, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *WorkflowListOut) SetVersion(v int32)`

SetVersion sets Version field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


