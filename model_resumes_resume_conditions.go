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
)

// checks if the ResumesResumeConditions type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ResumesResumeConditions{}

// ResumesResumeConditions Условия заполнения полей резюме.  Поле, представленное объектом с полями (`fields`), само по себе может не являться необходимым, но если заполнено хотя бы одно поле объекта, необходимо заполнить и другие его поля. Пример — желаемая зарплата (`salary`) может быть не указана, но если указана сумма, то необходимо указать и валюту 
type ResumesResumeConditions struct {
	Access OneOf `json:"access,omitempty"`
	Area NullableResumesResumeConditionFieldsRequiredWithTitle `json:"area,omitempty"`
	AutoHideTime NullableResumesResumeConditionFieldsRequiredWithTitle `json:"auto_hide_time,omitempty"`
	BirthDate NullableResumesResumeConditionFieldsRequiredDateWithTitle `json:"birth_date,omitempty"`
	BusinessTripReadiness NullableResumesResumeConditionFieldsRequiredWithTitle `json:"business_trip_readiness,omitempty"`
	Citizenship NullableResumesResumeConditionFieldsRequiredCountWithTitle `json:"citizenship,omitempty"`
	Contact OneOf `json:"contact,omitempty"`
	District NullableResumesResumeConditionFieldsRequiredWithTitle `json:"district,omitempty"`
	DriverLicenseTypes NullableResumesResumeConditionFieldsRequiredCountWithTitle `json:"driver_license_types,omitempty"`
	Education OneOf `json:"education,omitempty"`
	Employment NullableResumesResumeConditionFieldsRequiredWithTitle `json:"employment,omitempty"`
	Employments NullableResumesResumeConditionFieldsRequiredCountWithTitle `json:"employments,omitempty"`
	Experience OneOf `json:"experience,omitempty"`
	FirstName OneOf `json:"first_name,omitempty"`
	Gender NullableResumesResumeConditionFieldsRequiredWithTitle `json:"gender,omitempty"`
	HasVehicle NullableResumesResumeConditionFieldsRequiredWithTitle `json:"has_vehicle,omitempty"`
	HiddenFields NullableResumesResumeConditionFieldsRequiredCountWithTitle `json:"hidden_fields,omitempty"`
	Language OneOf `json:"language,omitempty"`
	LastName OneOf `json:"last_name,omitempty"`
	Metro NullableResumesResumeConditionFieldsRequiredWithTitle `json:"metro,omitempty"`
	MiddleName OneOf `json:"middle_name,omitempty"`
	Photo NullableResumesResumeConditionFieldsRequiredWithTitle `json:"photo,omitempty"`
	Portfolio NullableResumesResumeConditionFieldsRequiredCountWithTitle `json:"portfolio,omitempty"`
	ProfessionalRoles NullableResumesResumeConditionFieldsRequiredCountWithTitle `json:"professional_roles,omitempty"`
	Recommendation OneOf `json:"recommendation,omitempty"`
	Relocation OneOf `json:"relocation,omitempty"`
	ResumeLocale NullableResumesResumeConditionFieldsRequiredWithTitle `json:"resume_locale,omitempty"`
	Salary OneOf `json:"salary,omitempty"`
	Schedule NullableResumesResumeConditionFieldsRequiredWithTitle `json:"schedule,omitempty"`
	Schedules NullableResumesResumeConditionFieldsRequiredCountWithTitle `json:"schedules,omitempty"`
	Site OneOf `json:"site,omitempty"`
	SkillSet OneOf `json:"skill_set,omitempty"`
	Skills NullableResumesResumeConditionFieldsRequiredLengthWithTitle `json:"skills,omitempty"`
	Specialization NullableResumesResumeConditionFieldsRequiredCountWithTitle `json:"specialization,omitempty"`
	Title OneOf `json:"title,omitempty"`
	TravelTime NullableResumesResumeConditionFieldsRequiredWithTitle `json:"travel_time,omitempty"`
	WorkTicket NullableResumesResumeConditionFieldsRequiredCountWithTitle `json:"work_ticket,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _ResumesResumeConditions ResumesResumeConditions

// NewResumesResumeConditions instantiates a new ResumesResumeConditions object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewResumesResumeConditions() *ResumesResumeConditions {
	this := ResumesResumeConditions{}
	return &this
}

// NewResumesResumeConditionsWithDefaults instantiates a new ResumesResumeConditions object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewResumesResumeConditionsWithDefaults() *ResumesResumeConditions {
	this := ResumesResumeConditions{}
	return &this
}

// GetAccess returns the Access field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ResumesResumeConditions) GetAccess() OneOf {
	if o == nil {
		var ret OneOf
		return ret
	}
	return o.Access
}

// GetAccessOk returns a tuple with the Access field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ResumesResumeConditions) GetAccessOk() (OneOf, bool) {
	if o == nil || IsNil(o.Access) {
		return OneOf{}, false
	}
	return o.Access, true
}

// HasAccess returns a boolean if a field has been set.
func (o *ResumesResumeConditions) HasAccess() bool {
	if o != nil && IsNil(o.Access) {
		return true
	}

	return false
}

// SetAccess gets a reference to the given OneOf and assigns it to the Access field.
func (o *ResumesResumeConditions) SetAccess(v OneOf) {
	o.Access = v
}

// GetArea returns the Area field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ResumesResumeConditions) GetArea() ResumesResumeConditionFieldsRequiredWithTitle {
	if o == nil || IsNil(o.Area.Get()) {
		var ret ResumesResumeConditionFieldsRequiredWithTitle
		return ret
	}
	return *o.Area.Get()
}

// GetAreaOk returns a tuple with the Area field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ResumesResumeConditions) GetAreaOk() (*ResumesResumeConditionFieldsRequiredWithTitle, bool) {
	if o == nil {
		return nil, false
	}
	return o.Area.Get(), o.Area.IsSet()
}

// HasArea returns a boolean if a field has been set.
func (o *ResumesResumeConditions) HasArea() bool {
	if o != nil && o.Area.IsSet() {
		return true
	}

	return false
}

// SetArea gets a reference to the given NullableResumesResumeConditionFieldsRequiredWithTitle and assigns it to the Area field.
func (o *ResumesResumeConditions) SetArea(v ResumesResumeConditionFieldsRequiredWithTitle) {
	o.Area.Set(&v)
}
// SetAreaNil sets the value for Area to be an explicit nil
func (o *ResumesResumeConditions) SetAreaNil() {
	o.Area.Set(nil)
}

// UnsetArea ensures that no value is present for Area, not even an explicit nil
func (o *ResumesResumeConditions) UnsetArea() {
	o.Area.Unset()
}

// GetAutoHideTime returns the AutoHideTime field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ResumesResumeConditions) GetAutoHideTime() ResumesResumeConditionFieldsRequiredWithTitle {
	if o == nil || IsNil(o.AutoHideTime.Get()) {
		var ret ResumesResumeConditionFieldsRequiredWithTitle
		return ret
	}
	return *o.AutoHideTime.Get()
}

// GetAutoHideTimeOk returns a tuple with the AutoHideTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ResumesResumeConditions) GetAutoHideTimeOk() (*ResumesResumeConditionFieldsRequiredWithTitle, bool) {
	if o == nil {
		return nil, false
	}
	return o.AutoHideTime.Get(), o.AutoHideTime.IsSet()
}

// HasAutoHideTime returns a boolean if a field has been set.
func (o *ResumesResumeConditions) HasAutoHideTime() bool {
	if o != nil && o.AutoHideTime.IsSet() {
		return true
	}

	return false
}

// SetAutoHideTime gets a reference to the given NullableResumesResumeConditionFieldsRequiredWithTitle and assigns it to the AutoHideTime field.
func (o *ResumesResumeConditions) SetAutoHideTime(v ResumesResumeConditionFieldsRequiredWithTitle) {
	o.AutoHideTime.Set(&v)
}
// SetAutoHideTimeNil sets the value for AutoHideTime to be an explicit nil
func (o *ResumesResumeConditions) SetAutoHideTimeNil() {
	o.AutoHideTime.Set(nil)
}

// UnsetAutoHideTime ensures that no value is present for AutoHideTime, not even an explicit nil
func (o *ResumesResumeConditions) UnsetAutoHideTime() {
	o.AutoHideTime.Unset()
}

// GetBirthDate returns the BirthDate field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ResumesResumeConditions) GetBirthDate() ResumesResumeConditionFieldsRequiredDateWithTitle {
	if o == nil || IsNil(o.BirthDate.Get()) {
		var ret ResumesResumeConditionFieldsRequiredDateWithTitle
		return ret
	}
	return *o.BirthDate.Get()
}

// GetBirthDateOk returns a tuple with the BirthDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ResumesResumeConditions) GetBirthDateOk() (*ResumesResumeConditionFieldsRequiredDateWithTitle, bool) {
	if o == nil {
		return nil, false
	}
	return o.BirthDate.Get(), o.BirthDate.IsSet()
}

// HasBirthDate returns a boolean if a field has been set.
func (o *ResumesResumeConditions) HasBirthDate() bool {
	if o != nil && o.BirthDate.IsSet() {
		return true
	}

	return false
}

// SetBirthDate gets a reference to the given NullableResumesResumeConditionFieldsRequiredDateWithTitle and assigns it to the BirthDate field.
func (o *ResumesResumeConditions) SetBirthDate(v ResumesResumeConditionFieldsRequiredDateWithTitle) {
	o.BirthDate.Set(&v)
}
// SetBirthDateNil sets the value for BirthDate to be an explicit nil
func (o *ResumesResumeConditions) SetBirthDateNil() {
	o.BirthDate.Set(nil)
}

// UnsetBirthDate ensures that no value is present for BirthDate, not even an explicit nil
func (o *ResumesResumeConditions) UnsetBirthDate() {
	o.BirthDate.Unset()
}

// GetBusinessTripReadiness returns the BusinessTripReadiness field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ResumesResumeConditions) GetBusinessTripReadiness() ResumesResumeConditionFieldsRequiredWithTitle {
	if o == nil || IsNil(o.BusinessTripReadiness.Get()) {
		var ret ResumesResumeConditionFieldsRequiredWithTitle
		return ret
	}
	return *o.BusinessTripReadiness.Get()
}

// GetBusinessTripReadinessOk returns a tuple with the BusinessTripReadiness field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ResumesResumeConditions) GetBusinessTripReadinessOk() (*ResumesResumeConditionFieldsRequiredWithTitle, bool) {
	if o == nil {
		return nil, false
	}
	return o.BusinessTripReadiness.Get(), o.BusinessTripReadiness.IsSet()
}

// HasBusinessTripReadiness returns a boolean if a field has been set.
func (o *ResumesResumeConditions) HasBusinessTripReadiness() bool {
	if o != nil && o.BusinessTripReadiness.IsSet() {
		return true
	}

	return false
}

// SetBusinessTripReadiness gets a reference to the given NullableResumesResumeConditionFieldsRequiredWithTitle and assigns it to the BusinessTripReadiness field.
func (o *ResumesResumeConditions) SetBusinessTripReadiness(v ResumesResumeConditionFieldsRequiredWithTitle) {
	o.BusinessTripReadiness.Set(&v)
}
// SetBusinessTripReadinessNil sets the value for BusinessTripReadiness to be an explicit nil
func (o *ResumesResumeConditions) SetBusinessTripReadinessNil() {
	o.BusinessTripReadiness.Set(nil)
}

// UnsetBusinessTripReadiness ensures that no value is present for BusinessTripReadiness, not even an explicit nil
func (o *ResumesResumeConditions) UnsetBusinessTripReadiness() {
	o.BusinessTripReadiness.Unset()
}

// GetCitizenship returns the Citizenship field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ResumesResumeConditions) GetCitizenship() ResumesResumeConditionFieldsRequiredCountWithTitle {
	if o == nil || IsNil(o.Citizenship.Get()) {
		var ret ResumesResumeConditionFieldsRequiredCountWithTitle
		return ret
	}
	return *o.Citizenship.Get()
}

// GetCitizenshipOk returns a tuple with the Citizenship field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ResumesResumeConditions) GetCitizenshipOk() (*ResumesResumeConditionFieldsRequiredCountWithTitle, bool) {
	if o == nil {
		return nil, false
	}
	return o.Citizenship.Get(), o.Citizenship.IsSet()
}

// HasCitizenship returns a boolean if a field has been set.
func (o *ResumesResumeConditions) HasCitizenship() bool {
	if o != nil && o.Citizenship.IsSet() {
		return true
	}

	return false
}

// SetCitizenship gets a reference to the given NullableResumesResumeConditionFieldsRequiredCountWithTitle and assigns it to the Citizenship field.
func (o *ResumesResumeConditions) SetCitizenship(v ResumesResumeConditionFieldsRequiredCountWithTitle) {
	o.Citizenship.Set(&v)
}
// SetCitizenshipNil sets the value for Citizenship to be an explicit nil
func (o *ResumesResumeConditions) SetCitizenshipNil() {
	o.Citizenship.Set(nil)
}

// UnsetCitizenship ensures that no value is present for Citizenship, not even an explicit nil
func (o *ResumesResumeConditions) UnsetCitizenship() {
	o.Citizenship.Unset()
}

// GetContact returns the Contact field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ResumesResumeConditions) GetContact() OneOf {
	if o == nil {
		var ret OneOf
		return ret
	}
	return o.Contact
}

// GetContactOk returns a tuple with the Contact field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ResumesResumeConditions) GetContactOk() (OneOf, bool) {
	if o == nil || IsNil(o.Contact) {
		return OneOf{}, false
	}
	return o.Contact, true
}

// HasContact returns a boolean if a field has been set.
func (o *ResumesResumeConditions) HasContact() bool {
	if o != nil && IsNil(o.Contact) {
		return true
	}

	return false
}

// SetContact gets a reference to the given OneOf and assigns it to the Contact field.
func (o *ResumesResumeConditions) SetContact(v OneOf) {
	o.Contact = v
}

// GetDistrict returns the District field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ResumesResumeConditions) GetDistrict() ResumesResumeConditionFieldsRequiredWithTitle {
	if o == nil || IsNil(o.District.Get()) {
		var ret ResumesResumeConditionFieldsRequiredWithTitle
		return ret
	}
	return *o.District.Get()
}

// GetDistrictOk returns a tuple with the District field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ResumesResumeConditions) GetDistrictOk() (*ResumesResumeConditionFieldsRequiredWithTitle, bool) {
	if o == nil {
		return nil, false
	}
	return o.District.Get(), o.District.IsSet()
}

// HasDistrict returns a boolean if a field has been set.
func (o *ResumesResumeConditions) HasDistrict() bool {
	if o != nil && o.District.IsSet() {
		return true
	}

	return false
}

// SetDistrict gets a reference to the given NullableResumesResumeConditionFieldsRequiredWithTitle and assigns it to the District field.
func (o *ResumesResumeConditions) SetDistrict(v ResumesResumeConditionFieldsRequiredWithTitle) {
	o.District.Set(&v)
}
// SetDistrictNil sets the value for District to be an explicit nil
func (o *ResumesResumeConditions) SetDistrictNil() {
	o.District.Set(nil)
}

// UnsetDistrict ensures that no value is present for District, not even an explicit nil
func (o *ResumesResumeConditions) UnsetDistrict() {
	o.District.Unset()
}

// GetDriverLicenseTypes returns the DriverLicenseTypes field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ResumesResumeConditions) GetDriverLicenseTypes() ResumesResumeConditionFieldsRequiredCountWithTitle {
	if o == nil || IsNil(o.DriverLicenseTypes.Get()) {
		var ret ResumesResumeConditionFieldsRequiredCountWithTitle
		return ret
	}
	return *o.DriverLicenseTypes.Get()
}

// GetDriverLicenseTypesOk returns a tuple with the DriverLicenseTypes field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ResumesResumeConditions) GetDriverLicenseTypesOk() (*ResumesResumeConditionFieldsRequiredCountWithTitle, bool) {
	if o == nil {
		return nil, false
	}
	return o.DriverLicenseTypes.Get(), o.DriverLicenseTypes.IsSet()
}

// HasDriverLicenseTypes returns a boolean if a field has been set.
func (o *ResumesResumeConditions) HasDriverLicenseTypes() bool {
	if o != nil && o.DriverLicenseTypes.IsSet() {
		return true
	}

	return false
}

// SetDriverLicenseTypes gets a reference to the given NullableResumesResumeConditionFieldsRequiredCountWithTitle and assigns it to the DriverLicenseTypes field.
func (o *ResumesResumeConditions) SetDriverLicenseTypes(v ResumesResumeConditionFieldsRequiredCountWithTitle) {
	o.DriverLicenseTypes.Set(&v)
}
// SetDriverLicenseTypesNil sets the value for DriverLicenseTypes to be an explicit nil
func (o *ResumesResumeConditions) SetDriverLicenseTypesNil() {
	o.DriverLicenseTypes.Set(nil)
}

// UnsetDriverLicenseTypes ensures that no value is present for DriverLicenseTypes, not even an explicit nil
func (o *ResumesResumeConditions) UnsetDriverLicenseTypes() {
	o.DriverLicenseTypes.Unset()
}

// GetEducation returns the Education field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ResumesResumeConditions) GetEducation() OneOf {
	if o == nil {
		var ret OneOf
		return ret
	}
	return o.Education
}

// GetEducationOk returns a tuple with the Education field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ResumesResumeConditions) GetEducationOk() (OneOf, bool) {
	if o == nil || IsNil(o.Education) {
		return OneOf{}, false
	}
	return o.Education, true
}

// HasEducation returns a boolean if a field has been set.
func (o *ResumesResumeConditions) HasEducation() bool {
	if o != nil && IsNil(o.Education) {
		return true
	}

	return false
}

// SetEducation gets a reference to the given OneOf and assigns it to the Education field.
func (o *ResumesResumeConditions) SetEducation(v OneOf) {
	o.Education = v
}

// GetEmployment returns the Employment field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ResumesResumeConditions) GetEmployment() ResumesResumeConditionFieldsRequiredWithTitle {
	if o == nil || IsNil(o.Employment.Get()) {
		var ret ResumesResumeConditionFieldsRequiredWithTitle
		return ret
	}
	return *o.Employment.Get()
}

// GetEmploymentOk returns a tuple with the Employment field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ResumesResumeConditions) GetEmploymentOk() (*ResumesResumeConditionFieldsRequiredWithTitle, bool) {
	if o == nil {
		return nil, false
	}
	return o.Employment.Get(), o.Employment.IsSet()
}

// HasEmployment returns a boolean if a field has been set.
func (o *ResumesResumeConditions) HasEmployment() bool {
	if o != nil && o.Employment.IsSet() {
		return true
	}

	return false
}

// SetEmployment gets a reference to the given NullableResumesResumeConditionFieldsRequiredWithTitle and assigns it to the Employment field.
func (o *ResumesResumeConditions) SetEmployment(v ResumesResumeConditionFieldsRequiredWithTitle) {
	o.Employment.Set(&v)
}
// SetEmploymentNil sets the value for Employment to be an explicit nil
func (o *ResumesResumeConditions) SetEmploymentNil() {
	o.Employment.Set(nil)
}

// UnsetEmployment ensures that no value is present for Employment, not even an explicit nil
func (o *ResumesResumeConditions) UnsetEmployment() {
	o.Employment.Unset()
}

// GetEmployments returns the Employments field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ResumesResumeConditions) GetEmployments() ResumesResumeConditionFieldsRequiredCountWithTitle {
	if o == nil || IsNil(o.Employments.Get()) {
		var ret ResumesResumeConditionFieldsRequiredCountWithTitle
		return ret
	}
	return *o.Employments.Get()
}

// GetEmploymentsOk returns a tuple with the Employments field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ResumesResumeConditions) GetEmploymentsOk() (*ResumesResumeConditionFieldsRequiredCountWithTitle, bool) {
	if o == nil {
		return nil, false
	}
	return o.Employments.Get(), o.Employments.IsSet()
}

// HasEmployments returns a boolean if a field has been set.
func (o *ResumesResumeConditions) HasEmployments() bool {
	if o != nil && o.Employments.IsSet() {
		return true
	}

	return false
}

// SetEmployments gets a reference to the given NullableResumesResumeConditionFieldsRequiredCountWithTitle and assigns it to the Employments field.
func (o *ResumesResumeConditions) SetEmployments(v ResumesResumeConditionFieldsRequiredCountWithTitle) {
	o.Employments.Set(&v)
}
// SetEmploymentsNil sets the value for Employments to be an explicit nil
func (o *ResumesResumeConditions) SetEmploymentsNil() {
	o.Employments.Set(nil)
}

// UnsetEmployments ensures that no value is present for Employments, not even an explicit nil
func (o *ResumesResumeConditions) UnsetEmployments() {
	o.Employments.Unset()
}

// GetExperience returns the Experience field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ResumesResumeConditions) GetExperience() OneOf {
	if o == nil {
		var ret OneOf
		return ret
	}
	return o.Experience
}

// GetExperienceOk returns a tuple with the Experience field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ResumesResumeConditions) GetExperienceOk() (OneOf, bool) {
	if o == nil || IsNil(o.Experience) {
		return OneOf{}, false
	}
	return o.Experience, true
}

// HasExperience returns a boolean if a field has been set.
func (o *ResumesResumeConditions) HasExperience() bool {
	if o != nil && IsNil(o.Experience) {
		return true
	}

	return false
}

// SetExperience gets a reference to the given OneOf and assigns it to the Experience field.
func (o *ResumesResumeConditions) SetExperience(v OneOf) {
	o.Experience = v
}

// GetFirstName returns the FirstName field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ResumesResumeConditions) GetFirstName() OneOf {
	if o == nil {
		var ret OneOf
		return ret
	}
	return o.FirstName
}

// GetFirstNameOk returns a tuple with the FirstName field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ResumesResumeConditions) GetFirstNameOk() (OneOf, bool) {
	if o == nil || IsNil(o.FirstName) {
		return OneOf{}, false
	}
	return o.FirstName, true
}

// HasFirstName returns a boolean if a field has been set.
func (o *ResumesResumeConditions) HasFirstName() bool {
	if o != nil && IsNil(o.FirstName) {
		return true
	}

	return false
}

// SetFirstName gets a reference to the given OneOf and assigns it to the FirstName field.
func (o *ResumesResumeConditions) SetFirstName(v OneOf) {
	o.FirstName = v
}

// GetGender returns the Gender field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ResumesResumeConditions) GetGender() ResumesResumeConditionFieldsRequiredWithTitle {
	if o == nil || IsNil(o.Gender.Get()) {
		var ret ResumesResumeConditionFieldsRequiredWithTitle
		return ret
	}
	return *o.Gender.Get()
}

// GetGenderOk returns a tuple with the Gender field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ResumesResumeConditions) GetGenderOk() (*ResumesResumeConditionFieldsRequiredWithTitle, bool) {
	if o == nil {
		return nil, false
	}
	return o.Gender.Get(), o.Gender.IsSet()
}

// HasGender returns a boolean if a field has been set.
func (o *ResumesResumeConditions) HasGender() bool {
	if o != nil && o.Gender.IsSet() {
		return true
	}

	return false
}

// SetGender gets a reference to the given NullableResumesResumeConditionFieldsRequiredWithTitle and assigns it to the Gender field.
func (o *ResumesResumeConditions) SetGender(v ResumesResumeConditionFieldsRequiredWithTitle) {
	o.Gender.Set(&v)
}
// SetGenderNil sets the value for Gender to be an explicit nil
func (o *ResumesResumeConditions) SetGenderNil() {
	o.Gender.Set(nil)
}

// UnsetGender ensures that no value is present for Gender, not even an explicit nil
func (o *ResumesResumeConditions) UnsetGender() {
	o.Gender.Unset()
}

// GetHasVehicle returns the HasVehicle field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ResumesResumeConditions) GetHasVehicle() ResumesResumeConditionFieldsRequiredWithTitle {
	if o == nil || IsNil(o.HasVehicle.Get()) {
		var ret ResumesResumeConditionFieldsRequiredWithTitle
		return ret
	}
	return *o.HasVehicle.Get()
}

// GetHasVehicleOk returns a tuple with the HasVehicle field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ResumesResumeConditions) GetHasVehicleOk() (*ResumesResumeConditionFieldsRequiredWithTitle, bool) {
	if o == nil {
		return nil, false
	}
	return o.HasVehicle.Get(), o.HasVehicle.IsSet()
}

// HasHasVehicle returns a boolean if a field has been set.
func (o *ResumesResumeConditions) HasHasVehicle() bool {
	if o != nil && o.HasVehicle.IsSet() {
		return true
	}

	return false
}

// SetHasVehicle gets a reference to the given NullableResumesResumeConditionFieldsRequiredWithTitle and assigns it to the HasVehicle field.
func (o *ResumesResumeConditions) SetHasVehicle(v ResumesResumeConditionFieldsRequiredWithTitle) {
	o.HasVehicle.Set(&v)
}
// SetHasVehicleNil sets the value for HasVehicle to be an explicit nil
func (o *ResumesResumeConditions) SetHasVehicleNil() {
	o.HasVehicle.Set(nil)
}

// UnsetHasVehicle ensures that no value is present for HasVehicle, not even an explicit nil
func (o *ResumesResumeConditions) UnsetHasVehicle() {
	o.HasVehicle.Unset()
}

// GetHiddenFields returns the HiddenFields field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ResumesResumeConditions) GetHiddenFields() ResumesResumeConditionFieldsRequiredCountWithTitle {
	if o == nil || IsNil(o.HiddenFields.Get()) {
		var ret ResumesResumeConditionFieldsRequiredCountWithTitle
		return ret
	}
	return *o.HiddenFields.Get()
}

// GetHiddenFieldsOk returns a tuple with the HiddenFields field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ResumesResumeConditions) GetHiddenFieldsOk() (*ResumesResumeConditionFieldsRequiredCountWithTitle, bool) {
	if o == nil {
		return nil, false
	}
	return o.HiddenFields.Get(), o.HiddenFields.IsSet()
}

// HasHiddenFields returns a boolean if a field has been set.
func (o *ResumesResumeConditions) HasHiddenFields() bool {
	if o != nil && o.HiddenFields.IsSet() {
		return true
	}

	return false
}

// SetHiddenFields gets a reference to the given NullableResumesResumeConditionFieldsRequiredCountWithTitle and assigns it to the HiddenFields field.
func (o *ResumesResumeConditions) SetHiddenFields(v ResumesResumeConditionFieldsRequiredCountWithTitle) {
	o.HiddenFields.Set(&v)
}
// SetHiddenFieldsNil sets the value for HiddenFields to be an explicit nil
func (o *ResumesResumeConditions) SetHiddenFieldsNil() {
	o.HiddenFields.Set(nil)
}

// UnsetHiddenFields ensures that no value is present for HiddenFields, not even an explicit nil
func (o *ResumesResumeConditions) UnsetHiddenFields() {
	o.HiddenFields.Unset()
}

// GetLanguage returns the Language field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ResumesResumeConditions) GetLanguage() OneOf {
	if o == nil {
		var ret OneOf
		return ret
	}
	return o.Language
}

// GetLanguageOk returns a tuple with the Language field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ResumesResumeConditions) GetLanguageOk() (OneOf, bool) {
	if o == nil || IsNil(o.Language) {
		return OneOf{}, false
	}
	return o.Language, true
}

// HasLanguage returns a boolean if a field has been set.
func (o *ResumesResumeConditions) HasLanguage() bool {
	if o != nil && IsNil(o.Language) {
		return true
	}

	return false
}

// SetLanguage gets a reference to the given OneOf and assigns it to the Language field.
func (o *ResumesResumeConditions) SetLanguage(v OneOf) {
	o.Language = v
}

// GetLastName returns the LastName field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ResumesResumeConditions) GetLastName() OneOf {
	if o == nil {
		var ret OneOf
		return ret
	}
	return o.LastName
}

// GetLastNameOk returns a tuple with the LastName field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ResumesResumeConditions) GetLastNameOk() (OneOf, bool) {
	if o == nil || IsNil(o.LastName) {
		return OneOf{}, false
	}
	return o.LastName, true
}

// HasLastName returns a boolean if a field has been set.
func (o *ResumesResumeConditions) HasLastName() bool {
	if o != nil && IsNil(o.LastName) {
		return true
	}

	return false
}

// SetLastName gets a reference to the given OneOf and assigns it to the LastName field.
func (o *ResumesResumeConditions) SetLastName(v OneOf) {
	o.LastName = v
}

// GetMetro returns the Metro field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ResumesResumeConditions) GetMetro() ResumesResumeConditionFieldsRequiredWithTitle {
	if o == nil || IsNil(o.Metro.Get()) {
		var ret ResumesResumeConditionFieldsRequiredWithTitle
		return ret
	}
	return *o.Metro.Get()
}

// GetMetroOk returns a tuple with the Metro field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ResumesResumeConditions) GetMetroOk() (*ResumesResumeConditionFieldsRequiredWithTitle, bool) {
	if o == nil {
		return nil, false
	}
	return o.Metro.Get(), o.Metro.IsSet()
}

// HasMetro returns a boolean if a field has been set.
func (o *ResumesResumeConditions) HasMetro() bool {
	if o != nil && o.Metro.IsSet() {
		return true
	}

	return false
}

// SetMetro gets a reference to the given NullableResumesResumeConditionFieldsRequiredWithTitle and assigns it to the Metro field.
func (o *ResumesResumeConditions) SetMetro(v ResumesResumeConditionFieldsRequiredWithTitle) {
	o.Metro.Set(&v)
}
// SetMetroNil sets the value for Metro to be an explicit nil
func (o *ResumesResumeConditions) SetMetroNil() {
	o.Metro.Set(nil)
}

// UnsetMetro ensures that no value is present for Metro, not even an explicit nil
func (o *ResumesResumeConditions) UnsetMetro() {
	o.Metro.Unset()
}

// GetMiddleName returns the MiddleName field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ResumesResumeConditions) GetMiddleName() OneOf {
	if o == nil {
		var ret OneOf
		return ret
	}
	return o.MiddleName
}

// GetMiddleNameOk returns a tuple with the MiddleName field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ResumesResumeConditions) GetMiddleNameOk() (OneOf, bool) {
	if o == nil || IsNil(o.MiddleName) {
		return OneOf{}, false
	}
	return o.MiddleName, true
}

// HasMiddleName returns a boolean if a field has been set.
func (o *ResumesResumeConditions) HasMiddleName() bool {
	if o != nil && IsNil(o.MiddleName) {
		return true
	}

	return false
}

// SetMiddleName gets a reference to the given OneOf and assigns it to the MiddleName field.
func (o *ResumesResumeConditions) SetMiddleName(v OneOf) {
	o.MiddleName = v
}

// GetPhoto returns the Photo field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ResumesResumeConditions) GetPhoto() ResumesResumeConditionFieldsRequiredWithTitle {
	if o == nil || IsNil(o.Photo.Get()) {
		var ret ResumesResumeConditionFieldsRequiredWithTitle
		return ret
	}
	return *o.Photo.Get()
}

// GetPhotoOk returns a tuple with the Photo field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ResumesResumeConditions) GetPhotoOk() (*ResumesResumeConditionFieldsRequiredWithTitle, bool) {
	if o == nil {
		return nil, false
	}
	return o.Photo.Get(), o.Photo.IsSet()
}

// HasPhoto returns a boolean if a field has been set.
func (o *ResumesResumeConditions) HasPhoto() bool {
	if o != nil && o.Photo.IsSet() {
		return true
	}

	return false
}

// SetPhoto gets a reference to the given NullableResumesResumeConditionFieldsRequiredWithTitle and assigns it to the Photo field.
func (o *ResumesResumeConditions) SetPhoto(v ResumesResumeConditionFieldsRequiredWithTitle) {
	o.Photo.Set(&v)
}
// SetPhotoNil sets the value for Photo to be an explicit nil
func (o *ResumesResumeConditions) SetPhotoNil() {
	o.Photo.Set(nil)
}

// UnsetPhoto ensures that no value is present for Photo, not even an explicit nil
func (o *ResumesResumeConditions) UnsetPhoto() {
	o.Photo.Unset()
}

// GetPortfolio returns the Portfolio field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ResumesResumeConditions) GetPortfolio() ResumesResumeConditionFieldsRequiredCountWithTitle {
	if o == nil || IsNil(o.Portfolio.Get()) {
		var ret ResumesResumeConditionFieldsRequiredCountWithTitle
		return ret
	}
	return *o.Portfolio.Get()
}

// GetPortfolioOk returns a tuple with the Portfolio field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ResumesResumeConditions) GetPortfolioOk() (*ResumesResumeConditionFieldsRequiredCountWithTitle, bool) {
	if o == nil {
		return nil, false
	}
	return o.Portfolio.Get(), o.Portfolio.IsSet()
}

// HasPortfolio returns a boolean if a field has been set.
func (o *ResumesResumeConditions) HasPortfolio() bool {
	if o != nil && o.Portfolio.IsSet() {
		return true
	}

	return false
}

// SetPortfolio gets a reference to the given NullableResumesResumeConditionFieldsRequiredCountWithTitle and assigns it to the Portfolio field.
func (o *ResumesResumeConditions) SetPortfolio(v ResumesResumeConditionFieldsRequiredCountWithTitle) {
	o.Portfolio.Set(&v)
}
// SetPortfolioNil sets the value for Portfolio to be an explicit nil
func (o *ResumesResumeConditions) SetPortfolioNil() {
	o.Portfolio.Set(nil)
}

// UnsetPortfolio ensures that no value is present for Portfolio, not even an explicit nil
func (o *ResumesResumeConditions) UnsetPortfolio() {
	o.Portfolio.Unset()
}

// GetProfessionalRoles returns the ProfessionalRoles field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ResumesResumeConditions) GetProfessionalRoles() ResumesResumeConditionFieldsRequiredCountWithTitle {
	if o == nil || IsNil(o.ProfessionalRoles.Get()) {
		var ret ResumesResumeConditionFieldsRequiredCountWithTitle
		return ret
	}
	return *o.ProfessionalRoles.Get()
}

// GetProfessionalRolesOk returns a tuple with the ProfessionalRoles field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ResumesResumeConditions) GetProfessionalRolesOk() (*ResumesResumeConditionFieldsRequiredCountWithTitle, bool) {
	if o == nil {
		return nil, false
	}
	return o.ProfessionalRoles.Get(), o.ProfessionalRoles.IsSet()
}

// HasProfessionalRoles returns a boolean if a field has been set.
func (o *ResumesResumeConditions) HasProfessionalRoles() bool {
	if o != nil && o.ProfessionalRoles.IsSet() {
		return true
	}

	return false
}

// SetProfessionalRoles gets a reference to the given NullableResumesResumeConditionFieldsRequiredCountWithTitle and assigns it to the ProfessionalRoles field.
func (o *ResumesResumeConditions) SetProfessionalRoles(v ResumesResumeConditionFieldsRequiredCountWithTitle) {
	o.ProfessionalRoles.Set(&v)
}
// SetProfessionalRolesNil sets the value for ProfessionalRoles to be an explicit nil
func (o *ResumesResumeConditions) SetProfessionalRolesNil() {
	o.ProfessionalRoles.Set(nil)
}

// UnsetProfessionalRoles ensures that no value is present for ProfessionalRoles, not even an explicit nil
func (o *ResumesResumeConditions) UnsetProfessionalRoles() {
	o.ProfessionalRoles.Unset()
}

// GetRecommendation returns the Recommendation field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ResumesResumeConditions) GetRecommendation() OneOf {
	if o == nil {
		var ret OneOf
		return ret
	}
	return o.Recommendation
}

// GetRecommendationOk returns a tuple with the Recommendation field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ResumesResumeConditions) GetRecommendationOk() (OneOf, bool) {
	if o == nil || IsNil(o.Recommendation) {
		return OneOf{}, false
	}
	return o.Recommendation, true
}

// HasRecommendation returns a boolean if a field has been set.
func (o *ResumesResumeConditions) HasRecommendation() bool {
	if o != nil && IsNil(o.Recommendation) {
		return true
	}

	return false
}

// SetRecommendation gets a reference to the given OneOf and assigns it to the Recommendation field.
func (o *ResumesResumeConditions) SetRecommendation(v OneOf) {
	o.Recommendation = v
}

// GetRelocation returns the Relocation field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ResumesResumeConditions) GetRelocation() OneOf {
	if o == nil {
		var ret OneOf
		return ret
	}
	return o.Relocation
}

// GetRelocationOk returns a tuple with the Relocation field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ResumesResumeConditions) GetRelocationOk() (OneOf, bool) {
	if o == nil || IsNil(o.Relocation) {
		return OneOf{}, false
	}
	return o.Relocation, true
}

// HasRelocation returns a boolean if a field has been set.
func (o *ResumesResumeConditions) HasRelocation() bool {
	if o != nil && IsNil(o.Relocation) {
		return true
	}

	return false
}

// SetRelocation gets a reference to the given OneOf and assigns it to the Relocation field.
func (o *ResumesResumeConditions) SetRelocation(v OneOf) {
	o.Relocation = v
}

// GetResumeLocale returns the ResumeLocale field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ResumesResumeConditions) GetResumeLocale() ResumesResumeConditionFieldsRequiredWithTitle {
	if o == nil || IsNil(o.ResumeLocale.Get()) {
		var ret ResumesResumeConditionFieldsRequiredWithTitle
		return ret
	}
	return *o.ResumeLocale.Get()
}

// GetResumeLocaleOk returns a tuple with the ResumeLocale field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ResumesResumeConditions) GetResumeLocaleOk() (*ResumesResumeConditionFieldsRequiredWithTitle, bool) {
	if o == nil {
		return nil, false
	}
	return o.ResumeLocale.Get(), o.ResumeLocale.IsSet()
}

// HasResumeLocale returns a boolean if a field has been set.
func (o *ResumesResumeConditions) HasResumeLocale() bool {
	if o != nil && o.ResumeLocale.IsSet() {
		return true
	}

	return false
}

// SetResumeLocale gets a reference to the given NullableResumesResumeConditionFieldsRequiredWithTitle and assigns it to the ResumeLocale field.
func (o *ResumesResumeConditions) SetResumeLocale(v ResumesResumeConditionFieldsRequiredWithTitle) {
	o.ResumeLocale.Set(&v)
}
// SetResumeLocaleNil sets the value for ResumeLocale to be an explicit nil
func (o *ResumesResumeConditions) SetResumeLocaleNil() {
	o.ResumeLocale.Set(nil)
}

// UnsetResumeLocale ensures that no value is present for ResumeLocale, not even an explicit nil
func (o *ResumesResumeConditions) UnsetResumeLocale() {
	o.ResumeLocale.Unset()
}

// GetSalary returns the Salary field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ResumesResumeConditions) GetSalary() OneOf {
	if o == nil {
		var ret OneOf
		return ret
	}
	return o.Salary
}

// GetSalaryOk returns a tuple with the Salary field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ResumesResumeConditions) GetSalaryOk() (OneOf, bool) {
	if o == nil || IsNil(o.Salary) {
		return OneOf{}, false
	}
	return o.Salary, true
}

// HasSalary returns a boolean if a field has been set.
func (o *ResumesResumeConditions) HasSalary() bool {
	if o != nil && IsNil(o.Salary) {
		return true
	}

	return false
}

// SetSalary gets a reference to the given OneOf and assigns it to the Salary field.
func (o *ResumesResumeConditions) SetSalary(v OneOf) {
	o.Salary = v
}

// GetSchedule returns the Schedule field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ResumesResumeConditions) GetSchedule() ResumesResumeConditionFieldsRequiredWithTitle {
	if o == nil || IsNil(o.Schedule.Get()) {
		var ret ResumesResumeConditionFieldsRequiredWithTitle
		return ret
	}
	return *o.Schedule.Get()
}

// GetScheduleOk returns a tuple with the Schedule field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ResumesResumeConditions) GetScheduleOk() (*ResumesResumeConditionFieldsRequiredWithTitle, bool) {
	if o == nil {
		return nil, false
	}
	return o.Schedule.Get(), o.Schedule.IsSet()
}

// HasSchedule returns a boolean if a field has been set.
func (o *ResumesResumeConditions) HasSchedule() bool {
	if o != nil && o.Schedule.IsSet() {
		return true
	}

	return false
}

// SetSchedule gets a reference to the given NullableResumesResumeConditionFieldsRequiredWithTitle and assigns it to the Schedule field.
func (o *ResumesResumeConditions) SetSchedule(v ResumesResumeConditionFieldsRequiredWithTitle) {
	o.Schedule.Set(&v)
}
// SetScheduleNil sets the value for Schedule to be an explicit nil
func (o *ResumesResumeConditions) SetScheduleNil() {
	o.Schedule.Set(nil)
}

// UnsetSchedule ensures that no value is present for Schedule, not even an explicit nil
func (o *ResumesResumeConditions) UnsetSchedule() {
	o.Schedule.Unset()
}

// GetSchedules returns the Schedules field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ResumesResumeConditions) GetSchedules() ResumesResumeConditionFieldsRequiredCountWithTitle {
	if o == nil || IsNil(o.Schedules.Get()) {
		var ret ResumesResumeConditionFieldsRequiredCountWithTitle
		return ret
	}
	return *o.Schedules.Get()
}

// GetSchedulesOk returns a tuple with the Schedules field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ResumesResumeConditions) GetSchedulesOk() (*ResumesResumeConditionFieldsRequiredCountWithTitle, bool) {
	if o == nil {
		return nil, false
	}
	return o.Schedules.Get(), o.Schedules.IsSet()
}

// HasSchedules returns a boolean if a field has been set.
func (o *ResumesResumeConditions) HasSchedules() bool {
	if o != nil && o.Schedules.IsSet() {
		return true
	}

	return false
}

// SetSchedules gets a reference to the given NullableResumesResumeConditionFieldsRequiredCountWithTitle and assigns it to the Schedules field.
func (o *ResumesResumeConditions) SetSchedules(v ResumesResumeConditionFieldsRequiredCountWithTitle) {
	o.Schedules.Set(&v)
}
// SetSchedulesNil sets the value for Schedules to be an explicit nil
func (o *ResumesResumeConditions) SetSchedulesNil() {
	o.Schedules.Set(nil)
}

// UnsetSchedules ensures that no value is present for Schedules, not even an explicit nil
func (o *ResumesResumeConditions) UnsetSchedules() {
	o.Schedules.Unset()
}

// GetSite returns the Site field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ResumesResumeConditions) GetSite() OneOf {
	if o == nil {
		var ret OneOf
		return ret
	}
	return o.Site
}

// GetSiteOk returns a tuple with the Site field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ResumesResumeConditions) GetSiteOk() (OneOf, bool) {
	if o == nil || IsNil(o.Site) {
		return OneOf{}, false
	}
	return o.Site, true
}

// HasSite returns a boolean if a field has been set.
func (o *ResumesResumeConditions) HasSite() bool {
	if o != nil && IsNil(o.Site) {
		return true
	}

	return false
}

// SetSite gets a reference to the given OneOf and assigns it to the Site field.
func (o *ResumesResumeConditions) SetSite(v OneOf) {
	o.Site = v
}

// GetSkillSet returns the SkillSet field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ResumesResumeConditions) GetSkillSet() OneOf {
	if o == nil {
		var ret OneOf
		return ret
	}
	return o.SkillSet
}

// GetSkillSetOk returns a tuple with the SkillSet field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ResumesResumeConditions) GetSkillSetOk() (OneOf, bool) {
	if o == nil || IsNil(o.SkillSet) {
		return OneOf{}, false
	}
	return o.SkillSet, true
}

// HasSkillSet returns a boolean if a field has been set.
func (o *ResumesResumeConditions) HasSkillSet() bool {
	if o != nil && IsNil(o.SkillSet) {
		return true
	}

	return false
}

// SetSkillSet gets a reference to the given OneOf and assigns it to the SkillSet field.
func (o *ResumesResumeConditions) SetSkillSet(v OneOf) {
	o.SkillSet = v
}

// GetSkills returns the Skills field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ResumesResumeConditions) GetSkills() ResumesResumeConditionFieldsRequiredLengthWithTitle {
	if o == nil || IsNil(o.Skills.Get()) {
		var ret ResumesResumeConditionFieldsRequiredLengthWithTitle
		return ret
	}
	return *o.Skills.Get()
}

// GetSkillsOk returns a tuple with the Skills field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ResumesResumeConditions) GetSkillsOk() (*ResumesResumeConditionFieldsRequiredLengthWithTitle, bool) {
	if o == nil {
		return nil, false
	}
	return o.Skills.Get(), o.Skills.IsSet()
}

// HasSkills returns a boolean if a field has been set.
func (o *ResumesResumeConditions) HasSkills() bool {
	if o != nil && o.Skills.IsSet() {
		return true
	}

	return false
}

// SetSkills gets a reference to the given NullableResumesResumeConditionFieldsRequiredLengthWithTitle and assigns it to the Skills field.
func (o *ResumesResumeConditions) SetSkills(v ResumesResumeConditionFieldsRequiredLengthWithTitle) {
	o.Skills.Set(&v)
}
// SetSkillsNil sets the value for Skills to be an explicit nil
func (o *ResumesResumeConditions) SetSkillsNil() {
	o.Skills.Set(nil)
}

// UnsetSkills ensures that no value is present for Skills, not even an explicit nil
func (o *ResumesResumeConditions) UnsetSkills() {
	o.Skills.Unset()
}

// GetSpecialization returns the Specialization field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ResumesResumeConditions) GetSpecialization() ResumesResumeConditionFieldsRequiredCountWithTitle {
	if o == nil || IsNil(o.Specialization.Get()) {
		var ret ResumesResumeConditionFieldsRequiredCountWithTitle
		return ret
	}
	return *o.Specialization.Get()
}

// GetSpecializationOk returns a tuple with the Specialization field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ResumesResumeConditions) GetSpecializationOk() (*ResumesResumeConditionFieldsRequiredCountWithTitle, bool) {
	if o == nil {
		return nil, false
	}
	return o.Specialization.Get(), o.Specialization.IsSet()
}

// HasSpecialization returns a boolean if a field has been set.
func (o *ResumesResumeConditions) HasSpecialization() bool {
	if o != nil && o.Specialization.IsSet() {
		return true
	}

	return false
}

// SetSpecialization gets a reference to the given NullableResumesResumeConditionFieldsRequiredCountWithTitle and assigns it to the Specialization field.
func (o *ResumesResumeConditions) SetSpecialization(v ResumesResumeConditionFieldsRequiredCountWithTitle) {
	o.Specialization.Set(&v)
}
// SetSpecializationNil sets the value for Specialization to be an explicit nil
func (o *ResumesResumeConditions) SetSpecializationNil() {
	o.Specialization.Set(nil)
}

// UnsetSpecialization ensures that no value is present for Specialization, not even an explicit nil
func (o *ResumesResumeConditions) UnsetSpecialization() {
	o.Specialization.Unset()
}

// GetTitle returns the Title field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ResumesResumeConditions) GetTitle() OneOf {
	if o == nil {
		var ret OneOf
		return ret
	}
	return o.Title
}

// GetTitleOk returns a tuple with the Title field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ResumesResumeConditions) GetTitleOk() (OneOf, bool) {
	if o == nil || IsNil(o.Title) {
		return OneOf{}, false
	}
	return o.Title, true
}

// HasTitle returns a boolean if a field has been set.
func (o *ResumesResumeConditions) HasTitle() bool {
	if o != nil && IsNil(o.Title) {
		return true
	}

	return false
}

// SetTitle gets a reference to the given OneOf and assigns it to the Title field.
func (o *ResumesResumeConditions) SetTitle(v OneOf) {
	o.Title = v
}

// GetTravelTime returns the TravelTime field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ResumesResumeConditions) GetTravelTime() ResumesResumeConditionFieldsRequiredWithTitle {
	if o == nil || IsNil(o.TravelTime.Get()) {
		var ret ResumesResumeConditionFieldsRequiredWithTitle
		return ret
	}
	return *o.TravelTime.Get()
}

// GetTravelTimeOk returns a tuple with the TravelTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ResumesResumeConditions) GetTravelTimeOk() (*ResumesResumeConditionFieldsRequiredWithTitle, bool) {
	if o == nil {
		return nil, false
	}
	return o.TravelTime.Get(), o.TravelTime.IsSet()
}

// HasTravelTime returns a boolean if a field has been set.
func (o *ResumesResumeConditions) HasTravelTime() bool {
	if o != nil && o.TravelTime.IsSet() {
		return true
	}

	return false
}

// SetTravelTime gets a reference to the given NullableResumesResumeConditionFieldsRequiredWithTitle and assigns it to the TravelTime field.
func (o *ResumesResumeConditions) SetTravelTime(v ResumesResumeConditionFieldsRequiredWithTitle) {
	o.TravelTime.Set(&v)
}
// SetTravelTimeNil sets the value for TravelTime to be an explicit nil
func (o *ResumesResumeConditions) SetTravelTimeNil() {
	o.TravelTime.Set(nil)
}

// UnsetTravelTime ensures that no value is present for TravelTime, not even an explicit nil
func (o *ResumesResumeConditions) UnsetTravelTime() {
	o.TravelTime.Unset()
}

// GetWorkTicket returns the WorkTicket field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ResumesResumeConditions) GetWorkTicket() ResumesResumeConditionFieldsRequiredCountWithTitle {
	if o == nil || IsNil(o.WorkTicket.Get()) {
		var ret ResumesResumeConditionFieldsRequiredCountWithTitle
		return ret
	}
	return *o.WorkTicket.Get()
}

// GetWorkTicketOk returns a tuple with the WorkTicket field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ResumesResumeConditions) GetWorkTicketOk() (*ResumesResumeConditionFieldsRequiredCountWithTitle, bool) {
	if o == nil {
		return nil, false
	}
	return o.WorkTicket.Get(), o.WorkTicket.IsSet()
}

// HasWorkTicket returns a boolean if a field has been set.
func (o *ResumesResumeConditions) HasWorkTicket() bool {
	if o != nil && o.WorkTicket.IsSet() {
		return true
	}

	return false
}

// SetWorkTicket gets a reference to the given NullableResumesResumeConditionFieldsRequiredCountWithTitle and assigns it to the WorkTicket field.
func (o *ResumesResumeConditions) SetWorkTicket(v ResumesResumeConditionFieldsRequiredCountWithTitle) {
	o.WorkTicket.Set(&v)
}
// SetWorkTicketNil sets the value for WorkTicket to be an explicit nil
func (o *ResumesResumeConditions) SetWorkTicketNil() {
	o.WorkTicket.Set(nil)
}

// UnsetWorkTicket ensures that no value is present for WorkTicket, not even an explicit nil
func (o *ResumesResumeConditions) UnsetWorkTicket() {
	o.WorkTicket.Unset()
}

func (o ResumesResumeConditions) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ResumesResumeConditions) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.Access != nil {
		toSerialize["access"] = o.Access
	}
	if o.Area.IsSet() {
		toSerialize["area"] = o.Area.Get()
	}
	if o.AutoHideTime.IsSet() {
		toSerialize["auto_hide_time"] = o.AutoHideTime.Get()
	}
	if o.BirthDate.IsSet() {
		toSerialize["birth_date"] = o.BirthDate.Get()
	}
	if o.BusinessTripReadiness.IsSet() {
		toSerialize["business_trip_readiness"] = o.BusinessTripReadiness.Get()
	}
	if o.Citizenship.IsSet() {
		toSerialize["citizenship"] = o.Citizenship.Get()
	}
	if o.Contact != nil {
		toSerialize["contact"] = o.Contact
	}
	if o.District.IsSet() {
		toSerialize["district"] = o.District.Get()
	}
	if o.DriverLicenseTypes.IsSet() {
		toSerialize["driver_license_types"] = o.DriverLicenseTypes.Get()
	}
	if o.Education != nil {
		toSerialize["education"] = o.Education
	}
	if o.Employment.IsSet() {
		toSerialize["employment"] = o.Employment.Get()
	}
	if o.Employments.IsSet() {
		toSerialize["employments"] = o.Employments.Get()
	}
	if o.Experience != nil {
		toSerialize["experience"] = o.Experience
	}
	if o.FirstName != nil {
		toSerialize["first_name"] = o.FirstName
	}
	if o.Gender.IsSet() {
		toSerialize["gender"] = o.Gender.Get()
	}
	if o.HasVehicle.IsSet() {
		toSerialize["has_vehicle"] = o.HasVehicle.Get()
	}
	if o.HiddenFields.IsSet() {
		toSerialize["hidden_fields"] = o.HiddenFields.Get()
	}
	if o.Language != nil {
		toSerialize["language"] = o.Language
	}
	if o.LastName != nil {
		toSerialize["last_name"] = o.LastName
	}
	if o.Metro.IsSet() {
		toSerialize["metro"] = o.Metro.Get()
	}
	if o.MiddleName != nil {
		toSerialize["middle_name"] = o.MiddleName
	}
	if o.Photo.IsSet() {
		toSerialize["photo"] = o.Photo.Get()
	}
	if o.Portfolio.IsSet() {
		toSerialize["portfolio"] = o.Portfolio.Get()
	}
	if o.ProfessionalRoles.IsSet() {
		toSerialize["professional_roles"] = o.ProfessionalRoles.Get()
	}
	if o.Recommendation != nil {
		toSerialize["recommendation"] = o.Recommendation
	}
	if o.Relocation != nil {
		toSerialize["relocation"] = o.Relocation
	}
	if o.ResumeLocale.IsSet() {
		toSerialize["resume_locale"] = o.ResumeLocale.Get()
	}
	if o.Salary != nil {
		toSerialize["salary"] = o.Salary
	}
	if o.Schedule.IsSet() {
		toSerialize["schedule"] = o.Schedule.Get()
	}
	if o.Schedules.IsSet() {
		toSerialize["schedules"] = o.Schedules.Get()
	}
	if o.Site != nil {
		toSerialize["site"] = o.Site
	}
	if o.SkillSet != nil {
		toSerialize["skill_set"] = o.SkillSet
	}
	if o.Skills.IsSet() {
		toSerialize["skills"] = o.Skills.Get()
	}
	if o.Specialization.IsSet() {
		toSerialize["specialization"] = o.Specialization.Get()
	}
	if o.Title != nil {
		toSerialize["title"] = o.Title
	}
	if o.TravelTime.IsSet() {
		toSerialize["travel_time"] = o.TravelTime.Get()
	}
	if o.WorkTicket.IsSet() {
		toSerialize["work_ticket"] = o.WorkTicket.Get()
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *ResumesResumeConditions) UnmarshalJSON(data []byte) (err error) {
	varResumesResumeConditions := _ResumesResumeConditions{}

	err = json.Unmarshal(data, &varResumesResumeConditions)

	if err != nil {
		return err
	}

	*o = ResumesResumeConditions(varResumesResumeConditions)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "access")
		delete(additionalProperties, "area")
		delete(additionalProperties, "auto_hide_time")
		delete(additionalProperties, "birth_date")
		delete(additionalProperties, "business_trip_readiness")
		delete(additionalProperties, "citizenship")
		delete(additionalProperties, "contact")
		delete(additionalProperties, "district")
		delete(additionalProperties, "driver_license_types")
		delete(additionalProperties, "education")
		delete(additionalProperties, "employment")
		delete(additionalProperties, "employments")
		delete(additionalProperties, "experience")
		delete(additionalProperties, "first_name")
		delete(additionalProperties, "gender")
		delete(additionalProperties, "has_vehicle")
		delete(additionalProperties, "hidden_fields")
		delete(additionalProperties, "language")
		delete(additionalProperties, "last_name")
		delete(additionalProperties, "metro")
		delete(additionalProperties, "middle_name")
		delete(additionalProperties, "photo")
		delete(additionalProperties, "portfolio")
		delete(additionalProperties, "professional_roles")
		delete(additionalProperties, "recommendation")
		delete(additionalProperties, "relocation")
		delete(additionalProperties, "resume_locale")
		delete(additionalProperties, "salary")
		delete(additionalProperties, "schedule")
		delete(additionalProperties, "schedules")
		delete(additionalProperties, "site")
		delete(additionalProperties, "skill_set")
		delete(additionalProperties, "skills")
		delete(additionalProperties, "specialization")
		delete(additionalProperties, "title")
		delete(additionalProperties, "travel_time")
		delete(additionalProperties, "work_ticket")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableResumesResumeConditions struct {
	value *ResumesResumeConditions
	isSet bool
}

func (v NullableResumesResumeConditions) Get() *ResumesResumeConditions {
	return v.value
}

func (v *NullableResumesResumeConditions) Set(val *ResumesResumeConditions) {
	v.value = val
	v.isSet = true
}

func (v NullableResumesResumeConditions) IsSet() bool {
	return v.isSet
}

func (v *NullableResumesResumeConditions) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableResumesResumeConditions(val *ResumesResumeConditions) *NullableResumesResumeConditions {
	return &NullableResumesResumeConditions{value: val, isSet: true}
}

func (v NullableResumesResumeConditions) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableResumesResumeConditions) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


