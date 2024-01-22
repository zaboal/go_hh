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

// checks if the ResumeResumeForApplicant type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ResumeResumeForApplicant{}

// ResumeResumeForApplicant struct for ResumeResumeForApplicant
type ResumeResumeForApplicant struct {
	// URL резюме на сайте
	AlternateUrl string `json:"alternate_url"`
	// Идентификатор резюме
	Id string `json:"id"`
	// Желаемая должность
	Title NullableString `json:"title"`
	// Возраст
	Age NullableFloat32 `json:"age,omitempty"`
	Area NullableIncludesIdNameUrl `json:"area,omitempty"`
	// Доступен ли просмотр контактной информации в резюме текущему работодателю
	CanViewFullInfo NullableBool `json:"can_view_full_info,omitempty"`
	// Список сертификатов соискателя
	Certificate []ResumeObjectsCertificate `json:"certificate"`
	// Дата и время создания резюме
	CreatedAt string `json:"created_at"`
	Download ResumeResumeProfileAllOfDownload `json:"download"`
	Education ResumeResumeProfileAllOfEducation `json:"education"`
	// Опыт работы
	Experience []ResumeObjectsExperience `json:"experience"`
	// Имя
	FirstName NullableString `json:"first_name,omitempty"`
	Gender NullableIncludesIdName `json:"gender,omitempty"`
	// Справочник [Список скрытых полей](https://github.com/hhru/api/blob/master/docs/employer_resumes.md#hidden-fields). Возможные значения элементов приведены в поле `resume_hidden_fields` [справочника полей](#tag/Obshie-spravochniki/operation/get-dictionaries)
	HiddenFields []IncludesIdName `json:"hidden_fields"`
	// Фамилия
	LastName NullableString `json:"last_name,omitempty"`
	// Выделено ли резюме в поиске
	Marked *bool `json:"marked,omitempty"`
	// Отчество
	MiddleName NullableString `json:"middle_name,omitempty"`
	Platform *ResumeResumeProfileAllOfPlatform `json:"platform,omitempty"`
	Salary NullableResumeObjectsSalaryProperties `json:"salary,omitempty"`
	TotalExperience NullableResumeObjectsTotalExperience `json:"total_experience,omitempty"`
	// Дата и время обновления резюме
	UpdatedAt string `json:"updated_at"`
	// День рождения (в формате `ГГГГ-ММ-ДД`)
	BirthDate NullableString `json:"birth_date,omitempty"`
	BusinessTripReadiness ResumeResumeFullAllOfBusinessTripReadiness `json:"business_trip_readiness"`
	// Список гражданств соискателя. Элементы [справочника регионов](#tag/Obshie-spravochniki/operation/get-areas)
	Citizenship []IncludesIdNameUrl `json:"citizenship"`
	// Список контактов соискателя
	Contact []IncludesContact `json:"contact"`
	Creds NullableCredsCreds `json:"creds,omitempty"`
	// Список категорий водительских прав соискателя
	DriverLicenseTypes []ResumeObjectsDriverLicenseTypes `json:"driver_license_types"`
	// Deprecated
	Employment *ResumeResumeFullAllOfEmployment `json:"employment,omitempty"`
	// Список подходящих соискателю типов занятостей. Элементы справочника [employment](#tag/Obshie-spravochniki/operation/get-dictionaries)
	Employments []IncludesIdName `json:"employments"`
	// Наличие личного автомобиля у соискателя
	HasVehicle NullableBool `json:"has_vehicle,omitempty"`
	// Список языков, которыми владеет соискатель. Элементы справочника [languages](#tag/Obshie-spravochniki/operation/get-languages)
	Language []IncludesLanguageLevel `json:"language"`
	Metro NullableResumeObjectsMetroStation `json:"metro,omitempty"`
	// Платные услуги по резюме для автора
	PaidServices []ResumeObjectsPaidServices `json:"paid_services"`
	// Массив объектов профролей
	ProfessionalRoles []IncludesIdName `json:"professional_roles,omitempty"`
	// Список рекомендаций
	Recommendation []ResumeObjectsRecommendation `json:"recommendation"`
	Relocation ResumeResumeFullAllOfRelocation `json:"relocation"`
	ResumeLocale ResumeResumeFullAllOfResumeLocale `json:"resume_locale"`
	// Deprecated
	Schedule ResumeResumeFullAllOfSchedule `json:"schedule"`
	// Список подходящих соискателю графиков работы. Элементы справочника [schedule](#tag/Obshie-spravochniki/operation/get-dictionaries)
	Schedules []IncludesIdName `json:"schedules"`
	// Профили в соц. сетях и других сервисах
	Site []ResumeObjectsSite `json:"site"`
	// Ключевые навыки (список уникальных строк)
	SkillSet []string `json:"skill_set"`
	// Дополнительная информация, описание навыков в свободной форме
	Skills NullableString `json:"skills,omitempty"`
	TravelTime ResumeResumeFullAllOfTravelTime `json:"travel_time"`
	// Список регионов, в которых соискатель имеет разрешение на работу. Элементы [справочника регионов](#tag/Obshie-spravochniki/operation/get-areas) 
	WorkTicket []IncludesIdNameUrl `json:"work_ticket"`
	// Заблокировано ли резюме ([подробнее](#tag/Rezyume.-Prosmotr-informacii/Status-rezyume))
	Blocked bool `json:"blocked"`
	// Можно ли опубликовать или обновить данное резюме
	CanPublishOrUpdate NullableBool `json:"can_publish_or_update,omitempty"`
	// Заполнено ли резюме
	Finished bool `json:"finished"`
	Status IncludesIdName `json:"status"`
	// Замечания модератора. В некоторых случаях замечания могут сопровождаться [блокировкой резюме](#tag/Rezyume.-Prosmotr-informacii/Status-rezyume). Полный список возможных замечаний доступен в поле `resume_moderation_note` [в справочнике полей](#tag/Obshie-spravochniki/operation/get-dictionaries) 
	ModerationNote []ResumeObjectsModerationNote `json:"moderation_note"`
	Progress ResumeObjectsProgress `json:"progress"`
	// URL для публикации или обновления резюме
	PublishUrl string `json:"publish_url"`
	Access ResumeObjectsAccess `json:"access"`
	// Дополнительные действия
	Actions ResumeObjectsActionsForOwner `json:"actions"`
	// Число новых просмотров. Данный счетчик сбрасывается при получении [детальной истории просмотров](#tag/Rezyume.-Prosmotr-informacii/operation/get-resume-view-history) 
	NewViews float32 `json:"new_views"`
	// Дата и время следующей возможной публикации/обновления. Для неопубликованных резюме возвращается `null`
	NextPublishAt NullableString `json:"next_publish_at,omitempty"`
	// Число просмотров резюме
	TotalViews float32 `json:"total_views"`
	// URL, по которому необходимо сделать GET-запрос для получения [детальной истории просмотров](#tag/Rezyume.-Prosmotr-informacii/operation/get-resume-view-history) 
	ViewsUrl string `json:"views_url"`
	Photo NullableResumeObjectsPhoto `json:"photo,omitempty"`
	// Список изображений в портфолио пользователя
	Portfolio []ResumeObjectsPortfolio `json:"portfolio"`
}

type _ResumeResumeForApplicant ResumeResumeForApplicant

