# SAMLConfigIn

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**EntityId** | **string** |  | 
**SsoUrl** | **string** |  | 
**Certificate** | **string** |  | 
**CoveredDomain** | **string** |  | 
**IsEnabled** | Pointer to **bool** |  | [optional] [default to true]

## Methods

### NewSAMLConfigIn

`func NewSAMLConfigIn(entityId string, ssoUrl string, certificate string, coveredDomain string, ) *SAMLConfigIn`

NewSAMLConfigIn instantiates a new SAMLConfigIn object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSAMLConfigInWithDefaults

`func NewSAMLConfigInWithDefaults() *SAMLConfigIn`

NewSAMLConfigInWithDefaults instantiates a new SAMLConfigIn object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEntityId

`func (o *SAMLConfigIn) GetEntityId() string`

GetEntityId returns the EntityId field if non-nil, zero value otherwise.

### GetEntityIdOk

`func (o *SAMLConfigIn) GetEntityIdOk() (*string, bool)`

GetEntityIdOk returns a tuple with the EntityId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEntityId

`func (o *SAMLConfigIn) SetEntityId(v string)`

SetEntityId sets EntityId field to given value.


### GetSsoUrl

`func (o *SAMLConfigIn) GetSsoUrl() string`

GetSsoUrl returns the SsoUrl field if non-nil, zero value otherwise.

### GetSsoUrlOk

`func (o *SAMLConfigIn) GetSsoUrlOk() (*string, bool)`

GetSsoUrlOk returns a tuple with the SsoUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSsoUrl

`func (o *SAMLConfigIn) SetSsoUrl(v string)`

SetSsoUrl sets SsoUrl field to given value.


### GetCertificate

`func (o *SAMLConfigIn) GetCertificate() string`

GetCertificate returns the Certificate field if non-nil, zero value otherwise.

### GetCertificateOk

`func (o *SAMLConfigIn) GetCertificateOk() (*string, bool)`

GetCertificateOk returns a tuple with the Certificate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCertificate

`func (o *SAMLConfigIn) SetCertificate(v string)`

SetCertificate sets Certificate field to given value.


### GetCoveredDomain

`func (o *SAMLConfigIn) GetCoveredDomain() string`

GetCoveredDomain returns the CoveredDomain field if non-nil, zero value otherwise.

### GetCoveredDomainOk

`func (o *SAMLConfigIn) GetCoveredDomainOk() (*string, bool)`

GetCoveredDomainOk returns a tuple with the CoveredDomain field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCoveredDomain

`func (o *SAMLConfigIn) SetCoveredDomain(v string)`

SetCoveredDomain sets CoveredDomain field to given value.


### GetIsEnabled

`func (o *SAMLConfigIn) GetIsEnabled() bool`

GetIsEnabled returns the IsEnabled field if non-nil, zero value otherwise.

### GetIsEnabledOk

`func (o *SAMLConfigIn) GetIsEnabledOk() (*bool, bool)`

GetIsEnabledOk returns a tuple with the IsEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsEnabled

`func (o *SAMLConfigIn) SetIsEnabled(v bool)`

SetIsEnabled sets IsEnabled field to given value.

### HasIsEnabled

`func (o *SAMLConfigIn) HasIsEnabled() bool`

HasIsEnabled returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


