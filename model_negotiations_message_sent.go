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

// checks if the NegotiationsMessageSent type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &NegotiationsMessageSent{}

// NegotiationsMessageSent struct for NegotiationsMessageSent
type NegotiationsMessageSent struct {
	Address NullableVacancyAddressOutput `json:"address,omitempty"`
	Author NegotiationsAuthor `json:"author"`
	// Дата отправки сообщения
	CreatedAt string `json:"created_at"`
	// Можно ли редактировать сообщение
	Editable bool `json:"editable"`
	// Идентификатор сообщения
	Id string `json:"id"`
	// Можно ли прочитать сообщение
	Read *bool `json:"read,omitempty"`
	// Состояние сообщения
	State IncludesIdName `json:"state"`
	// Текст сообщения
	Text string `json:"text"`
	// Просмотрено ли сообщение отправителем
	ViewedByMe bool `json:"viewed_by_me"`
	// Просмотрено ли сообщение получателем
	ViewedByOpponent bool `json:"viewed_by_opponent"`
}

type _NegotiationsMessageSent NegotiationsMessageSent

// NewNegotiationsMessageSent instantiates a new NegotiationsMessageSent object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewNegotiationsMessageSent(author NegotiationsAuthor, createdAt string, editable bool, id string, state IncludesIdName, text string, viewedByMe bool, viewedByOpponent bool) *NegotiationsMessageSent {
	this := NegotiationsMessageSent{}
	this.Author = author
	this.CreatedAt = createdAt
	this.Editable = editable
	this.Id = id
	this.State = state
	this.Text = text
	this.ViewedByMe = viewedByMe
	this.ViewedByOpponent = viewedByOpponent
	return &this
}

// NewNegotiationsMessageSentWithDefaults instantiates a new NegotiationsMessageSent object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewNegotiationsMessageSentWithDefaults() *NegotiationsMessageSent {
	this := NegotiationsMessageSent{}
	return &this
}

// GetAddress returns the Address field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *NegotiationsMessageSent) GetAddress() VacancyAddressOutput {
	if o == nil || IsNil(o.Address.Get()) {
		var ret VacancyAddressOutput
		return ret
	}
	return *o.Address.Get()
}

// GetAddressOk returns a tuple with the Address field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *NegotiationsMessageSent) GetAddressOk() (*VacancyAddressOutput, bool) {
	if o == nil {
		return nil, false
	}
	return o.Address.Get(), o.Address.IsSet()
}

// HasAddress returns a boolean if a field has been set.
func (o *NegotiationsMessageSent) HasAddress() bool {
	if o != nil && o.Address.IsSet() {
		return true
	}

	return false
}

// SetAddress gets a reference to the given NullableVacancyAddressOutput and assigns it to the Address field.
func (o *NegotiationsMessageSent) SetAddress(v VacancyAddressOutput) {
	o.Address.Set(&v)
}
// SetAddressNil sets the value for Address to be an explicit nil
func (o *NegotiationsMessageSent) SetAddressNil() {
	o.Address.Set(nil)
}

// UnsetAddress ensures that no value is present for Address, not even an explicit nil
func (o *NegotiationsMessageSent) UnsetAddress() {
	o.Address.Unset()
}

// GetAuthor returns the Author field value
func (o *NegotiationsMessageSent) GetAuthor() NegotiationsAuthor {
	if o == nil {
		var ret NegotiationsAuthor
		return ret
	}

	return o.Author
}

// GetAuthorOk returns a tuple with the Author field value
// and a boolean to check if the value has been set.
func (o *NegotiationsMessageSent) GetAuthorOk() (*NegotiationsAuthor, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Author, true
}

// SetAuthor sets field value
func (o *NegotiationsMessageSent) SetAuthor(v NegotiationsAuthor) {
	o.Author = v
}

// GetCreatedAt returns the CreatedAt field value
func (o *NegotiationsMessageSent) GetCreatedAt() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value
// and a boolean to check if the value has been set.
func (o *NegotiationsMessageSent) GetCreatedAtOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CreatedAt, true
}

// SetCreatedAt sets field value
func (o *NegotiationsMessageSent) SetCreatedAt(v string) {
	o.CreatedAt = v
}

// GetEditable returns the Editable field value
func (o *NegotiationsMessageSent) GetEditable() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.Editable
}

// GetEditableOk returns a tuple with the Editable field value
// and a boolean to check if the value has been set.
func (o *NegotiationsMessageSent) GetEditableOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Editable, true
}

// SetEditable sets field value
func (o *NegotiationsMessageSent) SetEditable(v bool) {
	o.Editable = v
}

// GetId returns the Id field value
func (o *NegotiationsMessageSent) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *NegotiationsMessageSent) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *NegotiationsMessageSent) SetId(v string) {
	o.Id = v
}

