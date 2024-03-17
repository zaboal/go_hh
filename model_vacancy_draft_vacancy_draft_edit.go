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
)

// checks if the VacancyDraftVacancyDraftEdit type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &VacancyDraftVacancyDraftEdit{}

// VacancyDraftVacancyDraftEdit struct for VacancyDraftVacancyDraftEdit
type VacancyDraftVacancyDraftEdit struct {
	// Указание, что вакансия доступна для соискателей с инвалидностью
	AcceptHandicapped *bool `json:"accept_handicapped,omitempty"`
	// Разрешен ли отклик на вакансию неполным резюме
	AcceptIncompleteResumes *bool `json:"accept_incomplete_resumes,omitempty"`
	// Указание, что вакансия доступна для соискателей старше 14 лет [подробнее](https://github.com/hhru/api/blob/master/docs/employer_vacancies_accept_kids.md#accept-kids)
	AcceptKids *bool `json:"accept_kids,omitempty"`
	// Указание, что вакансия доступна с временным трудоустройством
	AcceptTemporary NullableBool `json:"accept_temporary,omitempty"`
	Address NullableVacancyDraftAddress `json:"address,omitempty"`
	// Возможность [переписки с кандидатами](https://inboxemp.hh.ru/) по данной вакансии
	AllowMessages *bool `json:"allow_messages,omitempty"`
	Area *VacancyArea `json:"area,omitempty"`
	// Список регионов
	Areas []VacancyArea `json:"areas,omitempty"`
	// Идентификатор рабочего аккаунта менеджера, которому перейдет вакансия после публикации
	AssignedManagerId NullableString `json:"assigned_manager_id,omitempty"`
	BillingType NullableVacancyDraftBillingType `json:"billing_type,omitempty"`
	BrandedTemplate *VacancyDraftBrandedTemplate `json:"branded_template,omitempty"`
	// Внутренний код вакансии
	Code NullableString `json:"code,omitempty"`
	Contacts NullableVacancyDraftContacts `json:"contacts,omitempty"`
	// Название компании для анонимных вакансий (`type.id=anonymous`), например \"крупный российский банк\". Поле становится обязательным при публикации вакансии **анонимного** типа
	CustomEmployerName NullableString `json:"custom_employer_name,omitempty"`
	Department NullableVacancyDepartment `json:"department,omitempty"`
	// Описание в html, минимум 1 символ, для успешной публикации вакансии не менее 200 символов
	Description NullableString `json:"description,omitempty"`
	// Список требуемых категорий водительских прав
	DriverLicenseTypes []VacancyDriverLicenceTypeItem `json:"driver_license_types,omitempty"`
	Employment NullableVacancyDraftEmployment `json:"employment,omitempty"`
	Experience NullableVacancyExperience `json:"experience,omitempty"`
	// Список ключевых навыков, не более 30
	KeySkills []VacancyDraftKeySkillItem `json:"key_skills,omitempty"`
	// Список языков вакансии
	Languages []VacancyLanguage `json:"languages,omitempty"`
	// Название
	Name NullableString `json:"name,omitempty"`
	// Список профессиональных ролей
	ProfessionalRoles []VacancyDraftProfessionalRoleItem `json:"professional_roles,omitempty"`
	// Обязательно ли заполнять сообщение при отклике на вакансию
	ResponseLetterRequired *bool `json:"response_letter_required,omitempty"`
	// Уведомлять ли менеджера о новых откликах
	ResponseNotifications *bool `json:"response_notifications,omitempty"`
	// URL отклика для прямых вакансий (`type.id=direct`)
	ResponseUrl NullableString `json:"response_url,omitempty"`
	Salary NullableVacancySalary `json:"salary,omitempty"`
	Schedule NullableVacancySchedule `json:"schedule,omitempty"`
	Test NullableVacancyDraftTest `json:"test,omitempty"`
	Type NullableVacancyDraftType `json:"type,omitempty"`
	// Вашу вакансию увидят больше людей. Мы разместим ее дополнительно на сервисе Зарплата.ру
	WithZp *bool `json:"with_zp,omitempty"`
	// Список рабочих дней
	WorkingDays []VacancyWorkingDayItem `json:"working_days,omitempty"`
	// Список с временными интервалами работы
	WorkingTimeIntervals []VacancyWorkingTimeIntervalItem `json:"working_time_intervals,omitempty"`
	// Список режимов времени работы
	WorkingTimeModes []VacancyWorkingTimeModeItem `json:"working_time_modes,omitempty"`
}

// NewVacancyDraftVacancyDraftEdit instantiates a new VacancyDraftVacancyDraftEdit object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewVacancyDraftVacancyDraftEdit() *VacancyDraftVacancyDraftEdit {
	this := VacancyDraftVacancyDraftEdit{}
	return &this
}

// NewVacancyDraftVacancyDraftEditWithDefaults instantiates a new VacancyDraftVacancyDraftEdit object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewVacancyDraftVacancyDraftEditWithDefaults() *VacancyDraftVacancyDraftEdit {
	this := VacancyDraftVacancyDraftEdit{}
	return &this
}

// GetAcceptHandicapped returns the AcceptHandicapped field value if set, zero value otherwise.
func (o *VacancyDraftVacancyDraftEdit) GetAcceptHandicapped() bool {
	if o == nil || IsNil(o.AcceptHandicapped) {
		var ret bool
		return ret
	}
	return *o.AcceptHandicapped
}

// GetAcceptHandicappedOk returns a tuple with the AcceptHandicapped field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VacancyDraftVacancyDraftEdit) GetAcceptHandicappedOk() (*bool, bool) {
	if o == nil || IsNil(o.AcceptHandicapped) {
		return nil, false
	}
	return o.AcceptHandicapped, true
}

// HasAcceptHandicapped returns a boolean if a field has been set.
func (o *VacancyDraftVacancyDraftEdit) HasAcceptHandicapped() bool {
	if o != nil && !IsNil(o.AcceptHandicapped) {
		return true
	}

	return false
}

// SetAcceptHandicapped gets a reference to the given bool and assigns it to the AcceptHandicapped field.
func (o *VacancyDraftVacancyDraftEdit) SetAcceptHandicapped(v bool) {
	o.AcceptHandicapped = &v
}

// GetAcceptIncompleteResumes returns the AcceptIncompleteResumes field value if set, zero value otherwise.
func (o *VacancyDraftVacancyDraftEdit) GetAcceptIncompleteResumes() bool {
	if o == nil || IsNil(o.AcceptIncompleteResumes) {
		var ret bool
		return ret
	}
	return *o.AcceptIncompleteResumes
}

