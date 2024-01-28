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

// checks if the VacanciesStandardVacancyFields type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &VacanciesStandardVacancyFields{}

// VacanciesStandardVacancyFields struct for VacanciesStandardVacancyFields
type VacanciesStandardVacancyFields struct {
	// Разрешен ли отклик на вакансию неполным резюме
	AcceptIncompleteResumes bool `json:"accept_incomplete_resumes"`
	// Указание, что вакансия доступна с временным трудоустройством
	AcceptTemporary NullableBool `json:"accept_temporary,omitempty"`
	Address NullableVacancyAddressRawOutput `json:"address,omitempty"`
	// URL для регистрации нажатия кнопки отклика (устаревшее, сейчас для регистрации нажатия кнопки отклика нужно заполнять хедер adv-context  в запросе [Резюме, сгруппированные по возможности отклика на данную вакансию](#tag/Rezyume.-Prosmotr-informacii/operation/get-resumes-by-status)) 
	// Deprecated
	AdvResponseUrl NullableString `json:"adv_response_url,omitempty"`
	// Ссылка на представление вакансии на сайте
	AlternateUrl string `json:"alternate_url"`
	// Ссылка на отклик на вакансию на сайте
	ApplyAlternateUrl string `json:"apply_alternate_url"`
	// Находится ли данная вакансия в архиве
	Archived NullableBool `json:"archived,omitempty"`
	Area IncludesArea `json:"area"`
	Contacts NullableVacancyContactsOutput `json:"contacts,omitempty"`
	// Дата и время публикации вакансии
	CreatedAt *string `json:"created_at,omitempty"`
	Department NullableVacancyDepartmentOutput `json:"department"`
	Employer VacanciesEmployerPublic `json:"employer"`
	// Информация о наличии прикрепленного тестового задании к вакансии
	HasTest bool `json:"has_test"`
	// Идентификатор вакансии
	Id string `json:"id"`
	InsiderInterview NullableVacancyInsiderInterview `json:"insider_interview,omitempty"`
	MetroStations *IncludesMetroStation `json:"metro_stations,omitempty"`
	// Название
	Name string `json:"name"`
	// Является ли данная вакансия премиум-вакансией
	Premium NullableBool `json:"premium,omitempty"`
	// Список профессиональных ролей
	ProfessionalRoles []VacancyProfessionalRoleItemOutput `json:"professional_roles"`
	// Дата и время публикации вакансии
	PublishedAt string `json:"published_at"`
	// Возвращает связи соискателя с вакансией. Значения из поля `vacancy_relation` в [справочнике полей](#tag/Obshie-spravochniki/operation/get-dictionaries)
	Relations []VacancyRelationItem `json:"relations"`
	// Обязательно ли заполнять сообщение при отклике на вакансию
	ResponseLetterRequired bool `json:"response_letter_required"`
	// URL отклика для прямых вакансий (`type.id=direct`)
	ResponseUrl NullableString `json:"response_url,omitempty"`
	Salary NullableVacancySalary `json:"salary"`
	Schedule NullableVacancyScheduleOutput `json:"schedule,omitempty"`
	// Расстояние в метрах между центром сортировки (заданной параметрами `sort_point_lat`, `sort_point_lng`) и указанным в вакансии адресом. В случае, если в адресе указаны только станции метро, выдается расстояние между центром сортировки и средней геометрической точкой указанных станций.  Значение `sort_point_distance` выдается только в случае, если заданы параметры `sort_point_lat`, `sort_point_lng`, `order_by=distance`
	SortPointDistance NullableFloat32 `json:"sort_point_distance,omitempty"`
	Type VacancyTypeOutput `json:"type"`
	// URL вакансии
	Url string `json:"url"`
	// Список рабочих дней
	WorkingDays []VacancyWorkingDayItemOutput `json:"working_days,omitempty"`
	// Список с временными интервалами работы
	WorkingTimeIntervals []VacancyWorkingTimeIntervalItemOutput `json:"working_time_intervals,omitempty"`
	// Список режимов времени работы
	WorkingTimeModes []VacancyWorkingTimeModeItemOutput `json:"working_time_modes,omitempty"`
}

type _VacanciesStandardVacancyFields VacanciesStandardVacancyFields

// NewVacanciesStandardVacancyFields instantiates a new VacanciesStandardVacancyFields object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewVacanciesStandardVacancyFields(acceptIncompleteResumes bool, alternateUrl string, applyAlternateUrl string, area IncludesArea, department NullableVacancyDepartmentOutput, employer VacanciesEmployerPublic, hasTest bool, id string, name string, professionalRoles []VacancyProfessionalRoleItemOutput, publishedAt string, relations []VacancyRelationItem, responseLetterRequired bool, salary NullableVacancySalary, type_ VacancyTypeOutput, url string) *VacanciesStandardVacancyFields {
	this := VacanciesStandardVacancyFields{}
	this.AcceptIncompleteResumes = acceptIncompleteResumes
	this.AlternateUrl = alternateUrl
	this.ApplyAlternateUrl = applyAlternateUrl
	this.Area = area
	this.Department = department
	this.Employer = employer
	this.HasTest = hasTest
	this.Id = id
	this.Name = name
	this.ProfessionalRoles = professionalRoles
	this.PublishedAt = publishedAt
	this.Relations = relations
	this.ResponseLetterRequired = responseLetterRequired
	this.Salary = salary
	this.Type = type_
	this.Url = url
	return &this
}

