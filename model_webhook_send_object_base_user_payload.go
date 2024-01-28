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
	"fmt"
	"bytes"
)

// WebhookSendObjectBaseUserPayload - struct for WebhookSendObjectBaseUserPayload
type WebhookSendObjectBaseUserPayload struct {
	WebhookPayloadNegotiationEmployerStateChange *WebhookPayloadNegotiationEmployerStateChange
	WebhookPayloadNewNegotiationVacancy *WebhookPayloadNewNegotiationVacancy
	WebhookPayloadNewResponseOrInvitationVacancy *WebhookPayloadNewResponseOrInvitationVacancy
	WebhookPayloadVacancyArchivation *WebhookPayloadVacancyArchivation
	WebhookPayloadVacancyChange *WebhookPayloadVacancyChange
	WebhookPayloadVacancyProlongation *WebhookPayloadVacancyProlongation
	WebhookPayloadVacancyPublicationForVacancyManager *WebhookPayloadVacancyPublicationForVacancyManager
}

// WebhookPayloadNegotiationEmployerStateChangeAsWebhookSendObjectBaseUserPayload is a convenience function that returns WebhookPayloadNegotiationEmployerStateChange wrapped in WebhookSendObjectBaseUserPayload
func WebhookPayloadNegotiationEmployerStateChangeAsWebhookSendObjectBaseUserPayload(v *WebhookPayloadNegotiationEmployerStateChange) WebhookSendObjectBaseUserPayload {
	return WebhookSendObjectBaseUserPayload{
		WebhookPayloadNegotiationEmployerStateChange: v,
	}
}

// WebhookPayloadNewNegotiationVacancyAsWebhookSendObjectBaseUserPayload is a convenience function that returns WebhookPayloadNewNegotiationVacancy wrapped in WebhookSendObjectBaseUserPayload
func WebhookPayloadNewNegotiationVacancyAsWebhookSendObjectBaseUserPayload(v *WebhookPayloadNewNegotiationVacancy) WebhookSendObjectBaseUserPayload {
	return WebhookSendObjectBaseUserPayload{
		WebhookPayloadNewNegotiationVacancy: v,
	}
}

// WebhookPayloadNewResponseOrInvitationVacancyAsWebhookSendObjectBaseUserPayload is a convenience function that returns WebhookPayloadNewResponseOrInvitationVacancy wrapped in WebhookSendObjectBaseUserPayload
func WebhookPayloadNewResponseOrInvitationVacancyAsWebhookSendObjectBaseUserPayload(v *WebhookPayloadNewResponseOrInvitationVacancy) WebhookSendObjectBaseUserPayload {
	return WebhookSendObjectBaseUserPayload{
		WebhookPayloadNewResponseOrInvitationVacancy: v,
	}
}

// WebhookPayloadVacancyArchivationAsWebhookSendObjectBaseUserPayload is a convenience function that returns WebhookPayloadVacancyArchivation wrapped in WebhookSendObjectBaseUserPayload
func WebhookPayloadVacancyArchivationAsWebhookSendObjectBaseUserPayload(v *WebhookPayloadVacancyArchivation) WebhookSendObjectBaseUserPayload {
	return WebhookSendObjectBaseUserPayload{
		WebhookPayloadVacancyArchivation: v,
	}
}

// WebhookPayloadVacancyChangeAsWebhookSendObjectBaseUserPayload is a convenience function that returns WebhookPayloadVacancyChange wrapped in WebhookSendObjectBaseUserPayload
func WebhookPayloadVacancyChangeAsWebhookSendObjectBaseUserPayload(v *WebhookPayloadVacancyChange) WebhookSendObjectBaseUserPayload {
	return WebhookSendObjectBaseUserPayload{
		WebhookPayloadVacancyChange: v,
	}
}

// WebhookPayloadVacancyProlongationAsWebhookSendObjectBaseUserPayload is a convenience function that returns WebhookPayloadVacancyProlongation wrapped in WebhookSendObjectBaseUserPayload
func WebhookPayloadVacancyProlongationAsWebhookSendObjectBaseUserPayload(v *WebhookPayloadVacancyProlongation) WebhookSendObjectBaseUserPayload {
	return WebhookSendObjectBaseUserPayload{
		WebhookPayloadVacancyProlongation: v,
	}
}

// WebhookPayloadVacancyPublicationForVacancyManagerAsWebhookSendObjectBaseUserPayload is a convenience function that returns WebhookPayloadVacancyPublicationForVacancyManager wrapped in WebhookSendObjectBaseUserPayload
func WebhookPayloadVacancyPublicationForVacancyManagerAsWebhookSendObjectBaseUserPayload(v *WebhookPayloadVacancyPublicationForVacancyManager) WebhookSendObjectBaseUserPayload {
	return WebhookSendObjectBaseUserPayload{
		WebhookPayloadVacancyPublicationForVacancyManager: v,
	}
}


