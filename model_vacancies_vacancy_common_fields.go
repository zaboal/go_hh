/*
HeadHunter API

По-русски | [Switch to English](https://api.hh.ru/openapi/en/redoc)  В OpenAPI ведется пока что только небольшая часть документации [Основная документация](https://github.com/hhru/api).  Для поиска по документации можно использовать Ctrl+F.  # Общая информация  * Всё API работает по протоколу HTTPS. * Авторизация осуществляется по протоколу OAuth2. * Все данные доступны только в формате JSON. * Базовый URL — `https://api.hh.ru/` * Возможны запросы к данным [любого сайта группы компаний HeadHunter](#section/Obshaya-informaciya/Vybor-sajta) * <a name=\"date-format\"></a> Даты форматируются в соответствии с [ISO 8601](http://en.wikipedia.org/wiki/ISO_8601): `YYYY-MM-DDThh:mm:ss±hhmm`.   <a name=\"request-requirements\"></a> ## Требования к запросам  В запросе необходимо передавать заголовок `User-Agent`, но если ваша реализация http клиента не позволяет, можно отправить `HH-User-Agent`. Если не отправлен ни один заголовок, то ответом будет `400 Bad Request`. Указание в заголовке названия приложения и контактной почты разработчика позволит нам оперативно с вами связаться в случае необходимости. Заголовки `User-Agent` и `HH-User-Agent` взаимозаменяемы, в случае, если вы отправите оба заголовка, обработан будет только `HH-User-Agent`.  ``` User-Agent: MyApp/1.0 (my-app-feedback@example.com) ```  Подробнее про [ошибки в заголовке User-Agent](https://github.com/hhru/api/blob/master/docs/errors.md#user-agent).   <a name=\"request-body\"></a> ## Формат тела запроса при отправке JSON  Данные, передающиеся в теле запроса, должны удовлетворять требованиям:  * Валидный JSON (допускается передача как минифицированного варианта, так и pretty print варианта с дополнительными пробелами и сбросами строк). * Рекомендуется использование кодировки UTF-8 без дополнительного экранирования (`{\"name\": \"Иванов Иван\"}`). * Также возможно использовать ascii кодировку с экранированием (`{\"name\": \"\\u0418\\u0432\\u0430\\u043d\\u043e\\u0432 \\u0418\\u0432\\u0430\\u043d\"}`). * К типам данных в определённым полях накладываются дополнительные условия, описанные в каждом конкретном методе. В JSON типами данных являются `string`, `number`, `boolean`, `null`, `object`, `array`.  ### Ответ Ответ свыше определенной длины будет сжиматься методом gzip.  ### Ошибки и коды ответов  API широко использует информирование при помощи кодов ответов. Приложение должно корректно их обрабатывать.  В случае неполадок и сбоев, возможны ответы с кодом `503` и `500`.  При каждой ошибке, помимо кода ответа, в теле ответа может быть выдана дополнительная информация, позволяющая разработчику понять причину соответствующего ответа.  [Более подробно про возможные ошибки](https://github.com/hhru/api/blob/master/docs/errors.md).   ## Недокументированные поля и параметры запросов  В ответах и параметрах API можно найти ключи, не описанные в документации. Обычно это означает, что они оставлены для совместимости со старыми версиями. Их использование не рекомендуется. Если ваше приложение использует такие ключи, перейдите на использование актуальных ключей, описанных в документации.   ## Пагинация  К любому запросу, подразумевающему выдачу списка объектов, можно в параметрах указать `page=N&per_page=M`. Нумерация идёт с нуля, по умолчанию выдаётся первая (нулевая) страница с 20 объектами на странице. Во всех ответах, где доступна пагинация, единообразный корневой объект:  ```json {   \"found\": 1,   \"per_page\": 1,   \"pages\": 1,   \"page\": 0,   \"items\": [{}] } ``` ## Выбор сайта  API HeadHunter позволяет получать данные со всех сайтов группы компании HeadHunter.  В частности:  * hh.ru * rabota.by * hh1.az * hh.uz * hh.kz * headhunter.ge * headhunter.kg  Запросы к данным на всех сайтах следует направлять на `https://api.hh.ru/`.  При необходимости учесть специфику сайта, можно добавить в запрос параметр `?host=`. По умолчанию используется `hh.ru`.  Например, для получения [локализаций](https://api.hh.ru/openapi/redoc#tag/Obshie-spravochniki/operation/get-locales), доступных на hh.kz необходимо сделать GET запрос на `https://api.hh.ru/locales?host=hh.kz`.  ## CORS (Cross-Origin Resource Sharing)  API поддерживает технологию CORS для запроса данных из браузера с произвольного домена. Этот метод более предпочтителен, чем использование JSONP. Он не ограничен методом GET. Для отладки CORS доступен [специальный метод](https://github.com/hhru/api/blob/master/docs/cors.md). Для использования JSONP передайте параметр `?callback=callback_name`.  * [CORS specification on w3.org](http://www.w3.org/TR/cors/) * [HTML5Rocks CORS Tutorial](http://www.html5rocks.com/en/tutorials/cors/) * [CORS on dev.opera.com](http://dev.opera.com/articles/view/dom-access-control-using-cross-origin-resource-sharing/) * [CORS on caniuse.com](http://caniuse.com/#feat=cors) * [CORS on en.wikipedia.org](http://en.wikipedia.org/wiki/Cross-origin_resource_sharing)   ## Внешние ссылки на статьи и стандарты  * [HTTP/1.1](http://tools.ietf.org/html/rfc2616) * [JSON](http://json.org/) * [URI Template](http://tools.ietf.org/html/rfc6570) * [OAuth 2.0](http://tools.ietf.org/html/rfc6749) * [REST](http://www.ics.uci.edu/~fielding/pubs/dissertation/rest_arch_style.htm) * [ISO 8601](http://en.wikipedia.org/wiki/ISO_8601)  # Авторизация  API поддерживает следующие уровни авторизации:   - [авторизация приложения](#tag/Avtorizaciya-prilozheniya)   - [авторизация пользователя](#section/Avtorizaciya/Avtorizaciya-polzovatelya)  * [Авторизация пользователя](#section/Avtorizaciya/Avtorizaciya-polzovatelya)   * [Правила формирования специального redirect_uri](#section/Avtorizaciya/Pravila-formirovaniya-specialnogo-redirect_uri)   * [Процесс авторизации](#section/Avtorizaciya/Process-avtorizacii)   * [Успешное получение временного `authorization_code`](#get-authorization_code)   * [Получение access и refresh токенов](#section/Avtorizaciya/Poluchenie-access-i-refresh-tokenov) * [Обновление пары access и refresh токенов](#section/Avtorizaciya/Obnovlenie-pary-access-i-refresh-tokenov) * [Инвалидация токена](#tag/Avtorizaciya-rabotodatelya/operation/invalidate-token) * [Запрос авторизации под другим пользователем](#section/Avtorizaciya/Zapros-avtorizacii-pod-drugim-polzovatelem) * [Авторизация под разными рабочими аккаунтами](#section/Avtorizaciya/Avtorizaciya-pod-raznymi-rabochimi-akkauntami) * [Авторизация приложения](#tag/Avtorizaciya-prilozheniya)       ## Авторизация пользователя Для выполнения запросов от имени пользователя необходимо пользоваться токеном пользователя.  В начале приложению необходимо направить пользователя (открыть страницу) по адресу:  ``` https://hh.ru/oauth/authorize? response_type=code& client_id={client_id}& state={state}& redirect_uri={redirect_uri} ```  Обязательные параметры:  * `response_type=code` — указание на способ получения авторизации, используя `authorization code` * `client_id` — идентификатор, полученный при создании приложения   Необязательные параметры:  * `state` — в случае указания, будет включен в ответный редирект. Это позволяет исключить возможность взлома путём подделки межсайтовых запросов. Подробнее об этом: [RFC 6749. Section 10.12](http://tools.ietf.org/html/rfc6749#section-10.12) * `redirect_uri` — uri для перенаправления пользователя после авторизации. Если не указать, используется из настроек приложения. При наличии происходит валидация значения. Вероятнее всего, потребуется сделать urlencode значения параметра.  ## Правила формирования специального redirect_uri  К примеру, если в настройках сохранен `http://example.com/oauth`, то разрешено указывать:  * `http://www.example.com/oauth` — поддомен; * `http://www.example.com/oauth/sub/path` — уточнение пути; * `http://example.com/oauth?lang=RU` — дополнительный параметр; * `http://www.example.com/oauth/sub/path?lang=RU` — всё вместе.  Запрещено:  * `https://example.com/oauth` — различные протоколы; * `http://wwwexample.com/oauth` — различные домены; * `http://wwwexample.com/` — другой путь; * `http://example.com/oauths` — другой путь; * `http://example.com:80/oauths` — указание изначально отсутствующего порта;  ## Процесс авторизации  Если пользователь не авторизован на сайте, ему будет показана форма авторизации на сайте. После прохождения авторизации на сайте, пользователю будет выведена форма с запросом разрешения доступа вашего приложения к его персональным данным.  Если пользователь не разрешает доступ приложению, пользователь будет перенаправлен на указанный `redirect_uri` с `?error=access_denied` и `state={state}`, если таковой был указан при первом запросе.  <a name=\"get-authorization_code\"></a> ### Успешное получение временного `authorization_code`  В случае разрешения прав, в редиректе будет указан временный `authorization_code`:  ```http HTTP/1.1 302 FOUND Location: {redirect_uri}?code={authorization_code} ```  Если пользователь авторизован на сайте и доступ данному приложению однажды ранее выдан, ответом будет сразу вышеописанный редирект с `authorization_code` (без показа формы логина и выдачи прав).  ## Получение access и refresh токенов  После получения `authorization_code` приложению необходимо сделать сервер-сервер запрос `POST https://api.hh.ru/token` для обмена полученного `authorization_code` на `access_token` (старый запрос `POST https://hh.ru/oauth/token` считается устаревшим).  В теле запроса необходимо передать [дополнительные параметры](#required_parameters).  Тело запроса необходимо передавать в стандартном `application/x-www-form-urlencoded` с указанием соответствующего заголовка `Content-Type`.  `authorization_code` имеет довольно короткий срок жизни, при его истечении необходимо запросить новый.  ## Обновление пары access и refresh токенов `access_token` также имеет срок жизни (ключ `expires_in`, в секундах), при его истечении приложение должно сделать запрос с `refresh_token` для получения нового.  Запрос необходимо делать в `application/x-www-form-urlencoded`.  ``` POST https://api.hh.ru/token ```  (старый запрос `POST https://hh.ru/oauth/token` считается устаревшим)  В теле запроса необходимо передать [дополнительные параметры](#required_parameters)  `refresh_token` можно использовать только один раз и только по истечению срока действия `access_token`.  После получения новой пары access и refresh токенов, их необходимо использовать в дальнейших запросах в api и запросах на продление токена.  ## Запрос авторизации под другим пользователем  Возможен следующий сценарий:  1. Приложение перенаправляет пользователя на сайт с запросом авторизации. 2. Пользователь на сайте уже авторизован и данному приложение доступ уже был разрешен. 3. Пользователю будет предложена возможность продолжить работу под текущим аккаунтом, либо зайти под другим аккаунтом.  Если есть необходимость, чтобы на шаге 3 сразу происходило перенаправление (redirect) с временным токеном, необходимо добавить к запросу `/oauth/authorize...` параметр `skip_choose_account=true`. В этом случае автоматически выдаётся доступ пользователю авторизованному на сайте.  Если есть необходимость всегда показывать форму авторизации, приложение может добавить к запросу `/oauth/authorize...` параметр `force_login=true`. В этом случае, пользователю будет показана форма авторизации с логином и паролем даже в случае, если пользователь уже авторизован.  Это может быть полезно приложениям, которые предоставляют сервис только для соискателей. Если пришел пользователь-работодатель, приложение может предложить пользователю повторно разрешить доступ на сайте, уже указав другую учетную запись.  Также, после авторизации приложение может показать пользователю сообщение:  ``` Вы вошли как %Имя_Фамилия%. Это не вы? ``` и предоставить ссылку с `force_login=true` для возможности захода под другим логином.  ## Авторизация под разными рабочими аккаунтами  Для получения списка рабочих аккаунтов менеджера и для работы под разными рабочими аккаунтами менеджера необходимо прочитать документацию по [рабочим аккаунтам менеджера](#tag/Menedzhery-rabotodatelya/operation/get-manager-accounts)  В случае компрометации токена необходимо [инвалидировать](#tag/Avtorizaciya-rabotodatelya/operation/invalidate-token) скомпрометированный токен и запросить токен заново!  <!-- ReDoc-Inject: <security-definitions> --> 

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

// checks if the VacanciesVacancyCommonFields type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &VacanciesVacancyCommonFields{}

// VacanciesVacancyCommonFields struct for VacanciesVacancyCommonFields
type VacanciesVacancyCommonFields struct {
	// Указание, что вакансия доступна для соискателей с инвалидностью
	AcceptHandicapped bool `json:"accept_handicapped"`
	// Разрешен ли отклик на вакансию неполным резюме
	AcceptIncompleteResumes bool `json:"accept_incomplete_resumes"`
	// Указание, что вакансия доступна для соискателей старше 14 лет [подробнее](https://github.com/hhru/api/blob/master/docs/employer_vacancies_accept_kids.md#accept-kids)
	AcceptKids bool `json:"accept_kids"`
	// Указание, что вакансия доступна с временным трудоустройством
	AcceptTemporary NullableBool `json:"accept_temporary,omitempty"`
	// Возможность [переписки с кандидатами](https://inboxemp.hh.ru/) по данной вакансии
	AllowMessages bool `json:"allow_messages"`
	// Ссылка на представление вакансии на сайте
	AlternateUrl string `json:"alternate_url"`
	// Ссылка на отклик на вакансию на сайте
	ApplyAlternateUrl string `json:"apply_alternate_url"`
	// Прошла ли вакансия модерацию
	Approved bool `json:"approved"`
	// Находится ли данная вакансия в архиве
	Archived bool `json:"archived"`
	Area IncludesArea `json:"area"`
	BillingType NullableVacancyBillingTypeOutput `json:"billing_type"`
	// Строка с кодом HTML (возможно наличие `<script/>` и `<style/>`), которая является альтернативой стандартному описанию вакансии.  HTML адаптирован для мобильных устройств и корректно отображается без поддержки JavaScript. При этом:  * Содержимое растягивается на 100% ширины контейнера и умещается в 300px без прокрутки. * Содержимое рассчитано на то, что оно будет вставлено в обвязку, в которую входит название, требуемый опыт работы, регион, тип занятости и рабочего дня вакансии, а также ссылка на компанию, опубликовавшую вакансию. * Изображения, которые могут встретиться в таком описании, адаптированы под Retina-дисплеи. * Размер шрифта не меньше 12px, размер межстрочного интервала не меньше 16px.  Значение может быть `null`, если у вакансии отсутствует индивидуальное описание
	BrandedDescription *string `json:"branded_description,omitempty"`
	// Внутренний код вакансии
	Code NullableString `json:"code,omitempty"`
	Contacts NullableVacancyContactsOutput `json:"contacts,omitempty"`
	// Дата и время публикации вакансии
	// Deprecated
	CreatedAt string `json:"created_at"`
	Department NullableVacancyDepartmentOutput `json:"department,omitempty"`
	// Описание в html, не менее 200 символов
	Description string `json:"description"`
	// Список требуемых категорий водительских прав
	DriverLicenseTypes []VacancyDriverLicenceTypeItem `json:"driver_license_types"`
	Employer NullableVacanciesVacancyEmployer `json:"employer,omitempty"`
	Employment NullableVacancyEmploymentOutput `json:"employment,omitempty"`
	Experience NullableVacancyExperienceOutput `json:"experience"`
	// Информация о наличии прикрепленного тестового задании к вакансии
	HasTest bool `json:"has_test"`
	// Идентификатор вакансии
	Id string `json:"id"`
	// Дата и время создания вакансии
	InitialCreatedAt string `json:"initial_created_at"`
	InsiderInterview NullableVacancyInsiderInterview `json:"insider_interview,omitempty"`
	// Список ключевых навыков, не более 30
	KeySkills []VacancyKeySkillItem `json:"key_skills"`
	Languages []string `json:"languages,omitempty"`
	// Название
	Name string `json:"name"`
	// Ссылка для получения списка откликов/приглашений
	NegotiationsUrl NullableString `json:"negotiations_url,omitempty"`
	// Является ли данная вакансия премиум-вакансией
	Premium bool `json:"premium"`
	// Список профессиональных ролей
	ProfessionalRoles []VacancyProfessionalRoleItemOutput `json:"professional_roles"`
	// Дата и время публикации вакансии
	PublishedAt string `json:"published_at"`
	// Возвращает связи соискателя с вакансией. Значения из поля `vacancy_relation` в [справочнике полей](#tag/Obshie-spravochniki/operation/get-dictionaries)
	Relations []VacancyRelationItem `json:"relations,omitempty"`
	// Обязательно ли заполнять сообщение при отклике на вакансию
	ResponseLetterRequired bool `json:"response_letter_required"`
	// URL отклика для прямых вакансий (`type.id=direct`)
	ResponseUrl NullableString `json:"response_url,omitempty"`
	Salary NullableVacancySalary `json:"salary,omitempty"`
	Schedule VacancyScheduleOutput `json:"schedule"`
	// Подходящие резюме на вакансию
	SuitableResumesUrl NullableString `json:"suitable_resumes_url,omitempty"`
	Test NullableVacancyDraftTest `json:"test,omitempty"`
	// Идентификатор типа вакансии из справочника [`vacancy_type`](https://api.hh.ru/openapi/redoc#tag/Obshie-spravochniki/operation/get-dictionaries)
	Type IncludesIdName `json:"type"`
	VacancyConstructorTemplate NullableVacancyVacancyConstructorTemplate `json:"vacancy_constructor_template,omitempty"`
	VideoVacancy NullableVacancyVideoVacancy `json:"video_vacancy,omitempty"`
	// Список рабочих дней
	WorkingDays []VacancyWorkingDayItemOutput `json:"working_days,omitempty"`
	// Список с временными интервалами работы
	WorkingTimeIntervals []VacancyWorkingTimeIntervalItemOutput `json:"working_time_intervals,omitempty"`
	// Список режимов времени работы
	WorkingTimeModes []VacancyWorkingTimeModeItemOutput `json:"working_time_modes,omitempty"`
}

type _VacanciesVacancyCommonFields VacanciesVacancyCommonFields

// NewVacanciesVacancyCommonFields instantiates a new VacanciesVacancyCommonFields object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewVacanciesVacancyCommonFields(acceptHandicapped bool, acceptIncompleteResumes bool, acceptKids bool, allowMessages bool, alternateUrl string, applyAlternateUrl string, approved bool, archived bool, area IncludesArea, billingType NullableVacancyBillingTypeOutput, createdAt string, description string, driverLicenseTypes []VacancyDriverLicenceTypeItem, experience NullableVacancyExperienceOutput, hasTest bool, id string, initialCreatedAt string, keySkills []VacancyKeySkillItem, name string, premium bool, professionalRoles []VacancyProfessionalRoleItemOutput, publishedAt string, responseLetterRequired bool, schedule VacancyScheduleOutput, type_ IncludesIdName) *VacanciesVacancyCommonFields {
	this := VacanciesVacancyCommonFields{}
	this.AcceptHandicapped = acceptHandicapped
	this.AcceptIncompleteResumes = acceptIncompleteResumes
	this.AcceptKids = acceptKids
	this.AllowMessages = allowMessages
	this.AlternateUrl = alternateUrl
	this.ApplyAlternateUrl = applyAlternateUrl
	this.Approved = approved
	this.Archived = archived
	this.Area = area
	this.BillingType = billingType
	this.CreatedAt = createdAt
	this.Description = description
	this.DriverLicenseTypes = driverLicenseTypes
	this.Experience = experience
	this.HasTest = hasTest
	this.Id = id
	this.InitialCreatedAt = initialCreatedAt
	this.KeySkills = keySkills
	this.Name = name
	this.Premium = premium
	this.ProfessionalRoles = professionalRoles
	this.PublishedAt = publishedAt
	this.ResponseLetterRequired = responseLetterRequired
	this.Schedule = schedule
	this.Type = type_
	return &this
}

// NewVacanciesVacancyCommonFieldsWithDefaults instantiates a new VacanciesVacancyCommonFields object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewVacanciesVacancyCommonFieldsWithDefaults() *VacanciesVacancyCommonFields {
	this := VacanciesVacancyCommonFields{}
	return &this
}

// GetAcceptHandicapped returns the AcceptHandicapped field value
func (o *VacanciesVacancyCommonFields) GetAcceptHandicapped() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.AcceptHandicapped
}

// GetAcceptHandicappedOk returns a tuple with the AcceptHandicapped field value
// and a boolean to check if the value has been set.
func (o *VacanciesVacancyCommonFields) GetAcceptHandicappedOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AcceptHandicapped, true
}

// SetAcceptHandicapped sets field value
func (o *VacanciesVacancyCommonFields) SetAcceptHandicapped(v bool) {
	o.AcceptHandicapped = v
}

// GetAcceptIncompleteResumes returns the AcceptIncompleteResumes field value
func (o *VacanciesVacancyCommonFields) GetAcceptIncompleteResumes() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.AcceptIncompleteResumes
}

// GetAcceptIncompleteResumesOk returns a tuple with the AcceptIncompleteResumes field value
// and a boolean to check if the value has been set.
func (o *VacanciesVacancyCommonFields) GetAcceptIncompleteResumesOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AcceptIncompleteResumes, true
}

// SetAcceptIncompleteResumes sets field value
func (o *VacanciesVacancyCommonFields) SetAcceptIncompleteResumes(v bool) {
	o.AcceptIncompleteResumes = v
}

// GetAcceptKids returns the AcceptKids field value
func (o *VacanciesVacancyCommonFields) GetAcceptKids() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.AcceptKids
}

// GetAcceptKidsOk returns a tuple with the AcceptKids field value
// and a boolean to check if the value has been set.
func (o *VacanciesVacancyCommonFields) GetAcceptKidsOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AcceptKids, true
}

// SetAcceptKids sets field value
func (o *VacanciesVacancyCommonFields) SetAcceptKids(v bool) {
	o.AcceptKids = v
}

// GetAcceptTemporary returns the AcceptTemporary field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *VacanciesVacancyCommonFields) GetAcceptTemporary() bool {
	if o == nil || IsNil(o.AcceptTemporary.Get()) {
		var ret bool
		return ret
	}
	return *o.AcceptTemporary.Get()
}

// GetAcceptTemporaryOk returns a tuple with the AcceptTemporary field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *VacanciesVacancyCommonFields) GetAcceptTemporaryOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return o.AcceptTemporary.Get(), o.AcceptTemporary.IsSet()
}

// HasAcceptTemporary returns a boolean if a field has been set.
func (o *VacanciesVacancyCommonFields) HasAcceptTemporary() bool {
	if o != nil && o.AcceptTemporary.IsSet() {
		return true
	}

	return false
}

// SetAcceptTemporary gets a reference to the given NullableBool and assigns it to the AcceptTemporary field.
func (o *VacanciesVacancyCommonFields) SetAcceptTemporary(v bool) {
	o.AcceptTemporary.Set(&v)
}
// SetAcceptTemporaryNil sets the value for AcceptTemporary to be an explicit nil
func (o *VacanciesVacancyCommonFields) SetAcceptTemporaryNil() {
	o.AcceptTemporary.Set(nil)
}

// UnsetAcceptTemporary ensures that no value is present for AcceptTemporary, not even an explicit nil
func (o *VacanciesVacancyCommonFields) UnsetAcceptTemporary() {
	o.AcceptTemporary.Unset()
}

// GetAllowMessages returns the AllowMessages field value
func (o *VacanciesVacancyCommonFields) GetAllowMessages() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.AllowMessages
}

// GetAllowMessagesOk returns a tuple with the AllowMessages field value
// and a boolean to check if the value has been set.
func (o *VacanciesVacancyCommonFields) GetAllowMessagesOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AllowMessages, true
}

// SetAllowMessages sets field value
func (o *VacanciesVacancyCommonFields) SetAllowMessages(v bool) {
	o.AllowMessages = v
}

// GetAlternateUrl returns the AlternateUrl field value
func (o *VacanciesVacancyCommonFields) GetAlternateUrl() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.AlternateUrl
}

// GetAlternateUrlOk returns a tuple with the AlternateUrl field value
// and a boolean to check if the value has been set.
func (o *VacanciesVacancyCommonFields) GetAlternateUrlOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AlternateUrl, true
}

// SetAlternateUrl sets field value
func (o *VacanciesVacancyCommonFields) SetAlternateUrl(v string) {
	o.AlternateUrl = v
}

// GetApplyAlternateUrl returns the ApplyAlternateUrl field value
func (o *VacanciesVacancyCommonFields) GetApplyAlternateUrl() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ApplyAlternateUrl
}

// GetApplyAlternateUrlOk returns a tuple with the ApplyAlternateUrl field value
// and a boolean to check if the value has been set.
func (o *VacanciesVacancyCommonFields) GetApplyAlternateUrlOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ApplyAlternateUrl, true
}

// SetApplyAlternateUrl sets field value
func (o *VacanciesVacancyCommonFields) SetApplyAlternateUrl(v string) {
	o.ApplyAlternateUrl = v
}

// GetApproved returns the Approved field value
func (o *VacanciesVacancyCommonFields) GetApproved() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.Approved
}

// GetApprovedOk returns a tuple with the Approved field value
// and a boolean to check if the value has been set.
func (o *VacanciesVacancyCommonFields) GetApprovedOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Approved, true
}

// SetApproved sets field value
func (o *VacanciesVacancyCommonFields) SetApproved(v bool) {
	o.Approved = v
}

// GetArchived returns the Archived field value
func (o *VacanciesVacancyCommonFields) GetArchived() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.Archived
}

// GetArchivedOk returns a tuple with the Archived field value
// and a boolean to check if the value has been set.
func (o *VacanciesVacancyCommonFields) GetArchivedOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Archived, true
}

// SetArchived sets field value
func (o *VacanciesVacancyCommonFields) SetArchived(v bool) {
	o.Archived = v
}

// GetArea returns the Area field value
func (o *VacanciesVacancyCommonFields) GetArea() IncludesArea {
	if o == nil {
		var ret IncludesArea
		return ret
	}

	return o.Area
}

// GetAreaOk returns a tuple with the Area field value
// and a boolean to check if the value has been set.
func (o *VacanciesVacancyCommonFields) GetAreaOk() (*IncludesArea, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Area, true
}

// SetArea sets field value
func (o *VacanciesVacancyCommonFields) SetArea(v IncludesArea) {
	o.Area = v
}

// GetBillingType returns the BillingType field value
// If the value is explicit nil, the zero value for VacancyBillingTypeOutput will be returned
func (o *VacanciesVacancyCommonFields) GetBillingType() VacancyBillingTypeOutput {
	if o == nil || o.BillingType.Get() == nil {
		var ret VacancyBillingTypeOutput
		return ret
	}

	return *o.BillingType.Get()
}

// GetBillingTypeOk returns a tuple with the BillingType field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *VacanciesVacancyCommonFields) GetBillingTypeOk() (*VacancyBillingTypeOutput, bool) {
	if o == nil {
		return nil, false
	}
	return o.BillingType.Get(), o.BillingType.IsSet()
}

// SetBillingType sets field value
func (o *VacanciesVacancyCommonFields) SetBillingType(v VacancyBillingTypeOutput) {
	o.BillingType.Set(&v)
}

// GetBrandedDescription returns the BrandedDescription field value if set, zero value otherwise.
func (o *VacanciesVacancyCommonFields) GetBrandedDescription() string {
	if o == nil || IsNil(o.BrandedDescription) {
		var ret string
		return ret
	}
	return *o.BrandedDescription
}

// GetBrandedDescriptionOk returns a tuple with the BrandedDescription field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VacanciesVacancyCommonFields) GetBrandedDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.BrandedDescription) {
		return nil, false
	}
	return o.BrandedDescription, true
}

// HasBrandedDescription returns a boolean if a field has been set.
func (o *VacanciesVacancyCommonFields) HasBrandedDescription() bool {
	if o != nil && !IsNil(o.BrandedDescription) {
		return true
	}

	return false
}

// SetBrandedDescription gets a reference to the given string and assigns it to the BrandedDescription field.
func (o *VacanciesVacancyCommonFields) SetBrandedDescription(v string) {
	o.BrandedDescription = &v
}

// GetCode returns the Code field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *VacanciesVacancyCommonFields) GetCode() string {
	if o == nil || IsNil(o.Code.Get()) {
		var ret string
		return ret
	}
	return *o.Code.Get()
}

// GetCodeOk returns a tuple with the Code field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *VacanciesVacancyCommonFields) GetCodeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Code.Get(), o.Code.IsSet()
}

// HasCode returns a boolean if a field has been set.
func (o *VacanciesVacancyCommonFields) HasCode() bool {
	if o != nil && o.Code.IsSet() {
		return true
	}

	return false
}

// SetCode gets a reference to the given NullableString and assigns it to the Code field.
func (o *VacanciesVacancyCommonFields) SetCode(v string) {
	o.Code.Set(&v)
}
// SetCodeNil sets the value for Code to be an explicit nil
func (o *VacanciesVacancyCommonFields) SetCodeNil() {
	o.Code.Set(nil)
}

// UnsetCode ensures that no value is present for Code, not even an explicit nil
func (o *VacanciesVacancyCommonFields) UnsetCode() {
	o.Code.Unset()
}

// GetContacts returns the Contacts field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *VacanciesVacancyCommonFields) GetContacts() VacancyContactsOutput {
	if o == nil || IsNil(o.Contacts.Get()) {
		var ret VacancyContactsOutput
		return ret
	}
	return *o.Contacts.Get()
}

// GetContactsOk returns a tuple with the Contacts field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *VacanciesVacancyCommonFields) GetContactsOk() (*VacancyContactsOutput, bool) {
	if o == nil {
		return nil, false
	}
	return o.Contacts.Get(), o.Contacts.IsSet()
}

// HasContacts returns a boolean if a field has been set.
func (o *VacanciesVacancyCommonFields) HasContacts() bool {
	if o != nil && o.Contacts.IsSet() {
		return true
	}

	return false
}

// SetContacts gets a reference to the given NullableVacancyContactsOutput and assigns it to the Contacts field.
func (o *VacanciesVacancyCommonFields) SetContacts(v VacancyContactsOutput) {
	o.Contacts.Set(&v)
}
// SetContactsNil sets the value for Contacts to be an explicit nil
func (o *VacanciesVacancyCommonFields) SetContactsNil() {
	o.Contacts.Set(nil)
}

// UnsetContacts ensures that no value is present for Contacts, not even an explicit nil
func (o *VacanciesVacancyCommonFields) UnsetContacts() {
	o.Contacts.Unset()
}

// GetCreatedAt returns the CreatedAt field value
// Deprecated
func (o *VacanciesVacancyCommonFields) GetCreatedAt() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value
// and a boolean to check if the value has been set.
// Deprecated
func (o *VacanciesVacancyCommonFields) GetCreatedAtOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CreatedAt, true
}

// SetCreatedAt sets field value
// Deprecated
func (o *VacanciesVacancyCommonFields) SetCreatedAt(v string) {
	o.CreatedAt = v
}

// GetDepartment returns the Department field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *VacanciesVacancyCommonFields) GetDepartment() VacancyDepartmentOutput {
	if o == nil || IsNil(o.Department.Get()) {
		var ret VacancyDepartmentOutput
		return ret
	}
	return *o.Department.Get()
}

// GetDepartmentOk returns a tuple with the Department field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *VacanciesVacancyCommonFields) GetDepartmentOk() (*VacancyDepartmentOutput, bool) {
	if o == nil {
		return nil, false
	}
	return o.Department.Get(), o.Department.IsSet()
}

// HasDepartment returns a boolean if a field has been set.
func (o *VacanciesVacancyCommonFields) HasDepartment() bool {
	if o != nil && o.Department.IsSet() {
		return true
	}

	return false
}

// SetDepartment gets a reference to the given NullableVacancyDepartmentOutput and assigns it to the Department field.
func (o *VacanciesVacancyCommonFields) SetDepartment(v VacancyDepartmentOutput) {
	o.Department.Set(&v)
}
// SetDepartmentNil sets the value for Department to be an explicit nil
func (o *VacanciesVacancyCommonFields) SetDepartmentNil() {
	o.Department.Set(nil)
}

// UnsetDepartment ensures that no value is present for Department, not even an explicit nil
func (o *VacanciesVacancyCommonFields) UnsetDepartment() {
	o.Department.Unset()
}

// GetDescription returns the Description field value
func (o *VacanciesVacancyCommonFields) GetDescription() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Description
}

// GetDescriptionOk returns a tuple with the Description field value
// and a boolean to check if the value has been set.
func (o *VacanciesVacancyCommonFields) GetDescriptionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Description, true
}

// SetDescription sets field value
func (o *VacanciesVacancyCommonFields) SetDescription(v string) {
	o.Description = v
}

// GetDriverLicenseTypes returns the DriverLicenseTypes field value
func (o *VacanciesVacancyCommonFields) GetDriverLicenseTypes() []VacancyDriverLicenceTypeItem {
	if o == nil {
		var ret []VacancyDriverLicenceTypeItem
		return ret
	}

	return o.DriverLicenseTypes
}

// GetDriverLicenseTypesOk returns a tuple with the DriverLicenseTypes field value
// and a boolean to check if the value has been set.
func (o *VacanciesVacancyCommonFields) GetDriverLicenseTypesOk() ([]VacancyDriverLicenceTypeItem, bool) {
	if o == nil {
		return nil, false
	}
	return o.DriverLicenseTypes, true
}

// SetDriverLicenseTypes sets field value
func (o *VacanciesVacancyCommonFields) SetDriverLicenseTypes(v []VacancyDriverLicenceTypeItem) {
	o.DriverLicenseTypes = v
}

// GetEmployer returns the Employer field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *VacanciesVacancyCommonFields) GetEmployer() VacanciesVacancyEmployer {
	if o == nil || IsNil(o.Employer.Get()) {
		var ret VacanciesVacancyEmployer
		return ret
	}
	return *o.Employer.Get()
}

// GetEmployerOk returns a tuple with the Employer field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *VacanciesVacancyCommonFields) GetEmployerOk() (*VacanciesVacancyEmployer, bool) {
	if o == nil {
		return nil, false
	}
	return o.Employer.Get(), o.Employer.IsSet()
}

// HasEmployer returns a boolean if a field has been set.
func (o *VacanciesVacancyCommonFields) HasEmployer() bool {
	if o != nil && o.Employer.IsSet() {
		return true
	}

	return false
}

// SetEmployer gets a reference to the given NullableVacanciesVacancyEmployer and assigns it to the Employer field.
func (o *VacanciesVacancyCommonFields) SetEmployer(v VacanciesVacancyEmployer) {
	o.Employer.Set(&v)
}
// SetEmployerNil sets the value for Employer to be an explicit nil
func (o *VacanciesVacancyCommonFields) SetEmployerNil() {
	o.Employer.Set(nil)
}

// UnsetEmployer ensures that no value is present for Employer, not even an explicit nil
func (o *VacanciesVacancyCommonFields) UnsetEmployer() {
	o.Employer.Unset()
}

// GetEmployment returns the Employment field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *VacanciesVacancyCommonFields) GetEmployment() VacancyEmploymentOutput {
	if o == nil || IsNil(o.Employment.Get()) {
		var ret VacancyEmploymentOutput
		return ret
	}
	return *o.Employment.Get()
}

// GetEmploymentOk returns a tuple with the Employment field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *VacanciesVacancyCommonFields) GetEmploymentOk() (*VacancyEmploymentOutput, bool) {
	if o == nil {
		return nil, false
	}
	return o.Employment.Get(), o.Employment.IsSet()
}

// HasEmployment returns a boolean if a field has been set.
func (o *VacanciesVacancyCommonFields) HasEmployment() bool {
	if o != nil && o.Employment.IsSet() {
		return true
	}

	return false
}

// SetEmployment gets a reference to the given NullableVacancyEmploymentOutput and assigns it to the Employment field.
func (o *VacanciesVacancyCommonFields) SetEmployment(v VacancyEmploymentOutput) {
	o.Employment.Set(&v)
}
// SetEmploymentNil sets the value for Employment to be an explicit nil
func (o *VacanciesVacancyCommonFields) SetEmploymentNil() {
	o.Employment.Set(nil)
}

// UnsetEmployment ensures that no value is present for Employment, not even an explicit nil
func (o *VacanciesVacancyCommonFields) UnsetEmployment() {
	o.Employment.Unset()
}

// GetExperience returns the Experience field value
// If the value is explicit nil, the zero value for VacancyExperienceOutput will be returned
func (o *VacanciesVacancyCommonFields) GetExperience() VacancyExperienceOutput {
	if o == nil || o.Experience.Get() == nil {
		var ret VacancyExperienceOutput
		return ret
	}

	return *o.Experience.Get()
}

// GetExperienceOk returns a tuple with the Experience field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *VacanciesVacancyCommonFields) GetExperienceOk() (*VacancyExperienceOutput, bool) {
	if o == nil {
		return nil, false
	}
	return o.Experience.Get(), o.Experience.IsSet()
}

// SetExperience sets field value
func (o *VacanciesVacancyCommonFields) SetExperience(v VacancyExperienceOutput) {
	o.Experience.Set(&v)
}

// GetHasTest returns the HasTest field value
func (o *VacanciesVacancyCommonFields) GetHasTest() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.HasTest
}

// GetHasTestOk returns a tuple with the HasTest field value
// and a boolean to check if the value has been set.
func (o *VacanciesVacancyCommonFields) GetHasTestOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.HasTest, true
}

// SetHasTest sets field value
func (o *VacanciesVacancyCommonFields) SetHasTest(v bool) {
	o.HasTest = v
}

// GetId returns the Id field value
func (o *VacanciesVacancyCommonFields) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *VacanciesVacancyCommonFields) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *VacanciesVacancyCommonFields) SetId(v string) {
	o.Id = v
}

// GetInitialCreatedAt returns the InitialCreatedAt field value
func (o *VacanciesVacancyCommonFields) GetInitialCreatedAt() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.InitialCreatedAt
}

// GetInitialCreatedAtOk returns a tuple with the InitialCreatedAt field value
// and a boolean to check if the value has been set.
func (o *VacanciesVacancyCommonFields) GetInitialCreatedAtOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.InitialCreatedAt, true
}

// SetInitialCreatedAt sets field value
func (o *VacanciesVacancyCommonFields) SetInitialCreatedAt(v string) {
	o.InitialCreatedAt = v
}

// GetInsiderInterview returns the InsiderInterview field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *VacanciesVacancyCommonFields) GetInsiderInterview() VacancyInsiderInterview {
	if o == nil || IsNil(o.InsiderInterview.Get()) {
		var ret VacancyInsiderInterview
		return ret
	}
	return *o.InsiderInterview.Get()
}

// GetInsiderInterviewOk returns a tuple with the InsiderInterview field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *VacanciesVacancyCommonFields) GetInsiderInterviewOk() (*VacancyInsiderInterview, bool) {
	if o == nil {
		return nil, false
	}
	return o.InsiderInterview.Get(), o.InsiderInterview.IsSet()
}

// HasInsiderInterview returns a boolean if a field has been set.
func (o *VacanciesVacancyCommonFields) HasInsiderInterview() bool {
	if o != nil && o.InsiderInterview.IsSet() {
		return true
	}

	return false
}

// SetInsiderInterview gets a reference to the given NullableVacancyInsiderInterview and assigns it to the InsiderInterview field.
func (o *VacanciesVacancyCommonFields) SetInsiderInterview(v VacancyInsiderInterview) {
	o.InsiderInterview.Set(&v)
}
// SetInsiderInterviewNil sets the value for InsiderInterview to be an explicit nil
func (o *VacanciesVacancyCommonFields) SetInsiderInterviewNil() {
	o.InsiderInterview.Set(nil)
}

// UnsetInsiderInterview ensures that no value is present for InsiderInterview, not even an explicit nil
func (o *VacanciesVacancyCommonFields) UnsetInsiderInterview() {
	o.InsiderInterview.Unset()
}

// GetKeySkills returns the KeySkills field value
func (o *VacanciesVacancyCommonFields) GetKeySkills() []VacancyKeySkillItem {
	if o == nil {
		var ret []VacancyKeySkillItem
		return ret
	}

	return o.KeySkills
}

// GetKeySkillsOk returns a tuple with the KeySkills field value
// and a boolean to check if the value has been set.
func (o *VacanciesVacancyCommonFields) GetKeySkillsOk() ([]VacancyKeySkillItem, bool) {
	if o == nil {
		return nil, false
	}
	return o.KeySkills, true
}

// SetKeySkills sets field value
func (o *VacanciesVacancyCommonFields) SetKeySkills(v []VacancyKeySkillItem) {
	o.KeySkills = v
}

// GetLanguages returns the Languages field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *VacanciesVacancyCommonFields) GetLanguages() []string {
	if o == nil {
		var ret []string
		return ret
	}
	return o.Languages
}

// GetLanguagesOk returns a tuple with the Languages field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *VacanciesVacancyCommonFields) GetLanguagesOk() ([]string, bool) {
	if o == nil || IsNil(o.Languages) {
		return nil, false
	}
	return o.Languages, true
}

// HasLanguages returns a boolean if a field has been set.
func (o *VacanciesVacancyCommonFields) HasLanguages() bool {
	if o != nil && !IsNil(o.Languages) {
		return true
	}

	return false
}

// SetLanguages gets a reference to the given []string and assigns it to the Languages field.
func (o *VacanciesVacancyCommonFields) SetLanguages(v []string) {
	o.Languages = v
}

// GetName returns the Name field value
func (o *VacanciesVacancyCommonFields) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *VacanciesVacancyCommonFields) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *VacanciesVacancyCommonFields) SetName(v string) {
	o.Name = v
}

// GetNegotiationsUrl returns the NegotiationsUrl field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *VacanciesVacancyCommonFields) GetNegotiationsUrl() string {
	if o == nil || IsNil(o.NegotiationsUrl.Get()) {
		var ret string
		return ret
	}
	return *o.NegotiationsUrl.Get()
}

// GetNegotiationsUrlOk returns a tuple with the NegotiationsUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *VacanciesVacancyCommonFields) GetNegotiationsUrlOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.NegotiationsUrl.Get(), o.NegotiationsUrl.IsSet()
}

// HasNegotiationsUrl returns a boolean if a field has been set.
func (o *VacanciesVacancyCommonFields) HasNegotiationsUrl() bool {
	if o != nil && o.NegotiationsUrl.IsSet() {
		return true
	}

	return false
}

// SetNegotiationsUrl gets a reference to the given NullableString and assigns it to the NegotiationsUrl field.
func (o *VacanciesVacancyCommonFields) SetNegotiationsUrl(v string) {
	o.NegotiationsUrl.Set(&v)
}
// SetNegotiationsUrlNil sets the value for NegotiationsUrl to be an explicit nil
func (o *VacanciesVacancyCommonFields) SetNegotiationsUrlNil() {
	o.NegotiationsUrl.Set(nil)
}

// UnsetNegotiationsUrl ensures that no value is present for NegotiationsUrl, not even an explicit nil
func (o *VacanciesVacancyCommonFields) UnsetNegotiationsUrl() {
	o.NegotiationsUrl.Unset()
}

// GetPremium returns the Premium field value
func (o *VacanciesVacancyCommonFields) GetPremium() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.Premium
}

// GetPremiumOk returns a tuple with the Premium field value
// and a boolean to check if the value has been set.
func (o *VacanciesVacancyCommonFields) GetPremiumOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Premium, true
}

// SetPremium sets field value
func (o *VacanciesVacancyCommonFields) SetPremium(v bool) {
	o.Premium = v
}

// GetProfessionalRoles returns the ProfessionalRoles field value
func (o *VacanciesVacancyCommonFields) GetProfessionalRoles() []VacancyProfessionalRoleItemOutput {
	if o == nil {
		var ret []VacancyProfessionalRoleItemOutput
		return ret
	}

	return o.ProfessionalRoles
}

// GetProfessionalRolesOk returns a tuple with the ProfessionalRoles field value
// and a boolean to check if the value has been set.
func (o *VacanciesVacancyCommonFields) GetProfessionalRolesOk() ([]VacancyProfessionalRoleItemOutput, bool) {
	if o == nil {
		return nil, false
	}
	return o.ProfessionalRoles, true
}

// SetProfessionalRoles sets field value
func (o *VacanciesVacancyCommonFields) SetProfessionalRoles(v []VacancyProfessionalRoleItemOutput) {
	o.ProfessionalRoles = v
}

// GetPublishedAt returns the PublishedAt field value
func (o *VacanciesVacancyCommonFields) GetPublishedAt() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.PublishedAt
}

// GetPublishedAtOk returns a tuple with the PublishedAt field value
// and a boolean to check if the value has been set.
func (o *VacanciesVacancyCommonFields) GetPublishedAtOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.PublishedAt, true
}

// SetPublishedAt sets field value
func (o *VacanciesVacancyCommonFields) SetPublishedAt(v string) {
	o.PublishedAt = v
}

// GetRelations returns the Relations field value if set, zero value otherwise.
func (o *VacanciesVacancyCommonFields) GetRelations() []VacancyRelationItem {
	if o == nil || IsNil(o.Relations) {
		var ret []VacancyRelationItem
		return ret
	}
	return o.Relations
}

// GetRelationsOk returns a tuple with the Relations field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VacanciesVacancyCommonFields) GetRelationsOk() ([]VacancyRelationItem, bool) {
	if o == nil || IsNil(o.Relations) {
		return nil, false
	}
	return o.Relations, true
}

// HasRelations returns a boolean if a field has been set.
func (o *VacanciesVacancyCommonFields) HasRelations() bool {
	if o != nil && !IsNil(o.Relations) {
		return true
	}

	return false
}

// SetRelations gets a reference to the given []VacancyRelationItem and assigns it to the Relations field.
func (o *VacanciesVacancyCommonFields) SetRelations(v []VacancyRelationItem) {
	o.Relations = v
}

// GetResponseLetterRequired returns the ResponseLetterRequired field value
func (o *VacanciesVacancyCommonFields) GetResponseLetterRequired() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.ResponseLetterRequired
}

// GetResponseLetterRequiredOk returns a tuple with the ResponseLetterRequired field value
// and a boolean to check if the value has been set.
func (o *VacanciesVacancyCommonFields) GetResponseLetterRequiredOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ResponseLetterRequired, true
}

// SetResponseLetterRequired sets field value
func (o *VacanciesVacancyCommonFields) SetResponseLetterRequired(v bool) {
	o.ResponseLetterRequired = v
}

// GetResponseUrl returns the ResponseUrl field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *VacanciesVacancyCommonFields) GetResponseUrl() string {
	if o == nil || IsNil(o.ResponseUrl.Get()) {
		var ret string
		return ret
	}
	return *o.ResponseUrl.Get()
}

// GetResponseUrlOk returns a tuple with the ResponseUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *VacanciesVacancyCommonFields) GetResponseUrlOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.ResponseUrl.Get(), o.ResponseUrl.IsSet()
}

// HasResponseUrl returns a boolean if a field has been set.
func (o *VacanciesVacancyCommonFields) HasResponseUrl() bool {
	if o != nil && o.ResponseUrl.IsSet() {
		return true
	}

	return false
}

// SetResponseUrl gets a reference to the given NullableString and assigns it to the ResponseUrl field.
func (o *VacanciesVacancyCommonFields) SetResponseUrl(v string) {
	o.ResponseUrl.Set(&v)
}
// SetResponseUrlNil sets the value for ResponseUrl to be an explicit nil
func (o *VacanciesVacancyCommonFields) SetResponseUrlNil() {
	o.ResponseUrl.Set(nil)
}

// UnsetResponseUrl ensures that no value is present for ResponseUrl, not even an explicit nil
func (o *VacanciesVacancyCommonFields) UnsetResponseUrl() {
	o.ResponseUrl.Unset()
}

// GetSalary returns the Salary field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *VacanciesVacancyCommonFields) GetSalary() VacancySalary {
	if o == nil || IsNil(o.Salary.Get()) {
		var ret VacancySalary
		return ret
	}
	return *o.Salary.Get()
}

// GetSalaryOk returns a tuple with the Salary field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *VacanciesVacancyCommonFields) GetSalaryOk() (*VacancySalary, bool) {
	if o == nil {
		return nil, false
	}
	return o.Salary.Get(), o.Salary.IsSet()
}

// HasSalary returns a boolean if a field has been set.
func (o *VacanciesVacancyCommonFields) HasSalary() bool {
	if o != nil && o.Salary.IsSet() {
		return true
	}

	return false
}

// SetSalary gets a reference to the given NullableVacancySalary and assigns it to the Salary field.
func (o *VacanciesVacancyCommonFields) SetSalary(v VacancySalary) {
	o.Salary.Set(&v)
}
// SetSalaryNil sets the value for Salary to be an explicit nil
func (o *VacanciesVacancyCommonFields) SetSalaryNil() {
	o.Salary.Set(nil)
}

// UnsetSalary ensures that no value is present for Salary, not even an explicit nil
func (o *VacanciesVacancyCommonFields) UnsetSalary() {
	o.Salary.Unset()
}

// GetSchedule returns the Schedule field value
func (o *VacanciesVacancyCommonFields) GetSchedule() VacancyScheduleOutput {
	if o == nil {
		var ret VacancyScheduleOutput
		return ret
	}

	return o.Schedule
}

// GetScheduleOk returns a tuple with the Schedule field value
// and a boolean to check if the value has been set.
func (o *VacanciesVacancyCommonFields) GetScheduleOk() (*VacancyScheduleOutput, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Schedule, true
}

// SetSchedule sets field value
func (o *VacanciesVacancyCommonFields) SetSchedule(v VacancyScheduleOutput) {
	o.Schedule = v
}

// GetSuitableResumesUrl returns the SuitableResumesUrl field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *VacanciesVacancyCommonFields) GetSuitableResumesUrl() string {
	if o == nil || IsNil(o.SuitableResumesUrl.Get()) {
		var ret string
		return ret
	}
	return *o.SuitableResumesUrl.Get()
}

// GetSuitableResumesUrlOk returns a tuple with the SuitableResumesUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *VacanciesVacancyCommonFields) GetSuitableResumesUrlOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.SuitableResumesUrl.Get(), o.SuitableResumesUrl.IsSet()
}

// HasSuitableResumesUrl returns a boolean if a field has been set.
func (o *VacanciesVacancyCommonFields) HasSuitableResumesUrl() bool {
	if o != nil && o.SuitableResumesUrl.IsSet() {
		return true
	}

	return false
}

// SetSuitableResumesUrl gets a reference to the given NullableString and assigns it to the SuitableResumesUrl field.
func (o *VacanciesVacancyCommonFields) SetSuitableResumesUrl(v string) {
	o.SuitableResumesUrl.Set(&v)
}
// SetSuitableResumesUrlNil sets the value for SuitableResumesUrl to be an explicit nil
func (o *VacanciesVacancyCommonFields) SetSuitableResumesUrlNil() {
	o.SuitableResumesUrl.Set(nil)
}

// UnsetSuitableResumesUrl ensures that no value is present for SuitableResumesUrl, not even an explicit nil
func (o *VacanciesVacancyCommonFields) UnsetSuitableResumesUrl() {
	o.SuitableResumesUrl.Unset()
}

// GetTest returns the Test field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *VacanciesVacancyCommonFields) GetTest() VacancyDraftTest {
	if o == nil || IsNil(o.Test.Get()) {
		var ret VacancyDraftTest
		return ret
	}
	return *o.Test.Get()
}

// GetTestOk returns a tuple with the Test field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *VacanciesVacancyCommonFields) GetTestOk() (*VacancyDraftTest, bool) {
	if o == nil {
		return nil, false
	}
	return o.Test.Get(), o.Test.IsSet()
}

// HasTest returns a boolean if a field has been set.
func (o *VacanciesVacancyCommonFields) HasTest() bool {
	if o != nil && o.Test.IsSet() {
		return true
	}

	return false
}

// SetTest gets a reference to the given NullableVacancyDraftTest and assigns it to the Test field.
func (o *VacanciesVacancyCommonFields) SetTest(v VacancyDraftTest) {
	o.Test.Set(&v)
}
// SetTestNil sets the value for Test to be an explicit nil
func (o *VacanciesVacancyCommonFields) SetTestNil() {
	o.Test.Set(nil)
}

// UnsetTest ensures that no value is present for Test, not even an explicit nil
func (o *VacanciesVacancyCommonFields) UnsetTest() {
	o.Test.Unset()
}

// GetType returns the Type field value
func (o *VacanciesVacancyCommonFields) GetType() IncludesIdName {
	if o == nil {
		var ret IncludesIdName
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *VacanciesVacancyCommonFields) GetTypeOk() (*IncludesIdName, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *VacanciesVacancyCommonFields) SetType(v IncludesIdName) {
	o.Type = v
}

// GetVacancyConstructorTemplate returns the VacancyConstructorTemplate field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *VacanciesVacancyCommonFields) GetVacancyConstructorTemplate() VacancyVacancyConstructorTemplate {
	if o == nil || IsNil(o.VacancyConstructorTemplate.Get()) {
		var ret VacancyVacancyConstructorTemplate
		return ret
	}
	return *o.VacancyConstructorTemplate.Get()
}

// GetVacancyConstructorTemplateOk returns a tuple with the VacancyConstructorTemplate field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *VacanciesVacancyCommonFields) GetVacancyConstructorTemplateOk() (*VacancyVacancyConstructorTemplate, bool) {
	if o == nil {
		return nil, false
	}
	return o.VacancyConstructorTemplate.Get(), o.VacancyConstructorTemplate.IsSet()
}

// HasVacancyConstructorTemplate returns a boolean if a field has been set.
func (o *VacanciesVacancyCommonFields) HasVacancyConstructorTemplate() bool {
	if o != nil && o.VacancyConstructorTemplate.IsSet() {
		return true
	}

	return false
}

// SetVacancyConstructorTemplate gets a reference to the given NullableVacancyVacancyConstructorTemplate and assigns it to the VacancyConstructorTemplate field.
func (o *VacanciesVacancyCommonFields) SetVacancyConstructorTemplate(v VacancyVacancyConstructorTemplate) {
	o.VacancyConstructorTemplate.Set(&v)
}
// SetVacancyConstructorTemplateNil sets the value for VacancyConstructorTemplate to be an explicit nil
func (o *VacanciesVacancyCommonFields) SetVacancyConstructorTemplateNil() {
	o.VacancyConstructorTemplate.Set(nil)
}

// UnsetVacancyConstructorTemplate ensures that no value is present for VacancyConstructorTemplate, not even an explicit nil
func (o *VacanciesVacancyCommonFields) UnsetVacancyConstructorTemplate() {
	o.VacancyConstructorTemplate.Unset()
}

// GetVideoVacancy returns the VideoVacancy field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *VacanciesVacancyCommonFields) GetVideoVacancy() VacancyVideoVacancy {
	if o == nil || IsNil(o.VideoVacancy.Get()) {
		var ret VacancyVideoVacancy
		return ret
	}
	return *o.VideoVacancy.Get()
}

// GetVideoVacancyOk returns a tuple with the VideoVacancy field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *VacanciesVacancyCommonFields) GetVideoVacancyOk() (*VacancyVideoVacancy, bool) {
	if o == nil {
		return nil, false
	}
	return o.VideoVacancy.Get(), o.VideoVacancy.IsSet()
}

// HasVideoVacancy returns a boolean if a field has been set.
func (o *VacanciesVacancyCommonFields) HasVideoVacancy() bool {
	if o != nil && o.VideoVacancy.IsSet() {
		return true
	}

	return false
}

// SetVideoVacancy gets a reference to the given NullableVacancyVideoVacancy and assigns it to the VideoVacancy field.
func (o *VacanciesVacancyCommonFields) SetVideoVacancy(v VacancyVideoVacancy) {
	o.VideoVacancy.Set(&v)
}
// SetVideoVacancyNil sets the value for VideoVacancy to be an explicit nil
func (o *VacanciesVacancyCommonFields) SetVideoVacancyNil() {
	o.VideoVacancy.Set(nil)
}

// UnsetVideoVacancy ensures that no value is present for VideoVacancy, not even an explicit nil
func (o *VacanciesVacancyCommonFields) UnsetVideoVacancy() {
	o.VideoVacancy.Unset()
}

// GetWorkingDays returns the WorkingDays field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *VacanciesVacancyCommonFields) GetWorkingDays() []VacancyWorkingDayItemOutput {
	if o == nil {
		var ret []VacancyWorkingDayItemOutput
		return ret
	}
	return o.WorkingDays
}

// GetWorkingDaysOk returns a tuple with the WorkingDays field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *VacanciesVacancyCommonFields) GetWorkingDaysOk() ([]VacancyWorkingDayItemOutput, bool) {
	if o == nil || IsNil(o.WorkingDays) {
		return nil, false
	}
	return o.WorkingDays, true
}

// HasWorkingDays returns a boolean if a field has been set.
func (o *VacanciesVacancyCommonFields) HasWorkingDays() bool {
	if o != nil && !IsNil(o.WorkingDays) {
		return true
	}

	return false
}

// SetWorkingDays gets a reference to the given []VacancyWorkingDayItemOutput and assigns it to the WorkingDays field.
func (o *VacanciesVacancyCommonFields) SetWorkingDays(v []VacancyWorkingDayItemOutput) {
	o.WorkingDays = v
}

// GetWorkingTimeIntervals returns the WorkingTimeIntervals field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *VacanciesVacancyCommonFields) GetWorkingTimeIntervals() []VacancyWorkingTimeIntervalItemOutput {
	if o == nil {
		var ret []VacancyWorkingTimeIntervalItemOutput
		return ret
	}
	return o.WorkingTimeIntervals
}

// GetWorkingTimeIntervalsOk returns a tuple with the WorkingTimeIntervals field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *VacanciesVacancyCommonFields) GetWorkingTimeIntervalsOk() ([]VacancyWorkingTimeIntervalItemOutput, bool) {
	if o == nil || IsNil(o.WorkingTimeIntervals) {
		return nil, false
	}
	return o.WorkingTimeIntervals, true
}

// HasWorkingTimeIntervals returns a boolean if a field has been set.
func (o *VacanciesVacancyCommonFields) HasWorkingTimeIntervals() bool {
	if o != nil && !IsNil(o.WorkingTimeIntervals) {
		return true
	}

	return false
}

// SetWorkingTimeIntervals gets a reference to the given []VacancyWorkingTimeIntervalItemOutput and assigns it to the WorkingTimeIntervals field.
func (o *VacanciesVacancyCommonFields) SetWorkingTimeIntervals(v []VacancyWorkingTimeIntervalItemOutput) {
	o.WorkingTimeIntervals = v
}

// GetWorkingTimeModes returns the WorkingTimeModes field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *VacanciesVacancyCommonFields) GetWorkingTimeModes() []VacancyWorkingTimeModeItemOutput {
	if o == nil {
		var ret []VacancyWorkingTimeModeItemOutput
		return ret
	}
	return o.WorkingTimeModes
}

// GetWorkingTimeModesOk returns a tuple with the WorkingTimeModes field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *VacanciesVacancyCommonFields) GetWorkingTimeModesOk() ([]VacancyWorkingTimeModeItemOutput, bool) {
	if o == nil || IsNil(o.WorkingTimeModes) {
		return nil, false
	}
	return o.WorkingTimeModes, true
}

// HasWorkingTimeModes returns a boolean if a field has been set.
func (o *VacanciesVacancyCommonFields) HasWorkingTimeModes() bool {
	if o != nil && !IsNil(o.WorkingTimeModes) {
		return true
	}

	return false
}

// SetWorkingTimeModes gets a reference to the given []VacancyWorkingTimeModeItemOutput and assigns it to the WorkingTimeModes field.
func (o *VacanciesVacancyCommonFields) SetWorkingTimeModes(v []VacancyWorkingTimeModeItemOutput) {
	o.WorkingTimeModes = v
}

func (o VacanciesVacancyCommonFields) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o VacanciesVacancyCommonFields) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["accept_handicapped"] = o.AcceptHandicapped
	toSerialize["accept_incomplete_resumes"] = o.AcceptIncompleteResumes
	toSerialize["accept_kids"] = o.AcceptKids
	if o.AcceptTemporary.IsSet() {
		toSerialize["accept_temporary"] = o.AcceptTemporary.Get()
	}
	toSerialize["allow_messages"] = o.AllowMessages
	toSerialize["alternate_url"] = o.AlternateUrl
	toSerialize["apply_alternate_url"] = o.ApplyAlternateUrl
	toSerialize["approved"] = o.Approved
	toSerialize["archived"] = o.Archived
	toSerialize["area"] = o.Area
	toSerialize["billing_type"] = o.BillingType.Get()
	if !IsNil(o.BrandedDescription) {
		toSerialize["branded_description"] = o.BrandedDescription
	}
	if o.Code.IsSet() {
		toSerialize["code"] = o.Code.Get()
	}
	if o.Contacts.IsSet() {
		toSerialize["contacts"] = o.Contacts.Get()
	}
	toSerialize["created_at"] = o.CreatedAt
	if o.Department.IsSet() {
		toSerialize["department"] = o.Department.Get()
	}
	toSerialize["description"] = o.Description
	toSerialize["driver_license_types"] = o.DriverLicenseTypes
	if o.Employer.IsSet() {
		toSerialize["employer"] = o.Employer.Get()
	}
	if o.Employment.IsSet() {
		toSerialize["employment"] = o.Employment.Get()
	}
	toSerialize["experience"] = o.Experience.Get()
	toSerialize["has_test"] = o.HasTest
	toSerialize["id"] = o.Id
	toSerialize["initial_created_at"] = o.InitialCreatedAt
	if o.InsiderInterview.IsSet() {
		toSerialize["insider_interview"] = o.InsiderInterview.Get()
	}
	toSerialize["key_skills"] = o.KeySkills
	if o.Languages != nil {
		toSerialize["languages"] = o.Languages
	}
	toSerialize["name"] = o.Name
	if o.NegotiationsUrl.IsSet() {
		toSerialize["negotiations_url"] = o.NegotiationsUrl.Get()
	}
	toSerialize["premium"] = o.Premium
	toSerialize["professional_roles"] = o.ProfessionalRoles
	toSerialize["published_at"] = o.PublishedAt
	if !IsNil(o.Relations) {
		toSerialize["relations"] = o.Relations
	}
	toSerialize["response_letter_required"] = o.ResponseLetterRequired
	if o.ResponseUrl.IsSet() {
		toSerialize["response_url"] = o.ResponseUrl.Get()
	}
	if o.Salary.IsSet() {
		toSerialize["salary"] = o.Salary.Get()
	}
	toSerialize["schedule"] = o.Schedule
	if o.SuitableResumesUrl.IsSet() {
		toSerialize["suitable_resumes_url"] = o.SuitableResumesUrl.Get()
	}
	if o.Test.IsSet() {
		toSerialize["test"] = o.Test.Get()
	}
	toSerialize["type"] = o.Type
	if o.VacancyConstructorTemplate.IsSet() {
		toSerialize["vacancy_constructor_template"] = o.VacancyConstructorTemplate.Get()
	}
	if o.VideoVacancy.IsSet() {
		toSerialize["video_vacancy"] = o.VideoVacancy.Get()
	}
	if o.WorkingDays != nil {
		toSerialize["working_days"] = o.WorkingDays
	}
	if o.WorkingTimeIntervals != nil {
		toSerialize["working_time_intervals"] = o.WorkingTimeIntervals
	}
	if o.WorkingTimeModes != nil {
		toSerialize["working_time_modes"] = o.WorkingTimeModes
	}
	return toSerialize, nil
}

func (o *VacanciesVacancyCommonFields) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"accept_handicapped",
		"accept_incomplete_resumes",
		"accept_kids",
		"allow_messages",
		"alternate_url",
		"apply_alternate_url",
		"approved",
		"archived",
		"area",
		"billing_type",
		"created_at",
		"description",
		"driver_license_types",
		"experience",
		"has_test",
		"id",
		"initial_created_at",
		"key_skills",
		"name",
		"premium",
		"professional_roles",
		"published_at",
		"response_letter_required",
		"schedule",
		"type",
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

	varVacanciesVacancyCommonFields := _VacanciesVacancyCommonFields{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varVacanciesVacancyCommonFields)

	if err != nil {
		return err
	}

	*o = VacanciesVacancyCommonFields(varVacanciesVacancyCommonFields)

	return err
}

type NullableVacanciesVacancyCommonFields struct {
	value *VacanciesVacancyCommonFields
	isSet bool
}

func (v NullableVacanciesVacancyCommonFields) Get() *VacanciesVacancyCommonFields {
	return v.value
}

func (v *NullableVacanciesVacancyCommonFields) Set(val *VacanciesVacancyCommonFields) {
	v.value = val
	v.isSet = true
}

func (v NullableVacanciesVacancyCommonFields) IsSet() bool {
	return v.isSet
}

func (v *NullableVacanciesVacancyCommonFields) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableVacanciesVacancyCommonFields(val *VacanciesVacancyCommonFields) *NullableVacanciesVacancyCommonFields {
	return &NullableVacanciesVacancyCommonFields{value: val, isSet: true}
}

func (v NullableVacanciesVacancyCommonFields) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableVacanciesVacancyCommonFields) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


