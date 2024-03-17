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

// checks if the EmployerAddressesEmployerAddressItem type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &EmployerAddressesEmployerAddressItem{}

// EmployerAddressesEmployerAddressItem struct for EmployerAddressesEmployerAddressItem
type EmployerAddressesEmployerAddressItem struct {
	// Номер дома
	Building NullableString `json:"building,omitempty"`
	// Имеет ли текущий пользователь право редактировать этот адрес
	CanEdit *bool `json:"can_edit,omitempty"`
	// Город
	City NullableString `json:"city"`
	// Удалён ли адрес
	Deleted *bool `json:"deleted,omitempty"`
	// Дополнительная информация об адресе
	Description NullableString `json:"description,omitempty"`
	// Идентификатор адреса
	Id *string `json:"id,omitempty"`
	// Географическая широта
	Lat NullableFloat32 `json:"lat"`
	// Географическая долгота
	Lng NullableFloat32 `json:"lng"`
	Manager NullableEmployerAddressesEmployerAddressItemManager `json:"manager,omitempty"`
	MetroStations []IncludesMetroStation `json:"metro_stations,omitempty"`
	// Полный адрес
	Raw NullableString `json:"raw,omitempty"`
	// Улица
	Street NullableString `json:"street,omitempty"`
}

type _EmployerAddressesEmployerAddressItem EmployerAddressesEmployerAddressItem

// NewEmployerAddressesEmployerAddressItem instantiates a new EmployerAddressesEmployerAddressItem object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewEmployerAddressesEmployerAddressItem(city NullableString, lat NullableFloat32, lng NullableFloat32) *EmployerAddressesEmployerAddressItem {
	this := EmployerAddressesEmployerAddressItem{}
	this.City = city
	this.Lat = lat
	this.Lng = lng
	return &this
}

// NewEmployerAddressesEmployerAddressItemWithDefaults instantiates a new EmployerAddressesEmployerAddressItem object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewEmployerAddressesEmployerAddressItemWithDefaults() *EmployerAddressesEmployerAddressItem {
	this := EmployerAddressesEmployerAddressItem{}
	return &this
}

// GetBuilding returns the Building field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *EmployerAddressesEmployerAddressItem) GetBuilding() string {
	if o == nil || IsNil(o.Building.Get()) {
		var ret string
		return ret
	}
	return *o.Building.Get()
}

// GetBuildingOk returns a tuple with the Building field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *EmployerAddressesEmployerAddressItem) GetBuildingOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Building.Get(), o.Building.IsSet()
}

// HasBuilding returns a boolean if a field has been set.
func (o *EmployerAddressesEmployerAddressItem) HasBuilding() bool {
	if o != nil && o.Building.IsSet() {
		return true
	}

	return false
}

// SetBuilding gets a reference to the given NullableString and assigns it to the Building field.
func (o *EmployerAddressesEmployerAddressItem) SetBuilding(v string) {
	o.Building.Set(&v)
}
// SetBuildingNil sets the value for Building to be an explicit nil
func (o *EmployerAddressesEmployerAddressItem) SetBuildingNil() {
	o.Building.Set(nil)
}

// UnsetBuilding ensures that no value is present for Building, not even an explicit nil
func (o *EmployerAddressesEmployerAddressItem) UnsetBuilding() {
	o.Building.Unset()
}

// GetCanEdit returns the CanEdit field value if set, zero value otherwise.
func (o *EmployerAddressesEmployerAddressItem) GetCanEdit() bool {
	if o == nil || IsNil(o.CanEdit) {
		var ret bool
		return ret
	}
	return *o.CanEdit
}

// GetCanEditOk returns a tuple with the CanEdit field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EmployerAddressesEmployerAddressItem) GetCanEditOk() (*bool, bool) {
	if o == nil || IsNil(o.CanEdit) {
		return nil, false
	}
	return o.CanEdit, true
}

// HasCanEdit returns a boolean if a field has been set.
func (o *EmployerAddressesEmployerAddressItem) HasCanEdit() bool {
	if o != nil && !IsNil(o.CanEdit) {
		return true
	}

	return false
}

