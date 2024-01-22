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
)

// checks if the ResumeEditResumeRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ResumeEditResumeRequest{}

// ResumeEditResumeRequest Тело запроса при редактировании резюме
type ResumeEditResumeRequest struct {
	Access *ResumeObjectsAccess `json:"access,omitempty"`
	// День рождения (в формате `ГГГГ-ММ-ДД`)
	BirthDate NullableString `json:"birth_date,omitempty"`
	BusinessTripReadiness *IncludesId `json:"business_trip_readiness,omitempty"`
	// Список сертификатов соискателя
	Certificate []ResumeObjectsCertificate `json:"certificate,omitempty"`
	// Список категорий водительских прав соискателя
	DriverLicenseTypes []ResumeObjectsDriverLicenseTypes `json:"driver_license_types,omitempty"`
	// Список подходящих соискателю типов занятостей. Элементы справочника [employment](#tag/Obshie-spravochniki/operation/get-dictionaries)
	Employments []IncludesIdName `json:"employments,omitempty"`
	// Имя
	FirstName NullableString `json:"first_name,omitempty"`
	// Наличие личного автомобиля у соискателя
	HasVehicle NullableBool `json:"has_vehicle,omitempty"`
	// Документация [Список скрытых полей](https://github.com/hhru/api/blob/master/docs/employer_resumes.md#hidden-fields). Возможные значения элементов приведены в поле `resume_hidden_fields` [справочника полей](#tag/Obshie-spravochniki/operation/get-dictionaries)
	HiddenFields []IncludesIdName `json:"hidden_fields,omitempty"`
	// Фамилия
	LastName NullableString `json:"last_name,omitempty"`
	Metro *IncludesId `json:"metro,omitempty"`
	// Отчество
	MiddleName NullableString `json:"middle_name,omitempty"`
	Photo NullableResumeObjectsPhoto `json:"photo,omitempty"`
	// Список изображений в портфолио пользователя
	Portfolio []ResumeObjectsPortfolio `json:"portfolio,omitempty"`
	// Массив объектов профролей. Элемент справочника [professional_roles](#tag/Obshie-spravochniki/operation/get-professional-roles-dictionary)
	ProfessionalRoles []IncludesId `json:"professional_roles,omitempty"`
	// Список рекомендаций
	Recommendation []ResumeObjectsRecommendation `json:"recommendation,omitempty"`
	Relocation *ResumeObjectsRelocationPublic `json:"relocation,omitempty"`
	ResumeLocale *IncludesIdName `json:"resume_locale,omitempty"`
	Salary *ResumeObjectsSalaryAddEdit `json:"salary,omitempty"`
	// Список подходящих соискателю графиков работы. Элементы справочника [schedule](#tag/Obshie-spravochniki/operation/get-dictionaries)
	Schedules []IncludesIdName `json:"schedules,omitempty"`
	// Профили в соц. сетях и других сервисах
	Site []ResumeObjectsSite `json:"site,omitempty"`
	// Ключевые навыки (список уникальных строк)
	SkillSet []string `json:"skill_set,omitempty"`
	// Дополнительная информация, описание навыков в свободной форме
	Skills NullableString `json:"skills,omitempty"`
	// Желаемая должность
	Title NullableString `json:"title,omitempty"`
	TotalExperience NullableResumeObjectsTotalExperience `json:"total_experience,omitempty"`
	TravelTime *IncludesId `json:"travel_time,omitempty"`
	// Список регионов, в который соискатель имеет разрешение на работу. Элементы [справочника регионов](#tag/Obshie-spravochniki/operation/get-areas) 
	WorkTicket []IncludesId `json:"work_ticket,omitempty"`
	Area *Id `json:"area,omitempty"`
	// Список гражданств соискателя. Элементы [справочника регионов](#tag/Obshie-spravochniki/operation/get-areas)
	Citizenship []IncludesId `json:"citizenship,omitempty"`
	// Список контактов соискателя.  При заполнении контактов в резюме необходимо учитывать следующие условия:  * В резюме обязательно должен быть указан e-mail. Он может быть только один. * В резюме должен быть указан хотя бы один телефон, причём можно указывать только один телефон каждого типа. * Комментарий можно указывать только для телефонов, для e-mail комментарий не сохранится * Обязательно указать либо телефон полностью в поле `formatted`, либо все три части телефона по отдельности в трёх полях: `country`, `city` и `number`. Если указано и то, и то, используются данные из трёх полей. В поле `formatted` допустимо использовать пробелы, скобки и дефисы. В остальных полях допустимы только цифры 
	Contact []IncludesContact `json:"contact,omitempty"`
	Education NullableResumeObjectsEducation `json:"education,omitempty"`
	// Опыт работы
	Experience []ResumeObjectsExperience `json:"experience,omitempty"`
	Gender *Id `json:"gender,omitempty"`
	// Список языков, которыми владеет соискатель. Элементы справочника [languages](#tag/Obshie-spravochniki/operation/get-languages)
	Language []IncludesLanguageLevel `json:"language,omitempty"`
}

// NewResumeEditResumeRequest instantiates a new ResumeEditResumeRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewResumeEditResumeRequest() *ResumeEditResumeRequest {
	this := ResumeEditResumeRequest{}
	return &this
}

// NewResumeEditResumeRequestWithDefaults instantiates a new ResumeEditResumeRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewResumeEditResumeRequestWithDefaults() *ResumeEditResumeRequest {
	this := ResumeEditResumeRequest{}
	return &this
}

// GetAccess returns the Access field value if set, zero value otherwise.
func (o *ResumeEditResumeRequest) GetAccess() ResumeObjectsAccess {
	if o == nil || IsNil(o.Access) {
		var ret ResumeObjectsAccess
		return ret
	}
	return *o.Access
}