// NewVacanciesStandardVacancyFieldsWithDefaults instantiates a new VacanciesStandardVacancyFields object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewVacanciesStandardVacancyFieldsWithDefaults() *VacanciesStandardVacancyFields {
	this := VacanciesStandardVacancyFields{}
	return &this
}

// GetAcceptIncompleteResumes returns the AcceptIncompleteResumes field value
func (o *VacanciesStandardVacancyFields) GetAcceptIncompleteResumes() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.AcceptIncompleteResumes
}

// GetAcceptIncompleteResumesOk returns a tuple with the AcceptIncompleteResumes field value
// and a boolean to check if the value has been set.
func (o *VacanciesStandardVacancyFields) GetAcceptIncompleteResumesOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AcceptIncompleteResumes, true
}

// SetAcceptIncompleteResumes sets field value
func (o *VacanciesStandardVacancyFields) SetAcceptIncompleteResumes(v bool) {
	o.AcceptIncompleteResumes = v
}

// GetAcceptTemporary returns the AcceptTemporary field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *VacanciesStandardVacancyFields) GetAcceptTemporary() bool {
	if o == nil || IsNil(o.AcceptTemporary.Get()) {
		var ret bool
		return ret
	}
	return *o.AcceptTemporary.Get()
}

// GetAcceptTemporaryOk returns a tuple with the AcceptTemporary field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *VacanciesStandardVacancyFields) GetAcceptTemporaryOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return o.AcceptTemporary.Get(), o.AcceptTemporary.IsSet()
}

// HasAcceptTemporary returns a boolean if a field has been set.
func (o *VacanciesStandardVacancyFields) HasAcceptTemporary() bool {
	if o != nil && o.AcceptTemporary.IsSet() {
		return true
	}

	return false
}

// SetAcceptTemporary gets a reference to the given NullableBool and assigns it to the AcceptTemporary field.
func (o *VacanciesStandardVacancyFields) SetAcceptTemporary(v bool) {
	o.AcceptTemporary.Set(&v)
}
// SetAcceptTemporaryNil sets the value for AcceptTemporary to be an explicit nil
func (o *VacanciesStandardVacancyFields) SetAcceptTemporaryNil() {
	o.AcceptTemporary.Set(nil)
}

// UnsetAcceptTemporary ensures that no value is present for AcceptTemporary, not even an explicit nil
func (o *VacanciesStandardVacancyFields) UnsetAcceptTemporary() {
	o.AcceptTemporary.Unset()
}

// GetAddress returns the Address field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *VacanciesStandardVacancyFields) GetAddress() VacancyAddressRawOutput {
	if o == nil || IsNil(o.Address.Get()) {
		var ret VacancyAddressRawOutput
		return ret
	}
	return *o.Address.Get()
}

// GetAddressOk returns a tuple with the Address field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *VacanciesStandardVacancyFields) GetAddressOk() (*VacancyAddressRawOutput, bool) {
	if o == nil {
		return nil, false
	}
	return o.Address.Get(), o.Address.IsSet()
}

// HasAddress returns a boolean if a field has been set.
func (o *VacanciesStandardVacancyFields) HasAddress() bool {
	if o != nil && o.Address.IsSet() {
		return true
	}

	return false
}

// SetAddress gets a reference to the given NullableVacancyAddressRawOutput and assigns it to the Address field.
func (o *VacanciesStandardVacancyFields) SetAddress(v VacancyAddressRawOutput) {
	o.Address.Set(&v)
}
// SetAddressNil sets the value for Address to be an explicit nil
func (o *VacanciesStandardVacancyFields) SetAddressNil() {
	o.Address.Set(nil)
}

// UnsetAddress ensures that no value is present for Address, not even an explicit nil
func (o *VacanciesStandardVacancyFields) UnsetAddress() {
	o.Address.Unset()
}

// GetAdvResponseUrl returns the AdvResponseUrl field value if set, zero value otherwise (both if not set or set to explicit null).
// Deprecated
func (o *VacanciesStandardVacancyFields) GetAdvResponseUrl() string {
	if o == nil || IsNil(o.AdvResponseUrl.Get()) {
		var ret string
		return ret
	}
	return *o.AdvResponseUrl.Get()
}