// NewResumeResumeForApplicant instantiates a new ResumeResumeForApplicant object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewResumeResumeForApplicant(alternateUrl string, id string, title NullableString, certificate []ResumeObjectsCertificate, createdAt string, download ResumeResumeProfileAllOfDownload, education ResumeResumeProfileAllOfEducation, experience []ResumeObjectsExperience, hiddenFields []IncludesIdName, updatedAt string, businessTripReadiness ResumeResumeFullAllOfBusinessTripReadiness, citizenship []IncludesIdNameUrl, contact []IncludesContact, driverLicenseTypes []ResumeObjectsDriverLicenseTypes, employments []IncludesIdName, language []IncludesLanguageLevel, paidServices []ResumeObjectsPaidServices, recommendation []ResumeObjectsRecommendation, relocation ResumeResumeFullAllOfRelocation, resumeLocale ResumeResumeFullAllOfResumeLocale, schedule ResumeResumeFullAllOfSchedule, schedules []IncludesIdName, site []ResumeObjectsSite, skillSet []string, travelTime ResumeResumeFullAllOfTravelTime, workTicket []IncludesIdNameUrl, blocked bool, finished bool, status IncludesIdName, moderationNote []ResumeObjectsModerationNote, progress ResumeObjectsProgress, publishUrl string, access ResumeObjectsAccess, actions ResumeObjectsActionsForOwner, newViews float32, totalViews float32, viewsUrl string, portfolio []ResumeObjectsPortfolio) *ResumeResumeForApplicant {
	this := ResumeResumeForApplicant{}
	this.AlternateUrl = alternateUrl
	this.Id = id
	this.Title = title
	this.Certificate = certificate
	this.CreatedAt = createdAt
	this.Download = download
	this.Education = education
	this.Experience = experience
	this.HiddenFields = hiddenFields
	var marked bool = false
	this.Marked = &marked
	this.UpdatedAt = updatedAt
	this.BusinessTripReadiness = businessTripReadiness
	this.Citizenship = citizenship
	this.Contact = contact
	this.DriverLicenseTypes = driverLicenseTypes
	this.Employments = employments
	this.Language = language
	this.PaidServices = paidServices
	this.Recommendation = recommendation
	this.Relocation = relocation
	this.ResumeLocale = resumeLocale
	this.Schedule = schedule
	this.Schedules = schedules
	this.Site = site
	this.SkillSet = skillSet
	this.TravelTime = travelTime
	this.WorkTicket = workTicket
	this.Blocked = blocked
	this.Finished = finished
	this.Status = status
	this.ModerationNote = moderationNote
	this.Progress = progress
	this.PublishUrl = publishUrl
	this.Access = access
	this.Actions = actions
	this.NewViews = newViews
	this.TotalViews = totalViews
	this.ViewsUrl = viewsUrl
	this.Portfolio = portfolio
	return &this
}

// NewResumeResumeForApplicantWithDefaults instantiates a new ResumeResumeForApplicant object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewResumeResumeForApplicantWithDefaults() *ResumeResumeForApplicant {
	this := ResumeResumeForApplicant{}
	var marked bool = false
	this.Marked = &marked
	return &this
}

// GetAlternateUrl returns the AlternateUrl field value
func (o *ResumeResumeForApplicant) GetAlternateUrl() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.AlternateUrl
}

// GetAlternateUrlOk returns a tuple with the AlternateUrl field value
// and a boolean to check if the value has been set.
func (o *ResumeResumeForApplicant) GetAlternateUrlOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AlternateUrl, true
}

// SetAlternateUrl sets field value
func (o *ResumeResumeForApplicant) SetAlternateUrl(v string) {
	o.AlternateUrl = v
}

// GetId returns the Id field value
func (o *ResumeResumeForApplicant) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *ResumeResumeForApplicant) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *ResumeResumeForApplicant) SetId(v string) {
	o.Id = v
}

// GetTitle returns the Title field value
// If the value is explicit nil, the zero value for string will be returned
func (o *ResumeResumeForApplicant) GetTitle() string {
	if o == nil || o.Title.Get() == nil {
		var ret string
		return ret
	}

	return *o.Title.Get()
}

// GetTitleOk returns a tuple with the Title field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ResumeResumeForApplicant) GetTitleOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Title.Get(), o.Title.IsSet()
}

// SetTitle sets field value
func (o *ResumeResumeForApplicant) SetTitle(v string) {
	o.Title.Set(&v)
}

// GetAge returns the Age field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ResumeResumeForApplicant) GetAge() float32 {
	if o == nil || IsNil(o.Age.Get()) {
		var ret float32
		return ret
	}
	return *o.Age.Get()
}

// GetAgeOk returns a tuple with the Age field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ResumeResumeForApplicant) GetAgeOk() (*float32, bool) {
	if o == nil {
		return nil, false
	}
	return o.Age.Get(), o.Age.IsSet()
}

// HasAge returns a boolean if a field has been set.
func (o *ResumeResumeForApplicant) HasAge() bool {
	if o != nil && o.Age.IsSet() {
		return true
	}

	return false
}

// SetAge gets a reference to the given NullableFloat32 and assigns it to the Age field.
func (o *ResumeResumeForApplicant) SetAge(v float32) {
	o.Age.Set(&v)
}
// SetAgeNil sets the value for Age to be an explicit nil
func (o *ResumeResumeForApplicant) SetAgeNil() {
	o.Age.Set(nil)
}

// UnsetAge ensures that no value is present for Age, not even an explicit nil
func (o *ResumeResumeForApplicant) UnsetAge() {
	o.Age.Unset()
}

// GetArea returns the Area field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ResumeResumeForApplicant) GetArea() IncludesIdNameUrl {
	if o == nil || IsNil(o.Area.Get()) {
		var ret IncludesIdNameUrl
		return ret
	}
	return *o.Area.Get()
}

// GetAreaOk returns a tuple with the Area field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ResumeResumeForApplicant) GetAreaOk() (*IncludesIdNameUrl, bool) {
	if o == nil {
		return nil, false
	}
	return o.Area.Get(), o.Area.IsSet()
}

// HasArea returns a boolean if a field has been set.
func (o *ResumeResumeForApplicant) HasArea() bool {
	if o != nil && o.Area.IsSet() {
		return true
	}

	return false
}

// SetArea gets a reference to the given NullableIncludesIdNameUrl and assigns it to the Area field.
func (o *ResumeResumeForApplicant) SetArea(v IncludesIdNameUrl) {
	o.Area.Set(&v)
}
// SetAreaNil sets the value for Area to be an explicit nil
func (o *ResumeResumeForApplicant) SetAreaNil() {
	o.Area.Set(nil)
}

// UnsetArea ensures that no value is present for Area, not even an explicit nil
func (o *ResumeResumeForApplicant) UnsetArea() {
	o.Area.Unset()
}

// GetCanViewFullInfo returns the CanViewFullInfo field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ResumeResumeForApplicant) GetCanViewFullInfo() bool {
	if o == nil || IsNil(o.CanViewFullInfo.Get()) {
		var ret bool
		return ret
	}
	return *o.CanViewFullInfo.Get()
}

// GetCanViewFullInfoOk returns a tuple with the CanViewFullInfo field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ResumeResumeForApplicant) GetCanViewFullInfoOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return o.CanViewFullInfo.Get(), o.CanViewFullInfo.IsSet()
}

// HasCanViewFullInfo returns a boolean if a field has been set.
func (o *ResumeResumeForApplicant) HasCanViewFullInfo() bool {
	if o != nil && o.CanViewFullInfo.IsSet() {
		return true
	}

	return false
}

// SetCanViewFullInfo gets a reference to the given NullableBool and assigns it to the CanViewFullInfo field.
func (o *ResumeResumeForApplicant) SetCanViewFullInfo(v bool) {
	o.CanViewFullInfo.Set(&v)
}
// SetCanViewFullInfoNil sets the value for CanViewFullInfo to be an explicit nil
func (o *ResumeResumeForApplicant) SetCanViewFullInfoNil() {
	o.CanViewFullInfo.Set(nil)
}

// UnsetCanViewFullInfo ensures that no value is present for CanViewFullInfo, not even an explicit nil
func (o *ResumeResumeForApplicant) UnsetCanViewFullInfo() {
	o.CanViewFullInfo.Unset()
}

// GetCertificate returns the Certificate field value
func (o *ResumeResumeForApplicant) GetCertificate() []ResumeObjectsCertificate {
	if o == nil {
		var ret []ResumeObjectsCertificate
		return ret
	}

	return o.Certificate
}

// GetCertificateOk returns a tuple with the Certificate field value
// and a boolean to check if the value has been set.
func (o *ResumeResumeForApplicant) GetCertificateOk() ([]ResumeObjectsCertificate, bool) {
	if o == nil {
		return nil, false
	}
	return o.Certificate, true
}

// SetCertificate sets field value
func (o *ResumeResumeForApplicant) SetCertificate(v []ResumeObjectsCertificate) {
	o.Certificate = v
}

// GetCreatedAt returns the CreatedAt field value
func (o *ResumeResumeForApplicant) GetCreatedAt() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value
// and a boolean to check if the value has been set.
func (o *ResumeResumeForApplicant) GetCreatedAtOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CreatedAt, true
}

