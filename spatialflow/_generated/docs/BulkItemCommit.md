# BulkItemCommit

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Index** | **int32** |  | 
**Address** | **string** |  | 
**Name** | Pointer to **NullableString** |  | [optional] 
**BufferMeters** | Pointer to **NullableInt32** |  | [optional] 
**Tags** | Pointer to **[]string** |  | [optional] 
**Status** | **string** |  | 
**OverrideDedup** | Pointer to **bool** |  | [optional] [default to false]

## Methods

### NewBulkItemCommit

`func NewBulkItemCommit(index int32, address string, status string, ) *BulkItemCommit`

NewBulkItemCommit instantiates a new BulkItemCommit object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBulkItemCommitWithDefaults

`func NewBulkItemCommitWithDefaults() *BulkItemCommit`

NewBulkItemCommitWithDefaults instantiates a new BulkItemCommit object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIndex

`func (o *BulkItemCommit) GetIndex() int32`

GetIndex returns the Index field if non-nil, zero value otherwise.

### GetIndexOk

`func (o *BulkItemCommit) GetIndexOk() (*int32, bool)`

GetIndexOk returns a tuple with the Index field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIndex

`func (o *BulkItemCommit) SetIndex(v int32)`

SetIndex sets Index field to given value.


### GetAddress

`func (o *BulkItemCommit) GetAddress() string`

GetAddress returns the Address field if non-nil, zero value otherwise.

### GetAddressOk

`func (o *BulkItemCommit) GetAddressOk() (*string, bool)`

GetAddressOk returns a tuple with the Address field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddress

`func (o *BulkItemCommit) SetAddress(v string)`

SetAddress sets Address field to given value.


### GetName

`func (o *BulkItemCommit) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *BulkItemCommit) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *BulkItemCommit) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *BulkItemCommit) HasName() bool`

HasName returns a boolean if a field has been set.

### SetNameNil

`func (o *BulkItemCommit) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *BulkItemCommit) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetBufferMeters

`func (o *BulkItemCommit) GetBufferMeters() int32`

GetBufferMeters returns the BufferMeters field if non-nil, zero value otherwise.

### GetBufferMetersOk

`func (o *BulkItemCommit) GetBufferMetersOk() (*int32, bool)`

GetBufferMetersOk returns a tuple with the BufferMeters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBufferMeters

`func (o *BulkItemCommit) SetBufferMeters(v int32)`

SetBufferMeters sets BufferMeters field to given value.

### HasBufferMeters

`func (o *BulkItemCommit) HasBufferMeters() bool`

HasBufferMeters returns a boolean if a field has been set.

### SetBufferMetersNil

`func (o *BulkItemCommit) SetBufferMetersNil(b bool)`

 SetBufferMetersNil sets the value for BufferMeters to be an explicit nil

### UnsetBufferMeters
`func (o *BulkItemCommit) UnsetBufferMeters()`

UnsetBufferMeters ensures that no value is present for BufferMeters, not even an explicit nil
### GetTags

`func (o *BulkItemCommit) GetTags() []string`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *BulkItemCommit) GetTagsOk() (*[]string, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *BulkItemCommit) SetTags(v []string)`

SetTags sets Tags field to given value.

### HasTags

`func (o *BulkItemCommit) HasTags() bool`

HasTags returns a boolean if a field has been set.

### SetTagsNil

`func (o *BulkItemCommit) SetTagsNil(b bool)`

 SetTagsNil sets the value for Tags to be an explicit nil

### UnsetTags
`func (o *BulkItemCommit) UnsetTags()`

UnsetTags ensures that no value is present for Tags, not even an explicit nil
### GetStatus

`func (o *BulkItemCommit) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *BulkItemCommit) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *BulkItemCommit) SetStatus(v string)`

SetStatus sets Status field to given value.


### GetOverrideDedup

`func (o *BulkItemCommit) GetOverrideDedup() bool`

GetOverrideDedup returns the OverrideDedup field if non-nil, zero value otherwise.

### GetOverrideDedupOk

`func (o *BulkItemCommit) GetOverrideDedupOk() (*bool, bool)`

GetOverrideDedupOk returns a tuple with the OverrideDedup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOverrideDedup

`func (o *BulkItemCommit) SetOverrideDedup(v bool)`

SetOverrideDedup sets OverrideDedup field to given value.

### HasOverrideDedup

`func (o *BulkItemCommit) HasOverrideDedup() bool`

HasOverrideDedup returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


