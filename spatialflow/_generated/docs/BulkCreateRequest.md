# BulkCreateRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Items** | [**[]BulkItemCommit**](BulkItemCommit.md) |  | 
**DedupStrategy** | Pointer to **string** |  | [optional] [default to "skip"]

## Methods

### NewBulkCreateRequest

`func NewBulkCreateRequest(items []BulkItemCommit, ) *BulkCreateRequest`

NewBulkCreateRequest instantiates a new BulkCreateRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBulkCreateRequestWithDefaults

`func NewBulkCreateRequestWithDefaults() *BulkCreateRequest`

NewBulkCreateRequestWithDefaults instantiates a new BulkCreateRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetItems

`func (o *BulkCreateRequest) GetItems() []BulkItemCommit`

GetItems returns the Items field if non-nil, zero value otherwise.

### GetItemsOk

`func (o *BulkCreateRequest) GetItemsOk() (*[]BulkItemCommit, bool)`

GetItemsOk returns a tuple with the Items field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItems

`func (o *BulkCreateRequest) SetItems(v []BulkItemCommit)`

SetItems sets Items field to given value.


### GetDedupStrategy

`func (o *BulkCreateRequest) GetDedupStrategy() string`

GetDedupStrategy returns the DedupStrategy field if non-nil, zero value otherwise.

### GetDedupStrategyOk

`func (o *BulkCreateRequest) GetDedupStrategyOk() (*string, bool)`

GetDedupStrategyOk returns a tuple with the DedupStrategy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDedupStrategy

`func (o *BulkCreateRequest) SetDedupStrategy(v string)`

SetDedupStrategy sets DedupStrategy field to given value.

### HasDedupStrategy

`func (o *BulkCreateRequest) HasDedupStrategy() bool`

HasDedupStrategy returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


