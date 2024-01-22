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

// checks if the EmployerManagersEmployerManagerItem type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &EmployerManagersEmployerManagerItem{}

// EmployerManagersEmployerManagerItem struct for EmployerManagersEmployerManagerItem
type EmployerManagersEmployerManagerItem struct {
	AdditionalPhone *EmployerManagersManagerDataAdditionalPhone `json:"additional_phone,omitempty"`
	Area EmployerManagersArea `json:"area"`
	// Адрес электронной почты менеджера
	Email string `json:"email"`
	// Имя менеджера
	FirstName *string `json:"first_name,omitempty"`
	// Полное имя менеджера
	FullName *string `json:"full_name,omitempty"`
	// Идентификатор менеджера
	Id string `json:"id"`
	// Является ли менеджер основным контактным лицом
	IsMainContactPerson *bool `json:"is_main_contact_person,omitempty"`
	// Фамилия менеджера
	LastName *string `json:"last_name,omitempty"`
	// Отчество менеджера
	MiddleName *string `json:"middle_name,omitempty"`
	// Полное имя менеджера
	// Deprecated
	Name *string `json:"name,omitempty"`
	Phone *EmployerManagersManagerDataPhone `json:"phone,omitempty"`
	// Должность менеджера
	Position string `json:"position"`
	// Количество опубликованных (активных) вакансий у данного менеджера. `null` — если у пользователя нет прав на просмотр вакансий этого менеджера
	VacanciesCount NullableFloat32 `json:"vacancies_count,omitempty"`
}

type _EmployerManagersEmployerManagerItem EmployerManagersEmployerManagerItem

// NewEmployerManagersEmployerManagerItem instantiates a new EmployerManagersEmployerManagerItem object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewEmployerManagersEmployerManagerItem(area EmployerManagersArea, email string, id string, position string) *EmployerManagersEmployerManagerItem {
	this := EmployerManagersEmployerManagerItem{}
	this.Area = area
	this.Email = email
	this.Id = id
	this.Position = position
	return &this
}

// NewEmployerManagersEmployerManagerItemWithDefaults instantiates a new EmployerManagersEmployerManagerItem object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewEmployerManagersEmployerManagerItemWithDefaults() *EmployerManagersEmployerManagerItem {
	this := EmployerManagersEmployerManagerItem{}
	return &this
}

// GetAdditionalPhone returns the AdditionalPhone field value if set, zero value otherwise.
func (o *EmployerManagersEmployerManagerItem) GetAdditionalPhone() EmployerManagersManagerDataAdditionalPhone {
	if o == nil || IsNil(o.AdditionalPhone) {
		var ret EmployerManagersManagerDataAdditionalPhone
		return ret
	}
	return *o.AdditionalPhone
}

// GetAdditionalPhoneOk returns a tuple with the AdditionalPhone field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EmployerManagersEmployerManagerItem) GetAdditionalPhoneOk() (*EmployerManagersManagerDataAdditionalPhone, bool) {
	if o == nil || IsNil(o.AdditionalPhone) {
		return nil, false
	}
	return o.AdditionalPhone, true
}

// HasAdditionalPhone returns a boolean if a field has been set.
func (o *EmployerManagersEmployerManagerItem) HasAdditionalPhone() bool {
	if o != nil && !IsNil(o.AdditionalPhone) {
		return true
	}

	return false
}

// SetAdditionalPhone gets a reference to the given EmployerManagersManagerDataAdditionalPhone and assigns it to the AdditionalPhone field.
func (o *EmployerManagersEmployerManagerItem) SetAdditionalPhone(v EmployerManagersManagerDataAdditionalPhone) {
	o.AdditionalPhone = &v
}

// GetArea returns the Area field value
func (o *EmployerManagersEmployerManagerItem) GetArea() EmployerManagersArea {
	if o == nil {
		var ret EmployerManagersArea
		return ret
	}

	return o.Area
}

