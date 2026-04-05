# LocationImportResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**JobId** | **string** |  | 
**Status** | **string** |  | 
**Filename** | **string** |  | 
**TotalRows** | Pointer to **int32** |  | [optional] [default to 0]
**ValidRows** | Pointer to **int32** |  | [optional] [default to 0]
**InvalidRows** | Pointer to **int32** |  | [optional] [default to 0]
**ProcessedRows** | Pointer to **int32** |  | [optional] [default to 0]
**ErrorRate** | Pointer to **float32** |  | [optional] [default to 0.0]
**Errors** | Pointer to **[]map[string]interface{}** |  | [optional] [default to []]
**CreatedAt** | **time.Time** |  | 
**StartedAt** | Pointer to **NullableTime** |  | [optional] 
**CompletedAt** | Pointer to **NullableTime** |  | [optional] 

## Methods

### NewLocationImportResponse

`func NewLocationImportResponse(jobId string, status string, filename string, createdAt time.Time, ) *LocationImportResponse`

NewLocationImportResponse instantiates a new LocationImportResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLocationImportResponseWithDefaults

`func NewLocationImportResponseWithDefaults() *LocationImportResponse`

NewLocationImportResponseWithDefaults instantiates a new LocationImportResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetJobId

`func (o *LocationImportResponse) GetJobId() string`

GetJobId returns the JobId field if non-nil, zero value otherwise.

### GetJobIdOk

`func (o *LocationImportResponse) GetJobIdOk() (*string, bool)`

GetJobIdOk returns a tuple with the JobId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJobId

`func (o *LocationImportResponse) SetJobId(v string)`

SetJobId sets JobId field to given value.


### GetStatus

`func (o *LocationImportResponse) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *LocationImportResponse) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *LocationImportResponse) SetStatus(v string)`

SetStatus sets Status field to given value.


### GetFilename

`func (o *LocationImportResponse) GetFilename() string`

GetFilename returns the Filename field if non-nil, zero value otherwise.

### GetFilenameOk

`func (o *LocationImportResponse) GetFilenameOk() (*string, bool)`

GetFilenameOk returns a tuple with the Filename field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilename

`func (o *LocationImportResponse) SetFilename(v string)`

SetFilename sets Filename field to given value.


### GetTotalRows

`func (o *LocationImportResponse) GetTotalRows() int32`

GetTotalRows returns the TotalRows field if non-nil, zero value otherwise.

### GetTotalRowsOk

`func (o *LocationImportResponse) GetTotalRowsOk() (*int32, bool)`

GetTotalRowsOk returns a tuple with the TotalRows field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalRows

`func (o *LocationImportResponse) SetTotalRows(v int32)`

SetTotalRows sets TotalRows field to given value.

### HasTotalRows

`func (o *LocationImportResponse) HasTotalRows() bool`

HasTotalRows returns a boolean if a field has been set.

### GetValidRows

`func (o *LocationImportResponse) GetValidRows() int32`

GetValidRows returns the ValidRows field if non-nil, zero value otherwise.

### GetValidRowsOk

`func (o *LocationImportResponse) GetValidRowsOk() (*int32, bool)`

GetValidRowsOk returns a tuple with the ValidRows field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValidRows

`func (o *LocationImportResponse) SetValidRows(v int32)`

SetValidRows sets ValidRows field to given value.

### HasValidRows

`func (o *LocationImportResponse) HasValidRows() bool`

HasValidRows returns a boolean if a field has been set.

### GetInvalidRows

`func (o *LocationImportResponse) GetInvalidRows() int32`

GetInvalidRows returns the InvalidRows field if non-nil, zero value otherwise.

### GetInvalidRowsOk

`func (o *LocationImportResponse) GetInvalidRowsOk() (*int32, bool)`

GetInvalidRowsOk returns a tuple with the InvalidRows field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInvalidRows

`func (o *LocationImportResponse) SetInvalidRows(v int32)`

SetInvalidRows sets InvalidRows field to given value.

### HasInvalidRows

`func (o *LocationImportResponse) HasInvalidRows() bool`

