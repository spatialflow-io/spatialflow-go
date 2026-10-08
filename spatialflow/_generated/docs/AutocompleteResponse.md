# AutocompleteResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Suggestions** | [**[]SuggestionSchema**](SuggestionSchema.md) |  | 

## Methods

### NewAutocompleteResponse

`func NewAutocompleteResponse(suggestions []SuggestionSchema, ) *AutocompleteResponse`

NewAutocompleteResponse instantiates a new AutocompleteResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAutocompleteResponseWithDefaults

`func NewAutocompleteResponseWithDefaults() *AutocompleteResponse`

NewAutocompleteResponseWithDefaults instantiates a new AutocompleteResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSuggestions

`func (o *AutocompleteResponse) GetSuggestions() []SuggestionSchema`

GetSuggestions returns the Suggestions field if non-nil, zero value otherwise.

### GetSuggestionsOk

`func (o *AutocompleteResponse) GetSuggestionsOk() (*[]SuggestionSchema, bool)`

GetSuggestionsOk returns a tuple with the Suggestions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSuggestions

`func (o *AutocompleteResponse) SetSuggestions(v []SuggestionSchema)`

SetSuggestions sets Suggestions field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


