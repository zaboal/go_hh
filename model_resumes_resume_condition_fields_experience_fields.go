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
)

// checks if the ResumesResumeConditionFieldsExperienceFields type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ResumesResumeConditionFieldsExperienceFields{}

// ResumesResumeConditionFieldsExperienceFields struct for ResumesResumeConditionFieldsExperienceFields
type ResumesResumeConditionFieldsExperienceFields struct {
	Area NullableResumesResumeConditionFieldsRequiredWithTitle `json:"area,omitempty"`
	Company NullableResumesResumeConditionFieldsRequiredLengthWithTitle `json:"company,omitempty"`
	CompanyUrl NullableResumesResumeConditionFieldsRequiredLengthWithTitle `json:"company_url,omitempty"`
	Description NullableResumesResumeConditionFieldsRequiredLengthWithTitle `json:"description,omitempty"`
	End NullableResumesResumeConditionFieldsRequiredDateWithTitle `json:"end,omitempty"`
	Industries NullableResumesResumeConditionFieldsRequiredCountWithTitle `json:"industries,omitempty"`
	Industry NullableResumesResumeConditionFieldsRequiredWithTitle `json:"industry,omitempty"`
	Position NullableResumesResumeConditionFieldsRequiredLengthWithTitle `json:"position,omitempty"`
	Start NullableResumesResumeConditionFieldsRequiredDateWithTitle `json:"start,omitempty"`
}

// NewResumesResumeConditionFieldsExperienceFields instantiates a new ResumesResumeConditionFieldsExperienceFields object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewResumesResumeConditionFieldsExperienceFields() *ResumesResumeConditionFieldsExperienceFields {
	this := ResumesResumeConditionFieldsExperienceFields{}
	return &this
}

// NewResumesResumeConditionFieldsExperienceFieldsWithDefaults instantiates a new ResumesResumeConditionFieldsExperienceFields object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewResumesResumeConditionFieldsExperienceFieldsWithDefaults() *ResumesResumeConditionFieldsExperienceFields {
	this := ResumesResumeConditionFieldsExperienceFields{}
	return &this
}

// GetArea returns the Area field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ResumesResumeConditionFieldsExperienceFields) GetArea() ResumesResumeConditionFieldsRequiredWithTitle {
	if o == nil || IsNil(o.Area.Get()) {
		var ret ResumesResumeConditionFieldsRequiredWithTitle
		return ret
	}
	return *o.Area.Get()
}

// GetAreaOk returns a tuple with the Area field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ResumesResumeConditionFieldsExperienceFields) GetAreaOk() (*ResumesResumeConditionFieldsRequiredWithTitle, bool) {
	if o == nil {
		return nil, false
	}
	return o.Area.Get(), o.Area.IsSet()
}

// HasArea returns a boolean if a field has been set.
func (o *ResumesResumeConditionFieldsExperienceFields) HasArea() bool {
	if o != nil && o.Area.IsSet() {
		return true
	}

	return false
}

// SetArea gets a reference to the given NullableResumesResumeConditionFieldsRequiredWithTitle and assigns it to the Area field.
func (o *ResumesResumeConditionFieldsExperienceFields) SetArea(v ResumesResumeConditionFieldsRequiredWithTitle) {
	o.Area.Set(&v)
}
// SetAreaNil sets the value for Area to be an explicit nil
func (o *ResumesResumeConditionFieldsExperienceFields) SetAreaNil() {
	o.Area.Set(nil)
}

// UnsetArea ensures that no value is present for Area, not even an explicit nil
func (o *ResumesResumeConditionFieldsExperienceFields) UnsetArea() {
	o.Area.Unset()
}

// GetCompany returns the Company field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ResumesResumeConditionFieldsExperienceFields) GetCompany() ResumesResumeConditionFieldsRequiredLengthWithTitle {
	if o == nil || IsNil(o.Company.Get()) {
		var ret ResumesResumeConditionFieldsRequiredLengthWithTitle
		return ret
	}
	return *o.Company.Get()
}

// GetCompanyOk returns a tuple with the Company field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ResumesResumeConditionFieldsExperienceFields) GetCompanyOk() (*ResumesResumeConditionFieldsRequiredLengthWithTitle, bool) {
	if o == nil {
		return nil, false
	}
	return o.Company.Get(), o.Company.IsSet()
}

// HasCompany returns a boolean if a field has been set.
func (o *ResumesResumeConditionFieldsExperienceFields) HasCompany() bool {
	if o != nil && o.Company.IsSet() {
		return true
	}

	return false
}

