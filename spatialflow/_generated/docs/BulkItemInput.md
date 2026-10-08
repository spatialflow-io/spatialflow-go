# BulkItemInput

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Address** | **string** |  | 
**Name** | Pointer to **NullableString** |  | [optional] 
**BufferMeters** | Pointer to **NullableInt32** |  | [optional] 
**Tags** | Pointer to **[]string** |  | [optional] 

## Methods

### NewBulkItemInput

`func NewBulkItemInput(address string, ) *BulkItemInput`

NewBulkItemInput instantiates a new BulkItemInput object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBulkItemInputWithDefaults

`func NewBulkItemInputWithDefaults() *BulkItemInput`

NewBulkItemInputWithDefaults instantiates a new BulkItemInput object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAddress

`func (o *BulkItemInput) GetAddress() string`

GetAddress returns the Address field if non-nil, zero value otherwise.

### GetAddressOk

`func (o *BulkItemInput) GetAddressOk() (*string, bool)`

GetAddressOk returns a tuple with the Address field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddress

`func (o *BulkItemInput) SetAddress(v string)`

SetAddress sets Address field to given value.


### GetName

`func (o *BulkItemInput) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *BulkItemInput) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *BulkItemInput) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *BulkItemInput) HasName() bool`

HasName returns a boolean if a field has been set.

### SetNameNil

`func (o *BulkItemInput) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *BulkItemInput) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetBufferMeters

`func (o *BulkItemInput) GetBufferMeters() int32`

GetBufferMeters returns the BufferMeters field if non-nil, zero value otherwise.

### GetBufferMetersOk

`func (o *BulkItemInput) GetBufferMetersOk() (*int32, bool)`

GetBufferMetersOk returns a tuple with the BufferMeters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBufferMeters

`func (o *BulkItemInput) SetBufferMeters(v int32)`

SetBufferMeters sets BufferMeters field to given value.

### HasBufferMeters

`func (o *BulkItemInput) HasBufferMeters() bool`

HasBufferMeters returns a boolean if a field has been set.

### SetBufferMetersNil

`func (o *BulkItemInput) SetBufferMetersNil(b bool)`

 SetBufferMetersNil sets the value for BufferMeters to be an explicit nil

### UnsetBufferMeters
`func (o *BulkItemInput) UnsetBufferMeters()`

UnsetBufferMeters ensures that no value is present for BufferMeters, not even an explicit nil
### GetTags

`func (o *BulkItemInput) GetTags() []string`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *BulkItemInput) GetTagsOk() (*[]string, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *BulkItemInput) SetTags(v []string)`

SetTags sets Tags field to given value.

### HasTags

`func (o *BulkItemInput) HasTags() bool`

HasTags returns a boolean if a field has been set.

### SetTagsNil

`func (o *BulkItemInput) SetTagsNil(b bool)`

 SetTagsNil sets the value for Tags to be an explicit nil

### UnsetTags
`func (o *BulkItemInput) UnsetTags()`

UnsetTags ensures that no value is present for Tags, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


