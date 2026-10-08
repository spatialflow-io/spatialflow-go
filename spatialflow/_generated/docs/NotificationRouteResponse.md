# NotificationRouteResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** |  | 
**Name** | **string** |  | 
**Provider** | **string** |  | 
**ProviderName** | **string** |  | 
**DestinationLabel** | **string** |  | 
**WebhookUrlConfigured** | **bool** |  | 
**EventTypes** | **[]string** |  | 
**IsEnabled** | **bool** |  | 
**IsDefault** | **bool** |  | 
**LastTestedAt** | Pointer to **NullableTime** |  | [optional] 
**LastSuccessAt** | Pointer to **NullableTime** |  | [optional] 
**LastFailureAt** | Pointer to **NullableTime** |  | [optional] 
**LastError** | Pointer to **NullableString** |  | [optional] 
**CreatedAt** | Pointer to **NullableTime** |  | [optional] 
**UpdatedAt** | Pointer to **NullableTime** |  | [optional] 
**UpdatedByEmail** | Pointer to **NullableString** |  | [optional] 

## Methods

### NewNotificationRouteResponse

`func NewNotificationRouteResponse(id string, name string, provider string, providerName string, destinationLabel string, webhookUrlConfigured bool, eventTypes []string, isEnabled bool, isDefault bool, ) *NotificationRouteResponse`

NewNotificationRouteResponse instantiates a new NotificationRouteResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNotificationRouteResponseWithDefaults

`func NewNotificationRouteResponseWithDefaults() *NotificationRouteResponse`

