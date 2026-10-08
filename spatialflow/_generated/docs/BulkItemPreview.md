# BulkItemPreview

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Index** | **int32** |  | 
**InputAddress** | **string** |  | 
**Status** | **string** |  | 
**NormalizedAddress** | Pointer to **NullableString** |  | [optional] 
**Lat** | Pointer to **NullableFloat32** |  | [optional] 
**Lng** | Pointer to **NullableFloat32** |  | [optional] 
**Confidence** | Pointer to **NullableFloat32** |  | [optional] 
**DedupMatch** | Pointer to **map[string]interface{}** |  | [optional] 
**ErrorReason** | Pointer to **NullableString** |  | [optional] 
**ErrorCode** | Pointer to **NullableString** |  | [optional] 
**Tags** | Pointer to **[]string** |  | [optional] 
**OutcomeClass** | Pointer to **NullableString** |  | [optional] 
**MatchType** | Pointer to **NullableString** |  | [optional] 
**Interpolated** | Pointer to **NullableBool** |  | [optional] 
**CandidateCount** | Pointer to **NullableInt32** |  | [optional] 

## Methods

### NewBulkItemPreview

`func NewBulkItemPreview(index int32, inputAddress string, status string, ) *BulkItemPreview`

NewBulkItemPreview instantiates a new BulkItemPreview object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBulkItemPreviewWithDefaults

`func NewBulkItemPreviewWithDefaults() *BulkItemPreview`

NewBulkItemPreviewWithDefaults instantiates a new BulkItemPreview object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIndex

`func (o *BulkItemPreview) GetIndex() int32`

GetIndex returns the Index field if non-nil, zero value otherwise.

### GetIndexOk

`func (o *BulkItemPreview) GetIndexOk() (*int32, bool)`

GetIndexOk returns a tuple with the Index field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIndex

`func (o *BulkItemPreview) SetIndex(v int32)`

SetIndex sets Index field to given value.


### GetInputAddress

`func (o *BulkItemPreview) GetInputAddress() string`

GetInputAddress returns the InputAddress field if non-nil, zero value otherwise.

### GetInputAddressOk

`func (o *BulkItemPreview) GetInputAddressOk() (*string, bool)`

GetInputAddressOk returns a tuple with the InputAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInputAddress

`func (o *BulkItemPreview) SetInputAddress(v string)`

SetInputAddress sets InputAddress field to given value.


### GetStatus

`func (o *BulkItemPreview) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *BulkItemPreview) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *BulkItemPreview) SetStatus(v string)`

SetStatus sets Status field to given value.


### GetNormalizedAddress

`func (o *BulkItemPreview) GetNormalizedAddress() string`

GetNormalizedAddress returns the NormalizedAddress field if non-nil, zero value otherwise.

### GetNormalizedAddressOk

`func (o *BulkItemPreview) GetNormalizedAddressOk() (*string, bool)`

GetNormalizedAddressOk returns a tuple with the NormalizedAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNormalizedAddress

`func (o *BulkItemPreview) SetNormalizedAddress(v string)`

SetNormalizedAddress sets NormalizedAddress field to given value.

### HasNormalizedAddress

`func (o *BulkItemPreview) HasNormalizedAddress() bool`

HasNormalizedAddress returns a boolean if a field has been set.

### SetNormalizedAddressNil

`func (o *BulkItemPreview) SetNormalizedAddressNil(b bool)`

 SetNormalizedAddressNil sets the value for NormalizedAddress to be an explicit nil

### UnsetNormalizedAddress
`func (o *BulkItemPreview) UnsetNormalizedAddress()`

UnsetNormalizedAddress ensures that no value is present for NormalizedAddress, not even an explicit nil
### GetLat

`func (o *BulkItemPreview) GetLat() float32`

GetLat returns the Lat field if non-nil, zero value otherwise.

### GetLatOk

`func (o *BulkItemPreview) GetLatOk() (*float32, bool)`

GetLatOk returns a tuple with the Lat field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLat

`func (o *BulkItemPreview) SetLat(v float32)`

SetLat sets Lat field to given value.

### HasLat

`func (o *BulkItemPreview) HasLat() bool`

HasLat returns a boolean if a field has been set.

### SetLatNil

`func (o *BulkItemPreview) SetLatNil(b bool)`

 SetLatNil sets the value for Lat to be an explicit nil

### UnsetLat
`func (o *BulkItemPreview) UnsetLat()`

UnsetLat ensures that no value is present for Lat, not even an explicit nil
### GetLng