// GetAcceptIncompleteResumesOk returns a tuple with the AcceptIncompleteResumes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VacancyDraftVacancyDraftEdit) GetAcceptIncompleteResumesOk() (*bool, bool) {
	if o == nil || IsNil(o.AcceptIncompleteResumes) {
		return nil, false
	}
	return o.AcceptIncompleteResumes, true
}

// HasAcceptIncompleteResumes returns a boolean if a field has been set.
func (o *VacancyDraftVacancyDraftEdit) HasAcceptIncompleteResumes() bool {
	if o != nil && !IsNil(o.AcceptIncompleteResumes) {
		return true
	}

	return false
}

// SetAcceptIncompleteResumes gets a reference to the given bool and assigns it to the AcceptIncompleteResumes field.
func (o *VacancyDraftVacancyDraftEdit) SetAcceptIncompleteResumes(v bool) {
	o.AcceptIncompleteResumes = &v
}

// GetAcceptKids returns the AcceptKids field value if set, zero value otherwise.
func (o *VacancyDraftVacancyDraftEdit) GetAcceptKids() bool {
	if o == nil || IsNil(o.AcceptKids) {
		var ret bool
		return ret
	}
	return *o.AcceptKids
}

// GetAcceptKidsOk returns a tuple with the AcceptKids field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VacancyDraftVacancyDraftEdit) GetAcceptKidsOk() (*bool, bool) {
	if o == nil || IsNil(o.AcceptKids) {
		return nil, false
	}
	return o.AcceptKids, true
}

// HasAcceptKids returns a boolean if a field has been set.
func (o *VacancyDraftVacancyDraftEdit) HasAcceptKids() bool {
	if o != nil && !IsNil(o.AcceptKids) {
		return true
	}

	return false
}

// SetAcceptKids gets a reference to the given bool and assigns it to the AcceptKids field.
func (o *VacancyDraftVacancyDraftEdit) SetAcceptKids(v bool) {
	o.AcceptKids = &v
}

// GetAcceptTemporary returns the AcceptTemporary field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *VacancyDraftVacancyDraftEdit) GetAcceptTemporary() bool {
	if o == nil || IsNil(o.AcceptTemporary.Get()) {
		var ret bool
		return ret
	}
	return *o.AcceptTemporary.Get()
}

// GetAcceptTemporaryOk returns a tuple with the AcceptTemporary field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *VacancyDraftVacancyDraftEdit) GetAcceptTemporaryOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return o.AcceptTemporary.Get(), o.AcceptTemporary.IsSet()
}

// HasAcceptTemporary returns a boolean if a field has been set.
func (o *VacancyDraftVacancyDraftEdit) HasAcceptTemporary() bool {
	if o != nil && o.AcceptTemporary.IsSet() {
		return true
	}

	return false
}

// SetAcceptTemporary gets a reference to the given NullableBool and assigns it to the AcceptTemporary field.
func (o *VacancyDraftVacancyDraftEdit) SetAcceptTemporary(v bool) {
	o.AcceptTemporary.Set(&v)
}
// SetAcceptTemporaryNil sets the value for AcceptTemporary to be an explicit nil
func (o *VacancyDraftVacancyDraftEdit) SetAcceptTemporaryNil() {
	o.AcceptTemporary.Set(nil)
}

// UnsetAcceptTemporary ensures that no value is present for AcceptTemporary, not even an explicit nil
func (o *VacancyDraftVacancyDraftEdit) UnsetAcceptTemporary() {
	o.AcceptTemporary.Unset()
}

// GetAddress returns the Address field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *VacancyDraftVacancyDraftEdit) GetAddress() VacancyDraftAddress {
	if o == nil || IsNil(o.Address.Get()) {
		var ret VacancyDraftAddress
		return ret
	}
	return *o.Address.Get()
}

// GetAddressOk returns a tuple with the Address field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *VacancyDraftVacancyDraftEdit) GetAddressOk() (*VacancyDraftAddress, bool) {
	if o == nil {
		return nil, false
	}
	return o.Address.Get(), o.Address.IsSet()
}

// HasAddress returns a boolean if a field has been set.
func (o *VacancyDraftVacancyDraftEdit) HasAddress() bool {
	if o != nil && o.Address.IsSet() {
		return true
	}

	return false
}

// SetAddress gets a reference to the given NullableVacancyDraftAddress and assigns it to the Address field.
func (o *VacancyDraftVacancyDraftEdit) SetAddress(v VacancyDraftAddress) {
	o.Address.Set(&v)
}
// SetAddressNil sets the value for Address to be an explicit nil
func (o *VacancyDraftVacancyDraftEdit) SetAddressNil() {
	o.Address.Set(nil)
}

// UnsetAddress ensures that no value is present for Address, not even an explicit nil
func (o *VacancyDraftVacancyDraftEdit) UnsetAddress() {
	o.Address.Unset()
}

// GetAllowMessages returns the AllowMessages field value if set, zero value otherwise.
func (o *VacancyDraftVacancyDraftEdit) GetAllowMessages() bool {
	if o == nil || IsNil(o.AllowMessages) {
		var ret bool
		return ret
	}
	return *o.AllowMessages
}

// GetAllowMessagesOk returns a tuple with the AllowMessages field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VacancyDraftVacancyDraftEdit) GetAllowMessagesOk() (*bool, bool) {
	if o == nil || IsNil(o.AllowMessages) {
		return nil, false
	}
	return o.AllowMessages, true
}

// HasAllowMessages returns a boolean if a field has been set.
func (o *VacancyDraftVacancyDraftEdit) HasAllowMessages() bool {
	if o != nil && !IsNil(o.AllowMessages) {
		return true
	}

	return false
}

// SetAllowMessages gets a reference to the given bool and assigns it to the AllowMessages field.
func (o *VacancyDraftVacancyDraftEdit) SetAllowMessages(v bool) {
	o.AllowMessages = &v
}

// GetArea returns the Area field value if set, zero value otherwise.
func (o *VacancyDraftVacancyDraftEdit) GetArea() VacancyArea {
	if o == nil || IsNil(o.Area) {
		var ret VacancyArea
		return ret
	}
	return *o.Area
}

// GetAreaOk returns a tuple with the Area field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VacancyDraftVacancyDraftEdit) GetAreaOk() (*VacancyArea, bool) {
	if o == nil || IsNil(o.Area) {
		return nil, false
	}
	return o.Area, true
}

