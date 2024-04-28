# NegotiationsNegotiationMessageTemplate

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**EmployerState** | Pointer to **string** | Работодательский статус, соответствующий данному шаблону; может отсутствовать если шаблон не связан с каким-либо статусом | [optional] 
**TemplateModified** | **bool** | Был ли изменен шаблон работодателем или же используется стандартный шаблон с текстом от Хэдхантер | 
**Text** | **string** | Текст шаблона | 

## Methods

### NewNegotiationsNegotiationMessageTemplate

`func NewNegotiationsNegotiationMessageTemplate(templateModified bool, text string, ) *NegotiationsNegotiationMessageTemplate`

NewNegotiationsNegotiationMessageTemplate instantiates a new NegotiationsNegotiationMessageTemplate object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNegotiationsNegotiationMessageTemplateWithDefaults

`func NewNegotiationsNegotiationMessageTemplateWithDefaults() *NegotiationsNegotiationMessageTemplate`

NewNegotiationsNegotiationMessageTemplateWithDefaults instantiates a new NegotiationsNegotiationMessageTemplate object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEmployerState

`func (o *NegotiationsNegotiationMessageTemplate) GetEmployerState() string`

GetEmployerState returns the EmployerState field if non-nil, zero value otherwise.

### GetEmployerStateOk

`func (o *NegotiationsNegotiationMessageTemplate) GetEmployerStateOk() (*string, bool)`

GetEmployerStateOk returns a tuple with the EmployerState field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmployerState

`func (o *NegotiationsNegotiationMessageTemplate) SetEmployerState(v string)`

SetEmployerState sets EmployerState field to given value.

### HasEmployerState

`func (o *NegotiationsNegotiationMessageTemplate) HasEmployerState() bool`

HasEmployerState returns a boolean if a field has been set.

### GetTemplateModified

`func (o *NegotiationsNegotiationMessageTemplate) GetTemplateModified() bool`

GetTemplateModified returns the TemplateModified field if non-nil, zero value otherwise.

### GetTemplateModifiedOk

`func (o *NegotiationsNegotiationMessageTemplate) GetTemplateModifiedOk() (*bool, bool)`

GetTemplateModifiedOk returns a tuple with the TemplateModified field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTemplateModified

`func (o *NegotiationsNegotiationMessageTemplate) SetTemplateModified(v bool)`

SetTemplateModified sets TemplateModified field to given value.


### GetText

`func (o *NegotiationsNegotiationMessageTemplate) GetText() string`

GetText returns the Text field if non-nil, zero value otherwise.

### GetTextOk

`func (o *NegotiationsNegotiationMessageTemplate) GetTextOk() (*string, bool)`

GetTextOk returns a tuple with the Text field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetText

`func (o *NegotiationsNegotiationMessageTemplate) SetText(v string)`

SetText sets Text field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


