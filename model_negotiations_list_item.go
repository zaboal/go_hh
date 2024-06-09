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

// checks if the NegotiationsListItem type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &NegotiationsListItem{}

// NegotiationsListItem struct for NegotiationsListItem
type NegotiationsListItem struct {
	Counters *NegotiationsObjectsCounters `json:"counters,omitempty"`
	// Дата и время создания отклика/приглашения
	CreatedAt string `json:"created_at"`
	// Есть ли в откликах/приглашениях по данной вакансии обновления, требующие внимания
	HasUpdates bool `json:"has_updates"`
	// Идентификатор отклика/приглашения
	Id string `json:"id"`
	// URL, на который необходимо делать GET-запрос для получения [списка сообщений в отклике/приглашении](#tag/Perepiska-(otklikipriglasheniya)-dlya-soiskatelya/operation/get-negotiation-messages). Если `can_edit` равно `false`, значение поля должно игнорироваться
	// Deprecated
	MessagesUrl *string `json:"messages_url,omitempty"`
	// Текущий статус переписки.  Возможные значения приведены в поле `messaging_status` [справочника полей](#tag/Obshie-spravochniki/operation/get-dictionaries) 
	MessagingStatus string `json:"messaging_status"`
	// Список профессиональных ролей
	ProfessionalRoles []VacancyProfessionalRoleItem `json:"professional_roles,omitempty"`
	// Источник отклика/приглашения
	Source *string `json:"source,omitempty"`
	// Текущее состояние отклика/приглашения.  Возможные значения приведены в поле `negotiations_state` [справочника полей](#tag/Obshie-spravochniki/operation/get-dictionaries) 
	State IncludesIdName `json:"state"`
	// Дата и время последнего обновления отклика/приглашения
	UpdatedAt string `json:"updated_at"`
	// Был ли отклик просмотрен работодателем
	ViewedByOpponent bool `json:"viewed_by_opponent"`
	// Можно ли [скрыть отклик](#tag/Perepiska-(otklikipriglasheniya)-dlya-soiskatelya/operation/hide-active-response) вместе с сообщением работодателю об отказе
	DeclineAllowed bool `json:"decline_allowed"`
	// Скрыт ли текущий отклик от соискателя
	Hidden bool `json:"hidden"`
	JobSearchStatus *IncludesIdName `json:"job_search_status,omitempty"`
	PhoneCalls *NegotiationsPhoneCalls `json:"phone_calls,omitempty"`
	Vacancy *VacanciesNegotiationsVacancyShort `json:"vacancy,omitempty"`
	Resume NullableResumeResumeNanoWithUrl `json:"resume,omitempty"`
	// Ссылка на полную версию отклика
	Url string `json:"url"`
}

type _NegotiationsListItem NegotiationsListItem

// NewNegotiationsListItem instantiates a new NegotiationsListItem object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewNegotiationsListItem(createdAt string, hasUpdates bool, id string, messagingStatus string, state IncludesIdName, updatedAt string, viewedByOpponent bool, declineAllowed bool, hidden bool, url string) *NegotiationsListItem {
	this := NegotiationsListItem{}
	this.CreatedAt = createdAt
	this.HasUpdates = hasUpdates
	this.Id = id
	this.MessagingStatus = messagingStatus
	this.State = state
	this.UpdatedAt = updatedAt
	this.ViewedByOpponent = viewedByOpponent
	this.DeclineAllowed = declineAllowed
	this.Hidden = hidden
	this.Url = url
	return &this
}

// NewNegotiationsListItemWithDefaults instantiates a new NegotiationsListItem object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewNegotiationsListItemWithDefaults() *NegotiationsListItem {
	this := NegotiationsListItem{}
	return &this
}

// GetCounters returns the Counters field value if set, zero value otherwise.
func (o *NegotiationsListItem) GetCounters() NegotiationsObjectsCounters {
	if o == nil || IsNil(o.Counters) {
		var ret NegotiationsObjectsCounters
		return ret
	}
	return *o.Counters
}

