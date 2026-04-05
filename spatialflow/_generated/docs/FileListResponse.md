# FileListResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Files** | **[]interface{}** |  | 
**Count** | **int32** |  | 
**FileType** | **string** |  | 

## Methods

### NewFileListResponse

`func NewFileListResponse(files []interface{}, count int32, fileType string, ) *FileListResponse`

NewFileListResponse instantiates a new FileListResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFileListResponseWithDefaults

`func NewFileListResponseWithDefaults() *FileListResponse`

NewFileListResponseWithDefaults instantiates a new FileListResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFiles

`func (o *FileListResponse) GetFiles() []interface{}`

GetFiles returns the Files field if non-nil, zero value otherwise.

### GetFilesOk

`func (o *FileListResponse) GetFilesOk() (*[]interface{}, bool)`

GetFilesOk returns a tuple with the Files field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFiles

`func (o *FileListResponse) SetFiles(v []interface{})`

SetFiles sets Files field to given value.


### GetCount

`func (o *FileListResponse) GetCount() int32`

GetCount returns the Count field if non-nil, zero value otherwise.

### GetCountOk

`func (o *FileListResponse) GetCountOk() (*int32, bool)`

GetCountOk returns a tuple with the Count field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCount

`func (o *FileListResponse) SetCount(v int32)`

SetCount sets Count field to given value.


### GetFileType

`func (o *FileListResponse) GetFileType() string`

GetFileType returns the FileType field if non-nil, zero value otherwise.

### GetFileTypeOk

`func (o *FileListResponse) GetFileTypeOk() (*string, bool)`

GetFileTypeOk returns a tuple with the FileType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFileType

`func (o *FileListResponse) SetFileType(v string)`

SetFileType sets FileType field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


