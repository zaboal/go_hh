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

// checks if the MeManagerProfile type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &MeManagerProfile{}

// MeManagerProfile Профиль текущего пользователя, авторизованного как менеджер
type MeManagerProfile struct {
	MeAnyUserProfile
	Employer *MeEmployerProfileCompanyDeprecated `json:"employer,omitempty"`
	// Deprecated
	IsInSearch map[string]interface{} `json:"is_in_search,omitempty"`
	Manager *MeEmployerProfileManager `json:"manager,omitempty"`
	// Deprecated
	NegotiationsUrl map[string]interface{} `json:"negotiations_url,omitempty"`
	PersonalManager *MeEmployerProfilePersonalManager `json:"personal_manager,omitempty"`
	// Deprecated
	ResumesUrl map[string]interface{} `json:"resumes_url,omitempty"`
}

type _MeManagerProfile MeManagerProfile

// NewMeManagerProfile instantiates a new MeManagerProfile object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMeManagerProfile(authType NullableString, isAdmin bool, isApplicant bool, isApplication bool, isEmployer bool, isEmployerIntegration bool, firstName string, id string, lastName string) *MeManagerProfile {
	this := MeManagerProfile{}
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

// NewMeManagerProfileWithDefaults instantiates a new MeManagerProfile object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMeManagerProfileWithDefaults() *MeManagerProfile {
	this := MeManagerProfile{}
	return &this
}

// GetEmployer returns the Employer field value if set, zero value otherwise.
func (o *MeManagerProfile) GetEmployer() MeEmployerProfileCompanyDeprecated {
	if o == nil || IsNil(o.Employer) {
		var ret MeEmployerProfileCompanyDeprecated
		return ret
	}
	return *o.Employer
}

// GetEmployerOk returns a tuple with the Employer field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MeManagerProfile) GetEmployerOk() (*MeEmployerProfileCompanyDeprecated, bool) {
	if o == nil || IsNil(o.Employer) {
		return nil, false
	}
	return o.Employer, true
}

// HasEmployer returns a boolean if a field has been set.
func (o *MeManagerProfile) HasEmployer() bool {
	if o != nil && !IsNil(o.Employer) {
		return true
	}

	return false
}

// SetEmployer gets a reference to the given MeEmployerProfileCompanyDeprecated and assigns it to the Employer field.
func (o *MeManagerProfile) SetEmployer(v MeEmployerProfileCompanyDeprecated) {
	o.Employer = &v
}

// GetIsInSearch returns the IsInSearch field value if set, zero value otherwise (both if not set or set to explicit null).
// Deprecated
func (o *MeManagerProfile) GetIsInSearch() map[string]interface{} {
	if o == nil {
		var ret map[string]interface{}
		return ret
	}
	return o.IsInSearch
}

// GetIsInSearchOk returns a tuple with the IsInSearch field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
// Deprecated
func (o *MeManagerProfile) GetIsInSearchOk() (map[string]interface{}, bool) {
	if o == nil || IsNil(o.IsInSearch) {
		return map[string]interface{}{}, false
	}
	return o.IsInSearch, true
}

// HasIsInSearch returns a boolean if a field has been set.
func (o *MeManagerProfile) HasIsInSearch() bool {
	if o != nil && IsNil(o.IsInSearch) {
		return true
	}

	return false
}

// SetIsInSearch gets a reference to the given map[string]interface{} and assigns it to the IsInSearch field.
// Deprecated
func (o *MeManagerProfile) SetIsInSearch(v map[string]interface{}) {
	o.IsInSearch = v
}

// GetManager returns the Manager field value if set, zero value otherwise.
func (o *MeManagerProfile) GetManager() MeEmployerProfileManager {
	if o == nil || IsNil(o.Manager) {
		var ret MeEmployerProfileManager
		return ret
	}
	return *o.Manager
}

// GetManagerOk returns a tuple with the Manager field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MeManagerProfile) GetManagerOk() (*MeEmployerProfileManager, bool) {
	if o == nil || IsNil(o.Manager) {
		return nil, false
	}
	return o.Manager, true
}

// HasManager returns a boolean if a field has been set.
func (o *MeManagerProfile) HasManager() bool {
	if o != nil && !IsNil(o.Manager) {
		return true
	}

	return false
}

// SetManager gets a reference to the given MeEmployerProfileManager and assigns it to the Manager field.
func (o *MeManagerProfile) SetManager(v MeEmployerProfileManager) {
	o.Manager = &v
}

// GetNegotiationsUrl returns the NegotiationsUrl field value if set, zero value otherwise (both if not set or set to explicit null).
// Deprecated
func (o *MeManagerProfile) GetNegotiationsUrl() map[string]interface{} {
	if o == nil {
		var ret map[string]interface{}
		return ret
	}
	return o.NegotiationsUrl
}

// GetNegotiationsUrlOk returns a tuple with the NegotiationsUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
// Deprecated
func (o *MeManagerProfile) GetNegotiationsUrlOk() (map[string]interface{}, bool) {
	if o == nil || IsNil(o.NegotiationsUrl) {
		return map[string]interface{}{}, false
	}
	return o.NegotiationsUrl, true
}

