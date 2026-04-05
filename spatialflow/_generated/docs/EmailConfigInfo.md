# EmailConfigInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Backend** | **string** |  | 
**Host** | Pointer to **NullableString** |  | [optional] 
**Port** | Pointer to **NullableInt32** |  | [optional] 

## Methods

### NewEmailConfigInfo

`func NewEmailConfigInfo(backend string, ) *EmailConfigInfo`

NewEmailConfigInfo instantiates a new EmailConfigInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEmailConfigInfoWithDefaults

`func NewEmailConfigInfoWithDefaults() *EmailConfigInfo`

NewEmailConfigInfoWithDefaults instantiates a new EmailConfigInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBackend

`func (o *EmailConfigInfo) GetBackend() string`

GetBackend returns the Backend field if non-nil, zero value otherwise.

### GetBackendOk

`func (o *EmailConfigInfo) GetBackendOk() (*string, bool)`

GetBackendOk returns a tuple with the Backend field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBackend

`func (o *EmailConfigInfo) SetBackend(v string)`

SetBackend sets Backend field to given value.


### GetHost

`func (o *EmailConfigInfo) GetHost() string`

GetHost returns the Host field if non-nil, zero value otherwise.

### GetHostOk

`func (o *EmailConfigInfo) GetHostOk() (*string, bool)`

GetHostOk returns a tuple with the Host field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHost

`func (o *EmailConfigInfo) SetHost(v string)`

SetHost sets Host field to given value.

### HasHost

`func (o *EmailConfigInfo) HasHost() bool`

HasHost returns a boolean if a field has been set.

### SetHostNil

`func (o *EmailConfigInfo) SetHostNil(b bool)`

 SetHostNil sets the value for Host to be an explicit nil

### UnsetHost
`func (o *EmailConfigInfo) UnsetHost()`

UnsetHost ensures that no value is present for Host, not even an explicit nil
### GetPort

`func (o *EmailConfigInfo) GetPort() int32`

GetPort returns the Port field if non-nil, zero value otherwise.

### GetPortOk

`func (o *EmailConfigInfo) GetPortOk() (*int32, bool)`

GetPortOk returns a tuple with the Port field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPort

`func (o *EmailConfigInfo) SetPort(v int32)`

SetPort sets Port field to given value.

### HasPort

`func (o *EmailConfigInfo) HasPort() bool`

HasPort returns a boolean if a field has been set.

### SetPortNil

`func (o *EmailConfigInfo) SetPortNil(b bool)`

 SetPortNil sets the value for Port to be an explicit nil

### UnsetPort
`func (o *EmailConfigInfo) UnsetPort()`

UnsetPort ensures that no value is present for Port, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