// GetAccessOk returns a tuple with the Access field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ResumeEditResumeRequest) GetAccessOk() (*ResumeObjectsAccess, bool) {
	if o == nil || IsNil(o.Access) {
		return nil, false
	}
	return o.Access, true
}

// HasAccess returns a boolean if a field has been set.
func (o *ResumeEditResumeRequest) HasAccess() bool {
	if o != nil && !IsNil(o.Access) {
		return true
	}

	return false
}

// SetAccess gets a reference to the given ResumeObjectsAccess and assigns it to the Access field.
func (o *ResumeEditResumeRequest) SetAccess(v ResumeObjectsAccess) {
	o.Access = &v
}

// GetBirthDate returns the BirthDate field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ResumeEditResumeRequest) GetBirthDate() string {
	if o == nil || IsNil(o.BirthDate.Get()) {
		var ret string
		return ret
	}
	return *o.BirthDate.Get()
}

// GetBirthDateOk returns a tuple with the BirthDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ResumeEditResumeRequest) GetBirthDateOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.BirthDate.Get(), o.BirthDate.IsSet()
}

// HasBirthDate returns a boolean if a field has been set.
func (o *ResumeEditResumeRequest) HasBirthDate() bool {
	if o != nil && o.BirthDate.IsSet() {
		return true
	}

	return false
}

// SetBirthDate gets a reference to the given NullableString and assigns it to the BirthDate field.
func (o *ResumeEditResumeRequest) SetBirthDate(v string) {
	o.BirthDate.Set(&v)
}
// SetBirthDateNil sets the value for BirthDate to be an explicit nil
func (o *ResumeEditResumeRequest) SetBirthDateNil() {
	o.BirthDate.Set(nil)
}

// UnsetBirthDate ensures that no value is present for BirthDate, not even an explicit nil
func (o *ResumeEditResumeRequest) UnsetBirthDate() {
	o.BirthDate.Unset()
}

// GetBusinessTripReadiness returns the BusinessTripReadiness field value if set, zero value otherwise.
func (o *ResumeEditResumeRequest) GetBusinessTripReadiness() IncludesId {
	if o == nil || IsNil(o.BusinessTripReadiness) {
		var ret IncludesId
		return ret
	}
	return *o.BusinessTripReadiness
}

// GetBusinessTripReadinessOk returns a tuple with the BusinessTripReadiness field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ResumeEditResumeRequest) GetBusinessTripReadinessOk() (*IncludesId, bool) {
	if o == nil || IsNil(o.BusinessTripReadiness) {
		return nil, false
	}
	return o.BusinessTripReadiness, true
}

// HasBusinessTripReadiness returns a boolean if a field has been set.
func (o *ResumeEditResumeRequest) HasBusinessTripReadiness() bool {
	if o != nil && !IsNil(o.BusinessTripReadiness) {
		return true
	}

	return false
}

// SetBusinessTripReadiness gets a reference to the given IncludesId and assigns it to the BusinessTripReadiness field.
func (o *ResumeEditResumeRequest) SetBusinessTripReadiness(v IncludesId) {
	o.BusinessTripReadiness = &v
}

// GetCertificate returns the Certificate field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ResumeEditResumeRequest) GetCertificate() []ResumeObjectsCertificate {
	if o == nil {
		var ret []ResumeObjectsCertificate
		return ret
	}
	return o.Certificate
}

// GetCertificateOk returns a tuple with the Certificate field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ResumeEditResumeRequest) GetCertificateOk() ([]ResumeObjectsCertificate, bool) {
	if o == nil || IsNil(o.Certificate) {
		return nil, false
	}
	return o.Certificate, true
}

// HasCertificate returns a boolean if a field has been set.
func (o *ResumeEditResumeRequest) HasCertificate() bool {
	if o != nil && IsNil(o.Certificate) {
		return true
	}

	return false
}

// SetCertificate gets a reference to the given []ResumeObjectsCertificate and assigns it to the Certificate field.
func (o *ResumeEditResumeRequest) SetCertificate(v []ResumeObjectsCertificate) {
	o.Certificate = v
}

// GetDriverLicenseTypes returns the DriverLicenseTypes field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ResumeEditResumeRequest) GetDriverLicenseTypes() []ResumeObjectsDriverLicenseTypes {
	if o == nil {
		var ret []ResumeObjectsDriverLicenseTypes
		return ret
	}
	return o.DriverLicenseTypes
}

// GetDriverLicenseTypesOk returns a tuple with the DriverLicenseTypes field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ResumeEditResumeRequest) GetDriverLicenseTypesOk() ([]ResumeObjectsDriverLicenseTypes, bool) {
	if o == nil || IsNil(o.DriverLicenseTypes) {
		return nil, false
	}
	return o.DriverLicenseTypes, true
}

// HasDriverLicenseTypes returns a boolean if a field has been set.
func (o *ResumeEditResumeRequest) HasDriverLicenseTypes() bool {
	if o != nil && IsNil(o.DriverLicenseTypes) {
		return true
	}

	return false
}

// SetDriverLicenseTypes gets a reference to the given []ResumeObjectsDriverLicenseTypes and assigns it to the DriverLicenseTypes field.
func (o *ResumeEditResumeRequest) SetDriverLicenseTypes(v []ResumeObjectsDriverLicenseTypes) {
	o.DriverLicenseTypes = v
}

// GetEmployments returns the Employments field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ResumeEditResumeRequest) GetEmployments() []IncludesIdName {
	if o == nil {
		var ret []IncludesIdName
		return ret
	}
	return o.Employments
}

