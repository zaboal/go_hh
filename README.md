# Go API client for hh

По-русски | [Switch to English](https://api.hh.ru/openapi/en/redoc)

В OpenAPI ведется пока что только небольшая часть документации
[Основная документация](https://github.com/hhru/api).

Для поиска по документации можно использовать Ctrl+F.

# Общая информация

* Всё API работает по протоколу HTTPS.
* Авторизация осуществляется по протоколу OAuth2.
* Все данные доступны только в формате JSON.
* Базовый URL — `https://api.hh.ru/`
* Возможны запросы к данным [любого сайта группы компаний HeadHunter](#section/Obshaya-informaciya/Vybor-sajta)
* <a name=\"date-format\"></a> Даты форматируются в соответствии с
[ISO 8601](http://en.wikipedia.org/wiki/ISO_8601): `YYYY-MM-DDThh:mm:ss±hhmm`.


<a name=\"request-requirements\"></a>
## Требования к запросам

В запросе необходимо передавать заголовок `User-Agent`, но если ваша реализация
http клиента не позволяет, можно отправить `HH-User-Agent`. Если не отправлен ни
один заголовок, то ответом будет `400 Bad Request`.
Указание в заголовке названия приложения и контактной почты разработчика
позволит нам оперативно с вами связаться в случае необходимости.
Заголовки `User-Agent` и `HH-User-Agent` взаимозаменяемы, в случае, если вы отправите оба заголовка,
обработан будет только `HH-User-Agent`.

```
User-Agent: MyApp/1.0 (my-app-feedback@example.com)
```

Подробнее про [ошибки в заголовке User-Agent](https://github.com/hhru/api/blob/master/docs/errors.md#user-agent).


<a name=\"request-body\"></a>
## Формат тела запроса при отправке JSON

Данные, передающиеся в теле запроса, должны удовлетворять требованиям:

* Валидный JSON (допускается передача как минифицированного варианта, так и
pretty print варианта с дополнительными пробелами и сбросами строк).
* Рекомендуется использование кодировки UTF-8 без дополнительного экранирования
(`{\"name\": \"Иванов Иван\"}`).
* Также возможно использовать ascii кодировку с экранированием
(`{\"name\": \"\\u0418\\u0432\\u0430\\u043d\\u043e\\u0432 \\u0418\\u0432\\u0430\\u043d\"}`).
* К типам данных в определённым полях накладываются дополнительные условия,
описанные в каждом конкретном методе. В JSON типами данных являются `string`,
`number`, `boolean`, `null`, `object`, `array`.

### Ответ
Ответ свыше определенной длины будет сжиматься методом gzip.

### Ошибки и коды ответов

API широко использует информирование при помощи кодов ответов.
Приложение должно корректно их обрабатывать.

В случае неполадок и сбоев, возможны ответы с кодом `503` и `500`.

При каждой ошибке, помимо кода ответа, в теле ответа может быть выдана
дополнительная информация, позволяющая разработчику понять
причину соответствующего ответа.

[Более подробно про возможные ошибки](https://github.com/hhru/api/blob/master/docs/errors.md).


## Недокументированные поля и параметры запросов

В ответах и параметрах API можно найти ключи, не описанные в документации.
Обычно это означает, что они оставлены для совместимости со старыми версиями.
Их использование не рекомендуется. Если ваше приложение использует такие ключи,
перейдите на использование актуальных ключей, описанных в документации.


## Пагинация

К любому запросу, подразумевающему выдачу списка объектов, можно в параметрах
указать `page=N&per_page=M`. Нумерация идёт с нуля, по умолчанию выдаётся
первая (нулевая) страница с 20 объектами на странице. Во всех ответах, где
доступна пагинация, единообразный корневой объект:

```json
{
  \"found\": 1,
  \"per_page\": 1,
  \"pages\": 1,
  \"page\": 0,
  \"items\": [{}]
}
```
## Выбор сайта

API HeadHunter позволяет получать данные со всех сайтов группы компании
HeadHunter.

В частности:

* hh.ru
* rabota.by
* hh1.az
* hh.uz
* hh.kz
* headhunter.ge
* headhunter.kg

Запросы к данным на всех сайтах следует направлять на `https://api.hh.ru/`.

При необходимости учесть специфику сайта, можно добавить в запрос параметр
`?host=`. По умолчанию используется `hh.ru`.

Например, для получения [локализаций](https://api.hh.ru/openapi/redoc#tag/Obshie-spravochniki/operation/get-locales), доступных на hh.kz необходимо
сделать GET запрос на `https://api.hh.ru/locales?host=hh.kz`.

## CORS (Cross-Origin Resource Sharing)

API поддерживает технологию CORS для запроса данных из
браузера с произвольного домена. Этот метод более предпочтителен, чем
использование JSONP. Он не ограничен методом GET. Для отладки CORS доступен
[специальный метод](https://github.com/hhru/api/blob/master/docs/cors.md). Для использования JSONP передайте параметр
`?callback=callback_name`.

* [CORS specification on w3.org](http://www.w3.org/TR/cors/)
* [HTML5Rocks CORS Tutorial](http://www.html5rocks.com/en/tutorials/cors/)
* [CORS on dev.opera.com](http://dev.opera.com/articles/view/dom-access-control-using-cross-origin-resource-sharing/)
* [CORS on caniuse.com](http://caniuse.com/#feat=cors)
* [CORS on en.wikipedia.org](http://en.wikipedia.org/wiki/Cross-origin_resource_sharing)


## Внешние ссылки на статьи и стандарты

* [HTTP/1.1](http://tools.ietf.org/html/rfc2616)
* [JSON](http://json.org/)
* [URI Template](http://tools.ietf.org/html/rfc6570)
* [OAuth 2.0](http://tools.ietf.org/html/rfc6749)
* [REST](http://www.ics.uci.edu/~fielding/pubs/dissertation/rest_arch_style.htm)
* [ISO 8601](http://en.wikipedia.org/wiki/ISO_8601)

# Авторизация

API поддерживает следующие уровни авторизации:
  - [авторизация приложения](#tag/Avtorizaciya-prilozheniya)
  - [авторизация пользователя](#section/Avtorizaciya/Avtorizaciya-polzovatelya)

* [Авторизация пользователя](#section/Avtorizaciya/Avtorizaciya-polzovatelya)
  * [Правила формирования специального redirect_uri](#section/Avtorizaciya/Pravila-formirovaniya-specialnogo-redirect_uri)
  * [Процесс авторизации](#section/Avtorizaciya/Process-avtorizacii)
  * [Успешное получение временного `authorization_code`](#get-authorization_code)
  * [Получение access и refresh токенов](#section/Avtorizaciya/Poluchenie-access-i-refresh-tokenov)
* [Обновление пары access и refresh токенов](#section/Avtorizaciya/Obnovlenie-pary-access-i-refresh-tokenov)
* [Инвалидация токена](#tag/Avtorizaciya-rabotodatelya/operation/invalidate-token)
* [Запрос авторизации под другим пользователем](#section/Avtorizaciya/Zapros-avtorizacii-pod-drugim-polzovatelem)
* [Авторизация под разными рабочими аккаунтами](#section/Avtorizaciya/Avtorizaciya-pod-raznymi-rabochimi-akkauntami)
* [Авторизация приложения](#tag/Avtorizaciya-prilozheniya)    


## Авторизация пользователя
Для выполнения запросов от имени пользователя необходимо пользоваться токеном пользователя.

В начале приложению необходимо направить пользователя (открыть страницу) по адресу:

```
https://hh.ru/oauth/authorize?
response_type=code&
client_id={client_id}&
state={state}&
redirect_uri={redirect_uri}
```

Обязательные параметры:

* `response_type=code` — указание на способ получения авторизации, используя `authorization code`
* `client_id` — идентификатор, полученный при создании приложения


Необязательные параметры:

* `state` — в случае указания, будет включен в ответный редирект.
Это позволяет исключить возможность взлома путём подделки межсайтовых
запросов. Подробнее об этом:
[RFC 6749. Section 10.12](http://tools.ietf.org/html/rfc6749#section-10.12)
* `redirect_uri` — uri для перенаправления пользователя после
авторизации. Если не указать, используется из настроек приложения. При
наличии происходит валидация значения. Вероятнее всего, потребуется сделать
urlencode значения параметра.

## Правила формирования специального redirect_uri

К примеру, если в настройках сохранен `http://example.com/oauth`, то разрешено
указывать:

* `http://www.example.com/oauth` — поддомен;
* `http://www.example.com/oauth/sub/path` — уточнение пути;
* `http://example.com/oauth?lang=RU` — дополнительный параметр;
* `http://www.example.com/oauth/sub/path?lang=RU` — всё вместе.

Запрещено:

* `https://example.com/oauth` — различные протоколы;
* `http://wwwexample.com/oauth` — различные домены;
* `http://wwwexample.com/` — другой путь;
* `http://example.com/oauths` — другой путь;
* `http://example.com:80/oauths` — указание изначально отсутствующего порта;

## Процесс авторизации

Если пользователь не авторизован на сайте, ему будет показана форма
авторизации на сайте. После прохождения авторизации на сайте, пользователю
будет выведена форма с запросом разрешения доступа вашего приложения к его
персональным данным.

Если пользователь не разрешает доступ приложению, пользователь будет
перенаправлен на указанный `redirect_uri` с `?error=access_denied` и
`state={state}`, если таковой был указан при первом запросе.

<a name=\"get-authorization_code\"></a>
### Успешное получение временного `authorization_code`

В случае разрешения прав, в редиректе будет указан
временный `authorization_code`:

```http
HTTP/1.1 302 FOUND
Location: {redirect_uri}?code={authorization_code}
```

Если пользователь авторизован на сайте и доступ данному приложению однажды
ранее выдан, ответом будет сразу вышеописанный редирект с `authorization_code`
(без показа формы логина и выдачи прав).

## Получение access и refresh токенов

После получения `authorization_code` приложению необходимо сделать сервер-сервер запрос `POST https://api.hh.ru/token` для обмена полученного
`authorization_code` на `access_token` (старый запрос `POST https://hh.ru/oauth/token` считается устаревшим).

В теле запроса необходимо передать [дополнительные параметры](#required_parameters).

Тело запроса необходимо передавать в стандартном
`application/x-www-form-urlencoded` с указанием соответствующего заголовка
`Content-Type`.

`authorization_code` имеет довольно короткий срок жизни, при его истечении
необходимо запросить новый.

## Обновление пары access и refresh токенов
`access_token` также имеет срок жизни (ключ `expires_in`, в секундах), при его
истечении приложение должно сделать запрос с `refresh_token` для получения
нового.

Запрос необходимо делать в `application/x-www-form-urlencoded`.

```
POST https://api.hh.ru/token
```

(старый запрос `POST https://hh.ru/oauth/token` считается устаревшим)

В теле запроса необходимо передать [дополнительные параметры](#required_parameters)

`refresh_token` можно использовать только один раз и только по истечению
срока действия `access_token`.

После получения новой пары access и refresh токенов, их необходимо использовать
в дальнейших запросах в api и запросах на продление токена.

## Запрос авторизации под другим пользователем

Возможен следующий сценарий:

1. Приложение перенаправляет пользователя на сайт с запросом авторизации.
2. Пользователь на сайте уже авторизован и данному приложение доступ уже был
разрешен.
3. Пользователю будет предложена возможность продолжить работу под текущим аккаунтом,
либо зайти под другим аккаунтом.

Если есть необходимость, чтобы на шаге 3 сразу происходило перенаправление (redirect) с временным токеном,
необходимо добавить к запросу `/oauth/authorize...` параметр `skip_choose_account=true`.
В этом случае автоматически выдаётся доступ пользователю авторизованному на сайте.

Если есть необходимость всегда показывать форму авторизации, приложение может
добавить к запросу `/oauth/authorize...` параметр `force_login=true`. В этом
случае, пользователю будет показана форма авторизации с логином и паролем
даже в случае, если пользователь уже авторизован.

Это может быть полезно приложениям, которые предоставляют сервис только для
соискателей. Если пришел пользователь-работодатель, приложение может предложить
пользователю повторно разрешить доступ на сайте, уже указав
другую учетную запись.

Также, после авторизации приложение может показать пользователю сообщение:

```
Вы вошли как %Имя_Фамилия%. Это не вы?
```
и предоставить ссылку с `force_login=true` для возможности захода под
другим логином.

## Авторизация под разными рабочими аккаунтами

Для получения списка рабочих аккаунтов менеджера и для работы под разными рабочими аккаунтами менеджера необходимо прочитать документацию по [рабочим аккаунтам менеджера](#tag/Menedzhery-rabotodatelya/operation/get-manager-accounts)

В случае компрометации токена необходимо [инвалидировать](#tag/Avtorizaciya-rabotodatelya/operation/invalidate-token) скомпрометированный токен и запросить токен заново!

<!-- ReDoc-Inject: <security-definitions> -->


## Overview
This API client was generated by the [OpenAPI Generator](https://openapi-generator.tech) project.  By using the [OpenAPI-spec](https://www.openapis.org/) from a remote server, you can easily generate an API client.

- API version: 1.0.0
- Package version: 1.0.0
- Generator version: 7.6.0-SNAPSHOT
- Build package: org.openapitools.codegen.languages.GoClientCodegen
For more information, please visit [https://api.hh.ru/openapi/redoc](https://api.hh.ru/openapi/redoc)

## Installation

Install the following dependencies:

```sh
go get github.com/stretchr/testify/assert
go get golang.org/x/net/context
```

Put the package under your project folder and add the following in import:

```go
import hh "github.com/zaboal/hh-go"
```

To use a proxy, set the environment variable `HTTP_PROXY`:

```go
os.Setenv("HTTP_PROXY", "http://proxy_name:proxy_port")
```

## Configuration of Server URL

Default configuration comes with `Servers` field that contains server objects as defined in the OpenAPI specification.

### Select Server Configuration

For using other server than the one defined on index 0 set context value `hh.ContextServerIndex` of type `int`.

```go
ctx := context.WithValue(context.Background(), hh.ContextServerIndex, 1)
```

### Templated Server URL

Templated server URL is formatted using default variables from configuration or from context value `hh.ContextServerVariables` of type `map[string]string`.

```go
ctx := context.WithValue(context.Background(), hh.ContextServerVariables, map[string]string{
	"basePath": "v2",
})
```

Note, enum values are always validated and all unused variables are silently ignored.

### URLs Configuration per Operation

Each operation can use different server URL defined using `OperationServers` map in the `Configuration`.
An operation is uniquely identified by `"{classname}Service.{nickname}"` string.
Similar rules for overriding default operation server index and variables applies by using `hh.ContextOperationServerIndices` and `hh.ContextOperationServerVariables` context maps.

```go
ctx := context.WithValue(context.Background(), hh.ContextOperationServerIndices, map[string]int{
	"{classname}Service.{nickname}": 2,
})
ctx = context.WithValue(context.Background(), hh.ContextOperationServerVariables, map[string]map[string]string{
	"{classname}Service.{nickname}": {
		"port": "8443",
	},
})
```

## Documentation for API Endpoints

All URIs are relative to *https://api.hh.ru*

Class | Method | HTTP request | Description
------------ | ------------- | ------------- | -------------
*DefaultApi* | [**AddApplicantComment**](docs/DefaultApi.md#addapplicantcomment) | **Post** /applicant_comments/{applicant_id} | Добавление комментария
*DefaultApi* | [**AddApplicantComment_0**](docs/DefaultApi.md#addapplicantcomment_0) | **Post** /applicant_comments/{applicant_id} | Добавление комментария
*DefaultApi* | [**AddEmployerManager**](docs/DefaultApi.md#addemployermanager) | **Post** /employers/{employer_id}/managers | Добавление менеджера
*DefaultApi* | [**AddEmployerToBlacklisted**](docs/DefaultApi.md#addemployertoblacklisted) | **Put** /employers/blacklisted/{employer_id} | Добавление работодателя в список скрытых
*DefaultApi* | [**AddResumeVisibilityList**](docs/DefaultApi.md#addresumevisibilitylist) | **Post** /resumes/{resume_id}/{list_type} | Добавление работодателей в список видимости
*DefaultApi* | [**AddVacancyToArchive**](docs/DefaultApi.md#addvacancytoarchive) | **Put** /employers/{employer_id}/vacancies/archived/{vacancy_id} | Архивация вакансии
*DefaultApi* | [**AddVacancyToArchive_0**](docs/DefaultApi.md#addvacancytoarchive_0) | **Put** /employers/{employer_id}/vacancies/archived/{vacancy_id} | Архивация вакансии
*DefaultApi* | [**AddVacancyToBlacklisted**](docs/DefaultApi.md#addvacancytoblacklisted) | **Put** /vacancies/blacklisted/{vacancy_id} | Добавление вакансии в список скрытых
*DefaultApi* | [**AddVacancyToFavorite**](docs/DefaultApi.md#addvacancytofavorite) | **Put** /vacancies/favorited/{vacancy_id} | Добавление вакансии в список отобранных
*DefaultApi* | [**AddVacancyToHidden**](docs/DefaultApi.md#addvacancytohidden) | **Put** /employers/{employer_id}/vacancies/hidden/{vacancy_id} | Удаление вакансий
*DefaultApi* | [**ApplyToVacancy**](docs/DefaultApi.md#applytovacancy) | **Post** /negotiations | Отклик на вакансию
*DefaultApi* | [**Authorize**](docs/DefaultApi.md#authorize) | **Post** /oauth/token | Получение access-токена
*DefaultApi* | [**Authorize_0**](docs/DefaultApi.md#authorize_0) | **Post** /oauth/token | Получение access-токена
*DefaultApi* | [**Authorize_1**](docs/DefaultApi.md#authorize_1) | **Post** /oauth/token | Получение access-токена
*DefaultApi* | [**ChangeNegotiationAction**](docs/DefaultApi.md#changenegotiationaction) | **Put** /negotiations/{collection_name}/{nid} | Действия по отклику/приглашению коллекции
*DefaultApi* | [**ChangeNegotiationAction_0**](docs/DefaultApi.md#changenegotiationaction_0) | **Put** /negotiations/{collection_name}/{nid} | Действия по отклику/приглашению коллекции
*DefaultApi* | [**ChangeVacancyDraft**](docs/DefaultApi.md#changevacancydraft) | **Put** /vacancies/drafts/{draft_id} | Изменение черновика вакансии
*DefaultApi* | [**ConfirmPhoneInResume**](docs/DefaultApi.md#confirmphoneinresume) | **Post** /resume_phone_confirm | Подтвердить телефон кодом
*DefaultApi* | [**CreateResume**](docs/DefaultApi.md#createresume) | **Post** /resumes | Создание резюме
*DefaultApi* | [**CreateSavedResumeSearch**](docs/DefaultApi.md#createsavedresumesearch) | **Post** /saved_searches/resumes | Создание нового сохраненного поиска резюме
*DefaultApi* | [**CreateSavedVacancySearch**](docs/DefaultApi.md#createsavedvacancysearch) | **Post** /saved_searches/vacancies | Создание нового сохраненного поиска вакансий
*DefaultApi* | [**CreateVacancyDraft**](docs/DefaultApi.md#createvacancydraft) | **Post** /vacancies/drafts | Создание черновика вакансии
*DefaultApi* | [**DeleteApplicantComment**](docs/DefaultApi.md#deleteapplicantcomment) | **Delete** /applicant_comments/{applicant_id}/{comment_id} | Удаление комментария
*DefaultApi* | [**DeleteApplicantComment_0**](docs/DefaultApi.md#deleteapplicantcomment_0) | **Delete** /applicant_comments/{applicant_id}/{comment_id} | Удаление комментария
*DefaultApi* | [**DeleteArtifact**](docs/DefaultApi.md#deleteartifact) | **Delete** /artifacts/{id} | Удаление артефакта
*DefaultApi* | [**DeleteEmployerFromBlacklisted**](docs/DefaultApi.md#deleteemployerfromblacklisted) | **Delete** /employers/blacklisted/{employer_id} | Удаление работодателя из списка скрытых
*DefaultApi* | [**DeleteEmployerFromResumeVisibilityList**](docs/DefaultApi.md#deleteemployerfromresumevisibilitylist) | **Delete** /resumes/{resume_id}/{list_type}/employer | Удаление работодателя из списка видимости
*DefaultApi* | [**DeleteEmployerManager**](docs/DefaultApi.md#deleteemployermanager) | **Delete** /employers/{employer_id}/managers/{manager_id} | Удаление менеджера
*DefaultApi* | [**DeleteResume**](docs/DefaultApi.md#deleteresume) | **Delete** /resumes/{resume_id} | Удаление резюме
*DefaultApi* | [**DeleteResumeVisibilityList**](docs/DefaultApi.md#deleteresumevisibilitylist) | **Delete** /resumes/{resume_id}/{list_type} | Очистка списка видимости
*DefaultApi* | [**DeleteSavedResumeSearch**](docs/DefaultApi.md#deletesavedresumesearch) | **Delete** /saved_searches/resumes/{id} | Удаление сохраненного поиска резюме
*DefaultApi* | [**DeleteSavedVacancySearch**](docs/DefaultApi.md#deletesavedvacancysearch) | **Delete** /saved_searches/vacancies/{id} | Удаление сохраненного поиска вакансий
*DefaultApi* | [**DeleteVacancyDraft**](docs/DefaultApi.md#deletevacancydraft) | **Delete** /vacancies/drafts/{draft_id} | Удаление черновика вакансии
*DefaultApi* | [**DeleteVacancyFromBlacklisted**](docs/DefaultApi.md#deletevacancyfromblacklisted) | **Delete** /vacancies/blacklisted/{vacancy_id} | Удаление вакансии из списка скрытых
*DefaultApi* | [**DeleteVacancyFromFavorite**](docs/DefaultApi.md#deletevacancyfromfavorite) | **Delete** /vacancies/favorited/{vacancy_id} | Удаление вакансии из списка отобранных
*DefaultApi* | [**DisableAutomaticVacancyPublication**](docs/DefaultApi.md#disableautomaticvacancypublication) | **Delete** /vacancies/auto_publication | Отмена автопубликации вакансии
*DefaultApi* | [**EditArtifact**](docs/DefaultApi.md#editartifact) | **Put** /artifacts/{id} | Редактирование артефакта
*DefaultApi* | [**EditCurrentUserInfo**](docs/DefaultApi.md#editcurrentuserinfo) | **Post** /me | Редактирование информации авторизованного пользователя
*DefaultApi* | [**EditEmployerManager**](docs/DefaultApi.md#editemployermanager) | **Put** /employers/{employer_id}/managers/{manager_id} | Редактирование менеджера
*DefaultApi* | [**EditNegotiationMessage**](docs/DefaultApi.md#editnegotiationmessage) | **Put** /negotiations/{nid}/messages/{mid} | Редактирование сообщения в отклике
*DefaultApi* | [**EditResume**](docs/DefaultApi.md#editresume) | **Put** /resumes/{resume_id} | Обновление резюме
*DefaultApi* | [**EditVacancy**](docs/DefaultApi.md#editvacancy) | **Put** /vacancies/{vacancy_id} | Редактирование вакансий
*DefaultApi* | [**EditVacancy_0**](docs/DefaultApi.md#editvacancy_0) | **Put** /vacancies/{vacancy_id} | Редактирование вакансий
*DefaultApi* | [**GetActiveNegotiations**](docs/DefaultApi.md#getactivenegotiations) | **Get** /negotiations/active | Список активных откликов
*DefaultApi* | [**GetActiveVacancyList**](docs/DefaultApi.md#getactivevacancylist) | **Get** /employers/{employer_id}/vacancies/active | Просмотр списка опубликованных вакансий
*DefaultApi* | [**GetActiveVacancyList_0**](docs/DefaultApi.md#getactivevacancylist_0) | **Get** /employers/{employer_id}/vacancies/active | Просмотр списка опубликованных вакансий
*DefaultApi* | [**GetAddress**](docs/DefaultApi.md#getaddress) | **Get** /employers/{employer_id}/addresses/{address_id} | Получение адреса
*DefaultApi* | [**GetAddress_0**](docs/DefaultApi.md#getaddress_0) | **Get** /employers/{employer_id}/addresses/{address_id} | Получение адреса
*DefaultApi* | [**GetAllDistricts**](docs/DefaultApi.md#getalldistricts) | **Get** /districts | Список районов во всех городах
*DefaultApi* | [**GetApplicantCommentsList**](docs/DefaultApi.md#getapplicantcommentslist) | **Get** /applicant_comments/{applicant_id} | Получение списка комментариев
*DefaultApi* | [**GetApplicantCommentsList_0**](docs/DefaultApi.md#getapplicantcommentslist_0) | **Get** /applicant_comments/{applicant_id} | Получение списка комментариев
*DefaultApi* | [**GetApplicantPhoneInfo**](docs/DefaultApi.md#getapplicantphoneinfo) | **Get** /resume_should_send_sms | Получить информацию о телефоне соискателя
*DefaultApi* | [**GetArchivedVacancies**](docs/DefaultApi.md#getarchivedvacancies) | **Get** /employers/{employer_id}/vacancies/archived | Список архивных вакансий
*DefaultApi* | [**GetArchivedVacancies_0**](docs/DefaultApi.md#getarchivedvacancies_0) | **Get** /employers/{employer_id}/vacancies/archived | Список архивных вакансий
*DefaultApi* | [**GetAreaLeavesSuggests**](docs/DefaultApi.md#getarealeavessuggests) | **Get** /suggests/area_leaves | Подсказки по регионам, являющимися листами в дереве регионов
*DefaultApi* | [**GetAreas**](docs/DefaultApi.md#getareas) | **Get** /areas | Дерево всех регионов
*DefaultApi* | [**GetAreasFromSpecified**](docs/DefaultApi.md#getareasfromspecified) | **Get** /areas/{area_id} | Справочник регионов, начиная с указанного
*DefaultApi* | [**GetAreasSuggests**](docs/DefaultApi.md#getareassuggests) | **Get** /suggests/areas | Подсказки по регионам
*DefaultApi* | [**GetArtifactPhotos**](docs/DefaultApi.md#getartifactphotos) | **Get** /artifacts/photo | Получение фотографий
*DefaultApi* | [**GetArtifactPhotosConditions**](docs/DefaultApi.md#getartifactphotosconditions) | **Get** /artifacts/photo/conditions | Условия загрузки фотографий
*DefaultApi* | [**GetArtifactsPortfolio**](docs/DefaultApi.md#getartifactsportfolio) | **Get** /artifacts/portfolio | Получение портфолио
*DefaultApi* | [**GetArtifactsPortfolioConditions**](docs/DefaultApi.md#getartifactsportfolioconditions) | **Get** /artifacts/portfolio/conditions | Условия загрузки портфолио
*DefaultApi* | [**GetAvailableVacancyTypes**](docs/DefaultApi.md#getavailablevacancytypes) | **Get** /employers/{employer_id}/managers/{manager_id}/vacancies/available_types | Варианты публикации вакансий у текущего менеджера
*DefaultApi* | [**GetBlacklistedEmployers**](docs/DefaultApi.md#getblacklistedemployers) | **Get** /employers/blacklisted | Список скрытых работодателей
*DefaultApi* | [**GetBlacklistedVacancies**](docs/DefaultApi.md#getblacklistedvacancies) | **Get** /vacancies/blacklisted | Список скрытых вакансий
*DefaultApi* | [**GetCollectionNegotiationsList**](docs/DefaultApi.md#getcollectionnegotiationslist) | **Get** /negotiations/response | Список откликов/приглашений коллекции
*DefaultApi* | [**GetCollectionNegotiationsList_0**](docs/DefaultApi.md#getcollectionnegotiationslist_0) | **Get** /negotiations/response | Список откликов/приглашений коллекции
*DefaultApi* | [**GetCountries**](docs/DefaultApi.md#getcountries) | **Get** /areas/countries | Справочник стран
*DefaultApi* | [**GetCurrentUserInfo**](docs/DefaultApi.md#getcurrentuserinfo) | **Get** /me | Информация о текущем пользователе
*DefaultApi* | [**GetCurrentUserInfo_0**](docs/DefaultApi.md#getcurrentuserinfo_0) | **Get** /me | Информация о текущем пользователе
*DefaultApi* | [**GetCurrentUserInfo_1**](docs/DefaultApi.md#getcurrentuserinfo_1) | **Get** /me | Информация о текущем пользователе
*DefaultApi* | [**GetCurrentUserInfo_2**](docs/DefaultApi.md#getcurrentuserinfo_2) | **Get** /me | Информация о текущем пользователе
*DefaultApi* | [**GetCurrentUserInfo_3**](docs/DefaultApi.md#getcurrentuserinfo_3) | **Get** /me | Информация о текущем пользователе
*DefaultApi* | [**GetDictionaries**](docs/DefaultApi.md#getdictionaries) | **Get** /dictionaries | Справочники полей
*DefaultApi* | [**GetEducationalInstitutionsDictionary**](docs/DefaultApi.md#geteducationalinstitutionsdictionary) | **Get** /educational_institutions | Основная информация об учебных заведениях
*DefaultApi* | [**GetEducationalInstitutionsSuggests**](docs/DefaultApi.md#geteducationalinstitutionssuggests) | **Get** /suggests/educational_institutions | Подсказки по названиям учебных заведений
*DefaultApi* | [**GetEmployerAddresses**](docs/DefaultApi.md#getemployeraddresses) | **Get** /employers/{employer_id}/addresses | Список адресов работодателя
*DefaultApi* | [**GetEmployerAddresses_0**](docs/DefaultApi.md#getemployeraddresses_0) | **Get** /employers/{employer_id}/addresses | Список адресов работодателя
*DefaultApi* | [**GetEmployerDepartments**](docs/DefaultApi.md#getemployerdepartments) | **Get** /employers/{employer_id}/departments | Справочник департаментов работодателя
*DefaultApi* | [**GetEmployerDepartments_0**](docs/DefaultApi.md#getemployerdepartments_0) | **Get** /employers/{employer_id}/departments | Справочник департаментов работодателя
*DefaultApi* | [**GetEmployerInfo**](docs/DefaultApi.md#getemployerinfo) | **Get** /employers/{employer_id} | Информация о работодателе
*DefaultApi* | [**GetEmployerManager**](docs/DefaultApi.md#getemployermanager) | **Get** /employers/{employer_id}/managers/{manager_id} | Получение информации о менеджере
*DefaultApi* | [**GetEmployerManagerLimits**](docs/DefaultApi.md#getemployermanagerlimits) | **Get** /employers/{employer_id}/managers/{manager_id}/limits/resume | Дневной лимит просмотра резюме для текущего менеджера
*DefaultApi* | [**GetEmployerManagerTypes**](docs/DefaultApi.md#getemployermanagertypes) | **Get** /employers/{employer_id}/manager_types | Справочник типов и прав менеджера
*DefaultApi* | [**GetEmployerManagerTypes_0**](docs/DefaultApi.md#getemployermanagertypes_0) | **Get** /employers/{employer_id}/manager_types | Справочник типов и прав менеджера
*DefaultApi* | [**GetEmployerManager_0**](docs/DefaultApi.md#getemployermanager_0) | **Get** /employers/{employer_id}/managers/{manager_id} | Получение информации о менеджере
*DefaultApi* | [**GetEmployerManagers**](docs/DefaultApi.md#getemployermanagers) | **Get** /employers/{employer_id}/managers | Список менеджеров работодателя
*DefaultApi* | [**GetEmployerManagers_0**](docs/DefaultApi.md#getemployermanagers_0) | **Get** /employers/{employer_id}/managers | Список менеджеров работодателя
*DefaultApi* | [**GetEmployerVacancyAreas**](docs/DefaultApi.md#getemployervacancyareas) | **Get** /employers/{employer_id}/vacancy_areas/active | Список регионов, в которых есть активные вакансии
*DefaultApi* | [**GetFaculties**](docs/DefaultApi.md#getfaculties) | **Get** /educational_institutions/{id}/faculties | Список факультетов учебного заведения
*DefaultApi* | [**GetFavoriteVacancies**](docs/DefaultApi.md#getfavoritevacancies) | **Get** /vacancies/favorited | Список отобранных вакансий
*DefaultApi* | [**GetFieldsOfStudySuggestions**](docs/DefaultApi.md#getfieldsofstudysuggestions) | **Get** /suggests/fields_of_study | Подсказки по специализациям
*DefaultApi* | [**GetHiddenVacancies**](docs/DefaultApi.md#gethiddenvacancies) | **Get** /employers/{employer_id}/vacancies/hidden | Список удаленных вакансий
*DefaultApi* | [**GetHiddenVacancies_0**](docs/DefaultApi.md#gethiddenvacancies_0) | **Get** /employers/{employer_id}/vacancies/hidden | Список удаленных вакансий
*DefaultApi* | [**GetIndustries**](docs/DefaultApi.md#getindustries) | **Get** /industries | Отрасли компаний
*DefaultApi* | [**GetLanguages**](docs/DefaultApi.md#getlanguages) | **Get** /languages | Список всех языков
*DefaultApi* | [**GetLocales**](docs/DefaultApi.md#getlocales) | **Get** /locales | Список доступных локалей
*DefaultApi* | [**GetLocalesForResume**](docs/DefaultApi.md#getlocalesforresume) | **Get** /locales/resume | Список доступных локалей для резюме
*DefaultApi* | [**GetMailTemplates**](docs/DefaultApi.md#getmailtemplates) | **Get** /employers/{employer_id}/mail_templates | Список доступных шаблонов ответов соискателю
*DefaultApi* | [**GetMailTemplates_0**](docs/DefaultApi.md#getmailtemplates_0) | **Get** /employers/{employer_id}/mail_templates | Список доступных шаблонов ответов соискателю
*DefaultApi* | [**GetManagerAccounts**](docs/DefaultApi.md#getmanageraccounts) | **Get** /manager_accounts/mine | Рабочие аккаунты менеджера
*DefaultApi* | [**GetManagerSettings**](docs/DefaultApi.md#getmanagersettings) | **Get** /employers/{employer_id}/managers/{manager_id}/settings | Предпочтения менеджера
*DefaultApi* | [**GetMetroStations**](docs/DefaultApi.md#getmetrostations) | **Get** /metro | Список станций метро во всех городах
*DefaultApi* | [**GetMetroStationsInCity**](docs/DefaultApi.md#getmetrostationsincity) | **Get** /metro/{city_id} | Список станций метро в указанном городе
*DefaultApi* | [**GetMineResumes**](docs/DefaultApi.md#getmineresumes) | **Get** /resumes/mine | Список резюме авторизованного пользователя
*DefaultApi* | [**GetNegotiationItem**](docs/DefaultApi.md#getnegotiationitem) | **Get** /negotiations/{nid} | Просмотр отклика/приглашения
*DefaultApi* | [**GetNegotiationItem_0**](docs/DefaultApi.md#getnegotiationitem_0) | **Get** /negotiations/{nid} | Просмотр отклика/приглашения
*DefaultApi* | [**GetNegotiationMessageTemplates**](docs/DefaultApi.md#getnegotiationmessagetemplates) | **Get** /message_templates/{template} | Список шаблонов ответов для отклика/приглашения
*DefaultApi* | [**GetNegotiationMessages**](docs/DefaultApi.md#getnegotiationmessages) | **Get** /negotiations/{nid}/messages | Просмотр списка сообщений в отклике
*DefaultApi* | [**GetNegotiationTestResults**](docs/DefaultApi.md#getnegotiationtestresults) | **Get** /negotiations/{nid}/test/solution | Получить результаты тестов, прикрепленных к вакансии
*DefaultApi* | [**GetNegotiationTestResults_0**](docs/DefaultApi.md#getnegotiationtestresults_0) | **Get** /negotiations/{nid}/test/solution | Получить результаты тестов, прикрепленных к вакансии
*DefaultApi* | [**GetNegotiations**](docs/DefaultApi.md#getnegotiations) | **Get** /negotiations | Список откликов/приглашений
*DefaultApi* | [**GetNegotiationsStatisticsEmployer**](docs/DefaultApi.md#getnegotiationsstatisticsemployer) | **Get** /employers/{employer_id}/negotiations_statistics | Статистика откликов для компании
*DefaultApi* | [**GetNegotiationsStatisticsEmployer_0**](docs/DefaultApi.md#getnegotiationsstatisticsemployer_0) | **Get** /employers/{employer_id}/negotiations_statistics | Статистика откликов для компании
*DefaultApi* | [**GetNegotiationsStatisticsManager**](docs/DefaultApi.md#getnegotiationsstatisticsmanager) | **Get** /employers/{employer_id}/managers/{manager_id}/negotiations_statistics | Статистика откликов для менеджера
*DefaultApi* | [**GetNegotiationsStatisticsManager_0**](docs/DefaultApi.md#getnegotiationsstatisticsmanager_0) | **Get** /employers/{employer_id}/managers/{manager_id}/negotiations_statistics | Статистика откликов для менеджера
*DefaultApi* | [**GetNegotiations_0**](docs/DefaultApi.md#getnegotiations_0) | **Get** /negotiations | Список откликов/приглашений
*DefaultApi* | [**GetNegotiations_1**](docs/DefaultApi.md#getnegotiations_1) | **Get** /negotiations | Список откликов/приглашений
*DefaultApi* | [**GetNewResumeConditions**](docs/DefaultApi.md#getnewresumeconditions) | **Get** /resume_conditions | Условия заполнения полей нового резюме
*DefaultApi* | [**GetPayableApiActions**](docs/DefaultApi.md#getpayableapiactions) | **Get** /employers/{employer_id}/services/payable_api_actions/active | Информация по активным услугам API для платных методов
*DefaultApi* | [**GetPayableApiMethodAccess**](docs/DefaultApi.md#getpayableapimethodaccess) | **Get** /employers/{employer_id}/managers/{manager_id}/method_access | Проверка доступа к платным методам
*DefaultApi* | [**GetPositionsSuggestions**](docs/DefaultApi.md#getpositionssuggestions) | **Get** /suggests/positions | Подсказки по должностям резюме
*DefaultApi* | [**GetPrefNegotiationsOrder**](docs/DefaultApi.md#getprefnegotiationsorder) | **Get** /vacancies/{id}/preferred_negotiations_order | Просмотр предпочитаемой сортировки откликов
*DefaultApi* | [**GetPrefNegotiationsOrder_0**](docs/DefaultApi.md#getprefnegotiationsorder_0) | **Get** /vacancies/{id}/preferred_negotiations_order | Просмотр предпочитаемой сортировки откликов
*DefaultApi* | [**GetProfessionalRolesDictionary**](docs/DefaultApi.md#getprofessionalrolesdictionary) | **Get** /professional_roles | Справочник профессиональных ролей
*DefaultApi* | [**GetProfessionalRolesSuggests**](docs/DefaultApi.md#getprofessionalrolessuggests) | **Get** /suggests/professional_roles | Подсказки по профессиональным ролям
*DefaultApi* | [**GetProlongationVacancyInfo**](docs/DefaultApi.md#getprolongationvacancyinfo) | **Get** /vacancies/{vacancy_id}/prolongate | Информация о возможности продления вакансии
*DefaultApi* | [**GetProlongationVacancyInfo_0**](docs/DefaultApi.md#getprolongationvacancyinfo_0) | **Get** /vacancies/{vacancy_id}/prolongate | Информация о возможности продления вакансии
*DefaultApi* | [**GetRegisteredCompaniesSuggests**](docs/DefaultApi.md#getregisteredcompaniessuggests) | **Get** /suggests/companies | Подсказки по зарегистрированным организациям
*DefaultApi* | [**GetResume**](docs/DefaultApi.md#getresume) | **Get** /resumes/{resume_id} | Просмотр резюме
*DefaultApi* | [**GetResumeAccessTypes**](docs/DefaultApi.md#getresumeaccesstypes) | **Get** /resumes/{resume_id}/access_types | Получение списка типов видимости резюме
*DefaultApi* | [**GetResumeConditions**](docs/DefaultApi.md#getresumeconditions) | **Get** /resumes/{resume_id}/conditions | Условия заполнения полей существующего резюме
*DefaultApi* | [**GetResumeCreationAvailability**](docs/DefaultApi.md#getresumecreationavailability) | **Get** /resumes/creation_availability | Проверка возможности создания резюме
*DefaultApi* | [**GetResumeNegotiationsHistory**](docs/DefaultApi.md#getresumenegotiationshistory) | **Get** /resumes/{resume_id}/negotiations_history | История откликов/приглашений по резюме
*DefaultApi* | [**GetResumeNegotiationsHistory_0**](docs/DefaultApi.md#getresumenegotiationshistory_0) | **Get** /resumes/{resume_id}/negotiations_history | История откликов/приглашений по резюме
*DefaultApi* | [**GetResumeSearchKeywordsSuggests**](docs/DefaultApi.md#getresumesearchkeywordssuggests) | **Get** /suggests/resume_search_keyword | Подсказки по ключевым словам поиска резюме
*DefaultApi* | [**GetResumeStatus**](docs/DefaultApi.md#getresumestatus) | **Get** /resumes/{resume_id}/status | Статус резюме и готовность к публикации
*DefaultApi* | [**GetResumeViewHistory**](docs/DefaultApi.md#getresumeviewhistory) | **Get** /resumes/{resume_id}/views | История просмотра резюме
*DefaultApi* | [**GetResumeVisibilityEmployersList**](docs/DefaultApi.md#getresumevisibilityemployerslist) | **Get** /resumes/{resume_id}/{list_type}/search | Поиск работодателей для добавления в список видимости
*DefaultApi* | [**GetResumeVisibilityList**](docs/DefaultApi.md#getresumevisibilitylist) | **Get** /resumes/{resume_id}/{list_type} | Получение списка видимости резюме
*DefaultApi* | [**GetResume_0**](docs/DefaultApi.md#getresume_0) | **Get** /resumes/{resume_id} | Просмотр резюме
*DefaultApi* | [**GetResume_1**](docs/DefaultApi.md#getresume_1) | **Get** /resumes/{resume_id} | Просмотр резюме
*DefaultApi* | [**GetResume_2**](docs/DefaultApi.md#getresume_2) | **Get** /resumes/{resume_id} | Просмотр резюме
*DefaultApi* | [**GetResumesByStatus**](docs/DefaultApi.md#getresumesbystatus) | **Get** /vacancies/{vacancy_id}/resumes_by_status | Резюме, сгруппированные по возможности отклика на данную вакансию
*DefaultApi* | [**GetSalaryEmployeeLevels**](docs/DefaultApi.md#getsalaryemployeelevels) | **Get** /salary_statistics/dictionaries/employee_levels | Уровни компетенций
*DefaultApi* | [**GetSalaryEvaluation**](docs/DefaultApi.md#getsalaryevaluation) | **Get** /salary_statistics/paid/salary_evaluation/{area_id} | Оценка заработной платы без прогноза
*DefaultApi* | [**GetSalaryEvaluation_0**](docs/DefaultApi.md#getsalaryevaluation_0) | **Get** /salary_statistics/paid/salary_evaluation/{area_id} | Оценка заработной платы без прогноза
*DefaultApi* | [**GetSalaryIndustries**](docs/DefaultApi.md#getsalaryindustries) | **Get** /salary_statistics/dictionaries/salary_industries | Отрасли и сферы деятельности
*DefaultApi* | [**GetSalaryProfessionalAreas**](docs/DefaultApi.md#getsalaryprofessionalareas) | **Get** /salary_statistics/dictionaries/professional_areas | Профобласти и специализации
*DefaultApi* | [**GetSalarySalaryAreas**](docs/DefaultApi.md#getsalarysalaryareas) | **Get** /salary_statistics/dictionaries/salary_areas | Регионы и города
*DefaultApi* | [**GetSavedResumeSearch**](docs/DefaultApi.md#getsavedresumesearch) | **Get** /saved_searches/resumes/{id} | Получение единичного сохраненного поиска резюме
*DefaultApi* | [**GetSavedResumeSearches**](docs/DefaultApi.md#getsavedresumesearches) | **Get** /saved_searches/resumes | Список сохраненных поисков резюме
*DefaultApi* | [**GetSavedVacancySearch**](docs/DefaultApi.md#getsavedvacancysearch) | **Get** /saved_searches/vacancies/{id} | Получение единичного сохраненного поиска вакансий
*DefaultApi* | [**GetSavedVacancySearches**](docs/DefaultApi.md#getsavedvacancysearches) | **Get** /saved_searches/vacancies | Список сохраненных поисков вакансий
*DefaultApi* | [**GetSkillSetSuggests**](docs/DefaultApi.md#getskillsetsuggests) | **Get** /suggests/skill_set | Подсказки по ключевым навыкам
*DefaultApi* | [**GetSkills**](docs/DefaultApi.md#getskills) | **Get** /skills | Справочник ключевых навыков
*DefaultApi* | [**GetSuitableResumes**](docs/DefaultApi.md#getsuitableresumes) | **Get** /vacancies/{vacancy_id}/suitable_resumes | Список подходящих для отклика резюме
*DefaultApi* | [**GetTestsDictionary**](docs/DefaultApi.md#gettestsdictionary) | **Get** /employers/{employer_id}/tests | Справочник тестов работодателя
*DefaultApi* | [**GetTestsDictionary_0**](docs/DefaultApi.md#gettestsdictionary_0) | **Get** /employers/{employer_id}/tests | Справочник тестов работодателя
*DefaultApi* | [**GetVacancies**](docs/DefaultApi.md#getvacancies) | **Get** /vacancies | Поиск по вакансиям
*DefaultApi* | [**GetVacanciesSimilarToResume**](docs/DefaultApi.md#getvacanciessimilartoresume) | **Get** /resumes/{resume_id}/similar_vacancies | Поиск по вакансиям, похожим на резюме
*DefaultApi* | [**GetVacanciesSimilarToVacancy**](docs/DefaultApi.md#getvacanciessimilartovacancy) | **Get** /vacancies/{vacancy_id}/similar_vacancies | Поиск по вакансиям, похожим на вакансию
*DefaultApi* | [**GetVacancy**](docs/DefaultApi.md#getvacancy) | **Get** /vacancies/{vacancy_id} | Просмотр вакансии
*DefaultApi* | [**GetVacancyBrandedTemplatesList**](docs/DefaultApi.md#getvacancybrandedtemplateslist) | **Get** /employers/{employer_id}/vacancy_branded_templates | Список доступных бренд шаблонов вакансий работодателя
*DefaultApi* | [**GetVacancyBrandedTemplatesList_0**](docs/DefaultApi.md#getvacancybrandedtemplateslist_0) | **Get** /employers/{employer_id}/vacancy_branded_templates | Список доступных бренд шаблонов вакансий работодателя
*DefaultApi* | [**GetVacancyConditions**](docs/DefaultApi.md#getvacancyconditions) | **Get** /vacancy_conditions | Условия заполнения полей при добавлении и редактировании вакансий
*DefaultApi* | [**GetVacancyDraft**](docs/DefaultApi.md#getvacancydraft) | **Get** /vacancies/drafts/{draft_id} | Получение черновика вакансии
*DefaultApi* | [**GetVacancyDraftList**](docs/DefaultApi.md#getvacancydraftlist) | **Get** /vacancies/drafts | Получение списка черновиков вакансий
*DefaultApi* | [**GetVacancyPositionsSuggests**](docs/DefaultApi.md#getvacancypositionssuggests) | **Get** /suggests/vacancy_positions | Подсказки по должностям вакансий
*DefaultApi* | [**GetVacancySearchKeywords**](docs/DefaultApi.md#getvacancysearchkeywords) | **Get** /suggests/vacancy_search_keyword | Подсказки по ключевым словам поиска вакансий
*DefaultApi* | [**GetVacancyStats**](docs/DefaultApi.md#getvacancystats) | **Get** /vacancies/{vacancy_id}/stats | Статистика по вакансии
*DefaultApi* | [**GetVacancyUpgradeList**](docs/DefaultApi.md#getvacancyupgradelist) | **Get** /vacancies/{vacancy_id}/upgrades | Список улучшений для вакансии
*DefaultApi* | [**GetVacancyVisitors**](docs/DefaultApi.md#getvacancyvisitors) | **Get** /vacancies/{vacancy_id}/visitors | Посмотревшие вакансию
*DefaultApi* | [**GetVacancy_0**](docs/DefaultApi.md#getvacancy_0) | **Get** /vacancies/{vacancy_id} | Просмотр вакансии
*DefaultApi* | [**HideActiveResponse**](docs/DefaultApi.md#hideactiveresponse) | **Delete** /negotiations/active/{nid} | Скрыть отклик
*DefaultApi* | [**InvalidateToken**](docs/DefaultApi.md#invalidatetoken) | **Delete** /oauth/token | Инвалидация токена
*DefaultApi* | [**InvalidateToken_0**](docs/DefaultApi.md#invalidatetoken_0) | **Delete** /oauth/token | Инвалидация токена
*DefaultApi* | [**InvalidateToken_1**](docs/DefaultApi.md#invalidatetoken_1) | **Delete** /oauth/token | Инвалидация токена
*DefaultApi* | [**LoadArtifact**](docs/DefaultApi.md#loadartifact) | **Post** /artifacts | Загрузка артефакта
*DefaultApi* | [**MoveSavedResumeSearch**](docs/DefaultApi.md#movesavedresumesearch) | **Put** /saved_searches/resumes/{saved_search_id}/managers/{manager_id} | Передача сохраненного поиска резюме другому менеджеру
*DefaultApi* | [**PostNegotiationsTopicsRead**](docs/DefaultApi.md#postnegotiationstopicsread) | **Post** /negotiations/read | Отметить отклики прочитанными
*DefaultApi* | [**PostNegotiationsTopicsRead_0**](docs/DefaultApi.md#postnegotiationstopicsread_0) | **Post** /negotiations/read | Отметить отклики прочитанными
*DefaultApi* | [**PublishResume**](docs/DefaultApi.md#publishresume) | **Post** /resumes/{resume_id}/publish | Публикация резюме
*DefaultApi* | [**PublishVacancy**](docs/DefaultApi.md#publishvacancy) | **Post** /vacancies | Публикация вакансии
*DefaultApi* | [**PublishVacancyFromDraft**](docs/DefaultApi.md#publishvacancyfromdraft) | **Post** /vacancies/drafts/{draft_id}/publish | Публикация вакансии на основе черновика
*DefaultApi* | [**PublishVacancy_0**](docs/DefaultApi.md#publishvacancy_0) | **Post** /vacancies | Публикация вакансии
*DefaultApi* | [**PutMailTemplatesItem**](docs/DefaultApi.md#putmailtemplatesitem) | **Put** /employers/{employer_id}/mail_templates/{template_id} | Изменение шаблона ответа соискателю
*DefaultApi* | [**PutMailTemplatesItem_0**](docs/DefaultApi.md#putmailtemplatesitem_0) | **Put** /employers/{employer_id}/mail_templates/{template_id} | Изменение шаблона ответа соискателю
*DefaultApi* | [**PutNegotiationsCollectionToNextState**](docs/DefaultApi.md#putnegotiationscollectiontonextstate) | **Put** /negotiations/{nid} | Действия по откликам/приглашениям
*DefaultApi* | [**PutNegotiationsCollectionToNextState_0**](docs/DefaultApi.md#putnegotiationscollectiontonextstate_0) | **Put** /negotiations/{nid} | Действия по откликам/приглашениям
*DefaultApi* | [**PutPrefNegotiationsOrder**](docs/DefaultApi.md#putprefnegotiationsorder) | **Put** /vacancies/{id}/preferred_negotiations_order | Изменение предпочитаемой сортировки откликов
*DefaultApi* | [**PutPrefNegotiationsOrder_0**](docs/DefaultApi.md#putprefnegotiationsorder_0) | **Put** /vacancies/{id}/preferred_negotiations_order | Изменение предпочитаемой сортировки откликов
*DefaultApi* | [**RestoreVacancyFromHidden**](docs/DefaultApi.md#restorevacancyfromhidden) | **Delete** /employers/{employer_id}/vacancies/hidden/{vacancy_id} | Восстановление вакансии из удаленных
*DefaultApi* | [**SearchEmployer**](docs/DefaultApi.md#searchemployer) | **Get** /employers | Поиск работодателя
*DefaultApi* | [**SearchForResumes**](docs/DefaultApi.md#searchforresumes) | **Get** /resumes | Поиск резюме
*DefaultApi* | [**SearchForResumes_0**](docs/DefaultApi.md#searchforresumes_0) | **Get** /resumes | Поиск резюме
*DefaultApi* | [**SearchForVacancyDraftDuplicates**](docs/DefaultApi.md#searchforvacancydraftduplicates) | **Get** /vacancies/drafts/{draft_id}/duplicates | Проверка наличия дубликатов вакансии
*DefaultApi* | [**SendCodeForVerifyPhoneInResume**](docs/DefaultApi.md#sendcodeforverifyphoneinresume) | **Post** /resume_phone_generate_code | Отправить код подтверждения для телефона резюме
*DefaultApi* | [**SendNegotiationMessage**](docs/DefaultApi.md#sendnegotiationmessage) | **Post** /negotiations/{nid}/messages | Отправка нового сообщения
*DefaultApi* | [**SendNegotiationMessage_0**](docs/DefaultApi.md#sendnegotiationmessage_0) | **Post** /negotiations/{nid}/messages | Отправка нового сообщения
*DefaultApi* | [**UpdateApplicantComment**](docs/DefaultApi.md#updateapplicantcomment) | **Put** /applicant_comments/{applicant_id}/{comment_id} | Обновление комментария
*DefaultApi* | [**UpdateApplicantComment_0**](docs/DefaultApi.md#updateapplicantcomment_0) | **Put** /applicant_comments/{applicant_id}/{comment_id} | Обновление комментария
*DefaultApi* | [**UpdateSavedResumeSearch**](docs/DefaultApi.md#updatesavedresumesearch) | **Put** /saved_searches/resumes/{id} | Обновление сохраненного поиска резюме
*DefaultApi* | [**UpdateSavedVacancySearch**](docs/DefaultApi.md#updatesavedvacancysearch) | **Put** /saved_searches/vacancies/{id} | Обновление сохраненного поиска вакансий
*DefaultApi* | [**VacancyProlongation**](docs/DefaultApi.md#vacancyprolongation) | **Post** /vacancies/{vacancy_id}/prolongate | Продление вакансии
*DefaultApi* | [**VacancyProlongation_0**](docs/DefaultApi.md#vacancyprolongation_0) | **Post** /vacancies/{vacancy_id}/prolongate | Продление вакансии
*WebhookAPIAPI* | [**CancelWebhookSubscription**](docs/WebhookAPIAPI.md#cancelwebhooksubscription) | **Delete** /webhook/subscriptions/{subscription_id} | Удалить подписку на уведомление
*WebhookAPIAPI* | [**ChangeWebhookSubscription**](docs/WebhookAPIAPI.md#changewebhooksubscription) | **Put** /webhook/subscriptions/{subscription_id} | Изменить подписку на уведомления
*WebhookAPIAPI* | [**GetWebhookSubscriptions**](docs/WebhookAPIAPI.md#getwebhooksubscriptions) | **Get** /webhook/subscriptions | Получить список уведомлений, на которые подписан пользователь
*WebhookAPIAPI* | [**PostWebhookSubscription**](docs/WebhookAPIAPI.md#postwebhooksubscription) | **Post** /webhook/subscriptions | Подписаться на уведомления


## Documentation For Models

 - [ApplicantCommentsAccessType](docs/ApplicantCommentsAccessType.md)
 - [ApplicantCommentsApplicantCommentItem](docs/ApplicantCommentsApplicantCommentItem.md)
 - [ApplicantCommentsApplicantCommentItemAccessType](docs/ApplicantCommentsApplicantCommentItemAccessType.md)
 - [ApplicantCommentsApplicantCommentItemAuthor](docs/ApplicantCommentsApplicantCommentItemAuthor.md)
 - [ApplicantCommentsApplicantCommentItems](docs/ApplicantCommentsApplicantCommentItems.md)
 - [ApplicantCommentsApplicantCommentsList](docs/ApplicantCommentsApplicantCommentsList.md)
 - [ApplicantCommentsAuthor](docs/ApplicantCommentsAuthor.md)
 - [ApplyToVacancy400Response](docs/ApplyToVacancy400Response.md)
 - [ArtifactsArtifactConditions](docs/ArtifactsArtifactConditions.md)
 - [ArtifactsArtifactConditionsCounters](docs/ArtifactsArtifactConditionsCounters.md)
 - [ArtifactsArtifactConditionsFields](docs/ArtifactsArtifactConditionsFields.md)
 - [ArtifactsArtifactItem](docs/ArtifactsArtifactItem.md)
 - [ArtifactsArtifactPhotoItems](docs/ArtifactsArtifactPhotoItems.md)
 - [ArtifactsArtifactPhotoResponse](docs/ArtifactsArtifactPhotoResponse.md)
 - [ArtifactsArtifactPortfolioItem](docs/ArtifactsArtifactPortfolioItem.md)
 - [ArtifactsArtifactPortfolioItems](docs/ArtifactsArtifactPortfolioItems.md)
 - [ArtifactsArtifactPortfolioResponse](docs/ArtifactsArtifactPortfolioResponse.md)
 - [ArtifactsCounters](docs/ArtifactsCounters.md)
 - [ArtifactsDescription](docs/ArtifactsDescription.md)
 - [ArtifactsFields](docs/ArtifactsFields.md)
 - [ArtifactsFieldsDescription](docs/ArtifactsFieldsDescription.md)
 - [ArtifactsFieldsFile](docs/ArtifactsFieldsFile.md)
 - [ArtifactsFieldsType](docs/ArtifactsFieldsType.md)
 - [ArtifactsFile](docs/ArtifactsFile.md)
 - [ArtifactsPortfolioDescription](docs/ArtifactsPortfolioDescription.md)
 - [ArtifactsState](docs/ArtifactsState.md)
 - [ArtifactsType](docs/ArtifactsType.md)
 - [AuthAppToken](docs/AuthAppToken.md)
 - [AuthUserToken](docs/AuthUserToken.md)
 - [AuthUserTokenAndAppToken](docs/AuthUserTokenAndAppToken.md)
 - [CredsAnswers](docs/CredsAnswers.md)
 - [CredsCreds](docs/CredsCreds.md)
 - [CredsQuestions](docs/CredsQuestions.md)
 - [DictionariesAreaItem](docs/DictionariesAreaItem.md)
 - [DictionariesBranchItem](docs/DictionariesBranchItem.md)
 - [DictionariesCurrencyItem](docs/DictionariesCurrencyItem.md)
 - [DictionariesDictResponse](docs/DictionariesDictResponse.md)
 - [DictionariesDistrictItem](docs/DictionariesDistrictItem.md)
 - [DictionariesLangItem](docs/DictionariesLangItem.md)
 - [DictionariesSalaryStatisticsAreaItem](docs/DictionariesSalaryStatisticsAreaItem.md)
 - [DictionariesSalaryStatisticsProfessionalAreasResponseInner](docs/DictionariesSalaryStatisticsProfessionalAreasResponseInner.md)
 - [DictionariesSalaryStatisticsSpecializations](docs/DictionariesSalaryStatisticsSpecializations.md)
 - [DictionariesSkillsResponse](docs/DictionariesSkillsResponse.md)
 - [EmployerAddressesEmployerAddressItem](docs/EmployerAddressesEmployerAddressItem.md)
 - [EmployerAddressesEmployerAddressItemManager](docs/EmployerAddressesEmployerAddressItemManager.md)
 - [EmployerAddressesEmployerAddressItemResponse](docs/EmployerAddressesEmployerAddressItemResponse.md)
 - [EmployerAddressesEmployerAddressItems](docs/EmployerAddressesEmployerAddressItems.md)
 - [EmployerAddressesEmployerAddressesResponse](docs/EmployerAddressesEmployerAddressesResponse.md)
 - [EmployerDictionariesTestItem](docs/EmployerDictionariesTestItem.md)
 - [EmployerDictionariesTestsResponse](docs/EmployerDictionariesTestsResponse.md)
 - [EmployerManagerTypesAvailablePermissions](docs/EmployerManagerTypesAvailablePermissions.md)
 - [EmployerManagerTypesEmployerManagerTypesItem](docs/EmployerManagerTypesEmployerManagerTypesItem.md)
 - [EmployerManagerTypesResponse](docs/EmployerManagerTypesResponse.md)
 - [EmployerManagersAddEmployerManager](docs/EmployerManagersAddEmployerManager.md)
 - [EmployerManagersAddEmployerManagerAdditionalPhone](docs/EmployerManagersAddEmployerManagerAdditionalPhone.md)
 - [EmployerManagersAddEmployerManagerPhone](docs/EmployerManagersAddEmployerManagerPhone.md)
 - [EmployerManagersArea](docs/EmployerManagersArea.md)
 - [EmployerManagersAreaId](docs/EmployerManagersAreaId.md)
 - [EmployerManagersEmployerManagerId](docs/EmployerManagersEmployerManagerId.md)
 - [EmployerManagersEmployerManagerInfo](docs/EmployerManagersEmployerManagerInfo.md)
 - [EmployerManagersEmployerManagerItem](docs/EmployerManagersEmployerManagerItem.md)
 - [EmployerManagersEmployerManagerItemAdditionalPhone](docs/EmployerManagersEmployerManagerItemAdditionalPhone.md)
 - [EmployerManagersEmployerManagerLimits](docs/EmployerManagersEmployerManagerLimits.md)
 - [EmployerManagersManagerData](docs/EmployerManagersManagerData.md)
 - [EmployerManagersManagerType](docs/EmployerManagersManagerType.md)
 - [EmployerManagersManagerTypeId](docs/EmployerManagersManagerTypeId.md)
 - [EmployerManagersPermissions](docs/EmployerManagersPermissions.md)
 - [EmployerManagersPhone](docs/EmployerManagersPhone.md)
 - [EmployerManagersResponse](docs/EmployerManagersResponse.md)
 - [EmployerManagersResumeView](docs/EmployerManagersResumeView.md)
 - [EmployerServicesBalance](docs/EmployerServicesBalance.md)
 - [EmployerServicesEmployerServiceItem](docs/EmployerServicesEmployerServiceItem.md)
 - [EmployerServicesEmployerServices](docs/EmployerServicesEmployerServices.md)
 - [EmployerServicesMethodAccess](docs/EmployerServicesMethodAccess.md)
 - [EmployerServicesMethodAccessItem](docs/EmployerServicesMethodAccessItem.md)
 - [EmployerServicesMethodAccessItemAccess](docs/EmployerServicesMethodAccessItemAccess.md)
 - [EmployerServicesServiceType](docs/EmployerServicesServiceType.md)
 - [EmployersBrandingConstructor](docs/EmployersBrandingConstructor.md)
 - [EmployersBrandingConstructorConstructor](docs/EmployersBrandingConstructorConstructor.md)
 - [EmployersBrandingConstructorConstructorWidgetsInner](docs/EmployersBrandingConstructorConstructorWidgetsInner.md)
 - [EmployersBrandingConstructorGalleryWidget](docs/EmployersBrandingConstructorGalleryWidget.md)
 - [EmployersBrandingConstructorGalleryWidgetItemsInner](docs/EmployersBrandingConstructorGalleryWidgetItemsInner.md)
 - [EmployersBrandingConstructorHeaderPicture](docs/EmployersBrandingConstructorHeaderPicture.md)
 - [EmployersBrandingEmployerBranding](docs/EmployersBrandingEmployerBranding.md)
 - [EmployersBrandingMakeup](docs/EmployersBrandingMakeup.md)
 - [EmployersBrandingMakeupMakeup](docs/EmployersBrandingMakeupMakeup.md)
 - [EmployersEmployerDepartmentsResponse](docs/EmployersEmployerDepartmentsResponse.md)
 - [EmployersEmployerInfo](docs/EmployersEmployerInfo.md)
 - [EmployersEmployerInfoArea](docs/EmployersEmployerInfoArea.md)
 - [EmployersEmployerInfoShort](docs/EmployersEmployerInfoShort.md)
 - [EmployersEmployerInfoShortLogoUrl90](docs/EmployersEmployerInfoShortLogoUrl90.md)
 - [EmployersEmployerItem](docs/EmployersEmployerItem.md)
 - [EmployersEmployerItemShort](docs/EmployersEmployerItemShort.md)
 - [EmployersEmployerVacancyAreasItems](docs/EmployersEmployerVacancyAreasItems.md)
 - [EmployersEmployerVacancyAreasResponse](docs/EmployersEmployerVacancyAreasResponse.md)
 - [EmployersEmployersBlacklisted](docs/EmployersEmployersBlacklisted.md)
 - [EmployersEmployersBlacklistedItem](docs/EmployersEmployersBlacklistedItem.md)
 - [EmployersEmployersBlacklistedResponse](docs/EmployersEmployersBlacklistedResponse.md)
 - [EmployersEmployersList](docs/EmployersEmployersList.md)
 - [EmployersEmployersState](docs/EmployersEmployersState.md)
 - [EmployersInsiderInterviews](docs/EmployersInsiderInterviews.md)
 - [EmployersVacancyBrandedTemplatesList](docs/EmployersVacancyBrandedTemplatesList.md)
 - [EmployersVacancyBrandedTemplatesListItemsInner](docs/EmployersVacancyBrandedTemplatesListItemsInner.md)
 - [ErrorsArtifactUploadBadJsonData](docs/ErrorsArtifactUploadBadJsonData.md)
 - [ErrorsArtifactUploadError](docs/ErrorsArtifactUploadError.md)
 - [ErrorsArtifactUploadErrors](docs/ErrorsArtifactUploadErrors.md)
 - [ErrorsBadRequestPublishResumeError](docs/ErrorsBadRequestPublishResumeError.md)
 - [ErrorsBadRequestPublishResumeErrors](docs/ErrorsBadRequestPublishResumeErrors.md)
 - [ErrorsCommonBadArgumentError](docs/ErrorsCommonBadArgumentError.md)
 - [ErrorsCommonBadArgumentErrors](docs/ErrorsCommonBadArgumentErrors.md)
 - [ErrorsCommonBadAuthorizationBadParameters](docs/ErrorsCommonBadAuthorizationBadParameters.md)
 - [ErrorsCommonBadAuthorizationCommonAndEmployerError](docs/ErrorsCommonBadAuthorizationCommonAndEmployerError.md)
 - [ErrorsCommonBadAuthorizationCommonAndPaymentMethodError](docs/ErrorsCommonBadAuthorizationCommonAndPaymentMethodError.md)
 - [ErrorsCommonBadAuthorizationEmployerError](docs/ErrorsCommonBadAuthorizationEmployerError.md)
 - [ErrorsCommonBadAuthorizationError](docs/ErrorsCommonBadAuthorizationError.md)
 - [ErrorsCommonBadAuthorizationErrors](docs/ErrorsCommonBadAuthorizationErrors.md)
 - [ErrorsCommonBadAuthorizationInvalidClientError](docs/ErrorsCommonBadAuthorizationInvalidClientError.md)
 - [ErrorsCommonBadAuthorizationInvalidGrantError](docs/ErrorsCommonBadAuthorizationInvalidGrantError.md)
 - [ErrorsCommonBadAuthorizationInvalidRequestError](docs/ErrorsCommonBadAuthorizationInvalidRequestError.md)
 - [ErrorsCommonBadAuthorizationPaymentMethodError](docs/ErrorsCommonBadAuthorizationPaymentMethodError.md)
 - [ErrorsCommonBadAuthorizationPaymentMethodErrors](docs/ErrorsCommonBadAuthorizationPaymentMethodErrors.md)
 - [ErrorsCommonBadAuthorizationUnsupportedGrantTypeError](docs/ErrorsCommonBadAuthorizationUnsupportedGrantTypeError.md)
 - [ErrorsCommonBadJsonDataError](docs/ErrorsCommonBadJsonDataError.md)
 - [ErrorsCommonBadJsonDataErrors](docs/ErrorsCommonBadJsonDataErrors.md)
 - [ErrorsCommonBadRequestBadArgumentBadRequestAndBadArgumentErrors](docs/ErrorsCommonBadRequestBadArgumentBadRequestAndBadArgumentErrors.md)
 - [ErrorsCommonBadRequestError](docs/ErrorsCommonBadRequestError.md)
 - [ErrorsCommonBadRequestErrors](docs/ErrorsCommonBadRequestErrors.md)
 - [ErrorsCommonCaptchaError](docs/ErrorsCommonCaptchaError.md)
 - [ErrorsCommonCaptchaErrors](docs/ErrorsCommonCaptchaErrors.md)
 - [ErrorsCommonConflictBothChangedError](docs/ErrorsCommonConflictBothChangedError.md)
 - [ErrorsCommonConflictBothChangedErrors](docs/ErrorsCommonConflictBothChangedErrors.md)
 - [ErrorsCommonErrorRequestId](docs/ErrorsCommonErrorRequestId.md)
 - [ErrorsCommonNotFoundError](docs/ErrorsCommonNotFoundError.md)
 - [ErrorsCommonNotFoundErrors](docs/ErrorsCommonNotFoundErrors.md)
 - [ErrorsDictionariesBadArgumentError](docs/ErrorsDictionariesBadArgumentError.md)
 - [ErrorsDictionariesBadArgumentErrors](docs/ErrorsDictionariesBadArgumentErrors.md)
 - [ErrorsDraftError](docs/ErrorsDraftError.md)
 - [ErrorsEmployerAddressError](docs/ErrorsEmployerAddressError.md)
 - [ErrorsEmployerAddressErrors](docs/ErrorsEmployerAddressErrors.md)
 - [ErrorsEmployerBlacklistedError](docs/ErrorsEmployerBlacklistedError.md)
 - [ErrorsEmployerBlacklistedErrors](docs/ErrorsEmployerBlacklistedErrors.md)
 - [ErrorsEmployerBlacklistedNotFoundError](docs/ErrorsEmployerBlacklistedNotFoundError.md)
 - [ErrorsEmployerBlacklistedNotFoundErrors](docs/ErrorsEmployerBlacklistedNotFoundErrors.md)
 - [ErrorsEmployerManagerBadArgumentError](docs/ErrorsEmployerManagerBadArgumentError.md)
 - [ErrorsEmployerManagerBadArgumentErrors](docs/ErrorsEmployerManagerBadArgumentErrors.md)
 - [ErrorsEmployerManagerBadAuthorizationError](docs/ErrorsEmployerManagerBadAuthorizationError.md)
 - [ErrorsEmployerManagerBadAuthorizationErrors](docs/ErrorsEmployerManagerBadAuthorizationErrors.md)
 - [ErrorsNegotiationEditMessageError](docs/ErrorsNegotiationEditMessageError.md)
 - [ErrorsNegotiationEditMessageErrors](docs/ErrorsNegotiationEditMessageErrors.md)
 - [ErrorsNegotiationEditMessageForbiddenErrors](docs/ErrorsNegotiationEditMessageForbiddenErrors.md)
 - [ErrorsNegotiationHideResponseError](docs/ErrorsNegotiationHideResponseError.md)
 - [ErrorsNegotiationHideResponseErrors](docs/ErrorsNegotiationHideResponseErrors.md)
 - [ErrorsNegotiationHideResponseForbiddenErrors](docs/ErrorsNegotiationHideResponseForbiddenErrors.md)
 - [ErrorsNegotiationNegotiationsNotFoundErrors](docs/ErrorsNegotiationNegotiationsNotFoundErrors.md)
 - [ErrorsNegotiationNotFoundError](docs/ErrorsNegotiationNotFoundError.md)
 - [ErrorsNegotiationNotFoundErrors](docs/ErrorsNegotiationNotFoundErrors.md)
 - [ErrorsRequestEntityTooLargeError](docs/ErrorsRequestEntityTooLargeError.md)
 - [ErrorsRequestEntityTooLargeErrors](docs/ErrorsRequestEntityTooLargeErrors.md)
 - [ErrorsResumeAddEditError](docs/ErrorsResumeAddEditError.md)
 - [ErrorsResumeAddEditErrors](docs/ErrorsResumeAddEditErrors.md)
 - [ErrorsResumeBadArgTooManyResumesErrors](docs/ErrorsResumeBadArgTooManyResumesErrors.md)
 - [ErrorsResumeBadArgumentResumeErrors](docs/ErrorsResumeBadArgumentResumeErrors.md)
 - [ErrorsResumeTooManyRequestsError](docs/ErrorsResumeTooManyRequestsError.md)
 - [ErrorsResumeTooManyRequestsErrors](docs/ErrorsResumeTooManyRequestsErrors.md)
 - [ErrorsResumeTooManyResumesError](docs/ErrorsResumeTooManyResumesError.md)
 - [ErrorsResumeTooManyResumesErrors](docs/ErrorsResumeTooManyResumesErrors.md)
 - [ErrorsResumeVisibilityError](docs/ErrorsResumeVisibilityError.md)
 - [ErrorsResumeVisibilityErrors](docs/ErrorsResumeVisibilityErrors.md)
 - [ErrorsResumeVisibilityErrorsBadRequest](docs/ErrorsResumeVisibilityErrorsBadRequest.md)
 - [ErrorsSavedSearchForbiddenError](docs/ErrorsSavedSearchForbiddenError.md)
 - [ErrorsSavedSearchForbiddenErrors](docs/ErrorsSavedSearchForbiddenErrors.md)
 - [ErrorsSavedSearchNotFoundError](docs/ErrorsSavedSearchNotFoundError.md)
 - [ErrorsSavedSearchNotFoundErrors](docs/ErrorsSavedSearchNotFoundErrors.md)
 - [ErrorsVacancyAddEditBadAuthForbiddenErrors](docs/ErrorsVacancyAddEditBadAuthForbiddenErrors.md)
 - [ErrorsVacancyAddEditBadJsonDataError](docs/ErrorsVacancyAddEditBadJsonDataError.md)
 - [ErrorsVacancyAddEditBadJsonDataErrors](docs/ErrorsVacancyAddEditBadJsonDataErrors.md)
 - [ErrorsVacancyAddEditCombinedBadJsonDataErrors](docs/ErrorsVacancyAddEditCombinedBadJsonDataErrors.md)
 - [ErrorsVacancyAddEditForbiddenError](docs/ErrorsVacancyAddEditForbiddenError.md)
 - [ErrorsVacancyAddEditForbiddenErrors](docs/ErrorsVacancyAddEditForbiddenErrors.md)
 - [ErrorsVacancyApplyBadRequestError](docs/ErrorsVacancyApplyBadRequestError.md)
 - [ErrorsVacancyApplyBadRequestErrors](docs/ErrorsVacancyApplyBadRequestErrors.md)
 - [ErrorsVacancyApplyForbiddenError](docs/ErrorsVacancyApplyForbiddenError.md)
 - [ErrorsVacancyApplyForbiddenErrors](docs/ErrorsVacancyApplyForbiddenErrors.md)
 - [ErrorsVacancyApplyForbiddenErrorsAllOfBadArguments](docs/ErrorsVacancyApplyForbiddenErrorsAllOfBadArguments.md)
 - [ErrorsVacancyBlacklistedBadRequestError](docs/ErrorsVacancyBlacklistedBadRequestError.md)
 - [ErrorsVacancyBlacklistedBadRequestErrors](docs/ErrorsVacancyBlacklistedBadRequestErrors.md)
 - [ErrorsVacancyBlacklistedNotFoundError](docs/ErrorsVacancyBlacklistedNotFoundError.md)
 - [ErrorsVacancyBlacklistedNotFoundErrors](docs/ErrorsVacancyBlacklistedNotFoundErrors.md)
 - [ErrorsVacancyError](docs/ErrorsVacancyError.md)
 - [ErrorsVacancyErrorItemsInner](docs/ErrorsVacancyErrorItemsInner.md)
 - [ErrorsVacancyErrors](docs/ErrorsVacancyErrors.md)
 - [ErrorsVacancyFavoritedBadAuthError](docs/ErrorsVacancyFavoritedBadAuthError.md)
 - [ErrorsVacancyFavoritedBadAuthErrors](docs/ErrorsVacancyFavoritedBadAuthErrors.md)
 - [ErrorsVacancyFavoritedCombinedBadAuthErrors](docs/ErrorsVacancyFavoritedCombinedBadAuthErrors.md)
 - [ErrorsVacancyProlongateError](docs/ErrorsVacancyProlongateError.md)
 - [ErrorsVacancyProlongateErrors](docs/ErrorsVacancyProlongateErrors.md)
 - [ErrorsVacancyProlongateForbidden](docs/ErrorsVacancyProlongateForbidden.md)
 - [GeocoderAddress](docs/GeocoderAddress.md)
 - [Id](docs/Id.md)
 - [IncludesArea](docs/IncludesArea.md)
 - [IncludesClusterMetroLine](docs/IncludesClusterMetroLine.md)
 - [IncludesClusterMetroLineArea](docs/IncludesClusterMetroLineArea.md)
 - [IncludesClusterMetroStation](docs/IncludesClusterMetroStation.md)
 - [IncludesClusterMetroStationArea](docs/IncludesClusterMetroStationArea.md)
 - [IncludesContact](docs/IncludesContact.md)
 - [IncludesContactPhoneValue](docs/IncludesContactPhoneValue.md)
 - [IncludesContactProperties](docs/IncludesContactProperties.md)
 - [IncludesContactPropertiesValue](docs/IncludesContactPropertiesValue.md)
 - [IncludesCountUrl](docs/IncludesCountUrl.md)
 - [IncludesEducationalInstitutionItem](docs/IncludesEducationalInstitutionItem.md)
 - [IncludesEducationalInstitutionItemArea](docs/IncludesEducationalInstitutionItemArea.md)
 - [IncludesEmployerApplicantServices](docs/IncludesEmployerApplicantServices.md)
 - [IncludesEmployerApplicantServicesObject](docs/IncludesEmployerApplicantServicesObject.md)
 - [IncludesEmployerApplicantServicesTargetEmployer](docs/IncludesEmployerApplicantServicesTargetEmployer.md)
 - [IncludesId](docs/IncludesId.md)
 - [IncludesIdName](docs/IncludesIdName.md)
 - [IncludesIdNameDesc](docs/IncludesIdNameDesc.md)
 - [IncludesIdNameLastChangeTime](docs/IncludesIdNameLastChangeTime.md)
 - [IncludesIdNameUid](docs/IncludesIdNameUid.md)
 - [IncludesIdNameUrl](docs/IncludesIdNameUrl.md)
 - [IncludesLanguageLevel](docs/IncludesLanguageLevel.md)
 - [IncludesLanguageProperties](docs/IncludesLanguageProperties.md)
 - [IncludesLogoUrls](docs/IncludesLogoUrls.md)
 - [IncludesLogoUrls90](docs/IncludesLogoUrls90.md)
 - [IncludesMetroStation](docs/IncludesMetroStation.md)
 - [IncludesName](docs/IncludesName.md)
 - [IncludesNameDescription](docs/IncludesNameDescription.md)
 - [IncludesNullableObject](docs/IncludesNullableObject.md)
 - [IncludesNumericId](docs/IncludesNumericId.md)
 - [IncludesPagination](docs/IncludesPagination.md)
 - [IncludesSkillSetItem](docs/IncludesSkillSetItem.md)
 - [IncludesSortingType](docs/IncludesSortingType.md)
 - [IncludesUrl](docs/IncludesUrl.md)
 - [LocalesLocaleItem](docs/LocalesLocaleItem.md)
 - [LocalesResumeLocaleItem](docs/LocalesResumeLocaleItem.md)
 - [MailTemplatesMailTemplate](docs/MailTemplatesMailTemplate.md)
 - [MailTemplatesMailTemplateInput](docs/MailTemplatesMailTemplateInput.md)
 - [ManagerAccount](docs/ManagerAccount.md)
 - [ManagerAccountCompany](docs/ManagerAccountCompany.md)
 - [ManagerAccounts](docs/ManagerAccounts.md)
 - [ManagerSettings](docs/ManagerSettings.md)
 - [ManagerSettingsCurrency](docs/ManagerSettingsCurrency.md)
 - [MeAnyProfile](docs/MeAnyProfile.md)
 - [MeAnyUserProfile](docs/MeAnyUserProfile.md)
 - [MeApplicantProfile](docs/MeApplicantProfile.md)
 - [MeApplicantProfileCounters](docs/MeApplicantProfileCounters.md)
 - [MeCommonProfile](docs/MeCommonProfile.md)
 - [MeEmployerProfile](docs/MeEmployerProfile.md)
 - [MeEmployerProfileCompany](docs/MeEmployerProfileCompany.md)
 - [MeEmployerProfileCompanyDeprecated](docs/MeEmployerProfileCompanyDeprecated.md)
 - [MeEmployerProfileManager](docs/MeEmployerProfileManager.md)
 - [MeEmployerProfilePersonalManager](docs/MeEmployerProfilePersonalManager.md)
 - [MeEmployerProfilePersonalManagerPhotoUrls](docs/MeEmployerProfilePersonalManagerPhotoUrls.md)
 - [MeEmployerProfilePersonalManagerUnavailable](docs/MeEmployerProfilePersonalManagerUnavailable.md)
 - [MeManagerProfile](docs/MeManagerProfile.md)
 - [MeProfile](docs/MeProfile.md)
 - [MetroCityMetroItem](docs/MetroCityMetroItem.md)
 - [MetroLineStation](docs/MetroLineStation.md)
 - [MetroMetroItem](docs/MetroMetroItem.md)
 - [MetroMetroLine](docs/MetroMetroLine.md)
 - [MetroMetroLineWithStations](docs/MetroMetroLineWithStations.md)
 - [NegotiationsAction](docs/NegotiationsAction.md)
 - [NegotiationsApplicantNegotiation](docs/NegotiationsApplicantNegotiation.md)
 - [NegotiationsApplicantNegotiationsResponse](docs/NegotiationsApplicantNegotiationsResponse.md)
 - [NegotiationsAssessment](docs/NegotiationsAssessment.md)
 - [NegotiationsAuthor](docs/NegotiationsAuthor.md)
 - [NegotiationsCollectionNegotiations](docs/NegotiationsCollectionNegotiations.md)
 - [NegotiationsCollectionNegotiationsItemsInner](docs/NegotiationsCollectionNegotiationsItemsInner.md)
 - [NegotiationsCollectionNegotiationsResponse](docs/NegotiationsCollectionNegotiationsResponse.md)
 - [NegotiationsEmployerNegotiation](docs/NegotiationsEmployerNegotiation.md)
 - [NegotiationsListItem](docs/NegotiationsListItem.md)
 - [NegotiationsListItems](docs/NegotiationsListItems.md)
 - [NegotiationsListResponse](docs/NegotiationsListResponse.md)
 - [NegotiationsMessage](docs/NegotiationsMessage.md)
 - [NegotiationsMessageSent](docs/NegotiationsMessageSent.md)
 - [NegotiationsMessagesGet](docs/NegotiationsMessagesGet.md)
 - [NegotiationsMessagesGetResponse](docs/NegotiationsMessagesGetResponse.md)
 - [NegotiationsNegotiationCommonFields](docs/NegotiationsNegotiationCommonFields.md)
 - [NegotiationsNegotiationGetResponse](docs/NegotiationsNegotiationGetResponse.md)
 - [NegotiationsNegotiationMessageTemplate](docs/NegotiationsNegotiationMessageTemplate.md)
 - [NegotiationsNegotiationMessageTemplates](docs/NegotiationsNegotiationMessageTemplates.md)
 - [NegotiationsNegotiationTestResultsResponse](docs/NegotiationsNegotiationTestResultsResponse.md)
 - [NegotiationsNegotiationTestResultsResponseTestResult](docs/NegotiationsNegotiationTestResultsResponseTestResult.md)
 - [NegotiationsNegotiationsCollection](docs/NegotiationsNegotiationsCollection.md)
 - [NegotiationsNegotiationsCollections](docs/NegotiationsNegotiationsCollections.md)
 - [NegotiationsNegotiationsOrderTypes](docs/NegotiationsNegotiationsOrderTypes.md)
 - [NegotiationsNegotiationsStatistics](docs/NegotiationsNegotiationsStatistics.md)
 - [NegotiationsNegotiationsStatisticsEmployerResponse](docs/NegotiationsNegotiationsStatisticsEmployerResponse.md)
 - [NegotiationsNegotiationsStatisticsManagerResponse](docs/NegotiationsNegotiationsStatisticsManagerResponse.md)
 - [NegotiationsObjectsCounters](docs/NegotiationsObjectsCounters.md)
 - [NegotiationsObjectsEmployerCounters](docs/NegotiationsObjectsEmployerCounters.md)
 - [NegotiationsObjectsEmployerTopicResume](docs/NegotiationsObjectsEmployerTopicResume.md)
 - [NegotiationsObjectsPoliteness](docs/NegotiationsObjectsPoliteness.md)
 - [NegotiationsObjectsStates](docs/NegotiationsObjectsStates.md)
 - [NegotiationsObjectsTopicItem](docs/NegotiationsObjectsTopicItem.md)
 - [NegotiationsObjectsTopicItemCommon](docs/NegotiationsObjectsTopicItemCommon.md)
 - [NegotiationsObjectsTopicResume](docs/NegotiationsObjectsTopicResume.md)
 - [NegotiationsPhoneCallItem](docs/NegotiationsPhoneCallItem.md)
 - [NegotiationsPhoneCalls](docs/NegotiationsPhoneCalls.md)
 - [ProfessionalRolesCatalog](docs/ProfessionalRolesCatalog.md)
 - [ProfessionalRolesCategory](docs/ProfessionalRolesCategory.md)
 - [ProfessionalRolesRole](docs/ProfessionalRolesRole.md)
 - [ProfilePhoto](docs/ProfilePhoto.md)
 - [ProfileVideosList](docs/ProfileVideosList.md)
 - [ProfileVideosListItemsInner](docs/ProfileVideosListItemsInner.md)
 - [ProfileVideosListItemsInnerDownloadUrl](docs/ProfileVideosListItemsInnerDownloadUrl.md)
 - [ResumeAddResumeRequest](docs/ResumeAddResumeRequest.md)
 - [ResumeApplicantFields](docs/ResumeApplicantFields.md)
 - [ResumeEditResumeRequest](docs/ResumeEditResumeRequest.md)
 - [ResumeEmployerFields](docs/ResumeEmployerFields.md)
 - [ResumeNullableFields](docs/ResumeNullableFields.md)
 - [ResumeObjectsAccess](docs/ResumeObjectsAccess.md)
 - [ResumeObjectsActions](docs/ResumeObjectsActions.md)
 - [ResumeObjectsActionsForOwner](docs/ResumeObjectsActionsForOwner.md)
 - [ResumeObjectsCertificate](docs/ResumeObjectsCertificate.md)
 - [ResumeObjectsContact](docs/ResumeObjectsContact.md)
 - [ResumeObjectsCounters](docs/ResumeObjectsCounters.md)
 - [ResumeObjectsDownload](docs/ResumeObjectsDownload.md)
 - [ResumeObjectsDownloadPdfRtf](docs/ResumeObjectsDownloadPdfRtf.md)
 - [ResumeObjectsDriverLicenseTypes](docs/ResumeObjectsDriverLicenseTypes.md)
 - [ResumeObjectsEducation](docs/ResumeObjectsEducation.md)
 - [ResumeObjectsEducationAdditional](docs/ResumeObjectsEducationAdditional.md)
 - [ResumeObjectsEducationElementary](docs/ResumeObjectsEducationElementary.md)
 - [ResumeObjectsEducationPrimary](docs/ResumeObjectsEducationPrimary.md)
 - [ResumeObjectsEmployerPaidServicesInner](docs/ResumeObjectsEmployerPaidServicesInner.md)
 - [ResumeObjectsEmployerPaidServicesInnerPriceList](docs/ResumeObjectsEmployerPaidServicesInnerPriceList.md)
 - [ResumeObjectsEmployerPaidServicesInnerQuickPurchase](docs/ResumeObjectsEmployerPaidServicesInnerQuickPurchase.md)
 - [ResumeObjectsEmployerPaidServicesInnerQuickPurchaseCurrency](docs/ResumeObjectsEmployerPaidServicesInnerQuickPurchaseCurrency.md)
 - [ResumeObjectsExperience](docs/ResumeObjectsExperience.md)
 - [ResumeObjectsExperienceCreateEditResume](docs/ResumeObjectsExperienceCreateEditResume.md)
 - [ResumeObjectsExperienceForOwner](docs/ResumeObjectsExperienceForOwner.md)
 - [ResumeObjectsExperienceLogoUrl90](docs/ResumeObjectsExperienceLogoUrl90.md)
 - [ResumeObjectsExperienceProperties](docs/ResumeObjectsExperienceProperties.md)
 - [ResumeObjectsExperiencePropertiesLogoUrl90](docs/ResumeObjectsExperiencePropertiesLogoUrl90.md)
 - [ResumeObjectsIndustry](docs/ResumeObjectsIndustry.md)
 - [ResumeObjectsLanguage](docs/ResumeObjectsLanguage.md)
 - [ResumeObjectsLastNegotiations](docs/ResumeObjectsLastNegotiations.md)
 - [ResumeObjectsMetroLine](docs/ResumeObjectsMetroLine.md)
 - [ResumeObjectsMetroStation](docs/ResumeObjectsMetroStation.md)
 - [ResumeObjectsModerationNote](docs/ResumeObjectsModerationNote.md)
 - [ResumeObjectsNegotiationsHistoryForEmployer](docs/ResumeObjectsNegotiationsHistoryForEmployer.md)
 - [ResumeObjectsNegotiationsHistoryUrl](docs/ResumeObjectsNegotiationsHistoryUrl.md)
 - [ResumeObjectsOwner](docs/ResumeObjectsOwner.md)
 - [ResumeObjectsOwnerComments](docs/ResumeObjectsOwnerComments.md)
 - [ResumeObjectsOwnerCommentsCounters](docs/ResumeObjectsOwnerCommentsCounters.md)
 - [ResumeObjectsPaidServices](docs/ResumeObjectsPaidServices.md)
 - [ResumeObjectsPhoto](docs/ResumeObjectsPhoto.md)
 - [ResumeObjectsPhotoNoId](docs/ResumeObjectsPhotoNoId.md)
 - [ResumeObjectsPortfolio](docs/ResumeObjectsPortfolio.md)
 - [ResumeObjectsPortfolioNoId](docs/ResumeObjectsPortfolioNoId.md)
 - [ResumeObjectsProgress](docs/ResumeObjectsProgress.md)
 - [ResumeObjectsRecommendation](docs/ResumeObjectsRecommendation.md)
 - [ResumeObjectsRelocationPublic](docs/ResumeObjectsRelocationPublic.md)
 - [ResumeObjectsSalaryAddEdit](docs/ResumeObjectsSalaryAddEdit.md)
 - [ResumeObjectsSalaryProperties](docs/ResumeObjectsSalaryProperties.md)
 - [ResumeObjectsSimilarVacancies](docs/ResumeObjectsSimilarVacancies.md)
 - [ResumeObjectsSite](docs/ResumeObjectsSite.md)
 - [ResumeObjectsTotalExperience](docs/ResumeObjectsTotalExperience.md)
 - [ResumePhoneGenerateCodeGenerateCode](docs/ResumePhoneGenerateCodeGenerateCode.md)
 - [ResumePhotoPortfolio](docs/ResumePhotoPortfolio.md)
 - [ResumeReadiness](docs/ResumeReadiness.md)
 - [ResumeResumeForApplicant](docs/ResumeResumeForApplicant.md)
 - [ResumeResumeForEmployer](docs/ResumeResumeForEmployer.md)
 - [ResumeResumeFull](docs/ResumeResumeFull.md)
 - [ResumeResumeNano](docs/ResumeResumeNano.md)
 - [ResumeResumeNanoWithUrl](docs/ResumeResumeNanoWithUrl.md)
 - [ResumeResumeProfile](docs/ResumeResumeProfile.md)
 - [ResumeResumeProfileLogoUrl90](docs/ResumeResumeProfileLogoUrl90.md)
 - [ResumeResumeShort](docs/ResumeResumeShort.md)
 - [ResumeResumeShortAdditionalFields](docs/ResumeResumeShortAdditionalFields.md)
 - [ResumeResumeShortForOwner](docs/ResumeResumeShortForOwner.md)
 - [ResumeResumeShortLogoUrl90](docs/ResumeResumeShortLogoUrl90.md)
 - [ResumeResumeViewResponse](docs/ResumeResumeViewResponse.md)
 - [ResumeShouldSendSmsContainer](docs/ResumeShouldSendSmsContainer.md)
 - [ResumeShouldSendSmsPhone](docs/ResumeShouldSendSmsPhone.md)
 - [ResumeStatus](docs/ResumeStatus.md)
 - [ResumeStatusReadiness](docs/ResumeStatusReadiness.md)
 - [ResumesAccessTypes](docs/ResumesAccessTypes.md)
 - [ResumesAccessTypesItem](docs/ResumesAccessTypesItem.md)
 - [ResumesAutoHideTimeOptions](docs/ResumesAutoHideTimeOptions.md)
 - [ResumesByStatusCounters](docs/ResumesByStatusCounters.md)
 - [ResumesByStatusResponse](docs/ResumesByStatusResponse.md)
 - [ResumesCreationAvailability](docs/ResumesCreationAvailability.md)
 - [ResumesGetResumeVisibilityListItem](docs/ResumesGetResumeVisibilityListItem.md)
 - [ResumesGetResumeVisibilityListProperties](docs/ResumesGetResumeVisibilityListProperties.md)
 - [ResumesGetResumeVisibilityListResponse](docs/ResumesGetResumeVisibilityListResponse.md)
 - [ResumesMineItem](docs/ResumesMineItem.md)
 - [ResumesMineItems](docs/ResumesMineItems.md)
 - [ResumesMineResponse](docs/ResumesMineResponse.md)
 - [ResumesNegotiationNano](docs/ResumesNegotiationNano.md)
 - [ResumesNegotiationNanoEmployerState](docs/ResumesNegotiationNanoEmployerState.md)
 - [ResumesPostResumeVisibilityListBody](docs/ResumesPostResumeVisibilityListBody.md)
 - [ResumesResumeConditionFieldsAccessCondition](docs/ResumesResumeConditionFieldsAccessCondition.md)
 - [ResumesResumeConditionFieldsAccessFields](docs/ResumesResumeConditionFieldsAccessFields.md)
 - [ResumesResumeConditionFieldsContactCondition](docs/ResumesResumeConditionFieldsContactCondition.md)
 - [ResumesResumeConditionFieldsContactFields](docs/ResumesResumeConditionFieldsContactFields.md)
 - [ResumesResumeConditionFieldsEducationCondition](docs/ResumesResumeConditionFieldsEducationCondition.md)
 - [ResumesResumeConditionFieldsEducationElementaryCondition](docs/ResumesResumeConditionFieldsEducationElementaryCondition.md)
 - [ResumesResumeConditionFieldsEducationFields](docs/ResumesResumeConditionFieldsEducationFields.md)
 - [ResumesResumeConditionFieldsEducationPrimaryCondition](docs/ResumesResumeConditionFieldsEducationPrimaryCondition.md)
 - [ResumesResumeConditionFieldsElementaryEducationFields](docs/ResumesResumeConditionFieldsElementaryEducationFields.md)
 - [ResumesResumeConditionFieldsExperienceCondition](docs/ResumesResumeConditionFieldsExperienceCondition.md)
 - [ResumesResumeConditionFieldsExperienceFields](docs/ResumesResumeConditionFieldsExperienceFields.md)
 - [ResumesResumeConditionFieldsLanguageCondition](docs/ResumesResumeConditionFieldsLanguageCondition.md)
 - [ResumesResumeConditionFieldsLanguageFields](docs/ResumesResumeConditionFieldsLanguageFields.md)
 - [ResumesResumeConditionFieldsMaxMinCount](docs/ResumesResumeConditionFieldsMaxMinCount.md)
 - [ResumesResumeConditionFieldsMaxMinDate](docs/ResumesResumeConditionFieldsMaxMinDate.md)
 - [ResumesResumeConditionFieldsMaxMinLength](docs/ResumesResumeConditionFieldsMaxMinLength.md)
 - [ResumesResumeConditionFieldsMaxMinValue](docs/ResumesResumeConditionFieldsMaxMinValue.md)
 - [ResumesResumeConditionFieldsNotIn](docs/ResumesResumeConditionFieldsNotIn.md)
 - [ResumesResumeConditionFieldsPrimaryEducationFields](docs/ResumesResumeConditionFieldsPrimaryEducationFields.md)
 - [ResumesResumeConditionFieldsRecommendationCondition](docs/ResumesResumeConditionFieldsRecommendationCondition.md)
 - [ResumesResumeConditionFieldsRecommendationFields](docs/ResumesResumeConditionFieldsRecommendationFields.md)
 - [ResumesResumeConditionFieldsRegexp](docs/ResumesResumeConditionFieldsRegexp.md)
 - [ResumesResumeConditionFieldsRelocationCondition](docs/ResumesResumeConditionFieldsRelocationCondition.md)
 - [ResumesResumeConditionFieldsRelocationFields](docs/ResumesResumeConditionFieldsRelocationFields.md)
 - [ResumesResumeConditionFieldsRequired](docs/ResumesResumeConditionFieldsRequired.md)
 - [ResumesResumeConditionFieldsRequiredCountAndLengthTitle](docs/ResumesResumeConditionFieldsRequiredCountAndLengthTitle.md)
 - [ResumesResumeConditionFieldsRequiredCountWithTitle](docs/ResumesResumeConditionFieldsRequiredCountWithTitle.md)
 - [ResumesResumeConditionFieldsRequiredDateWithTitle](docs/ResumesResumeConditionFieldsRequiredDateWithTitle.md)
 - [ResumesResumeConditionFieldsRequiredLengthTitleNotInt](docs/ResumesResumeConditionFieldsRequiredLengthTitleNotInt.md)
 - [ResumesResumeConditionFieldsRequiredLengthTitleRegexp](docs/ResumesResumeConditionFieldsRequiredLengthTitleRegexp.md)
 - [ResumesResumeConditionFieldsRequiredLengthWithTitle](docs/ResumesResumeConditionFieldsRequiredLengthWithTitle.md)
 - [ResumesResumeConditionFieldsRequiredValueWithTitle](docs/ResumesResumeConditionFieldsRequiredValueWithTitle.md)
 - [ResumesResumeConditionFieldsRequiredWithTitle](docs/ResumesResumeConditionFieldsRequiredWithTitle.md)
 - [ResumesResumeConditionFieldsSalaryCondition](docs/ResumesResumeConditionFieldsSalaryCondition.md)
 - [ResumesResumeConditionFieldsSalaryFields](docs/ResumesResumeConditionFieldsSalaryFields.md)
 - [ResumesResumeConditionFieldsSiteCondition](docs/ResumesResumeConditionFieldsSiteCondition.md)
 - [ResumesResumeConditionFieldsSiteFields](docs/ResumesResumeConditionFieldsSiteFields.md)
 - [ResumesResumeConditions](docs/ResumesResumeConditions.md)
 - [ResumesResumeNegotiationsHistoryResponse](docs/ResumesResumeNegotiationsHistoryResponse.md)
 - [ResumesResumeNegotiationsHistoryVacancy](docs/ResumesResumeNegotiationsHistoryVacancy.md)
 - [ResumesResumeNegotiationsHistoryVacancyItem](docs/ResumesResumeNegotiationsHistoryVacancyItem.md)
 - [ResumesResumeViewHistory](docs/ResumesResumeViewHistory.md)
 - [ResumesResumeViewHistoryItem](docs/ResumesResumeViewHistoryItem.md)
 - [ResumesResumeViewHistoryItemEmployer](docs/ResumesResumeViewHistoryItemEmployer.md)
 - [ResumesResumeViewHistoryResponse](docs/ResumesResumeViewHistoryResponse.md)
 - [ResumesResumeViewHistoryResume](docs/ResumesResumeViewHistoryResume.md)
 - [ResumesResumeVisibilityListSearchProperties](docs/ResumesResumeVisibilityListSearchProperties.md)
 - [ResumesResumeVisibilityListSearchPropertiesItemsInner](docs/ResumesResumeVisibilityListSearchPropertiesItemsInner.md)
 - [ResumesResumeVisibilityListSearchResponse](docs/ResumesResumeVisibilityListSearchResponse.md)
 - [ResumesSearchForEmployerAndApplicant](docs/ResumesSearchForEmployerAndApplicant.md)
 - [ResumesSearchForResumesItems](docs/ResumesSearchForResumesItems.md)
 - [ResumesSearchForResumesResponse](docs/ResumesSearchForResumesResponse.md)
 - [ResumesSelectedObject](docs/ResumesSelectedObject.md)
 - [ResumesSuitableResumeItem](docs/ResumesSuitableResumeItem.md)
 - [ResumesSuitableResumeItems](docs/ResumesSuitableResumeItems.md)
 - [ResumesSuitableResumeOverall](docs/ResumesSuitableResumeOverall.md)
 - [ResumesSuitableResumesResponse](docs/ResumesSuitableResumesResponse.md)
 - [SalaryStatisticsEvaluationResponse](docs/SalaryStatisticsEvaluationResponse.md)
 - [SalaryStatisticsIndirectCalculation](docs/SalaryStatisticsIndirectCalculation.md)
 - [SalaryStatisticsMarketSalary](docs/SalaryStatisticsMarketSalary.md)
 - [SalaryStatisticsResultingParameters](docs/SalaryStatisticsResultingParameters.md)
 - [SavedSearchesSavedSearchItem](docs/SavedSearchesSavedSearchItem.md)
 - [SavedSearchesSavedSearchItemItems](docs/SavedSearchesSavedSearchItemItems.md)
 - [SavedSearchesSavedSearchItemNewItems](docs/SavedSearchesSavedSearchItemNewItems.md)
 - [SavedSearchesSavedSearchItems](docs/SavedSearchesSavedSearchItems.md)
 - [SavedSearchesSavedSearchResponse](docs/SavedSearchesSavedSearchResponse.md)
 - [SendNegotiationMessage403Response](docs/SendNegotiationMessage403Response.md)
 - [SkillVerificationsOpenedAnswer](docs/SkillVerificationsOpenedAnswer.md)
 - [SkillVerificationsTestResultNano](docs/SkillVerificationsTestResultNano.md)
 - [SkillVerificationsTestResultTasks](docs/SkillVerificationsTestResultTasks.md)
 - [SkillVerificationsTestResultTasksClosedAnswersInner](docs/SkillVerificationsTestResultTasksClosedAnswersInner.md)
 - [SkillVerificationsTestResultWithUrl](docs/SkillVerificationsTestResultWithUrl.md)
 - [SuggestsArea](docs/SuggestsArea.md)
 - [SuggestsAreas](docs/SuggestsAreas.md)
 - [SuggestsAreasItem](docs/SuggestsAreasItem.md)
 - [SuggestsCompanies](docs/SuggestsCompanies.md)
 - [SuggestsCompanyItem](docs/SuggestsCompanyItem.md)
 - [SuggestsEducationalInstitutions](docs/SuggestsEducationalInstitutions.md)
 - [SuggestsEmployerItem](docs/SuggestsEmployerItem.md)
 - [SuggestsErrors](docs/SuggestsErrors.md)
 - [SuggestsErrorsAllOfBadArguments](docs/SuggestsErrorsAllOfBadArguments.md)
 - [SuggestsErrorsAllOfErrors](docs/SuggestsErrorsAllOfErrors.md)
 - [SuggestsFieldsOfStudy](docs/SuggestsFieldsOfStudy.md)
 - [SuggestsFieldsOfStudyItem](docs/SuggestsFieldsOfStudyItem.md)
 - [SuggestsLogoUrl](docs/SuggestsLogoUrl.md)
 - [SuggestsPositionItem](docs/SuggestsPositionItem.md)
 - [SuggestsPositions](docs/SuggestsPositions.md)
 - [SuggestsProfessionalRoleItem](docs/SuggestsProfessionalRoleItem.md)
 - [SuggestsProfessionalRoleItemWithName](docs/SuggestsProfessionalRoleItemWithName.md)
 - [SuggestsProfessionalRoles](docs/SuggestsProfessionalRoles.md)
 - [SuggestsSearchKeyword](docs/SuggestsSearchKeyword.md)
 - [SuggestsSearchKeywordItem](docs/SuggestsSearchKeywordItem.md)
 - [SuggestsSkillSet](docs/SuggestsSkillSet.md)
 - [SuggestsSpecializationsWithName](docs/SuggestsSpecializationsWithName.md)
 - [SuggestsVacancyPositionItem](docs/SuggestsVacancyPositionItem.md)
 - [SuggestsVacancyPositions](docs/SuggestsVacancyPositions.md)
 - [UserStatusesApplicant](docs/UserStatusesApplicant.md)
 - [UserStatusesJobSearchStatus](docs/UserStatusesJobSearchStatus.md)
 - [VacanciesActiveListItem](docs/VacanciesActiveListItem.md)
 - [VacanciesActiveListItems](docs/VacanciesActiveListItems.md)
 - [VacanciesActiveVacancyFields](docs/VacanciesActiveVacancyFields.md)
 - [VacanciesAdditionalVacancyFields](docs/VacanciesAdditionalVacancyFields.md)
 - [VacanciesAddress](docs/VacanciesAddress.md)
 - [VacanciesArchivedVacancyList](docs/VacanciesArchivedVacancyList.md)
 - [VacanciesArchivedVacancyListResponse](docs/VacanciesArchivedVacancyListResponse.md)
 - [VacanciesArgumentItem](docs/VacanciesArgumentItem.md)
 - [VacanciesArguments](docs/VacanciesArguments.md)
 - [VacanciesAvailableVacancyTypeItem](docs/VacanciesAvailableVacancyTypeItem.md)
 - [VacanciesAvailableVacancyTypeItemVacancyBillingType](docs/VacanciesAvailableVacancyTypeItemVacancyBillingType.md)
 - [VacanciesAvailableVacancyTypeResponse](docs/VacanciesAvailableVacancyTypeResponse.md)
 - [VacanciesClusterItem](docs/VacanciesClusterItem.md)
 - [VacanciesClusters](docs/VacanciesClusters.md)
 - [VacanciesDeletedVacancyListItems](docs/VacanciesDeletedVacancyListItems.md)
 - [VacanciesDeletedVacancyListResponse](docs/VacanciesDeletedVacancyListResponse.md)
 - [VacanciesEmployerPublic](docs/VacanciesEmployerPublic.md)
 - [VacanciesFixes](docs/VacanciesFixes.md)
 - [VacanciesItemsOfClusterItem](docs/VacanciesItemsOfClusterItem.md)
 - [VacanciesMatchListItem](docs/VacanciesMatchListItem.md)
 - [VacanciesMatchListItems](docs/VacanciesMatchListItems.md)
 - [VacanciesMatchVacancyFields](docs/VacanciesMatchVacancyFields.md)
 - [VacanciesNegotiationsVacancyShort](docs/VacanciesNegotiationsVacancyShort.md)
 - [VacanciesPreferredNegotiationsOrder](docs/VacanciesPreferredNegotiationsOrder.md)
 - [VacanciesStandardVacancyFields](docs/VacanciesStandardVacancyFields.md)
 - [VacanciesSuggests](docs/VacanciesSuggests.md)
 - [VacanciesUpgradeFieldsAction](docs/VacanciesUpgradeFieldsAction.md)
 - [VacanciesUpgradeFieldsBillingTypeFull](docs/VacanciesUpgradeFieldsBillingTypeFull.md)
 - [VacanciesUpgradeFieldsPrice](docs/VacanciesUpgradeFieldsPrice.md)
 - [VacanciesUpgradeFieldsWithoutAction](docs/VacanciesUpgradeFieldsWithoutAction.md)
 - [VacanciesVacanciesBlacklisted](docs/VacanciesVacanciesBlacklisted.md)
 - [VacanciesVacanciesBlacklistedItem](docs/VacanciesVacanciesBlacklistedItem.md)
 - [VacanciesVacanciesBlacklistedResponse](docs/VacanciesVacanciesBlacklistedResponse.md)
 - [VacanciesVacanciesFavorited](docs/VacanciesVacanciesFavorited.md)
 - [VacanciesVacanciesFavoritedItem](docs/VacanciesVacanciesFavoritedItem.md)
 - [VacanciesVacanciesFavoritedResponse](docs/VacanciesVacanciesFavoritedResponse.md)
 - [VacanciesVacanciesItem](docs/VacanciesVacanciesItem.md)
 - [VacanciesVacanciesItems](docs/VacanciesVacanciesItems.md)
 - [VacanciesVacanciesResponse](docs/VacanciesVacanciesResponse.md)
 - [VacanciesVacancy](docs/VacanciesVacancy.md)
 - [VacanciesVacancyArchived](docs/VacanciesVacancyArchived.md)
 - [VacanciesVacancyArchivedFields](docs/VacanciesVacancyArchivedFields.md)
 - [VacanciesVacancyCommonFields](docs/VacanciesVacancyCommonFields.md)
 - [VacanciesVacancyConditionFieldsAddressCondition](docs/VacanciesVacancyConditionFieldsAddressCondition.md)
 - [VacanciesVacancyConditionFieldsAddressFields](docs/VacanciesVacancyConditionFieldsAddressFields.md)
 - [VacanciesVacancyConditionFieldsCityCondition](docs/VacanciesVacancyConditionFieldsCityCondition.md)
 - [VacanciesVacancyConditionFieldsContactFields](docs/VacanciesVacancyConditionFieldsContactFields.md)
 - [VacanciesVacancyConditionFieldsContactsCondition](docs/VacanciesVacancyConditionFieldsContactsCondition.md)
 - [VacanciesVacancyConditionFieldsCountryCondition](docs/VacanciesVacancyConditionFieldsCountryCondition.md)
 - [VacanciesVacancyConditionFieldsFormattedCondition](docs/VacanciesVacancyConditionFieldsFormattedCondition.md)
 - [VacanciesVacancyConditionFieldsMaxMinCount](docs/VacanciesVacancyConditionFieldsMaxMinCount.md)
 - [VacanciesVacancyConditionFieldsMaxMinLength](docs/VacanciesVacancyConditionFieldsMaxMinLength.md)
 - [VacanciesVacancyConditionFieldsNumberCondition](docs/VacanciesVacancyConditionFieldsNumberCondition.md)
 - [VacanciesVacancyConditionFieldsPhoneCondition](docs/VacanciesVacancyConditionFieldsPhoneCondition.md)
 - [VacanciesVacancyConditionFieldsPhoneFields](docs/VacanciesVacancyConditionFieldsPhoneFields.md)
 - [VacanciesVacancyConditionFieldsRegexp](docs/VacanciesVacancyConditionFieldsRegexp.md)
 - [VacanciesVacancyConditionFieldsRequired](docs/VacanciesVacancyConditionFieldsRequired.md)
 - [VacanciesVacancyConditionFieldsRequiredCountWithTitle](docs/VacanciesVacancyConditionFieldsRequiredCountWithTitle.md)
 - [VacanciesVacancyConditionFieldsRequiredLengthWithTitle](docs/VacanciesVacancyConditionFieldsRequiredLengthWithTitle.md)
 - [VacanciesVacancyConditionFieldsRequiredWithTitle](docs/VacanciesVacancyConditionFieldsRequiredWithTitle.md)
 - [VacanciesVacancyConditionFieldsResponseUrlCondition](docs/VacanciesVacancyConditionFieldsResponseUrlCondition.md)
 - [VacanciesVacancyConditionFieldsSalaryCondition](docs/VacanciesVacancyConditionFieldsSalaryCondition.md)
 - [VacanciesVacancyConditionFieldsSalaryFields](docs/VacanciesVacancyConditionFieldsSalaryFields.md)
 - [VacanciesVacancyConditionFieldsTestCondition](docs/VacanciesVacancyConditionFieldsTestCondition.md)
 - [VacanciesVacancyConditionFieldsTestFields](docs/VacanciesVacancyConditionFieldsTestFields.md)
 - [VacanciesVacancyConditions](docs/VacanciesVacancyConditions.md)
 - [VacanciesVacancyEmployer](docs/VacanciesVacancyEmployer.md)
 - [VacanciesVacancyForApplicant](docs/VacanciesVacancyForApplicant.md)
 - [VacanciesVacancyForManager](docs/VacanciesVacancyForManager.md)
 - [VacanciesVacancyListResponse](docs/VacanciesVacancyListResponse.md)
 - [VacanciesVacancyManagerFields](docs/VacanciesVacancyManagerFields.md)
 - [VacanciesVacancyProlongate](docs/VacanciesVacancyProlongate.md)
 - [VacanciesVacancyProlongateActionsInner](docs/VacanciesVacancyProlongateActionsInner.md)
 - [VacanciesVacancyProlongateAvailableActions](docs/VacanciesVacancyProlongateAvailableActions.md)
 - [VacanciesVacancyProlongateUnavailableActions](docs/VacanciesVacancyProlongateUnavailableActions.md)
 - [VacanciesVacancyProlongateUnavailableActionsDisableReason](docs/VacanciesVacancyProlongateUnavailableActionsDisableReason.md)
 - [VacanciesVacancyShort](docs/VacanciesVacancyShort.md)
 - [VacanciesVacancyStatsItem](docs/VacanciesVacancyStatsItem.md)
 - [VacanciesVacancyStatsResponse](docs/VacanciesVacancyStatsResponse.md)
 - [VacanciesVacancyUpgradeListItem](docs/VacanciesVacancyUpgradeListItem.md)
 - [VacanciesVacancyUpgradeListItemWithoutActionInner](docs/VacanciesVacancyUpgradeListItemWithoutActionInner.md)
 - [VacanciesVacancyUpgradeListResponse](docs/VacanciesVacancyUpgradeListResponse.md)
 - [VacanciesVisitorsHiddenOnPage](docs/VacanciesVisitorsHiddenOnPage.md)
 - [VacanciesVisitorsResponse](docs/VacanciesVisitorsResponse.md)
 - [VacanciesVisitorsVisitorItems](docs/VacanciesVisitorsVisitorItems.md)
 - [VacanciesVisitorsVisitorItemsItemsInner](docs/VacanciesVisitorsVisitorItemsItemsInner.md)
 - [VacancyAddress](docs/VacancyAddress.md)
 - [VacancyAddressCommon](docs/VacancyAddressCommon.md)
 - [VacancyAddressOutput](docs/VacancyAddressOutput.md)
 - [VacancyAddressRawOutput](docs/VacancyAddressRawOutput.md)
 - [VacancyArea](docs/VacancyArea.md)
 - [VacancyAreaOutput](docs/VacancyAreaOutput.md)
 - [VacancyArguments](docs/VacancyArguments.md)
 - [VacancyBillingType](docs/VacancyBillingType.md)
 - [VacancyBillingTypeObject](docs/VacancyBillingTypeObject.md)
 - [VacancyBillingTypeOutput](docs/VacancyBillingTypeOutput.md)
 - [VacancyBrandedTemplate](docs/VacancyBrandedTemplate.md)
 - [VacancyCommonFields](docs/VacancyCommonFields.md)
 - [VacancyContacts](docs/VacancyContacts.md)
 - [VacancyContactsOutput](docs/VacancyContactsOutput.md)
 - [VacancyCounters](docs/VacancyCounters.md)
 - [VacancyCountersForActive](docs/VacancyCountersForActive.md)
 - [VacancyCountersForArchivedOrHidden](docs/VacancyCountersForArchivedOrHidden.md)
 - [VacancyCountersOutput](docs/VacancyCountersOutput.md)
 - [VacancyCreate](docs/VacancyCreate.md)
 - [VacancyCreateFields](docs/VacancyCreateFields.md)
 - [VacancyDepartment](docs/VacancyDepartment.md)
 - [VacancyDepartmentOutput](docs/VacancyDepartmentOutput.md)
 - [VacancyDescriptionAddress](docs/VacancyDescriptionAddress.md)
 - [VacancyDraftAddress](docs/VacancyDraftAddress.md)
 - [VacancyDraftAddressOutput](docs/VacancyDraftAddressOutput.md)
 - [VacancyDraftAssignedManager](docs/VacancyDraftAssignedManager.md)
 - [VacancyDraftAutoPublicationState](docs/VacancyDraftAutoPublicationState.md)
 - [VacancyDraftBillingType](docs/VacancyDraftBillingType.md)
 - [VacancyDraftBrandedTemplate](docs/VacancyDraftBrandedTemplate.md)
 - [VacancyDraftContacts](docs/VacancyDraftContacts.md)
 - [VacancyDraftContactsWithFullPhone](docs/VacancyDraftContactsWithFullPhone.md)
 - [VacancyDraftDraftResponseSchema](docs/VacancyDraftDraftResponseSchema.md)
 - [VacancyDraftDraftVacancyError](docs/VacancyDraftDraftVacancyError.md)
 - [VacancyDraftEmployer](docs/VacancyDraftEmployer.md)
 - [VacancyDraftEmployerLogoUrls](docs/VacancyDraftEmployerLogoUrls.md)
 - [VacancyDraftEmployment](docs/VacancyDraftEmployment.md)
 - [VacancyDraftKeySkillItem](docs/VacancyDraftKeySkillItem.md)
 - [VacancyDraftPhone](docs/VacancyDraftPhone.md)
 - [VacancyDraftPhoneFormatted](docs/VacancyDraftPhoneFormatted.md)
 - [VacancyDraftPhoneItem](docs/VacancyDraftPhoneItem.md)
 - [VacancyDraftPhoneItemWithFullPhone](docs/VacancyDraftPhoneItemWithFullPhone.md)
 - [VacancyDraftProfessionalRoleItem](docs/VacancyDraftProfessionalRoleItem.md)
 - [VacancyDraftPublications](docs/VacancyDraftPublications.md)
 - [VacancyDraftTest](docs/VacancyDraftTest.md)
 - [VacancyDraftType](docs/VacancyDraftType.md)
 - [VacancyDraftVacanciesDraftResponse](docs/VacancyDraftVacanciesDraftResponse.md)
 - [VacancyDraftVacancyDraftBase](docs/VacancyDraftVacancyDraftBase.md)
 - [VacancyDraftVacancyDraftBodyCommon](docs/VacancyDraftVacancyDraftBodyCommon.md)
 - [VacancyDraftVacancyDraftCommon](docs/VacancyDraftVacancyDraftCommon.md)
 - [VacancyDraftVacancyDraftCreate](docs/VacancyDraftVacancyDraftCreate.md)
 - [VacancyDraftVacancyDraftEdit](docs/VacancyDraftVacancyDraftEdit.md)
 - [VacancyDraftVacancyDraftFull](docs/VacancyDraftVacancyDraftFull.md)
 - [VacancyDraftVacancyDraftItem](docs/VacancyDraftVacancyDraftItem.md)
 - [VacancyDraftVacancyDraftItems](docs/VacancyDraftVacancyDraftItems.md)
 - [VacancyDriverLicenceTypeItem](docs/VacancyDriverLicenceTypeItem.md)
 - [VacancyDuplicates](docs/VacancyDuplicates.md)
 - [VacancyEdit](docs/VacancyEdit.md)
 - [VacancyEditBillingType](docs/VacancyEditBillingType.md)
 - [VacancyEditCommon](docs/VacancyEditCommon.md)
 - [VacancyEditFields](docs/VacancyEditFields.md)
 - [VacancyEditManager](docs/VacancyEditManager.md)
 - [VacancyEditManagerManager](docs/VacancyEditManagerManager.md)
 - [VacancyEmployerBlacklisted](docs/VacancyEmployerBlacklisted.md)
 - [VacancyEmployment](docs/VacancyEmployment.md)
 - [VacancyEmploymentOutput](docs/VacancyEmploymentOutput.md)
 - [VacancyExperience](docs/VacancyExperience.md)
 - [VacancyExperienceOutput](docs/VacancyExperienceOutput.md)
 - [VacancyHiddenDeprecated](docs/VacancyHiddenDeprecated.md)
 - [VacancyInsiderInterview](docs/VacancyInsiderInterview.md)
 - [VacancyKeySkillItem](docs/VacancyKeySkillItem.md)
 - [VacancyLanguage](docs/VacancyLanguage.md)
 - [VacancyLanguageLevel](docs/VacancyLanguageLevel.md)
 - [VacancyLanguageOutput](docs/VacancyLanguageOutput.md)
 - [VacancyLanguageOutputAllOfLevel](docs/VacancyLanguageOutputAllOfLevel.md)
 - [VacancyManager](docs/VacancyManager.md)
 - [VacancyManagerOutput](docs/VacancyManagerOutput.md)
 - [VacancyNegotiationActions](docs/VacancyNegotiationActions.md)
 - [VacancyPhoneItem](docs/VacancyPhoneItem.md)
 - [VacancyPicture](docs/VacancyPicture.md)
 - [VacancyProfessionalRoleItem](docs/VacancyProfessionalRoleItem.md)
 - [VacancyProfessionalRoleItemOutput](docs/VacancyProfessionalRoleItemOutput.md)
 - [VacancyPublication](docs/VacancyPublication.md)
 - [VacancyRelationItem](docs/VacancyRelationItem.md)
 - [VacancySalary](docs/VacancySalary.md)
 - [VacancySchedule](docs/VacancySchedule.md)
 - [VacancyScheduleOutput](docs/VacancyScheduleOutput.md)
 - [VacancySnippet](docs/VacancySnippet.md)
 - [VacancyTemplates](docs/VacancyTemplates.md)
 - [VacancyType](docs/VacancyType.md)
 - [VacancyTypeOutput](docs/VacancyTypeOutput.md)
 - [VacancyVacancyConstructorTemplate](docs/VacancyVacancyConstructorTemplate.md)
 - [VacancyVideoVacancy](docs/VacancyVideoVacancy.md)
 - [VacancyVideoVacancyCoverPicture](docs/VacancyVideoVacancyCoverPicture.md)
 - [VacancyWorkingDayItem](docs/VacancyWorkingDayItem.md)
 - [VacancyWorkingDayItemOutput](docs/VacancyWorkingDayItemOutput.md)
 - [VacancyWorkingTimeIntervalItem](docs/VacancyWorkingTimeIntervalItem.md)
 - [VacancyWorkingTimeIntervalItemOutput](docs/VacancyWorkingTimeIntervalItemOutput.md)
 - [VacancyWorkingTimeModeItem](docs/VacancyWorkingTimeModeItem.md)
 - [VacancyWorkingTimeModeItemOutput](docs/VacancyWorkingTimeModeItemOutput.md)
 - [WebhookActionNegotiationEmployerStateChange](docs/WebhookActionNegotiationEmployerStateChange.md)
 - [WebhookActionNewNegotiationVacancy](docs/WebhookActionNewNegotiationVacancy.md)
 - [WebhookActionNewResponseOrInvitationVacancy](docs/WebhookActionNewResponseOrInvitationVacancy.md)
 - [WebhookActionVacancyArchivation](docs/WebhookActionVacancyArchivation.md)
 - [WebhookActionVacancyChange](docs/WebhookActionVacancyChange.md)
 - [WebhookActionVacancyOnlyMineSettings](docs/WebhookActionVacancyOnlyMineSettings.md)
 - [WebhookActionVacancyProlongation](docs/WebhookActionVacancyProlongation.md)
 - [WebhookActionVacancyPublicationForVacancyManager](docs/WebhookActionVacancyPublicationForVacancyManager.md)
 - [WebhookBadDataError](docs/WebhookBadDataError.md)
 - [WebhookError](docs/WebhookError.md)
 - [WebhookErrors](docs/WebhookErrors.md)
 - [WebhookPayloadNegotiationEmployerStateChange](docs/WebhookPayloadNegotiationEmployerStateChange.md)
 - [WebhookPayloadNewNegotiationVacancy](docs/WebhookPayloadNewNegotiationVacancy.md)
 - [WebhookPayloadNewResponseOrInvitationVacancy](docs/WebhookPayloadNewResponseOrInvitationVacancy.md)
 - [WebhookPayloadVacancyArchivation](docs/WebhookPayloadVacancyArchivation.md)
 - [WebhookPayloadVacancyChange](docs/WebhookPayloadVacancyChange.md)
 - [WebhookPayloadVacancyProlongation](docs/WebhookPayloadVacancyProlongation.md)
 - [WebhookPayloadVacancyPublicationForVacancyManager](docs/WebhookPayloadVacancyPublicationForVacancyManager.md)
 - [WebhookSendObjectBaseUser](docs/WebhookSendObjectBaseUser.md)
 - [WebhookSendObjectBaseUserPayload](docs/WebhookSendObjectBaseUserPayload.md)
 - [WebhookSubscriptionCommonItem](docs/WebhookSubscriptionCommonItem.md)
 - [WebhookSubscriptionCommonItemActionsInner](docs/WebhookSubscriptionCommonItemActionsInner.md)
 - [WebhookSubscriptionCreate](docs/WebhookSubscriptionCreate.md)
 - [WebhookSubscriptionItem](docs/WebhookSubscriptionItem.md)
 - [WebhookSubscriptionUpdate](docs/WebhookSubscriptionUpdate.md)
 - [WebhookSubscriptionsOutput](docs/WebhookSubscriptionsOutput.md)
 - [WebhookSubscriptionsPostRequest](docs/WebhookSubscriptionsPostRequest.md)


## Documentation For Authorization


Authentication schemes defined for the API:
### OauthToken

- **Type**: HTTP Bearer token authentication

Example

```go
auth := context.WithValue(context.Background(), hh.ContextAccessToken, "BEARER_TOKEN_STRING")
r, err := client.Service.Operation(auth, args)
```


## Documentation for Utility Methods

Due to the fact that model structure members are all pointers, this package contains
a number of utility functions to easily obtain pointers to values of basic types.
Each of these functions takes a value of the given basic type and returns a pointer to it:

* `PtrBool`
* `PtrInt`
* `PtrInt32`
* `PtrInt64`
* `PtrFloat`
* `PtrFloat32`
* `PtrFloat64`
* `PtrString`
* `PtrTime`

## Author

api@hh.ru