// SetCreatedAt sets field value
func (o *ResumeResumeForApplicant) SetCreatedAt(v string) {
	o.CreatedAt = v
}

// GetDownload returns the Download field value
func (o *ResumeResumeForApplicant) GetDownload() ResumeResumeProfileAllOfDownload {
	if o == nil {
		var ret ResumeResumeProfileAllOfDownload
		return ret
	}

	return o.Download
}

// GetDownloadOk returns a tuple with the Download field value
// and a boolean to check if the value has been set.
func (o *ResumeResumeForApplicant) GetDownloadOk() (*ResumeResumeProfileAllOfDownload, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Download, true
}

// SetDownload sets field value
func (o *ResumeResumeForApplicant) SetDownload(v ResumeResumeProfileAllOfDownload) {
	o.Download = v
}

// GetEducation returns the Education field value
func (o *ResumeResumeForApplicant) GetEducation() ResumeResumeProfileAllOfEducation {
	if o == nil {
		var ret ResumeResumeProfileAllOfEducation
		return ret
	}

	return o.Education
}

// GetEducationOk returns a tuple with the Education field value
// and a boolean to check if the value has been set.
func (o *ResumeResumeForApplicant) GetEducationOk() (*ResumeResumeProfileAllOfEducation, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Education, true
}

// SetEducation sets field value
func (o *ResumeResumeForApplicant) SetEducation(v ResumeResumeProfileAllOfEducation) {
	o.Education = v
}

// GetExperience returns the Experience field value
func (o *ResumeResumeForApplicant) GetExperience() []ResumeObjectsExperience {
	if o == nil {
		var ret []ResumeObjectsExperience
		return ret
	}

	return o.Experience
}

// GetExperienceOk returns a tuple with the Experience field value
// and a boolean to check if the value has been set.
func (o *ResumeResumeForApplicant) GetExperienceOk() ([]ResumeObjectsExperience, bool) {
	if o == nil {
		return nil, false
	}
	return o.Experience, true
}

// SetExperience sets field value
func (o *ResumeResumeForApplicant) SetExperience(v []ResumeObjectsExperience) {
	o.Experience = v
}

// GetFirstName returns the FirstName field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ResumeResumeForApplicant) GetFirstName() string {
	if o == nil || IsNil(o.FirstName.Get()) {
		var ret string
		return ret
	}
	return *o.FirstName.Get()
}

// GetFirstNameOk returns a tuple with the FirstName field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ResumeResumeForApplicant) GetFirstNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.FirstName.Get(), o.FirstName.IsSet()
}

// HasFirstName returns a boolean if a field has been set.
func (o *ResumeResumeForApplicant) HasFirstName() bool {
	if o != nil && o.FirstName.IsSet() {
		return true
	}

	return false
}

// SetFirstName gets a reference to the given NullableString and assigns it to the FirstName field.
func (o *ResumeResumeForApplicant) SetFirstName(v string) {
	o.FirstName.Set(&v)
}
// SetFirstNameNil sets the value for FirstName to be an explicit nil
func (o *ResumeResumeForApplicant) SetFirstNameNil() {
	o.FirstName.Set(nil)
}

// UnsetFirstName ensures that no value is present for FirstName, not even an explicit nil
func (o *ResumeResumeForApplicant) UnsetFirstName() {
	o.FirstName.Unset()
}

// GetGender returns the Gender field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ResumeResumeForApplicant) GetGender() IncludesIdName {
	if o == nil || IsNil(o.Gender.Get()) {
		var ret IncludesIdName
		return ret
	}
	return *o.Gender.Get()
}

// GetGenderOk returns a tuple with the Gender field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ResumeResumeForApplicant) GetGenderOk() (*IncludesIdName, bool) {
	if o == nil {
		return nil, false
	}
	return o.Gender.Get(), o.Gender.IsSet()
}

// HasGender returns a boolean if a field has been set.
func (o *ResumeResumeForApplicant) HasGender() bool {
	if o != nil && o.Gender.IsSet() {
		return true
	}

	return false
}

// SetGender gets a reference to the given NullableIncludesIdName and assigns it to the Gender field.
func (o *ResumeResumeForApplicant) SetGender(v IncludesIdName) {
	o.Gender.Set(&v)
}
// SetGenderNil sets the value for Gender to be an explicit nil
func (o *ResumeResumeForApplicant) SetGenderNil() {
	o.Gender.Set(nil)
}

// UnsetGender ensures that no value is present for Gender, not even an explicit nil
func (o *ResumeResumeForApplicant) UnsetGender() {
	o.Gender.Unset()
}

// GetHiddenFields returns the HiddenFields field value
func (o *ResumeResumeForApplicant) GetHiddenFields() []IncludesIdName {
	if o == nil {
		var ret []IncludesIdName
		return ret
	}

	return o.HiddenFields
}

// GetHiddenFieldsOk returns a tuple with the HiddenFields field value
// and a boolean to check if the value has been set.
func (o *ResumeResumeForApplicant) GetHiddenFieldsOk() ([]IncludesIdName, bool) {
	if o == nil {
		return nil, false
	}
	return o.HiddenFields, true
}

// SetHiddenFields sets field value
func (o *ResumeResumeForApplicant) SetHiddenFields(v []IncludesIdName) {
	o.HiddenFields = v
}

// GetLastName returns the LastName field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ResumeResumeForApplicant) GetLastName() string {
	if o == nil || IsNil(o.LastName.Get()) {
		var ret string
		return ret
	}
	return *o.LastName.Get()
}

// GetLastNameOk returns a tuple with the LastName field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ResumeResumeForApplicant) GetLastNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.LastName.Get(), o.LastName.IsSet()
}

// HasLastName returns a boolean if a field has been set.
func (o *ResumeResumeForApplicant) HasLastName() bool {
	if o != nil && o.LastName.IsSet() {
		return true
	}

	return false
}

// SetLastName gets a reference to the given NullableString and assigns it to the LastName field.
func (o *ResumeResumeForApplicant) SetLastName(v string) {
	o.LastName.Set(&v)
}
// SetLastNameNil sets the value for LastName to be an explicit nil
func (o *ResumeResumeForApplicant) SetLastNameNil() {
	o.LastName.Set(nil)
}

// UnsetLastName ensures that no value is present for LastName, not even an explicit nil
func (o *ResumeResumeForApplicant) UnsetLastName() {
	o.LastName.Unset()
}

// GetMarked returns the Marked field value if set, zero value otherwise.
func (o *ResumeResumeForApplicant) GetMarked() bool {
	if o == nil || IsNil(o.Marked) {
		var ret bool
		return ret
	}
	return *o.Marked
}

// GetMarkedOk returns a tuple with the Marked field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ResumeResumeForApplicant) GetMarkedOk() (*bool, bool) {
	if o == nil || IsNil(o.Marked) {
		return nil, false
	}
	return o.Marked, true
}

// HasMarked returns a boolean if a field has been set.
func (o *ResumeResumeForApplicant) HasMarked() bool {
	if o != nil && !IsNil(o.Marked) {
		return true
	}

	return false
}

// SetMarked gets a reference to the given bool and assigns it to the Marked field.
func (o *ResumeResumeForApplicant) SetMarked(v bool) {
	o.Marked = &v
}

// GetMiddleName returns the MiddleName field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ResumeResumeForApplicant) GetMiddleName() string {
	if o == nil || IsNil(o.MiddleName.Get()) {
		var ret string
		return ret
	}
	return *o.MiddleName.Get()
}

// GetMiddleNameOk returns a tuple with the MiddleName field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ResumeResumeForApplicant) GetMiddleNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.MiddleName.Get(), o.MiddleName.IsSet()
}

// HasMiddleName returns a boolean if a field has been set.
func (o *ResumeResumeForApplicant) HasMiddleName() bool {
	if o != nil && o.MiddleName.IsSet() {
		return true
	}

	return false
}

// SetMiddleName gets a reference to the given NullableString and assigns it to the MiddleName field.
func (o *ResumeResumeForApplicant) SetMiddleName(v string) {
	o.MiddleName.Set(&v)
}
// SetMiddleNameNil sets the value for MiddleName to be an explicit nil
func (o *ResumeResumeForApplicant) SetMiddleNameNil() {
	o.MiddleName.Set(nil)
}

// UnsetMiddleName ensures that no value is present for MiddleName, not even an explicit nil
func (o *ResumeResumeForApplicant) UnsetMiddleName() {
	o.MiddleName.Unset()
}

