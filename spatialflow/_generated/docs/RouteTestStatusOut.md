# RouteTestStatusOut

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**TestId** | **string** |  | 
**Status** | **string** |  | 
**Progress** | Pointer to [**NullableRouteTestProgress**](RouteTestProgress.md) |  | [optional] 
**Results** | Pointer to **map[string]interface{}** |  | [optional] 
**ErrorMessage** | Pointer to **NullableString** |  | [optional] 
**ErrorCode** | Pointer to **NullableString** |  | [optional] 

## Methods

### NewRouteTestStatusOut

`func NewRouteTestStatusOut(testId string, status string, ) *RouteTestStatusOut`

NewRouteTestStatusOut instantiates a new RouteTestStatusOut object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRouteTestStatusOutWithDefaults

`func NewRouteTestStatusOutWithDefaults() *RouteTestStatusOut`

NewRouteTestStatusOutWithDefaults instantiates a new RouteTestStatusOut object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTestId

`func (o *RouteTestStatusOut) GetTestId() string`

GetTestId returns the TestId field if non-nil, zero value otherwise.

### GetTestIdOk

`func (o *RouteTestStatusOut) GetTestIdOk() (*string, bool)`

GetTestIdOk returns a tuple with the TestId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTestId

`func (o *RouteTestStatusOut) SetTestId(v string)`

SetTestId sets TestId field to given value.


### GetStatus

`func (o *RouteTestStatusOut) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *RouteTestStatusOut) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *RouteTestStatusOut) SetStatus(v string)`

SetStatus sets Status field to given value.


### GetProgress

`func (o *RouteTestStatusOut) GetProgress() RouteTestProgress`

GetProgress returns the Progress field if non-nil, zero value otherwise.

### GetProgressOk

`func (o *RouteTestStatusOut) GetProgressOk() (*RouteTestProgress, bool)`

GetProgressOk returns a tuple with the Progress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProgress

`func (o *RouteTestStatusOut) SetProgress(v RouteTestProgress)`

SetProgress sets Progress field to given value.

### HasProgress

`func (o *RouteTestStatusOut) HasProgress() bool`

HasProgress returns a boolean if a field has been set.

### SetProgressNil

`func (o *RouteTestStatusOut) SetProgressNil(b bool)`

 SetProgressNil sets the value for Progress to be an explicit nil

### UnsetProgress
`func (o *RouteTestStatusOut) UnsetProgress()`

UnsetProgress ensures that no value is present for Progress, not even an explicit nil
### GetResults

`func (o *RouteTestStatusOut) GetResults() map[string]interface{}`

GetResults returns the Results field if non-nil, zero value otherwise.

### GetResultsOk

`func (o *RouteTestStatusOut) GetResultsOk() (*map[string]interface{}, bool)`

GetResultsOk returns a tuple with the Results field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResults

`func (o *RouteTestStatusOut) SetResults(v map[string]interface{})`

SetResults sets Results field to given value.

### HasResults

`func (o *RouteTestStatusOut) HasResults() bool`

HasResults returns a boolean if a field has been set.

### SetResultsNil

`func (o *RouteTestStatusOut) SetResultsNil(b bool)`

 SetResultsNil sets the value for Results to be an explicit nil

### UnsetResults
`func (o *RouteTestStatusOut) UnsetResults()`

UnsetResults ensures that no value is present for Results, not even an explicit nil
### GetErrorMessage

`func (o *RouteTestStatusOut) GetErrorMessage() string`

GetErrorMessage returns the ErrorMessage field if non-nil, zero value otherwise.

### GetErrorMessageOk

`func (o *RouteTestStatusOut) GetErrorMessageOk() (*string, bool)`

GetErrorMessageOk returns a tuple with the ErrorMessage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrorMessage

`func (o *RouteTestStatusOut) SetErrorMessage(v string)`

SetErrorMessage sets ErrorMessage field to given value.

### HasErrorMessage

`func (o *RouteTestStatusOut) HasErrorMessage() bool`

HasErrorMessage returns a boolean if a field has been set.

### SetErrorMessageNil

`func (o *RouteTestStatusOut) SetErrorMessageNil(b bool)`

 SetErrorMessageNil sets the value for ErrorMessage to be an explicit nil

### UnsetErrorMessage
`func (o *RouteTestStatusOut) UnsetErrorMessage()`

UnsetErrorMessage ensures that no value is present for ErrorMessage, not even an explicit nil
### GetErrorCode

`func (o *RouteTestStatusOut) GetErrorCode() string`

GetErrorCode returns the ErrorCode field if non-nil, zero value otherwise.

### GetErrorCodeOk

`func (o *RouteTestStatusOut) GetErrorCodeOk() (*string, bool)`

GetErrorCodeOk returns a tuple with the ErrorCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrorCode

`func (o *RouteTestStatusOut) SetErrorCode(v string)`

SetErrorCode sets ErrorCode field to given value.

### HasErrorCode

`func (o *RouteTestStatusOut) HasErrorCode() bool`

HasErrorCode returns a boolean if a field has been set.

### SetErrorCodeNil

`func (o *RouteTestStatusOut) SetErrorCodeNil(b bool)`

 SetErrorCodeNil sets the value for ErrorCode to be an explicit nil

### UnsetErrorCode
`func (o *RouteTestStatusOut) UnsetErrorCode()`

UnsetErrorCode ensures that no value is present for ErrorCode, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


