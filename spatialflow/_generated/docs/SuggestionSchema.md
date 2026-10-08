# SuggestionSchema

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Text** | **string** |  | 
**PlaceId** | Pointer to **NullableString** |  | [optional] 

## Methods

### NewSuggestionSchema

`func NewSuggestionSchema(text string, ) *SuggestionSchema`

NewSuggestionSchema instantiates a new SuggestionSchema object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSuggestionSchemaWithDefaults

`func NewSuggestionSchemaWithDefaults() *SuggestionSchema`

NewSuggestionSchemaWithDefaults instantiates a new SuggestionSchema object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetText

`func (o *SuggestionSchema) GetText() string`

GetText returns the Text field if non-nil, zero value otherwise.

### GetTextOk

`func (o *SuggestionSchema) GetTextOk() (*string, bool)`

GetTextOk returns a tuple with the Text field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetText

`func (o *SuggestionSchema) SetText(v string)`

SetText sets Text field to given value.


### GetPlaceId

`func (o *SuggestionSchema) GetPlaceId() string`

GetPlaceId returns the PlaceId field if non-nil, zero value otherwise.

### GetPlaceIdOk

`func (o *SuggestionSchema) GetPlaceIdOk() (*string, bool)`

GetPlaceIdOk returns a tuple with the PlaceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlaceId

`func (o *SuggestionSchema) SetPlaceId(v string)`

SetPlaceId sets PlaceId field to given value.

### HasPlaceId

`func (o *SuggestionSchema) HasPlaceId() bool`

HasPlaceId returns a boolean if a field has been set.

### SetPlaceIdNil

`func (o *SuggestionSchema) SetPlaceIdNil(b bool)`

 SetPlaceIdNil sets the value for PlaceId to be an explicit nil

### UnsetPlaceId
`func (o *SuggestionSchema) UnsetPlaceId()`

UnsetPlaceId ensures that no value is present for PlaceId, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


