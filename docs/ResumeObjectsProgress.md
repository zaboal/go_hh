# ResumeObjectsProgress

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Mandatory** | [**[]IncludesIdName**](IncludesIdName.md) | Список полей, которые обязательны, но еще не заполнены | 
**Percentage** | **float32** | Процент заполненности резюме | 
**Recommended** | [**[]IncludesIdName**](IncludesIdName.md) | Список полей, которые рекомендованы к заполнению, но ещё не заполнены | 

## Methods

### NewResumeObjectsProgress

`func NewResumeObjectsProgress(mandatory []IncludesIdName, percentage float32, recommended []IncludesIdName, ) *ResumeObjectsProgress`

NewResumeObjectsProgress instantiates a new ResumeObjectsProgress object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewResumeObjectsProgressWithDefaults

`func NewResumeObjectsProgressWithDefaults() *ResumeObjectsProgress`

NewResumeObjectsProgressWithDefaults instantiates a new ResumeObjectsProgress object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMandatory

`func (o *ResumeObjectsProgress) GetMandatory() []IncludesIdName`

GetMandatory returns the Mandatory field if non-nil, zero value otherwise.

### GetMandatoryOk

`func (o *ResumeObjectsProgress) GetMandatoryOk() (*[]IncludesIdName, bool)`

GetMandatoryOk returns a tuple with the Mandatory field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMandatory

`func (o *ResumeObjectsProgress) SetMandatory(v []IncludesIdName)`

SetMandatory sets Mandatory field to given value.


### GetPercentage

`func (o *ResumeObjectsProgress) GetPercentage() float32`

GetPercentage returns the Percentage field if non-nil, zero value otherwise.

### GetPercentageOk

`func (o *ResumeObjectsProgress) GetPercentageOk() (*float32, bool)`

GetPercentageOk returns a tuple with the Percentage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPercentage

`func (o *ResumeObjectsProgress) SetPercentage(v float32)`

SetPercentage sets Percentage field to given value.


### GetRecommended

`func (o *ResumeObjectsProgress) GetRecommended() []IncludesIdName`

GetRecommended returns the Recommended field if non-nil, zero value otherwise.

### GetRecommendedOk

`func (o *ResumeObjectsProgress) GetRecommendedOk() (*[]IncludesIdName, bool)`

GetRecommendedOk returns a tuple with the Recommended field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecommended

`func (o *ResumeObjectsProgress) SetRecommended(v []IncludesIdName)`

SetRecommended sets Recommended field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


