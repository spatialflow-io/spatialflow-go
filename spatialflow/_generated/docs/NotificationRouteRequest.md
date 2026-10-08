# NotificationRouteRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** |  | 
**Provider** | Pointer to **string** |  | [optional] [default to "slack"]
**DestinationLabel** | Pointer to **NullableString** |  | [optional] 
**WebhookUrl** | Pointer to **NullableString** |  | [optional] 
**EventTypes** | Pointer to **[]string** |  | [optional] [default to []]
**IsEnabled** | Pointer to **bool** |  | [optional] [default to true]
**IsDefault** | Pointer to **bool** |  | [optional] [default to false]

## Methods

### NewNotificationRouteRequest

`func NewNotificationRouteRequest(name string, ) *NotificationRouteRequest`

NewNotificationRouteRequest instantiates a new NotificationRouteRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNotificationRouteRequestWithDefaults

`func NewNotificationRouteRequestWithDefaults() *NotificationRouteRequest`

NewNotificationRouteRequestWithDefaults instantiates a new NotificationRouteRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *NotificationRouteRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *NotificationRouteRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *NotificationRouteRequest) SetName(v string)`

SetName sets Name field to given value.


### GetProvider

`func (o *NotificationRouteRequest) GetProvider() string`

GetProvider returns the Provider field if non-nil, zero value otherwise.

### GetProviderOk

`func (o *NotificationRouteRequest) GetProviderOk() (*string, bool)`

GetProviderOk returns a tuple with the Provider field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProvider

`func (o *NotificationRouteRequest) SetProvider(v string)`

SetProvider sets Provider field to given value.

### HasProvider

`func (o *NotificationRouteRequest) HasProvider() bool`

HasProvider returns a boolean if a field has been set.

### GetDestinationLabel

`func (o *NotificationRouteRequest) GetDestinationLabel() string`

GetDestinationLabel returns the DestinationLabel field if non-nil, zero value otherwise.

### GetDestinationLabelOk

`func (o *NotificationRouteRequest) GetDestinationLabelOk() (*string, bool)`

GetDestinationLabelOk returns a tuple with the DestinationLabel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDestinationLabel

`func (o *NotificationRouteRequest) SetDestinationLabel(v string)`

SetDestinationLabel sets DestinationLabel field to given value.

### HasDestinationLabel

`func (o *NotificationRouteRequest) HasDestinationLabel() bool`

HasDestinationLabel returns a boolean if a field has been set.

### SetDestinationLabelNil

`func (o *NotificationRouteRequest) SetDestinationLabelNil(b bool)`

 SetDestinationLabelNil sets the value for DestinationLabel to be an explicit nil

### UnsetDestinationLabel
`func (o *NotificationRouteRequest) UnsetDestinationLabel()`

UnsetDestinationLabel ensures that no value is present for DestinationLabel, not even an explicit nil
### GetWebhookUrl

`func (o *NotificationRouteRequest) GetWebhookUrl() string`

GetWebhookUrl returns the WebhookUrl field if non-nil, zero value otherwise.

### GetWebhookUrlOk

`func (o *NotificationRouteRequest) GetWebhookUrlOk() (*string, bool)`

GetWebhookUrlOk returns a tuple with the WebhookUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWebhookUrl

`func (o *NotificationRouteRequest) SetWebhookUrl(v string)`

SetWebhookUrl sets WebhookUrl field to given value.

### HasWebhookUrl

`func (o *NotificationRouteRequest) HasWebhookUrl() bool`

HasWebhookUrl returns a boolean if a field has been set.

### SetWebhookUrlNil

`func (o *NotificationRouteRequest) SetWebhookUrlNil(b bool)`

 SetWebhookUrlNil sets the value for WebhookUrl to be an explicit nil

### UnsetWebhookUrl
`func (o *NotificationRouteRequest) UnsetWebhookUrl()`

UnsetWebhookUrl ensures that no value is present for WebhookUrl, not even an explicit nil
### GetEventTypes

`func (o *NotificationRouteRequest) GetEventTypes() []string`

GetEventTypes returns the EventTypes field if non-nil, zero value otherwise.

### GetEventTypesOk

`func (o *NotificationRouteRequest) GetEventTypesOk() (*[]string, bool)`

GetEventTypesOk returns a tuple with the EventTypes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEventTypes

`func (o *NotificationRouteRequest) SetEventTypes(v []string)`

SetEventTypes sets EventTypes field to given value.

### HasEventTypes

`func (o *NotificationRouteRequest) HasEventTypes() bool`

HasEventTypes returns a boolean if a field has been set.

### GetIsEnabled

`func (o *NotificationRouteRequest) GetIsEnabled() bool`

GetIsEnabled returns the IsEnabled field if non-nil, zero value otherwise.

### GetIsEnabledOk

`func (o *NotificationRouteRequest) GetIsEnabledOk() (*bool, bool)`

GetIsEnabledOk returns a tuple with the IsEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsEnabled

`func (o *NotificationRouteRequest) SetIsEnabled(v bool)`

SetIsEnabled sets IsEnabled field to given value.

### HasIsEnabled

`func (o *NotificationRouteRequest) HasIsEnabled() bool`

HasIsEnabled returns a boolean if a field has been set.

### GetIsDefault

`func (o *NotificationRouteRequest) GetIsDefault() bool`

GetIsDefault returns the IsDefault field if non-nil, zero value otherwise.

### GetIsDefaultOk

`func (o *NotificationRouteRequest) GetIsDefaultOk() (*bool, bool)`

GetIsDefaultOk returns a tuple with the IsDefault field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsDefault

`func (o *NotificationRouteRequest) SetIsDefault(v bool)`

SetIsDefault sets IsDefault field to given value.

### HasIsDefault

`func (o *NotificationRouteRequest) HasIsDefault() bool`

HasIsDefault returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


