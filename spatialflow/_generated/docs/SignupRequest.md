# SignupRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Email** | **string** |  | 
**Password** | **string** |  | 
**Name** | Pointer to **NullableString** |  | [optional] 
**Company** | Pointer to **NullableString** |  | [optional] 
**UseCase** | Pointer to **NullableString** |  | [optional] 
**SelectedPlan** | Pointer to **string** |  | [optional] [default to "free"]
**UtmSource** | Pointer to **NullableString** |  | [optional] 
**UtmMedium** | Pointer to **NullableString** |  | [optional] 
**UtmCampaign** | Pointer to **NullableString** |  | [optional] 
**UtmTerm** | Pointer to **NullableString** |  | [optional] 
**UtmContent** | Pointer to **NullableString** |  | [optional] 
**Website** | Pointer to **NullableString** |  | [optional] 

## Methods

### NewSignupRequest

`func NewSignupRequest(email string, password string, ) *SignupRequest`

NewSignupRequest instantiates a new SignupRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSignupRequestWithDefaults

`func NewSignupRequestWithDefaults() *SignupRequest`

NewSignupRequestWithDefaults instantiates a new SignupRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEmail

`func (o *SignupRequest) GetEmail() string`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *SignupRequest) GetEmailOk() (*string, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *SignupRequest) SetEmail(v string)`

SetEmail sets Email field to given value.


### GetPassword

`func (o *SignupRequest) GetPassword() string`

GetPassword returns the Password field if non-nil, zero value otherwise.

### GetPasswordOk

`func (o *SignupRequest) GetPasswordOk() (*string, bool)`

GetPasswordOk returns a tuple with the Password field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPassword

`func (o *SignupRequest) SetPassword(v string)`

SetPassword sets Password field to given value.


### GetName

`func (o *SignupRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *SignupRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *SignupRequest) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *SignupRequest) HasName() bool`

HasName returns a boolean if a field has been set.

### SetNameNil

`func (o *SignupRequest) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *SignupRequest) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetCompany

`func (o *SignupRequest) GetCompany() string`

GetCompany returns the Company field if non-nil, zero value otherwise.

### GetCompanyOk

`func (o *SignupRequest) GetCompanyOk() (*string, bool)`

GetCompanyOk returns a tuple with the Company field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompany

`func (o *SignupRequest) SetCompany(v string)`

SetCompany sets Company field to given value.

### HasCompany

`func (o *SignupRequest) HasCompany() bool`

HasCompany returns a boolean if a field has been set.

### SetCompanyNil

`func (o *SignupRequest) SetCompanyNil(b bool)`

 SetCompanyNil sets the value for Company to be an explicit nil

### UnsetCompany
`func (o *SignupRequest) UnsetCompany()`

UnsetCompany ensures that no value is present for Company, not even an explicit nil
### GetUseCase

`func (o *SignupRequest) GetUseCase() string`

GetUseCase returns the UseCase field if non-nil, zero value otherwise.

### GetUseCaseOk

`func (o *SignupRequest) GetUseCaseOk() (*string, bool)`

GetUseCaseOk returns a tuple with the UseCase field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUseCase

`func (o *SignupRequest) SetUseCase(v string)`

SetUseCase sets UseCase field to given value.

### HasUseCase

`func (o *SignupRequest) HasUseCase() bool`

HasUseCase returns a boolean if a field has been set.

### SetUseCaseNil

`func (o *SignupRequest) SetUseCaseNil(b bool)`

 SetUseCaseNil sets the value for UseCase to be an explicit nil

### UnsetUseCase
`func (o *SignupRequest) UnsetUseCase()`

UnsetUseCase ensures that no value is present for UseCase, not even an explicit nil
### GetSelectedPlan

`func (o *SignupRequest) GetSelectedPlan() string`

GetSelectedPlan returns the SelectedPlan field if non-nil, zero value otherwise.

### GetSelectedPlanOk

`func (o *SignupRequest) GetSelectedPlanOk() (*string, bool)`

GetSelectedPlanOk returns a tuple with the SelectedPlan field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSelectedPlan

`func (o *SignupRequest) SetSelectedPlan(v string)`

SetSelectedPlan sets SelectedPlan field to given value.

### HasSelectedPlan

`func (o *SignupRequest) HasSelectedPlan() bool`

HasSelectedPlan returns a boolean if a field has been set.

### GetUtmSource

`func (o *SignupRequest) GetUtmSource() string`

GetUtmSource returns the UtmSource field if non-nil, zero value otherwise.

### GetUtmSourceOk

`func (o *SignupRequest) GetUtmSourceOk() (*string, bool)`

GetUtmSourceOk returns a tuple with the UtmSource field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUtmSource

`func (o *SignupRequest) SetUtmSource(v string)`

SetUtmSource sets UtmSource field to given value.

### HasUtmSource

`func (o *SignupRequest) HasUtmSource() bool`

HasUtmSource returns a boolean if a field has been set.

### SetUtmSourceNil

