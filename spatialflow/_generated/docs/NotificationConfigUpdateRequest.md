# NotificationConfigUpdateRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Provider** | Pointer to **NullableString** |  | [optional] 
**WebhookUrl** | Pointer to **NullableString** |  | [optional] 
**IsEnabled** | Pointer to **NullableBool** |  | [optional] 
**NotifyNewSignups** | Pointer to **NullableBool** |  | [optional] 
**NotifyAdminApprovals** | Pointer to **NullableBool** |  | [optional] 
**NotifySubscriptionChanges** | Pointer to **NullableBool** |  | [optional] 
**NotifyPaymentFailures** | Pointer to **NullableBool** |  | [optional] 
**NotifyPrivacyErasures** | Pointer to **NullableBool** |  | [optional] 
**NotifyDlqThreshold** | Pointer to **NullableBool** |  | [optional] 
**NotifyServiceHealth** | Pointer to **NullableBool** |  | [optional] 
**DlqThreshold** | Pointer to **NullableInt32** |  | [optional] 

## Methods

### NewNotificationConfigUpdateRequest

`func NewNotificationConfigUpdateRequest() *NotificationConfigUpdateRequest`

NewNotificationConfigUpdateRequest instantiates a new NotificationConfigUpdateRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNotificationConfigUpdateRequestWithDefaults

`func NewNotificationConfigUpdateRequestWithDefaults() *NotificationConfigUpdateRequest`

NewNotificationConfigUpdateRequestWithDefaults instantiates a new NotificationConfigUpdateRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetProvider

`func (o *NotificationConfigUpdateRequest) GetProvider() string`

GetProvider returns the Provider field if non-nil, zero value otherwise.

### GetProviderOk

`func (o *NotificationConfigUpdateRequest) GetProviderOk() (*string, bool)`

GetProviderOk returns a tuple with the Provider field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProvider

`func (o *NotificationConfigUpdateRequest) SetProvider(v string)`

SetProvider sets Provider field to given value.

### HasProvider

`func (o *NotificationConfigUpdateRequest) HasProvider() bool`

HasProvider returns a boolean if a field has been set.

### SetProviderNil

`func (o *NotificationConfigUpdateRequest) SetProviderNil(b bool)`

 SetProviderNil sets the value for Provider to be an explicit nil

### UnsetProvider
`func (o *NotificationConfigUpdateRequest) UnsetProvider()`

UnsetProvider ensures that no value is present for Provider, not even an explicit nil
### GetWebhookUrl

`func (o *NotificationConfigUpdateRequest) GetWebhookUrl() string`

GetWebhookUrl returns the WebhookUrl field if non-nil, zero value otherwise.

### GetWebhookUrlOk

`func (o *NotificationConfigUpdateRequest) GetWebhookUrlOk() (*string, bool)`

GetWebhookUrlOk returns a tuple with the WebhookUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWebhookUrl

`func (o *NotificationConfigUpdateRequest) SetWebhookUrl(v string)`

SetWebhookUrl sets WebhookUrl field to given value.

### HasWebhookUrl

`func (o *NotificationConfigUpdateRequest) HasWebhookUrl() bool`

HasWebhookUrl returns a boolean if a field has been set.

### SetWebhookUrlNil

`func (o *NotificationConfigUpdateRequest) SetWebhookUrlNil(b bool)`

 SetWebhookUrlNil sets the value for WebhookUrl to be an explicit nil

### UnsetWebhookUrl
`func (o *NotificationConfigUpdateRequest) UnsetWebhookUrl()`

UnsetWebhookUrl ensures that no value is present for WebhookUrl, not even an explicit nil
### GetIsEnabled

`func (o *NotificationConfigUpdateRequest) GetIsEnabled() bool`

GetIsEnabled returns the IsEnabled field if non-nil, zero value otherwise.

### GetIsEnabledOk

`func (o *NotificationConfigUpdateRequest) GetIsEnabledOk() (*bool, bool)`

