# GeofenceResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** |  | 
**Name** | **string** |  | 
**Description** | **NullableString** |  | 
**Geometry** | [**Geometry**](Geometry.md) |  | 
**GeometryType** | **string** | Logical geometry type: Polygon, MultiPolygon, or Circle | 
**RadiusMeters** | Pointer to **NullableFloat32** |  | [optional] 
**WebhookUrl** | **NullableString** |  | 
**WebhookEvents** | **[]string** |  | 
**Metadata** | **map[string]interface{}** |  | 
**IsActive** | **bool** |  | 
**GroupId** | **NullableString** |  | 
**GroupName** | **NullableString** |  | 
**Source** | Pointer to **NullableString** |  | [optional] 
**Address** | Pointer to **NullableString** |  | [optional] 
**BufferMeters** | Pointer to **NullableInt32** |  | [optional] 
**Archived** | Pointer to **bool** | Whether the geofence is archived (hidden from default list and map). | [optional] [default to false]
**Point** | Pointer to **[]float32** |  | [optional] 
**Tags** | Pointer to **[]string** | Names of tags attached to this geofence (case as originally entered). | [optional] 
**SourceId** | Pointer to **NullableString** |  | [optional] 
**PreUpgradeGeometry** | Pointer to **map[string]interface{}** |  | [optional] 
**BuildingProvenance** | Pointer to **map[string]interface{}** |  | [optional] 
**EffectiveRadiusMeters** | Pointer to **NullableFloat32** |  | [optional] 
**BelowMinTriggerRadius** | Pointer to **bool** | True iff effective_radius_meters is strictly less than MIN_TRIGGER_RADIUS_METERS (50m). When True, the geofence fires within ~50m of the center. Non-blocking — the client should show a reassuring informational warning, not block Save. | [optional] [default to false]
**IsExample** | Pointer to **bool** |  | [optional] [default to false]
**CreatedAt** | **time.Time** |  | 
**UpdatedAt** | **time.Time** |  | 

## Methods

### NewGeofenceResponse

`func NewGeofenceResponse(id string, name string, description NullableString, geometry Geometry, geometryType string, webhookUrl NullableString, webhookEvents []string, metadata map[string]interface{}, isActive bool, groupId NullableString, groupName NullableString, createdAt time.Time, updatedAt time.Time, ) *GeofenceResponse`

NewGeofenceResponse instantiates a new GeofenceResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGeofenceResponseWithDefaults

`func NewGeofenceResponseWithDefaults() *GeofenceResponse`

NewGeofenceResponseWithDefaults instantiates a new GeofenceResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *GeofenceResponse) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *GeofenceResponse) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *GeofenceResponse) SetId(v string)`

SetId sets Id field to given value.


### GetName

`func (o *GeofenceResponse) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *GeofenceResponse) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *GeofenceResponse) SetName(v string)`

SetName sets Name field to given value.


### GetDescription

`func (o *GeofenceResponse) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *GeofenceResponse) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *GeofenceResponse) SetDescription(v string)`

SetDescription sets Description field to given value.


### SetDescriptionNil

`func (o *GeofenceResponse) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *GeofenceResponse) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetGeometry

`func (o *GeofenceResponse) GetGeometry() Geometry`

GetGeometry returns the Geometry field if non-nil, zero value otherwise.

### GetGeometryOk

`func (o *GeofenceResponse) GetGeometryOk() (*Geometry, bool)`

GetGeometryOk returns a tuple with the Geometry field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGeometry

`func (o *GeofenceResponse) SetGeometry(v Geometry)`

SetGeometry sets Geometry field to given value.


### GetGeometryType

`func (o *GeofenceResponse) GetGeometryType() string`

GetGeometryType returns the GeometryType field if non-nil, zero value otherwise.

### GetGeometryTypeOk

`func (o *GeofenceResponse) GetGeometryTypeOk() (*string, bool)`

GetGeometryTypeOk returns a tuple with the GeometryType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGeometryType

`func (o *GeofenceResponse) SetGeometryType(v string)`

SetGeometryType sets GeometryType field to given value.


### GetRadiusMeters

`func (o *GeofenceResponse) GetRadiusMeters() float32`

GetRadiusMeters returns the RadiusMeters field if non-nil, zero value otherwise.

### GetRadiusMetersOk

`func (o *GeofenceResponse) GetRadiusMetersOk() (*float32, bool)`

GetRadiusMetersOk returns a tuple with the RadiusMeters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRadiusMeters

`func (o *GeofenceResponse) SetRadiusMeters(v float32)`

SetRadiusMeters sets RadiusMeters field to given value.

### HasRadiusMeters

`func (o *GeofenceResponse) HasRadiusMeters() bool`

HasRadiusMeters returns a boolean if a field has been set.

### SetRadiusMetersNil

`func (o *GeofenceResponse) SetRadiusMetersNil(b bool)`

 SetRadiusMetersNil sets the value for RadiusMeters to be an explicit nil

### UnsetRadiusMeters
`func (o *GeofenceResponse) UnsetRadiusMeters()`

UnsetRadiusMeters ensures that no value is present for RadiusMeters, not even an explicit nil
### GetWebhookUrl

`func (o *GeofenceResponse) GetWebhookUrl() string`

GetWebhookUrl returns the WebhookUrl field if non-nil, zero value otherwise.

### GetWebhookUrlOk

`func (o *GeofenceResponse) GetWebhookUrlOk() (*string, bool)`

GetWebhookUrlOk returns a tuple with the WebhookUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWebhookUrl

`func (o *GeofenceResponse) SetWebhookUrl(v string)`

SetWebhookUrl sets WebhookUrl field to given value.


### SetWebhookUrlNil

`func (o *GeofenceResponse) SetWebhookUrlNil(b bool)`

 SetWebhookUrlNil sets the value for WebhookUrl to be an explicit nil

### UnsetWebhookUrl
`func (o *GeofenceResponse) UnsetWebhookUrl()`

UnsetWebhookUrl ensures that no value is present for WebhookUrl, not even an explicit nil
### GetWebhookEvents

`func (o *GeofenceResponse) GetWebhookEvents() []string`

GetWebhookEvents returns the WebhookEvents field if non-nil, zero value otherwise.

### GetWebhookEventsOk

`func (o *GeofenceResponse) GetWebhookEventsOk() (*[]string, bool)`

GetWebhookEventsOk returns a tuple with the WebhookEvents field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWebhookEvents

`func (o *GeofenceResponse) SetWebhookEvents(v []string)`

SetWebhookEvents sets WebhookEvents field to given value.


### GetMetadata

