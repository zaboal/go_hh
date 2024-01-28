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

// checks if the ResumesMineItem type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ResumesMineItem{}

// ResumesMineItem struct for ResumesMineItem
type ResumesMineItem struct {
	// Дополнительные действия
	Actions ResumeObjectsActionsForOwner `json:"actions"`
	// Возраст
	Age NullableFloat32 `json:"age,omitempty"`
	// URL резюме на сайте
	AlternateUrl string `json:"alternate_url"`
	Area *IncludesIdNameUrl `json:"area,omitempty"`
	AutoHideTime NullableIncludesIdName `json:"auto_hide_time,omitempty"`
	// Доступен ли просмотр контактной информации в резюме текущему работодателю
	CanViewFullInfo NullableBool `json:"can_view_full_info,omitempty"`
	// Список сертификатов соискателя
	Certificate []ResumeObjectsCertificate `json:"certificate"`
	// Дата и время создания резюме
	CreatedAt string `json:"created_at"`
	// Ссылки для скачивания резюме в разных форматах
	Download ResumeObjectsDownload `json:"download"`
	// Образование соискателя. 

Особенности сохранения образования:

* Если передать и высшее и среднее образование и уровень образования "средний", то сохранится только среднее образование.
* Если передать и высшее и среднее образование и уровень образования "высшее", то сохранится только высшее образование

	Education ResumeObjectsEducation `json:"education"`
	// Опыт работы
	Experience []ResumeObjectsExperienceForOwner `json:"experience"`
	// Имя
	FirstName NullableString `json:"first_name,omitempty"`
	Gender *IncludesIdName `json:"gender,omitempty"`
	// Документация [Список скрытых полей](https://github.com/hhru/api/blob/master/docs/employer_resumes.md#hidden-fields). Возможные значения элементов приведены в поле `resume_hidden_fields` [справочника полей](#tag/Obshie-spravochniki/operation/get-dictionaries)
	HiddenFields []IncludesIdName `json:"hidden_fields"`
	// Идентификатор резюме
	Id string `json:"id"`
	// Фамилия
	LastName NullableString `json:"last_name,omitempty"`
	// Выделено ли резюме в поиске
	Marked bool `json:"marked"`
	// Отчество
	MiddleName NullableString `json:"middle_name,omitempty"`
	Photo *ProfilePhoto `json:"photo,omitempty"`
	// Ресурс, на котором было размещено резюме
	Platform *IncludesId `json:"platform,omitempty"`
	Salary NullableResumeObjectsSalaryProperties `json:"salary,omitempty"`
	// Желаемая должность
	Title NullableString `json:"title,omitempty"`
	TotalExperience NullableResumeObjectsTotalExperience `json:"total_experience,omitempty"`
	// Дата и время обновления резюме
	UpdatedAt string `json:"updated_at"`
	// URL резюме на сайте
	Url string `json:"url"`
	// Заблокировано ли резюме ([подробнее](#tag/Rezyume.-Prosmotr-informacii/Status-rezyume))
	Blocked bool `json:"blocked"`
	// Можно ли опубликовать или обновить данное резюме
	CanPublishOrUpdate NullableBool `json:"can_publish_or_update,omitempty"`
	// Заполнено ли резюме
	Finished bool `json:"finished"`
	Status IncludesIdName `json:"status"`
	Access ResumeObjectsAccess `json:"access"`
	// Число новых просмотров. Данный счетчик сбрасывается при получении [детальной истории просмотров](#tag/Rezyume.-Prosmotr-informacii/operation/get-resume-view-history) 
	NewViews float32 `json:"new_views"`
	// Дата и время следующей возможной публикации/обновления. Для неопубликованных резюме возвращается `null`
	NextPublishAt NullableString `json:"next_publish_at,omitempty"`
	// Платные услуги по резюме для автора
	PaidServices []ResumeObjectsPaidServices `json:"paid_services"`
	// Число просмотров резюме
	TotalViews float32 `json:"total_views"`
	// URL, по которому необходимо сделать GET-запрос для получения [детальной истории просмотров](#tag/Rezyume.-Prosmotr-informacii/operation/get-resume-view-history) 
	ViewsUrl string `json:"views_url"`
	// Список контактов соискателя
	Contact []IncludesContact `json:"contact"`
	// Дата и время создания резюме
	// Deprecated
	Created string `json:"created"`
	SimilarVacancies ResumeObjectsSimilarVacancies `json:"similar_vacancies"`
	// Дата и время обновления резюме
	// Deprecated
	Updated string `json:"updated"`
	// Видно ли резюме в поиске
	// Deprecated
	Visible bool `json:"visible"`
}

type _ResumesMineItem ResumesMineItem

