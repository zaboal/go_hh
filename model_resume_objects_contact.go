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

// checks if the ResumeObjectsContact type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ResumeObjectsContact{}

// ResumeObjectsContact struct for ResumeObjectsContact
type ResumeObjectsContact struct {
	// Комментарий к контакту
	Comment NullableString `json:"comment,omitempty"`
	// Требуется ли подтверждение телефона
	NeedVerification NullableBool `json:"need_verification,omitempty"`
	// Является ли предпочтительным способом связи
	Preferred *bool `json:"preferred,omitempty"`
	// Тип контакта. Элемент справочника [preferred_contact_type](#tag/Obshie-spravochniki/operation/get-dictionaries)
	Type IncludesIdName `json:"type"`
	Value IncludesContactPropertiesValue `json:"value"`
	// Является ли телефон подтвержденным
	Verified NullableBool `json:"verified,omitempty"`
}

type _ResumeObjectsContact ResumeObjectsContact

// NewResumeObjectsContact instantiates a new ResumeObjectsContact object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewResumeObjectsContact(type_ IncludesIdName, value IncludesContactPropertiesValue) *ResumeObjectsContact {
	this := ResumeObjectsContact{}
	this.Type = type_
	this.Value = value
	return &this
}

// NewResumeObjectsContactWithDefaults instantiates a new ResumeObjectsContact object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewResumeObjectsContactWithDefaults() *ResumeObjectsContact {
	this := ResumeObjectsContact{}
	return &this
}

// GetComment returns the Comment field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ResumeObjectsContact) GetComment() string {
	if o == nil || IsNil(o.Comment.Get()) {
		var ret string
		return ret
	}
	return *o.Comment.Get()
}

// GetCommentOk returns a tuple with the Comment field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ResumeObjectsContact) GetCommentOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Comment.Get(), o.Comment.IsSet()
}

// HasComment returns a boolean if a field has been set.
func (o *ResumeObjectsContact) HasComment() bool {
	if o != nil && o.Comment.IsSet() {
		return true
	}

	return false
}

// SetComment gets a reference to the given NullableString and assigns it to the Comment field.
func (o *ResumeObjectsContact) SetComment(v string) {
	o.Comment.Set(&v)
}
// SetCommentNil sets the value for Comment to be an explicit nil
func (o *ResumeObjectsContact) SetCommentNil() {
	o.Comment.Set(nil)
}

// UnsetComment ensures that no value is present for Comment, not even an explicit nil
func (o *ResumeObjectsContact) UnsetComment() {
	o.Comment.Unset()
}

// GetNeedVerification returns the NeedVerification field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ResumeObjectsContact) GetNeedVerification() bool {
	if o == nil || IsNil(o.NeedVerification.Get()) {
		var ret bool
		return ret
	}
	return *o.NeedVerification.Get()
}

// GetNeedVerificationOk returns a tuple with the NeedVerification field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ResumeObjectsContact) GetNeedVerificationOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return o.NeedVerification.Get(), o.NeedVerification.IsSet()
}

// HasNeedVerification returns a boolean if a field has been set.
func (o *ResumeObjectsContact) HasNeedVerification() bool {
	if o != nil && o.NeedVerification.IsSet() {
		return true
	}

	return false
}

// SetNeedVerification gets a reference to the given NullableBool and assigns it to the NeedVerification field.
func (o *ResumeObjectsContact) SetNeedVerification(v bool) {
	o.NeedVerification.Set(&v)
}
// SetNeedVerificationNil sets the value for NeedVerification to be an explicit nil
func (o *ResumeObjectsContact) SetNeedVerificationNil() {
	o.NeedVerification.Set(nil)
}

// UnsetNeedVerification ensures that no value is present for NeedVerification, not even an explicit nil
func (o *ResumeObjectsContact) UnsetNeedVerification() {
	o.NeedVerification.Unset()
}

// GetPreferred returns the Preferred field value if set, zero value otherwise.
func (o *ResumeObjectsContact) GetPreferred() bool {
	if o == nil || IsNil(o.Preferred) {
		var ret bool
		return ret
	}
	return *o.Preferred
}

