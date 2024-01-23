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

// checks if the VacancyCreate type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &VacancyCreate{}

// VacancyCreate Поля, передаваемые в запросе на создание вакансии
type VacancyCreate struct {
	// Указание, что вакансия доступна для соискателей с инвалидностью
	AcceptHandicapped *bool `json:"accept_handicapped,omitempty"`
	// Разрешен ли отклик на вакансию неполным резюме
	AcceptIncompleteResumes *bool `json:"accept_incomplete_resumes,omitempty"`
	// Указание, что вакансия доступна для соискателей старше 14 лет [подробнее](https://github.com/hhru/api/blob/master/docs/employer_vacancies_accept_kids.md#accept-kids)
	AcceptKids *bool `json:"accept_kids,omitempty"`
	// Указание, что вакансия доступна с временным трудоустройством
	AcceptTemporary NullableBool `json:"accept_temporary,omitempty"`
	Address NullableVacancyAddress `json:"address,omitempty"`
	// Возможность [переписки с кандидатами](https://inboxemp.hh.ru/) по данной вакансии
	AllowMessages *bool `json:"allow_messages,omitempty"`
	BrandedTemplate NullableVacancyBrandedTemplate `json:"branded_template,omitempty"`
	// Внутренний код вакансии
	Code NullableString `json:"code,omitempty"`
	Contacts NullableVacancyContacts `json:"contacts,omitempty"`
	// Название компании для анонимных вакансий (`type.id=anonymous`), например \"крупный российский банк\". Поле становится обязательным при публикации вакансии **анонимного** типа
	CustomEmployerName NullableString `json:"custom_employer_name,omitempty"`
	Department NullableVacancyDepartment `json:"department,omitempty"`
	// Список требуемых категорий водительских прав
	DriverLicenseTypes []VacancyDriverLicenceTypeItem `json:"driver_license_types,omitempty"`
	Employment NullableVacancyEmployment `json:"employment,omitempty"`
	Experience NullableVacancyExperience `json:"experience,omitempty"`
	// Список ключевых навыков, не более 30
	KeySkills []VacancyKeySkillItem `json:"key_skills,omitempty"`
	// Список языков вакансии
	Languages []VacancyLanguage `json:"languages,omitempty"`
	// Список профессиональных ролей
	ProfessionalRoles []VacancyProfessionalRoleItem `json:"professional_roles"`
	// Обязательно ли заполнять сообщение при отклике на вакансию
	ResponseLetterRequired *bool `json:"response_letter_required,omitempty"`
	// Уведомлять ли менеджера о новых откликах
	ResponseNotifications *bool `json:"response_notifications,omitempty"`
	// URL отклика для прямых вакансий (`type.id=direct`)
	ResponseUrl NullableString `json:"response_url,omitempty"`
	Salary NullableVacancySalary `json:"salary,omitempty"`
	Schedule NullableVacancySchedule `json:"schedule,omitempty"`
	Test NullableVacancyDraftTest `json:"test,omitempty"`
	// Список рабочих дней
	WorkingDays []VacancyWorkingDayItem `json:"working_days,omitempty"`
	// Список с временными интервалами работы
	WorkingTimeIntervals []VacancyWorkingTimeIntervalItem `json:"working_time_intervals,omitempty"`
	// Список режимов времени работы
	WorkingTimeModes []VacancyWorkingTimeModeItem `json:"working_time_modes,omitempty"`
	Area VacancyArea `json:"area"`
	BillingType VacancyBillingType `json:"billing_type"`
	// Описание в html, не менее 200 символов
	Description string `json:"description"`
	Manager *VacancyManager `json:"manager,omitempty"`
	// Название
	Name string `json:"name"`
	// Если этот параметр передан, то у новой вакансии дополнительно будет создана связь с предыдущей вакансией (поле previous_id). Этот параметр не влияет на другие и не связан с ними, их всё равно необходимо передавать.  Должен быть равен только ID архивной вакансии. ID архивной вакансии можно получить, запросив [список архивных вакансий](#tag/Upravlenie-vakansiyami/operation/get-archived-vacancies) <a name='previous_id'></a> 
	PreviousId NullableString `json:"previous_id,omitempty"`
	Type VacancyType `json:"type"`
}

type _VacancyCreate VacancyCreate

// NewVacancyCreate instantiates a new VacancyCreate object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewVacancyCreate(professionalRoles []VacancyProfessionalRoleItem, area VacancyArea, billingType VacancyBillingType, description string, name string, type_ VacancyType) *VacancyCreate {
	this := VacancyCreate{}
	this.ProfessionalRoles = professionalRoles
	this.Area = area
	this.BillingType = billingType
	this.Description = description
	this.Name = name
	this.Type = type_
	return &this
}

// NewVacancyCreateWithDefaults instantiates a new VacancyCreate object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewVacancyCreateWithDefaults() *VacancyCreate {
	this := VacancyCreate{}
	return &this
}

// GetAcceptHandicapped returns the AcceptHandicapped field value if set, zero value otherwise.
func (o *VacancyCreate) GetAcceptHandicapped() bool {
	if o == nil || IsNil(o.AcceptHandicapped) {
		var ret bool
		return ret
	}
	return *o.AcceptHandicapped
}

// GetAcceptHandicappedOk returns a tuple with the AcceptHandicapped field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VacancyCreate) GetAcceptHandicappedOk() (*bool, bool) {
	if o == nil || IsNil(o.AcceptHandicapped) {
		return nil, false
	}
	return o.AcceptHandicapped, true
}