// GetAreaOk returns a tuple with the Area field value
// and a boolean to check if the value has been set.
func (o *EmployerManagersEmployerManagerItem) GetAreaOk() (*EmployerManagersArea, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Area, true
}

// SetArea sets field value
func (o *EmployerManagersEmployerManagerItem) SetArea(v EmployerManagersArea) {
	o.Area = v
}

// GetEmail returns the Email field value
func (o *EmployerManagersEmployerManagerItem) GetEmail() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Email
}

// GetEmailOk returns a tuple with the Email field value
// and a boolean to check if the value has been set.
func (o *EmployerManagersEmployerManagerItem) GetEmailOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Email, true
}

// SetEmail sets field value
func (o *EmployerManagersEmployerManagerItem) SetEmail(v string) {
	o.Email = v
}

// GetFirstName returns the FirstName field value if set, zero value otherwise.
func (o *EmployerManagersEmployerManagerItem) GetFirstName() string {
	if o == nil || IsNil(o.FirstName) {
		var ret string
		return ret
	}
	return *o.FirstName
}

// GetFirstNameOk returns a tuple with the FirstName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EmployerManagersEmployerManagerItem) GetFirstNameOk() (*string, bool) {
	if o == nil || IsNil(o.FirstName) {
		return nil, false
	}
	return o.FirstName, true
}

// HasFirstName returns a boolean if a field has been set.
func (o *EmployerManagersEmployerManagerItem) HasFirstName() bool {
	if o != nil && !IsNil(o.FirstName) {
		return true
	}

	return false
}

// SetFirstName gets a reference to the given string and assigns it to the FirstName field.
func (o *EmployerManagersEmployerManagerItem) SetFirstName(v string) {
	o.FirstName = &v
}

// GetFullName returns the FullName field value if set, zero value otherwise.
func (o *EmployerManagersEmployerManagerItem) GetFullName() string {
	if o == nil || IsNil(o.FullName) {
		var ret string
		return ret
	}
	return *o.FullName
}

// GetFullNameOk returns a tuple with the FullName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EmployerManagersEmployerManagerItem) GetFullNameOk() (*string, bool) {
	if o == nil || IsNil(o.FullName) {
		return nil, false
	}
	return o.FullName, true
}

// HasFullName returns a boolean if a field has been set.
func (o *EmployerManagersEmployerManagerItem) HasFullName() bool {
	if o != nil && !IsNil(o.FullName) {
		return true
	}

	return false
}

// SetFullName gets a reference to the given string and assigns it to the FullName field.
func (o *EmployerManagersEmployerManagerItem) SetFullName(v string) {
	o.FullName = &v
}

// GetId returns the Id field value
func (o *EmployerManagersEmployerManagerItem) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *EmployerManagersEmployerManagerItem) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *EmployerManagersEmployerManagerItem) SetId(v string) {
	o.Id = v
}

// GetIsMainContactPerson returns the IsMainContactPerson field value if set, zero value otherwise.
func (o *EmployerManagersEmployerManagerItem) GetIsMainContactPerson() bool {
	if o == nil || IsNil(o.IsMainContactPerson) {
		var ret bool
		return ret
	}
	return *o.IsMainContactPerson
}

// GetIsMainContactPersonOk returns a tuple with the IsMainContactPerson field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EmployerManagersEmployerManagerItem) GetIsMainContactPersonOk() (*bool, bool) {
	if o == nil || IsNil(o.IsMainContactPerson) {
		return nil, false
	}
	return o.IsMainContactPerson, true
}

// HasIsMainContactPerson returns a boolean if a field has been set.
func (o *EmployerManagersEmployerManagerItem) HasIsMainContactPerson() bool {
	if o != nil && !IsNil(o.IsMainContactPerson) {
		return true
	}

	return false
}

// SetIsMainContactPerson gets a reference to the given bool and assigns it to the IsMainContactPerson field.
func (o *EmployerManagersEmployerManagerItem) SetIsMainContactPerson(v bool) {
	o.IsMainContactPerson = &v
}