GetIsEnabledOk returns a tuple with the IsEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsEnabled

`func (o *NotificationConfigUpdateRequest) SetIsEnabled(v bool)`

SetIsEnabled sets IsEnabled field to given value.

### HasIsEnabled

`func (o *NotificationConfigUpdateRequest) HasIsEnabled() bool`

HasIsEnabled returns a boolean if a field has been set.

### SetIsEnabledNil

`func (o *NotificationConfigUpdateRequest) SetIsEnabledNil(b bool)`

 SetIsEnabledNil sets the value for IsEnabled to be an explicit nil

### UnsetIsEnabled
`func (o *NotificationConfigUpdateRequest) UnsetIsEnabled()`

UnsetIsEnabled ensures that no value is present for IsEnabled, not even an explicit nil
### GetNotifyNewSignups

`func (o *NotificationConfigUpdateRequest) GetNotifyNewSignups() bool`

GetNotifyNewSignups returns the NotifyNewSignups field if non-nil, zero value otherwise.

### GetNotifyNewSignupsOk

`func (o *NotificationConfigUpdateRequest) GetNotifyNewSignupsOk() (*bool, bool)`

GetNotifyNewSignupsOk returns a tuple with the NotifyNewSignups field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotifyNewSignups

`func (o *NotificationConfigUpdateRequest) SetNotifyNewSignups(v bool)`

SetNotifyNewSignups sets NotifyNewSignups field to given value.

### HasNotifyNewSignups

`func (o *NotificationConfigUpdateRequest) HasNotifyNewSignups() bool`

HasNotifyNewSignups returns a boolean if a field has been set.

### SetNotifyNewSignupsNil

`func (o *NotificationConfigUpdateRequest) SetNotifyNewSignupsNil(b bool)`

 SetNotifyNewSignupsNil sets the value for NotifyNewSignups to be an explicit nil

### UnsetNotifyNewSignups
`func (o *NotificationConfigUpdateRequest) UnsetNotifyNewSignups()`

UnsetNotifyNewSignups ensures that no value is present for NotifyNewSignups, not even an explicit nil
### GetNotifyAdminApprovals

`func (o *NotificationConfigUpdateRequest) GetNotifyAdminApprovals() bool`

GetNotifyAdminApprovals returns the NotifyAdminApprovals field if non-nil, zero value otherwise.

### GetNotifyAdminApprovalsOk

`func (o *NotificationConfigUpdateRequest) GetNotifyAdminApprovalsOk() (*bool, bool)`

GetNotifyAdminApprovalsOk returns a tuple with the NotifyAdminApprovals field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotifyAdminApprovals

`func (o *NotificationConfigUpdateRequest) SetNotifyAdminApprovals(v bool)`

SetNotifyAdminApprovals sets NotifyAdminApprovals field to given value.

### HasNotifyAdminApprovals

`func (o *NotificationConfigUpdateRequest) HasNotifyAdminApprovals() bool`

HasNotifyAdminApprovals returns a boolean if a field has been set.

### SetNotifyAdminApprovalsNil

`func (o *NotificationConfigUpdateRequest) SetNotifyAdminApprovalsNil(b bool)`

 SetNotifyAdminApprovalsNil sets the value for NotifyAdminApprovals to be an explicit nil

### UnsetNotifyAdminApprovals
`func (o *NotificationConfigUpdateRequest) UnsetNotifyAdminApprovals()`

UnsetNotifyAdminApprovals ensures that no value is present for NotifyAdminApprovals, not even an explicit nil
### GetNotifySubscriptionChanges

`func (o *NotificationConfigUpdateRequest) GetNotifySubscriptionChanges() bool`

GetNotifySubscriptionChanges returns the NotifySubscriptionChanges field if non-nil, zero value otherwise.

### GetNotifySubscriptionChangesOk

`func (o *NotificationConfigUpdateRequest) GetNotifySubscriptionChangesOk() (*bool, bool)`