// GetAdvResponseUrlOk returns a tuple with the AdvResponseUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
// Deprecated
func (o *VacanciesStandardVacancyFields) GetAdvResponseUrlOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.AdvResponseUrl.Get(), o.AdvResponseUrl.IsSet()
}

// HasAdvResponseUrl returns a boolean if a field has been set.
func (o *VacanciesStandardVacancyFields) HasAdvResponseUrl() bool {
	if o != nil && o.AdvResponseUrl.IsSet() {
		return true
	}

	return false
}

// SetAdvResponseUrl gets a reference to the given NullableString and assigns it to the AdvResponseUrl field.
// Deprecated
func (o *VacanciesStandardVacancyFields) SetAdvResponseUrl(v string) {
	o.AdvResponseUrl.Set(&v)
}
// SetAdvResponseUrlNil sets the value for AdvResponseUrl to be an explicit nil
func (o *VacanciesStandardVacancyFields) SetAdvResponseUrlNil() {
	o.AdvResponseUrl.Set(nil)
}

// UnsetAdvResponseUrl ensures that no value is present for AdvResponseUrl, not even an explicit nil
func (o *VacanciesStandardVacancyFields) UnsetAdvResponseUrl() {
	o.AdvResponseUrl.Unset()
}

// GetAlternateUrl returns the AlternateUrl field value
func (o *VacanciesStandardVacancyFields) GetAlternateUrl() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.AlternateUrl
}

// GetAlternateUrlOk returns a tuple with the AlternateUrl field value
// and a boolean to check if the value has been set.
func (o *VacanciesStandardVacancyFields) GetAlternateUrlOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AlternateUrl, true
}

// SetAlternateUrl sets field value
func (o *VacanciesStandardVacancyFields) SetAlternateUrl(v string) {
	o.AlternateUrl = v
}

// GetApplyAlternateUrl returns the ApplyAlternateUrl field value
func (o *VacanciesStandardVacancyFields) GetApplyAlternateUrl() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ApplyAlternateUrl
}

// GetApplyAlternateUrlOk returns a tuple with the ApplyAlternateUrl field value
// and a boolean to check if the value has been set.
func (o *VacanciesStandardVacancyFields) GetApplyAlternateUrlOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ApplyAlternateUrl, true
}

// SetApplyAlternateUrl sets field value
func (o *VacanciesStandardVacancyFields) SetApplyAlternateUrl(v string) {
	o.ApplyAlternateUrl = v
}

// GetArchived returns the Archived field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *VacanciesStandardVacancyFields) GetArchived() bool {
	if o == nil || IsNil(o.Archived.Get()) {
		var ret bool
		return ret
	}
	return *o.Archived.Get()
}

// GetArchivedOk returns a tuple with the Archived field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *VacanciesStandardVacancyFields) GetArchivedOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return o.Archived.Get(), o.Archived.IsSet()
}

// HasArchived returns a boolean if a field has been set.
func (o *VacanciesStandardVacancyFields) HasArchived() bool {
	if o != nil && o.Archived.IsSet() {
		return true
	}

	return false
}

// SetArchived gets a reference to the given NullableBool and assigns it to the Archived field.
func (o *VacanciesStandardVacancyFields) SetArchived(v bool) {
	o.Archived.Set(&v)
}
// SetArchivedNil sets the value for Archived to be an explicit nil
func (o *VacanciesStandardVacancyFields) SetArchivedNil() {
	o.Archived.Set(nil)
}

// UnsetArchived ensures that no value is present for Archived, not even an explicit nil
func (o *VacanciesStandardVacancyFields) UnsetArchived() {
	o.Archived.Unset()
}

// GetArea returns the Area field value
func (o *VacanciesStandardVacancyFields) GetArea() IncludesArea {
	if o == nil {
		var ret IncludesArea
		return ret
	}

	return o.Area
}

// GetAreaOk returns a tuple with the Area field value
// and a boolean to check if the value has been set.
func (o *VacanciesStandardVacancyFields) GetAreaOk() (IncludesArea, bool) {
	if o == nil {
		return IncludesArea{}, false
	}
	return o.Area, true
}

// SetArea sets field value
func (o *VacanciesStandardVacancyFields) SetArea(v IncludesArea) {
	o.Area = v
}

// GetContacts returns the Contacts field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *VacanciesStandardVacancyFields) GetContacts() VacancyContactsOutput {
	if o == nil || IsNil(o.Contacts.Get()) {
		var ret VacancyContactsOutput
		return ret
	}
	return *o.Contacts.Get()
}

// GetContactsOk returns a tuple with the Contacts field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *VacanciesStandardVacancyFields) GetContactsOk() (*VacancyContactsOutput, bool) {
	if o == nil {
		return nil, false
	}
	return o.Contacts.Get(), o.Contacts.IsSet()
}

// HasContacts returns a boolean if a field has been set.
func (o *VacanciesStandardVacancyFields) HasContacts() bool {
	if o != nil && o.Contacts.IsSet() {
		return true
	}

	return false
}

