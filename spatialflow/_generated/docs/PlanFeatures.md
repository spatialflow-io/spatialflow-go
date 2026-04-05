# PlanFeatures

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ApiCalls** | **string** |  | 
**Geofences** | **string** |  | 
**Webhooks** | **string** |  | 
**TestPoints** | **string** |  | 
**Support** | **string** |  | 
**CustomDomains** | Pointer to **NullableString** |  | [optional] 
**Sla** | Pointer to **NullableString** |  | [optional] 
**FeatureList** | Pointer to **[]string** | Feature list from model | [optional] 

## Methods

### NewPlanFeatures

`func NewPlanFeatures(apiCalls string, geofences string, webhooks string, testPoints string, support string, ) *PlanFeatures`

NewPlanFeatures instantiates a new PlanFeatures object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPlanFeaturesWithDefaults

`func NewPlanFeaturesWithDefaults() *PlanFeatures`

NewPlanFeaturesWithDefaults instantiates a new PlanFeatures object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetApiCalls

`func (o *PlanFeatures) GetApiCalls() string`

GetApiCalls returns the ApiCalls field if non-nil, zero value otherwise.

### GetApiCallsOk

`func (o *PlanFeatures) GetApiCallsOk() (*string, bool)`

GetApiCallsOk returns a tuple with the ApiCalls field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApiCalls

`func (o *PlanFeatures) SetApiCalls(v string)`

SetApiCalls sets ApiCalls field to given value.


### GetGeofences

`func (o *PlanFeatures) GetGeofences() string`

GetGeofences returns the Geofences field if non-nil, zero value otherwise.

### GetGeofencesOk

`func (o *PlanFeatures) GetGeofencesOk() (*string, bool)`

GetGeofencesOk returns a tuple with the Geofences field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGeofences

`func (o *PlanFeatures) SetGeofences(v string)`

SetGeofences sets Geofences field to given value.


### GetWebhooks

`func (o *PlanFeatures) GetWebhooks() string`

GetWebhooks returns the Webhooks field if non-nil, zero value otherwise.

### GetWebhooksOk

`func (o *PlanFeatures) GetWebhooksOk() (*string, bool)`

GetWebhooksOk returns a tuple with the Webhooks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWebhooks

`func (o *PlanFeatures) SetWebhooks(v string)`

SetWebhooks sets Webhooks field to given value.


### GetTestPoints

`func (o *PlanFeatures) GetTestPoints() string`

GetTestPoints returns the TestPoints field if non-nil, zero value otherwise.

### GetTestPointsOk

`func (o *PlanFeatures) GetTestPointsOk() (*string, bool)`

GetTestPointsOk returns a tuple with the TestPoints field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTestPoints

`func (o *PlanFeatures) SetTestPoints(v string)`

SetTestPoints sets TestPoints field to given value.


### GetSupport

`func (o *PlanFeatures) GetSupport() string`

GetSupport returns the Support field if non-nil, zero value otherwise.

### GetSupportOk

`func (o *PlanFeatures) GetSupportOk() (*string, bool)`

GetSupportOk returns a tuple with the Support field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSupport

`func (o *PlanFeatures) SetSupport(v string)`

SetSupport sets Support field to given value.


### GetCustomDomains

`func (o *PlanFeatures) GetCustomDomains() string`

GetCustomDomains returns the CustomDomains field if non-nil, zero value otherwise.

### GetCustomDomainsOk

`func (o *PlanFeatures) GetCustomDomainsOk() (*string, bool)`

GetCustomDomainsOk returns a tuple with the CustomDomains field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomDomains

`func (o *PlanFeatures) SetCustomDomains(v string)`

SetCustomDomains sets CustomDomains field to given value.

### HasCustomDomains

`func (o *PlanFeatures) HasCustomDomains() bool`

HasCustomDomains returns a boolean if a field has been set.

### SetCustomDomainsNil

`func (o *PlanFeatures) SetCustomDomainsNil(b bool)`

 SetCustomDomainsNil sets the value for CustomDomains to be an explicit nil

### UnsetCustomDomains
`func (o *PlanFeatures) UnsetCustomDomains()`

UnsetCustomDomains ensures that no value is present for CustomDomains, not even an explicit nil
### GetSla

`func (o *PlanFeatures) GetSla() string`

GetSla returns the Sla field if non-nil, zero value otherwise.

### GetSlaOk

`func (o *PlanFeatures) GetSlaOk() (*string, bool)`

GetSlaOk returns a tuple with the Sla field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSla

`func (o *PlanFeatures) SetSla(v string)`

SetSla sets Sla field to given value.

### HasSla

`func (o *PlanFeatures) HasSla() bool`

HasSla returns a boolean if a field has been set.

### SetSlaNil

`func (o *PlanFeatures) SetSlaNil(b bool)`

 SetSlaNil sets the value for Sla to be an explicit nil

### UnsetSla
`func (o *PlanFeatures) UnsetSla()`

UnsetSla ensures that no value is present for Sla, not even an explicit nil
### GetFeatureList

`func (o *PlanFeatures) GetFeatureList() []string`

GetFeatureList returns the FeatureList field if non-nil, zero value otherwise.

### GetFeatureListOk

`func (o *PlanFeatures) GetFeatureListOk() (*[]string, bool)`

GetFeatureListOk returns a tuple with the FeatureList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFeatureList

`func (o *PlanFeatures) SetFeatureList(v []string)`

SetFeatureList sets FeatureList field to given value.

### HasFeatureList

`func (o *PlanFeatures) HasFeatureList() bool`

HasFeatureList returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


