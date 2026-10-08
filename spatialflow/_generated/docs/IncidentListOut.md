# IncidentListOut

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Results** | [**[]IncidentOut**](IncidentOut.md) |  | 
**Total** | **int32** |  | 
**NextCursor** | Pointer to **NullableString** |  | [optional] 

## Methods

### NewIncidentListOut

`func NewIncidentListOut(results []IncidentOut, total int32, ) *IncidentListOut`

NewIncidentListOut instantiates a new IncidentListOut object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIncidentListOutWithDefaults

`func NewIncidentListOutWithDefaults() *IncidentListOut`

NewIncidentListOutWithDefaults instantiates a new IncidentListOut object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetResults

`func (o *IncidentListOut) GetResults() []IncidentOut`

GetResults returns the Results field if non-nil, zero value otherwise.

### GetResultsOk

`func (o *IncidentListOut) GetResultsOk() (*[]IncidentOut, bool)`

GetResultsOk returns a tuple with the Results field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResults

`func (o *IncidentListOut) SetResults(v []IncidentOut)`

SetResults sets Results field to given value.


### GetTotal

`func (o *IncidentListOut) GetTotal() int32`

GetTotal returns the Total field if non-nil, zero value otherwise.

### GetTotalOk

`func (o *IncidentListOut) GetTotalOk() (*int32, bool)`

GetTotalOk returns a tuple with the Total field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotal

`func (o *IncidentListOut) SetTotal(v int32)`

SetTotal sets Total field to given value.


### GetNextCursor

`func (o *IncidentListOut) GetNextCursor() string`

GetNextCursor returns the NextCursor field if non-nil, zero value otherwise.

### GetNextCursorOk

`func (o *IncidentListOut) GetNextCursorOk() (*string, bool)`

GetNextCursorOk returns a tuple with the NextCursor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNextCursor

`func (o *IncidentListOut) SetNextCursor(v string)`

SetNextCursor sets NextCursor field to given value.

### HasNextCursor

`func (o *IncidentListOut) HasNextCursor() bool`

HasNextCursor returns a boolean if a field has been set.

### SetNextCursorNil

`func (o *IncidentListOut) SetNextCursorNil(b bool)`

 SetNextCursorNil sets the value for NextCursor to be an explicit nil

### UnsetNextCursor
`func (o *IncidentListOut) UnsetNextCursor()`

UnsetNextCursor ensures that no value is present for NextCursor, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