// GetCountersOk returns a tuple with the Counters field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NegotiationsListItem) GetCountersOk() (*NegotiationsObjectsCounters, bool) {
	if o == nil || IsNil(o.Counters) {
		return nil, false
	}
	return o.Counters, true
}

// HasCounters returns a boolean if a field has been set.
func (o *NegotiationsListItem) HasCounters() bool {
	if o != nil && !IsNil(o.Counters) {
		return true
	}

	return false
}

// SetCounters gets a reference to the given NegotiationsObjectsCounters and assigns it to the Counters field.
func (o *NegotiationsListItem) SetCounters(v NegotiationsObjectsCounters) {
	o.Counters = &v
}

// GetCreatedAt returns the CreatedAt field value
func (o *NegotiationsListItem) GetCreatedAt() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value
// and a boolean to check if the value has been set.
func (o *NegotiationsListItem) GetCreatedAtOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CreatedAt, true
}

// SetCreatedAt sets field value
func (o *NegotiationsListItem) SetCreatedAt(v string) {
	o.CreatedAt = v
}

// GetHasUpdates returns the HasUpdates field value
func (o *NegotiationsListItem) GetHasUpdates() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.HasUpdates
}

// GetHasUpdatesOk returns a tuple with the HasUpdates field value
// and a boolean to check if the value has been set.
func (o *NegotiationsListItem) GetHasUpdatesOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.HasUpdates, true
}

// SetHasUpdates sets field value
func (o *NegotiationsListItem) SetHasUpdates(v bool) {
	o.HasUpdates = v
}

// GetId returns the Id field value
func (o *NegotiationsListItem) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *NegotiationsListItem) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *NegotiationsListItem) SetId(v string) {
	o.Id = v
}

// GetMessagesUrl returns the MessagesUrl field value if set, zero value otherwise.
// Deprecated
func (o *NegotiationsListItem) GetMessagesUrl() string {
	if o == nil || IsNil(o.MessagesUrl) {
		var ret string
		return ret
	}
	return *o.MessagesUrl
}

// GetMessagesUrlOk returns a tuple with the MessagesUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
// Deprecated
func (o *NegotiationsListItem) GetMessagesUrlOk() (*string, bool) {
	if o == nil || IsNil(o.MessagesUrl) {
		return nil, false
	}
	return o.MessagesUrl, true
}

// HasMessagesUrl returns a boolean if a field has been set.
func (o *NegotiationsListItem) HasMessagesUrl() bool {
	if o != nil && !IsNil(o.MessagesUrl) {
		return true
	}

	return false
}

// SetMessagesUrl gets a reference to the given string and assigns it to the MessagesUrl field.
// Deprecated
func (o *NegotiationsListItem) SetMessagesUrl(v string) {
	o.MessagesUrl = &v
}

// GetMessagingStatus returns the MessagingStatus field value
func (o *NegotiationsListItem) GetMessagingStatus() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.MessagingStatus
}

// GetMessagingStatusOk returns a tuple with the MessagingStatus field value
// and a boolean to check if the value has been set.
func (o *NegotiationsListItem) GetMessagingStatusOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.MessagingStatus, true
}

// SetMessagingStatus sets field value
func (o *NegotiationsListItem) SetMessagingStatus(v string) {
	o.MessagingStatus = v
}

// GetProfessionalRoles returns the ProfessionalRoles field value if set, zero value otherwise.
func (o *NegotiationsListItem) GetProfessionalRoles() []VacancyProfessionalRoleItem {
	if o == nil || IsNil(o.ProfessionalRoles) {
		var ret []VacancyProfessionalRoleItem
		return ret
	}
	return o.ProfessionalRoles
}