// GetEmploymentsOk returns a tuple with the Employments field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ResumeEditResumeRequest) GetEmploymentsOk() ([]IncludesIdName, bool) {
	if o == nil || IsNil(o.Employments) {
		return nil, false
	}
	return o.Employments, true
}

// HasEmployments returns a boolean if a field has been set.
func (o *ResumeEditResumeRequest) HasEmployments() bool {
	if o != nil && IsNil(o.Employments) {
		return true
	}

	return false
}

// SetEmployments gets a reference to the given []IncludesIdName and assigns it to the Employments field.
func (o *ResumeEditResumeRequest) SetEmployments(v []IncludesIdName) {
	o.Employments = v
}

// GetFirstName returns the FirstName field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ResumeEditResumeRequest) GetFirstName() string {
	if o == nil || IsNil(o.FirstName.Get()) {
		var ret string
		return ret
	}
	return *o.FirstName.Get()
}

// GetFirstNameOk returns a tuple with the FirstName field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ResumeEditResumeRequest) GetFirstNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.FirstName.Get(), o.FirstName.IsSet()
}

// HasFirstName returns a boolean if a field has been set.
func (o *ResumeEditResumeRequest) HasFirstName() bool {
	if o != nil && o.FirstName.IsSet() {
		return true
	}

	return false
}

// SetFirstName gets a reference to the given NullableString and assigns it to the FirstName field.
func (o *ResumeEditResumeRequest) SetFirstName(v string) {
	o.FirstName.Set(&v)
}
// SetFirstNameNil sets the value for FirstName to be an explicit nil
func (o *ResumeEditResumeRequest) SetFirstNameNil() {
	o.FirstName.Set(nil)
}

// UnsetFirstName ensures that no value is present for FirstName, not even an explicit nil
func (o *ResumeEditResumeRequest) UnsetFirstName() {
	o.FirstName.Unset()
}

// GetHasVehicle returns the HasVehicle field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ResumeEditResumeRequest) GetHasVehicle() bool {
	if o == nil || IsNil(o.HasVehicle.Get()) {
		var ret bool
		return ret
	}
	return *o.HasVehicle.Get()
}

// GetHasVehicleOk returns a tuple with the HasVehicle field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ResumeEditResumeRequest) GetHasVehicleOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return o.HasVehicle.Get(), o.HasVehicle.IsSet()
}

// HasHasVehicle returns a boolean if a field has been set.
func (o *ResumeEditResumeRequest) HasHasVehicle() bool {
	if o != nil && o.HasVehicle.IsSet() {
		return true
	}

	return false
}

// SetHasVehicle gets a reference to the given NullableBool and assigns it to the HasVehicle field.
func (o *ResumeEditResumeRequest) SetHasVehicle(v bool) {
	o.HasVehicle.Set(&v)
}
// SetHasVehicleNil sets the value for HasVehicle to be an explicit nil
func (o *ResumeEditResumeRequest) SetHasVehicleNil() {
	o.HasVehicle.Set(nil)
}

// UnsetHasVehicle ensures that no value is present for HasVehicle, not even an explicit nil
func (o *ResumeEditResumeRequest) UnsetHasVehicle() {
	o.HasVehicle.Unset()
}

// GetHiddenFields returns the HiddenFields field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ResumeEditResumeRequest) GetHiddenFields() []IncludesIdName {
	if o == nil {
		var ret []IncludesIdName
		return ret
	}
	return o.HiddenFields
}

// GetHiddenFieldsOk returns a tuple with the HiddenFields field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ResumeEditResumeRequest) GetHiddenFieldsOk() ([]IncludesIdName, bool) {
	if o == nil || IsNil(o.HiddenFields) {
		return nil, false
	}
	return o.HiddenFields, true
}

// HasHiddenFields returns a boolean if a field has been set.
func (o *ResumeEditResumeRequest) HasHiddenFields() bool {
	if o != nil && IsNil(o.HiddenFields) {
		return true
	}

	return false
}

// SetHiddenFields gets a reference to the given []IncludesIdName and assigns it to the HiddenFields field.
func (o *ResumeEditResumeRequest) SetHiddenFields(v []IncludesIdName) {
	o.HiddenFields = v
}

// GetLastName returns the LastName field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ResumeEditResumeRequest) GetLastName() string {
	if o == nil || IsNil(o.LastName.Get()) {
		var ret string
		return ret
	}
	return *o.LastName.Get()
}

// GetLastNameOk returns a tuple with the LastName field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ResumeEditResumeRequest) GetLastNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.LastName.Get(), o.LastName.IsSet()
}

// HasLastName returns a boolean if a field has been set.
func (o *ResumeEditResumeRequest) HasLastName() bool {
	if o != nil && o.LastName.IsSet() {
		return true
	}

	return false
}

// SetLastName gets a reference to the given NullableString and assigns it to the LastName field.
func (o *ResumeEditResumeRequest) SetLastName(v string) {
	o.LastName.Set(&v)
}
// SetLastNameNil sets the value for LastName to be an explicit nil
func (o *ResumeEditResumeRequest) SetLastNameNil() {
	o.LastName.Set(nil)
}

// UnsetLastName ensures that no value is present for LastName, not even an explicit nil
func (o *ResumeEditResumeRequest) UnsetLastName() {
	o.LastName.Unset()
}

// GetMetro returns the Metro field value if set, zero value otherwise.
func (o *ResumeEditResumeRequest) GetMetro() IncludesId {
	if o == nil || IsNil(o.Metro) {
		var ret IncludesId
		return ret
	}
	return *o.Metro
}