// SetContacts gets a reference to the given NullableVacancyContactsOutput and assigns it to the Contacts field.
func (o *VacanciesStandardVacancyFields) SetContacts(v VacancyContactsOutput) {
	o.Contacts.Set(&v)
}
// SetContactsNil sets the value for Contacts to be an explicit nil
func (o *VacanciesStandardVacancyFields) SetContactsNil() {
	o.Contacts.Set(nil)
}

// UnsetContacts ensures that no value is present for Contacts, not even an explicit nil
func (o *VacanciesStandardVacancyFields) UnsetContacts() {
	o.Contacts.Unset()
}

// GetCreatedAt returns the CreatedAt field value if set, zero value otherwise.
func (o *VacanciesStandardVacancyFields) GetCreatedAt() string {
	if o == nil || IsNil(o.CreatedAt) {
		var ret string
		return ret
	}
	return *o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VacanciesStandardVacancyFields) GetCreatedAtOk() (*string, bool) {
	if o == nil || IsNil(o.CreatedAt) {
		return nil, false
	}
	return o.CreatedAt, true
}

// HasCreatedAt returns a boolean if a field has been set.
func (o *VacanciesStandardVacancyFields) HasCreatedAt() bool {
	if o != nil && !IsNil(o.CreatedAt) {
		return true
	}

	return false
}

// SetCreatedAt gets a reference to the given string and assigns it to the CreatedAt field.
func (o *VacanciesStandardVacancyFields) SetCreatedAt(v string) {
	o.CreatedAt = &v
}

// GetDepartment returns the Department field value
// If the value is explicit nil, the zero value for VacancyDepartmentOutput will be returned
func (o *VacanciesStandardVacancyFields) GetDepartment() VacancyDepartmentOutput {
	if o == nil || o.Department.Get() == nil {
		var ret VacancyDepartmentOutput
		return ret
	}

	return *o.Department.Get()
}

// GetDepartmentOk returns a tuple with the Department field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *VacanciesStandardVacancyFields) GetDepartmentOk() (*VacancyDepartmentOutput, bool) {
	if o == nil {
		return nil, false
	}
	return o.Department.Get(), o.Department.IsSet()
}

// SetDepartment sets field value
func (o *VacanciesStandardVacancyFields) SetDepartment(v VacancyDepartmentOutput) {
	o.Department.Set(&v)
}

// GetEmployer returns the Employer field value
func (o *VacanciesStandardVacancyFields) GetEmployer() VacanciesEmployerPublic {
	if o == nil {
		var ret VacanciesEmployerPublic
		return ret
	}

	return o.Employer
}

// GetEmployerOk returns a tuple with the Employer field value
// and a boolean to check if the value has been set.
func (o *VacanciesStandardVacancyFields) GetEmployerOk() (*VacanciesEmployerPublic, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Employer, true
}

// SetEmployer sets field value
func (o *VacanciesStandardVacancyFields) SetEmployer(v VacanciesEmployerPublic) {
	o.Employer = v
}

// GetHasTest returns the HasTest field value
func (o *VacanciesStandardVacancyFields) GetHasTest() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.HasTest
}

// GetHasTestOk returns a tuple with the HasTest field value
// and a boolean to check if the value has been set.
func (o *VacanciesStandardVacancyFields) GetHasTestOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.HasTest, true
}

// SetHasTest sets field value
func (o *VacanciesStandardVacancyFields) SetHasTest(v bool) {
	o.HasTest = v
}

// GetId returns the Id field value
func (o *VacanciesStandardVacancyFields) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *VacanciesStandardVacancyFields) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *VacanciesStandardVacancyFields) SetId(v string) {
	o.Id = v
}

// GetInsiderInterview returns the InsiderInterview field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *VacanciesStandardVacancyFields) GetInsiderInterview() VacancyInsiderInterview {
	if o == nil || IsNil(o.InsiderInterview.Get()) {
		var ret VacancyInsiderInterview
		return ret
	}
	return *o.InsiderInterview.Get()
}

// GetInsiderInterviewOk returns a tuple with the InsiderInterview field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *VacanciesStandardVacancyFields) GetInsiderInterviewOk() (*VacancyInsiderInterview, bool) {
	if o == nil {
		return nil, false
	}
	return o.InsiderInterview.Get(), o.InsiderInterview.IsSet()
}

// HasInsiderInterview returns a boolean if a field has been set.
func (o *VacanciesStandardVacancyFields) HasInsiderInterview() bool {
	if o != nil && o.InsiderInterview.IsSet() {
		return true
	}

	return false
}

// SetInsiderInterview gets a reference to the given NullableVacancyInsiderInterview and assigns it to the InsiderInterview field.
func (o *VacanciesStandardVacancyFields) SetInsiderInterview(v VacancyInsiderInterview) {
	o.InsiderInterview.Set(&v)
}
// SetInsiderInterviewNil sets the value for InsiderInterview to be an explicit nil
func (o *VacanciesStandardVacancyFields) SetInsiderInterviewNil() {
	o.InsiderInterview.Set(nil)
}

