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

// checks if the EmployerAddressesEmployerAddressItemResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &EmployerAddressesEmployerAddressItemResponse{}

// EmployerAddressesEmployerAddressItemResponse struct for EmployerAddressesEmployerAddressItemResponse
type EmployerAddressesEmployerAddressItemResponse struct {
	// Номер дома
	Building NullableString `json:"building,omitempty"`
	// Имеет ли текущий пользователь право редактировать этот адрес
	CanEdit *bool `json:"can_edit,omitempty"`
	// Город
	City NullableString `json:"city"`
	// Удалён ли адрес
	Deleted *bool `json:"deleted,omitempty"`
	// Дополнительная информация об адресе
	Description NullableString `json:"description,omitempty"`
	// Идентификатор адреса
	Id *string `json:"id,omitempty"`
	// Географическая широта
	Lat NullableFloat32 `json:"lat"`
	// Географическая долгота
	Lng NullableFloat32 `json:"lng"`
	Manager NullableEmployerAddressesEmployerAddressItemManager `json:"manager,omitempty"`
	MetroStations []IncludesMetroStation `json:"metro_stations,omitempty"`
	// Полный адрес
	Raw NullableString `json:"raw,omitempty"`
	// Улица
	Street NullableString `json:"street,omitempty"`
}

type _EmployerAddressesEmployerAddressItemResponse EmployerAddressesEmployerAddressItemResponse

// NewEmployerAddressesEmployerAddressItemResponse instantiates a new EmployerAddressesEmployerAddressItemResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewEmployerAddressesEmployerAddressItemResponse(city NullableString, lat NullableFloat32, lng NullableFloat32) *EmployerAddressesEmployerAddressItemResponse {
	this := EmployerAddressesEmployerAddressItemResponse{}
	this.City = city
	this.Lat = lat
	this.Lng = lng
	return &this
}

// NewEmployerAddressesEmployerAddressItemResponseWithDefaults instantiates a new EmployerAddressesEmployerAddressItemResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewEmployerAddressesEmployerAddressItemResponseWithDefaults() *EmployerAddressesEmployerAddressItemResponse {
	this := EmployerAddressesEmployerAddressItemResponse{}
	return &this
}

// GetBuilding returns the Building field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *EmployerAddressesEmployerAddressItemResponse) GetBuilding() string {
	if o == nil || IsNil(o.Building.Get()) {
		var ret string
		return ret
	}
	return *o.Building.Get()
}

// GetBuildingOk returns a tuple with the Building field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *EmployerAddressesEmployerAddressItemResponse) GetBuildingOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Building.Get(), o.Building.IsSet()
}

// HasBuilding returns a boolean if a field has been set.
func (o *EmployerAddressesEmployerAddressItemResponse) HasBuilding() bool {
	if o != nil && o.Building.IsSet() {
		return true
	}

	return false
}

// SetBuilding gets a reference to the given NullableString and assigns it to the Building field.
func (o *EmployerAddressesEmployerAddressItemResponse) SetBuilding(v string) {
	o.Building.Set(&v)
}
// SetBuildingNil sets the value for Building to be an explicit nil
func (o *EmployerAddressesEmployerAddressItemResponse) SetBuildingNil() {
	o.Building.Set(nil)
}

// UnsetBuilding ensures that no value is present for Building, not even an explicit nil
func (o *EmployerAddressesEmployerAddressItemResponse) UnsetBuilding() {
	o.Building.Unset()
}

// GetCanEdit returns the CanEdit field value if set, zero value otherwise.
func (o *EmployerAddressesEmployerAddressItemResponse) GetCanEdit() bool {
	if o == nil || IsNil(o.CanEdit) {
		var ret bool
		return ret
	}
	return *o.CanEdit
}

// GetCanEditOk returns a tuple with the CanEdit field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EmployerAddressesEmployerAddressItemResponse) GetCanEditOk() (*bool, bool) {
	if o == nil || IsNil(o.CanEdit) {
		return nil, false
	}
	return o.CanEdit, true
}

// HasCanEdit returns a boolean if a field has been set.
func (o *EmployerAddressesEmployerAddressItemResponse) HasCanEdit() bool {
	if o != nil && !IsNil(o.CanEdit) {
		return true
	}

	return false
}