`func (o *GeofenceResponse) GetMetadata() map[string]interface{}`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *GeofenceResponse) GetMetadataOk() (*map[string]interface{}, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *GeofenceResponse) SetMetadata(v map[string]interface{})`

SetMetadata sets Metadata field to given value.


### GetIsActive

`func (o *GeofenceResponse) GetIsActive() bool`

GetIsActive returns the IsActive field if non-nil, zero value otherwise.

### GetIsActiveOk

`func (o *GeofenceResponse) GetIsActiveOk() (*bool, bool)`

GetIsActiveOk returns a tuple with the IsActive field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsActive

`func (o *GeofenceResponse) SetIsActive(v bool)`

SetIsActive sets IsActive field to given value.


### GetGroupId

`func (o *GeofenceResponse) GetGroupId() string`

GetGroupId returns the GroupId field if non-nil, zero value otherwise.

### GetGroupIdOk

`func (o *GeofenceResponse) GetGroupIdOk() (*string, bool)`

GetGroupIdOk returns a tuple with the GroupId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroupId

`func (o *GeofenceResponse) SetGroupId(v string)`

SetGroupId sets GroupId field to given value.


### SetGroupIdNil

`func (o *GeofenceResponse) SetGroupIdNil(b bool)`

 SetGroupIdNil sets the value for GroupId to be an explicit nil

### UnsetGroupId
`func (o *GeofenceResponse) UnsetGroupId()`

UnsetGroupId ensures that no value is present for GroupId, not even an explicit nil
### GetGroupName

`func (o *GeofenceResponse) GetGroupName() string`

GetGroupName returns the GroupName field if non-nil, zero value otherwise.

### GetGroupNameOk

`func (o *GeofenceResponse) GetGroupNameOk() (*string, bool)`

GetGroupNameOk returns a tuple with the GroupName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroupName

`func (o *GeofenceResponse) SetGroupName(v string)`

SetGroupName sets GroupName field to given value.


### SetGroupNameNil

`func (o *GeofenceResponse) SetGroupNameNil(b bool)`

 SetGroupNameNil sets the value for GroupName to be an explicit nil

### UnsetGroupName
`func (o *GeofenceResponse) UnsetGroupName()`

UnsetGroupName ensures that no value is present for GroupName, not even an explicit nil
### GetSource

`func (o *GeofenceResponse) GetSource() string`

GetSource returns the Source field if non-nil, zero value otherwise.

### GetSourceOk

`func (o *GeofenceResponse) GetSourceOk() (*string, bool)`

GetSourceOk returns a tuple with the Source field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSource

`func (o *GeofenceResponse) SetSource(v string)`

SetSource sets Source field to given value.

### HasSource

`func (o *GeofenceResponse) HasSource() bool`

HasSource returns a boolean if a field has been set.

### SetSourceNil

`func (o *GeofenceResponse) SetSourceNil(b bool)`

 SetSourceNil sets the value for Source to be an explicit nil

### UnsetSource
`func (o *GeofenceResponse) UnsetSource()`

UnsetSource ensures that no value is present for Source, not even an explicit nil
### GetAddress

`func (o *GeofenceResponse) GetAddress() string`

GetAddress returns the Address field if non-nil, zero value otherwise.

### GetAddressOk

`func (o *GeofenceResponse) GetAddressOk() (*string, bool)`

GetAddressOk returns a tuple with the Address field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddress

`func (o *GeofenceResponse) SetAddress(v string)`

SetAddress sets Address field to given value.

### HasAddress

`func (o *GeofenceResponse) HasAddress() bool`

HasAddress returns a boolean if a field has been set.

### SetAddressNil

`func (o *GeofenceResponse) SetAddressNil(b bool)`

 SetAddressNil sets the value for Address to be an explicit nil

### UnsetAddress
`func (o *GeofenceResponse) UnsetAddress()`

UnsetAddress ensures that no value is present for Address, not even an explicit nil
### GetBufferMeters

`func (o *GeofenceResponse) GetBufferMeters() int32`

GetBufferMeters returns the BufferMeters field if non-nil, zero value otherwise.

### GetBufferMetersOk

`func (o *GeofenceResponse) GetBufferMetersOk() (*int32, bool)`

GetBufferMetersOk returns a tuple with the BufferMeters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBufferMeters

`func (o *GeofenceResponse) SetBufferMeters(v int32)`

SetBufferMeters sets BufferMeters field to given value.

### HasBufferMeters

`func (o *GeofenceResponse) HasBufferMeters() bool`

HasBufferMeters returns a boolean if a field has been set.

### SetBufferMetersNil

`func (o *GeofenceResponse) SetBufferMetersNil(b bool)`

 SetBufferMetersNil sets the value for BufferMeters to be an explicit nil

### UnsetBufferMeters
`func (o *GeofenceResponse) UnsetBufferMeters()`

UnsetBufferMeters ensures that no value is present for BufferMeters, not even an explicit nil
### GetArchived

`func (o *GeofenceResponse) GetArchived() bool`

GetArchived returns the Archived field if non-nil, zero value otherwise.

### GetArchivedOk

`func (o *GeofenceResponse) GetArchivedOk() (*bool, bool)`

GetArchivedOk returns a tuple with the Archived field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArchived

`func (o *GeofenceResponse) SetArchived(v bool)`

SetArchived sets Archived field to given value.

### HasArchived

`func (o *GeofenceResponse) HasArchived() bool`

HasArchived returns a boolean if a field has been set.

### GetPoint

`func (o *GeofenceResponse) GetPoint() []float32`

GetPoint returns the Point field if non-nil, zero value otherwise.

### GetPointOk

`func (o *GeofenceResponse) GetPointOk() (*[]float32, bool)`

GetPointOk returns a tuple with the Point field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPoint

`func (o *GeofenceResponse) SetPoint(v []float32)`

SetPoint sets Point field to given value.

### HasPoint

`func (o *GeofenceResponse) HasPoint() bool`

HasPoint returns a boolean if a field has been set.

### SetPointNil

`func (o *GeofenceResponse) SetPointNil(b bool)`

 SetPointNil sets the value for Point to be an explicit nil

### UnsetPoint
`func (o *GeofenceResponse) UnsetPoint()`

UnsetPoint ensures that no value is present for Point, not even an explicit nil
### GetTags

`func (o *GeofenceResponse) GetTags() []string`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *GeofenceResponse) GetTagsOk() (*[]string, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *GeofenceResponse) SetTags(v []string)`

