# ConfigurationListResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Categories** | [**[]CategoryInfo**](CategoryInfo.md) |  | 
**Configurations** | [**map[string][]ConfigurationItem**](array.md) |  | 

## Methods

### NewConfigurationListResponse

`func NewConfigurationListResponse(categories []CategoryInfo, configurations map[string][]ConfigurationItem, ) *ConfigurationListResponse`

NewConfigurationListResponse instantiates a new ConfigurationListResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewConfigurationListResponseWithDefaults

`func NewConfigurationListResponseWithDefaults() *ConfigurationListResponse`

NewConfigurationListResponseWithDefaults instantiates a new ConfigurationListResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCategories

`func (o *ConfigurationListResponse) GetCategories() []CategoryInfo`

GetCategories returns the Categories field if non-nil, zero value otherwise.

### GetCategoriesOk

`func (o *ConfigurationListResponse) GetCategoriesOk() (*[]CategoryInfo, bool)`

GetCategoriesOk returns a tuple with the Categories field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCategories

`func (o *ConfigurationListResponse) SetCategories(v []CategoryInfo)`

SetCategories sets Categories field to given value.


### GetConfigurations

`func (o *ConfigurationListResponse) GetConfigurations() map[string][]ConfigurationItem`

GetConfigurations returns the Configurations field if non-nil, zero value otherwise.

### GetConfigurationsOk

`func (o *ConfigurationListResponse) GetConfigurationsOk() (*map[string][]ConfigurationItem, bool)`

GetConfigurationsOk returns a tuple with the Configurations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfigurations

`func (o *ConfigurationListResponse) SetConfigurations(v map[string][]ConfigurationItem)`

SetConfigurations sets Configurations field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


