# CreateGeofenceRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** |  | 
**Description** | Pointer to **NullableString** |  | [optional] 
**Geometry** | Pointer to [**NullableGeometry1**](Geometry1.md) |  | [optional] 
**WebhookUrl** | Pointer to **NullableString** |  | [optional] 
**WebhookEvents** | Pointer to **[]string** |  | [optional] 
**Metadata** | Pointer to **map[string]interface{}** |  | [optional] 
**GroupName** | Pointer to **NullableString** |  | [optional] 
**Source** | Pointer to **NullableString** |  | [optional] 
**Address** | Pointer to **NullableString** |  | [optional] 
**BufferMeters** | Pointer to **NullableInt32** |  | [optional] 
**ConfirmDuplicate** | Pointer to **bool** | Override dedup conflicts. | [optional] [default to false]
**Tags** | Pointer to **[]string** |  | [optional] 

## Methods

### NewCreateGeofenceRequest

`func NewCreateGeofenceRequest(name string, ) *CreateGeofenceRequest`

NewCreateGeofenceRequest instantiates a new CreateGeofenceRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateGeofenceRequestWithDefaults

`func NewCreateGeofenceRequestWithDefaults() *CreateGeofenceRequest`

NewCreateGeofenceRequestWithDefaults instantiates a new CreateGeofenceRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *CreateGeofenceRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CreateGeofenceRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CreateGeofenceRequest) SetName(v string)`

SetName sets Name field to given value.


### GetDescription

`func (o *CreateGeofenceRequest) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *CreateGeofenceRequest) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *CreateGeofenceRequest) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *CreateGeofenceRequest) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescriptionNil

`func (o *CreateGeofenceRequest) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *CreateGeofenceRequest) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetGeometry

`func (o *CreateGeofenceRequest) GetGeometry() Geometry1`

GetGeometry returns the Geometry field if non-nil, zero value otherwise.

### GetGeometryOk

`func (o *CreateGeofenceRequest) GetGeometryOk() (*Geometry1, bool)`

GetGeometryOk returns a tuple with the Geometry field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGeometry

`func (o *CreateGeofenceRequest) SetGeometry(v Geometry1)`

SetGeometry sets Geometry field to given value.

### HasGeometry

`func (o *CreateGeofenceRequest) HasGeometry() bool`

HasGeometry returns a boolean if a field has been set.

### SetGeometryNil

`func (o *CreateGeofenceRequest) SetGeometryNil(b bool)`

 SetGeometryNil sets the value for Geometry to be an explicit nil

### UnsetGeometry
`func (o *CreateGeofenceRequest) UnsetGeometry()`

UnsetGeometry ensures that no value is present for Geometry, not even an explicit nil
### GetWebhookUrl

`func (o *CreateGeofenceRequest) GetWebhookUrl() string`

GetWebhookUrl returns the WebhookUrl field if non-nil, zero value otherwise.

### GetWebhookUrlOk

`func (o *CreateGeofenceRequest) GetWebhookUrlOk() (*string, bool)`

GetWebhookUrlOk returns a tuple with the WebhookUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWebhookUrl

`func (o *CreateGeofenceRequest) SetWebhookUrl(v string)`

SetWebhookUrl sets WebhookUrl field to given value.

### HasWebhookUrl

`func (o *CreateGeofenceRequest) HasWebhookUrl() bool`

HasWebhookUrl returns a boolean if a field has been set.

### SetWebhookUrlNil

`func (o *CreateGeofenceRequest) SetWebhookUrlNil(b bool)`

 SetWebhookUrlNil sets the value for WebhookUrl to be an explicit nil

### UnsetWebhookUrl
`func (o *CreateGeofenceRequest) UnsetWebhookUrl()`

UnsetWebhookUrl ensures that no value is present for WebhookUrl, not even an explicit nil
### GetWebhookEvents

`func (o *CreateGeofenceRequest) GetWebhookEvents() []string`

GetWebhookEvents returns the WebhookEvents field if non-nil, zero value otherwise.

### GetWebhookEventsOk

`func (o *CreateGeofenceRequest) GetWebhookEventsOk() (*[]string, bool)`

GetWebhookEventsOk returns a tuple with the WebhookEvents field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWebhookEvents

`func (o *CreateGeofenceRequest) SetWebhookEvents(v []string)`

SetWebhookEvents sets WebhookEvents field to given value.

### HasWebhookEvents

`func (o *CreateGeofenceRequest) HasWebhookEvents() bool`

HasWebhookEvents returns a boolean if a field has been set.

### SetWebhookEventsNil

`func (o *CreateGeofenceRequest) SetWebhookEventsNil(b bool)`

 SetWebhookEventsNil sets the value for WebhookEvents to be an explicit nil

### UnsetWebhookEvents
`func (o *CreateGeofenceRequest) UnsetWebhookEvents()`

UnsetWebhookEvents ensures that no value is present for WebhookEvents, not even an explicit nil
### GetMetadata

`func (o *CreateGeofenceRequest) GetMetadata() map[string]interface{}`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *CreateGeofenceRequest) GetMetadataOk() (*map[string]interface{}, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *CreateGeofenceRequest) SetMetadata(v map[string]interface{})`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *CreateGeofenceRequest) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### SetMetadataNil

`func (o *CreateGeofenceRequest) SetMetadataNil(b bool)`

 SetMetadataNil sets the value for Metadata to be an explicit nil