// GetPlatform returns the Platform field value if set, zero value otherwise.
func (o *ResumeResumeForApplicant) GetPlatform() ResumeResumeProfileAllOfPlatform {
	if o == nil || IsNil(o.Platform) {
		var ret ResumeResumeProfileAllOfPlatform
		return ret
	}
	return *o.Platform
}

// GetPlatformOk returns a tuple with the Platform field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ResumeResumeForApplicant) GetPlatformOk() (*ResumeResumeProfileAllOfPlatform, bool) {
	if o == nil || IsNil(o.Platform) {
		return nil, false
	}
	return o.Platform, true
}

// HasPlatform returns a boolean if a field has been set.
func (o *ResumeResumeForApplicant) HasPlatform() bool {
	if o != nil && !IsNil(o.Platform) {
		return true
	}

	return false
}

// SetPlatform gets a reference to the given ResumeResumeProfileAllOfPlatform and assigns it to the Platform field.
func (o *ResumeResumeForApplicant) SetPlatform(v ResumeResumeProfileAllOfPlatform) {
	o.Platform = &v
}

// GetSalary returns the Salary field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ResumeResumeForApplicant) GetSalary() ResumeObjectsSalaryProperties {
	if o == nil || IsNil(o.Salary.Get()) {
		var ret ResumeObjectsSalaryProperties
		return ret
	}
	return *o.Salary.Get()
}

// GetSalaryOk returns a tuple with the Salary field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ResumeResumeForApplicant) GetSalaryOk() (*ResumeObjectsSalaryProperties, bool) {
	if o == nil {
		return nil, false
	}
	return o.Salary.Get(), o.Salary.IsSet()
}

// HasSalary returns a boolean if a field has been set.
func (o *ResumeResumeForApplicant) HasSalary() bool {
	if o != nil && o.Salary.IsSet() {
		return true
	}

	return false
}

// SetSalary gets a reference to the given NullableResumeObjectsSalaryProperties and assigns it to the Salary field.
func (o *ResumeResumeForApplicant) SetSalary(v ResumeObjectsSalaryProperties) {
	o.Salary.Set(&v)
}
// SetSalaryNil sets the value for Salary to be an explicit nil
func (o *ResumeResumeForApplicant) SetSalaryNil() {
	o.Salary.Set(nil)
}

// UnsetSalary ensures that no value is present for Salary, not even an explicit nil
func (o *ResumeResumeForApplicant) UnsetSalary() {
	o.Salary.Unset()
}

// GetTotalExperience returns the TotalExperience field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ResumeResumeForApplicant) GetTotalExperience() ResumeObjectsTotalExperience {
	if o == nil || IsNil(o.TotalExperience.Get()) {
		var ret ResumeObjectsTotalExperience
		return ret
	}
	return *o.TotalExperience.Get()
}

// GetTotalExperienceOk returns a tuple with the TotalExperience field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ResumeResumeForApplicant) GetTotalExperienceOk() (*ResumeObjectsTotalExperience, bool) {
	if o == nil {
		return nil, false
	}
	return o.TotalExperience.Get(), o.TotalExperience.IsSet()
}

// HasTotalExperience returns a boolean if a field has been set.
func (o *ResumeResumeForApplicant) HasTotalExperience() bool {
	if o != nil && o.TotalExperience.IsSet() {
		return true
	}

	return false
}

// SetTotalExperience gets a reference to the given NullableResumeObjectsTotalExperience and assigns it to the TotalExperience field.
func (o *ResumeResumeForApplicant) SetTotalExperience(v ResumeObjectsTotalExperience) {
	o.TotalExperience.Set(&v)
}
// SetTotalExperienceNil sets the value for TotalExperience to be an explicit nil
func (o *ResumeResumeForApplicant) SetTotalExperienceNil() {
	o.TotalExperience.Set(nil)
}

// UnsetTotalExperience ensures that no value is present for TotalExperience, not even an explicit nil
func (o *ResumeResumeForApplicant) UnsetTotalExperience() {
	o.TotalExperience.Unset()
}

// GetUpdatedAt returns the UpdatedAt field value
func (o *ResumeResumeForApplicant) GetUpdatedAt() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.UpdatedAt
}

// GetUpdatedAtOk returns a tuple with the UpdatedAt field value
// and a boolean to check if the value has been set.
func (o *ResumeResumeForApplicant) GetUpdatedAtOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.UpdatedAt, true
}

// SetUpdatedAt sets field value
func (o *ResumeResumeForApplicant) SetUpdatedAt(v string) {
	o.UpdatedAt = v
}

// GetBirthDate returns the BirthDate field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ResumeResumeForApplicant) GetBirthDate() string {
	if o == nil || IsNil(o.BirthDate.Get()) {
		var ret string
		return ret
	}
	return *o.BirthDate.Get()
}

// GetBirthDateOk returns a tuple with the BirthDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ResumeResumeForApplicant) GetBirthDateOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.BirthDate.Get(), o.BirthDate.IsSet()
}

// HasBirthDate returns a boolean if a field has been set.
func (o *ResumeResumeForApplicant) HasBirthDate() bool {
	if o != nil && o.BirthDate.IsSet() {
		return true
	}

	return false
}

// SetBirthDate gets a reference to the given NullableString and assigns it to the BirthDate field.
func (o *ResumeResumeForApplicant) SetBirthDate(v string) {
	o.BirthDate.Set(&v)
}
// SetBirthDateNil sets the value for BirthDate to be an explicit nil
func (o *ResumeResumeForApplicant) SetBirthDateNil() {
	o.BirthDate.Set(nil)
}

// UnsetBirthDate ensures that no value is present for BirthDate, not even an explicit nil
func (o *ResumeResumeForApplicant) UnsetBirthDate() {
	o.BirthDate.Unset()
}

// GetBusinessTripReadiness returns the BusinessTripReadiness field value
func (o *ResumeResumeForApplicant) GetBusinessTripReadiness() ResumeResumeFullAllOfBusinessTripReadiness {
	if o == nil {
		var ret ResumeResumeFullAllOfBusinessTripReadiness
		return ret
	}

	return o.BusinessTripReadiness
}

// GetBusinessTripReadinessOk returns a tuple with the BusinessTripReadiness field value
// and a boolean to check if the value has been set.
func (o *ResumeResumeForApplicant) GetBusinessTripReadinessOk() (*ResumeResumeFullAllOfBusinessTripReadiness, bool) {
	if o == nil {
		return nil, false
	}
	return &o.BusinessTripReadiness, true
}

// SetBusinessTripReadiness sets field value
func (o *ResumeResumeForApplicant) SetBusinessTripReadiness(v ResumeResumeFullAllOfBusinessTripReadiness) {
	o.BusinessTripReadiness = v
}

// GetCitizenship returns the Citizenship field value
func (o *ResumeResumeForApplicant) GetCitizenship() []IncludesIdNameUrl {
	if o == nil {
		var ret []IncludesIdNameUrl
		return ret
	}

	return o.Citizenship
}

// GetCitizenshipOk returns a tuple with the Citizenship field value
// and a boolean to check if the value has been set.
func (o *ResumeResumeForApplicant) GetCitizenshipOk() ([]IncludesIdNameUrl, bool) {
	if o == nil {
		return nil, false
	}
	return o.Citizenship, true
}

// SetCitizenship sets field value
func (o *ResumeResumeForApplicant) SetCitizenship(v []IncludesIdNameUrl) {
	o.Citizenship = v
}

// GetContact returns the Contact field value
func (o *ResumeResumeForApplicant) GetContact() []IncludesContact {
	if o == nil {
		var ret []IncludesContact
		return ret
	}

	return o.Contact
}

// GetContactOk returns a tuple with the Contact field value
// and a boolean to check if the value has been set.
func (o *ResumeResumeForApplicant) GetContactOk() ([]IncludesContact, bool) {
	if o == nil {
		return nil, false
	}
	return o.Contact, true
}

// SetContact sets field value
func (o *ResumeResumeForApplicant) SetContact(v []IncludesContact) {
	o.Contact = v
}

// GetCreds returns the Creds field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ResumeResumeForApplicant) GetCreds() CredsCreds {
	if o == nil || IsNil(o.Creds.Get()) {
		var ret CredsCreds
		return ret
	}
	return *o.Creds.Get()
}