// HasAcceptHandicapped returns a boolean if a field has been set.
func (o *VacancyCreate) HasAcceptHandicapped() bool {
	if o != nil && !IsNil(o.AcceptHandicapped) {
		return true
	}

	return false
}

// SetAcceptHandicapped gets a reference to the given bool and assigns it to the AcceptHandicapped field.
func (o *VacancyCreate) SetAcceptHandicapped(v bool) {
	o.AcceptHandicapped = &v
}

// GetAcceptIncompleteResumes returns the AcceptIncompleteResumes field value if set, zero value otherwise.
func (o *VacancyCreate) GetAcceptIncompleteResumes() bool {
	if o == nil || IsNil(o.AcceptIncompleteResumes) {
		var ret bool
		return ret
	}
	return *o.AcceptIncompleteResumes
}

// GetAcceptIncompleteResumesOk returns a tuple with the AcceptIncompleteResumes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VacancyCreate) GetAcceptIncompleteResumesOk() (*bool, bool) {
	if o == nil || IsNil(o.AcceptIncompleteResumes) {
		return nil, false
	}
	return o.AcceptIncompleteResumes, true
}

// HasAcceptIncompleteResumes returns a boolean if a field has been set.
func (o *VacancyCreate) HasAcceptIncompleteResumes() bool {
	if o != nil && !IsNil(o.AcceptIncompleteResumes) {
		return true
	}

	return false
}

// SetAcceptIncompleteResumes gets a reference to the given bool and assigns it to the AcceptIncompleteResumes field.
func (o *VacancyCreate) SetAcceptIncompleteResumes(v bool) {
	o.AcceptIncompleteResumes = &v
}

// GetAcceptKids returns the AcceptKids field value if set, zero value otherwise.
func (o *VacancyCreate) GetAcceptKids() bool {
	if o == nil || IsNil(o.AcceptKids) {
		var ret bool
		return ret
	}
	return *o.AcceptKids
}

// GetAcceptKidsOk returns a tuple with the AcceptKids field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VacancyCreate) GetAcceptKidsOk() (*bool, bool) {
	if o == nil || IsNil(o.AcceptKids) {
		return nil, false
	}
	return o.AcceptKids, true
}

// HasAcceptKids returns a boolean if a field has been set.
func (o *VacancyCreate) HasAcceptKids() bool {
	if o != nil && !IsNil(o.AcceptKids) {
		return true
	}

	return false
}

// SetAcceptKids gets a reference to the given bool and assigns it to the AcceptKids field.
func (o *VacancyCreate) SetAcceptKids(v bool) {
	o.AcceptKids = &v
}

// GetAcceptTemporary returns the AcceptTemporary field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *VacancyCreate) GetAcceptTemporary() bool {
	if o == nil || IsNil(o.AcceptTemporary.Get()) {
		var ret bool
		return ret
	}
	return *o.AcceptTemporary.Get()
}

// GetAcceptTemporaryOk returns a tuple with the AcceptTemporary field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *VacancyCreate) GetAcceptTemporaryOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return o.AcceptTemporary.Get(), o.AcceptTemporary.IsSet()
}

// HasAcceptTemporary returns a boolean if a field has been set.
func (o *VacancyCreate) HasAcceptTemporary() bool {
	if o != nil && o.AcceptTemporary.IsSet() {
		return true
	}

	return false
}

// SetAcceptTemporary gets a reference to the given NullableBool and assigns it to the AcceptTemporary field.
func (o *VacancyCreate) SetAcceptTemporary(v bool) {
	o.AcceptTemporary.Set(&v)
}
// SetAcceptTemporaryNil sets the value for AcceptTemporary to be an explicit nil
func (o *VacancyCreate) SetAcceptTemporaryNil() {
	o.AcceptTemporary.Set(nil)
}

// UnsetAcceptTemporary ensures that no value is present for AcceptTemporary, not even an explicit nil
func (o *VacancyCreate) UnsetAcceptTemporary() {
	o.AcceptTemporary.Unset()
}

// GetAddress returns the Address field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *VacancyCreate) GetAddress() VacancyAddress {
	if o == nil || IsNil(o.Address.Get()) {
		var ret VacancyAddress
		return ret
	}
	return *o.Address.Get()
}

// GetAddressOk returns a tuple with the Address field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *VacancyCreate) GetAddressOk() (*VacancyAddress, bool) {
	if o == nil {
		return nil, false
	}
	return o.Address.Get(), o.Address.IsSet()
}

// HasAddress returns a boolean if a field has been set.
func (o *VacancyCreate) HasAddress() bool {
	if o != nil && o.Address.IsSet() {
		return true
	}

	return false
}

// SetAddress gets a reference to the given NullableVacancyAddress and assigns it to the Address field.
func (o *VacancyCreate) SetAddress(v VacancyAddress) {
	o.Address.Set(&v)
}
// SetAddressNil sets the value for Address to be an explicit nil
func (o *VacancyCreate) SetAddressNil() {
	o.Address.Set(nil)
}

// UnsetAddress ensures that no value is present for Address, not even an explicit nil
func (o *VacancyCreate) UnsetAddress() {
	o.Address.Unset()
}

// GetAllowMessages returns the AllowMessages field value if set, zero value otherwise.
func (o *VacancyCreate) GetAllowMessages() bool {
	if o == nil || IsNil(o.AllowMessages) {
		var ret bool
		return ret
	}
	return *o.AllowMessages
}

