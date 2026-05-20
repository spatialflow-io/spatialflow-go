# WorkspaceAnalyticsResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Workspaces** | [**[]WorkspaceAnalyticsItem**](WorkspaceAnalyticsItem.md) |  | 
**Total** | **int32** |  | 
**Limit** | **int32** |  | 
**Offset** | **int32** |  | 

## Methods

### NewWorkspaceAnalyticsResponse

`func NewWorkspaceAnalyticsResponse(workspaces []WorkspaceAnalyticsItem, total int32, limit int32, offset int32, ) *WorkspaceAnalyticsResponse`

NewWorkspaceAnalyticsResponse instantiates a new WorkspaceAnalyticsResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWorkspaceAnalyticsResponseWithDefaults

`func NewWorkspaceAnalyticsResponseWithDefaults() *WorkspaceAnalyticsResponse`

NewWorkspaceAnalyticsResponseWithDefaults instantiates a new WorkspaceAnalyticsResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetWorkspaces

`func (o *WorkspaceAnalyticsResponse) GetWorkspaces() []WorkspaceAnalyticsItem`

GetWorkspaces returns the Workspaces field if non-nil, zero value otherwise.

### GetWorkspacesOk

`func (o *WorkspaceAnalyticsResponse) GetWorkspacesOk() (*[]WorkspaceAnalyticsItem, bool)`

GetWorkspacesOk returns a tuple with the Workspaces field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWorkspaces

`func (o *WorkspaceAnalyticsResponse) SetWorkspaces(v []WorkspaceAnalyticsItem)`

SetWorkspaces sets Workspaces field to given value.


### GetTotal

`func (o *WorkspaceAnalyticsResponse) GetTotal() int32`

GetTotal returns the Total field if non-nil, zero value otherwise.

### GetTotalOk

`func (o *WorkspaceAnalyticsResponse) GetTotalOk() (*int32, bool)`

GetTotalOk returns a tuple with the Total field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotal

`func (o *WorkspaceAnalyticsResponse) SetTotal(v int32)`

SetTotal sets Total field to given value.


### GetLimit

`func (o *WorkspaceAnalyticsResponse) GetLimit() int32`

GetLimit returns the Limit field if non-nil, zero value otherwise.

### GetLimitOk

`func (o *WorkspaceAnalyticsResponse) GetLimitOk() (*int32, bool)`

GetLimitOk returns a tuple with the Limit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLimit

`func (o *WorkspaceAnalyticsResponse) SetLimit(v int32)`

SetLimit sets Limit field to given value.


### GetOffset

`func (o *WorkspaceAnalyticsResponse) GetOffset() int32`

GetOffset returns the Offset field if non-nil, zero value otherwise.

### GetOffsetOk

`func (o *WorkspaceAnalyticsResponse) GetOffsetOk() (*int32, bool)`

GetOffsetOk returns a tuple with the Offset field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOffset

`func (o *WorkspaceAnalyticsResponse) SetOffset(v int32)`

SetOffset sets Offset field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


