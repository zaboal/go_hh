# ResumeObjectsOwnerComments

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Counters** | [**ResumeObjectsOwnerCommentsCounters**](ResumeObjectsOwnerCommentsCounters.md) | Информация о количестве комментариев | 
**Url** | **string** | URL, на который нужно сделать GET-запрос, чтобы получить список комментариев | 

## Methods

### NewResumeObjectsOwnerComments

`func NewResumeObjectsOwnerComments(counters ResumeObjectsOwnerCommentsCounters, url string, ) *ResumeObjectsOwnerComments`

NewResumeObjectsOwnerComments instantiates a new ResumeObjectsOwnerComments object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewResumeObjectsOwnerCommentsWithDefaults

`func NewResumeObjectsOwnerCommentsWithDefaults() *ResumeObjectsOwnerComments`

NewResumeObjectsOwnerCommentsWithDefaults instantiates a new ResumeObjectsOwnerComments object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCounters

`func (o *ResumeObjectsOwnerComments) GetCounters() ResumeObjectsOwnerCommentsCounters`

GetCounters returns the Counters field if non-nil, zero value otherwise.

### GetCountersOk

`func (o *ResumeObjectsOwnerComments) GetCountersOk() (*ResumeObjectsOwnerCommentsCounters, bool)`

GetCountersOk returns a tuple with the Counters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCounters

`func (o *ResumeObjectsOwnerComments) SetCounters(v ResumeObjectsOwnerCommentsCounters)`

SetCounters sets Counters field to given value.


### GetUrl

`func (o *ResumeObjectsOwnerComments) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *ResumeObjectsOwnerComments) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *ResumeObjectsOwnerComments) SetUrl(v string)`

SetUrl sets Url field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


