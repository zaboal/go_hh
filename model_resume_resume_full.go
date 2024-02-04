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

// checks if the ResumeResumeFull type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ResumeResumeFull{}

// ResumeResumeFull struct for ResumeResumeFull
type ResumeResumeFull struct {
	// URL резюме на сайте
	AlternateUrl string `json:"alternate_url"`
	// Идентификатор резюме
	Id string `json:"id"`
	// Желаемая должность
	Title NullableString `json:"title"`
	// Возраст
	Age NullableFloat32 `json:"age,omitempty"`
	Area NullableIncludesIdNameUrl `json:"area,omitempty"`
	// Доступен ли просмотр контактной информации в резюме текущему работодателю
	CanViewFullInfo NullableBool `json:"can_view_full_info,omitempty"`
	// Список сертификатов соискателя
	Certificate []ResumeObjectsCertificate `json:"certificate"`
	// Дата и время создания резюме
	CreatedAt string `json:"created_at"`
	Download ResumeResumeProfileAllOfDownload `json:"download"`
	Education ResumeResumeProfileAllOfEducation `json:"education"`
	// Опыт работы
	Experience []ResumeObjectsExperience `json:"experience"`
	// Имя
	FirstName NullableString `json:"first_name,omitempty"`
	Gender NullableIncludesIdName `json:"gender,omitempty"`
	// Справочник [Список скрытых полей](https://github.com/hhru/api/blob/master/docs/employer_resumes.md#hidden-fields). Возможные значения элементов приведены в поле `resume_hidden_fields` [справочника полей](#tag/Obshie-spravochniki/operation/get-dictionaries)
	HiddenFields []IncludesIdName `json:"hidden_fields"`
	// Фамилия
	LastName NullableString `json:"last_name,omitempty"`
	// Выделено ли резюме в поиске
	Marked *bool `json:"marked,omitempty"`
	// Отчество
	MiddleName NullableString `json:"middle_name,omitempty"`
	Platform *ResumeResumeProfileAllOfPlatform `json:"platform,omitempty"`
	Salary NullableResumeObjectsSalaryProperties `json:"salary,omitempty"`
	TotalExperience NullableResumeObjectsTotalExperience `json:"total_experience,omitempty"`
	// Дата и время обновления резюме
	UpdatedAt string `json:"updated_at"`
	// День рождения (в формате `ГГГГ-ММ-ДД`)
	BirthDate NullableString `json:"birth_date,omitempty"`
	BusinessTripReadiness ResumeResumeFullAllOfBusinessTripReadiness `json:"business_trip_readiness"`
	// Список гражданств соискателя. Элементы [справочника регионов](#tag/Obshie-spravochniki/operation/get-areas)
	Citizenship []IncludesIdNameUrl `json:"citizenship"`
	// Список контактов соискателя
	Contact []IncludesContact `json:"contact"`
	Creds NullableCredsCreds `json:"creds,omitempty"`
	// Список категорий водительских прав соискателя
	DriverLicenseTypes []ResumeObjectsDriverLicenseTypes `json:"driver_license_types"`
	// Deprecated
	Employment *ResumeResumeFullAllOfEmployment `json:"employment,omitempty"`
	// Список подходящих соискателю типов занятостей. Элементы справочника [employment](#tag/Obshie-spravochniki/operation/get-dictionaries)
	Employments []IncludesIdName `json:"employments"`
	// Наличие личного автомобиля у соискателя
	HasVehicle NullableBool `json:"has_vehicle,omitempty"`
	// Список языков, которыми владеет соискатель. Элементы справочника [languages](#tag/Obshie-spravochniki/operation/get-languages)
	Language []IncludesLanguageLevel `json:"language"`
	Metro NullableResumeObjectsMetroStation `json:"metro,omitempty"`
	// Платные услуги по резюме
	PaidServices []ResumeObjectsPaidServices `json:"paid_services"`
	// Массив объектов профролей
	ProfessionalRoles []IncludesIdName `json:"professional_roles,omitempty"`
	// Список рекомендаций
	Recommendation []ResumeObjectsRecommendation `json:"recommendation"`
	Relocation ResumeResumeFullAllOfRelocation `json:"relocation"`
	ResumeLocale ResumeResumeFullAllOfResumeLocale `json:"resume_locale"`
	// Deprecated
	Schedule ResumeResumeFullAllOfSchedule `json:"schedule"`
	// Список подходящих соискателю графиков работы. Элементы справочника [schedule](#tag/Obshie-spravochniki/operation/get-dictionaries)
	Schedules []IncludesIdName `json:"schedules"`
	// Профили в соц. сетях и других сервисах
	Site []ResumeObjectsSite `json:"site"`
	// Ключевые навыки (список уникальных строк)
	SkillSet []string `json:"skill_set"`
	// Дополнительная информация, описание навыков в свободной форме
	Skills NullableString `json:"skills,omitempty"`
	TravelTime ResumeResumeFullAllOfTravelTime `json:"travel_time"`
	// Список регионов, в которых соискатель имеет разрешение на работу. Элементы [справочника регионов](#tag/Obshie-spravochniki/operation/get-areas) 
	WorkTicket []IncludesIdNameUrl `json:"work_ticket"`
}

type _ResumeResumeFull ResumeResumeFull

// NewResumeResumeFull instantiates a new ResumeResumeFull object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewResumeResumeFull(alternateUrl string, id string, title NullableString, certificate []ResumeObjectsCertificate, createdAt string, download ResumeResumeProfileAllOfDownload, education ResumeResumeProfileAllOfEducation, experience []ResumeObjectsExperience, hiddenFields []IncludesIdName, updatedAt string, businessTripReadiness ResumeResumeFullAllOfBusinessTripReadiness, citizenship []IncludesIdNameUrl, contact []IncludesContact, driverLicenseTypes []ResumeObjectsDriverLicenseTypes, employments []IncludesIdName, language []IncludesLanguageLevel, paidServices []ResumeObjectsPaidServices, recommendation []ResumeObjectsRecommendation, relocation ResumeResumeFullAllOfRelocation, resumeLocale ResumeResumeFullAllOfResumeLocale, schedule ResumeResumeFullAllOfSchedule, schedules []IncludesIdName, site []ResumeObjectsSite, skillSet []string, travelTime ResumeResumeFullAllOfTravelTime, workTicket []IncludesIdNameUrl) *ResumeResumeFull {
	this := ResumeResumeFull{}
	this.AlternateUrl = alternateUrl
	this.Id = id
	this.Title = title
	this.Certificate = certificate
	this.CreatedAt = createdAt
	this.Download = download
	this.Education = education
	this.Experience = experience
	this.HiddenFields = hiddenFields
	var marked bool = false
	this.Marked = &marked
	this.UpdatedAt = updatedAt
	this.BusinessTripReadiness = businessTripReadiness
	this.Citizenship = citizenship
	this.Contact = contact
	this.DriverLicenseTypes = driverLicenseTypes
	this.Employments = employments
	this.Language = language
	this.PaidServices = paidServices
	this.Recommendation = recommendation
	this.Relocation = relocation
	this.ResumeLocale = resumeLocale
	this.Schedule = schedule
	this.Schedules = schedules
	this.Site = site
	this.SkillSet = skillSet
	this.TravelTime = travelTime
	this.WorkTicket = workTicket
	return &this
}

