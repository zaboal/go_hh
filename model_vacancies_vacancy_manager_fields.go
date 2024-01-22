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

// checks if the VacanciesVacancyManagerFields type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &VacanciesVacancyManagerFields{}

// VacanciesVacancyManagerFields struct for VacanciesVacancyManagerFields
type VacanciesVacancyManagerFields struct {
	Address VacanciesAddress `json:"address"`
	// Дата архивации вакансии
	ArchivedAt *string `json:"archived_at,omitempty"`
	BrandedTemplate NullableVacancyBrandedTemplate `json:"branded_template"`
	// Можно ли улучшить биллинговый тип вакансии
	CanUpgradeBillingType bool `json:"can_upgrade_billing_type"`
	Counters *VacancyCountersOutput `json:"counters,omitempty"`
	// Дата и время окончания публикации вакансии
	ExpiresAt string `json:"expires_at"`
	// Удалена ли вакансия (скрыта из архива)
	Hidden bool `json:"hidden"`
	Manager VacancyManager `json:"manager"`
	// Уведомлять ли менеджера о новых откликах
	ResponseNotifications bool `json:"response_notifications"`
}

type _VacanciesVacancyManagerFields VacanciesVacancyManagerFields

// NewVacanciesVacancyManagerFields instantiates a new VacanciesVacancyManagerFields object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewVacanciesVacancyManagerFields(address VacanciesAddress, brandedTemplate NullableVacancyBrandedTemplate, canUpgradeBillingType bool, expiresAt string, hidden bool, manager VacancyManager, responseNotifications bool) *VacanciesVacancyManagerFields {
	this := VacanciesVacancyManagerFields{}
	this.Address = address
	this.BrandedTemplate = brandedTemplate
	this.CanUpgradeBillingType = canUpgradeBillingType
	this.ExpiresAt = expiresAt
	this.Hidden = hidden
	this.Manager = manager
	this.ResponseNotifications = responseNotifications
	return &this
}

// NewVacanciesVacancyManagerFieldsWithDefaults instantiates a new VacanciesVacancyManagerFields object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewVacanciesVacancyManagerFieldsWithDefaults() *VacanciesVacancyManagerFields {
	this := VacanciesVacancyManagerFields{}
	return &this
}

// GetAddress returns the Address field value
func (o *VacanciesVacancyManagerFields) GetAddress() VacanciesAddress {
	if o == nil {
		var ret VacanciesAddress
		return ret
	}

	return o.Address
}

// GetAddressOk returns a tuple with the Address field value
// and a boolean to check if the value has been set.
func (o *VacanciesVacancyManagerFields) GetAddressOk() (*VacanciesAddress, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Address, true
}

// SetAddress sets field value
func (o *VacanciesVacancyManagerFields) SetAddress(v VacanciesAddress) {
	o.Address = v
}

// GetArchivedAt returns the ArchivedAt field value if set, zero value otherwise.
func (o *VacanciesVacancyManagerFields) GetArchivedAt() string {
	if o == nil || IsNil(o.ArchivedAt) {
		var ret string
		return ret
	}
	return *o.ArchivedAt
}

// GetArchivedAtOk returns a tuple with the ArchivedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VacanciesVacancyManagerFields) GetArchivedAtOk() (*string, bool) {
	if o == nil || IsNil(o.ArchivedAt) {
		return nil, false
	}
	return o.ArchivedAt, true
}

// HasArchivedAt returns a boolean if a field has been set.
func (o *VacanciesVacancyManagerFields) HasArchivedAt() bool {
	if o != nil && !IsNil(o.ArchivedAt) {
		return true
	}

	return false
}

// SetArchivedAt gets a reference to the given string and assigns it to the ArchivedAt field.
func (o *VacanciesVacancyManagerFields) SetArchivedAt(v string) {
	o.ArchivedAt = &v
}

// GetBrandedTemplate returns the BrandedTemplate field value
// If the value is explicit nil, the zero value for VacancyBrandedTemplate will be returned
func (o *VacanciesVacancyManagerFields) GetBrandedTemplate() VacancyBrandedTemplate {
	if o == nil || o.BrandedTemplate.Get() == nil {
		var ret VacancyBrandedTemplate
		return ret
	}

	return *o.BrandedTemplate.Get()
}

// GetBrandedTemplateOk returns a tuple with the BrandedTemplate field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *VacanciesVacancyManagerFields) GetBrandedTemplateOk() (*VacancyBrandedTemplate, bool) {
	if o == nil {
		return nil, false
	}
	return o.BrandedTemplate.Get(), o.BrandedTemplate.IsSet()
}

// SetBrandedTemplate sets field value
func (o *VacanciesVacancyManagerFields) SetBrandedTemplate(v VacancyBrandedTemplate) {
	o.BrandedTemplate.Set(&v)
}

// GetCanUpgradeBillingType returns the CanUpgradeBillingType field value
func (o *VacanciesVacancyManagerFields) GetCanUpgradeBillingType() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.CanUpgradeBillingType
}

// GetCanUpgradeBillingTypeOk returns a tuple with the CanUpgradeBillingType field value
// and a boolean to check if the value has been set.
func (o *VacanciesVacancyManagerFields) GetCanUpgradeBillingTypeOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CanUpgradeBillingType, true
}

