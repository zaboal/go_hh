# MailTemplatesMailTemplate

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**EmployerState** | Pointer to **string** | Работодательский статус, соответствующий данному шаблону; может отсутствовать если шаблон не связан с каким-либо статусом | [optional] 
**Id** | **string** | Идентификатор шаблона | 
**Name** | **string** | Имя шаблона | 
**TemplateModified** | **bool** | Был ли изменен шаблон работодателем или же используется стандартный шаблон с текстом от Хэдхантер | 
**Text** | **string** | Текст шаблона | 
**Type** | **string** | Способ доставки сообщения сформированного из шаблона | 
**Variables** | [**[]IncludesIdName**](IncludesIdName.md) | Переменные доступные для вставки в шаблон | 

## Methods

### NewMailTemplatesMailTemplate

`func NewMailTemplatesMailTemplate(id string, name string, templateModified bool, text string, type_ string, variables []IncludesIdName, ) *MailTemplatesMailTemplate`

NewMailTemplatesMailTemplate instantiates a new MailTemplatesMailTemplate object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMailTemplatesMailTemplateWithDefaults

`func NewMailTemplatesMailTemplateWithDefaults() *MailTemplatesMailTemplate`

NewMailTemplatesMailTemplateWithDefaults instantiates a new MailTemplatesMailTemplate object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEmployerState

`func (o *MailTemplatesMailTemplate) GetEmployerState() string`

GetEmployerState returns the EmployerState field if non-nil, zero value otherwise.

### GetEmployerStateOk

`func (o *MailTemplatesMailTemplate) GetEmployerStateOk() (*string, bool)`

GetEmployerStateOk returns a tuple with the EmployerState field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmployerState

`func (o *MailTemplatesMailTemplate) SetEmployerState(v string)`

SetEmployerState sets EmployerState field to given value.

### HasEmployerState

`func (o *MailTemplatesMailTemplate) HasEmployerState() bool`

HasEmployerState returns a boolean if a field has been set.

### GetId

`func (o *MailTemplatesMailTemplate) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *MailTemplatesMailTemplate) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *MailTemplatesMailTemplate) SetId(v string)`

SetId sets Id field to given value.


### GetName

`func (o *MailTemplatesMailTemplate) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *MailTemplatesMailTemplate) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *MailTemplatesMailTemplate) SetName(v string)`

SetName sets Name field to given value.


### GetTemplateModified

`func (o *MailTemplatesMailTemplate) GetTemplateModified() bool`

GetTemplateModified returns the TemplateModified field if non-nil, zero value otherwise.

### GetTemplateModifiedOk

`func (o *MailTemplatesMailTemplate) GetTemplateModifiedOk() (*bool, bool)`

GetTemplateModifiedOk returns a tuple with the TemplateModified field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTemplateModified

`func (o *MailTemplatesMailTemplate) SetTemplateModified(v bool)`

SetTemplateModified sets TemplateModified field to given value.


### GetText

`func (o *MailTemplatesMailTemplate) GetText() string`

GetText returns the Text field if non-nil, zero value otherwise.

### GetTextOk

`func (o *MailTemplatesMailTemplate) GetTextOk() (*string, bool)`

GetTextOk returns a tuple with the Text field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetText

`func (o *MailTemplatesMailTemplate) SetText(v string)`

SetText sets Text field to given value.


### GetType

`func (o *MailTemplatesMailTemplate) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *MailTemplatesMailTemplate) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *MailTemplatesMailTemplate) SetType(v string)`

SetType sets Type field to given value.


### GetVariables

`func (o *MailTemplatesMailTemplate) GetVariables() []IncludesIdName`

GetVariables returns the Variables field if non-nil, zero value otherwise.

### GetVariablesOk

`func (o *MailTemplatesMailTemplate) GetVariablesOk() (*[]IncludesIdName, bool)`

GetVariablesOk returns a tuple with the Variables field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVariables

`func (o *MailTemplatesMailTemplate) SetVariables(v []IncludesIdName)`

SetVariables sets Variables field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


