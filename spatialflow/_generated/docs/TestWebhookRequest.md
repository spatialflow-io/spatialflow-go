# TestWebhookRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**EventType** | Pointer to **NullableString** |  | [optional] 
**CustomPayload** | Pointer to **map[string]interface{}** |  | [optional] 
**TestPayload** | Pointer to **map[string]interface{}** |  | [optional] 

## Methods

### NewTestWebhookRequest

`func NewTestWebhookRequest() *TestWebhookRequest`

NewTestWebhookRequest instantiates a new TestWebhookRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTestWebhookRequestWithDefaults

`func NewTestWebhookRequestWithDefaults() *TestWebhookRequest`

NewTestWebhookRequestWithDefaults instantiates a new TestWebhookRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEventType

`func (o *TestWebhookRequest) GetEventType() string`

GetEventType returns the EventType field if non-nil, zero value otherwise.

### GetEventTypeOk

`func (o *TestWebhookRequest) GetEventTypeOk() (*string, bool)`

GetEventTypeOk returns a tuple with the EventType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEventType

`func (o *TestWebhookRequest) SetEventType(v string)`

SetEventType sets EventType field to given value.

### HasEventType

`func (o *TestWebhookRequest) HasEventType() bool`

HasEventType returns a boolean if a field has been set.

### SetEventTypeNil

`func (o *TestWebhookRequest) SetEventTypeNil(b bool)`

 SetEventTypeNil sets the value for EventType to be an explicit nil

### UnsetEventType
`func (o *TestWebhookRequest) UnsetEventType()`

UnsetEventType ensures that no value is present for EventType, not even an explicit nil
### GetCustomPayload

`func (o *TestWebhookRequest) GetCustomPayload() map[string]interface{}`

GetCustomPayload returns the CustomPayload field if non-nil, zero value otherwise.

### GetCustomPayloadOk

`func (o *TestWebhookRequest) GetCustomPayloadOk() (*map[string]interface{}, bool)`

GetCustomPayloadOk returns a tuple with the CustomPayload field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomPayload

`func (o *TestWebhookRequest) SetCustomPayload(v map[string]interface{})`

SetCustomPayload sets CustomPayload field to given value.

### HasCustomPayload

`func (o *TestWebhookRequest) HasCustomPayload() bool`

HasCustomPayload returns a boolean if a field has been set.

### SetCustomPayloadNil

`func (o *TestWebhookRequest) SetCustomPayloadNil(b bool)`

 SetCustomPayloadNil sets the value for CustomPayload to be an explicit nil

### UnsetCustomPayload
`func (o *TestWebhookRequest) UnsetCustomPayload()`

UnsetCustomPayload ensures that no value is present for CustomPayload, not even an explicit nil
### GetTestPayload

`func (o *TestWebhookRequest) GetTestPayload() map[string]interface{}`

GetTestPayload returns the TestPayload field if non-nil, zero value otherwise.

### GetTestPayloadOk

`func (o *TestWebhookRequest) GetTestPayloadOk() (*map[string]interface{}, bool)`

GetTestPayloadOk returns a tuple with the TestPayload field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTestPayload

`func (o *TestWebhookRequest) SetTestPayload(v map[string]interface{})`

SetTestPayload sets TestPayload field to given value.

### HasTestPayload

`func (o *TestWebhookRequest) HasTestPayload() bool`

HasTestPayload returns a boolean if a field has been set.

### SetTestPayloadNil

`func (o *TestWebhookRequest) SetTestPayloadNil(b bool)`

 SetTestPayloadNil sets the value for TestPayload to be an explicit nil

### UnsetTestPayload
`func (o *TestWebhookRequest) UnsetTestPayload()`

UnsetTestPayload ensures that no value is present for TestPayload, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