// Unmarshal JSON data into one of the pointers in the struct
func (dst *WebhookSendObjectBaseUserPayload) UnmarshalJSON(data []byte) error {
	var err error
	match := 0
	// try to unmarshal data into WebhookPayloadNegotiationEmployerStateChange
	err = newStrictDecoder(data).Decode(&dst.WebhookPayloadNegotiationEmployerStateChange)
	if err == nil {
		jsonWebhookPayloadNegotiationEmployerStateChange, _ := json.Marshal(dst.WebhookPayloadNegotiationEmployerStateChange)
		if string(jsonWebhookPayloadNegotiationEmployerStateChange) == "{}" { // empty struct
			dst.WebhookPayloadNegotiationEmployerStateChange = nil
		} else {
			match++
		}
	} else {
		dst.WebhookPayloadNegotiationEmployerStateChange = nil
	}

	// try to unmarshal data into WebhookPayloadNewNegotiationVacancy
	err = newStrictDecoder(data).Decode(&dst.WebhookPayloadNewNegotiationVacancy)
	if err == nil {
		jsonWebhookPayloadNewNegotiationVacancy, _ := json.Marshal(dst.WebhookPayloadNewNegotiationVacancy)
		if string(jsonWebhookPayloadNewNegotiationVacancy) == "{}" { // empty struct
			dst.WebhookPayloadNewNegotiationVacancy = nil
		} else {
			match++
		}
	} else {
		dst.WebhookPayloadNewNegotiationVacancy = nil
	}

	// try to unmarshal data into WebhookPayloadNewResponseOrInvitationVacancy
	err = newStrictDecoder(data).Decode(&dst.WebhookPayloadNewResponseOrInvitationVacancy)
	if err == nil {
		jsonWebhookPayloadNewResponseOrInvitationVacancy, _ := json.Marshal(dst.WebhookPayloadNewResponseOrInvitationVacancy)
		if string(jsonWebhookPayloadNewResponseOrInvitationVacancy) == "{}" { // empty struct
			dst.WebhookPayloadNewResponseOrInvitationVacancy = nil
		} else {
			match++
		}
	} else {
		dst.WebhookPayloadNewResponseOrInvitationVacancy = nil
	}

	// try to unmarshal data into WebhookPayloadVacancyArchivation
	err = newStrictDecoder(data).Decode(&dst.WebhookPayloadVacancyArchivation)
	if err == nil {
		jsonWebhookPayloadVacancyArchivation, _ := json.Marshal(dst.WebhookPayloadVacancyArchivation)
		if string(jsonWebhookPayloadVacancyArchivation) == "{}" { // empty struct
			dst.WebhookPayloadVacancyArchivation = nil
		} else {
			match++
		}
	} else {
		dst.WebhookPayloadVacancyArchivation = nil
	}

	// try to unmarshal data into WebhookPayloadVacancyChange
	err = newStrictDecoder(data).Decode(&dst.WebhookPayloadVacancyChange)
	if err == nil {
		jsonWebhookPayloadVacancyChange, _ := json.Marshal(dst.WebhookPayloadVacancyChange)
		if string(jsonWebhookPayloadVacancyChange) == "{}" { // empty struct
			dst.WebhookPayloadVacancyChange = nil
		} else {
			match++
		}
	} else {
		dst.WebhookPayloadVacancyChange = nil
	}

	// try to unmarshal data into WebhookPayloadVacancyProlongation
	err = newStrictDecoder(data).Decode(&dst.WebhookPayloadVacancyProlongation)
	if err == nil {
		jsonWebhookPayloadVacancyProlongation, _ := json.Marshal(dst.WebhookPayloadVacancyProlongation)
		if string(jsonWebhookPayloadVacancyProlongation) == "{}" { // empty struct
			dst.WebhookPayloadVacancyProlongation = nil
		} else {
			match++
		}
	} else {
		dst.WebhookPayloadVacancyProlongation = nil
	}

	// try to unmarshal data into WebhookPayloadVacancyPublicationForVacancyManager
	err = newStrictDecoder(data).Decode(&dst.WebhookPayloadVacancyPublicationForVacancyManager)
	if err == nil {
		jsonWebhookPayloadVacancyPublicationForVacancyManager, _ := json.Marshal(dst.WebhookPayloadVacancyPublicationForVacancyManager)
		if string(jsonWebhookPayloadVacancyPublicationForVacancyManager) == "{}" { // empty struct
			dst.WebhookPayloadVacancyPublicationForVacancyManager = nil
		} else {
			match++
		}
	} else {
		dst.WebhookPayloadVacancyPublicationForVacancyManager = nil
	}

	if match > 1 { // more than 1 match
		// reset to nil
		dst.WebhookPayloadNegotiationEmployerStateChange = nil
		dst.WebhookPayloadNewNegotiationVacancy = nil
		dst.WebhookPayloadNewResponseOrInvitationVacancy = nil
		dst.WebhookPayloadVacancyArchivation = nil
		dst.WebhookPayloadVacancyChange = nil
		dst.WebhookPayloadVacancyProlongation = nil
		dst.WebhookPayloadVacancyPublicationForVacancyManager = nil

		return fmt.Errorf("data matches more than one schema in oneOf(WebhookSendObjectBaseUserPayload)")
	} else if match == 1 {
		return nil // exactly one match
	} else { // no match
		return fmt.Errorf("data failed to match schemas in oneOf(WebhookSendObjectBaseUserPayload)")
	}
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src WebhookSendObjectBaseUserPayload) MarshalJSON() ([]byte, error) {
	if src.WebhookPayloadNegotiationEmployerStateChange != nil {
		return json.Marshal(&src.WebhookPayloadNegotiationEmployerStateChange)
	}

	if src.WebhookPayloadNewNegotiationVacancy != nil {
		return json.Marshal(&src.WebhookPayloadNewNegotiationVacancy)
	}

	if src.WebhookPayloadNewResponseOrInvitationVacancy != nil {
		return json.Marshal(&src.WebhookPayloadNewResponseOrInvitationVacancy)
	}

	if src.WebhookPayloadVacancyArchivation != nil {
		return json.Marshal(&src.WebhookPayloadVacancyArchivation)
	}

	if src.WebhookPayloadVacancyChange != nil {
		return json.Marshal(&src.WebhookPayloadVacancyChange)
	}

	if src.WebhookPayloadVacancyProlongation != nil {
		return json.Marshal(&src.WebhookPayloadVacancyProlongation)
	}

	if src.WebhookPayloadVacancyPublicationForVacancyManager != nil {
		return json.Marshal(&src.WebhookPayloadVacancyPublicationForVacancyManager)
	}

	return nil, nil // no data in oneOf schemas
}

