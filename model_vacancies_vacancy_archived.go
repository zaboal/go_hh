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

// checks if the VacanciesVacancyArchived type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &VacanciesVacancyArchived{}

// VacanciesVacancyArchived struct for VacanciesVacancyArchived
type VacanciesVacancyArchived struct {
	Address NullableVacancyAddressRawOutput `json:"address,omitempty"`
	// Ссылка на представление вакансии на сайте
	AlternateUrl string `json:"alternate_url"`
	// Ссылка на отклик на вакансию на сайте
	ApplyAlternateUrl string `json:"apply_alternate_url"`
	// Находится ли данная вакансия в архиве
	Archived bool `json:"archived"`
	Area IncludesArea `json:"area"`
	Department NullableVacancyDepartmentOutput `json:"department"`
	Employer NullableVacanciesEmployerPublic `json:"employer"`
	// Информация о наличии прикрепленного тестового задании к вакансии
	HasTest bool `json:"has_test"`
	// Идентификатор вакансии
	Id string `json:"id"`
	// Название вакансии
	Name string `json:"name"`
	// Является ли данная вакансия премиум-вакансией
	Premium bool `json:"premium"`
	// Дата и время публикации вакансии
	PublishedAt string `json:"published_at"`
	// Возвращает связи соискателя с вакансией. Значения из поля `vacancy_relation` в [справочнике полей](#tag/Obshie-spravochniki/operation/get-dictionaries)
	Relations []VacancyRelationItem `json:"relations"`
	// Обязательно ли заполнять сообщение при отклике на вакансию
	ResponseLetterRequired bool `json:"response_letter_required"`
	// URL отклика для прямых вакансий (`type.id=direct`)
	ResponseUrl NullableString `json:"response_url,omitempty"`
	Salary NullableVacancySalary `json:"salary"`
	// Отображать ли лого для вакансии в поисковой выдаче
	ShowLogoInSearch NullableBool `json:"show_logo_in_search,omitempty"`
	Type VacancyTypeOutput `json:"type"`
	// URL вакансии
	Url string `json:"url"`
	// Дата и время архивации вакансии
	ArchivedAt string `json:"archived_at"`
	Counters VacancyCountersForArchivedOrHidden `json:"counters"`
	// Дата и время публикации вакансии
	CreatedAt string `json:"created_at"`
	// Расстояние в метрах между центром сортировки (заданной параметрами `sort_point_lat`, `sort_point_lng`) и указанным в вакансии адресом. В случае, если в адресе указаны только станции метро, выдается расстояние между центром сортировки и средней геометрической точкой указанных станций.  Значение `sort_point_distance` выдается только в случае, если заданы параметры `sort_point_lat`, `sort_point_lng`, `order_by=distance` 
	SortPointDistance NullableFloat32 `json:"sort_point_distance,omitempty"`
}

type _VacanciesVacancyArchived VacanciesVacancyArchived

// NewVacanciesVacancyArchived instantiates a new VacanciesVacancyArchived object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewVacanciesVacancyArchived(alternateUrl string, applyAlternateUrl string, archived bool, area IncludesArea, department NullableVacancyDepartmentOutput, employer NullableVacanciesEmployerPublic, hasTest bool, id string, name string, premium bool, publishedAt string, relations []VacancyRelationItem, responseLetterRequired bool, salary NullableVacancySalary, type_ VacancyTypeOutput, url string, archivedAt string, counters VacancyCountersForArchivedOrHidden, createdAt string) *VacanciesVacancyArchived {
	this := VacanciesVacancyArchived{}
	this.AlternateUrl = alternateUrl
	this.ApplyAlternateUrl = applyAlternateUrl
	this.Archived = archived
	this.Area = area
	this.Department = department
	this.Employer = employer
	this.HasTest = hasTest
	this.Id = id
	this.Name = name
	this.Premium = premium
	this.PublishedAt = publishedAt
	this.Relations = relations
	this.ResponseLetterRequired = responseLetterRequired
	this.Salary = salary
	this.Type = type_
	this.Url = url
	this.ArchivedAt = archivedAt
	this.Counters = counters
	this.CreatedAt = createdAt
	return &this
}

// NewVacanciesVacancyArchivedWithDefaults instantiates a new VacanciesVacancyArchived object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewVacanciesVacancyArchivedWithDefaults() *VacanciesVacancyArchived {
	this := VacanciesVacancyArchived{}
	return &this
}

// GetAddress returns the Address field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *VacanciesVacancyArchived) GetAddress() VacancyAddressRawOutput {
	if o == nil || IsNil(o.Address.Get()) {
		var ret VacancyAddressRawOutput
		return ret
	}
	return *o.Address.Get()
}

