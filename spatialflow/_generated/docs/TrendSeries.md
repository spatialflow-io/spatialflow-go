# TrendSeries

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** |  | 
**Data** | [**[]TrendDataPoint**](TrendDataPoint.md) |  | 

## Methods

### NewTrendSeries

`func NewTrendSeries(name string, data []TrendDataPoint, ) *TrendSeries`

NewTrendSeries instantiates a new TrendSeries object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTrendSeriesWithDefaults

`func NewTrendSeriesWithDefaults() *TrendSeries`

NewTrendSeriesWithDefaults instantiates a new TrendSeries object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *TrendSeries) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *TrendSeries) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *TrendSeries) SetName(v string)`

SetName sets Name field to given value.


### GetData

`func (o *TrendSeries) GetData() []TrendDataPoint`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *TrendSeries) GetDataOk() (*[]TrendDataPoint, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *TrendSeries) SetData(v []TrendDataPoint)`

SetData sets Data field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


