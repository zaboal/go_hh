# NegotiationsMessagesGet

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Author** | [**NegotiationsAuthor**](NegotiationsAuthor.md) |  | 
**CreatedAt** | **string** | Дата и время создания сообщения | 
**Editable** | **bool** | Можно ли редактировать текст сообщения | 
**Id** | **string** | Идентификатор сообщения | 
**Read** | Pointer to **bool** | Можно ли прочитать сообщение | [optional] 
**State** | [**IncludesIdName**](IncludesIdName.md) | Состояние сообщения | 
**Text** | **NullableString** | Текст сообщения | 
**ViewedByMe** | **bool** | Прочитано ли сообщение смотрящим (для сообщений отправленных соискателем - всегда true) | 
**ViewedByOpponent** | **bool** | Прочитано ли сообщение работодателем (для сообщений работодателя - true) | 
**Address** | [**NullableVacancyAddressRawOutput**](VacancyAddressRawOutput.md) |  | 
**Assessments** | Pointer to [**[]NegotiationsAssessment**](NegotiationsAssessment.md) | Инструменты оценки, привязанные к сообщению | [optional] 

## Methods

### NewNegotiationsMessagesGet

`func NewNegotiationsMessagesGet(author NegotiationsAuthor, createdAt string, editable bool, id string, state IncludesIdName, text NullableString, viewedByMe bool, viewedByOpponent bool, address NullableVacancyAddressRawOutput, ) *NegotiationsMessagesGet`

NewNegotiationsMessagesGet instantiates a new NegotiationsMessagesGet object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNegotiationsMessagesGetWithDefaults

`func NewNegotiationsMessagesGetWithDefaults() *NegotiationsMessagesGet`

NewNegotiationsMessagesGetWithDefaults instantiates a new NegotiationsMessagesGet object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAuthor

`func (o *NegotiationsMessagesGet) GetAuthor() NegotiationsAuthor`

GetAuthor returns the Author field if non-nil, zero value otherwise.

### GetAuthorOk

`func (o *NegotiationsMessagesGet) GetAuthorOk() (*NegotiationsAuthor, bool)`

GetAuthorOk returns a tuple with the Author field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthor

`func (o *NegotiationsMessagesGet) SetAuthor(v NegotiationsAuthor)`

SetAuthor sets Author field to given value.


### GetCreatedAt

`func (o *NegotiationsMessagesGet) GetCreatedAt() string`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *NegotiationsMessagesGet) GetCreatedAtOk() (*string, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *NegotiationsMessagesGet) SetCreatedAt(v string)`

SetCreatedAt sets CreatedAt field to given value.


### GetEditable

`func (o *NegotiationsMessagesGet) GetEditable() bool`

GetEditable returns the Editable field if non-nil, zero value otherwise.

### GetEditableOk

`func (o *NegotiationsMessagesGet) GetEditableOk() (*bool, bool)`

GetEditableOk returns a tuple with the Editable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEditable

`func (o *NegotiationsMessagesGet) SetEditable(v bool)`

SetEditable sets Editable field to given value.


### GetId

`func (o *NegotiationsMessagesGet) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *NegotiationsMessagesGet) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *NegotiationsMessagesGet) SetId(v string)`

SetId sets Id field to given value.


### GetRead

`func (o *NegotiationsMessagesGet) GetRead() bool`

GetRead returns the Read field if non-nil, zero value otherwise.

### GetReadOk

`func (o *NegotiationsMessagesGet) GetReadOk() (*bool, bool)`

GetReadOk returns a tuple with the Read field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRead

`func (o *NegotiationsMessagesGet) SetRead(v bool)`

SetRead sets Read field to given value.

### HasRead

`func (o *NegotiationsMessagesGet) HasRead() bool`

HasRead returns a boolean if a field has been set.

### GetState

`func (o *NegotiationsMessagesGet) GetState() IncludesIdName`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *NegotiationsMessagesGet) GetStateOk() (*IncludesIdName, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *NegotiationsMessagesGet) SetState(v IncludesIdName)`

SetState sets State field to given value.


### GetText

`func (o *NegotiationsMessagesGet) GetText() string`

GetText returns the Text field if non-nil, zero value otherwise.

### GetTextOk

`func (o *NegotiationsMessagesGet) GetTextOk() (*string, bool)`

GetTextOk returns a tuple with the Text field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetText

`func (o *NegotiationsMessagesGet) SetText(v string)`

SetText sets Text field to given value.


### SetTextNil

`func (o *NegotiationsMessagesGet) SetTextNil(b bool)`

 SetTextNil sets the value for Text to be an explicit nil

### UnsetText
`func (o *NegotiationsMessagesGet) UnsetText()`

UnsetText ensures that no value is present for Text, not even an explicit nil
### GetViewedByMe

`func (o *NegotiationsMessagesGet) GetViewedByMe() bool`

GetViewedByMe returns the ViewedByMe field if non-nil, zero value otherwise.

### GetViewedByMeOk

`func (o *NegotiationsMessagesGet) GetViewedByMeOk() (*bool, bool)`

GetViewedByMeOk returns a tuple with the ViewedByMe field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetViewedByMe

`func (o *NegotiationsMessagesGet) SetViewedByMe(v bool)`

SetViewedByMe sets ViewedByMe field to given value.


### GetViewedByOpponent

`func (o *NegotiationsMessagesGet) GetViewedByOpponent() bool`

GetViewedByOpponent returns the ViewedByOpponent field if non-nil, zero value otherwise.

### GetViewedByOpponentOk

`func (o *NegotiationsMessagesGet) GetViewedByOpponentOk() (*bool, bool)`

GetViewedByOpponentOk returns a tuple with the ViewedByOpponent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetViewedByOpponent

`func (o *NegotiationsMessagesGet) SetViewedByOpponent(v bool)`

SetViewedByOpponent sets ViewedByOpponent field to given value.


### GetAddress

`func (o *NegotiationsMessagesGet) GetAddress() VacancyAddressRawOutput`

GetAddress returns the Address field if non-nil, zero value otherwise.

### GetAddressOk

`func (o *NegotiationsMessagesGet) GetAddressOk() (*VacancyAddressRawOutput, bool)`

GetAddressOk returns a tuple with the Address field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddress

`func (o *NegotiationsMessagesGet) SetAddress(v VacancyAddressRawOutput)`

SetAddress sets Address field to given value.


### SetAddressNil

`func (o *NegotiationsMessagesGet) SetAddressNil(b bool)`

 SetAddressNil sets the value for Address to be an explicit nil

### UnsetAddress
`func (o *NegotiationsMessagesGet) UnsetAddress()`

UnsetAddress ensures that no value is present for Address, not even an explicit nil
### GetAssessments

`func (o *NegotiationsMessagesGet) GetAssessments() []NegotiationsAssessment`

GetAssessments returns the Assessments field if non-nil, zero value otherwise.

### GetAssessmentsOk

`func (o *NegotiationsMessagesGet) GetAssessmentsOk() (*[]NegotiationsAssessment, bool)`

GetAssessmentsOk returns a tuple with the Assessments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssessments

`func (o *NegotiationsMessagesGet) SetAssessments(v []NegotiationsAssessment)`

SetAssessments sets Assessments field to given value.

### HasAssessments

`func (o *NegotiationsMessagesGet) HasAssessments() bool`

HasAssessments returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


