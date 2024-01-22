# EmployerManagersEmployerManagerLimits

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Left** | [**EmployerManagersResumeView**](EmployerManagersResumeView.md) | Количество оставшихся просмотров резюме. В этом параметре содержится лимит просмотров на компанию. Из-за этого он может быть меньше, чем доступно текущему пользователю | 
**Limits** | [**EmployerManagersResumeView**](EmployerManagersResumeView.md) | Лимит на просмотр резюме | 
**Spend** | [**EmployerManagersResumeView**](EmployerManagersResumeView.md) | Количество просмотренных резюме | 

## Methods

### NewEmployerManagersEmployerManagerLimits

`func NewEmployerManagersEmployerManagerLimits(left EmployerManagersResumeView, limits EmployerManagersResumeView, spend EmployerManagersResumeView, ) *EmployerManagersEmployerManagerLimits`

NewEmployerManagersEmployerManagerLimits instantiates a new EmployerManagersEmployerManagerLimits object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEmployerManagersEmployerManagerLimitsWithDefaults

`func NewEmployerManagersEmployerManagerLimitsWithDefaults() *EmployerManagersEmployerManagerLimits`

NewEmployerManagersEmployerManagerLimitsWithDefaults instantiates a new EmployerManagersEmployerManagerLimits object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLeft

`func (o *EmployerManagersEmployerManagerLimits) GetLeft() EmployerManagersResumeView`

GetLeft returns the Left field if non-nil, zero value otherwise.

### GetLeftOk

`func (o *EmployerManagersEmployerManagerLimits) GetLeftOk() (*EmployerManagersResumeView, bool)`

GetLeftOk returns a tuple with the Left field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLeft

`func (o *EmployerManagersEmployerManagerLimits) SetLeft(v EmployerManagersResumeView)`

SetLeft sets Left field to given value.


### GetLimits

`func (o *EmployerManagersEmployerManagerLimits) GetLimits() EmployerManagersResumeView`

GetLimits returns the Limits field if non-nil, zero value otherwise.

### GetLimitsOk

`func (o *EmployerManagersEmployerManagerLimits) GetLimitsOk() (*EmployerManagersResumeView, bool)`

GetLimitsOk returns a tuple with the Limits field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLimits

`func (o *EmployerManagersEmployerManagerLimits) SetLimits(v EmployerManagersResumeView)`

SetLimits sets Limits field to given value.


### GetSpend

`func (o *EmployerManagersEmployerManagerLimits) GetSpend() EmployerManagersResumeView`

GetSpend returns the Spend field if non-nil, zero value otherwise.

### GetSpendOk

`func (o *EmployerManagersEmployerManagerLimits) GetSpendOk() (*EmployerManagersResumeView, bool)`

GetSpendOk returns a tuple with the Spend field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpend

`func (o *EmployerManagersEmployerManagerLimits) SetSpend(v EmployerManagersResumeView)`

SetSpend sets Spend field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


