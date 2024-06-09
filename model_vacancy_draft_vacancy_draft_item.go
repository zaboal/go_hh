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

// checks if the VacancyDraftVacancyDraftItem type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &VacancyDraftVacancyDraftItem{}

// VacancyDraftVacancyDraftItem struct for VacancyDraftVacancyDraftItem
type VacancyDraftVacancyDraftItem struct {
	AutoPublication *VacancyDraftAutoPublicationState `json:"auto_publication,omitempty"`
	// Процент заполнения черновика
	CompletedFieldsPercentage float32 `json:"completed_fields_percentage"`
	// Идентификатор черновика
	DraftId string `json:"draft_id"`
	// Массив объектов с информацией о том, каких публикаций не хватает на счету для публикации вакансии из данного черновика
	InsufficientPublications []VacancyDraftPublications `json:"insufficient_publications,omitempty"`
	// Массив объектов с информацией о том, какие квоты превышены
	InsufficientQuotas []VacancyDraftPublications `json:"insufficient_quotas,omitempty"`
	// Время изменения черновика (в формате [ISO 8601](https://ru.wikipedia.org/wiki/ISO_8601) с точностью до секунды `YYYY-MM-DDThh:mm:ss±hhmm`)
	LastChangeTime NullableString `json:"last_change_time,omitempty"`
	// Готовность черновика к публикации
	PublicationReady bool `json:"publication_ready"`
	// Массив объектов с информацией о необходимых публикациях на счету
	RequiredPublications []VacancyDraftPublications `json:"required_publications,omitempty"`
	// Время запланированной публикации вакансии (в формате [ISO 8601](http://en.wikipedia.org/wiki/ISO_8601) с точностью до секунды: `YYYY-MM-DDThh:mm:ss±hhmm`
	ScheduledAt string `json:"scheduled_at"`
	// Коды и названия регионов (фед. округа, субъекты федерации, города)
	Areas []VacancyAreaOutput `json:"areas"`
	AssignedManager NullableVacancyDraftAssignedManager `json:"assigned_manager,omitempty"`
	BillingType NullableVacancyBillingTypeOutput `json:"billing_type"`
	// Название вакансии
	Name *string `json:"name,omitempty"`
	// Тип публикации (справочник [vacancy_billing_type](#tag/Obshie-spravochniki/operation/get-dictionaries))
	// Deprecated
	PublicationType string `json:"publication_type"`
	// Url для запроса полной информации черновика
	Url string `json:"url"`
	// Тип вакансии (справочник [vacancy_type](#tag/Obshie-spravochniki/operation/get-dictionaries))
	VacancyType NullableString `json:"vacancy_type"`
}

type _VacancyDraftVacancyDraftItem VacancyDraftVacancyDraftItem

// NewVacancyDraftVacancyDraftItem instantiates a new VacancyDraftVacancyDraftItem object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewVacancyDraftVacancyDraftItem(completedFieldsPercentage float32, draftId string, publicationReady bool, scheduledAt string, areas []VacancyAreaOutput, billingType NullableVacancyBillingTypeOutput, publicationType string, url string, vacancyType NullableString) *VacancyDraftVacancyDraftItem {
	this := VacancyDraftVacancyDraftItem{}
	this.CompletedFieldsPercentage = completedFieldsPercentage
	this.DraftId = draftId
	this.PublicationReady = publicationReady
	this.ScheduledAt = scheduledAt
	this.Areas = areas
	this.BillingType = billingType
	this.PublicationType = publicationType
	this.Url = url
	this.VacancyType = vacancyType
	return &this
}

// NewVacancyDraftVacancyDraftItemWithDefaults instantiates a new VacancyDraftVacancyDraftItem object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewVacancyDraftVacancyDraftItemWithDefaults() *VacancyDraftVacancyDraftItem {
	this := VacancyDraftVacancyDraftItem{}
	return &this
}

// GetAutoPublication returns the AutoPublication field value if set, zero value otherwise.
func (o *VacancyDraftVacancyDraftItem) GetAutoPublication() VacancyDraftAutoPublicationState {
	if o == nil || IsNil(o.AutoPublication) {
		var ret VacancyDraftAutoPublicationState
		return ret
	}
	return *o.AutoPublication
}