// NewResumeResumeFullWithDefaults instantiates a new ResumeResumeFull object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewResumeResumeFullWithDefaults() *ResumeResumeFull {
	this := ResumeResumeFull{}
	var marked bool = false
	this.Marked = &marked
	return &this
}

// GetAlternateUrl returns the AlternateUrl field value
func (o *ResumeResumeFull) GetAlternateUrl() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.AlternateUrl
}

// GetAlternateUrlOk returns a tuple with the AlternateUrl field value
// and a boolean to check if the value has been set.
func (o *ResumeResumeFull) GetAlternateUrlOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AlternateUrl, true
}

// SetAlternateUrl sets field value
func (o *ResumeResumeFull) SetAlternateUrl(v string) {
	o.AlternateUrl = v
}

// GetId returns the Id field value
func (o *ResumeResumeFull) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *ResumeResumeFull) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *ResumeResumeFull) SetId(v string) {
	o.Id = v
}

// GetTitle returns the Title field value
// If the value is explicit nil, the zero value for string will be returned
func (o *ResumeResumeFull) GetTitle() string {
	if o == nil || o.Title.Get() == nil {
		var ret string
		return ret
	}

	return *o.Title.Get()
}

// GetTitleOk returns a tuple with the Title field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ResumeResumeFull) GetTitleOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Title.Get(), o.Title.IsSet()
}

// SetTitle sets field value
func (o *ResumeResumeFull) SetTitle(v string) {
	o.Title.Set(&v)
}

// GetAge returns the Age field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ResumeResumeFull) GetAge() float32 {
	if o == nil || IsNil(o.Age.Get()) {
		var ret float32
		return ret
	}
	return *o.Age.Get()
}

// GetAgeOk returns a tuple with the Age field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ResumeResumeFull) GetAgeOk() (*float32, bool) {
	if o == nil {
		return nil, false
	}
	return o.Age.Get(), o.Age.IsSet()
}

// HasAge returns a boolean if a field has been set.
func (o *ResumeResumeFull) HasAge() bool {
	if o != nil && o.Age.IsSet() {
		return true
	}

	return false
}

// SetAge gets a reference to the given NullableFloat32 and assigns it to the Age field.
func (o *ResumeResumeFull) SetAge(v float32) {
	o.Age.Set(&v)
}
// SetAgeNil sets the value for Age to be an explicit nil
func (o *ResumeResumeFull) SetAgeNil() {
	o.Age.Set(nil)
}

// UnsetAge ensures that no value is present for Age, not even an explicit nil
func (o *ResumeResumeFull) UnsetAge() {
	o.Age.Unset()
}

// GetArea returns the Area field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ResumeResumeFull) GetArea() IncludesIdNameUrl {
	if o == nil || IsNil(o.Area.Get()) {
		var ret IncludesIdNameUrl
		return ret
	}
	return *o.Area.Get()
}

// GetAreaOk returns a tuple with the Area field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ResumeResumeFull) GetAreaOk() (*IncludesIdNameUrl, bool) {
	if o == nil {
		return nil, false
	}
	return o.Area.Get(), o.Area.IsSet()
}

// HasArea returns a boolean if a field has been set.
func (o *ResumeResumeFull) HasArea() bool {
	if o != nil && o.Area.IsSet() {
		return true
	}

	return false
}

// SetArea gets a reference to the given NullableIncludesIdNameUrl and assigns it to the Area field.
func (o *ResumeResumeFull) SetArea(v IncludesIdNameUrl) {
	o.Area.Set(&v)
}
// SetAreaNil sets the value for Area to be an explicit nil
func (o *ResumeResumeFull) SetAreaNil() {
	o.Area.Set(nil)
}

// UnsetArea ensures that no value is present for Area, not even an explicit nil
func (o *ResumeResumeFull) UnsetArea() {
	o.Area.Unset()
}

// GetCanViewFullInfo returns the CanViewFullInfo field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ResumeResumeFull) GetCanViewFullInfo() bool {
	if o == nil || IsNil(o.CanViewFullInfo.Get()) {
		var ret bool
		return ret
	}
	return *o.CanViewFullInfo.Get()
}

// GetCanViewFullInfoOk returns a tuple with the CanViewFullInfo field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ResumeResumeFull) GetCanViewFullInfoOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return o.CanViewFullInfo.Get(), o.CanViewFullInfo.IsSet()
}

// HasCanViewFullInfo returns a boolean if a field has been set.
func (o *ResumeResumeFull) HasCanViewFullInfo() bool {
	if o != nil && o.CanViewFullInfo.IsSet() {
		return true
	}

	return false
}

// SetCanViewFullInfo gets a reference to the given NullableBool and assigns it to the CanViewFullInfo field.
func (o *ResumeResumeFull) SetCanViewFullInfo(v bool) {
	o.CanViewFullInfo.Set(&v)
}
// SetCanViewFullInfoNil sets the value for CanViewFullInfo to be an explicit nil
func (o *ResumeResumeFull) SetCanViewFullInfoNil() {
	o.CanViewFullInfo.Set(nil)
}

// UnsetCanViewFullInfo ensures that no value is present for CanViewFullInfo, not even an explicit nil
func (o *ResumeResumeFull) UnsetCanViewFullInfo() {
	o.CanViewFullInfo.Unset()
}

