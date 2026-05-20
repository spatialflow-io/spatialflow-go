# MobileWorkspaceSelectionOut

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**EligibleWorkspaces** | [**[]MobileWorkspaceRow**](MobileWorkspaceRow.md) |  | 
**EligibleCount** | **int32** |  | 
**SelectedWorkspaceId** | Pointer to **NullableString** |  | [optional] 
**SelectedWorkspace** | Pointer to [**NullableMobileWorkspaceRow**](MobileWorkspaceRow.md) |  | [optional] 
**SkipPicker** | **bool** |  | 
**SelectionRequired** | **bool** |  | 
**AccessToken** | **string** |  | 
**RefreshToken** | **string** |  | 
**TokenType** | **string** |  | 
**ExpiresIn** | **int32** |  | 
**User** | **map[string]interface{}** |  | 

## Methods

### NewMobileWorkspaceSelectionOut

`func NewMobileWorkspaceSelectionOut(eligibleWorkspaces []MobileWorkspaceRow, eligibleCount int32, skipPicker bool, selectionRequired bool, accessToken string, refreshToken string, tokenType string, expiresIn int32, user map[string]interface{}, ) *MobileWorkspaceSelectionOut`

NewMobileWorkspaceSelectionOut instantiates a new MobileWorkspaceSelectionOut object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMobileWorkspaceSelectionOutWithDefaults

`func NewMobileWorkspaceSelectionOutWithDefaults() *MobileWorkspaceSelectionOut`

NewMobileWorkspaceSelectionOutWithDefaults instantiates a new MobileWorkspaceSelectionOut object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEligibleWorkspaces

`func (o *MobileWorkspaceSelectionOut) GetEligibleWorkspaces() []MobileWorkspaceRow`

GetEligibleWorkspaces returns the EligibleWorkspaces field if non-nil, zero value otherwise.

### GetEligibleWorkspacesOk

`func (o *MobileWorkspaceSelectionOut) GetEligibleWorkspacesOk() (*[]MobileWorkspaceRow, bool)`

GetEligibleWorkspacesOk returns a tuple with the EligibleWorkspaces field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEligibleWorkspaces

`func (o *MobileWorkspaceSelectionOut) SetEligibleWorkspaces(v []MobileWorkspaceRow)`

SetEligibleWorkspaces sets EligibleWorkspaces field to given value.


### GetEligibleCount

`func (o *MobileWorkspaceSelectionOut) GetEligibleCount() int32`

GetEligibleCount returns the EligibleCount field if non-nil, zero value otherwise.

### GetEligibleCountOk

`func (o *MobileWorkspaceSelectionOut) GetEligibleCountOk() (*int32, bool)`

GetEligibleCountOk returns a tuple with the EligibleCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEligibleCount

`func (o *MobileWorkspaceSelectionOut) SetEligibleCount(v int32)`

SetEligibleCount sets EligibleCount field to given value.


### GetSelectedWorkspaceId

`func (o *MobileWorkspaceSelectionOut) GetSelectedWorkspaceId() string`

GetSelectedWorkspaceId returns the SelectedWorkspaceId field if non-nil, zero value otherwise.

### GetSelectedWorkspaceIdOk

`func (o *MobileWorkspaceSelectionOut) GetSelectedWorkspaceIdOk() (*string, bool)`

GetSelectedWorkspaceIdOk returns a tuple with the SelectedWorkspaceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSelectedWorkspaceId

`func (o *MobileWorkspaceSelectionOut) SetSelectedWorkspaceId(v string)`

SetSelectedWorkspaceId sets SelectedWorkspaceId field to given value.

### HasSelectedWorkspaceId

`func (o *MobileWorkspaceSelectionOut) HasSelectedWorkspaceId() bool`

HasSelectedWorkspaceId returns a boolean if a field has been set.

### SetSelectedWorkspaceIdNil

`func (o *MobileWorkspaceSelectionOut) SetSelectedWorkspaceIdNil(b bool)`

 SetSelectedWorkspaceIdNil sets the value for SelectedWorkspaceId to be an explicit nil

### UnsetSelectedWorkspaceId
`func (o *MobileWorkspaceSelectionOut) UnsetSelectedWorkspaceId()`

UnsetSelectedWorkspaceId ensures that no value is present for SelectedWorkspaceId, not even an explicit nil
### GetSelectedWorkspace

`func (o *MobileWorkspaceSelectionOut) GetSelectedWorkspace() MobileWorkspaceRow`

GetSelectedWorkspace returns the SelectedWorkspace field if non-nil, zero value otherwise.

### GetSelectedWorkspaceOk

`func (o *MobileWorkspaceSelectionOut) GetSelectedWorkspaceOk() (*MobileWorkspaceRow, bool)`