// GetAllowMessagesOk returns a tuple with the AllowMessages field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VacancyCreate) GetAllowMessagesOk() (*bool, bool) {
	if o == nil || IsNil(o.AllowMessages) {
		return nil, false
	}
	return o.AllowMessages, true
}

// HasAllowMessages returns a boolean if a field has been set.
func (o *VacancyCreate) HasAllowMessages() bool {
	if o != nil && !IsNil(o.AllowMessages) {
		return true
	}

	return false
}

// SetAllowMessages gets a reference to the given bool and assigns it to the AllowMessages field.
func (o *VacancyCreate) SetAllowMessages(v bool) {
	o.AllowMessages = &v
}

// GetBrandedTemplate returns the BrandedTemplate field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *VacancyCreate) GetBrandedTemplate() VacancyBrandedTemplate {
	if o == nil || IsNil(o.BrandedTemplate.Get()) {
		var ret VacancyBrandedTemplate
		return ret
	}
	return *o.BrandedTemplate.Get()
}

// GetBrandedTemplateOk returns a tuple with the BrandedTemplate field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *VacancyCreate) GetBrandedTemplateOk() (*VacancyBrandedTemplate, bool) {
	if o == nil {
		return nil, false
	}
	return o.BrandedTemplate.Get(), o.BrandedTemplate.IsSet()
}

// HasBrandedTemplate returns a boolean if a field has been set.
func (o *VacancyCreate) HasBrandedTemplate() bool {
	if o != nil && o.BrandedTemplate.IsSet() {
		return true
	}

	return false
}

// SetBrandedTemplate gets a reference to the given NullableVacancyBrandedTemplate and assigns it to the BrandedTemplate field.
func (o *VacancyCreate) SetBrandedTemplate(v VacancyBrandedTemplate) {
	o.BrandedTemplate.Set(&v)
}
// SetBrandedTemplateNil sets the value for BrandedTemplate to be an explicit nil
func (o *VacancyCreate) SetBrandedTemplateNil() {
	o.BrandedTemplate.Set(nil)
}

// UnsetBrandedTemplate ensures that no value is present for BrandedTemplate, not even an explicit nil
func (o *VacancyCreate) UnsetBrandedTemplate() {
	o.BrandedTemplate.Unset()
}

// GetCode returns the Code field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *VacancyCreate) GetCode() string {
	if o == nil || IsNil(o.Code.Get()) {
		var ret string
		return ret
	}
	return *o.Code.Get()
}

// GetCodeOk returns a tuple with the Code field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *VacancyCreate) GetCodeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Code.Get(), o.Code.IsSet()
}

// HasCode returns a boolean if a field has been set.
func (o *VacancyCreate) HasCode() bool {
	if o != nil && o.Code.IsSet() {
		return true
	}

	return false
}

// SetCode gets a reference to the given NullableString and assigns it to the Code field.
func (o *VacancyCreate) SetCode(v string) {
	o.Code.Set(&v)
}
// SetCodeNil sets the value for Code to be an explicit nil
func (o *VacancyCreate) SetCodeNil() {
	o.Code.Set(nil)
}

// UnsetCode ensures that no value is present for Code, not even an explicit nil
func (o *VacancyCreate) UnsetCode() {
	o.Code.Unset()
}

// GetContacts returns the Contacts field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *VacancyCreate) GetContacts() VacancyContacts {
	if o == nil || IsNil(o.Contacts.Get()) {
		var ret VacancyContacts
		return ret
	}
	return *o.Contacts.Get()
}

// GetContactsOk returns a tuple with the Contacts field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *VacancyCreate) GetContactsOk() (*VacancyContacts, bool) {
	if o == nil {
		return nil, false
	}
	return o.Contacts.Get(), o.Contacts.IsSet()
}

// HasContacts returns a boolean if a field has been set.
func (o *VacancyCreate) HasContacts() bool {
	if o != nil && o.Contacts.IsSet() {
		return true
	}

	return false
}

// SetContacts gets a reference to the given NullableVacancyContacts and assigns it to the Contacts field.
func (o *VacancyCreate) SetContacts(v VacancyContacts) {
	o.Contacts.Set(&v)
}
// SetContactsNil sets the value for Contacts to be an explicit nil
func (o *VacancyCreate) SetContactsNil() {
	o.Contacts.Set(nil)
}

// UnsetContacts ensures that no value is present for Contacts, not even an explicit nil
func (o *VacancyCreate) UnsetContacts() {
	o.Contacts.Unset()
}

// GetCustomEmployerName returns the CustomEmployerName field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *VacancyCreate) GetCustomEmployerName() string {
	if o == nil || IsNil(o.CustomEmployerName.Get()) {
		var ret string
		return ret
	}
	return *o.CustomEmployerName.Get()
}

// GetCustomEmployerNameOk returns a tuple with the CustomEmployerName field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *VacancyCreate) GetCustomEmployerNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.CustomEmployerName.Get(), o.CustomEmployerName.IsSet()
}

// HasCustomEmployerName returns a boolean if a field has been set.
func (o *VacancyCreate) HasCustomEmployerName() bool {
	if o != nil && o.CustomEmployerName.IsSet() {
		return true
	}

	return false
}

