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

// checks if the VacanciesVacancyConditions type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &VacanciesVacancyConditions{}

// VacanciesVacancyConditions Условия заполнения полей при [публикации вакансии](#tag/Upravlenie-vakansiyami/operation/publish-vacancy)
type VacanciesVacancyConditions struct {
	AcceptHandicapped NullableVacanciesVacancyConditionFieldsRequiredWithTitle `json:"accept_handicapped,omitempty"`
	AcceptIncompleteResumes NullableVacanciesVacancyConditionFieldsRequiredWithTitle `json:"accept_incomplete_resumes,omitempty"`
	AcceptKids NullableVacanciesVacancyConditionFieldsRequiredWithTitle `json:"accept_kids,omitempty"`
	AcceptTemporary NullableVacanciesVacancyConditionFieldsRequiredWithTitle `json:"accept_temporary,omitempty"`
	Address NullableVacanciesVacancyConditionFieldsAddressCondition `json:"address,omitempty"`
	AllowMessages NullableVacanciesVacancyConditionFieldsRequiredWithTitle `json:"allow_messages,omitempty"`
	Area NullableVacanciesVacancyConditionFieldsRequiredWithTitle `json:"area,omitempty"`
	BillingType NullableVacanciesVacancyConditionFieldsRequiredWithTitle `json:"billing_type,omitempty"`
	BrandedTemplate NullableVacanciesVacancyConditionFieldsRequiredLengthWithTitle `json:"branded_template,omitempty"`
	Code NullableVacanciesVacancyConditionFieldsRequiredLengthWithTitle `json:"code,omitempty"`
	Contacts NullableVacanciesVacancyConditionFieldsContactsCondition `json:"contacts,omitempty"`
	CustomEmployerName NullableVacanciesVacancyConditionFieldsRequiredLengthWithTitle `json:"custom_employer_name,omitempty"`
	Department NullableVacanciesVacancyConditionFieldsRequiredLengthWithTitle `json:"department,omitempty"`
	Description NullableVacanciesVacancyConditionFieldsRequiredLengthWithTitle `json:"description,omitempty"`
	DriverLicenseTypes NullableVacanciesVacancyConditionFieldsRequiredCountWithTitle `json:"driver_license_types,omitempty"`
	Employment NullableVacanciesVacancyConditionFieldsRequiredWithTitle `json:"employment,omitempty"`
	Experience NullableVacanciesVacancyConditionFieldsRequiredWithTitle `json:"experience,omitempty"`
	KeySkills NullableVacanciesVacancyConditionFieldsRequiredCountWithTitle `json:"key_skills,omitempty"`
	Languages NullableVacanciesVacancyConditionFieldsRequiredCountWithTitle `json:"languages,omitempty"`
	Manager NullableVacanciesVacancyConditionFieldsRequiredWithTitle `json:"manager,omitempty"`
	Name NullableVacanciesVacancyConditionFieldsRequiredLengthWithTitle `json:"name,omitempty"`
	ProfessionalRoles NullableVacanciesVacancyConditionFieldsRequiredCountWithTitle `json:"professional_roles,omitempty"`
	ResponseLetterRequired NullableVacanciesVacancyConditionFieldsRequiredWithTitle `json:"response_letter_required,omitempty"`
	ResponseNotifications NullableVacanciesVacancyConditionFieldsRequiredWithTitle `json:"response_notifications,omitempty"`
	ResponseUrl NullableVacanciesVacancyConditionFieldsResponseUrlCondition `json:"response_url,omitempty"`
	Salary NullableVacanciesVacancyConditionFieldsSalaryCondition `json:"salary,omitempty"`
	Schedule NullableVacanciesVacancyConditionFieldsRequiredWithTitle `json:"schedule,omitempty"`
	Test NullableVacanciesVacancyConditionFieldsTestCondition `json:"test,omitempty"`
	Type NullableVacanciesVacancyConditionFieldsRequiredWithTitle `json:"type,omitempty"`
	WorkingDays NullableVacanciesVacancyConditionFieldsRequiredCountWithTitle `json:"working_days,omitempty"`
	WorkingTimeIntervals NullableVacanciesVacancyConditionFieldsRequiredCountWithTitle `json:"working_time_intervals,omitempty"`
	WorkingTimeModes NullableVacanciesVacancyConditionFieldsRequiredCountWithTitle `json:"working_time_modes,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _VacanciesVacancyConditions VacanciesVacancyConditions

// NewVacanciesVacancyConditions instantiates a new VacanciesVacancyConditions object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewVacanciesVacancyConditions() *VacanciesVacancyConditions {
	this := VacanciesVacancyConditions{}
	return &this
}

// NewVacanciesVacancyConditionsWithDefaults instantiates a new VacanciesVacancyConditions object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewVacanciesVacancyConditionsWithDefaults() *VacanciesVacancyConditions {
	this := VacanciesVacancyConditions{}
	return &this
}

// GetAcceptHandicapped returns the AcceptHandicapped field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *VacanciesVacancyConditions) GetAcceptHandicapped() VacanciesVacancyConditionFieldsRequiredWithTitle {
	if o == nil || IsNil(o.AcceptHandicapped.Get()) {
		var ret VacanciesVacancyConditionFieldsRequiredWithTitle
		return ret
	}
	return *o.AcceptHandicapped.Get()
}

// GetAcceptHandicappedOk returns a tuple with the AcceptHandicapped field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *VacanciesVacancyConditions) GetAcceptHandicappedOk() (*VacanciesVacancyConditionFieldsRequiredWithTitle, bool) {
	if o == nil {
		return nil, false
	}
	return o.AcceptHandicapped.Get(), o.AcceptHandicapped.IsSet()
}

// HasAcceptHandicapped returns a boolean if a field has been set.
func (o *VacanciesVacancyConditions) HasAcceptHandicapped() bool {
	if o != nil && o.AcceptHandicapped.IsSet() {
		return true
	}

	return false
}

// SetAcceptHandicapped gets a reference to the given NullableVacanciesVacancyConditionFieldsRequiredWithTitle and assigns it to the AcceptHandicapped field.
func (o *VacanciesVacancyConditions) SetAcceptHandicapped(v VacanciesVacancyConditionFieldsRequiredWithTitle) {
	o.AcceptHandicapped.Set(&v)
}
// SetAcceptHandicappedNil sets the value for AcceptHandicapped to be an explicit nil
func (o *VacanciesVacancyConditions) SetAcceptHandicappedNil() {
	o.AcceptHandicapped.Set(nil)
}

// UnsetAcceptHandicapped ensures that no value is present for AcceptHandicapped, not even an explicit nil
func (o *VacanciesVacancyConditions) UnsetAcceptHandicapped() {
	o.AcceptHandicapped.Unset()
}

// GetAcceptIncompleteResumes returns the AcceptIncompleteResumes field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *VacanciesVacancyConditions) GetAcceptIncompleteResumes() VacanciesVacancyConditionFieldsRequiredWithTitle {
	if o == nil || IsNil(o.AcceptIncompleteResumes.Get()) {
		var ret VacanciesVacancyConditionFieldsRequiredWithTitle
		return ret
	}
	return *o.AcceptIncompleteResumes.Get()
}

// GetAcceptIncompleteResumesOk returns a tuple with the AcceptIncompleteResumes field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *VacanciesVacancyConditions) GetAcceptIncompleteResumesOk() (*VacanciesVacancyConditionFieldsRequiredWithTitle, bool) {
	if o == nil {
		return nil, false
	}
	return o.AcceptIncompleteResumes.Get(), o.AcceptIncompleteResumes.IsSet()
}

// HasAcceptIncompleteResumes returns a boolean if a field has been set.
func (o *VacanciesVacancyConditions) HasAcceptIncompleteResumes() bool {
	if o != nil && o.AcceptIncompleteResumes.IsSet() {
		return true
	}

	return false
}

// SetAcceptIncompleteResumes gets a reference to the given NullableVacanciesVacancyConditionFieldsRequiredWithTitle and assigns it to the AcceptIncompleteResumes field.
func (o *VacanciesVacancyConditions) SetAcceptIncompleteResumes(v VacanciesVacancyConditionFieldsRequiredWithTitle) {
	o.AcceptIncompleteResumes.Set(&v)
}
// SetAcceptIncompleteResumesNil sets the value for AcceptIncompleteResumes to be an explicit nil
func (o *VacanciesVacancyConditions) SetAcceptIncompleteResumesNil() {
	o.AcceptIncompleteResumes.Set(nil)
}

// UnsetAcceptIncompleteResumes ensures that no value is present for AcceptIncompleteResumes, not even an explicit nil
func (o *VacanciesVacancyConditions) UnsetAcceptIncompleteResumes() {
	o.AcceptIncompleteResumes.Unset()
}

// GetAcceptKids returns the AcceptKids field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *VacanciesVacancyConditions) GetAcceptKids() VacanciesVacancyConditionFieldsRequiredWithTitle {
	if o == nil || IsNil(o.AcceptKids.Get()) {
		var ret VacanciesVacancyConditionFieldsRequiredWithTitle
		return ret
	}
	return *o.AcceptKids.Get()
}

// GetAcceptKidsOk returns a tuple with the AcceptKids field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *VacanciesVacancyConditions) GetAcceptKidsOk() (*VacanciesVacancyConditionFieldsRequiredWithTitle, bool) {
	if o == nil {
		return nil, false
	}
	return o.AcceptKids.Get(), o.AcceptKids.IsSet()
}

// HasAcceptKids returns a boolean if a field has been set.
func (o *VacanciesVacancyConditions) HasAcceptKids() bool {
	if o != nil && o.AcceptKids.IsSet() {
		return true
	}

	return false
}

// SetAcceptKids gets a reference to the given NullableVacanciesVacancyConditionFieldsRequiredWithTitle and assigns it to the AcceptKids field.
func (o *VacanciesVacancyConditions) SetAcceptKids(v VacanciesVacancyConditionFieldsRequiredWithTitle) {
	o.AcceptKids.Set(&v)
}
// SetAcceptKidsNil sets the value for AcceptKids to be an explicit nil
func (o *VacanciesVacancyConditions) SetAcceptKidsNil() {
	o.AcceptKids.Set(nil)
}

// UnsetAcceptKids ensures that no value is present for AcceptKids, not even an explicit nil
func (o *VacanciesVacancyConditions) UnsetAcceptKids() {
	o.AcceptKids.Unset()
}

// GetAcceptTemporary returns the AcceptTemporary field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *VacanciesVacancyConditions) GetAcceptTemporary() VacanciesVacancyConditionFieldsRequiredWithTitle {
	if o == nil || IsNil(o.AcceptTemporary.Get()) {
		var ret VacanciesVacancyConditionFieldsRequiredWithTitle
		return ret
	}
	return *o.AcceptTemporary.Get()
}

// GetAcceptTemporaryOk returns a tuple with the AcceptTemporary field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *VacanciesVacancyConditions) GetAcceptTemporaryOk() (*VacanciesVacancyConditionFieldsRequiredWithTitle, bool) {
	if o == nil {
		return nil, false
	}
	return o.AcceptTemporary.Get(), o.AcceptTemporary.IsSet()
}

// HasAcceptTemporary returns a boolean if a field has been set.
func (o *VacanciesVacancyConditions) HasAcceptTemporary() bool {
	if o != nil && o.AcceptTemporary.IsSet() {
		return true
	}

	return false
}

// SetAcceptTemporary gets a reference to the given NullableVacanciesVacancyConditionFieldsRequiredWithTitle and assigns it to the AcceptTemporary field.
func (o *VacanciesVacancyConditions) SetAcceptTemporary(v VacanciesVacancyConditionFieldsRequiredWithTitle) {
	o.AcceptTemporary.Set(&v)
}
// SetAcceptTemporaryNil sets the value for AcceptTemporary to be an explicit nil
func (o *VacanciesVacancyConditions) SetAcceptTemporaryNil() {
	o.AcceptTemporary.Set(nil)
}

// UnsetAcceptTemporary ensures that no value is present for AcceptTemporary, not even an explicit nil
func (o *VacanciesVacancyConditions) UnsetAcceptTemporary() {
	o.AcceptTemporary.Unset()
}

// GetAddress returns the Address field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *VacanciesVacancyConditions) GetAddress() VacanciesVacancyConditionFieldsAddressCondition {
	if o == nil || IsNil(o.Address.Get()) {
		var ret VacanciesVacancyConditionFieldsAddressCondition
		return ret
	}
	return *o.Address.Get()
}

// GetAddressOk returns a tuple with the Address field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *VacanciesVacancyConditions) GetAddressOk() (*VacanciesVacancyConditionFieldsAddressCondition, bool) {
	if o == nil {
		return nil, false
	}
	return o.Address.Get(), o.Address.IsSet()
}

// HasAddress returns a boolean if a field has been set.
func (o *VacanciesVacancyConditions) HasAddress() bool {
	if o != nil && o.Address.IsSet() {
		return true
	}

	return false
}

// SetAddress gets a reference to the given NullableVacanciesVacancyConditionFieldsAddressCondition and assigns it to the Address field.
func (o *VacanciesVacancyConditions) SetAddress(v VacanciesVacancyConditionFieldsAddressCondition) {
	o.Address.Set(&v)
}
// SetAddressNil sets the value for Address to be an explicit nil
func (o *VacanciesVacancyConditions) SetAddressNil() {
	o.Address.Set(nil)
}

// UnsetAddress ensures that no value is present for Address, not even an explicit nil
func (o *VacanciesVacancyConditions) UnsetAddress() {
	o.Address.Unset()
}

// GetAllowMessages returns the AllowMessages field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *VacanciesVacancyConditions) GetAllowMessages() VacanciesVacancyConditionFieldsRequiredWithTitle {
	if o == nil || IsNil(o.AllowMessages.Get()) {
		var ret VacanciesVacancyConditionFieldsRequiredWithTitle
		return ret
	}
	return *o.AllowMessages.Get()
}

// GetAllowMessagesOk returns a tuple with the AllowMessages field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *VacanciesVacancyConditions) GetAllowMessagesOk() (*VacanciesVacancyConditionFieldsRequiredWithTitle, bool) {
	if o == nil {
		return nil, false
	}
	return o.AllowMessages.Get(), o.AllowMessages.IsSet()
}

// HasAllowMessages returns a boolean if a field has been set.
func (o *VacanciesVacancyConditions) HasAllowMessages() bool {
	if o != nil && o.AllowMessages.IsSet() {
		return true
	}

	return false
}

// SetAllowMessages gets a reference to the given NullableVacanciesVacancyConditionFieldsRequiredWithTitle and assigns it to the AllowMessages field.
func (o *VacanciesVacancyConditions) SetAllowMessages(v VacanciesVacancyConditionFieldsRequiredWithTitle) {
	o.AllowMessages.Set(&v)
}
// SetAllowMessagesNil sets the value for AllowMessages to be an explicit nil
func (o *VacanciesVacancyConditions) SetAllowMessagesNil() {
	o.AllowMessages.Set(nil)
}

// UnsetAllowMessages ensures that no value is present for AllowMessages, not even an explicit nil
func (o *VacanciesVacancyConditions) UnsetAllowMessages() {
	o.AllowMessages.Unset()
}

// GetArea returns the Area field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *VacanciesVacancyConditions) GetArea() VacanciesVacancyConditionFieldsRequiredWithTitle {
	if o == nil || IsNil(o.Area.Get()) {
		var ret VacanciesVacancyConditionFieldsRequiredWithTitle
		return ret
	}
	return *o.Area.Get()
}

// GetAreaOk returns a tuple with the Area field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *VacanciesVacancyConditions) GetAreaOk() (*VacanciesVacancyConditionFieldsRequiredWithTitle, bool) {
	if o == nil {
		return nil, false
	}
	return o.Area.Get(), o.Area.IsSet()
}

// HasArea returns a boolean if a field has been set.
func (o *VacanciesVacancyConditions) HasArea() bool {
	if o != nil && o.Area.IsSet() {
		return true
	}

	return false
}

// SetArea gets a reference to the given NullableVacanciesVacancyConditionFieldsRequiredWithTitle and assigns it to the Area field.
func (o *VacanciesVacancyConditions) SetArea(v VacanciesVacancyConditionFieldsRequiredWithTitle) {
	o.Area.Set(&v)
}
// SetAreaNil sets the value for Area to be an explicit nil
func (o *VacanciesVacancyConditions) SetAreaNil() {
	o.Area.Set(nil)
}

// UnsetArea ensures that no value is present for Area, not even an explicit nil
func (o *VacanciesVacancyConditions) UnsetArea() {
	o.Area.Unset()
}

// GetBillingType returns the BillingType field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *VacanciesVacancyConditions) GetBillingType() VacanciesVacancyConditionFieldsRequiredWithTitle {
	if o == nil || IsNil(o.BillingType.Get()) {
		var ret VacanciesVacancyConditionFieldsRequiredWithTitle
		return ret
	}
	return *o.BillingType.Get()
}

// GetBillingTypeOk returns a tuple with the BillingType field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *VacanciesVacancyConditions) GetBillingTypeOk() (*VacanciesVacancyConditionFieldsRequiredWithTitle, bool) {
	if o == nil {
		return nil, false
	}
	return o.BillingType.Get(), o.BillingType.IsSet()
}

// HasBillingType returns a boolean if a field has been set.
func (o *VacanciesVacancyConditions) HasBillingType() bool {
	if o != nil && o.BillingType.IsSet() {
		return true
	}

	return false
}

// SetBillingType gets a reference to the given NullableVacanciesVacancyConditionFieldsRequiredWithTitle and assigns it to the BillingType field.
func (o *VacanciesVacancyConditions) SetBillingType(v VacanciesVacancyConditionFieldsRequiredWithTitle) {
	o.BillingType.Set(&v)
}
// SetBillingTypeNil sets the value for BillingType to be an explicit nil
func (o *VacanciesVacancyConditions) SetBillingTypeNil() {
	o.BillingType.Set(nil)
}

// UnsetBillingType ensures that no value is present for BillingType, not even an explicit nil
func (o *VacanciesVacancyConditions) UnsetBillingType() {
	o.BillingType.Unset()
}

// GetBrandedTemplate returns the BrandedTemplate field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *VacanciesVacancyConditions) GetBrandedTemplate() VacanciesVacancyConditionFieldsRequiredLengthWithTitle {
	if o == nil || IsNil(o.BrandedTemplate.Get()) {
		var ret VacanciesVacancyConditionFieldsRequiredLengthWithTitle
		return ret
	}
	return *o.BrandedTemplate.Get()
}

// GetBrandedTemplateOk returns a tuple with the BrandedTemplate field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *VacanciesVacancyConditions) GetBrandedTemplateOk() (*VacanciesVacancyConditionFieldsRequiredLengthWithTitle, bool) {
	if o == nil {
		return nil, false
	}
	return o.BrandedTemplate.Get(), o.BrandedTemplate.IsSet()
}

// HasBrandedTemplate returns a boolean if a field has been set.
func (o *VacanciesVacancyConditions) HasBrandedTemplate() bool {
	if o != nil && o.BrandedTemplate.IsSet() {
		return true
	}

	return false
}

// SetBrandedTemplate gets a reference to the given NullableVacanciesVacancyConditionFieldsRequiredLengthWithTitle and assigns it to the BrandedTemplate field.
func (o *VacanciesVacancyConditions) SetBrandedTemplate(v VacanciesVacancyConditionFieldsRequiredLengthWithTitle) {
	o.BrandedTemplate.Set(&v)
}
// SetBrandedTemplateNil sets the value for BrandedTemplate to be an explicit nil
func (o *VacanciesVacancyConditions) SetBrandedTemplateNil() {
	o.BrandedTemplate.Set(nil)
}

// UnsetBrandedTemplate ensures that no value is present for BrandedTemplate, not even an explicit nil
func (o *VacanciesVacancyConditions) UnsetBrandedTemplate() {
	o.BrandedTemplate.Unset()
}

// GetCode returns the Code field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *VacanciesVacancyConditions) GetCode() VacanciesVacancyConditionFieldsRequiredLengthWithTitle {
	if o == nil || IsNil(o.Code.Get()) {
		var ret VacanciesVacancyConditionFieldsRequiredLengthWithTitle
		return ret
	}
	return *o.Code.Get()
}

// GetCodeOk returns a tuple with the Code field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *VacanciesVacancyConditions) GetCodeOk() (*VacanciesVacancyConditionFieldsRequiredLengthWithTitle, bool) {
	if o == nil {
		return nil, false
	}
	return o.Code.Get(), o.Code.IsSet()
}

// HasCode returns a boolean if a field has been set.
func (o *VacanciesVacancyConditions) HasCode() bool {
	if o != nil && o.Code.IsSet() {
		return true
	}

	return false
}

// SetCode gets a reference to the given NullableVacanciesVacancyConditionFieldsRequiredLengthWithTitle and assigns it to the Code field.
func (o *VacanciesVacancyConditions) SetCode(v VacanciesVacancyConditionFieldsRequiredLengthWithTitle) {
	o.Code.Set(&v)
}
// SetCodeNil sets the value for Code to be an explicit nil
func (o *VacanciesVacancyConditions) SetCodeNil() {
	o.Code.Set(nil)
}

// UnsetCode ensures that no value is present for Code, not even an explicit nil
func (o *VacanciesVacancyConditions) UnsetCode() {
	o.Code.Unset()
}

// GetContacts returns the Contacts field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *VacanciesVacancyConditions) GetContacts() VacanciesVacancyConditionFieldsContactsCondition {
	if o == nil || IsNil(o.Contacts.Get()) {
		var ret VacanciesVacancyConditionFieldsContactsCondition
		return ret
	}
	return *o.Contacts.Get()
}

// GetContactsOk returns a tuple with the Contacts field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *VacanciesVacancyConditions) GetContactsOk() (*VacanciesVacancyConditionFieldsContactsCondition, bool) {
	if o == nil {
		return nil, false
	}
	return o.Contacts.Get(), o.Contacts.IsSet()
}

// HasContacts returns a boolean if a field has been set.
func (o *VacanciesVacancyConditions) HasContacts() bool {
	if o != nil && o.Contacts.IsSet() {
		return true
	}

	return false
}

// SetContacts gets a reference to the given NullableVacanciesVacancyConditionFieldsContactsCondition and assigns it to the Contacts field.
func (o *VacanciesVacancyConditions) SetContacts(v VacanciesVacancyConditionFieldsContactsCondition) {
	o.Contacts.Set(&v)
}
// SetContactsNil sets the value for Contacts to be an explicit nil
func (o *VacanciesVacancyConditions) SetContactsNil() {
	o.Contacts.Set(nil)
}

// UnsetContacts ensures that no value is present for Contacts, not even an explicit nil
func (o *VacanciesVacancyConditions) UnsetContacts() {
	o.Contacts.Unset()
}

// GetCustomEmployerName returns the CustomEmployerName field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *VacanciesVacancyConditions) GetCustomEmployerName() VacanciesVacancyConditionFieldsRequiredLengthWithTitle {
	if o == nil || IsNil(o.CustomEmployerName.Get()) {
		var ret VacanciesVacancyConditionFieldsRequiredLengthWithTitle
		return ret
	}
	return *o.CustomEmployerName.Get()
}

// GetCustomEmployerNameOk returns a tuple with the CustomEmployerName field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *VacanciesVacancyConditions) GetCustomEmployerNameOk() (*VacanciesVacancyConditionFieldsRequiredLengthWithTitle, bool) {
	if o == nil {
		return nil, false
	}
	return o.CustomEmployerName.Get(), o.CustomEmployerName.IsSet()
}

// HasCustomEmployerName returns a boolean if a field has been set.
func (o *VacanciesVacancyConditions) HasCustomEmployerName() bool {
	if o != nil && o.CustomEmployerName.IsSet() {
		return true
	}

	return false
}

// SetCustomEmployerName gets a reference to the given NullableVacanciesVacancyConditionFieldsRequiredLengthWithTitle and assigns it to the CustomEmployerName field.
func (o *VacanciesVacancyConditions) SetCustomEmployerName(v VacanciesVacancyConditionFieldsRequiredLengthWithTitle) {
	o.CustomEmployerName.Set(&v)
}
// SetCustomEmployerNameNil sets the value for CustomEmployerName to be an explicit nil
func (o *VacanciesVacancyConditions) SetCustomEmployerNameNil() {
	o.CustomEmployerName.Set(nil)
}

// UnsetCustomEmployerName ensures that no value is present for CustomEmployerName, not even an explicit nil
func (o *VacanciesVacancyConditions) UnsetCustomEmployerName() {
	o.CustomEmployerName.Unset()
}

// GetDepartment returns the Department field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *VacanciesVacancyConditions) GetDepartment() VacanciesVacancyConditionFieldsRequiredLengthWithTitle {
	if o == nil || IsNil(o.Department.Get()) {
		var ret VacanciesVacancyConditionFieldsRequiredLengthWithTitle
		return ret
	}
	return *o.Department.Get()
}

// GetDepartmentOk returns a tuple with the Department field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *VacanciesVacancyConditions) GetDepartmentOk() (*VacanciesVacancyConditionFieldsRequiredLengthWithTitle, bool) {
	if o == nil {
		return nil, false
	}
	return o.Department.Get(), o.Department.IsSet()
}

// HasDepartment returns a boolean if a field has been set.
func (o *VacanciesVacancyConditions) HasDepartment() bool {
	if o != nil && o.Department.IsSet() {
		return true
	}

	return false
}

// SetDepartment gets a reference to the given NullableVacanciesVacancyConditionFieldsRequiredLengthWithTitle and assigns it to the Department field.
func (o *VacanciesVacancyConditions) SetDepartment(v VacanciesVacancyConditionFieldsRequiredLengthWithTitle) {
	o.Department.Set(&v)
}
// SetDepartmentNil sets the value for Department to be an explicit nil
func (o *VacanciesVacancyConditions) SetDepartmentNil() {
	o.Department.Set(nil)
}

// UnsetDepartment ensures that no value is present for Department, not even an explicit nil
func (o *VacanciesVacancyConditions) UnsetDepartment() {
	o.Department.Unset()
}

// GetDescription returns the Description field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *VacanciesVacancyConditions) GetDescription() VacanciesVacancyConditionFieldsRequiredLengthWithTitle {
	if o == nil || IsNil(o.Description.Get()) {
		var ret VacanciesVacancyConditionFieldsRequiredLengthWithTitle
		return ret
	}
	return *o.Description.Get()
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *VacanciesVacancyConditions) GetDescriptionOk() (*VacanciesVacancyConditionFieldsRequiredLengthWithTitle, bool) {
	if o == nil {
		return nil, false
	}
	return o.Description.Get(), o.Description.IsSet()
}

// HasDescription returns a boolean if a field has been set.
func (o *VacanciesVacancyConditions) HasDescription() bool {
	if o != nil && o.Description.IsSet() {
		return true
	}

	return false
}

// SetDescription gets a reference to the given NullableVacanciesVacancyConditionFieldsRequiredLengthWithTitle and assigns it to the Description field.
func (o *VacanciesVacancyConditions) SetDescription(v VacanciesVacancyConditionFieldsRequiredLengthWithTitle) {
	o.Description.Set(&v)
}
// SetDescriptionNil sets the value for Description to be an explicit nil
func (o *VacanciesVacancyConditions) SetDescriptionNil() {
	o.Description.Set(nil)
}

// UnsetDescription ensures that no value is present for Description, not even an explicit nil
func (o *VacanciesVacancyConditions) UnsetDescription() {
	o.Description.Unset()
}

// GetDriverLicenseTypes returns the DriverLicenseTypes field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *VacanciesVacancyConditions) GetDriverLicenseTypes() VacanciesVacancyConditionFieldsRequiredCountWithTitle {
	if o == nil || IsNil(o.DriverLicenseTypes.Get()) {
		var ret VacanciesVacancyConditionFieldsRequiredCountWithTitle
		return ret
	}
	return *o.DriverLicenseTypes.Get()
}

// GetDriverLicenseTypesOk returns a tuple with the DriverLicenseTypes field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *VacanciesVacancyConditions) GetDriverLicenseTypesOk() (*VacanciesVacancyConditionFieldsRequiredCountWithTitle, bool) {
	if o == nil {
		return nil, false
	}
	return o.DriverLicenseTypes.Get(), o.DriverLicenseTypes.IsSet()
}

// HasDriverLicenseTypes returns a boolean if a field has been set.
func (o *VacanciesVacancyConditions) HasDriverLicenseTypes() bool {
	if o != nil && o.DriverLicenseTypes.IsSet() {
		return true
	}

	return false
}

// SetDriverLicenseTypes gets a reference to the given NullableVacanciesVacancyConditionFieldsRequiredCountWithTitle and assigns it to the DriverLicenseTypes field.
func (o *VacanciesVacancyConditions) SetDriverLicenseTypes(v VacanciesVacancyConditionFieldsRequiredCountWithTitle) {
	o.DriverLicenseTypes.Set(&v)
}
// SetDriverLicenseTypesNil sets the value for DriverLicenseTypes to be an explicit nil
func (o *VacanciesVacancyConditions) SetDriverLicenseTypesNil() {
	o.DriverLicenseTypes.Set(nil)
}

// UnsetDriverLicenseTypes ensures that no value is present for DriverLicenseTypes, not even an explicit nil
func (o *VacanciesVacancyConditions) UnsetDriverLicenseTypes() {
	o.DriverLicenseTypes.Unset()
}

// GetEmployment returns the Employment field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *VacanciesVacancyConditions) GetEmployment() VacanciesVacancyConditionFieldsRequiredWithTitle {
	if o == nil || IsNil(o.Employment.Get()) {
		var ret VacanciesVacancyConditionFieldsRequiredWithTitle
		return ret
	}
	return *o.Employment.Get()
}

// GetEmploymentOk returns a tuple with the Employment field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *VacanciesVacancyConditions) GetEmploymentOk() (*VacanciesVacancyConditionFieldsRequiredWithTitle, bool) {
	if o == nil {
		return nil, false
	}
	return o.Employment.Get(), o.Employment.IsSet()
}

// HasEmployment returns a boolean if a field has been set.
func (o *VacanciesVacancyConditions) HasEmployment() bool {
	if o != nil && o.Employment.IsSet() {
		return true
	}

	return false
}

// SetEmployment gets a reference to the given NullableVacanciesVacancyConditionFieldsRequiredWithTitle and assigns it to the Employment field.
func (o *VacanciesVacancyConditions) SetEmployment(v VacanciesVacancyConditionFieldsRequiredWithTitle) {
	o.Employment.Set(&v)
}
// SetEmploymentNil sets the value for Employment to be an explicit nil
func (o *VacanciesVacancyConditions) SetEmploymentNil() {
	o.Employment.Set(nil)
}

// UnsetEmployment ensures that no value is present for Employment, not even an explicit nil
func (o *VacanciesVacancyConditions) UnsetEmployment() {
	o.Employment.Unset()
}

// GetExperience returns the Experience field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *VacanciesVacancyConditions) GetExperience() VacanciesVacancyConditionFieldsRequiredWithTitle {
	if o == nil || IsNil(o.Experience.Get()) {
		var ret VacanciesVacancyConditionFieldsRequiredWithTitle
		return ret
	}
	return *o.Experience.Get()
}

// GetExperienceOk returns a tuple with the Experience field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *VacanciesVacancyConditions) GetExperienceOk() (*VacanciesVacancyConditionFieldsRequiredWithTitle, bool) {
	if o == nil {
		return nil, false
	}
	return o.Experience.Get(), o.Experience.IsSet()
}

// HasExperience returns a boolean if a field has been set.
func (o *VacanciesVacancyConditions) HasExperience() bool {
	if o != nil && o.Experience.IsSet() {
		return true
	}

	return false
}

// SetExperience gets a reference to the given NullableVacanciesVacancyConditionFieldsRequiredWithTitle and assigns it to the Experience field.
func (o *VacanciesVacancyConditions) SetExperience(v VacanciesVacancyConditionFieldsRequiredWithTitle) {
	o.Experience.Set(&v)
}
// SetExperienceNil sets the value for Experience to be an explicit nil
func (o *VacanciesVacancyConditions) SetExperienceNil() {
	o.Experience.Set(nil)
}

// UnsetExperience ensures that no value is present for Experience, not even an explicit nil
func (o *VacanciesVacancyConditions) UnsetExperience() {
	o.Experience.Unset()
}

// GetKeySkills returns the KeySkills field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *VacanciesVacancyConditions) GetKeySkills() VacanciesVacancyConditionFieldsRequiredCountWithTitle {
	if o == nil || IsNil(o.KeySkills.Get()) {
		var ret VacanciesVacancyConditionFieldsRequiredCountWithTitle
		return ret
	}
	return *o.KeySkills.Get()
}

// GetKeySkillsOk returns a tuple with the KeySkills field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *VacanciesVacancyConditions) GetKeySkillsOk() (*VacanciesVacancyConditionFieldsRequiredCountWithTitle, bool) {
	if o == nil {
		return nil, false
	}
	return o.KeySkills.Get(), o.KeySkills.IsSet()
}

// HasKeySkills returns a boolean if a field has been set.
func (o *VacanciesVacancyConditions) HasKeySkills() bool {
	if o != nil && o.KeySkills.IsSet() {
		return true
	}

	return false
}

// SetKeySkills gets a reference to the given NullableVacanciesVacancyConditionFieldsRequiredCountWithTitle and assigns it to the KeySkills field.
func (o *VacanciesVacancyConditions) SetKeySkills(v VacanciesVacancyConditionFieldsRequiredCountWithTitle) {
	o.KeySkills.Set(&v)
}
// SetKeySkillsNil sets the value for KeySkills to be an explicit nil
func (o *VacanciesVacancyConditions) SetKeySkillsNil() {
	o.KeySkills.Set(nil)
}

// UnsetKeySkills ensures that no value is present for KeySkills, not even an explicit nil
func (o *VacanciesVacancyConditions) UnsetKeySkills() {
	o.KeySkills.Unset()
}

// GetLanguages returns the Languages field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *VacanciesVacancyConditions) GetLanguages() VacanciesVacancyConditionFieldsRequiredCountWithTitle {
	if o == nil || IsNil(o.Languages.Get()) {
		var ret VacanciesVacancyConditionFieldsRequiredCountWithTitle
		return ret
	}
	return *o.Languages.Get()
}

// GetLanguagesOk returns a tuple with the Languages field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *VacanciesVacancyConditions) GetLanguagesOk() (*VacanciesVacancyConditionFieldsRequiredCountWithTitle, bool) {
	if o == nil {
		return nil, false
	}
	return o.Languages.Get(), o.Languages.IsSet()
}

// HasLanguages returns a boolean if a field has been set.
func (o *VacanciesVacancyConditions) HasLanguages() bool {
	if o != nil && o.Languages.IsSet() {
		return true
	}

	return false
}

// SetLanguages gets a reference to the given NullableVacanciesVacancyConditionFieldsRequiredCountWithTitle and assigns it to the Languages field.
func (o *VacanciesVacancyConditions) SetLanguages(v VacanciesVacancyConditionFieldsRequiredCountWithTitle) {
	o.Languages.Set(&v)
}
// SetLanguagesNil sets the value for Languages to be an explicit nil
func (o *VacanciesVacancyConditions) SetLanguagesNil() {
	o.Languages.Set(nil)
}

// UnsetLanguages ensures that no value is present for Languages, not even an explicit nil
func (o *VacanciesVacancyConditions) UnsetLanguages() {
	o.Languages.Unset()
}

// GetManager returns the Manager field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *VacanciesVacancyConditions) GetManager() VacanciesVacancyConditionFieldsRequiredWithTitle {
	if o == nil || IsNil(o.Manager.Get()) {
		var ret VacanciesVacancyConditionFieldsRequiredWithTitle
		return ret
	}
	return *o.Manager.Get()
}

// GetManagerOk returns a tuple with the Manager field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *VacanciesVacancyConditions) GetManagerOk() (*VacanciesVacancyConditionFieldsRequiredWithTitle, bool) {
	if o == nil {
		return nil, false
	}
	return o.Manager.Get(), o.Manager.IsSet()
}

// HasManager returns a boolean if a field has been set.
func (o *VacanciesVacancyConditions) HasManager() bool {
	if o != nil && o.Manager.IsSet() {
		return true
	}

	return false
}

// SetManager gets a reference to the given NullableVacanciesVacancyConditionFieldsRequiredWithTitle and assigns it to the Manager field.
func (o *VacanciesVacancyConditions) SetManager(v VacanciesVacancyConditionFieldsRequiredWithTitle) {
	o.Manager.Set(&v)
}
// SetManagerNil sets the value for Manager to be an explicit nil
func (o *VacanciesVacancyConditions) SetManagerNil() {
	o.Manager.Set(nil)
}

// UnsetManager ensures that no value is present for Manager, not even an explicit nil
func (o *VacanciesVacancyConditions) UnsetManager() {
	o.Manager.Unset()
}

// GetName returns the Name field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *VacanciesVacancyConditions) GetName() VacanciesVacancyConditionFieldsRequiredLengthWithTitle {
	if o == nil || IsNil(o.Name.Get()) {
		var ret VacanciesVacancyConditionFieldsRequiredLengthWithTitle
		return ret
	}
	return *o.Name.Get()
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *VacanciesVacancyConditions) GetNameOk() (*VacanciesVacancyConditionFieldsRequiredLengthWithTitle, bool) {
	if o == nil {
		return nil, false
	}
	return o.Name.Get(), o.Name.IsSet()
}

// HasName returns a boolean if a field has been set.
func (o *VacanciesVacancyConditions) HasName() bool {
	if o != nil && o.Name.IsSet() {
		return true
	}

	return false
}

// SetName gets a reference to the given NullableVacanciesVacancyConditionFieldsRequiredLengthWithTitle and assigns it to the Name field.
func (o *VacanciesVacancyConditions) SetName(v VacanciesVacancyConditionFieldsRequiredLengthWithTitle) {
	o.Name.Set(&v)
}
// SetNameNil sets the value for Name to be an explicit nil
func (o *VacanciesVacancyConditions) SetNameNil() {
	o.Name.Set(nil)
}

// UnsetName ensures that no value is present for Name, not even an explicit nil
func (o *VacanciesVacancyConditions) UnsetName() {
	o.Name.Unset()
}

// GetProfessionalRoles returns the ProfessionalRoles field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *VacanciesVacancyConditions) GetProfessionalRoles() VacanciesVacancyConditionFieldsRequiredCountWithTitle {
	if o == nil || IsNil(o.ProfessionalRoles.Get()) {
		var ret VacanciesVacancyConditionFieldsRequiredCountWithTitle
		return ret
	}
	return *o.ProfessionalRoles.Get()
}

// GetProfessionalRolesOk returns a tuple with the ProfessionalRoles field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *VacanciesVacancyConditions) GetProfessionalRolesOk() (*VacanciesVacancyConditionFieldsRequiredCountWithTitle, bool) {
	if o == nil {
		return nil, false
	}
	return o.ProfessionalRoles.Get(), o.ProfessionalRoles.IsSet()
}

// HasProfessionalRoles returns a boolean if a field has been set.
func (o *VacanciesVacancyConditions) HasProfessionalRoles() bool {
	if o != nil && o.ProfessionalRoles.IsSet() {
		return true
	}

	return false
}

// SetProfessionalRoles gets a reference to the given NullableVacanciesVacancyConditionFieldsRequiredCountWithTitle and assigns it to the ProfessionalRoles field.
func (o *VacanciesVacancyConditions) SetProfessionalRoles(v VacanciesVacancyConditionFieldsRequiredCountWithTitle) {
	o.ProfessionalRoles.Set(&v)
}
// SetProfessionalRolesNil sets the value for ProfessionalRoles to be an explicit nil
func (o *VacanciesVacancyConditions) SetProfessionalRolesNil() {
	o.ProfessionalRoles.Set(nil)
}

// UnsetProfessionalRoles ensures that no value is present for ProfessionalRoles, not even an explicit nil
func (o *VacanciesVacancyConditions) UnsetProfessionalRoles() {
	o.ProfessionalRoles.Unset()
}

// GetResponseLetterRequired returns the ResponseLetterRequired field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *VacanciesVacancyConditions) GetResponseLetterRequired() VacanciesVacancyConditionFieldsRequiredWithTitle {
	if o == nil || IsNil(o.ResponseLetterRequired.Get()) {
		var ret VacanciesVacancyConditionFieldsRequiredWithTitle
		return ret
	}
	return *o.ResponseLetterRequired.Get()
}

// GetResponseLetterRequiredOk returns a tuple with the ResponseLetterRequired field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *VacanciesVacancyConditions) GetResponseLetterRequiredOk() (*VacanciesVacancyConditionFieldsRequiredWithTitle, bool) {
	if o == nil {
		return nil, false
	}
	return o.ResponseLetterRequired.Get(), o.ResponseLetterRequired.IsSet()
}

// HasResponseLetterRequired returns a boolean if a field has been set.
func (o *VacanciesVacancyConditions) HasResponseLetterRequired() bool {
	if o != nil && o.ResponseLetterRequired.IsSet() {
		return true
	}

	return false
}

// SetResponseLetterRequired gets a reference to the given NullableVacanciesVacancyConditionFieldsRequiredWithTitle and assigns it to the ResponseLetterRequired field.
func (o *VacanciesVacancyConditions) SetResponseLetterRequired(v VacanciesVacancyConditionFieldsRequiredWithTitle) {
	o.ResponseLetterRequired.Set(&v)
}
// SetResponseLetterRequiredNil sets the value for ResponseLetterRequired to be an explicit nil
func (o *VacanciesVacancyConditions) SetResponseLetterRequiredNil() {
	o.ResponseLetterRequired.Set(nil)
}

// UnsetResponseLetterRequired ensures that no value is present for ResponseLetterRequired, not even an explicit nil
func (o *VacanciesVacancyConditions) UnsetResponseLetterRequired() {
	o.ResponseLetterRequired.Unset()
}

// GetResponseNotifications returns the ResponseNotifications field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *VacanciesVacancyConditions) GetResponseNotifications() VacanciesVacancyConditionFieldsRequiredWithTitle {
	if o == nil || IsNil(o.ResponseNotifications.Get()) {
		var ret VacanciesVacancyConditionFieldsRequiredWithTitle
		return ret
	}
	return *o.ResponseNotifications.Get()
}

// GetResponseNotificationsOk returns a tuple with the ResponseNotifications field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *VacanciesVacancyConditions) GetResponseNotificationsOk() (*VacanciesVacancyConditionFieldsRequiredWithTitle, bool) {
	if o == nil {
		return nil, false
	}
	return o.ResponseNotifications.Get(), o.ResponseNotifications.IsSet()
}

// HasResponseNotifications returns a boolean if a field has been set.
func (o *VacanciesVacancyConditions) HasResponseNotifications() bool {
	if o != nil && o.ResponseNotifications.IsSet() {
		return true
	}

	return false
}

// SetResponseNotifications gets a reference to the given NullableVacanciesVacancyConditionFieldsRequiredWithTitle and assigns it to the ResponseNotifications field.
func (o *VacanciesVacancyConditions) SetResponseNotifications(v VacanciesVacancyConditionFieldsRequiredWithTitle) {
	o.ResponseNotifications.Set(&v)
}
// SetResponseNotificationsNil sets the value for ResponseNotifications to be an explicit nil
func (o *VacanciesVacancyConditions) SetResponseNotificationsNil() {
	o.ResponseNotifications.Set(nil)
}

// UnsetResponseNotifications ensures that no value is present for ResponseNotifications, not even an explicit nil
func (o *VacanciesVacancyConditions) UnsetResponseNotifications() {
	o.ResponseNotifications.Unset()
}

// GetResponseUrl returns the ResponseUrl field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *VacanciesVacancyConditions) GetResponseUrl() VacanciesVacancyConditionFieldsResponseUrlCondition {
	if o == nil || IsNil(o.ResponseUrl.Get()) {
		var ret VacanciesVacancyConditionFieldsResponseUrlCondition
		return ret
	}
	return *o.ResponseUrl.Get()
}

// GetResponseUrlOk returns a tuple with the ResponseUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *VacanciesVacancyConditions) GetResponseUrlOk() (*VacanciesVacancyConditionFieldsResponseUrlCondition, bool) {
	if o == nil {
		return nil, false
	}
	return o.ResponseUrl.Get(), o.ResponseUrl.IsSet()
}

// HasResponseUrl returns a boolean if a field has been set.
func (o *VacanciesVacancyConditions) HasResponseUrl() bool {
	if o != nil && o.ResponseUrl.IsSet() {
		return true
	}

	return false
}

// SetResponseUrl gets a reference to the given NullableVacanciesVacancyConditionFieldsResponseUrlCondition and assigns it to the ResponseUrl field.
func (o *VacanciesVacancyConditions) SetResponseUrl(v VacanciesVacancyConditionFieldsResponseUrlCondition) {
	o.ResponseUrl.Set(&v)
}
// SetResponseUrlNil sets the value for ResponseUrl to be an explicit nil
func (o *VacanciesVacancyConditions) SetResponseUrlNil() {
	o.ResponseUrl.Set(nil)
}

// UnsetResponseUrl ensures that no value is present for ResponseUrl, not even an explicit nil
func (o *VacanciesVacancyConditions) UnsetResponseUrl() {
	o.ResponseUrl.Unset()
}

// GetSalary returns the Salary field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *VacanciesVacancyConditions) GetSalary() VacanciesVacancyConditionFieldsSalaryCondition {
	if o == nil || IsNil(o.Salary.Get()) {
		var ret VacanciesVacancyConditionFieldsSalaryCondition
		return ret
	}
	return *o.Salary.Get()
}

// GetSalaryOk returns a tuple with the Salary field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *VacanciesVacancyConditions) GetSalaryOk() (*VacanciesVacancyConditionFieldsSalaryCondition, bool) {
	if o == nil {
		return nil, false
	}
	return o.Salary.Get(), o.Salary.IsSet()
}

// HasSalary returns a boolean if a field has been set.
func (o *VacanciesVacancyConditions) HasSalary() bool {
	if o != nil && o.Salary.IsSet() {
		return true
	}

	return false
}

// SetSalary gets a reference to the given NullableVacanciesVacancyConditionFieldsSalaryCondition and assigns it to the Salary field.
func (o *VacanciesVacancyConditions) SetSalary(v VacanciesVacancyConditionFieldsSalaryCondition) {
	o.Salary.Set(&v)
}
// SetSalaryNil sets the value for Salary to be an explicit nil
func (o *VacanciesVacancyConditions) SetSalaryNil() {
	o.Salary.Set(nil)
}

// UnsetSalary ensures that no value is present for Salary, not even an explicit nil
func (o *VacanciesVacancyConditions) UnsetSalary() {
	o.Salary.Unset()
}

// GetSchedule returns the Schedule field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *VacanciesVacancyConditions) GetSchedule() VacanciesVacancyConditionFieldsRequiredWithTitle {
	if o == nil || IsNil(o.Schedule.Get()) {
		var ret VacanciesVacancyConditionFieldsRequiredWithTitle
		return ret
	}
	return *o.Schedule.Get()
}

// GetScheduleOk returns a tuple with the Schedule field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *VacanciesVacancyConditions) GetScheduleOk() (*VacanciesVacancyConditionFieldsRequiredWithTitle, bool) {
	if o == nil {
		return nil, false
	}
	return o.Schedule.Get(), o.Schedule.IsSet()
}

// HasSchedule returns a boolean if a field has been set.
func (o *VacanciesVacancyConditions) HasSchedule() bool {
	if o != nil && o.Schedule.IsSet() {
		return true
	}

	return false
}

// SetSchedule gets a reference to the given NullableVacanciesVacancyConditionFieldsRequiredWithTitle and assigns it to the Schedule field.
func (o *VacanciesVacancyConditions) SetSchedule(v VacanciesVacancyConditionFieldsRequiredWithTitle) {
	o.Schedule.Set(&v)
}
// SetScheduleNil sets the value for Schedule to be an explicit nil
func (o *VacanciesVacancyConditions) SetScheduleNil() {
	o.Schedule.Set(nil)
}

// UnsetSchedule ensures that no value is present for Schedule, not even an explicit nil
func (o *VacanciesVacancyConditions) UnsetSchedule() {
	o.Schedule.Unset()
}

// GetTest returns the Test field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *VacanciesVacancyConditions) GetTest() VacanciesVacancyConditionFieldsTestCondition {
	if o == nil || IsNil(o.Test.Get()) {
		var ret VacanciesVacancyConditionFieldsTestCondition
		return ret
	}
	return *o.Test.Get()
}

// GetTestOk returns a tuple with the Test field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *VacanciesVacancyConditions) GetTestOk() (*VacanciesVacancyConditionFieldsTestCondition, bool) {
	if o == nil {
		return nil, false
	}
	return o.Test.Get(), o.Test.IsSet()
}

// HasTest returns a boolean if a field has been set.
func (o *VacanciesVacancyConditions) HasTest() bool {
	if o != nil && o.Test.IsSet() {
		return true
	}

	return false
}

// SetTest gets a reference to the given NullableVacanciesVacancyConditionFieldsTestCondition and assigns it to the Test field.
func (o *VacanciesVacancyConditions) SetTest(v VacanciesVacancyConditionFieldsTestCondition) {
	o.Test.Set(&v)
}
// SetTestNil sets the value for Test to be an explicit nil
func (o *VacanciesVacancyConditions) SetTestNil() {
	o.Test.Set(nil)
}

// UnsetTest ensures that no value is present for Test, not even an explicit nil
func (o *VacanciesVacancyConditions) UnsetTest() {
	o.Test.Unset()
}

// GetType returns the Type field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *VacanciesVacancyConditions) GetType() VacanciesVacancyConditionFieldsRequiredWithTitle {
	if o == nil || IsNil(o.Type.Get()) {
		var ret VacanciesVacancyConditionFieldsRequiredWithTitle
		return ret
	}
	return *o.Type.Get()
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *VacanciesVacancyConditions) GetTypeOk() (*VacanciesVacancyConditionFieldsRequiredWithTitle, bool) {
	if o == nil {
		return nil, false
	}
	return o.Type.Get(), o.Type.IsSet()
}

// HasType returns a boolean if a field has been set.
func (o *VacanciesVacancyConditions) HasType() bool {
	if o != nil && o.Type.IsSet() {
		return true
	}

	return false
}

// SetType gets a reference to the given NullableVacanciesVacancyConditionFieldsRequiredWithTitle and assigns it to the Type field.
func (o *VacanciesVacancyConditions) SetType(v VacanciesVacancyConditionFieldsRequiredWithTitle) {
	o.Type.Set(&v)
}
// SetTypeNil sets the value for Type to be an explicit nil
func (o *VacanciesVacancyConditions) SetTypeNil() {
	o.Type.Set(nil)
}

// UnsetType ensures that no value is present for Type, not even an explicit nil
func (o *VacanciesVacancyConditions) UnsetType() {
	o.Type.Unset()
}

// GetWorkingDays returns the WorkingDays field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *VacanciesVacancyConditions) GetWorkingDays() VacanciesVacancyConditionFieldsRequiredCountWithTitle {
	if o == nil || IsNil(o.WorkingDays.Get()) {
		var ret VacanciesVacancyConditionFieldsRequiredCountWithTitle
		return ret
	}
	return *o.WorkingDays.Get()
}

// GetWorkingDaysOk returns a tuple with the WorkingDays field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *VacanciesVacancyConditions) GetWorkingDaysOk() (*VacanciesVacancyConditionFieldsRequiredCountWithTitle, bool) {
	if o == nil {
		return nil, false
	}
	return o.WorkingDays.Get(), o.WorkingDays.IsSet()
}

// HasWorkingDays returns a boolean if a field has been set.
func (o *VacanciesVacancyConditions) HasWorkingDays() bool {
	if o != nil && o.WorkingDays.IsSet() {
		return true
	}

	return false
}

// SetWorkingDays gets a reference to the given NullableVacanciesVacancyConditionFieldsRequiredCountWithTitle and assigns it to the WorkingDays field.
func (o *VacanciesVacancyConditions) SetWorkingDays(v VacanciesVacancyConditionFieldsRequiredCountWithTitle) {
	o.WorkingDays.Set(&v)
}
// SetWorkingDaysNil sets the value for WorkingDays to be an explicit nil
func (o *VacanciesVacancyConditions) SetWorkingDaysNil() {
	o.WorkingDays.Set(nil)
}

// UnsetWorkingDays ensures that no value is present for WorkingDays, not even an explicit nil
func (o *VacanciesVacancyConditions) UnsetWorkingDays() {
	o.WorkingDays.Unset()
}

// GetWorkingTimeIntervals returns the WorkingTimeIntervals field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *VacanciesVacancyConditions) GetWorkingTimeIntervals() VacanciesVacancyConditionFieldsRequiredCountWithTitle {
	if o == nil || IsNil(o.WorkingTimeIntervals.Get()) {
		var ret VacanciesVacancyConditionFieldsRequiredCountWithTitle
		return ret
	}
	return *o.WorkingTimeIntervals.Get()
}

// GetWorkingTimeIntervalsOk returns a tuple with the WorkingTimeIntervals field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *VacanciesVacancyConditions) GetWorkingTimeIntervalsOk() (*VacanciesVacancyConditionFieldsRequiredCountWithTitle, bool) {
	if o == nil {
		return nil, false
	}
	return o.WorkingTimeIntervals.Get(), o.WorkingTimeIntervals.IsSet()
}

// HasWorkingTimeIntervals returns a boolean if a field has been set.
func (o *VacanciesVacancyConditions) HasWorkingTimeIntervals() bool {
	if o != nil && o.WorkingTimeIntervals.IsSet() {
		return true
	}

	return false
}

// SetWorkingTimeIntervals gets a reference to the given NullableVacanciesVacancyConditionFieldsRequiredCountWithTitle and assigns it to the WorkingTimeIntervals field.
func (o *VacanciesVacancyConditions) SetWorkingTimeIntervals(v VacanciesVacancyConditionFieldsRequiredCountWithTitle) {
	o.WorkingTimeIntervals.Set(&v)
}
// SetWorkingTimeIntervalsNil sets the value for WorkingTimeIntervals to be an explicit nil
func (o *VacanciesVacancyConditions) SetWorkingTimeIntervalsNil() {
	o.WorkingTimeIntervals.Set(nil)
}

// UnsetWorkingTimeIntervals ensures that no value is present for WorkingTimeIntervals, not even an explicit nil
func (o *VacanciesVacancyConditions) UnsetWorkingTimeIntervals() {
	o.WorkingTimeIntervals.Unset()
}

// GetWorkingTimeModes returns the WorkingTimeModes field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *VacanciesVacancyConditions) GetWorkingTimeModes() VacanciesVacancyConditionFieldsRequiredCountWithTitle {
	if o == nil || IsNil(o.WorkingTimeModes.Get()) {
		var ret VacanciesVacancyConditionFieldsRequiredCountWithTitle
		return ret
	}
	return *o.WorkingTimeModes.Get()
}

// GetWorkingTimeModesOk returns a tuple with the WorkingTimeModes field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *VacanciesVacancyConditions) GetWorkingTimeModesOk() (*VacanciesVacancyConditionFieldsRequiredCountWithTitle, bool) {
	if o == nil {
		return nil, false
	}
	return o.WorkingTimeModes.Get(), o.WorkingTimeModes.IsSet()
}

// HasWorkingTimeModes returns a boolean if a field has been set.
func (o *VacanciesVacancyConditions) HasWorkingTimeModes() bool {
	if o != nil && o.WorkingTimeModes.IsSet() {
		return true
	}

	return false
}

// SetWorkingTimeModes gets a reference to the given NullableVacanciesVacancyConditionFieldsRequiredCountWithTitle and assigns it to the WorkingTimeModes field.
func (o *VacanciesVacancyConditions) SetWorkingTimeModes(v VacanciesVacancyConditionFieldsRequiredCountWithTitle) {
	o.WorkingTimeModes.Set(&v)
}
// SetWorkingTimeModesNil sets the value for WorkingTimeModes to be an explicit nil
func (o *VacanciesVacancyConditions) SetWorkingTimeModesNil() {
	o.WorkingTimeModes.Set(nil)
}

// UnsetWorkingTimeModes ensures that no value is present for WorkingTimeModes, not even an explicit nil
func (o *VacanciesVacancyConditions) UnsetWorkingTimeModes() {
	o.WorkingTimeModes.Unset()
}

func (o VacanciesVacancyConditions) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o VacanciesVacancyConditions) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.AcceptHandicapped.IsSet() {
		toSerialize["accept_handicapped"] = o.AcceptHandicapped.Get()
	}
	if o.AcceptIncompleteResumes.IsSet() {
		toSerialize["accept_incomplete_resumes"] = o.AcceptIncompleteResumes.Get()
	}
	if o.AcceptKids.IsSet() {
		toSerialize["accept_kids"] = o.AcceptKids.Get()
	}
	if o.AcceptTemporary.IsSet() {
		toSerialize["accept_temporary"] = o.AcceptTemporary.Get()
	}
	if o.Address.IsSet() {
		toSerialize["address"] = o.Address.Get()
	}
	if o.AllowMessages.IsSet() {
		toSerialize["allow_messages"] = o.AllowMessages.Get()
	}
	if o.Area.IsSet() {
		toSerialize["area"] = o.Area.Get()
	}
	if o.BillingType.IsSet() {
		toSerialize["billing_type"] = o.BillingType.Get()
	}
	if o.BrandedTemplate.IsSet() {
		toSerialize["branded_template"] = o.BrandedTemplate.Get()
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
	if o.DriverLicenseTypes.IsSet() {
		toSerialize["driver_license_types"] = o.DriverLicenseTypes.Get()
	}
	if o.Employment.IsSet() {
		toSerialize["employment"] = o.Employment.Get()
	}
	if o.Experience.IsSet() {
		toSerialize["experience"] = o.Experience.Get()
	}
	if o.KeySkills.IsSet() {
		toSerialize["key_skills"] = o.KeySkills.Get()
	}
	if o.Languages.IsSet() {
		toSerialize["languages"] = o.Languages.Get()
	}
	if o.Manager.IsSet() {
		toSerialize["manager"] = o.Manager.Get()
	}
	if o.Name.IsSet() {
		toSerialize["name"] = o.Name.Get()
	}
	if o.ProfessionalRoles.IsSet() {
		toSerialize["professional_roles"] = o.ProfessionalRoles.Get()
	}
	if o.ResponseLetterRequired.IsSet() {
		toSerialize["response_letter_required"] = o.ResponseLetterRequired.Get()
	}
	if o.ResponseNotifications.IsSet() {
		toSerialize["response_notifications"] = o.ResponseNotifications.Get()
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
	if o.WorkingDays.IsSet() {
		toSerialize["working_days"] = o.WorkingDays.Get()
	}
	if o.WorkingTimeIntervals.IsSet() {
		toSerialize["working_time_intervals"] = o.WorkingTimeIntervals.Get()
	}
	if o.WorkingTimeModes.IsSet() {
		toSerialize["working_time_modes"] = o.WorkingTimeModes.Get()
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *VacanciesVacancyConditions) UnmarshalJSON(data []byte) (err error) {
	varVacanciesVacancyConditions := _VacanciesVacancyConditions{}

	err = json.Unmarshal(data, &varVacanciesVacancyConditions)

	if err != nil {
		return err
	}

	*o = VacanciesVacancyConditions(varVacanciesVacancyConditions)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "accept_handicapped")
		delete(additionalProperties, "accept_incomplete_resumes")
		delete(additionalProperties, "accept_kids")
		delete(additionalProperties, "accept_temporary")
		delete(additionalProperties, "address")
		delete(additionalProperties, "allow_messages")
		delete(additionalProperties, "area")
		delete(additionalProperties, "billing_type")
		delete(additionalProperties, "branded_template")
		delete(additionalProperties, "code")
		delete(additionalProperties, "contacts")
		delete(additionalProperties, "custom_employer_name")
		delete(additionalProperties, "department")
		delete(additionalProperties, "description")
		delete(additionalProperties, "driver_license_types")
		delete(additionalProperties, "employment")
		delete(additionalProperties, "experience")
		delete(additionalProperties, "key_skills")
		delete(additionalProperties, "languages")
		delete(additionalProperties, "manager")
		delete(additionalProperties, "name")
		delete(additionalProperties, "professional_roles")
		delete(additionalProperties, "response_letter_required")
		delete(additionalProperties, "response_notifications")
		delete(additionalProperties, "response_url")
		delete(additionalProperties, "salary")
		delete(additionalProperties, "schedule")
		delete(additionalProperties, "test")
		delete(additionalProperties, "type")
		delete(additionalProperties, "working_days")
		delete(additionalProperties, "working_time_intervals")
		delete(additionalProperties, "working_time_modes")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableVacanciesVacancyConditions struct {
	value *VacanciesVacancyConditions
	isSet bool
}

func (v NullableVacanciesVacancyConditions) Get() *VacanciesVacancyConditions {
	return v.value
}

func (v *NullableVacanciesVacancyConditions) Set(val *VacanciesVacancyConditions) {
	v.value = val
	v.isSet = true
}

func (v NullableVacanciesVacancyConditions) IsSet() bool {
	return v.isSet
}

func (v *NullableVacanciesVacancyConditions) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableVacanciesVacancyConditions(val *VacanciesVacancyConditions) *NullableVacanciesVacancyConditions {
	return &NullableVacanciesVacancyConditions{value: val, isSet: true}
}

func (v NullableVacanciesVacancyConditions) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableVacanciesVacancyConditions) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


