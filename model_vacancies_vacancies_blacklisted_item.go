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

// checks if the VacanciesVacanciesBlacklistedItem type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &VacanciesVacanciesBlacklistedItem{}

// VacanciesVacanciesBlacklistedItem struct for VacanciesVacanciesBlacklistedItem
type VacanciesVacanciesBlacklistedItem struct {
	Address NullableVacancyAddress `json:"address,omitempty"`
	// URL для регистрации нажатия кнопки отклика
	AdvResponseUrl *string `json:"adv_response_url,omitempty"`
	// Ссылка на представление вакансии на сайте
	AlternateUrl string `json:"alternate_url"`
	// Ссылка на отклик на вакансию на сайте
	ApplyAlternateUrl string `json:"apply_alternate_url"`
	// Находится ли данная вакансия в архиве
	Archived bool `json:"archived"`
	Area IncludesArea `json:"area"`
	// Дата и время публикации вакансии
	CreatedAt *string `json:"created_at,omitempty"`
	Department NullableVacancyDepartment `json:"department"`
	Employer VacanciesEmployerPublic `json:"employer"`
	// Информация о наличии прикрепленного тестового задании к вакансии
	HasTest bool `json:"has_test"`
	// Идентификатор вакансии
	Id string `json:"id"`
	InsiderInterview NullableVacancyInsiderInterview `json:"insider_interview,omitempty"`
	// Название
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
	// Расстояние в метрах между центром сортировки (заданной параметрами `sort_point_lat`, `sort_point_lng`) и указанным в вакансии адресом. В случае, если в адресе указаны только станции метро, выдается расстояние между центром сортировки и средней геометрической точкой указанных станций.  Значение `sort_point_distance` выдается только в случае, если заданы параметры `sort_point_lat`, `sort_point_lng`, `order_by=distance`
	SortPointDistance NullableFloat32 `json:"sort_point_distance,omitempty"`
	Type VacancyTypeOutput `json:"type"`
	// URL вакансии
	Url string `json:"url"`
}

type _VacanciesVacanciesBlacklistedItem VacanciesVacanciesBlacklistedItem

// NewVacanciesVacanciesBlacklistedItem instantiates a new VacanciesVacanciesBlacklistedItem object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewVacanciesVacanciesBlacklistedItem(alternateUrl string, applyAlternateUrl string, archived bool, area IncludesArea, department NullableVacancyDepartment, employer VacanciesEmployerPublic, hasTest bool, id string, name string, premium bool, publishedAt string, relations []VacancyRelationItem, responseLetterRequired bool, salary NullableVacancySalary, type_ VacancyTypeOutput, url string) *VacanciesVacanciesBlacklistedItem {
	this := VacanciesVacanciesBlacklistedItem{}
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

// NewVacanciesVacanciesBlacklistedItemWithDefaults instantiates a new VacanciesVacanciesBlacklistedItem object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewVacanciesVacanciesBlacklistedItemWithDefaults() *VacanciesVacanciesBlacklistedItem {
	this := VacanciesVacanciesBlacklistedItem{}
	return &this
}

// GetAddress returns the Address field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *VacanciesVacanciesBlacklistedItem) GetAddress() VacancyAddress {
	if o == nil || IsNil(o.Address.Get()) {
		var ret VacancyAddress
		return ret
	}
	return *o.Address.Get()
}

// GetAddressOk returns a tuple with the Address field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *VacanciesVacanciesBlacklistedItem) GetAddressOk() (*VacancyAddress, bool) {
	if o == nil {
		return nil, false
	}
	return o.Address.Get(), o.Address.IsSet()
}

// HasAddress returns a boolean if a field has been set.
func (o *VacanciesVacanciesBlacklistedItem) HasAddress() bool {
	if o != nil && o.Address.IsSet() {
		return true
	}

	return false
}

// SetAddress gets a reference to the given NullableVacancyAddress and assigns it to the Address field.
func (o *VacanciesVacanciesBlacklistedItem) SetAddress(v VacancyAddress) {
	o.Address.Set(&v)
}
// SetAddressNil sets the value for Address to be an explicit nil
func (o *VacanciesVacanciesBlacklistedItem) SetAddressNil() {
	o.Address.Set(nil)
}

// UnsetAddress ensures that no value is present for Address, not even an explicit nil
func (o *VacanciesVacanciesBlacklistedItem) UnsetAddress() {
	o.Address.Unset()
}

