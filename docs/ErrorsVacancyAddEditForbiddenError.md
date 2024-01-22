# ErrorsVacancyAddEditForbiddenError

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Found** | Pointer to **NullableFloat32** | Общее количество дубликатов вакансии. Возвращается только для &#x60;\&quot;value\&quot;: \&quot;duplicate\&quot;&#x60;  | [optional] 
**Items** | Pointer to [**[]IncludesNumericId**](IncludesNumericId.md) | Ограниченное количество записей с информацией о дубликатах. Не гарантирует выдачу всех дубликатов. Возвращается только для &#x60;\&quot;value\&quot;: \&quot;duplicate\&quot;&#x60;  | [optional] 
**Type** | **string** | Текстовый идентификатор типа ошибки | 
**Value** | **string** | Ошибки при публикации и редактировании вакансии:   * &#x60;not_enough_purchased_services&#x60; — купленных услуг недостаточно для публикации или обновления данного типа вакансии   * &#x60;quota_exceeded&#x60; — квота менеджера на публикацию данного типа вакансии закончилась   * &#x60;duplicate&#x60; — аналогичная вакансия уже опубликована. В ответе передается информация по дубликатам вакансии. Данную ошибку можно форсировано отключить параметром &#x60;?ignore_duplicates&#x3D;true&#x60;   * &#x60;creation_forbidden&#x60; — публикация вакансий недоступна текущему менеджеру   * &#x60;unavailable_for_archived&#x60; — редактирование недоступно для архивной вакансии   * &#x60;conflict_changes&#x60; — [конфликтные изменения](https://github.com/hhru/api/blob/master/docs/employer_vacancies.md#%D1%81%D0%BC%D0%B5%D0%BD%D0%B0-%D0%B1%D0%B8%D0%BB%D0%BB%D0%B8%D0%BD%D0%B3-%D1%82%D0%B8%D0%BF%D0%B0-%D0%BC%D0%B5%D0%BD%D0%B5%D0%B4%D0%B6%D0%B5%D1%80%D0%B0-%D0%B2%D0%B0%D0%BA%D0%B0%D0%BD%D1%81%D0%B8%D0%B8) данных вакансии  | 

## Methods

### NewErrorsVacancyAddEditForbiddenError

`func NewErrorsVacancyAddEditForbiddenError(type_ string, value string, ) *ErrorsVacancyAddEditForbiddenError`

NewErrorsVacancyAddEditForbiddenError instantiates a new ErrorsVacancyAddEditForbiddenError object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewErrorsVacancyAddEditForbiddenErrorWithDefaults

`func NewErrorsVacancyAddEditForbiddenErrorWithDefaults() *ErrorsVacancyAddEditForbiddenError`

NewErrorsVacancyAddEditForbiddenErrorWithDefaults instantiates a new ErrorsVacancyAddEditForbiddenError object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFound

`func (o *ErrorsVacancyAddEditForbiddenError) GetFound() float32`

GetFound returns the Found field if non-nil, zero value otherwise.

### GetFoundOk

`func (o *ErrorsVacancyAddEditForbiddenError) GetFoundOk() (*float32, bool)`

GetFoundOk returns a tuple with the Found field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFound

`func (o *ErrorsVacancyAddEditForbiddenError) SetFound(v float32)`

SetFound sets Found field to given value.

### HasFound

`func (o *ErrorsVacancyAddEditForbiddenError) HasFound() bool`

HasFound returns a boolean if a field has been set.

### SetFoundNil

`func (o *ErrorsVacancyAddEditForbiddenError) SetFoundNil(b bool)`

 SetFoundNil sets the value for Found to be an explicit nil

### UnsetFound
`func (o *ErrorsVacancyAddEditForbiddenError) UnsetFound()`

UnsetFound ensures that no value is present for Found, not even an explicit nil
### GetItems

`func (o *ErrorsVacancyAddEditForbiddenError) GetItems() []IncludesNumericId`

GetItems returns the Items field if non-nil, zero value otherwise.

### GetItemsOk

`func (o *ErrorsVacancyAddEditForbiddenError) GetItemsOk() (*[]IncludesNumericId, bool)`

GetItemsOk returns a tuple with the Items field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItems

`func (o *ErrorsVacancyAddEditForbiddenError) SetItems(v []IncludesNumericId)`

SetItems sets Items field to given value.

### HasItems

`func (o *ErrorsVacancyAddEditForbiddenError) HasItems() bool`

HasItems returns a boolean if a field has been set.

### SetItemsNil

`func (o *ErrorsVacancyAddEditForbiddenError) SetItemsNil(b bool)`

 SetItemsNil sets the value for Items to be an explicit nil

### UnsetItems
`func (o *ErrorsVacancyAddEditForbiddenError) UnsetItems()`

UnsetItems ensures that no value is present for Items, not even an explicit nil
### GetType

`func (o *ErrorsVacancyAddEditForbiddenError) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *ErrorsVacancyAddEditForbiddenError) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *ErrorsVacancyAddEditForbiddenError) SetType(v string)`

SetType sets Type field to given value.


### GetValue

`func (o *ErrorsVacancyAddEditForbiddenError) GetValue() string`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *ErrorsVacancyAddEditForbiddenError) GetValueOk() (*string, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *ErrorsVacancyAddEditForbiddenError) SetValue(v string)`

SetValue sets Value field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


