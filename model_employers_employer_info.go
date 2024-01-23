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

// checks if the EmployersEmployerInfo type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &EmployersEmployerInfo{}

// EmployersEmployerInfo struct for EmployersEmployerInfo
type EmployersEmployerInfo struct {
	// Флаг, показывающий, прошел ли работодатель [IT аккредитацию](https://feedback.hh.ru/knowledge-base/article/что-означает-значок-«аккредитована-как-ит-компания»)
	AccreditedItEmployer *bool `json:"accredited_it_employer,omitempty"`
	// Ссылка на описание работодателя на сайте
	AlternateUrl string `json:"alternate_url"`
	ApplicantServices *IncludesEmployerApplicantServices `json:"applicant_services,omitempty"`
	Area EmployersEmployerInfoArea `json:"area"`
	// Строка с кодом HTML (возможно наличие `<script/>` и `<style/>`), которая является альтернативой стандартному описанию работодателя. HTML адаптирован для мобильных устройств и корректно отображается без поддержки Javascript.   При этом:  - Контент тянется по ширине на 100% ширины контейнера и умещается без прокрутки в 300px. - Контент рассчитан на то, что он будет вставлен в обвязку, в которую входит название, логотип, сайт и ссылка на вакансии работодателя. - Изображения, которые могут встретиться в таком описании, адаптированы под retina-дисплеи. - Размер шрифта не меньше 12px, размер межстрочного интервала не меньше 16px.  Значение может быть `null`, если у работодателя отсутствует индивидуальное описание 
	// Deprecated
	BrandedDescription NullableString `json:"branded_description,omitempty"`
	Branding NullableEmployersBrandingEmployerBranding `json:"branding,omitempty"`
	// Описание работодателя в виде строки с кодом HTML (без `<script/>` и `<style/>`)
	Description NullableString `json:"description,omitempty"`
	// Идентификатор работодателя
	Id string `json:"id"`
	// Список отраслей работодателя. Элементы [справочника индустрий](https://api.hh.ru/openapi/redoc#tag/Obshie-spravochniki/operation/get-industries)
	Industries []IncludesIdName `json:"industries"`
	// Список интервью
	InsiderInterviews []EmployersInsiderInterviews `json:"insider_interviews"`
	LogoUrls NullableIncludesLogoUrls `json:"logo_urls,omitempty"`
	// Название работодателя
	Name string `json:"name"`
	// Количество открытых вакансий у работодателя
	OpenVacancies NullableFloat32 `json:"open_vacancies,omitempty"`
	// Если работодатель добавлен в черный список, то вернется `['blacklisted']`, иначе `[]`
	Relations []*string `json:"relations"`
	// Адрес сайта работодателя
	SiteUrl string `json:"site_url"`
	// Флаг, показывающий, прошел ли работодатель [проверку на сайте](https://feedback.hh.ru/article/details/id/5951)
	Trusted bool `json:"trusted"`
	// Тип работодателя (прямой работодатель, кадровое агентство и т.п.). Возможные значения описаны в [справочнике](#tag/Obshie-spravochniki/operation/get-dictionaries) в поле `employer_type`. Возвращает `null`, если тип работодателя скрыт
	Type NullableString `json:"type,omitempty"`
	// URL для получения поисковой выдачи с вакансиями данного работодателя
	VacanciesUrl string `json:"vacancies_url"`
}

type _EmployersEmployerInfo EmployersEmployerInfo

// NewEmployersEmployerInfo instantiates a new EmployersEmployerInfo object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewEmployersEmployerInfo(alternateUrl string, area EmployersEmployerInfoArea, id string, industries []IncludesIdName, insiderInterviews []EmployersInsiderInterviews, name string, relations []*string, siteUrl string, trusted bool, vacanciesUrl string) *EmployersEmployerInfo {
	this := EmployersEmployerInfo{}
	this.AlternateUrl = alternateUrl
	this.Area = area
	this.Id = id
	this.Industries = industries
	this.InsiderInterviews = insiderInterviews
	this.Name = name
	this.Relations = relations
	this.SiteUrl = siteUrl
	this.Trusted = trusted
	this.VacanciesUrl = vacanciesUrl
	return &this
}

// NewEmployersEmployerInfoWithDefaults instantiates a new EmployersEmployerInfo object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewEmployersEmployerInfoWithDefaults() *EmployersEmployerInfo {
	this := EmployersEmployerInfo{}
	return &this
}