// GetProfessionalRolesOk returns a tuple with the ProfessionalRoles field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NegotiationsListItem) GetProfessionalRolesOk() ([]VacancyProfessionalRoleItem, bool) {
	if o == nil || IsNil(o.ProfessionalRoles) {
		return nil, false
	}
	return o.ProfessionalRoles, true
}

// HasProfessionalRoles returns a boolean if a field has been set.
func (o *NegotiationsListItem) HasProfessionalRoles() bool {
	if o != nil && !IsNil(o.ProfessionalRoles) {
		return true
	}

	return false
}

// SetProfessionalRoles gets a reference to the given []VacancyProfessionalRoleItem and assigns it to the ProfessionalRoles field.
func (o *NegotiationsListItem) SetProfessionalRoles(v []VacancyProfessionalRoleItem) {
	o.ProfessionalRoles = v
}

// GetSource returns the Source field value if set, zero value otherwise.
func (o *NegotiationsListItem) GetSource() string {
	if o == nil || IsNil(o.Source) {
		var ret string
		return ret
	}
	return *o.Source
}

// GetSourceOk returns a tuple with the Source field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NegotiationsListItem) GetSourceOk() (*string, bool) {
	if o == nil || IsNil(o.Source) {
		return nil, false
	}
	return o.Source, true
}

// HasSource returns a boolean if a field has been set.
func (o *NegotiationsListItem) HasSource() bool {
	if o != nil && !IsNil(o.Source) {
		return true
	}

	return false
}

// SetSource gets a reference to the given string and assigns it to the Source field.
func (o *NegotiationsListItem) SetSource(v string) {
	o.Source = &v
}

// GetState returns the State field value
func (o *NegotiationsListItem) GetState() IncludesIdName {
	if o == nil {
		var ret IncludesIdName
		return ret
	}

	return o.State
}

// GetStateOk returns a tuple with the State field value
// and a boolean to check if the value has been set.
func (o *NegotiationsListItem) GetStateOk() (*IncludesIdName, bool) {
	if o == nil {
		return nil, false
	}
	return &o.State, true
}

// SetState sets field value
func (o *NegotiationsListItem) SetState(v IncludesIdName) {
	o.State = v
}

// GetUpdatedAt returns the UpdatedAt field value
func (o *NegotiationsListItem) GetUpdatedAt() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.UpdatedAt
}

// GetUpdatedAtOk returns a tuple with the UpdatedAt field value
// and a boolean to check if the value has been set.
func (o *NegotiationsListItem) GetUpdatedAtOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.UpdatedAt, true
}

// SetUpdatedAt sets field value
func (o *NegotiationsListItem) SetUpdatedAt(v string) {
	o.UpdatedAt = v
}

// GetViewedByOpponent returns the ViewedByOpponent field value
func (o *NegotiationsListItem) GetViewedByOpponent() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.ViewedByOpponent
}

// GetViewedByOpponentOk returns a tuple with the ViewedByOpponent field value
// and a boolean to check if the value has been set.
func (o *NegotiationsListItem) GetViewedByOpponentOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ViewedByOpponent, true
}

// SetViewedByOpponent sets field value
func (o *NegotiationsListItem) SetViewedByOpponent(v bool) {
	o.ViewedByOpponent = v
}

// GetDeclineAllowed returns the DeclineAllowed field value
func (o *NegotiationsListItem) GetDeclineAllowed() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.DeclineAllowed
}

// GetDeclineAllowedOk returns a tuple with the DeclineAllowed field value
// and a boolean to check if the value has been set.
func (o *NegotiationsListItem) GetDeclineAllowedOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.DeclineAllowed, true
}

// SetDeclineAllowed sets field value
func (o *NegotiationsListItem) SetDeclineAllowed(v bool) {
	o.DeclineAllowed = v
}

// GetHidden returns the Hidden field value
func (o *NegotiationsListItem) GetHidden() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.Hidden
}

// GetHiddenOk returns a tuple with the Hidden field value
// and a boolean to check if the value has been set.
func (o *NegotiationsListItem) GetHiddenOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Hidden, true
}