// GetMetroOk returns a tuple with the Metro field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ResumeEditResumeRequest) GetMetroOk() (*IncludesId, bool) {
	if o == nil || IsNil(o.Metro) {
		return nil, false
	}
	return o.Metro, true
}

// HasMetro returns a boolean if a field has been set.
func (o *ResumeEditResumeRequest) HasMetro() bool {
	if o != nil && !IsNil(o.Metro) {
		return true
	}

	return false
}

// SetMetro gets a reference to the given IncludesId and assigns it to the Metro field.
func (o *ResumeEditResumeRequest) SetMetro(v IncludesId) {
	o.Metro = &v
}

// GetMiddleName returns the MiddleName field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ResumeEditResumeRequest) GetMiddleName() string {
	if o == nil || IsNil(o.MiddleName.Get()) {
		var ret string
		return ret
	}
	return *o.MiddleName.Get()
}

// GetMiddleNameOk returns a tuple with the MiddleName field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ResumeEditResumeRequest) GetMiddleNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.MiddleName.Get(), o.MiddleName.IsSet()
}

// HasMiddleName returns a boolean if a field has been set.
func (o *ResumeEditResumeRequest) HasMiddleName() bool {
	if o != nil && o.MiddleName.IsSet() {
		return true
	}

	return false
}

// SetMiddleName gets a reference to the given NullableString and assigns it to the MiddleName field.
func (o *ResumeEditResumeRequest) SetMiddleName(v string) {
	o.MiddleName.Set(&v)
}
// SetMiddleNameNil sets the value for MiddleName to be an explicit nil
func (o *ResumeEditResumeRequest) SetMiddleNameNil() {
	o.MiddleName.Set(nil)
}

// UnsetMiddleName ensures that no value is present for MiddleName, not even an explicit nil
func (o *ResumeEditResumeRequest) UnsetMiddleName() {
	o.MiddleName.Unset()
}

// GetPhoto returns the Photo field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ResumeEditResumeRequest) GetPhoto() ResumeObjectsPhoto {
	if o == nil || IsNil(o.Photo.Get()) {
		var ret ResumeObjectsPhoto
		return ret
	}
	return *o.Photo.Get()
}

// GetPhotoOk returns a tuple with the Photo field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ResumeEditResumeRequest) GetPhotoOk() (*ResumeObjectsPhoto, bool) {
	if o == nil {
		return nil, false
	}
	return o.Photo.Get(), o.Photo.IsSet()
}

// HasPhoto returns a boolean if a field has been set.
func (o *ResumeEditResumeRequest) HasPhoto() bool {
	if o != nil && o.Photo.IsSet() {
		return true
	}

	return false
}

// SetPhoto gets a reference to the given NullableResumeObjectsPhoto and assigns it to the Photo field.
func (o *ResumeEditResumeRequest) SetPhoto(v ResumeObjectsPhoto) {
	o.Photo.Set(&v)
}
// SetPhotoNil sets the value for Photo to be an explicit nil
func (o *ResumeEditResumeRequest) SetPhotoNil() {
	o.Photo.Set(nil)
}

// UnsetPhoto ensures that no value is present for Photo, not even an explicit nil
func (o *ResumeEditResumeRequest) UnsetPhoto() {
	o.Photo.Unset()
}

// GetPortfolio returns the Portfolio field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ResumeEditResumeRequest) GetPortfolio() []ResumeObjectsPortfolio {
	if o == nil {
		var ret []ResumeObjectsPortfolio
		return ret
	}
	return o.Portfolio
}

// GetPortfolioOk returns a tuple with the Portfolio field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ResumeEditResumeRequest) GetPortfolioOk() ([]ResumeObjectsPortfolio, bool) {
	if o == nil || IsNil(o.Portfolio) {
		return nil, false
	}
	return o.Portfolio, true
}

// HasPortfolio returns a boolean if a field has been set.
func (o *ResumeEditResumeRequest) HasPortfolio() bool {
	if o != nil && IsNil(o.Portfolio) {
		return true
	}

	return false
}

// SetPortfolio gets a reference to the given []ResumeObjectsPortfolio and assigns it to the Portfolio field.
func (o *ResumeEditResumeRequest) SetPortfolio(v []ResumeObjectsPortfolio) {
	o.Portfolio = v
}

// GetProfessionalRoles returns the ProfessionalRoles field value if set, zero value otherwise.
func (o *ResumeEditResumeRequest) GetProfessionalRoles() []IncludesId {
	if o == nil || IsNil(o.ProfessionalRoles) {
		var ret []IncludesId
		return ret
	}
	return o.ProfessionalRoles
}

// GetProfessionalRolesOk returns a tuple with the ProfessionalRoles field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ResumeEditResumeRequest) GetProfessionalRolesOk() ([]IncludesId, bool) {
	if o == nil || IsNil(o.ProfessionalRoles) {
		return nil, false
	}
	return o.ProfessionalRoles, true
}

// HasProfessionalRoles returns a boolean if a field has been set.
func (o *ResumeEditResumeRequest) HasProfessionalRoles() bool {
	if o != nil && !IsNil(o.ProfessionalRoles) {
		return true
	}

	return false
}

// SetProfessionalRoles gets a reference to the given []IncludesId and assigns it to the ProfessionalRoles field.
func (o *ResumeEditResumeRequest) SetProfessionalRoles(v []IncludesId) {
	o.ProfessionalRoles = v
}

// GetRecommendation returns the Recommendation field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ResumeEditResumeRequest) GetRecommendation() []ResumeObjectsRecommendation {
	if o == nil {
		var ret []ResumeObjectsRecommendation
		return ret
	}
	return o.Recommendation
}