`func (o *BulkItemPreview) GetLng() float32`

GetLng returns the Lng field if non-nil, zero value otherwise.

### GetLngOk

`func (o *BulkItemPreview) GetLngOk() (*float32, bool)`

GetLngOk returns a tuple with the Lng field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLng

`func (o *BulkItemPreview) SetLng(v float32)`

SetLng sets Lng field to given value.

### HasLng

`func (o *BulkItemPreview) HasLng() bool`

HasLng returns a boolean if a field has been set.

### SetLngNil

`func (o *BulkItemPreview) SetLngNil(b bool)`

 SetLngNil sets the value for Lng to be an explicit nil

### UnsetLng
`func (o *BulkItemPreview) UnsetLng()`

UnsetLng ensures that no value is present for Lng, not even an explicit nil
### GetConfidence

`func (o *BulkItemPreview) GetConfidence() float32`

GetConfidence returns the Confidence field if non-nil, zero value otherwise.

### GetConfidenceOk

`func (o *BulkItemPreview) GetConfidenceOk() (*float32, bool)`

GetConfidenceOk returns a tuple with the Confidence field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfidence

`func (o *BulkItemPreview) SetConfidence(v float32)`

SetConfidence sets Confidence field to given value.

### HasConfidence

`func (o *BulkItemPreview) HasConfidence() bool`

HasConfidence returns a boolean if a field has been set.

### SetConfidenceNil

`func (o *BulkItemPreview) SetConfidenceNil(b bool)`

 SetConfidenceNil sets the value for Confidence to be an explicit nil

### UnsetConfidence
`func (o *BulkItemPreview) UnsetConfidence()`

UnsetConfidence ensures that no value is present for Confidence, not even an explicit nil
### GetDedupMatch

`func (o *BulkItemPreview) GetDedupMatch() map[string]interface{}`

GetDedupMatch returns the DedupMatch field if non-nil, zero value otherwise.

### GetDedupMatchOk

`func (o *BulkItemPreview) GetDedupMatchOk() (*map[string]interface{}, bool)`

GetDedupMatchOk returns a tuple with the DedupMatch field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDedupMatch

`func (o *BulkItemPreview) SetDedupMatch(v map[string]interface{})`

SetDedupMatch sets DedupMatch field to given value.

### HasDedupMatch

`func (o *BulkItemPreview) HasDedupMatch() bool`

HasDedupMatch returns a boolean if a field has been set.

### SetDedupMatchNil

`func (o *BulkItemPreview) SetDedupMatchNil(b bool)`

 SetDedupMatchNil sets the value for DedupMatch to be an explicit nil

### UnsetDedupMatch
`func (o *BulkItemPreview) UnsetDedupMatch()`

UnsetDedupMatch ensures that no value is present for DedupMatch, not even an explicit nil
### GetErrorReason

`func (o *BulkItemPreview) GetErrorReason() string`

GetErrorReason returns the ErrorReason field if non-nil, zero value otherwise.

### GetErrorReasonOk

`func (o *BulkItemPreview) GetErrorReasonOk() (*string, bool)`

GetErrorReasonOk returns a tuple with the ErrorReason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrorReason

`func (o *BulkItemPreview) SetErrorReason(v string)`

SetErrorReason sets ErrorReason field to given value.

### HasErrorReason

`func (o *BulkItemPreview) HasErrorReason() bool`

HasErrorReason returns a boolean if a field has been set.

### SetErrorReasonNil

`func (o *BulkItemPreview) SetErrorReasonNil(b bool)`

 SetErrorReasonNil sets the value for ErrorReason to be an explicit nil

### UnsetErrorReason
`func (o *BulkItemPreview) UnsetErrorReason()`

UnsetErrorReason ensures that no value is present for ErrorReason, not even an explicit nil
### GetErrorCode

`func (o *BulkItemPreview) GetErrorCode() string`

GetErrorCode returns the ErrorCode field if non-nil, zero value otherwise.

### GetErrorCodeOk

`func (o *BulkItemPreview) GetErrorCodeOk() (*string, bool)`

GetErrorCodeOk returns a tuple with the ErrorCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrorCode

`func (o *BulkItemPreview) SetErrorCode(v string)`

SetErrorCode sets ErrorCode field to given value.

### HasErrorCode

`func (o *BulkItemPreview) HasErrorCode() bool`

HasErrorCode returns a boolean if a field has been set.

### SetErrorCodeNil

`func (o *BulkItemPreview) SetErrorCodeNil(b bool)`

 SetErrorCodeNil sets the value for ErrorCode to be an explicit nil

