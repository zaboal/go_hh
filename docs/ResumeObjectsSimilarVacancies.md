# ResumeObjectsSimilarVacancies

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Counters** | [**ResumeObjectsCounters**](ResumeObjectsCounters.md) |  | 
**Url** | **string** | URL, по которому необходимо сделать GET-запрос, для получения [вакансий, похожих на данное резюме](#tag/Poisk-vakansij-dlya-soiskatelya/operation/get-vacancies-similar-to-resume) | 

## Methods

### NewResumeObjectsSimilarVacancies

`func NewResumeObjectsSimilarVacancies(counters ResumeObjectsCounters, url string, ) *ResumeObjectsSimilarVacancies`

NewResumeObjectsSimilarVacancies instantiates a new ResumeObjectsSimilarVacancies object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewResumeObjectsSimilarVacanciesWithDefaults

`func NewResumeObjectsSimilarVacanciesWithDefaults() *ResumeObjectsSimilarVacancies`

NewResumeObjectsSimilarVacanciesWithDefaults instantiates a new ResumeObjectsSimilarVacancies object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCounters

`func (o *ResumeObjectsSimilarVacancies) GetCounters() ResumeObjectsCounters`

GetCounters returns the Counters field if non-nil, zero value otherwise.

### GetCountersOk

`func (o *ResumeObjectsSimilarVacancies) GetCountersOk() (*ResumeObjectsCounters, bool)`

GetCountersOk returns a tuple with the Counters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCounters

`func (o *ResumeObjectsSimilarVacancies) SetCounters(v ResumeObjectsCounters)`

SetCounters sets Counters field to given value.


### GetUrl

`func (o *ResumeObjectsSimilarVacancies) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *ResumeObjectsSimilarVacancies) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *ResumeObjectsSimilarVacancies) SetUrl(v string)`

SetUrl sets Url field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


