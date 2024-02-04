/*
HeadHunter API

По-русски | [Switch to English](https://api.hh.ru/openapi/en/redoc)  В OpenAPI ведется пока что только небольшая часть документации [Основная документация](https://github.com/hhru/api).  Для поиска по документации можно использовать Ctrl+F.  # Общая информация  * Всё API работает по протоколу HTTPS. * Авторизация осуществляется по протоколу OAuth2. * Все данные доступны только в формате JSON. * Базовый URL — `https://api.hh.ru/` * Возможны запросы к данным [любого сайта группы компаний HeadHunter](#section/Obshaya-informaciya/Vybor-sajta) * <a name=\"date-format\"></a> Даты форматируются в соответствии с [ISO 8601](http://en.wikipedia.org/wiki/ISO_8601): `YYYY-MM-DDThh:mm:ss±hhmm`.   <a name=\"request-requirements\"></a> ## Требования к запросам  В запросе необходимо передавать заголовок `User-Agent`, но если ваша реализация http клиента не позволяет, можно отправить `HH-User-Agent`. Если не отправлен ни один заголовок, то ответом будет `400 Bad Request`. Указание в заголовке названия приложения и контактной почты разработчика позволит нам оперативно с вами связаться в случае необходимости. Заголовки `User-Agent` и `HH-User-Agent` взаимозаменяемы, в случае, если вы отправите оба заголовка, обработан будет только `HH-User-Agent`.  ``` User-Agent: MyApp/1.0 (my-app-feedback@example.com) ```  Подробнее про [ошибки в заголовке User-Agent](https://github.com/hhru/api/blob/master/docs/errors.md#user-agent).   <a name=\"request-body\"></a> ## Формат тела запроса при отправке JSON  Данные, передающиеся в теле запроса, должны удовлетворять требованиям:  * Валидный JSON (допускается передача как минифицированного варианта, так и pretty print варианта с дополнительными пробелами и сбросами строк). * Рекомендуется использование кодировки UTF-8 без дополнительного экранирования (`{\"name\": \"Иванов Иван\"}`). * Также возможно использовать ascii кодировку с экранированием (`{\"name\": \"\\u0418\\u0432\\u0430\\u043d\\u043e\\u0432 \\u0418\\u0432\\u0430\\u043d\"}`). * К типам данных в определённым полях накладываются дополнительные условия, описанные в каждом конкретном методе. В JSON типами данных являются `string`, `number`, `boolean`, `null`, `object`, `array`.  ### Ответ Ответ свыше определенной длины будет сжиматься методом gzip.  ### Ошибки и коды ответов  API широко использует информирование при помощи кодов ответов. Приложение должно корректно их обрабатывать.  В случае неполадок и сбоев, возможны ответы с кодом `503` и `500`.  При каждой ошибке, помимо кода ответа, в теле ответа может быть выдана дополнительная информация, позволяющая разработчику понять причину соответствующего ответа.  [Более подробно про возможные ошибки](https://github.com/hhru/api/blob/master/docs/errors.md).   ## Недокументированные поля и параметры запросов  В ответах и параметрах API можно найти ключи, не описанные в документации. Обычно это означает, что они оставлены для совместимости со старыми версиями. Их использование не рекомендуется. Если ваше приложение использует такие ключи, перейдите на использование актуальных ключей, описанных в документации.   ## Пагинация  К любому запросу, подразумевающему выдачу списка объектов, можно в параметрах указать `page=N&per_page=M`. Нумерация идёт с нуля, по умолчанию выдаётся первая (нулевая) страница с 20 объектами на странице. Во всех ответах, где доступна пагинация, единообразный корневой объект:  ```json {   \"found\": 1,   \"per_page\": 1,   \"pages\": 1,   \"page\": 0,   \"items\": [{}] } ``` ## Выбор сайта  API HeadHunter позволяет получать данные со всех сайтов группы компании HeadHunter.  В частности:  * hh.ru * rabota.by * hh1.az * hh.uz * hh.kz * headhunter.ge * headhunter.kg  Запросы к данным на всех сайтах следует направлять на `https://api.hh.ru/`.  При необходимости учесть специфику сайта, можно добавить в запрос параметр `?host=`. По умолчанию используется `hh.ru`.  Например, для получения [локализаций](https://api.hh.ru/openapi/redoc#tag/Obshie-spravochniki/operation/get-locales), доступных на hh.kz необходимо сделать GET запрос на `https://api.hh.ru/locales?host=hh.kz`.  ## CORS (Cross-Origin Resource Sharing)  API поддерживает технологию CORS для запроса данных из браузера с произвольного домена. Этот метод более предпочтителен, чем использование JSONP. Он не ограничен методом GET. Для отладки CORS доступен [специальный метод](https://github.com/hhru/api/blob/master/docs/cors.md). Для использования JSONP передайте параметр `?callback=callback_name`.  * [CORS specification on w3.org](http://www.w3.org/TR/cors/) * [HTML5Rocks CORS Tutorial](http://www.html5rocks.com/en/tutorials/cors/) * [CORS on dev.opera.com](http://dev.opera.com/articles/view/dom-access-control-using-cross-origin-resource-sharing/) * [CORS on caniuse.com](http://caniuse.com/#feat=cors) * [CORS on en.wikipedia.org](http://en.wikipedia.org/wiki/Cross-origin_resource_sharing)   ## Внешние ссылки на статьи и стандарты  * [HTTP/1.1](http://tools.ietf.org/html/rfc2616) * [JSON](http://json.org/) * [URI Template](http://tools.ietf.org/html/rfc6570) * [OAuth 2.0](http://tools.ietf.org/html/rfc6749) * [REST](http://www.ics.uci.edu/~fielding/pubs/dissertation/rest_arch_style.htm) * [ISO 8601](http://en.wikipedia.org/wiki/ISO_8601)  # Авторизация  API поддерживает следующие уровни авторизации:   - [авторизация приложения](#section/Avtorizaciya/Avtorizaciya-prilozheniya)   - [авторизация пользователя](#section/Avtorizaciya/Avtorizaciya-polzovatelya)  * [Авторизация пользователя](#section/Avtorizaciya/Avtorizaciya-polzovatelya)   * [Правила формирования специального redirect_uri](#section/Avtorizaciya/Pravila-formirovaniya-specialnogo-redirect_uri)   * [Процесс авторизации](#section/Avtorizaciya/Process-avtorizacii)   * [Успешное получение временного `authorization_code`](#get-authorization_code)   * [Получение access и refresh токенов](#section/Avtorizaciya/Poluchenie-access-i-refresh-tokenov) * [Обновление пары access и refresh токенов](#section/Avtorizaciya/Obnovlenie-pary-access-i-refresh-tokenov) * [Инвалидация токена](#tag/Avtorizaciya-rabotodatelya/operation/invalidate-token) * [Запрос авторизации под другим пользователем](#section/Avtorizaciya/Zapros-avtorizacii-pod-drugim-polzovatelem) * [Авторизация под разными рабочими аккаунтами](#section/Avtorizaciya/Avtorizaciya-pod-raznymi-rabochimi-akkauntami) * [Авторизация приложения](#section/Avtorizaciya/Avtorizaciya-prilozheniya)       ## Авторизация пользователя Для выполнения запросов от имени пользователя необходимо пользоваться токеном пользователя.  В начале приложению необходимо направить пользователя (открыть страницу) по адресу:  ``` https://hh.ru/oauth/authorize? response_type=code& client_id={client_id}& state={state}& redirect_uri={redirect_uri} ```  Обязательные параметры:  * `response_type=code` — указание на способ получения авторизации, используя `authorization code` * `client_id` — идентификатор, полученный при создании приложения   Необязательные параметры:  * `state` — в случае указания, будет включен в ответный редирект. Это позволяет исключить возможность взлома путём подделки межсайтовых запросов. Подробнее об этом: [RFC 6749. Section 10.12](http://tools.ietf.org/html/rfc6749#section-10.12) * `redirect_uri` — uri для перенаправления пользователя после авторизации. Если не указать, используется из настроек приложения. При наличии происходит валидация значения. Вероятнее всего, потребуется сделать urlencode значения параметра.  ## Правила формирования специального redirect_uri  К примеру, если в настройках сохранен `http://example.com/oauth`, то разрешено указывать:  * `http://www.example.com/oauth` — поддомен; * `http://www.example.com/oauth/sub/path` — уточнение пути; * `http://example.com/oauth?lang=RU` — дополнительный параметр; * `http://www.example.com/oauth/sub/path?lang=RU` — всё вместе.  Запрещено:  * `https://example.com/oauth` — различные протоколы; * `http://wwwexample.com/oauth` — различные домены; * `http://wwwexample.com/` — другой путь; * `http://example.com/oauths` — другой путь; * `http://example.com:80/oauths` — указание изначально отсутствующего порта;  ## Процесс авторизации  Если пользователь не авторизован на сайте, ему будет показана форма авторизации на сайте. После прохождения авторизации на сайте, пользователю будет выведена форма с запросом разрешения доступа вашего приложения к его персональным данным.  Если пользователь не разрешает доступ приложению, пользователь будет перенаправлен на указанный `redirect_uri` с `?error=access_denied` и `state={state}`, если таковой был указан при первом запросе.  <a name=\"get-authorization_code\"></a> ### Успешное получение временного `authorization_code`  В случае разрешения прав, в редиректе будет указан временный `authorization_code`:  ```http HTTP/1.1 302 FOUND Location: {redirect_uri}?code={authorization_code} ```  Если пользователь авторизован на сайте и доступ данному приложению однажды ранее выдан, ответом будет сразу вышеописанный редирект с `authorization_code` (без показа формы логина и выдачи прав).  ## Получение access и refresh токенов  После получения `authorization_code` приложению необходимо сделать сервер-сервер запрос `POST https://api.hh.ru/token` для обмена полученного `authorization_code` на `access_token` (старый запрос `POST https://hh.ru/oauth/token` считается устаревшим).  В теле запроса необходимо передать [дополнительные параметры](#required_parameters).  Тело запроса необходимо передавать в стандартном `application/x-www-form-urlencoded` с указанием соответствующего заголовка `Content-Type`.  `authorization_code` имеет довольно короткий срок жизни, при его истечении необходимо запросить новый.  ## Обновление пары access и refresh токенов `access_token` также имеет срок жизни (ключ `expires_in`, в секундах), при его истечении приложение должно сделать запрос с `refresh_token` для получения нового.  Запрос необходимо делать в `application/x-www-form-urlencoded`.  ``` POST https://api.hh.ru/token ```  (старый запрос `POST https://hh.ru/oauth/token` считается устаревшим)  В теле запроса необходимо передать [дополнительные параметры](#required_parameters)  `refresh_token` можно использовать только один раз и только по истечению срока действия `access_token`.  После получения новой пары access и refresh токенов, их необходимо использовать в дальнейших запросах в api и запросах на продление токена.  ## Запрос авторизации под другим пользователем  Возможен следующий сценарий:  1. Приложение перенаправляет пользователя на сайт с запросом авторизации. 2. Пользователь на сайте уже авторизован и данному приложение доступ уже был разрешен. 3. Пользователю будет предложена возможность продолжить работу под текущим аккаунтом, либо зайти под другим аккаунтом.  Если есть необходимость, чтобы на шаге 3 сразу происходило перенаправление (redirect) с временным токеном, необходимо добавить к запросу `/oauth/authorize...` параметр `skip_choose_account=true`. В этом случае автоматически выдаётся доступ пользователю авторизованному на сайте.  Если есть необходимость всегда показывать форму авторизации, приложение может добавить к запросу `/oauth/authorize...` параметр `force_login=true`. В этом случае, пользователю будет показана форма авторизации с логином и паролем даже в случае, если пользователь уже авторизован.  Это может быть полезно приложениям, которые предоставляют сервис только для соискателей. Если пришел пользователь-работодатель, приложение может предложить пользователю повторно разрешить доступ на сайте, уже указав другую учетную запись.  Также, после авторизации приложение может показать пользователю сообщение:  ``` Вы вошли как %Имя_Фамилия%. Это не вы? ``` и предоставить ссылку с `force_login=true` для возможности захода под другим логином.  ## Авторизация под разными рабочими аккаунтами  Для получения списка рабочих аккаунтов менеджера и для работы под разными рабочими аккаунтами менеджера необходимо прочитать документацию по [рабочим аккаунтам менеджера](#tag/Menedzhery-rabotodatelya/operation/get-manager-accounts)  ## Авторизация приложения  Токен приложения необходимо сгенерировать 1 раз. В случае, если токен был скомпрометирован, его нужно запросить еще раз. При этом ранее выданный токен отзывается. Владелец приложения может посмотреть актуальный `access_token` для приложения на сайте [https://dev.hh.ru/admin](https://dev.hh.ru/admin). В случае, если вы еще ни разу [не получали токен приложения](#section/Avtorizaciya/Avtorizaciya-prilozheniya), токен отображаться не будет.  <a name=\"get-client-token\"></a> ### Получение токена приложения Для получения `access_token` необходимо сделать запрос:  ``` POST https://api.hh.ru/token ```  (старый запрос `POST https://hh.ru/oauth/token` считается устаревшим)  В теле запроса необходимо передать [дополнительные параметры](#required_parameters). Тело запроса необходимо передавать в стандартном `application/x-www-form-urlencoded` с указанием соответствующего заголовка `Content-Type`.  Данный `access_token` имеет **неограниченный** срок жизни. При повторном запросе ранее выданный токен отзывается и выдается новый. Запрашивать `access_token` можно не чаще, чем один раз в 5 минут.  В случае компрометации токена необходимо [инвалидировать](#tag/Avtorizaciya-rabotodatelya/operation/invalidate-token) скомпроментированный токен и запросить токен заново!  <!-- ReDoc-Inject: <security-definitions> --> 

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

// checks if the VacancyDraftVacancyDraftFull type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &VacancyDraftVacancyDraftFull{}

// VacancyDraftVacancyDraftFull struct for VacancyDraftVacancyDraftFull
type VacancyDraftVacancyDraftFull struct {
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
	BillingType NullableVacancyBillingTypeOutput `json:"billing_type"`
	// Внутренний код вакансии
	Code NullableString `json:"code,omitempty"`
	Department NullableVacancyDepartmentOutput `json:"department,omitempty"`
	// Описание в html, не менее 200 символов
	Description string `json:"description"`
	// Список требуемых категорий водительских прав
	DriverLicenseTypes []VacancyDriverLicenceTypeItem `json:"driver_license_types"`
	Employment NullableVacancyEmploymentOutput `json:"employment,omitempty"`
	Experience NullableVacancyExperienceOutput `json:"experience"`
	// Информация о наличии прикрепленного тестового задании к вакансии
	HasTest bool `json:"has_test"`
	// Список ключевых навыков, не более 30
	KeySkills []VacancyKeySkillItem `json:"key_skills"`
	// Список языков вакансии. Значения из справочника [/languages](#tag/Obshie-spravochniki/operation/get-dictionaries)
	Languages []VacancyLanguageOutput `json:"languages"`
	Manager VacancyManager `json:"manager"`
	// Название
	Name string `json:"name"`
	// Список профессиональных ролей
	ProfessionalRoles []VacancyProfessionalRoleItemOutput `json:"professional_roles"`
	// Обязательно ли заполнять сообщение при отклике на вакансию
	ResponseLetterRequired bool `json:"response_letter_required"`
	// Уведомлять ли менеджера о новых откликах
	ResponseNotifications bool `json:"response_notifications"`
	// URL отклика для прямых вакансий (`type.id=direct`)
	ResponseUrl NullableString `json:"response_url,omitempty"`
	Salary NullableVacancySalary `json:"salary,omitempty"`
	Schedule VacancyScheduleOutput `json:"schedule"`
	Test NullableVacancyDraftTest `json:"test,omitempty"`
	Type VacancyTypeOutput `json:"type"`
	// Список рабочих дней
	WorkingDays []VacancyWorkingDayItemOutput `json:"working_days,omitempty"`
	// Список с временными интервалами работы
	WorkingTimeIntervals []VacancyWorkingTimeIntervalItemOutput `json:"working_time_intervals,omitempty"`
	// Список режимов времени работы
	WorkingTimeModes []VacancyWorkingTimeModeItemOutput `json:"working_time_modes,omitempty"`
	Address VacancyDraftAddressOutput `json:"address"`
	// Коды и названия регионов (фед. округа, субъекты федерации, города)
	Areas []VacancyAreaOutput `json:"areas"`
	AssignedManager NullableVacancyDraftAssignedManager `json:"assigned_manager,omitempty"`
	BrandedTemplate NullableVacancyBrandedTemplate `json:"branded_template,omitempty"`
	Contacts VacancyDraftContactsWithFullPhone `json:"contacts"`
	// Название компании для анонимных вакансий (`type.id=anonymous`), например \"крупный российский банк\". Поле становится обязательным при публикации вакансии **анонимного** типа
	CustomEmployerName NullableString `json:"custom_employer_name,omitempty"`
	Employer VacancyDraftEmployer `json:"employer"`
	MetaInfo VacancyDraftVacancyDraftBase `json:"meta_info"`
	// Вашу вакансию увидят больше людей. Мы разместим ее дополнительно на сервисе Зарплата.ру
	WithZp bool `json:"with_zp"`
}

type _VacancyDraftVacancyDraftFull VacancyDraftVacancyDraftFull

// NewVacancyDraftVacancyDraftFull instantiates a new VacancyDraftVacancyDraftFull object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewVacancyDraftVacancyDraftFull(acceptHandicapped bool, acceptIncompleteResumes bool, acceptKids bool, allowMessages bool, billingType NullableVacancyBillingTypeOutput, description string, driverLicenseTypes []VacancyDriverLicenceTypeItem, experience NullableVacancyExperienceOutput, hasTest bool, keySkills []VacancyKeySkillItem, languages []VacancyLanguageOutput, manager VacancyManager, name string, professionalRoles []VacancyProfessionalRoleItemOutput, responseLetterRequired bool, responseNotifications bool, schedule VacancyScheduleOutput, type_ VacancyTypeOutput, address VacancyDraftAddressOutput, areas []VacancyAreaOutput, contacts VacancyDraftContactsWithFullPhone, employer VacancyDraftEmployer, metaInfo VacancyDraftVacancyDraftBase, withZp bool) *VacancyDraftVacancyDraftFull {
	this := VacancyDraftVacancyDraftFull{}
	this.AcceptHandicapped = acceptHandicapped
	this.AcceptIncompleteResumes = acceptIncompleteResumes
	this.AcceptKids = acceptKids
	this.AllowMessages = allowMessages
	this.BillingType = billingType
	this.Description = description
	this.DriverLicenseTypes = driverLicenseTypes
	this.Experience = experience
	this.HasTest = hasTest
	this.KeySkills = keySkills
	this.Languages = languages
	this.Manager = manager
	this.Name = name
	this.ProfessionalRoles = professionalRoles
	this.ResponseLetterRequired = responseLetterRequired
	this.ResponseNotifications = responseNotifications
	this.Schedule = schedule
	this.Type = type_
	this.Address = address
	this.Areas = areas
	this.Contacts = contacts
	this.Employer = employer
	this.MetaInfo = metaInfo
	this.WithZp = withZp
	return &this
}

// NewVacancyDraftVacancyDraftFullWithDefaults instantiates a new VacancyDraftVacancyDraftFull object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewVacancyDraftVacancyDraftFullWithDefaults() *VacancyDraftVacancyDraftFull {
	this := VacancyDraftVacancyDraftFull{}
	return &this
}

// GetAcceptHandicapped returns the AcceptHandicapped field value
func (o *VacancyDraftVacancyDraftFull) GetAcceptHandicapped() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.AcceptHandicapped
}

// GetAcceptHandicappedOk returns a tuple with the AcceptHandicapped field value
// and a boolean to check if the value has been set.
func (o *VacancyDraftVacancyDraftFull) GetAcceptHandicappedOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AcceptHandicapped, true
}

// SetAcceptHandicapped sets field value
func (o *VacancyDraftVacancyDraftFull) SetAcceptHandicapped(v bool) {
	o.AcceptHandicapped = v
}

// GetAcceptIncompleteResumes returns the AcceptIncompleteResumes field value
func (o *VacancyDraftVacancyDraftFull) GetAcceptIncompleteResumes() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.AcceptIncompleteResumes
}

// GetAcceptIncompleteResumesOk returns a tuple with the AcceptIncompleteResumes field value
// and a boolean to check if the value has been set.
func (o *VacancyDraftVacancyDraftFull) GetAcceptIncompleteResumesOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AcceptIncompleteResumes, true
}

// SetAcceptIncompleteResumes sets field value
func (o *VacancyDraftVacancyDraftFull) SetAcceptIncompleteResumes(v bool) {
	o.AcceptIncompleteResumes = v
}

// GetAcceptKids returns the AcceptKids field value
func (o *VacancyDraftVacancyDraftFull) GetAcceptKids() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.AcceptKids
}

// GetAcceptKidsOk returns a tuple with the AcceptKids field value
// and a boolean to check if the value has been set.
func (o *VacancyDraftVacancyDraftFull) GetAcceptKidsOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AcceptKids, true
}

// SetAcceptKids sets field value
func (o *VacancyDraftVacancyDraftFull) SetAcceptKids(v bool) {
	o.AcceptKids = v
}

// GetAcceptTemporary returns the AcceptTemporary field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *VacancyDraftVacancyDraftFull) GetAcceptTemporary() bool {
	if o == nil || IsNil(o.AcceptTemporary.Get()) {
		var ret bool
		return ret
	}
	return *o.AcceptTemporary.Get()
}

// GetAcceptTemporaryOk returns a tuple with the AcceptTemporary field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *VacancyDraftVacancyDraftFull) GetAcceptTemporaryOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return o.AcceptTemporary.Get(), o.AcceptTemporary.IsSet()
}

// HasAcceptTemporary returns a boolean if a field has been set.
func (o *VacancyDraftVacancyDraftFull) HasAcceptTemporary() bool {
	if o != nil && o.AcceptTemporary.IsSet() {
		return true
	}

	return false
}

// SetAcceptTemporary gets a reference to the given NullableBool and assigns it to the AcceptTemporary field.
func (o *VacancyDraftVacancyDraftFull) SetAcceptTemporary(v bool) {
	o.AcceptTemporary.Set(&v)
}
// SetAcceptTemporaryNil sets the value for AcceptTemporary to be an explicit nil
func (o *VacancyDraftVacancyDraftFull) SetAcceptTemporaryNil() {
	o.AcceptTemporary.Set(nil)
}

// UnsetAcceptTemporary ensures that no value is present for AcceptTemporary, not even an explicit nil
func (o *VacancyDraftVacancyDraftFull) UnsetAcceptTemporary() {
	o.AcceptTemporary.Unset()
}

// GetAllowMessages returns the AllowMessages field value
func (o *VacancyDraftVacancyDraftFull) GetAllowMessages() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.AllowMessages
}

// GetAllowMessagesOk returns a tuple with the AllowMessages field value
// and a boolean to check if the value has been set.
func (o *VacancyDraftVacancyDraftFull) GetAllowMessagesOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AllowMessages, true
}

// SetAllowMessages sets field value
func (o *VacancyDraftVacancyDraftFull) SetAllowMessages(v bool) {
	o.AllowMessages = v
}

// GetBillingType returns the BillingType field value
// If the value is explicit nil, the zero value for VacancyBillingTypeOutput will be returned
func (o *VacancyDraftVacancyDraftFull) GetBillingType() VacancyBillingTypeOutput {
	if o == nil || o.BillingType.Get() == nil {
		var ret VacancyBillingTypeOutput
		return ret
	}

	return *o.BillingType.Get()
}

// GetBillingTypeOk returns a tuple with the BillingType field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *VacancyDraftVacancyDraftFull) GetBillingTypeOk() (*VacancyBillingTypeOutput, bool) {
	if o == nil {
		return nil, false
	}
	return o.BillingType.Get(), o.BillingType.IsSet()
}

// SetBillingType sets field value
func (o *VacancyDraftVacancyDraftFull) SetBillingType(v VacancyBillingTypeOutput) {
	o.BillingType.Set(&v)
}

// GetCode returns the Code field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *VacancyDraftVacancyDraftFull) GetCode() string {
	if o == nil || IsNil(o.Code.Get()) {
		var ret string
		return ret
	}
	return *o.Code.Get()
}

// GetCodeOk returns a tuple with the Code field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *VacancyDraftVacancyDraftFull) GetCodeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Code.Get(), o.Code.IsSet()
}

// HasCode returns a boolean if a field has been set.
func (o *VacancyDraftVacancyDraftFull) HasCode() bool {
	if o != nil && o.Code.IsSet() {
		return true
	}

	return false
}

// SetCode gets a reference to the given NullableString and assigns it to the Code field.
func (o *VacancyDraftVacancyDraftFull) SetCode(v string) {
	o.Code.Set(&v)
}
// SetCodeNil sets the value for Code to be an explicit nil
func (o *VacancyDraftVacancyDraftFull) SetCodeNil() {
	o.Code.Set(nil)
}

// UnsetCode ensures that no value is present for Code, not even an explicit nil
func (o *VacancyDraftVacancyDraftFull) UnsetCode() {
	o.Code.Unset()
}

// GetDepartment returns the Department field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *VacancyDraftVacancyDraftFull) GetDepartment() VacancyDepartmentOutput {
	if o == nil || IsNil(o.Department.Get()) {
		var ret VacancyDepartmentOutput
		return ret
	}
	return *o.Department.Get()
}

// GetDepartmentOk returns a tuple with the Department field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *VacancyDraftVacancyDraftFull) GetDepartmentOk() (*VacancyDepartmentOutput, bool) {
	if o == nil {
		return nil, false
	}
	return o.Department.Get(), o.Department.IsSet()
}

// HasDepartment returns a boolean if a field has been set.
func (o *VacancyDraftVacancyDraftFull) HasDepartment() bool {
	if o != nil && o.Department.IsSet() {
		return true
	}

	return false
}

// SetDepartment gets a reference to the given NullableVacancyDepartmentOutput and assigns it to the Department field.
func (o *VacancyDraftVacancyDraftFull) SetDepartment(v VacancyDepartmentOutput) {
	o.Department.Set(&v)
}
// SetDepartmentNil sets the value for Department to be an explicit nil
func (o *VacancyDraftVacancyDraftFull) SetDepartmentNil() {
	o.Department.Set(nil)
}

// UnsetDepartment ensures that no value is present for Department, not even an explicit nil
func (o *VacancyDraftVacancyDraftFull) UnsetDepartment() {
	o.Department.Unset()
}

// GetDescription returns the Description field value
func (o *VacancyDraftVacancyDraftFull) GetDescription() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Description
}

// GetDescriptionOk returns a tuple with the Description field value
// and a boolean to check if the value has been set.
func (o *VacancyDraftVacancyDraftFull) GetDescriptionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Description, true
}

// SetDescription sets field value
func (o *VacancyDraftVacancyDraftFull) SetDescription(v string) {
	o.Description = v
}

// GetDriverLicenseTypes returns the DriverLicenseTypes field value
func (o *VacancyDraftVacancyDraftFull) GetDriverLicenseTypes() []VacancyDriverLicenceTypeItem {
	if o == nil {
		var ret []VacancyDriverLicenceTypeItem
		return ret
	}

	return o.DriverLicenseTypes
}

// GetDriverLicenseTypesOk returns a tuple with the DriverLicenseTypes field value
// and a boolean to check if the value has been set.
func (o *VacancyDraftVacancyDraftFull) GetDriverLicenseTypesOk() ([]VacancyDriverLicenceTypeItem, bool) {
	if o == nil {
		return nil, false
	}
	return o.DriverLicenseTypes, true
}

// SetDriverLicenseTypes sets field value
func (o *VacancyDraftVacancyDraftFull) SetDriverLicenseTypes(v []VacancyDriverLicenceTypeItem) {
	o.DriverLicenseTypes = v
}

// GetEmployment returns the Employment field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *VacancyDraftVacancyDraftFull) GetEmployment() VacancyEmploymentOutput {
	if o == nil || IsNil(o.Employment.Get()) {
		var ret VacancyEmploymentOutput
		return ret
	}
	return *o.Employment.Get()
}

// GetEmploymentOk returns a tuple with the Employment field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *VacancyDraftVacancyDraftFull) GetEmploymentOk() (*VacancyEmploymentOutput, bool) {
	if o == nil {
		return nil, false
	}
	return o.Employment.Get(), o.Employment.IsSet()
}

// HasEmployment returns a boolean if a field has been set.
func (o *VacancyDraftVacancyDraftFull) HasEmployment() bool {
	if o != nil && o.Employment.IsSet() {
		return true
	}

	return false
}

// SetEmployment gets a reference to the given NullableVacancyEmploymentOutput and assigns it to the Employment field.
func (o *VacancyDraftVacancyDraftFull) SetEmployment(v VacancyEmploymentOutput) {
	o.Employment.Set(&v)
}
// SetEmploymentNil sets the value for Employment to be an explicit nil
func (o *VacancyDraftVacancyDraftFull) SetEmploymentNil() {
	o.Employment.Set(nil)
}

// UnsetEmployment ensures that no value is present for Employment, not even an explicit nil
func (o *VacancyDraftVacancyDraftFull) UnsetEmployment() {
	o.Employment.Unset()
}

// GetExperience returns the Experience field value
// If the value is explicit nil, the zero value for VacancyExperienceOutput will be returned
func (o *VacancyDraftVacancyDraftFull) GetExperience() VacancyExperienceOutput {
	if o == nil || o.Experience.Get() == nil {
		var ret VacancyExperienceOutput
		return ret
	}

	return *o.Experience.Get()
}

// GetExperienceOk returns a tuple with the Experience field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *VacancyDraftVacancyDraftFull) GetExperienceOk() (*VacancyExperienceOutput, bool) {
	if o == nil {
		return nil, false
	}
	return o.Experience.Get(), o.Experience.IsSet()
}

// SetExperience sets field value
func (o *VacancyDraftVacancyDraftFull) SetExperience(v VacancyExperienceOutput) {
	o.Experience.Set(&v)
}

// GetHasTest returns the HasTest field value
func (o *VacancyDraftVacancyDraftFull) GetHasTest() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.HasTest
}

// GetHasTestOk returns a tuple with the HasTest field value
// and a boolean to check if the value has been set.
func (o *VacancyDraftVacancyDraftFull) GetHasTestOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.HasTest, true
}

// SetHasTest sets field value
func (o *VacancyDraftVacancyDraftFull) SetHasTest(v bool) {
	o.HasTest = v
}

// GetKeySkills returns the KeySkills field value
func (o *VacancyDraftVacancyDraftFull) GetKeySkills() []VacancyKeySkillItem {
	if o == nil {
		var ret []VacancyKeySkillItem
		return ret
	}

	return o.KeySkills
}

// GetKeySkillsOk returns a tuple with the KeySkills field value
// and a boolean to check if the value has been set.
func (o *VacancyDraftVacancyDraftFull) GetKeySkillsOk() ([]VacancyKeySkillItem, bool) {
	if o == nil {
		return nil, false
	}
	return o.KeySkills, true
}

// SetKeySkills sets field value
func (o *VacancyDraftVacancyDraftFull) SetKeySkills(v []VacancyKeySkillItem) {
	o.KeySkills = v
}

// GetLanguages returns the Languages field value
func (o *VacancyDraftVacancyDraftFull) GetLanguages() []VacancyLanguageOutput {
	if o == nil {
		var ret []VacancyLanguageOutput
		return ret
	}

	return o.Languages
}

// GetLanguagesOk returns a tuple with the Languages field value
// and a boolean to check if the value has been set.
func (o *VacancyDraftVacancyDraftFull) GetLanguagesOk() ([]VacancyLanguageOutput, bool) {
	if o == nil {
		return nil, false
	}
	return o.Languages, true
}

// SetLanguages sets field value
func (o *VacancyDraftVacancyDraftFull) SetLanguages(v []VacancyLanguageOutput) {
	o.Languages = v
}

// GetManager returns the Manager field value
func (o *VacancyDraftVacancyDraftFull) GetManager() VacancyManager {
	if o == nil {
		var ret VacancyManager
		return ret
	}

	return o.Manager
}

// GetManagerOk returns a tuple with the Manager field value
// and a boolean to check if the value has been set.
func (o *VacancyDraftVacancyDraftFull) GetManagerOk() (*VacancyManager, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Manager, true
}

// SetManager sets field value
func (o *VacancyDraftVacancyDraftFull) SetManager(v VacancyManager) {
	o.Manager = v
}

// GetName returns the Name field value
func (o *VacancyDraftVacancyDraftFull) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *VacancyDraftVacancyDraftFull) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *VacancyDraftVacancyDraftFull) SetName(v string) {
	o.Name = v
}

// GetProfessionalRoles returns the ProfessionalRoles field value
func (o *VacancyDraftVacancyDraftFull) GetProfessionalRoles() []VacancyProfessionalRoleItemOutput {
	if o == nil {
		var ret []VacancyProfessionalRoleItemOutput
		return ret
	}

	return o.ProfessionalRoles
}

// GetProfessionalRolesOk returns a tuple with the ProfessionalRoles field value
// and a boolean to check if the value has been set.
func (o *VacancyDraftVacancyDraftFull) GetProfessionalRolesOk() ([]VacancyProfessionalRoleItemOutput, bool) {
	if o == nil {
		return nil, false
	}
	return o.ProfessionalRoles, true
}

// SetProfessionalRoles sets field value
func (o *VacancyDraftVacancyDraftFull) SetProfessionalRoles(v []VacancyProfessionalRoleItemOutput) {
	o.ProfessionalRoles = v
}

// GetResponseLetterRequired returns the ResponseLetterRequired field value
func (o *VacancyDraftVacancyDraftFull) GetResponseLetterRequired() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.ResponseLetterRequired
}

// GetResponseLetterRequiredOk returns a tuple with the ResponseLetterRequired field value
// and a boolean to check if the value has been set.
func (o *VacancyDraftVacancyDraftFull) GetResponseLetterRequiredOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ResponseLetterRequired, true
}

// SetResponseLetterRequired sets field value
func (o *VacancyDraftVacancyDraftFull) SetResponseLetterRequired(v bool) {
	o.ResponseLetterRequired = v
}

// GetResponseNotifications returns the ResponseNotifications field value
func (o *VacancyDraftVacancyDraftFull) GetResponseNotifications() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.ResponseNotifications
}

// GetResponseNotificationsOk returns a tuple with the ResponseNotifications field value
// and a boolean to check if the value has been set.
func (o *VacancyDraftVacancyDraftFull) GetResponseNotificationsOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ResponseNotifications, true
}

// SetResponseNotifications sets field value
func (o *VacancyDraftVacancyDraftFull) SetResponseNotifications(v bool) {
	o.ResponseNotifications = v
}

// GetResponseUrl returns the ResponseUrl field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *VacancyDraftVacancyDraftFull) GetResponseUrl() string {
	if o == nil || IsNil(o.ResponseUrl.Get()) {
		var ret string
		return ret
	}
	return *o.ResponseUrl.Get()
}

// GetResponseUrlOk returns a tuple with the ResponseUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *VacancyDraftVacancyDraftFull) GetResponseUrlOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.ResponseUrl.Get(), o.ResponseUrl.IsSet()
}

// HasResponseUrl returns a boolean if a field has been set.
func (o *VacancyDraftVacancyDraftFull) HasResponseUrl() bool {
	if o != nil && o.ResponseUrl.IsSet() {
		return true
	}

	return false
}

// SetResponseUrl gets a reference to the given NullableString and assigns it to the ResponseUrl field.
func (o *VacancyDraftVacancyDraftFull) SetResponseUrl(v string) {
	o.ResponseUrl.Set(&v)
}
// SetResponseUrlNil sets the value for ResponseUrl to be an explicit nil
func (o *VacancyDraftVacancyDraftFull) SetResponseUrlNil() {
	o.ResponseUrl.Set(nil)
}

// UnsetResponseUrl ensures that no value is present for ResponseUrl, not even an explicit nil
func (o *VacancyDraftVacancyDraftFull) UnsetResponseUrl() {
	o.ResponseUrl.Unset()
}

// GetSalary returns the Salary field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *VacancyDraftVacancyDraftFull) GetSalary() VacancySalary {
	if o == nil || IsNil(o.Salary.Get()) {
		var ret VacancySalary
		return ret
	}
	return *o.Salary.Get()
}

// GetSalaryOk returns a tuple with the Salary field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *VacancyDraftVacancyDraftFull) GetSalaryOk() (*VacancySalary, bool) {
	if o == nil {
		return nil, false
	}
	return o.Salary.Get(), o.Salary.IsSet()
}

// HasSalary returns a boolean if a field has been set.
func (o *VacancyDraftVacancyDraftFull) HasSalary() bool {
	if o != nil && o.Salary.IsSet() {
		return true
	}

	return false
}

// SetSalary gets a reference to the given NullableVacancySalary and assigns it to the Salary field.
func (o *VacancyDraftVacancyDraftFull) SetSalary(v VacancySalary) {
	o.Salary.Set(&v)
}
// SetSalaryNil sets the value for Salary to be an explicit nil
func (o *VacancyDraftVacancyDraftFull) SetSalaryNil() {
	o.Salary.Set(nil)
}

// UnsetSalary ensures that no value is present for Salary, not even an explicit nil
func (o *VacancyDraftVacancyDraftFull) UnsetSalary() {
	o.Salary.Unset()
}

// GetSchedule returns the Schedule field value
func (o *VacancyDraftVacancyDraftFull) GetSchedule() VacancyScheduleOutput {
	if o == nil {
		var ret VacancyScheduleOutput
		return ret
	}

	return o.Schedule
}

// GetScheduleOk returns a tuple with the Schedule field value
// and a boolean to check if the value has been set.
func (o *VacancyDraftVacancyDraftFull) GetScheduleOk() (*VacancyScheduleOutput, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Schedule, true
}

// SetSchedule sets field value
func (o *VacancyDraftVacancyDraftFull) SetSchedule(v VacancyScheduleOutput) {
	o.Schedule = v
}

// GetTest returns the Test field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *VacancyDraftVacancyDraftFull) GetTest() VacancyDraftTest {
	if o == nil || IsNil(o.Test.Get()) {
		var ret VacancyDraftTest
		return ret
	}
	return *o.Test.Get()
}

// GetTestOk returns a tuple with the Test field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *VacancyDraftVacancyDraftFull) GetTestOk() (*VacancyDraftTest, bool) {
	if o == nil {
		return nil, false
	}
	return o.Test.Get(), o.Test.IsSet()
}

// HasTest returns a boolean if a field has been set.
func (o *VacancyDraftVacancyDraftFull) HasTest() bool {
	if o != nil && o.Test.IsSet() {
		return true
	}

	return false
}

// SetTest gets a reference to the given NullableVacancyDraftTest and assigns it to the Test field.
func (o *VacancyDraftVacancyDraftFull) SetTest(v VacancyDraftTest) {
	o.Test.Set(&v)
}
// SetTestNil sets the value for Test to be an explicit nil
func (o *VacancyDraftVacancyDraftFull) SetTestNil() {
	o.Test.Set(nil)
}

// UnsetTest ensures that no value is present for Test, not even an explicit nil
func (o *VacancyDraftVacancyDraftFull) UnsetTest() {
	o.Test.Unset()
}

// GetType returns the Type field value
func (o *VacancyDraftVacancyDraftFull) GetType() VacancyTypeOutput {
	if o == nil {
		var ret VacancyTypeOutput
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *VacancyDraftVacancyDraftFull) GetTypeOk() (*VacancyTypeOutput, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *VacancyDraftVacancyDraftFull) SetType(v VacancyTypeOutput) {
	o.Type = v
}

// GetWorkingDays returns the WorkingDays field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *VacancyDraftVacancyDraftFull) GetWorkingDays() []VacancyWorkingDayItemOutput {
	if o == nil {
		var ret []VacancyWorkingDayItemOutput
		return ret
	}
	return o.WorkingDays
}

// GetWorkingDaysOk returns a tuple with the WorkingDays field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *VacancyDraftVacancyDraftFull) GetWorkingDaysOk() ([]VacancyWorkingDayItemOutput, bool) {
	if o == nil || IsNil(o.WorkingDays) {
		return nil, false
	}
	return o.WorkingDays, true
}

// HasWorkingDays returns a boolean if a field has been set.
func (o *VacancyDraftVacancyDraftFull) HasWorkingDays() bool {
	if o != nil && IsNil(o.WorkingDays) {
		return true
	}

	return false
}

// SetWorkingDays gets a reference to the given []VacancyWorkingDayItemOutput and assigns it to the WorkingDays field.
func (o *VacancyDraftVacancyDraftFull) SetWorkingDays(v []VacancyWorkingDayItemOutput) {
	o.WorkingDays = v
}

// GetWorkingTimeIntervals returns the WorkingTimeIntervals field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *VacancyDraftVacancyDraftFull) GetWorkingTimeIntervals() []VacancyWorkingTimeIntervalItemOutput {
	if o == nil {
		var ret []VacancyWorkingTimeIntervalItemOutput
		return ret
	}
	return o.WorkingTimeIntervals
}

// GetWorkingTimeIntervalsOk returns a tuple with the WorkingTimeIntervals field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *VacancyDraftVacancyDraftFull) GetWorkingTimeIntervalsOk() ([]VacancyWorkingTimeIntervalItemOutput, bool) {
	if o == nil || IsNil(o.WorkingTimeIntervals) {
		return nil, false
	}
	return o.WorkingTimeIntervals, true
}

// HasWorkingTimeIntervals returns a boolean if a field has been set.
func (o *VacancyDraftVacancyDraftFull) HasWorkingTimeIntervals() bool {
	if o != nil && IsNil(o.WorkingTimeIntervals) {
		return true
	}

	return false
}

// SetWorkingTimeIntervals gets a reference to the given []VacancyWorkingTimeIntervalItemOutput and assigns it to the WorkingTimeIntervals field.
func (o *VacancyDraftVacancyDraftFull) SetWorkingTimeIntervals(v []VacancyWorkingTimeIntervalItemOutput) {
	o.WorkingTimeIntervals = v
}

// GetWorkingTimeModes returns the WorkingTimeModes field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *VacancyDraftVacancyDraftFull) GetWorkingTimeModes() []VacancyWorkingTimeModeItemOutput {
	if o == nil {
		var ret []VacancyWorkingTimeModeItemOutput
		return ret
	}
	return o.WorkingTimeModes
}

// GetWorkingTimeModesOk returns a tuple with the WorkingTimeModes field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *VacancyDraftVacancyDraftFull) GetWorkingTimeModesOk() ([]VacancyWorkingTimeModeItemOutput, bool) {
	if o == nil || IsNil(o.WorkingTimeModes) {
		return nil, false
	}
	return o.WorkingTimeModes, true
}

// HasWorkingTimeModes returns a boolean if a field has been set.
func (o *VacancyDraftVacancyDraftFull) HasWorkingTimeModes() bool {
	if o != nil && IsNil(o.WorkingTimeModes) {
		return true
	}

	return false
}

// SetWorkingTimeModes gets a reference to the given []VacancyWorkingTimeModeItemOutput and assigns it to the WorkingTimeModes field.
func (o *VacancyDraftVacancyDraftFull) SetWorkingTimeModes(v []VacancyWorkingTimeModeItemOutput) {
	o.WorkingTimeModes = v
}

// GetAddress returns the Address field value
func (o *VacancyDraftVacancyDraftFull) GetAddress() VacancyDraftAddressOutput {
	if o == nil {
		var ret VacancyDraftAddressOutput
		return ret
	}

	return o.Address
}

// GetAddressOk returns a tuple with the Address field value
// and a boolean to check if the value has been set.
func (o *VacancyDraftVacancyDraftFull) GetAddressOk() (*VacancyDraftAddressOutput, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Address, true
}

// SetAddress sets field value
func (o *VacancyDraftVacancyDraftFull) SetAddress(v VacancyDraftAddressOutput) {
	o.Address = v
}

// GetAreas returns the Areas field value
func (o *VacancyDraftVacancyDraftFull) GetAreas() []VacancyAreaOutput {
	if o == nil {
		var ret []VacancyAreaOutput
		return ret
	}

	return o.Areas
}

// GetAreasOk returns a tuple with the Areas field value
// and a boolean to check if the value has been set.
func (o *VacancyDraftVacancyDraftFull) GetAreasOk() ([]VacancyAreaOutput, bool) {
	if o == nil {
		return nil, false
	}
	return o.Areas, true
}

// SetAreas sets field value
func (o *VacancyDraftVacancyDraftFull) SetAreas(v []VacancyAreaOutput) {
	o.Areas = v
}

// GetAssignedManager returns the AssignedManager field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *VacancyDraftVacancyDraftFull) GetAssignedManager() VacancyDraftAssignedManager {
	if o == nil || IsNil(o.AssignedManager.Get()) {
		var ret VacancyDraftAssignedManager
		return ret
	}
	return *o.AssignedManager.Get()
}

// GetAssignedManagerOk returns a tuple with the AssignedManager field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *VacancyDraftVacancyDraftFull) GetAssignedManagerOk() (*VacancyDraftAssignedManager, bool) {
	if o == nil {
		return nil, false
	}
	return o.AssignedManager.Get(), o.AssignedManager.IsSet()
}

// HasAssignedManager returns a boolean if a field has been set.
func (o *VacancyDraftVacancyDraftFull) HasAssignedManager() bool {
	if o != nil && o.AssignedManager.IsSet() {
		return true
	}

	return false
}

// SetAssignedManager gets a reference to the given NullableVacancyDraftAssignedManager and assigns it to the AssignedManager field.
func (o *VacancyDraftVacancyDraftFull) SetAssignedManager(v VacancyDraftAssignedManager) {
	o.AssignedManager.Set(&v)
}
// SetAssignedManagerNil sets the value for AssignedManager to be an explicit nil
func (o *VacancyDraftVacancyDraftFull) SetAssignedManagerNil() {
	o.AssignedManager.Set(nil)
}

// UnsetAssignedManager ensures that no value is present for AssignedManager, not even an explicit nil
func (o *VacancyDraftVacancyDraftFull) UnsetAssignedManager() {
	o.AssignedManager.Unset()
}

// GetBrandedTemplate returns the BrandedTemplate field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *VacancyDraftVacancyDraftFull) GetBrandedTemplate() VacancyBrandedTemplate {
	if o == nil || IsNil(o.BrandedTemplate.Get()) {
		var ret VacancyBrandedTemplate
		return ret
	}
	return *o.BrandedTemplate.Get()
}

// GetBrandedTemplateOk returns a tuple with the BrandedTemplate field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *VacancyDraftVacancyDraftFull) GetBrandedTemplateOk() (*VacancyBrandedTemplate, bool) {
	if o == nil {
		return nil, false
	}
	return o.BrandedTemplate.Get(), o.BrandedTemplate.IsSet()
}

// HasBrandedTemplate returns a boolean if a field has been set.
func (o *VacancyDraftVacancyDraftFull) HasBrandedTemplate() bool {
	if o != nil && o.BrandedTemplate.IsSet() {
		return true
	}

	return false
}

// SetBrandedTemplate gets a reference to the given NullableVacancyBrandedTemplate and assigns it to the BrandedTemplate field.
func (o *VacancyDraftVacancyDraftFull) SetBrandedTemplate(v VacancyBrandedTemplate) {
	o.BrandedTemplate.Set(&v)
}
// SetBrandedTemplateNil sets the value for BrandedTemplate to be an explicit nil
func (o *VacancyDraftVacancyDraftFull) SetBrandedTemplateNil() {
	o.BrandedTemplate.Set(nil)
}

// UnsetBrandedTemplate ensures that no value is present for BrandedTemplate, not even an explicit nil
func (o *VacancyDraftVacancyDraftFull) UnsetBrandedTemplate() {
	o.BrandedTemplate.Unset()
}

// GetContacts returns the Contacts field value
func (o *VacancyDraftVacancyDraftFull) GetContacts() VacancyDraftContactsWithFullPhone {
	if o == nil {
		var ret VacancyDraftContactsWithFullPhone
		return ret
	}

	return o.Contacts
}

// GetContactsOk returns a tuple with the Contacts field value
// and a boolean to check if the value has been set.
func (o *VacancyDraftVacancyDraftFull) GetContactsOk() (*VacancyDraftContactsWithFullPhone, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Contacts, true
}

// SetContacts sets field value
func (o *VacancyDraftVacancyDraftFull) SetContacts(v VacancyDraftContactsWithFullPhone) {
	o.Contacts = v
}

// GetCustomEmployerName returns the CustomEmployerName field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *VacancyDraftVacancyDraftFull) GetCustomEmployerName() string {
	if o == nil || IsNil(o.CustomEmployerName.Get()) {
		var ret string
		return ret
	}
	return *o.CustomEmployerName.Get()
}

// GetCustomEmployerNameOk returns a tuple with the CustomEmployerName field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *VacancyDraftVacancyDraftFull) GetCustomEmployerNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.CustomEmployerName.Get(), o.CustomEmployerName.IsSet()
}

// HasCustomEmployerName returns a boolean if a field has been set.
func (o *VacancyDraftVacancyDraftFull) HasCustomEmployerName() bool {
	if o != nil && o.CustomEmployerName.IsSet() {
		return true
	}

	return false
}

// SetCustomEmployerName gets a reference to the given NullableString and assigns it to the CustomEmployerName field.
func (o *VacancyDraftVacancyDraftFull) SetCustomEmployerName(v string) {
	o.CustomEmployerName.Set(&v)
}
// SetCustomEmployerNameNil sets the value for CustomEmployerName to be an explicit nil
func (o *VacancyDraftVacancyDraftFull) SetCustomEmployerNameNil() {
	o.CustomEmployerName.Set(nil)
}

// UnsetCustomEmployerName ensures that no value is present for CustomEmployerName, not even an explicit nil
func (o *VacancyDraftVacancyDraftFull) UnsetCustomEmployerName() {
	o.CustomEmployerName.Unset()
}

// GetEmployer returns the Employer field value
func (o *VacancyDraftVacancyDraftFull) GetEmployer() VacancyDraftEmployer {
	if o == nil {
		var ret VacancyDraftEmployer
		return ret
	}

	return o.Employer
}

// GetEmployerOk returns a tuple with the Employer field value
// and a boolean to check if the value has been set.
func (o *VacancyDraftVacancyDraftFull) GetEmployerOk() (*VacancyDraftEmployer, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Employer, true
}

// SetEmployer sets field value
func (o *VacancyDraftVacancyDraftFull) SetEmployer(v VacancyDraftEmployer) {
	o.Employer = v
}

// GetMetaInfo returns the MetaInfo field value
func (o *VacancyDraftVacancyDraftFull) GetMetaInfo() VacancyDraftVacancyDraftBase {
	if o == nil {
		var ret VacancyDraftVacancyDraftBase
		return ret
	}

	return o.MetaInfo
}

// GetMetaInfoOk returns a tuple with the MetaInfo field value
// and a boolean to check if the value has been set.
func (o *VacancyDraftVacancyDraftFull) GetMetaInfoOk() (*VacancyDraftVacancyDraftBase, bool) {
	if o == nil {
		return nil, false
	}
	return &o.MetaInfo, true
}

// SetMetaInfo sets field value
func (o *VacancyDraftVacancyDraftFull) SetMetaInfo(v VacancyDraftVacancyDraftBase) {
	o.MetaInfo = v
}

// GetWithZp returns the WithZp field value
func (o *VacancyDraftVacancyDraftFull) GetWithZp() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.WithZp
}

// GetWithZpOk returns a tuple with the WithZp field value
// and a boolean to check if the value has been set.
func (o *VacancyDraftVacancyDraftFull) GetWithZpOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.WithZp, true
}

// SetWithZp sets field value
func (o *VacancyDraftVacancyDraftFull) SetWithZp(v bool) {
	o.WithZp = v
}

func (o VacancyDraftVacancyDraftFull) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o VacancyDraftVacancyDraftFull) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["accept_handicapped"] = o.AcceptHandicapped
	toSerialize["accept_incomplete_resumes"] = o.AcceptIncompleteResumes
	toSerialize["accept_kids"] = o.AcceptKids
	if o.AcceptTemporary.IsSet() {
		toSerialize["accept_temporary"] = o.AcceptTemporary.Get()
	}
	toSerialize["allow_messages"] = o.AllowMessages
	toSerialize["billing_type"] = o.BillingType.Get()
	if o.Code.IsSet() {
		toSerialize["code"] = o.Code.Get()
	}
	if o.Department.IsSet() {
		toSerialize["department"] = o.Department.Get()
	}
	toSerialize["description"] = o.Description
	toSerialize["driver_license_types"] = o.DriverLicenseTypes
	if o.Employment.IsSet() {
		toSerialize["employment"] = o.Employment.Get()
	}
	toSerialize["experience"] = o.Experience.Get()
	toSerialize["has_test"] = o.HasTest
	toSerialize["key_skills"] = o.KeySkills
	toSerialize["languages"] = o.Languages
	toSerialize["manager"] = o.Manager
	toSerialize["name"] = o.Name
	toSerialize["professional_roles"] = o.ProfessionalRoles
	toSerialize["response_letter_required"] = o.ResponseLetterRequired
	toSerialize["response_notifications"] = o.ResponseNotifications
	if o.ResponseUrl.IsSet() {
		toSerialize["response_url"] = o.ResponseUrl.Get()
	}
	if o.Salary.IsSet() {
		toSerialize["salary"] = o.Salary.Get()
	}
	toSerialize["schedule"] = o.Schedule
	if o.Test.IsSet() {
		toSerialize["test"] = o.Test.Get()
	}
	toSerialize["type"] = o.Type
	if o.WorkingDays != nil {
		toSerialize["working_days"] = o.WorkingDays
	}
	if o.WorkingTimeIntervals != nil {
		toSerialize["working_time_intervals"] = o.WorkingTimeIntervals
	}
	if o.WorkingTimeModes != nil {
		toSerialize["working_time_modes"] = o.WorkingTimeModes
	}
	toSerialize["address"] = o.Address
	toSerialize["areas"] = o.Areas
	if o.AssignedManager.IsSet() {
		toSerialize["assigned_manager"] = o.AssignedManager.Get()
	}
	if o.BrandedTemplate.IsSet() {
		toSerialize["branded_template"] = o.BrandedTemplate.Get()
	}
	toSerialize["contacts"] = o.Contacts
	if o.CustomEmployerName.IsSet() {
		toSerialize["custom_employer_name"] = o.CustomEmployerName.Get()
	}
	toSerialize["employer"] = o.Employer
	toSerialize["meta_info"] = o.MetaInfo
	toSerialize["with_zp"] = o.WithZp
	return toSerialize, nil
}

func (o *VacancyDraftVacancyDraftFull) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"accept_handicapped",
		"accept_incomplete_resumes",
		"accept_kids",
		"allow_messages",
		"billing_type",
		"description",
		"driver_license_types",
		"experience",
		"has_test",
		"key_skills",
		"languages",
		"manager",
		"name",
		"professional_roles",
		"response_letter_required",
		"response_notifications",
		"schedule",
		"type",
		"address",
		"areas",
		"contacts",
		"employer",
		"meta_info",
		"with_zp",
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

	varVacancyDraftVacancyDraftFull := _VacancyDraftVacancyDraftFull{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varVacancyDraftVacancyDraftFull)

	if err != nil {
		return err
	}

	*o = VacancyDraftVacancyDraftFull(varVacancyDraftVacancyDraftFull)

	return err
}

type NullableVacancyDraftVacancyDraftFull struct {
	value *VacancyDraftVacancyDraftFull
	isSet bool
}

func (v NullableVacancyDraftVacancyDraftFull) Get() *VacancyDraftVacancyDraftFull {
	return v.value
}

func (v *NullableVacancyDraftVacancyDraftFull) Set(val *VacancyDraftVacancyDraftFull) {
	v.value = val
	v.isSet = true
}

func (v NullableVacancyDraftVacancyDraftFull) IsSet() bool {
	return v.isSet
}

func (v *NullableVacancyDraftVacancyDraftFull) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableVacancyDraftVacancyDraftFull(val *VacancyDraftVacancyDraftFull) *NullableVacancyDraftVacancyDraftFull {
	return &NullableVacancyDraftVacancyDraftFull{value: val, isSet: true}
}

func (v NullableVacancyDraftVacancyDraftFull) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableVacancyDraftVacancyDraftFull) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


