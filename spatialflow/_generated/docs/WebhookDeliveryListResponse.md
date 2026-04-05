# WebhookDeliveryListResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Deliveries** | [**[]WebhookDeliveryResponse**](WebhookDeliveryResponse.md) |  | 
**Pagination** | **map[string]interface{}** |  | 

## Methods

### NewWebhookDeliveryListResponse

`func NewWebhookDeliveryListResponse(deliveries []WebhookDeliveryResponse, pagination map[string]interface{}, ) *WebhookDeliveryListResponse`

NewWebhookDeliveryListResponse instantiates a new WebhookDeliveryListResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWebhookDeliveryListResponseWithDefaults

`func NewWebhookDeliveryListResponseWithDefaults() *WebhookDeliveryListResponse`

NewWebhookDeliveryListResponseWithDefaults instantiates a new WebhookDeliveryListResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDeliveries

`func (o *WebhookDeliveryListResponse) GetDeliveries() []WebhookDeliveryResponse`

GetDeliveries returns the Deliveries field if non-nil, zero value otherwise.

### GetDeliveriesOk

`func (o *WebhookDeliveryListResponse) GetDeliveriesOk() (*[]WebhookDeliveryResponse, bool)`

GetDeliveriesOk returns a tuple with the Deliveries field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeliveries

`func (o *WebhookDeliveryListResponse) SetDeliveries(v []WebhookDeliveryResponse)`

SetDeliveries sets Deliveries field to given value.


### GetPagination

`func (o *WebhookDeliveryListResponse) GetPagination() map[string]interface{}`

GetPagination returns the Pagination field if non-nil, zero value otherwise.

### GetPaginationOk

`func (o *WebhookDeliveryListResponse) GetPaginationOk() (*map[string]interface{}, bool)`

GetPaginationOk returns a tuple with the Pagination field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPagination

`func (o *WebhookDeliveryListResponse) SetPagination(v map[string]interface{})`

SetPagination sets Pagination field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