// GetAutoPublicationOk returns a tuple with the AutoPublication field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VacancyDraftVacancyDraftItem) GetAutoPublicationOk() (*VacancyDraftAutoPublicationState, bool) {
	if o == nil || IsNil(o.AutoPublication) {
		return nil, false
	}
	return o.AutoPublication, true
}

// HasAutoPublication returns a boolean if a field has been set.
func (o *VacancyDraftVacancyDraftItem) HasAutoPublication() bool {
	if o != nil && !IsNil(o.AutoPublication) {
		return true
	}

	return false
}

// SetAutoPublication gets a reference to the given VacancyDraftAutoPublicationState and assigns it to the AutoPublication field.
func (o *VacancyDraftVacancyDraftItem) SetAutoPublication(v VacancyDraftAutoPublicationState) {
	o.AutoPublication = &v
}

// GetCompletedFieldsPercentage returns the CompletedFieldsPercentage field value
func (o *VacancyDraftVacancyDraftItem) GetCompletedFieldsPercentage() float32 {
	if o == nil {
		var ret float32
		return ret
	}

	return o.CompletedFieldsPercentage
}

// GetCompletedFieldsPercentageOk returns a tuple with the CompletedFieldsPercentage field value
// and a boolean to check if the value has been set.
func (o *VacancyDraftVacancyDraftItem) GetCompletedFieldsPercentageOk() (*float32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CompletedFieldsPercentage, true
}

// SetCompletedFieldsPercentage sets field value
func (o *VacancyDraftVacancyDraftItem) SetCompletedFieldsPercentage(v float32) {
	o.CompletedFieldsPercentage = v
}

// GetDraftId returns the DraftId field value
func (o *VacancyDraftVacancyDraftItem) GetDraftId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.DraftId
}

// GetDraftIdOk returns a tuple with the DraftId field value
// and a boolean to check if the value has been set.
func (o *VacancyDraftVacancyDraftItem) GetDraftIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.DraftId, true
}

// SetDraftId sets field value
func (o *VacancyDraftVacancyDraftItem) SetDraftId(v string) {
	o.DraftId = v
}

// GetInsufficientPublications returns the InsufficientPublications field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *VacancyDraftVacancyDraftItem) GetInsufficientPublications() []VacancyDraftPublications {
	if o == nil {
		var ret []VacancyDraftPublications
		return ret
	}
	return o.InsufficientPublications
}

// GetInsufficientPublicationsOk returns a tuple with the InsufficientPublications field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *VacancyDraftVacancyDraftItem) GetInsufficientPublicationsOk() ([]VacancyDraftPublications, bool) {
	if o == nil || IsNil(o.InsufficientPublications) {
		return nil, false
	}
	return o.InsufficientPublications, true
}

// HasInsufficientPublications returns a boolean if a field has been set.
func (o *VacancyDraftVacancyDraftItem) HasInsufficientPublications() bool {
	if o != nil && !IsNil(o.InsufficientPublications) {
		return true
	}

	return false
}

// SetInsufficientPublications gets a reference to the given []VacancyDraftPublications and assigns it to the InsufficientPublications field.
func (o *VacancyDraftVacancyDraftItem) SetInsufficientPublications(v []VacancyDraftPublications) {
	o.InsufficientPublications = v
}

// GetInsufficientQuotas returns the InsufficientQuotas field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *VacancyDraftVacancyDraftItem) GetInsufficientQuotas() []VacancyDraftPublications {
	if o == nil {
		var ret []VacancyDraftPublications
		return ret
	}
	return o.InsufficientQuotas
}

// GetInsufficientQuotasOk returns a tuple with the InsufficientQuotas field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *VacancyDraftVacancyDraftItem) GetInsufficientQuotasOk() ([]VacancyDraftPublications, bool) {
	if o == nil || IsNil(o.InsufficientQuotas) {
		return nil, false
	}
	return o.InsufficientQuotas, true
}

