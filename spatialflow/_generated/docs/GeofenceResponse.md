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


