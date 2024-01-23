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

// checks if the ResumesByStatusResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ResumesByStatusResponse{}

// ResumesByStatusResponse struct for ResumesByStatusResponse
type ResumesByStatusResponse struct {
	// Список резюме, уже использованных для отклика на данную вакансию
	AlreadyApplied []ResumesSuitableResumeItem `json:"already_applied"`
	Counters NullableResumesByStatusCounters `json:"counters,omitempty"`
	// Список неопубликованных резюме (в [статусе](#tag/Rezyume.-Prosmotr-informacii/Status-rezyume) `not_published` или `blocked`)
	NotPublished []ResumesSuitableResumeItem `json:"not_published"`
	// Список резюме, которыми можно откликнуться на данную вакансию
	Suitable []ResumesSuitableResumeItem `json:"suitable"`
	// Список резюме, которыми невозможно откликнуться на данную вакансию (например, из-за конфликтующих настроек видимости)
	Unavailable []ResumesSuitableResumeItem `json:"unavailable"`
}

type _ResumesByStatusResponse ResumesByStatusResponse

// NewResumesByStatusResponse instantiates a new ResumesByStatusResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewResumesByStatusResponse(alreadyApplied []ResumesSuitableResumeItem, notPublished []ResumesSuitableResumeItem, suitable []ResumesSuitableResumeItem, unavailable []ResumesSuitableResumeItem) *ResumesByStatusResponse {
	this := ResumesByStatusResponse{}
	this.AlreadyApplied = alreadyApplied
	this.NotPublished = notPublished
	this.Suitable = suitable
	this.Unavailable = unavailable
	return &this
}

// NewResumesByStatusResponseWithDefaults instantiates a new ResumesByStatusResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewResumesByStatusResponseWithDefaults() *ResumesByStatusResponse {
	this := ResumesByStatusResponse{}
	return &this
}

// GetAlreadyApplied returns the AlreadyApplied field value
func (o *ResumesByStatusResponse) GetAlreadyApplied() []ResumesSuitableResumeItem {
	if o == nil {
		var ret []ResumesSuitableResumeItem
		return ret
	}

	return o.AlreadyApplied
}

// GetAlreadyAppliedOk returns a tuple with the AlreadyApplied field value
// and a boolean to check if the value has been set.
func (o *ResumesByStatusResponse) GetAlreadyAppliedOk() ([]ResumesSuitableResumeItem, bool) {
	if o == nil {
		return nil, false
	}
	return o.AlreadyApplied, true
}

// SetAlreadyApplied sets field value
func (o *ResumesByStatusResponse) SetAlreadyApplied(v []ResumesSuitableResumeItem) {
	o.AlreadyApplied = v
}

// GetCounters returns the Counters field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ResumesByStatusResponse) GetCounters() ResumesByStatusCounters {
	if o == nil || IsNil(o.Counters.Get()) {
		var ret ResumesByStatusCounters
		return ret
	}
	return *o.Counters.Get()
}

// GetCountersOk returns a tuple with the Counters field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ResumesByStatusResponse) GetCountersOk() (*ResumesByStatusCounters, bool) {
	if o == nil {
		return nil, false
	}
	return o.Counters.Get(), o.Counters.IsSet()
}

// HasCounters returns a boolean if a field has been set.
func (o *ResumesByStatusResponse) HasCounters() bool {
	if o != nil && o.Counters.IsSet() {
		return true
	}

	return false
}

// SetCounters gets a reference to the given NullableResumesByStatusCounters and assigns it to the Counters field.
func (o *ResumesByStatusResponse) SetCounters(v ResumesByStatusCounters) {
	o.Counters.Set(&v)
}
// SetCountersNil sets the value for Counters to be an explicit nil
func (o *ResumesByStatusResponse) SetCountersNil() {
	o.Counters.Set(nil)
}

// UnsetCounters ensures that no value is present for Counters, not even an explicit nil
func (o *ResumesByStatusResponse) UnsetCounters() {
	o.Counters.Unset()
}

// GetNotPublished returns the NotPublished field value
func (o *ResumesByStatusResponse) GetNotPublished() []ResumesSuitableResumeItem {
	if o == nil {
		var ret []ResumesSuitableResumeItem
		return ret
	}

	return o.NotPublished
}

// GetNotPublishedOk returns a tuple with the NotPublished field value
// and a boolean to check if the value has been set.
func (o *ResumesByStatusResponse) GetNotPublishedOk() ([]ResumesSuitableResumeItem, bool) {
	if o == nil {
		return nil, false
	}
	return o.NotPublished, true
}

// SetNotPublished sets field value
func (o *ResumesByStatusResponse) SetNotPublished(v []ResumesSuitableResumeItem) {
	o.NotPublished = v
}

// GetSuitable returns the Suitable field value
func (o *ResumesByStatusResponse) GetSuitable() []ResumesSuitableResumeItem {
	if o == nil {
		var ret []ResumesSuitableResumeItem
		return ret
	}

	return o.Suitable
}

// GetSuitableOk returns a tuple with the Suitable field value
// and a boolean to check if the value has been set.
func (o *ResumesByStatusResponse) GetSuitableOk() ([]ResumesSuitableResumeItem, bool) {
	if o == nil {
		return nil, false
	}
	return o.Suitable, true
}

// SetSuitable sets field value
func (o *ResumesByStatusResponse) SetSuitable(v []ResumesSuitableResumeItem) {
	o.Suitable = v
}

// GetUnavailable returns the Unavailable field value
func (o *ResumesByStatusResponse) GetUnavailable() []ResumesSuitableResumeItem {
	if o == nil {
		var ret []ResumesSuitableResumeItem
		return ret
	}

	return o.Unavailable
}

// GetUnavailableOk returns a tuple with the Unavailable field value
// and a boolean to check if the value has been set.
func (o *ResumesByStatusResponse) GetUnavailableOk() ([]ResumesSuitableResumeItem, bool) {
	if o == nil {
		return nil, false
	}
	return o.Unavailable, true
}

// SetUnavailable sets field value
func (o *ResumesByStatusResponse) SetUnavailable(v []ResumesSuitableResumeItem) {
	o.Unavailable = v
}

func (o ResumesByStatusResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ResumesByStatusResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["already_applied"] = o.AlreadyApplied
	if o.Counters.IsSet() {
		toSerialize["counters"] = o.Counters.Get()
	}
	toSerialize["not_published"] = o.NotPublished
	toSerialize["suitable"] = o.Suitable
	toSerialize["unavailable"] = o.Unavailable
	return toSerialize, nil
}

func (o *ResumesByStatusResponse) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"already_applied",
		"not_published",
		"suitable",
		"unavailable",
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

	varResumesByStatusResponse := _ResumesByStatusResponse{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varResumesByStatusResponse)

	if err != nil {
		return err
	}

	*o = ResumesByStatusResponse(varResumesByStatusResponse)

	return err
}

type NullableResumesByStatusResponse struct {
	value *ResumesByStatusResponse
	isSet bool
}

func (v NullableResumesByStatusResponse) Get() *ResumesByStatusResponse {
	return v.value
}

func (v *NullableResumesByStatusResponse) Set(val *ResumesByStatusResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableResumesByStatusResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableResumesByStatusResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableResumesByStatusResponse(val *ResumesByStatusResponse) *NullableResumesByStatusResponse {
	return &NullableResumesByStatusResponse{value: val, isSet: true}
}

func (v NullableResumesByStatusResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableResumesByStatusResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