// HasInsufficientQuotas returns a boolean if a field has been set.
func (o *VacancyDraftVacancyDraftItem) HasInsufficientQuotas() bool {
	if o != nil && !IsNil(o.InsufficientQuotas) {
		return true
	}

	return false
}

// SetInsufficientQuotas gets a reference to the given []VacancyDraftPublications and assigns it to the InsufficientQuotas field.
func (o *VacancyDraftVacancyDraftItem) SetInsufficientQuotas(v []VacancyDraftPublications) {
	o.InsufficientQuotas = v
}

// GetLastChangeTime returns the LastChangeTime field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *VacancyDraftVacancyDraftItem) GetLastChangeTime() string {
	if o == nil || IsNil(o.LastChangeTime.Get()) {
		var ret string
		return ret
	}
	return *o.LastChangeTime.Get()
}

// GetLastChangeTimeOk returns a tuple with the LastChangeTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *VacancyDraftVacancyDraftItem) GetLastChangeTimeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.LastChangeTime.Get(), o.LastChangeTime.IsSet()
}

// HasLastChangeTime returns a boolean if a field has been set.
func (o *VacancyDraftVacancyDraftItem) HasLastChangeTime() bool {
	if o != nil && o.LastChangeTime.IsSet() {
		return true
	}

	return false
}

// SetLastChangeTime gets a reference to the given NullableString and assigns it to the LastChangeTime field.
func (o *VacancyDraftVacancyDraftItem) SetLastChangeTime(v string) {
	o.LastChangeTime.Set(&v)
}
// SetLastChangeTimeNil sets the value for LastChangeTime to be an explicit nil
func (o *VacancyDraftVacancyDraftItem) SetLastChangeTimeNil() {
	o.LastChangeTime.Set(nil)
}

// UnsetLastChangeTime ensures that no value is present for LastChangeTime, not even an explicit nil
func (o *VacancyDraftVacancyDraftItem) UnsetLastChangeTime() {
	o.LastChangeTime.Unset()
}

// GetPublicationReady returns the PublicationReady field value
func (o *VacancyDraftVacancyDraftItem) GetPublicationReady() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.PublicationReady
}

// GetPublicationReadyOk returns a tuple with the PublicationReady field value
// and a boolean to check if the value has been set.
func (o *VacancyDraftVacancyDraftItem) GetPublicationReadyOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.PublicationReady, true
}

// SetPublicationReady sets field value
func (o *VacancyDraftVacancyDraftItem) SetPublicationReady(v bool) {
	o.PublicationReady = v
}

// GetRequiredPublications returns the RequiredPublications field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *VacancyDraftVacancyDraftItem) GetRequiredPublications() []VacancyDraftPublications {
	if o == nil {
		var ret []VacancyDraftPublications
		return ret
	}
	return o.RequiredPublications
}

// GetRequiredPublicationsOk returns a tuple with the RequiredPublications field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *VacancyDraftVacancyDraftItem) GetRequiredPublicationsOk() ([]VacancyDraftPublications, bool) {
	if o == nil || IsNil(o.RequiredPublications) {
		return nil, false
	}
	return o.RequiredPublications, true
}

// HasRequiredPublications returns a boolean if a field has been set.
func (o *VacancyDraftVacancyDraftItem) HasRequiredPublications() bool {
	if o != nil && !IsNil(o.RequiredPublications) {
		return true
	}

	return false
}

// SetRequiredPublications gets a reference to the given []VacancyDraftPublications and assigns it to the RequiredPublications field.
func (o *VacancyDraftVacancyDraftItem) SetRequiredPublications(v []VacancyDraftPublications) {
	o.RequiredPublications = v
}

// GetScheduledAt returns the ScheduledAt field value
func (o *VacancyDraftVacancyDraftItem) GetScheduledAt() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ScheduledAt
}

// GetScheduledAtOk returns a tuple with the ScheduledAt field value
// and a boolean to check if the value has been set.
func (o *VacancyDraftVacancyDraftItem) GetScheduledAtOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ScheduledAt, true
}

