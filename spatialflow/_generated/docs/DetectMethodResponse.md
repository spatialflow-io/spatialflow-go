# DetectMethodResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Method** | **string** |  | 
**WorkspaceSlug** | Pointer to **NullableString** |  | [optional] 
**IdpName** | Pointer to **NullableString** |  | [optional] 

## Methods

### NewDetectMethodResponse

`func NewDetectMethodResponse(method string, ) *DetectMethodResponse`

NewDetectMethodResponse instantiates a new DetectMethodResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDetectMethodResponseWithDefaults

`func NewDetectMethodResponseWithDefaults() *DetectMethodResponse`

NewDetectMethodResponseWithDefaults instantiates a new DetectMethodResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMethod

`func (o *DetectMethodResponse) GetMethod() string`

GetMethod returns the Method field if non-nil, zero value otherwise.

### GetMethodOk

`func (o *DetectMethodResponse) GetMethodOk() (*string, bool)`

GetMethodOk returns a tuple with the Method field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMethod

`func (o *DetectMethodResponse) SetMethod(v string)`

SetMethod sets Method field to given value.


### GetWorkspaceSlug

`func (o *DetectMethodResponse) GetWorkspaceSlug() string`

GetWorkspaceSlug returns the WorkspaceSlug field if non-nil, zero value otherwise.

### GetWorkspaceSlugOk

`func (o *DetectMethodResponse) GetWorkspaceSlugOk() (*string, bool)`

GetWorkspaceSlugOk returns a tuple with the WorkspaceSlug field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWorkspaceSlug

`func (o *DetectMethodResponse) SetWorkspaceSlug(v string)`

SetWorkspaceSlug sets WorkspaceSlug field to given value.

### HasWorkspaceSlug

`func (o *DetectMethodResponse) HasWorkspaceSlug() bool`

HasWorkspaceSlug returns a boolean if a field has been set.

### SetWorkspaceSlugNil

`func (o *DetectMethodResponse) SetWorkspaceSlugNil(b bool)`

 SetWorkspaceSlugNil sets the value for WorkspaceSlug to be an explicit nil

### UnsetWorkspaceSlug
`func (o *DetectMethodResponse) UnsetWorkspaceSlug()`

UnsetWorkspaceSlug ensures that no value is present for WorkspaceSlug, not even an explicit nil
### GetIdpName

`func (o *DetectMethodResponse) GetIdpName() string`

GetIdpName returns the IdpName field if non-nil, zero value otherwise.

### GetIdpNameOk

`func (o *DetectMethodResponse) GetIdpNameOk() (*string, bool)`

GetIdpNameOk returns a tuple with the IdpName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdpName

`func (o *DetectMethodResponse) SetIdpName(v string)`

SetIdpName sets IdpName field to given value.

### HasIdpName

`func (o *DetectMethodResponse) HasIdpName() bool`

HasIdpName returns a boolean if a field has been set.

### SetIdpNameNil

`func (o *DetectMethodResponse) SetIdpNameNil(b bool)`

 SetIdpNameNil sets the value for IdpName to be an explicit nil

### UnsetIdpName
`func (o *DetectMethodResponse) UnsetIdpName()`

UnsetIdpName ensures that no value is present for IdpName, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


