# WebhookListResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Webhooks** | [**[]WebhookResponse**](WebhookResponse.md) |  | 
**Pagination** | **map[string]interface{}** |  | 

## Methods

### NewWebhookListResponse

`func NewWebhookListResponse(webhooks []WebhookResponse, pagination map[string]interface{}, ) *WebhookListResponse`

NewWebhookListResponse instantiates a new WebhookListResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWebhookListResponseWithDefaults

`func NewWebhookListResponseWithDefaults() *WebhookListResponse`

NewWebhookListResponseWithDefaults instantiates a new WebhookListResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetWebhooks

`func (o *WebhookListResponse) GetWebhooks() []WebhookResponse`

GetWebhooks returns the Webhooks field if non-nil, zero value otherwise.

### GetWebhooksOk

`func (o *WebhookListResponse) GetWebhooksOk() (*[]WebhookResponse, bool)`

GetWebhooksOk returns a tuple with the Webhooks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWebhooks

`func (o *WebhookListResponse) SetWebhooks(v []WebhookResponse)`

SetWebhooks sets Webhooks field to given value.


### GetPagination

`func (o *WebhookListResponse) GetPagination() map[string]interface{}`

GetPagination returns the Pagination field if non-nil, zero value otherwise.

### GetPaginationOk

`func (o *WebhookListResponse) GetPaginationOk() (*map[string]interface{}, bool)`

GetPaginationOk returns a tuple with the Pagination field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPagination

`func (o *WebhookListResponse) SetPagination(v map[string]interface{})`

SetPagination sets Pagination field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


