# ArtifactsCounters

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Max** | **float32** | Максимально возможное количество артефактов данного типа, которое можно загрузить | 
**Uploaded** | **float32** | Количество уже загруженных артефактов данного типа | 

## Methods

### NewArtifactsCounters

`func NewArtifactsCounters(max float32, uploaded float32, ) *ArtifactsCounters`

NewArtifactsCounters instantiates a new ArtifactsCounters object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewArtifactsCountersWithDefaults

`func NewArtifactsCountersWithDefaults() *ArtifactsCounters`

NewArtifactsCountersWithDefaults instantiates a new ArtifactsCounters object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMax

`func (o *ArtifactsCounters) GetMax() float32`

GetMax returns the Max field if non-nil, zero value otherwise.

### GetMaxOk

`func (o *ArtifactsCounters) GetMaxOk() (*float32, bool)`

GetMaxOk returns a tuple with the Max field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMax

`func (o *ArtifactsCounters) SetMax(v float32)`

SetMax sets Max field to given value.


### GetUploaded

`func (o *ArtifactsCounters) GetUploaded() float32`

GetUploaded returns the Uploaded field if non-nil, zero value otherwise.

### GetUploadedOk

`func (o *ArtifactsCounters) GetUploadedOk() (*float32, bool)`

GetUploadedOk returns a tuple with the Uploaded field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUploaded

`func (o *ArtifactsCounters) SetUploaded(v float32)`

SetUploaded sets Uploaded field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


