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

// checks if the ResumeObjectsExperience type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ResumeObjectsExperience{}

// ResumeObjectsExperience struct for ResumeObjectsExperience
type ResumeObjectsExperience struct {
	Area *IncludesIdNameUrl `json:"area,omitempty"`
	// Название организации
	Company NullableString `json:"company,omitempty"`
	// Уникальный идентификатор организации
	CompanyId NullableString `json:"company_id,omitempty"`
	// Сайт компании
	CompanyUrl NullableString `json:"company_url,omitempty"`
	// Обязанности, функции, достижения
	Description NullableString `json:"description,omitempty"`
	Employer *EmployersEmployerInfoShort `json:"employer,omitempty"`
	// Окончание работы (дата в формате `ГГГГ-ММ-ДД`)
	End NullableString `json:"end,omitempty"`
	// Список отраслей компании. Возможные значения приведены в [справочнике индустрий](#tag/Obshie-spravochniki/operation/get-industries)
	Industries []IncludesIdName `json:"industries"`
	// Deprecated
	Industry *ResumeObjectsIndustry `json:"industry,omitempty"`
	// Должность
	Position string `json:"position"`
	// Начало работы (дата в формате `ГГГГ-ММ-ДД`)
	Start string `json:"start"`
}

type _ResumeObjectsExperience ResumeObjectsExperience

// NewResumeObjectsExperience instantiates a new ResumeObjectsExperience object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewResumeObjectsExperience(industries []IncludesIdName, position string, start string) *ResumeObjectsExperience {
	this := ResumeObjectsExperience{}
	this.Industries = industries
	this.Position = position
	this.Start = start
	return &this
}

// NewResumeObjectsExperienceWithDefaults instantiates a new ResumeObjectsExperience object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewResumeObjectsExperienceWithDefaults() *ResumeObjectsExperience {
	this := ResumeObjectsExperience{}
	return &this
}

// GetArea returns the Area field value if set, zero value otherwise.
func (o *ResumeObjectsExperience) GetArea() IncludesIdNameUrl {
	if o == nil || IsNil(o.Area) {
		var ret IncludesIdNameUrl
		return ret
	}
	return *o.Area
}

// GetAreaOk returns a tuple with the Area field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ResumeObjectsExperience) GetAreaOk() (*IncludesIdNameUrl, bool) {
	if o == nil || IsNil(o.Area) {
		return nil, false
	}
	return o.Area, true
}

// HasArea returns a boolean if a field has been set.
func (o *ResumeObjectsExperience) HasArea() bool {
	if o != nil && !IsNil(o.Area) {
		return true
	}

	return false
}

// SetArea gets a reference to the given IncludesIdNameUrl and assigns it to the Area field.
func (o *ResumeObjectsExperience) SetArea(v IncludesIdNameUrl) {
	o.Area = &v
}

// GetCompany returns the Company field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ResumeObjectsExperience) GetCompany() string {
	if o == nil || IsNil(o.Company.Get()) {
		var ret string
		return ret
	}
	return *o.Company.Get()
}

// GetCompanyOk returns a tuple with the Company field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ResumeObjectsExperience) GetCompanyOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Company.Get(), o.Company.IsSet()
}

// HasCompany returns a boolean if a field has been set.
func (o *ResumeObjectsExperience) HasCompany() bool {
	if o != nil && o.Company.IsSet() {
		return true
	}

	return false
}

// SetCompany gets a reference to the given NullableString and assigns it to the Company field.
func (o *ResumeObjectsExperience) SetCompany(v string) {
	o.Company.Set(&v)
}
// SetCompanyNil sets the value for Company to be an explicit nil
func (o *ResumeObjectsExperience) SetCompanyNil() {
	o.Company.Set(nil)
}

// UnsetCompany ensures that no value is present for Company, not even an explicit nil
func (o *ResumeObjectsExperience) UnsetCompany() {
	o.Company.Unset()
}

// GetCompanyId returns the CompanyId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ResumeObjectsExperience) GetCompanyId() string {
	if o == nil || IsNil(o.CompanyId.Get()) {
		var ret string
		return ret
	}
	return *o.CompanyId.Get()
}

// GetCompanyIdOk returns a tuple with the CompanyId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ResumeObjectsExperience) GetCompanyIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.CompanyId.Get(), o.CompanyId.IsSet()
}

// HasCompanyId returns a boolean if a field has been set.
func (o *ResumeObjectsExperience) HasCompanyId() bool {
	if o != nil && o.CompanyId.IsSet() {
		return true
	}

	return false
}