// GetCertificate returns the Certificate field value
func (o *ResumeResumeFull) GetCertificate() []ResumeObjectsCertificate {
	if o == nil {
		var ret []ResumeObjectsCertificate
		return ret
	}

	return o.Certificate
}

// GetCertificateOk returns a tuple with the Certificate field value
// and a boolean to check if the value has been set.
func (o *ResumeResumeFull) GetCertificateOk() ([]ResumeObjectsCertificate, bool) {
	if o == nil {
		return nil, false
	}
	return o.Certificate, true
}

// SetCertificate sets field value
func (o *ResumeResumeFull) SetCertificate(v []ResumeObjectsCertificate) {
	o.Certificate = v
}

// GetCreatedAt returns the CreatedAt field value
func (o *ResumeResumeFull) GetCreatedAt() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value
// and a boolean to check if the value has been set.
func (o *ResumeResumeFull) GetCreatedAtOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CreatedAt, true
}

// SetCreatedAt sets field value
func (o *ResumeResumeFull) SetCreatedAt(v string) {
	o.CreatedAt = v
}

// GetDownload returns the Download field value
func (o *ResumeResumeFull) GetDownload() ResumeResumeProfileAllOfDownload {
	if o == nil {
		var ret ResumeResumeProfileAllOfDownload
		return ret
	}

	return o.Download
}

// GetDownloadOk returns a tuple with the Download field value
// and a boolean to check if the value has been set.
func (o *ResumeResumeFull) GetDownloadOk() (*ResumeResumeProfileAllOfDownload, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Download, true
}

// SetDownload sets field value
func (o *ResumeResumeFull) SetDownload(v ResumeResumeProfileAllOfDownload) {
	o.Download = v
}

// GetEducation returns the Education field value
func (o *ResumeResumeFull) GetEducation() ResumeResumeProfileAllOfEducation {
	if o == nil {
		var ret ResumeResumeProfileAllOfEducation
		return ret
	}

	return o.Education
}

// GetEducationOk returns a tuple with the Education field value
// and a boolean to check if the value has been set.
func (o *ResumeResumeFull) GetEducationOk() (*ResumeResumeProfileAllOfEducation, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Education, true
}

// SetEducation sets field value
func (o *ResumeResumeFull) SetEducation(v ResumeResumeProfileAllOfEducation) {
	o.Education = v
}

// GetExperience returns the Experience field value
func (o *ResumeResumeFull) GetExperience() []ResumeObjectsExperience {
	if o == nil {
		var ret []ResumeObjectsExperience
		return ret
	}

	return o.Experience
}

// GetExperienceOk returns a tuple with the Experience field value
// and a boolean to check if the value has been set.
func (o *ResumeResumeFull) GetExperienceOk() ([]ResumeObjectsExperience, bool) {
	if o == nil {
		return nil, false
	}
	return o.Experience, true
}

// SetExperience sets field value
func (o *ResumeResumeFull) SetExperience(v []ResumeObjectsExperience) {
	o.Experience = v
}

// GetFirstName returns the FirstName field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ResumeResumeFull) GetFirstName() string {
	if o == nil || IsNil(o.FirstName.Get()) {
		var ret string
		return ret
	}
	return *o.FirstName.Get()
}

// GetFirstNameOk returns a tuple with the FirstName field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ResumeResumeFull) GetFirstNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.FirstName.Get(), o.FirstName.IsSet()
}

// HasFirstName returns a boolean if a field has been set.
func (o *ResumeResumeFull) HasFirstName() bool {
	if o != nil && o.FirstName.IsSet() {
		return true
	}

	return false
}

// SetFirstName gets a reference to the given NullableString and assigns it to the FirstName field.
func (o *ResumeResumeFull) SetFirstName(v string) {
	o.FirstName.Set(&v)
}
// SetFirstNameNil sets the value for FirstName to be an explicit nil
func (o *ResumeResumeFull) SetFirstNameNil() {
	o.FirstName.Set(nil)
}

// UnsetFirstName ensures that no value is present for FirstName, not even an explicit nil
func (o *ResumeResumeFull) UnsetFirstName() {
	o.FirstName.Unset()
}

// GetGender returns the Gender field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ResumeResumeFull) GetGender() IncludesIdName {
	if o == nil || IsNil(o.Gender.Get()) {
		var ret IncludesIdName
		return ret
	}
	return *o.Gender.Get()
}

// GetGenderOk returns a tuple with the Gender field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ResumeResumeFull) GetGenderOk() (*IncludesIdName, bool) {
	if o == nil {
		return nil, false
	}
	return o.Gender.Get(), o.Gender.IsSet()
}

// HasGender returns a boolean if a field has been set.
func (o *ResumeResumeFull) HasGender() bool {
	if o != nil && o.Gender.IsSet() {
		return true
	}

	return false
}

// SetGender gets a reference to the given NullableIncludesIdName and assigns it to the Gender field.
func (o *ResumeResumeFull) SetGender(v IncludesIdName) {
	o.Gender.Set(&v)
}
// SetGenderNil sets the value for Gender to be an explicit nil
func (o *ResumeResumeFull) SetGenderNil() {
	o.Gender.Set(nil)
}

// UnsetGender ensures that no value is present for Gender, not even an explicit nil
func (o *ResumeResumeFull) UnsetGender() {
	o.Gender.Unset()
}

// GetHiddenFields returns the HiddenFields field value
func (o *ResumeResumeFull) GetHiddenFields() []IncludesIdName {
	if o == nil {
		var ret []IncludesIdName
		return ret
	}

	return o.HiddenFields
}

// GetHiddenFieldsOk returns a tuple with the HiddenFields field value
// and a boolean to check if the value has been set.
func (o *ResumeResumeFull) GetHiddenFieldsOk() ([]IncludesIdName, bool) {
	if o == nil {
		return nil, false
	}
	return o.HiddenFields, true
}

// SetHiddenFields sets field value
func (o *ResumeResumeFull) SetHiddenFields(v []IncludesIdName) {
	o.HiddenFields = v
}

// GetLastName returns the LastName field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ResumeResumeFull) GetLastName() string {
	if o == nil || IsNil(o.LastName.Get()) {
		var ret string
		return ret
	}
	return *o.LastName.Get()
}

// GetLastNameOk returns a tuple with the LastName field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ResumeResumeFull) GetLastNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.LastName.Get(), o.LastName.IsSet()
}