// SetCustomEmployerName gets a reference to the given NullableString and assigns it to the CustomEmployerName field.
func (o *VacancyCreate) SetCustomEmployerName(v string) {
	o.CustomEmployerName.Set(&v)
}
// SetCustomEmployerNameNil sets the value for CustomEmployerName to be an explicit nil
func (o *VacancyCreate) SetCustomEmployerNameNil() {
	o.CustomEmployerName.Set(nil)
}

// UnsetCustomEmployerName ensures that no value is present for CustomEmployerName, not even an explicit nil
func (o *VacancyCreate) UnsetCustomEmployerName() {
	o.CustomEmployerName.Unset()
}

// GetDepartment returns the Department field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *VacancyCreate) GetDepartment() VacancyDepartment {
	if o == nil || IsNil(o.Department.Get()) {
		var ret VacancyDepartment
		return ret
	}
	return *o.Department.Get()
}

// GetDepartmentOk returns a tuple with the Department field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *VacancyCreate) GetDepartmentOk() (*VacancyDepartment, bool) {
	if o == nil {
		return nil, false
	}
	return o.Department.Get(), o.Department.IsSet()
}

// HasDepartment returns a boolean if a field has been set.
func (o *VacancyCreate) HasDepartment() bool {
	if o != nil && o.Department.IsSet() {
		return true
	}

	return false
}

// SetDepartment gets a reference to the given NullableVacancyDepartment and assigns it to the Department field.
func (o *VacancyCreate) SetDepartment(v VacancyDepartment) {
	o.Department.Set(&v)
}
// SetDepartmentNil sets the value for Department to be an explicit nil
func (o *VacancyCreate) SetDepartmentNil() {
	o.Department.Set(nil)
}

// UnsetDepartment ensures that no value is present for Department, not even an explicit nil
func (o *VacancyCreate) UnsetDepartment() {
	o.Department.Unset()
}

// GetDriverLicenseTypes returns the DriverLicenseTypes field value if set, zero value otherwise.
func (o *VacancyCreate) GetDriverLicenseTypes() []VacancyDriverLicenceTypeItem {
	if o == nil || IsNil(o.DriverLicenseTypes) {
		var ret []VacancyDriverLicenceTypeItem
		return ret
	}
	return o.DriverLicenseTypes
}

// GetDriverLicenseTypesOk returns a tuple with the DriverLicenseTypes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VacancyCreate) GetDriverLicenseTypesOk() ([]VacancyDriverLicenceTypeItem, bool) {
	if o == nil || IsNil(o.DriverLicenseTypes) {
		return nil, false
	}
	return o.DriverLicenseTypes, true
}

// HasDriverLicenseTypes returns a boolean if a field has been set.
func (o *VacancyCreate) HasDriverLicenseTypes() bool {
	if o != nil && !IsNil(o.DriverLicenseTypes) {
		return true
	}

	return false
}

// SetDriverLicenseTypes gets a reference to the given []VacancyDriverLicenceTypeItem and assigns it to the DriverLicenseTypes field.
func (o *VacancyCreate) SetDriverLicenseTypes(v []VacancyDriverLicenceTypeItem) {
	o.DriverLicenseTypes = v
}

// GetEmployment returns the Employment field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *VacancyCreate) GetEmployment() VacancyEmployment {
	if o == nil || IsNil(o.Employment.Get()) {
		var ret VacancyEmployment
		return ret
	}
	return *o.Employment.Get()
}

// GetEmploymentOk returns a tuple with the Employment field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *VacancyCreate) GetEmploymentOk() (*VacancyEmployment, bool) {
	if o == nil {
		return nil, false
	}
	return o.Employment.Get(), o.Employment.IsSet()
}

// HasEmployment returns a boolean if a field has been set.
func (o *VacancyCreate) HasEmployment() bool {
	if o != nil && o.Employment.IsSet() {
		return true
	}

	return false
}

// SetEmployment gets a reference to the given NullableVacancyEmployment and assigns it to the Employment field.
func (o *VacancyCreate) SetEmployment(v VacancyEmployment) {
	o.Employment.Set(&v)
}
// SetEmploymentNil sets the value for Employment to be an explicit nil
func (o *VacancyCreate) SetEmploymentNil() {
	o.Employment.Set(nil)
}

// UnsetEmployment ensures that no value is present for Employment, not even an explicit nil
func (o *VacancyCreate) UnsetEmployment() {
	o.Employment.Unset()
}

// GetExperience returns the Experience field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *VacancyCreate) GetExperience() VacancyExperience {
	if o == nil || IsNil(o.Experience.Get()) {
		var ret VacancyExperience
		return ret
	}
	return *o.Experience.Get()
}

// GetExperienceOk returns a tuple with the Experience field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *VacancyCreate) GetExperienceOk() (*VacancyExperience, bool) {
	if o == nil {
		return nil, false
	}
	return o.Experience.Get(), o.Experience.IsSet()
}

// HasExperience returns a boolean if a field has been set.
func (o *VacancyCreate) HasExperience() bool {
	if o != nil && o.Experience.IsSet() {
		return true
	}

	return false
}

// SetExperience gets a reference to the given NullableVacancyExperience and assigns it to the Experience field.
func (o *VacancyCreate) SetExperience(v VacancyExperience) {
	o.Experience.Set(&v)
}
// SetExperienceNil sets the value for Experience to be an explicit nil
func (o *VacancyCreate) SetExperienceNil() {
	o.Experience.Set(nil)
}