SetTags sets Tags field to given value.

### HasTags

`func (o *GeofenceResponse) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetSourceId

`func (o *GeofenceResponse) GetSourceId() string`

GetSourceId returns the SourceId field if non-nil, zero value otherwise.

### GetSourceIdOk

`func (o *GeofenceResponse) GetSourceIdOk() (*string, bool)`

GetSourceIdOk returns a tuple with the SourceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceId

`func (o *GeofenceResponse) SetSourceId(v string)`

SetSourceId sets SourceId field to given value.

### HasSourceId

`func (o *GeofenceResponse) HasSourceId() bool`

HasSourceId returns a boolean if a field has been set.

### SetSourceIdNil

`func (o *GeofenceResponse) SetSourceIdNil(b bool)`

 SetSourceIdNil sets the value for SourceId to be an explicit nil

### UnsetSourceId
`func (o *GeofenceResponse) UnsetSourceId()`

UnsetSourceId ensures that no value is present for SourceId, not even an explicit nil
### GetPreUpgradeGeometry

`func (o *GeofenceResponse) GetPreUpgradeGeometry() map[string]interface{}`

GetPreUpgradeGeometry returns the PreUpgradeGeometry field if non-nil, zero value otherwise.

### GetPreUpgradeGeometryOk

`func (o *GeofenceResponse) GetPreUpgradeGeometryOk() (*map[string]interface{}, bool)`

GetPreUpgradeGeometryOk returns a tuple with the PreUpgradeGeometry field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPreUpgradeGeometry

`func (o *GeofenceResponse) SetPreUpgradeGeometry(v map[string]interface{})`

SetPreUpgradeGeometry sets PreUpgradeGeometry field to given value.

### HasPreUpgradeGeometry

`func (o *GeofenceResponse) HasPreUpgradeGeometry() bool`

HasPreUpgradeGeometry returns a boolean if a field has been set.

### SetPreUpgradeGeometryNil

`func (o *GeofenceResponse) SetPreUpgradeGeometryNil(b bool)`

 SetPreUpgradeGeometryNil sets the value for PreUpgradeGeometry to be an explicit nil

### UnsetPreUpgradeGeometry
`func (o *GeofenceResponse) UnsetPreUpgradeGeometry()`

UnsetPreUpgradeGeometry ensures that no value is present for PreUpgradeGeometry, not even an explicit nil
### GetBuildingProvenance

`func (o *GeofenceResponse) GetBuildingProvenance() map[string]interface{}`

GetBuildingProvenance returns the BuildingProvenance field if non-nil, zero value otherwise.

### GetBuildingProvenanceOk

`func (o *GeofenceResponse) GetBuildingProvenanceOk() (*map[string]interface{}, bool)`

GetBuildingProvenanceOk returns a tuple with the BuildingProvenance field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBuildingProvenance

`func (o *GeofenceResponse) SetBuildingProvenance(v map[string]interface{})`

SetBuildingProvenance sets BuildingProvenance field to given value.

### HasBuildingProvenance

`func (o *GeofenceResponse) HasBuildingProvenance() bool`

HasBuildingProvenance returns a boolean if a field has been set.

### SetBuildingProvenanceNil

`func (o *GeofenceResponse) SetBuildingProvenanceNil(b bool)`

 SetBuildingProvenanceNil sets the value for BuildingProvenance to be an explicit nil

### UnsetBuildingProvenance
`func (o *GeofenceResponse) UnsetBuildingProvenance()`

UnsetBuildingProvenance ensures that no value is present for BuildingProvenance, not even an explicit nil
### GetEffectiveRadiusMeters

`func (o *GeofenceResponse) GetEffectiveRadiusMeters() float32`

GetEffectiveRadiusMeters returns the EffectiveRadiusMeters field if non-nil, zero value otherwise.

### GetEffectiveRadiusMetersOk

`func (o *GeofenceResponse) GetEffectiveRadiusMetersOk() (*float32, bool)`

GetEffectiveRadiusMetersOk returns a tuple with the EffectiveRadiusMeters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEffectiveRadiusMeters

`func (o *GeofenceResponse) SetEffectiveRadiusMeters(v float32)`

SetEffectiveRadiusMeters sets EffectiveRadiusMeters field to given value.

### HasEffectiveRadiusMeters

`func (o *GeofenceResponse) HasEffectiveRadiusMeters() bool`

HasEffectiveRadiusMeters returns a boolean if a field has been set.

### SetEffectiveRadiusMetersNil

`func (o *GeofenceResponse) SetEffectiveRadiusMetersNil(b bool)`

 SetEffectiveRadiusMetersNil sets the value for EffectiveRadiusMeters to be an explicit nil

### UnsetEffectiveRadiusMeters
`func (o *GeofenceResponse) UnsetEffectiveRadiusMeters()`

UnsetEffectiveRadiusMeters ensures that no value is present for EffectiveRadiusMeters, not even an explicit nil
### GetBelowMinTriggerRadius

`func (o *GeofenceResponse) GetBelowMinTriggerRadius() bool`

GetBelowMinTriggerRadius returns the BelowMinTriggerRadius field if non-nil, zero value otherwise.

### GetBelowMinTriggerRadiusOk

`func (o *GeofenceResponse) GetBelowMinTriggerRadiusOk() (*bool, bool)`

GetBelowMinTriggerRadiusOk returns a tuple with the BelowMinTriggerRadius field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBelowMinTriggerRadius

`func (o *GeofenceResponse) SetBelowMinTriggerRadius(v bool)`

SetBelowMinTriggerRadius sets BelowMinTriggerRadius field to given value.

### HasBelowMinTriggerRadius

`func (o *GeofenceResponse) HasBelowMinTriggerRadius() bool`

HasBelowMinTriggerRadius returns a boolean if a field has been set.

### GetIsExample

`func (o *GeofenceResponse) GetIsExample() bool`

GetIsExample returns the IsExample field if non-nil, zero value otherwise.

### GetIsExampleOk

`func (o *GeofenceResponse) GetIsExampleOk() (*bool, bool)`

GetIsExampleOk returns a tuple with the IsExample field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsExample

`func (o *GeofenceResponse) SetIsExample(v bool)`

SetIsExample sets IsExample field to given value.

### HasIsExample

`func (o *GeofenceResponse) HasIsExample() bool`

HasIsExample returns a boolean if a field has been set.

### GetCreatedAt

`func (o *GeofenceResponse) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *GeofenceResponse) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *GeofenceResponse) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.


### GetUpdatedAt

`func (o *GeofenceResponse) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *GeofenceResponse) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *GeofenceResponse) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


