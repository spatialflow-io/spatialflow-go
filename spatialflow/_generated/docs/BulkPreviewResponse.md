# BulkPreviewResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Items** | [**[]BulkItemPreview**](BulkItemPreview.md) |  | 
**Total** | **int32** |  | 
**OkCount** | **int32** |  | 
**DedupCount** | **int32** |  | 
**ErrorCount** | **int32** |  | 

## Methods

### NewBulkPreviewResponse

`func NewBulkPreviewResponse(items []BulkItemPreview, total int32, okCount int32, dedupCount int32, errorCount int32, ) *BulkPreviewResponse`

NewBulkPreviewResponse instantiates a new BulkPreviewResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBulkPreviewResponseWithDefaults

`func NewBulkPreviewResponseWithDefaults() *BulkPreviewResponse`

NewBulkPreviewResponseWithDefaults instantiates a new BulkPreviewResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetItems

`func (o *BulkPreviewResponse) GetItems() []BulkItemPreview`

GetItems returns the Items field if non-nil, zero value otherwise.

### GetItemsOk

`func (o *BulkPreviewResponse) GetItemsOk() (*[]BulkItemPreview, bool)`

GetItemsOk returns a tuple with the Items field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItems

`func (o *BulkPreviewResponse) SetItems(v []BulkItemPreview)`

SetItems sets Items field to given value.


### GetTotal

`func (o *BulkPreviewResponse) GetTotal() int32`

GetTotal returns the Total field if non-nil, zero value otherwise.

### GetTotalOk

`func (o *BulkPreviewResponse) GetTotalOk() (*int32, bool)`

GetTotalOk returns a tuple with the Total field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotal

`func (o *BulkPreviewResponse) SetTotal(v int32)`

SetTotal sets Total field to given value.


### GetOkCount

`func (o *BulkPreviewResponse) GetOkCount() int32`

GetOkCount returns the OkCount field if non-nil, zero value otherwise.

### GetOkCountOk

`func (o *BulkPreviewResponse) GetOkCountOk() (*int32, bool)`

GetOkCountOk returns a tuple with the OkCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOkCount

`func (o *BulkPreviewResponse) SetOkCount(v int32)`

SetOkCount sets OkCount field to given value.


### GetDedupCount

`func (o *BulkPreviewResponse) GetDedupCount() int32`

GetDedupCount returns the DedupCount field if non-nil, zero value otherwise.

### GetDedupCountOk

`func (o *BulkPreviewResponse) GetDedupCountOk() (*int32, bool)`

GetDedupCountOk returns a tuple with the DedupCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDedupCount

`func (o *BulkPreviewResponse) SetDedupCount(v int32)`

SetDedupCount sets DedupCount field to given value.


### GetErrorCount

`func (o *BulkPreviewResponse) GetErrorCount() int32`

GetErrorCount returns the ErrorCount field if non-nil, zero value otherwise.

### GetErrorCountOk

`func (o *BulkPreviewResponse) GetErrorCountOk() (*int32, bool)`

GetErrorCountOk returns a tuple with the ErrorCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrorCount

`func (o *BulkPreviewResponse) SetErrorCount(v int32)`

SetErrorCount sets ErrorCount field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