HasInvalidRows returns a boolean if a field has been set.

### GetProcessedRows

`func (o *LocationImportResponse) GetProcessedRows() int32`

GetProcessedRows returns the ProcessedRows field if non-nil, zero value otherwise.

### GetProcessedRowsOk

`func (o *LocationImportResponse) GetProcessedRowsOk() (*int32, bool)`

GetProcessedRowsOk returns a tuple with the ProcessedRows field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProcessedRows

`func (o *LocationImportResponse) SetProcessedRows(v int32)`

SetProcessedRows sets ProcessedRows field to given value.

### HasProcessedRows

`func (o *LocationImportResponse) HasProcessedRows() bool`

HasProcessedRows returns a boolean if a field has been set.

### GetErrorRate

`func (o *LocationImportResponse) GetErrorRate() float32`

GetErrorRate returns the ErrorRate field if non-nil, zero value otherwise.

### GetErrorRateOk

`func (o *LocationImportResponse) GetErrorRateOk() (*float32, bool)`

GetErrorRateOk returns a tuple with the ErrorRate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrorRate

`func (o *LocationImportResponse) SetErrorRate(v float32)`

SetErrorRate sets ErrorRate field to given value.

### HasErrorRate

`func (o *LocationImportResponse) HasErrorRate() bool`

HasErrorRate returns a boolean if a field has been set.

### GetErrors

`func (o *LocationImportResponse) GetErrors() []*map[string]interface{}`

GetErrors returns the Errors field if non-nil, zero value otherwise.

### GetErrorsOk

`func (o *LocationImportResponse) GetErrorsOk() (*[]*map[string]interface{}, bool)`

GetErrorsOk returns a tuple with the Errors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrors

`func (o *LocationImportResponse) SetErrors(v []*map[string]interface{})`

SetErrors sets Errors field to given value.

### HasErrors

`func (o *LocationImportResponse) HasErrors() bool`

HasErrors returns a boolean if a field has been set.

### GetCreatedAt

`func (o *LocationImportResponse) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *LocationImportResponse) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *LocationImportResponse) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.


### GetStartedAt

`func (o *LocationImportResponse) GetStartedAt() time.Time`

GetStartedAt returns the StartedAt field if non-nil, zero value otherwise.

### GetStartedAtOk

`func (o *LocationImportResponse) GetStartedAtOk() (*time.Time, bool)`

GetStartedAtOk returns a tuple with the StartedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartedAt

`func (o *LocationImportResponse) SetStartedAt(v time.Time)`

SetStartedAt sets StartedAt field to given value.

### HasStartedAt

`func (o *LocationImportResponse) HasStartedAt() bool`

HasStartedAt returns a boolean if a field has been set.

### SetStartedAtNil

`func (o *LocationImportResponse) SetStartedAtNil(b bool)`

 SetStartedAtNil sets the value for StartedAt to be an explicit nil

### UnsetStartedAt
`func (o *LocationImportResponse) UnsetStartedAt()`

UnsetStartedAt ensures that no value is present for StartedAt, not even an explicit nil
### GetCompletedAt

`func (o *LocationImportResponse) GetCompletedAt() time.Time`

GetCompletedAt returns the CompletedAt field if non-nil, zero value otherwise.

### GetCompletedAtOk

`func (o *LocationImportResponse) GetCompletedAtOk() (*time.Time, bool)`

GetCompletedAtOk returns a tuple with the CompletedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompletedAt

`func (o *LocationImportResponse) SetCompletedAt(v time.Time)`

SetCompletedAt sets CompletedAt field to given value.

### HasCompletedAt

`func (o *LocationImportResponse) HasCompletedAt() bool`

HasCompletedAt returns a boolean if a field has been set.

### SetCompletedAtNil

`func (o *LocationImportResponse) SetCompletedAtNil(b bool)`

 SetCompletedAtNil sets the value for CompletedAt to be an explicit nil

### UnsetCompletedAt
`func (o *LocationImportResponse) UnsetCompletedAt()`

UnsetCompletedAt ensures that no value is present for CompletedAt, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


