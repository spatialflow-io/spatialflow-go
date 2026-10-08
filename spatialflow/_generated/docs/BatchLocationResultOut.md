# BatchLocationResultOut

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Index** | **int32** |  | 
**ClientLocationId** | Pointer to **NullableString** |  | [optional] 
**Status** | **string** |  | 
**ErrorCode** | Pointer to **NullableString** |  | [optional] 
**Retryable** | Pointer to **bool** |  | [optional] [default to false]

## Methods

### NewBatchLocationResultOut

`func NewBatchLocationResultOut(index int32, status string, ) *BatchLocationResultOut`

NewBatchLocationResultOut instantiates a new BatchLocationResultOut object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBatchLocationResultOutWithDefaults

`func NewBatchLocationResultOutWithDefaults() *BatchLocationResultOut`

NewBatchLocationResultOutWithDefaults instantiates a new BatchLocationResultOut object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIndex

`func (o *BatchLocationResultOut) GetIndex() int32`

GetIndex returns the Index field if non-nil, zero value otherwise.

### GetIndexOk

`func (o *BatchLocationResultOut) GetIndexOk() (*int32, bool)`

GetIndexOk returns a tuple with the Index field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIndex

`func (o *BatchLocationResultOut) SetIndex(v int32)`

SetIndex sets Index field to given value.


### GetClientLocationId

`func (o *BatchLocationResultOut) GetClientLocationId() string`

GetClientLocationId returns the ClientLocationId field if non-nil, zero value otherwise.

### GetClientLocationIdOk

`func (o *BatchLocationResultOut) GetClientLocationIdOk() (*string, bool)`

GetClientLocationIdOk returns a tuple with the ClientLocationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientLocationId

`func (o *BatchLocationResultOut) SetClientLocationId(v string)`

SetClientLocationId sets ClientLocationId field to given value.

### HasClientLocationId

`func (o *BatchLocationResultOut) HasClientLocationId() bool`

HasClientLocationId returns a boolean if a field has been set.

### SetClientLocationIdNil

`func (o *BatchLocationResultOut) SetClientLocationIdNil(b bool)`

 SetClientLocationIdNil sets the value for ClientLocationId to be an explicit nil

### UnsetClientLocationId
`func (o *BatchLocationResultOut) UnsetClientLocationId()`

UnsetClientLocationId ensures that no value is present for ClientLocationId, not even an explicit nil
### GetStatus

`func (o *BatchLocationResultOut) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *BatchLocationResultOut) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *BatchLocationResultOut) SetStatus(v string)`

SetStatus sets Status field to given value.


### GetErrorCode

`func (o *BatchLocationResultOut) GetErrorCode() string`

GetErrorCode returns the ErrorCode field if non-nil, zero value otherwise.

### GetErrorCodeOk

`func (o *BatchLocationResultOut) GetErrorCodeOk() (*string, bool)`

GetErrorCodeOk returns a tuple with the ErrorCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrorCode

`func (o *BatchLocationResultOut) SetErrorCode(v string)`

SetErrorCode sets ErrorCode field to given value.

### HasErrorCode

`func (o *BatchLocationResultOut) HasErrorCode() bool`

HasErrorCode returns a boolean if a field has been set.

### SetErrorCodeNil

`func (o *BatchLocationResultOut) SetErrorCodeNil(b bool)`

 SetErrorCodeNil sets the value for ErrorCode to be an explicit nil

### UnsetErrorCode
`func (o *BatchLocationResultOut) UnsetErrorCode()`

UnsetErrorCode ensures that no value is present for ErrorCode, not even an explicit nil
### GetRetryable

`func (o *BatchLocationResultOut) GetRetryable() bool`

GetRetryable returns the Retryable field if non-nil, zero value otherwise.

### GetRetryableOk

`func (o *BatchLocationResultOut) GetRetryableOk() (*bool, bool)`

GetRetryableOk returns a tuple with the Retryable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRetryable

`func (o *BatchLocationResultOut) SetRetryable(v bool)`

SetRetryable sets Retryable field to given value.

### HasRetryable

`func (o *BatchLocationResultOut) HasRetryable() bool`

HasRetryable returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


