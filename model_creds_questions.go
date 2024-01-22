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
)

// checks if the CredsQuestions type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CredsQuestions{}

// CredsQuestions struct for CredsQuestions
type CredsQuestions struct {
	// Описание вопроса
	Description NullableString `json:"description,omitempty"`
	// Показан ли вопрос изначально, актуально для динамических вопросов
	IsActive *bool `json:"is_active,omitempty"`
	// Возможные ответы на вопрос, гарантировано придут в поле answers
	PossibleAnswers []string `json:"possible_answers,omitempty"`
	// Идентификатор вопроса (совпадает с ключом объекта)
	QuestionId *string `json:"question_id,omitempty"`
	// Текст вопроса отображаемый на форме
	QuestionTitle *string `json:"question_title,omitempty"`
	// Возможность мульти выбора ответов на данный вопрос \"single_choice\" / \"multi_select\"
	QuestionType *string `json:"question_type,omitempty"`
	// Обязателен ли вопрос для получения ответа
	Required *bool `json:"required,omitempty"`
	// Пропускать ли текст вопроса на просмотре, если false - ответы внутри placeholder, если true - просто перечисляем без текста вопроса
	SkipTitleAtView *bool `json:"skip_title_at_view,omitempty"`
	// Текст вопроса на просмотре
	ViewTitle NullableString `json:"view_title,omitempty"`
}

// NewCredsQuestions instantiates a new CredsQuestions object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCredsQuestions() *CredsQuestions {
	this := CredsQuestions{}
	return &this
}

// NewCredsQuestionsWithDefaults instantiates a new CredsQuestions object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCredsQuestionsWithDefaults() *CredsQuestions {
	this := CredsQuestions{}
	return &this
}

// GetDescription returns the Description field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CredsQuestions) GetDescription() string {
	if o == nil || IsNil(o.Description.Get()) {
		var ret string
		return ret
	}
	return *o.Description.Get()
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CredsQuestions) GetDescriptionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Description.Get(), o.Description.IsSet()
}

// HasDescription returns a boolean if a field has been set.
func (o *CredsQuestions) HasDescription() bool {
	if o != nil && o.Description.IsSet() {
		return true
	}

	return false
}

// SetDescription gets a reference to the given NullableString and assigns it to the Description field.
func (o *CredsQuestions) SetDescription(v string) {
	o.Description.Set(&v)
}
// SetDescriptionNil sets the value for Description to be an explicit nil
func (o *CredsQuestions) SetDescriptionNil() {
	o.Description.Set(nil)
}

// UnsetDescription ensures that no value is present for Description, not even an explicit nil
func (o *CredsQuestions) UnsetDescription() {
	o.Description.Unset()
}

// GetIsActive returns the IsActive field value if set, zero value otherwise.
func (o *CredsQuestions) GetIsActive() bool {
	if o == nil || IsNil(o.IsActive) {
		var ret bool
		return ret
	}
	return *o.IsActive
}

// GetIsActiveOk returns a tuple with the IsActive field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CredsQuestions) GetIsActiveOk() (*bool, bool) {
	if o == nil || IsNil(o.IsActive) {
		return nil, false
	}
	return o.IsActive, true
}

// HasIsActive returns a boolean if a field has been set.
func (o *CredsQuestions) HasIsActive() bool {
	if o != nil && !IsNil(o.IsActive) {
		return true
	}

	return false
}

// SetIsActive gets a reference to the given bool and assigns it to the IsActive field.
func (o *CredsQuestions) SetIsActive(v bool) {
	o.IsActive = &v
}

// GetPossibleAnswers returns the PossibleAnswers field value if set, zero value otherwise.
func (o *CredsQuestions) GetPossibleAnswers() []string {
	if o == nil || IsNil(o.PossibleAnswers) {
		var ret []string
		return ret
	}
	return o.PossibleAnswers
}

// GetPossibleAnswersOk returns a tuple with the PossibleAnswers field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CredsQuestions) GetPossibleAnswersOk() ([]string, bool) {
	if o == nil || IsNil(o.PossibleAnswers) {
		return nil, false
	}
	return o.PossibleAnswers, true
}