// SetCanEdit gets a reference to the given bool and assigns it to the CanEdit field.
func (o *EmployerAddressesEmployerAddressItem) SetCanEdit(v bool) {
	o.CanEdit = &v
}

// GetCity returns the City field value
// If the value is explicit nil, the zero value for string will be returned
func (o *EmployerAddressesEmployerAddressItem) GetCity() string {
	if o == nil || o.City.Get() == nil {
		var ret string
		return ret
	}

	return *o.City.Get()
}

// GetCityOk returns a tuple with the City field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *EmployerAddressesEmployerAddressItem) GetCityOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.City.Get(), o.City.IsSet()
}

// SetCity sets field value
func (o *EmployerAddressesEmployerAddressItem) SetCity(v string) {
	o.City.Set(&v)
}

// GetDeleted returns the Deleted field value if set, zero value otherwise.
func (o *EmployerAddressesEmployerAddressItem) GetDeleted() bool {
	if o == nil || IsNil(o.Deleted) {
		var ret bool
		return ret
	}
	return *o.Deleted
}

// GetDeletedOk returns a tuple with the Deleted field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EmployerAddressesEmployerAddressItem) GetDeletedOk() (*bool, bool) {
	if o == nil || IsNil(o.Deleted) {
		return nil, false
	}
	return o.Deleted, true
}

// HasDeleted returns a boolean if a field has been set.
func (o *EmployerAddressesEmployerAddressItem) HasDeleted() bool {
	if o != nil && !IsNil(o.Deleted) {
		return true
	}

	return false
}

// SetDeleted gets a reference to the given bool and assigns it to the Deleted field.
func (o *EmployerAddressesEmployerAddressItem) SetDeleted(v bool) {
	o.Deleted = &v
}

// GetDescription returns the Description field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *EmployerAddressesEmployerAddressItem) GetDescription() string {
	if o == nil || IsNil(o.Description.Get()) {
		var ret string
		return ret
	}
	return *o.Description.Get()
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *EmployerAddressesEmployerAddressItem) GetDescriptionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Description.Get(), o.Description.IsSet()
}

// HasDescription returns a boolean if a field has been set.
func (o *EmployerAddressesEmployerAddressItem) HasDescription() bool {
	if o != nil && o.Description.IsSet() {
		return true
	}

	return false
}

// SetDescription gets a reference to the given NullableString and assigns it to the Description field.
func (o *EmployerAddressesEmployerAddressItem) SetDescription(v string) {
	o.Description.Set(&v)
}
// SetDescriptionNil sets the value for Description to be an explicit nil
func (o *EmployerAddressesEmployerAddressItem) SetDescriptionNil() {
	o.Description.Set(nil)
}

// UnsetDescription ensures that no value is present for Description, not even an explicit nil
func (o *EmployerAddressesEmployerAddressItem) UnsetDescription() {
	o.Description.Unset()
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *EmployerAddressesEmployerAddressItem) GetId() string {
	if o == nil || IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EmployerAddressesEmployerAddressItem) GetIdOk() (*string, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *EmployerAddressesEmployerAddressItem) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *EmployerAddressesEmployerAddressItem) SetId(v string) {
	o.Id = &v
}

// GetLat returns the Lat field value
// If the value is explicit nil, the zero value for float32 will be returned
func (o *EmployerAddressesEmployerAddressItem) GetLat() float32 {
	if o == nil || o.Lat.Get() == nil {
		var ret float32
		return ret
	}

	return *o.Lat.Get()
}

// GetLatOk returns a tuple with the Lat field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *EmployerAddressesEmployerAddressItem) GetLatOk() (*float32, bool) {
	if o == nil {
		return nil, false
	}
	return o.Lat.Get(), o.Lat.IsSet()
}

// SetLat sets field value
func (o *EmployerAddressesEmployerAddressItem) SetLat(v float32) {
	o.Lat.Set(&v)
}

