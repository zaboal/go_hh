# ErrorsResumeTooManyRequestsErrors

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**RequestId** | **string** | Идентификатор запроса | 
**Errors** | [**[]ErrorsResumeTooManyRequestsError**](ErrorsResumeTooManyRequestsError.md) | Массив с данными ошибок | 

## Methods

### NewErrorsResumeTooManyRequestsErrors

`func NewErrorsResumeTooManyRequestsErrors(requestId string, errors []ErrorsResumeTooManyRequestsError, ) *ErrorsResumeTooManyRequestsErrors`

NewErrorsResumeTooManyRequestsErrors instantiates a new ErrorsResumeTooManyRequestsErrors object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewErrorsResumeTooManyRequestsErrorsWithDefaults

`func NewErrorsResumeTooManyRequestsErrorsWithDefaults() *ErrorsResumeTooManyRequestsErrors`

NewErrorsResumeTooManyRequestsErrorsWithDefaults instantiates a new ErrorsResumeTooManyRequestsErrors object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRequestId

`func (o *ErrorsResumeTooManyRequestsErrors) GetRequestId() string`

GetRequestId returns the RequestId field if non-nil, zero value otherwise.

### GetRequestIdOk

`func (o *ErrorsResumeTooManyRequestsErrors) GetRequestIdOk() (*string, bool)`

GetRequestIdOk returns a tuple with the RequestId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestId

`func (o *ErrorsResumeTooManyRequestsErrors) SetRequestId(v string)`

SetRequestId sets RequestId field to given value.


### GetErrors

`func (o *ErrorsResumeTooManyRequestsErrors) GetErrors() []ErrorsResumeTooManyRequestsError`

GetErrors returns the Errors field if non-nil, zero value otherwise.

### GetErrorsOk

`func (o *ErrorsResumeTooManyRequestsErrors) GetErrorsOk() (*[]ErrorsResumeTooManyRequestsError, bool)`

GetErrorsOk returns a tuple with the Errors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrors

`func (o *ErrorsResumeTooManyRequestsErrors) SetErrors(v []ErrorsResumeTooManyRequestsError)`

SetErrors sets Errors field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