// UnsetExperience ensures that no value is present for Experience, not even an explicit nil
func (o *VacancyCreate) UnsetExperience() {
	o.Experience.Unset()
}

// GetKeySkills returns the KeySkills field value if set, zero value otherwise.
func (o *VacancyCreate) GetKeySkills() []VacancyKeySkillItem {
	if o == nil || IsNil(o.KeySkills) {
		var ret []VacancyKeySkillItem
		return ret
	}
	return o.KeySkills
}

// GetKeySkillsOk returns a tuple with the KeySkills field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VacancyCreate) GetKeySkillsOk() ([]VacancyKeySkillItem, bool) {
	if o == nil || IsNil(o.KeySkills) {
		return nil, false
	}
	return o.KeySkills, true
}

// HasKeySkills returns a boolean if a field has been set.
func (o *VacancyCreate) HasKeySkills() bool {
	if o != nil && !IsNil(o.KeySkills) {
		return true
	}

	return false
}

// SetKeySkills gets a reference to the given []VacancyKeySkillItem and assigns it to the KeySkills field.
func (o *VacancyCreate) SetKeySkills(v []VacancyKeySkillItem) {
	o.KeySkills = v
}

// GetLanguages returns the Languages field value if set, zero value otherwise.
func (o *VacancyCreate) GetLanguages() []VacancyLanguage {
	if o == nil || IsNil(o.Languages) {
		var ret []VacancyLanguage
		return ret
	}
	return o.Languages
}

// GetLanguagesOk returns a tuple with the Languages field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VacancyCreate) GetLanguagesOk() ([]VacancyLanguage, bool) {
	if o == nil || IsNil(o.Languages) {
		return nil, false
	}
	return o.Languages, true
}

// HasLanguages returns a boolean if a field has been set.
func (o *VacancyCreate) HasLanguages() bool {
	if o != nil && !IsNil(o.Languages) {
		return true
	}

	return false
}

// SetLanguages gets a reference to the given []VacancyLanguage and assigns it to the Languages field.
func (o *VacancyCreate) SetLanguages(v []VacancyLanguage) {
	o.Languages = v
}

// GetProfessionalRoles returns the ProfessionalRoles field value
func (o *VacancyCreate) GetProfessionalRoles() []VacancyProfessionalRoleItem {
	if o == nil {
		var ret []VacancyProfessionalRoleItem
		return ret
	}

	return o.ProfessionalRoles
}

// GetProfessionalRolesOk returns a tuple with the ProfessionalRoles field value
// and a boolean to check if the value has been set.
func (o *VacancyCreate) GetProfessionalRolesOk() ([]VacancyProfessionalRoleItem, bool) {
	if o == nil {
		return nil, false
	}
	return o.ProfessionalRoles, true
}

// SetProfessionalRoles sets field value
func (o *VacancyCreate) SetProfessionalRoles(v []VacancyProfessionalRoleItem) {
	o.ProfessionalRoles = v
}

// GetResponseLetterRequired returns the ResponseLetterRequired field value if set, zero value otherwise.
func (o *VacancyCreate) GetResponseLetterRequired() bool {
	if o == nil || IsNil(o.ResponseLetterRequired) {
		var ret bool
		return ret
	}
	return *o.ResponseLetterRequired
}

// GetResponseLetterRequiredOk returns a tuple with the ResponseLetterRequired field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VacancyCreate) GetResponseLetterRequiredOk() (*bool, bool) {
	if o == nil || IsNil(o.ResponseLetterRequired) {
		return nil, false
	}
	return o.ResponseLetterRequired, true
}

// HasResponseLetterRequired returns a boolean if a field has been set.
func (o *VacancyCreate) HasResponseLetterRequired() bool {
	if o != nil && !IsNil(o.ResponseLetterRequired) {
		return true
	}

	return false
}

// SetResponseLetterRequired gets a reference to the given bool and assigns it to the ResponseLetterRequired field.
func (o *VacancyCreate) SetResponseLetterRequired(v bool) {
	o.ResponseLetterRequired = &v
}

// GetResponseNotifications returns the ResponseNotifications field value if set, zero value otherwise.
func (o *VacancyCreate) GetResponseNotifications() bool {
	if o == nil || IsNil(o.ResponseNotifications) {
		var ret bool
		return ret
	}
	return *o.ResponseNotifications
}

// GetResponseNotificationsOk returns a tuple with the ResponseNotifications field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VacancyCreate) GetResponseNotificationsOk() (*bool, bool) {
	if o == nil || IsNil(o.ResponseNotifications) {
		return nil, false
	}
	return o.ResponseNotifications, true
}

// HasResponseNotifications returns a boolean if a field has been set.
func (o *VacancyCreate) HasResponseNotifications() bool {
	if o != nil && !IsNil(o.ResponseNotifications) {
		return true
	}

	return false
}

// SetResponseNotifications gets a reference to the given bool and assigns it to the ResponseNotifications field.
func (o *VacancyCreate) SetResponseNotifications(v bool) {
	o.ResponseNotifications = &v
}

// GetResponseUrl returns the ResponseUrl field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *VacancyCreate) GetResponseUrl() string {
	if o == nil || IsNil(o.ResponseUrl.Get()) {
		var ret string
		return ret
	}
	return *o.ResponseUrl.Get()
}

// GetResponseUrlOk returns a tuple with the ResponseUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *VacancyCreate) GetResponseUrlOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.ResponseUrl.Get(), o.ResponseUrl.IsSet()
}

