# RateLimitResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Error** | **string** |  | 
**Limit** | **int32** |  | 
**ResetAt** | **string** |  | 

## Methods

### NewRateLimitResponse

`func NewRateLimitResponse(error_ string, limit int32, resetAt string, ) *RateLimitResponse`

NewRateLimitResponse instantiates a new RateLimitResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRateLimitResponseWithDefaults

`func NewRateLimitResponseWithDefaults() *RateLimitResponse`

NewRateLimitResponseWithDefaults instantiates a new RateLimitResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetError

`func (o *RateLimitResponse) GetError() string`

GetError returns the Error field if non-nil, zero value otherwise.

### GetErrorOk

`func (o *RateLimitResponse) GetErrorOk() (*string, bool)`

GetErrorOk returns a tuple with the Error field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetError

`func (o *RateLimitResponse) SetError(v string)`

SetError sets Error field to given value.


### GetLimit

`func (o *RateLimitResponse) GetLimit() int32`

GetLimit returns the Limit field if non-nil, zero value otherwise.

### GetLimitOk

`func (o *RateLimitResponse) GetLimitOk() (*int32, bool)`

GetLimitOk returns a tuple with the Limit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLimit

`func (o *RateLimitResponse) SetLimit(v int32)`

SetLimit sets Limit field to given value.


### GetResetAt

`func (o *RateLimitResponse) GetResetAt() string`

GetResetAt returns the ResetAt field if non-nil, zero value otherwise.

### GetResetAtOk

`func (o *RateLimitResponse) GetResetAtOk() (*string, bool)`

GetResetAtOk returns a tuple with the ResetAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResetAt

`func (o *RateLimitResponse) SetResetAt(v string)`

SetResetAt sets ResetAt field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