// GetRecommendationOk returns a tuple with the Recommendation field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ResumeEditResumeRequest) GetRecommendationOk() ([]ResumeObjectsRecommendation, bool) {
	if o == nil || IsNil(o.Recommendation) {
		return nil, false
	}
	return o.Recommendation, true
}

// HasRecommendation returns a boolean if a field has been set.
func (o *ResumeEditResumeRequest) HasRecommendation() bool {
	if o != nil && IsNil(o.Recommendation) {
		return true
	}

	return false
}

// SetRecommendation gets a reference to the given []ResumeObjectsRecommendation and assigns it to the Recommendation field.
func (o *ResumeEditResumeRequest) SetRecommendation(v []ResumeObjectsRecommendation) {
	o.Recommendation = v
}

// GetRelocation returns the Relocation field value if set, zero value otherwise.
func (o *ResumeEditResumeRequest) GetRelocation() ResumeObjectsRelocationPublic {
	if o == nil || IsNil(o.Relocation) {
		var ret ResumeObjectsRelocationPublic
		return ret
	}
	return *o.Relocation
}

// GetRelocationOk returns a tuple with the Relocation field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ResumeEditResumeRequest) GetRelocationOk() (*ResumeObjectsRelocationPublic, bool) {
	if o == nil || IsNil(o.Relocation) {
		return nil, false
	}
	return o.Relocation, true
}

// HasRelocation returns a boolean if a field has been set.
func (o *ResumeEditResumeRequest) HasRelocation() bool {
	if o != nil && !IsNil(o.Relocation) {
		return true
	}

	return false
}

// SetRelocation gets a reference to the given ResumeObjectsRelocationPublic and assigns it to the Relocation field.
func (o *ResumeEditResumeRequest) SetRelocation(v ResumeObjectsRelocationPublic) {
	o.Relocation = &v
}

// GetResumeLocale returns the ResumeLocale field value if set, zero value otherwise.
func (o *ResumeEditResumeRequest) GetResumeLocale() IncludesIdName {
	if o == nil || IsNil(o.ResumeLocale) {
		var ret IncludesIdName
		return ret
	}
	return *o.ResumeLocale
}

// GetResumeLocaleOk returns a tuple with the ResumeLocale field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ResumeEditResumeRequest) GetResumeLocaleOk() (*IncludesIdName, bool) {
	if o == nil || IsNil(o.ResumeLocale) {
		return nil, false
	}
	return o.ResumeLocale, true
}

// HasResumeLocale returns a boolean if a field has been set.
func (o *ResumeEditResumeRequest) HasResumeLocale() bool {
	if o != nil && !IsNil(o.ResumeLocale) {
		return true
	}

	return false
}

// SetResumeLocale gets a reference to the given IncludesIdName and assigns it to the ResumeLocale field.
func (o *ResumeEditResumeRequest) SetResumeLocale(v IncludesIdName) {
	o.ResumeLocale = &v
}

// GetSalary returns the Salary field value if set, zero value otherwise.
func (o *ResumeEditResumeRequest) GetSalary() ResumeObjectsSalaryAddEdit {
	if o == nil || IsNil(o.Salary) {
		var ret ResumeObjectsSalaryAddEdit
		return ret
	}
	return *o.Salary
}

// GetSalaryOk returns a tuple with the Salary field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ResumeEditResumeRequest) GetSalaryOk() (*ResumeObjectsSalaryAddEdit, bool) {
	if o == nil || IsNil(o.Salary) {
		return nil, false
	}
	return o.Salary, true
}

// HasSalary returns a boolean if a field has been set.
func (o *ResumeEditResumeRequest) HasSalary() bool {
	if o != nil && !IsNil(o.Salary) {
		return true
	}

	return false
}

// SetSalary gets a reference to the given ResumeObjectsSalaryAddEdit and assigns it to the Salary field.
func (o *ResumeEditResumeRequest) SetSalary(v ResumeObjectsSalaryAddEdit) {
	o.Salary = &v
}

// GetSchedules returns the Schedules field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ResumeEditResumeRequest) GetSchedules() []IncludesIdName {
	if o == nil {
		var ret []IncludesIdName
		return ret
	}
	return o.Schedules
}

// GetSchedulesOk returns a tuple with the Schedules field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ResumeEditResumeRequest) GetSchedulesOk() ([]IncludesIdName, bool) {
	if o == nil || IsNil(o.Schedules) {
		return nil, false
	}
	return o.Schedules, true
}

// HasSchedules returns a boolean if a field has been set.
func (o *ResumeEditResumeRequest) HasSchedules() bool {
	if o != nil && IsNil(o.Schedules) {
		return true
	}

	return false
}

// SetSchedules gets a reference to the given []IncludesIdName and assigns it to the Schedules field.
func (o *ResumeEditResumeRequest) SetSchedules(v []IncludesIdName) {
	o.Schedules = v
}

// GetSite returns the Site field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ResumeEditResumeRequest) GetSite() []ResumeObjectsSite {
	if o == nil {
		var ret []ResumeObjectsSite
		return ret
	}
	return o.Site
}

// GetSiteOk returns a tuple with the Site field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ResumeEditResumeRequest) GetSiteOk() ([]ResumeObjectsSite, bool) {
	if o == nil || IsNil(o.Site) {
		return nil, false
	}
	return o.Site, true
}

// HasSite returns a boolean if a field has been set.
func (o *ResumeEditResumeRequest) HasSite() bool {
	if o != nil && IsNil(o.Site) {
		return true
	}

	return false
}

// SetSite gets a reference to the given []ResumeObjectsSite and assigns it to the Site field.
func (o *ResumeEditResumeRequest) SetSite(v []ResumeObjectsSite) {
	o.Site = v
}