// GetAddressOk returns a tuple with the Address field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *VacanciesVacancyArchived) GetAddressOk() (*VacancyAddressRawOutput, bool) {
	if o == nil {
		return nil, false
	}
	return o.Address.Get(), o.Address.IsSet()
}

// HasAddress returns a boolean if a field has been set.
func (o *VacanciesVacancyArchived) HasAddress() bool {
	if o != nil && o.Address.IsSet() {
		return true
	}

	return false
}

// SetAddress gets a reference to the given NullableVacancyAddressRawOutput and assigns it to the Address field.
func (o *VacanciesVacancyArchived) SetAddress(v VacancyAddressRawOutput) {
	o.Address.Set(&v)
}
// SetAddressNil sets the value for Address to be an explicit nil
func (o *VacanciesVacancyArchived) SetAddressNil() {
	o.Address.Set(nil)
}

// UnsetAddress ensures that no value is present for Address, not even an explicit nil
func (o *VacanciesVacancyArchived) UnsetAddress() {
	o.Address.Unset()
}

// GetAlternateUrl returns the AlternateUrl field value
func (o *VacanciesVacancyArchived) GetAlternateUrl() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.AlternateUrl
}

// GetAlternateUrlOk returns a tuple with the AlternateUrl field value
// and a boolean to check if the value has been set.
func (o *VacanciesVacancyArchived) GetAlternateUrlOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AlternateUrl, true
}

// SetAlternateUrl sets field value
func (o *VacanciesVacancyArchived) SetAlternateUrl(v string) {
	o.AlternateUrl = v
}

// GetApplyAlternateUrl returns the ApplyAlternateUrl field value
func (o *VacanciesVacancyArchived) GetApplyAlternateUrl() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ApplyAlternateUrl
}

// GetApplyAlternateUrlOk returns a tuple with the ApplyAlternateUrl field value
// and a boolean to check if the value has been set.
func (o *VacanciesVacancyArchived) GetApplyAlternateUrlOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ApplyAlternateUrl, true
}

// SetApplyAlternateUrl sets field value
func (o *VacanciesVacancyArchived) SetApplyAlternateUrl(v string) {
	o.ApplyAlternateUrl = v
}

// GetArchived returns the Archived field value
func (o *VacanciesVacancyArchived) GetArchived() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.Archived
}

// GetArchivedOk returns a tuple with the Archived field value
// and a boolean to check if the value has been set.
func (o *VacanciesVacancyArchived) GetArchivedOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Archived, true
}

// SetArchived sets field value
func (o *VacanciesVacancyArchived) SetArchived(v bool) {
	o.Archived = v
}

// GetArea returns the Area field value
func (o *VacanciesVacancyArchived) GetArea() IncludesArea {
	if o == nil {
		var ret IncludesArea
		return ret
	}

	return o.Area
}

// GetAreaOk returns a tuple with the Area field value
// and a boolean to check if the value has been set.
func (o *VacanciesVacancyArchived) GetAreaOk() (*IncludesArea, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Area, true
}

// SetArea sets field value
func (o *VacanciesVacancyArchived) SetArea(v IncludesArea) {
	o.Area = v
}

// GetDepartment returns the Department field value
// If the value is explicit nil, the zero value for VacancyDepartmentOutput will be returned
func (o *VacanciesVacancyArchived) GetDepartment() VacancyDepartmentOutput {
	if o == nil || o.Department.Get() == nil {
		var ret VacancyDepartmentOutput
		return ret
	}

	return *o.Department.Get()
}

// GetDepartmentOk returns a tuple with the Department field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *VacanciesVacancyArchived) GetDepartmentOk() (*VacancyDepartmentOutput, bool) {
	if o == nil {
		return nil, false
	}
	return o.Department.Get(), o.Department.IsSet()
}

// SetDepartment sets field value
func (o *VacanciesVacancyArchived) SetDepartment(v VacancyDepartmentOutput) {
	o.Department.Set(&v)
}

// GetEmployer returns the Employer field value
// If the value is explicit nil, the zero value for VacanciesEmployerPublic will be returned
func (o *VacanciesVacancyArchived) GetEmployer() VacanciesEmployerPublic {
	if o == nil || o.Employer.Get() == nil {
		var ret VacanciesEmployerPublic
		return ret
	}

	return *o.Employer.Get()
}

// GetEmployerOk returns a tuple with the Employer field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *VacanciesVacancyArchived) GetEmployerOk() (*VacanciesEmployerPublic, bool) {
	if o == nil {
		return nil, false
	}
	return o.Employer.Get(), o.Employer.IsSet()
}

