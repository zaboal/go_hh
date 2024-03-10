# ErrorsNegotiationSendMessageNotFoundErrors

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**RequestId** | **string** | Идентификатор запроса | 
**Description** | Pointer to **string** | Описание типа ошибки | [optional] 
**Errors** | [**[]ErrorsNegotiationSendMessageNotFoundError**](ErrorsNegotiationSendMessageNotFoundError.md) | Массив с данными ошибок | 

## Methods

### NewErrorsNegotiationSendMessageNotFoundErrors

`func NewErrorsNegotiationSendMessageNotFoundErrors(requestId string, errors []ErrorsNegotiationSendMessageNotFoundError, ) *ErrorsNegotiationSendMessageNotFoundErrors`

NewErrorsNegotiationSendMessageNotFoundErrors instantiates a new ErrorsNegotiationSendMessageNotFoundErrors object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewErrorsNegotiationSendMessageNotFoundErrorsWithDefaults

`func NewErrorsNegotiationSendMessageNotFoundErrorsWithDefaults() *ErrorsNegotiationSendMessageNotFoundErrors`

NewErrorsNegotiationSendMessageNotFoundErrorsWithDefaults instantiates a new ErrorsNegotiationSendMessageNotFoundErrors object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRequestId

`func (o *ErrorsNegotiationSendMessageNotFoundErrors) GetRequestId() string`

GetRequestId returns the RequestId field if non-nil, zero value otherwise.

### GetRequestIdOk

`func (o *ErrorsNegotiationSendMessageNotFoundErrors) GetRequestIdOk() (*string, bool)`

GetRequestIdOk returns a tuple with the RequestId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestId

`func (o *ErrorsNegotiationSendMessageNotFoundErrors) SetRequestId(v string)`

SetRequestId sets RequestId field to given value.


### GetDescription

`func (o *ErrorsNegotiationSendMessageNotFoundErrors) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *ErrorsNegotiationSendMessageNotFoundErrors) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *ErrorsNegotiationSendMessageNotFoundErrors) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *ErrorsNegotiationSendMessageNotFoundErrors) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetErrors

`func (o *ErrorsNegotiationSendMessageNotFoundErrors) GetErrors() []ErrorsNegotiationSendMessageNotFoundError`

GetErrors returns the Errors field if non-nil, zero value otherwise.

### GetErrorsOk

`func (o *ErrorsNegotiationSendMessageNotFoundErrors) GetErrorsOk() (*[]ErrorsNegotiationSendMessageNotFoundError, bool)`

GetErrorsOk returns a tuple with the Errors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrors

`func (o *ErrorsNegotiationSendMessageNotFoundErrors) SetErrors(v []ErrorsNegotiationSendMessageNotFoundError)`

SetErrors sets Errors field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