// HasResponseUrl returns a boolean if a field has been set.
func (o *VacancyCreate) HasResponseUrl() bool {
	if o != nil && o.ResponseUrl.IsSet() {
		return true
	}

	return false
}

// SetResponseUrl gets a reference to the given NullableString and assigns it to the ResponseUrl field.
func (o *VacancyCreate) SetResponseUrl(v string) {
	o.ResponseUrl.Set(&v)
}
// SetResponseUrlNil sets the value for ResponseUrl to be an explicit nil
func (o *VacancyCreate) SetResponseUrlNil() {
	o.ResponseUrl.Set(nil)
}

// UnsetResponseUrl ensures that no value is present for ResponseUrl, not even an explicit nil
func (o *VacancyCreate) UnsetResponseUrl() {
	o.ResponseUrl.Unset()
}

// GetSalary returns the Salary field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *VacancyCreate) GetSalary() VacancySalary {
	if o == nil || IsNil(o.Salary.Get()) {
		var ret VacancySalary
		return ret
	}
	return *o.Salary.Get()
}

// GetSalaryOk returns a tuple with the Salary field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *VacancyCreate) GetSalaryOk() (*VacancySalary, bool) {
	if o == nil {
		return nil, false
	}
	return o.Salary.Get(), o.Salary.IsSet()
}

// HasSalary returns a boolean if a field has been set.
func (o *VacancyCreate) HasSalary() bool {
	if o != nil && o.Salary.IsSet() {
		return true
	}

	return false
}

// SetSalary gets a reference to the given NullableVacancySalary and assigns it to the Salary field.
func (o *VacancyCreate) SetSalary(v VacancySalary) {
	o.Salary.Set(&v)
}
// SetSalaryNil sets the value for Salary to be an explicit nil
func (o *VacancyCreate) SetSalaryNil() {
	o.Salary.Set(nil)
}

// UnsetSalary ensures that no value is present for Salary, not even an explicit nil
func (o *VacancyCreate) UnsetSalary() {
	o.Salary.Unset()
}

// GetSchedule returns the Schedule field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *VacancyCreate) GetSchedule() VacancySchedule {
	if o == nil || IsNil(o.Schedule.Get()) {
		var ret VacancySchedule
		return ret
	}
	return *o.Schedule.Get()
}

// GetScheduleOk returns a tuple with the Schedule field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *VacancyCreate) GetScheduleOk() (*VacancySchedule, bool) {
	if o == nil {
		return nil, false
	}
	return o.Schedule.Get(), o.Schedule.IsSet()
}

// HasSchedule returns a boolean if a field has been set.
func (o *VacancyCreate) HasSchedule() bool {
	if o != nil && o.Schedule.IsSet() {
		return true
	}

	return false
}

// SetSchedule gets a reference to the given NullableVacancySchedule and assigns it to the Schedule field.
func (o *VacancyCreate) SetSchedule(v VacancySchedule) {
	o.Schedule.Set(&v)
}
// SetScheduleNil sets the value for Schedule to be an explicit nil
func (o *VacancyCreate) SetScheduleNil() {
	o.Schedule.Set(nil)
}

// UnsetSchedule ensures that no value is present for Schedule, not even an explicit nil
func (o *VacancyCreate) UnsetSchedule() {
	o.Schedule.Unset()
}

// GetTest returns the Test field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *VacancyCreate) GetTest() VacancyDraftTest {
	if o == nil || IsNil(o.Test.Get()) {
		var ret VacancyDraftTest
		return ret
	}
	return *o.Test.Get()
}

// GetTestOk returns a tuple with the Test field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *VacancyCreate) GetTestOk() (*VacancyDraftTest, bool) {
	if o == nil {
		return nil, false
	}
	return o.Test.Get(), o.Test.IsSet()
}

// HasTest returns a boolean if a field has been set.
func (o *VacancyCreate) HasTest() bool {
	if o != nil && o.Test.IsSet() {
		return true
	}

	return false
}

// SetTest gets a reference to the given NullableVacancyDraftTest and assigns it to the Test field.
func (o *VacancyCreate) SetTest(v VacancyDraftTest) {
	o.Test.Set(&v)
}
// SetTestNil sets the value for Test to be an explicit nil
func (o *VacancyCreate) SetTestNil() {
	o.Test.Set(nil)
}

// UnsetTest ensures that no value is present for Test, not even an explicit nil
func (o *VacancyCreate) UnsetTest() {
	o.Test.Unset()
}

// GetWorkingDays returns the WorkingDays field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *VacancyCreate) GetWorkingDays() []VacancyWorkingDayItem {
	if o == nil {
		var ret []VacancyWorkingDayItem
		return ret
	}
	return o.WorkingDays
}

// GetWorkingDaysOk returns a tuple with the WorkingDays field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *VacancyCreate) GetWorkingDaysOk() ([]VacancyWorkingDayItem, bool) {
	if o == nil || IsNil(o.WorkingDays) {
		return nil, false
	}
	return o.WorkingDays, true
}

// HasWorkingDays returns a boolean if a field has been set.
func (o *VacancyCreate) HasWorkingDays() bool {
	if o != nil && IsNil(o.WorkingDays) {
		return true
	}

	return false
}

// SetWorkingDays gets a reference to the given []VacancyWorkingDayItem and assigns it to the WorkingDays field.
func (o *VacancyCreate) SetWorkingDays(v []VacancyWorkingDayItem) {
	o.WorkingDays = v
}

