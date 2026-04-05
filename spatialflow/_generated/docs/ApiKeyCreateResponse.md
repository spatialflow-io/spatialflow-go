# ApiKeyCreateResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Message** | **string** |  | 
**ApiKey** | [**ApiKeyDetailSchema**](ApiKeyDetailSchema.md) |  | 

## Methods

### NewApiKeyCreateResponse

`func NewApiKeyCreateResponse(message string, apiKey ApiKeyDetailSchema, ) *ApiKeyCreateResponse`

NewApiKeyCreateResponse instantiates a new ApiKeyCreateResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewApiKeyCreateResponseWithDefaults

`func NewApiKeyCreateResponseWithDefaults() *ApiKeyCreateResponse`

NewApiKeyCreateResponseWithDefaults instantiates a new ApiKeyCreateResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMessage

`func (o *ApiKeyCreateResponse) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *ApiKeyCreateResponse) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *ApiKeyCreateResponse) SetMessage(v string)`

SetMessage sets Message field to given value.


### GetApiKey

`func (o *ApiKeyCreateResponse) GetApiKey() ApiKeyDetailSchema`

GetApiKey returns the ApiKey field if non-nil, zero value otherwise.

### GetApiKeyOk

`func (o *ApiKeyCreateResponse) GetApiKeyOk() (*ApiKeyDetailSchema, bool)`

GetApiKeyOk returns a tuple with the ApiKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApiKey

`func (o *ApiKeyCreateResponse) SetApiKey(v ApiKeyDetailSchema)`

SetApiKey sets ApiKey field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


