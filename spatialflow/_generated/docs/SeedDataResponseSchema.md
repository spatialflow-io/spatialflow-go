# SeedDataResponseSchema

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Success** | **bool** |  | 
**Message** | **string** |  | 
**Details** | Pointer to **map[string]interface{}** |  | [optional] [default to {}]

## Methods

### NewSeedDataResponseSchema

`func NewSeedDataResponseSchema(success bool, message string, ) *SeedDataResponseSchema`

NewSeedDataResponseSchema instantiates a new SeedDataResponseSchema object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSeedDataResponseSchemaWithDefaults

`func NewSeedDataResponseSchemaWithDefaults() *SeedDataResponseSchema`

NewSeedDataResponseSchemaWithDefaults instantiates a new SeedDataResponseSchema object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSuccess

`func (o *SeedDataResponseSchema) GetSuccess() bool`

GetSuccess returns the Success field if non-nil, zero value otherwise.

### GetSuccessOk

`func (o *SeedDataResponseSchema) GetSuccessOk() (*bool, bool)`

GetSuccessOk returns a tuple with the Success field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSuccess

`func (o *SeedDataResponseSchema) SetSuccess(v bool)`

SetSuccess sets Success field to given value.


### GetMessage

`func (o *SeedDataResponseSchema) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *SeedDataResponseSchema) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *SeedDataResponseSchema) SetMessage(v string)`

SetMessage sets Message field to given value.


### GetDetails

`func (o *SeedDataResponseSchema) GetDetails() map[string]interface{}`

GetDetails returns the Details field if non-nil, zero value otherwise.

### GetDetailsOk

`func (o *SeedDataResponseSchema) GetDetailsOk() (*map[string]interface{}, bool)`

GetDetailsOk returns a tuple with the Details field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDetails

`func (o *SeedDataResponseSchema) SetDetails(v map[string]interface{})`

SetDetails sets Details field to given value.

### HasDetails

`func (o *SeedDataResponseSchema) HasDetails() bool`

HasDetails returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


