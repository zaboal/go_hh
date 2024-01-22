# CreateResume400Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**RequestId** | **string** | Идентификатор запроса | 
**Errors** | [**[]ErrorsResumeTooManyResumesError**](ErrorsResumeTooManyResumesError.md) | Массив с данными ошибок | 

## Methods

### NewCreateResume400Response

`func NewCreateResume400Response(requestId string, errors []ErrorsResumeTooManyResumesError, ) *CreateResume400Response`

NewCreateResume400Response instantiates a new CreateResume400Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateResume400ResponseWithDefaults

`func NewCreateResume400ResponseWithDefaults() *CreateResume400Response`

NewCreateResume400ResponseWithDefaults instantiates a new CreateResume400Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRequestId

`func (o *CreateResume400Response) GetRequestId() string`

GetRequestId returns the RequestId field if non-nil, zero value otherwise.

### GetRequestIdOk

`func (o *CreateResume400Response) GetRequestIdOk() (*string, bool)`

GetRequestIdOk returns a tuple with the RequestId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestId

`func (o *CreateResume400Response) SetRequestId(v string)`

SetRequestId sets RequestId field to given value.


### GetErrors

`func (o *CreateResume400Response) GetErrors() []ErrorsResumeTooManyResumesError`

GetErrors returns the Errors field if non-nil, zero value otherwise.

### GetErrorsOk

`func (o *CreateResume400Response) GetErrorsOk() (*[]ErrorsResumeTooManyResumesError, bool)`

GetErrorsOk returns a tuple with the Errors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrors

`func (o *CreateResume400Response) SetErrors(v []ErrorsResumeTooManyResumesError)`

SetErrors sets Errors field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