// UnsetInsiderInterview ensures that no value is present for InsiderInterview, not even an explicit nil
func (o *VacanciesStandardVacancyFields) UnsetInsiderInterview() {
	o.InsiderInterview.Unset()
}

// GetMetroStations returns the MetroStations field value if set, zero value otherwise.
func (o *VacanciesStandardVacancyFields) GetMetroStations() IncludesMetroStation {
	if o == nil || IsNil(o.MetroStations) {
		var ret IncludesMetroStation
		return ret
	}
	return *o.MetroStations
}

// GetMetroStationsOk returns a tuple with the MetroStations field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VacanciesStandardVacancyFields) GetMetroStationsOk() (*IncludesMetroStation, bool) {
	if o == nil || IsNil(o.MetroStations) {
		return nil, false
	}
	return o.MetroStations, true
}

// HasMetroStations returns a boolean if a field has been set.
func (o *VacanciesStandardVacancyFields) HasMetroStations() bool {
	if o != nil && !IsNil(o.MetroStations) {
		return true
	}

	return false
}

// SetMetroStations gets a reference to the given IncludesMetroStation and assigns it to the MetroStations field.
func (o *VacanciesStandardVacancyFields) SetMetroStations(v IncludesMetroStation) {
	o.MetroStations = &v
}

// GetName returns the Name field value
func (o *VacanciesStandardVacancyFields) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *VacanciesStandardVacancyFields) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *VacanciesStandardVacancyFields) SetName(v string) {
	o.Name = v
}

// GetPremium returns the Premium field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *VacanciesStandardVacancyFields) GetPremium() bool {
	if o == nil || IsNil(o.Premium.Get()) {
		var ret bool
		return ret
	}
	return *o.Premium.Get()
}

// GetPremiumOk returns a tuple with the Premium field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *VacanciesStandardVacancyFields) GetPremiumOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return o.Premium.Get(), o.Premium.IsSet()
}

// HasPremium returns a boolean if a field has been set.
func (o *VacanciesStandardVacancyFields) HasPremium() bool {
	if o != nil && o.Premium.IsSet() {
		return true
	}

	return false
}

// SetPremium gets a reference to the given NullableBool and assigns it to the Premium field.
func (o *VacanciesStandardVacancyFields) SetPremium(v bool) {
	o.Premium.Set(&v)
}
// SetPremiumNil sets the value for Premium to be an explicit nil
func (o *VacanciesStandardVacancyFields) SetPremiumNil() {
	o.Premium.Set(nil)
}

// UnsetPremium ensures that no value is present for Premium, not even an explicit nil
func (o *VacanciesStandardVacancyFields) UnsetPremium() {
	o.Premium.Unset()
}

// GetProfessionalRoles returns the ProfessionalRoles field value
func (o *VacanciesStandardVacancyFields) GetProfessionalRoles() []VacancyProfessionalRoleItemOutput {
	if o == nil {
		var ret []VacancyProfessionalRoleItemOutput
		return ret
	}

	return o.ProfessionalRoles
}

// GetProfessionalRolesOk returns a tuple with the ProfessionalRoles field value
// and a boolean to check if the value has been set.
func (o *VacanciesStandardVacancyFields) GetProfessionalRolesOk() ([]VacancyProfessionalRoleItemOutput, bool) {
	if o == nil {
		return nil, false
	}
	return o.ProfessionalRoles, true
}

// SetProfessionalRoles sets field value
func (o *VacanciesStandardVacancyFields) SetProfessionalRoles(v []VacancyProfessionalRoleItemOutput) {
	o.ProfessionalRoles = v
}

// GetPublishedAt returns the PublishedAt field value
func (o *VacanciesStandardVacancyFields) GetPublishedAt() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.PublishedAt
}

// GetPublishedAtOk returns a tuple with the PublishedAt field value
// and a boolean to check if the value has been set.
func (o *VacanciesStandardVacancyFields) GetPublishedAtOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.PublishedAt, true
}

// SetPublishedAt sets field value
func (o *VacanciesStandardVacancyFields) SetPublishedAt(v string) {
	o.PublishedAt = v
}

// GetRelations returns the Relations field value
func (o *VacanciesStandardVacancyFields) GetRelations() []VacancyRelationItem {
	if o == nil {
		var ret []VacancyRelationItem
		return ret
	}

	return o.Relations
}

// GetRelationsOk returns a tuple with the Relations field value
// and a boolean to check if the value has been set.
func (o *VacanciesStandardVacancyFields) GetRelationsOk() ([]VacancyRelationItem, bool) {
	if o == nil {
		return nil, false
	}
	return o.Relations, true
}