// GetAdvResponseUrl returns the AdvResponseUrl field value if set, zero value otherwise.
func (o *VacanciesVacanciesBlacklistedItem) GetAdvResponseUrl() string {
	if o == nil || IsNil(o.AdvResponseUrl) {
		var ret string
		return ret
	}
	return *o.AdvResponseUrl
}

// GetAdvResponseUrlOk returns a tuple with the AdvResponseUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VacanciesVacanciesBlacklistedItem) GetAdvResponseUrlOk() (*string, bool) {
	if o == nil || IsNil(o.AdvResponseUrl) {
		return nil, false
	}
	return o.AdvResponseUrl, true
}

// HasAdvResponseUrl returns a boolean if a field has been set.
func (o *VacanciesVacanciesBlacklistedItem) HasAdvResponseUrl() bool {
	if o != nil && !IsNil(o.AdvResponseUrl) {
		return true
	}

	return false
}

// SetAdvResponseUrl gets a reference to the given string and assigns it to the AdvResponseUrl field.
func (o *VacanciesVacanciesBlacklistedItem) SetAdvResponseUrl(v string) {
	o.AdvResponseUrl = &v
}

// GetAlternateUrl returns the AlternateUrl field value
func (o *VacanciesVacanciesBlacklistedItem) GetAlternateUrl() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.AlternateUrl
}

// GetAlternateUrlOk returns a tuple with the AlternateUrl field value
// and a boolean to check if the value has been set.
func (o *VacanciesVacanciesBlacklistedItem) GetAlternateUrlOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AlternateUrl, true
}

// SetAlternateUrl sets field value
func (o *VacanciesVacanciesBlacklistedItem) SetAlternateUrl(v string) {
	o.AlternateUrl = v
}

// GetApplyAlternateUrl returns the ApplyAlternateUrl field value
func (o *VacanciesVacanciesBlacklistedItem) GetApplyAlternateUrl() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ApplyAlternateUrl
}

// GetApplyAlternateUrlOk returns a tuple with the ApplyAlternateUrl field value
// and a boolean to check if the value has been set.
func (o *VacanciesVacanciesBlacklistedItem) GetApplyAlternateUrlOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ApplyAlternateUrl, true
}

// SetApplyAlternateUrl sets field value
func (o *VacanciesVacanciesBlacklistedItem) SetApplyAlternateUrl(v string) {
	o.ApplyAlternateUrl = v
}

// GetArchived returns the Archived field value
func (o *VacanciesVacanciesBlacklistedItem) GetArchived() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.Archived
}

// GetArchivedOk returns a tuple with the Archived field value
// and a boolean to check if the value has been set.
func (o *VacanciesVacanciesBlacklistedItem) GetArchivedOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Archived, true
}

// SetArchived sets field value
func (o *VacanciesVacanciesBlacklistedItem) SetArchived(v bool) {
	o.Archived = v
}

// GetArea returns the Area field value
func (o *VacanciesVacanciesBlacklistedItem) GetArea() IncludesArea {
	if o == nil {
		var ret IncludesArea
		return ret
	}

	return o.Area
}

// GetAreaOk returns a tuple with the Area field value
// and a boolean to check if the value has been set.
func (o *VacanciesVacanciesBlacklistedItem) GetAreaOk() (IncludesArea, bool) {
	if o == nil {
		return IncludesArea{}, false
	}
	return o.Area, true
}

// SetArea sets field value
func (o *VacanciesVacanciesBlacklistedItem) SetArea(v IncludesArea) {
	o.Area = v
}

// GetCreatedAt returns the CreatedAt field value if set, zero value otherwise.
func (o *VacanciesVacanciesBlacklistedItem) GetCreatedAt() string {
	if o == nil || IsNil(o.CreatedAt) {
		var ret string
		return ret
	}
	return *o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VacanciesVacanciesBlacklistedItem) GetCreatedAtOk() (*string, bool) {
	if o == nil || IsNil(o.CreatedAt) {
		return nil, false
	}
	return o.CreatedAt, true
}

// HasCreatedAt returns a boolean if a field has been set.
func (o *VacanciesVacanciesBlacklistedItem) HasCreatedAt() bool {
	if o != nil && !IsNil(o.CreatedAt) {
		return true
	}

	return false
}

// SetCreatedAt gets a reference to the given string and assigns it to the CreatedAt field.
func (o *VacanciesVacanciesBlacklistedItem) SetCreatedAt(v string) {
	o.CreatedAt = &v
}

// GetDepartment returns the Department field value
// If the value is explicit nil, the zero value for VacancyDepartment will be returned
func (o *VacanciesVacanciesBlacklistedItem) GetDepartment() VacancyDepartment {
	if o == nil || o.Department.Get() == nil {
		var ret VacancyDepartment
		return ret
	}

	return *o.Department.Get()
}