// HasPossibleAnswers returns a boolean if a field has been set.
func (o *CredsQuestions) HasPossibleAnswers() bool {
	if o != nil && !IsNil(o.PossibleAnswers) {
		return true
	}

	return false
}

// SetPossibleAnswers gets a reference to the given []string and assigns it to the PossibleAnswers field.
func (o *CredsQuestions) SetPossibleAnswers(v []string) {
	o.PossibleAnswers = v
}

// GetQuestionId returns the QuestionId field value if set, zero value otherwise.
func (o *CredsQuestions) GetQuestionId() string {
	if o == nil || IsNil(o.QuestionId) {
		var ret string
		return ret
	}
	return *o.QuestionId
}

// GetQuestionIdOk returns a tuple with the QuestionId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CredsQuestions) GetQuestionIdOk() (*string, bool) {
	if o == nil || IsNil(o.QuestionId) {
		return nil, false
	}
	return o.QuestionId, true
}

// HasQuestionId returns a boolean if a field has been set.
func (o *CredsQuestions) HasQuestionId() bool {
	if o != nil && !IsNil(o.QuestionId) {
		return true
	}

	return false
}

// SetQuestionId gets a reference to the given string and assigns it to the QuestionId field.
func (o *CredsQuestions) SetQuestionId(v string) {
	o.QuestionId = &v
}

// GetQuestionTitle returns the QuestionTitle field value if set, zero value otherwise.
func (o *CredsQuestions) GetQuestionTitle() string {
	if o == nil || IsNil(o.QuestionTitle) {
		var ret string
		return ret
	}
	return *o.QuestionTitle
}

// GetQuestionTitleOk returns a tuple with the QuestionTitle field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CredsQuestions) GetQuestionTitleOk() (*string, bool) {
	if o == nil || IsNil(o.QuestionTitle) {
		return nil, false
	}
	return o.QuestionTitle, true
}

// HasQuestionTitle returns a boolean if a field has been set.
func (o *CredsQuestions) HasQuestionTitle() bool {
	if o != nil && !IsNil(o.QuestionTitle) {
		return true
	}

	return false
}

// SetQuestionTitle gets a reference to the given string and assigns it to the QuestionTitle field.
func (o *CredsQuestions) SetQuestionTitle(v string) {
	o.QuestionTitle = &v
}

// GetQuestionType returns the QuestionType field value if set, zero value otherwise.
func (o *CredsQuestions) GetQuestionType() string {
	if o == nil || IsNil(o.QuestionType) {
		var ret string
		return ret
	}
	return *o.QuestionType
}

// GetQuestionTypeOk returns a tuple with the QuestionType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CredsQuestions) GetQuestionTypeOk() (*string, bool) {
	if o == nil || IsNil(o.QuestionType) {
		return nil, false
	}
	return o.QuestionType, true
}

// HasQuestionType returns a boolean if a field has been set.
func (o *CredsQuestions) HasQuestionType() bool {
	if o != nil && !IsNil(o.QuestionType) {
		return true
	}

	return false
}

// SetQuestionType gets a reference to the given string and assigns it to the QuestionType field.
func (o *CredsQuestions) SetQuestionType(v string) {
	o.QuestionType = &v
}

// GetRequired returns the Required field value if set, zero value otherwise.
func (o *CredsQuestions) GetRequired() bool {
	if o == nil || IsNil(o.Required) {
		var ret bool
		return ret
	}
	return *o.Required
}

// GetRequiredOk returns a tuple with the Required field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CredsQuestions) GetRequiredOk() (*bool, bool) {
	if o == nil || IsNil(o.Required) {
		return nil, false
	}
	return o.Required, true
}

// HasRequired returns a boolean if a field has been set.
func (o *CredsQuestions) HasRequired() bool {
	if o != nil && !IsNil(o.Required) {
		return true
	}

	return false
}

// SetRequired gets a reference to the given bool and assigns it to the Required field.
func (o *CredsQuestions) SetRequired(v bool) {
	o.Required = &v
}