// GetAccreditedItEmployer returns the AccreditedItEmployer field value if set, zero value otherwise.
func (o *EmployersEmployerInfo) GetAccreditedItEmployer() bool {
	if o == nil || IsNil(o.AccreditedItEmployer) {
		var ret bool
		return ret
	}
	return *o.AccreditedItEmployer
}

// GetAccreditedItEmployerOk returns a tuple with the AccreditedItEmployer field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EmployersEmployerInfo) GetAccreditedItEmployerOk() (*bool, bool) {
	if o == nil || IsNil(o.AccreditedItEmployer) {
		return nil, false
	}
	return o.AccreditedItEmployer, true
}

// HasAccreditedItEmployer returns a boolean if a field has been set.
func (o *EmployersEmployerInfo) HasAccreditedItEmployer() bool {
	if o != nil && !IsNil(o.AccreditedItEmployer) {
		return true
	}

	return false
}

// SetAccreditedItEmployer gets a reference to the given bool and assigns it to the AccreditedItEmployer field.
func (o *EmployersEmployerInfo) SetAccreditedItEmployer(v bool) {
	o.AccreditedItEmployer = &v
}

// GetAlternateUrl returns the AlternateUrl field value
func (o *EmployersEmployerInfo) GetAlternateUrl() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.AlternateUrl
}

// GetAlternateUrlOk returns a tuple with the AlternateUrl field value
// and a boolean to check if the value has been set.
func (o *EmployersEmployerInfo) GetAlternateUrlOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AlternateUrl, true
}

// SetAlternateUrl sets field value
func (o *EmployersEmployerInfo) SetAlternateUrl(v string) {
	o.AlternateUrl = v
}

// GetApplicantServices returns the ApplicantServices field value if set, zero value otherwise.
func (o *EmployersEmployerInfo) GetApplicantServices() IncludesEmployerApplicantServices {
	if o == nil || IsNil(o.ApplicantServices) {
		var ret IncludesEmployerApplicantServices
		return ret
	}
	return *o.ApplicantServices
}

// GetApplicantServicesOk returns a tuple with the ApplicantServices field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EmployersEmployerInfo) GetApplicantServicesOk() (*IncludesEmployerApplicantServices, bool) {
	if o == nil || IsNil(o.ApplicantServices) {
		return nil, false
	}
	return o.ApplicantServices, true
}

// HasApplicantServices returns a boolean if a field has been set.
func (o *EmployersEmployerInfo) HasApplicantServices() bool {
	if o != nil && !IsNil(o.ApplicantServices) {
		return true
	}

	return false
}

// SetApplicantServices gets a reference to the given IncludesEmployerApplicantServices and assigns it to the ApplicantServices field.
func (o *EmployersEmployerInfo) SetApplicantServices(v IncludesEmployerApplicantServices) {
	o.ApplicantServices = &v
}

// GetArea returns the Area field value
func (o *EmployersEmployerInfo) GetArea() EmployersEmployerInfoArea {
	if o == nil {
		var ret EmployersEmployerInfoArea
		return ret
	}

	return o.Area
}

// GetAreaOk returns a tuple with the Area field value
// and a boolean to check if the value has been set.
func (o *EmployersEmployerInfo) GetAreaOk() (*EmployersEmployerInfoArea, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Area, true
}

// SetArea sets field value
func (o *EmployersEmployerInfo) SetArea(v EmployersEmployerInfoArea) {
	o.Area = v
}

// GetBrandedDescription returns the BrandedDescription field value if set, zero value otherwise (both if not set or set to explicit null).
// Deprecated
func (o *EmployersEmployerInfo) GetBrandedDescription() string {
	if o == nil || IsNil(o.BrandedDescription.Get()) {
		var ret string
		return ret
	}
	return *o.BrandedDescription.Get()
}

// GetBrandedDescriptionOk returns a tuple with the BrandedDescription field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
// Deprecated
func (o *EmployersEmployerInfo) GetBrandedDescriptionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.BrandedDescription.Get(), o.BrandedDescription.IsSet()
}

// HasBrandedDescription returns a boolean if a field has been set.
func (o *EmployersEmployerInfo) HasBrandedDescription() bool {
	if o != nil && o.BrandedDescription.IsSet() {
		return true
	}

	return false
}

// SetBrandedDescription gets a reference to the given NullableString and assigns it to the BrandedDescription field.
// Deprecated
func (o *EmployersEmployerInfo) SetBrandedDescription(v string) {
	o.BrandedDescription.Set(&v)
}
// SetBrandedDescriptionNil sets the value for BrandedDescription to be an explicit nil
func (o *EmployersEmployerInfo) SetBrandedDescriptionNil() {
	o.BrandedDescription.Set(nil)
}