// SetCanEdit gets a reference to the given bool and assigns it to the CanEdit field.
func (o *EmployerAddressesEmployerAddressItemResponse) SetCanEdit(v bool) {
	o.CanEdit = &v
}

// GetCity returns the City field value
// If the value is explicit nil, the zero value for string will be returned
func (o *EmployerAddressesEmployerAddressItemResponse) GetCity() string {
	if o == nil || o.City.Get() == nil {
		var ret string
		return ret
	}

	return *o.City.Get()
}

// GetCityOk returns a tuple with the City field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *EmployerAddressesEmployerAddressItemResponse) GetCityOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.City.Get(), o.City.IsSet()
}

// SetCity sets field value
func (o *EmployerAddressesEmployerAddressItemResponse) SetCity(v string) {
	o.City.Set(&v)
}

// GetDeleted returns the Deleted field value if set, zero value otherwise.
func (o *EmployerAddressesEmployerAddressItemResponse) GetDeleted() bool {
	if o == nil || IsNil(o.Deleted) {
		var ret bool
		return ret
	}
	return *o.Deleted
}

// GetDeletedOk returns a tuple with the Deleted field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EmployerAddressesEmployerAddressItemResponse) GetDeletedOk() (*bool, bool) {
	if o == nil || IsNil(o.Deleted) {
		return nil, false
	}
	return o.Deleted, true
}

// HasDeleted returns a boolean if a field has been set.
func (o *EmployerAddressesEmployerAddressItemResponse) HasDeleted() bool {
	if o != nil && !IsNil(o.Deleted) {
		return true
	}

	return false
}

// SetDeleted gets a reference to the given bool and assigns it to the Deleted field.
func (o *EmployerAddressesEmployerAddressItemResponse) SetDeleted(v bool) {
	o.Deleted = &v
}

// GetDescription returns the Description field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *EmployerAddressesEmployerAddressItemResponse) GetDescription() string {
	if o == nil || IsNil(o.Description.Get()) {
		var ret string
		return ret
	}
	return *o.Description.Get()
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *EmployerAddressesEmployerAddressItemResponse) GetDescriptionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Description.Get(), o.Description.IsSet()
}

// HasDescription returns a boolean if a field has been set.
func (o *EmployerAddressesEmployerAddressItemResponse) HasDescription() bool {
	if o != nil && o.Description.IsSet() {
		return true
	}

	return false
}

// SetDescription gets a reference to the given NullableString and assigns it to the Description field.
func (o *EmployerAddressesEmployerAddressItemResponse) SetDescription(v string) {
	o.Description.Set(&v)
}
// SetDescriptionNil sets the value for Description to be an explicit nil
func (o *EmployerAddressesEmployerAddressItemResponse) SetDescriptionNil() {
	o.Description.Set(nil)
}

// UnsetDescription ensures that no value is present for Description, not even an explicit nil
func (o *EmployerAddressesEmployerAddressItemResponse) UnsetDescription() {
	o.Description.Unset()
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *EmployerAddressesEmployerAddressItemResponse) GetId() string {
	if o == nil || IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EmployerAddressesEmployerAddressItemResponse) GetIdOk() (*string, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *EmployerAddressesEmployerAddressItemResponse) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *EmployerAddressesEmployerAddressItemResponse) SetId(v string) {
	o.Id = &v
}

// GetLat returns the Lat field value
// If the value is explicit nil, the zero value for float32 will be returned
func (o *EmployerAddressesEmployerAddressItemResponse) GetLat() float32 {
	if o == nil || o.Lat.Get() == nil {
		var ret float32
		return ret
	}

	return *o.Lat.Get()
}

// GetLatOk returns a tuple with the Lat field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *EmployerAddressesEmployerAddressItemResponse) GetLatOk() (*float32, bool) {
	if o == nil {
		return nil, false
	}
	return o.Lat.Get(), o.Lat.IsSet()
}

// SetLat sets field value
func (o *EmployerAddressesEmployerAddressItemResponse) SetLat(v float32) {
	o.Lat.Set(&v)
}

