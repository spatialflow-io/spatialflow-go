# TagSchema

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** |  | 
**Name** | **string** |  | 
**NormalizedName** | **string** |  | 
**UsageCount** | Pointer to **int32** | Count of geofences in the workspace carrying this tag. | [optional] [default to 0]

## Methods

### NewTagSchema

`func NewTagSchema(id string, name string, normalizedName string, ) *TagSchema`

NewTagSchema instantiates a new TagSchema object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTagSchemaWithDefaults

`func NewTagSchemaWithDefaults() *TagSchema`

NewTagSchemaWithDefaults instantiates a new TagSchema object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *TagSchema) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *TagSchema) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *TagSchema) SetId(v string)`

SetId sets Id field to given value.


### GetName

`func (o *TagSchema) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *TagSchema) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *TagSchema) SetName(v string)`

SetName sets Name field to given value.


### GetNormalizedName

`func (o *TagSchema) GetNormalizedName() string`

GetNormalizedName returns the NormalizedName field if non-nil, zero value otherwise.

### GetNormalizedNameOk

`func (o *TagSchema) GetNormalizedNameOk() (*string, bool)`

GetNormalizedNameOk returns a tuple with the NormalizedName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNormalizedName

`func (o *TagSchema) SetNormalizedName(v string)`

SetNormalizedName sets NormalizedName field to given value.


### GetUsageCount

`func (o *TagSchema) GetUsageCount() int32`

GetUsageCount returns the UsageCount field if non-nil, zero value otherwise.

### GetUsageCountOk

`func (o *TagSchema) GetUsageCountOk() (*int32, bool)`

GetUsageCountOk returns a tuple with the UsageCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsageCount

`func (o *TagSchema) SetUsageCount(v int32)`

SetUsageCount sets UsageCount field to given value.

### HasUsageCount

`func (o *TagSchema) HasUsageCount() bool`

HasUsageCount returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


