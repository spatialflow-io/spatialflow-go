# WorkspacePhotosOut

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Photos** | [**[]PhotoOut**](PhotoOut.md) |  | 
**HasMore** | **bool** |  | 

## Methods

### NewWorkspacePhotosOut

`func NewWorkspacePhotosOut(photos []PhotoOut, hasMore bool, ) *WorkspacePhotosOut`

NewWorkspacePhotosOut instantiates a new WorkspacePhotosOut object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWorkspacePhotosOutWithDefaults

`func NewWorkspacePhotosOutWithDefaults() *WorkspacePhotosOut`

NewWorkspacePhotosOutWithDefaults instantiates a new WorkspacePhotosOut object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPhotos

`func (o *WorkspacePhotosOut) GetPhotos() []PhotoOut`

GetPhotos returns the Photos field if non-nil, zero value otherwise.

### GetPhotosOk

`func (o *WorkspacePhotosOut) GetPhotosOk() (*[]PhotoOut, bool)`

GetPhotosOk returns a tuple with the Photos field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPhotos

`func (o *WorkspacePhotosOut) SetPhotos(v []PhotoOut)`

SetPhotos sets Photos field to given value.


### GetHasMore

`func (o *WorkspacePhotosOut) GetHasMore() bool`

GetHasMore returns the HasMore field if non-nil, zero value otherwise.

### GetHasMoreOk

`func (o *WorkspacePhotosOut) GetHasMoreOk() (*bool, bool)`

GetHasMoreOk returns a tuple with the HasMore field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasMore

`func (o *WorkspacePhotosOut) SetHasMore(v bool)`

SetHasMore sets HasMore field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


