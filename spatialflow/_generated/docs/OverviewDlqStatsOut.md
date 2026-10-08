# OverviewDlqStatsOut

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**TotalEntries** | **int32** |  | 
**NotRequeued** | **int32** |  | 
**Requeued** | **int32** |  | 
**TopFailedWebhooks** | [**[]OverviewFailedWebhookOut**](OverviewFailedWebhookOut.md) |  | 
**WorkspaceId** | **string** |  | 

## Methods

### NewOverviewDlqStatsOut

`func NewOverviewDlqStatsOut(totalEntries int32, notRequeued int32, requeued int32, topFailedWebhooks []OverviewFailedWebhookOut, workspaceId string, ) *OverviewDlqStatsOut`

NewOverviewDlqStatsOut instantiates a new OverviewDlqStatsOut object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOverviewDlqStatsOutWithDefaults

`func NewOverviewDlqStatsOutWithDefaults() *OverviewDlqStatsOut`

NewOverviewDlqStatsOutWithDefaults instantiates a new OverviewDlqStatsOut object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTotalEntries

`func (o *OverviewDlqStatsOut) GetTotalEntries() int32`

GetTotalEntries returns the TotalEntries field if non-nil, zero value otherwise.

### GetTotalEntriesOk

`func (o *OverviewDlqStatsOut) GetTotalEntriesOk() (*int32, bool)`

GetTotalEntriesOk returns a tuple with the TotalEntries field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalEntries

`func (o *OverviewDlqStatsOut) SetTotalEntries(v int32)`

SetTotalEntries sets TotalEntries field to given value.


### GetNotRequeued

`func (o *OverviewDlqStatsOut) GetNotRequeued() int32`

GetNotRequeued returns the NotRequeued field if non-nil, zero value otherwise.

### GetNotRequeuedOk

`func (o *OverviewDlqStatsOut) GetNotRequeuedOk() (*int32, bool)`

GetNotRequeuedOk returns a tuple with the NotRequeued field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotRequeued

`func (o *OverviewDlqStatsOut) SetNotRequeued(v int32)`

SetNotRequeued sets NotRequeued field to given value.


### GetRequeued

`func (o *OverviewDlqStatsOut) GetRequeued() int32`

GetRequeued returns the Requeued field if non-nil, zero value otherwise.

### GetRequeuedOk

`func (o *OverviewDlqStatsOut) GetRequeuedOk() (*int32, bool)`

GetRequeuedOk returns a tuple with the Requeued field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequeued

`func (o *OverviewDlqStatsOut) SetRequeued(v int32)`

SetRequeued sets Requeued field to given value.


### GetTopFailedWebhooks

`func (o *OverviewDlqStatsOut) GetTopFailedWebhooks() []OverviewFailedWebhookOut`

GetTopFailedWebhooks returns the TopFailedWebhooks field if non-nil, zero value otherwise.

### GetTopFailedWebhooksOk

`func (o *OverviewDlqStatsOut) GetTopFailedWebhooksOk() (*[]OverviewFailedWebhookOut, bool)`

GetTopFailedWebhooksOk returns a tuple with the TopFailedWebhooks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTopFailedWebhooks

`func (o *OverviewDlqStatsOut) SetTopFailedWebhooks(v []OverviewFailedWebhookOut)`

SetTopFailedWebhooks sets TopFailedWebhooks field to given value.


### GetWorkspaceId

`func (o *OverviewDlqStatsOut) GetWorkspaceId() string`

GetWorkspaceId returns the WorkspaceId field if non-nil, zero value otherwise.

### GetWorkspaceIdOk

`func (o *OverviewDlqStatsOut) GetWorkspaceIdOk() (*string, bool)`

GetWorkspaceIdOk returns a tuple with the WorkspaceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWorkspaceId

`func (o *OverviewDlqStatsOut) SetWorkspaceId(v string)`

SetWorkspaceId sets WorkspaceId field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