// GetLng returns the Lng field value
// If the value is explicit nil, the zero value for float32 will be returned
func (o *EmployerAddressesEmployerAddressItem) GetLng() float32 {
	if o == nil || o.Lng.Get() == nil {
		var ret float32
		return ret
	}

	return *o.Lng.Get()
}

// GetLngOk returns a tuple with the Lng field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *EmployerAddressesEmployerAddressItem) GetLngOk() (*float32, bool) {
	if o == nil {
		return nil, false
	}
	return o.Lng.Get(), o.Lng.IsSet()
}

// SetLng sets field value
func (o *EmployerAddressesEmployerAddressItem) SetLng(v float32) {
	o.Lng.Set(&v)
}

// GetManager returns the Manager field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *EmployerAddressesEmployerAddressItem) GetManager() EmployerAddressesEmployerAddressItemManager {
	if o == nil || IsNil(o.Manager.Get()) {
		var ret EmployerAddressesEmployerAddressItemManager
		return ret
	}
	return *o.Manager.Get()
}

// GetManagerOk returns a tuple with the Manager field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *EmployerAddressesEmployerAddressItem) GetManagerOk() (*EmployerAddressesEmployerAddressItemManager, bool) {
	if o == nil {
		return nil, false
	}
	return o.Manager.Get(), o.Manager.IsSet()
}

// HasManager returns a boolean if a field has been set.
func (o *EmployerAddressesEmployerAddressItem) HasManager() bool {
	if o != nil && o.Manager.IsSet() {
		return true
	}

	return false
}

// SetManager gets a reference to the given NullableEmployerAddressesEmployerAddressItemManager and assigns it to the Manager field.
func (o *EmployerAddressesEmployerAddressItem) SetManager(v EmployerAddressesEmployerAddressItemManager) {
	o.Manager.Set(&v)
}
// SetManagerNil sets the value for Manager to be an explicit nil
func (o *EmployerAddressesEmployerAddressItem) SetManagerNil() {
	o.Manager.Set(nil)
}

// UnsetManager ensures that no value is present for Manager, not even an explicit nil
func (o *EmployerAddressesEmployerAddressItem) UnsetManager() {
	o.Manager.Unset()
}

// GetMetroStations returns the MetroStations field value if set, zero value otherwise.
func (o *EmployerAddressesEmployerAddressItem) GetMetroStations() []IncludesMetroStation {
	if o == nil || IsNil(o.MetroStations) {
		var ret []IncludesMetroStation
		return ret
	}
	return o.MetroStations
}

// GetMetroStationsOk returns a tuple with the MetroStations field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EmployerAddressesEmployerAddressItem) GetMetroStationsOk() ([]IncludesMetroStation, bool) {
	if o == nil || IsNil(o.MetroStations) {
		return nil, false
	}
	return o.MetroStations, true
}

// HasMetroStations returns a boolean if a field has been set.
func (o *EmployerAddressesEmployerAddressItem) HasMetroStations() bool {
	if o != nil && !IsNil(o.MetroStations) {
		return true
	}

	return false
}

// SetMetroStations gets a reference to the given []IncludesMetroStation and assigns it to the MetroStations field.
func (o *EmployerAddressesEmployerAddressItem) SetMetroStations(v []IncludesMetroStation) {
	o.MetroStations = v
}

// GetRaw returns the Raw field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *EmployerAddressesEmployerAddressItem) GetRaw() string {
	if o == nil || IsNil(o.Raw.Get()) {
		var ret string
		return ret
	}
	return *o.Raw.Get()
}

// GetRawOk returns a tuple with the Raw field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *EmployerAddressesEmployerAddressItem) GetRawOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Raw.Get(), o.Raw.IsSet()
}

// HasRaw returns a boolean if a field has been set.
func (o *EmployerAddressesEmployerAddressItem) HasRaw() bool {
	if o != nil && o.Raw.IsSet() {
		return true
	}

	return false
}

// SetRaw gets a reference to the given NullableString and assigns it to the Raw field.
func (o *EmployerAddressesEmployerAddressItem) SetRaw(v string) {
	o.Raw.Set(&v)
}
// SetRawNil sets the value for Raw to be an explicit nil
func (o *EmployerAddressesEmployerAddressItem) SetRawNil() {
	o.Raw.Set(nil)
}

