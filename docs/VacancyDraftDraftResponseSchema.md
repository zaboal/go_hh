# VacancyDraftDraftResponseSchema

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | Идентификатор созданного объекта | 
**IgnoredFields** | Pointer to **[]string** | Поля, которые не были сохранены при создании черновика, вследствие не верного заполнения | [optional] 
**Name** | **string** | Заголовок черновика | 
**PublicationReady** | **bool** | Готовность черновика к публикации | 
**ValidationErrors** | Pointer to [**[]VacancyDraftDraftVacancyError**](VacancyDraftDraftVacancyError.md) | Поля, которые были сохранены при создании черновика, с соответствующими ошибками, которые необходимо поправить для успешной публикации вакансии на основе черновика | [optional] 

## Methods

### NewVacancyDraftDraftResponseSchema

`func NewVacancyDraftDraftResponseSchema(id string, name string, publicationReady bool, ) *VacancyDraftDraftResponseSchema`

NewVacancyDraftDraftResponseSchema instantiates a new VacancyDraftDraftResponseSchema object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVacancyDraftDraftResponseSchemaWithDefaults

`func NewVacancyDraftDraftResponseSchemaWithDefaults() *VacancyDraftDraftResponseSchema`

NewVacancyDraftDraftResponseSchemaWithDefaults instantiates a new VacancyDraftDraftResponseSchema object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *VacancyDraftDraftResponseSchema) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *VacancyDraftDraftResponseSchema) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *VacancyDraftDraftResponseSchema) SetId(v string)`

SetId sets Id field to given value.


### GetIgnoredFields

`func (o *VacancyDraftDraftResponseSchema) GetIgnoredFields() []string`

GetIgnoredFields returns the IgnoredFields field if non-nil, zero value otherwise.

### GetIgnoredFieldsOk

`func (o *VacancyDraftDraftResponseSchema) GetIgnoredFieldsOk() (*[]string, bool)`

GetIgnoredFieldsOk returns a tuple with the IgnoredFields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIgnoredFields

`func (o *VacancyDraftDraftResponseSchema) SetIgnoredFields(v []string)`

SetIgnoredFields sets IgnoredFields field to given value.

### HasIgnoredFields

`func (o *VacancyDraftDraftResponseSchema) HasIgnoredFields() bool`

HasIgnoredFields returns a boolean if a field has been set.

### GetName

`func (o *VacancyDraftDraftResponseSchema) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *VacancyDraftDraftResponseSchema) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *VacancyDraftDraftResponseSchema) SetName(v string)`

SetName sets Name field to given value.


### GetPublicationReady

`func (o *VacancyDraftDraftResponseSchema) GetPublicationReady() bool`

GetPublicationReady returns the PublicationReady field if non-nil, zero value otherwise.

### GetPublicationReadyOk

`func (o *VacancyDraftDraftResponseSchema) GetPublicationReadyOk() (*bool, bool)`

GetPublicationReadyOk returns a tuple with the PublicationReady field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPublicationReady

`func (o *VacancyDraftDraftResponseSchema) SetPublicationReady(v bool)`

SetPublicationReady sets PublicationReady field to given value.


### GetValidationErrors

`func (o *VacancyDraftDraftResponseSchema) GetValidationErrors() []VacancyDraftDraftVacancyError`

GetValidationErrors returns the ValidationErrors field if non-nil, zero value otherwise.

### GetValidationErrorsOk

`func (o *VacancyDraftDraftResponseSchema) GetValidationErrorsOk() (*[]VacancyDraftDraftVacancyError, bool)`

GetValidationErrorsOk returns a tuple with the ValidationErrors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValidationErrors

`func (o *VacancyDraftDraftResponseSchema) SetValidationErrors(v []VacancyDraftDraftVacancyError)`

SetValidationErrors sets ValidationErrors field to given value.

### HasValidationErrors

`func (o *VacancyDraftDraftResponseSchema) HasValidationErrors() bool`

HasValidationErrors returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