// HasArea returns a boolean if a field has been set.
func (o *VacancyDraftVacancyDraftEdit) HasArea() bool {
	if o != nil && !IsNil(o.Area) {
		return true
	}

	return false
}

// SetArea gets a reference to the given VacancyArea and assigns it to the Area field.
func (o *VacancyDraftVacancyDraftEdit) SetArea(v VacancyArea) {
	o.Area = &v
}

// GetAreas returns the Areas field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *VacancyDraftVacancyDraftEdit) GetAreas() []VacancyArea {
	if o == nil {
		var ret []VacancyArea
		return ret
	}
	return o.Areas
}

// GetAreasOk returns a tuple with the Areas field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *VacancyDraftVacancyDraftEdit) GetAreasOk() ([]VacancyArea, bool) {
	if o == nil || IsNil(o.Areas) {
		return nil, false
	}
	return o.Areas, true
}

// HasAreas returns a boolean if a field has been set.
func (o *VacancyDraftVacancyDraftEdit) HasAreas() bool {
	if o != nil && !IsNil(o.Areas) {
		return true
	}

	return false
}

// SetAreas gets a reference to the given []VacancyArea and assigns it to the Areas field.
func (o *VacancyDraftVacancyDraftEdit) SetAreas(v []VacancyArea) {
	o.Areas = v
}

// GetAssignedManagerId returns the AssignedManagerId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *VacancyDraftVacancyDraftEdit) GetAssignedManagerId() string {
	if o == nil || IsNil(o.AssignedManagerId.Get()) {
		var ret string
		return ret
	}
	return *o.AssignedManagerId.Get()
}

// GetAssignedManagerIdOk returns a tuple with the AssignedManagerId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *VacancyDraftVacancyDraftEdit) GetAssignedManagerIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.AssignedManagerId.Get(), o.AssignedManagerId.IsSet()
}

// HasAssignedManagerId returns a boolean if a field has been set.
func (o *VacancyDraftVacancyDraftEdit) HasAssignedManagerId() bool {
	if o != nil && o.AssignedManagerId.IsSet() {
		return true
	}

	return false
}

// SetAssignedManagerId gets a reference to the given NullableString and assigns it to the AssignedManagerId field.
func (o *VacancyDraftVacancyDraftEdit) SetAssignedManagerId(v string) {
	o.AssignedManagerId.Set(&v)
}
// SetAssignedManagerIdNil sets the value for AssignedManagerId to be an explicit nil
func (o *VacancyDraftVacancyDraftEdit) SetAssignedManagerIdNil() {
	o.AssignedManagerId.Set(nil)
}

// UnsetAssignedManagerId ensures that no value is present for AssignedManagerId, not even an explicit nil
func (o *VacancyDraftVacancyDraftEdit) UnsetAssignedManagerId() {
	o.AssignedManagerId.Unset()
}

// GetBillingType returns the BillingType field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *VacancyDraftVacancyDraftEdit) GetBillingType() VacancyDraftBillingType {
	if o == nil || IsNil(o.BillingType.Get()) {
		var ret VacancyDraftBillingType
		return ret
	}
	return *o.BillingType.Get()
}

// GetBillingTypeOk returns a tuple with the BillingType field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *VacancyDraftVacancyDraftEdit) GetBillingTypeOk() (*VacancyDraftBillingType, bool) {
	if o == nil {
		return nil, false
	}
	return o.BillingType.Get(), o.BillingType.IsSet()
}

// HasBillingType returns a boolean if a field has been set.
func (o *VacancyDraftVacancyDraftEdit) HasBillingType() bool {
	if o != nil && o.BillingType.IsSet() {
		return true
	}

	return false
}

// SetBillingType gets a reference to the given NullableVacancyDraftBillingType and assigns it to the BillingType field.
func (o *VacancyDraftVacancyDraftEdit) SetBillingType(v VacancyDraftBillingType) {
	o.BillingType.Set(&v)
}
// SetBillingTypeNil sets the value for BillingType to be an explicit nil
func (o *VacancyDraftVacancyDraftEdit) SetBillingTypeNil() {
	o.BillingType.Set(nil)
}

// UnsetBillingType ensures that no value is present for BillingType, not even an explicit nil
func (o *VacancyDraftVacancyDraftEdit) UnsetBillingType() {
	o.BillingType.Unset()
}

// GetBrandedTemplate returns the BrandedTemplate field value if set, zero value otherwise.
func (o *VacancyDraftVacancyDraftEdit) GetBrandedTemplate() VacancyDraftBrandedTemplate {
	if o == nil || IsNil(o.BrandedTemplate) {
		var ret VacancyDraftBrandedTemplate
		return ret
	}
	return *o.BrandedTemplate
}

// GetBrandedTemplateOk returns a tuple with the BrandedTemplate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VacancyDraftVacancyDraftEdit) GetBrandedTemplateOk() (*VacancyDraftBrandedTemplate, bool) {
	if o == nil || IsNil(o.BrandedTemplate) {
		return nil, false
	}
	return o.BrandedTemplate, true
}

// HasBrandedTemplate returns a boolean if a field has been set.
func (o *VacancyDraftVacancyDraftEdit) HasBrandedTemplate() bool {
	if o != nil && !IsNil(o.BrandedTemplate) {
		return true
	}

	return false
}

// SetBrandedTemplate gets a reference to the given VacancyDraftBrandedTemplate and assigns it to the BrandedTemplate field.
func (o *VacancyDraftVacancyDraftEdit) SetBrandedTemplate(v VacancyDraftBrandedTemplate) {
	o.BrandedTemplate = &v
}

// GetCode returns the Code field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *VacancyDraftVacancyDraftEdit) GetCode() string {
	if o == nil || IsNil(o.Code.Get()) {
		var ret string
		return ret
	}
	return *o.Code.Get()
}

// GetCodeOk returns a tuple with the Code field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *VacancyDraftVacancyDraftEdit) GetCodeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Code.Get(), o.Code.IsSet()
}

// HasCode returns a boolean if a field has been set.
func (o *VacancyDraftVacancyDraftEdit) HasCode() bool {
	if o != nil && o.Code.IsSet() {
		return true
	}

	return false
}

// SetCode gets a reference to the given NullableString and assigns it to the Code field.
func (o *VacancyDraftVacancyDraftEdit) SetCode(v string) {
	o.Code.Set(&v)
}
// SetCodeNil sets the value for Code to be an explicit nil
func (o *VacancyDraftVacancyDraftEdit) SetCodeNil() {
	o.Code.Set(nil)
}

