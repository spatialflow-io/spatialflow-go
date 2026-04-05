# RecentActivity

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ActivityType** | **string** |  | 
**Timestamp** | **string** |  | 
**Metadata** | **map[string]interface{}** |  | 

## Methods

### NewRecentActivity

`func NewRecentActivity(activityType string, timestamp string, metadata map[string]interface{}, ) *RecentActivity`

NewRecentActivity instantiates a new RecentActivity object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRecentActivityWithDefaults

`func NewRecentActivityWithDefaults() *RecentActivity`

NewRecentActivityWithDefaults instantiates a new RecentActivity object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetActivityType

`func (o *RecentActivity) GetActivityType() string`

GetActivityType returns the ActivityType field if non-nil, zero value otherwise.

### GetActivityTypeOk

`func (o *RecentActivity) GetActivityTypeOk() (*string, bool)`

GetActivityTypeOk returns a tuple with the ActivityType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActivityType

`func (o *RecentActivity) SetActivityType(v string)`

SetActivityType sets ActivityType field to given value.


### GetTimestamp

`func (o *RecentActivity) GetTimestamp() string`

GetTimestamp returns the Timestamp field if non-nil, zero value otherwise.

### GetTimestampOk

`func (o *RecentActivity) GetTimestampOk() (*string, bool)`

GetTimestampOk returns a tuple with the Timestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimestamp

`func (o *RecentActivity) SetTimestamp(v string)`

SetTimestamp sets Timestamp field to given value.


### GetMetadata

`func (o *RecentActivity) GetMetadata() map[string]interface{}`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *RecentActivity) GetMetadataOk() (*map[string]interface{}, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *RecentActivity) SetMetadata(v map[string]interface{})`

SetMetadata sets Metadata field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