// UnsetRaw ensures that no value is present for Raw, not even an explicit nil
func (o *EmployerAddressesEmployerAddressItem) UnsetRaw() {
	o.Raw.Unset()
}

// GetStreet returns the Street field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *EmployerAddressesEmployerAddressItem) GetStreet() string {
	if o == nil || IsNil(o.Street.Get()) {
		var ret string
		return ret
	}
	return *o.Street.Get()
}

// GetStreetOk returns a tuple with the Street field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *EmployerAddressesEmployerAddressItem) GetStreetOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Street.Get(), o.Street.IsSet()
}

// HasStreet returns a boolean if a field has been set.
func (o *EmployerAddressesEmployerAddressItem) HasStreet() bool {
	if o != nil && o.Street.IsSet() {
		return true
	}

	return false
}

// SetStreet gets a reference to the given NullableString and assigns it to the Street field.
func (o *EmployerAddressesEmployerAddressItem) SetStreet(v string) {
	o.Street.Set(&v)
}
// SetStreetNil sets the value for Street to be an explicit nil
func (o *EmployerAddressesEmployerAddressItem) SetStreetNil() {
	o.Street.Set(nil)
}

// UnsetStreet ensures that no value is present for Street, not even an explicit nil
func (o *EmployerAddressesEmployerAddressItem) UnsetStreet() {
	o.Street.Unset()
}

func (o EmployerAddressesEmployerAddressItem) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o EmployerAddressesEmployerAddressItem) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.Building.IsSet() {
		toSerialize["building"] = o.Building.Get()
	}
	if !IsNil(o.CanEdit) {
		toSerialize["can_edit"] = o.CanEdit
	}
	toSerialize["city"] = o.City.Get()
	if !IsNil(o.Deleted) {
		toSerialize["deleted"] = o.Deleted
	}
	if o.Description.IsSet() {
		toSerialize["description"] = o.Description.Get()
	}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	toSerialize["lat"] = o.Lat.Get()
	toSerialize["lng"] = o.Lng.Get()
	if o.Manager.IsSet() {
		toSerialize["manager"] = o.Manager.Get()
	}
	if !IsNil(o.MetroStations) {
		toSerialize["metro_stations"] = o.MetroStations
	}
	if o.Raw.IsSet() {
		toSerialize["raw"] = o.Raw.Get()
	}
	if o.Street.IsSet() {
		toSerialize["street"] = o.Street.Get()
	}
	return toSerialize, nil
}

func (o *EmployerAddressesEmployerAddressItem) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"city",
		"lat",
		"lng",
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

	varEmployerAddressesEmployerAddressItem := _EmployerAddressesEmployerAddressItem{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varEmployerAddressesEmployerAddressItem)

	if err != nil {
		return err
	}

	*o = EmployerAddressesEmployerAddressItem(varEmployerAddressesEmployerAddressItem)

	return err
}

type NullableEmployerAddressesEmployerAddressItem struct {
	value *EmployerAddressesEmployerAddressItem
	isSet bool
}

func (v NullableEmployerAddressesEmployerAddressItem) Get() *EmployerAddressesEmployerAddressItem {
	return v.value
}

func (v *NullableEmployerAddressesEmployerAddressItem) Set(val *EmployerAddressesEmployerAddressItem) {
	v.value = val
	v.isSet = true
}

func (v NullableEmployerAddressesEmployerAddressItem) IsSet() bool {
	return v.isSet
}

func (v *NullableEmployerAddressesEmployerAddressItem) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEmployerAddressesEmployerAddressItem(val *EmployerAddressesEmployerAddressItem) *NullableEmployerAddressesEmployerAddressItem {
	return &NullableEmployerAddressesEmployerAddressItem{value: val, isSet: true}
}

func (v NullableEmployerAddressesEmployerAddressItem) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEmployerAddressesEmployerAddressItem) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


