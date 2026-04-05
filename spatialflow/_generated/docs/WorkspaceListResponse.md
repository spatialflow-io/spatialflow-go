# WorkspaceListResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Workspaces** | [**[]WorkspaceListItem**](WorkspaceListItem.md) |  | 
**Total** | **int32** |  | 
**Page** | **int32** |  | 
**Limit** | **int32** |  | 
**Pages** | **int32** |  | 

## Methods

### NewWorkspaceListResponse

`func NewWorkspaceListResponse(workspaces []WorkspaceListItem, total int32, page int32, limit int32, pages int32, ) *WorkspaceListResponse`

NewWorkspaceListResponse instantiates a new WorkspaceListResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWorkspaceListResponseWithDefaults

`func NewWorkspaceListResponseWithDefaults() *WorkspaceListResponse`

NewWorkspaceListResponseWithDefaults instantiates a new WorkspaceListResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetWorkspaces

`func (o *WorkspaceListResponse) GetWorkspaces() []WorkspaceListItem`

GetWorkspaces returns the Workspaces field if non-nil, zero value otherwise.

### GetWorkspacesOk

`func (o *WorkspaceListResponse) GetWorkspacesOk() (*[]WorkspaceListItem, bool)`

GetWorkspacesOk returns a tuple with the Workspaces field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWorkspaces

`func (o *WorkspaceListResponse) SetWorkspaces(v []WorkspaceListItem)`

SetWorkspaces sets Workspaces field to given value.


### GetTotal

`func (o *WorkspaceListResponse) GetTotal() int32`

GetTotal returns the Total field if non-nil, zero value otherwise.

### GetTotalOk

`func (o *WorkspaceListResponse) GetTotalOk() (*int32, bool)`

GetTotalOk returns a tuple with the Total field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotal

`func (o *WorkspaceListResponse) SetTotal(v int32)`

SetTotal sets Total field to given value.


### GetPage

`func (o *WorkspaceListResponse) GetPage() int32`

GetPage returns the Page field if non-nil, zero value otherwise.

### GetPageOk

`func (o *WorkspaceListResponse) GetPageOk() (*int32, bool)`

GetPageOk returns a tuple with the Page field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPage

`func (o *WorkspaceListResponse) SetPage(v int32)`

SetPage sets Page field to given value.


### GetLimit

`func (o *WorkspaceListResponse) GetLimit() int32`

GetLimit returns the Limit field if non-nil, zero value otherwise.

### GetLimitOk

`func (o *WorkspaceListResponse) GetLimitOk() (*int32, bool)`

GetLimitOk returns a tuple with the Limit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLimit

`func (o *WorkspaceListResponse) SetLimit(v int32)`

SetLimit sets Limit field to given value.


### GetPages

`func (o *WorkspaceListResponse) GetPages() int32`

GetPages returns the Pages field if non-nil, zero value otherwise.

### GetPagesOk

`func (o *WorkspaceListResponse) GetPagesOk() (*int32, bool)`

GetPagesOk returns a tuple with the Pages field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPages

`func (o *WorkspaceListResponse) SetPages(v int32)`

SetPages sets Pages field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


