# TrendDataPoint

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Date** | **string** |  | 
**Value** | **float32** |  | 

## Methods

### NewTrendDataPoint

`func NewTrendDataPoint(date string, value float32, ) *TrendDataPoint`

NewTrendDataPoint instantiates a new TrendDataPoint object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTrendDataPointWithDefaults

`func NewTrendDataPointWithDefaults() *TrendDataPoint`

NewTrendDataPointWithDefaults instantiates a new TrendDataPoint object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDate

`func (o *TrendDataPoint) GetDate() string`

GetDate returns the Date field if non-nil, zero value otherwise.

### GetDateOk

`func (o *TrendDataPoint) GetDateOk() (*string, bool)`

GetDateOk returns a tuple with the Date field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDate

`func (o *TrendDataPoint) SetDate(v string)`

SetDate sets Date field to given value.


### GetValue

`func (o *TrendDataPoint) GetValue() float32`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *TrendDataPoint) GetValueOk() (*float32, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *TrendDataPoint) SetValue(v float32)`

SetValue sets Value field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