GetNotifySubscriptionChangesOk returns a tuple with the NotifySubscriptionChanges field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotifySubscriptionChanges

`func (o *NotificationConfigUpdateRequest) SetNotifySubscriptionChanges(v bool)`

SetNotifySubscriptionChanges sets NotifySubscriptionChanges field to given value.

### HasNotifySubscriptionChanges

`func (o *NotificationConfigUpdateRequest) HasNotifySubscriptionChanges() bool`

HasNotifySubscriptionChanges returns a boolean if a field has been set.

### SetNotifySubscriptionChangesNil

`func (o *NotificationConfigUpdateRequest) SetNotifySubscriptionChangesNil(b bool)`

 SetNotifySubscriptionChangesNil sets the value for NotifySubscriptionChanges to be an explicit nil

### UnsetNotifySubscriptionChanges
`func (o *NotificationConfigUpdateRequest) UnsetNotifySubscriptionChanges()`

UnsetNotifySubscriptionChanges ensures that no value is present for NotifySubscriptionChanges, not even an explicit nil
### GetNotifyPaymentFailures

`func (o *NotificationConfigUpdateRequest) GetNotifyPaymentFailures() bool`

GetNotifyPaymentFailures returns the NotifyPaymentFailures field if non-nil, zero value otherwise.

### GetNotifyPaymentFailuresOk

`func (o *NotificationConfigUpdateRequest) GetNotifyPaymentFailuresOk() (*bool, bool)`

GetNotifyPaymentFailuresOk returns a tuple with the NotifyPaymentFailures field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotifyPaymentFailures

`func (o *NotificationConfigUpdateRequest) SetNotifyPaymentFailures(v bool)`

SetNotifyPaymentFailures sets NotifyPaymentFailures field to given value.

### HasNotifyPaymentFailures

`func (o *NotificationConfigUpdateRequest) HasNotifyPaymentFailures() bool`

HasNotifyPaymentFailures returns a boolean if a field has been set.

### SetNotifyPaymentFailuresNil

`func (o *NotificationConfigUpdateRequest) SetNotifyPaymentFailuresNil(b bool)`

 SetNotifyPaymentFailuresNil sets the value for NotifyPaymentFailures to be an explicit nil

### UnsetNotifyPaymentFailures
`func (o *NotificationConfigUpdateRequest) UnsetNotifyPaymentFailures()`

UnsetNotifyPaymentFailures ensures that no value is present for NotifyPaymentFailures, not even an explicit nil
### GetNotifyPrivacyErasures

`func (o *NotificationConfigUpdateRequest) GetNotifyPrivacyErasures() bool`

GetNotifyPrivacyErasures returns the NotifyPrivacyErasures field if non-nil, zero value otherwise.

### GetNotifyPrivacyErasuresOk

`func (o *NotificationConfigUpdateRequest) GetNotifyPrivacyErasuresOk() (*bool, bool)`

GetNotifyPrivacyErasuresOk returns a tuple with the NotifyPrivacyErasures field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotifyPrivacyErasures

`func (o *NotificationConfigUpdateRequest) SetNotifyPrivacyErasures(v bool)`

SetNotifyPrivacyErasures sets NotifyPrivacyErasures field to given value.

### HasNotifyPrivacyErasures

`func (o *NotificationConfigUpdateRequest) HasNotifyPrivacyErasures() bool`

HasNotifyPrivacyErasures returns a boolean if a field has been set.

### SetNotifyPrivacyErasuresNil

`func (o *NotificationConfigUpdateRequest) SetNotifyPrivacyErasuresNil(b bool)`

 SetNotifyPrivacyErasuresNil sets the value for NotifyPrivacyErasures to be an explicit nil

### UnsetNotifyPrivacyErasures
`func (o *NotificationConfigUpdateRequest) UnsetNotifyPrivacyErasures()`

UnsetNotifyPrivacyErasures ensures that no value is present for NotifyPrivacyErasures, not even an explicit nil
### GetNotifyDlqThreshold

