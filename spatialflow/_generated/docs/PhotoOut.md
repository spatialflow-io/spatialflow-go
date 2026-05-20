# PhotoOut

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** |  | 
**FileKey** | **string** |  | 
**OriginalName** | **string** |  | 
**ContentType** | **string** |  | 
**SizeBytes** | **int32** |  | 
**Latitude** | Pointer to **NullableFloat32** |  | [optional] 
**Longitude** | Pointer to **NullableFloat32** |  | [optional] 
**DownloadUrl** | **string** |  | 
**DeviceUuid** | **string** |  | 
**DeviceName** | **string** |  | 
**SessionId** | **string** |  | 
**CapturedAt** | **time.Time** |  | 

## Methods

### NewPhotoOut

`func NewPhotoOut(id string, fileKey string, originalName string, contentType string, sizeBytes int32, downloadUrl string, deviceUuid string, deviceName string, sessionId string, capturedAt time.Time, ) *PhotoOut`

NewPhotoOut instantiates a new PhotoOut object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPhotoOutWithDefaults

`func NewPhotoOutWithDefaults() *PhotoOut`

NewPhotoOutWithDefaults instantiates a new PhotoOut object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *PhotoOut) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *PhotoOut) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *PhotoOut) SetId(v string)`

SetId sets Id field to given value.


### GetFileKey

`func (o *PhotoOut) GetFileKey() string`

GetFileKey returns the FileKey field if non-nil, zero value otherwise.

### GetFileKeyOk

`func (o *PhotoOut) GetFileKeyOk() (*string, bool)`

GetFileKeyOk returns a tuple with the FileKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFileKey

`func (o *PhotoOut) SetFileKey(v string)`

SetFileKey sets FileKey field to given value.


### GetOriginalName

`func (o *PhotoOut) GetOriginalName() string`

GetOriginalName returns the OriginalName field if non-nil, zero value otherwise.

### GetOriginalNameOk

`func (o *PhotoOut) GetOriginalNameOk() (*string, bool)`

GetOriginalNameOk returns a tuple with the OriginalName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOriginalName

`func (o *PhotoOut) SetOriginalName(v string)`

SetOriginalName sets OriginalName field to given value.


### GetContentType

`func (o *PhotoOut) GetContentType() string`

GetContentType returns the ContentType field if non-nil, zero value otherwise.

### GetContentTypeOk

`func (o *PhotoOut) GetContentTypeOk() (*string, bool)`

GetContentTypeOk returns a tuple with the ContentType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContentType

`func (o *PhotoOut) SetContentType(v string)`

SetContentType sets ContentType field to given value.


### GetSizeBytes

`func (o *PhotoOut) GetSizeBytes() int32`

GetSizeBytes returns the SizeBytes field if non-nil, zero value otherwise.

### GetSizeBytesOk

`func (o *PhotoOut) GetSizeBytesOk() (*int32, bool)`

GetSizeBytesOk returns a tuple with the SizeBytes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSizeBytes

`func (o *PhotoOut) SetSizeBytes(v int32)`

SetSizeBytes sets SizeBytes field to given value.


### GetLatitude

`func (o *PhotoOut) GetLatitude() float32`

GetLatitude returns the Latitude field if non-nil, zero value otherwise.

### GetLatitudeOk

`func (o *PhotoOut) GetLatitudeOk() (*float32, bool)`

GetLatitudeOk returns a tuple with the Latitude field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLatitude

`func (o *PhotoOut) SetLatitude(v float32)`

SetLatitude sets Latitude field to given value.

### HasLatitude

`func (o *PhotoOut) HasLatitude() bool`

HasLatitude returns a boolean if a field has been set.

### SetLatitudeNil

`func (o *PhotoOut) SetLatitudeNil(b bool)`

 SetLatitudeNil sets the value for Latitude to be an explicit nil

### UnsetLatitude
`func (o *PhotoOut) UnsetLatitude()`

UnsetLatitude ensures that no value is present for Latitude, not even an explicit nil
### GetLongitude

`func (o *PhotoOut) GetLongitude() float32`

GetLongitude returns the Longitude field if non-nil, zero value otherwise.

### GetLongitudeOk

`func (o *PhotoOut) GetLongitudeOk() (*float32, bool)`

GetLongitudeOk returns a tuple with the Longitude field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLongitude

`func (o *PhotoOut) SetLongitude(v float32)`

SetLongitude sets Longitude field to given value.

### HasLongitude

`func (o *PhotoOut) HasLongitude() bool`

HasLongitude returns a boolean if a field has been set.

### SetLongitudeNil

`func (o *PhotoOut) SetLongitudeNil(b bool)`

 SetLongitudeNil sets the value for Longitude to be an explicit nil

### UnsetLongitude
`func (o *PhotoOut) UnsetLongitude()`

UnsetLongitude ensures that no value is present for Longitude, not even an explicit nil
### GetDownloadUrl

`func (o *PhotoOut) GetDownloadUrl() string`

GetDownloadUrl returns the DownloadUrl field if non-nil, zero value otherwise.

### GetDownloadUrlOk

`func (o *PhotoOut) GetDownloadUrlOk() (*string, bool)`

GetDownloadUrlOk returns a tuple with the DownloadUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDownloadUrl

`func (o *PhotoOut) SetDownloadUrl(v string)`

SetDownloadUrl sets DownloadUrl field to given value.


### GetDeviceUuid

`func (o *PhotoOut) GetDeviceUuid() string`

GetDeviceUuid returns the DeviceUuid field if non-nil, zero value otherwise.

### GetDeviceUuidOk

`func (o *PhotoOut) GetDeviceUuidOk() (*string, bool)`

GetDeviceUuidOk returns a tuple with the DeviceUuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeviceUuid

`func (o *PhotoOut) SetDeviceUuid(v string)`

SetDeviceUuid sets DeviceUuid field to given value.


### GetDeviceName

`func (o *PhotoOut) GetDeviceName() string`

GetDeviceName returns the DeviceName field if non-nil, zero value otherwise.

### GetDeviceNameOk

`func (o *PhotoOut) GetDeviceNameOk() (*string, bool)`

GetDeviceNameOk returns a tuple with the DeviceName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeviceName

`func (o *PhotoOut) SetDeviceName(v string)`

SetDeviceName sets DeviceName field to given value.


### GetSessionId

`func (o *PhotoOut) GetSessionId() string`

GetSessionId returns the SessionId field if non-nil, zero value otherwise.

### GetSessionIdOk

`func (o *PhotoOut) GetSessionIdOk() (*string, bool)`

GetSessionIdOk returns a tuple with the SessionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSessionId

`func (o *PhotoOut) SetSessionId(v string)`

SetSessionId sets SessionId field to given value.


### GetCapturedAt

`func (o *PhotoOut) GetCapturedAt() time.Time`

GetCapturedAt returns the CapturedAt field if non-nil, zero value otherwise.

### GetCapturedAtOk

`func (o *PhotoOut) GetCapturedAtOk() (*time.Time, bool)`

GetCapturedAtOk returns a tuple with the CapturedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCapturedAt

`func (o *PhotoOut) SetCapturedAt(v time.Time)`

SetCapturedAt sets CapturedAt field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