// GetSkillSet returns the SkillSet field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ResumeEditResumeRequest) GetSkillSet() []string {
	if o == nil {
		var ret []string
		return ret
	}
	return o.SkillSet
}

// GetSkillSetOk returns a tuple with the SkillSet field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ResumeEditResumeRequest) GetSkillSetOk() ([]string, bool) {
	if o == nil || IsNil(o.SkillSet) {
		return nil, false
	}
	return o.SkillSet, true
}

// HasSkillSet returns a boolean if a field has been set.
func (o *ResumeEditResumeRequest) HasSkillSet() bool {
	if o != nil && IsNil(o.SkillSet) {
		return true
	}

	return false
}

// SetSkillSet gets a reference to the given []string and assigns it to the SkillSet field.
func (o *ResumeEditResumeRequest) SetSkillSet(v []string) {
	o.SkillSet = v
}

// GetSkills returns the Skills field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ResumeEditResumeRequest) GetSkills() string {
	if o == nil || IsNil(o.Skills.Get()) {
		var ret string
		return ret
	}
	return *o.Skills.Get()
}

// GetSkillsOk returns a tuple with the Skills field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ResumeEditResumeRequest) GetSkillsOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Skills.Get(), o.Skills.IsSet()
}

// HasSkills returns a boolean if a field has been set.
func (o *ResumeEditResumeRequest) HasSkills() bool {
	if o != nil && o.Skills.IsSet() {
		return true
	}

	return false
}

// SetSkills gets a reference to the given NullableString and assigns it to the Skills field.
func (o *ResumeEditResumeRequest) SetSkills(v string) {
	o.Skills.Set(&v)
}
// SetSkillsNil sets the value for Skills to be an explicit nil
func (o *ResumeEditResumeRequest) SetSkillsNil() {
	o.Skills.Set(nil)
}

// UnsetSkills ensures that no value is present for Skills, not even an explicit nil
func (o *ResumeEditResumeRequest) UnsetSkills() {
	o.Skills.Unset()
}

// GetTitle returns the Title field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ResumeEditResumeRequest) GetTitle() string {
	if o == nil || IsNil(o.Title.Get()) {
		var ret string
		return ret
	}
	return *o.Title.Get()
}

// GetTitleOk returns a tuple with the Title field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ResumeEditResumeRequest) GetTitleOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Title.Get(), o.Title.IsSet()
}

// HasTitle returns a boolean if a field has been set.
func (o *ResumeEditResumeRequest) HasTitle() bool {
	if o != nil && o.Title.IsSet() {
		return true
	}

	return false
}

// SetTitle gets a reference to the given NullableString and assigns it to the Title field.
func (o *ResumeEditResumeRequest) SetTitle(v string) {
	o.Title.Set(&v)
}
// SetTitleNil sets the value for Title to be an explicit nil
func (o *ResumeEditResumeRequest) SetTitleNil() {
	o.Title.Set(nil)
}

// UnsetTitle ensures that no value is present for Title, not even an explicit nil
func (o *ResumeEditResumeRequest) UnsetTitle() {
	o.Title.Unset()
}

// GetTotalExperience returns the TotalExperience field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ResumeEditResumeRequest) GetTotalExperience() ResumeObjectsTotalExperience {
	if o == nil || IsNil(o.TotalExperience.Get()) {
		var ret ResumeObjectsTotalExperience
		return ret
	}
	return *o.TotalExperience.Get()
}

// GetTotalExperienceOk returns a tuple with the TotalExperience field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ResumeEditResumeRequest) GetTotalExperienceOk() (*ResumeObjectsTotalExperience, bool) {
	if o == nil {
		return nil, false
	}
	return o.TotalExperience.Get(), o.TotalExperience.IsSet()
}

// HasTotalExperience returns a boolean if a field has been set.
func (o *ResumeEditResumeRequest) HasTotalExperience() bool {
	if o != nil && o.TotalExperience.IsSet() {
		return true
	}

	return false
}

// SetTotalExperience gets a reference to the given NullableResumeObjectsTotalExperience and assigns it to the TotalExperience field.
func (o *ResumeEditResumeRequest) SetTotalExperience(v ResumeObjectsTotalExperience) {
	o.TotalExperience.Set(&v)
}
// SetTotalExperienceNil sets the value for TotalExperience to be an explicit nil
func (o *ResumeEditResumeRequest) SetTotalExperienceNil() {
	o.TotalExperience.Set(nil)
}

// UnsetTotalExperience ensures that no value is present for TotalExperience, not even an explicit nil
func (o *ResumeEditResumeRequest) UnsetTotalExperience() {
	o.TotalExperience.Unset()
}

// GetTravelTime returns the TravelTime field value if set, zero value otherwise.
func (o *ResumeEditResumeRequest) GetTravelTime() IncludesId {
	if o == nil || IsNil(o.TravelTime) {
		var ret IncludesId
		return ret
	}
	return *o.TravelTime
}

// GetTravelTimeOk returns a tuple with the TravelTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ResumeEditResumeRequest) GetTravelTimeOk() (*IncludesId, bool) {
	if o == nil || IsNil(o.TravelTime) {
		return nil, false
	}
	return o.TravelTime, true
}

// HasTravelTime returns a boolean if a field has been set.
func (o *ResumeEditResumeRequest) HasTravelTime() bool {
	if o != nil && !IsNil(o.TravelTime) {
		return true
	}

	return false
}

// SetTravelTime gets a reference to the given IncludesId and assigns it to the TravelTime field.
func (o *ResumeEditResumeRequest) SetTravelTime(v IncludesId) {
	o.TravelTime = &v
}

