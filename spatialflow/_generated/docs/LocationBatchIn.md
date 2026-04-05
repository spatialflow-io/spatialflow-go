# LocationBatchIn

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Locations** | [**[]LocationPointIn**](LocationPointIn.md) |  | 
**IdempotencyKey** | Pointer to **NullableString** |  | [optional] 

## Methods

### NewLocationBatchIn

`func NewLocationBatchIn(locations []LocationPointIn, ) *LocationBatchIn`

NewLocationBatchIn instantiates a new LocationBatchIn object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLocationBatchInWithDefaults

`func NewLocationBatchInWithDefaults() *LocationBatchIn`

NewLocationBatchInWithDefaults instantiates a new LocationBatchIn object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLocations

`func (o *LocationBatchIn) GetLocations() []LocationPointIn`

GetLocations returns the Locations field if non-nil, zero value otherwise.

### GetLocationsOk

`func (o *LocationBatchIn) GetLocationsOk() (*[]LocationPointIn, bool)`

GetLocationsOk returns a tuple with the Locations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocations

`func (o *LocationBatchIn) SetLocations(v []LocationPointIn)`

SetLocations sets Locations field to given value.


### GetIdempotencyKey

`func (o *LocationBatchIn) GetIdempotencyKey() string`

GetIdempotencyKey returns the IdempotencyKey field if non-nil, zero value otherwise.

### GetIdempotencyKeyOk

`func (o *LocationBatchIn) GetIdempotencyKeyOk() (*string, bool)`

GetIdempotencyKeyOk returns a tuple with the IdempotencyKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdempotencyKey

`func (o *LocationBatchIn) SetIdempotencyKey(v string)`

SetIdempotencyKey sets IdempotencyKey field to given value.

### HasIdempotencyKey

`func (o *LocationBatchIn) HasIdempotencyKey() bool`

HasIdempotencyKey returns a boolean if a field has been set.

### SetIdempotencyKeyNil

`func (o *LocationBatchIn) SetIdempotencyKeyNil(b bool)`

 SetIdempotencyKeyNil sets the value for IdempotencyKey to be an explicit nil

### UnsetIdempotencyKey
`func (o *LocationBatchIn) UnsetIdempotencyKey()`

UnsetIdempotencyKey ensures that no value is present for IdempotencyKey, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