// SetRelations sets field value
func (o *VacanciesStandardVacancyFields) SetRelations(v []VacancyRelationItem) {
	o.Relations = v
}

// GetResponseLetterRequired returns the ResponseLetterRequired field value
func (o *VacanciesStandardVacancyFields) GetResponseLetterRequired() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.ResponseLetterRequired
}

// GetResponseLetterRequiredOk returns a tuple with the ResponseLetterRequired field value
// and a boolean to check if the value has been set.
func (o *VacanciesStandardVacancyFields) GetResponseLetterRequiredOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ResponseLetterRequired, true
}

// SetResponseLetterRequired sets field value
func (o *VacanciesStandardVacancyFields) SetResponseLetterRequired(v bool) {
	o.ResponseLetterRequired = v
}

// GetResponseUrl returns the ResponseUrl field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *VacanciesStandardVacancyFields) GetResponseUrl() string {
	if o == nil || IsNil(o.ResponseUrl.Get()) {
		var ret string
		return ret
	}
	return *o.ResponseUrl.Get()
}

// GetResponseUrlOk returns a tuple with the ResponseUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *VacanciesStandardVacancyFields) GetResponseUrlOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.ResponseUrl.Get(), o.ResponseUrl.IsSet()
}

// HasResponseUrl returns a boolean if a field has been set.
func (o *VacanciesStandardVacancyFields) HasResponseUrl() bool {
	if o != nil && o.ResponseUrl.IsSet() {
		return true
	}

	return false
}

// SetResponseUrl gets a reference to the given NullableString and assigns it to the ResponseUrl field.
func (o *VacanciesStandardVacancyFields) SetResponseUrl(v string) {
	o.ResponseUrl.Set(&v)
}
// SetResponseUrlNil sets the value for ResponseUrl to be an explicit nil
func (o *VacanciesStandardVacancyFields) SetResponseUrlNil() {
	o.ResponseUrl.Set(nil)
}

// UnsetResponseUrl ensures that no value is present for ResponseUrl, not even an explicit nil
func (o *VacanciesStandardVacancyFields) UnsetResponseUrl() {
	o.ResponseUrl.Unset()
}

// GetSalary returns the Salary field value
// If the value is explicit nil, the zero value for VacancySalary will be returned
func (o *VacanciesStandardVacancyFields) GetSalary() VacancySalary {
	if o == nil || o.Salary.Get() == nil {
		var ret VacancySalary
		return ret
	}

	return *o.Salary.Get()
}

// GetSalaryOk returns a tuple with the Salary field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *VacanciesStandardVacancyFields) GetSalaryOk() (*VacancySalary, bool) {
	if o == nil {
		return nil, false
	}
	return o.Salary.Get(), o.Salary.IsSet()
}

// SetSalary sets field value
func (o *VacanciesStandardVacancyFields) SetSalary(v VacancySalary) {
	o.Salary.Set(&v)
}

// GetSchedule returns the Schedule field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *VacanciesStandardVacancyFields) GetSchedule() VacancyScheduleOutput {
	if o == nil || IsNil(o.Schedule.Get()) {
		var ret VacancyScheduleOutput
		return ret
	}
	return *o.Schedule.Get()
}

// GetScheduleOk returns a tuple with the Schedule field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *VacanciesStandardVacancyFields) GetScheduleOk() (*VacancyScheduleOutput, bool) {
	if o == nil {
		return nil, false
	}
	return o.Schedule.Get(), o.Schedule.IsSet()
}

// HasSchedule returns a boolean if a field has been set.
func (o *VacanciesStandardVacancyFields) HasSchedule() bool {
	if o != nil && o.Schedule.IsSet() {
		return true
	}

	return false
}

// SetSchedule gets a reference to the given NullableVacancyScheduleOutput and assigns it to the Schedule field.
func (o *VacanciesStandardVacancyFields) SetSchedule(v VacancyScheduleOutput) {
	o.Schedule.Set(&v)
}
// SetScheduleNil sets the value for Schedule to be an explicit nil
func (o *VacanciesStandardVacancyFields) SetScheduleNil() {
	o.Schedule.Set(nil)
}

// UnsetSchedule ensures that no value is present for Schedule, not even an explicit nil
func (o *VacanciesStandardVacancyFields) UnsetSchedule() {
	o.Schedule.Unset()
}

// GetSortPointDistance returns the SortPointDistance field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *VacanciesStandardVacancyFields) GetSortPointDistance() float32 {
	if o == nil || IsNil(o.SortPointDistance.Get()) {
		var ret float32
		return ret
	}
	return *o.SortPointDistance.Get()
}

// GetSortPointDistanceOk returns a tuple with the SortPointDistance field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *VacanciesStandardVacancyFields) GetSortPointDistanceOk() (*float32, bool) {
	if o == nil {
		return nil, false
	}
	return o.SortPointDistance.Get(), o.SortPointDistance.IsSet()
}