// GetWorkingTimeIntervals returns the WorkingTimeIntervals field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *VacancyCreate) GetWorkingTimeIntervals() []VacancyWorkingTimeIntervalItem {
	if o == nil {
		var ret []VacancyWorkingTimeIntervalItem
		return ret
	}
	return o.WorkingTimeIntervals
}

// GetWorkingTimeIntervalsOk returns a tuple with the WorkingTimeIntervals field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *VacancyCreate) GetWorkingTimeIntervalsOk() ([]VacancyWorkingTimeIntervalItem, bool) {
	if o == nil || IsNil(o.WorkingTimeIntervals) {
		return nil, false
	}
	return o.WorkingTimeIntervals, true
}

// HasWorkingTimeIntervals returns a boolean if a field has been set.
func (o *VacancyCreate) HasWorkingTimeIntervals() bool {
	if o != nil && IsNil(o.WorkingTimeIntervals) {
		return true
	}

	return false
}

// SetWorkingTimeIntervals gets a reference to the given []VacancyWorkingTimeIntervalItem and assigns it to the WorkingTimeIntervals field.
func (o *VacancyCreate) SetWorkingTimeIntervals(v []VacancyWorkingTimeIntervalItem) {
	o.WorkingTimeIntervals = v
}

// GetWorkingTimeModes returns the WorkingTimeModes field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *VacancyCreate) GetWorkingTimeModes() []VacancyWorkingTimeModeItem {
	if o == nil {
		var ret []VacancyWorkingTimeModeItem
		return ret
	}
	return o.WorkingTimeModes
}

// GetWorkingTimeModesOk returns a tuple with the WorkingTimeModes field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *VacancyCreate) GetWorkingTimeModesOk() ([]VacancyWorkingTimeModeItem, bool) {
	if o == nil || IsNil(o.WorkingTimeModes) {
		return nil, false
	}
	return o.WorkingTimeModes, true
}

// HasWorkingTimeModes returns a boolean if a field has been set.
func (o *VacancyCreate) HasWorkingTimeModes() bool {
	if o != nil && IsNil(o.WorkingTimeModes) {
		return true
	}

	return false
}

// SetWorkingTimeModes gets a reference to the given []VacancyWorkingTimeModeItem and assigns it to the WorkingTimeModes field.
func (o *VacancyCreate) SetWorkingTimeModes(v []VacancyWorkingTimeModeItem) {
	o.WorkingTimeModes = v
}

// GetArea returns the Area field value
func (o *VacancyCreate) GetArea() VacancyArea {
	if o == nil {
		var ret VacancyArea
		return ret
	}

	return o.Area
}

// GetAreaOk returns a tuple with the Area field value
// and a boolean to check if the value has been set.
func (o *VacancyCreate) GetAreaOk() (*VacancyArea, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Area, true
}

// SetArea sets field value
func (o *VacancyCreate) SetArea(v VacancyArea) {
	o.Area = v
}

// GetBillingType returns the BillingType field value
func (o *VacancyCreate) GetBillingType() VacancyBillingType {
	if o == nil {
		var ret VacancyBillingType
		return ret
	}

	return o.BillingType
}

// GetBillingTypeOk returns a tuple with the BillingType field value
// and a boolean to check if the value has been set.
func (o *VacancyCreate) GetBillingTypeOk() (*VacancyBillingType, bool) {
	if o == nil {
		return nil, false
	}
	return &o.BillingType, true
}

// SetBillingType sets field value
func (o *VacancyCreate) SetBillingType(v VacancyBillingType) {
	o.BillingType = v
}

// GetDescription returns the Description field value
func (o *VacancyCreate) GetDescription() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Description
}

// GetDescriptionOk returns a tuple with the Description field value
// and a boolean to check if the value has been set.
func (o *VacancyCreate) GetDescriptionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Description, true
}

// SetDescription sets field value
func (o *VacancyCreate) SetDescription(v string) {
	o.Description = v
}

// GetManager returns the Manager field value if set, zero value otherwise.
func (o *VacancyCreate) GetManager() VacancyManager {
	if o == nil || IsNil(o.Manager) {
		var ret VacancyManager
		return ret
	}
	return *o.Manager
}

// GetManagerOk returns a tuple with the Manager field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VacancyCreate) GetManagerOk() (*VacancyManager, bool) {
	if o == nil || IsNil(o.Manager) {
		return nil, false
	}
	return o.Manager, true
}

// HasManager returns a boolean if a field has been set.
func (o *VacancyCreate) HasManager() bool {
	if o != nil && !IsNil(o.Manager) {
		return true
	}

	return false
}

// SetManager gets a reference to the given VacancyManager and assigns it to the Manager field.
func (o *VacancyCreate) SetManager(v VacancyManager) {
	o.Manager = &v
}

// GetName returns the Name field value
func (o *VacancyCreate) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *VacancyCreate) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *VacancyCreate) SetName(v string) {
	o.Name = v
}

// GetPreviousId returns the PreviousId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *VacancyCreate) GetPreviousId() string {
	if o == nil || IsNil(o.PreviousId.Get()) {
		var ret string
		return ret
	}
	return *o.PreviousId.Get()
}

// GetPreviousIdOk returns a tuple with the PreviousId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *VacancyCreate) GetPreviousIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.PreviousId.Get(), o.PreviousId.IsSet()
}

