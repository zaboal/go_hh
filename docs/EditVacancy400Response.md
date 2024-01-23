# EditVacancy400Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**RequestId** | **string** | Идентификатор запроса | 
**Errors** | [**[]ErrorsVacancyAddEditBadJsonDataError**](ErrorsVacancyAddEditBadJsonDataError.md) | Массив с данными ошибок | 

## Methods

### NewEditVacancy400Response

`func NewEditVacancy400Response(requestId string, errors []ErrorsVacancyAddEditBadJsonDataError, ) *EditVacancy400Response`

NewEditVacancy400Response instantiates a new EditVacancy400Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEditVacancy400ResponseWithDefaults

`func NewEditVacancy400ResponseWithDefaults() *EditVacancy400Response`

NewEditVacancy400ResponseWithDefaults instantiates a new EditVacancy400Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRequestId

`func (o *EditVacancy400Response) GetRequestId() string`

GetRequestId returns the RequestId field if non-nil, zero value otherwise.

### GetRequestIdOk

`func (o *EditVacancy400Response) GetRequestIdOk() (*string, bool)`

GetRequestIdOk returns a tuple with the RequestId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestId

`func (o *EditVacancy400Response) SetRequestId(v string)`

SetRequestId sets RequestId field to given value.


### GetErrors

`func (o *EditVacancy400Response) GetErrors() []ErrorsVacancyAddEditBadJsonDataError`

GetErrors returns the Errors field if non-nil, zero value otherwise.

### GetErrorsOk

`func (o *EditVacancy400Response) GetErrorsOk() (*[]ErrorsVacancyAddEditBadJsonDataError, bool)`

GetErrorsOk returns a tuple with the Errors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrors

`func (o *EditVacancy400Response) SetErrors(v []ErrorsVacancyAddEditBadJsonDataError)`

SetErrors sets Errors field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


