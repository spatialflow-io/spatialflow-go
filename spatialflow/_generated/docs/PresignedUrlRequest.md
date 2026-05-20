# PresignedUrlRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**FileType** | **string** |  | 
**Filename** | **string** |  | 
**FileSize** | **int32** |  | 
**RelatedObjectType** | Pointer to **NullableString** |  | [optional] 
**RelatedObjectId** | Pointer to **NullableString** |  | [optional] 

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


### GetRelatedObjectType

`func (o *PresignedUrlRequest) GetRelatedObjectType() string`

GetRelatedObjectType returns the RelatedObjectType field if non-nil, zero value otherwise.

### GetRelatedObjectTypeOk

`func (o *PresignedUrlRequest) GetRelatedObjectTypeOk() (*string, bool)`

GetRelatedObjectTypeOk returns a tuple with the RelatedObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRelatedObjectType

`func (o *PresignedUrlRequest) SetRelatedObjectType(v string)`

SetRelatedObjectType sets RelatedObjectType field to given value.

### HasRelatedObjectType

`func (o *PresignedUrlRequest) HasRelatedObjectType() bool`

HasRelatedObjectType returns a boolean if a field has been set.

### SetRelatedObjectTypeNil

`func (o *PresignedUrlRequest) SetRelatedObjectTypeNil(b bool)`

 SetRelatedObjectTypeNil sets the value for RelatedObjectType to be an explicit nil

### UnsetRelatedObjectType
`func (o *PresignedUrlRequest) UnsetRelatedObjectType()`

UnsetRelatedObjectType ensures that no value is present for RelatedObjectType, not even an explicit nil
### GetRelatedObjectId

`func (o *PresignedUrlRequest) GetRelatedObjectId() string`

GetRelatedObjectId returns the RelatedObjectId field if non-nil, zero value otherwise.

### GetRelatedObjectIdOk

`func (o *PresignedUrlRequest) GetRelatedObjectIdOk() (*string, bool)`

GetRelatedObjectIdOk returns a tuple with the RelatedObjectId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRelatedObjectId

`func (o *PresignedUrlRequest) SetRelatedObjectId(v string)`

SetRelatedObjectId sets RelatedObjectId field to given value.

### HasRelatedObjectId

`func (o *PresignedUrlRequest) HasRelatedObjectId() bool`

HasRelatedObjectId returns a boolean if a field has been set.

### SetRelatedObjectIdNil

`func (o *PresignedUrlRequest) SetRelatedObjectIdNil(b bool)`

 SetRelatedObjectIdNil sets the value for RelatedObjectId to be an explicit nil

### UnsetRelatedObjectId
`func (o *PresignedUrlRequest) UnsetRelatedObjectId()`

UnsetRelatedObjectId ensures that no value is present for RelatedObjectId, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


