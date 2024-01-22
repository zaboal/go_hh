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

// checks if the ResumeObjectsPhotoNoId type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ResumeObjectsPhotoNoId{}

// ResumeObjectsPhotoNoId struct for ResumeObjectsPhotoNoId
type ResumeObjectsPhotoNoId struct {
	// URL изображения размером 40x40 пикселей. Изображение по данному URL доступно ограниченное время после получения ответа. Приложение должно быть готово к тому, что на запрос изображения вернется ошибка `404 Not Found`
	// Deprecated
	Var40 NullableString `json:"40,omitempty"`
	// URL изображения размером 100x100 пикселей. Изображение по данному URL доступно ограниченное время после получения ответа. Приложение должно быть готово к тому, что на запрос изображения вернется ошибка `404 Not Found`
	// Deprecated
	Var100 NullableString `json:"100,omitempty"`
	// URL изображения размером 500x500 пикселей. Изображение по данному URL доступно ограниченное время после получения ответа. Приложение должно быть готово к тому, что на запрос изображения вернется ошибка `404 Not Found`
	// Deprecated
	Var500 NullableString `json:"500,omitempty"`
	// URL среднего по размеру изображения. Изображение по данному URL доступно ограниченное время после получения ответа. Приложение должно быть готово к тому, что на запрос изображения вернется ошибка `404 Not Found`
	Medium string `json:"medium"`
	// URL уменьшенного изображения. Изображение по данному URL доступно ограниченное время после получения ответа. Приложение должно быть готово к тому, что на запрос изображения вернется ошибка `404 Not Found`
	Small string `json:"small"`
}

type _ResumeObjectsPhotoNoId ResumeObjectsPhotoNoId

// NewResumeObjectsPhotoNoId instantiates a new ResumeObjectsPhotoNoId object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewResumeObjectsPhotoNoId(medium string, small string) *ResumeObjectsPhotoNoId {
	this := ResumeObjectsPhotoNoId{}
	this.Medium = medium
	this.Small = small
	return &this
}

// NewResumeObjectsPhotoNoIdWithDefaults instantiates a new ResumeObjectsPhotoNoId object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewResumeObjectsPhotoNoIdWithDefaults() *ResumeObjectsPhotoNoId {
	this := ResumeObjectsPhotoNoId{}
	return &this
}

// GetVar40 returns the Var40 field value if set, zero value otherwise (both if not set or set to explicit null).
// Deprecated
func (o *ResumeObjectsPhotoNoId) GetVar40() string {
	if o == nil || IsNil(o.Var40.Get()) {
		var ret string
		return ret
	}
	return *o.Var40.Get()
}

// GetVar40Ok returns a tuple with the Var40 field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
// Deprecated
func (o *ResumeObjectsPhotoNoId) GetVar40Ok() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Var40.Get(), o.Var40.IsSet()
}

// HasVar40 returns a boolean if a field has been set.
func (o *ResumeObjectsPhotoNoId) HasVar40() bool {
	if o != nil && o.Var40.IsSet() {
		return true
	}

	return false
}

// SetVar40 gets a reference to the given NullableString and assigns it to the Var40 field.
// Deprecated
func (o *ResumeObjectsPhotoNoId) SetVar40(v string) {
	o.Var40.Set(&v)
}
// SetVar40Nil sets the value for Var40 to be an explicit nil
func (o *ResumeObjectsPhotoNoId) SetVar40Nil() {
	o.Var40.Set(nil)
}

// UnsetVar40 ensures that no value is present for Var40, not even an explicit nil
func (o *ResumeObjectsPhotoNoId) UnsetVar40() {
	o.Var40.Unset()
}

// GetVar100 returns the Var100 field value if set, zero value otherwise (both if not set or set to explicit null).
// Deprecated
func (o *ResumeObjectsPhotoNoId) GetVar100() string {
	if o == nil || IsNil(o.Var100.Get()) {
		var ret string
		return ret
	}
	return *o.Var100.Get()
}

// GetVar100Ok returns a tuple with the Var100 field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
// Deprecated
func (o *ResumeObjectsPhotoNoId) GetVar100Ok() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Var100.Get(), o.Var100.IsSet()
}

// HasVar100 returns a boolean if a field has been set.
func (o *ResumeObjectsPhotoNoId) HasVar100() bool {
	if o != nil && o.Var100.IsSet() {
		return true
	}

	return false
}