// UnsetCode ensures that no value is present for Code, not even an explicit nil
func (o *VacancyDraftVacancyDraftEdit) UnsetCode() {
	o.Code.Unset()
}

// GetContacts returns the Contacts field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *VacancyDraftVacancyDraftEdit) GetContacts() VacancyDraftContacts {
	if o == nil || IsNil(o.Contacts.Get()) {
		var ret VacancyDraftContacts
		return ret
	}
	return *o.Contacts.Get()
}

// GetContactsOk returns a tuple with the Contacts field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *VacancyDraftVacancyDraftEdit) GetContactsOk() (*VacancyDraftContacts, bool) {
	if o == nil {
		return nil, false
	}
	return o.Contacts.Get(), o.Contacts.IsSet()
}

// HasContacts returns a boolean if a field has been set.
func (o *VacancyDraftVacancyDraftEdit) HasContacts() bool {
	if o != nil && o.Contacts.IsSet() {
		return true
	}

	return false
}

// SetContacts gets a reference to the given NullableVacancyDraftContacts and assigns it to the Contacts field.
func (o *VacancyDraftVacancyDraftEdit) SetContacts(v VacancyDraftContacts) {
	o.Contacts.Set(&v)
}
// SetContactsNil sets the value for Contacts to be an explicit nil
func (o *VacancyDraftVacancyDraftEdit) SetContactsNil() {
	o.Contacts.Set(nil)
}

// UnsetContacts ensures that no value is present for Contacts, not even an explicit nil
func (o *VacancyDraftVacancyDraftEdit) UnsetContacts() {
	o.Contacts.Unset()
}

// GetCustomEmployerName returns the CustomEmployerName field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *VacancyDraftVacancyDraftEdit) GetCustomEmployerName() string {
	if o == nil || IsNil(o.CustomEmployerName.Get()) {
		var ret string
		return ret
	}
	return *o.CustomEmployerName.Get()
}

// GetCustomEmployerNameOk returns a tuple with the CustomEmployerName field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *VacancyDraftVacancyDraftEdit) GetCustomEmployerNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.CustomEmployerName.Get(), o.CustomEmployerName.IsSet()
}

// HasCustomEmployerName returns a boolean if a field has been set.
func (o *VacancyDraftVacancyDraftEdit) HasCustomEmployerName() bool {
	if o != nil && o.CustomEmployerName.IsSet() {
		return true
	}

	return false
}

// SetCustomEmployerName gets a reference to the given NullableString and assigns it to the CustomEmployerName field.
func (o *VacancyDraftVacancyDraftEdit) SetCustomEmployerName(v string) {
	o.CustomEmployerName.Set(&v)
}
// SetCustomEmployerNameNil sets the value for CustomEmployerName to be an explicit nil
func (o *VacancyDraftVacancyDraftEdit) SetCustomEmployerNameNil() {
	o.CustomEmployerName.Set(nil)
}

// UnsetCustomEmployerName ensures that no value is present for CustomEmployerName, not even an explicit nil
func (o *VacancyDraftVacancyDraftEdit) UnsetCustomEmployerName() {
	o.CustomEmployerName.Unset()
}

// GetDepartment returns the Department field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *VacancyDraftVacancyDraftEdit) GetDepartment() VacancyDepartment {
	if o == nil || IsNil(o.Department.Get()) {
		var ret VacancyDepartment
		return ret
	}
	return *o.Department.Get()
}

// GetDepartmentOk returns a tuple with the Department field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *VacancyDraftVacancyDraftEdit) GetDepartmentOk() (*VacancyDepartment, bool) {
	if o == nil {
		return nil, false
	}
	return o.Department.Get(), o.Department.IsSet()
}

// HasDepartment returns a boolean if a field has been set.
func (o *VacancyDraftVacancyDraftEdit) HasDepartment() bool {
	if o != nil && o.Department.IsSet() {
		return true
	}

	return false
}

// SetDepartment gets a reference to the given NullableVacancyDepartment and assigns it to the Department field.
func (o *VacancyDraftVacancyDraftEdit) SetDepartment(v VacancyDepartment) {
	o.Department.Set(&v)
}
// SetDepartmentNil sets the value for Department to be an explicit nil
func (o *VacancyDraftVacancyDraftEdit) SetDepartmentNil() {
	o.Department.Set(nil)
}

// UnsetDepartment ensures that no value is present for Department, not even an explicit nil
func (o *VacancyDraftVacancyDraftEdit) UnsetDepartment() {
	o.Department.Unset()
}

// GetDescription returns the Description field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *VacancyDraftVacancyDraftEdit) GetDescription() string {
	if o == nil || IsNil(o.Description.Get()) {
		var ret string
		return ret
	}
	return *o.Description.Get()
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *VacancyDraftVacancyDraftEdit) GetDescriptionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Description.Get(), o.Description.IsSet()
}

// HasDescription returns a boolean if a field has been set.
func (o *VacancyDraftVacancyDraftEdit) HasDescription() bool {
	if o != nil && o.Description.IsSet() {
		return true
	}

	return false
}

// SetDescription gets a reference to the given NullableString and assigns it to the Description field.
func (o *VacancyDraftVacancyDraftEdit) SetDescription(v string) {
	o.Description.Set(&v)
}
// SetDescriptionNil sets the value for Description to be an explicit nil
func (o *VacancyDraftVacancyDraftEdit) SetDescriptionNil() {
	o.Description.Set(nil)
}

// UnsetDescription ensures that no value is present for Description, not even an explicit nil
func (o *VacancyDraftVacancyDraftEdit) UnsetDescription() {
	o.Description.Unset()
}

// GetDriverLicenseTypes returns the DriverLicenseTypes field value if set, zero value otherwise.
func (o *VacancyDraftVacancyDraftEdit) GetDriverLicenseTypes() []VacancyDriverLicenceTypeItem {
	if o == nil || IsNil(o.DriverLicenseTypes) {
		var ret []VacancyDriverLicenceTypeItem
		return ret
	}
	return o.DriverLicenseTypes
}

// GetDriverLicenseTypesOk returns a tuple with the DriverLicenseTypes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VacancyDraftVacancyDraftEdit) GetDriverLicenseTypesOk() ([]VacancyDriverLicenceTypeItem, bool) {
	if o == nil || IsNil(o.DriverLicenseTypes) {
		return nil, false
	}
	return o.DriverLicenseTypes, true
}