NewNotificationRouteResponseWithDefaults instantiates a new NotificationRouteResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *NotificationRouteResponse) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *NotificationRouteResponse) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *NotificationRouteResponse) SetId(v string)`

SetId sets Id field to given value.


### GetName

`func (o *NotificationRouteResponse) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *NotificationRouteResponse) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *NotificationRouteResponse) SetName(v string)`

SetName sets Name field to given value.


### GetProvider

`func (o *NotificationRouteResponse) GetProvider() string`

GetProvider returns the Provider field if non-nil, zero value otherwise.

### GetProviderOk

`func (o *NotificationRouteResponse) GetProviderOk() (*string, bool)`

GetProviderOk returns a tuple with the Provider field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProvider

`func (o *NotificationRouteResponse) SetProvider(v string)`

SetProvider sets Provider field to given value.


### GetProviderName

`func (o *NotificationRouteResponse) GetProviderName() string`

GetProviderName returns the ProviderName field if non-nil, zero value otherwise.

### GetProviderNameOk

`func (o *NotificationRouteResponse) GetProviderNameOk() (*string, bool)`

GetProviderNameOk returns a tuple with the ProviderName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProviderName

`func (o *NotificationRouteResponse) SetProviderName(v string)`

SetProviderName sets ProviderName field to given value.


### GetDestinationLabel

`func (o *NotificationRouteResponse) GetDestinationLabel() string`

GetDestinationLabel returns the DestinationLabel field if non-nil, zero value otherwise.

### GetDestinationLabelOk

`func (o *NotificationRouteResponse) GetDestinationLabelOk() (*string, bool)`

GetDestinationLabelOk returns a tuple with the DestinationLabel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDestinationLabel

`func (o *NotificationRouteResponse) SetDestinationLabel(v string)`

SetDestinationLabel sets DestinationLabel field to given value.


### GetWebhookUrlConfigured

`func (o *NotificationRouteResponse) GetWebhookUrlConfigured() bool`

GetWebhookUrlConfigured returns the WebhookUrlConfigured field if non-nil, zero value otherwise.

### GetWebhookUrlConfiguredOk

`func (o *NotificationRouteResponse) GetWebhookUrlConfiguredOk() (*bool, bool)`

GetWebhookUrlConfiguredOk returns a tuple with the WebhookUrlConfigured field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWebhookUrlConfigured

`func (o *NotificationRouteResponse) SetWebhookUrlConfigured(v bool)`

SetWebhookUrlConfigured sets WebhookUrlConfigured field to given value.


### GetEventTypes

`func (o *NotificationRouteResponse) GetEventTypes() []string`

GetEventTypes returns the EventTypes field if non-nil, zero value otherwise.

### GetEventTypesOk

`func (o *NotificationRouteResponse) GetEventTypesOk() (*[]string, bool)`

GetEventTypesOk returns a tuple with the EventTypes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEventTypes

`func (o *NotificationRouteResponse) SetEventTypes(v []string)`

SetEventTypes sets EventTypes field to given value.


### GetIsEnabled

`func (o *NotificationRouteResponse) GetIsEnabled() bool`

GetIsEnabled returns the IsEnabled field if non-nil, zero value otherwise.

### GetIsEnabledOk

`func (o *NotificationRouteResponse) GetIsEnabledOk() (*bool, bool)`

GetIsEnabledOk returns a tuple with the IsEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsEnabled

`func (o *NotificationRouteResponse) SetIsEnabled(v bool)`

SetIsEnabled sets IsEnabled field to given value.


### GetIsDefault

`func (o *NotificationRouteResponse) GetIsDefault() bool`

GetIsDefault returns the IsDefault field if non-nil, zero value otherwise.

### GetIsDefaultOk

`func (o *NotificationRouteResponse) GetIsDefaultOk() (*bool, bool)`

GetIsDefaultOk returns a tuple with the IsDefault field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsDefault

`func (o *NotificationRouteResponse) SetIsDefault(v bool)`

SetIsDefault sets IsDefault field to given value.


### GetLastTestedAt

`func (o *NotificationRouteResponse) GetLastTestedAt() time.Time`

GetLastTestedAt returns the LastTestedAt field if non-nil, zero value otherwise.

### GetLastTestedAtOk

`func (o *NotificationRouteResponse) GetLastTestedAtOk() (*time.Time, bool)`

GetLastTestedAtOk returns a tuple with the LastTestedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastTestedAt

`func (o *NotificationRouteResponse) SetLastTestedAt(v time.Time)`

SetLastTestedAt sets LastTestedAt field to given value.

### HasLastTestedAt

`func (o *NotificationRouteResponse) HasLastTestedAt() bool`

HasLastTestedAt returns a boolean if a field has been set.

### SetLastTestedAtNil

`func (o *NotificationRouteResponse) SetLastTestedAtNil(b bool)`

 SetLastTestedAtNil sets the value for LastTestedAt to be an explicit nil

### UnsetLastTestedAt
`func (o *NotificationRouteResponse) UnsetLastTestedAt()`

UnsetLastTestedAt ensures that no value is present for LastTestedAt, not even an explicit nil
### GetLastSuccessAt

`func (o *NotificationRouteResponse) GetLastSuccessAt() time.Time`

GetLastSuccessAt returns the LastSuccessAt field if non-nil, zero value otherwise.

### GetLastSuccessAtOk

`func (o *NotificationRouteResponse) GetLastSuccessAtOk() (*time.Time, bool)`

GetLastSuccessAtOk returns a tuple with the LastSuccessAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastSuccessAt

`func (o *NotificationRouteResponse) SetLastSuccessAt(v time.Time)`

SetLastSuccessAt sets LastSuccessAt field to given value.

### HasLastSuccessAt

`func (o *NotificationRouteResponse) HasLastSuccessAt() bool`

HasLastSuccessAt returns a boolean if a field has been set.

### SetLastSuccessAtNil

`func (o *NotificationRouteResponse) SetLastSuccessAtNil(b bool)`

 SetLastSuccessAtNil sets the value for LastSuccessAt to be an explicit nil

### UnsetLastSuccessAt
`func (o *NotificationRouteResponse) UnsetLastSuccessAt()`

UnsetLastSuccessAt ensures that no value is present for LastSuccessAt, not even an explicit nil
### GetLastFailureAt

`func (o *NotificationRouteResponse) GetLastFailureAt() time.Time`

GetLastFailureAt returns the LastFailureAt field if non-nil, zero value otherwise.

### GetLastFailureAtOk

`func (o *NotificationRouteResponse) GetLastFailureAtOk() (*time.Time, bool)`

GetLastFailureAtOk returns a tuple with the LastFailureAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastFailureAt

`func (o *NotificationRouteResponse) SetLastFailureAt(v time.Time)`

SetLastFailureAt sets LastFailureAt field to given value.

### HasLastFailureAt

`func (o *NotificationRouteResponse) HasLastFailureAt() bool`

HasLastFailureAt returns a boolean if a field has been set.

### SetLastFailureAtNil

`func (o *NotificationRouteResponse) SetLastFailureAtNil(b bool)`

 SetLastFailureAtNil sets the value for LastFailureAt to be an explicit nil

### UnsetLastFailureAt
`func (o *NotificationRouteResponse) UnsetLastFailureAt()`

UnsetLastFailureAt ensures that no value is present for LastFailureAt, not even an explicit nil
### GetLastError

`func (o *NotificationRouteResponse) GetLastError() string`

GetLastError returns the LastError field if non-nil, zero value otherwise.

### GetLastErrorOk

`func (o *NotificationRouteResponse) GetLastErrorOk() (*string, bool)`

GetLastErrorOk returns a tuple with the LastError field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastError

`func (o *NotificationRouteResponse) SetLastError(v string)`

SetLastError sets LastError field to given value.

### HasLastError

`func (o *NotificationRouteResponse) HasLastError() bool`

HasLastError returns a boolean if a field has been set.

### SetLastErrorNil

`func (o *NotificationRouteResponse) SetLastErrorNil(b bool)`

 SetLastErrorNil sets the value for LastError to be an explicit nil

### UnsetLastError
`func (o *NotificationRouteResponse) UnsetLastError()`

UnsetLastError ensures that no value is present for LastError, not even an explicit nil
### GetCreatedAt

`func (o *NotificationRouteResponse) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *NotificationRouteResponse) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *NotificationRouteResponse) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *NotificationRouteResponse) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### SetCreatedAtNil