// HasLastName returns a boolean if a field has been set.
func (o *ResumeResumeFull) HasLastName() bool {
	if o != nil && o.LastName.IsSet() {
		return true
	}

	return false
}

// SetLastName gets a reference to the given NullableString and assigns it to the LastName field.
func (o *ResumeResumeFull) SetLastName(v string) {
	o.LastName.Set(&v)
}
// SetLastNameNil sets the value for LastName to be an explicit nil
func (o *ResumeResumeFull) SetLastNameNil() {
	o.LastName.Set(nil)
}

// UnsetLastName ensures that no value is present for LastName, not even an explicit nil
func (o *ResumeResumeFull) UnsetLastName() {
	o.LastName.Unset()
}

// GetMarked returns the Marked field value if set, zero value otherwise.
func (o *ResumeResumeFull) GetMarked() bool {
	if o == nil || IsNil(o.Marked) {
		var ret bool
		return ret
	}
	return *o.Marked
}

// GetMarkedOk returns a tuple with the Marked field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ResumeResumeFull) GetMarkedOk() (*bool, bool) {
	if o == nil || IsNil(o.Marked) {
		return nil, false
	}
	return o.Marked, true
}

// HasMarked returns a boolean if a field has been set.
func (o *ResumeResumeFull) HasMarked() bool {
	if o != nil && !IsNil(o.Marked) {
		return true
	}

	return false
}

// SetMarked gets a reference to the given bool and assigns it to the Marked field.
func (o *ResumeResumeFull) SetMarked(v bool) {
	o.Marked = &v
}

// GetMiddleName returns the MiddleName field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ResumeResumeFull) GetMiddleName() string {
	if o == nil || IsNil(o.MiddleName.Get()) {
		var ret string
		return ret
	}
	return *o.MiddleName.Get()
}

// GetMiddleNameOk returns a tuple with the MiddleName field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ResumeResumeFull) GetMiddleNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.MiddleName.Get(), o.MiddleName.IsSet()
}

// HasMiddleName returns a boolean if a field has been set.
func (o *ResumeResumeFull) HasMiddleName() bool {
	if o != nil && o.MiddleName.IsSet() {
		return true
	}

	return false
}

// SetMiddleName gets a reference to the given NullableString and assigns it to the MiddleName field.
func (o *ResumeResumeFull) SetMiddleName(v string) {
	o.MiddleName.Set(&v)
}
// SetMiddleNameNil sets the value for MiddleName to be an explicit nil
func (o *ResumeResumeFull) SetMiddleNameNil() {
	o.MiddleName.Set(nil)
}

// UnsetMiddleName ensures that no value is present for MiddleName, not even an explicit nil
func (o *ResumeResumeFull) UnsetMiddleName() {
	o.MiddleName.Unset()
}

// GetPlatform returns the Platform field value if set, zero value otherwise.
func (o *ResumeResumeFull) GetPlatform() ResumeResumeProfileAllOfPlatform {
	if o == nil || IsNil(o.Platform) {
		var ret ResumeResumeProfileAllOfPlatform
		return ret
	}
	return *o.Platform
}

// GetPlatformOk returns a tuple with the Platform field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ResumeResumeFull) GetPlatformOk() (*ResumeResumeProfileAllOfPlatform, bool) {
	if o == nil || IsNil(o.Platform) {
		return nil, false
	}
	return o.Platform, true
}

// HasPlatform returns a boolean if a field has been set.
func (o *ResumeResumeFull) HasPlatform() bool {
	if o != nil && !IsNil(o.Platform) {
		return true
	}

	return false
}

// SetPlatform gets a reference to the given ResumeResumeProfileAllOfPlatform and assigns it to the Platform field.
func (o *ResumeResumeFull) SetPlatform(v ResumeResumeProfileAllOfPlatform) {
	o.Platform = &v
}

// GetSalary returns the Salary field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ResumeResumeFull) GetSalary() ResumeObjectsSalaryProperties {
	if o == nil || IsNil(o.Salary.Get()) {
		var ret ResumeObjectsSalaryProperties
		return ret
	}
	return *o.Salary.Get()
}

// GetSalaryOk returns a tuple with the Salary field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ResumeResumeFull) GetSalaryOk() (*ResumeObjectsSalaryProperties, bool) {
	if o == nil {
		return nil, false
	}
	return o.Salary.Get(), o.Salary.IsSet()
}

// HasSalary returns a boolean if a field has been set.
func (o *ResumeResumeFull) HasSalary() bool {
	if o != nil && o.Salary.IsSet() {
		return true
	}

	return false
}

// SetSalary gets a reference to the given NullableResumeObjectsSalaryProperties and assigns it to the Salary field.
func (o *ResumeResumeFull) SetSalary(v ResumeObjectsSalaryProperties) {
	o.Salary.Set(&v)
}
// SetSalaryNil sets the value for Salary to be an explicit nil
func (o *ResumeResumeFull) SetSalaryNil() {
	o.Salary.Set(nil)
}

// UnsetSalary ensures that no value is present for Salary, not even an explicit nil
func (o *ResumeResumeFull) UnsetSalary() {
	o.Salary.Unset()
}

// GetTotalExperience returns the TotalExperience field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ResumeResumeFull) GetTotalExperience() ResumeObjectsTotalExperience {
	if o == nil || IsNil(o.TotalExperience.Get()) {
		var ret ResumeObjectsTotalExperience
		return ret
	}
	return *o.TotalExperience.Get()
}

// GetTotalExperienceOk returns a tuple with the TotalExperience field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ResumeResumeFull) GetTotalExperienceOk() (*ResumeObjectsTotalExperience, bool) {
	if o == nil {
		return nil, false
	}
	return o.TotalExperience.Get(), o.TotalExperience.IsSet()
}

// HasTotalExperience returns a boolean if a field has been set.
func (o *ResumeResumeFull) HasTotalExperience() bool {
	if o != nil && o.TotalExperience.IsSet() {
		return true
	}

	return false
}

// SetTotalExperience gets a reference to the given NullableResumeObjectsTotalExperience and assigns it to the TotalExperience field.
func (o *ResumeResumeFull) SetTotalExperience(v ResumeObjectsTotalExperience) {
	o.TotalExperience.Set(&v)
}
// SetTotalExperienceNil sets the value for TotalExperience to be an explicit nil
func (o *ResumeResumeFull) SetTotalExperienceNil() {
	o.TotalExperience.Set(nil)
}

