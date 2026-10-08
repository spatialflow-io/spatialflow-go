# BulkPreviewRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Items** | [**[]BulkItemInput**](BulkItemInput.md) |  | 

## Methods

### NewBulkPreviewRequest

`func NewBulkPreviewRequest(items []BulkItemInput, ) *BulkPreviewRequest`

NewBulkPreviewRequest instantiates a new BulkPreviewRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBulkPreviewRequestWithDefaults

`func NewBulkPreviewRequestWithDefaults() *BulkPreviewRequest`

NewBulkPreviewRequestWithDefaults instantiates a new BulkPreviewRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetItems

`func (o *BulkPreviewRequest) GetItems() []BulkItemInput`

GetItems returns the Items field if non-nil, zero value otherwise.

### GetItemsOk

`func (o *BulkPreviewRequest) GetItemsOk() (*[]BulkItemInput, bool)`

GetItemsOk returns a tuple with the Items field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItems

`func (o *BulkPreviewRequest) SetItems(v []BulkItemInput)`

SetItems sets Items field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


