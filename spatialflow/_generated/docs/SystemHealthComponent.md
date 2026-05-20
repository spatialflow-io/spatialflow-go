# SystemHealthComponent

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Label** | **string** |  | 
**Status** | **string** |  | 
**Detail** | **string** |  | 
**Value** | Pointer to **NullableString** |  | [optional] 
**Metadata** | Pointer to **map[string]interface{}** |  | [optional] 

## Methods

### NewSystemHealthComponent

`func NewSystemHealthComponent(label string, status string, detail string, ) *SystemHealthComponent`

NewSystemHealthComponent instantiates a new SystemHealthComponent object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSystemHealthComponentWithDefaults

`func NewSystemHealthComponentWithDefaults() *SystemHealthComponent`

NewSystemHealthComponentWithDefaults instantiates a new SystemHealthComponent object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLabel

`func (o *SystemHealthComponent) GetLabel() string`

GetLabel returns the Label field if non-nil, zero value otherwise.

### GetLabelOk

`func (o *SystemHealthComponent) GetLabelOk() (*string, bool)`

GetLabelOk returns a tuple with the Label field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabel

`func (o *SystemHealthComponent) SetLabel(v string)`

SetLabel sets Label field to given value.


### GetStatus

`func (o *SystemHealthComponent) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *SystemHealthComponent) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *SystemHealthComponent) SetStatus(v string)`

SetStatus sets Status field to given value.


### GetDetail

`func (o *SystemHealthComponent) GetDetail() string`

GetDetail returns the Detail field if non-nil, zero value otherwise.

### GetDetailOk

`func (o *SystemHealthComponent) GetDetailOk() (*string, bool)`

GetDetailOk returns a tuple with the Detail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDetail

`func (o *SystemHealthComponent) SetDetail(v string)`

SetDetail sets Detail field to given value.


### GetValue

`func (o *SystemHealthComponent) GetValue() string`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *SystemHealthComponent) GetValueOk() (*string, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *SystemHealthComponent) SetValue(v string)`

SetValue sets Value field to given value.

### HasValue

`func (o *SystemHealthComponent) HasValue() bool`

HasValue returns a boolean if a field has been set.

### SetValueNil

`func (o *SystemHealthComponent) SetValueNil(b bool)`

 SetValueNil sets the value for Value to be an explicit nil

### UnsetValue
`func (o *SystemHealthComponent) UnsetValue()`

UnsetValue ensures that no value is present for Value, not even an explicit nil
### GetMetadata

`func (o *SystemHealthComponent) GetMetadata() map[string]interface{}`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *SystemHealthComponent) GetMetadataOk() (*map[string]interface{}, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *SystemHealthComponent) SetMetadata(v map[string]interface{})`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *SystemHealthComponent) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### SetMetadataNil

`func (o *SystemHealthComponent) SetMetadataNil(b bool)`

 SetMetadataNil sets the value for Metadata to be an explicit nil

### UnsetMetadata
`func (o *SystemHealthComponent) UnsetMetadata()`

UnsetMetadata ensures that no value is present for Metadata, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


