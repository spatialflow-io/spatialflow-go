# RegisterSchema

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Email** | **string** |  | 
**Password** | **string** |  | 
**Name** | **string** |  | 
**Company** | Pointer to **NullableString** |  | [optional] 
**SignupType** | Pointer to **NullableString** |  | [optional] 
**UseCase** | Pointer to **NullableString** |  | [optional] 
**UtmSource** | Pointer to **NullableString** |  | [optional] 
**UtmMedium** | Pointer to **NullableString** |  | [optional] 
**UtmCampaign** | Pointer to **NullableString** |  | [optional] 

## Methods

### NewRegisterSchema

`func NewRegisterSchema(email string, password string, name string, ) *RegisterSchema`

NewRegisterSchema instantiates a new RegisterSchema object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRegisterSchemaWithDefaults

`func NewRegisterSchemaWithDefaults() *RegisterSchema`

NewRegisterSchemaWithDefaults instantiates a new RegisterSchema object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEmail

`func (o *RegisterSchema) GetEmail() string`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *RegisterSchema) GetEmailOk() (*string, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *RegisterSchema) SetEmail(v string)`

SetEmail sets Email field to given value.


### GetPassword

`func (o *RegisterSchema) GetPassword() string`

GetPassword returns the Password field if non-nil, zero value otherwise.

### GetPasswordOk

`func (o *RegisterSchema) GetPasswordOk() (*string, bool)`

GetPasswordOk returns a tuple with the Password field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPassword

`func (o *RegisterSchema) SetPassword(v string)`

SetPassword sets Password field to given value.


### GetName

`func (o *RegisterSchema) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *RegisterSchema) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *RegisterSchema) SetName(v string)`

SetName sets Name field to given value.


### GetCompany

`func (o *RegisterSchema) GetCompany() string`

GetCompany returns the Company field if non-nil, zero value otherwise.

### GetCompanyOk

`func (o *RegisterSchema) GetCompanyOk() (*string, bool)`

GetCompanyOk returns a tuple with the Company field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompany

`func (o *RegisterSchema) SetCompany(v string)`

SetCompany sets Company field to given value.

### HasCompany

`func (o *RegisterSchema) HasCompany() bool`

HasCompany returns a boolean if a field has been set.

### SetCompanyNil

`func (o *RegisterSchema) SetCompanyNil(b bool)`

 SetCompanyNil sets the value for Company to be an explicit nil

### UnsetCompany
`func (o *RegisterSchema) UnsetCompany()`

UnsetCompany ensures that no value is present for Company, not even an explicit nil
### GetSignupType

`func (o *RegisterSchema) GetSignupType() string`

GetSignupType returns the SignupType field if non-nil, zero value otherwise.

### GetSignupTypeOk

`func (o *RegisterSchema) GetSignupTypeOk() (*string, bool)`

GetSignupTypeOk returns a tuple with the SignupType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSignupType

`func (o *RegisterSchema) SetSignupType(v string)`

SetSignupType sets SignupType field to given value.

### HasSignupType

`func (o *RegisterSchema) HasSignupType() bool`

HasSignupType returns a boolean if a field has been set.

### SetSignupTypeNil

`func (o *RegisterSchema) SetSignupTypeNil(b bool)`

 SetSignupTypeNil sets the value for SignupType to be an explicit nil

### UnsetSignupType
`func (o *RegisterSchema) UnsetSignupType()`

UnsetSignupType ensures that no value is present for SignupType, not even an explicit nil
### GetUseCase

`func (o *RegisterSchema) GetUseCase() string`

GetUseCase returns the UseCase field if non-nil, zero value otherwise.

### GetUseCaseOk

`func (o *RegisterSchema) GetUseCaseOk() (*string, bool)`

GetUseCaseOk returns a tuple with the UseCase field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUseCase

`func (o *RegisterSchema) SetUseCase(v string)`

SetUseCase sets UseCase field to given value.

### HasUseCase

`func (o *RegisterSchema) HasUseCase() bool`

HasUseCase returns a boolean if a field has been set.

### SetUseCaseNil

`func (o *RegisterSchema) SetUseCaseNil(b bool)`

 SetUseCaseNil sets the value for UseCase to be an explicit nil

### UnsetUseCase
`func (o *RegisterSchema) UnsetUseCase()`

UnsetUseCase ensures that no value is present for UseCase, not even an explicit nil
### GetUtmSource

`func (o *RegisterSchema) GetUtmSource() string`

GetUtmSource returns the UtmSource field if non-nil, zero value otherwise.

### GetUtmSourceOk

`func (o *RegisterSchema) GetUtmSourceOk() (*string, bool)`

GetUtmSourceOk returns a tuple with the UtmSource field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUtmSource

`func (o *RegisterSchema) SetUtmSource(v string)`

SetUtmSource sets UtmSource field to given value.

### HasUtmSource

`func (o *RegisterSchema) HasUtmSource() bool`

HasUtmSource returns a boolean if a field has been set.

### SetUtmSourceNil

`func (o *RegisterSchema) SetUtmSourceNil(b bool)`

 SetUtmSourceNil sets the value for UtmSource to be an explicit nil

### UnsetUtmSource
`func (o *RegisterSchema) UnsetUtmSource()`

UnsetUtmSource ensures that no value is present for UtmSource, not even an explicit nil
### GetUtmMedium

`func (o *RegisterSchema) GetUtmMedium() string`

GetUtmMedium returns the UtmMedium field if non-nil, zero value otherwise.

### GetUtmMediumOk

`func (o *RegisterSchema) GetUtmMediumOk() (*string, bool)`

GetUtmMediumOk returns a tuple with the UtmMedium field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUtmMedium

`func (o *RegisterSchema) SetUtmMedium(v string)`

SetUtmMedium sets UtmMedium field to given value.

### HasUtmMedium

`func (o *RegisterSchema) HasUtmMedium() bool`

HasUtmMedium returns a boolean if a field has been set.

### SetUtmMediumNil

`func (o *RegisterSchema) SetUtmMediumNil(b bool)`

 SetUtmMediumNil sets the value for UtmMedium to be an explicit nil

### UnsetUtmMedium
`func (o *RegisterSchema) UnsetUtmMedium()`

UnsetUtmMedium ensures that no value is present for UtmMedium, not even an explicit nil
### GetUtmCampaign

`func (o *RegisterSchema) GetUtmCampaign() string`

GetUtmCampaign returns the UtmCampaign field if non-nil, zero value otherwise.

### GetUtmCampaignOk

`func (o *RegisterSchema) GetUtmCampaignOk() (*string, bool)`

GetUtmCampaignOk returns a tuple with the UtmCampaign field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUtmCampaign

`func (o *RegisterSchema) SetUtmCampaign(v string)`

SetUtmCampaign sets UtmCampaign field to given value.

### HasUtmCampaign

`func (o *RegisterSchema) HasUtmCampaign() bool`

HasUtmCampaign returns a boolean if a field has been set.

### SetUtmCampaignNil

`func (o *RegisterSchema) SetUtmCampaignNil(b bool)`

 SetUtmCampaignNil sets the value for UtmCampaign to be an explicit nil

### UnsetUtmCampaign
`func (o *RegisterSchema) UnsetUtmCampaign()`

UnsetUtmCampaign ensures that no value is present for UtmCampaign, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


