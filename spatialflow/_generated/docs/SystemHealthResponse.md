# SystemHealthResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Status** | **string** |  | 
**CheckedAt** | **string** |  | 
**Environment** | **string** |  | 
**Version** | **string** |  | 
**Components** | [**map[string]SystemHealthComponent**](SystemHealthComponent.md) |  | 

## Methods

### NewSystemHealthResponse

`func NewSystemHealthResponse(status string, checkedAt string, environment string, version string, components map[string]SystemHealthComponent, ) *SystemHealthResponse`

NewSystemHealthResponse instantiates a new SystemHealthResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSystemHealthResponseWithDefaults

`func NewSystemHealthResponseWithDefaults() *SystemHealthResponse`

NewSystemHealthResponseWithDefaults instantiates a new SystemHealthResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetStatus

`func (o *SystemHealthResponse) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *SystemHealthResponse) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *SystemHealthResponse) SetStatus(v string)`

SetStatus sets Status field to given value.


### GetCheckedAt

`func (o *SystemHealthResponse) GetCheckedAt() string`

GetCheckedAt returns the CheckedAt field if non-nil, zero value otherwise.

### GetCheckedAtOk

`func (o *SystemHealthResponse) GetCheckedAtOk() (*string, bool)`

GetCheckedAtOk returns a tuple with the CheckedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCheckedAt

`func (o *SystemHealthResponse) SetCheckedAt(v string)`

SetCheckedAt sets CheckedAt field to given value.


### GetEnvironment

`func (o *SystemHealthResponse) GetEnvironment() string`

GetEnvironment returns the Environment field if non-nil, zero value otherwise.

### GetEnvironmentOk

`func (o *SystemHealthResponse) GetEnvironmentOk() (*string, bool)`

GetEnvironmentOk returns a tuple with the Environment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnvironment

`func (o *SystemHealthResponse) SetEnvironment(v string)`

SetEnvironment sets Environment field to given value.


### GetVersion

`func (o *SystemHealthResponse) GetVersion() string`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *SystemHealthResponse) GetVersionOk() (*string, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *SystemHealthResponse) SetVersion(v string)`

SetVersion sets Version field to given value.


### GetComponents

`func (o *SystemHealthResponse) GetComponents() map[string]SystemHealthComponent`

GetComponents returns the Components field if non-nil, zero value otherwise.

### GetComponentsOk

`func (o *SystemHealthResponse) GetComponentsOk() (*map[string]SystemHealthComponent, bool)`

GetComponentsOk returns a tuple with the Components field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComponents

`func (o *SystemHealthResponse) SetComponents(v map[string]SystemHealthComponent)`

SetComponents sets Components field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


