# TagListResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Tags** | [**[]TagSchema**](TagSchema.md) |  | 
**Total** | **int32** |  | 
**Truncated** | Pointer to **bool** | True if the workspace has more than the cap (500) tags and the list was truncated. | [optional] [default to false]

## Methods

### NewTagListResponse

`func NewTagListResponse(tags []TagSchema, total int32, ) *TagListResponse`

NewTagListResponse instantiates a new TagListResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTagListResponseWithDefaults

`func NewTagListResponseWithDefaults() *TagListResponse`

NewTagListResponseWithDefaults instantiates a new TagListResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTags

`func (o *TagListResponse) GetTags() []TagSchema`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *TagListResponse) GetTagsOk() (*[]TagSchema, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *TagListResponse) SetTags(v []TagSchema)`

SetTags sets Tags field to given value.


### GetTotal

`func (o *TagListResponse) GetTotal() int32`

GetTotal returns the Total field if non-nil, zero value otherwise.

### GetTotalOk

`func (o *TagListResponse) GetTotalOk() (*int32, bool)`

GetTotalOk returns a tuple with the Total field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotal

`func (o *TagListResponse) SetTotal(v int32)`

SetTotal sets Total field to given value.


### GetTruncated

`func (o *TagListResponse) GetTruncated() bool`

GetTruncated returns the Truncated field if non-nil, zero value otherwise.

### GetTruncatedOk

`func (o *TagListResponse) GetTruncatedOk() (*bool, bool)`

GetTruncatedOk returns a tuple with the Truncated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTruncated

`func (o *TagListResponse) SetTruncated(v bool)`

SetTruncated sets Truncated field to given value.

### HasTruncated

`func (o *TagListResponse) HasTruncated() bool`

HasTruncated returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