// GetCredsOk returns a tuple with the Creds field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ResumeResumeForApplicant) GetCredsOk() (*CredsCreds, bool) {
	if o == nil {
		return nil, false
	}
	return o.Creds.Get(), o.Creds.IsSet()
}

// HasCreds returns a boolean if a field has been set.
func (o *ResumeResumeForApplicant) HasCreds() bool {
	if o != nil && o.Creds.IsSet() {
		return true
	}

	return false
}

// SetCreds gets a reference to the given NullableCredsCreds and assigns it to the Creds field.
func (o *ResumeResumeForApplicant) SetCreds(v CredsCreds) {
	o.Creds.Set(&v)
}
// SetCredsNil sets the value for Creds to be an explicit nil
func (o *ResumeResumeForApplicant) SetCredsNil() {
	o.Creds.Set(nil)
}

// UnsetCreds ensures that no value is present for Creds, not even an explicit nil
func (o *ResumeResumeForApplicant) UnsetCreds() {
	o.Creds.Unset()
}

// GetDriverLicenseTypes returns the DriverLicenseTypes field value
func (o *ResumeResumeForApplicant) GetDriverLicenseTypes() []ResumeObjectsDriverLicenseTypes {
	if o == nil {
		var ret []ResumeObjectsDriverLicenseTypes
		return ret
	}

	return o.DriverLicenseTypes
}

// GetDriverLicenseTypesOk returns a tuple with the DriverLicenseTypes field value
// and a boolean to check if the value has been set.
func (o *ResumeResumeForApplicant) GetDriverLicenseTypesOk() ([]ResumeObjectsDriverLicenseTypes, bool) {
	if o == nil {
		return nil, false
	}
	return o.DriverLicenseTypes, true
}

// SetDriverLicenseTypes sets field value
func (o *ResumeResumeForApplicant) SetDriverLicenseTypes(v []ResumeObjectsDriverLicenseTypes) {
	o.DriverLicenseTypes = v
}

// GetEmployment returns the Employment field value if set, zero value otherwise.
// Deprecated
func (o *ResumeResumeForApplicant) GetEmployment() ResumeResumeFullAllOfEmployment {
	if o == nil || IsNil(o.Employment) {
		var ret ResumeResumeFullAllOfEmployment
		return ret
	}
	return *o.Employment
}

// GetEmploymentOk returns a tuple with the Employment field value if set, nil otherwise
// and a boolean to check if the value has been set.
// Deprecated
func (o *ResumeResumeForApplicant) GetEmploymentOk() (*ResumeResumeFullAllOfEmployment, bool) {
	if o == nil || IsNil(o.Employment) {
		return nil, false
	}
	return o.Employment, true
}

// HasEmployment returns a boolean if a field has been set.
func (o *ResumeResumeForApplicant) HasEmployment() bool {
	if o != nil && !IsNil(o.Employment) {
		return true
	}

	return false
}

// SetEmployment gets a reference to the given ResumeResumeFullAllOfEmployment and assigns it to the Employment field.
// Deprecated
func (o *ResumeResumeForApplicant) SetEmployment(v ResumeResumeFullAllOfEmployment) {
	o.Employment = &v
}

// GetEmployments returns the Employments field value
func (o *ResumeResumeForApplicant) GetEmployments() []IncludesIdName {
	if o == nil {
		var ret []IncludesIdName
		return ret
	}

	return o.Employments
}

// GetEmploymentsOk returns a tuple with the Employments field value
// and a boolean to check if the value has been set.
func (o *ResumeResumeForApplicant) GetEmploymentsOk() ([]IncludesIdName, bool) {
	if o == nil {
		return nil, false
	}
	return o.Employments, true
}

// SetEmployments sets field value
func (o *ResumeResumeForApplicant) SetEmployments(v []IncludesIdName) {
	o.Employments = v
}

// GetHasVehicle returns the HasVehicle field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ResumeResumeForApplicant) GetHasVehicle() bool {
	if o == nil || IsNil(o.HasVehicle.Get()) {
		var ret bool
		return ret
	}
	return *o.HasVehicle.Get()
}

// GetHasVehicleOk returns a tuple with the HasVehicle field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ResumeResumeForApplicant) GetHasVehicleOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return o.HasVehicle.Get(), o.HasVehicle.IsSet()
}

// HasHasVehicle returns a boolean if a field has been set.
func (o *ResumeResumeForApplicant) HasHasVehicle() bool {
	if o != nil && o.HasVehicle.IsSet() {
		return true
	}

	return false
}

// SetHasVehicle gets a reference to the given NullableBool and assigns it to the HasVehicle field.
func (o *ResumeResumeForApplicant) SetHasVehicle(v bool) {
	o.HasVehicle.Set(&v)
}
// SetHasVehicleNil sets the value for HasVehicle to be an explicit nil
func (o *ResumeResumeForApplicant) SetHasVehicleNil() {
	o.HasVehicle.Set(nil)
}

// UnsetHasVehicle ensures that no value is present for HasVehicle, not even an explicit nil
func (o *ResumeResumeForApplicant) UnsetHasVehicle() {
	o.HasVehicle.Unset()
}

// GetLanguage returns the Language field value
func (o *ResumeResumeForApplicant) GetLanguage() []IncludesLanguageLevel {
	if o == nil {
		var ret []IncludesLanguageLevel
		return ret
	}

	return o.Language
}

// GetLanguageOk returns a tuple with the Language field value
// and a boolean to check if the value has been set.
func (o *ResumeResumeForApplicant) GetLanguageOk() ([]IncludesLanguageLevel, bool) {
	if o == nil {
		return nil, false
	}
	return o.Language, true
}

// SetLanguage sets field value
func (o *ResumeResumeForApplicant) SetLanguage(v []IncludesLanguageLevel) {
	o.Language = v
}

// GetMetro returns the Metro field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ResumeResumeForApplicant) GetMetro() ResumeObjectsMetroStation {
	if o == nil || IsNil(o.Metro.Get()) {
		var ret ResumeObjectsMetroStation
		return ret
	}
	return *o.Metro.Get()
}

// GetMetroOk returns a tuple with the Metro field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ResumeResumeForApplicant) GetMetroOk() (*ResumeObjectsMetroStation, bool) {
	if o == nil {
		return nil, false
	}
	return o.Metro.Get(), o.Metro.IsSet()
}

// HasMetro returns a boolean if a field has been set.
func (o *ResumeResumeForApplicant) HasMetro() bool {
	if o != nil && o.Metro.IsSet() {
		return true
	}

	return false
}

// SetMetro gets a reference to the given NullableResumeObjectsMetroStation and assigns it to the Metro field.
func (o *ResumeResumeForApplicant) SetMetro(v ResumeObjectsMetroStation) {
	o.Metro.Set(&v)
}
// SetMetroNil sets the value for Metro to be an explicit nil
func (o *ResumeResumeForApplicant) SetMetroNil() {
	o.Metro.Set(nil)
}

// UnsetMetro ensures that no value is present for Metro, not even an explicit nil
func (o *ResumeResumeForApplicant) UnsetMetro() {
	o.Metro.Unset()
}

// GetPaidServices returns the PaidServices field value
func (o *ResumeResumeForApplicant) GetPaidServices() []ResumeObjectsPaidServices {
	if o == nil {
		var ret []ResumeObjectsPaidServices
		return ret
	}

	return o.PaidServices
}

// GetPaidServicesOk returns a tuple with the PaidServices field value
// and a boolean to check if the value has been set.
func (o *ResumeResumeForApplicant) GetPaidServicesOk() ([]ResumeObjectsPaidServices, bool) {
	if o == nil {
		return nil, false
	}
	return o.PaidServices, true
}

// SetPaidServices sets field value
func (o *ResumeResumeForApplicant) SetPaidServices(v []ResumeObjectsPaidServices) {
	o.PaidServices = v
}

// GetProfessionalRoles returns the ProfessionalRoles field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ResumeResumeForApplicant) GetProfessionalRoles() []IncludesIdName {
	if o == nil {
		var ret []IncludesIdName
		return ret
	}
	return o.ProfessionalRoles
}

