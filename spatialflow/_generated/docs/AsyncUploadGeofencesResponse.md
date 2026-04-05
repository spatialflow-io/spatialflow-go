# AsyncUploadGeofencesResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**JobId** | **string** | Upload job ID to track progress | 
**Status** | **string** | Job status (pending, processing, completed, failed) | 
**Message** | **string** | Human-readable status message | 
**StatusUrl** | **string** | URL to check job status | 

## Methods

### NewAsyncUploadGeofencesResponse

`func NewAsyncUploadGeofencesResponse(jobId string, status string, message string, statusUrl string, ) *AsyncUploadGeofencesResponse`

NewAsyncUploadGeofencesResponse instantiates a new AsyncUploadGeofencesResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAsyncUploadGeofencesResponseWithDefaults

`func NewAsyncUploadGeofencesResponseWithDefaults() *AsyncUploadGeofencesResponse`

NewAsyncUploadGeofencesResponseWithDefaults instantiates a new AsyncUploadGeofencesResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetJobId

`func (o *AsyncUploadGeofencesResponse) GetJobId() string`

GetJobId returns the JobId field if non-nil, zero value otherwise.

### GetJobIdOk

`func (o *AsyncUploadGeofencesResponse) GetJobIdOk() (*string, bool)`

GetJobIdOk returns a tuple with the JobId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJobId

`func (o *AsyncUploadGeofencesResponse) SetJobId(v string)`

SetJobId sets JobId field to given value.


### GetStatus

`func (o *AsyncUploadGeofencesResponse) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *AsyncUploadGeofencesResponse) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *AsyncUploadGeofencesResponse) SetStatus(v string)`

SetStatus sets Status field to given value.


### GetMessage

`func (o *AsyncUploadGeofencesResponse) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *AsyncUploadGeofencesResponse) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *AsyncUploadGeofencesResponse) SetMessage(v string)`

SetMessage sets Message field to given value.


### GetStatusUrl

`func (o *AsyncUploadGeofencesResponse) GetStatusUrl() string`

GetStatusUrl returns the StatusUrl field if non-nil, zero value otherwise.

### GetStatusUrlOk

`func (o *AsyncUploadGeofencesResponse) GetStatusUrlOk() (*string, bool)`

GetStatusUrlOk returns a tuple with the StatusUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatusUrl

`func (o *AsyncUploadGeofencesResponse) SetStatusUrl(v string)`

SetStatusUrl sets StatusUrl field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


