# NotificationRouteListResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Routes** | [**[]NotificationRouteResponse**](NotificationRouteResponse.md) |  | 
**SupportedEventTypes** | **[]string** |  | 
**SlackWebhookChannelBound** | Pointer to **bool** |  | [optional] [default to true]

## Methods

### NewNotificationRouteListResponse

`func NewNotificationRouteListResponse(routes []NotificationRouteResponse, supportedEventTypes []string, ) *NotificationRouteListResponse`

NewNotificationRouteListResponse instantiates a new NotificationRouteListResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNotificationRouteListResponseWithDefaults

`func NewNotificationRouteListResponseWithDefaults() *NotificationRouteListResponse`

NewNotificationRouteListResponseWithDefaults instantiates a new NotificationRouteListResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRoutes

`func (o *NotificationRouteListResponse) GetRoutes() []NotificationRouteResponse`

GetRoutes returns the Routes field if non-nil, zero value otherwise.

### GetRoutesOk

`func (o *NotificationRouteListResponse) GetRoutesOk() (*[]NotificationRouteResponse, bool)`

GetRoutesOk returns a tuple with the Routes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRoutes

`func (o *NotificationRouteListResponse) SetRoutes(v []NotificationRouteResponse)`

SetRoutes sets Routes field to given value.


### GetSupportedEventTypes

`func (o *NotificationRouteListResponse) GetSupportedEventTypes() []string`

GetSupportedEventTypes returns the SupportedEventTypes field if non-nil, zero value otherwise.

### GetSupportedEventTypesOk

`func (o *NotificationRouteListResponse) GetSupportedEventTypesOk() (*[]string, bool)`

GetSupportedEventTypesOk returns a tuple with the SupportedEventTypes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSupportedEventTypes

`func (o *NotificationRouteListResponse) SetSupportedEventTypes(v []string)`

SetSupportedEventTypes sets SupportedEventTypes field to given value.


### GetSlackWebhookChannelBound

`func (o *NotificationRouteListResponse) GetSlackWebhookChannelBound() bool`

GetSlackWebhookChannelBound returns the SlackWebhookChannelBound field if non-nil, zero value otherwise.

### GetSlackWebhookChannelBoundOk

`func (o *NotificationRouteListResponse) GetSlackWebhookChannelBoundOk() (*bool, bool)`

GetSlackWebhookChannelBoundOk returns a tuple with the SlackWebhookChannelBound field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSlackWebhookChannelBound

`func (o *NotificationRouteListResponse) SetSlackWebhookChannelBound(v bool)`

SetSlackWebhookChannelBound sets SlackWebhookChannelBound field to given value.

### HasSlackWebhookChannelBound

`func (o *NotificationRouteListResponse) HasSlackWebhookChannelBound() bool`

HasSlackWebhookChannelBound returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


