# ErrorsVacancyAddEditBadJsonDataError

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Description** | Pointer to **NullableString** | Описание ошибки | [optional] 
**Pointer** | Pointer to **string** | Путь до параметра, в котором возникла ошибка.  Для указания параметра используется формат JsonPointer [RFC 6901](https://tools.ietf.org/html/rfc6901)  | [optional] 
**Reason** | Pointer to **string** | Причина ошибки. Возможные значения:   * &#x60;required&#x60; - отстутствует поле в запросе   * &#x60;invalid&#x60; - недопустимое значение в поле запроса   * &#x60;is_empty&#x60; — пустое значение   * &#x60;wrong_size&#x60; — значение имеет неправильный размер   * &#x60;is_too_short&#x60; — значение имеет слишком маленький размер   * &#x60;is_too_long&#x60; — значение имеет слишком большой размер   * &#x60;currency_code_is_invalid&#x60; — валюта заработной платы введена некорректно   * &#x60;chosen_area_is_not_a_leaf_or_not_exist&#x60; — местоположение вакансии введено неверно (например, передан несуществующий ID) или не является конечным регионом (город, населенный пункт)   * &#x60;email_in_description&#x60; — в описании вакансии содержится email   * &#x60;anonymous_vacancy_contains_address&#x60; — в анонимной вакансии содержится адрес работодателя   * &#x60;anonymous_vacancy_has_real_company_name&#x60; — в названии вакансии содержится название компании работодателя   * &#x60;only_for_anonymous_type&#x60; — действие доступно только для анонимных вакансий   * &#x60;address_is_disabled&#x60; — адрес недоступен   * &#x60;vacancy_type_employer_billing_type_mismatch&#x60; — тип вакансии не совместим с текущим биллинг-типом   * &#x60;only_for_direct_type&#x60; — действие доступно только для прямых вакансий   * &#x60;address_is_empty_with_checked_show_metro_flag&#x60; — введен пустой адрес, но указана опция показывать метро   * &#x60;address_has_no_metro_but_checked_show_metro_flag&#x60; — по введенному адресу не доступно метро, но указана опция показывать метро   * &#x60;default_vacancy_branded_template_is_invalid_or_not_enough_purchased_services&#x60; — в запросе указан шаблон, который отсутствует в списке доступных шаблонов (этот список можно получить [запросом](#tag/Informaciya-o-rabotodatele/operation/get-vacancy-branded-templates-list)). Также шаблон может отсутствовать в списке доступных шаблонов, если не оплачена услуга использования [брендированного шаблона вакансии](https://hh.ru/price/branding)   * &#x60;department_code_prohibited_in_anonymous_vacancy&#x60; — нельзя указать код подразделения для анонимной вакансии   * &#x60;branded_template_prohibited_in_anonymous_vacancy&#x60; — использование брендированного шаблона невозможно для анонимной вакансии   * &#x60;value_conflict_with_business_rules&#x60; — публикация вакансии с указанным &#x60;billing_type&#x60; запрещена   * &#x60;can_not_accept_kids&#x60; — вакансия недоступна несовершеннолетним  | [optional] 
**Type** | **string** | Текстовый идентификатор типа ошибки | 
**Value** | Pointer to **string** | Название поля с ошибкой | [optional] 

## Methods

### NewErrorsVacancyAddEditBadJsonDataError

`func NewErrorsVacancyAddEditBadJsonDataError(type_ string, ) *ErrorsVacancyAddEditBadJsonDataError`

NewErrorsVacancyAddEditBadJsonDataError instantiates a new ErrorsVacancyAddEditBadJsonDataError object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewErrorsVacancyAddEditBadJsonDataErrorWithDefaults

`func NewErrorsVacancyAddEditBadJsonDataErrorWithDefaults() *ErrorsVacancyAddEditBadJsonDataError`

NewErrorsVacancyAddEditBadJsonDataErrorWithDefaults instantiates a new ErrorsVacancyAddEditBadJsonDataError object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDescription

`func (o *ErrorsVacancyAddEditBadJsonDataError) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *ErrorsVacancyAddEditBadJsonDataError) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *ErrorsVacancyAddEditBadJsonDataError) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *ErrorsVacancyAddEditBadJsonDataError) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescriptionNil

`func (o *ErrorsVacancyAddEditBadJsonDataError) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *ErrorsVacancyAddEditBadJsonDataError) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetPointer

`func (o *ErrorsVacancyAddEditBadJsonDataError) GetPointer() string`

GetPointer returns the Pointer field if non-nil, zero value otherwise.

### GetPointerOk

`func (o *ErrorsVacancyAddEditBadJsonDataError) GetPointerOk() (*string, bool)`

GetPointerOk returns a tuple with the Pointer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPointer

`func (o *ErrorsVacancyAddEditBadJsonDataError) SetPointer(v string)`

SetPointer sets Pointer field to given value.

### HasPointer

`func (o *ErrorsVacancyAddEditBadJsonDataError) HasPointer() bool`

HasPointer returns a boolean if a field has been set.

### GetReason

`func (o *ErrorsVacancyAddEditBadJsonDataError) GetReason() string`

GetReason returns the Reason field if non-nil, zero value otherwise.

### GetReasonOk

`func (o *ErrorsVacancyAddEditBadJsonDataError) GetReasonOk() (*string, bool)`

GetReasonOk returns a tuple with the Reason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReason

`func (o *ErrorsVacancyAddEditBadJsonDataError) SetReason(v string)`

SetReason sets Reason field to given value.

### HasReason

`func (o *ErrorsVacancyAddEditBadJsonDataError) HasReason() bool`

HasReason returns a boolean if a field has been set.

### GetType

`func (o *ErrorsVacancyAddEditBadJsonDataError) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *ErrorsVacancyAddEditBadJsonDataError) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *ErrorsVacancyAddEditBadJsonDataError) SetType(v string)`

SetType sets Type field to given value.


### GetValue

`func (o *ErrorsVacancyAddEditBadJsonDataError) GetValue() string`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *ErrorsVacancyAddEditBadJsonDataError) GetValueOk() (*string, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *ErrorsVacancyAddEditBadJsonDataError) SetValue(v string)`

SetValue sets Value field to given value.

### HasValue

`func (o *ErrorsVacancyAddEditBadJsonDataError) HasValue() bool`

HasValue returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


