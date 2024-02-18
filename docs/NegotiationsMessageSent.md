# NegotiationsMessageSent

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Address** | [**NullableVacancyAddressOutput**](VacancyAddressOutput.md) |  | 
**Author** | [**NegotiationsAuthor**](NegotiationsAuthor.md) |  | 
**CreatedAt** | **string** | Дата и время создания сообщения | 
**Editable** | **bool** | Можно ли редактировать текст сообщения | 
**Id** | **string** | Идентификатор сообщения | 
**Read** | Pointer to **bool** | Можно ли прочитать сообщение | [optional] 
**State** | [**IncludesIdName**](IncludesIdName.md) | Состояние сообщения | 
**Text** | **string** | Текст сообщения | 
**ViewedByMe** | **bool** | Прочитано ли сообщение смотрящим (для сообщений отправленных соискателем - всегда true) | 
**ViewedByOpponent** | **bool** | Прочитано ли сообщение работодателем (для сообщений работодателя - true) | 

## Methods

### NewNegotiationsMessageSent

`func NewNegotiationsMessageSent(address NullableVacancyAddressOutput, author NegotiationsAuthor, createdAt string, editable bool, id string, state IncludesIdName, text string, viewedByMe bool, viewedByOpponent bool, ) *NegotiationsMessageSent`

NewNegotiationsMessageSent instantiates a new NegotiationsMessageSent object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNegotiationsMessageSentWithDefaults

`func NewNegotiationsMessageSentWithDefaults() *NegotiationsMessageSent`

NewNegotiationsMessageSentWithDefaults instantiates a new NegotiationsMessageSent object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAddress

`func (o *NegotiationsMessageSent) GetAddress() VacancyAddressOutput`

GetAddress returns the Address field if non-nil, zero value otherwise.

### GetAddressOk

`func (o *NegotiationsMessageSent) GetAddressOk() (*VacancyAddressOutput, bool)`

GetAddressOk returns a tuple with the Address field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddress

`func (o *NegotiationsMessageSent) SetAddress(v VacancyAddressOutput)`

SetAddress sets Address field to given value.


### SetAddressNil

`func (o *NegotiationsMessageSent) SetAddressNil(b bool)`

 SetAddressNil sets the value for Address to be an explicit nil

### UnsetAddress
`func (o *NegotiationsMessageSent) UnsetAddress()`

UnsetAddress ensures that no value is present for Address, not even an explicit nil
### GetAuthor

`func (o *NegotiationsMessageSent) GetAuthor() NegotiationsAuthor`

GetAuthor returns the Author field if non-nil, zero value otherwise.

### GetAuthorOk

`func (o *NegotiationsMessageSent) GetAuthorOk() (*NegotiationsAuthor, bool)`

GetAuthorOk returns a tuple with the Author field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthor

`func (o *NegotiationsMessageSent) SetAuthor(v NegotiationsAuthor)`

SetAuthor sets Author field to given value.


### GetCreatedAt

`func (o *NegotiationsMessageSent) GetCreatedAt() string`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *NegotiationsMessageSent) GetCreatedAtOk() (*string, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *NegotiationsMessageSent) SetCreatedAt(v string)`

SetCreatedAt sets CreatedAt field to given value.


### GetEditable

`func (o *NegotiationsMessageSent) GetEditable() bool`

GetEditable returns the Editable field if non-nil, zero value otherwise.

### GetEditableOk

`func (o *NegotiationsMessageSent) GetEditableOk() (*bool, bool)`

GetEditableOk returns a tuple with the Editable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEditable

`func (o *NegotiationsMessageSent) SetEditable(v bool)`

SetEditable sets Editable field to given value.


### GetId

`func (o *NegotiationsMessageSent) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *NegotiationsMessageSent) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *NegotiationsMessageSent) SetId(v string)`

SetId sets Id field to given value.


### GetRead

`func (o *NegotiationsMessageSent) GetRead() bool`

GetRead returns the Read field if non-nil, zero value otherwise.

### GetReadOk

`func (o *NegotiationsMessageSent) GetReadOk() (*bool, bool)`

GetReadOk returns a tuple with the Read field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRead

`func (o *NegotiationsMessageSent) SetRead(v bool)`

SetRead sets Read field to given value.

### HasRead

`func (o *NegotiationsMessageSent) HasRead() bool`

HasRead returns a boolean if a field has been set.

### GetState

`func (o *NegotiationsMessageSent) GetState() IncludesIdName`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *NegotiationsMessageSent) GetStateOk() (*IncludesIdName, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *NegotiationsMessageSent) SetState(v IncludesIdName)`

SetState sets State field to given value.


### GetText

`func (o *NegotiationsMessageSent) GetText() string`

GetText returns the Text field if non-nil, zero value otherwise.

### GetTextOk

`func (o *NegotiationsMessageSent) GetTextOk() (*string, bool)`

GetTextOk returns a tuple with the Text field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetText

`func (o *NegotiationsMessageSent) SetText(v string)`

SetText sets Text field to given value.


### GetViewedByMe

`func (o *NegotiationsMessageSent) GetViewedByMe() bool`

GetViewedByMe returns the ViewedByMe field if non-nil, zero value otherwise.

### GetViewedByMeOk

`func (o *NegotiationsMessageSent) GetViewedByMeOk() (*bool, bool)`

GetViewedByMeOk returns a tuple with the ViewedByMe field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetViewedByMe

`func (o *NegotiationsMessageSent) SetViewedByMe(v bool)`

SetViewedByMe sets ViewedByMe field to given value.


### GetViewedByOpponent

`func (o *NegotiationsMessageSent) GetViewedByOpponent() bool`

GetViewedByOpponent returns the ViewedByOpponent field if non-nil, zero value otherwise.

### GetViewedByOpponentOk

`func (o *NegotiationsMessageSent) GetViewedByOpponentOk() (*bool, bool)`

GetViewedByOpponentOk returns a tuple with the ViewedByOpponent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetViewedByOpponent

`func (o *NegotiationsMessageSent) SetViewedByOpponent(v bool)`

SetViewedByOpponent sets ViewedByOpponent field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