// GetRead returns the Read field value if set, zero value otherwise.
func (o *NegotiationsMessageSent) GetRead() bool {
	if o == nil || IsNil(o.Read) {
		var ret bool
		return ret
	}
	return *o.Read
}

// GetReadOk returns a tuple with the Read field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NegotiationsMessageSent) GetReadOk() (*bool, bool) {
	if o == nil || IsNil(o.Read) {
		return nil, false
	}
	return o.Read, true
}

// HasRead returns a boolean if a field has been set.
func (o *NegotiationsMessageSent) HasRead() bool {
	if o != nil && !IsNil(o.Read) {
		return true
	}

	return false
}

// SetRead gets a reference to the given bool and assigns it to the Read field.
func (o *NegotiationsMessageSent) SetRead(v bool) {
	o.Read = &v
}

// GetState returns the State field value
func (o *NegotiationsMessageSent) GetState() IncludesIdName {
	if o == nil {
		var ret IncludesIdName
		return ret
	}

	return o.State
}

// GetStateOk returns a tuple with the State field value
// and a boolean to check if the value has been set.
func (o *NegotiationsMessageSent) GetStateOk() (*IncludesIdName, bool) {
	if o == nil {
		return nil, false
	}
	return &o.State, true
}

// SetState sets field value
func (o *NegotiationsMessageSent) SetState(v IncludesIdName) {
	o.State = v
}

// GetText returns the Text field value
func (o *NegotiationsMessageSent) GetText() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Text
}

// GetTextOk returns a tuple with the Text field value
// and a boolean to check if the value has been set.
func (o *NegotiationsMessageSent) GetTextOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Text, true
}

// SetText sets field value
func (o *NegotiationsMessageSent) SetText(v string) {
	o.Text = v
}

// GetViewedByMe returns the ViewedByMe field value
func (o *NegotiationsMessageSent) GetViewedByMe() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.ViewedByMe
}

// GetViewedByMeOk returns a tuple with the ViewedByMe field value
// and a boolean to check if the value has been set.
func (o *NegotiationsMessageSent) GetViewedByMeOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ViewedByMe, true
}

// SetViewedByMe sets field value
func (o *NegotiationsMessageSent) SetViewedByMe(v bool) {
	o.ViewedByMe = v
}

// GetViewedByOpponent returns the ViewedByOpponent field value
func (o *NegotiationsMessageSent) GetViewedByOpponent() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.ViewedByOpponent
}

// GetViewedByOpponentOk returns a tuple with the ViewedByOpponent field value
// and a boolean to check if the value has been set.
func (o *NegotiationsMessageSent) GetViewedByOpponentOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ViewedByOpponent, true
}

// SetViewedByOpponent sets field value
func (o *NegotiationsMessageSent) SetViewedByOpponent(v bool) {
	o.ViewedByOpponent = v
}

func (o NegotiationsMessageSent) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o NegotiationsMessageSent) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.Address.IsSet() {
		toSerialize["address"] = o.Address.Get()
	}
	toSerialize["author"] = o.Author
	toSerialize["created_at"] = o.CreatedAt
	toSerialize["editable"] = o.Editable
	toSerialize["id"] = o.Id
	if !IsNil(o.Read) {
		toSerialize["read"] = o.Read
	}
	toSerialize["state"] = o.State
	toSerialize["text"] = o.Text
	toSerialize["viewed_by_me"] = o.ViewedByMe
	toSerialize["viewed_by_opponent"] = o.ViewedByOpponent
	return toSerialize, nil
}

func (o *NegotiationsMessageSent) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"author",
		"created_at",
		"editable",
		"id",
		"state",
		"text",
		"viewed_by_me",
		"viewed_by_opponent",
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

	varNegotiationsMessageSent := _NegotiationsMessageSent{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varNegotiationsMessageSent)

	if err != nil {
		return err
	}

	*o = NegotiationsMessageSent(varNegotiationsMessageSent)

	return err
}

type NullableNegotiationsMessageSent struct {
	value *NegotiationsMessageSent
	isSet bool
}

func (v NullableNegotiationsMessageSent) Get() *NegotiationsMessageSent {
	return v.value
}

func (v *NullableNegotiationsMessageSent) Set(val *NegotiationsMessageSent) {
	v.value = val
	v.isSet = true
}

func (v NullableNegotiationsMessageSent) IsSet() bool {
	return v.isSet
}

func (v *NullableNegotiationsMessageSent) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableNegotiationsMessageSent(val *NegotiationsMessageSent) *NullableNegotiationsMessageSent {
	return &NullableNegotiationsMessageSent{value: val, isSet: true}
}

func (v NullableNegotiationsMessageSent) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableNegotiationsMessageSent) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


