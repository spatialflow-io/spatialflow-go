# SAMLConfigOut

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** |  | 
**EntityId** | **string** |  | 
**SsoUrl** | **string** |  | 
**Certificate** | **string** |  | 
**CoveredDomain** | **string** |  | 
**IsEnabled** | **bool** |  | 
**DomainVerified** | **bool** |  | 
**DomainVerifiedAt** | Pointer to **NullableTime** |  | [optional] 
**DomainVerificationDnsName** | **string** |  | 
**DomainVerificationDnsValue** | **string** |  | 
**CreatedAt** | **time.Time** |  | 
**UpdatedAt** | **time.Time** |  | 

## Methods

### NewSAMLConfigOut

`func NewSAMLConfigOut(id string, entityId string, ssoUrl string, certificate string, coveredDomain string, isEnabled bool, domainVerified bool, domainVerificationDnsName string, domainVerificationDnsValue string, createdAt time.Time, updatedAt time.Time, ) *SAMLConfigOut`

NewSAMLConfigOut instantiates a new SAMLConfigOut object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSAMLConfigOutWithDefaults

`func NewSAMLConfigOutWithDefaults() *SAMLConfigOut`

NewSAMLConfigOutWithDefaults instantiates a new SAMLConfigOut object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *SAMLConfigOut) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *SAMLConfigOut) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *SAMLConfigOut) SetId(v string)`

SetId sets Id field to given value.


### GetEntityId

`func (o *SAMLConfigOut) GetEntityId() string`

GetEntityId returns the EntityId field if non-nil, zero value otherwise.

### GetEntityIdOk

`func (o *SAMLConfigOut) GetEntityIdOk() (*string, bool)`

GetEntityIdOk returns a tuple with the EntityId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEntityId

`func (o *SAMLConfigOut) SetEntityId(v string)`

SetEntityId sets EntityId field to given value.


### GetSsoUrl

`func (o *SAMLConfigOut) GetSsoUrl() string`

GetSsoUrl returns the SsoUrl field if non-nil, zero value otherwise.

### GetSsoUrlOk

`func (o *SAMLConfigOut) GetSsoUrlOk() (*string, bool)`

GetSsoUrlOk returns a tuple with the SsoUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSsoUrl

`func (o *SAMLConfigOut) SetSsoUrl(v string)`

SetSsoUrl sets SsoUrl field to given value.


### GetCertificate

`func (o *SAMLConfigOut) GetCertificate() string`

GetCertificate returns the Certificate field if non-nil, zero value otherwise.

### GetCertificateOk

`func (o *SAMLConfigOut) GetCertificateOk() (*string, bool)`

GetCertificateOk returns a tuple with the Certificate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCertificate

`func (o *SAMLConfigOut) SetCertificate(v string)`

SetCertificate sets Certificate field to given value.


### GetCoveredDomain

`func (o *SAMLConfigOut) GetCoveredDomain() string`

GetCoveredDomain returns the CoveredDomain field if non-nil, zero value otherwise.

### GetCoveredDomainOk

`func (o *SAMLConfigOut) GetCoveredDomainOk() (*string, bool)`

GetCoveredDomainOk returns a tuple with the CoveredDomain field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCoveredDomain

`func (o *SAMLConfigOut) SetCoveredDomain(v string)`

SetCoveredDomain sets CoveredDomain field to given value.


### GetIsEnabled

`func (o *SAMLConfigOut) GetIsEnabled() bool`

GetIsEnabled returns the IsEnabled field if non-nil, zero value otherwise.

### GetIsEnabledOk

`func (o *SAMLConfigOut) GetIsEnabledOk() (*bool, bool)`

GetIsEnabledOk returns a tuple with the IsEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsEnabled

`func (o *SAMLConfigOut) SetIsEnabled(v bool)`

SetIsEnabled sets IsEnabled field to given value.


### GetDomainVerified

`func (o *SAMLConfigOut) GetDomainVerified() bool`

GetDomainVerified returns the DomainVerified field if non-nil, zero value otherwise.

### GetDomainVerifiedOk

`func (o *SAMLConfigOut) GetDomainVerifiedOk() (*bool, bool)`

GetDomainVerifiedOk returns a tuple with the DomainVerified field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDomainVerified

`func (o *SAMLConfigOut) SetDomainVerified(v bool)`

SetDomainVerified sets DomainVerified field to given value.


### GetDomainVerifiedAt

`func (o *SAMLConfigOut) GetDomainVerifiedAt() time.Time`

GetDomainVerifiedAt returns the DomainVerifiedAt field if non-nil, zero value otherwise.

### GetDomainVerifiedAtOk

`func (o *SAMLConfigOut) GetDomainVerifiedAtOk() (*time.Time, bool)`

GetDomainVerifiedAtOk returns a tuple with the DomainVerifiedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDomainVerifiedAt

`func (o *SAMLConfigOut) SetDomainVerifiedAt(v time.Time)`

SetDomainVerifiedAt sets DomainVerifiedAt field to given value.

### HasDomainVerifiedAt

`func (o *SAMLConfigOut) HasDomainVerifiedAt() bool`

HasDomainVerifiedAt returns a boolean if a field has been set.

### SetDomainVerifiedAtNil

`func (o *SAMLConfigOut) SetDomainVerifiedAtNil(b bool)`

 SetDomainVerifiedAtNil sets the value for DomainVerifiedAt to be an explicit nil

### UnsetDomainVerifiedAt
`func (o *SAMLConfigOut) UnsetDomainVerifiedAt()`

UnsetDomainVerifiedAt ensures that no value is present for DomainVerifiedAt, not even an explicit nil
### GetDomainVerificationDnsName

`func (o *SAMLConfigOut) GetDomainVerificationDnsName() string`

GetDomainVerificationDnsName returns the DomainVerificationDnsName field if non-nil, zero value otherwise.

### GetDomainVerificationDnsNameOk

`func (o *SAMLConfigOut) GetDomainVerificationDnsNameOk() (*string, bool)`

GetDomainVerificationDnsNameOk returns a tuple with the DomainVerificationDnsName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDomainVerificationDnsName

`func (o *SAMLConfigOut) SetDomainVerificationDnsName(v string)`

SetDomainVerificationDnsName sets DomainVerificationDnsName field to given value.


### GetDomainVerificationDnsValue

`func (o *SAMLConfigOut) GetDomainVerificationDnsValue() string`

GetDomainVerificationDnsValue returns the DomainVerificationDnsValue field if non-nil, zero value otherwise.

### GetDomainVerificationDnsValueOk

`func (o *SAMLConfigOut) GetDomainVerificationDnsValueOk() (*string, bool)`

GetDomainVerificationDnsValueOk returns a tuple with the DomainVerificationDnsValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDomainVerificationDnsValue

`func (o *SAMLConfigOut) SetDomainVerificationDnsValue(v string)`

SetDomainVerificationDnsValue sets DomainVerificationDnsValue field to given value.


### GetCreatedAt

`func (o *SAMLConfigOut) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *SAMLConfigOut) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *SAMLConfigOut) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.


### GetUpdatedAt

`func (o *SAMLConfigOut) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *SAMLConfigOut) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *SAMLConfigOut) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


