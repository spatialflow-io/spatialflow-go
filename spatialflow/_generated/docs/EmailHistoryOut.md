# EmailHistoryOut

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Items** | [**[]EmailHistoryItemOut**](EmailHistoryItemOut.md) |  | 
**Total** | **int32** |  | 
**Limit** | **int32** |  | 
**Offset** | **int32** |  | 

## Methods

### NewEmailHistoryOut

`func NewEmailHistoryOut(items []EmailHistoryItemOut, total int32, limit int32, offset int32, ) *EmailHistoryOut`

NewEmailHistoryOut instantiates a new EmailHistoryOut object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEmailHistoryOutWithDefaults

`func NewEmailHistoryOutWithDefaults() *EmailHistoryOut`

NewEmailHistoryOutWithDefaults instantiates a new EmailHistoryOut object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetItems

`func (o *EmailHistoryOut) GetItems() []EmailHistoryItemOut`

GetItems returns the Items field if non-nil, zero value otherwise.

### GetItemsOk

`func (o *EmailHistoryOut) GetItemsOk() (*[]EmailHistoryItemOut, bool)`

GetItemsOk returns a tuple with the Items field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItems

`func (o *EmailHistoryOut) SetItems(v []EmailHistoryItemOut)`

SetItems sets Items field to given value.


### GetTotal

`func (o *EmailHistoryOut) GetTotal() int32`

GetTotal returns the Total field if non-nil, zero value otherwise.

### GetTotalOk

`func (o *EmailHistoryOut) GetTotalOk() (*int32, bool)`

GetTotalOk returns a tuple with the Total field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotal

`func (o *EmailHistoryOut) SetTotal(v int32)`

SetTotal sets Total field to given value.


### GetLimit

`func (o *EmailHistoryOut) GetLimit() int32`

GetLimit returns the Limit field if non-nil, zero value otherwise.

### GetLimitOk

`func (o *EmailHistoryOut) GetLimitOk() (*int32, bool)`

GetLimitOk returns a tuple with the Limit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLimit

`func (o *EmailHistoryOut) SetLimit(v int32)`

SetLimit sets Limit field to given value.


### GetOffset

`func (o *EmailHistoryOut) GetOffset() int32`

GetOffset returns the Offset field if non-nil, zero value otherwise.

### GetOffsetOk

`func (o *EmailHistoryOut) GetOffsetOk() (*int32, bool)`

GetOffsetOk returns a tuple with the Offset field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOffset

`func (o *EmailHistoryOut) SetOffset(v int32)`

SetOffset sets Offset field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