// SetScheduledAt sets field value
func (o *VacancyDraftVacancyDraftItem) SetScheduledAt(v string) {
	o.ScheduledAt = v
}

// GetAreas returns the Areas field value
func (o *VacancyDraftVacancyDraftItem) GetAreas() []VacancyAreaOutput {
	if o == nil {
		var ret []VacancyAreaOutput
		return ret
	}

	return o.Areas
}

// GetAreasOk returns a tuple with the Areas field value
// and a boolean to check if the value has been set.
func (o *VacancyDraftVacancyDraftItem) GetAreasOk() ([]VacancyAreaOutput, bool) {
	if o == nil {
		return nil, false
	}
	return o.Areas, true
}

// SetAreas sets field value
func (o *VacancyDraftVacancyDraftItem) SetAreas(v []VacancyAreaOutput) {
	o.Areas = v
}

// GetAssignedManager returns the AssignedManager field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *VacancyDraftVacancyDraftItem) GetAssignedManager() VacancyDraftAssignedManager {
	if o == nil || IsNil(o.AssignedManager.Get()) {
		var ret VacancyDraftAssignedManager
		return ret
	}
	return *o.AssignedManager.Get()
}

// GetAssignedManagerOk returns a tuple with the AssignedManager field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *VacancyDraftVacancyDraftItem) GetAssignedManagerOk() (*VacancyDraftAssignedManager, bool) {
	if o == nil {
		return nil, false
	}
	return o.AssignedManager.Get(), o.AssignedManager.IsSet()
}

// HasAssignedManager returns a boolean if a field has been set.
func (o *VacancyDraftVacancyDraftItem) HasAssignedManager() bool {
	if o != nil && o.AssignedManager.IsSet() {
		return true
	}

	return false
}

// SetAssignedManager gets a reference to the given NullableVacancyDraftAssignedManager and assigns it to the AssignedManager field.
func (o *VacancyDraftVacancyDraftItem) SetAssignedManager(v VacancyDraftAssignedManager) {
	o.AssignedManager.Set(&v)
}
// SetAssignedManagerNil sets the value for AssignedManager to be an explicit nil
func (o *VacancyDraftVacancyDraftItem) SetAssignedManagerNil() {
	o.AssignedManager.Set(nil)
}

// UnsetAssignedManager ensures that no value is present for AssignedManager, not even an explicit nil
func (o *VacancyDraftVacancyDraftItem) UnsetAssignedManager() {
	o.AssignedManager.Unset()
}

// GetBillingType returns the BillingType field value
// If the value is explicit nil, the zero value for VacancyBillingTypeOutput will be returned
func (o *VacancyDraftVacancyDraftItem) GetBillingType() VacancyBillingTypeOutput {
	if o == nil || o.BillingType.Get() == nil {
		var ret VacancyBillingTypeOutput
		return ret
	}

	return *o.BillingType.Get()
}

// GetBillingTypeOk returns a tuple with the BillingType field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *VacancyDraftVacancyDraftItem) GetBillingTypeOk() (*VacancyBillingTypeOutput, bool) {
	if o == nil {
		return nil, false
	}
	return o.BillingType.Get(), o.BillingType.IsSet()
}

// SetBillingType sets field value
func (o *VacancyDraftVacancyDraftItem) SetBillingType(v VacancyBillingTypeOutput) {
	o.BillingType.Set(&v)
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *VacancyDraftVacancyDraftItem) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VacancyDraftVacancyDraftItem) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *VacancyDraftVacancyDraftItem) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *VacancyDraftVacancyDraftItem) SetName(v string) {
	o.Name = &v
}

// GetPublicationType returns the PublicationType field value
// Deprecated
func (o *VacancyDraftVacancyDraftItem) GetPublicationType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.PublicationType
}

// GetPublicationTypeOk returns a tuple with the PublicationType field value
// and a boolean to check if the value has been set.
// Deprecated
func (o *VacancyDraftVacancyDraftItem) GetPublicationTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.PublicationType, true
}

// SetPublicationType sets field value
// Deprecated
func (o *VacancyDraftVacancyDraftItem) SetPublicationType(v string) {
	o.PublicationType = v
}