// UnsetBrandedDescription ensures that no value is present for BrandedDescription, not even an explicit nil
func (o *EmployersEmployerInfo) UnsetBrandedDescription() {
	o.BrandedDescription.Unset()
}

// GetBranding returns the Branding field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *EmployersEmployerInfo) GetBranding() EmployersBrandingEmployerBranding {
	if o == nil || IsNil(o.Branding.Get()) {
		var ret EmployersBrandingEmployerBranding
		return ret
	}
	return *o.Branding.Get()
}

// GetBrandingOk returns a tuple with the Branding field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *EmployersEmployerInfo) GetBrandingOk() (*EmployersBrandingEmployerBranding, bool) {
	if o == nil {
		return nil, false
	}
	return o.Branding.Get(), o.Branding.IsSet()
}

// HasBranding returns a boolean if a field has been set.
func (o *EmployersEmployerInfo) HasBranding() bool {
	if o != nil && o.Branding.IsSet() {
		return true
	}

	return false
}

// SetBranding gets a reference to the given NullableEmployersBrandingEmployerBranding and assigns it to the Branding field.
func (o *EmployersEmployerInfo) SetBranding(v EmployersBrandingEmployerBranding) {
	o.Branding.Set(&v)
}
// SetBrandingNil sets the value for Branding to be an explicit nil
func (o *EmployersEmployerInfo) SetBrandingNil() {
	o.Branding.Set(nil)
}

// UnsetBranding ensures that no value is present for Branding, not even an explicit nil
func (o *EmployersEmployerInfo) UnsetBranding() {
	o.Branding.Unset()
}

// GetDescription returns the Description field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *EmployersEmployerInfo) GetDescription() string {
	if o == nil || IsNil(o.Description.Get()) {
		var ret string
		return ret
	}
	return *o.Description.Get()
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *EmployersEmployerInfo) GetDescriptionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Description.Get(), o.Description.IsSet()
}

// HasDescription returns a boolean if a field has been set.
func (o *EmployersEmployerInfo) HasDescription() bool {
	if o != nil && o.Description.IsSet() {
		return true
	}

	return false
}

// SetDescription gets a reference to the given NullableString and assigns it to the Description field.
func (o *EmployersEmployerInfo) SetDescription(v string) {
	o.Description.Set(&v)
}
// SetDescriptionNil sets the value for Description to be an explicit nil
func (o *EmployersEmployerInfo) SetDescriptionNil() {
	o.Description.Set(nil)
}

// UnsetDescription ensures that no value is present for Description, not even an explicit nil
func (o *EmployersEmployerInfo) UnsetDescription() {
	o.Description.Unset()
}

// GetId returns the Id field value
func (o *EmployersEmployerInfo) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *EmployersEmployerInfo) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *EmployersEmployerInfo) SetId(v string) {
	o.Id = v
}

// GetIndustries returns the Industries field value
func (o *EmployersEmployerInfo) GetIndustries() []IncludesIdName {
	if o == nil {
		var ret []IncludesIdName
		return ret
	}

	return o.Industries
}

// GetIndustriesOk returns a tuple with the Industries field value
// and a boolean to check if the value has been set.
func (o *EmployersEmployerInfo) GetIndustriesOk() ([]IncludesIdName, bool) {
	if o == nil {
		return nil, false
	}
	return o.Industries, true
}

// SetIndustries sets field value
func (o *EmployersEmployerInfo) SetIndustries(v []IncludesIdName) {
	o.Industries = v
}

// GetInsiderInterviews returns the InsiderInterviews field value
func (o *EmployersEmployerInfo) GetInsiderInterviews() []EmployersInsiderInterviews {
	if o == nil {
		var ret []EmployersInsiderInterviews
		return ret
	}

	return o.InsiderInterviews
}

// GetInsiderInterviewsOk returns a tuple with the InsiderInterviews field value
// and a boolean to check if the value has been set.
func (o *EmployersEmployerInfo) GetInsiderInterviewsOk() ([]EmployersInsiderInterviews, bool) {
	if o == nil {
		return nil, false
	}
	return o.InsiderInterviews, true
}

// SetInsiderInterviews sets field value
func (o *EmployersEmployerInfo) SetInsiderInterviews(v []EmployersInsiderInterviews) {
	o.InsiderInterviews = v
}