// NewResumesMineItem instantiates a new ResumesMineItem object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewResumesMineItem(actions ResumeObjectsActionsForOwner, alternateUrl string, certificate []ResumeObjectsCertificate, createdAt string, download ResumeObjectsDownload, education ResumeObjectsEducation, experience []ResumeObjectsExperienceForOwner, hiddenFields []IncludesIdName, id string, marked bool, updatedAt string, url string, blocked bool, finished bool, status IncludesIdName, access ResumeObjectsAccess, newViews float32, paidServices []ResumeObjectsPaidServices, totalViews float32, viewsUrl string, contact []IncludesContact, created string, similarVacancies ResumeObjectsSimilarVacancies, updated string, visible bool) *ResumesMineItem {
	this := ResumesMineItem{}
	this.Actions = actions
	this.AlternateUrl = alternateUrl
	this.Certificate = certificate
	this.CreatedAt = createdAt
	this.Download = download
	this.Education = education
	this.Experience = experience
	this.HiddenFields = hiddenFields
	this.Id = id
	this.Marked = marked
	this.UpdatedAt = updatedAt
	this.Url = url
	this.Blocked = blocked
	this.Finished = finished
	this.Status = status
	this.Access = access
	this.NewViews = newViews
	this.PaidServices = paidServices
	this.TotalViews = totalViews
	this.ViewsUrl = viewsUrl
	this.Contact = contact
	this.Created = created
	this.SimilarVacancies = similarVacancies
	this.Updated = updated
	this.Visible = visible
	return &this
}

// NewResumesMineItemWithDefaults instantiates a new ResumesMineItem object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewResumesMineItemWithDefaults() *ResumesMineItem {
	this := ResumesMineItem{}
	return &this
}

// GetActions returns the Actions field value
func (o *ResumesMineItem) GetActions() ResumeObjectsActionsForOwner {
	if o == nil {
		var ret ResumeObjectsActionsForOwner
		return ret
	}

	return o.Actions
}

// GetActionsOk returns a tuple with the Actions field value
// and a boolean to check if the value has been set.
func (o *ResumesMineItem) GetActionsOk() (*ResumeObjectsActionsForOwner, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Actions, true
}

// SetActions sets field value
func (o *ResumesMineItem) SetActions(v ResumeObjectsActionsForOwner) {
	o.Actions = v
}

// GetAge returns the Age field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ResumesMineItem) GetAge() float32 {
	if o == nil || IsNil(o.Age.Get()) {
		var ret float32
		return ret
	}
	return *o.Age.Get()
}

// GetAgeOk returns a tuple with the Age field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ResumesMineItem) GetAgeOk() (*float32, bool) {
	if o == nil {
		return nil, false
	}
	return o.Age.Get(), o.Age.IsSet()
}

// HasAge returns a boolean if a field has been set.
func (o *ResumesMineItem) HasAge() bool {
	if o != nil && o.Age.IsSet() {
		return true
	}

	return false
}

// SetAge gets a reference to the given NullableFloat32 and assigns it to the Age field.
func (o *ResumesMineItem) SetAge(v float32) {
	o.Age.Set(&v)
}
// SetAgeNil sets the value for Age to be an explicit nil
func (o *ResumesMineItem) SetAgeNil() {
	o.Age.Set(nil)
}

// UnsetAge ensures that no value is present for Age, not even an explicit nil
func (o *ResumesMineItem) UnsetAge() {
	o.Age.Unset()
}

// GetAlternateUrl returns the AlternateUrl field value
func (o *ResumesMineItem) GetAlternateUrl() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.AlternateUrl
}

// GetAlternateUrlOk returns a tuple with the AlternateUrl field value
// and a boolean to check if the value has been set.
func (o *ResumesMineItem) GetAlternateUrlOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AlternateUrl, true
}

// SetAlternateUrl sets field value
func (o *ResumesMineItem) SetAlternateUrl(v string) {
	o.AlternateUrl = v
}

// GetArea returns the Area field value if set, zero value otherwise.
func (o *ResumesMineItem) GetArea() IncludesIdNameUrl {
	if o == nil || IsNil(o.Area) {
		var ret IncludesIdNameUrl
		return ret
	}
	return *o.Area
}

// GetAreaOk returns a tuple with the Area field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ResumesMineItem) GetAreaOk() (*IncludesIdNameUrl, bool) {
	if o == nil || IsNil(o.Area) {
		return nil, false
	}
	return o.Area, true
}

// HasArea returns a boolean if a field has been set.
func (o *ResumesMineItem) HasArea() bool {
	if o != nil && !IsNil(o.Area) {
		return true
	}

	return false
}

// SetArea gets a reference to the given IncludesIdNameUrl and assigns it to the Area field.
func (o *ResumesMineItem) SetArea(v IncludesIdNameUrl) {
	o.Area = &v
}

// GetAutoHideTime returns the AutoHideTime field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ResumesMineItem) GetAutoHideTime() IncludesIdName {
	if o == nil || IsNil(o.AutoHideTime.Get()) {
		var ret IncludesIdName
		return ret
	}
	return *o.AutoHideTime.Get()
}

// GetAutoHideTimeOk returns a tuple with the AutoHideTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ResumesMineItem) GetAutoHideTimeOk() (*IncludesIdName, bool) {
	if o == nil {
		return nil, false
	}
	return o.AutoHideTime.Get(), o.AutoHideTime.IsSet()
}

// HasAutoHideTime returns a boolean if a field has been set.
func (o *ResumesMineItem) HasAutoHideTime() bool {
	if o != nil && o.AutoHideTime.IsSet() {
		return true
	}

	return false
}

