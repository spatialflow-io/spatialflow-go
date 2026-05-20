# HealthOut

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Status** | **string** |  | 
**Service** | **string** |  | 
**Timestamp** | **string** |  | 
**Checks** | [**HealthChecksOut**](HealthChecksOut.md) |  | 

## Methods

### NewHealthOut

`func NewHealthOut(status string, service string, timestamp string, checks HealthChecksOut, ) *HealthOut`

NewHealthOut instantiates a new HealthOut object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewHealthOutWithDefaults

`func NewHealthOutWithDefaults() *HealthOut`

NewHealthOutWithDefaults instantiates a new HealthOut object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetStatus

`func (o *HealthOut) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *HealthOut) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *HealthOut) SetStatus(v string)`

SetStatus sets Status field to given value.


### GetService

`func (o *HealthOut) GetService() string`

GetService returns the Service field if non-nil, zero value otherwise.

### GetServiceOk

`func (o *HealthOut) GetServiceOk() (*string, bool)`

GetServiceOk returns a tuple with the Service field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetService

`func (o *HealthOut) SetService(v string)`

SetService sets Service field to given value.


### GetTimestamp

`func (o *HealthOut) GetTimestamp() string`

GetTimestamp returns the Timestamp field if non-nil, zero value otherwise.

### GetTimestampOk

`func (o *HealthOut) GetTimestampOk() (*string, bool)`

GetTimestampOk returns a tuple with the Timestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimestamp

`func (o *HealthOut) SetTimestamp(v string)`

SetTimestamp sets Timestamp field to given value.


### GetChecks

`func (o *HealthOut) GetChecks() HealthChecksOut`

GetChecks returns the Checks field if non-nil, zero value otherwise.

### GetChecksOk

`func (o *HealthOut) GetChecksOk() (*HealthChecksOut, bool)`

GetChecksOk returns a tuple with the Checks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChecks

`func (o *HealthOut) SetChecks(v HealthChecksOut)`

SetChecks sets Checks field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


