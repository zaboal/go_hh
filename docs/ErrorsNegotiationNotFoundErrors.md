# ErrorsNegotiationNotFoundErrors

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**RequestId** | **string** | Идентификатор запроса | 
**Description** | Pointer to **string** | Описание типа ошибки | [optional] 
**Errors** | [**[]ErrorsNegotiationNotFoundError**](ErrorsNegotiationNotFoundError.md) | Массив с данными ошибок | 

## Methods

### NewErrorsNegotiationNotFoundErrors

`func NewErrorsNegotiationNotFoundErrors(requestId string, errors []ErrorsNegotiationNotFoundError, ) *ErrorsNegotiationNotFoundErrors`

NewErrorsNegotiationNotFoundErrors instantiates a new ErrorsNegotiationNotFoundErrors object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewErrorsNegotiationNotFoundErrorsWithDefaults

`func NewErrorsNegotiationNotFoundErrorsWithDefaults() *ErrorsNegotiationNotFoundErrors`

NewErrorsNegotiationNotFoundErrorsWithDefaults instantiates a new ErrorsNegotiationNotFoundErrors object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRequestId

`func (o *ErrorsNegotiationNotFoundErrors) GetRequestId() string`

GetRequestId returns the RequestId field if non-nil, zero value otherwise.

### GetRequestIdOk

`func (o *ErrorsNegotiationNotFoundErrors) GetRequestIdOk() (*string, bool)`

GetRequestIdOk returns a tuple with the RequestId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestId

`func (o *ErrorsNegotiationNotFoundErrors) SetRequestId(v string)`

SetRequestId sets RequestId field to given value.


### GetDescription

`func (o *ErrorsNegotiationNotFoundErrors) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *ErrorsNegotiationNotFoundErrors) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *ErrorsNegotiationNotFoundErrors) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *ErrorsNegotiationNotFoundErrors) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetErrors

`func (o *ErrorsNegotiationNotFoundErrors) GetErrors() []ErrorsNegotiationNotFoundError`

GetErrors returns the Errors field if non-nil, zero value otherwise.

### GetErrorsOk

`func (o *ErrorsNegotiationNotFoundErrors) GetErrorsOk() (*[]ErrorsNegotiationNotFoundError, bool)`

GetErrorsOk returns a tuple with the Errors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrors

`func (o *ErrorsNegotiationNotFoundErrors) SetErrors(v []ErrorsNegotiationNotFoundError)`

SetErrors sets Errors field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