// SetAutoHideTime gets a reference to the given NullableIncludesIdName and assigns it to the AutoHideTime field.
func (o *ResumesMineItem) SetAutoHideTime(v IncludesIdName) {
	o.AutoHideTime.Set(&v)
}
// SetAutoHideTimeNil sets the value for AutoHideTime to be an explicit nil
func (o *ResumesMineItem) SetAutoHideTimeNil() {
	o.AutoHideTime.Set(nil)
}

// UnsetAutoHideTime ensures that no value is present for AutoHideTime, not even an explicit nil
func (o *ResumesMineItem) UnsetAutoHideTime() {
	o.AutoHideTime.Unset()
}

// GetCanViewFullInfo returns the CanViewFullInfo field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ResumesMineItem) GetCanViewFullInfo() bool {
	if o == nil || IsNil(o.CanViewFullInfo.Get()) {
		var ret bool
		return ret
	}
	return *o.CanViewFullInfo.Get()
}

// GetCanViewFullInfoOk returns a tuple with the CanViewFullInfo field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ResumesMineItem) GetCanViewFullInfoOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return o.CanViewFullInfo.Get(), o.CanViewFullInfo.IsSet()
}

// HasCanViewFullInfo returns a boolean if a field has been set.
func (o *ResumesMineItem) HasCanViewFullInfo() bool {
	if o != nil && o.CanViewFullInfo.IsSet() {
		return true
	}

	return false
}

// SetCanViewFullInfo gets a reference to the given NullableBool and assigns it to the CanViewFullInfo field.
func (o *ResumesMineItem) SetCanViewFullInfo(v bool) {
	o.CanViewFullInfo.Set(&v)
}
// SetCanViewFullInfoNil sets the value for CanViewFullInfo to be an explicit nil
func (o *ResumesMineItem) SetCanViewFullInfoNil() {
	o.CanViewFullInfo.Set(nil)
}

// UnsetCanViewFullInfo ensures that no value is present for CanViewFullInfo, not even an explicit nil
func (o *ResumesMineItem) UnsetCanViewFullInfo() {
	o.CanViewFullInfo.Unset()
}

// GetCertificate returns the Certificate field value
func (o *ResumesMineItem) GetCertificate() []ResumeObjectsCertificate {
	if o == nil {
		var ret []ResumeObjectsCertificate
		return ret
	}

	return o.Certificate
}

// GetCertificateOk returns a tuple with the Certificate field value
// and a boolean to check if the value has been set.
func (o *ResumesMineItem) GetCertificateOk() ([]ResumeObjectsCertificate, bool) {
	if o == nil {
		return nil, false
	}
	return o.Certificate, true
}

// SetCertificate sets field value
func (o *ResumesMineItem) SetCertificate(v []ResumeObjectsCertificate) {
	o.Certificate = v
}

// GetCreatedAt returns the CreatedAt field value
func (o *ResumesMineItem) GetCreatedAt() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value
// and a boolean to check if the value has been set.
func (o *ResumesMineItem) GetCreatedAtOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CreatedAt, true
}

// SetCreatedAt sets field value
func (o *ResumesMineItem) SetCreatedAt(v string) {
	o.CreatedAt = v
}

// GetDownload returns the Download field value
func (o *ResumesMineItem) GetDownload() ResumeObjectsDownload {
	if o == nil {
		var ret ResumeObjectsDownload
		return ret
	}

	return o.Download
}

// GetDownloadOk returns a tuple with the Download field value
// and a boolean to check if the value has been set.
func (o *ResumesMineItem) GetDownloadOk() (*ResumeObjectsDownload, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Download, true
}

// SetDownload sets field value
func (o *ResumesMineItem) SetDownload(v ResumeObjectsDownload) {
	o.Download = v
}

// GetEducation returns the Education field value
func (o *ResumesMineItem) GetEducation() ResumeObjectsEducation {
	if o == nil {
		var ret ResumeObjectsEducation
		return ret
	}

	return o.Education
}

// GetEducationOk returns a tuple with the Education field value
// and a boolean to check if the value has been set.
func (o *ResumesMineItem) GetEducationOk() (*ResumeObjectsEducation, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Education, true
}

// SetEducation sets field value
func (o *ResumesMineItem) SetEducation(v ResumeObjectsEducation) {
	o.Education = v
}

// GetExperience returns the Experience field value
func (o *ResumesMineItem) GetExperience() []ResumeObjectsExperienceForOwner {
	if o == nil {
		var ret []ResumeObjectsExperienceForOwner
		return ret
	}

	return o.Experience
}

// GetExperienceOk returns a tuple with the Experience field value
// and a boolean to check if the value has been set.
func (o *ResumesMineItem) GetExperienceOk() ([]ResumeObjectsExperienceForOwner, bool) {
	if o == nil {
		return nil, false
	}
	return o.Experience, true
}

// SetExperience sets field value
func (o *ResumesMineItem) SetExperience(v []ResumeObjectsExperienceForOwner) {
	o.Experience = v
}

// GetFirstName returns the FirstName field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ResumesMineItem) GetFirstName() string {
	if o == nil || IsNil(o.FirstName.Get()) {
		var ret string
		return ret
	}
	return *o.FirstName.Get()
}