// GetLastName returns the LastName field value if set, zero value otherwise.
func (o *EmployerManagersEmployerManagerItem) GetLastName() string {
	if o == nil || IsNil(o.LastName) {
		var ret string
		return ret
	}
	return *o.LastName
}

// GetLastNameOk returns a tuple with the LastName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EmployerManagersEmployerManagerItem) GetLastNameOk() (*string, bool) {
	if o == nil || IsNil(o.LastName) {
		return nil, false
	}
	return o.LastName, true
}

// HasLastName returns a boolean if a field has been set.
func (o *EmployerManagersEmployerManagerItem) HasLastName() bool {
	if o != nil && !IsNil(o.LastName) {
		return true
	}

	return false
}

// SetLastName gets a reference to the given string and assigns it to the LastName field.
func (o *EmployerManagersEmployerManagerItem) SetLastName(v string) {
	o.LastName = &v
}

// GetMiddleName returns the MiddleName field value if set, zero value otherwise.
func (o *EmployerManagersEmployerManagerItem) GetMiddleName() string {
	if o == nil || IsNil(o.MiddleName) {
		var ret string
		return ret
	}
	return *o.MiddleName
}

// GetMiddleNameOk returns a tuple with the MiddleName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EmployerManagersEmployerManagerItem) GetMiddleNameOk() (*string, bool) {
	if o == nil || IsNil(o.MiddleName) {
		return nil, false
	}
	return o.MiddleName, true
}

// HasMiddleName returns a boolean if a field has been set.
func (o *EmployerManagersEmployerManagerItem) HasMiddleName() bool {
	if o != nil && !IsNil(o.MiddleName) {
		return true
	}

	return false
}

// SetMiddleName gets a reference to the given string and assigns it to the MiddleName field.
func (o *EmployerManagersEmployerManagerItem) SetMiddleName(v string) {
	o.MiddleName = &v
}

// GetName returns the Name field value if set, zero value otherwise.
// Deprecated
func (o *EmployerManagersEmployerManagerItem) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
// Deprecated
func (o *EmployerManagersEmployerManagerItem) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *EmployerManagersEmployerManagerItem) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
// Deprecated
func (o *EmployerManagersEmployerManagerItem) SetName(v string) {
	o.Name = &v
}

// GetPhone returns the Phone field value if set, zero value otherwise.
func (o *EmployerManagersEmployerManagerItem) GetPhone() EmployerManagersManagerDataPhone {
	if o == nil || IsNil(o.Phone) {
		var ret EmployerManagersManagerDataPhone
		return ret
	}
	return *o.Phone
}

// GetPhoneOk returns a tuple with the Phone field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EmployerManagersEmployerManagerItem) GetPhoneOk() (*EmployerManagersManagerDataPhone, bool) {
	if o == nil || IsNil(o.Phone) {
		return nil, false
	}
	return o.Phone, true
}

// HasPhone returns a boolean if a field has been set.
func (o *EmployerManagersEmployerManagerItem) HasPhone() bool {
	if o != nil && !IsNil(o.Phone) {
		return true
	}

	return false
}

// SetPhone gets a reference to the given EmployerManagersManagerDataPhone and assigns it to the Phone field.
func (o *EmployerManagersEmployerManagerItem) SetPhone(v EmployerManagersManagerDataPhone) {
	o.Phone = &v
}

// GetPosition returns the Position field value
func (o *EmployerManagersEmployerManagerItem) GetPosition() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Position
}

// GetPositionOk returns a tuple with the Position field value
// and a boolean to check if the value has been set.
func (o *EmployerManagersEmployerManagerItem) GetPositionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Position, true
}

// SetPosition sets field value
func (o *EmployerManagersEmployerManagerItem) SetPosition(v string) {
	o.Position = v
}