// GetLogoUrls returns the LogoUrls field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *EmployersEmployerInfo) GetLogoUrls() IncludesLogoUrls {
	if o == nil || IsNil(o.LogoUrls.Get()) {
		var ret IncludesLogoUrls
		return ret
	}
	return *o.LogoUrls.Get()
}

// GetLogoUrlsOk returns a tuple with the LogoUrls field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *EmployersEmployerInfo) GetLogoUrlsOk() (*IncludesLogoUrls, bool) {
	if o == nil {
		return nil, false
	}
	return o.LogoUrls.Get(), o.LogoUrls.IsSet()
}

// HasLogoUrls returns a boolean if a field has been set.
func (o *EmployersEmployerInfo) HasLogoUrls() bool {
	if o != nil && o.LogoUrls.IsSet() {
		return true
	}

	return false
}

// SetLogoUrls gets a reference to the given NullableIncludesLogoUrls and assigns it to the LogoUrls field.
func (o *EmployersEmployerInfo) SetLogoUrls(v IncludesLogoUrls) {
	o.LogoUrls.Set(&v)
}
// SetLogoUrlsNil sets the value for LogoUrls to be an explicit nil
func (o *EmployersEmployerInfo) SetLogoUrlsNil() {
	o.LogoUrls.Set(nil)
}

// UnsetLogoUrls ensures that no value is present for LogoUrls, not even an explicit nil
func (o *EmployersEmployerInfo) UnsetLogoUrls() {
	o.LogoUrls.Unset()
}

// GetName returns the Name field value
func (o *EmployersEmployerInfo) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *EmployersEmployerInfo) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *EmployersEmployerInfo) SetName(v string) {
	o.Name = v
}

// GetOpenVacancies returns the OpenVacancies field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *EmployersEmployerInfo) GetOpenVacancies() float32 {
	if o == nil || IsNil(o.OpenVacancies.Get()) {
		var ret float32
		return ret
	}
	return *o.OpenVacancies.Get()
}

// GetOpenVacanciesOk returns a tuple with the OpenVacancies field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *EmployersEmployerInfo) GetOpenVacanciesOk() (*float32, bool) {
	if o == nil {
		return nil, false
	}
	return o.OpenVacancies.Get(), o.OpenVacancies.IsSet()
}

// HasOpenVacancies returns a boolean if a field has been set.
func (o *EmployersEmployerInfo) HasOpenVacancies() bool {
	if o != nil && o.OpenVacancies.IsSet() {
		return true
	}

	return false
}

// SetOpenVacancies gets a reference to the given NullableFloat32 and assigns it to the OpenVacancies field.
func (o *EmployersEmployerInfo) SetOpenVacancies(v float32) {
	o.OpenVacancies.Set(&v)
}
// SetOpenVacanciesNil sets the value for OpenVacancies to be an explicit nil
func (o *EmployersEmployerInfo) SetOpenVacanciesNil() {
	o.OpenVacancies.Set(nil)
}

// UnsetOpenVacancies ensures that no value is present for OpenVacancies, not even an explicit nil
func (o *EmployersEmployerInfo) UnsetOpenVacancies() {
	o.OpenVacancies.Unset()
}

// GetRelations returns the Relations field value
func (o *EmployersEmployerInfo) GetRelations() []*string {
	if o == nil {
		var ret []*string
		return ret
	}

	return o.Relations
}

// GetRelationsOk returns a tuple with the Relations field value
// and a boolean to check if the value has been set.
func (o *EmployersEmployerInfo) GetRelationsOk() ([]*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Relations, true
}

// SetRelations sets field value
func (o *EmployersEmployerInfo) SetRelations(v []*string) {
	o.Relations = v
}

// GetSiteUrl returns the SiteUrl field value
func (o *EmployersEmployerInfo) GetSiteUrl() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.SiteUrl
}

// GetSiteUrlOk returns a tuple with the SiteUrl field value
// and a boolean to check if the value has been set.
func (o *EmployersEmployerInfo) GetSiteUrlOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SiteUrl, true
}

// SetSiteUrl sets field value
func (o *EmployersEmployerInfo) SetSiteUrl(v string) {
	o.SiteUrl = v
}

// GetTrusted returns the Trusted field value
func (o *EmployersEmployerInfo) GetTrusted() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.Trusted
}

// GetTrustedOk returns a tuple with the Trusted field value
// and a boolean to check if the value has been set.
func (o *EmployersEmployerInfo) GetTrustedOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Trusted, true
}

