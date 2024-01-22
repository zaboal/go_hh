/*
HeadHunter API

По-русски | [Switch to English](https://api.hh.ru/openapi/en/redoc)  В OpenAPI ведется пока что только небольшая часть документации [Основная документация](https://github.com/hhru/api).  Для поиска по документации можно использовать Ctrl+F.  # Авторизация  API поддерживает следующие уровни авторизации:   - [авторизация приложения](#section/Avtorizaciya/Avtorizaciya-prilozheniya)   - [авторизация пользователя](#section/Avtorizaciya/Avtorizaciya-polzovatelya)  * [Авторизация пользователя](#section/Avtorizaciya/Avtorizaciya-polzovatelya)   * [Правила формирования специального redirect_uri](#section/Avtorizaciya/Pravila-formirovaniya-specialnogo-redirect_uri)   * [Процесс авторизации](#section/Avtorizaciya/Process-avtorizacii)   * [Успешное получение временного `authorization_code`](#get-authorization_code)   * [Получение access и refresh токенов](#section/Avtorizaciya/Poluchenie-access-i-refresh-tokenov) * [Обновление пары access и refresh токенов](#section/Avtorizaciya/Obnovlenie-pary-access-i-refresh-tokenov) * [Инвалидация токена](#tag/Avtorizaciya-rabotodatelya/operation/invalidate-token) * [Запрос авторизации под другим пользователем](#section/Avtorizaciya/Zapros-avtorizacii-pod-drugim-polzovatelem) * [Авторизация под разными рабочими аккаунтами](#section/Avtorizaciya/Avtorizaciya-pod-raznymi-rabochimi-akkauntami) * [Авторизация приложения](#section/Avtorizaciya/Avtorizaciya-prilozheniya)       ## Авторизация пользователя Для выполнения запросов от имени пользователя необходимо пользоваться токеном пользователя.  В начале приложению необходимо направить пользователя (открыть страницу) по адресу:  ``` https://hh.ru/oauth/authorize? response_type=code& client_id={client_id}& state={state}& redirect_uri={redirect_uri} ```  Обязательные параметры:  * `response_type=code` — указание на способ получения авторизации, используя `authorization code` * `client_id` — идентификатор, полученный при создании приложения   Необязательные параметры:  * `state` — в случае указания, будет включен в ответный редирект. Это позволяет исключить возможность взлома путём подделки межсайтовых запросов. Подробнее об этом: [RFC 6749. Section 10.12](http://tools.ietf.org/html/rfc6749#section-10.12) * `redirect_uri` — uri для перенаправления пользователя после авторизации. Если не указать, используется из настроек приложения. При наличии происходит валидация значения. Вероятнее всего, потребуется сделать urlencode значения параметра.  ## Правила формирования специального redirect_uri  К примеру, если в настройках сохранен `http://example.com/oauth`, то разрешено указывать:  * `http://www.example.com/oauth` — поддомен; * `http://www.example.com/oauth/sub/path` — уточнение пути; * `http://example.com/oauth?lang=RU` — дополнительный параметр; * `http://www.example.com/oauth/sub/path?lang=RU` — всё вместе.  Запрещено:  * `https://example.com/oauth` — различные протоколы; * `http://wwwexample.com/oauth` — различные домены; * `http://wwwexample.com/` — другой путь; * `http://example.com/oauths` — другой путь; * `http://example.com:80/oauths` — указание изначально отсутствующего порта;  ## Процесс авторизации  Если пользователь не авторизован на сайте, ему будет показана форма авторизации на сайте. После прохождения авторизации на сайте, пользователю будет выведена форма с запросом разрешения доступа вашего приложения к его персональным данным.  Если пользователь не разрешает доступ приложению, пользователь будет перенаправлен на указанный `redirect_uri` с `?error=access_denied` и `state={state}`, если таковой был указан при первом запросе.  <a name=\"get-authorization_code\"></a> ### Успешное получение временного `authorization_code`  В случае разрешения прав, в редиректе будет указан временный `authorization_code`:  ```http HTTP/1.1 302 FOUND Location: {redirect_uri}?code={authorization_code} ```  Если пользователь авторизован на сайте и доступ данному приложению однажды ранее выдан, ответом будет сразу вышеописанный редирект с `authorization_code` (без показа формы логина и выдачи прав).  ## Получение access и refresh токенов  После получения `authorization_code` приложению необходимо сделать сервер-сервер запрос `POST https://hh.ru/oauth/token` для обмена полученного `authorization_code` на `access_token`.  В теле запроса необходимо передать [дополнительные параметры](#required_parameters).  Тело запроса необходимо передавать в стандартном `application/x-www-form-urlencoded` с указанием соответствующего заголовка `Content-Type`.  `authorization_code` имеет довольно короткий срок жизни, при его истечении необходимо запросить новый.  ## Обновление пары access и refresh токенов `access_token` также имеет срок жизни (ключ `expires_in`, в секундах), при его истечении приложение должно сделать запрос с `refresh_token` для получения нового.  Запрос необходимо делать в `application/x-www-form-urlencoded`.  ``` POST https://hh.ru/oauth/token ```  В теле запроса необходимо передать [дополнительные параметры](#required_parameters)  `refresh_token` можно использовать только один раз и только по истечению срока действия `access_token`.  После получения новой пары access и refresh токенов, их необходимо использовать в дальнейших запросах в api и запросах на продление токена.  ## Запрос авторизации под другим пользователем  Возможен следующий сценарий:  1. Приложение перенаправляет пользователя на сайт с запросом авторизации. 2. Пользователь на сайте уже авторизован и данному приложение доступ уже был разрешен. 3. Пользователю будет предложена возможность продолжить работу под текущим аккаунтом, либо зайти под другим аккаунтом.  Если есть необходимость, чтобы на шаге 3 сразу происходило перенаправление (redirect) с временным токеном, необходимо добавить к запросу `/oauth/authorize...` параметр `skip_choose_account=true`. В этом случае автоматически выдаётся доступ пользователю авторизованному на сайте.  Если есть необходимость всегда показывать форму авторизации, приложение может добавить к запросу `/oauth/authorize...` параметр `force_login=true`. В этом случае, пользователю будет показана форма авторизации с логином и паролем даже в случае, если пользователь уже авторизован.  Это может быть полезно приложениям, которые предоставляют сервис только для соискателей. Если пришел пользователь-работодатель, приложение может предложить пользователю повторно разрешить доступ на сайте, уже указав другую учетную запись.  Также, после авторизации приложение может показать пользователю сообщение:  ``` Вы вошли как %Имя_Фамилия%. Это не вы? ``` и предоставить ссылку с `force_login=true` для возможности захода под другим логином.  ## Авторизация под разными рабочими аккаунтами  Для получения списка рабочих аккаунтов менеджера и для работы под разными рабочими аккаунтами менеджера необходимо прочитать документацию по [рабочим аккаунтам менеджера](#tag/Menedzhery-rabotodatelya/operation/get-manager-accounts)  ## Авторизация приложения  Токен приложения необходимо сгенерировать 1 раз. В случае, если токен был скомпрометирован, его нужно запросить еще раз. При этом ранее выданный токен отзывается. Владелец приложения может посмотреть актуальный `access_token` для приложения на сайте [https://dev.hh.ru/admin](https://dev.hh.ru/admin). В случае, если вы еще ни разу [не получали токен приложения](#section/Avtorizaciya/Avtorizaciya-prilozheniya), токен отображаться не будет.  <a name=\"get-client-token\"></a> ### Получение токена приложения Для получения `access_token` необходимо сделать запрос:  ``` POST https://hh.ru/oauth/token ```  В теле запроса необходимо передать [дополнительные параметры](#required_parameters). Тело запроса необходимо передавать в стандартном `application/x-www-form-urlencoded` с указанием соответствующего заголовка `Content-Type`.  Данный `access_token` имеет **неограниченный** срок жизни. При повторном запросе ранее выданный токен отзывается и выдается новый. Запрашивать `access_token` можно не чаще, чем один раз в 5 минут.  В случае компрометации токена необходимо [инвалидировать](#tag/Avtorizaciya-rabotodatelya/operation/invalidate-token) скомпроментированный токен и запросить токен заново!  <!-- ReDoc-Inject: <security-definitions> --> 

API version: 1.0.0
Contact: api@hh.ru
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package github.com/zaboal/hh-go

import (
	"encoding/json"
	"bytes"
	"fmt"
)

// checks if the ResumeShouldSendSmsPhone type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ResumeShouldSendSmsPhone{}

// ResumeShouldSendSmsPhone Информация о телефоне соискателя
type ResumeShouldSendSmsPhone struct {
	// Префикс города номера телефона
	City string `json:"city"`
	// Префикс страны номера телефона
	Country string `json:"country"`
	// Номер телефона отформатированный
	Formatted string `json:"formatted"`
	// Нужно ли верифицировать номер телефона
	NeedVerification bool `json:"need_verification"`
	// Номер телефона без префиксов
	Number string `json:"number"`
	// Находится ли номер телефона в черном списке по отправке смс
	RestrictedCountry bool `json:"restricted_country"`
	// Верифицирован ли номер телефона пользователем
	Verified bool `json:"verified"`
}

type _ResumeShouldSendSmsPhone ResumeShouldSendSmsPhone

// NewResumeShouldSendSmsPhone instantiates a new ResumeShouldSendSmsPhone object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewResumeShouldSendSmsPhone(city string, country string, formatted string, needVerification bool, number string, restrictedCountry bool, verified bool) *ResumeShouldSendSmsPhone {
	this := ResumeShouldSendSmsPhone{}
	this.City = city
	this.Country = country
	this.Formatted = formatted
	this.NeedVerification = needVerification
	this.Number = number
	this.RestrictedCountry = restrictedCountry
	this.Verified = verified
	return &this
}

// NewResumeShouldSendSmsPhoneWithDefaults instantiates a new ResumeShouldSendSmsPhone object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewResumeShouldSendSmsPhoneWithDefaults() *ResumeShouldSendSmsPhone {
	this := ResumeShouldSendSmsPhone{}
	return &this
}

// GetCity returns the City field value
func (o *ResumeShouldSendSmsPhone) GetCity() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.City
}

// GetCityOk returns a tuple with the City field value
// and a boolean to check if the value has been set.
func (o *ResumeShouldSendSmsPhone) GetCityOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.City, true
}

// SetCity sets field value
func (o *ResumeShouldSendSmsPhone) SetCity(v string) {
	o.City = v
}

// GetCountry returns the Country field value
func (o *ResumeShouldSendSmsPhone) GetCountry() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Country
}

// GetCountryOk returns a tuple with the Country field value
// and a boolean to check if the value has been set.
func (o *ResumeShouldSendSmsPhone) GetCountryOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Country, true
}

// SetCountry sets field value
func (o *ResumeShouldSendSmsPhone) SetCountry(v string) {
	o.Country = v
}

// GetFormatted returns the Formatted field value
func (o *ResumeShouldSendSmsPhone) GetFormatted() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Formatted
}

// GetFormattedOk returns a tuple with the Formatted field value
// and a boolean to check if the value has been set.
func (o *ResumeShouldSendSmsPhone) GetFormattedOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Formatted, true
}

// SetFormatted sets field value
func (o *ResumeShouldSendSmsPhone) SetFormatted(v string) {
	o.Formatted = v
}

// GetNeedVerification returns the NeedVerification field value
func (o *ResumeShouldSendSmsPhone) GetNeedVerification() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.NeedVerification
}

// GetNeedVerificationOk returns a tuple with the NeedVerification field value
// and a boolean to check if the value has been set.
func (o *ResumeShouldSendSmsPhone) GetNeedVerificationOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.NeedVerification, true
}

// SetNeedVerification sets field value
func (o *ResumeShouldSendSmsPhone) SetNeedVerification(v bool) {
	o.NeedVerification = v
}

// GetNumber returns the Number field value
func (o *ResumeShouldSendSmsPhone) GetNumber() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Number
}

// GetNumberOk returns a tuple with the Number field value
// and a boolean to check if the value has been set.
func (o *ResumeShouldSendSmsPhone) GetNumberOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Number, true
}

// SetNumber sets field value
func (o *ResumeShouldSendSmsPhone) SetNumber(v string) {
	o.Number = v
}

// GetRestrictedCountry returns the RestrictedCountry field value
func (o *ResumeShouldSendSmsPhone) GetRestrictedCountry() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.RestrictedCountry
}

// GetRestrictedCountryOk returns a tuple with the RestrictedCountry field value
// and a boolean to check if the value has been set.
func (o *ResumeShouldSendSmsPhone) GetRestrictedCountryOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.RestrictedCountry, true
}

// SetRestrictedCountry sets field value
func (o *ResumeShouldSendSmsPhone) SetRestrictedCountry(v bool) {
	o.RestrictedCountry = v
}

// GetVerified returns the Verified field value
func (o *ResumeShouldSendSmsPhone) GetVerified() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.Verified
}

// GetVerifiedOk returns a tuple with the Verified field value
// and a boolean to check if the value has been set.
func (o *ResumeShouldSendSmsPhone) GetVerifiedOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Verified, true
}

// SetVerified sets field value
func (o *ResumeShouldSendSmsPhone) SetVerified(v bool) {
	o.Verified = v
}

func (o ResumeShouldSendSmsPhone) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ResumeShouldSendSmsPhone) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["city"] = o.City
	toSerialize["country"] = o.Country
	toSerialize["formatted"] = o.Formatted
	toSerialize["need_verification"] = o.NeedVerification
	toSerialize["number"] = o.Number
	toSerialize["restricted_country"] = o.RestrictedCountry
	toSerialize["verified"] = o.Verified
	return toSerialize, nil
}

func (o *ResumeShouldSendSmsPhone) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"city",
		"country",
		"formatted",
		"need_verification",
		"number",
		"restricted_country",
		"verified",
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

	varResumeShouldSendSmsPhone := _ResumeShouldSendSmsPhone{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varResumeShouldSendSmsPhone)

	if err != nil {
		return err
	}

	*o = ResumeShouldSendSmsPhone(varResumeShouldSendSmsPhone)

	return err
}

type NullableResumeShouldSendSmsPhone struct {
	value *ResumeShouldSendSmsPhone
	isSet bool
}

func (v NullableResumeShouldSendSmsPhone) Get() *ResumeShouldSendSmsPhone {
	return v.value
}

func (v *NullableResumeShouldSendSmsPhone) Set(val *ResumeShouldSendSmsPhone) {
	v.value = val
	v.isSet = true
}

func (v NullableResumeShouldSendSmsPhone) IsSet() bool {
	return v.isSet
}

func (v *NullableResumeShouldSendSmsPhone) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableResumeShouldSendSmsPhone(val *ResumeShouldSendSmsPhone) *NullableResumeShouldSendSmsPhone {
	return &NullableResumeShouldSendSmsPhone{value: val, isSet: true}
}

func (v NullableResumeShouldSendSmsPhone) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableResumeShouldSendSmsPhone) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


