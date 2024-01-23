# ErrorsResumeAddEditError

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Description** | **string** | Описание ошибки для пользователя | 
**Pointer** | **string** | Путь до параметра, в котором возникла ошибка.  Для указания параметра используется формат JsonPointer [RFC 6901](https://tools.ietf.org/html/rfc6901)  | 
**Reason** | **string** | Причина ошибки. Возможные значения:    * &#x60;required&#x60; — поле является обязательным для заполнения.   * &#x60;not_found&#x60; — не найдено значение по переданному ID.   * &#x60;faculty_without_university&#x60; — нельзя установить факультет без университета.   * &#x60;not_in_dictionary&#x60; — не найдено значение по переданному ID в справочнике.   * &#x60;not_a_leaf&#x60; — значение не должно содержать потомков.   * &#x60;end_date_before_start_date&#x60; — значение &#x60;end&#x60; меньше &#x60;start&#x60;.   * &#x60;not_country&#x60; — значение &#x60;area&#x60; должно быть страной (см. [справочник стран](#tag/Obshie-spravochniki/operation/get-countries)).   * &#x60;more_than_one_native_language&#x60; — указано более одного родного языка.   * &#x60;must_contain_unique&#x60; — переданные значения должны быть уникальны.   * &#x60;from_different_profareas&#x60; — переданы значения из разных отраслей.   * &#x60;duplicate&#x60; — значение уже было использовано.   * &#x60;bad_image_type&#x60; — передано значение изображения неправильного типа (для &#x60;portfolio&#x60; необходимы значения из [GET /artifacts/portfolio](#tag/Rabota-s-artefaktami/operation/get-artifacts-portfolio), для photo — [GET /artifacts/photo](#tag/Rabota-s-artefaktami/operation/get-artifact-photos)) .   * &#x60;processing&#x60; — объект в процессе обработки.   * &#x60;preferred_must_be_unique&#x60; — предпочитаемый тип связи должен быть уникальным.   * &#x60;preferred_contact_not_specified&#x60; — предпочитаемый тип связи не указан или не указано значение контакта.   * &#x60;need_country_city_number_or_formatted&#x60; — телефон в контактах указан в неверном формате (см. [условия заполнения контактов в резюме](#tag/Rezyume.-Usloviya-zapolneniya-polej/operation/get-new-resume-conditions)).   * &#x60;invalid&#x60; — ошибка в значении поля (поля должны соответствовать [условиям заполнения](#tag/Rezyume.-Usloviya-zapolneniya-polej/operation/get-resume-conditions)).   * &#x60;greater_than_max&#x60; — значение больше максимума .   * &#x60;less_than_min&#x60; — значение меньше минимума.   * &#x60;earlier_than_min&#x60; — указанная дата раньше минимально возможной.   * &#x60;later_than_max&#x60; — указанная дата позже максимально возможной.   * &#x60;length_less_than_min&#x60; — количество символов в поле меньше минимума.   * &#x60;length_greater_than_max&#x60; — количество символов в поле больше максимума.   * &#x60;size_less_than_min&#x60; — количество элементов меньше минимума.   * &#x60;size_greater_than_max&#x60; — количество элементов больше максимума.   * &#x60;send_metro_without_area&#x60; — не передано значение поля &#x60;area&#x60; при заполненном метро.   * &#x60;not_belong_this_city&#x60; — указанного метро нет в указанном городе.   * &#x60;required_with_not_started_career&#x60; — необходимо отправлять опыт работы, если специализация не начало карьеры.   * &#x60;not_match_regexp&#x60; — значение не соответствует регулярному выражению.   * &#x60;more_than_one&#x60; — передано более одного email.   * &#x60;not_available&#x60; — недопустимое значение  | 
**Type** | **string** | Текстовый идентификатор типа ошибки | 
**Value** | **string** | Название поля с ошибкой | 

## Methods

### NewErrorsResumeAddEditError

`func NewErrorsResumeAddEditError(description string, pointer string, reason string, type_ string, value string, ) *ErrorsResumeAddEditError`

NewErrorsResumeAddEditError instantiates a new ErrorsResumeAddEditError object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewErrorsResumeAddEditErrorWithDefaults

`func NewErrorsResumeAddEditErrorWithDefaults() *ErrorsResumeAddEditError`

NewErrorsResumeAddEditErrorWithDefaults instantiates a new ErrorsResumeAddEditError object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDescription

`func (o *ErrorsResumeAddEditError) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *ErrorsResumeAddEditError) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *ErrorsResumeAddEditError) SetDescription(v string)`

SetDescription sets Description field to given value.


### GetPointer

`func (o *ErrorsResumeAddEditError) GetPointer() string`

GetPointer returns the Pointer field if non-nil, zero value otherwise.

### GetPointerOk

`func (o *ErrorsResumeAddEditError) GetPointerOk() (*string, bool)`

GetPointerOk returns a tuple with the Pointer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPointer

`func (o *ErrorsResumeAddEditError) SetPointer(v string)`

SetPointer sets Pointer field to given value.


### GetReason

`func (o *ErrorsResumeAddEditError) GetReason() string`

GetReason returns the Reason field if non-nil, zero value otherwise.

### GetReasonOk

`func (o *ErrorsResumeAddEditError) GetReasonOk() (*string, bool)`

GetReasonOk returns a tuple with the Reason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReason

`func (o *ErrorsResumeAddEditError) SetReason(v string)`

SetReason sets Reason field to given value.


### GetType

`func (o *ErrorsResumeAddEditError) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *ErrorsResumeAddEditError) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *ErrorsResumeAddEditError) SetType(v string)`

SetType sets Type field to given value.


### GetValue

`func (o *ErrorsResumeAddEditError) GetValue() string`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *ErrorsResumeAddEditError) GetValueOk() (*string, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *ErrorsResumeAddEditError) SetValue(v string)`

SetValue sets Value field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


