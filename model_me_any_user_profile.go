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

// checks if the MeAnyUserProfile type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &MeAnyUserProfile{}

// MeAnyUserProfile Базовый профиль текущего пользователя, авторизованного как соискатель или работодатель
type MeAnyUserProfile struct {
	MeAnyProfile
	// Email текущего пользователя
	Email NullableString `json:"email,omitempty"`
	// Имя текущего пользователя
	FirstName string `json:"first_name"`
	// Идентификатор текущего пользователя
	Id string `json:"id"`
	// Deprecated
	IsAnonymous *bool `json:"is_anonymous,omitempty"`
	// Фамилия текущего пользователя
	LastName string `json:"last_name"`
	// Deprecated
	MidName *string `json:"mid_name,omitempty"`
	// Отчество текущего пользователя
	MiddleName NullableString `json:"middle_name,omitempty"`
	// Телефон текущего пользователя
	Phone NullableString `json:"phone,omitempty"`
}

type _MeAnyUserProfile MeAnyUserProfile

// NewMeAnyUserProfile instantiates a new MeAnyUserProfile object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMeAnyUserProfile(firstName string, id string, lastName string, authType NullableString, isAdmin bool, isApplicant bool, isApplication bool, isEmployer bool, isEmployerIntegration bool) *MeAnyUserProfile {
	this := MeAnyUserProfile{}
	this.AuthType = authType
	this.IsAdmin = isAdmin
	this.IsApplicant = isApplicant
	this.IsApplication = isApplication
	this.IsEmployer = isEmployer
	this.IsEmployerIntegration = isEmployerIntegration
	this.FirstName = firstName
	this.Id = id
	this.LastName = lastName
	return &this
}

// NewMeAnyUserProfileWithDefaults instantiates a new MeAnyUserProfile object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMeAnyUserProfileWithDefaults() *MeAnyUserProfile {
	this := MeAnyUserProfile{}
	return &this
}

// GetEmail returns the Email field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *MeAnyUserProfile) GetEmail() string {
	if o == nil || IsNil(o.Email.Get()) {
		var ret string
		return ret
	}
	return *o.Email.Get()
}

// GetEmailOk returns a tuple with the Email field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *MeAnyUserProfile) GetEmailOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Email.Get(), o.Email.IsSet()
}

// HasEmail returns a boolean if a field has been set.
func (o *MeAnyUserProfile) HasEmail() bool {
	if o != nil && o.Email.IsSet() {
		return true
	}

	return false
}

// SetEmail gets a reference to the given NullableString and assigns it to the Email field.
func (o *MeAnyUserProfile) SetEmail(v string) {
	o.Email.Set(&v)
}
// SetEmailNil sets the value for Email to be an explicit nil
func (o *MeAnyUserProfile) SetEmailNil() {
	o.Email.Set(nil)
}

// UnsetEmail ensures that no value is present for Email, not even an explicit nil
func (o *MeAnyUserProfile) UnsetEmail() {
	o.Email.Unset()
}

// GetFirstName returns the FirstName field value
func (o *MeAnyUserProfile) GetFirstName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.FirstName
}

// GetFirstNameOk returns a tuple with the FirstName field value
// and a boolean to check if the value has been set.
func (o *MeAnyUserProfile) GetFirstNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.FirstName, true
}

// SetFirstName sets field value
func (o *MeAnyUserProfile) SetFirstName(v string) {
	o.FirstName = v
}

// GetId returns the Id field value
func (o *MeAnyUserProfile) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *MeAnyUserProfile) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *MeAnyUserProfile) SetId(v string) {
	o.Id = v
}

// GetIsAnonymous returns the IsAnonymous field value if set, zero value otherwise.
// Deprecated
func (o *MeAnyUserProfile) GetIsAnonymous() bool {
	if o == nil || IsNil(o.IsAnonymous) {
		var ret bool
		return ret
	}
	return *o.IsAnonymous
}

// GetIsAnonymousOk returns a tuple with the IsAnonymous field value if set, nil otherwise
// and a boolean to check if the value has been set.
// Deprecated
func (o *MeAnyUserProfile) GetIsAnonymousOk() (*bool, bool) {
	if o == nil || IsNil(o.IsAnonymous) {
		return nil, false
	}
	return o.IsAnonymous, true
}

// HasIsAnonymous returns a boolean if a field has been set.
func (o *MeAnyUserProfile) HasIsAnonymous() bool {
	if o != nil && !IsNil(o.IsAnonymous) {
		return true
	}

	return false
}

// SetIsAnonymous gets a reference to the given bool and assigns it to the IsAnonymous field.
// Deprecated
func (o *MeAnyUserProfile) SetIsAnonymous(v bool) {
	o.IsAnonymous = &v
}

// GetLastName returns the LastName field value
func (o *MeAnyUserProfile) GetLastName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.LastName
}

// GetLastNameOk returns a tuple with the LastName field value
// and a boolean to check if the value has been set.
func (o *MeAnyUserProfile) GetLastNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.LastName, true
}

// SetLastName sets field value
func (o *MeAnyUserProfile) SetLastName(v string) {
	o.LastName = v
}

// GetMidName returns the MidName field value if set, zero value otherwise.
// Deprecated
func (o *MeAnyUserProfile) GetMidName() string {
	if o == nil || IsNil(o.MidName) {
		var ret string
		return ret
	}
	return *o.MidName
}

