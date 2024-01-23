# ApplicantCommentsApplicantCommentItem

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccessType** | [**ApplicantCommentsApplicantCommentItemAccessType**](ApplicantCommentsApplicantCommentItemAccessType.md) |  | 
**Author** | [**ApplicantCommentsApplicantCommentItemAuthor**](ApplicantCommentsApplicantCommentItemAuthor.md) |  | 
**CreatedAt** | **string** | Дата создания комментария | 
**Id** | **string** | Уникальный идентификатор комментария | 
**IsMine** | **bool** | Комментарий написан текущим пользователем? | 
**Text** | **string** | Текст комментария. Может содержать символы новой строки | 

## Methods

### NewApplicantCommentsApplicantCommentItem

`func NewApplicantCommentsApplicantCommentItem(accessType ApplicantCommentsApplicantCommentItemAccessType, author ApplicantCommentsApplicantCommentItemAuthor, createdAt string, id string, isMine bool, text string, ) *ApplicantCommentsApplicantCommentItem`

NewApplicantCommentsApplicantCommentItem instantiates a new ApplicantCommentsApplicantCommentItem object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewApplicantCommentsApplicantCommentItemWithDefaults

`func NewApplicantCommentsApplicantCommentItemWithDefaults() *ApplicantCommentsApplicantCommentItem`

NewApplicantCommentsApplicantCommentItemWithDefaults instantiates a new ApplicantCommentsApplicantCommentItem object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccessType

`func (o *ApplicantCommentsApplicantCommentItem) GetAccessType() ApplicantCommentsApplicantCommentItemAccessType`

GetAccessType returns the AccessType field if non-nil, zero value otherwise.

### GetAccessTypeOk

`func (o *ApplicantCommentsApplicantCommentItem) GetAccessTypeOk() (*ApplicantCommentsApplicantCommentItemAccessType, bool)`

GetAccessTypeOk returns a tuple with the AccessType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccessType

`func (o *ApplicantCommentsApplicantCommentItem) SetAccessType(v ApplicantCommentsApplicantCommentItemAccessType)`

SetAccessType sets AccessType field to given value.


### GetAuthor

`func (o *ApplicantCommentsApplicantCommentItem) GetAuthor() ApplicantCommentsApplicantCommentItemAuthor`

GetAuthor returns the Author field if non-nil, zero value otherwise.

### GetAuthorOk

`func (o *ApplicantCommentsApplicantCommentItem) GetAuthorOk() (*ApplicantCommentsApplicantCommentItemAuthor, bool)`

GetAuthorOk returns a tuple with the Author field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthor

`func (o *ApplicantCommentsApplicantCommentItem) SetAuthor(v ApplicantCommentsApplicantCommentItemAuthor)`

SetAuthor sets Author field to given value.


### GetCreatedAt

`func (o *ApplicantCommentsApplicantCommentItem) GetCreatedAt() string`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *ApplicantCommentsApplicantCommentItem) GetCreatedAtOk() (*string, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *ApplicantCommentsApplicantCommentItem) SetCreatedAt(v string)`

SetCreatedAt sets CreatedAt field to given value.


### GetId

`func (o *ApplicantCommentsApplicantCommentItem) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ApplicantCommentsApplicantCommentItem) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ApplicantCommentsApplicantCommentItem) SetId(v string)`

SetId sets Id field to given value.


### GetIsMine

`func (o *ApplicantCommentsApplicantCommentItem) GetIsMine() bool`

GetIsMine returns the IsMine field if non-nil, zero value otherwise.

### GetIsMineOk

`func (o *ApplicantCommentsApplicantCommentItem) GetIsMineOk() (*bool, bool)`

GetIsMineOk returns a tuple with the IsMine field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsMine

`func (o *ApplicantCommentsApplicantCommentItem) SetIsMine(v bool)`

SetIsMine sets IsMine field to given value.


### GetText

`func (o *ApplicantCommentsApplicantCommentItem) GetText() string`

GetText returns the Text field if non-nil, zero value otherwise.

### GetTextOk

`func (o *ApplicantCommentsApplicantCommentItem) GetTextOk() (*string, bool)`

GetTextOk returns a tuple with the Text field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetText

`func (o *ApplicantCommentsApplicantCommentItem) SetText(v string)`

SetText sets Text field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


