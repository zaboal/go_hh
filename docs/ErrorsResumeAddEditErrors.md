# ErrorsResumeAddEditErrors

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**RequestId** | **string** | Идентификатор запроса | 
**Errors** | [**[]ErrorsResumeAddEditError**](ErrorsResumeAddEditError.md) | Массив с данными ошибок | 

## Methods

### NewErrorsResumeAddEditErrors

`func NewErrorsResumeAddEditErrors(requestId string, errors []ErrorsResumeAddEditError, ) *ErrorsResumeAddEditErrors`

NewErrorsResumeAddEditErrors instantiates a new ErrorsResumeAddEditErrors object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewErrorsResumeAddEditErrorsWithDefaults

`func NewErrorsResumeAddEditErrorsWithDefaults() *ErrorsResumeAddEditErrors`

NewErrorsResumeAddEditErrorsWithDefaults instantiates a new ErrorsResumeAddEditErrors object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRequestId

`func (o *ErrorsResumeAddEditErrors) GetRequestId() string`

GetRequestId returns the RequestId field if non-nil, zero value otherwise.

### GetRequestIdOk

`func (o *ErrorsResumeAddEditErrors) GetRequestIdOk() (*string, bool)`

GetRequestIdOk returns a tuple with the RequestId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestId

`func (o *ErrorsResumeAddEditErrors) SetRequestId(v string)`

SetRequestId sets RequestId field to given value.


### GetErrors

`func (o *ErrorsResumeAddEditErrors) GetErrors() []ErrorsResumeAddEditError`

GetErrors returns the Errors field if non-nil, zero value otherwise.

### GetErrorsOk

`func (o *ErrorsResumeAddEditErrors) GetErrorsOk() (*[]ErrorsResumeAddEditError, bool)`

GetErrorsOk returns a tuple with the Errors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrors

`func (o *ErrorsResumeAddEditErrors) SetErrors(v []ErrorsResumeAddEditError)`

SetErrors sets Errors field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


