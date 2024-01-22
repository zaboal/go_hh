/*
HeadHunter API

По-русски | [Switch to English](https://api.hh.ru/openapi/en/redoc)  В OpenAPI ведется пока что только небольшая часть документации [Основная документация](https://github.com/hhru/api).  Для поиска по документации можно использовать Ctrl+F.  # Авторизация  API поддерживает следующие уровни авторизации:   - [авторизация приложения](#section/Avtorizaciya/Avtorizaciya-prilozheniya)   - [авторизация пользователя](#section/Avtorizaciya/Avtorizaciya-polzovatelya)  * [Авторизация пользователя](#section/Avtorizaciya/Avtorizaciya-polzovatelya)   * [Правила формирования специального redirect_uri](#section/Avtorizaciya/Pravila-formirovaniya-specialnogo-redirect_uri)   * [Процесс авторизации](#section/Avtorizaciya/Process-avtorizacii)   * [Успешное получение временного `authorization_code`](#get-authorization_code)   * [Получение access и refresh токенов](#section/Avtorizaciya/Poluchenie-access-i-refresh-tokenov) * [Обновление пары access и refresh токенов](#section/Avtorizaciya/Obnovlenie-pary-access-i-refresh-tokenov) * [Инвалидация токена](#tag/Avtorizaciya-rabotodatelya/operation/invalidate-token) * [Запрос авторизации под другим пользователем](#section/Avtorizaciya/Zapros-avtorizacii-pod-drugim-polzovatelem) * [Авторизация под разными рабочими аккаунтами](#section/Avtorizaciya/Avtorizaciya-pod-raznymi-rabochimi-akkauntami) * [Авторизация приложения](#section/Avtorizaciya/Avtorizaciya-prilozheniya)       ## Авторизация пользователя Для выполнения запросов от имени пользователя необходимо пользоваться токеном пользователя.  В начале приложению необходимо направить пользователя (открыть страницу) по адресу:  ``` https://hh.ru/oauth/authorize? response_type=code& client_id={client_id}& state={state}& redirect_uri={redirect_uri} ```  Обязательные параметры:  * `response_type=code` — указание на способ получения авторизации, используя `authorization code` * `client_id` — идентификатор, полученный при создании приложения   Необязательные параметры:  * `state` — в случае указания, будет включен в ответный редирект. Это позволяет исключить возможность взлома путём подделки межсайтовых запросов. Подробнее об этом: [RFC 6749. Section 10.12](http://tools.ietf.org/html/rfc6749#section-10.12) * `redirect_uri` — uri для перенаправления пользователя после авторизации. Если не указать, используется из настроек приложения. При наличии происходит валидация значения. Вероятнее всего, потребуется сделать urlencode значения параметра.  ## Правила формирования специального redirect_uri  К примеру, если в настройках сохранен `http://example.com/oauth`, то разрешено указывать:  * `http://www.example.com/oauth` — поддомен; * `http://www.example.com/oauth/sub/path` — уточнение пути; * `http://example.com/oauth?lang=RU` — дополнительный параметр; * `http://www.example.com/oauth/sub/path?lang=RU` — всё вместе.  Запрещено:  * `https://example.com/oauth` — различные протоколы; * `http://wwwexample.com/oauth` — различные домены; * `http://wwwexample.com/` — другой путь; * `http://example.com/oauths` — другой путь; * `http://example.com:80/oauths` — указание изначально отсутствующего порта;  ## Процесс авторизации  Если пользователь не авторизован на сайте, ему будет показана форма авторизации на сайте. После прохождения авторизации на сайте, пользователю будет выведена форма с запросом разрешения доступа вашего приложения к его персональным данным.  Если пользователь не разрешает доступ приложению, пользователь будет перенаправлен на указанный `redirect_uri` с `?error=access_denied` и `state={state}`, если таковой был указан при первом запросе.  <a name=\"get-authorization_code\"></a> ### Успешное получение временного `authorization_code`  В случае разрешения прав, в редиректе будет указан временный `authorization_code`:  ```http HTTP/1.1 302 FOUND Location: {redirect_uri}?code={authorization_code} ```  Если пользователь авторизован на сайте и доступ данному приложению однажды ранее выдан, ответом будет сразу вышеописанный редирект с `authorization_code` (без показа формы логина и выдачи прав).  ## Получение access и refresh токенов  После получения `authorization_code` приложению необходимо сделать сервер-сервер запрос `POST https://hh.ru/oauth/token` для обмена полученного `authorization_code` на `access_token`.  В теле запроса необходимо передать [дополнительные параметры](#required_parameters).  Тело запроса необходимо передавать в стандартном `application/x-www-form-urlencoded` с указанием соответствующего заголовка `Content-Type`.  `authorization_code` имеет довольно короткий срок жизни, при его истечении необходимо запросить новый.  ## Обновление пары access и refresh токенов `access_token` также имеет срок жизни (ключ `expires_in`, в секундах), при его истечении приложение должно сделать запрос с `refresh_token` для получения нового.  Запрос необходимо делать в `application/x-www-form-urlencoded`.  ``` POST https://hh.ru/oauth/token ```  В теле запроса необходимо передать [дополнительные параметры](#required_parameters)  `refresh_token` можно использовать только один раз и только по истечению срока действия `access_token`.  После получения новой пары access и refresh токенов, их необходимо использовать в дальнейших запросах в api и запросах на продление токена.  ## Запрос авторизации под другим пользователем  Возможен следующий сценарий:  1. Приложение перенаправляет пользователя на сайт с запросом авторизации. 2. Пользователь на сайте уже авторизован и данному приложение доступ уже был разрешен. 3. Пользователю будет предложена возможность продолжить работу под текущим аккаунтом, либо зайти под другим аккаунтом.  Если есть необходимость, чтобы на шаге 3 сразу происходило перенаправление (redirect) с временным токеном, необходимо добавить к запросу `/oauth/authorize...` параметр `skip_choose_account=true`. В этом случае автоматически выдаётся доступ пользователю авторизованному на сайте.  Если есть необходимость всегда показывать форму авторизации, приложение может добавить к запросу `/oauth/authorize...` параметр `force_login=true`. В этом случае, пользователю будет показана форма авторизации с логином и паролем даже в случае, если пользователь уже авторизован.  Это может быть полезно приложениям, которые предоставляют сервис только для соискателей. Если пришел пользователь-работодатель, приложение может предложить пользователю повторно разрешить доступ на сайте, уже указав другую учетную запись.  Также, после авторизации приложение может показать пользователю сообщение:  ``` Вы вошли как %Имя_Фамилия%. Это не вы? ``` и предоставить ссылку с `force_login=true` для возможности захода под другим логином.  ## Авторизация под разными рабочими аккаунтами  Для получения списка рабочих аккаунтов менеджера и для работы под разными рабочими аккаунтами менеджера необходимо прочитать документацию по [рабочим аккаунтам менеджера](#tag/Menedzhery-rabotodatelya/operation/get-manager-accounts)  ## Авторизация приложения  Токен приложения необходимо сгенерировать 1 раз. В случае, если токен был скомпрометирован, его нужно запросить еще раз. При этом ранее выданный токен отзывается. Владелец приложения может посмотреть актуальный `access_token` для приложения на сайте [https://dev.hh.ru/admin](https://dev.hh.ru/admin). В случае, если вы еще ни разу [не получали токен приложения](#section/Avtorizaciya/Avtorizaciya-prilozheniya), токен отображаться не будет.  <a name=\"get-client-token\"></a> ### Получение токена приложения Для получения `access_token` необходимо сделать запрос:  ``` POST https://hh.ru/oauth/token ```  В теле запроса необходимо передать [дополнительные параметры](#required_parameters). Тело запроса необходимо передавать в стандартном `application/x-www-form-urlencoded` с указанием соответствующего заголовка `Content-Type`.  Данный `access_token` имеет **неограниченный** срок жизни. При повторном запросе ранее выданный токен отзывается и выдается новый. Запрашивать `access_token` можно не чаще, чем один раз в 5 минут.  В случае компрометации токена необходимо [инвалидировать](#tag/Avtorizaciya-rabotodatelya/operation/invalidate-token) скомпроментированный токен и запросить токен заново!  <!-- ReDoc-Inject: <security-definitions> --> 

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

// checks if the ResumesAccessTypesItem type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ResumesAccessTypesItem{}

// ResumesAccessTypesItem struct for ResumesAccessTypesItem
type ResumesAccessTypesItem struct {
	// Выбран ли тип видимости
	Active NullableBool `json:"active,omitempty"`
	// Идентификатор типа видимости
	Id string `json:"id"`
	// Максимальное количество компаний в списке видимости. Возвращается только для типов `blacklist` и `whitelist`
	Limit NullableFloat32 `json:"limit,omitempty"`
	// Ссылка на список видимости. Возвращается только для типов `blacklist` и `whitelist`
	ListUrl NullableString `json:"list_url,omitempty"`
	// Имя типа видимости
	Name string `json:"name"`
	// Количество компаний, добавленных в соответствующий список видимости. Возвращается только для типов `blacklist` и `whitelist`
	Total NullableFloat32 `json:"total,omitempty"`
}

type _ResumesAccessTypesItem ResumesAccessTypesItem

// NewResumesAccessTypesItem instantiates a new ResumesAccessTypesItem object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewResumesAccessTypesItem(id string, name string) *ResumesAccessTypesItem {
	this := ResumesAccessTypesItem{}
	this.Id = id
	this.Name = name
	return &this
}

// NewResumesAccessTypesItemWithDefaults instantiates a new ResumesAccessTypesItem object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewResumesAccessTypesItemWithDefaults() *ResumesAccessTypesItem {
	this := ResumesAccessTypesItem{}
	return &this
}

// GetActive returns the Active field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ResumesAccessTypesItem) GetActive() bool {
	if o == nil || IsNil(o.Active.Get()) {
		var ret bool
		return ret
	}
	return *o.Active.Get()
}

// GetActiveOk returns a tuple with the Active field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ResumesAccessTypesItem) GetActiveOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return o.Active.Get(), o.Active.IsSet()
}

// HasActive returns a boolean if a field has been set.
func (o *ResumesAccessTypesItem) HasActive() bool {
	if o != nil && o.Active.IsSet() {
		return true
	}

	return false
}

// SetActive gets a reference to the given NullableBool and assigns it to the Active field.
func (o *ResumesAccessTypesItem) SetActive(v bool) {
	o.Active.Set(&v)
}
// SetActiveNil sets the value for Active to be an explicit nil
func (o *ResumesAccessTypesItem) SetActiveNil() {
	o.Active.Set(nil)
}

// UnsetActive ensures that no value is present for Active, not even an explicit nil
func (o *ResumesAccessTypesItem) UnsetActive() {
	o.Active.Unset()
}

// GetId returns the Id field value
func (o *ResumesAccessTypesItem) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *ResumesAccessTypesItem) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *ResumesAccessTypesItem) SetId(v string) {
	o.Id = v
}

// GetLimit returns the Limit field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ResumesAccessTypesItem) GetLimit() float32 {
	if o == nil || IsNil(o.Limit.Get()) {
		var ret float32
		return ret
	}
	return *o.Limit.Get()
}

// GetLimitOk returns a tuple with the Limit field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ResumesAccessTypesItem) GetLimitOk() (*float32, bool) {
	if o == nil {
		return nil, false
	}
	return o.Limit.Get(), o.Limit.IsSet()
}

// HasLimit returns a boolean if a field has been set.
func (o *ResumesAccessTypesItem) HasLimit() bool {
	if o != nil && o.Limit.IsSet() {
		return true
	}

	return false
}

// SetLimit gets a reference to the given NullableFloat32 and assigns it to the Limit field.
func (o *ResumesAccessTypesItem) SetLimit(v float32) {
	o.Limit.Set(&v)
}
// SetLimitNil sets the value for Limit to be an explicit nil
func (o *ResumesAccessTypesItem) SetLimitNil() {
	o.Limit.Set(nil)
}

// UnsetLimit ensures that no value is present for Limit, not even an explicit nil
func (o *ResumesAccessTypesItem) UnsetLimit() {
	o.Limit.Unset()
}

// GetListUrl returns the ListUrl field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ResumesAccessTypesItem) GetListUrl() string {
	if o == nil || IsNil(o.ListUrl.Get()) {
		var ret string
		return ret
	}
	return *o.ListUrl.Get()
}

// GetListUrlOk returns a tuple with the ListUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ResumesAccessTypesItem) GetListUrlOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.ListUrl.Get(), o.ListUrl.IsSet()
}

// HasListUrl returns a boolean if a field has been set.
func (o *ResumesAccessTypesItem) HasListUrl() bool {
	if o != nil && o.ListUrl.IsSet() {
		return true
	}

	return false
}

// SetListUrl gets a reference to the given NullableString and assigns it to the ListUrl field.
func (o *ResumesAccessTypesItem) SetListUrl(v string) {
	o.ListUrl.Set(&v)
}
// SetListUrlNil sets the value for ListUrl to be an explicit nil
func (o *ResumesAccessTypesItem) SetListUrlNil() {
	o.ListUrl.Set(nil)
}

// UnsetListUrl ensures that no value is present for ListUrl, not even an explicit nil
func (o *ResumesAccessTypesItem) UnsetListUrl() {
	o.ListUrl.Unset()
}

// GetName returns the Name field value
func (o *ResumesAccessTypesItem) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *ResumesAccessTypesItem) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *ResumesAccessTypesItem) SetName(v string) {
	o.Name = v
}

// GetTotal returns the Total field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ResumesAccessTypesItem) GetTotal() float32 {
	if o == nil || IsNil(o.Total.Get()) {
		var ret float32
		return ret
	}
	return *o.Total.Get()
}

// GetTotalOk returns a tuple with the Total field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ResumesAccessTypesItem) GetTotalOk() (*float32, bool) {
	if o == nil {
		return nil, false
	}
	return o.Total.Get(), o.Total.IsSet()
}

// HasTotal returns a boolean if a field has been set.
func (o *ResumesAccessTypesItem) HasTotal() bool {
	if o != nil && o.Total.IsSet() {
		return true
	}

	return false
}

// SetTotal gets a reference to the given NullableFloat32 and assigns it to the Total field.
func (o *ResumesAccessTypesItem) SetTotal(v float32) {
	o.Total.Set(&v)
}
// SetTotalNil sets the value for Total to be an explicit nil
func (o *ResumesAccessTypesItem) SetTotalNil() {
	o.Total.Set(nil)
}

// UnsetTotal ensures that no value is present for Total, not even an explicit nil
func (o *ResumesAccessTypesItem) UnsetTotal() {
	o.Total.Unset()
}

func (o ResumesAccessTypesItem) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ResumesAccessTypesItem) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.Active.IsSet() {
		toSerialize["active"] = o.Active.Get()
	}
	toSerialize["id"] = o.Id
	if o.Limit.IsSet() {
		toSerialize["limit"] = o.Limit.Get()
	}
	if o.ListUrl.IsSet() {
		toSerialize["list_url"] = o.ListUrl.Get()
	}
	toSerialize["name"] = o.Name
	if o.Total.IsSet() {
		toSerialize["total"] = o.Total.Get()
	}
	return toSerialize, nil
}

func (o *ResumesAccessTypesItem) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"id",
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

	varResumesAccessTypesItem := _ResumesAccessTypesItem{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varResumesAccessTypesItem)

	if err != nil {
		return err
	}

	*o = ResumesAccessTypesItem(varResumesAccessTypesItem)

	return err
}

type NullableResumesAccessTypesItem struct {
	value *ResumesAccessTypesItem
	isSet bool
}

func (v NullableResumesAccessTypesItem) Get() *ResumesAccessTypesItem {
	return v.value
}

func (v *NullableResumesAccessTypesItem) Set(val *ResumesAccessTypesItem) {
	v.value = val
	v.isSet = true
}

func (v NullableResumesAccessTypesItem) IsSet() bool {
	return v.isSet
}

func (v *NullableResumesAccessTypesItem) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableResumesAccessTypesItem(val *ResumesAccessTypesItem) *NullableResumesAccessTypesItem {
	return &NullableResumesAccessTypesItem{value: val, isSet: true}
}

func (v NullableResumesAccessTypesItem) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableResumesAccessTypesItem) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


