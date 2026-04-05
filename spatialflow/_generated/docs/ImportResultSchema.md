# ImportResultSchema

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Success** | **bool** |  | 
**IntegrationId** | Pointer to **NullableString** |  | [optional] 
**Message** | **string** |  | 
**Conflicts** | Pointer to **[]string** |  | [optional] 

## Methods

### NewImportResultSchema

`func NewImportResultSchema(success bool, message string, ) *ImportResultSchema`

NewImportResultSchema instantiates a new ImportResultSchema object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewImportResultSchemaWithDefaults

`func NewImportResultSchemaWithDefaults() *ImportResultSchema`

NewImportResultSchemaWithDefaults instantiates a new ImportResultSchema object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSuccess

`func (o *ImportResultSchema) GetSuccess() bool`

GetSuccess returns the Success field if non-nil, zero value otherwise.

### GetSuccessOk

`func (o *ImportResultSchema) GetSuccessOk() (*bool, bool)`

GetSuccessOk returns a tuple with the Success field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSuccess

`func (o *ImportResultSchema) SetSuccess(v bool)`

SetSuccess sets Success field to given value.


### GetIntegrationId

`func (o *ImportResultSchema) GetIntegrationId() string`

GetIntegrationId returns the IntegrationId field if non-nil, zero value otherwise.

### GetIntegrationIdOk

`func (o *ImportResultSchema) GetIntegrationIdOk() (*string, bool)`

GetIntegrationIdOk returns a tuple with the IntegrationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIntegrationId

`func (o *ImportResultSchema) SetIntegrationId(v string)`

SetIntegrationId sets IntegrationId field to given value.

### HasIntegrationId

`func (o *ImportResultSchema) HasIntegrationId() bool`

HasIntegrationId returns a boolean if a field has been set.

### SetIntegrationIdNil

`func (o *ImportResultSchema) SetIntegrationIdNil(b bool)`

 SetIntegrationIdNil sets the value for IntegrationId to be an explicit nil

### UnsetIntegrationId
`func (o *ImportResultSchema) UnsetIntegrationId()`

UnsetIntegrationId ensures that no value is present for IntegrationId, not even an explicit nil
### GetMessage

`func (o *ImportResultSchema) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *ImportResultSchema) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *ImportResultSchema) SetMessage(v string)`

SetMessage sets Message field to given value.


### GetConflicts

`func (o *ImportResultSchema) GetConflicts() []string`

GetConflicts returns the Conflicts field if non-nil, zero value otherwise.

### GetConflictsOk

`func (o *ImportResultSchema) GetConflictsOk() (*[]string, bool)`

GetConflictsOk returns a tuple with the Conflicts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConflicts

`func (o *ImportResultSchema) SetConflicts(v []string)`

SetConflicts sets Conflicts field to given value.

### HasConflicts

`func (o *ImportResultSchema) HasConflicts() bool`

HasConflicts returns a boolean if a field has been set.

### SetConflictsNil

`func (o *ImportResultSchema) SetConflictsNil(b bool)`

 SetConflictsNil sets the value for Conflicts to be an explicit nil

### UnsetConflicts
`func (o *ImportResultSchema) UnsetConflicts()`

UnsetConflicts ensures that no value is present for Conflicts, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


