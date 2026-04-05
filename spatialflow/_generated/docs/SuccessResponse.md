# SuccessResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Message** | **string** |  | 
**Data** | Pointer to **map[string]interface{}** |  | [optional] 

## Methods

### NewSuccessResponse

`func NewSuccessResponse(message string, ) *SuccessResponse`

NewSuccessResponse instantiates a new SuccessResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSuccessResponseWithDefaults

`func NewSuccessResponseWithDefaults() *SuccessResponse`

NewSuccessResponseWithDefaults instantiates a new SuccessResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMessage

`func (o *SuccessResponse) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *SuccessResponse) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *SuccessResponse) SetMessage(v string)`

SetMessage sets Message field to given value.


### GetData

`func (o *SuccessResponse) GetData() map[string]interface{}`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *SuccessResponse) GetDataOk() (*map[string]interface{}, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *SuccessResponse) SetData(v map[string]interface{})`

SetData sets Data field to given value.

### HasData

`func (o *SuccessResponse) HasData() bool`

HasData returns a boolean if a field has been set.

### SetDataNil

`func (o *SuccessResponse) SetDataNil(b bool)`

 SetDataNil sets the value for Data to be an explicit nil

### UnsetData
`func (o *SuccessResponse) UnsetData()`

UnsetData ensures that no value is present for Data, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