// SetCompany gets a reference to the given NullableResumesResumeConditionFieldsRequiredLengthWithTitle and assigns it to the Company field.
func (o *ResumesResumeConditionFieldsExperienceFields) SetCompany(v ResumesResumeConditionFieldsRequiredLengthWithTitle) {
	o.Company.Set(&v)
}
// SetCompanyNil sets the value for Company to be an explicit nil
func (o *ResumesResumeConditionFieldsExperienceFields) SetCompanyNil() {
	o.Company.Set(nil)
}

// UnsetCompany ensures that no value is present for Company, not even an explicit nil
func (o *ResumesResumeConditionFieldsExperienceFields) UnsetCompany() {
	o.Company.Unset()
}

// GetCompanyUrl returns the CompanyUrl field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ResumesResumeConditionFieldsExperienceFields) GetCompanyUrl() ResumesResumeConditionFieldsRequiredLengthWithTitle {
	if o == nil || IsNil(o.CompanyUrl.Get()) {
		var ret ResumesResumeConditionFieldsRequiredLengthWithTitle
		return ret
	}
	return *o.CompanyUrl.Get()
}

// GetCompanyUrlOk returns a tuple with the CompanyUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ResumesResumeConditionFieldsExperienceFields) GetCompanyUrlOk() (*ResumesResumeConditionFieldsRequiredLengthWithTitle, bool) {
	if o == nil {
		return nil, false
	}
	return o.CompanyUrl.Get(), o.CompanyUrl.IsSet()
}

// HasCompanyUrl returns a boolean if a field has been set.
func (o *ResumesResumeConditionFieldsExperienceFields) HasCompanyUrl() bool {
	if o != nil && o.CompanyUrl.IsSet() {
		return true
	}

	return false
}

// SetCompanyUrl gets a reference to the given NullableResumesResumeConditionFieldsRequiredLengthWithTitle and assigns it to the CompanyUrl field.
func (o *ResumesResumeConditionFieldsExperienceFields) SetCompanyUrl(v ResumesResumeConditionFieldsRequiredLengthWithTitle) {
	o.CompanyUrl.Set(&v)
}
// SetCompanyUrlNil sets the value for CompanyUrl to be an explicit nil
func (o *ResumesResumeConditionFieldsExperienceFields) SetCompanyUrlNil() {
	o.CompanyUrl.Set(nil)
}

// UnsetCompanyUrl ensures that no value is present for CompanyUrl, not even an explicit nil
func (o *ResumesResumeConditionFieldsExperienceFields) UnsetCompanyUrl() {
	o.CompanyUrl.Unset()
}

// GetDescription returns the Description field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ResumesResumeConditionFieldsExperienceFields) GetDescription() ResumesResumeConditionFieldsRequiredLengthWithTitle {
	if o == nil || IsNil(o.Description.Get()) {
		var ret ResumesResumeConditionFieldsRequiredLengthWithTitle
		return ret
	}
	return *o.Description.Get()
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ResumesResumeConditionFieldsExperienceFields) GetDescriptionOk() (*ResumesResumeConditionFieldsRequiredLengthWithTitle, bool) {
	if o == nil {
		return nil, false
	}
	return o.Description.Get(), o.Description.IsSet()
}

// HasDescription returns a boolean if a field has been set.
func (o *ResumesResumeConditionFieldsExperienceFields) HasDescription() bool {
	if o != nil && o.Description.IsSet() {
		return true
	}

	return false
}

// SetDescription gets a reference to the given NullableResumesResumeConditionFieldsRequiredLengthWithTitle and assigns it to the Description field.
func (o *ResumesResumeConditionFieldsExperienceFields) SetDescription(v ResumesResumeConditionFieldsRequiredLengthWithTitle) {
	o.Description.Set(&v)
}
// SetDescriptionNil sets the value for Description to be an explicit nil
func (o *ResumesResumeConditionFieldsExperienceFields) SetDescriptionNil() {
	o.Description.Set(nil)
}

// UnsetDescription ensures that no value is present for Description, not even an explicit nil
func (o *ResumesResumeConditionFieldsExperienceFields) UnsetDescription() {
	o.Description.Unset()
}

// GetEnd returns the End field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ResumesResumeConditionFieldsExperienceFields) GetEnd() ResumesResumeConditionFieldsRequiredDateWithTitle {
	if o == nil || IsNil(o.End.Get()) {
		var ret ResumesResumeConditionFieldsRequiredDateWithTitle
		return ret
	}
	return *o.End.Get()
}

// GetEndOk returns a tuple with the End field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ResumesResumeConditionFieldsExperienceFields) GetEndOk() (*ResumesResumeConditionFieldsRequiredDateWithTitle, bool) {
	if o == nil {
		return nil, false
	}
	return o.End.Get(), o.End.IsSet()
}

