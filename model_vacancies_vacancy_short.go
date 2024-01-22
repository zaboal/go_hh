/*
HeadHunter API

По-русски | [Switch to English](https://api.hh.ru/openapi/en/redoc)  В OpenAPI ведется пока что только небольшая часть документации [Основная документация](https://github.com/hhru/api).  Для поиска по документации можно использовать Ctrl+F.  # Авторизация  API поддерживает следующие уровни авторизации:   - [авторизация приложения](#section/Avtorizaciya/Avtorizaciya-prilozheniya)   - [авторизация пользователя](#section/Avtorizaciya/Avtorizaciya-polzovatelya)  * [Авторизация пользователя](#section/Avtorizaciya/Avtorizaciya-polzovatelya)   * [Правила формирования специального redirect_uri](#section/Avtorizaciya/Pravila-formirovaniya-specialnogo-redirect_uri)   * [Процесс авторизации](#section/Avtorizaciya/Process-avtorizacii)   * [Успешное получение временного `authorization_code`](#get-authorization_code)   * [Получение access и refresh токенов](#section/Avtorizaciya/Poluchenie-access-i-refresh-tokenov) * [Обновление пары access и refresh токенов](#section/Avtorizaciya/Obnovlenie-pary-access-i-refresh-tokenov) * [Инвалидация токена](#tag/Avtorizaciya-rabotodatelya/operation/invalidate-token) * [Запрос авторизации под другим пользователем](#section/Avtorizaciya/Zapros-avtorizacii-pod-drugim-polzovatelem) * [Авторизация под разными рабочими аккаунтами](#section/Avtorizaciya/Avtorizaciya-pod-raznymi-rabochimi-akkauntami) * [Авторизация приложения](#section/Avtorizaciya/Avtorizaciya-prilozheniya)       ## Авторизация пользователя Для выполнения запросов от имени пользователя необходимо пользоваться токеном пользователя.  В начале приложению необходимо направить пользователя (открыть страницу) по адресу:  ``` https://hh.ru/oauth/authorize? response_type=code& client_id={client_id}& state={state}& redirect_uri={redirect_uri} ```  Обязательные параметры:  * `response_type=code` — указание на способ получения авторизации, используя `authorization code` * `client_id` — идентификатор, полученный при создании приложения   Необязательные параметры:  * `state` — в случае указания, будет включен в ответный редирект. Это позволяет исключить возможность взлома путём подделки межсайтовых запросов. Подробнее об этом: [RFC 6749. Section 10.12](http://tools.ietf.org/html/rfc6749#section-10.12) * `redirect_uri` — uri для перенаправления пользователя после авторизации. Если не указать, используется из настроек приложения. При наличии происходит валидация значения. Вероятнее всего, потребуется сделать urlencode значения параметра.  ## Правила формирования специального redirect_uri  К примеру, если в настройках сохранен `http://example.com/oauth`, то разрешено указывать:  * `http://www.example.com/oauth` — поддомен; * `http://www.example.com/oauth/sub/path` — уточнение пути; * `http://example.com/oauth?lang=RU` — дополнительный параметр; * `http://www.example.com/oauth/sub/path?lang=RU` — всё вместе.  Запрещено:  * `https://example.com/oauth` — различные протоколы; * `http://wwwexample.com/oauth` — различные домены; * `http://wwwexample.com/` — другой путь; * `http://example.com/oauths` — другой путь; * `http://example.com:80/oauths` — указание изначально отсутствующего порта;  ## Процесс авторизации  Если пользователь не авторизован на сайте, ему будет показана форма авторизации на сайте. После прохождения авторизации на сайте, пользователю будет выведена форма с запросом разрешения доступа вашего приложения к его персональным данным.  Если пользователь не разрешает доступ приложению, пользователь будет перенаправлен на указанный `redirect_uri` с `?error=access_denied` и `state={state}`, если таковой был указан при первом запросе.  <a name=\"get-authorization_code\"></a> ### Успешное получение временного `authorization_code`  В случае разрешения прав, в редиректе будет указан временный `authorization_code`:  ```http HTTP/1.1 302 FOUND Location: {redirect_uri}?code={authorization_code} ```  Если пользователь авторизован на сайте и доступ данному приложению однажды ранее выдан, ответом будет сразу вышеописанный редирект с `authorization_code` (без показа формы логина и выдачи прав).  ## Получение access и refresh токенов  После получения `authorization_code` приложению необходимо сделать сервер-сервер запрос `POST https://hh.ru/oauth/token` для обмена полученного `authorization_code` на `access_token`.  В теле запроса необходимо передать [дополнительные параметры](#required_parameters).  Тело запроса необходимо передавать в стандартном `application/x-www-form-urlencoded` с указанием соответствующего заголовка `Content-Type`.  `authorization_code` имеет довольно короткий срок жизни, при его истечении необходимо запросить новый.  ## Обновление пары access и refresh токенов `access_token` также имеет срок жизни (ключ `expires_in`, в секундах), при его истечении приложение должно сделать запрос с `refresh_token` для получения нового.  Запрос необходимо делать в `application/x-www-form-urlencoded`.  ``` POST https://hh.ru/oauth/token ```  В теле запроса необходимо передать [дополнительные параметры](#required_parameters)  `refresh_token` можно использовать только один раз и только по истечению срока действия `access_token`.  После получения новой пары access и refresh токенов, их необходимо использовать в дальнейших запросах в api и запросах на продление токена.  ## Запрос авторизации под другим пользователем  Возможен следующий сценарий:  1. Приложение перенаправляет пользователя на сайт с запросом авторизации. 2. Пользователь на сайте уже авторизован и данному приложение доступ уже был разрешен. 3. Пользователю будет предложена возможность продолжить работу под текущим аккаунтом, либо зайти под другим аккаунтом.  Если есть необходимость, чтобы на шаге 3 сразу происходило перенаправление (redirect) с временным токеном, необходимо добавить к запросу `/oauth/authorize...` параметр `skip_choose_account=true`. В этом случае автоматически выдаётся доступ пользователю авторизованному на сайте.  Если есть необходимость всегда показывать форму авторизации, приложение может добавить к запросу `/oauth/authorize...` параметр `force_login=true`. В этом случае, пользователю будет показана форма авторизации с логином и паролем даже в случае, если пользователь уже авторизован.  Это может быть полезно приложениям, которые предоставляют сервис только для соискателей. Если пришел пользователь-работодатель, приложение может предложить пользователю повторно разрешить доступ на сайте, уже указав другую учетную запись.  Также, после авторизации приложение может показать пользователю сообщение:  ``` Вы вошли как %Имя_Фамилия%. Это не вы? ``` и предоставить ссылку с `force_login=true` для возможности захода под другим логином.  ## Авторизация под разными рабочими аккаунтами  Для получения списка рабочих аккаунтов менеджера и для работы под разными рабочими аккаунтами менеджера необходимо прочитать документацию по [рабочим аккаунтам менеджера](#tag/Menedzhery-rabotodatelya/operation/get-manager-accounts)  ## Авторизация приложения  Токен приложения необходимо сгенерировать 1 раз. В случае, если токен был скомпрометирован, его нужно запросить еще раз. При этом ранее выданный токен отзывается. Владелец приложения может посмотреть актуальный `access_token` для приложения на сайте [https://dev.hh.ru/admin](https://dev.hh.ru/admin). В случае, если вы еще ни разу [не получали токен приложения](#section/Avtorizaciya/Avtorizaciya-prilozheniya), токен отображаться не будет.  <a name=\"get-client-token\"></a> ### Получение токена приложения Для получения `access_token` необходимо сделать запрос:  ``` POST https://hh.ru/oauth/token ```  В теле запроса необходимо передать [дополнительные параметры](#required_parameters). Тело запроса необходимо передавать в стандартном `application/x-www-form-urlencoded` с указанием соответствующего заголовка `Content-Type`.  Данный `access_token` имеет **неограниченный** срок жизни. При повторном запросе ранее выданный токен отзывается и выдается новый. Запрашивать `access_token` можно не чаще, чем один раз в 5 минут.  В случае компрометации токена необходимо [инвалидировать](#tag/Avtorizaciya-rabotodatelya/operation/invalidate-token) скомпроментированный токен и запросить токен заново!  <!-- ReDoc-Inject: <security-definitions> --> 

API version: 1.0.0
Contact: api@hh.ru
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package github.com/zaboal/hh-go

import (
	"encoding/json"
	"bytes"
	"fmt"
)

// checks if the VacanciesVacancyShort type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &VacanciesVacancyShort{}

// VacanciesVacancyShort struct for VacanciesVacancyShort
type VacanciesVacancyShort struct {
	Address NullableVacancyAddressRawOutput `json:"address,omitempty"`
	// Ссылка на представление вакансии на сайте
	AlternateUrl string `json:"alternate_url"`
	// Ссылка на отклик на вакансию на сайте
	ApplyAlternateUrl string `json:"apply_alternate_url"`
	// Находится ли данная вакансия в архиве
	Archived bool `json:"archived"`
	Area IncludesArea `json:"area"`
	Department NullableVacancyDepartmentOutput `json:"department"`
	Employer NullableVacanciesEmployerPublic `json:"employer"`
	// Информация о наличии прикрепленного тестового задании к вакансии
	HasTest bool `json:"has_test"`
	// Идентификатор вакансии
	Id string `json:"id"`
	// Название вакансии
	Name string `json:"name"`
	// Является ли данная вакансия премиум-вакансией
	Premium bool `json:"premium"`
	// Дата и время публикации вакансии
	PublishedAt string `json:"published_at"`
	// Возвращает связи соискателя с вакансией. Значения из поля `vacancy_relation` в [справочнике полей](#tag/Obshie-spravochniki/operation/get-dictionaries)
	Relations []VacancyRelationItem `json:"relations"`
	// Обязательно ли заполнять сообщение при отклике на вакансию
	ResponseLetterRequired bool `json:"response_letter_required"`
	// URL отклика для прямых вакансий (`type.id=direct`)
	ResponseUrl NullableString `json:"response_url,omitempty"`
	Salary NullableVacancySalary `json:"salary"`
	// Отображать ли лого для вакансии в поисковой выдаче
	ShowLogoInSearch NullableBool `json:"show_logo_in_search,omitempty"`
	Type VacancyTypeOutput `json:"type"`
	// URL вакансии
	Url string `json:"url"`
}

type _VacanciesVacancyShort VacanciesVacancyShort

// NewVacanciesVacancyShort instantiates a new VacanciesVacancyShort object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewVacanciesVacancyShort(alternateUrl string, applyAlternateUrl string, archived bool, area IncludesArea, department NullableVacancyDepartmentOutput, employer NullableVacanciesEmployerPublic, hasTest bool, id string, name string, premium bool, publishedAt string, relations []VacancyRelationItem, responseLetterRequired bool, salary NullableVacancySalary, type_ VacancyTypeOutput, url string) *VacanciesVacancyShort {
	this := VacanciesVacancyShort{}
	this.AlternateUrl = alternateUrl
	this.ApplyAlternateUrl = applyAlternateUrl
	this.Archived = archived
	this.Area = area
	this.Department = department
	this.Employer = employer
	this.HasTest = hasTest
	this.Id = id
	this.Name = name
	this.Premium = premium
	this.PublishedAt = publishedAt
	this.Relations = relations
	this.ResponseLetterRequired = responseLetterRequired
	this.Salary = salary
	this.Type = type_
	this.Url = url
	return &this
}

// NewVacanciesVacancyShortWithDefaults instantiates a new VacanciesVacancyShort object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewVacanciesVacancyShortWithDefaults() *VacanciesVacancyShort {
	this := VacanciesVacancyShort{}
	return &this
}

// GetAddress returns the Address field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *VacanciesVacancyShort) GetAddress() VacancyAddressRawOutput {
	if o == nil || IsNil(o.Address.Get()) {
		var ret VacancyAddressRawOutput
		return ret
	}
	return *o.Address.Get()
}

// GetAddressOk returns a tuple with the Address field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *VacanciesVacancyShort) GetAddressOk() (*VacancyAddressRawOutput, bool) {
	if o == nil {
		return nil, false
	}
	return o.Address.Get(), o.Address.IsSet()
}

// HasAddress returns a boolean if a field has been set.
func (o *VacanciesVacancyShort) HasAddress() bool {
	if o != nil && o.Address.IsSet() {
		return true
	}

	return false
}

// SetAddress gets a reference to the given NullableVacancyAddressRawOutput and assigns it to the Address field.
func (o *VacanciesVacancyShort) SetAddress(v VacancyAddressRawOutput) {
	o.Address.Set(&v)
}
// SetAddressNil sets the value for Address to be an explicit nil
func (o *VacanciesVacancyShort) SetAddressNil() {
	o.Address.Set(nil)
}

// UnsetAddress ensures that no value is present for Address, not even an explicit nil
func (o *VacanciesVacancyShort) UnsetAddress() {
	o.Address.Unset()
}

// GetAlternateUrl returns the AlternateUrl field value
func (o *VacanciesVacancyShort) GetAlternateUrl() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.AlternateUrl
}

// GetAlternateUrlOk returns a tuple with the AlternateUrl field value
// and a boolean to check if the value has been set.
func (o *VacanciesVacancyShort) GetAlternateUrlOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AlternateUrl, true
}

// SetAlternateUrl sets field value
func (o *VacanciesVacancyShort) SetAlternateUrl(v string) {
	o.AlternateUrl = v
}

// GetApplyAlternateUrl returns the ApplyAlternateUrl field value
func (o *VacanciesVacancyShort) GetApplyAlternateUrl() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ApplyAlternateUrl
}

// GetApplyAlternateUrlOk returns a tuple with the ApplyAlternateUrl field value
// and a boolean to check if the value has been set.
func (o *VacanciesVacancyShort) GetApplyAlternateUrlOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ApplyAlternateUrl, true
}

// SetApplyAlternateUrl sets field value
func (o *VacanciesVacancyShort) SetApplyAlternateUrl(v string) {
	o.ApplyAlternateUrl = v
}

// GetArchived returns the Archived field value
func (o *VacanciesVacancyShort) GetArchived() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.Archived
}

// GetArchivedOk returns a tuple with the Archived field value
// and a boolean to check if the value has been set.
func (o *VacanciesVacancyShort) GetArchivedOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Archived, true
}

// SetArchived sets field value
func (o *VacanciesVacancyShort) SetArchived(v bool) {
	o.Archived = v
}

// GetArea returns the Area field value
func (o *VacanciesVacancyShort) GetArea() IncludesArea {
	if o == nil {
		var ret IncludesArea
		return ret
	}

	return o.Area
}

// GetAreaOk returns a tuple with the Area field value
// and a boolean to check if the value has been set.
func (o *VacanciesVacancyShort) GetAreaOk() (*IncludesArea, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Area, true
}

// SetArea sets field value
func (o *VacanciesVacancyShort) SetArea(v IncludesArea) {
	o.Area = v
}

// GetDepartment returns the Department field value
// If the value is explicit nil, the zero value for VacancyDepartmentOutput will be returned
func (o *VacanciesVacancyShort) GetDepartment() VacancyDepartmentOutput {
	if o == nil || o.Department.Get() == nil {
		var ret VacancyDepartmentOutput
		return ret
	}

	return *o.Department.Get()
}

// GetDepartmentOk returns a tuple with the Department field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *VacanciesVacancyShort) GetDepartmentOk() (*VacancyDepartmentOutput, bool) {
	if o == nil {
		return nil, false
	}
	return o.Department.Get(), o.Department.IsSet()
}

// SetDepartment sets field value
func (o *VacanciesVacancyShort) SetDepartment(v VacancyDepartmentOutput) {
	o.Department.Set(&v)
}

// GetEmployer returns the Employer field value
// If the value is explicit nil, the zero value for VacanciesEmployerPublic will be returned
func (o *VacanciesVacancyShort) GetEmployer() VacanciesEmployerPublic {
	if o == nil || o.Employer.Get() == nil {
		var ret VacanciesEmployerPublic
		return ret
	}

	return *o.Employer.Get()
}

// GetEmployerOk returns a tuple with the Employer field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *VacanciesVacancyShort) GetEmployerOk() (*VacanciesEmployerPublic, bool) {
	if o == nil {
		return nil, false
	}
	return o.Employer.Get(), o.Employer.IsSet()
}

// SetEmployer sets field value
func (o *VacanciesVacancyShort) SetEmployer(v VacanciesEmployerPublic) {
	o.Employer.Set(&v)
}

// GetHasTest returns the HasTest field value
func (o *VacanciesVacancyShort) GetHasTest() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.HasTest
}

// GetHasTestOk returns a tuple with the HasTest field value
// and a boolean to check if the value has been set.
func (o *VacanciesVacancyShort) GetHasTestOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.HasTest, true
}

// SetHasTest sets field value
func (o *VacanciesVacancyShort) SetHasTest(v bool) {
	o.HasTest = v
}

// GetId returns the Id field value
func (o *VacanciesVacancyShort) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *VacanciesVacancyShort) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *VacanciesVacancyShort) SetId(v string) {
	o.Id = v
}

// GetName returns the Name field value
func (o *VacanciesVacancyShort) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *VacanciesVacancyShort) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *VacanciesVacancyShort) SetName(v string) {
	o.Name = v
}

// GetPremium returns the Premium field value
func (o *VacanciesVacancyShort) GetPremium() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.Premium
}

// GetPremiumOk returns a tuple with the Premium field value
// and a boolean to check if the value has been set.
func (o *VacanciesVacancyShort) GetPremiumOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Premium, true
}

// SetPremium sets field value
func (o *VacanciesVacancyShort) SetPremium(v bool) {
	o.Premium = v
}

// GetPublishedAt returns the PublishedAt field value
func (o *VacanciesVacancyShort) GetPublishedAt() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.PublishedAt
}

// GetPublishedAtOk returns a tuple with the PublishedAt field value
// and a boolean to check if the value has been set.
func (o *VacanciesVacancyShort) GetPublishedAtOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.PublishedAt, true
}

// SetPublishedAt sets field value
func (o *VacanciesVacancyShort) SetPublishedAt(v string) {
	o.PublishedAt = v
}

// GetRelations returns the Relations field value
func (o *VacanciesVacancyShort) GetRelations() []VacancyRelationItem {
	if o == nil {
		var ret []VacancyRelationItem
		return ret
	}

	return o.Relations
}

// GetRelationsOk returns a tuple with the Relations field value
// and a boolean to check if the value has been set.
func (o *VacanciesVacancyShort) GetRelationsOk() ([]VacancyRelationItem, bool) {
	if o == nil {
		return nil, false
	}
	return o.Relations, true
}

// SetRelations sets field value
func (o *VacanciesVacancyShort) SetRelations(v []VacancyRelationItem) {
	o.Relations = v
}

// GetResponseLetterRequired returns the ResponseLetterRequired field value
func (o *VacanciesVacancyShort) GetResponseLetterRequired() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.ResponseLetterRequired
}

// GetResponseLetterRequiredOk returns a tuple with the ResponseLetterRequired field value
// and a boolean to check if the value has been set.
func (o *VacanciesVacancyShort) GetResponseLetterRequiredOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ResponseLetterRequired, true
}

// SetResponseLetterRequired sets field value
func (o *VacanciesVacancyShort) SetResponseLetterRequired(v bool) {
	o.ResponseLetterRequired = v
}

// GetResponseUrl returns the ResponseUrl field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *VacanciesVacancyShort) GetResponseUrl() string {
	if o == nil || IsNil(o.ResponseUrl.Get()) {
		var ret string
		return ret
	}
	return *o.ResponseUrl.Get()
}

// GetResponseUrlOk returns a tuple with the ResponseUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *VacanciesVacancyShort) GetResponseUrlOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.ResponseUrl.Get(), o.ResponseUrl.IsSet()
}

// HasResponseUrl returns a boolean if a field has been set.
func (o *VacanciesVacancyShort) HasResponseUrl() bool {
	if o != nil && o.ResponseUrl.IsSet() {
		return true
	}

	return false
}

// SetResponseUrl gets a reference to the given NullableString and assigns it to the ResponseUrl field.
func (o *VacanciesVacancyShort) SetResponseUrl(v string) {
	o.ResponseUrl.Set(&v)
}
// SetResponseUrlNil sets the value for ResponseUrl to be an explicit nil
func (o *VacanciesVacancyShort) SetResponseUrlNil() {
	o.ResponseUrl.Set(nil)
}

// UnsetResponseUrl ensures that no value is present for ResponseUrl, not even an explicit nil
func (o *VacanciesVacancyShort) UnsetResponseUrl() {
	o.ResponseUrl.Unset()
}

// GetSalary returns the Salary field value
// If the value is explicit nil, the zero value for VacancySalary will be returned
func (o *VacanciesVacancyShort) GetSalary() VacancySalary {
	if o == nil || o.Salary.Get() == nil {
		var ret VacancySalary
		return ret
	}

	return *o.Salary.Get()
}

// GetSalaryOk returns a tuple with the Salary field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *VacanciesVacancyShort) GetSalaryOk() (*VacancySalary, bool) {
	if o == nil {
		return nil, false
	}
	return o.Salary.Get(), o.Salary.IsSet()
}

// SetSalary sets field value
func (o *VacanciesVacancyShort) SetSalary(v VacancySalary) {
	o.Salary.Set(&v)
}

// GetShowLogoInSearch returns the ShowLogoInSearch field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *VacanciesVacancyShort) GetShowLogoInSearch() bool {
	if o == nil || IsNil(o.ShowLogoInSearch.Get()) {
		var ret bool
		return ret
	}
	return *o.ShowLogoInSearch.Get()
}

// GetShowLogoInSearchOk returns a tuple with the ShowLogoInSearch field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *VacanciesVacancyShort) GetShowLogoInSearchOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return o.ShowLogoInSearch.Get(), o.ShowLogoInSearch.IsSet()
}

// HasShowLogoInSearch returns a boolean if a field has been set.
func (o *VacanciesVacancyShort) HasShowLogoInSearch() bool {
	if o != nil && o.ShowLogoInSearch.IsSet() {
		return true
	}

	return false
}

// SetShowLogoInSearch gets a reference to the given NullableBool and assigns it to the ShowLogoInSearch field.
func (o *VacanciesVacancyShort) SetShowLogoInSearch(v bool) {
	o.ShowLogoInSearch.Set(&v)
}
// SetShowLogoInSearchNil sets the value for ShowLogoInSearch to be an explicit nil
func (o *VacanciesVacancyShort) SetShowLogoInSearchNil() {
	o.ShowLogoInSearch.Set(nil)
}

// UnsetShowLogoInSearch ensures that no value is present for ShowLogoInSearch, not even an explicit nil
func (o *VacanciesVacancyShort) UnsetShowLogoInSearch() {
	o.ShowLogoInSearch.Unset()
}

// GetType returns the Type field value
func (o *VacanciesVacancyShort) GetType() VacancyTypeOutput {
	if o == nil {
		var ret VacancyTypeOutput
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *VacanciesVacancyShort) GetTypeOk() (*VacancyTypeOutput, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *VacanciesVacancyShort) SetType(v VacancyTypeOutput) {
	o.Type = v
}

// GetUrl returns the Url field value
func (o *VacanciesVacancyShort) GetUrl() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Url
}

// GetUrlOk returns a tuple with the Url field value
// and a boolean to check if the value has been set.
func (o *VacanciesVacancyShort) GetUrlOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Url, true
}

// SetUrl sets field value
func (o *VacanciesVacancyShort) SetUrl(v string) {
	o.Url = v
}

func (o VacanciesVacancyShort) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o VacanciesVacancyShort) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.Address.IsSet() {
		toSerialize["address"] = o.Address.Get()
	}
	toSerialize["alternate_url"] = o.AlternateUrl
	toSerialize["apply_alternate_url"] = o.ApplyAlternateUrl
	toSerialize["archived"] = o.Archived
	toSerialize["area"] = o.Area
	toSerialize["department"] = o.Department.Get()
	toSerialize["employer"] = o.Employer.Get()
	toSerialize["has_test"] = o.HasTest
	toSerialize["id"] = o.Id
	toSerialize["name"] = o.Name
	toSerialize["premium"] = o.Premium
	toSerialize["published_at"] = o.PublishedAt
	toSerialize["relations"] = o.Relations
	toSerialize["response_letter_required"] = o.ResponseLetterRequired
	if o.ResponseUrl.IsSet() {
		toSerialize["response_url"] = o.ResponseUrl.Get()
	}
	toSerialize["salary"] = o.Salary.Get()
	if o.ShowLogoInSearch.IsSet() {
		toSerialize["show_logo_in_search"] = o.ShowLogoInSearch.Get()
	}
	toSerialize["type"] = o.Type
	toSerialize["url"] = o.Url
	return toSerialize, nil
}

func (o *VacanciesVacancyShort) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"alternate_url",
		"apply_alternate_url",
		"archived",
		"area",
		"department",
		"employer",
		"has_test",
		"id",
		"name",
		"premium",
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

	varVacanciesVacancyShort := _VacanciesVacancyShort{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varVacanciesVacancyShort)

	if err != nil {
		return err
	}

	*o = VacanciesVacancyShort(varVacanciesVacancyShort)

	return err
}

type NullableVacanciesVacancyShort struct {
	value *VacanciesVacancyShort
	isSet bool
}

func (v NullableVacanciesVacancyShort) Get() *VacanciesVacancyShort {
	return v.value
}

func (v *NullableVacanciesVacancyShort) Set(val *VacanciesVacancyShort) {
	v.value = val
	v.isSet = true
}

func (v NullableVacanciesVacancyShort) IsSet() bool {
	return v.isSet
}

func (v *NullableVacanciesVacancyShort) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableVacanciesVacancyShort(val *VacanciesVacancyShort) *NullableVacanciesVacancyShort {
	return &NullableVacanciesVacancyShort{value: val, isSet: true}
}

func (v NullableVacanciesVacancyShort) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableVacanciesVacancyShort) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


