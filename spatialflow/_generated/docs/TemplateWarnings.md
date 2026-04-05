# TemplateWarnings

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Warnings** | **[]string** |  | 
**Suggestions** | **map[string]string** |  | 

## Methods

### NewTemplateWarnings

`func NewTemplateWarnings(warnings []string, suggestions map[string]string, ) *TemplateWarnings`

NewTemplateWarnings instantiates a new TemplateWarnings object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTemplateWarningsWithDefaults

`func NewTemplateWarningsWithDefaults() *TemplateWarnings`

NewTemplateWarningsWithDefaults instantiates a new TemplateWarnings object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetWarnings

`func (o *TemplateWarnings) GetWarnings() []string`

GetWarnings returns the Warnings field if non-nil, zero value otherwise.

### GetWarningsOk

`func (o *TemplateWarnings) GetWarningsOk() (*[]string, bool)`

GetWarningsOk returns a tuple with the Warnings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWarnings

`func (o *TemplateWarnings) SetWarnings(v []string)`

SetWarnings sets Warnings field to given value.


### GetSuggestions

`func (o *TemplateWarnings) GetSuggestions() map[string]string`

GetSuggestions returns the Suggestions field if non-nil, zero value otherwise.

### GetSuggestionsOk

`func (o *TemplateWarnings) GetSuggestionsOk() (*map[string]string, bool)`

GetSuggestionsOk returns a tuple with the Suggestions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSuggestions

`func (o *TemplateWarnings) SetSuggestions(v map[string]string)`

SetSuggestions sets Suggestions field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