// HasNegotiationsUrl returns a boolean if a field has been set.
func (o *MeManagerProfile) HasNegotiationsUrl() bool {
	if o != nil && IsNil(o.NegotiationsUrl) {
		return true
	}

	return false
}

// SetNegotiationsUrl gets a reference to the given map[string]interface{} and assigns it to the NegotiationsUrl field.
// Deprecated
func (o *MeManagerProfile) SetNegotiationsUrl(v map[string]interface{}) {
	o.NegotiationsUrl = v
}

// GetPersonalManager returns the PersonalManager field value if set, zero value otherwise.
func (o *MeManagerProfile) GetPersonalManager() MeEmployerProfilePersonalManager {
	if o == nil || IsNil(o.PersonalManager) {
		var ret MeEmployerProfilePersonalManager
		return ret
	}
	return *o.PersonalManager
}

// GetPersonalManagerOk returns a tuple with the PersonalManager field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MeManagerProfile) GetPersonalManagerOk() (*MeEmployerProfilePersonalManager, bool) {
	if o == nil || IsNil(o.PersonalManager) {
		return nil, false
	}
	return o.PersonalManager, true
}

// HasPersonalManager returns a boolean if a field has been set.
func (o *MeManagerProfile) HasPersonalManager() bool {
	if o != nil && !IsNil(o.PersonalManager) {
		return true
	}

	return false
}

// SetPersonalManager gets a reference to the given MeEmployerProfilePersonalManager and assigns it to the PersonalManager field.
func (o *MeManagerProfile) SetPersonalManager(v MeEmployerProfilePersonalManager) {
	o.PersonalManager = &v
}

// GetResumesUrl returns the ResumesUrl field value if set, zero value otherwise (both if not set or set to explicit null).
// Deprecated
func (o *MeManagerProfile) GetResumesUrl() map[string]interface{} {
	if o == nil {
		var ret map[string]interface{}
		return ret
	}
	return o.ResumesUrl
}

// GetResumesUrlOk returns a tuple with the ResumesUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
// Deprecated
func (o *MeManagerProfile) GetResumesUrlOk() (map[string]interface{}, bool) {
	if o == nil || IsNil(o.ResumesUrl) {
		return map[string]interface{}{}, false
	}
	return o.ResumesUrl, true
}

// HasResumesUrl returns a boolean if a field has been set.
func (o *MeManagerProfile) HasResumesUrl() bool {
	if o != nil && IsNil(o.ResumesUrl) {
		return true
	}

	return false
}

// SetResumesUrl gets a reference to the given map[string]interface{} and assigns it to the ResumesUrl field.
// Deprecated
func (o *MeManagerProfile) SetResumesUrl(v map[string]interface{}) {
	o.ResumesUrl = v
}

func (o MeManagerProfile) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o MeManagerProfile) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	serializedMeAnyUserProfile, errMeAnyUserProfile := json.Marshal(o.MeAnyUserProfile)
	if errMeAnyUserProfile != nil {
		return map[string]interface{}{}, errMeAnyUserProfile
	}
	errMeAnyUserProfile = json.Unmarshal([]byte(serializedMeAnyUserProfile), &toSerialize)
	if errMeAnyUserProfile != nil {
		return map[string]interface{}{}, errMeAnyUserProfile
	}
	if !IsNil(o.Employer) {
		toSerialize["employer"] = o.Employer
	}
	if o.IsInSearch != nil {
		toSerialize["is_in_search"] = o.IsInSearch
	}
	if !IsNil(o.Manager) {
		toSerialize["manager"] = o.Manager
	}
	if o.NegotiationsUrl != nil {
		toSerialize["negotiations_url"] = o.NegotiationsUrl
	}
	if !IsNil(o.PersonalManager) {
		toSerialize["personal_manager"] = o.PersonalManager
	}
	if o.ResumesUrl != nil {
		toSerialize["resumes_url"] = o.ResumesUrl
	}
	return toSerialize, nil
}

func (o *MeManagerProfile) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"auth_type",
		"is_admin",
		"is_applicant",
		"is_application",
		"is_employer",
		"is_employer_integration",
		"first_name",
		"id",
		"last_name",
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

	varMeManagerProfile := _MeManagerProfile{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varMeManagerProfile)

	if err != nil {
		return err
	}

	*o = MeManagerProfile(varMeManagerProfile)

	return err
}

type NullableMeManagerProfile struct {
	value *MeManagerProfile
	isSet bool
}

func (v NullableMeManagerProfile) Get() *MeManagerProfile {
	return v.value
}

func (v *NullableMeManagerProfile) Set(val *MeManagerProfile) {
	v.value = val
	v.isSet = true
}

func (v NullableMeManagerProfile) IsSet() bool {
	return v.isSet
}

func (v *NullableMeManagerProfile) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMeManagerProfile(val *MeManagerProfile) *NullableMeManagerProfile {
	return &NullableMeManagerProfile{value: val, isSet: true}
}

func (v NullableMeManagerProfile) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMeManagerProfile) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


