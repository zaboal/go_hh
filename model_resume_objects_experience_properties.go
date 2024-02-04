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
)

// checks if the ResumeObjectsExperienceProperties type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ResumeObjectsExperienceProperties{}

// ResumeObjectsExperienceProperties struct for ResumeObjectsExperienceProperties
type ResumeObjectsExperienceProperties struct {
	Area NullableIncludesIdNameUrl `json:"area,omitempty"`
	// Название организации
	Company NullableString `json:"company,omitempty"`
	// Уникальный идентификатор организации
	CompanyId NullableString `json:"company_id,omitempty"`
	// Сайт компании
	CompanyUrl NullableString `json:"company_url,omitempty"`
	// Обязанности, функции, достижения
	Description NullableString `json:"description,omitempty"`
	Employer NullableEmployersEmployerInfoShort `json:"employer,omitempty"`
	// Окончание работы (дата в формате `ГГГГ-ММ-ДД`)
	End NullableString `json:"end,omitempty"`
	// Список отраслей компании. Возможные значения приведены в [справочнике индустрий](#tag/Obshie-spravochniki/operation/get-industries)
	Industries []IncludesIdName `json:"industries,omitempty"`
	// Deprecated
	Industry NullableResumeObjectsIndustry `json:"industry,omitempty"`
	// Должность
	Position *string `json:"position,omitempty"`
	// Начало работы (дата в формате `ГГГГ-ММ-ДД`)
	Start *string `json:"start,omitempty"`
}

// NewResumeObjectsExperienceProperties instantiates a new ResumeObjectsExperienceProperties object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewResumeObjectsExperienceProperties() *ResumeObjectsExperienceProperties {
	this := ResumeObjectsExperienceProperties{}
	return &this
}

// NewResumeObjectsExperiencePropertiesWithDefaults instantiates a new ResumeObjectsExperienceProperties object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewResumeObjectsExperiencePropertiesWithDefaults() *ResumeObjectsExperienceProperties {
	this := ResumeObjectsExperienceProperties{}
	return &this
}

// GetArea returns the Area field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ResumeObjectsExperienceProperties) GetArea() IncludesIdNameUrl {
	if o == nil || IsNil(o.Area.Get()) {
		var ret IncludesIdNameUrl
		return ret
	}
	return *o.Area.Get()
}

// GetAreaOk returns a tuple with the Area field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ResumeObjectsExperienceProperties) GetAreaOk() (*IncludesIdNameUrl, bool) {
	if o == nil {
		return nil, false
	}
	return o.Area.Get(), o.Area.IsSet()
}

// HasArea returns a boolean if a field has been set.
func (o *ResumeObjectsExperienceProperties) HasArea() bool {
	if o != nil && o.Area.IsSet() {
		return true
	}

	return false
}

// SetArea gets a reference to the given NullableIncludesIdNameUrl and assigns it to the Area field.
func (o *ResumeObjectsExperienceProperties) SetArea(v IncludesIdNameUrl) {
	o.Area.Set(&v)
}
// SetAreaNil sets the value for Area to be an explicit nil
func (o *ResumeObjectsExperienceProperties) SetAreaNil() {
	o.Area.Set(nil)
}

// UnsetArea ensures that no value is present for Area, not even an explicit nil
func (o *ResumeObjectsExperienceProperties) UnsetArea() {
	o.Area.Unset()
}

// GetCompany returns the Company field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ResumeObjectsExperienceProperties) GetCompany() string {
	if o == nil || IsNil(o.Company.Get()) {
		var ret string
		return ret
	}
	return *o.Company.Get()
}

// GetCompanyOk returns a tuple with the Company field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ResumeObjectsExperienceProperties) GetCompanyOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Company.Get(), o.Company.IsSet()
}

// HasCompany returns a boolean if a field has been set.
func (o *ResumeObjectsExperienceProperties) HasCompany() bool {
	if o != nil && o.Company.IsSet() {
		return true
	}

	return false
}

// SetCompany gets a reference to the given NullableString and assigns it to the Company field.
func (o *ResumeObjectsExperienceProperties) SetCompany(v string) {
	o.Company.Set(&v)
}
// SetCompanyNil sets the value for Company to be an explicit nil
func (o *ResumeObjectsExperienceProperties) SetCompanyNil() {
	o.Company.Set(nil)
}

