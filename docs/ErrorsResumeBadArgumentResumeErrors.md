# ErrorsResumeBadArgumentResumeErrors

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**RequestId** | **string** | Идентификатор запроса | 
**Errors** | [**[]ErrorsResumeAddEditError**](ErrorsResumeAddEditError.md) | Массив с данными ошибок | 

## Methods

### NewErrorsResumeBadArgumentResumeErrors

`func NewErrorsResumeBadArgumentResumeErrors(requestId string, errors []ErrorsResumeAddEditError, ) *ErrorsResumeBadArgumentResumeErrors`

NewErrorsResumeBadArgumentResumeErrors instantiates a new ErrorsResumeBadArgumentResumeErrors object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewErrorsResumeBadArgumentResumeErrorsWithDefaults

`func NewErrorsResumeBadArgumentResumeErrorsWithDefaults() *ErrorsResumeBadArgumentResumeErrors`

NewErrorsResumeBadArgumentResumeErrorsWithDefaults instantiates a new ErrorsResumeBadArgumentResumeErrors object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRequestId

`func (o *ErrorsResumeBadArgumentResumeErrors) GetRequestId() string`

GetRequestId returns the RequestId field if non-nil, zero value otherwise.

### GetRequestIdOk

`func (o *ErrorsResumeBadArgumentResumeErrors) GetRequestIdOk() (*string, bool)`

GetRequestIdOk returns a tuple with the RequestId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestId

`func (o *ErrorsResumeBadArgumentResumeErrors) SetRequestId(v string)`

SetRequestId sets RequestId field to given value.


### GetErrors

`func (o *ErrorsResumeBadArgumentResumeErrors) GetErrors() []ErrorsResumeAddEditError`

GetErrors returns the Errors field if non-nil, zero value otherwise.

### GetErrorsOk

`func (o *ErrorsResumeBadArgumentResumeErrors) GetErrorsOk() (*[]ErrorsResumeAddEditError, bool)`

GetErrorsOk returns a tuple with the Errors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrors

`func (o *ErrorsResumeBadArgumentResumeErrors) SetErrors(v []ErrorsResumeAddEditError)`

SetErrors sets Errors field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