// GetWorkTicket returns the WorkTicket field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ResumeEditResumeRequest) GetWorkTicket() []IncludesId {
	if o == nil {
		var ret []IncludesId
		return ret
	}
	return o.WorkTicket
}

// GetWorkTicketOk returns a tuple with the WorkTicket field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ResumeEditResumeRequest) GetWorkTicketOk() ([]IncludesId, bool) {
	if o == nil || IsNil(o.WorkTicket) {
		return nil, false
	}
	return o.WorkTicket, true
}

// HasWorkTicket returns a boolean if a field has been set.
func (o *ResumeEditResumeRequest) HasWorkTicket() bool {
	if o != nil && IsNil(o.WorkTicket) {
		return true
	}

	return false
}

// SetWorkTicket gets a reference to the given []IncludesId and assigns it to the WorkTicket field.
func (o *ResumeEditResumeRequest) SetWorkTicket(v []IncludesId) {
	o.WorkTicket = v
}

// GetArea returns the Area field value if set, zero value otherwise.
func (o *ResumeEditResumeRequest) GetArea() Id {
	if o == nil || IsNil(o.Area) {
		var ret Id
		return ret
	}
	return *o.Area
}

// GetAreaOk returns a tuple with the Area field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ResumeEditResumeRequest) GetAreaOk() (*Id, bool) {
	if o == nil || IsNil(o.Area) {
		return nil, false
	}
	return o.Area, true
}

// HasArea returns a boolean if a field has been set.
func (o *ResumeEditResumeRequest) HasArea() bool {
	if o != nil && !IsNil(o.Area) {
		return true
	}

	return false
}

// SetArea gets a reference to the given Id and assigns it to the Area field.
func (o *ResumeEditResumeRequest) SetArea(v Id) {
	o.Area = &v
}

// GetCitizenship returns the Citizenship field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ResumeEditResumeRequest) GetCitizenship() []IncludesId {
	if o == nil {
		var ret []IncludesId
		return ret
	}
	return o.Citizenship
}

// GetCitizenshipOk returns a tuple with the Citizenship field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ResumeEditResumeRequest) GetCitizenshipOk() ([]IncludesId, bool) {
	if o == nil || IsNil(o.Citizenship) {
		return nil, false
	}
	return o.Citizenship, true
}

// HasCitizenship returns a boolean if a field has been set.
func (o *ResumeEditResumeRequest) HasCitizenship() bool {
	if o != nil && IsNil(o.Citizenship) {
		return true
	}

	return false
}

// SetCitizenship gets a reference to the given []IncludesId and assigns it to the Citizenship field.
func (o *ResumeEditResumeRequest) SetCitizenship(v []IncludesId) {
	o.Citizenship = v
}

// GetContact returns the Contact field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ResumeEditResumeRequest) GetContact() []IncludesContact {
	if o == nil {
		var ret []IncludesContact
		return ret
	}
	return o.Contact
}

// GetContactOk returns a tuple with the Contact field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ResumeEditResumeRequest) GetContactOk() ([]IncludesContact, bool) {
	if o == nil || IsNil(o.Contact) {
		return nil, false
	}
	return o.Contact, true
}

// HasContact returns a boolean if a field has been set.
func (o *ResumeEditResumeRequest) HasContact() bool {
	if o != nil && IsNil(o.Contact) {
		return true
	}

	return false
}

// SetContact gets a reference to the given []IncludesContact and assigns it to the Contact field.
func (o *ResumeEditResumeRequest) SetContact(v []IncludesContact) {
	o.Contact = v
}

// GetEducation returns the Education field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ResumeEditResumeRequest) GetEducation() ResumeObjectsEducation {
	if o == nil || IsNil(o.Education.Get()) {
		var ret ResumeObjectsEducation
		return ret
	}
	return *o.Education.Get()
}

// GetEducationOk returns a tuple with the Education field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ResumeEditResumeRequest) GetEducationOk() (*ResumeObjectsEducation, bool) {
	if o == nil {
		return nil, false
	}
	return o.Education.Get(), o.Education.IsSet()
}

// HasEducation returns a boolean if a field has been set.
func (o *ResumeEditResumeRequest) HasEducation() bool {
	if o != nil && o.Education.IsSet() {
		return true
	}

	return false
}

// SetEducation gets a reference to the given NullableResumeObjectsEducation and assigns it to the Education field.
func (o *ResumeEditResumeRequest) SetEducation(v ResumeObjectsEducation) {
	o.Education.Set(&v)
}
// SetEducationNil sets the value for Education to be an explicit nil
func (o *ResumeEditResumeRequest) SetEducationNil() {
	o.Education.Set(nil)
}

// UnsetEducation ensures that no value is present for Education, not even an explicit nil
func (o *ResumeEditResumeRequest) UnsetEducation() {
	o.Education.Unset()
}

// GetExperience returns the Experience field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ResumeEditResumeRequest) GetExperience() []ResumeObjectsExperience {
	if o == nil {
		var ret []ResumeObjectsExperience
		return ret
	}
	return o.Experience
}

// GetExperienceOk returns a tuple with the Experience field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ResumeEditResumeRequest) GetExperienceOk() ([]ResumeObjectsExperience, bool) {
	if o == nil || IsNil(o.Experience) {
		return nil, false
	}
	return o.Experience, true
}

// HasExperience returns a boolean if a field has been set.
func (o *ResumeEditResumeRequest) HasExperience() bool {
	if o != nil && IsNil(o.Experience) {
		return true
	}

	return false
}

// SetExperience gets a reference to the given []ResumeObjectsExperience and assigns it to the Experience field.
func (o *ResumeEditResumeRequest) SetExperience(v []ResumeObjectsExperience) {
	o.Experience = v
}

