/*
HeadHunter API

По-русски | [Switch to English](https://api.hh.ru/openapi/en/redoc)  В OpenAPI ведется пока что только небольшая часть документации [Основная документация](https://github.com/hhru/api).  Для поиска по документации можно использовать Ctrl+F.  # Общая информация  * Всё API работает по протоколу HTTPS. * Авторизация осуществляется по протоколу OAuth2. * Все данные доступны только в формате JSON. * Базовый URL — `https://api.hh.ru/` * Возможны запросы к данным [любого сайта группы компаний HeadHunter](#section/Obshaya-informaciya/Vybor-sajta) * <a name=\"date-format\"></a> Даты форматируются в соответствии с [ISO 8601](http://en.wikipedia.org/wiki/ISO_8601): `YYYY-MM-DDThh:mm:ss±hhmm`.   <a name=\"request-requirements\"></a> ## Требования к запросам  В запросе необходимо передавать заголовок `User-Agent`, но если ваша реализация http клиента не позволяет, можно отправить `HH-User-Agent`. Если не отправлен ни один заголовок, то ответом будет `400 Bad Request`. Указание в заголовке названия приложения и контактной почты разработчика позволит нам оперативно с вами связаться в случае необходимости. Заголовки `User-Agent` и `HH-User-Agent` взаимозаменяемы, в случае, если вы отправите оба заголовка, обработан будет только `HH-User-Agent`.  ``` User-Agent: MyApp/1.0 (my-app-feedback@example.com) ```  Подробнее про [ошибки в заголовке User-Agent](https://github.com/hhru/api/blob/master/docs/errors.md#user-agent).   <a name=\"request-body\"></a> ## Формат тела запроса при отправке JSON  Данные, передающиеся в теле запроса, должны удовлетворять требованиям:  * Валидный JSON (допускается передача как минифицированного варианта, так и pretty print варианта с дополнительными пробелами и сбросами строк). * Рекомендуется использование кодировки UTF-8 без дополнительного экранирования (`{\"name\": \"Иванов Иван\"}`). * Также возможно использовать ascii кодировку с экранированием (`{\"name\": \"\\u0418\\u0432\\u0430\\u043d\\u043e\\u0432 \\u0418\\u0432\\u0430\\u043d\"}`). * К типам данных в определённым полях накладываются дополнительные условия, описанные в каждом конкретном методе. В JSON типами данных являются `string`, `number`, `boolean`, `null`, `object`, `array`.  ### Ответ Ответ свыше определенной длины будет сжиматься методом gzip.  ### Ошибки и коды ответов  API широко использует информирование при помощи кодов ответов. Приложение должно корректно их обрабатывать.  В случае неполадок и сбоев, возможны ответы с кодом `503` и `500`.  При каждой ошибке, помимо кода ответа, в теле ответа может быть выдана дополнительная информация, позволяющая разработчику понять причину соответствующего ответа.  [Более подробно про возможные ошибки](https://github.com/hhru/api/blob/master/docs/errors.md).   ## Недокументированные поля и параметры запросов  В ответах и параметрах API можно найти ключи, не описанные в документации. Обычно это означает, что они оставлены для совместимости со старыми версиями. Их использование не рекомендуется. Если ваше приложение использует такие ключи, перейдите на использование актуальных ключей, описанных в документации.   ## Пагинация  К любому запросу, подразумевающему выдачу списка объектов, можно в параметрах указать `page=N&per_page=M`. Нумерация идёт с нуля, по умолчанию выдаётся первая (нулевая) страница с 20 объектами на странице. Во всех ответах, где доступна пагинация, единообразный корневой объект:  ```json {   \"found\": 1,   \"per_page\": 1,   \"pages\": 1,   \"page\": 0,   \"items\": [{}] } ``` ## Выбор сайта  API HeadHunter позволяет получать данные со всех сайтов группы компании HeadHunter.  В частности:  * hh.ru * rabota.by * hh1.az * hh.uz * hh.kz * headhunter.ge * headhunter.kg  Запросы к данным на всех сайтах следует направлять на `https://api.hh.ru/`.  При необходимости учесть специфику сайта, можно добавить в запрос параметр `?host=`. По умолчанию используется `hh.ru`.  Например, для получения [локализаций](https://api.hh.ru/openapi/redoc#tag/Obshie-spravochniki/operation/get-locales), доступных на hh.kz необходимо сделать GET запрос на `https://api.hh.ru/locales?host=hh.kz`.  ## CORS (Cross-Origin Resource Sharing)  API поддерживает технологию CORS для запроса данных из браузера с произвольного домена. Этот метод более предпочтителен, чем использование JSONP. Он не ограничен методом GET. Для отладки CORS доступен [специальный метод](https://github.com/hhru/api/blob/master/docs/cors.md). Для использования JSONP передайте параметр `?callback=callback_name`.  * [CORS specification on w3.org](http://www.w3.org/TR/cors/) * [HTML5Rocks CORS Tutorial](http://www.html5rocks.com/en/tutorials/cors/) * [CORS on dev.opera.com](http://dev.opera.com/articles/view/dom-access-control-using-cross-origin-resource-sharing/) * [CORS on caniuse.com](http://caniuse.com/#feat=cors) * [CORS on en.wikipedia.org](http://en.wikipedia.org/wiki/Cross-origin_resource_sharing)   ## Внешние ссылки на статьи и стандарты  * [HTTP/1.1](http://tools.ietf.org/html/rfc2616) * [JSON](http://json.org/) * [URI Template](http://tools.ietf.org/html/rfc6570) * [OAuth 2.0](http://tools.ietf.org/html/rfc6749) * [REST](http://www.ics.uci.edu/~fielding/pubs/dissertation/rest_arch_style.htm) * [ISO 8601](http://en.wikipedia.org/wiki/ISO_8601)  # Авторизация  API поддерживает следующие уровни авторизации:   - [авторизация приложения](#tag/Avtorizaciya-prilozheniya)   - [авторизация пользователя](#section/Avtorizaciya/Avtorizaciya-polzovatelya)  * [Авторизация пользователя](#section/Avtorizaciya/Avtorizaciya-polzovatelya)   * [Правила формирования специального redirect_uri](#section/Avtorizaciya/Pravila-formirovaniya-specialnogo-redirect_uri)   * [Процесс авторизации](#section/Avtorizaciya/Process-avtorizacii)   * [Успешное получение временного `authorization_code`](#get-authorization_code)   * [Получение access и refresh токенов](#section/Avtorizaciya/Poluchenie-access-i-refresh-tokenov) * [Обновление пары access и refresh токенов](#section/Avtorizaciya/Obnovlenie-pary-access-i-refresh-tokenov) * [Инвалидация токена](#tag/Avtorizaciya-rabotodatelya/operation/invalidate-token) * [Запрос авторизации под другим пользователем](#section/Avtorizaciya/Zapros-avtorizacii-pod-drugim-polzovatelem) * [Авторизация под разными рабочими аккаунтами](#section/Avtorizaciya/Avtorizaciya-pod-raznymi-rabochimi-akkauntami) * [Авторизация приложения](#tag/Avtorizaciya-prilozheniya)       ## Авторизация пользователя Для выполнения запросов от имени пользователя необходимо пользоваться токеном пользователя.  В начале приложению необходимо направить пользователя (открыть страницу) по адресу:  ``` https://hh.ru/oauth/authorize? response_type=code& client_id={client_id}& state={state}& redirect_uri={redirect_uri} ```  Обязательные параметры:  * `response_type=code` — указание на способ получения авторизации, используя `authorization code` * `client_id` — идентификатор, полученный при создании приложения   Необязательные параметры:  * `state` — в случае указания, будет включен в ответный редирект. Это позволяет исключить возможность взлома путём подделки межсайтовых запросов. Подробнее об этом: [RFC 6749. Section 10.12](http://tools.ietf.org/html/rfc6749#section-10.12) * `redirect_uri` — uri для перенаправления пользователя после авторизации. Если не указать, используется из настроек приложения. При наличии происходит валидация значения. Вероятнее всего, потребуется сделать urlencode значения параметра.  ## Правила формирования специального redirect_uri  К примеру, если в настройках сохранен `http://example.com/oauth`, то разрешено указывать:  * `http://www.example.com/oauth` — поддомен; * `http://www.example.com/oauth/sub/path` — уточнение пути; * `http://example.com/oauth?lang=RU` — дополнительный параметр; * `http://www.example.com/oauth/sub/path?lang=RU` — всё вместе.  Запрещено:  * `https://example.com/oauth` — различные протоколы; * `http://wwwexample.com/oauth` — различные домены; * `http://wwwexample.com/` — другой путь; * `http://example.com/oauths` — другой путь; * `http://example.com:80/oauths` — указание изначально отсутствующего порта;  ## Процесс авторизации  Если пользователь не авторизован на сайте, ему будет показана форма авторизации на сайте. После прохождения авторизации на сайте, пользователю будет выведена форма с запросом разрешения доступа вашего приложения к его персональным данным.  Если пользователь не разрешает доступ приложению, пользователь будет перенаправлен на указанный `redirect_uri` с `?error=access_denied` и `state={state}`, если таковой был указан при первом запросе.  <a name=\"get-authorization_code\"></a> ### Успешное получение временного `authorization_code`  В случае разрешения прав, в редиректе будет указан временный `authorization_code`:  ```http HTTP/1.1 302 FOUND Location: {redirect_uri}?code={authorization_code} ```  Если пользователь авторизован на сайте и доступ данному приложению однажды ранее выдан, ответом будет сразу вышеописанный редирект с `authorization_code` (без показа формы логина и выдачи прав).  ## Получение access и refresh токенов  После получения `authorization_code` приложению необходимо сделать сервер-сервер запрос `POST https://api.hh.ru/token` для обмена полученного `authorization_code` на `access_token` (старый запрос `POST https://hh.ru/oauth/token` считается устаревшим).  В теле запроса необходимо передать [дополнительные параметры](#required_parameters).  Тело запроса необходимо передавать в стандартном `application/x-www-form-urlencoded` с указанием соответствующего заголовка `Content-Type`.  `authorization_code` имеет довольно короткий срок жизни, при его истечении необходимо запросить новый.  ## Обновление пары access и refresh токенов `access_token` также имеет срок жизни (ключ `expires_in`, в секундах), при его истечении приложение должно сделать запрос с `refresh_token` для получения нового.  Запрос необходимо делать в `application/x-www-form-urlencoded`.  ``` POST https://api.hh.ru/token ```  (старый запрос `POST https://hh.ru/oauth/token` считается устаревшим)  В теле запроса необходимо передать [дополнительные параметры](#required_parameters)  `refresh_token` можно использовать только один раз и только по истечению срока действия `access_token`.  После получения новой пары access и refresh токенов, их необходимо использовать в дальнейших запросах в api и запросах на продление токена.  ## Запрос авторизации под другим пользователем  Возможен следующий сценарий:  1. Приложение перенаправляет пользователя на сайт с запросом авторизации. 2. Пользователь на сайте уже авторизован и данному приложение доступ уже был разрешен. 3. Пользователю будет предложена возможность продолжить работу под текущим аккаунтом, либо зайти под другим аккаунтом.  Если есть необходимость, чтобы на шаге 3 сразу происходило перенаправление (redirect) с временным токеном, необходимо добавить к запросу `/oauth/authorize...` параметр `skip_choose_account=true`. В этом случае автоматически выдаётся доступ пользователю авторизованному на сайте.  Если есть необходимость всегда показывать форму авторизации, приложение может добавить к запросу `/oauth/authorize...` параметр `force_login=true`. В этом случае, пользователю будет показана форма авторизации с логином и паролем даже в случае, если пользователь уже авторизован.  Это может быть полезно приложениям, которые предоставляют сервис только для соискателей. Если пришел пользователь-работодатель, приложение может предложить пользователю повторно разрешить доступ на сайте, уже указав другую учетную запись.  Также, после авторизации приложение может показать пользователю сообщение:  ``` Вы вошли как %Имя_Фамилия%. Это не вы? ``` и предоставить ссылку с `force_login=true` для возможности захода под другим логином.  ## Авторизация под разными рабочими аккаунтами  Для получения списка рабочих аккаунтов менеджера и для работы под разными рабочими аккаунтами менеджера необходимо прочитать документацию по [рабочим аккаунтам менеджера](#tag/Menedzhery-rabotodatelya/operation/get-manager-accounts)  В случае компрометации токена необходимо [инвалидировать](#tag/Avtorizaciya-rabotodatelya/operation/invalidate-token) скомпрометированный токен и запросить токен заново!  <!-- ReDoc-Inject: <security-definitions> --> 

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

// checks if the MeCommonProfile type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &MeCommonProfile{}

// MeCommonProfile Содержит флаги, актуальные для любого типа авторизации: соискатель, работодатель, приложение 
type MeCommonProfile struct {
	// Тип авторизации
	AuthType NullableString `json:"auth_type"`
	// Является ли текущий пользователь администратором сайта
	IsAdmin bool `json:"is_admin"`
	// Является ли текущий пользователь соискателем
	IsApplicant bool `json:"is_applicant"`
	// Является ли авторизованный клиент приложением
	IsApplication bool `json:"is_application"`
	// Является ли текущий пользователь менеджером
	IsEmployer bool `json:"is_employer"`
	// Является ли текущий пользователь работодателем
	IsEmployerIntegration bool `json:"is_employer_integration"`
}

type _MeCommonProfile MeCommonProfile

// NewMeCommonProfile instantiates a new MeCommonProfile object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMeCommonProfile(authType NullableString, isAdmin bool, isApplicant bool, isApplication bool, isEmployer bool, isEmployerIntegration bool) *MeCommonProfile {
	this := MeCommonProfile{}
	this.AuthType = authType
	this.IsAdmin = isAdmin
	this.IsApplicant = isApplicant
	this.IsApplication = isApplication
	this.IsEmployer = isEmployer
	this.IsEmployerIntegration = isEmployerIntegration
	return &this
}

// NewMeCommonProfileWithDefaults instantiates a new MeCommonProfile object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMeCommonProfileWithDefaults() *MeCommonProfile {
	this := MeCommonProfile{}
	return &this
}

// GetAuthType returns the AuthType field value
// If the value is explicit nil, the zero value for string will be returned
func (o *MeCommonProfile) GetAuthType() string {
	if o == nil || o.AuthType.Get() == nil {
		var ret string
		return ret
	}

	return *o.AuthType.Get()
}

// GetAuthTypeOk returns a tuple with the AuthType field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *MeCommonProfile) GetAuthTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.AuthType.Get(), o.AuthType.IsSet()
}

// SetAuthType sets field value
func (o *MeCommonProfile) SetAuthType(v string) {
	o.AuthType.Set(&v)
}

// GetIsAdmin returns the IsAdmin field value
func (o *MeCommonProfile) GetIsAdmin() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.IsAdmin
}

// GetIsAdminOk returns a tuple with the IsAdmin field value
// and a boolean to check if the value has been set.
func (o *MeCommonProfile) GetIsAdminOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.IsAdmin, true
}

// SetIsAdmin sets field value
func (o *MeCommonProfile) SetIsAdmin(v bool) {
	o.IsAdmin = v
}

// GetIsApplicant returns the IsApplicant field value
func (o *MeCommonProfile) GetIsApplicant() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.IsApplicant
}

// GetIsApplicantOk returns a tuple with the IsApplicant field value
// and a boolean to check if the value has been set.
func (o *MeCommonProfile) GetIsApplicantOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.IsApplicant, true
}

// SetIsApplicant sets field value
func (o *MeCommonProfile) SetIsApplicant(v bool) {
	o.IsApplicant = v
}

// GetIsApplication returns the IsApplication field value
func (o *MeCommonProfile) GetIsApplication() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.IsApplication
}

// GetIsApplicationOk returns a tuple with the IsApplication field value
// and a boolean to check if the value has been set.
func (o *MeCommonProfile) GetIsApplicationOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.IsApplication, true
}

// SetIsApplication sets field value
func (o *MeCommonProfile) SetIsApplication(v bool) {
	o.IsApplication = v
}

// GetIsEmployer returns the IsEmployer field value
func (o *MeCommonProfile) GetIsEmployer() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.IsEmployer
}

// GetIsEmployerOk returns a tuple with the IsEmployer field value
// and a boolean to check if the value has been set.
func (o *MeCommonProfile) GetIsEmployerOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.IsEmployer, true
}

// SetIsEmployer sets field value
func (o *MeCommonProfile) SetIsEmployer(v bool) {
	o.IsEmployer = v
}

// GetIsEmployerIntegration returns the IsEmployerIntegration field value
func (o *MeCommonProfile) GetIsEmployerIntegration() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.IsEmployerIntegration
}

// GetIsEmployerIntegrationOk returns a tuple with the IsEmployerIntegration field value
// and a boolean to check if the value has been set.
func (o *MeCommonProfile) GetIsEmployerIntegrationOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.IsEmployerIntegration, true
}

// SetIsEmployerIntegration sets field value
func (o *MeCommonProfile) SetIsEmployerIntegration(v bool) {
	o.IsEmployerIntegration = v
}

func (o MeCommonProfile) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o MeCommonProfile) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["auth_type"] = o.AuthType.Get()
	toSerialize["is_admin"] = o.IsAdmin
	toSerialize["is_applicant"] = o.IsApplicant
	toSerialize["is_application"] = o.IsApplication
	toSerialize["is_employer"] = o.IsEmployer
	toSerialize["is_employer_integration"] = o.IsEmployerIntegration
	return toSerialize, nil
}

func (o *MeCommonProfile) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"auth_type",
		"is_admin",
		"is_applicant",
		"is_application",
		"is_employer",
		"is_employer_integration",
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

	varMeCommonProfile := _MeCommonProfile{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varMeCommonProfile)

	if err != nil {
		return err
	}

	*o = MeCommonProfile(varMeCommonProfile)

	return err
}

type NullableMeCommonProfile struct {
	value *MeCommonProfile
	isSet bool
}

func (v NullableMeCommonProfile) Get() *MeCommonProfile {
	return v.value
}

func (v *NullableMeCommonProfile) Set(val *MeCommonProfile) {
	v.value = val
	v.isSet = true
}

func (v NullableMeCommonProfile) IsSet() bool {
	return v.isSet
}

func (v *NullableMeCommonProfile) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMeCommonProfile(val *MeCommonProfile) *NullableMeCommonProfile {
	return &NullableMeCommonProfile{value: val, isSet: true}
}

func (v NullableMeCommonProfile) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMeCommonProfile) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


