# ArtifactsState

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | Идентификатор текущего статуса изображения:  - &#x60;processing&#x60; — в процессе обработки; - &#x60;failed&#x60; — ошибка, скорее всего неподдерживаемый формат; - &#x60;ok&#x60; — обработан, доступен для использования в резюме  | 
**Name** | **string** | Название текущего статуса изображения | 

## Methods

### NewArtifactsState

`func NewArtifactsState(id string, name string, ) *ArtifactsState`

NewArtifactsState instantiates a new ArtifactsState object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewArtifactsStateWithDefaults

`func NewArtifactsStateWithDefaults() *ArtifactsState`

NewArtifactsStateWithDefaults instantiates a new ArtifactsState object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ArtifactsState) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ArtifactsState) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ArtifactsState) SetId(v string)`

SetId sets Id field to given value.


### GetName

`func (o *ArtifactsState) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ArtifactsState) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ArtifactsState) SetName(v string)`

SetName sets Name field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


