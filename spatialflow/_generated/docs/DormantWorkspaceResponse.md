# DormantWorkspaceResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Workspaces** | [**[]DormantWorkspaceItem**](DormantWorkspaceItem.md) |  | 
**Total** | **int32** |  | 
**InactiveDaysThreshold** | **int32** |  | 

## Methods

### NewDormantWorkspaceResponse

`func NewDormantWorkspaceResponse(workspaces []DormantWorkspaceItem, total int32, inactiveDaysThreshold int32, ) *DormantWorkspaceResponse`

NewDormantWorkspaceResponse instantiates a new DormantWorkspaceResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDormantWorkspaceResponseWithDefaults

`func NewDormantWorkspaceResponseWithDefaults() *DormantWorkspaceResponse`

NewDormantWorkspaceResponseWithDefaults instantiates a new DormantWorkspaceResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetWorkspaces

`func (o *DormantWorkspaceResponse) GetWorkspaces() []DormantWorkspaceItem`

GetWorkspaces returns the Workspaces field if non-nil, zero value otherwise.

### GetWorkspacesOk

`func (o *DormantWorkspaceResponse) GetWorkspacesOk() (*[]DormantWorkspaceItem, bool)`

GetWorkspacesOk returns a tuple with the Workspaces field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWorkspaces

`func (o *DormantWorkspaceResponse) SetWorkspaces(v []DormantWorkspaceItem)`

SetWorkspaces sets Workspaces field to given value.


### GetTotal

`func (o *DormantWorkspaceResponse) GetTotal() int32`

GetTotal returns the Total field if non-nil, zero value otherwise.

### GetTotalOk

`func (o *DormantWorkspaceResponse) GetTotalOk() (*int32, bool)`

GetTotalOk returns a tuple with the Total field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotal

`func (o *DormantWorkspaceResponse) SetTotal(v int32)`

SetTotal sets Total field to given value.


### GetInactiveDaysThreshold

`func (o *DormantWorkspaceResponse) GetInactiveDaysThreshold() int32`

GetInactiveDaysThreshold returns the InactiveDaysThreshold field if non-nil, zero value otherwise.

### GetInactiveDaysThresholdOk

`func (o *DormantWorkspaceResponse) GetInactiveDaysThresholdOk() (*int32, bool)`

GetInactiveDaysThresholdOk returns a tuple with the InactiveDaysThreshold field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInactiveDaysThreshold

`func (o *DormantWorkspaceResponse) SetInactiveDaysThreshold(v int32)`

SetInactiveDaysThreshold sets InactiveDaysThreshold field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