GetSelectedWorkspaceOk returns a tuple with the SelectedWorkspace field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSelectedWorkspace

`func (o *MobileWorkspaceSelectionOut) SetSelectedWorkspace(v MobileWorkspaceRow)`

SetSelectedWorkspace sets SelectedWorkspace field to given value.

### HasSelectedWorkspace

`func (o *MobileWorkspaceSelectionOut) HasSelectedWorkspace() bool`

HasSelectedWorkspace returns a boolean if a field has been set.

### SetSelectedWorkspaceNil

`func (o *MobileWorkspaceSelectionOut) SetSelectedWorkspaceNil(b bool)`

 SetSelectedWorkspaceNil sets the value for SelectedWorkspace to be an explicit nil

### UnsetSelectedWorkspace
`func (o *MobileWorkspaceSelectionOut) UnsetSelectedWorkspace()`

UnsetSelectedWorkspace ensures that no value is present for SelectedWorkspace, not even an explicit nil
### GetSkipPicker

`func (o *MobileWorkspaceSelectionOut) GetSkipPicker() bool`

GetSkipPicker returns the SkipPicker field if non-nil, zero value otherwise.

### GetSkipPickerOk

`func (o *MobileWorkspaceSelectionOut) GetSkipPickerOk() (*bool, bool)`

GetSkipPickerOk returns a tuple with the SkipPicker field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSkipPicker

`func (o *MobileWorkspaceSelectionOut) SetSkipPicker(v bool)`

SetSkipPicker sets SkipPicker field to given value.


### GetSelectionRequired

`func (o *MobileWorkspaceSelectionOut) GetSelectionRequired() bool`

GetSelectionRequired returns the SelectionRequired field if non-nil, zero value otherwise.

### GetSelectionRequiredOk

`func (o *MobileWorkspaceSelectionOut) GetSelectionRequiredOk() (*bool, bool)`

GetSelectionRequiredOk returns a tuple with the SelectionRequired field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSelectionRequired

`func (o *MobileWorkspaceSelectionOut) SetSelectionRequired(v bool)`

SetSelectionRequired sets SelectionRequired field to given value.


### GetAccessToken

`func (o *MobileWorkspaceSelectionOut) GetAccessToken() string`

GetAccessToken returns the AccessToken field if non-nil, zero value otherwise.

### GetAccessTokenOk

`func (o *MobileWorkspaceSelectionOut) GetAccessTokenOk() (*string, bool)`

GetAccessTokenOk returns a tuple with the AccessToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccessToken

`func (o *MobileWorkspaceSelectionOut) SetAccessToken(v string)`

SetAccessToken sets AccessToken field to given value.


### GetRefreshToken

`func (o *MobileWorkspaceSelectionOut) GetRefreshToken() string`

GetRefreshToken returns the RefreshToken field if non-nil, zero value otherwise.

### GetRefreshTokenOk

`func (o *MobileWorkspaceSelectionOut) GetRefreshTokenOk() (*string, bool)`

GetRefreshTokenOk returns a tuple with the RefreshToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRefreshToken

`func (o *MobileWorkspaceSelectionOut) SetRefreshToken(v string)`

SetRefreshToken sets RefreshToken field to given value.


### GetTokenType

`func (o *MobileWorkspaceSelectionOut) GetTokenType() string`

GetTokenType returns the TokenType field if non-nil, zero value otherwise.

### GetTokenTypeOk

`func (o *MobileWorkspaceSelectionOut) GetTokenTypeOk() (*string, bool)`

GetTokenTypeOk returns a tuple with the TokenType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTokenType

`func (o *MobileWorkspaceSelectionOut) SetTokenType(v string)`

SetTokenType sets TokenType field to given value.


### GetExpiresIn

`func (o *MobileWorkspaceSelectionOut) GetExpiresIn() int32`

GetExpiresIn returns the ExpiresIn field if non-nil, zero value otherwise.

### GetExpiresInOk

`func (o *MobileWorkspaceSelectionOut) GetExpiresInOk() (*int32, bool)`

GetExpiresInOk returns a tuple with the ExpiresIn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpiresIn

`func (o *MobileWorkspaceSelectionOut) SetExpiresIn(v int32)`

SetExpiresIn sets ExpiresIn field to given value.


### GetUser

`func (o *MobileWorkspaceSelectionOut) GetUser() map[string]interface{}`

GetUser returns the User field if non-nil, zero value otherwise.

### GetUserOk

`func (o *MobileWorkspaceSelectionOut) GetUserOk() (*map[string]interface{}, bool)`

GetUserOk returns a tuple with the User field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUser

`func (o *MobileWorkspaceSelectionOut) SetUser(v map[string]interface{})`

SetUser sets User field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