// SetEmployer sets field value
func (o *VacanciesVacancyArchived) SetEmployer(v VacanciesEmployerPublic) {
	o.Employer.Set(&v)
}

// GetHasTest returns the HasTest field value
func (o *VacanciesVacancyArchived) GetHasTest() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.HasTest
}

// GetHasTestOk returns a tuple with the HasTest field value
// and a boolean to check if the value has been set.
func (o *VacanciesVacancyArchived) GetHasTestOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.HasTest, true
}

// SetHasTest sets field value
func (o *VacanciesVacancyArchived) SetHasTest(v bool) {
	o.HasTest = v
}

// GetId returns the Id field value
func (o *VacanciesVacancyArchived) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *VacanciesVacancyArchived) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *VacanciesVacancyArchived) SetId(v string) {
	o.Id = v
}

// GetName returns the Name field value
func (o *VacanciesVacancyArchived) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *VacanciesVacancyArchived) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *VacanciesVacancyArchived) SetName(v string) {
	o.Name = v
}

// GetPremium returns the Premium field value
func (o *VacanciesVacancyArchived) GetPremium() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.Premium
}

// GetPremiumOk returns a tuple with the Premium field value
// and a boolean to check if the value has been set.
func (o *VacanciesVacancyArchived) GetPremiumOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Premium, true
}

// SetPremium sets field value
func (o *VacanciesVacancyArchived) SetPremium(v bool) {
	o.Premium = v
}

// GetPublishedAt returns the PublishedAt field value
func (o *VacanciesVacancyArchived) GetPublishedAt() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.PublishedAt
}

// GetPublishedAtOk returns a tuple with the PublishedAt field value
// and a boolean to check if the value has been set.
func (o *VacanciesVacancyArchived) GetPublishedAtOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.PublishedAt, true
}

// SetPublishedAt sets field value
func (o *VacanciesVacancyArchived) SetPublishedAt(v string) {
	o.PublishedAt = v
}

// GetRelations returns the Relations field value
func (o *VacanciesVacancyArchived) GetRelations() []VacancyRelationItem {
	if o == nil {
		var ret []VacancyRelationItem
		return ret
	}

	return o.Relations
}

// GetRelationsOk returns a tuple with the Relations field value
// and a boolean to check if the value has been set.
func (o *VacanciesVacancyArchived) GetRelationsOk() ([]VacancyRelationItem, bool) {
	if o == nil {
		return nil, false
	}
	return o.Relations, true
}

// SetRelations sets field value
func (o *VacanciesVacancyArchived) SetRelations(v []VacancyRelationItem) {
	o.Relations = v
}

// GetResponseLetterRequired returns the ResponseLetterRequired field value
func (o *VacanciesVacancyArchived) GetResponseLetterRequired() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.ResponseLetterRequired
}

// GetResponseLetterRequiredOk returns a tuple with the ResponseLetterRequired field value
// and a boolean to check if the value has been set.
func (o *VacanciesVacancyArchived) GetResponseLetterRequiredOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ResponseLetterRequired, true
}

// SetResponseLetterRequired sets field value
func (o *VacanciesVacancyArchived) SetResponseLetterRequired(v bool) {
	o.ResponseLetterRequired = v
}

// GetResponseUrl returns the ResponseUrl field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *VacanciesVacancyArchived) GetResponseUrl() string {
	if o == nil || IsNil(o.ResponseUrl.Get()) {
		var ret string
		return ret
	}
	return *o.ResponseUrl.Get()
}

// GetResponseUrlOk returns a tuple with the ResponseUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *VacanciesVacancyArchived) GetResponseUrlOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.ResponseUrl.Get(), o.ResponseUrl.IsSet()
}

// HasResponseUrl returns a boolean if a field has been set.
func (o *VacanciesVacancyArchived) HasResponseUrl() bool {
	if o != nil && o.ResponseUrl.IsSet() {
		return true
	}

	return false
}

// SetResponseUrl gets a reference to the given NullableString and assigns it to the ResponseUrl field.
func (o *VacanciesVacancyArchived) SetResponseUrl(v string) {
	o.ResponseUrl.Set(&v)
}
// SetResponseUrlNil sets the value for ResponseUrl to be an explicit nil
func (o *VacanciesVacancyArchived) SetResponseUrlNil() {
	o.ResponseUrl.Set(nil)
}

// UnsetResponseUrl ensures that no value is present for ResponseUrl, not even an explicit nil
func (o *VacanciesVacancyArchived) UnsetResponseUrl() {
	o.ResponseUrl.Unset()
}

