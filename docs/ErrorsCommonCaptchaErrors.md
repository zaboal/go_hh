# ErrorsCommonCaptchaErrors

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**RequestId** | **string** | Идентификатор запроса | 
**Description** | Pointer to **string** | Описание ошибки | [optional] 
**Errors** | [**[]ErrorsCommonCaptchaError**](ErrorsCommonCaptchaError.md) | Массив с данными ошибок | 

## Methods

### NewErrorsCommonCaptchaErrors

`func NewErrorsCommonCaptchaErrors(requestId string, errors []ErrorsCommonCaptchaError, ) *ErrorsCommonCaptchaErrors`

NewErrorsCommonCaptchaErrors instantiates a new ErrorsCommonCaptchaErrors object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewErrorsCommonCaptchaErrorsWithDefaults

`func NewErrorsCommonCaptchaErrorsWithDefaults() *ErrorsCommonCaptchaErrors`

NewErrorsCommonCaptchaErrorsWithDefaults instantiates a new ErrorsCommonCaptchaErrors object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRequestId

`func (o *ErrorsCommonCaptchaErrors) GetRequestId() string`

GetRequestId returns the RequestId field if non-nil, zero value otherwise.

### GetRequestIdOk

`func (o *ErrorsCommonCaptchaErrors) GetRequestIdOk() (*string, bool)`

GetRequestIdOk returns a tuple with the RequestId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestId

`func (o *ErrorsCommonCaptchaErrors) SetRequestId(v string)`

SetRequestId sets RequestId field to given value.


### GetDescription

`func (o *ErrorsCommonCaptchaErrors) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *ErrorsCommonCaptchaErrors) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *ErrorsCommonCaptchaErrors) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *ErrorsCommonCaptchaErrors) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetErrors

`func (o *ErrorsCommonCaptchaErrors) GetErrors() []ErrorsCommonCaptchaError`

GetErrors returns the Errors field if non-nil, zero value otherwise.

### GetErrorsOk

`func (o *ErrorsCommonCaptchaErrors) GetErrorsOk() (*[]ErrorsCommonCaptchaError, bool)`

GetErrorsOk returns a tuple with the Errors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrors

`func (o *ErrorsCommonCaptchaErrors) SetErrors(v []ErrorsCommonCaptchaError)`

SetErrors sets Errors field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