// GetFirstNameOk returns a tuple with the FirstName field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ResumesMineItem) GetFirstNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.FirstName.Get(), o.FirstName.IsSet()
}

// HasFirstName returns a boolean if a field has been set.
func (o *ResumesMineItem) HasFirstName() bool {
	if o != nil && o.FirstName.IsSet() {
		return true
	}

	return false
}

// SetFirstName gets a reference to the given NullableString and assigns it to the FirstName field.
func (o *ResumesMineItem) SetFirstName(v string) {
	o.FirstName.Set(&v)
}
// SetFirstNameNil sets the value for FirstName to be an explicit nil
func (o *ResumesMineItem) SetFirstNameNil() {
	o.FirstName.Set(nil)
}

// UnsetFirstName ensures that no value is present for FirstName, not even an explicit nil
func (o *ResumesMineItem) UnsetFirstName() {
	o.FirstName.Unset()
}

// GetGender returns the Gender field value if set, zero value otherwise.
func (o *ResumesMineItem) GetGender() IncludesIdName {
	if o == nil || IsNil(o.Gender) {
		var ret IncludesIdName
		return ret
	}
	return *o.Gender
}

// GetGenderOk returns a tuple with the Gender field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ResumesMineItem) GetGenderOk() (*IncludesIdName, bool) {
	if o == nil || IsNil(o.Gender) {
		return nil, false
	}
	return o.Gender, true
}

// HasGender returns a boolean if a field has been set.
func (o *ResumesMineItem) HasGender() bool {
	if o != nil && !IsNil(o.Gender) {
		return true
	}

	return false
}

// SetGender gets a reference to the given IncludesIdName and assigns it to the Gender field.
func (o *ResumesMineItem) SetGender(v IncludesIdName) {
	o.Gender = &v
}

// GetHiddenFields returns the HiddenFields field value
func (o *ResumesMineItem) GetHiddenFields() []IncludesIdName {
	if o == nil {
		var ret []IncludesIdName
		return ret
	}

	return o.HiddenFields
}

// GetHiddenFieldsOk returns a tuple with the HiddenFields field value
// and a boolean to check if the value has been set.
func (o *ResumesMineItem) GetHiddenFieldsOk() ([]IncludesIdName, bool) {
	if o == nil {
		return nil, false
	}
	return o.HiddenFields, true
}

// SetHiddenFields sets field value
func (o *ResumesMineItem) SetHiddenFields(v []IncludesIdName) {
	o.HiddenFields = v
}

// GetId returns the Id field value
func (o *ResumesMineItem) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *ResumesMineItem) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *ResumesMineItem) SetId(v string) {
	o.Id = v
}

// GetLastName returns the LastName field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ResumesMineItem) GetLastName() string {
	if o == nil || IsNil(o.LastName.Get()) {
		var ret string
		return ret
	}
	return *o.LastName.Get()
}

// GetLastNameOk returns a tuple with the LastName field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ResumesMineItem) GetLastNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.LastName.Get(), o.LastName.IsSet()
}

// HasLastName returns a boolean if a field has been set.
func (o *ResumesMineItem) HasLastName() bool {
	if o != nil && o.LastName.IsSet() {
		return true
	}

	return false
}

// SetLastName gets a reference to the given NullableString and assigns it to the LastName field.
func (o *ResumesMineItem) SetLastName(v string) {
	o.LastName.Set(&v)
}
// SetLastNameNil sets the value for LastName to be an explicit nil
func (o *ResumesMineItem) SetLastNameNil() {
	o.LastName.Set(nil)
}

// UnsetLastName ensures that no value is present for LastName, not even an explicit nil
func (o *ResumesMineItem) UnsetLastName() {
	o.LastName.Unset()
}

// GetMarked returns the Marked field value
func (o *ResumesMineItem) GetMarked() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.Marked
}

// GetMarkedOk returns a tuple with the Marked field value
// and a boolean to check if the value has been set.
func (o *ResumesMineItem) GetMarkedOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Marked, true
}

// SetMarked sets field value
func (o *ResumesMineItem) SetMarked(v bool) {
	o.Marked = v
}

// GetMiddleName returns the MiddleName field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ResumesMineItem) GetMiddleName() string {
	if o == nil || IsNil(o.MiddleName.Get()) {
		var ret string
		return ret
	}
	return *o.MiddleName.Get()
}

// GetMiddleNameOk returns a tuple with the MiddleName field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ResumesMineItem) GetMiddleNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.MiddleName.Get(), o.MiddleName.IsSet()
}

// HasMiddleName returns a boolean if a field has been set.
func (o *ResumesMineItem) HasMiddleName() bool {
	if o != nil && o.MiddleName.IsSet() {
		return true
	}

	return false
}

// SetMiddleName gets a reference to the given NullableString and assigns it to the MiddleName field.
func (o *ResumesMineItem) SetMiddleName(v string) {
	o.MiddleName.Set(&v)
}
// SetMiddleNameNil sets the value for MiddleName to be an explicit nil
func (o *ResumesMineItem) SetMiddleNameNil() {
	o.MiddleName.Set(nil)
}