// UnsetTotalExperience ensures that no value is present for TotalExperience, not even an explicit nil
func (o *ResumeResumeFull) UnsetTotalExperience() {
	o.TotalExperience.Unset()
}

// GetUpdatedAt returns the UpdatedAt field value
func (o *ResumeResumeFull) GetUpdatedAt() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.UpdatedAt
}

// GetUpdatedAtOk returns a tuple with the UpdatedAt field value
// and a boolean to check if the value has been set.
func (o *ResumeResumeFull) GetUpdatedAtOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.UpdatedAt, true
}

// SetUpdatedAt sets field value
func (o *ResumeResumeFull) SetUpdatedAt(v string) {
	o.UpdatedAt = v
}

// GetBirthDate returns the BirthDate field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ResumeResumeFull) GetBirthDate() string {
	if o == nil || IsNil(o.BirthDate.Get()) {
		var ret string
		return ret
	}
	return *o.BirthDate.Get()
}

// GetBirthDateOk returns a tuple with the BirthDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ResumeResumeFull) GetBirthDateOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.BirthDate.Get(), o.BirthDate.IsSet()
}

// HasBirthDate returns a boolean if a field has been set.
func (o *ResumeResumeFull) HasBirthDate() bool {
	if o != nil && o.BirthDate.IsSet() {
		return true
	}

	return false
}

// SetBirthDate gets a reference to the given NullableString and assigns it to the BirthDate field.
func (o *ResumeResumeFull) SetBirthDate(v string) {
	o.BirthDate.Set(&v)
}
// SetBirthDateNil sets the value for BirthDate to be an explicit nil
func (o *ResumeResumeFull) SetBirthDateNil() {
	o.BirthDate.Set(nil)
}

// UnsetBirthDate ensures that no value is present for BirthDate, not even an explicit nil
func (o *ResumeResumeFull) UnsetBirthDate() {
	o.BirthDate.Unset()
}

// GetBusinessTripReadiness returns the BusinessTripReadiness field value
func (o *ResumeResumeFull) GetBusinessTripReadiness() ResumeResumeFullAllOfBusinessTripReadiness {
	if o == nil {
		var ret ResumeResumeFullAllOfBusinessTripReadiness
		return ret
	}

	return o.BusinessTripReadiness
}

// GetBusinessTripReadinessOk returns a tuple with the BusinessTripReadiness field value
// and a boolean to check if the value has been set.
func (o *ResumeResumeFull) GetBusinessTripReadinessOk() (*ResumeResumeFullAllOfBusinessTripReadiness, bool) {
	if o == nil {
		return nil, false
	}
	return &o.BusinessTripReadiness, true
}

// SetBusinessTripReadiness sets field value
func (o *ResumeResumeFull) SetBusinessTripReadiness(v ResumeResumeFullAllOfBusinessTripReadiness) {
	o.BusinessTripReadiness = v
}

// GetCitizenship returns the Citizenship field value
func (o *ResumeResumeFull) GetCitizenship() []IncludesIdNameUrl {
	if o == nil {
		var ret []IncludesIdNameUrl
		return ret
	}

	return o.Citizenship
}

// GetCitizenshipOk returns a tuple with the Citizenship field value
// and a boolean to check if the value has been set.
func (o *ResumeResumeFull) GetCitizenshipOk() ([]IncludesIdNameUrl, bool) {
	if o == nil {
		return nil, false
	}
	return o.Citizenship, true
}

// SetCitizenship sets field value
func (o *ResumeResumeFull) SetCitizenship(v []IncludesIdNameUrl) {
	o.Citizenship = v
}

// GetContact returns the Contact field value
func (o *ResumeResumeFull) GetContact() []IncludesContact {
	if o == nil {
		var ret []IncludesContact
		return ret
	}

	return o.Contact
}

// GetContactOk returns a tuple with the Contact field value
// and a boolean to check if the value has been set.
func (o *ResumeResumeFull) GetContactOk() ([]IncludesContact, bool) {
	if o == nil {
		return nil, false
	}
	return o.Contact, true
}

// SetContact sets field value
func (o *ResumeResumeFull) SetContact(v []IncludesContact) {
	o.Contact = v
}

// GetCreds returns the Creds field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ResumeResumeFull) GetCreds() CredsCreds {
	if o == nil || IsNil(o.Creds.Get()) {
		var ret CredsCreds
		return ret
	}
	return *o.Creds.Get()
}

// GetCredsOk returns a tuple with the Creds field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ResumeResumeFull) GetCredsOk() (*CredsCreds, bool) {
	if o == nil {
		return nil, false
	}
	return o.Creds.Get(), o.Creds.IsSet()
}

// HasCreds returns a boolean if a field has been set.
func (o *ResumeResumeFull) HasCreds() bool {
	if o != nil && o.Creds.IsSet() {
		return true
	}

	return false
}

// SetCreds gets a reference to the given NullableCredsCreds and assigns it to the Creds field.
func (o *ResumeResumeFull) SetCreds(v CredsCreds) {
	o.Creds.Set(&v)
}
// SetCredsNil sets the value for Creds to be an explicit nil
func (o *ResumeResumeFull) SetCredsNil() {
	o.Creds.Set(nil)
}

// UnsetCreds ensures that no value is present for Creds, not even an explicit nil
func (o *ResumeResumeFull) UnsetCreds() {
	o.Creds.Unset()
}

// GetDriverLicenseTypes returns the DriverLicenseTypes field value
func (o *ResumeResumeFull) GetDriverLicenseTypes() []ResumeObjectsDriverLicenseTypes {
	if o == nil {
		var ret []ResumeObjectsDriverLicenseTypes
		return ret
	}

	return o.DriverLicenseTypes
}

// GetDriverLicenseTypesOk returns a tuple with the DriverLicenseTypes field value
// and a boolean to check if the value has been set.
func (o *ResumeResumeFull) GetDriverLicenseTypesOk() ([]ResumeObjectsDriverLicenseTypes, bool) {
	if o == nil {
		return nil, false
	}
	return o.DriverLicenseTypes, true
}

// SetDriverLicenseTypes sets field value
func (o *ResumeResumeFull) SetDriverLicenseTypes(v []ResumeObjectsDriverLicenseTypes) {
	o.DriverLicenseTypes = v
}

