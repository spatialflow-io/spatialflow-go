# EnhancedWorkspaceListResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Workspaces** | [**[]EnhancedWorkspaceListItem**](EnhancedWorkspaceListItem.md) |  | 
**Total** | **int32** |  | 
**Page** | **int32** |  | 
**Limit** | **int32** |  | 
**Pages** | **int32** |  | 

## Methods

### NewEnhancedWorkspaceListResponse

`func NewEnhancedWorkspaceListResponse(workspaces []EnhancedWorkspaceListItem, total int32, page int32, limit int32, pages int32, ) *EnhancedWorkspaceListResponse`

NewEnhancedWorkspaceListResponse instantiates a new EnhancedWorkspaceListResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEnhancedWorkspaceListResponseWithDefaults

`func NewEnhancedWorkspaceListResponseWithDefaults() *EnhancedWorkspaceListResponse`

NewEnhancedWorkspaceListResponseWithDefaults instantiates a new EnhancedWorkspaceListResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetWorkspaces

`func (o *EnhancedWorkspaceListResponse) GetWorkspaces() []EnhancedWorkspaceListItem`

GetWorkspaces returns the Workspaces field if non-nil, zero value otherwise.

### GetWorkspacesOk

`func (o *EnhancedWorkspaceListResponse) GetWorkspacesOk() (*[]EnhancedWorkspaceListItem, bool)`

GetWorkspacesOk returns a tuple with the Workspaces field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWorkspaces

`func (o *EnhancedWorkspaceListResponse) SetWorkspaces(v []EnhancedWorkspaceListItem)`

SetWorkspaces sets Workspaces field to given value.


### GetTotal

`func (o *EnhancedWorkspaceListResponse) GetTotal() int32`

GetTotal returns the Total field if non-nil, zero value otherwise.

### GetTotalOk

`func (o *EnhancedWorkspaceListResponse) GetTotalOk() (*int32, bool)`

GetTotalOk returns a tuple with the Total field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotal

`func (o *EnhancedWorkspaceListResponse) SetTotal(v int32)`

SetTotal sets Total field to given value.


### GetPage

`func (o *EnhancedWorkspaceListResponse) GetPage() int32`

GetPage returns the Page field if non-nil, zero value otherwise.

### GetPageOk

`func (o *EnhancedWorkspaceListResponse) GetPageOk() (*int32, bool)`

GetPageOk returns a tuple with the Page field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPage

`func (o *EnhancedWorkspaceListResponse) SetPage(v int32)`

SetPage sets Page field to given value.


### GetLimit

`func (o *EnhancedWorkspaceListResponse) GetLimit() int32`

GetLimit returns the Limit field if non-nil, zero value otherwise.

### GetLimitOk

`func (o *EnhancedWorkspaceListResponse) GetLimitOk() (*int32, bool)`

GetLimitOk returns a tuple with the Limit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLimit

`func (o *EnhancedWorkspaceListResponse) SetLimit(v int32)`

SetLimit sets Limit field to given value.


### GetPages

`func (o *EnhancedWorkspaceListResponse) GetPages() int32`

GetPages returns the Pages field if non-nil, zero value otherwise.

### GetPagesOk

`func (o *EnhancedWorkspaceListResponse) GetPagesOk() (*int32, bool)`

GetPagesOk returns a tuple with the Pages field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPages

`func (o *EnhancedWorkspaceListResponse) SetPages(v int32)`

SetPages sets Pages field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


