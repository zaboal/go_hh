/*
HeadHunter API

По-русски | [Switch to English](https://api.hh.ru/openapi/en/redoc)  В OpenAPI ведется пока что только небольшая часть документации [Основная документация](https://github.com/hhru/api).  Для поиска по документации можно использовать Ctrl+F.  # Авторизация  API поддерживает следующие уровни авторизации:   - [авторизация приложения](#section/Avtorizaciya/Avtorizaciya-prilozheniya)   - [авторизация пользователя](#section/Avtorizaciya/Avtorizaciya-polzovatelya)  * [Авторизация пользователя](#section/Avtorizaciya/Avtorizaciya-polzovatelya)   * [Правила формирования специального redirect_uri](#section/Avtorizaciya/Pravila-formirovaniya-specialnogo-redirect_uri)   * [Процесс авторизации](#section/Avtorizaciya/Process-avtorizacii)   * [Успешное получение временного `authorization_code`](#get-authorization_code)   * [Получение access и refresh токенов](#section/Avtorizaciya/Poluchenie-access-i-refresh-tokenov) * [Обновление пары access и refresh токенов](#section/Avtorizaciya/Obnovlenie-pary-access-i-refresh-tokenov) * [Инвалидация токена](#tag/Avtorizaciya-rabotodatelya/operation/invalidate-token) * [Запрос авторизации под другим пользователем](#section/Avtorizaciya/Zapros-avtorizacii-pod-drugim-polzovatelem) * [Авторизация под разными рабочими аккаунтами](#section/Avtorizaciya/Avtorizaciya-pod-raznymi-rabochimi-akkauntami) * [Авторизация приложения](#section/Avtorizaciya/Avtorizaciya-prilozheniya)       ## Авторизация пользователя Для выполнения запросов от имени пользователя необходимо пользоваться токеном пользователя.  В начале приложению необходимо направить пользователя (открыть страницу) по адресу:  ``` https://hh.ru/oauth/authorize? response_type=code& client_id={client_id}& state={state}& redirect_uri={redirect_uri} ```  Обязательные параметры:  * `response_type=code` — указание на способ получения авторизации, используя `authorization code` * `client_id` — идентификатор, полученный при создании приложения   Необязательные параметры:  * `state` — в случае указания, будет включен в ответный редирект. Это позволяет исключить возможность взлома путём подделки межсайтовых запросов. Подробнее об этом: [RFC 6749. Section 10.12](http://tools.ietf.org/html/rfc6749#section-10.12) * `redirect_uri` — uri для перенаправления пользователя после авторизации. Если не указать, используется из настроек приложения. При наличии происходит валидация значения. Вероятнее всего, потребуется сделать urlencode значения параметра.  ## Правила формирования специального redirect_uri  К примеру, если в настройках сохранен `http://example.com/oauth`, то разрешено указывать:  * `http://www.example.com/oauth` — поддомен; * `http://www.example.com/oauth/sub/path` — уточнение пути; * `http://example.com/oauth?lang=RU` — дополнительный параметр; * `http://www.example.com/oauth/sub/path?lang=RU` — всё вместе.  Запрещено:  * `https://example.com/oauth` — различные протоколы; * `http://wwwexample.com/oauth` — различные домены; * `http://wwwexample.com/` — другой путь; * `http://example.com/oauths` — другой путь; * `http://example.com:80/oauths` — указание изначально отсутствующего порта;  ## Процесс авторизации  Если пользователь не авторизован на сайте, ему будет показана форма авторизации на сайте. После прохождения авторизации на сайте, пользователю будет выведена форма с запросом разрешения доступа вашего приложения к его персональным данным.  Если пользователь не разрешает доступ приложению, пользователь будет перенаправлен на указанный `redirect_uri` с `?error=access_denied` и `state={state}`, если таковой был указан при первом запросе.  <a name=\"get-authorization_code\"></a> ### Успешное получение временного `authorization_code`  В случае разрешения прав, в редиректе будет указан временный `authorization_code`:  ```http HTTP/1.1 302 FOUND Location: {redirect_uri}?code={authorization_code} ```  Если пользователь авторизован на сайте и доступ данному приложению однажды ранее выдан, ответом будет сразу вышеописанный редирект с `authorization_code` (без показа формы логина и выдачи прав).  ## Получение access и refresh токенов  После получения `authorization_code` приложению необходимо сделать сервер-сервер запрос `POST https://hh.ru/oauth/token` для обмена полученного `authorization_code` на `access_token`.  В теле запроса необходимо передать [дополнительные параметры](#required_parameters).  Тело запроса необходимо передавать в стандартном `application/x-www-form-urlencoded` с указанием соответствующего заголовка `Content-Type`.  `authorization_code` имеет довольно короткий срок жизни, при его истечении необходимо запросить новый.  ## Обновление пары access и refresh токенов `access_token` также имеет срок жизни (ключ `expires_in`, в секундах), при его истечении приложение должно сделать запрос с `refresh_token` для получения нового.  Запрос необходимо делать в `application/x-www-form-urlencoded`.  ``` POST https://hh.ru/oauth/token ```  В теле запроса необходимо передать [дополнительные параметры](#required_parameters)  `refresh_token` можно использовать только один раз и только по истечению срока действия `access_token`.  После получения новой пары access и refresh токенов, их необходимо использовать в дальнейших запросах в api и запросах на продление токена.  ## Запрос авторизации под другим пользователем  Возможен следующий сценарий:  1. Приложение перенаправляет пользователя на сайт с запросом авторизации. 2. Пользователь на сайте уже авторизован и данному приложение доступ уже был разрешен. 3. Пользователю будет предложена возможность продолжить работу под текущим аккаунтом, либо зайти под другим аккаунтом.  Если есть необходимость, чтобы на шаге 3 сразу происходило перенаправление (redirect) с временным токеном, необходимо добавить к запросу `/oauth/authorize...` параметр `skip_choose_account=true`. В этом случае автоматически выдаётся доступ пользователю авторизованному на сайте.  Если есть необходимость всегда показывать форму авторизации, приложение может добавить к запросу `/oauth/authorize...` параметр `force_login=true`. В этом случае, пользователю будет показана форма авторизации с логином и паролем даже в случае, если пользователь уже авторизован.  Это может быть полезно приложениям, которые предоставляют сервис только для соискателей. Если пришел пользователь-работодатель, приложение может предложить пользователю повторно разрешить доступ на сайте, уже указав другую учетную запись.  Также, после авторизации приложение может показать пользователю сообщение:  ``` Вы вошли как %Имя_Фамилия%. Это не вы? ``` и предоставить ссылку с `force_login=true` для возможности захода под другим логином.  ## Авторизация под разными рабочими аккаунтами  Для получения списка рабочих аккаунтов менеджера и для работы под разными рабочими аккаунтами менеджера необходимо прочитать документацию по [рабочим аккаунтам менеджера](#tag/Menedzhery-rabotodatelya/operation/get-manager-accounts)  ## Авторизация приложения  Токен приложения необходимо сгенерировать 1 раз. В случае, если токен был скомпрометирован, его нужно запросить еще раз. При этом ранее выданный токен отзывается. Владелец приложения может посмотреть актуальный `access_token` для приложения на сайте [https://dev.hh.ru/admin](https://dev.hh.ru/admin). В случае, если вы еще ни разу [не получали токен приложения](#section/Avtorizaciya/Avtorizaciya-prilozheniya), токен отображаться не будет.  <a name=\"get-client-token\"></a> ### Получение токена приложения Для получения `access_token` необходимо сделать запрос:  ``` POST https://hh.ru/oauth/token ```  В теле запроса необходимо передать [дополнительные параметры](#required_parameters). Тело запроса необходимо передавать в стандартном `application/x-www-form-urlencoded` с указанием соответствующего заголовка `Content-Type`.  Данный `access_token` имеет **неограниченный** срок жизни. При повторном запросе ранее выданный токен отзывается и выдается новый. Запрашивать `access_token` можно не чаще, чем один раз в 5 минут.  В случае компрометации токена необходимо [инвалидировать](#tag/Avtorizaciya-rabotodatelya/operation/invalidate-token) скомпроментированный токен и запросить токен заново!  <!-- ReDoc-Inject: <security-definitions> --> 

API version: 1.0.0
Contact: api@hh.ru
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package hh-go

import (
	"encoding/json"
	"bytes"
	"fmt"
)

// checks if the ErrorsCommonBadAuthorizationErrors type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ErrorsCommonBadAuthorizationErrors{}

// ErrorsCommonBadAuthorizationErrors Информация о возникших ошибках
type ErrorsCommonBadAuthorizationErrors struct {
	// Идентификатор запроса
	RequestId string `json:"request_id"`
	// Описание ошибки
	Description *string `json:"description,omitempty"`
	// Массив с данными ошибок
	Errors []ErrorsCommonBadAuthorizationCommonAndEmployerError `json:"errors"`
	// Ошибки авторизации:   * `token-revoked` — Токен отозван пользователем, приложению необходимо [запросить новую авторизацию](#tag/Avtorizaciya-rabotodatelya/operation/authorize)   * `token-expired` — Время жизни `access_token` завершилось, необходимо [получить `refresh_token`](#tag/Avtorizaciya-rabotodatelya/operation/authorize) 
	OauthError *string `json:"oauth_error,omitempty"`
}

type _ErrorsCommonBadAuthorizationErrors ErrorsCommonBadAuthorizationErrors

// NewErrorsCommonBadAuthorizationErrors instantiates a new ErrorsCommonBadAuthorizationErrors object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewErrorsCommonBadAuthorizationErrors(requestId string, errors []ErrorsCommonBadAuthorizationCommonAndEmployerError) *ErrorsCommonBadAuthorizationErrors {
	this := ErrorsCommonBadAuthorizationErrors{}
	this.RequestId = requestId
	this.Errors = errors
	return &this
}

// NewErrorsCommonBadAuthorizationErrorsWithDefaults instantiates a new ErrorsCommonBadAuthorizationErrors object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewErrorsCommonBadAuthorizationErrorsWithDefaults() *ErrorsCommonBadAuthorizationErrors {
	this := ErrorsCommonBadAuthorizationErrors{}
	return &this
}

// GetRequestId returns the RequestId field value
func (o *ErrorsCommonBadAuthorizationErrors) GetRequestId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.RequestId
}

// GetRequestIdOk returns a tuple with the RequestId field value
// and a boolean to check if the value has been set.
func (o *ErrorsCommonBadAuthorizationErrors) GetRequestIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.RequestId, true
}

// SetRequestId sets field value
func (o *ErrorsCommonBadAuthorizationErrors) SetRequestId(v string) {
	o.RequestId = v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *ErrorsCommonBadAuthorizationErrors) GetDescription() string {
	if o == nil || IsNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ErrorsCommonBadAuthorizationErrors) GetDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *ErrorsCommonBadAuthorizationErrors) HasDescription() bool {
	if o != nil && !IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *ErrorsCommonBadAuthorizationErrors) SetDescription(v string) {
	o.Description = &v
}

// GetErrors returns the Errors field value
func (o *ErrorsCommonBadAuthorizationErrors) GetErrors() []ErrorsCommonBadAuthorizationCommonAndEmployerError {
	if o == nil {
		var ret []ErrorsCommonBadAuthorizationCommonAndEmployerError
		return ret
	}

	return o.Errors
}

// GetErrorsOk returns a tuple with the Errors field value
// and a boolean to check if the value has been set.
func (o *ErrorsCommonBadAuthorizationErrors) GetErrorsOk() ([]ErrorsCommonBadAuthorizationCommonAndEmployerError, bool) {
	if o == nil {
		return nil, false
	}
	return o.Errors, true
}

// SetErrors sets field value
func (o *ErrorsCommonBadAuthorizationErrors) SetErrors(v []ErrorsCommonBadAuthorizationCommonAndEmployerError) {
	o.Errors = v
}

// GetOauthError returns the OauthError field value if set, zero value otherwise.
func (o *ErrorsCommonBadAuthorizationErrors) GetOauthError() string {
	if o == nil || IsNil(o.OauthError) {
		var ret string
		return ret
	}
	return *o.OauthError
}

// GetOauthErrorOk returns a tuple with the OauthError field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ErrorsCommonBadAuthorizationErrors) GetOauthErrorOk() (*string, bool) {
	if o == nil || IsNil(o.OauthError) {
		return nil, false
	}
	return o.OauthError, true
}

// HasOauthError returns a boolean if a field has been set.
func (o *ErrorsCommonBadAuthorizationErrors) HasOauthError() bool {
	if o != nil && !IsNil(o.OauthError) {
		return true
	}

	return false
}

// SetOauthError gets a reference to the given string and assigns it to the OauthError field.
func (o *ErrorsCommonBadAuthorizationErrors) SetOauthError(v string) {
	o.OauthError = &v
}

func (o ErrorsCommonBadAuthorizationErrors) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ErrorsCommonBadAuthorizationErrors) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["request_id"] = o.RequestId
	if !IsNil(o.Description) {
		toSerialize["description"] = o.Description
	}
	toSerialize["errors"] = o.Errors
	if !IsNil(o.OauthError) {
		toSerialize["oauth_error"] = o.OauthError
	}
	return toSerialize, nil
}

func (o *ErrorsCommonBadAuthorizationErrors) UnmarshalJSON(data []byte) (err error) {
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

	varErrorsCommonBadAuthorizationErrors := _ErrorsCommonBadAuthorizationErrors{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varErrorsCommonBadAuthorizationErrors)

	if err != nil {
		return err
	}

	*o = ErrorsCommonBadAuthorizationErrors(varErrorsCommonBadAuthorizationErrors)

	return err
}

type NullableErrorsCommonBadAuthorizationErrors struct {
	value *ErrorsCommonBadAuthorizationErrors
	isSet bool
}

func (v NullableErrorsCommonBadAuthorizationErrors) Get() *ErrorsCommonBadAuthorizationErrors {
	return v.value
}

func (v *NullableErrorsCommonBadAuthorizationErrors) Set(val *ErrorsCommonBadAuthorizationErrors) {
	v.value = val
	v.isSet = true
}

func (v NullableErrorsCommonBadAuthorizationErrors) IsSet() bool {
	return v.isSet
}

func (v *NullableErrorsCommonBadAuthorizationErrors) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableErrorsCommonBadAuthorizationErrors(val *ErrorsCommonBadAuthorizationErrors) *NullableErrorsCommonBadAuthorizationErrors {
	return &NullableErrorsCommonBadAuthorizationErrors{value: val, isSet: true}
}

func (v NullableErrorsCommonBadAuthorizationErrors) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableErrorsCommonBadAuthorizationErrors) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