// UnsetMiddleName ensures that no value is present for MiddleName, not even an explicit nil
func (o *ResumesMineItem) UnsetMiddleName() {
	o.MiddleName.Unset()
}

// GetPhoto returns the Photo field value if set, zero value otherwise.
func (o *ResumesMineItem) GetPhoto() ProfilePhoto {
	if o == nil || IsNil(o.Photo) {
		var ret ProfilePhoto
		return ret
	}
	return *o.Photo
}

// GetPhotoOk returns a tuple with the Photo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ResumesMineItem) GetPhotoOk() (*ProfilePhoto, bool) {
	if o == nil || IsNil(o.Photo) {
		return nil, false
	}
	return o.Photo, true
}

// HasPhoto returns a boolean if a field has been set.
func (o *ResumesMineItem) HasPhoto() bool {
	if o != nil && !IsNil(o.Photo) {
		return true
	}

	return false
}

// SetPhoto gets a reference to the given ProfilePhoto and assigns it to the Photo field.
func (o *ResumesMineItem) SetPhoto(v ProfilePhoto) {
	o.Photo = &v
}

// GetPlatform returns the Platform field value if set, zero value otherwise.
func (o *ResumesMineItem) GetPlatform() IncludesId {
	if o == nil || IsNil(o.Platform) {
		var ret IncludesId
		return ret
	}
	return *o.Platform
}

// GetPlatformOk returns a tuple with the Platform field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ResumesMineItem) GetPlatformOk() (*IncludesId, bool) {
	if o == nil || IsNil(o.Platform) {
		return nil, false
	}
	return o.Platform, true
}

// HasPlatform returns a boolean if a field has been set.
func (o *ResumesMineItem) HasPlatform() bool {
	if o != nil && !IsNil(o.Platform) {
		return true
	}

	return false
}

// SetPlatform gets a reference to the given IncludesId and assigns it to the Platform field.
func (o *ResumesMineItem) SetPlatform(v IncludesId) {
	o.Platform = &v
}

// GetSalary returns the Salary field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ResumesMineItem) GetSalary() ResumeObjectsSalaryProperties {
	if o == nil || IsNil(o.Salary.Get()) {
		var ret ResumeObjectsSalaryProperties
		return ret
	}
	return *o.Salary.Get()
}

// GetSalaryOk returns a tuple with the Salary field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ResumesMineItem) GetSalaryOk() (*ResumeObjectsSalaryProperties, bool) {
	if o == nil {
		return nil, false
	}
	return o.Salary.Get(), o.Salary.IsSet()
}

// HasSalary returns a boolean if a field has been set.
func (o *ResumesMineItem) HasSalary() bool {
	if o != nil && o.Salary.IsSet() {
		return true
	}

	return false
}

// SetSalary gets a reference to the given NullableResumeObjectsSalaryProperties and assigns it to the Salary field.
func (o *ResumesMineItem) SetSalary(v ResumeObjectsSalaryProperties) {
	o.Salary.Set(&v)
}
// SetSalaryNil sets the value for Salary to be an explicit nil
func (o *ResumesMineItem) SetSalaryNil() {
	o.Salary.Set(nil)
}

// UnsetSalary ensures that no value is present for Salary, not even an explicit nil
func (o *ResumesMineItem) UnsetSalary() {
	o.Salary.Unset()
}

// GetTitle returns the Title field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ResumesMineItem) GetTitle() string {
	if o == nil || IsNil(o.Title.Get()) {
		var ret string
		return ret
	}
	return *o.Title.Get()
}

// GetTitleOk returns a tuple with the Title field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ResumesMineItem) GetTitleOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Title.Get(), o.Title.IsSet()
}

// HasTitle returns a boolean if a field has been set.
func (o *ResumesMineItem) HasTitle() bool {
	if o != nil && o.Title.IsSet() {
		return true
	}

	return false
}

// SetTitle gets a reference to the given NullableString and assigns it to the Title field.
func (o *ResumesMineItem) SetTitle(v string) {
	o.Title.Set(&v)
}
// SetTitleNil sets the value for Title to be an explicit nil
func (o *ResumesMineItem) SetTitleNil() {
	o.Title.Set(nil)
}

// UnsetTitle ensures that no value is present for Title, not even an explicit nil
func (o *ResumesMineItem) UnsetTitle() {
	o.Title.Unset()
}

// GetTotalExperience returns the TotalExperience field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ResumesMineItem) GetTotalExperience() ResumeObjectsTotalExperience {
	if o == nil || IsNil(o.TotalExperience.Get()) {
		var ret ResumeObjectsTotalExperience
		return ret
	}
	return *o.TotalExperience.Get()
}

// GetTotalExperienceOk returns a tuple with the TotalExperience field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ResumesMineItem) GetTotalExperienceOk() (*ResumeObjectsTotalExperience, bool) {
	if o == nil {
		return nil, false
	}
	return o.TotalExperience.Get(), o.TotalExperience.IsSet()
}

// HasTotalExperience returns a boolean if a field has been set.
func (o *ResumesMineItem) HasTotalExperience() bool {
	if o != nil && o.TotalExperience.IsSet() {
		return true
	}

	return false
}

