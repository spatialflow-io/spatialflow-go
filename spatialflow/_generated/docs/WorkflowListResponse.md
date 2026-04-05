# WorkflowListResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Workflows** | [**[]WorkflowListOut**](WorkflowListOut.md) |  | 
**Total** | **int32** |  | 
**Page** | **int32** |  | 
**PageSize** | **int32** |  | 

## Methods

### NewWorkflowListResponse

`func NewWorkflowListResponse(workflows []WorkflowListOut, total int32, page int32, pageSize int32, ) *WorkflowListResponse`

NewWorkflowListResponse instantiates a new WorkflowListResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWorkflowListResponseWithDefaults

`func NewWorkflowListResponseWithDefaults() *WorkflowListResponse`

NewWorkflowListResponseWithDefaults instantiates a new WorkflowListResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetWorkflows

`func (o *WorkflowListResponse) GetWorkflows() []WorkflowListOut`

GetWorkflows returns the Workflows field if non-nil, zero value otherwise.

### GetWorkflowsOk

`func (o *WorkflowListResponse) GetWorkflowsOk() (*[]WorkflowListOut, bool)`

GetWorkflowsOk returns a tuple with the Workflows field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWorkflows

`func (o *WorkflowListResponse) SetWorkflows(v []WorkflowListOut)`

SetWorkflows sets Workflows field to given value.


### GetTotal

`func (o *WorkflowListResponse) GetTotal() int32`

GetTotal returns the Total field if non-nil, zero value otherwise.

### GetTotalOk

`func (o *WorkflowListResponse) GetTotalOk() (*int32, bool)`

GetTotalOk returns a tuple with the Total field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotal

`func (o *WorkflowListResponse) SetTotal(v int32)`

SetTotal sets Total field to given value.


### GetPage

`func (o *WorkflowListResponse) GetPage() int32`

GetPage returns the Page field if non-nil, zero value otherwise.

### GetPageOk

`func (o *WorkflowListResponse) GetPageOk() (*int32, bool)`

GetPageOk returns a tuple with the Page field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPage

`func (o *WorkflowListResponse) SetPage(v int32)`

SetPage sets Page field to given value.


### GetPageSize

`func (o *WorkflowListResponse) GetPageSize() int32`

GetPageSize returns the PageSize field if non-nil, zero value otherwise.

### GetPageSizeOk

`func (o *WorkflowListResponse) GetPageSizeOk() (*int32, bool)`

GetPageSizeOk returns a tuple with the PageSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPageSize

`func (o *WorkflowListResponse) SetPageSize(v int32)`

SetPageSize sets PageSize field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


