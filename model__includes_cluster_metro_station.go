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

// checks if the IncludesClusterMetroStation type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &IncludesClusterMetroStation{}

// IncludesClusterMetroStation struct for IncludesClusterMetroStation
type IncludesClusterMetroStation struct {
	Area IncludesClusterMetroStationArea `json:"area"`
	// Цвет линии в HEX-формате `RRGGBB` (от `000000` до `FFFFFF`)
	HexColor string `json:"hex_color"`
	// Идентификатор станции метро
	Id string `json:"id"`
	// Широта
	Lat float32 `json:"lat"`
	// Долгота
	Lng float32 `json:"lng"`
	// Порядковый номер станции в линии метро
	Order float32 `json:"order"`
}

type _IncludesClusterMetroStation IncludesClusterMetroStation

// NewIncludesClusterMetroStation instantiates a new IncludesClusterMetroStation object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewIncludesClusterMetroStation(area IncludesClusterMetroStationArea, hexColor string, id string, lat float32, lng float32, order float32) *IncludesClusterMetroStation {
	this := IncludesClusterMetroStation{}
	this.Area = area
	this.HexColor = hexColor
	this.Id = id
	this.Lat = lat
	this.Lng = lng
	this.Order = order
	return &this
}

// NewIncludesClusterMetroStationWithDefaults instantiates a new IncludesClusterMetroStation object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewIncludesClusterMetroStationWithDefaults() *IncludesClusterMetroStation {
	this := IncludesClusterMetroStation{}
	return &this
}

// GetArea returns the Area field value
func (o *IncludesClusterMetroStation) GetArea() IncludesClusterMetroStationArea {
	if o == nil {
		var ret IncludesClusterMetroStationArea
		return ret
	}

	return o.Area
}

// GetAreaOk returns a tuple with the Area field value
// and a boolean to check if the value has been set.
func (o *IncludesClusterMetroStation) GetAreaOk() (*IncludesClusterMetroStationArea, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Area, true
}

// SetArea sets field value
func (o *IncludesClusterMetroStation) SetArea(v IncludesClusterMetroStationArea) {
	o.Area = v
}

// GetHexColor returns the HexColor field value
func (o *IncludesClusterMetroStation) GetHexColor() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.HexColor
}

// GetHexColorOk returns a tuple with the HexColor field value
// and a boolean to check if the value has been set.
func (o *IncludesClusterMetroStation) GetHexColorOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.HexColor, true
}

// SetHexColor sets field value
func (o *IncludesClusterMetroStation) SetHexColor(v string) {
	o.HexColor = v
}

// GetId returns the Id field value
func (o *IncludesClusterMetroStation) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *IncludesClusterMetroStation) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *IncludesClusterMetroStation) SetId(v string) {
	o.Id = v
}

// GetLat returns the Lat field value
func (o *IncludesClusterMetroStation) GetLat() float32 {
	if o == nil {
		var ret float32
		return ret
	}

	return o.Lat
}

// GetLatOk returns a tuple with the Lat field value
// and a boolean to check if the value has been set.
func (o *IncludesClusterMetroStation) GetLatOk() (*float32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Lat, true
}

// SetLat sets field value
func (o *IncludesClusterMetroStation) SetLat(v float32) {
	o.Lat = v
}

// GetLng returns the Lng field value
func (o *IncludesClusterMetroStation) GetLng() float32 {
	if o == nil {
		var ret float32
		return ret
	}

	return o.Lng
}

// GetLngOk returns a tuple with the Lng field value
// and a boolean to check if the value has been set.
func (o *IncludesClusterMetroStation) GetLngOk() (*float32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Lng, true
}

// SetLng sets field value
func (o *IncludesClusterMetroStation) SetLng(v float32) {
	o.Lng = v
}

// GetOrder returns the Order field value
func (o *IncludesClusterMetroStation) GetOrder() float32 {
	if o == nil {
		var ret float32
		return ret
	}

	return o.Order
}

// GetOrderOk returns a tuple with the Order field value
// and a boolean to check if the value has been set.
func (o *IncludesClusterMetroStation) GetOrderOk() (*float32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Order, true
}

// SetOrder sets field value
func (o *IncludesClusterMetroStation) SetOrder(v float32) {
	o.Order = v
}

func (o IncludesClusterMetroStation) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o IncludesClusterMetroStation) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["area"] = o.Area
	toSerialize["hex_color"] = o.HexColor
	toSerialize["id"] = o.Id
	toSerialize["lat"] = o.Lat
	toSerialize["lng"] = o.Lng
	toSerialize["order"] = o.Order
	return toSerialize, nil
}

func (o *IncludesClusterMetroStation) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"area",
		"hex_color",
		"id",
		"lat",
		"lng",
		"order",
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

	varIncludesClusterMetroStation := _IncludesClusterMetroStation{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varIncludesClusterMetroStation)

	if err != nil {
		return err
	}

	*o = IncludesClusterMetroStation(varIncludesClusterMetroStation)

	return err
}

type NullableIncludesClusterMetroStation struct {
	value *IncludesClusterMetroStation
	isSet bool
}

func (v NullableIncludesClusterMetroStation) Get() *IncludesClusterMetroStation {
	return v.value
}

func (v *NullableIncludesClusterMetroStation) Set(val *IncludesClusterMetroStation) {
	v.value = val
	v.isSet = true
}

func (v NullableIncludesClusterMetroStation) IsSet() bool {
	return v.isSet
}

func (v *NullableIncludesClusterMetroStation) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableIncludesClusterMetroStation(val *IncludesClusterMetroStation) *NullableIncludesClusterMetroStation {
	return &NullableIncludesClusterMetroStation{value: val, isSet: true}
}

func (v NullableIncludesClusterMetroStation) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableIncludesClusterMetroStation) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