// HasDriverLicenseTypes returns a boolean if a field has been set.
func (o *VacancyDraftVacancyDraftEdit) HasDriverLicenseTypes() bool {
	if o != nil && !IsNil(o.DriverLicenseTypes) {
		return true
	}

	return false
}

// SetDriverLicenseTypes gets a reference to the given []VacancyDriverLicenceTypeItem and assigns it to the DriverLicenseTypes field.
func (o *VacancyDraftVacancyDraftEdit) SetDriverLicenseTypes(v []VacancyDriverLicenceTypeItem) {
	o.DriverLicenseTypes = v
}

// GetEmployment returns the Employment field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *VacancyDraftVacancyDraftEdit) GetEmployment() VacancyDraftEmployment {
	if o == nil || IsNil(o.Employment.Get()) {
		var ret VacancyDraftEmployment
		return ret
	}
	return *o.Employment.Get()
}

// GetEmploymentOk returns a tuple with the Employment field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *VacancyDraftVacancyDraftEdit) GetEmploymentOk() (*VacancyDraftEmployment, bool) {
	if o == nil {
		return nil, false
	}
	return o.Employment.Get(), o.Employment.IsSet()
}

// HasEmployment returns a boolean if a field has been set.
func (o *VacancyDraftVacancyDraftEdit) HasEmployment() bool {
	if o != nil && o.Employment.IsSet() {
		return true
	}

	return false
}

// SetEmployment gets a reference to the given NullableVacancyDraftEmployment and assigns it to the Employment field.
func (o *VacancyDraftVacancyDraftEdit) SetEmployment(v VacancyDraftEmployment) {
	o.Employment.Set(&v)
}
// SetEmploymentNil sets the value for Employment to be an explicit nil
func (o *VacancyDraftVacancyDraftEdit) SetEmploymentNil() {
	o.Employment.Set(nil)
}

// UnsetEmployment ensures that no value is present for Employment, not even an explicit nil
func (o *VacancyDraftVacancyDraftEdit) UnsetEmployment() {
	o.Employment.Unset()
}

// GetExperience returns the Experience field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *VacancyDraftVacancyDraftEdit) GetExperience() VacancyExperience {
	if o == nil || IsNil(o.Experience.Get()) {
		var ret VacancyExperience
		return ret
	}
	return *o.Experience.Get()
}

// GetExperienceOk returns a tuple with the Experience field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *VacancyDraftVacancyDraftEdit) GetExperienceOk() (*VacancyExperience, bool) {
	if o == nil {
		return nil, false
	}
	return o.Experience.Get(), o.Experience.IsSet()
}

// HasExperience returns a boolean if a field has been set.
func (o *VacancyDraftVacancyDraftEdit) HasExperience() bool {
	if o != nil && o.Experience.IsSet() {
		return true
	}

	return false
}

// SetExperience gets a reference to the given NullableVacancyExperience and assigns it to the Experience field.
func (o *VacancyDraftVacancyDraftEdit) SetExperience(v VacancyExperience) {
	o.Experience.Set(&v)
}
// SetExperienceNil sets the value for Experience to be an explicit nil
func (o *VacancyDraftVacancyDraftEdit) SetExperienceNil() {
	o.Experience.Set(nil)
}

// UnsetExperience ensures that no value is present for Experience, not even an explicit nil
func (o *VacancyDraftVacancyDraftEdit) UnsetExperience() {
	o.Experience.Unset()
}

// GetKeySkills returns the KeySkills field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *VacancyDraftVacancyDraftEdit) GetKeySkills() []VacancyDraftKeySkillItem {
	if o == nil {
		var ret []VacancyDraftKeySkillItem
		return ret
	}
	return o.KeySkills
}

// GetKeySkillsOk returns a tuple with the KeySkills field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *VacancyDraftVacancyDraftEdit) GetKeySkillsOk() ([]VacancyDraftKeySkillItem, bool) {
	if o == nil || IsNil(o.KeySkills) {
		return nil, false
	}
	return o.KeySkills, true
}

// HasKeySkills returns a boolean if a field has been set.
func (o *VacancyDraftVacancyDraftEdit) HasKeySkills() bool {
	if o != nil && !IsNil(o.KeySkills) {
		return true
	}

	return false
}

// SetKeySkills gets a reference to the given []VacancyDraftKeySkillItem and assigns it to the KeySkills field.
func (o *VacancyDraftVacancyDraftEdit) SetKeySkills(v []VacancyDraftKeySkillItem) {
	o.KeySkills = v
}

// GetLanguages returns the Languages field value if set, zero value otherwise.
func (o *VacancyDraftVacancyDraftEdit) GetLanguages() []VacancyLanguage {
	if o == nil || IsNil(o.Languages) {
		var ret []VacancyLanguage
		return ret
	}
	return o.Languages
}

// GetLanguagesOk returns a tuple with the Languages field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VacancyDraftVacancyDraftEdit) GetLanguagesOk() ([]VacancyLanguage, bool) {
	if o == nil || IsNil(o.Languages) {
		return nil, false
	}
	return o.Languages, true
}

// HasLanguages returns a boolean if a field has been set.
func (o *VacancyDraftVacancyDraftEdit) HasLanguages() bool {
	if o != nil && !IsNil(o.Languages) {
		return true
	}

	return false
}

// SetLanguages gets a reference to the given []VacancyLanguage and assigns it to the Languages field.
func (o *VacancyDraftVacancyDraftEdit) SetLanguages(v []VacancyLanguage) {
	o.Languages = v
}

// GetName returns the Name field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *VacancyDraftVacancyDraftEdit) GetName() string {
	if o == nil || IsNil(o.Name.Get()) {
		var ret string
		return ret
	}
	return *o.Name.Get()
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *VacancyDraftVacancyDraftEdit) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Name.Get(), o.Name.IsSet()
}

// HasName returns a boolean if a field has been set.
func (o *VacancyDraftVacancyDraftEdit) HasName() bool {
	if o != nil && o.Name.IsSet() {
		return true
	}

	return false
}

// SetName gets a reference to the given NullableString and assigns it to the Name field.
func (o *VacancyDraftVacancyDraftEdit) SetName(v string) {
	o.Name.Set(&v)
}
// SetNameNil sets the value for Name to be an explicit nil
func (o *VacancyDraftVacancyDraftEdit) SetNameNil() {
	o.Name.Set(nil)
}

// UnsetName ensures that no value is present for Name, not even an explicit nil
func (o *VacancyDraftVacancyDraftEdit) UnsetName() {
	o.Name.Unset()
}

