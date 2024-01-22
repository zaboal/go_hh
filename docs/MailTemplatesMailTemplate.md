# MailTemplatesMailTemplate

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | Идентификатор шаблона | 
**Name** | **string** | Имя шаблона | 
**Text** | **string** | Текст шаблона | 
**Type** | **string** | Способ доставки сообщения сформированного из шаблона | 
**Variables** | [**[]IncludesIdName**](IncludesIdName.md) | Переменные доступные для вставки в шаблон | 

## Methods

### NewMailTemplatesMailTemplate

`func NewMailTemplatesMailTemplate(id string, name string, text string, type_ string, variables []IncludesIdName, ) *MailTemplatesMailTemplate`

NewMailTemplatesMailTemplate instantiates a new MailTemplatesMailTemplate object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMailTemplatesMailTemplateWithDefaults

`func NewMailTemplatesMailTemplateWithDefaults() *MailTemplatesMailTemplate`

NewMailTemplatesMailTemplateWithDefaults instantiates a new MailTemplatesMailTemplate object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

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


