# BulkResult

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Index** | **int32** |  | 
**Status** | **string** |  | 
**GeofenceId** | Pointer to **NullableString** |  | [optional] 
**ErrorCode** | Pointer to **NullableString** |  | [optional] 
**Error** | Pointer to **map[string]string** |  | [optional] 
**DedupMatch** | Pointer to **map[string]interface{}** |  | [optional] 

## Methods

### NewBulkResult

`func NewBulkResult(index int32, status string, ) *BulkResult`

NewBulkResult instantiates a new BulkResult object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBulkResultWithDefaults

`func NewBulkResultWithDefaults() *BulkResult`

NewBulkResultWithDefaults instantiates a new BulkResult object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIndex

`func (o *BulkResult) GetIndex() int32`

GetIndex returns the Index field if non-nil, zero value otherwise.

### GetIndexOk

`func (o *BulkResult) GetIndexOk() (*int32, bool)`

GetIndexOk returns a tuple with the Index field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIndex

`func (o *BulkResult) SetIndex(v int32)`

SetIndex sets Index field to given value.


### GetStatus

`func (o *BulkResult) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *BulkResult) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *BulkResult) SetStatus(v string)`

SetStatus sets Status field to given value.


### GetGeofenceId

`func (o *BulkResult) GetGeofenceId() string`

GetGeofenceId returns the GeofenceId field if non-nil, zero value otherwise.

### GetGeofenceIdOk

`func (o *BulkResult) GetGeofenceIdOk() (*string, bool)`

GetGeofenceIdOk returns a tuple with the GeofenceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGeofenceId

`func (o *BulkResult) SetGeofenceId(v string)`

SetGeofenceId sets GeofenceId field to given value.

### HasGeofenceId

`func (o *BulkResult) HasGeofenceId() bool`

HasGeofenceId returns a boolean if a field has been set.

### SetGeofenceIdNil

`func (o *BulkResult) SetGeofenceIdNil(b bool)`

 SetGeofenceIdNil sets the value for GeofenceId to be an explicit nil

### UnsetGeofenceId
`func (o *BulkResult) UnsetGeofenceId()`

UnsetGeofenceId ensures that no value is present for GeofenceId, not even an explicit nil
### GetErrorCode

`func (o *BulkResult) GetErrorCode() string`

GetErrorCode returns the ErrorCode field if non-nil, zero value otherwise.

### GetErrorCodeOk

`func (o *BulkResult) GetErrorCodeOk() (*string, bool)`

GetErrorCodeOk returns a tuple with the ErrorCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrorCode

`func (o *BulkResult) SetErrorCode(v string)`

SetErrorCode sets ErrorCode field to given value.

### HasErrorCode

`func (o *BulkResult) HasErrorCode() bool`

HasErrorCode returns a boolean if a field has been set.

### SetErrorCodeNil

`func (o *BulkResult) SetErrorCodeNil(b bool)`

 SetErrorCodeNil sets the value for ErrorCode to be an explicit nil

### UnsetErrorCode
`func (o *BulkResult) UnsetErrorCode()`

UnsetErrorCode ensures that no value is present for ErrorCode, not even an explicit nil
### GetError

`func (o *BulkResult) GetError() map[string]string`

GetError returns the Error field if non-nil, zero value otherwise.

### GetErrorOk

`func (o *BulkResult) GetErrorOk() (*map[string]string, bool)`

GetErrorOk returns a tuple with the Error field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetError

`func (o *BulkResult) SetError(v map[string]string)`

SetError sets Error field to given value.

### HasError

`func (o *BulkResult) HasError() bool`

HasError returns a boolean if a field has been set.

### SetErrorNil

`func (o *BulkResult) SetErrorNil(b bool)`

 SetErrorNil sets the value for Error to be an explicit nil

### UnsetError
`func (o *BulkResult) UnsetError()`

UnsetError ensures that no value is present for Error, not even an explicit nil
### GetDedupMatch

`func (o *BulkResult) GetDedupMatch() map[string]interface{}`

GetDedupMatch returns the DedupMatch field if non-nil, zero value otherwise.

### GetDedupMatchOk

`func (o *BulkResult) GetDedupMatchOk() (*map[string]interface{}, bool)`

GetDedupMatchOk returns a tuple with the DedupMatch field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDedupMatch

`func (o *BulkResult) SetDedupMatch(v map[string]interface{})`

SetDedupMatch sets DedupMatch field to given value.

### HasDedupMatch

`func (o *BulkResult) HasDedupMatch() bool`

HasDedupMatch returns a boolean if a field has been set.

### SetDedupMatchNil

`func (o *BulkResult) SetDedupMatchNil(b bool)`

 SetDedupMatchNil sets the value for DedupMatch to be an explicit nil

### UnsetDedupMatch
`func (o *BulkResult) UnsetDedupMatch()`

UnsetDedupMatch ensures that no value is present for DedupMatch, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


