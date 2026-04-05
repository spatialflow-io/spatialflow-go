# WorkspaceMembersResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Members** | [**[]MemberSummary**](MemberSummary.md) |  | 
**Total** | **int32** |  | 
**Page** | **int32** |  | 
**Limit** | **int32** |  | 
**Pages** | **int32** |  | 
**Workspace** | [**WorkspaceSummary**](WorkspaceSummary.md) |  | 

## Methods

### NewWorkspaceMembersResponse

`func NewWorkspaceMembersResponse(members []MemberSummary, total int32, page int32, limit int32, pages int32, workspace WorkspaceSummary, ) *WorkspaceMembersResponse`

NewWorkspaceMembersResponse instantiates a new WorkspaceMembersResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWorkspaceMembersResponseWithDefaults

`func NewWorkspaceMembersResponseWithDefaults() *WorkspaceMembersResponse`

NewWorkspaceMembersResponseWithDefaults instantiates a new WorkspaceMembersResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMembers

`func (o *WorkspaceMembersResponse) GetMembers() []MemberSummary`

GetMembers returns the Members field if non-nil, zero value otherwise.

### GetMembersOk

`func (o *WorkspaceMembersResponse) GetMembersOk() (*[]MemberSummary, bool)`

GetMembersOk returns a tuple with the Members field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMembers

`func (o *WorkspaceMembersResponse) SetMembers(v []MemberSummary)`

SetMembers sets Members field to given value.


### GetTotal

`func (o *WorkspaceMembersResponse) GetTotal() int32`

GetTotal returns the Total field if non-nil, zero value otherwise.

### GetTotalOk

`func (o *WorkspaceMembersResponse) GetTotalOk() (*int32, bool)`

GetTotalOk returns a tuple with the Total field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotal

`func (o *WorkspaceMembersResponse) SetTotal(v int32)`

SetTotal sets Total field to given value.


### GetPage

`func (o *WorkspaceMembersResponse) GetPage() int32`

GetPage returns the Page field if non-nil, zero value otherwise.

### GetPageOk

`func (o *WorkspaceMembersResponse) GetPageOk() (*int32, bool)`

GetPageOk returns a tuple with the Page field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPage

`func (o *WorkspaceMembersResponse) SetPage(v int32)`

SetPage sets Page field to given value.


### GetLimit

`func (o *WorkspaceMembersResponse) GetLimit() int32`

GetLimit returns the Limit field if non-nil, zero value otherwise.

### GetLimitOk

`func (o *WorkspaceMembersResponse) GetLimitOk() (*int32, bool)`

GetLimitOk returns a tuple with the Limit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLimit

`func (o *WorkspaceMembersResponse) SetLimit(v int32)`

SetLimit sets Limit field to given value.


### GetPages

`func (o *WorkspaceMembersResponse) GetPages() int32`

GetPages returns the Pages field if non-nil, zero value otherwise.

### GetPagesOk

`func (o *WorkspaceMembersResponse) GetPagesOk() (*int32, bool)`

GetPagesOk returns a tuple with the Pages field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPages

`func (o *WorkspaceMembersResponse) SetPages(v int32)`

SetPages sets Pages field to given value.


### GetWorkspace

`func (o *WorkspaceMembersResponse) GetWorkspace() WorkspaceSummary`

GetWorkspace returns the Workspace field if non-nil, zero value otherwise.

### GetWorkspaceOk

`func (o *WorkspaceMembersResponse) GetWorkspaceOk() (*WorkspaceSummary, bool)`

GetWorkspaceOk returns a tuple with the Workspace field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWorkspace

`func (o *WorkspaceMembersResponse) SetWorkspace(v WorkspaceSummary)`

SetWorkspace sets Workspace field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