// HasEnd returns a boolean if a field has been set.
func (o *ResumesResumeConditionFieldsExperienceFields) HasEnd() bool {
	if o != nil && o.End.IsSet() {
		return true
	}

	return false
}

// SetEnd gets a reference to the given NullableResumesResumeConditionFieldsRequiredDateWithTitle and assigns it to the End field.
func (o *ResumesResumeConditionFieldsExperienceFields) SetEnd(v ResumesResumeConditionFieldsRequiredDateWithTitle) {
	o.End.Set(&v)
}
// SetEndNil sets the value for End to be an explicit nil
func (o *ResumesResumeConditionFieldsExperienceFields) SetEndNil() {
	o.End.Set(nil)
}

// UnsetEnd ensures that no value is present for End, not even an explicit nil
func (o *ResumesResumeConditionFieldsExperienceFields) UnsetEnd() {
	o.End.Unset()
}

// GetIndustries returns the Industries field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ResumesResumeConditionFieldsExperienceFields) GetIndustries() ResumesResumeConditionFieldsRequiredCountWithTitle {
	if o == nil || IsNil(o.Industries.Get()) {
		var ret ResumesResumeConditionFieldsRequiredCountWithTitle
		return ret
	}
	return *o.Industries.Get()
}

// GetIndustriesOk returns a tuple with the Industries field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ResumesResumeConditionFieldsExperienceFields) GetIndustriesOk() (*ResumesResumeConditionFieldsRequiredCountWithTitle, bool) {
	if o == nil {
		return nil, false
	}
	return o.Industries.Get(), o.Industries.IsSet()
}

// HasIndustries returns a boolean if a field has been set.
func (o *ResumesResumeConditionFieldsExperienceFields) HasIndustries() bool {
	if o != nil && o.Industries.IsSet() {
		return true
	}

	return false
}

// SetIndustries gets a reference to the given NullableResumesResumeConditionFieldsRequiredCountWithTitle and assigns it to the Industries field.
func (o *ResumesResumeConditionFieldsExperienceFields) SetIndustries(v ResumesResumeConditionFieldsRequiredCountWithTitle) {
	o.Industries.Set(&v)
}
// SetIndustriesNil sets the value for Industries to be an explicit nil
func (o *ResumesResumeConditionFieldsExperienceFields) SetIndustriesNil() {
	o.Industries.Set(nil)
}

// UnsetIndustries ensures that no value is present for Industries, not even an explicit nil
func (o *ResumesResumeConditionFieldsExperienceFields) UnsetIndustries() {
	o.Industries.Unset()
}

// GetIndustry returns the Industry field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ResumesResumeConditionFieldsExperienceFields) GetIndustry() ResumesResumeConditionFieldsRequiredWithTitle {
	if o == nil || IsNil(o.Industry.Get()) {
		var ret ResumesResumeConditionFieldsRequiredWithTitle
		return ret
	}
	return *o.Industry.Get()
}

// GetIndustryOk returns a tuple with the Industry field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ResumesResumeConditionFieldsExperienceFields) GetIndustryOk() (*ResumesResumeConditionFieldsRequiredWithTitle, bool) {
	if o == nil {
		return nil, false
	}
	return o.Industry.Get(), o.Industry.IsSet()
}

// HasIndustry returns a boolean if a field has been set.
func (o *ResumesResumeConditionFieldsExperienceFields) HasIndustry() bool {
	if o != nil && o.Industry.IsSet() {
		return true
	}

	return false
}

// SetIndustry gets a reference to the given NullableResumesResumeConditionFieldsRequiredWithTitle and assigns it to the Industry field.
func (o *ResumesResumeConditionFieldsExperienceFields) SetIndustry(v ResumesResumeConditionFieldsRequiredWithTitle) {
	o.Industry.Set(&v)
}
// SetIndustryNil sets the value for Industry to be an explicit nil
func (o *ResumesResumeConditionFieldsExperienceFields) SetIndustryNil() {
	o.Industry.Set(nil)
}

// UnsetIndustry ensures that no value is present for Industry, not even an explicit nil
func (o *ResumesResumeConditionFieldsExperienceFields) UnsetIndustry() {
	o.Industry.Unset()
}

// GetPosition returns the Position field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ResumesResumeConditionFieldsExperienceFields) GetPosition() ResumesResumeConditionFieldsRequiredLengthWithTitle {
	if o == nil || IsNil(o.Position.Get()) {
		var ret ResumesResumeConditionFieldsRequiredLengthWithTitle
		return ret
	}
	return *o.Position.Get()
}

// GetPositionOk returns a tuple with the Position field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ResumesResumeConditionFieldsExperienceFields) GetPositionOk() (*ResumesResumeConditionFieldsRequiredLengthWithTitle, bool) {
	if o == nil {
		return nil, false
	}
	return o.Position.Get(), o.Position.IsSet()
}

