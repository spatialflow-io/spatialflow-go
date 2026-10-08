# GeocodePlaceResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Point** | **[]float32** | [longitude, latitude] of the geocoded address. | 
**Confidence** | **float32** | Geocoder confidence score (0-1). | 
**NormalizedAddress** | **string** | Full normalized address string from geocoder. | 
**NormalizedComponents** | **map[string]string** | Structured address components from geocoder. | 

## Methods

### NewGeocodePlaceResponse

`func NewGeocodePlaceResponse(point []float32, confidence float32, normalizedAddress string, normalizedComponents map[string]string, ) *GeocodePlaceResponse`

NewGeocodePlaceResponse instantiates a new GeocodePlaceResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGeocodePlaceResponseWithDefaults

`func NewGeocodePlaceResponseWithDefaults() *GeocodePlaceResponse`

NewGeocodePlaceResponseWithDefaults instantiates a new GeocodePlaceResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPoint

`func (o *GeocodePlaceResponse) GetPoint() []float32`

GetPoint returns the Point field if non-nil, zero value otherwise.

### GetPointOk

`func (o *GeocodePlaceResponse) GetPointOk() (*[]float32, bool)`

GetPointOk returns a tuple with the Point field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPoint

`func (o *GeocodePlaceResponse) SetPoint(v []float32)`

SetPoint sets Point field to given value.


### GetConfidence

`func (o *GeocodePlaceResponse) GetConfidence() float32`

GetConfidence returns the Confidence field if non-nil, zero value otherwise.

### GetConfidenceOk

`func (o *GeocodePlaceResponse) GetConfidenceOk() (*float32, bool)`

GetConfidenceOk returns a tuple with the Confidence field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfidence

`func (o *GeocodePlaceResponse) SetConfidence(v float32)`

SetConfidence sets Confidence field to given value.


### GetNormalizedAddress

`func (o *GeocodePlaceResponse) GetNormalizedAddress() string`

GetNormalizedAddress returns the NormalizedAddress field if non-nil, zero value otherwise.

### GetNormalizedAddressOk

`func (o *GeocodePlaceResponse) GetNormalizedAddressOk() (*string, bool)`

GetNormalizedAddressOk returns a tuple with the NormalizedAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNormalizedAddress

`func (o *GeocodePlaceResponse) SetNormalizedAddress(v string)`

SetNormalizedAddress sets NormalizedAddress field to given value.


### GetNormalizedComponents

`func (o *GeocodePlaceResponse) GetNormalizedComponents() map[string]string`

GetNormalizedComponents returns the NormalizedComponents field if non-nil, zero value otherwise.

### GetNormalizedComponentsOk

`func (o *GeocodePlaceResponse) GetNormalizedComponentsOk() (*map[string]string, bool)`

GetNormalizedComponentsOk returns a tuple with the NormalizedComponents field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNormalizedComponents

`func (o *GeocodePlaceResponse) SetNormalizedComponents(v map[string]string)`

SetNormalizedComponents sets NormalizedComponents field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


