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
	"time"
	"bytes"
	"fmt"
)

// checks if the ProfessionalRolesRole type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ProfessionalRolesRole{}

// ProfessionalRolesRole struct for ProfessionalRolesRole
type ProfessionalRolesRole struct {
	// На роль принимаются отклики неполным резюме
	AcceptIncompleteResumes bool `json:"accept_incomplete_resumes"`
	// Идентификатор профессиональной роли
	Id string `json:"id"`
	// Дефолтная роль
	IsDefault bool `json:"is_default"`
	// Имя профессиональной роли
	Name string `json:"name"`
	// Наличие запрета на использование в поиске при составлении поискового запроса
	SearchDeprecated *bool `json:"search_deprecated,omitempty"`
	// Время, с которого действует запрет на использование роли в поиске при составлении поискового запроса,  в формате [ISO 8601](https://en.wikipedia.org/wiki/ISO_8601) с точностью до секунды: `YYYY-MM-DDThh:mm:ss±hhmm` 
	SearchDeprecatedDatetime NullableTime `json:"search_deprecated_datetime,omitempty"`
	// Наличие запрета на использование при создании новых сущностей (резюме или вакансии)
	SelectDeprecated *bool `json:"select_deprecated,omitempty"`
	// Время, с которого действует запрет на использование роли при создании новых сущностей,  в формате [ISO 8601](https://en.wikipedia.org/wiki/ISO_8601) с точностью до секунды: `YYYY-MM-DDThh:mm:ss±hhmm` 
	SelectDeprecatedDatetime NullableTime `json:"select_deprecated_datetime,omitempty"`
}

type _ProfessionalRolesRole ProfessionalRolesRole

// NewProfessionalRolesRole instantiates a new ProfessionalRolesRole object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewProfessionalRolesRole(acceptIncompleteResumes bool, id string, isDefault bool, name string) *ProfessionalRolesRole {
	this := ProfessionalRolesRole{}
	this.AcceptIncompleteResumes = acceptIncompleteResumes
	this.Id = id
	this.IsDefault = isDefault
	this.Name = name
	return &this
}

// NewProfessionalRolesRoleWithDefaults instantiates a new ProfessionalRolesRole object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewProfessionalRolesRoleWithDefaults() *ProfessionalRolesRole {
	this := ProfessionalRolesRole{}
	return &this
}

// GetAcceptIncompleteResumes returns the AcceptIncompleteResumes field value
func (o *ProfessionalRolesRole) GetAcceptIncompleteResumes() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.AcceptIncompleteResumes
}

// GetAcceptIncompleteResumesOk returns a tuple with the AcceptIncompleteResumes field value
// and a boolean to check if the value has been set.
func (o *ProfessionalRolesRole) GetAcceptIncompleteResumesOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AcceptIncompleteResumes, true
}

// SetAcceptIncompleteResumes sets field value
func (o *ProfessionalRolesRole) SetAcceptIncompleteResumes(v bool) {
	o.AcceptIncompleteResumes = v
}

// GetId returns the Id field value
func (o *ProfessionalRolesRole) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *ProfessionalRolesRole) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *ProfessionalRolesRole) SetId(v string) {
	o.Id = v
}

// GetIsDefault returns the IsDefault field value
func (o *ProfessionalRolesRole) GetIsDefault() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.IsDefault
}

// GetIsDefaultOk returns a tuple with the IsDefault field value
// and a boolean to check if the value has been set.
func (o *ProfessionalRolesRole) GetIsDefaultOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.IsDefault, true
}

// SetIsDefault sets field value
func (o *ProfessionalRolesRole) SetIsDefault(v bool) {
	o.IsDefault = v
}

// GetName returns the Name field value
func (o *ProfessionalRolesRole) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *ProfessionalRolesRole) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *ProfessionalRolesRole) SetName(v string) {
	o.Name = v
}

// GetSearchDeprecated returns the SearchDeprecated field value if set, zero value otherwise.
func (o *ProfessionalRolesRole) GetSearchDeprecated() bool {
	if o == nil || IsNil(o.SearchDeprecated) {
		var ret bool
		return ret
	}
	return *o.SearchDeprecated
}

// GetSearchDeprecatedOk returns a tuple with the SearchDeprecated field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProfessionalRolesRole) GetSearchDeprecatedOk() (*bool, bool) {
	if o == nil || IsNil(o.SearchDeprecated) {
		return nil, false
	}
	return o.SearchDeprecated, true
}

// HasSearchDeprecated returns a boolean if a field has been set.
func (o *ProfessionalRolesRole) HasSearchDeprecated() bool {
	if o != nil && !IsNil(o.SearchDeprecated) {
		return true
	}

	return false
}

// SetSearchDeprecated gets a reference to the given bool and assigns it to the SearchDeprecated field.
func (o *ProfessionalRolesRole) SetSearchDeprecated(v bool) {
	o.SearchDeprecated = &v
}

// GetSearchDeprecatedDatetime returns the SearchDeprecatedDatetime field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ProfessionalRolesRole) GetSearchDeprecatedDatetime() time.Time {
	if o == nil || IsNil(o.SearchDeprecatedDatetime.Get()) {
		var ret time.Time
		return ret
	}
	return *o.SearchDeprecatedDatetime.Get()
}

// GetSearchDeprecatedDatetimeOk returns a tuple with the SearchDeprecatedDatetime field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ProfessionalRolesRole) GetSearchDeprecatedDatetimeOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return o.SearchDeprecatedDatetime.Get(), o.SearchDeprecatedDatetime.IsSet()
}

// HasSearchDeprecatedDatetime returns a boolean if a field has been set.
func (o *ProfessionalRolesRole) HasSearchDeprecatedDatetime() bool {
	if o != nil && o.SearchDeprecatedDatetime.IsSet() {
		return true
	}

	return false
}

