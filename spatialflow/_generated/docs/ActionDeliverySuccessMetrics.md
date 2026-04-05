# ActionDeliverySuccessMetrics

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Rate** | Pointer to **NullableFloat32** |  | [optional] 
**Successful** | **int32** | Number of successful action deliveries | 
**Total** | **int32** | Total number of action attempts | 

## Methods

### NewActionDeliverySuccessMetrics

`func NewActionDeliverySuccessMetrics(successful int32, total int32, ) *ActionDeliverySuccessMetrics`

NewActionDeliverySuccessMetrics instantiates a new ActionDeliverySuccessMetrics object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewActionDeliverySuccessMetricsWithDefaults

`func NewActionDeliverySuccessMetricsWithDefaults() *ActionDeliverySuccessMetrics`

NewActionDeliverySuccessMetricsWithDefaults instantiates a new ActionDeliverySuccessMetrics object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRate

`func (o *ActionDeliverySuccessMetrics) GetRate() float32`

GetRate returns the Rate field if non-nil, zero value otherwise.

### GetRateOk

`func (o *ActionDeliverySuccessMetrics) GetRateOk() (*float32, bool)`

GetRateOk returns a tuple with the Rate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRate

`func (o *ActionDeliverySuccessMetrics) SetRate(v float32)`

SetRate sets Rate field to given value.

### HasRate

`func (o *ActionDeliverySuccessMetrics) HasRate() bool`

HasRate returns a boolean if a field has been set.

### SetRateNil

`func (o *ActionDeliverySuccessMetrics) SetRateNil(b bool)`

 SetRateNil sets the value for Rate to be an explicit nil

### UnsetRate
`func (o *ActionDeliverySuccessMetrics) UnsetRate()`

UnsetRate ensures that no value is present for Rate, not even an explicit nil
### GetSuccessful

`func (o *ActionDeliverySuccessMetrics) GetSuccessful() int32`

GetSuccessful returns the Successful field if non-nil, zero value otherwise.

### GetSuccessfulOk

`func (o *ActionDeliverySuccessMetrics) GetSuccessfulOk() (*int32, bool)`

GetSuccessfulOk returns a tuple with the Successful field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSuccessful

`func (o *ActionDeliverySuccessMetrics) SetSuccessful(v int32)`

SetSuccessful sets Successful field to given value.


### GetTotal

`func (o *ActionDeliverySuccessMetrics) GetTotal() int32`

GetTotal returns the Total field if non-nil, zero value otherwise.

### GetTotalOk

`func (o *ActionDeliverySuccessMetrics) GetTotalOk() (*int32, bool)`

GetTotalOk returns a tuple with the Total field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotal

`func (o *ActionDeliverySuccessMetrics) SetTotal(v int32)`

SetTotal sets Total field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