`func (o *NotificationRouteResponse) SetCreatedAtNil(b bool)`

 SetCreatedAtNil sets the value for CreatedAt to be an explicit nil

### UnsetCreatedAt
`func (o *NotificationRouteResponse) UnsetCreatedAt()`

UnsetCreatedAt ensures that no value is present for CreatedAt, not even an explicit nil
### GetUpdatedAt

`func (o *NotificationRouteResponse) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *NotificationRouteResponse) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *NotificationRouteResponse) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *NotificationRouteResponse) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.

### SetUpdatedAtNil

`func (o *NotificationRouteResponse) SetUpdatedAtNil(b bool)`

 SetUpdatedAtNil sets the value for UpdatedAt to be an explicit nil

### UnsetUpdatedAt
`func (o *NotificationRouteResponse) UnsetUpdatedAt()`

UnsetUpdatedAt ensures that no value is present for UpdatedAt, not even an explicit nil
### GetUpdatedByEmail

`func (o *NotificationRouteResponse) GetUpdatedByEmail() string`

GetUpdatedByEmail returns the UpdatedByEmail field if non-nil, zero value otherwise.

### GetUpdatedByEmailOk

`func (o *NotificationRouteResponse) GetUpdatedByEmailOk() (*string, bool)`

GetUpdatedByEmailOk returns a tuple with the UpdatedByEmail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedByEmail

`func (o *NotificationRouteResponse) SetUpdatedByEmail(v string)`

SetUpdatedByEmail sets UpdatedByEmail field to given value.

### HasUpdatedByEmail

`func (o *NotificationRouteResponse) HasUpdatedByEmail() bool`

HasUpdatedByEmail returns a boolean if a field has been set.

### SetUpdatedByEmailNil

`func (o *NotificationRouteResponse) SetUpdatedByEmailNil(b bool)`

 SetUpdatedByEmailNil sets the value for UpdatedByEmail to be an explicit nil

### UnsetUpdatedByEmail
`func (o *NotificationRouteResponse) UnsetUpdatedByEmail()`

UnsetUpdatedByEmail ensures that no value is present for UpdatedByEmail, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


