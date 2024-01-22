/*
HeadHunter API

По-русски | [Switch to English](https://api.hh.ru/openapi/en/redoc)  В OpenAPI ведется пока что только небольшая часть документации [Основная документация](https://github.com/hhru/api).  Для поиска по документации можно использовать Ctrl+F.  # Авторизация  API поддерживает следующие уровни авторизации:   - [авторизация приложения](#section/Avtorizaciya/Avtorizaciya-prilozheniya)   - [авторизация пользователя](#section/Avtorizaciya/Avtorizaciya-polzovatelya)  * [Авторизация пользователя](#section/Avtorizaciya/Avtorizaciya-polzovatelya)   * [Правила формирования специального redirect_uri](#section/Avtorizaciya/Pravila-formirovaniya-specialnogo-redirect_uri)   * [Процесс авторизации](#section/Avtorizaciya/Process-avtorizacii)   * [Успешное получение временного `authorization_code`](#get-authorization_code)   * [Получение access и refresh токенов](#section/Avtorizaciya/Poluchenie-access-i-refresh-tokenov) * [Обновление пары access и refresh токенов](#section/Avtorizaciya/Obnovlenie-pary-access-i-refresh-tokenov) * [Инвалидация токена](#tag/Avtorizaciya-rabotodatelya/operation/invalidate-token) * [Запрос авторизации под другим пользователем](#section/Avtorizaciya/Zapros-avtorizacii-pod-drugim-polzovatelem) * [Авторизация под разными рабочими аккаунтами](#section/Avtorizaciya/Avtorizaciya-pod-raznymi-rabochimi-akkauntami) * [Авторизация приложения](#section/Avtorizaciya/Avtorizaciya-prilozheniya)       ## Авторизация пользователя Для выполнения запросов от имени пользователя необходимо пользоваться токеном пользователя.  В начале приложению необходимо направить пользователя (открыть страницу) по адресу:  ``` https://hh.ru/oauth/authorize? response_type=code& client_id={client_id}& state={state}& redirect_uri={redirect_uri} ```  Обязательные параметры:  * `response_type=code` — указание на способ получения авторизации, используя `authorization code` * `client_id` — идентификатор, полученный при создании приложения   Необязательные параметры:  * `state` — в случае указания, будет включен в ответный редирект. Это позволяет исключить возможность взлома путём подделки межсайтовых запросов. Подробнее об этом: [RFC 6749. Section 10.12](http://tools.ietf.org/html/rfc6749#section-10.12) * `redirect_uri` — uri для перенаправления пользователя после авторизации. Если не указать, используется из настроек приложения. При наличии происходит валидация значения. Вероятнее всего, потребуется сделать urlencode значения параметра.  ## Правила формирования специального redirect_uri  К примеру, если в настройках сохранен `http://example.com/oauth`, то разрешено указывать:  * `http://www.example.com/oauth` — поддомен; * `http://www.example.com/oauth/sub/path` — уточнение пути; * `http://example.com/oauth?lang=RU` — дополнительный параметр; * `http://www.example.com/oauth/sub/path?lang=RU` — всё вместе.  Запрещено:  * `https://example.com/oauth` — различные протоколы; * `http://wwwexample.com/oauth` — различные домены; * `http://wwwexample.com/` — другой путь; * `http://example.com/oauths` — другой путь; * `http://example.com:80/oauths` — указание изначально отсутствующего порта;  ## Процесс авторизации  Если пользователь не авторизован на сайте, ему будет показана форма авторизации на сайте. После прохождения авторизации на сайте, пользователю будет выведена форма с запросом разрешения доступа вашего приложения к его персональным данным.  Если пользователь не разрешает доступ приложению, пользователь будет перенаправлен на указанный `redirect_uri` с `?error=access_denied` и `state={state}`, если таковой был указан при первом запросе.  <a name=\"get-authorization_code\"></a> ### Успешное получение временного `authorization_code`  В случае разрешения прав, в редиректе будет указан временный `authorization_code`:  ```http HTTP/1.1 302 FOUND Location: {redirect_uri}?code={authorization_code} ```  Если пользователь авторизован на сайте и доступ данному приложению однажды ранее выдан, ответом будет сразу вышеописанный редирект с `authorization_code` (без показа формы логина и выдачи прав).  ## Получение access и refresh токенов  После получения `authorization_code` приложению необходимо сделать сервер-сервер запрос `POST https://hh.ru/oauth/token` для обмена полученного `authorization_code` на `access_token`.  В теле запроса необходимо передать [дополнительные параметры](#required_parameters).  Тело запроса необходимо передавать в стандартном `application/x-www-form-urlencoded` с указанием соответствующего заголовка `Content-Type`.  `authorization_code` имеет довольно короткий срок жизни, при его истечении необходимо запросить новый.  ## Обновление пары access и refresh токенов `access_token` также имеет срок жизни (ключ `expires_in`, в секундах), при его истечении приложение должно сделать запрос с `refresh_token` для получения нового.  Запрос необходимо делать в `application/x-www-form-urlencoded`.  ``` POST https://hh.ru/oauth/token ```  В теле запроса необходимо передать [дополнительные параметры](#required_parameters)  `refresh_token` можно использовать только один раз и только по истечению срока действия `access_token`.  После получения новой пары access и refresh токенов, их необходимо использовать в дальнейших запросах в api и запросах на продление токена.  ## Запрос авторизации под другим пользователем  Возможен следующий сценарий:  1. Приложение перенаправляет пользователя на сайт с запросом авторизации. 2. Пользователь на сайте уже авторизован и данному приложение доступ уже был разрешен. 3. Пользователю будет предложена возможность продолжить работу под текущим аккаунтом, либо зайти под другим аккаунтом.  Если есть необходимость, чтобы на шаге 3 сразу происходило перенаправление (redirect) с временным токеном, необходимо добавить к запросу `/oauth/authorize...` параметр `skip_choose_account=true`. В этом случае автоматически выдаётся доступ пользователю авторизованному на сайте.  Если есть необходимость всегда показывать форму авторизации, приложение может добавить к запросу `/oauth/authorize...` параметр `force_login=true`. В этом случае, пользователю будет показана форма авторизации с логином и паролем даже в случае, если пользователь уже авторизован.  Это может быть полезно приложениям, которые предоставляют сервис только для соискателей. Если пришел пользователь-работодатель, приложение может предложить пользователю повторно разрешить доступ на сайте, уже указав другую учетную запись.  Также, после авторизации приложение может показать пользователю сообщение:  ``` Вы вошли как %Имя_Фамилия%. Это не вы? ``` и предоставить ссылку с `force_login=true` для возможности захода под другим логином.  ## Авторизация под разными рабочими аккаунтами  Для получения списка рабочих аккаунтов менеджера и для работы под разными рабочими аккаунтами менеджера необходимо прочитать документацию по [рабочим аккаунтам менеджера](#tag/Menedzhery-rabotodatelya/operation/get-manager-accounts)  ## Авторизация приложения  Токен приложения необходимо сгенерировать 1 раз. В случае, если токен был скомпрометирован, его нужно запросить еще раз. При этом ранее выданный токен отзывается. Владелец приложения может посмотреть актуальный `access_token` для приложения на сайте [https://dev.hh.ru/admin](https://dev.hh.ru/admin). В случае, если вы еще ни разу [не получали токен приложения](#section/Avtorizaciya/Avtorizaciya-prilozheniya), токен отображаться не будет.  <a name=\"get-client-token\"></a> ### Получение токена приложения Для получения `access_token` необходимо сделать запрос:  ``` POST https://hh.ru/oauth/token ```  В теле запроса необходимо передать [дополнительные параметры](#required_parameters). Тело запроса необходимо передавать в стандартном `application/x-www-form-urlencoded` с указанием соответствующего заголовка `Content-Type`.  Данный `access_token` имеет **неограниченный** срок жизни. При повторном запросе ранее выданный токен отзывается и выдается новый. Запрашивать `access_token` можно не чаще, чем один раз в 5 минут.  В случае компрометации токена необходимо [инвалидировать](#tag/Avtorizaciya-rabotodatelya/operation/invalidate-token) скомпроментированный токен и запросить токен заново!  <!-- ReDoc-Inject: <security-definitions> --> 

API version: 1.0.0
Contact: api@hh.ru
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package hh

import (
	"encoding/json"
	"bytes"
	"fmt"
)

// checks if the DictionariesDictResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &DictionariesDictResponse{}

// DictionariesDictResponse Справочники полей и сущностей, используемых в API
type DictionariesDictResponse struct {
	// Тип доступа для комментария к соискателю
	ApplicantCommentAccessType []IncludesIdName `json:"applicant_comment_access_type"`
	// Типы сортировки [списка комментариев к соискателю](#tag/Kommentarii-k-soiskatelyu/operation/get-applicant-comments-list)
	ApplicantCommentsOrder []IncludesIdName `json:"applicant_comments_order"`
	// Статусы откликов/приглашений
	ApplicantNegotiationStatus []IncludesIdName `json:"applicant_negotiation_status"`
	// Готовность к командировкам
	BusinessTripReadiness []IncludesIdName `json:"business_trip_readiness"`
	// Справочник валют
	Currency []DictionariesCurrencyItem `json:"currency"`
	// Категории водительских прав
	DriverLicenseTypes []IncludesId `json:"driver_license_types"`
	// Образование в резюме
	EducationLevel []IncludesIdName `json:"education_level"`
	// Тип сортировки списка опубликованных вакансий
	EmployerActiveVacanciesOrder []IncludesIdName `json:"employer_active_vacancies_order"`
	// Тип сортировки списка архивных вакансий
	EmployerArchivedVacanciesOrder []IncludesIdName `json:"employer_archived_vacancies_order"`
	// Тип сортировки скрытых вакансий
	EmployerHiddenVacanciesOrder []IncludesIdName `json:"employer_hidden_vacancies_order,omitempty"`
	// Типы связи компании с пользователем
	EmployerRelation []IncludesIdName `json:"employer_relation"`
	// Тип работодателя
	EmployerType []IncludesIdName `json:"employer_type"`
	// Тип занятости
	Employment []IncludesIdName `json:"employment"`
	// Опыт работы
	Experience []IncludesIdName `json:"experience"`
	// Пол
	Gender []IncludesIdName `json:"gender"`
	// Статусы поиска соискателей для установки и отображения самому соискателю
	JobSearchStatusesApplicant []IncludesIdName `json:"job_search_statuses_applicant"`
	// Статусы поиска соискателей для отображения работодателям
	JobSearchStatusesEmployer []IncludesIdName `json:"job_search_statuses_employer"`
	// Уровень владения языком
	LanguageLevel []IncludesIdName `json:"language_level"`
	// Статус возможности отправки сообщения в переписке
	MessagingStatus []IncludesIdName `json:"messaging_status"`
	// Типы порядка отображения откликов
	NegotiationsOrder []IncludesIdName `json:"negotiations_order"`
	// Типы участников переписки
	NegotiationsParticipantType []IncludesIdName `json:"negotiations_participant_type"`
	// Типы состояний откликов
	NegotiationsState []IncludesIdName `json:"negotiations_state"`
	// Статус звонка, зафиксированного в системе кол-трекинг
	PhoneCallStatus []IncludesIdName `json:"phone_call_status"`
	// Предпочитаемый способ связи
	PreferredContactType []IncludesIdName `json:"preferred_contact_type"`
	// Готовность к переезду
	RelocationType []IncludesIdName `json:"relocation_type"`
	// Уровень доступа к резюме
	ResumeAccessType []IncludesIdName `json:"resume_access_type"`
	// Тип сайта в поле «контакты»
	ResumeContactsSiteType []IncludesIdName `json:"resume_contacts_site_type"`
	// Поля резюме, которые могут быть [скрыты](https://github.com/hhru/api/blob/master/docs/employer_resumes.md#hidden-fields)
	ResumeHiddenFields []IncludesIdName `json:"resume_hidden_fields"`
	// Комментарий модератора
	ResumeModerationNote []IncludesIdName `json:"resume_moderation_note"`
	// Условие поиска по опыту работы
	ResumeSearchExperiencePeriod []IncludesIdName `json:"resume_search_experience_period,omitempty"`
	// Область поиска в резюме
	ResumeSearchFields []IncludesIdName `json:"resume_search_fields,omitempty"`
	// Метки поиска резюме
	ResumeSearchLabel []IncludesIdName `json:"resume_search_label,omitempty"`
	// Условие поиска резюме
	ResumeSearchLogic []IncludesIdName `json:"resume_search_logic,omitempty"`
	// Тип сортировки резюме
	ResumeSearchOrder []IncludesIdName `json:"resume_search_order,omitempty"`
	// Условие поиска по проживанию в регионе и готовности к переезду
	ResumeSearchRelocation []IncludesIdName `json:"resume_search_relocation,omitempty"`
	// Статус резюме
	ResumeStatus []IncludesIdName `json:"resume_status"`
	// График работы
	Schedule []IncludesIdName `json:"schedule"`
	// Время в пути
	TravelTime []IncludesIdName `json:"travel_time"`
	// Варианты размещения вакансии с точки зрения биллинга
	VacancyBillingType []IncludesIdName `json:"vacancy_billing_type"`
	// Тип кластеров
	VacancyCluster []IncludesIdName `json:"vacancy_cluster"`
	// Метки вакансии
	VacancyLabel []IncludesIdName `json:"vacancy_label"`
	// Причины, из-за которых невозможно [продлить вакансию](#tag/Upravlenie-vakansiyami/operation/get-prolongation-vacancy-info)
	VacancyNotProlongedReason []IncludesIdName `json:"vacancy_not_prolonged_reason"`
	// Типы связи вакансии с пользователем
	VacancyRelation []IncludesIdName `json:"vacancy_relation"`
	// Область поиска в вакансии
	VacancySearchFields []IncludesIdName `json:"vacancy_search_fields"`
	// Тип сортировки вакансии
	VacancySearchOrder []IncludesIdName `json:"vacancy_search_order"`
	// Тип вакансии
	VacancyType []IncludesIdName `json:"vacancy_type"`
	// Рабочие дни
	WorkingDays []IncludesIdName `json:"working_days"`
	// Временные интервалы работы
	WorkingTimeIntervals []IncludesIdName `json:"working_time_intervals"`
	// Режимы времени работы
	WorkingTimeModes []IncludesIdName `json:"working_time_modes"`
}

type _DictionariesDictResponse DictionariesDictResponse

// NewDictionariesDictResponse instantiates a new DictionariesDictResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDictionariesDictResponse(applicantCommentAccessType []IncludesIdName, applicantCommentsOrder []IncludesIdName, applicantNegotiationStatus []IncludesIdName, businessTripReadiness []IncludesIdName, currency []DictionariesCurrencyItem, driverLicenseTypes []IncludesId, educationLevel []IncludesIdName, employerActiveVacanciesOrder []IncludesIdName, employerArchivedVacanciesOrder []IncludesIdName, employerRelation []IncludesIdName, employerType []IncludesIdName, employment []IncludesIdName, experience []IncludesIdName, gender []IncludesIdName, jobSearchStatusesApplicant []IncludesIdName, jobSearchStatusesEmployer []IncludesIdName, languageLevel []IncludesIdName, messagingStatus []IncludesIdName, negotiationsOrder []IncludesIdName, negotiationsParticipantType []IncludesIdName, negotiationsState []IncludesIdName, phoneCallStatus []IncludesIdName, preferredContactType []IncludesIdName, relocationType []IncludesIdName, resumeAccessType []IncludesIdName, resumeContactsSiteType []IncludesIdName, resumeHiddenFields []IncludesIdName, resumeModerationNote []IncludesIdName, resumeStatus []IncludesIdName, schedule []IncludesIdName, travelTime []IncludesIdName, vacancyBillingType []IncludesIdName, vacancyCluster []IncludesIdName, vacancyLabel []IncludesIdName, vacancyNotProlongedReason []IncludesIdName, vacancyRelation []IncludesIdName, vacancySearchFields []IncludesIdName, vacancySearchOrder []IncludesIdName, vacancyType []IncludesIdName, workingDays []IncludesIdName, workingTimeIntervals []IncludesIdName, workingTimeModes []IncludesIdName) *DictionariesDictResponse {
	this := DictionariesDictResponse{}
	this.ApplicantCommentAccessType = applicantCommentAccessType
	this.ApplicantCommentsOrder = applicantCommentsOrder
	this.ApplicantNegotiationStatus = applicantNegotiationStatus
	this.BusinessTripReadiness = businessTripReadiness
	this.Currency = currency
	this.DriverLicenseTypes = driverLicenseTypes
	this.EducationLevel = educationLevel
	this.EmployerActiveVacanciesOrder = employerActiveVacanciesOrder
	this.EmployerArchivedVacanciesOrder = employerArchivedVacanciesOrder
	this.EmployerRelation = employerRelation
	this.EmployerType = employerType
	this.Employment = employment
	this.Experience = experience
	this.Gender = gender
	this.JobSearchStatusesApplicant = jobSearchStatusesApplicant
	this.JobSearchStatusesEmployer = jobSearchStatusesEmployer
	this.LanguageLevel = languageLevel
	this.MessagingStatus = messagingStatus
	this.NegotiationsOrder = negotiationsOrder
	this.NegotiationsParticipantType = negotiationsParticipantType
	this.NegotiationsState = negotiationsState
	this.PhoneCallStatus = phoneCallStatus
	this.PreferredContactType = preferredContactType
	this.RelocationType = relocationType
	this.ResumeAccessType = resumeAccessType
	this.ResumeContactsSiteType = resumeContactsSiteType
	this.ResumeHiddenFields = resumeHiddenFields
	this.ResumeModerationNote = resumeModerationNote
	this.ResumeStatus = resumeStatus
	this.Schedule = schedule
	this.TravelTime = travelTime
	this.VacancyBillingType = vacancyBillingType
	this.VacancyCluster = vacancyCluster
	this.VacancyLabel = vacancyLabel
	this.VacancyNotProlongedReason = vacancyNotProlongedReason
	this.VacancyRelation = vacancyRelation
	this.VacancySearchFields = vacancySearchFields
	this.VacancySearchOrder = vacancySearchOrder
	this.VacancyType = vacancyType
	this.WorkingDays = workingDays
	this.WorkingTimeIntervals = workingTimeIntervals
	this.WorkingTimeModes = workingTimeModes
	return &this
}

// NewDictionariesDictResponseWithDefaults instantiates a new DictionariesDictResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDictionariesDictResponseWithDefaults() *DictionariesDictResponse {
	this := DictionariesDictResponse{}
	return &this
}

// GetApplicantCommentAccessType returns the ApplicantCommentAccessType field value
func (o *DictionariesDictResponse) GetApplicantCommentAccessType() []IncludesIdName {
	if o == nil {
		var ret []IncludesIdName
		return ret
	}

	return o.ApplicantCommentAccessType
}

// GetApplicantCommentAccessTypeOk returns a tuple with the ApplicantCommentAccessType field value
// and a boolean to check if the value has been set.
func (o *DictionariesDictResponse) GetApplicantCommentAccessTypeOk() ([]IncludesIdName, bool) {
	if o == nil {
		return nil, false
	}
	return o.ApplicantCommentAccessType, true
}

// SetApplicantCommentAccessType sets field value
func (o *DictionariesDictResponse) SetApplicantCommentAccessType(v []IncludesIdName) {
	o.ApplicantCommentAccessType = v
}

// GetApplicantCommentsOrder returns the ApplicantCommentsOrder field value
func (o *DictionariesDictResponse) GetApplicantCommentsOrder() []IncludesIdName {
	if o == nil {
		var ret []IncludesIdName
		return ret
	}

	return o.ApplicantCommentsOrder
}

// GetApplicantCommentsOrderOk returns a tuple with the ApplicantCommentsOrder field value
// and a boolean to check if the value has been set.
func (o *DictionariesDictResponse) GetApplicantCommentsOrderOk() ([]IncludesIdName, bool) {
	if o == nil {
		return nil, false
	}
	return o.ApplicantCommentsOrder, true
}

// SetApplicantCommentsOrder sets field value
func (o *DictionariesDictResponse) SetApplicantCommentsOrder(v []IncludesIdName) {
	o.ApplicantCommentsOrder = v
}

// GetApplicantNegotiationStatus returns the ApplicantNegotiationStatus field value
func (o *DictionariesDictResponse) GetApplicantNegotiationStatus() []IncludesIdName {
	if o == nil {
		var ret []IncludesIdName
		return ret
	}

	return o.ApplicantNegotiationStatus
}

// GetApplicantNegotiationStatusOk returns a tuple with the ApplicantNegotiationStatus field value
// and a boolean to check if the value has been set.
func (o *DictionariesDictResponse) GetApplicantNegotiationStatusOk() ([]IncludesIdName, bool) {
	if o == nil {
		return nil, false
	}
	return o.ApplicantNegotiationStatus, true
}

// SetApplicantNegotiationStatus sets field value
func (o *DictionariesDictResponse) SetApplicantNegotiationStatus(v []IncludesIdName) {
	o.ApplicantNegotiationStatus = v
}

// GetBusinessTripReadiness returns the BusinessTripReadiness field value
func (o *DictionariesDictResponse) GetBusinessTripReadiness() []IncludesIdName {
	if o == nil {
		var ret []IncludesIdName
		return ret
	}

	return o.BusinessTripReadiness
}

// GetBusinessTripReadinessOk returns a tuple with the BusinessTripReadiness field value
// and a boolean to check if the value has been set.
func (o *DictionariesDictResponse) GetBusinessTripReadinessOk() ([]IncludesIdName, bool) {
	if o == nil {
		return nil, false
	}
	return o.BusinessTripReadiness, true
}

// SetBusinessTripReadiness sets field value
func (o *DictionariesDictResponse) SetBusinessTripReadiness(v []IncludesIdName) {
	o.BusinessTripReadiness = v
}

// GetCurrency returns the Currency field value
func (o *DictionariesDictResponse) GetCurrency() []DictionariesCurrencyItem {
	if o == nil {
		var ret []DictionariesCurrencyItem
		return ret
	}

	return o.Currency
}

// GetCurrencyOk returns a tuple with the Currency field value
// and a boolean to check if the value has been set.
func (o *DictionariesDictResponse) GetCurrencyOk() ([]DictionariesCurrencyItem, bool) {
	if o == nil {
		return nil, false
	}
	return o.Currency, true
}

// SetCurrency sets field value
func (o *DictionariesDictResponse) SetCurrency(v []DictionariesCurrencyItem) {
	o.Currency = v
}

// GetDriverLicenseTypes returns the DriverLicenseTypes field value
func (o *DictionariesDictResponse) GetDriverLicenseTypes() []IncludesId {
	if o == nil {
		var ret []IncludesId
		return ret
	}

	return o.DriverLicenseTypes
}

// GetDriverLicenseTypesOk returns a tuple with the DriverLicenseTypes field value
// and a boolean to check if the value has been set.
func (o *DictionariesDictResponse) GetDriverLicenseTypesOk() ([]IncludesId, bool) {
	if o == nil {
		return nil, false
	}
	return o.DriverLicenseTypes, true
}

// SetDriverLicenseTypes sets field value
func (o *DictionariesDictResponse) SetDriverLicenseTypes(v []IncludesId) {
	o.DriverLicenseTypes = v
}

// GetEducationLevel returns the EducationLevel field value
func (o *DictionariesDictResponse) GetEducationLevel() []IncludesIdName {
	if o == nil {
		var ret []IncludesIdName
		return ret
	}

	return o.EducationLevel
}

// GetEducationLevelOk returns a tuple with the EducationLevel field value
// and a boolean to check if the value has been set.
func (o *DictionariesDictResponse) GetEducationLevelOk() ([]IncludesIdName, bool) {
	if o == nil {
		return nil, false
	}
	return o.EducationLevel, true
}

// SetEducationLevel sets field value
func (o *DictionariesDictResponse) SetEducationLevel(v []IncludesIdName) {
	o.EducationLevel = v
}

// GetEmployerActiveVacanciesOrder returns the EmployerActiveVacanciesOrder field value
func (o *DictionariesDictResponse) GetEmployerActiveVacanciesOrder() []IncludesIdName {
	if o == nil {
		var ret []IncludesIdName
		return ret
	}

	return o.EmployerActiveVacanciesOrder
}

// GetEmployerActiveVacanciesOrderOk returns a tuple with the EmployerActiveVacanciesOrder field value
// and a boolean to check if the value has been set.
func (o *DictionariesDictResponse) GetEmployerActiveVacanciesOrderOk() ([]IncludesIdName, bool) {
	if o == nil {
		return nil, false
	}
	return o.EmployerActiveVacanciesOrder, true
}

// SetEmployerActiveVacanciesOrder sets field value
func (o *DictionariesDictResponse) SetEmployerActiveVacanciesOrder(v []IncludesIdName) {
	o.EmployerActiveVacanciesOrder = v
}

// GetEmployerArchivedVacanciesOrder returns the EmployerArchivedVacanciesOrder field value
func (o *DictionariesDictResponse) GetEmployerArchivedVacanciesOrder() []IncludesIdName {
	if o == nil {
		var ret []IncludesIdName
		return ret
	}

	return o.EmployerArchivedVacanciesOrder
}

// GetEmployerArchivedVacanciesOrderOk returns a tuple with the EmployerArchivedVacanciesOrder field value
// and a boolean to check if the value has been set.
func (o *DictionariesDictResponse) GetEmployerArchivedVacanciesOrderOk() ([]IncludesIdName, bool) {
	if o == nil {
		return nil, false
	}
	return o.EmployerArchivedVacanciesOrder, true
}

// SetEmployerArchivedVacanciesOrder sets field value
func (o *DictionariesDictResponse) SetEmployerArchivedVacanciesOrder(v []IncludesIdName) {
	o.EmployerArchivedVacanciesOrder = v
}

// GetEmployerHiddenVacanciesOrder returns the EmployerHiddenVacanciesOrder field value if set, zero value otherwise.
func (o *DictionariesDictResponse) GetEmployerHiddenVacanciesOrder() []IncludesIdName {
	if o == nil || IsNil(o.EmployerHiddenVacanciesOrder) {
		var ret []IncludesIdName
		return ret
	}
	return o.EmployerHiddenVacanciesOrder
}

// GetEmployerHiddenVacanciesOrderOk returns a tuple with the EmployerHiddenVacanciesOrder field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DictionariesDictResponse) GetEmployerHiddenVacanciesOrderOk() ([]IncludesIdName, bool) {
	if o == nil || IsNil(o.EmployerHiddenVacanciesOrder) {
		return nil, false
	}
	return o.EmployerHiddenVacanciesOrder, true
}

// HasEmployerHiddenVacanciesOrder returns a boolean if a field has been set.
func (o *DictionariesDictResponse) HasEmployerHiddenVacanciesOrder() bool {
	if o != nil && !IsNil(o.EmployerHiddenVacanciesOrder) {
		return true
	}

	return false
}

// SetEmployerHiddenVacanciesOrder gets a reference to the given []IncludesIdName and assigns it to the EmployerHiddenVacanciesOrder field.
func (o *DictionariesDictResponse) SetEmployerHiddenVacanciesOrder(v []IncludesIdName) {
	o.EmployerHiddenVacanciesOrder = v
}

// GetEmployerRelation returns the EmployerRelation field value
func (o *DictionariesDictResponse) GetEmployerRelation() []IncludesIdName {
	if o == nil {
		var ret []IncludesIdName
		return ret
	}

	return o.EmployerRelation
}

// GetEmployerRelationOk returns a tuple with the EmployerRelation field value
// and a boolean to check if the value has been set.
func (o *DictionariesDictResponse) GetEmployerRelationOk() ([]IncludesIdName, bool) {
	if o == nil {
		return nil, false
	}
	return o.EmployerRelation, true
}

// SetEmployerRelation sets field value
func (o *DictionariesDictResponse) SetEmployerRelation(v []IncludesIdName) {
	o.EmployerRelation = v
}

// GetEmployerType returns the EmployerType field value
func (o *DictionariesDictResponse) GetEmployerType() []IncludesIdName {
	if o == nil {
		var ret []IncludesIdName
		return ret
	}

	return o.EmployerType
}

// GetEmployerTypeOk returns a tuple with the EmployerType field value
// and a boolean to check if the value has been set.
func (o *DictionariesDictResponse) GetEmployerTypeOk() ([]IncludesIdName, bool) {
	if o == nil {
		return nil, false
	}
	return o.EmployerType, true
}

// SetEmployerType sets field value
func (o *DictionariesDictResponse) SetEmployerType(v []IncludesIdName) {
	o.EmployerType = v
}

// GetEmployment returns the Employment field value
func (o *DictionariesDictResponse) GetEmployment() []IncludesIdName {
	if o == nil {
		var ret []IncludesIdName
		return ret
	}

	return o.Employment
}

// GetEmploymentOk returns a tuple with the Employment field value
// and a boolean to check if the value has been set.
func (o *DictionariesDictResponse) GetEmploymentOk() ([]IncludesIdName, bool) {
	if o == nil {
		return nil, false
	}
	return o.Employment, true
}

// SetEmployment sets field value
func (o *DictionariesDictResponse) SetEmployment(v []IncludesIdName) {
	o.Employment = v
}

// GetExperience returns the Experience field value
func (o *DictionariesDictResponse) GetExperience() []IncludesIdName {
	if o == nil {
		var ret []IncludesIdName
		return ret
	}

	return o.Experience
}

// GetExperienceOk returns a tuple with the Experience field value
// and a boolean to check if the value has been set.
func (o *DictionariesDictResponse) GetExperienceOk() ([]IncludesIdName, bool) {
	if o == nil {
		return nil, false
	}
	return o.Experience, true
}

// SetExperience sets field value
func (o *DictionariesDictResponse) SetExperience(v []IncludesIdName) {
	o.Experience = v
}

// GetGender returns the Gender field value
func (o *DictionariesDictResponse) GetGender() []IncludesIdName {
	if o == nil {
		var ret []IncludesIdName
		return ret
	}

	return o.Gender
}

// GetGenderOk returns a tuple with the Gender field value
// and a boolean to check if the value has been set.
func (o *DictionariesDictResponse) GetGenderOk() ([]IncludesIdName, bool) {
	if o == nil {
		return nil, false
	}
	return o.Gender, true
}

// SetGender sets field value
func (o *DictionariesDictResponse) SetGender(v []IncludesIdName) {
	o.Gender = v
}

// GetJobSearchStatusesApplicant returns the JobSearchStatusesApplicant field value
func (o *DictionariesDictResponse) GetJobSearchStatusesApplicant() []IncludesIdName {
	if o == nil {
		var ret []IncludesIdName
		return ret
	}

	return o.JobSearchStatusesApplicant
}

// GetJobSearchStatusesApplicantOk returns a tuple with the JobSearchStatusesApplicant field value
// and a boolean to check if the value has been set.
func (o *DictionariesDictResponse) GetJobSearchStatusesApplicantOk() ([]IncludesIdName, bool) {
	if o == nil {
		return nil, false
	}
	return o.JobSearchStatusesApplicant, true
}

// SetJobSearchStatusesApplicant sets field value
func (o *DictionariesDictResponse) SetJobSearchStatusesApplicant(v []IncludesIdName) {
	o.JobSearchStatusesApplicant = v
}

// GetJobSearchStatusesEmployer returns the JobSearchStatusesEmployer field value
func (o *DictionariesDictResponse) GetJobSearchStatusesEmployer() []IncludesIdName {
	if o == nil {
		var ret []IncludesIdName
		return ret
	}

	return o.JobSearchStatusesEmployer
}

// GetJobSearchStatusesEmployerOk returns a tuple with the JobSearchStatusesEmployer field value
// and a boolean to check if the value has been set.
func (o *DictionariesDictResponse) GetJobSearchStatusesEmployerOk() ([]IncludesIdName, bool) {
	if o == nil {
		return nil, false
	}
	return o.JobSearchStatusesEmployer, true
}

// SetJobSearchStatusesEmployer sets field value
func (o *DictionariesDictResponse) SetJobSearchStatusesEmployer(v []IncludesIdName) {
	o.JobSearchStatusesEmployer = v
}

// GetLanguageLevel returns the LanguageLevel field value
func (o *DictionariesDictResponse) GetLanguageLevel() []IncludesIdName {
	if o == nil {
		var ret []IncludesIdName
		return ret
	}

	return o.LanguageLevel
}

// GetLanguageLevelOk returns a tuple with the LanguageLevel field value
// and a boolean to check if the value has been set.
func (o *DictionariesDictResponse) GetLanguageLevelOk() ([]IncludesIdName, bool) {
	if o == nil {
		return nil, false
	}
	return o.LanguageLevel, true
}

// SetLanguageLevel sets field value
func (o *DictionariesDictResponse) SetLanguageLevel(v []IncludesIdName) {
	o.LanguageLevel = v
}

// GetMessagingStatus returns the MessagingStatus field value
func (o *DictionariesDictResponse) GetMessagingStatus() []IncludesIdName {
	if o == nil {
		var ret []IncludesIdName
		return ret
	}

	return o.MessagingStatus
}

// GetMessagingStatusOk returns a tuple with the MessagingStatus field value
// and a boolean to check if the value has been set.
func (o *DictionariesDictResponse) GetMessagingStatusOk() ([]IncludesIdName, bool) {
	if o == nil {
		return nil, false
	}
	return o.MessagingStatus, true
}

// SetMessagingStatus sets field value
func (o *DictionariesDictResponse) SetMessagingStatus(v []IncludesIdName) {
	o.MessagingStatus = v
}

// GetNegotiationsOrder returns the NegotiationsOrder field value
func (o *DictionariesDictResponse) GetNegotiationsOrder() []IncludesIdName {
	if o == nil {
		var ret []IncludesIdName
		return ret
	}

	return o.NegotiationsOrder
}

// GetNegotiationsOrderOk returns a tuple with the NegotiationsOrder field value
// and a boolean to check if the value has been set.
func (o *DictionariesDictResponse) GetNegotiationsOrderOk() ([]IncludesIdName, bool) {
	if o == nil {
		return nil, false
	}
	return o.NegotiationsOrder, true
}

// SetNegotiationsOrder sets field value
func (o *DictionariesDictResponse) SetNegotiationsOrder(v []IncludesIdName) {
	o.NegotiationsOrder = v
}

// GetNegotiationsParticipantType returns the NegotiationsParticipantType field value
func (o *DictionariesDictResponse) GetNegotiationsParticipantType() []IncludesIdName {
	if o == nil {
		var ret []IncludesIdName
		return ret
	}

	return o.NegotiationsParticipantType
}

// GetNegotiationsParticipantTypeOk returns a tuple with the NegotiationsParticipantType field value
// and a boolean to check if the value has been set.
func (o *DictionariesDictResponse) GetNegotiationsParticipantTypeOk() ([]IncludesIdName, bool) {
	if o == nil {
		return nil, false
	}
	return o.NegotiationsParticipantType, true
}

// SetNegotiationsParticipantType sets field value
func (o *DictionariesDictResponse) SetNegotiationsParticipantType(v []IncludesIdName) {
	o.NegotiationsParticipantType = v
}

// GetNegotiationsState returns the NegotiationsState field value
func (o *DictionariesDictResponse) GetNegotiationsState() []IncludesIdName {
	if o == nil {
		var ret []IncludesIdName
		return ret
	}

	return o.NegotiationsState
}

// GetNegotiationsStateOk returns a tuple with the NegotiationsState field value
// and a boolean to check if the value has been set.
func (o *DictionariesDictResponse) GetNegotiationsStateOk() ([]IncludesIdName, bool) {
	if o == nil {
		return nil, false
	}
	return o.NegotiationsState, true
}

// SetNegotiationsState sets field value
func (o *DictionariesDictResponse) SetNegotiationsState(v []IncludesIdName) {
	o.NegotiationsState = v
}

// GetPhoneCallStatus returns the PhoneCallStatus field value
func (o *DictionariesDictResponse) GetPhoneCallStatus() []IncludesIdName {
	if o == nil {
		var ret []IncludesIdName
		return ret
	}

	return o.PhoneCallStatus
}

// GetPhoneCallStatusOk returns a tuple with the PhoneCallStatus field value
// and a boolean to check if the value has been set.
func (o *DictionariesDictResponse) GetPhoneCallStatusOk() ([]IncludesIdName, bool) {
	if o == nil {
		return nil, false
	}
	return o.PhoneCallStatus, true
}

// SetPhoneCallStatus sets field value
func (o *DictionariesDictResponse) SetPhoneCallStatus(v []IncludesIdName) {
	o.PhoneCallStatus = v
}

// GetPreferredContactType returns the PreferredContactType field value
func (o *DictionariesDictResponse) GetPreferredContactType() []IncludesIdName {
	if o == nil {
		var ret []IncludesIdName
		return ret
	}

	return o.PreferredContactType
}

// GetPreferredContactTypeOk returns a tuple with the PreferredContactType field value
// and a boolean to check if the value has been set.
func (o *DictionariesDictResponse) GetPreferredContactTypeOk() ([]IncludesIdName, bool) {
	if o == nil {
		return nil, false
	}
	return o.PreferredContactType, true
}

// SetPreferredContactType sets field value
func (o *DictionariesDictResponse) SetPreferredContactType(v []IncludesIdName) {
	o.PreferredContactType = v
}

// GetRelocationType returns the RelocationType field value
func (o *DictionariesDictResponse) GetRelocationType() []IncludesIdName {
	if o == nil {
		var ret []IncludesIdName
		return ret
	}

	return o.RelocationType
}

// GetRelocationTypeOk returns a tuple with the RelocationType field value
// and a boolean to check if the value has been set.
func (o *DictionariesDictResponse) GetRelocationTypeOk() ([]IncludesIdName, bool) {
	if o == nil {
		return nil, false
	}
	return o.RelocationType, true
}

// SetRelocationType sets field value
func (o *DictionariesDictResponse) SetRelocationType(v []IncludesIdName) {
	o.RelocationType = v
}

// GetResumeAccessType returns the ResumeAccessType field value
func (o *DictionariesDictResponse) GetResumeAccessType() []IncludesIdName {
	if o == nil {
		var ret []IncludesIdName
		return ret
	}

	return o.ResumeAccessType
}

// GetResumeAccessTypeOk returns a tuple with the ResumeAccessType field value
// and a boolean to check if the value has been set.
func (o *DictionariesDictResponse) GetResumeAccessTypeOk() ([]IncludesIdName, bool) {
	if o == nil {
		return nil, false
	}
	return o.ResumeAccessType, true
}

// SetResumeAccessType sets field value
func (o *DictionariesDictResponse) SetResumeAccessType(v []IncludesIdName) {
	o.ResumeAccessType = v
}

// GetResumeContactsSiteType returns the ResumeContactsSiteType field value
func (o *DictionariesDictResponse) GetResumeContactsSiteType() []IncludesIdName {
	if o == nil {
		var ret []IncludesIdName
		return ret
	}

	return o.ResumeContactsSiteType
}

// GetResumeContactsSiteTypeOk returns a tuple with the ResumeContactsSiteType field value
// and a boolean to check if the value has been set.
func (o *DictionariesDictResponse) GetResumeContactsSiteTypeOk() ([]IncludesIdName, bool) {
	if o == nil {
		return nil, false
	}
	return o.ResumeContactsSiteType, true
}

// SetResumeContactsSiteType sets field value
func (o *DictionariesDictResponse) SetResumeContactsSiteType(v []IncludesIdName) {
	o.ResumeContactsSiteType = v
}

// GetResumeHiddenFields returns the ResumeHiddenFields field value
func (o *DictionariesDictResponse) GetResumeHiddenFields() []IncludesIdName {
	if o == nil {
		var ret []IncludesIdName
		return ret
	}

	return o.ResumeHiddenFields
}

// GetResumeHiddenFieldsOk returns a tuple with the ResumeHiddenFields field value
// and a boolean to check if the value has been set.
func (o *DictionariesDictResponse) GetResumeHiddenFieldsOk() ([]IncludesIdName, bool) {
	if o == nil {
		return nil, false
	}
	return o.ResumeHiddenFields, true
}

// SetResumeHiddenFields sets field value
func (o *DictionariesDictResponse) SetResumeHiddenFields(v []IncludesIdName) {
	o.ResumeHiddenFields = v
}

// GetResumeModerationNote returns the ResumeModerationNote field value
func (o *DictionariesDictResponse) GetResumeModerationNote() []IncludesIdName {
	if o == nil {
		var ret []IncludesIdName
		return ret
	}

	return o.ResumeModerationNote
}

// GetResumeModerationNoteOk returns a tuple with the ResumeModerationNote field value
// and a boolean to check if the value has been set.
func (o *DictionariesDictResponse) GetResumeModerationNoteOk() ([]IncludesIdName, bool) {
	if o == nil {
		return nil, false
	}
	return o.ResumeModerationNote, true
}

// SetResumeModerationNote sets field value
func (o *DictionariesDictResponse) SetResumeModerationNote(v []IncludesIdName) {
	o.ResumeModerationNote = v
}

// GetResumeSearchExperiencePeriod returns the ResumeSearchExperiencePeriod field value if set, zero value otherwise.
func (o *DictionariesDictResponse) GetResumeSearchExperiencePeriod() []IncludesIdName {
	if o == nil || IsNil(o.ResumeSearchExperiencePeriod) {
		var ret []IncludesIdName
		return ret
	}
	return o.ResumeSearchExperiencePeriod
}

// GetResumeSearchExperiencePeriodOk returns a tuple with the ResumeSearchExperiencePeriod field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DictionariesDictResponse) GetResumeSearchExperiencePeriodOk() ([]IncludesIdName, bool) {
	if o == nil || IsNil(o.ResumeSearchExperiencePeriod) {
		return nil, false
	}
	return o.ResumeSearchExperiencePeriod, true
}

// HasResumeSearchExperiencePeriod returns a boolean if a field has been set.
func (o *DictionariesDictResponse) HasResumeSearchExperiencePeriod() bool {
	if o != nil && !IsNil(o.ResumeSearchExperiencePeriod) {
		return true
	}

	return false
}

// SetResumeSearchExperiencePeriod gets a reference to the given []IncludesIdName and assigns it to the ResumeSearchExperiencePeriod field.
func (o *DictionariesDictResponse) SetResumeSearchExperiencePeriod(v []IncludesIdName) {
	o.ResumeSearchExperiencePeriod = v
}

// GetResumeSearchFields returns the ResumeSearchFields field value if set, zero value otherwise.
func (o *DictionariesDictResponse) GetResumeSearchFields() []IncludesIdName {
	if o == nil || IsNil(o.ResumeSearchFields) {
		var ret []IncludesIdName
		return ret
	}
	return o.ResumeSearchFields
}

// GetResumeSearchFieldsOk returns a tuple with the ResumeSearchFields field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DictionariesDictResponse) GetResumeSearchFieldsOk() ([]IncludesIdName, bool) {
	if o == nil || IsNil(o.ResumeSearchFields) {
		return nil, false
	}
	return o.ResumeSearchFields, true
}

// HasResumeSearchFields returns a boolean if a field has been set.
func (o *DictionariesDictResponse) HasResumeSearchFields() bool {
	if o != nil && !IsNil(o.ResumeSearchFields) {
		return true
	}

	return false
}

// SetResumeSearchFields gets a reference to the given []IncludesIdName and assigns it to the ResumeSearchFields field.
func (o *DictionariesDictResponse) SetResumeSearchFields(v []IncludesIdName) {
	o.ResumeSearchFields = v
}

// GetResumeSearchLabel returns the ResumeSearchLabel field value if set, zero value otherwise.
func (o *DictionariesDictResponse) GetResumeSearchLabel() []IncludesIdName {
	if o == nil || IsNil(o.ResumeSearchLabel) {
		var ret []IncludesIdName
		return ret
	}
	return o.ResumeSearchLabel
}

// GetResumeSearchLabelOk returns a tuple with the ResumeSearchLabel field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DictionariesDictResponse) GetResumeSearchLabelOk() ([]IncludesIdName, bool) {
	if o == nil || IsNil(o.ResumeSearchLabel) {
		return nil, false
	}
	return o.ResumeSearchLabel, true
}

// HasResumeSearchLabel returns a boolean if a field has been set.
func (o *DictionariesDictResponse) HasResumeSearchLabel() bool {
	if o != nil && !IsNil(o.ResumeSearchLabel) {
		return true
	}

	return false
}

// SetResumeSearchLabel gets a reference to the given []IncludesIdName and assigns it to the ResumeSearchLabel field.
func (o *DictionariesDictResponse) SetResumeSearchLabel(v []IncludesIdName) {
	o.ResumeSearchLabel = v
}

// GetResumeSearchLogic returns the ResumeSearchLogic field value if set, zero value otherwise.
func (o *DictionariesDictResponse) GetResumeSearchLogic() []IncludesIdName {
	if o == nil || IsNil(o.ResumeSearchLogic) {
		var ret []IncludesIdName
		return ret
	}
	return o.ResumeSearchLogic
}

// GetResumeSearchLogicOk returns a tuple with the ResumeSearchLogic field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DictionariesDictResponse) GetResumeSearchLogicOk() ([]IncludesIdName, bool) {
	if o == nil || IsNil(o.ResumeSearchLogic) {
		return nil, false
	}
	return o.ResumeSearchLogic, true
}

// HasResumeSearchLogic returns a boolean if a field has been set.
func (o *DictionariesDictResponse) HasResumeSearchLogic() bool {
	if o != nil && !IsNil(o.ResumeSearchLogic) {
		return true
	}

	return false
}

// SetResumeSearchLogic gets a reference to the given []IncludesIdName and assigns it to the ResumeSearchLogic field.
func (o *DictionariesDictResponse) SetResumeSearchLogic(v []IncludesIdName) {
	o.ResumeSearchLogic = v
}

// GetResumeSearchOrder returns the ResumeSearchOrder field value if set, zero value otherwise.
func (o *DictionariesDictResponse) GetResumeSearchOrder() []IncludesIdName {
	if o == nil || IsNil(o.ResumeSearchOrder) {
		var ret []IncludesIdName
		return ret
	}
	return o.ResumeSearchOrder
}

// GetResumeSearchOrderOk returns a tuple with the ResumeSearchOrder field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DictionariesDictResponse) GetResumeSearchOrderOk() ([]IncludesIdName, bool) {
	if o == nil || IsNil(o.ResumeSearchOrder) {
		return nil, false
	}
	return o.ResumeSearchOrder, true
}

// HasResumeSearchOrder returns a boolean if a field has been set.
func (o *DictionariesDictResponse) HasResumeSearchOrder() bool {
	if o != nil && !IsNil(o.ResumeSearchOrder) {
		return true
	}

	return false
}

// SetResumeSearchOrder gets a reference to the given []IncludesIdName and assigns it to the ResumeSearchOrder field.
func (o *DictionariesDictResponse) SetResumeSearchOrder(v []IncludesIdName) {
	o.ResumeSearchOrder = v
}

// GetResumeSearchRelocation returns the ResumeSearchRelocation field value if set, zero value otherwise.
func (o *DictionariesDictResponse) GetResumeSearchRelocation() []IncludesIdName {
	if o == nil || IsNil(o.ResumeSearchRelocation) {
		var ret []IncludesIdName
		return ret
	}
	return o.ResumeSearchRelocation
}

// GetResumeSearchRelocationOk returns a tuple with the ResumeSearchRelocation field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DictionariesDictResponse) GetResumeSearchRelocationOk() ([]IncludesIdName, bool) {
	if o == nil || IsNil(o.ResumeSearchRelocation) {
		return nil, false
	}
	return o.ResumeSearchRelocation, true
}

// HasResumeSearchRelocation returns a boolean if a field has been set.
func (o *DictionariesDictResponse) HasResumeSearchRelocation() bool {
	if o != nil && !IsNil(o.ResumeSearchRelocation) {
		return true
	}

	return false
}

// SetResumeSearchRelocation gets a reference to the given []IncludesIdName and assigns it to the ResumeSearchRelocation field.
func (o *DictionariesDictResponse) SetResumeSearchRelocation(v []IncludesIdName) {
	o.ResumeSearchRelocation = v
}

// GetResumeStatus returns the ResumeStatus field value
func (o *DictionariesDictResponse) GetResumeStatus() []IncludesIdName {
	if o == nil {
		var ret []IncludesIdName
		return ret
	}

	return o.ResumeStatus
}

// GetResumeStatusOk returns a tuple with the ResumeStatus field value
// and a boolean to check if the value has been set.
func (o *DictionariesDictResponse) GetResumeStatusOk() ([]IncludesIdName, bool) {
	if o == nil {
		return nil, false
	}
	return o.ResumeStatus, true
}

// SetResumeStatus sets field value
func (o *DictionariesDictResponse) SetResumeStatus(v []IncludesIdName) {
	o.ResumeStatus = v
}

// GetSchedule returns the Schedule field value
func (o *DictionariesDictResponse) GetSchedule() []IncludesIdName {
	if o == nil {
		var ret []IncludesIdName
		return ret
	}

	return o.Schedule
}

// GetScheduleOk returns a tuple with the Schedule field value
// and a boolean to check if the value has been set.
func (o *DictionariesDictResponse) GetScheduleOk() ([]IncludesIdName, bool) {
	if o == nil {
		return nil, false
	}
	return o.Schedule, true
}

// SetSchedule sets field value
func (o *DictionariesDictResponse) SetSchedule(v []IncludesIdName) {
	o.Schedule = v
}

// GetTravelTime returns the TravelTime field value
func (o *DictionariesDictResponse) GetTravelTime() []IncludesIdName {
	if o == nil {
		var ret []IncludesIdName
		return ret
	}

	return o.TravelTime
}

// GetTravelTimeOk returns a tuple with the TravelTime field value
// and a boolean to check if the value has been set.
func (o *DictionariesDictResponse) GetTravelTimeOk() ([]IncludesIdName, bool) {
	if o == nil {
		return nil, false
	}
	return o.TravelTime, true
}

// SetTravelTime sets field value
func (o *DictionariesDictResponse) SetTravelTime(v []IncludesIdName) {
	o.TravelTime = v
}

// GetVacancyBillingType returns the VacancyBillingType field value
func (o *DictionariesDictResponse) GetVacancyBillingType() []IncludesIdName {
	if o == nil {
		var ret []IncludesIdName
		return ret
	}

	return o.VacancyBillingType
}

// GetVacancyBillingTypeOk returns a tuple with the VacancyBillingType field value
// and a boolean to check if the value has been set.
func (o *DictionariesDictResponse) GetVacancyBillingTypeOk() ([]IncludesIdName, bool) {
	if o == nil {
		return nil, false
	}
	return o.VacancyBillingType, true
}

// SetVacancyBillingType sets field value
func (o *DictionariesDictResponse) SetVacancyBillingType(v []IncludesIdName) {
	o.VacancyBillingType = v
}

// GetVacancyCluster returns the VacancyCluster field value
func (o *DictionariesDictResponse) GetVacancyCluster() []IncludesIdName {
	if o == nil {
		var ret []IncludesIdName
		return ret
	}

	return o.VacancyCluster
}

// GetVacancyClusterOk returns a tuple with the VacancyCluster field value
// and a boolean to check if the value has been set.
func (o *DictionariesDictResponse) GetVacancyClusterOk() ([]IncludesIdName, bool) {
	if o == nil {
		return nil, false
	}
	return o.VacancyCluster, true
}

// SetVacancyCluster sets field value
func (o *DictionariesDictResponse) SetVacancyCluster(v []IncludesIdName) {
	o.VacancyCluster = v
}

// GetVacancyLabel returns the VacancyLabel field value
func (o *DictionariesDictResponse) GetVacancyLabel() []IncludesIdName {
	if o == nil {
		var ret []IncludesIdName
		return ret
	}

	return o.VacancyLabel
}

// GetVacancyLabelOk returns a tuple with the VacancyLabel field value
// and a boolean to check if the value has been set.
func (o *DictionariesDictResponse) GetVacancyLabelOk() ([]IncludesIdName, bool) {
	if o == nil {
		return nil, false
	}
	return o.VacancyLabel, true
}

// SetVacancyLabel sets field value
func (o *DictionariesDictResponse) SetVacancyLabel(v []IncludesIdName) {
	o.VacancyLabel = v
}

// GetVacancyNotProlongedReason returns the VacancyNotProlongedReason field value
func (o *DictionariesDictResponse) GetVacancyNotProlongedReason() []IncludesIdName {
	if o == nil {
		var ret []IncludesIdName
		return ret
	}

	return o.VacancyNotProlongedReason
}

// GetVacancyNotProlongedReasonOk returns a tuple with the VacancyNotProlongedReason field value
// and a boolean to check if the value has been set.
func (o *DictionariesDictResponse) GetVacancyNotProlongedReasonOk() ([]IncludesIdName, bool) {
	if o == nil {
		return nil, false
	}
	return o.VacancyNotProlongedReason, true
}

// SetVacancyNotProlongedReason sets field value
func (o *DictionariesDictResponse) SetVacancyNotProlongedReason(v []IncludesIdName) {
	o.VacancyNotProlongedReason = v
}

// GetVacancyRelation returns the VacancyRelation field value
func (o *DictionariesDictResponse) GetVacancyRelation() []IncludesIdName {
	if o == nil {
		var ret []IncludesIdName
		return ret
	}

	return o.VacancyRelation
}

// GetVacancyRelationOk returns a tuple with the VacancyRelation field value
// and a boolean to check if the value has been set.
func (o *DictionariesDictResponse) GetVacancyRelationOk() ([]IncludesIdName, bool) {
	if o == nil {
		return nil, false
	}
	return o.VacancyRelation, true
}

// SetVacancyRelation sets field value
func (o *DictionariesDictResponse) SetVacancyRelation(v []IncludesIdName) {
	o.VacancyRelation = v
}

// GetVacancySearchFields returns the VacancySearchFields field value
func (o *DictionariesDictResponse) GetVacancySearchFields() []IncludesIdName {
	if o == nil {
		var ret []IncludesIdName
		return ret
	}

	return o.VacancySearchFields
}

// GetVacancySearchFieldsOk returns a tuple with the VacancySearchFields field value
// and a boolean to check if the value has been set.
func (o *DictionariesDictResponse) GetVacancySearchFieldsOk() ([]IncludesIdName, bool) {
	if o == nil {
		return nil, false
	}
	return o.VacancySearchFields, true
}

// SetVacancySearchFields sets field value
func (o *DictionariesDictResponse) SetVacancySearchFields(v []IncludesIdName) {
	o.VacancySearchFields = v
}

// GetVacancySearchOrder returns the VacancySearchOrder field value
func (o *DictionariesDictResponse) GetVacancySearchOrder() []IncludesIdName {
	if o == nil {
		var ret []IncludesIdName
		return ret
	}

	return o.VacancySearchOrder
}

// GetVacancySearchOrderOk returns a tuple with the VacancySearchOrder field value
// and a boolean to check if the value has been set.
func (o *DictionariesDictResponse) GetVacancySearchOrderOk() ([]IncludesIdName, bool) {
	if o == nil {
		return nil, false
	}
	return o.VacancySearchOrder, true
}

// SetVacancySearchOrder sets field value
func (o *DictionariesDictResponse) SetVacancySearchOrder(v []IncludesIdName) {
	o.VacancySearchOrder = v
}

// GetVacancyType returns the VacancyType field value
func (o *DictionariesDictResponse) GetVacancyType() []IncludesIdName {
	if o == nil {
		var ret []IncludesIdName
		return ret
	}

	return o.VacancyType
}

// GetVacancyTypeOk returns a tuple with the VacancyType field value
// and a boolean to check if the value has been set.
func (o *DictionariesDictResponse) GetVacancyTypeOk() ([]IncludesIdName, bool) {
	if o == nil {
		return nil, false
	}
	return o.VacancyType, true
}

// SetVacancyType sets field value
func (o *DictionariesDictResponse) SetVacancyType(v []IncludesIdName) {
	o.VacancyType = v
}

// GetWorkingDays returns the WorkingDays field value
func (o *DictionariesDictResponse) GetWorkingDays() []IncludesIdName {
	if o == nil {
		var ret []IncludesIdName
		return ret
	}

	return o.WorkingDays
}

// GetWorkingDaysOk returns a tuple with the WorkingDays field value
// and a boolean to check if the value has been set.
func (o *DictionariesDictResponse) GetWorkingDaysOk() ([]IncludesIdName, bool) {
	if o == nil {
		return nil, false
	}
	return o.WorkingDays, true
}

// SetWorkingDays sets field value
func (o *DictionariesDictResponse) SetWorkingDays(v []IncludesIdName) {
	o.WorkingDays = v
}

// GetWorkingTimeIntervals returns the WorkingTimeIntervals field value
func (o *DictionariesDictResponse) GetWorkingTimeIntervals() []IncludesIdName {
	if o == nil {
		var ret []IncludesIdName
		return ret
	}

	return o.WorkingTimeIntervals
}

// GetWorkingTimeIntervalsOk returns a tuple with the WorkingTimeIntervals field value
// and a boolean to check if the value has been set.
func (o *DictionariesDictResponse) GetWorkingTimeIntervalsOk() ([]IncludesIdName, bool) {
	if o == nil {
		return nil, false
	}
	return o.WorkingTimeIntervals, true
}

// SetWorkingTimeIntervals sets field value
func (o *DictionariesDictResponse) SetWorkingTimeIntervals(v []IncludesIdName) {
	o.WorkingTimeIntervals = v
}

// GetWorkingTimeModes returns the WorkingTimeModes field value
func (o *DictionariesDictResponse) GetWorkingTimeModes() []IncludesIdName {
	if o == nil {
		var ret []IncludesIdName
		return ret
	}

	return o.WorkingTimeModes
}

// GetWorkingTimeModesOk returns a tuple with the WorkingTimeModes field value
// and a boolean to check if the value has been set.
func (o *DictionariesDictResponse) GetWorkingTimeModesOk() ([]IncludesIdName, bool) {
	if o == nil {
		return nil, false
	}
	return o.WorkingTimeModes, true
}

// SetWorkingTimeModes sets field value
func (o *DictionariesDictResponse) SetWorkingTimeModes(v []IncludesIdName) {
	o.WorkingTimeModes = v
}

func (o DictionariesDictResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o DictionariesDictResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["applicant_comment_access_type"] = o.ApplicantCommentAccessType
	toSerialize["applicant_comments_order"] = o.ApplicantCommentsOrder
	toSerialize["applicant_negotiation_status"] = o.ApplicantNegotiationStatus
	toSerialize["business_trip_readiness"] = o.BusinessTripReadiness
	toSerialize["currency"] = o.Currency
	toSerialize["driver_license_types"] = o.DriverLicenseTypes
	toSerialize["education_level"] = o.EducationLevel
	toSerialize["employer_active_vacancies_order"] = o.EmployerActiveVacanciesOrder
	toSerialize["employer_archived_vacancies_order"] = o.EmployerArchivedVacanciesOrder
	if !IsNil(o.EmployerHiddenVacanciesOrder) {
		toSerialize["employer_hidden_vacancies_order"] = o.EmployerHiddenVacanciesOrder
	}
	toSerialize["employer_relation"] = o.EmployerRelation
	toSerialize["employer_type"] = o.EmployerType
	toSerialize["employment"] = o.Employment
	toSerialize["experience"] = o.Experience
	toSerialize["gender"] = o.Gender
	toSerialize["job_search_statuses_applicant"] = o.JobSearchStatusesApplicant
	toSerialize["job_search_statuses_employer"] = o.JobSearchStatusesEmployer
	toSerialize["language_level"] = o.LanguageLevel
	toSerialize["messaging_status"] = o.MessagingStatus
	toSerialize["negotiations_order"] = o.NegotiationsOrder
	toSerialize["negotiations_participant_type"] = o.NegotiationsParticipantType
	toSerialize["negotiations_state"] = o.NegotiationsState
	toSerialize["phone_call_status"] = o.PhoneCallStatus
	toSerialize["preferred_contact_type"] = o.PreferredContactType
	toSerialize["relocation_type"] = o.RelocationType
	toSerialize["resume_access_type"] = o.ResumeAccessType
	toSerialize["resume_contacts_site_type"] = o.ResumeContactsSiteType
	toSerialize["resume_hidden_fields"] = o.ResumeHiddenFields
	toSerialize["resume_moderation_note"] = o.ResumeModerationNote
	if !IsNil(o.ResumeSearchExperiencePeriod) {
		toSerialize["resume_search_experience_period"] = o.ResumeSearchExperiencePeriod
	}
	if !IsNil(o.ResumeSearchFields) {
		toSerialize["resume_search_fields"] = o.ResumeSearchFields
	}
	if !IsNil(o.ResumeSearchLabel) {
		toSerialize["resume_search_label"] = o.ResumeSearchLabel
	}
	if !IsNil(o.ResumeSearchLogic) {
		toSerialize["resume_search_logic"] = o.ResumeSearchLogic
	}
	if !IsNil(o.ResumeSearchOrder) {
		toSerialize["resume_search_order"] = o.ResumeSearchOrder
	}
	if !IsNil(o.ResumeSearchRelocation) {
		toSerialize["resume_search_relocation"] = o.ResumeSearchRelocation
	}
	toSerialize["resume_status"] = o.ResumeStatus
	toSerialize["schedule"] = o.Schedule
	toSerialize["travel_time"] = o.TravelTime
	toSerialize["vacancy_billing_type"] = o.VacancyBillingType
	toSerialize["vacancy_cluster"] = o.VacancyCluster
	toSerialize["vacancy_label"] = o.VacancyLabel
	toSerialize["vacancy_not_prolonged_reason"] = o.VacancyNotProlongedReason
	toSerialize["vacancy_relation"] = o.VacancyRelation
	toSerialize["vacancy_search_fields"] = o.VacancySearchFields
	toSerialize["vacancy_search_order"] = o.VacancySearchOrder
	toSerialize["vacancy_type"] = o.VacancyType
	toSerialize["working_days"] = o.WorkingDays
	toSerialize["working_time_intervals"] = o.WorkingTimeIntervals
	toSerialize["working_time_modes"] = o.WorkingTimeModes
	return toSerialize, nil
}

func (o *DictionariesDictResponse) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"applicant_comment_access_type",
		"applicant_comments_order",
		"applicant_negotiation_status",
		"business_trip_readiness",
		"currency",
		"driver_license_types",
		"education_level",
		"employer_active_vacancies_order",
		"employer_archived_vacancies_order",
		"employer_relation",
		"employer_type",
		"employment",
		"experience",
		"gender",
		"job_search_statuses_applicant",
		"job_search_statuses_employer",
		"language_level",
		"messaging_status",
		"negotiations_order",
		"negotiations_participant_type",
		"negotiations_state",
		"phone_call_status",
		"preferred_contact_type",
		"relocation_type",
		"resume_access_type",
		"resume_contacts_site_type",
		"resume_hidden_fields",
		"resume_moderation_note",
		"resume_status",
		"schedule",
		"travel_time",
		"vacancy_billing_type",
		"vacancy_cluster",
		"vacancy_label",
		"vacancy_not_prolonged_reason",
		"vacancy_relation",
		"vacancy_search_fields",
		"vacancy_search_order",
		"vacancy_type",
		"working_days",
		"working_time_intervals",
		"working_time_modes",
	}

	allProperties := make(map[string]interface{})

	err = json.Unmarshal(data, &allProperties)

	if err != nil {
		return err;
	}

	for _, requiredProperty := range(requiredProperties) {
		if _, exists := allProperties[requiredProperty]; !exists {
			return fmt.Errorf("no value given for required property %v", requiredProperty)
		}
	}

	varDictionariesDictResponse := _DictionariesDictResponse{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varDictionariesDictResponse)

	if err != nil {
		return err
	}

	*o = DictionariesDictResponse(varDictionariesDictResponse)

	return err
}

type NullableDictionariesDictResponse struct {
	value *DictionariesDictResponse
	isSet bool
}

func (v NullableDictionariesDictResponse) Get() *DictionariesDictResponse {
	return v.value
}

func (v *NullableDictionariesDictResponse) Set(val *DictionariesDictResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableDictionariesDictResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableDictionariesDictResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDictionariesDictResponse(val *DictionariesDictResponse) *NullableDictionariesDictResponse {
	return &NullableDictionariesDictResponse{value: val, isSet: true}
}

func (v NullableDictionariesDictResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDictionariesDictResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


