# ErrorsResumeVisibilityErrorsBadRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**RequestId** | **string** | Идентификатор запроса | 
**Errors** | [**[]ErrorsResumeVisibilityError**](ErrorsResumeVisibilityError.md) | Массив с данными ошибок | 

## Methods

### NewErrorsResumeVisibilityErrorsBadRequest

`func NewErrorsResumeVisibilityErrorsBadRequest(requestId string, errors []ErrorsResumeVisibilityError, ) *ErrorsResumeVisibilityErrorsBadRequest`

NewErrorsResumeVisibilityErrorsBadRequest instantiates a new ErrorsResumeVisibilityErrorsBadRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewErrorsResumeVisibilityErrorsBadRequestWithDefaults

`func NewErrorsResumeVisibilityErrorsBadRequestWithDefaults() *ErrorsResumeVisibilityErrorsBadRequest`

NewErrorsResumeVisibilityErrorsBadRequestWithDefaults instantiates a new ErrorsResumeVisibilityErrorsBadRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRequestId

`func (o *ErrorsResumeVisibilityErrorsBadRequest) GetRequestId() string`

GetRequestId returns the RequestId field if non-nil, zero value otherwise.

### GetRequestIdOk

`func (o *ErrorsResumeVisibilityErrorsBadRequest) GetRequestIdOk() (*string, bool)`

GetRequestIdOk returns a tuple with the RequestId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestId

`func (o *ErrorsResumeVisibilityErrorsBadRequest) SetRequestId(v string)`

SetRequestId sets RequestId field to given value.


### GetErrors

`func (o *ErrorsResumeVisibilityErrorsBadRequest) GetErrors() []ErrorsResumeVisibilityError`

GetErrors returns the Errors field if non-nil, zero value otherwise.

### GetErrorsOk

`func (o *ErrorsResumeVisibilityErrorsBadRequest) GetErrorsOk() (*[]ErrorsResumeVisibilityError, bool)`

GetErrorsOk returns a tuple with the Errors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrors

`func (o *ErrorsResumeVisibilityErrorsBadRequest) SetErrors(v []ErrorsResumeVisibilityError)`

SetErrors sets Errors field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