// SetHidden sets field value
func (o *NegotiationsListItem) SetHidden(v bool) {
	o.Hidden = v
}

// GetJobSearchStatus returns the JobSearchStatus field value if set, zero value otherwise.
func (o *NegotiationsListItem) GetJobSearchStatus() IncludesIdName {
	if o == nil || IsNil(o.JobSearchStatus) {
		var ret IncludesIdName
		return ret
	}
	return *o.JobSearchStatus
}

// GetJobSearchStatusOk returns a tuple with the JobSearchStatus field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NegotiationsListItem) GetJobSearchStatusOk() (*IncludesIdName, bool) {
	if o == nil || IsNil(o.JobSearchStatus) {
		return nil, false
	}
	return o.JobSearchStatus, true
}

// HasJobSearchStatus returns a boolean if a field has been set.
func (o *NegotiationsListItem) HasJobSearchStatus() bool {
	if o != nil && !IsNil(o.JobSearchStatus) {
		return true
	}

	return false
}

// SetJobSearchStatus gets a reference to the given IncludesIdName and assigns it to the JobSearchStatus field.
func (o *NegotiationsListItem) SetJobSearchStatus(v IncludesIdName) {
	o.JobSearchStatus = &v
}

// GetPhoneCalls returns the PhoneCalls field value if set, zero value otherwise.
func (o *NegotiationsListItem) GetPhoneCalls() NegotiationsPhoneCalls {
	if o == nil || IsNil(o.PhoneCalls) {
		var ret NegotiationsPhoneCalls
		return ret
	}
	return *o.PhoneCalls
}

// GetPhoneCallsOk returns a tuple with the PhoneCalls field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NegotiationsListItem) GetPhoneCallsOk() (*NegotiationsPhoneCalls, bool) {
	if o == nil || IsNil(o.PhoneCalls) {
		return nil, false
	}
	return o.PhoneCalls, true
}

// HasPhoneCalls returns a boolean if a field has been set.
func (o *NegotiationsListItem) HasPhoneCalls() bool {
	if o != nil && !IsNil(o.PhoneCalls) {
		return true
	}

	return false
}

// SetPhoneCalls gets a reference to the given NegotiationsPhoneCalls and assigns it to the PhoneCalls field.
func (o *NegotiationsListItem) SetPhoneCalls(v NegotiationsPhoneCalls) {
	o.PhoneCalls = &v
}

// GetVacancy returns the Vacancy field value if set, zero value otherwise.
func (o *NegotiationsListItem) GetVacancy() VacanciesNegotiationsVacancyShort {
	if o == nil || IsNil(o.Vacancy) {
		var ret VacanciesNegotiationsVacancyShort
		return ret
	}
	return *o.Vacancy
}

// GetVacancyOk returns a tuple with the Vacancy field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NegotiationsListItem) GetVacancyOk() (*VacanciesNegotiationsVacancyShort, bool) {
	if o == nil || IsNil(o.Vacancy) {
		return nil, false
	}
	return o.Vacancy, true
}

// HasVacancy returns a boolean if a field has been set.
func (o *NegotiationsListItem) HasVacancy() bool {
	if o != nil && !IsNil(o.Vacancy) {
		return true
	}

	return false
}

// SetVacancy gets a reference to the given VacanciesNegotiationsVacancyShort and assigns it to the Vacancy field.
func (o *NegotiationsListItem) SetVacancy(v VacanciesNegotiationsVacancyShort) {
	o.Vacancy = &v
}

// GetResume returns the Resume field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *NegotiationsListItem) GetResume() ResumeResumeNanoWithUrl {
	if o == nil || IsNil(o.Resume.Get()) {
		var ret ResumeResumeNanoWithUrl
		return ret
	}
	return *o.Resume.Get()
}

// GetResumeOk returns a tuple with the Resume field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *NegotiationsListItem) GetResumeOk() (*ResumeResumeNanoWithUrl, bool) {
	if o == nil {
		return nil, false
	}
	return o.Resume.Get(), o.Resume.IsSet()
}

