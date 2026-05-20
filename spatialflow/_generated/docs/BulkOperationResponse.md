# BulkOperationResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Total** | **int32** |  | 
**Succeeded** | **int32** |  | 
**Failed** | **int32** |  | 
**Results** | [**[]BulkResultItem**](BulkResultItem.md) |  | 

## Methods

### NewBulkOperationResponse

`func NewBulkOperationResponse(total int32, succeeded int32, failed int32, results []BulkResultItem, ) *BulkOperationResponse`

NewBulkOperationResponse instantiates a new BulkOperationResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBulkOperationResponseWithDefaults

`func NewBulkOperationResponseWithDefaults() *BulkOperationResponse`

NewBulkOperationResponseWithDefaults instantiates a new BulkOperationResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTotal

`func (o *BulkOperationResponse) GetTotal() int32`

GetTotal returns the Total field if non-nil, zero value otherwise.

### GetTotalOk

`func (o *BulkOperationResponse) GetTotalOk() (*int32, bool)`

GetTotalOk returns a tuple with the Total field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotal

`func (o *BulkOperationResponse) SetTotal(v int32)`

SetTotal sets Total field to given value.


### GetSucceeded

`func (o *BulkOperationResponse) GetSucceeded() int32`

GetSucceeded returns the Succeeded field if non-nil, zero value otherwise.

### GetSucceededOk

`func (o *BulkOperationResponse) GetSucceededOk() (*int32, bool)`

GetSucceededOk returns a tuple with the Succeeded field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSucceeded

`func (o *BulkOperationResponse) SetSucceeded(v int32)`

SetSucceeded sets Succeeded field to given value.


### GetFailed

`func (o *BulkOperationResponse) GetFailed() int32`

GetFailed returns the Failed field if non-nil, zero value otherwise.

### GetFailedOk

`func (o *BulkOperationResponse) GetFailedOk() (*int32, bool)`

GetFailedOk returns a tuple with the Failed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFailed

`func (o *BulkOperationResponse) SetFailed(v int32)`

SetFailed sets Failed field to given value.


### GetResults

`func (o *BulkOperationResponse) GetResults() []BulkResultItem`

GetResults returns the Results field if non-nil, zero value otherwise.

### GetResultsOk

`func (o *BulkOperationResponse) GetResultsOk() (*[]BulkResultItem, bool)`

GetResultsOk returns a tuple with the Results field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResults

`func (o *BulkOperationResponse) SetResults(v []BulkResultItem)`

SetResults sets Results field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


