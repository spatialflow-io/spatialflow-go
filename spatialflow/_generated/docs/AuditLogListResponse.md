# AuditLogListResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Logs** | [**[]AuditLogOut**](AuditLogOut.md) |  | 
**Total** | **int32** |  | 
**Page** | **int32** |  | 
**PageSize** | **int32** |  | 
**HasMore** | **bool** |  | 

## Methods

### NewAuditLogListResponse

`func NewAuditLogListResponse(logs []AuditLogOut, total int32, page int32, pageSize int32, hasMore bool, ) *AuditLogListResponse`

NewAuditLogListResponse instantiates a new AuditLogListResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAuditLogListResponseWithDefaults

`func NewAuditLogListResponseWithDefaults() *AuditLogListResponse`

NewAuditLogListResponseWithDefaults instantiates a new AuditLogListResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLogs

`func (o *AuditLogListResponse) GetLogs() []AuditLogOut`

GetLogs returns the Logs field if non-nil, zero value otherwise.

### GetLogsOk

`func (o *AuditLogListResponse) GetLogsOk() (*[]AuditLogOut, bool)`

GetLogsOk returns a tuple with the Logs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLogs

`func (o *AuditLogListResponse) SetLogs(v []AuditLogOut)`

SetLogs sets Logs field to given value.


### GetTotal

`func (o *AuditLogListResponse) GetTotal() int32`

GetTotal returns the Total field if non-nil, zero value otherwise.

### GetTotalOk

`func (o *AuditLogListResponse) GetTotalOk() (*int32, bool)`

GetTotalOk returns a tuple with the Total field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotal

`func (o *AuditLogListResponse) SetTotal(v int32)`

SetTotal sets Total field to given value.


### GetPage

`func (o *AuditLogListResponse) GetPage() int32`

GetPage returns the Page field if non-nil, zero value otherwise.

### GetPageOk

`func (o *AuditLogListResponse) GetPageOk() (*int32, bool)`

GetPageOk returns a tuple with the Page field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPage

`func (o *AuditLogListResponse) SetPage(v int32)`

SetPage sets Page field to given value.


### GetPageSize

`func (o *AuditLogListResponse) GetPageSize() int32`

GetPageSize returns the PageSize field if non-nil, zero value otherwise.

### GetPageSizeOk

`func (o *AuditLogListResponse) GetPageSizeOk() (*int32, bool)`

GetPageSizeOk returns a tuple with the PageSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPageSize

`func (o *AuditLogListResponse) SetPageSize(v int32)`

SetPageSize sets PageSize field to given value.


### GetHasMore

`func (o *AuditLogListResponse) GetHasMore() bool`

GetHasMore returns the HasMore field if non-nil, zero value otherwise.

### GetHasMoreOk

`func (o *AuditLogListResponse) GetHasMoreOk() (*bool, bool)`

GetHasMoreOk returns a tuple with the HasMore field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasMore

`func (o *AuditLogListResponse) SetHasMore(v bool)`

SetHasMore sets HasMore field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


