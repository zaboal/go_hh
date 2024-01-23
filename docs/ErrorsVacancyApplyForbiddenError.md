# ErrorsVacancyApplyForbiddenError

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **string** | Текстовый идентификатор типа ошибки | 
**Value** | **string** | Ошибки при отклике на вакансию:   * &#x60;invalid_vacancy&#x60; — вакансия из отклика/приглашения была архивирована или скрыта.   * &#x60;resume_not_found&#x60; — резюме из отклика/приглашения скрыто, удалено или не найдено.   * &#x60;already_applied&#x60; — в системе уже есть отклик/приглашение на указанную вакансию (связка &#x60;resume_id&#x60; + &#x60;vacancy_id&#x60;)   * &#x60;test_required&#x60; — для отклика необходимо пройти тест.   * &#x60;resume_visibility_conflict&#x60; — невозможно откликнуться на анонимную вакансию, используя резюме с [типом видимости](#tag/Rezyume.-Spiski-vidimosti/operation/get-resume-access-types) &#x60;whitelist&#x60;.   * &#x60;edit_forbidden&#x60; — редактирование сообщения недоступно.   * &#x60;application_denied&#x60; — общая ошибка запрета отклика в случае, когда дополнительная информация недоступна.   * &#x60;limit_exceeded&#x60; — превышен лимит количества откликов/приглашений.   * &#x60;wrong_state&#x60; — действие по отклику/приглашению в данном статусе невозможно.   * &#x60;empty_message&#x60; — передан пустой текст сопроводительного письма.   * &#x60;too_long_message&#x60; — передан слишком длинный текст сопроводительного письма.   * &#x60;address_not_found&#x60; — переданный адрес не существует, либо принадлежит другому работодателю.   * &#x60;not_enough_purchased_services&#x60; — не хватает оплаченных услуг. Например, [доступа к базе резюме](https://hh.ru/price#dbaccess).   * &#x60;in_a_row_limit&#x60; — превышено количество последовательных сообщений в переписке. Необходимо дождаться ответа собеседника на сообщение   * &#x60;overall_limit&#x60; — превышен лимит сообщений.   * &#x60;no_invitation&#x60; — переписка недоступна, так как в отклике ещё не было приглашения.   * &#x60;message_cannot_be_empty&#x60; — сообщение в переписке не может быть пустым.   * &#x60;disabled_by_employer&#x60; — возможность переписки по отклику отключена работодателем.   * &#x60;resume_deleted&#x60; — резюме, с которым совершался отклик, удалено или скрыто.   * &#x60;archived&#x60; — вакансия, на которую совершался отклик, заархивирована.   * &#x60;chat_archived&#x60; — отклик/приглашение заархивировано  | 

## Methods

### NewErrorsVacancyApplyForbiddenError

`func NewErrorsVacancyApplyForbiddenError(type_ string, value string, ) *ErrorsVacancyApplyForbiddenError`

NewErrorsVacancyApplyForbiddenError instantiates a new ErrorsVacancyApplyForbiddenError object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewErrorsVacancyApplyForbiddenErrorWithDefaults

`func NewErrorsVacancyApplyForbiddenErrorWithDefaults() *ErrorsVacancyApplyForbiddenError`

NewErrorsVacancyApplyForbiddenErrorWithDefaults instantiates a new ErrorsVacancyApplyForbiddenError object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *ErrorsVacancyApplyForbiddenError) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *ErrorsVacancyApplyForbiddenError) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *ErrorsVacancyApplyForbiddenError) SetType(v string)`

SetType sets Type field to given value.


### GetValue

`func (o *ErrorsVacancyApplyForbiddenError) GetValue() string`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *ErrorsVacancyApplyForbiddenError) GetValueOk() (*string, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *ErrorsVacancyApplyForbiddenError) SetValue(v string)`

SetValue sets Value field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


