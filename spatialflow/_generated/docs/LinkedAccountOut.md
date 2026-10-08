# LinkedAccountOut

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Provider** | **string** |  | 
**DisplayName** | **string** |  | 
**LinkedAt** | **time.Time** |  | 
**AvatarUrl** | Pointer to **NullableString** |  | [optional] 
**ProfileUrl** | Pointer to **NullableString** |  | [optional] 

## Methods

### NewLinkedAccountOut

`func NewLinkedAccountOut(provider string, displayName string, linkedAt time.Time, ) *LinkedAccountOut`

NewLinkedAccountOut instantiates a new LinkedAccountOut object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLinkedAccountOutWithDefaults

`func NewLinkedAccountOutWithDefaults() *LinkedAccountOut`

NewLinkedAccountOutWithDefaults instantiates a new LinkedAccountOut object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetProvider

`func (o *LinkedAccountOut) GetProvider() string`

GetProvider returns the Provider field if non-nil, zero value otherwise.

### GetProviderOk

`func (o *LinkedAccountOut) GetProviderOk() (*string, bool)`

GetProviderOk returns a tuple with the Provider field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProvider

`func (o *LinkedAccountOut) SetProvider(v string)`

SetProvider sets Provider field to given value.


### GetDisplayName

`func (o *LinkedAccountOut) GetDisplayName() string`

GetDisplayName returns the DisplayName field if non-nil, zero value otherwise.

### GetDisplayNameOk

`func (o *LinkedAccountOut) GetDisplayNameOk() (*string, bool)`

GetDisplayNameOk returns a tuple with the DisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayName

`func (o *LinkedAccountOut) SetDisplayName(v string)`

SetDisplayName sets DisplayName field to given value.


### GetLinkedAt

`func (o *LinkedAccountOut) GetLinkedAt() time.Time`

GetLinkedAt returns the LinkedAt field if non-nil, zero value otherwise.

### GetLinkedAtOk

`func (o *LinkedAccountOut) GetLinkedAtOk() (*time.Time, bool)`

GetLinkedAtOk returns a tuple with the LinkedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinkedAt

`func (o *LinkedAccountOut) SetLinkedAt(v time.Time)`

SetLinkedAt sets LinkedAt field to given value.


### GetAvatarUrl

`func (o *LinkedAccountOut) GetAvatarUrl() string`

GetAvatarUrl returns the AvatarUrl field if non-nil, zero value otherwise.

### GetAvatarUrlOk

`func (o *LinkedAccountOut) GetAvatarUrlOk() (*string, bool)`

GetAvatarUrlOk returns a tuple with the AvatarUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvatarUrl

`func (o *LinkedAccountOut) SetAvatarUrl(v string)`

SetAvatarUrl sets AvatarUrl field to given value.

### HasAvatarUrl

`func (o *LinkedAccountOut) HasAvatarUrl() bool`

HasAvatarUrl returns a boolean if a field has been set.

### SetAvatarUrlNil

`func (o *LinkedAccountOut) SetAvatarUrlNil(b bool)`

 SetAvatarUrlNil sets the value for AvatarUrl to be an explicit nil

### UnsetAvatarUrl
`func (o *LinkedAccountOut) UnsetAvatarUrl()`

UnsetAvatarUrl ensures that no value is present for AvatarUrl, not even an explicit nil
### GetProfileUrl

`func (o *LinkedAccountOut) GetProfileUrl() string`

GetProfileUrl returns the ProfileUrl field if non-nil, zero value otherwise.

### GetProfileUrlOk

`func (o *LinkedAccountOut) GetProfileUrlOk() (*string, bool)`

GetProfileUrlOk returns a tuple with the ProfileUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProfileUrl

`func (o *LinkedAccountOut) SetProfileUrl(v string)`

SetProfileUrl sets ProfileUrl field to given value.

### HasProfileUrl

`func (o *LinkedAccountOut) HasProfileUrl() bool`

HasProfileUrl returns a boolean if a field has been set.

### SetProfileUrlNil

`func (o *LinkedAccountOut) SetProfileUrlNil(b bool)`

 SetProfileUrlNil sets the value for ProfileUrl to be an explicit nil

### UnsetProfileUrl
`func (o *LinkedAccountOut) UnsetProfileUrl()`

UnsetProfileUrl ensures that no value is present for ProfileUrl, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