// GetGender returns the Gender field value if set, zero value otherwise.
func (o *ResumeEditResumeRequest) GetGender() Id {
	if o == nil || IsNil(o.Gender) {
		var ret Id
		return ret
	}
	return *o.Gender
}

// GetGenderOk returns a tuple with the Gender field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ResumeEditResumeRequest) GetGenderOk() (*Id, bool) {
	if o == nil || IsNil(o.Gender) {
		return nil, false
	}
	return o.Gender, true
}

// HasGender returns a boolean if a field has been set.
func (o *ResumeEditResumeRequest) HasGender() bool {
	if o != nil && !IsNil(o.Gender) {
		return true
	}

	return false
}

// SetGender gets a reference to the given Id and assigns it to the Gender field.
func (o *ResumeEditResumeRequest) SetGender(v Id) {
	o.Gender = &v
}

// GetLanguage returns the Language field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ResumeEditResumeRequest) GetLanguage() []IncludesLanguageLevel {
	if o == nil {
		var ret []IncludesLanguageLevel
		return ret
	}
	return o.Language
}

// GetLanguageOk returns a tuple with the Language field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ResumeEditResumeRequest) GetLanguageOk() ([]IncludesLanguageLevel, bool) {
	if o == nil || IsNil(o.Language) {
		return nil, false
	}
	return o.Language, true
}

// HasLanguage returns a boolean if a field has been set.
func (o *ResumeEditResumeRequest) HasLanguage() bool {
	if o != nil && IsNil(o.Language) {
		return true
	}

	return false
}

// SetLanguage gets a reference to the given []IncludesLanguageLevel and assigns it to the Language field.
func (o *ResumeEditResumeRequest) SetLanguage(v []IncludesLanguageLevel) {
	o.Language = v
}

func (o ResumeEditResumeRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ResumeEditResumeRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Access) {
		toSerialize["access"] = o.Access
	}
	if o.BirthDate.IsSet() {
		toSerialize["birth_date"] = o.BirthDate.Get()
	}
	if !IsNil(o.BusinessTripReadiness) {
		toSerialize["business_trip_readiness"] = o.BusinessTripReadiness
	}
	if o.Certificate != nil {
		toSerialize["certificate"] = o.Certificate
	}
	if o.DriverLicenseTypes != nil {
		toSerialize["driver_license_types"] = o.DriverLicenseTypes
	}
	if o.Employments != nil {
		toSerialize["employments"] = o.Employments
	}
	if o.FirstName.IsSet() {
		toSerialize["first_name"] = o.FirstName.Get()
	}
	if o.HasVehicle.IsSet() {
		toSerialize["has_vehicle"] = o.HasVehicle.Get()
	}
	if o.HiddenFields != nil {
		toSerialize["hidden_fields"] = o.HiddenFields
	}
	if o.LastName.IsSet() {
		toSerialize["last_name"] = o.LastName.Get()
	}
	if !IsNil(o.Metro) {
		toSerialize["metro"] = o.Metro
	}
	if o.MiddleName.IsSet() {
		toSerialize["middle_name"] = o.MiddleName.Get()
	}
	if o.Photo.IsSet() {
		toSerialize["photo"] = o.Photo.Get()
	}
	if o.Portfolio != nil {
		toSerialize["portfolio"] = o.Portfolio
	}
	if !IsNil(o.ProfessionalRoles) {
		toSerialize["professional_roles"] = o.ProfessionalRoles
	}
	if o.Recommendation != nil {
		toSerialize["recommendation"] = o.Recommendation
	}
	if !IsNil(o.Relocation) {
		toSerialize["relocation"] = o.Relocation
	}
	if !IsNil(o.ResumeLocale) {
		toSerialize["resume_locale"] = o.ResumeLocale
	}
	if !IsNil(o.Salary) {
		toSerialize["salary"] = o.Salary
	}
	if o.Schedules != nil {
		toSerialize["schedules"] = o.Schedules
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
	if o.Title.IsSet() {
		toSerialize["title"] = o.Title.Get()
	}
	if o.TotalExperience.IsSet() {
		toSerialize["total_experience"] = o.TotalExperience.Get()
	}
	if !IsNil(o.TravelTime) {
		toSerialize["travel_time"] = o.TravelTime
	}
	if o.WorkTicket != nil {
		toSerialize["work_ticket"] = o.WorkTicket
	}
	if !IsNil(o.Area) {
		toSerialize["area"] = o.Area
	}
	if o.Citizenship != nil {
		toSerialize["citizenship"] = o.Citizenship
	}
	if o.Contact != nil {
		toSerialize["contact"] = o.Contact
	}
	if o.Education.IsSet() {
		toSerialize["education"] = o.Education.Get()
	}
	if o.Experience != nil {
		toSerialize["experience"] = o.Experience
	}
	if !IsNil(o.Gender) {
		toSerialize["gender"] = o.Gender
	}
	if o.Language != nil {
		toSerialize["language"] = o.Language
	}
	return toSerialize, nil
}

type NullableResumeEditResumeRequest struct {
	value *ResumeEditResumeRequest
	isSet bool
}

func (v NullableResumeEditResumeRequest) Get() *ResumeEditResumeRequest {
	return v.value
}

func (v *NullableResumeEditResumeRequest) Set(val *ResumeEditResumeRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableResumeEditResumeRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableResumeEditResumeRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableResumeEditResumeRequest(val *ResumeEditResumeRequest) *NullableResumeEditResumeRequest {
	return &NullableResumeEditResumeRequest{value: val, isSet: true}
}

func (v NullableResumeEditResumeRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableResumeEditResumeRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


