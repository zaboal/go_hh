# CredsAnswers

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AnswerGroup** | Pointer to **string** | Группа данного ответа, positive, negative, neutral | [optional] 
**AnswerId** | Pointer to **string** | Идентификатор ответа (совпадает с ключом объекта) | [optional] 
**AskQuestionsAfter** | Pointer to **[]string** | Вопросы которые нужно задать после использования пользователем данного ответа | [optional] 
**Description** | Pointer to **NullableString** | Описание ответа | [optional] 
**PositiveTitle** | Pointer to **string** | Текст ответа который можно использовать для отображения без самого вопроса | [optional] 
**SkipAtResult** | Pointer to **bool** | Нужно ли пропускать данный ответ на форме с отображением кредов пользователя | [optional] 
**Title** | Pointer to **string** | Текст ответа который нужно отрисовать для сбора ответов от пользователя | [optional] 

## Methods

### NewCredsAnswers

`func NewCredsAnswers() *CredsAnswers`

NewCredsAnswers instantiates a new CredsAnswers object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCredsAnswersWithDefaults

`func NewCredsAnswersWithDefaults() *CredsAnswers`

NewCredsAnswersWithDefaults instantiates a new CredsAnswers object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAnswerGroup

`func (o *CredsAnswers) GetAnswerGroup() string`

GetAnswerGroup returns the AnswerGroup field if non-nil, zero value otherwise.

### GetAnswerGroupOk

`func (o *CredsAnswers) GetAnswerGroupOk() (*string, bool)`

GetAnswerGroupOk returns a tuple with the AnswerGroup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAnswerGroup

`func (o *CredsAnswers) SetAnswerGroup(v string)`

SetAnswerGroup sets AnswerGroup field to given value.

### HasAnswerGroup

`func (o *CredsAnswers) HasAnswerGroup() bool`

HasAnswerGroup returns a boolean if a field has been set.

### GetAnswerId

`func (o *CredsAnswers) GetAnswerId() string`

GetAnswerId returns the AnswerId field if non-nil, zero value otherwise.

### GetAnswerIdOk

`func (o *CredsAnswers) GetAnswerIdOk() (*string, bool)`

GetAnswerIdOk returns a tuple with the AnswerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAnswerId

`func (o *CredsAnswers) SetAnswerId(v string)`

SetAnswerId sets AnswerId field to given value.

### HasAnswerId

`func (o *CredsAnswers) HasAnswerId() bool`

HasAnswerId returns a boolean if a field has been set.

### GetAskQuestionsAfter

`func (o *CredsAnswers) GetAskQuestionsAfter() []string`

GetAskQuestionsAfter returns the AskQuestionsAfter field if non-nil, zero value otherwise.

### GetAskQuestionsAfterOk

`func (o *CredsAnswers) GetAskQuestionsAfterOk() (*[]string, bool)`

GetAskQuestionsAfterOk returns a tuple with the AskQuestionsAfter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAskQuestionsAfter

`func (o *CredsAnswers) SetAskQuestionsAfter(v []string)`

SetAskQuestionsAfter sets AskQuestionsAfter field to given value.

### HasAskQuestionsAfter

`func (o *CredsAnswers) HasAskQuestionsAfter() bool`

HasAskQuestionsAfter returns a boolean if a field has been set.

### GetDescription

`func (o *CredsAnswers) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *CredsAnswers) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *CredsAnswers) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *CredsAnswers) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescriptionNil

`func (o *CredsAnswers) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *CredsAnswers) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetPositiveTitle

`func (o *CredsAnswers) GetPositiveTitle() string`

GetPositiveTitle returns the PositiveTitle field if non-nil, zero value otherwise.

### GetPositiveTitleOk

`func (o *CredsAnswers) GetPositiveTitleOk() (*string, bool)`

GetPositiveTitleOk returns a tuple with the PositiveTitle field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPositiveTitle

`func (o *CredsAnswers) SetPositiveTitle(v string)`

SetPositiveTitle sets PositiveTitle field to given value.

### HasPositiveTitle

`func (o *CredsAnswers) HasPositiveTitle() bool`

HasPositiveTitle returns a boolean if a field has been set.

### GetSkipAtResult

`func (o *CredsAnswers) GetSkipAtResult() bool`

GetSkipAtResult returns the SkipAtResult field if non-nil, zero value otherwise.

### GetSkipAtResultOk

`func (o *CredsAnswers) GetSkipAtResultOk() (*bool, bool)`

GetSkipAtResultOk returns a tuple with the SkipAtResult field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSkipAtResult

`func (o *CredsAnswers) SetSkipAtResult(v bool)`

SetSkipAtResult sets SkipAtResult field to given value.

### HasSkipAtResult

`func (o *CredsAnswers) HasSkipAtResult() bool`

HasSkipAtResult returns a boolean if a field has been set.

### GetTitle

`func (o *CredsAnswers) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *CredsAnswers) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *CredsAnswers) SetTitle(v string)`

SetTitle sets Title field to given value.

### HasTitle

`func (o *CredsAnswers) HasTitle() bool`

HasTitle returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


