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

// checks if the VacanciesArgumentItem type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &VacanciesArgumentItem{}

// VacanciesArgumentItem struct for VacanciesArgumentItem
type VacanciesArgumentItem struct {
	// Параметр поиска вакансии
	Argument string `json:"argument"`
	ClusterGroup NullableIncludesIdName `json:"cluster_group,omitempty"`
	// URL поиска вакансий, который получится, если перестать учитывать в поиске данный параметр
	DisableUrl string `json:"disable_url"`
	// Цвет линии в HEX-формате `RRGGBB` (от `000000` до `FFFFFF`). Возвращается только для аргумента `metro`
	HexColor NullableString `json:"hex_color,omitempty"`
	// Станция или линия метро (`station`/`line`). Возвращается только для аргумента `metro`
	MetroType NullableString `json:"metro_type,omitempty"`
	// Название значения
	Name NullableString `json:"name,omitempty"`
	// Значение параметра
	Value string `json:"value"`
	// Описание параметра
	ValueDescription NullableString `json:"value_description,omitempty"`
}

type _VacanciesArgumentItem VacanciesArgumentItem

// NewVacanciesArgumentItem instantiates a new VacanciesArgumentItem object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewVacanciesArgumentItem(argument string, disableUrl string, value string) *VacanciesArgumentItem {
	this := VacanciesArgumentItem{}
	this.Argument = argument
	this.DisableUrl = disableUrl
	this.Value = value
	return &this
}

// NewVacanciesArgumentItemWithDefaults instantiates a new VacanciesArgumentItem object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewVacanciesArgumentItemWithDefaults() *VacanciesArgumentItem {
	this := VacanciesArgumentItem{}
	return &this
}

// GetArgument returns the Argument field value
func (o *VacanciesArgumentItem) GetArgument() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Argument
}

// GetArgumentOk returns a tuple with the Argument field value
// and a boolean to check if the value has been set.
func (o *VacanciesArgumentItem) GetArgumentOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Argument, true
}

// SetArgument sets field value
func (o *VacanciesArgumentItem) SetArgument(v string) {
	o.Argument = v
}

// GetClusterGroup returns the ClusterGroup field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *VacanciesArgumentItem) GetClusterGroup() IncludesIdName {
	if o == nil || IsNil(o.ClusterGroup.Get()) {
		var ret IncludesIdName
		return ret
	}
	return *o.ClusterGroup.Get()
}

// GetClusterGroupOk returns a tuple with the ClusterGroup field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *VacanciesArgumentItem) GetClusterGroupOk() (*IncludesIdName, bool) {
	if o == nil {
		return nil, false
	}
	return o.ClusterGroup.Get(), o.ClusterGroup.IsSet()
}

// HasClusterGroup returns a boolean if a field has been set.
func (o *VacanciesArgumentItem) HasClusterGroup() bool {
	if o != nil && o.ClusterGroup.IsSet() {
		return true
	}

	return false
}

// SetClusterGroup gets a reference to the given NullableIncludesIdName and assigns it to the ClusterGroup field.
func (o *VacanciesArgumentItem) SetClusterGroup(v IncludesIdName) {
	o.ClusterGroup.Set(&v)
}
// SetClusterGroupNil sets the value for ClusterGroup to be an explicit nil
func (o *VacanciesArgumentItem) SetClusterGroupNil() {
	o.ClusterGroup.Set(nil)
}

// UnsetClusterGroup ensures that no value is present for ClusterGroup, not even an explicit nil
func (o *VacanciesArgumentItem) UnsetClusterGroup() {
	o.ClusterGroup.Unset()
}

// GetDisableUrl returns the DisableUrl field value
func (o *VacanciesArgumentItem) GetDisableUrl() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.DisableUrl
}

// GetDisableUrlOk returns a tuple with the DisableUrl field value
// and a boolean to check if the value has been set.
func (o *VacanciesArgumentItem) GetDisableUrlOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.DisableUrl, true
}

// SetDisableUrl sets field value
func (o *VacanciesArgumentItem) SetDisableUrl(v string) {
	o.DisableUrl = v
}

// GetHexColor returns the HexColor field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *VacanciesArgumentItem) GetHexColor() string {
	if o == nil || IsNil(o.HexColor.Get()) {
		var ret string
		return ret
	}
	return *o.HexColor.Get()
}

// GetHexColorOk returns a tuple with the HexColor field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *VacanciesArgumentItem) GetHexColorOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.HexColor.Get(), o.HexColor.IsSet()
}

// HasHexColor returns a boolean if a field has been set.
func (o *VacanciesArgumentItem) HasHexColor() bool {
	if o != nil && o.HexColor.IsSet() {
		return true
	}

	return false
}

// SetHexColor gets a reference to the given NullableString and assigns it to the HexColor field.
func (o *VacanciesArgumentItem) SetHexColor(v string) {
	o.HexColor.Set(&v)
}
// SetHexColorNil sets the value for HexColor to be an explicit nil
func (o *VacanciesArgumentItem) SetHexColorNil() {
	o.HexColor.Set(nil)
}

// UnsetHexColor ensures that no value is present for HexColor, not even an explicit nil
func (o *VacanciesArgumentItem) UnsetHexColor() {
	o.HexColor.Unset()
}

// GetMetroType returns the MetroType field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *VacanciesArgumentItem) GetMetroType() string {
	if o == nil || IsNil(o.MetroType.Get()) {
		var ret string
		return ret
	}
	return *o.MetroType.Get()
}

