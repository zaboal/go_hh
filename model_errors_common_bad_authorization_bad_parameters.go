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
	"fmt"
	"bytes"
)

// ErrorsCommonBadAuthorizationBadParameters - struct for ErrorsCommonBadAuthorizationBadParameters
type ErrorsCommonBadAuthorizationBadParameters struct {
	ErrorsCommonBadAuthorizationInvalidClientError *ErrorsCommonBadAuthorizationInvalidClientError
	ErrorsCommonBadAuthorizationInvalidGrantError *ErrorsCommonBadAuthorizationInvalidGrantError
	ErrorsCommonBadAuthorizationInvalidRequestError *ErrorsCommonBadAuthorizationInvalidRequestError
	ErrorsCommonBadAuthorizationUnsupportedGrantTypeError *ErrorsCommonBadAuthorizationUnsupportedGrantTypeError
}

// ErrorsCommonBadAuthorizationInvalidClientErrorAsErrorsCommonBadAuthorizationBadParameters is a convenience function that returns ErrorsCommonBadAuthorizationInvalidClientError wrapped in ErrorsCommonBadAuthorizationBadParameters
func ErrorsCommonBadAuthorizationInvalidClientErrorAsErrorsCommonBadAuthorizationBadParameters(v *ErrorsCommonBadAuthorizationInvalidClientError) ErrorsCommonBadAuthorizationBadParameters {
	return ErrorsCommonBadAuthorizationBadParameters{
		ErrorsCommonBadAuthorizationInvalidClientError: v,
	}
}

// ErrorsCommonBadAuthorizationInvalidGrantErrorAsErrorsCommonBadAuthorizationBadParameters is a convenience function that returns ErrorsCommonBadAuthorizationInvalidGrantError wrapped in ErrorsCommonBadAuthorizationBadParameters
func ErrorsCommonBadAuthorizationInvalidGrantErrorAsErrorsCommonBadAuthorizationBadParameters(v *ErrorsCommonBadAuthorizationInvalidGrantError) ErrorsCommonBadAuthorizationBadParameters {
	return ErrorsCommonBadAuthorizationBadParameters{
		ErrorsCommonBadAuthorizationInvalidGrantError: v,
	}
}

// ErrorsCommonBadAuthorizationInvalidRequestErrorAsErrorsCommonBadAuthorizationBadParameters is a convenience function that returns ErrorsCommonBadAuthorizationInvalidRequestError wrapped in ErrorsCommonBadAuthorizationBadParameters
func ErrorsCommonBadAuthorizationInvalidRequestErrorAsErrorsCommonBadAuthorizationBadParameters(v *ErrorsCommonBadAuthorizationInvalidRequestError) ErrorsCommonBadAuthorizationBadParameters {
	return ErrorsCommonBadAuthorizationBadParameters{
		ErrorsCommonBadAuthorizationInvalidRequestError: v,
	}
}

// ErrorsCommonBadAuthorizationUnsupportedGrantTypeErrorAsErrorsCommonBadAuthorizationBadParameters is a convenience function that returns ErrorsCommonBadAuthorizationUnsupportedGrantTypeError wrapped in ErrorsCommonBadAuthorizationBadParameters
func ErrorsCommonBadAuthorizationUnsupportedGrantTypeErrorAsErrorsCommonBadAuthorizationBadParameters(v *ErrorsCommonBadAuthorizationUnsupportedGrantTypeError) ErrorsCommonBadAuthorizationBadParameters {
	return ErrorsCommonBadAuthorizationBadParameters{
		ErrorsCommonBadAuthorizationUnsupportedGrantTypeError: v,
	}
}


