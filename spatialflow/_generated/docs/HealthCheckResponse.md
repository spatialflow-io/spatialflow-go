# HealthCheckResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Status** | **string** |  | 
**Timestamp** | **string** |  | 
**Services** | **map[string]string** |  | 

## Methods

### NewHealthCheckResponse

`func NewHealthCheckResponse(status string, timestamp string, services map[string]string, ) *HealthCheckResponse`

NewHealthCheckResponse instantiates a new HealthCheckResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewHealthCheckResponseWithDefaults

`func NewHealthCheckResponseWithDefaults() *HealthCheckResponse`

NewHealthCheckResponseWithDefaults instantiates a new HealthCheckResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetStatus

`func (o *HealthCheckResponse) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *HealthCheckResponse) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *HealthCheckResponse) SetStatus(v string)`

SetStatus sets Status field to given value.


### GetTimestamp

`func (o *HealthCheckResponse) GetTimestamp() string`

GetTimestamp returns the Timestamp field if non-nil, zero value otherwise.

### GetTimestampOk

`func (o *HealthCheckResponse) GetTimestampOk() (*string, bool)`

GetTimestampOk returns a tuple with the Timestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimestamp

`func (o *HealthCheckResponse) SetTimestamp(v string)`

SetTimestamp sets Timestamp field to given value.


### GetServices

`func (o *HealthCheckResponse) GetServices() map[string]string`

GetServices returns the Services field if non-nil, zero value otherwise.

### GetServicesOk

`func (o *HealthCheckResponse) GetServicesOk() (*map[string]string, bool)`

GetServicesOk returns a tuple with the Services field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServices

`func (o *HealthCheckResponse) SetServices(v map[string]string)`

SetServices sets Services field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