// GetEmployment returns the Employment field value if set, zero value otherwise.
// Deprecated
func (o *ResumeResumeFull) GetEmployment() ResumeResumeFullAllOfEmployment {
	if o == nil || IsNil(o.Employment) {
		var ret ResumeResumeFullAllOfEmployment
		return ret
	}
	return *o.Employment
}

// GetEmploymentOk returns a tuple with the Employment field value if set, nil otherwise
// and a boolean to check if the value has been set.
// Deprecated
func (o *ResumeResumeFull) GetEmploymentOk() (*ResumeResumeFullAllOfEmployment, bool) {
	if o == nil || IsNil(o.Employment) {
		return nil, false
	}
	return o.Employment, true
}

// HasEmployment returns a boolean if a field has been set.
func (o *ResumeResumeFull) HasEmployment() bool {
	if o != nil && !IsNil(o.Employment) {
		return true
	}

	return false
}

// SetEmployment gets a reference to the given ResumeResumeFullAllOfEmployment and assigns it to the Employment field.
// Deprecated
func (o *ResumeResumeFull) SetEmployment(v ResumeResumeFullAllOfEmployment) {
	o.Employment = &v
}

// GetEmployments returns the Employments field value
func (o *ResumeResumeFull) GetEmployments() []IncludesIdName {
	if o == nil {
		var ret []IncludesIdName
		return ret
	}

	return o.Employments
}

// GetEmploymentsOk returns a tuple with the Employments field value
// and a boolean to check if the value has been set.
func (o *ResumeResumeFull) GetEmploymentsOk() ([]IncludesIdName, bool) {
	if o == nil {
		return nil, false
	}
	return o.Employments, true
}

// SetEmployments sets field value
func (o *ResumeResumeFull) SetEmployments(v []IncludesIdName) {
	o.Employments = v
}

// GetHasVehicle returns the HasVehicle field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ResumeResumeFull) GetHasVehicle() bool {
	if o == nil || IsNil(o.HasVehicle.Get()) {
		var ret bool
		return ret
	}
	return *o.HasVehicle.Get()
}

// GetHasVehicleOk returns a tuple with the HasVehicle field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ResumeResumeFull) GetHasVehicleOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return o.HasVehicle.Get(), o.HasVehicle.IsSet()
}

// HasHasVehicle returns a boolean if a field has been set.
func (o *ResumeResumeFull) HasHasVehicle() bool {
	if o != nil && o.HasVehicle.IsSet() {
		return true
	}

	return false
}

// SetHasVehicle gets a reference to the given NullableBool and assigns it to the HasVehicle field.
func (o *ResumeResumeFull) SetHasVehicle(v bool) {
	o.HasVehicle.Set(&v)
}
// SetHasVehicleNil sets the value for HasVehicle to be an explicit nil
func (o *ResumeResumeFull) SetHasVehicleNil() {
	o.HasVehicle.Set(nil)
}

// UnsetHasVehicle ensures that no value is present for HasVehicle, not even an explicit nil
func (o *ResumeResumeFull) UnsetHasVehicle() {
	o.HasVehicle.Unset()
}

// GetLanguage returns the Language field value
func (o *ResumeResumeFull) GetLanguage() []IncludesLanguageLevel {
	if o == nil {
		var ret []IncludesLanguageLevel
		return ret
	}

	return o.Language
}

// GetLanguageOk returns a tuple with the Language field value
// and a boolean to check if the value has been set.
func (o *ResumeResumeFull) GetLanguageOk() ([]IncludesLanguageLevel, bool) {
	if o == nil {
		return nil, false
	}
	return o.Language, true
}

// SetLanguage sets field value
func (o *ResumeResumeFull) SetLanguage(v []IncludesLanguageLevel) {
	o.Language = v
}

// GetMetro returns the Metro field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ResumeResumeFull) GetMetro() ResumeObjectsMetroStation {
	if o == nil || IsNil(o.Metro.Get()) {
		var ret ResumeObjectsMetroStation
		return ret
	}
	return *o.Metro.Get()
}

// GetMetroOk returns a tuple with the Metro field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ResumeResumeFull) GetMetroOk() (*ResumeObjectsMetroStation, bool) {
	if o == nil {
		return nil, false
	}
	return o.Metro.Get(), o.Metro.IsSet()
}

// HasMetro returns a boolean if a field has been set.
func (o *ResumeResumeFull) HasMetro() bool {
	if o != nil && o.Metro.IsSet() {
		return true
	}

	return false
}

// SetMetro gets a reference to the given NullableResumeObjectsMetroStation and assigns it to the Metro field.
func (o *ResumeResumeFull) SetMetro(v ResumeObjectsMetroStation) {
	o.Metro.Set(&v)
}
// SetMetroNil sets the value for Metro to be an explicit nil
func (o *ResumeResumeFull) SetMetroNil() {
	o.Metro.Set(nil)
}

// UnsetMetro ensures that no value is present for Metro, not even an explicit nil
func (o *ResumeResumeFull) UnsetMetro() {
	o.Metro.Unset()
}

// GetPaidServices returns the PaidServices field value
func (o *ResumeResumeFull) GetPaidServices() []ResumeObjectsPaidServices {
	if o == nil {
		var ret []ResumeObjectsPaidServices
		return ret
	}

	return o.PaidServices
}

// GetPaidServicesOk returns a tuple with the PaidServices field value
// and a boolean to check if the value has been set.
func (o *ResumeResumeFull) GetPaidServicesOk() ([]ResumeObjectsPaidServices, bool) {
	if o == nil {
		return nil, false
	}
	return o.PaidServices, true
}

// SetPaidServices sets field value
func (o *ResumeResumeFull) SetPaidServices(v []ResumeObjectsPaidServices) {
	o.PaidServices = v
}

// GetProfessionalRoles returns the ProfessionalRoles field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ResumeResumeFull) GetProfessionalRoles() []IncludesIdName {
	if o == nil {
		var ret []IncludesIdName
		return ret
	}
	return o.ProfessionalRoles
}

// GetProfessionalRolesOk returns a tuple with the ProfessionalRoles field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ResumeResumeFull) GetProfessionalRolesOk() ([]IncludesIdName, bool) {
	if o == nil || IsNil(o.ProfessionalRoles) {
		return nil, false
	}
	return o.ProfessionalRoles, true
}