// GetProfessionalRoles returns the ProfessionalRoles field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *VacancyDraftVacancyDraftEdit) GetProfessionalRoles() []VacancyDraftProfessionalRoleItem {
	if o == nil {
		var ret []VacancyDraftProfessionalRoleItem
		return ret
	}
	return o.ProfessionalRoles
}

// GetProfessionalRolesOk returns a tuple with the ProfessionalRoles field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *VacancyDraftVacancyDraftEdit) GetProfessionalRolesOk() ([]VacancyDraftProfessionalRoleItem, bool) {
	if o == nil || IsNil(o.ProfessionalRoles) {
		return nil, false
	}
	return o.ProfessionalRoles, true
}

// HasProfessionalRoles returns a boolean if a field has been set.
func (o *VacancyDraftVacancyDraftEdit) HasProfessionalRoles() bool {
	if o != nil && !IsNil(o.ProfessionalRoles) {
		return true
	}

	return false
}

// SetProfessionalRoles gets a reference to the given []VacancyDraftProfessionalRoleItem and assigns it to the ProfessionalRoles field.
func (o *VacancyDraftVacancyDraftEdit) SetProfessionalRoles(v []VacancyDraftProfessionalRoleItem) {
	o.ProfessionalRoles = v
}

// GetResponseLetterRequired returns the ResponseLetterRequired field value if set, zero value otherwise.
func (o *VacancyDraftVacancyDraftEdit) GetResponseLetterRequired() bool {
	if o == nil || IsNil(o.ResponseLetterRequired) {
		var ret bool
		return ret
	}
	return *o.ResponseLetterRequired
}

// GetResponseLetterRequiredOk returns a tuple with the ResponseLetterRequired field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VacancyDraftVacancyDraftEdit) GetResponseLetterRequiredOk() (*bool, bool) {
	if o == nil || IsNil(o.ResponseLetterRequired) {
		return nil, false
	}
	return o.ResponseLetterRequired, true
}

// HasResponseLetterRequired returns a boolean if a field has been set.
func (o *VacancyDraftVacancyDraftEdit) HasResponseLetterRequired() bool {
	if o != nil && !IsNil(o.ResponseLetterRequired) {
		return true
	}

	return false
}

// SetResponseLetterRequired gets a reference to the given bool and assigns it to the ResponseLetterRequired field.
func (o *VacancyDraftVacancyDraftEdit) SetResponseLetterRequired(v bool) {
	o.ResponseLetterRequired = &v
}

// GetResponseNotifications returns the ResponseNotifications field value if set, zero value otherwise.
func (o *VacancyDraftVacancyDraftEdit) GetResponseNotifications() bool {
	if o == nil || IsNil(o.ResponseNotifications) {
		var ret bool
		return ret
	}
	return *o.ResponseNotifications
}

// GetResponseNotificationsOk returns a tuple with the ResponseNotifications field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VacancyDraftVacancyDraftEdit) GetResponseNotificationsOk() (*bool, bool) {
	if o == nil || IsNil(o.ResponseNotifications) {
		return nil, false
	}
	return o.ResponseNotifications, true
}

// HasResponseNotifications returns a boolean if a field has been set.
func (o *VacancyDraftVacancyDraftEdit) HasResponseNotifications() bool {
	if o != nil && !IsNil(o.ResponseNotifications) {
		return true
	}

	return false
}

// SetResponseNotifications gets a reference to the given bool and assigns it to the ResponseNotifications field.
func (o *VacancyDraftVacancyDraftEdit) SetResponseNotifications(v bool) {
	o.ResponseNotifications = &v
}

// GetResponseUrl returns the ResponseUrl field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *VacancyDraftVacancyDraftEdit) GetResponseUrl() string {
	if o == nil || IsNil(o.ResponseUrl.Get()) {
		var ret string
		return ret
	}
	return *o.ResponseUrl.Get()
}

// GetResponseUrlOk returns a tuple with the ResponseUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *VacancyDraftVacancyDraftEdit) GetResponseUrlOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.ResponseUrl.Get(), o.ResponseUrl.IsSet()
}

// HasResponseUrl returns a boolean if a field has been set.
func (o *VacancyDraftVacancyDraftEdit) HasResponseUrl() bool {
	if o != nil && o.ResponseUrl.IsSet() {
		return true
	}

	return false
}

// SetResponseUrl gets a reference to the given NullableString and assigns it to the ResponseUrl field.
func (o *VacancyDraftVacancyDraftEdit) SetResponseUrl(v string) {
	o.ResponseUrl.Set(&v)
}
// SetResponseUrlNil sets the value for ResponseUrl to be an explicit nil
func (o *VacancyDraftVacancyDraftEdit) SetResponseUrlNil() {
	o.ResponseUrl.Set(nil)
}

// UnsetResponseUrl ensures that no value is present for ResponseUrl, not even an explicit nil
func (o *VacancyDraftVacancyDraftEdit) UnsetResponseUrl() {
	o.ResponseUrl.Unset()
}

// GetSalary returns the Salary field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *VacancyDraftVacancyDraftEdit) GetSalary() VacancySalary {
	if o == nil || IsNil(o.Salary.Get()) {
		var ret VacancySalary
		return ret
	}
	return *o.Salary.Get()
}

// GetSalaryOk returns a tuple with the Salary field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *VacancyDraftVacancyDraftEdit) GetSalaryOk() (*VacancySalary, bool) {
	if o == nil {
		return nil, false
	}
	return o.Salary.Get(), o.Salary.IsSet()
}

// HasSalary returns a boolean if a field has been set.
func (o *VacancyDraftVacancyDraftEdit) HasSalary() bool {
	if o != nil && o.Salary.IsSet() {
		return true
	}

	return false
}

// SetSalary gets a reference to the given NullableVacancySalary and assigns it to the Salary field.
func (o *VacancyDraftVacancyDraftEdit) SetSalary(v VacancySalary) {
	o.Salary.Set(&v)
}
// SetSalaryNil sets the value for Salary to be an explicit nil
func (o *VacancyDraftVacancyDraftEdit) SetSalaryNil() {
	o.Salary.Set(nil)
}

// UnsetSalary ensures that no value is present for Salary, not even an explicit nil
func (o *VacancyDraftVacancyDraftEdit) UnsetSalary() {
	o.Salary.Unset()
}

// GetSchedule returns the Schedule field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *VacancyDraftVacancyDraftEdit) GetSchedule() VacancySchedule {
	if o == nil || IsNil(o.Schedule.Get()) {
		var ret VacancySchedule
		return ret
	}
	return *o.Schedule.Get()
}

