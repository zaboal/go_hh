# WebhookBadDataError

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**RequestId** | **string** | Идентификатор запроса | 
**Errors** | [**[]WebhookError**](WebhookError.md) | Массив с данными ошибок | 

## Methods

### NewWebhookBadDataError

`func NewWebhookBadDataError(requestId string, errors []WebhookError, ) *WebhookBadDataError`

NewWebhookBadDataError instantiates a new WebhookBadDataError object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWebhookBadDataErrorWithDefaults

`func NewWebhookBadDataErrorWithDefaults() *WebhookBadDataError`

NewWebhookBadDataErrorWithDefaults instantiates a new WebhookBadDataError object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRequestId

`func (o *WebhookBadDataError) GetRequestId() string`

GetRequestId returns the RequestId field if non-nil, zero value otherwise.

### GetRequestIdOk

`func (o *WebhookBadDataError) GetRequestIdOk() (*string, bool)`

GetRequestIdOk returns a tuple with the RequestId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestId

`func (o *WebhookBadDataError) SetRequestId(v string)`

SetRequestId sets RequestId field to given value.


### GetErrors

`func (o *WebhookBadDataError) GetErrors() []WebhookError`

GetErrors returns the Errors field if non-nil, zero value otherwise.

### GetErrorsOk

`func (o *WebhookBadDataError) GetErrorsOk() (*[]WebhookError, bool)`

GetErrorsOk returns a tuple with the Errors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrors

`func (o *WebhookBadDataError) SetErrors(v []WebhookError)`

SetErrors sets Errors field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


