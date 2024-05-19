# NegotiationsMessage

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

## Methods

### NewNegotiationsMessage

`func NewNegotiationsMessage(author NegotiationsAuthor, createdAt string, editable bool, id string, state IncludesIdName, text NullableString, viewedByMe bool, viewedByOpponent bool, ) *NegotiationsMessage`

NewNegotiationsMessage instantiates a new NegotiationsMessage object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNegotiationsMessageWithDefaults

`func NewNegotiationsMessageWithDefaults() *NegotiationsMessage`

NewNegotiationsMessageWithDefaults instantiates a new NegotiationsMessage object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAuthor

`func (o *NegotiationsMessage) GetAuthor() NegotiationsAuthor`

GetAuthor returns the Author field if non-nil, zero value otherwise.

### GetAuthorOk

`func (o *NegotiationsMessage) GetAuthorOk() (*NegotiationsAuthor, bool)`

GetAuthorOk returns a tuple with the Author field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthor

`func (o *NegotiationsMessage) SetAuthor(v NegotiationsAuthor)`

SetAuthor sets Author field to given value.


### GetCreatedAt

`func (o *NegotiationsMessage) GetCreatedAt() string`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *NegotiationsMessage) GetCreatedAtOk() (*string, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *NegotiationsMessage) SetCreatedAt(v string)`

SetCreatedAt sets CreatedAt field to given value.


### GetEditable

`func (o *NegotiationsMessage) GetEditable() bool`

GetEditable returns the Editable field if non-nil, zero value otherwise.

### GetEditableOk

`func (o *NegotiationsMessage) GetEditableOk() (*bool, bool)`

GetEditableOk returns a tuple with the Editable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEditable

`func (o *NegotiationsMessage) SetEditable(v bool)`

SetEditable sets Editable field to given value.


### GetId

`func (o *NegotiationsMessage) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *NegotiationsMessage) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *NegotiationsMessage) SetId(v string)`

SetId sets Id field to given value.


### GetRead

`func (o *NegotiationsMessage) GetRead() bool`

GetRead returns the Read field if non-nil, zero value otherwise.

### GetReadOk

`func (o *NegotiationsMessage) GetReadOk() (*bool, bool)`

GetReadOk returns a tuple with the Read field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRead

`func (o *NegotiationsMessage) SetRead(v bool)`

SetRead sets Read field to given value.

### HasRead

`func (o *NegotiationsMessage) HasRead() bool`

HasRead returns a boolean if a field has been set.

### GetState

`func (o *NegotiationsMessage) GetState() IncludesIdName`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *NegotiationsMessage) GetStateOk() (*IncludesIdName, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *NegotiationsMessage) SetState(v IncludesIdName)`

SetState sets State field to given value.


### GetText

`func (o *NegotiationsMessage) GetText() string`

GetText returns the Text field if non-nil, zero value otherwise.

### GetTextOk

`func (o *NegotiationsMessage) GetTextOk() (*string, bool)`

GetTextOk returns a tuple with the Text field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetText

`func (o *NegotiationsMessage) SetText(v string)`

SetText sets Text field to given value.


### SetTextNil

`func (o *NegotiationsMessage) SetTextNil(b bool)`

 SetTextNil sets the value for Text to be an explicit nil

### UnsetText
`func (o *NegotiationsMessage) UnsetText()`

UnsetText ensures that no value is present for Text, not even an explicit nil
### GetViewedByMe

`func (o *NegotiationsMessage) GetViewedByMe() bool`

GetViewedByMe returns the ViewedByMe field if non-nil, zero value otherwise.

### GetViewedByMeOk

`func (o *NegotiationsMessage) GetViewedByMeOk() (*bool, bool)`

GetViewedByMeOk returns a tuple with the ViewedByMe field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetViewedByMe

`func (o *NegotiationsMessage) SetViewedByMe(v bool)`

SetViewedByMe sets ViewedByMe field to given value.


### GetViewedByOpponent

`func (o *NegotiationsMessage) GetViewedByOpponent() bool`

GetViewedByOpponent returns the ViewedByOpponent field if non-nil, zero value otherwise.

### GetViewedByOpponentOk

`func (o *NegotiationsMessage) GetViewedByOpponentOk() (*bool, bool)`

GetViewedByOpponentOk returns a tuple with the ViewedByOpponent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetViewedByOpponent

`func (o *NegotiationsMessage) SetViewedByOpponent(v bool)`

SetViewedByOpponent sets ViewedByOpponent field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


