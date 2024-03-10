# SendNegotiationMessage404Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**RequestId** | **string** | Идентификатор запроса | 
**Errors** | [**[]ErrorsNegotiationSendMessageNotFoundError**](ErrorsNegotiationSendMessageNotFoundError.md) | Массив с данными ошибок | 
**Description** | Pointer to **string** | Описание типа ошибки | [optional] 

## Methods

### NewSendNegotiationMessage404Response

`func NewSendNegotiationMessage404Response(requestId string, errors []ErrorsNegotiationSendMessageNotFoundError, ) *SendNegotiationMessage404Response`

NewSendNegotiationMessage404Response instantiates a new SendNegotiationMessage404Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSendNegotiationMessage404ResponseWithDefaults

`func NewSendNegotiationMessage404ResponseWithDefaults() *SendNegotiationMessage404Response`

NewSendNegotiationMessage404ResponseWithDefaults instantiates a new SendNegotiationMessage404Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRequestId

`func (o *SendNegotiationMessage404Response) GetRequestId() string`

GetRequestId returns the RequestId field if non-nil, zero value otherwise.

### GetRequestIdOk

`func (o *SendNegotiationMessage404Response) GetRequestIdOk() (*string, bool)`

GetRequestIdOk returns a tuple with the RequestId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestId

`func (o *SendNegotiationMessage404Response) SetRequestId(v string)`

SetRequestId sets RequestId field to given value.


### GetErrors

`func (o *SendNegotiationMessage404Response) GetErrors() []ErrorsNegotiationSendMessageNotFoundError`

GetErrors returns the Errors field if non-nil, zero value otherwise.

### GetErrorsOk

`func (o *SendNegotiationMessage404Response) GetErrorsOk() (*[]ErrorsNegotiationSendMessageNotFoundError, bool)`

GetErrorsOk returns a tuple with the Errors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrors

`func (o *SendNegotiationMessage404Response) SetErrors(v []ErrorsNegotiationSendMessageNotFoundError)`

SetErrors sets Errors field to given value.


### GetDescription

`func (o *SendNegotiationMessage404Response) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *SendNegotiationMessage404Response) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *SendNegotiationMessage404Response) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *SendNegotiationMessage404Response) HasDescription() bool`

HasDescription returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