// GetProfessionalRolesOk returns a tuple with the ProfessionalRoles field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ResumeResumeForApplicant) GetProfessionalRolesOk() ([]IncludesIdName, bool) {
	if o == nil || IsNil(o.ProfessionalRoles) {
		return nil, false
	}
	return o.ProfessionalRoles, true
}

// HasProfessionalRoles returns a boolean if a field has been set.
func (o *ResumeResumeForApplicant) HasProfessionalRoles() bool {
	if o != nil && IsNil(o.ProfessionalRoles) {
		return true
	}

	return false
}

// SetProfessionalRoles gets a reference to the given []IncludesIdName and assigns it to the ProfessionalRoles field.
func (o *ResumeResumeForApplicant) SetProfessionalRoles(v []IncludesIdName) {
	o.ProfessionalRoles = v
}

// GetRecommendation returns the Recommendation field value
func (o *ResumeResumeForApplicant) GetRecommendation() []ResumeObjectsRecommendation {
	if o == nil {
		var ret []ResumeObjectsRecommendation
		return ret
	}

	return o.Recommendation
}

// GetRecommendationOk returns a tuple with the Recommendation field value
// and a boolean to check if the value has been set.
func (o *ResumeResumeForApplicant) GetRecommendationOk() ([]ResumeObjectsRecommendation, bool) {
	if o == nil {
		return nil, false
	}
	return o.Recommendation, true
}

// SetRecommendation sets field value
func (o *ResumeResumeForApplicant) SetRecommendation(v []ResumeObjectsRecommendation) {
	o.Recommendation = v
}

// GetRelocation returns the Relocation field value
func (o *ResumeResumeForApplicant) GetRelocation() ResumeResumeFullAllOfRelocation {
	if o == nil {
		var ret ResumeResumeFullAllOfRelocation
		return ret
	}

	return o.Relocation
}

// GetRelocationOk returns a tuple with the Relocation field value
// and a boolean to check if the value has been set.
func (o *ResumeResumeForApplicant) GetRelocationOk() (*ResumeResumeFullAllOfRelocation, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Relocation, true
}

// SetRelocation sets field value
func (o *ResumeResumeForApplicant) SetRelocation(v ResumeResumeFullAllOfRelocation) {
	o.Relocation = v
}

// GetResumeLocale returns the ResumeLocale field value
func (o *ResumeResumeForApplicant) GetResumeLocale() ResumeResumeFullAllOfResumeLocale {
	if o == nil {
		var ret ResumeResumeFullAllOfResumeLocale
		return ret
	}

	return o.ResumeLocale
}

// GetResumeLocaleOk returns a tuple with the ResumeLocale field value
// and a boolean to check if the value has been set.
func (o *ResumeResumeForApplicant) GetResumeLocaleOk() (*ResumeResumeFullAllOfResumeLocale, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ResumeLocale, true
}

// SetResumeLocale sets field value
func (o *ResumeResumeForApplicant) SetResumeLocale(v ResumeResumeFullAllOfResumeLocale) {
	o.ResumeLocale = v
}

// GetSchedule returns the Schedule field value
// Deprecated
func (o *ResumeResumeForApplicant) GetSchedule() ResumeResumeFullAllOfSchedule {
	if o == nil {
		var ret ResumeResumeFullAllOfSchedule
		return ret
	}

	return o.Schedule
}

// GetScheduleOk returns a tuple with the Schedule field value
// and a boolean to check if the value has been set.
// Deprecated
func (o *ResumeResumeForApplicant) GetScheduleOk() (*ResumeResumeFullAllOfSchedule, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Schedule, true
}

// SetSchedule sets field value
// Deprecated
func (o *ResumeResumeForApplicant) SetSchedule(v ResumeResumeFullAllOfSchedule) {
	o.Schedule = v
}

// GetSchedules returns the Schedules field value
func (o *ResumeResumeForApplicant) GetSchedules() []IncludesIdName {
	if o == nil {
		var ret []IncludesIdName
		return ret
	}

	return o.Schedules
}

// GetSchedulesOk returns a tuple with the Schedules field value
// and a boolean to check if the value has been set.
func (o *ResumeResumeForApplicant) GetSchedulesOk() ([]IncludesIdName, bool) {
	if o == nil {
		return nil, false
	}
	return o.Schedules, true
}

// SetSchedules sets field value
func (o *ResumeResumeForApplicant) SetSchedules(v []IncludesIdName) {
	o.Schedules = v
}

// GetSite returns the Site field value
func (o *ResumeResumeForApplicant) GetSite() []ResumeObjectsSite {
	if o == nil {
		var ret []ResumeObjectsSite
		return ret
	}

	return o.Site
}

// GetSiteOk returns a tuple with the Site field value
// and a boolean to check if the value has been set.
func (o *ResumeResumeForApplicant) GetSiteOk() ([]ResumeObjectsSite, bool) {
	if o == nil {
		return nil, false
	}
	return o.Site, true
}

// SetSite sets field value
func (o *ResumeResumeForApplicant) SetSite(v []ResumeObjectsSite) {
	o.Site = v
}

// GetSkillSet returns the SkillSet field value
func (o *ResumeResumeForApplicant) GetSkillSet() []string {
	if o == nil {
		var ret []string
		return ret
	}

	return o.SkillSet
}

// GetSkillSetOk returns a tuple with the SkillSet field value
// and a boolean to check if the value has been set.
func (o *ResumeResumeForApplicant) GetSkillSetOk() ([]string, bool) {
	if o == nil {
		return nil, false
	}
	return o.SkillSet, true
}

// SetSkillSet sets field value
func (o *ResumeResumeForApplicant) SetSkillSet(v []string) {
	o.SkillSet = v
}

// GetSkills returns the Skills field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ResumeResumeForApplicant) GetSkills() string {
	if o == nil || IsNil(o.Skills.Get()) {
		var ret string
		return ret
	}
	return *o.Skills.Get()
}

// GetSkillsOk returns a tuple with the Skills field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ResumeResumeForApplicant) GetSkillsOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Skills.Get(), o.Skills.IsSet()
}

// HasSkills returns a boolean if a field has been set.
func (o *ResumeResumeForApplicant) HasSkills() bool {
	if o != nil && o.Skills.IsSet() {
		return true
	}

	return false
}

// SetSkills gets a reference to the given NullableString and assigns it to the Skills field.
func (o *ResumeResumeForApplicant) SetSkills(v string) {
	o.Skills.Set(&v)
}
// SetSkillsNil sets the value for Skills to be an explicit nil
func (o *ResumeResumeForApplicant) SetSkillsNil() {
	o.Skills.Set(nil)
}

// UnsetSkills ensures that no value is present for Skills, not even an explicit nil
func (o *ResumeResumeForApplicant) UnsetSkills() {
	o.Skills.Unset()
}

// GetTravelTime returns the TravelTime field value
func (o *ResumeResumeForApplicant) GetTravelTime() ResumeResumeFullAllOfTravelTime {
	if o == nil {
		var ret ResumeResumeFullAllOfTravelTime
		return ret
	}

	return o.TravelTime
}

// GetTravelTimeOk returns a tuple with the TravelTime field value
// and a boolean to check if the value has been set.
func (o *ResumeResumeForApplicant) GetTravelTimeOk() (*ResumeResumeFullAllOfTravelTime, bool) {
	if o == nil {
		return nil, false
	}
	return &o.TravelTime, true
}

// SetTravelTime sets field value
func (o *ResumeResumeForApplicant) SetTravelTime(v ResumeResumeFullAllOfTravelTime) {
	o.TravelTime = v
}

// GetWorkTicket returns the WorkTicket field value
func (o *ResumeResumeForApplicant) GetWorkTicket() []IncludesIdNameUrl {
	if o == nil {
		var ret []IncludesIdNameUrl
		return ret
	}

	return o.WorkTicket
}

// GetWorkTicketOk returns a tuple with the WorkTicket field value
// and a boolean to check if the value has been set.
func (o *ResumeResumeForApplicant) GetWorkTicketOk() ([]IncludesIdNameUrl, bool) {
	if o == nil {
		return nil, false
	}
	return o.WorkTicket, true
}

// SetWorkTicket sets field value
func (o *ResumeResumeForApplicant) SetWorkTicket(v []IncludesIdNameUrl) {
	o.WorkTicket = v
}

// GetBlocked returns the Blocked field value
func (o *ResumeResumeForApplicant) GetBlocked() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.Blocked
}

