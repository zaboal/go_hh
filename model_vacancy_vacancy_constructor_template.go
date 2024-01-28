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
)

// checks if the VacancyVacancyConstructorTemplate type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &VacancyVacancyConstructorTemplate{}

// VacancyVacancyConstructorTemplate Информация по примененному брендированному шаблону конструктора вакансии. Поддерживается только две картинки: верхняя — в шапке вакансии, нижняя — в самом низу
type VacancyVacancyConstructorTemplate struct {
	BottomPicture NullableVacancyPicture `json:"bottom_picture,omitempty"`
	// ID шаблона
	Id *float32 `json:"id,omitempty"`
	// Название шаблона
	Name *string `json:"name,omitempty"`
	TopPicture NullableVacancyPicture `json:"top_picture,omitempty"`
}

// NewVacancyVacancyConstructorTemplate instantiates a new VacancyVacancyConstructorTemplate object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewVacancyVacancyConstructorTemplate() *VacancyVacancyConstructorTemplate {
	this := VacancyVacancyConstructorTemplate{}
	return &this
}

// NewVacancyVacancyConstructorTemplateWithDefaults instantiates a new VacancyVacancyConstructorTemplate object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewVacancyVacancyConstructorTemplateWithDefaults() *VacancyVacancyConstructorTemplate {
	this := VacancyVacancyConstructorTemplate{}
	return &this
}

// GetBottomPicture returns the BottomPicture field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *VacancyVacancyConstructorTemplate) GetBottomPicture() VacancyPicture {
	if o == nil || IsNil(o.BottomPicture.Get()) {
		var ret VacancyPicture
		return ret
	}
	return *o.BottomPicture.Get()
}

// GetBottomPictureOk returns a tuple with the BottomPicture field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *VacancyVacancyConstructorTemplate) GetBottomPictureOk() (*VacancyPicture, bool) {
	if o == nil {
		return nil, false
	}
	return o.BottomPicture.Get(), o.BottomPicture.IsSet()
}

// HasBottomPicture returns a boolean if a field has been set.
func (o *VacancyVacancyConstructorTemplate) HasBottomPicture() bool {
	if o != nil && o.BottomPicture.IsSet() {
		return true
	}

	return false
}

// SetBottomPicture gets a reference to the given NullableVacancyPicture and assigns it to the BottomPicture field.
func (o *VacancyVacancyConstructorTemplate) SetBottomPicture(v VacancyPicture) {
	o.BottomPicture.Set(&v)
}
// SetBottomPictureNil sets the value for BottomPicture to be an explicit nil
func (o *VacancyVacancyConstructorTemplate) SetBottomPictureNil() {
	o.BottomPicture.Set(nil)
}

// UnsetBottomPicture ensures that no value is present for BottomPicture, not even an explicit nil
func (o *VacancyVacancyConstructorTemplate) UnsetBottomPicture() {
	o.BottomPicture.Unset()
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *VacancyVacancyConstructorTemplate) GetId() float32 {
	if o == nil || IsNil(o.Id) {
		var ret float32
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VacancyVacancyConstructorTemplate) GetIdOk() (*float32, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *VacancyVacancyConstructorTemplate) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given float32 and assigns it to the Id field.
func (o *VacancyVacancyConstructorTemplate) SetId(v float32) {
	o.Id = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *VacancyVacancyConstructorTemplate) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VacancyVacancyConstructorTemplate) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *VacancyVacancyConstructorTemplate) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *VacancyVacancyConstructorTemplate) SetName(v string) {
	o.Name = &v
}

// GetTopPicture returns the TopPicture field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *VacancyVacancyConstructorTemplate) GetTopPicture() VacancyPicture {
	if o == nil || IsNil(o.TopPicture.Get()) {
		var ret VacancyPicture
		return ret
	}
	return *o.TopPicture.Get()
}

// GetTopPictureOk returns a tuple with the TopPicture field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *VacancyVacancyConstructorTemplate) GetTopPictureOk() (*VacancyPicture, bool) {
	if o == nil {
		return nil, false
	}
	return o.TopPicture.Get(), o.TopPicture.IsSet()
}

// HasTopPicture returns a boolean if a field has been set.
func (o *VacancyVacancyConstructorTemplate) HasTopPicture() bool {
	if o != nil && o.TopPicture.IsSet() {
		return true
	}

	return false
}

// SetTopPicture gets a reference to the given NullableVacancyPicture and assigns it to the TopPicture field.
func (o *VacancyVacancyConstructorTemplate) SetTopPicture(v VacancyPicture) {
	o.TopPicture.Set(&v)
}
// SetTopPictureNil sets the value for TopPicture to be an explicit nil
func (o *VacancyVacancyConstructorTemplate) SetTopPictureNil() {
	o.TopPicture.Set(nil)
}

// UnsetTopPicture ensures that no value is present for TopPicture, not even an explicit nil
func (o *VacancyVacancyConstructorTemplate) UnsetTopPicture() {
	o.TopPicture.Unset()
}

func (o VacancyVacancyConstructorTemplate) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o VacancyVacancyConstructorTemplate) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.BottomPicture.IsSet() {
		toSerialize["bottom_picture"] = o.BottomPicture.Get()
	}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if o.TopPicture.IsSet() {
		toSerialize["top_picture"] = o.TopPicture.Get()
	}
	return toSerialize, nil
}

type NullableVacancyVacancyConstructorTemplate struct {
	value *VacancyVacancyConstructorTemplate
	isSet bool
}

func (v NullableVacancyVacancyConstructorTemplate) Get() *VacancyVacancyConstructorTemplate {
	return v.value
}

func (v *NullableVacancyVacancyConstructorTemplate) Set(val *VacancyVacancyConstructorTemplate) {
	v.value = val
	v.isSet = true
}

func (v NullableVacancyVacancyConstructorTemplate) IsSet() bool {
	return v.isSet
}

func (v *NullableVacancyVacancyConstructorTemplate) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableVacancyVacancyConstructorTemplate(val *VacancyVacancyConstructorTemplate) *NullableVacancyVacancyConstructorTemplate {
	return &NullableVacancyVacancyConstructorTemplate{value: val, isSet: true}
}

func (v NullableVacancyVacancyConstructorTemplate) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableVacancyVacancyConstructorTemplate) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