// SetTotalExperience gets a reference to the given NullableResumeObjectsTotalExperience and assigns it to the TotalExperience field.
func (o *ResumesMineItem) SetTotalExperience(v ResumeObjectsTotalExperience) {
	o.TotalExperience.Set(&v)
}
// SetTotalExperienceNil sets the value for TotalExperience to be an explicit nil
func (o *ResumesMineItem) SetTotalExperienceNil() {
	o.TotalExperience.Set(nil)
}

// UnsetTotalExperience ensures that no value is present for TotalExperience, not even an explicit nil
func (o *ResumesMineItem) UnsetTotalExperience() {
	o.TotalExperience.Unset()
}

// GetUpdatedAt returns the UpdatedAt field value
func (o *ResumesMineItem) GetUpdatedAt() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.UpdatedAt
}

// GetUpdatedAtOk returns a tuple with the UpdatedAt field value
// and a boolean to check if the value has been set.
func (o *ResumesMineItem) GetUpdatedAtOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.UpdatedAt, true
}

// SetUpdatedAt sets field value
func (o *ResumesMineItem) SetUpdatedAt(v string) {
	o.UpdatedAt = v
}

// GetUrl returns the Url field value
func (o *ResumesMineItem) GetUrl() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Url
}

// GetUrlOk returns a tuple with the Url field value
// and a boolean to check if the value has been set.
func (o *ResumesMineItem) GetUrlOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Url, true
}

// SetUrl sets field value
func (o *ResumesMineItem) SetUrl(v string) {
	o.Url = v
}

// GetBlocked returns the Blocked field value
func (o *ResumesMineItem) GetBlocked() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.Blocked
}

// GetBlockedOk returns a tuple with the Blocked field value
// and a boolean to check if the value has been set.
func (o *ResumesMineItem) GetBlockedOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Blocked, true
}

// SetBlocked sets field value
func (o *ResumesMineItem) SetBlocked(v bool) {
	o.Blocked = v
}

// GetCanPublishOrUpdate returns the CanPublishOrUpdate field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ResumesMineItem) GetCanPublishOrUpdate() bool {
	if o == nil || IsNil(o.CanPublishOrUpdate.Get()) {
		var ret bool
		return ret
	}
	return *o.CanPublishOrUpdate.Get()
}

// GetCanPublishOrUpdateOk returns a tuple with the CanPublishOrUpdate field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ResumesMineItem) GetCanPublishOrUpdateOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return o.CanPublishOrUpdate.Get(), o.CanPublishOrUpdate.IsSet()
}

// HasCanPublishOrUpdate returns a boolean if a field has been set.
func (o *ResumesMineItem) HasCanPublishOrUpdate() bool {
	if o != nil && o.CanPublishOrUpdate.IsSet() {
		return true
	}

	return false
}

// SetCanPublishOrUpdate gets a reference to the given NullableBool and assigns it to the CanPublishOrUpdate field.
func (o *ResumesMineItem) SetCanPublishOrUpdate(v bool) {
	o.CanPublishOrUpdate.Set(&v)
}
// SetCanPublishOrUpdateNil sets the value for CanPublishOrUpdate to be an explicit nil
func (o *ResumesMineItem) SetCanPublishOrUpdateNil() {
	o.CanPublishOrUpdate.Set(nil)
}

// UnsetCanPublishOrUpdate ensures that no value is present for CanPublishOrUpdate, not even an explicit nil
func (o *ResumesMineItem) UnsetCanPublishOrUpdate() {
	o.CanPublishOrUpdate.Unset()
}

// GetFinished returns the Finished field value
func (o *ResumesMineItem) GetFinished() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.Finished
}

// GetFinishedOk returns a tuple with the Finished field value
// and a boolean to check if the value has been set.
func (o *ResumesMineItem) GetFinishedOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Finished, true
}

// SetFinished sets field value
func (o *ResumesMineItem) SetFinished(v bool) {
	o.Finished = v
}

// GetStatus returns the Status field value
func (o *ResumesMineItem) GetStatus() IncludesIdName {
	if o == nil {
		var ret IncludesIdName
		return ret
	}

	return o.Status
}

// GetStatusOk returns a tuple with the Status field value
// and a boolean to check if the value has been set.
func (o *ResumesMineItem) GetStatusOk() (*IncludesIdName, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Status, true
}

// SetStatus sets field value
func (o *ResumesMineItem) SetStatus(v IncludesIdName) {
	o.Status = v
}

// GetAccess returns the Access field value
func (o *ResumesMineItem) GetAccess() ResumeObjectsAccess {
	if o == nil {
		var ret ResumeObjectsAccess
		return ret
	}

	return o.Access
}

// GetAccessOk returns a tuple with the Access field value
// and a boolean to check if the value has been set.
func (o *ResumesMineItem) GetAccessOk() (*ResumeObjectsAccess, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Access, true
}

// SetAccess sets field value
func (o *ResumesMineItem) SetAccess(v ResumeObjectsAccess) {
	o.Access = v
}

// GetNewViews returns the NewViews field value
func (o *ResumesMineItem) GetNewViews() float32 {
	if o == nil {
		var ret float32
		return ret
	}

	return o.NewViews
}