// SetCompanyId gets a reference to the given NullableString and assigns it to the CompanyId field.
func (o *ResumeObjectsExperience) SetCompanyId(v string) {
	o.CompanyId.Set(&v)
}
// SetCompanyIdNil sets the value for CompanyId to be an explicit nil
func (o *ResumeObjectsExperience) SetCompanyIdNil() {
	o.CompanyId.Set(nil)
}

// UnsetCompanyId ensures that no value is present for CompanyId, not even an explicit nil
func (o *ResumeObjectsExperience) UnsetCompanyId() {
	o.CompanyId.Unset()
}

// GetCompanyUrl returns the CompanyUrl field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ResumeObjectsExperience) GetCompanyUrl() string {
	if o == nil || IsNil(o.CompanyUrl.Get()) {
		var ret string
		return ret
	}
	return *o.CompanyUrl.Get()
}

// GetCompanyUrlOk returns a tuple with the CompanyUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ResumeObjectsExperience) GetCompanyUrlOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.CompanyUrl.Get(), o.CompanyUrl.IsSet()
}

// HasCompanyUrl returns a boolean if a field has been set.
func (o *ResumeObjectsExperience) HasCompanyUrl() bool {
	if o != nil && o.CompanyUrl.IsSet() {
		return true
	}

	return false
}

// SetCompanyUrl gets a reference to the given NullableString and assigns it to the CompanyUrl field.
func (o *ResumeObjectsExperience) SetCompanyUrl(v string) {
	o.CompanyUrl.Set(&v)
}
// SetCompanyUrlNil sets the value for CompanyUrl to be an explicit nil
func (o *ResumeObjectsExperience) SetCompanyUrlNil() {
	o.CompanyUrl.Set(nil)
}

// UnsetCompanyUrl ensures that no value is present for CompanyUrl, not even an explicit nil
func (o *ResumeObjectsExperience) UnsetCompanyUrl() {
	o.CompanyUrl.Unset()
}

// GetDescription returns the Description field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ResumeObjectsExperience) GetDescription() string {
	if o == nil || IsNil(o.Description.Get()) {
		var ret string
		return ret
	}
	return *o.Description.Get()
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ResumeObjectsExperience) GetDescriptionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Description.Get(), o.Description.IsSet()
}

// HasDescription returns a boolean if a field has been set.
func (o *ResumeObjectsExperience) HasDescription() bool {
	if o != nil && o.Description.IsSet() {
		return true
	}

	return false
}

// SetDescription gets a reference to the given NullableString and assigns it to the Description field.
func (o *ResumeObjectsExperience) SetDescription(v string) {
	o.Description.Set(&v)
}
// SetDescriptionNil sets the value for Description to be an explicit nil
func (o *ResumeObjectsExperience) SetDescriptionNil() {
	o.Description.Set(nil)
}

// UnsetDescription ensures that no value is present for Description, not even an explicit nil
func (o *ResumeObjectsExperience) UnsetDescription() {
	o.Description.Unset()
}

// GetEmployer returns the Employer field value if set, zero value otherwise.
func (o *ResumeObjectsExperience) GetEmployer() EmployersEmployerInfoShort {
	if o == nil || IsNil(o.Employer) {
		var ret EmployersEmployerInfoShort
		return ret
	}
	return *o.Employer
}

// GetEmployerOk returns a tuple with the Employer field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ResumeObjectsExperience) GetEmployerOk() (*EmployersEmployerInfoShort, bool) {
	if o == nil || IsNil(o.Employer) {
		return nil, false
	}
	return o.Employer, true
}

// HasEmployer returns a boolean if a field has been set.
func (o *ResumeObjectsExperience) HasEmployer() bool {
	if o != nil && !IsNil(o.Employer) {
		return true
	}

	return false
}

// SetEmployer gets a reference to the given EmployersEmployerInfoShort and assigns it to the Employer field.
func (o *ResumeObjectsExperience) SetEmployer(v EmployersEmployerInfoShort) {
	o.Employer = &v
}

// GetEnd returns the End field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ResumeObjectsExperience) GetEnd() string {
	if o == nil || IsNil(o.End.Get()) {
		var ret string
		return ret
	}
	return *o.End.Get()
}

// GetEndOk returns a tuple with the End field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ResumeObjectsExperience) GetEndOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.End.Get(), o.End.IsSet()
}

// HasEnd returns a boolean if a field has been set.
func (o *ResumeObjectsExperience) HasEnd() bool {
	if o != nil && o.End.IsSet() {
		return true
	}

	return false
}

// SetEnd gets a reference to the given NullableString and assigns it to the End field.
func (o *ResumeObjectsExperience) SetEnd(v string) {
	o.End.Set(&v)
}
// SetEndNil sets the value for End to be an explicit nil
func (o *ResumeObjectsExperience) SetEndNil() {
	o.End.Set(nil)
}

