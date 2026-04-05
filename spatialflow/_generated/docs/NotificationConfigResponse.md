# NotificationConfigResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Provider** | **string** |  | 
**ProviderName** | **string** |  | 
**WebhookUrlPlaceholder** | **string** |  | 
**IsEnabled** | **bool** |  | 
**WebhookUrlConfigured** | **bool** |  | 
**NotifyNewSignups** | **bool** |  | 
**NotifyAdminApprovals** | **bool** |  | 
**NotifySubscriptionChanges** | **bool** |  | 
**NotifyPaymentFailures** | **bool** |  | 
**NotifyPrivacyErasures** | **bool** |  | 
**NotifyDlqThreshold** | **bool** |  | 
**NotifyServiceHealth** | **bool** |  | 
**DlqThreshold** | **int32** |  | 
**UpdatedAt** | Pointer to **NullableString** |  | [optional] 
**UpdatedByEmail** | Pointer to **NullableString** |  | [optional] 

## Methods

### NewNotificationConfigResponse

`func NewNotificationConfigResponse(provider string, providerName string, webhookUrlPlaceholder string, isEnabled bool, webhookUrlConfigured bool, notifyNewSignups bool, notifyAdminApprovals bool, notifySubscriptionChanges bool, notifyPaymentFailures bool, notifyPrivacyErasures bool, notifyDlqThreshold bool, notifyServiceHealth bool, dlqThreshold int32, ) *NotificationConfigResponse`

NewNotificationConfigResponse instantiates a new NotificationConfigResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNotificationConfigResponseWithDefaults

`func NewNotificationConfigResponseWithDefaults() *NotificationConfigResponse`

NewNotificationConfigResponseWithDefaults instantiates a new NotificationConfigResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetProvider

`func (o *NotificationConfigResponse) GetProvider() string`

GetProvider returns the Provider field if non-nil, zero value otherwise.

### GetProviderOk

`func (o *NotificationConfigResponse) GetProviderOk() (*string, bool)`

GetProviderOk returns a tuple with the Provider field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProvider

`func (o *NotificationConfigResponse) SetProvider(v string)`

SetProvider sets Provider field to given value.


### GetProviderName

`func (o *NotificationConfigResponse) GetProviderName() string`

GetProviderName returns the ProviderName field if non-nil, zero value otherwise.

### GetProviderNameOk

`func (o *NotificationConfigResponse) GetProviderNameOk() (*string, bool)`

GetProviderNameOk returns a tuple with the ProviderName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProviderName

`func (o *NotificationConfigResponse) SetProviderName(v string)`

SetProviderName sets ProviderName field to given value.


### GetWebhookUrlPlaceholder

`func (o *NotificationConfigResponse) GetWebhookUrlPlaceholder() string`

GetWebhookUrlPlaceholder returns the WebhookUrlPlaceholder field if non-nil, zero value otherwise.

### GetWebhookUrlPlaceholderOk

`func (o *NotificationConfigResponse) GetWebhookUrlPlaceholderOk() (*string, bool)`

GetWebhookUrlPlaceholderOk returns a tuple with the WebhookUrlPlaceholder field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWebhookUrlPlaceholder

`func (o *NotificationConfigResponse) SetWebhookUrlPlaceholder(v string)`

SetWebhookUrlPlaceholder sets WebhookUrlPlaceholder field to given value.


### GetIsEnabled

`func (o *NotificationConfigResponse) GetIsEnabled() bool`

GetIsEnabled returns the IsEnabled field if non-nil, zero value otherwise.

### GetIsEnabledOk

`func (o *NotificationConfigResponse) GetIsEnabledOk() (*bool, bool)`

GetIsEnabledOk returns a tuple with the IsEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsEnabled

`func (o *NotificationConfigResponse) SetIsEnabled(v bool)`

SetIsEnabled sets IsEnabled field to given value.


### GetWebhookUrlConfigured