// GetNewViewsOk returns a tuple with the NewViews field value
// and a boolean to check if the value has been set.
func (o *ResumesMineItem) GetNewViewsOk() (*float32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.NewViews, true
}

// SetNewViews sets field value
func (o *ResumesMineItem) SetNewViews(v float32) {
	o.NewViews = v
}

// GetNextPublishAt returns the NextPublishAt field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ResumesMineItem) GetNextPublishAt() string {
	if o == nil || IsNil(o.NextPublishAt.Get()) {
		var ret string
		return ret
	}
	return *o.NextPublishAt.Get()
}

// GetNextPublishAtOk returns a tuple with the NextPublishAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ResumesMineItem) GetNextPublishAtOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.NextPublishAt.Get(), o.NextPublishAt.IsSet()
}

// HasNextPublishAt returns a boolean if a field has been set.
func (o *ResumesMineItem) HasNextPublishAt() bool {
	if o != nil && o.NextPublishAt.IsSet() {
		return true
	}

	return false
}

// SetNextPublishAt gets a reference to the given NullableString and assigns it to the NextPublishAt field.
func (o *ResumesMineItem) SetNextPublishAt(v string) {
	o.NextPublishAt.Set(&v)
}
// SetNextPublishAtNil sets the value for NextPublishAt to be an explicit nil
func (o *ResumesMineItem) SetNextPublishAtNil() {
	o.NextPublishAt.Set(nil)
}

// UnsetNextPublishAt ensures that no value is present for NextPublishAt, not even an explicit nil
func (o *ResumesMineItem) UnsetNextPublishAt() {
	o.NextPublishAt.Unset()
}

// GetPaidServices returns the PaidServices field value
func (o *ResumesMineItem) GetPaidServices() []ResumeObjectsPaidServices {
	if o == nil {
		var ret []ResumeObjectsPaidServices
		return ret
	}

	return o.PaidServices
}

// GetPaidServicesOk returns a tuple with the PaidServices field value
// and a boolean to check if the value has been set.
func (o *ResumesMineItem) GetPaidServicesOk() ([]ResumeObjectsPaidServices, bool) {
	if o == nil {
		return nil, false
	}
	return o.PaidServices, true
}

// SetPaidServices sets field value
func (o *ResumesMineItem) SetPaidServices(v []ResumeObjectsPaidServices) {
	o.PaidServices = v
}

// GetTotalViews returns the TotalViews field value
func (o *ResumesMineItem) GetTotalViews() float32 {
	if o == nil {
		var ret float32
		return ret
	}

	return o.TotalViews
}

// GetTotalViewsOk returns a tuple with the TotalViews field value
// and a boolean to check if the value has been set.
func (o *ResumesMineItem) GetTotalViewsOk() (*float32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.TotalViews, true
}

// SetTotalViews sets field value
func (o *ResumesMineItem) SetTotalViews(v float32) {
	o.TotalViews = v
}

// GetViewsUrl returns the ViewsUrl field value
func (o *ResumesMineItem) GetViewsUrl() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ViewsUrl
}

// GetViewsUrlOk returns a tuple with the ViewsUrl field value
// and a boolean to check if the value has been set.
func (o *ResumesMineItem) GetViewsUrlOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ViewsUrl, true
}

// SetViewsUrl sets field value
func (o *ResumesMineItem) SetViewsUrl(v string) {
	o.ViewsUrl = v
}

// GetContact returns the Contact field value
func (o *ResumesMineItem) GetContact() []IncludesContact {
	if o == nil {
		var ret []IncludesContact
		return ret
	}

	return o.Contact
}

// GetContactOk returns a tuple with the Contact field value
// and a boolean to check if the value has been set.
func (o *ResumesMineItem) GetContactOk() ([]IncludesContact, bool) {
	if o == nil {
		return nil, false
	}
	return o.Contact, true
}

// SetContact sets field value
func (o *ResumesMineItem) SetContact(v []IncludesContact) {
	o.Contact = v
}

// GetCreated returns the Created field value
// Deprecated
func (o *ResumesMineItem) GetCreated() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Created
}

// GetCreatedOk returns a tuple with the Created field value
// and a boolean to check if the value has been set.
// Deprecated
func (o *ResumesMineItem) GetCreatedOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Created, true
}

// SetCreated sets field value
// Deprecated
func (o *ResumesMineItem) SetCreated(v string) {
	o.Created = v
}

// GetSimilarVacancies returns the SimilarVacancies field value
func (o *ResumesMineItem) GetSimilarVacancies() ResumeObjectsSimilarVacancies {
	if o == nil {
		var ret ResumeObjectsSimilarVacancies
		return ret
	}

	return o.SimilarVacancies
}

// GetSimilarVacanciesOk returns a tuple with the SimilarVacancies field value
// and a boolean to check if the value has been set.
func (o *ResumesMineItem) GetSimilarVacanciesOk() (*ResumeObjectsSimilarVacancies, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SimilarVacancies, true
}

// SetSimilarVacancies sets field value
func (o *ResumesMineItem) SetSimilarVacancies(v ResumeObjectsSimilarVacancies) {
	o.SimilarVacancies = v
}