// SetTrusted sets field value
func (o *EmployersEmployerInfo) SetTrusted(v bool) {
	o.Trusted = v
}

// GetType returns the Type field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *EmployersEmployerInfo) GetType() string {
	if o == nil || IsNil(o.Type.Get()) {
		var ret string
		return ret
	}
	return *o.Type.Get()
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *EmployersEmployerInfo) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Type.Get(), o.Type.IsSet()
}

// HasType returns a boolean if a field has been set.
func (o *EmployersEmployerInfo) HasType() bool {
	if o != nil && o.Type.IsSet() {
		return true
	}

	return false
}

// SetType gets a reference to the given NullableString and assigns it to the Type field.
func (o *EmployersEmployerInfo) SetType(v string) {
	o.Type.Set(&v)
}
// SetTypeNil sets the value for Type to be an explicit nil
func (o *EmployersEmployerInfo) SetTypeNil() {
	o.Type.Set(nil)
}

// UnsetType ensures that no value is present for Type, not even an explicit nil
func (o *EmployersEmployerInfo) UnsetType() {
	o.Type.Unset()
}

// GetVacanciesUrl returns the VacanciesUrl field value
func (o *EmployersEmployerInfo) GetVacanciesUrl() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.VacanciesUrl
}

// GetVacanciesUrlOk returns a tuple with the VacanciesUrl field value
// and a boolean to check if the value has been set.
func (o *EmployersEmployerInfo) GetVacanciesUrlOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.VacanciesUrl, true
}

// SetVacanciesUrl sets field value
func (o *EmployersEmployerInfo) SetVacanciesUrl(v string) {
	o.VacanciesUrl = v
}

func (o EmployersEmployerInfo) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o EmployersEmployerInfo) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.AccreditedItEmployer) {
		toSerialize["accredited_it_employer"] = o.AccreditedItEmployer
	}
	toSerialize["alternate_url"] = o.AlternateUrl
	if !IsNil(o.ApplicantServices) {
		toSerialize["applicant_services"] = o.ApplicantServices
	}
	toSerialize["area"] = o.Area
	if o.BrandedDescription.IsSet() {
		toSerialize["branded_description"] = o.BrandedDescription.Get()
	}
	if o.Branding.IsSet() {
		toSerialize["branding"] = o.Branding.Get()
	}
	if o.Description.IsSet() {
		toSerialize["description"] = o.Description.Get()
	}
	toSerialize["id"] = o.Id
	toSerialize["industries"] = o.Industries
	toSerialize["insider_interviews"] = o.InsiderInterviews
	if o.LogoUrls.IsSet() {
		toSerialize["logo_urls"] = o.LogoUrls.Get()
	}
	toSerialize["name"] = o.Name
	if o.OpenVacancies.IsSet() {
		toSerialize["open_vacancies"] = o.OpenVacancies.Get()
	}
	toSerialize["relations"] = o.Relations
	toSerialize["site_url"] = o.SiteUrl
	toSerialize["trusted"] = o.Trusted
	if o.Type.IsSet() {
		toSerialize["type"] = o.Type.Get()
	}
	toSerialize["vacancies_url"] = o.VacanciesUrl
	return toSerialize, nil
}

func (o *EmployersEmployerInfo) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"alternate_url",
		"area",
		"id",
		"industries",
		"insider_interviews",
		"name",
		"relations",
		"site_url",
		"trusted",
		"vacancies_url",
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

	varEmployersEmployerInfo := _EmployersEmployerInfo{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varEmployersEmployerInfo)

	if err != nil {
		return err
	}

	*o = EmployersEmployerInfo(varEmployersEmployerInfo)

	return err
}

type NullableEmployersEmployerInfo struct {
	value *EmployersEmployerInfo
	isSet bool
}

func (v NullableEmployersEmployerInfo) Get() *EmployersEmployerInfo {
	return v.value
}

func (v *NullableEmployersEmployerInfo) Set(val *EmployersEmployerInfo) {
	v.value = val
	v.isSet = true
}

func (v NullableEmployersEmployerInfo) IsSet() bool {
	return v.isSet
}

func (v *NullableEmployersEmployerInfo) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEmployersEmployerInfo(val *EmployersEmployerInfo) *NullableEmployersEmployerInfo {
	return &NullableEmployersEmployerInfo{value: val, isSet: true}
}

func (v NullableEmployersEmployerInfo) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEmployersEmployerInfo) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