// HasProfessionalRoles returns a boolean if a field has been set.
func (o *ResumeResumeFull) HasProfessionalRoles() bool {
	if o != nil && IsNil(o.ProfessionalRoles) {
		return true
	}

	return false
}

// SetProfessionalRoles gets a reference to the given []IncludesIdName and assigns it to the ProfessionalRoles field.
func (o *ResumeResumeFull) SetProfessionalRoles(v []IncludesIdName) {
	o.ProfessionalRoles = v
}

// GetRecommendation returns the Recommendation field value
func (o *ResumeResumeFull) GetRecommendation() []ResumeObjectsRecommendation {
	if o == nil {
		var ret []ResumeObjectsRecommendation
		return ret
	}

	return o.Recommendation
}

// GetRecommendationOk returns a tuple with the Recommendation field value
// and a boolean to check if the value has been set.
func (o *ResumeResumeFull) GetRecommendationOk() ([]ResumeObjectsRecommendation, bool) {
	if o == nil {
		return nil, false
	}
	return o.Recommendation, true
}

// SetRecommendation sets field value
func (o *ResumeResumeFull) SetRecommendation(v []ResumeObjectsRecommendation) {
	o.Recommendation = v
}

// GetRelocation returns the Relocation field value
func (o *ResumeResumeFull) GetRelocation() ResumeResumeFullAllOfRelocation {
	if o == nil {
		var ret ResumeResumeFullAllOfRelocation
		return ret
	}

	return o.Relocation
}

// GetRelocationOk returns a tuple with the Relocation field value
// and a boolean to check if the value has been set.
func (o *ResumeResumeFull) GetRelocationOk() (*ResumeResumeFullAllOfRelocation, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Relocation, true
}

// SetRelocation sets field value
func (o *ResumeResumeFull) SetRelocation(v ResumeResumeFullAllOfRelocation) {
	o.Relocation = v
}

// GetResumeLocale returns the ResumeLocale field value
func (o *ResumeResumeFull) GetResumeLocale() ResumeResumeFullAllOfResumeLocale {
	if o == nil {
		var ret ResumeResumeFullAllOfResumeLocale
		return ret
	}

	return o.ResumeLocale
}

// GetResumeLocaleOk returns a tuple with the ResumeLocale field value
// and a boolean to check if the value has been set.
func (o *ResumeResumeFull) GetResumeLocaleOk() (*ResumeResumeFullAllOfResumeLocale, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ResumeLocale, true
}

// SetResumeLocale sets field value
func (o *ResumeResumeFull) SetResumeLocale(v ResumeResumeFullAllOfResumeLocale) {
	o.ResumeLocale = v
}

// GetSchedule returns the Schedule field value
// Deprecated
func (o *ResumeResumeFull) GetSchedule() ResumeResumeFullAllOfSchedule {
	if o == nil {
		var ret ResumeResumeFullAllOfSchedule
		return ret
	}

	return o.Schedule
}

// GetScheduleOk returns a tuple with the Schedule field value
// and a boolean to check if the value has been set.
// Deprecated
func (o *ResumeResumeFull) GetScheduleOk() (*ResumeResumeFullAllOfSchedule, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Schedule, true
}

// SetSchedule sets field value
// Deprecated
func (o *ResumeResumeFull) SetSchedule(v ResumeResumeFullAllOfSchedule) {
	o.Schedule = v
}

// GetSchedules returns the Schedules field value
func (o *ResumeResumeFull) GetSchedules() []IncludesIdName {
	if o == nil {
		var ret []IncludesIdName
		return ret
	}

	return o.Schedules
}

// GetSchedulesOk returns a tuple with the Schedules field value
// and a boolean to check if the value has been set.
func (o *ResumeResumeFull) GetSchedulesOk() ([]IncludesIdName, bool) {
	if o == nil {
		return nil, false
	}
	return o.Schedules, true
}

// SetSchedules sets field value
func (o *ResumeResumeFull) SetSchedules(v []IncludesIdName) {
	o.Schedules = v
}

// GetSite returns the Site field value
func (o *ResumeResumeFull) GetSite() []ResumeObjectsSite {
	if o == nil {
		var ret []ResumeObjectsSite
		return ret
	}

	return o.Site
}

// GetSiteOk returns a tuple with the Site field value
// and a boolean to check if the value has been set.
func (o *ResumeResumeFull) GetSiteOk() ([]ResumeObjectsSite, bool) {
	if o == nil {
		return nil, false
	}
	return o.Site, true
}

// SetSite sets field value
func (o *ResumeResumeFull) SetSite(v []ResumeObjectsSite) {
	o.Site = v
}

// GetSkillSet returns the SkillSet field value
func (o *ResumeResumeFull) GetSkillSet() []string {
	if o == nil {
		var ret []string
		return ret
	}

	return o.SkillSet
}

// GetSkillSetOk returns a tuple with the SkillSet field value
// and a boolean to check if the value has been set.
func (o *ResumeResumeFull) GetSkillSetOk() ([]string, bool) {
	if o == nil {
		return nil, false
	}
	return o.SkillSet, true
}

// SetSkillSet sets field value
func (o *ResumeResumeFull) SetSkillSet(v []string) {
	o.SkillSet = v
}

// GetSkills returns the Skills field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ResumeResumeFull) GetSkills() string {
	if o == nil || IsNil(o.Skills.Get()) {
		var ret string
		return ret
	}
	return *o.Skills.Get()
}

// GetSkillsOk returns a tuple with the Skills field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ResumeResumeFull) GetSkillsOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Skills.Get(), o.Skills.IsSet()
}

// HasSkills returns a boolean if a field has been set.
func (o *ResumeResumeFull) HasSkills() bool {
	if o != nil && o.Skills.IsSet() {
		return true
	}

	return false
}

// SetSkills gets a reference to the given NullableString and assigns it to the Skills field.
func (o *ResumeResumeFull) SetSkills(v string) {
	o.Skills.Set(&v)
}
// SetSkillsNil sets the value for Skills to be an explicit nil
func (o *ResumeResumeFull) SetSkillsNil() {
	o.Skills.Set(nil)
}

// UnsetSkills ensures that no value is present for Skills, not even an explicit nil
func (o *ResumeResumeFull) UnsetSkills() {
	o.Skills.Unset()
}