// Unmarshal JSON data into one of the pointers in the struct
func (dst *ErrorsCommonBadAuthorizationBadParameters) UnmarshalJSON(data []byte) error {
	var err error
	match := 0
	// try to unmarshal data into ErrorsCommonBadAuthorizationInvalidClientError
	err = newStrictDecoder(data).Decode(&dst.ErrorsCommonBadAuthorizationInvalidClientError)
	if err == nil {
		jsonErrorsCommonBadAuthorizationInvalidClientError, _ := json.Marshal(dst.ErrorsCommonBadAuthorizationInvalidClientError)
		if string(jsonErrorsCommonBadAuthorizationInvalidClientError) == "{}" { // empty struct
			dst.ErrorsCommonBadAuthorizationInvalidClientError = nil
		} else {
			match++
		}
	} else {
		dst.ErrorsCommonBadAuthorizationInvalidClientError = nil
	}

	// try to unmarshal data into ErrorsCommonBadAuthorizationInvalidGrantError
	err = newStrictDecoder(data).Decode(&dst.ErrorsCommonBadAuthorizationInvalidGrantError)
	if err == nil {
		jsonErrorsCommonBadAuthorizationInvalidGrantError, _ := json.Marshal(dst.ErrorsCommonBadAuthorizationInvalidGrantError)
		if string(jsonErrorsCommonBadAuthorizationInvalidGrantError) == "{}" { // empty struct
			dst.ErrorsCommonBadAuthorizationInvalidGrantError = nil
		} else {
			match++
		}
	} else {
		dst.ErrorsCommonBadAuthorizationInvalidGrantError = nil
	}

	// try to unmarshal data into ErrorsCommonBadAuthorizationInvalidRequestError
	err = newStrictDecoder(data).Decode(&dst.ErrorsCommonBadAuthorizationInvalidRequestError)
	if err == nil {
		jsonErrorsCommonBadAuthorizationInvalidRequestError, _ := json.Marshal(dst.ErrorsCommonBadAuthorizationInvalidRequestError)
		if string(jsonErrorsCommonBadAuthorizationInvalidRequestError) == "{}" { // empty struct
			dst.ErrorsCommonBadAuthorizationInvalidRequestError = nil
		} else {
			match++
		}
	} else {
		dst.ErrorsCommonBadAuthorizationInvalidRequestError = nil
	}

	// try to unmarshal data into ErrorsCommonBadAuthorizationUnsupportedGrantTypeError
	err = newStrictDecoder(data).Decode(&dst.ErrorsCommonBadAuthorizationUnsupportedGrantTypeError)
	if err == nil {
		jsonErrorsCommonBadAuthorizationUnsupportedGrantTypeError, _ := json.Marshal(dst.ErrorsCommonBadAuthorizationUnsupportedGrantTypeError)
		if string(jsonErrorsCommonBadAuthorizationUnsupportedGrantTypeError) == "{}" { // empty struct
			dst.ErrorsCommonBadAuthorizationUnsupportedGrantTypeError = nil
		} else {
			match++
		}
	} else {
		dst.ErrorsCommonBadAuthorizationUnsupportedGrantTypeError = nil
	}

	if match > 1 { // more than 1 match
		// reset to nil
		dst.ErrorsCommonBadAuthorizationInvalidClientError = nil
		dst.ErrorsCommonBadAuthorizationInvalidGrantError = nil
		dst.ErrorsCommonBadAuthorizationInvalidRequestError = nil
		dst.ErrorsCommonBadAuthorizationUnsupportedGrantTypeError = nil

		return fmt.Errorf("data matches more than one schema in oneOf(ErrorsCommonBadAuthorizationBadParameters)")
	} else if match == 1 {
		return nil // exactly one match
	} else { // no match
		return fmt.Errorf("data failed to match schemas in oneOf(ErrorsCommonBadAuthorizationBadParameters)")
	}
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src ErrorsCommonBadAuthorizationBadParameters) MarshalJSON() ([]byte, error) {
	if src.ErrorsCommonBadAuthorizationInvalidClientError != nil {
		return json.Marshal(&src.ErrorsCommonBadAuthorizationInvalidClientError)
	}

	if src.ErrorsCommonBadAuthorizationInvalidGrantError != nil {
		return json.Marshal(&src.ErrorsCommonBadAuthorizationInvalidGrantError)
	}

	if src.ErrorsCommonBadAuthorizationInvalidRequestError != nil {
		return json.Marshal(&src.ErrorsCommonBadAuthorizationInvalidRequestError)
	}

	if src.ErrorsCommonBadAuthorizationUnsupportedGrantTypeError != nil {
		return json.Marshal(&src.ErrorsCommonBadAuthorizationUnsupportedGrantTypeError)
	}

	return nil, nil // no data in oneOf schemas
}

// Get the actual instance
func (obj *ErrorsCommonBadAuthorizationBadParameters) GetActualInstance() (interface{}) {
	if obj == nil {
		return nil
	}
	if obj.ErrorsCommonBadAuthorizationInvalidClientError != nil {
		return obj.ErrorsCommonBadAuthorizationInvalidClientError
	}

	if obj.ErrorsCommonBadAuthorizationInvalidGrantError != nil {
		return obj.ErrorsCommonBadAuthorizationInvalidGrantError
	}

	if obj.ErrorsCommonBadAuthorizationInvalidRequestError != nil {
		return obj.ErrorsCommonBadAuthorizationInvalidRequestError
	}

	if obj.ErrorsCommonBadAuthorizationUnsupportedGrantTypeError != nil {
		return obj.ErrorsCommonBadAuthorizationUnsupportedGrantTypeError
	}

	// all schemas are nil
	return nil
}

type NullableErrorsCommonBadAuthorizationBadParameters struct {
	value *ErrorsCommonBadAuthorizationBadParameters
	isSet bool
}

func (v NullableErrorsCommonBadAuthorizationBadParameters) Get() *ErrorsCommonBadAuthorizationBadParameters {
	return v.value
}

func (v *NullableErrorsCommonBadAuthorizationBadParameters) Set(val *ErrorsCommonBadAuthorizationBadParameters) {
	v.value = val
	v.isSet = true
}

func (v NullableErrorsCommonBadAuthorizationBadParameters) IsSet() bool {
	return v.isSet
}

func (v *NullableErrorsCommonBadAuthorizationBadParameters) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableErrorsCommonBadAuthorizationBadParameters(val *ErrorsCommonBadAuthorizationBadParameters) *NullableErrorsCommonBadAuthorizationBadParameters {
	return &NullableErrorsCommonBadAuthorizationBadParameters{value: val, isSet: true}
}

func (v NullableErrorsCommonBadAuthorizationBadParameters) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableErrorsCommonBadAuthorizationBadParameters) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