// GetLng returns the Lng field value
// If the value is explicit nil, the zero value for float32 will be returned
func (o *EmployerAddressesEmployerAddressItemResponse) GetLng() float32 {
	if o == nil || o.Lng.Get() == nil {
		var ret float32
		return ret
	}

	return *o.Lng.Get()
}

// GetLngOk returns a tuple with the Lng field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *EmployerAddressesEmployerAddressItemResponse) GetLngOk() (*float32, bool) {
	if o == nil {
		return nil, false
	}
	return o.Lng.Get(), o.Lng.IsSet()
}

// SetLng sets field value
func (o *EmployerAddressesEmployerAddressItemResponse) SetLng(v float32) {
	o.Lng.Set(&v)
}

// GetManager returns the Manager field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *EmployerAddressesEmployerAddressItemResponse) GetManager() EmployerAddressesEmployerAddressItemManager {
	if o == nil || IsNil(o.Manager.Get()) {
		var ret EmployerAddressesEmployerAddressItemManager
		return ret
	}
	return *o.Manager.Get()
}

// GetManagerOk returns a tuple with the Manager field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *EmployerAddressesEmployerAddressItemResponse) GetManagerOk() (*EmployerAddressesEmployerAddressItemManager, bool) {
	if o == nil {
		return nil, false
	}
	return o.Manager.Get(), o.Manager.IsSet()
}

// HasManager returns a boolean if a field has been set.
func (o *EmployerAddressesEmployerAddressItemResponse) HasManager() bool {
	if o != nil && o.Manager.IsSet() {
		return true
	}

	return false
}

// SetManager gets a reference to the given NullableEmployerAddressesEmployerAddressItemManager and assigns it to the Manager field.
func (o *EmployerAddressesEmployerAddressItemResponse) SetManager(v EmployerAddressesEmployerAddressItemManager) {
	o.Manager.Set(&v)
}
// SetManagerNil sets the value for Manager to be an explicit nil
func (o *EmployerAddressesEmployerAddressItemResponse) SetManagerNil() {
	o.Manager.Set(nil)
}

// UnsetManager ensures that no value is present for Manager, not even an explicit nil
func (o *EmployerAddressesEmployerAddressItemResponse) UnsetManager() {
	o.Manager.Unset()
}

// GetMetroStations returns the MetroStations field value if set, zero value otherwise.
func (o *EmployerAddressesEmployerAddressItemResponse) GetMetroStations() []IncludesMetroStation {
	if o == nil || IsNil(o.MetroStations) {
		var ret []IncludesMetroStation
		return ret
	}
	return o.MetroStations
}

// GetMetroStationsOk returns a tuple with the MetroStations field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EmployerAddressesEmployerAddressItemResponse) GetMetroStationsOk() ([]IncludesMetroStation, bool) {
	if o == nil || IsNil(o.MetroStations) {
		return nil, false
	}
	return o.MetroStations, true
}

// HasMetroStations returns a boolean if a field has been set.
func (o *EmployerAddressesEmployerAddressItemResponse) HasMetroStations() bool {
	if o != nil && !IsNil(o.MetroStations) {
		return true
	}

	return false
}

// SetMetroStations gets a reference to the given []IncludesMetroStation and assigns it to the MetroStations field.
func (o *EmployerAddressesEmployerAddressItemResponse) SetMetroStations(v []IncludesMetroStation) {
	o.MetroStations = v
}

// GetRaw returns the Raw field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *EmployerAddressesEmployerAddressItemResponse) GetRaw() string {
	if o == nil || IsNil(o.Raw.Get()) {
		var ret string
		return ret
	}
	return *o.Raw.Get()
}

// GetRawOk returns a tuple with the Raw field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *EmployerAddressesEmployerAddressItemResponse) GetRawOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Raw.Get(), o.Raw.IsSet()
}

// HasRaw returns a boolean if a field has been set.
func (o *EmployerAddressesEmployerAddressItemResponse) HasRaw() bool {
	if o != nil && o.Raw.IsSet() {
		return true
	}

	return false
}