// GetUpdated returns the Updated field value
// Deprecated
func (o *ResumesMineItem) GetUpdated() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Updated
}

// GetUpdatedOk returns a tuple with the Updated field value
// and a boolean to check if the value has been set.
// Deprecated
func (o *ResumesMineItem) GetUpdatedOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Updated, true
}

// SetUpdated sets field value
// Deprecated
func (o *ResumesMineItem) SetUpdated(v string) {
	o.Updated = v
}

// GetVisible returns the Visible field value
// Deprecated
func (o *ResumesMineItem) GetVisible() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.Visible
}

// GetVisibleOk returns a tuple with the Visible field value
// and a boolean to check if the value has been set.
// Deprecated
func (o *ResumesMineItem) GetVisibleOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Visible, true
}

// SetVisible sets field value
// Deprecated
func (o *ResumesMineItem) SetVisible(v bool) {
	o.Visible = v
}

func (o ResumesMineItem) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ResumesMineItem) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["actions"] = o.Actions
	if o.Age.IsSet() {
		toSerialize["age"] = o.Age.Get()
	}
	toSerialize["alternate_url"] = o.AlternateUrl
	if !IsNil(o.Area) {
		toSerialize["area"] = o.Area
	}
	if o.AutoHideTime.IsSet() {
		toSerialize["auto_hide_time"] = o.AutoHideTime.Get()
	}
	if o.CanViewFullInfo.IsSet() {
		toSerialize["can_view_full_info"] = o.CanViewFullInfo.Get()
	}
	toSerialize["certificate"] = o.Certificate
	toSerialize["created_at"] = o.CreatedAt
	toSerialize["download"] = o.Download
	toSerialize["education"] = o.Education
	toSerialize["experience"] = o.Experience
	if o.FirstName.IsSet() {
		toSerialize["first_name"] = o.FirstName.Get()
	}
	if !IsNil(o.Gender) {
		toSerialize["gender"] = o.Gender
	}
	toSerialize["hidden_fields"] = o.HiddenFields
	toSerialize["id"] = o.Id
	if o.LastName.IsSet() {
		toSerialize["last_name"] = o.LastName.Get()
	}
	toSerialize["marked"] = o.Marked
	if o.MiddleName.IsSet() {
		toSerialize["middle_name"] = o.MiddleName.Get()
	}
	if !IsNil(o.Photo) {
		toSerialize["photo"] = o.Photo
	}
	if !IsNil(o.Platform) {
		toSerialize["platform"] = o.Platform
	}
	if o.Salary.IsSet() {
		toSerialize["salary"] = o.Salary.Get()
	}
	if o.Title.IsSet() {
		toSerialize["title"] = o.Title.Get()
	}
	if o.TotalExperience.IsSet() {
		toSerialize["total_experience"] = o.TotalExperience.Get()
	}
	toSerialize["updated_at"] = o.UpdatedAt
	toSerialize["url"] = o.Url
	toSerialize["blocked"] = o.Blocked
	if o.CanPublishOrUpdate.IsSet() {
		toSerialize["can_publish_or_update"] = o.CanPublishOrUpdate.Get()
	}
	toSerialize["finished"] = o.Finished
	toSerialize["status"] = o.Status
	toSerialize["access"] = o.Access
	toSerialize["new_views"] = o.NewViews
	if o.NextPublishAt.IsSet() {
		toSerialize["next_publish_at"] = o.NextPublishAt.Get()
	}
	toSerialize["paid_services"] = o.PaidServices
	toSerialize["total_views"] = o.TotalViews
	toSerialize["views_url"] = o.ViewsUrl
	toSerialize["contact"] = o.Contact
	toSerialize["created"] = o.Created
	toSerialize["similar_vacancies"] = o.SimilarVacancies
	toSerialize["updated"] = o.Updated
	toSerialize["visible"] = o.Visible
	return toSerialize, nil
}

func (o *ResumesMineItem) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"actions",
		"alternate_url",
		"certificate",
		"created_at",
		"download",
		"education",
		"experience",
		"hidden_fields",
		"id",
		"marked",
		"updated_at",
		"url",
		"blocked",
		"finished",
		"status",
		"access",
		"new_views",
		"paid_services",
		"total_views",
		"views_url",
		"contact",
		"created",
		"similar_vacancies",
		"updated",
		"visible",
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

	varResumesMineItem := _ResumesMineItem{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varResumesMineItem)

	if err != nil {
		return err
	}

	*o = ResumesMineItem(varResumesMineItem)

	return err
}

type NullableResumesMineItem struct {
	value *ResumesMineItem
	isSet bool
}

func (v NullableResumesMineItem) Get() *ResumesMineItem {
	return v.value
}

func (v *NullableResumesMineItem) Set(val *ResumesMineItem) {
	v.value = val
	v.isSet = true
}

func (v NullableResumesMineItem) IsSet() bool {
	return v.isSet
}

func (v *NullableResumesMineItem) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableResumesMineItem(val *ResumesMineItem) *NullableResumesMineItem {
	return &NullableResumesMineItem{value: val, isSet: true}
}

func (v NullableResumesMineItem) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableResumesMineItem) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


