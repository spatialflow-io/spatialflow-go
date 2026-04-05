# SessionLocationsOut

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SessionId** | **string** |  | 
**Locations** | [**[]LocationPointOut**](LocationPointOut.md) |  | 
**TotalCount** | **int32** |  | 
**Offset** | **int32** |  | 
**Limit** | **int32** |  | 
**Simplified** | Pointer to **bool** |  | [optional] [default to false]

## Methods

### NewSessionLocationsOut

`func NewSessionLocationsOut(sessionId string, locations []LocationPointOut, totalCount int32, offset int32, limit int32, ) *SessionLocationsOut`

NewSessionLocationsOut instantiates a new SessionLocationsOut object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSessionLocationsOutWithDefaults

`func NewSessionLocationsOutWithDefaults() *SessionLocationsOut`

NewSessionLocationsOutWithDefaults instantiates a new SessionLocationsOut object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSessionId

`func (o *SessionLocationsOut) GetSessionId() string`

GetSessionId returns the SessionId field if non-nil, zero value otherwise.

### GetSessionIdOk

`func (o *SessionLocationsOut) GetSessionIdOk() (*string, bool)`

GetSessionIdOk returns a tuple with the SessionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSessionId

`func (o *SessionLocationsOut) SetSessionId(v string)`

SetSessionId sets SessionId field to given value.


### GetLocations

`func (o *SessionLocationsOut) GetLocations() []LocationPointOut`

GetLocations returns the Locations field if non-nil, zero value otherwise.

### GetLocationsOk

`func (o *SessionLocationsOut) GetLocationsOk() (*[]LocationPointOut, bool)`

GetLocationsOk returns a tuple with the Locations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocations

`func (o *SessionLocationsOut) SetLocations(v []LocationPointOut)`

SetLocations sets Locations field to given value.


### GetTotalCount

`func (o *SessionLocationsOut) GetTotalCount() int32`

GetTotalCount returns the TotalCount field if non-nil, zero value otherwise.

### GetTotalCountOk

`func (o *SessionLocationsOut) GetTotalCountOk() (*int32, bool)`

GetTotalCountOk returns a tuple with the TotalCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalCount

`func (o *SessionLocationsOut) SetTotalCount(v int32)`

SetTotalCount sets TotalCount field to given value.


### GetOffset

`func (o *SessionLocationsOut) GetOffset() int32`

GetOffset returns the Offset field if non-nil, zero value otherwise.

### GetOffsetOk

`func (o *SessionLocationsOut) GetOffsetOk() (*int32, bool)`

GetOffsetOk returns a tuple with the Offset field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOffset

`func (o *SessionLocationsOut) SetOffset(v int32)`

SetOffset sets Offset field to given value.


### GetLimit

`func (o *SessionLocationsOut) GetLimit() int32`

GetLimit returns the Limit field if non-nil, zero value otherwise.

### GetLimitOk

`func (o *SessionLocationsOut) GetLimitOk() (*int32, bool)`

GetLimitOk returns a tuple with the Limit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLimit

`func (o *SessionLocationsOut) SetLimit(v int32)`

SetLimit sets Limit field to given value.


### GetSimplified

`func (o *SessionLocationsOut) GetSimplified() bool`

GetSimplified returns the Simplified field if non-nil, zero value otherwise.

### GetSimplifiedOk

`func (o *SessionLocationsOut) GetSimplifiedOk() (*bool, bool)`

GetSimplifiedOk returns a tuple with the Simplified field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSimplified

`func (o *SessionLocationsOut) SetSimplified(v bool)`

SetSimplified sets Simplified field to given value.

### HasSimplified

`func (o *SessionLocationsOut) HasSimplified() bool`

HasSimplified returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