// SetRaw gets a reference to the given NullableString and assigns it to the Raw field.
func (o *EmployerAddressesEmployerAddressItemResponse) SetRaw(v string) {
	o.Raw.Set(&v)
}
// SetRawNil sets the value for Raw to be an explicit nil
func (o *EmployerAddressesEmployerAddressItemResponse) SetRawNil() {
	o.Raw.Set(nil)
}

// UnsetRaw ensures that no value is present for Raw, not even an explicit nil
func (o *EmployerAddressesEmployerAddressItemResponse) UnsetRaw() {
	o.Raw.Unset()
}

// GetStreet returns the Street field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *EmployerAddressesEmployerAddressItemResponse) GetStreet() string {
	if o == nil || IsNil(o.Street.Get()) {
		var ret string
		return ret
	}
	return *o.Street.Get()
}

// GetStreetOk returns a tuple with the Street field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *EmployerAddressesEmployerAddressItemResponse) GetStreetOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Street.Get(), o.Street.IsSet()
}

// HasStreet returns a boolean if a field has been set.
func (o *EmployerAddressesEmployerAddressItemResponse) HasStreet() bool {
	if o != nil && o.Street.IsSet() {
		return true
	}

	return false
}

// SetStreet gets a reference to the given NullableString and assigns it to the Street field.
func (o *EmployerAddressesEmployerAddressItemResponse) SetStreet(v string) {
	o.Street.Set(&v)
}
// SetStreetNil sets the value for Street to be an explicit nil
func (o *EmployerAddressesEmployerAddressItemResponse) SetStreetNil() {
	o.Street.Set(nil)
}

// UnsetStreet ensures that no value is present for Street, not even an explicit nil
func (o *EmployerAddressesEmployerAddressItemResponse) UnsetStreet() {
	o.Street.Unset()
}

func (o EmployerAddressesEmployerAddressItemResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o EmployerAddressesEmployerAddressItemResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.Building.IsSet() {
		toSerialize["building"] = o.Building.Get()
	}
	if !IsNil(o.CanEdit) {
		toSerialize["can_edit"] = o.CanEdit
	}
	toSerialize["city"] = o.City.Get()
	if !IsNil(o.Deleted) {
		toSerialize["deleted"] = o.Deleted
	}
	if o.Description.IsSet() {
		toSerialize["description"] = o.Description.Get()
	}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	toSerialize["lat"] = o.Lat.Get()
	toSerialize["lng"] = o.Lng.Get()
	if o.Manager.IsSet() {
		toSerialize["manager"] = o.Manager.Get()
	}
	if !IsNil(o.MetroStations) {
		toSerialize["metro_stations"] = o.MetroStations
	}
	if o.Raw.IsSet() {
		toSerialize["raw"] = o.Raw.Get()
	}
	if o.Street.IsSet() {
		toSerialize["street"] = o.Street.Get()
	}
	return toSerialize, nil
}

func (o *EmployerAddressesEmployerAddressItemResponse) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"city",
		"lat",
		"lng",
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

	varEmployerAddressesEmployerAddressItemResponse := _EmployerAddressesEmployerAddressItemResponse{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varEmployerAddressesEmployerAddressItemResponse)

	if err != nil {
		return err
	}

	*o = EmployerAddressesEmployerAddressItemResponse(varEmployerAddressesEmployerAddressItemResponse)

	return err
}

type NullableEmployerAddressesEmployerAddressItemResponse struct {
	value *EmployerAddressesEmployerAddressItemResponse
	isSet bool
}

func (v NullableEmployerAddressesEmployerAddressItemResponse) Get() *EmployerAddressesEmployerAddressItemResponse {
	return v.value
}

func (v *NullableEmployerAddressesEmployerAddressItemResponse) Set(val *EmployerAddressesEmployerAddressItemResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableEmployerAddressesEmployerAddressItemResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableEmployerAddressesEmployerAddressItemResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEmployerAddressesEmployerAddressItemResponse(val *EmployerAddressesEmployerAddressItemResponse) *NullableEmployerAddressesEmployerAddressItemResponse {
	return &NullableEmployerAddressesEmployerAddressItemResponse{value: val, isSet: true}
}

func (v NullableEmployerAddressesEmployerAddressItemResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEmployerAddressesEmployerAddressItemResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


