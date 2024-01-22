/*
HeadHunter API

По-русски | [Switch to English](https://api.hh.ru/openapi/en/redoc)  В OpenAPI ведется пока что только небольшая часть документации [Основная документация](https://github.com/hhru/api).  Для поиска по документации можно использовать Ctrl+F.  # Авторизация  API поддерживает следующие уровни авторизации:   - [авторизация приложения](#section/Avtorizaciya/Avtorizaciya-prilozheniya)   - [авторизация пользователя](#section/Avtorizaciya/Avtorizaciya-polzovatelya)  * [Авторизация пользователя](#section/Avtorizaciya/Avtorizaciya-polzovatelya)   * [Правила формирования специального redirect_uri](#section/Avtorizaciya/Pravila-formirovaniya-specialnogo-redirect_uri)   * [Процесс авторизации](#section/Avtorizaciya/Process-avtorizacii)   * [Успешное получение временного `authorization_code`](#get-authorization_code)   * [Получение access и refresh токенов](#section/Avtorizaciya/Poluchenie-access-i-refresh-tokenov) * [Обновление пары access и refresh токенов](#section/Avtorizaciya/Obnovlenie-pary-access-i-refresh-tokenov) * [Инвалидация токена](#tag/Avtorizaciya-rabotodatelya/operation/invalidate-token) * [Запрос авторизации под другим пользователем](#section/Avtorizaciya/Zapros-avtorizacii-pod-drugim-polzovatelem) * [Авторизация под разными рабочими аккаунтами](#section/Avtorizaciya/Avtorizaciya-pod-raznymi-rabochimi-akkauntami) * [Авторизация приложения](#section/Avtorizaciya/Avtorizaciya-prilozheniya)       ## Авторизация пользователя Для выполнения запросов от имени пользователя необходимо пользоваться токеном пользователя.  В начале приложению необходимо направить пользователя (открыть страницу) по адресу:  ``` https://hh.ru/oauth/authorize? response_type=code& client_id={client_id}& state={state}& redirect_uri={redirect_uri} ```  Обязательные параметры:  * `response_type=code` — указание на способ получения авторизации, используя `authorization code` * `client_id` — идентификатор, полученный при создании приложения   Необязательные параметры:  * `state` — в случае указания, будет включен в ответный редирект. Это позволяет исключить возможность взлома путём подделки межсайтовых запросов. Подробнее об этом: [RFC 6749. Section 10.12](http://tools.ietf.org/html/rfc6749#section-10.12) * `redirect_uri` — uri для перенаправления пользователя после авторизации. Если не указать, используется из настроек приложения. При наличии происходит валидация значения. Вероятнее всего, потребуется сделать urlencode значения параметра.  ## Правила формирования специального redirect_uri  К примеру, если в настройках сохранен `http://example.com/oauth`, то разрешено указывать:  * `http://www.example.com/oauth` — поддомен; * `http://www.example.com/oauth/sub/path` — уточнение пути; * `http://example.com/oauth?lang=RU` — дополнительный параметр; * `http://www.example.com/oauth/sub/path?lang=RU` — всё вместе.  Запрещено:  * `https://example.com/oauth` — различные протоколы; * `http://wwwexample.com/oauth` — различные домены; * `http://wwwexample.com/` — другой путь; * `http://example.com/oauths` — другой путь; * `http://example.com:80/oauths` — указание изначально отсутствующего порта;  ## Процесс авторизации  Если пользователь не авторизован на сайте, ему будет показана форма авторизации на сайте. После прохождения авторизации на сайте, пользователю будет выведена форма с запросом разрешения доступа вашего приложения к его персональным данным.  Если пользователь не разрешает доступ приложению, пользователь будет перенаправлен на указанный `redirect_uri` с `?error=access_denied` и `state={state}`, если таковой был указан при первом запросе.  <a name=\"get-authorization_code\"></a> ### Успешное получение временного `authorization_code`  В случае разрешения прав, в редиректе будет указан временный `authorization_code`:  ```http HTTP/1.1 302 FOUND Location: {redirect_uri}?code={authorization_code} ```  Если пользователь авторизован на сайте и доступ данному приложению однажды ранее выдан, ответом будет сразу вышеописанный редирект с `authorization_code` (без показа формы логина и выдачи прав).  ## Получение access и refresh токенов  После получения `authorization_code` приложению необходимо сделать сервер-сервер запрос `POST https://hh.ru/oauth/token` для обмена полученного `authorization_code` на `access_token`.  В теле запроса необходимо передать [дополнительные параметры](#required_parameters).  Тело запроса необходимо передавать в стандартном `application/x-www-form-urlencoded` с указанием соответствующего заголовка `Content-Type`.  `authorization_code` имеет довольно короткий срок жизни, при его истечении необходимо запросить новый.  ## Обновление пары access и refresh токенов `access_token` также имеет срок жизни (ключ `expires_in`, в секундах), при его истечении приложение должно сделать запрос с `refresh_token` для получения нового.  Запрос необходимо делать в `application/x-www-form-urlencoded`.  ``` POST https://hh.ru/oauth/token ```  В теле запроса необходимо передать [дополнительные параметры](#required_parameters)  `refresh_token` можно использовать только один раз и только по истечению срока действия `access_token`.  После получения новой пары access и refresh токенов, их необходимо использовать в дальнейших запросах в api и запросах на продление токена.  ## Запрос авторизации под другим пользователем  Возможен следующий сценарий:  1. Приложение перенаправляет пользователя на сайт с запросом авторизации. 2. Пользователь на сайте уже авторизован и данному приложение доступ уже был разрешен. 3. Пользователю будет предложена возможность продолжить работу под текущим аккаунтом, либо зайти под другим аккаунтом.  Если есть необходимость, чтобы на шаге 3 сразу происходило перенаправление (redirect) с временным токеном, необходимо добавить к запросу `/oauth/authorize...` параметр `skip_choose_account=true`. В этом случае автоматически выдаётся доступ пользователю авторизованному на сайте.  Если есть необходимость всегда показывать форму авторизации, приложение может добавить к запросу `/oauth/authorize...` параметр `force_login=true`. В этом случае, пользователю будет показана форма авторизации с логином и паролем даже в случае, если пользователь уже авторизован.  Это может быть полезно приложениям, которые предоставляют сервис только для соискателей. Если пришел пользователь-работодатель, приложение может предложить пользователю повторно разрешить доступ на сайте, уже указав другую учетную запись.  Также, после авторизации приложение может показать пользователю сообщение:  ``` Вы вошли как %Имя_Фамилия%. Это не вы? ``` и предоставить ссылку с `force_login=true` для возможности захода под другим логином.  ## Авторизация под разными рабочими аккаунтами  Для получения списка рабочих аккаунтов менеджера и для работы под разными рабочими аккаунтами менеджера необходимо прочитать документацию по [рабочим аккаунтам менеджера](#tag/Menedzhery-rabotodatelya/operation/get-manager-accounts)  ## Авторизация приложения  Токен приложения необходимо сгенерировать 1 раз. В случае, если токен был скомпрометирован, его нужно запросить еще раз. При этом ранее выданный токен отзывается. Владелец приложения может посмотреть актуальный `access_token` для приложения на сайте [https://dev.hh.ru/admin](https://dev.hh.ru/admin). В случае, если вы еще ни разу [не получали токен приложения](#section/Avtorizaciya/Avtorizaciya-prilozheniya), токен отображаться не будет.  <a name=\"get-client-token\"></a> ### Получение токена приложения Для получения `access_token` необходимо сделать запрос:  ``` POST https://hh.ru/oauth/token ```  В теле запроса необходимо передать [дополнительные параметры](#required_parameters). Тело запроса необходимо передавать в стандартном `application/x-www-form-urlencoded` с указанием соответствующего заголовка `Content-Type`.  Данный `access_token` имеет **неограниченный** срок жизни. При повторном запросе ранее выданный токен отзывается и выдается новый. Запрашивать `access_token` можно не чаще, чем один раз в 5 минут.  В случае компрометации токена необходимо [инвалидировать](#tag/Avtorizaciya-rabotodatelya/operation/invalidate-token) скомпроментированный токен и запросить токен заново!  <!-- ReDoc-Inject: <security-definitions> --> 

API version: 1.0.0
Contact: api@hh.ru
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package hh-go

import (
	"encoding/json"
)

// checks if the VacancyAddressOutput type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &VacancyAddressOutput{}

// VacancyAddressOutput Адрес
type VacancyAddressOutput struct {
	// Дом
	Building NullableString `json:"building,omitempty"`
	// Город
	City NullableString `json:"city,omitempty"`
	// Широта
	Lat NullableFloat32 `json:"lat,omitempty"`
	// Долгота
	Lng NullableFloat32 `json:"lng,omitempty"`
	// Улица
	Street NullableString `json:"street,omitempty"`
	// Описание
	Description NullableString `json:"description,omitempty"`
	MetroStations []IncludesMetroStation `json:"metro_stations,omitempty"`
}

// NewVacancyAddressOutput instantiates a new VacancyAddressOutput object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewVacancyAddressOutput() *VacancyAddressOutput {
	this := VacancyAddressOutput{}
	return &this
}

// NewVacancyAddressOutputWithDefaults instantiates a new VacancyAddressOutput object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewVacancyAddressOutputWithDefaults() *VacancyAddressOutput {
	this := VacancyAddressOutput{}
	return &this
}

// GetBuilding returns the Building field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *VacancyAddressOutput) GetBuilding() string {
	if o == nil || IsNil(o.Building.Get()) {
		var ret string
		return ret
	}
	return *o.Building.Get()
}

// GetBuildingOk returns a tuple with the Building field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *VacancyAddressOutput) GetBuildingOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Building.Get(), o.Building.IsSet()
}

// HasBuilding returns a boolean if a field has been set.
func (o *VacancyAddressOutput) HasBuilding() bool {
	if o != nil && o.Building.IsSet() {
		return true
	}

	return false
}

// SetBuilding gets a reference to the given NullableString and assigns it to the Building field.
func (o *VacancyAddressOutput) SetBuilding(v string) {
	o.Building.Set(&v)
}
// SetBuildingNil sets the value for Building to be an explicit nil
func (o *VacancyAddressOutput) SetBuildingNil() {
	o.Building.Set(nil)
}

// UnsetBuilding ensures that no value is present for Building, not even an explicit nil
func (o *VacancyAddressOutput) UnsetBuilding() {
	o.Building.Unset()
}

// GetCity returns the City field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *VacancyAddressOutput) GetCity() string {
	if o == nil || IsNil(o.City.Get()) {
		var ret string
		return ret
	}
	return *o.City.Get()
}

// GetCityOk returns a tuple with the City field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *VacancyAddressOutput) GetCityOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.City.Get(), o.City.IsSet()
}

// HasCity returns a boolean if a field has been set.
func (o *VacancyAddressOutput) HasCity() bool {
	if o != nil && o.City.IsSet() {
		return true
	}

	return false
}

// SetCity gets a reference to the given NullableString and assigns it to the City field.
func (o *VacancyAddressOutput) SetCity(v string) {
	o.City.Set(&v)
}
// SetCityNil sets the value for City to be an explicit nil
func (o *VacancyAddressOutput) SetCityNil() {
	o.City.Set(nil)
}

// UnsetCity ensures that no value is present for City, not even an explicit nil
func (o *VacancyAddressOutput) UnsetCity() {
	o.City.Unset()
}

// GetLat returns the Lat field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *VacancyAddressOutput) GetLat() float32 {
	if o == nil || IsNil(o.Lat.Get()) {
		var ret float32
		return ret
	}
	return *o.Lat.Get()
}

// GetLatOk returns a tuple with the Lat field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *VacancyAddressOutput) GetLatOk() (*float32, bool) {
	if o == nil {
		return nil, false
	}
	return o.Lat.Get(), o.Lat.IsSet()
}

// HasLat returns a boolean if a field has been set.
func (o *VacancyAddressOutput) HasLat() bool {
	if o != nil && o.Lat.IsSet() {
		return true
	}

	return false
}

// SetLat gets a reference to the given NullableFloat32 and assigns it to the Lat field.
func (o *VacancyAddressOutput) SetLat(v float32) {
	o.Lat.Set(&v)
}
// SetLatNil sets the value for Lat to be an explicit nil
func (o *VacancyAddressOutput) SetLatNil() {
	o.Lat.Set(nil)
}

// UnsetLat ensures that no value is present for Lat, not even an explicit nil
func (o *VacancyAddressOutput) UnsetLat() {
	o.Lat.Unset()
}

// GetLng returns the Lng field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *VacancyAddressOutput) GetLng() float32 {
	if o == nil || IsNil(o.Lng.Get()) {
		var ret float32
		return ret
	}
	return *o.Lng.Get()
}

// GetLngOk returns a tuple with the Lng field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *VacancyAddressOutput) GetLngOk() (*float32, bool) {
	if o == nil {
		return nil, false
	}
	return o.Lng.Get(), o.Lng.IsSet()
}

// HasLng returns a boolean if a field has been set.
func (o *VacancyAddressOutput) HasLng() bool {
	if o != nil && o.Lng.IsSet() {
		return true
	}

	return false
}

// SetLng gets a reference to the given NullableFloat32 and assigns it to the Lng field.
func (o *VacancyAddressOutput) SetLng(v float32) {
	o.Lng.Set(&v)
}
// SetLngNil sets the value for Lng to be an explicit nil
func (o *VacancyAddressOutput) SetLngNil() {
	o.Lng.Set(nil)
}

// UnsetLng ensures that no value is present for Lng, not even an explicit nil
func (o *VacancyAddressOutput) UnsetLng() {
	o.Lng.Unset()
}

// GetStreet returns the Street field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *VacancyAddressOutput) GetStreet() string {
	if o == nil || IsNil(o.Street.Get()) {
		var ret string
		return ret
	}
	return *o.Street.Get()
}

// GetStreetOk returns a tuple with the Street field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *VacancyAddressOutput) GetStreetOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Street.Get(), o.Street.IsSet()
}

// HasStreet returns a boolean if a field has been set.
func (o *VacancyAddressOutput) HasStreet() bool {
	if o != nil && o.Street.IsSet() {
		return true
	}

	return false
}

// SetStreet gets a reference to the given NullableString and assigns it to the Street field.
func (o *VacancyAddressOutput) SetStreet(v string) {
	o.Street.Set(&v)
}
// SetStreetNil sets the value for Street to be an explicit nil
func (o *VacancyAddressOutput) SetStreetNil() {
	o.Street.Set(nil)
}

// UnsetStreet ensures that no value is present for Street, not even an explicit nil
func (o *VacancyAddressOutput) UnsetStreet() {
	o.Street.Unset()
}

// GetDescription returns the Description field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *VacancyAddressOutput) GetDescription() string {
	if o == nil || IsNil(o.Description.Get()) {
		var ret string
		return ret
	}
	return *o.Description.Get()
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *VacancyAddressOutput) GetDescriptionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Description.Get(), o.Description.IsSet()
}

// HasDescription returns a boolean if a field has been set.
func (o *VacancyAddressOutput) HasDescription() bool {
	if o != nil && o.Description.IsSet() {
		return true
	}

	return false
}

// SetDescription gets a reference to the given NullableString and assigns it to the Description field.
func (o *VacancyAddressOutput) SetDescription(v string) {
	o.Description.Set(&v)
}
// SetDescriptionNil sets the value for Description to be an explicit nil
func (o *VacancyAddressOutput) SetDescriptionNil() {
	o.Description.Set(nil)
}

// UnsetDescription ensures that no value is present for Description, not even an explicit nil
func (o *VacancyAddressOutput) UnsetDescription() {
	o.Description.Unset()
}

// GetMetroStations returns the MetroStations field value if set, zero value otherwise.
func (o *VacancyAddressOutput) GetMetroStations() []IncludesMetroStation {
	if o == nil || IsNil(o.MetroStations) {
		var ret []IncludesMetroStation
		return ret
	}
	return o.MetroStations
}

// GetMetroStationsOk returns a tuple with the MetroStations field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VacancyAddressOutput) GetMetroStationsOk() ([]IncludesMetroStation, bool) {
	if o == nil || IsNil(o.MetroStations) {
		return nil, false
	}
	return o.MetroStations, true
}

// HasMetroStations returns a boolean if a field has been set.
func (o *VacancyAddressOutput) HasMetroStations() bool {
	if o != nil && !IsNil(o.MetroStations) {
		return true
	}

	return false
}

// SetMetroStations gets a reference to the given []IncludesMetroStation and assigns it to the MetroStations field.
func (o *VacancyAddressOutput) SetMetroStations(v []IncludesMetroStation) {
	o.MetroStations = v
}

func (o VacancyAddressOutput) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o VacancyAddressOutput) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.Building.IsSet() {
		toSerialize["building"] = o.Building.Get()
	}
	if o.City.IsSet() {
		toSerialize["city"] = o.City.Get()
	}
	if o.Lat.IsSet() {
		toSerialize["lat"] = o.Lat.Get()
	}
	if o.Lng.IsSet() {
		toSerialize["lng"] = o.Lng.Get()
	}
	if o.Street.IsSet() {
		toSerialize["street"] = o.Street.Get()
	}
	if o.Description.IsSet() {
		toSerialize["description"] = o.Description.Get()
	}
	if !IsNil(o.MetroStations) {
		toSerialize["metro_stations"] = o.MetroStations
	}
	return toSerialize, nil
}

type NullableVacancyAddressOutput struct {
	value *VacancyAddressOutput
	isSet bool
}

func (v NullableVacancyAddressOutput) Get() *VacancyAddressOutput {
	return v.value
}

func (v *NullableVacancyAddressOutput) Set(val *VacancyAddressOutput) {
	v.value = val
	v.isSet = true
}

func (v NullableVacancyAddressOutput) IsSet() bool {
	return v.isSet
}

func (v *NullableVacancyAddressOutput) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableVacancyAddressOutput(val *VacancyAddressOutput) *NullableVacancyAddressOutput {
	return &NullableVacancyAddressOutput{value: val, isSet: true}
}

func (v NullableVacancyAddressOutput) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableVacancyAddressOutput) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


