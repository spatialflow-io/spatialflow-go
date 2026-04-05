# UpdateGeofenceRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **NullableString** |  | [optional] 
**Description** | Pointer to **NullableString** |  | [optional] 
**Geometry** | Pointer to **map[string]interface{}** |  | [optional] 
**WebhookUrl** | Pointer to **NullableString** |  | [optional] 
**WebhookEvents** | Pointer to **[]string** |  | [optional] 
**Metadata** | Pointer to **map[string]interface{}** |  | [optional] 
**IsActive** | Pointer to **NullableBool** |  | [optional] 
**GroupName** | Pointer to **NullableString** |  | [optional] 

## Methods

### NewUpdateGeofenceRequest

`func NewUpdateGeofenceRequest() *UpdateGeofenceRequest`

NewUpdateGeofenceRequest instantiates a new UpdateGeofenceRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateGeofenceRequestWithDefaults

`func NewUpdateGeofenceRequestWithDefaults() *UpdateGeofenceRequest`

NewUpdateGeofenceRequestWithDefaults instantiates a new UpdateGeofenceRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *UpdateGeofenceRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *UpdateGeofenceRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *UpdateGeofenceRequest) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *UpdateGeofenceRequest) HasName() bool`

HasName returns a boolean if a field has been set.

### SetNameNil

`func (o *UpdateGeofenceRequest) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *UpdateGeofenceRequest) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetDescription

`func (o *UpdateGeofenceRequest) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *UpdateGeofenceRequest) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *UpdateGeofenceRequest) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *UpdateGeofenceRequest) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescriptionNil

`func (o *UpdateGeofenceRequest) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *UpdateGeofenceRequest) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetGeometry

`func (o *UpdateGeofenceRequest) GetGeometry() map[string]interface{}`

GetGeometry returns the Geometry field if non-nil, zero value otherwise.

### GetGeometryOk

`func (o *UpdateGeofenceRequest) GetGeometryOk() (*map[string]interface{}, bool)`

GetGeometryOk returns a tuple with the Geometry field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGeometry

`func (o *UpdateGeofenceRequest) SetGeometry(v map[string]interface{})`

SetGeometry sets Geometry field to given value.

### HasGeometry

`func (o *UpdateGeofenceRequest) HasGeometry() bool`

HasGeometry returns a boolean if a field has been set.

### SetGeometryNil

`func (o *UpdateGeofenceRequest) SetGeometryNil(b bool)`

 SetGeometryNil sets the value for Geometry to be an explicit nil

### UnsetGeometry
`func (o *UpdateGeofenceRequest) UnsetGeometry()`

UnsetGeometry ensures that no value is present for Geometry, not even an explicit nil
### GetWebhookUrl

`func (o *UpdateGeofenceRequest) GetWebhookUrl() string`

GetWebhookUrl returns the WebhookUrl field if non-nil, zero value otherwise.

### GetWebhookUrlOk

`func (o *UpdateGeofenceRequest) GetWebhookUrlOk() (*string, bool)`

GetWebhookUrlOk returns a tuple with the WebhookUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWebhookUrl

`func (o *UpdateGeofenceRequest) SetWebhookUrl(v string)`

SetWebhookUrl sets WebhookUrl field to given value.

### HasWebhookUrl

`func (o *UpdateGeofenceRequest) HasWebhookUrl() bool`

HasWebhookUrl returns a boolean if a field has been set.

### SetWebhookUrlNil

`func (o *UpdateGeofenceRequest) SetWebhookUrlNil(b bool)`

 SetWebhookUrlNil sets the value for WebhookUrl to be an explicit nil

### UnsetWebhookUrl
`func (o *UpdateGeofenceRequest) UnsetWebhookUrl()`

UnsetWebhookUrl ensures that no value is present for WebhookUrl, not even an explicit nil
### GetWebhookEvents

`func (o *UpdateGeofenceRequest) GetWebhookEvents() []string`

GetWebhookEvents returns the WebhookEvents field if non-nil, zero value otherwise.

### GetWebhookEventsOk

`func (o *UpdateGeofenceRequest) GetWebhookEventsOk() (*[]string, bool)`

GetWebhookEventsOk returns a tuple with the WebhookEvents field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWebhookEvents

`func (o *UpdateGeofenceRequest) SetWebhookEvents(v []string)`

SetWebhookEvents sets WebhookEvents field to given value.

### HasWebhookEvents

`func (o *UpdateGeofenceRequest) HasWebhookEvents() bool`

HasWebhookEvents returns a boolean if a field has been set.

### SetWebhookEventsNil

`func (o *UpdateGeofenceRequest) SetWebhookEventsNil(b bool)`

 SetWebhookEventsNil sets the value for WebhookEvents to be an explicit nil

### UnsetWebhookEvents
`func (o *UpdateGeofenceRequest) UnsetWebhookEvents()`

UnsetWebhookEvents ensures that no value is present for WebhookEvents, not even an explicit nil
### GetMetadata

`func (o *UpdateGeofenceRequest) GetMetadata() map[string]interface{}`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *UpdateGeofenceRequest) GetMetadataOk() (*map[string]interface{}, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *UpdateGeofenceRequest) SetMetadata(v map[string]interface{})`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *UpdateGeofenceRequest) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### SetMetadataNil

`func (o *UpdateGeofenceRequest) SetMetadataNil(b bool)`

 SetMetadataNil sets the value for Metadata to be an explicit nil

### UnsetMetadata
`func (o *UpdateGeofenceRequest) UnsetMetadata()`

UnsetMetadata ensures that no value is present for Metadata, not even an explicit nil
### GetIsActive

`func (o *UpdateGeofenceRequest) GetIsActive() bool`

GetIsActive returns the IsActive field if non-nil, zero value otherwise.

### GetIsActiveOk

`func (o *UpdateGeofenceRequest) GetIsActiveOk() (*bool, bool)`

GetIsActiveOk returns a tuple with the IsActive field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsActive

`func (o *UpdateGeofenceRequest) SetIsActive(v bool)`

SetIsActive sets IsActive field to given value.

### HasIsActive

`func (o *UpdateGeofenceRequest) HasIsActive() bool`

HasIsActive returns a boolean if a field has been set.

### SetIsActiveNil

`func (o *UpdateGeofenceRequest) SetIsActiveNil(b bool)`

 SetIsActiveNil sets the value for IsActive to be an explicit nil

### UnsetIsActive
`func (o *UpdateGeofenceRequest) UnsetIsActive()`

UnsetIsActive ensures that no value is present for IsActive, not even an explicit nil
### GetGroupName

`func (o *UpdateGeofenceRequest) GetGroupName() string`

GetGroupName returns the GroupName field if non-nil, zero value otherwise.

### GetGroupNameOk

`func (o *UpdateGeofenceRequest) GetGroupNameOk() (*string, bool)`

GetGroupNameOk returns a tuple with the GroupName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroupName

`func (o *UpdateGeofenceRequest) SetGroupName(v string)`

SetGroupName sets GroupName field to given value.

### HasGroupName

`func (o *UpdateGeofenceRequest) HasGroupName() bool`

HasGroupName returns a boolean if a field has been set.

### SetGroupNameNil

`func (o *UpdateGeofenceRequest) SetGroupNameNil(b bool)`

 SetGroupNameNil sets the value for GroupName to be an explicit nil

### UnsetGroupName
`func (o *UpdateGeofenceRequest) UnsetGroupName()`

UnsetGroupName ensures that no value is present for GroupName, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