// HasSortPointDistance returns a boolean if a field has been set.
func (o *VacanciesStandardVacancyFields) HasSortPointDistance() bool {
	if o != nil && o.SortPointDistance.IsSet() {
		return true
	}

	return false
}

// SetSortPointDistance gets a reference to the given NullableFloat32 and assigns it to the SortPointDistance field.
func (o *VacanciesStandardVacancyFields) SetSortPointDistance(v float32) {
	o.SortPointDistance.Set(&v)
}
// SetSortPointDistanceNil sets the value for SortPointDistance to be an explicit nil
func (o *VacanciesStandardVacancyFields) SetSortPointDistanceNil() {
	o.SortPointDistance.Set(nil)
}

// UnsetSortPointDistance ensures that no value is present for SortPointDistance, not even an explicit nil
func (o *VacanciesStandardVacancyFields) UnsetSortPointDistance() {
	o.SortPointDistance.Unset()
}

// GetType returns the Type field value
func (o *VacanciesStandardVacancyFields) GetType() VacancyTypeOutput {
	if o == nil {
		var ret VacancyTypeOutput
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *VacanciesStandardVacancyFields) GetTypeOk() (*VacancyTypeOutput, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *VacanciesStandardVacancyFields) SetType(v VacancyTypeOutput) {
	o.Type = v
}

// GetUrl returns the Url field value
func (o *VacanciesStandardVacancyFields) GetUrl() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Url
}

// GetUrlOk returns a tuple with the Url field value
// and a boolean to check if the value has been set.
func (o *VacanciesStandardVacancyFields) GetUrlOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Url, true
}

// SetUrl sets field value
func (o *VacanciesStandardVacancyFields) SetUrl(v string) {
	o.Url = v
}

// GetWorkingDays returns the WorkingDays field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *VacanciesStandardVacancyFields) GetWorkingDays() []VacancyWorkingDayItemOutput {
	if o == nil {
		var ret []VacancyWorkingDayItemOutput
		return ret
	}
	return o.WorkingDays
}

// GetWorkingDaysOk returns a tuple with the WorkingDays field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *VacanciesStandardVacancyFields) GetWorkingDaysOk() ([]VacancyWorkingDayItemOutput, bool) {
	if o == nil || IsNil(o.WorkingDays) {
		return nil, false
	}
	return o.WorkingDays, true
}

// HasWorkingDays returns a boolean if a field has been set.
func (o *VacanciesStandardVacancyFields) HasWorkingDays() bool {
	if o != nil && IsNil(o.WorkingDays) {
		return true
	}

	return false
}

// SetWorkingDays gets a reference to the given []VacancyWorkingDayItemOutput and assigns it to the WorkingDays field.
func (o *VacanciesStandardVacancyFields) SetWorkingDays(v []VacancyWorkingDayItemOutput) {
	o.WorkingDays = v
}

// GetWorkingTimeIntervals returns the WorkingTimeIntervals field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *VacanciesStandardVacancyFields) GetWorkingTimeIntervals() []VacancyWorkingTimeIntervalItemOutput {
	if o == nil {
		var ret []VacancyWorkingTimeIntervalItemOutput
		return ret
	}
	return o.WorkingTimeIntervals
}

// GetWorkingTimeIntervalsOk returns a tuple with the WorkingTimeIntervals field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *VacanciesStandardVacancyFields) GetWorkingTimeIntervalsOk() ([]VacancyWorkingTimeIntervalItemOutput, bool) {
	if o == nil || IsNil(o.WorkingTimeIntervals) {
		return nil, false
	}
	return o.WorkingTimeIntervals, true
}

// HasWorkingTimeIntervals returns a boolean if a field has been set.
func (o *VacanciesStandardVacancyFields) HasWorkingTimeIntervals() bool {
	if o != nil && IsNil(o.WorkingTimeIntervals) {
		return true
	}

	return false
}

// SetWorkingTimeIntervals gets a reference to the given []VacancyWorkingTimeIntervalItemOutput and assigns it to the WorkingTimeIntervals field.
func (o *VacanciesStandardVacancyFields) SetWorkingTimeIntervals(v []VacancyWorkingTimeIntervalItemOutput) {
	o.WorkingTimeIntervals = v
}

// GetWorkingTimeModes returns the WorkingTimeModes field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *VacanciesStandardVacancyFields) GetWorkingTimeModes() []VacancyWorkingTimeModeItemOutput {
	if o == nil {
		var ret []VacancyWorkingTimeModeItemOutput
		return ret
	}
	return o.WorkingTimeModes
}

// GetWorkingTimeModesOk returns a tuple with the WorkingTimeModes field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *VacanciesStandardVacancyFields) GetWorkingTimeModesOk() ([]VacancyWorkingTimeModeItemOutput, bool) {
	if o == nil || IsNil(o.WorkingTimeModes) {
		return nil, false
	}
	return o.WorkingTimeModes, true
}