// GetScheduleOk returns a tuple with the Schedule field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *VacancyDraftVacancyDraftEdit) GetScheduleOk() (*VacancySchedule, bool) {
	if o == nil {
		return nil, false
	}
	return o.Schedule.Get(), o.Schedule.IsSet()
}

// HasSchedule returns a boolean if a field has been set.
func (o *VacancyDraftVacancyDraftEdit) HasSchedule() bool {
	if o != nil && o.Schedule.IsSet() {
		return true
	}

	return false
}

// SetSchedule gets a reference to the given NullableVacancySchedule and assigns it to the Schedule field.
func (o *VacancyDraftVacancyDraftEdit) SetSchedule(v VacancySchedule) {
	o.Schedule.Set(&v)
}
// SetScheduleNil sets the value for Schedule to be an explicit nil
func (o *VacancyDraftVacancyDraftEdit) SetScheduleNil() {
	o.Schedule.Set(nil)
}

// UnsetSchedule ensures that no value is present for Schedule, not even an explicit nil
func (o *VacancyDraftVacancyDraftEdit) UnsetSchedule() {
	o.Schedule.Unset()
}

// GetTest returns the Test field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *VacancyDraftVacancyDraftEdit) GetTest() VacancyDraftTest {
	if o == nil || IsNil(o.Test.Get()) {
		var ret VacancyDraftTest
		return ret
	}
	return *o.Test.Get()
}

// GetTestOk returns a tuple with the Test field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *VacancyDraftVacancyDraftEdit) GetTestOk() (*VacancyDraftTest, bool) {
	if o == nil {
		return nil, false
	}
	return o.Test.Get(), o.Test.IsSet()
}

// HasTest returns a boolean if a field has been set.
func (o *VacancyDraftVacancyDraftEdit) HasTest() bool {
	if o != nil && o.Test.IsSet() {
		return true
	}

	return false
}

// SetTest gets a reference to the given NullableVacancyDraftTest and assigns it to the Test field.
func (o *VacancyDraftVacancyDraftEdit) SetTest(v VacancyDraftTest) {
	o.Test.Set(&v)
}
// SetTestNil sets the value for Test to be an explicit nil
func (o *VacancyDraftVacancyDraftEdit) SetTestNil() {
	o.Test.Set(nil)
}

// UnsetTest ensures that no value is present for Test, not even an explicit nil
func (o *VacancyDraftVacancyDraftEdit) UnsetTest() {
	o.Test.Unset()
}

// GetType returns the Type field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *VacancyDraftVacancyDraftEdit) GetType() VacancyDraftType {
	if o == nil || IsNil(o.Type.Get()) {
		var ret VacancyDraftType
		return ret
	}
	return *o.Type.Get()
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *VacancyDraftVacancyDraftEdit) GetTypeOk() (*VacancyDraftType, bool) {
	if o == nil {
		return nil, false
	}
	return o.Type.Get(), o.Type.IsSet()
}

// HasType returns a boolean if a field has been set.
func (o *VacancyDraftVacancyDraftEdit) HasType() bool {
	if o != nil && o.Type.IsSet() {
		return true
	}

	return false
}

// SetType gets a reference to the given NullableVacancyDraftType and assigns it to the Type field.
func (o *VacancyDraftVacancyDraftEdit) SetType(v VacancyDraftType) {
	o.Type.Set(&v)
}
// SetTypeNil sets the value for Type to be an explicit nil
func (o *VacancyDraftVacancyDraftEdit) SetTypeNil() {
	o.Type.Set(nil)
}

// UnsetType ensures that no value is present for Type, not even an explicit nil
func (o *VacancyDraftVacancyDraftEdit) UnsetType() {
	o.Type.Unset()
}

// GetWithZp returns the WithZp field value if set, zero value otherwise.
func (o *VacancyDraftVacancyDraftEdit) GetWithZp() bool {
	if o == nil || IsNil(o.WithZp) {
		var ret bool
		return ret
	}
	return *o.WithZp
}

// GetWithZpOk returns a tuple with the WithZp field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VacancyDraftVacancyDraftEdit) GetWithZpOk() (*bool, bool) {
	if o == nil || IsNil(o.WithZp) {
		return nil, false
	}
	return o.WithZp, true
}

// HasWithZp returns a boolean if a field has been set.
func (o *VacancyDraftVacancyDraftEdit) HasWithZp() bool {
	if o != nil && !IsNil(o.WithZp) {
		return true
	}

	return false
}

// SetWithZp gets a reference to the given bool and assigns it to the WithZp field.
func (o *VacancyDraftVacancyDraftEdit) SetWithZp(v bool) {
	o.WithZp = &v
}

// GetWorkingDays returns the WorkingDays field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *VacancyDraftVacancyDraftEdit) GetWorkingDays() []VacancyWorkingDayItem {
	if o == nil {
		var ret []VacancyWorkingDayItem
		return ret
	}
	return o.WorkingDays
}

// GetWorkingDaysOk returns a tuple with the WorkingDays field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *VacancyDraftVacancyDraftEdit) GetWorkingDaysOk() ([]VacancyWorkingDayItem, bool) {
	if o == nil || IsNil(o.WorkingDays) {
		return nil, false
	}
	return o.WorkingDays, true
}

// HasWorkingDays returns a boolean if a field has been set.
func (o *VacancyDraftVacancyDraftEdit) HasWorkingDays() bool {
	if o != nil && !IsNil(o.WorkingDays) {
		return true
	}

	return false
}

// SetWorkingDays gets a reference to the given []VacancyWorkingDayItem and assigns it to the WorkingDays field.
func (o *VacancyDraftVacancyDraftEdit) SetWorkingDays(v []VacancyWorkingDayItem) {
	o.WorkingDays = v
}

// GetWorkingTimeIntervals returns the WorkingTimeIntervals field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *VacancyDraftVacancyDraftEdit) GetWorkingTimeIntervals() []VacancyWorkingTimeIntervalItem {
	if o == nil {
		var ret []VacancyWorkingTimeIntervalItem
		return ret
	}
	return o.WorkingTimeIntervals
}

// GetWorkingTimeIntervalsOk returns a tuple with the WorkingTimeIntervals field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *VacancyDraftVacancyDraftEdit) GetWorkingTimeIntervalsOk() ([]VacancyWorkingTimeIntervalItem, bool) {
	if o == nil || IsNil(o.WorkingTimeIntervals) {
		return nil, false
	}
	return o.WorkingTimeIntervals, true
}

