# PrivacyErasureResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**JobId** | **string** |  | 
**Status** | **string** |  | 
**EstimatedCount** | **int32** |  | 
**DeletedCount** | **int32** |  | 
**StartedAt** | Pointer to **NullableTime** |  | [optional] 
**CompletedAt** | Pointer to **NullableTime** |  | [optional] 
**DryRun** | **bool** |  | 

## Methods

### NewPrivacyErasureResponse

`func NewPrivacyErasureResponse(jobId string, status string, estimatedCount int32, deletedCount int32, dryRun bool, ) *PrivacyErasureResponse`

NewPrivacyErasureResponse instantiates a new PrivacyErasureResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPrivacyErasureResponseWithDefaults

`func NewPrivacyErasureResponseWithDefaults() *PrivacyErasureResponse`

NewPrivacyErasureResponseWithDefaults instantiates a new PrivacyErasureResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetJobId

`func (o *PrivacyErasureResponse) GetJobId() string`

GetJobId returns the JobId field if non-nil, zero value otherwise.

### GetJobIdOk

`func (o *PrivacyErasureResponse) GetJobIdOk() (*string, bool)`

GetJobIdOk returns a tuple with the JobId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJobId

`func (o *PrivacyErasureResponse) SetJobId(v string)`

SetJobId sets JobId field to given value.


### GetStatus

`func (o *PrivacyErasureResponse) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *PrivacyErasureResponse) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *PrivacyErasureResponse) SetStatus(v string)`

SetStatus sets Status field to given value.


### GetEstimatedCount

`func (o *PrivacyErasureResponse) GetEstimatedCount() int32`

GetEstimatedCount returns the EstimatedCount field if non-nil, zero value otherwise.

### GetEstimatedCountOk

`func (o *PrivacyErasureResponse) GetEstimatedCountOk() (*int32, bool)`

GetEstimatedCountOk returns a tuple with the EstimatedCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEstimatedCount

`func (o *PrivacyErasureResponse) SetEstimatedCount(v int32)`

SetEstimatedCount sets EstimatedCount field to given value.


### GetDeletedCount

`func (o *PrivacyErasureResponse) GetDeletedCount() int32`

GetDeletedCount returns the DeletedCount field if non-nil, zero value otherwise.

### GetDeletedCountOk

`func (o *PrivacyErasureResponse) GetDeletedCountOk() (*int32, bool)`

GetDeletedCountOk returns a tuple with the DeletedCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeletedCount

`func (o *PrivacyErasureResponse) SetDeletedCount(v int32)`

SetDeletedCount sets DeletedCount field to given value.


### GetStartedAt

`func (o *PrivacyErasureResponse) GetStartedAt() time.Time`

GetStartedAt returns the StartedAt field if non-nil, zero value otherwise.

### GetStartedAtOk

`func (o *PrivacyErasureResponse) GetStartedAtOk() (*time.Time, bool)`

GetStartedAtOk returns a tuple with the StartedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartedAt

`func (o *PrivacyErasureResponse) SetStartedAt(v time.Time)`

SetStartedAt sets StartedAt field to given value.

### HasStartedAt

`func (o *PrivacyErasureResponse) HasStartedAt() bool`

HasStartedAt returns a boolean if a field has been set.

### SetStartedAtNil

`func (o *PrivacyErasureResponse) SetStartedAtNil(b bool)`

 SetStartedAtNil sets the value for StartedAt to be an explicit nil

### UnsetStartedAt
`func (o *PrivacyErasureResponse) UnsetStartedAt()`

UnsetStartedAt ensures that no value is present for StartedAt, not even an explicit nil
### GetCompletedAt

`func (o *PrivacyErasureResponse) GetCompletedAt() time.Time`

GetCompletedAt returns the CompletedAt field if non-nil, zero value otherwise.

### GetCompletedAtOk

`func (o *PrivacyErasureResponse) GetCompletedAtOk() (*time.Time, bool)`

GetCompletedAtOk returns a tuple with the CompletedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompletedAt

`func (o *PrivacyErasureResponse) SetCompletedAt(v time.Time)`

SetCompletedAt sets CompletedAt field to given value.

### HasCompletedAt

`func (o *PrivacyErasureResponse) HasCompletedAt() bool`

HasCompletedAt returns a boolean if a field has been set.

### SetCompletedAtNil

`func (o *PrivacyErasureResponse) SetCompletedAtNil(b bool)`

 SetCompletedAtNil sets the value for CompletedAt to be an explicit nil

### UnsetCompletedAt
`func (o *PrivacyErasureResponse) UnsetCompletedAt()`

UnsetCompletedAt ensures that no value is present for CompletedAt, not even an explicit nil
### GetDryRun

`func (o *PrivacyErasureResponse) GetDryRun() bool`

GetDryRun returns the DryRun field if non-nil, zero value otherwise.

### GetDryRunOk

`func (o *PrivacyErasureResponse) GetDryRunOk() (*bool, bool)`

GetDryRunOk returns a tuple with the DryRun field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDryRun

`func (o *PrivacyErasureResponse) SetDryRun(v bool)`

SetDryRun sets DryRun field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


