# UploadJobStatus

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**JobId** | **string** |  | 
**Status** | **string** |  | 
**FileName** | **string** |  | 
**FileFormat** | **string** |  | 
**CreatedAt** | **time.Time** |  | 
**StartedAt** | **NullableTime** |  | 
**CompletedAt** | **NullableTime** |  | 
**Duration** | Pointer to **NullableFloat32** |  | [optional] 
**TotalFeatures** | **int32** |  | 
**CreatedCount** | **int32** |  | 
**FailedCount** | **int32** |  | 
**Results** | Pointer to **map[string]interface{}** |  | [optional] 
**ErrorMessage** | Pointer to **NullableString** |  | [optional] 

## Methods

### NewUploadJobStatus

`func NewUploadJobStatus(jobId string, status string, fileName string, fileFormat string, createdAt time.Time, startedAt NullableTime, completedAt NullableTime, totalFeatures int32, createdCount int32, failedCount int32, ) *UploadJobStatus`

NewUploadJobStatus instantiates a new UploadJobStatus object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUploadJobStatusWithDefaults

`func NewUploadJobStatusWithDefaults() *UploadJobStatus`

NewUploadJobStatusWithDefaults instantiates a new UploadJobStatus object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetJobId

`func (o *UploadJobStatus) GetJobId() string`

GetJobId returns the JobId field if non-nil, zero value otherwise.

### GetJobIdOk

`func (o *UploadJobStatus) GetJobIdOk() (*string, bool)`

GetJobIdOk returns a tuple with the JobId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJobId

`func (o *UploadJobStatus) SetJobId(v string)`

SetJobId sets JobId field to given value.


### GetStatus

`func (o *UploadJobStatus) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *UploadJobStatus) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *UploadJobStatus) SetStatus(v string)`

SetStatus sets Status field to given value.


### GetFileName

`func (o *UploadJobStatus) GetFileName() string`

GetFileName returns the FileName field if non-nil, zero value otherwise.

### GetFileNameOk

`func (o *UploadJobStatus) GetFileNameOk() (*string, bool)`

GetFileNameOk returns a tuple with the FileName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFileName

`func (o *UploadJobStatus) SetFileName(v string)`

SetFileName sets FileName field to given value.


### GetFileFormat

`func (o *UploadJobStatus) GetFileFormat() string`

GetFileFormat returns the FileFormat field if non-nil, zero value otherwise.

### GetFileFormatOk

`func (o *UploadJobStatus) GetFileFormatOk() (*string, bool)`

GetFileFormatOk returns a tuple with the FileFormat field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFileFormat

`func (o *UploadJobStatus) SetFileFormat(v string)`

SetFileFormat sets FileFormat field to given value.


### GetCreatedAt