// GetTravelTime returns the TravelTime field value
func (o *ResumeResumeFull) GetTravelTime() ResumeResumeFullAllOfTravelTime {
	if o == nil {
		var ret ResumeResumeFullAllOfTravelTime
		return ret
	}

	return o.TravelTime
}

// GetTravelTimeOk returns a tuple with the TravelTime field value
// and a boolean to check if the value has been set.
func (o *ResumeResumeFull) GetTravelTimeOk() (*ResumeResumeFullAllOfTravelTime, bool) {
	if o == nil {
		return nil, false
	}
	return &o.TravelTime, true
}

// SetTravelTime sets field value
func (o *ResumeResumeFull) SetTravelTime(v ResumeResumeFullAllOfTravelTime) {
	o.TravelTime = v
}

// GetWorkTicket returns the WorkTicket field value
func (o *ResumeResumeFull) GetWorkTicket() []IncludesIdNameUrl {
	if o == nil {
		var ret []IncludesIdNameUrl
		return ret
	}

	return o.WorkTicket
}

// GetWorkTicketOk returns a tuple with the WorkTicket field value
// and a boolean to check if the value has been set.
func (o *ResumeResumeFull) GetWorkTicketOk() ([]IncludesIdNameUrl, bool) {
	if o == nil {
		return nil, false
	}
	return o.WorkTicket, true
}

// SetWorkTicket sets field value
func (o *ResumeResumeFull) SetWorkTicket(v []IncludesIdNameUrl) {
	o.WorkTicket = v
}

func (o ResumeResumeFull) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ResumeResumeFull) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["alternate_url"] = o.AlternateUrl
	toSerialize["id"] = o.Id
	toSerialize["title"] = o.Title.Get()
	if o.Age.IsSet() {
		toSerialize["age"] = o.Age.Get()
	}
	if o.Area.IsSet() {
		toSerialize["area"] = o.Area.Get()
	}
	if o.CanViewFullInfo.IsSet() {
		toSerialize["can_view_full_info"] = o.CanViewFullInfo.Get()
	}
	toSerialize["certificate"] = o.Certificate
	toSerialize["created_at"] = o.CreatedAt
	toSerialize["download"] = o.Download
	toSerialize["education"] = o.Education
	toSerialize["experience"] = o.Experience
	if o.FirstName.IsSet() {
		toSerialize["first_name"] = o.FirstName.Get()
	}
	if o.Gender.IsSet() {
		toSerialize["gender"] = o.Gender.Get()
	}
	toSerialize["hidden_fields"] = o.HiddenFields
	if o.LastName.IsSet() {
		toSerialize["last_name"] = o.LastName.Get()
	}
	if !IsNil(o.Marked) {
		toSerialize["marked"] = o.Marked
	}
	if o.MiddleName.IsSet() {
		toSerialize["middle_name"] = o.MiddleName.Get()
	}
	if !IsNil(o.Platform) {
		toSerialize["platform"] = o.Platform
	}
	if o.Salary.IsSet() {
		toSerialize["salary"] = o.Salary.Get()
	}
	if o.TotalExperience.IsSet() {
		toSerialize["total_experience"] = o.TotalExperience.Get()
	}
	toSerialize["updated_at"] = o.UpdatedAt
	if o.BirthDate.IsSet() {
		toSerialize["birth_date"] = o.BirthDate.Get()
	}
	toSerialize["business_trip_readiness"] = o.BusinessTripReadiness
	toSerialize["citizenship"] = o.Citizenship
	toSerialize["contact"] = o.Contact
	if o.Creds.IsSet() {
		toSerialize["creds"] = o.Creds.Get()
	}
	toSerialize["driver_license_types"] = o.DriverLicenseTypes
	if !IsNil(o.Employment) {
		toSerialize["employment"] = o.Employment
	}
	toSerialize["employments"] = o.Employments
	if o.HasVehicle.IsSet() {
		toSerialize["has_vehicle"] = o.HasVehicle.Get()
	}
	toSerialize["language"] = o.Language
	if o.Metro.IsSet() {
		toSerialize["metro"] = o.Metro.Get()
	}
	toSerialize["paid_services"] = o.PaidServices
	if o.ProfessionalRoles != nil {
		toSerialize["professional_roles"] = o.ProfessionalRoles
	}
	toSerialize["recommendation"] = o.Recommendation
	toSerialize["relocation"] = o.Relocation
	toSerialize["resume_locale"] = o.ResumeLocale
	toSerialize["schedule"] = o.Schedule
	toSerialize["schedules"] = o.Schedules
	toSerialize["site"] = o.Site
	toSerialize["skill_set"] = o.SkillSet
	if o.Skills.IsSet() {
		toSerialize["skills"] = o.Skills.Get()
	}
	toSerialize["travel_time"] = o.TravelTime
	toSerialize["work_ticket"] = o.WorkTicket
	return toSerialize, nil
}

func (o *ResumeResumeFull) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"alternate_url",
		"id",
		"title",
		"certificate",
		"created_at",
		"download",
		"education",
		"experience",
		"hidden_fields",
		"updated_at",
		"business_trip_readiness",
		"citizenship",
		"contact",
		"driver_license_types",
		"employments",
		"language",
		"paid_services",
		"recommendation",
		"relocation",
		"resume_locale",
		"schedule",
		"schedules",
		"site",
		"skill_set",
		"travel_time",
		"work_ticket",
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

	varResumeResumeFull := _ResumeResumeFull{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varResumeResumeFull)

	if err != nil {
		return err
	}

	*o = ResumeResumeFull(varResumeResumeFull)

	return err
}

type NullableResumeResumeFull struct {
	value *ResumeResumeFull
	isSet bool
}

func (v NullableResumeResumeFull) Get() *ResumeResumeFull {
	return v.value
}

func (v *NullableResumeResumeFull) Set(val *ResumeResumeFull) {
	v.value = val
	v.isSet = true
}

func (v NullableResumeResumeFull) IsSet() bool {
	return v.isSet
}

func (v *NullableResumeResumeFull) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableResumeResumeFull(val *ResumeResumeFull) *NullableResumeResumeFull {
	return &NullableResumeResumeFull{value: val, isSet: true}
}

func (v NullableResumeResumeFull) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableResumeResumeFull) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


