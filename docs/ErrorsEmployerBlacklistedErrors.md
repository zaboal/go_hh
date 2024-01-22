# ErrorsEmployerBlacklistedErrors

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**RequestId** | **string** | Идентификатор запроса | 
**Description** | Pointer to **string** | Описание ошибки | [optional] 
**Errors** | [**[]ErrorsEmployerBlacklistedError**](ErrorsEmployerBlacklistedError.md) | Массив с данными ошибок | 

## Methods

### NewErrorsEmployerBlacklistedErrors

`func NewErrorsEmployerBlacklistedErrors(requestId string, errors []ErrorsEmployerBlacklistedError, ) *ErrorsEmployerBlacklistedErrors`

NewErrorsEmployerBlacklistedErrors instantiates a new ErrorsEmployerBlacklistedErrors object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewErrorsEmployerBlacklistedErrorsWithDefaults

`func NewErrorsEmployerBlacklistedErrorsWithDefaults() *ErrorsEmployerBlacklistedErrors`

NewErrorsEmployerBlacklistedErrorsWithDefaults instantiates a new ErrorsEmployerBlacklistedErrors object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRequestId

`func (o *ErrorsEmployerBlacklistedErrors) GetRequestId() string`

GetRequestId returns the RequestId field if non-nil, zero value otherwise.

### GetRequestIdOk

`func (o *ErrorsEmployerBlacklistedErrors) GetRequestIdOk() (*string, bool)`

GetRequestIdOk returns a tuple with the RequestId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestId

`func (o *ErrorsEmployerBlacklistedErrors) SetRequestId(v string)`

SetRequestId sets RequestId field to given value.


### GetDescription

`func (o *ErrorsEmployerBlacklistedErrors) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *ErrorsEmployerBlacklistedErrors) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *ErrorsEmployerBlacklistedErrors) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *ErrorsEmployerBlacklistedErrors) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetErrors

`func (o *ErrorsEmployerBlacklistedErrors) GetErrors() []ErrorsEmployerBlacklistedError`

GetErrors returns the Errors field if non-nil, zero value otherwise.

### GetErrorsOk

`func (o *ErrorsEmployerBlacklistedErrors) GetErrorsOk() (*[]ErrorsEmployerBlacklistedError, bool)`

GetErrorsOk returns a tuple with the Errors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrors

`func (o *ErrorsEmployerBlacklistedErrors) SetErrors(v []ErrorsEmployerBlacklistedError)`

SetErrors sets Errors field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