// GetUrl returns the Url field value
func (o *VacancyDraftVacancyDraftItem) GetUrl() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Url
}

// GetUrlOk returns a tuple with the Url field value
// and a boolean to check if the value has been set.
func (o *VacancyDraftVacancyDraftItem) GetUrlOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Url, true
}

// SetUrl sets field value
func (o *VacancyDraftVacancyDraftItem) SetUrl(v string) {
	o.Url = v
}

// GetVacancyType returns the VacancyType field value
// If the value is explicit nil, the zero value for string will be returned
func (o *VacancyDraftVacancyDraftItem) GetVacancyType() string {
	if o == nil || o.VacancyType.Get() == nil {
		var ret string
		return ret
	}

	return *o.VacancyType.Get()
}

// GetVacancyTypeOk returns a tuple with the VacancyType field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *VacancyDraftVacancyDraftItem) GetVacancyTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.VacancyType.Get(), o.VacancyType.IsSet()
}

// SetVacancyType sets field value
func (o *VacancyDraftVacancyDraftItem) SetVacancyType(v string) {
	o.VacancyType.Set(&v)
}

func (o VacancyDraftVacancyDraftItem) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o VacancyDraftVacancyDraftItem) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.AutoPublication) {
		toSerialize["auto_publication"] = o.AutoPublication
	}
	toSerialize["completed_fields_percentage"] = o.CompletedFieldsPercentage
	toSerialize["draft_id"] = o.DraftId
	if o.InsufficientPublications != nil {
		toSerialize["insufficient_publications"] = o.InsufficientPublications
	}
	if o.InsufficientQuotas != nil {
		toSerialize["insufficient_quotas"] = o.InsufficientQuotas
	}
	if o.LastChangeTime.IsSet() {
		toSerialize["last_change_time"] = o.LastChangeTime.Get()
	}
	toSerialize["publication_ready"] = o.PublicationReady
	if o.RequiredPublications != nil {
		toSerialize["required_publications"] = o.RequiredPublications
	}
	toSerialize["scheduled_at"] = o.ScheduledAt
	toSerialize["areas"] = o.Areas
	if o.AssignedManager.IsSet() {
		toSerialize["assigned_manager"] = o.AssignedManager.Get()
	}
	toSerialize["billing_type"] = o.BillingType.Get()
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	toSerialize["publication_type"] = o.PublicationType
	toSerialize["url"] = o.Url
	toSerialize["vacancy_type"] = o.VacancyType.Get()
	return toSerialize, nil
}

func (o *VacancyDraftVacancyDraftItem) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"completed_fields_percentage",
		"draft_id",
		"publication_ready",
		"scheduled_at",
		"areas",
		"billing_type",
		"publication_type",
		"url",
		"vacancy_type",
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

	varVacancyDraftVacancyDraftItem := _VacancyDraftVacancyDraftItem{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varVacancyDraftVacancyDraftItem)

	if err != nil {
		return err
	}

	*o = VacancyDraftVacancyDraftItem(varVacancyDraftVacancyDraftItem)

	return err
}

type NullableVacancyDraftVacancyDraftItem struct {
	value *VacancyDraftVacancyDraftItem
	isSet bool
}

func (v NullableVacancyDraftVacancyDraftItem) Get() *VacancyDraftVacancyDraftItem {
	return v.value
}

func (v *NullableVacancyDraftVacancyDraftItem) Set(val *VacancyDraftVacancyDraftItem) {
	v.value = val
	v.isSet = true
}

func (v NullableVacancyDraftVacancyDraftItem) IsSet() bool {
	return v.isSet
}

func (v *NullableVacancyDraftVacancyDraftItem) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableVacancyDraftVacancyDraftItem(val *VacancyDraftVacancyDraftItem) *NullableVacancyDraftVacancyDraftItem {
	return &NullableVacancyDraftVacancyDraftItem{value: val, isSet: true}
}

func (v NullableVacancyDraftVacancyDraftItem) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableVacancyDraftVacancyDraftItem) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


