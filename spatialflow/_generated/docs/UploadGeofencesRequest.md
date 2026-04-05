# UploadGeofencesRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**FileId** | **string** | ID of the uploaded file from storage service | 
**GroupName** | Pointer to **NullableString** |  | [optional] 

## Methods

### NewUploadGeofencesRequest

`func NewUploadGeofencesRequest(fileId string, ) *UploadGeofencesRequest`

NewUploadGeofencesRequest instantiates a new UploadGeofencesRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUploadGeofencesRequestWithDefaults

`func NewUploadGeofencesRequestWithDefaults() *UploadGeofencesRequest`

NewUploadGeofencesRequestWithDefaults instantiates a new UploadGeofencesRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFileId

`func (o *UploadGeofencesRequest) GetFileId() string`

GetFileId returns the FileId field if non-nil, zero value otherwise.

### GetFileIdOk

`func (o *UploadGeofencesRequest) GetFileIdOk() (*string, bool)`

GetFileIdOk returns a tuple with the FileId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFileId

`func (o *UploadGeofencesRequest) SetFileId(v string)`

SetFileId sets FileId field to given value.


### GetGroupName

`func (o *UploadGeofencesRequest) GetGroupName() string`

GetGroupName returns the GroupName field if non-nil, zero value otherwise.

### GetGroupNameOk

`func (o *UploadGeofencesRequest) GetGroupNameOk() (*string, bool)`

GetGroupNameOk returns a tuple with the GroupName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroupName

`func (o *UploadGeofencesRequest) SetGroupName(v string)`

SetGroupName sets GroupName field to given value.

### HasGroupName

`func (o *UploadGeofencesRequest) HasGroupName() bool`

HasGroupName returns a boolean if a field has been set.

### SetGroupNameNil

`func (o *UploadGeofencesRequest) SetGroupNameNil(b bool)`

 SetGroupNameNil sets the value for GroupName to be an explicit nil

### UnsetGroupName
`func (o *UploadGeofencesRequest) UnsetGroupName()`

UnsetGroupName ensures that no value is present for GroupName, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


