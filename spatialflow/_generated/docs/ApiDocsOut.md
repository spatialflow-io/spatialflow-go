# ApiDocsOut

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Openapi** | **string** |  | 
**Info** | **map[string]interface{}** |  | 
**Servers** | **[]map[string]interface{}** |  | 

## Methods

### NewApiDocsOut

`func NewApiDocsOut(openapi string, info map[string]interface{}, servers []map[string]interface{}, ) *ApiDocsOut`

NewApiDocsOut instantiates a new ApiDocsOut object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewApiDocsOutWithDefaults

`func NewApiDocsOutWithDefaults() *ApiDocsOut`

NewApiDocsOutWithDefaults instantiates a new ApiDocsOut object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetOpenapi

`func (o *ApiDocsOut) GetOpenapi() string`

GetOpenapi returns the Openapi field if non-nil, zero value otherwise.

### GetOpenapiOk

`func (o *ApiDocsOut) GetOpenapiOk() (*string, bool)`

GetOpenapiOk returns a tuple with the Openapi field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOpenapi

`func (o *ApiDocsOut) SetOpenapi(v string)`

SetOpenapi sets Openapi field to given value.


### GetInfo

`func (o *ApiDocsOut) GetInfo() map[string]interface{}`

GetInfo returns the Info field if non-nil, zero value otherwise.

### GetInfoOk

`func (o *ApiDocsOut) GetInfoOk() (*map[string]interface{}, bool)`

GetInfoOk returns a tuple with the Info field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInfo

`func (o *ApiDocsOut) SetInfo(v map[string]interface{})`

SetInfo sets Info field to given value.


### GetServers

`func (o *ApiDocsOut) GetServers() []map[string]interface{}`

GetServers returns the Servers field if non-nil, zero value otherwise.

### GetServersOk

`func (o *ApiDocsOut) GetServersOk() (*[]map[string]interface{}, bool)`

GetServersOk returns a tuple with the Servers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServers

`func (o *ApiDocsOut) SetServers(v []map[string]interface{})`

SetServers sets Servers field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