// UnsetCompany ensures that no value is present for Company, not even an explicit nil
func (o *ResumeObjectsExperienceProperties) UnsetCompany() {
	o.Company.Unset()
}

// GetCompanyId returns the CompanyId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ResumeObjectsExperienceProperties) GetCompanyId() string {
	if o == nil || IsNil(o.CompanyId.Get()) {
		var ret string
		return ret
	}
	return *o.CompanyId.Get()
}

// GetCompanyIdOk returns a tuple with the CompanyId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ResumeObjectsExperienceProperties) GetCompanyIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.CompanyId.Get(), o.CompanyId.IsSet()
}

// HasCompanyId returns a boolean if a field has been set.
func (o *ResumeObjectsExperienceProperties) HasCompanyId() bool {
	if o != nil && o.CompanyId.IsSet() {
		return true
	}

	return false
}

// SetCompanyId gets a reference to the given NullableString and assigns it to the CompanyId field.
func (o *ResumeObjectsExperienceProperties) SetCompanyId(v string) {
	o.CompanyId.Set(&v)
}
// SetCompanyIdNil sets the value for CompanyId to be an explicit nil
func (o *ResumeObjectsExperienceProperties) SetCompanyIdNil() {
	o.CompanyId.Set(nil)
}

// UnsetCompanyId ensures that no value is present for CompanyId, not even an explicit nil
func (o *ResumeObjectsExperienceProperties) UnsetCompanyId() {
	o.CompanyId.Unset()
}

// GetCompanyUrl returns the CompanyUrl field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ResumeObjectsExperienceProperties) GetCompanyUrl() string {
	if o == nil || IsNil(o.CompanyUrl.Get()) {
		var ret string
		return ret
	}
	return *o.CompanyUrl.Get()
}

// GetCompanyUrlOk returns a tuple with the CompanyUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ResumeObjectsExperienceProperties) GetCompanyUrlOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.CompanyUrl.Get(), o.CompanyUrl.IsSet()
}

// HasCompanyUrl returns a boolean if a field has been set.
func (o *ResumeObjectsExperienceProperties) HasCompanyUrl() bool {
	if o != nil && o.CompanyUrl.IsSet() {
		return true
	}

	return false
}

// SetCompanyUrl gets a reference to the given NullableString and assigns it to the CompanyUrl field.
func (o *ResumeObjectsExperienceProperties) SetCompanyUrl(v string) {
	o.CompanyUrl.Set(&v)
}
// SetCompanyUrlNil sets the value for CompanyUrl to be an explicit nil
func (o *ResumeObjectsExperienceProperties) SetCompanyUrlNil() {
	o.CompanyUrl.Set(nil)
}

// UnsetCompanyUrl ensures that no value is present for CompanyUrl, not even an explicit nil
func (o *ResumeObjectsExperienceProperties) UnsetCompanyUrl() {
	o.CompanyUrl.Unset()
}

// GetDescription returns the Description field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ResumeObjectsExperienceProperties) GetDescription() string {
	if o == nil || IsNil(o.Description.Get()) {
		var ret string
		return ret
	}
	return *o.Description.Get()
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ResumeObjectsExperienceProperties) GetDescriptionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Description.Get(), o.Description.IsSet()
}

// HasDescription returns a boolean if a field has been set.
func (o *ResumeObjectsExperienceProperties) HasDescription() bool {
	if o != nil && o.Description.IsSet() {
		return true
	}

	return false
}

// SetDescription gets a reference to the given NullableString and assigns it to the Description field.
func (o *ResumeObjectsExperienceProperties) SetDescription(v string) {
	o.Description.Set(&v)
}
// SetDescriptionNil sets the value for Description to be an explicit nil
func (o *ResumeObjectsExperienceProperties) SetDescriptionNil() {
	o.Description.Set(nil)
}

// UnsetDescription ensures that no value is present for Description, not even an explicit nil
func (o *ResumeObjectsExperienceProperties) UnsetDescription() {
	o.Description.Unset()
}

// GetEmployer returns the Employer field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ResumeObjectsExperienceProperties) GetEmployer() EmployersEmployerInfoShort {
	if o == nil || IsNil(o.Employer.Get()) {
		var ret EmployersEmployerInfoShort
		return ret
	}
	return *o.Employer.Get()
}

