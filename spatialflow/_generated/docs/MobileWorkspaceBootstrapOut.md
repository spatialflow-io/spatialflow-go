# MobileWorkspaceBootstrapOut

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**EligibleWorkspaces** | [**[]MobileWorkspaceRow**](MobileWorkspaceRow.md) |  | 
**EligibleCount** | **int32** |  | 
**SelectedWorkspaceId** | Pointer to **NullableString** |  | [optional] 
**SelectedWorkspace** | Pointer to [**NullableMobileWorkspaceRow**](MobileWorkspaceRow.md) |  | [optional] 
**SkipPicker** | **bool** |  | 
**SelectionRequired** | **bool** |  | 

## Methods

### NewMobileWorkspaceBootstrapOut

`func NewMobileWorkspaceBootstrapOut(eligibleWorkspaces []MobileWorkspaceRow, eligibleCount int32, skipPicker bool, selectionRequired bool, ) *MobileWorkspaceBootstrapOut`

NewMobileWorkspaceBootstrapOut instantiates a new MobileWorkspaceBootstrapOut object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMobileWorkspaceBootstrapOutWithDefaults

`func NewMobileWorkspaceBootstrapOutWithDefaults() *MobileWorkspaceBootstrapOut`

NewMobileWorkspaceBootstrapOutWithDefaults instantiates a new MobileWorkspaceBootstrapOut object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEligibleWorkspaces

`func (o *MobileWorkspaceBootstrapOut) GetEligibleWorkspaces() []MobileWorkspaceRow`

GetEligibleWorkspaces returns the EligibleWorkspaces field if non-nil, zero value otherwise.

### GetEligibleWorkspacesOk

`func (o *MobileWorkspaceBootstrapOut) GetEligibleWorkspacesOk() (*[]MobileWorkspaceRow, bool)`

GetEligibleWorkspacesOk returns a tuple with the EligibleWorkspaces field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEligibleWorkspaces

`func (o *MobileWorkspaceBootstrapOut) SetEligibleWorkspaces(v []MobileWorkspaceRow)`

SetEligibleWorkspaces sets EligibleWorkspaces field to given value.


### GetEligibleCount

`func (o *MobileWorkspaceBootstrapOut) GetEligibleCount() int32`

GetEligibleCount returns the EligibleCount field if non-nil, zero value otherwise.

### GetEligibleCountOk

`func (o *MobileWorkspaceBootstrapOut) GetEligibleCountOk() (*int32, bool)`

GetEligibleCountOk returns a tuple with the EligibleCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEligibleCount

`func (o *MobileWorkspaceBootstrapOut) SetEligibleCount(v int32)`

SetEligibleCount sets EligibleCount field to given value.


### GetSelectedWorkspaceId

`func (o *MobileWorkspaceBootstrapOut) GetSelectedWorkspaceId() string`

GetSelectedWorkspaceId returns the SelectedWorkspaceId field if non-nil, zero value otherwise.

### GetSelectedWorkspaceIdOk

`func (o *MobileWorkspaceBootstrapOut) GetSelectedWorkspaceIdOk() (*string, bool)`

GetSelectedWorkspaceIdOk returns a tuple with the SelectedWorkspaceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSelectedWorkspaceId

`func (o *MobileWorkspaceBootstrapOut) SetSelectedWorkspaceId(v string)`

SetSelectedWorkspaceId sets SelectedWorkspaceId field to given value.

### HasSelectedWorkspaceId

`func (o *MobileWorkspaceBootstrapOut) HasSelectedWorkspaceId() bool`

HasSelectedWorkspaceId returns a boolean if a field has been set.

### SetSelectedWorkspaceIdNil

`func (o *MobileWorkspaceBootstrapOut) SetSelectedWorkspaceIdNil(b bool)`

 SetSelectedWorkspaceIdNil sets the value for SelectedWorkspaceId to be an explicit nil

### UnsetSelectedWorkspaceId
`func (o *MobileWorkspaceBootstrapOut) UnsetSelectedWorkspaceId()`

UnsetSelectedWorkspaceId ensures that no value is present for SelectedWorkspaceId, not even an explicit nil
### GetSelectedWorkspace

`func (o *MobileWorkspaceBootstrapOut) GetSelectedWorkspace() MobileWorkspaceRow`

GetSelectedWorkspace returns the SelectedWorkspace field if non-nil, zero value otherwise.

### GetSelectedWorkspaceOk

`func (o *MobileWorkspaceBootstrapOut) GetSelectedWorkspaceOk() (*MobileWorkspaceRow, bool)`

GetSelectedWorkspaceOk returns a tuple with the SelectedWorkspace field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSelectedWorkspace

`func (o *MobileWorkspaceBootstrapOut) SetSelectedWorkspace(v MobileWorkspaceRow)`

SetSelectedWorkspace sets SelectedWorkspace field to given value.

### HasSelectedWorkspace

`func (o *MobileWorkspaceBootstrapOut) HasSelectedWorkspace() bool`

HasSelectedWorkspace returns a boolean if a field has been set.

### SetSelectedWorkspaceNil

`func (o *MobileWorkspaceBootstrapOut) SetSelectedWorkspaceNil(b bool)`

 SetSelectedWorkspaceNil sets the value for SelectedWorkspace to be an explicit nil

### UnsetSelectedWorkspace
`func (o *MobileWorkspaceBootstrapOut) UnsetSelectedWorkspace()`

UnsetSelectedWorkspace ensures that no value is present for SelectedWorkspace, not even an explicit nil
### GetSkipPicker

`func (o *MobileWorkspaceBootstrapOut) GetSkipPicker() bool`

GetSkipPicker returns the SkipPicker field if non-nil, zero value otherwise.

### GetSkipPickerOk

`func (o *MobileWorkspaceBootstrapOut) GetSkipPickerOk() (*bool, bool)`

GetSkipPickerOk returns a tuple with the SkipPicker field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSkipPicker

`func (o *MobileWorkspaceBootstrapOut) SetSkipPicker(v bool)`

SetSkipPicker sets SkipPicker field to given value.


### GetSelectionRequired

`func (o *MobileWorkspaceBootstrapOut) GetSelectionRequired() bool`

GetSelectionRequired returns the SelectionRequired field if non-nil, zero value otherwise.

### GetSelectionRequiredOk

`func (o *MobileWorkspaceBootstrapOut) GetSelectionRequiredOk() (*bool, bool)`

GetSelectionRequiredOk returns a tuple with the SelectionRequired field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSelectionRequired

`func (o *MobileWorkspaceBootstrapOut) SetSelectionRequired(v bool)`

SetSelectionRequired sets SelectionRequired field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


