# RuntimeConfigAnalyticsOut

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PosthogEnabled** | **bool** |  | 
**PosthogHost** | Pointer to **NullableString** |  | [optional] 

## Methods

### NewRuntimeConfigAnalyticsOut

`func NewRuntimeConfigAnalyticsOut(posthogEnabled bool, ) *RuntimeConfigAnalyticsOut`

NewRuntimeConfigAnalyticsOut instantiates a new RuntimeConfigAnalyticsOut object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRuntimeConfigAnalyticsOutWithDefaults

`func NewRuntimeConfigAnalyticsOutWithDefaults() *RuntimeConfigAnalyticsOut`

NewRuntimeConfigAnalyticsOutWithDefaults instantiates a new RuntimeConfigAnalyticsOut object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPosthogEnabled

`func (o *RuntimeConfigAnalyticsOut) GetPosthogEnabled() bool`

GetPosthogEnabled returns the PosthogEnabled field if non-nil, zero value otherwise.

### GetPosthogEnabledOk

`func (o *RuntimeConfigAnalyticsOut) GetPosthogEnabledOk() (*bool, bool)`

GetPosthogEnabledOk returns a tuple with the PosthogEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPosthogEnabled

`func (o *RuntimeConfigAnalyticsOut) SetPosthogEnabled(v bool)`

SetPosthogEnabled sets PosthogEnabled field to given value.


### GetPosthogHost

`func (o *RuntimeConfigAnalyticsOut) GetPosthogHost() string`

GetPosthogHost returns the PosthogHost field if non-nil, zero value otherwise.

### GetPosthogHostOk

`func (o *RuntimeConfigAnalyticsOut) GetPosthogHostOk() (*string, bool)`

GetPosthogHostOk returns a tuple with the PosthogHost field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPosthogHost

`func (o *RuntimeConfigAnalyticsOut) SetPosthogHost(v string)`

SetPosthogHost sets PosthogHost field to given value.

### HasPosthogHost

`func (o *RuntimeConfigAnalyticsOut) HasPosthogHost() bool`

HasPosthogHost returns a boolean if a field has been set.

### SetPosthogHostNil

`func (o *RuntimeConfigAnalyticsOut) SetPosthogHostNil(b bool)`

 SetPosthogHostNil sets the value for PosthogHost to be an explicit nil

### UnsetPosthogHost
`func (o *RuntimeConfigAnalyticsOut) UnsetPosthogHost()`

UnsetPosthogHost ensures that no value is present for PosthogHost, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