// HasPreviousId returns a boolean if a field has been set.
func (o *VacancyCreate) HasPreviousId() bool {
	if o != nil && o.PreviousId.IsSet() {
		return true
	}

	return false
}

// SetPreviousId gets a reference to the given NullableString and assigns it to the PreviousId field.
func (o *VacancyCreate) SetPreviousId(v string) {
	o.PreviousId.Set(&v)
}
// SetPreviousIdNil sets the value for PreviousId to be an explicit nil
func (o *VacancyCreate) SetPreviousIdNil() {
	o.PreviousId.Set(nil)
}

// UnsetPreviousId ensures that no value is present for PreviousId, not even an explicit nil
func (o *VacancyCreate) UnsetPreviousId() {
	o.PreviousId.Unset()
}

// GetType returns the Type field value
func (o *VacancyCreate) GetType() VacancyType {
	if o == nil {
		var ret VacancyType
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *VacancyCreate) GetTypeOk() (*VacancyType, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *VacancyCreate) SetType(v VacancyType) {
	o.Type = v
}

func (o VacancyCreate) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o VacancyCreate) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.AcceptHandicapped) {
		toSerialize["accept_handicapped"] = o.AcceptHandicapped
	}
	if !IsNil(o.AcceptIncompleteResumes) {
		toSerialize["accept_incomplete_resumes"] = o.AcceptIncompleteResumes
	}
	if !IsNil(o.AcceptKids) {
		toSerialize["accept_kids"] = o.AcceptKids
	}
	if o.AcceptTemporary.IsSet() {
		toSerialize["accept_temporary"] = o.AcceptTemporary.Get()
	}
	if o.Address.IsSet() {
		toSerialize["address"] = o.Address.Get()
	}
	if !IsNil(o.AllowMessages) {
		toSerialize["allow_messages"] = o.AllowMessages
	}
	if o.BrandedTemplate.IsSet() {
		toSerialize["branded_template"] = o.BrandedTemplate.Get()
	}
	if o.Code.IsSet() {
		toSerialize["code"] = o.Code.Get()
	}
	if o.Contacts.IsSet() {
		toSerialize["contacts"] = o.Contacts.Get()
	}
	if o.CustomEmployerName.IsSet() {
		toSerialize["custom_employer_name"] = o.CustomEmployerName.Get()
	}
	if o.Department.IsSet() {
		toSerialize["department"] = o.Department.Get()
	}
	if !IsNil(o.DriverLicenseTypes) {
		toSerialize["driver_license_types"] = o.DriverLicenseTypes
	}
	if o.Employment.IsSet() {
		toSerialize["employment"] = o.Employment.Get()
	}
	if o.Experience.IsSet() {
		toSerialize["experience"] = o.Experience.Get()
	}
	if !IsNil(o.KeySkills) {
		toSerialize["key_skills"] = o.KeySkills
	}
	if !IsNil(o.Languages) {
		toSerialize["languages"] = o.Languages
	}
	toSerialize["professional_roles"] = o.ProfessionalRoles
	if !IsNil(o.ResponseLetterRequired) {
		toSerialize["response_letter_required"] = o.ResponseLetterRequired
	}
	if !IsNil(o.ResponseNotifications) {
		toSerialize["response_notifications"] = o.ResponseNotifications
	}
	if o.ResponseUrl.IsSet() {
		toSerialize["response_url"] = o.ResponseUrl.Get()
	}
	if o.Salary.IsSet() {
		toSerialize["salary"] = o.Salary.Get()
	}
	if o.Schedule.IsSet() {
		toSerialize["schedule"] = o.Schedule.Get()
	}
	if o.Test.IsSet() {
		toSerialize["test"] = o.Test.Get()
	}
	if o.WorkingDays != nil {
		toSerialize["working_days"] = o.WorkingDays
	}
	if o.WorkingTimeIntervals != nil {
		toSerialize["working_time_intervals"] = o.WorkingTimeIntervals
	}
	if o.WorkingTimeModes != nil {
		toSerialize["working_time_modes"] = o.WorkingTimeModes
	}
	toSerialize["area"] = o.Area
	toSerialize["billing_type"] = o.BillingType
	toSerialize["description"] = o.Description
	if !IsNil(o.Manager) {
		toSerialize["manager"] = o.Manager
	}
	toSerialize["name"] = o.Name
	if o.PreviousId.IsSet() {
		toSerialize["previous_id"] = o.PreviousId.Get()
	}
	toSerialize["type"] = o.Type
	return toSerialize, nil
}

func (o *VacancyCreate) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"professional_roles",
		"area",
		"billing_type",
		"description",
		"name",
		"type",
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

	varVacancyCreate := _VacancyCreate{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varVacancyCreate)

	if err != nil {
		return err
	}

	*o = VacancyCreate(varVacancyCreate)

	return err
}

type NullableVacancyCreate struct {
	value *VacancyCreate
	isSet bool
}

func (v NullableVacancyCreate) Get() *VacancyCreate {
	return v.value
}

func (v *NullableVacancyCreate) Set(val *VacancyCreate) {
	v.value = val
	v.isSet = true
}

func (v NullableVacancyCreate) IsSet() bool {
	return v.isSet
}

func (v *NullableVacancyCreate) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableVacancyCreate(val *VacancyCreate) *NullableVacancyCreate {
	return &NullableVacancyCreate{value: val, isSet: true}
}

func (v NullableVacancyCreate) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableVacancyCreate) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


