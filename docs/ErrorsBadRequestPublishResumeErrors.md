# ErrorsBadRequestPublishResumeErrors

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**RequestId** | **string** | Идентификатор запроса | 
**Description** | Pointer to **string** | Описание причины ошибки. Возможные причины:  * Не заполнены обязательные поля.    Чтобы понять, какие именно поля не заполнены, воспользуйтесь методом [просмотр резюме](#tag/Rezyume.-Prosmotr-informacii/operation/get-resume). Обязательные поля перечислены в поле &#x60;progress.mandatory&#x60;).  * Не отредактированы поля после блокировки модератором. * Резюме находится на проверке у модератора  | [optional] 
**Errors** | [**[]ErrorsBadRequestPublishResumeError**](ErrorsBadRequestPublishResumeError.md) | Массив с данными ошибок | 

## Methods

### NewErrorsBadRequestPublishResumeErrors

`func NewErrorsBadRequestPublishResumeErrors(requestId string, errors []ErrorsBadRequestPublishResumeError, ) *ErrorsBadRequestPublishResumeErrors`

NewErrorsBadRequestPublishResumeErrors instantiates a new ErrorsBadRequestPublishResumeErrors object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewErrorsBadRequestPublishResumeErrorsWithDefaults

`func NewErrorsBadRequestPublishResumeErrorsWithDefaults() *ErrorsBadRequestPublishResumeErrors`

NewErrorsBadRequestPublishResumeErrorsWithDefaults instantiates a new ErrorsBadRequestPublishResumeErrors object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRequestId

`func (o *ErrorsBadRequestPublishResumeErrors) GetRequestId() string`

GetRequestId returns the RequestId field if non-nil, zero value otherwise.

### GetRequestIdOk

`func (o *ErrorsBadRequestPublishResumeErrors) GetRequestIdOk() (*string, bool)`

GetRequestIdOk returns a tuple with the RequestId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestId

`func (o *ErrorsBadRequestPublishResumeErrors) SetRequestId(v string)`

SetRequestId sets RequestId field to given value.


### GetDescription

`func (o *ErrorsBadRequestPublishResumeErrors) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *ErrorsBadRequestPublishResumeErrors) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *ErrorsBadRequestPublishResumeErrors) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *ErrorsBadRequestPublishResumeErrors) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetErrors

`func (o *ErrorsBadRequestPublishResumeErrors) GetErrors() []ErrorsBadRequestPublishResumeError`

GetErrors returns the Errors field if non-nil, zero value otherwise.

### GetErrorsOk

`func (o *ErrorsBadRequestPublishResumeErrors) GetErrorsOk() (*[]ErrorsBadRequestPublishResumeError, bool)`

GetErrorsOk returns a tuple with the Errors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrors

`func (o *ErrorsBadRequestPublishResumeErrors) SetErrors(v []ErrorsBadRequestPublishResumeError)`

SetErrors sets Errors field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