// UnsetEnd ensures that no value is present for End, not even an explicit nil
func (o *ResumeObjectsExperience) UnsetEnd() {
	o.End.Unset()
}

// GetIndustries returns the Industries field value
func (o *ResumeObjectsExperience) GetIndustries() []IncludesIdName {
	if o == nil {
		var ret []IncludesIdName
		return ret
	}

	return o.Industries
}

// GetIndustriesOk returns a tuple with the Industries field value
// and a boolean to check if the value has been set.
func (o *ResumeObjectsExperience) GetIndustriesOk() ([]IncludesIdName, bool) {
	if o == nil {
		return nil, false
	}
	return o.Industries, true
}

// SetIndustries sets field value
func (o *ResumeObjectsExperience) SetIndustries(v []IncludesIdName) {
	o.Industries = v
}

// GetIndustry returns the Industry field value if set, zero value otherwise.
// Deprecated
func (o *ResumeObjectsExperience) GetIndustry() ResumeObjectsIndustry {
	if o == nil || IsNil(o.Industry) {
		var ret ResumeObjectsIndustry
		return ret
	}
	return *o.Industry
}

// GetIndustryOk returns a tuple with the Industry field value if set, nil otherwise
// and a boolean to check if the value has been set.
// Deprecated
func (o *ResumeObjectsExperience) GetIndustryOk() (*ResumeObjectsIndustry, bool) {
	if o == nil || IsNil(o.Industry) {
		return nil, false
	}
	return o.Industry, true
}

// HasIndustry returns a boolean if a field has been set.
func (o *ResumeObjectsExperience) HasIndustry() bool {
	if o != nil && !IsNil(o.Industry) {
		return true
	}

	return false
}

// SetIndustry gets a reference to the given ResumeObjectsIndustry and assigns it to the Industry field.
// Deprecated
func (o *ResumeObjectsExperience) SetIndustry(v ResumeObjectsIndustry) {
	o.Industry = &v
}

// GetPosition returns the Position field value
func (o *ResumeObjectsExperience) GetPosition() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Position
}

// GetPositionOk returns a tuple with the Position field value
// and a boolean to check if the value has been set.
func (o *ResumeObjectsExperience) GetPositionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Position, true
}

// SetPosition sets field value
func (o *ResumeObjectsExperience) SetPosition(v string) {
	o.Position = v
}

// GetStart returns the Start field value
func (o *ResumeObjectsExperience) GetStart() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Start
}

// GetStartOk returns a tuple with the Start field value
// and a boolean to check if the value has been set.
func (o *ResumeObjectsExperience) GetStartOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Start, true
}

// SetStart sets field value
func (o *ResumeObjectsExperience) SetStart(v string) {
	o.Start = v
}

func (o ResumeObjectsExperience) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ResumeObjectsExperience) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Area) {
		toSerialize["area"] = o.Area
	}
	if o.Company.IsSet() {
		toSerialize["company"] = o.Company.Get()
	}
	if o.CompanyId.IsSet() {
		toSerialize["company_id"] = o.CompanyId.Get()
	}
	if o.CompanyUrl.IsSet() {
		toSerialize["company_url"] = o.CompanyUrl.Get()
	}
	if o.Description.IsSet() {
		toSerialize["description"] = o.Description.Get()
	}
	if !IsNil(o.Employer) {
		toSerialize["employer"] = o.Employer
	}
	if o.End.IsSet() {
		toSerialize["end"] = o.End.Get()
	}
	toSerialize["industries"] = o.Industries
	if !IsNil(o.Industry) {
		toSerialize["industry"] = o.Industry
	}
	toSerialize["position"] = o.Position
	toSerialize["start"] = o.Start
	return toSerialize, nil
}

func (o *ResumeObjectsExperience) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"industries",
		"position",
		"start",
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

	varResumeObjectsExperience := _ResumeObjectsExperience{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varResumeObjectsExperience)

	if err != nil {
		return err
	}

	*o = ResumeObjectsExperience(varResumeObjectsExperience)

	return err
}

type NullableResumeObjectsExperience struct {
	value *ResumeObjectsExperience
	isSet bool
}

func (v NullableResumeObjectsExperience) Get() *ResumeObjectsExperience {
	return v.value
}

func (v *NullableResumeObjectsExperience) Set(val *ResumeObjectsExperience) {
	v.value = val
	v.isSet = true
}

func (v NullableResumeObjectsExperience) IsSet() bool {
	return v.isSet
}

func (v *NullableResumeObjectsExperience) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableResumeObjectsExperience(val *ResumeObjectsExperience) *NullableResumeObjectsExperience {
	return &NullableResumeObjectsExperience{value: val, isSet: true}
}

func (v NullableResumeObjectsExperience) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableResumeObjectsExperience) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


