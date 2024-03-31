# ErrorsNegotiationNegotiationsNotFoundErrors

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**RequestId** | **string** | Идентификатор запроса | 
**Errors** | [**[]ErrorsNegotiationNotFoundError**](ErrorsNegotiationNotFoundError.md) | Массив с данными ошибок | 
**Description** | Pointer to **string** | Описание типа ошибки | [optional] 

## Methods

### NewErrorsNegotiationNegotiationsNotFoundErrors

`func NewErrorsNegotiationNegotiationsNotFoundErrors(requestId string, errors []ErrorsNegotiationNotFoundError, ) *ErrorsNegotiationNegotiationsNotFoundErrors`

NewErrorsNegotiationNegotiationsNotFoundErrors instantiates a new ErrorsNegotiationNegotiationsNotFoundErrors object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewErrorsNegotiationNegotiationsNotFoundErrorsWithDefaults

`func NewErrorsNegotiationNegotiationsNotFoundErrorsWithDefaults() *ErrorsNegotiationNegotiationsNotFoundErrors`

NewErrorsNegotiationNegotiationsNotFoundErrorsWithDefaults instantiates a new ErrorsNegotiationNegotiationsNotFoundErrors object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRequestId

`func (o *ErrorsNegotiationNegotiationsNotFoundErrors) GetRequestId() string`

GetRequestId returns the RequestId field if non-nil, zero value otherwise.

### GetRequestIdOk

`func (o *ErrorsNegotiationNegotiationsNotFoundErrors) GetRequestIdOk() (*string, bool)`

GetRequestIdOk returns a tuple with the RequestId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestId

`func (o *ErrorsNegotiationNegotiationsNotFoundErrors) SetRequestId(v string)`

SetRequestId sets RequestId field to given value.


### GetErrors

`func (o *ErrorsNegotiationNegotiationsNotFoundErrors) GetErrors() []ErrorsNegotiationNotFoundError`

GetErrors returns the Errors field if non-nil, zero value otherwise.

### GetErrorsOk

`func (o *ErrorsNegotiationNegotiationsNotFoundErrors) GetErrorsOk() (*[]ErrorsNegotiationNotFoundError, bool)`

GetErrorsOk returns a tuple with the Errors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrors

`func (o *ErrorsNegotiationNegotiationsNotFoundErrors) SetErrors(v []ErrorsNegotiationNotFoundError)`

SetErrors sets Errors field to given value.


### GetDescription

`func (o *ErrorsNegotiationNegotiationsNotFoundErrors) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *ErrorsNegotiationNegotiationsNotFoundErrors) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *ErrorsNegotiationNegotiationsNotFoundErrors) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *ErrorsNegotiationNegotiationsNotFoundErrors) HasDescription() bool`

HasDescription returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