// GetBlockedOk returns a tuple with the Blocked field value
// and a boolean to check if the value has been set.
func (o *ResumeResumeForApplicant) GetBlockedOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Blocked, true
}

// SetBlocked sets field value
func (o *ResumeResumeForApplicant) SetBlocked(v bool) {
	o.Blocked = v
}

// GetCanPublishOrUpdate returns the CanPublishOrUpdate field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ResumeResumeForApplicant) GetCanPublishOrUpdate() bool {
	if o == nil || IsNil(o.CanPublishOrUpdate.Get()) {
		var ret bool
		return ret
	}
	return *o.CanPublishOrUpdate.Get()
}

// GetCanPublishOrUpdateOk returns a tuple with the CanPublishOrUpdate field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ResumeResumeForApplicant) GetCanPublishOrUpdateOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return o.CanPublishOrUpdate.Get(), o.CanPublishOrUpdate.IsSet()
}

// HasCanPublishOrUpdate returns a boolean if a field has been set.
func (o *ResumeResumeForApplicant) HasCanPublishOrUpdate() bool {
	if o != nil && o.CanPublishOrUpdate.IsSet() {
		return true
	}

	return false
}

// SetCanPublishOrUpdate gets a reference to the given NullableBool and assigns it to the CanPublishOrUpdate field.
func (o *ResumeResumeForApplicant) SetCanPublishOrUpdate(v bool) {
	o.CanPublishOrUpdate.Set(&v)
}
// SetCanPublishOrUpdateNil sets the value for CanPublishOrUpdate to be an explicit nil
func (o *ResumeResumeForApplicant) SetCanPublishOrUpdateNil() {
	o.CanPublishOrUpdate.Set(nil)
}

// UnsetCanPublishOrUpdate ensures that no value is present for CanPublishOrUpdate, not even an explicit nil
func (o *ResumeResumeForApplicant) UnsetCanPublishOrUpdate() {
	o.CanPublishOrUpdate.Unset()
}

// GetFinished returns the Finished field value
func (o *ResumeResumeForApplicant) GetFinished() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.Finished
}

// GetFinishedOk returns a tuple with the Finished field value
// and a boolean to check if the value has been set.
func (o *ResumeResumeForApplicant) GetFinishedOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Finished, true
}

// SetFinished sets field value
func (o *ResumeResumeForApplicant) SetFinished(v bool) {
	o.Finished = v
}

// GetStatus returns the Status field value
func (o *ResumeResumeForApplicant) GetStatus() IncludesIdName {
	if o == nil {
		var ret IncludesIdName
		return ret
	}

	return o.Status
}

// GetStatusOk returns a tuple with the Status field value
// and a boolean to check if the value has been set.
func (o *ResumeResumeForApplicant) GetStatusOk() (*IncludesIdName, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Status, true
}

// SetStatus sets field value
func (o *ResumeResumeForApplicant) SetStatus(v IncludesIdName) {
	o.Status = v
}

// GetModerationNote returns the ModerationNote field value
func (o *ResumeResumeForApplicant) GetModerationNote() []ResumeObjectsModerationNote {
	if o == nil {
		var ret []ResumeObjectsModerationNote
		return ret
	}

	return o.ModerationNote
}

// GetModerationNoteOk returns a tuple with the ModerationNote field value
// and a boolean to check if the value has been set.
func (o *ResumeResumeForApplicant) GetModerationNoteOk() ([]ResumeObjectsModerationNote, bool) {
	if o == nil {
		return nil, false
	}
	return o.ModerationNote, true
}

// SetModerationNote sets field value
func (o *ResumeResumeForApplicant) SetModerationNote(v []ResumeObjectsModerationNote) {
	o.ModerationNote = v
}

// GetProgress returns the Progress field value
func (o *ResumeResumeForApplicant) GetProgress() ResumeObjectsProgress {
	if o == nil {
		var ret ResumeObjectsProgress
		return ret
	}

	return o.Progress
}

// GetProgressOk returns a tuple with the Progress field value
// and a boolean to check if the value has been set.
func (o *ResumeResumeForApplicant) GetProgressOk() (*ResumeObjectsProgress, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Progress, true
}

// SetProgress sets field value
func (o *ResumeResumeForApplicant) SetProgress(v ResumeObjectsProgress) {
	o.Progress = v
}

// GetPublishUrl returns the PublishUrl field value
func (o *ResumeResumeForApplicant) GetPublishUrl() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.PublishUrl
}

// GetPublishUrlOk returns a tuple with the PublishUrl field value
// and a boolean to check if the value has been set.
func (o *ResumeResumeForApplicant) GetPublishUrlOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.PublishUrl, true
}

// SetPublishUrl sets field value
func (o *ResumeResumeForApplicant) SetPublishUrl(v string) {
	o.PublishUrl = v
}

// GetAccess returns the Access field value
func (o *ResumeResumeForApplicant) GetAccess() ResumeObjectsAccess {
	if o == nil {
		var ret ResumeObjectsAccess
		return ret
	}

	return o.Access
}

// GetAccessOk returns a tuple with the Access field value
// and a boolean to check if the value has been set.
func (o *ResumeResumeForApplicant) GetAccessOk() (*ResumeObjectsAccess, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Access, true
}

// SetAccess sets field value
func (o *ResumeResumeForApplicant) SetAccess(v ResumeObjectsAccess) {
	o.Access = v
}

// GetActions returns the Actions field value
func (o *ResumeResumeForApplicant) GetActions() ResumeObjectsActionsForOwner {
	if o == nil {
		var ret ResumeObjectsActionsForOwner
		return ret
	}

	return o.Actions
}

// GetActionsOk returns a tuple with the Actions field value
// and a boolean to check if the value has been set.
func (o *ResumeResumeForApplicant) GetActionsOk() (*ResumeObjectsActionsForOwner, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Actions, true
}

// SetActions sets field value
func (o *ResumeResumeForApplicant) SetActions(v ResumeObjectsActionsForOwner) {
	o.Actions = v
}

// GetNewViews returns the NewViews field value
func (o *ResumeResumeForApplicant) GetNewViews() float32 {
	if o == nil {
		var ret float32
		return ret
	}

	return o.NewViews
}

// GetNewViewsOk returns a tuple with the NewViews field value
// and a boolean to check if the value has been set.
func (o *ResumeResumeForApplicant) GetNewViewsOk() (*float32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.NewViews, true
}

// SetNewViews sets field value
func (o *ResumeResumeForApplicant) SetNewViews(v float32) {
	o.NewViews = v
}

// GetNextPublishAt returns the NextPublishAt field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ResumeResumeForApplicant) GetNextPublishAt() string {
	if o == nil || IsNil(o.NextPublishAt.Get()) {
		var ret string
		return ret
	}
	return *o.NextPublishAt.Get()
}

// GetNextPublishAtOk returns a tuple with the NextPublishAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ResumeResumeForApplicant) GetNextPublishAtOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.NextPublishAt.Get(), o.NextPublishAt.IsSet()
}

// HasNextPublishAt returns a boolean if a field has been set.
func (o *ResumeResumeForApplicant) HasNextPublishAt() bool {
	if o != nil && o.NextPublishAt.IsSet() {
		return true
	}

	return false
}

// SetNextPublishAt gets a reference to the given NullableString and assigns it to the NextPublishAt field.
func (o *ResumeResumeForApplicant) SetNextPublishAt(v string) {
	o.NextPublishAt.Set(&v)
}
// SetNextPublishAtNil sets the value for NextPublishAt to be an explicit nil
func (o *ResumeResumeForApplicant) SetNextPublishAtNil() {
	o.NextPublishAt.Set(nil)
}

// UnsetNextPublishAt ensures that no value is present for NextPublishAt, not even an explicit nil
func (o *ResumeResumeForApplicant) UnsetNextPublishAt() {
	o.NextPublishAt.Unset()
}

// GetTotalViews returns the TotalViews field value
func (o *ResumeResumeForApplicant) GetTotalViews() float32 {
	if o == nil {
		var ret float32
		return ret
	}

	return o.TotalViews
}

// GetTotalViewsOk returns a tuple with the TotalViews field value
// and a boolean to check if the value has been set.
func (o *ResumeResumeForApplicant) GetTotalViewsOk() (*float32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.TotalViews, true
}