// GetDepartmentOk returns a tuple with the Department field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *VacanciesVacanciesBlacklistedItem) GetDepartmentOk() (*VacancyDepartment, bool) {
	if o == nil {
		return nil, false
	}
	return o.Department.Get(), o.Department.IsSet()
}

// SetDepartment sets field value
func (o *VacanciesVacanciesBlacklistedItem) SetDepartment(v VacancyDepartment) {
	o.Department.Set(&v)
}

// GetEmployer returns the Employer field value
func (o *VacanciesVacanciesBlacklistedItem) GetEmployer() VacanciesEmployerPublic {
	if o == nil {
		var ret VacanciesEmployerPublic
		return ret
	}

	return o.Employer
}

// GetEmployerOk returns a tuple with the Employer field value
// and a boolean to check if the value has been set.
func (o *VacanciesVacanciesBlacklistedItem) GetEmployerOk() (*VacanciesEmployerPublic, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Employer, true
}

// SetEmployer sets field value
func (o *VacanciesVacanciesBlacklistedItem) SetEmployer(v VacanciesEmployerPublic) {
	o.Employer = v
}

// GetHasTest returns the HasTest field value
func (o *VacanciesVacanciesBlacklistedItem) GetHasTest() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.HasTest
}

// GetHasTestOk returns a tuple with the HasTest field value
// and a boolean to check if the value has been set.
func (o *VacanciesVacanciesBlacklistedItem) GetHasTestOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.HasTest, true
}

// SetHasTest sets field value
func (o *VacanciesVacanciesBlacklistedItem) SetHasTest(v bool) {
	o.HasTest = v
}

// GetId returns the Id field value
func (o *VacanciesVacanciesBlacklistedItem) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *VacanciesVacanciesBlacklistedItem) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *VacanciesVacanciesBlacklistedItem) SetId(v string) {
	o.Id = v
}

// GetInsiderInterview returns the InsiderInterview field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *VacanciesVacanciesBlacklistedItem) GetInsiderInterview() VacancyInsiderInterview {
	if o == nil || IsNil(o.InsiderInterview.Get()) {
		var ret VacancyInsiderInterview
		return ret
	}
	return *o.InsiderInterview.Get()
}

// GetInsiderInterviewOk returns a tuple with the InsiderInterview field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *VacanciesVacanciesBlacklistedItem) GetInsiderInterviewOk() (*VacancyInsiderInterview, bool) {
	if o == nil {
		return nil, false
	}
	return o.InsiderInterview.Get(), o.InsiderInterview.IsSet()
}

// HasInsiderInterview returns a boolean if a field has been set.
func (o *VacanciesVacanciesBlacklistedItem) HasInsiderInterview() bool {
	if o != nil && o.InsiderInterview.IsSet() {
		return true
	}

	return false
}

// SetInsiderInterview gets a reference to the given NullableVacancyInsiderInterview and assigns it to the InsiderInterview field.
func (o *VacanciesVacanciesBlacklistedItem) SetInsiderInterview(v VacancyInsiderInterview) {
	o.InsiderInterview.Set(&v)
}
// SetInsiderInterviewNil sets the value for InsiderInterview to be an explicit nil
func (o *VacanciesVacanciesBlacklistedItem) SetInsiderInterviewNil() {
	o.InsiderInterview.Set(nil)
}

// UnsetInsiderInterview ensures that no value is present for InsiderInterview, not even an explicit nil
func (o *VacanciesVacanciesBlacklistedItem) UnsetInsiderInterview() {
	o.InsiderInterview.Unset()
}

// GetName returns the Name field value
func (o *VacanciesVacanciesBlacklistedItem) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *VacanciesVacanciesBlacklistedItem) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *VacanciesVacanciesBlacklistedItem) SetName(v string) {
	o.Name = v
}

// GetPremium returns the Premium field value
func (o *VacanciesVacanciesBlacklistedItem) GetPremium() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.Premium
}

// GetPremiumOk returns a tuple with the Premium field value
// and a boolean to check if the value has been set.
func (o *VacanciesVacanciesBlacklistedItem) GetPremiumOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Premium, true
}

// SetPremium sets field value
func (o *VacanciesVacanciesBlacklistedItem) SetPremium(v bool) {
	o.Premium = v
}

// GetPublishedAt returns the PublishedAt field value
func (o *VacanciesVacanciesBlacklistedItem) GetPublishedAt() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.PublishedAt
}