// HasPosition returns a boolean if a field has been set.
func (o *ResumesResumeConditionFieldsExperienceFields) HasPosition() bool {
	if o != nil && o.Position.IsSet() {
		return true
	}

	return false
}

// SetPosition gets a reference to the given NullableResumesResumeConditionFieldsRequiredLengthWithTitle and assigns it to the Position field.
func (o *ResumesResumeConditionFieldsExperienceFields) SetPosition(v ResumesResumeConditionFieldsRequiredLengthWithTitle) {
	o.Position.Set(&v)
}
// SetPositionNil sets the value for Position to be an explicit nil
func (o *ResumesResumeConditionFieldsExperienceFields) SetPositionNil() {
	o.Position.Set(nil)
}

// UnsetPosition ensures that no value is present for Position, not even an explicit nil
func (o *ResumesResumeConditionFieldsExperienceFields) UnsetPosition() {
	o.Position.Unset()
}

// GetStart returns the Start field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ResumesResumeConditionFieldsExperienceFields) GetStart() ResumesResumeConditionFieldsRequiredDateWithTitle {
	if o == nil || IsNil(o.Start.Get()) {
		var ret ResumesResumeConditionFieldsRequiredDateWithTitle
		return ret
	}
	return *o.Start.Get()
}

// GetStartOk returns a tuple with the Start field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ResumesResumeConditionFieldsExperienceFields) GetStartOk() (*ResumesResumeConditionFieldsRequiredDateWithTitle, bool) {
	if o == nil {
		return nil, false
	}
	return o.Start.Get(), o.Start.IsSet()
}

// HasStart returns a boolean if a field has been set.
func (o *ResumesResumeConditionFieldsExperienceFields) HasStart() bool {
	if o != nil && o.Start.IsSet() {
		return true
	}

	return false
}

// SetStart gets a reference to the given NullableResumesResumeConditionFieldsRequiredDateWithTitle and assigns it to the Start field.
func (o *ResumesResumeConditionFieldsExperienceFields) SetStart(v ResumesResumeConditionFieldsRequiredDateWithTitle) {
	o.Start.Set(&v)
}
// SetStartNil sets the value for Start to be an explicit nil
func (o *ResumesResumeConditionFieldsExperienceFields) SetStartNil() {
	o.Start.Set(nil)
}

// UnsetStart ensures that no value is present for Start, not even an explicit nil
func (o *ResumesResumeConditionFieldsExperienceFields) UnsetStart() {
	o.Start.Unset()
}

func (o ResumesResumeConditionFieldsExperienceFields) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ResumesResumeConditionFieldsExperienceFields) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.Area.IsSet() {
		toSerialize["area"] = o.Area.Get()
	}
	if o.Company.IsSet() {
		toSerialize["company"] = o.Company.Get()
	}
	if o.CompanyUrl.IsSet() {
		toSerialize["company_url"] = o.CompanyUrl.Get()
	}
	if o.Description.IsSet() {
		toSerialize["description"] = o.Description.Get()
	}
	if o.End.IsSet() {
		toSerialize["end"] = o.End.Get()
	}
	if o.Industries.IsSet() {
		toSerialize["industries"] = o.Industries.Get()
	}
	if o.Industry.IsSet() {
		toSerialize["industry"] = o.Industry.Get()
	}
	if o.Position.IsSet() {
		toSerialize["position"] = o.Position.Get()
	}
	if o.Start.IsSet() {
		toSerialize["start"] = o.Start.Get()
	}
	return toSerialize, nil
}

type NullableResumesResumeConditionFieldsExperienceFields struct {
	value *ResumesResumeConditionFieldsExperienceFields
	isSet bool
}

func (v NullableResumesResumeConditionFieldsExperienceFields) Get() *ResumesResumeConditionFieldsExperienceFields {
	return v.value
}

func (v *NullableResumesResumeConditionFieldsExperienceFields) Set(val *ResumesResumeConditionFieldsExperienceFields) {
	v.value = val
	v.isSet = true
}

func (v NullableResumesResumeConditionFieldsExperienceFields) IsSet() bool {
	return v.isSet
}

func (v *NullableResumesResumeConditionFieldsExperienceFields) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableResumesResumeConditionFieldsExperienceFields(val *ResumesResumeConditionFieldsExperienceFields) *NullableResumesResumeConditionFieldsExperienceFields {
	return &NullableResumesResumeConditionFieldsExperienceFields{value: val, isSet: true}
}

func (v NullableResumesResumeConditionFieldsExperienceFields) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableResumesResumeConditionFieldsExperienceFields) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


