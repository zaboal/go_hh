/*
HeadHunter API

По-русски | [Switch to English](https://api.hh.ru/openapi/en/redoc)  В OpenAPI ведется пока что только небольшая часть документации [Основная документация](https://github.com/hhru/api).  Для поиска по документации можно использовать Ctrl+F.  # Общая информация  * Всё API работает по протоколу HTTPS. * Авторизация осуществляется по протоколу OAuth2. * Все данные доступны только в формате JSON. * Базовый URL — `https://api.hh.ru/` * Возможны запросы к данным [любого сайта группы компаний HeadHunter](#section/Obshaya-informaciya/Vybor-sajta) * <a name=\"date-format\"></a> Даты форматируются в соответствии с [ISO 8601](http://en.wikipedia.org/wiki/ISO_8601): `YYYY-MM-DDThh:mm:ss±hhmm`.   <a name=\"request-requirements\"></a> ## Требования к запросам  В запросе необходимо передавать заголовок `User-Agent`, но если ваша реализация http клиента не позволяет, можно отправить `HH-User-Agent`. Если не отправлен ни один заголовок, то ответом будет `400 Bad Request`. Указание в заголовке названия приложения и контактной почты разработчика позволит нам оперативно с вами связаться в случае необходимости. Заголовки `User-Agent` и `HH-User-Agent` взаимозаменяемы, в случае, если вы отправите оба заголовка, обработан будет только `HH-User-Agent`.  ``` User-Agent: MyApp/1.0 (my-app-feedback@example.com) ```  Подробнее про [ошибки в заголовке User-Agent](https://github.com/hhru/api/blob/master/docs/errors.md#user-agent).   <a name=\"request-body\"></a> ## Формат тела запроса при отправке JSON  Данные, передающиеся в теле запроса, должны удовлетворять требованиям:  * Валидный JSON (допускается передача как минифицированного варианта, так и pretty print варианта с дополнительными пробелами и сбросами строк). * Рекомендуется использование кодировки UTF-8 без дополнительного экранирования (`{\"name\": \"Иванов Иван\"}`). * Также возможно использовать ascii кодировку с экранированием (`{\"name\": \"\\u0418\\u0432\\u0430\\u043d\\u043e\\u0432 \\u0418\\u0432\\u0430\\u043d\"}`). * К типам данных в определённым полях накладываются дополнительные условия, описанные в каждом конкретном методе. В JSON типами данных являются `string`, `number`, `boolean`, `null`, `object`, `array`.  ### Ответ Ответ свыше определенной длины будет сжиматься методом gzip.  ### Ошибки и коды ответов  API широко использует информирование при помощи кодов ответов. Приложение должно корректно их обрабатывать.  В случае неполадок и сбоев, возможны ответы с кодом `503` и `500`.  При каждой ошибке, помимо кода ответа, в теле ответа может быть выдана дополнительная информация, позволяющая разработчику понять причину соответствующего ответа.  [Более подробно про возможные ошибки](https://github.com/hhru/api/blob/master/docs/errors.md).   ## Недокументированные поля и параметры запросов  В ответах и параметрах API можно найти ключи, не описанные в документации. Обычно это означает, что они оставлены для совместимости со старыми версиями. Их использование не рекомендуется. Если ваше приложение использует такие ключи, перейдите на использование актуальных ключей, описанных в документации.   ## Пагинация  К любому запросу, подразумевающему выдачу списка объектов, можно в параметрах указать `page=N&per_page=M`. Нумерация идёт с нуля, по умолчанию выдаётся первая (нулевая) страница с 20 объектами на странице. Во всех ответах, где доступна пагинация, единообразный корневой объект:  ```json {   \"found\": 1,   \"per_page\": 1,   \"pages\": 1,   \"page\": 0,   \"items\": [{}] } ``` ## Выбор сайта  API HeadHunter позволяет получать данные со всех сайтов группы компании HeadHunter.  В частности:  * hh.ru * rabota.by * hh1.az * hh.uz * hh.kz * headhunter.ge * headhunter.kg  Запросы к данным на всех сайтах следует направлять на `https://api.hh.ru/`.  При необходимости учесть специфику сайта, можно добавить в запрос параметр `?host=`. По умолчанию используется `hh.ru`.  Например, для получения [локализаций](https://api.hh.ru/openapi/redoc#tag/Obshie-spravochniki/operation/get-locales), доступных на hh.kz необходимо сделать GET запрос на `https://api.hh.ru/locales?host=hh.kz`.  ## CORS (Cross-Origin Resource Sharing)  API поддерживает технологию CORS для запроса данных из браузера с произвольного домена. Этот метод более предпочтителен, чем использование JSONP. Он не ограничен методом GET. Для отладки CORS доступен [специальный метод](https://github.com/hhru/api/blob/master/docs/cors.md). Для использования JSONP передайте параметр `?callback=callback_name`.  * [CORS specification on w3.org](http://www.w3.org/TR/cors/) * [HTML5Rocks CORS Tutorial](http://www.html5rocks.com/en/tutorials/cors/) * [CORS on dev.opera.com](http://dev.opera.com/articles/view/dom-access-control-using-cross-origin-resource-sharing/) * [CORS on caniuse.com](http://caniuse.com/#feat=cors) * [CORS on en.wikipedia.org](http://en.wikipedia.org/wiki/Cross-origin_resource_sharing)   ## Внешние ссылки на статьи и стандарты  * [HTTP/1.1](http://tools.ietf.org/html/rfc2616) * [JSON](http://json.org/) * [URI Template](http://tools.ietf.org/html/rfc6570) * [OAuth 2.0](http://tools.ietf.org/html/rfc6749) * [REST](http://www.ics.uci.edu/~fielding/pubs/dissertation/rest_arch_style.htm) * [ISO 8601](http://en.wikipedia.org/wiki/ISO_8601)  # Авторизация  API поддерживает следующие уровни авторизации:   - [авторизация приложения](#section/Avtorizaciya/Avtorizaciya-prilozheniya)   - [авторизация пользователя](#section/Avtorizaciya/Avtorizaciya-polzovatelya)  * [Авторизация пользователя](#section/Avtorizaciya/Avtorizaciya-polzovatelya)   * [Правила формирования специального redirect_uri](#section/Avtorizaciya/Pravila-formirovaniya-specialnogo-redirect_uri)   * [Процесс авторизации](#section/Avtorizaciya/Process-avtorizacii)   * [Успешное получение временного `authorization_code`](#get-authorization_code)   * [Получение access и refresh токенов](#section/Avtorizaciya/Poluchenie-access-i-refresh-tokenov) * [Обновление пары access и refresh токенов](#section/Avtorizaciya/Obnovlenie-pary-access-i-refresh-tokenov) * [Инвалидация токена](#tag/Avtorizaciya-rabotodatelya/operation/invalidate-token) * [Запрос авторизации под другим пользователем](#section/Avtorizaciya/Zapros-avtorizacii-pod-drugim-polzovatelem) * [Авторизация под разными рабочими аккаунтами](#section/Avtorizaciya/Avtorizaciya-pod-raznymi-rabochimi-akkauntami) * [Авторизация приложения](#section/Avtorizaciya/Avtorizaciya-prilozheniya)       ## Авторизация пользователя Для выполнения запросов от имени пользователя необходимо пользоваться токеном пользователя.  В начале приложению необходимо направить пользователя (открыть страницу) по адресу:  ``` https://hh.ru/oauth/authorize? response_type=code& client_id={client_id}& state={state}& redirect_uri={redirect_uri} ```  Обязательные параметры:  * `response_type=code` — указание на способ получения авторизации, используя `authorization code` * `client_id` — идентификатор, полученный при создании приложения   Необязательные параметры:  * `state` — в случае указания, будет включен в ответный редирект. Это позволяет исключить возможность взлома путём подделки межсайтовых запросов. Подробнее об этом: [RFC 6749. Section 10.12](http://tools.ietf.org/html/rfc6749#section-10.12) * `redirect_uri` — uri для перенаправления пользователя после авторизации. Если не указать, используется из настроек приложения. При наличии происходит валидация значения. Вероятнее всего, потребуется сделать urlencode значения параметра.  ## Правила формирования специального redirect_uri  К примеру, если в настройках сохранен `http://example.com/oauth`, то разрешено указывать:  * `http://www.example.com/oauth` — поддомен; * `http://www.example.com/oauth/sub/path` — уточнение пути; * `http://example.com/oauth?lang=RU` — дополнительный параметр; * `http://www.example.com/oauth/sub/path?lang=RU` — всё вместе.  Запрещено:  * `https://example.com/oauth` — различные протоколы; * `http://wwwexample.com/oauth` — различные домены; * `http://wwwexample.com/` — другой путь; * `http://example.com/oauths` — другой путь; * `http://example.com:80/oauths` — указание изначально отсутствующего порта;  ## Процесс авторизации  Если пользователь не авторизован на сайте, ему будет показана форма авторизации на сайте. После прохождения авторизации на сайте, пользователю будет выведена форма с запросом разрешения доступа вашего приложения к его персональным данным.  Если пользователь не разрешает доступ приложению, пользователь будет перенаправлен на указанный `redirect_uri` с `?error=access_denied` и `state={state}`, если таковой был указан при первом запросе.  <a name=\"get-authorization_code\"></a> ### Успешное получение временного `authorization_code`  В случае разрешения прав, в редиректе будет указан временный `authorization_code`:  ```http HTTP/1.1 302 FOUND Location: {redirect_uri}?code={authorization_code} ```  Если пользователь авторизован на сайте и доступ данному приложению однажды ранее выдан, ответом будет сразу вышеописанный редирект с `authorization_code` (без показа формы логина и выдачи прав).  ## Получение access и refresh токенов  После получения `authorization_code` приложению необходимо сделать сервер-сервер запрос `POST https://api.hh.ru/token` для обмена полученного `authorization_code` на `access_token` (старый запрос `POST https://hh.ru/oauth/token` считается устаревшим).  В теле запроса необходимо передать [дополнительные параметры](#required_parameters).  Тело запроса необходимо передавать в стандартном `application/x-www-form-urlencoded` с указанием соответствующего заголовка `Content-Type`.  `authorization_code` имеет довольно короткий срок жизни, при его истечении необходимо запросить новый.  ## Обновление пары access и refresh токенов `access_token` также имеет срок жизни (ключ `expires_in`, в секундах), при его истечении приложение должно сделать запрос с `refresh_token` для получения нового.  Запрос необходимо делать в `application/x-www-form-urlencoded`.  ``` POST https://api.hh.ru/token ```  (старый запрос `POST https://hh.ru/oauth/token` считается устаревшим)  В теле запроса необходимо передать [дополнительные параметры](#required_parameters)  `refresh_token` можно использовать только один раз и только по истечению срока действия `access_token`.  После получения новой пары access и refresh токенов, их необходимо использовать в дальнейших запросах в api и запросах на продление токена.  ## Запрос авторизации под другим пользователем  Возможен следующий сценарий:  1. Приложение перенаправляет пользователя на сайт с запросом авторизации. 2. Пользователь на сайте уже авторизован и данному приложение доступ уже был разрешен. 3. Пользователю будет предложена возможность продолжить работу под текущим аккаунтом, либо зайти под другим аккаунтом.  Если есть необходимость, чтобы на шаге 3 сразу происходило перенаправление (redirect) с временным токеном, необходимо добавить к запросу `/oauth/authorize...` параметр `skip_choose_account=true`. В этом случае автоматически выдаётся доступ пользователю авторизованному на сайте.  Если есть необходимость всегда показывать форму авторизации, приложение может добавить к запросу `/oauth/authorize...` параметр `force_login=true`. В этом случае, пользователю будет показана форма авторизации с логином и паролем даже в случае, если пользователь уже авторизован.  Это может быть полезно приложениям, которые предоставляют сервис только для соискателей. Если пришел пользователь-работодатель, приложение может предложить пользователю повторно разрешить доступ на сайте, уже указав другую учетную запись.  Также, после авторизации приложение может показать пользователю сообщение:  ``` Вы вошли как %Имя_Фамилия%. Это не вы? ``` и предоставить ссылку с `force_login=true` для возможности захода под другим логином.  ## Авторизация под разными рабочими аккаунтами  Для получения списка рабочих аккаунтов менеджера и для работы под разными рабочими аккаунтами менеджера необходимо прочитать документацию по [рабочим аккаунтам менеджера](#tag/Menedzhery-rabotodatelya/operation/get-manager-accounts)  ## Авторизация приложения  Токен приложения необходимо сгенерировать 1 раз. В случае, если токен был скомпрометирован, его нужно запросить еще раз. При этом ранее выданный токен отзывается. Владелец приложения может посмотреть актуальный `access_token` для приложения на сайте [https://dev.hh.ru/admin](https://dev.hh.ru/admin). В случае, если вы еще ни разу [не получали токен приложения](#section/Avtorizaciya/Avtorizaciya-prilozheniya), токен отображаться не будет.  <a name=\"get-client-token\"></a> ### Получение токена приложения Для получения `access_token` необходимо сделать запрос:  ``` POST https://api.hh.ru/token ```  (старый запрос `POST https://hh.ru/oauth/token` считается устаревшим)  В теле запроса необходимо передать [дополнительные параметры](#required_parameters). Тело запроса необходимо передавать в стандартном `application/x-www-form-urlencoded` с указанием соответствующего заголовка `Content-Type`.  Данный `access_token` имеет **неограниченный** срок жизни. При повторном запросе ранее выданный токен отзывается и выдается новый. Запрашивать `access_token` можно не чаще, чем один раз в 5 минут.  В случае компрометации токена необходимо [инвалидировать](#tag/Avtorizaciya-rabotodatelya/operation/invalidate-token) скомпроментированный токен и запросить токен заново!  <!-- ReDoc-Inject: <security-definitions> --> 

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

// checks if the VacanciesVisitorsVisitorItemsItemsInner type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &VacanciesVisitorsVisitorItemsItemsInner{}

// VacanciesVisitorsVisitorItemsItemsInner struct for VacanciesVisitorsVisitorItemsItemsInner
type VacanciesVisitorsVisitorItemsItemsInner struct {
	// URL резюме на сайте
	AlternateUrl string `json:"alternate_url"`
	// Идентификатор резюме
	Id string `json:"id"`
	// Желаемая должность
	Title NullableString `json:"title"`
	// Возраст
	Age NullableFloat32 `json:"age,omitempty"`
	Area NullableIncludesIdNameUrl `json:"area,omitempty"`
	// Доступен ли просмотр контактной информации в резюме текущему работодателю
	CanViewFullInfo NullableBool `json:"can_view_full_info,omitempty"`
	// Список сертификатов соискателя
	Certificate []ResumeObjectsCertificate `json:"certificate"`
	// Дата и время создания резюме
	CreatedAt string `json:"created_at"`
	Download ResumeResumeProfileAllOfDownload `json:"download"`
	Education ResumeResumeProfileAllOfEducation `json:"education"`
	// Опыт работы
	Experience []ResumeObjectsExperience `json:"experience"`
	// Имя
	FirstName NullableString `json:"first_name,omitempty"`
	Gender NullableIncludesIdName `json:"gender,omitempty"`
	// Справочник [Список скрытых полей](https://github.com/hhru/api/blob/master/docs/employer_resumes.md#hidden-fields). Возможные значения элементов приведены в поле `resume_hidden_fields` [справочника полей](#tag/Obshie-spravochniki/operation/get-dictionaries)
	HiddenFields []IncludesIdName `json:"hidden_fields"`
	// Фамилия
	LastName NullableString `json:"last_name,omitempty"`
	// Выделено ли резюме в поиске
	Marked *bool `json:"marked,omitempty"`
	// Отчество
	MiddleName NullableString `json:"middle_name,omitempty"`
	Platform *ResumeResumeProfileAllOfPlatform `json:"platform,omitempty"`
	Salary NullableResumeObjectsSalaryProperties `json:"salary,omitempty"`
	TotalExperience NullableResumeObjectsTotalExperience `json:"total_experience,omitempty"`
	// Дата и время обновления резюме
	UpdatedAt string `json:"updated_at"`
	Actions ResumeResumeShortAllOfActions `json:"actions"`
	// Добавлено ли резюме в избранные
	Favorited bool `json:"favorited"`
	NegotiationsHistory ResumeResumeShortAllOfNegotiationsHistory `json:"negotiations_history"`
	Owner ResumeResumeShortAllOfOwner `json:"owner"`
	Photo NullableResumeResumeShortAllOfPhoto `json:"photo,omitempty"`
	// Было ли резюме уже просмотрено работодателем
	Viewed bool `json:"viewed"`
	// Ссылка на получение элементов
	Url string `json:"url"`
}

type _VacanciesVisitorsVisitorItemsItemsInner VacanciesVisitorsVisitorItemsItemsInner

// NewVacanciesVisitorsVisitorItemsItemsInner instantiates a new VacanciesVisitorsVisitorItemsItemsInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewVacanciesVisitorsVisitorItemsItemsInner(alternateUrl string, id string, title NullableString, certificate []ResumeObjectsCertificate, createdAt string, download ResumeResumeProfileAllOfDownload, education ResumeResumeProfileAllOfEducation, experience []ResumeObjectsExperience, hiddenFields []IncludesIdName, updatedAt string, actions ResumeResumeShortAllOfActions, favorited bool, negotiationsHistory ResumeResumeShortAllOfNegotiationsHistory, owner ResumeResumeShortAllOfOwner, viewed bool, url string) *VacanciesVisitorsVisitorItemsItemsInner {
	this := VacanciesVisitorsVisitorItemsItemsInner{}
	this.AlternateUrl = alternateUrl
	this.Id = id
	this.Title = title
	this.Certificate = certificate
	this.CreatedAt = createdAt
	this.Download = download
	this.Education = education
	this.Experience = experience
	this.HiddenFields = hiddenFields
	this.UpdatedAt = updatedAt
	this.Actions = actions
	this.Favorited = favorited
	this.NegotiationsHistory = negotiationsHistory
	this.Owner = owner
	this.Viewed = viewed
	this.Url = url
	return &this
}

// NewVacanciesVisitorsVisitorItemsItemsInnerWithDefaults instantiates a new VacanciesVisitorsVisitorItemsItemsInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewVacanciesVisitorsVisitorItemsItemsInnerWithDefaults() *VacanciesVisitorsVisitorItemsItemsInner {
	this := VacanciesVisitorsVisitorItemsItemsInner{}
	return &this
}

// GetAlternateUrl returns the AlternateUrl field value
func (o *VacanciesVisitorsVisitorItemsItemsInner) GetAlternateUrl() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.AlternateUrl
}

// GetAlternateUrlOk returns a tuple with the AlternateUrl field value
// and a boolean to check if the value has been set.
func (o *VacanciesVisitorsVisitorItemsItemsInner) GetAlternateUrlOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AlternateUrl, true
}

// SetAlternateUrl sets field value
func (o *VacanciesVisitorsVisitorItemsItemsInner) SetAlternateUrl(v string) {
	o.AlternateUrl = v
}

// GetId returns the Id field value
func (o *VacanciesVisitorsVisitorItemsItemsInner) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *VacanciesVisitorsVisitorItemsItemsInner) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *VacanciesVisitorsVisitorItemsItemsInner) SetId(v string) {
	o.Id = v
}

// GetTitle returns the Title field value
// If the value is explicit nil, the zero value for string will be returned
func (o *VacanciesVisitorsVisitorItemsItemsInner) GetTitle() string {
	if o == nil || o.Title.Get() == nil {
		var ret string
		return ret
	}

	return *o.Title.Get()
}

// GetTitleOk returns a tuple with the Title field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *VacanciesVisitorsVisitorItemsItemsInner) GetTitleOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Title.Get(), o.Title.IsSet()
}

// SetTitle sets field value
func (o *VacanciesVisitorsVisitorItemsItemsInner) SetTitle(v string) {
	o.Title.Set(&v)
}

// GetAge returns the Age field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *VacanciesVisitorsVisitorItemsItemsInner) GetAge() float32 {
	if o == nil || IsNil(o.Age.Get()) {
		var ret float32
		return ret
	}
	return *o.Age.Get()
}

// GetAgeOk returns a tuple with the Age field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *VacanciesVisitorsVisitorItemsItemsInner) GetAgeOk() (*float32, bool) {
	if o == nil {
		return nil, false
	}
	return o.Age.Get(), o.Age.IsSet()
}

// HasAge returns a boolean if a field has been set.
func (o *VacanciesVisitorsVisitorItemsItemsInner) HasAge() bool {
	if o != nil && o.Age.IsSet() {
		return true
	}

	return false
}

// SetAge gets a reference to the given NullableFloat32 and assigns it to the Age field.
func (o *VacanciesVisitorsVisitorItemsItemsInner) SetAge(v float32) {
	o.Age.Set(&v)
}
// SetAgeNil sets the value for Age to be an explicit nil
func (o *VacanciesVisitorsVisitorItemsItemsInner) SetAgeNil() {
	o.Age.Set(nil)
}

// UnsetAge ensures that no value is present for Age, not even an explicit nil
func (o *VacanciesVisitorsVisitorItemsItemsInner) UnsetAge() {
	o.Age.Unset()
}

// GetArea returns the Area field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *VacanciesVisitorsVisitorItemsItemsInner) GetArea() IncludesIdNameUrl {
	if o == nil || IsNil(o.Area.Get()) {
		var ret IncludesIdNameUrl
		return ret
	}
	return *o.Area.Get()
}

// GetAreaOk returns a tuple with the Area field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *VacanciesVisitorsVisitorItemsItemsInner) GetAreaOk() (*IncludesIdNameUrl, bool) {
	if o == nil {
		return nil, false
	}
	return o.Area.Get(), o.Area.IsSet()
}

// HasArea returns a boolean if a field has been set.
func (o *VacanciesVisitorsVisitorItemsItemsInner) HasArea() bool {
	if o != nil && o.Area.IsSet() {
		return true
	}

	return false
}

// SetArea gets a reference to the given NullableIncludesIdNameUrl and assigns it to the Area field.
func (o *VacanciesVisitorsVisitorItemsItemsInner) SetArea(v IncludesIdNameUrl) {
	o.Area.Set(&v)
}
// SetAreaNil sets the value for Area to be an explicit nil
func (o *VacanciesVisitorsVisitorItemsItemsInner) SetAreaNil() {
	o.Area.Set(nil)
}

// UnsetArea ensures that no value is present for Area, not even an explicit nil
func (o *VacanciesVisitorsVisitorItemsItemsInner) UnsetArea() {
	o.Area.Unset()
}

// GetCanViewFullInfo returns the CanViewFullInfo field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *VacanciesVisitorsVisitorItemsItemsInner) GetCanViewFullInfo() bool {
	if o == nil || IsNil(o.CanViewFullInfo.Get()) {
		var ret bool
		return ret
	}
	return *o.CanViewFullInfo.Get()
}

// GetCanViewFullInfoOk returns a tuple with the CanViewFullInfo field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *VacanciesVisitorsVisitorItemsItemsInner) GetCanViewFullInfoOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return o.CanViewFullInfo.Get(), o.CanViewFullInfo.IsSet()
}

// HasCanViewFullInfo returns a boolean if a field has been set.
func (o *VacanciesVisitorsVisitorItemsItemsInner) HasCanViewFullInfo() bool {
	if o != nil && o.CanViewFullInfo.IsSet() {
		return true
	}

	return false
}

// SetCanViewFullInfo gets a reference to the given NullableBool and assigns it to the CanViewFullInfo field.
func (o *VacanciesVisitorsVisitorItemsItemsInner) SetCanViewFullInfo(v bool) {
	o.CanViewFullInfo.Set(&v)
}
// SetCanViewFullInfoNil sets the value for CanViewFullInfo to be an explicit nil
func (o *VacanciesVisitorsVisitorItemsItemsInner) SetCanViewFullInfoNil() {
	o.CanViewFullInfo.Set(nil)
}

// UnsetCanViewFullInfo ensures that no value is present for CanViewFullInfo, not even an explicit nil
func (o *VacanciesVisitorsVisitorItemsItemsInner) UnsetCanViewFullInfo() {
	o.CanViewFullInfo.Unset()
}

// GetCertificate returns the Certificate field value
func (o *VacanciesVisitorsVisitorItemsItemsInner) GetCertificate() []ResumeObjectsCertificate {
	if o == nil {
		var ret []ResumeObjectsCertificate
		return ret
	}

	return o.Certificate
}

// GetCertificateOk returns a tuple with the Certificate field value
// and a boolean to check if the value has been set.
func (o *VacanciesVisitorsVisitorItemsItemsInner) GetCertificateOk() ([]ResumeObjectsCertificate, bool) {
	if o == nil {
		return nil, false
	}
	return o.Certificate, true
}

// SetCertificate sets field value
func (o *VacanciesVisitorsVisitorItemsItemsInner) SetCertificate(v []ResumeObjectsCertificate) {
	o.Certificate = v
}

// GetCreatedAt returns the CreatedAt field value
func (o *VacanciesVisitorsVisitorItemsItemsInner) GetCreatedAt() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value
// and a boolean to check if the value has been set.
func (o *VacanciesVisitorsVisitorItemsItemsInner) GetCreatedAtOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CreatedAt, true
}

// SetCreatedAt sets field value
func (o *VacanciesVisitorsVisitorItemsItemsInner) SetCreatedAt(v string) {
	o.CreatedAt = v
}

// GetDownload returns the Download field value
func (o *VacanciesVisitorsVisitorItemsItemsInner) GetDownload() ResumeResumeProfileAllOfDownload {
	if o == nil {
		var ret ResumeResumeProfileAllOfDownload
		return ret
	}

	return o.Download
}

// GetDownloadOk returns a tuple with the Download field value
// and a boolean to check if the value has been set.
func (o *VacanciesVisitorsVisitorItemsItemsInner) GetDownloadOk() (*ResumeResumeProfileAllOfDownload, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Download, true
}

// SetDownload sets field value
func (o *VacanciesVisitorsVisitorItemsItemsInner) SetDownload(v ResumeResumeProfileAllOfDownload) {
	o.Download = v
}

// GetEducation returns the Education field value
func (o *VacanciesVisitorsVisitorItemsItemsInner) GetEducation() ResumeResumeProfileAllOfEducation {
	if o == nil {
		var ret ResumeResumeProfileAllOfEducation
		return ret
	}

	return o.Education
}

// GetEducationOk returns a tuple with the Education field value
// and a boolean to check if the value has been set.
func (o *VacanciesVisitorsVisitorItemsItemsInner) GetEducationOk() (*ResumeResumeProfileAllOfEducation, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Education, true
}

// SetEducation sets field value
func (o *VacanciesVisitorsVisitorItemsItemsInner) SetEducation(v ResumeResumeProfileAllOfEducation) {
	o.Education = v
}

// GetExperience returns the Experience field value
func (o *VacanciesVisitorsVisitorItemsItemsInner) GetExperience() []ResumeObjectsExperience {
	if o == nil {
		var ret []ResumeObjectsExperience
		return ret
	}

	return o.Experience
}

// GetExperienceOk returns a tuple with the Experience field value
// and a boolean to check if the value has been set.
func (o *VacanciesVisitorsVisitorItemsItemsInner) GetExperienceOk() ([]ResumeObjectsExperience, bool) {
	if o == nil {
		return nil, false
	}
	return o.Experience, true
}

// SetExperience sets field value
func (o *VacanciesVisitorsVisitorItemsItemsInner) SetExperience(v []ResumeObjectsExperience) {
	o.Experience = v
}

// GetFirstName returns the FirstName field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *VacanciesVisitorsVisitorItemsItemsInner) GetFirstName() string {
	if o == nil || IsNil(o.FirstName.Get()) {
		var ret string
		return ret
	}
	return *o.FirstName.Get()
}

// GetFirstNameOk returns a tuple with the FirstName field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *VacanciesVisitorsVisitorItemsItemsInner) GetFirstNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.FirstName.Get(), o.FirstName.IsSet()
}

// HasFirstName returns a boolean if a field has been set.
func (o *VacanciesVisitorsVisitorItemsItemsInner) HasFirstName() bool {
	if o != nil && o.FirstName.IsSet() {
		return true
	}

	return false
}

// SetFirstName gets a reference to the given NullableString and assigns it to the FirstName field.
func (o *VacanciesVisitorsVisitorItemsItemsInner) SetFirstName(v string) {
	o.FirstName.Set(&v)
}
// SetFirstNameNil sets the value for FirstName to be an explicit nil
func (o *VacanciesVisitorsVisitorItemsItemsInner) SetFirstNameNil() {
	o.FirstName.Set(nil)
}

// UnsetFirstName ensures that no value is present for FirstName, not even an explicit nil
func (o *VacanciesVisitorsVisitorItemsItemsInner) UnsetFirstName() {
	o.FirstName.Unset()
}

// GetGender returns the Gender field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *VacanciesVisitorsVisitorItemsItemsInner) GetGender() IncludesIdName {
	if o == nil || IsNil(o.Gender.Get()) {
		var ret IncludesIdName
		return ret
	}
	return *o.Gender.Get()
}

// GetGenderOk returns a tuple with the Gender field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *VacanciesVisitorsVisitorItemsItemsInner) GetGenderOk() (*IncludesIdName, bool) {
	if o == nil {
		return nil, false
	}
	return o.Gender.Get(), o.Gender.IsSet()
}

// HasGender returns a boolean if a field has been set.
func (o *VacanciesVisitorsVisitorItemsItemsInner) HasGender() bool {
	if o != nil && o.Gender.IsSet() {
		return true
	}

	return false
}

// SetGender gets a reference to the given NullableIncludesIdName and assigns it to the Gender field.
func (o *VacanciesVisitorsVisitorItemsItemsInner) SetGender(v IncludesIdName) {
	o.Gender.Set(&v)
}
// SetGenderNil sets the value for Gender to be an explicit nil
func (o *VacanciesVisitorsVisitorItemsItemsInner) SetGenderNil() {
	o.Gender.Set(nil)
}

// UnsetGender ensures that no value is present for Gender, not even an explicit nil
func (o *VacanciesVisitorsVisitorItemsItemsInner) UnsetGender() {
	o.Gender.Unset()
}

// GetHiddenFields returns the HiddenFields field value
func (o *VacanciesVisitorsVisitorItemsItemsInner) GetHiddenFields() []IncludesIdName {
	if o == nil {
		var ret []IncludesIdName
		return ret
	}

	return o.HiddenFields
}

// GetHiddenFieldsOk returns a tuple with the HiddenFields field value
// and a boolean to check if the value has been set.
func (o *VacanciesVisitorsVisitorItemsItemsInner) GetHiddenFieldsOk() ([]IncludesIdName, bool) {
	if o == nil {
		return nil, false
	}
	return o.HiddenFields, true
}

// SetHiddenFields sets field value
func (o *VacanciesVisitorsVisitorItemsItemsInner) SetHiddenFields(v []IncludesIdName) {
	o.HiddenFields = v
}

// GetLastName returns the LastName field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *VacanciesVisitorsVisitorItemsItemsInner) GetLastName() string {
	if o == nil || IsNil(o.LastName.Get()) {
		var ret string
		return ret
	}
	return *o.LastName.Get()
}

// GetLastNameOk returns a tuple with the LastName field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *VacanciesVisitorsVisitorItemsItemsInner) GetLastNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.LastName.Get(), o.LastName.IsSet()
}

// HasLastName returns a boolean if a field has been set.
func (o *VacanciesVisitorsVisitorItemsItemsInner) HasLastName() bool {
	if o != nil && o.LastName.IsSet() {
		return true
	}

	return false
}

// SetLastName gets a reference to the given NullableString and assigns it to the LastName field.
func (o *VacanciesVisitorsVisitorItemsItemsInner) SetLastName(v string) {
	o.LastName.Set(&v)
}
// SetLastNameNil sets the value for LastName to be an explicit nil
func (o *VacanciesVisitorsVisitorItemsItemsInner) SetLastNameNil() {
	o.LastName.Set(nil)
}

// UnsetLastName ensures that no value is present for LastName, not even an explicit nil
func (o *VacanciesVisitorsVisitorItemsItemsInner) UnsetLastName() {
	o.LastName.Unset()
}

// GetMarked returns the Marked field value if set, zero value otherwise.
func (o *VacanciesVisitorsVisitorItemsItemsInner) GetMarked() bool {
	if o == nil || IsNil(o.Marked) {
		var ret bool
		return ret
	}
	return *o.Marked
}

// GetMarkedOk returns a tuple with the Marked field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VacanciesVisitorsVisitorItemsItemsInner) GetMarkedOk() (*bool, bool) {
	if o == nil || IsNil(o.Marked) {
		return nil, false
	}
	return o.Marked, true
}

// HasMarked returns a boolean if a field has been set.
func (o *VacanciesVisitorsVisitorItemsItemsInner) HasMarked() bool {
	if o != nil && !IsNil(o.Marked) {
		return true
	}

	return false
}

// SetMarked gets a reference to the given bool and assigns it to the Marked field.
func (o *VacanciesVisitorsVisitorItemsItemsInner) SetMarked(v bool) {
	o.Marked = &v
}

// GetMiddleName returns the MiddleName field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *VacanciesVisitorsVisitorItemsItemsInner) GetMiddleName() string {
	if o == nil || IsNil(o.MiddleName.Get()) {
		var ret string
		return ret
	}
	return *o.MiddleName.Get()
}

// GetMiddleNameOk returns a tuple with the MiddleName field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *VacanciesVisitorsVisitorItemsItemsInner) GetMiddleNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.MiddleName.Get(), o.MiddleName.IsSet()
}

// HasMiddleName returns a boolean if a field has been set.
func (o *VacanciesVisitorsVisitorItemsItemsInner) HasMiddleName() bool {
	if o != nil && o.MiddleName.IsSet() {
		return true
	}

	return false
}

// SetMiddleName gets a reference to the given NullableString and assigns it to the MiddleName field.
func (o *VacanciesVisitorsVisitorItemsItemsInner) SetMiddleName(v string) {
	o.MiddleName.Set(&v)
}
// SetMiddleNameNil sets the value for MiddleName to be an explicit nil
func (o *VacanciesVisitorsVisitorItemsItemsInner) SetMiddleNameNil() {
	o.MiddleName.Set(nil)
}

// UnsetMiddleName ensures that no value is present for MiddleName, not even an explicit nil
func (o *VacanciesVisitorsVisitorItemsItemsInner) UnsetMiddleName() {
	o.MiddleName.Unset()
}

// GetPlatform returns the Platform field value if set, zero value otherwise.
func (o *VacanciesVisitorsVisitorItemsItemsInner) GetPlatform() ResumeResumeProfileAllOfPlatform {
	if o == nil || IsNil(o.Platform) {
		var ret ResumeResumeProfileAllOfPlatform
		return ret
	}
	return *o.Platform
}

// GetPlatformOk returns a tuple with the Platform field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VacanciesVisitorsVisitorItemsItemsInner) GetPlatformOk() (*ResumeResumeProfileAllOfPlatform, bool) {
	if o == nil || IsNil(o.Platform) {
		return nil, false
	}
	return o.Platform, true
}

// HasPlatform returns a boolean if a field has been set.
func (o *VacanciesVisitorsVisitorItemsItemsInner) HasPlatform() bool {
	if o != nil && !IsNil(o.Platform) {
		return true
	}

	return false
}

// SetPlatform gets a reference to the given ResumeResumeProfileAllOfPlatform and assigns it to the Platform field.
func (o *VacanciesVisitorsVisitorItemsItemsInner) SetPlatform(v ResumeResumeProfileAllOfPlatform) {
	o.Platform = &v
}

// GetSalary returns the Salary field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *VacanciesVisitorsVisitorItemsItemsInner) GetSalary() ResumeObjectsSalaryProperties {
	if o == nil || IsNil(o.Salary.Get()) {
		var ret ResumeObjectsSalaryProperties
		return ret
	}
	return *o.Salary.Get()
}

// GetSalaryOk returns a tuple with the Salary field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *VacanciesVisitorsVisitorItemsItemsInner) GetSalaryOk() (*ResumeObjectsSalaryProperties, bool) {
	if o == nil {
		return nil, false
	}
	return o.Salary.Get(), o.Salary.IsSet()
}

// HasSalary returns a boolean if a field has been set.
func (o *VacanciesVisitorsVisitorItemsItemsInner) HasSalary() bool {
	if o != nil && o.Salary.IsSet() {
		return true
	}

	return false
}

// SetSalary gets a reference to the given NullableResumeObjectsSalaryProperties and assigns it to the Salary field.
func (o *VacanciesVisitorsVisitorItemsItemsInner) SetSalary(v ResumeObjectsSalaryProperties) {
	o.Salary.Set(&v)
}
// SetSalaryNil sets the value for Salary to be an explicit nil
func (o *VacanciesVisitorsVisitorItemsItemsInner) SetSalaryNil() {
	o.Salary.Set(nil)
}

// UnsetSalary ensures that no value is present for Salary, not even an explicit nil
func (o *VacanciesVisitorsVisitorItemsItemsInner) UnsetSalary() {
	o.Salary.Unset()
}

// GetTotalExperience returns the TotalExperience field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *VacanciesVisitorsVisitorItemsItemsInner) GetTotalExperience() ResumeObjectsTotalExperience {
	if o == nil || IsNil(o.TotalExperience.Get()) {
		var ret ResumeObjectsTotalExperience
		return ret
	}
	return *o.TotalExperience.Get()
}

// GetTotalExperienceOk returns a tuple with the TotalExperience field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *VacanciesVisitorsVisitorItemsItemsInner) GetTotalExperienceOk() (*ResumeObjectsTotalExperience, bool) {
	if o == nil {
		return nil, false
	}
	return o.TotalExperience.Get(), o.TotalExperience.IsSet()
}

// HasTotalExperience returns a boolean if a field has been set.
func (o *VacanciesVisitorsVisitorItemsItemsInner) HasTotalExperience() bool {
	if o != nil && o.TotalExperience.IsSet() {
		return true
	}

	return false
}

// SetTotalExperience gets a reference to the given NullableResumeObjectsTotalExperience and assigns it to the TotalExperience field.
func (o *VacanciesVisitorsVisitorItemsItemsInner) SetTotalExperience(v ResumeObjectsTotalExperience) {
	o.TotalExperience.Set(&v)
}
// SetTotalExperienceNil sets the value for TotalExperience to be an explicit nil
func (o *VacanciesVisitorsVisitorItemsItemsInner) SetTotalExperienceNil() {
	o.TotalExperience.Set(nil)
}

// UnsetTotalExperience ensures that no value is present for TotalExperience, not even an explicit nil
func (o *VacanciesVisitorsVisitorItemsItemsInner) UnsetTotalExperience() {
	o.TotalExperience.Unset()
}

// GetUpdatedAt returns the UpdatedAt field value
func (o *VacanciesVisitorsVisitorItemsItemsInner) GetUpdatedAt() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.UpdatedAt
}

// GetUpdatedAtOk returns a tuple with the UpdatedAt field value
// and a boolean to check if the value has been set.
func (o *VacanciesVisitorsVisitorItemsItemsInner) GetUpdatedAtOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.UpdatedAt, true
}

// SetUpdatedAt sets field value
func (o *VacanciesVisitorsVisitorItemsItemsInner) SetUpdatedAt(v string) {
	o.UpdatedAt = v
}

// GetActions returns the Actions field value
func (o *VacanciesVisitorsVisitorItemsItemsInner) GetActions() ResumeResumeShortAllOfActions {
	if o == nil {
		var ret ResumeResumeShortAllOfActions
		return ret
	}

	return o.Actions
}

// GetActionsOk returns a tuple with the Actions field value
// and a boolean to check if the value has been set.
func (o *VacanciesVisitorsVisitorItemsItemsInner) GetActionsOk() (*ResumeResumeShortAllOfActions, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Actions, true
}

// SetActions sets field value
func (o *VacanciesVisitorsVisitorItemsItemsInner) SetActions(v ResumeResumeShortAllOfActions) {
	o.Actions = v
}

// GetFavorited returns the Favorited field value
func (o *VacanciesVisitorsVisitorItemsItemsInner) GetFavorited() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.Favorited
}

// GetFavoritedOk returns a tuple with the Favorited field value
// and a boolean to check if the value has been set.
func (o *VacanciesVisitorsVisitorItemsItemsInner) GetFavoritedOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Favorited, true
}

// SetFavorited sets field value
func (o *VacanciesVisitorsVisitorItemsItemsInner) SetFavorited(v bool) {
	o.Favorited = v
}

// GetNegotiationsHistory returns the NegotiationsHistory field value
func (o *VacanciesVisitorsVisitorItemsItemsInner) GetNegotiationsHistory() ResumeResumeShortAllOfNegotiationsHistory {
	if o == nil {
		var ret ResumeResumeShortAllOfNegotiationsHistory
		return ret
	}

	return o.NegotiationsHistory
}

// GetNegotiationsHistoryOk returns a tuple with the NegotiationsHistory field value
// and a boolean to check if the value has been set.
func (o *VacanciesVisitorsVisitorItemsItemsInner) GetNegotiationsHistoryOk() (*ResumeResumeShortAllOfNegotiationsHistory, bool) {
	if o == nil {
		return nil, false
	}
	return &o.NegotiationsHistory, true
}

// SetNegotiationsHistory sets field value
func (o *VacanciesVisitorsVisitorItemsItemsInner) SetNegotiationsHistory(v ResumeResumeShortAllOfNegotiationsHistory) {
	o.NegotiationsHistory = v
}

// GetOwner returns the Owner field value
func (o *VacanciesVisitorsVisitorItemsItemsInner) GetOwner() ResumeResumeShortAllOfOwner {
	if o == nil {
		var ret ResumeResumeShortAllOfOwner
		return ret
	}

	return o.Owner
}

// GetOwnerOk returns a tuple with the Owner field value
// and a boolean to check if the value has been set.
func (o *VacanciesVisitorsVisitorItemsItemsInner) GetOwnerOk() (*ResumeResumeShortAllOfOwner, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Owner, true
}

// SetOwner sets field value
func (o *VacanciesVisitorsVisitorItemsItemsInner) SetOwner(v ResumeResumeShortAllOfOwner) {
	o.Owner = v
}

// GetPhoto returns the Photo field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *VacanciesVisitorsVisitorItemsItemsInner) GetPhoto() ResumeResumeShortAllOfPhoto {
	if o == nil || IsNil(o.Photo.Get()) {
		var ret ResumeResumeShortAllOfPhoto
		return ret
	}
	return *o.Photo.Get()
}

// GetPhotoOk returns a tuple with the Photo field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *VacanciesVisitorsVisitorItemsItemsInner) GetPhotoOk() (*ResumeResumeShortAllOfPhoto, bool) {
	if o == nil {
		return nil, false
	}
	return o.Photo.Get(), o.Photo.IsSet()
}

// HasPhoto returns a boolean if a field has been set.
func (o *VacanciesVisitorsVisitorItemsItemsInner) HasPhoto() bool {
	if o != nil && o.Photo.IsSet() {
		return true
	}

	return false
}

// SetPhoto gets a reference to the given NullableResumeResumeShortAllOfPhoto and assigns it to the Photo field.
func (o *VacanciesVisitorsVisitorItemsItemsInner) SetPhoto(v ResumeResumeShortAllOfPhoto) {
	o.Photo.Set(&v)
}
// SetPhotoNil sets the value for Photo to be an explicit nil
func (o *VacanciesVisitorsVisitorItemsItemsInner) SetPhotoNil() {
	o.Photo.Set(nil)
}

// UnsetPhoto ensures that no value is present for Photo, not even an explicit nil
func (o *VacanciesVisitorsVisitorItemsItemsInner) UnsetPhoto() {
	o.Photo.Unset()
}

// GetViewed returns the Viewed field value
func (o *VacanciesVisitorsVisitorItemsItemsInner) GetViewed() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.Viewed
}

// GetViewedOk returns a tuple with the Viewed field value
// and a boolean to check if the value has been set.
func (o *VacanciesVisitorsVisitorItemsItemsInner) GetViewedOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Viewed, true
}

// SetViewed sets field value
func (o *VacanciesVisitorsVisitorItemsItemsInner) SetViewed(v bool) {
	o.Viewed = v
}

// GetUrl returns the Url field value
func (o *VacanciesVisitorsVisitorItemsItemsInner) GetUrl() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Url
}

// GetUrlOk returns a tuple with the Url field value
// and a boolean to check if the value has been set.
func (o *VacanciesVisitorsVisitorItemsItemsInner) GetUrlOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Url, true
}

// SetUrl sets field value
func (o *VacanciesVisitorsVisitorItemsItemsInner) SetUrl(v string) {
	o.Url = v
}

func (o VacanciesVisitorsVisitorItemsItemsInner) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o VacanciesVisitorsVisitorItemsItemsInner) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["alternate_url"] = o.AlternateUrl
	toSerialize["id"] = o.Id
	toSerialize["title"] = o.Title.Get()
	if o.Age.IsSet() {
		toSerialize["age"] = o.Age.Get()
	}
	if o.Area.IsSet() {
		toSerialize["area"] = o.Area.Get()
	}
	if o.CanViewFullInfo.IsSet() {
		toSerialize["can_view_full_info"] = o.CanViewFullInfo.Get()
	}
	toSerialize["certificate"] = o.Certificate
	toSerialize["created_at"] = o.CreatedAt
	toSerialize["download"] = o.Download
	toSerialize["education"] = o.Education
	toSerialize["experience"] = o.Experience
	if o.FirstName.IsSet() {
		toSerialize["first_name"] = o.FirstName.Get()
	}
	if o.Gender.IsSet() {
		toSerialize["gender"] = o.Gender.Get()
	}
	toSerialize["hidden_fields"] = o.HiddenFields
	if o.LastName.IsSet() {
		toSerialize["last_name"] = o.LastName.Get()
	}
	if !IsNil(o.Marked) {
		toSerialize["marked"] = o.Marked
	}
	if o.MiddleName.IsSet() {
		toSerialize["middle_name"] = o.MiddleName.Get()
	}
	if !IsNil(o.Platform) {
		toSerialize["platform"] = o.Platform
	}
	if o.Salary.IsSet() {
		toSerialize["salary"] = o.Salary.Get()
	}
	if o.TotalExperience.IsSet() {
		toSerialize["total_experience"] = o.TotalExperience.Get()
	}
	toSerialize["updated_at"] = o.UpdatedAt
	toSerialize["actions"] = o.Actions
	toSerialize["favorited"] = o.Favorited
	toSerialize["negotiations_history"] = o.NegotiationsHistory
	toSerialize["owner"] = o.Owner
	if o.Photo.IsSet() {
		toSerialize["photo"] = o.Photo.Get()
	}
	toSerialize["viewed"] = o.Viewed
	toSerialize["url"] = o.Url
	return toSerialize, nil
}

func (o *VacanciesVisitorsVisitorItemsItemsInner) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"alternate_url",
		"id",
		"title",
		"certificate",
		"created_at",
		"download",
		"education",
		"experience",
		"hidden_fields",
		"updated_at",
		"actions",
		"favorited",
		"negotiations_history",
		"owner",
		"viewed",
		"url",
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

	varVacanciesVisitorsVisitorItemsItemsInner := _VacanciesVisitorsVisitorItemsItemsInner{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varVacanciesVisitorsVisitorItemsItemsInner)

	if err != nil {
		return err
	}

	*o = VacanciesVisitorsVisitorItemsItemsInner(varVacanciesVisitorsVisitorItemsItemsInner)

	return err
}

type NullableVacanciesVisitorsVisitorItemsItemsInner struct {
	value *VacanciesVisitorsVisitorItemsItemsInner
	isSet bool
}

func (v NullableVacanciesVisitorsVisitorItemsItemsInner) Get() *VacanciesVisitorsVisitorItemsItemsInner {
	return v.value
}

func (v *NullableVacanciesVisitorsVisitorItemsItemsInner) Set(val *VacanciesVisitorsVisitorItemsItemsInner) {
	v.value = val
	v.isSet = true
}

func (v NullableVacanciesVisitorsVisitorItemsItemsInner) IsSet() bool {
	return v.isSet
}

func (v *NullableVacanciesVisitorsVisitorItemsItemsInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableVacanciesVisitorsVisitorItemsItemsInner(val *VacanciesVisitorsVisitorItemsItemsInner) *NullableVacanciesVisitorsVisitorItemsItemsInner {
	return &NullableVacanciesVisitorsVisitorItemsItemsInner{value: val, isSet: true}
}

func (v NullableVacanciesVisitorsVisitorItemsItemsInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableVacanciesVisitorsVisitorItemsItemsInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