// SetTotalViews sets field value
func (o *ResumeResumeForApplicant) SetTotalViews(v float32) {
	o.TotalViews = v
}

// GetViewsUrl returns the ViewsUrl field value
func (o *ResumeResumeForApplicant) GetViewsUrl() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ViewsUrl
}

// GetViewsUrlOk returns a tuple with the ViewsUrl field value
// and a boolean to check if the value has been set.
func (o *ResumeResumeForApplicant) GetViewsUrlOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ViewsUrl, true
}

// SetViewsUrl sets field value
func (o *ResumeResumeForApplicant) SetViewsUrl(v string) {
	o.ViewsUrl = v
}

// GetPhoto returns the Photo field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ResumeResumeForApplicant) GetPhoto() ResumeObjectsPhoto {
	if o == nil || IsNil(o.Photo.Get()) {
		var ret ResumeObjectsPhoto
		return ret
	}
	return *o.Photo.Get()
}

// GetPhotoOk returns a tuple with the Photo field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ResumeResumeForApplicant) GetPhotoOk() (*ResumeObjectsPhoto, bool) {
	if o == nil {
		return nil, false
	}
	return o.Photo.Get(), o.Photo.IsSet()
}

// HasPhoto returns a boolean if a field has been set.
func (o *ResumeResumeForApplicant) HasPhoto() bool {
	if o != nil && o.Photo.IsSet() {
		return true
	}

	return false
}

// SetPhoto gets a reference to the given NullableResumeObjectsPhoto and assigns it to the Photo field.
func (o *ResumeResumeForApplicant) SetPhoto(v ResumeObjectsPhoto) {
	o.Photo.Set(&v)
}
// SetPhotoNil sets the value for Photo to be an explicit nil
func (o *ResumeResumeForApplicant) SetPhotoNil() {
	o.Photo.Set(nil)
}

// UnsetPhoto ensures that no value is present for Photo, not even an explicit nil
func (o *ResumeResumeForApplicant) UnsetPhoto() {
	o.Photo.Unset()
}

// GetPortfolio returns the Portfolio field value
func (o *ResumeResumeForApplicant) GetPortfolio() []ResumeObjectsPortfolio {
	if o == nil {
		var ret []ResumeObjectsPortfolio
		return ret
	}

	return o.Portfolio
}

// GetPortfolioOk returns a tuple with the Portfolio field value
// and a boolean to check if the value has been set.
func (o *ResumeResumeForApplicant) GetPortfolioOk() ([]ResumeObjectsPortfolio, bool) {
	if o == nil {
		return nil, false
	}
	return o.Portfolio, true
}

// SetPortfolio sets field value
func (o *ResumeResumeForApplicant) SetPortfolio(v []ResumeObjectsPortfolio) {
	o.Portfolio = v
}

func (o ResumeResumeForApplicant) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ResumeResumeForApplicant) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["alternate_url"] = o.AlternateUrl
	toSerialize["id"] = o.Id
	toSerialize["title"] = o.Title.Get()
	if o.Age.IsSet() {
		toSerialize["age"] = o.Age.Get()
	}
	if o.Area.IsSet() {
		toSerialize["area"] = o.Area.Get()
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
	if o.Gender.IsSet() {
		toSerialize["gender"] = o.Gender.Get()
	}
	toSerialize["hidden_fields"] = o.HiddenFields
	if o.LastName.IsSet() {
		toSerialize["last_name"] = o.LastName.Get()
	}
	if !IsNil(o.Marked) {
		toSerialize["marked"] = o.Marked
	}
	if o.MiddleName.IsSet() {
		toSerialize["middle_name"] = o.MiddleName.Get()
	}
	if !IsNil(o.Platform) {
		toSerialize["platform"] = o.Platform
	}
	if o.Salary.IsSet() {
		toSerialize["salary"] = o.Salary.Get()
	}
	if o.TotalExperience.IsSet() {
		toSerialize["total_experience"] = o.TotalExperience.Get()
	}
	toSerialize["updated_at"] = o.UpdatedAt
	if o.BirthDate.IsSet() {
		toSerialize["birth_date"] = o.BirthDate.Get()
	}
	toSerialize["business_trip_readiness"] = o.BusinessTripReadiness
	toSerialize["citizenship"] = o.Citizenship
	toSerialize["contact"] = o.Contact
	if o.Creds.IsSet() {
		toSerialize["creds"] = o.Creds.Get()
	}
	toSerialize["driver_license_types"] = o.DriverLicenseTypes
	if !IsNil(o.Employment) {
		toSerialize["employment"] = o.Employment
	}
	toSerialize["employments"] = o.Employments
	if o.HasVehicle.IsSet() {
		toSerialize["has_vehicle"] = o.HasVehicle.Get()
	}
	toSerialize["language"] = o.Language
	if o.Metro.IsSet() {
		toSerialize["metro"] = o.Metro.Get()
	}
	toSerialize["paid_services"] = o.PaidServices
	if o.ProfessionalRoles != nil {
		toSerialize["professional_roles"] = o.ProfessionalRoles
	}
	toSerialize["recommendation"] = o.Recommendation
	toSerialize["relocation"] = o.Relocation
	toSerialize["resume_locale"] = o.ResumeLocale
	toSerialize["schedule"] = o.Schedule
	toSerialize["schedules"] = o.Schedules
	toSerialize["site"] = o.Site
	toSerialize["skill_set"] = o.SkillSet
	if o.Skills.IsSet() {
		toSerialize["skills"] = o.Skills.Get()
	}
	toSerialize["travel_time"] = o.TravelTime
	toSerialize["work_ticket"] = o.WorkTicket
	toSerialize["blocked"] = o.Blocked
	if o.CanPublishOrUpdate.IsSet() {
		toSerialize["can_publish_or_update"] = o.CanPublishOrUpdate.Get()
	}
	toSerialize["finished"] = o.Finished
	toSerialize["status"] = o.Status
	toSerialize["moderation_note"] = o.ModerationNote
	toSerialize["progress"] = o.Progress
	toSerialize["publish_url"] = o.PublishUrl
	toSerialize["access"] = o.Access
	toSerialize["actions"] = o.Actions
	toSerialize["new_views"] = o.NewViews
	if o.NextPublishAt.IsSet() {
		toSerialize["next_publish_at"] = o.NextPublishAt.Get()
	}
	toSerialize["total_views"] = o.TotalViews
	toSerialize["views_url"] = o.ViewsUrl
	if o.Photo.IsSet() {
		toSerialize["photo"] = o.Photo.Get()
	}
	toSerialize["portfolio"] = o.Portfolio
	return toSerialize, nil
}

func (o *ResumeResumeForApplicant) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"alternate_url",
		"id",
		"title",
		"certificate",
		"created_at",
		"download",
		"education",
		"experience",
		"hidden_fields",
		"updated_at",
		"business_trip_readiness",
		"citizenship",
		"contact",
		"driver_license_types",
		"employments",
		"language",
		"paid_services",
		"recommendation",
		"relocation",
		"resume_locale",
		"schedule",
		"schedules",
		"site",
		"skill_set",
		"travel_time",
		"work_ticket",
		"blocked",
		"finished",
		"status",
		"moderation_note",
		"progress",
		"publish_url",
		"access",
		"actions",
		"new_views",
		"total_views",
		"views_url",
		"portfolio",
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

	varResumeResumeForApplicant := _ResumeResumeForApplicant{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varResumeResumeForApplicant)

	if err != nil {
		return err
	}

	*o = ResumeResumeForApplicant(varResumeResumeForApplicant)

	return err
}

type NullableResumeResumeForApplicant struct {
	value *ResumeResumeForApplicant
	isSet bool
}

func (v NullableResumeResumeForApplicant) Get() *ResumeResumeForApplicant {
	return v.value
}

func (v *NullableResumeResumeForApplicant) Set(val *ResumeResumeForApplicant) {
	v.value = val
	v.isSet = true
}

func (v NullableResumeResumeForApplicant) IsSet() bool {
	return v.isSet
}

func (v *NullableResumeResumeForApplicant) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableResumeResumeForApplicant(val *ResumeResumeForApplicant) *NullableResumeResumeForApplicant {
	return &NullableResumeResumeForApplicant{value: val, isSet: true}
}

func (v NullableResumeResumeForApplicant) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableResumeResumeForApplicant) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