// SetSearchDeprecatedDatetime gets a reference to the given NullableTime and assigns it to the SearchDeprecatedDatetime field.
func (o *ProfessionalRolesRole) SetSearchDeprecatedDatetime(v time.Time) {
	o.SearchDeprecatedDatetime.Set(&v)
}
// SetSearchDeprecatedDatetimeNil sets the value for SearchDeprecatedDatetime to be an explicit nil
func (o *ProfessionalRolesRole) SetSearchDeprecatedDatetimeNil() {
	o.SearchDeprecatedDatetime.Set(nil)
}

// UnsetSearchDeprecatedDatetime ensures that no value is present for SearchDeprecatedDatetime, not even an explicit nil
func (o *ProfessionalRolesRole) UnsetSearchDeprecatedDatetime() {
	o.SearchDeprecatedDatetime.Unset()
}

// GetSelectDeprecated returns the SelectDeprecated field value if set, zero value otherwise.
func (o *ProfessionalRolesRole) GetSelectDeprecated() bool {
	if o == nil || IsNil(o.SelectDeprecated) {
		var ret bool
		return ret
	}
	return *o.SelectDeprecated
}

// GetSelectDeprecatedOk returns a tuple with the SelectDeprecated field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProfessionalRolesRole) GetSelectDeprecatedOk() (*bool, bool) {
	if o == nil || IsNil(o.SelectDeprecated) {
		return nil, false
	}
	return o.SelectDeprecated, true
}

// HasSelectDeprecated returns a boolean if a field has been set.
func (o *ProfessionalRolesRole) HasSelectDeprecated() bool {
	if o != nil && !IsNil(o.SelectDeprecated) {
		return true
	}

	return false
}

// SetSelectDeprecated gets a reference to the given bool and assigns it to the SelectDeprecated field.
func (o *ProfessionalRolesRole) SetSelectDeprecated(v bool) {
	o.SelectDeprecated = &v
}

// GetSelectDeprecatedDatetime returns the SelectDeprecatedDatetime field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ProfessionalRolesRole) GetSelectDeprecatedDatetime() time.Time {
	if o == nil || IsNil(o.SelectDeprecatedDatetime.Get()) {
		var ret time.Time
		return ret
	}
	return *o.SelectDeprecatedDatetime.Get()
}

// GetSelectDeprecatedDatetimeOk returns a tuple with the SelectDeprecatedDatetime field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ProfessionalRolesRole) GetSelectDeprecatedDatetimeOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return o.SelectDeprecatedDatetime.Get(), o.SelectDeprecatedDatetime.IsSet()
}

// HasSelectDeprecatedDatetime returns a boolean if a field has been set.
func (o *ProfessionalRolesRole) HasSelectDeprecatedDatetime() bool {
	if o != nil && o.SelectDeprecatedDatetime.IsSet() {
		return true
	}

	return false
}

// SetSelectDeprecatedDatetime gets a reference to the given NullableTime and assigns it to the SelectDeprecatedDatetime field.
func (o *ProfessionalRolesRole) SetSelectDeprecatedDatetime(v time.Time) {
	o.SelectDeprecatedDatetime.Set(&v)
}
// SetSelectDeprecatedDatetimeNil sets the value for SelectDeprecatedDatetime to be an explicit nil
func (o *ProfessionalRolesRole) SetSelectDeprecatedDatetimeNil() {
	o.SelectDeprecatedDatetime.Set(nil)
}

// UnsetSelectDeprecatedDatetime ensures that no value is present for SelectDeprecatedDatetime, not even an explicit nil
func (o *ProfessionalRolesRole) UnsetSelectDeprecatedDatetime() {
	o.SelectDeprecatedDatetime.Unset()
}

func (o ProfessionalRolesRole) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ProfessionalRolesRole) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["accept_incomplete_resumes"] = o.AcceptIncompleteResumes
	toSerialize["id"] = o.Id
	toSerialize["is_default"] = o.IsDefault
	toSerialize["name"] = o.Name
	if !IsNil(o.SearchDeprecated) {
		toSerialize["search_deprecated"] = o.SearchDeprecated
	}
	if o.SearchDeprecatedDatetime.IsSet() {
		toSerialize["search_deprecated_datetime"] = o.SearchDeprecatedDatetime.Get()
	}
	if !IsNil(o.SelectDeprecated) {
		toSerialize["select_deprecated"] = o.SelectDeprecated
	}
	if o.SelectDeprecatedDatetime.IsSet() {
		toSerialize["select_deprecated_datetime"] = o.SelectDeprecatedDatetime.Get()
	}
	return toSerialize, nil
}

func (o *ProfessionalRolesRole) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"accept_incomplete_resumes",
		"id",
		"is_default",
		"name",
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

	varProfessionalRolesRole := _ProfessionalRolesRole{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varProfessionalRolesRole)

	if err != nil {
		return err
	}

	*o = ProfessionalRolesRole(varProfessionalRolesRole)

	return err
}

type NullableProfessionalRolesRole struct {
	value *ProfessionalRolesRole
	isSet bool
}

func (v NullableProfessionalRolesRole) Get() *ProfessionalRolesRole {
	return v.value
}

func (v *NullableProfessionalRolesRole) Set(val *ProfessionalRolesRole) {
	v.value = val
	v.isSet = true
}

func (v NullableProfessionalRolesRole) IsSet() bool {
	return v.isSet
}

func (v *NullableProfessionalRolesRole) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableProfessionalRolesRole(val *ProfessionalRolesRole) *NullableProfessionalRolesRole {
	return &NullableProfessionalRolesRole{value: val, isSet: true}
}

func (v NullableProfessionalRolesRole) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableProfessionalRolesRole) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