// GetPublishedAtOk returns a tuple with the PublishedAt field value
// and a boolean to check if the value has been set.
func (o *VacanciesVacanciesBlacklistedItem) GetPublishedAtOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.PublishedAt, true
}

// SetPublishedAt sets field value
func (o *VacanciesVacanciesBlacklistedItem) SetPublishedAt(v string) {
	o.PublishedAt = v
}

// GetRelations returns the Relations field value
func (o *VacanciesVacanciesBlacklistedItem) GetRelations() []VacancyRelationItem {
	if o == nil {
		var ret []VacancyRelationItem
		return ret
	}

	return o.Relations
}

// GetRelationsOk returns a tuple with the Relations field value
// and a boolean to check if the value has been set.
func (o *VacanciesVacanciesBlacklistedItem) GetRelationsOk() ([]VacancyRelationItem, bool) {
	if o == nil {
		return nil, false
	}
	return o.Relations, true
}

// SetRelations sets field value
func (o *VacanciesVacanciesBlacklistedItem) SetRelations(v []VacancyRelationItem) {
	o.Relations = v
}

// GetResponseLetterRequired returns the ResponseLetterRequired field value
func (o *VacanciesVacanciesBlacklistedItem) GetResponseLetterRequired() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.ResponseLetterRequired
}

// GetResponseLetterRequiredOk returns a tuple with the ResponseLetterRequired field value
// and a boolean to check if the value has been set.
func (o *VacanciesVacanciesBlacklistedItem) GetResponseLetterRequiredOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ResponseLetterRequired, true
}

// SetResponseLetterRequired sets field value
func (o *VacanciesVacanciesBlacklistedItem) SetResponseLetterRequired(v bool) {
	o.ResponseLetterRequired = v
}

// GetResponseUrl returns the ResponseUrl field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *VacanciesVacanciesBlacklistedItem) GetResponseUrl() string {
	if o == nil || IsNil(o.ResponseUrl.Get()) {
		var ret string
		return ret
	}
	return *o.ResponseUrl.Get()
}

// GetResponseUrlOk returns a tuple with the ResponseUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *VacanciesVacanciesBlacklistedItem) GetResponseUrlOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.ResponseUrl.Get(), o.ResponseUrl.IsSet()
}

// HasResponseUrl returns a boolean if a field has been set.
func (o *VacanciesVacanciesBlacklistedItem) HasResponseUrl() bool {
	if o != nil && o.ResponseUrl.IsSet() {
		return true
	}

	return false
}

// SetResponseUrl gets a reference to the given NullableString and assigns it to the ResponseUrl field.
func (o *VacanciesVacanciesBlacklistedItem) SetResponseUrl(v string) {
	o.ResponseUrl.Set(&v)
}
// SetResponseUrlNil sets the value for ResponseUrl to be an explicit nil
func (o *VacanciesVacanciesBlacklistedItem) SetResponseUrlNil() {
	o.ResponseUrl.Set(nil)
}

// UnsetResponseUrl ensures that no value is present for ResponseUrl, not even an explicit nil
func (o *VacanciesVacanciesBlacklistedItem) UnsetResponseUrl() {
	o.ResponseUrl.Unset()
}

// GetSalary returns the Salary field value
// If the value is explicit nil, the zero value for VacancySalary will be returned
func (o *VacanciesVacanciesBlacklistedItem) GetSalary() VacancySalary {
	if o == nil || o.Salary.Get() == nil {
		var ret VacancySalary
		return ret
	}

	return *o.Salary.Get()
}

// GetSalaryOk returns a tuple with the Salary field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *VacanciesVacanciesBlacklistedItem) GetSalaryOk() (*VacancySalary, bool) {
	if o == nil {
		return nil, false
	}
	return o.Salary.Get(), o.Salary.IsSet()
}

// SetSalary sets field value
func (o *VacanciesVacanciesBlacklistedItem) SetSalary(v VacancySalary) {
	o.Salary.Set(&v)
}

// GetSortPointDistance returns the SortPointDistance field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *VacanciesVacanciesBlacklistedItem) GetSortPointDistance() float32 {
	if o == nil || IsNil(o.SortPointDistance.Get()) {
		var ret float32
		return ret
	}
	return *o.SortPointDistance.Get()
}

// GetSortPointDistanceOk returns a tuple with the SortPointDistance field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *VacanciesVacanciesBlacklistedItem) GetSortPointDistanceOk() (*float32, bool) {
	if o == nil {
		return nil, false
	}
	return o.SortPointDistance.Get(), o.SortPointDistance.IsSet()
}