// GetPreferredOk returns a tuple with the Preferred field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ResumeObjectsContact) GetPreferredOk() (*bool, bool) {
	if o == nil || IsNil(o.Preferred) {
		return nil, false
	}
	return o.Preferred, true
}

// HasPreferred returns a boolean if a field has been set.
func (o *ResumeObjectsContact) HasPreferred() bool {
	if o != nil && !IsNil(o.Preferred) {
		return true
	}

	return false
}

// SetPreferred gets a reference to the given bool and assigns it to the Preferred field.
func (o *ResumeObjectsContact) SetPreferred(v bool) {
	o.Preferred = &v
}

// GetType returns the Type field value
func (o *ResumeObjectsContact) GetType() IncludesIdName {
	if o == nil {
		var ret IncludesIdName
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *ResumeObjectsContact) GetTypeOk() (*IncludesIdName, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *ResumeObjectsContact) SetType(v IncludesIdName) {
	o.Type = v
}

// GetValue returns the Value field value
func (o *ResumeObjectsContact) GetValue() IncludesContactPropertiesValue {
	if o == nil {
		var ret IncludesContactPropertiesValue
		return ret
	}

	return o.Value
}

// GetValueOk returns a tuple with the Value field value
// and a boolean to check if the value has been set.
func (o *ResumeObjectsContact) GetValueOk() (*IncludesContactPropertiesValue, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Value, true
}

// SetValue sets field value
func (o *ResumeObjectsContact) SetValue(v IncludesContactPropertiesValue) {
	o.Value = v
}

// GetVerified returns the Verified field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ResumeObjectsContact) GetVerified() bool {
	if o == nil || IsNil(o.Verified.Get()) {
		var ret bool
		return ret
	}
	return *o.Verified.Get()
}

// GetVerifiedOk returns a tuple with the Verified field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ResumeObjectsContact) GetVerifiedOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return o.Verified.Get(), o.Verified.IsSet()
}

// HasVerified returns a boolean if a field has been set.
func (o *ResumeObjectsContact) HasVerified() bool {
	if o != nil && o.Verified.IsSet() {
		return true
	}

	return false
}

// SetVerified gets a reference to the given NullableBool and assigns it to the Verified field.
func (o *ResumeObjectsContact) SetVerified(v bool) {
	o.Verified.Set(&v)
}
// SetVerifiedNil sets the value for Verified to be an explicit nil
func (o *ResumeObjectsContact) SetVerifiedNil() {
	o.Verified.Set(nil)
}

// UnsetVerified ensures that no value is present for Verified, not even an explicit nil
func (o *ResumeObjectsContact) UnsetVerified() {
	o.Verified.Unset()
}

func (o ResumeObjectsContact) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ResumeObjectsContact) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.Comment.IsSet() {
		toSerialize["comment"] = o.Comment.Get()
	}
	if o.NeedVerification.IsSet() {
		toSerialize["need_verification"] = o.NeedVerification.Get()
	}
	if !IsNil(o.Preferred) {
		toSerialize["preferred"] = o.Preferred
	}
	toSerialize["type"] = o.Type
	toSerialize["value"] = o.Value
	if o.Verified.IsSet() {
		toSerialize["verified"] = o.Verified.Get()
	}
	return toSerialize, nil
}

func (o *ResumeObjectsContact) UnmarshalJSON(data []byte) (err error) {
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

	varResumeObjectsContact := _ResumeObjectsContact{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varResumeObjectsContact)

	if err != nil {
		return err
	}

	*o = ResumeObjectsContact(varResumeObjectsContact)

	return err
}

type NullableResumeObjectsContact struct {
	value *ResumeObjectsContact
	isSet bool
}

func (v NullableResumeObjectsContact) Get() *ResumeObjectsContact {
	return v.value
}

func (v *NullableResumeObjectsContact) Set(val *ResumeObjectsContact) {
	v.value = val
	v.isSet = true
}

func (v NullableResumeObjectsContact) IsSet() bool {
	return v.isSet
}

func (v *NullableResumeObjectsContact) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableResumeObjectsContact(val *ResumeObjectsContact) *NullableResumeObjectsContact {
	return &NullableResumeObjectsContact{value: val, isSet: true}
}

func (v NullableResumeObjectsContact) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableResumeObjectsContact) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


