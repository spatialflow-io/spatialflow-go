# EmailPreviewErrorResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Detail** | **string** |  | 
**ErrorCode** | Pointer to **NullableString** |  | [optional] 
**Details** | Pointer to [**NullableDetails**](Details.md) |  | [optional] 
**Template** | Pointer to **NullableString** |  | [optional] 
**Format** | Pointer to **NullableString** |  | [optional] 

## Methods

### NewEmailPreviewErrorResponse

`func NewEmailPreviewErrorResponse(detail string, ) *EmailPreviewErrorResponse`

NewEmailPreviewErrorResponse instantiates a new EmailPreviewErrorResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEmailPreviewErrorResponseWithDefaults

`func NewEmailPreviewErrorResponseWithDefaults() *EmailPreviewErrorResponse`

NewEmailPreviewErrorResponseWithDefaults instantiates a new EmailPreviewErrorResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDetail

`func (o *EmailPreviewErrorResponse) GetDetail() string`

GetDetail returns the Detail field if non-nil, zero value otherwise.

### GetDetailOk

`func (o *EmailPreviewErrorResponse) GetDetailOk() (*string, bool)`

GetDetailOk returns a tuple with the Detail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDetail

`func (o *EmailPreviewErrorResponse) SetDetail(v string)`

SetDetail sets Detail field to given value.


### GetErrorCode

`func (o *EmailPreviewErrorResponse) GetErrorCode() string`

GetErrorCode returns the ErrorCode field if non-nil, zero value otherwise.

### GetErrorCodeOk

`func (o *EmailPreviewErrorResponse) GetErrorCodeOk() (*string, bool)`

GetErrorCodeOk returns a tuple with the ErrorCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrorCode

`func (o *EmailPreviewErrorResponse) SetErrorCode(v string)`

SetErrorCode sets ErrorCode field to given value.

### HasErrorCode

`func (o *EmailPreviewErrorResponse) HasErrorCode() bool`

HasErrorCode returns a boolean if a field has been set.

### SetErrorCodeNil

`func (o *EmailPreviewErrorResponse) SetErrorCodeNil(b bool)`

 SetErrorCodeNil sets the value for ErrorCode to be an explicit nil

### UnsetErrorCode
`func (o *EmailPreviewErrorResponse) UnsetErrorCode()`

UnsetErrorCode ensures that no value is present for ErrorCode, not even an explicit nil
### GetDetails

`func (o *EmailPreviewErrorResponse) GetDetails() Details`

GetDetails returns the Details field if non-nil, zero value otherwise.

### GetDetailsOk

`func (o *EmailPreviewErrorResponse) GetDetailsOk() (*Details, bool)`

GetDetailsOk returns a tuple with the Details field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDetails

`func (o *EmailPreviewErrorResponse) SetDetails(v Details)`

SetDetails sets Details field to given value.

### HasDetails

`func (o *EmailPreviewErrorResponse) HasDetails() bool`

HasDetails returns a boolean if a field has been set.

### SetDetailsNil

`func (o *EmailPreviewErrorResponse) SetDetailsNil(b bool)`

 SetDetailsNil sets the value for Details to be an explicit nil

### UnsetDetails
`func (o *EmailPreviewErrorResponse) UnsetDetails()`

UnsetDetails ensures that no value is present for Details, not even an explicit nil
### GetTemplate

`func (o *EmailPreviewErrorResponse) GetTemplate() string`

GetTemplate returns the Template field if non-nil, zero value otherwise.

### GetTemplateOk

`func (o *EmailPreviewErrorResponse) GetTemplateOk() (*string, bool)`

GetTemplateOk returns a tuple with the Template field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTemplate

`func (o *EmailPreviewErrorResponse) SetTemplate(v string)`

SetTemplate sets Template field to given value.

### HasTemplate

`func (o *EmailPreviewErrorResponse) HasTemplate() bool`

HasTemplate returns a boolean if a field has been set.

### SetTemplateNil

`func (o *EmailPreviewErrorResponse) SetTemplateNil(b bool)`

 SetTemplateNil sets the value for Template to be an explicit nil

### UnsetTemplate
`func (o *EmailPreviewErrorResponse) UnsetTemplate()`

UnsetTemplate ensures that no value is present for Template, not even an explicit nil
### GetFormat

`func (o *EmailPreviewErrorResponse) GetFormat() string`

GetFormat returns the Format field if non-nil, zero value otherwise.

### GetFormatOk

`func (o *EmailPreviewErrorResponse) GetFormatOk() (*string, bool)`

GetFormatOk returns a tuple with the Format field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFormat

`func (o *EmailPreviewErrorResponse) SetFormat(v string)`

SetFormat sets Format field to given value.

### HasFormat

`func (o *EmailPreviewErrorResponse) HasFormat() bool`

HasFormat returns a boolean if a field has been set.

### SetFormatNil

`func (o *EmailPreviewErrorResponse) SetFormatNil(b bool)`

 SetFormatNil sets the value for Format to be an explicit nil

### UnsetFormat
`func (o *EmailPreviewErrorResponse) UnsetFormat()`

UnsetFormat ensures that no value is present for Format, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