`func (o *NotificationConfigResponse) GetWebhookUrlConfigured() bool`

GetWebhookUrlConfigured returns the WebhookUrlConfigured field if non-nil, zero value otherwise.

### GetWebhookUrlConfiguredOk

`func (o *NotificationConfigResponse) GetWebhookUrlConfiguredOk() (*bool, bool)`

GetWebhookUrlConfiguredOk returns a tuple with the WebhookUrlConfigured field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWebhookUrlConfigured

`func (o *NotificationConfigResponse) SetWebhookUrlConfigured(v bool)`

SetWebhookUrlConfigured sets WebhookUrlConfigured field to given value.


### GetNotifyNewSignups

`func (o *NotificationConfigResponse) GetNotifyNewSignups() bool`

GetNotifyNewSignups returns the NotifyNewSignups field if non-nil, zero value otherwise.

### GetNotifyNewSignupsOk

`func (o *NotificationConfigResponse) GetNotifyNewSignupsOk() (*bool, bool)`

GetNotifyNewSignupsOk returns a tuple with the NotifyNewSignups field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotifyNewSignups

`func (o *NotificationConfigResponse) SetNotifyNewSignups(v bool)`

SetNotifyNewSignups sets NotifyNewSignups field to given value.


### GetNotifyAdminApprovals

`func (o *NotificationConfigResponse) GetNotifyAdminApprovals() bool`

GetNotifyAdminApprovals returns the NotifyAdminApprovals field if non-nil, zero value otherwise.

### GetNotifyAdminApprovalsOk

`func (o *NotificationConfigResponse) GetNotifyAdminApprovalsOk() (*bool, bool)`

GetNotifyAdminApprovalsOk returns a tuple with the NotifyAdminApprovals field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotifyAdminApprovals

`func (o *NotificationConfigResponse) SetNotifyAdminApprovals(v bool)`

SetNotifyAdminApprovals sets NotifyAdminApprovals field to given value.


### GetNotifySubscriptionChanges

`func (o *NotificationConfigResponse) GetNotifySubscriptionChanges() bool`

GetNotifySubscriptionChanges returns the NotifySubscriptionChanges field if non-nil, zero value otherwise.

### GetNotifySubscriptionChangesOk

`func (o *NotificationConfigResponse) GetNotifySubscriptionChangesOk() (*bool, bool)`

GetNotifySubscriptionChangesOk returns a tuple with the NotifySubscriptionChanges field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotifySubscriptionChanges

`func (o *NotificationConfigResponse) SetNotifySubscriptionChanges(v bool)`

SetNotifySubscriptionChanges sets NotifySubscriptionChanges field to given value.


### GetNotifyPaymentFailures

`func (o *NotificationConfigResponse) GetNotifyPaymentFailures() bool`

GetNotifyPaymentFailures returns the NotifyPaymentFailures field if non-nil, zero value otherwise.

### GetNotifyPaymentFailuresOk

`func (o *NotificationConfigResponse) GetNotifyPaymentFailuresOk() (*bool, bool)`

GetNotifyPaymentFailuresOk returns a tuple with the NotifyPaymentFailures field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotifyPaymentFailures

`func (o *NotificationConfigResponse) SetNotifyPaymentFailures(v bool)`

SetNotifyPaymentFailures sets NotifyPaymentFailures field to given value.


### GetNotifyPrivacyErasures

`func (o *NotificationConfigResponse) GetNotifyPrivacyErasures() bool`

GetNotifyPrivacyErasures returns the NotifyPrivacyErasures field if non-nil, zero value otherwise.

### GetNotifyPrivacyErasuresOk

`func (o *NotificationConfigResponse) GetNotifyPrivacyErasuresOk() (*bool, bool)`

GetNotifyPrivacyErasuresOk returns a tuple with the NotifyPrivacyErasures field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotifyPrivacyErasures

`func (o *NotificationConfigResponse) SetNotifyPrivacyErasures(v bool)`

SetNotifyPrivacyErasures sets NotifyPrivacyErasures field to given value.


### GetNotifyDlqThreshold

`func (o *NotificationConfigResponse) GetNotifyDlqThreshold() bool`

GetNotifyDlqThreshold returns the NotifyDlqThreshold field if non-nil, zero value otherwise.

### GetNotifyDlqThresholdOk

`func (o *NotificationConfigResponse) GetNotifyDlqThresholdOk() (*bool, bool)`

GetNotifyDlqThresholdOk returns a tuple with the NotifyDlqThreshold field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotifyDlqThreshold

`func (o *NotificationConfigResponse) SetNotifyDlqThreshold(v bool)`

SetNotifyDlqThreshold sets NotifyDlqThreshold field to given value.


### GetNotifyServiceHealth

`func (o *NotificationConfigResponse) GetNotifyServiceHealth() bool`

GetNotifyServiceHealth returns the NotifyServiceHealth field if non-nil, zero value otherwise.

### GetNotifyServiceHealthOk

`func (o *NotificationConfigResponse) GetNotifyServiceHealthOk() (*bool, bool)`

GetNotifyServiceHealthOk returns a tuple with the NotifyServiceHealth field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotifyServiceHealth

`func (o *NotificationConfigResponse) SetNotifyServiceHealth(v bool)`

SetNotifyServiceHealth sets NotifyServiceHealth field to given value.


### GetDlqThreshold

`func (o *NotificationConfigResponse) GetDlqThreshold() int32`

GetDlqThreshold returns the DlqThreshold field if non-nil, zero value otherwise.

### GetDlqThresholdOk

`func (o *NotificationConfigResponse) GetDlqThresholdOk() (*int32, bool)`

GetDlqThresholdOk returns a tuple with the DlqThreshold field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDlqThreshold

`func (o *NotificationConfigResponse) SetDlqThreshold(v int32)`

SetDlqThreshold sets DlqThreshold field to given value.


### GetUpdatedAt

`func (o *NotificationConfigResponse) GetUpdatedAt() string`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *NotificationConfigResponse) GetUpdatedAtOk() (*string, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *NotificationConfigResponse) SetUpdatedAt(v string)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *NotificationConfigResponse) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.

### SetUpdatedAtNil

`func (o *NotificationConfigResponse) SetUpdatedAtNil(b bool)`

 SetUpdatedAtNil sets the value for UpdatedAt to be an explicit nil

### UnsetUpdatedAt
`func (o *NotificationConfigResponse) UnsetUpdatedAt()`

UnsetUpdatedAt ensures that no value is present for UpdatedAt, not even an explicit nil
### GetUpdatedByEmail

`func (o *NotificationConfigResponse) GetUpdatedByEmail() string`

GetUpdatedByEmail returns the UpdatedByEmail field if non-nil, zero value otherwise.

### GetUpdatedByEmailOk

`func (o *NotificationConfigResponse) GetUpdatedByEmailOk() (*string, bool)`

GetUpdatedByEmailOk returns a tuple with the UpdatedByEmail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedByEmail

`func (o *NotificationConfigResponse) SetUpdatedByEmail(v string)`

SetUpdatedByEmail sets UpdatedByEmail field to given value.

### HasUpdatedByEmail

`func (o *NotificationConfigResponse) HasUpdatedByEmail() bool`

HasUpdatedByEmail returns a boolean if a field has been set.

### SetUpdatedByEmailNil

`func (o *NotificationConfigResponse) SetUpdatedByEmailNil(b bool)`

 SetUpdatedByEmailNil sets the value for UpdatedByEmail to be an explicit nil

### UnsetUpdatedByEmail
`func (o *NotificationConfigResponse) UnsetUpdatedByEmail()`

UnsetUpdatedByEmail ensures that no value is present for UpdatedByEmail, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