// GetEmployerOk returns a tuple with the Employer field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ResumeObjectsExperienceProperties) GetEmployerOk() (*EmployersEmployerInfoShort, bool) {
	if o == nil {
		return nil, false
	}
	return o.Employer.Get(), o.Employer.IsSet()
}

// HasEmployer returns a boolean if a field has been set.
func (o *ResumeObjectsExperienceProperties) HasEmployer() bool {
	if o != nil && o.Employer.IsSet() {
		return true
	}

	return false
}

// SetEmployer gets a reference to the given NullableEmployersEmployerInfoShort and assigns it to the Employer field.
func (o *ResumeObjectsExperienceProperties) SetEmployer(v EmployersEmployerInfoShort) {
	o.Employer.Set(&v)
}
// SetEmployerNil sets the value for Employer to be an explicit nil
func (o *ResumeObjectsExperienceProperties) SetEmployerNil() {
	o.Employer.Set(nil)
}

// UnsetEmployer ensures that no value is present for Employer, not even an explicit nil
func (o *ResumeObjectsExperienceProperties) UnsetEmployer() {
	o.Employer.Unset()
}

// GetEnd returns the End field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ResumeObjectsExperienceProperties) GetEnd() string {
	if o == nil || IsNil(o.End.Get()) {
		var ret string
		return ret
	}
	return *o.End.Get()
}

// GetEndOk returns a tuple with the End field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ResumeObjectsExperienceProperties) GetEndOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.End.Get(), o.End.IsSet()
}

// HasEnd returns a boolean if a field has been set.
func (o *ResumeObjectsExperienceProperties) HasEnd() bool {
	if o != nil && o.End.IsSet() {
		return true
	}

	return false
}

// SetEnd gets a reference to the given NullableString and assigns it to the End field.
func (o *ResumeObjectsExperienceProperties) SetEnd(v string) {
	o.End.Set(&v)
}
// SetEndNil sets the value for End to be an explicit nil
func (o *ResumeObjectsExperienceProperties) SetEndNil() {
	o.End.Set(nil)
}

// UnsetEnd ensures that no value is present for End, not even an explicit nil
func (o *ResumeObjectsExperienceProperties) UnsetEnd() {
	o.End.Unset()
}

// GetIndustries returns the Industries field value if set, zero value otherwise.
func (o *ResumeObjectsExperienceProperties) GetIndustries() []IncludesIdName {
	if o == nil || IsNil(o.Industries) {
		var ret []IncludesIdName
		return ret
	}
	return o.Industries
}

// GetIndustriesOk returns a tuple with the Industries field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ResumeObjectsExperienceProperties) GetIndustriesOk() ([]IncludesIdName, bool) {
	if o == nil || IsNil(o.Industries) {
		return nil, false
	}
	return o.Industries, true
}

// HasIndustries returns a boolean if a field has been set.
func (o *ResumeObjectsExperienceProperties) HasIndustries() bool {
	if o != nil && !IsNil(o.Industries) {
		return true
	}

	return false
}

// SetIndustries gets a reference to the given []IncludesIdName and assigns it to the Industries field.
func (o *ResumeObjectsExperienceProperties) SetIndustries(v []IncludesIdName) {
	o.Industries = v
}

// GetIndustry returns the Industry field value if set, zero value otherwise (both if not set or set to explicit null).
// Deprecated
func (o *ResumeObjectsExperienceProperties) GetIndustry() ResumeObjectsIndustry {
	if o == nil || IsNil(o.Industry.Get()) {
		var ret ResumeObjectsIndustry
		return ret
	}
	return *o.Industry.Get()
}

// GetIndustryOk returns a tuple with the Industry field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
// Deprecated
func (o *ResumeObjectsExperienceProperties) GetIndustryOk() (*ResumeObjectsIndustry, bool) {
	if o == nil {
		return nil, false
	}
	return o.Industry.Get(), o.Industry.IsSet()
}

// HasIndustry returns a boolean if a field has been set.
func (o *ResumeObjectsExperienceProperties) HasIndustry() bool {
	if o != nil && o.Industry.IsSet() {
		return true
	}

	return false
}