// GetMidNameOk returns a tuple with the MidName field value if set, nil otherwise
// and a boolean to check if the value has been set.
// Deprecated
func (o *MeAnyUserProfile) GetMidNameOk() (*string, bool) {
	if o == nil || IsNil(o.MidName) {
		return nil, false
	}
	return o.MidName, true
}

// HasMidName returns a boolean if a field has been set.
func (o *MeAnyUserProfile) HasMidName() bool {
	if o != nil && !IsNil(o.MidName) {
		return true
	}

	return false
}

// SetMidName gets a reference to the given string and assigns it to the MidName field.
// Deprecated
func (o *MeAnyUserProfile) SetMidName(v string) {
	o.MidName = &v
}

// GetMiddleName returns the MiddleName field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *MeAnyUserProfile) GetMiddleName() string {
	if o == nil || IsNil(o.MiddleName.Get()) {
		var ret string
		return ret
	}
	return *o.MiddleName.Get()
}

// GetMiddleNameOk returns a tuple with the MiddleName field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *MeAnyUserProfile) GetMiddleNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.MiddleName.Get(), o.MiddleName.IsSet()
}

// HasMiddleName returns a boolean if a field has been set.
func (o *MeAnyUserProfile) HasMiddleName() bool {
	if o != nil && o.MiddleName.IsSet() {
		return true
	}

	return false
}

// SetMiddleName gets a reference to the given NullableString and assigns it to the MiddleName field.
func (o *MeAnyUserProfile) SetMiddleName(v string) {
	o.MiddleName.Set(&v)
}
// SetMiddleNameNil sets the value for MiddleName to be an explicit nil
func (o *MeAnyUserProfile) SetMiddleNameNil() {
	o.MiddleName.Set(nil)
}

// UnsetMiddleName ensures that no value is present for MiddleName, not even an explicit nil
func (o *MeAnyUserProfile) UnsetMiddleName() {
	o.MiddleName.Unset()
}

// GetPhone returns the Phone field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *MeAnyUserProfile) GetPhone() string {
	if o == nil || IsNil(o.Phone.Get()) {
		var ret string
		return ret
	}
	return *o.Phone.Get()
}

// GetPhoneOk returns a tuple with the Phone field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *MeAnyUserProfile) GetPhoneOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Phone.Get(), o.Phone.IsSet()
}

// HasPhone returns a boolean if a field has been set.
func (o *MeAnyUserProfile) HasPhone() bool {
	if o != nil && o.Phone.IsSet() {
		return true
	}

	return false
}

// SetPhone gets a reference to the given NullableString and assigns it to the Phone field.
func (o *MeAnyUserProfile) SetPhone(v string) {
	o.Phone.Set(&v)
}
// SetPhoneNil sets the value for Phone to be an explicit nil
func (o *MeAnyUserProfile) SetPhoneNil() {
	o.Phone.Set(nil)
}

// UnsetPhone ensures that no value is present for Phone, not even an explicit nil
func (o *MeAnyUserProfile) UnsetPhone() {
	o.Phone.Unset()
}

func (o MeAnyUserProfile) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o MeAnyUserProfile) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	serializedMeAnyProfile, errMeAnyProfile := json.Marshal(o.MeAnyProfile)
	if errMeAnyProfile != nil {
		return map[string]interface{}{}, errMeAnyProfile
	}
	errMeAnyProfile = json.Unmarshal([]byte(serializedMeAnyProfile), &toSerialize)
	if errMeAnyProfile != nil {
		return map[string]interface{}{}, errMeAnyProfile
	}
	if o.Email.IsSet() {
		toSerialize["email"] = o.Email.Get()
	}
	toSerialize["first_name"] = o.FirstName
	toSerialize["id"] = o.Id
	if !IsNil(o.IsAnonymous) {
		toSerialize["is_anonymous"] = o.IsAnonymous
	}
	toSerialize["last_name"] = o.LastName
	if !IsNil(o.MidName) {
		toSerialize["mid_name"] = o.MidName
	}
	if o.MiddleName.IsSet() {
		toSerialize["middle_name"] = o.MiddleName.Get()
	}
	if o.Phone.IsSet() {
		toSerialize["phone"] = o.Phone.Get()
	}
	return toSerialize, nil
}

func (o *MeAnyUserProfile) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"first_name",
		"id",
		"last_name",
		"auth_type",
		"is_admin",
		"is_applicant",
		"is_application",
		"is_employer",
		"is_employer_integration",
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

	varMeAnyUserProfile := _MeAnyUserProfile{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varMeAnyUserProfile)

	if err != nil {
		return err
	}

	*o = MeAnyUserProfile(varMeAnyUserProfile)

	return err
}

type NullableMeAnyUserProfile struct {
	value *MeAnyUserProfile
	isSet bool
}

func (v NullableMeAnyUserProfile) Get() *MeAnyUserProfile {
	return v.value
}

func (v *NullableMeAnyUserProfile) Set(val *MeAnyUserProfile) {
	v.value = val
	v.isSet = true
}

func (v NullableMeAnyUserProfile) IsSet() bool {
	return v.isSet
}

func (v *NullableMeAnyUserProfile) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMeAnyUserProfile(val *MeAnyUserProfile) *NullableMeAnyUserProfile {
	return &NullableMeAnyUserProfile{value: val, isSet: true}
}

func (v NullableMeAnyUserProfile) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMeAnyUserProfile) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