// HasSortPointDistance returns a boolean if a field has been set.
func (o *VacanciesVacanciesBlacklistedItem) HasSortPointDistance() bool {
	if o != nil && o.SortPointDistance.IsSet() {
		return true
	}

	return false
}

// SetSortPointDistance gets a reference to the given NullableFloat32 and assigns it to the SortPointDistance field.
func (o *VacanciesVacanciesBlacklistedItem) SetSortPointDistance(v float32) {
	o.SortPointDistance.Set(&v)
}
// SetSortPointDistanceNil sets the value for SortPointDistance to be an explicit nil
func (o *VacanciesVacanciesBlacklistedItem) SetSortPointDistanceNil() {
	o.SortPointDistance.Set(nil)
}

// UnsetSortPointDistance ensures that no value is present for SortPointDistance, not even an explicit nil
func (o *VacanciesVacanciesBlacklistedItem) UnsetSortPointDistance() {
	o.SortPointDistance.Unset()
}

// GetType returns the Type field value
func (o *VacanciesVacanciesBlacklistedItem) GetType() VacancyTypeOutput {
	if o == nil {
		var ret VacancyTypeOutput
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *VacanciesVacanciesBlacklistedItem) GetTypeOk() (*VacancyTypeOutput, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *VacanciesVacanciesBlacklistedItem) SetType(v VacancyTypeOutput) {
	o.Type = v
}

// GetUrl returns the Url field value
func (o *VacanciesVacanciesBlacklistedItem) GetUrl() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Url
}

// GetUrlOk returns a tuple with the Url field value
// and a boolean to check if the value has been set.
func (o *VacanciesVacanciesBlacklistedItem) GetUrlOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Url, true
}

// SetUrl sets field value
func (o *VacanciesVacanciesBlacklistedItem) SetUrl(v string) {
	o.Url = v
}

func (o VacanciesVacanciesBlacklistedItem) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o VacanciesVacanciesBlacklistedItem) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.Address.IsSet() {
		toSerialize["address"] = o.Address.Get()
	}
	if !IsNil(o.AdvResponseUrl) {
		toSerialize["adv_response_url"] = o.AdvResponseUrl
	}
	toSerialize["alternate_url"] = o.AlternateUrl
	toSerialize["apply_alternate_url"] = o.ApplyAlternateUrl
	toSerialize["archived"] = o.Archived
	toSerialize["area"] = o.Area
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
	toSerialize["name"] = o.Name
	toSerialize["premium"] = o.Premium
	toSerialize["published_at"] = o.PublishedAt
	toSerialize["relations"] = o.Relations
	toSerialize["response_letter_required"] = o.ResponseLetterRequired
	if o.ResponseUrl.IsSet() {
		toSerialize["response_url"] = o.ResponseUrl.Get()
	}
	toSerialize["salary"] = o.Salary.Get()
	if o.SortPointDistance.IsSet() {
		toSerialize["sort_point_distance"] = o.SortPointDistance.Get()
	}
	toSerialize["type"] = o.Type
	toSerialize["url"] = o.Url
	return toSerialize, nil
}

func (o *VacanciesVacanciesBlacklistedItem) UnmarshalJSON(data []byte) (err error) {
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

	varVacanciesVacanciesBlacklistedItem := _VacanciesVacanciesBlacklistedItem{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varVacanciesVacanciesBlacklistedItem)

	if err != nil {
		return err
	}

	*o = VacanciesVacanciesBlacklistedItem(varVacanciesVacanciesBlacklistedItem)

	return err
}

type NullableVacanciesVacanciesBlacklistedItem struct {
	value *VacanciesVacanciesBlacklistedItem
	isSet bool
}

func (v NullableVacanciesVacanciesBlacklistedItem) Get() *VacanciesVacanciesBlacklistedItem {
	return v.value
}

func (v *NullableVacanciesVacanciesBlacklistedItem) Set(val *VacanciesVacanciesBlacklistedItem) {
	v.value = val
	v.isSet = true
}

func (v NullableVacanciesVacanciesBlacklistedItem) IsSet() bool {
	return v.isSet
}

func (v *NullableVacanciesVacanciesBlacklistedItem) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableVacanciesVacanciesBlacklistedItem(val *VacanciesVacanciesBlacklistedItem) *NullableVacanciesVacanciesBlacklistedItem {
	return &NullableVacanciesVacanciesBlacklistedItem{value: val, isSet: true}
}

func (v NullableVacanciesVacanciesBlacklistedItem) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableVacanciesVacanciesBlacklistedItem) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


