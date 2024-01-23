# ErrorsArtifactUploadBadJsonData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**RequestId** | **string** | Идентификатор запроса | 
**Errors** | [**[]ErrorsCommonBadJsonDataError**](ErrorsCommonBadJsonDataError.md) | Массив с данными ошибок | 

## Methods

### NewErrorsArtifactUploadBadJsonData

`func NewErrorsArtifactUploadBadJsonData(requestId string, errors []ErrorsCommonBadJsonDataError, ) *ErrorsArtifactUploadBadJsonData`

NewErrorsArtifactUploadBadJsonData instantiates a new ErrorsArtifactUploadBadJsonData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewErrorsArtifactUploadBadJsonDataWithDefaults

`func NewErrorsArtifactUploadBadJsonDataWithDefaults() *ErrorsArtifactUploadBadJsonData`

NewErrorsArtifactUploadBadJsonDataWithDefaults instantiates a new ErrorsArtifactUploadBadJsonData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRequestId

`func (o *ErrorsArtifactUploadBadJsonData) GetRequestId() string`

GetRequestId returns the RequestId field if non-nil, zero value otherwise.

### GetRequestIdOk

`func (o *ErrorsArtifactUploadBadJsonData) GetRequestIdOk() (*string, bool)`

GetRequestIdOk returns a tuple with the RequestId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestId

`func (o *ErrorsArtifactUploadBadJsonData) SetRequestId(v string)`

SetRequestId sets RequestId field to given value.


### GetErrors

`func (o *ErrorsArtifactUploadBadJsonData) GetErrors() []ErrorsCommonBadJsonDataError`

GetErrors returns the Errors field if non-nil, zero value otherwise.

### GetErrorsOk

`func (o *ErrorsArtifactUploadBadJsonData) GetErrorsOk() (*[]ErrorsCommonBadJsonDataError, bool)`

GetErrorsOk returns a tuple with the Errors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrors

`func (o *ErrorsArtifactUploadBadJsonData) SetErrors(v []ErrorsCommonBadJsonDataError)`

SetErrors sets Errors field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