// GetVacanciesCount returns the VacanciesCount field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *EmployerManagersEmployerManagerItem) GetVacanciesCount() float32 {
	if o == nil || IsNil(o.VacanciesCount.Get()) {
		var ret float32
		return ret
	}
	return *o.VacanciesCount.Get()
}

// GetVacanciesCountOk returns a tuple with the VacanciesCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *EmployerManagersEmployerManagerItem) GetVacanciesCountOk() (*float32, bool) {
	if o == nil {
		return nil, false
	}
	return o.VacanciesCount.Get(), o.VacanciesCount.IsSet()
}

// HasVacanciesCount returns a boolean if a field has been set.
func (o *EmployerManagersEmployerManagerItem) HasVacanciesCount() bool {
	if o != nil && o.VacanciesCount.IsSet() {
		return true
	}

	return false
}

// SetVacanciesCount gets a reference to the given NullableFloat32 and assigns it to the VacanciesCount field.
func (o *EmployerManagersEmployerManagerItem) SetVacanciesCount(v float32) {
	o.VacanciesCount.Set(&v)
}
// SetVacanciesCountNil sets the value for VacanciesCount to be an explicit nil
func (o *EmployerManagersEmployerManagerItem) SetVacanciesCountNil() {
	o.VacanciesCount.Set(nil)
}

// UnsetVacanciesCount ensures that no value is present for VacanciesCount, not even an explicit nil
func (o *EmployerManagersEmployerManagerItem) UnsetVacanciesCount() {
	o.VacanciesCount.Unset()
}

func (o EmployerManagersEmployerManagerItem) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o EmployerManagersEmployerManagerItem) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.AdditionalPhone) {
		toSerialize["additional_phone"] = o.AdditionalPhone
	}
	toSerialize["area"] = o.Area
	toSerialize["email"] = o.Email
	if !IsNil(o.FirstName) {
		toSerialize["first_name"] = o.FirstName
	}
	if !IsNil(o.FullName) {
		toSerialize["full_name"] = o.FullName
	}
	toSerialize["id"] = o.Id
	if !IsNil(o.IsMainContactPerson) {
		toSerialize["is_main_contact_person"] = o.IsMainContactPerson
	}
	if !IsNil(o.LastName) {
		toSerialize["last_name"] = o.LastName
	}
	if !IsNil(o.MiddleName) {
		toSerialize["middle_name"] = o.MiddleName
	}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !IsNil(o.Phone) {
		toSerialize["phone"] = o.Phone
	}
	toSerialize["position"] = o.Position
	if o.VacanciesCount.IsSet() {
		toSerialize["vacancies_count"] = o.VacanciesCount.Get()
	}
	return toSerialize, nil
}

func (o *EmployerManagersEmployerManagerItem) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"area",
		"email",
		"id",
		"position",
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

	varEmployerManagersEmployerManagerItem := _EmployerManagersEmployerManagerItem{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varEmployerManagersEmployerManagerItem)

	if err != nil {
		return err
	}

	*o = EmployerManagersEmployerManagerItem(varEmployerManagersEmployerManagerItem)

	return err
}

type NullableEmployerManagersEmployerManagerItem struct {
	value *EmployerManagersEmployerManagerItem
	isSet bool
}

func (v NullableEmployerManagersEmployerManagerItem) Get() *EmployerManagersEmployerManagerItem {
	return v.value
}

func (v *NullableEmployerManagersEmployerManagerItem) Set(val *EmployerManagersEmployerManagerItem) {
	v.value = val
	v.isSet = true
}

func (v NullableEmployerManagersEmployerManagerItem) IsSet() bool {
	return v.isSet
}

func (v *NullableEmployerManagersEmployerManagerItem) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEmployerManagersEmployerManagerItem(val *EmployerManagersEmployerManagerItem) *NullableEmployerManagersEmployerManagerItem {
	return &NullableEmployerManagersEmployerManagerItem{value: val, isSet: true}
}

func (v NullableEmployerManagersEmployerManagerItem) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEmployerManagersEmployerManagerItem) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


