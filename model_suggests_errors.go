/*
HeadHunter API

По-русски | [Switch to English](https://api.hh.ru/openapi/en/redoc)  В OpenAPI ведется пока что только небольшая часть документации [Основная документация](https://github.com/hhru/api).  Для поиска по документации можно использовать Ctrl+F.  # Авторизация  API поддерживает следующие уровни авторизации:   - [авторизация приложения](#section/Avtorizaciya/Avtorizaciya-prilozheniya)   - [авторизация пользователя](#section/Avtorizaciya/Avtorizaciya-polzovatelya)  * [Авторизация пользователя](#section/Avtorizaciya/Avtorizaciya-polzovatelya)   * [Правила формирования специального redirect_uri](#section/Avtorizaciya/Pravila-formirovaniya-specialnogo-redirect_uri)   * [Процесс авторизации](#section/Avtorizaciya/Process-avtorizacii)   * [Успешное получение временного `authorization_code`](#get-authorization_code)   * [Получение access и refresh токенов](#section/Avtorizaciya/Poluchenie-access-i-refresh-tokenov) * [Обновление пары access и refresh токенов](#section/Avtorizaciya/Obnovlenie-pary-access-i-refresh-tokenov) * [Инвалидация токена](#tag/Avtorizaciya-rabotodatelya/operation/invalidate-token) * [Запрос авторизации под другим пользователем](#section/Avtorizaciya/Zapros-avtorizacii-pod-drugim-polzovatelem) * [Авторизация под разными рабочими аккаунтами](#section/Avtorizaciya/Avtorizaciya-pod-raznymi-rabochimi-akkauntami) * [Авторизация приложения](#section/Avtorizaciya/Avtorizaciya-prilozheniya)       ## Авторизация пользователя Для выполнения запросов от имени пользователя необходимо пользоваться токеном пользователя.  В начале приложению необходимо направить пользователя (открыть страницу) по адресу:  ``` https://hh.ru/oauth/authorize? response_type=code& client_id={client_id}& state={state}& redirect_uri={redirect_uri} ```  Обязательные параметры:  * `response_type=code` — указание на способ получения авторизации, используя `authorization code` * `client_id` — идентификатор, полученный при создании приложения   Необязательные параметры:  * `state` — в случае указания, будет включен в ответный редирект. Это позволяет исключить возможность взлома путём подделки межсайтовых запросов. Подробнее об этом: [RFC 6749. Section 10.12](http://tools.ietf.org/html/rfc6749#section-10.12) * `redirect_uri` — uri для перенаправления пользователя после авторизации. Если не указать, используется из настроек приложения. При наличии происходит валидация значения. Вероятнее всего, потребуется сделать urlencode значения параметра.  ## Правила формирования специального redirect_uri  К примеру, если в настройках сохранен `http://example.com/oauth`, то разрешено указывать:  * `http://www.example.com/oauth` — поддомен; * `http://www.example.com/oauth/sub/path` — уточнение пути; * `http://example.com/oauth?lang=RU` — дополнительный параметр; * `http://www.example.com/oauth/sub/path?lang=RU` — всё вместе.  Запрещено:  * `https://example.com/oauth` — различные протоколы; * `http://wwwexample.com/oauth` — различные домены; * `http://wwwexample.com/` — другой путь; * `http://example.com/oauths` — другой путь; * `http://example.com:80/oauths` — указание изначально отсутствующего порта;  ## Процесс авторизации  Если пользователь не авторизован на сайте, ему будет показана форма авторизации на сайте. После прохождения авторизации на сайте, пользователю будет выведена форма с запросом разрешения доступа вашего приложения к его персональным данным.  Если пользователь не разрешает доступ приложению, пользователь будет перенаправлен на указанный `redirect_uri` с `?error=access_denied` и `state={state}`, если таковой был указан при первом запросе.  <a name=\"get-authorization_code\"></a> ### Успешное получение временного `authorization_code`  В случае разрешения прав, в редиректе будет указан временный `authorization_code`:  ```http HTTP/1.1 302 FOUND Location: {redirect_uri}?code={authorization_code} ```  Если пользователь авторизован на сайте и доступ данному приложению однажды ранее выдан, ответом будет сразу вышеописанный редирект с `authorization_code` (без показа формы логина и выдачи прав).  ## Получение access и refresh токенов  После получения `authorization_code` приложению необходимо сделать сервер-сервер запрос `POST https://api.hh.ru/token` для обмена полученного `authorization_code` на `access_token` (старый запрос `POST https://hh.ru/oauth/token` считается устаревшим).  В теле запроса необходимо передать [дополнительные параметры](#required_parameters).  Тело запроса необходимо передавать в стандартном `application/x-www-form-urlencoded` с указанием соответствующего заголовка `Content-Type`.  `authorization_code` имеет довольно короткий срок жизни, при его истечении необходимо запросить новый.  ## Обновление пары access и refresh токенов `access_token` также имеет срок жизни (ключ `expires_in`, в секундах), при его истечении приложение должно сделать запрос с `refresh_token` для получения нового.  Запрос необходимо делать в `application/x-www-form-urlencoded`.  ``` POST https://api.hh.ru/token ```  (старый запрос `POST https://hh.ru/oauth/token` считается устаревшим)  В теле запроса необходимо передать [дополнительные параметры](#required_parameters)  `refresh_token` можно использовать только один раз и только по истечению срока действия `access_token`.  После получения новой пары access и refresh токенов, их необходимо использовать в дальнейших запросах в api и запросах на продление токена.  ## Запрос авторизации под другим пользователем  Возможен следующий сценарий:  1. Приложение перенаправляет пользователя на сайт с запросом авторизации. 2. Пользователь на сайте уже авторизован и данному приложение доступ уже был разрешен. 3. Пользователю будет предложена возможность продолжить работу под текущим аккаунтом, либо зайти под другим аккаунтом.  Если есть необходимость, чтобы на шаге 3 сразу происходило перенаправление (redirect) с временным токеном, необходимо добавить к запросу `/oauth/authorize...` параметр `skip_choose_account=true`. В этом случае автоматически выдаётся доступ пользователю авторизованному на сайте.  Если есть необходимость всегда показывать форму авторизации, приложение может добавить к запросу `/oauth/authorize...` параметр `force_login=true`. В этом случае, пользователю будет показана форма авторизации с логином и паролем даже в случае, если пользователь уже авторизован.  Это может быть полезно приложениям, которые предоставляют сервис только для соискателей. Если пришел пользователь-работодатель, приложение может предложить пользователю повторно разрешить доступ на сайте, уже указав другую учетную запись.  Также, после авторизации приложение может показать пользователю сообщение:  ``` Вы вошли как %Имя_Фамилия%. Это не вы? ``` и предоставить ссылку с `force_login=true` для возможности захода под другим логином.  ## Авторизация под разными рабочими аккаунтами  Для получения списка рабочих аккаунтов менеджера и для работы под разными рабочими аккаунтами менеджера необходимо прочитать документацию по [рабочим аккаунтам менеджера](#tag/Menedzhery-rabotodatelya/operation/get-manager-accounts)  ## Авторизация приложения  Токен приложения необходимо сгенерировать 1 раз. В случае, если токен был скомпрометирован, его нужно запросить еще раз. При этом ранее выданный токен отзывается. Владелец приложения может посмотреть актуальный `access_token` для приложения на сайте [https://dev.hh.ru/admin](https://dev.hh.ru/admin). В случае, если вы еще ни разу [не получали токен приложения](#section/Avtorizaciya/Avtorizaciya-prilozheniya), токен отображаться не будет.  <a name=\"get-client-token\"></a> ### Получение токена приложения Для получения `access_token` необходимо сделать запрос:  ``` POST https://api.hh.ru/token ```  (старый запрос `POST https://hh.ru/oauth/token` считается устаревшим)  В теле запроса необходимо передать [дополнительные параметры](#required_parameters). Тело запроса необходимо передавать в стандартном `application/x-www-form-urlencoded` с указанием соответствующего заголовка `Content-Type`.  Данный `access_token` имеет **неограниченный** срок жизни. При повторном запросе ранее выданный токен отзывается и выдается новый. Запрашивать `access_token` можно не чаще, чем один раз в 5 минут.  В случае компрометации токена необходимо [инвалидировать](#tag/Avtorizaciya-rabotodatelya/operation/invalidate-token) скомпроментированный токен и запросить токен заново!  <!-- ReDoc-Inject: <security-definitions> --> 

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

// checks if the SuggestsErrors type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SuggestsErrors{}

// SuggestsErrors struct for SuggestsErrors
type SuggestsErrors struct {
	// Идентификатор запроса
	RequestId string `json:"request_id"`
	// Deprecated
	BadArgument *string `json:"bad_argument,omitempty"`
	// Deprecated
	BadArguments []SuggestsErrorsAllOfBadArguments `json:"bad_arguments,omitempty"`
	// Deprecated
	Description *string `json:"description,omitempty"`
	Errors []SuggestsErrorsAllOfErrors `json:"errors"`
}

type _SuggestsErrors SuggestsErrors

// NewSuggestsErrors instantiates a new SuggestsErrors object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSuggestsErrors(requestId string, errors []SuggestsErrorsAllOfErrors) *SuggestsErrors {
	this := SuggestsErrors{}
	this.RequestId = requestId
	this.Errors = errors
	return &this
}

// NewSuggestsErrorsWithDefaults instantiates a new SuggestsErrors object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSuggestsErrorsWithDefaults() *SuggestsErrors {
	this := SuggestsErrors{}
	return &this
}

// GetRequestId returns the RequestId field value
func (o *SuggestsErrors) GetRequestId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.RequestId
}

// GetRequestIdOk returns a tuple with the RequestId field value
// and a boolean to check if the value has been set.
func (o *SuggestsErrors) GetRequestIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.RequestId, true
}

// SetRequestId sets field value
func (o *SuggestsErrors) SetRequestId(v string) {
	o.RequestId = v
}

// GetBadArgument returns the BadArgument field value if set, zero value otherwise.
// Deprecated
func (o *SuggestsErrors) GetBadArgument() string {
	if o == nil || IsNil(o.BadArgument) {
		var ret string
		return ret
	}
	return *o.BadArgument
}

// GetBadArgumentOk returns a tuple with the BadArgument field value if set, nil otherwise
// and a boolean to check if the value has been set.
// Deprecated
func (o *SuggestsErrors) GetBadArgumentOk() (*string, bool) {
	if o == nil || IsNil(o.BadArgument) {
		return nil, false
	}
	return o.BadArgument, true
}

// HasBadArgument returns a boolean if a field has been set.
func (o *SuggestsErrors) HasBadArgument() bool {
	if o != nil && !IsNil(o.BadArgument) {
		return true
	}

	return false
}

// SetBadArgument gets a reference to the given string and assigns it to the BadArgument field.
// Deprecated
func (o *SuggestsErrors) SetBadArgument(v string) {
	o.BadArgument = &v
}

// GetBadArguments returns the BadArguments field value if set, zero value otherwise.
// Deprecated
func (o *SuggestsErrors) GetBadArguments() []SuggestsErrorsAllOfBadArguments {
	if o == nil || IsNil(o.BadArguments) {
		var ret []SuggestsErrorsAllOfBadArguments
		return ret
	}
	return o.BadArguments
}

// GetBadArgumentsOk returns a tuple with the BadArguments field value if set, nil otherwise
// and a boolean to check if the value has been set.
// Deprecated
func (o *SuggestsErrors) GetBadArgumentsOk() ([]SuggestsErrorsAllOfBadArguments, bool) {
	if o == nil || IsNil(o.BadArguments) {
		return nil, false
	}
	return o.BadArguments, true
}

// HasBadArguments returns a boolean if a field has been set.
func (o *SuggestsErrors) HasBadArguments() bool {
	if o != nil && !IsNil(o.BadArguments) {
		return true
	}

	return false
}

// SetBadArguments gets a reference to the given []SuggestsErrorsAllOfBadArguments and assigns it to the BadArguments field.
// Deprecated
func (o *SuggestsErrors) SetBadArguments(v []SuggestsErrorsAllOfBadArguments) {
	o.BadArguments = v
}

// GetDescription returns the Description field value if set, zero value otherwise.
// Deprecated
func (o *SuggestsErrors) GetDescription() string {
	if o == nil || IsNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
// Deprecated
func (o *SuggestsErrors) GetDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *SuggestsErrors) HasDescription() bool {
	if o != nil && !IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
// Deprecated
func (o *SuggestsErrors) SetDescription(v string) {
	o.Description = &v
}

// GetErrors returns the Errors field value
func (o *SuggestsErrors) GetErrors() []SuggestsErrorsAllOfErrors {
	if o == nil {
		var ret []SuggestsErrorsAllOfErrors
		return ret
	}

	return o.Errors
}

// GetErrorsOk returns a tuple with the Errors field value
// and a boolean to check if the value has been set.
func (o *SuggestsErrors) GetErrorsOk() ([]SuggestsErrorsAllOfErrors, bool) {
	if o == nil {
		return nil, false
	}
	return o.Errors, true
}

// SetErrors sets field value
func (o *SuggestsErrors) SetErrors(v []SuggestsErrorsAllOfErrors) {
	o.Errors = v
}

func (o SuggestsErrors) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SuggestsErrors) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["request_id"] = o.RequestId
	if !IsNil(o.BadArgument) {
		toSerialize["bad_argument"] = o.BadArgument
	}
	if !IsNil(o.BadArguments) {
		toSerialize["bad_arguments"] = o.BadArguments
	}
	if !IsNil(o.Description) {
		toSerialize["description"] = o.Description
	}
	toSerialize["errors"] = o.Errors
	return toSerialize, nil
}

func (o *SuggestsErrors) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"request_id",
		"errors",
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

	varSuggestsErrors := _SuggestsErrors{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varSuggestsErrors)

	if err != nil {
		return err
	}

	*o = SuggestsErrors(varSuggestsErrors)

	return err
}

type NullableSuggestsErrors struct {
	value *SuggestsErrors
	isSet bool
}

func (v NullableSuggestsErrors) Get() *SuggestsErrors {
	return v.value
}

func (v *NullableSuggestsErrors) Set(val *SuggestsErrors) {
	v.value = val
	v.isSet = true
}

func (v NullableSuggestsErrors) IsSet() bool {
	return v.isSet
}

func (v *NullableSuggestsErrors) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSuggestsErrors(val *SuggestsErrors) *NullableSuggestsErrors {
	return &NullableSuggestsErrors{value: val, isSet: true}
}

func (v NullableSuggestsErrors) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSuggestsErrors) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