// HasWorkingTimeModes returns a boolean if a field has been set.
func (o *VacanciesStandardVacancyFields) HasWorkingTimeModes() bool {
	if o != nil && IsNil(o.WorkingTimeModes) {
		return true
	}

	return false
}

// SetWorkingTimeModes gets a reference to the given []VacancyWorkingTimeModeItemOutput and assigns it to the WorkingTimeModes field.
func (o *VacanciesStandardVacancyFields) SetWorkingTimeModes(v []VacancyWorkingTimeModeItemOutput) {
	o.WorkingTimeModes = v
}

func (o VacanciesStandardVacancyFields) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o VacanciesStandardVacancyFields) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["accept_incomplete_resumes"] = o.AcceptIncompleteResumes
	if o.AcceptTemporary.IsSet() {
		toSerialize["accept_temporary"] = o.AcceptTemporary.Get()
	}
	if o.Address.IsSet() {
		toSerialize["address"] = o.Address.Get()
	}
	if o.AdvResponseUrl.IsSet() {
		toSerialize["adv_response_url"] = o.AdvResponseUrl.Get()
	}
	toSerialize["alternate_url"] = o.AlternateUrl
	toSerialize["apply_alternate_url"] = o.ApplyAlternateUrl
	if o.Archived.IsSet() {
		toSerialize["archived"] = o.Archived.Get()
	}
	toSerialize["area"] = o.Area
	if o.Contacts.IsSet() {
		toSerialize["contacts"] = o.Contacts.Get()
	}
	if !IsNil(o.CreatedAt) {
		toSerialize["created_at"] = o.CreatedAt
	}
	toSerialize["department"] = o.Department.Get()
	toSerialize["employer"] = o.Employer
	toSerialize["has_test"] = o.HasTest
	toSerialize["id"] = o.Id
	if o.InsiderInterview.IsSet() {
		toSerialize["insider_interview"] = o.InsiderInterview.Get()
	}
	if !IsNil(o.MetroStations) {
		toSerialize["metro_stations"] = o.MetroStations
	}
	toSerialize["name"] = o.Name
	if o.Premium.IsSet() {
		toSerialize["premium"] = o.Premium.Get()
	}
	toSerialize["professional_roles"] = o.ProfessionalRoles
	toSerialize["published_at"] = o.PublishedAt
	toSerialize["relations"] = o.Relations
	toSerialize["response_letter_required"] = o.ResponseLetterRequired
	if o.ResponseUrl.IsSet() {
		toSerialize["response_url"] = o.ResponseUrl.Get()
	}
	toSerialize["salary"] = o.Salary.Get()
	if o.Schedule.IsSet() {
		toSerialize["schedule"] = o.Schedule.Get()
	}
	if o.SortPointDistance.IsSet() {
		toSerialize["sort_point_distance"] = o.SortPointDistance.Get()
	}
	toSerialize["type"] = o.Type
	toSerialize["url"] = o.Url
	if o.WorkingDays != nil {
		toSerialize["working_days"] = o.WorkingDays
	}
	if o.WorkingTimeIntervals != nil {
		toSerialize["working_time_intervals"] = o.WorkingTimeIntervals
	}
	if o.WorkingTimeModes != nil {
		toSerialize["working_time_modes"] = o.WorkingTimeModes
	}
	return toSerialize, nil
}

func (o *VacanciesStandardVacancyFields) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"accept_incomplete_resumes",
		"alternate_url",
		"apply_alternate_url",
		"area",
		"department",
		"employer",
		"has_test",
		"id",
		"name",
		"professional_roles",
		"published_at",
		"relations",
		"response_letter_required",
		"salary",
		"type",
		"url",
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

	varVacanciesStandardVacancyFields := _VacanciesStandardVacancyFields{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varVacanciesStandardVacancyFields)

	if err != nil {
		return err
	}

	*o = VacanciesStandardVacancyFields(varVacanciesStandardVacancyFields)

	return err
}

type NullableVacanciesStandardVacancyFields struct {
	value *VacanciesStandardVacancyFields
	isSet bool
}

func (v NullableVacanciesStandardVacancyFields) Get() *VacanciesStandardVacancyFields {
	return v.value
}

func (v *NullableVacanciesStandardVacancyFields) Set(val *VacanciesStandardVacancyFields) {
	v.value = val
	v.isSet = true
}

func (v NullableVacanciesStandardVacancyFields) IsSet() bool {
	return v.isSet
}

func (v *NullableVacanciesStandardVacancyFields) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableVacanciesStandardVacancyFields(val *VacanciesStandardVacancyFields) *NullableVacanciesStandardVacancyFields {
	return &NullableVacanciesStandardVacancyFields{value: val, isSet: true}
}

func (v NullableVacanciesStandardVacancyFields) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableVacanciesStandardVacancyFields) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