`func (o *NotificationConfigUpdateRequest) GetNotifyDlqThreshold() bool`

GetNotifyDlqThreshold returns the NotifyDlqThreshold field if non-nil, zero value otherwise.

### GetNotifyDlqThresholdOk

`func (o *NotificationConfigUpdateRequest) GetNotifyDlqThresholdOk() (*bool, bool)`

GetNotifyDlqThresholdOk returns a tuple with the NotifyDlqThreshold field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotifyDlqThreshold

`func (o *NotificationConfigUpdateRequest) SetNotifyDlqThreshold(v bool)`

SetNotifyDlqThreshold sets NotifyDlqThreshold field to given value.

### HasNotifyDlqThreshold

`func (o *NotificationConfigUpdateRequest) HasNotifyDlqThreshold() bool`

HasNotifyDlqThreshold returns a boolean if a field has been set.

### SetNotifyDlqThresholdNil

`func (o *NotificationConfigUpdateRequest) SetNotifyDlqThresholdNil(b bool)`

 SetNotifyDlqThresholdNil sets the value for NotifyDlqThreshold to be an explicit nil

### UnsetNotifyDlqThreshold
`func (o *NotificationConfigUpdateRequest) UnsetNotifyDlqThreshold()`

UnsetNotifyDlqThreshold ensures that no value is present for NotifyDlqThreshold, not even an explicit nil
### GetNotifyServiceHealth

`func (o *NotificationConfigUpdateRequest) GetNotifyServiceHealth() bool`

GetNotifyServiceHealth returns the NotifyServiceHealth field if non-nil, zero value otherwise.

### GetNotifyServiceHealthOk

`func (o *NotificationConfigUpdateRequest) GetNotifyServiceHealthOk() (*bool, bool)`

GetNotifyServiceHealthOk returns a tuple with the NotifyServiceHealth field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotifyServiceHealth

`func (o *NotificationConfigUpdateRequest) SetNotifyServiceHealth(v bool)`

SetNotifyServiceHealth sets NotifyServiceHealth field to given value.

### HasNotifyServiceHealth

`func (o *NotificationConfigUpdateRequest) HasNotifyServiceHealth() bool`

HasNotifyServiceHealth returns a boolean if a field has been set.

### SetNotifyServiceHealthNil

`func (o *NotificationConfigUpdateRequest) SetNotifyServiceHealthNil(b bool)`

 SetNotifyServiceHealthNil sets the value for NotifyServiceHealth to be an explicit nil

### UnsetNotifyServiceHealth
`func (o *NotificationConfigUpdateRequest) UnsetNotifyServiceHealth()`

UnsetNotifyServiceHealth ensures that no value is present for NotifyServiceHealth, not even an explicit nil
### GetDlqThreshold

`func (o *NotificationConfigUpdateRequest) GetDlqThreshold() int32`

GetDlqThreshold returns the DlqThreshold field if non-nil, zero value otherwise.

### GetDlqThresholdOk

`func (o *NotificationConfigUpdateRequest) GetDlqThresholdOk() (*int32, bool)`

GetDlqThresholdOk returns a tuple with the DlqThreshold field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDlqThreshold

`func (o *NotificationConfigUpdateRequest) SetDlqThreshold(v int32)`

SetDlqThreshold sets DlqThreshold field to given value.

### HasDlqThreshold

`func (o *NotificationConfigUpdateRequest) HasDlqThreshold() bool`

HasDlqThreshold returns a boolean if a field has been set.

### SetDlqThresholdNil

`func (o *NotificationConfigUpdateRequest) SetDlqThresholdNil(b bool)`

 SetDlqThresholdNil sets the value for DlqThreshold to be an explicit nil

### UnsetDlqThreshold
`func (o *NotificationConfigUpdateRequest) UnsetDlqThreshold()`

UnsetDlqThreshold ensures that no value is present for DlqThreshold, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