// GetSalary returns the Salary field value
// If the value is explicit nil, the zero value for VacancySalary will be returned
func (o *VacanciesVacancyArchived) GetSalary() VacancySalary {
	if o == nil || o.Salary.Get() == nil {
		var ret VacancySalary
		return ret
	}

	return *o.Salary.Get()
}

// GetSalaryOk returns a tuple with the Salary field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *VacanciesVacancyArchived) GetSalaryOk() (*VacancySalary, bool) {
	if o == nil {
		return nil, false
	}
	return o.Salary.Get(), o.Salary.IsSet()
}

// SetSalary sets field value
func (o *VacanciesVacancyArchived) SetSalary(v VacancySalary) {
	o.Salary.Set(&v)
}

// GetShowLogoInSearch returns the ShowLogoInSearch field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *VacanciesVacancyArchived) GetShowLogoInSearch() bool {
	if o == nil || IsNil(o.ShowLogoInSearch.Get()) {
		var ret bool
		return ret
	}
	return *o.ShowLogoInSearch.Get()
}

// GetShowLogoInSearchOk returns a tuple with the ShowLogoInSearch field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *VacanciesVacancyArchived) GetShowLogoInSearchOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return o.ShowLogoInSearch.Get(), o.ShowLogoInSearch.IsSet()
}

// HasShowLogoInSearch returns a boolean if a field has been set.
func (o *VacanciesVacancyArchived) HasShowLogoInSearch() bool {
	if o != nil && o.ShowLogoInSearch.IsSet() {
		return true
	}

	return false
}

// SetShowLogoInSearch gets a reference to the given NullableBool and assigns it to the ShowLogoInSearch field.
func (o *VacanciesVacancyArchived) SetShowLogoInSearch(v bool) {
	o.ShowLogoInSearch.Set(&v)
}
// SetShowLogoInSearchNil sets the value for ShowLogoInSearch to be an explicit nil
func (o *VacanciesVacancyArchived) SetShowLogoInSearchNil() {
	o.ShowLogoInSearch.Set(nil)
}

// UnsetShowLogoInSearch ensures that no value is present for ShowLogoInSearch, not even an explicit nil
func (o *VacanciesVacancyArchived) UnsetShowLogoInSearch() {
	o.ShowLogoInSearch.Unset()
}

// GetType returns the Type field value
func (o *VacanciesVacancyArchived) GetType() VacancyTypeOutput {
	if o == nil {
		var ret VacancyTypeOutput
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *VacanciesVacancyArchived) GetTypeOk() (*VacancyTypeOutput, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *VacanciesVacancyArchived) SetType(v VacancyTypeOutput) {
	o.Type = v
}

// GetUrl returns the Url field value
func (o *VacanciesVacancyArchived) GetUrl() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Url
}

// GetUrlOk returns a tuple with the Url field value
// and a boolean to check if the value has been set.
func (o *VacanciesVacancyArchived) GetUrlOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Url, true
}

// SetUrl sets field value
func (o *VacanciesVacancyArchived) SetUrl(v string) {
	o.Url = v
}

// GetArchivedAt returns the ArchivedAt field value
func (o *VacanciesVacancyArchived) GetArchivedAt() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ArchivedAt
}

// GetArchivedAtOk returns a tuple with the ArchivedAt field value
// and a boolean to check if the value has been set.
func (o *VacanciesVacancyArchived) GetArchivedAtOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ArchivedAt, true
}

// SetArchivedAt sets field value
func (o *VacanciesVacancyArchived) SetArchivedAt(v string) {
	o.ArchivedAt = v
}

// GetCounters returns the Counters field value
func (o *VacanciesVacancyArchived) GetCounters() VacancyCountersForArchivedOrHidden {
	if o == nil {
		var ret VacancyCountersForArchivedOrHidden
		return ret
	}

	return o.Counters
}

// GetCountersOk returns a tuple with the Counters field value
// and a boolean to check if the value has been set.
func (o *VacanciesVacancyArchived) GetCountersOk() (*VacancyCountersForArchivedOrHidden, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Counters, true
}

// SetCounters sets field value
func (o *VacanciesVacancyArchived) SetCounters(v VacancyCountersForArchivedOrHidden) {
	o.Counters = v
}

// GetCreatedAt returns the CreatedAt field value
func (o *VacanciesVacancyArchived) GetCreatedAt() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value
// and a boolean to check if the value has been set.
func (o *VacanciesVacancyArchived) GetCreatedAtOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CreatedAt, true
}

