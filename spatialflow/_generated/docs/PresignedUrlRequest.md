# PresignedUrlRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**FileType** | **string** |  | 
**Filename** | **string** |  | 
**FileSize** | **int32** |  | 

## Methods

### NewPresignedUrlRequest

`func NewPresignedUrlRequest(fileType string, filename string, fileSize int32, ) *PresignedUrlRequest`

NewPresignedUrlRequest instantiates a new PresignedUrlRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPresignedUrlRequestWithDefaults

`func NewPresignedUrlRequestWithDefaults() *PresignedUrlRequest`

NewPresignedUrlRequestWithDefaults instantiates a new PresignedUrlRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFileType

`func (o *PresignedUrlRequest) GetFileType() string`

GetFileType returns the FileType field if non-nil, zero value otherwise.

### GetFileTypeOk

`func (o *PresignedUrlRequest) GetFileTypeOk() (*string, bool)`

GetFileTypeOk returns a tuple with the FileType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFileType

`func (o *PresignedUrlRequest) SetFileType(v string)`

SetFileType sets FileType field to given value.


### GetFilename

`func (o *PresignedUrlRequest) GetFilename() string`

GetFilename returns the Filename field if non-nil, zero value otherwise.

### GetFilenameOk

`func (o *PresignedUrlRequest) GetFilenameOk() (*string, bool)`

GetFilenameOk returns a tuple with the Filename field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilename

`func (o *PresignedUrlRequest) SetFilename(v string)`

SetFilename sets Filename field to given value.


### GetFileSize

`func (o *PresignedUrlRequest) GetFileSize() int32`

GetFileSize returns the FileSize field if non-nil, zero value otherwise.

### GetFileSizeOk

`func (o *PresignedUrlRequest) GetFileSizeOk() (*int32, bool)`

GetFileSizeOk returns a tuple with the FileSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFileSize

`func (o *PresignedUrlRequest) SetFileSize(v int32)`

SetFileSize sets FileSize field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


