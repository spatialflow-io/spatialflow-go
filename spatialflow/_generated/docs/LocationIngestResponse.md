# LocationIngestResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Accepted** | **int32** |  | 
**Rejected** | **int32** |  | 
**Ids** | Pointer to **[]string** |  | [optional] 
**IdempotencyKey** | Pointer to **NullableString** |  | [optional] 
**Errors** | Pointer to **[]map[string]interface{}** |  | [optional] 
**TaskIds** | Pointer to **[]string** |  | [optional] 

## Methods

### NewLocationIngestResponse

`func NewLocationIngestResponse(accepted int32, rejected int32, ) *LocationIngestResponse`

NewLocationIngestResponse instantiates a new LocationIngestResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLocationIngestResponseWithDefaults

`func NewLocationIngestResponseWithDefaults() *LocationIngestResponse`

NewLocationIngestResponseWithDefaults instantiates a new LocationIngestResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccepted

`func (o *LocationIngestResponse) GetAccepted() int32`

GetAccepted returns the Accepted field if non-nil, zero value otherwise.

### GetAcceptedOk

`func (o *LocationIngestResponse) GetAcceptedOk() (*int32, bool)`

GetAcceptedOk returns a tuple with the Accepted field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccepted

`func (o *LocationIngestResponse) SetAccepted(v int32)`

SetAccepted sets Accepted field to given value.


### GetRejected

`func (o *LocationIngestResponse) GetRejected() int32`

GetRejected returns the Rejected field if non-nil, zero value otherwise.

### GetRejectedOk

`func (o *LocationIngestResponse) GetRejectedOk() (*int32, bool)`

GetRejectedOk returns a tuple with the Rejected field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRejected

`func (o *LocationIngestResponse) SetRejected(v int32)`

SetRejected sets Rejected field to given value.


### GetIds

`func (o *LocationIngestResponse) GetIds() []string`

GetIds returns the Ids field if non-nil, zero value otherwise.

### GetIdsOk

`func (o *LocationIngestResponse) GetIdsOk() (*[]string, bool)`

GetIdsOk returns a tuple with the Ids field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIds

`func (o *LocationIngestResponse) SetIds(v []string)`

SetIds sets Ids field to given value.

### HasIds

`func (o *LocationIngestResponse) HasIds() bool`

HasIds returns a boolean if a field has been set.

### SetIdsNil

`func (o *LocationIngestResponse) SetIdsNil(b bool)`

 SetIdsNil sets the value for Ids to be an explicit nil

### UnsetIds
`func (o *LocationIngestResponse) UnsetIds()`

UnsetIds ensures that no value is present for Ids, not even an explicit nil
### GetIdempotencyKey

`func (o *LocationIngestResponse) GetIdempotencyKey() string`

GetIdempotencyKey returns the IdempotencyKey field if non-nil, zero value otherwise.

### GetIdempotencyKeyOk

`func (o *LocationIngestResponse) GetIdempotencyKeyOk() (*string, bool)`

GetIdempotencyKeyOk returns a tuple with the IdempotencyKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdempotencyKey

`func (o *LocationIngestResponse) SetIdempotencyKey(v string)`

SetIdempotencyKey sets IdempotencyKey field to given value.

### HasIdempotencyKey

`func (o *LocationIngestResponse) HasIdempotencyKey() bool`

HasIdempotencyKey returns a boolean if a field has been set.

### SetIdempotencyKeyNil

`func (o *LocationIngestResponse) SetIdempotencyKeyNil(b bool)`

 SetIdempotencyKeyNil sets the value for IdempotencyKey to be an explicit nil

### UnsetIdempotencyKey
`func (o *LocationIngestResponse) UnsetIdempotencyKey()`

UnsetIdempotencyKey ensures that no value is present for IdempotencyKey, not even an explicit nil
### GetErrors

`func (o *LocationIngestResponse) GetErrors() []map[string]interface{}`

GetErrors returns the Errors field if non-nil, zero value otherwise.

### GetErrorsOk

`func (o *LocationIngestResponse) GetErrorsOk() (*[]map[string]interface{}, bool)`

GetErrorsOk returns a tuple with the Errors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrors

`func (o *LocationIngestResponse) SetErrors(v []map[string]interface{})`

SetErrors sets Errors field to given value.

### HasErrors

`func (o *LocationIngestResponse) HasErrors() bool`

HasErrors returns a boolean if a field has been set.

### SetErrorsNil

`func (o *LocationIngestResponse) SetErrorsNil(b bool)`

 SetErrorsNil sets the value for Errors to be an explicit nil

### UnsetErrors
`func (o *LocationIngestResponse) UnsetErrors()`

UnsetErrors ensures that no value is present for Errors, not even an explicit nil
### GetTaskIds

`func (o *LocationIngestResponse) GetTaskIds() []string`

GetTaskIds returns the TaskIds field if non-nil, zero value otherwise.

### GetTaskIdsOk

`func (o *LocationIngestResponse) GetTaskIdsOk() (*[]string, bool)`

GetTaskIdsOk returns a tuple with the TaskIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaskIds

`func (o *LocationIngestResponse) SetTaskIds(v []string)`

SetTaskIds sets TaskIds field to given value.

### HasTaskIds

`func (o *LocationIngestResponse) HasTaskIds() bool`

HasTaskIds returns a boolean if a field has been set.

### SetTaskIdsNil

`func (o *LocationIngestResponse) SetTaskIdsNil(b bool)`

 SetTaskIdsNil sets the value for TaskIds to be an explicit nil

### UnsetTaskIds
`func (o *LocationIngestResponse) UnsetTaskIds()`

UnsetTaskIds ensures that no value is present for TaskIds, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