// SetCreatedAt sets field value
func (o *VacanciesVacancyArchived) SetCreatedAt(v string) {
	o.CreatedAt = v
}

// GetSortPointDistance returns the SortPointDistance field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *VacanciesVacancyArchived) GetSortPointDistance() float32 {
	if o == nil || IsNil(o.SortPointDistance.Get()) {
		var ret float32
		return ret
	}
	return *o.SortPointDistance.Get()
}

// GetSortPointDistanceOk returns a tuple with the SortPointDistance field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *VacanciesVacancyArchived) GetSortPointDistanceOk() (*float32, bool) {
	if o == nil {
		return nil, false
	}
	return o.SortPointDistance.Get(), o.SortPointDistance.IsSet()
}

// HasSortPointDistance returns a boolean if a field has been set.
func (o *VacanciesVacancyArchived) HasSortPointDistance() bool {
	if o != nil && o.SortPointDistance.IsSet() {
		return true
	}

	return false
}

// SetSortPointDistance gets a reference to the given NullableFloat32 and assigns it to the SortPointDistance field.
func (o *VacanciesVacancyArchived) SetSortPointDistance(v float32) {
	o.SortPointDistance.Set(&v)
}
// SetSortPointDistanceNil sets the value for SortPointDistance to be an explicit nil
func (o *VacanciesVacancyArchived) SetSortPointDistanceNil() {
	o.SortPointDistance.Set(nil)
}

// UnsetSortPointDistance ensures that no value is present for SortPointDistance, not even an explicit nil
func (o *VacanciesVacancyArchived) UnsetSortPointDistance() {
	o.SortPointDistance.Unset()
}

func (o VacanciesVacancyArchived) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o VacanciesVacancyArchived) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.Address.IsSet() {
		toSerialize["address"] = o.Address.Get()
	}
	toSerialize["alternate_url"] = o.AlternateUrl
	toSerialize["apply_alternate_url"] = o.ApplyAlternateUrl
	toSerialize["archived"] = o.Archived
	toSerialize["area"] = o.Area
	toSerialize["department"] = o.Department.Get()
	toSerialize["employer"] = o.Employer.Get()
	toSerialize["has_test"] = o.HasTest
	toSerialize["id"] = o.Id
	toSerialize["name"] = o.Name
	toSerialize["premium"] = o.Premium
	toSerialize["published_at"] = o.PublishedAt
	toSerialize["relations"] = o.Relations
	toSerialize["response_letter_required"] = o.ResponseLetterRequired
	if o.ResponseUrl.IsSet() {
		toSerialize["response_url"] = o.ResponseUrl.Get()
	}
	toSerialize["salary"] = o.Salary.Get()
	if o.ShowLogoInSearch.IsSet() {
		toSerialize["show_logo_in_search"] = o.ShowLogoInSearch.Get()
	}
	toSerialize["type"] = o.Type
	toSerialize["url"] = o.Url
	toSerialize["archived_at"] = o.ArchivedAt
	toSerialize["counters"] = o.Counters
	toSerialize["created_at"] = o.CreatedAt
	if o.SortPointDistance.IsSet() {
		toSerialize["sort_point_distance"] = o.SortPointDistance.Get()
	}
	return toSerialize, nil
}

func (o *VacanciesVacancyArchived) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"alternate_url",
		"apply_alternate_url",
		"archived",
		"area",
		"department",
		"employer",
		"has_test",
		"id",
		"name",
		"premium",
		"published_at",
		"relations",
		"response_letter_required",
		"salary",
		"type",
		"url",
		"archived_at",
		"counters",
		"created_at",
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

	varVacanciesVacancyArchived := _VacanciesVacancyArchived{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varVacanciesVacancyArchived)

	if err != nil {
		return err
	}

	*o = VacanciesVacancyArchived(varVacanciesVacancyArchived)

	return err
}

type NullableVacanciesVacancyArchived struct {
	value *VacanciesVacancyArchived
	isSet bool
}

func (v NullableVacanciesVacancyArchived) Get() *VacanciesVacancyArchived {
	return v.value
}

func (v *NullableVacanciesVacancyArchived) Set(val *VacanciesVacancyArchived) {
	v.value = val
	v.isSet = true
}

func (v NullableVacanciesVacancyArchived) IsSet() bool {
	return v.isSet
}

func (v *NullableVacanciesVacancyArchived) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableVacanciesVacancyArchived(val *VacanciesVacancyArchived) *NullableVacanciesVacancyArchived {
	return &NullableVacanciesVacancyArchived{value: val, isSet: true}
}

func (v NullableVacanciesVacancyArchived) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableVacanciesVacancyArchived) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