### UnsetErrorCode
`func (o *BulkItemPreview) UnsetErrorCode()`

UnsetErrorCode ensures that no value is present for ErrorCode, not even an explicit nil
### GetTags

`func (o *BulkItemPreview) GetTags() []string`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *BulkItemPreview) GetTagsOk() (*[]string, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *BulkItemPreview) SetTags(v []string)`

SetTags sets Tags field to given value.

### HasTags

`func (o *BulkItemPreview) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetOutcomeClass

`func (o *BulkItemPreview) GetOutcomeClass() string`

GetOutcomeClass returns the OutcomeClass field if non-nil, zero value otherwise.

### GetOutcomeClassOk

`func (o *BulkItemPreview) GetOutcomeClassOk() (*string, bool)`

GetOutcomeClassOk returns a tuple with the OutcomeClass field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOutcomeClass

`func (o *BulkItemPreview) SetOutcomeClass(v string)`

SetOutcomeClass sets OutcomeClass field to given value.

### HasOutcomeClass

`func (o *BulkItemPreview) HasOutcomeClass() bool`

HasOutcomeClass returns a boolean if a field has been set.

### SetOutcomeClassNil

`func (o *BulkItemPreview) SetOutcomeClassNil(b bool)`

 SetOutcomeClassNil sets the value for OutcomeClass to be an explicit nil

### UnsetOutcomeClass
`func (o *BulkItemPreview) UnsetOutcomeClass()`

UnsetOutcomeClass ensures that no value is present for OutcomeClass, not even an explicit nil
### GetMatchType

`func (o *BulkItemPreview) GetMatchType() string`

GetMatchType returns the MatchType field if non-nil, zero value otherwise.

### GetMatchTypeOk

`func (o *BulkItemPreview) GetMatchTypeOk() (*string, bool)`

GetMatchTypeOk returns a tuple with the MatchType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMatchType

`func (o *BulkItemPreview) SetMatchType(v string)`

SetMatchType sets MatchType field to given value.

### HasMatchType

`func (o *BulkItemPreview) HasMatchType() bool`

HasMatchType returns a boolean if a field has been set.

### SetMatchTypeNil

`func (o *BulkItemPreview) SetMatchTypeNil(b bool)`

 SetMatchTypeNil sets the value for MatchType to be an explicit nil

### UnsetMatchType
`func (o *BulkItemPreview) UnsetMatchType()`

UnsetMatchType ensures that no value is present for MatchType, not even an explicit nil
### GetInterpolated

`func (o *BulkItemPreview) GetInterpolated() bool`

GetInterpolated returns the Interpolated field if non-nil, zero value otherwise.

### GetInterpolatedOk

`func (o *BulkItemPreview) GetInterpolatedOk() (*bool, bool)`

GetInterpolatedOk returns a tuple with the Interpolated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInterpolated

`func (o *BulkItemPreview) SetInterpolated(v bool)`

SetInterpolated sets Interpolated field to given value.

### HasInterpolated

`func (o *BulkItemPreview) HasInterpolated() bool`

HasInterpolated returns a boolean if a field has been set.

### SetInterpolatedNil

`func (o *BulkItemPreview) SetInterpolatedNil(b bool)`

 SetInterpolatedNil sets the value for Interpolated to be an explicit nil

### UnsetInterpolated
`func (o *BulkItemPreview) UnsetInterpolated()`

UnsetInterpolated ensures that no value is present for Interpolated, not even an explicit nil
### GetCandidateCount

`func (o *BulkItemPreview) GetCandidateCount() int32`

GetCandidateCount returns the CandidateCount field if non-nil, zero value otherwise.

### GetCandidateCountOk

`func (o *BulkItemPreview) GetCandidateCountOk() (*int32, bool)`

GetCandidateCountOk returns a tuple with the CandidateCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCandidateCount

`func (o *BulkItemPreview) SetCandidateCount(v int32)`

SetCandidateCount sets CandidateCount field to given value.

### HasCandidateCount

`func (o *BulkItemPreview) HasCandidateCount() bool`

HasCandidateCount returns a boolean if a field has been set.

### SetCandidateCountNil

`func (o *BulkItemPreview) SetCandidateCountNil(b bool)`

 SetCandidateCountNil sets the value for CandidateCount to be an explicit nil

### UnsetCandidateCount
`func (o *BulkItemPreview) UnsetCandidateCount()`

UnsetCandidateCount ensures that no value is present for CandidateCount, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