### UnsetMetadata
`func (o *CreateGeofenceRequest) UnsetMetadata()`

UnsetMetadata ensures that no value is present for Metadata, not even an explicit nil
### GetGroupName

`func (o *CreateGeofenceRequest) GetGroupName() string`

GetGroupName returns the GroupName field if non-nil, zero value otherwise.

### GetGroupNameOk

`func (o *CreateGeofenceRequest) GetGroupNameOk() (*string, bool)`

GetGroupNameOk returns a tuple with the GroupName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroupName

`func (o *CreateGeofenceRequest) SetGroupName(v string)`

SetGroupName sets GroupName field to given value.

### HasGroupName

`func (o *CreateGeofenceRequest) HasGroupName() bool`

HasGroupName returns a boolean if a field has been set.

### SetGroupNameNil

`func (o *CreateGeofenceRequest) SetGroupNameNil(b bool)`

 SetGroupNameNil sets the value for GroupName to be an explicit nil

### UnsetGroupName
`func (o *CreateGeofenceRequest) UnsetGroupName()`

UnsetGroupName ensures that no value is present for GroupName, not even an explicit nil
### GetSource

`func (o *CreateGeofenceRequest) GetSource() string`

GetSource returns the Source field if non-nil, zero value otherwise.

### GetSourceOk

`func (o *CreateGeofenceRequest) GetSourceOk() (*string, bool)`

GetSourceOk returns a tuple with the Source field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSource

`func (o *CreateGeofenceRequest) SetSource(v string)`

SetSource sets Source field to given value.

### HasSource

`func (o *CreateGeofenceRequest) HasSource() bool`

HasSource returns a boolean if a field has been set.

### SetSourceNil

`func (o *CreateGeofenceRequest) SetSourceNil(b bool)`

 SetSourceNil sets the value for Source to be an explicit nil

### UnsetSource
`func (o *CreateGeofenceRequest) UnsetSource()`

UnsetSource ensures that no value is present for Source, not even an explicit nil
### GetAddress

`func (o *CreateGeofenceRequest) GetAddress() string`

GetAddress returns the Address field if non-nil, zero value otherwise.

### GetAddressOk

`func (o *CreateGeofenceRequest) GetAddressOk() (*string, bool)`

GetAddressOk returns a tuple with the Address field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddress

`func (o *CreateGeofenceRequest) SetAddress(v string)`

SetAddress sets Address field to given value.

### HasAddress

`func (o *CreateGeofenceRequest) HasAddress() bool`

HasAddress returns a boolean if a field has been set.

### SetAddressNil

`func (o *CreateGeofenceRequest) SetAddressNil(b bool)`

 SetAddressNil sets the value for Address to be an explicit nil

### UnsetAddress
`func (o *CreateGeofenceRequest) UnsetAddress()`

UnsetAddress ensures that no value is present for Address, not even an explicit nil
### GetBufferMeters

`func (o *CreateGeofenceRequest) GetBufferMeters() int32`

GetBufferMeters returns the BufferMeters field if non-nil, zero value otherwise.

### GetBufferMetersOk

`func (o *CreateGeofenceRequest) GetBufferMetersOk() (*int32, bool)`

GetBufferMetersOk returns a tuple with the BufferMeters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBufferMeters

`func (o *CreateGeofenceRequest) SetBufferMeters(v int32)`

SetBufferMeters sets BufferMeters field to given value.

### HasBufferMeters

`func (o *CreateGeofenceRequest) HasBufferMeters() bool`

HasBufferMeters returns a boolean if a field has been set.

### SetBufferMetersNil

`func (o *CreateGeofenceRequest) SetBufferMetersNil(b bool)`

 SetBufferMetersNil sets the value for BufferMeters to be an explicit nil

### UnsetBufferMeters
`func (o *CreateGeofenceRequest) UnsetBufferMeters()`

UnsetBufferMeters ensures that no value is present for BufferMeters, not even an explicit nil
### GetConfirmDuplicate

`func (o *CreateGeofenceRequest) GetConfirmDuplicate() bool`

GetConfirmDuplicate returns the ConfirmDuplicate field if non-nil, zero value otherwise.

### GetConfirmDuplicateOk

`func (o *CreateGeofenceRequest) GetConfirmDuplicateOk() (*bool, bool)`

GetConfirmDuplicateOk returns a tuple with the ConfirmDuplicate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfirmDuplicate

`func (o *CreateGeofenceRequest) SetConfirmDuplicate(v bool)`

SetConfirmDuplicate sets ConfirmDuplicate field to given value.

### HasConfirmDuplicate

`func (o *CreateGeofenceRequest) HasConfirmDuplicate() bool`

HasConfirmDuplicate returns a boolean if a field has been set.

### GetTags

`func (o *CreateGeofenceRequest) GetTags() []string`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *CreateGeofenceRequest) GetTagsOk() (*[]string, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *CreateGeofenceRequest) SetTags(v []string)`

SetTags sets Tags field to given value.

### HasTags

`func (o *CreateGeofenceRequest) HasTags() bool`

HasTags returns a boolean if a field has been set.

### SetTagsNil

`func (o *CreateGeofenceRequest) SetTagsNil(b bool)`

 SetTagsNil sets the value for Tags to be an explicit nil

### UnsetTags
`func (o *CreateGeofenceRequest) UnsetTags()`

UnsetTags ensures that no value is present for Tags, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