// SetVar100 gets a reference to the given NullableString and assigns it to the Var100 field.
// Deprecated
func (o *ResumeObjectsPhotoNoId) SetVar100(v string) {
	o.Var100.Set(&v)
}
// SetVar100Nil sets the value for Var100 to be an explicit nil
func (o *ResumeObjectsPhotoNoId) SetVar100Nil() {
	o.Var100.Set(nil)
}

// UnsetVar100 ensures that no value is present for Var100, not even an explicit nil
func (o *ResumeObjectsPhotoNoId) UnsetVar100() {
	o.Var100.Unset()
}

// GetVar500 returns the Var500 field value if set, zero value otherwise (both if not set or set to explicit null).
// Deprecated
func (o *ResumeObjectsPhotoNoId) GetVar500() string {
	if o == nil || IsNil(o.Var500.Get()) {
		var ret string
		return ret
	}
	return *o.Var500.Get()
}

// GetVar500Ok returns a tuple with the Var500 field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
// Deprecated
func (o *ResumeObjectsPhotoNoId) GetVar500Ok() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Var500.Get(), o.Var500.IsSet()
}

// HasVar500 returns a boolean if a field has been set.
func (o *ResumeObjectsPhotoNoId) HasVar500() bool {
	if o != nil && o.Var500.IsSet() {
		return true
	}

	return false
}

// SetVar500 gets a reference to the given NullableString and assigns it to the Var500 field.
// Deprecated
func (o *ResumeObjectsPhotoNoId) SetVar500(v string) {
	o.Var500.Set(&v)
}
// SetVar500Nil sets the value for Var500 to be an explicit nil
func (o *ResumeObjectsPhotoNoId) SetVar500Nil() {
	o.Var500.Set(nil)
}

// UnsetVar500 ensures that no value is present for Var500, not even an explicit nil
func (o *ResumeObjectsPhotoNoId) UnsetVar500() {
	o.Var500.Unset()
}

// GetMedium returns the Medium field value
func (o *ResumeObjectsPhotoNoId) GetMedium() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Medium
}

// GetMediumOk returns a tuple with the Medium field value
// and a boolean to check if the value has been set.
func (o *ResumeObjectsPhotoNoId) GetMediumOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Medium, true
}

// SetMedium sets field value
func (o *ResumeObjectsPhotoNoId) SetMedium(v string) {
	o.Medium = v
}

// GetSmall returns the Small field value
func (o *ResumeObjectsPhotoNoId) GetSmall() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Small
}

// GetSmallOk returns a tuple with the Small field value
// and a boolean to check if the value has been set.
func (o *ResumeObjectsPhotoNoId) GetSmallOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Small, true
}

// SetSmall sets field value
func (o *ResumeObjectsPhotoNoId) SetSmall(v string) {
	o.Small = v
}

func (o ResumeObjectsPhotoNoId) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ResumeObjectsPhotoNoId) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.Var40.IsSet() {
		toSerialize["40"] = o.Var40.Get()
	}
	if o.Var100.IsSet() {
		toSerialize["100"] = o.Var100.Get()
	}
	if o.Var500.IsSet() {
		toSerialize["500"] = o.Var500.Get()
	}
	toSerialize["medium"] = o.Medium
	toSerialize["small"] = o.Small
	return toSerialize, nil
}

func (o *ResumeObjectsPhotoNoId) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"medium",
		"small",
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

	varResumeObjectsPhotoNoId := _ResumeObjectsPhotoNoId{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varResumeObjectsPhotoNoId)

	if err != nil {
		return err
	}

	*o = ResumeObjectsPhotoNoId(varResumeObjectsPhotoNoId)

	return err
}

type NullableResumeObjectsPhotoNoId struct {
	value *ResumeObjectsPhotoNoId
	isSet bool
}

func (v NullableResumeObjectsPhotoNoId) Get() *ResumeObjectsPhotoNoId {
	return v.value
}

func (v *NullableResumeObjectsPhotoNoId) Set(val *ResumeObjectsPhotoNoId) {
	v.value = val
	v.isSet = true
}

func (v NullableResumeObjectsPhotoNoId) IsSet() bool {
	return v.isSet
}

func (v *NullableResumeObjectsPhotoNoId) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableResumeObjectsPhotoNoId(val *ResumeObjectsPhotoNoId) *NullableResumeObjectsPhotoNoId {
	return &NullableResumeObjectsPhotoNoId{value: val, isSet: true}
}

func (v NullableResumeObjectsPhotoNoId) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableResumeObjectsPhotoNoId) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


