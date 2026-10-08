# NotificationRoutePatchRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **NullableString** |  | [optional] 
**Provider** | Pointer to **NullableString** |  | [optional] 
**DestinationLabel** | Pointer to **NullableString** |  | [optional] 
**WebhookUrl** | Pointer to **NullableString** |  | [optional] 
**EventTypes** | Pointer to **[]string** |  | [optional] 
**IsEnabled** | Pointer to **NullableBool** |  | [optional] 
**IsDefault** | Pointer to **NullableBool** |  | [optional] 

## Methods

### NewNotificationRoutePatchRequest

`func NewNotificationRoutePatchRequest() *NotificationRoutePatchRequest`

NewNotificationRoutePatchRequest instantiates a new NotificationRoutePatchRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNotificationRoutePatchRequestWithDefaults

`func NewNotificationRoutePatchRequestWithDefaults() *NotificationRoutePatchRequest`

NewNotificationRoutePatchRequestWithDefaults instantiates a new NotificationRoutePatchRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *NotificationRoutePatchRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *NotificationRoutePatchRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *NotificationRoutePatchRequest) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *NotificationRoutePatchRequest) HasName() bool`

HasName returns a boolean if a field has been set.

### SetNameNil

`func (o *NotificationRoutePatchRequest) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *NotificationRoutePatchRequest) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetProvider

`func (o *NotificationRoutePatchRequest) GetProvider() string`

GetProvider returns the Provider field if non-nil, zero value otherwise.

### GetProviderOk

`func (o *NotificationRoutePatchRequest) GetProviderOk() (*string, bool)`

GetProviderOk returns a tuple with the Provider field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProvider

`func (o *NotificationRoutePatchRequest) SetProvider(v string)`

SetProvider sets Provider field to given value.

### HasProvider

`func (o *NotificationRoutePatchRequest) HasProvider() bool`

HasProvider returns a boolean if a field has been set.

### SetProviderNil

`func (o *NotificationRoutePatchRequest) SetProviderNil(b bool)`

 SetProviderNil sets the value for Provider to be an explicit nil

### UnsetProvider
`func (o *NotificationRoutePatchRequest) UnsetProvider()`

UnsetProvider ensures that no value is present for Provider, not even an explicit nil
### GetDestinationLabel

`func (o *NotificationRoutePatchRequest) GetDestinationLabel() string`

GetDestinationLabel returns the DestinationLabel field if non-nil, zero value otherwise.

### GetDestinationLabelOk

`func (o *NotificationRoutePatchRequest) GetDestinationLabelOk() (*string, bool)`

GetDestinationLabelOk returns a tuple with the DestinationLabel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDestinationLabel

`func (o *NotificationRoutePatchRequest) SetDestinationLabel(v string)`

SetDestinationLabel sets DestinationLabel field to given value.

### HasDestinationLabel

`func (o *NotificationRoutePatchRequest) HasDestinationLabel() bool`

HasDestinationLabel returns a boolean if a field has been set.

### SetDestinationLabelNil

`func (o *NotificationRoutePatchRequest) SetDestinationLabelNil(b bool)`

 SetDestinationLabelNil sets the value for DestinationLabel to be an explicit nil

### UnsetDestinationLabel
`func (o *NotificationRoutePatchRequest) UnsetDestinationLabel()`

UnsetDestinationLabel ensures that no value is present for DestinationLabel, not even an explicit nil
### GetWebhookUrl

`func (o *NotificationRoutePatchRequest) GetWebhookUrl() string`

GetWebhookUrl returns the WebhookUrl field if non-nil, zero value otherwise.

### GetWebhookUrlOk

`func (o *NotificationRoutePatchRequest) GetWebhookUrlOk() (*string, bool)`

GetWebhookUrlOk returns a tuple with the WebhookUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWebhookUrl

`func (o *NotificationRoutePatchRequest) SetWebhookUrl(v string)`

SetWebhookUrl sets WebhookUrl field to given value.

### HasWebhookUrl

`func (o *NotificationRoutePatchRequest) HasWebhookUrl() bool`

HasWebhookUrl returns a boolean if a field has been set.

### SetWebhookUrlNil

`func (o *NotificationRoutePatchRequest) SetWebhookUrlNil(b bool)`

 SetWebhookUrlNil sets the value for WebhookUrl to be an explicit nil

### UnsetWebhookUrl
`func (o *NotificationRoutePatchRequest) UnsetWebhookUrl()`

UnsetWebhookUrl ensures that no value is present for WebhookUrl, not even an explicit nil
### GetEventTypes

`func (o *NotificationRoutePatchRequest) GetEventTypes() []string`

GetEventTypes returns the EventTypes field if non-nil, zero value otherwise.

### GetEventTypesOk

`func (o *NotificationRoutePatchRequest) GetEventTypesOk() (*[]string, bool)`

GetEventTypesOk returns a tuple with the EventTypes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEventTypes

`func (o *NotificationRoutePatchRequest) SetEventTypes(v []string)`

SetEventTypes sets EventTypes field to given value.

### HasEventTypes

`func (o *NotificationRoutePatchRequest) HasEventTypes() bool`

HasEventTypes returns a boolean if a field has been set.

### SetEventTypesNil

`func (o *NotificationRoutePatchRequest) SetEventTypesNil(b bool)`

 SetEventTypesNil sets the value for EventTypes to be an explicit nil

### UnsetEventTypes
`func (o *NotificationRoutePatchRequest) UnsetEventTypes()`

UnsetEventTypes ensures that no value is present for EventTypes, not even an explicit nil
### GetIsEnabled

`func (o *NotificationRoutePatchRequest) GetIsEnabled() bool`

GetIsEnabled returns the IsEnabled field if non-nil, zero value otherwise.

### GetIsEnabledOk

`func (o *NotificationRoutePatchRequest) GetIsEnabledOk() (*bool, bool)`

GetIsEnabledOk returns a tuple with the IsEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsEnabled

`func (o *NotificationRoutePatchRequest) SetIsEnabled(v bool)`

SetIsEnabled sets IsEnabled field to given value.

### HasIsEnabled

`func (o *NotificationRoutePatchRequest) HasIsEnabled() bool`

HasIsEnabled returns a boolean if a field has been set.

### SetIsEnabledNil

`func (o *NotificationRoutePatchRequest) SetIsEnabledNil(b bool)`

 SetIsEnabledNil sets the value for IsEnabled to be an explicit nil

### UnsetIsEnabled
`func (o *NotificationRoutePatchRequest) UnsetIsEnabled()`

UnsetIsEnabled ensures that no value is present for IsEnabled, not even an explicit nil
### GetIsDefault

`func (o *NotificationRoutePatchRequest) GetIsDefault() bool`

GetIsDefault returns the IsDefault field if non-nil, zero value otherwise.

### GetIsDefaultOk

`func (o *NotificationRoutePatchRequest) GetIsDefaultOk() (*bool, bool)`

GetIsDefaultOk returns a tuple with the IsDefault field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsDefault

`func (o *NotificationRoutePatchRequest) SetIsDefault(v bool)`

SetIsDefault sets IsDefault field to given value.

### HasIsDefault

`func (o *NotificationRoutePatchRequest) HasIsDefault() bool`

HasIsDefault returns a boolean if a field has been set.

### SetIsDefaultNil

`func (o *NotificationRoutePatchRequest) SetIsDefaultNil(b bool)`

 SetIsDefaultNil sets the value for IsDefault to be an explicit nil

### UnsetIsDefault
`func (o *NotificationRoutePatchRequest) UnsetIsDefault()`

UnsetIsDefault ensures that no value is present for IsDefault, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