`func (o *UploadJobStatus) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *UploadJobStatus) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *UploadJobStatus) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.


### GetStartedAt

`func (o *UploadJobStatus) GetStartedAt() time.Time`

GetStartedAt returns the StartedAt field if non-nil, zero value otherwise.

### GetStartedAtOk

`func (o *UploadJobStatus) GetStartedAtOk() (*time.Time, bool)`

GetStartedAtOk returns a tuple with the StartedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartedAt

`func (o *UploadJobStatus) SetStartedAt(v time.Time)`

SetStartedAt sets StartedAt field to given value.


### SetStartedAtNil

`func (o *UploadJobStatus) SetStartedAtNil(b bool)`

 SetStartedAtNil sets the value for StartedAt to be an explicit nil

### UnsetStartedAt
`func (o *UploadJobStatus) UnsetStartedAt()`

UnsetStartedAt ensures that no value is present for StartedAt, not even an explicit nil
### GetCompletedAt

`func (o *UploadJobStatus) GetCompletedAt() time.Time`

GetCompletedAt returns the CompletedAt field if non-nil, zero value otherwise.

### GetCompletedAtOk

`func (o *UploadJobStatus) GetCompletedAtOk() (*time.Time, bool)`

GetCompletedAtOk returns a tuple with the CompletedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompletedAt

`func (o *UploadJobStatus) SetCompletedAt(v time.Time)`

SetCompletedAt sets CompletedAt field to given value.


### SetCompletedAtNil

`func (o *UploadJobStatus) SetCompletedAtNil(b bool)`

 SetCompletedAtNil sets the value for CompletedAt to be an explicit nil

### UnsetCompletedAt
`func (o *UploadJobStatus) UnsetCompletedAt()`

UnsetCompletedAt ensures that no value is present for CompletedAt, not even an explicit nil
### GetDuration

`func (o *UploadJobStatus) GetDuration() float32`

GetDuration returns the Duration field if non-nil, zero value otherwise.

### GetDurationOk

`func (o *UploadJobStatus) GetDurationOk() (*float32, bool)`

GetDurationOk returns a tuple with the Duration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDuration

`func (o *UploadJobStatus) SetDuration(v float32)`

SetDuration sets Duration field to given value.

### HasDuration

`func (o *UploadJobStatus) HasDuration() bool`

HasDuration returns a boolean if a field has been set.

### SetDurationNil

`func (o *UploadJobStatus) SetDurationNil(b bool)`

 SetDurationNil sets the value for Duration to be an explicit nil

### UnsetDuration
`func (o *UploadJobStatus) UnsetDuration()`

UnsetDuration ensures that no value is present for Duration, not even an explicit nil
### GetTotalFeatures

`func (o *UploadJobStatus) GetTotalFeatures() int32`

GetTotalFeatures returns the TotalFeatures field if non-nil, zero value otherwise.

### GetTotalFeaturesOk

`func (o *UploadJobStatus) GetTotalFeaturesOk() (*int32, bool)`

GetTotalFeaturesOk returns a tuple with the TotalFeatures field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalFeatures

`func (o *UploadJobStatus) SetTotalFeatures(v int32)`

SetTotalFeatures sets TotalFeatures field to given value.


### GetCreatedCount

`func (o *UploadJobStatus) GetCreatedCount() int32`

GetCreatedCount returns the CreatedCount field if non-nil, zero value otherwise.

### GetCreatedCountOk

`func (o *UploadJobStatus) GetCreatedCountOk() (*int32, bool)`

GetCreatedCountOk returns a tuple with the CreatedCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedCount

`func (o *UploadJobStatus) SetCreatedCount(v int32)`

SetCreatedCount sets CreatedCount field to given value.


### GetFailedCount

`func (o *UploadJobStatus) GetFailedCount() int32`

GetFailedCount returns the FailedCount field if non-nil, zero value otherwise.

### GetFailedCountOk

`func (o *UploadJobStatus) GetFailedCountOk() (*int32, bool)`

GetFailedCountOk returns a tuple with the FailedCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFailedCount

`func (o *UploadJobStatus) SetFailedCount(v int32)`

SetFailedCount sets FailedCount field to given value.


### GetResults

`func (o *UploadJobStatus) GetResults() map[string]interface{}`

GetResults returns the Results field if non-nil, zero value otherwise.

### GetResultsOk

`func (o *UploadJobStatus) GetResultsOk() (*map[string]interface{}, bool)`

GetResultsOk returns a tuple with the Results field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResults

`func (o *UploadJobStatus) SetResults(v map[string]interface{})`

SetResults sets Results field to given value.

### HasResults

`func (o *UploadJobStatus) HasResults() bool`

HasResults returns a boolean if a field has been set.

### SetResultsNil

`func (o *UploadJobStatus) SetResultsNil(b bool)`

 SetResultsNil sets the value for Results to be an explicit nil

### UnsetResults
`func (o *UploadJobStatus) UnsetResults()`

UnsetResults ensures that no value is present for Results, not even an explicit nil
### GetErrorMessage

`func (o *UploadJobStatus) GetErrorMessage() string`

GetErrorMessage returns the ErrorMessage field if non-nil, zero value otherwise.

### GetErrorMessageOk

`func (o *UploadJobStatus) GetErrorMessageOk() (*string, bool)`

GetErrorMessageOk returns a tuple with the ErrorMessage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrorMessage

`func (o *UploadJobStatus) SetErrorMessage(v string)`

SetErrorMessage sets ErrorMessage field to given value.

### HasErrorMessage

`func (o *UploadJobStatus) HasErrorMessage() bool`

HasErrorMessage returns a boolean if a field has been set.

### SetErrorMessageNil

`func (o *UploadJobStatus) SetErrorMessageNil(b bool)`

 SetErrorMessageNil sets the value for ErrorMessage to be an explicit nil

### UnsetErrorMessage
`func (o *UploadJobStatus) UnsetErrorMessage()`

UnsetErrorMessage ensures that no value is present for ErrorMessage, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


