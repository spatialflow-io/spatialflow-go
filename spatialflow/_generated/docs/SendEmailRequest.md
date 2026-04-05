# SendEmailRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ToEmail** | **string** |  | 
**Template** | **string** |  | 
**Context** | **map[string]interface{}** |  | 
**Subject** | Pointer to **NullableString** |  | [optional] 
**Priority** | Pointer to **bool** |  | [optional] [default to false]

## Methods

### NewSendEmailRequest

`func NewSendEmailRequest(toEmail string, template string, context map[string]interface{}, ) *SendEmailRequest`

NewSendEmailRequest instantiates a new SendEmailRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSendEmailRequestWithDefaults

`func NewSendEmailRequestWithDefaults() *SendEmailRequest`

NewSendEmailRequestWithDefaults instantiates a new SendEmailRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetToEmail

`func (o *SendEmailRequest) GetToEmail() string`

GetToEmail returns the ToEmail field if non-nil, zero value otherwise.

### GetToEmailOk

`func (o *SendEmailRequest) GetToEmailOk() (*string, bool)`

GetToEmailOk returns a tuple with the ToEmail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToEmail

`func (o *SendEmailRequest) SetToEmail(v string)`

SetToEmail sets ToEmail field to given value.


### GetTemplate

`func (o *SendEmailRequest) GetTemplate() string`

GetTemplate returns the Template field if non-nil, zero value otherwise.

### GetTemplateOk

`func (o *SendEmailRequest) GetTemplateOk() (*string, bool)`

GetTemplateOk returns a tuple with the Template field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTemplate

`func (o *SendEmailRequest) SetTemplate(v string)`

SetTemplate sets Template field to given value.


### GetContext

`func (o *SendEmailRequest) GetContext() map[string]interface{}`

GetContext returns the Context field if non-nil, zero value otherwise.

### GetContextOk

`func (o *SendEmailRequest) GetContextOk() (*map[string]interface{}, bool)`

GetContextOk returns a tuple with the Context field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContext

`func (o *SendEmailRequest) SetContext(v map[string]interface{})`

SetContext sets Context field to given value.


### GetSubject

`func (o *SendEmailRequest) GetSubject() string`

GetSubject returns the Subject field if non-nil, zero value otherwise.

### GetSubjectOk

`func (o *SendEmailRequest) GetSubjectOk() (*string, bool)`

GetSubjectOk returns a tuple with the Subject field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubject

`func (o *SendEmailRequest) SetSubject(v string)`

SetSubject sets Subject field to given value.

### HasSubject

`func (o *SendEmailRequest) HasSubject() bool`

HasSubject returns a boolean if a field has been set.

### SetSubjectNil

`func (o *SendEmailRequest) SetSubjectNil(b bool)`

 SetSubjectNil sets the value for Subject to be an explicit nil

### UnsetSubject
`func (o *SendEmailRequest) UnsetSubject()`

UnsetSubject ensures that no value is present for Subject, not even an explicit nil
### GetPriority

`func (o *SendEmailRequest) GetPriority() bool`

GetPriority returns the Priority field if non-nil, zero value otherwise.

### GetPriorityOk

`func (o *SendEmailRequest) GetPriorityOk() (*bool, bool)`

GetPriorityOk returns a tuple with the Priority field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPriority

`func (o *SendEmailRequest) SetPriority(v bool)`

SetPriority sets Priority field to given value.

### HasPriority

`func (o *SendEmailRequest) HasPriority() bool`

HasPriority returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