// SetIndustry gets a reference to the given NullableResumeObjectsIndustry and assigns it to the Industry field.
// Deprecated
func (o *ResumeObjectsExperienceProperties) SetIndustry(v ResumeObjectsIndustry) {
	o.Industry.Set(&v)
}
// SetIndustryNil sets the value for Industry to be an explicit nil
func (o *ResumeObjectsExperienceProperties) SetIndustryNil() {
	o.Industry.Set(nil)
}

// UnsetIndustry ensures that no value is present for Industry, not even an explicit nil
func (o *ResumeObjectsExperienceProperties) UnsetIndustry() {
	o.Industry.Unset()
}

// GetPosition returns the Position field value if set, zero value otherwise.
func (o *ResumeObjectsExperienceProperties) GetPosition() string {
	if o == nil || IsNil(o.Position) {
		var ret string
		return ret
	}
	return *o.Position
}

// GetPositionOk returns a tuple with the Position field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ResumeObjectsExperienceProperties) GetPositionOk() (*string, bool) {
	if o == nil || IsNil(o.Position) {
		return nil, false
	}
	return o.Position, true
}

// HasPosition returns a boolean if a field has been set.
func (o *ResumeObjectsExperienceProperties) HasPosition() bool {
	if o != nil && !IsNil(o.Position) {
		return true
	}

	return false
}

// SetPosition gets a reference to the given string and assigns it to the Position field.
func (o *ResumeObjectsExperienceProperties) SetPosition(v string) {
	o.Position = &v
}

// GetStart returns the Start field value if set, zero value otherwise.
func (o *ResumeObjectsExperienceProperties) GetStart() string {
	if o == nil || IsNil(o.Start) {
		var ret string
		return ret
	}
	return *o.Start
}

// GetStartOk returns a tuple with the Start field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ResumeObjectsExperienceProperties) GetStartOk() (*string, bool) {
	if o == nil || IsNil(o.Start) {
		return nil, false
	}
	return o.Start, true
}

// HasStart returns a boolean if a field has been set.
func (o *ResumeObjectsExperienceProperties) HasStart() bool {
	if o != nil && !IsNil(o.Start) {
		return true
	}

	return false
}

// SetStart gets a reference to the given string and assigns it to the Start field.
func (o *ResumeObjectsExperienceProperties) SetStart(v string) {
	o.Start = &v
}

func (o ResumeObjectsExperienceProperties) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ResumeObjectsExperienceProperties) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.Area.IsSet() {
		toSerialize["area"] = o.Area.Get()
	}
	if o.Company.IsSet() {
		toSerialize["company"] = o.Company.Get()
	}
	if o.CompanyId.IsSet() {
		toSerialize["company_id"] = o.CompanyId.Get()
	}
	if o.CompanyUrl.IsSet() {
		toSerialize["company_url"] = o.CompanyUrl.Get()
	}
	if o.Description.IsSet() {
		toSerialize["description"] = o.Description.Get()
	}
	if o.Employer.IsSet() {
		toSerialize["employer"] = o.Employer.Get()
	}
	if o.End.IsSet() {
		toSerialize["end"] = o.End.Get()
	}
	if !IsNil(o.Industries) {
		toSerialize["industries"] = o.Industries
	}
	if o.Industry.IsSet() {
		toSerialize["industry"] = o.Industry.Get()
	}
	if !IsNil(o.Position) {
		toSerialize["position"] = o.Position
	}
	if !IsNil(o.Start) {
		toSerialize["start"] = o.Start
	}
	return toSerialize, nil
}

type NullableResumeObjectsExperienceProperties struct {
	value *ResumeObjectsExperienceProperties
	isSet bool
}

func (v NullableResumeObjectsExperienceProperties) Get() *ResumeObjectsExperienceProperties {
	return v.value
}

func (v *NullableResumeObjectsExperienceProperties) Set(val *ResumeObjectsExperienceProperties) {
	v.value = val
	v.isSet = true
}

func (v NullableResumeObjectsExperienceProperties) IsSet() bool {
	return v.isSet
}

func (v *NullableResumeObjectsExperienceProperties) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableResumeObjectsExperienceProperties(val *ResumeObjectsExperienceProperties) *NullableResumeObjectsExperienceProperties {
	return &NullableResumeObjectsExperienceProperties{value: val, isSet: true}
}

func (v NullableResumeObjectsExperienceProperties) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableResumeObjectsExperienceProperties) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