// SetCanUpgradeBillingType sets field value
func (o *VacanciesVacancyManagerFields) SetCanUpgradeBillingType(v bool) {
	o.CanUpgradeBillingType = v
}

// GetCounters returns the Counters field value if set, zero value otherwise.
func (o *VacanciesVacancyManagerFields) GetCounters() VacancyCountersOutput {
	if o == nil || IsNil(o.Counters) {
		var ret VacancyCountersOutput
		return ret
	}
	return *o.Counters
}

// GetCountersOk returns a tuple with the Counters field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VacanciesVacancyManagerFields) GetCountersOk() (*VacancyCountersOutput, bool) {
	if o == nil || IsNil(o.Counters) {
		return nil, false
	}
	return o.Counters, true
}

// HasCounters returns a boolean if a field has been set.
func (o *VacanciesVacancyManagerFields) HasCounters() bool {
	if o != nil && !IsNil(o.Counters) {
		return true
	}

	return false
}

// SetCounters gets a reference to the given VacancyCountersOutput and assigns it to the Counters field.
func (o *VacanciesVacancyManagerFields) SetCounters(v VacancyCountersOutput) {
	o.Counters = &v
}

// GetExpiresAt returns the ExpiresAt field value
func (o *VacanciesVacancyManagerFields) GetExpiresAt() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ExpiresAt
}

// GetExpiresAtOk returns a tuple with the ExpiresAt field value
// and a boolean to check if the value has been set.
func (o *VacanciesVacancyManagerFields) GetExpiresAtOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ExpiresAt, true
}

// SetExpiresAt sets field value
func (o *VacanciesVacancyManagerFields) SetExpiresAt(v string) {
	o.ExpiresAt = v
}

// GetHidden returns the Hidden field value
func (o *VacanciesVacancyManagerFields) GetHidden() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.Hidden
}

// GetHiddenOk returns a tuple with the Hidden field value
// and a boolean to check if the value has been set.
func (o *VacanciesVacancyManagerFields) GetHiddenOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Hidden, true
}

// SetHidden sets field value
func (o *VacanciesVacancyManagerFields) SetHidden(v bool) {
	o.Hidden = v
}

// GetManager returns the Manager field value
func (o *VacanciesVacancyManagerFields) GetManager() VacancyManager {
	if o == nil {
		var ret VacancyManager
		return ret
	}

	return o.Manager
}

// GetManagerOk returns a tuple with the Manager field value
// and a boolean to check if the value has been set.
func (o *VacanciesVacancyManagerFields) GetManagerOk() (*VacancyManager, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Manager, true
}

// SetManager sets field value
func (o *VacanciesVacancyManagerFields) SetManager(v VacancyManager) {
	o.Manager = v
}

// GetResponseNotifications returns the ResponseNotifications field value
func (o *VacanciesVacancyManagerFields) GetResponseNotifications() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.ResponseNotifications
}

// GetResponseNotificationsOk returns a tuple with the ResponseNotifications field value
// and a boolean to check if the value has been set.
func (o *VacanciesVacancyManagerFields) GetResponseNotificationsOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ResponseNotifications, true
}

// SetResponseNotifications sets field value
func (o *VacanciesVacancyManagerFields) SetResponseNotifications(v bool) {
	o.ResponseNotifications = v
}

func (o VacanciesVacancyManagerFields) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o VacanciesVacancyManagerFields) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["address"] = o.Address
	if !IsNil(o.ArchivedAt) {
		toSerialize["archived_at"] = o.ArchivedAt
	}
	toSerialize["branded_template"] = o.BrandedTemplate.Get()
	toSerialize["can_upgrade_billing_type"] = o.CanUpgradeBillingType
	if !IsNil(o.Counters) {
		toSerialize["counters"] = o.Counters
	}
	toSerialize["expires_at"] = o.ExpiresAt
	toSerialize["hidden"] = o.Hidden
	toSerialize["manager"] = o.Manager
	toSerialize["response_notifications"] = o.ResponseNotifications
	return toSerialize, nil
}

func (o *VacanciesVacancyManagerFields) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"address",
		"branded_template",
		"can_upgrade_billing_type",
		"expires_at",
		"hidden",
		"manager",
		"response_notifications",
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

	varVacanciesVacancyManagerFields := _VacanciesVacancyManagerFields{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varVacanciesVacancyManagerFields)

	if err != nil {
		return err
	}

	*o = VacanciesVacancyManagerFields(varVacanciesVacancyManagerFields)

	return err
}

type NullableVacanciesVacancyManagerFields struct {
	value *VacanciesVacancyManagerFields
	isSet bool
}

func (v NullableVacanciesVacancyManagerFields) Get() *VacanciesVacancyManagerFields {
	return v.value
}

func (v *NullableVacanciesVacancyManagerFields) Set(val *VacanciesVacancyManagerFields) {
	v.value = val
	v.isSet = true
}

func (v NullableVacanciesVacancyManagerFields) IsSet() bool {
	return v.isSet
}

func (v *NullableVacanciesVacancyManagerFields) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableVacanciesVacancyManagerFields(val *VacanciesVacancyManagerFields) *NullableVacanciesVacancyManagerFields {
	return &NullableVacanciesVacancyManagerFields{value: val, isSet: true}
}

func (v NullableVacanciesVacancyManagerFields) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableVacanciesVacancyManagerFields) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


