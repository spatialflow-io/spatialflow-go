# IncidentGroupListOut

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Results** | [**[]IncidentGroupOut**](IncidentGroupOut.md) |  | 
**AsOf** | **time.Time** |  | 

## Methods

### NewIncidentGroupListOut

`func NewIncidentGroupListOut(results []IncidentGroupOut, asOf time.Time, ) *IncidentGroupListOut`

NewIncidentGroupListOut instantiates a new IncidentGroupListOut object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIncidentGroupListOutWithDefaults

`func NewIncidentGroupListOutWithDefaults() *IncidentGroupListOut`

NewIncidentGroupListOutWithDefaults instantiates a new IncidentGroupListOut object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetResults

`func (o *IncidentGroupListOut) GetResults() []IncidentGroupOut`

GetResults returns the Results field if non-nil, zero value otherwise.

### GetResultsOk

`func (o *IncidentGroupListOut) GetResultsOk() (*[]IncidentGroupOut, bool)`

GetResultsOk returns a tuple with the Results field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResults

`func (o *IncidentGroupListOut) SetResults(v []IncidentGroupOut)`

SetResults sets Results field to given value.


### GetAsOf

`func (o *IncidentGroupListOut) GetAsOf() time.Time`

GetAsOf returns the AsOf field if non-nil, zero value otherwise.

### GetAsOfOk

`func (o *IncidentGroupListOut) GetAsOfOk() (*time.Time, bool)`

GetAsOfOk returns a tuple with the AsOf field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAsOf

`func (o *IncidentGroupListOut) SetAsOf(v time.Time)`

SetAsOf sets AsOf field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