// HasResume returns a boolean if a field has been set.
func (o *NegotiationsListItem) HasResume() bool {
	if o != nil && o.Resume.IsSet() {
		return true
	}

	return false
}

// SetResume gets a reference to the given NullableResumeResumeNanoWithUrl and assigns it to the Resume field.
func (o *NegotiationsListItem) SetResume(v ResumeResumeNanoWithUrl) {
	o.Resume.Set(&v)
}
// SetResumeNil sets the value for Resume to be an explicit nil
func (o *NegotiationsListItem) SetResumeNil() {
	o.Resume.Set(nil)
}

// UnsetResume ensures that no value is present for Resume, not even an explicit nil
func (o *NegotiationsListItem) UnsetResume() {
	o.Resume.Unset()
}

// GetUrl returns the Url field value
func (o *NegotiationsListItem) GetUrl() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Url
}

// GetUrlOk returns a tuple with the Url field value
// and a boolean to check if the value has been set.
func (o *NegotiationsListItem) GetUrlOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Url, true
}

// SetUrl sets field value
func (o *NegotiationsListItem) SetUrl(v string) {
	o.Url = v
}

func (o NegotiationsListItem) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o NegotiationsListItem) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Counters) {
		toSerialize["counters"] = o.Counters
	}
	toSerialize["created_at"] = o.CreatedAt
	toSerialize["has_updates"] = o.HasUpdates
	toSerialize["id"] = o.Id
	if !IsNil(o.MessagesUrl) {
		toSerialize["messages_url"] = o.MessagesUrl
	}
	toSerialize["messaging_status"] = o.MessagingStatus
	if !IsNil(o.ProfessionalRoles) {
		toSerialize["professional_roles"] = o.ProfessionalRoles
	}
	if !IsNil(o.Source) {
		toSerialize["source"] = o.Source
	}
	toSerialize["state"] = o.State
	toSerialize["updated_at"] = o.UpdatedAt
	toSerialize["viewed_by_opponent"] = o.ViewedByOpponent
	toSerialize["decline_allowed"] = o.DeclineAllowed
	toSerialize["hidden"] = o.Hidden
	if !IsNil(o.JobSearchStatus) {
		toSerialize["job_search_status"] = o.JobSearchStatus
	}
	if !IsNil(o.PhoneCalls) {
		toSerialize["phone_calls"] = o.PhoneCalls
	}
	if !IsNil(o.Vacancy) {
		toSerialize["vacancy"] = o.Vacancy
	}
	if o.Resume.IsSet() {
		toSerialize["resume"] = o.Resume.Get()
	}
	toSerialize["url"] = o.Url
	return toSerialize, nil
}

func (o *NegotiationsListItem) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"created_at",
		"has_updates",
		"id",
		"messaging_status",
		"state",
		"updated_at",
		"viewed_by_opponent",
		"decline_allowed",
		"hidden",
		"url",
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

	varNegotiationsListItem := _NegotiationsListItem{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varNegotiationsListItem)

	if err != nil {
		return err
	}

	*o = NegotiationsListItem(varNegotiationsListItem)

	return err
}

type NullableNegotiationsListItem struct {
	value *NegotiationsListItem
	isSet bool
}

func (v NullableNegotiationsListItem) Get() *NegotiationsListItem {
	return v.value
}

func (v *NullableNegotiationsListItem) Set(val *NegotiationsListItem) {
	v.value = val
	v.isSet = true
}

func (v NullableNegotiationsListItem) IsSet() bool {
	return v.isSet
}

func (v *NullableNegotiationsListItem) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableNegotiationsListItem(val *NegotiationsListItem) *NullableNegotiationsListItem {
	return &NullableNegotiationsListItem{value: val, isSet: true}
}

func (v NullableNegotiationsListItem) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableNegotiationsListItem) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


