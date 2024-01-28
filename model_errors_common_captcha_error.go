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

// checks if the ErrorsCommonCaptchaError type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ErrorsCommonCaptchaError{}

// ErrorsCommonCaptchaError struct for ErrorsCommonCaptchaError
type ErrorsCommonCaptchaError struct {
	// Адрес веб-страницы, на которой можно пройти капчу.  После прохождения капчи аналогичный запрос в API должен выполниться успешно.  Приложение должно добавить в captcha_url обязательный параметр backurl,на который произойдет редирект   после прохождения капчи.  Backurl должен обязательно содержать схему, например, https:// или схему приложения 
	CaptchaUrl *string `json:"captcha_url,omitempty"`
	// Адрес веб-страницы, на котором можно капчу. Аналогично параметру captcha_url
	FallbackUrl *string `json:"fallback_url,omitempty"`
	// Текстовый идентификатор типа ошибки
	Type string `json:"type"`
	// Необходимо пройти капчу - `captcha_required` 
	Value string `json:"value"`
}

type _ErrorsCommonCaptchaError ErrorsCommonCaptchaError

// NewErrorsCommonCaptchaError instantiates a new ErrorsCommonCaptchaError object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewErrorsCommonCaptchaError(type_ string, value string) *ErrorsCommonCaptchaError {
	this := ErrorsCommonCaptchaError{}
	this.Type = type_
	this.Value = value
	return &this
}

// NewErrorsCommonCaptchaErrorWithDefaults instantiates a new ErrorsCommonCaptchaError object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewErrorsCommonCaptchaErrorWithDefaults() *ErrorsCommonCaptchaError {
	this := ErrorsCommonCaptchaError{}
	return &this
}

// GetCaptchaUrl returns the CaptchaUrl field value if set, zero value otherwise.
func (o *ErrorsCommonCaptchaError) GetCaptchaUrl() string {
	if o == nil || IsNil(o.CaptchaUrl) {
		var ret string
		return ret
	}
	return *o.CaptchaUrl
}

// GetCaptchaUrlOk returns a tuple with the CaptchaUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ErrorsCommonCaptchaError) GetCaptchaUrlOk() (*string, bool) {
	if o == nil || IsNil(o.CaptchaUrl) {
		return nil, false
	}
	return o.CaptchaUrl, true
}

// HasCaptchaUrl returns a boolean if a field has been set.
func (o *ErrorsCommonCaptchaError) HasCaptchaUrl() bool {
	if o != nil && !IsNil(o.CaptchaUrl) {
		return true
	}

	return false
}

// SetCaptchaUrl gets a reference to the given string and assigns it to the CaptchaUrl field.
func (o *ErrorsCommonCaptchaError) SetCaptchaUrl(v string) {
	o.CaptchaUrl = &v
}

// GetFallbackUrl returns the FallbackUrl field value if set, zero value otherwise.
func (o *ErrorsCommonCaptchaError) GetFallbackUrl() string {
	if o == nil || IsNil(o.FallbackUrl) {
		var ret string
		return ret
	}
	return *o.FallbackUrl
}

// GetFallbackUrlOk returns a tuple with the FallbackUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ErrorsCommonCaptchaError) GetFallbackUrlOk() (*string, bool) {
	if o == nil || IsNil(o.FallbackUrl) {
		return nil, false
	}
	return o.FallbackUrl, true
}

// HasFallbackUrl returns a boolean if a field has been set.
func (o *ErrorsCommonCaptchaError) HasFallbackUrl() bool {
	if o != nil && !IsNil(o.FallbackUrl) {
		return true
	}

	return false
}

// SetFallbackUrl gets a reference to the given string and assigns it to the FallbackUrl field.
func (o *ErrorsCommonCaptchaError) SetFallbackUrl(v string) {
	o.FallbackUrl = &v
}

// GetType returns the Type field value
func (o *ErrorsCommonCaptchaError) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *ErrorsCommonCaptchaError) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *ErrorsCommonCaptchaError) SetType(v string) {
	o.Type = v
}

// GetValue returns the Value field value
func (o *ErrorsCommonCaptchaError) GetValue() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Value
}

// GetValueOk returns a tuple with the Value field value
// and a boolean to check if the value has been set.
func (o *ErrorsCommonCaptchaError) GetValueOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Value, true
}

// SetValue sets field value
func (o *ErrorsCommonCaptchaError) SetValue(v string) {
	o.Value = v
}

func (o ErrorsCommonCaptchaError) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ErrorsCommonCaptchaError) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.CaptchaUrl) {
		toSerialize["captcha_url"] = o.CaptchaUrl
	}
	if !IsNil(o.FallbackUrl) {
		toSerialize["fallback_url"] = o.FallbackUrl
	}
	toSerialize["type"] = o.Type
	toSerialize["value"] = o.Value
	return toSerialize, nil
}

func (o *ErrorsCommonCaptchaError) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"type",
		"value",
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

	varErrorsCommonCaptchaError := _ErrorsCommonCaptchaError{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varErrorsCommonCaptchaError)

	if err != nil {
		return err
	}

	*o = ErrorsCommonCaptchaError(varErrorsCommonCaptchaError)

	return err
}

type NullableErrorsCommonCaptchaError struct {
	value *ErrorsCommonCaptchaError
	isSet bool
}

func (v NullableErrorsCommonCaptchaError) Get() *ErrorsCommonCaptchaError {
	return v.value
}

func (v *NullableErrorsCommonCaptchaError) Set(val *ErrorsCommonCaptchaError) {
	v.value = val
	v.isSet = true
}

func (v NullableErrorsCommonCaptchaError) IsSet() bool {
	return v.isSet
}

func (v *NullableErrorsCommonCaptchaError) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableErrorsCommonCaptchaError(val *ErrorsCommonCaptchaError) *NullableErrorsCommonCaptchaError {
	return &NullableErrorsCommonCaptchaError{value: val, isSet: true}
}

func (v NullableErrorsCommonCaptchaError) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableErrorsCommonCaptchaError) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