// Get the actual instance
func (obj *WebhookSendObjectBaseUserPayload) GetActualInstance() (interface{}) {
	if obj == nil {
		return nil
	}
	if obj.WebhookPayloadNegotiationEmployerStateChange != nil {
		return obj.WebhookPayloadNegotiationEmployerStateChange
	}

	if obj.WebhookPayloadNewNegotiationVacancy != nil {
		return obj.WebhookPayloadNewNegotiationVacancy
	}

	if obj.WebhookPayloadNewResponseOrInvitationVacancy != nil {
		return obj.WebhookPayloadNewResponseOrInvitationVacancy
	}

	if obj.WebhookPayloadVacancyArchivation != nil {
		return obj.WebhookPayloadVacancyArchivation
	}

	if obj.WebhookPayloadVacancyChange != nil {
		return obj.WebhookPayloadVacancyChange
	}

	if obj.WebhookPayloadVacancyProlongation != nil {
		return obj.WebhookPayloadVacancyProlongation
	}

	if obj.WebhookPayloadVacancyPublicationForVacancyManager != nil {
		return obj.WebhookPayloadVacancyPublicationForVacancyManager
	}

	// all schemas are nil
	return nil
}

type NullableWebhookSendObjectBaseUserPayload struct {
	value *WebhookSendObjectBaseUserPayload
	isSet bool
}

func (v NullableWebhookSendObjectBaseUserPayload) Get() *WebhookSendObjectBaseUserPayload {
	return v.value
}

func (v *NullableWebhookSendObjectBaseUserPayload) Set(val *WebhookSendObjectBaseUserPayload) {
	v.value = val
	v.isSet = true
}

func (v NullableWebhookSendObjectBaseUserPayload) IsSet() bool {
	return v.isSet
}

func (v *NullableWebhookSendObjectBaseUserPayload) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableWebhookSendObjectBaseUserPayload(val *WebhookSendObjectBaseUserPayload) *NullableWebhookSendObjectBaseUserPayload {
	return &NullableWebhookSendObjectBaseUserPayload{value: val, isSet: true}
}

func (v NullableWebhookSendObjectBaseUserPayload) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableWebhookSendObjectBaseUserPayload) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


