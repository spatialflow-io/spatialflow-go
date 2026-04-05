# RouteTestErrorOut

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Detail** | **string** |  | 
**ErrorCode** | **string** |  | 
**LineNumber** | Pointer to **NullableInt32** |  | [optional] 

## Methods

### NewRouteTestErrorOut

`func NewRouteTestErrorOut(detail string, errorCode string, ) *RouteTestErrorOut`

NewRouteTestErrorOut instantiates a new RouteTestErrorOut object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRouteTestErrorOutWithDefaults

`func NewRouteTestErrorOutWithDefaults() *RouteTestErrorOut`

NewRouteTestErrorOutWithDefaults instantiates a new RouteTestErrorOut object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDetail

`func (o *RouteTestErrorOut) GetDetail() string`

GetDetail returns the Detail field if non-nil, zero value otherwise.

### GetDetailOk

`func (o *RouteTestErrorOut) GetDetailOk() (*string, bool)`

GetDetailOk returns a tuple with the Detail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDetail

`func (o *RouteTestErrorOut) SetDetail(v string)`

SetDetail sets Detail field to given value.


### GetErrorCode

`func (o *RouteTestErrorOut) GetErrorCode() string`

GetErrorCode returns the ErrorCode field if non-nil, zero value otherwise.

### GetErrorCodeOk

`func (o *RouteTestErrorOut) GetErrorCodeOk() (*string, bool)`

GetErrorCodeOk returns a tuple with the ErrorCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrorCode

`func (o *RouteTestErrorOut) SetErrorCode(v string)`

SetErrorCode sets ErrorCode field to given value.


### GetLineNumber

`func (o *RouteTestErrorOut) GetLineNumber() int32`

GetLineNumber returns the LineNumber field if non-nil, zero value otherwise.

### GetLineNumberOk

`func (o *RouteTestErrorOut) GetLineNumberOk() (*int32, bool)`

GetLineNumberOk returns a tuple with the LineNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLineNumber

`func (o *RouteTestErrorOut) SetLineNumber(v int32)`

SetLineNumber sets LineNumber field to given value.

### HasLineNumber

`func (o *RouteTestErrorOut) HasLineNumber() bool`

HasLineNumber returns a boolean if a field has been set.

### SetLineNumberNil

`func (o *RouteTestErrorOut) SetLineNumberNil(b bool)`

 SetLineNumberNil sets the value for LineNumber to be an explicit nil

### UnsetLineNumber
`func (o *RouteTestErrorOut) UnsetLineNumber()`

UnsetLineNumber ensures that no value is present for LineNumber, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


