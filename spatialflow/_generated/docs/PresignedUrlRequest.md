# PresignedUrlRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**FileType** | **string** |  | 
**Filename** | **string** |  | 
**FileSize** | **int32** |  | 
**RelatedObjectType** | Pointer to **NullableString** |  | [optional] 
**RelatedObjectId** | Pointer to **NullableString** |  | [optional] 
**CapturedAt** | Pointer to **NullableTime** |  | [optional] 
**CaptureLatitude** | Pointer to **NullableFloat32** |  | [optional] 
**CaptureLongitude** | Pointer to **NullableFloat32** |  | [optional] 
**CaptureAccuracyM** | Pointer to **NullableFloat32** |  | [optional] 

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
### GetCapturedAt

`func (o *PresignedUrlRequest) GetCapturedAt() time.Time`

GetCapturedAt returns the CapturedAt field if non-nil, zero value otherwise.

### GetCapturedAtOk

`func (o *PresignedUrlRequest) GetCapturedAtOk() (*time.Time, bool)`

GetCapturedAtOk returns a tuple with the CapturedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCapturedAt

`func (o *PresignedUrlRequest) SetCapturedAt(v time.Time)`

SetCapturedAt sets CapturedAt field to given value.

### HasCapturedAt

`func (o *PresignedUrlRequest) HasCapturedAt() bool`

HasCapturedAt returns a boolean if a field has been set.

### SetCapturedAtNil

`func (o *PresignedUrlRequest) SetCapturedAtNil(b bool)`

 SetCapturedAtNil sets the value for CapturedAt to be an explicit nil

### UnsetCapturedAt
`func (o *PresignedUrlRequest) UnsetCapturedAt()`

UnsetCapturedAt ensures that no value is present for CapturedAt, not even an explicit nil
### GetCaptureLatitude

`func (o *PresignedUrlRequest) GetCaptureLatitude() float32`

GetCaptureLatitude returns the CaptureLatitude field if non-nil, zero value otherwise.

### GetCaptureLatitudeOk

`func (o *PresignedUrlRequest) GetCaptureLatitudeOk() (*float32, bool)`

GetCaptureLatitudeOk returns a tuple with the CaptureLatitude field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCaptureLatitude

`func (o *PresignedUrlRequest) SetCaptureLatitude(v float32)`

SetCaptureLatitude sets CaptureLatitude field to given value.

### HasCaptureLatitude

`func (o *PresignedUrlRequest) HasCaptureLatitude() bool`

HasCaptureLatitude returns a boolean if a field has been set.

### SetCaptureLatitudeNil

`func (o *PresignedUrlRequest) SetCaptureLatitudeNil(b bool)`

 SetCaptureLatitudeNil sets the value for CaptureLatitude to be an explicit nil

### UnsetCaptureLatitude
`func (o *PresignedUrlRequest) UnsetCaptureLatitude()`

UnsetCaptureLatitude ensures that no value is present for CaptureLatitude, not even an explicit nil
### GetCaptureLongitude

`func (o *PresignedUrlRequest) GetCaptureLongitude() float32`

GetCaptureLongitude returns the CaptureLongitude field if non-nil, zero value otherwise.

### GetCaptureLongitudeOk

`func (o *PresignedUrlRequest) GetCaptureLongitudeOk() (*float32, bool)`

GetCaptureLongitudeOk returns a tuple with the CaptureLongitude field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCaptureLongitude

`func (o *PresignedUrlRequest) SetCaptureLongitude(v float32)`

SetCaptureLongitude sets CaptureLongitude field to given value.

### HasCaptureLongitude

`func (o *PresignedUrlRequest) HasCaptureLongitude() bool`

HasCaptureLongitude returns a boolean if a field has been set.

### SetCaptureLongitudeNil

`func (o *PresignedUrlRequest) SetCaptureLongitudeNil(b bool)`

 SetCaptureLongitudeNil sets the value for CaptureLongitude to be an explicit nil

### UnsetCaptureLongitude
`func (o *PresignedUrlRequest) UnsetCaptureLongitude()`

UnsetCaptureLongitude ensures that no value is present for CaptureLongitude, not even an explicit nil
### GetCaptureAccuracyM

`func (o *PresignedUrlRequest) GetCaptureAccuracyM() float32`

GetCaptureAccuracyM returns the CaptureAccuracyM field if non-nil, zero value otherwise.

### GetCaptureAccuracyMOk

`func (o *PresignedUrlRequest) GetCaptureAccuracyMOk() (*float32, bool)`

GetCaptureAccuracyMOk returns a tuple with the CaptureAccuracyM field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCaptureAccuracyM

`func (o *PresignedUrlRequest) SetCaptureAccuracyM(v float32)`

SetCaptureAccuracyM sets CaptureAccuracyM field to given value.

### HasCaptureAccuracyM

`func (o *PresignedUrlRequest) HasCaptureAccuracyM() bool`

HasCaptureAccuracyM returns a boolean if a field has been set.

### SetCaptureAccuracyMNil

`func (o *PresignedUrlRequest) SetCaptureAccuracyMNil(b bool)`

 SetCaptureAccuracyMNil sets the value for CaptureAccuracyM to be an explicit nil

### UnsetCaptureAccuracyM
`func (o *PresignedUrlRequest) UnsetCaptureAccuracyM()`

UnsetCaptureAccuracyM ensures that no value is present for CaptureAccuracyM, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