// HasWorkingTimeIntervals returns a boolean if a field has been set.
func (o *VacancyDraftVacancyDraftEdit) HasWorkingTimeIntervals() bool {
	if o != nil && !IsNil(o.WorkingTimeIntervals) {
		return true
	}

	return false
}

// SetWorkingTimeIntervals gets a reference to the given []VacancyWorkingTimeIntervalItem and assigns it to the WorkingTimeIntervals field.
func (o *VacancyDraftVacancyDraftEdit) SetWorkingTimeIntervals(v []VacancyWorkingTimeIntervalItem) {
	o.WorkingTimeIntervals = v
}

// GetWorkingTimeModes returns the WorkingTimeModes field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *VacancyDraftVacancyDraftEdit) GetWorkingTimeModes() []VacancyWorkingTimeModeItem {
	if o == nil {
		var ret []VacancyWorkingTimeModeItem
		return ret
	}
	return o.WorkingTimeModes
}

// GetWorkingTimeModesOk returns a tuple with the WorkingTimeModes field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *VacancyDraftVacancyDraftEdit) GetWorkingTimeModesOk() ([]VacancyWorkingTimeModeItem, bool) {
	if o == nil || IsNil(o.WorkingTimeModes) {
		return nil, false
	}
	return o.WorkingTimeModes, true
}

// HasWorkingTimeModes returns a boolean if a field has been set.
func (o *VacancyDraftVacancyDraftEdit) HasWorkingTimeModes() bool {
	if o != nil && !IsNil(o.WorkingTimeModes) {
		return true
	}

	return false
}

// SetWorkingTimeModes gets a reference to the given []VacancyWorkingTimeModeItem and assigns it to the WorkingTimeModes field.
func (o *VacancyDraftVacancyDraftEdit) SetWorkingTimeModes(v []VacancyWorkingTimeModeItem) {
	o.WorkingTimeModes = v
}

func (o VacancyDraftVacancyDraftEdit) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o VacancyDraftVacancyDraftEdit) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.AcceptHandicapped) {
		toSerialize["accept_handicapped"] = o.AcceptHandicapped
	}
	if !IsNil(o.AcceptIncompleteResumes) {
		toSerialize["accept_incomplete_resumes"] = o.AcceptIncompleteResumes
	}
	if !IsNil(o.AcceptKids) {
		toSerialize["accept_kids"] = o.AcceptKids
	}
	if o.AcceptTemporary.IsSet() {
		toSerialize["accept_temporary"] = o.AcceptTemporary.Get()
	}
	if o.Address.IsSet() {
		toSerialize["address"] = o.Address.Get()
	}
	if !IsNil(o.AllowMessages) {
		toSerialize["allow_messages"] = o.AllowMessages
	}
	if !IsNil(o.Area) {
		toSerialize["area"] = o.Area
	}
	if o.Areas != nil {
		toSerialize["areas"] = o.Areas
	}
	if o.AssignedManagerId.IsSet() {
		toSerialize["assigned_manager_id"] = o.AssignedManagerId.Get()
	}
	if o.BillingType.IsSet() {
		toSerialize["billing_type"] = o.BillingType.Get()
	}
	if !IsNil(o.BrandedTemplate) {
		toSerialize["branded_template"] = o.BrandedTemplate
	}
	if o.Code.IsSet() {
		toSerialize["code"] = o.Code.Get()
	}
	if o.Contacts.IsSet() {
		toSerialize["contacts"] = o.Contacts.Get()
	}
	if o.CustomEmployerName.IsSet() {
		toSerialize["custom_employer_name"] = o.CustomEmployerName.Get()
	}
	if o.Department.IsSet() {
		toSerialize["department"] = o.Department.Get()
	}
	if o.Description.IsSet() {
		toSerialize["description"] = o.Description.Get()
	}
	if !IsNil(o.DriverLicenseTypes) {
		toSerialize["driver_license_types"] = o.DriverLicenseTypes
	}
	if o.Employment.IsSet() {
		toSerialize["employment"] = o.Employment.Get()
	}
	if o.Experience.IsSet() {
		toSerialize["experience"] = o.Experience.Get()
	}
	if o.KeySkills != nil {
		toSerialize["key_skills"] = o.KeySkills
	}
	if !IsNil(o.Languages) {
		toSerialize["languages"] = o.Languages
	}
	if o.Name.IsSet() {
		toSerialize["name"] = o.Name.Get()
	}
	if o.ProfessionalRoles != nil {
		toSerialize["professional_roles"] = o.ProfessionalRoles
	}
	if !IsNil(o.ResponseLetterRequired) {
		toSerialize["response_letter_required"] = o.ResponseLetterRequired
	}
	if !IsNil(o.ResponseNotifications) {
		toSerialize["response_notifications"] = o.ResponseNotifications
	}
	if o.ResponseUrl.IsSet() {
		toSerialize["response_url"] = o.ResponseUrl.Get()
	}
	if o.Salary.IsSet() {
		toSerialize["salary"] = o.Salary.Get()
	}
	if o.Schedule.IsSet() {
		toSerialize["schedule"] = o.Schedule.Get()
	}
	if o.Test.IsSet() {
		toSerialize["test"] = o.Test.Get()
	}
	if o.Type.IsSet() {
		toSerialize["type"] = o.Type.Get()
	}
	if !IsNil(o.WithZp) {
		toSerialize["with_zp"] = o.WithZp
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

type NullableVacancyDraftVacancyDraftEdit struct {
	value *VacancyDraftVacancyDraftEdit
	isSet bool
}

func (v NullableVacancyDraftVacancyDraftEdit) Get() *VacancyDraftVacancyDraftEdit {
	return v.value
}

func (v *NullableVacancyDraftVacancyDraftEdit) Set(val *VacancyDraftVacancyDraftEdit) {
	v.value = val
	v.isSet = true
}

func (v NullableVacancyDraftVacancyDraftEdit) IsSet() bool {
	return v.isSet
}

func (v *NullableVacancyDraftVacancyDraftEdit) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableVacancyDraftVacancyDraftEdit(val *VacancyDraftVacancyDraftEdit) *NullableVacancyDraftVacancyDraftEdit {
	return &NullableVacancyDraftVacancyDraftEdit{value: val, isSet: true}
}

func (v NullableVacancyDraftVacancyDraftEdit) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableVacancyDraftVacancyDraftEdit) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


