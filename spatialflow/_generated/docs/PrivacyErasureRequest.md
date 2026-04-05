# PrivacyErasureRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Scope** | **string** | Erasure scope: org, device, date_range, tag | 
**FromDate** | Pointer to **NullableTime** |  | [optional] 
**ToDate** | Pointer to **NullableTime** |  | [optional] 
**DeviceIds** | Pointer to **[]string** |  | [optional] 
**Tags** | Pointer to **[]string** |  | [optional] 
**DryRun** | Pointer to **bool** | If true, estimate deletions without actually deleting | [optional] [default to false]

## Methods

### NewPrivacyErasureRequest

`func NewPrivacyErasureRequest(scope string, ) *PrivacyErasureRequest`

NewPrivacyErasureRequest instantiates a new PrivacyErasureRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPrivacyErasureRequestWithDefaults

`func NewPrivacyErasureRequestWithDefaults() *PrivacyErasureRequest`

NewPrivacyErasureRequestWithDefaults instantiates a new PrivacyErasureRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetScope

`func (o *PrivacyErasureRequest) GetScope() string`

GetScope returns the Scope field if non-nil, zero value otherwise.

### GetScopeOk

`func (o *PrivacyErasureRequest) GetScopeOk() (*string, bool)`

GetScopeOk returns a tuple with the Scope field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScope

`func (o *PrivacyErasureRequest) SetScope(v string)`

SetScope sets Scope field to given value.


### GetFromDate

`func (o *PrivacyErasureRequest) GetFromDate() time.Time`

GetFromDate returns the FromDate field if non-nil, zero value otherwise.

### GetFromDateOk

`func (o *PrivacyErasureRequest) GetFromDateOk() (*time.Time, bool)`

GetFromDateOk returns a tuple with the FromDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFromDate

`func (o *PrivacyErasureRequest) SetFromDate(v time.Time)`

SetFromDate sets FromDate field to given value.

### HasFromDate

`func (o *PrivacyErasureRequest) HasFromDate() bool`

HasFromDate returns a boolean if a field has been set.

### SetFromDateNil

`func (o *PrivacyErasureRequest) SetFromDateNil(b bool)`

 SetFromDateNil sets the value for FromDate to be an explicit nil

### UnsetFromDate
`func (o *PrivacyErasureRequest) UnsetFromDate()`

UnsetFromDate ensures that no value is present for FromDate, not even an explicit nil
### GetToDate

`func (o *PrivacyErasureRequest) GetToDate() time.Time`

GetToDate returns the ToDate field if non-nil, zero value otherwise.

### GetToDateOk

`func (o *PrivacyErasureRequest) GetToDateOk() (*time.Time, bool)`

GetToDateOk returns a tuple with the ToDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToDate

`func (o *PrivacyErasureRequest) SetToDate(v time.Time)`

SetToDate sets ToDate field to given value.

### HasToDate

`func (o *PrivacyErasureRequest) HasToDate() bool`

HasToDate returns a boolean if a field has been set.

### SetToDateNil

`func (o *PrivacyErasureRequest) SetToDateNil(b bool)`

 SetToDateNil sets the value for ToDate to be an explicit nil

### UnsetToDate
`func (o *PrivacyErasureRequest) UnsetToDate()`

UnsetToDate ensures that no value is present for ToDate, not even an explicit nil
### GetDeviceIds

`func (o *PrivacyErasureRequest) GetDeviceIds() []string`

GetDeviceIds returns the DeviceIds field if non-nil, zero value otherwise.

### GetDeviceIdsOk

`func (o *PrivacyErasureRequest) GetDeviceIdsOk() (*[]string, bool)`

GetDeviceIdsOk returns a tuple with the DeviceIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeviceIds

`func (o *PrivacyErasureRequest) SetDeviceIds(v []string)`

SetDeviceIds sets DeviceIds field to given value.

### HasDeviceIds

`func (o *PrivacyErasureRequest) HasDeviceIds() bool`

HasDeviceIds returns a boolean if a field has been set.

### SetDeviceIdsNil

`func (o *PrivacyErasureRequest) SetDeviceIdsNil(b bool)`

 SetDeviceIdsNil sets the value for DeviceIds to be an explicit nil

### UnsetDeviceIds
`func (o *PrivacyErasureRequest) UnsetDeviceIds()`

UnsetDeviceIds ensures that no value is present for DeviceIds, not even an explicit nil
### GetTags

`func (o *PrivacyErasureRequest) GetTags() []string`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *PrivacyErasureRequest) GetTagsOk() (*[]string, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *PrivacyErasureRequest) SetTags(v []string)`

SetTags sets Tags field to given value.

### HasTags

`func (o *PrivacyErasureRequest) HasTags() bool`

HasTags returns a boolean if a field has been set.

### SetTagsNil

`func (o *PrivacyErasureRequest) SetTagsNil(b bool)`

 SetTagsNil sets the value for Tags to be an explicit nil

### UnsetTags
`func (o *PrivacyErasureRequest) UnsetTags()`

UnsetTags ensures that no value is present for Tags, not even an explicit nil
### GetDryRun

`func (o *PrivacyErasureRequest) GetDryRun() bool`

GetDryRun returns the DryRun field if non-nil, zero value otherwise.

### GetDryRunOk

`func (o *PrivacyErasureRequest) GetDryRunOk() (*bool, bool)`

GetDryRunOk returns a tuple with the DryRun field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDryRun

`func (o *PrivacyErasureRequest) SetDryRun(v bool)`

SetDryRun sets DryRun field to given value.

### HasDryRun

`func (o *PrivacyErasureRequest) HasDryRun() bool`

HasDryRun returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