// GetMetroTypeOk returns a tuple with the MetroType field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *VacanciesArgumentItem) GetMetroTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.MetroType.Get(), o.MetroType.IsSet()
}

// HasMetroType returns a boolean if a field has been set.
func (o *VacanciesArgumentItem) HasMetroType() bool {
	if o != nil && o.MetroType.IsSet() {
		return true
	}

	return false
}

// SetMetroType gets a reference to the given NullableString and assigns it to the MetroType field.
func (o *VacanciesArgumentItem) SetMetroType(v string) {
	o.MetroType.Set(&v)
}
// SetMetroTypeNil sets the value for MetroType to be an explicit nil
func (o *VacanciesArgumentItem) SetMetroTypeNil() {
	o.MetroType.Set(nil)
}

// UnsetMetroType ensures that no value is present for MetroType, not even an explicit nil
func (o *VacanciesArgumentItem) UnsetMetroType() {
	o.MetroType.Unset()
}

// GetName returns the Name field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *VacanciesArgumentItem) GetName() string {
	if o == nil || IsNil(o.Name.Get()) {
		var ret string
		return ret
	}
	return *o.Name.Get()
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *VacanciesArgumentItem) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Name.Get(), o.Name.IsSet()
}

// HasName returns a boolean if a field has been set.
func (o *VacanciesArgumentItem) HasName() bool {
	if o != nil && o.Name.IsSet() {
		return true
	}

	return false
}

// SetName gets a reference to the given NullableString and assigns it to the Name field.
func (o *VacanciesArgumentItem) SetName(v string) {
	o.Name.Set(&v)
}
// SetNameNil sets the value for Name to be an explicit nil
func (o *VacanciesArgumentItem) SetNameNil() {
	o.Name.Set(nil)
}

// UnsetName ensures that no value is present for Name, not even an explicit nil
func (o *VacanciesArgumentItem) UnsetName() {
	o.Name.Unset()
}

// GetValue returns the Value field value
func (o *VacanciesArgumentItem) GetValue() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Value
}

// GetValueOk returns a tuple with the Value field value
// and a boolean to check if the value has been set.
func (o *VacanciesArgumentItem) GetValueOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Value, true
}

// SetValue sets field value
func (o *VacanciesArgumentItem) SetValue(v string) {
	o.Value = v
}

// GetValueDescription returns the ValueDescription field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *VacanciesArgumentItem) GetValueDescription() string {
	if o == nil || IsNil(o.ValueDescription.Get()) {
		var ret string
		return ret
	}
	return *o.ValueDescription.Get()
}

// GetValueDescriptionOk returns a tuple with the ValueDescription field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *VacanciesArgumentItem) GetValueDescriptionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.ValueDescription.Get(), o.ValueDescription.IsSet()
}

// HasValueDescription returns a boolean if a field has been set.
func (o *VacanciesArgumentItem) HasValueDescription() bool {
	if o != nil && o.ValueDescription.IsSet() {
		return true
	}

	return false
}

// SetValueDescription gets a reference to the given NullableString and assigns it to the ValueDescription field.
func (o *VacanciesArgumentItem) SetValueDescription(v string) {
	o.ValueDescription.Set(&v)
}
// SetValueDescriptionNil sets the value for ValueDescription to be an explicit nil
func (o *VacanciesArgumentItem) SetValueDescriptionNil() {
	o.ValueDescription.Set(nil)
}

// UnsetValueDescription ensures that no value is present for ValueDescription, not even an explicit nil
func (o *VacanciesArgumentItem) UnsetValueDescription() {
	o.ValueDescription.Unset()
}

func (o VacanciesArgumentItem) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o VacanciesArgumentItem) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["argument"] = o.Argument
	if o.ClusterGroup.IsSet() {
		toSerialize["cluster_group"] = o.ClusterGroup.Get()
	}
	toSerialize["disable_url"] = o.DisableUrl
	if o.HexColor.IsSet() {
		toSerialize["hex_color"] = o.HexColor.Get()
	}
	if o.MetroType.IsSet() {
		toSerialize["metro_type"] = o.MetroType.Get()
	}
	if o.Name.IsSet() {
		toSerialize["name"] = o.Name.Get()
	}
	toSerialize["value"] = o.Value
	if o.ValueDescription.IsSet() {
		toSerialize["value_description"] = o.ValueDescription.Get()
	}
	return toSerialize, nil
}

func (o *VacanciesArgumentItem) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"argument",
		"disable_url",
		"value",
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

	varVacanciesArgumentItem := _VacanciesArgumentItem{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varVacanciesArgumentItem)

	if err != nil {
		return err
	}

	*o = VacanciesArgumentItem(varVacanciesArgumentItem)

	return err
}

type NullableVacanciesArgumentItem struct {
	value *VacanciesArgumentItem
	isSet bool
}

func (v NullableVacanciesArgumentItem) Get() *VacanciesArgumentItem {
	return v.value
}

func (v *NullableVacanciesArgumentItem) Set(val *VacanciesArgumentItem) {
	v.value = val
	v.isSet = true
}

func (v NullableVacanciesArgumentItem) IsSet() bool {
	return v.isSet
}

func (v *NullableVacanciesArgumentItem) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableVacanciesArgumentItem(val *VacanciesArgumentItem) *NullableVacanciesArgumentItem {
	return &NullableVacanciesArgumentItem{value: val, isSet: true}
}

func (v NullableVacanciesArgumentItem) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableVacanciesArgumentItem) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