`func (o *SignupRequest) SetUtmSourceNil(b bool)`

 SetUtmSourceNil sets the value for UtmSource to be an explicit nil

### UnsetUtmSource
`func (o *SignupRequest) UnsetUtmSource()`

UnsetUtmSource ensures that no value is present for UtmSource, not even an explicit nil
### GetUtmMedium

`func (o *SignupRequest) GetUtmMedium() string`

GetUtmMedium returns the UtmMedium field if non-nil, zero value otherwise.

### GetUtmMediumOk

`func (o *SignupRequest) GetUtmMediumOk() (*string, bool)`

GetUtmMediumOk returns a tuple with the UtmMedium field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUtmMedium

`func (o *SignupRequest) SetUtmMedium(v string)`

SetUtmMedium sets UtmMedium field to given value.

### HasUtmMedium

`func (o *SignupRequest) HasUtmMedium() bool`

HasUtmMedium returns a boolean if a field has been set.

### SetUtmMediumNil

`func (o *SignupRequest) SetUtmMediumNil(b bool)`

 SetUtmMediumNil sets the value for UtmMedium to be an explicit nil

### UnsetUtmMedium
`func (o *SignupRequest) UnsetUtmMedium()`

UnsetUtmMedium ensures that no value is present for UtmMedium, not even an explicit nil
### GetUtmCampaign

`func (o *SignupRequest) GetUtmCampaign() string`

GetUtmCampaign returns the UtmCampaign field if non-nil, zero value otherwise.

### GetUtmCampaignOk

`func (o *SignupRequest) GetUtmCampaignOk() (*string, bool)`

GetUtmCampaignOk returns a tuple with the UtmCampaign field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUtmCampaign

`func (o *SignupRequest) SetUtmCampaign(v string)`

SetUtmCampaign sets UtmCampaign field to given value.

### HasUtmCampaign

`func (o *SignupRequest) HasUtmCampaign() bool`

HasUtmCampaign returns a boolean if a field has been set.

### SetUtmCampaignNil

`func (o *SignupRequest) SetUtmCampaignNil(b bool)`

 SetUtmCampaignNil sets the value for UtmCampaign to be an explicit nil

### UnsetUtmCampaign
`func (o *SignupRequest) UnsetUtmCampaign()`

UnsetUtmCampaign ensures that no value is present for UtmCampaign, not even an explicit nil
### GetUtmTerm

`func (o *SignupRequest) GetUtmTerm() string`

GetUtmTerm returns the UtmTerm field if non-nil, zero value otherwise.

### GetUtmTermOk

`func (o *SignupRequest) GetUtmTermOk() (*string, bool)`

GetUtmTermOk returns a tuple with the UtmTerm field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUtmTerm

`func (o *SignupRequest) SetUtmTerm(v string)`

SetUtmTerm sets UtmTerm field to given value.

### HasUtmTerm

`func (o *SignupRequest) HasUtmTerm() bool`

HasUtmTerm returns a boolean if a field has been set.

### SetUtmTermNil

`func (o *SignupRequest) SetUtmTermNil(b bool)`

 SetUtmTermNil sets the value for UtmTerm to be an explicit nil

### UnsetUtmTerm
`func (o *SignupRequest) UnsetUtmTerm()`

UnsetUtmTerm ensures that no value is present for UtmTerm, not even an explicit nil
### GetUtmContent

`func (o *SignupRequest) GetUtmContent() string`

GetUtmContent returns the UtmContent field if non-nil, zero value otherwise.

### GetUtmContentOk

`func (o *SignupRequest) GetUtmContentOk() (*string, bool)`

GetUtmContentOk returns a tuple with the UtmContent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUtmContent

`func (o *SignupRequest) SetUtmContent(v string)`

SetUtmContent sets UtmContent field to given value.

### HasUtmContent

`func (o *SignupRequest) HasUtmContent() bool`

HasUtmContent returns a boolean if a field has been set.

### SetUtmContentNil

`func (o *SignupRequest) SetUtmContentNil(b bool)`

 SetUtmContentNil sets the value for UtmContent to be an explicit nil

### UnsetUtmContent
`func (o *SignupRequest) UnsetUtmContent()`

UnsetUtmContent ensures that no value is present for UtmContent, not even an explicit nil
### GetWebsite

`func (o *SignupRequest) GetWebsite() string`

GetWebsite returns the Website field if non-nil, zero value otherwise.

### GetWebsiteOk

`func (o *SignupRequest) GetWebsiteOk() (*string, bool)`

GetWebsiteOk returns a tuple with the Website field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWebsite

`func (o *SignupRequest) SetWebsite(v string)`

SetWebsite sets Website field to given value.

### HasWebsite

`func (o *SignupRequest) HasWebsite() bool`

HasWebsite returns a boolean if a field has been set.

### SetWebsiteNil

`func (o *SignupRequest) SetWebsiteNil(b bool)`

 SetWebsiteNil sets the value for Website to be an explicit nil

### UnsetWebsite
`func (o *SignupRequest) UnsetWebsite()`

UnsetWebsite ensures that no value is present for Website, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


