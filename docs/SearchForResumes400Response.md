# SearchForResumes400Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**RequestId** | **string** | Идентификатор запроса | 
**Errors** | [**[]ErrorsCommonBadRequestError**](ErrorsCommonBadRequestError.md) | Массив с данными ошибок | 

## Methods

### NewSearchForResumes400Response

`func NewSearchForResumes400Response(requestId string, errors []ErrorsCommonBadRequestError, ) *SearchForResumes400Response`

NewSearchForResumes400Response instantiates a new SearchForResumes400Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSearchForResumes400ResponseWithDefaults

`func NewSearchForResumes400ResponseWithDefaults() *SearchForResumes400Response`

NewSearchForResumes400ResponseWithDefaults instantiates a new SearchForResumes400Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRequestId

`func (o *SearchForResumes400Response) GetRequestId() string`

GetRequestId returns the RequestId field if non-nil, zero value otherwise.

### GetRequestIdOk

`func (o *SearchForResumes400Response) GetRequestIdOk() (*string, bool)`

GetRequestIdOk returns a tuple with the RequestId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestId

`func (o *SearchForResumes400Response) SetRequestId(v string)`

SetRequestId sets RequestId field to given value.


### GetErrors

`func (o *SearchForResumes400Response) GetErrors() []ErrorsCommonBadRequestError`

GetErrors returns the Errors field if non-nil, zero value otherwise.

### GetErrorsOk

`func (o *SearchForResumes400Response) GetErrorsOk() (*[]ErrorsCommonBadRequestError, bool)`

GetErrorsOk returns a tuple with the Errors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrors

`func (o *SearchForResumes400Response) SetErrors(v []ErrorsCommonBadRequestError)`

SetErrors sets Errors field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


