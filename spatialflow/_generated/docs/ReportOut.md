# ReportOut

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Rows** | **[]map[string]interface{}** |  | 
**Summary** | **map[string]interface{}** |  | 
**TotalCount** | **int32** |  | 

## Methods

### NewReportOut

`func NewReportOut(rows []map[string]interface{}, summary map[string]interface{}, totalCount int32, ) *ReportOut`

NewReportOut instantiates a new ReportOut object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewReportOutWithDefaults

`func NewReportOutWithDefaults() *ReportOut`

NewReportOutWithDefaults instantiates a new ReportOut object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRows

`func (o *ReportOut) GetRows() []map[string]interface{}`

GetRows returns the Rows field if non-nil, zero value otherwise.

### GetRowsOk

`func (o *ReportOut) GetRowsOk() (*[]map[string]interface{}, bool)`

GetRowsOk returns a tuple with the Rows field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRows

`func (o *ReportOut) SetRows(v []map[string]interface{})`

SetRows sets Rows field to given value.


### GetSummary

`func (o *ReportOut) GetSummary() map[string]interface{}`

GetSummary returns the Summary field if non-nil, zero value otherwise.

### GetSummaryOk

`func (o *ReportOut) GetSummaryOk() (*map[string]interface{}, bool)`

GetSummaryOk returns a tuple with the Summary field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSummary

`func (o *ReportOut) SetSummary(v map[string]interface{})`

SetSummary sets Summary field to given value.


### GetTotalCount

`func (o *ReportOut) GetTotalCount() int32`

GetTotalCount returns the TotalCount field if non-nil, zero value otherwise.

### GetTotalCountOk

`func (o *ReportOut) GetTotalCountOk() (*int32, bool)`

GetTotalCountOk returns a tuple with the TotalCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalCount

`func (o *ReportOut) SetTotalCount(v int32)`

SetTotalCount sets TotalCount field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


