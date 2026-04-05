# CreateFromTemplateIn

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** |  | 
**Customizations** | Pointer to **map[string]interface{}** |  | [optional] 

## Methods

### NewCreateFromTemplateIn

`func NewCreateFromTemplateIn(name string, ) *CreateFromTemplateIn`

NewCreateFromTemplateIn instantiates a new CreateFromTemplateIn object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateFromTemplateInWithDefaults

`func NewCreateFromTemplateInWithDefaults() *CreateFromTemplateIn`

NewCreateFromTemplateInWithDefaults instantiates a new CreateFromTemplateIn object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *CreateFromTemplateIn) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CreateFromTemplateIn) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CreateFromTemplateIn) SetName(v string)`

SetName sets Name field to given value.


### GetCustomizations

`func (o *CreateFromTemplateIn) GetCustomizations() map[string]interface{}`

GetCustomizations returns the Customizations field if non-nil, zero value otherwise.

### GetCustomizationsOk

`func (o *CreateFromTemplateIn) GetCustomizationsOk() (*map[string]interface{}, bool)`

GetCustomizationsOk returns a tuple with the Customizations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomizations

`func (o *CreateFromTemplateIn) SetCustomizations(v map[string]interface{})`

SetCustomizations sets Customizations field to given value.

### HasCustomizations

`func (o *CreateFromTemplateIn) HasCustomizations() bool`

HasCustomizations returns a boolean if a field has been set.

### SetCustomizationsNil

`func (o *CreateFromTemplateIn) SetCustomizationsNil(b bool)`

 SetCustomizationsNil sets the value for Customizations to be an explicit nil

### UnsetCustomizations
`func (o *CreateFromTemplateIn) UnsetCustomizations()`

UnsetCustomizations ensures that no value is present for Customizations, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


