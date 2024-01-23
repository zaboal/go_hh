# EmployerManagersResumeView

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ResumeView** | **int32** | Просмотры резюме через сайт | 
**ResumeViewFromApi** | Pointer to **int32** | Просмотры резюме из API. Возвращаются, если они предусмотрены в текущей конфигурации доступа к резюме для работодателя | [optional] 

## Methods

### NewEmployerManagersResumeView

`func NewEmployerManagersResumeView(resumeView int32, ) *EmployerManagersResumeView`

NewEmployerManagersResumeView instantiates a new EmployerManagersResumeView object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEmployerManagersResumeViewWithDefaults

`func NewEmployerManagersResumeViewWithDefaults() *EmployerManagersResumeView`

NewEmployerManagersResumeViewWithDefaults instantiates a new EmployerManagersResumeView object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetResumeView

`func (o *EmployerManagersResumeView) GetResumeView() int32`

GetResumeView returns the ResumeView field if non-nil, zero value otherwise.

### GetResumeViewOk

`func (o *EmployerManagersResumeView) GetResumeViewOk() (*int32, bool)`

GetResumeViewOk returns a tuple with the ResumeView field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResumeView

`func (o *EmployerManagersResumeView) SetResumeView(v int32)`

SetResumeView sets ResumeView field to given value.


### GetResumeViewFromApi

`func (o *EmployerManagersResumeView) GetResumeViewFromApi() int32`

GetResumeViewFromApi returns the ResumeViewFromApi field if non-nil, zero value otherwise.

### GetResumeViewFromApiOk

`func (o *EmployerManagersResumeView) GetResumeViewFromApiOk() (*int32, bool)`

GetResumeViewFromApiOk returns a tuple with the ResumeViewFromApi field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResumeViewFromApi

`func (o *EmployerManagersResumeView) SetResumeViewFromApi(v int32)`

SetResumeViewFromApi sets ResumeViewFromApi field to given value.

### HasResumeViewFromApi

`func (o *EmployerManagersResumeView) HasResumeViewFromApi() bool`

HasResumeViewFromApi returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