// GetSkipTitleAtView returns the SkipTitleAtView field value if set, zero value otherwise.
func (o *CredsQuestions) GetSkipTitleAtView() bool {
	if o == nil || IsNil(o.SkipTitleAtView) {
		var ret bool
		return ret
	}
	return *o.SkipTitleAtView
}

// GetSkipTitleAtViewOk returns a tuple with the SkipTitleAtView field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CredsQuestions) GetSkipTitleAtViewOk() (*bool, bool) {
	if o == nil || IsNil(o.SkipTitleAtView) {
		return nil, false
	}
	return o.SkipTitleAtView, true
}

// HasSkipTitleAtView returns a boolean if a field has been set.
func (o *CredsQuestions) HasSkipTitleAtView() bool {
	if o != nil && !IsNil(o.SkipTitleAtView) {
		return true
	}

	return false
}

// SetSkipTitleAtView gets a reference to the given bool and assigns it to the SkipTitleAtView field.
func (o *CredsQuestions) SetSkipTitleAtView(v bool) {
	o.SkipTitleAtView = &v
}

// GetViewTitle returns the ViewTitle field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CredsQuestions) GetViewTitle() string {
	if o == nil || IsNil(o.ViewTitle.Get()) {
		var ret string
		return ret
	}
	return *o.ViewTitle.Get()
}

// GetViewTitleOk returns a tuple with the ViewTitle field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CredsQuestions) GetViewTitleOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.ViewTitle.Get(), o.ViewTitle.IsSet()
}

// HasViewTitle returns a boolean if a field has been set.
func (o *CredsQuestions) HasViewTitle() bool {
	if o != nil && o.ViewTitle.IsSet() {
		return true
	}

	return false
}

// SetViewTitle gets a reference to the given NullableString and assigns it to the ViewTitle field.
func (o *CredsQuestions) SetViewTitle(v string) {
	o.ViewTitle.Set(&v)
}
// SetViewTitleNil sets the value for ViewTitle to be an explicit nil
func (o *CredsQuestions) SetViewTitleNil() {
	o.ViewTitle.Set(nil)
}

// UnsetViewTitle ensures that no value is present for ViewTitle, not even an explicit nil
func (o *CredsQuestions) UnsetViewTitle() {
	o.ViewTitle.Unset()
}

func (o CredsQuestions) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CredsQuestions) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.Description.IsSet() {
		toSerialize["description"] = o.Description.Get()
	}
	if !IsNil(o.IsActive) {
		toSerialize["is_active"] = o.IsActive
	}
	if !IsNil(o.PossibleAnswers) {
		toSerialize["possible_answers"] = o.PossibleAnswers
	}
	if !IsNil(o.QuestionId) {
		toSerialize["question_id"] = o.QuestionId
	}
	if !IsNil(o.QuestionTitle) {
		toSerialize["question_title"] = o.QuestionTitle
	}
	if !IsNil(o.QuestionType) {
		toSerialize["question_type"] = o.QuestionType
	}
	if !IsNil(o.Required) {
		toSerialize["required"] = o.Required
	}
	if !IsNil(o.SkipTitleAtView) {
		toSerialize["skip_title_at_view"] = o.SkipTitleAtView
	}
	if o.ViewTitle.IsSet() {
		toSerialize["view_title"] = o.ViewTitle.Get()
	}
	return toSerialize, nil
}

type NullableCredsQuestions struct {
	value *CredsQuestions
	isSet bool
}

func (v NullableCredsQuestions) Get() *CredsQuestions {
	return v.value
}

func (v *NullableCredsQuestions) Set(val *CredsQuestions) {
	v.value = val
	v.isSet = true
}

func (v NullableCredsQuestions) IsSet() bool {
	return v.isSet
}

func (v *NullableCredsQuestions) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCredsQuestions(val *CredsQuestions) *NullableCredsQuestions {
	return &NullableCredsQuestions{value: val, isSet: true}
}

func (v NullableCredsQuestions) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCredsQuestions) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


