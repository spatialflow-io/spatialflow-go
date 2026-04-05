# ImportIntegrationSchema

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Data** | [**ExportIntegrationSchema**](ExportIntegrationSchema.md) |  | 
**Overwrite** | Pointer to **bool** | Overwrite existing integration with same name | [optional] [default to false]
**DecryptKey** | Pointer to **NullableString** |  | [optional] 

## Methods

### NewImportIntegrationSchema

`func NewImportIntegrationSchema(data ExportIntegrationSchema, ) *ImportIntegrationSchema`

NewImportIntegrationSchema instantiates a new ImportIntegrationSchema object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewImportIntegrationSchemaWithDefaults

`func NewImportIntegrationSchemaWithDefaults() *ImportIntegrationSchema`

NewImportIntegrationSchemaWithDefaults instantiates a new ImportIntegrationSchema object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetData

`func (o *ImportIntegrationSchema) GetData() ExportIntegrationSchema`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *ImportIntegrationSchema) GetDataOk() (*ExportIntegrationSchema, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *ImportIntegrationSchema) SetData(v ExportIntegrationSchema)`

SetData sets Data field to given value.


### GetOverwrite

`func (o *ImportIntegrationSchema) GetOverwrite() bool`

GetOverwrite returns the Overwrite field if non-nil, zero value otherwise.

### GetOverwriteOk

`func (o *ImportIntegrationSchema) GetOverwriteOk() (*bool, bool)`

GetOverwriteOk returns a tuple with the Overwrite field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOverwrite

`func (o *ImportIntegrationSchema) SetOverwrite(v bool)`

SetOverwrite sets Overwrite field to given value.

### HasOverwrite

`func (o *ImportIntegrationSchema) HasOverwrite() bool`

HasOverwrite returns a boolean if a field has been set.

### GetDecryptKey

`func (o *ImportIntegrationSchema) GetDecryptKey() string`

GetDecryptKey returns the DecryptKey field if non-nil, zero value otherwise.

### GetDecryptKeyOk

`func (o *ImportIntegrationSchema) GetDecryptKeyOk() (*string, bool)`

GetDecryptKeyOk returns a tuple with the DecryptKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDecryptKey

`func (o *ImportIntegrationSchema) SetDecryptKey(v string)`

SetDecryptKey sets DecryptKey field to given value.

### HasDecryptKey

`func (o *ImportIntegrationSchema) HasDecryptKey() bool`

HasDecryptKey returns a boolean if a field has been set.

### SetDecryptKeyNil

`func (o *ImportIntegrationSchema) SetDecryptKeyNil(b bool)`

 SetDecryptKeyNil sets the value for DecryptKey to be an explicit nil

### UnsetDecryptKey
`func (o *ImportIntegrationSchema) UnsetDecryptKey()`

UnsetDecryptKey ensures that no value is present for DecryptKey, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


