# ResumeApplicantFields

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Access** | [**ResumeObjectsAccess**](ResumeObjectsAccess.md) |  | 
**Actions** | [**ResumeObjectsActionsForOwner**](ResumeObjectsActionsForOwner.md) | Дополнительные действия | 
**NewViews** | **float32** | Число новых просмотров. Данный счетчик сбрасывается при получении [детальной истории просмотров](#tag/Rezyume.-Prosmotr-informacii/operation/get-resume-view-history)  | 
**NextPublishAt** | Pointer to **NullableString** | Дата и время следующей возможной публикации/обновления. Для неопубликованных резюме возвращается &#x60;null&#x60; | [optional] 
**PaidServices** | [**[]ResumeObjectsPaidServices**](ResumeObjectsPaidServices.md) | Платные услуги по резюме для автора | 
**TotalViews** | **float32** | Число просмотров резюме | 
**ViewsUrl** | **string** | URL, по которому необходимо сделать GET-запрос для получения [детальной истории просмотров](#tag/Rezyume.-Prosmotr-informacii/operation/get-resume-view-history)  | 

## Methods

### NewResumeApplicantFields

`func NewResumeApplicantFields(access ResumeObjectsAccess, actions ResumeObjectsActionsForOwner, newViews float32, paidServices []ResumeObjectsPaidServices, totalViews float32, viewsUrl string, ) *ResumeApplicantFields`

NewResumeApplicantFields instantiates a new ResumeApplicantFields object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewResumeApplicantFieldsWithDefaults

`func NewResumeApplicantFieldsWithDefaults() *ResumeApplicantFields`

NewResumeApplicantFieldsWithDefaults instantiates a new ResumeApplicantFields object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccess

`func (o *ResumeApplicantFields) GetAccess() ResumeObjectsAccess`

GetAccess returns the Access field if non-nil, zero value otherwise.

### GetAccessOk

`func (o *ResumeApplicantFields) GetAccessOk() (*ResumeObjectsAccess, bool)`

GetAccessOk returns a tuple with the Access field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccess

`func (o *ResumeApplicantFields) SetAccess(v ResumeObjectsAccess)`

SetAccess sets Access field to given value.


### GetActions

`func (o *ResumeApplicantFields) GetActions() ResumeObjectsActionsForOwner`

GetActions returns the Actions field if non-nil, zero value otherwise.

### GetActionsOk

`func (o *ResumeApplicantFields) GetActionsOk() (*ResumeObjectsActionsForOwner, bool)`

GetActionsOk returns a tuple with the Actions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActions

`func (o *ResumeApplicantFields) SetActions(v ResumeObjectsActionsForOwner)`

SetActions sets Actions field to given value.


### GetNewViews

`func (o *ResumeApplicantFields) GetNewViews() float32`

GetNewViews returns the NewViews field if non-nil, zero value otherwise.

### GetNewViewsOk

`func (o *ResumeApplicantFields) GetNewViewsOk() (*float32, bool)`

GetNewViewsOk returns a tuple with the NewViews field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNewViews

`func (o *ResumeApplicantFields) SetNewViews(v float32)`

SetNewViews sets NewViews field to given value.


### GetNextPublishAt

`func (o *ResumeApplicantFields) GetNextPublishAt() string`

GetNextPublishAt returns the NextPublishAt field if non-nil, zero value otherwise.

### GetNextPublishAtOk

`func (o *ResumeApplicantFields) GetNextPublishAtOk() (*string, bool)`

GetNextPublishAtOk returns a tuple with the NextPublishAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNextPublishAt

`func (o *ResumeApplicantFields) SetNextPublishAt(v string)`

SetNextPublishAt sets NextPublishAt field to given value.

### HasNextPublishAt

`func (o *ResumeApplicantFields) HasNextPublishAt() bool`

HasNextPublishAt returns a boolean if a field has been set.

### SetNextPublishAtNil

`func (o *ResumeApplicantFields) SetNextPublishAtNil(b bool)`

 SetNextPublishAtNil sets the value for NextPublishAt to be an explicit nil

### UnsetNextPublishAt
`func (o *ResumeApplicantFields) UnsetNextPublishAt()`

UnsetNextPublishAt ensures that no value is present for NextPublishAt, not even an explicit nil
### GetPaidServices

`func (o *ResumeApplicantFields) GetPaidServices() []ResumeObjectsPaidServices`

GetPaidServices returns the PaidServices field if non-nil, zero value otherwise.

### GetPaidServicesOk

`func (o *ResumeApplicantFields) GetPaidServicesOk() (*[]ResumeObjectsPaidServices, bool)`

GetPaidServicesOk returns a tuple with the PaidServices field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaidServices

`func (o *ResumeApplicantFields) SetPaidServices(v []ResumeObjectsPaidServices)`

SetPaidServices sets PaidServices field to given value.


### GetTotalViews

`func (o *ResumeApplicantFields) GetTotalViews() float32`

GetTotalViews returns the TotalViews field if non-nil, zero value otherwise.

### GetTotalViewsOk

`func (o *ResumeApplicantFields) GetTotalViewsOk() (*float32, bool)`

GetTotalViewsOk returns a tuple with the TotalViews field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalViews

`func (o *ResumeApplicantFields) SetTotalViews(v float32)`

SetTotalViews sets TotalViews field to given value.


### GetViewsUrl

`func (o *ResumeApplicantFields) GetViewsUrl() string`

GetViewsUrl returns the ViewsUrl field if non-nil, zero value otherwise.

### GetViewsUrlOk

`func (o *ResumeApplicantFields) GetViewsUrlOk() (*string, bool)`

GetViewsUrlOk returns a tuple with the ViewsUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetViewsUrl

`func (o *ResumeApplicantFields) SetViewsUrl(v string)`

SetViewsUrl sets ViewsUrl field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


