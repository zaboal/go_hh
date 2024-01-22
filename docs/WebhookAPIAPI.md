# \WebhookAPIAPI

All URIs are relative to *https://api.hh.ru*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CancelWebhookSubscription**](WebhookAPIAPI.md#CancelWebhookSubscription) | **Delete** /webhook/subscriptions/{subscription_id} | Удалить подписку на уведомление
[**ChangeWebhookSubscription**](WebhookAPIAPI.md#ChangeWebhookSubscription) | **Put** /webhook/subscriptions/{subscription_id} | Изменить подписку на уведомления
[**GetWebhookSubscriptions**](WebhookAPIAPI.md#GetWebhookSubscriptions) | **Get** /webhook/subscriptions | Получить список уведомлений, на которые подписан пользователь
[**PostWebhookSubscription**](WebhookAPIAPI.md#PostWebhookSubscription) | **Post** /webhook/subscriptions | Подписаться на уведомления



## CancelWebhookSubscription

> CancelWebhookSubscription(ctx, subscriptionId).HHUserAgent(hHUserAgent).Locale(locale).Host(host).Execute()

Удалить подписку на уведомление

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/zaboal/hh-go"
)

func main() {
	subscriptionId := "1455" // string | Идентификатор подписки
	hHUserAgent := "MyApp/1.0 (my-app-feedback@example.com)" // string | Название приложения и контактная почта разработчика (см. [Информация о клиенте](https://github.com/hhru/api/blob/master/docs/general.md#%D1%82%D1%80%D0%B5%D0%B1%D0%BE%D0%B2%D0%B0%D0%BD%D0%B8%D1%8F-%D0%BA-%D0%B7%D0%B0%D0%BF%D1%80%D0%BE%D1%81%D0%B0%D0%BC)) 
	locale := "EN" // string | Идентификатор локали (см. [Локализация](#tag/Obshie-spravochniki/operation/get-locales))  (optional) (default to "RU")
	host := "hh.ru" // string | Доменное имя сайта (см. [Выбор сайта](https://github.com/hhru/api/blob/master/docs/hosts.md))  (optional) (default to "hh.ru")

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.WebhookAPIAPI.CancelWebhookSubscription(context.Background(), subscriptionId).HHUserAgent(hHUserAgent).Locale(locale).Host(host).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WebhookAPIAPI.CancelWebhookSubscription``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**subscriptionId** | **string** | Идентификатор подписки | 

### Other Parameters

Other parameters are passed through a pointer to a apiCancelWebhookSubscriptionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **hHUserAgent** | **string** | Название приложения и контактная почта разработчика (см. [Информация о клиенте](https://github.com/hhru/api/blob/master/docs/general.md#%D1%82%D1%80%D0%B5%D0%B1%D0%BE%D0%B2%D0%B0%D0%BD%D0%B8%D1%8F-%D0%BA-%D0%B7%D0%B0%D0%BF%D1%80%D0%BE%D1%81%D0%B0%D0%BC))  | 
 **locale** | **string** | Идентификатор локали (см. [Локализация](#tag/Obshie-spravochniki/operation/get-locales))  | [default to &quot;RU&quot;]
 **host** | **string** | Доменное имя сайта (см. [Выбор сайта](https://github.com/hhru/api/blob/master/docs/hosts.md))  | [default to &quot;hh.ru&quot;]

### Return type

 (empty response body)

### Authorization

[OauthToken](../README.md#OauthToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ChangeWebhookSubscription

> ChangeWebhookSubscription(ctx, subscriptionId).HHUserAgent(hHUserAgent).WebhookSubscriptionUpdate(webhookSubscriptionUpdate).Locale(locale).Host(host).Execute()

Изменить подписку на уведомления



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/zaboal/hh-go"
)

func main() {
	subscriptionId := "1455" // string | Идентификатор подписки
	hHUserAgent := "MyApp/1.0 (my-app-feedback@example.com)" // string | Название приложения и контактная почта разработчика (см. [Информация о клиенте](https://github.com/hhru/api/blob/master/docs/general.md#%D1%82%D1%80%D0%B5%D0%B1%D0%BE%D0%B2%D0%B0%D0%BD%D0%B8%D1%8F-%D0%BA-%D0%B7%D0%B0%D0%BF%D1%80%D0%BE%D1%81%D0%B0%D0%BC)) 
	webhookSubscriptionUpdate := *openapiclient.NewWebhookSubscriptionUpdate() // WebhookSubscriptionUpdate | 
	locale := "EN" // string | Идентификатор локали (см. [Локализация](#tag/Obshie-spravochniki/operation/get-locales))  (optional) (default to "RU")
	host := "hh.ru" // string | Доменное имя сайта (см. [Выбор сайта](https://github.com/hhru/api/blob/master/docs/hosts.md))  (optional) (default to "hh.ru")

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.WebhookAPIAPI.ChangeWebhookSubscription(context.Background(), subscriptionId).HHUserAgent(hHUserAgent).WebhookSubscriptionUpdate(webhookSubscriptionUpdate).Locale(locale).Host(host).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WebhookAPIAPI.ChangeWebhookSubscription``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**subscriptionId** | **string** | Идентификатор подписки | 

### Other Parameters

Other parameters are passed through a pointer to a apiChangeWebhookSubscriptionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **hHUserAgent** | **string** | Название приложения и контактная почта разработчика (см. [Информация о клиенте](https://github.com/hhru/api/blob/master/docs/general.md#%D1%82%D1%80%D0%B5%D0%B1%D0%BE%D0%B2%D0%B0%D0%BD%D0%B8%D1%8F-%D0%BA-%D0%B7%D0%B0%D0%BF%D1%80%D0%BE%D1%81%D0%B0%D0%BC))  | 
 **webhookSubscriptionUpdate** | [**WebhookSubscriptionUpdate**](WebhookSubscriptionUpdate.md) |  | 
 **locale** | **string** | Идентификатор локали (см. [Локализация](#tag/Obshie-spravochniki/operation/get-locales))  | [default to &quot;RU&quot;]
 **host** | **string** | Доменное имя сайта (см. [Выбор сайта](https://github.com/hhru/api/blob/master/docs/hosts.md))  | [default to &quot;hh.ru&quot;]

### Return type

 (empty response body)

### Authorization

[OauthToken](../README.md#OauthToken)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetWebhookSubscriptions

> WebhookSubscriptionsOutput GetWebhookSubscriptions(ctx).HHUserAgent(hHUserAgent).Locale(locale).Host(host).Execute()

Получить список уведомлений, на которые подписан пользователь

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/zaboal/hh-go"
)

func main() {
	hHUserAgent := "MyApp/1.0 (my-app-feedback@example.com)" // string | Название приложения и контактная почта разработчика (см. [Информация о клиенте](https://github.com/hhru/api/blob/master/docs/general.md#%D1%82%D1%80%D0%B5%D0%B1%D0%BE%D0%B2%D0%B0%D0%BD%D0%B8%D1%8F-%D0%BA-%D0%B7%D0%B0%D0%BF%D1%80%D0%BE%D1%81%D0%B0%D0%BC)) 
	locale := "EN" // string | Идентификатор локали (см. [Локализация](#tag/Obshie-spravochniki/operation/get-locales))  (optional) (default to "RU")
	host := "hh.ru" // string | Доменное имя сайта (см. [Выбор сайта](https://github.com/hhru/api/blob/master/docs/hosts.md))  (optional) (default to "hh.ru")

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.WebhookAPIAPI.GetWebhookSubscriptions(context.Background()).HHUserAgent(hHUserAgent).Locale(locale).Host(host).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WebhookAPIAPI.GetWebhookSubscriptions``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetWebhookSubscriptions`: WebhookSubscriptionsOutput
	fmt.Fprintf(os.Stdout, "Response from `WebhookAPIAPI.GetWebhookSubscriptions`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetWebhookSubscriptionsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **hHUserAgent** | **string** | Название приложения и контактная почта разработчика (см. [Информация о клиенте](https://github.com/hhru/api/blob/master/docs/general.md#%D1%82%D1%80%D0%B5%D0%B1%D0%BE%D0%B2%D0%B0%D0%BD%D0%B8%D1%8F-%D0%BA-%D0%B7%D0%B0%D0%BF%D1%80%D0%BE%D1%81%D0%B0%D0%BC))  | 
 **locale** | **string** | Идентификатор локали (см. [Локализация](#tag/Obshie-spravochniki/operation/get-locales))  | [default to &quot;RU&quot;]
 **host** | **string** | Доменное имя сайта (см. [Выбор сайта](https://github.com/hhru/api/blob/master/docs/hosts.md))  | [default to &quot;hh.ru&quot;]

### Return type

[**WebhookSubscriptionsOutput**](WebhookSubscriptionsOutput.md)

### Authorization

[OauthToken](../README.md#OauthToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostWebhookSubscription

> Id PostWebhookSubscription(ctx).HHUserAgent(hHUserAgent).WebhookSubscriptionCreate(webhookSubscriptionCreate).Locale(locale).Host(host).Execute()

Подписаться на уведомления



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/zaboal/hh-go"
)

func main() {
	hHUserAgent := "MyApp/1.0 (my-app-feedback@example.com)" // string | Название приложения и контактная почта разработчика (см. [Информация о клиенте](https://github.com/hhru/api/blob/master/docs/general.md#%D1%82%D1%80%D0%B5%D0%B1%D0%BE%D0%B2%D0%B0%D0%BD%D0%B8%D1%8F-%D0%BA-%D0%B7%D0%B0%D0%BF%D1%80%D0%BE%D1%81%D0%B0%D0%BC)) 
	webhookSubscriptionCreate := *openapiclient.NewWebhookSubscriptionCreate([]openapiclient.WebhookSubscriptionCommonItemActionsInner{*openapiclient.NewWebhookSubscriptionCommonItemActionsInner("Type_example")}, "http://www.example.com") // WebhookSubscriptionCreate | 
	locale := "EN" // string | Идентификатор локали (см. [Локализация](#tag/Obshie-spravochniki/operation/get-locales))  (optional) (default to "RU")
	host := "hh.ru" // string | Доменное имя сайта (см. [Выбор сайта](https://github.com/hhru/api/blob/master/docs/hosts.md))  (optional) (default to "hh.ru")

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.WebhookAPIAPI.PostWebhookSubscription(context.Background()).HHUserAgent(hHUserAgent).WebhookSubscriptionCreate(webhookSubscriptionCreate).Locale(locale).Host(host).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WebhookAPIAPI.PostWebhookSubscription``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostWebhookSubscription`: Id
	fmt.Fprintf(os.Stdout, "Response from `WebhookAPIAPI.PostWebhookSubscription`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPostWebhookSubscriptionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **hHUserAgent** | **string** | Название приложения и контактная почта разработчика (см. [Информация о клиенте](https://github.com/hhru/api/blob/master/docs/general.md#%D1%82%D1%80%D0%B5%D0%B1%D0%BE%D0%B2%D0%B0%D0%BD%D0%B8%D1%8F-%D0%BA-%D0%B7%D0%B0%D0%BF%D1%80%D0%BE%D1%81%D0%B0%D0%BC))  | 
 **webhookSubscriptionCreate** | [**WebhookSubscriptionCreate**](WebhookSubscriptionCreate.md) |  | 
 **locale** | **string** | Идентификатор локали (см. [Локализация](#tag/Obshie-spravochniki/operation/get-locales))  | [default to &quot;RU&quot;]
 **host** | **string** | Доменное имя сайта (см. [Выбор сайта](https://github.com/hhru/api/blob/master/docs/hosts.md))  | [default to &quot;hh.ru&quot;]

### Return type

[**Id**](Id.md)

### Authorization

[OauthToken](../README.md#OauthToken)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

